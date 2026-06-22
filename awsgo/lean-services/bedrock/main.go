package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrock"
)

var fields_batch_delete_evaluation_job = []leanruntime.Field{
	{Name: "JobIdentifiers", Flag: "job-identifiers", Type: "[]string", Required: true},
}

var fields_cancel_automated_reasoning_policy_build_workflow = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_create_automated_reasoning_policy = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyDefinition", Flag: "policy-definition", Type: "*types.AutomatedReasoningPolicyDefinition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_automated_reasoning_policy_test_case = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ConfidenceThreshold", Flag: "confidence-threshold", Type: "*float64", Required: false},
	{Name: "ExpectedAggregatedFindingsResult", Flag: "expected-aggregated-findings-result", Type: "types.AutomatedReasoningCheckResult", Required: true},
	{Name: "GuardContent", Flag: "guard-content", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "QueryContent", Flag: "query-content", Type: "*string", Required: false},
}

var fields_create_automated_reasoning_policy_version = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LastUpdatedDefinitionHash", Flag: "last-updated-definition-hash", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_model = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ModelKmsKeyArn", Flag: "model-kms-key-arn", Type: "*string", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "ModelSourceConfig", Flag: "model-source-config", Type: "types.ModelDataSource", Required: true},
	{Name: "ModelTags", Flag: "model-tags", Type: "[]types.Tag", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_create_custom_model_deployment = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModelArn", Flag: "model-arn", Type: "*string", Required: true},
	{Name: "ModelDeploymentName", Flag: "model-deployment-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_evaluation_job = []leanruntime.Field{
	{Name: "ApplicationType", Flag: "application-type", Type: "types.ApplicationType", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CustomerEncryptionKeyId", Flag: "customer-encryption-key-id", Type: "*string", Required: false},
	{Name: "EvaluationConfig", Flag: "evaluation-config", Type: "types.EvaluationConfig", Required: true},
	{Name: "InferenceConfig", Flag: "inference-config", Type: "types.EvaluationInferenceConfig", Required: true},
	{Name: "JobDescription", Flag: "job-description", Type: "*string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobTags", Flag: "job-tags", Type: "[]types.Tag", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.EvaluationOutputDataConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_create_foundation_model_agreement = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "OfferToken", Flag: "offer-token", Type: "*string", Required: true},
}

var fields_create_guardrail = []leanruntime.Field{
	{Name: "AutomatedReasoningPolicyConfig", Flag: "automated-reasoning-policy-config", Type: "*types.GuardrailAutomatedReasoningPolicyConfig", Required: false},
	{Name: "BlockedInputMessaging", Flag: "blocked-input-messaging", Type: "*string", Required: true},
	{Name: "BlockedOutputsMessaging", Flag: "blocked-outputs-messaging", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ContentPolicyConfig", Flag: "content-policy-config", Type: "*types.GuardrailContentPolicyConfig", Required: false},
	{Name: "ContextualGroundingPolicyConfig", Flag: "contextual-grounding-policy-config", Type: "*types.GuardrailContextualGroundingPolicyConfig", Required: false},
	{Name: "CrossRegionConfig", Flag: "cross-region-config", Type: "*types.GuardrailCrossRegionConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SensitiveInformationPolicyConfig", Flag: "sensitive-information-policy-config", Type: "*types.GuardrailSensitiveInformationPolicyConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TopicPolicyConfig", Flag: "topic-policy-config", Type: "*types.GuardrailTopicPolicyConfig", Required: false},
	{Name: "WordPolicyConfig", Flag: "word-policy-config", Type: "*types.GuardrailWordPolicyConfig", Required: false},
}

var fields_create_guardrail_version = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: true},
}

var fields_create_inference_profile = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InferenceProfileName", Flag: "inference-profile-name", Type: "*string", Required: true},
	{Name: "ModelSource", Flag: "model-source", Type: "types.InferenceProfileModelSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_marketplace_model_endpoint = []leanruntime.Field{
	{Name: "AcceptEula", Flag: "accept-eula", Type: "bool", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "EndpointConfig", Flag: "endpoint-config", Type: "types.EndpointConfig", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "ModelSourceIdentifier", Flag: "model-source-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_copy_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ModelKmsKeyId", Flag: "model-kms-key-id", Type: "*string", Required: false},
	{Name: "SourceModelArn", Flag: "source-model-arn", Type: "*string", Required: true},
	{Name: "TargetModelName", Flag: "target-model-name", Type: "*string", Required: true},
	{Name: "TargetModelTags", Flag: "target-model-tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_customization_job = []leanruntime.Field{
	{Name: "BaseModelIdentifier", Flag: "base-model-identifier", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CustomModelKmsKeyId", Flag: "custom-model-kms-key-id", Type: "*string", Required: false},
	{Name: "CustomModelName", Flag: "custom-model-name", Type: "*string", Required: true},
	{Name: "CustomModelTags", Flag: "custom-model-tags", Type: "[]types.Tag", Required: false},
	{Name: "CustomizationConfig", Flag: "customization-config", Type: "types.CustomizationConfig", Required: false},
	{Name: "CustomizationType", Flag: "customization-type", Type: "types.CustomizationType", Required: false},
	{Name: "HyperParameters", Flag: "hyper-parameters", Type: "map[string]string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobTags", Flag: "job-tags", Type: "[]types.Tag", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "TrainingDataConfig", Flag: "training-data-config", Type: "*types.TrainingDataConfig", Required: true},
	{Name: "ValidationDataConfig", Flag: "validation-data-config", Type: "*types.ValidationDataConfig", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_model_import_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ImportedModelKmsKeyId", Flag: "imported-model-kms-key-id", Type: "*string", Required: false},
	{Name: "ImportedModelName", Flag: "imported-model-name", Type: "*string", Required: true},
	{Name: "ImportedModelTags", Flag: "imported-model-tags", Type: "[]types.Tag", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobTags", Flag: "job-tags", Type: "[]types.Tag", Required: false},
	{Name: "ModelDataSource", Flag: "model-data-source", Type: "types.ModelDataSource", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_model_invocation_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "types.ModelInvocationJobInputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelInvocationType", Flag: "model-invocation-type", Type: "types.ModelInvocationType", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "types.ModelInvocationJobOutputDataConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutDurationInHours", Flag: "timeout-duration-in-hours", Type: "*int32", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_prompt_router = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FallbackModel", Flag: "fallback-model", Type: "*types.PromptRouterTargetModel", Required: true},
	{Name: "Models", Flag: "models", Type: "[]types.PromptRouterTargetModel", Required: true},
	{Name: "PromptRouterName", Flag: "prompt-router-name", Type: "*string", Required: true},
	{Name: "RoutingCriteria", Flag: "routing-criteria", Type: "*types.RoutingCriteria", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_provisioned_model_throughput = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CommitmentDuration", Flag: "commitment-duration", Type: "types.CommitmentDuration", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelUnits", Flag: "model-units", Type: "*int32", Required: true},
	{Name: "ProvisionedModelName", Flag: "provisioned-model-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_automated_reasoning_policy = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_delete_automated_reasoning_policy_build_workflow = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "LastUpdatedAt", Flag: "last-updated-at", Type: "*time.Time", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_delete_automated_reasoning_policy_test_case = []leanruntime.Field{
	{Name: "LastUpdatedAt", Flag: "last-updated-at", Type: "*time.Time", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_delete_custom_model = []leanruntime.Field{
	{Name: "ModelIdentifier", Flag: "model-identifier", Type: "*string", Required: true},
}

var fields_delete_custom_model_deployment = []leanruntime.Field{
	{Name: "CustomModelDeploymentIdentifier", Flag: "custom-model-deployment-identifier", Type: "*string", Required: true},
}

var fields_delete_enforced_guardrail_configuration = []leanruntime.Field{
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
}

var fields_delete_foundation_model_agreement = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_delete_guardrail = []leanruntime.Field{
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: true},
	{Name: "GuardrailVersion", Flag: "guardrail-version", Type: "*string", Required: false},
}

var fields_delete_imported_model = []leanruntime.Field{
	{Name: "ModelIdentifier", Flag: "model-identifier", Type: "*string", Required: true},
}

var fields_delete_inference_profile = []leanruntime.Field{
	{Name: "InferenceProfileIdentifier", Flag: "inference-profile-identifier", Type: "*string", Required: true},
}

var fields_delete_marketplace_model_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_delete_model_invocation_logging_configuration = []leanruntime.Field{}

var fields_delete_prompt_router = []leanruntime.Field{
	{Name: "PromptRouterArn", Flag: "prompt-router-arn", Type: "*string", Required: true},
}

var fields_delete_provisioned_model_throughput = []leanruntime.Field{
	{Name: "ProvisionedModelId", Flag: "provisioned-model-id", Type: "*string", Required: true},
}

var fields_deregister_marketplace_model_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_export_automated_reasoning_policy_version = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy_annotations = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy_build_workflow = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy_build_workflow_result_assets = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "AssetType", Flag: "asset-type", Type: "types.AutomatedReasoningPolicyBuildResultAssetType", Required: true},
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy_next_scenario = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy_test_case = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_get_automated_reasoning_policy_test_result = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_get_custom_model = []leanruntime.Field{
	{Name: "ModelIdentifier", Flag: "model-identifier", Type: "*string", Required: true},
}

var fields_get_custom_model_deployment = []leanruntime.Field{
	{Name: "CustomModelDeploymentIdentifier", Flag: "custom-model-deployment-identifier", Type: "*string", Required: true},
}

var fields_get_evaluation_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_get_foundation_model = []leanruntime.Field{
	{Name: "ModelIdentifier", Flag: "model-identifier", Type: "*string", Required: true},
}

var fields_get_foundation_model_availability = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_get_guardrail = []leanruntime.Field{
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: true},
	{Name: "GuardrailVersion", Flag: "guardrail-version", Type: "*string", Required: false},
}

var fields_get_imported_model = []leanruntime.Field{
	{Name: "ModelIdentifier", Flag: "model-identifier", Type: "*string", Required: true},
}

var fields_get_inference_profile = []leanruntime.Field{
	{Name: "InferenceProfileIdentifier", Flag: "inference-profile-identifier", Type: "*string", Required: true},
}

var fields_get_marketplace_model_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_get_model_copy_job = []leanruntime.Field{
	{Name: "JobArn", Flag: "job-arn", Type: "*string", Required: true},
}

var fields_get_model_customization_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_get_model_import_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_get_model_invocation_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_get_model_invocation_logging_configuration = []leanruntime.Field{}

var fields_get_prompt_router = []leanruntime.Field{
	{Name: "PromptRouterArn", Flag: "prompt-router-arn", Type: "*string", Required: true},
}

var fields_get_provisioned_model_throughput = []leanruntime.Field{
	{Name: "ProvisionedModelId", Flag: "provisioned-model-id", Type: "*string", Required: true},
}

var fields_get_use_case_for_model_access = []leanruntime.Field{}

var fields_list_automated_reasoning_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: false},
}

var fields_list_automated_reasoning_policy_build_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_list_automated_reasoning_policy_test_cases = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_list_automated_reasoning_policy_test_results = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_list_custom_model_deployments = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelArnEquals", Flag: "model-arn-equals", Type: "*string", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortModelsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.CustomModelDeploymentStatus", Required: false},
}

var fields_list_custom_models = []leanruntime.Field{
	{Name: "BaseModelArnEquals", Flag: "base-model-arn-equals", Type: "*string", Required: false},
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "FoundationModelArnEquals", Flag: "foundation-model-arn-equals", Type: "*string", Required: false},
	{Name: "IsOwned", Flag: "is-owned", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelStatus", Flag: "model-status", Type: "types.ModelStatus", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortModelsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_enforced_guardrails_configuration = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_evaluation_jobs = []leanruntime.Field{
	{Name: "ApplicationTypeEquals", Flag: "application-type-equals", Type: "types.ApplicationType", Required: false},
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortJobsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.EvaluationJobStatus", Required: false},
}

var fields_list_foundation_model_agreement_offers = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "OfferType", Flag: "offer-type", Type: "types.OfferType", Required: false},
}

