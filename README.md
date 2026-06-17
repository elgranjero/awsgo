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

Build the dispatcher and all generated service binaries:

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

The harness:

- Uses read-only AWS APIs.
- Disables pagination for both CLIs by default and caps result sizes where the
  underlying API exposes a limit field.
- Redirects command output and records only timing, status, and byte counts.
- Writes local TSV files under `benchmarks/`, which are ignored by git.
- Still uses real AWS API calls, so normal AWS API rate limits and any
  service-specific API pricing still apply.

Use this for a one-time data point before writing performance claims. The useful
question is not “is Go always faster?” but “for common AWS read-only commands,
how much startup/runtime overhead do we avoid compared with the Python AWS CLI?”

### Preliminary Local Result

On a local development machine, three read-only operations were each repeated
five times against the same AWS profile and region:

| Case | AWS CLI avg real | awsgo avg real | Approx speedup |
| --- | ---: | ---: | ---: |
| `sts get-caller-identity` | 0.682s | 0.442s | 1.54x |
| `ec2 describe-instances` | 0.778s | 0.524s | 1.48x |
| `s3api/s3 list-buckets` | 0.666s | 0.430s | 1.55x |

Across that small sample, AWS CLI averaged 0.709s real time while `awsgo`
averaged 0.465s, roughly a 1.52x wall-clock improvement. The bigger local
difference was CPU time: AWS CLI averaged 0.234s user CPU per command, while
`awsgo` averaged 0.010s. That supports the expected thesis: for short-lived AWS
read calls, a native Go CLI avoids much of the Python process/runtime overhead,
even when the network call itself is still the dominant cost.

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
