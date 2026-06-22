package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol"
)

var fields_create_agent_runtime = []leanruntime.Field{
	{Name: "AgentRuntimeArtifact", Flag: "agent-runtime-artifact", Type: "types.AgentRuntimeArtifact", Required: true},
	{Name: "AgentRuntimeName", Flag: "agent-runtime-name", Type: "*string", Required: true},
	{Name: "AuthorizerConfiguration", Flag: "authorizer-configuration", Type: "types.AuthorizerConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "LifecycleConfiguration", Flag: "lifecycle-configuration", Type: "*types.LifecycleConfiguration", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: true},
	{Name: "ProtocolConfiguration", Flag: "protocol-configuration", Type: "*types.ProtocolConfiguration", Required: false},
	{Name: "RequestHeaderConfiguration", Flag: "request-header-configuration", Type: "types.RequestHeaderConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_agent_runtime_endpoint = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "AgentRuntimeVersion", Flag: "agent-runtime-version", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_api_key_credential_provider = []leanruntime.Field{
	{Name: "ApiKey", Flag: "api-key", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_browser = []leanruntime.Field{
	{Name: "BrowserSigning", Flag: "browser-signing", Type: "*types.BrowserSigningConfigInput", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.BrowserNetworkConfiguration", Required: true},
	{Name: "Recording", Flag: "recording", Type: "*types.RecordingConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_browser_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_code_interpreter = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.CodeInterpreterNetworkConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_evaluator = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EvaluatorConfig", Flag: "evaluator-config", Type: "types.EvaluatorConfig", Required: true},
	{Name: "EvaluatorName", Flag: "evaluator-name", Type: "*string", Required: true},
	{Name: "Level", Flag: "level", Type: "types.EvaluatorLevel", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_gateway = []leanruntime.Field{
	{Name: "AuthorizerConfiguration", Flag: "authorizer-configuration", Type: "types.AuthorizerConfiguration", Required: false},
	{Name: "AuthorizerType", Flag: "authorizer-type", Type: "types.AuthorizerType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExceptionLevel", Flag: "exception-level", Type: "types.ExceptionLevel", Required: false},
	{Name: "InterceptorConfigurations", Flag: "interceptor-configurations", Type: "[]types.GatewayInterceptorConfiguration", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyEngineConfiguration", Flag: "policy-engine-configuration", Type: "*types.GatewayPolicyEngineConfiguration", Required: false},
	{Name: "ProtocolConfiguration", Flag: "protocol-configuration", Type: "types.GatewayProtocolConfiguration", Required: false},
	{Name: "ProtocolType", Flag: "protocol-type", Type: "types.GatewayProtocolType", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_gateway_target = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CredentialProviderConfigurations", Flag: "credential-provider-configurations", Type: "[]types.CredentialProviderConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "MetadataConfiguration", Flag: "metadata-configuration", Type: "*types.MetadataConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TargetConfiguration", Flag: "target-configuration", Type: "types.TargetConfiguration", Required: true},
}

var fields_create_memory = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "EventExpiryDuration", Flag: "event-expiry-duration", Type: "*int32", Required: true},
	{Name: "MemoryExecutionRoleArn", Flag: "memory-execution-role-arn", Type: "*string", Required: false},
	{Name: "MemoryStrategies", Flag: "memory-strategies", Type: "[]types.MemoryStrategyInput", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_oauth2_credential_provider = []leanruntime.Field{
	{Name: "CredentialProviderVendor", Flag: "credential-provider-vendor", Type: "types.CredentialProviderVendorType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Oauth2ProviderConfigInput", Flag: "oauth2-provider-config-input", Type: "types.Oauth2ProviderConfigInput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_online_evaluation_config = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSourceConfig", Flag: "data-source-config", Type: "types.DataSourceConfig", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnableOnCreate", Flag: "enable-on-create", Type: "*bool", Required: true},
	{Name: "EvaluationExecutionRoleArn", Flag: "evaluation-execution-role-arn", Type: "*string", Required: true},
	{Name: "Evaluators", Flag: "evaluators", Type: "[]types.EvaluatorReference", Required: true},
	{Name: "OnlineEvaluationConfigName", Flag: "online-evaluation-config-name", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*types.Rule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "types.PolicyDefinition", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "ValidationMode", Flag: "validation-mode", Type: "types.PolicyValidationMode", Required: false},
}

var fields_create_policy_engine = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_workload_identity = []leanruntime.Field{
	{Name: "AllowedResourceOauth2ReturnUrls", Flag: "allowed-resource-oauth2-return-urls", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_agent_runtime = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_agent_runtime_endpoint = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_delete_api_key_credential_provider = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_browser = []leanruntime.Field{
	{Name: "BrowserId", Flag: "browser-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_browser_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_delete_code_interpreter = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CodeInterpreterId", Flag: "code-interpreter-id", Type: "*string", Required: true},
}

var fields_delete_evaluator = []leanruntime.Field{
	{Name: "EvaluatorId", Flag: "evaluator-id", Type: "*string", Required: true},
}

var fields_delete_gateway = []leanruntime.Field{
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
}

var fields_delete_gateway_target = []leanruntime.Field{
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_delete_memory = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
}

var fields_delete_oauth2_credential_provider = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_online_evaluation_config = []leanruntime.Field{
	{Name: "OnlineEvaluationConfigId", Flag: "online-evaluation-config-id", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_delete_policy_engine = []leanruntime.Field{
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_workload_identity = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_agent_runtime = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "AgentRuntimeVersion", Flag: "agent-runtime-version", Type: "*string", Required: false},
}

var fields_get_agent_runtime_endpoint = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_get_api_key_credential_provider = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_browser = []leanruntime.Field{
	{Name: "BrowserId", Flag: "browser-id", Type: "*string", Required: true},
}

var fields_get_browser_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_get_code_interpreter = []leanruntime.Field{
	{Name: "CodeInterpreterId", Flag: "code-interpreter-id", Type: "*string", Required: true},
}

var fields_get_evaluator = []leanruntime.Field{
	{Name: "EvaluatorId", Flag: "evaluator-id", Type: "*string", Required: true},
}

var fields_get_gateway = []leanruntime.Field{
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
}

var fields_get_gateway_target = []leanruntime.Field{
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_get_memory = []leanruntime.Field{
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "View", Flag: "view", Type: "types.MemoryView", Required: false},
}

var fields_get_oauth2_credential_provider = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_online_evaluation_config = []leanruntime.Field{
	{Name: "OnlineEvaluationConfigId", Flag: "online-evaluation-config-id", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_get_policy_engine = []leanruntime.Field{
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
}

var fields_get_policy_generation = []leanruntime.Field{
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "PolicyGenerationId", Flag: "policy-generation-id", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_token_vault = []leanruntime.Field{
	{Name: "TokenVaultId", Flag: "token-vault-id", Type: "*string", Required: false},
}

var fields_get_workload_identity = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_agent_runtime_endpoints = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agent_runtime_versions = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agent_runtimes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_api_key_credential_providers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_browser_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_browsers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ResourceType", Required: false},
}

var fields_list_code_interpreters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ResourceType", Required: false},
}

var fields_list_evaluators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_gateway_targets = []leanruntime.Field{
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_memories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_oauth2_credential_providers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_online_evaluation_configs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "TargetResourceScope", Flag: "target-resource-scope", Type: "*string", Required: false},
}

var fields_list_policy_engines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policy_generation_assets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "PolicyGenerationId", Flag: "policy-generation-id", Type: "*string", Required: true},
}

var fields_list_policy_generations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_workload_identities = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_set_token_vault_cmk = []leanruntime.Field{
	{Name: "KmsConfiguration", Flag: "kms-configuration", Type: "*types.KmsConfiguration", Required: true},
	{Name: "TokenVaultId", Flag: "token-vault-id", Type: "*string", Required: false},
}

var fields_start_policy_generation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "types.Content", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "types.Resource", Required: true},
}

var fields_synchronize_gateway_targets = []leanruntime.Field{
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "TargetIdList", Flag: "target-id-list", Type: "[]string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_agent_runtime = []leanruntime.Field{
	{Name: "AgentRuntimeArtifact", Flag: "agent-runtime-artifact", Type: "types.AgentRuntimeArtifact", Required: true},
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "AuthorizerConfiguration", Flag: "authorizer-configuration", Type: "types.AuthorizerConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "LifecycleConfiguration", Flag: "lifecycle-configuration", Type: "*types.LifecycleConfiguration", Required: false},
	{Name: "MetadataConfiguration", Flag: "metadata-configuration", Type: "*types.RuntimeMetadataConfiguration", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: true},
	{Name: "ProtocolConfiguration", Flag: "protocol-configuration", Type: "*types.ProtocolConfiguration", Required: false},
	{Name: "RequestHeaderConfiguration", Flag: "request-header-configuration", Type: "types.RequestHeaderConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_update_agent_runtime_endpoint = []leanruntime.Field{
	{Name: "AgentRuntimeId", Flag: "agent-runtime-id", Type: "*string", Required: true},
	{Name: "AgentRuntimeVersion", Flag: "agent-runtime-version", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_update_api_key_credential_provider = []leanruntime.Field{
	{Name: "ApiKey", Flag: "api-key", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_evaluator = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EvaluatorConfig", Flag: "evaluator-config", Type: "types.EvaluatorConfig", Required: false},
	{Name: "EvaluatorId", Flag: "evaluator-id", Type: "*string", Required: true},
	{Name: "Level", Flag: "level", Type: "types.EvaluatorLevel", Required: false},
}

var fields_update_gateway = []leanruntime.Field{
	{Name: "AuthorizerConfiguration", Flag: "authorizer-configuration", Type: "types.AuthorizerConfiguration", Required: false},
	{Name: "AuthorizerType", Flag: "authorizer-type", Type: "types.AuthorizerType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExceptionLevel", Flag: "exception-level", Type: "types.ExceptionLevel", Required: false},
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "InterceptorConfigurations", Flag: "interceptor-configurations", Type: "[]types.GatewayInterceptorConfiguration", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyEngineConfiguration", Flag: "policy-engine-configuration", Type: "*types.GatewayPolicyEngineConfiguration", Required: false},
	{Name: "ProtocolConfiguration", Flag: "protocol-configuration", Type: "types.GatewayProtocolConfiguration", Required: false},
	{Name: "ProtocolType", Flag: "protocol-type", Type: "types.GatewayProtocolType", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_update_gateway_target = []leanruntime.Field{
	{Name: "CredentialProviderConfigurations", Flag: "credential-provider-configurations", Type: "[]types.CredentialProviderConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GatewayIdentifier", Flag: "gateway-identifier", Type: "*string", Required: true},
	{Name: "MetadataConfiguration", Flag: "metadata-configuration", Type: "*types.MetadataConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TargetConfiguration", Flag: "target-configuration", Type: "types.TargetConfiguration", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_update_memory = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventExpiryDuration", Flag: "event-expiry-duration", Type: "*int32", Required: false},
	{Name: "MemoryExecutionRoleArn", Flag: "memory-execution-role-arn", Type: "*string", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "MemoryStrategies", Flag: "memory-strategies", Type: "*types.ModifyMemoryStrategies", Required: false},
}

var fields_update_oauth2_credential_provider = []leanruntime.Field{
	{Name: "CredentialProviderVendor", Flag: "credential-provider-vendor", Type: "types.CredentialProviderVendorType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Oauth2ProviderConfigInput", Flag: "oauth2-provider-config-input", Type: "types.Oauth2ProviderConfigInput", Required: true},
}

var fields_update_online_evaluation_config = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSourceConfig", Flag: "data-source-config", Type: "types.DataSourceConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EvaluationExecutionRoleArn", Flag: "evaluation-execution-role-arn", Type: "*string", Required: false},
	{Name: "Evaluators", Flag: "evaluators", Type: "[]types.EvaluatorReference", Required: false},
	{Name: "ExecutionStatus", Flag: "execution-status", Type: "types.OnlineEvaluationExecutionStatus", Required: false},
	{Name: "OnlineEvaluationConfigId", Flag: "online-evaluation-config-id", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*types.Rule", Required: false},
}

var fields_update_policy = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "types.PolicyDefinition", Required: false},
	{Name: "Description", Flag: "description", Type: "*types.UpdatedDescription", Required: false},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "ValidationMode", Flag: "validation-mode", Type: "types.PolicyValidationMode", Required: false},
}

var fields_update_policy_engine = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*types.UpdatedDescription", Required: false},
	{Name: "PolicyEngineId", Flag: "policy-engine-id", Type: "*string", Required: true},
}

var fields_update_workload_identity = []leanruntime.Field{
	{Name: "AllowedResourceOauth2ReturnUrls", Flag: "allowed-resource-oauth2-return-urls", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-agent-runtime": {
			Name:   "create-agent-runtime",
			Fields: fields_create_agent_runtime,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentRuntimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent_runtime, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgentRuntime(ctx, input)
			},
		},
		"create-agent-runtime-endpoint": {
			Name:   "create-agent-runtime-endpoint",
			Fields: fields_create_agent_runtime_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentRuntimeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent_runtime_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgentRuntimeEndpoint(ctx, input)
			},
		},
		"create-api-key-credential-provider": {
			Name:   "create-api-key-credential-provider",
			Fields: fields_create_api_key_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApiKeyCredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api_key_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApiKeyCredentialProvider(ctx, input)
			},
		},
		"create-browser": {
			Name:   "create-browser",
			Fields: fields_create_browser,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBrowserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_browser, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBrowser(ctx, input)
			},
		},
		"create-browser-profile": {
			Name:   "create-browser-profile",
			Fields: fields_create_browser_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBrowserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_browser_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBrowserProfile(ctx, input)
			},
		},
		"create-code-interpreter": {
			Name:   "create-code-interpreter",
			Fields: fields_create_code_interpreter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCodeInterpreterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_code_interpreter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCodeInterpreter(ctx, input)
			},
		},
		"create-evaluator": {
			Name:   "create-evaluator",
			Fields: fields_create_evaluator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEvaluatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_evaluator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEvaluator(ctx, input)
			},
		},
		"create-gateway": {
			Name:   "create-gateway",
			Fields: fields_create_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGateway(ctx, input)
			},
		},
		"create-gateway-target": {
			Name:   "create-gateway-target",
			Fields: fields_create_gateway_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGatewayTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gateway_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGatewayTarget(ctx, input)
			},
		},
		"create-memory": {
			Name:   "create-memory",
			Fields: fields_create_memory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMemoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_memory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMemory(ctx, input)
			},
		},
		"create-oauth2-credential-provider": {
			Name:   "create-oauth2-credential-provider",
			Fields: fields_create_oauth2_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOauth2CredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_oauth2_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOauth2CredentialProvider(ctx, input)
			},
		},
		"create-online-evaluation-config": {
			Name:   "create-online-evaluation-config",
			Fields: fields_create_online_evaluation_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOnlineEvaluationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_online_evaluation_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOnlineEvaluationConfig(ctx, input)
			},
		},
		"create-policy": {
			Name:   "create-policy",
			Fields: fields_create_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicy(ctx, input)
			},
		},
		"create-policy-engine": {
			Name:   "create-policy-engine",
			Fields: fields_create_policy_engine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyEngineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy_engine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicyEngine(ctx, input)
			},
		},
		"create-workload-identity": {
			Name:   "create-workload-identity",
			Fields: fields_create_workload_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkloadIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workload_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkloadIdentity(ctx, input)
			},
		},
		"delete-agent-runtime": {
			Name:   "delete-agent-runtime",
			Fields: fields_delete_agent_runtime,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentRuntimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent_runtime, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgentRuntime(ctx, input)
			},
		},
		"delete-agent-runtime-endpoint": {
			Name:   "delete-agent-runtime-endpoint",
			Fields: fields_delete_agent_runtime_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentRuntimeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent_runtime_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgentRuntimeEndpoint(ctx, input)
			},
		},
		"delete-api-key-credential-provider": {
			Name:   "delete-api-key-credential-provider",
			Fields: fields_delete_api_key_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApiKeyCredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api_key_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApiKeyCredentialProvider(ctx, input)
			},
		},
		"delete-browser": {
			Name:   "delete-browser",
			Fields: fields_delete_browser,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBrowserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_browser, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBrowser(ctx, input)
			},
		},
		"delete-browser-profile": {
			Name:   "delete-browser-profile",
			Fields: fields_delete_browser_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBrowserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_browser_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBrowserProfile(ctx, input)
			},
		},
		"delete-code-interpreter": {
			Name:   "delete-code-interpreter",
			Fields: fields_delete_code_interpreter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCodeInterpreterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_code_interpreter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCodeInterpreter(ctx, input)
			},
		},
		"delete-evaluator": {
			Name:   "delete-evaluator",
			Fields: fields_delete_evaluator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEvaluatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_evaluator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEvaluator(ctx, input)
			},
		},
		"delete-gateway": {
			Name:   "delete-gateway",
			Fields: fields_delete_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGateway(ctx, input)
			},
		},
		"delete-gateway-target": {
			Name:   "delete-gateway-target",
			Fields: fields_delete_gateway_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGatewayTarget(ctx, input)
			},
		},
		"delete-memory": {
			Name:   "delete-memory",
			Fields: fields_delete_memory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMemoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_memory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMemory(ctx, input)
			},
		},
		"delete-oauth2-credential-provider": {
			Name:   "delete-oauth2-credential-provider",
			Fields: fields_delete_oauth2_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOauth2CredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_oauth2_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOauth2CredentialProvider(ctx, input)
			},
		},
		"delete-online-evaluation-config": {
			Name:   "delete-online-evaluation-config",
			Fields: fields_delete_online_evaluation_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOnlineEvaluationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_online_evaluation_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOnlineEvaluationConfig(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-policy-engine": {
			Name:   "delete-policy-engine",
			Fields: fields_delete_policy_engine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyEngineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy_engine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicyEngine(ctx, input)
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
		"delete-workload-identity": {
			Name:   "delete-workload-identity",
			Fields: fields_delete_workload_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkloadIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workload_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkloadIdentity(ctx, input)
			},
		},
		"get-agent-runtime": {
			Name:   "get-agent-runtime",
			Fields: fields_get_agent_runtime,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentRuntimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_runtime, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentRuntime(ctx, input)
			},
		},
		"get-agent-runtime-endpoint": {
			Name:   "get-agent-runtime-endpoint",
			Fields: fields_get_agent_runtime_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentRuntimeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_runtime_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentRuntimeEndpoint(ctx, input)
			},
		},
		"get-api-key-credential-provider": {
			Name:   "get-api-key-credential-provider",
			Fields: fields_get_api_key_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiKeyCredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api_key_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApiKeyCredentialProvider(ctx, input)
			},
		},
		"get-browser": {
			Name:   "get-browser",
			Fields: fields_get_browser,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBrowserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_browser, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBrowser(ctx, input)
			},
		},
		"get-browser-profile": {
			Name:   "get-browser-profile",
			Fields: fields_get_browser_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBrowserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_browser_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBrowserProfile(ctx, input)
			},
		},
		"get-code-interpreter": {
			Name:   "get-code-interpreter",
			Fields: fields_get_code_interpreter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeInterpreterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_interpreter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeInterpreter(ctx, input)
			},
		},
		"get-evaluator": {
			Name:   "get-evaluator",
			Fields: fields_get_evaluator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvaluatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_evaluator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvaluator(ctx, input)
			},
		},
		"get-gateway": {
			Name:   "get-gateway",
			Fields: fields_get_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGateway(ctx, input)
			},
		},
		"get-gateway-target": {
			Name:   "get-gateway-target",
			Fields: fields_get_gateway_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGatewayTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gateway_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGatewayTarget(ctx, input)
			},
		},
		"get-memory": {
			Name:   "get-memory",
			Fields: fields_get_memory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMemoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_memory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMemory(ctx, input)
			},
		},
		"get-oauth2-credential-provider": {
			Name:   "get-oauth2-credential-provider",
			Fields: fields_get_oauth2_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOauth2CredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_oauth2_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOauth2CredentialProvider(ctx, input)
			},
		},
		"get-online-evaluation-config": {
			Name:   "get-online-evaluation-config",
			Fields: fields_get_online_evaluation_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOnlineEvaluationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_online_evaluation_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOnlineEvaluationConfig(ctx, input)
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
		"get-policy-engine": {
			Name:   "get-policy-engine",
			Fields: fields_get_policy_engine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyEngineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_engine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyEngine(ctx, input)
			},
		},
		"get-policy-generation": {
			Name:   "get-policy-generation",
			Fields: fields_get_policy_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyGeneration(ctx, input)
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
		"get-token-vault": {
			Name:   "get-token-vault",
			Fields: fields_get_token_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTokenVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_token_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTokenVault(ctx, input)
			},
		},
		"get-workload-identity": {
			Name:   "get-workload-identity",
			Fields: fields_get_workload_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkloadIdentity(ctx, input)
			},
		},
		"list-agent-runtime-endpoints": {
			Name:   "list-agent-runtime-endpoints",
			Fields: fields_list_agent_runtime_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentRuntimeEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_runtime_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentRuntimeEndpoints(ctx, input)
				}
				var results []*svc.ListAgentRuntimeEndpointsOutput
				p := svc.NewListAgentRuntimeEndpointsPaginator(client, input)
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
		"list-agent-runtime-versions": {
			Name:   "list-agent-runtime-versions",
			Fields: fields_list_agent_runtime_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentRuntimeVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_runtime_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentRuntimeVersions(ctx, input)
				}
				var results []*svc.ListAgentRuntimeVersionsOutput
				p := svc.NewListAgentRuntimeVersionsPaginator(client, input)
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
		"list-agent-runtimes": {
			Name:   "list-agent-runtimes",
			Fields: fields_list_agent_runtimes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentRuntimesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_runtimes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentRuntimes(ctx, input)
				}
				var results []*svc.ListAgentRuntimesOutput
				p := svc.NewListAgentRuntimesPaginator(client, input)
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
		"list-api-key-credential-providers": {
			Name:   "list-api-key-credential-providers",
			Fields: fields_list_api_key_credential_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApiKeyCredentialProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_api_key_credential_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApiKeyCredentialProviders(ctx, input)
				}
				var results []*svc.ListApiKeyCredentialProvidersOutput
				p := svc.NewListApiKeyCredentialProvidersPaginator(client, input)
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
		"list-browser-profiles": {
			Name:   "list-browser-profiles",
			Fields: fields_list_browser_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBrowserProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_browser_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBrowserProfiles(ctx, input)
				}
				var results []*svc.ListBrowserProfilesOutput
				p := svc.NewListBrowserProfilesPaginator(client, input)
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
		"list-browsers": {
			Name:   "list-browsers",
			Fields: fields_list_browsers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBrowsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_browsers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBrowsers(ctx, input)
				}
				var results []*svc.ListBrowsersOutput
				p := svc.NewListBrowsersPaginator(client, input)
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
		"list-code-interpreters": {
			Name:   "list-code-interpreters",
			Fields: fields_list_code_interpreters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeInterpretersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_code_interpreters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCodeInterpreters(ctx, input)
				}
				var results []*svc.ListCodeInterpretersOutput
				p := svc.NewListCodeInterpretersPaginator(client, input)
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
		"list-evaluators": {
			Name:   "list-evaluators",
			Fields: fields_list_evaluators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEvaluatorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_evaluators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEvaluators(ctx, input)
				}
				var results []*svc.ListEvaluatorsOutput
				p := svc.NewListEvaluatorsPaginator(client, input)
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
		"list-gateway-targets": {
			Name:   "list-gateway-targets",
			Fields: fields_list_gateway_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewayTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateway_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGatewayTargets(ctx, input)
				}
				var results []*svc.ListGatewayTargetsOutput
				p := svc.NewListGatewayTargetsPaginator(client, input)
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
		"list-gateways": {
			Name:   "list-gateways",
			Fields: fields_list_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGateways(ctx, input)
				}
				var results []*svc.ListGatewaysOutput
				p := svc.NewListGatewaysPaginator(client, input)
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
		"list-memories": {
			Name:   "list-memories",
			Fields: fields_list_memories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMemoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_memories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMemories(ctx, input)
				}
				var results []*svc.ListMemoriesOutput
				p := svc.NewListMemoriesPaginator(client, input)
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
		"list-oauth2-credential-providers": {
			Name:   "list-oauth2-credential-providers",
			Fields: fields_list_oauth2_credential_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOauth2CredentialProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_oauth2_credential_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOauth2CredentialProviders(ctx, input)
				}
				var results []*svc.ListOauth2CredentialProvidersOutput
				p := svc.NewListOauth2CredentialProvidersPaginator(client, input)
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
		"list-online-evaluation-configs": {
			Name:   "list-online-evaluation-configs",
			Fields: fields_list_online_evaluation_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOnlineEvaluationConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_online_evaluation_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOnlineEvaluationConfigs(ctx, input)
				}
				var results []*svc.ListOnlineEvaluationConfigsOutput
				p := svc.NewListOnlineEvaluationConfigsPaginator(client, input)
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
		"list-policies": {
			Name:   "list-policies",
			Fields: fields_list_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicies(ctx, input)
				}
				var results []*svc.ListPoliciesOutput
				p := svc.NewListPoliciesPaginator(client, input)
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
		"list-policy-engines": {
			Name:   "list-policy-engines",
			Fields: fields_list_policy_engines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyEnginesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_engines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyEngines(ctx, input)
				}
				var results []*svc.ListPolicyEnginesOutput
				p := svc.NewListPolicyEnginesPaginator(client, input)
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
		"list-policy-generation-assets": {
			Name:   "list-policy-generation-assets",
			Fields: fields_list_policy_generation_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyGenerationAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_generation_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyGenerationAssets(ctx, input)
				}
				var results []*svc.ListPolicyGenerationAssetsOutput
				p := svc.NewListPolicyGenerationAssetsPaginator(client, input)
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
		"list-policy-generations": {
			Name:   "list-policy-generations",
			Fields: fields_list_policy_generations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyGenerationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_generations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyGenerations(ctx, input)
				}
				var results []*svc.ListPolicyGenerationsOutput
				p := svc.NewListPolicyGenerationsPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-workload-identities": {
			Name:   "list-workload-identities",
			Fields: fields_list_workload_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workload_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloadIdentities(ctx, input)
				}
				var results []*svc.ListWorkloadIdentitiesOutput
				p := svc.NewListWorkloadIdentitiesPaginator(client, input)
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
		"set-token-vault-cmk": {
			Name:   "set-token-vault-cmk",
			Fields: fields_set_token_vault_cmk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTokenVaultCMKInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_token_vault_cmk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTokenVaultCMK(ctx, input)
			},
		},
		"start-policy-generation": {
			Name:   "start-policy-generation",
			Fields: fields_start_policy_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPolicyGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_policy_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPolicyGeneration(ctx, input)
			},
		},
		"synchronize-gateway-targets": {
			Name:   "synchronize-gateway-targets",
			Fields: fields_synchronize_gateway_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SynchronizeGatewayTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_synchronize_gateway_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SynchronizeGatewayTargets(ctx, input)
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
		"update-agent-runtime": {
			Name:   "update-agent-runtime",
			Fields: fields_update_agent_runtime,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentRuntimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_runtime, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentRuntime(ctx, input)
			},
		},
		"update-agent-runtime-endpoint": {
			Name:   "update-agent-runtime-endpoint",
			Fields: fields_update_agent_runtime_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentRuntimeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_runtime_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentRuntimeEndpoint(ctx, input)
			},
		},
		"update-api-key-credential-provider": {
			Name:   "update-api-key-credential-provider",
			Fields: fields_update_api_key_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApiKeyCredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_api_key_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApiKeyCredentialProvider(ctx, input)
			},
		},
		"update-evaluator": {
			Name:   "update-evaluator",
			Fields: fields_update_evaluator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEvaluatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_evaluator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEvaluator(ctx, input)
			},
		},
		"update-gateway": {
			Name:   "update-gateway",
			Fields: fields_update_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGateway(ctx, input)
			},
		},
		"update-gateway-target": {
			Name:   "update-gateway-target",
			Fields: fields_update_gateway_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayTarget(ctx, input)
			},
		},
		"update-memory": {
			Name:   "update-memory",
			Fields: fields_update_memory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMemoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_memory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMemory(ctx, input)
			},
		},
		"update-oauth2-credential-provider": {
			Name:   "update-oauth2-credential-provider",
			Fields: fields_update_oauth2_credential_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOauth2CredentialProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_oauth2_credential_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOauth2CredentialProvider(ctx, input)
			},
		},
		"update-online-evaluation-config": {
			Name:   "update-online-evaluation-config",
			Fields: fields_update_online_evaluation_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOnlineEvaluationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_online_evaluation_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOnlineEvaluationConfig(ctx, input)
			},
		},
		"update-policy": {
			Name:   "update-policy",
			Fields: fields_update_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePolicy(ctx, input)
			},
		},
		"update-policy-engine": {
			Name:   "update-policy-engine",
			Fields: fields_update_policy_engine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePolicyEngineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_policy_engine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePolicyEngine(ctx, input)
			},
		},
		"update-workload-identity": {
			Name:   "update-workload-identity",
			Fields: fields_update_workload_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkloadIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workload_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkloadIdentity(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bedrockagentcorecontrol", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
