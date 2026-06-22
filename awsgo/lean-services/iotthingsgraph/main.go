package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
)

var fields_associate_entity_to_thing = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "NamespaceVersion", Flag: "namespace-version", Type: "*int64", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_create_flow_template = []leanruntime.Field{
	{Name: "CompatibleNamespaceVersion", Flag: "compatible-namespace-version", Type: "*int64", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.DefinitionDocument", Required: true},
}

var fields_create_system_instance = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "*types.DefinitionDocument", Required: true},
	{Name: "FlowActionsRoleArn", Flag: "flow-actions-role-arn", Type: "*string", Required: false},
	{Name: "GreengrassGroupName", Flag: "greengrass-group-name", Type: "*string", Required: false},
	{Name: "MetricsConfiguration", Flag: "metrics-configuration", Type: "*types.MetricsConfiguration", Required: false},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Target", Flag: "target", Type: "types.DeploymentTarget", Required: true},
}

var fields_create_system_template = []leanruntime.Field{
	{Name: "CompatibleNamespaceVersion", Flag: "compatible-namespace-version", Type: "*int64", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.DefinitionDocument", Required: true},
}

var fields_delete_flow_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_namespace = []leanruntime.Field{}

var fields_delete_system_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
}

var fields_delete_system_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_deploy_system_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
}

var fields_deprecate_flow_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_deprecate_system_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_namespace = []leanruntime.Field{
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: false},
}

var fields_dissociate_entity_from_thing = []leanruntime.Field{
	{Name: "EntityType", Flag: "entity-type", Type: "types.EntityType", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_get_entities = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
	{Name: "NamespaceVersion", Flag: "namespace-version", Type: "*int64", Required: false},
}

var fields_get_flow_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RevisionNumber", Flag: "revision-number", Type: "*int64", Required: false},
}

var fields_get_flow_template_revisions = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_namespace_deletion_status = []leanruntime.Field{}

var fields_get_system_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_system_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RevisionNumber", Flag: "revision-number", Type: "*int64", Required: false},
}

