SHELL := /bin/bash

ROOT := $(abspath .)
HOME_DIR := $(HOME)
GOIMPORTS ?= $(shell command -v goimports 2>/dev/null)
CACHE_ROOT ?= $(ROOT)
GOCACHE ?= $(CACHE_ROOT)/.gocache
GOMODCACHE ?= $(CACHE_ROOT)/.gomodcache
GO := GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go
BIN_DIR ?= $(ROOT)/bin
SERVICE_BIN_DIR ?= $(BIN_DIR)/awsgo-services
SDK_ROOT ?= $(HOME_DIR)/git/aws-sdk-go-v2/service
OUT_ROOT ?= $(ROOT)/generated
JOBS ?= 16
SPLIT_SERVICES ?= sts iam ec2 s3 cloudwatch rds
SAFE_JOBS ?= 4
SAFE_PROCS ?= 2
SAFE_TEST_P ?= 2
SAFE_TEST_TIMEOUT ?= 45m
SAFE_LIVE_TIMEOUT ?= 2h

.PHONY: help regen regen-service regen-safe gen fmt goimports install-tools tidy build build-monolith build-dispatcher build-service build-split build-safe test test-safe test-parity-offline test-parity-offline-safe test-parity-live test-parity-live-safe parity-safe clean clean-cache check all

help:
	@echo "Targets:"
	@echo "  regen         Regenerate all service output under ./service using servicegen"
	@echo "  regen-service Regenerate one service (SERVICE=ec2)"
	@echo "                Optional: SDK_ROOT=/path/to/aws-sdk-go-v2/service OUT_ROOT=/path JOBS=16"
	@echo "  gen           Regenerate awsgo command registry"
	@echo "  tidy          Run go mod tidy"
	@echo "  fmt           Run gofmt"
	@echo "  goimports     Run goimports"
	@echo "  build         Build monolithic awsgo binary"
	@echo "  build-split   Build split dispatcher + selected services (SPLIT_SERVICES=\"ec2 s3\")"
	@echo "  build-service Build one service binary (SERVICE=ec2)"
	@echo "  test          Run tests under ./tests"
	@echo "  test-parity-offline Run offline CLI parity tests"
	@echo "  test-parity-live    Run live AWS parity tests (opt-in)"
	@echo "  parity-safe         CPU/time-capped parity path (offline)"
	@echo "  test-parity-live-safe Live parity with low CPU/time caps"
	@echo "  clean         Remove build artifacts"
	@echo "  clean-cache   Remove local Go build/module caches"
	@echo "  check         Regen + gen + tidy + fmt + goimports + build + test"
	@echo "  all           Alias for check"

regen:
	cd $(ROOT) && $(GO) run ./servicegen.go \
		--sdk-root "$(SDK_ROOT)" \
		--out-root "$(OUT_ROOT)" \
		--jobs "$(JOBS)" \
		--write

regen-service:
	test -n "$(SERVICE)"
	cd $(ROOT) && $(GO) run ./servicegen.go \
		--sdk-root "$(SDK_ROOT)" \
		--out-root "$(OUT_ROOT)" \
		--jobs "$(JOBS)" \
		--services "$(SERVICE)" \
		--write

regen-safe:
	$(MAKE) regen JOBS="$(SAFE_JOBS)"

gen:
	cd $(ROOT) && $(GO) run ./awsgo/gen

tidy:
	cd $(ROOT) && $(GO) mod tidy

fmt:
	cd $(ROOT) && gofmt -w $$(find awsgo generated tests -type f -name '*.go' | sort)

install-tools:
	cd $(ROOT) && $(GO) install golang.org/x/tools/cmd/goimports@latest

goimports:
	@if [ -z "$(GOIMPORTS)" ]; then \
		echo "goimports not found. Installing..."; \
		$(MAKE) install-tools; \
	fi
	cd $(ROOT) && $${GOIMPORTS:-$$(command -v goimports)} -w $$(find awsgo generated tests -type f -name '*.go' | sort)

build: build-monolith

build-monolith: gen
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && $(GO) build -o $(BIN_DIR)/awsgo ./awsgo

build-dispatcher: gen
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && $(GO) build -o $(BIN_DIR)/awsgo-split ./awsgo/dispatcher

build-service: gen
	test -n "$(SERVICE)"
	mkdir -p $(SERVICE_BIN_DIR)
	cd $(ROOT) && $(GO) build -o $(SERVICE_BIN_DIR)/awsgo-$(SERVICE) ./awsgo/services/$(SERVICE)

build-split: build-dispatcher
	mkdir -p $(SERVICE_BIN_DIR)
	@for svc in $(SPLIT_SERVICES); do \
		echo "building split service $$svc"; \
		(cd $(ROOT) && $(GO) build -o $(SERVICE_BIN_DIR)/awsgo-$$svc ./awsgo/services/$$svc) || exit $$?; \
	done

build-safe: gen
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && GOMAXPROCS=$(SAFE_PROCS) $(GO) build -o $(BIN_DIR)/awsgo ./awsgo

test:
	cd $(ROOT) && $(GO) test ./tests/...

test-safe:
	cd $(ROOT) && GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_TEST_TIMEOUT) ./awsgo/cmd ./tests/...

test-parity-offline:
	cd $(ROOT) && $(GO) test ./tests/... -run 'TestParityHighValueServicesPresentInManifest|TestParityHelpFlagNamesAcronymsAndRequiredOptional|TestParityHelpHasOperationFlagsSection'

test-parity-offline-safe:
	cd $(ROOT) && GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_TEST_TIMEOUT) ./tests/... -run 'TestParityHighValueServicesPresentInManifest|TestParityHelpFlagNamesAcronymsAndRequiredOptional|TestParityHelpHasOperationFlagsSection'

test-parity-live:
	cd $(ROOT) && AWSGO_PARITY_LIVE=1 $(GO) test ./tests/... -run TestParityLiveReadOnlyHighValue -count=1 -v

test-parity-live-safe:
	cd $(ROOT) && AWSGO_PARITY_LIVE=1 GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_LIVE_TIMEOUT) ./tests/... -run TestParityLiveReadOnlyHighValue -count=1 -v

parity-safe: regen-safe gen build-safe test-safe test-parity-offline-safe

clean:
	rm -rf $(BIN_DIR)
	find $(ROOT) -type f \( -name '*.test' -o -name 'coverage.out' -o -name '*.prof' \) -delete

clean-cache:
	rm -rf $(GOCACHE) $(GOMODCACHE)

check: regen gen tidy fmt goimports build test

all: check
