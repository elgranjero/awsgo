package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/grafana"
)

var fields_associate_license = []leanruntime.Field{
	{Name: "GrafanaToken", Flag: "grafana-token", Type: "*string", Required: false},
	{Name: "LicenseType", Flag: "license-type", Type: "types.LicenseType", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_workspace = []leanruntime.Field{
	{Name: "AccountAccessType", Flag: "account-access-type", Type: "types.AccountAccessType", Required: true},
	{Name: "AuthenticationProviders", Flag: "authentication-providers", Type: "[]types.AuthenticationProviderTypes", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: false},
	{Name: "GrafanaVersion", Flag: "grafana-version", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "NetworkAccessControl", Flag: "network-access-control", Type: "*types.NetworkAccessConfiguration", Required: false},
	{Name: "OrganizationRoleName", Flag: "organization-role-name", Type: "*string", Required: false},
	{Name: "PermissionType", Flag: "permission-type", Type: "types.PermissionType", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: false},
	{Name: "WorkspaceDataSources", Flag: "workspace-data-sources", Type: "[]types.DataSourceType", Required: false},
	{Name: "WorkspaceDescription", Flag: "workspace-description", Type: "*string", Required: false},
	{Name: "WorkspaceName", Flag: "workspace-name", Type: "*string", Required: false},
	{Name: "WorkspaceNotificationDestinations", Flag: "workspace-notification-destinations", Type: "[]types.NotificationDestinationType", Required: false},
	{Name: "WorkspaceOrganizationalUnits", Flag: "workspace-organizational-units", Type: "[]string", Required: false},
	{Name: "WorkspaceRoleArn", Flag: "workspace-role-arn", Type: "*string", Required: false},
}

var fields_create_workspace_api_key = []leanruntime.Field{
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "KeyRole", Flag: "key-role", Type: "*string", Required: true},
	{Name: "SecondsToLive", Flag: "seconds-to-live", Type: "*int32", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_workspace_service_account = []leanruntime.Field{
	{Name: "GrafanaRole", Flag: "grafana-role", Type: "types.Role", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_workspace_service_account_token = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecondsToLive", Flag: "seconds-to-live", Type: "*int32", Required: true},
	{Name: "ServiceAccountId", Flag: "service-account-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace_api_key = []leanruntime.Field{
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace_service_account = []leanruntime.Field{
	{Name: "ServiceAccountId", Flag: "service-account-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace_service_account_token = []leanruntime.Field{
	{Name: "ServiceAccountId", Flag: "service-account-id", Type: "*string", Required: true},
	{Name: "TokenId", Flag: "token-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_workspace = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_workspace_authentication = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_workspace_configuration = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_disassociate_license = []leanruntime.Field{
	{Name: "LicenseType", Flag: "license-type", Type: "types.LicenseType", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_permissions = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
	{Name: "UserType", Flag: "user-type", Type: "types.UserType", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: false},
}

var fields_list_workspace_service_account_tokens = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceAccountId", Flag: "service-account-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_workspace_service_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_workspaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_permissions = []leanruntime.Field{
	{Name: "UpdateInstructionBatch", Flag: "update-instruction-batch", Type: "[]types.UpdateInstruction", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace = []leanruntime.Field{
	{Name: "AccountAccessType", Flag: "account-access-type", Type: "types.AccountAccessType", Required: false},
	{Name: "NetworkAccessControl", Flag: "network-access-control", Type: "*types.NetworkAccessConfiguration", Required: false},
	{Name: "OrganizationRoleName", Flag: "organization-role-name", Type: "*string", Required: false},
	{Name: "PermissionType", Flag: "permission-type", Type: "types.PermissionType", Required: false},
	{Name: "RemoveNetworkAccessConfiguration", Flag: "remove-network-access-configuration", Type: "*bool", Required: false},
	{Name: "RemoveVpcConfiguration", Flag: "remove-vpc-configuration", Type: "*bool", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: false},
	{Name: "WorkspaceDataSources", Flag: "workspace-data-sources", Type: "[]types.DataSourceType", Required: false},
	{Name: "WorkspaceDescription", Flag: "workspace-description", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
	{Name: "WorkspaceName", Flag: "workspace-name", Type: "*string", Required: false},
	{Name: "WorkspaceNotificationDestinations", Flag: "workspace-notification-destinations", Type: "[]types.NotificationDestinationType", Required: false},
	{Name: "WorkspaceOrganizationalUnits", Flag: "workspace-organizational-units", Type: "[]string", Required: false},
	{Name: "WorkspaceRoleArn", Flag: "workspace-role-arn", Type: "*string", Required: false},
}

var fields_update_workspace_authentication = []leanruntime.Field{
	{Name: "AuthenticationProviders", Flag: "authentication-providers", Type: "[]types.AuthenticationProviderTypes", Required: true},
	{Name: "SamlConfiguration", Flag: "saml-configuration", Type: "*types.SamlConfiguration", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: true},
	{Name: "GrafanaVersion", Flag: "grafana-version", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-license": {
			Name:   "associate-license",
			Fields: fields_associate_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateLicense(ctx, input)
			},
		},
		"create-workspace": {
			Name:   "create-workspace",
			Fields: fields_create_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspace(ctx, input)
			},
		},
		"create-workspace-api-key": {
			Name:   "create-workspace-api-key",
			Fields: fields_create_workspace_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaceApiKey(ctx, input)
			},
		},
		"create-workspace-service-account": {
			Name:   "create-workspace-service-account",
			Fields: fields_create_workspace_service_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceServiceAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_service_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaceServiceAccount(ctx, input)
			},
		},
		"create-workspace-service-account-token": {
			Name:   "create-workspace-service-account-token",
			Fields: fields_create_workspace_service_account_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceServiceAccountTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_service_account_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaceServiceAccountToken(ctx, input)
			},
		},
		"delete-workspace": {
			Name:   "delete-workspace",
			Fields: fields_delete_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspace(ctx, input)
			},
		},
		"delete-workspace-api-key": {
			Name:   "delete-workspace-api-key",
			Fields: fields_delete_workspace_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceApiKey(ctx, input)
			},
		},
		"delete-workspace-service-account": {
			Name:   "delete-workspace-service-account",
			Fields: fields_delete_workspace_service_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceServiceAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_service_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceServiceAccount(ctx, input)
			},
		},
		"delete-workspace-service-account-token": {
			Name:   "delete-workspace-service-account-token",
			Fields: fields_delete_workspace_service_account_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceServiceAccountTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_service_account_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceServiceAccountToken(ctx, input)
			},
		},
		"describe-workspace": {
			Name:   "describe-workspace",
			Fields: fields_describe_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspace(ctx, input)
			},
		},
		"describe-workspace-authentication": {
			Name:   "describe-workspace-authentication",
			Fields: fields_describe_workspace_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceAuthentication(ctx, input)
			},
		},
		"describe-workspace-configuration": {
			Name:   "describe-workspace-configuration",
			Fields: fields_describe_workspace_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceConfiguration(ctx, input)
			},
		},
		"disassociate-license": {
			Name:   "disassociate-license",
			Fields: fields_disassociate_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateLicense(ctx, input)
			},
		},
		"list-permissions": {
			Name:   "list-permissions",
			Fields: fields_list_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissions(ctx, input)
				}
				var results []*svc.ListPermissionsOutput
				p := svc.NewListPermissionsPaginator(client, input)
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
		"list-versions": {
			Name:   "list-versions",
			Fields: fields_list_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVersions(ctx, input)
				}
				var results []*svc.ListVersionsOutput
				p := svc.NewListVersionsPaginator(client, input)
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
		"list-workspace-service-account-tokens": {
			Name:   "list-workspace-service-account-tokens",
			Fields: fields_list_workspace_service_account_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspaceServiceAccountTokensInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspace_service_account_tokens, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspaceServiceAccountTokens(ctx, input)
				}
				var results []*svc.ListWorkspaceServiceAccountTokensOutput
				p := svc.NewListWorkspaceServiceAccountTokensPaginator(client, input)
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
		"list-workspace-service-accounts": {
			Name:   "list-workspace-service-accounts",
			Fields: fields_list_workspace_service_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspaceServiceAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspace_service_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspaceServiceAccounts(ctx, input)
				}
				var results []*svc.ListWorkspaceServiceAccountsOutput
				p := svc.NewListWorkspaceServiceAccountsPaginator(client, input)
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
		"list-workspaces": {
			Name:   "list-workspaces",
			Fields: fields_list_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspaces(ctx, input)
				}
				var results []*svc.ListWorkspacesOutput
				p := svc.NewListWorkspacesPaginator(client, input)
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
		"update-permissions": {
			Name:   "update-permissions",
			Fields: fields_update_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePermissions(ctx, input)
			},
		},
		"update-workspace": {
			Name:   "update-workspace",
			Fields: fields_update_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspace(ctx, input)
			},
		},
		"update-workspace-authentication": {
			Name:   "update-workspace-authentication",
			Fields: fields_update_workspace_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceAuthentication(ctx, input)
			},
		},
		"update-workspace-configuration": {
			Name:   "update-workspace-configuration",
			Fields: fields_update_workspace_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("grafana", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
