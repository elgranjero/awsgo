package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/swf"
)

var fields_count_closed_workflow_executions = []leanruntime.Field{
	{Name: "CloseStatusFilter", Flag: "close-status-filter", Type: "*types.CloseStatusFilter", Required: false},
	{Name: "CloseTimeFilter", Flag: "close-time-filter", Type: "*types.ExecutionTimeFilter", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ExecutionFilter", Flag: "execution-filter", Type: "*types.WorkflowExecutionFilter", Required: false},
	{Name: "StartTimeFilter", Flag: "start-time-filter", Type: "*types.ExecutionTimeFilter", Required: false},
	{Name: "TagFilter", Flag: "tag-filter", Type: "*types.TagFilter", Required: false},
	{Name: "TypeFilter", Flag: "type-filter", Type: "*types.WorkflowTypeFilter", Required: false},
}

var fields_count_open_workflow_executions = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ExecutionFilter", Flag: "execution-filter", Type: "*types.WorkflowExecutionFilter", Required: false},
	{Name: "StartTimeFilter", Flag: "start-time-filter", Type: "*types.ExecutionTimeFilter", Required: true},
	{Name: "TagFilter", Flag: "tag-filter", Type: "*types.TagFilter", Required: false},
	{Name: "TypeFilter", Flag: "type-filter", Type: "*types.WorkflowTypeFilter", Required: false},
}

var fields_count_pending_activity_tasks = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "TaskList", Flag: "task-list", Type: "*types.TaskList", Required: true},
}

var fields_count_pending_decision_tasks = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "TaskList", Flag: "task-list", Type: "*types.TaskList", Required: true},
}

var fields_delete_activity_type = []leanruntime.Field{
	{Name: "ActivityType", Flag: "activity-type", Type: "*types.ActivityType", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_delete_workflow_type = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "*types.WorkflowType", Required: true},
}

var fields_deprecate_activity_type = []leanruntime.Field{
	{Name: "ActivityType", Flag: "activity-type", Type: "*types.ActivityType", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_deprecate_domain = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_deprecate_workflow_type = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "*types.WorkflowType", Required: true},
}

var fields_describe_activity_type = []leanruntime.Field{
	{Name: "ActivityType", Flag: "activity-type", Type: "*types.ActivityType", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_describe_domain = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_workflow_execution = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Execution", Flag: "execution", Type: "*types.WorkflowExecution", Required: true},
}

var fields_describe_workflow_type = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "*types.WorkflowType", Required: true},
}

var fields_get_workflow_execution_history = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Execution", Flag: "execution", Type: "*types.WorkflowExecution", Required: true},
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
}

var fields_list_activity_types = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "RegistrationStatus", Flag: "registration-status", Type: "types.RegistrationStatus", Required: true},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
}

var fields_list_closed_workflow_executions = []leanruntime.Field{
	{Name: "CloseStatusFilter", Flag: "close-status-filter", Type: "*types.CloseStatusFilter", Required: false},
	{Name: "CloseTimeFilter", Flag: "close-time-filter", Type: "*types.ExecutionTimeFilter", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ExecutionFilter", Flag: "execution-filter", Type: "*types.WorkflowExecutionFilter", Required: false},
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
	{Name: "StartTimeFilter", Flag: "start-time-filter", Type: "*types.ExecutionTimeFilter", Required: false},
	{Name: "TagFilter", Flag: "tag-filter", Type: "*types.TagFilter", Required: false},
	{Name: "TypeFilter", Flag: "type-filter", Type: "*types.WorkflowTypeFilter", Required: false},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "RegistrationStatus", Flag: "registration-status", Type: "types.RegistrationStatus", Required: true},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
}

var fields_list_open_workflow_executions = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ExecutionFilter", Flag: "execution-filter", Type: "*types.WorkflowExecutionFilter", Required: false},
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
	{Name: "StartTimeFilter", Flag: "start-time-filter", Type: "*types.ExecutionTimeFilter", Required: true},
	{Name: "TagFilter", Flag: "tag-filter", Type: "*types.TagFilter", Required: false},
	{Name: "TypeFilter", Flag: "type-filter", Type: "*types.WorkflowTypeFilter", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_workflow_types = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "RegistrationStatus", Flag: "registration-status", Type: "types.RegistrationStatus", Required: true},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
}

