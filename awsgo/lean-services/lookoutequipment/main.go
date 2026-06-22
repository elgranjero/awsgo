package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
)

var fields_create_dataset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DatasetSchema", Flag: "dataset-schema", Type: "*types.DatasetSchema", Required: false},
	{Name: "ServerSideKmsKeyId", Flag: "server-side-kms-key-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_inference_scheduler = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DataDelayOffsetInMinutes", Flag: "data-delay-offset-in-minutes", Type: "*int64", Required: false},
	{Name: "DataInputConfiguration", Flag: "data-input-configuration", Type: "*types.InferenceInputConfiguration", Required: true},
	{Name: "DataOutputConfiguration", Flag: "data-output-configuration", Type: "*types.InferenceOutputConfiguration", Required: true},
	{Name: "DataUploadFrequency", Flag: "data-upload-frequency", Type: "types.DataUploadFrequency", Required: true},
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "ServerSideKmsKeyId", Flag: "server-side-kms-key-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_label = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Equipment", Flag: "equipment", Type: "*string", Required: false},
	{Name: "FaultCode", Flag: "fault-code", Type: "*string", Required: false},
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "Rating", Flag: "rating", Type: "types.LabelRating", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_create_label_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "FaultCodes", Flag: "fault-codes", Type: "[]string", Required: false},
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DataPreProcessingConfiguration", Flag: "data-pre-processing-configuration", Type: "*types.DataPreProcessingConfiguration", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DatasetSchema", Flag: "dataset-schema", Type: "*types.DatasetSchema", Required: false},
	{Name: "EvaluationDataEndTime", Flag: "evaluation-data-end-time", Type: "*time.Time", Required: false},
	{Name: "EvaluationDataStartTime", Flag: "evaluation-data-start-time", Type: "*time.Time", Required: false},
	{Name: "LabelsInputConfiguration", Flag: "labels-input-configuration", Type: "*types.LabelsInputConfiguration", Required: false},
	{Name: "ModelDiagnosticsOutputConfiguration", Flag: "model-diagnostics-output-configuration", Type: "*types.ModelDiagnosticsOutputConfiguration", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "OffCondition", Flag: "off-condition", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ServerSideKmsKeyId", Flag: "server-side-kms-key-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingDataEndTime", Flag: "training-data-end-time", Type: "*time.Time", Required: false},
	{Name: "TrainingDataStartTime", Flag: "training-data-start-time", Type: "*time.Time", Required: false},
}

var fields_create_retraining_scheduler = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "LookbackWindow", Flag: "lookback-window", Type: "*string", Required: true},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "PromoteMode", Flag: "promote-mode", Type: "types.ModelPromoteMode", Required: false},
	{Name: "RetrainingFrequency", Flag: "retraining-frequency", Type: "*string", Required: true},
	{Name: "RetrainingStartDate", Flag: "retraining-start-date", Type: "*time.Time", Required: false},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
}

var fields_delete_inference_scheduler = []leanruntime.Field{
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
}

var fields_delete_label = []leanruntime.Field{
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
	{Name: "LabelId", Flag: "label-id", Type: "*string", Required: true},
}

var fields_delete_label_group = []leanruntime.Field{
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
}

var fields_delete_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_retraining_scheduler = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_describe_data_ingestion_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
}

var fields_describe_inference_scheduler = []leanruntime.Field{
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
}

var fields_describe_label = []leanruntime.Field{
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
	{Name: "LabelId", Flag: "label-id", Type: "*string", Required: true},
}

var fields_describe_label_group = []leanruntime.Field{
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
}

var fields_describe_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_describe_model_version = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "ModelVersion", Flag: "model-version", Type: "*int64", Required: true},
}

