package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
)

var fields_associate_default_view = []leanruntime.Field{
	{Name: "ViewArn", Flag: "view-arn", Type: "*string", Required: true},
}

var fields_batch_get_view = []leanruntime.Field{
	{Name: "ViewArns", Flag: "view-arns", Type: "[]string", Required: false},
}

var fields_create_index = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_resource_explorer_setup = []leanruntime.Field{
	{Name: "AggregatorRegions", Flag: "aggregator-regions", Type: "[]string", Required: false},
	{Name: "RegionList", Flag: "region-list", Type: "[]string", Required: true},
	{Name: "ViewName", Flag: "view-name", Type: "*string", Required: true},
}

var fields_create_view = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.SearchFilter", Required: false},
	{Name: "IncludedProperties", Flag: "included-properties", Type: "[]types.IncludedProperty", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "ViewName", Flag: "view-name", Type: "*string", Required: true},
}

var fields_delete_index = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_resource_explorer_setup = []leanruntime.Field{
	{Name: "DeleteInAllRegions", Flag: "delete-in-all-regions", Type: "*bool", Required: false},
	{Name: "RegionList", Flag: "region-list", Type: "[]string", Required: false},
}

var fields_delete_view = []leanruntime.Field{
	{Name: "ViewArn", Flag: "view-arn", Type: "*string", Required: true},
}

var fields_disassociate_default_view = []leanruntime.Field{}

var fields_get_account_level_service_configuration = []leanruntime.Field{}

var fields_get_default_view = []leanruntime.Field{}

var fields_get_index = []leanruntime.Field{}

var fields_get_managed_view = []leanruntime.Field{
	{Name: "ManagedViewArn", Flag: "managed-view-arn", Type: "*string", Required: true},
}

var fields_get_resource_explorer_setup = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_get_service_index = []leanruntime.Field{}

var fields_get_service_view = []leanruntime.Field{
	{Name: "ServiceViewArn", Flag: "service-view-arn", Type: "*string", Required: true},
}

var fields_get_view = []leanruntime.Field{
	{Name: "ViewArn", Flag: "view-arn", Type: "*string", Required: true},
}

var fields_list_indexes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.IndexType", Required: false},
}

var fields_list_indexes_for_members = []leanruntime.Field{
	{Name: "AccountIdList", Flag: "account-id-list", Type: "[]string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_views = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: false},
}

var fields_list_resources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.SearchFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ViewArn", Flag: "view-arn", Type: "*string", Required: false},
}

var fields_list_service_indexes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
}

var fields_list_service_views = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_streaming_access_for_services = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_supported_resource_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_views = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "ViewArn", Flag: "view-arn", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_index_type = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.IndexType", Required: true},
}

