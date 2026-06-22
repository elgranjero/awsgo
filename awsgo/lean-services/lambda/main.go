package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lambda"
)

var fields_add_layer_version_permission = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: true},
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_add_permission = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: true},
	{Name: "EventSourceToken", Flag: "event-source-token", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "FunctionUrlAuthType", Flag: "function-url-auth-type", Type: "types.FunctionUrlAuthType", Required: false},
	{Name: "InvokedViaFunctionUrl", Flag: "invoked-via-function-url", Type: "*bool", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "PrincipalOrgID", Flag: "principal-org-id", Type: "*string", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "SourceAccount", Flag: "source-account", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_checkpoint_durable_execution = []leanruntime.Field{
	{Name: "CheckpointToken", Flag: "checkpoint-token", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DurableExecutionArn", Flag: "durable-execution-arn", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.OperationUpdate", Required: false},
}

var fields_create_alias = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "FunctionVersion", Flag: "function-version", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoutingConfig", Flag: "routing-config", Type: "*types.AliasRoutingConfiguration", Required: false},
}

var fields_create_capacity_provider = []leanruntime.Field{
	{Name: "CapacityProviderName", Flag: "capacity-provider-name", Type: "*string", Required: true},
	{Name: "CapacityProviderScalingConfig", Flag: "capacity-provider-scaling-config", Type: "*types.CapacityProviderScalingConfig", Required: false},
	{Name: "InstanceRequirements", Flag: "instance-requirements", Type: "*types.InstanceRequirements", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "PermissionsConfig", Flag: "permissions-config", Type: "*types.CapacityProviderPermissionsConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.CapacityProviderVpcConfig", Required: true},
}

