package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/imagebuilder"
)

var fields_cancel_image_creation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
}

var fields_cancel_lifecycle_execution = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "LifecycleExecutionId", Flag: "lifecycle-execution-id", Type: "*string", Required: true},
}

var fields_create_component = []leanruntime.Field{
	{Name: "ChangeDescription", Flag: "change-description", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Data", Flag: "data", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "SupportedOsVersions", Flag: "supported-os-versions", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: false},
}

var fields_create_container_recipe = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Components", Flag: "components", Type: "[]types.ComponentConfiguration", Required: false},
	{Name: "ContainerType", Flag: "container-type", Type: "types.ContainerType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DockerfileTemplateData", Flag: "dockerfile-template-data", Type: "*string", Required: false},
	{Name: "DockerfileTemplateUri", Flag: "dockerfile-template-uri", Type: "*string", Required: false},
	{Name: "ImageOsVersionOverride", Flag: "image-os-version-override", Type: "*string", Required: false},
	{Name: "InstanceConfiguration", Flag: "instance-configuration", Type: "*types.InstanceConfiguration", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentImage", Flag: "parent-image", Type: "*string", Required: true},
	{Name: "PlatformOverride", Flag: "platform-override", Type: "types.Platform", Required: false},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetRepository", Flag: "target-repository", Type: "*types.TargetContainerRepository", Required: true},
	{Name: "WorkingDirectory", Flag: "working-directory", Type: "*string", Required: false},
}

var fields_create_distribution_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Distributions", Flag: "distributions", Type: "[]types.Distribution", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: false},
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: false},
	{Name: "EnhancedImageMetadataEnabled", Flag: "enhanced-image-metadata-enabled", Type: "*bool", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: false},
	{Name: "ImageScanningConfiguration", Flag: "image-scanning-configuration", Type: "*types.ImageScanningConfiguration", Required: false},
	{Name: "ImageTestsConfiguration", Flag: "image-tests-configuration", Type: "*types.ImageTestsConfiguration", Required: false},
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.ImageLoggingConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Workflows", Flag: "workflows", Type: "[]types.WorkflowConfiguration", Required: false},
}

var fields_create_image_pipeline = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: false},
	{Name: "EnhancedImageMetadataEnabled", Flag: "enhanced-image-metadata-enabled", Type: "*bool", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: false},
	{Name: "ImageScanningConfiguration", Flag: "image-scanning-configuration", Type: "*types.ImageScanningConfiguration", Required: false},
	{Name: "ImageTestsConfiguration", Flag: "image-tests-configuration", Type: "*types.ImageTestsConfiguration", Required: false},
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.PipelineLoggingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.Schedule", Required: false},
	{Name: "Status", Flag: "status", Type: "types.PipelineStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Workflows", Flag: "workflows", Type: "[]types.WorkflowConfiguration", Required: false},
}

var fields_create_image_recipe = []leanruntime.Field{
	{Name: "AdditionalInstanceConfiguration", Flag: "additional-instance-configuration", Type: "*types.AdditionalInstanceConfiguration", Required: false},
	{Name: "AmiTags", Flag: "ami-tags", Type: "map[string]string", Required: false},
	{Name: "BlockDeviceMappings", Flag: "block-device-mappings", Type: "[]types.InstanceBlockDeviceMapping", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Components", Flag: "components", Type: "[]types.ComponentConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentImage", Flag: "parent-image", Type: "*string", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkingDirectory", Flag: "working-directory", Type: "*string", Required: false},
}

var fields_create_infrastructure_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceMetadataOptions", Flag: "instance-metadata-options", Type: "*types.InstanceMetadataOptions", Required: false},
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "InstanceTypes", Flag: "instance-types", Type: "[]string", Required: false},
	{Name: "KeyPair", Flag: "key-pair", Type: "*string", Required: false},
	{Name: "Logging", Flag: "logging", Type: "*types.Logging", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Placement", Flag: "placement", Type: "*types.Placement", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "map[string]string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TerminateInstanceOnFailure", Flag: "terminate-instance-on-failure", Type: "*bool", Required: false},
}