var fields_list_foundation_models = []leanruntime.Field{
	{Name: "ByCustomizationType", Flag: "by-customization-type", Type: "types.ModelCustomization", Required: false},
	{Name: "ByInferenceType", Flag: "by-inference-type", Type: "types.InferenceType", Required: false},
	{Name: "ByOutputModality", Flag: "by-output-modality", Type: "types.ModelModality", Required: false},
	{Name: "ByProvider", Flag: "by-provider", Type: "*string", Required: false},
}

var fields_list_guardrails = []leanruntime.Field{
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_imported_models = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortModelsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_inference_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TypeEquals", Flag: "type-equals", Type: "types.InferenceProfileType", Required: false},
}

var fields_list_marketplace_model_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelSourceEquals", Flag: "model-source-equals", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_model_copy_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortJobsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SourceAccountEquals", Flag: "source-account-equals", Type: "*string", Required: false},
	{Name: "SourceModelArnEquals", Flag: "source-model-arn-equals", Type: "*string", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ModelCopyJobStatus", Required: false},
	{Name: "TargetModelNameContains", Flag: "target-model-name-contains", Type: "*string", Required: false},
}

var fields_list_model_customization_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortJobsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.FineTuningJobStatus", Required: false},
}

var fields_list_model_import_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortJobsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ModelImportJobStatus", Required: false},
}

