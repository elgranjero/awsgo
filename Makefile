SHELL := /bin/bash

ROOT := $(abspath .)
HOME_DIR := $(HOME)
GOIMPORTS ?= $(shell command -v goimports 2>/dev/null)
CACHE_ROOT ?= $(ROOT)
GOCACHE ?= $(CACHE_ROOT)/.gocache
GOMODCACHE ?= $(CACHE_ROOT)/.gomodcache
GO := GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go
GO_BUILD_FLAGS ?= -trimpath
GO_LDFLAGS ?= -s -w
BIN_DIR ?= $(ROOT)/bin
SERVICE_BIN_DIR ?= $(BIN_DIR)/awsgo-services
SERVICE_MANIFEST ?= $(SERVICE_BIN_DIR)/manifest.txt
DISPATCHER_BIN ?= $(BIN_DIR)/awsgo
MONOLITH_BIN ?= $(BIN_DIR)/awsgo-monolith
LEAN_BIN ?= $(BIN_DIR)/awsgo-lean
LEAN_SERVICE_BIN_DIR ?= $(BIN_DIR)/awsgo-lean-services
LEAN_SERVICE_MANIFEST ?= $(LEAN_SERVICE_BIN_DIR)/manifest.txt
SDK_ROOT ?= $(HOME_DIR)/git/aws-sdk-go-v2/service
OUT_ROOT ?= $(ROOT)/generated
JOBS ?= 16
SERVICE_FIND = find awsgo/services -mindepth 1 -maxdepth 1 -type d -exec basename {} \; | sort
LEAN_SERVICE_FIND = find awsgo/lean-services -mindepth 1 -maxdepth 1 -type d -exec basename {} \; | sort
SPLIT_SERVICES ?= all
LEAN_SERVICES ?= all
SPLIT_BUILD_JOBS ?= 2
SPLIT_GOMAXPROCS ?= 1
SAFE_JOBS ?= 4
SAFE_PROCS ?= 2
SAFE_TEST_P ?= 2
SAFE_TEST_TIMEOUT ?= 45m
SAFE_LIVE_TIMEOUT ?= 2h
BENCH_PROFILE ?= default
BENCH_REGION ?= us-east-1
BENCH_REPEAT ?= 3
BENCH_CASES ?=
BENCH_SERVICES ?= acm backup cloudfront cloudtrail cloudwatch configservice costexplorer datasync dynamodb ec2 ecr ecs efs elasticache elasticloadbalancing elasticloadbalancingv2 firehose glacier glue guardduty inspector inspector2 kms lambda opensearch rds route53 s3 secretsmanager securityhub servicecatalog sfn sns sqs sts wafv2
LEAN_BENCH_CASES ?= sts_get_caller_identity ec2_describe_instances s3_list_buckets
BENCH_STARTUP_SUITE ?= lean
BENCH_STARTUP_COMPARE ?= aws
BENCH_STARTUP_REPEAT ?= 10
BENCH_STARTUP_WORKERS ?= 1
BENCH_STARTUP_REST_MS ?= 100
BENCH_STARTUP_TIMEOUT ?= 15

.PHONY: help list-services regen regen-service regen-safe gen fmt goimports install-tools tidy build build-monolith build-dispatcher build-service build-service-bin build-split build-split-all build-lean build-lean-dispatcher build-lean-service build-lean-service-bin build-safe test test-safe test-parity-offline test-parity-offline-safe test-parity-live test-parity-live-safe test-lean-full-live test-lean-help-all parity-safe bench-readonly bench-lean-readonly bench-readonly-dry-run bench-startup-local bench-startup-local-dry-run bench-lean-startup-local bench-lean-vs-full-startup-local bench-lean-vs-full-startup-local-dry-run bench-services size-report scrub-history scrub-history-check clean clean-cache clean-all check all

