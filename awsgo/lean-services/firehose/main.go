package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/firehose"
)

var fields_create_delivery_stream = []leanruntime.Field{
	{Name: "AmazonOpenSearchServerlessDestinationConfiguration", Flag: "amazon-open-search-serverless-destination-configuration", Type: "*types.AmazonOpenSearchServerlessDestinationConfiguration", Required: false},
	{Name: "AmazonopensearchserviceDestinationConfiguration", Flag: "amazonopensearchservice-destination-configuration", Type: "*types.AmazonopensearchserviceDestinationConfiguration", Required: false},
	{Name: "DatabaseSourceConfiguration", Flag: "database-source-configuration", Type: "*types.DatabaseSourceConfiguration", Required: false},
	{Name: "DeliveryStreamEncryptionConfigurationInput", Flag: "delivery-stream-encryption-configuration-input", Type: "*types.DeliveryStreamEncryptionConfigurationInput", Required: false},
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "DeliveryStreamType", Flag: "delivery-stream-type", Type: "types.DeliveryStreamType", Required: false},
	{Name: "DirectPutSourceConfiguration", Flag: "direct-put-source-configuration", Type: "*types.DirectPutSourceConfiguration", Required: false},
	{Name: "ElasticsearchDestinationConfiguration", Flag: "elasticsearch-destination-configuration", Type: "*types.ElasticsearchDestinationConfiguration", Required: false},
	{Name: "ExtendedS3DestinationConfiguration", Flag: "extended-s3-destination-configuration", Type: "*types.ExtendedS3DestinationConfiguration", Required: false},
	{Name: "HttpEndpointDestinationConfiguration", Flag: "http-endpoint-destination-configuration", Type: "*types.HttpEndpointDestinationConfiguration", Required: false},
	{Name: "IcebergDestinationConfiguration", Flag: "iceberg-destination-configuration", Type: "*types.IcebergDestinationConfiguration", Required: false},
	{Name: "KinesisStreamSourceConfiguration", Flag: "kinesis-stream-source-configuration", Type: "*types.KinesisStreamSourceConfiguration", Required: false},
	{Name: "MSKSourceConfiguration", Flag: "msk-source-configuration", Type: "*types.MSKSourceConfiguration", Required: false},
	{Name: "RedshiftDestinationConfiguration", Flag: "redshift-destination-configuration", Type: "*types.RedshiftDestinationConfiguration", Required: false},
	{Name: "S3DestinationConfiguration", Flag: "s3-destination-configuration", Type: "*types.S3DestinationConfiguration", Required: false},
	{Name: "SnowflakeDestinationConfiguration", Flag: "snowflake-destination-configuration", Type: "*types.SnowflakeDestinationConfiguration", Required: false},
	{Name: "SplunkDestinationConfiguration", Flag: "splunk-destination-configuration", Type: "*types.SplunkDestinationConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_delivery_stream = []leanruntime.Field{
	{Name: "AllowForceDelete", Flag: "allow-force-delete", Type: "*bool", Required: false},
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
}

var fields_describe_delivery_stream = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "ExclusiveStartDestinationId", Flag: "exclusive-start-destination-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
}

var fields_list_delivery_streams = []leanruntime.Field{
	{Name: "DeliveryStreamType", Flag: "delivery-stream-type", Type: "types.DeliveryStreamType", Required: false},
	{Name: "ExclusiveStartDeliveryStreamName", Flag: "exclusive-start-delivery-stream-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
}

var fields_list_tags_for_delivery_stream = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "ExclusiveStartTagKey", Flag: "exclusive-start-tag-key", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
}

var fields_put_record = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "Record", Flag: "record", Type: "*types.Record", Required: true},
}

var fields_put_record_batch = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "Records", Flag: "records", Type: "[]types.Record", Required: true},
}

var fields_start_delivery_stream_encryption = []leanruntime.Field{
	{Name: "DeliveryStreamEncryptionConfigurationInput", Flag: "delivery-stream-encryption-configuration-input", Type: "*types.DeliveryStreamEncryptionConfigurationInput", Required: false},
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
}

var fields_stop_delivery_stream_encryption = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
}