var fields_list_model_invocation_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortJobsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ModelInvocationJobStatus", Required: false},
	{Name: "SubmitTimeAfter", Flag: "submit-time-after", Type: "*time.Time", Required: false},
	{Name: "SubmitTimeBefore", Flag: "submit-time-before", Type: "*time.Time", Required: false},
}

var fields_list_prompt_routers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.PromptRouterType", Required: false},
}

var fields_list_provisioned_model_throughputs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelArnEquals", Flag: "model-arn-equals", Type: "*string", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortByProvisionedModels", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ProvisionedModelStatus", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_enforced_guardrail_configuration = []leanruntime.Field{
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: false},
	{Name: "GuardrailInferenceConfig", Flag: "guardrail-inference-config", Type: "*types.AccountEnforcedGuardrailInferenceInputConfiguration", Required: true},
}

var fields_put_model_invocation_logging_configuration = []leanruntime.Field{
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.LoggingConfig", Required: true},
}

var fields_put_use_case_for_model_access = []leanruntime.Field{
	{Name: "FormData", Flag: "form-data", Type: "[]byte", Required: true},
}

var fields_register_marketplace_model_endpoint = []leanruntime.Field{
	{Name: "EndpointIdentifier", Flag: "endpoint-identifier", Type: "*string", Required: true},
	{Name: "ModelSourceIdentifier", Flag: "model-source-identifier", Type: "*string", Required: true},
}

var fields_start_automated_reasoning_policy_build_workflow = []leanruntime.Field{
	{Name: "BuildWorkflowType", Flag: "build-workflow-type", Type: "types.AutomatedReasoningPolicyBuildWorkflowType", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "SourceContent", Flag: "source-content", Type: "*types.AutomatedReasoningPolicyBuildWorkflowSource", Required: true},
}

var fields_start_automated_reasoning_policy_test_workflow = []leanruntime.Field{
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "TestCaseIds", Flag: "test-case-ids", Type: "[]string", Required: false},
}

var fields_stop_evaluation_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_stop_model_customization_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_stop_model_invocation_job = []leanruntime.Field{
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_automated_reasoning_policy = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "PolicyDefinition", Flag: "policy-definition", Type: "*types.AutomatedReasoningPolicyDefinition", Required: true},
}

