package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

var fields_create_api = []leanruntime.Field{
	{Name: "ApiKeySelectionExpression", Flag: "api-key-selection-expression", Type: "*string", Required: false},
	{Name: "CorsConfiguration", Flag: "cors-configuration", Type: "*types.Cors", Required: false},
	{Name: "CredentialsArn", Flag: "credentials-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableExecuteApiEndpoint", Flag: "disable-execute-api-endpoint", Type: "*bool", Required: false},
	{Name: "DisableSchemaValidation", Flag: "disable-schema-validation", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProtocolType", Flag: "protocol-type", Type: "types.ProtocolType", Required: true},
	{Name: "RouteKey", Flag: "route-key", Type: "*string", Required: false},
	{Name: "RouteSelectionExpression", Flag: "route-selection-expression", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_create_api_mapping = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ApiMappingKey", Flag: "api-mapping-key", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "*string", Required: true},
}

var fields_create_authorizer = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AuthorizerCredentialsArn", Flag: "authorizer-credentials-arn", Type: "*string", Required: false},
	{Name: "AuthorizerPayloadFormatVersion", Flag: "authorizer-payload-format-version", Type: "*string", Required: false},
	{Name: "AuthorizerResultTtlInSeconds", Flag: "authorizer-result-ttl-in-seconds", Type: "*int32", Required: false},
	{Name: "AuthorizerType", Flag: "authorizer-type", Type: "types.AuthorizerType", Required: true},
	{Name: "AuthorizerUri", Flag: "authorizer-uri", Type: "*string", Required: false},
	{Name: "EnableSimpleResponses", Flag: "enable-simple-responses", Type: "*bool", Required: false},
	{Name: "IdentitySource", Flag: "identity-source", Type: "[]string", Required: true},
	{Name: "IdentityValidationExpression", Flag: "identity-validation-expression", Type: "*string", Required: false},
	{Name: "JwtConfiguration", Flag: "jwt-configuration", Type: "*types.JWTConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: false},
}

var fields_create_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameConfigurations", Flag: "domain-name-configurations", Type: "[]types.DomainNameConfiguration", Required: false},
	{Name: "MutualTlsAuthentication", Flag: "mutual-tls-authentication", Type: "*types.MutualTlsAuthenticationInput", Required: false},
	{Name: "RoutingMode", Flag: "routing-mode", Type: "types.RoutingMode", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_integration = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: false},
	{Name: "ContentHandlingStrategy", Flag: "content-handling-strategy", Type: "types.ContentHandlingStrategy", Required: false},
	{Name: "CredentialsArn", Flag: "credentials-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationMethod", Flag: "integration-method", Type: "*string", Required: false},
	{Name: "IntegrationSubtype", Flag: "integration-subtype", Type: "*string", Required: false},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: true},
	{Name: "IntegrationUri", Flag: "integration-uri", Type: "*string", Required: false},
	{Name: "PassthroughBehavior", Flag: "passthrough-behavior", Type: "types.PassthroughBehavior", Required: false},
	{Name: "PayloadFormatVersion", Flag: "payload-format-version", Type: "*string", Required: false},
	{Name: "RequestParameters", Flag: "request-parameters", Type: "map[string]string", Required: false},
	{Name: "RequestTemplates", Flag: "request-templates", Type: "map[string]string", Required: false},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]map[string]string", Required: false},
	{Name: "TemplateSelectionExpression", Flag: "template-selection-expression", Type: "*string", Required: false},
	{Name: "TimeoutInMillis", Flag: "timeout-in-millis", Type: "*int32", Required: false},
	{Name: "TlsConfig", Flag: "tls-config", Type: "*types.TlsConfigInput", Required: false},
}

var fields_create_integration_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ContentHandlingStrategy", Flag: "content-handling-strategy", Type: "types.ContentHandlingStrategy", Required: false},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
	{Name: "IntegrationResponseKey", Flag: "integration-response-key", Type: "*string", Required: true},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]string", Required: false},
	{Name: "ResponseTemplates", Flag: "response-templates", Type: "map[string]string", Required: false},
	{Name: "TemplateSelectionExpression", Flag: "template-selection-expression", Type: "*string", Required: false},
}

