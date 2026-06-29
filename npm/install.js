#!/usr/bin/env node
// Postinstall: downloads the datpaq binary from the matching GitHub release.
'use strict';

const https = require('https');
const fs = require('fs');
const path = require('path');
const os = require('os');
const { execFileSync } = require('child_process');

const pkg = require('./package.json');
const VERSION = pkg.version;

const PLATFORM_MAP = { win32: 'windows', darwin: 'darwin', linux: 'linux' };
const ARCH_MAP = { x64: 'amd64', arm64: 'arm64' };

function platform() {
  const p = PLATFORM_MAP[process.platform];
  if (!p) throw new Error(`Unsupported platform: ${process.platform}`);
  return p;
}

function arch() {
  const a = ARCH_MAP[process.arch];
  if (!a) throw new Error(`Unsupported architecture: ${process.arch}`);
  return a;
}

function follow(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    function get(u) {
      https.get(u, (res) => {
        if (res.statusCode === 301 || res.statusCode === 302) {
          res.resume();
          return get(res.headers.location);
        }
        if (res.statusCode !== 200) {
          file.close();
          fs.unlink(dest, () => {});
          return reject(new Error(`HTTP ${res.statusCode} downloading ${u}`));
        }
        res.pipe(file);
        file.on('finish', () => file.close(resolve));
        file.on('error', (err) => { fs.unlink(dest, () => {}); reject(err); });
      }).on('error', (err) => { fs.unlink(dest, () => {}); reject(err); });
    }
    get(url);
  });
}

async function main() {
  const plat = platform();
  const ar = arch();
  const isWin = plat === 'windows';
  const ext = isWin ? '.exe' : '';
  const archiveExt = isWin ? '.zip' : '.tar.gz';
  const archiveName = `datpaq_${VERSION}_${plat}_${ar}${archiveExt}`;
  const url = `https://github.com/datpaq/cli/releases/download/v${VERSION}/${archiveName}`;

  const binDir = path.join(__dirname, 'bin');
  fs.mkdirSync(binDir, { recursive: true });

  const tmpArchive = path.join(os.tmpdir(), archiveName);

  console.log(`Downloading datpaq v${VERSION} (${plat}/${ar})...`);
  await follow(url, tmpArchive);

  const binaryDest = path.join(binDir, `datpaq${ext}`);

  if (isWin) {
    // tar.exe ships with Windows 10 1803+ and handles zip natively
    execFileSync('tar', ['-xf', tmpArchive, '-C', binDir, `datpaq${ext}`], { stdio: 'inherit' });
  } else {
    execFileSync('tar', ['-xzf', tmpArchive, '-C', binDir, `datpaq${ext}`], { stdio: 'inherit' });
    fs.chmodSync(binaryDest, 0o755);
  }

  fs.rmSync(tmpArchive, { force: true });

  if (!fs.existsSync(binaryDest)) {
    throw new Error(`Binary not found after extraction: ${binaryDest}`);
  }

  console.log('datpaq installed successfully. Run: datpaq --version');
}

main().catch((err) => {
  console.error(`\ndatpaq install failed: ${err.message}`);
  console.error('You can also install via Go: go install github.com/datpaq/cli/cmd/datpaq@latest');
  process.exit(1);
});
