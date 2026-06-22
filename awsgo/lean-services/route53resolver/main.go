package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

var fields_associate_firewall_rule_group = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
	{Name: "MutationProtection", Flag: "mutation-protection", Type: "types.MutationProtectionStatus", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_associate_resolver_endpoint_ip_address = []leanruntime.Field{
	{Name: "IpAddress", Flag: "ip-address", Type: "*types.IpAddressUpdate", Required: true},
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: true},
}

var fields_associate_resolver_query_log_config = []leanruntime.Field{
	{Name: "ResolverQueryLogConfigId", Flag: "resolver-query-log-config-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_associate_resolver_rule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResolverRuleId", Flag: "resolver-rule-id", Type: "*string", Required: true},
	{Name: "VPCId", Flag: "vpcid", Type: "*string", Required: true},
}

var fields_create_firewall_domain_list = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_firewall_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.Action", Required: true},
	{Name: "BlockOverrideDnsType", Flag: "block-override-dns-type", Type: "types.BlockOverrideDnsType", Required: false},
	{Name: "BlockOverrideDomain", Flag: "block-override-domain", Type: "*string", Required: false},
	{Name: "BlockOverrideTtl", Flag: "block-override-ttl", Type: "*int32", Required: false},
	{Name: "BlockResponse", Flag: "block-response", Type: "types.BlockResponse", Required: false},
	{Name: "ConfidenceThreshold", Flag: "confidence-threshold", Type: "types.ConfidenceThreshold", Required: false},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "DnsThreatProtection", Flag: "dns-threat-protection", Type: "types.DnsThreatProtection", Required: false},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: false},
	{Name: "FirewallDomainRedirectionAction", Flag: "firewall-domain-redirection-action", Type: "types.FirewallDomainRedirectionAction", Required: false},
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "Qtype", Flag: "qtype", Type: "*string", Required: false},
}

var fields_create_firewall_rule_group = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_outpost_resolver = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: true},
	{Name: "PreferredInstanceType", Flag: "preferred-instance-type", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_resolver_endpoint = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "Direction", Flag: "direction", Type: "types.ResolverEndpointDirection", Required: true},
	{Name: "IpAddresses", Flag: "ip-addresses", Type: "[]types.IpAddressRequest", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "PreferredInstanceType", Flag: "preferred-instance-type", Type: "*string", Required: false},
	{Name: "Protocols", Flag: "protocols", Type: "[]types.Protocol", Required: false},
	{Name: "ResolverEndpointType", Flag: "resolver-endpoint-type", Type: "types.ResolverEndpointType", Required: false},
	{Name: "RniEnhancedMetricsEnabled", Flag: "rni-enhanced-metrics-enabled", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetNameServerMetricsEnabled", Flag: "target-name-server-metrics-enabled", Type: "*bool", Required: false},
}

var fields_create_resolver_query_log_config = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_resolver_rule = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: true},
	{Name: "DelegationRecord", Flag: "delegation-record", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: false},
	{Name: "RuleType", Flag: "rule-type", Type: "types.RuleTypeOption", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetIps", Flag: "target-ips", Type: "[]types.TargetAddress", Required: false},
}

var fields_delete_firewall_domain_list = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
}

var fields_delete_firewall_rule = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: false},
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
	{Name: "FirewallThreatProtectionId", Flag: "firewall-threat-protection-id", Type: "*string", Required: false},
	{Name: "Qtype", Flag: "qtype", Type: "*string", Required: false},
}

var fields_delete_firewall_rule_group = []leanruntime.Field{
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
}

var fields_delete_outpost_resolver = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_resolver_endpoint = []leanruntime.Field{
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: true},
}

var fields_delete_resolver_query_log_config = []leanruntime.Field{
	{Name: "ResolverQueryLogConfigId", Flag: "resolver-query-log-config-id", Type: "*string", Required: true},
}

var fields_delete_resolver_rule = []leanruntime.Field{
	{Name: "ResolverRuleId", Flag: "resolver-rule-id", Type: "*string", Required: true},
}

var fields_disassociate_firewall_rule_group = []leanruntime.Field{
	{Name: "FirewallRuleGroupAssociationId", Flag: "firewall-rule-group-association-id", Type: "*string", Required: true},
}

var fields_disassociate_resolver_endpoint_ip_address = []leanruntime.Field{
	{Name: "IpAddress", Flag: "ip-address", Type: "*types.IpAddressUpdate", Required: true},
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: true},
}

