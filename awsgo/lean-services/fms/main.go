package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/fms"
)

var fields_associate_admin_account = []leanruntime.Field{
	{Name: "AdminAccount", Flag: "admin-account", Type: "*string", Required: true},
}

var fields_associate_third_party_firewall = []leanruntime.Field{
	{Name: "ThirdPartyFirewall", Flag: "third-party-firewall", Type: "types.ThirdPartyFirewall", Required: true},
}

var fields_batch_associate_resource = []leanruntime.Field{
	{Name: "Items", Flag: "items", Type: "[]string", Required: true},
	{Name: "ResourceSetIdentifier", Flag: "resource-set-identifier", Type: "*string", Required: true},
}

var fields_batch_disassociate_resource = []leanruntime.Field{
	{Name: "Items", Flag: "items", Type: "[]string", Required: true},
	{Name: "ResourceSetIdentifier", Flag: "resource-set-identifier", Type: "*string", Required: true},
}

var fields_delete_apps_list = []leanruntime.Field{
	{Name: "ListId", Flag: "list-id", Type: "*string", Required: true},
}

var fields_delete_notification_channel = []leanruntime.Field{}

var fields_delete_policy = []leanruntime.Field{
	{Name: "DeleteAllPolicyResources", Flag: "delete-all-policy-resources", Type: "bool", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_delete_protocols_list = []leanruntime.Field{
	{Name: "ListId", Flag: "list-id", Type: "*string", Required: true},
}

var fields_delete_resource_set = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_disassociate_admin_account = []leanruntime.Field{}

var fields_disassociate_third_party_firewall = []leanruntime.Field{
	{Name: "ThirdPartyFirewall", Flag: "third-party-firewall", Type: "types.ThirdPartyFirewall", Required: true},
}

var fields_get_admin_account = []leanruntime.Field{}

var fields_get_admin_scope = []leanruntime.Field{
	{Name: "AdminAccount", Flag: "admin-account", Type: "*string", Required: true},
}

var fields_get_apps_list = []leanruntime.Field{
	{Name: "DefaultList", Flag: "default-list", Type: "bool", Required: false},
	{Name: "ListId", Flag: "list-id", Type: "*string", Required: true},
}

var fields_get_compliance_detail = []leanruntime.Field{
	{Name: "MemberAccount", Flag: "member-account", Type: "*string", Required: true},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_get_notification_channel = []leanruntime.Field{}

var fields_get_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_get_protection_status = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberAccountId", Flag: "member-account-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_protocols_list = []leanruntime.Field{
	{Name: "DefaultList", Flag: "default-list", Type: "bool", Required: false},
	{Name: "ListId", Flag: "list-id", Type: "*string", Required: true},
}

var fields_get_resource_set = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_third_party_firewall_association_status = []leanruntime.Field{
	{Name: "ThirdPartyFirewall", Flag: "third-party-firewall", Type: "types.ThirdPartyFirewall", Required: true},
}

var fields_get_violation_details = []leanruntime.Field{
	{Name: "MemberAccount", Flag: "member-account", Type: "*string", Required: true},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_list_admin_accounts_for_organization = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_admins_managing_account = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_apps_lists = []leanruntime.Field{
	{Name: "DefaultLists", Flag: "default-lists", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_compliance_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_list_discovered_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberAccountIds", Flag: "member-account-ids", Type: "[]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_list_member_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_protocols_lists = []leanruntime.Field{
	{Name: "DefaultLists", Flag: "default-lists", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_set_resources = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_sets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_third_party_firewall_firewall_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThirdPartyFirewall", Flag: "third-party-firewall", Type: "types.ThirdPartyFirewall", Required: true},
}

var fields_put_admin_account = []leanruntime.Field{
	{Name: "AdminAccount", Flag: "admin-account", Type: "*string", Required: true},
	{Name: "AdminScope", Flag: "admin-scope", Type: "*types.AdminScope", Required: false},
}

var fields_put_apps_list = []leanruntime.Field{
	{Name: "AppsList", Flag: "apps-list", Type: "*types.AppsListData", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_put_notification_channel = []leanruntime.Field{
	{Name: "SnsRoleName", Flag: "sns-role-name", Type: "*string", Required: true},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: true},
}

var fields_put_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*types.Policy", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_put_protocols_list = []leanruntime.Field{
	{Name: "ProtocolsList", Flag: "protocols-list", Type: "*types.ProtocolsListData", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_put_resource_set = []leanruntime.Field{
	{Name: "ResourceSet", Flag: "resource-set", Type: "*types.ResourceSet", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-admin-account": {
			Name:   "associate-admin-account",
			Fields: fields_associate_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAdminAccount(ctx, input)
			},
		},
		"associate-third-party-firewall": {
			Name:   "associate-third-party-firewall",
			Fields: fields_associate_third_party_firewall,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateThirdPartyFirewallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_third_party_firewall, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateThirdPartyFirewall(ctx, input)
			},
		},
		"batch-associate-resource": {
			Name:   "batch-associate-resource",
			Fields: fields_batch_associate_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateResource(ctx, input)
			},
		},
		"batch-disassociate-resource": {
			Name:   "batch-disassociate-resource",
			Fields: fields_batch_disassociate_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateResource(ctx, input)
			},
		},
		"delete-apps-list": {
			Name:   "delete-apps-list",
			Fields: fields_delete_apps_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppsListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_apps_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppsList(ctx, input)
			},
		},
		"delete-notification-channel": {
			Name:   "delete-notification-channel",
			Fields: fields_delete_notification_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotificationChannel(ctx, input)
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
		"delete-protocols-list": {
			Name:   "delete-protocols-list",
			Fields: fields_delete_protocols_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProtocolsListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_protocols_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProtocolsList(ctx, input)
			},
		},
		"delete-resource-set": {
			Name:   "delete-resource-set",
			Fields: fields_delete_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceSet(ctx, input)
			},
		},
		"disassociate-admin-account": {
			Name:   "disassociate-admin-account",
			Fields: fields_disassociate_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAdminAccount(ctx, input)
			},
		},
		"disassociate-third-party-firewall": {
			Name:   "disassociate-third-party-firewall",
			Fields: fields_disassociate_third_party_firewall,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateThirdPartyFirewallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_third_party_firewall, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateThirdPartyFirewall(ctx, input)
			},
		},
		"get-admin-account": {
			Name:   "get-admin-account",
			Fields: fields_get_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAdminAccount(ctx, input)
			},
		},
		"get-admin-scope": {
			Name:   "get-admin-scope",
			Fields: fields_get_admin_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAdminScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_admin_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAdminScope(ctx, input)
			},
		},
		"get-apps-list": {
			Name:   "get-apps-list",
			Fields: fields_get_apps_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppsListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apps_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAppsList(ctx, input)
			},
		},
		"get-compliance-detail": {
			Name:   "get-compliance-detail",
			Fields: fields_get_compliance_detail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComplianceDetailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compliance_detail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComplianceDetail(ctx, input)
			},
		},
		"get-notification-channel": {
			Name:   "get-notification-channel",
			Fields: fields_get_notification_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotificationChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notification_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotificationChannel(ctx, input)
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
		"get-protection-status": {
			Name:   "get-protection-status",
			Fields: fields_get_protection_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProtectionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_protection_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProtectionStatus(ctx, input)
			},
		},
		"get-protocols-list": {
			Name:   "get-protocols-list",
			Fields: fields_get_protocols_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProtocolsListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_protocols_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProtocolsList(ctx, input)
			},
		},
		"get-resource-set": {
			Name:   "get-resource-set",
			Fields: fields_get_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceSet(ctx, input)
			},
		},
		"get-third-party-firewall-association-status": {
			Name:   "get-third-party-firewall-association-status",
			Fields: fields_get_third_party_firewall_association_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThirdPartyFirewallAssociationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_third_party_firewall_association_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThirdPartyFirewallAssociationStatus(ctx, input)
			},
		},
		"get-violation-details": {
			Name:   "get-violation-details",
			Fields: fields_get_violation_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetViolationDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_violation_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetViolationDetails(ctx, input)
			},
		},
		"list-admin-accounts-for-organization": {
			Name:   "list-admin-accounts-for-organization",
			Fields: fields_list_admin_accounts_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAdminAccountsForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_admin_accounts_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAdminAccountsForOrganization(ctx, input)
				}
				var results []*svc.ListAdminAccountsForOrganizationOutput
				p := svc.NewListAdminAccountsForOrganizationPaginator(client, input)
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
		"list-admins-managing-account": {
			Name:   "list-admins-managing-account",
			Fields: fields_list_admins_managing_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAdminsManagingAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_admins_managing_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAdminsManagingAccount(ctx, input)
				}
				var results []*svc.ListAdminsManagingAccountOutput
				p := svc.NewListAdminsManagingAccountPaginator(client, input)
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
		"list-apps-lists": {
			Name:   "list-apps-lists",
			Fields: fields_list_apps_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppsListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_apps_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppsLists(ctx, input)
				}
				var results []*svc.ListAppsListsOutput
				p := svc.NewListAppsListsPaginator(client, input)
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
		"list-compliance-status": {
			Name:   "list-compliance-status",
			Fields: fields_list_compliance_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComplianceStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compliance_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComplianceStatus(ctx, input)
				}
				var results []*svc.ListComplianceStatusOutput
				p := svc.NewListComplianceStatusPaginator(client, input)
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
		"list-discovered-resources": {
			Name:   "list-discovered-resources",
			Fields: fields_list_discovered_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDiscoveredResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_discovered_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDiscoveredResources(ctx, input)
			},
		},
		"list-member-accounts": {
			Name:   "list-member-accounts",
			Fields: fields_list_member_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMemberAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_member_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMemberAccounts(ctx, input)
				}
				var results []*svc.ListMemberAccountsOutput
				p := svc.NewListMemberAccountsPaginator(client, input)
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
		"list-protocols-lists": {
			Name:   "list-protocols-lists",
			Fields: fields_list_protocols_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtocolsListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protocols_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtocolsLists(ctx, input)
				}
				var results []*svc.ListProtocolsListsOutput
				p := svc.NewListProtocolsListsPaginator(client, input)
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
		"list-resource-set-resources": {
			Name:   "list-resource-set-resources",
			Fields: fields_list_resource_set_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceSetResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resource_set_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourceSetResources(ctx, input)
			},
		},
		"list-resource-sets": {
			Name:   "list-resource-sets",
			Fields: fields_list_resource_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resource_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourceSets(ctx, input)
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
		"list-third-party-firewall-firewall-policies": {
			Name:   "list-third-party-firewall-firewall-policies",
			Fields: fields_list_third_party_firewall_firewall_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThirdPartyFirewallFirewallPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_third_party_firewall_firewall_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThirdPartyFirewallFirewallPolicies(ctx, input)
				}
				var results []*svc.ListThirdPartyFirewallFirewallPoliciesOutput
				p := svc.NewListThirdPartyFirewallFirewallPoliciesPaginator(client, input)
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
		"put-admin-account": {
			Name:   "put-admin-account",
			Fields: fields_put_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAdminAccount(ctx, input)
			},
		},
		"put-apps-list": {
			Name:   "put-apps-list",
			Fields: fields_put_apps_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAppsListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_apps_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAppsList(ctx, input)
			},
		},
		"put-notification-channel": {
			Name:   "put-notification-channel",
			Fields: fields_put_notification_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutNotificationChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_notification_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutNotificationChannel(ctx, input)
			},
		},
		"put-policy": {
			Name:   "put-policy",
			Fields: fields_put_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPolicy(ctx, input)
			},
		},
		"put-protocols-list": {
			Name:   "put-protocols-list",
			Fields: fields_put_protocols_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProtocolsListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_protocols_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProtocolsList(ctx, input)
			},
		},
		"put-resource-set": {
			Name:   "put-resource-set",
			Fields: fields_put_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourceSet(ctx, input)
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
	}
	if err := leanruntime.Execute("fms", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
