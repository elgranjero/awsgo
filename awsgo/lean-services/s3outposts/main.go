package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/s3outposts"
)

var fields_create_endpoint = []leanruntime.Field{
	{Name: "AccessType", Flag: "access-type", Type: "types.EndpointAccessType", Required: false},
	{Name: "CustomerOwnedIpv4Pool", Flag: "customer-owned-ipv4-pool", Type: "*string", Required: false},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
	{Name: "SecurityGroupId", Flag: "security-group-id", Type: "*string", Required: true},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
}

var fields_list_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_outposts_with_s3 = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_shared_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-endpoint": {
			Name:   "create-endpoint",
			Fields: fields_create_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpoint(ctx, input)
			},
		},
		"delete-endpoint": {
			Name:   "delete-endpoint",
			Fields: fields_delete_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpoint(ctx, input)
			},
		},
		"list-endpoints": {
			Name:   "list-endpoints",
			Fields: fields_list_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpoints(ctx, input)
				}
				var results []*svc.ListEndpointsOutput
				p := svc.NewListEndpointsPaginator(client, input)
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
		"list-outposts-with-s3": {
			Name:   "list-outposts-with-s3",
			Fields: fields_list_outposts_with_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOutpostsWithS3Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_outposts_with_s3, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOutpostsWithS3(ctx, input)
				}
				var results []*svc.ListOutpostsWithS3Output
				p := svc.NewListOutpostsWithS3Paginator(client, input)
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
		"list-shared-endpoints": {
			Name:   "list-shared-endpoints",
			Fields: fields_list_shared_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSharedEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_shared_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSharedEndpoints(ctx, input)
				}
				var results []*svc.ListSharedEndpointsOutput
				p := svc.NewListSharedEndpointsPaginator(client, input)
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
	}
	if err := leanruntime.Execute("s3outposts", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
