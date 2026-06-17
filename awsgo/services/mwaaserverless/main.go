package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mwaaserverless/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-workflow", "delete-workflow", "get-task-instance", "get-workflow", "get-workflow-run", "list-tags-for-resource", "list-task-instances", "list-workflow-runs", "list-workflow-versions", "list-workflows", "start-workflow-run", "stop-workflow-run", "tag-resource", "untag-resource", "update-workflow"},
		OperationSet: map[string]bool{"create-workflow": true, "delete-workflow": true, "get-task-instance": true, "get-workflow": true, "get-workflow-run": true, "list-tags-for-resource": true, "list-task-instances": true, "list-workflow-runs": true, "list-workflow-versions": true, "list-workflows": true, "start-workflow-run": true, "stop-workflow-run": true, "tag-resource": true, "untag-resource": true, "update-workflow": true},
		OperationInputs: map[string][]string{
			"create-workflow":        {"ClientToken", "DefinitionS3Location", "Description", "EncryptionConfiguration", "EngineVersion", "LoggingConfiguration", "Name", "NetworkConfiguration", "RoleArn", "Tags", "TriggerMode"},
			"delete-workflow":        {"WorkflowArn", "WorkflowVersion"},
			"get-task-instance":      {"RunId", "TaskInstanceId", "WorkflowArn"},
			"get-workflow":           {"WorkflowArn", "WorkflowVersion"},
			"get-workflow-run":       {"RunId", "WorkflowArn"},
			"list-tags-for-resource": {"ResourceArn"},
			"list-task-instances":    {"MaxResults", "NextToken", "RunId", "WorkflowArn"},
			"list-workflow-runs":     {"MaxResults", "NextToken", "WorkflowArn", "WorkflowVersion"},
			"list-workflow-versions": {"MaxResults", "NextToken", "WorkflowArn"},
			"list-workflows":         {"MaxResults", "NextToken"},
			"start-workflow-run":     {"ClientToken", "OverrideParameters", "WorkflowArn", "WorkflowVersion"},
			"stop-workflow-run":      {"RunId", "WorkflowArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-workflow":        {"DefinitionS3Location", "Description", "EngineVersion", "LoggingConfiguration", "NetworkConfiguration", "RoleArn", "TriggerMode", "WorkflowArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-workflow":        {"ClientToken": "*string", "DefinitionS3Location": "*types.DefinitionS3Location", "Description": "*string", "EncryptionConfiguration": "*types.EncryptionConfiguration", "EngineVersion": "types.EngineVersion", "LoggingConfiguration": "*types.LoggingConfiguration", "Name": "*string", "NetworkConfiguration": "*types.NetworkConfiguration", "RoleArn": "*string", "Tags": "map[string]string", "TriggerMode": "*string"},
			"delete-workflow":        {"WorkflowArn": "*string", "WorkflowVersion": "*string"},
			"get-task-instance":      {"RunId": "*string", "TaskInstanceId": "*string", "WorkflowArn": "*string"},
			"get-workflow":           {"WorkflowArn": "*string", "WorkflowVersion": "*string"},
			"get-workflow-run":       {"RunId": "*string", "WorkflowArn": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"list-task-instances":    {"MaxResults": "*int32", "NextToken": "*string", "RunId": "*string", "WorkflowArn": "*string"},
			"list-workflow-runs":     {"MaxResults": "*int32", "NextToken": "*string", "WorkflowArn": "*string", "WorkflowVersion": "*string"},
			"list-workflow-versions": {"MaxResults": "*int32", "NextToken": "*string", "WorkflowArn": "*string"},
			"list-workflows":         {"MaxResults": "*int32", "NextToken": "*string"},
			"start-workflow-run":     {"ClientToken": "*string", "OverrideParameters": "map[string]document.Interface", "WorkflowArn": "*string", "WorkflowVersion": "*string"},
			"stop-workflow-run":      {"RunId": "*string", "WorkflowArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-workflow":        {"DefinitionS3Location": "*types.DefinitionS3Location", "Description": "*string", "EngineVersion": "types.EngineVersion", "LoggingConfiguration": "*types.LoggingConfiguration", "NetworkConfiguration": "*types.NetworkConfiguration", "RoleArn": "*string", "TriggerMode": "*string", "WorkflowArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-workflow":        {"DefinitionS3Location", "Name", "RoleArn"},
			"delete-workflow":        {"WorkflowArn"},
			"get-task-instance":      {"RunId", "TaskInstanceId", "WorkflowArn"},
			"get-workflow":           {"WorkflowArn"},
			"get-workflow-run":       {"RunId", "WorkflowArn"},
			"list-tags-for-resource": {"ResourceArn"},
			"list-task-instances":    {"RunId", "WorkflowArn"},
			"list-workflow-runs":     {"WorkflowArn"},
			"list-workflow-versions": {"WorkflowArn"},
			"list-workflows":         {},
			"start-workflow-run":     {"WorkflowArn"},
			"stop-workflow-run":      {"RunId", "WorkflowArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-workflow":        {"DefinitionS3Location", "RoleArn", "WorkflowArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mwaaserverless", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
