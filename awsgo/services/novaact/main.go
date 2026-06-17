package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/novaact/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-act", "create-session", "create-workflow-definition", "create-workflow-run", "delete-workflow-definition", "delete-workflow-run", "get-workflow-definition", "get-workflow-run", "invoke-act-step", "list-acts", "list-models", "list-sessions", "list-workflow-definitions", "list-workflow-runs", "update-act", "update-workflow-run"},
		OperationSet: map[string]bool{"create-act": true, "create-session": true, "create-workflow-definition": true, "create-workflow-run": true, "delete-workflow-definition": true, "delete-workflow-run": true, "get-workflow-definition": true, "get-workflow-run": true, "invoke-act-step": true, "list-acts": true, "list-models": true, "list-sessions": true, "list-workflow-definitions": true, "list-workflow-runs": true, "update-act": true, "update-workflow-run": true},
		OperationInputs: map[string][]string{
			"create-act":                 {"ClientToken", "SessionId", "Task", "ToolSpecs", "WorkflowDefinitionName", "WorkflowRunId"},
			"create-session":             {"ClientToken", "WorkflowDefinitionName", "WorkflowRunId"},
			"create-workflow-definition": {"ClientToken", "Description", "ExportConfig", "Name"},
			"create-workflow-run":        {"ClientInfo", "ClientToken", "LogGroupName", "ModelId", "WorkflowDefinitionName"},
			"delete-workflow-definition": {"WorkflowDefinitionName"},
			"delete-workflow-run":        {"WorkflowDefinitionName", "WorkflowRunId"},
			"get-workflow-definition":    {"WorkflowDefinitionName"},
			"get-workflow-run":           {"WorkflowDefinitionName", "WorkflowRunId"},
			"invoke-act-step":            {"ActId", "CallResults", "PreviousStepId", "SessionId", "WorkflowDefinitionName", "WorkflowRunId"},
			"list-acts":                  {"MaxResults", "NextToken", "SessionId", "SortOrder", "WorkflowDefinitionName", "WorkflowRunId"},
			"list-models":                {"ClientCompatibilityVersion"},
			"list-sessions":              {"MaxResults", "NextToken", "SortOrder", "WorkflowDefinitionName", "WorkflowRunId"},
			"list-workflow-definitions":  {"MaxResults", "NextToken", "SortOrder"},
			"list-workflow-runs":         {"MaxResults", "NextToken", "SortOrder", "WorkflowDefinitionName"},
			"update-act":                 {"ActId", "Error", "SessionId", "Status", "WorkflowDefinitionName", "WorkflowRunId"},
			"update-workflow-run":        {"Status", "WorkflowDefinitionName", "WorkflowRunId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-act":                 {"ClientToken": "*string", "SessionId": "*string", "Task": "*string", "ToolSpecs": "[]types.ToolSpec", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"create-session":             {"ClientToken": "*string", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"create-workflow-definition": {"ClientToken": "*string", "Description": "*string", "ExportConfig": "*types.WorkflowExportConfig", "Name": "*string"},
			"create-workflow-run":        {"ClientInfo": "*types.ClientInfo", "ClientToken": "*string", "LogGroupName": "*string", "ModelId": "*string", "WorkflowDefinitionName": "*string"},
			"delete-workflow-definition": {"WorkflowDefinitionName": "*string"},
			"delete-workflow-run":        {"WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"get-workflow-definition":    {"WorkflowDefinitionName": "*string"},
			"get-workflow-run":           {"WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"invoke-act-step":            {"ActId": "*string", "CallResults": "[]types.CallResult", "PreviousStepId": "*string", "SessionId": "*string", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"list-acts":                  {"MaxResults": "*int32", "NextToken": "*string", "SessionId": "*string", "SortOrder": "types.SortOrder", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"list-models":                {"ClientCompatibilityVersion": "*int32"},
			"list-sessions":              {"MaxResults": "*int32", "NextToken": "*string", "SortOrder": "types.SortOrder", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"list-workflow-definitions":  {"MaxResults": "*int32", "NextToken": "*string", "SortOrder": "types.SortOrder"},
			"list-workflow-runs":         {"MaxResults": "*int32", "NextToken": "*string", "SortOrder": "types.SortOrder", "WorkflowDefinitionName": "*string"},
			"update-act":                 {"ActId": "*string", "Error": "*types.ActError", "SessionId": "*string", "Status": "types.ActStatus", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
			"update-workflow-run":        {"Status": "types.WorkflowRunStatus", "WorkflowDefinitionName": "*string", "WorkflowRunId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-act":                 {"SessionId", "Task", "WorkflowDefinitionName", "WorkflowRunId"},
			"create-session":             {"WorkflowDefinitionName", "WorkflowRunId"},
			"create-workflow-definition": {"Name"},
			"create-workflow-run":        {"ClientInfo", "ModelId", "WorkflowDefinitionName"},
			"delete-workflow-definition": {"WorkflowDefinitionName"},
			"delete-workflow-run":        {"WorkflowDefinitionName", "WorkflowRunId"},
			"get-workflow-definition":    {"WorkflowDefinitionName"},
			"get-workflow-run":           {"WorkflowDefinitionName", "WorkflowRunId"},
			"invoke-act-step":            {"ActId", "CallResults", "SessionId", "WorkflowDefinitionName", "WorkflowRunId"},
			"list-acts":                  {"WorkflowDefinitionName"},
			"list-models":                {"ClientCompatibilityVersion"},
			"list-sessions":              {"WorkflowDefinitionName", "WorkflowRunId"},
			"list-workflow-definitions":  {},
			"list-workflow-runs":         {"WorkflowDefinitionName"},
			"update-act":                 {"ActId", "SessionId", "Status", "WorkflowDefinitionName", "WorkflowRunId"},
			"update-workflow-run":        {"Status", "WorkflowDefinitionName", "WorkflowRunId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("novaact", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
