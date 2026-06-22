package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

var fields_add_association = []leanruntime.Field{
	{Name: "AssociationType", Flag: "association-type", Type: "types.AssociationEdgeType", Required: false},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
}

var fields_add_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_associate_trial_component = []leanruntime.Field{
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: true},
}

var fields_attach_cluster_node_volume = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_batch_add_cluster_nodes = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodesToAdd", Flag: "nodes-to-add", Type: "[]types.AddClusterNodeSpecification", Required: true},
}

var fields_batch_delete_cluster_nodes = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodeIds", Flag: "node-ids", Type: "[]string", Required: false},
	{Name: "NodeLogicalIds", Flag: "node-logical-ids", Type: "[]string", Required: false},
}

var fields_batch_describe_model_package = []leanruntime.Field{
	{Name: "ModelPackageArnList", Flag: "model-package-arn-list", Type: "[]string", Required: true},
}

var fields_batch_reboot_cluster_nodes = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodeIds", Flag: "node-ids", Type: "[]string", Required: false},
	{Name: "NodeLogicalIds", Flag: "node-logical-ids", Type: "[]string", Required: false},
}

var fields_batch_replace_cluster_nodes = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodeIds", Flag: "node-ids", Type: "[]string", Required: false},
	{Name: "NodeLogicalIds", Flag: "node-logical-ids", Type: "[]string", Required: false},
}