var fields_create_model = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: true},
}

var fields_create_portal = []leanruntime.Field{
	{Name: "Authorization", Flag: "authorization", Type: "*types.Authorization", Required: true},
	{Name: "EndpointConfiguration", Flag: "endpoint-configuration", Type: "*types.EndpointConfigurationRequest", Required: true},
	{Name: "IncludedPortalProductArns", Flag: "included-portal-product-arns", Type: "[]string", Required: false},
	{Name: "LogoUri", Flag: "logo-uri", Type: "*string", Required: false},
	{Name: "PortalContent", Flag: "portal-content", Type: "*types.PortalContent", Required: true},
	{Name: "RumAppMonitorName", Flag: "rum-app-monitor-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_portal_product = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_product_page = []leanruntime.Field{
	{Name: "DisplayContent", Flag: "display-content", Type: "*types.DisplayContent", Required: true},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
}

var fields_create_product_rest_endpoint_page = []leanruntime.Field{
	{Name: "DisplayContent", Flag: "display-content", Type: "*types.EndpointDisplayContent", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "RestEndpointIdentifier", Flag: "rest-endpoint-identifier", Type: "*types.RestEndpointIdentifier", Required: true},
	{Name: "TryItState", Flag: "try-it-state", Type: "types.TryItState", Required: false},
}

var fields_create_route = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ApiKeyRequired", Flag: "api-key-required", Type: "*bool", Required: false},
	{Name: "AuthorizationScopes", Flag: "authorization-scopes", Type: "[]string", Required: false},
	{Name: "AuthorizationType", Flag: "authorization-type", Type: "types.AuthorizationType", Required: false},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: false},
	{Name: "ModelSelectionExpression", Flag: "model-selection-expression", Type: "*string", Required: false},
	{Name: "OperationName", Flag: "operation-name", Type: "*string", Required: false},
	{Name: "RequestModels", Flag: "request-models", Type: "map[string]string", Required: false},
	{Name: "RequestParameters", Flag: "request-parameters", Type: "map[string]types.ParameterConstraints", Required: false},
	{Name: "RouteKey", Flag: "route-key", Type: "*string", Required: true},
	{Name: "RouteResponseSelectionExpression", Flag: "route-response-selection-expression", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: false},
}

var fields_create_route_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ModelSelectionExpression", Flag: "model-selection-expression", Type: "*string", Required: false},
	{Name: "ResponseModels", Flag: "response-models", Type: "map[string]string", Required: false},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]types.ParameterConstraints", Required: false},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
	{Name: "RouteResponseKey", Flag: "route-response-key", Type: "*string", Required: true},
}

var fields_create_routing_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.RoutingRuleAction", Required: true},
	{Name: "Conditions", Flag: "conditions", Type: "[]types.RoutingRuleCondition", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
}

var fields_create_stage = []leanruntime.Field{
	{Name: "AccessLogSettings", Flag: "access-log-settings", Type: "*types.AccessLogSettings", Required: false},
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AutoDeploy", Flag: "auto-deploy", Type: "*bool", Required: false},
	{Name: "ClientCertificateId", Flag: "client-certificate-id", Type: "*string", Required: false},
	{Name: "DefaultRouteSettings", Flag: "default-route-settings", Type: "*types.RouteSettings", Required: false},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RouteSettings", Flag: "route-settings", Type: "map[string]types.RouteSettings", Required: false},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "StageVariables", Flag: "stage-variables", Type: "map[string]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_vpc_link = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_access_log_settings = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_delete_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_delete_api_mapping = []leanruntime.Field{
	{Name: "ApiMappingId", Flag: "api-mapping-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_authorizer = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
}

var fields_delete_cors_configuration = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_delete_deployment = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_delete_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_integration = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
}

var fields_delete_integration_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
	{Name: "IntegrationResponseId", Flag: "integration-response-id", Type: "*string", Required: true},
}

