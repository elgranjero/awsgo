package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appconfig"
)

var fields_create_application = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configuration_profile = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "LocationUri", Flag: "location-uri", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RetrievalRoleArn", Flag: "retrieval-role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
	{Name: "Validators", Flag: "validators", Type: "[]types.Validator", Required: false},
}

var fields_create_deployment_strategy = []leanruntime.Field{
	{Name: "DeploymentDurationInMinutes", Flag: "deployment-duration-in-minutes", Type: "*int32", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FinalBakeTimeInMinutes", Flag: "final-bake-time-in-minutes", Type: "int32", Required: false},
	{Name: "GrowthFactor", Flag: "growth-factor", Type: "*float32", Required: true},
	{Name: "GrowthType", Flag: "growth-type", Type: "types.GrowthType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReplicateTo", Flag: "replicate-to", Type: "types.ReplicateTo", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Monitors", Flag: "monitors", Type: "[]types.Monitor", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_extension = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "map[string][]types.Action", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LatestVersionNumber", Flag: "latest-version-number", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.Parameter", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_extension_association = []leanruntime.Field{
	{Name: "ExtensionIdentifier", Flag: "extension-identifier", Type: "*string", Required: true},
	{Name: "ExtensionVersionNumber", Flag: "extension-version-number", Type: "*int32", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_hosted_configuration_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "[]byte", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LatestVersionNumber", Flag: "latest-version-number", Type: "*int32", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_configuration_profile = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "DeletionProtectionCheck", Flag: "deletion-protection-check", Type: "types.DeletionProtectionCheck", Required: false},
}

var fields_delete_deployment_strategy = []leanruntime.Field{
	{Name: "DeploymentStrategyId", Flag: "deployment-strategy-id", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DeletionProtectionCheck", Flag: "deletion-protection-check", Type: "types.DeletionProtectionCheck", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_extension = []leanruntime.Field{
	{Name: "ExtensionIdentifier", Flag: "extension-identifier", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: false},
}

var fields_delete_extension_association = []leanruntime.Field{
	{Name: "ExtensionAssociationId", Flag: "extension-association-id", Type: "*string", Required: true},
}

var fields_delete_hosted_configuration_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: true},
}

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_configuration = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "ClientConfigurationVersion", Flag: "client-configuration-version", Type: "*string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: true},
	{Name: "Environment", Flag: "environment", Type: "*string", Required: true},
}

var fields_get_configuration_profile = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DeploymentNumber", Flag: "deployment-number", Type: "*int32", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_deployment_strategy = []leanruntime.Field{
	{Name: "DeploymentStrategyId", Flag: "deployment-strategy-id", Type: "*string", Required: true},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_extension = []leanruntime.Field{
	{Name: "ExtensionIdentifier", Flag: "extension-identifier", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: false},
}

var fields_get_extension_association = []leanruntime.Field{
	{Name: "ExtensionAssociationId", Flag: "extension-association-id", Type: "*string", Required: true},
}

var fields_get_hosted_configuration_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_profiles = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_list_deployment_strategies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_extension_associations = []leanruntime.Field{
	{Name: "ExtensionIdentifier", Flag: "extension-identifier", Type: "*string", Required: false},
	{Name: "ExtensionVersionNumber", Flag: "extension-version-number", Type: "*int32", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
}

var fields_list_extensions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hosted_configuration_versions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VersionLabel", Flag: "version-label", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_deployment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "ConfigurationVersion", Flag: "configuration-version", Type: "*string", Required: true},
	{Name: "DeploymentStrategyId", Flag: "deployment-strategy-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DynamicExtensionParameters", Flag: "dynamic-extension-parameters", Type: "map[string]string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_deployment = []leanruntime.Field{
	{Name: "AllowRevert", Flag: "allow-revert", Type: "*bool", Required: false},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DeploymentNumber", Flag: "deployment-number", Type: "*int32", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*types.DeletionProtectionSettings", Required: false},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_configuration_profile = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RetrievalRoleArn", Flag: "retrieval-role-arn", Type: "*string", Required: false},
	{Name: "Validators", Flag: "validators", Type: "[]types.Validator", Required: false},
}

var fields_update_deployment_strategy = []leanruntime.Field{
	{Name: "DeploymentDurationInMinutes", Flag: "deployment-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "DeploymentStrategyId", Flag: "deployment-strategy-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FinalBakeTimeInMinutes", Flag: "final-bake-time-in-minutes", Type: "*int32", Required: false},
	{Name: "GrowthFactor", Flag: "growth-factor", Type: "*float32", Required: false},
	{Name: "GrowthType", Flag: "growth-type", Type: "types.GrowthType", Required: false},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Monitors", Flag: "monitors", Type: "[]types.Monitor", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_extension = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "map[string][]types.Action", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExtensionIdentifier", Flag: "extension-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.Parameter", Required: false},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: false},
}

