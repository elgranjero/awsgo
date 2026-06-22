package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mpa"
)

var fields_cancel_session = []leanruntime.Field{
	{Name: "SessionArn", Flag: "session-arn", Type: "*string", Required: true},
}

var fields_create_approval_team = []leanruntime.Field{
	{Name: "ApprovalStrategy", Flag: "approval-strategy", Type: "types.ApprovalStrategy", Required: true},
	{Name: "Approvers", Flag: "approvers", Type: "[]types.ApprovalTeamRequestApprover", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policies", Flag: "policies", Type: "[]types.PolicyReference", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_identity_source = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IdentitySourceParameters", Flag: "identity-source-parameters", Type: "*types.IdentitySourceParameters", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_identity_source = []leanruntime.Field{
	{Name: "IdentitySourceArn", Flag: "identity-source-arn", Type: "*string", Required: true},
}

var fields_delete_inactive_approval_team_version = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_get_approval_team = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_identity_source = []leanruntime.Field{
	{Name: "IdentitySourceArn", Flag: "identity-source-arn", Type: "*string", Required: true},
}

var fields_get_policy_version = []leanruntime.Field{
	{Name: "PolicyVersionArn", Flag: "policy-version-arn", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "SessionArn", Flag: "session-arn", Type: "*string", Required: true},
}

var fields_list_approval_teams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_sources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policy_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_list_resource_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "ApprovalTeamArn", Flag: "approval-team-arn", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_active_approval_team_deletion = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "PendingWindowDays", Flag: "pending-window-days", Type: "*int32", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_approval_team = []leanruntime.Field{
	{Name: "ApprovalStrategy", Flag: "approval-strategy", Type: "types.ApprovalStrategy", Required: false},
	{Name: "Approvers", Flag: "approvers", Type: "[]types.ApprovalTeamRequestApprover", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "UpdateActions", Flag: "update-actions", Type: "[]types.UpdateAction", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-session": {
			Name:   "cancel-session",
			Fields: fields_cancel_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSession(ctx, input)
			},
		},
		"create-approval-team": {
			Name:   "create-approval-team",
			Fields: fields_create_approval_team,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApprovalTeamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_approval_team, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApprovalTeam(ctx, input)
			},
		},
		"create-identity-source": {
			Name:   "create-identity-source",
			Fields: fields_create_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdentitySource(ctx, input)
			},
		},
		"delete-identity-source": {
			Name:   "delete-identity-source",
			Fields: fields_delete_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentitySource(ctx, input)
			},
		},
		"delete-inactive-approval-team-version": {
			Name:   "delete-inactive-approval-team-version",
			Fields: fields_delete_inactive_approval_team_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInactiveApprovalTeamVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inactive_approval_team_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInactiveApprovalTeamVersion(ctx, input)
			},
		},
		"get-approval-team": {
			Name:   "get-approval-team",
			Fields: fields_get_approval_team,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApprovalTeamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_approval_team, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApprovalTeam(ctx, input)
			},
		},
		"get-identity-source": {
			Name:   "get-identity-source",
			Fields: fields_get_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentitySource(ctx, input)
			},
		},
		"get-policy-version": {
			Name:   "get-policy-version",
			Fields: fields_get_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyVersion(ctx, input)
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
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"list-approval-teams": {
			Name:   "list-approval-teams",
			Fields: fields_list_approval_teams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApprovalTeamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_approval_teams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApprovalTeams(ctx, input)
				}
				var results []*svc.ListApprovalTeamsOutput
				p := svc.NewListApprovalTeamsPaginator(client, input)
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
		"list-identity-sources": {
			Name:   "list-identity-sources",
			Fields: fields_list_identity_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentitySourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identity_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentitySources(ctx, input)
				}
				var results []*svc.ListIdentitySourcesOutput
				p := svc.NewListIdentitySourcesPaginator(client, input)
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
		"list-policy-versions": {
			Name:   "list-policy-versions",
			Fields: fields_list_policy_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyVersions(ctx, input)
				}
				var results []*svc.ListPolicyVersionsOutput
				p := svc.NewListPolicyVersionsPaginator(client, input)
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
		"list-resource-policies": {
			Name:   "list-resource-policies",
			Fields: fields_list_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourcePolicies(ctx, input)
				}
				var results []*svc.ListResourcePoliciesOutput
				p := svc.NewListResourcePoliciesPaginator(client, input)
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
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"start-active-approval-team-deletion": {
			Name:   "start-active-approval-team-deletion",
			Fields: fields_start_active_approval_team_deletion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartActiveApprovalTeamDeletionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_active_approval_team_deletion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartActiveApprovalTeamDeletion(ctx, input)
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
		"update-approval-team": {
			Name:   "update-approval-team",
			Fields: fields_update_approval_team,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApprovalTeamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_approval_team, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApprovalTeam(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mpa", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
