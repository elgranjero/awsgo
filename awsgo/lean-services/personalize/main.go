package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/personalize"
)

var fields_create_batch_inference_job = []leanruntime.Field{
	{Name: "BatchInferenceJobConfig", Flag: "batch-inference-job-config", Type: "*types.BatchInferenceJobConfig", Required: false},
	{Name: "BatchInferenceJobMode", Flag: "batch-inference-job-mode", Type: "types.BatchInferenceJobMode", Required: false},
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "JobInput", Flag: "job-input", Type: "*types.BatchInferenceJobInput", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobOutput", Flag: "job-output", Type: "*types.BatchInferenceJobOutput", Required: true},
	{Name: "NumResults", Flag: "num-results", Type: "*int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThemeGenerationConfig", Flag: "theme-generation-config", Type: "*types.ThemeGenerationConfig", Required: false},
}

var fields_create_batch_segment_job = []leanruntime.Field{
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "JobInput", Flag: "job-input", Type: "*types.BatchSegmentJobInput", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobOutput", Flag: "job-output", Type: "*types.BatchSegmentJobOutput", Required: true},
	{Name: "NumResults", Flag: "num-results", Type: "*int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_campaign = []leanruntime.Field{
	{Name: "CampaignConfig", Flag: "campaign-config", Type: "*types.CampaignConfig", Required: false},
	{Name: "MinProvisionedTPS", Flag: "min-provisioned-tps", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_data_deletion_job = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: true},
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "DatasetType", Flag: "dataset-type", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dataset_export_job = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "IngestionMode", Flag: "ingestion-mode", Type: "types.IngestionMode", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobOutput", Flag: "job-output", Type: "*types.DatasetExportJobOutput", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dataset_group = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "types.Domain", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dataset_import_job = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: true},
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "ImportMode", Flag: "import-mode", Type: "types.ImportMode", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "PublishAttributionMetricsToS3", Flag: "publish-attribution-metrics-to-s3", Type: "*bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_event_tracker = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_filter = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_metric_attribution = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.MetricAttribute", Required: true},
	{Name: "MetricsOutputConfig", Flag: "metrics-output-config", Type: "*types.MetricAttributionOutput", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_recommender = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecipeArn", Flag: "recipe-arn", Type: "*string", Required: true},
	{Name: "RecommenderConfig", Flag: "recommender-config", Type: "*types.RecommenderConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_schema = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "types.Domain", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: true},
}

var fields_create_solution = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
	{Name: "EventType", Flag: "event-type", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PerformAutoML", Flag: "perform-auto-ml", Type: "bool", Required: false},
	{Name: "PerformAutoTraining", Flag: "perform-auto-training", Type: "*bool", Required: false},
	{Name: "PerformHPO", Flag: "perform-hpo", Type: "*bool", Required: false},
	{Name: "PerformIncrementalUpdate", Flag: "perform-incremental-update", Type: "*bool", Required: false},
	{Name: "RecipeArn", Flag: "recipe-arn", Type: "*string", Required: false},
	{Name: "SolutionConfig", Flag: "solution-config", Type: "*types.SolutionConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_solution_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SolutionArn", Flag: "solution-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingMode", Flag: "training-mode", Type: "types.TrainingMode", Required: false},
}

var fields_delete_campaign = []leanruntime.Field{
	{Name: "CampaignArn", Flag: "campaign-arn", Type: "*string", Required: true},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_delete_dataset_group = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
}

var fields_delete_event_tracker = []leanruntime.Field{
	{Name: "EventTrackerArn", Flag: "event-tracker-arn", Type: "*string", Required: true},
}

var fields_delete_filter = []leanruntime.Field{
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: true},
}

var fields_delete_metric_attribution = []leanruntime.Field{
	{Name: "MetricAttributionArn", Flag: "metric-attribution-arn", Type: "*string", Required: true},
}

var fields_delete_recommender = []leanruntime.Field{
	{Name: "RecommenderArn", Flag: "recommender-arn", Type: "*string", Required: true},
}

var fields_delete_schema = []leanruntime.Field{
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_delete_solution = []leanruntime.Field{
	{Name: "SolutionArn", Flag: "solution-arn", Type: "*string", Required: true},
}

var fields_describe_algorithm = []leanruntime.Field{
	{Name: "AlgorithmArn", Flag: "algorithm-arn", Type: "*string", Required: true},
}

var fields_describe_batch_inference_job = []leanruntime.Field{
	{Name: "BatchInferenceJobArn", Flag: "batch-inference-job-arn", Type: "*string", Required: true},
}

var fields_describe_batch_segment_job = []leanruntime.Field{
	{Name: "BatchSegmentJobArn", Flag: "batch-segment-job-arn", Type: "*string", Required: true},
}

var fields_describe_campaign = []leanruntime.Field{
	{Name: "CampaignArn", Flag: "campaign-arn", Type: "*string", Required: true},
}

var fields_describe_data_deletion_job = []leanruntime.Field{
	{Name: "DataDeletionJobArn", Flag: "data-deletion-job-arn", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_describe_dataset_export_job = []leanruntime.Field{
	{Name: "DatasetExportJobArn", Flag: "dataset-export-job-arn", Type: "*string", Required: true},
}

var fields_describe_dataset_group = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
}

var fields_describe_dataset_import_job = []leanruntime.Field{
	{Name: "DatasetImportJobArn", Flag: "dataset-import-job-arn", Type: "*string", Required: true},
}

var fields_describe_event_tracker = []leanruntime.Field{
	{Name: "EventTrackerArn", Flag: "event-tracker-arn", Type: "*string", Required: true},
}

var fields_describe_feature_transformation = []leanruntime.Field{
	{Name: "FeatureTransformationArn", Flag: "feature-transformation-arn", Type: "*string", Required: true},
}

var fields_describe_filter = []leanruntime.Field{
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: true},
}

var fields_describe_metric_attribution = []leanruntime.Field{
	{Name: "MetricAttributionArn", Flag: "metric-attribution-arn", Type: "*string", Required: true},
}

var fields_describe_recipe = []leanruntime.Field{
	{Name: "RecipeArn", Flag: "recipe-arn", Type: "*string", Required: true},
}

var fields_describe_recommender = []leanruntime.Field{
	{Name: "RecommenderArn", Flag: "recommender-arn", Type: "*string", Required: true},
}

var fields_describe_schema = []leanruntime.Field{
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_describe_solution = []leanruntime.Field{
	{Name: "SolutionArn", Flag: "solution-arn", Type: "*string", Required: true},
}

var fields_describe_solution_version = []leanruntime.Field{
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: true},
}

var fields_get_solution_metrics = []leanruntime.Field{
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: true},
}

var fields_list_batch_inference_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: false},
}

