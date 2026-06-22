package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

var fields_abort_environment_update = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
}

var fields_apply_environment_managed_action = []leanruntime.Field{
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
}

var fields_associate_environment_operations_role = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "OperationsRole", Flag: "operations-role", Type: "*string", Required: true},
}

var fields_check_dns_availability = []leanruntime.Field{
	{Name: "CNAMEPrefix", Flag: "cname-prefix", Type: "*string", Required: true},
}

var fields_compose_environments = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "VersionLabels", Flag: "version-labels", Type: "[]string", Required: false},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ResourceLifecycleConfig", Flag: "resource-lifecycle-config", Type: "*types.ApplicationResourceLifecycleConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_application_version = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "AutoCreateApplication", Flag: "auto-create-application", Type: "*bool", Required: false},
	{Name: "BuildConfiguration", Flag: "build-configuration", Type: "*types.BuildConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Process", Flag: "process", Type: "*bool", Required: false},
	{Name: "SourceBuildInformation", Flag: "source-build-information", Type: "*types.SourceBuildInformation", Required: false},
	{Name: "SourceBundle", Flag: "source-bundle", Type: "*types.S3Location", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: true},
}

var fields_create_configuration_template = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "OptionSettings", Flag: "option-settings", Type: "[]types.ConfigurationOptionSetting", Required: false},
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
	{Name: "SolutionStackName", Flag: "solution-stack-name", Type: "*string", Required: false},
	{Name: "SourceConfiguration", Flag: "source-configuration", Type: "*types.SourceConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CNAMEPrefix", Flag: "cname-prefix", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "OperationsRole", Flag: "operations-role", Type: "*string", Required: false},
	{Name: "OptionSettings", Flag: "option-settings", Type: "[]types.ConfigurationOptionSetting", Required: false},
	{Name: "OptionsToRemove", Flag: "options-to-remove", Type: "[]types.OptionSpecification", Required: false},
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
	{Name: "SolutionStackName", Flag: "solution-stack-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "*types.EnvironmentTier", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: false},
}

var fields_create_platform_version = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "OptionSettings", Flag: "option-settings", Type: "[]types.ConfigurationOptionSetting", Required: false},
	{Name: "PlatformDefinitionBundle", Flag: "platform-definition-bundle", Type: "*types.S3Location", Required: true},
	{Name: "PlatformName", Flag: "platform-name", Type: "*string", Required: true},
	{Name: "PlatformVersion", Flag: "platform-version", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_storage_location = []leanruntime.Field{}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "TerminateEnvByForce", Flag: "terminate-env-by-force", Type: "*bool", Required: false},
}

var fields_delete_application_version = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "DeleteSourceBundle", Flag: "delete-source-bundle", Type: "*bool", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: true},
}

var fields_delete_configuration_template = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_environment_configuration = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_delete_platform_version = []leanruntime.Field{
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
}

var fields_describe_account_attributes = []leanruntime.Field{}

var fields_describe_application_versions = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VersionLabels", Flag: "version-labels", Type: "[]string", Required: false},
}

var fields_describe_applications = []leanruntime.Field{
	{Name: "ApplicationNames", Flag: "application-names", Type: "[]string", Required: false},
}

var fields_describe_configuration_options = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "Options", Flag: "options", Type: "[]types.OptionSpecification", Required: false},
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
	{Name: "SolutionStackName", Flag: "solution-stack-name", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

var fields_describe_configuration_settings = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

var fields_describe_environment_health = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]types.EnvironmentHealthAttribute", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
}

var fields_describe_environment_managed_action_history = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_environment_managed_actions = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ActionStatus", Required: false},
}

var fields_describe_environment_resources = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
}

var fields_describe_environments = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "EnvironmentIds", Flag: "environment-ids", Type: "[]string", Required: false},
	{Name: "EnvironmentNames", Flag: "environment-names", Type: "[]string", Required: false},
	{Name: "IncludeDeleted", Flag: "include-deleted", Type: "*bool", Required: false},
	{Name: "IncludedDeletedBackTo", Flag: "included-deleted-back-to", Type: "*time.Time", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: false},
}

var fields_describe_events = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Severity", Flag: "severity", Type: "types.EventSeverity", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: false},
}

var fields_describe_instances_health = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]types.InstancesHealthAttribute", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_platform_version = []leanruntime.Field{
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
}

var fields_disassociate_environment_operations_role = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_list_available_solution_stacks = []leanruntime.Field{}

var fields_list_platform_branches = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SearchFilter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_platform_versions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PlatformFilter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_rebuild_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
}

var fields_request_environment_info = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "InfoType", Flag: "info-type", Type: "types.EnvironmentInfoType", Required: true},
}

var fields_restart_app_server = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
}

var fields_retrieve_environment_info = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "InfoType", Flag: "info-type", Type: "types.EnvironmentInfoType", Required: true},
}

