package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions"
)

var fields_associate_user = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_create_license_server_endpoint = []leanruntime.Field{
	{Name: "IdentityProviderArn", Flag: "identity-provider-arn", Type: "*string", Required: true},
	{Name: "LicenseServerSettings", Flag: "license-server-settings", Type: "*types.LicenseServerSettings", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_license_server_endpoint = []leanruntime.Field{
	{Name: "LicenseServerEndpointArn", Flag: "license-server-endpoint-arn", Type: "*string", Required: true},
	{Name: "ServerType", Flag: "server-type", Type: "types.ServerType", Required: true},
}

var fields_deregister_identity_provider = []leanruntime.Field{
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: false},
	{Name: "IdentityProviderArn", Flag: "identity-provider-arn", Type: "*string", Required: false},
	{Name: "Product", Flag: "product", Type: "*string", Required: false},
}

var fields_disassociate_user = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "InstanceUserArn", Flag: "instance-user-arn", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_list_identity_providers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_instances = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_server_endpoints = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_product_subscriptions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Product", Flag: "product", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_user_associations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_register_identity_provider = []leanruntime.Field{
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: true},
	{Name: "Product", Flag: "product", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.Settings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_product_subscription = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: true},
	{Name: "Product", Flag: "product", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_stop_product_subscription = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: false},
	{Name: "Product", Flag: "product", Type: "*string", Required: false},
	{Name: "ProductUserArn", Flag: "product-user-arn", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_identity_provider_settings = []leanruntime.Field{
	{Name: "IdentityProvider", Flag: "identity-provider", Type: "types.IdentityProvider", Required: false},
	{Name: "IdentityProviderArn", Flag: "identity-provider-arn", Type: "*string", Required: false},
	{Name: "Product", Flag: "product", Type: "*string", Required: false},
	{Name: "UpdateSettings", Flag: "update-settings", Type: "*types.UpdateSettings", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-user": {
			Name:   "associate-user",
			Fields: fields_associate_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateUser(ctx, input)
			},
		},
		"create-license-server-endpoint": {
			Name:   "create-license-server-endpoint",
			Fields: fields_create_license_server_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseServerEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_server_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseServerEndpoint(ctx, input)
			},
		},
		"delete-license-server-endpoint": {
			Name:   "delete-license-server-endpoint",
			Fields: fields_delete_license_server_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseServerEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license_server_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicenseServerEndpoint(ctx, input)
			},
		},
		"deregister-identity-provider": {
			Name:   "deregister-identity-provider",
			Fields: fields_deregister_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterIdentityProvider(ctx, input)
			},
		},
		"disassociate-user": {
			Name:   "disassociate-user",
			Fields: fields_disassociate_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateUser(ctx, input)
			},
		},
		"list-identity-providers": {
			Name:   "list-identity-providers",
			Fields: fields_list_identity_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identity_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentityProviders(ctx, input)
				}
				var results []*svc.ListIdentityProvidersOutput
				p := svc.NewListIdentityProvidersPaginator(client, input)
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
		"list-instances": {
			Name:   "list-instances",
			Fields: fields_list_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstances(ctx, input)
				}
				var results []*svc.ListInstancesOutput
				p := svc.NewListInstancesPaginator(client, input)
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
		"list-license-server-endpoints": {
			Name:   "list-license-server-endpoints",
			Fields: fields_list_license_server_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseServerEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_license_server_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLicenseServerEndpoints(ctx, input)
				}
				var results []*svc.ListLicenseServerEndpointsOutput
				p := svc.NewListLicenseServerEndpointsPaginator(client, input)
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
		"list-product-subscriptions": {
			Name:   "list-product-subscriptions",
			Fields: fields_list_product_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProductSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_product_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProductSubscriptions(ctx, input)
				}
				var results []*svc.ListProductSubscriptionsOutput
				p := svc.NewListProductSubscriptionsPaginator(client, input)
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
		"list-user-associations": {
			Name:   "list-user-associations",
			Fields: fields_list_user_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserAssociations(ctx, input)
				}
				var results []*svc.ListUserAssociationsOutput
				p := svc.NewListUserAssociationsPaginator(client, input)
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
		"register-identity-provider": {
			Name:   "register-identity-provider",
			Fields: fields_register_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterIdentityProvider(ctx, input)
			},
		},
		"start-product-subscription": {
			Name:   "start-product-subscription",
			Fields: fields_start_product_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartProductSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_product_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartProductSubscription(ctx, input)
			},
		},
		"stop-product-subscription": {
			Name:   "stop-product-subscription",
			Fields: fields_stop_product_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopProductSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_product_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopProductSubscription(ctx, input)
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
		"update-identity-provider-settings": {
			Name:   "update-identity-provider-settings",
			Fields: fields_update_identity_provider_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdentityProviderSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_identity_provider_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdentityProviderSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("licensemanagerusersubscriptions", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
