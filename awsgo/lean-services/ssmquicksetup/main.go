package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssmquicksetup"
)

var fields_create_configuration_manager = []leanruntime.Field{
	{Name: "ConfigurationDefinitions", Flag: "configuration-definitions", Type: "[]types.ConfigurationDefinitionInput", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_configuration_manager = []leanruntime.Field{
	{Name: "ManagerArn", Flag: "manager-arn", Type: "*string", Required: true},
}

var fields_get_configuration = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
}

var fields_get_configuration_manager = []leanruntime.Field{
	{Name: "ManagerArn", Flag: "manager-arn", Type: "*string", Required: true},
}

var fields_get_service_settings = []leanruntime.Field{}

var fields_list_configuration_managers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
}

var fields_list_configurations = []leanruntime.Field{
	{Name: "ConfigurationDefinitionId", Flag: "configuration-definition-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ManagerArn", Flag: "manager-arn", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
}

var fields_list_quick_setup_types = []leanruntime.Field{}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_configuration_definition = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LocalDeploymentAdministrationRoleArn", Flag: "local-deployment-administration-role-arn", Type: "*string", Required: false},
	{Name: "LocalDeploymentExecutionRoleName", Flag: "local-deployment-execution-role-name", Type: "*string", Required: false},
	{Name: "ManagerArn", Flag: "manager-arn", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "TypeVersion", Flag: "type-version", Type: "*string", Required: false},
}

var fields_update_configuration_manager = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ManagerArn", Flag: "manager-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_service_settings = []leanruntime.Field{
	{Name: "ExplorerEnablingRoleArn", Flag: "explorer-enabling-role-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-configuration-manager": {
			Name:   "create-configuration-manager",
			Fields: fields_create_configuration_manager,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationManagerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_manager, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationManager(ctx, input)
			},
		},
		"delete-configuration-manager": {
			Name:   "delete-configuration-manager",
			Fields: fields_delete_configuration_manager,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationManagerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_manager, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationManager(ctx, input)
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
		"get-configuration-manager": {
			Name:   "get-configuration-manager",
			Fields: fields_get_configuration_manager,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationManagerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_manager, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationManager(ctx, input)
			},
		},
		"get-service-settings": {
			Name:   "get-service-settings",
			Fields: fields_get_service_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceSettings(ctx, input)
			},
		},
		"list-configuration-managers": {
			Name:   "list-configuration-managers",
			Fields: fields_list_configuration_managers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationManagersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_managers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationManagers(ctx, input)
				}
				var results []*svc.ListConfigurationManagersOutput
				p := svc.NewListConfigurationManagersPaginator(client, input)
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
		"list-configurations": {
			Name:   "list-configurations",
			Fields: fields_list_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurations(ctx, input)
				}
				var results []*svc.ListConfigurationsOutput
				p := svc.NewListConfigurationsPaginator(client, input)
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
		"list-quick-setup-types": {
			Name:   "list-quick-setup-types",
			Fields: fields_list_quick_setup_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQuickSetupTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_quick_setup_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListQuickSetupTypes(ctx, input)
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
		"update-configuration-definition": {
			Name:   "update-configuration-definition",
			Fields: fields_update_configuration_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationDefinition(ctx, input)
			},
		},
		"update-configuration-manager": {
			Name:   "update-configuration-manager",
			Fields: fields_update_configuration_manager,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationManagerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_manager, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationManager(ctx, input)
			},
		},
		"update-service-settings": {
			Name:   "update-service-settings",
			Fields: fields_update_service_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssmquicksetup", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
