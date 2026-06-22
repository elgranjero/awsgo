package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cognitosync"
)

var fields_bulk_publish = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_describe_identity_pool_usage = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_describe_identity_usage = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_get_bulk_publish_details = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_get_cognito_events = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_get_identity_pool_configuration = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_pool_usage = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_records = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "LastSyncCount", Flag: "last-sync-count", Type: "*int64", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SyncSessionToken", Flag: "sync-session-token", Type: "*string", Required: false},
}

var fields_register_device = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
}

var fields_set_cognito_events = []leanruntime.Field{
	{Name: "Events", Flag: "events", Type: "map[string]string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_set_identity_pool_configuration = []leanruntime.Field{
	{Name: "CognitoStreams", Flag: "cognito-streams", Type: "*types.CognitoStreams", Required: false},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "PushSync", Flag: "push-sync", Type: "*types.PushSync", Required: false},
}

var fields_subscribe_to_dataset = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_unsubscribe_from_dataset = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_update_records = []leanruntime.Field{
	{Name: "ClientContext", Flag: "client-context", Type: "*string", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: false},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "RecordPatches", Flag: "record-patches", Type: "[]types.RecordPatch", Required: false},
	{Name: "SyncSessionToken", Flag: "sync-session-token", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"bulk-publish": {
			Name:   "bulk-publish",
			Fields: fields_bulk_publish,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BulkPublishInput{}
				if _, err := leanruntime.ApplyInput(input, fields_bulk_publish, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BulkPublish(ctx, input)
			},
		},
		"delete-dataset": {
			Name:   "delete-dataset",
			Fields: fields_delete_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataset(ctx, input)
			},
		},
		"describe-dataset": {
			Name:   "describe-dataset",
			Fields: fields_describe_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataset(ctx, input)
			},
		},
		"describe-identity-pool-usage": {
			Name:   "describe-identity-pool-usage",
			Fields: fields_describe_identity_pool_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityPoolUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_pool_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityPoolUsage(ctx, input)
			},
		},
		"describe-identity-usage": {
			Name:   "describe-identity-usage",
			Fields: fields_describe_identity_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityUsage(ctx, input)
			},
		},
		"get-bulk-publish-details": {
			Name:   "get-bulk-publish-details",
			Fields: fields_get_bulk_publish_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBulkPublishDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bulk_publish_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBulkPublishDetails(ctx, input)
			},
		},
		"get-cognito-events": {
			Name:   "get-cognito-events",
			Fields: fields_get_cognito_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCognitoEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cognito_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCognitoEvents(ctx, input)
			},
		},
		"get-identity-pool-configuration": {
			Name:   "get-identity-pool-configuration",
			Fields: fields_get_identity_pool_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityPoolConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_pool_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityPoolConfiguration(ctx, input)
			},
		},
		"list-datasets": {
			Name:   "list-datasets",
			Fields: fields_list_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_datasets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDatasets(ctx, input)
			},
		},
		"list-identity-pool-usage": {
			Name:   "list-identity-pool-usage",
			Fields: fields_list_identity_pool_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityPoolUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_identity_pool_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIdentityPoolUsage(ctx, input)
			},
		},
		"list-records": {
			Name:   "list-records",
			Fields: fields_list_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRecords(ctx, input)
			},
		},
		"register-device": {
			Name:   "register-device",
			Fields: fields_register_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDevice(ctx, input)
			},
		},
		"set-cognito-events": {
			Name:   "set-cognito-events",
			Fields: fields_set_cognito_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetCognitoEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_cognito_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetCognitoEvents(ctx, input)
			},
		},
		"set-identity-pool-configuration": {
			Name:   "set-identity-pool-configuration",
			Fields: fields_set_identity_pool_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityPoolConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_pool_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityPoolConfiguration(ctx, input)
			},
		},
		"subscribe-to-dataset": {
			Name:   "subscribe-to-dataset",
			Fields: fields_subscribe_to_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubscribeToDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_subscribe_to_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubscribeToDataset(ctx, input)
			},
		},
		"unsubscribe-from-dataset": {
			Name:   "unsubscribe-from-dataset",
			Fields: fields_unsubscribe_from_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnsubscribeFromDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unsubscribe_from_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnsubscribeFromDataset(ctx, input)
			},
		},
		"update-records": {
			Name:   "update-records",
			Fields: fields_update_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecords(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cognitosync", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