var fields_create_code_signing_config = []leanruntime.Field{
	{Name: "AllowedPublishers", Flag: "allowed-publishers", Type: "*types.AllowedPublishers", Required: true},
	{Name: "CodeSigningPolicies", Flag: "code-signing-policies", Type: "*types.CodeSigningPolicies", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_event_source_mapping = []leanruntime.Field{
	{Name: "AmazonManagedKafkaEventSourceConfig", Flag: "amazon-managed-kafka-event-source-config", Type: "*types.AmazonManagedKafkaEventSourceConfig", Required: false},
	{Name: "BatchSize", Flag: "batch-size", Type: "*int32", Required: false},
	{Name: "BisectBatchOnFunctionError", Flag: "bisect-batch-on-function-error", Type: "*bool", Required: false},
	{Name: "DestinationConfig", Flag: "destination-config", Type: "*types.DestinationConfig", Required: false},
	{Name: "DocumentDBEventSourceConfig", Flag: "document-db-event-source-config", Type: "*types.DocumentDBEventSourceConfig", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventSourceArn", Flag: "event-source-arn", Type: "*string", Required: false},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "FunctionResponseTypes", Flag: "function-response-types", Type: "[]types.FunctionResponseType", Required: false},
	{Name: "KMSKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.EventSourceMappingLoggingConfig", Required: false},
	{Name: "MaximumBatchingWindowInSeconds", Flag: "maximum-batching-window-in-seconds", Type: "*int32", Required: false},
	{Name: "MaximumRecordAgeInSeconds", Flag: "maximum-record-age-in-seconds", Type: "*int32", Required: false},
	{Name: "MaximumRetryAttempts", Flag: "maximum-retry-attempts", Type: "*int32", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "*types.EventSourceMappingMetricsConfig", Required: false},
	{Name: "ParallelizationFactor", Flag: "parallelization-factor", Type: "*int32", Required: false},
	{Name: "ProvisionedPollerConfig", Flag: "provisioned-poller-config", Type: "*types.ProvisionedPollerConfig", Required: false},
	{Name: "Queues", Flag: "queues", Type: "[]string", Required: false},
	{Name: "ScalingConfig", Flag: "scaling-config", Type: "*types.ScalingConfig", Required: false},
	{Name: "SelfManagedEventSource", Flag: "self-managed-event-source", Type: "*types.SelfManagedEventSource", Required: false},
	{Name: "SelfManagedKafkaEventSourceConfig", Flag: "self-managed-kafka-event-source-config", Type: "*types.SelfManagedKafkaEventSourceConfig", Required: false},
	{Name: "SourceAccessConfigurations", Flag: "source-access-configurations", Type: "[]types.SourceAccessConfiguration", Required: false},
	{Name: "StartingPosition", Flag: "starting-position", Type: "types.EventSourcePosition", Required: false},
	{Name: "StartingPositionTimestamp", Flag: "starting-position-timestamp", Type: "*time.Time", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Topics", Flag: "topics", Type: "[]string", Required: false},
	{Name: "TumblingWindowInSeconds", Flag: "tumbling-window-in-seconds", Type: "*int32", Required: false},
}

var fields_create_function = []leanruntime.Field{
	{Name: "Architectures", Flag: "architectures", Type: "[]types.Architecture", Required: false},
	{Name: "CapacityProviderConfig", Flag: "capacity-provider-config", Type: "*types.CapacityProviderConfig", Required: false},
	{Name: "Code", Flag: "code", Type: "*types.FunctionCode", Required: true},
	{Name: "CodeSigningConfigArn", Flag: "code-signing-config-arn", Type: "*string", Required: false},
	{Name: "DeadLetterConfig", Flag: "dead-letter-config", Type: "*types.DeadLetterConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DurableConfig", Flag: "durable-config", Type: "*types.DurableConfig", Required: false},
	{Name: "Environment", Flag: "environment", Type: "*types.Environment", Required: false},
	{Name: "EphemeralStorage", Flag: "ephemeral-storage", Type: "*types.EphemeralStorage", Required: false},
	{Name: "FileSystemConfigs", Flag: "file-system-configs", Type: "[]types.FileSystemConfig", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Handler", Flag: "handler", Type: "*string", Required: false},
	{Name: "ImageConfig", Flag: "image-config", Type: "*types.ImageConfig", Required: false},
	{Name: "KMSKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Layers", Flag: "layers", Type: "[]string", Required: false},
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.LoggingConfig", Required: false},
	{Name: "MemorySize", Flag: "memory-size", Type: "*int32", Required: false},
	{Name: "PackageType", Flag: "package-type", Type: "types.PackageType", Required: false},
	{Name: "Publish", Flag: "publish", Type: "bool", Required: false},
	{Name: "PublishTo", Flag: "publish-to", Type: "types.FunctionVersionLatestPublished", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Runtime", Flag: "runtime", Type: "types.Runtime", Required: false},
	{Name: "SnapStart", Flag: "snap-start", Type: "*types.SnapStart", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TenancyConfig", Flag: "tenancy-config", Type: "*types.TenancyConfig", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "TracingConfig", Flag: "tracing-config", Type: "*types.TracingConfig", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_function_url_config = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.FunctionUrlAuthType", Required: true},
	{Name: "Cors", Flag: "cors", Type: "*types.Cors", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "InvokeMode", Flag: "invoke-mode", Type: "types.InvokeMode", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_delete_alias = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_capacity_provider = []leanruntime.Field{
	{Name: "CapacityProviderName", Flag: "capacity-provider-name", Type: "*string", Required: true},
}

var fields_delete_code_signing_config = []leanruntime.Field{
	{Name: "CodeSigningConfigArn", Flag: "code-signing-config-arn", Type: "*string", Required: true},
}

var fields_delete_event_source_mapping = []leanruntime.Field{
	{Name: "UUID", Flag: "uuid", Type: "*string", Required: true},
}

var fields_delete_function = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_delete_function_code_signing_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_delete_function_concurrency = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_delete_function_event_invoke_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_delete_function_url_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_delete_layer_version = []leanruntime.Field{
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_delete_provisioned_concurrency_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: true},
}

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_alias = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_capacity_provider = []leanruntime.Field{
	{Name: "CapacityProviderName", Flag: "capacity-provider-name", Type: "*string", Required: true},
}

var fields_get_code_signing_config = []leanruntime.Field{
	{Name: "CodeSigningConfigArn", Flag: "code-signing-config-arn", Type: "*string", Required: true},
}

var fields_get_durable_execution = []leanruntime.Field{
	{Name: "DurableExecutionArn", Flag: "durable-execution-arn", Type: "*string", Required: true},
}

var fields_get_durable_execution_history = []leanruntime.Field{
	{Name: "DurableExecutionArn", Flag: "durable-execution-arn", Type: "*string", Required: true},
	{Name: "IncludeExecutionData", Flag: "include-execution-data", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "int32", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "*bool", Required: false},
}

var fields_get_durable_execution_state = []leanruntime.Field{
	{Name: "CheckpointToken", Flag: "checkpoint-token", Type: "*string", Required: true},
	{Name: "DurableExecutionArn", Flag: "durable-execution-arn", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "int32", Required: false},
}

var fields_get_event_source_mapping = []leanruntime.Field{
	{Name: "UUID", Flag: "uuid", Type: "*string", Required: true},
}

var fields_get_function = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_get_function_code_signing_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_get_function_concurrency = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_get_function_configuration = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_get_function_event_invoke_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_get_function_recursion_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_get_function_scaling_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: true},
}

var fields_get_function_url_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_get_layer_version = []leanruntime.Field{
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_get_layer_version_by_arn = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_layer_version_policy = []leanruntime.Field{
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_get_provisioned_concurrency_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: true},
}

var fields_get_runtime_management_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_invoke = []leanruntime.Field{
	{Name: "ClientContext", Flag: "client-context", Type: "*string", Required: false},
	{Name: "DurableExecutionName", Flag: "durable-execution-name", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "InvocationType", Flag: "invocation-type", Type: "types.InvocationType", Required: false},
	{Name: "LogType", Flag: "log-type", Type: "types.LogType", Required: false},
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "TenantId", Flag: "tenant-id", Type: "*string", Required: false},
}

var fields_invoke_async = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "InvokeArgs", Flag: "invoke-args", Type: "io.Reader", Required: true},
}

var fields_invoke_with_response_stream = []leanruntime.Field{
	{Name: "ClientContext", Flag: "client-context", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "InvocationType", Flag: "invocation-type", Type: "types.ResponseStreamingInvocationType", Required: false},
	{Name: "LogType", Flag: "log-type", Type: "types.LogType", Required: false},
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "TenantId", Flag: "tenant-id", Type: "*string", Required: false},
}

var fields_list_aliases = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "FunctionVersion", Flag: "function-version", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_capacity_providers = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "State", Flag: "state", Type: "types.CapacityProviderState", Required: false},
}