var fields_delete_model = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_delete_portal = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_delete_portal_product = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
}

var fields_delete_portal_product_sharing_policy = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
}

var fields_delete_product_page = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ProductPageId", Flag: "product-page-id", Type: "*string", Required: true},
}

var fields_delete_product_rest_endpoint_page = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ProductRestEndpointPageId", Flag: "product-rest-endpoint-page-id", Type: "*string", Required: true},
}

var fields_delete_route = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
}

var fields_delete_route_request_parameter = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "RequestParameterKey", Flag: "request-parameter-key", Type: "*string", Required: true},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
}

var fields_delete_route_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
	{Name: "RouteResponseId", Flag: "route-response-id", Type: "*string", Required: true},
}

var fields_delete_route_settings = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "RouteKey", Flag: "route-key", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_delete_routing_rule = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "RoutingRuleId", Flag: "routing-rule-id", Type: "*string", Required: true},
}

var fields_delete_stage = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_delete_vpc_link = []leanruntime.Field{
	{Name: "VpcLinkId", Flag: "vpc-link-id", Type: "*string", Required: true},
}

var fields_disable_portal = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_export_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ExportVersion", Flag: "export-version", Type: "*string", Required: false},
	{Name: "IncludeExtensions", Flag: "include-extensions", Type: "*bool", Required: false},
	{Name: "OutputType", Flag: "output-type", Type: "*string", Required: true},
	{Name: "Specification", Flag: "specification", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: false},
}

var fields_get_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
}

var fields_get_api_mapping = []leanruntime.Field{
	{Name: "ApiMappingId", Flag: "api-mapping-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_api_mappings = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_apis = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_authorizer = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
}

var fields_get_authorizers = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_get_deployments = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_domain_names = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_integration = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
}

var fields_get_integration_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
	{Name: "IntegrationResponseId", Flag: "integration-response-id", Type: "*string", Required: true},
}

var fields_get_integration_responses = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_integrations = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_model = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_get_model_template = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_get_models = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_portal = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_get_portal_product = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ResourceOwnerAccountId", Flag: "resource-owner-account-id", Type: "*string", Required: false},
}

var fields_get_portal_product_sharing_policy = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
}

var fields_get_product_page = []leanruntime.Field{
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ProductPageId", Flag: "product-page-id", Type: "*string", Required: true},
	{Name: "ResourceOwnerAccountId", Flag: "resource-owner-account-id", Type: "*string", Required: false},
}

var fields_get_product_rest_endpoint_page = []leanruntime.Field{
	{Name: "IncludeRawDisplayContent", Flag: "include-raw-display-content", Type: "*string", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ProductRestEndpointPageId", Flag: "product-rest-endpoint-page-id", Type: "*string", Required: true},
	{Name: "ResourceOwnerAccountId", Flag: "resource-owner-account-id", Type: "*string", Required: false},
}

var fields_get_route = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
}

var fields_get_route_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
	{Name: "RouteResponseId", Flag: "route-response-id", Type: "*string", Required: true},
}

var fields_get_route_responses = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
}

var fields_get_routes = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_routing_rule = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "RoutingRuleId", Flag: "routing-rule-id", Type: "*string", Required: true},
}

var fields_get_stage = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_get_stages = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_vpc_link = []leanruntime.Field{
	{Name: "VpcLinkId", Flag: "vpc-link-id", Type: "*string", Required: true},
}

var fields_get_vpc_links = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_import_api = []leanruntime.Field{
	{Name: "Basepath", Flag: "basepath", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "*string", Required: true},
	{Name: "FailOnWarnings", Flag: "fail-on-warnings", Type: "*bool", Required: false},
}

var fields_list_portal_products = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "*string", Required: false},
}

var fields_list_portals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_product_pages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ResourceOwnerAccountId", Flag: "resource-owner-account-id", Type: "*string", Required: false},
}

