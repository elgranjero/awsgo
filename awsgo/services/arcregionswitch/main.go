package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/arcregionswitch/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"approve-plan-execution-step", "cancel-plan-execution", "create-plan", "delete-plan", "get-plan", "get-plan-evaluation-status", "get-plan-execution", "get-plan-in-region", "list-plan-execution-events", "list-plan-executions", "list-plans", "list-plans-in-region", "list-route53-health-checks", "list-route53-health-checks-in-region", "list-tags-for-resource", "start-plan-execution", "tag-resource", "untag-resource", "update-plan", "update-plan-execution", "update-plan-execution-step"},
		OperationSet: map[string]bool{"approve-plan-execution-step": true, "cancel-plan-execution": true, "create-plan": true, "delete-plan": true, "get-plan": true, "get-plan-evaluation-status": true, "get-plan-execution": true, "get-plan-in-region": true, "list-plan-execution-events": true, "list-plan-executions": true, "list-plans": true, "list-plans-in-region": true, "list-route53-health-checks": true, "list-route53-health-checks-in-region": true, "list-tags-for-resource": true, "start-plan-execution": true, "tag-resource": true, "untag-resource": true, "update-plan": true, "update-plan-execution": true, "update-plan-execution-step": true},
		OperationInputs: map[string][]string{
			"approve-plan-execution-step":          {"Approval", "Comment", "ExecutionId", "PlanArn", "StepName"},
			"cancel-plan-execution":                {"Comment", "ExecutionId", "PlanArn"},
			"create-plan":                          {"AssociatedAlarms", "Description", "ExecutionRole", "Name", "PrimaryRegion", "RecoveryApproach", "RecoveryTimeObjectiveMinutes", "Regions", "ReportConfiguration", "Tags", "Triggers", "Workflows"},
			"delete-plan":                          {"Arn"},
			"get-plan":                             {"Arn"},
			"get-plan-evaluation-status":           {"MaxResults", "NextToken", "PlanArn"},
			"get-plan-execution":                   {"ExecutionId", "MaxResults", "NextToken", "PlanArn"},
			"get-plan-in-region":                   {"Arn"},
			"list-plan-execution-events":           {"ExecutionId", "MaxResults", "Name", "NextToken", "PlanArn"},
			"list-plan-executions":                 {"MaxResults", "NextToken", "PlanArn", "State"},
			"list-plans":                           {"MaxResults", "NextToken"},
			"list-plans-in-region":                 {"MaxResults", "NextToken"},
			"list-route53-health-checks":           {"Arn", "HostedZoneId", "MaxResults", "NextToken", "RecordName"},
			"list-route53-health-checks-in-region": {"Arn", "HostedZoneId", "MaxResults", "NextToken", "RecordName"},
			"list-tags-for-resource":               {"Arn"},
			"start-plan-execution":                 {"Action", "Comment", "LatestVersion", "Mode", "PlanArn", "RecoveryExecutionId", "TargetRegion"},
			"tag-resource":                         {"Arn", "Tags"},
			"untag-resource":                       {"Arn", "ResourceTagKeys"},
			"update-plan":                          {"Arn", "AssociatedAlarms", "Description", "ExecutionRole", "RecoveryTimeObjectiveMinutes", "ReportConfiguration", "Triggers", "Workflows"},
			"update-plan-execution":                {"Action", "Comment", "ExecutionId", "PlanArn"},
			"update-plan-execution-step":           {"ActionToTake", "Comment", "ExecutionId", "PlanArn", "StepName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"approve-plan-execution-step":          {"Approval": "types.Approval", "Comment": "*string", "ExecutionId": "*string", "PlanArn": "*string", "StepName": "*string"},
			"cancel-plan-execution":                {"Comment": "*string", "ExecutionId": "*string", "PlanArn": "*string"},
			"create-plan":                          {"AssociatedAlarms": "map[string]types.AssociatedAlarm", "Description": "*string", "ExecutionRole": "*string", "Name": "*string", "PrimaryRegion": "*string", "RecoveryApproach": "types.RecoveryApproach", "RecoveryTimeObjectiveMinutes": "*int32", "Regions": "[]string", "ReportConfiguration": "*types.ReportConfiguration", "Tags": "map[string]string", "Triggers": "[]types.Trigger", "Workflows": "[]types.Workflow"},
			"delete-plan":                          {"Arn": "*string"},
			"get-plan":                             {"Arn": "*string"},
			"get-plan-evaluation-status":           {"MaxResults": "*int32", "NextToken": "*string", "PlanArn": "*string"},
			"get-plan-execution":                   {"ExecutionId": "*string", "MaxResults": "*int32", "NextToken": "*string", "PlanArn": "*string"},
			"get-plan-in-region":                   {"Arn": "*string"},
			"list-plan-execution-events":           {"ExecutionId": "*string", "MaxResults": "*int32", "Name": "*string", "NextToken": "*string", "PlanArn": "*string"},
			"list-plan-executions":                 {"MaxResults": "*int32", "NextToken": "*string", "PlanArn": "*string", "State": "types.ExecutionState"},
			"list-plans":                           {"MaxResults": "*int32", "NextToken": "*string"},
			"list-plans-in-region":                 {"MaxResults": "*int32", "NextToken": "*string"},
			"list-route53-health-checks":           {"Arn": "*string", "HostedZoneId": "*string", "MaxResults": "*int32", "NextToken": "*string", "RecordName": "*string"},
			"list-route53-health-checks-in-region": {"Arn": "*string", "HostedZoneId": "*string", "MaxResults": "*int32", "NextToken": "*string", "RecordName": "*string"},
			"list-tags-for-resource":               {"Arn": "*string"},
			"start-plan-execution":                 {"Action": "types.ExecutionAction", "Comment": "*string", "LatestVersion": "*string", "Mode": "types.ExecutionMode", "PlanArn": "*string", "RecoveryExecutionId": "*string", "TargetRegion": "*string"},
			"tag-resource":                         {"Arn": "*string", "Tags": "map[string]string"},
			"untag-resource":                       {"Arn": "*string", "ResourceTagKeys": "[]string"},
			"update-plan":                          {"Arn": "*string", "AssociatedAlarms": "map[string]types.AssociatedAlarm", "Description": "*string", "ExecutionRole": "*string", "RecoveryTimeObjectiveMinutes": "*int32", "ReportConfiguration": "*types.ReportConfiguration", "Triggers": "[]types.Trigger", "Workflows": "[]types.Workflow"},
			"update-plan-execution":                {"Action": "types.UpdatePlanExecutionAction", "Comment": "*string", "ExecutionId": "*string", "PlanArn": "*string"},
			"update-plan-execution-step":           {"ActionToTake": "types.UpdatePlanExecutionStepAction", "Comment": "*string", "ExecutionId": "*string", "PlanArn": "*string", "StepName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"approve-plan-execution-step":          {"Approval", "ExecutionId", "PlanArn", "StepName"},
			"cancel-plan-execution":                {"ExecutionId", "PlanArn"},
			"create-plan":                          {"ExecutionRole", "Name", "RecoveryApproach", "Regions", "Workflows"},
			"delete-plan":                          {"Arn"},
			"get-plan":                             {"Arn"},
			"get-plan-evaluation-status":           {"PlanArn"},
			"get-plan-execution":                   {"ExecutionId", "PlanArn"},
			"get-plan-in-region":                   {"Arn"},
			"list-plan-execution-events":           {"ExecutionId", "PlanArn"},
			"list-plan-executions":                 {"PlanArn"},
			"list-plans":                           {},
			"list-plans-in-region":                 {},
			"list-route53-health-checks":           {"Arn"},
			"list-route53-health-checks-in-region": {"Arn"},
			"list-tags-for-resource":               {"Arn"},
			"start-plan-execution":                 {"Action", "PlanArn", "TargetRegion"},
			"tag-resource":                         {"Arn", "Tags"},
			"untag-resource":                       {"Arn", "ResourceTagKeys"},
			"update-plan":                          {"Arn", "ExecutionRole", "Workflows"},
			"update-plan-execution":                {"Action", "ExecutionId", "PlanArn"},
			"update-plan-execution-step":           {"ActionToTake", "Comment", "ExecutionId", "PlanArn", "StepName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("arcregionswitch", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
