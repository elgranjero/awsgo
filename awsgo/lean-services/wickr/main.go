package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/wickr"
)

var fields_batch_create_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Users", Flag: "users", Type: "[]types.BatchCreateUserRequestItem", Required: true},
}

var fields_batch_delete_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: true},
}

var fields_batch_lookup_user_uname = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Unames", Flag: "unames", Type: "[]string", Required: true},
}

var fields_batch_reinvite_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: true},
}

var fields_batch_reset_devices_for_user = []leanruntime.Field{
	{Name: "AppIds", Flag: "app-ids", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_batch_toggle_user_suspend_status = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Suspend", Flag: "suspend", Type: "*bool", Required: true},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: true},
}

var fields_create_bot = []leanruntime.Field{
	{Name: "Challenge", Flag: "challenge", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_create_data_retention_bot = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_create_data_retention_bot_challenge = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_create_network = []leanruntime.Field{
	{Name: "AccessLevel", Flag: "access-level", Type: "types.AccessLevel", Required: true},
	{Name: "EnablePremiumFreeTrial", Flag: "enable-premium-free-trial", Type: "*bool", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "NetworkName", Flag: "network-name", Type: "*string", Required: true},
}

var fields_create_security_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "SecurityGroupSettings", Flag: "security-group-settings", Type: "*types.SecurityGroupSettingsRequest", Required: true},
}

var fields_delete_bot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_delete_data_retention_bot = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_delete_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_delete_security_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_bot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_bots_count = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_data_retention_bot = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_guest_user_history_count = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_network = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_network_settings = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_oidc_info = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: false},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "CodeVerifier", Flag: "code-verifier", Type: "*string", Required: false},
	{Name: "GrantType", Flag: "grant-type", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "RedirectUri", Flag: "redirect-uri", Type: "*string", Required: false},
	{Name: "Url", Flag: "url", Type: "*string", Required: false},
}

var fields_get_opentdf_config = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_security_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_user = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_users_count = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_list_blocked_guest_users = []leanruntime.Field{
	{Name: "Admin", Flag: "admin", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_list_bots = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BotStatus", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_list_devices_for_user = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_list_guest_users = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_list_networks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
}

var fields_list_security_group_users = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
}

var fields_list_security_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
}

var fields_list_users = []leanruntime.Field{
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortDirection", Flag: "sort-direction", Type: "types.SortDirection", Required: false},
	{Name: "SortFields", Flag: "sort-fields", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.UserStatus", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_register_oidc_config = []leanruntime.Field{
	{Name: "CompanyId", Flag: "company-id", Type: "*string", Required: true},
	{Name: "CustomUsername", Flag: "custom-username", Type: "*string", Required: false},
	{Name: "ExtraAuthParams", Flag: "extra-auth-params", Type: "*string", Required: false},
	{Name: "Issuer", Flag: "issuer", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Scopes", Flag: "scopes", Type: "*string", Required: true},
	{Name: "Secret", Flag: "secret", Type: "*string", Required: false},
	{Name: "SsoTokenBufferMinutes", Flag: "sso-token-buffer-minutes", Type: "*int32", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_register_oidc_config_test = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: false},
	{Name: "ExtraAuthParams", Flag: "extra-auth-params", Type: "*string", Required: false},
	{Name: "Issuer", Flag: "issuer", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Scopes", Flag: "scopes", Type: "*string", Required: true},
}

var fields_register_opentdf_config = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: true},
}

var fields_update_bot = []leanruntime.Field{
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "Challenge", Flag: "challenge", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Suspend", Flag: "suspend", Type: "*bool", Required: false},
}

var fields_update_data_retention = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.DataRetentionActionType", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_update_guest_user = []leanruntime.Field{
	{Name: "Block", Flag: "block", Type: "*bool", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "UsernameHash", Flag: "username-hash", Type: "*string", Required: true},
}

var fields_update_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NetworkName", Flag: "network-name", Type: "*string", Required: true},
}

var fields_update_network_settings = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.NetworkSettings", Required: true},
}