var fields_disassociate_resolver_query_log_config = []leanruntime.Field{
	{Name: "ResolverQueryLogConfigId", Flag: "resolver-query-log-config-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_disassociate_resolver_rule = []leanruntime.Field{
	{Name: "ResolverRuleId", Flag: "resolver-rule-id", Type: "*string", Required: true},
	{Name: "VPCId", Flag: "vpcid", Type: "*string", Required: true},
}

var fields_get_firewall_config = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_get_firewall_domain_list = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
}

var fields_get_firewall_rule_group = []leanruntime.Field{
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
}

var fields_get_firewall_rule_group_association = []leanruntime.Field{
	{Name: "FirewallRuleGroupAssociationId", Flag: "firewall-rule-group-association-id", Type: "*string", Required: true},
}

var fields_get_firewall_rule_group_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_outpost_resolver = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_resolver_config = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_get_resolver_dnssec_config = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_get_resolver_endpoint = []leanruntime.Field{
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: true},
}

var fields_get_resolver_query_log_config = []leanruntime.Field{
	{Name: "ResolverQueryLogConfigId", Flag: "resolver-query-log-config-id", Type: "*string", Required: true},
}

var fields_get_resolver_query_log_config_association = []leanruntime.Field{
	{Name: "ResolverQueryLogConfigAssociationId", Flag: "resolver-query-log-config-association-id", Type: "*string", Required: true},
}

var fields_get_resolver_query_log_config_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_resolver_rule = []leanruntime.Field{
	{Name: "ResolverRuleId", Flag: "resolver-rule-id", Type: "*string", Required: true},
}

var fields_get_resolver_rule_association = []leanruntime.Field{
	{Name: "ResolverRuleAssociationId", Flag: "resolver-rule-association-id", Type: "*string", Required: true},
}

var fields_get_resolver_rule_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_import_firewall_domains = []leanruntime.Field{
	{Name: "DomainFileUrl", Flag: "domain-file-url", Type: "*string", Required: true},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "types.FirewallDomainImportOperation", Required: true},
}

var fields_list_firewall_configs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_domain_lists = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_domains = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_rule_group_associations = []leanruntime.Field{
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "Status", Flag: "status", Type: "types.FirewallRuleGroupAssociationStatus", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_list_firewall_rule_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_rules = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.Action", Required: false},
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
}

var fields_list_outpost_resolvers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
}

var fields_list_resolver_configs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resolver_dnssec_configs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resolver_endpoint_ip_addresses = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: true},
}

var fields_list_resolver_endpoints = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resolver_query_log_config_associations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_resolver_query_log_configs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_resolver_rule_associations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resolver_rules = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_firewall_rule_group_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "FirewallRuleGroupPolicy", Flag: "firewall-rule-group-policy", Type: "*string", Required: true},
}

var fields_put_resolver_query_log_config_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ResolverQueryLogConfigPolicy", Flag: "resolver-query-log-config-policy", Type: "*string", Required: true},
}

var fields_put_resolver_rule_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ResolverRulePolicy", Flag: "resolver-rule-policy", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_firewall_config = []leanruntime.Field{
	{Name: "FirewallFailOpen", Flag: "firewall-fail-open", Type: "types.FirewallFailOpenStatus", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_update_firewall_domains = []leanruntime.Field{
	{Name: "Domains", Flag: "domains", Type: "[]string", Required: true},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "types.FirewallDomainUpdateOperation", Required: true},
}

var fields_update_firewall_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.Action", Required: false},
	{Name: "BlockOverrideDnsType", Flag: "block-override-dns-type", Type: "types.BlockOverrideDnsType", Required: false},
	{Name: "BlockOverrideDomain", Flag: "block-override-domain", Type: "*string", Required: false},
	{Name: "BlockOverrideTtl", Flag: "block-override-ttl", Type: "*int32", Required: false},
	{Name: "BlockResponse", Flag: "block-response", Type: "types.BlockResponse", Required: false},
	{Name: "ConfidenceThreshold", Flag: "confidence-threshold", Type: "types.ConfidenceThreshold", Required: false},
	{Name: "DnsThreatProtection", Flag: "dns-threat-protection", Type: "types.DnsThreatProtection", Required: false},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: false},
	{Name: "FirewallDomainRedirectionAction", Flag: "firewall-domain-redirection-action", Type: "types.FirewallDomainRedirectionAction", Required: false},
	{Name: "FirewallRuleGroupId", Flag: "firewall-rule-group-id", Type: "*string", Required: true},
	{Name: "FirewallThreatProtectionId", Flag: "firewall-threat-protection-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "Qtype", Flag: "qtype", Type: "*string", Required: false},
}

