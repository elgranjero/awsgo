# awsgo

`awsgo` is an experimental AWS CLI-style command-line tool generated from the
AWS SDK for Go v2.

The goal is AWS CLI parity first: familiar command shapes, generated operation
flags, shared config/profile behavior, JMESPath queries, and multiple output
formats. Once parity is solid, the project can start improving on AWS CLI
ergonomics without giving up compatibility.

This is not an official AWS project.

## Status

Pre-release. The CLI is usable for many read/list/describe workflows, but parity
work is still active.

Current focus areas:

- Broad AWS SDK v2 service coverage.
- AWS-style commands: `awsgo <service> <operation> [flags]`.
- Generated operation input flags from SDK input shapes.
- AWS shared config/profile support.
- JMESPath `--query` support.
- Output formats: `json`, `yaml`, `text`, `table`, `csv`, `markdown`, `html`.
- Split service binaries to avoid loading every AWS service into one giant
  process.

## Why Split Binaries?

A single Go binary containing every AWS SDK service is large and expensive to
load. `awsgo` now builds a small dispatcher plus one binary per service:

```text
bin/awsgo
bin/awsgo-services/awsgo-ec2
bin/awsgo-services/awsgo-iam
bin/awsgo-services/awsgo-rds
...
```

The dispatcher finds and executes the service binary only when that service is
called.

## Lean Runtime Pilot

`bin/awsgo` prefers the experimental lean runtime for operations that have a
lean implementation, then falls back to the full split service binary for every
other operation.

```text
bin/awsgo
bin/awsgo-lean-services/awsgo-lean-sts
bin/awsgo-lean-services/awsgo-lean-ec2
bin/awsgo-lean-services/awsgo-lean-s3
```

The lean path skips the generated Cobra command package for each service and
calls selected AWS SDK operations directly through a smaller shared runtime. It
is intentionally narrow today: `sts`, `ec2`, and `s3`, JSON output, and a small
set of read-only operations. It exists to test whether a future generated direct
SDK invoker can replace the heavier generated service command packages.

Build the default dispatcher, full split service binaries, and lean pilot:

```bash
make build
```

Example:

```bash
bin/awsgo ec2 describe-instances --max-results 20 --profile default
```

Force the full split runtime instead of lean:

```bash
AWSGO_DISABLE_LEAN=1 bin/awsgo ec2 describe-instances --max-results 20
```

Current takeaway: the lean approach is a major binary-size win, but only a
modest wall-clock runtime win over the current split architecture because most
live AWS command time is spent waiting on AWS service/network latency.

### Lean Runtime: Tested So Far

The lean path has been tested as a pilot, not as a full replacement yet.

Implemented lean operations:

| Service | Operations |
| --- | --- |
| `sts` | `get-caller-identity` |
| `ec2` | `describe-instances`, `describe-volumes` |
| `s3` | `list-buckets` |

Verified locally:

- `bin/awsgo` routes JSON-compatible supported operations to lean by default.
- `AWSGO_DISABLE_LEAN=1` forces the full split service runtime.
- `--output text` falls back to the full split runtime because lean currently
  supports JSON output only.
- AWS shared config `output = text` also falls back to the full split runtime.
- Unsupported lean operations, for example `iam list-roles`, use the full split
  runtime.
- `make build-dispatcher`, `make build-lean`, and `make test-safe` pass with
  dispatcher, lean runtime, split runtime, and parity tests included.

Live read-only benchmarks have been run for:

- `sts get-caller-identity`
- `ec2 describe-instances`
- `s3 list-buckets`

Not yet tested:

- Generated lean implementations for every AWS service.
- Mutating operations through the lean runtime.
- Full AWS CLI output-format parity directly inside lean.
- Full AWS CLI shorthand parity directly inside lean.

The intended direction is to generate lean service binaries for every service
and make direct SDK invocation the primary path. Until that generated lean path
reaches full parity, the heavier split runtime remains the compatibility
fallback.

## Requirements

- Go 1.24 or newer.
- AWS credentials configured using the usual AWS SDK/AWS CLI mechanisms.
- Optional for regeneration: a local checkout of `aws-sdk-go-v2`.

The default generator expects the SDK checkout here:

```bash
~/git/aws-sdk-go-v2/service
```