help:
	@echo "Targets:"
	@echo "  regen         Regenerate all service output under ./service using servicegen"
	@echo "  regen-service Regenerate one service (SERVICE=ec2)"
	@echo "                Optional: SDK_ROOT=/path/to/aws-sdk-go-v2/service OUT_ROOT=/path JOBS=16"
	@echo "  gen           Regenerate awsgo command registry"
	@echo "  tidy          Run go mod tidy"
	@echo "  fmt           Run gofmt"
	@echo "  goimports     Run goimports"
	@echo "  build         Build split awsgo dispatcher + all service binaries + lean pilot"
	@echo "  build-monolith Build legacy monolithic binary at bin/awsgo-monolith"
	@echo "  build-split   Build split dispatcher + services (default SPLIT_SERVICES=all)"
	@echo "  build-split-all Build split dispatcher + every service binary"
	@echo "  build-service Build one service binary (SERVICE=ec2)"
	@echo "  build-lean    Build lean dispatcher + services (default LEAN_SERVICES=all)"
	@echo "  list-services Print all split service names"
	@echo "  test          Run tests under ./tests"
	@echo "  test-parity-offline Run offline CLI parity tests"
	@echo "  test-parity-live    Run live AWS parity tests (opt-in)"
	@echo "  test-lean-full-live Run live lean-vs-full output parity tests (opt-in)"
	@echo "  test-lean-help-all  Run every generated operation --help path locally"
	@echo "  parity-safe         CPU/time-capped parity path (offline)"
	@echo "  test-parity-live-safe Live parity with low CPU/time caps"
	@echo "  bench-readonly      One-time non-destructive AWS CLI vs awsgo timing benchmark"
	@echo "  bench-lean-readonly Benchmark lean pilot against AWS CLI (sts/ec2/s3)"
	@echo "  bench-startup-local No-network startup/help benchmark (BENCH_STARTUP_WORKERS=4)"
	@echo "  bench-lean-vs-full-startup-local Compare lean awsgo against full-runtime awsgo"
	@echo "  size-report         Show source/cache/binary disk usage"
	@echo "  scrub-history       Rewrite local Git history before public release"
	@echo "  scrub-history-check Check current/history for private identifier leaks"
	@echo "  clean         Remove build artifacts"
	@echo "  clean-cache   Remove local Go build/module caches"
	@echo "  clean-all     Remove build artifacts and local Go caches (next build is cold/slow)"
	@echo "  check         Regen + gen + tidy + fmt + goimports + build + test"
	@echo "  all           Alias for check"

list-services:
	@$(SERVICE_FIND)

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

build: build-split build-lean

