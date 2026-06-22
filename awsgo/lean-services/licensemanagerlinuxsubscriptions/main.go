package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions"
)

var fields_deregister_subscription_provider = []leanruntime.Field{
	{Name: "SubscriptionProviderArn", Flag: "subscription-provider-arn", Type: "*string", Required: true},
}

var fields_get_registered_subscription_provider = []leanruntime.Field{
	{Name: "SubscriptionProviderArn", Flag: "subscription-provider-arn", Type: "*string", Required: true},
}

var fields_get_service_settings = []leanruntime.Field{}

var fields_list_linux_subscription_instances = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_linux_subscriptions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_registered_subscription_providers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubscriptionProviderSources", Flag: "subscription-provider-sources", Type: "[]types.SubscriptionProviderSource", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_subscription_provider = []leanruntime.Field{
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
	{Name: "SubscriptionProviderSource", Flag: "subscription-provider-source", Type: "types.SubscriptionProviderSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_service_settings = []leanruntime.Field{
	{Name: "AllowUpdate", Flag: "allow-update", Type: "*bool", Required: false},
	{Name: "LinuxSubscriptionsDiscovery", Flag: "linux-subscriptions-discovery", Type: "types.LinuxSubscriptionsDiscovery", Required: true},
	{Name: "LinuxSubscriptionsDiscoverySettings", Flag: "linux-subscriptions-discovery-settings", Type: "*types.LinuxSubscriptionsDiscoverySettings", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"deregister-subscription-provider": {
			Name:   "deregister-subscription-provider",
			Fields: fields_deregister_subscription_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterSubscriptionProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_subscription_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterSubscriptionProvider(ctx, input)
			},
		},
		"get-registered-subscription-provider": {
			Name:   "get-registered-subscription-provider",
			Fields: fields_get_registered_subscription_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegisteredSubscriptionProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_registered_subscription_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegisteredSubscriptionProvider(ctx, input)
			},
		},
		"get-service-settings": {
			Name:   "get-service-settings",
			Fields: fields_get_service_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceSettings(ctx, input)
			},
		},
		"list-linux-subscription-instances": {
			Name:   "list-linux-subscription-instances",
			Fields: fields_list_linux_subscription_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLinuxSubscriptionInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_linux_subscription_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLinuxSubscriptionInstances(ctx, input)
				}
				var results []*svc.ListLinuxSubscriptionInstancesOutput
				p := svc.NewListLinuxSubscriptionInstancesPaginator(client, input)
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
		"list-linux-subscriptions": {
			Name:   "list-linux-subscriptions",
			Fields: fields_list_linux_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLinuxSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_linux_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLinuxSubscriptions(ctx, input)
				}
				var results []*svc.ListLinuxSubscriptionsOutput
				p := svc.NewListLinuxSubscriptionsPaginator(client, input)
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
		"list-registered-subscription-providers": {
			Name:   "list-registered-subscription-providers",
			Fields: fields_list_registered_subscription_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegisteredSubscriptionProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_registered_subscription_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegisteredSubscriptionProviders(ctx, input)
				}
				var results []*svc.ListRegisteredSubscriptionProvidersOutput
				p := svc.NewListRegisteredSubscriptionProvidersPaginator(client, input)
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
		"register-subscription-provider": {
			Name:   "register-subscription-provider",
			Fields: fields_register_subscription_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterSubscriptionProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_subscription_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterSubscriptionProvider(ctx, input)
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
		"update-service-settings": {
			Name:   "update-service-settings",
			Fields: fields_update_service_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("licensemanagerlinuxsubscriptions", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
