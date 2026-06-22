package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/greengrass"
)

var fields_associate_role_to_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_associate_service_role_to_account = []leanruntime.Field{
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_create_connector_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.ConnectorDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_connector_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "ConnectorDefinitionId", Flag: "connector-definition-id", Type: "*string", Required: true},
	{Name: "Connectors", Flag: "connectors", Type: "[]types.Connector", Required: false},
}

var fields_create_core_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.CoreDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_core_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "CoreDefinitionId", Flag: "core-definition-id", Type: "*string", Required: true},
	{Name: "Cores", Flag: "cores", Type: "[]types.Core", Required: false},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.DeploymentType", Required: true},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "GroupVersionId", Flag: "group-version-id", Type: "*string", Required: false},
}

var fields_create_device_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.DeviceDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_device_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "DeviceDefinitionId", Flag: "device-definition-id", Type: "*string", Required: true},
	{Name: "Devices", Flag: "devices", Type: "[]types.Device", Required: false},
}

var fields_create_function_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.FunctionDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_function_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "DefaultConfig", Flag: "default-config", Type: "*types.FunctionDefaultConfig", Required: false},
	{Name: "FunctionDefinitionId", Flag: "function-definition-id", Type: "*string", Required: true},
	{Name: "Functions", Flag: "functions", Type: "[]types.Function", Required: false},
}

var fields_create_group = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.GroupVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_group_certificate_authority = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_create_group_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "ConnectorDefinitionVersionArn", Flag: "connector-definition-version-arn", Type: "*string", Required: false},
	{Name: "CoreDefinitionVersionArn", Flag: "core-definition-version-arn", Type: "*string", Required: false},
	{Name: "DeviceDefinitionVersionArn", Flag: "device-definition-version-arn", Type: "*string", Required: false},
	{Name: "FunctionDefinitionVersionArn", Flag: "function-definition-version-arn", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "LoggerDefinitionVersionArn", Flag: "logger-definition-version-arn", Type: "*string", Required: false},
	{Name: "ResourceDefinitionVersionArn", Flag: "resource-definition-version-arn", Type: "*string", Required: false},
	{Name: "SubscriptionDefinitionVersionArn", Flag: "subscription-definition-version-arn", Type: "*string", Required: false},
}

var fields_create_logger_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.LoggerDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_logger_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "LoggerDefinitionId", Flag: "logger-definition-id", Type: "*string", Required: true},
	{Name: "Loggers", Flag: "loggers", Type: "[]types.Logger", Required: false},
}

var fields_create_resource_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.ResourceDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_resource_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "ResourceDefinitionId", Flag: "resource-definition-id", Type: "*string", Required: true},
	{Name: "Resources", Flag: "resources", Type: "[]types.Resource", Required: false},
}

var fields_create_software_update_job = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "S3UrlSignerRole", Flag: "s3-url-signer-role", Type: "*string", Required: true},
	{Name: "SoftwareToUpdate", Flag: "software-to-update", Type: "types.SoftwareToUpdate", Required: true},
	{Name: "UpdateAgentLogLevel", Flag: "update-agent-log-level", Type: "types.UpdateAgentLogLevel", Required: false},
	{Name: "UpdateTargets", Flag: "update-targets", Type: "[]string", Required: true},
	{Name: "UpdateTargetsArchitecture", Flag: "update-targets-architecture", Type: "types.UpdateTargetsArchitecture", Required: true},
	{Name: "UpdateTargetsOperatingSystem", Flag: "update-targets-operating-system", Type: "types.UpdateTargetsOperatingSystem", Required: true},
}

var fields_create_subscription_definition = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "InitialVersion", Flag: "initial-version", Type: "*types.SubscriptionDefinitionVersion", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_subscription_definition_version = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "SubscriptionDefinitionId", Flag: "subscription-definition-id", Type: "*string", Required: true},
	{Name: "Subscriptions", Flag: "subscriptions", Type: "[]types.Subscription", Required: false},
}

