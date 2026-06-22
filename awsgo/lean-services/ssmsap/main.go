package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssmsap"
)

var fields_delete_resource_permission = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.PermissionActionType", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceResourceArn", Flag: "source-resource-arn", Type: "*string", Required: false},
}

var fields_deregister_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "AppRegistryArn", Flag: "app-registry-arn", Type: "*string", Required: false},
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: false},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
}

var fields_get_component = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ComponentId", Flag: "component-id", Type: "*string", Required: true},
}

var fields_get_configuration_check_operation = []leanruntime.Field{
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_get_database = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
	{Name: "ComponentId", Flag: "component-id", Type: "*string", Required: false},
	{Name: "DatabaseArn", Flag: "database-arn", Type: "*string", Required: false},
	{Name: "DatabaseId", Flag: "database-id", Type: "*string", Required: false},
}

var fields_get_operation = []leanruntime.Field{
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_get_resource_permission = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.PermissionActionType", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_components = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_check_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_check_operations = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ListMode", Flag: "list-mode", Type: "types.ConfigurationCheckOperationListingMode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_databases = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
	{Name: "ComponentId", Flag: "component-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_operation_events = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_list_operations = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sub_check_results = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_list_sub_check_rule_results = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubCheckResultId", Flag: "sub-check-result-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_resource_permission = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.PermissionActionType", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceResourceArn", Flag: "source-resource-arn", Type: "*string", Required: true},
}

var fields_register_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ApplicationType", Flag: "application-type", Type: "types.ApplicationType", Required: true},
	{Name: "ComponentsInfo", Flag: "components-info", Type: "[]types.ComponentInfo", Required: false},
	{Name: "Credentials", Flag: "credentials", Type: "[]types.ApplicationCredential", Required: false},
	{Name: "DatabaseArn", Flag: "database-arn", Type: "*string", Required: false},
	{Name: "Instances", Flag: "instances", Type: "[]string", Required: true},
	{Name: "SapInstanceNumber", Flag: "sap-instance-number", Type: "*string", Required: false},
	{Name: "Sid", Flag: "sid", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_start_application_refresh = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_start_configuration_checks = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConfigurationCheckIds", Flag: "configuration-check-ids", Type: "[]types.ConfigurationCheckType", Required: false},
}

var fields_stop_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "IncludeEc2InstanceShutdown", Flag: "include-ec2-instance-shutdown", Type: "*bool", Required: false},
	{Name: "StopConnectedEntity", Flag: "stop-connected-entity", Type: "types.ConnectedEntityType", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application_settings = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Backint", Flag: "backint", Type: "*types.BackintConfig", Required: false},
	{Name: "CredentialsToAddOrUpdate", Flag: "credentials-to-add-or-update", Type: "[]types.ApplicationCredential", Required: false},
	{Name: "CredentialsToRemove", Flag: "credentials-to-remove", Type: "[]types.ApplicationCredential", Required: false},
	{Name: "DatabaseArn", Flag: "database-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-resource-permission": {
			Name:   "delete-resource-permission",
			Fields: fields_delete_resource_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePermission(ctx, input)
			},
		},
		"deregister-application": {
			Name:   "deregister-application",
			Fields: fields_deregister_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterApplication(ctx, input)
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
		"get-configuration-check-operation": {
			Name:   "get-configuration-check-operation",
			Fields: fields_get_configuration_check_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationCheckOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_check_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationCheckOperation(ctx, input)
			},
		},
		"get-database": {
			Name:   "get-database",
			Fields: fields_get_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDatabase(ctx, input)
			},
		},
		"get-operation": {
			Name:   "get-operation",
			Fields: fields_get_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOperation(ctx, input)
			},
		},
		"get-resource-permission": {
			Name:   "get-resource-permission",
			Fields: fields_get_resource_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePermission(ctx, input)
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
		"list-configuration-check-definitions": {
			Name:   "list-configuration-check-definitions",
			Fields: fields_list_configuration_check_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationCheckDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_check_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationCheckDefinitions(ctx, input)
				}
				var results []*svc.ListConfigurationCheckDefinitionsOutput
				p := svc.NewListConfigurationCheckDefinitionsPaginator(client, input)
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
		"list-configuration-check-operations": {
			Name:   "list-configuration-check-operations",
			Fields: fields_list_configuration_check_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationCheckOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_check_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationCheckOperations(ctx, input)
				}
				var results []*svc.ListConfigurationCheckOperationsOutput
				p := svc.NewListConfigurationCheckOperationsPaginator(client, input)
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
		"list-databases": {
			Name:   "list-databases",
			Fields: fields_list_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatabases(ctx, input)
				}
				var results []*svc.ListDatabasesOutput
				p := svc.NewListDatabasesPaginator(client, input)
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
		"list-operation-events": {
			Name:   "list-operation-events",
			Fields: fields_list_operation_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOperationEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_operation_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOperationEvents(ctx, input)
				}
				var results []*svc.ListOperationEventsOutput
				p := svc.NewListOperationEventsPaginator(client, input)
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
		"list-operations": {
			Name:   "list-operations",
			Fields: fields_list_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOperations(ctx, input)
				}
				var results []*svc.ListOperationsOutput
				p := svc.NewListOperationsPaginator(client, input)
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
		"list-sub-check-results": {
			Name:   "list-sub-check-results",
			Fields: fields_list_sub_check_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubCheckResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sub_check_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubCheckResults(ctx, input)
				}
				var results []*svc.ListSubCheckResultsOutput
				p := svc.NewListSubCheckResultsPaginator(client, input)
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
		"list-sub-check-rule-results": {
			Name:   "list-sub-check-rule-results",
			Fields: fields_list_sub_check_rule_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubCheckRuleResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sub_check_rule_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubCheckRuleResults(ctx, input)
				}
				var results []*svc.ListSubCheckRuleResultsOutput
				p := svc.NewListSubCheckRuleResultsPaginator(client, input)
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
		"put-resource-permission": {
			Name:   "put-resource-permission",
			Fields: fields_put_resource_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePermission(ctx, input)
			},
		},
		"register-application": {
			Name:   "register-application",
			Fields: fields_register_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterApplication(ctx, input)
			},
		},
		"start-application": {
			Name:   "start-application",
			Fields: fields_start_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartApplication(ctx, input)
			},
		},
		"start-application-refresh": {
			Name:   "start-application-refresh",
			Fields: fields_start_application_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartApplicationRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_application_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartApplicationRefresh(ctx, input)
			},
		},
		"start-configuration-checks": {
			Name:   "start-configuration-checks",
			Fields: fields_start_configuration_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConfigurationChecksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_configuration_checks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConfigurationChecks(ctx, input)
			},
		},
		"stop-application": {
			Name:   "stop-application",
			Fields: fields_stop_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopApplication(ctx, input)
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
		"update-application-settings": {
			Name:   "update-application-settings",
			Fields: fields_update_application_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssmsap", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
