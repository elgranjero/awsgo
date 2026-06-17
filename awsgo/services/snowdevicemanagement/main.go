package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/snowdevicemanagement/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-task", "create-task", "describe-device", "describe-device-ec2-instances", "describe-execution", "describe-task", "list-device-resources", "list-devices", "list-executions", "list-tags-for-resource", "list-tasks", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"cancel-task": true, "create-task": true, "describe-device": true, "describe-device-ec2-instances": true, "describe-execution": true, "describe-task": true, "list-device-resources": true, "list-devices": true, "list-executions": true, "list-tags-for-resource": true, "list-tasks": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"cancel-task":                   {"TaskId"},
			"create-task":                   {"ClientToken", "Command", "Description", "Tags", "Targets"},
			"describe-device":               {"ManagedDeviceId"},
			"describe-device-ec2-instances": {"InstanceIds", "ManagedDeviceId"},
			"describe-execution":            {"ManagedDeviceId", "TaskId"},
			"describe-task":                 {"TaskId"},
			"list-device-resources":         {"ManagedDeviceId", "MaxResults", "NextToken", "Type"},
			"list-devices":                  {"JobId", "MaxResults", "NextToken"},
			"list-executions":               {"MaxResults", "NextToken", "State", "TaskId"},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-tasks":                    {"MaxResults", "NextToken", "State"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-task":                   {"TaskId": "*string"},
			"create-task":                   {"ClientToken": "*string", "Command": "types.Command", "Description": "*string", "Tags": "map[string]string", "Targets": "[]string"},
			"describe-device":               {"ManagedDeviceId": "*string"},
			"describe-device-ec2-instances": {"InstanceIds": "[]string", "ManagedDeviceId": "*string"},
			"describe-execution":            {"ManagedDeviceId": "*string", "TaskId": "*string"},
			"describe-task":                 {"TaskId": "*string"},
			"list-device-resources":         {"ManagedDeviceId": "*string", "MaxResults": "*int32", "NextToken": "*string", "Type": "*string"},
			"list-devices":                  {"JobId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-executions":               {"MaxResults": "*int32", "NextToken": "*string", "State": "types.ExecutionState", "TaskId": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"list-tasks":                    {"MaxResults": "*int32", "NextToken": "*string", "State": "types.TaskState"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-task":                   {"TaskId"},
			"create-task":                   {"Command", "Targets"},
			"describe-device":               {"ManagedDeviceId"},
			"describe-device-ec2-instances": {"InstanceIds", "ManagedDeviceId"},
			"describe-execution":            {"ManagedDeviceId", "TaskId"},
			"describe-task":                 {"TaskId"},
			"list-device-resources":         {"ManagedDeviceId"},
			"list-devices":                  {},
			"list-executions":               {"TaskId"},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-tasks":                    {},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("snowdevicemanagement", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
