package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/bcmdataexports/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-export", "delete-export", "get-execution", "get-export", "get-table", "list-executions", "list-exports", "list-tables", "list-tags-for-resource", "tag-resource", "untag-resource", "update-export"},
		OperationSet: map[string]bool{"create-export": true, "delete-export": true, "get-execution": true, "get-export": true, "get-table": true, "list-executions": true, "list-exports": true, "list-tables": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-export": true},
		OperationInputs: map[string][]string{
			"create-export":          {"Export", "ResourceTags"},
			"delete-export":          {"ExportArn"},
			"get-execution":          {"ExecutionId", "ExportArn"},
			"get-export":             {"ExportArn"},
			"get-table":              {"TableName", "TableProperties"},
			"list-executions":        {"ExportArn", "MaxResults", "NextToken"},
			"list-exports":           {"MaxResults", "NextToken"},
			"list-tables":            {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"MaxResults", "NextToken", "ResourceArn"},
			"tag-resource":           {"ResourceArn", "ResourceTags"},
			"untag-resource":         {"ResourceArn", "ResourceTagKeys"},
			"update-export":          {"Export", "ExportArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-export":          {"Export": "*types.Export", "ResourceTags": "[]types.ResourceTag"},
			"delete-export":          {"ExportArn": "*string"},
			"get-execution":          {"ExecutionId": "*string", "ExportArn": "*string"},
			"get-export":             {"ExportArn": "*string"},
			"get-table":              {"TableName": "*string", "TableProperties": "map[string]string"},
			"list-executions":        {"ExportArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-exports":           {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tables":            {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"MaxResults": "*int32", "NextToken": "*string", "ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "ResourceTags": "[]types.ResourceTag"},
			"untag-resource":         {"ResourceArn": "*string", "ResourceTagKeys": "[]string"},
			"update-export":          {"Export": "*types.Export", "ExportArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-export":          {"Export"},
			"delete-export":          {"ExportArn"},
			"get-execution":          {"ExecutionId", "ExportArn"},
			"get-export":             {"ExportArn"},
			"get-table":              {"TableName"},
			"list-executions":        {"ExportArn"},
			"list-exports":           {},
			"list-tables":            {},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "ResourceTags"},
			"untag-resource":         {"ResourceArn", "ResourceTagKeys"},
			"update-export":          {"Export", "ExportArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("bcmdataexports", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