var fields_swap_environment_cnames = []leanruntime.Field{
	{Name: "DestinationEnvironmentId", Flag: "destination-environment-id", Type: "*string", Required: false},
	{Name: "DestinationEnvironmentName", Flag: "destination-environment-name", Type: "*string", Required: false},
	{Name: "SourceEnvironmentId", Flag: "source-environment-id", Type: "*string", Required: false},
	{Name: "SourceEnvironmentName", Flag: "source-environment-name", Type: "*string", Required: false},
}

var fields_terminate_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "ForceTerminate", Flag: "force-terminate", Type: "*bool", Required: false},
	{Name: "TerminateResources", Flag: "terminate-resources", Type: "*bool", Required: false},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_application_resource_lifecycle = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "ResourceLifecycleConfig", Flag: "resource-lifecycle-config", Type: "*types.ApplicationResourceLifecycleConfig", Required: true},
}

var fields_update_application_version = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: true},
}

var fields_update_configuration_template = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "OptionSettings", Flag: "option-settings", Type: "[]types.ConfigurationOptionSetting", Required: false},
	{Name: "OptionsToRemove", Flag: "options-to-remove", Type: "[]types.OptionSpecification", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "OptionSettings", Flag: "option-settings", Type: "[]types.ConfigurationOptionSetting", Required: false},
	{Name: "OptionsToRemove", Flag: "options-to-remove", Type: "[]types.OptionSpecification", Required: false},
	{Name: "PlatformArn", Flag: "platform-arn", Type: "*string", Required: false},
	{Name: "SolutionStackName", Flag: "solution-stack-name", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "*types.EnvironmentTier", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: false},
}

var fields_update_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagsToAdd", Flag: "tags-to-add", Type: "[]types.Tag", Required: false},
	{Name: "TagsToRemove", Flag: "tags-to-remove", Type: "[]string", Required: false},
}

