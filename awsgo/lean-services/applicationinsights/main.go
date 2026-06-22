package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/applicationinsights"
)

var fields_add_workload = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "WorkloadConfiguration", Flag: "workload-configuration", Type: "*types.WorkloadConfiguration", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "AttachMissingPermission", Flag: "attach-missing-permission", Type: "*bool", Required: false},
	{Name: "AutoConfigEnabled", Flag: "auto-config-enabled", Type: "*bool", Required: false},
	{Name: "AutoCreate", Flag: "auto-create", Type: "*bool", Required: false},
	{Name: "CWEMonitorEnabled", Flag: "cwe-monitor-enabled", Type: "*bool", Required: false},
	{Name: "GroupingType", Flag: "grouping-type", Type: "types.GroupingType", Required: false},
	{Name: "OpsCenterEnabled", Flag: "ops-center-enabled", Type: "*bool", Required: false},
	{Name: "OpsItemSNSTopicArn", Flag: "ops-item-sns-topic-arn", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: false},
	{Name: "SNSNotificationArn", Flag: "sns-notification-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_component = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "ResourceList", Flag: "resource-list", Type: "[]string", Required: true},
}

var fields_create_log_pattern = []leanruntime.Field{
	{Name: "Pattern", Flag: "pattern", Type: "*string", Required: true},
	{Name: "PatternName", Flag: "pattern-name", Type: "*string", Required: true},
	{Name: "PatternSetName", Flag: "pattern-set-name", Type: "*string", Required: true},
	{Name: "Rank", Flag: "rank", Type: "int32", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_delete_component = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_delete_log_pattern = []leanruntime.Field{
	{Name: "PatternName", Flag: "pattern-name", Type: "*string", Required: true},
	{Name: "PatternSetName", Flag: "pattern-set-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_describe_application = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_describe_component = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_describe_component_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_describe_component_configuration_recommendation = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "RecommendationType", Flag: "recommendation-type", Type: "types.RecommendationType", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "Tier", Flag: "tier", Type: "types.Tier", Required: true},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: false},
}

var fields_describe_log_pattern = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "PatternName", Flag: "pattern-name", Type: "*string", Required: true},
	{Name: "PatternSetName", Flag: "pattern-set-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_describe_observation = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ObservationId", Flag: "observation-id", Type: "*string", Required: true},
}

var fields_describe_problem = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ProblemId", Flag: "problem-id", Type: "*string", Required: true},
}

var fields_describe_problem_observations = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ProblemId", Flag: "problem-id", Type: "*string", Required: true},
}

var fields_describe_workload = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_components = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_list_configuration_history = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EventStatus", Flag: "event-status", Type: "types.ConfigurationEventStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_log_pattern_sets = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_list_log_patterns = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PatternSetName", Flag: "pattern-set-name", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_list_problems = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Visibility", Flag: "visibility", Type: "types.Visibility", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_workloads = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_remove_workload = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "AttachMissingPermission", Flag: "attach-missing-permission", Type: "*bool", Required: false},
	{Name: "AutoConfigEnabled", Flag: "auto-config-enabled", Type: "*bool", Required: false},
	{Name: "CWEMonitorEnabled", Flag: "cwe-monitor-enabled", Type: "*bool", Required: false},
	{Name: "OpsCenterEnabled", Flag: "ops-center-enabled", Type: "*bool", Required: false},
	{Name: "OpsItemSNSTopicArn", Flag: "ops-item-sns-topic-arn", Type: "*string", Required: false},
	{Name: "RemoveSNSTopic", Flag: "remove-sns-topic", Type: "*bool", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "SNSNotificationArn", Flag: "sns-notification-arn", Type: "*string", Required: false},
}

var fields_update_component = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "NewComponentName", Flag: "new-component-name", Type: "*string", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "ResourceList", Flag: "resource-list", Type: "[]string", Required: false},
}

var fields_update_component_configuration = []leanruntime.Field{
	{Name: "AutoConfigEnabled", Flag: "auto-config-enabled", Type: "*bool", Required: false},
	{Name: "ComponentConfiguration", Flag: "component-configuration", Type: "*string", Required: false},
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "Monitor", Flag: "monitor", Type: "*bool", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "Tier", Flag: "tier", Type: "types.Tier", Required: false},
}

var fields_update_log_pattern = []leanruntime.Field{
	{Name: "Pattern", Flag: "pattern", Type: "*string", Required: false},
	{Name: "PatternName", Flag: "pattern-name", Type: "*string", Required: true},
	{Name: "PatternSetName", Flag: "pattern-set-name", Type: "*string", Required: true},
	{Name: "Rank", Flag: "rank", Type: "int32", Required: false},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
}

