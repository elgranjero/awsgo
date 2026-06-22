package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/apigateway"
)

var fields_create_api_key = []leanruntime.Field{
	{Name: "CustomerId", Flag: "customer-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "bool", Required: false},
	{Name: "GenerateDistinctId", Flag: "generate-distinct-id", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StageKeys", Flag: "stage-keys", Type: "[]types.StageKey", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Value", Flag: "value", Type: "*string", Required: false},
}

var fields_create_authorizer = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "*string", Required: false},
	{Name: "AuthorizerCredentials", Flag: "authorizer-credentials", Type: "*string", Required: false},
	{Name: "AuthorizerResultTtlInSeconds", Flag: "authorizer-result-ttl-in-seconds", Type: "*int32", Required: false},
	{Name: "AuthorizerUri", Flag: "authorizer-uri", Type: "*string", Required: false},
	{Name: "IdentitySource", Flag: "identity-source", Type: "*string", Required: false},
	{Name: "IdentityValidationExpression", Flag: "identity-validation-expression", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProviderARNs", Flag: "provider-arns", Type: "[]string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AuthorizerType", Required: true},
}

var fields_create_base_path_mapping = []leanruntime.Field{
	{Name: "BasePath", Flag: "base-path", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "*string", Required: false},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "CacheClusterEnabled", Flag: "cache-cluster-enabled", Type: "*bool", Required: false},
	{Name: "CacheClusterSize", Flag: "cache-cluster-size", Type: "types.CacheClusterSize", Required: false},
	{Name: "CanarySettings", Flag: "canary-settings", Type: "*types.DeploymentCanarySettings", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageDescription", Flag: "stage-description", Type: "*string", Required: false},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: false},
	{Name: "TracingEnabled", Flag: "tracing-enabled", Type: "*bool", Required: false},
	{Name: "Variables", Flag: "variables", Type: "map[string]string", Required: false},
}

