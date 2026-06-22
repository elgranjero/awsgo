package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/machinelearning"
)

var fields_add_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.TaggableResourceType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_create_batch_prediction = []leanruntime.Field{
	{Name: "BatchPredictionDataSourceId", Flag: "batch-prediction-data-source-id", Type: "*string", Required: true},
	{Name: "BatchPredictionId", Flag: "batch-prediction-id", Type: "*string", Required: true},
	{Name: "BatchPredictionName", Flag: "batch-prediction-name", Type: "*string", Required: false},
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
	{Name: "OutputUri", Flag: "output-uri", Type: "*string", Required: true},
}

var fields_create_data_source_from_rds = []leanruntime.Field{
	{Name: "ComputeStatistics", Flag: "compute-statistics", Type: "bool", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: false},
	{Name: "RDSData", Flag: "rds-data", Type: "*types.RDSDataSpec", Required: true},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_create_data_source_from_redshift = []leanruntime.Field{
	{Name: "ComputeStatistics", Flag: "compute-statistics", Type: "bool", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: false},
	{Name: "DataSpec", Flag: "data-spec", Type: "*types.RedshiftDataSpec", Required: true},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_create_data_source_from_s3 = []leanruntime.Field{
	{Name: "ComputeStatistics", Flag: "compute-statistics", Type: "bool", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: false},
	{Name: "DataSpec", Flag: "data-spec", Type: "*types.S3DataSpec", Required: true},
}

var fields_create_evaluation = []leanruntime.Field{
	{Name: "EvaluationDataSourceId", Flag: "evaluation-data-source-id", Type: "*string", Required: true},
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
	{Name: "EvaluationName", Flag: "evaluation-name", Type: "*string", Required: false},
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
}

var fields_create_ml_model = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
	{Name: "MLModelName", Flag: "ml-model-name", Type: "*string", Required: false},
	{Name: "MLModelType", Flag: "ml-model-type", Type: "types.MLModelType", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "Recipe", Flag: "recipe", Type: "*string", Required: false},
	{Name: "RecipeUri", Flag: "recipe-uri", Type: "*string", Required: false},
	{Name: "TrainingDataSourceId", Flag: "training-data-source-id", Type: "*string", Required: true},
}

var fields_create_realtime_endpoint = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
}

var fields_delete_batch_prediction = []leanruntime.Field{
	{Name: "BatchPredictionId", Flag: "batch-prediction-id", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
}

var fields_delete_evaluation = []leanruntime.Field{
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
}

var fields_delete_ml_model = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
}

var fields_delete_realtime_endpoint = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.TaggableResourceType", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_describe_batch_predictions = []leanruntime.Field{
	{Name: "EQ", Flag: "eq", Type: "*string", Required: false},
	{Name: "FilterVariable", Flag: "filter-variable", Type: "types.BatchPredictionFilterVariable", Required: false},
	{Name: "GE", Flag: "ge", Type: "*string", Required: false},
	{Name: "GT", Flag: "gt", Type: "*string", Required: false},
	{Name: "LE", Flag: "le", Type: "*string", Required: false},
	{Name: "LT", Flag: "lt", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NE", Flag: "ne", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_describe_data_sources = []leanruntime.Field{
	{Name: "EQ", Flag: "eq", Type: "*string", Required: false},
	{Name: "FilterVariable", Flag: "filter-variable", Type: "types.DataSourceFilterVariable", Required: false},
	{Name: "GE", Flag: "ge", Type: "*string", Required: false},
	{Name: "GT", Flag: "gt", Type: "*string", Required: false},
	{Name: "LE", Flag: "le", Type: "*string", Required: false},
	{Name: "LT", Flag: "lt", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NE", Flag: "ne", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_describe_evaluations = []leanruntime.Field{
	{Name: "EQ", Flag: "eq", Type: "*string", Required: false},
	{Name: "FilterVariable", Flag: "filter-variable", Type: "types.EvaluationFilterVariable", Required: false},
	{Name: "GE", Flag: "ge", Type: "*string", Required: false},
	{Name: "GT", Flag: "gt", Type: "*string", Required: false},
	{Name: "LE", Flag: "le", Type: "*string", Required: false},
	{Name: "LT", Flag: "lt", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NE", Flag: "ne", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_describe_ml_models = []leanruntime.Field{
	{Name: "EQ", Flag: "eq", Type: "*string", Required: false},
	{Name: "FilterVariable", Flag: "filter-variable", Type: "types.MLModelFilterVariable", Required: false},
	{Name: "GE", Flag: "ge", Type: "*string", Required: false},
	{Name: "GT", Flag: "gt", Type: "*string", Required: false},
	{Name: "LE", Flag: "le", Type: "*string", Required: false},
	{Name: "LT", Flag: "lt", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NE", Flag: "ne", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.TaggableResourceType", Required: true},
}

var fields_get_batch_prediction = []leanruntime.Field{
	{Name: "BatchPredictionId", Flag: "batch-prediction-id", Type: "*string", Required: true},
}

var fields_get_data_source = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "Verbose", Flag: "verbose", Type: "bool", Required: false},
}

var fields_get_evaluation = []leanruntime.Field{
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
}

var fields_get_ml_model = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
	{Name: "Verbose", Flag: "verbose", Type: "bool", Required: false},
}

var fields_predict = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
	{Name: "PredictEndpoint", Flag: "predict-endpoint", Type: "*string", Required: true},
	{Name: "Record", Flag: "record", Type: "map[string]string", Required: true},
}

