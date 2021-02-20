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

## Build and run
dev.run: dev run

## Execute development pipeline
dev: format lint swagger build

## Run
run:
	./bin/conductor

#-------------------------
# Checks
#-------------------------
.PHONY: format lint stats.loc

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
.PHONY: clean clean.conductor

## Clean build files
clean:
	rm -rf $(DIRS_TO_CLEAN)
	rm -f $(FILES_TO_CLEAN)

#-------------------------
# Target: swagger
#-------------------------
.PHONY: swagger

swagger: swagger.gen swagger.validate

## Generate swagger json
swagger.gen:
	swagger generate spec -o ./internal/ui/swagger.json

## Validate swagger
swagger.validate:
	swagger validate ./internal/ui/swagger.json