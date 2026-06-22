package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/osis"
)

var fields_create_pipeline = []leanruntime.Field{
	{Name: "BufferOptions", Flag: "buffer-options", Type: "*types.BufferOptions", Required: false},
	{Name: "EncryptionAtRestOptions", Flag: "encryption-at-rest-options", Type: "*types.EncryptionAtRestOptions", Required: false},
	{Name: "LogPublishingOptions", Flag: "log-publishing-options", Type: "*types.LogPublishingOptions", Required: false},
	{Name: "MaxUnits", Flag: "max-units", Type: "*int32", Required: true},
	{Name: "MinUnits", Flag: "min-units", Type: "*int32", Required: true},
	{Name: "PipelineConfigurationBody", Flag: "pipeline-configuration-body", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "PipelineRoleArn", Flag: "pipeline-role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcOptions", Flag: "vpc-options", Type: "*types.VpcOptions", Required: false},
}

var fields_create_pipeline_endpoint = []leanruntime.Field{
	{Name: "PipelineArn", Flag: "pipeline-arn", Type: "*string", Required: true},
	{Name: "VpcOptions", Flag: "vpc-options", Type: "*types.PipelineEndpointVpcOptions", Required: true},
}

var fields_delete_pipeline = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_delete_pipeline_endpoint = []leanruntime.Field{
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_pipeline = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_get_pipeline_blueprint = []leanruntime.Field{
	{Name: "BlueprintName", Flag: "blueprint-name", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "*string", Required: false},
}

var fields_get_pipeline_change_progress = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_pipeline_blueprints = []leanruntime.Field{}

var fields_list_pipeline_endpoint_connections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pipeline_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pipelines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_revoke_pipeline_endpoint_connections = []leanruntime.Field{
	{Name: "EndpointIds", Flag: "endpoint-ids", Type: "[]string", Required: true},
	{Name: "PipelineArn", Flag: "pipeline-arn", Type: "*string", Required: true},
}

var fields_start_pipeline = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_stop_pipeline = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_pipeline = []leanruntime.Field{
	{Name: "BufferOptions", Flag: "buffer-options", Type: "*types.BufferOptions", Required: false},
	{Name: "EncryptionAtRestOptions", Flag: "encryption-at-rest-options", Type: "*types.EncryptionAtRestOptions", Required: false},
	{Name: "LogPublishingOptions", Flag: "log-publishing-options", Type: "*types.LogPublishingOptions", Required: false},
	{Name: "MaxUnits", Flag: "max-units", Type: "*int32", Required: false},
	{Name: "MinUnits", Flag: "min-units", Type: "*int32", Required: false},
	{Name: "PipelineConfigurationBody", Flag: "pipeline-configuration-body", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "PipelineRoleArn", Flag: "pipeline-role-arn", Type: "*string", Required: false},
}

var fields_validate_pipeline = []leanruntime.Field{
	{Name: "PipelineConfigurationBody", Flag: "pipeline-configuration-body", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-pipeline": {
			Name:   "create-pipeline",
			Fields: fields_create_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePipeline(ctx, input)
			},
		},
		"create-pipeline-endpoint": {
			Name:   "create-pipeline-endpoint",
			Fields: fields_create_pipeline_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePipelineEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pipeline_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePipelineEndpoint(ctx, input)
			},
		},
		"delete-pipeline": {
			Name:   "delete-pipeline",
			Fields: fields_delete_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePipeline(ctx, input)
			},
		},
		"delete-pipeline-endpoint": {
			Name:   "delete-pipeline-endpoint",
			Fields: fields_delete_pipeline_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePipelineEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pipeline_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePipelineEndpoint(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"get-pipeline": {
			Name:   "get-pipeline",
			Fields: fields_get_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipeline(ctx, input)
			},
		},
		"get-pipeline-blueprint": {
			Name:   "get-pipeline-blueprint",
			Fields: fields_get_pipeline_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipelineBlueprint(ctx, input)
			},
		},
		"get-pipeline-change-progress": {
			Name:   "get-pipeline-change-progress",
			Fields: fields_get_pipeline_change_progress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineChangeProgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline_change_progress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipelineChangeProgress(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"list-pipeline-blueprints": {
			Name:   "list-pipeline-blueprints",
			Fields: fields_list_pipeline_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineBlueprintsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_pipeline_blueprints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPipelineBlueprints(ctx, input)
			},
		},
		"list-pipeline-endpoint-connections": {
			Name:   "list-pipeline-endpoint-connections",
			Fields: fields_list_pipeline_endpoint_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineEndpointConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_endpoint_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineEndpointConnections(ctx, input)
				}
				var results []*svc.ListPipelineEndpointConnectionsOutput
				p := svc.NewListPipelineEndpointConnectionsPaginator(client, input)
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
		"list-pipeline-endpoints": {
			Name:   "list-pipeline-endpoints",
			Fields: fields_list_pipeline_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineEndpoints(ctx, input)
				}
				var results []*svc.ListPipelineEndpointsOutput
				p := svc.NewListPipelineEndpointsPaginator(client, input)
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
		"list-pipelines": {
			Name:   "list-pipelines",
			Fields: fields_list_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelines(ctx, input)
				}
				var results []*svc.ListPipelinesOutput
				p := svc.NewListPipelinesPaginator(client, input)
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
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"revoke-pipeline-endpoint-connections": {
			Name:   "revoke-pipeline-endpoint-connections",
			Fields: fields_revoke_pipeline_endpoint_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokePipelineEndpointConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_pipeline_endpoint_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokePipelineEndpointConnections(ctx, input)
			},
		},
		"start-pipeline": {
			Name:   "start-pipeline",
			Fields: fields_start_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPipeline(ctx, input)
			},
		},
		"stop-pipeline": {
			Name:   "stop-pipeline",
			Fields: fields_stop_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopPipeline(ctx, input)
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
		"update-pipeline": {
			Name:   "update-pipeline",
			Fields: fields_update_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePipeline(ctx, input)
			},
		},
		"validate-pipeline": {
			Name:   "validate-pipeline",
			Fields: fields_validate_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidatePipeline(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("osis", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
