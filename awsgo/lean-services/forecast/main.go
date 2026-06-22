package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/forecast"
)

var fields_create_auto_predictor = []leanruntime.Field{
	{Name: "DataConfig", Flag: "data-config", Type: "*types.DataConfig", Required: false},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "*types.EncryptionConfig", Required: false},
	{Name: "ExplainPredictor", Flag: "explain-predictor", Type: "*bool", Required: false},
	{Name: "ForecastDimensions", Flag: "forecast-dimensions", Type: "[]string", Required: false},
	{Name: "ForecastFrequency", Flag: "forecast-frequency", Type: "*string", Required: false},
	{Name: "ForecastHorizon", Flag: "forecast-horizon", Type: "*int32", Required: false},
	{Name: "ForecastTypes", Flag: "forecast-types", Type: "[]string", Required: false},
	{Name: "MonitorConfig", Flag: "monitor-config", Type: "*types.MonitorConfig", Required: false},
	{Name: "OptimizationMetric", Flag: "optimization-metric", Type: "types.OptimizationMetric", Required: false},
	{Name: "PredictorName", Flag: "predictor-name", Type: "*string", Required: true},
	{Name: "ReferencePredictorArn", Flag: "reference-predictor-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeAlignmentBoundary", Flag: "time-alignment-boundary", Type: "*types.TimeAlignmentBoundary", Required: false},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "DataFrequency", Flag: "data-frequency", Type: "*string", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DatasetType", Flag: "dataset-type", Type: "types.DatasetType", Required: true},
	{Name: "Domain", Flag: "domain", Type: "types.Domain", Required: true},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "*types.EncryptionConfig", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*types.Schema", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dataset_group = []leanruntime.Field{
	{Name: "DatasetArns", Flag: "dataset-arns", Type: "[]string", Required: false},
	{Name: "DatasetGroupName", Flag: "dataset-group-name", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "types.Domain", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dataset_import_job = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: true},
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "DatasetImportJobName", Flag: "dataset-import-job-name", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "*string", Required: false},
	{Name: "GeolocationFormat", Flag: "geolocation-format", Type: "*string", Required: false},
	{Name: "ImportMode", Flag: "import-mode", Type: "types.ImportMode", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: false},
	{Name: "TimestampFormat", Flag: "timestamp-format", Type: "*string", Required: false},
	{Name: "UseGeolocationForTimeZone", Flag: "use-geolocation-for-time-zone", Type: "bool", Required: false},
}