var fields_update_firewall_rule_group_association = []leanruntime.Field{
	{Name: "FirewallRuleGroupAssociationId", Flag: "firewall-rule-group-association-id", Type: "*string", Required: true},
	{Name: "MutationProtection", Flag: "mutation-protection", Type: "types.MutationProtectionStatus", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
}

var fields_update_outpost_resolver = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PreferredInstanceType", Flag: "preferred-instance-type", Type: "*string", Required: false},
}

var fields_update_resolver_config = []leanruntime.Field{
	{Name: "AutodefinedReverseFlag", Flag: "autodefined-reverse-flag", Type: "types.AutodefinedReverseFlag", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_update_resolver_dnssec_config = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Validation", Flag: "validation", Type: "types.Validation", Required: true},
}

var fields_update_resolver_endpoint = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Protocols", Flag: "protocols", Type: "[]types.Protocol", Required: false},
	{Name: "ResolverEndpointId", Flag: "resolver-endpoint-id", Type: "*string", Required: true},
	{Name: "ResolverEndpointType", Flag: "resolver-endpoint-type", Type: "types.ResolverEndpointType", Required: false},
	{Name: "RniEnhancedMetricsEnabled", Flag: "rni-enhanced-metrics-enabled", Type: "*bool", Required: false},
	{Name: "TargetNameServerMetricsEnabled", Flag: "target-name-server-metrics-enabled", Type: "*bool", Required: false},
	{Name: "UpdateIpAddresses", Flag: "update-ip-addresses", Type: "[]types.UpdateIpAddress", Required: false},
}