var fields_update_view = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.SearchFilter", Required: false},
	{Name: "IncludedProperties", Flag: "included-properties", Type: "[]types.IncludedProperty", Required: false},
	{Name: "ViewArn", Flag: "view-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-default-view": {
			Name:   "associate-default-view",
			Fields: fields_associate_default_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDefaultViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_default_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDefaultView(ctx, input)
			},
		},
		"batch-get-view": {
			Name:   "batch-get-view",
			Fields: fields_batch_get_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetView(ctx, input)
			},
		},
		"create-index": {
			Name:   "create-index",
			Fields: fields_create_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIndex(ctx, input)
			},
		},
		"create-resource-explorer-setup": {
			Name:   "create-resource-explorer-setup",
			Fields: fields_create_resource_explorer_setup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceExplorerSetupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_explorer_setup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceExplorerSetup(ctx, input)
			},
		},
		"create-view": {
			Name:   "create-view",
			Fields: fields_create_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateView(ctx, input)
			},
		},
		"delete-index": {
			Name:   "delete-index",
			Fields: fields_delete_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndex(ctx, input)
			},
		},
		"delete-resource-explorer-setup": {
			Name:   "delete-resource-explorer-setup",
			Fields: fields_delete_resource_explorer_setup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceExplorerSetupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_explorer_setup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceExplorerSetup(ctx, input)
			},
		},
		"delete-view": {
			Name:   "delete-view",
			Fields: fields_delete_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteView(ctx, input)
			},
		},
		"disassociate-default-view": {
			Name:   "disassociate-default-view",
			Fields: fields_disassociate_default_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDefaultViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_default_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDefaultView(ctx, input)
			},
		},
		"get-account-level-service-configuration": {
			Name:   "get-account-level-service-configuration",
			Fields: fields_get_account_level_service_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountLevelServiceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_level_service_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountLevelServiceConfiguration(ctx, input)
			},
		},
		"get-default-view": {
			Name:   "get-default-view",
			Fields: fields_get_default_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultView(ctx, input)
			},
		},
		"get-index": {
			Name:   "get-index",
			Fields: fields_get_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndex(ctx, input)
			},
		},
		"get-managed-view": {
			Name:   "get-managed-view",
			Fields: fields_get_managed_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedView(ctx, input)
			},
		},
		"get-resource-explorer-setup": {
			Name:   "get-resource-explorer-setup",
			Fields: fields_get_resource_explorer_setup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceExplorerSetupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_explorer_setup, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceExplorerSetup(ctx, input)
				}
				var results []*svc.GetResourceExplorerSetupOutput
				p := svc.NewGetResourceExplorerSetupPaginator(client, input)
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
		"get-service-index": {
			Name:   "get-service-index",
			Fields: fields_get_service_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceIndex(ctx, input)
			},
		},
		"get-service-view": {
			Name:   "get-service-view",
			Fields: fields_get_service_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceView(ctx, input)
			},
		},
		"get-view": {
			Name:   "get-view",
			Fields: fields_get_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetView(ctx, input)
			},
		},
		"list-indexes": {
			Name:   "list-indexes",
			Fields: fields_list_indexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndexesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indexes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndexes(ctx, input)
				}
				var results []*svc.ListIndexesOutput
				p := svc.NewListIndexesPaginator(client, input)
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
		"list-indexes-for-members": {
			Name:   "list-indexes-for-members",
			Fields: fields_list_indexes_for_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndexesForMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indexes_for_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndexesForMembers(ctx, input)
				}
				var results []*svc.ListIndexesForMembersOutput
				p := svc.NewListIndexesForMembersPaginator(client, input)
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
		"list-managed-views": {
			Name:   "list-managed-views",
			Fields: fields_list_managed_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedViews(ctx, input)
				}
				var results []*svc.ListManagedViewsOutput
				p := svc.NewListManagedViewsPaginator(client, input)
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
		"list-resources": {
			Name:   "list-resources",
			Fields: fields_list_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResources(ctx, input)
				}
				var results []*svc.ListResourcesOutput
				p := svc.NewListResourcesPaginator(client, input)
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
		"list-service-indexes": {
			Name:   "list-service-indexes",
			Fields: fields_list_service_indexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceIndexesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_indexes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceIndexes(ctx, input)
				}
				var results []*svc.ListServiceIndexesOutput
				p := svc.NewListServiceIndexesPaginator(client, input)
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
		"list-service-views": {
			Name:   "list-service-views",
			Fields: fields_list_service_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceViews(ctx, input)
				}
				var results []*svc.ListServiceViewsOutput
				p := svc.NewListServiceViewsPaginator(client, input)
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
		"list-streaming-access-for-services": {
			Name:   "list-streaming-access-for-services",
			Fields: fields_list_streaming_access_for_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamingAccessForServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_streaming_access_for_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamingAccessForServices(ctx, input)
				}
				var results []*svc.ListStreamingAccessForServicesOutput
				p := svc.NewListStreamingAccessForServicesPaginator(client, input)
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
		"list-supported-resource-types": {
			Name:   "list-supported-resource-types",
			Fields: fields_list_supported_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSupportedResourceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_supported_resource_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSupportedResourceTypes(ctx, input)
				}
				var results []*svc.ListSupportedResourceTypesOutput
				p := svc.NewListSupportedResourceTypesPaginator(client, input)
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
		"list-views": {
			Name:   "list-views",
			Fields: fields_list_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListViews(ctx, input)
				}
				var results []*svc.ListViewsOutput
				p := svc.NewListViewsPaginator(client, input)
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
		"search": {
			Name:   "search",
			Fields: fields_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Search(ctx, input)
				}
				var results []*svc.SearchOutput
				p := svc.NewSearchPaginator(client, input)
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
		"update-index-type": {
			Name:   "update-index-type",
			Fields: fields_update_index_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_index_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndexType(ctx, input)
			},
		},
		"update-view": {
			Name:   "update-view",
			Fields: fields_update_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateView(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("resourceexplorer2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
