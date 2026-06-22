package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

var fields_associate_alias = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "TargetDistributionId", Flag: "target-distribution-id", Type: "*string", Required: true},
}

var fields_associate_distribution_tenant_web_acl = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "WebACLArn", Flag: "web-acl-arn", Type: "*string", Required: true},
}

var fields_associate_distribution_web_acl = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "WebACLArn", Flag: "web-acl-arn", Type: "*string", Required: true},
}

var fields_copy_distribution = []leanruntime.Field{
	{Name: "CallerReference", Flag: "caller-reference", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "PrimaryDistributionId", Flag: "primary-distribution-id", Type: "*string", Required: true},
	{Name: "Staging", Flag: "staging", Type: "*bool", Required: false},
}

var fields_create_anycast_ip_list = []leanruntime.Field{
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "IpCount", Flag: "ip-count", Type: "*int32", Required: true},
	{Name: "IpamCidrConfigs", Flag: "ipam-cidr-configs", Type: "[]types.IpamCidrConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: false},
}

var fields_create_cache_policy = []leanruntime.Field{
	{Name: "CachePolicyConfig", Flag: "cache-policy-config", Type: "*types.CachePolicyConfig", Required: true},
}

var fields_create_cloud_front_origin_access_identity = []leanruntime.Field{
	{Name: "CloudFrontOriginAccessIdentityConfig", Flag: "cloud-front-origin-access-identity-config", Type: "*types.CloudFrontOriginAccessIdentityConfig", Required: true},
}

var fields_create_connection_function = []leanruntime.Field{
	{Name: "ConnectionFunctionCode", Flag: "connection-function-code", Type: "[]byte", Required: true},
	{Name: "ConnectionFunctionConfig", Flag: "connection-function-config", Type: "*types.FunctionConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: false},
}

var fields_create_connection_group = []leanruntime.Field{
	{Name: "AnycastIpListId", Flag: "anycast-ip-list-id", Type: "*string", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Ipv6Enabled", Flag: "ipv6-enabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: false},
}

var fields_create_continuous_deployment_policy = []leanruntime.Field{
	{Name: "ContinuousDeploymentPolicyConfig", Flag: "continuous-deployment-policy-config", Type: "*types.ContinuousDeploymentPolicyConfig", Required: true},
}

var fields_create_distribution = []leanruntime.Field{
	{Name: "DistributionConfig", Flag: "distribution-config", Type: "*types.DistributionConfig", Required: true},
}

var fields_create_distribution_tenant = []leanruntime.Field{
	{Name: "ConnectionGroupId", Flag: "connection-group-id", Type: "*string", Required: false},
	{Name: "Customizations", Flag: "customizations", Type: "*types.Customizations", Required: false},
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
	{Name: "Domains", Flag: "domains", Type: "[]types.DomainItem", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "ManagedCertificateRequest", Flag: "managed-certificate-request", Type: "*types.ManagedCertificateRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: false},
}

var fields_create_distribution_with_tags = []leanruntime.Field{
	{Name: "DistributionConfigWithTags", Flag: "distribution-config-with-tags", Type: "*types.DistributionConfigWithTags", Required: true},
}

var fields_create_field_level_encryption_config = []leanruntime.Field{
	{Name: "FieldLevelEncryptionConfig", Flag: "field-level-encryption-config", Type: "*types.FieldLevelEncryptionConfig", Required: true},
}

var fields_create_field_level_encryption_profile = []leanruntime.Field{
	{Name: "FieldLevelEncryptionProfileConfig", Flag: "field-level-encryption-profile-config", Type: "*types.FieldLevelEncryptionProfileConfig", Required: true},
}

var fields_create_function = []leanruntime.Field{
	{Name: "FunctionCode", Flag: "function-code", Type: "[]byte", Required: true},
	{Name: "FunctionConfig", Flag: "function-config", Type: "*types.FunctionConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_invalidation = []leanruntime.Field{
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
	{Name: "InvalidationBatch", Flag: "invalidation-batch", Type: "*types.InvalidationBatch", Required: true},
}

var fields_create_invalidation_for_distribution_tenant = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "InvalidationBatch", Flag: "invalidation-batch", Type: "*types.InvalidationBatch", Required: true},
}

var fields_create_key_group = []leanruntime.Field{
	{Name: "KeyGroupConfig", Flag: "key-group-config", Type: "*types.KeyGroupConfig", Required: true},
}

var fields_create_key_value_store = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "ImportSource", Flag: "import-source", Type: "*types.ImportSource", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_monitoring_subscription = []leanruntime.Field{
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
	{Name: "MonitoringSubscription", Flag: "monitoring-subscription", Type: "*types.MonitoringSubscription", Required: true},
}

var fields_create_origin_access_control = []leanruntime.Field{
	{Name: "OriginAccessControlConfig", Flag: "origin-access-control-config", Type: "*types.OriginAccessControlConfig", Required: true},
}

var fields_create_origin_request_policy = []leanruntime.Field{
	{Name: "OriginRequestPolicyConfig", Flag: "origin-request-policy-config", Type: "*types.OriginRequestPolicyConfig", Required: true},
}

var fields_create_public_key = []leanruntime.Field{
	{Name: "PublicKeyConfig", Flag: "public-key-config", Type: "*types.PublicKeyConfig", Required: true},
}

var fields_create_realtime_log_config = []leanruntime.Field{
	{Name: "EndPoints", Flag: "end-points", Type: "[]types.EndPoint", Required: true},
	{Name: "Fields", Flag: "fields", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SamplingRate", Flag: "sampling-rate", Type: "*int64", Required: true},
}

var fields_create_response_headers_policy = []leanruntime.Field{
	{Name: "ResponseHeadersPolicyConfig", Flag: "response-headers-policy-config", Type: "*types.ResponseHeadersPolicyConfig", Required: true},
}

var fields_create_streaming_distribution = []leanruntime.Field{
	{Name: "StreamingDistributionConfig", Flag: "streaming-distribution-config", Type: "*types.StreamingDistributionConfig", Required: true},
}

var fields_create_streaming_distribution_with_tags = []leanruntime.Field{
	{Name: "StreamingDistributionConfigWithTags", Flag: "streaming-distribution-config-with-tags", Type: "*types.StreamingDistributionConfigWithTags", Required: true},
}

var fields_create_trust_store = []leanruntime.Field{
	{Name: "CaCertificatesBundleSource", Flag: "ca-certificates-bundle-source", Type: "types.CaCertificatesBundleSource", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: false},
}

var fields_create_vpc_origin = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: false},
	{Name: "VpcOriginEndpointConfig", Flag: "vpc-origin-endpoint-config", Type: "*types.VpcOriginEndpointConfig", Required: true},
}

