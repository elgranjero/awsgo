package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cleanroomsml"
)

var fields_cancel_trained_model = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "VersionIdentifier", Flag: "version-identifier", Type: "*string", Required: false},
}

var fields_cancel_trained_model_inference_job = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "TrainedModelInferenceJobArn", Flag: "trained-model-inference-job-arn", Type: "*string", Required: true},
}

var fields_create_audience_model = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrainingDataEndTime", Flag: "training-data-end-time", Type: "*time.Time", Required: false},
	{Name: "TrainingDataStartTime", Flag: "training-data-start-time", Type: "*time.Time", Required: false},
	{Name: "TrainingDatasetArn", Flag: "training-dataset-arn", Type: "*string", Required: true},
}

var fields_create_configured_audience_model = []leanruntime.Field{
	{Name: "AudienceModelArn", Flag: "audience-model-arn", Type: "*string", Required: true},
	{Name: "AudienceSizeConfig", Flag: "audience-size-config", Type: "*types.AudienceSizeConfig", Required: false},
	{Name: "ChildResourceTagOnCreatePolicy", Flag: "child-resource-tag-on-create-policy", Type: "types.TagOnCreatePolicy", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MinMatchingSeedSize", Flag: "min-matching-seed-size", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.ConfiguredAudienceModelOutputConfig", Required: true},
	{Name: "SharedAudienceMetrics", Flag: "shared-audience-metrics", Type: "[]types.SharedAudienceMetrics", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configured_model_algorithm = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InferenceContainerConfig", Flag: "inference-container-config", Type: "*types.InferenceContainerConfig", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrainingContainerConfig", Flag: "training-container-config", Type: "*types.ContainerConfig", Required: false},
}

var fields_create_configured_model_algorithm_association = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmArn", Flag: "configured-model-algorithm-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PrivacyConfiguration", Flag: "privacy-configuration", Type: "*types.PrivacyConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_ml_input_channel = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmAssociations", Flag: "configured-model-algorithm-associations", Type: "[]string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InputChannel", Flag: "input-channel", Type: "*types.InputChannel", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RetentionInDays", Flag: "retention-in-days", Type: "*int32", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_trained_model = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmAssociationArn", Flag: "configured-model-algorithm-association-arn", Type: "*string", Required: true},
	{Name: "DataChannels", Flag: "data-channels", Type: "[]types.ModelTrainingDataChannel", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Environment", Flag: "environment", Type: "map[string]string", Required: false},
	{Name: "Hyperparameters", Flag: "hyperparameters", Type: "map[string]string", Required: false},
	{Name: "IncrementalTrainingDataChannels", Flag: "incremental-training-data-channels", Type: "[]types.IncrementalTrainingDataChannel", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.ResourceConfig", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.StoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrainingInputMode", Flag: "training-input-mode", Type: "types.TrainingInputMode", Required: false},
}

var fields_create_training_dataset = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrainingData", Flag: "training-data", Type: "[]types.Dataset", Required: true},
}

var fields_delete_audience_generation_job = []leanruntime.Field{
	{Name: "AudienceGenerationJobArn", Flag: "audience-generation-job-arn", Type: "*string", Required: true},
}

var fields_delete_audience_model = []leanruntime.Field{
	{Name: "AudienceModelArn", Flag: "audience-model-arn", Type: "*string", Required: true},
}

var fields_delete_configured_audience_model = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
}

var fields_delete_configured_audience_model_policy = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
}

var fields_delete_configured_model_algorithm = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmArn", Flag: "configured-model-algorithm-arn", Type: "*string", Required: true},
}

var fields_delete_configured_model_algorithm_association = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmAssociationArn", Flag: "configured-model-algorithm-association-arn", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_ml_configuration = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_ml_input_channel_data = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "MlInputChannelArn", Flag: "ml-input-channel-arn", Type: "*string", Required: true},
}

var fields_delete_trained_model_output = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "VersionIdentifier", Flag: "version-identifier", Type: "*string", Required: false},
}

var fields_delete_training_dataset = []leanruntime.Field{
	{Name: "TrainingDatasetArn", Flag: "training-dataset-arn", Type: "*string", Required: true},
}

