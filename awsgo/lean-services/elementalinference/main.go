package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/elementalinference"
)

var fields_associate_feed = []leanruntime.Field{
	{Name: "AssociatedResourceName", Flag: "associated-resource-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.CreateOutput", Required: true},
}

var fields_create_feed = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.CreateOutput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_feed = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_disassociate_feed = []leanruntime.Field{
	{Name: "AssociatedResourceName", Flag: "associated-resource-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_feed = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_feeds = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_feed = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.UpdateOutput", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-feed": {
			Name:   "associate-feed",
			Fields: fields_associate_feed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFeedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_feed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFeed(ctx, input)
			},
		},
		"create-feed": {
			Name:   "create-feed",
			Fields: fields_create_feed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFeedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_feed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFeed(ctx, input)
			},
		},
		"delete-feed": {
			Name:   "delete-feed",
			Fields: fields_delete_feed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFeedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_feed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFeed(ctx, input)
			},
		},
		"disassociate-feed": {
			Name:   "disassociate-feed",
			Fields: fields_disassociate_feed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFeedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_feed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFeed(ctx, input)
			},
		},
		"get-feed": {
			Name:   "get-feed",
			Fields: fields_get_feed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFeedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_feed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFeed(ctx, input)
			},
		},
		"list-feeds": {
			Name:   "list-feeds",
			Fields: fields_list_feeds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFeedsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_feeds, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFeeds(ctx, input)
				}
				var results []*svc.ListFeedsOutput
				p := svc.NewListFeedsPaginator(client, input)
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
		"update-feed": {
			Name:   "update-feed",
			Fields: fields_update_feed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFeedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_feed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFeed(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("elementalinference", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
