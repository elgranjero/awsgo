package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/firehose/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-delivery-stream", "delete-delivery-stream", "describe-delivery-stream", "list-delivery-streams", "list-tags-for-delivery-stream", "put-record", "put-record-batch", "start-delivery-stream-encryption", "stop-delivery-stream-encryption", "tag-delivery-stream", "untag-delivery-stream", "update-destination"},
		OperationSet: map[string]bool{"create-delivery-stream": true, "delete-delivery-stream": true, "describe-delivery-stream": true, "list-delivery-streams": true, "list-tags-for-delivery-stream": true, "put-record": true, "put-record-batch": true, "start-delivery-stream-encryption": true, "stop-delivery-stream-encryption": true, "tag-delivery-stream": true, "untag-delivery-stream": true, "update-destination": true},
		OperationInputs: map[string][]string{
			"create-delivery-stream":           {"AmazonOpenSearchServerlessDestinationConfiguration", "AmazonopensearchserviceDestinationConfiguration", "DatabaseSourceConfiguration", "DeliveryStreamEncryptionConfigurationInput", "DeliveryStreamName", "DeliveryStreamType", "DirectPutSourceConfiguration", "ElasticsearchDestinationConfiguration", "ExtendedS3DestinationConfiguration", "HttpEndpointDestinationConfiguration", "IcebergDestinationConfiguration", "KinesisStreamSourceConfiguration", "MSKSourceConfiguration", "RedshiftDestinationConfiguration", "S3DestinationConfiguration", "SnowflakeDestinationConfiguration", "SplunkDestinationConfiguration", "Tags"},
			"delete-delivery-stream":           {"AllowForceDelete", "DeliveryStreamName"},
			"describe-delivery-stream":         {"DeliveryStreamName", "ExclusiveStartDestinationId", "Limit"},
			"list-delivery-streams":            {"DeliveryStreamType", "ExclusiveStartDeliveryStreamName", "Limit"},
			"list-tags-for-delivery-stream":    {"DeliveryStreamName", "ExclusiveStartTagKey", "Limit"},
			"put-record":                       {"DeliveryStreamName", "Record"},
			"put-record-batch":                 {"DeliveryStreamName", "Records"},
			"start-delivery-stream-encryption": {"DeliveryStreamEncryptionConfigurationInput", "DeliveryStreamName"},
			"stop-delivery-stream-encryption":  {"DeliveryStreamName"},
			"tag-delivery-stream":              {"DeliveryStreamName", "Tags"},
			"untag-delivery-stream":            {"DeliveryStreamName", "TagKeys"},
			"update-destination":               {"AmazonOpenSearchServerlessDestinationUpdate", "AmazonopensearchserviceDestinationUpdate", "CurrentDeliveryStreamVersionId", "DeliveryStreamName", "DestinationId", "ElasticsearchDestinationUpdate", "ExtendedS3DestinationUpdate", "HttpEndpointDestinationUpdate", "IcebergDestinationUpdate", "RedshiftDestinationUpdate", "S3DestinationUpdate", "SnowflakeDestinationUpdate", "SplunkDestinationUpdate"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-delivery-stream":           {"AmazonOpenSearchServerlessDestinationConfiguration": "*types.AmazonOpenSearchServerlessDestinationConfiguration", "AmazonopensearchserviceDestinationConfiguration": "*types.AmazonopensearchserviceDestinationConfiguration", "DatabaseSourceConfiguration": "*types.DatabaseSourceConfiguration", "DeliveryStreamEncryptionConfigurationInput": "*types.DeliveryStreamEncryptionConfigurationInput", "DeliveryStreamName": "*string", "DeliveryStreamType": "types.DeliveryStreamType", "DirectPutSourceConfiguration": "*types.DirectPutSourceConfiguration", "ElasticsearchDestinationConfiguration": "*types.ElasticsearchDestinationConfiguration", "ExtendedS3DestinationConfiguration": "*types.ExtendedS3DestinationConfiguration", "HttpEndpointDestinationConfiguration": "*types.HttpEndpointDestinationConfiguration", "IcebergDestinationConfiguration": "*types.IcebergDestinationConfiguration", "KinesisStreamSourceConfiguration": "*types.KinesisStreamSourceConfiguration", "MSKSourceConfiguration": "*types.MSKSourceConfiguration", "RedshiftDestinationConfiguration": "*types.RedshiftDestinationConfiguration", "S3DestinationConfiguration": "*types.S3DestinationConfiguration", "SnowflakeDestinationConfiguration": "*types.SnowflakeDestinationConfiguration", "SplunkDestinationConfiguration": "*types.SplunkDestinationConfiguration", "Tags": "[]types.Tag"},
			"delete-delivery-stream":           {"AllowForceDelete": "*bool", "DeliveryStreamName": "*string"},
			"describe-delivery-stream":         {"DeliveryStreamName": "*string", "ExclusiveStartDestinationId": "*string", "Limit": "*int32"},
			"list-delivery-streams":            {"DeliveryStreamType": "types.DeliveryStreamType", "ExclusiveStartDeliveryStreamName": "*string", "Limit": "*int32"},
			"list-tags-for-delivery-stream":    {"DeliveryStreamName": "*string", "ExclusiveStartTagKey": "*string", "Limit": "*int32"},
			"put-record":                       {"DeliveryStreamName": "*string", "Record": "*types.Record"},
			"put-record-batch":                 {"DeliveryStreamName": "*string", "Records": "[]types.Record"},
			"start-delivery-stream-encryption": {"DeliveryStreamEncryptionConfigurationInput": "*types.DeliveryStreamEncryptionConfigurationInput", "DeliveryStreamName": "*string"},
			"stop-delivery-stream-encryption":  {"DeliveryStreamName": "*string"},
			"tag-delivery-stream":              {"DeliveryStreamName": "*string", "Tags": "[]types.Tag"},
			"untag-delivery-stream":            {"DeliveryStreamName": "*string", "TagKeys": "[]string"},
			"update-destination":               {"AmazonOpenSearchServerlessDestinationUpdate": "*types.AmazonOpenSearchServerlessDestinationUpdate", "AmazonopensearchserviceDestinationUpdate": "*types.AmazonopensearchserviceDestinationUpdate", "CurrentDeliveryStreamVersionId": "*string", "DeliveryStreamName": "*string", "DestinationId": "*string", "ElasticsearchDestinationUpdate": "*types.ElasticsearchDestinationUpdate", "ExtendedS3DestinationUpdate": "*types.ExtendedS3DestinationUpdate", "HttpEndpointDestinationUpdate": "*types.HttpEndpointDestinationUpdate", "IcebergDestinationUpdate": "*types.IcebergDestinationUpdate", "RedshiftDestinationUpdate": "*types.RedshiftDestinationUpdate", "S3DestinationUpdate": "*types.S3DestinationUpdate", "SnowflakeDestinationUpdate": "*types.SnowflakeDestinationUpdate", "SplunkDestinationUpdate": "*types.SplunkDestinationUpdate"},
		},
		OperationInputRequired: map[string][]string{
			"create-delivery-stream":           {"DeliveryStreamName"},
			"delete-delivery-stream":           {"DeliveryStreamName"},
			"describe-delivery-stream":         {"DeliveryStreamName"},
			"list-delivery-streams":            {},
			"list-tags-for-delivery-stream":    {"DeliveryStreamName"},
			"put-record":                       {"DeliveryStreamName", "Record"},
			"put-record-batch":                 {"DeliveryStreamName", "Records"},
			"start-delivery-stream-encryption": {"DeliveryStreamName"},
			"stop-delivery-stream-encryption":  {"DeliveryStreamName"},
			"tag-delivery-stream":              {"DeliveryStreamName", "Tags"},
			"untag-delivery-stream":            {"DeliveryStreamName", "TagKeys"},
			"update-destination":               {"CurrentDeliveryStreamVersionId", "DeliveryStreamName", "DestinationId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("firehose", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
