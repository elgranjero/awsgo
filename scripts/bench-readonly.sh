#!/usr/bin/env bash
set -o pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
AWS_BIN="${AWS_BIN:-aws}"
AWSGO_BIN="${AWSGO_BIN:-$ROOT/bin/awsgo}"
PROFILE="${PROFILE:-${BENCH_PROFILE:-default}}"
REGION="${REGION:-${BENCH_REGION:-us-east-1}}"
REPEAT="${REPEAT:-${BENCH_REPEAT:-3}}"
OUT="${OUT:-${BENCH_OUT:-}}"
CASES="${BENCH_CASES:-}"
DRY_RUN=0
LIST_ONLY=0

ALL_CASES=(
	sts_get_caller_identity
	ecs_list_clusters
	rds_describe_db_instances
	cloudwatch_describe_alarms
	vpc_describe_vpcs
	s3_list_buckets
	ec2_describe_instances
	ec2_describe_volumes
	opensearch_list_domain_names
	elbv2_describe_load_balancers
	elb_describe_load_balancers
	guardduty_list_detectors
	elasticache_describe_cache_clusters
	datasync_list_agents
	wafv2_list_web_acls
	route53_list_hosted_zones
	configservice_describe_configuration_recorders
	cloudtrail_lookup_events
	securityhub_get_findings
	secretsmanager_list_secrets
	kms_list_keys
	ecr_describe_repositories
	efs_describe_file_systems
	inspector_list_assessment_templates
	inspector2_list_findings
	servicecatalog_search_products
	costexplorer_get_cost_and_usage
	firehose_list_delivery_streams
	sqs_list_queues
	backup_list_backup_vaults
	sns_list_topics
	glacier_list_vaults
	stepfunctions_list_state_machines
	cloudfront_list_distributions
	dynamodb_list_tables
	glue_get_databases
	acm_list_certificates
	lambda_list_functions
)

usage() {
	cat <<USAGE
Usage: scripts/bench-readonly.sh [flags]

Non-destructive AWS CLI vs awsgo timing harness for read/list/describe calls.

Flags:
  --profile name     AWS profile to use (default: $PROFILE)
  --region region    AWS region to use (default: $REGION)
  --repeat n         Runs per CLI/case (default: $REPEAT)
  --cases "a b"      Space-separated case names; default is all cases
  --aws-bin path     AWS CLI binary (default: $AWS_BIN)
  --awsgo-bin path   awsgo binary (default: $AWSGO_BIN)
  --out path         TSV output file
  --list             List benchmark case names
  --dry-run          Print commands without running them
  -h, --help         Show this help

Environment equivalents:
  BENCH_PROFILE, BENCH_REGION, BENCH_REPEAT, BENCH_CASES, BENCH_OUT
  AWS_BIN, AWSGO_BIN
USAGE
}

while [[ $# -gt 0 ]]; do
	case "$1" in
		--profile)
			PROFILE="$2"
			shift 2
			;;
		--region)
			REGION="$2"
			shift 2
			;;
		--repeat)
			REPEAT="$2"
			shift 2
			;;
		--cases)
			CASES="$2"
			shift 2
			;;
		--aws-bin)
			AWS_BIN="$2"
			shift 2
			;;
		--awsgo-bin)
			AWSGO_BIN="$2"
			shift 2
			;;
		--out)
			OUT="$2"
			shift 2
			;;
		--list)
			LIST_ONLY=1
			shift
			;;
		--dry-run)
			DRY_RUN=1
			shift
			;;
		-h|--help)
			usage
			exit 0
			;;
		*)
			echo "unknown argument: $1" >&2
			usage >&2
			exit 2
			;;
	esac
done

if [[ "$LIST_ONLY" -eq 1 ]]; then
	printf '%s\n' "${ALL_CASES[@]}"
	exit 0
fi

if ! [[ "$REPEAT" =~ ^[0-9]+$ ]] || [[ "$REPEAT" -lt 1 ]]; then
	echo "--repeat must be a positive integer" >&2
	exit 2
fi

if [[ -z "$CASES" ]]; then
	CASES="${ALL_CASES[*]}"
fi

if [[ -z "$OUT" ]]; then
	mkdir -p "$ROOT/benchmarks"
	OUT="$ROOT/benchmarks/readonly-$(date -u +%Y%m%d-%H%M%S).tsv"