var fields_validate_configuration_settings = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "OptionSettings", Flag: "option-settings", Type: "[]types.ConfigurationOptionSetting", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"abort-environment-update": {
			Name:   "abort-environment-update",
			Fields: fields_abort_environment_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AbortEnvironmentUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_abort_environment_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AbortEnvironmentUpdate(ctx, input)
			},
		},
		"apply-environment-managed-action": {
			Name:   "apply-environment-managed-action",
			Fields: fields_apply_environment_managed_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplyEnvironmentManagedActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_environment_managed_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplyEnvironmentManagedAction(ctx, input)
			},
		},
		"associate-environment-operations-role": {
			Name:   "associate-environment-operations-role",
			Fields: fields_associate_environment_operations_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEnvironmentOperationsRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_environment_operations_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEnvironmentOperationsRole(ctx, input)
			},
		},
		"check-dns-availability": {
			Name:   "check-dns-availability",
			Fields: fields_check_dns_availability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckDNSAvailabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_dns_availability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckDNSAvailability(ctx, input)
			},
		},
		"compose-environments": {
			Name:   "compose-environments",
			Fields: fields_compose_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ComposeEnvironmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_compose_environments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ComposeEnvironments(ctx, input)
			},
		},
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-application-version": {
			Name:   "create-application-version",
			Fields: fields_create_application_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplicationVersion(ctx, input)
			},
		},
		"create-configuration-template": {
			Name:   "create-configuration-template",
			Fields: fields_create_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationTemplate(ctx, input)
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
		"create-platform-version": {
			Name:   "create-platform-version",
			Fields: fields_create_platform_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlatformVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_platform_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlatformVersion(ctx, input)
			},
		},
		"create-storage-location": {
			Name:   "create-storage-location",
			Fields: fields_create_storage_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStorageLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_storage_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStorageLocation(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-application-version": {
			Name:   "delete-application-version",
			Fields: fields_delete_application_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationVersion(ctx, input)
			},
		},
		"delete-configuration-template": {
			Name:   "delete-configuration-template",
			Fields: fields_delete_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationTemplate(ctx, input)
			},
		},
		"delete-environment-configuration": {
			Name:   "delete-environment-configuration",
			Fields: fields_delete_environment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentConfiguration(ctx, input)
			},
		},
		"delete-platform-version": {
			Name:   "delete-platform-version",
			Fields: fields_delete_platform_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlatformVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_platform_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlatformVersion(ctx, input)
			},
		},
		"describe-account-attributes": {
			Name:   "describe-account-attributes",
			Fields: fields_describe_account_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountAttributes(ctx, input)
			},
		},
		"describe-application-versions": {
			Name:   "describe-application-versions",
			Fields: fields_describe_application_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationVersions(ctx, input)
			},
		},
		"describe-applications": {
			Name:   "describe-applications",
			Fields: fields_describe_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplications(ctx, input)
			},
		},
		"describe-configuration-options": {
			Name:   "describe-configuration-options",
			Fields: fields_describe_configuration_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationOptions(ctx, input)
			},
		},
		"describe-configuration-settings": {
			Name:   "describe-configuration-settings",
			Fields: fields_describe_configuration_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationSettings(ctx, input)
			},
		},
		"describe-environment-health": {
			Name:   "describe-environment-health",
			Fields: fields_describe_environment_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_environment_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEnvironmentHealth(ctx, input)
			},
		},
		"describe-environment-managed-action-history": {
			Name:   "describe-environment-managed-action-history",
			Fields: fields_describe_environment_managed_action_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentManagedActionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_environment_managed_action_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEnvironmentManagedActionHistory(ctx, input)
				}
				var results []*svc.DescribeEnvironmentManagedActionHistoryOutput
				p := svc.NewDescribeEnvironmentManagedActionHistoryPaginator(client, input)
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
		"describe-environment-managed-actions": {
			Name:   "describe-environment-managed-actions",
			Fields: fields_describe_environment_managed_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentManagedActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_environment_managed_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEnvironmentManagedActions(ctx, input)
			},
		},
		"describe-environment-resources": {
			Name:   "describe-environment-resources",
			Fields: fields_describe_environment_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_environment_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEnvironmentResources(ctx, input)
			},
		},
		"describe-environments": {
			Name:   "describe-environments",
			Fields: fields_describe_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_environments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEnvironments(ctx, input)
			},
		},
		"describe-events": {
			Name:   "describe-events",
			Fields: fields_describe_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEvents(ctx, input)
				}
				var results []*svc.DescribeEventsOutput
				p := svc.NewDescribeEventsPaginator(client, input)
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
		"describe-instances-health": {
			Name:   "describe-instances-health",
			Fields: fields_describe_instances_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancesHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instances_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstancesHealth(ctx, input)
			},
		},
		"describe-platform-version": {
			Name:   "describe-platform-version",
			Fields: fields_describe_platform_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePlatformVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_platform_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePlatformVersion(ctx, input)
			},
		},
		"disassociate-environment-operations-role": {
			Name:   "disassociate-environment-operations-role",
			Fields: fields_disassociate_environment_operations_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateEnvironmentOperationsRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_environment_operations_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateEnvironmentOperationsRole(ctx, input)
			},
		},
		"list-available-solution-stacks": {
			Name:   "list-available-solution-stacks",
			Fields: fields_list_available_solution_stacks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableSolutionStacksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_available_solution_stacks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAvailableSolutionStacks(ctx, input)
			},
		},
		"list-platform-branches": {
			Name:   "list-platform-branches",
			Fields: fields_list_platform_branches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlatformBranchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_platform_branches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlatformBranches(ctx, input)
				}
				var results []*svc.ListPlatformBranchesOutput
				p := svc.NewListPlatformBranchesPaginator(client, input)
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
		"list-platform-versions": {
			Name:   "list-platform-versions",
			Fields: fields_list_platform_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlatformVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_platform_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlatformVersions(ctx, input)
				}
				var results []*svc.ListPlatformVersionsOutput
				p := svc.NewListPlatformVersionsPaginator(client, input)
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
		"rebuild-environment": {
			Name:   "rebuild-environment",
			Fields: fields_rebuild_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebuildEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rebuild_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebuildEnvironment(ctx, input)
			},
		},
		"request-environment-info": {
			Name:   "request-environment-info",
			Fields: fields_request_environment_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestEnvironmentInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_environment_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestEnvironmentInfo(ctx, input)
			},
		},
		"restart-app-server": {
			Name:   "restart-app-server",
			Fields: fields_restart_app_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestartAppServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restart_app_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestartAppServer(ctx, input)
			},
		},
		"retrieve-environment-info": {
			Name:   "retrieve-environment-info",
			Fields: fields_retrieve_environment_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveEnvironmentInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve_environment_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetrieveEnvironmentInfo(ctx, input)
			},
		},
		"swap-environment-cnames": {
			Name:   "swap-environment-cnames",
			Fields: fields_swap_environment_cnames,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SwapEnvironmentCNAMEsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_swap_environment_cnames, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SwapEnvironmentCNAMEs(ctx, input)
			},
		},
		"terminate-environment": {
			Name:   "terminate-environment",
			Fields: fields_terminate_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateEnvironment(ctx, input)
			},
		},
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
		"update-application-resource-lifecycle": {
			Name:   "update-application-resource-lifecycle",
			Fields: fields_update_application_resource_lifecycle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationResourceLifecycleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_resource_lifecycle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationResourceLifecycle(ctx, input)
			},
		},
		"update-application-version": {
			Name:   "update-application-version",
			Fields: fields_update_application_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationVersion(ctx, input)
			},
		},
		"update-configuration-template": {
			Name:   "update-configuration-template",
			Fields: fields_update_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationTemplate(ctx, input)
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
		"update-tags-for-resource": {
			Name:   "update-tags-for-resource",
			Fields: fields_update_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTagsForResource(ctx, input)
			},
		},
		"validate-configuration-settings": {
			Name:   "validate-configuration-settings",
			Fields: fields_validate_configuration_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateConfigurationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_configuration_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateConfigurationSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("elasticbeanstalk", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