var fields_get_system_template_revisions = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_upload_status = []leanruntime.Field{
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_list_flow_execution_messages = []leanruntime.Field{
	{Name: "FlowExecutionId", Flag: "flow-execution-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_search_entities = []leanruntime.Field{
	{Name: "EntityTypes", Flag: "entity-types", Type: "[]types.EntityType", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.EntityFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceVersion", Flag: "namespace-version", Type: "*int64", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_flow_executions = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "FlowExecutionId", Flag: "flow-execution-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "SystemInstanceId", Flag: "system-instance-id", Type: "*string", Required: true},
}

var fields_search_flow_templates = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.FlowTemplateFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_system_instances = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SystemInstanceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_system_templates = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SystemTemplateFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_things = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceVersion", Flag: "namespace-version", Type: "*int64", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_undeploy_system_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_flow_template = []leanruntime.Field{
	{Name: "CompatibleNamespaceVersion", Flag: "compatible-namespace-version", Type: "*int64", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.DefinitionDocument", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_system_template = []leanruntime.Field{
	{Name: "CompatibleNamespaceVersion", Flag: "compatible-namespace-version", Type: "*int64", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.DefinitionDocument", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_upload_entity_definitions = []leanruntime.Field{
	{Name: "DeprecateExistingEntities", Flag: "deprecate-existing-entities", Type: "bool", Required: false},
	{Name: "Document", Flag: "document", Type: "*types.DefinitionDocument", Required: false},
	{Name: "SyncWithPublicNamespace", Flag: "sync-with-public-namespace", Type: "bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-entity-to-thing": {
			Name:   "associate-entity-to-thing",
			Fields: fields_associate_entity_to_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEntityToThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_entity_to_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEntityToThing(ctx, input)
			},
		},
		"create-flow-template": {
			Name:   "create-flow-template",
			Fields: fields_create_flow_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlowTemplate(ctx, input)
			},
		},
		"create-system-instance": {
			Name:   "create-system-instance",
			Fields: fields_create_system_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSystemInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_system_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSystemInstance(ctx, input)
			},
		},
		"create-system-template": {
			Name:   "create-system-template",
			Fields: fields_create_system_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSystemTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_system_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSystemTemplate(ctx, input)
			},
		},
		"delete-flow-template": {
			Name:   "delete-flow-template",
			Fields: fields_delete_flow_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlowTemplate(ctx, input)
			},
		},
		"delete-namespace": {
			Name:   "delete-namespace",
			Fields: fields_delete_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNamespace(ctx, input)
			},
		},
		"delete-system-instance": {
			Name:   "delete-system-instance",
			Fields: fields_delete_system_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSystemInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_system_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSystemInstance(ctx, input)
			},
		},
		"delete-system-template": {
			Name:   "delete-system-template",
			Fields: fields_delete_system_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSystemTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_system_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSystemTemplate(ctx, input)
			},
		},
		"deploy-system-instance": {
			Name:   "deploy-system-instance",
			Fields: fields_deploy_system_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeploySystemInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deploy_system_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeploySystemInstance(ctx, input)
			},
		},
		"deprecate-flow-template": {
			Name:   "deprecate-flow-template",
			Fields: fields_deprecate_flow_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprecateFlowTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprecate_flow_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprecateFlowTemplate(ctx, input)
			},
		},
		"deprecate-system-template": {
			Name:   "deprecate-system-template",
			Fields: fields_deprecate_system_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprecateSystemTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprecate_system_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprecateSystemTemplate(ctx, input)
			},
		},
		"describe-namespace": {
			Name:   "describe-namespace",
			Fields: fields_describe_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNamespace(ctx, input)
			},
		},
		"dissociate-entity-from-thing": {
			Name:   "dissociate-entity-from-thing",
			Fields: fields_dissociate_entity_from_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DissociateEntityFromThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dissociate_entity_from_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DissociateEntityFromThing(ctx, input)
			},
		},
		"get-entities": {
			Name:   "get-entities",
			Fields: fields_get_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEntities(ctx, input)
			},
		},
		"get-flow-template": {
			Name:   "get-flow-template",
			Fields: fields_get_flow_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowTemplate(ctx, input)
			},
		},
		"get-flow-template-revisions": {
			Name:   "get-flow-template-revisions",
			Fields: fields_get_flow_template_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowTemplateRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_flow_template_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFlowTemplateRevisions(ctx, input)
				}
				var results []*svc.GetFlowTemplateRevisionsOutput
				p := svc.NewGetFlowTemplateRevisionsPaginator(client, input)
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
		"get-namespace-deletion-status": {
			Name:   "get-namespace-deletion-status",
			Fields: fields_get_namespace_deletion_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNamespaceDeletionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_namespace_deletion_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNamespaceDeletionStatus(ctx, input)
			},
		},
		"get-system-instance": {
			Name:   "get-system-instance",
			Fields: fields_get_system_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSystemInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_system_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSystemInstance(ctx, input)
			},
		},
		"get-system-template": {
			Name:   "get-system-template",
			Fields: fields_get_system_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSystemTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_system_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSystemTemplate(ctx, input)
			},
		},
		"get-system-template-revisions": {
			Name:   "get-system-template-revisions",
			Fields: fields_get_system_template_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSystemTemplateRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_system_template_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSystemTemplateRevisions(ctx, input)
				}
				var results []*svc.GetSystemTemplateRevisionsOutput
				p := svc.NewGetSystemTemplateRevisionsPaginator(client, input)
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
		"get-upload-status": {
			Name:   "get-upload-status",
			Fields: fields_get_upload_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUploadStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_upload_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUploadStatus(ctx, input)
			},
		},
		"list-flow-execution-messages": {
			Name:   "list-flow-execution-messages",
			Fields: fields_list_flow_execution_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowExecutionMessagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_execution_messages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowExecutionMessages(ctx, input)
				}
				var results []*svc.ListFlowExecutionMessagesOutput
				p := svc.NewListFlowExecutionMessagesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"search-entities": {
			Name:   "search-entities",
			Fields: fields_search_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchEntities(ctx, input)
				}
				var results []*svc.SearchEntitiesOutput
				p := svc.NewSearchEntitiesPaginator(client, input)
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
		"search-flow-executions": {
			Name:   "search-flow-executions",
			Fields: fields_search_flow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchFlowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_flow_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchFlowExecutions(ctx, input)
				}
				var results []*svc.SearchFlowExecutionsOutput
				p := svc.NewSearchFlowExecutionsPaginator(client, input)
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
		"search-flow-templates": {
			Name:   "search-flow-templates",
			Fields: fields_search_flow_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchFlowTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_flow_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchFlowTemplates(ctx, input)
				}
				var results []*svc.SearchFlowTemplatesOutput
				p := svc.NewSearchFlowTemplatesPaginator(client, input)
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
		"search-system-instances": {
			Name:   "search-system-instances",
			Fields: fields_search_system_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSystemInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_system_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchSystemInstances(ctx, input)
				}
				var results []*svc.SearchSystemInstancesOutput
				p := svc.NewSearchSystemInstancesPaginator(client, input)
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
		"search-system-templates": {
			Name:   "search-system-templates",
			Fields: fields_search_system_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSystemTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_system_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchSystemTemplates(ctx, input)
				}
				var results []*svc.SearchSystemTemplatesOutput
				p := svc.NewSearchSystemTemplatesPaginator(client, input)
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
		"search-things": {
			Name:   "search-things",
			Fields: fields_search_things,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchThingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_things, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchThings(ctx, input)
				}
				var results []*svc.SearchThingsOutput
				p := svc.NewSearchThingsPaginator(client, input)
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
		"undeploy-system-instance": {
			Name:   "undeploy-system-instance",
			Fields: fields_undeploy_system_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UndeploySystemInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_undeploy_system_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UndeploySystemInstance(ctx, input)
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
		"update-flow-template": {
			Name:   "update-flow-template",
			Fields: fields_update_flow_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowTemplate(ctx, input)
			},
		},
		"update-system-template": {
			Name:   "update-system-template",
			Fields: fields_update_system_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSystemTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_system_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSystemTemplate(ctx, input)
			},
		},
		"upload-entity-definitions": {
			Name:   "upload-entity-definitions",
			Fields: fields_upload_entity_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadEntityDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_entity_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadEntityDefinitions(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotthingsgraph", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
