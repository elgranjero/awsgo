package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/securitylake"
)

var fields_create_aws_log_source = []leanruntime.Field{
	{Name: "Sources", Flag: "sources", Type: "[]types.AwsLogSourceConfiguration", Required: true},
}

var fields_create_custom_log_source = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.CustomLogSourceConfiguration", Required: true},
	{Name: "EventClasses", Flag: "event-classes", Type: "[]string", Required: false},
	{Name: "SourceName", Flag: "source-name", Type: "*string", Required: true},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
}

var fields_create_data_lake = []leanruntime.Field{
	{Name: "Configurations", Flag: "configurations", Type: "[]types.DataLakeConfiguration", Required: true},
	{Name: "MetaStoreManagerRoleArn", Flag: "meta-store-manager-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_data_lake_exception_subscription = []leanruntime.Field{
	{Name: "ExceptionTimeToLive", Flag: "exception-time-to-live", Type: "*int64", Required: false},
	{Name: "NotificationEndpoint", Flag: "notification-endpoint", Type: "*string", Required: true},
	{Name: "SubscriptionProtocol", Flag: "subscription-protocol", Type: "*string", Required: true},
}

var fields_create_data_lake_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnableNewAccount", Flag: "auto-enable-new-account", Type: "[]types.DataLakeAutoEnableNewAccountConfiguration", Required: false},
}

var fields_create_subscriber = []leanruntime.Field{
	{Name: "AccessTypes", Flag: "access-types", Type: "[]types.AccessType", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.LogSourceResource", Required: true},
	{Name: "SubscriberDescription", Flag: "subscriber-description", Type: "*string", Required: false},
	{Name: "SubscriberIdentity", Flag: "subscriber-identity", Type: "*types.AwsIdentity", Required: true},
	{Name: "SubscriberName", Flag: "subscriber-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_subscriber_notification = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "types.NotificationConfiguration", Required: true},
	{Name: "SubscriberId", Flag: "subscriber-id", Type: "*string", Required: true},
}

var fields_delete_aws_log_source = []leanruntime.Field{
	{Name: "Sources", Flag: "sources", Type: "[]types.AwsLogSourceConfiguration", Required: true},
}

var fields_delete_custom_log_source = []leanruntime.Field{
	{Name: "SourceName", Flag: "source-name", Type: "*string", Required: true},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
}

var fields_delete_data_lake = []leanruntime.Field{
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
}

var fields_delete_data_lake_exception_subscription = []leanruntime.Field{}

var fields_delete_data_lake_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnableNewAccount", Flag: "auto-enable-new-account", Type: "[]types.DataLakeAutoEnableNewAccountConfiguration", Required: false},
}

var fields_delete_subscriber = []leanruntime.Field{
	{Name: "SubscriberId", Flag: "subscriber-id", Type: "*string", Required: true},
}

var fields_delete_subscriber_notification = []leanruntime.Field{
	{Name: "SubscriberId", Flag: "subscriber-id", Type: "*string", Required: true},
}

var fields_deregister_data_lake_delegated_administrator = []leanruntime.Field{}

var fields_get_data_lake_exception_subscription = []leanruntime.Field{}

var fields_get_data_lake_organization_configuration = []leanruntime.Field{}

var fields_get_data_lake_sources = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_subscriber = []leanruntime.Field{
	{Name: "SubscriberId", Flag: "subscriber-id", Type: "*string", Required: true},
}

var fields_list_data_lake_exceptions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
}

var fields_list_data_lakes = []leanruntime.Field{
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
}

var fields_list_log_sources = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.LogSourceResource", Required: false},
}

var fields_list_subscribers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_data_lake_delegated_administrator = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_data_lake = []leanruntime.Field{
	{Name: "Configurations", Flag: "configurations", Type: "[]types.DataLakeConfiguration", Required: true},
	{Name: "MetaStoreManagerRoleArn", Flag: "meta-store-manager-role-arn", Type: "*string", Required: false},
}

var fields_update_data_lake_exception_subscription = []leanruntime.Field{
	{Name: "ExceptionTimeToLive", Flag: "exception-time-to-live", Type: "*int64", Required: false},
	{Name: "NotificationEndpoint", Flag: "notification-endpoint", Type: "*string", Required: true},
	{Name: "SubscriptionProtocol", Flag: "subscription-protocol", Type: "*string", Required: true},
}

var fields_update_subscriber = []leanruntime.Field{
	{Name: "Sources", Flag: "sources", Type: "[]types.LogSourceResource", Required: false},
	{Name: "SubscriberDescription", Flag: "subscriber-description", Type: "*string", Required: false},
	{Name: "SubscriberId", Flag: "subscriber-id", Type: "*string", Required: true},
	{Name: "SubscriberIdentity", Flag: "subscriber-identity", Type: "*types.AwsIdentity", Required: false},
	{Name: "SubscriberName", Flag: "subscriber-name", Type: "*string", Required: false},
}

