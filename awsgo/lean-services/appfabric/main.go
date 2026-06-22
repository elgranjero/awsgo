package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appfabric"
)

var fields_batch_get_user_access_tasks = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "TaskIdList", Flag: "task-id-list", Type: "[]string", Required: true},
}

var fields_connect_app_authorization = []leanruntime.Field{
	{Name: "AppAuthorizationIdentifier", Flag: "app-authorization-identifier", Type: "*string", Required: true},
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "AuthRequest", Flag: "auth-request", Type: "*types.AuthRequest", Required: false},
}

var fields_create_app_authorization = []leanruntime.Field{
	{Name: "App", Flag: "app", Type: "*string", Required: true},
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Credential", Flag: "credential", Type: "types.Credential", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Tenant", Flag: "tenant", Type: "*types.Tenant", Required: true},
}

var fields_create_app_bundle = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerManagedKeyIdentifier", Flag: "customer-managed-key-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_ingestion = []leanruntime.Field{
	{Name: "App", Flag: "app", Type: "*string", Required: true},
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IngestionType", Flag: "ingestion-type", Type: "types.IngestionType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TenantId", Flag: "tenant-id", Type: "*string", Required: true},
}

var fields_create_ingestion_destination = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "types.DestinationConfiguration", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
	{Name: "ProcessingConfiguration", Flag: "processing-configuration", Type: "types.ProcessingConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_app_authorization = []leanruntime.Field{
	{Name: "AppAuthorizationIdentifier", Flag: "app-authorization-identifier", Type: "*string", Required: true},
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
}

var fields_delete_app_bundle = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
}

var fields_delete_ingestion = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

var fields_delete_ingestion_destination = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionDestinationIdentifier", Flag: "ingestion-destination-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

var fields_get_app_authorization = []leanruntime.Field{
	{Name: "AppAuthorizationIdentifier", Flag: "app-authorization-identifier", Type: "*string", Required: true},
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
}

var fields_get_app_bundle = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
}

var fields_get_ingestion = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

var fields_get_ingestion_destination = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionDestinationIdentifier", Flag: "ingestion-destination-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

var fields_list_app_authorizations = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_bundles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ingestion_destinations = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ingestions = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_ingestion = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

var fields_start_user_access_tasks = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
}

var fields_stop_ingestion = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_app_authorization = []leanruntime.Field{
	{Name: "AppAuthorizationIdentifier", Flag: "app-authorization-identifier", Type: "*string", Required: true},
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "Credential", Flag: "credential", Type: "types.Credential", Required: false},
	{Name: "Tenant", Flag: "tenant", Type: "*types.Tenant", Required: false},
}

var fields_update_ingestion_destination = []leanruntime.Field{
	{Name: "AppBundleIdentifier", Flag: "app-bundle-identifier", Type: "*string", Required: true},
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "types.DestinationConfiguration", Required: true},
	{Name: "IngestionDestinationIdentifier", Flag: "ingestion-destination-identifier", Type: "*string", Required: true},
	{Name: "IngestionIdentifier", Flag: "ingestion-identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-user-access-tasks": {
			Name:   "batch-get-user-access-tasks",
			Fields: fields_batch_get_user_access_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetUserAccessTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_user_access_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetUserAccessTasks(ctx, input)
			},
		},
		"connect-app-authorization": {
			Name:   "connect-app-authorization",
			Fields: fields_connect_app_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConnectAppAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_connect_app_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConnectAppAuthorization(ctx, input)
			},
		},
		"create-app-authorization": {
			Name:   "create-app-authorization",
			Fields: fields_create_app_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppAuthorization(ctx, input)
			},
		},
		"create-app-bundle": {
			Name:   "create-app-bundle",
			Fields: fields_create_app_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppBundle(ctx, input)
			},
		},
		"create-ingestion": {
			Name:   "create-ingestion",
			Fields: fields_create_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIngestion(ctx, input)
			},
		},
		"create-ingestion-destination": {
			Name:   "create-ingestion-destination",
			Fields: fields_create_ingestion_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIngestionDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ingestion_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIngestionDestination(ctx, input)
			},
		},
		"delete-app-authorization": {
			Name:   "delete-app-authorization",
			Fields: fields_delete_app_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppAuthorization(ctx, input)
			},
		},
		"delete-app-bundle": {
			Name:   "delete-app-bundle",
			Fields: fields_delete_app_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppBundle(ctx, input)
			},
		},
		"delete-ingestion": {
			Name:   "delete-ingestion",
			Fields: fields_delete_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIngestion(ctx, input)
			},
		},
		"delete-ingestion-destination": {
			Name:   "delete-ingestion-destination",
			Fields: fields_delete_ingestion_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIngestionDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ingestion_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIngestionDestination(ctx, input)
			},
		},
		"get-app-authorization": {
			Name:   "get-app-authorization",
			Fields: fields_get_app_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_app_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAppAuthorization(ctx, input)
			},
		},
		"get-app-bundle": {
			Name:   "get-app-bundle",
			Fields: fields_get_app_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_app_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAppBundle(ctx, input)
			},
		},
		"get-ingestion": {
			Name:   "get-ingestion",
			Fields: fields_get_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIngestion(ctx, input)
			},
		},
		"get-ingestion-destination": {
			Name:   "get-ingestion-destination",
			Fields: fields_get_ingestion_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIngestionDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ingestion_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIngestionDestination(ctx, input)
			},
		},
		"list-app-authorizations": {
			Name:   "list-app-authorizations",
			Fields: fields_list_app_authorizations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppAuthorizationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_authorizations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppAuthorizations(ctx, input)
				}
				var results []*svc.ListAppAuthorizationsOutput
				p := svc.NewListAppAuthorizationsPaginator(client, input)
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
		"list-app-bundles": {
			Name:   "list-app-bundles",
			Fields: fields_list_app_bundles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppBundlesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_bundles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppBundles(ctx, input)
				}
				var results []*svc.ListAppBundlesOutput
				p := svc.NewListAppBundlesPaginator(client, input)
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
		"list-ingestion-destinations": {
			Name:   "list-ingestion-destinations",
			Fields: fields_list_ingestion_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIngestionDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ingestion_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIngestionDestinations(ctx, input)
				}
				var results []*svc.ListIngestionDestinationsOutput
				p := svc.NewListIngestionDestinationsPaginator(client, input)
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
		"list-ingestions": {
			Name:   "list-ingestions",
			Fields: fields_list_ingestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIngestionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ingestions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIngestions(ctx, input)
				}
				var results []*svc.ListIngestionsOutput
				p := svc.NewListIngestionsPaginator(client, input)
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
		"start-ingestion": {
			Name:   "start-ingestion",
			Fields: fields_start_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartIngestion(ctx, input)
			},
		},
		"start-user-access-tasks": {
			Name:   "start-user-access-tasks",
			Fields: fields_start_user_access_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartUserAccessTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_user_access_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartUserAccessTasks(ctx, input)
			},
		},
		"stop-ingestion": {
			Name:   "stop-ingestion",
			Fields: fields_stop_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopIngestion(ctx, input)
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
		"update-app-authorization": {
			Name:   "update-app-authorization",
			Fields: fields_update_app_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppAuthorization(ctx, input)
			},
		},
		"update-ingestion-destination": {
			Name:   "update-ingestion-destination",
			Fields: fields_update_ingestion_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIngestionDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ingestion_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIngestionDestination(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appfabric", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
