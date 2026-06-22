package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/vpclattice"
)

var fields_batch_update_rule = []leanruntime.Field{
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.RuleUpdate", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_create_access_log_subscription = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ServiceNetworkLogType", Flag: "service-network-log-type", Type: "types.ServiceNetworkLogType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_listener = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultAction", Flag: "default-action", Type: "types.RuleAction", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.ListenerProtocol", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_resource_configuration = []leanruntime.Field{
	{Name: "AllowAssociationToShareableServiceNetwork", Flag: "allow-association-to-shareable-service-network", Type: "*bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "DomainVerificationIdentifier", Flag: "domain-verification-identifier", Type: "*string", Required: false},
	{Name: "GroupDomain", Flag: "group-domain", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PortRanges", Flag: "port-ranges", Type: "[]string", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.ProtocolType", Required: false},
	{Name: "ResourceConfigurationDefinition", Flag: "resource-configuration-definition", Type: "types.ResourceConfigurationDefinition", Required: false},
	{Name: "ResourceConfigurationGroupIdentifier", Flag: "resource-configuration-group-identifier", Type: "*string", Required: false},
	{Name: "ResourceGatewayIdentifier", Flag: "resource-gateway-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ResourceConfigurationType", Required: true},
}

var fields_create_resource_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.ResourceGatewayIpAddressType", Required: false},
	{Name: "Ipv4AddressesPerEni", Flag: "ipv4-addresses-per-eni", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcIdentifier", Flag: "vpc-identifier", Type: "*string", Required: false},
}

var fields_create_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.RuleAction", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "Match", Flag: "match", Type: "types.RuleMatch", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: false},
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service_network = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SharingConfig", Flag: "sharing-config", Type: "*types.SharingConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service_network_resource_association = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PrivateDnsEnabled", Flag: "private-dns-enabled", Type: "*bool", Required: false},
	{Name: "ResourceConfigurationIdentifier", Flag: "resource-configuration-identifier", Type: "*string", Required: true},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service_network_service_association = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service_network_vpc_association = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DnsOptions", Flag: "dns-options", Type: "*types.DnsOptions", Required: false},
	{Name: "PrivateDnsEnabled", Flag: "private-dns-enabled", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcIdentifier", Flag: "vpc-identifier", Type: "*string", Required: true},
}

var fields_create_target_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Config", Flag: "config", Type: "*types.TargetGroupConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.TargetGroupType", Required: true},
}

var fields_delete_access_log_subscription = []leanruntime.Field{
	{Name: "AccessLogSubscriptionIdentifier", Flag: "access-log-subscription-identifier", Type: "*string", Required: true},
}

var fields_delete_auth_policy = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_delete_domain_verification = []leanruntime.Field{
	{Name: "DomainVerificationIdentifier", Flag: "domain-verification-identifier", Type: "*string", Required: true},
}

