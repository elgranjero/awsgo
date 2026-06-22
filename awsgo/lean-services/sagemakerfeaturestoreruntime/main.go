package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime"
)

var fields_batch_get_record = []leanruntime.Field{
	{Name: "ExpirationTimeResponse", Flag: "expiration-time-response", Type: "types.ExpirationTimeResponse", Required: false},
	{Name: "Identifiers", Flag: "identifiers", Type: "[]types.BatchGetRecordIdentifier", Required: true},
}

var fields_delete_record = []leanruntime.Field{
	{Name: "DeletionMode", Flag: "deletion-mode", Type: "types.DeletionMode", Required: false},
	{Name: "EventTime", Flag: "event-time", Type: "*string", Required: true},
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "RecordIdentifierValueAsString", Flag: "record-identifier-value-as-string", Type: "*string", Required: true},
	{Name: "TargetStores", Flag: "target-stores", Type: "[]types.TargetStore", Required: false},
}

var fields_get_record = []leanruntime.Field{
	{Name: "ExpirationTimeResponse", Flag: "expiration-time-response", Type: "types.ExpirationTimeResponse", Required: false},
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "FeatureNames", Flag: "feature-names", Type: "[]string", Required: false},
	{Name: "RecordIdentifierValueAsString", Flag: "record-identifier-value-as-string", Type: "*string", Required: true},
}

var fields_put_record = []leanruntime.Field{
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "Record", Flag: "record", Type: "[]types.FeatureValue", Required: true},
	{Name: "TargetStores", Flag: "target-stores", Type: "[]types.TargetStore", Required: false},
	{Name: "TtlDuration", Flag: "ttl-duration", Type: "*types.TtlDuration", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-record": {
			Name:   "batch-get-record",
			Fields: fields_batch_get_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetRecord(ctx, input)
			},
		},
		"delete-record": {
			Name:   "delete-record",
			Fields: fields_delete_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecord(ctx, input)
			},
		},
		"get-record": {
			Name:   "get-record",
			Fields: fields_get_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecord(ctx, input)
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
	}
	if err := leanruntime.Execute("sagemakerfeaturestoreruntime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