var fields_create_documentation_part = []leanruntime.Field{
	{Name: "Location", Flag: "location", Type: "*types.DocumentationPartLocation", Required: true},
	{Name: "Properties", Flag: "properties", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_create_documentation_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DocumentationVersion", Flag: "documentation-version", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: false},
}

var fields_create_domain_name = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "CertificateBody", Flag: "certificate-body", Type: "*string", Required: false},
	{Name: "CertificateChain", Flag: "certificate-chain", Type: "*string", Required: false},
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: false},
	{Name: "CertificatePrivateKey", Flag: "certificate-private-key", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EndpointAccessMode", Flag: "endpoint-access-mode", Type: "types.EndpointAccessMode", Required: false},
	{Name: "EndpointConfiguration", Flag: "endpoint-configuration", Type: "*types.EndpointConfiguration", Required: false},
	{Name: "MutualTlsAuthentication", Flag: "mutual-tls-authentication", Type: "*types.MutualTlsAuthenticationInput", Required: false},
	{Name: "OwnershipVerificationCertificateArn", Flag: "ownership-verification-certificate-arn", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "RegionalCertificateArn", Flag: "regional-certificate-arn", Type: "*string", Required: false},
	{Name: "RegionalCertificateName", Flag: "regional-certificate-name", Type: "*string", Required: false},
	{Name: "RoutingMode", Flag: "routing-mode", Type: "types.RoutingMode", Required: false},
	{Name: "SecurityPolicy", Flag: "security-policy", Type: "types.SecurityPolicy", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_domain_name_access_association = []leanruntime.Field{
	{Name: "AccessAssociationSource", Flag: "access-association-source", Type: "*string", Required: true},
	{Name: "AccessAssociationSourceType", Flag: "access-association-source-type", Type: "types.AccessAssociationSourceType", Required: true},
	{Name: "DomainNameArn", Flag: "domain-name-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_model = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
}

var fields_create_request_validator = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "ValidateRequestBody", Flag: "validate-request-body", Type: "bool", Required: false},
	{Name: "ValidateRequestParameters", Flag: "validate-request-parameters", Type: "bool", Required: false},
}

var fields_create_resource = []leanruntime.Field{
	{Name: "ParentId", Flag: "parent-id", Type: "*string", Required: true},
	{Name: "PathPart", Flag: "path-part", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_create_rest_api = []leanruntime.Field{
	{Name: "ApiKeySource", Flag: "api-key-source", Type: "types.ApiKeySourceType", Required: false},
	{Name: "BinaryMediaTypes", Flag: "binary-media-types", Type: "[]string", Required: false},
	{Name: "CloneFrom", Flag: "clone-from", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableExecuteApiEndpoint", Flag: "disable-execute-api-endpoint", Type: "bool", Required: false},
	{Name: "EndpointAccessMode", Flag: "endpoint-access-mode", Type: "types.EndpointAccessMode", Required: false},
	{Name: "EndpointConfiguration", Flag: "endpoint-configuration", Type: "*types.EndpointConfiguration", Required: false},
	{Name: "MinimumCompressionSize", Flag: "minimum-compression-size", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "SecurityPolicy", Flag: "security-policy", Type: "types.SecurityPolicy", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_create_stage = []leanruntime.Field{
	{Name: "CacheClusterEnabled", Flag: "cache-cluster-enabled", Type: "bool", Required: false},
	{Name: "CacheClusterSize", Flag: "cache-cluster-size", Type: "types.CacheClusterSize", Required: false},
	{Name: "CanarySettings", Flag: "canary-settings", Type: "*types.CanarySettings", Required: false},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DocumentationVersion", Flag: "documentation-version", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TracingEnabled", Flag: "tracing-enabled", Type: "bool", Required: false},
	{Name: "Variables", Flag: "variables", Type: "map[string]string", Required: false},
}

var fields_create_usage_plan = []leanruntime.Field{
	{Name: "ApiStages", Flag: "api-stages", Type: "[]types.ApiStage", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Quota", Flag: "quota", Type: "*types.QuotaSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Throttle", Flag: "throttle", Type: "*types.ThrottleSettings", Required: false},
}

var fields_create_usage_plan_key = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeyType", Flag: "key-type", Type: "*string", Required: true},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_create_vpc_link = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetArns", Flag: "target-arns", Type: "[]string", Required: true},
}

var fields_delete_api_key = []leanruntime.Field{
	{Name: "ApiKey", Flag: "api-key", Type: "*string", Required: true},
}

var fields_delete_authorizer = []leanruntime.Field{
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_base_path_mapping = []leanruntime.Field{
	{Name: "BasePath", Flag: "base-path", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
}

var fields_delete_client_certificate = []leanruntime.Field{
	{Name: "ClientCertificateId", Flag: "client-certificate-id", Type: "*string", Required: true},
}

var fields_delete_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_documentation_part = []leanruntime.Field{
	{Name: "DocumentationPartId", Flag: "documentation-part-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_documentation_version = []leanruntime.Field{
	{Name: "DocumentationVersion", Flag: "documentation-version", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
}

var fields_delete_domain_name_access_association = []leanruntime.Field{
	{Name: "DomainNameAccessAssociationArn", Flag: "domain-name-access-association-arn", Type: "*string", Required: true},
}

var fields_delete_gateway_response = []leanruntime.Field{
	{Name: "ResponseType", Flag: "response-type", Type: "types.GatewayResponseType", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_integration = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_integration_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_delete_method = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_method_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_delete_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_request_validator = []leanruntime.Field{
	{Name: "RequestValidatorId", Flag: "request-validator-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_rest_api = []leanruntime.Field{
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_delete_stage = []leanruntime.Field{
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_delete_usage_plan = []leanruntime.Field{
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_delete_usage_plan_key = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_delete_vpc_link = []leanruntime.Field{
	{Name: "VpcLinkId", Flag: "vpc-link-id", Type: "*string", Required: true},
}

var fields_flush_stage_authorizers_cache = []leanruntime.Field{
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_flush_stage_cache = []leanruntime.Field{
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_generate_client_certificate = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_get_account = []leanruntime.Field{}

var fields_get_api_key = []leanruntime.Field{
	{Name: "ApiKey", Flag: "api-key", Type: "*string", Required: true},
	{Name: "IncludeValue", Flag: "include-value", Type: "*bool", Required: false},
}

var fields_get_api_keys = []leanruntime.Field{
	{Name: "CustomerId", Flag: "customer-id", Type: "*string", Required: false},
	{Name: "IncludeValues", Flag: "include-values", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NameQuery", Flag: "name-query", Type: "*string", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_get_authorizer = []leanruntime.Field{
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_authorizers = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_base_path_mapping = []leanruntime.Field{
	{Name: "BasePath", Flag: "base-path", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
}

var fields_get_base_path_mappings = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_get_client_certificate = []leanruntime.Field{
	{Name: "ClientCertificateId", Flag: "client-certificate-id", Type: "*string", Required: true},
}

var fields_get_client_certificates = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "Embed", Flag: "embed", Type: "[]string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_deployments = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_documentation_part = []leanruntime.Field{
	{Name: "DocumentationPartId", Flag: "documentation-part-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_documentation_parts = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LocationStatus", Flag: "location-status", Type: "types.LocationStatusType", Required: false},
	{Name: "NameQuery", Flag: "name-query", Type: "*string", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.DocumentationPartType", Required: false},
}

var fields_get_documentation_version = []leanruntime.Field{
	{Name: "DocumentationVersion", Flag: "documentation-version", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_documentation_versions = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
}

var fields_get_domain_name_access_associations = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: false},
}

var fields_get_domain_names = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: false},
}

var fields_get_export = []leanruntime.Field{
	{Name: "Accepts", Flag: "accepts", Type: "*string", Required: false},
	{Name: "ExportType", Flag: "export-type", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_get_gateway_response = []leanruntime.Field{
	{Name: "ResponseType", Flag: "response-type", Type: "types.GatewayResponseType", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_gateway_responses = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_integration = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_integration_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_get_method = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_method_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_get_model = []leanruntime.Field{
	{Name: "Flatten", Flag: "flatten", Type: "bool", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_model_template = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_models = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_request_validator = []leanruntime.Field{
	{Name: "RequestValidatorId", Flag: "request-validator-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_request_validators = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_resource = []leanruntime.Field{
	{Name: "Embed", Flag: "embed", Type: "[]string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_resources = []leanruntime.Field{
	{Name: "Embed", Flag: "embed", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_rest_api = []leanruntime.Field{
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_rest_apis = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_get_sdk = []leanruntime.Field{
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "SdkType", Flag: "sdk-type", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_get_sdk_type = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_sdk_types = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_get_stage = []leanruntime.Field{
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_get_stages = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_get_tags = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_usage = []leanruntime.Field{
	{Name: "EndDate", Flag: "end-date", Type: "*string", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*string", Required: true},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_get_usage_plan = []leanruntime.Field{
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_get_usage_plan_key = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_get_usage_plan_keys = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NameQuery", Flag: "name-query", Type: "*string", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_get_usage_plans = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_get_vpc_link = []leanruntime.Field{
	{Name: "VpcLinkId", Flag: "vpc-link-id", Type: "*string", Required: true},
}

var fields_get_vpc_links = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "*string", Required: false},
}

var fields_import_api_keys = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "[]byte", Required: true},
	{Name: "FailOnWarnings", Flag: "fail-on-warnings", Type: "bool", Required: false},
	{Name: "Format", Flag: "format", Type: "types.ApiKeysFormat", Required: true},
}

var fields_import_documentation_parts = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "[]byte", Required: true},
	{Name: "FailOnWarnings", Flag: "fail-on-warnings", Type: "bool", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.PutMode", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_import_rest_api = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "[]byte", Required: true},
	{Name: "FailOnWarnings", Flag: "fail-on-warnings", Type: "bool", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
}

var fields_put_gateway_response = []leanruntime.Field{
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]string", Required: false},
	{Name: "ResponseTemplates", Flag: "response-templates", Type: "map[string]string", Required: false},
	{Name: "ResponseType", Flag: "response-type", Type: "types.GatewayResponseType", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: false},
}

var fields_put_integration = []leanruntime.Field{
	{Name: "CacheKeyParameters", Flag: "cache-key-parameters", Type: "[]string", Required: false},
	{Name: "CacheNamespace", Flag: "cache-namespace", Type: "*string", Required: false},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: false},
	{Name: "ContentHandling", Flag: "content-handling", Type: "types.ContentHandlingStrategy", Required: false},
	{Name: "Credentials", Flag: "credentials", Type: "*string", Required: false},
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "IntegrationHttpMethod", Flag: "integration-http-method", Type: "*string", Required: false},
	{Name: "IntegrationTarget", Flag: "integration-target", Type: "*string", Required: false},
	{Name: "PassthroughBehavior", Flag: "passthrough-behavior", Type: "*string", Required: false},
	{Name: "RequestParameters", Flag: "request-parameters", Type: "map[string]string", Required: false},
	{Name: "RequestTemplates", Flag: "request-templates", Type: "map[string]string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResponseTransferMode", Flag: "response-transfer-mode", Type: "types.ResponseTransferMode", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "TimeoutInMillis", Flag: "timeout-in-millis", Type: "*int32", Required: false},
	{Name: "TlsConfig", Flag: "tls-config", Type: "*types.TlsConfig", Required: false},
	{Name: "Type", Flag: "type", Type: "types.IntegrationType", Required: true},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: false},
}

var fields_put_integration_response = []leanruntime.Field{
	{Name: "ContentHandling", Flag: "content-handling", Type: "types.ContentHandlingStrategy", Required: false},
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]string", Required: false},
	{Name: "ResponseTemplates", Flag: "response-templates", Type: "map[string]string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "SelectionPattern", Flag: "selection-pattern", Type: "*string", Required: false},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_put_method = []leanruntime.Field{
	{Name: "ApiKeyRequired", Flag: "api-key-required", Type: "bool", Required: false},
	{Name: "AuthorizationScopes", Flag: "authorization-scopes", Type: "[]string", Required: false},
	{Name: "AuthorizationType", Flag: "authorization-type", Type: "*string", Required: true},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: false},
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "OperationName", Flag: "operation-name", Type: "*string", Required: false},
	{Name: "RequestModels", Flag: "request-models", Type: "map[string]string", Required: false},
	{Name: "RequestParameters", Flag: "request-parameters", Type: "map[string]bool", Required: false},
	{Name: "RequestValidatorId", Flag: "request-validator-id", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_put_method_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResponseModels", Flag: "response-models", Type: "map[string]string", Required: false},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]bool", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_put_rest_api = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "[]byte", Required: true},
	{Name: "FailOnWarnings", Flag: "fail-on-warnings", Type: "bool", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.PutMode", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_reject_domain_name_access_association = []leanruntime.Field{
	{Name: "DomainNameAccessAssociationArn", Flag: "domain-name-access-association-arn", Type: "*string", Required: true},
	{Name: "DomainNameArn", Flag: "domain-name-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_test_invoke_authorizer = []leanruntime.Field{
	{Name: "AdditionalContext", Flag: "additional-context", Type: "map[string]string", Required: false},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
	{Name: "Body", Flag: "body", Type: "*string", Required: false},
	{Name: "Headers", Flag: "headers", Type: "map[string]string", Required: false},
	{Name: "MultiValueHeaders", Flag: "multi-value-headers", Type: "map[string][]string", Required: false},
	{Name: "PathWithQueryString", Flag: "path-with-query-string", Type: "*string", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageVariables", Flag: "stage-variables", Type: "map[string]string", Required: false},
}

var fields_test_invoke_method = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "*string", Required: false},
	{Name: "ClientCertificateId", Flag: "client-certificate-id", Type: "*string", Required: false},
	{Name: "Headers", Flag: "headers", Type: "map[string]string", Required: false},
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "MultiValueHeaders", Flag: "multi-value-headers", Type: "map[string][]string", Required: false},
	{Name: "PathWithQueryString", Flag: "path-with-query-string", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageVariables", Flag: "stage-variables", Type: "map[string]string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
}

var fields_update_api_key = []leanruntime.Field{
	{Name: "ApiKey", Flag: "api-key", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
}

var fields_update_authorizer = []leanruntime.Field{
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_base_path_mapping = []leanruntime.Field{
	{Name: "BasePath", Flag: "base-path", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
}

var fields_update_client_certificate = []leanruntime.Field{
	{Name: "ClientCertificateId", Flag: "client-certificate-id", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
}

var fields_update_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_documentation_part = []leanruntime.Field{
	{Name: "DocumentationPartId", Flag: "documentation-part-id", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_documentation_version = []leanruntime.Field{
	{Name: "DocumentationVersion", Flag: "documentation-version", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
}

var fields_update_gateway_response = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "ResponseType", Flag: "response-type", Type: "types.GatewayResponseType", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_integration = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_integration_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_update_method = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_method_response = []leanruntime.Field{
	{Name: "HttpMethod", Flag: "http-method", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StatusCode", Flag: "status-code", Type: "*string", Required: true},
}

var fields_update_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_request_validator = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RequestValidatorId", Flag: "request-validator-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_resource = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_rest_api = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
}

var fields_update_stage = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "RestApiId", Flag: "rest-api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_update_usage = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_update_usage_plan = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "UsagePlanId", Flag: "usage-plan-id", Type: "*string", Required: true},
}

var fields_update_vpc_link = []leanruntime.Field{
	{Name: "PatchOperations", Flag: "patch-operations", Type: "[]types.PatchOperation", Required: false},
	{Name: "VpcLinkId", Flag: "vpc-link-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-api-key": {
			Name:   "create-api-key",
			Fields: fields_create_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApiKey(ctx, input)
			},
		},
		"create-authorizer": {
			Name:   "create-authorizer",
			Fields: fields_create_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAuthorizer(ctx, input)
			},
		},
		"create-base-path-mapping": {
			Name:   "create-base-path-mapping",
			Fields: fields_create_base_path_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBasePathMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_base_path_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBasePathMapping(ctx, input)
			},
		},
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
			},
		},
		"create-documentation-part": {
			Name:   "create-documentation-part",
			Fields: fields_create_documentation_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDocumentationPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_documentation_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDocumentationPart(ctx, input)
			},
		},
		"create-documentation-version": {
			Name:   "create-documentation-version",
			Fields: fields_create_documentation_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDocumentationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_documentation_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDocumentationVersion(ctx, input)
			},
		},
		"create-domain-name": {
			Name:   "create-domain-name",
			Fields: fields_create_domain_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainName(ctx, input)
			},
		},
		"create-domain-name-access-association": {
			Name:   "create-domain-name-access-association",
			Fields: fields_create_domain_name_access_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainNameAccessAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_name_access_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainNameAccessAssociation(ctx, input)
			},
		},
		"create-model": {
			Name:   "create-model",
			Fields: fields_create_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModel(ctx, input)
			},
		},
		"create-request-validator": {
			Name:   "create-request-validator",
			Fields: fields_create_request_validator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRequestValidatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_request_validator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRequestValidator(ctx, input)
			},
		},
		"create-resource": {
			Name:   "create-resource",
			Fields: fields_create_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResource(ctx, input)
			},
		},
		"create-rest-api": {
			Name:   "create-rest-api",
			Fields: fields_create_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRestApi(ctx, input)
			},
		},
		"create-stage": {
			Name:   "create-stage",
			Fields: fields_create_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStage(ctx, input)
			},
		},
		"create-usage-plan": {
			Name:   "create-usage-plan",
			Fields: fields_create_usage_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUsagePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_usage_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUsagePlan(ctx, input)
			},
		},
		"create-usage-plan-key": {
			Name:   "create-usage-plan-key",
			Fields: fields_create_usage_plan_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUsagePlanKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_usage_plan_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUsagePlanKey(ctx, input)
			},
		},
		"create-vpc-link": {
			Name:   "create-vpc-link",
			Fields: fields_create_vpc_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcLink(ctx, input)
			},
		},
		"delete-api-key": {
			Name:   "delete-api-key",
			Fields: fields_delete_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApiKey(ctx, input)
			},
		},
		"delete-authorizer": {
			Name:   "delete-authorizer",
			Fields: fields_delete_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAuthorizer(ctx, input)
			},
		},
		"delete-base-path-mapping": {
			Name:   "delete-base-path-mapping",
			Fields: fields_delete_base_path_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBasePathMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_base_path_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBasePathMapping(ctx, input)
			},
		},
		"delete-client-certificate": {
			Name:   "delete-client-certificate",
			Fields: fields_delete_client_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClientCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_client_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClientCertificate(ctx, input)
			},
		},
		"delete-deployment": {
			Name:   "delete-deployment",
			Fields: fields_delete_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeployment(ctx, input)
			},
		},
		"delete-documentation-part": {
			Name:   "delete-documentation-part",
			Fields: fields_delete_documentation_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDocumentationPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_documentation_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDocumentationPart(ctx, input)
			},
		},
		"delete-documentation-version": {
			Name:   "delete-documentation-version",
			Fields: fields_delete_documentation_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDocumentationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_documentation_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDocumentationVersion(ctx, input)
			},
		},
		"delete-domain-name": {
			Name:   "delete-domain-name",
			Fields: fields_delete_domain_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainName(ctx, input)
			},
		},
		"delete-domain-name-access-association": {
			Name:   "delete-domain-name-access-association",
			Fields: fields_delete_domain_name_access_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainNameAccessAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_name_access_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainNameAccessAssociation(ctx, input)
			},
		},
		"delete-gateway-response": {
			Name:   "delete-gateway-response",
			Fields: fields_delete_gateway_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGatewayResponse(ctx, input)
			},
		},
		"delete-integration": {
			Name:   "delete-integration",
			Fields: fields_delete_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegration(ctx, input)
			},
		},
		"delete-integration-response": {
			Name:   "delete-integration-response",
			Fields: fields_delete_integration_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegrationResponse(ctx, input)
			},
		},
		"delete-method": {
			Name:   "delete-method",
			Fields: fields_delete_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMethod(ctx, input)
			},
		},
		"delete-method-response": {
			Name:   "delete-method-response",
			Fields: fields_delete_method_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMethodResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_method_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMethodResponse(ctx, input)
			},
		},
		"delete-model": {
			Name:   "delete-model",
			Fields: fields_delete_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModel(ctx, input)
			},
		},
		"delete-request-validator": {
			Name:   "delete-request-validator",
			Fields: fields_delete_request_validator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRequestValidatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_request_validator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRequestValidator(ctx, input)
			},
		},
		"delete-resource": {
			Name:   "delete-resource",
			Fields: fields_delete_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResource(ctx, input)
			},
		},
		"delete-rest-api": {
			Name:   "delete-rest-api",
			Fields: fields_delete_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRestApi(ctx, input)
			},
		},
		"delete-stage": {
			Name:   "delete-stage",
			Fields: fields_delete_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStage(ctx, input)
			},
		},
		"delete-usage-plan": {
			Name:   "delete-usage-plan",
			Fields: fields_delete_usage_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUsagePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_usage_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUsagePlan(ctx, input)
			},
		},
		"delete-usage-plan-key": {
			Name:   "delete-usage-plan-key",
			Fields: fields_delete_usage_plan_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUsagePlanKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_usage_plan_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUsagePlanKey(ctx, input)
			},
		},
		"delete-vpc-link": {
			Name:   "delete-vpc-link",
			Fields: fields_delete_vpc_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcLink(ctx, input)
			},
		},
		"flush-stage-authorizers-cache": {
			Name:   "flush-stage-authorizers-cache",
			Fields: fields_flush_stage_authorizers_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FlushStageAuthorizersCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_flush_stage_authorizers_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FlushStageAuthorizersCache(ctx, input)
			},
		},
		"flush-stage-cache": {
			Name:   "flush-stage-cache",
			Fields: fields_flush_stage_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FlushStageCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_flush_stage_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FlushStageCache(ctx, input)
			},
		},
		"generate-client-certificate": {
			Name:   "generate-client-certificate",
			Fields: fields_generate_client_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateClientCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_client_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateClientCertificate(ctx, input)
			},
		},
		"get-account": {
			Name:   "get-account",
			Fields: fields_get_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccount(ctx, input)
			},
		},
		"get-api-key": {
			Name:   "get-api-key",
			Fields: fields_get_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApiKey(ctx, input)
			},
		},
		"get-api-keys": {
			Name:   "get-api-keys",
			Fields: fields_get_api_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_api_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetApiKeys(ctx, input)
				}
				var results []*svc.GetApiKeysOutput
				p := svc.NewGetApiKeysPaginator(client, input)
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
		"get-authorizer": {
			Name:   "get-authorizer",
			Fields: fields_get_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAuthorizer(ctx, input)
			},
		},
		"get-authorizers": {
			Name:   "get-authorizers",
			Fields: fields_get_authorizers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAuthorizersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_authorizers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAuthorizers(ctx, input)
			},
		},
		"get-base-path-mapping": {
			Name:   "get-base-path-mapping",
			Fields: fields_get_base_path_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBasePathMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_base_path_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBasePathMapping(ctx, input)
			},
		},
		"get-base-path-mappings": {
			Name:   "get-base-path-mappings",
			Fields: fields_get_base_path_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBasePathMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_base_path_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBasePathMappings(ctx, input)
				}
				var results []*svc.GetBasePathMappingsOutput
				p := svc.NewGetBasePathMappingsPaginator(client, input)
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
		"get-client-certificate": {
			Name:   "get-client-certificate",
			Fields: fields_get_client_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClientCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_client_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClientCertificate(ctx, input)
			},
		},
		"get-client-certificates": {
			Name:   "get-client-certificates",
			Fields: fields_get_client_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClientCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_client_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetClientCertificates(ctx, input)
				}
				var results []*svc.GetClientCertificatesOutput
				p := svc.NewGetClientCertificatesPaginator(client, input)
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
		"get-deployment": {
			Name:   "get-deployment",
			Fields: fields_get_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployment(ctx, input)
			},
		},
		"get-deployments": {
			Name:   "get-deployments",
			Fields: fields_get_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDeployments(ctx, input)
				}
				var results []*svc.GetDeploymentsOutput
				p := svc.NewGetDeploymentsPaginator(client, input)
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
		"get-documentation-part": {
			Name:   "get-documentation-part",
			Fields: fields_get_documentation_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentationPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_documentation_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentationPart(ctx, input)
			},
		},
		"get-documentation-parts": {
			Name:   "get-documentation-parts",
			Fields: fields_get_documentation_parts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentationPartsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_documentation_parts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentationParts(ctx, input)
			},
		},
		"get-documentation-version": {
			Name:   "get-documentation-version",
			Fields: fields_get_documentation_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_documentation_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentationVersion(ctx, input)
			},
		},
		"get-documentation-versions": {
			Name:   "get-documentation-versions",
			Fields: fields_get_documentation_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentationVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_documentation_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentationVersions(ctx, input)
			},
		},
		"get-domain-name": {
			Name:   "get-domain-name",
			Fields: fields_get_domain_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainName(ctx, input)
			},
		},
		"get-domain-name-access-associations": {
			Name:   "get-domain-name-access-associations",
			Fields: fields_get_domain_name_access_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainNameAccessAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_name_access_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainNameAccessAssociations(ctx, input)
			},
		},
		"get-domain-names": {
			Name:   "get-domain-names",
			Fields: fields_get_domain_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainNamesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_domain_names, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDomainNames(ctx, input)
				}
				var results []*svc.GetDomainNamesOutput
				p := svc.NewGetDomainNamesPaginator(client, input)
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
		"get-export": {
			Name:   "get-export",
			Fields: fields_get_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExport(ctx, input)
			},
		},
		"get-gateway-response": {
			Name:   "get-gateway-response",
			Fields: fields_get_gateway_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGatewayResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gateway_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGatewayResponse(ctx, input)
			},
		},
		"get-gateway-responses": {
			Name:   "get-gateway-responses",
			Fields: fields_get_gateway_responses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGatewayResponsesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gateway_responses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGatewayResponses(ctx, input)
			},
		},
		"get-integration": {
			Name:   "get-integration",
			Fields: fields_get_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegration(ctx, input)
			},
		},
		"get-integration-response": {
			Name:   "get-integration-response",
			Fields: fields_get_integration_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegrationResponse(ctx, input)
			},
		},
		"get-method": {
			Name:   "get-method",
			Fields: fields_get_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMethod(ctx, input)
			},
		},
		"get-method-response": {
			Name:   "get-method-response",
			Fields: fields_get_method_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMethodResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_method_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMethodResponse(ctx, input)
			},
		},
		"get-model": {
			Name:   "get-model",
			Fields: fields_get_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModel(ctx, input)
			},
		},
		"get-model-template": {
			Name:   "get-model-template",
			Fields: fields_get_model_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelTemplate(ctx, input)
			},
		},
		"get-models": {
			Name:   "get-models",
			Fields: fields_get_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetModels(ctx, input)
				}
				var results []*svc.GetModelsOutput
				p := svc.NewGetModelsPaginator(client, input)
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
		"get-request-validator": {
			Name:   "get-request-validator",
			Fields: fields_get_request_validator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRequestValidatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_request_validator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRequestValidator(ctx, input)
			},
		},
		"get-request-validators": {
			Name:   "get-request-validators",
			Fields: fields_get_request_validators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRequestValidatorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_request_validators, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRequestValidators(ctx, input)
			},
		},
		"get-resource": {
			Name:   "get-resource",
			Fields: fields_get_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResource(ctx, input)
			},
		},
		"get-resources": {
			Name:   "get-resources",
			Fields: fields_get_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResources(ctx, input)
				}
				var results []*svc.GetResourcesOutput
				p := svc.NewGetResourcesPaginator(client, input)
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
		"get-rest-api": {
			Name:   "get-rest-api",
			Fields: fields_get_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRestApi(ctx, input)
			},
		},
		"get-rest-apis": {
			Name:   "get-rest-apis",
			Fields: fields_get_rest_apis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRestApisInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_rest_apis, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetRestApis(ctx, input)
				}
				var results []*svc.GetRestApisOutput
				p := svc.NewGetRestApisPaginator(client, input)
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
		"get-sdk": {
			Name:   "get-sdk",
			Fields: fields_get_sdk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSdkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sdk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSdk(ctx, input)
			},
		},
		"get-sdk-type": {
			Name:   "get-sdk-type",
			Fields: fields_get_sdk_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSdkTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sdk_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSdkType(ctx, input)
			},
		},
		"get-sdk-types": {
			Name:   "get-sdk-types",
			Fields: fields_get_sdk_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSdkTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sdk_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSdkTypes(ctx, input)
			},
		},
		"get-stage": {
			Name:   "get-stage",
			Fields: fields_get_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStage(ctx, input)
			},
		},
		"get-stages": {
			Name:   "get-stages",
			Fields: fields_get_stages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStages(ctx, input)
			},
		},
		"get-tags": {
			Name:   "get-tags",
			Fields: fields_get_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTags(ctx, input)
			},
		},
		"get-usage": {
			Name:   "get-usage",
			Fields: fields_get_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_usage, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUsage(ctx, input)
				}
				var results []*svc.GetUsageOutput
				p := svc.NewGetUsagePaginator(client, input)
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
		"get-usage-plan": {
			Name:   "get-usage-plan",
			Fields: fields_get_usage_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsagePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_usage_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsagePlan(ctx, input)
			},
		},
		"get-usage-plan-key": {
			Name:   "get-usage-plan-key",
			Fields: fields_get_usage_plan_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsagePlanKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_usage_plan_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsagePlanKey(ctx, input)
			},
		},
		"get-usage-plan-keys": {
			Name:   "get-usage-plan-keys",
			Fields: fields_get_usage_plan_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsagePlanKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_usage_plan_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUsagePlanKeys(ctx, input)
				}
				var results []*svc.GetUsagePlanKeysOutput
				p := svc.NewGetUsagePlanKeysPaginator(client, input)
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
		"get-usage-plans": {
			Name:   "get-usage-plans",
			Fields: fields_get_usage_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsagePlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_usage_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUsagePlans(ctx, input)
				}
				var results []*svc.GetUsagePlansOutput
				p := svc.NewGetUsagePlansPaginator(client, input)
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
		"get-vpc-link": {
			Name:   "get-vpc-link",
			Fields: fields_get_vpc_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpcLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpc_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpcLink(ctx, input)
			},
		},
		"get-vpc-links": {
			Name:   "get-vpc-links",
			Fields: fields_get_vpc_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpcLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_vpc_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetVpcLinks(ctx, input)
				}
				var results []*svc.GetVpcLinksOutput
				p := svc.NewGetVpcLinksPaginator(client, input)
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
		"import-api-keys": {
			Name:   "import-api-keys",
			Fields: fields_import_api_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportApiKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_api_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportApiKeys(ctx, input)
			},
		},
		"import-documentation-parts": {
			Name:   "import-documentation-parts",
			Fields: fields_import_documentation_parts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportDocumentationPartsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_documentation_parts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportDocumentationParts(ctx, input)
			},
		},
		"import-rest-api": {
			Name:   "import-rest-api",
			Fields: fields_import_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportRestApi(ctx, input)
			},
		},
		"put-gateway-response": {
			Name:   "put-gateway-response",
			Fields: fields_put_gateway_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGatewayResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_gateway_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGatewayResponse(ctx, input)
			},
		},
		"put-integration": {
			Name:   "put-integration",
			Fields: fields_put_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIntegration(ctx, input)
			},
		},
		"put-integration-response": {
			Name:   "put-integration-response",
			Fields: fields_put_integration_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIntegrationResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_integration_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIntegrationResponse(ctx, input)
			},
		},
		"put-method": {
			Name:   "put-method",
			Fields: fields_put_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMethod(ctx, input)
			},
		},
		"put-method-response": {
			Name:   "put-method-response",
			Fields: fields_put_method_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMethodResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_method_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMethodResponse(ctx, input)
			},
		},
		"put-rest-api": {
			Name:   "put-rest-api",
			Fields: fields_put_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRestApi(ctx, input)
			},
		},
		"reject-domain-name-access-association": {
			Name:   "reject-domain-name-access-association",
			Fields: fields_reject_domain_name_access_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectDomainNameAccessAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_domain_name_access_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectDomainNameAccessAssociation(ctx, input)
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
		"test-invoke-authorizer": {
			Name:   "test-invoke-authorizer",
			Fields: fields_test_invoke_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestInvokeAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_invoke_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestInvokeAuthorizer(ctx, input)
			},
		},
		"test-invoke-method": {
			Name:   "test-invoke-method",
			Fields: fields_test_invoke_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestInvokeMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_invoke_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestInvokeMethod(ctx, input)
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
		"update-account": {
			Name:   "update-account",
			Fields: fields_update_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccount(ctx, input)
			},
		},
		"update-api-key": {
			Name:   "update-api-key",
			Fields: fields_update_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApiKey(ctx, input)
			},
		},
		"update-authorizer": {
			Name:   "update-authorizer",
			Fields: fields_update_authorizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAuthorizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_authorizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAuthorizer(ctx, input)
			},
		},
		"update-base-path-mapping": {
			Name:   "update-base-path-mapping",
			Fields: fields_update_base_path_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBasePathMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_base_path_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBasePathMapping(ctx, input)
			},
		},
		"update-client-certificate": {
			Name:   "update-client-certificate",
			Fields: fields_update_client_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClientCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_client_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClientCertificate(ctx, input)
			},
		},
		"update-deployment": {
			Name:   "update-deployment",
			Fields: fields_update_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeployment(ctx, input)
			},
		},
		"update-documentation-part": {
			Name:   "update-documentation-part",
			Fields: fields_update_documentation_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentationPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_documentation_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocumentationPart(ctx, input)
			},
		},
		"update-documentation-version": {
			Name:   "update-documentation-version",
			Fields: fields_update_documentation_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_documentation_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocumentationVersion(ctx, input)
			},
		},
		"update-domain-name": {
			Name:   "update-domain-name",
			Fields: fields_update_domain_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainName(ctx, input)
			},
		},
		"update-gateway-response": {
			Name:   "update-gateway-response",
			Fields: fields_update_gateway_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayResponse(ctx, input)
			},
		},
		"update-integration": {
			Name:   "update-integration",
			Fields: fields_update_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIntegration(ctx, input)
			},
		},
		"update-integration-response": {
			Name:   "update-integration-response",
			Fields: fields_update_integration_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIntegrationResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_integration_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIntegrationResponse(ctx, input)
			},
		},
		"update-method": {
			Name:   "update-method",
			Fields: fields_update_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMethod(ctx, input)
			},
		},
		"update-method-response": {
			Name:   "update-method-response",
			Fields: fields_update_method_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMethodResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_method_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMethodResponse(ctx, input)
			},
		},
		"update-model": {
			Name:   "update-model",
			Fields: fields_update_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModel(ctx, input)
			},
		},
		"update-request-validator": {
			Name:   "update-request-validator",
			Fields: fields_update_request_validator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRequestValidatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_request_validator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRequestValidator(ctx, input)
			},
		},
		"update-resource": {
			Name:   "update-resource",
			Fields: fields_update_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResource(ctx, input)
			},
		},
		"update-rest-api": {
			Name:   "update-rest-api",
			Fields: fields_update_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRestApi(ctx, input)
			},
		},
		"update-stage": {
			Name:   "update-stage",
			Fields: fields_update_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStage(ctx, input)
			},
		},
		"update-usage": {
			Name:   "update-usage",
			Fields: fields_update_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUsage(ctx, input)
			},
		},
		"update-usage-plan": {
			Name:   "update-usage-plan",
			Fields: fields_update_usage_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUsagePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_usage_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUsagePlan(ctx, input)
			},
		},
		"update-vpc-link": {
			Name:   "update-vpc-link",
			Fields: fields_update_vpc_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcLink(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("apigateway", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
