.PHONY: build test lint install clean

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
