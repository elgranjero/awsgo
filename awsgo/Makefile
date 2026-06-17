SHELL := /bin/bash

ROOT := $(abspath ..)
GOIMPORTS ?= $(shell command -v goimports 2>/dev/null)
GOCACHE ?= $(ROOT)/.gocache
GOMODCACHE ?= $(ROOT)/.gomodcache
GO := GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go
BIN_DIR ?= $(ROOT)/bin
BIN ?= $(BIN_DIR)/awsgo

.PHONY: help gen tidy fmt goimports install-tools build test check all

help:
	@echo "Targets:"
	@echo "  gen           Generate awsgo registry and adapters"
	@echo "  tidy          Run go mod tidy"
	@echo "  fmt           Run gofmt"
	@echo "  goimports     Run goimports"
	@echo "  install-tools Install goimports"
	@echo "  build         Build awsgo"
	@echo "  test          Run shared test suite under tests/"
	@echo "  check         Run gen + tidy + fmt + goimports + build + test"
	@echo "  all           Alias for check"

gen:
	cd $(ROOT) && $(GO) run ./awsgo/gen

tidy:
	cd $(ROOT) && $(GO) mod tidy

fmt:
	cd $(ROOT) && gofmt -w $$(find awsgo -type f -name '*.go' | sort)

install-tools:
	cd $(ROOT) && $(GO) install golang.org/x/tools/cmd/goimports@latest

goimports:
	@if [ -z "$(GOIMPORTS)" ]; then \
		echo "goimports not found. Installing..."; \
		$(MAKE) install-tools; \
	fi
	cd $(ROOT) && $${GOIMPORTS:-$$(command -v goimports)} -w $$(find awsgo -type f -name '*.go' | sort)

build:
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && $(GO) build -o $(BIN) ./awsgo

test:
	cd $(ROOT) && $(GO) test ./tests

check: gen tidy fmt goimports build test

all: check
