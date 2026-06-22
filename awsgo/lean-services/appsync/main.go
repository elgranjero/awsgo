package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appsync"
)

var fields_associate_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_associate_merged_graphql_api = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
	{Name: "SourceApiAssociationConfig", Flag: "source-api-association-config", Type: "*types.SourceApiAssociationConfig", Required: false},
	{Name: "SourceApiIdentifier", Flag: "source-api-identifier", Type: "*string", Required: true},
}

var fields_associate_source_graphql_api = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
	{Name: "SourceApiAssociationConfig", Flag: "source-api-association-config", Type: "*types.SourceApiAssociationConfig", Required: false},
	{Name: "SourceApiIdentifier", Flag: "source-api-identifier", Type: "*string", Required: true},
}

var fields_create_api = []leanruntime.Field{
	{Name: "EventConfig", Flag: "event-config", Type: "*types.EventConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwnerContact", Flag: "owner-contact", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_api_cache = []leanruntime.Field{
	{Name: "ApiCachingBehavior", Flag: "api-caching-behavior", Type: "types.ApiCachingBehavior", Required: true},
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AtRestEncryptionEnabled", Flag: "at-rest-encryption-enabled", Type: "bool", Required: false},
	{Name: "HealthMetricsConfig", Flag: "health-metrics-config", Type: "types.CacheHealthMetricsConfig", Required: false},
	{Name: "TransitEncryptionEnabled", Flag: "transit-encryption-enabled", Type: "bool", Required: false},
	{Name: "Ttl", Flag: "ttl", Type: "int64", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ApiCacheType", Required: true},
}

var fields_create_api_key = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expires", Flag: "expires", Type: "int64", Required: false},
}

var fields_create_channel_namespace = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "CodeHandlers", Flag: "code-handlers", Type: "*string", Required: false},
	{Name: "HandlerConfigs", Flag: "handler-configs", Type: "*types.HandlerConfigs", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PublishAuthModes", Flag: "publish-auth-modes", Type: "[]types.AuthMode", Required: false},
	{Name: "SubscribeAuthModes", Flag: "subscribe-auth-modes", Type: "[]types.AuthMode", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_source = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DynamodbConfig", Flag: "dynamodb-config", Type: "*types.DynamodbDataSourceConfig", Required: false},
	{Name: "ElasticsearchConfig", Flag: "elasticsearch-config", Type: "*types.ElasticsearchDataSourceConfig", Required: false},
	{Name: "EventBridgeConfig", Flag: "event-bridge-config", Type: "*types.EventBridgeDataSourceConfig", Required: false},
	{Name: "HttpConfig", Flag: "http-config", Type: "*types.HttpDataSourceConfig", Required: false},
	{Name: "LambdaConfig", Flag: "lambda-config", Type: "*types.LambdaDataSourceConfig", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "types.DataSourceLevelMetricsConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OpenSearchServiceConfig", Flag: "open-search-service-config", Type: "*types.OpenSearchServiceDataSourceConfig", Required: false},
	{Name: "RelationalDatabaseConfig", Flag: "relational-database-config", Type: "*types.RelationalDatabaseDataSourceConfig", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataSourceType", Required: true},
}

var fields_create_domain_name = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_function = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionVersion", Flag: "function-version", Type: "*string", Required: false},
	{Name: "MaxBatchSize", Flag: "max-batch-size", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestMappingTemplate", Flag: "request-mapping-template", Type: "*string", Required: false},
	{Name: "ResponseMappingTemplate", Flag: "response-mapping-template", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "*types.AppSyncRuntime", Required: false},
	{Name: "SyncConfig", Flag: "sync-config", Type: "*types.SyncConfig", Required: false},
}

var fields_create_graphql_api = []leanruntime.Field{
	{Name: "AdditionalAuthenticationProviders", Flag: "additional-authentication-providers", Type: "[]types.AdditionalAuthenticationProvider", Required: false},
	{Name: "ApiType", Flag: "api-type", Type: "types.GraphQLApiType", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "EnhancedMetricsConfig", Flag: "enhanced-metrics-config", Type: "*types.EnhancedMetricsConfig", Required: false},
	{Name: "IntrospectionConfig", Flag: "introspection-config", Type: "types.GraphQLApiIntrospectionConfig", Required: false},
	{Name: "LambdaAuthorizerConfig", Flag: "lambda-authorizer-config", Type: "*types.LambdaAuthorizerConfig", Required: false},
	{Name: "LogConfig", Flag: "log-config", Type: "*types.LogConfig", Required: false},
	{Name: "MergedApiExecutionRoleArn", Flag: "merged-api-execution-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OpenIDConnectConfig", Flag: "open-id-connect-config", Type: "*types.OpenIDConnectConfig", Required: false},
	{Name: "OwnerContact", Flag: "owner-contact", Type: "*string", Required: false},
	{Name: "QueryDepthLimit", Flag: "query-depth-limit", Type: "int32", Required: false},
	{Name: "ResolverCountLimit", Flag: "resolver-count-limit", Type: "int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UserPoolConfig", Flag: "user-pool-config", Type: "*types.UserPoolConfig", Required: false},
	{Name: "Visibility", Flag: "visibility", Type: "types.GraphQLApiVisibility", Required: false},
	{Name: "XrayEnabled", Flag: "xray-enabled", Type: "bool", Required: false},
}

var fields_create_resolver = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "CachingConfig", Flag: "caching-config", Type: "*types.CachingConfig", Required: false},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: false},
	{Name: "FieldName", Flag: "field-name", Type: "*string", Required: true},
	{Name: "Kind", Flag: "kind", Type: "types.ResolverKind", Required: false},
	{Name: "MaxBatchSize", Flag: "max-batch-size", Type: "int32", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "types.ResolverLevelMetricsConfig", Required: false},
	{Name: "PipelineConfig", Flag: "pipeline-config", Type: "*types.PipelineConfig", Required: false},
	{Name: "RequestMappingTemplate", Flag: "request-mapping-template", Type: "*string", Required: false},
	{Name: "ResponseMappingTemplate", Flag: "response-mapping-template", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "*types.AppSyncRuntime", Required: false},
	{Name: "SyncConfig", Flag: "sync-config", Type: "*types.SyncConfig", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_create_type = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "types.TypeDefinitionFormat", Required: true},
}