var fields_describe_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_retraining_scheduler = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_import_dataset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: false},
	{Name: "ServerSideKmsKeyId", Flag: "server-side-kms-key-id", Type: "*string", Required: false},
	{Name: "SourceDatasetArn", Flag: "source-dataset-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_import_model_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "InferenceDataImportStrategy", Flag: "inference-data-import-strategy", Type: "types.InferenceDataImportStrategy", Required: false},
	{Name: "LabelsInputConfiguration", Flag: "labels-input-configuration", Type: "*types.LabelsInputConfiguration", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ServerSideKmsKeyId", Flag: "server-side-kms-key-id", Type: "*string", Required: false},
	{Name: "SourceModelVersionArn", Flag: "source-model-version-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_data_ingestion_jobs = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.IngestionJobStatus", Required: false},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "DatasetNameBeginsWith", Flag: "dataset-name-begins-with", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_inference_events = []leanruntime.Field{
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
	{Name: "IntervalEndTime", Flag: "interval-end-time", Type: "*time.Time", Required: true},
	{Name: "IntervalStartTime", Flag: "interval-start-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_inference_executions = []leanruntime.Field{
	{Name: "DataEndTimeBefore", Flag: "data-end-time-before", Type: "*time.Time", Required: false},
	{Name: "DataStartTimeAfter", Flag: "data-start-time-after", Type: "*time.Time", Required: false},
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.InferenceExecutionStatus", Required: false},
}

var fields_list_inference_schedulers = []leanruntime.Field{
	{Name: "InferenceSchedulerNameBeginsWith", Flag: "inference-scheduler-name-begins-with", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.InferenceSchedulerStatus", Required: false},
}

var fields_list_label_groups = []leanruntime.Field{
	{Name: "LabelGroupNameBeginsWith", Flag: "label-group-name-begins-with", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_labels = []leanruntime.Field{
	{Name: "Equipment", Flag: "equipment", Type: "*string", Required: false},
	{Name: "FaultCode", Flag: "fault-code", Type: "*string", Required: false},
	{Name: "IntervalEndTime", Flag: "interval-end-time", Type: "*time.Time", Required: false},
	{Name: "IntervalStartTime", Flag: "interval-start-time", Type: "*time.Time", Required: false},
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_model_versions = []leanruntime.Field{
	{Name: "CreatedAtEndTime", Flag: "created-at-end-time", Type: "*time.Time", Required: false},
	{Name: "CreatedAtStartTime", Flag: "created-at-start-time", Type: "*time.Time", Required: false},
	{Name: "MaxModelVersion", Flag: "max-model-version", Type: "*int64", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MinModelVersion", Flag: "min-model-version", Type: "*int64", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.ModelVersionSourceType", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ModelVersionStatus", Required: false},
}

var fields_list_models = []leanruntime.Field{
	{Name: "DatasetNameBeginsWith", Flag: "dataset-name-begins-with", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelNameBeginsWith", Flag: "model-name-begins-with", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ModelStatus", Required: false},
}

var fields_list_retraining_schedulers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelNameBeginsWith", Flag: "model-name-begins-with", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RetrainingSchedulerStatus", Required: false},
}

var fields_list_sensor_statistics = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "IngestionJobId", Flag: "ingestion-job-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
}

var fields_start_data_ingestion_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "IngestionInputConfiguration", Flag: "ingestion-input-configuration", Type: "*types.IngestionInputConfiguration", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_start_inference_scheduler = []leanruntime.Field{
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
}

var fields_start_retraining_scheduler = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_stop_inference_scheduler = []leanruntime.Field{
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
}

var fields_stop_retraining_scheduler = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_active_model_version = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "ModelVersion", Flag: "model-version", Type: "*int64", Required: true},
}

var fields_update_inference_scheduler = []leanruntime.Field{
	{Name: "DataDelayOffsetInMinutes", Flag: "data-delay-offset-in-minutes", Type: "*int64", Required: false},
	{Name: "DataInputConfiguration", Flag: "data-input-configuration", Type: "*types.InferenceInputConfiguration", Required: false},
	{Name: "DataOutputConfiguration", Flag: "data-output-configuration", Type: "*types.InferenceOutputConfiguration", Required: false},
	{Name: "DataUploadFrequency", Flag: "data-upload-frequency", Type: "types.DataUploadFrequency", Required: false},
	{Name: "InferenceSchedulerName", Flag: "inference-scheduler-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_label_group = []leanruntime.Field{
	{Name: "FaultCodes", Flag: "fault-codes", Type: "[]string", Required: false},
	{Name: "LabelGroupName", Flag: "label-group-name", Type: "*string", Required: true},
}

var fields_update_model = []leanruntime.Field{
	{Name: "LabelsInputConfiguration", Flag: "labels-input-configuration", Type: "*types.LabelsInputConfiguration", Required: false},
	{Name: "ModelDiagnosticsOutputConfiguration", Flag: "model-diagnostics-output-configuration", Type: "*types.ModelDiagnosticsOutputConfiguration", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_retraining_scheduler = []leanruntime.Field{
	{Name: "LookbackWindow", Flag: "lookback-window", Type: "*string", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "PromoteMode", Flag: "promote-mode", Type: "types.ModelPromoteMode", Required: false},
	{Name: "RetrainingFrequency", Flag: "retraining-frequency", Type: "*string", Required: false},
	{Name: "RetrainingStartDate", Flag: "retraining-start-date", Type: "*time.Time", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-dataset": {
			Name:   "create-dataset",
			Fields: fields_create_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataset(ctx, input)
			},
		},
		"create-inference-scheduler": {
			Name:   "create-inference-scheduler",
			Fields: fields_create_inference_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInferenceSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_inference_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInferenceScheduler(ctx, input)
			},
		},
		"create-label": {
			Name:   "create-label",
			Fields: fields_create_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLabel(ctx, input)
			},
		},
		"create-label-group": {
			Name:   "create-label-group",
			Fields: fields_create_label_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLabelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_label_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLabelGroup(ctx, input)
			},
		},
		"create-model": {
			Name:   "create-model",
			Fields: fields_create_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModel(ctx, input)
			},
		},
		"create-retraining-scheduler": {
			Name:   "create-retraining-scheduler",
			Fields: fields_create_retraining_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRetrainingSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_retraining_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRetrainingScheduler(ctx, input)
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
		"delete-inference-scheduler": {
			Name:   "delete-inference-scheduler",
			Fields: fields_delete_inference_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInferenceSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inference_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInferenceScheduler(ctx, input)
			},
		},
		"delete-label": {
			Name:   "delete-label",
			Fields: fields_delete_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLabel(ctx, input)
			},
		},
		"delete-label-group": {
			Name:   "delete-label-group",
			Fields: fields_delete_label_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLabelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_label_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLabelGroup(ctx, input)
			},
		},
		"delete-model": {
			Name:   "delete-model",
			Fields: fields_delete_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModel(ctx, input)
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
		"delete-retraining-scheduler": {
			Name:   "delete-retraining-scheduler",
			Fields: fields_delete_retraining_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRetrainingSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_retraining_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRetrainingScheduler(ctx, input)
			},
		},
		"describe-data-ingestion-job": {
			Name:   "describe-data-ingestion-job",
			Fields: fields_describe_data_ingestion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataIngestionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_ingestion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataIngestionJob(ctx, input)
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
		"describe-inference-scheduler": {
			Name:   "describe-inference-scheduler",
			Fields: fields_describe_inference_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInferenceSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_inference_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInferenceScheduler(ctx, input)
			},
		},
		"describe-label": {
			Name:   "describe-label",
			Fields: fields_describe_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLabel(ctx, input)
			},
		},
		"describe-label-group": {
			Name:   "describe-label-group",
			Fields: fields_describe_label_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLabelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_label_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLabelGroup(ctx, input)
			},
		},
		"describe-model": {
			Name:   "describe-model",
			Fields: fields_describe_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModel(ctx, input)
			},
		},
		"describe-model-version": {
			Name:   "describe-model-version",
			Fields: fields_describe_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelVersion(ctx, input)
			},
		},
		"describe-resource-policy": {
			Name:   "describe-resource-policy",
			Fields: fields_describe_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicy(ctx, input)
			},
		},
		"describe-retraining-scheduler": {
			Name:   "describe-retraining-scheduler",
			Fields: fields_describe_retraining_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRetrainingSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_retraining_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRetrainingScheduler(ctx, input)
			},
		},
		"import-dataset": {
			Name:   "import-dataset",
			Fields: fields_import_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportDataset(ctx, input)
			},
		},
		"import-model-version": {
			Name:   "import-model-version",
			Fields: fields_import_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportModelVersion(ctx, input)
			},
		},
		"list-data-ingestion-jobs": {
			Name:   "list-data-ingestion-jobs",
			Fields: fields_list_data_ingestion_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataIngestionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_ingestion_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataIngestionJobs(ctx, input)
				}
				var results []*svc.ListDataIngestionJobsOutput
				p := svc.NewListDataIngestionJobsPaginator(client, input)
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
		"list-datasets": {
			Name:   "list-datasets",
			Fields: fields_list_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasets(ctx, input)
				}
				var results []*svc.ListDatasetsOutput
				p := svc.NewListDatasetsPaginator(client, input)
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
		"list-inference-events": {
			Name:   "list-inference-events",
			Fields: fields_list_inference_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceEvents(ctx, input)
				}
				var results []*svc.ListInferenceEventsOutput
				p := svc.NewListInferenceEventsPaginator(client, input)
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
		"list-inference-executions": {
			Name:   "list-inference-executions",
			Fields: fields_list_inference_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceExecutions(ctx, input)
				}
				var results []*svc.ListInferenceExecutionsOutput
				p := svc.NewListInferenceExecutionsPaginator(client, input)
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
		"list-inference-schedulers": {
			Name:   "list-inference-schedulers",
			Fields: fields_list_inference_schedulers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceSchedulersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_schedulers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceSchedulers(ctx, input)
				}
				var results []*svc.ListInferenceSchedulersOutput
				p := svc.NewListInferenceSchedulersPaginator(client, input)
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
		"list-label-groups": {
			Name:   "list-label-groups",
			Fields: fields_list_label_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLabelGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_label_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLabelGroups(ctx, input)
				}
				var results []*svc.ListLabelGroupsOutput
				p := svc.NewListLabelGroupsPaginator(client, input)
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
		"list-labels": {
			Name:   "list-labels",
			Fields: fields_list_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLabelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_labels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLabels(ctx, input)
				}
				var results []*svc.ListLabelsOutput
				p := svc.NewListLabelsPaginator(client, input)
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
		"list-model-versions": {
			Name:   "list-model-versions",
			Fields: fields_list_model_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelVersions(ctx, input)
				}
				var results []*svc.ListModelVersionsOutput
				p := svc.NewListModelVersionsPaginator(client, input)
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
		"list-models": {
			Name:   "list-models",
			Fields: fields_list_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModels(ctx, input)
				}
				var results []*svc.ListModelsOutput
				p := svc.NewListModelsPaginator(client, input)
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
		"list-retraining-schedulers": {
			Name:   "list-retraining-schedulers",
			Fields: fields_list_retraining_schedulers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRetrainingSchedulersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_retraining_schedulers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRetrainingSchedulers(ctx, input)
				}
				var results []*svc.ListRetrainingSchedulersOutput
				p := svc.NewListRetrainingSchedulersPaginator(client, input)
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
		"list-sensor-statistics": {
			Name:   "list-sensor-statistics",
			Fields: fields_list_sensor_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSensorStatisticsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sensor_statistics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSensorStatistics(ctx, input)
				}
				var results []*svc.ListSensorStatisticsOutput
				p := svc.NewListSensorStatisticsPaginator(client, input)
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
		"start-data-ingestion-job": {
			Name:   "start-data-ingestion-job",
			Fields: fields_start_data_ingestion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataIngestionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_ingestion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataIngestionJob(ctx, input)
			},
		},
		"start-inference-scheduler": {
			Name:   "start-inference-scheduler",
			Fields: fields_start_inference_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInferenceSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_inference_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInferenceScheduler(ctx, input)
			},
		},
		"start-retraining-scheduler": {
			Name:   "start-retraining-scheduler",
			Fields: fields_start_retraining_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRetrainingSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_retraining_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRetrainingScheduler(ctx, input)
			},
		},
		"stop-inference-scheduler": {
			Name:   "stop-inference-scheduler",
			Fields: fields_stop_inference_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopInferenceSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_inference_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopInferenceScheduler(ctx, input)
			},
		},
		"stop-retraining-scheduler": {
			Name:   "stop-retraining-scheduler",
			Fields: fields_stop_retraining_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRetrainingSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_retraining_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRetrainingScheduler(ctx, input)
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
		"update-active-model-version": {
			Name:   "update-active-model-version",
			Fields: fields_update_active_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActiveModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_active_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateActiveModelVersion(ctx, input)
			},
		},
		"update-inference-scheduler": {
			Name:   "update-inference-scheduler",
			Fields: fields_update_inference_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInferenceSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_inference_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInferenceScheduler(ctx, input)
			},
		},
		"update-label-group": {
			Name:   "update-label-group",
			Fields: fields_update_label_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLabelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_label_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLabelGroup(ctx, input)
			},
		},
		"update-model": {
			Name:   "update-model",
			Fields: fields_update_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModel(ctx, input)
			},
		},
		"update-retraining-scheduler": {
			Name:   "update-retraining-scheduler",
			Fields: fields_update_retraining_scheduler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRetrainingSchedulerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_retraining_scheduler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRetrainingScheduler(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lookoutequipment", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
