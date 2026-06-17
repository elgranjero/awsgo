package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/datapipeline/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"activate-pipeline", "add-tags", "create-pipeline", "deactivate-pipeline", "delete-pipeline", "describe-objects", "describe-pipelines", "evaluate-expression", "get-pipeline-definition", "list-pipelines", "poll-for-task", "put-pipeline-definition", "query-objects", "remove-tags", "report-task-progress", "report-task-runner-heartbeat", "set-status", "set-task-status", "validate-pipeline-definition"},
		OperationSet: map[string]bool{"activate-pipeline": true, "add-tags": true, "create-pipeline": true, "deactivate-pipeline": true, "delete-pipeline": true, "describe-objects": true, "describe-pipelines": true, "evaluate-expression": true, "get-pipeline-definition": true, "list-pipelines": true, "poll-for-task": true, "put-pipeline-definition": true, "query-objects": true, "remove-tags": true, "report-task-progress": true, "report-task-runner-heartbeat": true, "set-status": true, "set-task-status": true, "validate-pipeline-definition": true},
		OperationInputs: map[string][]string{
			"activate-pipeline":            {"ParameterValues", "PipelineId", "StartTimestamp"},
			"add-tags":                     {"PipelineId", "Tags"},
			"create-pipeline":              {"Description", "Name", "Tags", "UniqueId"},
			"deactivate-pipeline":          {"CancelActive", "PipelineId"},
			"delete-pipeline":              {"PipelineId"},
			"describe-objects":             {"EvaluateExpressions", "Marker", "ObjectIds", "PipelineId"},
			"describe-pipelines":           {"PipelineIds"},
			"evaluate-expression":          {"Expression", "ObjectId", "PipelineId"},
			"get-pipeline-definition":      {"PipelineId", "Version"},
			"list-pipelines":               {"Marker"},
			"poll-for-task":                {"Hostname", "InstanceIdentity", "WorkerGroup"},
			"put-pipeline-definition":      {"ParameterObjects", "ParameterValues", "PipelineId", "PipelineObjects"},
			"query-objects":                {"Limit", "Marker", "PipelineId", "Query", "Sphere"},
			"remove-tags":                  {"PipelineId", "TagKeys"},
			"report-task-progress":         {"Fields", "TaskId"},
			"report-task-runner-heartbeat": {"Hostname", "TaskrunnerId", "WorkerGroup"},
			"set-status":                   {"ObjectIds", "PipelineId", "Status"},
			"set-task-status":              {"ErrorId", "ErrorMessage", "ErrorStackTrace", "TaskId", "TaskStatus"},
			"validate-pipeline-definition": {"ParameterObjects", "ParameterValues", "PipelineId", "PipelineObjects"},
		},
		OperationInputTypes: map[string]map[string]string{
			"activate-pipeline":            {"ParameterValues": "[]types.ParameterValue", "PipelineId": "*string", "StartTimestamp": "*time.Time"},
			"add-tags":                     {"PipelineId": "*string", "Tags": "[]types.Tag"},
			"create-pipeline":              {"Description": "*string", "Name": "*string", "Tags": "[]types.Tag", "UniqueId": "*string"},
			"deactivate-pipeline":          {"CancelActive": "*bool", "PipelineId": "*string"},
			"delete-pipeline":              {"PipelineId": "*string"},
			"describe-objects":             {"EvaluateExpressions": "bool", "Marker": "*string", "ObjectIds": "[]string", "PipelineId": "*string"},
			"describe-pipelines":           {"PipelineIds": "[]string"},
			"evaluate-expression":          {"Expression": "*string", "ObjectId": "*string", "PipelineId": "*string"},
			"get-pipeline-definition":      {"PipelineId": "*string", "Version": "*string"},
			"list-pipelines":               {"Marker": "*string"},
			"poll-for-task":                {"Hostname": "*string", "InstanceIdentity": "*types.InstanceIdentity", "WorkerGroup": "*string"},
			"put-pipeline-definition":      {"ParameterObjects": "[]types.ParameterObject", "ParameterValues": "[]types.ParameterValue", "PipelineId": "*string", "PipelineObjects": "[]types.PipelineObject"},
			"query-objects":                {"Limit": "*int32", "Marker": "*string", "PipelineId": "*string", "Query": "*types.Query", "Sphere": "*string"},
			"remove-tags":                  {"PipelineId": "*string", "TagKeys": "[]string"},
			"report-task-progress":         {"Fields": "[]types.Field", "TaskId": "*string"},
			"report-task-runner-heartbeat": {"Hostname": "*string", "TaskrunnerId": "*string", "WorkerGroup": "*string"},
			"set-status":                   {"ObjectIds": "[]string", "PipelineId": "*string", "Status": "*string"},
			"set-task-status":              {"ErrorId": "*string", "ErrorMessage": "*string", "ErrorStackTrace": "*string", "TaskId": "*string", "TaskStatus": "types.TaskStatus"},
			"validate-pipeline-definition": {"ParameterObjects": "[]types.ParameterObject", "ParameterValues": "[]types.ParameterValue", "PipelineId": "*string", "PipelineObjects": "[]types.PipelineObject"},
		},
		OperationInputRequired: map[string][]string{
			"activate-pipeline":            {"PipelineId"},
			"add-tags":                     {"PipelineId", "Tags"},
			"create-pipeline":              {"Name", "UniqueId"},
			"deactivate-pipeline":          {"PipelineId"},
			"delete-pipeline":              {"PipelineId"},
			"describe-objects":             {"ObjectIds", "PipelineId"},
			"describe-pipelines":           {"PipelineIds"},
			"evaluate-expression":          {"Expression", "ObjectId", "PipelineId"},
			"get-pipeline-definition":      {"PipelineId"},
			"list-pipelines":               {},
			"poll-for-task":                {"WorkerGroup"},
			"put-pipeline-definition":      {"PipelineId", "PipelineObjects"},
			"query-objects":                {"PipelineId", "Sphere"},
			"remove-tags":                  {"PipelineId", "TagKeys"},
			"report-task-progress":         {"TaskId"},
			"report-task-runner-heartbeat": {"TaskrunnerId"},
			"set-status":                   {"ObjectIds", "PipelineId", "Status"},
			"set-task-status":              {"TaskId", "TaskStatus"},
			"validate-pipeline-definition": {"PipelineId", "PipelineObjects"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("datapipeline", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
