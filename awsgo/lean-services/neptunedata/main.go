package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/neptunedata"
)

var fields_cancel_gremlin_query = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_cancel_loader_job = []leanruntime.Field{
	{Name: "LoadId", Flag: "load-id", Type: "*string", Required: true},
}

var fields_cancel_ml_data_processing_job = []leanruntime.Field{
	{Name: "Clean", Flag: "clean", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_cancel_ml_model_training_job = []leanruntime.Field{
	{Name: "Clean", Flag: "clean", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_cancel_ml_model_transform_job = []leanruntime.Field{
	{Name: "Clean", Flag: "clean", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_cancel_open_cypher_query = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "Silent", Flag: "silent", Type: "*bool", Required: false},
}

var fields_create_ml_endpoint = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "MlModelTrainingJobId", Flag: "ml-model-training-job-id", Type: "*string", Required: false},
	{Name: "MlModelTransformJobId", Flag: "ml-model-transform-job-id", Type: "*string", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
	{Name: "Update", Flag: "update", Type: "*bool", Required: false},
	{Name: "VolumeEncryptionKMSKey", Flag: "volume-encryption-kms-key", Type: "*string", Required: false},
}

var fields_delete_ml_endpoint = []leanruntime.Field{
	{Name: "Clean", Flag: "clean", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_delete_propertygraph_statistics = []leanruntime.Field{}

var fields_delete_sparql_statistics = []leanruntime.Field{}

var fields_execute_fast_reset = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.Action", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_execute_gremlin_explain_query = []leanruntime.Field{
	{Name: "GremlinQuery", Flag: "gremlin-query", Type: "*string", Required: true},
}

var fields_execute_gremlin_profile_query = []leanruntime.Field{
	{Name: "Chop", Flag: "chop", Type: "*int32", Required: false},
	{Name: "GremlinQuery", Flag: "gremlin-query", Type: "*string", Required: true},
	{Name: "IndexOps", Flag: "index-ops", Type: "*bool", Required: false},
	{Name: "Results", Flag: "results", Type: "*bool", Required: false},
	{Name: "Serializer", Flag: "serializer", Type: "*string", Required: false},
}

var fields_execute_gremlin_query = []leanruntime.Field{
	{Name: "GremlinQuery", Flag: "gremlin-query", Type: "*string", Required: true},
	{Name: "Serializer", Flag: "serializer", Type: "*string", Required: false},
}

var fields_execute_open_cypher_explain_query = []leanruntime.Field{
	{Name: "ExplainMode", Flag: "explain-mode", Type: "types.OpenCypherExplainMode", Required: true},
	{Name: "OpenCypherQuery", Flag: "open-cypher-query", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*string", Required: false},
}

var fields_execute_open_cypher_query = []leanruntime.Field{
	{Name: "OpenCypherQuery", Flag: "open-cypher-query", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*string", Required: false},
}

var fields_get_engine_status = []leanruntime.Field{}

var fields_get_gremlin_query_status = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_loader_job_status = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "*bool", Required: false},
	{Name: "Errors", Flag: "errors", Type: "*bool", Required: false},
	{Name: "ErrorsPerPage", Flag: "errors-per-page", Type: "*int32", Required: false},
	{Name: "LoadId", Flag: "load-id", Type: "*string", Required: true},
	{Name: "Page", Flag: "page", Type: "*int32", Required: false},
}

var fields_get_ml_data_processing_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_get_ml_endpoint = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_get_ml_model_training_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_get_ml_model_transform_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_get_open_cypher_query_status = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_propertygraph_statistics = []leanruntime.Field{}

var fields_get_propertygraph_stream = []leanruntime.Field{
	{Name: "CommitNum", Flag: "commit-num", Type: "*int64", Required: false},
	{Name: "Encoding", Flag: "encoding", Type: "types.Encoding", Required: false},
	{Name: "IteratorType", Flag: "iterator-type", Type: "types.IteratorType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int64", Required: false},
	{Name: "OpNum", Flag: "op-num", Type: "*int64", Required: false},
}

var fields_get_propertygraph_summary = []leanruntime.Field{
	{Name: "Mode", Flag: "mode", Type: "types.GraphSummaryType", Required: false},
}

var fields_get_rdf_graph_summary = []leanruntime.Field{
	{Name: "Mode", Flag: "mode", Type: "types.GraphSummaryType", Required: false},
}

var fields_get_sparql_statistics = []leanruntime.Field{}

var fields_get_sparql_stream = []leanruntime.Field{
	{Name: "CommitNum", Flag: "commit-num", Type: "*int64", Required: false},
	{Name: "Encoding", Flag: "encoding", Type: "types.Encoding", Required: false},
	{Name: "IteratorType", Flag: "iterator-type", Type: "types.IteratorType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int64", Required: false},
	{Name: "OpNum", Flag: "op-num", Type: "*int64", Required: false},
}

var fields_list_gremlin_queries = []leanruntime.Field{
	{Name: "IncludeWaiting", Flag: "include-waiting", Type: "*bool", Required: false},
}

var fields_list_loader_jobs = []leanruntime.Field{
	{Name: "IncludeQueuedLoads", Flag: "include-queued-loads", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
}

var fields_list_ml_data_processing_jobs = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_list_ml_endpoints = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_list_ml_model_training_jobs = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_list_ml_model_transform_jobs = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
}

var fields_list_open_cypher_queries = []leanruntime.Field{
	{Name: "IncludeWaiting", Flag: "include-waiting", Type: "*bool", Required: false},
}

var fields_manage_propertygraph_statistics = []leanruntime.Field{
	{Name: "Mode", Flag: "mode", Type: "types.StatisticsAutoGenerationMode", Required: false},
}

var fields_manage_sparql_statistics = []leanruntime.Field{
	{Name: "Mode", Flag: "mode", Type: "types.StatisticsAutoGenerationMode", Required: false},
}

var fields_start_loader_job = []leanruntime.Field{
	{Name: "Dependencies", Flag: "dependencies", Type: "[]string", Required: false},
	{Name: "EdgeOnlyLoad", Flag: "edge-only-load", Type: "*bool", Required: false},
	{Name: "FailOnError", Flag: "fail-on-error", Type: "*bool", Required: false},
	{Name: "Format", Flag: "format", Type: "types.Format", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.Mode", Required: false},
	{Name: "Parallelism", Flag: "parallelism", Type: "types.Parallelism", Required: false},
	{Name: "ParserConfiguration", Flag: "parser-configuration", Type: "map[string]string", Required: false},
	{Name: "QueueRequest", Flag: "queue-request", Type: "*bool", Required: false},
	{Name: "S3BucketRegion", Flag: "s3-bucket-region", Type: "types.S3BucketRegion", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "UpdateSingleCardinalityProperties", Flag: "update-single-cardinality-properties", Type: "*bool", Required: false},
	{Name: "UserProvidedEdgeIds", Flag: "user-provided-edge-ids", Type: "*bool", Required: false},
}

var fields_start_ml_data_processing_job = []leanruntime.Field{
	{Name: "ConfigFileName", Flag: "config-file-name", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "InputDataS3Location", Flag: "input-data-s3-location", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "*string", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
	{Name: "PreviousDataProcessingJobId", Flag: "previous-data-processing-job-id", Type: "*string", Required: false},
	{Name: "ProcessedDataS3Location", Flag: "processed-data-s3-location", Type: "*string", Required: true},
	{Name: "ProcessingInstanceType", Flag: "processing-instance-type", Type: "*string", Required: false},
	{Name: "ProcessingInstanceVolumeSizeInGB", Flag: "processing-instance-volume-size-in-gb", Type: "*int32", Required: false},
	{Name: "ProcessingTimeOutInSeconds", Flag: "processing-time-out-in-seconds", Type: "*int32", Required: false},
	{Name: "S3OutputEncryptionKMSKey", Flag: "s3-output-encryption-kms-key", Type: "*string", Required: false},
	{Name: "SagemakerIamRoleArn", Flag: "sagemaker-iam-role-arn", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
	{Name: "VolumeEncryptionKMSKey", Flag: "volume-encryption-kms-key", Type: "*string", Required: false},
}

var fields_start_ml_model_training_job = []leanruntime.Field{
	{Name: "BaseProcessingInstanceType", Flag: "base-processing-instance-type", Type: "*string", Required: false},
	{Name: "CustomModelTrainingParameters", Flag: "custom-model-training-parameters", Type: "*types.CustomModelTrainingParameters", Required: false},
	{Name: "DataProcessingJobId", Flag: "data-processing-job-id", Type: "*string", Required: true},
	{Name: "EnableManagedSpotTraining", Flag: "enable-managed-spot-training", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "MaxHPONumberOfTrainingJobs", Flag: "max-hpo-number-of-training-jobs", Type: "*int32", Required: false},
	{Name: "MaxHPOParallelTrainingJobs", Flag: "max-hpo-parallel-training-jobs", Type: "*int32", Required: false},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
	{Name: "PreviousModelTrainingJobId", Flag: "previous-model-training-job-id", Type: "*string", Required: false},
	{Name: "S3OutputEncryptionKMSKey", Flag: "s3-output-encryption-kms-key", Type: "*string", Required: false},
	{Name: "SagemakerIamRoleArn", Flag: "sagemaker-iam-role-arn", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
	{Name: "TrainModelS3Location", Flag: "train-model-s3-location", Type: "*string", Required: true},
	{Name: "TrainingInstanceType", Flag: "training-instance-type", Type: "*string", Required: false},
	{Name: "TrainingInstanceVolumeSizeInGB", Flag: "training-instance-volume-size-in-gb", Type: "*int32", Required: false},
	{Name: "TrainingTimeOutInSeconds", Flag: "training-time-out-in-seconds", Type: "*int32", Required: false},
	{Name: "VolumeEncryptionKMSKey", Flag: "volume-encryption-kms-key", Type: "*string", Required: false},
}

var fields_start_ml_model_transform_job = []leanruntime.Field{
	{Name: "BaseProcessingInstanceType", Flag: "base-processing-instance-type", Type: "*string", Required: false},
	{Name: "BaseProcessingInstanceVolumeSizeInGB", Flag: "base-processing-instance-volume-size-in-gb", Type: "*int32", Required: false},
	{Name: "CustomModelTransformParameters", Flag: "custom-model-transform-parameters", Type: "*types.CustomModelTransformParameters", Required: false},
	{Name: "DataProcessingJobId", Flag: "data-processing-job-id", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "MlModelTrainingJobId", Flag: "ml-model-training-job-id", Type: "*string", Required: false},
	{Name: "ModelTransformOutputS3Location", Flag: "model-transform-output-s3-location", Type: "*string", Required: true},
	{Name: "NeptuneIamRoleArn", Flag: "neptune-iam-role-arn", Type: "*string", Required: false},
	{Name: "S3OutputEncryptionKMSKey", Flag: "s3-output-encryption-kms-key", Type: "*string", Required: false},
	{Name: "SagemakerIamRoleArn", Flag: "sagemaker-iam-role-arn", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
	{Name: "TrainingJobName", Flag: "training-job-name", Type: "*string", Required: false},
	{Name: "VolumeEncryptionKMSKey", Flag: "volume-encryption-kms-key", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-gremlin-query": {
			Name:   "cancel-gremlin-query",
			Fields: fields_cancel_gremlin_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelGremlinQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_gremlin_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelGremlinQuery(ctx, input)
			},
		},
		"cancel-loader-job": {
			Name:   "cancel-loader-job",
			Fields: fields_cancel_loader_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelLoaderJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_loader_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelLoaderJob(ctx, input)
			},
		},
		"cancel-ml-data-processing-job": {
			Name:   "cancel-ml-data-processing-job",
			Fields: fields_cancel_ml_data_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMLDataProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_ml_data_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMLDataProcessingJob(ctx, input)
			},
		},
		"cancel-ml-model-training-job": {
			Name:   "cancel-ml-model-training-job",
			Fields: fields_cancel_ml_model_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMLModelTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_ml_model_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMLModelTrainingJob(ctx, input)
			},
		},
		"cancel-ml-model-transform-job": {
			Name:   "cancel-ml-model-transform-job",
			Fields: fields_cancel_ml_model_transform_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMLModelTransformJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_ml_model_transform_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMLModelTransformJob(ctx, input)
			},
		},
		"cancel-open-cypher-query": {
			Name:   "cancel-open-cypher-query",
			Fields: fields_cancel_open_cypher_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelOpenCypherQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_open_cypher_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelOpenCypherQuery(ctx, input)
			},
		},
		"create-ml-endpoint": {
			Name:   "create-ml-endpoint",
			Fields: fields_create_ml_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMLEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ml_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMLEndpoint(ctx, input)
			},
		},
		"delete-ml-endpoint": {
			Name:   "delete-ml-endpoint",
			Fields: fields_delete_ml_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMLEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ml_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMLEndpoint(ctx, input)
			},
		},
		"delete-propertygraph-statistics": {
			Name:   "delete-propertygraph-statistics",
			Fields: fields_delete_propertygraph_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePropertygraphStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_propertygraph_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePropertygraphStatistics(ctx, input)
			},
		},
		"delete-sparql-statistics": {
			Name:   "delete-sparql-statistics",
			Fields: fields_delete_sparql_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSparqlStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sparql_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSparqlStatistics(ctx, input)
			},
		},
		"execute-fast-reset": {
			Name:   "execute-fast-reset",
			Fields: fields_execute_fast_reset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteFastResetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_fast_reset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteFastReset(ctx, input)
			},
		},
		"execute-gremlin-explain-query": {
			Name:   "execute-gremlin-explain-query",
			Fields: fields_execute_gremlin_explain_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteGremlinExplainQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_gremlin_explain_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteGremlinExplainQuery(ctx, input)
			},
		},
		"execute-gremlin-profile-query": {
			Name:   "execute-gremlin-profile-query",
			Fields: fields_execute_gremlin_profile_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteGremlinProfileQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_gremlin_profile_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteGremlinProfileQuery(ctx, input)
			},
		},
		"execute-gremlin-query": {
			Name:   "execute-gremlin-query",
			Fields: fields_execute_gremlin_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteGremlinQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_gremlin_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteGremlinQuery(ctx, input)
			},
		},
		"execute-open-cypher-explain-query": {
			Name:   "execute-open-cypher-explain-query",
			Fields: fields_execute_open_cypher_explain_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteOpenCypherExplainQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_open_cypher_explain_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteOpenCypherExplainQuery(ctx, input)
			},
		},
		"execute-open-cypher-query": {
			Name:   "execute-open-cypher-query",
			Fields: fields_execute_open_cypher_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteOpenCypherQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_open_cypher_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteOpenCypherQuery(ctx, input)
			},
		},
		"get-engine-status": {
			Name:   "get-engine-status",
			Fields: fields_get_engine_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEngineStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_engine_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEngineStatus(ctx, input)
			},
		},
		"get-gremlin-query-status": {
			Name:   "get-gremlin-query-status",
			Fields: fields_get_gremlin_query_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGremlinQueryStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gremlin_query_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGremlinQueryStatus(ctx, input)
			},
		},
		"get-loader-job-status": {
			Name:   "get-loader-job-status",
			Fields: fields_get_loader_job_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoaderJobStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_loader_job_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoaderJobStatus(ctx, input)
			},
		},
		"get-ml-data-processing-job": {
			Name:   "get-ml-data-processing-job",
			Fields: fields_get_ml_data_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLDataProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_data_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLDataProcessingJob(ctx, input)
			},
		},
		"get-ml-endpoint": {
			Name:   "get-ml-endpoint",
			Fields: fields_get_ml_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLEndpoint(ctx, input)
			},
		},
		"get-ml-model-training-job": {
			Name:   "get-ml-model-training-job",
			Fields: fields_get_ml_model_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLModelTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_model_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLModelTrainingJob(ctx, input)
			},
		},
		"get-ml-model-transform-job": {
			Name:   "get-ml-model-transform-job",
			Fields: fields_get_ml_model_transform_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLModelTransformJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_model_transform_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLModelTransformJob(ctx, input)
			},
		},
		"get-open-cypher-query-status": {
			Name:   "get-open-cypher-query-status",
			Fields: fields_get_open_cypher_query_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpenCypherQueryStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_open_cypher_query_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpenCypherQueryStatus(ctx, input)
			},
		},
		"get-propertygraph-statistics": {
			Name:   "get-propertygraph-statistics",
			Fields: fields_get_propertygraph_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPropertygraphStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_propertygraph_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPropertygraphStatistics(ctx, input)
			},
		},
		"get-propertygraph-stream": {
			Name:   "get-propertygraph-stream",
			Fields: fields_get_propertygraph_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPropertygraphStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_propertygraph_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPropertygraphStream(ctx, input)
			},
		},
		"get-propertygraph-summary": {
			Name:   "get-propertygraph-summary",
			Fields: fields_get_propertygraph_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPropertygraphSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_propertygraph_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPropertygraphSummary(ctx, input)
			},
		},
		"get-rdf-graph-summary": {
			Name:   "get-rdf-graph-summary",
			Fields: fields_get_rdf_graph_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRDFGraphSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rdf_graph_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRDFGraphSummary(ctx, input)
			},
		},
		"get-sparql-statistics": {
			Name:   "get-sparql-statistics",
			Fields: fields_get_sparql_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSparqlStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sparql_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSparqlStatistics(ctx, input)
			},
		},
		"get-sparql-stream": {
			Name:   "get-sparql-stream",
			Fields: fields_get_sparql_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSparqlStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sparql_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSparqlStream(ctx, input)
			},
		},
		"list-gremlin-queries": {
			Name:   "list-gremlin-queries",
			Fields: fields_list_gremlin_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGremlinQueriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_gremlin_queries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGremlinQueries(ctx, input)
			},
		},
		"list-loader-jobs": {
			Name:   "list-loader-jobs",
			Fields: fields_list_loader_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLoaderJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_loader_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLoaderJobs(ctx, input)
			},
		},
		"list-ml-data-processing-jobs": {
			Name:   "list-ml-data-processing-jobs",
			Fields: fields_list_ml_data_processing_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMLDataProcessingJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_ml_data_processing_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMLDataProcessingJobs(ctx, input)
			},
		},
		"list-ml-endpoints": {
			Name:   "list-ml-endpoints",
			Fields: fields_list_ml_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMLEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_ml_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMLEndpoints(ctx, input)
			},
		},
		"list-ml-model-training-jobs": {
			Name:   "list-ml-model-training-jobs",
			Fields: fields_list_ml_model_training_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMLModelTrainingJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_ml_model_training_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMLModelTrainingJobs(ctx, input)
			},
		},
		"list-ml-model-transform-jobs": {
			Name:   "list-ml-model-transform-jobs",
			Fields: fields_list_ml_model_transform_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMLModelTransformJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_ml_model_transform_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMLModelTransformJobs(ctx, input)
			},
		},
		"list-open-cypher-queries": {
			Name:   "list-open-cypher-queries",
			Fields: fields_list_open_cypher_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpenCypherQueriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_open_cypher_queries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOpenCypherQueries(ctx, input)
			},
		},
		"manage-propertygraph-statistics": {
			Name:   "manage-propertygraph-statistics",
			Fields: fields_manage_propertygraph_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ManagePropertygraphStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_manage_propertygraph_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ManagePropertygraphStatistics(ctx, input)
			},
		},
		"manage-sparql-statistics": {
			Name:   "manage-sparql-statistics",
			Fields: fields_manage_sparql_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ManageSparqlStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_manage_sparql_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ManageSparqlStatistics(ctx, input)
			},
		},
		"start-loader-job": {
			Name:   "start-loader-job",
			Fields: fields_start_loader_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartLoaderJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_loader_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartLoaderJob(ctx, input)
			},
		},
		"start-ml-data-processing-job": {
			Name:   "start-ml-data-processing-job",
			Fields: fields_start_ml_data_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMLDataProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ml_data_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMLDataProcessingJob(ctx, input)
			},
		},
		"start-ml-model-training-job": {
			Name:   "start-ml-model-training-job",
			Fields: fields_start_ml_model_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMLModelTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ml_model_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMLModelTrainingJob(ctx, input)
			},
		},
		"start-ml-model-transform-job": {
			Name:   "start-ml-model-transform-job",
			Fields: fields_start_ml_model_transform_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMLModelTransformJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ml_model_transform_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMLModelTransformJob(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("neptunedata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
