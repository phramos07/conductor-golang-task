GO   := go

DIRS_TO_CLEAN:=  ./tmp
FILES_TO_CLEAN:= ./bin/*

ifeq ($(origin GO), undefined)
  GO:=$(shell which go)
endif
ifeq ($(GO),)
  $(error Could not find 'go' in path. Please install go, or if already installed either add it to your path or set GO to point to its directory)
endif

pkgs  = $(shell GOFLAGS=-mod=vendor $(GO) list ./... | grep -vE -e /vendor/ -e /pkg/swagger/)
pkgDirs = $(shell GOFLAGS=-mod=vendor $(GO) list -f {{.Dir}} ./... | grep -vE -e /vendor/ -e /pkg/swagger/)

GOLANGCI:=$(shell command -v golangci-lint 2> /dev/null)

#-------------------------
# Final targets
#-------------------------
.PHONY: dev

## Execute development pipeline
dev: format lint build

#-------------------------
# Checks
#-------------------------
.PHONY: format lint stats.loc

check: format lint

## Apply code format, import reorganization and code simplification on source code
format:
	@echo "==> formatting code"
	@$(GO) fmt $(pkgs)
	@echo "==> clean imports"
	@goimports -w $(pkgDirs)
	@echo "==> simplify code"
	@gofmt -s -w $(pkgDirs)

## Validate code
lint:
ifndef GOLANGCI
	$(error "Please install golangci! make get-tools")
endif
	@golangci-lint run -v $(pkgDirs)

#-------------------------
# Build artefacts
#-------------------------
.PHONY: build build.conductor

## Build all binaries
build:
	$(GO) build -o bin/conductor internal/app.go

## Compress all binaries
pack:
	@echo ">> packing all binaries"
	@upx -7 -qq bin/*

#-------------------------
# Target: clean
#-------------------------

## Clean build files
clean:
	rm -rf $(DIRS_TO_CLEAN)
	rm -f $(FILES_TO_CLEAN)