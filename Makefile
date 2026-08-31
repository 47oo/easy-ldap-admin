# Binary name, taken from the module name
BINARY   := ela
BIN_DIR  := bin

# Build metadata injected via -ldflags
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)

GO       ?= go
LDFLAGS  := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT)

# Cross-compile targets for `make build-all`
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

.DEFAULT_GOAL := help

## build: Build the binary into ./bin/
.PHONY: build
build:
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY) .

## build-all: Cross-compile release binaries into ./bin/
.PHONY: build-all
build-all: $(PLATFORMS)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@mkdir -p $(BIN_DIR)
	@GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) \
		$(GO) build -trimpath -ldflags "$(LDFLAGS)" \
		-o $(BIN_DIR)/$(BINARY)-$(word 1, $(subst /, ,$@))-$(word 2, $(subst /, ,$@))$(if $(findstring windows,$@),.exe) .

## test: Run all tests
.PHONY: test
test:
	$(GO) test ./...

## vet: Run go vet
.PHONY: vet
vet:
	$(GO) vet ./...

## fmt: Format all Go code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

## clean: Remove build artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

## help: Show available targets
.PHONY: help
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  make /'