fi

if [[ "$DRY_RUN" -eq 0 ]]; then
	if ! command -v "$AWS_BIN" >/dev/null 2>&1; then
		echo "AWS CLI not found: $AWS_BIN" >&2
		exit 127
	fi
	if [[ ! -x "$AWSGO_BIN" ]]; then
		echo "awsgo binary not found or not executable: $AWSGO_BIN" >&2
		echo 'Run: make build-split SPLIT_SERVICES="$(make bench-services)"' >&2
		exit 127
	fi
fi

BENCH_START_DATE="${BENCH_START_DATE:-}"
BENCH_END_DATE="${BENCH_END_DATE:-}"
if [[ -z "$BENCH_START_DATE" || -z "$BENCH_END_DATE" ]]; then
	BENCH_END_DATE="$(date -u +%Y-%m-%d)"
	BENCH_START_DATE="$(date -u -v-1d +%Y-%m-%d 2>/dev/null || date -u -d yesterday +%Y-%m-%d 2>/dev/null || true)"
fi

shell_join() {
	local out="" arg
	for arg in "$@"; do
		printf -v out '%s%q ' "$out" "$arg"
	done
	printf '%s' "${out% }"
}

redact_local_paths() {
	local value="$1"
	value="${value//$ROOT/\$REPO}"
	if [[ -n "${HOME:-}" ]]; then
		value="${value//$HOME/\$HOME}"
	fi
	printf '%s' "$value"
}

