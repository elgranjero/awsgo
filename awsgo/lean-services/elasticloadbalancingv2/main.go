package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

var fields_add_listener_certificates = []leanruntime.Field{
	{Name: "Certificates", Flag: "certificates", Type: "[]types.Certificate", Required: true},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_add_tags = []leanruntime.Field{
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_add_trust_store_revocations = []leanruntime.Field{
	{Name: "RevocationContents", Flag: "revocation-contents", Type: "[]types.RevocationContent", Required: false},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_create_listener = []leanruntime.Field{
	{Name: "AlpnPolicy", Flag: "alpn-policy", Type: "[]string", Required: false},
	{Name: "Certificates", Flag: "certificates", Type: "[]types.Certificate", Required: false},
	{Name: "DefaultActions", Flag: "default-actions", Type: "[]types.Action", Required: true},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
	{Name: "MutualAuthentication", Flag: "mutual-authentication", Type: "*types.MutualAuthenticationAttributes", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.ProtocolEnum", Required: false},
	{Name: "SslPolicy", Flag: "ssl-policy", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_load_balancer = []leanruntime.Field{
	{Name: "CustomerOwnedIpv4Pool", Flag: "customer-owned-ipv4-pool", Type: "*string", Required: false},
	{Name: "EnablePrefixForIpv6SourceNat", Flag: "enable-prefix-for-ipv6-source-nat", Type: "types.EnablePrefixForIpv6SourceNatEnum", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "IpamPools", Flag: "ipam-pools", Type: "*types.IpamPools", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scheme", Flag: "scheme", Type: "types.LoadBalancerSchemeEnum", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "SubnetMappings", Flag: "subnet-mappings", Type: "[]types.SubnetMapping", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.LoadBalancerTypeEnum", Required: false},
}

var fields_create_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.Action", Required: true},
	{Name: "Conditions", Flag: "conditions", Type: "[]types.RuleCondition", Required: true},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Transforms", Flag: "transforms", Type: "[]types.RuleTransform", Required: false},
}

var fields_create_target_group = []leanruntime.Field{
	{Name: "HealthCheckEnabled", Flag: "health-check-enabled", Type: "*bool", Required: false},
	{Name: "HealthCheckIntervalSeconds", Flag: "health-check-interval-seconds", Type: "*int32", Required: false},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "HealthCheckPort", Flag: "health-check-port", Type: "*string", Required: false},
	{Name: "HealthCheckProtocol", Flag: "health-check-protocol", Type: "types.ProtocolEnum", Required: false},
	{Name: "HealthCheckTimeoutSeconds", Flag: "health-check-timeout-seconds", Type: "*int32", Required: false},
	{Name: "HealthyThresholdCount", Flag: "healthy-threshold-count", Type: "*int32", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.TargetGroupIpAddressTypeEnum", Required: false},
	{Name: "Matcher", Flag: "matcher", Type: "*types.Matcher", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.ProtocolEnum", Required: false},
	{Name: "ProtocolVersion", Flag: "protocol-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetControlPort", Flag: "target-control-port", Type: "*int32", Required: false},
	{Name: "TargetType", Flag: "target-type", Type: "types.TargetTypeEnum", Required: false},
	{Name: "UnhealthyThresholdCount", Flag: "unhealthy-threshold-count", Type: "*int32", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_create_trust_store = []leanruntime.Field{
	{Name: "CaCertificatesBundleS3Bucket", Flag: "ca-certificates-bundle-s3-bucket", Type: "*string", Required: true},
	{Name: "CaCertificatesBundleS3Key", Flag: "ca-certificates-bundle-s3-key", Type: "*string", Required: true},
	{Name: "CaCertificatesBundleS3ObjectVersion", Flag: "ca-certificates-bundle-s3-object-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_listener = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_delete_load_balancer = []leanruntime.Field{
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "RuleArn", Flag: "rule-arn", Type: "*string", Required: true},
}

var fields_delete_shared_trust_store_association = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_delete_target_group = []leanruntime.Field{
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
}

var fields_delete_trust_store = []leanruntime.Field{
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_deregister_targets = []leanruntime.Field{
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.TargetDescription", Required: true},
}

var fields_describe_account_limits = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_capacity_reservation = []leanruntime.Field{
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
}

var fields_describe_listener_attributes = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_describe_listener_certificates = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_listeners = []leanruntime.Field{
	{Name: "ListenerArns", Flag: "listener-arns", Type: "[]string", Required: false},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_load_balancer_attributes = []leanruntime.Field{
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
}

var fields_describe_load_balancers = []leanruntime.Field{
	{Name: "LoadBalancerArns", Flag: "load-balancer-arns", Type: "[]string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_rules = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "RuleArns", Flag: "rule-arns", Type: "[]string", Required: false},
}

var fields_describe_ssl_policies = []leanruntime.Field{
	{Name: "LoadBalancerType", Flag: "load-balancer-type", Type: "types.LoadBalancerTypeEnum", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
}

var fields_describe_target_group_attributes = []leanruntime.Field{
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
}

var fields_describe_target_groups = []leanruntime.Field{
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "TargetGroupArns", Flag: "target-group-arns", Type: "[]string", Required: false},
}

var fields_describe_target_health = []leanruntime.Field{
	{Name: "Include", Flag: "include", Type: "[]types.DescribeTargetHealthInputIncludeEnum", Required: false},
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.TargetDescription", Required: false},
}

var fields_describe_trust_store_associations = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_describe_trust_store_revocations = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "RevocationIds", Flag: "revocation-ids", Type: "[]int64", Required: false},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_describe_trust_stores = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "TrustStoreArns", Flag: "trust-store-arns", Type: "[]string", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_trust_store_ca_certificates_bundle = []leanruntime.Field{
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_get_trust_store_revocation_content = []leanruntime.Field{
	{Name: "RevocationId", Flag: "revocation-id", Type: "*int64", Required: true},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_modify_capacity_reservation = []leanruntime.Field{
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
	{Name: "MinimumLoadBalancerCapacity", Flag: "minimum-load-balancer-capacity", Type: "*types.MinimumLoadBalancerCapacity", Required: false},
	{Name: "ResetCapacityReservation", Flag: "reset-capacity-reservation", Type: "*bool", Required: false},
}

var fields_modify_ip_pools = []leanruntime.Field{
	{Name: "IpamPools", Flag: "ipam-pools", Type: "*types.IpamPools", Required: false},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
	{Name: "RemoveIpamPools", Flag: "remove-ipam-pools", Type: "[]types.RemoveIpamPoolEnum", Required: false},
}

var fields_modify_listener = []leanruntime.Field{
	{Name: "AlpnPolicy", Flag: "alpn-policy", Type: "[]string", Required: false},
	{Name: "Certificates", Flag: "certificates", Type: "[]types.Certificate", Required: false},
	{Name: "DefaultActions", Flag: "default-actions", Type: "[]types.Action", Required: false},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "MutualAuthentication", Flag: "mutual-authentication", Type: "*types.MutualAuthenticationAttributes", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.ProtocolEnum", Required: false},
	{Name: "SslPolicy", Flag: "ssl-policy", Type: "*string", Required: false},
}

var fields_modify_listener_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.ListenerAttribute", Required: true},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_modify_load_balancer_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.LoadBalancerAttribute", Required: true},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
}

var fields_modify_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.Action", Required: false},
	{Name: "Conditions", Flag: "conditions", Type: "[]types.RuleCondition", Required: false},
	{Name: "ResetTransforms", Flag: "reset-transforms", Type: "*bool", Required: false},
	{Name: "RuleArn", Flag: "rule-arn", Type: "*string", Required: true},
	{Name: "Transforms", Flag: "transforms", Type: "[]types.RuleTransform", Required: false},
}

var fields_modify_target_group = []leanruntime.Field{
	{Name: "HealthCheckEnabled", Flag: "health-check-enabled", Type: "*bool", Required: false},
	{Name: "HealthCheckIntervalSeconds", Flag: "health-check-interval-seconds", Type: "*int32", Required: false},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "HealthCheckPort", Flag: "health-check-port", Type: "*string", Required: false},
	{Name: "HealthCheckProtocol", Flag: "health-check-protocol", Type: "types.ProtocolEnum", Required: false},
	{Name: "HealthCheckTimeoutSeconds", Flag: "health-check-timeout-seconds", Type: "*int32", Required: false},
	{Name: "HealthyThresholdCount", Flag: "healthy-threshold-count", Type: "*int32", Required: false},
	{Name: "Matcher", Flag: "matcher", Type: "*types.Matcher", Required: false},
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
	{Name: "UnhealthyThresholdCount", Flag: "unhealthy-threshold-count", Type: "*int32", Required: false},
}

var fields_modify_target_group_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.TargetGroupAttribute", Required: true},
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
}

var fields_modify_trust_store = []leanruntime.Field{
	{Name: "CaCertificatesBundleS3Bucket", Flag: "ca-certificates-bundle-s3-bucket", Type: "*string", Required: true},
	{Name: "CaCertificatesBundleS3Key", Flag: "ca-certificates-bundle-s3-key", Type: "*string", Required: true},
	{Name: "CaCertificatesBundleS3ObjectVersion", Flag: "ca-certificates-bundle-s3-object-version", Type: "*string", Required: false},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_register_targets = []leanruntime.Field{
	{Name: "TargetGroupArn", Flag: "target-group-arn", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.TargetDescription", Required: true},
}

var fields_remove_listener_certificates = []leanruntime.Field{
	{Name: "Certificates", Flag: "certificates", Type: "[]types.Certificate", Required: true},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_remove_trust_store_revocations = []leanruntime.Field{
	{Name: "RevocationIds", Flag: "revocation-ids", Type: "[]int64", Required: true},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_set_ip_address_type = []leanruntime.Field{
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: true},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
}

var fields_set_rule_priorities = []leanruntime.Field{
	{Name: "RulePriorities", Flag: "rule-priorities", Type: "[]types.RulePriorityPair", Required: true},
}

var fields_set_security_groups = []leanruntime.Field{
	{Name: "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic", Flag: "enforce-security-group-inbound-rules-on-private-link-traffic", Type: "types.EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficEnum", Required: false},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: true},
}

var fields_set_subnets = []leanruntime.Field{
	{Name: "EnablePrefixForIpv6SourceNat", Flag: "enable-prefix-for-ipv6-source-nat", Type: "types.EnablePrefixForIpv6SourceNatEnum", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "LoadBalancerArn", Flag: "load-balancer-arn", Type: "*string", Required: true},
	{Name: "SubnetMappings", Flag: "subnet-mappings", Type: "[]types.SubnetMapping", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-listener-certificates": {
			Name:   "add-listener-certificates",
			Fields: fields_add_listener_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddListenerCertificatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_listener_certificates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddListenerCertificates(ctx, input)
			},
		},
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"add-trust-store-revocations": {
			Name:   "add-trust-store-revocations",
			Fields: fields_add_trust_store_revocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTrustStoreRevocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_trust_store_revocations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTrustStoreRevocations(ctx, input)
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
		"create-load-balancer": {
			Name:   "create-load-balancer",
			Fields: fields_create_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoadBalancer(ctx, input)
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
		"delete-load-balancer": {
			Name:   "delete-load-balancer",
			Fields: fields_delete_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoadBalancer(ctx, input)
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
		"delete-shared-trust-store-association": {
			Name:   "delete-shared-trust-store-association",
			Fields: fields_delete_shared_trust_store_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSharedTrustStoreAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_shared_trust_store_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSharedTrustStoreAssociation(ctx, input)
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
		"describe-account-limits": {
			Name:   "describe-account-limits",
			Fields: fields_describe_account_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_account_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAccountLimits(ctx, input)
				}
				var results []*svc.DescribeAccountLimitsOutput
				p := svc.NewDescribeAccountLimitsPaginator(client, input)
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
		"describe-capacity-reservation": {
			Name:   "describe-capacity-reservation",
			Fields: fields_describe_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCapacityReservation(ctx, input)
			},
		},
		"describe-listener-attributes": {
			Name:   "describe-listener-attributes",
			Fields: fields_describe_listener_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeListenerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_listener_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeListenerAttributes(ctx, input)
			},
		},
		"describe-listener-certificates": {
			Name:   "describe-listener-certificates",
			Fields: fields_describe_listener_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeListenerCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_listener_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeListenerCertificates(ctx, input)
				}
				var results []*svc.DescribeListenerCertificatesOutput
				p := svc.NewDescribeListenerCertificatesPaginator(client, input)
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
		"describe-listeners": {
			Name:   "describe-listeners",
			Fields: fields_describe_listeners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeListenersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_listeners, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeListeners(ctx, input)
				}
				var results []*svc.DescribeListenersOutput
				p := svc.NewDescribeListenersPaginator(client, input)
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
		"describe-load-balancer-attributes": {
			Name:   "describe-load-balancer-attributes",
			Fields: fields_describe_load_balancer_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_load_balancer_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoadBalancerAttributes(ctx, input)
			},
		},
		"describe-load-balancers": {
			Name:   "describe-load-balancers",
			Fields: fields_describe_load_balancers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_load_balancers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLoadBalancers(ctx, input)
				}
				var results []*svc.DescribeLoadBalancersOutput
				p := svc.NewDescribeLoadBalancersPaginator(client, input)
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
		"describe-rules": {
			Name:   "describe-rules",
			Fields: fields_describe_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRules(ctx, input)
				}
				var results []*svc.DescribeRulesOutput
				p := svc.NewDescribeRulesPaginator(client, input)
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
		"describe-ssl-policies": {
			Name:   "describe-ssl-policies",
			Fields: fields_describe_ssl_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSSLPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ssl_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSSLPolicies(ctx, input)
			},
		},
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTags(ctx, input)
			},
		},
		"describe-target-group-attributes": {
			Name:   "describe-target-group-attributes",
			Fields: fields_describe_target_group_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTargetGroupAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_target_group_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTargetGroupAttributes(ctx, input)
			},
		},
		"describe-target-groups": {
			Name:   "describe-target-groups",
			Fields: fields_describe_target_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTargetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_target_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTargetGroups(ctx, input)
				}
				var results []*svc.DescribeTargetGroupsOutput
				p := svc.NewDescribeTargetGroupsPaginator(client, input)
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
		"describe-target-health": {
			Name:   "describe-target-health",
			Fields: fields_describe_target_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTargetHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_target_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTargetHealth(ctx, input)
			},
		},
		"describe-trust-store-associations": {
			Name:   "describe-trust-store-associations",
			Fields: fields_describe_trust_store_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustStoreAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_trust_store_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrustStoreAssociations(ctx, input)
				}
				var results []*svc.DescribeTrustStoreAssociationsOutput
				p := svc.NewDescribeTrustStoreAssociationsPaginator(client, input)
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
		"describe-trust-store-revocations": {
			Name:   "describe-trust-store-revocations",
			Fields: fields_describe_trust_store_revocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustStoreRevocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_trust_store_revocations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrustStoreRevocations(ctx, input)
				}
				var results []*svc.DescribeTrustStoreRevocationsOutput
				p := svc.NewDescribeTrustStoreRevocationsPaginator(client, input)
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
		"describe-trust-stores": {
			Name:   "describe-trust-stores",
			Fields: fields_describe_trust_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_trust_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrustStores(ctx, input)
				}
				var results []*svc.DescribeTrustStoresOutput
				p := svc.NewDescribeTrustStoresPaginator(client, input)
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
		"get-trust-store-ca-certificates-bundle": {
			Name:   "get-trust-store-ca-certificates-bundle",
			Fields: fields_get_trust_store_ca_certificates_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustStoreCaCertificatesBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trust_store_ca_certificates_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustStoreCaCertificatesBundle(ctx, input)
			},
		},
		"get-trust-store-revocation-content": {
			Name:   "get-trust-store-revocation-content",
			Fields: fields_get_trust_store_revocation_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustStoreRevocationContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trust_store_revocation_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustStoreRevocationContent(ctx, input)
			},
		},
		"modify-capacity-reservation": {
			Name:   "modify-capacity-reservation",
			Fields: fields_modify_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCapacityReservation(ctx, input)
			},
		},
		"modify-ip-pools": {
			Name:   "modify-ip-pools",
			Fields: fields_modify_ip_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpPoolsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ip_pools, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpPools(ctx, input)
			},
		},
		"modify-listener": {
			Name:   "modify-listener",
			Fields: fields_modify_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyListener(ctx, input)
			},
		},
		"modify-listener-attributes": {
			Name:   "modify-listener-attributes",
			Fields: fields_modify_listener_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyListenerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_listener_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyListenerAttributes(ctx, input)
			},
		},
		"modify-load-balancer-attributes": {
			Name:   "modify-load-balancer-attributes",
			Fields: fields_modify_load_balancer_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyLoadBalancerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_load_balancer_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyLoadBalancerAttributes(ctx, input)
			},
		},
		"modify-rule": {
			Name:   "modify-rule",
			Fields: fields_modify_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyRule(ctx, input)
			},
		},
		"modify-target-group": {
			Name:   "modify-target-group",
			Fields: fields_modify_target_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTargetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_target_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTargetGroup(ctx, input)
			},
		},
		"modify-target-group-attributes": {
			Name:   "modify-target-group-attributes",
			Fields: fields_modify_target_group_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTargetGroupAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_target_group_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTargetGroupAttributes(ctx, input)
			},
		},
		"modify-trust-store": {
			Name:   "modify-trust-store",
			Fields: fields_modify_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTrustStore(ctx, input)
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
		"remove-listener-certificates": {
			Name:   "remove-listener-certificates",
			Fields: fields_remove_listener_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveListenerCertificatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_listener_certificates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveListenerCertificates(ctx, input)
			},
		},
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"remove-trust-store-revocations": {
			Name:   "remove-trust-store-revocations",
			Fields: fields_remove_trust_store_revocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTrustStoreRevocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_trust_store_revocations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTrustStoreRevocations(ctx, input)
			},
		},
		"set-ip-address-type": {
			Name:   "set-ip-address-type",
			Fields: fields_set_ip_address_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIpAddressTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_ip_address_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIpAddressType(ctx, input)
			},
		},
		"set-rule-priorities": {
			Name:   "set-rule-priorities",
			Fields: fields_set_rule_priorities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetRulePrioritiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_rule_priorities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetRulePriorities(ctx, input)
			},
		},
		"set-security-groups": {
			Name:   "set-security-groups",
			Fields: fields_set_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetSecurityGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_security_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetSecurityGroups(ctx, input)
			},
		},
		"set-subnets": {
			Name:   "set-subnets",
			Fields: fields_set_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetSubnetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_subnets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetSubnets(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("elasticloadbalancingv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