var fields_update_batch_prediction = []leanruntime.Field{
	{Name: "BatchPredictionId", Flag: "batch-prediction-id", Type: "*string", Required: true},
	{Name: "BatchPredictionName", Flag: "batch-prediction-name", Type: "*string", Required: true},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
}

var fields_update_evaluation = []leanruntime.Field{
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
	{Name: "EvaluationName", Flag: "evaluation-name", Type: "*string", Required: true},
}

var fields_update_ml_model = []leanruntime.Field{
	{Name: "MLModelId", Flag: "ml-model-id", Type: "*string", Required: true},
	{Name: "MLModelName", Flag: "ml-model-name", Type: "*string", Required: false},
	{Name: "ScoreThreshold", Flag: "score-threshold", Type: "*float32", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"create-batch-prediction": {
			Name:   "create-batch-prediction",
			Fields: fields_create_batch_prediction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBatchPredictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_batch_prediction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBatchPrediction(ctx, input)
			},
		},
		"create-data-source-from-rds": {
			Name:   "create-data-source-from-rds",
			Fields: fields_create_data_source_from_rds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceFromRDSInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source_from_rds, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSourceFromRDS(ctx, input)
			},
		},
		"create-data-source-from-redshift": {
			Name:   "create-data-source-from-redshift",
			Fields: fields_create_data_source_from_redshift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceFromRedshiftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source_from_redshift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSourceFromRedshift(ctx, input)
			},
		},
		"create-data-source-from-s3": {
			Name:   "create-data-source-from-s3",
			Fields: fields_create_data_source_from_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceFromS3Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source_from_s3, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSourceFromS3(ctx, input)
			},
		},
		"create-evaluation": {
			Name:   "create-evaluation",
			Fields: fields_create_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEvaluation(ctx, input)
			},
		},
		"create-ml-model": {
			Name:   "create-ml-model",
			Fields: fields_create_ml_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMLModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ml_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMLModel(ctx, input)
			},
		},
		"create-realtime-endpoint": {
			Name:   "create-realtime-endpoint",
			Fields: fields_create_realtime_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRealtimeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_realtime_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRealtimeEndpoint(ctx, input)
			},
		},
		"delete-batch-prediction": {
			Name:   "delete-batch-prediction",
			Fields: fields_delete_batch_prediction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBatchPredictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_batch_prediction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBatchPrediction(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
			},
		},
		"delete-evaluation": {
			Name:   "delete-evaluation",
			Fields: fields_delete_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEvaluation(ctx, input)
			},
		},
		"delete-ml-model": {
			Name:   "delete-ml-model",
			Fields: fields_delete_ml_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMLModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ml_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMLModel(ctx, input)
			},
		},
		"delete-realtime-endpoint": {
			Name:   "delete-realtime-endpoint",
			Fields: fields_delete_realtime_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRealtimeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_realtime_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRealtimeEndpoint(ctx, input)
			},
		},
		"delete-tags": {
			Name:   "delete-tags",
			Fields: fields_delete_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTags(ctx, input)
			},
		},
		"describe-batch-predictions": {
			Name:   "describe-batch-predictions",
			Fields: fields_describe_batch_predictions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBatchPredictionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_batch_predictions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBatchPredictions(ctx, input)
				}
				var results []*svc.DescribeBatchPredictionsOutput
				p := svc.NewDescribeBatchPredictionsPaginator(client, input)
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
		"describe-data-sources": {
			Name:   "describe-data-sources",
			Fields: fields_describe_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataSources(ctx, input)
				}
				var results []*svc.DescribeDataSourcesOutput
				p := svc.NewDescribeDataSourcesPaginator(client, input)
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
		"describe-evaluations": {
			Name:   "describe-evaluations",
			Fields: fields_describe_evaluations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEvaluationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_evaluations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEvaluations(ctx, input)
				}
				var results []*svc.DescribeEvaluationsOutput
				p := svc.NewDescribeEvaluationsPaginator(client, input)
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
		"describe-ml-models": {
			Name:   "describe-ml-models",
			Fields: fields_describe_ml_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMLModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ml_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMLModels(ctx, input)
				}
				var results []*svc.DescribeMLModelsOutput
				p := svc.NewDescribeMLModelsPaginator(client, input)
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
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTags(ctx, input)
			},
		},
		"get-batch-prediction": {
			Name:   "get-batch-prediction",
			Fields: fields_get_batch_prediction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBatchPredictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_batch_prediction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBatchPrediction(ctx, input)
			},
		},
		"get-data-source": {
			Name:   "get-data-source",
			Fields: fields_get_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSource(ctx, input)
			},
		},
		"get-evaluation": {
			Name:   "get-evaluation",
			Fields: fields_get_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvaluation(ctx, input)
			},
		},
		"get-ml-model": {
			Name:   "get-ml-model",
			Fields: fields_get_ml_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLModel(ctx, input)
			},
		},
		"predict": {
			Name:   "predict",
			Fields: fields_predict,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PredictInput{}
				if _, err := leanruntime.ApplyInput(input, fields_predict, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Predict(ctx, input)
			},
		},
		"update-batch-prediction": {
			Name:   "update-batch-prediction",
			Fields: fields_update_batch_prediction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBatchPredictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_batch_prediction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBatchPrediction(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
			},
		},
		"update-evaluation": {
			Name:   "update-evaluation",
			Fields: fields_update_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEvaluation(ctx, input)
			},
		},
		"update-ml-model": {
			Name:   "update-ml-model",
			Fields: fields_update_ml_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMLModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ml_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMLModel(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("machinelearning", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