var fields_list_batch_segment_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: false},
}

var fields_list_campaigns = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SolutionArn", Flag: "solution-arn", Type: "*string", Required: false},
}

var fields_list_data_deletion_jobs = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dataset_export_jobs = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dataset_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dataset_import_jobs = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_trackers = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_filters = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_metric_attribution_metrics = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricAttributionArn", Flag: "metric-attribution-arn", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_metric_attributions = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recipes = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "types.Domain", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecipeProvider", Flag: "recipe-provider", Type: "types.RecipeProvider", Required: false},
}

var fields_list_recommenders = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_schemas = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_solution_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SolutionArn", Flag: "solution-arn", Type: "*string", Required: false},
}

var fields_list_solutions = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_recommender = []leanruntime.Field{
	{Name: "RecommenderArn", Flag: "recommender-arn", Type: "*string", Required: true},
}

var fields_stop_recommender = []leanruntime.Field{
	{Name: "RecommenderArn", Flag: "recommender-arn", Type: "*string", Required: true},
}

var fields_stop_solution_version_creation = []leanruntime.Field{
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_campaign = []leanruntime.Field{
	{Name: "CampaignArn", Flag: "campaign-arn", Type: "*string", Required: true},
	{Name: "CampaignConfig", Flag: "campaign-config", Type: "*types.CampaignConfig", Required: false},
	{Name: "MinProvisionedTPS", Flag: "min-provisioned-tps", Type: "*int32", Required: false},
	{Name: "SolutionVersionArn", Flag: "solution-version-arn", Type: "*string", Required: false},
}

var fields_update_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_update_metric_attribution = []leanruntime.Field{
	{Name: "AddMetrics", Flag: "add-metrics", Type: "[]types.MetricAttribute", Required: false},
	{Name: "MetricAttributionArn", Flag: "metric-attribution-arn", Type: "*string", Required: false},
	{Name: "MetricsOutputConfig", Flag: "metrics-output-config", Type: "*types.MetricAttributionOutput", Required: false},
	{Name: "RemoveMetrics", Flag: "remove-metrics", Type: "[]string", Required: false},
}

var fields_update_recommender = []leanruntime.Field{
	{Name: "RecommenderArn", Flag: "recommender-arn", Type: "*string", Required: true},
	{Name: "RecommenderConfig", Flag: "recommender-config", Type: "*types.RecommenderConfig", Required: true},
}

var fields_update_solution = []leanruntime.Field{
	{Name: "PerformAutoTraining", Flag: "perform-auto-training", Type: "*bool", Required: false},
	{Name: "PerformIncrementalUpdate", Flag: "perform-incremental-update", Type: "*bool", Required: false},
	{Name: "SolutionArn", Flag: "solution-arn", Type: "*string", Required: true},
	{Name: "SolutionUpdateConfig", Flag: "solution-update-config", Type: "*types.SolutionUpdateConfig", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-batch-inference-job": {
			Name:   "create-batch-inference-job",
			Fields: fields_create_batch_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBatchInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_batch_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBatchInferenceJob(ctx, input)
			},
		},
		"create-batch-segment-job": {
			Name:   "create-batch-segment-job",
			Fields: fields_create_batch_segment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBatchSegmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_batch_segment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBatchSegmentJob(ctx, input)
			},
		},
		"create-campaign": {
			Name:   "create-campaign",
			Fields: fields_create_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCampaign(ctx, input)
			},
		},
		"create-data-deletion-job": {
			Name:   "create-data-deletion-job",
			Fields: fields_create_data_deletion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataDeletionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_deletion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataDeletionJob(ctx, input)
			},
		},
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
		"create-dataset-export-job": {
			Name:   "create-dataset-export-job",
			Fields: fields_create_dataset_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDatasetExportJob(ctx, input)
			},
		},
		"create-dataset-group": {
			Name:   "create-dataset-group",
			Fields: fields_create_dataset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDatasetGroup(ctx, input)
			},
		},
		"create-dataset-import-job": {
			Name:   "create-dataset-import-job",
			Fields: fields_create_dataset_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDatasetImportJob(ctx, input)
			},
		},
		"create-event-tracker": {
			Name:   "create-event-tracker",
			Fields: fields_create_event_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventTracker(ctx, input)
			},
		},
		"create-filter": {
			Name:   "create-filter",
			Fields: fields_create_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFilter(ctx, input)
			},
		},
		"create-metric-attribution": {
			Name:   "create-metric-attribution",
			Fields: fields_create_metric_attribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMetricAttributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_metric_attribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMetricAttribution(ctx, input)
			},
		},
		"create-recommender": {
			Name:   "create-recommender",
			Fields: fields_create_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecommender(ctx, input)
			},
		},
		"create-schema": {
			Name:   "create-schema",
			Fields: fields_create_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchema(ctx, input)
			},
		},
		"create-solution": {
			Name:   "create-solution",
			Fields: fields_create_solution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSolutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_solution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSolution(ctx, input)
			},
		},
		"create-solution-version": {
			Name:   "create-solution-version",
			Fields: fields_create_solution_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSolutionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_solution_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSolutionVersion(ctx, input)
			},
		},
		"delete-campaign": {
			Name:   "delete-campaign",
			Fields: fields_delete_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCampaign(ctx, input)
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
		"delete-dataset-group": {
			Name:   "delete-dataset-group",
			Fields: fields_delete_dataset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDatasetGroup(ctx, input)
			},
		},
		"delete-event-tracker": {
			Name:   "delete-event-tracker",
			Fields: fields_delete_event_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventTracker(ctx, input)
			},
		},
		"delete-filter": {
			Name:   "delete-filter",
			Fields: fields_delete_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFilter(ctx, input)
			},
		},
		"delete-metric-attribution": {
			Name:   "delete-metric-attribution",
			Fields: fields_delete_metric_attribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMetricAttributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_metric_attribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMetricAttribution(ctx, input)
			},
		},
		"delete-recommender": {
			Name:   "delete-recommender",
			Fields: fields_delete_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecommender(ctx, input)
			},
		},
		"delete-schema": {
			Name:   "delete-schema",
			Fields: fields_delete_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchema(ctx, input)
			},
		},
		"delete-solution": {
			Name:   "delete-solution",
			Fields: fields_delete_solution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSolutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_solution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSolution(ctx, input)
			},
		},
		"describe-algorithm": {
			Name:   "describe-algorithm",
			Fields: fields_describe_algorithm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlgorithmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_algorithm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlgorithm(ctx, input)
			},
		},
		"describe-batch-inference-job": {
			Name:   "describe-batch-inference-job",
			Fields: fields_describe_batch_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBatchInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_batch_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBatchInferenceJob(ctx, input)
			},
		},
		"describe-batch-segment-job": {
			Name:   "describe-batch-segment-job",
			Fields: fields_describe_batch_segment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBatchSegmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_batch_segment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBatchSegmentJob(ctx, input)
			},
		},
		"describe-campaign": {
			Name:   "describe-campaign",
			Fields: fields_describe_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCampaign(ctx, input)
			},
		},
		"describe-data-deletion-job": {
			Name:   "describe-data-deletion-job",
			Fields: fields_describe_data_deletion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataDeletionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_deletion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataDeletionJob(ctx, input)
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
		"describe-dataset-export-job": {
			Name:   "describe-dataset-export-job",
			Fields: fields_describe_dataset_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDatasetExportJob(ctx, input)
			},
		},
		"describe-dataset-group": {
			Name:   "describe-dataset-group",
			Fields: fields_describe_dataset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDatasetGroup(ctx, input)
			},
		},
		"describe-dataset-import-job": {
			Name:   "describe-dataset-import-job",
			Fields: fields_describe_dataset_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDatasetImportJob(ctx, input)
			},
		},
		"describe-event-tracker": {
			Name:   "describe-event-tracker",
			Fields: fields_describe_event_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventTracker(ctx, input)
			},
		},
		"describe-feature-transformation": {
			Name:   "describe-feature-transformation",
			Fields: fields_describe_feature_transformation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFeatureTransformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_feature_transformation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFeatureTransformation(ctx, input)
			},
		},
		"describe-filter": {
			Name:   "describe-filter",
			Fields: fields_describe_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFilter(ctx, input)
			},
		},
		"describe-metric-attribution": {
			Name:   "describe-metric-attribution",
			Fields: fields_describe_metric_attribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetricAttributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_metric_attribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMetricAttribution(ctx, input)
			},
		},
		"describe-recipe": {
			Name:   "describe-recipe",
			Fields: fields_describe_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRecipe(ctx, input)
			},
		},
		"describe-recommender": {
			Name:   "describe-recommender",
			Fields: fields_describe_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRecommender(ctx, input)
			},
		},
		"describe-schema": {
			Name:   "describe-schema",
			Fields: fields_describe_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSchema(ctx, input)
			},
		},
		"describe-solution": {
			Name:   "describe-solution",
			Fields: fields_describe_solution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSolutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_solution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSolution(ctx, input)
			},
		},
		"describe-solution-version": {
			Name:   "describe-solution-version",
			Fields: fields_describe_solution_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSolutionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_solution_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSolutionVersion(ctx, input)
			},
		},
		"get-solution-metrics": {
			Name:   "get-solution-metrics",
			Fields: fields_get_solution_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolutionMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_solution_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolutionMetrics(ctx, input)
			},
		},
		"list-batch-inference-jobs": {
			Name:   "list-batch-inference-jobs",
			Fields: fields_list_batch_inference_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBatchInferenceJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_batch_inference_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBatchInferenceJobs(ctx, input)
				}
				var results []*svc.ListBatchInferenceJobsOutput
				p := svc.NewListBatchInferenceJobsPaginator(client, input)
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
		"list-batch-segment-jobs": {
			Name:   "list-batch-segment-jobs",
			Fields: fields_list_batch_segment_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBatchSegmentJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_batch_segment_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBatchSegmentJobs(ctx, input)
				}
				var results []*svc.ListBatchSegmentJobsOutput
				p := svc.NewListBatchSegmentJobsPaginator(client, input)
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
		"list-campaigns": {
			Name:   "list-campaigns",
			Fields: fields_list_campaigns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCampaignsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_campaigns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCampaigns(ctx, input)
				}
				var results []*svc.ListCampaignsOutput
				p := svc.NewListCampaignsPaginator(client, input)
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
		"list-data-deletion-jobs": {
			Name:   "list-data-deletion-jobs",
			Fields: fields_list_data_deletion_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataDeletionJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_data_deletion_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDataDeletionJobs(ctx, input)
			},
		},
		"list-dataset-export-jobs": {
			Name:   "list-dataset-export-jobs",
			Fields: fields_list_dataset_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dataset_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasetExportJobs(ctx, input)
				}
				var results []*svc.ListDatasetExportJobsOutput
				p := svc.NewListDatasetExportJobsPaginator(client, input)
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
		"list-dataset-groups": {
			Name:   "list-dataset-groups",
			Fields: fields_list_dataset_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dataset_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasetGroups(ctx, input)
				}
				var results []*svc.ListDatasetGroupsOutput
				p := svc.NewListDatasetGroupsPaginator(client, input)
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
		"list-dataset-import-jobs": {
			Name:   "list-dataset-import-jobs",
			Fields: fields_list_dataset_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dataset_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasetImportJobs(ctx, input)
				}
				var results []*svc.ListDatasetImportJobsOutput
				p := svc.NewListDatasetImportJobsPaginator(client, input)
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
		"list-event-trackers": {
			Name:   "list-event-trackers",
			Fields: fields_list_event_trackers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventTrackersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_trackers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventTrackers(ctx, input)
				}
				var results []*svc.ListEventTrackersOutput
				p := svc.NewListEventTrackersPaginator(client, input)
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
		"list-filters": {
			Name:   "list-filters",
			Fields: fields_list_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFilters(ctx, input)
				}
				var results []*svc.ListFiltersOutput
				p := svc.NewListFiltersPaginator(client, input)
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
		"list-metric-attribution-metrics": {
			Name:   "list-metric-attribution-metrics",
			Fields: fields_list_metric_attribution_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetricAttributionMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metric_attribution_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetricAttributionMetrics(ctx, input)
				}
				var results []*svc.ListMetricAttributionMetricsOutput
				p := svc.NewListMetricAttributionMetricsPaginator(client, input)
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
		"list-metric-attributions": {
			Name:   "list-metric-attributions",
			Fields: fields_list_metric_attributions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetricAttributionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metric_attributions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetricAttributions(ctx, input)
				}
				var results []*svc.ListMetricAttributionsOutput
				p := svc.NewListMetricAttributionsPaginator(client, input)
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
		"list-recipes": {
			Name:   "list-recipes",
			Fields: fields_list_recipes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecipesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recipes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecipes(ctx, input)
				}
				var results []*svc.ListRecipesOutput
				p := svc.NewListRecipesPaginator(client, input)
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
		"list-recommenders": {
			Name:   "list-recommenders",
			Fields: fields_list_recommenders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommenders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommenders(ctx, input)
				}
				var results []*svc.ListRecommendersOutput
				p := svc.NewListRecommendersPaginator(client, input)
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
		"list-schemas": {
			Name:   "list-schemas",
			Fields: fields_list_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemas(ctx, input)
				}
				var results []*svc.ListSchemasOutput
				p := svc.NewListSchemasPaginator(client, input)
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
		"list-solution-versions": {
			Name:   "list-solution-versions",
			Fields: fields_list_solution_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolutionVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_solution_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolutionVersions(ctx, input)
				}
				var results []*svc.ListSolutionVersionsOutput
				p := svc.NewListSolutionVersionsPaginator(client, input)
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
		"list-solutions": {
			Name:   "list-solutions",
			Fields: fields_list_solutions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_solutions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolutions(ctx, input)
				}
				var results []*svc.ListSolutionsOutput
				p := svc.NewListSolutionsPaginator(client, input)
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
		"start-recommender": {
			Name:   "start-recommender",
			Fields: fields_start_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRecommender(ctx, input)
			},
		},
		"stop-recommender": {
			Name:   "stop-recommender",
			Fields: fields_stop_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRecommender(ctx, input)
			},
		},
		"stop-solution-version-creation": {
			Name:   "stop-solution-version-creation",
			Fields: fields_stop_solution_version_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSolutionVersionCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_solution_version_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSolutionVersionCreation(ctx, input)
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
		"update-campaign": {
			Name:   "update-campaign",
			Fields: fields_update_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCampaign(ctx, input)
			},
		},
		"update-dataset": {
			Name:   "update-dataset",
			Fields: fields_update_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataset(ctx, input)
			},
		},
		"update-metric-attribution": {
			Name:   "update-metric-attribution",
			Fields: fields_update_metric_attribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMetricAttributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_metric_attribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMetricAttribution(ctx, input)
			},
		},
		"update-recommender": {
			Name:   "update-recommender",
			Fields: fields_update_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecommender(ctx, input)
			},
		},
		"update-solution": {
			Name:   "update-solution",
			Fields: fields_update_solution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSolutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_solution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSolution(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("personalize", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