You can override that path with `SDK_ROOT`.

## Build

Build the dispatcher, all generated service binaries, and the lean pilot:

```bash
make build
```

Build one service binary:

```bash
make build-service SERVICE=ec2
```

Build a limited split set:

```bash
make build-split SPLIT_SERVICES="sts iam ec2"
```

Limited split builds write `bin/awsgo-services/manifest.txt`; `bin/awsgo
--help` and dispatch only expose services listed in that manifest.

Build the legacy monolithic binary:

```bash
make build-monolith
```

List generated service wrappers:

```bash
make list-services
```

Clean build artifacts and caches:

```bash
make clean
make clean-cache
make clean-all
```

`make clean` removes build artifacts but keeps local Go caches, so rebuilds are
usually much faster. `make clean-all` also removes `.gocache/` and
`.gomodcache/`; use it to reclaim disk space or test cold-cache behavior, but
expect the next build to be substantially slower.

## Usage

Show top-level help:

```bash
bin/awsgo --help
```

Call an operation:

```bash
bin/awsgo sts get-caller-identity
```

Use a profile and region:

```bash
bin/awsgo ec2 describe-instances --profile dev --region us-east-1
```

Use generated operation flags:

```bash
bin/awsgo iam list-roles --max-items 2
```

Use JSON input for complex shapes:

```bash
bin/awsgo cloudwatch get-metric-statistics \
  --profile prod \
  --cli-input-json '{
    "Namespace": "AWS/RDS",
    "MetricName": "VolumeWriteIOPs",
    "Dimensions": [
      {"Name": "DBInstanceIdentifier", "Value": "example-db"}
    ],
    "StartTime": "2026-06-17T00:00:00Z",
    "EndTime": "2026-06-17T01:00:00Z",
    "Period": 3600,
    "Statistics": ["Average"]
  }'
```

Apply a JMESPath query:

```bash
bin/awsgo rds describe-db-instances \
  --query 'DBInstances[*].{Identifier:DBInstanceIdentifier,StorageType:StorageType,Size:AllocatedStorage}'
```

Choose output format:

```bash
bin/awsgo ec2 describe-volumes --output table
bin/awsgo ec2 describe-volumes --output yaml
bin/awsgo ec2 describe-volumes --output markdown
bin/awsgo ec2 describe-volumes --output html
```

`awsgo` honors output defaults in this order:

1. explicit `--output`
2. `AWS_DEFAULT_OUTPUT`
3. selected profile `output` from AWS shared config
4. `json`

For example:

```ini
[default]
region = us-east-1
output = text
```

Then:

```bash
bin/awsgo ec2 describe-instances --profile default
```

uses `text` unless `--output` is provided.

## Input Flags

Operation flags are generated from SDK input fields and converted to
kebab-case:

```bash
bin/awsgo cloudwatch get-metric-statistics \
  --namespace AWS/RDS \
  --metric-name VolumeWriteIOPs \
  --dimensions '[{"Name":"DBInstanceIdentifier","Value":"example-db"}]' \
  --start-time 2026-06-17T00:00:00Z \
  --end-time 2026-06-17T01:00:00Z \
  --period 3600 \
  --statistics '["Average"]'
```

Primitive fields can usually be passed directly. Complex objects, lists, and
maps should be passed as JSON strings or through `--cli-input-json`.

Show operation help:

```bash
bin/awsgo wafv2 create-web-acl --help
```

The help output includes operation flags, inferred types, and required/optional
status when available.

## Regeneration

Regenerate all service output from the AWS SDK checkout:

```bash
make regen
make gen
```

Regenerate one service:

```bash
make regen-service SERVICE=iam
make gen
```

Use a custom SDK checkout:

```bash
make regen SDK_ROOT=/path/to/aws-sdk-go-v2/service
```

Safe, lower-concurrency regeneration:

```bash
make regen-safe
```

## Tests

Run the main test suite:

```bash
make test
```

Run safer bounded tests:

```bash
make test-safe
```

Run offline parity checks:

```bash
make test-parity-offline
```

Live parity tests are opt-in because they can call AWS APIs:

```bash
make test-parity-live
```

Lean-vs-full live parity is also opt-in. It compares the current lean-default
path against the full generated service runtime for the lean read-only pilot
operations:

```bash
make test-lean-full-live
```

## One-Time Benchmarking

The project includes a small, opt-in benchmark harness for comparing AWS CLI
against `awsgo` on non-destructive read/list/describe calls.

Run a dry-run first to see the exact commands:

```bash
make bench-readonly-dry-run BENCH_PROFILE=default BENCH_REGION=us-east-1
```

Run the benchmark:

```bash
make bench-readonly BENCH_PROFILE=default BENCH_REGION=us-east-1 BENCH_REPEAT=3 SPLIT_BUILD_JOBS=8
```

Run a smaller subset:

```bash
make bench-readonly \
  BENCH_CASES="sts_get_caller_identity ec2_describe_instances s3_list_buckets" \
  BENCH_PROFILE=default \
  BENCH_REGION=us-east-1 \
  BENCH_REPEAT=5
```

Run the lean pilot benchmark subset:

```bash
make bench-lean-readonly \
  BENCH_PROFILE=default \
  BENCH_REGION=us-east-1 \
  BENCH_REPEAT=5
```

The harness:

- Uses read-only AWS APIs.
- Disables pagination for both CLIs by default and caps result sizes where the
  underlying API exposes a limit field.
- Is implemented as a Bash harness; it does not use Python or Perl helper
  scripts.
- Redirects command output and records only timing, status, and byte counts.
- Writes local TSV files under `benchmarks/`, which are ignored by git.
- Still uses real AWS API calls, so normal AWS API rate limits and any
  service-specific API pricing still apply.

Use this for a one-time data point before writing performance claims. The useful
question is not “is Go always faster?” but “for common AWS read-only commands,
how much startup/runtime overhead do we avoid compared with the Python AWS CLI?”

### Planned Local Startup Benchmark

The next benchmark should remove AWS service latency entirely and measure local
startup/shutdown overhead only. The proposed local suite compares `aws` and
`awsgo` by running help/parse-only commands that do not make network calls:

- top-level help: `aws help` vs `bin/awsgo --help`
- service help: `aws <service> help` vs `bin/awsgo <service> --help`
- operation help: `aws <service> <operation> help` vs
  `bin/awsgo <service> <operation> --help`

This should be run with a bounded worker count, one command process per worker
at a time, alternating CLI order between repeats, and with a small rest interval
between commands so neither runtime gets an artificial warm-process advantage.
The benchmark runner itself is implemented in Go under `awsgo/benchstartup`;
Python is only present as part of the AWS CLI being measured.

Run the lean local startup suite:

```bash
make bench-startup-local \
  BENCH_STARTUP_SUITE=lean \
  BENCH_STARTUP_REPEAT=25 \
  BENCH_STARTUP_WORKERS=4 \
  BENCH_STARTUP_REST_MS=100
```

Compare a more parallel run:

```bash
make bench-startup-local \
  BENCH_STARTUP_SUITE=lean \
  BENCH_STARTUP_REPEAT=25 \
  BENCH_STARTUP_WORKERS=8 \
  BENCH_STARTUP_REST_MS=100
```

Compare lean-default `awsgo` against the full generated service runtime without
calling AWS APIs:

```bash
make bench-lean-vs-full-startup-local \
  BENCH_STARTUP_REPEAT=25 \
  BENCH_STARTUP_WORKERS=4 \
  BENCH_STARTUP_REST_MS=100
```

The lean-vs-full startup report prints an explicit verdict:

```text
lean-vs-full verdict: same-or-better
```

Planning estimates before running the local startup suite:

| Scope | Rest | Workers | Estimated runtime |
| --- | ---: | ---: | ---: |
| Lean pilot operation help, 25 repeats | `100ms` | 4 | under 5 minutes |
| Lean pilot operation help, 25 repeats | `100ms` | 8 | under 5 minutes |
| Zero-required read-only operation help | `100ms` | 1 | 40-100 minutes |
| Zero-required read-only operation help | `100ms` | 2 | 20-60 minutes |
| All generated operation help | `100ms` | 1 | 4-9 hours |
| All generated operation help | `100ms` | 4 | 1-3 hours |

Planned reporting fields:

| Field | Purpose |
| --- | --- |
| `service` / `operation` | Benchmark case identity |
| `cli` | `aws` or `awsgo` |
| `repeat` / `order` | Detect order bias |
| `real` / `user` / `sys` | Wall-clock and CPU cost |
| `exit` | Capture failures without hiding them |
| `stdout_bytes` / `stderr_bytes` | Sanity-check comparable output paths |

Measured lean local startup result:

| Suite | Repeats | Workers | Rest | Commands | AWS CLI avg real | awsgo avg real | Approx speedup |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Lean operation help | 25 | 4 | `100ms` | 200 | 0.400s | 0.029s | 13.93x |
| Lean operation help | 25 | 8 | `100ms` | 200 | 0.456s | 0.022s | 21.00x |

Measured lean-vs-full local startup result:

| Suite | Repeats | Workers | Rest | Commands | Full avg real | Lean avg real | Full median | Lean median | Verdict |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Lean operation help | 25 | 4 | `100ms` | 200 | 0.029s | 0.029s | 0.029s | 0.027s | same-or-better |

The local startup test does not call AWS APIs. It only invokes help/parse paths
for the lean pilot operations and measures process startup, CLI initialization,
argument parsing, help rendering, and shutdown. The result is intentionally not
a live API benchmark; it isolates the AWS CLI Python startup/runtime cost that
can be hidden by network latency in real AWS calls.

### Preliminary Local Result

On a local development machine, three read-only operations were each repeated
five times against the same AWS profile and region.

Current split architecture, using the full generated service runtime:

| Case | AWS CLI avg real | awsgo avg real | Approx speedup |
| --- | ---: | ---: | ---: |
| `sts get-caller-identity` | 0.632s | 0.360s | 1.76x |
| `ec2 describe-instances` | 0.788s | 0.406s | 1.94x |
| `s3api/s3 list-buckets` | 0.656s | 0.606s | 1.08x |

Lean runtime pilot, using direct SDK invocation:

| Case | AWS CLI avg real | awsgo avg real | Approx speedup |
| --- | ---: | ---: | ---: |
| `sts get-caller-identity` | 0.600s | 0.452s | 1.33x |
| `ec2 describe-instances` | 0.748s | 0.432s | 1.73x |
| `s3api/s3 list-buckets` | 0.700s | 0.370s | 1.89x |

These are early numbers, not a universal guarantee. The defensible claim is
that `awsgo` avoids most of the Python AWS CLI startup/runtime overhead. The
working performance thesis is that a generated lean path should beat the Python
AWS CLI for most short-lived AWS commands, especially where startup/runtime
overhead is a meaningful share of total latency. Calls dominated by AWS service
latency or network latency will show smaller wall-clock gains.

The CPU difference is more dramatic. In the same benchmark runs, AWS CLI used
roughly 0.20s-0.28s user CPU per command, while `awsgo` and the lean pilot
typically used about 0.00s-0.01s. That supports the thesis that a native Go CLI
is much lighter for short-lived AWS commands, even when total runtime is still
bounded by AWS service latency.

The lean pilot also demonstrates the binary-size opportunity:

| Service | Current split service binary | Lean service binary |
| --- | ---: | ---: |
| `sts` | 17M | 8.7M |
| `ec2` | 39M | 9.1M |
| `s3` | 17M | 9.8M |

The next step is to generate lean service binaries from SDK metadata instead of
hand-wiring selected operations.

## Known Gaps

`awsgo` is still chasing AWS CLI parity. Known areas that need continued work:

- Pagination behavior is still being aligned operation-by-operation.
- Shorthand parsing is incomplete compared with AWS CLI.
- Some generated input flag types may still need service-specific refinement.
- Split mode currently prioritizes memory/load behavior; some help/output paths
  still need to be fully aligned with the monolithic AWS-style runtime.
- Generated output is intentionally broad and noisy while coverage is being
  expanded.

## Development Notes

The generated source lives under `generated/`. Manual edits should usually go
into `servicegen.go`, `awsgo/`, or tests, then generated files should be
refreshed.

Useful targets:

```bash
make all
make check
make fmt
make goimports
make tidy
```

## License

This repository currently uses an all-rights-reserved license. Before publishing
for public use or accepting outside contributions, replace `LICENSE` with the
intended open-source license.

## AI Assistance

This project was built with assistance from OpenAI Codex using a GPT-5 Codex
model.