var fields_list_product_rest_endpoint_pages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ResourceOwnerAccountId", Flag: "resource-owner-account-id", Type: "*string", Required: false},
}

var fields_list_routing_rules = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_preview_portal = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_publish_portal = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_put_portal_product_sharing_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
}

var fields_put_routing_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.RoutingRuleAction", Required: true},
	{Name: "Conditions", Flag: "conditions", Type: "[]types.RoutingRuleCondition", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameId", Flag: "domain-name-id", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "RoutingRuleId", Flag: "routing-rule-id", Type: "*string", Required: true},
}

var fields_reimport_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "Basepath", Flag: "basepath", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "*string", Required: true},
	{Name: "FailOnWarnings", Flag: "fail-on-warnings", Type: "*bool", Required: false},
}

var fields_reset_authorizers_cache = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_api = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ApiKeySelectionExpression", Flag: "api-key-selection-expression", Type: "*string", Required: false},
	{Name: "CorsConfiguration", Flag: "cors-configuration", Type: "*types.Cors", Required: false},
	{Name: "CredentialsArn", Flag: "credentials-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableExecuteApiEndpoint", Flag: "disable-execute-api-endpoint", Type: "*bool", Required: false},
	{Name: "DisableSchemaValidation", Flag: "disable-schema-validation", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RouteKey", Flag: "route-key", Type: "*string", Required: false},
	{Name: "RouteSelectionExpression", Flag: "route-selection-expression", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_api_mapping = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ApiMappingId", Flag: "api-mapping-id", Type: "*string", Required: true},
	{Name: "ApiMappingKey", Flag: "api-mapping-key", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "*string", Required: false},
}

var fields_update_authorizer = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AuthorizerCredentialsArn", Flag: "authorizer-credentials-arn", Type: "*string", Required: false},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: true},
	{Name: "AuthorizerPayloadFormatVersion", Flag: "authorizer-payload-format-version", Type: "*string", Required: false},
	{Name: "AuthorizerResultTtlInSeconds", Flag: "authorizer-result-ttl-in-seconds", Type: "*int32", Required: false},
	{Name: "AuthorizerType", Flag: "authorizer-type", Type: "types.AuthorizerType", Required: false},
	{Name: "AuthorizerUri", Flag: "authorizer-uri", Type: "*string", Required: false},
	{Name: "EnableSimpleResponses", Flag: "enable-simple-responses", Type: "*bool", Required: false},
	{Name: "IdentitySource", Flag: "identity-source", Type: "[]string", Required: false},
	{Name: "IdentityValidationExpression", Flag: "identity-validation-expression", Type: "*string", Required: false},
	{Name: "JwtConfiguration", Flag: "jwt-configuration", Type: "*types.JWTConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_deployment = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_domain_name = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainNameConfigurations", Flag: "domain-name-configurations", Type: "[]types.DomainNameConfiguration", Required: false},
	{Name: "MutualTlsAuthentication", Flag: "mutual-tls-authentication", Type: "*types.MutualTlsAuthenticationInput", Required: false},
	{Name: "RoutingMode", Flag: "routing-mode", Type: "types.RoutingMode", Required: false},
}

var fields_update_integration = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: false},
	{Name: "ContentHandlingStrategy", Flag: "content-handling-strategy", Type: "types.ContentHandlingStrategy", Required: false},
	{Name: "CredentialsArn", Flag: "credentials-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
	{Name: "IntegrationMethod", Flag: "integration-method", Type: "*string", Required: false},
	{Name: "IntegrationSubtype", Flag: "integration-subtype", Type: "*string", Required: false},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: false},
	{Name: "IntegrationUri", Flag: "integration-uri", Type: "*string", Required: false},
	{Name: "PassthroughBehavior", Flag: "passthrough-behavior", Type: "types.PassthroughBehavior", Required: false},
	{Name: "PayloadFormatVersion", Flag: "payload-format-version", Type: "*string", Required: false},
	{Name: "RequestParameters", Flag: "request-parameters", Type: "map[string]string", Required: false},
	{Name: "RequestTemplates", Flag: "request-templates", Type: "map[string]string", Required: false},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]map[string]string", Required: false},
	{Name: "TemplateSelectionExpression", Flag: "template-selection-expression", Type: "*string", Required: false},
	{Name: "TimeoutInMillis", Flag: "timeout-in-millis", Type: "*int32", Required: false},
	{Name: "TlsConfig", Flag: "tls-config", Type: "*types.TlsConfigInput", Required: false},
}