var fields_create_lifecycle_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyDetails", Flag: "policy-details", Type: "[]types.LifecyclePolicyDetail", Required: true},
	{Name: "ResourceSelection", Flag: "resource-selection", Type: "*types.LifecyclePolicyResourceSelection", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.LifecyclePolicyResourceType", Required: true},
	{Name: "Status", Flag: "status", Type: "types.LifecyclePolicyStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_workflow = []leanruntime.Field{
	{Name: "ChangeDescription", Flag: "change-description", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Data", Flag: "data", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.WorkflowType", Required: true},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: false},
}

var fields_delete_component = []leanruntime.Field{
	{Name: "ComponentBuildVersionArn", Flag: "component-build-version-arn", Type: "*string", Required: true},
}

var fields_delete_container_recipe = []leanruntime.Field{
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: true},
}

var fields_delete_distribution_configuration = []leanruntime.Field{
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_image = []leanruntime.Field{
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
}

var fields_delete_image_pipeline = []leanruntime.Field{
	{Name: "ImagePipelineArn", Flag: "image-pipeline-arn", Type: "*string", Required: true},
}

var fields_delete_image_recipe = []leanruntime.Field{
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: true},
}

var fields_delete_infrastructure_configuration = []leanruntime.Field{
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_lifecycle_policy = []leanruntime.Field{
	{Name: "LifecyclePolicyArn", Flag: "lifecycle-policy-arn", Type: "*string", Required: true},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "WorkflowBuildVersionArn", Flag: "workflow-build-version-arn", Type: "*string", Required: true},
}

var fields_distribute_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: true},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: true},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.ImageLoggingConfiguration", Required: false},
	{Name: "SourceImage", Flag: "source-image", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_get_component = []leanruntime.Field{
	{Name: "ComponentBuildVersionArn", Flag: "component-build-version-arn", Type: "*string", Required: true},
}

var fields_get_component_policy = []leanruntime.Field{
	{Name: "ComponentArn", Flag: "component-arn", Type: "*string", Required: true},
}

var fields_get_container_recipe = []leanruntime.Field{
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: true},
}

var fields_get_container_recipe_policy = []leanruntime.Field{
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: true},
}

var fields_get_distribution_configuration = []leanruntime.Field{
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: true},
}

var fields_get_image = []leanruntime.Field{
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
}

var fields_get_image_pipeline = []leanruntime.Field{
	{Name: "ImagePipelineArn", Flag: "image-pipeline-arn", Type: "*string", Required: true},
}

var fields_get_image_policy = []leanruntime.Field{
	{Name: "ImageArn", Flag: "image-arn", Type: "*string", Required: true},
}

var fields_get_image_recipe = []leanruntime.Field{
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: true},
}

var fields_get_image_recipe_policy = []leanruntime.Field{
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: true},
}

var fields_get_infrastructure_configuration = []leanruntime.Field{
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
}

var fields_get_lifecycle_execution = []leanruntime.Field{
	{Name: "LifecycleExecutionId", Flag: "lifecycle-execution-id", Type: "*string", Required: true},
}

var fields_get_lifecycle_policy = []leanruntime.Field{
	{Name: "LifecyclePolicyArn", Flag: "lifecycle-policy-arn", Type: "*string", Required: true},
}

var fields_get_marketplace_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceLocation", Flag: "resource-location", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.MarketplaceResourceType", Required: true},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "WorkflowBuildVersionArn", Flag: "workflow-build-version-arn", Type: "*string", Required: true},
}

var fields_get_workflow_execution = []leanruntime.Field{
	{Name: "WorkflowExecutionId", Flag: "workflow-execution-id", Type: "*string", Required: true},
}

var fields_get_workflow_step_execution = []leanruntime.Field{
	{Name: "StepExecutionId", Flag: "step-execution-id", Type: "*string", Required: true},
}

var fields_import_component = []leanruntime.Field{
	{Name: "ChangeDescription", Flag: "change-description", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Data", Flag: "data", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.ComponentFormat", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ComponentType", Required: true},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: false},
}

var fields_import_disk_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.ImageLoggingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OsVersion", Flag: "os-version", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "*string", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: true},
}