var fields_poll_for_activity_task = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Identity", Flag: "identity", Type: "*string", Required: false},
	{Name: "TaskList", Flag: "task-list", Type: "*types.TaskList", Required: true},
}

var fields_poll_for_decision_task = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Identity", Flag: "identity", Type: "*string", Required: false},
	{Name: "MaximumPageSize", Flag: "maximum-page-size", Type: "int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "bool", Required: false},
	{Name: "StartAtPreviousStartedEvent", Flag: "start-at-previous-started-event", Type: "bool", Required: false},
	{Name: "TaskList", Flag: "task-list", Type: "*types.TaskList", Required: true},
}

var fields_record_activity_task_heartbeat = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "*string", Required: false},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_register_activity_type = []leanruntime.Field{
	{Name: "DefaultTaskHeartbeatTimeout", Flag: "default-task-heartbeat-timeout", Type: "*string", Required: false},
	{Name: "DefaultTaskList", Flag: "default-task-list", Type: "*types.TaskList", Required: false},
	{Name: "DefaultTaskPriority", Flag: "default-task-priority", Type: "*string", Required: false},
	{Name: "DefaultTaskScheduleToCloseTimeout", Flag: "default-task-schedule-to-close-timeout", Type: "*string", Required: false},
	{Name: "DefaultTaskScheduleToStartTimeout", Flag: "default-task-schedule-to-start-timeout", Type: "*string", Required: false},
	{Name: "DefaultTaskStartToCloseTimeout", Flag: "default-task-start-to-close-timeout", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_register_domain = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "WorkflowExecutionRetentionPeriodInDays", Flag: "workflow-execution-retention-period-in-days", Type: "*string", Required: true},
}

