#!/usr/bin/env python3
"""Catalog bookkeeping for table-declared APIs. Run from the repo root.

  catalog.py activate <slug>...                 add slugs to active-apis.json + active_apis_test.go
  catalog.py splice-spec <fragment.yaml>        insert "# TAGS / # PATHS" fragment into spec.yaml
  catalog.py manifest <id> <summary> <reason> <validated> <file>...
                                                append a .printing-press-patches.json entry
"""
import datetime
import json
import re
import sys

ACTIVE = "internal/cli/active-apis.json"
ACTIVE_TEST = "internal/cli/active_apis_test.go"
SPEC = "spec.yaml"
PATCHES = ".printing-press-patches.json"


def insert_in_place(items, add):
    """Insert without reordering existing entries (keeps diffs minimal)."""
    out = list(items)
    for a in add:
        if a in out:
            continue
        i = next((k for k, x in enumerate(out) if x > a), len(out))
        out.insert(i, a)
    return out


def activate(slugs):
    d = json.load(open(ACTIVE))
    d["active"] = insert_in_place(d["active"], slugs)
    open(ACTIVE, "w").write(json.dumps(d, indent=2) + "\n")
    s = open(ACTIVE_TEST).read()
    m = re.search(r"(func TestSupportedWebsiteActiveAPIsAreActive.*?range \[\]string\{\n)(.*?)(\t\} \{)", s, re.S)
    cur = re.findall(r'"([^"]+)"', m.group(2))
    s = s[:m.start(2)] + "".join(f'\t\t"{x}",\n' for x in insert_in_place(cur, slugs)) + s[m.end(2):]
    n = len(d["active"])
    s = re.sub(r'(if got != )\d+( \{\n\t\tt\.Errorf\("activeAPICount\(\) = %d, want )\d+', rf"\g<1>{n}\g<2>{n}", s)
    open(ACTIVE_TEST, "w").write(s)
    print(f"active APIs: {n}")


def group_tags(block):
    out = []
    for ln in block:
        if ln.startswith("  - name: "):
            out.append([ln[len("  - name: "):].strip(), [ln]])
        elif out and ln.strip():
            out[-1][1].append(ln)
    return out


def splice_spec(frag_path):
    lines = open(SPEC).read().split("\n")
    sec, tags, paths = None, [], []
    for ln in open(frag_path).read().split("\n"):
        if ln.strip() in ("# TAGS", "# PATHS"):
            sec = tags if ln.strip() == "# TAGS" else paths
        elif sec is not None:
            sec.append(ln)
    while paths and paths[-1] == "":
        paths.pop()
    t0 = lines.index("tags:") + 1
    t1 = t0
    while t1 < len(lines) and (lines[t1].startswith("  ") or lines[t1] == ""):
        t1 += 1
    existing = group_tags(lines[t0:t1])
    new_tags = [t for t in group_tags(tags) if t[0].lower() not in {n.lower() for n, _ in existing}]
    merged = sorted(existing + new_tags, key=lambda x: x[0].lower())
    lines = lines[:t0] + [l for _, blk in merged for l in blk] + lines[t1:]
    ci = lines.index("components:")
    have = {l.strip() for l in lines[:ci] if l.startswith("  /")}
    for l in paths:
        if l.startswith("  /") and l.strip() in have:
            sys.exit(f"path already in spec.yaml: {l.strip()} (add the missing method by hand)")
    lines = lines[:ci] + paths + lines[ci:]
    open(SPEC, "w").write("\n".join(lines))
    print(f"spec.yaml: +{len(new_tags)} tags, +{sum(l.startswith('  /') for l in paths)} paths")


def manifest(pid, summary, reason, validated, files):
    d = json.load(open(PATCHES))
    d["applied_at"] = datetime.date.today().isoformat()
    d["patches"].append({"id": pid, "summary": summary, "reason": reason, "files": files, "validated_outcome": validated})
    open(PATCHES, "w").write(json.dumps(d, indent=2) + "\n")


if __name__ == "__main__":
    cmd, args = sys.argv[1], sys.argv[2:]
    if cmd == "activate":
        activate(args)
    elif cmd == "splice-spec":
        splice_spec(args[0])
    elif cmd == "manifest":
        manifest(args[0], args[1], args[2], args[3], args[4:])
    else:
        sys.exit(__doc__)
