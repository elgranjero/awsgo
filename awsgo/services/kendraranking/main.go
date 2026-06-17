package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kendraranking/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-rescore-execution-plan", "delete-rescore-execution-plan", "describe-rescore-execution-plan", "list-rescore-execution-plans", "list-tags-for-resource", "rescore", "tag-resource", "untag-resource", "update-rescore-execution-plan"},
		OperationSet: map[string]bool{"create-rescore-execution-plan": true, "delete-rescore-execution-plan": true, "describe-rescore-execution-plan": true, "list-rescore-execution-plans": true, "list-tags-for-resource": true, "rescore": true, "tag-resource": true, "untag-resource": true, "update-rescore-execution-plan": true},
		OperationInputs: map[string][]string{
			"create-rescore-execution-plan":   {"CapacityUnits", "ClientToken", "Description", "Name", "Tags"},
			"delete-rescore-execution-plan":   {"Id"},
			"describe-rescore-execution-plan": {"Id"},
			"list-rescore-execution-plans":    {"MaxResults", "NextToken"},
			"list-tags-for-resource":          {"ResourceARN"},
			"rescore":                         {"Documents", "RescoreExecutionPlanId", "SearchQuery"},
			"tag-resource":                    {"ResourceARN", "Tags"},
			"untag-resource":                  {"ResourceARN", "TagKeys"},
			"update-rescore-execution-plan":   {"CapacityUnits", "Description", "Id", "Name"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-rescore-execution-plan":   {"CapacityUnits": "*types.CapacityUnitsConfiguration", "ClientToken": "*string", "Description": "*string", "Name": "*string", "Tags": "[]types.Tag"},
			"delete-rescore-execution-plan":   {"Id": "*string"},
			"describe-rescore-execution-plan": {"Id": "*string"},
			"list-rescore-execution-plans":    {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":          {"ResourceARN": "*string"},
			"rescore":                         {"Documents": "[]types.Document", "RescoreExecutionPlanId": "*string", "SearchQuery": "*string"},
			"tag-resource":                    {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                  {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-rescore-execution-plan":   {"CapacityUnits": "*types.CapacityUnitsConfiguration", "Description": "*string", "Id": "*string", "Name": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-rescore-execution-plan":   {"Name"},
			"delete-rescore-execution-plan":   {"Id"},
			"describe-rescore-execution-plan": {"Id"},
			"list-rescore-execution-plans":    {},
			"list-tags-for-resource":          {"ResourceARN"},
			"rescore":                         {"Documents", "RescoreExecutionPlanId", "SearchQuery"},
			"tag-resource":                    {"ResourceARN", "Tags"},
			"untag-resource":                  {"ResourceARN", "TagKeys"},
			"update-rescore-execution-plan":   {"Id"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kendraranking", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