var fields_create_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "ActionType", Flag: "action-type", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MetadataProperties", Flag: "metadata-properties", Type: "*types.MetadataProperties", Required: false},
	{Name: "Properties", Flag: "properties", Type: "map[string]string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.ActionSource", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ActionStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_algorithm = []leanruntime.Field{
	{Name: "AlgorithmDescription", Flag: "algorithm-description", Type: "*string", Required: false},
	{Name: "AlgorithmName", Flag: "algorithm-name", Type: "*string", Required: true},
	{Name: "CertifyForMarketplace", Flag: "certify-for-marketplace", Type: "*bool", Required: false},
	{Name: "InferenceSpecification", Flag: "inference-specification", Type: "*types.InferenceSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingSpecification", Flag: "training-specification", Type: "*types.TrainingSpecification", Required: true},
	{Name: "ValidationSpecification", Flag: "validation-specification", Type: "*types.AlgorithmValidationSpecification", Required: false},
}

var fields_create_app = []leanruntime.Field{
	{Name: "AppName", Flag: "app-name", Type: "*string", Required: true},
	{Name: "AppType", Flag: "app-type", Type: "types.AppType", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "RecoveryMode", Flag: "recovery-mode", Type: "*bool", Required: false},
	{Name: "ResourceSpec", Flag: "resource-spec", Type: "*types.ResourceSpec", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: false},
}

var fields_create_app_image_config = []leanruntime.Field{
	{Name: "AppImageConfigName", Flag: "app-image-config-name", Type: "*string", Required: true},
	{Name: "CodeEditorAppImageConfig", Flag: "code-editor-app-image-config", Type: "*types.CodeEditorAppImageConfig", Required: false},
	{Name: "JupyterLabAppImageConfig", Flag: "jupyter-lab-app-image-config", Type: "*types.JupyterLabAppImageConfig", Required: false},
	{Name: "KernelGatewayImageConfig", Flag: "kernel-gateway-image-config", Type: "*types.KernelGatewayImageConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_artifact = []leanruntime.Field{
	{Name: "ArtifactName", Flag: "artifact-name", Type: "*string", Required: false},
	{Name: "ArtifactType", Flag: "artifact-type", Type: "*string", Required: true},
	{Name: "MetadataProperties", Flag: "metadata-properties", Type: "*types.MetadataProperties", Required: false},
	{Name: "Properties", Flag: "properties", Type: "map[string]string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.ArtifactSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_auto_ml_job = []leanruntime.Field{
	{Name: "AutoMLJobConfig", Flag: "auto-ml-job-config", Type: "*types.AutoMLJobConfig", Required: false},
	{Name: "AutoMLJobName", Flag: "auto-ml-job-name", Type: "*string", Required: true},
	{Name: "AutoMLJobObjective", Flag: "auto-ml-job-objective", Type: "*types.AutoMLJobObjective", Required: false},
	{Name: "GenerateCandidateDefinitionsOnly", Flag: "generate-candidate-definitions-only", Type: "*bool", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "[]types.AutoMLChannel", Required: true},
	{Name: "ModelDeployConfig", Flag: "model-deploy-config", Type: "*types.ModelDeployConfig", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.AutoMLOutputDataConfig", Required: true},
	{Name: "ProblemType", Flag: "problem-type", Type: "types.ProblemType", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_auto_ml_job_v2 = []leanruntime.Field{
	{Name: "AutoMLComputeConfig", Flag: "auto-ml-compute-config", Type: "*types.AutoMLComputeConfig", Required: false},
	{Name: "AutoMLJobInputDataConfig", Flag: "auto-ml-job-input-data-config", Type: "[]types.AutoMLJobChannel", Required: true},
	{Name: "AutoMLJobName", Flag: "auto-ml-job-name", Type: "*string", Required: true},
	{Name: "AutoMLJobObjective", Flag: "auto-ml-job-objective", Type: "*types.AutoMLJobObjective", Required: false},
	{Name: "AutoMLProblemTypeConfig", Flag: "auto-ml-problem-type-config", Type: "types.AutoMLProblemTypeConfig", Required: true},
	{Name: "DataSplitConfig", Flag: "data-split-config", Type: "*types.AutoMLDataSplitConfig", Required: false},
	{Name: "ModelDeployConfig", Flag: "model-deploy-config", Type: "*types.ModelDeployConfig", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.AutoMLOutputDataConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SecurityConfig", Flag: "security-config", Type: "*types.AutoMLSecurityConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "AutoScaling", Flag: "auto-scaling", Type: "*types.ClusterAutoScalingConfig", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ClusterRole", Flag: "cluster-role", Type: "*string", Required: false},
	{Name: "InstanceGroups", Flag: "instance-groups", Type: "[]types.ClusterInstanceGroupSpecification", Required: false},
	{Name: "NodeProvisioningMode", Flag: "node-provisioning-mode", Type: "types.ClusterNodeProvisioningMode", Required: false},
	{Name: "NodeRecovery", Flag: "node-recovery", Type: "types.ClusterNodeRecovery", Required: false},
	{Name: "Orchestrator", Flag: "orchestrator", Type: "*types.ClusterOrchestrator", Required: false},
	{Name: "RestrictedInstanceGroups", Flag: "restricted-instance-groups", Type: "[]types.ClusterRestrictedInstanceGroupSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TieredStorageConfig", Flag: "tiered-storage-config", Type: "*types.ClusterTieredStorageConfig", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_cluster_scheduler_config = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchedulerConfig", Flag: "scheduler-config", Type: "*types.SchedulerConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_code_repository = []leanruntime.Field{
	{Name: "CodeRepositoryName", Flag: "code-repository-name", Type: "*string", Required: true},
	{Name: "GitConfig", Flag: "git-config", Type: "*types.GitConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_compilation_job = []leanruntime.Field{
	{Name: "CompilationJobName", Flag: "compilation-job-name", Type: "*string", Required: true},
	{Name: "InputConfig", Flag: "input-config", Type: "*types.InputConfig", Required: false},
	{Name: "ModelPackageVersionArn", Flag: "model-package-version-arn", Type: "*string", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.StoppingCondition", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.NeoVpcConfig", Required: false},
}

var fields_create_compute_quota = []leanruntime.Field{
	{Name: "ActivationState", Flag: "activation-state", Type: "types.ActivationState", Required: false},
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "ComputeQuotaConfig", Flag: "compute-quota-config", Type: "*types.ComputeQuotaConfig", Required: true},
	{Name: "ComputeQuotaTarget", Flag: "compute-quota-target", Type: "*types.ComputeQuotaTarget", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_context = []leanruntime.Field{
	{Name: "ContextName", Flag: "context-name", Type: "*string", Required: true},
	{Name: "ContextType", Flag: "context-type", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Properties", Flag: "properties", Type: "map[string]string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.ContextSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_data_quality_job_definition = []leanruntime.Field{
	{Name: "DataQualityAppSpecification", Flag: "data-quality-app-specification", Type: "*types.DataQualityAppSpecification", Required: true},
	{Name: "DataQualityBaselineConfig", Flag: "data-quality-baseline-config", Type: "*types.DataQualityBaselineConfig", Required: false},
	{Name: "DataQualityJobInput", Flag: "data-quality-job-input", Type: "*types.DataQualityJobInput", Required: true},
	{Name: "DataQualityJobOutputConfig", Flag: "data-quality-job-output-config", Type: "*types.MonitoringOutputConfig", Required: true},
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
	{Name: "JobResources", Flag: "job-resources", Type: "*types.MonitoringResources", Required: true},
	{Name: "NetworkConfig", Flag: "network-config", Type: "*types.MonitoringNetworkConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.MonitoringStoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_device_fleet = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "EnableIotRoleAlias", Flag: "enable-iot-role-alias", Type: "*bool", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.EdgeOutputConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "AppNetworkAccessType", Flag: "app-network-access-type", Type: "types.AppNetworkAccessType", Required: false},
	{Name: "AppSecurityGroupManagement", Flag: "app-security-group-management", Type: "types.AppSecurityGroupManagement", Required: false},
	{Name: "AuthMode", Flag: "auth-mode", Type: "types.AuthMode", Required: true},
	{Name: "DefaultSpaceSettings", Flag: "default-space-settings", Type: "*types.DefaultSpaceSettings", Required: false},
	{Name: "DefaultUserSettings", Flag: "default-user-settings", Type: "*types.UserSettings", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainSettings", Flag: "domain-settings", Type: "*types.DomainSettings", Required: false},
	{Name: "HomeEfsFileSystemKmsKeyId", Flag: "home-efs-file-system-kms-key-id", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "TagPropagation", Flag: "tag-propagation", Type: "types.TagPropagation", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_create_edge_deployment_plan = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "ModelConfigs", Flag: "model-configs", Type: "[]types.EdgeDeploymentModelConfig", Required: true},
	{Name: "Stages", Flag: "stages", Type: "[]types.DeploymentStage", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_edge_deployment_stage = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "Stages", Flag: "stages", Type: "[]types.DeploymentStage", Required: true},
}

var fields_create_edge_packaging_job = []leanruntime.Field{
	{Name: "CompilationJobName", Flag: "compilation-job-name", Type: "*string", Required: true},
	{Name: "EdgePackagingJobName", Flag: "edge-packaging-job-name", Type: "*string", Required: true},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "ModelVersion", Flag: "model-version", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.EdgeOutputConfig", Required: true},
	{Name: "ResourceKey", Flag: "resource-key", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_endpoint = []leanruntime.Field{
	{Name: "DeploymentConfig", Flag: "deployment-config", Type: "*types.DeploymentConfig", Required: false},
	{Name: "EndpointConfigName", Flag: "endpoint-config-name", Type: "*string", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_endpoint_config = []leanruntime.Field{
	{Name: "AsyncInferenceConfig", Flag: "async-inference-config", Type: "*types.AsyncInferenceConfig", Required: false},
	{Name: "DataCaptureConfig", Flag: "data-capture-config", Type: "*types.DataCaptureConfig", Required: false},
	{Name: "EnableNetworkIsolation", Flag: "enable-network-isolation", Type: "*bool", Required: false},
	{Name: "EndpointConfigName", Flag: "endpoint-config-name", Type: "*string", Required: true},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "ExplainerConfig", Flag: "explainer-config", Type: "*types.ExplainerConfig", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MetricsConfig", Flag: "metrics-config", Type: "*types.MetricsConfig", Required: false},
	{Name: "ProductionVariants", Flag: "production-variants", Type: "[]types.ProductionVariant", Required: true},
	{Name: "ShadowProductionVariants", Flag: "shadow-production-variants", Type: "[]types.ProductionVariant", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_experiment = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_feature_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventTimeFeatureName", Flag: "event-time-feature-name", Type: "*string", Required: true},
	{Name: "FeatureDefinitions", Flag: "feature-definitions", Type: "[]types.FeatureDefinition", Required: true},
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "OfflineStoreConfig", Flag: "offline-store-config", Type: "*types.OfflineStoreConfig", Required: false},
	{Name: "OnlineStoreConfig", Flag: "online-store-config", Type: "*types.OnlineStoreConfig", Required: false},
	{Name: "RecordIdentifierFeatureName", Flag: "record-identifier-feature-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThroughputConfig", Flag: "throughput-config", Type: "*types.ThroughputConfig", Required: false},
}

var fields_create_flow_definition = []leanruntime.Field{
	{Name: "FlowDefinitionName", Flag: "flow-definition-name", Type: "*string", Required: true},
	{Name: "HumanLoopActivationConfig", Flag: "human-loop-activation-config", Type: "*types.HumanLoopActivationConfig", Required: false},
	{Name: "HumanLoopConfig", Flag: "human-loop-config", Type: "*types.HumanLoopConfig", Required: false},
	{Name: "HumanLoopRequestSource", Flag: "human-loop-request-source", Type: "*types.HumanLoopRequestSource", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.FlowDefinitionOutputConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_hub = []leanruntime.Field{
	{Name: "HubDescription", Flag: "hub-description", Type: "*string", Required: true},
	{Name: "HubDisplayName", Flag: "hub-display-name", Type: "*string", Required: false},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "HubSearchKeywords", Flag: "hub-search-keywords", Type: "[]string", Required: false},
	{Name: "S3StorageConfig", Flag: "s3-storage-config", Type: "*types.HubS3StorageConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_hub_content_presigned_urls = []leanruntime.Field{
	{Name: "AccessConfig", Flag: "access-config", Type: "*types.PresignedUrlAccessConfig", Required: false},
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubContentVersion", Flag: "hub-content-version", Type: "*string", Required: false},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_create_hub_content_reference = []leanruntime.Field{
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: false},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "MinVersion", Flag: "min-version", Type: "*string", Required: false},
	{Name: "SageMakerPublicHubContentArn", Flag: "sage-maker-public-hub-content-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_human_task_ui = []leanruntime.Field{
	{Name: "HumanTaskUiName", Flag: "human-task-ui-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UiTemplate", Flag: "ui-template", Type: "*types.UiTemplate", Required: true},
}

var fields_create_hyper_parameter_tuning_job = []leanruntime.Field{
	{Name: "Autotune", Flag: "autotune", Type: "*types.Autotune", Required: false},
	{Name: "HyperParameterTuningJobConfig", Flag: "hyper-parameter-tuning-job-config", Type: "*types.HyperParameterTuningJobConfig", Required: true},
	{Name: "HyperParameterTuningJobName", Flag: "hyper-parameter-tuning-job-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingJobDefinition", Flag: "training-job-definition", Type: "*types.HyperParameterTrainingJobDefinition", Required: false},
	{Name: "TrainingJobDefinitions", Flag: "training-job-definitions", Type: "[]types.HyperParameterTrainingJobDefinition", Required: false},
	{Name: "WarmStartConfig", Flag: "warm-start-config", Type: "*types.HyperParameterTuningJobWarmStartConfig", Required: false},
}

var fields_create_image = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_image_version = []leanruntime.Field{
	{Name: "Aliases", Flag: "aliases", Type: "[]string", Required: false},
	{Name: "BaseImage", Flag: "base-image", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Horovod", Flag: "horovod", Type: "*bool", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: false},
	{Name: "MLFramework", Flag: "ml-framework", Type: "*string", Required: false},
	{Name: "Processor", Flag: "processor", Type: "types.Processor", Required: false},
	{Name: "ProgrammingLang", Flag: "programming-lang", Type: "*string", Required: false},
	{Name: "ReleaseNotes", Flag: "release-notes", Type: "*string", Required: false},
	{Name: "VendorGuidance", Flag: "vendor-guidance", Type: "types.VendorGuidance", Required: false},
}

var fields_create_inference_component = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: true},
	{Name: "RuntimeConfig", Flag: "runtime-config", Type: "*types.InferenceComponentRuntimeConfig", Required: false},
	{Name: "Specification", Flag: "specification", Type: "*types.InferenceComponentSpecification", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VariantName", Flag: "variant-name", Type: "*string", Required: false},
}

var fields_create_inference_experiment = []leanruntime.Field{
	{Name: "DataStorageConfig", Flag: "data-storage-config", Type: "*types.InferenceExperimentDataStorageConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "KmsKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "ModelVariants", Flag: "model-variants", Type: "[]types.ModelVariantConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.InferenceExperimentSchedule", Required: false},
	{Name: "ShadowModeConfig", Flag: "shadow-mode-config", Type: "*types.ShadowModeConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.InferenceExperimentType", Required: true},
}

var fields_create_inference_recommendations_job = []leanruntime.Field{
	{Name: "InputConfig", Flag: "input-config", Type: "*types.RecommendationJobInputConfig", Required: true},
	{Name: "JobDescription", Flag: "job-description", Type: "*string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobType", Flag: "job-type", Type: "types.RecommendationJobType", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.RecommendationJobOutputConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingConditions", Flag: "stopping-conditions", Type: "*types.RecommendationJobStoppingConditions", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_labeling_job = []leanruntime.Field{
	{Name: "HumanTaskConfig", Flag: "human-task-config", Type: "*types.HumanTaskConfig", Required: true},
	{Name: "InputConfig", Flag: "input-config", Type: "*types.LabelingJobInputConfig", Required: true},
	{Name: "LabelAttributeName", Flag: "label-attribute-name", Type: "*string", Required: true},
	{Name: "LabelCategoryConfigS3Uri", Flag: "label-category-config-s3-uri", Type: "*string", Required: false},
	{Name: "LabelingJobAlgorithmsConfig", Flag: "labeling-job-algorithms-config", Type: "*types.LabelingJobAlgorithmsConfig", Required: false},
	{Name: "LabelingJobName", Flag: "labeling-job-name", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.LabelingJobOutputConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingConditions", Flag: "stopping-conditions", Type: "*types.LabelingJobStoppingConditions", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_mlflow_app = []leanruntime.Field{
	{Name: "AccountDefaultStatus", Flag: "account-default-status", Type: "types.AccountDefaultStatus", Required: false},
	{Name: "ArtifactStoreUri", Flag: "artifact-store-uri", Type: "*string", Required: true},
	{Name: "DefaultDomainIdList", Flag: "default-domain-id-list", Type: "[]string", Required: false},
	{Name: "ModelRegistrationMode", Flag: "model-registration-mode", Type: "types.ModelRegistrationMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WeeklyMaintenanceWindowStart", Flag: "weekly-maintenance-window-start", Type: "*string", Required: false},
}

var fields_create_mlflow_tracking_server = []leanruntime.Field{
	{Name: "ArtifactStoreUri", Flag: "artifact-store-uri", Type: "*string", Required: true},
	{Name: "AutomaticModelRegistration", Flag: "automatic-model-registration", Type: "*bool", Required: false},
	{Name: "MlflowVersion", Flag: "mlflow-version", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
	{Name: "TrackingServerSize", Flag: "tracking-server-size", Type: "types.TrackingServerSize", Required: false},
	{Name: "WeeklyMaintenanceWindowStart", Flag: "weekly-maintenance-window-start", Type: "*string", Required: false},
}

var fields_create_model = []leanruntime.Field{
	{Name: "Containers", Flag: "containers", Type: "[]types.ContainerDefinition", Required: false},
	{Name: "EnableNetworkIsolation", Flag: "enable-network-isolation", Type: "*bool", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "InferenceExecutionConfig", Flag: "inference-execution-config", Type: "*types.InferenceExecutionConfig", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "PrimaryContainer", Flag: "primary-container", Type: "*types.ContainerDefinition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_model_bias_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
	{Name: "JobResources", Flag: "job-resources", Type: "*types.MonitoringResources", Required: true},
	{Name: "ModelBiasAppSpecification", Flag: "model-bias-app-specification", Type: "*types.ModelBiasAppSpecification", Required: true},
	{Name: "ModelBiasBaselineConfig", Flag: "model-bias-baseline-config", Type: "*types.ModelBiasBaselineConfig", Required: false},
	{Name: "ModelBiasJobInput", Flag: "model-bias-job-input", Type: "*types.ModelBiasJobInput", Required: true},
	{Name: "ModelBiasJobOutputConfig", Flag: "model-bias-job-output-config", Type: "*types.MonitoringOutputConfig", Required: true},
	{Name: "NetworkConfig", Flag: "network-config", Type: "*types.MonitoringNetworkConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.MonitoringStoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_card = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
	{Name: "ModelCardStatus", Flag: "model-card-status", Type: "types.ModelCardStatus", Required: true},
	{Name: "SecurityConfig", Flag: "security-config", Type: "*types.ModelCardSecurityConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_card_export_job = []leanruntime.Field{
	{Name: "ModelCardExportJobName", Flag: "model-card-export-job-name", Type: "*string", Required: true},
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
	{Name: "ModelCardVersion", Flag: "model-card-version", Type: "*int32", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.ModelCardExportOutputConfig", Required: true},
}

var fields_create_model_explainability_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
	{Name: "JobResources", Flag: "job-resources", Type: "*types.MonitoringResources", Required: true},
	{Name: "ModelExplainabilityAppSpecification", Flag: "model-explainability-app-specification", Type: "*types.ModelExplainabilityAppSpecification", Required: true},
	{Name: "ModelExplainabilityBaselineConfig", Flag: "model-explainability-baseline-config", Type: "*types.ModelExplainabilityBaselineConfig", Required: false},
	{Name: "ModelExplainabilityJobInput", Flag: "model-explainability-job-input", Type: "*types.ModelExplainabilityJobInput", Required: true},
	{Name: "ModelExplainabilityJobOutputConfig", Flag: "model-explainability-job-output-config", Type: "*types.MonitoringOutputConfig", Required: true},
	{Name: "NetworkConfig", Flag: "network-config", Type: "*types.MonitoringNetworkConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.MonitoringStoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_package = []leanruntime.Field{
	{Name: "AdditionalInferenceSpecifications", Flag: "additional-inference-specifications", Type: "[]types.AdditionalInferenceSpecificationDefinition", Required: false},
	{Name: "CertifyForMarketplace", Flag: "certify-for-marketplace", Type: "*bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerMetadataProperties", Flag: "customer-metadata-properties", Type: "map[string]string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DriftCheckBaselines", Flag: "drift-check-baselines", Type: "*types.DriftCheckBaselines", Required: false},
	{Name: "InferenceSpecification", Flag: "inference-specification", Type: "*types.InferenceSpecification", Required: false},
	{Name: "MetadataProperties", Flag: "metadata-properties", Type: "*types.MetadataProperties", Required: false},
	{Name: "ModelApprovalStatus", Flag: "model-approval-status", Type: "types.ModelApprovalStatus", Required: false},
	{Name: "ModelCard", Flag: "model-card", Type: "*types.ModelPackageModelCard", Required: false},
	{Name: "ModelLifeCycle", Flag: "model-life-cycle", Type: "*types.ModelLifeCycle", Required: false},
	{Name: "ModelMetrics", Flag: "model-metrics", Type: "*types.ModelMetrics", Required: false},
	{Name: "ModelPackageDescription", Flag: "model-package-description", Type: "*string", Required: false},
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: false},
	{Name: "ModelPackageName", Flag: "model-package-name", Type: "*string", Required: false},
	{Name: "ModelPackageRegistrationType", Flag: "model-package-registration-type", Type: "types.ModelPackageRegistrationType", Required: false},
	{Name: "SamplePayloadUrl", Flag: "sample-payload-url", Type: "*string", Required: false},
	{Name: "SecurityConfig", Flag: "security-config", Type: "*types.ModelPackageSecurityConfig", Required: false},
	{Name: "SkipModelValidation", Flag: "skip-model-validation", Type: "types.SkipModelValidation", Required: false},
	{Name: "SourceAlgorithmSpecification", Flag: "source-algorithm-specification", Type: "*types.SourceAlgorithmSpecification", Required: false},
	{Name: "SourceUri", Flag: "source-uri", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Task", Flag: "task", Type: "*string", Required: false},
	{Name: "ValidationSpecification", Flag: "validation-specification", Type: "*types.ModelPackageValidationSpecification", Required: false},
}

var fields_create_model_package_group = []leanruntime.Field{
	{Name: "ModelPackageGroupDescription", Flag: "model-package-group-description", Type: "*string", Required: false},
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_quality_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
	{Name: "JobResources", Flag: "job-resources", Type: "*types.MonitoringResources", Required: true},
	{Name: "ModelQualityAppSpecification", Flag: "model-quality-app-specification", Type: "*types.ModelQualityAppSpecification", Required: true},
	{Name: "ModelQualityBaselineConfig", Flag: "model-quality-baseline-config", Type: "*types.ModelQualityBaselineConfig", Required: false},
	{Name: "ModelQualityJobInput", Flag: "model-quality-job-input", Type: "*types.ModelQualityJobInput", Required: true},
	{Name: "ModelQualityJobOutputConfig", Flag: "model-quality-job-output-config", Type: "*types.MonitoringOutputConfig", Required: true},
	{Name: "NetworkConfig", Flag: "network-config", Type: "*types.MonitoringNetworkConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.MonitoringStoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_monitoring_schedule = []leanruntime.Field{
	{Name: "MonitoringScheduleConfig", Flag: "monitoring-schedule-config", Type: "*types.MonitoringScheduleConfig", Required: true},
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_notebook_instance = []leanruntime.Field{
	{Name: "AcceleratorTypes", Flag: "accelerator-types", Type: "[]types.NotebookInstanceAcceleratorType", Required: false},
	{Name: "AdditionalCodeRepositories", Flag: "additional-code-repositories", Type: "[]string", Required: false},
	{Name: "DefaultCodeRepository", Flag: "default-code-repository", Type: "*string", Required: false},
	{Name: "DirectInternetAccess", Flag: "direct-internet-access", Type: "types.DirectInternetAccess", Required: false},
	{Name: "InstanceMetadataServiceConfiguration", Flag: "instance-metadata-service-configuration", Type: "*types.InstanceMetadataServiceConfiguration", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IPAddressType", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LifecycleConfigName", Flag: "lifecycle-config-name", Type: "*string", Required: false},
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
	{Name: "PlatformIdentifier", Flag: "platform-identifier", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "RootAccess", Flag: "root-access", Type: "types.RootAccess", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeSizeInGB", Flag: "volume-size-in-gb", Type: "*int32", Required: false},
}

var fields_create_notebook_instance_lifecycle_config = []leanruntime.Field{
	{Name: "NotebookInstanceLifecycleConfigName", Flag: "notebook-instance-lifecycle-config-name", Type: "*string", Required: true},
	{Name: "OnCreate", Flag: "on-create", Type: "[]types.NotebookInstanceLifecycleHook", Required: false},
	{Name: "OnStart", Flag: "on-start", Type: "[]types.NotebookInstanceLifecycleHook", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_optimization_job = []leanruntime.Field{
	{Name: "DeploymentInstanceType", Flag: "deployment-instance-type", Type: "types.OptimizationJobDeploymentInstanceType", Required: true},
	{Name: "MaxInstanceCount", Flag: "max-instance-count", Type: "*int32", Required: false},
	{Name: "ModelSource", Flag: "model-source", Type: "*types.OptimizationJobModelSource", Required: true},
	{Name: "OptimizationConfigs", Flag: "optimization-configs", Type: "[]types.OptimizationConfig", Required: true},
	{Name: "OptimizationEnvironment", Flag: "optimization-environment", Type: "map[string]string", Required: false},
	{Name: "OptimizationJobName", Flag: "optimization-job-name", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OptimizationJobOutputConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.StoppingCondition", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.OptimizationVpcConfig", Required: false},
}

var fields_create_partner_app = []leanruntime.Field{
	{Name: "ApplicationConfig", Flag: "application-config", Type: "*types.PartnerAppConfig", Required: false},
	{Name: "AuthType", Flag: "auth-type", Type: "types.PartnerAppAuthType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnableAutoMinorVersionUpgrade", Flag: "enable-auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "EnableIamSessionBasedIdentity", Flag: "enable-iam-session-based-identity", Type: "*bool", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MaintenanceConfig", Flag: "maintenance-config", Type: "*types.PartnerAppMaintenanceConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Tier", Flag: "tier", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.PartnerAppType", Required: true},
}

var fields_create_partner_app_presigned_url = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ExpiresInSeconds", Flag: "expires-in-seconds", Type: "*int32", Required: false},
	{Name: "SessionExpirationDurationInSeconds", Flag: "session-expiration-duration-in-seconds", Type: "*int32", Required: false},
}

var fields_create_pipeline = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ParallelismConfiguration", Flag: "parallelism-configuration", Type: "*types.ParallelismConfiguration", Required: false},
	{Name: "PipelineDefinition", Flag: "pipeline-definition", Type: "*string", Required: false},
	{Name: "PipelineDefinitionS3Location", Flag: "pipeline-definition-s3-location", Type: "*types.PipelineDefinitionS3Location", Required: false},
	{Name: "PipelineDescription", Flag: "pipeline-description", Type: "*string", Required: false},
	{Name: "PipelineDisplayName", Flag: "pipeline-display-name", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_presigned_domain_url = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "ExpiresInSeconds", Flag: "expires-in-seconds", Type: "*int32", Required: false},
	{Name: "LandingUri", Flag: "landing-uri", Type: "*string", Required: false},
	{Name: "SessionExpirationDurationInSeconds", Flag: "session-expiration-duration-in-seconds", Type: "*int32", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: false},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: true},
}

var fields_create_presigned_mlflow_app_url = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ExpiresInSeconds", Flag: "expires-in-seconds", Type: "*int32", Required: false},
	{Name: "SessionExpirationDurationInSeconds", Flag: "session-expiration-duration-in-seconds", Type: "*int32", Required: false},
}

var fields_create_presigned_mlflow_tracking_server_url = []leanruntime.Field{
	{Name: "ExpiresInSeconds", Flag: "expires-in-seconds", Type: "*int32", Required: false},
	{Name: "SessionExpirationDurationInSeconds", Flag: "session-expiration-duration-in-seconds", Type: "*int32", Required: false},
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
}

var fields_create_presigned_notebook_instance_url = []leanruntime.Field{
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
	{Name: "SessionExpirationDurationInSeconds", Flag: "session-expiration-duration-in-seconds", Type: "*int32", Required: false},
}

var fields_create_processing_job = []leanruntime.Field{
	{Name: "AppSpecification", Flag: "app-specification", Type: "*types.AppSpecification", Required: true},
	{Name: "Environment", Flag: "environment", Type: "map[string]string", Required: false},
	{Name: "ExperimentConfig", Flag: "experiment-config", Type: "*types.ExperimentConfig", Required: false},
	{Name: "NetworkConfig", Flag: "network-config", Type: "*types.NetworkConfig", Required: false},
	{Name: "ProcessingInputs", Flag: "processing-inputs", Type: "[]types.ProcessingInput", Required: false},
	{Name: "ProcessingJobName", Flag: "processing-job-name", Type: "*string", Required: true},
	{Name: "ProcessingOutputConfig", Flag: "processing-output-config", Type: "*types.ProcessingOutputConfig", Required: false},
	{Name: "ProcessingResources", Flag: "processing-resources", Type: "*types.ProcessingResources", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.ProcessingStoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "ProjectDescription", Flag: "project-description", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "ServiceCatalogProvisioningDetails", Flag: "service-catalog-provisioning-details", Type: "*types.ServiceCatalogProvisioningDetails", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateProviders", Flag: "template-providers", Type: "[]types.CreateTemplateProvider", Required: false},
}

var fields_create_space = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "OwnershipSettings", Flag: "ownership-settings", Type: "*types.OwnershipSettings", Required: false},
	{Name: "SpaceDisplayName", Flag: "space-display-name", Type: "*string", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
	{Name: "SpaceSettings", Flag: "space-settings", Type: "*types.SpaceSettings", Required: false},
	{Name: "SpaceSharingSettings", Flag: "space-sharing-settings", Type: "*types.SpaceSharingSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_studio_lifecycle_config = []leanruntime.Field{
	{Name: "StudioLifecycleConfigAppType", Flag: "studio-lifecycle-config-app-type", Type: "types.StudioLifecycleConfigAppType", Required: true},
	{Name: "StudioLifecycleConfigContent", Flag: "studio-lifecycle-config-content", Type: "*string", Required: true},
	{Name: "StudioLifecycleConfigName", Flag: "studio-lifecycle-config-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_training_job = []leanruntime.Field{
	{Name: "AlgorithmSpecification", Flag: "algorithm-specification", Type: "*types.AlgorithmSpecification", Required: false},
	{Name: "CheckpointConfig", Flag: "checkpoint-config", Type: "*types.CheckpointConfig", Required: false},
	{Name: "DebugHookConfig", Flag: "debug-hook-config", Type: "*types.DebugHookConfig", Required: false},
	{Name: "DebugRuleConfigurations", Flag: "debug-rule-configurations", Type: "[]types.DebugRuleConfiguration", Required: false},
	{Name: "EnableInterContainerTrafficEncryption", Flag: "enable-inter-container-traffic-encryption", Type: "*bool", Required: false},
	{Name: "EnableManagedSpotTraining", Flag: "enable-managed-spot-training", Type: "*bool", Required: false},
	{Name: "EnableNetworkIsolation", Flag: "enable-network-isolation", Type: "*bool", Required: false},
	{Name: "Environment", Flag: "environment", Type: "map[string]string", Required: false},
	{Name: "ExperimentConfig", Flag: "experiment-config", Type: "*types.ExperimentConfig", Required: false},
	{Name: "HyperParameters", Flag: "hyper-parameters", Type: "map[string]string", Required: false},
	{Name: "InfraCheckConfig", Flag: "infra-check-config", Type: "*types.InfraCheckConfig", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "[]types.Channel", Required: false},
	{Name: "MlflowConfig", Flag: "mlflow-config", Type: "*types.MlflowConfig", Required: false},
	{Name: "ModelPackageConfig", Flag: "model-package-config", Type: "*types.ModelPackageConfig", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "ProfilerConfig", Flag: "profiler-config", Type: "*types.ProfilerConfig", Required: false},
	{Name: "ProfilerRuleConfigurations", Flag: "profiler-rule-configurations", Type: "[]types.ProfilerRuleConfiguration", Required: false},
	{Name: "RemoteDebugConfig", Flag: "remote-debug-config", Type: "*types.RemoteDebugConfig", Required: false},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.ResourceConfig", Required: false},
	{Name: "RetryStrategy", Flag: "retry-strategy", Type: "*types.RetryStrategy", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "ServerlessJobConfig", Flag: "serverless-job-config", Type: "*types.ServerlessJobConfig", Required: false},
	{Name: "SessionChainingConfig", Flag: "session-chaining-config", Type: "*types.SessionChainingConfig", Required: false},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.StoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TensorBoardOutputConfig", Flag: "tensor-board-output-config", Type: "*types.TensorBoardOutputConfig", Required: false},
	{Name: "TrainingJobName", Flag: "training-job-name", Type: "*string", Required: true},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_training_plan = []leanruntime.Field{
	{Name: "SpareInstanceCountPerUltraServer", Flag: "spare-instance-count-per-ultra-server", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingPlanName", Flag: "training-plan-name", Type: "*string", Required: true},
	{Name: "TrainingPlanOfferingId", Flag: "training-plan-offering-id", Type: "*string", Required: true},
}

var fields_create_transform_job = []leanruntime.Field{
	{Name: "BatchStrategy", Flag: "batch-strategy", Type: "types.BatchStrategy", Required: false},
	{Name: "DataCaptureConfig", Flag: "data-capture-config", Type: "*types.BatchDataCaptureConfig", Required: false},
	{Name: "DataProcessing", Flag: "data-processing", Type: "*types.DataProcessing", Required: false},
	{Name: "Environment", Flag: "environment", Type: "map[string]string", Required: false},
	{Name: "ExperimentConfig", Flag: "experiment-config", Type: "*types.ExperimentConfig", Required: false},
	{Name: "MaxConcurrentTransforms", Flag: "max-concurrent-transforms", Type: "*int32", Required: false},
	{Name: "MaxPayloadInMB", Flag: "max-payload-in-mb", Type: "*int32", Required: false},
	{Name: "ModelClientConfig", Flag: "model-client-config", Type: "*types.ModelClientConfig", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransformInput", Flag: "transform-input", Type: "*types.TransformInput", Required: true},
	{Name: "TransformJobName", Flag: "transform-job-name", Type: "*string", Required: true},
	{Name: "TransformOutput", Flag: "transform-output", Type: "*types.TransformOutput", Required: true},
	{Name: "TransformResources", Flag: "transform-resources", Type: "*types.TransformResources", Required: true},
}

var fields_create_trial = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: true},
	{Name: "MetadataProperties", Flag: "metadata-properties", Type: "*types.MetadataProperties", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: true},
}

var fields_create_trial_component = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "InputArtifacts", Flag: "input-artifacts", Type: "map[string]types.TrialComponentArtifact", Required: false},
	{Name: "MetadataProperties", Flag: "metadata-properties", Type: "*types.MetadataProperties", Required: false},
	{Name: "OutputArtifacts", Flag: "output-artifacts", Type: "map[string]types.TrialComponentArtifact", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.TrialComponentParameterValue", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "*types.TrialComponentStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
}

var fields_create_user_profile = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SingleSignOnUserIdentifier", Flag: "single-sign-on-user-identifier", Type: "*string", Required: false},
	{Name: "SingleSignOnUserValue", Flag: "single-sign-on-user-value", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: true},
	{Name: "UserSettings", Flag: "user-settings", Type: "*types.UserSettings", Required: false},
}

var fields_create_workforce = []leanruntime.Field{
	{Name: "CognitoConfig", Flag: "cognito-config", Type: "*types.CognitoConfig", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.WorkforceIpAddressType", Required: false},
	{Name: "OidcConfig", Flag: "oidc-config", Type: "*types.OidcConfig", Required: false},
	{Name: "SourceIpConfig", Flag: "source-ip-config", Type: "*types.SourceIpConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WorkforceName", Flag: "workforce-name", Type: "*string", Required: true},
	{Name: "WorkforceVpcConfig", Flag: "workforce-vpc-config", Type: "*types.WorkforceVpcConfigRequest", Required: false},
}

var fields_create_workteam = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "MemberDefinitions", Flag: "member-definitions", Type: "[]types.MemberDefinition", Required: true},
	{Name: "NotificationConfiguration", Flag: "notification-configuration", Type: "*types.NotificationConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WorkerAccessConfiguration", Flag: "worker-access-configuration", Type: "*types.WorkerAccessConfiguration", Required: false},
	{Name: "WorkforceName", Flag: "workforce-name", Type: "*string", Required: false},
	{Name: "WorkteamName", Flag: "workteam-name", Type: "*string", Required: true},
}

var fields_delete_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
}

var fields_delete_algorithm = []leanruntime.Field{
	{Name: "AlgorithmName", Flag: "algorithm-name", Type: "*string", Required: true},
}

var fields_delete_app = []leanruntime.Field{
	{Name: "AppName", Flag: "app-name", Type: "*string", Required: true},
	{Name: "AppType", Flag: "app-type", Type: "types.AppType", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: false},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: false},
}

var fields_delete_app_image_config = []leanruntime.Field{
	{Name: "AppImageConfigName", Flag: "app-image-config-name", Type: "*string", Required: true},
}

var fields_delete_artifact = []leanruntime.Field{
	{Name: "ArtifactArn", Flag: "artifact-arn", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.ArtifactSource", Required: false},
}

var fields_delete_association = []leanruntime.Field{
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_delete_cluster_scheduler_config = []leanruntime.Field{
	{Name: "ClusterSchedulerConfigId", Flag: "cluster-scheduler-config-id", Type: "*string", Required: true},
}

var fields_delete_code_repository = []leanruntime.Field{
	{Name: "CodeRepositoryName", Flag: "code-repository-name", Type: "*string", Required: true},
}

var fields_delete_compilation_job = []leanruntime.Field{
	{Name: "CompilationJobName", Flag: "compilation-job-name", Type: "*string", Required: true},
}

var fields_delete_compute_quota = []leanruntime.Field{
	{Name: "ComputeQuotaId", Flag: "compute-quota-id", Type: "*string", Required: true},
}

var fields_delete_context = []leanruntime.Field{
	{Name: "ContextName", Flag: "context-name", Type: "*string", Required: true},
}

var fields_delete_data_quality_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_delete_device_fleet = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "RetentionPolicy", Flag: "retention-policy", Type: "*types.RetentionPolicy", Required: false},
}

var fields_delete_edge_deployment_plan = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
}

var fields_delete_edge_deployment_stage = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_delete_endpoint_config = []leanruntime.Field{
	{Name: "EndpointConfigName", Flag: "endpoint-config-name", Type: "*string", Required: true},
}

var fields_delete_experiment = []leanruntime.Field{
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: true},
}

var fields_delete_feature_group = []leanruntime.Field{
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
}

var fields_delete_flow_definition = []leanruntime.Field{
	{Name: "FlowDefinitionName", Flag: "flow-definition-name", Type: "*string", Required: true},
}

var fields_delete_hub = []leanruntime.Field{
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
}

var fields_delete_hub_content = []leanruntime.Field{
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubContentVersion", Flag: "hub-content-version", Type: "*string", Required: true},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
}

var fields_delete_hub_content_reference = []leanruntime.Field{
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
}

var fields_delete_human_task_ui = []leanruntime.Field{
	{Name: "HumanTaskUiName", Flag: "human-task-ui-name", Type: "*string", Required: true},
}

var fields_delete_hyper_parameter_tuning_job = []leanruntime.Field{
	{Name: "HyperParameterTuningJobName", Flag: "hyper-parameter-tuning-job-name", Type: "*string", Required: true},
}

var fields_delete_image = []leanruntime.Field{
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
}

var fields_delete_image_version = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*int32", Required: false},
}

var fields_delete_inference_component = []leanruntime.Field{
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: true},
}

var fields_delete_inference_experiment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_mlflow_app = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_mlflow_tracking_server = []leanruntime.Field{
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
}

var fields_delete_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_delete_model_bias_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_delete_model_card = []leanruntime.Field{
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
}

var fields_delete_model_explainability_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_delete_model_package = []leanruntime.Field{
	{Name: "ModelPackageName", Flag: "model-package-name", Type: "*string", Required: true},
}

var fields_delete_model_package_group = []leanruntime.Field{
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: true},
}

var fields_delete_model_package_group_policy = []leanruntime.Field{
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: true},
}

var fields_delete_model_quality_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_delete_monitoring_schedule = []leanruntime.Field{
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
}

var fields_delete_notebook_instance = []leanruntime.Field{
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
}

var fields_delete_notebook_instance_lifecycle_config = []leanruntime.Field{
	{Name: "NotebookInstanceLifecycleConfigName", Flag: "notebook-instance-lifecycle-config-name", Type: "*string", Required: true},
}

var fields_delete_optimization_job = []leanruntime.Field{
	{Name: "OptimizationJobName", Flag: "optimization-job-name", Type: "*string", Required: true},
}

var fields_delete_partner_app = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_pipeline = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_delete_processing_job = []leanruntime.Field{
	{Name: "ProcessingJobName", Flag: "processing-job-name", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
}

var fields_delete_space = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_delete_studio_lifecycle_config = []leanruntime.Field{
	{Name: "StudioLifecycleConfigName", Flag: "studio-lifecycle-config-name", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_delete_training_job = []leanruntime.Field{
	{Name: "TrainingJobName", Flag: "training-job-name", Type: "*string", Required: true},
}

var fields_delete_trial = []leanruntime.Field{
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: true},
}

var fields_delete_trial_component = []leanruntime.Field{
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
}

var fields_delete_user_profile = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: true},
}

var fields_delete_workforce = []leanruntime.Field{
	{Name: "WorkforceName", Flag: "workforce-name", Type: "*string", Required: true},
}

var fields_delete_workteam = []leanruntime.Field{
	{Name: "WorkteamName", Flag: "workteam-name", Type: "*string", Required: true},
}

var fields_deregister_devices = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "DeviceNames", Flag: "device-names", Type: "[]string", Required: true},
}

var fields_describe_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
}

var fields_describe_algorithm = []leanruntime.Field{
	{Name: "AlgorithmName", Flag: "algorithm-name", Type: "*string", Required: true},
}

var fields_describe_app = []leanruntime.Field{
	{Name: "AppName", Flag: "app-name", Type: "*string", Required: true},
	{Name: "AppType", Flag: "app-type", Type: "types.AppType", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: false},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: false},
}

var fields_describe_app_image_config = []leanruntime.Field{
	{Name: "AppImageConfigName", Flag: "app-image-config-name", Type: "*string", Required: true},
}

var fields_describe_artifact = []leanruntime.Field{
	{Name: "ArtifactArn", Flag: "artifact-arn", Type: "*string", Required: true},
}

var fields_describe_auto_ml_job = []leanruntime.Field{
	{Name: "AutoMLJobName", Flag: "auto-ml-job-name", Type: "*string", Required: true},
}

var fields_describe_auto_ml_job_v2 = []leanruntime.Field{
	{Name: "AutoMLJobName", Flag: "auto-ml-job-name", Type: "*string", Required: true},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_describe_cluster_event = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
}

var fields_describe_cluster_node = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: false},
	{Name: "NodeLogicalId", Flag: "node-logical-id", Type: "*string", Required: false},
}

var fields_describe_cluster_scheduler_config = []leanruntime.Field{
	{Name: "ClusterSchedulerConfigId", Flag: "cluster-scheduler-config-id", Type: "*string", Required: true},
	{Name: "ClusterSchedulerConfigVersion", Flag: "cluster-scheduler-config-version", Type: "*int32", Required: false},
}

var fields_describe_code_repository = []leanruntime.Field{
	{Name: "CodeRepositoryName", Flag: "code-repository-name", Type: "*string", Required: true},
}

var fields_describe_compilation_job = []leanruntime.Field{
	{Name: "CompilationJobName", Flag: "compilation-job-name", Type: "*string", Required: true},
}

var fields_describe_compute_quota = []leanruntime.Field{
	{Name: "ComputeQuotaId", Flag: "compute-quota-id", Type: "*string", Required: true},
	{Name: "ComputeQuotaVersion", Flag: "compute-quota-version", Type: "*int32", Required: false},
}

var fields_describe_context = []leanruntime.Field{
	{Name: "ContextName", Flag: "context-name", Type: "*string", Required: true},
}

var fields_describe_data_quality_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_describe_device = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_device_fleet = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
}

var fields_describe_domain = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_describe_edge_deployment_plan = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_edge_packaging_job = []leanruntime.Field{
	{Name: "EdgePackagingJobName", Flag: "edge-packaging-job-name", Type: "*string", Required: true},
}

var fields_describe_endpoint = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_describe_endpoint_config = []leanruntime.Field{
	{Name: "EndpointConfigName", Flag: "endpoint-config-name", Type: "*string", Required: true},
}

var fields_describe_experiment = []leanruntime.Field{
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: true},
}

var fields_describe_feature_group = []leanruntime.Field{
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_feature_metadata = []leanruntime.Field{
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: true},
}

var fields_describe_flow_definition = []leanruntime.Field{
	{Name: "FlowDefinitionName", Flag: "flow-definition-name", Type: "*string", Required: true},
}

var fields_describe_hub = []leanruntime.Field{
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
}

var fields_describe_hub_content = []leanruntime.Field{
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubContentVersion", Flag: "hub-content-version", Type: "*string", Required: false},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
}

var fields_describe_human_task_ui = []leanruntime.Field{
	{Name: "HumanTaskUiName", Flag: "human-task-ui-name", Type: "*string", Required: true},
}

var fields_describe_hyper_parameter_tuning_job = []leanruntime.Field{
	{Name: "HyperParameterTuningJobName", Flag: "hyper-parameter-tuning-job-name", Type: "*string", Required: true},
}

var fields_describe_image = []leanruntime.Field{
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
}

var fields_describe_image_version = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*int32", Required: false},
}

var fields_describe_inference_component = []leanruntime.Field{
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: true},
}

var fields_describe_inference_experiment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_inference_recommendations_job = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
}

var fields_describe_labeling_job = []leanruntime.Field{
	{Name: "LabelingJobName", Flag: "labeling-job-name", Type: "*string", Required: true},
}

var fields_describe_lineage_group = []leanruntime.Field{
	{Name: "LineageGroupName", Flag: "lineage-group-name", Type: "*string", Required: true},
}

var fields_describe_mlflow_app = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_describe_mlflow_tracking_server = []leanruntime.Field{
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
}

var fields_describe_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_describe_model_bias_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_describe_model_card = []leanruntime.Field{
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
	{Name: "ModelCardVersion", Flag: "model-card-version", Type: "*int32", Required: false},
}

var fields_describe_model_card_export_job = []leanruntime.Field{
	{Name: "ModelCardExportJobArn", Flag: "model-card-export-job-arn", Type: "*string", Required: true},
}

var fields_describe_model_explainability_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_describe_model_package = []leanruntime.Field{
	{Name: "ModelPackageName", Flag: "model-package-name", Type: "*string", Required: true},
}

var fields_describe_model_package_group = []leanruntime.Field{
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: true},
}

var fields_describe_model_quality_job_definition = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
}

var fields_describe_monitoring_schedule = []leanruntime.Field{
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
}

var fields_describe_notebook_instance = []leanruntime.Field{
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
}

var fields_describe_notebook_instance_lifecycle_config = []leanruntime.Field{
	{Name: "NotebookInstanceLifecycleConfigName", Flag: "notebook-instance-lifecycle-config-name", Type: "*string", Required: true},
}

var fields_describe_optimization_job = []leanruntime.Field{
	{Name: "OptimizationJobName", Flag: "optimization-job-name", Type: "*string", Required: true},
}

var fields_describe_partner_app = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "IncludeAvailableUpgrade", Flag: "include-available-upgrade", Type: "*bool", Required: false},
}

var fields_describe_pipeline = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "PipelineVersionId", Flag: "pipeline-version-id", Type: "*int64", Required: false},
}

var fields_describe_pipeline_definition_for_execution = []leanruntime.Field{
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: true},
}

var fields_describe_pipeline_execution = []leanruntime.Field{
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: true},
}

var fields_describe_processing_job = []leanruntime.Field{
	{Name: "ProcessingJobName", Flag: "processing-job-name", Type: "*string", Required: true},
}

var fields_describe_project = []leanruntime.Field{
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
}

var fields_describe_reserved_capacity = []leanruntime.Field{
	{Name: "ReservedCapacityArn", Flag: "reserved-capacity-arn", Type: "*string", Required: true},
}

var fields_describe_space = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_describe_studio_lifecycle_config = []leanruntime.Field{
	{Name: "StudioLifecycleConfigName", Flag: "studio-lifecycle-config-name", Type: "*string", Required: true},
}

var fields_describe_subscribed_workteam = []leanruntime.Field{
	{Name: "WorkteamArn", Flag: "workteam-arn", Type: "*string", Required: true},
}

var fields_describe_training_job = []leanruntime.Field{
	{Name: "TrainingJobName", Flag: "training-job-name", Type: "*string", Required: true},
}

var fields_describe_training_plan = []leanruntime.Field{
	{Name: "TrainingPlanName", Flag: "training-plan-name", Type: "*string", Required: true},
}

var fields_describe_transform_job = []leanruntime.Field{
	{Name: "TransformJobName", Flag: "transform-job-name", Type: "*string", Required: true},
}

var fields_describe_trial = []leanruntime.Field{
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: true},
}

var fields_describe_trial_component = []leanruntime.Field{
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
}

var fields_describe_user_profile = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: true},
}