var fields_create_explainability = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: false},
	{Name: "EnableVisualization", Flag: "enable-visualization", Type: "*bool", Required: false},
	{Name: "EndDateTime", Flag: "end-date-time", Type: "*string", Required: false},
	{Name: "ExplainabilityConfig", Flag: "explainability-config", Type: "*types.ExplainabilityConfig", Required: true},
	{Name: "ExplainabilityName", Flag: "explainability-name", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*types.Schema", Required: false},
	{Name: "StartDateTime", Flag: "start-date-time", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_explainability_export = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*types.DataDestination", Required: true},
	{Name: "ExplainabilityArn", Flag: "explainability-arn", Type: "*string", Required: true},
	{Name: "ExplainabilityExportName", Flag: "explainability-export-name", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_forecast = []leanruntime.Field{
	{Name: "ForecastName", Flag: "forecast-name", Type: "*string", Required: true},
	{Name: "ForecastTypes", Flag: "forecast-types", Type: "[]string", Required: false},
	{Name: "PredictorArn", Flag: "predictor-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeSeriesSelector", Flag: "time-series-selector", Type: "*types.TimeSeriesSelector", Required: false},
}

var fields_create_forecast_export_job = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*types.DataDestination", Required: true},
	{Name: "ForecastArn", Flag: "forecast-arn", Type: "*string", Required: true},
	{Name: "ForecastExportJobName", Flag: "forecast-export-job-name", Type: "*string", Required: true},
	{Name: "Format", Flag: "format", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_monitor = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_predictor = []leanruntime.Field{
	{Name: "AlgorithmArn", Flag: "algorithm-arn", Type: "*string", Required: false},
	{Name: "AutoMLOverrideStrategy", Flag: "auto-ml-override-strategy", Type: "types.AutoMLOverrideStrategy", Required: false},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "*types.EncryptionConfig", Required: false},
	{Name: "EvaluationParameters", Flag: "evaluation-parameters", Type: "*types.EvaluationParameters", Required: false},
	{Name: "FeaturizationConfig", Flag: "featurization-config", Type: "*types.FeaturizationConfig", Required: true},
	{Name: "ForecastHorizon", Flag: "forecast-horizon", Type: "*int32", Required: true},
	{Name: "ForecastTypes", Flag: "forecast-types", Type: "[]string", Required: false},
	{Name: "HPOConfig", Flag: "hpo-config", Type: "*types.HyperParameterTuningJobConfig", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "OptimizationMetric", Flag: "optimization-metric", Type: "types.OptimizationMetric", Required: false},
	{Name: "PerformAutoML", Flag: "perform-auto-ml", Type: "*bool", Required: false},
	{Name: "PerformHPO", Flag: "perform-hpo", Type: "*bool", Required: false},
	{Name: "PredictorName", Flag: "predictor-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingParameters", Flag: "training-parameters", Type: "map[string]string", Required: false},
}

var fields_create_predictor_backtest_export_job = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*types.DataDestination", Required: true},
	{Name: "Format", Flag: "format", Type: "*string", Required: false},
	{Name: "PredictorArn", Flag: "predictor-arn", Type: "*string", Required: true},
	{Name: "PredictorBacktestExportJobName", Flag: "predictor-backtest-export-job-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_what_if_analysis = []leanruntime.Field{
	{Name: "ForecastArn", Flag: "forecast-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeSeriesSelector", Flag: "time-series-selector", Type: "*types.TimeSeriesSelector", Required: false},
	{Name: "WhatIfAnalysisName", Flag: "what-if-analysis-name", Type: "*string", Required: true},
}

var fields_create_what_if_forecast = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeSeriesReplacementsDataSource", Flag: "time-series-replacements-data-source", Type: "*types.TimeSeriesReplacementsDataSource", Required: false},
	{Name: "TimeSeriesTransformations", Flag: "time-series-transformations", Type: "[]types.TimeSeriesTransformation", Required: false},
	{Name: "WhatIfAnalysisArn", Flag: "what-if-analysis-arn", Type: "*string", Required: true},
	{Name: "WhatIfForecastName", Flag: "what-if-forecast-name", Type: "*string", Required: true},
}

var fields_create_what_if_forecast_export = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*types.DataDestination", Required: true},
	{Name: "Format", Flag: "format", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WhatIfForecastArns", Flag: "what-if-forecast-arns", Type: "[]string", Required: true},
	{Name: "WhatIfForecastExportName", Flag: "what-if-forecast-export-name", Type: "*string", Required: true},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_delete_dataset_group = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
}

var fields_delete_dataset_import_job = []leanruntime.Field{
	{Name: "DatasetImportJobArn", Flag: "dataset-import-job-arn", Type: "*string", Required: true},
}

var fields_delete_explainability = []leanruntime.Field{
	{Name: "ExplainabilityArn", Flag: "explainability-arn", Type: "*string", Required: true},
}

var fields_delete_explainability_export = []leanruntime.Field{
	{Name: "ExplainabilityExportArn", Flag: "explainability-export-arn", Type: "*string", Required: true},
}

var fields_delete_forecast = []leanruntime.Field{
	{Name: "ForecastArn", Flag: "forecast-arn", Type: "*string", Required: true},
}

var fields_delete_forecast_export_job = []leanruntime.Field{
	{Name: "ForecastExportJobArn", Flag: "forecast-export-job-arn", Type: "*string", Required: true},
}

var fields_delete_monitor = []leanruntime.Field{
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: true},
}

var fields_delete_predictor = []leanruntime.Field{
	{Name: "PredictorArn", Flag: "predictor-arn", Type: "*string", Required: true},
}

var fields_delete_predictor_backtest_export_job = []leanruntime.Field{
	{Name: "PredictorBacktestExportJobArn", Flag: "predictor-backtest-export-job-arn", Type: "*string", Required: true},
}

var fields_delete_resource_tree = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_what_if_analysis = []leanruntime.Field{
	{Name: "WhatIfAnalysisArn", Flag: "what-if-analysis-arn", Type: "*string", Required: true},
}

var fields_delete_what_if_forecast = []leanruntime.Field{
	{Name: "WhatIfForecastArn", Flag: "what-if-forecast-arn", Type: "*string", Required: true},
}

var fields_delete_what_if_forecast_export = []leanruntime.Field{
	{Name: "WhatIfForecastExportArn", Flag: "what-if-forecast-export-arn", Type: "*string", Required: true},
}

var fields_describe_auto_predictor = []leanruntime.Field{
	{Name: "PredictorArn", Flag: "predictor-arn", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_describe_dataset_group = []leanruntime.Field{
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
}

var fields_describe_dataset_import_job = []leanruntime.Field{
	{Name: "DatasetImportJobArn", Flag: "dataset-import-job-arn", Type: "*string", Required: true},
}

var fields_describe_explainability = []leanruntime.Field{
	{Name: "ExplainabilityArn", Flag: "explainability-arn", Type: "*string", Required: true},
}

var fields_describe_explainability_export = []leanruntime.Field{
	{Name: "ExplainabilityExportArn", Flag: "explainability-export-arn", Type: "*string", Required: true},
}

var fields_describe_forecast = []leanruntime.Field{
	{Name: "ForecastArn", Flag: "forecast-arn", Type: "*string", Required: true},
}

var fields_describe_forecast_export_job = []leanruntime.Field{
	{Name: "ForecastExportJobArn", Flag: "forecast-export-job-arn", Type: "*string", Required: true},
}

var fields_describe_monitor = []leanruntime.Field{
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: true},
}

var fields_describe_predictor = []leanruntime.Field{
	{Name: "PredictorArn", Flag: "predictor-arn", Type: "*string", Required: true},
}

var fields_describe_predictor_backtest_export_job = []leanruntime.Field{
	{Name: "PredictorBacktestExportJobArn", Flag: "predictor-backtest-export-job-arn", Type: "*string", Required: true},
}

var fields_describe_what_if_analysis = []leanruntime.Field{
	{Name: "WhatIfAnalysisArn", Flag: "what-if-analysis-arn", Type: "*string", Required: true},
}

var fields_describe_what_if_forecast = []leanruntime.Field{
	{Name: "WhatIfForecastArn", Flag: "what-if-forecast-arn", Type: "*string", Required: true},
}

var fields_describe_what_if_forecast_export = []leanruntime.Field{
	{Name: "WhatIfForecastExportArn", Flag: "what-if-forecast-export-arn", Type: "*string", Required: true},
}

var fields_get_accuracy_metrics = []leanruntime.Field{
	{Name: "PredictorArn", Flag: "predictor-arn", Type: "*string", Required: true},
}

var fields_list_dataset_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dataset_import_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_explainabilities = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_explainability_exports = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_forecast_export_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_forecasts = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_monitor_evaluations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_monitors = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_predictor_backtest_export_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_predictors = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_what_if_analyses = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_what_if_forecast_exports = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_what_if_forecasts = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_resume_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_stop_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_dataset_group = []leanruntime.Field{
	{Name: "DatasetArns", Flag: "dataset-arns", Type: "[]string", Required: true},
	{Name: "DatasetGroupArn", Flag: "dataset-group-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-auto-predictor": {
			Name:   "create-auto-predictor",
			Fields: fields_create_auto_predictor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutoPredictorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_auto_predictor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutoPredictor(ctx, input)
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
		"create-explainability": {
			Name:   "create-explainability",
			Fields: fields_create_explainability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExplainabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_explainability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExplainability(ctx, input)
			},
		},
		"create-explainability-export": {
			Name:   "create-explainability-export",
			Fields: fields_create_explainability_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExplainabilityExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_explainability_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExplainabilityExport(ctx, input)
			},
		},
		"create-forecast": {
			Name:   "create-forecast",
			Fields: fields_create_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateForecast(ctx, input)
			},
		},
		"create-forecast-export-job": {
			Name:   "create-forecast-export-job",
			Fields: fields_create_forecast_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateForecastExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_forecast_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateForecastExportJob(ctx, input)
			},
		},
		"create-monitor": {
			Name:   "create-monitor",
			Fields: fields_create_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMonitor(ctx, input)
			},
		},
		"create-predictor": {
			Name:   "create-predictor",
			Fields: fields_create_predictor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePredictorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_predictor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePredictor(ctx, input)
			},
		},
		"create-predictor-backtest-export-job": {
			Name:   "create-predictor-backtest-export-job",
			Fields: fields_create_predictor_backtest_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePredictorBacktestExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_predictor_backtest_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePredictorBacktestExportJob(ctx, input)
			},
		},
		"create-what-if-analysis": {
			Name:   "create-what-if-analysis",
			Fields: fields_create_what_if_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWhatIfAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_what_if_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWhatIfAnalysis(ctx, input)
			},
		},
		"create-what-if-forecast": {
			Name:   "create-what-if-forecast",
			Fields: fields_create_what_if_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWhatIfForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_what_if_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWhatIfForecast(ctx, input)
			},
		},
		"create-what-if-forecast-export": {
			Name:   "create-what-if-forecast-export",
			Fields: fields_create_what_if_forecast_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWhatIfForecastExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_what_if_forecast_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWhatIfForecastExport(ctx, input)
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
		"delete-dataset-import-job": {
			Name:   "delete-dataset-import-job",
			Fields: fields_delete_dataset_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDatasetImportJob(ctx, input)
			},
		},
		"delete-explainability": {
			Name:   "delete-explainability",
			Fields: fields_delete_explainability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExplainabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_explainability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExplainability(ctx, input)
			},
		},
		"delete-explainability-export": {
			Name:   "delete-explainability-export",
			Fields: fields_delete_explainability_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExplainabilityExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_explainability_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExplainabilityExport(ctx, input)
			},
		},
		"delete-forecast": {
			Name:   "delete-forecast",
			Fields: fields_delete_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteForecast(ctx, input)
			},
		},
		"delete-forecast-export-job": {
			Name:   "delete-forecast-export-job",
			Fields: fields_delete_forecast_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteForecastExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_forecast_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteForecastExportJob(ctx, input)
			},
		},
		"delete-monitor": {
			Name:   "delete-monitor",
			Fields: fields_delete_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMonitor(ctx, input)
			},
		},
		"delete-predictor": {
			Name:   "delete-predictor",
			Fields: fields_delete_predictor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePredictorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_predictor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePredictor(ctx, input)
			},
		},
		"delete-predictor-backtest-export-job": {
			Name:   "delete-predictor-backtest-export-job",
			Fields: fields_delete_predictor_backtest_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePredictorBacktestExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_predictor_backtest_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePredictorBacktestExportJob(ctx, input)
			},
		},
		"delete-resource-tree": {
			Name:   "delete-resource-tree",
			Fields: fields_delete_resource_tree,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceTreeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_tree, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceTree(ctx, input)
			},
		},
		"delete-what-if-analysis": {
			Name:   "delete-what-if-analysis",
			Fields: fields_delete_what_if_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWhatIfAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_what_if_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWhatIfAnalysis(ctx, input)
			},
		},
		"delete-what-if-forecast": {
			Name:   "delete-what-if-forecast",
			Fields: fields_delete_what_if_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWhatIfForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_what_if_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWhatIfForecast(ctx, input)
			},
		},
		"delete-what-if-forecast-export": {
			Name:   "delete-what-if-forecast-export",
			Fields: fields_delete_what_if_forecast_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWhatIfForecastExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_what_if_forecast_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWhatIfForecastExport(ctx, input)
			},
		},
		"describe-auto-predictor": {
			Name:   "describe-auto-predictor",
			Fields: fields_describe_auto_predictor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoPredictorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_auto_predictor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAutoPredictor(ctx, input)
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
		"describe-explainability": {
			Name:   "describe-explainability",
			Fields: fields_describe_explainability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExplainabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_explainability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExplainability(ctx, input)
			},
		},
		"describe-explainability-export": {
			Name:   "describe-explainability-export",
			Fields: fields_describe_explainability_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExplainabilityExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_explainability_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExplainabilityExport(ctx, input)
			},
		},
		"describe-forecast": {
			Name:   "describe-forecast",
			Fields: fields_describe_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeForecast(ctx, input)
			},
		},
		"describe-forecast-export-job": {
			Name:   "describe-forecast-export-job",
			Fields: fields_describe_forecast_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeForecastExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_forecast_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeForecastExportJob(ctx, input)
			},
		},
		"describe-monitor": {
			Name:   "describe-monitor",
			Fields: fields_describe_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMonitor(ctx, input)
			},
		},
		"describe-predictor": {
			Name:   "describe-predictor",
			Fields: fields_describe_predictor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePredictorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_predictor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePredictor(ctx, input)
			},
		},
		"describe-predictor-backtest-export-job": {
			Name:   "describe-predictor-backtest-export-job",
			Fields: fields_describe_predictor_backtest_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePredictorBacktestExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_predictor_backtest_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePredictorBacktestExportJob(ctx, input)
			},
		},
		"describe-what-if-analysis": {
			Name:   "describe-what-if-analysis",
			Fields: fields_describe_what_if_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWhatIfAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_what_if_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWhatIfAnalysis(ctx, input)
			},
		},
		"describe-what-if-forecast": {
			Name:   "describe-what-if-forecast",
			Fields: fields_describe_what_if_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWhatIfForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_what_if_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWhatIfForecast(ctx, input)
			},
		},
		"describe-what-if-forecast-export": {
			Name:   "describe-what-if-forecast-export",
			Fields: fields_describe_what_if_forecast_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWhatIfForecastExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_what_if_forecast_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWhatIfForecastExport(ctx, input)
			},
		},
		"get-accuracy-metrics": {
			Name:   "get-accuracy-metrics",
			Fields: fields_get_accuracy_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccuracyMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_accuracy_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccuracyMetrics(ctx, input)
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
		"list-explainabilities": {
			Name:   "list-explainabilities",
			Fields: fields_list_explainabilities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExplainabilitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_explainabilities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExplainabilities(ctx, input)
				}
				var results []*svc.ListExplainabilitiesOutput
				p := svc.NewListExplainabilitiesPaginator(client, input)
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
		"list-explainability-exports": {
			Name:   "list-explainability-exports",
			Fields: fields_list_explainability_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExplainabilityExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_explainability_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExplainabilityExports(ctx, input)
				}
				var results []*svc.ListExplainabilityExportsOutput
				p := svc.NewListExplainabilityExportsPaginator(client, input)
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
		"list-forecast-export-jobs": {
			Name:   "list-forecast-export-jobs",
			Fields: fields_list_forecast_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListForecastExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_forecast_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListForecastExportJobs(ctx, input)
				}
				var results []*svc.ListForecastExportJobsOutput
				p := svc.NewListForecastExportJobsPaginator(client, input)
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
		"list-forecasts": {
			Name:   "list-forecasts",
			Fields: fields_list_forecasts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListForecastsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_forecasts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListForecasts(ctx, input)
				}
				var results []*svc.ListForecastsOutput
				p := svc.NewListForecastsPaginator(client, input)
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
		"list-monitor-evaluations": {
			Name:   "list-monitor-evaluations",
			Fields: fields_list_monitor_evaluations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitorEvaluationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitor_evaluations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitorEvaluations(ctx, input)
				}
				var results []*svc.ListMonitorEvaluationsOutput
				p := svc.NewListMonitorEvaluationsPaginator(client, input)
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
		"list-monitors": {
			Name:   "list-monitors",
			Fields: fields_list_monitors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitors(ctx, input)
				}
				var results []*svc.ListMonitorsOutput
				p := svc.NewListMonitorsPaginator(client, input)
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
		"list-predictor-backtest-export-jobs": {
			Name:   "list-predictor-backtest-export-jobs",
			Fields: fields_list_predictor_backtest_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPredictorBacktestExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_predictor_backtest_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPredictorBacktestExportJobs(ctx, input)
				}
				var results []*svc.ListPredictorBacktestExportJobsOutput
				p := svc.NewListPredictorBacktestExportJobsPaginator(client, input)
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
		"list-predictors": {
			Name:   "list-predictors",
			Fields: fields_list_predictors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPredictorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_predictors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPredictors(ctx, input)
				}
				var results []*svc.ListPredictorsOutput
				p := svc.NewListPredictorsPaginator(client, input)
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
		"list-what-if-analyses": {
			Name:   "list-what-if-analyses",
			Fields: fields_list_what_if_analyses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWhatIfAnalysesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_what_if_analyses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWhatIfAnalyses(ctx, input)
				}
				var results []*svc.ListWhatIfAnalysesOutput
				p := svc.NewListWhatIfAnalysesPaginator(client, input)
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
		"list-what-if-forecast-exports": {
			Name:   "list-what-if-forecast-exports",
			Fields: fields_list_what_if_forecast_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWhatIfForecastExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_what_if_forecast_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWhatIfForecastExports(ctx, input)
				}
				var results []*svc.ListWhatIfForecastExportsOutput
				p := svc.NewListWhatIfForecastExportsPaginator(client, input)
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
		"list-what-if-forecasts": {
			Name:   "list-what-if-forecasts",
			Fields: fields_list_what_if_forecasts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWhatIfForecastsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_what_if_forecasts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWhatIfForecasts(ctx, input)
				}
				var results []*svc.ListWhatIfForecastsOutput
				p := svc.NewListWhatIfForecastsPaginator(client, input)
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
		"resume-resource": {
			Name:   "resume-resource",
			Fields: fields_resume_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeResource(ctx, input)
			},
		},
		"stop-resource": {
			Name:   "stop-resource",
			Fields: fields_stop_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopResource(ctx, input)
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
		"update-dataset-group": {
			Name:   "update-dataset-group",
			Fields: fields_update_dataset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dataset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDatasetGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("forecast", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
