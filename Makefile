.PHONY: build test lint install clean screenshots

build:
	go build -o bin/datpaq ./cmd/datpaq

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/datpaq

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/datpaq-mcp ./cmd/datpaq-mcp

install-mcp:
	go install ./cmd/datpaq-mcp

build-all: build build-mcp

# Regenerate the CLI hero banner + social-share card from docs/cli-terminal.tape.
# Requires: brew install vhs webp
screenshots:
	./generate-cli-screenshots.sh