var fields_describe_workforce = []leanruntime.Field{
	{Name: "WorkforceName", Flag: "workforce-name", Type: "*string", Required: true},
}

var fields_describe_workteam = []leanruntime.Field{
	{Name: "WorkteamName", Flag: "workteam-name", Type: "*string", Required: true},
}

var fields_detach_cluster_node_volume = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_disable_sagemaker_servicecatalog_portfolio = []leanruntime.Field{}

var fields_disassociate_trial_component = []leanruntime.Field{
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: true},
}

var fields_enable_sagemaker_servicecatalog_portfolio = []leanruntime.Field{}

var fields_get_device_fleet_report = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
}

var fields_get_lineage_group_policy = []leanruntime.Field{
	{Name: "LineageGroupName", Flag: "lineage-group-name", Type: "*string", Required: true},
}

var fields_get_model_package_group_policy = []leanruntime.Field{
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: true},
}

var fields_get_sagemaker_servicecatalog_portfolio_status = []leanruntime.Field{}

var fields_get_scaling_configuration_recommendation = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "InferenceRecommendationsJobName", Flag: "inference-recommendations-job-name", Type: "*string", Required: true},
	{Name: "RecommendationId", Flag: "recommendation-id", Type: "*string", Required: false},
	{Name: "ScalingPolicyObjective", Flag: "scaling-policy-objective", Type: "*types.ScalingPolicyObjective", Required: false},
	{Name: "TargetCpuUtilizationPerCore", Flag: "target-cpu-utilization-per-core", Type: "*int32", Required: false},
}

var fields_get_search_suggestions = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "types.ResourceType", Required: true},
	{Name: "SuggestionQuery", Flag: "suggestion-query", Type: "*types.SuggestionQuery", Required: false},
}

var fields_import_hub_content = []leanruntime.Field{
	{Name: "DocumentSchemaVersion", Flag: "document-schema-version", Type: "*string", Required: true},
	{Name: "HubContentDescription", Flag: "hub-content-description", Type: "*string", Required: false},
	{Name: "HubContentDisplayName", Flag: "hub-content-display-name", Type: "*string", Required: false},
	{Name: "HubContentDocument", Flag: "hub-content-document", Type: "*string", Required: true},
	{Name: "HubContentMarkdown", Flag: "hub-content-markdown", Type: "*string", Required: false},
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentSearchKeywords", Flag: "hub-content-search-keywords", Type: "[]string", Required: false},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubContentVersion", Flag: "hub-content-version", Type: "*string", Required: false},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "SupportStatus", Flag: "support-status", Type: "types.HubContentSupportStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_actions = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "*string", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortActionsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SourceUri", Flag: "source-uri", Type: "*string", Required: false},
}

var fields_list_algorithms = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.AlgorithmSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_aliases = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*int32", Required: false},
}

var fields_list_app_image_configs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModifiedTimeAfter", Flag: "modified-time-after", Type: "*time.Time", Required: false},
	{Name: "ModifiedTimeBefore", Flag: "modified-time-before", Type: "*time.Time", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.AppImageConfigSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_apps = []leanruntime.Field{
	{Name: "DomainIdEquals", Flag: "domain-id-equals", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.AppSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SpaceNameEquals", Flag: "space-name-equals", Type: "*string", Required: false},
	{Name: "UserProfileNameEquals", Flag: "user-profile-name-equals", Type: "*string", Required: false},
}

var fields_list_artifacts = []leanruntime.Field{
	{Name: "ArtifactType", Flag: "artifact-type", Type: "*string", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortArtifactsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SourceUri", Flag: "source-uri", Type: "*string", Required: false},
}

var fields_list_associations = []leanruntime.Field{
	{Name: "AssociationType", Flag: "association-type", Type: "types.AssociationEdgeType", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "DestinationType", Flag: "destination-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortAssociationsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
}

var fields_list_auto_ml_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.AutoMLSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.AutoMLSortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.AutoMLJobStatus", Required: false},
}

var fields_list_candidates_for_auto_ml_job = []leanruntime.Field{
	{Name: "AutoMLJobName", Flag: "auto-ml-job-name", Type: "*string", Required: true},
	{Name: "CandidateNameEquals", Flag: "candidate-name-equals", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CandidateSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.AutoMLSortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.CandidateStatus", Required: false},
}

var fields_list_cluster_events = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EventTimeAfter", Flag: "event-time-after", Type: "*time.Time", Required: false},
	{Name: "EventTimeBefore", Flag: "event-time-before", Type: "*time.Time", Required: false},
	{Name: "InstanceGroupName", Flag: "instance-group-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ClusterEventResourceType", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.EventSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_cluster_nodes = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "IncludeNodeLogicalIds", Flag: "include-node-logical-ids", Type: "*bool", Required: false},
	{Name: "InstanceGroupNameContains", Flag: "instance-group-name-contains", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ClusterSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_cluster_scheduler_configs = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortClusterSchedulerConfigBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.SchedulerResourceStatus", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ClusterSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "TrainingPlanArn", Flag: "training-plan-arn", Type: "*string", Required: false},
}

var fields_list_code_repositories = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CodeRepositorySortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.CodeRepositorySortOrder", Required: false},
}

var fields_list_compilation_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListCompilationJobsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.CompilationJobStatus", Required: false},
}

var fields_list_compute_quotas = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortQuotaBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.SchedulerResourceStatus", Required: false},
}

var fields_list_contexts = []leanruntime.Field{
	{Name: "ContextType", Flag: "context-type", Type: "*string", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortContextsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SourceUri", Flag: "source-uri", Type: "*string", Required: false},
}

var fields_list_data_quality_job_definitions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringJobDefinitionSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_device_fleets = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListDeviceFleetsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_devices = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: false},
	{Name: "LatestHeartbeatAfter", Flag: "latest-heartbeat-after", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_edge_deployment_plans = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "DeviceFleetNameContains", Flag: "device-fleet-name-contains", Type: "*string", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListEdgeDeploymentPlansSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_edge_packaging_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelNameContains", Flag: "model-name-contains", Type: "*string", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListEdgePackagingJobsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.EdgePackagingJobStatus", Required: false},
}

var fields_list_endpoint_configs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.EndpointConfigSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.OrderKey", Required: false},
}

var fields_list_endpoints = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.EndpointSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.OrderKey", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.EndpointStatus", Required: false},
}

var fields_list_experiments = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortExperimentsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_feature_groups = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "FeatureGroupStatusEquals", Flag: "feature-group-status-equals", Type: "types.FeatureGroupStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OfflineStoreStatusEquals", Flag: "offline-store-status-equals", Type: "types.OfflineStoreStatusValue", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.FeatureGroupSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.FeatureGroupSortOrder", Required: false},
}

var fields_list_flow_definitions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_hub_content_versions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MaxSchemaVersion", Flag: "max-schema-version", Type: "*string", Required: false},
	{Name: "MinVersion", Flag: "min-version", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.HubContentSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_hub_contents = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MaxSchemaVersion", Flag: "max-schema-version", Type: "*string", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.HubContentSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_hubs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.HubSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_human_task_uis = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_hyper_parameter_tuning_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.HyperParameterTuningJobSortByOptions", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.HyperParameterTuningJobStatus", Required: false},
}

var fields_list_image_versions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ImageVersionSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ImageVersionSortOrder", Required: false},
}

var fields_list_images = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ImageSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ImageSortOrder", Required: false},
}

var fields_list_inference_components = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointNameEquals", Flag: "endpoint-name-equals", Type: "*string", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.InferenceComponentSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.OrderKey", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.InferenceComponentStatus", Required: false},
	{Name: "VariantNameEquals", Flag: "variant-name-equals", Type: "*string", Required: false},
}

var fields_list_inference_experiments = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortInferenceExperimentsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.InferenceExperimentStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "types.InferenceExperimentType", Required: false},
}

var fields_list_inference_recommendations_job_steps = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RecommendationJobStatus", Required: false},
	{Name: "StepType", Flag: "step-type", Type: "types.RecommendationStepType", Required: false},
}

var fields_list_inference_recommendations_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelNameEquals", Flag: "model-name-equals", Type: "*string", Required: false},
	{Name: "ModelPackageVersionArnEquals", Flag: "model-package-version-arn-equals", Type: "*string", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListInferenceRecommendationsJobsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.RecommendationJobStatus", Required: false},
}

var fields_list_labeling_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.LabelingJobStatus", Required: false},
}

var fields_list_labeling_jobs_for_workteam = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "JobReferenceCodeContains", Flag: "job-reference-code-contains", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListLabelingJobsForWorkteamSortByOptions", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "WorkteamArn", Flag: "workteam-arn", Type: "*string", Required: true},
}

var fields_list_lineage_groups = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortLineageGroupsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_mlflow_apps = []leanruntime.Field{
	{Name: "AccountDefaultStatus", Flag: "account-default-status", Type: "types.AccountDefaultStatus", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "DefaultForDomainId", Flag: "default-for-domain-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MlflowVersion", Flag: "mlflow-version", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortMlflowAppBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MlflowAppStatus", Required: false},
}

var fields_list_mlflow_tracking_servers = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MlflowVersion", Flag: "mlflow-version", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortTrackingServerBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "TrackingServerStatus", Flag: "tracking-server-status", Type: "types.TrackingServerStatus", Required: false},
}

var fields_list_model_bias_job_definitions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringJobDefinitionSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_model_card_export_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelCardExportJobNameContains", Flag: "model-card-export-job-name-contains", Type: "*string", Required: false},
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
	{Name: "ModelCardVersion", Flag: "model-card-version", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ModelCardExportJobSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ModelCardExportJobSortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ModelCardExportJobStatus", Required: false},
}

var fields_list_model_card_versions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
	{Name: "ModelCardStatus", Flag: "model-card-status", Type: "types.ModelCardStatus", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ModelCardVersionSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ModelCardSortOrder", Required: false},
}

var fields_list_model_cards = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelCardStatus", Flag: "model-card-status", Type: "types.ModelCardStatus", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ModelCardSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ModelCardSortOrder", Required: false},
}

var fields_list_model_explainability_job_definitions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringJobDefinitionSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_model_metadata = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.ModelMetadataSearchExpression", Required: false},
}

var fields_list_model_package_groups = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "CrossAccountFilterOption", Flag: "cross-account-filter-option", Type: "types.CrossAccountFilterOption", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ModelPackageGroupSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_model_packages = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelApprovalStatus", Flag: "model-approval-status", Type: "types.ModelApprovalStatus", Required: false},
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: false},
	{Name: "ModelPackageType", Flag: "model-package-type", Type: "types.ModelPackageType", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ModelPackageSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_model_quality_job_definitions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringJobDefinitionSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_models = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ModelSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.OrderKey", Required: false},
}

var fields_list_monitoring_alert_history = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitoringAlertName", Flag: "monitoring-alert-name", Type: "*string", Required: false},
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringAlertHistorySortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.MonitoringAlertStatus", Required: false},
}

var fields_list_monitoring_alerts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_monitoring_executions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitoringJobDefinitionName", Flag: "monitoring-job-definition-name", Type: "*string", Required: false},
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: false},
	{Name: "MonitoringTypeEquals", Flag: "monitoring-type-equals", Type: "types.MonitoringType", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScheduledTimeAfter", Flag: "scheduled-time-after", Type: "*time.Time", Required: false},
	{Name: "ScheduledTimeBefore", Flag: "scheduled-time-before", Type: "*time.Time", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringExecutionSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ExecutionStatus", Required: false},
}

var fields_list_monitoring_schedules = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitoringJobDefinitionName", Flag: "monitoring-job-definition-name", Type: "*string", Required: false},
	{Name: "MonitoringTypeEquals", Flag: "monitoring-type-equals", Type: "types.MonitoringType", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.MonitoringScheduleSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ScheduleStatus", Required: false},
}

var fields_list_notebook_instance_lifecycle_configs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.NotebookInstanceLifecycleConfigSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.NotebookInstanceLifecycleConfigSortOrder", Required: false},
}

var fields_list_notebook_instances = []leanruntime.Field{
	{Name: "AdditionalCodeRepositoryEquals", Flag: "additional-code-repository-equals", Type: "*string", Required: false},
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "DefaultCodeRepositoryContains", Flag: "default-code-repository-contains", Type: "*string", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotebookInstanceLifecycleConfigNameContains", Flag: "notebook-instance-lifecycle-config-name-contains", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.NotebookInstanceSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.NotebookInstanceSortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.NotebookInstanceStatus", Required: false},
}

var fields_list_optimization_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OptimizationContains", Flag: "optimization-contains", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListOptimizationJobsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.OptimizationJobStatus", Required: false},
}

var fields_list_partner_apps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pipeline_execution_steps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_pipeline_executions = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortPipelineExecutionsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_pipeline_parameters_for_execution = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: true},
}

var fields_list_pipeline_versions = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_pipelines = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineNamePrefix", Flag: "pipeline-name-prefix", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortPipelinesBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_processing_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ProcessingJobStatus", Required: false},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ProjectSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ProjectSortOrder", Required: false},
}

