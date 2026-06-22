package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/detective"
)

var fields_accept_invitation = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_batch_get_graph_member_datasources = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_batch_get_membership_datasources = []leanruntime.Field{
	{Name: "GraphArns", Flag: "graph-arns", Type: "[]string", Required: true},
}

var fields_create_graph = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_members = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]types.Account", Required: true},
	{Name: "DisableEmailNotification", Flag: "disable-email-notification", Type: "bool", Required: false},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "Message", Flag: "message", Type: "*string", Required: false},
}

var fields_delete_graph = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_delete_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_describe_organization_configuration = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_disable_organization_admin_account = []leanruntime.Field{}

var fields_disassociate_membership = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_enable_organization_admin_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_investigation = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "InvestigationId", Flag: "investigation-id", Type: "*string", Required: true},
}

var fields_get_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_list_datasource_packages = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_graphs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_indicators = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "IndicatorType", Flag: "indicator-type", Type: "types.IndicatorType", Required: false},
	{Name: "InvestigationId", Flag: "investigation-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_investigations = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_list_invitations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_organization_admin_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reject_invitation = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_start_investigation = []leanruntime.Field{
	{Name: "EntityArn", Flag: "entity-arn", Type: "*string", Required: true},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "ScopeEndTime", Flag: "scope-end-time", Type: "*time.Time", Required: true},
	{Name: "ScopeStartTime", Flag: "scope-start-time", Type: "*time.Time", Required: true},
}

var fields_start_monitoring_member = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_datasource_packages = []leanruntime.Field{
	{Name: "DatasourcePackages", Flag: "datasource-packages", Type: "[]types.DatasourcePackage", Required: true},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

var fields_update_investigation_state = []leanruntime.Field{
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
	{Name: "InvestigationId", Flag: "investigation-id", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.State", Required: true},
}

var fields_update_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnable", Flag: "auto-enable", Type: "bool", Required: false},
	{Name: "GraphArn", Flag: "graph-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-invitation": {
			Name:   "accept-invitation",
			Fields: fields_accept_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptInvitation(ctx, input)
			},
		},
		"batch-get-graph-member-datasources": {
			Name:   "batch-get-graph-member-datasources",
			Fields: fields_batch_get_graph_member_datasources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetGraphMemberDatasourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_graph_member_datasources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetGraphMemberDatasources(ctx, input)
			},
		},
		"batch-get-membership-datasources": {
			Name:   "batch-get-membership-datasources",
			Fields: fields_batch_get_membership_datasources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetMembershipDatasourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_membership_datasources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetMembershipDatasources(ctx, input)
			},
		},
		"create-graph": {
			Name:   "create-graph",
			Fields: fields_create_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGraph(ctx, input)
			},
		},
		"create-members": {
			Name:   "create-members",
			Fields: fields_create_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMembers(ctx, input)
			},
		},
		"delete-graph": {
			Name:   "delete-graph",
			Fields: fields_delete_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGraph(ctx, input)
			},
		},
		"delete-members": {
			Name:   "delete-members",
			Fields: fields_delete_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMembers(ctx, input)
			},
		},
		"describe-organization-configuration": {
			Name:   "describe-organization-configuration",
			Fields: fields_describe_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganizationConfiguration(ctx, input)
			},
		},
		"disable-organization-admin-account": {
			Name:   "disable-organization-admin-account",
			Fields: fields_disable_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableOrganizationAdminAccount(ctx, input)
			},
		},
		"disassociate-membership": {
			Name:   "disassociate-membership",
			Fields: fields_disassociate_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMembership(ctx, input)
			},
		},
		"enable-organization-admin-account": {
			Name:   "enable-organization-admin-account",
			Fields: fields_enable_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableOrganizationAdminAccount(ctx, input)
			},
		},
		"get-investigation": {
			Name:   "get-investigation",
			Fields: fields_get_investigation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvestigationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_investigation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvestigation(ctx, input)
			},
		},
		"get-members": {
			Name:   "get-members",
			Fields: fields_get_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMembers(ctx, input)
			},
		},
		"list-datasource-packages": {
			Name:   "list-datasource-packages",
			Fields: fields_list_datasource_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasourcePackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datasource_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasourcePackages(ctx, input)
				}
				var results []*svc.ListDatasourcePackagesOutput
				p := svc.NewListDatasourcePackagesPaginator(client, input)
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
		"list-graphs": {
			Name:   "list-graphs",
			Fields: fields_list_graphs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGraphsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_graphs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGraphs(ctx, input)
				}
				var results []*svc.ListGraphsOutput
				p := svc.NewListGraphsPaginator(client, input)
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
		"list-indicators": {
			Name:   "list-indicators",
			Fields: fields_list_indicators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndicatorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_indicators, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIndicators(ctx, input)
			},
		},
		"list-investigations": {
			Name:   "list-investigations",
			Fields: fields_list_investigations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvestigationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_investigations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListInvestigations(ctx, input)
			},
		},
		"list-invitations": {
			Name:   "list-invitations",
			Fields: fields_list_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvitations(ctx, input)
				}
				var results []*svc.ListInvitationsOutput
				p := svc.NewListInvitationsPaginator(client, input)
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
		"list-members": {
			Name:   "list-members",
			Fields: fields_list_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMembers(ctx, input)
				}
				var results []*svc.ListMembersOutput
				p := svc.NewListMembersPaginator(client, input)
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
		"list-organization-admin-accounts": {
			Name:   "list-organization-admin-accounts",
			Fields: fields_list_organization_admin_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationAdminAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_admin_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationAdminAccounts(ctx, input)
				}
				var results []*svc.ListOrganizationAdminAccountsOutput
				p := svc.NewListOrganizationAdminAccountsPaginator(client, input)
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
		"reject-invitation": {
			Name:   "reject-invitation",
			Fields: fields_reject_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectInvitation(ctx, input)
			},
		},
		"start-investigation": {
			Name:   "start-investigation",
			Fields: fields_start_investigation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInvestigationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_investigation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInvestigation(ctx, input)
			},
		},
		"start-monitoring-member": {
			Name:   "start-monitoring-member",
			Fields: fields_start_monitoring_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMonitoringMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_monitoring_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMonitoringMember(ctx, input)
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
		"update-datasource-packages": {
			Name:   "update-datasource-packages",
			Fields: fields_update_datasource_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasourcePackagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_datasource_packages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDatasourcePackages(ctx, input)
			},
		},
		"update-investigation-state": {
			Name:   "update-investigation-state",
			Fields: fields_update_investigation_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInvestigationStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_investigation_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInvestigationState(ctx, input)
			},
		},
		"update-organization-configuration": {
			Name:   "update-organization-configuration",
			Fields: fields_update_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOrganizationConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("detective", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