var fields_delete_anycast_ip_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_delete_cache_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_cloud_front_origin_access_identity = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_connection_function = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_delete_connection_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_delete_continuous_deployment_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_distribution = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_distribution_tenant = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_delete_field_level_encryption_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_field_level_encryption_profile = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_function = []leanruntime.Field{
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_key_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_key_value_store = []leanruntime.Field{
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_monitoring_subscription = []leanruntime.Field{
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
}

var fields_delete_origin_access_control = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_origin_request_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_public_key = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_realtime_log_config = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_response_headers_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_streaming_distribution = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_delete_trust_store = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_delete_vpc_origin = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_describe_connection_function = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_describe_function = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_describe_key_value_store = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_disassociate_distribution_tenant_web_acl = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_disassociate_distribution_web_acl = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_get_anycast_ip_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_cache_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_cache_policy_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_cloud_front_origin_access_identity = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_cloud_front_origin_access_identity_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_connection_function = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_get_connection_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_connection_group_by_routing_endpoint = []leanruntime.Field{
	{Name: "RoutingEndpoint", Flag: "routing-endpoint", Type: "*string", Required: true},
}

var fields_get_continuous_deployment_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_continuous_deployment_policy_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_distribution = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_distribution_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_distribution_tenant = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_distribution_tenant_by_domain = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_get_field_level_encryption = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_field_level_encryption_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_field_level_encryption_profile = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_field_level_encryption_profile_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_function = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_get_invalidation = []leanruntime.Field{
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_invalidation_for_distribution_tenant = []leanruntime.Field{
	{Name: "DistributionTenantId", Flag: "distribution-tenant-id", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_key_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_key_group_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_managed_certificate_details = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_monitoring_subscription = []leanruntime.Field{
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
}

var fields_get_origin_access_control = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_origin_access_control_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_origin_request_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_origin_request_policy_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_public_key = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_public_key_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_realtime_log_config = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_response_headers_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_response_headers_policy_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_streaming_distribution = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_streaming_distribution_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_trust_store = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_vpc_origin = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_anycast_ip_lists = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_cache_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CachePolicyType", Required: false},
}

var fields_list_cloud_front_origin_access_identities = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_conflicting_aliases = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_connection_functions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_list_connection_groups = []leanruntime.Field{
	{Name: "AssociationFilter", Flag: "association-filter", Type: "*types.ConnectionGroupAssociationFilter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_continuous_deployment_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distribution_tenants = []leanruntime.Field{
	{Name: "AssociationFilter", Flag: "association-filter", Type: "*types.DistributionTenantAssociationFilter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distribution_tenants_by_customization = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "WebACLArn", Flag: "web-acl-arn", Type: "*string", Required: false},
}

var fields_list_distributions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distributions_by_anycast_ip_list_id = []leanruntime.Field{
	{Name: "AnycastIpListId", Flag: "anycast-ip-list-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distributions_by_cache_policy_id = []leanruntime.Field{
	{Name: "CachePolicyId", Flag: "cache-policy-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distributions_by_connection_function = []leanruntime.Field{
	{Name: "ConnectionFunctionIdentifier", Flag: "connection-function-identifier", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distributions_by_connection_mode = []leanruntime.Field{
	{Name: "ConnectionMode", Flag: "connection-mode", Type: "types.ConnectionMode", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distributions_by_key_group = []leanruntime.Field{
	{Name: "KeyGroupId", Flag: "key-group-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_distributions_by_origin_request_policy_id = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "OriginRequestPolicyId", Flag: "origin-request-policy-id", Type: "*string", Required: true},
}

var fields_list_distributions_by_owned_resource = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_distributions_by_realtime_log_config = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "RealtimeLogConfigArn", Flag: "realtime-log-config-arn", Type: "*string", Required: false},
	{Name: "RealtimeLogConfigName", Flag: "realtime-log-config-name", Type: "*string", Required: false},
}

var fields_list_distributions_by_response_headers_policy_id = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "ResponseHeadersPolicyId", Flag: "response-headers-policy-id", Type: "*string", Required: true},
}

var fields_list_distributions_by_trust_store = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "TrustStoreIdentifier", Flag: "trust-store-identifier", Type: "*string", Required: true},
}

var fields_list_distributions_by_vpc_origin_id = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "VpcOriginId", Flag: "vpc-origin-id", Type: "*string", Required: true},
}

var fields_list_distributions_by_web_aclid = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_list_domain_conflicts = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainControlValidationResource", Flag: "domain-control-validation-resource", Type: "*types.DistributionResourceId", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_field_level_encryption_configs = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_field_level_encryption_profiles = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_functions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_list_invalidations = []leanruntime.Field{
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_invalidations_for_distribution_tenant = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_key_groups = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_key_value_stores = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_list_origin_access_controls = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_origin_request_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Type", Flag: "type", Type: "types.OriginRequestPolicyType", Required: false},
}

var fields_list_public_keys = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_realtime_log_configs = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_response_headers_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ResponseHeadersPolicyType", Required: false},
}

var fields_list_streaming_distributions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
}

var fields_list_trust_stores = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_vpc_origins = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_publish_connection_function = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_publish_function = []leanruntime.Field{
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "*types.Tags", Required: true},
}

var fields_test_connection_function = []leanruntime.Field{
	{Name: "ConnectionObject", Flag: "connection-object", Type: "[]byte", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_test_function = []leanruntime.Field{
	{Name: "EventObject", Flag: "event-object", Type: "[]byte", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Stage", Flag: "stage", Type: "types.FunctionStage", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "*types.TagKeys", Required: true},
}

var fields_update_anycast_ip_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
}

var fields_update_cache_policy = []leanruntime.Field{
	{Name: "CachePolicyConfig", Flag: "cache-policy-config", Type: "*types.CachePolicyConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_update_cloud_front_origin_access_identity = []leanruntime.Field{
	{Name: "CloudFrontOriginAccessIdentityConfig", Flag: "cloud-front-origin-access-identity-config", Type: "*types.CloudFrontOriginAccessIdentityConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_update_connection_function = []leanruntime.Field{
	{Name: "ConnectionFunctionCode", Flag: "connection-function-code", Type: "[]byte", Required: true},
	{Name: "ConnectionFunctionConfig", Flag: "connection-function-config", Type: "*types.FunctionConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_update_connection_group = []leanruntime.Field{
	{Name: "AnycastIpListId", Flag: "anycast-ip-list-id", Type: "*string", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Ipv6Enabled", Flag: "ipv6-enabled", Type: "*bool", Required: false},
}

var fields_update_continuous_deployment_policy = []leanruntime.Field{
	{Name: "ContinuousDeploymentPolicyConfig", Flag: "continuous-deployment-policy-config", Type: "*types.ContinuousDeploymentPolicyConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_update_distribution = []leanruntime.Field{
	{Name: "DistributionConfig", Flag: "distribution-config", Type: "*types.DistributionConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_update_distribution_tenant = []leanruntime.Field{
	{Name: "ConnectionGroupId", Flag: "connection-group-id", Type: "*string", Required: false},
	{Name: "Customizations", Flag: "customizations", Type: "*types.Customizations", Required: false},
	{Name: "DistributionId", Flag: "distribution-id", Type: "*string", Required: false},
	{Name: "Domains", Flag: "domains", Type: "[]types.DomainItem", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "ManagedCertificateRequest", Flag: "managed-certificate-request", Type: "*types.ManagedCertificateRequest", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
}

var fields_update_distribution_with_staging_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "StagingDistributionId", Flag: "staging-distribution-id", Type: "*string", Required: false},
}

var fields_update_domain_association = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "TargetResource", Flag: "target-resource", Type: "*types.DistributionResourceId", Required: true},
}

var fields_update_field_level_encryption_config = []leanruntime.Field{
	{Name: "FieldLevelEncryptionConfig", Flag: "field-level-encryption-config", Type: "*types.FieldLevelEncryptionConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_update_field_level_encryption_profile = []leanruntime.Field{
	{Name: "FieldLevelEncryptionProfileConfig", Flag: "field-level-encryption-profile-config", Type: "*types.FieldLevelEncryptionProfileConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
}

var fields_update_function = []leanruntime.Field{
	{Name: "FunctionCode", Flag: "function-code", Type: "[]byte", Required: true},
	{Name: "FunctionConfig", Flag: "function-config", Type: "*types.FunctionConfig", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_key_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "KeyGroupConfig", Flag: "key-group-config", Type: "*types.KeyGroupConfig", Required: true},
}

var fields_update_key_value_store = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_origin_access_control = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "OriginAccessControlConfig", Flag: "origin-access-control-config", Type: "*types.OriginAccessControlConfig", Required: true},
}

var fields_update_origin_request_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "OriginRequestPolicyConfig", Flag: "origin-request-policy-config", Type: "*types.OriginRequestPolicyConfig", Required: true},
}

var fields_update_public_key = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "PublicKeyConfig", Flag: "public-key-config", Type: "*types.PublicKeyConfig", Required: true},
}

var fields_update_realtime_log_config = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: false},
	{Name: "EndPoints", Flag: "end-points", Type: "[]types.EndPoint", Required: false},
	{Name: "Fields", Flag: "fields", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SamplingRate", Flag: "sampling-rate", Type: "*int64", Required: false},
}

var fields_update_response_headers_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "ResponseHeadersPolicyConfig", Flag: "response-headers-policy-config", Type: "*types.ResponseHeadersPolicyConfig", Required: true},
}

