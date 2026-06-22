package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chimesdkidentity"
)

var fields_create_app_instance = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_app_instance_admin = []leanruntime.Field{
	{Name: "AppInstanceAdminArn", Flag: "app-instance-admin-arn", Type: "*string", Required: true},
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_create_app_instance_bot = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.Configuration", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "AppInstanceUserId", Flag: "app-instance-user-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ExpirationSettings", Flag: "expiration-settings", Type: "*types.ExpirationSettings", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_app_instance = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_delete_app_instance_admin = []leanruntime.Field{
	{Name: "AppInstanceAdminArn", Flag: "app-instance-admin-arn", Type: "*string", Required: true},
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_delete_app_instance_bot = []leanruntime.Field{
	{Name: "AppInstanceBotArn", Flag: "app-instance-bot-arn", Type: "*string", Required: true},
}

var fields_delete_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
}

var fields_deregister_app_instance_user_endpoint = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_describe_app_instance = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_describe_app_instance_admin = []leanruntime.Field{
	{Name: "AppInstanceAdminArn", Flag: "app-instance-admin-arn", Type: "*string", Required: true},
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_describe_app_instance_bot = []leanruntime.Field{
	{Name: "AppInstanceBotArn", Flag: "app-instance-bot-arn", Type: "*string", Required: true},
}

var fields_describe_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
}

var fields_describe_app_instance_user_endpoint = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_get_app_instance_retention_settings = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_list_app_instance_admins = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_instance_bots = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_instance_user_endpoints = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_instance_users = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_app_instance_retention_settings = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "AppInstanceRetentionSettings", Flag: "app-instance-retention-settings", Type: "*types.AppInstanceRetentionSettings", Required: true},
}

var fields_put_app_instance_user_expiration_settings = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "ExpirationSettings", Flag: "expiration-settings", Type: "*types.ExpirationSettings", Required: false},
}