var fields_update_subscriber_notification = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "types.NotificationConfiguration", Required: true},
	{Name: "SubscriberId", Flag: "subscriber-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-aws-log-source": {
			Name:   "create-aws-log-source",
			Fields: fields_create_aws_log_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAwsLogSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_aws_log_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAwsLogSource(ctx, input)
			},
		},
		"create-custom-log-source": {
			Name:   "create-custom-log-source",
			Fields: fields_create_custom_log_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomLogSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_log_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomLogSource(ctx, input)
			},
		},
		"create-data-lake": {
			Name:   "create-data-lake",
			Fields: fields_create_data_lake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataLakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_lake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataLake(ctx, input)
			},
		},
		"create-data-lake-exception-subscription": {
			Name:   "create-data-lake-exception-subscription",
			Fields: fields_create_data_lake_exception_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataLakeExceptionSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_lake_exception_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataLakeExceptionSubscription(ctx, input)
			},
		},
		"create-data-lake-organization-configuration": {
			Name:   "create-data-lake-organization-configuration",
			Fields: fields_create_data_lake_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataLakeOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_lake_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataLakeOrganizationConfiguration(ctx, input)
			},
		},
		"create-subscriber": {
			Name:   "create-subscriber",
			Fields: fields_create_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriber(ctx, input)
			},
		},
		"create-subscriber-notification": {
			Name:   "create-subscriber-notification",
			Fields: fields_create_subscriber_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriberNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscriber_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriberNotification(ctx, input)
			},
		},
		"delete-aws-log-source": {
			Name:   "delete-aws-log-source",
			Fields: fields_delete_aws_log_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAwsLogSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_aws_log_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAwsLogSource(ctx, input)
			},
		},
		"delete-custom-log-source": {
			Name:   "delete-custom-log-source",
			Fields: fields_delete_custom_log_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomLogSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_log_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomLogSource(ctx, input)
			},
		},
		"delete-data-lake": {
			Name:   "delete-data-lake",
			Fields: fields_delete_data_lake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataLakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_lake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataLake(ctx, input)
			},
		},
		"delete-data-lake-exception-subscription": {
			Name:   "delete-data-lake-exception-subscription",
			Fields: fields_delete_data_lake_exception_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataLakeExceptionSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_lake_exception_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataLakeExceptionSubscription(ctx, input)
			},
		},
		"delete-data-lake-organization-configuration": {
			Name:   "delete-data-lake-organization-configuration",
			Fields: fields_delete_data_lake_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataLakeOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_lake_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataLakeOrganizationConfiguration(ctx, input)
			},
		},
		"delete-subscriber": {
			Name:   "delete-subscriber",
			Fields: fields_delete_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriber(ctx, input)
			},
		},
		"delete-subscriber-notification": {
			Name:   "delete-subscriber-notification",
			Fields: fields_delete_subscriber_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriberNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscriber_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriberNotification(ctx, input)
			},
		},
		"deregister-data-lake-delegated-administrator": {
			Name:   "deregister-data-lake-delegated-administrator",
			Fields: fields_deregister_data_lake_delegated_administrator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterDataLakeDelegatedAdministratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_data_lake_delegated_administrator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterDataLakeDelegatedAdministrator(ctx, input)
			},
		},
		"get-data-lake-exception-subscription": {
			Name:   "get-data-lake-exception-subscription",
			Fields: fields_get_data_lake_exception_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakeExceptionSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_lake_exception_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataLakeExceptionSubscription(ctx, input)
			},
		},
		"get-data-lake-organization-configuration": {
			Name:   "get-data-lake-organization-configuration",
			Fields: fields_get_data_lake_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakeOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_lake_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataLakeOrganizationConfiguration(ctx, input)
			},
		},
		"get-data-lake-sources": {
			Name:   "get-data-lake-sources",
			Fields: fields_get_data_lake_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakeSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_data_lake_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDataLakeSources(ctx, input)
				}
				var results []*svc.GetDataLakeSourcesOutput
				p := svc.NewGetDataLakeSourcesPaginator(client, input)
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
		"get-subscriber": {
			Name:   "get-subscriber",
			Fields: fields_get_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriber(ctx, input)
			},
		},
		"list-data-lake-exceptions": {
			Name:   "list-data-lake-exceptions",
			Fields: fields_list_data_lake_exceptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataLakeExceptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_lake_exceptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataLakeExceptions(ctx, input)
				}
				var results []*svc.ListDataLakeExceptionsOutput
				p := svc.NewListDataLakeExceptionsPaginator(client, input)
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
		"list-data-lakes": {
			Name:   "list-data-lakes",
			Fields: fields_list_data_lakes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataLakesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_data_lakes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDataLakes(ctx, input)
			},
		},
		"list-log-sources": {
			Name:   "list-log-sources",
			Fields: fields_list_log_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_log_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLogSources(ctx, input)
				}
				var results []*svc.ListLogSourcesOutput
				p := svc.NewListLogSourcesPaginator(client, input)
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
		"list-subscribers": {
			Name:   "list-subscribers",
			Fields: fields_list_subscribers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscribersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscribers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscribers(ctx, input)
				}
				var results []*svc.ListSubscribersOutput
				p := svc.NewListSubscribersPaginator(client, input)
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
		"register-data-lake-delegated-administrator": {
			Name:   "register-data-lake-delegated-administrator",
			Fields: fields_register_data_lake_delegated_administrator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDataLakeDelegatedAdministratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_data_lake_delegated_administrator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDataLakeDelegatedAdministrator(ctx, input)
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
		"update-data-lake": {
			Name:   "update-data-lake",
			Fields: fields_update_data_lake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataLakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_lake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataLake(ctx, input)
			},
		},
		"update-data-lake-exception-subscription": {
			Name:   "update-data-lake-exception-subscription",
			Fields: fields_update_data_lake_exception_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataLakeExceptionSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_lake_exception_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataLakeExceptionSubscription(ctx, input)
			},
		},
		"update-subscriber": {
			Name:   "update-subscriber",
			Fields: fields_update_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriber(ctx, input)
			},
		},
		"update-subscriber-notification": {
			Name:   "update-subscriber-notification",
			Fields: fields_update_subscriber_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriberNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscriber_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriberNotification(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("securitylake", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
