package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ebs"
)

var fields_complete_snapshot = []leanruntime.Field{
	{Name: "ChangedBlocksCount", Flag: "changed-blocks-count", Type: "*int32", Required: true},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "ChecksumAggregationMethod", Flag: "checksum-aggregation-method", Type: "types.ChecksumAggregationMethod", Required: false},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_get_snapshot_block = []leanruntime.Field{
	{Name: "BlockIndex", Flag: "block-index", Type: "*int32", Required: true},
	{Name: "BlockToken", Flag: "block-token", Type: "*string", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_list_changed_blocks = []leanruntime.Field{
	{Name: "FirstSnapshotId", Flag: "first-snapshot-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecondSnapshotId", Flag: "second-snapshot-id", Type: "*string", Required: true},
	{Name: "StartingBlockIndex", Flag: "starting-block-index", Type: "*int32", Required: false},
}

var fields_list_snapshot_blocks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
	{Name: "StartingBlockIndex", Flag: "starting-block-index", Type: "*int32", Required: false},
}

var fields_put_snapshot_block = []leanruntime.Field{
	{Name: "BlockData", Flag: "block-data", Type: "io.Reader", Required: true},
	{Name: "BlockIndex", Flag: "block-index", Type: "*int32", Required: true},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: true},
	{Name: "DataLength", Flag: "data-length", Type: "*int32", Required: true},
	{Name: "Progress", Flag: "progress", Type: "*int32", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_start_snapshot = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "ParentSnapshotId", Flag: "parent-snapshot-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "VolumeSize", Flag: "volume-size", Type: "*int64", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"complete-snapshot": {
			Name:   "complete-snapshot",
			Fields: fields_complete_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteSnapshot(ctx, input)
			},
		},
		"get-snapshot-block": {
			Name:   "get-snapshot-block",
			Fields: fields_get_snapshot_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSnapshotBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_snapshot_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSnapshotBlock(ctx, input)
			},
		},
		"list-changed-blocks": {
			Name:   "list-changed-blocks",
			Fields: fields_list_changed_blocks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChangedBlocksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_changed_blocks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChangedBlocks(ctx, input)
				}
				var results []*svc.ListChangedBlocksOutput
				p := svc.NewListChangedBlocksPaginator(client, input)
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
		"list-snapshot-blocks": {
			Name:   "list-snapshot-blocks",
			Fields: fields_list_snapshot_blocks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSnapshotBlocksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_snapshot_blocks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSnapshotBlocks(ctx, input)
				}
				var results []*svc.ListSnapshotBlocksOutput
				p := svc.NewListSnapshotBlocksPaginator(client, input)
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
		"put-snapshot-block": {
			Name:   "put-snapshot-block",
			Fields: fields_put_snapshot_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSnapshotBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_snapshot_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSnapshotBlock(ctx, input)
			},
		},
		"start-snapshot": {
			Name:   "start-snapshot",
			Fields: fields_start_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSnapshot(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ebs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
