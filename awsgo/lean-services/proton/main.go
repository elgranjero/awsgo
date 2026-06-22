package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/proton"
)

var fields_accept_environment_account_connection = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_cancel_component_deployment = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
}

var fields_cancel_environment_deployment = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_cancel_service_instance_deployment = []leanruntime.Field{
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_cancel_service_pipeline_deployment = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_create_component = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "Manifest", Flag: "manifest", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "ServiceSpec", Flag: "service-spec", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateFile", Flag: "template-file", Type: "*string", Required: true},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "CodebuildRoleArn", Flag: "codebuild-role-arn", Type: "*string", Required: false},
	{Name: "ComponentRoleArn", Flag: "component-role-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentAccountConnectionId", Flag: "environment-account-connection-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProtonServiceRoleArn", Flag: "proton-service-role-arn", Type: "*string", Required: false},
	{Name: "ProvisioningRepository", Flag: "provisioning-repository", Type: "*types.RepositoryBranchInput", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateMajorVersion", Flag: "template-major-version", Type: "*string", Required: true},
	{Name: "TemplateMinorVersion", Flag: "template-minor-version", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_environment_account_connection = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CodebuildRoleArn", Flag: "codebuild-role-arn", Type: "*string", Required: false},
	{Name: "ComponentRoleArn", Flag: "component-role-arn", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "ManagementAccountId", Flag: "management-account-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_environment_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Provisioning", Flag: "provisioning", Type: "types.Provisioning", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_environment_template_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "types.TemplateVersionSourceInput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_repository = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: true},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_service = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RepositoryConnectionArn", Flag: "repository-connection-arn", Type: "*string", Required: false},
	{Name: "RepositoryId", Flag: "repository-id", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateMajorVersion", Flag: "template-major-version", Type: "*string", Required: true},
	{Name: "TemplateMinorVersion", Flag: "template-minor-version", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_service_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateMajorVersion", Flag: "template-major-version", Type: "*string", Required: false},
	{Name: "TemplateMinorVersion", Flag: "template-minor-version", Type: "*string", Required: false},
}

var fields_create_service_sync_config = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "RepositoryProvider", Flag: "repository-provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_create_service_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PipelineProvisioning", Flag: "pipeline-provisioning", Type: "types.Provisioning", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_service_template_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CompatibleEnvironmentTemplates", Flag: "compatible-environment-templates", Type: "[]types.CompatibleEnvironmentTemplateInput", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "types.TemplateVersionSourceInput", Required: true},
	{Name: "SupportedComponentSources", Flag: "supported-component-sources", Type: "[]types.ServiceTemplateSupportedComponentSourceType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_template_sync_config = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "RepositoryProvider", Flag: "repository-provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.TemplateType", Required: true},
}

var fields_delete_component = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_deployment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_environment_account_connection = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_environment_template = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_environment_template_version = []leanruntime.Field{
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_repository = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "types.RepositoryProvider", Required: true},
}

var fields_delete_service = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_service_sync_config = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_delete_service_template = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_service_template_version = []leanruntime.Field{
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_template_sync_config = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.TemplateType", Required: true},
}

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_component = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_environment_account_connection = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_environment_template = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_environment_template_version = []leanruntime.Field{
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_get_repository = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "types.RepositoryProvider", Required: true},
}

var fields_get_repository_sync_status = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "RepositoryProvider", Flag: "repository-provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncType", Required: true},
}

var fields_get_resources_summary = []leanruntime.Field{}

var fields_get_service = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_service_instance = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_service_instance_sync_status = []leanruntime.Field{
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_service_sync_blocker_summary = []leanruntime.Field{
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_service_sync_config = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_service_template = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_service_template_version = []leanruntime.Field{
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_get_template_sync_config = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.TemplateType", Required: true},
}

var fields_get_template_sync_status = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.TemplateType", Required: true},
	{Name: "TemplateVersion", Flag: "template-version", Type: "*string", Required: true},
}

var fields_list_component_outputs = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_component_provisioned_resources = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_components = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
}