var fields_delete_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_delete_api_cache = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_delete_api_key = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_channel_namespace = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_function = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "FunctionId", Flag: "function-id", Type: "*string", Required: true},
}

var fields_delete_graphql_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_delete_resolver = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "FieldName", Flag: "field-name", Type: "*string", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_delete_type = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_disassociate_api = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_disassociate_merged_graphql_api = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "SourceApiIdentifier", Flag: "source-api-identifier", Type: "*string", Required: true},
}

var fields_disassociate_source_graphql_api = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
}

var fields_evaluate_code = []leanruntime.Field{
	{Name: "Code", Flag: "code", Type: "*string", Required: true},
	{Name: "Context", Flag: "context", Type: "*string", Required: true},
	{Name: "Function", Flag: "function", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "*types.AppSyncRuntime", Required: true},
}

var fields_evaluate_mapping_template = []leanruntime.Field{
	{Name: "Context", Flag: "context", Type: "*string", Required: true},
	{Name: "Template", Flag: "template", Type: "*string", Required: true},
}

var fields_flush_api_cache = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_api_association = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_api_cache = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_channel_namespace = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_data_source = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_data_source_introspection = []leanruntime.Field{
	{Name: "IncludeModelsSDL", Flag: "include-models-sdl", Type: "bool", Required: false},
	{Name: "IntrospectionId", Flag: "introspection-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_function = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "FunctionId", Flag: "function-id", Type: "*string", Required: true},
}

var fields_get_graphql_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_graphql_api_environment_variables = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_introspection_schema = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "types.OutputType", Required: true},
	{Name: "IncludeDirectives", Flag: "include-directives", Type: "*bool", Required: false},
}

var fields_get_resolver = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "FieldName", Flag: "field-name", Type: "*string", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_get_schema_creation_status = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_source_api_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
}

var fields_get_type = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "types.TypeDefinitionFormat", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_list_api_keys = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_apis = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channel_namespaces = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domain_names = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_functions = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_graphql_apis = []leanruntime.Field{
	{Name: "ApiType", Flag: "api-type", Type: "types.GraphQLApiType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Ownership", Required: false},
}

