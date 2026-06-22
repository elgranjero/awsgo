package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/billing"
)

var fields_associate_source_views = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "SourceViews", Flag: "source-views", Type: "[]string", Required: true},
}

var fields_create_billing_view = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataFilterExpression", Flag: "data-filter-expression", Type: "*types.Expression", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "SourceViews", Flag: "source-views", Type: "[]string", Required: true},
}

var fields_delete_billing_view = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
}

var fields_disassociate_source_views = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "SourceViews", Flag: "source-views", Type: "[]string", Required: true},
}

var fields_get_billing_view = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_billing_views = []leanruntime.Field{
	{Name: "ActiveTimeRange", Flag: "active-time-range", Type: "*types.ActiveTimeRange", Required: false},
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: false},
	{Name: "BillingViewTypes", Flag: "billing-view-types", Type: "[]types.BillingViewType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]types.StringSearch", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerAccountId", Flag: "owner-account-id", Type: "*string", Required: false},
	{Name: "SourceAccountId", Flag: "source-account-id", Type: "*string", Required: false},
}

var fields_list_source_views_for_billing_view = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
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

var fields_update_billing_view = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "DataFilterExpression", Flag: "data-filter-expression", Type: "*types.Expression", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-source-views": {
			Name:   "associate-source-views",
			Fields: fields_associate_source_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSourceViewsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_source_views, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSourceViews(ctx, input)
			},
		},
		"create-billing-view": {
			Name:   "create-billing-view",
			Fields: fields_create_billing_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBillingViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_billing_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBillingView(ctx, input)
			},
		},
		"delete-billing-view": {
			Name:   "delete-billing-view",
			Fields: fields_delete_billing_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBillingViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_billing_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBillingView(ctx, input)
			},
		},
		"disassociate-source-views": {
			Name:   "disassociate-source-views",
			Fields: fields_disassociate_source_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSourceViewsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_source_views, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSourceViews(ctx, input)
			},
		},
		"get-billing-view": {
			Name:   "get-billing-view",
			Fields: fields_get_billing_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBillingViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_billing_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBillingView(ctx, input)
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
		"list-billing-views": {
			Name:   "list-billing-views",
			Fields: fields_list_billing_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillingViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_billing_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillingViews(ctx, input)
				}
				var results []*svc.ListBillingViewsOutput
				p := svc.NewListBillingViewsPaginator(client, input)
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
		"list-source-views-for-billing-view": {
			Name:   "list-source-views-for-billing-view",
			Fields: fields_list_source_views_for_billing_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceViewsForBillingViewInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_views_for_billing_view, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceViewsForBillingView(ctx, input)
				}
				var results []*svc.ListSourceViewsForBillingViewOutput
				p := svc.NewListSourceViewsForBillingViewPaginator(client, input)
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
		"update-billing-view": {
			Name:   "update-billing-view",
			Fields: fields_update_billing_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBillingViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_billing_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBillingView(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("billing", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
