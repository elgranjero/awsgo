package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sfn"
)

var fields_create_activity = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_state_machine = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "*string", Required: true},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Publish", Flag: "publish", Type: "bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TracingConfiguration", Flag: "tracing-configuration", Type: "*types.TracingConfiguration", Required: false},
	{Name: "Type", Flag: "type", Type: "types.StateMachineType", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_state_machine_alias = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoutingConfiguration", Flag: "routing-configuration", Type: "[]types.RoutingConfigurationListItem", Required: true},
}

var fields_delete_activity = []leanruntime.Field{
	{Name: "ActivityArn", Flag: "activity-arn", Type: "*string", Required: true},
}

var fields_delete_state_machine = []leanruntime.Field{
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
}

var fields_delete_state_machine_alias = []leanruntime.Field{
	{Name: "StateMachineAliasArn", Flag: "state-machine-alias-arn", Type: "*string", Required: true},
}

var fields_delete_state_machine_version = []leanruntime.Field{
	{Name: "StateMachineVersionArn", Flag: "state-machine-version-arn", Type: "*string", Required: true},
}

var fields_describe_activity = []leanruntime.Field{
	{Name: "ActivityArn", Flag: "activity-arn", Type: "*string", Required: true},
}

var fields_describe_execution = []leanruntime.Field{
	{Name: "ExecutionArn", Flag: "execution-arn", Type: "*string", Required: true},
	{Name: "IncludedData", Flag: "included-data", Type: "types.IncludedData", Required: false},
}

var fields_describe_map_run = []leanruntime.Field{
	{Name: "MapRunArn", Flag: "map-run-arn", Type: "*string", Required: true},
}

