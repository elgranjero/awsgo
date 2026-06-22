package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/oam"
)

var fields_create_link = []leanruntime.Field{
	{Name: "LabelTemplate", Flag: "label-template", Type: "*string", Required: true},
	{Name: "LinkConfiguration", Flag: "link-configuration", Type: "*types.LinkConfiguration", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceType", Required: true},
	{Name: "SinkIdentifier", Flag: "sink-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sink = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_link = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_sink = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_link = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IncludeTags", Flag: "include-tags", Type: "*bool", Required: false},
}

var fields_get_sink = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IncludeTags", Flag: "include-tags", Type: "*bool", Required: false},
}

var fields_get_sink_policy = []leanruntime.Field{
	{Name: "SinkIdentifier", Flag: "sink-identifier", Type: "*string", Required: true},
}

var fields_list_attached_links = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SinkIdentifier", Flag: "sink-identifier", Type: "*string", Required: true},
}

var fields_list_links = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sinks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_sink_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "SinkIdentifier", Flag: "sink-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_link = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IncludeTags", Flag: "include-tags", Type: "*bool", Required: false},
	{Name: "LinkConfiguration", Flag: "link-configuration", Type: "*types.LinkConfiguration", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceType", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-link": {
			Name:   "create-link",
			Fields: fields_create_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLink(ctx, input)
			},
		},
		"create-sink": {
			Name:   "create-sink",
			Fields: fields_create_sink,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sink, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSink(ctx, input)
			},
		},
		"delete-link": {
			Name:   "delete-link",
			Fields: fields_delete_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLink(ctx, input)
			},
		},
		"delete-sink": {
			Name:   "delete-sink",
			Fields: fields_delete_sink,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sink, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSink(ctx, input)
			},
		},
		"get-link": {
			Name:   "get-link",
			Fields: fields_get_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLink(ctx, input)
			},
		},
		"get-sink": {
			Name:   "get-sink",
			Fields: fields_get_sink,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sink, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSink(ctx, input)
			},
		},
		"get-sink-policy": {
			Name:   "get-sink-policy",
			Fields: fields_get_sink_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSinkPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sink_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSinkPolicy(ctx, input)
			},
		},
		"list-attached-links": {
			Name:   "list-attached-links",
			Fields: fields_list_attached_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachedLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attached_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachedLinks(ctx, input)
				}
				var results []*svc.ListAttachedLinksOutput
				p := svc.NewListAttachedLinksPaginator(client, input)
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
		"list-links": {
			Name:   "list-links",
			Fields: fields_list_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLinks(ctx, input)
				}
				var results []*svc.ListLinksOutput
				p := svc.NewListLinksPaginator(client, input)
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
		"list-sinks": {
			Name:   "list-sinks",
			Fields: fields_list_sinks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sinks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSinks(ctx, input)
				}
				var results []*svc.ListSinksOutput
				p := svc.NewListSinksPaginator(client, input)
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
		"put-sink-policy": {
			Name:   "put-sink-policy",
			Fields: fields_put_sink_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSinkPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_sink_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSinkPolicy(ctx, input)
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
		"update-link": {
			Name:   "update-link",
			Fields: fields_update_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLink(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("oam", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
