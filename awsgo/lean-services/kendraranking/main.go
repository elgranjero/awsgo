package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kendraranking"
)

var fields_create_rescore_execution_plan = []leanruntime.Field{
	{Name: "CapacityUnits", Flag: "capacity-units", Type: "*types.CapacityUnitsConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_rescore_execution_plan = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_rescore_execution_plan = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_rescore_execution_plans = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_rescore = []leanruntime.Field{
	{Name: "Documents", Flag: "documents", Type: "[]types.Document", Required: true},
	{Name: "RescoreExecutionPlanId", Flag: "rescore-execution-plan-id", Type: "*string", Required: true},
	{Name: "SearchQuery", Flag: "search-query", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_rescore_execution_plan = []leanruntime.Field{
	{Name: "CapacityUnits", Flag: "capacity-units", Type: "*types.CapacityUnitsConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-rescore-execution-plan": {
			Name:   "create-rescore-execution-plan",
			Fields: fields_create_rescore_execution_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRescoreExecutionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rescore_execution_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRescoreExecutionPlan(ctx, input)
			},
		},
		"delete-rescore-execution-plan": {
			Name:   "delete-rescore-execution-plan",
			Fields: fields_delete_rescore_execution_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRescoreExecutionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rescore_execution_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRescoreExecutionPlan(ctx, input)
			},
		},
		"describe-rescore-execution-plan": {
			Name:   "describe-rescore-execution-plan",
			Fields: fields_describe_rescore_execution_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRescoreExecutionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rescore_execution_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRescoreExecutionPlan(ctx, input)
			},
		},
		"list-rescore-execution-plans": {
			Name:   "list-rescore-execution-plans",
			Fields: fields_list_rescore_execution_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRescoreExecutionPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rescore_execution_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRescoreExecutionPlans(ctx, input)
				}
				var results []*svc.ListRescoreExecutionPlansOutput
				p := svc.NewListRescoreExecutionPlansPaginator(client, input)
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
		"rescore": {
			Name:   "rescore",
			Fields: fields_rescore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RescoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rescore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Rescore(ctx, input)
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
		"update-rescore-execution-plan": {
			Name:   "update-rescore-execution-plan",
			Fields: fields_update_rescore_execution_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRescoreExecutionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rescore_execution_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRescoreExecutionPlan(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kendraranking", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