var fields_update_automated_reasoning_policy_annotations = []leanruntime.Field{
	{Name: "Annotations", Flag: "annotations", Type: "[]types.AutomatedReasoningPolicyAnnotation", Required: true},
	{Name: "BuildWorkflowId", Flag: "build-workflow-id", Type: "*string", Required: true},
	{Name: "LastUpdatedAnnotationSetHash", Flag: "last-updated-annotation-set-hash", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_update_automated_reasoning_policy_test_case = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ConfidenceThreshold", Flag: "confidence-threshold", Type: "*float64", Required: false},
	{Name: "ExpectedAggregatedFindingsResult", Flag: "expected-aggregated-findings-result", Type: "types.AutomatedReasoningCheckResult", Required: true},
	{Name: "GuardContent", Flag: "guard-content", Type: "*string", Required: true},
	{Name: "LastUpdatedAt", Flag: "last-updated-at", Type: "*time.Time", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "QueryContent", Flag: "query-content", Type: "*string", Required: false},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_update_custom_model_deployment = []leanruntime.Field{
	{Name: "CustomModelDeploymentIdentifier", Flag: "custom-model-deployment-identifier", Type: "*string", Required: true},
	{Name: "ModelArn", Flag: "model-arn", Type: "*string", Required: true},
}

var fields_update_guardrail = []leanruntime.Field{
	{Name: "AutomatedReasoningPolicyConfig", Flag: "automated-reasoning-policy-config", Type: "*types.GuardrailAutomatedReasoningPolicyConfig", Required: false},
	{Name: "BlockedInputMessaging", Flag: "blocked-input-messaging", Type: "*string", Required: true},
	{Name: "BlockedOutputsMessaging", Flag: "blocked-outputs-messaging", Type: "*string", Required: true},
	{Name: "ContentPolicyConfig", Flag: "content-policy-config", Type: "*types.GuardrailContentPolicyConfig", Required: false},
	{Name: "ContextualGroundingPolicyConfig", Flag: "contextual-grounding-policy-config", Type: "*types.GuardrailContextualGroundingPolicyConfig", Required: false},
	{Name: "CrossRegionConfig", Flag: "cross-region-config", Type: "*types.GuardrailCrossRegionConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SensitiveInformationPolicyConfig", Flag: "sensitive-information-policy-config", Type: "*types.GuardrailSensitiveInformationPolicyConfig", Required: false},
	{Name: "TopicPolicyConfig", Flag: "topic-policy-config", Type: "*types.GuardrailTopicPolicyConfig", Required: false},
	{Name: "WordPolicyConfig", Flag: "word-policy-config", Type: "*types.GuardrailWordPolicyConfig", Required: false},
}

var fields_update_marketplace_model_endpoint = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "EndpointConfig", Flag: "endpoint-config", Type: "types.EndpointConfig", Required: true},
}

