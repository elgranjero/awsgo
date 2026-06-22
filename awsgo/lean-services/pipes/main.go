package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pipes"
)

var fields_create_pipe = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DesiredState", Flag: "desired-state", Type: "types.RequestedPipeState", Required: false},
	{Name: "Enrichment", Flag: "enrichment", Type: "*string", Required: false},
	{Name: "EnrichmentParameters", Flag: "enrichment-parameters", Type: "*types.PipeEnrichmentParameters", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.PipeLogConfigurationParameters", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "SourceParameters", Flag: "source-parameters", Type: "*types.PipeSourceParameters", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
	{Name: "TargetParameters", Flag: "target-parameters", Type: "*types.PipeTargetParameters", Required: false},
}

var fields_delete_pipe = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_pipe = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_pipes = []leanruntime.Field{
	{Name: "CurrentState", Flag: "current-state", Type: "types.PipeState", Required: false},
	{Name: "DesiredState", Flag: "desired-state", Type: "types.RequestedPipeState", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourcePrefix", Flag: "source-prefix", Type: "*string", Required: false},
	{Name: "TargetPrefix", Flag: "target-prefix", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_pipe = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_pipe = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_pipe = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DesiredState", Flag: "desired-state", Type: "types.RequestedPipeState", Required: false},
	{Name: "Enrichment", Flag: "enrichment", Type: "*string", Required: false},
	{Name: "EnrichmentParameters", Flag: "enrichment-parameters", Type: "*types.PipeEnrichmentParameters", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.PipeLogConfigurationParameters", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SourceParameters", Flag: "source-parameters", Type: "*types.UpdatePipeSourceParameters", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: false},
	{Name: "TargetParameters", Flag: "target-parameters", Type: "*types.PipeTargetParameters", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-pipe": {
			Name:   "create-pipe",
			Fields: fields_create_pipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePipe(ctx, input)
			},
		},
		"delete-pipe": {
			Name:   "delete-pipe",
			Fields: fields_delete_pipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePipe(ctx, input)
			},
		},
		"describe-pipe": {
			Name:   "describe-pipe",
			Fields: fields_describe_pipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePipe(ctx, input)
			},
		},
		"list-pipes": {
			Name:   "list-pipes",
			Fields: fields_list_pipes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipes(ctx, input)
				}
				var results []*svc.ListPipesOutput
				p := svc.NewListPipesPaginator(client, input)
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
		"start-pipe": {
			Name:   "start-pipe",
			Fields: fields_start_pipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_pipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPipe(ctx, input)
			},
		},
		"stop-pipe": {
			Name:   "stop-pipe",
			Fields: fields_stop_pipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopPipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_pipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopPipe(ctx, input)
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
		"update-pipe": {
			Name:   "update-pipe",
			Fields: fields_update_pipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePipe(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pipes", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