var fields_register_app_instance_user_endpoint = []leanruntime.Field{
	{Name: "AllowMessages", Flag: "allow-messages", Type: "types.AllowMessages", Required: false},
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "EndpointAttributes", Flag: "endpoint-attributes", Type: "*types.EndpointAttributes", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AppInstanceUserEndpointType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_app_instance = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_app_instance_bot = []leanruntime.Field{
	{Name: "AppInstanceBotArn", Flag: "app-instance-bot-arn", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.Configuration", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_app_instance_user_endpoint = []leanruntime.Field{
	{Name: "AllowMessages", Flag: "allow-messages", Type: "types.AllowMessages", Required: false},
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-app-instance": {
			Name:   "create-app-instance",
			Fields: fields_create_app_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppInstance(ctx, input)
			},
		},
		"create-app-instance-admin": {
			Name:   "create-app-instance-admin",
			Fields: fields_create_app_instance_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInstanceAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_instance_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppInstanceAdmin(ctx, input)
			},
		},
		"create-app-instance-bot": {
			Name:   "create-app-instance-bot",
			Fields: fields_create_app_instance_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInstanceBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_instance_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppInstanceBot(ctx, input)
			},
		},
		"create-app-instance-user": {
			Name:   "create-app-instance-user",
			Fields: fields_create_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInstanceUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_instance_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppInstanceUser(ctx, input)
			},
		},
		"delete-app-instance": {
			Name:   "delete-app-instance",
			Fields: fields_delete_app_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppInstance(ctx, input)
			},
		},
		"delete-app-instance-admin": {
			Name:   "delete-app-instance-admin",
			Fields: fields_delete_app_instance_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInstanceAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_instance_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppInstanceAdmin(ctx, input)
			},
		},
		"delete-app-instance-bot": {
			Name:   "delete-app-instance-bot",
			Fields: fields_delete_app_instance_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInstanceBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_instance_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppInstanceBot(ctx, input)
			},
		},
		"delete-app-instance-user": {
			Name:   "delete-app-instance-user",
			Fields: fields_delete_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInstanceUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_instance_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppInstanceUser(ctx, input)
			},
		},
		"deregister-app-instance-user-endpoint": {
			Name:   "deregister-app-instance-user-endpoint",
			Fields: fields_deregister_app_instance_user_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterAppInstanceUserEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_app_instance_user_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterAppInstanceUserEndpoint(ctx, input)
			},
		},
		"describe-app-instance": {
			Name:   "describe-app-instance",
			Fields: fields_describe_app_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppInstance(ctx, input)
			},
		},
		"describe-app-instance-admin": {
			Name:   "describe-app-instance-admin",
			Fields: fields_describe_app_instance_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInstanceAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_instance_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppInstanceAdmin(ctx, input)
			},
		},
		"describe-app-instance-bot": {
			Name:   "describe-app-instance-bot",
			Fields: fields_describe_app_instance_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInstanceBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_instance_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppInstanceBot(ctx, input)
			},
		},
		"describe-app-instance-user": {
			Name:   "describe-app-instance-user",
			Fields: fields_describe_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInstanceUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_instance_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppInstanceUser(ctx, input)
			},
		},
		"describe-app-instance-user-endpoint": {
			Name:   "describe-app-instance-user-endpoint",
			Fields: fields_describe_app_instance_user_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInstanceUserEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_instance_user_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppInstanceUserEndpoint(ctx, input)
			},
		},
		"get-app-instance-retention-settings": {
			Name:   "get-app-instance-retention-settings",
			Fields: fields_get_app_instance_retention_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppInstanceRetentionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_app_instance_retention_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAppInstanceRetentionSettings(ctx, input)
			},
		},
		"list-app-instance-admins": {
			Name:   "list-app-instance-admins",
			Fields: fields_list_app_instance_admins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppInstanceAdminsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_instance_admins, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppInstanceAdmins(ctx, input)
				}
				var results []*svc.ListAppInstanceAdminsOutput
				p := svc.NewListAppInstanceAdminsPaginator(client, input)
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
		"list-app-instance-bots": {
			Name:   "list-app-instance-bots",
			Fields: fields_list_app_instance_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppInstanceBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_instance_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppInstanceBots(ctx, input)
				}
				var results []*svc.ListAppInstanceBotsOutput
				p := svc.NewListAppInstanceBotsPaginator(client, input)
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
		"list-app-instance-user-endpoints": {
			Name:   "list-app-instance-user-endpoints",
			Fields: fields_list_app_instance_user_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppInstanceUserEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_instance_user_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppInstanceUserEndpoints(ctx, input)
				}
				var results []*svc.ListAppInstanceUserEndpointsOutput
				p := svc.NewListAppInstanceUserEndpointsPaginator(client, input)
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
		"list-app-instance-users": {
			Name:   "list-app-instance-users",
			Fields: fields_list_app_instance_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppInstanceUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_instance_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppInstanceUsers(ctx, input)
				}
				var results []*svc.ListAppInstanceUsersOutput
				p := svc.NewListAppInstanceUsersPaginator(client, input)
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
		"list-app-instances": {
			Name:   "list-app-instances",
			Fields: fields_list_app_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppInstances(ctx, input)
				}
				var results []*svc.ListAppInstancesOutput
				p := svc.NewListAppInstancesPaginator(client, input)
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
		"put-app-instance-retention-settings": {
			Name:   "put-app-instance-retention-settings",
			Fields: fields_put_app_instance_retention_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAppInstanceRetentionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_app_instance_retention_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAppInstanceRetentionSettings(ctx, input)
			},
		},
		"put-app-instance-user-expiration-settings": {
			Name:   "put-app-instance-user-expiration-settings",
			Fields: fields_put_app_instance_user_expiration_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAppInstanceUserExpirationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_app_instance_user_expiration_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAppInstanceUserExpirationSettings(ctx, input)
			},
		},
		"register-app-instance-user-endpoint": {
			Name:   "register-app-instance-user-endpoint",
			Fields: fields_register_app_instance_user_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterAppInstanceUserEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_app_instance_user_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterAppInstanceUserEndpoint(ctx, input)
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
		"update-app-instance": {
			Name:   "update-app-instance",
			Fields: fields_update_app_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppInstance(ctx, input)
			},
		},
		"update-app-instance-bot": {
			Name:   "update-app-instance-bot",
			Fields: fields_update_app_instance_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppInstanceBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_instance_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppInstanceBot(ctx, input)
			},
		},
		"update-app-instance-user": {
			Name:   "update-app-instance-user",
			Fields: fields_update_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppInstanceUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_instance_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppInstanceUser(ctx, input)
			},
		},
		"update-app-instance-user-endpoint": {
			Name:   "update-app-instance-user-endpoint",
			Fields: fields_update_app_instance_user_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppInstanceUserEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_instance_user_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppInstanceUserEndpoint(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("chimesdkidentity", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