var fields_update_security_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "SecurityGroupSettings", Flag: "security-group-settings", Type: "*types.SecurityGroupSettings", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "UserDetails", Flag: "user-details", Type: "*types.UpdateUserDetails", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-user": {
			Name:   "batch-create-user",
			Fields: fields_batch_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateUser(ctx, input)
			},
		},
		"batch-delete-user": {
			Name:   "batch-delete-user",
			Fields: fields_batch_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteUser(ctx, input)
			},
		},
		"batch-lookup-user-uname": {
			Name:   "batch-lookup-user-uname",
			Fields: fields_batch_lookup_user_uname,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchLookupUserUnameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_lookup_user_uname, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchLookupUserUname(ctx, input)
			},
		},
		"batch-reinvite-user": {
			Name:   "batch-reinvite-user",
			Fields: fields_batch_reinvite_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchReinviteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_reinvite_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchReinviteUser(ctx, input)
			},
		},
		"batch-reset-devices-for-user": {
			Name:   "batch-reset-devices-for-user",
			Fields: fields_batch_reset_devices_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchResetDevicesForUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_reset_devices_for_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchResetDevicesForUser(ctx, input)
			},
		},
		"batch-toggle-user-suspend-status": {
			Name:   "batch-toggle-user-suspend-status",
			Fields: fields_batch_toggle_user_suspend_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchToggleUserSuspendStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_toggle_user_suspend_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchToggleUserSuspendStatus(ctx, input)
			},
		},
		"create-bot": {
			Name:   "create-bot",
			Fields: fields_create_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBot(ctx, input)
			},
		},
		"create-data-retention-bot": {
			Name:   "create-data-retention-bot",
			Fields: fields_create_data_retention_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataRetentionBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_retention_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataRetentionBot(ctx, input)
			},
		},
		"create-data-retention-bot-challenge": {
			Name:   "create-data-retention-bot-challenge",
			Fields: fields_create_data_retention_bot_challenge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataRetentionBotChallengeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_retention_bot_challenge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataRetentionBotChallenge(ctx, input)
			},
		},
		"create-network": {
			Name:   "create-network",
			Fields: fields_create_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetwork(ctx, input)
			},
		},
		"create-security-group": {
			Name:   "create-security-group",
			Fields: fields_create_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityGroup(ctx, input)
			},
		},
		"delete-bot": {
			Name:   "delete-bot",
			Fields: fields_delete_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBot(ctx, input)
			},
		},
		"delete-data-retention-bot": {
			Name:   "delete-data-retention-bot",
			Fields: fields_delete_data_retention_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataRetentionBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_retention_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataRetentionBot(ctx, input)
			},
		},
		"delete-network": {
			Name:   "delete-network",
			Fields: fields_delete_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetwork(ctx, input)
			},
		},
		"delete-security-group": {
			Name:   "delete-security-group",
			Fields: fields_delete_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityGroup(ctx, input)
			},
		},
		"get-bot": {
			Name:   "get-bot",
			Fields: fields_get_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBot(ctx, input)
			},
		},
		"get-bots-count": {
			Name:   "get-bots-count",
			Fields: fields_get_bots_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotsCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bots_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBotsCount(ctx, input)
			},
		},
		"get-data-retention-bot": {
			Name:   "get-data-retention-bot",
			Fields: fields_get_data_retention_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataRetentionBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_retention_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataRetentionBot(ctx, input)
			},
		},
		"get-guest-user-history-count": {
			Name:   "get-guest-user-history-count",
			Fields: fields_get_guest_user_history_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGuestUserHistoryCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_guest_user_history_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGuestUserHistoryCount(ctx, input)
			},
		},
		"get-network": {
			Name:   "get-network",
			Fields: fields_get_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetwork(ctx, input)
			},
		},
		"get-network-settings": {
			Name:   "get-network-settings",
			Fields: fields_get_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetworkSettings(ctx, input)
			},
		},
		"get-oidc-info": {
			Name:   "get-oidc-info",
			Fields: fields_get_oidc_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOidcInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_oidc_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOidcInfo(ctx, input)
			},
		},
		"get-opentdf-config": {
			Name:   "get-opentdf-config",
			Fields: fields_get_opentdf_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpentdfConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_opentdf_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpentdfConfig(ctx, input)
			},
		},
		"get-security-group": {
			Name:   "get-security-group",
			Fields: fields_get_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSecurityGroup(ctx, input)
			},
		},
		"get-user": {
			Name:   "get-user",
			Fields: fields_get_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUser(ctx, input)
			},
		},
		"get-users-count": {
			Name:   "get-users-count",
			Fields: fields_get_users_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsersCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_users_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsersCount(ctx, input)
			},
		},
		"list-blocked-guest-users": {
			Name:   "list-blocked-guest-users",
			Fields: fields_list_blocked_guest_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBlockedGuestUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_blocked_guest_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBlockedGuestUsers(ctx, input)
				}
				var results []*svc.ListBlockedGuestUsersOutput
				p := svc.NewListBlockedGuestUsersPaginator(client, input)
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
		"list-bots": {
			Name:   "list-bots",
			Fields: fields_list_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBots(ctx, input)
				}
				var results []*svc.ListBotsOutput
				p := svc.NewListBotsPaginator(client, input)
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
		"list-devices-for-user": {
			Name:   "list-devices-for-user",
			Fields: fields_list_devices_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesForUserInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_devices_for_user, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevicesForUser(ctx, input)
				}
				var results []*svc.ListDevicesForUserOutput
				p := svc.NewListDevicesForUserPaginator(client, input)
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
		"list-guest-users": {
			Name:   "list-guest-users",
			Fields: fields_list_guest_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGuestUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_guest_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGuestUsers(ctx, input)
				}
				var results []*svc.ListGuestUsersOutput
				p := svc.NewListGuestUsersPaginator(client, input)
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
		"list-networks": {
			Name:   "list-networks",
			Fields: fields_list_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNetworks(ctx, input)
				}
				var results []*svc.ListNetworksOutput
				p := svc.NewListNetworksPaginator(client, input)
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
		"list-security-group-users": {
			Name:   "list-security-group-users",
			Fields: fields_list_security_group_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityGroupUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_group_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityGroupUsers(ctx, input)
				}
				var results []*svc.ListSecurityGroupUsersOutput
				p := svc.NewListSecurityGroupUsersPaginator(client, input)
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
		"list-security-groups": {
			Name:   "list-security-groups",
			Fields: fields_list_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityGroups(ctx, input)
				}
				var results []*svc.ListSecurityGroupsOutput
				p := svc.NewListSecurityGroupsPaginator(client, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"register-oidc-config": {
			Name:   "register-oidc-config",
			Fields: fields_register_oidc_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterOidcConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_oidc_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterOidcConfig(ctx, input)
			},
		},
		"register-oidc-config-test": {
			Name:   "register-oidc-config-test",
			Fields: fields_register_oidc_config_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterOidcConfigTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_oidc_config_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterOidcConfigTest(ctx, input)
			},
		},
		"register-opentdf-config": {
			Name:   "register-opentdf-config",
			Fields: fields_register_opentdf_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterOpentdfConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_opentdf_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterOpentdfConfig(ctx, input)
			},
		},
		"update-bot": {
			Name:   "update-bot",
			Fields: fields_update_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBot(ctx, input)
			},
		},
		"update-data-retention": {
			Name:   "update-data-retention",
			Fields: fields_update_data_retention,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataRetentionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_retention, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataRetention(ctx, input)
			},
		},
		"update-guest-user": {
			Name:   "update-guest-user",
			Fields: fields_update_guest_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGuestUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_guest_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGuestUser(ctx, input)
			},
		},
		"update-network": {
			Name:   "update-network",
			Fields: fields_update_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetwork(ctx, input)
			},
		},
		"update-network-settings": {
			Name:   "update-network-settings",
			Fields: fields_update_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetworkSettings(ctx, input)
			},
		},
		"update-security-group": {
			Name:   "update-security-group",
			Fields: fields_update_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityGroup(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("wickr", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
