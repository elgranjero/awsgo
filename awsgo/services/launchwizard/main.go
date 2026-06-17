package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/launchwizard/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-deployment", "delete-deployment", "get-deployment", "get-deployment-pattern-version", "get-workload", "get-workload-deployment-pattern", "list-deployment-events", "list-deployment-pattern-versions", "list-deployments", "list-tags-for-resource", "list-workload-deployment-patterns", "list-workloads", "tag-resource", "untag-resource", "update-deployment"},
		OperationSet: map[string]bool{"create-deployment": true, "delete-deployment": true, "get-deployment": true, "get-deployment-pattern-version": true, "get-workload": true, "get-workload-deployment-pattern": true, "list-deployment-events": true, "list-deployment-pattern-versions": true, "list-deployments": true, "list-tags-for-resource": true, "list-workload-deployment-patterns": true, "list-workloads": true, "tag-resource": true, "untag-resource": true, "update-deployment": true},
		OperationInputs: map[string][]string{
			"create-deployment":                 {"DeploymentPatternName", "DryRun", "Name", "Specifications", "Tags", "WorkloadName"},
			"delete-deployment":                 {"DeploymentId"},
			"get-deployment":                    {"DeploymentId"},
			"get-deployment-pattern-version":    {"DeploymentPatternName", "DeploymentPatternVersionName", "WorkloadName"},
			"get-workload":                      {"WorkloadName"},
			"get-workload-deployment-pattern":   {"DeploymentPatternName", "WorkloadName"},
			"list-deployment-events":            {"DeploymentId", "MaxResults", "NextToken"},
			"list-deployment-pattern-versions":  {"DeploymentPatternName", "Filters", "MaxResults", "NextToken", "WorkloadName"},
			"list-deployments":                  {"Filters", "MaxResults", "NextToken"},
			"list-tags-for-resource":            {"ResourceArn"},
			"list-workload-deployment-patterns": {"MaxResults", "NextToken", "WorkloadName"},
			"list-workloads":                    {"MaxResults", "NextToken"},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-deployment":                 {"DeploymentId", "DeploymentPatternVersionName", "DryRun", "Force", "Specifications", "WorkloadVersionName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-deployment":                 {"DeploymentPatternName": "*string", "DryRun": "bool", "Name": "*string", "Specifications": "map[string]string", "Tags": "map[string]string", "WorkloadName": "*string"},
			"delete-deployment":                 {"DeploymentId": "*string"},
			"get-deployment":                    {"DeploymentId": "*string"},
			"get-deployment-pattern-version":    {"DeploymentPatternName": "*string", "DeploymentPatternVersionName": "*string", "WorkloadName": "*string"},
			"get-workload":                      {"WorkloadName": "*string"},
			"get-workload-deployment-pattern":   {"DeploymentPatternName": "*string", "WorkloadName": "*string"},
			"list-deployment-events":            {"DeploymentId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-deployment-pattern-versions":  {"DeploymentPatternName": "*string", "Filters": "[]types.DeploymentPatternVersionFilter", "MaxResults": "*int32", "NextToken": "*string", "WorkloadName": "*string"},
			"list-deployments":                  {"Filters": "[]types.DeploymentFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":            {"ResourceArn": "*string"},
			"list-workload-deployment-patterns": {"MaxResults": "*int32", "NextToken": "*string", "WorkloadName": "*string"},
			"list-workloads":                    {"MaxResults": "*int32", "NextToken": "*string"},
			"tag-resource":                      {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                    {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-deployment":                 {"DeploymentId": "*string", "DeploymentPatternVersionName": "*string", "DryRun": "bool", "Force": "bool", "Specifications": "map[string]string", "WorkloadVersionName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-deployment":                 {"DeploymentPatternName", "Name", "Specifications", "WorkloadName"},
			"delete-deployment":                 {"DeploymentId"},
			"get-deployment":                    {"DeploymentId"},
			"get-deployment-pattern-version":    {"DeploymentPatternName", "DeploymentPatternVersionName", "WorkloadName"},
			"get-workload":                      {"WorkloadName"},
			"get-workload-deployment-pattern":   {"DeploymentPatternName", "WorkloadName"},
			"list-deployment-events":            {"DeploymentId"},
			"list-deployment-pattern-versions":  {"DeploymentPatternName", "WorkloadName"},
			"list-deployments":                  {},
			"list-tags-for-resource":            {"ResourceArn"},
			"list-workload-deployment-patterns": {"WorkloadName"},
			"list-workloads":                    {},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-deployment":                 {"DeploymentId", "Specifications"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("launchwizard", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