var fields_delete_listener = []leanruntime.Field{
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_delete_resource_configuration = []leanruntime.Field{
	{Name: "ResourceConfigurationIdentifier", Flag: "resource-configuration-identifier", Type: "*string", Required: true},
}

var fields_delete_resource_endpoint_association = []leanruntime.Field{
	{Name: "ResourceEndpointAssociationIdentifier", Flag: "resource-endpoint-association-identifier", Type: "*string", Required: true},
}

var fields_delete_resource_gateway = []leanruntime.Field{
	{Name: "ResourceGatewayIdentifier", Flag: "resource-gateway-identifier", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_delete_service = []leanruntime.Field{
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_delete_service_network = []leanruntime.Field{
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
}

var fields_delete_service_network_resource_association = []leanruntime.Field{
	{Name: "ServiceNetworkResourceAssociationIdentifier", Flag: "service-network-resource-association-identifier", Type: "*string", Required: true},
}

var fields_delete_service_network_service_association = []leanruntime.Field{
	{Name: "ServiceNetworkServiceAssociationIdentifier", Flag: "service-network-service-association-identifier", Type: "*string", Required: true},
}

var fields_delete_service_network_vpc_association = []leanruntime.Field{
	{Name: "ServiceNetworkVpcAssociationIdentifier", Flag: "service-network-vpc-association-identifier", Type: "*string", Required: true},
}

var fields_delete_target_group = []leanruntime.Field{
	{Name: "TargetGroupIdentifier", Flag: "target-group-identifier", Type: "*string", Required: true},
}

var fields_deregister_targets = []leanruntime.Field{
	{Name: "TargetGroupIdentifier", Flag: "target-group-identifier", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
}

var fields_get_access_log_subscription = []leanruntime.Field{
	{Name: "AccessLogSubscriptionIdentifier", Flag: "access-log-subscription-identifier", Type: "*string", Required: true},
}

var fields_get_auth_policy = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_get_domain_verification = []leanruntime.Field{
	{Name: "DomainVerificationIdentifier", Flag: "domain-verification-identifier", Type: "*string", Required: true},
}

var fields_get_listener = []leanruntime.Field{
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_get_resource_configuration = []leanruntime.Field{
	{Name: "ResourceConfigurationIdentifier", Flag: "resource-configuration-identifier", Type: "*string", Required: true},
}

var fields_get_resource_gateway = []leanruntime.Field{
	{Name: "ResourceGatewayIdentifier", Flag: "resource-gateway-identifier", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_rule = []leanruntime.Field{
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_get_service = []leanruntime.Field{
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_get_service_network = []leanruntime.Field{
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
}

var fields_get_service_network_resource_association = []leanruntime.Field{
	{Name: "ServiceNetworkResourceAssociationIdentifier", Flag: "service-network-resource-association-identifier", Type: "*string", Required: true},
}

var fields_get_service_network_service_association = []leanruntime.Field{
	{Name: "ServiceNetworkServiceAssociationIdentifier", Flag: "service-network-service-association-identifier", Type: "*string", Required: true},
}

var fields_get_service_network_vpc_association = []leanruntime.Field{
	{Name: "ServiceNetworkVpcAssociationIdentifier", Flag: "service-network-vpc-association-identifier", Type: "*string", Required: true},
}

var fields_get_target_group = []leanruntime.Field{
	{Name: "TargetGroupIdentifier", Flag: "target-group-identifier", Type: "*string", Required: true},
}

var fields_list_access_log_subscriptions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_list_domain_verifications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_listeners = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_list_resource_configurations = []leanruntime.Field{
	{Name: "DomainVerificationIdentifier", Flag: "domain-verification-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceConfigurationGroupIdentifier", Flag: "resource-configuration-group-identifier", Type: "*string", Required: false},
	{Name: "ResourceGatewayIdentifier", Flag: "resource-gateway-identifier", Type: "*string", Required: false},
}

var fields_list_resource_endpoint_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceConfigurationIdentifier", Flag: "resource-configuration-identifier", Type: "*string", Required: true},
	{Name: "ResourceEndpointAssociationIdentifier", Flag: "resource-endpoint-association-identifier", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
	{Name: "VpcEndpointOwner", Flag: "vpc-endpoint-owner", Type: "*string", Required: false},
}

var fields_list_resource_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rules = []leanruntime.Field{
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_list_service_network_resource_associations = []leanruntime.Field{
	{Name: "IncludeChildren", Flag: "include-children", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceConfigurationIdentifier", Flag: "resource-configuration-identifier", Type: "*string", Required: false},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: false},
}

var fields_list_service_network_service_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: false},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: false},
}

var fields_list_service_network_vpc_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: false},
	{Name: "VpcIdentifier", Flag: "vpc-identifier", Type: "*string", Required: false},
}

var fields_list_service_network_vpc_endpoint_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
}

var fields_list_service_networks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_services = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_target_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetGroupType", Flag: "target-group-type", Type: "types.TargetGroupType", Required: false},
	{Name: "VpcIdentifier", Flag: "vpc-identifier", Type: "*string", Required: false},
}

var fields_list_targets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetGroupIdentifier", Flag: "target-group-identifier", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
}

var fields_put_auth_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_targets = []leanruntime.Field{
	{Name: "TargetGroupIdentifier", Flag: "target-group-identifier", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
}

var fields_start_domain_verification = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_log_subscription = []leanruntime.Field{
	{Name: "AccessLogSubscriptionIdentifier", Flag: "access-log-subscription-identifier", Type: "*string", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: true},
}

var fields_update_listener = []leanruntime.Field{
	{Name: "DefaultAction", Flag: "default-action", Type: "types.RuleAction", Required: true},
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_update_resource_configuration = []leanruntime.Field{
	{Name: "AllowAssociationToShareableServiceNetwork", Flag: "allow-association-to-shareable-service-network", Type: "*bool", Required: false},
	{Name: "PortRanges", Flag: "port-ranges", Type: "[]string", Required: false},
	{Name: "ResourceConfigurationDefinition", Flag: "resource-configuration-definition", Type: "types.ResourceConfigurationDefinition", Required: false},
	{Name: "ResourceConfigurationIdentifier", Flag: "resource-configuration-identifier", Type: "*string", Required: true},
}

var fields_update_resource_gateway = []leanruntime.Field{
	{Name: "ResourceGatewayIdentifier", Flag: "resource-gateway-identifier", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
}

var fields_update_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.RuleAction", Required: false},
	{Name: "ListenerIdentifier", Flag: "listener-identifier", Type: "*string", Required: true},
	{Name: "Match", Flag: "match", Type: "types.RuleMatch", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_update_service = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: false},
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_update_service_network = []leanruntime.Field{
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: true},
	{Name: "ServiceNetworkIdentifier", Flag: "service-network-identifier", Type: "*string", Required: true},
}

var fields_update_service_network_vpc_association = []leanruntime.Field{
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "ServiceNetworkVpcAssociationIdentifier", Flag: "service-network-vpc-association-identifier", Type: "*string", Required: true},
}

var fields_update_target_group = []leanruntime.Field{
	{Name: "HealthCheck", Flag: "health-check", Type: "*types.HealthCheckConfig", Required: true},
	{Name: "TargetGroupIdentifier", Flag: "target-group-identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-update-rule": {
			Name:   "batch-update-rule",
			Fields: fields_batch_update_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateRule(ctx, input)
			},
		},
		"create-access-log-subscription": {
			Name:   "create-access-log-subscription",
			Fields: fields_create_access_log_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessLogSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_log_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessLogSubscription(ctx, input)
			},
		},
		"create-listener": {
			Name:   "create-listener",
			Fields: fields_create_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateListener(ctx, input)
			},
		},
		"create-resource-configuration": {
			Name:   "create-resource-configuration",
			Fields: fields_create_resource_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceConfiguration(ctx, input)
			},
		},
		"create-resource-gateway": {
			Name:   "create-resource-gateway",
			Fields: fields_create_resource_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceGateway(ctx, input)
			},
		},
		"create-rule": {
			Name:   "create-rule",
			Fields: fields_create_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRule(ctx, input)
			},
		},
		"create-service": {
			Name:   "create-service",
			Fields: fields_create_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateService(ctx, input)
			},
		},
		"create-service-network": {
			Name:   "create-service-network",
			Fields: fields_create_service_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceNetwork(ctx, input)
			},
		},
		"create-service-network-resource-association": {
			Name:   "create-service-network-resource-association",
			Fields: fields_create_service_network_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceNetworkResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_network_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceNetworkResourceAssociation(ctx, input)
			},
		},
		"create-service-network-service-association": {
			Name:   "create-service-network-service-association",
			Fields: fields_create_service_network_service_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceNetworkServiceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_network_service_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceNetworkServiceAssociation(ctx, input)
			},
		},
		"create-service-network-vpc-association": {
			Name:   "create-service-network-vpc-association",
			Fields: fields_create_service_network_vpc_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceNetworkVpcAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_network_vpc_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceNetworkVpcAssociation(ctx, input)
			},
		},
		"create-target-group": {
			Name:   "create-target-group",
			Fields: fields_create_target_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTargetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_target_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTargetGroup(ctx, input)
			},
		},
		"delete-access-log-subscription": {
			Name:   "delete-access-log-subscription",
			Fields: fields_delete_access_log_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessLogSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_log_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessLogSubscription(ctx, input)
			},
		},
		"delete-auth-policy": {
			Name:   "delete-auth-policy",
			Fields: fields_delete_auth_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAuthPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_auth_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAuthPolicy(ctx, input)
			},
		},
		"delete-domain-verification": {
			Name:   "delete-domain-verification",
			Fields: fields_delete_domain_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainVerification(ctx, input)
			},
		},
		"delete-listener": {
			Name:   "delete-listener",
			Fields: fields_delete_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteListener(ctx, input)
			},
		},
		"delete-resource-configuration": {
			Name:   "delete-resource-configuration",
			Fields: fields_delete_resource_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceConfiguration(ctx, input)
			},
		},
		"delete-resource-endpoint-association": {
			Name:   "delete-resource-endpoint-association",
			Fields: fields_delete_resource_endpoint_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceEndpointAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_endpoint_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceEndpointAssociation(ctx, input)
			},
		},
		"delete-resource-gateway": {
			Name:   "delete-resource-gateway",
			Fields: fields_delete_resource_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceGateway(ctx, input)
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
		"delete-rule": {
			Name:   "delete-rule",
			Fields: fields_delete_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRule(ctx, input)
			},
		},
		"delete-service": {
			Name:   "delete-service",
			Fields: fields_delete_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteService(ctx, input)
			},
		},
		"delete-service-network": {
			Name:   "delete-service-network",
			Fields: fields_delete_service_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceNetwork(ctx, input)
			},
		},
		"delete-service-network-resource-association": {
			Name:   "delete-service-network-resource-association",
			Fields: fields_delete_service_network_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceNetworkResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_network_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceNetworkResourceAssociation(ctx, input)
			},
		},
		"delete-service-network-service-association": {
			Name:   "delete-service-network-service-association",
			Fields: fields_delete_service_network_service_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceNetworkServiceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_network_service_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceNetworkServiceAssociation(ctx, input)
			},
		},
		"delete-service-network-vpc-association": {
			Name:   "delete-service-network-vpc-association",
			Fields: fields_delete_service_network_vpc_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceNetworkVpcAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_network_vpc_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceNetworkVpcAssociation(ctx, input)
			},
		},
		"delete-target-group": {
			Name:   "delete-target-group",
			Fields: fields_delete_target_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTargetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_target_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTargetGroup(ctx, input)
			},
		},
		"deregister-targets": {
			Name:   "deregister-targets",
			Fields: fields_deregister_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTargets(ctx, input)
			},
		},
		"get-access-log-subscription": {
			Name:   "get-access-log-subscription",
			Fields: fields_get_access_log_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessLogSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_log_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessLogSubscription(ctx, input)
			},
		},
		"get-auth-policy": {
			Name:   "get-auth-policy",
			Fields: fields_get_auth_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAuthPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_auth_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAuthPolicy(ctx, input)
			},
		},
		"get-domain-verification": {
			Name:   "get-domain-verification",
			Fields: fields_get_domain_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainVerification(ctx, input)
			},
		},
		"get-listener": {
			Name:   "get-listener",
			Fields: fields_get_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetListener(ctx, input)
			},
		},
		"get-resource-configuration": {
			Name:   "get-resource-configuration",
			Fields: fields_get_resource_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceConfiguration(ctx, input)
			},
		},
		"get-resource-gateway": {
			Name:   "get-resource-gateway",
			Fields: fields_get_resource_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceGateway(ctx, input)
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
		"get-rule": {
			Name:   "get-rule",
			Fields: fields_get_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRule(ctx, input)
			},
		},
		"get-service": {
			Name:   "get-service",
			Fields: fields_get_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetService(ctx, input)
			},
		},
		"get-service-network": {
			Name:   "get-service-network",
			Fields: fields_get_service_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceNetwork(ctx, input)
			},
		},
		"get-service-network-resource-association": {
			Name:   "get-service-network-resource-association",
			Fields: fields_get_service_network_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceNetworkResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_network_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceNetworkResourceAssociation(ctx, input)
			},
		},
		"get-service-network-service-association": {
			Name:   "get-service-network-service-association",
			Fields: fields_get_service_network_service_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceNetworkServiceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_network_service_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceNetworkServiceAssociation(ctx, input)
			},
		},
		"get-service-network-vpc-association": {
			Name:   "get-service-network-vpc-association",
			Fields: fields_get_service_network_vpc_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceNetworkVpcAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_network_vpc_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceNetworkVpcAssociation(ctx, input)
			},
		},
		"get-target-group": {
			Name:   "get-target-group",
			Fields: fields_get_target_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTargetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_target_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTargetGroup(ctx, input)
			},
		},
		"list-access-log-subscriptions": {
			Name:   "list-access-log-subscriptions",
			Fields: fields_list_access_log_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessLogSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_log_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessLogSubscriptions(ctx, input)
				}
				var results []*svc.ListAccessLogSubscriptionsOutput
				p := svc.NewListAccessLogSubscriptionsPaginator(client, input)
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
		"list-domain-verifications": {
			Name:   "list-domain-verifications",
			Fields: fields_list_domain_verifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainVerificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_verifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainVerifications(ctx, input)
				}
				var results []*svc.ListDomainVerificationsOutput
				p := svc.NewListDomainVerificationsPaginator(client, input)
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
		"list-listeners": {
			Name:   "list-listeners",
			Fields: fields_list_listeners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListListenersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_listeners, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListListeners(ctx, input)
				}
				var results []*svc.ListListenersOutput
				p := svc.NewListListenersPaginator(client, input)
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
		"list-resource-configurations": {
			Name:   "list-resource-configurations",
			Fields: fields_list_resource_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceConfigurations(ctx, input)
				}
				var results []*svc.ListResourceConfigurationsOutput
				p := svc.NewListResourceConfigurationsPaginator(client, input)
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
		"list-resource-endpoint-associations": {
			Name:   "list-resource-endpoint-associations",
			Fields: fields_list_resource_endpoint_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceEndpointAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_endpoint_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceEndpointAssociations(ctx, input)
				}
				var results []*svc.ListResourceEndpointAssociationsOutput
				p := svc.NewListResourceEndpointAssociationsPaginator(client, input)
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
		"list-resource-gateways": {
			Name:   "list-resource-gateways",
			Fields: fields_list_resource_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceGateways(ctx, input)
				}
				var results []*svc.ListResourceGatewaysOutput
				p := svc.NewListResourceGatewaysPaginator(client, input)
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
		"list-rules": {
			Name:   "list-rules",
			Fields: fields_list_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRules(ctx, input)
				}
				var results []*svc.ListRulesOutput
				p := svc.NewListRulesPaginator(client, input)
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
		"list-service-network-resource-associations": {
			Name:   "list-service-network-resource-associations",
			Fields: fields_list_service_network_resource_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceNetworkResourceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_network_resource_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceNetworkResourceAssociations(ctx, input)
				}
				var results []*svc.ListServiceNetworkResourceAssociationsOutput
				p := svc.NewListServiceNetworkResourceAssociationsPaginator(client, input)
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
		"list-service-network-service-associations": {
			Name:   "list-service-network-service-associations",
			Fields: fields_list_service_network_service_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceNetworkServiceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_network_service_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceNetworkServiceAssociations(ctx, input)
				}
				var results []*svc.ListServiceNetworkServiceAssociationsOutput
				p := svc.NewListServiceNetworkServiceAssociationsPaginator(client, input)
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
		"list-service-network-vpc-associations": {
			Name:   "list-service-network-vpc-associations",
			Fields: fields_list_service_network_vpc_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceNetworkVpcAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_network_vpc_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceNetworkVpcAssociations(ctx, input)
				}
				var results []*svc.ListServiceNetworkVpcAssociationsOutput
				p := svc.NewListServiceNetworkVpcAssociationsPaginator(client, input)
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
		"list-service-network-vpc-endpoint-associations": {
			Name:   "list-service-network-vpc-endpoint-associations",
			Fields: fields_list_service_network_vpc_endpoint_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceNetworkVpcEndpointAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_network_vpc_endpoint_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceNetworkVpcEndpointAssociations(ctx, input)
				}
				var results []*svc.ListServiceNetworkVpcEndpointAssociationsOutput
				p := svc.NewListServiceNetworkVpcEndpointAssociationsPaginator(client, input)
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
		"list-service-networks": {
			Name:   "list-service-networks",
			Fields: fields_list_service_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceNetworks(ctx, input)
				}
				var results []*svc.ListServiceNetworksOutput
				p := svc.NewListServiceNetworksPaginator(client, input)
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
		"list-services": {
			Name:   "list-services",
			Fields: fields_list_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServices(ctx, input)
				}
				var results []*svc.ListServicesOutput
				p := svc.NewListServicesPaginator(client, input)
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
		"list-target-groups": {
			Name:   "list-target-groups",
			Fields: fields_list_target_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_target_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetGroups(ctx, input)
				}
				var results []*svc.ListTargetGroupsOutput
				p := svc.NewListTargetGroupsPaginator(client, input)
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
		"list-targets": {
			Name:   "list-targets",
			Fields: fields_list_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargets(ctx, input)
				}
				var results []*svc.ListTargetsOutput
				p := svc.NewListTargetsPaginator(client, input)
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
		"put-auth-policy": {
			Name:   "put-auth-policy",
			Fields: fields_put_auth_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAuthPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_auth_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAuthPolicy(ctx, input)
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
		"register-targets": {
			Name:   "register-targets",
			Fields: fields_register_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTargets(ctx, input)
			},
		},
		"start-domain-verification": {
			Name:   "start-domain-verification",
			Fields: fields_start_domain_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDomainVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_domain_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDomainVerification(ctx, input)
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
		"update-access-log-subscription": {
			Name:   "update-access-log-subscription",
			Fields: fields_update_access_log_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessLogSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_log_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessLogSubscription(ctx, input)
			},
		},
		"update-listener": {
			Name:   "update-listener",
			Fields: fields_update_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateListener(ctx, input)
			},
		},
		"update-resource-configuration": {
			Name:   "update-resource-configuration",
			Fields: fields_update_resource_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceConfiguration(ctx, input)
			},
		},
		"update-resource-gateway": {
			Name:   "update-resource-gateway",
			Fields: fields_update_resource_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceGateway(ctx, input)
			},
		},
		"update-rule": {
			Name:   "update-rule",
			Fields: fields_update_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRule(ctx, input)
			},
		},
		"update-service": {
			Name:   "update-service",
			Fields: fields_update_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateService(ctx, input)
			},
		},
		"update-service-network": {
			Name:   "update-service-network",
			Fields: fields_update_service_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceNetwork(ctx, input)
			},
		},
		"update-service-network-vpc-association": {
			Name:   "update-service-network-vpc-association",
			Fields: fields_update_service_network_vpc_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceNetworkVpcAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_network_vpc_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceNetworkVpcAssociation(ctx, input)
			},
		},
		"update-target-group": {
			Name:   "update-target-group",
			Fields: fields_update_target_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTargetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_target_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTargetGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("vpclattice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
