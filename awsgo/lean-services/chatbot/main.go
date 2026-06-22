package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chatbot"
)

var fields_associate_to_configuration = []leanruntime.Field{
	{Name: "ChatConfiguration", Flag: "chat-configuration", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
}

var fields_create_chime_webhook_configuration = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "LoggingLevel", Flag: "logging-level", Type: "*string", Required: false},
	{Name: "SnsTopicArns", Flag: "sns-topic-arns", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WebhookDescription", Flag: "webhook-description", Type: "*string", Required: true},
	{Name: "WebhookUrl", Flag: "webhook-url", Type: "*string", Required: true},
}

var fields_create_custom_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "Attachments", Flag: "attachments", Type: "[]types.CustomActionAttachment", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.CustomActionDefinition", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_microsoft_teams_channel_configuration = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
	{Name: "GuardrailPolicyArns", Flag: "guardrail-policy-arns", Type: "[]string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "LoggingLevel", Flag: "logging-level", Type: "*string", Required: false},
	{Name: "SnsTopicArns", Flag: "sns-topic-arns", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
	{Name: "TeamName", Flag: "team-name", Type: "*string", Required: false},
	{Name: "TenantId", Flag: "tenant-id", Type: "*string", Required: true},
	{Name: "UserAuthorizationRequired", Flag: "user-authorization-required", Type: "*bool", Required: false},
}

var fields_create_slack_channel_configuration = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
	{Name: "GuardrailPolicyArns", Flag: "guardrail-policy-arns", Type: "[]string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "LoggingLevel", Flag: "logging-level", Type: "*string", Required: false},
	{Name: "SlackChannelId", Flag: "slack-channel-id", Type: "*string", Required: true},
	{Name: "SlackChannelName", Flag: "slack-channel-name", Type: "*string", Required: false},
	{Name: "SlackTeamId", Flag: "slack-team-id", Type: "*string", Required: true},
	{Name: "SnsTopicArns", Flag: "sns-topic-arns", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserAuthorizationRequired", Flag: "user-authorization-required", Type: "*bool", Required: false},
}

var fields_delete_chime_webhook_configuration = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_custom_action = []leanruntime.Field{
	{Name: "CustomActionArn", Flag: "custom-action-arn", Type: "*string", Required: true},
}

var fields_delete_microsoft_teams_channel_configuration = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_microsoft_teams_configured_team = []leanruntime.Field{
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: true},
}

var fields_delete_microsoft_teams_user_identity = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_slack_channel_configuration = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_slack_user_identity = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
	{Name: "SlackTeamId", Flag: "slack-team-id", Type: "*string", Required: true},
	{Name: "SlackUserId", Flag: "slack-user-id", Type: "*string", Required: true},
}

var fields_delete_slack_workspace_authorization = []leanruntime.Field{
	{Name: "SlackTeamId", Flag: "slack-team-id", Type: "*string", Required: true},
}

var fields_describe_chime_webhook_configurations = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_slack_channel_configurations = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_slack_user_identities = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_slack_workspaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_from_configuration = []leanruntime.Field{
	{Name: "ChatConfiguration", Flag: "chat-configuration", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
}

var fields_get_account_preferences = []leanruntime.Field{}

var fields_get_custom_action = []leanruntime.Field{
	{Name: "CustomActionArn", Flag: "custom-action-arn", Type: "*string", Required: true},
}

var fields_get_microsoft_teams_channel_configuration = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
}

var fields_list_associations = []leanruntime.Field{
	{Name: "ChatConfiguration", Flag: "chat-configuration", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_microsoft_teams_channel_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TeamId", Flag: "team-id", Type: "*string", Required: false},
}

var fields_list_microsoft_teams_configured_teams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_microsoft_teams_user_identities = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_preferences = []leanruntime.Field{
	{Name: "TrainingDataCollectionEnabled", Flag: "training-data-collection-enabled", Type: "*bool", Required: false},
	{Name: "UserAuthorizationRequired", Flag: "user-authorization-required", Type: "*bool", Required: false},
}

var fields_update_chime_webhook_configuration = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "LoggingLevel", Flag: "logging-level", Type: "*string", Required: false},
	{Name: "SnsTopicArns", Flag: "sns-topic-arns", Type: "[]string", Required: false},
	{Name: "WebhookDescription", Flag: "webhook-description", Type: "*string", Required: false},
	{Name: "WebhookUrl", Flag: "webhook-url", Type: "*string", Required: false},
}

var fields_update_custom_action = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "Attachments", Flag: "attachments", Type: "[]types.CustomActionAttachment", Required: false},
	{Name: "CustomActionArn", Flag: "custom-action-arn", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.CustomActionDefinition", Required: true},
}

