package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

var fields_describe_stream = []leanruntime.Field{
	{Name: "ExclusiveStartShardId", Flag: "exclusive-start-shard-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "ShardFilter", Flag: "shard-filter", Type: "*types.ShardFilter", Required: false},
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
}

var fields_get_records = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "ShardIterator", Flag: "shard-iterator", Type: "*string", Required: true},
}

var fields_get_shard_iterator = []leanruntime.Field{
	{Name: "SequenceNumber", Flag: "sequence-number", Type: "*string", Required: false},
	{Name: "ShardId", Flag: "shard-id", Type: "*string", Required: true},
	{Name: "ShardIteratorType", Flag: "shard-iterator-type", Type: "types.ShardIteratorType", Required: true},
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
}

var fields_list_streams = []leanruntime.Field{
	{Name: "ExclusiveStartStreamArn", Flag: "exclusive-start-stream-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-stream": {
			Name:   "describe-stream",
			Fields: fields_describe_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStream(ctx, input)
			},
		},
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
		"list-streams": {
			Name:   "list-streams",
			Fields: fields_list_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_streams, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStreams(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("dynamodbstreams", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
