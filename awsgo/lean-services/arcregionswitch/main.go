package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/arcregionswitch"
)

var fields_approve_plan_execution_step = []leanruntime.Field{
	{Name: "Approval", Flag: "approval", Type: "types.Approval", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
	{Name: "StepName", Flag: "step-name", Type: "*string", Required: true},
}

var fields_cancel_plan_execution = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
}

var fields_create_plan = []leanruntime.Field{
	{Name: "AssociatedAlarms", Flag: "associated-alarms", Type: "map[string]types.AssociatedAlarm", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PrimaryRegion", Flag: "primary-region", Type: "*string", Required: false},
	{Name: "RecoveryApproach", Flag: "recovery-approach", Type: "types.RecoveryApproach", Required: true},
	{Name: "RecoveryTimeObjectiveMinutes", Flag: "recovery-time-objective-minutes", Type: "*int32", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
	{Name: "ReportConfiguration", Flag: "report-configuration", Type: "*types.ReportConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Triggers", Flag: "triggers", Type: "[]types.Trigger", Required: false},
	{Name: "Workflows", Flag: "workflows", Type: "[]types.Workflow", Required: true},
}

var fields_delete_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_plan_evaluation_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
}

var fields_get_plan_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
}

var fields_get_plan_in_region = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_plan_execution_events = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
}

var fields_list_plan_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.ExecutionState", Required: false},
}

var fields_list_plans = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_plans_in_region = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_route53_health_checks = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecordName", Flag: "record-name", Type: "*string", Required: false},
}

var fields_list_route53_health_checks_in_region = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecordName", Flag: "record-name", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_start_plan_execution = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.ExecutionAction", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "LatestVersion", Flag: "latest-version", Type: "*string", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.ExecutionMode", Required: false},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
	{Name: "RecoveryExecutionId", Flag: "recovery-execution-id", Type: "*string", Required: false},
	{Name: "TargetRegion", Flag: "target-region", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ResourceTagKeys", Flag: "resource-tag-keys", Type: "[]string", Required: true},
}

var fields_update_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "AssociatedAlarms", Flag: "associated-alarms", Type: "map[string]types.AssociatedAlarm", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: true},
	{Name: "RecoveryTimeObjectiveMinutes", Flag: "recovery-time-objective-minutes", Type: "*int32", Required: false},
	{Name: "ReportConfiguration", Flag: "report-configuration", Type: "*types.ReportConfiguration", Required: false},
	{Name: "Triggers", Flag: "triggers", Type: "[]types.Trigger", Required: false},
	{Name: "Workflows", Flag: "workflows", Type: "[]types.Workflow", Required: true},
}

var fields_update_plan_execution = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.UpdatePlanExecutionAction", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
}

var fields_update_plan_execution_step = []leanruntime.Field{
	{Name: "ActionToTake", Flag: "action-to-take", Type: "types.UpdatePlanExecutionStepAction", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: true},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "PlanArn", Flag: "plan-arn", Type: "*string", Required: true},
	{Name: "StepName", Flag: "step-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"approve-plan-execution-step": {
			Name:   "approve-plan-execution-step",
			Fields: fields_approve_plan_execution_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApprovePlanExecutionStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_approve_plan_execution_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApprovePlanExecutionStep(ctx, input)
			},
		},
		"cancel-plan-execution": {
			Name:   "cancel-plan-execution",
			Fields: fields_cancel_plan_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelPlanExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_plan_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelPlanExecution(ctx, input)
			},
		},
		"create-plan": {
			Name:   "create-plan",
			Fields: fields_create_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlan(ctx, input)
			},
		},
		"delete-plan": {
			Name:   "delete-plan",
			Fields: fields_delete_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlan(ctx, input)
			},
		},
		"get-plan": {
			Name:   "get-plan",
			Fields: fields_get_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlan(ctx, input)
			},
		},
		"get-plan-evaluation-status": {
			Name:   "get-plan-evaluation-status",
			Fields: fields_get_plan_evaluation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlanEvaluationStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_plan_evaluation_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPlanEvaluationStatus(ctx, input)
				}
				var results []*svc.GetPlanEvaluationStatusOutput
				p := svc.NewGetPlanEvaluationStatusPaginator(client, input)
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
		"get-plan-execution": {
			Name:   "get-plan-execution",
			Fields: fields_get_plan_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlanExecutionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_plan_execution, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPlanExecution(ctx, input)
				}
				var results []*svc.GetPlanExecutionOutput
				p := svc.NewGetPlanExecutionPaginator(client, input)
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
		"get-plan-in-region": {
			Name:   "get-plan-in-region",
			Fields: fields_get_plan_in_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlanInRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_plan_in_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlanInRegion(ctx, input)
			},
		},
		"list-plan-execution-events": {
			Name:   "list-plan-execution-events",
			Fields: fields_list_plan_execution_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlanExecutionEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plan_execution_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlanExecutionEvents(ctx, input)
				}
				var results []*svc.ListPlanExecutionEventsOutput
				p := svc.NewListPlanExecutionEventsPaginator(client, input)
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
		"list-plan-executions": {
			Name:   "list-plan-executions",
			Fields: fields_list_plan_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlanExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plan_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlanExecutions(ctx, input)
				}
				var results []*svc.ListPlanExecutionsOutput
				p := svc.NewListPlanExecutionsPaginator(client, input)
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
		"list-plans": {
			Name:   "list-plans",
			Fields: fields_list_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlans(ctx, input)
				}
				var results []*svc.ListPlansOutput
				p := svc.NewListPlansPaginator(client, input)
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
		"list-plans-in-region": {
			Name:   "list-plans-in-region",
			Fields: fields_list_plans_in_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlansInRegionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plans_in_region, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlansInRegion(ctx, input)
				}
				var results []*svc.ListPlansInRegionOutput
				p := svc.NewListPlansInRegionPaginator(client, input)
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
		"list-route53-health-checks": {
			Name:   "list-route53-health-checks",
			Fields: fields_list_route53_health_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoute53HealthChecksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_route53_health_checks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoute53HealthChecks(ctx, input)
				}
				var results []*svc.ListRoute53HealthChecksOutput
				p := svc.NewListRoute53HealthChecksPaginator(client, input)
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
		"list-route53-health-checks-in-region": {
			Name:   "list-route53-health-checks-in-region",
			Fields: fields_list_route53_health_checks_in_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoute53HealthChecksInRegionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_route53_health_checks_in_region, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoute53HealthChecksInRegion(ctx, input)
				}
				var results []*svc.ListRoute53HealthChecksInRegionOutput
				p := svc.NewListRoute53HealthChecksInRegionPaginator(client, input)
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
		"start-plan-execution": {
			Name:   "start-plan-execution",
			Fields: fields_start_plan_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPlanExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_plan_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPlanExecution(ctx, input)
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
		"update-plan": {
			Name:   "update-plan",
			Fields: fields_update_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePlan(ctx, input)
			},
		},
		"update-plan-execution": {
			Name:   "update-plan-execution",
			Fields: fields_update_plan_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePlanExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_plan_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePlanExecution(ctx, input)
			},
		},
		"update-plan-execution-step": {
			Name:   "update-plan-execution-step",
			Fields: fields_update_plan_execution_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePlanExecutionStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_plan_execution_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePlanExecutionStep(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("arcregionswitch", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
