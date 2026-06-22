package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iottwinmaker"
)

var fields_batch_put_property_values = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.PropertyValueEntry", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_cancel_metadata_transfer_job = []leanruntime.Field{
	{Name: "MetadataTransferJobId", Flag: "metadata-transfer-job-id", Type: "*string", Required: true},
}

var fields_create_component_type = []leanruntime.Field{
	{Name: "ComponentTypeId", Flag: "component-type-id", Type: "*string", Required: true},
	{Name: "ComponentTypeName", Flag: "component-type-name", Type: "*string", Required: false},
	{Name: "CompositeComponentTypes", Flag: "composite-component-types", Type: "map[string]types.CompositeComponentTypeRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExtendsFrom", Flag: "extends-from", Type: "[]string", Required: false},
	{Name: "Functions", Flag: "functions", Type: "map[string]types.FunctionRequest", Required: false},
	{Name: "IsSingleton", Flag: "is-singleton", Type: "*bool", Required: false},
	{Name: "PropertyDefinitions", Flag: "property-definitions", Type: "map[string]types.PropertyDefinitionRequest", Required: false},
	{Name: "PropertyGroups", Flag: "property-groups", Type: "map[string]types.PropertyGroupRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_entity = []leanruntime.Field{
	{Name: "Components", Flag: "components", Type: "map[string]types.ComponentRequest", Required: false},
	{Name: "CompositeComponents", Flag: "composite-components", Type: "map[string]types.CompositeComponentRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: false},
	{Name: "EntityName", Flag: "entity-name", Type: "*string", Required: true},
	{Name: "ParentEntityId", Flag: "parent-entity-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_metadata_transfer_job = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*types.DestinationConfiguration", Required: true},
	{Name: "MetadataTransferJobId", Flag: "metadata-transfer-job-id", Type: "*string", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.SourceConfiguration", Required: true},
}

var fields_create_scene = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]string", Required: false},
	{Name: "ContentLocation", Flag: "content-location", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SceneId", Flag: "scene-id", Type: "*string", Required: true},
	{Name: "SceneMetadata", Flag: "scene-metadata", Type: "map[string]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_sync_job = []leanruntime.Field{
	{Name: "SyncRole", Flag: "sync-role", Type: "*string", Required: true},
	{Name: "SyncSource", Flag: "sync-source", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_workspace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "S3Location", Flag: "s3-location", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_component_type = []leanruntime.Field{
	{Name: "ComponentTypeId", Flag: "component-type-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_entity = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "IsRecursive", Flag: "is-recursive", Type: "*bool", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_scene = []leanruntime.Field{
	{Name: "SceneId", Flag: "scene-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_sync_job = []leanruntime.Field{
	{Name: "SyncSource", Flag: "sync-source", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_execute_query = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryStatement", Flag: "query-statement", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_component_type = []leanruntime.Field{
	{Name: "ComponentTypeId", Flag: "component-type-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_entity = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_metadata_transfer_job = []leanruntime.Field{
	{Name: "MetadataTransferJobId", Flag: "metadata-transfer-job-id", Type: "*string", Required: true},
}

var fields_get_pricing_plan = []leanruntime.Field{}

var fields_get_property_value = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: false},
	{Name: "ComponentPath", Flag: "component-path", Type: "*string", Required: false},
	{Name: "ComponentTypeId", Flag: "component-type-id", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PropertyGroupName", Flag: "property-group-name", Type: "*string", Required: false},
	{Name: "SelectedProperties", Flag: "selected-properties", Type: "[]string", Required: true},
	{Name: "TabularConditions", Flag: "tabular-conditions", Type: "*types.TabularConditions", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_property_value_history = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: false},
	{Name: "ComponentPath", Flag: "component-path", Type: "*string", Required: false},
	{Name: "ComponentTypeId", Flag: "component-type-id", Type: "*string", Required: false},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*time.Time", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: false},
	{Name: "Interpolation", Flag: "interpolation", Type: "*types.InterpolationParameters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderByTime", Flag: "order-by-time", Type: "types.OrderByTime", Required: false},
	{Name: "PropertyFilters", Flag: "property-filters", Type: "[]types.PropertyFilter", Required: false},
	{Name: "SelectedProperties", Flag: "selected-properties", Type: "[]string", Required: true},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*time.Time", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_scene = []leanruntime.Field{
	{Name: "SceneId", Flag: "scene-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_sync_job = []leanruntime.Field{
	{Name: "SyncSource", Flag: "sync-source", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: false},
}

var fields_get_workspace = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_component_types = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListComponentTypesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_components = []leanruntime.Field{
	{Name: "ComponentPath", Flag: "component-path", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_entities = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListEntitiesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_metadata_transfer_jobs = []leanruntime.Field{
	{Name: "DestinationType", Flag: "destination-type", Type: "types.DestinationType", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListMetadataTransferJobsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.SourceType", Required: true},
}

var fields_list_properties = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: false},
	{Name: "ComponentPath", Flag: "component-path", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_scenes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_sync_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_sync_resources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SyncResourceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SyncSource", Flag: "sync-source", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_workspaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_component_type = []leanruntime.Field{
	{Name: "ComponentTypeId", Flag: "component-type-id", Type: "*string", Required: true},
	{Name: "ComponentTypeName", Flag: "component-type-name", Type: "*string", Required: false},
	{Name: "CompositeComponentTypes", Flag: "composite-component-types", Type: "map[string]types.CompositeComponentTypeRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExtendsFrom", Flag: "extends-from", Type: "[]string", Required: false},
	{Name: "Functions", Flag: "functions", Type: "map[string]types.FunctionRequest", Required: false},
	{Name: "IsSingleton", Flag: "is-singleton", Type: "*bool", Required: false},
	{Name: "PropertyDefinitions", Flag: "property-definitions", Type: "map[string]types.PropertyDefinitionRequest", Required: false},
	{Name: "PropertyGroups", Flag: "property-groups", Type: "map[string]types.PropertyGroupRequest", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_entity = []leanruntime.Field{
	{Name: "ComponentUpdates", Flag: "component-updates", Type: "map[string]types.ComponentUpdateRequest", Required: false},
	{Name: "CompositeComponentUpdates", Flag: "composite-component-updates", Type: "map[string]types.CompositeComponentUpdateRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "EntityName", Flag: "entity-name", Type: "*string", Required: false},
	{Name: "ParentEntityUpdate", Flag: "parent-entity-update", Type: "*types.ParentEntityUpdateRequest", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_pricing_plan = []leanruntime.Field{
	{Name: "BundleNames", Flag: "bundle-names", Type: "[]string", Required: false},
	{Name: "PricingMode", Flag: "pricing-mode", Type: "types.PricingMode", Required: true},
}

var fields_update_scene = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]string", Required: false},
	{Name: "ContentLocation", Flag: "content-location", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SceneId", Flag: "scene-id", Type: "*string", Required: true},
	{Name: "SceneMetadata", Flag: "scene-metadata", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "S3Location", Flag: "s3-location", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-put-property-values": {
			Name:   "batch-put-property-values",
			Fields: fields_batch_put_property_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutPropertyValuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_property_values, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutPropertyValues(ctx, input)
			},
		},
		"cancel-metadata-transfer-job": {
			Name:   "cancel-metadata-transfer-job",
			Fields: fields_cancel_metadata_transfer_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMetadataTransferJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_metadata_transfer_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMetadataTransferJob(ctx, input)
			},
		},
		"create-component-type": {
			Name:   "create-component-type",
			Fields: fields_create_component_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComponentTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_component_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComponentType(ctx, input)
			},
		},
		"create-entity": {
			Name:   "create-entity",
			Fields: fields_create_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEntity(ctx, input)
			},
		},
		"create-metadata-transfer-job": {
			Name:   "create-metadata-transfer-job",
			Fields: fields_create_metadata_transfer_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMetadataTransferJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_metadata_transfer_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMetadataTransferJob(ctx, input)
			},
		},
		"create-scene": {
			Name:   "create-scene",
			Fields: fields_create_scene,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSceneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scene, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScene(ctx, input)
			},
		},
		"create-sync-job": {
			Name:   "create-sync-job",
			Fields: fields_create_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSyncJob(ctx, input)
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
		"delete-component-type": {
			Name:   "delete-component-type",
			Fields: fields_delete_component_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComponentTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_component_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComponentType(ctx, input)
			},
		},
		"delete-entity": {
			Name:   "delete-entity",
			Fields: fields_delete_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEntity(ctx, input)
			},
		},
		"delete-scene": {
			Name:   "delete-scene",
			Fields: fields_delete_scene,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSceneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scene, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScene(ctx, input)
			},
		},
		"delete-sync-job": {
			Name:   "delete-sync-job",
			Fields: fields_delete_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSyncJob(ctx, input)
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
		"execute-query": {
			Name:   "execute-query",
			Fields: fields_execute_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteQueryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_execute_query, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ExecuteQuery(ctx, input)
				}
				var results []*svc.ExecuteQueryOutput
				p := svc.NewExecuteQueryPaginator(client, input)
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
		"get-component-type": {
			Name:   "get-component-type",
			Fields: fields_get_component_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComponentTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_component_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComponentType(ctx, input)
			},
		},
		"get-entity": {
			Name:   "get-entity",
			Fields: fields_get_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEntity(ctx, input)
			},
		},
		"get-metadata-transfer-job": {
			Name:   "get-metadata-transfer-job",
			Fields: fields_get_metadata_transfer_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetadataTransferJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metadata_transfer_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetadataTransferJob(ctx, input)
			},
		},
		"get-pricing-plan": {
			Name:   "get-pricing-plan",
			Fields: fields_get_pricing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPricingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pricing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPricingPlan(ctx, input)
			},
		},
		"get-property-value": {
			Name:   "get-property-value",
			Fields: fields_get_property_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPropertyValueInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_property_value, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPropertyValue(ctx, input)
				}
				var results []*svc.GetPropertyValueOutput
				p := svc.NewGetPropertyValuePaginator(client, input)
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
		"get-property-value-history": {
			Name:   "get-property-value-history",
			Fields: fields_get_property_value_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPropertyValueHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_property_value_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPropertyValueHistory(ctx, input)
				}
				var results []*svc.GetPropertyValueHistoryOutput
				p := svc.NewGetPropertyValueHistoryPaginator(client, input)
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
		"get-scene": {
			Name:   "get-scene",
			Fields: fields_get_scene,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSceneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scene, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScene(ctx, input)
			},
		},
		"get-sync-job": {
			Name:   "get-sync-job",
			Fields: fields_get_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSyncJob(ctx, input)
			},
		},
		"get-workspace": {
			Name:   "get-workspace",
			Fields: fields_get_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkspace(ctx, input)
			},
		},
		"list-component-types": {
			Name:   "list-component-types",
			Fields: fields_list_component_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_component_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponentTypes(ctx, input)
				}
				var results []*svc.ListComponentTypesOutput
				p := svc.NewListComponentTypesPaginator(client, input)
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
		"list-components": {
			Name:   "list-components",
			Fields: fields_list_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponents(ctx, input)
				}
				var results []*svc.ListComponentsOutput
				p := svc.NewListComponentsPaginator(client, input)
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
		"list-entities": {
			Name:   "list-entities",
			Fields: fields_list_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntities(ctx, input)
				}
				var results []*svc.ListEntitiesOutput
				p := svc.NewListEntitiesPaginator(client, input)
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
		"list-metadata-transfer-jobs": {
			Name:   "list-metadata-transfer-jobs",
			Fields: fields_list_metadata_transfer_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetadataTransferJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metadata_transfer_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetadataTransferJobs(ctx, input)
				}
				var results []*svc.ListMetadataTransferJobsOutput
				p := svc.NewListMetadataTransferJobsPaginator(client, input)
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
		"list-properties": {
			Name:   "list-properties",
			Fields: fields_list_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPropertiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_properties, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProperties(ctx, input)
				}
				var results []*svc.ListPropertiesOutput
				p := svc.NewListPropertiesPaginator(client, input)
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
		"list-scenes": {
			Name:   "list-scenes",
			Fields: fields_list_scenes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScenesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scenes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScenes(ctx, input)
				}
				var results []*svc.ListScenesOutput
				p := svc.NewListScenesPaginator(client, input)
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
		"list-sync-jobs": {
			Name:   "list-sync-jobs",
			Fields: fields_list_sync_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSyncJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sync_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSyncJobs(ctx, input)
				}
				var results []*svc.ListSyncJobsOutput
				p := svc.NewListSyncJobsPaginator(client, input)
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
		"list-sync-resources": {
			Name:   "list-sync-resources",
			Fields: fields_list_sync_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSyncResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sync_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSyncResources(ctx, input)
				}
				var results []*svc.ListSyncResourcesOutput
				p := svc.NewListSyncResourcesPaginator(client, input)
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
		"update-component-type": {
			Name:   "update-component-type",
			Fields: fields_update_component_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComponentTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_component_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComponentType(ctx, input)
			},
		},
		"update-entity": {
			Name:   "update-entity",
			Fields: fields_update_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEntity(ctx, input)
			},
		},
		"update-pricing-plan": {
			Name:   "update-pricing-plan",
			Fields: fields_update_pricing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePricingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pricing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePricingPlan(ctx, input)
			},
		},
		"update-scene": {
			Name:   "update-scene",
			Fields: fields_update_scene,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSceneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scene, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScene(ctx, input)
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
	}
	if err := leanruntime.Execute("iottwinmaker", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