var fields_register_workflow_type = []leanruntime.Field{
	{Name: "DefaultChildPolicy", Flag: "default-child-policy", Type: "types.ChildPolicy", Required: false},
	{Name: "DefaultExecutionStartToCloseTimeout", Flag: "default-execution-start-to-close-timeout", Type: "*string", Required: false},
	{Name: "DefaultLambdaRole", Flag: "default-lambda-role", Type: "*string", Required: false},
	{Name: "DefaultTaskList", Flag: "default-task-list", Type: "*types.TaskList", Required: false},
	{Name: "DefaultTaskPriority", Flag: "default-task-priority", Type: "*string", Required: false},
	{Name: "DefaultTaskStartToCloseTimeout", Flag: "default-task-start-to-close-timeout", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_request_cancel_workflow_execution = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_respond_activity_task_canceled = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "*string", Required: false},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_respond_activity_task_completed = []leanruntime.Field{
	{Name: "Result", Flag: "result", Type: "*string", Required: false},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_respond_activity_task_failed = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "*string", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_respond_decision_task_completed = []leanruntime.Field{
	{Name: "Decisions", Flag: "decisions", Type: "[]types.Decision", Required: false},
	{Name: "ExecutionContext", Flag: "execution-context", Type: "*string", Required: false},
	{Name: "TaskList", Flag: "task-list", Type: "*types.TaskList", Required: false},
	{Name: "TaskListScheduleToStartTimeout", Flag: "task-list-schedule-to-start-timeout", Type: "*string", Required: false},
	{Name: "TaskToken", Flag: "task-token", Type: "*string", Required: true},
}

var fields_signal_workflow_execution = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Input", Flag: "input", Type: "*string", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
	{Name: "SignalName", Flag: "signal-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_start_workflow_execution = []leanruntime.Field{
	{Name: "ChildPolicy", Flag: "child-policy", Type: "types.ChildPolicy", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ExecutionStartToCloseTimeout", Flag: "execution-start-to-close-timeout", Type: "*string", Required: false},
	{Name: "Input", Flag: "input", Type: "*string", Required: false},
	{Name: "LambdaRole", Flag: "lambda-role", Type: "*string", Required: false},
	{Name: "TagList", Flag: "tag-list", Type: "[]string", Required: false},
	{Name: "TaskList", Flag: "task-list", Type: "*types.TaskList", Required: false},
	{Name: "TaskPriority", Flag: "task-priority", Type: "*string", Required: false},
	{Name: "TaskStartToCloseTimeout", Flag: "task-start-to-close-timeout", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "*types.WorkflowType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.ResourceTag", Required: true},
}

var fields_terminate_workflow_execution = []leanruntime.Field{
	{Name: "ChildPolicy", Flag: "child-policy", Type: "types.ChildPolicy", Required: false},
	{Name: "Details", Flag: "details", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_undeprecate_activity_type = []leanruntime.Field{
	{Name: "ActivityType", Flag: "activity-type", Type: "*types.ActivityType", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_undeprecate_domain = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_undeprecate_workflow_type = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "*types.WorkflowType", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"count-closed-workflow-executions": {
			Name:   "count-closed-workflow-executions",
			Fields: fields_count_closed_workflow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CountClosedWorkflowExecutionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_count_closed_workflow_executions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CountClosedWorkflowExecutions(ctx, input)
			},
		},
		"count-open-workflow-executions": {
			Name:   "count-open-workflow-executions",
			Fields: fields_count_open_workflow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CountOpenWorkflowExecutionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_count_open_workflow_executions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CountOpenWorkflowExecutions(ctx, input)
			},
		},
		"count-pending-activity-tasks": {
			Name:   "count-pending-activity-tasks",
			Fields: fields_count_pending_activity_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CountPendingActivityTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_count_pending_activity_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CountPendingActivityTasks(ctx, input)
			},
		},
		"count-pending-decision-tasks": {
			Name:   "count-pending-decision-tasks",
			Fields: fields_count_pending_decision_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CountPendingDecisionTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_count_pending_decision_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CountPendingDecisionTasks(ctx, input)
			},
		},
		"delete-activity-type": {
			Name:   "delete-activity-type",
			Fields: fields_delete_activity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteActivityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_activity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteActivityType(ctx, input)
			},
		},
		"delete-workflow-type": {
			Name:   "delete-workflow-type",
			Fields: fields_delete_workflow_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflowType(ctx, input)
			},
		},
		"deprecate-activity-type": {
			Name:   "deprecate-activity-type",
			Fields: fields_deprecate_activity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprecateActivityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprecate_activity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprecateActivityType(ctx, input)
			},
		},
		"deprecate-domain": {
			Name:   "deprecate-domain",
			Fields: fields_deprecate_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprecateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprecate_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprecateDomain(ctx, input)
			},
		},
		"deprecate-workflow-type": {
			Name:   "deprecate-workflow-type",
			Fields: fields_deprecate_workflow_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprecateWorkflowTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprecate_workflow_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprecateWorkflowType(ctx, input)
			},
		},
		"describe-activity-type": {
			Name:   "describe-activity-type",
			Fields: fields_describe_activity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActivityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_activity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeActivityType(ctx, input)
			},
		},
		"describe-domain": {
			Name:   "describe-domain",
			Fields: fields_describe_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomain(ctx, input)
			},
		},
		"describe-workflow-execution": {
			Name:   "describe-workflow-execution",
			Fields: fields_describe_workflow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkflowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workflow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkflowExecution(ctx, input)
			},
		},
		"describe-workflow-type": {
			Name:   "describe-workflow-type",
			Fields: fields_describe_workflow_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkflowTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workflow_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkflowType(ctx, input)
			},
		},
		"get-workflow-execution-history": {
			Name:   "get-workflow-execution-history",
			Fields: fields_get_workflow_execution_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowExecutionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_workflow_execution_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetWorkflowExecutionHistory(ctx, input)
				}
				var results []*svc.GetWorkflowExecutionHistoryOutput
				p := svc.NewGetWorkflowExecutionHistoryPaginator(client, input)
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
		"list-activity-types": {
			Name:   "list-activity-types",
			Fields: fields_list_activity_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActivityTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_activity_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActivityTypes(ctx, input)
				}
				var results []*svc.ListActivityTypesOutput
				p := svc.NewListActivityTypesPaginator(client, input)
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
		"list-closed-workflow-executions": {
			Name:   "list-closed-workflow-executions",
			Fields: fields_list_closed_workflow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClosedWorkflowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_closed_workflow_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClosedWorkflowExecutions(ctx, input)
				}
				var results []*svc.ListClosedWorkflowExecutionsOutput
				p := svc.NewListClosedWorkflowExecutionsPaginator(client, input)
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
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-open-workflow-executions": {
			Name:   "list-open-workflow-executions",
			Fields: fields_list_open_workflow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpenWorkflowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_open_workflow_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpenWorkflowExecutions(ctx, input)
				}
				var results []*svc.ListOpenWorkflowExecutionsOutput
				p := svc.NewListOpenWorkflowExecutionsPaginator(client, input)
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
		"list-workflow-types": {
			Name:   "list-workflow-types",
			Fields: fields_list_workflow_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowTypes(ctx, input)
				}
				var results []*svc.ListWorkflowTypesOutput
				p := svc.NewListWorkflowTypesPaginator(client, input)
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
		"poll-for-activity-task": {
			Name:   "poll-for-activity-task",
			Fields: fields_poll_for_activity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PollForActivityTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_poll_for_activity_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PollForActivityTask(ctx, input)
			},
		},
		"poll-for-decision-task": {
			Name:   "poll-for-decision-task",
			Fields: fields_poll_for_decision_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PollForDecisionTaskInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_poll_for_decision_task, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.PollForDecisionTask(ctx, input)
				}
				var results []*svc.PollForDecisionTaskOutput
				p := svc.NewPollForDecisionTaskPaginator(client, input)
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
		"record-activity-task-heartbeat": {
			Name:   "record-activity-task-heartbeat",
			Fields: fields_record_activity_task_heartbeat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecordActivityTaskHeartbeatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_record_activity_task_heartbeat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecordActivityTaskHeartbeat(ctx, input)
			},
		},
		"register-activity-type": {
			Name:   "register-activity-type",
			Fields: fields_register_activity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterActivityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_activity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterActivityType(ctx, input)
			},
		},
		"register-domain": {
			Name:   "register-domain",
			Fields: fields_register_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDomain(ctx, input)
			},
		},
		"register-workflow-type": {
			Name:   "register-workflow-type",
			Fields: fields_register_workflow_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterWorkflowTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_workflow_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterWorkflowType(ctx, input)
			},
		},
		"request-cancel-workflow-execution": {
			Name:   "request-cancel-workflow-execution",
			Fields: fields_request_cancel_workflow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestCancelWorkflowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_cancel_workflow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestCancelWorkflowExecution(ctx, input)
			},
		},
		"respond-activity-task-canceled": {
			Name:   "respond-activity-task-canceled",
			Fields: fields_respond_activity_task_canceled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RespondActivityTaskCanceledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_respond_activity_task_canceled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RespondActivityTaskCanceled(ctx, input)
			},
		},
		"respond-activity-task-completed": {
			Name:   "respond-activity-task-completed",
			Fields: fields_respond_activity_task_completed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RespondActivityTaskCompletedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_respond_activity_task_completed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RespondActivityTaskCompleted(ctx, input)
			},
		},
		"respond-activity-task-failed": {
			Name:   "respond-activity-task-failed",
			Fields: fields_respond_activity_task_failed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RespondActivityTaskFailedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_respond_activity_task_failed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RespondActivityTaskFailed(ctx, input)
			},
		},
		"respond-decision-task-completed": {
			Name:   "respond-decision-task-completed",
			Fields: fields_respond_decision_task_completed,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RespondDecisionTaskCompletedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_respond_decision_task_completed, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RespondDecisionTaskCompleted(ctx, input)
			},
		},
		"signal-workflow-execution": {
			Name:   "signal-workflow-execution",
			Fields: fields_signal_workflow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SignalWorkflowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_signal_workflow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SignalWorkflowExecution(ctx, input)
			},
		},
		"start-workflow-execution": {
			Name:   "start-workflow-execution",
			Fields: fields_start_workflow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkflowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workflow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkflowExecution(ctx, input)
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
		"terminate-workflow-execution": {
			Name:   "terminate-workflow-execution",
			Fields: fields_terminate_workflow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateWorkflowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_workflow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateWorkflowExecution(ctx, input)
			},
		},
		"undeprecate-activity-type": {
			Name:   "undeprecate-activity-type",
			Fields: fields_undeprecate_activity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UndeprecateActivityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_undeprecate_activity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UndeprecateActivityType(ctx, input)
			},
		},
		"undeprecate-domain": {
			Name:   "undeprecate-domain",
			Fields: fields_undeprecate_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UndeprecateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_undeprecate_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UndeprecateDomain(ctx, input)
			},
		},
		"undeprecate-workflow-type": {
			Name:   "undeprecate-workflow-type",
			Fields: fields_undeprecate_workflow_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UndeprecateWorkflowTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_undeprecate_workflow_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UndeprecateWorkflowType(ctx, input)
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
	}
	if err := leanruntime.Execute("swf", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