var fields_update_resolver_rule = []leanruntime.Field{
	{Name: "Config", Flag: "config", Type: "*types.ResolverRuleConfig", Required: true},
	{Name: "ResolverRuleId", Flag: "resolver-rule-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-firewall-rule-group": {
			Name:   "associate-firewall-rule-group",
			Fields: fields_associate_firewall_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFirewallRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_firewall_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFirewallRuleGroup(ctx, input)
			},
		},
		"associate-resolver-endpoint-ip-address": {
			Name:   "associate-resolver-endpoint-ip-address",
			Fields: fields_associate_resolver_endpoint_ip_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResolverEndpointIpAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resolver_endpoint_ip_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResolverEndpointIpAddress(ctx, input)
			},
		},
		"associate-resolver-query-log-config": {
			Name:   "associate-resolver-query-log-config",
			Fields: fields_associate_resolver_query_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResolverQueryLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resolver_query_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResolverQueryLogConfig(ctx, input)
			},
		},
		"associate-resolver-rule": {
			Name:   "associate-resolver-rule",
			Fields: fields_associate_resolver_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResolverRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resolver_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResolverRule(ctx, input)
			},
		},
		"create-firewall-domain-list": {
			Name:   "create-firewall-domain-list",
			Fields: fields_create_firewall_domain_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFirewallDomainListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_firewall_domain_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFirewallDomainList(ctx, input)
			},
		},
		"create-firewall-rule": {
			Name:   "create-firewall-rule",
			Fields: fields_create_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFirewallRule(ctx, input)
			},
		},
		"create-firewall-rule-group": {
			Name:   "create-firewall-rule-group",
			Fields: fields_create_firewall_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFirewallRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_firewall_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFirewallRuleGroup(ctx, input)
			},
		},
		"create-outpost-resolver": {
			Name:   "create-outpost-resolver",
			Fields: fields_create_outpost_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOutpostResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_outpost_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOutpostResolver(ctx, input)
			},
		},
		"create-resolver-endpoint": {
			Name:   "create-resolver-endpoint",
			Fields: fields_create_resolver_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResolverEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resolver_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResolverEndpoint(ctx, input)
			},
		},
		"create-resolver-query-log-config": {
			Name:   "create-resolver-query-log-config",
			Fields: fields_create_resolver_query_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResolverQueryLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resolver_query_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResolverQueryLogConfig(ctx, input)
			},
		},
		"create-resolver-rule": {
			Name:   "create-resolver-rule",
			Fields: fields_create_resolver_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResolverRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resolver_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResolverRule(ctx, input)
			},
		},
		"delete-firewall-domain-list": {
			Name:   "delete-firewall-domain-list",
			Fields: fields_delete_firewall_domain_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFirewallDomainListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_firewall_domain_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFirewallDomainList(ctx, input)
			},
		},
		"delete-firewall-rule": {
			Name:   "delete-firewall-rule",
			Fields: fields_delete_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFirewallRule(ctx, input)
			},
		},
		"delete-firewall-rule-group": {
			Name:   "delete-firewall-rule-group",
			Fields: fields_delete_firewall_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFirewallRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_firewall_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFirewallRuleGroup(ctx, input)
			},
		},
		"delete-outpost-resolver": {
			Name:   "delete-outpost-resolver",
			Fields: fields_delete_outpost_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOutpostResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_outpost_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOutpostResolver(ctx, input)
			},
		},
		"delete-resolver-endpoint": {
			Name:   "delete-resolver-endpoint",
			Fields: fields_delete_resolver_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResolverEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resolver_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResolverEndpoint(ctx, input)
			},
		},
		"delete-resolver-query-log-config": {
			Name:   "delete-resolver-query-log-config",
			Fields: fields_delete_resolver_query_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResolverQueryLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resolver_query_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResolverQueryLogConfig(ctx, input)
			},
		},
		"delete-resolver-rule": {
			Name:   "delete-resolver-rule",
			Fields: fields_delete_resolver_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResolverRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resolver_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResolverRule(ctx, input)
			},
		},
		"disassociate-firewall-rule-group": {
			Name:   "disassociate-firewall-rule-group",
			Fields: fields_disassociate_firewall_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFirewallRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_firewall_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFirewallRuleGroup(ctx, input)
			},
		},
		"disassociate-resolver-endpoint-ip-address": {
			Name:   "disassociate-resolver-endpoint-ip-address",
			Fields: fields_disassociate_resolver_endpoint_ip_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResolverEndpointIpAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resolver_endpoint_ip_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResolverEndpointIpAddress(ctx, input)
			},
		},
		"disassociate-resolver-query-log-config": {
			Name:   "disassociate-resolver-query-log-config",
			Fields: fields_disassociate_resolver_query_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResolverQueryLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resolver_query_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResolverQueryLogConfig(ctx, input)
			},
		},
		"disassociate-resolver-rule": {
			Name:   "disassociate-resolver-rule",
			Fields: fields_disassociate_resolver_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResolverRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resolver_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResolverRule(ctx, input)
			},
		},
		"get-firewall-config": {
			Name:   "get-firewall-config",
			Fields: fields_get_firewall_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFirewallConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_firewall_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFirewallConfig(ctx, input)
			},
		},
		"get-firewall-domain-list": {
			Name:   "get-firewall-domain-list",
			Fields: fields_get_firewall_domain_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFirewallDomainListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_firewall_domain_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFirewallDomainList(ctx, input)
			},
		},
		"get-firewall-rule-group": {
			Name:   "get-firewall-rule-group",
			Fields: fields_get_firewall_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFirewallRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_firewall_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFirewallRuleGroup(ctx, input)
			},
		},
		"get-firewall-rule-group-association": {
			Name:   "get-firewall-rule-group-association",
			Fields: fields_get_firewall_rule_group_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFirewallRuleGroupAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_firewall_rule_group_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFirewallRuleGroupAssociation(ctx, input)
			},
		},
		"get-firewall-rule-group-policy": {
			Name:   "get-firewall-rule-group-policy",
			Fields: fields_get_firewall_rule_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFirewallRuleGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_firewall_rule_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFirewallRuleGroupPolicy(ctx, input)
			},
		},
		"get-outpost-resolver": {
			Name:   "get-outpost-resolver",
			Fields: fields_get_outpost_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutpostResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_outpost_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOutpostResolver(ctx, input)
			},
		},
		"get-resolver-config": {
			Name:   "get-resolver-config",
			Fields: fields_get_resolver_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverConfig(ctx, input)
			},
		},
		"get-resolver-dnssec-config": {
			Name:   "get-resolver-dnssec-config",
			Fields: fields_get_resolver_dnssec_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverDnssecConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_dnssec_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverDnssecConfig(ctx, input)
			},
		},
		"get-resolver-endpoint": {
			Name:   "get-resolver-endpoint",
			Fields: fields_get_resolver_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverEndpoint(ctx, input)
			},
		},
		"get-resolver-query-log-config": {
			Name:   "get-resolver-query-log-config",
			Fields: fields_get_resolver_query_log_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverQueryLogConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_query_log_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverQueryLogConfig(ctx, input)
			},
		},
		"get-resolver-query-log-config-association": {
			Name:   "get-resolver-query-log-config-association",
			Fields: fields_get_resolver_query_log_config_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverQueryLogConfigAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_query_log_config_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverQueryLogConfigAssociation(ctx, input)
			},
		},
		"get-resolver-query-log-config-policy": {
			Name:   "get-resolver-query-log-config-policy",
			Fields: fields_get_resolver_query_log_config_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverQueryLogConfigPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_query_log_config_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverQueryLogConfigPolicy(ctx, input)
			},
		},
		"get-resolver-rule": {
			Name:   "get-resolver-rule",
			Fields: fields_get_resolver_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverRule(ctx, input)
			},
		},
		"get-resolver-rule-association": {
			Name:   "get-resolver-rule-association",
			Fields: fields_get_resolver_rule_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverRuleAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_rule_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverRuleAssociation(ctx, input)
			},
		},
		"get-resolver-rule-policy": {
			Name:   "get-resolver-rule-policy",
			Fields: fields_get_resolver_rule_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResolverRulePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resolver_rule_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResolverRulePolicy(ctx, input)
			},
		},
		"import-firewall-domains": {
			Name:   "import-firewall-domains",
			Fields: fields_import_firewall_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportFirewallDomainsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_firewall_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportFirewallDomains(ctx, input)
			},
		},
		"list-firewall-configs": {
			Name:   "list-firewall-configs",
			Fields: fields_list_firewall_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallConfigs(ctx, input)
				}
				var results []*svc.ListFirewallConfigsOutput
				p := svc.NewListFirewallConfigsPaginator(client, input)
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
		"list-firewall-domain-lists": {
			Name:   "list-firewall-domain-lists",
			Fields: fields_list_firewall_domain_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallDomainListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_domain_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallDomainLists(ctx, input)
				}
				var results []*svc.ListFirewallDomainListsOutput
				p := svc.NewListFirewallDomainListsPaginator(client, input)
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
		"list-firewall-domains": {
			Name:   "list-firewall-domains",
			Fields: fields_list_firewall_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallDomains(ctx, input)
				}
				var results []*svc.ListFirewallDomainsOutput
				p := svc.NewListFirewallDomainsPaginator(client, input)
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
		"list-firewall-rule-group-associations": {
			Name:   "list-firewall-rule-group-associations",
			Fields: fields_list_firewall_rule_group_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallRuleGroupAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_rule_group_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallRuleGroupAssociations(ctx, input)
				}
				var results []*svc.ListFirewallRuleGroupAssociationsOutput
				p := svc.NewListFirewallRuleGroupAssociationsPaginator(client, input)
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
		"list-firewall-rule-groups": {
			Name:   "list-firewall-rule-groups",
			Fields: fields_list_firewall_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallRuleGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_rule_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallRuleGroups(ctx, input)
				}
				var results []*svc.ListFirewallRuleGroupsOutput
				p := svc.NewListFirewallRuleGroupsPaginator(client, input)
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
		"list-firewall-rules": {
			Name:   "list-firewall-rules",
			Fields: fields_list_firewall_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallRules(ctx, input)
				}
				var results []*svc.ListFirewallRulesOutput
				p := svc.NewListFirewallRulesPaginator(client, input)
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
		"list-outpost-resolvers": {
			Name:   "list-outpost-resolvers",
			Fields: fields_list_outpost_resolvers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOutpostResolversInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_outpost_resolvers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOutpostResolvers(ctx, input)
				}
				var results []*svc.ListOutpostResolversOutput
				p := svc.NewListOutpostResolversPaginator(client, input)
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
		"list-resolver-configs": {
			Name:   "list-resolver-configs",
			Fields: fields_list_resolver_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverConfigs(ctx, input)
				}
				var results []*svc.ListResolverConfigsOutput
				p := svc.NewListResolverConfigsPaginator(client, input)
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
		"list-resolver-dnssec-configs": {
			Name:   "list-resolver-dnssec-configs",
			Fields: fields_list_resolver_dnssec_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverDnssecConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_dnssec_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverDnssecConfigs(ctx, input)
				}
				var results []*svc.ListResolverDnssecConfigsOutput
				p := svc.NewListResolverDnssecConfigsPaginator(client, input)
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
		"list-resolver-endpoint-ip-addresses": {
			Name:   "list-resolver-endpoint-ip-addresses",
			Fields: fields_list_resolver_endpoint_ip_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverEndpointIpAddressesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_endpoint_ip_addresses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverEndpointIpAddresses(ctx, input)
				}
				var results []*svc.ListResolverEndpointIpAddressesOutput
				p := svc.NewListResolverEndpointIpAddressesPaginator(client, input)
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
		"list-resolver-endpoints": {
			Name:   "list-resolver-endpoints",
			Fields: fields_list_resolver_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverEndpoints(ctx, input)
				}
				var results []*svc.ListResolverEndpointsOutput
				p := svc.NewListResolverEndpointsPaginator(client, input)
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
		"list-resolver-query-log-config-associations": {
			Name:   "list-resolver-query-log-config-associations",
			Fields: fields_list_resolver_query_log_config_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverQueryLogConfigAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_query_log_config_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverQueryLogConfigAssociations(ctx, input)
				}
				var results []*svc.ListResolverQueryLogConfigAssociationsOutput
				p := svc.NewListResolverQueryLogConfigAssociationsPaginator(client, input)
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
		"list-resolver-query-log-configs": {
			Name:   "list-resolver-query-log-configs",
			Fields: fields_list_resolver_query_log_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverQueryLogConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_query_log_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverQueryLogConfigs(ctx, input)
				}
				var results []*svc.ListResolverQueryLogConfigsOutput
				p := svc.NewListResolverQueryLogConfigsPaginator(client, input)
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
		"list-resolver-rule-associations": {
			Name:   "list-resolver-rule-associations",
			Fields: fields_list_resolver_rule_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverRuleAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_rule_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverRuleAssociations(ctx, input)
				}
				var results []*svc.ListResolverRuleAssociationsOutput
				p := svc.NewListResolverRuleAssociationsPaginator(client, input)
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
		"list-resolver-rules": {
			Name:   "list-resolver-rules",
			Fields: fields_list_resolver_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResolverRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resolver_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResolverRules(ctx, input)
				}
				var results []*svc.ListResolverRulesOutput
				p := svc.NewListResolverRulesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"put-firewall-rule-group-policy": {
			Name:   "put-firewall-rule-group-policy",
			Fields: fields_put_firewall_rule_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFirewallRuleGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_firewall_rule_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFirewallRuleGroupPolicy(ctx, input)
			},
		},
		"put-resolver-query-log-config-policy": {
			Name:   "put-resolver-query-log-config-policy",
			Fields: fields_put_resolver_query_log_config_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResolverQueryLogConfigPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resolver_query_log_config_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResolverQueryLogConfigPolicy(ctx, input)
			},
		},
		"put-resolver-rule-policy": {
			Name:   "put-resolver-rule-policy",
			Fields: fields_put_resolver_rule_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResolverRulePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resolver_rule_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResolverRulePolicy(ctx, input)
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
		"update-firewall-config": {
			Name:   "update-firewall-config",
			Fields: fields_update_firewall_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallConfig(ctx, input)
			},
		},
		"update-firewall-domains": {
			Name:   "update-firewall-domains",
			Fields: fields_update_firewall_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallDomainsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallDomains(ctx, input)
			},
		},
		"update-firewall-rule": {
			Name:   "update-firewall-rule",
			Fields: fields_update_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallRule(ctx, input)
			},
		},
		"update-firewall-rule-group-association": {
			Name:   "update-firewall-rule-group-association",
			Fields: fields_update_firewall_rule_group_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallRuleGroupAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_rule_group_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallRuleGroupAssociation(ctx, input)
			},
		},
		"update-outpost-resolver": {
			Name:   "update-outpost-resolver",
			Fields: fields_update_outpost_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOutpostResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_outpost_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOutpostResolver(ctx, input)
			},
		},
		"update-resolver-config": {
			Name:   "update-resolver-config",
			Fields: fields_update_resolver_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResolverConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resolver_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResolverConfig(ctx, input)
			},
		},
		"update-resolver-dnssec-config": {
			Name:   "update-resolver-dnssec-config",
			Fields: fields_update_resolver_dnssec_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResolverDnssecConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resolver_dnssec_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResolverDnssecConfig(ctx, input)
			},
		},
		"update-resolver-endpoint": {
			Name:   "update-resolver-endpoint",
			Fields: fields_update_resolver_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResolverEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resolver_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResolverEndpoint(ctx, input)
			},
		},
		"update-resolver-rule": {
			Name:   "update-resolver-rule",
			Fields: fields_update_resolver_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResolverRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resolver_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResolverRule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53resolver", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
