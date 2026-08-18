# 1panel-sdk Makefile
# Common tasks: build, test, lint, codegen.

# Use bash for consistent behaviour across platforms.
SHELL := /usr/bin/env bash

GO          ?= go
GOLANGCI    ?= golangci-lint
BIN_DIR     ?= bin

.PHONY: help build test test-race cover lint fmt vet codegen clean tidy all

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build the SDK and example.
	$(GO) build ./...
	cd example && $(GO) build ./...

test: ## Run all unit tests.
	$(GO) test -count=1 ./...

test-race: ## Run tests with the race detector.
	$(GO) test -count=1 -race ./...

cover: ## Print test coverage to stdout.
	$(GO) test -coverprofile=cover.out -coverpkg=./... ./...
	$(GO) tool cover -func=cover.out | tail -1

lint: ## Run golangci-lint (requires `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`).
	$(GOLANGCI) run ./...

fmt: ## Run gofmt -s and goimports.
	$(GO) fmt ./...
	gofmt -s -w .

vet: ## Run go vet.
	$(GO) vet ./...

codegen: ## Regenerate services_swagger.go from 1Panel's swagger.json.
	python3 scripts/gen-from-swagger.py
	$(GO) build ./...

split: ## Re-split services2.go into per-domain files (one-time).
	python3 scripts/split-services2.py
	$(GO) build ./...

tidy: ## Tidy go.mod files.
	$(GO) mod tidy
	cd example && $(GO) mod tidy

clean: ## Remove build artefacts.
	rm -f cover.out
	rm -rf $(BIN_DIR)

all: build test lint ## Build, test, and lint.