var fields_update_integration_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ContentHandlingStrategy", Flag: "content-handling-strategy", Type: "types.ContentHandlingStrategy", Required: false},
	{Name: "IntegrationId", Flag: "integration-id", Type: "*string", Required: true},
	{Name: "IntegrationResponseId", Flag: "integration-response-id", Type: "*string", Required: true},
	{Name: "IntegrationResponseKey", Flag: "integration-response-key", Type: "*string", Required: false},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]string", Required: false},
	{Name: "ResponseTemplates", Flag: "response-templates", Type: "map[string]string", Required: false},
	{Name: "TemplateSelectionExpression", Flag: "template-selection-expression", Type: "*string", Required: false},
}

var fields_update_model = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: false},
}

var fields_update_portal = []leanruntime.Field{
	{Name: "Authorization", Flag: "authorization", Type: "*types.Authorization", Required: false},
	{Name: "EndpointConfiguration", Flag: "endpoint-configuration", Type: "*types.EndpointConfigurationRequest", Required: false},
	{Name: "IncludedPortalProductArns", Flag: "included-portal-product-arns", Type: "[]string", Required: false},
	{Name: "LogoUri", Flag: "logo-uri", Type: "*string", Required: false},
	{Name: "PortalContent", Flag: "portal-content", Type: "*types.PortalContent", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
	{Name: "RumAppMonitorName", Flag: "rum-app-monitor-name", Type: "*string", Required: false},
}

var fields_update_portal_product = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DisplayOrder", Flag: "display-order", Type: "*types.DisplayOrder", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
}

var fields_update_product_page = []leanruntime.Field{
	{Name: "DisplayContent", Flag: "display-content", Type: "*types.DisplayContent", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ProductPageId", Flag: "product-page-id", Type: "*string", Required: true},
}

var fields_update_product_rest_endpoint_page = []leanruntime.Field{
	{Name: "DisplayContent", Flag: "display-content", Type: "*types.EndpointDisplayContent", Required: false},
	{Name: "PortalProductId", Flag: "portal-product-id", Type: "*string", Required: true},
	{Name: "ProductRestEndpointPageId", Flag: "product-rest-endpoint-page-id", Type: "*string", Required: true},
	{Name: "TryItState", Flag: "try-it-state", Type: "types.TryItState", Required: false},
}

var fields_update_route = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ApiKeyRequired", Flag: "api-key-required", Type: "*bool", Required: false},
	{Name: "AuthorizationScopes", Flag: "authorization-scopes", Type: "[]string", Required: false},
	{Name: "AuthorizationType", Flag: "authorization-type", Type: "types.AuthorizationType", Required: false},
	{Name: "AuthorizerId", Flag: "authorizer-id", Type: "*string", Required: false},
	{Name: "ModelSelectionExpression", Flag: "model-selection-expression", Type: "*string", Required: false},
	{Name: "OperationName", Flag: "operation-name", Type: "*string", Required: false},
	{Name: "RequestModels", Flag: "request-models", Type: "map[string]string", Required: false},
	{Name: "RequestParameters", Flag: "request-parameters", Type: "map[string]types.ParameterConstraints", Required: false},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
	{Name: "RouteKey", Flag: "route-key", Type: "*string", Required: false},
	{Name: "RouteResponseSelectionExpression", Flag: "route-response-selection-expression", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: false},
}