var fields_import_vm_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.ImageLoggingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OsVersion", Flag: "os-version", Type: "*string", Required: false},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VmImportTaskId", Flag: "vm-import-task-id", Type: "*string", Required: true},
}

var fields_list_component_build_versions = []leanruntime.Field{
	{Name: "ComponentVersionArn", Flag: "component-version-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_components = []leanruntime.Field{
	{Name: "ByName", Flag: "by-name", Type: "bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Ownership", Required: false},
}

var fields_list_container_recipes = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Ownership", Required: false},
}

var fields_list_distribution_configurations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_build_versions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImageVersionArn", Flag: "image-version-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_packages = []leanruntime.Field{
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_pipeline_images = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImagePipelineArn", Flag: "image-pipeline-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_pipelines = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_recipes = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Ownership", Required: false},
}

var fields_list_image_scan_finding_aggregations = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Filter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_scan_findings = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ImageScanFindingsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_images = []leanruntime.Field{
	{Name: "ByName", Flag: "by-name", Type: "bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeDeprecated", Flag: "include-deprecated", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Ownership", Required: false},
}

var fields_list_infrastructure_configurations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_lifecycle_execution_resources = []leanruntime.Field{
	{Name: "LifecycleExecutionId", Flag: "lifecycle-execution-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentResourceId", Flag: "parent-resource-id", Type: "*string", Required: false},
}

var fields_list_lifecycle_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_lifecycle_policies = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_waiting_workflow_steps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflow_build_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowVersionArn", Flag: "workflow-version-arn", Type: "*string", Required: false},
}

var fields_list_workflow_executions = []leanruntime.Field{
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflow_step_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowExecutionId", Flag: "workflow-execution-id", Type: "*string", Required: true},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "ByName", Flag: "by-name", Type: "bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Ownership", Required: false},
}

var fields_put_component_policy = []leanruntime.Field{
	{Name: "ComponentArn", Flag: "component-arn", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_container_recipe_policy = []leanruntime.Field{
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_image_policy = []leanruntime.Field{
	{Name: "ImageArn", Flag: "image-arn", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_image_recipe_policy = []leanruntime.Field{
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_retry_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
}

var fields_send_workflow_step_action = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.WorkflowStepActionType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ImageBuildVersionArn", Flag: "image-build-version-arn", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "StepExecutionId", Flag: "step-execution-id", Type: "*string", Required: true},
}

var fields_start_image_pipeline_execution = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ImagePipelineArn", Flag: "image-pipeline-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_resource_state_update = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ExclusionRules", Flag: "exclusion-rules", Type: "*types.ResourceStateUpdateExclusionRules", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "IncludeResources", Flag: "include-resources", Type: "*types.ResourceStateUpdateIncludeResources", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "*types.ResourceState", Required: true},
	{Name: "UpdateAt", Flag: "update-at", Type: "*time.Time", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_distribution_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: true},
	{Name: "Distributions", Flag: "distributions", Type: "[]types.Distribution", Required: true},
}

var fields_update_image_pipeline = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContainerRecipeArn", Flag: "container-recipe-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DistributionConfigurationArn", Flag: "distribution-configuration-arn", Type: "*string", Required: false},
	{Name: "EnhancedImageMetadataEnabled", Flag: "enhanced-image-metadata-enabled", Type: "*bool", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "ImagePipelineArn", Flag: "image-pipeline-arn", Type: "*string", Required: true},
	{Name: "ImageRecipeArn", Flag: "image-recipe-arn", Type: "*string", Required: false},
	{Name: "ImageScanningConfiguration", Flag: "image-scanning-configuration", Type: "*types.ImageScanningConfiguration", Required: false},
	{Name: "ImageTestsConfiguration", Flag: "image-tests-configuration", Type: "*types.ImageTestsConfiguration", Required: false},
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.PipelineLoggingConfiguration", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.Schedule", Required: false},
	{Name: "Status", Flag: "status", Type: "types.PipelineStatus", Required: false},
	{Name: "Workflows", Flag: "workflows", Type: "[]types.WorkflowConfiguration", Required: false},
}

var fields_update_infrastructure_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
	{Name: "InstanceMetadataOptions", Flag: "instance-metadata-options", Type: "*types.InstanceMetadataOptions", Required: false},
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "InstanceTypes", Flag: "instance-types", Type: "[]string", Required: false},
	{Name: "KeyPair", Flag: "key-pair", Type: "*string", Required: false},
	{Name: "Logging", Flag: "logging", Type: "*types.Logging", Required: false},
	{Name: "Placement", Flag: "placement", Type: "*types.Placement", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "map[string]string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "TerminateInstanceOnFailure", Flag: "terminate-instance-on-failure", Type: "*bool", Required: false},
}

