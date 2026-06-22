package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codebuild"
)

var fields_batch_delete_builds = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_batch_get_build_batches = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_batch_get_builds = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_batch_get_command_executions = []leanruntime.Field{
	{Name: "CommandExecutionIds", Flag: "command-execution-ids", Type: "[]string", Required: true},
	{Name: "SandboxId", Flag: "sandbox-id", Type: "*string", Required: true},
}

var fields_batch_get_fleets = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_batch_get_projects = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_batch_get_report_groups = []leanruntime.Field{
	{Name: "ReportGroupArns", Flag: "report-group-arns", Type: "[]string", Required: true},
}

var fields_batch_get_reports = []leanruntime.Field{
	{Name: "ReportArns", Flag: "report-arns", Type: "[]string", Required: true},
}

var fields_batch_get_sandboxes = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_create_fleet = []leanruntime.Field{
	{Name: "BaseCapacity", Flag: "base-capacity", Type: "*int32", Required: true},
	{Name: "ComputeConfiguration", Flag: "compute-configuration", Type: "*types.ComputeConfiguration", Required: false},
	{Name: "ComputeType", Flag: "compute-type", Type: "types.ComputeType", Required: true},
	{Name: "EnvironmentType", Flag: "environment-type", Type: "types.EnvironmentType", Required: true},
	{Name: "FleetServiceRole", Flag: "fleet-service-role", Type: "*string", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OverflowBehavior", Flag: "overflow-behavior", Type: "types.FleetOverflowBehavior", Required: false},
	{Name: "ProxyConfiguration", Flag: "proxy-configuration", Type: "*types.ProxyConfiguration", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfigurationInput", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "Artifacts", Flag: "artifacts", Type: "*types.ProjectArtifacts", Required: true},
	{Name: "AutoRetryLimit", Flag: "auto-retry-limit", Type: "*int32", Required: false},
	{Name: "BadgeEnabled", Flag: "badge-enabled", Type: "*bool", Required: false},
	{Name: "BuildBatchConfig", Flag: "build-batch-config", Type: "*types.ProjectBuildBatchConfig", Required: false},
	{Name: "Cache", Flag: "cache", Type: "*types.ProjectCache", Required: false},
	{Name: "ConcurrentBuildLimit", Flag: "concurrent-build-limit", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Environment", Flag: "environment", Type: "*types.ProjectEnvironment", Required: true},
	{Name: "FileSystemLocations", Flag: "file-system-locations", Type: "[]types.ProjectFileSystemLocation", Required: false},
	{Name: "LogsConfig", Flag: "logs-config", Type: "*types.LogsConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueuedTimeoutInMinutes", Flag: "queued-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "SecondaryArtifacts", Flag: "secondary-artifacts", Type: "[]types.ProjectArtifacts", Required: false},
	{Name: "SecondarySourceVersions", Flag: "secondary-source-versions", Type: "[]types.ProjectSourceVersion", Required: false},
	{Name: "SecondarySources", Flag: "secondary-sources", Type: "[]types.ProjectSource", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*types.ProjectSource", Required: true},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutInMinutes", Flag: "timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_report_group = []leanruntime.Field{
	{Name: "ExportConfig", Flag: "export-config", Type: "*types.ReportExportConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ReportType", Required: true},
}

var fields_create_webhook = []leanruntime.Field{
	{Name: "BranchFilter", Flag: "branch-filter", Type: "*string", Required: false},
	{Name: "BuildType", Flag: "build-type", Type: "types.WebhookBuildType", Required: false},
	{Name: "FilterGroups", Flag: "filter-groups", Type: "[][]types.WebhookFilter", Required: false},
	{Name: "ManualCreation", Flag: "manual-creation", Type: "*bool", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "PullRequestBuildPolicy", Flag: "pull-request-build-policy", Type: "*types.PullRequestBuildPolicy", Required: false},
	{Name: "ScopeConfiguration", Flag: "scope-configuration", Type: "*types.ScopeConfiguration", Required: false},
}

var fields_delete_build_batch = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_fleet = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_report = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_report_group = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "DeleteReports", Flag: "delete-reports", Type: "bool", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_source_credentials = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_webhook = []leanruntime.Field{
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
}

var fields_describe_code_coverages = []leanruntime.Field{
	{Name: "MaxLineCoveragePercentage", Flag: "max-line-coverage-percentage", Type: "*float64", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MinLineCoveragePercentage", Flag: "min-line-coverage-percentage", Type: "*float64", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportArn", Flag: "report-arn", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ReportCodeCoverageSortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_describe_test_cases = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TestCaseFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportArn", Flag: "report-arn", Type: "*string", Required: true},
}

var fields_get_report_group_trend = []leanruntime.Field{
	{Name: "NumOfReports", Flag: "num-of-reports", Type: "*int32", Required: false},
	{Name: "ReportGroupArn", Flag: "report-group-arn", Type: "*string", Required: true},
	{Name: "TrendField", Flag: "trend-field", Type: "types.ReportGroupTrendFieldType", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_import_source_credentials = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: true},
	{Name: "ServerType", Flag: "server-type", Type: "types.ServerType", Required: true},
	{Name: "ShouldOverwrite", Flag: "should-overwrite", Type: "*bool", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_invalidate_project_cache = []leanruntime.Field{
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
}

var fields_list_build_batches = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.BuildBatchFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_build_batches_for_project = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.BuildBatchFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_builds = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_builds_for_project = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_command_executions_for_sandbox = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SandboxId", Flag: "sandbox-id", Type: "*string", Required: true},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_curated_environment_images = []leanruntime.Field{}

var fields_list_fleets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.FleetSortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ProjectSortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_report_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ReportGroupSortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_reports = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReportFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_reports_for_report_group = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReportFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportGroupArn", Flag: "report-group-arn", Type: "*string", Required: true},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_sandboxes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_sandboxes_for_project = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_shared_projects = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SharedResourceSortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_shared_report_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SharedResourceSortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_list_source_credentials = []leanruntime.Field{}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_retry_build = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
}

var fields_retry_build_batch = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "RetryType", Flag: "retry-type", Type: "types.RetryBuildBatchType", Required: false},
}

var fields_start_build = []leanruntime.Field{
	{Name: "ArtifactsOverride", Flag: "artifacts-override", Type: "*types.ProjectArtifacts", Required: false},
	{Name: "AutoRetryLimitOverride", Flag: "auto-retry-limit-override", Type: "*int32", Required: false},
	{Name: "BuildStatusConfigOverride", Flag: "build-status-config-override", Type: "*types.BuildStatusConfig", Required: false},
	{Name: "BuildspecOverride", Flag: "buildspec-override", Type: "*string", Required: false},
	{Name: "CacheOverride", Flag: "cache-override", Type: "*types.ProjectCache", Required: false},
	{Name: "CertificateOverride", Flag: "certificate-override", Type: "*string", Required: false},
	{Name: "ComputeTypeOverride", Flag: "compute-type-override", Type: "types.ComputeType", Required: false},
	{Name: "DebugSessionEnabled", Flag: "debug-session-enabled", Type: "*bool", Required: false},
	{Name: "EncryptionKeyOverride", Flag: "encryption-key-override", Type: "*string", Required: false},
	{Name: "EnvironmentTypeOverride", Flag: "environment-type-override", Type: "types.EnvironmentType", Required: false},
	{Name: "EnvironmentVariablesOverride", Flag: "environment-variables-override", Type: "[]types.EnvironmentVariable", Required: false},
	{Name: "FleetOverride", Flag: "fleet-override", Type: "*types.ProjectFleet", Required: false},
	{Name: "GitCloneDepthOverride", Flag: "git-clone-depth-override", Type: "*int32", Required: false},
	{Name: "GitSubmodulesConfigOverride", Flag: "git-submodules-config-override", Type: "*types.GitSubmodulesConfig", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ImageOverride", Flag: "image-override", Type: "*string", Required: false},
	{Name: "ImagePullCredentialsTypeOverride", Flag: "image-pull-credentials-type-override", Type: "types.ImagePullCredentialsType", Required: false},
	{Name: "InsecureSslOverride", Flag: "insecure-ssl-override", Type: "*bool", Required: false},
	{Name: "LogsConfigOverride", Flag: "logs-config-override", Type: "*types.LogsConfig", Required: false},
	{Name: "PrivilegedModeOverride", Flag: "privileged-mode-override", Type: "*bool", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "QueuedTimeoutInMinutesOverride", Flag: "queued-timeout-in-minutes-override", Type: "*int32", Required: false},
	{Name: "RegistryCredentialOverride", Flag: "registry-credential-override", Type: "*types.RegistryCredential", Required: false},
	{Name: "ReportBuildStatusOverride", Flag: "report-build-status-override", Type: "*bool", Required: false},
	{Name: "SecondaryArtifactsOverride", Flag: "secondary-artifacts-override", Type: "[]types.ProjectArtifacts", Required: false},
	{Name: "SecondarySourcesOverride", Flag: "secondary-sources-override", Type: "[]types.ProjectSource", Required: false},
	{Name: "SecondarySourcesVersionOverride", Flag: "secondary-sources-version-override", Type: "[]types.ProjectSourceVersion", Required: false},
	{Name: "ServiceRoleOverride", Flag: "service-role-override", Type: "*string", Required: false},
	{Name: "SourceAuthOverride", Flag: "source-auth-override", Type: "*types.SourceAuth", Required: false},
	{Name: "SourceLocationOverride", Flag: "source-location-override", Type: "*string", Required: false},
	{Name: "SourceTypeOverride", Flag: "source-type-override", Type: "types.SourceType", Required: false},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
	{Name: "TimeoutInMinutesOverride", Flag: "timeout-in-minutes-override", Type: "*int32", Required: false},
}

var fields_start_build_batch = []leanruntime.Field{
	{Name: "ArtifactsOverride", Flag: "artifacts-override", Type: "*types.ProjectArtifacts", Required: false},
	{Name: "BuildBatchConfigOverride", Flag: "build-batch-config-override", Type: "*types.ProjectBuildBatchConfig", Required: false},
	{Name: "BuildTimeoutInMinutesOverride", Flag: "build-timeout-in-minutes-override", Type: "*int32", Required: false},
	{Name: "BuildspecOverride", Flag: "buildspec-override", Type: "*string", Required: false},
	{Name: "CacheOverride", Flag: "cache-override", Type: "*types.ProjectCache", Required: false},
	{Name: "CertificateOverride", Flag: "certificate-override", Type: "*string", Required: false},
	{Name: "ComputeTypeOverride", Flag: "compute-type-override", Type: "types.ComputeType", Required: false},
	{Name: "DebugSessionEnabled", Flag: "debug-session-enabled", Type: "*bool", Required: false},
	{Name: "EncryptionKeyOverride", Flag: "encryption-key-override", Type: "*string", Required: false},
	{Name: "EnvironmentTypeOverride", Flag: "environment-type-override", Type: "types.EnvironmentType", Required: false},
	{Name: "EnvironmentVariablesOverride", Flag: "environment-variables-override", Type: "[]types.EnvironmentVariable", Required: false},
	{Name: "GitCloneDepthOverride", Flag: "git-clone-depth-override", Type: "*int32", Required: false},
	{Name: "GitSubmodulesConfigOverride", Flag: "git-submodules-config-override", Type: "*types.GitSubmodulesConfig", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ImageOverride", Flag: "image-override", Type: "*string", Required: false},
	{Name: "ImagePullCredentialsTypeOverride", Flag: "image-pull-credentials-type-override", Type: "types.ImagePullCredentialsType", Required: false},
	{Name: "InsecureSslOverride", Flag: "insecure-ssl-override", Type: "*bool", Required: false},
	{Name: "LogsConfigOverride", Flag: "logs-config-override", Type: "*types.LogsConfig", Required: false},
	{Name: "PrivilegedModeOverride", Flag: "privileged-mode-override", Type: "*bool", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "QueuedTimeoutInMinutesOverride", Flag: "queued-timeout-in-minutes-override", Type: "*int32", Required: false},
	{Name: "RegistryCredentialOverride", Flag: "registry-credential-override", Type: "*types.RegistryCredential", Required: false},
	{Name: "ReportBuildBatchStatusOverride", Flag: "report-build-batch-status-override", Type: "*bool", Required: false},
	{Name: "SecondaryArtifactsOverride", Flag: "secondary-artifacts-override", Type: "[]types.ProjectArtifacts", Required: false},
	{Name: "SecondarySourcesOverride", Flag: "secondary-sources-override", Type: "[]types.ProjectSource", Required: false},
	{Name: "SecondarySourcesVersionOverride", Flag: "secondary-sources-version-override", Type: "[]types.ProjectSourceVersion", Required: false},
	{Name: "ServiceRoleOverride", Flag: "service-role-override", Type: "*string", Required: false},
	{Name: "SourceAuthOverride", Flag: "source-auth-override", Type: "*types.SourceAuth", Required: false},
	{Name: "SourceLocationOverride", Flag: "source-location-override", Type: "*string", Required: false},
	{Name: "SourceTypeOverride", Flag: "source-type-override", Type: "types.SourceType", Required: false},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
}

var fields_start_command_execution = []leanruntime.Field{
	{Name: "Command", Flag: "command", Type: "*string", Required: true},
	{Name: "SandboxId", Flag: "sandbox-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.CommandType", Required: false},
}

var fields_start_sandbox = []leanruntime.Field{
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: false},
}

var fields_start_sandbox_connection = []leanruntime.Field{
	{Name: "SandboxId", Flag: "sandbox-id", Type: "*string", Required: true},
}

var fields_stop_build = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_stop_build_batch = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_stop_sandbox = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_fleet = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "BaseCapacity", Flag: "base-capacity", Type: "*int32", Required: false},
	{Name: "ComputeConfiguration", Flag: "compute-configuration", Type: "*types.ComputeConfiguration", Required: false},
	{Name: "ComputeType", Flag: "compute-type", Type: "types.ComputeType", Required: false},
	{Name: "EnvironmentType", Flag: "environment-type", Type: "types.EnvironmentType", Required: false},
	{Name: "FleetServiceRole", Flag: "fleet-service-role", Type: "*string", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "OverflowBehavior", Flag: "overflow-behavior", Type: "types.FleetOverflowBehavior", Required: false},
	{Name: "ProxyConfiguration", Flag: "proxy-configuration", Type: "*types.ProxyConfiguration", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfigurationInput", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_update_project = []leanruntime.Field{
	{Name: "Artifacts", Flag: "artifacts", Type: "*types.ProjectArtifacts", Required: false},
	{Name: "AutoRetryLimit", Flag: "auto-retry-limit", Type: "*int32", Required: false},
	{Name: "BadgeEnabled", Flag: "badge-enabled", Type: "*bool", Required: false},
	{Name: "BuildBatchConfig", Flag: "build-batch-config", Type: "*types.ProjectBuildBatchConfig", Required: false},
	{Name: "Cache", Flag: "cache", Type: "*types.ProjectCache", Required: false},
	{Name: "ConcurrentBuildLimit", Flag: "concurrent-build-limit", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Environment", Flag: "environment", Type: "*types.ProjectEnvironment", Required: false},
	{Name: "FileSystemLocations", Flag: "file-system-locations", Type: "[]types.ProjectFileSystemLocation", Required: false},
	{Name: "LogsConfig", Flag: "logs-config", Type: "*types.LogsConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueuedTimeoutInMinutes", Flag: "queued-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "SecondaryArtifacts", Flag: "secondary-artifacts", Type: "[]types.ProjectArtifacts", Required: false},
	{Name: "SecondarySourceVersions", Flag: "secondary-source-versions", Type: "[]types.ProjectSourceVersion", Required: false},
	{Name: "SecondarySources", Flag: "secondary-sources", Type: "[]types.ProjectSource", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.ProjectSource", Required: false},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutInMinutes", Flag: "timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_update_project_visibility = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "ProjectVisibility", Flag: "project-visibility", Type: "types.ProjectVisibilityType", Required: true},
	{Name: "ResourceAccessRole", Flag: "resource-access-role", Type: "*string", Required: false},
}

var fields_update_report_group = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ExportConfig", Flag: "export-config", Type: "*types.ReportExportConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_update_webhook = []leanruntime.Field{
	{Name: "BranchFilter", Flag: "branch-filter", Type: "*string", Required: false},
	{Name: "BuildType", Flag: "build-type", Type: "types.WebhookBuildType", Required: false},
	{Name: "FilterGroups", Flag: "filter-groups", Type: "[][]types.WebhookFilter", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "PullRequestBuildPolicy", Flag: "pull-request-build-policy", Type: "*types.PullRequestBuildPolicy", Required: false},
	{Name: "RotateSecret", Flag: "rotate-secret", Type: "bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-delete-builds": {
			Name:   "batch-delete-builds",
			Fields: fields_batch_delete_builds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteBuildsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_builds, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteBuilds(ctx, input)
			},
		},
		"batch-get-build-batches": {
			Name:   "batch-get-build-batches",
			Fields: fields_batch_get_build_batches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetBuildBatchesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_build_batches, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetBuildBatches(ctx, input)
			},
		},
		"batch-get-builds": {
			Name:   "batch-get-builds",
			Fields: fields_batch_get_builds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetBuildsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_builds, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetBuilds(ctx, input)
			},
		},
		"batch-get-command-executions": {
			Name:   "batch-get-command-executions",
			Fields: fields_batch_get_command_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCommandExecutionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_command_executions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCommandExecutions(ctx, input)
			},
		},
		"batch-get-fleets": {
			Name:   "batch-get-fleets",
			Fields: fields_batch_get_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFleetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_fleets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetFleets(ctx, input)
			},
		},
		"batch-get-projects": {
			Name:   "batch-get-projects",
			Fields: fields_batch_get_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetProjectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_projects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetProjects(ctx, input)
			},
		},
		"batch-get-report-groups": {
			Name:   "batch-get-report-groups",
			Fields: fields_batch_get_report_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetReportGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_report_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetReportGroups(ctx, input)
			},
		},
		"batch-get-reports": {
			Name:   "batch-get-reports",
			Fields: fields_batch_get_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetReportsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_reports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetReports(ctx, input)
			},
		},
		"batch-get-sandboxes": {
			Name:   "batch-get-sandboxes",
			Fields: fields_batch_get_sandboxes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetSandboxesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_sandboxes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetSandboxes(ctx, input)
			},
		},
		"create-fleet": {
			Name:   "create-fleet",
			Fields: fields_create_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleet(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"create-report-group": {
			Name:   "create-report-group",
			Fields: fields_create_report_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReportGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_report_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReportGroup(ctx, input)
			},
		},
		"create-webhook": {
			Name:   "create-webhook",
			Fields: fields_create_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebhook(ctx, input)
			},
		},
		"delete-build-batch": {
			Name:   "delete-build-batch",
			Fields: fields_delete_build_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBuildBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_build_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBuildBatch(ctx, input)
			},
		},
		"delete-fleet": {
			Name:   "delete-fleet",
			Fields: fields_delete_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleet(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-report": {
			Name:   "delete-report",
			Fields: fields_delete_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReport(ctx, input)
			},
		},
		"delete-report-group": {
			Name:   "delete-report-group",
			Fields: fields_delete_report_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReportGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_report_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReportGroup(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-source-credentials": {
			Name:   "delete-source-credentials",
			Fields: fields_delete_source_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSourceCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_source_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSourceCredentials(ctx, input)
			},
		},
		"delete-webhook": {
			Name:   "delete-webhook",
			Fields: fields_delete_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebhook(ctx, input)
			},
		},
		"describe-code-coverages": {
			Name:   "describe-code-coverages",
			Fields: fields_describe_code_coverages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCodeCoveragesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_code_coverages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCodeCoverages(ctx, input)
				}
				var results []*svc.DescribeCodeCoveragesOutput
				p := svc.NewDescribeCodeCoveragesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-test-cases": {
			Name:   "describe-test-cases",
			Fields: fields_describe_test_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTestCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_test_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTestCases(ctx, input)
				}
				var results []*svc.DescribeTestCasesOutput
				p := svc.NewDescribeTestCasesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-report-group-trend": {
			Name:   "get-report-group-trend",
			Fields: fields_get_report_group_trend,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReportGroupTrendInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_report_group_trend, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReportGroupTrend(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"import-source-credentials": {
			Name:   "import-source-credentials",
			Fields: fields_import_source_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportSourceCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_source_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportSourceCredentials(ctx, input)
			},
		},
		"invalidate-project-cache": {
			Name:   "invalidate-project-cache",
			Fields: fields_invalidate_project_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvalidateProjectCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invalidate_project_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvalidateProjectCache(ctx, input)
			},
		},
		"list-build-batches": {
			Name:   "list-build-batches",
			Fields: fields_list_build_batches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuildBatchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_build_batches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuildBatches(ctx, input)
				}
				var results []*svc.ListBuildBatchesOutput
				p := svc.NewListBuildBatchesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-build-batches-for-project": {
			Name:   "list-build-batches-for-project",
			Fields: fields_list_build_batches_for_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuildBatchesForProjectInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_build_batches_for_project, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuildBatchesForProject(ctx, input)
				}
				var results []*svc.ListBuildBatchesForProjectOutput
				p := svc.NewListBuildBatchesForProjectPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-builds": {
			Name:   "list-builds",
			Fields: fields_list_builds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuildsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_builds, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuilds(ctx, input)
				}
				var results []*svc.ListBuildsOutput
				p := svc.NewListBuildsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-builds-for-project": {
			Name:   "list-builds-for-project",
			Fields: fields_list_builds_for_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuildsForProjectInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_builds_for_project, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuildsForProject(ctx, input)
				}
				var results []*svc.ListBuildsForProjectOutput
				p := svc.NewListBuildsForProjectPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-command-executions-for-sandbox": {
			Name:   "list-command-executions-for-sandbox",
			Fields: fields_list_command_executions_for_sandbox,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommandExecutionsForSandboxInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_command_executions_for_sandbox, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommandExecutionsForSandbox(ctx, input)
				}
				var results []*svc.ListCommandExecutionsForSandboxOutput
				p := svc.NewListCommandExecutionsForSandboxPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-curated-environment-images": {
			Name:   "list-curated-environment-images",
			Fields: fields_list_curated_environment_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCuratedEnvironmentImagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_curated_environment_images, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCuratedEnvironmentImages(ctx, input)
			},
		},
		"list-fleets": {
			Name:   "list-fleets",
			Fields: fields_list_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleets(ctx, input)
				}
				var results []*svc.ListFleetsOutput
				p := svc.NewListFleetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-projects": {
			Name:   "list-projects",
			Fields: fields_list_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjects(ctx, input)
				}
				var results []*svc.ListProjectsOutput
				p := svc.NewListProjectsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-report-groups": {
			Name:   "list-report-groups",
			Fields: fields_list_report_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_report_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReportGroups(ctx, input)
				}
				var results []*svc.ListReportGroupsOutput
				p := svc.NewListReportGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-reports": {
			Name:   "list-reports",
			Fields: fields_list_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReports(ctx, input)
				}
				var results []*svc.ListReportsOutput
				p := svc.NewListReportsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-reports-for-report-group": {
			Name:   "list-reports-for-report-group",
			Fields: fields_list_reports_for_report_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportsForReportGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reports_for_report_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReportsForReportGroup(ctx, input)
				}
				var results []*svc.ListReportsForReportGroupOutput
				p := svc.NewListReportsForReportGroupPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-sandboxes": {
			Name:   "list-sandboxes",
			Fields: fields_list_sandboxes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSandboxesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sandboxes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSandboxes(ctx, input)
				}
				var results []*svc.ListSandboxesOutput
				p := svc.NewListSandboxesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-sandboxes-for-project": {
			Name:   "list-sandboxes-for-project",
			Fields: fields_list_sandboxes_for_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSandboxesForProjectInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sandboxes_for_project, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSandboxesForProject(ctx, input)
				}
				var results []*svc.ListSandboxesForProjectOutput
				p := svc.NewListSandboxesForProjectPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-shared-projects": {
			Name:   "list-shared-projects",
			Fields: fields_list_shared_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSharedProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_shared_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSharedProjects(ctx, input)
				}
				var results []*svc.ListSharedProjectsOutput
				p := svc.NewListSharedProjectsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-shared-report-groups": {
			Name:   "list-shared-report-groups",
			Fields: fields_list_shared_report_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSharedReportGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_shared_report_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSharedReportGroups(ctx, input)
				}
				var results []*svc.ListSharedReportGroupsOutput
				p := svc.NewListSharedReportGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-source-credentials": {
			Name:   "list-source-credentials",
			Fields: fields_list_source_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_source_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSourceCredentials(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"retry-build": {
			Name:   "retry-build",
			Fields: fields_retry_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryBuild(ctx, input)
			},
		},
		"retry-build-batch": {
			Name:   "retry-build-batch",
			Fields: fields_retry_build_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryBuildBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_build_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryBuildBatch(ctx, input)
			},
		},
		"start-build": {
			Name:   "start-build",
			Fields: fields_start_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBuild(ctx, input)
			},
		},
		"start-build-batch": {
			Name:   "start-build-batch",
			Fields: fields_start_build_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBuildBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_build_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBuildBatch(ctx, input)
			},
		},
		"start-command-execution": {
			Name:   "start-command-execution",
			Fields: fields_start_command_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCommandExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_command_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCommandExecution(ctx, input)
			},
		},
		"start-sandbox": {
			Name:   "start-sandbox",
			Fields: fields_start_sandbox,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSandboxInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_sandbox, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSandbox(ctx, input)
			},
		},
		"start-sandbox-connection": {
			Name:   "start-sandbox-connection",
			Fields: fields_start_sandbox_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSandboxConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_sandbox_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSandboxConnection(ctx, input)
			},
		},
		"stop-build": {
			Name:   "stop-build",
			Fields: fields_stop_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBuild(ctx, input)
			},
		},
		"stop-build-batch": {
			Name:   "stop-build-batch",
			Fields: fields_stop_build_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBuildBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_build_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBuildBatch(ctx, input)
			},
		},
		"stop-sandbox": {
			Name:   "stop-sandbox",
			Fields: fields_stop_sandbox,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSandboxInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_sandbox, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSandbox(ctx, input)
			},
		},
		"update-fleet": {
			Name:   "update-fleet",
			Fields: fields_update_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleet(ctx, input)
			},
		},
		"update-project": {
			Name:   "update-project",
			Fields: fields_update_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProject(ctx, input)
			},
		},
		"update-project-visibility": {
			Name:   "update-project-visibility",
			Fields: fields_update_project_visibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectVisibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project_visibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProjectVisibility(ctx, input)
			},
		},
		"update-report-group": {
			Name:   "update-report-group",
			Fields: fields_update_report_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReportGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_report_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReportGroup(ctx, input)
			},
		},
		"update-webhook": {
			Name:   "update-webhook",
			Fields: fields_update_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWebhook(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codebuild", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