case_definition() {
	local name="$1"
	label="$name"
	aws_service=""
	awsgo_service=""
	operation=""
	aws_args=()
	awsgo_args=()

	case "$name" in
		sts_get_caller_identity)
			aws_service="sts"; awsgo_service="sts"; operation="get-caller-identity"
			;;
		ecs_list_clusters)
			aws_service="ecs"; awsgo_service="ecs"; operation="list-clusters"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		rds_describe_db_instances)
			aws_service="rds"; awsgo_service="rds"; operation="describe-db-instances"
			aws_args=(--max-records 20); awsgo_args=(--max-records 20)
			;;
		cloudwatch_describe_alarms)
			aws_service="cloudwatch"; awsgo_service="cloudwatch"; operation="describe-alarms"
			aws_args=(--max-records 20); awsgo_args=(--max-records 20)
			;;
		vpc_describe_vpcs)
			aws_service="ec2"; awsgo_service="ec2"; operation="describe-vpcs"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		s3_list_buckets)
			aws_service="s3api"; awsgo_service="s3"; operation="list-buckets"
			aws_args=(--max-buckets 20); awsgo_args=(--max-buckets 20)
			;;
		ec2_describe_instances)
			aws_service="ec2"; awsgo_service="ec2"; operation="describe-instances"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		ec2_describe_volumes)
			aws_service="ec2"; awsgo_service="ec2"; operation="describe-volumes"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		opensearch_list_domain_names)
			aws_service="opensearch"; awsgo_service="opensearch"; operation="list-domain-names"
			;;
		elbv2_describe_load_balancers)
			aws_service="elbv2"; awsgo_service="elasticloadbalancingv2"; operation="describe-load-balancers"
			aws_args=(--page-size 20); awsgo_args=(--page-size 20)
			;;
		elb_describe_load_balancers)
			aws_service="elb"; awsgo_service="elasticloadbalancing"; operation="describe-load-balancers"
			aws_args=(--page-size 20); awsgo_args=(--page-size 20)
			;;
		guardduty_list_detectors)
			aws_service="guardduty"; awsgo_service="guardduty"; operation="list-detectors"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		elasticache_describe_cache_clusters)
			aws_service="elasticache"; awsgo_service="elasticache"; operation="describe-cache-clusters"
			aws_args=(--max-records 20); awsgo_args=(--max-records 20)
			;;
		datasync_list_agents)
			aws_service="datasync"; awsgo_service="datasync"; operation="list-agents"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		wafv2_list_web_acls)
			aws_service="wafv2"; awsgo_service="wafv2"; operation="list-web-acls"
			aws_args=(--scope REGIONAL --limit 20); awsgo_args=(--scope REGIONAL --limit 20)
			;;
		route53_list_hosted_zones)
			aws_service="route53"; awsgo_service="route53"; operation="list-hosted-zones"
			aws_args=(--max-items 20); awsgo_args=(--max-items 20)
			;;
		configservice_describe_configuration_recorders)
			aws_service="configservice"; awsgo_service="configservice"; operation="describe-configuration-recorders"
			;;
		cloudtrail_lookup_events)
			aws_service="cloudtrail"; awsgo_service="cloudtrail"; operation="lookup-events"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		securityhub_get_findings)
			aws_service="securityhub"; awsgo_service="securityhub"; operation="get-findings"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		secretsmanager_list_secrets)
			aws_service="secretsmanager"; awsgo_service="secretsmanager"; operation="list-secrets"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		kms_list_keys)
			aws_service="kms"; awsgo_service="kms"; operation="list-keys"
			aws_args=(--limit 20); awsgo_args=(--limit 20)
			;;
		ecr_describe_repositories)
			aws_service="ecr"; awsgo_service="ecr"; operation="describe-repositories"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		efs_describe_file_systems)
			aws_service="efs"; awsgo_service="efs"; operation="describe-file-systems"
			aws_args=(--max-items 20); awsgo_args=(--max-items 20)
			;;
		inspector_list_assessment_templates)
			aws_service="inspector"; awsgo_service="inspector"; operation="list-assessment-templates"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		inspector2_list_findings)
			aws_service="inspector2"; awsgo_service="inspector2"; operation="list-findings"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		servicecatalog_search_products)
			aws_service="servicecatalog"; awsgo_service="servicecatalog"; operation="search-products"
			aws_args=(--page-size 20); awsgo_args=(--page-size 20)
			;;
		costexplorer_get_cost_and_usage)
			aws_service="ce"; awsgo_service="costexplorer"; operation="get-cost-and-usage"
			aws_args=(--time-period "Start=$BENCH_START_DATE,End=$BENCH_END_DATE" --granularity DAILY --metrics UnblendedCost)
			awsgo_args=(--time-period "{\"Start\":\"$BENCH_START_DATE\",\"End\":\"$BENCH_END_DATE\"}" --granularity DAILY --metrics '["UnblendedCost"]')
			;;
		firehose_list_delivery_streamS|firehose_list_delivery_streams)
			aws_service="firehose"; awsgo_service="firehose"; operation="list-delivery-streams"
			label="firehose_list_delivery_streams"
			aws_args=(--limit 20); awsgo_args=(--limit 20)
			;;
		sqs_list_queues)
			aws_service="sqs"; awsgo_service="sqs"; operation="list-queues"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		backup_list_backup_vaults)
			aws_service="backup"; awsgo_service="backup"; operation="list-backup-vaults"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		sns_list_topics)
			aws_service="sns"; awsgo_service="sns"; operation="list-topics"
			;;
		glacier_list_vaults)
			aws_service="glacier"; awsgo_service="glacier"; operation="list-vaults"
			aws_args=(--account-id - --limit 20); awsgo_args=(--account-id - --limit 20)
			;;
		stepfunctions_list_state_machines)
			aws_service="stepfunctions"; awsgo_service="sfn"; operation="list-state-machines"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		cloudfront_list_distributions)
			aws_service="cloudfront"; awsgo_service="cloudfront"; operation="list-distributions"
			aws_args=(--max-items 20); awsgo_args=(--max-items 20)
			;;
		dynamodb_list_tables)
			aws_service="dynamodb"; awsgo_service="dynamodb"; operation="list-tables"
			aws_args=(--limit 20); awsgo_args=(--limit 20)
			;;
		glue_get_databases)
			aws_service="glue"; awsgo_service="glue"; operation="get-databases"
			aws_args=(--max-results 20); awsgo_args=(--max-results 20)
			;;
		acm_list_certificates)
			aws_service="acm"; awsgo_service="acm"; operation="list-certificates"
			aws_args=(--max-items 20); awsgo_args=(--max-items 20)
			;;
		lambda_list_functions)
			aws_service="lambda"; awsgo_service="lambda"; operation="list-functions"
			aws_args=(--max-items 20); awsgo_args=(--max-items 20)
			;;
		*)
			echo "unknown benchmark case: $name" >&2
			return 1
			;;
	esac
}

build_aws_cmd() {
	aws_cmd=("$AWS_BIN" "$aws_service" "$operation" --output json --no-cli-pager --no-paginate)
	if [[ -n "$PROFILE" ]]; then
		aws_cmd+=(--profile "$PROFILE")
	fi
	if [[ -n "$REGION" ]]; then
		aws_cmd+=(--region "$REGION")
	fi
	aws_cmd+=("${aws_args[@]}")
}

