package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
)

var fields_generate_data_set = []leanruntime.Field{
	{Name: "CustomerDefinedValues", Flag: "customer-defined-values", Type: "map[string]string", Required: false},
	{Name: "DataSetPublicationDate", Flag: "data-set-publication-date", Type: "*time.Time", Required: true},
	{Name: "DataSetType", Flag: "data-set-type", Type: "types.DataSetType", Required: true},
	{Name: "DestinationS3BucketName", Flag: "destination-s3-bucket-name", Type: "*string", Required: true},
	{Name: "DestinationS3Prefix", Flag: "destination-s3-prefix", Type: "*string", Required: false},
	{Name: "RoleNameArn", Flag: "role-name-arn", Type: "*string", Required: true},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: true},
}

var fields_start_support_data_export = []leanruntime.Field{
	{Name: "CustomerDefinedValues", Flag: "customer-defined-values", Type: "map[string]string", Required: false},
	{Name: "DataSetType", Flag: "data-set-type", Type: "types.SupportDataSetType", Required: true},
	{Name: "DestinationS3BucketName", Flag: "destination-s3-bucket-name", Type: "*string", Required: true},
	{Name: "DestinationS3Prefix", Flag: "destination-s3-prefix", Type: "*string", Required: false},
	{Name: "FromDate", Flag: "from-date", Type: "*time.Time", Required: true},
	{Name: "RoleNameArn", Flag: "role-name-arn", Type: "*string", Required: true},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"generate-data-set": {
			Name:   "generate-data-set",
			Fields: fields_generate_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateDataSet(ctx, input)
			},
		},
		"start-support-data-export": {
			Name:   "start-support-data-export",
			Fields: fields_start_support_data_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSupportDataExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_support_data_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSupportDataExport(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("marketplacecommerceanalytics", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
