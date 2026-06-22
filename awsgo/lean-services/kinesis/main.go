package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesis"
)

var fields_add_tags_to_stream = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_create_stream = []leanruntime.Field{
	{Name: "MaxRecordSizeInKiB", Flag: "max-record-size-in-ki-b", Type: "*int32", Required: false},
	{Name: "ShardCount", Flag: "shard-count", Type: "*int32", Required: false},
	{Name: "StreamModeDetails", Flag: "stream-mode-details", Type: "*types.StreamModeDetails", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WarmThroughputMiBps", Flag: "warm-throughput-mi-bps", Type: "*int32", Required: false},
}

var fields_decrease_stream_retention_period = []leanruntime.Field{
	{Name: "RetentionPeriodHours", Flag: "retention-period-hours", Type: "*int32", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_delete_stream = []leanruntime.Field{
	{Name: "EnforceConsumerDeletion", Flag: "enforce-consumer-deletion", Type: "*bool", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_deregister_stream_consumer = []leanruntime.Field{
	{Name: "ConsumerARN", Flag: "consumer-arn", Type: "*string", Required: false},
	{Name: "ConsumerName", Flag: "consumer-name", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_describe_account_settings = []leanruntime.Field{}

var fields_describe_limits = []leanruntime.Field{}

var fields_describe_stream = []leanruntime.Field{
	{Name: "ExclusiveStartShardId", Flag: "exclusive-start-shard-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_describe_stream_consumer = []leanruntime.Field{
	{Name: "ConsumerARN", Flag: "consumer-arn", Type: "*string", Required: false},
	{Name: "ConsumerName", Flag: "consumer-name", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_describe_stream_summary = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_disable_enhanced_monitoring = []leanruntime.Field{
	{Name: "ShardLevelMetrics", Flag: "shard-level-metrics", Type: "[]types.MetricsName", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_enable_enhanced_monitoring = []leanruntime.Field{
	{Name: "ShardLevelMetrics", Flag: "shard-level-metrics", Type: "[]types.MetricsName", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_get_records = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "ShardIterator", Flag: "shard-iterator", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_get_shard_iterator = []leanruntime.Field{
	{Name: "ShardId", Flag: "shard-id", Type: "*string", Required: true},
	{Name: "ShardIteratorType", Flag: "shard-iterator-type", Type: "types.ShardIteratorType", Required: true},
	{Name: "StartingSequenceNumber", Flag: "starting-sequence-number", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "Timestamp", Flag: "timestamp", Type: "*time.Time", Required: false},
}

var fields_increase_stream_retention_period = []leanruntime.Field{
	{Name: "RetentionPeriodHours", Flag: "retention-period-hours", Type: "*int32", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_list_shards = []leanruntime.Field{
	{Name: "ExclusiveStartShardId", Flag: "exclusive-start-shard-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShardFilter", Flag: "shard-filter", Type: "*types.ShardFilter", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamCreationTimestamp", Flag: "stream-creation-timestamp", Type: "*time.Time", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_list_stream_consumers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: true},
	{Name: "StreamCreationTimestamp", Flag: "stream-creation-timestamp", Type: "*time.Time", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_list_streams = []leanruntime.Field{
	{Name: "ExclusiveStartStreamName", Flag: "exclusive-start-stream-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_list_tags_for_stream = []leanruntime.Field{
	{Name: "ExclusiveStartTagKey", Flag: "exclusive-start-tag-key", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_merge_shards = []leanruntime.Field{
	{Name: "AdjacentShardToMerge", Flag: "adjacent-shard-to-merge", Type: "*string", Required: true},
	{Name: "ShardToMerge", Flag: "shard-to-merge", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_put_record = []leanruntime.Field{
	{Name: "Data", Flag: "data", Type: "[]byte", Required: true},
	{Name: "ExplicitHashKey", Flag: "explicit-hash-key", Type: "*string", Required: false},
	{Name: "PartitionKey", Flag: "partition-key", Type: "*string", Required: true},
	{Name: "SequenceNumberForOrdering", Flag: "sequence-number-for-ordering", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_put_records = []leanruntime.Field{
	{Name: "Records", Flag: "records", Type: "[]types.PutRecordsRequestEntry", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_register_stream_consumer = []leanruntime.Field{
	{Name: "ConsumerName", Flag: "consumer-name", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_remove_tags_from_stream = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_split_shard = []leanruntime.Field{
	{Name: "NewStartingHashKey", Flag: "new-starting-hash-key", Type: "*string", Required: true},
	{Name: "ShardToSplit", Flag: "shard-to-split", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_start_stream_encryption = []leanruntime.Field{
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_stop_stream_encryption = []leanruntime.Field{
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_subscribe_to_shard = []leanruntime.Field{
	{Name: "ConsumerARN", Flag: "consumer-arn", Type: "*string", Required: true},
	{Name: "ShardId", Flag: "shard-id", Type: "*string", Required: true},
	{Name: "StartingPosition", Flag: "starting-position", Type: "*types.StartingPosition", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "MinimumThroughputBillingCommitment", Flag: "minimum-throughput-billing-commitment", Type: "*types.MinimumThroughputBillingCommitmentInput", Required: true},
}

var fields_update_max_record_size = []leanruntime.Field{
	{Name: "MaxRecordSizeInKiB", Flag: "max-record-size-in-ki-b", Type: "*int32", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_update_shard_count = []leanruntime.Field{
	{Name: "ScalingType", Flag: "scaling-type", Type: "types.ScalingType", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "TargetShardCount", Flag: "target-shard-count", Type: "*int32", Required: true},
}

var fields_update_stream_mode = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamModeDetails", Flag: "stream-mode-details", Type: "*types.StreamModeDetails", Required: true},
	{Name: "WarmThroughputMiBps", Flag: "warm-throughput-mi-bps", Type: "*int32", Required: false},
}

var fields_update_stream_warm_throughput = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "WarmThroughputMiBps", Flag: "warm-throughput-mi-bps", Type: "*int32", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags-to-stream": {
			Name:   "add-tags-to-stream",
			Fields: fields_add_tags_to_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToStream(ctx, input)
			},
		},
		"create-stream": {
			Name:   "create-stream",
			Fields: fields_create_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStream(ctx, input)
			},
		},
		"decrease-stream-retention-period": {
			Name:   "decrease-stream-retention-period",
			Fields: fields_decrease_stream_retention_period,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecreaseStreamRetentionPeriodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decrease_stream_retention_period, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DecreaseStreamRetentionPeriod(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-stream": {
			Name:   "delete-stream",
			Fields: fields_delete_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStream(ctx, input)
			},
		},
		"deregister-stream-consumer": {
			Name:   "deregister-stream-consumer",
			Fields: fields_deregister_stream_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterStreamConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_stream_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterStreamConsumer(ctx, input)
			},
		},
		"describe-account-settings": {
			Name:   "describe-account-settings",
			Fields: fields_describe_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountSettings(ctx, input)
			},
		},
		"describe-limits": {
			Name:   "describe-limits",
			Fields: fields_describe_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLimits(ctx, input)
			},
		},
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
		"describe-stream-consumer": {
			Name:   "describe-stream-consumer",
			Fields: fields_describe_stream_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStreamConsumer(ctx, input)
			},
		},
		"describe-stream-summary": {
			Name:   "describe-stream-summary",
			Fields: fields_describe_stream_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStreamSummary(ctx, input)
			},
		},
		"disable-enhanced-monitoring": {
			Name:   "disable-enhanced-monitoring",
			Fields: fields_disable_enhanced_monitoring,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableEnhancedMonitoringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_enhanced_monitoring, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableEnhancedMonitoring(ctx, input)
			},
		},
		"enable-enhanced-monitoring": {
			Name:   "enable-enhanced-monitoring",
			Fields: fields_enable_enhanced_monitoring,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableEnhancedMonitoringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_enhanced_monitoring, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableEnhancedMonitoring(ctx, input)
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
		"increase-stream-retention-period": {
			Name:   "increase-stream-retention-period",
			Fields: fields_increase_stream_retention_period,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IncreaseStreamRetentionPeriodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_increase_stream_retention_period, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IncreaseStreamRetentionPeriod(ctx, input)
			},
		},
		"list-shards": {
			Name:   "list-shards",
			Fields: fields_list_shards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListShardsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_shards, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListShards(ctx, input)
			},
		},
		"list-stream-consumers": {
			Name:   "list-stream-consumers",
			Fields: fields_list_stream_consumers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamConsumersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stream_consumers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamConsumers(ctx, input)
				}
				var results []*svc.ListStreamConsumersOutput
				p := svc.NewListStreamConsumersPaginator(client, input)
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
		"list-tags-for-stream": {
			Name:   "list-tags-for-stream",
			Fields: fields_list_tags_for_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForStream(ctx, input)
			},
		},
		"merge-shards": {
			Name:   "merge-shards",
			Fields: fields_merge_shards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergeShardsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_shards, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergeShards(ctx, input)
			},
		},
		"put-record": {
			Name:   "put-record",
			Fields: fields_put_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRecord(ctx, input)
			},
		},
		"put-records": {
			Name:   "put-records",
			Fields: fields_put_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRecords(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"register-stream-consumer": {
			Name:   "register-stream-consumer",
			Fields: fields_register_stream_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterStreamConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_stream_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterStreamConsumer(ctx, input)
			},
		},
		"remove-tags-from-stream": {
			Name:   "remove-tags-from-stream",
			Fields: fields_remove_tags_from_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromStream(ctx, input)
			},
		},
		"split-shard": {
			Name:   "split-shard",
			Fields: fields_split_shard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SplitShardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_split_shard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SplitShard(ctx, input)
			},
		},
		"start-stream-encryption": {
			Name:   "start-stream-encryption",
			Fields: fields_start_stream_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartStreamEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_stream_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartStreamEncryption(ctx, input)
			},
		},
		"stop-stream-encryption": {
			Name:   "stop-stream-encryption",
			Fields: fields_stop_stream_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopStreamEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_stream_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopStreamEncryption(ctx, input)
			},
		},
		"subscribe-to-shard": {
			Name:   "subscribe-to-shard",
			Fields: fields_subscribe_to_shard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubscribeToShardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_subscribe_to_shard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubscribeToShard(ctx, input)
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
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-max-record-size": {
			Name:   "update-max-record-size",
			Fields: fields_update_max_record_size,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMaxRecordSizeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_max_record_size, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMaxRecordSize(ctx, input)
			},
		},
		"update-shard-count": {
			Name:   "update-shard-count",
			Fields: fields_update_shard_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateShardCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_shard_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateShardCount(ctx, input)
			},
		},
		"update-stream-mode": {
			Name:   "update-stream-mode",
			Fields: fields_update_stream_mode,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamModeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream_mode, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStreamMode(ctx, input)
			},
		},
		"update-stream-warm-throughput": {
			Name:   "update-stream-warm-throughput",
			Fields: fields_update_stream_warm_throughput,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamWarmThroughputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream_warm_throughput, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStreamWarmThroughput(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kinesis", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