var fields_list_environment_account_connections = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RequestedBy", Flag: "requested-by", Type: "types.EnvironmentAccountConnectionRequesterAccountType", Required: true},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.EnvironmentAccountConnectionStatus", Required: false},
}

var fields_list_environment_outputs = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_provisioned_resources = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_template_versions = []leanruntime.Field{
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_list_environment_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "EnvironmentTemplates", Flag: "environment-templates", Type: "[]types.EnvironmentTemplateFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_repositories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_repository_sync_definitions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "RepositoryProvider", Flag: "repository-provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncType", Required: true},
}

var fields_list_service_instance_outputs = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_list_service_instance_provisioned_resources = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_list_service_instances = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListServiceInstancesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListServiceInstancesSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_service_pipeline_outputs = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_list_service_pipeline_provisioned_resources = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_list_service_template_versions = []leanruntime.Field{
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_list_service_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_services = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_notify_resource_deployment_status_change = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.Output", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ResourceDeploymentStatus", Required: false},
	{Name: "StatusMessage", Flag: "status-message", Type: "*string", Required: false},
}

var fields_reject_environment_account_connection = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "DeletePipelineProvisioningRepository", Flag: "delete-pipeline-provisioning-repository", Type: "*bool", Required: false},
	{Name: "PipelineCodebuildRoleArn", Flag: "pipeline-codebuild-role-arn", Type: "*string", Required: false},
	{Name: "PipelineProvisioningRepository", Flag: "pipeline-provisioning-repository", Type: "*types.RepositoryBranchInput", Required: false},
	{Name: "PipelineServiceRoleArn", Flag: "pipeline-service-role-arn", Type: "*string", Required: false},
}

var fields_update_component = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.ComponentDeploymentUpdateType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceInstanceName", Flag: "service-instance-name", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "ServiceSpec", Flag: "service-spec", Type: "*string", Required: false},
	{Name: "TemplateFile", Flag: "template-file", Type: "*string", Required: false},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "CodebuildRoleArn", Flag: "codebuild-role-arn", Type: "*string", Required: false},
	{Name: "ComponentRoleArn", Flag: "component-role-arn", Type: "*string", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.DeploymentUpdateType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentAccountConnectionId", Flag: "environment-account-connection-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProtonServiceRoleArn", Flag: "proton-service-role-arn", Type: "*string", Required: false},
	{Name: "ProvisioningRepository", Flag: "provisioning-repository", Type: "*types.RepositoryBranchInput", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: false},
	{Name: "TemplateMajorVersion", Flag: "template-major-version", Type: "*string", Required: false},
	{Name: "TemplateMinorVersion", Flag: "template-minor-version", Type: "*string", Required: false},
}

var fields_update_environment_account_connection = []leanruntime.Field{
	{Name: "CodebuildRoleArn", Flag: "codebuild-role-arn", Type: "*string", Required: false},
	{Name: "ComponentRoleArn", Flag: "component-role-arn", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_environment_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_environment_template_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.TemplateVersionStatus", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_update_service = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: false},
}

var fields_update_service_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.DeploymentUpdateType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: false},
	{Name: "TemplateMajorVersion", Flag: "template-major-version", Type: "*string", Required: false},
	{Name: "TemplateMinorVersion", Flag: "template-minor-version", Type: "*string", Required: false},
}

var fields_update_service_pipeline = []leanruntime.Field{
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.DeploymentUpdateType", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*string", Required: true},
	{Name: "TemplateMajorVersion", Flag: "template-major-version", Type: "*string", Required: false},
	{Name: "TemplateMinorVersion", Flag: "template-minor-version", Type: "*string", Required: false},
}

var fields_update_service_sync_blocker = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ResolvedReason", Flag: "resolved-reason", Type: "*string", Required: true},
}