var fields_list_resource_catalogs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ResourceCatalogSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.ResourceCatalogSortOrder", Required: false},
}

var fields_list_spaces = []leanruntime.Field{
	{Name: "DomainIdEquals", Flag: "domain-id-equals", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SpaceSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SpaceNameContains", Flag: "space-name-contains", Type: "*string", Required: false},
}

var fields_list_stage_devices = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "ExcludeDevicesDeployedInOtherStage", Flag: "exclude-devices-deployed-in-other-stage", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_list_studio_lifecycle_configs = []leanruntime.Field{
	{Name: "AppTypeEquals", Flag: "app-type-equals", Type: "types.StudioLifecycleConfigAppType", Required: false},
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModifiedTimeAfter", Flag: "modified-time-after", Type: "*time.Time", Required: false},
	{Name: "ModifiedTimeBefore", Flag: "modified-time-before", Type: "*time.Time", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.StudioLifecycleConfigSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_subscribed_workteams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_training_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.TrainingJobStatus", Required: false},
	{Name: "TrainingPlanArnEquals", Flag: "training-plan-arn-equals", Type: "*string", Required: false},
	{Name: "WarmPoolStatusEquals", Flag: "warm-pool-status-equals", Type: "types.WarmPoolResourceStatus", Required: false},
}

var fields_list_training_jobs_for_hyper_parameter_tuning_job = []leanruntime.Field{
	{Name: "HyperParameterTuningJobName", Flag: "hyper-parameter-tuning-job-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.TrainingJobSortByOptions", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.TrainingJobStatus", Required: false},
}

var fields_list_training_plans = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.TrainingPlanFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.TrainingPlanSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.TrainingPlanSortOrder", Required: false},
	{Name: "StartTimeAfter", Flag: "start-time-after", Type: "*time.Time", Required: false},
	{Name: "StartTimeBefore", Flag: "start-time-before", Type: "*time.Time", Required: false},
}

var fields_list_transform_jobs = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeAfter", Flag: "last-modified-time-after", Type: "*time.Time", Required: false},
	{Name: "LastModifiedTimeBefore", Flag: "last-modified-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.TransformJobStatus", Required: false},
}

var fields_list_trial_components = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortTrialComponentsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: false},
}

var fields_list_trials = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortTrialsBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: false},
}

var fields_list_ultra_servers_by_reserved_capacity = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReservedCapacityArn", Flag: "reserved-capacity-arn", Type: "*string", Required: true},
}

var fields_list_user_profiles = []leanruntime.Field{
	{Name: "DomainIdEquals", Flag: "domain-id-equals", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.UserProfileSortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "UserProfileNameContains", Flag: "user-profile-name-contains", Type: "*string", Required: false},
}

var fields_list_workforces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListWorkforcesSortByOptions", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_workteams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListWorkteamsSortByOptions", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_put_model_package_group_policy = []leanruntime.Field{
	{Name: "ModelPackageGroupName", Flag: "model-package-group-name", Type: "*string", Required: true},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
}

var fields_query_lineage = []leanruntime.Field{
	{Name: "Direction", Flag: "direction", Type: "types.Direction", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.QueryFilters", Required: false},
	{Name: "IncludeEdges", Flag: "include-edges", Type: "*bool", Required: false},
	{Name: "MaxDepth", Flag: "max-depth", Type: "*int32", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartArns", Flag: "start-arns", Type: "[]string", Required: false},
}

var fields_register_devices = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "Devices", Flag: "devices", Type: "[]types.Device", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_render_ui_template = []leanruntime.Field{
	{Name: "HumanTaskUiArn", Flag: "human-task-ui-arn", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Task", Flag: "task", Type: "*types.RenderableTask", Required: true},
	{Name: "UiTemplate", Flag: "ui-template", Type: "*types.UiTemplate", Required: false},
}

var fields_retry_pipeline_execution = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ParallelismConfiguration", Flag: "parallelism-configuration", Type: "*types.ParallelismConfiguration", Required: false},
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: true},
}

var fields_search = []leanruntime.Field{
	{Name: "CrossAccountFilterOption", Flag: "cross-account-filter-option", Type: "types.CrossAccountFilterOption", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resource", Flag: "resource", Type: "types.ResourceType", Required: true},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.SearchExpression", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SearchSortOrder", Required: false},
	{Name: "VisibilityConditions", Flag: "visibility-conditions", Type: "[]types.VisibilityConditions", Required: false},
}

var fields_search_training_plan_offerings = []leanruntime.Field{
	{Name: "DurationHours", Flag: "duration-hours", Type: "*int64", Required: false},
	{Name: "EndTimeBefore", Flag: "end-time-before", Type: "*time.Time", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.ReservedCapacityInstanceType", Required: false},
	{Name: "StartTimeAfter", Flag: "start-time-after", Type: "*time.Time", Required: false},
	{Name: "TargetResources", Flag: "target-resources", Type: "[]types.SageMakerResourceName", Required: false},
	{Name: "UltraServerCount", Flag: "ultra-server-count", Type: "*int32", Required: false},
	{Name: "UltraServerType", Flag: "ultra-server-type", Type: "*string", Required: false},
}

var fields_send_pipeline_execution_step_failure = []leanruntime.Field{
	{Name: "CallbackToken", Flag: "callback-token", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FailureReason", Flag: "failure-reason", Type: "*string", Required: false},
}

var fields_send_pipeline_execution_step_success = []leanruntime.Field{
	{Name: "CallbackToken", Flag: "callback-token", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "OutputParameters", Flag: "output-parameters", Type: "[]types.OutputParameter", Required: false},
}

var fields_start_edge_deployment_stage = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_start_inference_experiment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_mlflow_tracking_server = []leanruntime.Field{
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
}

var fields_start_monitoring_schedule = []leanruntime.Field{
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
}

var fields_start_notebook_instance = []leanruntime.Field{
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
}

var fields_start_pipeline_execution = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "MlflowExperimentName", Flag: "mlflow-experiment-name", Type: "*string", Required: false},
	{Name: "ParallelismConfiguration", Flag: "parallelism-configuration", Type: "*types.ParallelismConfiguration", Required: false},
	{Name: "PipelineExecutionDescription", Flag: "pipeline-execution-description", Type: "*string", Required: false},
	{Name: "PipelineExecutionDisplayName", Flag: "pipeline-execution-display-name", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "PipelineParameters", Flag: "pipeline-parameters", Type: "[]types.Parameter", Required: false},
	{Name: "PipelineVersionId", Flag: "pipeline-version-id", Type: "*int64", Required: false},
	{Name: "SelectiveExecutionConfig", Flag: "selective-execution-config", Type: "*types.SelectiveExecutionConfig", Required: false},
}

var fields_start_session = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_stop_auto_ml_job = []leanruntime.Field{
	{Name: "AutoMLJobName", Flag: "auto-ml-job-name", Type: "*string", Required: true},
}

var fields_stop_compilation_job = []leanruntime.Field{
	{Name: "CompilationJobName", Flag: "compilation-job-name", Type: "*string", Required: true},
}

var fields_stop_edge_deployment_stage = []leanruntime.Field{
	{Name: "EdgeDeploymentPlanName", Flag: "edge-deployment-plan-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_stop_edge_packaging_job = []leanruntime.Field{
	{Name: "EdgePackagingJobName", Flag: "edge-packaging-job-name", Type: "*string", Required: true},
}

var fields_stop_hyper_parameter_tuning_job = []leanruntime.Field{
	{Name: "HyperParameterTuningJobName", Flag: "hyper-parameter-tuning-job-name", Type: "*string", Required: true},
}

var fields_stop_inference_experiment = []leanruntime.Field{
	{Name: "DesiredModelVariants", Flag: "desired-model-variants", Type: "[]types.ModelVariantConfig", Required: false},
	{Name: "DesiredState", Flag: "desired-state", Type: "types.InferenceExperimentStopDesiredState", Required: false},
	{Name: "ModelVariantActions", Flag: "model-variant-actions", Type: "map[string]types.ModelVariantAction", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_stop_inference_recommendations_job = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
}

var fields_stop_labeling_job = []leanruntime.Field{
	{Name: "LabelingJobName", Flag: "labeling-job-name", Type: "*string", Required: true},
}

var fields_stop_mlflow_tracking_server = []leanruntime.Field{
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
}

var fields_stop_monitoring_schedule = []leanruntime.Field{
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
}

var fields_stop_notebook_instance = []leanruntime.Field{
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
}

var fields_stop_optimization_job = []leanruntime.Field{
	{Name: "OptimizationJobName", Flag: "optimization-job-name", Type: "*string", Required: true},
}

var fields_stop_pipeline_execution = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: true},
}

var fields_stop_processing_job = []leanruntime.Field{
	{Name: "ProcessingJobName", Flag: "processing-job-name", Type: "*string", Required: true},
}

var fields_stop_training_job = []leanruntime.Field{
	{Name: "TrainingJobName", Flag: "training-job-name", Type: "*string", Required: true},
}

var fields_stop_transform_job = []leanruntime.Field{
	{Name: "TransformJobName", Flag: "transform-job-name", Type: "*string", Required: true},
}

var fields_update_action = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Properties", Flag: "properties", Type: "map[string]string", Required: false},
	{Name: "PropertiesToRemove", Flag: "properties-to-remove", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ActionStatus", Required: false},
}

var fields_update_app_image_config = []leanruntime.Field{
	{Name: "AppImageConfigName", Flag: "app-image-config-name", Type: "*string", Required: true},
	{Name: "CodeEditorAppImageConfig", Flag: "code-editor-app-image-config", Type: "*types.CodeEditorAppImageConfig", Required: false},
	{Name: "JupyterLabAppImageConfig", Flag: "jupyter-lab-app-image-config", Type: "*types.JupyterLabAppImageConfig", Required: false},
	{Name: "KernelGatewayImageConfig", Flag: "kernel-gateway-image-config", Type: "*types.KernelGatewayImageConfig", Required: false},
}

var fields_update_artifact = []leanruntime.Field{
	{Name: "ArtifactArn", Flag: "artifact-arn", Type: "*string", Required: true},
	{Name: "ArtifactName", Flag: "artifact-name", Type: "*string", Required: false},
	{Name: "Properties", Flag: "properties", Type: "map[string]string", Required: false},
	{Name: "PropertiesToRemove", Flag: "properties-to-remove", Type: "[]string", Required: false},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "AutoScaling", Flag: "auto-scaling", Type: "*types.ClusterAutoScalingConfig", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ClusterRole", Flag: "cluster-role", Type: "*string", Required: false},
	{Name: "InstanceGroups", Flag: "instance-groups", Type: "[]types.ClusterInstanceGroupSpecification", Required: false},
	{Name: "InstanceGroupsToDelete", Flag: "instance-groups-to-delete", Type: "[]string", Required: false},
	{Name: "NodeProvisioningMode", Flag: "node-provisioning-mode", Type: "types.ClusterNodeProvisioningMode", Required: false},
	{Name: "NodeRecovery", Flag: "node-recovery", Type: "types.ClusterNodeRecovery", Required: false},
	{Name: "Orchestrator", Flag: "orchestrator", Type: "*types.ClusterOrchestrator", Required: false},
	{Name: "RestrictedInstanceGroups", Flag: "restricted-instance-groups", Type: "[]types.ClusterRestrictedInstanceGroupSpecification", Required: false},
	{Name: "TieredStorageConfig", Flag: "tiered-storage-config", Type: "*types.ClusterTieredStorageConfig", Required: false},
}

var fields_update_cluster_scheduler_config = []leanruntime.Field{
	{Name: "ClusterSchedulerConfigId", Flag: "cluster-scheduler-config-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SchedulerConfig", Flag: "scheduler-config", Type: "*types.SchedulerConfig", Required: false},
	{Name: "TargetVersion", Flag: "target-version", Type: "*int32", Required: true},
}

var fields_update_cluster_software = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "DeploymentConfig", Flag: "deployment-config", Type: "*types.DeploymentConfiguration", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "InstanceGroups", Flag: "instance-groups", Type: "[]types.UpdateClusterSoftwareInstanceGroupSpecification", Required: false},
}

var fields_update_code_repository = []leanruntime.Field{
	{Name: "CodeRepositoryName", Flag: "code-repository-name", Type: "*string", Required: true},
	{Name: "GitConfig", Flag: "git-config", Type: "*types.GitConfigForUpdate", Required: false},
}

var fields_update_compute_quota = []leanruntime.Field{
	{Name: "ActivationState", Flag: "activation-state", Type: "types.ActivationState", Required: false},
	{Name: "ComputeQuotaConfig", Flag: "compute-quota-config", Type: "*types.ComputeQuotaConfig", Required: false},
	{Name: "ComputeQuotaId", Flag: "compute-quota-id", Type: "*string", Required: true},
	{Name: "ComputeQuotaTarget", Flag: "compute-quota-target", Type: "*types.ComputeQuotaTarget", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "TargetVersion", Flag: "target-version", Type: "*int32", Required: true},
}

var fields_update_context = []leanruntime.Field{
	{Name: "ContextName", Flag: "context-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Properties", Flag: "properties", Type: "map[string]string", Required: false},
	{Name: "PropertiesToRemove", Flag: "properties-to-remove", Type: "[]string", Required: false},
}

var fields_update_device_fleet = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "EnableIotRoleAlias", Flag: "enable-iot-role-alias", Type: "*bool", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.EdgeOutputConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_devices = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "Devices", Flag: "devices", Type: "[]types.Device", Required: true},
}

var fields_update_domain = []leanruntime.Field{
	{Name: "AppNetworkAccessType", Flag: "app-network-access-type", Type: "types.AppNetworkAccessType", Required: false},
	{Name: "AppSecurityGroupManagement", Flag: "app-security-group-management", Type: "types.AppSecurityGroupManagement", Required: false},
	{Name: "DefaultSpaceSettings", Flag: "default-space-settings", Type: "*types.DefaultSpaceSettings", Required: false},
	{Name: "DefaultUserSettings", Flag: "default-user-settings", Type: "*types.UserSettings", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "DomainSettingsForUpdate", Flag: "domain-settings-for-update", Type: "*types.DomainSettingsForUpdate", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "TagPropagation", Flag: "tag-propagation", Type: "types.TagPropagation", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_update_endpoint = []leanruntime.Field{
	{Name: "DeploymentConfig", Flag: "deployment-config", Type: "*types.DeploymentConfig", Required: false},
	{Name: "EndpointConfigName", Flag: "endpoint-config-name", Type: "*string", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "ExcludeRetainedVariantProperties", Flag: "exclude-retained-variant-properties", Type: "[]types.VariantProperty", Required: false},
	{Name: "RetainAllVariantProperties", Flag: "retain-all-variant-properties", Type: "*bool", Required: false},
	{Name: "RetainDeploymentConfig", Flag: "retain-deployment-config", Type: "*bool", Required: false},
}

var fields_update_endpoint_weights_and_capacities = []leanruntime.Field{
	{Name: "DesiredWeightsAndCapacities", Flag: "desired-weights-and-capacities", Type: "[]types.DesiredWeightAndCapacity", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_update_experiment = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "ExperimentName", Flag: "experiment-name", Type: "*string", Required: true},
}

var fields_update_feature_group = []leanruntime.Field{
	{Name: "FeatureAdditions", Flag: "feature-additions", Type: "[]types.FeatureDefinition", Required: false},
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "OnlineStoreConfig", Flag: "online-store-config", Type: "*types.OnlineStoreConfigUpdate", Required: false},
	{Name: "ThroughputConfig", Flag: "throughput-config", Type: "*types.ThroughputConfigUpdate", Required: false},
}

var fields_update_feature_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FeatureGroupName", Flag: "feature-group-name", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: true},
	{Name: "ParameterAdditions", Flag: "parameter-additions", Type: "[]types.FeatureParameter", Required: false},
	{Name: "ParameterRemovals", Flag: "parameter-removals", Type: "[]string", Required: false},
}

var fields_update_hub = []leanruntime.Field{
	{Name: "HubDescription", Flag: "hub-description", Type: "*string", Required: false},
	{Name: "HubDisplayName", Flag: "hub-display-name", Type: "*string", Required: false},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "HubSearchKeywords", Flag: "hub-search-keywords", Type: "[]string", Required: false},
}

var fields_update_hub_content = []leanruntime.Field{
	{Name: "HubContentDescription", Flag: "hub-content-description", Type: "*string", Required: false},
	{Name: "HubContentDisplayName", Flag: "hub-content-display-name", Type: "*string", Required: false},
	{Name: "HubContentMarkdown", Flag: "hub-content-markdown", Type: "*string", Required: false},
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentSearchKeywords", Flag: "hub-content-search-keywords", Type: "[]string", Required: false},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubContentVersion", Flag: "hub-content-version", Type: "*string", Required: true},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "SupportStatus", Flag: "support-status", Type: "types.HubContentSupportStatus", Required: false},
}

var fields_update_hub_content_reference = []leanruntime.Field{
	{Name: "HubContentName", Flag: "hub-content-name", Type: "*string", Required: true},
	{Name: "HubContentType", Flag: "hub-content-type", Type: "types.HubContentType", Required: true},
	{Name: "HubName", Flag: "hub-name", Type: "*string", Required: true},
	{Name: "MinVersion", Flag: "min-version", Type: "*string", Required: false},
}

var fields_update_image = []leanruntime.Field{
	{Name: "DeleteProperties", Flag: "delete-properties", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_image_version = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "AliasesToAdd", Flag: "aliases-to-add", Type: "[]string", Required: false},
	{Name: "AliasesToDelete", Flag: "aliases-to-delete", Type: "[]string", Required: false},
	{Name: "Horovod", Flag: "horovod", Type: "*bool", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: false},
	{Name: "MLFramework", Flag: "ml-framework", Type: "*string", Required: false},
	{Name: "Processor", Flag: "processor", Type: "types.Processor", Required: false},
	{Name: "ProgrammingLang", Flag: "programming-lang", Type: "*string", Required: false},
	{Name: "ReleaseNotes", Flag: "release-notes", Type: "*string", Required: false},
	{Name: "VendorGuidance", Flag: "vendor-guidance", Type: "types.VendorGuidance", Required: false},
	{Name: "Version", Flag: "version", Type: "*int32", Required: false},
}

var fields_update_inference_component = []leanruntime.Field{
	{Name: "DeploymentConfig", Flag: "deployment-config", Type: "*types.InferenceComponentDeploymentConfig", Required: false},
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: true},
	{Name: "RuntimeConfig", Flag: "runtime-config", Type: "*types.InferenceComponentRuntimeConfig", Required: false},
	{Name: "Specification", Flag: "specification", Type: "*types.InferenceComponentSpecification", Required: false},
}

var fields_update_inference_component_runtime_config = []leanruntime.Field{
	{Name: "DesiredRuntimeConfig", Flag: "desired-runtime-config", Type: "*types.InferenceComponentRuntimeConfig", Required: true},
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: true},
}

var fields_update_inference_experiment = []leanruntime.Field{
	{Name: "DataStorageConfig", Flag: "data-storage-config", Type: "*types.InferenceExperimentDataStorageConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModelVariants", Flag: "model-variants", Type: "[]types.ModelVariantConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.InferenceExperimentSchedule", Required: false},
	{Name: "ShadowModeConfig", Flag: "shadow-mode-config", Type: "*types.ShadowModeConfig", Required: false},
}

var fields_update_mlflow_app = []leanruntime.Field{
	{Name: "AccountDefaultStatus", Flag: "account-default-status", Type: "types.AccountDefaultStatus", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ArtifactStoreUri", Flag: "artifact-store-uri", Type: "*string", Required: false},
	{Name: "DefaultDomainIdList", Flag: "default-domain-id-list", Type: "[]string", Required: false},
	{Name: "ModelRegistrationMode", Flag: "model-registration-mode", Type: "types.ModelRegistrationMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "WeeklyMaintenanceWindowStart", Flag: "weekly-maintenance-window-start", Type: "*string", Required: false},
}

var fields_update_mlflow_tracking_server = []leanruntime.Field{
	{Name: "ArtifactStoreUri", Flag: "artifact-store-uri", Type: "*string", Required: false},
	{Name: "AutomaticModelRegistration", Flag: "automatic-model-registration", Type: "*bool", Required: false},
	{Name: "TrackingServerName", Flag: "tracking-server-name", Type: "*string", Required: true},
	{Name: "TrackingServerSize", Flag: "tracking-server-size", Type: "types.TrackingServerSize", Required: false},
	{Name: "WeeklyMaintenanceWindowStart", Flag: "weekly-maintenance-window-start", Type: "*string", Required: false},
}

var fields_update_model_card = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: false},
	{Name: "ModelCardName", Flag: "model-card-name", Type: "*string", Required: true},
	{Name: "ModelCardStatus", Flag: "model-card-status", Type: "types.ModelCardStatus", Required: false},
}