var fields_get_audience_generation_job = []leanruntime.Field{
	{Name: "AudienceGenerationJobArn", Flag: "audience-generation-job-arn", Type: "*string", Required: true},
}

var fields_get_audience_model = []leanruntime.Field{
	{Name: "AudienceModelArn", Flag: "audience-model-arn", Type: "*string", Required: true},
}

var fields_get_collaboration_configured_model_algorithm_association = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "ConfiguredModelAlgorithmAssociationArn", Flag: "configured-model-algorithm-association-arn", Type: "*string", Required: true},
}

var fields_get_collaboration_ml_input_channel = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MlInputChannelArn", Flag: "ml-input-channel-arn", Type: "*string", Required: true},
}

var fields_get_collaboration_trained_model = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "VersionIdentifier", Flag: "version-identifier", Type: "*string", Required: false},
}

var fields_get_configured_audience_model = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
}

var fields_get_configured_audience_model_policy = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
}

var fields_get_configured_model_algorithm = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmArn", Flag: "configured-model-algorithm-arn", Type: "*string", Required: true},
}

var fields_get_configured_model_algorithm_association = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmAssociationArn", Flag: "configured-model-algorithm-association-arn", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_ml_configuration = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_ml_input_channel = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "MlInputChannelArn", Flag: "ml-input-channel-arn", Type: "*string", Required: true},
}

var fields_get_trained_model = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "VersionIdentifier", Flag: "version-identifier", Type: "*string", Required: false},
}

var fields_get_trained_model_inference_job = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "TrainedModelInferenceJobArn", Flag: "trained-model-inference-job-arn", Type: "*string", Required: true},
}

var fields_get_training_dataset = []leanruntime.Field{
	{Name: "TrainingDatasetArn", Flag: "training-dataset-arn", Type: "*string", Required: true},
}

var fields_list_audience_export_jobs = []leanruntime.Field{
	{Name: "AudienceGenerationJobArn", Flag: "audience-generation-job-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_audience_generation_jobs = []leanruntime.Field{
	{Name: "CollaborationId", Flag: "collaboration-id", Type: "*string", Required: false},
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_audience_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_configured_model_algorithm_associations = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_ml_input_channels = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_trained_model_export_jobs = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "TrainedModelVersionIdentifier", Flag: "trained-model-version-identifier", Type: "*string", Required: false},
}

var fields_list_collaboration_trained_model_inference_jobs = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: false},
	{Name: "TrainedModelVersionIdentifier", Flag: "trained-model-version-identifier", Type: "*string", Required: false},
}

var fields_list_collaboration_trained_models = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configured_audience_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configured_model_algorithm_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configured_model_algorithms = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ml_input_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_trained_model_inference_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: false},
	{Name: "TrainedModelVersionIdentifier", Flag: "trained-model-version-identifier", Type: "*string", Required: false},
}

var fields_list_trained_model_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TrainedModelStatus", Required: false},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
}

var fields_list_trained_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_training_datasets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_configured_audience_model_policy = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
	{Name: "ConfiguredAudienceModelPolicy", Flag: "configured-audience-model-policy", Type: "*string", Required: true},
	{Name: "PolicyExistenceCondition", Flag: "policy-existence-condition", Type: "types.PolicyExistenceCondition", Required: false},
	{Name: "PreviousPolicyHash", Flag: "previous-policy-hash", Type: "*string", Required: false},
}