build-monolith: gen
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && $(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(MONOLITH_BIN) ./awsgo

build-dispatcher: gen
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && $(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(DISPATCHER_BIN) ./awsgo/dispatcher

build-service: gen
	$(MAKE) build-service-bin SERVICE="$(SERVICE)"
	@mkdir -p $(SERVICE_BIN_DIR)
	@{ if [ -f "$(SERVICE_MANIFEST)" ]; then cat "$(SERVICE_MANIFEST)"; fi; printf '%s\n' "$(SERVICE)"; } | sed '/^$$/d' | sort -u > "$(SERVICE_MANIFEST).tmp"
	@mv "$(SERVICE_MANIFEST).tmp" "$(SERVICE_MANIFEST)"

build-service-bin:
	test -n "$(SERVICE)"
	mkdir -p $(SERVICE_BIN_DIR)
	cd $(ROOT) && GOMAXPROCS=$(SPLIT_GOMAXPROCS) $(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(SERVICE_BIN_DIR)/awsgo-$(SERVICE) ./awsgo/services/$(SERVICE)

build-split: build-dispatcher
	mkdir -p $(SERVICE_BIN_DIR)
	@test -n "$(SPLIT_SERVICES)"
	@if [ "$(SPLIT_SERVICES)" = "all" ]; then \
		$(SERVICE_FIND); \
	else \
		printf '%s\n' $(SPLIT_SERVICES); \
	fi | sed '/^$$/d' | sort -u > "$(SERVICE_MANIFEST)"
	@cat "$(SERVICE_MANIFEST)" | xargs -I{} -P $(SPLIT_BUILD_JOBS) $(MAKE) --no-print-directory build-service-bin SERVICE={}

build-split-all:
	$(MAKE) build-split SPLIT_SERVICES=all

build-lean: build-lean-dispatcher
	mkdir -p $(LEAN_SERVICE_BIN_DIR)
	@test -n "$(LEAN_SERVICES)"
	@if [ "$(LEAN_SERVICES)" = "all" ]; then \
		$(LEAN_SERVICE_FIND); \
	else \
		printf '%s\n' $(LEAN_SERVICES); \
	fi | sed '/^$$/d' | sort -u > "$(LEAN_SERVICE_MANIFEST)"
	@cat "$(LEAN_SERVICE_MANIFEST)" | xargs -I{} -P $(SPLIT_BUILD_JOBS) $(MAKE) --no-print-directory build-lean-service-bin SERVICE={}

build-lean-dispatcher:
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && $(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(LEAN_BIN) ./awsgo/lean-dispatcher

build-lean-service:
	test -n "$(SERVICE)"
	$(MAKE) build-lean-dispatcher
	$(MAKE) build-lean-service-bin SERVICE="$(SERVICE)"

build-lean-service-bin:
	test -n "$(SERVICE)"
	mkdir -p $(LEAN_SERVICE_BIN_DIR)
	cd $(ROOT) && GOMAXPROCS=$(SPLIT_GOMAXPROCS) $(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(LEAN_SERVICE_BIN_DIR)/awsgo-lean-$(SERVICE) ./awsgo/lean-services/$(SERVICE)

build-safe: gen
	mkdir -p $(BIN_DIR)
	cd $(ROOT) && GOMAXPROCS=$(SAFE_PROCS) $(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(DISPATCHER_BIN) ./awsgo/dispatcher

test:
	cd $(ROOT) && $(GO) test ./awsgo/dispatcher ./awsgo/leanruntime ./awsgo/svcruntime ./tests/...

test-safe:
	cd $(ROOT) && GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_TEST_TIMEOUT) ./awsgo/cmd ./awsgo/dispatcher ./awsgo/leanruntime ./awsgo/svcruntime ./tests/...

test-parity-offline:
	cd $(ROOT) && $(GO) test ./tests/... -run 'TestParityHighValueServicesPresentInManifest|TestParityHelpFlagNamesAcronymsAndRequiredOptional|TestParityHelpHasOperationFlagsSection'

test-parity-offline-safe:
	cd $(ROOT) && GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_TEST_TIMEOUT) ./tests/... -run 'TestParityHighValueServicesPresentInManifest|TestParityHelpFlagNamesAcronymsAndRequiredOptional|TestParityHelpHasOperationFlagsSection'

test-parity-live:
	cd $(ROOT) && AWSGO_PARITY_LIVE=1 $(GO) test ./tests/... -run TestParityLiveReadOnlyHighValue -count=1 -v

test-parity-live-safe:
	cd $(ROOT) && AWSGO_PARITY_LIVE=1 GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_LIVE_TIMEOUT) ./tests/... -run TestParityLiveReadOnlyHighValue -count=1 -v

test-lean-full-live:
	cd $(ROOT) && AWSGO_LEAN_FULL_LIVE=1 GOMAXPROCS=$(SAFE_PROCS) $(GO) test -p $(SAFE_TEST_P) -timeout $(SAFE_LIVE_TIMEOUT) ./tests/... -run TestLeanFullLiveReadOnlyParity -count=1 -v

test-lean-help-all:
	cd $(ROOT) && AWSGO_LEAN_HELP_ALL=1 GOMAXPROCS=1 $(GO) test -p 1 -timeout $(SAFE_TEST_TIMEOUT) ./tests/... -run TestLeanHelpAllGeneratedOperations -count=1 -v

parity-safe: regen-safe gen build-safe test-safe test-parity-offline-safe

bench-services:
	@printf '%s\n' $(BENCH_SERVICES)

bench-readonly-dry-run:
	@PROFILE="$(BENCH_PROFILE)" REGION="$(BENCH_REGION)" REPEAT="$(BENCH_REPEAT)" BENCH_CASES="$(BENCH_CASES)" AWSGO_BIN="$(DISPATCHER_BIN)" ./scripts/bench-readonly.sh --dry-run

bench-readonly:
	$(MAKE) build-split SPLIT_SERVICES="$(BENCH_SERVICES)" SPLIT_BUILD_JOBS="$(SPLIT_BUILD_JOBS)"
	@PROFILE="$(BENCH_PROFILE)" REGION="$(BENCH_REGION)" REPEAT="$(BENCH_REPEAT)" BENCH_CASES="$(BENCH_CASES)" AWSGO_BIN="$(DISPATCHER_BIN)" ./scripts/bench-readonly.sh

bench-lean-readonly:
	$(MAKE) build-lean
	@PROFILE="$(BENCH_PROFILE)" REGION="$(BENCH_REGION)" REPEAT="$(BENCH_REPEAT)" BENCH_CASES="$(LEAN_BENCH_CASES)" AWSGO_BIN="$(LEAN_BIN)" ./scripts/bench-readonly.sh

bench-startup-local-dry-run:
	@$(GO) run ./awsgo/benchstartup \
		--suite "$(BENCH_STARTUP_SUITE)" \
		--compare "$(BENCH_STARTUP_COMPARE)" \
		--repeat "$(BENCH_STARTUP_REPEAT)" \
		--workers "$(BENCH_STARTUP_WORKERS)" \
		--rest-ms "$(BENCH_STARTUP_REST_MS)" \
		--timeout "$(BENCH_STARTUP_TIMEOUT)" \
		--awsgo-bin "$(DISPATCHER_BIN)" \
		--dry-run

bench-startup-local: build-dispatcher build-lean
	@$(GO) run ./awsgo/benchstartup \
		--suite "$(BENCH_STARTUP_SUITE)" \
		--compare "$(BENCH_STARTUP_COMPARE)" \
		--repeat "$(BENCH_STARTUP_REPEAT)" \
		--workers "$(BENCH_STARTUP_WORKERS)" \
		--rest-ms "$(BENCH_STARTUP_REST_MS)" \
		--timeout "$(BENCH_STARTUP_TIMEOUT)" \
		--awsgo-bin "$(DISPATCHER_BIN)"

bench-lean-startup-local:
	$(MAKE) bench-startup-local BENCH_STARTUP_SUITE=lean

bench-lean-vs-full-startup-local-dry-run:
	$(MAKE) bench-startup-local-dry-run BENCH_STARTUP_SUITE=lean BENCH_STARTUP_COMPARE=full

bench-lean-vs-full-startup-local:
	$(MAKE) bench-startup-local BENCH_STARTUP_SUITE=lean BENCH_STARTUP_COMPARE=full

size-report:
	@du -sh $(ROOT) $(BIN_DIR) $(SERVICE_BIN_DIR) $(LEAN_SERVICE_BIN_DIR) $(GOCACHE) $(GOMODCACHE) $(ROOT)/generated $(ROOT)/awsgo 2>/dev/null || true
	@if [ -d "$(SERVICE_BIN_DIR)" ]; then \
		find "$(SERVICE_BIN_DIR)" -maxdepth 1 -type f -exec stat -f '%z %N' {} \; 2>/dev/null | \
			awk '{sum+=$$1; count++; if($$1>max) max=$$1} END {if(count>0) printf "service-binaries count=%d total=%.1fM avg=%.1fM max=%.1fM\n", count, sum/1024/1024, (sum/count)/1024/1024, max/1024/1024}'; \
		find "$(SERVICE_BIN_DIR)" -maxdepth 1 -type f -exec stat -f '%z %N' {} \; 2>/dev/null | \
			sort -nr | head -n 15 | awk '{size=$$1; $$1=""; sub(/^ /, ""); printf "%.1fM %s\n", size/1024/1024, $$0}'; \
	fi
	@if [ -d "$(LEAN_SERVICE_BIN_DIR)" ]; then \
		find "$(LEAN_SERVICE_BIN_DIR)" -maxdepth 1 -type f -exec stat -f '%z %N' {} \; 2>/dev/null | \
			awk '{sum+=$$1; count++; if($$1>max) max=$$1} END {if(count>0) printf "lean-service-binaries count=%d total=%.1fM avg=%.1fM max=%.1fM\n", count, sum/1024/1024, (sum/count)/1024/1024, max/1024/1024}'; \
		find "$(LEAN_SERVICE_BIN_DIR)" -maxdepth 1 -type f -exec stat -f '%z %N' {} \; 2>/dev/null | \
			sort -nr | head -n 15 | awk '{size=$$1; $$1=""; sub(/^ /, ""); printf "%.1fM %s\n", size/1024/1024, $$0}'; \
	fi

scrub-history:
	./scripts/scrub-private-history.sh --checkpoint "Make lean runtime cover all generated AWS services"

scrub-history-check:
	./scripts/scrub-private-history.sh --check-only

clean:
	rm -rf $(BIN_DIR)
	find $(ROOT) -type f \( -name '*.test' -o -name 'coverage.out' -o -name '*.prof' \) -delete

clean-cache:
	if [ -d "$(GOCACHE)" ]; then chmod -R u+w "$(GOCACHE)"; fi
	if [ -d "$(GOMODCACHE)" ]; then chmod -R u+w "$(GOMODCACHE)"; fi
	rm -rf $(GOCACHE) $(GOMODCACHE)

clean-all: clean clean-cache

check: regen gen tidy fmt goimports build test

all: check