var fields_update_model_package = []leanruntime.Field{
	{Name: "AdditionalInferenceSpecificationsToAdd", Flag: "additional-inference-specifications-to-add", Type: "[]types.AdditionalInferenceSpecificationDefinition", Required: false},
	{Name: "ApprovalDescription", Flag: "approval-description", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerMetadataProperties", Flag: "customer-metadata-properties", Type: "map[string]string", Required: false},
	{Name: "CustomerMetadataPropertiesToRemove", Flag: "customer-metadata-properties-to-remove", Type: "[]string", Required: false},
	{Name: "InferenceSpecification", Flag: "inference-specification", Type: "*types.InferenceSpecification", Required: false},
	{Name: "ModelApprovalStatus", Flag: "model-approval-status", Type: "types.ModelApprovalStatus", Required: false},
	{Name: "ModelCard", Flag: "model-card", Type: "*types.ModelPackageModelCard", Required: false},
	{Name: "ModelLifeCycle", Flag: "model-life-cycle", Type: "*types.ModelLifeCycle", Required: false},
	{Name: "ModelPackageArn", Flag: "model-package-arn", Type: "*string", Required: true},
	{Name: "ModelPackageRegistrationType", Flag: "model-package-registration-type", Type: "types.ModelPackageRegistrationType", Required: false},
	{Name: "SourceUri", Flag: "source-uri", Type: "*string", Required: false},
}

var fields_update_monitoring_alert = []leanruntime.Field{
	{Name: "DatapointsToAlert", Flag: "datapoints-to-alert", Type: "*int32", Required: true},
	{Name: "EvaluationPeriod", Flag: "evaluation-period", Type: "*int32", Required: true},
	{Name: "MonitoringAlertName", Flag: "monitoring-alert-name", Type: "*string", Required: true},
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
}

var fields_update_monitoring_schedule = []leanruntime.Field{
	{Name: "MonitoringScheduleConfig", Flag: "monitoring-schedule-config", Type: "*types.MonitoringScheduleConfig", Required: true},
	{Name: "MonitoringScheduleName", Flag: "monitoring-schedule-name", Type: "*string", Required: true},
}

var fields_update_notebook_instance = []leanruntime.Field{
	{Name: "AcceleratorTypes", Flag: "accelerator-types", Type: "[]types.NotebookInstanceAcceleratorType", Required: false},
	{Name: "AdditionalCodeRepositories", Flag: "additional-code-repositories", Type: "[]string", Required: false},
	{Name: "DefaultCodeRepository", Flag: "default-code-repository", Type: "*string", Required: false},
	{Name: "DisassociateAcceleratorTypes", Flag: "disassociate-accelerator-types", Type: "*bool", Required: false},
	{Name: "DisassociateAdditionalCodeRepositories", Flag: "disassociate-additional-code-repositories", Type: "*bool", Required: false},
	{Name: "DisassociateDefaultCodeRepository", Flag: "disassociate-default-code-repository", Type: "*bool", Required: false},
	{Name: "DisassociateLifecycleConfig", Flag: "disassociate-lifecycle-config", Type: "*bool", Required: false},
	{Name: "InstanceMetadataServiceConfiguration", Flag: "instance-metadata-service-configuration", Type: "*types.InstanceMetadataServiceConfiguration", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IPAddressType", Required: false},
	{Name: "LifecycleConfigName", Flag: "lifecycle-config-name", Type: "*string", Required: false},
	{Name: "NotebookInstanceName", Flag: "notebook-instance-name", Type: "*string", Required: true},
	{Name: "PlatformIdentifier", Flag: "platform-identifier", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RootAccess", Flag: "root-access", Type: "types.RootAccess", Required: false},
	{Name: "VolumeSizeInGB", Flag: "volume-size-in-gb", Type: "*int32", Required: false},
}

var fields_update_notebook_instance_lifecycle_config = []leanruntime.Field{
	{Name: "NotebookInstanceLifecycleConfigName", Flag: "notebook-instance-lifecycle-config-name", Type: "*string", Required: true},
	{Name: "OnCreate", Flag: "on-create", Type: "[]types.NotebookInstanceLifecycleHook", Required: false},
	{Name: "OnStart", Flag: "on-start", Type: "[]types.NotebookInstanceLifecycleHook", Required: false},
}

var fields_update_partner_app = []leanruntime.Field{
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: false},
	{Name: "ApplicationConfig", Flag: "application-config", Type: "*types.PartnerAppConfig", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnableAutoMinorVersionUpgrade", Flag: "enable-auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "EnableIamSessionBasedIdentity", Flag: "enable-iam-session-based-identity", Type: "*bool", Required: false},
	{Name: "MaintenanceConfig", Flag: "maintenance-config", Type: "*types.PartnerAppMaintenanceConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Tier", Flag: "tier", Type: "*string", Required: false},
}

var fields_update_pipeline = []leanruntime.Field{
	{Name: "ParallelismConfiguration", Flag: "parallelism-configuration", Type: "*types.ParallelismConfiguration", Required: false},
	{Name: "PipelineDefinition", Flag: "pipeline-definition", Type: "*string", Required: false},
	{Name: "PipelineDefinitionS3Location", Flag: "pipeline-definition-s3-location", Type: "*types.PipelineDefinitionS3Location", Required: false},
	{Name: "PipelineDescription", Flag: "pipeline-description", Type: "*string", Required: false},
	{Name: "PipelineDisplayName", Flag: "pipeline-display-name", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_pipeline_execution = []leanruntime.Field{
	{Name: "ParallelismConfiguration", Flag: "parallelism-configuration", Type: "*types.ParallelismConfiguration", Required: false},
	{Name: "PipelineExecutionArn", Flag: "pipeline-execution-arn", Type: "*string", Required: true},
	{Name: "PipelineExecutionDescription", Flag: "pipeline-execution-description", Type: "*string", Required: false},
	{Name: "PipelineExecutionDisplayName", Flag: "pipeline-execution-display-name", Type: "*string", Required: false},
}

var fields_update_pipeline_version = []leanruntime.Field{
	{Name: "PipelineArn", Flag: "pipeline-arn", Type: "*string", Required: true},
	{Name: "PipelineVersionDescription", Flag: "pipeline-version-description", Type: "*string", Required: false},
	{Name: "PipelineVersionDisplayName", Flag: "pipeline-version-display-name", Type: "*string", Required: false},
	{Name: "PipelineVersionId", Flag: "pipeline-version-id", Type: "*int64", Required: true},
}

var fields_update_project = []leanruntime.Field{
	{Name: "ProjectDescription", Flag: "project-description", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "ServiceCatalogProvisioningUpdateDetails", Flag: "service-catalog-provisioning-update-details", Type: "*types.ServiceCatalogProvisioningUpdateDetails", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateProvidersToUpdate", Flag: "template-providers-to-update", Type: "[]types.UpdateTemplateProvider", Required: false},
}

var fields_update_space = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpaceDisplayName", Flag: "space-display-name", Type: "*string", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
	{Name: "SpaceSettings", Flag: "space-settings", Type: "*types.SpaceSettings", Required: false},
}

var fields_update_training_job = []leanruntime.Field{
	{Name: "ProfilerConfig", Flag: "profiler-config", Type: "*types.ProfilerConfigForUpdate", Required: false},
	{Name: "ProfilerRuleConfigurations", Flag: "profiler-rule-configurations", Type: "[]types.ProfilerRuleConfiguration", Required: false},
	{Name: "RemoteDebugConfig", Flag: "remote-debug-config", Type: "*types.RemoteDebugConfigForUpdate", Required: false},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "*types.ResourceConfigForUpdate", Required: false},
	{Name: "TrainingJobName", Flag: "training-job-name", Type: "*string", Required: true},
}

var fields_update_trial = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "TrialName", Flag: "trial-name", Type: "*string", Required: true},
}

var fields_update_trial_component = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "InputArtifacts", Flag: "input-artifacts", Type: "map[string]types.TrialComponentArtifact", Required: false},
	{Name: "InputArtifactsToRemove", Flag: "input-artifacts-to-remove", Type: "[]string", Required: false},
	{Name: "OutputArtifacts", Flag: "output-artifacts", Type: "map[string]types.TrialComponentArtifact", Required: false},
	{Name: "OutputArtifactsToRemove", Flag: "output-artifacts-to-remove", Type: "[]string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.TrialComponentParameterValue", Required: false},
	{Name: "ParametersToRemove", Flag: "parameters-to-remove", Type: "[]string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "*types.TrialComponentStatus", Required: false},
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
}

var fields_update_user_profile = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "UserProfileName", Flag: "user-profile-name", Type: "*string", Required: true},
	{Name: "UserSettings", Flag: "user-settings", Type: "*types.UserSettings", Required: false},
}

var fields_update_workforce = []leanruntime.Field{
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.WorkforceIpAddressType", Required: false},
	{Name: "OidcConfig", Flag: "oidc-config", Type: "*types.OidcConfig", Required: false},
	{Name: "SourceIpConfig", Flag: "source-ip-config", Type: "*types.SourceIpConfig", Required: false},
	{Name: "WorkforceName", Flag: "workforce-name", Type: "*string", Required: true},
	{Name: "WorkforceVpcConfig", Flag: "workforce-vpc-config", Type: "*types.WorkforceVpcConfigRequest", Required: false},
}