var fields_tag_delivery_stream = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_delivery_stream = []leanruntime.Field{
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_destination = []leanruntime.Field{
	{Name: "AmazonOpenSearchServerlessDestinationUpdate", Flag: "amazon-open-search-serverless-destination-update", Type: "*types.AmazonOpenSearchServerlessDestinationUpdate", Required: false},
	{Name: "AmazonopensearchserviceDestinationUpdate", Flag: "amazonopensearchservice-destination-update", Type: "*types.AmazonopensearchserviceDestinationUpdate", Required: false},
	{Name: "CurrentDeliveryStreamVersionId", Flag: "current-delivery-stream-version-id", Type: "*string", Required: true},
	{Name: "DeliveryStreamName", Flag: "delivery-stream-name", Type: "*string", Required: true},
	{Name: "DestinationId", Flag: "destination-id", Type: "*string", Required: true},
	{Name: "ElasticsearchDestinationUpdate", Flag: "elasticsearch-destination-update", Type: "*types.ElasticsearchDestinationUpdate", Required: false},
	{Name: "ExtendedS3DestinationUpdate", Flag: "extended-s3-destination-update", Type: "*types.ExtendedS3DestinationUpdate", Required: false},
	{Name: "HttpEndpointDestinationUpdate", Flag: "http-endpoint-destination-update", Type: "*types.HttpEndpointDestinationUpdate", Required: false},
	{Name: "IcebergDestinationUpdate", Flag: "iceberg-destination-update", Type: "*types.IcebergDestinationUpdate", Required: false},
	{Name: "RedshiftDestinationUpdate", Flag: "redshift-destination-update", Type: "*types.RedshiftDestinationUpdate", Required: false},
	{Name: "S3DestinationUpdate", Flag: "s3-destination-update", Type: "*types.S3DestinationUpdate", Required: false},
	{Name: "SnowflakeDestinationUpdate", Flag: "snowflake-destination-update", Type: "*types.SnowflakeDestinationUpdate", Required: false},
	{Name: "SplunkDestinationUpdate", Flag: "splunk-destination-update", Type: "*types.SplunkDestinationUpdate", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-delivery-stream": {
			Name:   "create-delivery-stream",
			Fields: fields_create_delivery_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeliveryStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_delivery_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeliveryStream(ctx, input)
			},
		},
		"delete-delivery-stream": {
			Name:   "delete-delivery-stream",
			Fields: fields_delete_delivery_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeliveryStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_delivery_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeliveryStream(ctx, input)
			},
		},
		"describe-delivery-stream": {
			Name:   "describe-delivery-stream",
			Fields: fields_describe_delivery_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeliveryStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_delivery_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeliveryStream(ctx, input)
			},
		},
		"list-delivery-streams": {
			Name:   "list-delivery-streams",
			Fields: fields_list_delivery_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeliveryStreamsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_delivery_streams, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDeliveryStreams(ctx, input)
			},
		},
		"list-tags-for-delivery-stream": {
			Name:   "list-tags-for-delivery-stream",
			Fields: fields_list_tags_for_delivery_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForDeliveryStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_delivery_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForDeliveryStream(ctx, input)
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
		"put-record-batch": {
			Name:   "put-record-batch",
			Fields: fields_put_record_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRecordBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_record_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRecordBatch(ctx, input)
			},
		},
		"start-delivery-stream-encryption": {
			Name:   "start-delivery-stream-encryption",
			Fields: fields_start_delivery_stream_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeliveryStreamEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_delivery_stream_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeliveryStreamEncryption(ctx, input)
			},
		},
		"stop-delivery-stream-encryption": {
			Name:   "stop-delivery-stream-encryption",
			Fields: fields_stop_delivery_stream_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDeliveryStreamEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_delivery_stream_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDeliveryStreamEncryption(ctx, input)
			},
		},
		"tag-delivery-stream": {
			Name:   "tag-delivery-stream",
			Fields: fields_tag_delivery_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagDeliveryStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_delivery_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagDeliveryStream(ctx, input)
			},
		},
		"untag-delivery-stream": {
			Name:   "untag-delivery-stream",
			Fields: fields_untag_delivery_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagDeliveryStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_delivery_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagDeliveryStream(ctx, input)
			},
		},
		"update-destination": {
			Name:   "update-destination",
			Fields: fields_update_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDestination(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("firehose", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
