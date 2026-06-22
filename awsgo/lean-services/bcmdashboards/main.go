package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bcmdashboards"
)

var fields_create_dashboard = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "Widgets", Flag: "widgets", Type: "[]types.Widget", Required: true},
}

var fields_delete_dashboard = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_dashboard = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_dashboards = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTagKeys", Flag: "resource-tag-keys", Type: "[]string", Required: true},
}

var fields_update_dashboard = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Widgets", Flag: "widgets", Type: "[]types.Widget", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-dashboard": {
			Name:   "create-dashboard",
			Fields: fields_create_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDashboard(ctx, input)
			},
		},
		"delete-dashboard": {
			Name:   "delete-dashboard",
			Fields: fields_delete_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDashboard(ctx, input)
			},
		},
		"get-dashboard": {
			Name:   "get-dashboard",
			Fields: fields_get_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDashboard(ctx, input)
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
		"list-dashboards": {
			Name:   "list-dashboards",
			Fields: fields_list_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDashboardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dashboards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDashboards(ctx, input)
				}
				var results []*svc.ListDashboardsOutput
				p := svc.NewListDashboardsPaginator(client, input)
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
		"update-dashboard": {
			Name:   "update-dashboard",
			Fields: fields_update_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboard(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bcmdashboards", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