var fields_update_provisioned_model_throughput = []leanruntime.Field{
	{Name: "DesiredModelId", Flag: "desired-model-id", Type: "*string", Required: false},
	{Name: "DesiredProvisionedModelName", Flag: "desired-provisioned-model-name", Type: "*string", Required: false},
	{Name: "ProvisionedModelId", Flag: "provisioned-model-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-delete-evaluation-job": {
			Name:   "batch-delete-evaluation-job",
			Fields: fields_batch_delete_evaluation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteEvaluationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_evaluation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteEvaluationJob(ctx, input)
			},
		},
		"cancel-automated-reasoning-policy-build-workflow": {
			Name:   "cancel-automated-reasoning-policy-build-workflow",
			Fields: fields_cancel_automated_reasoning_policy_build_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelAutomatedReasoningPolicyBuildWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_automated_reasoning_policy_build_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelAutomatedReasoningPolicyBuildWorkflow(ctx, input)
			},
		},
		"create-automated-reasoning-policy": {
			Name:   "create-automated-reasoning-policy",
			Fields: fields_create_automated_reasoning_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutomatedReasoningPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_automated_reasoning_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutomatedReasoningPolicy(ctx, input)
			},
		},
		"create-automated-reasoning-policy-test-case": {
			Name:   "create-automated-reasoning-policy-test-case",
			Fields: fields_create_automated_reasoning_policy_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutomatedReasoningPolicyTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_automated_reasoning_policy_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutomatedReasoningPolicyTestCase(ctx, input)
			},
		},
		"create-automated-reasoning-policy-version": {
			Name:   "create-automated-reasoning-policy-version",
			Fields: fields_create_automated_reasoning_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutomatedReasoningPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_automated_reasoning_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutomatedReasoningPolicyVersion(ctx, input)
			},
		},
		"create-custom-model": {
			Name:   "create-custom-model",
			Fields: fields_create_custom_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomModel(ctx, input)
			},
		},
		"create-custom-model-deployment": {
			Name:   "create-custom-model-deployment",
			Fields: fields_create_custom_model_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomModelDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_model_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomModelDeployment(ctx, input)
			},
		},
		"create-evaluation-job": {
			Name:   "create-evaluation-job",
			Fields: fields_create_evaluation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEvaluationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_evaluation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEvaluationJob(ctx, input)
			},
		},
		"create-foundation-model-agreement": {
			Name:   "create-foundation-model-agreement",
			Fields: fields_create_foundation_model_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFoundationModelAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_foundation_model_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFoundationModelAgreement(ctx, input)
			},
		},
		"create-guardrail": {
			Name:   "create-guardrail",
			Fields: fields_create_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGuardrail(ctx, input)
			},
		},
		"create-guardrail-version": {
			Name:   "create-guardrail-version",
			Fields: fields_create_guardrail_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGuardrailVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_guardrail_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGuardrailVersion(ctx, input)
			},
		},
		"create-inference-profile": {
			Name:   "create-inference-profile",
			Fields: fields_create_inference_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInferenceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_inference_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInferenceProfile(ctx, input)
			},
		},
		"create-marketplace-model-endpoint": {
			Name:   "create-marketplace-model-endpoint",
			Fields: fields_create_marketplace_model_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMarketplaceModelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_marketplace_model_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMarketplaceModelEndpoint(ctx, input)
			},
		},
		"create-model-copy-job": {
			Name:   "create-model-copy-job",
			Fields: fields_create_model_copy_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelCopyJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_copy_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelCopyJob(ctx, input)
			},
		},
		"create-model-customization-job": {
			Name:   "create-model-customization-job",
			Fields: fields_create_model_customization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelCustomizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_customization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelCustomizationJob(ctx, input)
			},
		},
		"create-model-import-job": {
			Name:   "create-model-import-job",
			Fields: fields_create_model_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelImportJob(ctx, input)
			},
		},
		"create-model-invocation-job": {
			Name:   "create-model-invocation-job",
			Fields: fields_create_model_invocation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelInvocationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_invocation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelInvocationJob(ctx, input)
			},
		},
		"create-prompt-router": {
			Name:   "create-prompt-router",
			Fields: fields_create_prompt_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePromptRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_prompt_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePromptRouter(ctx, input)
			},
		},
		"create-provisioned-model-throughput": {
			Name:   "create-provisioned-model-throughput",
			Fields: fields_create_provisioned_model_throughput,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisionedModelThroughputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioned_model_throughput, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisionedModelThroughput(ctx, input)
			},
		},
		"delete-automated-reasoning-policy": {
			Name:   "delete-automated-reasoning-policy",
			Fields: fields_delete_automated_reasoning_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutomatedReasoningPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_automated_reasoning_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutomatedReasoningPolicy(ctx, input)
			},
		},
		"delete-automated-reasoning-policy-build-workflow": {
			Name:   "delete-automated-reasoning-policy-build-workflow",
			Fields: fields_delete_automated_reasoning_policy_build_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutomatedReasoningPolicyBuildWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_automated_reasoning_policy_build_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutomatedReasoningPolicyBuildWorkflow(ctx, input)
			},
		},
		"delete-automated-reasoning-policy-test-case": {
			Name:   "delete-automated-reasoning-policy-test-case",
			Fields: fields_delete_automated_reasoning_policy_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutomatedReasoningPolicyTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_automated_reasoning_policy_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutomatedReasoningPolicyTestCase(ctx, input)
			},
		},
		"delete-custom-model": {
			Name:   "delete-custom-model",
			Fields: fields_delete_custom_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomModel(ctx, input)
			},
		},
		"delete-custom-model-deployment": {
			Name:   "delete-custom-model-deployment",
			Fields: fields_delete_custom_model_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomModelDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_model_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomModelDeployment(ctx, input)
			},
		},
		"delete-enforced-guardrail-configuration": {
			Name:   "delete-enforced-guardrail-configuration",
			Fields: fields_delete_enforced_guardrail_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnforcedGuardrailConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_enforced_guardrail_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnforcedGuardrailConfiguration(ctx, input)
			},
		},
		"delete-foundation-model-agreement": {
			Name:   "delete-foundation-model-agreement",
			Fields: fields_delete_foundation_model_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFoundationModelAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_foundation_model_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFoundationModelAgreement(ctx, input)
			},
		},
		"delete-guardrail": {
			Name:   "delete-guardrail",
			Fields: fields_delete_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGuardrail(ctx, input)
			},
		},
		"delete-imported-model": {
			Name:   "delete-imported-model",
			Fields: fields_delete_imported_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImportedModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_imported_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImportedModel(ctx, input)
			},
		},
		"delete-inference-profile": {
			Name:   "delete-inference-profile",
			Fields: fields_delete_inference_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInferenceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inference_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInferenceProfile(ctx, input)
			},
		},
		"delete-marketplace-model-endpoint": {
			Name:   "delete-marketplace-model-endpoint",
			Fields: fields_delete_marketplace_model_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMarketplaceModelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_marketplace_model_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMarketplaceModelEndpoint(ctx, input)
			},
		},
		"delete-model-invocation-logging-configuration": {
			Name:   "delete-model-invocation-logging-configuration",
			Fields: fields_delete_model_invocation_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelInvocationLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_invocation_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelInvocationLoggingConfiguration(ctx, input)
			},
		},
		"delete-prompt-router": {
			Name:   "delete-prompt-router",
			Fields: fields_delete_prompt_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePromptRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_prompt_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePromptRouter(ctx, input)
			},
		},
		"delete-provisioned-model-throughput": {
			Name:   "delete-provisioned-model-throughput",
			Fields: fields_delete_provisioned_model_throughput,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisionedModelThroughputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioned_model_throughput, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisionedModelThroughput(ctx, input)
			},
		},
		"deregister-marketplace-model-endpoint": {
			Name:   "deregister-marketplace-model-endpoint",
			Fields: fields_deregister_marketplace_model_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterMarketplaceModelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_marketplace_model_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterMarketplaceModelEndpoint(ctx, input)
			},
		},
		"export-automated-reasoning-policy-version": {
			Name:   "export-automated-reasoning-policy-version",
			Fields: fields_export_automated_reasoning_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportAutomatedReasoningPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_automated_reasoning_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportAutomatedReasoningPolicyVersion(ctx, input)
			},
		},
		"get-automated-reasoning-policy": {
			Name:   "get-automated-reasoning-policy",
			Fields: fields_get_automated_reasoning_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicy(ctx, input)
			},
		},
		"get-automated-reasoning-policy-annotations": {
			Name:   "get-automated-reasoning-policy-annotations",
			Fields: fields_get_automated_reasoning_policy_annotations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyAnnotationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy_annotations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicyAnnotations(ctx, input)
			},
		},
		"get-automated-reasoning-policy-build-workflow": {
			Name:   "get-automated-reasoning-policy-build-workflow",
			Fields: fields_get_automated_reasoning_policy_build_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyBuildWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy_build_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicyBuildWorkflow(ctx, input)
			},
		},
		"get-automated-reasoning-policy-build-workflow-result-assets": {
			Name:   "get-automated-reasoning-policy-build-workflow-result-assets",
			Fields: fields_get_automated_reasoning_policy_build_workflow_result_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyBuildWorkflowResultAssetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy_build_workflow_result_assets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicyBuildWorkflowResultAssets(ctx, input)
			},
		},
		"get-automated-reasoning-policy-next-scenario": {
			Name:   "get-automated-reasoning-policy-next-scenario",
			Fields: fields_get_automated_reasoning_policy_next_scenario,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyNextScenarioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy_next_scenario, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicyNextScenario(ctx, input)
			},
		},
		"get-automated-reasoning-policy-test-case": {
			Name:   "get-automated-reasoning-policy-test-case",
			Fields: fields_get_automated_reasoning_policy_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicyTestCase(ctx, input)
			},
		},
		"get-automated-reasoning-policy-test-result": {
			Name:   "get-automated-reasoning-policy-test-result",
			Fields: fields_get_automated_reasoning_policy_test_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedReasoningPolicyTestResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_reasoning_policy_test_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedReasoningPolicyTestResult(ctx, input)
			},
		},
		"get-custom-model": {
			Name:   "get-custom-model",
			Fields: fields_get_custom_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomModel(ctx, input)
			},
		},
		"get-custom-model-deployment": {
			Name:   "get-custom-model-deployment",
			Fields: fields_get_custom_model_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomModelDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_model_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomModelDeployment(ctx, input)
			},
		},
		"get-evaluation-job": {
			Name:   "get-evaluation-job",
			Fields: fields_get_evaluation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvaluationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_evaluation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvaluationJob(ctx, input)
			},
		},
		"get-foundation-model": {
			Name:   "get-foundation-model",
			Fields: fields_get_foundation_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFoundationModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_foundation_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFoundationModel(ctx, input)
			},
		},
		"get-foundation-model-availability": {
			Name:   "get-foundation-model-availability",
			Fields: fields_get_foundation_model_availability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFoundationModelAvailabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_foundation_model_availability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFoundationModelAvailability(ctx, input)
			},
		},
		"get-guardrail": {
			Name:   "get-guardrail",
			Fields: fields_get_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGuardrail(ctx, input)
			},
		},
		"get-imported-model": {
			Name:   "get-imported-model",
			Fields: fields_get_imported_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportedModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_imported_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImportedModel(ctx, input)
			},
		},
		"get-inference-profile": {
			Name:   "get-inference-profile",
			Fields: fields_get_inference_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInferenceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_inference_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInferenceProfile(ctx, input)
			},
		},
		"get-marketplace-model-endpoint": {
			Name:   "get-marketplace-model-endpoint",
			Fields: fields_get_marketplace_model_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMarketplaceModelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_marketplace_model_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMarketplaceModelEndpoint(ctx, input)
			},
		},
		"get-model-copy-job": {
			Name:   "get-model-copy-job",
			Fields: fields_get_model_copy_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelCopyJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_copy_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelCopyJob(ctx, input)
			},
		},
		"get-model-customization-job": {
			Name:   "get-model-customization-job",
			Fields: fields_get_model_customization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelCustomizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_customization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelCustomizationJob(ctx, input)
			},
		},
		"get-model-import-job": {
			Name:   "get-model-import-job",
			Fields: fields_get_model_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelImportJob(ctx, input)
			},
		},
		"get-model-invocation-job": {
			Name:   "get-model-invocation-job",
			Fields: fields_get_model_invocation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelInvocationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_invocation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelInvocationJob(ctx, input)
			},
		},
		"get-model-invocation-logging-configuration": {
			Name:   "get-model-invocation-logging-configuration",
			Fields: fields_get_model_invocation_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelInvocationLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_invocation_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelInvocationLoggingConfiguration(ctx, input)
			},
		},
		"get-prompt-router": {
			Name:   "get-prompt-router",
			Fields: fields_get_prompt_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPromptRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_prompt_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPromptRouter(ctx, input)
			},
		},
		"get-provisioned-model-throughput": {
			Name:   "get-provisioned-model-throughput",
			Fields: fields_get_provisioned_model_throughput,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProvisionedModelThroughputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_provisioned_model_throughput, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProvisionedModelThroughput(ctx, input)
			},
		},
		"get-use-case-for-model-access": {
			Name:   "get-use-case-for-model-access",
			Fields: fields_get_use_case_for_model_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUseCaseForModelAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_use_case_for_model_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUseCaseForModelAccess(ctx, input)
			},
		},
		"list-automated-reasoning-policies": {
			Name:   "list-automated-reasoning-policies",
			Fields: fields_list_automated_reasoning_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomatedReasoningPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automated_reasoning_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomatedReasoningPolicies(ctx, input)
				}
				var results []*svc.ListAutomatedReasoningPoliciesOutput
				p := svc.NewListAutomatedReasoningPoliciesPaginator(client, input)
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
		"list-automated-reasoning-policy-build-workflows": {
			Name:   "list-automated-reasoning-policy-build-workflows",
			Fields: fields_list_automated_reasoning_policy_build_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomatedReasoningPolicyBuildWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automated_reasoning_policy_build_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomatedReasoningPolicyBuildWorkflows(ctx, input)
				}
				var results []*svc.ListAutomatedReasoningPolicyBuildWorkflowsOutput
				p := svc.NewListAutomatedReasoningPolicyBuildWorkflowsPaginator(client, input)
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
		"list-automated-reasoning-policy-test-cases": {
			Name:   "list-automated-reasoning-policy-test-cases",
			Fields: fields_list_automated_reasoning_policy_test_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomatedReasoningPolicyTestCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automated_reasoning_policy_test_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomatedReasoningPolicyTestCases(ctx, input)
				}
				var results []*svc.ListAutomatedReasoningPolicyTestCasesOutput
				p := svc.NewListAutomatedReasoningPolicyTestCasesPaginator(client, input)
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
		"list-automated-reasoning-policy-test-results": {
			Name:   "list-automated-reasoning-policy-test-results",
			Fields: fields_list_automated_reasoning_policy_test_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomatedReasoningPolicyTestResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automated_reasoning_policy_test_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomatedReasoningPolicyTestResults(ctx, input)
				}
				var results []*svc.ListAutomatedReasoningPolicyTestResultsOutput
				p := svc.NewListAutomatedReasoningPolicyTestResultsPaginator(client, input)
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
		"list-custom-model-deployments": {
			Name:   "list-custom-model-deployments",
			Fields: fields_list_custom_model_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomModelDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_model_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomModelDeployments(ctx, input)
				}
				var results []*svc.ListCustomModelDeploymentsOutput
				p := svc.NewListCustomModelDeploymentsPaginator(client, input)
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
		"list-custom-models": {
			Name:   "list-custom-models",
			Fields: fields_list_custom_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomModels(ctx, input)
				}
				var results []*svc.ListCustomModelsOutput
				p := svc.NewListCustomModelsPaginator(client, input)
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
		"list-enforced-guardrails-configuration": {
			Name:   "list-enforced-guardrails-configuration",
			Fields: fields_list_enforced_guardrails_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnforcedGuardrailsConfigurationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_enforced_guardrails_configuration, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnforcedGuardrailsConfiguration(ctx, input)
				}
				var results []*svc.ListEnforcedGuardrailsConfigurationOutput
				p := svc.NewListEnforcedGuardrailsConfigurationPaginator(client, input)
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
		"list-evaluation-jobs": {
			Name:   "list-evaluation-jobs",
			Fields: fields_list_evaluation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEvaluationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_evaluation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEvaluationJobs(ctx, input)
				}
				var results []*svc.ListEvaluationJobsOutput
				p := svc.NewListEvaluationJobsPaginator(client, input)
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
		"list-foundation-model-agreement-offers": {
			Name:   "list-foundation-model-agreement-offers",
			Fields: fields_list_foundation_model_agreement_offers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFoundationModelAgreementOffersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_foundation_model_agreement_offers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFoundationModelAgreementOffers(ctx, input)
			},
		},
		"list-foundation-models": {
			Name:   "list-foundation-models",
			Fields: fields_list_foundation_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFoundationModelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_foundation_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFoundationModels(ctx, input)
			},
		},
		"list-guardrails": {
			Name:   "list-guardrails",
			Fields: fields_list_guardrails,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGuardrailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_guardrails, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGuardrails(ctx, input)
				}
				var results []*svc.ListGuardrailsOutput
				p := svc.NewListGuardrailsPaginator(client, input)
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
		"list-imported-models": {
			Name:   "list-imported-models",
			Fields: fields_list_imported_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportedModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_imported_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportedModels(ctx, input)
				}
				var results []*svc.ListImportedModelsOutput
				p := svc.NewListImportedModelsPaginator(client, input)
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
		"list-inference-profiles": {
			Name:   "list-inference-profiles",
			Fields: fields_list_inference_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceProfiles(ctx, input)
				}
				var results []*svc.ListInferenceProfilesOutput
				p := svc.NewListInferenceProfilesPaginator(client, input)
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
		"list-marketplace-model-endpoints": {
			Name:   "list-marketplace-model-endpoints",
			Fields: fields_list_marketplace_model_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMarketplaceModelEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_marketplace_model_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMarketplaceModelEndpoints(ctx, input)
				}
				var results []*svc.ListMarketplaceModelEndpointsOutput
				p := svc.NewListMarketplaceModelEndpointsPaginator(client, input)
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
		"list-model-copy-jobs": {
			Name:   "list-model-copy-jobs",
			Fields: fields_list_model_copy_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelCopyJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_copy_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelCopyJobs(ctx, input)
				}
				var results []*svc.ListModelCopyJobsOutput
				p := svc.NewListModelCopyJobsPaginator(client, input)
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
		"list-model-customization-jobs": {
			Name:   "list-model-customization-jobs",
			Fields: fields_list_model_customization_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelCustomizationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_customization_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelCustomizationJobs(ctx, input)
				}
				var results []*svc.ListModelCustomizationJobsOutput
				p := svc.NewListModelCustomizationJobsPaginator(client, input)
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
		"list-model-import-jobs": {
			Name:   "list-model-import-jobs",
			Fields: fields_list_model_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelImportJobs(ctx, input)
				}
				var results []*svc.ListModelImportJobsOutput
				p := svc.NewListModelImportJobsPaginator(client, input)
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
		"list-model-invocation-jobs": {
			Name:   "list-model-invocation-jobs",
			Fields: fields_list_model_invocation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelInvocationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_invocation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelInvocationJobs(ctx, input)
				}
				var results []*svc.ListModelInvocationJobsOutput
				p := svc.NewListModelInvocationJobsPaginator(client, input)
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
		"list-prompt-routers": {
			Name:   "list-prompt-routers",
			Fields: fields_list_prompt_routers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPromptRoutersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_prompt_routers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPromptRouters(ctx, input)
				}
				var results []*svc.ListPromptRoutersOutput
				p := svc.NewListPromptRoutersPaginator(client, input)
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
		"list-provisioned-model-throughputs": {
			Name:   "list-provisioned-model-throughputs",
			Fields: fields_list_provisioned_model_throughputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisionedModelThroughputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provisioned_model_throughputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProvisionedModelThroughputs(ctx, input)
				}
				var results []*svc.ListProvisionedModelThroughputsOutput
				p := svc.NewListProvisionedModelThroughputsPaginator(client, input)
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
		"put-enforced-guardrail-configuration": {
			Name:   "put-enforced-guardrail-configuration",
			Fields: fields_put_enforced_guardrail_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEnforcedGuardrailConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_enforced_guardrail_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEnforcedGuardrailConfiguration(ctx, input)
			},
		},
		"put-model-invocation-logging-configuration": {
			Name:   "put-model-invocation-logging-configuration",
			Fields: fields_put_model_invocation_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutModelInvocationLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_model_invocation_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutModelInvocationLoggingConfiguration(ctx, input)
			},
		},
		"put-use-case-for-model-access": {
			Name:   "put-use-case-for-model-access",
			Fields: fields_put_use_case_for_model_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutUseCaseForModelAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_use_case_for_model_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutUseCaseForModelAccess(ctx, input)
			},
		},
		"register-marketplace-model-endpoint": {
			Name:   "register-marketplace-model-endpoint",
			Fields: fields_register_marketplace_model_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterMarketplaceModelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_marketplace_model_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterMarketplaceModelEndpoint(ctx, input)
			},
		},
		"start-automated-reasoning-policy-build-workflow": {
			Name:   "start-automated-reasoning-policy-build-workflow",
			Fields: fields_start_automated_reasoning_policy_build_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAutomatedReasoningPolicyBuildWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_automated_reasoning_policy_build_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAutomatedReasoningPolicyBuildWorkflow(ctx, input)
			},
		},
		"start-automated-reasoning-policy-test-workflow": {
			Name:   "start-automated-reasoning-policy-test-workflow",
			Fields: fields_start_automated_reasoning_policy_test_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAutomatedReasoningPolicyTestWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_automated_reasoning_policy_test_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAutomatedReasoningPolicyTestWorkflow(ctx, input)
			},
		},
		"stop-evaluation-job": {
			Name:   "stop-evaluation-job",
			Fields: fields_stop_evaluation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEvaluationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_evaluation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEvaluationJob(ctx, input)
			},
		},
		"stop-model-customization-job": {
			Name:   "stop-model-customization-job",
			Fields: fields_stop_model_customization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopModelCustomizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_model_customization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopModelCustomizationJob(ctx, input)
			},
		},
		"stop-model-invocation-job": {
			Name:   "stop-model-invocation-job",
			Fields: fields_stop_model_invocation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopModelInvocationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_model_invocation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopModelInvocationJob(ctx, input)
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
		"update-automated-reasoning-policy": {
			Name:   "update-automated-reasoning-policy",
			Fields: fields_update_automated_reasoning_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomatedReasoningPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automated_reasoning_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomatedReasoningPolicy(ctx, input)
			},
		},
		"update-automated-reasoning-policy-annotations": {
			Name:   "update-automated-reasoning-policy-annotations",
			Fields: fields_update_automated_reasoning_policy_annotations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomatedReasoningPolicyAnnotationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automated_reasoning_policy_annotations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomatedReasoningPolicyAnnotations(ctx, input)
			},
		},
		"update-automated-reasoning-policy-test-case": {
			Name:   "update-automated-reasoning-policy-test-case",
			Fields: fields_update_automated_reasoning_policy_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomatedReasoningPolicyTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automated_reasoning_policy_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomatedReasoningPolicyTestCase(ctx, input)
			},
		},
		"update-custom-model-deployment": {
			Name:   "update-custom-model-deployment",
			Fields: fields_update_custom_model_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomModelDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_model_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomModelDeployment(ctx, input)
			},
		},
		"update-guardrail": {
			Name:   "update-guardrail",
			Fields: fields_update_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGuardrail(ctx, input)
			},
		},
		"update-marketplace-model-endpoint": {
			Name:   "update-marketplace-model-endpoint",
			Fields: fields_update_marketplace_model_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMarketplaceModelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_marketplace_model_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMarketplaceModelEndpoint(ctx, input)
			},
		},
		"update-provisioned-model-throughput": {
			Name:   "update-provisioned-model-throughput",
			Fields: fields_update_provisioned_model_throughput,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProvisionedModelThroughputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_provisioned_model_throughput, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProvisionedModelThroughput(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bedrock", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