var fields_describe_state_machine = []leanruntime.Field{
	{Name: "IncludedData", Flag: "included-data", Type: "types.IncludedData", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
}

var fields_describe_state_machine_alias = []leanruntime.Field{
	{Name: "StateMachineAliasArn", Flag: "state-machine-alias-arn", Type: "*string", Required: true},
}

var fields_describe_state_machine_for_execution = []leanruntime.Field{
	{Name: "ExecutionArn", Flag: "execution-arn", Type: "*string", Required: true},
	{Name: "IncludedData", Flag: "included-data", Type: "types.IncludedData", Required: false},
}

var fields_get_activity_task = []leanruntime.Field{
	{Name: "ActivityArn", Flag: "activity-arn", Type: "*string", Required: true},
	{Name: "WorkerName", Flag: "worker-name", Type: "*string", Required: false},
}

var fields_get_execution_history = []leanruntime.Field{
	{Name: "ExecutionArn", Flag: "execution-arn", Type: "*string", Required: true},
	{Name: "IncludeExecutionData", Flag: "include-execution-data", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
}

var fields_list_activities = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_executions = []leanruntime.Field{
	{Name: "MapRunArn", Flag: "map-run-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RedriveFilter", Flag: "redrive-filter", Type: "types.ExecutionRedriveFilter", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.ExecutionStatus", Required: false},
}

var fields_list_map_runs = []leanruntime.Field{
	{Name: "ExecutionArn", Flag: "execution-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_state_machine_aliases = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
}

var fields_list_state_machine_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
}

var fields_list_state_machines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_publish_state_machine_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
}

var fields_redrive_execution = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExecutionArn", Flag: "execution-arn", Type: "*string", Required: true},
}

var fields_send_task_failure = []leanruntime.Field{
	{Name: "Cause", Flag: "cause", Type: "*string", Required: false},
	{Name: "Error", Flag: "error", Type: "*string", Required: false},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_send_task_heartbeat = []leanruntime.Field{
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_send_task_success = []leanruntime.Field{
	{Name: "Output", Flag: "output", Type: "*string", Required: true},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_start_execution = []leanruntime.Field{
	{Name: "Input", Flag: "input", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
	{Name: "TraceHeader", Flag: "trace-header", Type: "*string", Required: false},
}

var fields_start_sync_execution = []leanruntime.Field{
	{Name: "IncludedData", Flag: "included-data", Type: "types.IncludedData", Required: false},
	{Name: "Input", Flag: "input", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
	{Name: "TraceHeader", Flag: "trace-header", Type: "*string", Required: false},
}

var fields_stop_execution = []leanruntime.Field{
	{Name: "Cause", Flag: "cause", Type: "*string", Required: false},
	{Name: "Error", Flag: "error", Type: "*string", Required: false},
	{Name: "ExecutionArn", Flag: "execution-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_state = []leanruntime.Field{
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*string", Required: true},
	{Name: "Input", Flag: "input", Type: "*string", Required: false},
	{Name: "InspectionLevel", Flag: "inspection-level", Type: "types.InspectionLevel", Required: false},
	{Name: "Mock", Flag: "mock", Type: "*types.MockInput", Required: false},
	{Name: "RevealSecrets", Flag: "reveal-secrets", Type: "bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StateConfiguration", Flag: "state-configuration", Type: "*types.TestStateConfiguration", Required: false},
	{Name: "StateName", Flag: "state-name", Type: "*string", Required: false},
	{Name: "Variables", Flag: "variables", Type: "*string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_map_run = []leanruntime.Field{
	{Name: "MapRunArn", Flag: "map-run-arn", Type: "*string", Required: true},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*int32", Required: false},
	{Name: "ToleratedFailureCount", Flag: "tolerated-failure-count", Type: "*int64", Required: false},
	{Name: "ToleratedFailurePercentage", Flag: "tolerated-failure-percentage", Type: "*float32", Required: false},
}

var fields_update_state_machine = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: false},
	{Name: "Publish", Flag: "publish", Type: "bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StateMachineArn", Flag: "state-machine-arn", Type: "*string", Required: true},
	{Name: "TracingConfiguration", Flag: "tracing-configuration", Type: "*types.TracingConfiguration", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_update_state_machine_alias = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RoutingConfiguration", Flag: "routing-configuration", Type: "[]types.RoutingConfigurationListItem", Required: false},
	{Name: "StateMachineAliasArn", Flag: "state-machine-alias-arn", Type: "*string", Required: true},
}

var fields_validate_state_machine_definition = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "Severity", Flag: "severity", Type: "types.ValidateStateMachineDefinitionSeverity", Required: false},
	{Name: "Type", Flag: "type", Type: "types.StateMachineType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-activity": {
			Name:   "create-activity",
			Fields: fields_create_activity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateActivityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_activity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateActivity(ctx, input)
			},
		},
		"create-state-machine": {
			Name:   "create-state-machine",
			Fields: fields_create_state_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStateMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_state_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStateMachine(ctx, input)
			},
		},
		"create-state-machine-alias": {
			Name:   "create-state-machine-alias",
			Fields: fields_create_state_machine_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStateMachineAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_state_machine_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStateMachineAlias(ctx, input)
			},
		},
		"delete-activity": {
			Name:   "delete-activity",
			Fields: fields_delete_activity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteActivityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_activity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteActivity(ctx, input)
			},
		},
		"delete-state-machine": {
			Name:   "delete-state-machine",
			Fields: fields_delete_state_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStateMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_state_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStateMachine(ctx, input)
			},
		},
		"delete-state-machine-alias": {
			Name:   "delete-state-machine-alias",
			Fields: fields_delete_state_machine_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStateMachineAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_state_machine_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStateMachineAlias(ctx, input)
			},
		},
		"delete-state-machine-version": {
			Name:   "delete-state-machine-version",
			Fields: fields_delete_state_machine_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStateMachineVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_state_machine_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStateMachineVersion(ctx, input)
			},
		},
		"describe-activity": {
			Name:   "describe-activity",
			Fields: fields_describe_activity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActivityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_activity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeActivity(ctx, input)
			},
		},
		"describe-execution": {
			Name:   "describe-execution",
			Fields: fields_describe_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExecution(ctx, input)
			},
		},
		"describe-map-run": {
			Name:   "describe-map-run",
			Fields: fields_describe_map_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMapRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_map_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMapRun(ctx, input)
			},
		},
		"describe-state-machine": {
			Name:   "describe-state-machine",
			Fields: fields_describe_state_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStateMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_state_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStateMachine(ctx, input)
			},
		},
		"describe-state-machine-alias": {
			Name:   "describe-state-machine-alias",
			Fields: fields_describe_state_machine_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStateMachineAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_state_machine_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStateMachineAlias(ctx, input)
			},
		},
		"describe-state-machine-for-execution": {
			Name:   "describe-state-machine-for-execution",
			Fields: fields_describe_state_machine_for_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStateMachineForExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_state_machine_for_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStateMachineForExecution(ctx, input)
			},
		},
		"get-activity-task": {
			Name:   "get-activity-task",
			Fields: fields_get_activity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetActivityTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_activity_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetActivityTask(ctx, input)
			},
		},
		"get-execution-history": {
			Name:   "get-execution-history",
			Fields: fields_get_execution_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExecutionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_execution_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetExecutionHistory(ctx, input)
				}
				var results []*svc.GetExecutionHistoryOutput
				p := svc.NewGetExecutionHistoryPaginator(client, input)
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
		"list-activities": {
			Name:   "list-activities",
			Fields: fields_list_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActivitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_activities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActivities(ctx, input)
				}
				var results []*svc.ListActivitiesOutput
				p := svc.NewListActivitiesPaginator(client, input)
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
		"list-executions": {
			Name:   "list-executions",
			Fields: fields_list_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExecutions(ctx, input)
				}
				var results []*svc.ListExecutionsOutput
				p := svc.NewListExecutionsPaginator(client, input)
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
		"list-map-runs": {
			Name:   "list-map-runs",
			Fields: fields_list_map_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMapRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_map_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMapRuns(ctx, input)
				}
				var results []*svc.ListMapRunsOutput
				p := svc.NewListMapRunsPaginator(client, input)
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
		"list-state-machine-aliases": {
			Name:   "list-state-machine-aliases",
			Fields: fields_list_state_machine_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStateMachineAliasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_state_machine_aliases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStateMachineAliases(ctx, input)
			},
		},
		"list-state-machine-versions": {
			Name:   "list-state-machine-versions",
			Fields: fields_list_state_machine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStateMachineVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_state_machine_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStateMachineVersions(ctx, input)
			},
		},
		"list-state-machines": {
			Name:   "list-state-machines",
			Fields: fields_list_state_machines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStateMachinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_state_machines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStateMachines(ctx, input)
				}
				var results []*svc.ListStateMachinesOutput
				p := svc.NewListStateMachinesPaginator(client, input)
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
		"publish-state-machine-version": {
			Name:   "publish-state-machine-version",
			Fields: fields_publish_state_machine_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishStateMachineVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_state_machine_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishStateMachineVersion(ctx, input)
			},
		},
		"redrive-execution": {
			Name:   "redrive-execution",
			Fields: fields_redrive_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RedriveExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_redrive_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RedriveExecution(ctx, input)
			},
		},
		"send-task-failure": {
			Name:   "send-task-failure",
			Fields: fields_send_task_failure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendTaskFailureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_task_failure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendTaskFailure(ctx, input)
			},
		},
		"send-task-heartbeat": {
			Name:   "send-task-heartbeat",
			Fields: fields_send_task_heartbeat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendTaskHeartbeatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_task_heartbeat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendTaskHeartbeat(ctx, input)
			},
		},
		"send-task-success": {
			Name:   "send-task-success",
			Fields: fields_send_task_success,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendTaskSuccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_task_success, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendTaskSuccess(ctx, input)
			},
		},
		"start-execution": {
			Name:   "start-execution",
			Fields: fields_start_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExecution(ctx, input)
			},
		},
		"start-sync-execution": {
			Name:   "start-sync-execution",
			Fields: fields_start_sync_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSyncExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_sync_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSyncExecution(ctx, input)
			},
		},
		"stop-execution": {
			Name:   "stop-execution",
			Fields: fields_stop_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopExecution(ctx, input)
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
		"test-state": {
			Name:   "test-state",
			Fields: fields_test_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestState(ctx, input)
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
		"update-map-run": {
			Name:   "update-map-run",
			Fields: fields_update_map_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMapRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_map_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMapRun(ctx, input)
			},
		},
		"update-state-machine": {
			Name:   "update-state-machine",
			Fields: fields_update_state_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStateMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_state_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStateMachine(ctx, input)
			},
		},
		"update-state-machine-alias": {
			Name:   "update-state-machine-alias",
			Fields: fields_update_state_machine_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStateMachineAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_state_machine_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStateMachineAlias(ctx, input)
			},
		},
		"validate-state-machine-definition": {
			Name:   "validate-state-machine-definition",
			Fields: fields_validate_state_machine_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateStateMachineDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_state_machine_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateStateMachineDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sfn", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