var fields_list_resolvers = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_list_resolvers_by_function = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "FunctionId", Flag: "function-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_source_api_associations = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_types = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "types.TypeDefinitionFormat", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_types_by_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "types.TypeDefinitionFormat", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_graphql_api_environment_variables = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: true},
}

var fields_start_data_source_introspection = []leanruntime.Field{
	{Name: "RdsDataApiConfig", Flag: "rds-data-api-config", Type: "*types.RdsDataApiConfig", Required: false},
}

var fields_start_schema_creation = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "[]byte", Required: true},
}

var fields_start_schema_merge = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "EventConfig", Flag: "event-config", Type: "*types.EventConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwnerContact", Flag: "owner-contact", Type: "*string", Required: false},
}

var fields_update_api_cache = []leanruntime.Field{
	{Name: "ApiCachingBehavior", Flag: "api-caching-behavior", Type: "types.ApiCachingBehavior", Required: true},
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "HealthMetricsConfig", Flag: "health-metrics-config", Type: "types.CacheHealthMetricsConfig", Required: false},
	{Name: "Ttl", Flag: "ttl", Type: "int64", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ApiCacheType", Required: true},
}

var fields_update_api_key = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expires", Flag: "expires", Type: "int64", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_channel_namespace = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "CodeHandlers", Flag: "code-handlers", Type: "*string", Required: false},
	{Name: "HandlerConfigs", Flag: "handler-configs", Type: "*types.HandlerConfigs", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PublishAuthModes", Flag: "publish-auth-modes", Type: "[]types.AuthMode", Required: false},
	{Name: "SubscribeAuthModes", Flag: "subscribe-auth-modes", Type: "[]types.AuthMode", Required: false},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DynamodbConfig", Flag: "dynamodb-config", Type: "*types.DynamodbDataSourceConfig", Required: false},
	{Name: "ElasticsearchConfig", Flag: "elasticsearch-config", Type: "*types.ElasticsearchDataSourceConfig", Required: false},
	{Name: "EventBridgeConfig", Flag: "event-bridge-config", Type: "*types.EventBridgeDataSourceConfig", Required: false},
	{Name: "HttpConfig", Flag: "http-config", Type: "*types.HttpDataSourceConfig", Required: false},
	{Name: "LambdaConfig", Flag: "lambda-config", Type: "*types.LambdaDataSourceConfig", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "types.DataSourceLevelMetricsConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OpenSearchServiceConfig", Flag: "open-search-service-config", Type: "*types.OpenSearchServiceDataSourceConfig", Required: false},
	{Name: "RelationalDatabaseConfig", Flag: "relational-database-config", Type: "*types.RelationalDatabaseDataSourceConfig", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataSourceType", Required: true},
}

var fields_update_domain_name = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_update_function = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionId", Flag: "function-id", Type: "*string", Required: true},
	{Name: "FunctionVersion", Flag: "function-version", Type: "*string", Required: false},
	{Name: "MaxBatchSize", Flag: "max-batch-size", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestMappingTemplate", Flag: "request-mapping-template", Type: "*string", Required: false},
	{Name: "ResponseMappingTemplate", Flag: "response-mapping-template", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "*types.AppSyncRuntime", Required: false},
	{Name: "SyncConfig", Flag: "sync-config", Type: "*types.SyncConfig", Required: false},
}

var fields_update_graphql_api = []leanruntime.Field{
	{Name: "AdditionalAuthenticationProviders", Flag: "additional-authentication-providers", Type: "[]types.AdditionalAuthenticationProvider", Required: false},
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "EnhancedMetricsConfig", Flag: "enhanced-metrics-config", Type: "*types.EnhancedMetricsConfig", Required: false},
	{Name: "IntrospectionConfig", Flag: "introspection-config", Type: "types.GraphQLApiIntrospectionConfig", Required: false},
	{Name: "LambdaAuthorizerConfig", Flag: "lambda-authorizer-config", Type: "*types.LambdaAuthorizerConfig", Required: false},
	{Name: "LogConfig", Flag: "log-config", Type: "*types.LogConfig", Required: false},
	{Name: "MergedApiExecutionRoleArn", Flag: "merged-api-execution-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OpenIDConnectConfig", Flag: "open-id-connect-config", Type: "*types.OpenIDConnectConfig", Required: false},
	{Name: "OwnerContact", Flag: "owner-contact", Type: "*string", Required: false},
	{Name: "QueryDepthLimit", Flag: "query-depth-limit", Type: "int32", Required: false},
	{Name: "ResolverCountLimit", Flag: "resolver-count-limit", Type: "int32", Required: false},
	{Name: "UserPoolConfig", Flag: "user-pool-config", Type: "*types.UserPoolConfig", Required: false},
	{Name: "XrayEnabled", Flag: "xray-enabled", Type: "bool", Required: false},
}