var fields_update_extension_association = []leanruntime.Field{
	{Name: "ExtensionAssociationId", Flag: "extension-association-id", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
}

var fields_validate_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationProfileId", Flag: "configuration-profile-id", Type: "*string", Required: true},
	{Name: "ConfigurationVersion", Flag: "configuration-version", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"create-configuration-profile": {
			Name:   "create-configuration-profile",
			Fields: fields_create_configuration_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationProfile(ctx, input)
			},
		},
		"create-deployment-strategy": {
			Name:   "create-deployment-strategy",
			Fields: fields_create_deployment_strategy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentStrategyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment_strategy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeploymentStrategy(ctx, input)
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
		"create-extension": {
			Name:   "create-extension",
			Fields: fields_create_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExtension(ctx, input)
			},
		},
		"create-extension-association": {
			Name:   "create-extension-association",
			Fields: fields_create_extension_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExtensionAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_extension_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExtensionAssociation(ctx, input)
			},
		},
		"create-hosted-configuration-version": {
			Name:   "create-hosted-configuration-version",
			Fields: fields_create_hosted_configuration_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHostedConfigurationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hosted_configuration_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHostedConfigurationVersion(ctx, input)
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
		"delete-configuration-profile": {
			Name:   "delete-configuration-profile",
			Fields: fields_delete_configuration_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationProfile(ctx, input)
			},
		},
		"delete-deployment-strategy": {
			Name:   "delete-deployment-strategy",
			Fields: fields_delete_deployment_strategy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentStrategyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment_strategy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeploymentStrategy(ctx, input)
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
		"delete-extension": {
			Name:   "delete-extension",
			Fields: fields_delete_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExtension(ctx, input)
			},
		},
		"delete-extension-association": {
			Name:   "delete-extension-association",
			Fields: fields_delete_extension_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExtensionAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_extension_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExtensionAssociation(ctx, input)
			},
		},
		"delete-hosted-configuration-version": {
			Name:   "delete-hosted-configuration-version",
			Fields: fields_delete_hosted_configuration_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHostedConfigurationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hosted_configuration_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHostedConfigurationVersion(ctx, input)
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
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-configuration": {
			Name:   "get-configuration",
			Fields: fields_get_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguration(ctx, input)
			},
		},
		"get-configuration-profile": {
			Name:   "get-configuration-profile",
			Fields: fields_get_configuration_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationProfile(ctx, input)
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
		"get-deployment-strategy": {
			Name:   "get-deployment-strategy",
			Fields: fields_get_deployment_strategy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentStrategyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_strategy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentStrategy(ctx, input)
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
		"get-extension": {
			Name:   "get-extension",
			Fields: fields_get_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExtension(ctx, input)
			},
		},
		"get-extension-association": {
			Name:   "get-extension-association",
			Fields: fields_get_extension_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExtensionAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_extension_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExtensionAssociation(ctx, input)
			},
		},
		"get-hosted-configuration-version": {
			Name:   "get-hosted-configuration-version",
			Fields: fields_get_hosted_configuration_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostedConfigurationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hosted_configuration_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHostedConfigurationVersion(ctx, input)
			},
		},
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-configuration-profiles": {
			Name:   "list-configuration-profiles",
			Fields: fields_list_configuration_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationProfiles(ctx, input)
				}
				var results []*svc.ListConfigurationProfilesOutput
				p := svc.NewListConfigurationProfilesPaginator(client, input)
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
		"list-deployment-strategies": {
			Name:   "list-deployment-strategies",
			Fields: fields_list_deployment_strategies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentStrategiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployment_strategies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeploymentStrategies(ctx, input)
				}
				var results []*svc.ListDeploymentStrategiesOutput
				p := svc.NewListDeploymentStrategiesPaginator(client, input)
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
		"list-extension-associations": {
			Name:   "list-extension-associations",
			Fields: fields_list_extension_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExtensionAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_extension_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExtensionAssociations(ctx, input)
				}
				var results []*svc.ListExtensionAssociationsOutput
				p := svc.NewListExtensionAssociationsPaginator(client, input)
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
		"list-extensions": {
			Name:   "list-extensions",
			Fields: fields_list_extensions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExtensionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_extensions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExtensions(ctx, input)
				}
				var results []*svc.ListExtensionsOutput
				p := svc.NewListExtensionsPaginator(client, input)
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
		"list-hosted-configuration-versions": {
			Name:   "list-hosted-configuration-versions",
			Fields: fields_list_hosted_configuration_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostedConfigurationVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hosted_configuration_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHostedConfigurationVersions(ctx, input)
				}
				var results []*svc.ListHostedConfigurationVersionsOutput
				p := svc.NewListHostedConfigurationVersionsPaginator(client, input)
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
		"start-deployment": {
			Name:   "start-deployment",
			Fields: fields_start_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeployment(ctx, input)
			},
		},
		"stop-deployment": {
			Name:   "stop-deployment",
			Fields: fields_stop_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDeployment(ctx, input)
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
		"update-configuration-profile": {
			Name:   "update-configuration-profile",
			Fields: fields_update_configuration_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationProfile(ctx, input)
			},
		},
		"update-deployment-strategy": {
			Name:   "update-deployment-strategy",
			Fields: fields_update_deployment_strategy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeploymentStrategyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_deployment_strategy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeploymentStrategy(ctx, input)
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
		"update-extension": {
			Name:   "update-extension",
			Fields: fields_update_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExtension(ctx, input)
			},
		},
		"update-extension-association": {
			Name:   "update-extension-association",
			Fields: fields_update_extension_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExtensionAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_extension_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExtensionAssociation(ctx, input)
			},
		},
		"validate-configuration": {
			Name:   "validate-configuration",
			Fields: fields_validate_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appconfig", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
