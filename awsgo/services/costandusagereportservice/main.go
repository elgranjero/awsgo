package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/costandusagereportservice/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-report-definition", "describe-report-definitions", "list-tags-for-resource", "modify-report-definition", "put-report-definition", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"delete-report-definition": true, "describe-report-definitions": true, "list-tags-for-resource": true, "modify-report-definition": true, "put-report-definition": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"delete-report-definition":    {"ReportName"},
			"describe-report-definitions": {"MaxResults", "NextToken"},
			"list-tags-for-resource":      {"ReportName"},
			"modify-report-definition":    {"ReportDefinition", "ReportName"},
			"put-report-definition":       {"ReportDefinition", "Tags"},
			"tag-resource":                {"ReportName", "Tags"},
			"untag-resource":              {"ReportName", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-report-definition":    {"ReportName": "*string"},
			"describe-report-definitions": {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":      {"ReportName": "*string"},
			"modify-report-definition":    {"ReportDefinition": "*types.ReportDefinition", "ReportName": "*string"},
			"put-report-definition":       {"ReportDefinition": "*types.ReportDefinition", "Tags": "[]types.Tag"},
			"tag-resource":                {"ReportName": "*string", "Tags": "[]types.Tag"},
			"untag-resource":              {"ReportName": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-report-definition":    {"ReportName"},
			"describe-report-definitions": {},
			"list-tags-for-resource":      {"ReportName"},
			"modify-report-definition":    {"ReportDefinition", "ReportName"},
			"put-report-definition":       {"ReportDefinition"},
			"tag-resource":                {"ReportName", "Tags"},
			"untag-resource":              {"ReportName", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("costandusagereportservice", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
