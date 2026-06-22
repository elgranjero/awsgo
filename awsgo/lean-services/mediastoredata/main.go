package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediastoredata"
)

var fields_delete_object = []leanruntime.Field{
	{Name: "Path", Flag: "path", Type: "*string", Required: true},
}

var fields_describe_object = []leanruntime.Field{
	{Name: "Path", Flag: "path", Type: "*string", Required: true},
}

var fields_get_object = []leanruntime.Field{
	{Name: "Path", Flag: "path", Type: "*string", Required: true},
	{Name: "Range", Flag: "range", Type: "*string", Required: false},
}

var fields_list_items = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
}

var fields_put_object = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "io.Reader", Required: true},
	{Name: "CacheControl", Flag: "cache-control", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: true},
	{Name: "StorageClass", Flag: "storage-class", Type: "types.StorageClass", Required: false},
	{Name: "UploadAvailability", Flag: "upload-availability", Type: "types.UploadAvailability", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-object": {
			Name:   "delete-object",
			Fields: fields_delete_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObject(ctx, input)
			},
		},
		"describe-object": {
			Name:   "describe-object",
			Fields: fields_describe_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeObject(ctx, input)
			},
		},
		"get-object": {
			Name:   "get-object",
			Fields: fields_get_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObject(ctx, input)
			},
		},
		"list-items": {
			Name:   "list-items",
			Fields: fields_list_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListItems(ctx, input)
				}
				var results []*svc.ListItemsOutput
				p := svc.NewListItemsPaginator(client, input)
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
		"put-object": {
			Name:   "put-object",
			Fields: fields_put_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObject(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mediastoredata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
