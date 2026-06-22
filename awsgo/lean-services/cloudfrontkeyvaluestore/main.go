package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore"
)

var fields_delete_key = []leanruntime.Field{
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "KvsARN", Flag: "kvs-arn", Type: "*string", Required: true},
}

var fields_describe_key_value_store = []leanruntime.Field{
	{Name: "KvsARN", Flag: "kvs-arn", Type: "*string", Required: true},
}

var fields_get_key = []leanruntime.Field{
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "KvsARN", Flag: "kvs-arn", Type: "*string", Required: true},
}

var fields_list_keys = []leanruntime.Field{
	{Name: "KvsARN", Flag: "kvs-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_key = []leanruntime.Field{
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "KvsARN", Flag: "kvs-arn", Type: "*string", Required: true},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_update_keys = []leanruntime.Field{
	{Name: "Deletes", Flag: "deletes", Type: "[]types.DeleteKeyRequestListItem", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: true},
	{Name: "KvsARN", Flag: "kvs-arn", Type: "*string", Required: true},
	{Name: "Puts", Flag: "puts", Type: "[]types.PutKeyRequestListItem", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-key": {
			Name:   "delete-key",
			Fields: fields_delete_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKey(ctx, input)
			},
		},
		"describe-key-value-store": {
			Name:   "describe-key-value-store",
			Fields: fields_describe_key_value_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyValueStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key_value_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKeyValueStore(ctx, input)
			},
		},
		"get-key": {
			Name:   "get-key",
			Fields: fields_get_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKey(ctx, input)
			},
		},
		"list-keys": {
			Name:   "list-keys",
			Fields: fields_list_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeys(ctx, input)
				}
				var results []*svc.ListKeysOutput
				p := svc.NewListKeysPaginator(client, input)
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
		"put-key": {
			Name:   "put-key",
			Fields: fields_put_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutKey(ctx, input)
			},
		},
		"update-keys": {
			Name:   "update-keys",
			Fields: fields_update_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKeys(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudfrontkeyvaluestore", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