var fields_update_service_sync_config = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "RepositoryProvider", Flag: "repository-provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_update_service_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_service_template_version = []leanruntime.Field{
	{Name: "CompatibleEnvironmentTemplates", Flag: "compatible-environment-templates", Type: "[]types.CompatibleEnvironmentTemplateInput", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MajorVersion", Flag: "major-version", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.TemplateVersionStatus", Required: false},
	{Name: "SupportedComponentSources", Flag: "supported-component-sources", Type: "[]types.ServiceTemplateSupportedComponentSourceType", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_update_template_sync_config = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "RepositoryProvider", Flag: "repository-provider", Type: "types.RepositoryProvider", Required: true},
	{Name: "Subdirectory", Flag: "subdirectory", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.TemplateType", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-environment-account-connection": {
			Name:   "accept-environment-account-connection",
			Fields: fields_accept_environment_account_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptEnvironmentAccountConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_environment_account_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptEnvironmentAccountConnection(ctx, input)
			},
		},
		"cancel-component-deployment": {
			Name:   "cancel-component-deployment",
			Fields: fields_cancel_component_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelComponentDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_component_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelComponentDeployment(ctx, input)
			},
		},
		"cancel-environment-deployment": {
			Name:   "cancel-environment-deployment",
			Fields: fields_cancel_environment_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelEnvironmentDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_environment_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelEnvironmentDeployment(ctx, input)
			},
		},
		"cancel-service-instance-deployment": {
			Name:   "cancel-service-instance-deployment",
			Fields: fields_cancel_service_instance_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelServiceInstanceDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_service_instance_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelServiceInstanceDeployment(ctx, input)
			},
		},
		"cancel-service-pipeline-deployment": {
			Name:   "cancel-service-pipeline-deployment",
			Fields: fields_cancel_service_pipeline_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelServicePipelineDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_service_pipeline_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelServicePipelineDeployment(ctx, input)
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
		"create-environment": {
			Name:   "create-environment",
			Fields: fields_create_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironment(ctx, input)
			},
		},
		"create-environment-account-connection": {
			Name:   "create-environment-account-connection",
			Fields: fields_create_environment_account_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentAccountConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_account_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentAccountConnection(ctx, input)
			},
		},
		"create-environment-template": {
			Name:   "create-environment-template",
			Fields: fields_create_environment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentTemplate(ctx, input)
			},
		},
		"create-environment-template-version": {
			Name:   "create-environment-template-version",
			Fields: fields_create_environment_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentTemplateVersion(ctx, input)
			},
		},
		"create-repository": {
			Name:   "create-repository",
			Fields: fields_create_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRepository(ctx, input)
			},
		},
		"create-service": {
			Name:   "create-service",
			Fields: fields_create_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateService(ctx, input)
			},
		},
		"create-service-instance": {
			Name:   "create-service-instance",
			Fields: fields_create_service_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceInstance(ctx, input)
			},
		},
		"create-service-sync-config": {
			Name:   "create-service-sync-config",
			Fields: fields_create_service_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceSyncConfig(ctx, input)
			},
		},
		"create-service-template": {
			Name:   "create-service-template",
			Fields: fields_create_service_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceTemplate(ctx, input)
			},
		},
		"create-service-template-version": {
			Name:   "create-service-template-version",
			Fields: fields_create_service_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceTemplateVersion(ctx, input)
			},
		},
		"create-template-sync-config": {
			Name:   "create-template-sync-config",
			Fields: fields_create_template_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplateSyncConfig(ctx, input)
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
		"delete-deployment": {
			Name:   "delete-deployment",
			Fields: fields_delete_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeployment(ctx, input)
			},
		},
		"delete-environment": {
			Name:   "delete-environment",
			Fields: fields_delete_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironment(ctx, input)
			},
		},
		"delete-environment-account-connection": {
			Name:   "delete-environment-account-connection",
			Fields: fields_delete_environment_account_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentAccountConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_account_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentAccountConnection(ctx, input)
			},
		},
		"delete-environment-template": {
			Name:   "delete-environment-template",
			Fields: fields_delete_environment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentTemplate(ctx, input)
			},
		},
		"delete-environment-template-version": {
			Name:   "delete-environment-template-version",
			Fields: fields_delete_environment_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentTemplateVersion(ctx, input)
			},
		},
		"delete-repository": {
			Name:   "delete-repository",
			Fields: fields_delete_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepository(ctx, input)
			},
		},
		"delete-service": {
			Name:   "delete-service",
			Fields: fields_delete_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteService(ctx, input)
			},
		},
		"delete-service-sync-config": {
			Name:   "delete-service-sync-config",
			Fields: fields_delete_service_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceSyncConfig(ctx, input)
			},
		},
		"delete-service-template": {
			Name:   "delete-service-template",
			Fields: fields_delete_service_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceTemplate(ctx, input)
			},
		},
		"delete-service-template-version": {
			Name:   "delete-service-template-version",
			Fields: fields_delete_service_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceTemplateVersion(ctx, input)
			},
		},
		"delete-template-sync-config": {
			Name:   "delete-template-sync-config",
			Fields: fields_delete_template_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplateSyncConfig(ctx, input)
			},
		},
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
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
		"get-deployment": {
			Name:   "get-deployment",
			Fields: fields_get_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployment(ctx, input)
			},
		},
		"get-environment": {
			Name:   "get-environment",
			Fields: fields_get_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironment(ctx, input)
			},
		},
		"get-environment-account-connection": {
			Name:   "get-environment-account-connection",
			Fields: fields_get_environment_account_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentAccountConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_account_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentAccountConnection(ctx, input)
			},
		},
		"get-environment-template": {
			Name:   "get-environment-template",
			Fields: fields_get_environment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentTemplate(ctx, input)
			},
		},
		"get-environment-template-version": {
			Name:   "get-environment-template-version",
			Fields: fields_get_environment_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentTemplateVersion(ctx, input)
			},
		},
		"get-repository": {
			Name:   "get-repository",
			Fields: fields_get_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepository(ctx, input)
			},
		},
		"get-repository-sync-status": {
			Name:   "get-repository-sync-status",
			Fields: fields_get_repository_sync_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositorySyncStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_sync_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositorySyncStatus(ctx, input)
			},
		},
		"get-resources-summary": {
			Name:   "get-resources-summary",
			Fields: fields_get_resources_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resources_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcesSummary(ctx, input)
			},
		},
		"get-service": {
			Name:   "get-service",
			Fields: fields_get_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetService(ctx, input)
			},
		},
		"get-service-instance": {
			Name:   "get-service-instance",
			Fields: fields_get_service_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceInstance(ctx, input)
			},
		},
		"get-service-instance-sync-status": {
			Name:   "get-service-instance-sync-status",
			Fields: fields_get_service_instance_sync_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceInstanceSyncStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_instance_sync_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceInstanceSyncStatus(ctx, input)
			},
		},
		"get-service-sync-blocker-summary": {
			Name:   "get-service-sync-blocker-summary",
			Fields: fields_get_service_sync_blocker_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceSyncBlockerSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_sync_blocker_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceSyncBlockerSummary(ctx, input)
			},
		},
		"get-service-sync-config": {
			Name:   "get-service-sync-config",
			Fields: fields_get_service_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceSyncConfig(ctx, input)
			},
		},
		"get-service-template": {
			Name:   "get-service-template",
			Fields: fields_get_service_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceTemplate(ctx, input)
			},
		},
		"get-service-template-version": {
			Name:   "get-service-template-version",
			Fields: fields_get_service_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceTemplateVersion(ctx, input)
			},
		},
		"get-template-sync-config": {
			Name:   "get-template-sync-config",
			Fields: fields_get_template_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplateSyncConfig(ctx, input)
			},
		},
		"get-template-sync-status": {
			Name:   "get-template-sync-status",
			Fields: fields_get_template_sync_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateSyncStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template_sync_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplateSyncStatus(ctx, input)
			},
		},
		"list-component-outputs": {
			Name:   "list-component-outputs",
			Fields: fields_list_component_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentOutputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_component_outputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponentOutputs(ctx, input)
				}
				var results []*svc.ListComponentOutputsOutput
				p := svc.NewListComponentOutputsPaginator(client, input)
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
		"list-component-provisioned-resources": {
			Name:   "list-component-provisioned-resources",
			Fields: fields_list_component_provisioned_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentProvisionedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_component_provisioned_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponentProvisionedResources(ctx, input)
				}
				var results []*svc.ListComponentProvisionedResourcesOutput
				p := svc.NewListComponentProvisionedResourcesPaginator(client, input)
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
		"list-deployments": {
			Name:   "list-deployments",
			Fields: fields_list_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeployments(ctx, input)
				}
				var results []*svc.ListDeploymentsOutput
				p := svc.NewListDeploymentsPaginator(client, input)
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
		"list-environment-account-connections": {
			Name:   "list-environment-account-connections",
			Fields: fields_list_environment_account_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentAccountConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_account_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentAccountConnections(ctx, input)
				}
				var results []*svc.ListEnvironmentAccountConnectionsOutput
				p := svc.NewListEnvironmentAccountConnectionsPaginator(client, input)
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
		"list-environment-outputs": {
			Name:   "list-environment-outputs",
			Fields: fields_list_environment_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentOutputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_outputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentOutputs(ctx, input)
				}
				var results []*svc.ListEnvironmentOutputsOutput
				p := svc.NewListEnvironmentOutputsPaginator(client, input)
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
		"list-environment-provisioned-resources": {
			Name:   "list-environment-provisioned-resources",
			Fields: fields_list_environment_provisioned_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentProvisionedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_provisioned_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentProvisionedResources(ctx, input)
				}
				var results []*svc.ListEnvironmentProvisionedResourcesOutput
				p := svc.NewListEnvironmentProvisionedResourcesPaginator(client, input)
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
		"list-environment-template-versions": {
			Name:   "list-environment-template-versions",
			Fields: fields_list_environment_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentTemplateVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_template_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentTemplateVersions(ctx, input)
				}
				var results []*svc.ListEnvironmentTemplateVersionsOutput
				p := svc.NewListEnvironmentTemplateVersionsPaginator(client, input)
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
		"list-environment-templates": {
			Name:   "list-environment-templates",
			Fields: fields_list_environment_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentTemplates(ctx, input)
				}
				var results []*svc.ListEnvironmentTemplatesOutput
				p := svc.NewListEnvironmentTemplatesPaginator(client, input)
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
		"list-environments": {
			Name:   "list-environments",
			Fields: fields_list_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironments(ctx, input)
				}
				var results []*svc.ListEnvironmentsOutput
				p := svc.NewListEnvironmentsPaginator(client, input)
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
		"list-repositories": {
			Name:   "list-repositories",
			Fields: fields_list_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repositories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositories(ctx, input)
				}
				var results []*svc.ListRepositoriesOutput
				p := svc.NewListRepositoriesPaginator(client, input)
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
		"list-repository-sync-definitions": {
			Name:   "list-repository-sync-definitions",
			Fields: fields_list_repository_sync_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositorySyncDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repository_sync_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositorySyncDefinitions(ctx, input)
				}
				var results []*svc.ListRepositorySyncDefinitionsOutput
				p := svc.NewListRepositorySyncDefinitionsPaginator(client, input)
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
		"list-service-instance-outputs": {
			Name:   "list-service-instance-outputs",
			Fields: fields_list_service_instance_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceInstanceOutputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_instance_outputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceInstanceOutputs(ctx, input)
				}
				var results []*svc.ListServiceInstanceOutputsOutput
				p := svc.NewListServiceInstanceOutputsPaginator(client, input)
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
		"list-service-instance-provisioned-resources": {
			Name:   "list-service-instance-provisioned-resources",
			Fields: fields_list_service_instance_provisioned_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceInstanceProvisionedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_instance_provisioned_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceInstanceProvisionedResources(ctx, input)
				}
				var results []*svc.ListServiceInstanceProvisionedResourcesOutput
				p := svc.NewListServiceInstanceProvisionedResourcesPaginator(client, input)
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
		"list-service-instances": {
			Name:   "list-service-instances",
			Fields: fields_list_service_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceInstances(ctx, input)
				}
				var results []*svc.ListServiceInstancesOutput
				p := svc.NewListServiceInstancesPaginator(client, input)
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
		"list-service-pipeline-outputs": {
			Name:   "list-service-pipeline-outputs",
			Fields: fields_list_service_pipeline_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicePipelineOutputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_pipeline_outputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServicePipelineOutputs(ctx, input)
				}
				var results []*svc.ListServicePipelineOutputsOutput
				p := svc.NewListServicePipelineOutputsPaginator(client, input)
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
		"list-service-pipeline-provisioned-resources": {
			Name:   "list-service-pipeline-provisioned-resources",
			Fields: fields_list_service_pipeline_provisioned_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicePipelineProvisionedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_pipeline_provisioned_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServicePipelineProvisionedResources(ctx, input)
				}
				var results []*svc.ListServicePipelineProvisionedResourcesOutput
				p := svc.NewListServicePipelineProvisionedResourcesPaginator(client, input)
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
		"list-service-template-versions": {
			Name:   "list-service-template-versions",
			Fields: fields_list_service_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceTemplateVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_template_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceTemplateVersions(ctx, input)
				}
				var results []*svc.ListServiceTemplateVersionsOutput
				p := svc.NewListServiceTemplateVersionsPaginator(client, input)
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
		"list-service-templates": {
			Name:   "list-service-templates",
			Fields: fields_list_service_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceTemplates(ctx, input)
				}
				var results []*svc.ListServiceTemplatesOutput
				p := svc.NewListServiceTemplatesPaginator(client, input)
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
		"list-services": {
			Name:   "list-services",
			Fields: fields_list_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServices(ctx, input)
				}
				var results []*svc.ListServicesOutput
				p := svc.NewListServicesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"notify-resource-deployment-status-change": {
			Name:   "notify-resource-deployment-status-change",
			Fields: fields_notify_resource_deployment_status_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyResourceDeploymentStatusChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_resource_deployment_status_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyResourceDeploymentStatusChange(ctx, input)
			},
		},
		"reject-environment-account-connection": {
			Name:   "reject-environment-account-connection",
			Fields: fields_reject_environment_account_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectEnvironmentAccountConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_environment_account_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectEnvironmentAccountConnection(ctx, input)
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
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-component": {
			Name:   "update-component",
			Fields: fields_update_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComponent(ctx, input)
			},
		},
		"update-environment": {
			Name:   "update-environment",
			Fields: fields_update_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironment(ctx, input)
			},
		},
		"update-environment-account-connection": {
			Name:   "update-environment-account-connection",
			Fields: fields_update_environment_account_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentAccountConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_account_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentAccountConnection(ctx, input)
			},
		},
		"update-environment-template": {
			Name:   "update-environment-template",
			Fields: fields_update_environment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentTemplate(ctx, input)
			},
		},
		"update-environment-template-version": {
			Name:   "update-environment-template-version",
			Fields: fields_update_environment_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentTemplateVersion(ctx, input)
			},
		},
		"update-service": {
			Name:   "update-service",
			Fields: fields_update_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateService(ctx, input)
			},
		},
		"update-service-instance": {
			Name:   "update-service-instance",
			Fields: fields_update_service_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceInstance(ctx, input)
			},
		},
		"update-service-pipeline": {
			Name:   "update-service-pipeline",
			Fields: fields_update_service_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServicePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServicePipeline(ctx, input)
			},
		},
		"update-service-sync-blocker": {
			Name:   "update-service-sync-blocker",
			Fields: fields_update_service_sync_blocker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSyncBlockerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_sync_blocker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSyncBlocker(ctx, input)
			},
		},
		"update-service-sync-config": {
			Name:   "update-service-sync-config",
			Fields: fields_update_service_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSyncConfig(ctx, input)
			},
		},
		"update-service-template": {
			Name:   "update-service-template",
			Fields: fields_update_service_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceTemplate(ctx, input)
			},
		},
		"update-service-template-version": {
			Name:   "update-service-template-version",
			Fields: fields_update_service_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceTemplateVersion(ctx, input)
			},
		},
		"update-template-sync-config": {
			Name:   "update-template-sync-config",
			Fields: fields_update_template_sync_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateSyncConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template_sync_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplateSyncConfig(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("proton", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