var fields_update_resolver = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "CachingConfig", Flag: "caching-config", Type: "*types.CachingConfig", Required: false},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: false},
	{Name: "FieldName", Flag: "field-name", Type: "*string", Required: true},
	{Name: "Kind", Flag: "kind", Type: "types.ResolverKind", Required: false},
	{Name: "MaxBatchSize", Flag: "max-batch-size", Type: "int32", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "types.ResolverLevelMetricsConfig", Required: false},
	{Name: "PipelineConfig", Flag: "pipeline-config", Type: "*types.PipelineConfig", Required: false},
	{Name: "RequestMappingTemplate", Flag: "request-mapping-template", Type: "*string", Required: false},
	{Name: "ResponseMappingTemplate", Flag: "response-mapping-template", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "*types.AppSyncRuntime", Required: false},
	{Name: "SyncConfig", Flag: "sync-config", Type: "*types.SyncConfig", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_update_source_api_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MergedApiIdentifier", Flag: "merged-api-identifier", Type: "*string", Required: true},
	{Name: "SourceApiAssociationConfig", Flag: "source-api-association-config", Type: "*types.SourceApiAssociationConfig", Required: false},
}

var fields_update_type = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.TypeDefinitionFormat", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-api": {
			Name:   "associate-api",
			Fields: fields_associate_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApi(ctx, input)
			},
		},
		"associate-merged-graphql-api": {
			Name:   "associate-merged-graphql-api",
			Fields: fields_associate_merged_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMergedGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_merged_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMergedGraphqlApi(ctx, input)
			},
		},
		"associate-source-graphql-api": {
			Name:   "associate-source-graphql-api",
			Fields: fields_associate_source_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSourceGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_source_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSourceGraphqlApi(ctx, input)
			},
		},
		"create-api": {
			Name:   "create-api",
			Fields: fields_create_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApi(ctx, input)
			},
		},
		"create-api-cache": {
			Name:   "create-api-cache",
			Fields: fields_create_api_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApiCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApiCache(ctx, input)
			},
		},
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
		"create-channel-namespace": {
			Name:   "create-channel-namespace",
			Fields: fields_create_channel_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelNamespace(ctx, input)
			},
		},
		"create-data-source": {
			Name:   "create-data-source",
			Fields: fields_create_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSource(ctx, input)
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
		"create-graphql-api": {
			Name:   "create-graphql-api",
			Fields: fields_create_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGraphqlApi(ctx, input)
			},
		},
		"create-resolver": {
			Name:   "create-resolver",
			Fields: fields_create_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResolver(ctx, input)
			},
		},
		"create-type": {
			Name:   "create-type",
			Fields: fields_create_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateType(ctx, input)
			},
		},
		"delete-api": {
			Name:   "delete-api",
			Fields: fields_delete_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApi(ctx, input)
			},
		},
		"delete-api-cache": {
			Name:   "delete-api-cache",
			Fields: fields_delete_api_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApiCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApiCache(ctx, input)
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
		"delete-channel-namespace": {
			Name:   "delete-channel-namespace",
			Fields: fields_delete_channel_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelNamespace(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
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
		"delete-graphql-api": {
			Name:   "delete-graphql-api",
			Fields: fields_delete_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGraphqlApi(ctx, input)
			},
		},
		"delete-resolver": {
			Name:   "delete-resolver",
			Fields: fields_delete_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResolver(ctx, input)
			},
		},
		"delete-type": {
			Name:   "delete-type",
			Fields: fields_delete_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteType(ctx, input)
			},
		},
		"disassociate-api": {
			Name:   "disassociate-api",
			Fields: fields_disassociate_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApi(ctx, input)
			},
		},
		"disassociate-merged-graphql-api": {
			Name:   "disassociate-merged-graphql-api",
			Fields: fields_disassociate_merged_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMergedGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_merged_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMergedGraphqlApi(ctx, input)
			},
		},
		"disassociate-source-graphql-api": {
			Name:   "disassociate-source-graphql-api",
			Fields: fields_disassociate_source_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSourceGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_source_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSourceGraphqlApi(ctx, input)
			},
		},
		"evaluate-code": {
			Name:   "evaluate-code",
			Fields: fields_evaluate_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluateCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evaluate_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EvaluateCode(ctx, input)
			},
		},
		"evaluate-mapping-template": {
			Name:   "evaluate-mapping-template",
			Fields: fields_evaluate_mapping_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluateMappingTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evaluate_mapping_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EvaluateMappingTemplate(ctx, input)
			},
		},
		"flush-api-cache": {
			Name:   "flush-api-cache",
			Fields: fields_flush_api_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FlushApiCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_flush_api_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FlushApiCache(ctx, input)
			},
		},
		"get-api": {
			Name:   "get-api",
			Fields: fields_get_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApi(ctx, input)
			},
		},
		"get-api-association": {
			Name:   "get-api-association",
			Fields: fields_get_api_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApiAssociation(ctx, input)
			},
		},
		"get-api-cache": {
			Name:   "get-api-cache",
			Fields: fields_get_api_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApiCache(ctx, input)
			},
		},
		"get-channel-namespace": {
			Name:   "get-channel-namespace",
			Fields: fields_get_channel_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelNamespace(ctx, input)
			},
		},
		"get-data-source": {
			Name:   "get-data-source",
			Fields: fields_get_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSource(ctx, input)
			},
		},
		"get-data-source-introspection": {
			Name:   "get-data-source-introspection",
			Fields: fields_get_data_source_introspection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceIntrospectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source_introspection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSourceIntrospection(ctx, input)
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
		"get-graphql-api": {
			Name:   "get-graphql-api",
			Fields: fields_get_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGraphqlApi(ctx, input)
			},
		},
		"get-graphql-api-environment-variables": {
			Name:   "get-graphql-api-environment-variables",
			Fields: fields_get_graphql_api_environment_variables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGraphqlApiEnvironmentVariablesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_graphql_api_environment_variables, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGraphqlApiEnvironmentVariables(ctx, input)
			},
		},
		"get-introspection-schema": {
			Name:   "get-introspection-schema",
			Fields: fields_get_introspection_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntrospectionSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_introspection_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntrospectionSchema(ctx, input)
			},
		},
		"get-resolver": {
			Name:   "get-resolver",
			Fields: fields_get_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolver(ctx, input)
			},
		},
		"get-schema-creation-status": {
			Name:   "get-schema-creation-status",
			Fields: fields_get_schema_creation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaCreationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_creation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaCreationStatus(ctx, input)
			},
		},
		"get-source-api-association": {
			Name:   "get-source-api-association",
			Fields: fields_get_source_api_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSourceApiAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_source_api_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSourceApiAssociation(ctx, input)
			},
		},
		"get-type": {
			Name:   "get-type",
			Fields: fields_get_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetType(ctx, input)
			},
		},
		"list-api-keys": {
			Name:   "list-api-keys",
			Fields: fields_list_api_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApiKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_api_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApiKeys(ctx, input)
				}
				var results []*svc.ListApiKeysOutput
				p := svc.NewListApiKeysPaginator(client, input)
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
		"list-apis": {
			Name:   "list-apis",
			Fields: fields_list_apis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApisInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_apis, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApis(ctx, input)
				}
				var results []*svc.ListApisOutput
				p := svc.NewListApisPaginator(client, input)
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
		"list-channel-namespaces": {
			Name:   "list-channel-namespaces",
			Fields: fields_list_channel_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelNamespaces(ctx, input)
				}
				var results []*svc.ListChannelNamespacesOutput
				p := svc.NewListChannelNamespacesPaginator(client, input)
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
		"list-data-sources": {
			Name:   "list-data-sources",
			Fields: fields_list_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSources(ctx, input)
				}
				var results []*svc.ListDataSourcesOutput
				p := svc.NewListDataSourcesPaginator(client, input)
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
		"list-domain-names": {
			Name:   "list-domain-names",
			Fields: fields_list_domain_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainNamesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_names, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainNames(ctx, input)
				}
				var results []*svc.ListDomainNamesOutput
				p := svc.NewListDomainNamesPaginator(client, input)
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
		"list-graphql-apis": {
			Name:   "list-graphql-apis",
			Fields: fields_list_graphql_apis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGraphqlApisInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_graphql_apis, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGraphqlApis(ctx, input)
				}
				var results []*svc.ListGraphqlApisOutput
				p := svc.NewListGraphqlApisPaginator(client, input)
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
		"list-resolvers": {
			Name:   "list-resolvers",
			Fields: fields_list_resolvers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolversInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolvers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolvers(ctx, input)
				}
				var results []*svc.ListResolversOutput
				p := svc.NewListResolversPaginator(client, input)
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
		"list-resolvers-by-function": {
			Name:   "list-resolvers-by-function",
			Fields: fields_list_resolvers_by_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolversByFunctionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolvers_by_function, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolversByFunction(ctx, input)
				}
				var results []*svc.ListResolversByFunctionOutput
				p := svc.NewListResolversByFunctionPaginator(client, input)
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
		"list-source-api-associations": {
			Name:   "list-source-api-associations",
			Fields: fields_list_source_api_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceApiAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_api_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceApiAssociations(ctx, input)
				}
				var results []*svc.ListSourceApiAssociationsOutput
				p := svc.NewListSourceApiAssociationsPaginator(client, input)
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
		"list-types": {
			Name:   "list-types",
			Fields: fields_list_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypes(ctx, input)
				}
				var results []*svc.ListTypesOutput
				p := svc.NewListTypesPaginator(client, input)
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
		"list-types-by-association": {
			Name:   "list-types-by-association",
			Fields: fields_list_types_by_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypesByAssociationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_types_by_association, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypesByAssociation(ctx, input)
				}
				var results []*svc.ListTypesByAssociationOutput
				p := svc.NewListTypesByAssociationPaginator(client, input)
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
		"put-graphql-api-environment-variables": {
			Name:   "put-graphql-api-environment-variables",
			Fields: fields_put_graphql_api_environment_variables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGraphqlApiEnvironmentVariablesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_graphql_api_environment_variables, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGraphqlApiEnvironmentVariables(ctx, input)
			},
		},
		"start-data-source-introspection": {
			Name:   "start-data-source-introspection",
			Fields: fields_start_data_source_introspection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataSourceIntrospectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_source_introspection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataSourceIntrospection(ctx, input)
			},
		},
		"start-schema-creation": {
			Name:   "start-schema-creation",
			Fields: fields_start_schema_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSchemaCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_schema_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSchemaCreation(ctx, input)
			},
		},
		"start-schema-merge": {
			Name:   "start-schema-merge",
			Fields: fields_start_schema_merge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSchemaMergeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_schema_merge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSchemaMerge(ctx, input)
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
		"update-api": {
			Name:   "update-api",
			Fields: fields_update_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApi(ctx, input)
			},
		},
		"update-api-cache": {
			Name:   "update-api-cache",
			Fields: fields_update_api_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApiCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_api_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApiCache(ctx, input)
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
		"update-channel-namespace": {
			Name:   "update-channel-namespace",
			Fields: fields_update_channel_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelNamespace(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
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
		"update-function": {
			Name:   "update-function",
			Fields: fields_update_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFunction(ctx, input)
			},
		},
		"update-graphql-api": {
			Name:   "update-graphql-api",
			Fields: fields_update_graphql_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGraphqlApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_graphql_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGraphqlApi(ctx, input)
			},
		},
		"update-resolver": {
			Name:   "update-resolver",
			Fields: fields_update_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResolver(ctx, input)
			},
		},
		"update-source-api-association": {
			Name:   "update-source-api-association",
			Fields: fields_update_source_api_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSourceApiAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_source_api_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSourceApiAssociation(ctx, input)
			},
		},
		"update-type": {
			Name:   "update-type",
			Fields: fields_update_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateType(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appsync", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