var fields_update_streaming_distribution = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "StreamingDistributionConfig", Flag: "streaming-distribution-config", Type: "*types.StreamingDistributionConfig", Required: true},
}

var fields_update_trust_store = []leanruntime.Field{
	{Name: "CaCertificatesBundleSource", Flag: "ca-certificates-bundle-source", Type: "types.CaCertificatesBundleSource", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
}

var fields_update_vpc_origin = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "VpcOriginEndpointConfig", Flag: "vpc-origin-endpoint-config", Type: "*types.VpcOriginEndpointConfig", Required: true},
}

var fields_verify_dns_configuration = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-alias": {
			Name:   "associate-alias",
			Fields: fields_associate_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAlias(ctx, input)
			},
		},
		"associate-distribution-tenant-web-acl": {
			Name:   "associate-distribution-tenant-web-acl",
			Fields: fields_associate_distribution_tenant_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDistributionTenantWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_distribution_tenant_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDistributionTenantWebACL(ctx, input)
			},
		},
		"associate-distribution-web-acl": {
			Name:   "associate-distribution-web-acl",
			Fields: fields_associate_distribution_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDistributionWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_distribution_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDistributionWebACL(ctx, input)
			},
		},
		"copy-distribution": {
			Name:   "copy-distribution",
			Fields: fields_copy_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyDistribution(ctx, input)
			},
		},
		"create-anycast-ip-list": {
			Name:   "create-anycast-ip-list",
			Fields: fields_create_anycast_ip_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnycastIpListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_anycast_ip_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnycastIpList(ctx, input)
			},
		},
		"create-cache-policy": {
			Name:   "create-cache-policy",
			Fields: fields_create_cache_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCachePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cache_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCachePolicy(ctx, input)
			},
		},
		"create-cloud-front-origin-access-identity": {
			Name:   "create-cloud-front-origin-access-identity",
			Fields: fields_create_cloud_front_origin_access_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudFrontOriginAccessIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_front_origin_access_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudFrontOriginAccessIdentity(ctx, input)
			},
		},
		"create-connection-function": {
			Name:   "create-connection-function",
			Fields: fields_create_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectionFunction(ctx, input)
			},
		},
		"create-connection-group": {
			Name:   "create-connection-group",
			Fields: fields_create_connection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectionGroup(ctx, input)
			},
		},
		"create-continuous-deployment-policy": {
			Name:   "create-continuous-deployment-policy",
			Fields: fields_create_continuous_deployment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContinuousDeploymentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_continuous_deployment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContinuousDeploymentPolicy(ctx, input)
			},
		},
		"create-distribution": {
			Name:   "create-distribution",
			Fields: fields_create_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDistribution(ctx, input)
			},
		},
		"create-distribution-tenant": {
			Name:   "create-distribution-tenant",
			Fields: fields_create_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDistributionTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_distribution_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDistributionTenant(ctx, input)
			},
		},
		"create-distribution-with-tags": {
			Name:   "create-distribution-with-tags",
			Fields: fields_create_distribution_with_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDistributionWithTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_distribution_with_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDistributionWithTags(ctx, input)
			},
		},
		"create-field-level-encryption-config": {
			Name:   "create-field-level-encryption-config",
			Fields: fields_create_field_level_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFieldLevelEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_field_level_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFieldLevelEncryptionConfig(ctx, input)
			},
		},
		"create-field-level-encryption-profile": {
			Name:   "create-field-level-encryption-profile",
			Fields: fields_create_field_level_encryption_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFieldLevelEncryptionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_field_level_encryption_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFieldLevelEncryptionProfile(ctx, input)
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
		"create-invalidation": {
			Name:   "create-invalidation",
			Fields: fields_create_invalidation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInvalidationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_invalidation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInvalidation(ctx, input)
			},
		},
		"create-invalidation-for-distribution-tenant": {
			Name:   "create-invalidation-for-distribution-tenant",
			Fields: fields_create_invalidation_for_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInvalidationForDistributionTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_invalidation_for_distribution_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInvalidationForDistributionTenant(ctx, input)
			},
		},
		"create-key-group": {
			Name:   "create-key-group",
			Fields: fields_create_key_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeyGroup(ctx, input)
			},
		},
		"create-key-value-store": {
			Name:   "create-key-value-store",
			Fields: fields_create_key_value_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyValueStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key_value_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeyValueStore(ctx, input)
			},
		},
		"create-monitoring-subscription": {
			Name:   "create-monitoring-subscription",
			Fields: fields_create_monitoring_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMonitoringSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_monitoring_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMonitoringSubscription(ctx, input)
			},
		},
		"create-origin-access-control": {
			Name:   "create-origin-access-control",
			Fields: fields_create_origin_access_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOriginAccessControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_origin_access_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOriginAccessControl(ctx, input)
			},
		},
		"create-origin-request-policy": {
			Name:   "create-origin-request-policy",
			Fields: fields_create_origin_request_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOriginRequestPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_origin_request_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOriginRequestPolicy(ctx, input)
			},
		},
		"create-public-key": {
			Name:   "create-public-key",
			Fields: fields_create_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePublicKey(ctx, input)
			},
		},
		"create-realtime-log-config": {
			Name:   "create-realtime-log-config",
			Fields: fields_create_realtime_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRealtimeLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_realtime_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRealtimeLogConfig(ctx, input)
			},
		},
		"create-response-headers-policy": {
			Name:   "create-response-headers-policy",
			Fields: fields_create_response_headers_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResponseHeadersPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_response_headers_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResponseHeadersPolicy(ctx, input)
			},
		},
		"create-streaming-distribution": {
			Name:   "create-streaming-distribution",
			Fields: fields_create_streaming_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamingDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_streaming_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamingDistribution(ctx, input)
			},
		},
		"create-streaming-distribution-with-tags": {
			Name:   "create-streaming-distribution-with-tags",
			Fields: fields_create_streaming_distribution_with_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamingDistributionWithTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_streaming_distribution_with_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamingDistributionWithTags(ctx, input)
			},
		},
		"create-trust-store": {
			Name:   "create-trust-store",
			Fields: fields_create_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrustStore(ctx, input)
			},
		},
		"create-vpc-origin": {
			Name:   "create-vpc-origin",
			Fields: fields_create_vpc_origin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcOriginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_origin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcOrigin(ctx, input)
			},
		},
		"delete-anycast-ip-list": {
			Name:   "delete-anycast-ip-list",
			Fields: fields_delete_anycast_ip_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnycastIpListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_anycast_ip_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnycastIpList(ctx, input)
			},
		},
		"delete-cache-policy": {
			Name:   "delete-cache-policy",
			Fields: fields_delete_cache_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCachePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cache_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCachePolicy(ctx, input)
			},
		},
		"delete-cloud-front-origin-access-identity": {
			Name:   "delete-cloud-front-origin-access-identity",
			Fields: fields_delete_cloud_front_origin_access_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudFrontOriginAccessIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_front_origin_access_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudFrontOriginAccessIdentity(ctx, input)
			},
		},
		"delete-connection-function": {
			Name:   "delete-connection-function",
			Fields: fields_delete_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectionFunction(ctx, input)
			},
		},
		"delete-connection-group": {
			Name:   "delete-connection-group",
			Fields: fields_delete_connection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectionGroup(ctx, input)
			},
		},
		"delete-continuous-deployment-policy": {
			Name:   "delete-continuous-deployment-policy",
			Fields: fields_delete_continuous_deployment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContinuousDeploymentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_continuous_deployment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContinuousDeploymentPolicy(ctx, input)
			},
		},
		"delete-distribution": {
			Name:   "delete-distribution",
			Fields: fields_delete_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDistribution(ctx, input)
			},
		},
		"delete-distribution-tenant": {
			Name:   "delete-distribution-tenant",
			Fields: fields_delete_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDistributionTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_distribution_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDistributionTenant(ctx, input)
			},
		},
		"delete-field-level-encryption-config": {
			Name:   "delete-field-level-encryption-config",
			Fields: fields_delete_field_level_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFieldLevelEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_field_level_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFieldLevelEncryptionConfig(ctx, input)
			},
		},
		"delete-field-level-encryption-profile": {
			Name:   "delete-field-level-encryption-profile",
			Fields: fields_delete_field_level_encryption_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFieldLevelEncryptionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_field_level_encryption_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFieldLevelEncryptionProfile(ctx, input)
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
		"delete-key-group": {
			Name:   "delete-key-group",
			Fields: fields_delete_key_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeyGroup(ctx, input)
			},
		},
		"delete-key-value-store": {
			Name:   "delete-key-value-store",
			Fields: fields_delete_key_value_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyValueStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key_value_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeyValueStore(ctx, input)
			},
		},
		"delete-monitoring-subscription": {
			Name:   "delete-monitoring-subscription",
			Fields: fields_delete_monitoring_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMonitoringSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_monitoring_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMonitoringSubscription(ctx, input)
			},
		},
		"delete-origin-access-control": {
			Name:   "delete-origin-access-control",
			Fields: fields_delete_origin_access_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOriginAccessControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_origin_access_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOriginAccessControl(ctx, input)
			},
		},
		"delete-origin-request-policy": {
			Name:   "delete-origin-request-policy",
			Fields: fields_delete_origin_request_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOriginRequestPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_origin_request_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOriginRequestPolicy(ctx, input)
			},
		},
		"delete-public-key": {
			Name:   "delete-public-key",
			Fields: fields_delete_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePublicKey(ctx, input)
			},
		},
		"delete-realtime-log-config": {
			Name:   "delete-realtime-log-config",
			Fields: fields_delete_realtime_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRealtimeLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_realtime_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRealtimeLogConfig(ctx, input)
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
		"delete-response-headers-policy": {
			Name:   "delete-response-headers-policy",
			Fields: fields_delete_response_headers_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResponseHeadersPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_response_headers_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResponseHeadersPolicy(ctx, input)
			},
		},
		"delete-streaming-distribution": {
			Name:   "delete-streaming-distribution",
			Fields: fields_delete_streaming_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamingDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_streaming_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStreamingDistribution(ctx, input)
			},
		},
		"delete-trust-store": {
			Name:   "delete-trust-store",
			Fields: fields_delete_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrustStore(ctx, input)
			},
		},
		"delete-vpc-origin": {
			Name:   "delete-vpc-origin",
			Fields: fields_delete_vpc_origin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcOriginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_origin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcOrigin(ctx, input)
			},
		},
		"describe-connection-function": {
			Name:   "describe-connection-function",
			Fields: fields_describe_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectionFunction(ctx, input)
			},
		},
		"describe-function": {
			Name:   "describe-function",
			Fields: fields_describe_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFunction(ctx, input)
			},
		},
		"describe-key-value-store": {
			Name:   "describe-key-value-store",
			Fields: fields_describe_key_value_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyValueStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key_value_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKeyValueStore(ctx, input)
			},
		},
		"disassociate-distribution-tenant-web-acl": {
			Name:   "disassociate-distribution-tenant-web-acl",
			Fields: fields_disassociate_distribution_tenant_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDistributionTenantWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_distribution_tenant_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDistributionTenantWebACL(ctx, input)
			},
		},
		"disassociate-distribution-web-acl": {
			Name:   "disassociate-distribution-web-acl",
			Fields: fields_disassociate_distribution_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDistributionWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_distribution_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDistributionWebACL(ctx, input)
			},
		},
		"get-anycast-ip-list": {
			Name:   "get-anycast-ip-list",
			Fields: fields_get_anycast_ip_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnycastIpListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_anycast_ip_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnycastIpList(ctx, input)
			},
		},
		"get-cache-policy": {
			Name:   "get-cache-policy",
			Fields: fields_get_cache_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCachePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cache_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCachePolicy(ctx, input)
			},
		},
		"get-cache-policy-config": {
			Name:   "get-cache-policy-config",
			Fields: fields_get_cache_policy_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCachePolicyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cache_policy_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCachePolicyConfig(ctx, input)
			},
		},
		"get-cloud-front-origin-access-identity": {
			Name:   "get-cloud-front-origin-access-identity",
			Fields: fields_get_cloud_front_origin_access_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudFrontOriginAccessIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_front_origin_access_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudFrontOriginAccessIdentity(ctx, input)
			},
		},
		"get-cloud-front-origin-access-identity-config": {
			Name:   "get-cloud-front-origin-access-identity-config",
			Fields: fields_get_cloud_front_origin_access_identity_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudFrontOriginAccessIdentityConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_front_origin_access_identity_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudFrontOriginAccessIdentityConfig(ctx, input)
			},
		},
		"get-connection-function": {
			Name:   "get-connection-function",
			Fields: fields_get_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionFunction(ctx, input)
			},
		},
		"get-connection-group": {
			Name:   "get-connection-group",
			Fields: fields_get_connection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionGroup(ctx, input)
			},
		},
		"get-connection-group-by-routing-endpoint": {
			Name:   "get-connection-group-by-routing-endpoint",
			Fields: fields_get_connection_group_by_routing_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionGroupByRoutingEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_group_by_routing_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionGroupByRoutingEndpoint(ctx, input)
			},
		},
		"get-continuous-deployment-policy": {
			Name:   "get-continuous-deployment-policy",
			Fields: fields_get_continuous_deployment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContinuousDeploymentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_continuous_deployment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContinuousDeploymentPolicy(ctx, input)
			},
		},
		"get-continuous-deployment-policy-config": {
			Name:   "get-continuous-deployment-policy-config",
			Fields: fields_get_continuous_deployment_policy_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContinuousDeploymentPolicyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_continuous_deployment_policy_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContinuousDeploymentPolicyConfig(ctx, input)
			},
		},
		"get-distribution": {
			Name:   "get-distribution",
			Fields: fields_get_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistribution(ctx, input)
			},
		},
		"get-distribution-config": {
			Name:   "get-distribution-config",
			Fields: fields_get_distribution_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionConfig(ctx, input)
			},
		},
		"get-distribution-tenant": {
			Name:   "get-distribution-tenant",
			Fields: fields_get_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionTenant(ctx, input)
			},
		},
		"get-distribution-tenant-by-domain": {
			Name:   "get-distribution-tenant-by-domain",
			Fields: fields_get_distribution_tenant_by_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionTenantByDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_tenant_by_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionTenantByDomain(ctx, input)
			},
		},
		"get-field-level-encryption": {
			Name:   "get-field-level-encryption",
			Fields: fields_get_field_level_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFieldLevelEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_field_level_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFieldLevelEncryption(ctx, input)
			},
		},
		"get-field-level-encryption-config": {
			Name:   "get-field-level-encryption-config",
			Fields: fields_get_field_level_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFieldLevelEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_field_level_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFieldLevelEncryptionConfig(ctx, input)
			},
		},
		"get-field-level-encryption-profile": {
			Name:   "get-field-level-encryption-profile",
			Fields: fields_get_field_level_encryption_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFieldLevelEncryptionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_field_level_encryption_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFieldLevelEncryptionProfile(ctx, input)
			},
		},
		"get-field-level-encryption-profile-config": {
			Name:   "get-field-level-encryption-profile-config",
			Fields: fields_get_field_level_encryption_profile_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFieldLevelEncryptionProfileConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_field_level_encryption_profile_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFieldLevelEncryptionProfileConfig(ctx, input)
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
		"get-invalidation": {
			Name:   "get-invalidation",
			Fields: fields_get_invalidation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvalidationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_invalidation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvalidation(ctx, input)
			},
		},
		"get-invalidation-for-distribution-tenant": {
			Name:   "get-invalidation-for-distribution-tenant",
			Fields: fields_get_invalidation_for_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvalidationForDistributionTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_invalidation_for_distribution_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvalidationForDistributionTenant(ctx, input)
			},
		},
		"get-key-group": {
			Name:   "get-key-group",
			Fields: fields_get_key_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyGroup(ctx, input)
			},
		},
		"get-key-group-config": {
			Name:   "get-key-group-config",
			Fields: fields_get_key_group_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyGroupConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key_group_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyGroupConfig(ctx, input)
			},
		},
		"get-managed-certificate-details": {
			Name:   "get-managed-certificate-details",
			Fields: fields_get_managed_certificate_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedCertificateDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_certificate_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedCertificateDetails(ctx, input)
			},
		},
		"get-monitoring-subscription": {
			Name:   "get-monitoring-subscription",
			Fields: fields_get_monitoring_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMonitoringSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_monitoring_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMonitoringSubscription(ctx, input)
			},
		},
		"get-origin-access-control": {
			Name:   "get-origin-access-control",
			Fields: fields_get_origin_access_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOriginAccessControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_origin_access_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOriginAccessControl(ctx, input)
			},
		},
		"get-origin-access-control-config": {
			Name:   "get-origin-access-control-config",
			Fields: fields_get_origin_access_control_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOriginAccessControlConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_origin_access_control_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOriginAccessControlConfig(ctx, input)
			},
		},
		"get-origin-request-policy": {
			Name:   "get-origin-request-policy",
			Fields: fields_get_origin_request_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOriginRequestPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_origin_request_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOriginRequestPolicy(ctx, input)
			},
		},
		"get-origin-request-policy-config": {
			Name:   "get-origin-request-policy-config",
			Fields: fields_get_origin_request_policy_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOriginRequestPolicyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_origin_request_policy_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOriginRequestPolicyConfig(ctx, input)
			},
		},
		"get-public-key": {
			Name:   "get-public-key",
			Fields: fields_get_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPublicKey(ctx, input)
			},
		},
		"get-public-key-config": {
			Name:   "get-public-key-config",
			Fields: fields_get_public_key_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPublicKeyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_public_key_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPublicKeyConfig(ctx, input)
			},
		},
		"get-realtime-log-config": {
			Name:   "get-realtime-log-config",
			Fields: fields_get_realtime_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRealtimeLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_realtime_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRealtimeLogConfig(ctx, input)
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
		"get-response-headers-policy": {
			Name:   "get-response-headers-policy",
			Fields: fields_get_response_headers_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResponseHeadersPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_response_headers_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResponseHeadersPolicy(ctx, input)
			},
		},
		"get-response-headers-policy-config": {
			Name:   "get-response-headers-policy-config",
			Fields: fields_get_response_headers_policy_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResponseHeadersPolicyConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_response_headers_policy_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResponseHeadersPolicyConfig(ctx, input)
			},
		},
		"get-streaming-distribution": {
			Name:   "get-streaming-distribution",
			Fields: fields_get_streaming_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamingDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_streaming_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStreamingDistribution(ctx, input)
			},
		},
		"get-streaming-distribution-config": {
			Name:   "get-streaming-distribution-config",
			Fields: fields_get_streaming_distribution_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamingDistributionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_streaming_distribution_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStreamingDistributionConfig(ctx, input)
			},
		},
		"get-trust-store": {
			Name:   "get-trust-store",
			Fields: fields_get_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustStore(ctx, input)
			},
		},
		"get-vpc-origin": {
			Name:   "get-vpc-origin",
			Fields: fields_get_vpc_origin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpcOriginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpc_origin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpcOrigin(ctx, input)
			},
		},
		"list-anycast-ip-lists": {
			Name:   "list-anycast-ip-lists",
			Fields: fields_list_anycast_ip_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnycastIpListsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_anycast_ip_lists, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAnycastIpLists(ctx, input)
			},
		},
		"list-cache-policies": {
			Name:   "list-cache-policies",
			Fields: fields_list_cache_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCachePoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_cache_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCachePolicies(ctx, input)
			},
		},
		"list-cloud-front-origin-access-identities": {
			Name:   "list-cloud-front-origin-access-identities",
			Fields: fields_list_cloud_front_origin_access_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudFrontOriginAccessIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_front_origin_access_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudFrontOriginAccessIdentities(ctx, input)
				}
				var results []*svc.ListCloudFrontOriginAccessIdentitiesOutput
				p := svc.NewListCloudFrontOriginAccessIdentitiesPaginator(client, input)
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
		"list-conflicting-aliases": {
			Name:   "list-conflicting-aliases",
			Fields: fields_list_conflicting_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConflictingAliasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_conflicting_aliases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConflictingAliases(ctx, input)
			},
		},
		"list-connection-functions": {
			Name:   "list-connection-functions",
			Fields: fields_list_connection_functions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionFunctionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connection_functions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectionFunctions(ctx, input)
				}
				var results []*svc.ListConnectionFunctionsOutput
				p := svc.NewListConnectionFunctionsPaginator(client, input)
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
		"list-connection-groups": {
			Name:   "list-connection-groups",
			Fields: fields_list_connection_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connection_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectionGroups(ctx, input)
				}
				var results []*svc.ListConnectionGroupsOutput
				p := svc.NewListConnectionGroupsPaginator(client, input)
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
		"list-continuous-deployment-policies": {
			Name:   "list-continuous-deployment-policies",
			Fields: fields_list_continuous_deployment_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContinuousDeploymentPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_continuous_deployment_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListContinuousDeploymentPolicies(ctx, input)
			},
		},
		"list-distribution-tenants": {
			Name:   "list-distribution-tenants",
			Fields: fields_list_distribution_tenants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionTenantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distribution_tenants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributionTenants(ctx, input)
				}
				var results []*svc.ListDistributionTenantsOutput
				p := svc.NewListDistributionTenantsPaginator(client, input)
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
		"list-distribution-tenants-by-customization": {
			Name:   "list-distribution-tenants-by-customization",
			Fields: fields_list_distribution_tenants_by_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionTenantsByCustomizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distribution_tenants_by_customization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributionTenantsByCustomization(ctx, input)
				}
				var results []*svc.ListDistributionTenantsByCustomizationOutput
				p := svc.NewListDistributionTenantsByCustomizationPaginator(client, input)
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
		"list-distributions": {
			Name:   "list-distributions",
			Fields: fields_list_distributions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distributions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributions(ctx, input)
				}
				var results []*svc.ListDistributionsOutput
				p := svc.NewListDistributionsPaginator(client, input)
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
		"list-distributions-by-anycast-ip-list-id": {
			Name:   "list-distributions-by-anycast-ip-list-id",
			Fields: fields_list_distributions_by_anycast_ip_list_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByAnycastIpListIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_anycast_ip_list_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByAnycastIpListId(ctx, input)
			},
		},
		"list-distributions-by-cache-policy-id": {
			Name:   "list-distributions-by-cache-policy-id",
			Fields: fields_list_distributions_by_cache_policy_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByCachePolicyIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_cache_policy_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByCachePolicyId(ctx, input)
			},
		},
		"list-distributions-by-connection-function": {
			Name:   "list-distributions-by-connection-function",
			Fields: fields_list_distributions_by_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByConnectionFunctionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distributions_by_connection_function, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributionsByConnectionFunction(ctx, input)
				}
				var results []*svc.ListDistributionsByConnectionFunctionOutput
				p := svc.NewListDistributionsByConnectionFunctionPaginator(client, input)
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
		"list-distributions-by-connection-mode": {
			Name:   "list-distributions-by-connection-mode",
			Fields: fields_list_distributions_by_connection_mode,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByConnectionModeInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distributions_by_connection_mode, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributionsByConnectionMode(ctx, input)
				}
				var results []*svc.ListDistributionsByConnectionModeOutput
				p := svc.NewListDistributionsByConnectionModePaginator(client, input)
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
		"list-distributions-by-key-group": {
			Name:   "list-distributions-by-key-group",
			Fields: fields_list_distributions_by_key_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByKeyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_key_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByKeyGroup(ctx, input)
			},
		},
		"list-distributions-by-origin-request-policy-id": {
			Name:   "list-distributions-by-origin-request-policy-id",
			Fields: fields_list_distributions_by_origin_request_policy_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByOriginRequestPolicyIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_origin_request_policy_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByOriginRequestPolicyId(ctx, input)
			},
		},
		"list-distributions-by-owned-resource": {
			Name:   "list-distributions-by-owned-resource",
			Fields: fields_list_distributions_by_owned_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByOwnedResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_owned_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByOwnedResource(ctx, input)
			},
		},
		"list-distributions-by-realtime-log-config": {
			Name:   "list-distributions-by-realtime-log-config",
			Fields: fields_list_distributions_by_realtime_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByRealtimeLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_realtime_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByRealtimeLogConfig(ctx, input)
			},
		},
		"list-distributions-by-response-headers-policy-id": {
			Name:   "list-distributions-by-response-headers-policy-id",
			Fields: fields_list_distributions_by_response_headers_policy_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByResponseHeadersPolicyIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_response_headers_policy_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByResponseHeadersPolicyId(ctx, input)
			},
		},
		"list-distributions-by-trust-store": {
			Name:   "list-distributions-by-trust-store",
			Fields: fields_list_distributions_by_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByTrustStoreInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distributions_by_trust_store, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributionsByTrustStore(ctx, input)
				}
				var results []*svc.ListDistributionsByTrustStoreOutput
				p := svc.NewListDistributionsByTrustStorePaginator(client, input)
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
		"list-distributions-by-vpc-origin-id": {
			Name:   "list-distributions-by-vpc-origin-id",
			Fields: fields_list_distributions_by_vpc_origin_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByVpcOriginIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_vpc_origin_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByVpcOriginId(ctx, input)
			},
		},
		"list-distributions-by-web-aclid": {
			Name:   "list-distributions-by-web-aclid",
			Fields: fields_list_distributions_by_web_aclid,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionsByWebACLIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributions_by_web_aclid, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributionsByWebACLId(ctx, input)
			},
		},
		"list-domain-conflicts": {
			Name:   "list-domain-conflicts",
			Fields: fields_list_domain_conflicts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainConflictsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_conflicts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainConflicts(ctx, input)
				}
				var results []*svc.ListDomainConflictsOutput
				p := svc.NewListDomainConflictsPaginator(client, input)
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
		"list-field-level-encryption-configs": {
			Name:   "list-field-level-encryption-configs",
			Fields: fields_list_field_level_encryption_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFieldLevelEncryptionConfigsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_field_level_encryption_configs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFieldLevelEncryptionConfigs(ctx, input)
			},
		},
		"list-field-level-encryption-profiles": {
			Name:   "list-field-level-encryption-profiles",
			Fields: fields_list_field_level_encryption_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFieldLevelEncryptionProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_field_level_encryption_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFieldLevelEncryptionProfiles(ctx, input)
			},
		},
		"list-functions": {
			Name:   "list-functions",
			Fields: fields_list_functions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_functions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFunctions(ctx, input)
			},
		},
		"list-invalidations": {
			Name:   "list-invalidations",
			Fields: fields_list_invalidations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvalidationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invalidations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvalidations(ctx, input)
				}
				var results []*svc.ListInvalidationsOutput
				p := svc.NewListInvalidationsPaginator(client, input)
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
		"list-invalidations-for-distribution-tenant": {
			Name:   "list-invalidations-for-distribution-tenant",
			Fields: fields_list_invalidations_for_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvalidationsForDistributionTenantInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invalidations_for_distribution_tenant, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvalidationsForDistributionTenant(ctx, input)
				}
				var results []*svc.ListInvalidationsForDistributionTenantOutput
				p := svc.NewListInvalidationsForDistributionTenantPaginator(client, input)
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
		"list-key-groups": {
			Name:   "list-key-groups",
			Fields: fields_list_key_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeyGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_key_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListKeyGroups(ctx, input)
			},
		},
		"list-key-value-stores": {
			Name:   "list-key-value-stores",
			Fields: fields_list_key_value_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeyValueStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_key_value_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeyValueStores(ctx, input)
				}
				var results []*svc.ListKeyValueStoresOutput
				p := svc.NewListKeyValueStoresPaginator(client, input)
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
		"list-origin-access-controls": {
			Name:   "list-origin-access-controls",
			Fields: fields_list_origin_access_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOriginAccessControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_origin_access_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOriginAccessControls(ctx, input)
				}
				var results []*svc.ListOriginAccessControlsOutput
				p := svc.NewListOriginAccessControlsPaginator(client, input)
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
		"list-origin-request-policies": {
			Name:   "list-origin-request-policies",
			Fields: fields_list_origin_request_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOriginRequestPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_origin_request_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOriginRequestPolicies(ctx, input)
			},
		},
		"list-public-keys": {
			Name:   "list-public-keys",
			Fields: fields_list_public_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPublicKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_public_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPublicKeys(ctx, input)
				}
				var results []*svc.ListPublicKeysOutput
				p := svc.NewListPublicKeysPaginator(client, input)
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
		"list-realtime-log-configs": {
			Name:   "list-realtime-log-configs",
			Fields: fields_list_realtime_log_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRealtimeLogConfigsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_realtime_log_configs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRealtimeLogConfigs(ctx, input)
			},
		},
		"list-response-headers-policies": {
			Name:   "list-response-headers-policies",
			Fields: fields_list_response_headers_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResponseHeadersPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_response_headers_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResponseHeadersPolicies(ctx, input)
			},
		},
		"list-streaming-distributions": {
			Name:   "list-streaming-distributions",
			Fields: fields_list_streaming_distributions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamingDistributionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_streaming_distributions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamingDistributions(ctx, input)
				}
				var results []*svc.ListStreamingDistributionsOutput
				p := svc.NewListStreamingDistributionsPaginator(client, input)
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
		"list-trust-stores": {
			Name:   "list-trust-stores",
			Fields: fields_list_trust_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrustStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trust_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrustStores(ctx, input)
				}
				var results []*svc.ListTrustStoresOutput
				p := svc.NewListTrustStoresPaginator(client, input)
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
		"list-vpc-origins": {
			Name:   "list-vpc-origins",
			Fields: fields_list_vpc_origins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcOriginsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_origins, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcOrigins(ctx, input)
			},
		},
		"publish-connection-function": {
			Name:   "publish-connection-function",
			Fields: fields_publish_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishConnectionFunction(ctx, input)
			},
		},
		"publish-function": {
			Name:   "publish-function",
			Fields: fields_publish_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishFunction(ctx, input)
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
		"test-connection-function": {
			Name:   "test-connection-function",
			Fields: fields_test_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestConnectionFunction(ctx, input)
			},
		},
		"test-function": {
			Name:   "test-function",
			Fields: fields_test_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestFunction(ctx, input)
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
		"update-anycast-ip-list": {
			Name:   "update-anycast-ip-list",
			Fields: fields_update_anycast_ip_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnycastIpListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_anycast_ip_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnycastIpList(ctx, input)
			},
		},
		"update-cache-policy": {
			Name:   "update-cache-policy",
			Fields: fields_update_cache_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCachePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cache_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCachePolicy(ctx, input)
			},
		},
		"update-cloud-front-origin-access-identity": {
			Name:   "update-cloud-front-origin-access-identity",
			Fields: fields_update_cloud_front_origin_access_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCloudFrontOriginAccessIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cloud_front_origin_access_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCloudFrontOriginAccessIdentity(ctx, input)
			},
		},
		"update-connection-function": {
			Name:   "update-connection-function",
			Fields: fields_update_connection_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectionFunction(ctx, input)
			},
		},
		"update-connection-group": {
			Name:   "update-connection-group",
			Fields: fields_update_connection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectionGroup(ctx, input)
			},
		},
		"update-continuous-deployment-policy": {
			Name:   "update-continuous-deployment-policy",
			Fields: fields_update_continuous_deployment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContinuousDeploymentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_continuous_deployment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContinuousDeploymentPolicy(ctx, input)
			},
		},
		"update-distribution": {
			Name:   "update-distribution",
			Fields: fields_update_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDistribution(ctx, input)
			},
		},
		"update-distribution-tenant": {
			Name:   "update-distribution-tenant",
			Fields: fields_update_distribution_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDistributionTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_distribution_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDistributionTenant(ctx, input)
			},
		},
		"update-distribution-with-staging-config": {
			Name:   "update-distribution-with-staging-config",
			Fields: fields_update_distribution_with_staging_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDistributionWithStagingConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_distribution_with_staging_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDistributionWithStagingConfig(ctx, input)
			},
		},
		"update-domain-association": {
			Name:   "update-domain-association",
			Fields: fields_update_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainAssociation(ctx, input)
			},
		},
		"update-field-level-encryption-config": {
			Name:   "update-field-level-encryption-config",
			Fields: fields_update_field_level_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFieldLevelEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_field_level_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFieldLevelEncryptionConfig(ctx, input)
			},
		},
		"update-field-level-encryption-profile": {
			Name:   "update-field-level-encryption-profile",
			Fields: fields_update_field_level_encryption_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFieldLevelEncryptionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_field_level_encryption_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFieldLevelEncryptionProfile(ctx, input)
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
		"update-key-group": {
			Name:   "update-key-group",
			Fields: fields_update_key_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_key_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKeyGroup(ctx, input)
			},
		},
		"update-key-value-store": {
			Name:   "update-key-value-store",
			Fields: fields_update_key_value_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeyValueStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_key_value_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKeyValueStore(ctx, input)
			},
		},
		"update-origin-access-control": {
			Name:   "update-origin-access-control",
			Fields: fields_update_origin_access_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOriginAccessControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_origin_access_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOriginAccessControl(ctx, input)
			},
		},
		"update-origin-request-policy": {
			Name:   "update-origin-request-policy",
			Fields: fields_update_origin_request_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOriginRequestPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_origin_request_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOriginRequestPolicy(ctx, input)
			},
		},
		"update-public-key": {
			Name:   "update-public-key",
			Fields: fields_update_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePublicKey(ctx, input)
			},
		},
		"update-realtime-log-config": {
			Name:   "update-realtime-log-config",
			Fields: fields_update_realtime_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRealtimeLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_realtime_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRealtimeLogConfig(ctx, input)
			},
		},
		"update-response-headers-policy": {
			Name:   "update-response-headers-policy",
			Fields: fields_update_response_headers_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResponseHeadersPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_response_headers_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResponseHeadersPolicy(ctx, input)
			},
		},
		"update-streaming-distribution": {
			Name:   "update-streaming-distribution",
			Fields: fields_update_streaming_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamingDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_streaming_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStreamingDistribution(ctx, input)
			},
		},
		"update-trust-store": {
			Name:   "update-trust-store",
			Fields: fields_update_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrustStore(ctx, input)
			},
		},
		"update-vpc-origin": {
			Name:   "update-vpc-origin",
			Fields: fields_update_vpc_origin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcOriginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_origin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcOrigin(ctx, input)
			},
		},
		"verify-dns-configuration": {
			Name:   "verify-dns-configuration",
			Fields: fields_verify_dns_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyDnsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_dns_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyDnsConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudfront", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