var fields_update_lifecycle_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: true},
	{Name: "LifecyclePolicyArn", Flag: "lifecycle-policy-arn", Type: "*string", Required: true},
	{Name: "PolicyDetails", Flag: "policy-details", Type: "[]types.LifecyclePolicyDetail", Required: true},
	{Name: "ResourceSelection", Flag: "resource-selection", Type: "*types.LifecyclePolicyResourceSelection", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.LifecyclePolicyResourceType", Required: true},
	{Name: "Status", Flag: "status", Type: "types.LifecyclePolicyStatus", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-image-creation": {
			Name:   "cancel-image-creation",
			Fields: fields_cancel_image_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelImageCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_image_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelImageCreation(ctx, input)
			},
		},
		"cancel-lifecycle-execution": {
			Name:   "cancel-lifecycle-execution",
			Fields: fields_cancel_lifecycle_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelLifecycleExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_lifecycle_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelLifecycleExecution(ctx, input)
			},
		},
		"create-component": {
			Name:   "create-component",
			Fields: fields_create_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComponent(ctx, input)
			},
		},
		"create-container-recipe": {
			Name:   "create-container-recipe",
			Fields: fields_create_container_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainerRecipe(ctx, input)
			},
		},
		"create-distribution-configuration": {
			Name:   "create-distribution-configuration",
			Fields: fields_create_distribution_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDistributionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_distribution_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDistributionConfiguration(ctx, input)
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
		"create-image-pipeline": {
			Name:   "create-image-pipeline",
			Fields: fields_create_image_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImagePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImagePipeline(ctx, input)
			},
		},
		"create-image-recipe": {
			Name:   "create-image-recipe",
			Fields: fields_create_image_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImageRecipe(ctx, input)
			},
		},
		"create-infrastructure-configuration": {
			Name:   "create-infrastructure-configuration",
			Fields: fields_create_infrastructure_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInfrastructureConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_infrastructure_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInfrastructureConfiguration(ctx, input)
			},
		},
		"create-lifecycle-policy": {
			Name:   "create-lifecycle-policy",
			Fields: fields_create_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLifecyclePolicy(ctx, input)
			},
		},
		"create-workflow": {
			Name:   "create-workflow",
			Fields: fields_create_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflow(ctx, input)
			},
		},
		"delete-component": {
			Name:   "delete-component",
			Fields: fields_delete_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComponent(ctx, input)
			},
		},
		"delete-container-recipe": {
			Name:   "delete-container-recipe",
			Fields: fields_delete_container_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainerRecipe(ctx, input)
			},
		},
		"delete-distribution-configuration": {
			Name:   "delete-distribution-configuration",
			Fields: fields_delete_distribution_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDistributionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_distribution_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDistributionConfiguration(ctx, input)
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
		"delete-image-pipeline": {
			Name:   "delete-image-pipeline",
			Fields: fields_delete_image_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImagePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImagePipeline(ctx, input)
			},
		},
		"delete-image-recipe": {
			Name:   "delete-image-recipe",
			Fields: fields_delete_image_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImageRecipe(ctx, input)
			},
		},
		"delete-infrastructure-configuration": {
			Name:   "delete-infrastructure-configuration",
			Fields: fields_delete_infrastructure_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInfrastructureConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_infrastructure_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInfrastructureConfiguration(ctx, input)
			},
		},
		"delete-lifecycle-policy": {
			Name:   "delete-lifecycle-policy",
			Fields: fields_delete_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLifecyclePolicy(ctx, input)
			},
		},
		"delete-workflow": {
			Name:   "delete-workflow",
			Fields: fields_delete_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflow(ctx, input)
			},
		},
		"distribute-image": {
			Name:   "distribute-image",
			Fields: fields_distribute_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DistributeImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_distribute_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DistributeImage(ctx, input)
			},
		},
		"get-component": {
			Name:   "get-component",
			Fields: fields_get_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComponent(ctx, input)
			},
		},
		"get-component-policy": {
			Name:   "get-component-policy",
			Fields: fields_get_component_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComponentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_component_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComponentPolicy(ctx, input)
			},
		},
		"get-container-recipe": {
			Name:   "get-container-recipe",
			Fields: fields_get_container_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerRecipe(ctx, input)
			},
		},
		"get-container-recipe-policy": {
			Name:   "get-container-recipe-policy",
			Fields: fields_get_container_recipe_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerRecipePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_recipe_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerRecipePolicy(ctx, input)
			},
		},
		"get-distribution-configuration": {
			Name:   "get-distribution-configuration",
			Fields: fields_get_distribution_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionConfiguration(ctx, input)
			},
		},
		"get-image": {
			Name:   "get-image",
			Fields: fields_get_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImage(ctx, input)
			},
		},
		"get-image-pipeline": {
			Name:   "get-image-pipeline",
			Fields: fields_get_image_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImagePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImagePipeline(ctx, input)
			},
		},
		"get-image-policy": {
			Name:   "get-image-policy",
			Fields: fields_get_image_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImagePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImagePolicy(ctx, input)
			},
		},
		"get-image-recipe": {
			Name:   "get-image-recipe",
			Fields: fields_get_image_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageRecipe(ctx, input)
			},
		},
		"get-image-recipe-policy": {
			Name:   "get-image-recipe-policy",
			Fields: fields_get_image_recipe_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageRecipePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_recipe_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageRecipePolicy(ctx, input)
			},
		},
		"get-infrastructure-configuration": {
			Name:   "get-infrastructure-configuration",
			Fields: fields_get_infrastructure_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInfrastructureConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_infrastructure_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInfrastructureConfiguration(ctx, input)
			},
		},
		"get-lifecycle-execution": {
			Name:   "get-lifecycle-execution",
			Fields: fields_get_lifecycle_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecycleExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lifecycle_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLifecycleExecution(ctx, input)
			},
		},
		"get-lifecycle-policy": {
			Name:   "get-lifecycle-policy",
			Fields: fields_get_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLifecyclePolicy(ctx, input)
			},
		},
		"get-marketplace-resource": {
			Name:   "get-marketplace-resource",
			Fields: fields_get_marketplace_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMarketplaceResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_marketplace_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMarketplaceResource(ctx, input)
			},
		},
		"get-workflow": {
			Name:   "get-workflow",
			Fields: fields_get_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflow(ctx, input)
			},
		},
		"get-workflow-execution": {
			Name:   "get-workflow-execution",
			Fields: fields_get_workflow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowExecution(ctx, input)
			},
		},
		"get-workflow-step-execution": {
			Name:   "get-workflow-step-execution",
			Fields: fields_get_workflow_step_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowStepExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_step_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowStepExecution(ctx, input)
			},
		},
		"import-component": {
			Name:   "import-component",
			Fields: fields_import_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportComponent(ctx, input)
			},
		},
		"import-disk-image": {
			Name:   "import-disk-image",
			Fields: fields_import_disk_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportDiskImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_disk_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportDiskImage(ctx, input)
			},
		},
		"import-vm-image": {
			Name:   "import-vm-image",
			Fields: fields_import_vm_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportVmImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_vm_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportVmImage(ctx, input)
			},
		},
		"list-component-build-versions": {
			Name:   "list-component-build-versions",
			Fields: fields_list_component_build_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentBuildVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_component_build_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponentBuildVersions(ctx, input)
				}
				var results []*svc.ListComponentBuildVersionsOutput
				p := svc.NewListComponentBuildVersionsPaginator(client, input)
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
		"list-components": {
			Name:   "list-components",
			Fields: fields_list_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponents(ctx, input)
				}
				var results []*svc.ListComponentsOutput
				p := svc.NewListComponentsPaginator(client, input)
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
		"list-container-recipes": {
			Name:   "list-container-recipes",
			Fields: fields_list_container_recipes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContainerRecipesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_container_recipes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContainerRecipes(ctx, input)
				}
				var results []*svc.ListContainerRecipesOutput
				p := svc.NewListContainerRecipesPaginator(client, input)
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
		"list-distribution-configurations": {
			Name:   "list-distribution-configurations",
			Fields: fields_list_distribution_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributionConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_distribution_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDistributionConfigurations(ctx, input)
				}
				var results []*svc.ListDistributionConfigurationsOutput
				p := svc.NewListDistributionConfigurationsPaginator(client, input)
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
		"list-image-build-versions": {
			Name:   "list-image-build-versions",
			Fields: fields_list_image_build_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageBuildVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_build_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImageBuildVersions(ctx, input)
				}
				var results []*svc.ListImageBuildVersionsOutput
				p := svc.NewListImageBuildVersionsPaginator(client, input)
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
		"list-image-packages": {
			Name:   "list-image-packages",
			Fields: fields_list_image_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImagePackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImagePackages(ctx, input)
				}
				var results []*svc.ListImagePackagesOutput
				p := svc.NewListImagePackagesPaginator(client, input)
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
		"list-image-pipeline-images": {
			Name:   "list-image-pipeline-images",
			Fields: fields_list_image_pipeline_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImagePipelineImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_pipeline_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImagePipelineImages(ctx, input)
				}
				var results []*svc.ListImagePipelineImagesOutput
				p := svc.NewListImagePipelineImagesPaginator(client, input)
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
		"list-image-pipelines": {
			Name:   "list-image-pipelines",
			Fields: fields_list_image_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImagePipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImagePipelines(ctx, input)
				}
				var results []*svc.ListImagePipelinesOutput
				p := svc.NewListImagePipelinesPaginator(client, input)
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
		"list-image-recipes": {
			Name:   "list-image-recipes",
			Fields: fields_list_image_recipes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageRecipesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_recipes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImageRecipes(ctx, input)
				}
				var results []*svc.ListImageRecipesOutput
				p := svc.NewListImageRecipesPaginator(client, input)
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
		"list-image-scan-finding-aggregations": {
			Name:   "list-image-scan-finding-aggregations",
			Fields: fields_list_image_scan_finding_aggregations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageScanFindingAggregationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_scan_finding_aggregations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImageScanFindingAggregations(ctx, input)
				}
				var results []*svc.ListImageScanFindingAggregationsOutput
				p := svc.NewListImageScanFindingAggregationsPaginator(client, input)
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
		"list-image-scan-findings": {
			Name:   "list-image-scan-findings",
			Fields: fields_list_image_scan_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageScanFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_scan_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImageScanFindings(ctx, input)
				}
				var results []*svc.ListImageScanFindingsOutput
				p := svc.NewListImageScanFindingsPaginator(client, input)
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
		"list-infrastructure-configurations": {
			Name:   "list-infrastructure-configurations",
			Fields: fields_list_infrastructure_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInfrastructureConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_infrastructure_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInfrastructureConfigurations(ctx, input)
				}
				var results []*svc.ListInfrastructureConfigurationsOutput
				p := svc.NewListInfrastructureConfigurationsPaginator(client, input)
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
		"list-lifecycle-execution-resources": {
			Name:   "list-lifecycle-execution-resources",
			Fields: fields_list_lifecycle_execution_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLifecycleExecutionResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lifecycle_execution_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLifecycleExecutionResources(ctx, input)
				}
				var results []*svc.ListLifecycleExecutionResourcesOutput
				p := svc.NewListLifecycleExecutionResourcesPaginator(client, input)
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
		"list-lifecycle-executions": {
			Name:   "list-lifecycle-executions",
			Fields: fields_list_lifecycle_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLifecycleExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lifecycle_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLifecycleExecutions(ctx, input)
				}
				var results []*svc.ListLifecycleExecutionsOutput
				p := svc.NewListLifecycleExecutionsPaginator(client, input)
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
		"list-lifecycle-policies": {
			Name:   "list-lifecycle-policies",
			Fields: fields_list_lifecycle_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLifecyclePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lifecycle_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLifecyclePolicies(ctx, input)
				}
				var results []*svc.ListLifecyclePoliciesOutput
				p := svc.NewListLifecyclePoliciesPaginator(client, input)
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
		"list-waiting-workflow-steps": {
			Name:   "list-waiting-workflow-steps",
			Fields: fields_list_waiting_workflow_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWaitingWorkflowStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_waiting_workflow_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWaitingWorkflowSteps(ctx, input)
				}
				var results []*svc.ListWaitingWorkflowStepsOutput
				p := svc.NewListWaitingWorkflowStepsPaginator(client, input)
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
		"list-workflow-build-versions": {
			Name:   "list-workflow-build-versions",
			Fields: fields_list_workflow_build_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowBuildVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_build_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowBuildVersions(ctx, input)
				}
				var results []*svc.ListWorkflowBuildVersionsOutput
				p := svc.NewListWorkflowBuildVersionsPaginator(client, input)
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
		"list-workflow-executions": {
			Name:   "list-workflow-executions",
			Fields: fields_list_workflow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowExecutions(ctx, input)
				}
				var results []*svc.ListWorkflowExecutionsOutput
				p := svc.NewListWorkflowExecutionsPaginator(client, input)
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
		"list-workflow-step-executions": {
			Name:   "list-workflow-step-executions",
			Fields: fields_list_workflow_step_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowStepExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_step_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowStepExecutions(ctx, input)
				}
				var results []*svc.ListWorkflowStepExecutionsOutput
				p := svc.NewListWorkflowStepExecutionsPaginator(client, input)
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
		"list-workflows": {
			Name:   "list-workflows",
			Fields: fields_list_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflows(ctx, input)
				}
				var results []*svc.ListWorkflowsOutput
				p := svc.NewListWorkflowsPaginator(client, input)
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
		"put-component-policy": {
			Name:   "put-component-policy",
			Fields: fields_put_component_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutComponentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_component_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutComponentPolicy(ctx, input)
			},
		},
		"put-container-recipe-policy": {
			Name:   "put-container-recipe-policy",
			Fields: fields_put_container_recipe_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutContainerRecipePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_container_recipe_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutContainerRecipePolicy(ctx, input)
			},
		},
		"put-image-policy": {
			Name:   "put-image-policy",
			Fields: fields_put_image_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutImagePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_image_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutImagePolicy(ctx, input)
			},
		},
		"put-image-recipe-policy": {
			Name:   "put-image-recipe-policy",
			Fields: fields_put_image_recipe_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutImageRecipePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_image_recipe_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutImageRecipePolicy(ctx, input)
			},
		},
		"retry-image": {
			Name:   "retry-image",
			Fields: fields_retry_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryImage(ctx, input)
			},
		},
		"send-workflow-step-action": {
			Name:   "send-workflow-step-action",
			Fields: fields_send_workflow_step_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendWorkflowStepActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_workflow_step_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendWorkflowStepAction(ctx, input)
			},
		},
		"start-image-pipeline-execution": {
			Name:   "start-image-pipeline-execution",
			Fields: fields_start_image_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImagePipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_image_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImagePipelineExecution(ctx, input)
			},
		},
		"start-resource-state-update": {
			Name:   "start-resource-state-update",
			Fields: fields_start_resource_state_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartResourceStateUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_resource_state_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartResourceStateUpdate(ctx, input)
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
		"update-distribution-configuration": {
			Name:   "update-distribution-configuration",
			Fields: fields_update_distribution_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDistributionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_distribution_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDistributionConfiguration(ctx, input)
			},
		},
		"update-image-pipeline": {
			Name:   "update-image-pipeline",
			Fields: fields_update_image_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImagePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImagePipeline(ctx, input)
			},
		},
		"update-infrastructure-configuration": {
			Name:   "update-infrastructure-configuration",
			Fields: fields_update_infrastructure_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInfrastructureConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_infrastructure_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInfrastructureConfiguration(ctx, input)
			},
		},
		"update-lifecycle-policy": {
			Name:   "update-lifecycle-policy",
			Fields: fields_update_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLifecyclePolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("imagebuilder", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
