package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/scheduler/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-schedule", "create-schedule-group", "delete-schedule", "delete-schedule-group", "get-schedule", "get-schedule-group", "list-schedule-groups", "list-schedules", "list-tags-for-resource", "tag-resource", "untag-resource", "update-schedule"},
		OperationSet: map[string]bool{"create-schedule": true, "create-schedule-group": true, "delete-schedule": true, "delete-schedule-group": true, "get-schedule": true, "get-schedule-group": true, "list-schedule-groups": true, "list-schedules": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-schedule": true},
		OperationInputs: map[string][]string{
			"create-schedule":        {"ActionAfterCompletion", "ClientToken", "Description", "EndDate", "FlexibleTimeWindow", "GroupName", "KmsKeyArn", "Name", "ScheduleExpression", "ScheduleExpressionTimezone", "StartDate", "State", "Target"},
			"create-schedule-group":  {"ClientToken", "Name", "Tags"},
			"delete-schedule":        {"ClientToken", "GroupName", "Name"},
			"delete-schedule-group":  {"ClientToken", "Name"},
			"get-schedule":           {"GroupName", "Name"},
			"get-schedule-group":     {"Name"},
			"list-schedule-groups":   {"MaxResults", "NamePrefix", "NextToken"},
			"list-schedules":         {"GroupName", "MaxResults", "NamePrefix", "NextToken", "State"},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-schedule":        {"ActionAfterCompletion", "ClientToken", "Description", "EndDate", "FlexibleTimeWindow", "GroupName", "KmsKeyArn", "Name", "ScheduleExpression", "ScheduleExpressionTimezone", "StartDate", "State", "Target"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-schedule":        {"ActionAfterCompletion": "types.ActionAfterCompletion", "ClientToken": "*string", "Description": "*string", "EndDate": "*time.Time", "FlexibleTimeWindow": "*types.FlexibleTimeWindow", "GroupName": "*string", "KmsKeyArn": "*string", "Name": "*string", "ScheduleExpression": "*string", "ScheduleExpressionTimezone": "*string", "StartDate": "*time.Time", "State": "types.ScheduleState", "Target": "*types.Target"},
			"create-schedule-group":  {"ClientToken": "*string", "Name": "*string", "Tags": "[]types.Tag"},
			"delete-schedule":        {"ClientToken": "*string", "GroupName": "*string", "Name": "*string"},
			"delete-schedule-group":  {"ClientToken": "*string", "Name": "*string"},
			"get-schedule":           {"GroupName": "*string", "Name": "*string"},
			"get-schedule-group":     {"Name": "*string"},
			"list-schedule-groups":   {"MaxResults": "*int32", "NamePrefix": "*string", "NextToken": "*string"},
			"list-schedules":         {"GroupName": "*string", "MaxResults": "*int32", "NamePrefix": "*string", "NextToken": "*string", "State": "types.ScheduleState"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-schedule":        {"ActionAfterCompletion": "types.ActionAfterCompletion", "ClientToken": "*string", "Description": "*string", "EndDate": "*time.Time", "FlexibleTimeWindow": "*types.FlexibleTimeWindow", "GroupName": "*string", "KmsKeyArn": "*string", "Name": "*string", "ScheduleExpression": "*string", "ScheduleExpressionTimezone": "*string", "StartDate": "*time.Time", "State": "types.ScheduleState", "Target": "*types.Target"},
		},
		OperationInputRequired: map[string][]string{
			"create-schedule":        {"FlexibleTimeWindow", "Name", "ScheduleExpression", "Target"},
			"create-schedule-group":  {"Name"},
			"delete-schedule":        {"Name"},
			"delete-schedule-group":  {"Name"},
			"get-schedule":           {"Name"},
			"get-schedule-group":     {"Name"},
			"list-schedule-groups":   {},
			"list-schedules":         {},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-schedule":        {"FlexibleTimeWindow", "Name", "ScheduleExpression", "Target"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("scheduler", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