var fields_delete_connector_definition = []leanruntime.Field{
	{Name: "ConnectorDefinitionId", Flag: "connector-definition-id", Type: "*string", Required: true},
}

var fields_delete_core_definition = []leanruntime.Field{
	{Name: "CoreDefinitionId", Flag: "core-definition-id", Type: "*string", Required: true},
}

var fields_delete_device_definition = []leanruntime.Field{
	{Name: "DeviceDefinitionId", Flag: "device-definition-id", Type: "*string", Required: true},
}

var fields_delete_function_definition = []leanruntime.Field{
	{Name: "FunctionDefinitionId", Flag: "function-definition-id", Type: "*string", Required: true},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_delete_logger_definition = []leanruntime.Field{
	{Name: "LoggerDefinitionId", Flag: "logger-definition-id", Type: "*string", Required: true},
}

var fields_delete_resource_definition = []leanruntime.Field{
	{Name: "ResourceDefinitionId", Flag: "resource-definition-id", Type: "*string", Required: true},
}

var fields_delete_subscription_definition = []leanruntime.Field{
	{Name: "SubscriptionDefinitionId", Flag: "subscription-definition-id", Type: "*string", Required: true},
}

var fields_disassociate_role_from_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_disassociate_service_role_from_account = []leanruntime.Field{}

var fields_get_associated_role = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_get_bulk_deployment_status = []leanruntime.Field{
	{Name: "BulkDeploymentId", Flag: "bulk-deployment-id", Type: "*string", Required: true},
}

var fields_get_connectivity_info = []leanruntime.Field{
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_get_connector_definition = []leanruntime.Field{
	{Name: "ConnectorDefinitionId", Flag: "connector-definition-id", Type: "*string", Required: true},
}

var fields_get_connector_definition_version = []leanruntime.Field{
	{Name: "ConnectorDefinitionId", Flag: "connector-definition-id", Type: "*string", Required: true},
	{Name: "ConnectorDefinitionVersionId", Flag: "connector-definition-version-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_core_definition = []leanruntime.Field{
	{Name: "CoreDefinitionId", Flag: "core-definition-id", Type: "*string", Required: true},
}

var fields_get_core_definition_version = []leanruntime.Field{
	{Name: "CoreDefinitionId", Flag: "core-definition-id", Type: "*string", Required: true},
	{Name: "CoreDefinitionVersionId", Flag: "core-definition-version-id", Type: "*string", Required: true},
}

var fields_get_deployment_status = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_get_device_definition = []leanruntime.Field{
	{Name: "DeviceDefinitionId", Flag: "device-definition-id", Type: "*string", Required: true},
}

var fields_get_device_definition_version = []leanruntime.Field{
	{Name: "DeviceDefinitionId", Flag: "device-definition-id", Type: "*string", Required: true},
	{Name: "DeviceDefinitionVersionId", Flag: "device-definition-version-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_function_definition = []leanruntime.Field{
	{Name: "FunctionDefinitionId", Flag: "function-definition-id", Type: "*string", Required: true},
}

var fields_get_function_definition_version = []leanruntime.Field{
	{Name: "FunctionDefinitionId", Flag: "function-definition-id", Type: "*string", Required: true},
	{Name: "FunctionDefinitionVersionId", Flag: "function-definition-version-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_get_group_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityId", Flag: "certificate-authority-id", Type: "*string", Required: true},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_get_group_certificate_configuration = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_get_group_version = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "GroupVersionId", Flag: "group-version-id", Type: "*string", Required: true},
}

var fields_get_logger_definition = []leanruntime.Field{
	{Name: "LoggerDefinitionId", Flag: "logger-definition-id", Type: "*string", Required: true},
}

var fields_get_logger_definition_version = []leanruntime.Field{
	{Name: "LoggerDefinitionId", Flag: "logger-definition-id", Type: "*string", Required: true},
	{Name: "LoggerDefinitionVersionId", Flag: "logger-definition-version-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_resource_definition = []leanruntime.Field{
	{Name: "ResourceDefinitionId", Flag: "resource-definition-id", Type: "*string", Required: true},
}

var fields_get_resource_definition_version = []leanruntime.Field{
	{Name: "ResourceDefinitionId", Flag: "resource-definition-id", Type: "*string", Required: true},
	{Name: "ResourceDefinitionVersionId", Flag: "resource-definition-version-id", Type: "*string", Required: true},
}

var fields_get_service_role_for_account = []leanruntime.Field{}

var fields_get_subscription_definition = []leanruntime.Field{
	{Name: "SubscriptionDefinitionId", Flag: "subscription-definition-id", Type: "*string", Required: true},
}

var fields_get_subscription_definition_version = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubscriptionDefinitionId", Flag: "subscription-definition-id", Type: "*string", Required: true},
	{Name: "SubscriptionDefinitionVersionId", Flag: "subscription-definition-version-id", Type: "*string", Required: true},
}

var fields_get_thing_runtime_configuration = []leanruntime.Field{
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_bulk_deployment_detailed_reports = []leanruntime.Field{
	{Name: "BulkDeploymentId", Flag: "bulk-deployment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bulk_deployments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connector_definition_versions = []leanruntime.Field{
	{Name: "ConnectorDefinitionId", Flag: "connector-definition-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connector_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_core_definition_versions = []leanruntime.Field{
	{Name: "CoreDefinitionId", Flag: "core-definition-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_core_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_device_definition_versions = []leanruntime.Field{
	{Name: "DeviceDefinitionId", Flag: "device-definition-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_device_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_function_definition_versions = []leanruntime.Field{
	{Name: "FunctionDefinitionId", Flag: "function-definition-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_function_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_group_certificate_authorities = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_list_group_versions = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_logger_definition_versions = []leanruntime.Field{
	{Name: "LoggerDefinitionId", Flag: "logger-definition-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_logger_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_definition_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceDefinitionId", Flag: "resource-definition-id", Type: "*string", Required: true},
}

var fields_list_resource_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_subscription_definition_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubscriptionDefinitionId", Flag: "subscription-definition-id", Type: "*string", Required: true},
}

var fields_list_subscription_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reset_deployments = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_start_bulk_deployment = []leanruntime.Field{
	{Name: "AmznClientToken", Flag: "amzn-client-token", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "InputFileUri", Flag: "input-file-uri", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_bulk_deployment = []leanruntime.Field{
	{Name: "BulkDeploymentId", Flag: "bulk-deployment-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_connectivity_info = []leanruntime.Field{
	{Name: "ConnectivityInfo", Flag: "connectivity-info", Type: "[]types.ConnectivityInfo", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_update_connector_definition = []leanruntime.Field{
	{Name: "ConnectorDefinitionId", Flag: "connector-definition-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_core_definition = []leanruntime.Field{
	{Name: "CoreDefinitionId", Flag: "core-definition-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_device_definition = []leanruntime.Field{
	{Name: "DeviceDefinitionId", Flag: "device-definition-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_function_definition = []leanruntime.Field{
	{Name: "FunctionDefinitionId", Flag: "function-definition-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_group_certificate_configuration = []leanruntime.Field{
	{Name: "CertificateExpiryInMilliseconds", Flag: "certificate-expiry-in-milliseconds", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_update_logger_definition = []leanruntime.Field{
	{Name: "LoggerDefinitionId", Flag: "logger-definition-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_resource_definition = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResourceDefinitionId", Flag: "resource-definition-id", Type: "*string", Required: true},
}

var fields_update_subscription_definition = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SubscriptionDefinitionId", Flag: "subscription-definition-id", Type: "*string", Required: true},
}

var fields_update_thing_runtime_configuration = []leanruntime.Field{
	{Name: "TelemetryConfiguration", Flag: "telemetry-configuration", Type: "*types.TelemetryConfigurationUpdate", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-role-to-group": {
			Name:   "associate-role-to-group",
			Fields: fields_associate_role_to_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateRoleToGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_role_to_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateRoleToGroup(ctx, input)
			},
		},
		"associate-service-role-to-account": {
			Name:   "associate-service-role-to-account",
			Fields: fields_associate_service_role_to_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateServiceRoleToAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_service_role_to_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateServiceRoleToAccount(ctx, input)
			},
		},
		"create-connector-definition": {
			Name:   "create-connector-definition",
			Fields: fields_create_connector_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectorDefinition(ctx, input)
			},
		},
		"create-connector-definition-version": {
			Name:   "create-connector-definition-version",
			Fields: fields_create_connector_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectorDefinitionVersion(ctx, input)
			},
		},
		"create-core-definition": {
			Name:   "create-core-definition",
			Fields: fields_create_core_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCoreDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_core_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCoreDefinition(ctx, input)
			},
		},
		"create-core-definition-version": {
			Name:   "create-core-definition-version",
			Fields: fields_create_core_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCoreDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_core_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCoreDefinitionVersion(ctx, input)
			},
		},
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
			},
		},
		"create-device-definition": {
			Name:   "create-device-definition",
			Fields: fields_create_device_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeviceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_device_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeviceDefinition(ctx, input)
			},
		},
		"create-device-definition-version": {
			Name:   "create-device-definition-version",
			Fields: fields_create_device_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeviceDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_device_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeviceDefinitionVersion(ctx, input)
			},
		},
		"create-function-definition": {
			Name:   "create-function-definition",
			Fields: fields_create_function_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFunctionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_function_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFunctionDefinition(ctx, input)
			},
		},
		"create-function-definition-version": {
			Name:   "create-function-definition-version",
			Fields: fields_create_function_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFunctionDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_function_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFunctionDefinitionVersion(ctx, input)
			},
		},
		"create-group": {
			Name:   "create-group",
			Fields: fields_create_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroup(ctx, input)
			},
		},
		"create-group-certificate-authority": {
			Name:   "create-group-certificate-authority",
			Fields: fields_create_group_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroupCertificateAuthority(ctx, input)
			},
		},
		"create-group-version": {
			Name:   "create-group-version",
			Fields: fields_create_group_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroupVersion(ctx, input)
			},
		},
		"create-logger-definition": {
			Name:   "create-logger-definition",
			Fields: fields_create_logger_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoggerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_logger_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoggerDefinition(ctx, input)
			},
		},
		"create-logger-definition-version": {
			Name:   "create-logger-definition-version",
			Fields: fields_create_logger_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoggerDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_logger_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoggerDefinitionVersion(ctx, input)
			},
		},
		"create-resource-definition": {
			Name:   "create-resource-definition",
			Fields: fields_create_resource_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceDefinition(ctx, input)
			},
		},
		"create-resource-definition-version": {
			Name:   "create-resource-definition-version",
			Fields: fields_create_resource_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceDefinitionVersion(ctx, input)
			},
		},
		"create-software-update-job": {
			Name:   "create-software-update-job",
			Fields: fields_create_software_update_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSoftwareUpdateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_software_update_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSoftwareUpdateJob(ctx, input)
			},
		},
		"create-subscription-definition": {
			Name:   "create-subscription-definition",
			Fields: fields_create_subscription_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriptionDefinition(ctx, input)
			},
		},
		"create-subscription-definition-version": {
			Name:   "create-subscription-definition-version",
			Fields: fields_create_subscription_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriptionDefinitionVersion(ctx, input)
			},
		},
		"delete-connector-definition": {
			Name:   "delete-connector-definition",
			Fields: fields_delete_connector_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectorDefinition(ctx, input)
			},
		},
		"delete-core-definition": {
			Name:   "delete-core-definition",
			Fields: fields_delete_core_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoreDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_core_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoreDefinition(ctx, input)
			},
		},
		"delete-device-definition": {
			Name:   "delete-device-definition",
			Fields: fields_delete_device_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeviceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_device_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeviceDefinition(ctx, input)
			},
		},
		"delete-function-definition": {
			Name:   "delete-function-definition",
			Fields: fields_delete_function_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFunctionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_function_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFunctionDefinition(ctx, input)
			},
		},
		"delete-group": {
			Name:   "delete-group",
			Fields: fields_delete_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroup(ctx, input)
			},
		},
		"delete-logger-definition": {
			Name:   "delete-logger-definition",
			Fields: fields_delete_logger_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoggerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_logger_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoggerDefinition(ctx, input)
			},
		},
		"delete-resource-definition": {
			Name:   "delete-resource-definition",
			Fields: fields_delete_resource_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceDefinition(ctx, input)
			},
		},
		"delete-subscription-definition": {
			Name:   "delete-subscription-definition",
			Fields: fields_delete_subscription_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriptionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscription_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriptionDefinition(ctx, input)
			},
		},
		"disassociate-role-from-group": {
			Name:   "disassociate-role-from-group",
			Fields: fields_disassociate_role_from_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRoleFromGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_role_from_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRoleFromGroup(ctx, input)
			},
		},
		"disassociate-service-role-from-account": {
			Name:   "disassociate-service-role-from-account",
			Fields: fields_disassociate_service_role_from_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateServiceRoleFromAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_service_role_from_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateServiceRoleFromAccount(ctx, input)
			},
		},
		"get-associated-role": {
			Name:   "get-associated-role",
			Fields: fields_get_associated_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssociatedRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_associated_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssociatedRole(ctx, input)
			},
		},
		"get-bulk-deployment-status": {
			Name:   "get-bulk-deployment-status",
			Fields: fields_get_bulk_deployment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBulkDeploymentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bulk_deployment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBulkDeploymentStatus(ctx, input)
			},
		},
		"get-connectivity-info": {
			Name:   "get-connectivity-info",
			Fields: fields_get_connectivity_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectivityInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connectivity_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectivityInfo(ctx, input)
			},
		},
		"get-connector-definition": {
			Name:   "get-connector-definition",
			Fields: fields_get_connector_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectorDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connector_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectorDefinition(ctx, input)
			},
		},
		"get-connector-definition-version": {
			Name:   "get-connector-definition-version",
			Fields: fields_get_connector_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectorDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connector_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectorDefinitionVersion(ctx, input)
			},
		},
		"get-core-definition": {
			Name:   "get-core-definition",
			Fields: fields_get_core_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_core_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoreDefinition(ctx, input)
			},
		},
		"get-core-definition-version": {
			Name:   "get-core-definition-version",
			Fields: fields_get_core_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_core_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoreDefinitionVersion(ctx, input)
			},
		},
		"get-deployment-status": {
			Name:   "get-deployment-status",
			Fields: fields_get_deployment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentStatus(ctx, input)
			},
		},
		"get-device-definition": {
			Name:   "get-device-definition",
			Fields: fields_get_device_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceDefinition(ctx, input)
			},
		},
		"get-device-definition-version": {
			Name:   "get-device-definition-version",
			Fields: fields_get_device_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceDefinitionVersion(ctx, input)
			},
		},
		"get-function-definition": {
			Name:   "get-function-definition",
			Fields: fields_get_function_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionDefinition(ctx, input)
			},
		},
		"get-function-definition-version": {
			Name:   "get-function-definition-version",
			Fields: fields_get_function_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFunctionDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_function_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFunctionDefinitionVersion(ctx, input)
			},
		},
		"get-group": {
			Name:   "get-group",
			Fields: fields_get_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroup(ctx, input)
			},
		},
		"get-group-certificate-authority": {
			Name:   "get-group-certificate-authority",
			Fields: fields_get_group_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupCertificateAuthority(ctx, input)
			},
		},
		"get-group-certificate-configuration": {
			Name:   "get-group-certificate-configuration",
			Fields: fields_get_group_certificate_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupCertificateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_certificate_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupCertificateConfiguration(ctx, input)
			},
		},
		"get-group-version": {
			Name:   "get-group-version",
			Fields: fields_get_group_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupVersion(ctx, input)
			},
		},
		"get-logger-definition": {
			Name:   "get-logger-definition",
			Fields: fields_get_logger_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoggerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_logger_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoggerDefinition(ctx, input)
			},
		},
		"get-logger-definition-version": {
			Name:   "get-logger-definition-version",
			Fields: fields_get_logger_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoggerDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_logger_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoggerDefinitionVersion(ctx, input)
			},
		},
		"get-resource-definition": {
			Name:   "get-resource-definition",
			Fields: fields_get_resource_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceDefinition(ctx, input)
			},
		},
		"get-resource-definition-version": {
			Name:   "get-resource-definition-version",
			Fields: fields_get_resource_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceDefinitionVersion(ctx, input)
			},
		},
		"get-service-role-for-account": {
			Name:   "get-service-role-for-account",
			Fields: fields_get_service_role_for_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceRoleForAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_role_for_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceRoleForAccount(ctx, input)
			},
		},
		"get-subscription-definition": {
			Name:   "get-subscription-definition",
			Fields: fields_get_subscription_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionDefinition(ctx, input)
			},
		},
		"get-subscription-definition-version": {
			Name:   "get-subscription-definition-version",
			Fields: fields_get_subscription_definition_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionDefinitionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_definition_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionDefinitionVersion(ctx, input)
			},
		},
		"get-thing-runtime-configuration": {
			Name:   "get-thing-runtime-configuration",
			Fields: fields_get_thing_runtime_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThingRuntimeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_thing_runtime_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThingRuntimeConfiguration(ctx, input)
			},
		},
		"list-bulk-deployment-detailed-reports": {
			Name:   "list-bulk-deployment-detailed-reports",
			Fields: fields_list_bulk_deployment_detailed_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBulkDeploymentDetailedReportsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bulk_deployment_detailed_reports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBulkDeploymentDetailedReports(ctx, input)
			},
		},
		"list-bulk-deployments": {
			Name:   "list-bulk-deployments",
			Fields: fields_list_bulk_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBulkDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bulk_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBulkDeployments(ctx, input)
			},
		},
		"list-connector-definition-versions": {
			Name:   "list-connector-definition-versions",
			Fields: fields_list_connector_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_connector_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConnectorDefinitionVersions(ctx, input)
			},
		},
		"list-connector-definitions": {
			Name:   "list-connector-definitions",
			Fields: fields_list_connector_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_connector_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConnectorDefinitions(ctx, input)
			},
		},
		"list-core-definition-versions": {
			Name:   "list-core-definition-versions",
			Fields: fields_list_core_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_core_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCoreDefinitionVersions(ctx, input)
			},
		},
		"list-core-definitions": {
			Name:   "list-core-definitions",
			Fields: fields_list_core_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_core_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCoreDefinitions(ctx, input)
			},
		},
		"list-deployments": {
			Name:   "list-deployments",
			Fields: fields_list_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDeployments(ctx, input)
			},
		},
		"list-device-definition-versions": {
			Name:   "list-device-definition-versions",
			Fields: fields_list_device_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_device_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDeviceDefinitionVersions(ctx, input)
			},
		},
		"list-device-definitions": {
			Name:   "list-device-definitions",
			Fields: fields_list_device_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_device_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDeviceDefinitions(ctx, input)
			},
		},
		"list-function-definition-versions": {
			Name:   "list-function-definition-versions",
			Fields: fields_list_function_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_function_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFunctionDefinitionVersions(ctx, input)
			},
		},
		"list-function-definitions": {
			Name:   "list-function-definitions",
			Fields: fields_list_function_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFunctionDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_function_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFunctionDefinitions(ctx, input)
			},
		},
		"list-group-certificate-authorities": {
			Name:   "list-group-certificate-authorities",
			Fields: fields_list_group_certificate_authorities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupCertificateAuthoritiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_group_certificate_authorities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGroupCertificateAuthorities(ctx, input)
			},
		},
		"list-group-versions": {
			Name:   "list-group-versions",
			Fields: fields_list_group_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_group_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGroupVersions(ctx, input)
			},
		},
		"list-groups": {
			Name:   "list-groups",
			Fields: fields_list_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGroups(ctx, input)
			},
		},
		"list-logger-definition-versions": {
			Name:   "list-logger-definition-versions",
			Fields: fields_list_logger_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLoggerDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_logger_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLoggerDefinitionVersions(ctx, input)
			},
		},
		"list-logger-definitions": {
			Name:   "list-logger-definitions",
			Fields: fields_list_logger_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLoggerDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_logger_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLoggerDefinitions(ctx, input)
			},
		},
		"list-resource-definition-versions": {
			Name:   "list-resource-definition-versions",
			Fields: fields_list_resource_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resource_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourceDefinitionVersions(ctx, input)
			},
		},
		"list-resource-definitions": {
			Name:   "list-resource-definitions",
			Fields: fields_list_resource_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resource_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourceDefinitions(ctx, input)
			},
		},
		"list-subscription-definition-versions": {
			Name:   "list-subscription-definition-versions",
			Fields: fields_list_subscription_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionDefinitionVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_subscription_definition_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSubscriptionDefinitionVersions(ctx, input)
			},
		},
		"list-subscription-definitions": {
			Name:   "list-subscription-definitions",
			Fields: fields_list_subscription_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_subscription_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSubscriptionDefinitions(ctx, input)
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
		"reset-deployments": {
			Name:   "reset-deployments",
			Fields: fields_reset_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetDeployments(ctx, input)
			},
		},
		"start-bulk-deployment": {
			Name:   "start-bulk-deployment",
			Fields: fields_start_bulk_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBulkDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_bulk_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBulkDeployment(ctx, input)
			},
		},
		"stop-bulk-deployment": {
			Name:   "stop-bulk-deployment",
			Fields: fields_stop_bulk_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBulkDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_bulk_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBulkDeployment(ctx, input)
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
		"update-connectivity-info": {
			Name:   "update-connectivity-info",
			Fields: fields_update_connectivity_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectivityInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connectivity_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectivityInfo(ctx, input)
			},
		},
		"update-connector-definition": {
			Name:   "update-connector-definition",
			Fields: fields_update_connector_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectorDefinition(ctx, input)
			},
		},
		"update-core-definition": {
			Name:   "update-core-definition",
			Fields: fields_update_core_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCoreDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_core_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCoreDefinition(ctx, input)
			},
		},
		"update-device-definition": {
			Name:   "update-device-definition",
			Fields: fields_update_device_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeviceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeviceDefinition(ctx, input)
			},
		},
		"update-function-definition": {
			Name:   "update-function-definition",
			Fields: fields_update_function_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFunctionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_function_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFunctionDefinition(ctx, input)
			},
		},
		"update-group": {
			Name:   "update-group",
			Fields: fields_update_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroup(ctx, input)
			},
		},
		"update-group-certificate-configuration": {
			Name:   "update-group-certificate-configuration",
			Fields: fields_update_group_certificate_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupCertificateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group_certificate_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroupCertificateConfiguration(ctx, input)
			},
		},
		"update-logger-definition": {
			Name:   "update-logger-definition",
			Fields: fields_update_logger_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLoggerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_logger_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLoggerDefinition(ctx, input)
			},
		},
		"update-resource-definition": {
			Name:   "update-resource-definition",
			Fields: fields_update_resource_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceDefinition(ctx, input)
			},
		},
		"update-subscription-definition": {
			Name:   "update-subscription-definition",
			Fields: fields_update_subscription_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscription_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriptionDefinition(ctx, input)
			},
		},
		"update-thing-runtime-configuration": {
			Name:   "update-thing-runtime-configuration",
			Fields: fields_update_thing_runtime_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThingRuntimeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thing_runtime_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThingRuntimeConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("greengrass", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