var fields_update_microsoft_teams_channel_configuration = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
	{Name: "GuardrailPolicyArns", Flag: "guardrail-policy-arns", Type: "[]string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "LoggingLevel", Flag: "logging-level", Type: "*string", Required: false},
	{Name: "SnsTopicArns", Flag: "sns-topic-arns", Type: "[]string", Required: false},
	{Name: "UserAuthorizationRequired", Flag: "user-authorization-required", Type: "*bool", Required: false},
}

var fields_update_slack_channel_configuration = []leanruntime.Field{
	{Name: "ChatConfigurationArn", Flag: "chat-configuration-arn", Type: "*string", Required: true},
	{Name: "GuardrailPolicyArns", Flag: "guardrail-policy-arns", Type: "[]string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "LoggingLevel", Flag: "logging-level", Type: "*string", Required: false},
	{Name: "SlackChannelId", Flag: "slack-channel-id", Type: "*string", Required: true},
	{Name: "SlackChannelName", Flag: "slack-channel-name", Type: "*string", Required: false},
	{Name: "SnsTopicArns", Flag: "sns-topic-arns", Type: "[]string", Required: false},
	{Name: "UserAuthorizationRequired", Flag: "user-authorization-required", Type: "*bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-to-configuration": {
			Name:   "associate-to-configuration",
			Fields: fields_associate_to_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateToConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_to_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateToConfiguration(ctx, input)
			},
		},
		"create-chime-webhook-configuration": {
			Name:   "create-chime-webhook-configuration",
			Fields: fields_create_chime_webhook_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChimeWebhookConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_chime_webhook_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChimeWebhookConfiguration(ctx, input)
			},
		},
		"create-custom-action": {
			Name:   "create-custom-action",
			Fields: fields_create_custom_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomAction(ctx, input)
			},
		},
		"create-microsoft-teams-channel-configuration": {
			Name:   "create-microsoft-teams-channel-configuration",
			Fields: fields_create_microsoft_teams_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMicrosoftTeamsChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_microsoft_teams_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMicrosoftTeamsChannelConfiguration(ctx, input)
			},
		},
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
		"delete-chime-webhook-configuration": {
			Name:   "delete-chime-webhook-configuration",
			Fields: fields_delete_chime_webhook_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChimeWebhookConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_chime_webhook_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChimeWebhookConfiguration(ctx, input)
			},
		},
		"delete-custom-action": {
			Name:   "delete-custom-action",
			Fields: fields_delete_custom_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomAction(ctx, input)
			},
		},
		"delete-microsoft-teams-channel-configuration": {
			Name:   "delete-microsoft-teams-channel-configuration",
			Fields: fields_delete_microsoft_teams_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMicrosoftTeamsChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_microsoft_teams_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMicrosoftTeamsChannelConfiguration(ctx, input)
			},
		},
		"delete-microsoft-teams-configured-team": {
			Name:   "delete-microsoft-teams-configured-team",
			Fields: fields_delete_microsoft_teams_configured_team,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMicrosoftTeamsConfiguredTeamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_microsoft_teams_configured_team, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMicrosoftTeamsConfiguredTeam(ctx, input)
			},
		},
		"delete-microsoft-teams-user-identity": {
			Name:   "delete-microsoft-teams-user-identity",
			Fields: fields_delete_microsoft_teams_user_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMicrosoftTeamsUserIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_microsoft_teams_user_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMicrosoftTeamsUserIdentity(ctx, input)
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
		"delete-slack-user-identity": {
			Name:   "delete-slack-user-identity",
			Fields: fields_delete_slack_user_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlackUserIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slack_user_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlackUserIdentity(ctx, input)
			},
		},
		"delete-slack-workspace-authorization": {
			Name:   "delete-slack-workspace-authorization",
			Fields: fields_delete_slack_workspace_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlackWorkspaceAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slack_workspace_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlackWorkspaceAuthorization(ctx, input)
			},
		},
		"describe-chime-webhook-configurations": {
			Name:   "describe-chime-webhook-configurations",
			Fields: fields_describe_chime_webhook_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChimeWebhookConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_chime_webhook_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeChimeWebhookConfigurations(ctx, input)
				}
				var results []*svc.DescribeChimeWebhookConfigurationsOutput
				p := svc.NewDescribeChimeWebhookConfigurationsPaginator(client, input)
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
		"describe-slack-channel-configurations": {
			Name:   "describe-slack-channel-configurations",
			Fields: fields_describe_slack_channel_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSlackChannelConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_slack_channel_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSlackChannelConfigurations(ctx, input)
				}
				var results []*svc.DescribeSlackChannelConfigurationsOutput
				p := svc.NewDescribeSlackChannelConfigurationsPaginator(client, input)
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
		"describe-slack-user-identities": {
			Name:   "describe-slack-user-identities",
			Fields: fields_describe_slack_user_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSlackUserIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_slack_user_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSlackUserIdentities(ctx, input)
				}
				var results []*svc.DescribeSlackUserIdentitiesOutput
				p := svc.NewDescribeSlackUserIdentitiesPaginator(client, input)
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
		"describe-slack-workspaces": {
			Name:   "describe-slack-workspaces",
			Fields: fields_describe_slack_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSlackWorkspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_slack_workspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSlackWorkspaces(ctx, input)
				}
				var results []*svc.DescribeSlackWorkspacesOutput
				p := svc.NewDescribeSlackWorkspacesPaginator(client, input)
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
		"disassociate-from-configuration": {
			Name:   "disassociate-from-configuration",
			Fields: fields_disassociate_from_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFromConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_from_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFromConfiguration(ctx, input)
			},
		},
		"get-account-preferences": {
			Name:   "get-account-preferences",
			Fields: fields_get_account_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountPreferences(ctx, input)
			},
		},
		"get-custom-action": {
			Name:   "get-custom-action",
			Fields: fields_get_custom_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomAction(ctx, input)
			},
		},
		"get-microsoft-teams-channel-configuration": {
			Name:   "get-microsoft-teams-channel-configuration",
			Fields: fields_get_microsoft_teams_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMicrosoftTeamsChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_microsoft_teams_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMicrosoftTeamsChannelConfiguration(ctx, input)
			},
		},
		"list-associations": {
			Name:   "list-associations",
			Fields: fields_list_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociations(ctx, input)
				}
				var results []*svc.ListAssociationsOutput
				p := svc.NewListAssociationsPaginator(client, input)
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
		"list-custom-actions": {
			Name:   "list-custom-actions",
			Fields: fields_list_custom_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomActions(ctx, input)
				}
				var results []*svc.ListCustomActionsOutput
				p := svc.NewListCustomActionsPaginator(client, input)
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
		"list-microsoft-teams-channel-configurations": {
			Name:   "list-microsoft-teams-channel-configurations",
			Fields: fields_list_microsoft_teams_channel_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMicrosoftTeamsChannelConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_microsoft_teams_channel_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMicrosoftTeamsChannelConfigurations(ctx, input)
				}
				var results []*svc.ListMicrosoftTeamsChannelConfigurationsOutput
				p := svc.NewListMicrosoftTeamsChannelConfigurationsPaginator(client, input)
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
		"list-microsoft-teams-configured-teams": {
			Name:   "list-microsoft-teams-configured-teams",
			Fields: fields_list_microsoft_teams_configured_teams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMicrosoftTeamsConfiguredTeamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_microsoft_teams_configured_teams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMicrosoftTeamsConfiguredTeams(ctx, input)
				}
				var results []*svc.ListMicrosoftTeamsConfiguredTeamsOutput
				p := svc.NewListMicrosoftTeamsConfiguredTeamsPaginator(client, input)
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
		"list-microsoft-teams-user-identities": {
			Name:   "list-microsoft-teams-user-identities",
			Fields: fields_list_microsoft_teams_user_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMicrosoftTeamsUserIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_microsoft_teams_user_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMicrosoftTeamsUserIdentities(ctx, input)
				}
				var results []*svc.ListMicrosoftTeamsUserIdentitiesOutput
				p := svc.NewListMicrosoftTeamsUserIdentitiesPaginator(client, input)
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
		"update-account-preferences": {
			Name:   "update-account-preferences",
			Fields: fields_update_account_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountPreferences(ctx, input)
			},
		},
		"update-chime-webhook-configuration": {
			Name:   "update-chime-webhook-configuration",
			Fields: fields_update_chime_webhook_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChimeWebhookConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_chime_webhook_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChimeWebhookConfiguration(ctx, input)
			},
		},
		"update-custom-action": {
			Name:   "update-custom-action",
			Fields: fields_update_custom_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomAction(ctx, input)
			},
		},
		"update-microsoft-teams-channel-configuration": {
			Name:   "update-microsoft-teams-channel-configuration",
			Fields: fields_update_microsoft_teams_channel_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMicrosoftTeamsChannelConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_microsoft_teams_channel_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMicrosoftTeamsChannelConfiguration(ctx, input)
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
	if err := leanruntime.Execute("chatbot", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
