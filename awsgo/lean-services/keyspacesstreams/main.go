package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/keyspacesstreams"
)

var fields_get_records = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ShardIterator", Flag: "shard-iterator", Type: "*string", Required: true},
}

var fields_get_shard_iterator = []leanruntime.Field{
	{Name: "SequenceNumber", Flag: "sequence-number", Type: "*string", Required: false},
	{Name: "ShardId", Flag: "shard-id", Type: "*string", Required: true},
	{Name: "ShardIteratorType", Flag: "shard-iterator-type", Type: "types.ShardIteratorType", Required: true},
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
}

var fields_get_stream = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShardFilter", Flag: "shard-filter", Type: "*types.ShardFilter", Required: false},
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
}

var fields_list_streams = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-records": {
			Name:   "get-records",
			Fields: fields_get_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecords(ctx, input)
			},
		},
		"get-shard-iterator": {
			Name:   "get-shard-iterator",
			Fields: fields_get_shard_iterator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetShardIteratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_shard_iterator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetShardIterator(ctx, input)
			},
		},
		"get-stream": {
			Name:   "get-stream",
			Fields: fields_get_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_stream, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetStream(ctx, input)
				}
				var results []*svc.GetStreamOutput
				p := svc.NewGetStreamPaginator(client, input)
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
		"list-streams": {
			Name:   "list-streams",
			Fields: fields_list_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreams(ctx, input)
				}
				var results []*svc.ListStreamsOutput
				p := svc.NewListStreamsPaginator(client, input)
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
	if err := leanruntime.Execute("keyspacesstreams", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
