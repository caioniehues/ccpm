VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS := -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all build clean test install cross-compile dev lint fmt

all: build

build:
	go build $(LDFLAGS) -o ccpm-tui ./cmd/ccpm-tui

install:
	go install $(LDFLAGS) ./cmd/ccpm-tui

test:
	go test -v -cover ./internal/...

clean:
	rm -f ccpm-tui
	rm -rf dist/

cross-compile: clean
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/ccpm-tui-linux-amd64 ./cmd/ccpm-tui
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/ccpm-tui-linux-arm64 ./cmd/ccpm-tui
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/ccpm-tui-darwin-amd64 ./cmd/ccpm-tui
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/ccpm-tui-darwin-arm64 ./cmd/ccpm-tui
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/ccpm-tui-windows-amd64.exe ./cmd/ccpm-tui
	cd dist && gzip -k ccpm-tui-linux-* ccpm-tui-darwin-* 2>/dev/null || true
	cd dist && sha256sum ccpm-tui-* > checksums.txt 2>/dev/null || shasum -a 256 ccpm-tui-* > checksums.txt

dev:
	go build -race -o ccpm-tui ./cmd/ccpm-tui

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...
	gofmt -s -w .

coverage:
	go test -coverprofile=coverage.out ./internal/...
	go tool cover -html=coverage.out -o coverage.html