build_awsgo_cmd() {
	awsgo_cmd=("$AWSGO_BIN" "$awsgo_service" "$operation" --output json)
	if [[ -n "$PROFILE" ]]; then
		awsgo_cmd+=(--profile "$PROFILE")
	fi
	if [[ -n "$REGION" ]]; then
		awsgo_cmd+=(--region "$REGION")
	fi
	awsgo_cmd+=("${awsgo_args[@]}")
}

run_one() {
	local case_name="$1" cli="$2" iter="$3"
	shift 3
	local stdout_tmp stderr_tmp status real user sys out_bytes err_bytes command
	stdout_tmp="$(mktemp "${TMPDIR:-/tmp}/awsgo-bench-out.XXXXXX")"
	stderr_tmp="$(mktemp "${TMPDIR:-/tmp}/awsgo-bench-err.XXXXXX")"
	command="$(redact_local_paths "$(shell_join "$@")")"

	/usr/bin/time -p "$@" >"$stdout_tmp" 2>"$stderr_tmp"
	status=$?
	real="$(awk '$1=="real"{v=$2} END{print v+0}' "$stderr_tmp")"
	user="$(awk '$1=="user"{v=$2} END{print v+0}' "$stderr_tmp")"
	sys="$(awk '$1=="sys"{v=$2} END{print v+0}' "$stderr_tmp")"
	out_bytes="$(wc -c <"$stdout_tmp" | tr -d ' ')"
	err_bytes="$(awk '$1!="real" && $1!="user" && $1!="sys"{n+=length($0)+1} END{print n+0}' "$stderr_tmp")"

	printf '%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n' \
		"$case_name" "$cli" "$iter" "$status" "$real" "$user" "$sys" "$out_bytes" "$err_bytes" "$command" >>"$OUT"

	rm -f "$stdout_tmp" "$stderr_tmp"
}

if [[ "$DRY_RUN" -eq 0 ]]; then
	mkdir -p "$(dirname "$OUT")"
	printf 'case\tcli\titeration\tstatus\treal\tuser\tsys\tstdout_bytes\tstderr_bytes\tcommand\n' >"$OUT"
	echo "writing benchmark results to $OUT"
fi

for case_name in $CASES; do
	case_definition "$case_name" || exit 2
	build_aws_cmd
	build_awsgo_cmd

	if [[ "$DRY_RUN" -eq 1 ]]; then
		printf '%s\taws\t%s\n' "$label" "$(redact_local_paths "$(shell_join "${aws_cmd[@]}")")"
		printf '%s\tawsgo\tAWSGO_DISABLE_PAGINATOR=1 %s\n' "$label" "$(redact_local_paths "$(shell_join "${awsgo_cmd[@]}")")"
		continue
	fi

	echo "benchmarking $label"
	for ((i = 1; i <= REPEAT; i++)); do
		run_one "$label" aws "$i" "${aws_cmd[@]}"
		run_one "$label" awsgo "$i" env AWSGO_DISABLE_PAGINATOR=1 "${awsgo_cmd[@]}"
	done
done

if [[ "$DRY_RUN" -eq 1 ]]; then
	exit 0
fi

echo
echo "summary (successful runs only):"
awk -F '\t' '
NR == 1 { next }
{
	key = $1 "\t" $2
	total[key]++
	if ($4 == 0) {
		ok[key]++
		real[key] += $5
		user[key] += $6
		sys[key] += $7
	} else {
		fail[key]++
	}
}
END {
	printf "%-45s %-8s %8s %8s %8s %8s\n", "case", "cli", "ok/total", "avg_real", "avg_user", "avg_sys"
	for (key in total) {
		split(key, parts, "\t")
		if (ok[key] > 0) {
			printf "%-45s %-8s %3d/%-4d %8.3f %8.3f %8.3f\n", parts[1], parts[2], ok[key], total[key], real[key]/ok[key], user[key]/ok[key], sys[key]/ok[key]
		} else {
			printf "%-45s %-8s %3d/%-4d %8s %8s %8s\n", parts[1], parts[2], 0, total[key], "failed", "failed", "failed"
		}
	}
}
' "$OUT" | sort