var fields_update_route_response = []leanruntime.Field{
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "ModelSelectionExpression", Flag: "model-selection-expression", Type: "*string", Required: false},
	{Name: "ResponseModels", Flag: "response-models", Type: "map[string]string", Required: false},
	{Name: "ResponseParameters", Flag: "response-parameters", Type: "map[string]types.ParameterConstraints", Required: false},
	{Name: "RouteId", Flag: "route-id", Type: "*string", Required: true},
	{Name: "RouteResponseId", Flag: "route-response-id", Type: "*string", Required: true},
	{Name: "RouteResponseKey", Flag: "route-response-key", Type: "*string", Required: false},
}

var fields_update_stage = []leanruntime.Field{
	{Name: "AccessLogSettings", Flag: "access-log-settings", Type: "*types.AccessLogSettings", Required: false},
	{Name: "ApiId", Flag: "api-id", Type: "*string", Required: true},
	{Name: "AutoDeploy", Flag: "auto-deploy", Type: "*bool", Required: false},
	{Name: "ClientCertificateId", Flag: "client-certificate-id", Type: "*string", Required: false},
	{Name: "DefaultRouteSettings", Flag: "default-route-settings", Type: "*types.RouteSettings", Required: false},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RouteSettings", Flag: "route-settings", Type: "map[string]types.RouteSettings", Required: false},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "StageVariables", Flag: "stage-variables", Type: "map[string]string", Required: false},
}

