package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/supportapp"
)

var fields_create_slack_channel_configuration = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
	{Name: "ChannelRoleArn", Flag: "channel-role-arn", Type: "*string", Required: true},
	{Name: "NotifyOnAddCorrespondenceToCase", Flag: "notify-on-add-correspondence-to-case", Type: "*bool", Required: false},
	{Name: "NotifyOnCaseSeverity", Flag: "notify-on-case-severity", Type: "types.NotificationSeverityLevel", Required: true},
	{Name: "NotifyOnCreateOrReopenCase", Flag: "notify-on-create-or-reopen-case", Type: "*bool", Required: false},
	{Name: "NotifyOnResolveCase", Flag: "notify-on-resolve-case", Type: "*bool", Required: false},
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
}

var fields_delete_account_alias = []leanruntime.Field{}

var fields_delete_slack_channel_configuration = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
}

var fields_delete_slack_workspace_configuration = []leanruntime.Field{
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
}

var fields_get_account_alias = []leanruntime.Field{}

var fields_list_slack_channel_configurations = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_slack_workspace_configurations = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_account_alias = []leanruntime.Field{
	{Name: "AccountAlias", Flag: "account-alias", Type: "*string", Required: true},
}

var fields_register_slack_workspace_for_organization = []leanruntime.Field{
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
}

var fields_update_slack_channel_configuration = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
	{Name: "ChannelRoleArn", Flag: "channel-role-arn", Type: "*string", Required: false},
	{Name: "NotifyOnAddCorrespondenceToCase", Flag: "notify-on-add-correspondence-to-case", Type: "*bool", Required: false},
	{Name: "NotifyOnCaseSeverity", Flag: "notify-on-case-severity", Type: "types.NotificationSeverityLevel", Required: false},
	{Name: "NotifyOnCreateOrReopenCase", Flag: "notify-on-create-or-reopen-case", Type: "*bool", Required: false},
	{Name: "NotifyOnResolveCase", Flag: "notify-on-resolve-case", Type: "*bool", Required: false},
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-slack-channel-configuration": {
			Name:   "create-slack-channel-configuration",
			Fields: fields_create_slack_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSlackChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_slack_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSlackChannelConfiguration(ctx, input)
			},
		},
		"delete-account-alias": {
			Name:   "delete-account-alias",
			Fields: fields_delete_account_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountAlias(ctx, input)
			},
		},
		"delete-slack-channel-configuration": {
			Name:   "delete-slack-channel-configuration",
			Fields: fields_delete_slack_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlackChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slack_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlackChannelConfiguration(ctx, input)
			},
		},
		"delete-slack-workspace-configuration": {
			Name:   "delete-slack-workspace-configuration",
			Fields: fields_delete_slack_workspace_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlackWorkspaceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slack_workspace_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlackWorkspaceConfiguration(ctx, input)
			},
		},
		"get-account-alias": {
			Name:   "get-account-alias",
			Fields: fields_get_account_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountAlias(ctx, input)
			},
		},
		"list-slack-channel-configurations": {
			Name:   "list-slack-channel-configurations",
			Fields: fields_list_slack_channel_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSlackChannelConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_slack_channel_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSlackChannelConfigurations(ctx, input)
				}
				var results []*svc.ListSlackChannelConfigurationsOutput
				p := svc.NewListSlackChannelConfigurationsPaginator(client, input)
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
		"list-slack-workspace-configurations": {
			Name:   "list-slack-workspace-configurations",
			Fields: fields_list_slack_workspace_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSlackWorkspaceConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_slack_workspace_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSlackWorkspaceConfigurations(ctx, input)
				}
				var results []*svc.ListSlackWorkspaceConfigurationsOutput
				p := svc.NewListSlackWorkspaceConfigurationsPaginator(client, input)
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
		"put-account-alias": {
			Name:   "put-account-alias",
			Fields: fields_put_account_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountAlias(ctx, input)
			},
		},
		"register-slack-workspace-for-organization": {
			Name:   "register-slack-workspace-for-organization",
			Fields: fields_register_slack_workspace_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterSlackWorkspaceForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_slack_workspace_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterSlackWorkspaceForOrganization(ctx, input)
			},
		},
		"update-slack-channel-configuration": {
			Name:   "update-slack-channel-configuration",
			Fields: fields_update_slack_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSlackChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_slack_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSlackChannelConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("supportapp", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
