package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/applicationsignals"
)

var fields_batch_get_service_level_objective_budget_report = []leanruntime.Field{
	{Name: "SloIds", Flag: "slo-ids", Type: "[]string", Required: true},
	{Name: "Timestamp", Flag: "timestamp", Type: "*time.Time", Required: true},
}

var fields_batch_update_exclusion_windows = []leanruntime.Field{
	{Name: "AddExclusionWindows", Flag: "add-exclusion-windows", Type: "[]types.ExclusionWindow", Required: false},
	{Name: "RemoveExclusionWindows", Flag: "remove-exclusion-windows", Type: "[]types.ExclusionWindow", Required: false},
	{Name: "SloIds", Flag: "slo-ids", Type: "[]string", Required: true},
}

var fields_create_service_level_objective = []leanruntime.Field{
	{Name: "BurnRateConfigurations", Flag: "burn-rate-configurations", Type: "[]types.BurnRateConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Goal", Flag: "goal", Type: "*types.Goal", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestBasedSliConfig", Flag: "request-based-sli-config", Type: "*types.RequestBasedServiceLevelIndicatorConfig", Required: false},
	{Name: "SliConfig", Flag: "sli-config", Type: "*types.ServiceLevelIndicatorConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_grouping_configuration = []leanruntime.Field{}

var fields_delete_service_level_objective = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_service = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "KeyAttributes", Flag: "key-attributes", Type: "map[string]string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_service_level_objective = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_audit_findings = []leanruntime.Field{
	{Name: "AuditTargets", Flag: "audit-targets", Type: "[]types.AuditTarget", Required: true},
	{Name: "Auditors", Flag: "auditors", Type: "[]string", Required: false},
	{Name: "DetailLevel", Flag: "detail-level", Type: "types.DetailLevel", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_entity_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Entity", Flag: "entity", Type: "map[string]string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_grouping_attribute_definitions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "bool", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_service_dependencies = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "KeyAttributes", Flag: "key-attributes", Type: "map[string]string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_service_dependents = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "KeyAttributes", Flag: "key-attributes", Type: "map[string]string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_service_level_objective_exclusion_windows = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_service_level_objectives = []leanruntime.Field{
	{Name: "DependencyConfig", Flag: "dependency-config", Type: "*types.DependencyConfig", Required: false},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "bool", Required: false},
	{Name: "KeyAttributes", Flag: "key-attributes", Type: "map[string]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricSourceTypes", Flag: "metric-source-types", Type: "[]types.MetricSourceType", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperationName", Flag: "operation-name", Type: "*string", Required: false},
	{Name: "SloOwnerAwsAccountId", Flag: "slo-owner-aws-account-id", Type: "*string", Required: false},
}

var fields_list_service_operations = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "KeyAttributes", Flag: "key-attributes", Type: "map[string]string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_service_states = []leanruntime.Field{
	{Name: "AttributeFilters", Flag: "attribute-filters", Type: "[]types.AttributeFilter", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_services = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_grouping_configuration = []leanruntime.Field{
	{Name: "GroupingAttributeDefinitions", Flag: "grouping-attribute-definitions", Type: "[]types.GroupingAttributeDefinition", Required: true},
}

var fields_start_discovery = []leanruntime.Field{}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_service_level_objective = []leanruntime.Field{
	{Name: "BurnRateConfigurations", Flag: "burn-rate-configurations", Type: "[]types.BurnRateConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Goal", Flag: "goal", Type: "*types.Goal", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RequestBasedSliConfig", Flag: "request-based-sli-config", Type: "*types.RequestBasedServiceLevelIndicatorConfig", Required: false},
	{Name: "SliConfig", Flag: "sli-config", Type: "*types.ServiceLevelIndicatorConfig", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-service-level-objective-budget-report": {
			Name:   "batch-get-service-level-objective-budget-report",
			Fields: fields_batch_get_service_level_objective_budget_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetServiceLevelObjectiveBudgetReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_service_level_objective_budget_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetServiceLevelObjectiveBudgetReport(ctx, input)
			},
		},
		"batch-update-exclusion-windows": {
			Name:   "batch-update-exclusion-windows",
			Fields: fields_batch_update_exclusion_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateExclusionWindowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_exclusion_windows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateExclusionWindows(ctx, input)
			},
		},
		"create-service-level-objective": {
			Name:   "create-service-level-objective",
			Fields: fields_create_service_level_objective,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceLevelObjectiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_level_objective, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceLevelObjective(ctx, input)
			},
		},
		"delete-grouping-configuration": {
			Name:   "delete-grouping-configuration",
			Fields: fields_delete_grouping_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_grouping_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroupingConfiguration(ctx, input)
			},
		},
		"delete-service-level-objective": {
			Name:   "delete-service-level-objective",
			Fields: fields_delete_service_level_objective,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceLevelObjectiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_level_objective, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceLevelObjective(ctx, input)
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
		"get-service-level-objective": {
			Name:   "get-service-level-objective",
			Fields: fields_get_service_level_objective,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceLevelObjectiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_level_objective, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceLevelObjective(ctx, input)
			},
		},
		"list-audit-findings": {
			Name:   "list-audit-findings",
			Fields: fields_list_audit_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuditFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_audit_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAuditFindings(ctx, input)
			},
		},
		"list-entity-events": {
			Name:   "list-entity-events",
			Fields: fields_list_entity_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntityEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entity_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntityEvents(ctx, input)
				}
				var results []*svc.ListEntityEventsOutput
				p := svc.NewListEntityEventsPaginator(client, input)
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
		"list-grouping-attribute-definitions": {
			Name:   "list-grouping-attribute-definitions",
			Fields: fields_list_grouping_attribute_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupingAttributeDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_grouping_attribute_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGroupingAttributeDefinitions(ctx, input)
			},
		},
		"list-service-dependencies": {
			Name:   "list-service-dependencies",
			Fields: fields_list_service_dependencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceDependenciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_dependencies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceDependencies(ctx, input)
				}
				var results []*svc.ListServiceDependenciesOutput
				p := svc.NewListServiceDependenciesPaginator(client, input)
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
		"list-service-dependents": {
			Name:   "list-service-dependents",
			Fields: fields_list_service_dependents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceDependentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_dependents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceDependents(ctx, input)
				}
				var results []*svc.ListServiceDependentsOutput
				p := svc.NewListServiceDependentsPaginator(client, input)
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
		"list-service-level-objective-exclusion-windows": {
			Name:   "list-service-level-objective-exclusion-windows",
			Fields: fields_list_service_level_objective_exclusion_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceLevelObjectiveExclusionWindowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_level_objective_exclusion_windows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceLevelObjectiveExclusionWindows(ctx, input)
				}
				var results []*svc.ListServiceLevelObjectiveExclusionWindowsOutput
				p := svc.NewListServiceLevelObjectiveExclusionWindowsPaginator(client, input)
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
		"list-service-level-objectives": {
			Name:   "list-service-level-objectives",
			Fields: fields_list_service_level_objectives,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceLevelObjectivesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_level_objectives, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceLevelObjectives(ctx, input)
				}
				var results []*svc.ListServiceLevelObjectivesOutput
				p := svc.NewListServiceLevelObjectivesPaginator(client, input)
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
		"list-service-operations": {
			Name:   "list-service-operations",
			Fields: fields_list_service_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceOperations(ctx, input)
				}
				var results []*svc.ListServiceOperationsOutput
				p := svc.NewListServiceOperationsPaginator(client, input)
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
		"list-service-states": {
			Name:   "list-service-states",
			Fields: fields_list_service_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceStatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_states, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceStates(ctx, input)
				}
				var results []*svc.ListServiceStatesOutput
				p := svc.NewListServiceStatesPaginator(client, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"put-grouping-configuration": {
			Name:   "put-grouping-configuration",
			Fields: fields_put_grouping_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGroupingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_grouping_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGroupingConfiguration(ctx, input)
			},
		},
		"start-discovery": {
			Name:   "start-discovery",
			Fields: fields_start_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDiscovery(ctx, input)
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
		"update-service-level-objective": {
			Name:   "update-service-level-objective",
			Fields: fields_update_service_level_objective,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceLevelObjectiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_level_objective, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceLevelObjective(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("applicationsignals", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