var fields_put_ml_configuration = []leanruntime.Field{
	{Name: "DefaultOutputLocation", Flag: "default-output-location", Type: "*types.MLOutputConfiguration", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_start_audience_export_job = []leanruntime.Field{
	{Name: "AudienceGenerationJobArn", Flag: "audience-generation-job-arn", Type: "*string", Required: true},
	{Name: "AudienceSize", Flag: "audience-size", Type: "*types.AudienceSize", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_audience_generation_job = []leanruntime.Field{
	{Name: "CollaborationId", Flag: "collaboration-id", Type: "*string", Required: false},
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IncludeSeedInOutput", Flag: "include-seed-in-output", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SeedAudience", Flag: "seed-audience", Type: "*types.AudienceGenerationJobDataSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_trained_model_export_job = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputConfiguration", Flag: "output-configuration", Type: "*types.TrainedModelExportOutputConfiguration", Required: true},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "TrainedModelVersionIdentifier", Flag: "trained-model-version-identifier", Type: "*string", Required: false},
}

var fields_start_trained_model_inference_job = []leanruntime.Field{
	{Name: "ConfiguredModelAlgorithmAssociationArn", Flag: "configured-model-algorithm-association-arn", Type: "*string", Required: false},
	{Name: "ContainerExecutionParameters", Flag: "container-execution-parameters", Type: "*types.InferenceContainerExecutionParameters", Required: false},
	{Name: "DataSource", Flag: "data-source", Type: "*types.ModelInferenceDataSource", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Environment", Flag: "environment", Type: "map[string]string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputConfiguration", Flag: "output-configuration", Type: "*types.InferenceOutputConfiguration", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.InferenceResourceConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrainedModelArn", Flag: "trained-model-arn", Type: "*string", Required: true},
	{Name: "TrainedModelVersionIdentifier", Flag: "trained-model-version-identifier", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_configured_audience_model = []leanruntime.Field{
	{Name: "AudienceModelArn", Flag: "audience-model-arn", Type: "*string", Required: false},
	{Name: "AudienceSizeConfig", Flag: "audience-size-config", Type: "*types.AudienceSizeConfig", Required: false},
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MinMatchingSeedSize", Flag: "min-matching-seed-size", Type: "*int32", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.ConfiguredAudienceModelOutputConfig", Required: false},
	{Name: "SharedAudienceMetrics", Flag: "shared-audience-metrics", Type: "[]types.SharedAudienceMetrics", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-trained-model": {
			Name:   "cancel-trained-model",
			Fields: fields_cancel_trained_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTrainedModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_trained_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTrainedModel(ctx, input)
			},
		},
		"cancel-trained-model-inference-job": {
			Name:   "cancel-trained-model-inference-job",
			Fields: fields_cancel_trained_model_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTrainedModelInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_trained_model_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTrainedModelInferenceJob(ctx, input)
			},
		},
		"create-audience-model": {
			Name:   "create-audience-model",
			Fields: fields_create_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAudienceModel(ctx, input)
			},
		},
		"create-configured-audience-model": {
			Name:   "create-configured-audience-model",
			Fields: fields_create_configured_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredAudienceModel(ctx, input)
			},
		},
		"create-configured-model-algorithm": {
			Name:   "create-configured-model-algorithm",
			Fields: fields_create_configured_model_algorithm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredModelAlgorithmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_model_algorithm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredModelAlgorithm(ctx, input)
			},
		},
		"create-configured-model-algorithm-association": {
			Name:   "create-configured-model-algorithm-association",
			Fields: fields_create_configured_model_algorithm_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredModelAlgorithmAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_model_algorithm_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredModelAlgorithmAssociation(ctx, input)
			},
		},
		"create-ml-input-channel": {
			Name:   "create-ml-input-channel",
			Fields: fields_create_ml_input_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMLInputChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ml_input_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMLInputChannel(ctx, input)
			},
		},
		"create-trained-model": {
			Name:   "create-trained-model",
			Fields: fields_create_trained_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrainedModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trained_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrainedModel(ctx, input)
			},
		},
		"create-training-dataset": {
			Name:   "create-training-dataset",
			Fields: fields_create_training_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrainingDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_training_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrainingDataset(ctx, input)
			},
		},
		"delete-audience-generation-job": {
			Name:   "delete-audience-generation-job",
			Fields: fields_delete_audience_generation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAudienceGenerationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_audience_generation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAudienceGenerationJob(ctx, input)
			},
		},
		"delete-audience-model": {
			Name:   "delete-audience-model",
			Fields: fields_delete_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAudienceModel(ctx, input)
			},
		},
		"delete-configured-audience-model": {
			Name:   "delete-configured-audience-model",
			Fields: fields_delete_configured_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredAudienceModel(ctx, input)
			},
		},
		"delete-configured-audience-model-policy": {
			Name:   "delete-configured-audience-model-policy",
			Fields: fields_delete_configured_audience_model_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredAudienceModelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_audience_model_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredAudienceModelPolicy(ctx, input)
			},
		},
		"delete-configured-model-algorithm": {
			Name:   "delete-configured-model-algorithm",
			Fields: fields_delete_configured_model_algorithm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredModelAlgorithmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_model_algorithm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredModelAlgorithm(ctx, input)
			},
		},
		"delete-configured-model-algorithm-association": {
			Name:   "delete-configured-model-algorithm-association",
			Fields: fields_delete_configured_model_algorithm_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredModelAlgorithmAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_model_algorithm_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredModelAlgorithmAssociation(ctx, input)
			},
		},
		"delete-ml-configuration": {
			Name:   "delete-ml-configuration",
			Fields: fields_delete_ml_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMLConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ml_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMLConfiguration(ctx, input)
			},
		},
		"delete-ml-input-channel-data": {
			Name:   "delete-ml-input-channel-data",
			Fields: fields_delete_ml_input_channel_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMLInputChannelDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ml_input_channel_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMLInputChannelData(ctx, input)
			},
		},
		"delete-trained-model-output": {
			Name:   "delete-trained-model-output",
			Fields: fields_delete_trained_model_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrainedModelOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trained_model_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrainedModelOutput(ctx, input)
			},
		},
		"delete-training-dataset": {
			Name:   "delete-training-dataset",
			Fields: fields_delete_training_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrainingDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_training_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrainingDataset(ctx, input)
			},
		},
		"get-audience-generation-job": {
			Name:   "get-audience-generation-job",
			Fields: fields_get_audience_generation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAudienceGenerationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_audience_generation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAudienceGenerationJob(ctx, input)
			},
		},
		"get-audience-model": {
			Name:   "get-audience-model",
			Fields: fields_get_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAudienceModel(ctx, input)
			},
		},
		"get-collaboration-configured-model-algorithm-association": {
			Name:   "get-collaboration-configured-model-algorithm-association",
			Fields: fields_get_collaboration_configured_model_algorithm_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationConfiguredModelAlgorithmAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_configured_model_algorithm_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationConfiguredModelAlgorithmAssociation(ctx, input)
			},
		},
		"get-collaboration-ml-input-channel": {
			Name:   "get-collaboration-ml-input-channel",
			Fields: fields_get_collaboration_ml_input_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationMLInputChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_ml_input_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationMLInputChannel(ctx, input)
			},
		},
		"get-collaboration-trained-model": {
			Name:   "get-collaboration-trained-model",
			Fields: fields_get_collaboration_trained_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationTrainedModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_trained_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationTrainedModel(ctx, input)
			},
		},
		"get-configured-audience-model": {
			Name:   "get-configured-audience-model",
			Fields: fields_get_configured_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredAudienceModel(ctx, input)
			},
		},
		"get-configured-audience-model-policy": {
			Name:   "get-configured-audience-model-policy",
			Fields: fields_get_configured_audience_model_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredAudienceModelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_audience_model_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredAudienceModelPolicy(ctx, input)
			},
		},
		"get-configured-model-algorithm": {
			Name:   "get-configured-model-algorithm",
			Fields: fields_get_configured_model_algorithm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredModelAlgorithmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_model_algorithm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredModelAlgorithm(ctx, input)
			},
		},
		"get-configured-model-algorithm-association": {
			Name:   "get-configured-model-algorithm-association",
			Fields: fields_get_configured_model_algorithm_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredModelAlgorithmAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_model_algorithm_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredModelAlgorithmAssociation(ctx, input)
			},
		},
		"get-ml-configuration": {
			Name:   "get-ml-configuration",
			Fields: fields_get_ml_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLConfiguration(ctx, input)
			},
		},
		"get-ml-input-channel": {
			Name:   "get-ml-input-channel",
			Fields: fields_get_ml_input_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLInputChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_input_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLInputChannel(ctx, input)
			},
		},
		"get-trained-model": {
			Name:   "get-trained-model",
			Fields: fields_get_trained_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrainedModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trained_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrainedModel(ctx, input)
			},
		},
		"get-trained-model-inference-job": {
			Name:   "get-trained-model-inference-job",
			Fields: fields_get_trained_model_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrainedModelInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trained_model_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrainedModelInferenceJob(ctx, input)
			},
		},
		"get-training-dataset": {
			Name:   "get-training-dataset",
			Fields: fields_get_training_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrainingDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_training_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrainingDataset(ctx, input)
			},
		},
		"list-audience-export-jobs": {
			Name:   "list-audience-export-jobs",
			Fields: fields_list_audience_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAudienceExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audience_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAudienceExportJobs(ctx, input)
				}
				var results []*svc.ListAudienceExportJobsOutput
				p := svc.NewListAudienceExportJobsPaginator(client, input)
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
		"list-audience-generation-jobs": {
			Name:   "list-audience-generation-jobs",
			Fields: fields_list_audience_generation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAudienceGenerationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audience_generation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAudienceGenerationJobs(ctx, input)
				}
				var results []*svc.ListAudienceGenerationJobsOutput
				p := svc.NewListAudienceGenerationJobsPaginator(client, input)
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
		"list-audience-models": {
			Name:   "list-audience-models",
			Fields: fields_list_audience_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAudienceModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_audience_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAudienceModels(ctx, input)
				}
				var results []*svc.ListAudienceModelsOutput
				p := svc.NewListAudienceModelsPaginator(client, input)
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
		"list-collaboration-configured-model-algorithm-associations": {
			Name:   "list-collaboration-configured-model-algorithm-associations",
			Fields: fields_list_collaboration_configured_model_algorithm_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationConfiguredModelAlgorithmAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_configured_model_algorithm_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationConfiguredModelAlgorithmAssociations(ctx, input)
				}
				var results []*svc.ListCollaborationConfiguredModelAlgorithmAssociationsOutput
				p := svc.NewListCollaborationConfiguredModelAlgorithmAssociationsPaginator(client, input)
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
		"list-collaboration-ml-input-channels": {
			Name:   "list-collaboration-ml-input-channels",
			Fields: fields_list_collaboration_ml_input_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationMLInputChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_ml_input_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationMLInputChannels(ctx, input)
				}
				var results []*svc.ListCollaborationMLInputChannelsOutput
				p := svc.NewListCollaborationMLInputChannelsPaginator(client, input)
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
		"list-collaboration-trained-model-export-jobs": {
			Name:   "list-collaboration-trained-model-export-jobs",
			Fields: fields_list_collaboration_trained_model_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationTrainedModelExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_trained_model_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationTrainedModelExportJobs(ctx, input)
				}
				var results []*svc.ListCollaborationTrainedModelExportJobsOutput
				p := svc.NewListCollaborationTrainedModelExportJobsPaginator(client, input)
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
		"list-collaboration-trained-model-inference-jobs": {
			Name:   "list-collaboration-trained-model-inference-jobs",
			Fields: fields_list_collaboration_trained_model_inference_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationTrainedModelInferenceJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_trained_model_inference_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationTrainedModelInferenceJobs(ctx, input)
				}
				var results []*svc.ListCollaborationTrainedModelInferenceJobsOutput
				p := svc.NewListCollaborationTrainedModelInferenceJobsPaginator(client, input)
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
		"list-collaboration-trained-models": {
			Name:   "list-collaboration-trained-models",
			Fields: fields_list_collaboration_trained_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationTrainedModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_trained_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationTrainedModels(ctx, input)
				}
				var results []*svc.ListCollaborationTrainedModelsOutput
				p := svc.NewListCollaborationTrainedModelsPaginator(client, input)
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
		"list-configured-audience-models": {
			Name:   "list-configured-audience-models",
			Fields: fields_list_configured_audience_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfiguredAudienceModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configured_audience_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfiguredAudienceModels(ctx, input)
				}
				var results []*svc.ListConfiguredAudienceModelsOutput
				p := svc.NewListConfiguredAudienceModelsPaginator(client, input)
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
		"list-configured-model-algorithm-associations": {
			Name:   "list-configured-model-algorithm-associations",
			Fields: fields_list_configured_model_algorithm_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfiguredModelAlgorithmAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configured_model_algorithm_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfiguredModelAlgorithmAssociations(ctx, input)
				}
				var results []*svc.ListConfiguredModelAlgorithmAssociationsOutput
				p := svc.NewListConfiguredModelAlgorithmAssociationsPaginator(client, input)
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
		"list-configured-model-algorithms": {
			Name:   "list-configured-model-algorithms",
			Fields: fields_list_configured_model_algorithms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfiguredModelAlgorithmsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configured_model_algorithms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfiguredModelAlgorithms(ctx, input)
				}
				var results []*svc.ListConfiguredModelAlgorithmsOutput
				p := svc.NewListConfiguredModelAlgorithmsPaginator(client, input)
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
		"list-ml-input-channels": {
			Name:   "list-ml-input-channels",
			Fields: fields_list_ml_input_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMLInputChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ml_input_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMLInputChannels(ctx, input)
				}
				var results []*svc.ListMLInputChannelsOutput
				p := svc.NewListMLInputChannelsPaginator(client, input)
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
		"list-trained-model-inference-jobs": {
			Name:   "list-trained-model-inference-jobs",
			Fields: fields_list_trained_model_inference_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainedModelInferenceJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trained_model_inference_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainedModelInferenceJobs(ctx, input)
				}
				var results []*svc.ListTrainedModelInferenceJobsOutput
				p := svc.NewListTrainedModelInferenceJobsPaginator(client, input)
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
		"list-trained-model-versions": {
			Name:   "list-trained-model-versions",
			Fields: fields_list_trained_model_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainedModelVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trained_model_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainedModelVersions(ctx, input)
				}
				var results []*svc.ListTrainedModelVersionsOutput
				p := svc.NewListTrainedModelVersionsPaginator(client, input)
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
		"list-trained-models": {
			Name:   "list-trained-models",
			Fields: fields_list_trained_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainedModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trained_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainedModels(ctx, input)
				}
				var results []*svc.ListTrainedModelsOutput
				p := svc.NewListTrainedModelsPaginator(client, input)
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
		"list-training-datasets": {
			Name:   "list-training-datasets",
			Fields: fields_list_training_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainingDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_training_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainingDatasets(ctx, input)
				}
				var results []*svc.ListTrainingDatasetsOutput
				p := svc.NewListTrainingDatasetsPaginator(client, input)
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
		"put-configured-audience-model-policy": {
			Name:   "put-configured-audience-model-policy",
			Fields: fields_put_configured_audience_model_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfiguredAudienceModelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configured_audience_model_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfiguredAudienceModelPolicy(ctx, input)
			},
		},
		"put-ml-configuration": {
			Name:   "put-ml-configuration",
			Fields: fields_put_ml_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMLConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_ml_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMLConfiguration(ctx, input)
			},
		},
		"start-audience-export-job": {
			Name:   "start-audience-export-job",
			Fields: fields_start_audience_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAudienceExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_audience_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAudienceExportJob(ctx, input)
			},
		},
		"start-audience-generation-job": {
			Name:   "start-audience-generation-job",
			Fields: fields_start_audience_generation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAudienceGenerationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_audience_generation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAudienceGenerationJob(ctx, input)
			},
		},
		"start-trained-model-export-job": {
			Name:   "start-trained-model-export-job",
			Fields: fields_start_trained_model_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTrainedModelExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_trained_model_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTrainedModelExportJob(ctx, input)
			},
		},
		"start-trained-model-inference-job": {
			Name:   "start-trained-model-inference-job",
			Fields: fields_start_trained_model_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTrainedModelInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_trained_model_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTrainedModelInferenceJob(ctx, input)
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
		"update-configured-audience-model": {
			Name:   "update-configured-audience-model",
			Fields: fields_update_configured_audience_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfiguredAudienceModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configured_audience_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguredAudienceModel(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cleanroomsml", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
