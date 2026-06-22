package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
)

var fields_close_tunnel = []leanruntime.Field{
	{Name: "Delete", Flag: "delete", Type: "*bool", Required: false},
	{Name: "TunnelId", Flag: "tunnel-id", Type: "*string", Required: true},
}

var fields_describe_tunnel = []leanruntime.Field{
	{Name: "TunnelId", Flag: "tunnel-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tunnels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: false},
}

var fields_open_tunnel = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationConfig", Flag: "destination-config", Type: "*types.DestinationConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutConfig", Flag: "timeout-config", Type: "*types.TimeoutConfig", Required: false},
}

var fields_rotate_tunnel_access_token = []leanruntime.Field{
	{Name: "ClientMode", Flag: "client-mode", Type: "types.ClientMode", Required: true},
	{Name: "DestinationConfig", Flag: "destination-config", Type: "*types.DestinationConfig", Required: false},
	{Name: "TunnelId", Flag: "tunnel-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"close-tunnel": {
			Name:   "close-tunnel",
			Fields: fields_close_tunnel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CloseTunnelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_close_tunnel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CloseTunnel(ctx, input)
			},
		},
		"describe-tunnel": {
			Name:   "describe-tunnel",
			Fields: fields_describe_tunnel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTunnelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tunnel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTunnel(ctx, input)
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
		"list-tunnels": {
			Name:   "list-tunnels",
			Fields: fields_list_tunnels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTunnelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tunnels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTunnels(ctx, input)
				}
				var results []*svc.ListTunnelsOutput
				p := svc.NewListTunnelsPaginator(client, input)
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
		"open-tunnel": {
			Name:   "open-tunnel",
			Fields: fields_open_tunnel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OpenTunnelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_open_tunnel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OpenTunnel(ctx, input)
			},
		},
		"rotate-tunnel-access-token": {
			Name:   "rotate-tunnel-access-token",
			Fields: fields_rotate_tunnel_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RotateTunnelAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rotate_tunnel_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RotateTunnelAccessToken(ctx, input)
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
	}
	if err := leanruntime.Execute("iotsecuretunneling", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