var fields_list_code_signing_configs = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_durable_executions_by_function = []leanruntime.Field{
	{Name: "DurableExecutionName", Flag: "durable-execution-name", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "int32", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "*bool", Required: false},
	{Name: "StartedAfter", Flag: "started-after", Type: "*time.Time", Required: false},
	{Name: "StartedBefore", Flag: "started-before", Type: "*time.Time", Required: false},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.ExecutionStatus", Required: false},
}

var fields_list_event_source_mappings = []leanruntime.Field{
	{Name: "EventSourceArn", Flag: "event-source-arn", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_function_event_invoke_configs = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_function_url_configs = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_function_versions_by_capacity_provider = []leanruntime.Field{
	{Name: "CapacityProviderName", Flag: "capacity-provider-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_functions = []leanruntime.Field{
	{Name: "FunctionVersion", Flag: "function-version", Type: "types.FunctionVersion", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MasterRegion", Flag: "master-region", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_functions_by_code_signing_config = []leanruntime.Field{
	{Name: "CodeSigningConfigArn", Flag: "code-signing-config-arn", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_layer_versions = []leanruntime.Field{
	{Name: "CompatibleArchitecture", Flag: "compatible-architecture", Type: "types.Architecture", Required: false},
	{Name: "CompatibleRuntime", Flag: "compatible-runtime", Type: "types.Runtime", Required: false},
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_layers = []leanruntime.Field{
	{Name: "CompatibleArchitecture", Flag: "compatible-architecture", Type: "types.Architecture", Required: false},
	{Name: "CompatibleRuntime", Flag: "compatible-runtime", Type: "types.Runtime", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_provisioned_concurrency_configs = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
}

var fields_list_versions_by_function = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_publish_layer_version = []leanruntime.Field{
	{Name: "CompatibleArchitectures", Flag: "compatible-architectures", Type: "[]types.Architecture", Required: false},
	{Name: "CompatibleRuntimes", Flag: "compatible-runtimes", Type: "[]types.Runtime", Required: false},
	{Name: "Content", Flag: "content", Type: "*types.LayerVersionContentInput", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "LicenseInfo", Flag: "license-info", Type: "*string", Required: false},
}

var fields_publish_version = []leanruntime.Field{
	{Name: "CodeSha256", Flag: "code-sha256", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "PublishTo", Flag: "publish-to", Type: "types.FunctionVersionLatestPublished", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
}

var fields_put_function_code_signing_config = []leanruntime.Field{
	{Name: "CodeSigningConfigArn", Flag: "code-signing-config-arn", Type: "*string", Required: true},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_put_function_concurrency = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "ReservedConcurrentExecutions", Flag: "reserved-concurrent-executions", Type: "*int32", Required: true},
}

var fields_put_function_event_invoke_config = []leanruntime.Field{
	{Name: "DestinationConfig", Flag: "destination-config", Type: "*types.DestinationConfig", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "MaximumEventAgeInSeconds", Flag: "maximum-event-age-in-seconds", Type: "*int32", Required: false},
	{Name: "MaximumRetryAttempts", Flag: "maximum-retry-attempts", Type: "*int32", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_put_function_recursion_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "RecursiveLoop", Flag: "recursive-loop", Type: "types.RecursiveLoop", Required: true},
}

var fields_put_function_scaling_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "FunctionScalingConfig", Flag: "function-scaling-config", Type: "*types.FunctionScalingConfig", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: true},
}

var fields_put_provisioned_concurrency_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "ProvisionedConcurrentExecutions", Flag: "provisioned-concurrent-executions", Type: "*int32", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: true},
}

var fields_put_runtime_management_config = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "RuntimeVersionArn", Flag: "runtime-version-arn", Type: "*string", Required: false},
	{Name: "UpdateRuntimeOn", Flag: "update-runtime-on", Type: "types.UpdateRuntimeOn", Required: true},
}

var fields_remove_layer_version_permission = []leanruntime.Field{
	{Name: "LayerName", Flag: "layer-name", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_remove_permission = []leanruntime.Field{
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_send_durable_execution_callback_failure = []leanruntime.Field{
	{Name: "CallbackId", Flag: "callback-id", Type: "*string", Required: true},
	{Name: "Error", Flag: "error", Type: "*types.ErrorObject", Required: false},
}

var fields_send_durable_execution_callback_heartbeat = []leanruntime.Field{
	{Name: "CallbackId", Flag: "callback-id", Type: "*string", Required: true},
}

var fields_send_durable_execution_callback_success = []leanruntime.Field{
	{Name: "CallbackId", Flag: "callback-id", Type: "*string", Required: true},
	{Name: "Result", Flag: "result", Type: "[]byte", Required: false},
}

var fields_stop_durable_execution = []leanruntime.Field{
	{Name: "DurableExecutionArn", Flag: "durable-execution-arn", Type: "*string", Required: true},
	{Name: "Error", Flag: "error", Type: "*types.ErrorObject", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_alias = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "FunctionVersion", Flag: "function-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "RoutingConfig", Flag: "routing-config", Type: "*types.AliasRoutingConfiguration", Required: false},
}

var fields_update_capacity_provider = []leanruntime.Field{
	{Name: "CapacityProviderName", Flag: "capacity-provider-name", Type: "*string", Required: true},
	{Name: "CapacityProviderScalingConfig", Flag: "capacity-provider-scaling-config", Type: "*types.CapacityProviderScalingConfig", Required: false},
}

var fields_update_code_signing_config = []leanruntime.Field{
	{Name: "AllowedPublishers", Flag: "allowed-publishers", Type: "*types.AllowedPublishers", Required: false},
	{Name: "CodeSigningConfigArn", Flag: "code-signing-config-arn", Type: "*string", Required: true},
	{Name: "CodeSigningPolicies", Flag: "code-signing-policies", Type: "*types.CodeSigningPolicies", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_event_source_mapping = []leanruntime.Field{
	{Name: "AmazonManagedKafkaEventSourceConfig", Flag: "amazon-managed-kafka-event-source-config", Type: "*types.AmazonManagedKafkaEventSourceConfig", Required: false},
	{Name: "BatchSize", Flag: "batch-size", Type: "*int32", Required: false},
	{Name: "BisectBatchOnFunctionError", Flag: "bisect-batch-on-function-error", Type: "*bool", Required: false},
	{Name: "DestinationConfig", Flag: "destination-config", Type: "*types.DestinationConfig", Required: false},
	{Name: "DocumentDBEventSourceConfig", Flag: "document-db-event-source-config", Type: "*types.DocumentDBEventSourceConfig", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: false},
	{Name: "FunctionResponseTypes", Flag: "function-response-types", Type: "[]types.FunctionResponseType", Required: false},
	{Name: "KMSKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.EventSourceMappingLoggingConfig", Required: false},
	{Name: "MaximumBatchingWindowInSeconds", Flag: "maximum-batching-window-in-seconds", Type: "*int32", Required: false},
	{Name: "MaximumRecordAgeInSeconds", Flag: "maximum-record-age-in-seconds", Type: "*int32", Required: false},
	{Name: "MaximumRetryAttempts", Flag: "maximum-retry-attempts", Type: "*int32", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "*types.EventSourceMappingMetricsConfig", Required: false},
	{Name: "ParallelizationFactor", Flag: "parallelization-factor", Type: "*int32", Required: false},
	{Name: "ProvisionedPollerConfig", Flag: "provisioned-poller-config", Type: "*types.ProvisionedPollerConfig", Required: false},
	{Name: "ScalingConfig", Flag: "scaling-config", Type: "*types.ScalingConfig", Required: false},
	{Name: "SelfManagedKafkaEventSourceConfig", Flag: "self-managed-kafka-event-source-config", Type: "*types.SelfManagedKafkaEventSourceConfig", Required: false},
	{Name: "SourceAccessConfigurations", Flag: "source-access-configurations", Type: "[]types.SourceAccessConfiguration", Required: false},
	{Name: "TumblingWindowInSeconds", Flag: "tumbling-window-in-seconds", Type: "*int32", Required: false},
	{Name: "UUID", Flag: "uuid", Type: "*string", Required: true},
}

var fields_update_function_code = []leanruntime.Field{
	{Name: "Architectures", Flag: "architectures", Type: "[]types.Architecture", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "ImageUri", Flag: "image-uri", Type: "*string", Required: false},
	{Name: "Publish", Flag: "publish", Type: "bool", Required: false},
	{Name: "PublishTo", Flag: "publish-to", Type: "types.FunctionVersionLatestPublished", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: false},
	{Name: "S3Key", Flag: "s3-key", Type: "*string", Required: false},
	{Name: "S3ObjectVersion", Flag: "s3-object-version", Type: "*string", Required: false},
	{Name: "SourceKMSKeyArn", Flag: "source-kms-key-arn", Type: "*string", Required: false},
	{Name: "ZipFile", Flag: "zip-file", Type: "[]byte", Required: false},
}

var fields_update_function_configuration = []leanruntime.Field{
	{Name: "CapacityProviderConfig", Flag: "capacity-provider-config", Type: "*types.CapacityProviderConfig", Required: false},
	{Name: "DeadLetterConfig", Flag: "dead-letter-config", Type: "*types.DeadLetterConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DurableConfig", Flag: "durable-config", Type: "*types.DurableConfig", Required: false},
	{Name: "Environment", Flag: "environment", Type: "*types.Environment", Required: false},
	{Name: "EphemeralStorage", Flag: "ephemeral-storage", Type: "*types.EphemeralStorage", Required: false},
	{Name: "FileSystemConfigs", Flag: "file-system-configs", Type: "[]types.FileSystemConfig", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "Handler", Flag: "handler", Type: "*string", Required: false},
	{Name: "ImageConfig", Flag: "image-config", Type: "*types.ImageConfig", Required: false},
	{Name: "KMSKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Layers", Flag: "layers", Type: "[]string", Required: false},
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.LoggingConfig", Required: false},
	{Name: "MemorySize", Flag: "memory-size", Type: "*int32", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "types.Runtime", Required: false},
	{Name: "SnapStart", Flag: "snap-start", Type: "*types.SnapStart", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "TracingConfig", Flag: "tracing-config", Type: "*types.TracingConfig", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_update_function_event_invoke_config = []leanruntime.Field{
	{Name: "DestinationConfig", Flag: "destination-config", Type: "*types.DestinationConfig", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "MaximumEventAgeInSeconds", Flag: "maximum-event-age-in-seconds", Type: "*int32", Required: false},
	{Name: "MaximumRetryAttempts", Flag: "maximum-retry-attempts", Type: "*int32", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

var fields_update_function_url_config = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.FunctionUrlAuthType", Required: false},
	{Name: "Cors", Flag: "cors", Type: "*types.Cors", Required: false},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
	{Name: "InvokeMode", Flag: "invoke-mode", Type: "types.InvokeMode", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-layer-version-permission": {
			Name:   "add-layer-version-permission",
			Fields: fields_add_layer_version_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddLayerVersionPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_layer_version_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddLayerVersionPermission(ctx, input)
			},
		},
		"add-permission": {
			Name:   "add-permission",
			Fields: fields_add_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddPermission(ctx, input)
			},
		},
		"checkpoint-durable-execution": {
			Name:   "checkpoint-durable-execution",
			Fields: fields_checkpoint_durable_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckpointDurableExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_checkpoint_durable_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckpointDurableExecution(ctx, input)
			},
		},
		"create-alias": {
			Name:   "create-alias",
			Fields: fields_create_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlias(ctx, input)
			},
		},
		"create-capacity-provider": {
			Name:   "create-capacity-provider",
			Fields: fields_create_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityProvider(ctx, input)
			},
		},
		"create-code-signing-config": {
			Name:   "create-code-signing-config",
			Fields: fields_create_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCodeSigningConfig(ctx, input)
			},
		},
		"create-event-source-mapping": {
			Name:   "create-event-source-mapping",
			Fields: fields_create_event_source_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventSourceMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_source_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventSourceMapping(ctx, input)
			},
		},
		"create-function": {
			Name:   "create-function",
			Fields: fields_create_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFunction(ctx, input)
			},
		},
		"create-function-url-config": {
			Name:   "create-function-url-config",
			Fields: fields_create_function_url_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFunctionUrlConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_function_url_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFunctionUrlConfig(ctx, input)
			},
		},
		"delete-alias": {
			Name:   "delete-alias",
			Fields: fields_delete_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlias(ctx, input)
			},
		},
		"delete-capacity-provider": {
			Name:   "delete-capacity-provider",
			Fields: fields_delete_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCapacityProvider(ctx, input)
			},
		},
		"delete-code-signing-config": {
			Name:   "delete-code-signing-config",
			Fields: fields_delete_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCodeSigningConfig(ctx, input)
			},
		},
		"delete-event-source-mapping": {
			Name:   "delete-event-source-mapping",
			Fields: fields_delete_event_source_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventSourceMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_source_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventSourceMapping(ctx, input)
			},
		},
		"delete-function": {
			Name:   "delete-function",
			Fields: fields_delete_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFunction(ctx, input)
			},
		},
		"delete-function-code-signing-config": {
			Name:   "delete-function-code-signing-config",
			Fields: fields_delete_function_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFunctionCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_function_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFunctionCodeSigningConfig(ctx, input)
			},
		},
		"delete-function-concurrency": {
			Name:   "delete-function-concurrency",
			Fields: fields_delete_function_concurrency,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFunctionConcurrencyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_function_concurrency, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFunctionConcurrency(ctx, input)
			},
		},
		"delete-function-event-invoke-config": {
			Name:   "delete-function-event-invoke-config",
			Fields: fields_delete_function_event_invoke_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFunctionEventInvokeConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_function_event_invoke_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFunctionEventInvokeConfig(ctx, input)
			},
		},
		"delete-function-url-config": {
			Name:   "delete-function-url-config",
			Fields: fields_delete_function_url_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFunctionUrlConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_function_url_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFunctionUrlConfig(ctx, input)
			},
		},
		"delete-layer-version": {
			Name:   "delete-layer-version",
			Fields: fields_delete_layer_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLayerVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_layer_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLayerVersion(ctx, input)
			},
		},
		"delete-provisioned-concurrency-config": {
			Name:   "delete-provisioned-concurrency-config",
			Fields: fields_delete_provisioned_concurrency_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisionedConcurrencyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioned_concurrency_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisionedConcurrencyConfig(ctx, input)
			},
		},
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
			},
		},
		"get-alias": {
			Name:   "get-alias",
			Fields: fields_get_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAlias(ctx, input)
			},
		},
		"get-capacity-provider": {
			Name:   "get-capacity-provider",
			Fields: fields_get_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapacityProvider(ctx, input)
			},
		},
		"get-code-signing-config": {
			Name:   "get-code-signing-config",
			Fields: fields_get_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeSigningConfig(ctx, input)
			},
		},
		"get-durable-execution": {
			Name:   "get-durable-execution",
			Fields: fields_get_durable_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDurableExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_durable_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDurableExecution(ctx, input)
			},
		},
		"get-durable-execution-history": {
			Name:   "get-durable-execution-history",
			Fields: fields_get_durable_execution_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDurableExecutionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_durable_execution_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDurableExecutionHistory(ctx, input)
				}
				var results []*svc.GetDurableExecutionHistoryOutput
				p := svc.NewGetDurableExecutionHistoryPaginator(client, input)
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
		"get-durable-execution-state": {
			Name:   "get-durable-execution-state",
			Fields: fields_get_durable_execution_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDurableExecutionStateInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_durable_execution_state, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDurableExecutionState(ctx, input)
				}
				var results []*svc.GetDurableExecutionStateOutput
				p := svc.NewGetDurableExecutionStatePaginator(client, input)
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
		"get-event-source-mapping": {
			Name:   "get-event-source-mapping",
			Fields: fields_get_event_source_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventSourceMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_source_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventSourceMapping(ctx, input)
			},
		},
		"get-function": {
			Name:   "get-function",
			Fields: fields_get_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunction(ctx, input)
			},
		},
		"get-function-code-signing-config": {
			Name:   "get-function-code-signing-config",
			Fields: fields_get_function_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionCodeSigningConfig(ctx, input)
			},
		},
		"get-function-concurrency": {
			Name:   "get-function-concurrency",
			Fields: fields_get_function_concurrency,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionConcurrencyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_concurrency, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionConcurrency(ctx, input)
			},
		},
		"get-function-configuration": {
			Name:   "get-function-configuration",
			Fields: fields_get_function_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionConfiguration(ctx, input)
			},
		},
		"get-function-event-invoke-config": {
			Name:   "get-function-event-invoke-config",
			Fields: fields_get_function_event_invoke_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionEventInvokeConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_event_invoke_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionEventInvokeConfig(ctx, input)
			},
		},
		"get-function-recursion-config": {
			Name:   "get-function-recursion-config",
			Fields: fields_get_function_recursion_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionRecursionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_recursion_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionRecursionConfig(ctx, input)
			},
		},
		"get-function-scaling-config": {
			Name:   "get-function-scaling-config",
			Fields: fields_get_function_scaling_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionScalingConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_scaling_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionScalingConfig(ctx, input)
			},
		},
		"get-function-url-config": {
			Name:   "get-function-url-config",
			Fields: fields_get_function_url_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionUrlConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_url_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionUrlConfig(ctx, input)
			},
		},
		"get-layer-version": {
			Name:   "get-layer-version",
			Fields: fields_get_layer_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLayerVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_layer_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLayerVersion(ctx, input)
			},
		},
		"get-layer-version-by-arn": {
			Name:   "get-layer-version-by-arn",
			Fields: fields_get_layer_version_by_arn,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLayerVersionByArnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_layer_version_by_arn, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLayerVersionByArn(ctx, input)
			},
		},
		"get-layer-version-policy": {
			Name:   "get-layer-version-policy",
			Fields: fields_get_layer_version_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLayerVersionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_layer_version_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLayerVersionPolicy(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-provisioned-concurrency-config": {
			Name:   "get-provisioned-concurrency-config",
			Fields: fields_get_provisioned_concurrency_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProvisionedConcurrencyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_provisioned_concurrency_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProvisionedConcurrencyConfig(ctx, input)
			},
		},
		"get-runtime-management-config": {
			Name:   "get-runtime-management-config",
			Fields: fields_get_runtime_management_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuntimeManagementConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_runtime_management_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRuntimeManagementConfig(ctx, input)
			},
		},
		"invoke": {
			Name:   "invoke",
			Fields: fields_invoke,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Invoke(ctx, input)
			},
		},
		"invoke-async": {
			Name:   "invoke-async",
			Fields: fields_invoke_async,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeAsyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_async, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeAsync(ctx, input)
			},
		},
		"invoke-with-response-stream": {
			Name:   "invoke-with-response-stream",
			Fields: fields_invoke_with_response_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeWithResponseStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_with_response_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeWithResponseStream(ctx, input)
			},
		},
		"list-aliases": {
			Name:   "list-aliases",
			Fields: fields_list_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAliases(ctx, input)
				}
				var results []*svc.ListAliasesOutput
				p := svc.NewListAliasesPaginator(client, input)
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
		"list-capacity-providers": {
			Name:   "list-capacity-providers",
			Fields: fields_list_capacity_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCapacityProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_capacity_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCapacityProviders(ctx, input)
				}
				var results []*svc.ListCapacityProvidersOutput
				p := svc.NewListCapacityProvidersPaginator(client, input)
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
		"list-code-signing-configs": {
			Name:   "list-code-signing-configs",
			Fields: fields_list_code_signing_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeSigningConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_code_signing_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCodeSigningConfigs(ctx, input)
				}
				var results []*svc.ListCodeSigningConfigsOutput
				p := svc.NewListCodeSigningConfigsPaginator(client, input)
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
		"list-durable-executions-by-function": {
			Name:   "list-durable-executions-by-function",
			Fields: fields_list_durable_executions_by_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDurableExecutionsByFunctionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_durable_executions_by_function, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDurableExecutionsByFunction(ctx, input)
				}
				var results []*svc.ListDurableExecutionsByFunctionOutput
				p := svc.NewListDurableExecutionsByFunctionPaginator(client, input)
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
		"list-event-source-mappings": {
			Name:   "list-event-source-mappings",
			Fields: fields_list_event_source_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventSourceMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_source_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventSourceMappings(ctx, input)
				}
				var results []*svc.ListEventSourceMappingsOutput
				p := svc.NewListEventSourceMappingsPaginator(client, input)
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
		"list-function-event-invoke-configs": {
			Name:   "list-function-event-invoke-configs",
			Fields: fields_list_function_event_invoke_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionEventInvokeConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_function_event_invoke_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFunctionEventInvokeConfigs(ctx, input)
				}
				var results []*svc.ListFunctionEventInvokeConfigsOutput
				p := svc.NewListFunctionEventInvokeConfigsPaginator(client, input)
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
		"list-function-url-configs": {
			Name:   "list-function-url-configs",
			Fields: fields_list_function_url_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionUrlConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_function_url_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFunctionUrlConfigs(ctx, input)
				}
				var results []*svc.ListFunctionUrlConfigsOutput
				p := svc.NewListFunctionUrlConfigsPaginator(client, input)
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
		"list-function-versions-by-capacity-provider": {
			Name:   "list-function-versions-by-capacity-provider",
			Fields: fields_list_function_versions_by_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionVersionsByCapacityProviderInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_function_versions_by_capacity_provider, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFunctionVersionsByCapacityProvider(ctx, input)
				}
				var results []*svc.ListFunctionVersionsByCapacityProviderOutput
				p := svc.NewListFunctionVersionsByCapacityProviderPaginator(client, input)
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
		"list-functions": {
			Name:   "list-functions",
			Fields: fields_list_functions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_functions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFunctions(ctx, input)
				}
				var results []*svc.ListFunctionsOutput
				p := svc.NewListFunctionsPaginator(client, input)
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
		"list-functions-by-code-signing-config": {
			Name:   "list-functions-by-code-signing-config",
			Fields: fields_list_functions_by_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionsByCodeSigningConfigInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_functions_by_code_signing_config, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFunctionsByCodeSigningConfig(ctx, input)
				}
				var results []*svc.ListFunctionsByCodeSigningConfigOutput
				p := svc.NewListFunctionsByCodeSigningConfigPaginator(client, input)
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
		"list-layer-versions": {
			Name:   "list-layer-versions",
			Fields: fields_list_layer_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLayerVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_layer_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLayerVersions(ctx, input)
				}
				var results []*svc.ListLayerVersionsOutput
				p := svc.NewListLayerVersionsPaginator(client, input)
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
		"list-layers": {
			Name:   "list-layers",
			Fields: fields_list_layers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLayersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_layers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLayers(ctx, input)
				}
				var results []*svc.ListLayersOutput
				p := svc.NewListLayersPaginator(client, input)
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
		"list-provisioned-concurrency-configs": {
			Name:   "list-provisioned-concurrency-configs",
			Fields: fields_list_provisioned_concurrency_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisionedConcurrencyConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provisioned_concurrency_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProvisionedConcurrencyConfigs(ctx, input)
				}
				var results []*svc.ListProvisionedConcurrencyConfigsOutput
				p := svc.NewListProvisionedConcurrencyConfigsPaginator(client, input)
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
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTags(ctx, input)
			},
		},
		"list-versions-by-function": {
			Name:   "list-versions-by-function",
			Fields: fields_list_versions_by_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVersionsByFunctionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_versions_by_function, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVersionsByFunction(ctx, input)
				}
				var results []*svc.ListVersionsByFunctionOutput
				p := svc.NewListVersionsByFunctionPaginator(client, input)
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
		"publish-layer-version": {
			Name:   "publish-layer-version",
			Fields: fields_publish_layer_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishLayerVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_layer_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishLayerVersion(ctx, input)
			},
		},
		"publish-version": {
			Name:   "publish-version",
			Fields: fields_publish_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishVersion(ctx, input)
			},
		},
		"put-function-code-signing-config": {
			Name:   "put-function-code-signing-config",
			Fields: fields_put_function_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFunctionCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_function_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFunctionCodeSigningConfig(ctx, input)
			},
		},
		"put-function-concurrency": {
			Name:   "put-function-concurrency",
			Fields: fields_put_function_concurrency,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFunctionConcurrencyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_function_concurrency, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFunctionConcurrency(ctx, input)
			},
		},
		"put-function-event-invoke-config": {
			Name:   "put-function-event-invoke-config",
			Fields: fields_put_function_event_invoke_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFunctionEventInvokeConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_function_event_invoke_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFunctionEventInvokeConfig(ctx, input)
			},
		},
		"put-function-recursion-config": {
			Name:   "put-function-recursion-config",
			Fields: fields_put_function_recursion_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFunctionRecursionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_function_recursion_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFunctionRecursionConfig(ctx, input)
			},
		},
		"put-function-scaling-config": {
			Name:   "put-function-scaling-config",
			Fields: fields_put_function_scaling_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFunctionScalingConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_function_scaling_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFunctionScalingConfig(ctx, input)
			},
		},
		"put-provisioned-concurrency-config": {
			Name:   "put-provisioned-concurrency-config",
			Fields: fields_put_provisioned_concurrency_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProvisionedConcurrencyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_provisioned_concurrency_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProvisionedConcurrencyConfig(ctx, input)
			},
		},
		"put-runtime-management-config": {
			Name:   "put-runtime-management-config",
			Fields: fields_put_runtime_management_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRuntimeManagementConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_runtime_management_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRuntimeManagementConfig(ctx, input)
			},
		},
		"remove-layer-version-permission": {
			Name:   "remove-layer-version-permission",
			Fields: fields_remove_layer_version_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveLayerVersionPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_layer_version_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveLayerVersionPermission(ctx, input)
			},
		},
		"remove-permission": {
			Name:   "remove-permission",
			Fields: fields_remove_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemovePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemovePermission(ctx, input)
			},
		},
		"send-durable-execution-callback-failure": {
			Name:   "send-durable-execution-callback-failure",
			Fields: fields_send_durable_execution_callback_failure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDurableExecutionCallbackFailureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_durable_execution_callback_failure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDurableExecutionCallbackFailure(ctx, input)
			},
		},
		"send-durable-execution-callback-heartbeat": {
			Name:   "send-durable-execution-callback-heartbeat",
			Fields: fields_send_durable_execution_callback_heartbeat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDurableExecutionCallbackHeartbeatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_durable_execution_callback_heartbeat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDurableExecutionCallbackHeartbeat(ctx, input)
			},
		},
		"send-durable-execution-callback-success": {
			Name:   "send-durable-execution-callback-success",
			Fields: fields_send_durable_execution_callback_success,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDurableExecutionCallbackSuccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_durable_execution_callback_success, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDurableExecutionCallbackSuccess(ctx, input)
			},
		},
		"stop-durable-execution": {
			Name:   "stop-durable-execution",
			Fields: fields_stop_durable_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDurableExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_durable_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDurableExecution(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-alias": {
			Name:   "update-alias",
			Fields: fields_update_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAlias(ctx, input)
			},
		},
		"update-capacity-provider": {
			Name:   "update-capacity-provider",
			Fields: fields_update_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCapacityProvider(ctx, input)
			},
		},
		"update-code-signing-config": {
			Name:   "update-code-signing-config",
			Fields: fields_update_code_signing_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCodeSigningConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_code_signing_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCodeSigningConfig(ctx, input)
			},
		},
		"update-event-source-mapping": {
			Name:   "update-event-source-mapping",
			Fields: fields_update_event_source_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventSourceMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_source_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventSourceMapping(ctx, input)
			},
		},
		"update-function-code": {
			Name:   "update-function-code",
			Fields: fields_update_function_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFunctionCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_function_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFunctionCode(ctx, input)
			},
		},
		"update-function-configuration": {
			Name:   "update-function-configuration",
			Fields: fields_update_function_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFunctionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_function_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFunctionConfiguration(ctx, input)
			},
		},
		"update-function-event-invoke-config": {
			Name:   "update-function-event-invoke-config",
			Fields: fields_update_function_event_invoke_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFunctionEventInvokeConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_function_event_invoke_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFunctionEventInvokeConfig(ctx, input)
			},
		},
		"update-function-url-config": {
			Name:   "update-function-url-config",
			Fields: fields_update_function_url_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFunctionUrlConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_function_url_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFunctionUrlConfig(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lambda", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
