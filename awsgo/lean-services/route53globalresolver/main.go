package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53globalresolver"
)

var fields_associate_hosted_zone = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_batch_create_firewall_rule = []leanruntime.Field{
	{Name: "FirewallRules", Flag: "firewall-rules", Type: "[]types.BatchCreateFirewallRuleInputItem", Required: true},
}

var fields_batch_delete_firewall_rule = []leanruntime.Field{
	{Name: "FirewallRules", Flag: "firewall-rules", Type: "[]types.BatchDeleteFirewallRuleInputItem", Required: true},
}

var fields_batch_update_firewall_rule = []leanruntime.Field{
	{Name: "FirewallRules", Flag: "firewall-rules", Type: "[]types.BatchUpdateFirewallRuleInputItem", Required: true},
}

var fields_create_access_source = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.DnsProtocol", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_access_token = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
	{Name: "ExpiresAt", Flag: "expires-at", Type: "*time.Time", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dns_view = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DnssecValidation", Flag: "dnssec-validation", Type: "types.DnsSecValidationType", Required: false},
	{Name: "EdnsClientSubnet", Flag: "edns-client-subnet", Type: "types.EdnsClientSubnetType", Required: false},
	{Name: "FirewallRulesFailOpen", Flag: "firewall-rules-fail-open", Type: "types.FirewallRulesFailOpenType", Required: false},
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_firewall_domain_list = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_firewall_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FirewallRuleAction", Required: true},
	{Name: "BlockOverrideDnsType", Flag: "block-override-dns-type", Type: "types.BlockOverrideDnsQueryType", Required: false},
	{Name: "BlockOverrideDomain", Flag: "block-override-domain", Type: "*string", Required: false},
	{Name: "BlockOverrideTtl", Flag: "block-override-ttl", Type: "*int32", Required: false},
	{Name: "BlockResponse", Flag: "block-response", Type: "types.FirewallBlockResponse", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConfidenceThreshold", Flag: "confidence-threshold", Type: "types.ConfidenceThreshold", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DnsAdvancedProtection", Flag: "dns-advanced-protection", Type: "types.DnsAdvancedProtection", Required: false},
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int64", Required: false},
	{Name: "QType", Flag: "qtype", Type: "*string", Required: false},
}

var fields_create_global_resolver = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ObservabilityRegion", Flag: "observability-region", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_access_source = []leanruntime.Field{
	{Name: "AccessSourceId", Flag: "access-source-id", Type: "*string", Required: true},
}

var fields_delete_access_token = []leanruntime.Field{
	{Name: "AccessTokenId", Flag: "access-token-id", Type: "*string", Required: true},
}

var fields_delete_dns_view = []leanruntime.Field{
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
}

var fields_delete_firewall_domain_list = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
}

var fields_delete_firewall_rule = []leanruntime.Field{
	{Name: "FirewallRuleId", Flag: "firewall-rule-id", Type: "*string", Required: true},
}

var fields_delete_global_resolver = []leanruntime.Field{
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: true},
}

var fields_disable_dns_view = []leanruntime.Field{
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
}

var fields_disassociate_hosted_zone = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_enable_dns_view = []leanruntime.Field{
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
}

var fields_get_access_source = []leanruntime.Field{
	{Name: "AccessSourceId", Flag: "access-source-id", Type: "*string", Required: true},
}

var fields_get_access_token = []leanruntime.Field{
	{Name: "AccessTokenId", Flag: "access-token-id", Type: "*string", Required: true},
}

var fields_get_dns_view = []leanruntime.Field{
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
}

var fields_get_firewall_domain_list = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
}

var fields_get_firewall_rule = []leanruntime.Field{
	{Name: "FirewallRuleId", Flag: "firewall-rule-id", Type: "*string", Required: true},
}

var fields_get_global_resolver = []leanruntime.Field{
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: true},
}

var fields_get_hosted_zone_association = []leanruntime.Field{
	{Name: "HostedZoneAssociationId", Flag: "hosted-zone-association-id", Type: "*string", Required: true},
}

var fields_get_managed_firewall_domain_list = []leanruntime.Field{
	{Name: "ManagedFirewallDomainListId", Flag: "managed-firewall-domain-list-id", Type: "*string", Required: true},
}

var fields_import_firewall_domains = []leanruntime.Field{
	{Name: "DomainFileUrl", Flag: "domain-file-url", Type: "*string", Required: true},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: true},
}