var fields_update_vpc_link = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "VpcLinkId", Flag: "vpc-link-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"create-api-mapping": {
			Name:   "create-api-mapping",
			Fields: fields_create_api_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApiMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApiMapping(ctx, input)
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
		"create-integration": {
			Name:   "create-integration",
			Fields: fields_create_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegration(ctx, input)
			},
		},
		"create-integration-response": {
			Name:   "create-integration-response",
			Fields: fields_create_integration_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegrationResponse(ctx, input)
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
		"create-portal": {
			Name:   "create-portal",
			Fields: fields_create_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePortal(ctx, input)
			},
		},
		"create-portal-product": {
			Name:   "create-portal-product",
			Fields: fields_create_portal_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePortalProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_portal_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePortalProduct(ctx, input)
			},
		},
		"create-product-page": {
			Name:   "create-product-page",
			Fields: fields_create_product_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProductPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_product_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProductPage(ctx, input)
			},
		},
		"create-product-rest-endpoint-page": {
			Name:   "create-product-rest-endpoint-page",
			Fields: fields_create_product_rest_endpoint_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProductRestEndpointPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_product_rest_endpoint_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProductRestEndpointPage(ctx, input)
			},
		},
		"create-route": {
			Name:   "create-route",
			Fields: fields_create_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoute(ctx, input)
			},
		},
		"create-route-response": {
			Name:   "create-route-response",
			Fields: fields_create_route_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouteResponse(ctx, input)
			},
		},
		"create-routing-rule": {
			Name:   "create-routing-rule",
			Fields: fields_create_routing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoutingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_routing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoutingRule(ctx, input)
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
		"delete-access-log-settings": {
			Name:   "delete-access-log-settings",
			Fields: fields_delete_access_log_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessLogSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_log_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessLogSettings(ctx, input)
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
		"delete-api-mapping": {
			Name:   "delete-api-mapping",
			Fields: fields_delete_api_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApiMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApiMapping(ctx, input)
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
		"delete-cors-configuration": {
			Name:   "delete-cors-configuration",
			Fields: fields_delete_cors_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCorsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cors_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCorsConfiguration(ctx, input)
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
		"delete-portal": {
			Name:   "delete-portal",
			Fields: fields_delete_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortal(ctx, input)
			},
		},
		"delete-portal-product": {
			Name:   "delete-portal-product",
			Fields: fields_delete_portal_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortalProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portal_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortalProduct(ctx, input)
			},
		},
		"delete-portal-product-sharing-policy": {
			Name:   "delete-portal-product-sharing-policy",
			Fields: fields_delete_portal_product_sharing_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortalProductSharingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portal_product_sharing_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortalProductSharingPolicy(ctx, input)
			},
		},
		"delete-product-page": {
			Name:   "delete-product-page",
			Fields: fields_delete_product_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProductPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_product_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProductPage(ctx, input)
			},
		},
		"delete-product-rest-endpoint-page": {
			Name:   "delete-product-rest-endpoint-page",
			Fields: fields_delete_product_rest_endpoint_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProductRestEndpointPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_product_rest_endpoint_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProductRestEndpointPage(ctx, input)
			},
		},
		"delete-route": {
			Name:   "delete-route",
			Fields: fields_delete_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoute(ctx, input)
			},
		},
		"delete-route-request-parameter": {
			Name:   "delete-route-request-parameter",
			Fields: fields_delete_route_request_parameter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteRequestParameterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_request_parameter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteRequestParameter(ctx, input)
			},
		},
		"delete-route-response": {
			Name:   "delete-route-response",
			Fields: fields_delete_route_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteResponse(ctx, input)
			},
		},
		"delete-route-settings": {
			Name:   "delete-route-settings",
			Fields: fields_delete_route_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteSettings(ctx, input)
			},
		},
		"delete-routing-rule": {
			Name:   "delete-routing-rule",
			Fields: fields_delete_routing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoutingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_routing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoutingRule(ctx, input)
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
		"disable-portal": {
			Name:   "disable-portal",
			Fields: fields_disable_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisablePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisablePortal(ctx, input)
			},
		},
		"export-api": {
			Name:   "export-api",
			Fields: fields_export_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportApi(ctx, input)
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
		"get-api-mapping": {
			Name:   "get-api-mapping",
			Fields: fields_get_api_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApiMapping(ctx, input)
			},
		},
		"get-api-mappings": {
			Name:   "get-api-mappings",
			Fields: fields_get_api_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApiMappingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_api_mappings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApiMappings(ctx, input)
			},
		},
		"get-apis": {
			Name:   "get-apis",
			Fields: fields_get_apis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApis(ctx, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_get_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployments(ctx, input)
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
		"get-domain-names": {
			Name:   "get-domain-names",
			Fields: fields_get_domain_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainNamesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_names, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainNames(ctx, input)
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
		"get-integration-responses": {
			Name:   "get-integration-responses",
			Fields: fields_get_integration_responses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationResponsesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration_responses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegrationResponses(ctx, input)
			},
		},
		"get-integrations": {
			Name:   "get-integrations",
			Fields: fields_get_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegrations(ctx, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_get_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModels(ctx, input)
			},
		},
		"get-portal": {
			Name:   "get-portal",
			Fields: fields_get_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortal(ctx, input)
			},
		},
		"get-portal-product": {
			Name:   "get-portal-product",
			Fields: fields_get_portal_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortalProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portal_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortalProduct(ctx, input)
			},
		},
		"get-portal-product-sharing-policy": {
			Name:   "get-portal-product-sharing-policy",
			Fields: fields_get_portal_product_sharing_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortalProductSharingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portal_product_sharing_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortalProductSharingPolicy(ctx, input)
			},
		},
		"get-product-page": {
			Name:   "get-product-page",
			Fields: fields_get_product_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProductPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_product_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProductPage(ctx, input)
			},
		},
		"get-product-rest-endpoint-page": {
			Name:   "get-product-rest-endpoint-page",
			Fields: fields_get_product_rest_endpoint_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProductRestEndpointPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_product_rest_endpoint_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProductRestEndpointPage(ctx, input)
			},
		},
		"get-route": {
			Name:   "get-route",
			Fields: fields_get_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoute(ctx, input)
			},
		},
		"get-route-response": {
			Name:   "get-route-response",
			Fields: fields_get_route_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouteResponse(ctx, input)
			},
		},
		"get-route-responses": {
			Name:   "get-route-responses",
			Fields: fields_get_route_responses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteResponsesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route_responses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouteResponses(ctx, input)
			},
		},
		"get-routes": {
			Name:   "get-routes",
			Fields: fields_get_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoutes(ctx, input)
			},
		},
		"get-routing-rule": {
			Name:   "get-routing-rule",
			Fields: fields_get_routing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoutingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_routing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoutingRule(ctx, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_get_vpc_links, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpcLinks(ctx, input)
			},
		},
		"import-api": {
			Name:   "import-api",
			Fields: fields_import_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportApi(ctx, input)
			},
		},
		"list-portal-products": {
			Name:   "list-portal-products",
			Fields: fields_list_portal_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortalProductsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_portal_products, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPortalProducts(ctx, input)
			},
		},
		"list-portals": {
			Name:   "list-portals",
			Fields: fields_list_portals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortalsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_portals, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPortals(ctx, input)
			},
		},
		"list-product-pages": {
			Name:   "list-product-pages",
			Fields: fields_list_product_pages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProductPagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_product_pages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProductPages(ctx, input)
			},
		},
		"list-product-rest-endpoint-pages": {
			Name:   "list-product-rest-endpoint-pages",
			Fields: fields_list_product_rest_endpoint_pages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProductRestEndpointPagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_product_rest_endpoint_pages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProductRestEndpointPages(ctx, input)
			},
		},
		"list-routing-rules": {
			Name:   "list-routing-rules",
			Fields: fields_list_routing_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutingRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routing_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutingRules(ctx, input)
				}
				var results []*svc.ListRoutingRulesOutput
				p := svc.NewListRoutingRulesPaginator(client, input)
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
		"preview-portal": {
			Name:   "preview-portal",
			Fields: fields_preview_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PreviewPortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_preview_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PreviewPortal(ctx, input)
			},
		},
		"publish-portal": {
			Name:   "publish-portal",
			Fields: fields_publish_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishPortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishPortal(ctx, input)
			},
		},
		"put-portal-product-sharing-policy": {
			Name:   "put-portal-product-sharing-policy",
			Fields: fields_put_portal_product_sharing_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPortalProductSharingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_portal_product_sharing_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPortalProductSharingPolicy(ctx, input)
			},
		},
		"put-routing-rule": {
			Name:   "put-routing-rule",
			Fields: fields_put_routing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRoutingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_routing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRoutingRule(ctx, input)
			},
		},
		"reimport-api": {
			Name:   "reimport-api",
			Fields: fields_reimport_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReimportApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reimport_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReimportApi(ctx, input)
			},
		},
		"reset-authorizers-cache": {
			Name:   "reset-authorizers-cache",
			Fields: fields_reset_authorizers_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetAuthorizersCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_authorizers_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetAuthorizersCache(ctx, input)
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
		"update-api-mapping": {
			Name:   "update-api-mapping",
			Fields: fields_update_api_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApiMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_api_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApiMapping(ctx, input)
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
		"update-portal": {
			Name:   "update-portal",
			Fields: fields_update_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePortal(ctx, input)
			},
		},
		"update-portal-product": {
			Name:   "update-portal-product",
			Fields: fields_update_portal_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePortalProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_portal_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePortalProduct(ctx, input)
			},
		},
		"update-product-page": {
			Name:   "update-product-page",
			Fields: fields_update_product_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProductPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_product_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProductPage(ctx, input)
			},
		},
		"update-product-rest-endpoint-page": {
			Name:   "update-product-rest-endpoint-page",
			Fields: fields_update_product_rest_endpoint_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProductRestEndpointPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_product_rest_endpoint_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProductRestEndpointPage(ctx, input)
			},
		},
		"update-route": {
			Name:   "update-route",
			Fields: fields_update_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoute(ctx, input)
			},
		},
		"update-route-response": {
			Name:   "update-route-response",
			Fields: fields_update_route_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouteResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_route_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRouteResponse(ctx, input)
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
	if err := leanruntime.Execute("apigatewayv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