var fields_update_problem = []leanruntime.Field{
	{Name: "ProblemId", Flag: "problem-id", Type: "*string", Required: true},
	{Name: "UpdateStatus", Flag: "update-status", Type: "types.UpdateStatus", Required: false},
	{Name: "Visibility", Flag: "visibility", Type: "types.Visibility", Required: false},
}

var fields_update_workload = []leanruntime.Field{
	{Name: "ComponentName", Flag: "component-name", Type: "*string", Required: true},
	{Name: "ResourceGroupName", Flag: "resource-group-name", Type: "*string", Required: true},
	{Name: "WorkloadConfiguration", Flag: "workload-configuration", Type: "*types.WorkloadConfiguration", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-workload": {
			Name:   "add-workload",
			Fields: fields_add_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddWorkload(ctx, input)
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
		"create-log-pattern": {
			Name:   "create-log-pattern",
			Fields: fields_create_log_pattern,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLogPatternInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_log_pattern, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLogPattern(ctx, input)
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
		"delete-log-pattern": {
			Name:   "delete-log-pattern",
			Fields: fields_delete_log_pattern,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLogPatternInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_log_pattern, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLogPattern(ctx, input)
			},
		},
		"describe-application": {
			Name:   "describe-application",
			Fields: fields_describe_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplication(ctx, input)
			},
		},
		"describe-component": {
			Name:   "describe-component",
			Fields: fields_describe_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComponent(ctx, input)
			},
		},
		"describe-component-configuration": {
			Name:   "describe-component-configuration",
			Fields: fields_describe_component_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComponentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_component_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComponentConfiguration(ctx, input)
			},
		},
		"describe-component-configuration-recommendation": {
			Name:   "describe-component-configuration-recommendation",
			Fields: fields_describe_component_configuration_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComponentConfigurationRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_component_configuration_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComponentConfigurationRecommendation(ctx, input)
			},
		},
		"describe-log-pattern": {
			Name:   "describe-log-pattern",
			Fields: fields_describe_log_pattern,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLogPatternInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_log_pattern, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLogPattern(ctx, input)
			},
		},
		"describe-observation": {
			Name:   "describe-observation",
			Fields: fields_describe_observation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeObservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_observation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeObservation(ctx, input)
			},
		},
		"describe-problem": {
			Name:   "describe-problem",
			Fields: fields_describe_problem,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProblemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_problem, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProblem(ctx, input)
			},
		},
		"describe-problem-observations": {
			Name:   "describe-problem-observations",
			Fields: fields_describe_problem_observations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProblemObservationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_problem_observations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProblemObservations(ctx, input)
			},
		},
		"describe-workload": {
			Name:   "describe-workload",
			Fields: fields_describe_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkload(ctx, input)
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
		"list-configuration-history": {
			Name:   "list-configuration-history",
			Fields: fields_list_configuration_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationHistory(ctx, input)
				}
				var results []*svc.ListConfigurationHistoryOutput
				p := svc.NewListConfigurationHistoryPaginator(client, input)
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
		"list-log-pattern-sets": {
			Name:   "list-log-pattern-sets",
			Fields: fields_list_log_pattern_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogPatternSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_log_pattern_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLogPatternSets(ctx, input)
				}
				var results []*svc.ListLogPatternSetsOutput
				p := svc.NewListLogPatternSetsPaginator(client, input)
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
		"list-log-patterns": {
			Name:   "list-log-patterns",
			Fields: fields_list_log_patterns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogPatternsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_log_patterns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLogPatterns(ctx, input)
				}
				var results []*svc.ListLogPatternsOutput
				p := svc.NewListLogPatternsPaginator(client, input)
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
		"list-problems": {
			Name:   "list-problems",
			Fields: fields_list_problems,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProblemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_problems, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProblems(ctx, input)
				}
				var results []*svc.ListProblemsOutput
				p := svc.NewListProblemsPaginator(client, input)
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
		"list-workloads": {
			Name:   "list-workloads",
			Fields: fields_list_workloads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workloads, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloads(ctx, input)
				}
				var results []*svc.ListWorkloadsOutput
				p := svc.NewListWorkloadsPaginator(client, input)
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
		"remove-workload": {
			Name:   "remove-workload",
			Fields: fields_remove_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveWorkload(ctx, input)
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
		"update-component-configuration": {
			Name:   "update-component-configuration",
			Fields: fields_update_component_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComponentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_component_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComponentConfiguration(ctx, input)
			},
		},
		"update-log-pattern": {
			Name:   "update-log-pattern",
			Fields: fields_update_log_pattern,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLogPatternInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_log_pattern, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLogPattern(ctx, input)
			},
		},
		"update-problem": {
			Name:   "update-problem",
			Fields: fields_update_problem,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProblemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_problem, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProblem(ctx, input)
			},
		},
		"update-workload": {
			Name:   "update-workload",
			Fields: fields_update_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkload(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("applicationinsights", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