var fields_list_access_sources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_tokens = []leanruntime.Field{
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dns_views = []leanruntime.Field{
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_domain_lists = []leanruntime.Field{
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_domains = []leanruntime.Field{
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_rules = []leanruntime.Field{
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_global_resolvers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hosted_zone_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_managed_firewall_domain_lists = []leanruntime.Field{
	{Name: "ManagedFirewallDomainListType", Flag: "managed-firewall-domain-list-type", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_source = []leanruntime.Field{
	{Name: "AccessSourceId", Flag: "access-source-id", Type: "*string", Required: true},
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.DnsProtocol", Required: false},
}

var fields_update_access_token = []leanruntime.Field{
	{Name: "AccessTokenId", Flag: "access-token-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_dns_view = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DnsViewId", Flag: "dns-view-id", Type: "*string", Required: true},
	{Name: "DnssecValidation", Flag: "dnssec-validation", Type: "types.DnsSecValidationType", Required: false},
	{Name: "EdnsClientSubnet", Flag: "edns-client-subnet", Type: "types.EdnsClientSubnetType", Required: false},
	{Name: "FirewallRulesFailOpen", Flag: "firewall-rules-fail-open", Type: "types.FirewallRulesFailOpenType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_firewall_domains = []leanruntime.Field{
	{Name: "Domains", Flag: "domains", Type: "[]string", Required: true},
	{Name: "FirewallDomainListId", Flag: "firewall-domain-list-id", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: true},
}

var fields_update_firewall_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FirewallRuleAction", Required: false},
	{Name: "BlockOverrideDnsType", Flag: "block-override-dns-type", Type: "types.BlockOverrideDnsQueryType", Required: false},
	{Name: "BlockOverrideDomain", Flag: "block-override-domain", Type: "*string", Required: false},
	{Name: "BlockOverrideTtl", Flag: "block-override-ttl", Type: "*int32", Required: false},
	{Name: "BlockResponse", Flag: "block-response", Type: "types.FirewallBlockResponse", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConfidenceThreshold", Flag: "confidence-threshold", Type: "types.ConfidenceThreshold", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DnsAdvancedProtection", Flag: "dns-advanced-protection", Type: "types.DnsAdvancedProtection", Required: false},
	{Name: "FirewallRuleId", Flag: "firewall-rule-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int64", Required: false},
}

var fields_update_global_resolver = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalResolverId", Flag: "global-resolver-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ObservabilityRegion", Flag: "observability-region", Type: "*string", Required: false},
}

var fields_update_hosted_zone_association = []leanruntime.Field{
	{Name: "HostedZoneAssociationId", Flag: "hosted-zone-association-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-hosted-zone": {
			Name:   "associate-hosted-zone",
			Fields: fields_associate_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateHostedZone(ctx, input)
			},
		},
		"batch-create-firewall-rule": {
			Name:   "batch-create-firewall-rule",
			Fields: fields_batch_create_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateFirewallRule(ctx, input)
			},
		},
		"batch-delete-firewall-rule": {
			Name:   "batch-delete-firewall-rule",
			Fields: fields_batch_delete_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteFirewallRule(ctx, input)
			},
		},
		"batch-update-firewall-rule": {
			Name:   "batch-update-firewall-rule",
			Fields: fields_batch_update_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateFirewallRule(ctx, input)
			},
		},
		"create-access-source": {
			Name:   "create-access-source",
			Fields: fields_create_access_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessSource(ctx, input)
			},
		},
		"create-access-token": {
			Name:   "create-access-token",
			Fields: fields_create_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessToken(ctx, input)
			},
		},
		"create-dns-view": {
			Name:   "create-dns-view",
			Fields: fields_create_dns_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDNSViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dns_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDNSView(ctx, input)
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
		"create-global-resolver": {
			Name:   "create-global-resolver",
			Fields: fields_create_global_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlobalResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_global_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlobalResolver(ctx, input)
			},
		},
		"delete-access-source": {
			Name:   "delete-access-source",
			Fields: fields_delete_access_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessSource(ctx, input)
			},
		},
		"delete-access-token": {
			Name:   "delete-access-token",
			Fields: fields_delete_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessToken(ctx, input)
			},
		},
		"delete-dns-view": {
			Name:   "delete-dns-view",
			Fields: fields_delete_dns_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDNSViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dns_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDNSView(ctx, input)
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
		"delete-global-resolver": {
			Name:   "delete-global-resolver",
			Fields: fields_delete_global_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlobalResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_global_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlobalResolver(ctx, input)
			},
		},
		"disable-dns-view": {
			Name:   "disable-dns-view",
			Fields: fields_disable_dns_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDNSViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_dns_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDNSView(ctx, input)
			},
		},
		"disassociate-hosted-zone": {
			Name:   "disassociate-hosted-zone",
			Fields: fields_disassociate_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateHostedZone(ctx, input)
			},
		},
		"enable-dns-view": {
			Name:   "enable-dns-view",
			Fields: fields_enable_dns_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDNSViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_dns_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDNSView(ctx, input)
			},
		},
		"get-access-source": {
			Name:   "get-access-source",
			Fields: fields_get_access_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessSource(ctx, input)
			},
		},
		"get-access-token": {
			Name:   "get-access-token",
			Fields: fields_get_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessToken(ctx, input)
			},
		},
		"get-dns-view": {
			Name:   "get-dns-view",
			Fields: fields_get_dns_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDNSViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dns_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDNSView(ctx, input)
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
		"get-firewall-rule": {
			Name:   "get-firewall-rule",
			Fields: fields_get_firewall_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFirewallRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_firewall_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFirewallRule(ctx, input)
			},
		},
		"get-global-resolver": {
			Name:   "get-global-resolver",
			Fields: fields_get_global_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlobalResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_global_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlobalResolver(ctx, input)
			},
		},
		"get-hosted-zone-association": {
			Name:   "get-hosted-zone-association",
			Fields: fields_get_hosted_zone_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostedZoneAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hosted_zone_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHostedZoneAssociation(ctx, input)
			},
		},
		"get-managed-firewall-domain-list": {
			Name:   "get-managed-firewall-domain-list",
			Fields: fields_get_managed_firewall_domain_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedFirewallDomainListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_firewall_domain_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedFirewallDomainList(ctx, input)
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
		"list-access-sources": {
			Name:   "list-access-sources",
			Fields: fields_list_access_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessSources(ctx, input)
				}
				var results []*svc.ListAccessSourcesOutput
				p := svc.NewListAccessSourcesPaginator(client, input)
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
		"list-access-tokens": {
			Name:   "list-access-tokens",
			Fields: fields_list_access_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessTokensInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_tokens, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessTokens(ctx, input)
				}
				var results []*svc.ListAccessTokensOutput
				p := svc.NewListAccessTokensPaginator(client, input)
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
		"list-dns-views": {
			Name:   "list-dns-views",
			Fields: fields_list_dns_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDNSViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dns_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDNSViews(ctx, input)
				}
				var results []*svc.ListDNSViewsOutput
				p := svc.NewListDNSViewsPaginator(client, input)
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
		"list-global-resolvers": {
			Name:   "list-global-resolvers",
			Fields: fields_list_global_resolvers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGlobalResolversInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_global_resolvers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGlobalResolvers(ctx, input)
				}
				var results []*svc.ListGlobalResolversOutput
				p := svc.NewListGlobalResolversPaginator(client, input)
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
		"list-hosted-zone-associations": {
			Name:   "list-hosted-zone-associations",
			Fields: fields_list_hosted_zone_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostedZoneAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hosted_zone_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHostedZoneAssociations(ctx, input)
				}
				var results []*svc.ListHostedZoneAssociationsOutput
				p := svc.NewListHostedZoneAssociationsPaginator(client, input)
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
		"list-managed-firewall-domain-lists": {
			Name:   "list-managed-firewall-domain-lists",
			Fields: fields_list_managed_firewall_domain_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedFirewallDomainListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_firewall_domain_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedFirewallDomainLists(ctx, input)
				}
				var results []*svc.ListManagedFirewallDomainListsOutput
				p := svc.NewListManagedFirewallDomainListsPaginator(client, input)
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
		"update-access-source": {
			Name:   "update-access-source",
			Fields: fields_update_access_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessSource(ctx, input)
			},
		},
		"update-access-token": {
			Name:   "update-access-token",
			Fields: fields_update_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessToken(ctx, input)
			},
		},
		"update-dns-view": {
			Name:   "update-dns-view",
			Fields: fields_update_dns_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDNSViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dns_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDNSView(ctx, input)
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
		"update-global-resolver": {
			Name:   "update-global-resolver",
			Fields: fields_update_global_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalResolver(ctx, input)
			},
		},
		"update-hosted-zone-association": {
			Name:   "update-hosted-zone-association",
			Fields: fields_update_hosted_zone_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHostedZoneAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hosted_zone_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHostedZoneAssociation(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53globalresolver", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