var fields_update_workteam = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MemberDefinitions", Flag: "member-definitions", Type: "[]types.MemberDefinition", Required: false},
	{Name: "NotificationConfiguration", Flag: "notification-configuration", Type: "*types.NotificationConfiguration", Required: false},
	{Name: "WorkerAccessConfiguration", Flag: "worker-access-configuration", Type: "*types.WorkerAccessConfiguration", Required: false},
	{Name: "WorkteamName", Flag: "workteam-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-association": {
			Name:   "add-association",
			Fields: fields_add_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddAssociation(ctx, input)
			},
		},
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
		"associate-trial-component": {
			Name:   "associate-trial-component",
			Fields: fields_associate_trial_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTrialComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_trial_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTrialComponent(ctx, input)
			},
		},
		"attach-cluster-node-volume": {
			Name:   "attach-cluster-node-volume",
			Fields: fields_attach_cluster_node_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachClusterNodeVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_cluster_node_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachClusterNodeVolume(ctx, input)
			},
		},
		"batch-add-cluster-nodes": {
			Name:   "batch-add-cluster-nodes",
			Fields: fields_batch_add_cluster_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAddClusterNodesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_add_cluster_nodes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAddClusterNodes(ctx, input)
			},
		},
		"batch-delete-cluster-nodes": {
			Name:   "batch-delete-cluster-nodes",
			Fields: fields_batch_delete_cluster_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteClusterNodesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_cluster_nodes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteClusterNodes(ctx, input)
			},
		},
		"batch-describe-model-package": {
			Name:   "batch-describe-model-package",
			Fields: fields_batch_describe_model_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDescribeModelPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_describe_model_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDescribeModelPackage(ctx, input)
			},
		},
		"batch-reboot-cluster-nodes": {
			Name:   "batch-reboot-cluster-nodes",
			Fields: fields_batch_reboot_cluster_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchRebootClusterNodesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_reboot_cluster_nodes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchRebootClusterNodes(ctx, input)
			},
		},
		"batch-replace-cluster-nodes": {
			Name:   "batch-replace-cluster-nodes",
			Fields: fields_batch_replace_cluster_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchReplaceClusterNodesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_replace_cluster_nodes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchReplaceClusterNodes(ctx, input)
			},
		},
		"create-action": {
			Name:   "create-action",
			Fields: fields_create_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAction(ctx, input)
			},
		},
		"create-algorithm": {
			Name:   "create-algorithm",
			Fields: fields_create_algorithm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAlgorithmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_algorithm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlgorithm(ctx, input)
			},
		},
		"create-app": {
			Name:   "create-app",
			Fields: fields_create_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApp(ctx, input)
			},
		},
		"create-app-image-config": {
			Name:   "create-app-image-config",
			Fields: fields_create_app_image_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppImageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_image_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppImageConfig(ctx, input)
			},
		},
		"create-artifact": {
			Name:   "create-artifact",
			Fields: fields_create_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateArtifact(ctx, input)
			},
		},
		"create-auto-ml-job": {
			Name:   "create-auto-ml-job",
			Fields: fields_create_auto_ml_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutoMLJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_auto_ml_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutoMLJob(ctx, input)
			},
		},
		"create-auto-ml-job-v2": {
			Name:   "create-auto-ml-job-v2",
			Fields: fields_create_auto_ml_job_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutoMLJobV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_auto_ml_job_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutoMLJobV2(ctx, input)
			},
		},
		"create-cluster": {
			Name:   "create-cluster",
			Fields: fields_create_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCluster(ctx, input)
			},
		},
		"create-cluster-scheduler-config": {
			Name:   "create-cluster-scheduler-config",
			Fields: fields_create_cluster_scheduler_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterSchedulerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster_scheduler_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClusterSchedulerConfig(ctx, input)
			},
		},
		"create-code-repository": {
			Name:   "create-code-repository",
			Fields: fields_create_code_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCodeRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_code_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCodeRepository(ctx, input)
			},
		},
		"create-compilation-job": {
			Name:   "create-compilation-job",
			Fields: fields_create_compilation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCompilationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_compilation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCompilationJob(ctx, input)
			},
		},
		"create-compute-quota": {
			Name:   "create-compute-quota",
			Fields: fields_create_compute_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComputeQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_compute_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComputeQuota(ctx, input)
			},
		},
		"create-context": {
			Name:   "create-context",
			Fields: fields_create_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContext(ctx, input)
			},
		},
		"create-data-quality-job-definition": {
			Name:   "create-data-quality-job-definition",
			Fields: fields_create_data_quality_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataQualityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_quality_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataQualityJobDefinition(ctx, input)
			},
		},
		"create-device-fleet": {
			Name:   "create-device-fleet",
			Fields: fields_create_device_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeviceFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_device_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeviceFleet(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-edge-deployment-plan": {
			Name:   "create-edge-deployment-plan",
			Fields: fields_create_edge_deployment_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEdgeDeploymentPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_edge_deployment_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEdgeDeploymentPlan(ctx, input)
			},
		},
		"create-edge-deployment-stage": {
			Name:   "create-edge-deployment-stage",
			Fields: fields_create_edge_deployment_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEdgeDeploymentStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_edge_deployment_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEdgeDeploymentStage(ctx, input)
			},
		},
		"create-edge-packaging-job": {
			Name:   "create-edge-packaging-job",
			Fields: fields_create_edge_packaging_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEdgePackagingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_edge_packaging_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEdgePackagingJob(ctx, input)
			},
		},
		"create-endpoint": {
			Name:   "create-endpoint",
			Fields: fields_create_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpoint(ctx, input)
			},
		},
		"create-endpoint-config": {
			Name:   "create-endpoint-config",
			Fields: fields_create_endpoint_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpointConfig(ctx, input)
			},
		},
		"create-experiment": {
			Name:   "create-experiment",
			Fields: fields_create_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExperiment(ctx, input)
			},
		},
		"create-feature-group": {
			Name:   "create-feature-group",
			Fields: fields_create_feature_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFeatureGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_feature_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFeatureGroup(ctx, input)
			},
		},
		"create-flow-definition": {
			Name:   "create-flow-definition",
			Fields: fields_create_flow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlowDefinition(ctx, input)
			},
		},
		"create-hub": {
			Name:   "create-hub",
			Fields: fields_create_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHub(ctx, input)
			},
		},
		"create-hub-content-presigned-urls": {
			Name:   "create-hub-content-presigned-urls",
			Fields: fields_create_hub_content_presigned_urls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHubContentPresignedUrlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_create_hub_content_presigned_urls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.CreateHubContentPresignedUrls(ctx, input)
				}
				var results []*svc.CreateHubContentPresignedUrlsOutput
				p := svc.NewCreateHubContentPresignedUrlsPaginator(client, input)
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
		"create-hub-content-reference": {
			Name:   "create-hub-content-reference",
			Fields: fields_create_hub_content_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHubContentReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hub_content_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHubContentReference(ctx, input)
			},
		},
		"create-human-task-ui": {
			Name:   "create-human-task-ui",
			Fields: fields_create_human_task_ui,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHumanTaskUiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_human_task_ui, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHumanTaskUi(ctx, input)
			},
		},
		"create-hyper-parameter-tuning-job": {
			Name:   "create-hyper-parameter-tuning-job",
			Fields: fields_create_hyper_parameter_tuning_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHyperParameterTuningJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hyper_parameter_tuning_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHyperParameterTuningJob(ctx, input)
			},
		},
		"create-image": {
			Name:   "create-image",
			Fields: fields_create_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImage(ctx, input)
			},
		},
		"create-image-version": {
			Name:   "create-image-version",
			Fields: fields_create_image_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImageVersion(ctx, input)
			},
		},
		"create-inference-component": {
			Name:   "create-inference-component",
			Fields: fields_create_inference_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInferenceComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_inference_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInferenceComponent(ctx, input)
			},
		},
		"create-inference-experiment": {
			Name:   "create-inference-experiment",
			Fields: fields_create_inference_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInferenceExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_inference_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInferenceExperiment(ctx, input)
			},
		},
		"create-inference-recommendations-job": {
			Name:   "create-inference-recommendations-job",
			Fields: fields_create_inference_recommendations_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInferenceRecommendationsJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_inference_recommendations_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInferenceRecommendationsJob(ctx, input)
			},
		},
		"create-labeling-job": {
			Name:   "create-labeling-job",
			Fields: fields_create_labeling_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLabelingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_labeling_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLabelingJob(ctx, input)
			},
		},
		"create-mlflow-app": {
			Name:   "create-mlflow-app",
			Fields: fields_create_mlflow_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMlflowAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mlflow_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMlflowApp(ctx, input)
			},
		},
		"create-mlflow-tracking-server": {
			Name:   "create-mlflow-tracking-server",
			Fields: fields_create_mlflow_tracking_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMlflowTrackingServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mlflow_tracking_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMlflowTrackingServer(ctx, input)
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
		"create-model-bias-job-definition": {
			Name:   "create-model-bias-job-definition",
			Fields: fields_create_model_bias_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelBiasJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_bias_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelBiasJobDefinition(ctx, input)
			},
		},
		"create-model-card": {
			Name:   "create-model-card",
			Fields: fields_create_model_card,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelCardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_card, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelCard(ctx, input)
			},
		},
		"create-model-card-export-job": {
			Name:   "create-model-card-export-job",
			Fields: fields_create_model_card_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelCardExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_card_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelCardExportJob(ctx, input)
			},
		},
		"create-model-explainability-job-definition": {
			Name:   "create-model-explainability-job-definition",
			Fields: fields_create_model_explainability_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelExplainabilityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_explainability_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelExplainabilityJobDefinition(ctx, input)
			},
		},
		"create-model-package": {
			Name:   "create-model-package",
			Fields: fields_create_model_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelPackage(ctx, input)
			},
		},
		"create-model-package-group": {
			Name:   "create-model-package-group",
			Fields: fields_create_model_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelPackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelPackageGroup(ctx, input)
			},
		},
		"create-model-quality-job-definition": {
			Name:   "create-model-quality-job-definition",
			Fields: fields_create_model_quality_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelQualityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_quality_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelQualityJobDefinition(ctx, input)
			},
		},
		"create-monitoring-schedule": {
			Name:   "create-monitoring-schedule",
			Fields: fields_create_monitoring_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMonitoringScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_monitoring_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMonitoringSchedule(ctx, input)
			},
		},
		"create-notebook-instance": {
			Name:   "create-notebook-instance",
			Fields: fields_create_notebook_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotebookInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notebook_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotebookInstance(ctx, input)
			},
		},
		"create-notebook-instance-lifecycle-config": {
			Name:   "create-notebook-instance-lifecycle-config",
			Fields: fields_create_notebook_instance_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotebookInstanceLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notebook_instance_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotebookInstanceLifecycleConfig(ctx, input)
			},
		},
		"create-optimization-job": {
			Name:   "create-optimization-job",
			Fields: fields_create_optimization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOptimizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_optimization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOptimizationJob(ctx, input)
			},
		},
		"create-partner-app": {
			Name:   "create-partner-app",
			Fields: fields_create_partner_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartnerAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partner_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartnerApp(ctx, input)
			},
		},
		"create-partner-app-presigned-url": {
			Name:   "create-partner-app-presigned-url",
			Fields: fields_create_partner_app_presigned_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartnerAppPresignedUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partner_app_presigned_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartnerAppPresignedUrl(ctx, input)
			},
		},
		"create-pipeline": {
			Name:   "create-pipeline",
			Fields: fields_create_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePipeline(ctx, input)
			},
		},
		"create-presigned-domain-url": {
			Name:   "create-presigned-domain-url",
			Fields: fields_create_presigned_domain_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePresignedDomainUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_presigned_domain_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePresignedDomainUrl(ctx, input)
			},
		},
		"create-presigned-mlflow-app-url": {
			Name:   "create-presigned-mlflow-app-url",
			Fields: fields_create_presigned_mlflow_app_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePresignedMlflowAppUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_presigned_mlflow_app_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePresignedMlflowAppUrl(ctx, input)
			},
		},
		"create-presigned-mlflow-tracking-server-url": {
			Name:   "create-presigned-mlflow-tracking-server-url",
			Fields: fields_create_presigned_mlflow_tracking_server_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePresignedMlflowTrackingServerUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_presigned_mlflow_tracking_server_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePresignedMlflowTrackingServerUrl(ctx, input)
			},
		},
		"create-presigned-notebook-instance-url": {
			Name:   "create-presigned-notebook-instance-url",
			Fields: fields_create_presigned_notebook_instance_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePresignedNotebookInstanceUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_presigned_notebook_instance_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePresignedNotebookInstanceUrl(ctx, input)
			},
		},
		"create-processing-job": {
			Name:   "create-processing-job",
			Fields: fields_create_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProcessingJob(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"create-space": {
			Name:   "create-space",
			Fields: fields_create_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSpace(ctx, input)
			},
		},
		"create-studio-lifecycle-config": {
			Name:   "create-studio-lifecycle-config",
			Fields: fields_create_studio_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStudioLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_studio_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStudioLifecycleConfig(ctx, input)
			},
		},
		"create-training-job": {
			Name:   "create-training-job",
			Fields: fields_create_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrainingJob(ctx, input)
			},
		},
		"create-training-plan": {
			Name:   "create-training-plan",
			Fields: fields_create_training_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrainingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_training_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrainingPlan(ctx, input)
			},
		},
		"create-transform-job": {
			Name:   "create-transform-job",
			Fields: fields_create_transform_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransformJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transform_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransformJob(ctx, input)
			},
		},
		"create-trial": {
			Name:   "create-trial",
			Fields: fields_create_trial,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trial, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrial(ctx, input)
			},
		},
		"create-trial-component": {
			Name:   "create-trial-component",
			Fields: fields_create_trial_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrialComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trial_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrialComponent(ctx, input)
			},
		},
		"create-user-profile": {
			Name:   "create-user-profile",
			Fields: fields_create_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserProfile(ctx, input)
			},
		},
		"create-workforce": {
			Name:   "create-workforce",
			Fields: fields_create_workforce,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkforceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workforce, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkforce(ctx, input)
			},
		},
		"create-workteam": {
			Name:   "create-workteam",
			Fields: fields_create_workteam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkteamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workteam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkteam(ctx, input)
			},
		},
		"delete-action": {
			Name:   "delete-action",
			Fields: fields_delete_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAction(ctx, input)
			},
		},
		"delete-algorithm": {
			Name:   "delete-algorithm",
			Fields: fields_delete_algorithm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlgorithmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_algorithm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlgorithm(ctx, input)
			},
		},
		"delete-app": {
			Name:   "delete-app",
			Fields: fields_delete_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApp(ctx, input)
			},
		},
		"delete-app-image-config": {
			Name:   "delete-app-image-config",
			Fields: fields_delete_app_image_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppImageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_image_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppImageConfig(ctx, input)
			},
		},
		"delete-artifact": {
			Name:   "delete-artifact",
			Fields: fields_delete_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteArtifact(ctx, input)
			},
		},
		"delete-association": {
			Name:   "delete-association",
			Fields: fields_delete_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssociation(ctx, input)
			},
		},
		"delete-cluster": {
			Name:   "delete-cluster",
			Fields: fields_delete_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCluster(ctx, input)
			},
		},
		"delete-cluster-scheduler-config": {
			Name:   "delete-cluster-scheduler-config",
			Fields: fields_delete_cluster_scheduler_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterSchedulerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster_scheduler_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClusterSchedulerConfig(ctx, input)
			},
		},
		"delete-code-repository": {
			Name:   "delete-code-repository",
			Fields: fields_delete_code_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCodeRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_code_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCodeRepository(ctx, input)
			},
		},
		"delete-compilation-job": {
			Name:   "delete-compilation-job",
			Fields: fields_delete_compilation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCompilationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_compilation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCompilationJob(ctx, input)
			},
		},
		"delete-compute-quota": {
			Name:   "delete-compute-quota",
			Fields: fields_delete_compute_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComputeQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_compute_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComputeQuota(ctx, input)
			},
		},
		"delete-context": {
			Name:   "delete-context",
			Fields: fields_delete_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContext(ctx, input)
			},
		},
		"delete-data-quality-job-definition": {
			Name:   "delete-data-quality-job-definition",
			Fields: fields_delete_data_quality_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataQualityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_quality_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataQualityJobDefinition(ctx, input)
			},
		},
		"delete-device-fleet": {
			Name:   "delete-device-fleet",
			Fields: fields_delete_device_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeviceFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_device_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeviceFleet(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-edge-deployment-plan": {
			Name:   "delete-edge-deployment-plan",
			Fields: fields_delete_edge_deployment_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEdgeDeploymentPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_edge_deployment_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEdgeDeploymentPlan(ctx, input)
			},
		},
		"delete-edge-deployment-stage": {
			Name:   "delete-edge-deployment-stage",
			Fields: fields_delete_edge_deployment_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEdgeDeploymentStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_edge_deployment_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEdgeDeploymentStage(ctx, input)
			},
		},
		"delete-endpoint": {
			Name:   "delete-endpoint",
			Fields: fields_delete_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpoint(ctx, input)
			},
		},
		"delete-endpoint-config": {
			Name:   "delete-endpoint-config",
			Fields: fields_delete_endpoint_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpointConfig(ctx, input)
			},
		},
		"delete-experiment": {
			Name:   "delete-experiment",
			Fields: fields_delete_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExperiment(ctx, input)
			},
		},
		"delete-feature-group": {
			Name:   "delete-feature-group",
			Fields: fields_delete_feature_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFeatureGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_feature_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFeatureGroup(ctx, input)
			},
		},
		"delete-flow-definition": {
			Name:   "delete-flow-definition",
			Fields: fields_delete_flow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlowDefinition(ctx, input)
			},
		},
		"delete-hub": {
			Name:   "delete-hub",
			Fields: fields_delete_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHub(ctx, input)
			},
		},
		"delete-hub-content": {
			Name:   "delete-hub-content",
			Fields: fields_delete_hub_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHubContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hub_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHubContent(ctx, input)
			},
		},
		"delete-hub-content-reference": {
			Name:   "delete-hub-content-reference",
			Fields: fields_delete_hub_content_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHubContentReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hub_content_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHubContentReference(ctx, input)
			},
		},
		"delete-human-task-ui": {
			Name:   "delete-human-task-ui",
			Fields: fields_delete_human_task_ui,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHumanTaskUiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_human_task_ui, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHumanTaskUi(ctx, input)
			},
		},
		"delete-hyper-parameter-tuning-job": {
			Name:   "delete-hyper-parameter-tuning-job",
			Fields: fields_delete_hyper_parameter_tuning_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHyperParameterTuningJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hyper_parameter_tuning_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHyperParameterTuningJob(ctx, input)
			},
		},
		"delete-image": {
			Name:   "delete-image",
			Fields: fields_delete_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImage(ctx, input)
			},
		},
		"delete-image-version": {
			Name:   "delete-image-version",
			Fields: fields_delete_image_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImageVersion(ctx, input)
			},
		},
		"delete-inference-component": {
			Name:   "delete-inference-component",
			Fields: fields_delete_inference_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInferenceComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inference_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInferenceComponent(ctx, input)
			},
		},
		"delete-inference-experiment": {
			Name:   "delete-inference-experiment",
			Fields: fields_delete_inference_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInferenceExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inference_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInferenceExperiment(ctx, input)
			},
		},
		"delete-mlflow-app": {
			Name:   "delete-mlflow-app",
			Fields: fields_delete_mlflow_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMlflowAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mlflow_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMlflowApp(ctx, input)
			},
		},
		"delete-mlflow-tracking-server": {
			Name:   "delete-mlflow-tracking-server",
			Fields: fields_delete_mlflow_tracking_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMlflowTrackingServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mlflow_tracking_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMlflowTrackingServer(ctx, input)
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
		"delete-model-bias-job-definition": {
			Name:   "delete-model-bias-job-definition",
			Fields: fields_delete_model_bias_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelBiasJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_bias_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelBiasJobDefinition(ctx, input)
			},
		},
		"delete-model-card": {
			Name:   "delete-model-card",
			Fields: fields_delete_model_card,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelCardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_card, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelCard(ctx, input)
			},
		},
		"delete-model-explainability-job-definition": {
			Name:   "delete-model-explainability-job-definition",
			Fields: fields_delete_model_explainability_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelExplainabilityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_explainability_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelExplainabilityJobDefinition(ctx, input)
			},
		},
		"delete-model-package": {
			Name:   "delete-model-package",
			Fields: fields_delete_model_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelPackage(ctx, input)
			},
		},
		"delete-model-package-group": {
			Name:   "delete-model-package-group",
			Fields: fields_delete_model_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelPackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelPackageGroup(ctx, input)
			},
		},
		"delete-model-package-group-policy": {
			Name:   "delete-model-package-group-policy",
			Fields: fields_delete_model_package_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelPackageGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_package_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelPackageGroupPolicy(ctx, input)
			},
		},
		"delete-model-quality-job-definition": {
			Name:   "delete-model-quality-job-definition",
			Fields: fields_delete_model_quality_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelQualityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_quality_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelQualityJobDefinition(ctx, input)
			},
		},
		"delete-monitoring-schedule": {
			Name:   "delete-monitoring-schedule",
			Fields: fields_delete_monitoring_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMonitoringScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_monitoring_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMonitoringSchedule(ctx, input)
			},
		},
		"delete-notebook-instance": {
			Name:   "delete-notebook-instance",
			Fields: fields_delete_notebook_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotebookInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notebook_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotebookInstance(ctx, input)
			},
		},
		"delete-notebook-instance-lifecycle-config": {
			Name:   "delete-notebook-instance-lifecycle-config",
			Fields: fields_delete_notebook_instance_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotebookInstanceLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notebook_instance_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotebookInstanceLifecycleConfig(ctx, input)
			},
		},
		"delete-optimization-job": {
			Name:   "delete-optimization-job",
			Fields: fields_delete_optimization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOptimizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_optimization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOptimizationJob(ctx, input)
			},
		},
		"delete-partner-app": {
			Name:   "delete-partner-app",
			Fields: fields_delete_partner_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePartnerAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_partner_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePartnerApp(ctx, input)
			},
		},
		"delete-pipeline": {
			Name:   "delete-pipeline",
			Fields: fields_delete_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePipeline(ctx, input)
			},
		},
		"delete-processing-job": {
			Name:   "delete-processing-job",
			Fields: fields_delete_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProcessingJob(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-space": {
			Name:   "delete-space",
			Fields: fields_delete_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSpace(ctx, input)
			},
		},
		"delete-studio-lifecycle-config": {
			Name:   "delete-studio-lifecycle-config",
			Fields: fields_delete_studio_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStudioLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_studio_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStudioLifecycleConfig(ctx, input)
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
		"delete-training-job": {
			Name:   "delete-training-job",
			Fields: fields_delete_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrainingJob(ctx, input)
			},
		},
		"delete-trial": {
			Name:   "delete-trial",
			Fields: fields_delete_trial,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trial, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrial(ctx, input)
			},
		},
		"delete-trial-component": {
			Name:   "delete-trial-component",
			Fields: fields_delete_trial_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrialComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trial_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrialComponent(ctx, input)
			},
		},
		"delete-user-profile": {
			Name:   "delete-user-profile",
			Fields: fields_delete_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserProfile(ctx, input)
			},
		},
		"delete-workforce": {
			Name:   "delete-workforce",
			Fields: fields_delete_workforce,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkforceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workforce, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkforce(ctx, input)
			},
		},
		"delete-workteam": {
			Name:   "delete-workteam",
			Fields: fields_delete_workteam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkteamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workteam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkteam(ctx, input)
			},
		},
		"deregister-devices": {
			Name:   "deregister-devices",
			Fields: fields_deregister_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterDevicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_devices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterDevices(ctx, input)
			},
		},
		"describe-action": {
			Name:   "describe-action",
			Fields: fields_describe_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAction(ctx, input)
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
		"describe-app": {
			Name:   "describe-app",
			Fields: fields_describe_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApp(ctx, input)
			},
		},
		"describe-app-image-config": {
			Name:   "describe-app-image-config",
			Fields: fields_describe_app_image_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppImageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_image_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppImageConfig(ctx, input)
			},
		},
		"describe-artifact": {
			Name:   "describe-artifact",
			Fields: fields_describe_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeArtifact(ctx, input)
			},
		},
		"describe-auto-ml-job": {
			Name:   "describe-auto-ml-job",
			Fields: fields_describe_auto_ml_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoMLJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_auto_ml_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAutoMLJob(ctx, input)
			},
		},
		"describe-auto-ml-job-v2": {
			Name:   "describe-auto-ml-job-v2",
			Fields: fields_describe_auto_ml_job_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoMLJobV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_auto_ml_job_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAutoMLJobV2(ctx, input)
			},
		},
		"describe-cluster": {
			Name:   "describe-cluster",
			Fields: fields_describe_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCluster(ctx, input)
			},
		},
		"describe-cluster-event": {
			Name:   "describe-cluster-event",
			Fields: fields_describe_cluster_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusterEvent(ctx, input)
			},
		},
		"describe-cluster-node": {
			Name:   "describe-cluster-node",
			Fields: fields_describe_cluster_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusterNode(ctx, input)
			},
		},
		"describe-cluster-scheduler-config": {
			Name:   "describe-cluster-scheduler-config",
			Fields: fields_describe_cluster_scheduler_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterSchedulerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster_scheduler_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusterSchedulerConfig(ctx, input)
			},
		},
		"describe-code-repository": {
			Name:   "describe-code-repository",
			Fields: fields_describe_code_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCodeRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_code_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCodeRepository(ctx, input)
			},
		},
		"describe-compilation-job": {
			Name:   "describe-compilation-job",
			Fields: fields_describe_compilation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCompilationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_compilation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCompilationJob(ctx, input)
			},
		},
		"describe-compute-quota": {
			Name:   "describe-compute-quota",
			Fields: fields_describe_compute_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComputeQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_compute_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComputeQuota(ctx, input)
			},
		},
		"describe-context": {
			Name:   "describe-context",
			Fields: fields_describe_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContext(ctx, input)
			},
		},
		"describe-data-quality-job-definition": {
			Name:   "describe-data-quality-job-definition",
			Fields: fields_describe_data_quality_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataQualityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_quality_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataQualityJobDefinition(ctx, input)
			},
		},
		"describe-device": {
			Name:   "describe-device",
			Fields: fields_describe_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDevice(ctx, input)
			},
		},
		"describe-device-fleet": {
			Name:   "describe-device-fleet",
			Fields: fields_describe_device_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeviceFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_device_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeviceFleet(ctx, input)
			},
		},
		"describe-domain": {
			Name:   "describe-domain",
			Fields: fields_describe_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomain(ctx, input)
			},
		},
		"describe-edge-deployment-plan": {
			Name:   "describe-edge-deployment-plan",
			Fields: fields_describe_edge_deployment_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEdgeDeploymentPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_edge_deployment_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEdgeDeploymentPlan(ctx, input)
			},
		},
		"describe-edge-packaging-job": {
			Name:   "describe-edge-packaging-job",
			Fields: fields_describe_edge_packaging_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEdgePackagingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_edge_packaging_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEdgePackagingJob(ctx, input)
			},
		},
		"describe-endpoint": {
			Name:   "describe-endpoint",
			Fields: fields_describe_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoint(ctx, input)
			},
		},
		"describe-endpoint-config": {
			Name:   "describe-endpoint-config",
			Fields: fields_describe_endpoint_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoint_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpointConfig(ctx, input)
			},
		},
		"describe-experiment": {
			Name:   "describe-experiment",
			Fields: fields_describe_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExperiment(ctx, input)
			},
		},
		"describe-feature-group": {
			Name:   "describe-feature-group",
			Fields: fields_describe_feature_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFeatureGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_feature_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFeatureGroup(ctx, input)
			},
		},
		"describe-feature-metadata": {
			Name:   "describe-feature-metadata",
			Fields: fields_describe_feature_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFeatureMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_feature_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFeatureMetadata(ctx, input)
			},
		},
		"describe-flow-definition": {
			Name:   "describe-flow-definition",
			Fields: fields_describe_flow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlowDefinition(ctx, input)
			},
		},
		"describe-hub": {
			Name:   "describe-hub",
			Fields: fields_describe_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHub(ctx, input)
			},
		},
		"describe-hub-content": {
			Name:   "describe-hub-content",
			Fields: fields_describe_hub_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHubContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hub_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHubContent(ctx, input)
			},
		},
		"describe-human-task-ui": {
			Name:   "describe-human-task-ui",
			Fields: fields_describe_human_task_ui,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHumanTaskUiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_human_task_ui, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHumanTaskUi(ctx, input)
			},
		},
		"describe-hyper-parameter-tuning-job": {
			Name:   "describe-hyper-parameter-tuning-job",
			Fields: fields_describe_hyper_parameter_tuning_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHyperParameterTuningJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hyper_parameter_tuning_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHyperParameterTuningJob(ctx, input)
			},
		},
		"describe-image": {
			Name:   "describe-image",
			Fields: fields_describe_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImage(ctx, input)
			},
		},
		"describe-image-version": {
			Name:   "describe-image-version",
			Fields: fields_describe_image_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageVersion(ctx, input)
			},
		},
		"describe-inference-component": {
			Name:   "describe-inference-component",
			Fields: fields_describe_inference_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInferenceComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_inference_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInferenceComponent(ctx, input)
			},
		},
		"describe-inference-experiment": {
			Name:   "describe-inference-experiment",
			Fields: fields_describe_inference_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInferenceExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_inference_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInferenceExperiment(ctx, input)
			},
		},
		"describe-inference-recommendations-job": {
			Name:   "describe-inference-recommendations-job",
			Fields: fields_describe_inference_recommendations_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInferenceRecommendationsJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_inference_recommendations_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInferenceRecommendationsJob(ctx, input)
			},
		},
		"describe-labeling-job": {
			Name:   "describe-labeling-job",
			Fields: fields_describe_labeling_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLabelingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_labeling_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLabelingJob(ctx, input)
			},
		},
		"describe-lineage-group": {
			Name:   "describe-lineage-group",
			Fields: fields_describe_lineage_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLineageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_lineage_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLineageGroup(ctx, input)
			},
		},
		"describe-mlflow-app": {
			Name:   "describe-mlflow-app",
			Fields: fields_describe_mlflow_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMlflowAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_mlflow_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMlflowApp(ctx, input)
			},
		},
		"describe-mlflow-tracking-server": {
			Name:   "describe-mlflow-tracking-server",
			Fields: fields_describe_mlflow_tracking_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMlflowTrackingServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_mlflow_tracking_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMlflowTrackingServer(ctx, input)
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
		"describe-model-bias-job-definition": {
			Name:   "describe-model-bias-job-definition",
			Fields: fields_describe_model_bias_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelBiasJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_bias_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelBiasJobDefinition(ctx, input)
			},
		},
		"describe-model-card": {
			Name:   "describe-model-card",
			Fields: fields_describe_model_card,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelCardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_card, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelCard(ctx, input)
			},
		},
		"describe-model-card-export-job": {
			Name:   "describe-model-card-export-job",
			Fields: fields_describe_model_card_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelCardExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_card_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelCardExportJob(ctx, input)
			},
		},
		"describe-model-explainability-job-definition": {
			Name:   "describe-model-explainability-job-definition",
			Fields: fields_describe_model_explainability_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelExplainabilityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_explainability_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelExplainabilityJobDefinition(ctx, input)
			},
		},
		"describe-model-package": {
			Name:   "describe-model-package",
			Fields: fields_describe_model_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelPackage(ctx, input)
			},
		},
		"describe-model-package-group": {
			Name:   "describe-model-package-group",
			Fields: fields_describe_model_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelPackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelPackageGroup(ctx, input)
			},
		},
		"describe-model-quality-job-definition": {
			Name:   "describe-model-quality-job-definition",
			Fields: fields_describe_model_quality_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelQualityJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_model_quality_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeModelQualityJobDefinition(ctx, input)
			},
		},
		"describe-monitoring-schedule": {
			Name:   "describe-monitoring-schedule",
			Fields: fields_describe_monitoring_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMonitoringScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_monitoring_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMonitoringSchedule(ctx, input)
			},
		},
		"describe-notebook-instance": {
			Name:   "describe-notebook-instance",
			Fields: fields_describe_notebook_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotebookInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_notebook_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNotebookInstance(ctx, input)
			},
		},
		"describe-notebook-instance-lifecycle-config": {
			Name:   "describe-notebook-instance-lifecycle-config",
			Fields: fields_describe_notebook_instance_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotebookInstanceLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_notebook_instance_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNotebookInstanceLifecycleConfig(ctx, input)
			},
		},
		"describe-optimization-job": {
			Name:   "describe-optimization-job",
			Fields: fields_describe_optimization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOptimizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_optimization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOptimizationJob(ctx, input)
			},
		},
		"describe-partner-app": {
			Name:   "describe-partner-app",
			Fields: fields_describe_partner_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePartnerAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_partner_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePartnerApp(ctx, input)
			},
		},
		"describe-pipeline": {
			Name:   "describe-pipeline",
			Fields: fields_describe_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePipeline(ctx, input)
			},
		},
		"describe-pipeline-definition-for-execution": {
			Name:   "describe-pipeline-definition-for-execution",
			Fields: fields_describe_pipeline_definition_for_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePipelineDefinitionForExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pipeline_definition_for_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePipelineDefinitionForExecution(ctx, input)
			},
		},
		"describe-pipeline-execution": {
			Name:   "describe-pipeline-execution",
			Fields: fields_describe_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePipelineExecution(ctx, input)
			},
		},
		"describe-processing-job": {
			Name:   "describe-processing-job",
			Fields: fields_describe_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProcessingJob(ctx, input)
			},
		},
		"describe-project": {
			Name:   "describe-project",
			Fields: fields_describe_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProject(ctx, input)
			},
		},
		"describe-reserved-capacity": {
			Name:   "describe-reserved-capacity",
			Fields: fields_describe_reserved_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_reserved_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReservedCapacity(ctx, input)
			},
		},
		"describe-space": {
			Name:   "describe-space",
			Fields: fields_describe_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSpace(ctx, input)
			},
		},
		"describe-studio-lifecycle-config": {
			Name:   "describe-studio-lifecycle-config",
			Fields: fields_describe_studio_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStudioLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_studio_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStudioLifecycleConfig(ctx, input)
			},
		},
		"describe-subscribed-workteam": {
			Name:   "describe-subscribed-workteam",
			Fields: fields_describe_subscribed_workteam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubscribedWorkteamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_subscribed_workteam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSubscribedWorkteam(ctx, input)
			},
		},
		"describe-training-job": {
			Name:   "describe-training-job",
			Fields: fields_describe_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrainingJob(ctx, input)
			},
		},
		"describe-training-plan": {
			Name:   "describe-training-plan",
			Fields: fields_describe_training_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrainingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_training_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrainingPlan(ctx, input)
			},
		},
		"describe-transform-job": {
			Name:   "describe-transform-job",
			Fields: fields_describe_transform_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransformJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_transform_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTransformJob(ctx, input)
			},
		},
		"describe-trial": {
			Name:   "describe-trial",
			Fields: fields_describe_trial,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trial, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrial(ctx, input)
			},
		},
		"describe-trial-component": {
			Name:   "describe-trial-component",
			Fields: fields_describe_trial_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrialComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trial_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrialComponent(ctx, input)
			},
		},
		"describe-user-profile": {
			Name:   "describe-user-profile",
			Fields: fields_describe_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserProfile(ctx, input)
			},
		},
		"describe-workforce": {
			Name:   "describe-workforce",
			Fields: fields_describe_workforce,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkforceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workforce, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkforce(ctx, input)
			},
		},
		"describe-workteam": {
			Name:   "describe-workteam",
			Fields: fields_describe_workteam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkteamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workteam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkteam(ctx, input)
			},
		},
		"detach-cluster-node-volume": {
			Name:   "detach-cluster-node-volume",
			Fields: fields_detach_cluster_node_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachClusterNodeVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_cluster_node_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachClusterNodeVolume(ctx, input)
			},
		},
		"disable-sagemaker-servicecatalog-portfolio": {
			Name:   "disable-sagemaker-servicecatalog-portfolio",
			Fields: fields_disable_sagemaker_servicecatalog_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSagemakerServicecatalogPortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_sagemaker_servicecatalog_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSagemakerServicecatalogPortfolio(ctx, input)
			},
		},
		"disassociate-trial-component": {
			Name:   "disassociate-trial-component",
			Fields: fields_disassociate_trial_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTrialComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_trial_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTrialComponent(ctx, input)
			},
		},
		"enable-sagemaker-servicecatalog-portfolio": {
			Name:   "enable-sagemaker-servicecatalog-portfolio",
			Fields: fields_enable_sagemaker_servicecatalog_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSagemakerServicecatalogPortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_sagemaker_servicecatalog_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSagemakerServicecatalogPortfolio(ctx, input)
			},
		},
		"get-device-fleet-report": {
			Name:   "get-device-fleet-report",
			Fields: fields_get_device_fleet_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceFleetReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_fleet_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceFleetReport(ctx, input)
			},
		},
		"get-lineage-group-policy": {
			Name:   "get-lineage-group-policy",
			Fields: fields_get_lineage_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLineageGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lineage_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLineageGroupPolicy(ctx, input)
			},
		},
		"get-model-package-group-policy": {
			Name:   "get-model-package-group-policy",
			Fields: fields_get_model_package_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelPackageGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_package_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelPackageGroupPolicy(ctx, input)
			},
		},
		"get-sagemaker-servicecatalog-portfolio-status": {
			Name:   "get-sagemaker-servicecatalog-portfolio-status",
			Fields: fields_get_sagemaker_servicecatalog_portfolio_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSagemakerServicecatalogPortfolioStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sagemaker_servicecatalog_portfolio_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSagemakerServicecatalogPortfolioStatus(ctx, input)
			},
		},
		"get-scaling-configuration-recommendation": {
			Name:   "get-scaling-configuration-recommendation",
			Fields: fields_get_scaling_configuration_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScalingConfigurationRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scaling_configuration_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScalingConfigurationRecommendation(ctx, input)
			},
		},
		"get-search-suggestions": {
			Name:   "get-search-suggestions",
			Fields: fields_get_search_suggestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSearchSuggestionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_search_suggestions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSearchSuggestions(ctx, input)
			},
		},
		"import-hub-content": {
			Name:   "import-hub-content",
			Fields: fields_import_hub_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportHubContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_hub_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportHubContent(ctx, input)
			},
		},
		"list-actions": {
			Name:   "list-actions",
			Fields: fields_list_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActions(ctx, input)
				}
				var results []*svc.ListActionsOutput
				p := svc.NewListActionsPaginator(client, input)
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
		"list-algorithms": {
			Name:   "list-algorithms",
			Fields: fields_list_algorithms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlgorithmsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_algorithms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAlgorithms(ctx, input)
				}
				var results []*svc.ListAlgorithmsOutput
				p := svc.NewListAlgorithmsPaginator(client, input)
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
		"list-aliases": {
			Name:   "list-aliases",
			Fields: fields_list_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAliases(ctx, input)
				}
				var results []*svc.ListAliasesOutput
				p := svc.NewListAliasesPaginator(client, input)
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
		"list-app-image-configs": {
			Name:   "list-app-image-configs",
			Fields: fields_list_app_image_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppImageConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_image_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppImageConfigs(ctx, input)
				}
				var results []*svc.ListAppImageConfigsOutput
				p := svc.NewListAppImageConfigsPaginator(client, input)
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
		"list-apps": {
			Name:   "list-apps",
			Fields: fields_list_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApps(ctx, input)
				}
				var results []*svc.ListAppsOutput
				p := svc.NewListAppsPaginator(client, input)
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
		"list-artifacts": {
			Name:   "list-artifacts",
			Fields: fields_list_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArtifactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_artifacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListArtifacts(ctx, input)
				}
				var results []*svc.ListArtifactsOutput
				p := svc.NewListArtifactsPaginator(client, input)
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
		"list-associations": {
			Name:   "list-associations",
			Fields: fields_list_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociations(ctx, input)
				}
				var results []*svc.ListAssociationsOutput
				p := svc.NewListAssociationsPaginator(client, input)
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
		"list-auto-ml-jobs": {
			Name:   "list-auto-ml-jobs",
			Fields: fields_list_auto_ml_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutoMLJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_auto_ml_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutoMLJobs(ctx, input)
				}
				var results []*svc.ListAutoMLJobsOutput
				p := svc.NewListAutoMLJobsPaginator(client, input)
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
		"list-candidates-for-auto-ml-job": {
			Name:   "list-candidates-for-auto-ml-job",
			Fields: fields_list_candidates_for_auto_ml_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCandidatesForAutoMLJobInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_candidates_for_auto_ml_job, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCandidatesForAutoMLJob(ctx, input)
				}
				var results []*svc.ListCandidatesForAutoMLJobOutput
				p := svc.NewListCandidatesForAutoMLJobPaginator(client, input)
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
		"list-cluster-events": {
			Name:   "list-cluster-events",
			Fields: fields_list_cluster_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterEvents(ctx, input)
				}
				var results []*svc.ListClusterEventsOutput
				p := svc.NewListClusterEventsPaginator(client, input)
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
		"list-cluster-nodes": {
			Name:   "list-cluster-nodes",
			Fields: fields_list_cluster_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterNodes(ctx, input)
				}
				var results []*svc.ListClusterNodesOutput
				p := svc.NewListClusterNodesPaginator(client, input)
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
		"list-cluster-scheduler-configs": {
			Name:   "list-cluster-scheduler-configs",
			Fields: fields_list_cluster_scheduler_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterSchedulerConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_scheduler_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterSchedulerConfigs(ctx, input)
				}
				var results []*svc.ListClusterSchedulerConfigsOutput
				p := svc.NewListClusterSchedulerConfigsPaginator(client, input)
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
		"list-clusters": {
			Name:   "list-clusters",
			Fields: fields_list_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusters(ctx, input)
				}
				var results []*svc.ListClustersOutput
				p := svc.NewListClustersPaginator(client, input)
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
		"list-code-repositories": {
			Name:   "list-code-repositories",
			Fields: fields_list_code_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeRepositoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_code_repositories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCodeRepositories(ctx, input)
				}
				var results []*svc.ListCodeRepositoriesOutput
				p := svc.NewListCodeRepositoriesPaginator(client, input)
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
		"list-compilation-jobs": {
			Name:   "list-compilation-jobs",
			Fields: fields_list_compilation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCompilationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compilation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCompilationJobs(ctx, input)
				}
				var results []*svc.ListCompilationJobsOutput
				p := svc.NewListCompilationJobsPaginator(client, input)
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
		"list-compute-quotas": {
			Name:   "list-compute-quotas",
			Fields: fields_list_compute_quotas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComputeQuotasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compute_quotas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComputeQuotas(ctx, input)
				}
				var results []*svc.ListComputeQuotasOutput
				p := svc.NewListComputeQuotasPaginator(client, input)
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
		"list-contexts": {
			Name:   "list-contexts",
			Fields: fields_list_contexts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContextsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contexts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContexts(ctx, input)
				}
				var results []*svc.ListContextsOutput
				p := svc.NewListContextsPaginator(client, input)
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
		"list-data-quality-job-definitions": {
			Name:   "list-data-quality-job-definitions",
			Fields: fields_list_data_quality_job_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityJobDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_quality_job_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataQualityJobDefinitions(ctx, input)
				}
				var results []*svc.ListDataQualityJobDefinitionsOutput
				p := svc.NewListDataQualityJobDefinitionsPaginator(client, input)
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
		"list-device-fleets": {
			Name:   "list-device-fleets",
			Fields: fields_list_device_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_device_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeviceFleets(ctx, input)
				}
				var results []*svc.ListDeviceFleetsOutput
				p := svc.NewListDeviceFleetsPaginator(client, input)
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
		"list-devices": {
			Name:   "list-devices",
			Fields: fields_list_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevices(ctx, input)
				}
				var results []*svc.ListDevicesOutput
				p := svc.NewListDevicesPaginator(client, input)
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
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-edge-deployment-plans": {
			Name:   "list-edge-deployment-plans",
			Fields: fields_list_edge_deployment_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEdgeDeploymentPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_edge_deployment_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEdgeDeploymentPlans(ctx, input)
				}
				var results []*svc.ListEdgeDeploymentPlansOutput
				p := svc.NewListEdgeDeploymentPlansPaginator(client, input)
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
		"list-edge-packaging-jobs": {
			Name:   "list-edge-packaging-jobs",
			Fields: fields_list_edge_packaging_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEdgePackagingJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_edge_packaging_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEdgePackagingJobs(ctx, input)
				}
				var results []*svc.ListEdgePackagingJobsOutput
				p := svc.NewListEdgePackagingJobsPaginator(client, input)
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
		"list-endpoint-configs": {
			Name:   "list-endpoint-configs",
			Fields: fields_list_endpoint_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoint_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpointConfigs(ctx, input)
				}
				var results []*svc.ListEndpointConfigsOutput
				p := svc.NewListEndpointConfigsPaginator(client, input)
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
		"list-endpoints": {
			Name:   "list-endpoints",
			Fields: fields_list_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpoints(ctx, input)
				}
				var results []*svc.ListEndpointsOutput
				p := svc.NewListEndpointsPaginator(client, input)
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
		"list-experiments": {
			Name:   "list-experiments",
			Fields: fields_list_experiments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperimentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_experiments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExperiments(ctx, input)
				}
				var results []*svc.ListExperimentsOutput
				p := svc.NewListExperimentsPaginator(client, input)
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
		"list-feature-groups": {
			Name:   "list-feature-groups",
			Fields: fields_list_feature_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFeatureGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_feature_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFeatureGroups(ctx, input)
				}
				var results []*svc.ListFeatureGroupsOutput
				p := svc.NewListFeatureGroupsPaginator(client, input)
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
		"list-flow-definitions": {
			Name:   "list-flow-definitions",
			Fields: fields_list_flow_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowDefinitions(ctx, input)
				}
				var results []*svc.ListFlowDefinitionsOutput
				p := svc.NewListFlowDefinitionsPaginator(client, input)
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
		"list-hub-content-versions": {
			Name:   "list-hub-content-versions",
			Fields: fields_list_hub_content_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHubContentVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hub_content_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHubContentVersions(ctx, input)
			},
		},
		"list-hub-contents": {
			Name:   "list-hub-contents",
			Fields: fields_list_hub_contents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHubContentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hub_contents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHubContents(ctx, input)
			},
		},
		"list-hubs": {
			Name:   "list-hubs",
			Fields: fields_list_hubs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHubsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hubs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHubs(ctx, input)
			},
		},
		"list-human-task-uis": {
			Name:   "list-human-task-uis",
			Fields: fields_list_human_task_uis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHumanTaskUisInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_human_task_uis, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHumanTaskUis(ctx, input)
				}
				var results []*svc.ListHumanTaskUisOutput
				p := svc.NewListHumanTaskUisPaginator(client, input)
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
		"list-hyper-parameter-tuning-jobs": {
			Name:   "list-hyper-parameter-tuning-jobs",
			Fields: fields_list_hyper_parameter_tuning_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHyperParameterTuningJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hyper_parameter_tuning_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHyperParameterTuningJobs(ctx, input)
				}
				var results []*svc.ListHyperParameterTuningJobsOutput
				p := svc.NewListHyperParameterTuningJobsPaginator(client, input)
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
		"list-image-versions": {
			Name:   "list-image-versions",
			Fields: fields_list_image_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImageVersions(ctx, input)
				}
				var results []*svc.ListImageVersionsOutput
				p := svc.NewListImageVersionsPaginator(client, input)
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
		"list-images": {
			Name:   "list-images",
			Fields: fields_list_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImages(ctx, input)
				}
				var results []*svc.ListImagesOutput
				p := svc.NewListImagesPaginator(client, input)
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
		"list-inference-components": {
			Name:   "list-inference-components",
			Fields: fields_list_inference_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceComponents(ctx, input)
				}
				var results []*svc.ListInferenceComponentsOutput
				p := svc.NewListInferenceComponentsPaginator(client, input)
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
		"list-inference-experiments": {
			Name:   "list-inference-experiments",
			Fields: fields_list_inference_experiments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceExperimentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_experiments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceExperiments(ctx, input)
				}
				var results []*svc.ListInferenceExperimentsOutput
				p := svc.NewListInferenceExperimentsPaginator(client, input)
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
		"list-inference-recommendations-job-steps": {
			Name:   "list-inference-recommendations-job-steps",
			Fields: fields_list_inference_recommendations_job_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceRecommendationsJobStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_recommendations_job_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceRecommendationsJobSteps(ctx, input)
				}
				var results []*svc.ListInferenceRecommendationsJobStepsOutput
				p := svc.NewListInferenceRecommendationsJobStepsPaginator(client, input)
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
		"list-inference-recommendations-jobs": {
			Name:   "list-inference-recommendations-jobs",
			Fields: fields_list_inference_recommendations_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInferenceRecommendationsJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inference_recommendations_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInferenceRecommendationsJobs(ctx, input)
				}
				var results []*svc.ListInferenceRecommendationsJobsOutput
				p := svc.NewListInferenceRecommendationsJobsPaginator(client, input)
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
		"list-labeling-jobs": {
			Name:   "list-labeling-jobs",
			Fields: fields_list_labeling_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLabelingJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_labeling_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLabelingJobs(ctx, input)
				}
				var results []*svc.ListLabelingJobsOutput
				p := svc.NewListLabelingJobsPaginator(client, input)
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
		"list-labeling-jobs-for-workteam": {
			Name:   "list-labeling-jobs-for-workteam",
			Fields: fields_list_labeling_jobs_for_workteam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLabelingJobsForWorkteamInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_labeling_jobs_for_workteam, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLabelingJobsForWorkteam(ctx, input)
				}
				var results []*svc.ListLabelingJobsForWorkteamOutput
				p := svc.NewListLabelingJobsForWorkteamPaginator(client, input)
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
		"list-lineage-groups": {
			Name:   "list-lineage-groups",
			Fields: fields_list_lineage_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLineageGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lineage_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLineageGroups(ctx, input)
				}
				var results []*svc.ListLineageGroupsOutput
				p := svc.NewListLineageGroupsPaginator(client, input)
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
		"list-mlflow-apps": {
			Name:   "list-mlflow-apps",
			Fields: fields_list_mlflow_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMlflowAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mlflow_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMlflowApps(ctx, input)
				}
				var results []*svc.ListMlflowAppsOutput
				p := svc.NewListMlflowAppsPaginator(client, input)
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
		"list-mlflow-tracking-servers": {
			Name:   "list-mlflow-tracking-servers",
			Fields: fields_list_mlflow_tracking_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMlflowTrackingServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mlflow_tracking_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMlflowTrackingServers(ctx, input)
				}
				var results []*svc.ListMlflowTrackingServersOutput
				p := svc.NewListMlflowTrackingServersPaginator(client, input)
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
		"list-model-bias-job-definitions": {
			Name:   "list-model-bias-job-definitions",
			Fields: fields_list_model_bias_job_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelBiasJobDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_bias_job_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelBiasJobDefinitions(ctx, input)
				}
				var results []*svc.ListModelBiasJobDefinitionsOutput
				p := svc.NewListModelBiasJobDefinitionsPaginator(client, input)
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
		"list-model-card-export-jobs": {
			Name:   "list-model-card-export-jobs",
			Fields: fields_list_model_card_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelCardExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_card_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelCardExportJobs(ctx, input)
				}
				var results []*svc.ListModelCardExportJobsOutput
				p := svc.NewListModelCardExportJobsPaginator(client, input)
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
		"list-model-card-versions": {
			Name:   "list-model-card-versions",
			Fields: fields_list_model_card_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelCardVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_card_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelCardVersions(ctx, input)
				}
				var results []*svc.ListModelCardVersionsOutput
				p := svc.NewListModelCardVersionsPaginator(client, input)
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
		"list-model-cards": {
			Name:   "list-model-cards",
			Fields: fields_list_model_cards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelCardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_cards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelCards(ctx, input)
				}
				var results []*svc.ListModelCardsOutput
				p := svc.NewListModelCardsPaginator(client, input)
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
		"list-model-explainability-job-definitions": {
			Name:   "list-model-explainability-job-definitions",
			Fields: fields_list_model_explainability_job_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelExplainabilityJobDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_explainability_job_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelExplainabilityJobDefinitions(ctx, input)
				}
				var results []*svc.ListModelExplainabilityJobDefinitionsOutput
				p := svc.NewListModelExplainabilityJobDefinitionsPaginator(client, input)
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
		"list-model-metadata": {
			Name:   "list-model-metadata",
			Fields: fields_list_model_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelMetadata(ctx, input)
				}
				var results []*svc.ListModelMetadataOutput
				p := svc.NewListModelMetadataPaginator(client, input)
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
		"list-model-package-groups": {
			Name:   "list-model-package-groups",
			Fields: fields_list_model_package_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelPackageGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_package_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelPackageGroups(ctx, input)
				}
				var results []*svc.ListModelPackageGroupsOutput
				p := svc.NewListModelPackageGroupsPaginator(client, input)
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
		"list-model-packages": {
			Name:   "list-model-packages",
			Fields: fields_list_model_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelPackages(ctx, input)
				}
				var results []*svc.ListModelPackagesOutput
				p := svc.NewListModelPackagesPaginator(client, input)
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
		"list-model-quality-job-definitions": {
			Name:   "list-model-quality-job-definitions",
			Fields: fields_list_model_quality_job_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelQualityJobDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_quality_job_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelQualityJobDefinitions(ctx, input)
				}
				var results []*svc.ListModelQualityJobDefinitionsOutput
				p := svc.NewListModelQualityJobDefinitionsPaginator(client, input)
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
		"list-monitoring-alert-history": {
			Name:   "list-monitoring-alert-history",
			Fields: fields_list_monitoring_alert_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitoringAlertHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitoring_alert_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitoringAlertHistory(ctx, input)
				}
				var results []*svc.ListMonitoringAlertHistoryOutput
				p := svc.NewListMonitoringAlertHistoryPaginator(client, input)
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
		"list-monitoring-alerts": {
			Name:   "list-monitoring-alerts",
			Fields: fields_list_monitoring_alerts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitoringAlertsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitoring_alerts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitoringAlerts(ctx, input)
				}
				var results []*svc.ListMonitoringAlertsOutput
				p := svc.NewListMonitoringAlertsPaginator(client, input)
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
		"list-monitoring-executions": {
			Name:   "list-monitoring-executions",
			Fields: fields_list_monitoring_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitoringExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitoring_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitoringExecutions(ctx, input)
				}
				var results []*svc.ListMonitoringExecutionsOutput
				p := svc.NewListMonitoringExecutionsPaginator(client, input)
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
		"list-monitoring-schedules": {
			Name:   "list-monitoring-schedules",
			Fields: fields_list_monitoring_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitoringSchedulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitoring_schedules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitoringSchedules(ctx, input)
				}
				var results []*svc.ListMonitoringSchedulesOutput
				p := svc.NewListMonitoringSchedulesPaginator(client, input)
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
		"list-notebook-instance-lifecycle-configs": {
			Name:   "list-notebook-instance-lifecycle-configs",
			Fields: fields_list_notebook_instance_lifecycle_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotebookInstanceLifecycleConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notebook_instance_lifecycle_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotebookInstanceLifecycleConfigs(ctx, input)
				}
				var results []*svc.ListNotebookInstanceLifecycleConfigsOutput
				p := svc.NewListNotebookInstanceLifecycleConfigsPaginator(client, input)
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
		"list-notebook-instances": {
			Name:   "list-notebook-instances",
			Fields: fields_list_notebook_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotebookInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notebook_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotebookInstances(ctx, input)
				}
				var results []*svc.ListNotebookInstancesOutput
				p := svc.NewListNotebookInstancesPaginator(client, input)
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
		"list-optimization-jobs": {
			Name:   "list-optimization-jobs",
			Fields: fields_list_optimization_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOptimizationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_optimization_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOptimizationJobs(ctx, input)
				}
				var results []*svc.ListOptimizationJobsOutput
				p := svc.NewListOptimizationJobsPaginator(client, input)
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
		"list-partner-apps": {
			Name:   "list-partner-apps",
			Fields: fields_list_partner_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartnerAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_partner_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPartnerApps(ctx, input)
				}
				var results []*svc.ListPartnerAppsOutput
				p := svc.NewListPartnerAppsPaginator(client, input)
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
		"list-pipeline-execution-steps": {
			Name:   "list-pipeline-execution-steps",
			Fields: fields_list_pipeline_execution_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineExecutionStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_execution_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineExecutionSteps(ctx, input)
				}
				var results []*svc.ListPipelineExecutionStepsOutput
				p := svc.NewListPipelineExecutionStepsPaginator(client, input)
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
		"list-pipeline-executions": {
			Name:   "list-pipeline-executions",
			Fields: fields_list_pipeline_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineExecutions(ctx, input)
				}
				var results []*svc.ListPipelineExecutionsOutput
				p := svc.NewListPipelineExecutionsPaginator(client, input)
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
		"list-pipeline-parameters-for-execution": {
			Name:   "list-pipeline-parameters-for-execution",
			Fields: fields_list_pipeline_parameters_for_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineParametersForExecutionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_parameters_for_execution, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineParametersForExecution(ctx, input)
				}
				var results []*svc.ListPipelineParametersForExecutionOutput
				p := svc.NewListPipelineParametersForExecutionPaginator(client, input)
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
		"list-pipeline-versions": {
			Name:   "list-pipeline-versions",
			Fields: fields_list_pipeline_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineVersions(ctx, input)
				}
				var results []*svc.ListPipelineVersionsOutput
				p := svc.NewListPipelineVersionsPaginator(client, input)
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
		"list-pipelines": {
			Name:   "list-pipelines",
			Fields: fields_list_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelines(ctx, input)
				}
				var results []*svc.ListPipelinesOutput
				p := svc.NewListPipelinesPaginator(client, input)
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
		"list-processing-jobs": {
			Name:   "list-processing-jobs",
			Fields: fields_list_processing_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProcessingJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_processing_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProcessingJobs(ctx, input)
				}
				var results []*svc.ListProcessingJobsOutput
				p := svc.NewListProcessingJobsPaginator(client, input)
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
		"list-projects": {
			Name:   "list-projects",
			Fields: fields_list_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjects(ctx, input)
				}
				var results []*svc.ListProjectsOutput
				p := svc.NewListProjectsPaginator(client, input)
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
		"list-resource-catalogs": {
			Name:   "list-resource-catalogs",
			Fields: fields_list_resource_catalogs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceCatalogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_catalogs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceCatalogs(ctx, input)
				}
				var results []*svc.ListResourceCatalogsOutput
				p := svc.NewListResourceCatalogsPaginator(client, input)
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
		"list-spaces": {
			Name:   "list-spaces",
			Fields: fields_list_spaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_spaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpaces(ctx, input)
				}
				var results []*svc.ListSpacesOutput
				p := svc.NewListSpacesPaginator(client, input)
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
		"list-stage-devices": {
			Name:   "list-stage-devices",
			Fields: fields_list_stage_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStageDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stage_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStageDevices(ctx, input)
				}
				var results []*svc.ListStageDevicesOutput
				p := svc.NewListStageDevicesPaginator(client, input)
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
		"list-studio-lifecycle-configs": {
			Name:   "list-studio-lifecycle-configs",
			Fields: fields_list_studio_lifecycle_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStudioLifecycleConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_studio_lifecycle_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStudioLifecycleConfigs(ctx, input)
				}
				var results []*svc.ListStudioLifecycleConfigsOutput
				p := svc.NewListStudioLifecycleConfigsPaginator(client, input)
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
		"list-subscribed-workteams": {
			Name:   "list-subscribed-workteams",
			Fields: fields_list_subscribed_workteams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscribedWorkteamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscribed_workteams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscribedWorkteams(ctx, input)
				}
				var results []*svc.ListSubscribedWorkteamsOutput
				p := svc.NewListSubscribedWorkteamsPaginator(client, input)
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
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTags(ctx, input)
				}
				var results []*svc.ListTagsOutput
				p := svc.NewListTagsPaginator(client, input)
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
		"list-training-jobs": {
			Name:   "list-training-jobs",
			Fields: fields_list_training_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainingJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_training_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainingJobs(ctx, input)
				}
				var results []*svc.ListTrainingJobsOutput
				p := svc.NewListTrainingJobsPaginator(client, input)
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
		"list-training-jobs-for-hyper-parameter-tuning-job": {
			Name:   "list-training-jobs-for-hyper-parameter-tuning-job",
			Fields: fields_list_training_jobs_for_hyper_parameter_tuning_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainingJobsForHyperParameterTuningJobInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_training_jobs_for_hyper_parameter_tuning_job, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainingJobsForHyperParameterTuningJob(ctx, input)
				}
				var results []*svc.ListTrainingJobsForHyperParameterTuningJobOutput
				p := svc.NewListTrainingJobsForHyperParameterTuningJobPaginator(client, input)
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
		"list-training-plans": {
			Name:   "list-training-plans",
			Fields: fields_list_training_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrainingPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_training_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrainingPlans(ctx, input)
				}
				var results []*svc.ListTrainingPlansOutput
				p := svc.NewListTrainingPlansPaginator(client, input)
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
		"list-transform-jobs": {
			Name:   "list-transform-jobs",
			Fields: fields_list_transform_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTransformJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_transform_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTransformJobs(ctx, input)
				}
				var results []*svc.ListTransformJobsOutput
				p := svc.NewListTransformJobsPaginator(client, input)
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
		"list-trial-components": {
			Name:   "list-trial-components",
			Fields: fields_list_trial_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrialComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trial_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrialComponents(ctx, input)
				}
				var results []*svc.ListTrialComponentsOutput
				p := svc.NewListTrialComponentsPaginator(client, input)
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
		"list-trials": {
			Name:   "list-trials",
			Fields: fields_list_trials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrialsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trials, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrials(ctx, input)
				}
				var results []*svc.ListTrialsOutput
				p := svc.NewListTrialsPaginator(client, input)
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
		"list-ultra-servers-by-reserved-capacity": {
			Name:   "list-ultra-servers-by-reserved-capacity",
			Fields: fields_list_ultra_servers_by_reserved_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUltraServersByReservedCapacityInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ultra_servers_by_reserved_capacity, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUltraServersByReservedCapacity(ctx, input)
				}
				var results []*svc.ListUltraServersByReservedCapacityOutput
				p := svc.NewListUltraServersByReservedCapacityPaginator(client, input)
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
		"list-user-profiles": {
			Name:   "list-user-profiles",
			Fields: fields_list_user_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserProfiles(ctx, input)
				}
				var results []*svc.ListUserProfilesOutput
				p := svc.NewListUserProfilesPaginator(client, input)
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
		"list-workforces": {
			Name:   "list-workforces",
			Fields: fields_list_workforces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkforcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workforces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkforces(ctx, input)
				}
				var results []*svc.ListWorkforcesOutput
				p := svc.NewListWorkforcesPaginator(client, input)
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
		"list-workteams": {
			Name:   "list-workteams",
			Fields: fields_list_workteams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkteamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workteams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkteams(ctx, input)
				}
				var results []*svc.ListWorkteamsOutput
				p := svc.NewListWorkteamsPaginator(client, input)
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
		"put-model-package-group-policy": {
			Name:   "put-model-package-group-policy",
			Fields: fields_put_model_package_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutModelPackageGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_model_package_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutModelPackageGroupPolicy(ctx, input)
			},
		},
		"query-lineage": {
			Name:   "query-lineage",
			Fields: fields_query_lineage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryLineageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_query_lineage, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.QueryLineage(ctx, input)
				}
				var results []*svc.QueryLineageOutput
				p := svc.NewQueryLineagePaginator(client, input)
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
		"register-devices": {
			Name:   "register-devices",
			Fields: fields_register_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDevicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_devices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDevices(ctx, input)
			},
		},
		"render-ui-template": {
			Name:   "render-ui-template",
			Fields: fields_render_ui_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenderUiTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_render_ui_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenderUiTemplate(ctx, input)
			},
		},
		"retry-pipeline-execution": {
			Name:   "retry-pipeline-execution",
			Fields: fields_retry_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryPipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryPipelineExecution(ctx, input)
			},
		},
		"search": {
			Name:   "search",
			Fields: fields_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Search(ctx, input)
				}
				var results []*svc.SearchOutput
				p := svc.NewSearchPaginator(client, input)
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
		"search-training-plan-offerings": {
			Name:   "search-training-plan-offerings",
			Fields: fields_search_training_plan_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTrainingPlanOfferingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_training_plan_offerings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchTrainingPlanOfferings(ctx, input)
			},
		},
		"send-pipeline-execution-step-failure": {
			Name:   "send-pipeline-execution-step-failure",
			Fields: fields_send_pipeline_execution_step_failure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendPipelineExecutionStepFailureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_pipeline_execution_step_failure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendPipelineExecutionStepFailure(ctx, input)
			},
		},
		"send-pipeline-execution-step-success": {
			Name:   "send-pipeline-execution-step-success",
			Fields: fields_send_pipeline_execution_step_success,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendPipelineExecutionStepSuccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_pipeline_execution_step_success, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendPipelineExecutionStepSuccess(ctx, input)
			},
		},
		"start-edge-deployment-stage": {
			Name:   "start-edge-deployment-stage",
			Fields: fields_start_edge_deployment_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEdgeDeploymentStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_edge_deployment_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEdgeDeploymentStage(ctx, input)
			},
		},
		"start-inference-experiment": {
			Name:   "start-inference-experiment",
			Fields: fields_start_inference_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInferenceExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_inference_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInferenceExperiment(ctx, input)
			},
		},
		"start-mlflow-tracking-server": {
			Name:   "start-mlflow-tracking-server",
			Fields: fields_start_mlflow_tracking_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMlflowTrackingServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_mlflow_tracking_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMlflowTrackingServer(ctx, input)
			},
		},
		"start-monitoring-schedule": {
			Name:   "start-monitoring-schedule",
			Fields: fields_start_monitoring_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMonitoringScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_monitoring_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMonitoringSchedule(ctx, input)
			},
		},
		"start-notebook-instance": {
			Name:   "start-notebook-instance",
			Fields: fields_start_notebook_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartNotebookInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_notebook_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartNotebookInstance(ctx, input)
			},
		},
		"start-pipeline-execution": {
			Name:   "start-pipeline-execution",
			Fields: fields_start_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPipelineExecution(ctx, input)
			},
		},
		"start-session": {
			Name:   "start-session",
			Fields: fields_start_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSession(ctx, input)
			},
		},
		"stop-auto-ml-job": {
			Name:   "stop-auto-ml-job",
			Fields: fields_stop_auto_ml_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAutoMLJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_auto_ml_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAutoMLJob(ctx, input)
			},
		},
		"stop-compilation-job": {
			Name:   "stop-compilation-job",
			Fields: fields_stop_compilation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCompilationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_compilation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCompilationJob(ctx, input)
			},
		},
		"stop-edge-deployment-stage": {
			Name:   "stop-edge-deployment-stage",
			Fields: fields_stop_edge_deployment_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEdgeDeploymentStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_edge_deployment_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEdgeDeploymentStage(ctx, input)
			},
		},
		"stop-edge-packaging-job": {
			Name:   "stop-edge-packaging-job",
			Fields: fields_stop_edge_packaging_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEdgePackagingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_edge_packaging_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEdgePackagingJob(ctx, input)
			},
		},
		"stop-hyper-parameter-tuning-job": {
			Name:   "stop-hyper-parameter-tuning-job",
			Fields: fields_stop_hyper_parameter_tuning_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopHyperParameterTuningJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_hyper_parameter_tuning_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopHyperParameterTuningJob(ctx, input)
			},
		},
		"stop-inference-experiment": {
			Name:   "stop-inference-experiment",
			Fields: fields_stop_inference_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopInferenceExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_inference_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopInferenceExperiment(ctx, input)
			},
		},
		"stop-inference-recommendations-job": {
			Name:   "stop-inference-recommendations-job",
			Fields: fields_stop_inference_recommendations_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopInferenceRecommendationsJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_inference_recommendations_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopInferenceRecommendationsJob(ctx, input)
			},
		},
		"stop-labeling-job": {
			Name:   "stop-labeling-job",
			Fields: fields_stop_labeling_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopLabelingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_labeling_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopLabelingJob(ctx, input)
			},
		},
		"stop-mlflow-tracking-server": {
			Name:   "stop-mlflow-tracking-server",
			Fields: fields_stop_mlflow_tracking_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMlflowTrackingServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_mlflow_tracking_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMlflowTrackingServer(ctx, input)
			},
		},
		"stop-monitoring-schedule": {
			Name:   "stop-monitoring-schedule",
			Fields: fields_stop_monitoring_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMonitoringScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_monitoring_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMonitoringSchedule(ctx, input)
			},
		},
		"stop-notebook-instance": {
			Name:   "stop-notebook-instance",
			Fields: fields_stop_notebook_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopNotebookInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_notebook_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopNotebookInstance(ctx, input)
			},
		},
		"stop-optimization-job": {
			Name:   "stop-optimization-job",
			Fields: fields_stop_optimization_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopOptimizationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_optimization_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopOptimizationJob(ctx, input)
			},
		},
		"stop-pipeline-execution": {
			Name:   "stop-pipeline-execution",
			Fields: fields_stop_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopPipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopPipelineExecution(ctx, input)
			},
		},
		"stop-processing-job": {
			Name:   "stop-processing-job",
			Fields: fields_stop_processing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopProcessingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_processing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopProcessingJob(ctx, input)
			},
		},
		"stop-training-job": {
			Name:   "stop-training-job",
			Fields: fields_stop_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTrainingJob(ctx, input)
			},
		},
		"stop-transform-job": {
			Name:   "stop-transform-job",
			Fields: fields_stop_transform_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTransformJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_transform_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTransformJob(ctx, input)
			},
		},
		"update-action": {
			Name:   "update-action",
			Fields: fields_update_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAction(ctx, input)
			},
		},
		"update-app-image-config": {
			Name:   "update-app-image-config",
			Fields: fields_update_app_image_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppImageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_image_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppImageConfig(ctx, input)
			},
		},
		"update-artifact": {
			Name:   "update-artifact",
			Fields: fields_update_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateArtifact(ctx, input)
			},
		},
		"update-cluster": {
			Name:   "update-cluster",
			Fields: fields_update_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCluster(ctx, input)
			},
		},
		"update-cluster-scheduler-config": {
			Name:   "update-cluster-scheduler-config",
			Fields: fields_update_cluster_scheduler_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterSchedulerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_scheduler_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterSchedulerConfig(ctx, input)
			},
		},
		"update-cluster-software": {
			Name:   "update-cluster-software",
			Fields: fields_update_cluster_software,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterSoftwareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_software, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterSoftware(ctx, input)
			},
		},
		"update-code-repository": {
			Name:   "update-code-repository",
			Fields: fields_update_code_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCodeRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_code_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCodeRepository(ctx, input)
			},
		},
		"update-compute-quota": {
			Name:   "update-compute-quota",
			Fields: fields_update_compute_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComputeQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_compute_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComputeQuota(ctx, input)
			},
		},
		"update-context": {
			Name:   "update-context",
			Fields: fields_update_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContext(ctx, input)
			},
		},
		"update-device-fleet": {
			Name:   "update-device-fleet",
			Fields: fields_update_device_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeviceFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeviceFleet(ctx, input)
			},
		},
		"update-devices": {
			Name:   "update-devices",
			Fields: fields_update_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDevicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_devices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDevices(ctx, input)
			},
		},
		"update-domain": {
			Name:   "update-domain",
			Fields: fields_update_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomain(ctx, input)
			},
		},
		"update-endpoint": {
			Name:   "update-endpoint",
			Fields: fields_update_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpoint(ctx, input)
			},
		},
		"update-endpoint-weights-and-capacities": {
			Name:   "update-endpoint-weights-and-capacities",
			Fields: fields_update_endpoint_weights_and_capacities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointWeightsAndCapacitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint_weights_and_capacities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpointWeightsAndCapacities(ctx, input)
			},
		},
		"update-experiment": {
			Name:   "update-experiment",
			Fields: fields_update_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExperiment(ctx, input)
			},
		},
		"update-feature-group": {
			Name:   "update-feature-group",
			Fields: fields_update_feature_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFeatureGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_feature_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFeatureGroup(ctx, input)
			},
		},
		"update-feature-metadata": {
			Name:   "update-feature-metadata",
			Fields: fields_update_feature_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFeatureMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_feature_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFeatureMetadata(ctx, input)
			},
		},
		"update-hub": {
			Name:   "update-hub",
			Fields: fields_update_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHub(ctx, input)
			},
		},
		"update-hub-content": {
			Name:   "update-hub-content",
			Fields: fields_update_hub_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHubContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hub_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHubContent(ctx, input)
			},
		},
		"update-hub-content-reference": {
			Name:   "update-hub-content-reference",
			Fields: fields_update_hub_content_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHubContentReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hub_content_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHubContentReference(ctx, input)
			},
		},
		"update-image": {
			Name:   "update-image",
			Fields: fields_update_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImage(ctx, input)
			},
		},
		"update-image-version": {
			Name:   "update-image-version",
			Fields: fields_update_image_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImageVersion(ctx, input)
			},
		},
		"update-inference-component": {
			Name:   "update-inference-component",
			Fields: fields_update_inference_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInferenceComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_inference_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInferenceComponent(ctx, input)
			},
		},
		"update-inference-component-runtime-config": {
			Name:   "update-inference-component-runtime-config",
			Fields: fields_update_inference_component_runtime_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInferenceComponentRuntimeConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_inference_component_runtime_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInferenceComponentRuntimeConfig(ctx, input)
			},
		},
		"update-inference-experiment": {
			Name:   "update-inference-experiment",
			Fields: fields_update_inference_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInferenceExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_inference_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInferenceExperiment(ctx, input)
			},
		},
		"update-mlflow-app": {
			Name:   "update-mlflow-app",
			Fields: fields_update_mlflow_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMlflowAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mlflow_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMlflowApp(ctx, input)
			},
		},
		"update-mlflow-tracking-server": {
			Name:   "update-mlflow-tracking-server",
			Fields: fields_update_mlflow_tracking_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMlflowTrackingServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mlflow_tracking_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMlflowTrackingServer(ctx, input)
			},
		},
		"update-model-card": {
			Name:   "update-model-card",
			Fields: fields_update_model_card,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelCardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model_card, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModelCard(ctx, input)
			},
		},
		"update-model-package": {
			Name:   "update-model-package",
			Fields: fields_update_model_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModelPackage(ctx, input)
			},
		},
		"update-monitoring-alert": {
			Name:   "update-monitoring-alert",
			Fields: fields_update_monitoring_alert,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMonitoringAlertInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_monitoring_alert, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMonitoringAlert(ctx, input)
			},
		},
		"update-monitoring-schedule": {
			Name:   "update-monitoring-schedule",
			Fields: fields_update_monitoring_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMonitoringScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_monitoring_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMonitoringSchedule(ctx, input)
			},
		},
		"update-notebook-instance": {
			Name:   "update-notebook-instance",
			Fields: fields_update_notebook_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotebookInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notebook_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotebookInstance(ctx, input)
			},
		},
		"update-notebook-instance-lifecycle-config": {
			Name:   "update-notebook-instance-lifecycle-config",
			Fields: fields_update_notebook_instance_lifecycle_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotebookInstanceLifecycleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notebook_instance_lifecycle_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotebookInstanceLifecycleConfig(ctx, input)
			},
		},
		"update-partner-app": {
			Name:   "update-partner-app",
			Fields: fields_update_partner_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePartnerAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_partner_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePartnerApp(ctx, input)
			},
		},
		"update-pipeline": {
			Name:   "update-pipeline",
			Fields: fields_update_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePipeline(ctx, input)
			},
		},
		"update-pipeline-execution": {
			Name:   "update-pipeline-execution",
			Fields: fields_update_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePipelineExecution(ctx, input)
			},
		},
		"update-pipeline-version": {
			Name:   "update-pipeline-version",
			Fields: fields_update_pipeline_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePipelineVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pipeline_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePipelineVersion(ctx, input)
			},
		},
		"update-project": {
			Name:   "update-project",
			Fields: fields_update_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProject(ctx, input)
			},
		},
		"update-space": {
			Name:   "update-space",
			Fields: fields_update_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSpace(ctx, input)
			},
		},
		"update-training-job": {
			Name:   "update-training-job",
			Fields: fields_update_training_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrainingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_training_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrainingJob(ctx, input)
			},
		},
		"update-trial": {
			Name:   "update-trial",
			Fields: fields_update_trial,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trial, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrial(ctx, input)
			},
		},
		"update-trial-component": {
			Name:   "update-trial-component",
			Fields: fields_update_trial_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrialComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trial_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrialComponent(ctx, input)
			},
		},
		"update-user-profile": {
			Name:   "update-user-profile",
			Fields: fields_update_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserProfile(ctx, input)
			},
		},
		"update-workforce": {
			Name:   "update-workforce",
			Fields: fields_update_workforce,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkforceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workforce, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkforce(ctx, input)
			},
		},
		"update-workteam": {
			Name:   "update-workteam",
			Fields: fields_update_workteam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkteamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workteam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkteam(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sagemaker", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
