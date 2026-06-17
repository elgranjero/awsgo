package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/resourcegroupstaggingapi/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-report-creation", "get-compliance-summary", "get-resources", "get-tag-keys", "get-tag-values", "list-required-tags", "start-report-creation", "tag-resources", "untag-resources"},
		OperationSet: map[string]bool{"describe-report-creation": true, "get-compliance-summary": true, "get-resources": true, "get-tag-keys": true, "get-tag-values": true, "list-required-tags": true, "start-report-creation": true, "tag-resources": true, "untag-resources": true},
		OperationInputs: map[string][]string{
			"describe-report-creation": {},
			"get-compliance-summary":   {"GroupBy", "MaxResults", "PaginationToken", "RegionFilters", "ResourceTypeFilters", "TagKeyFilters", "TargetIdFilters"},
			"get-resources":            {"ExcludeCompliantResources", "IncludeComplianceDetails", "PaginationToken", "ResourceARNList", "ResourceTypeFilters", "ResourcesPerPage", "TagFilters", "TagsPerPage"},
			"get-tag-keys":             {"PaginationToken"},
			"get-tag-values":           {"Key", "PaginationToken"},
			"list-required-tags":       {"MaxResults", "NextToken"},
			"start-report-creation":    {"S3Bucket"},
			"tag-resources":            {"ResourceARNList", "Tags"},
			"untag-resources":          {"ResourceARNList", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-report-creation": {},
			"get-compliance-summary":   {"GroupBy": "[]types.GroupByAttribute", "MaxResults": "*int32", "PaginationToken": "*string", "RegionFilters": "[]string", "ResourceTypeFilters": "[]string", "TagKeyFilters": "[]string", "TargetIdFilters": "[]string"},
			"get-resources":            {"ExcludeCompliantResources": "*bool", "IncludeComplianceDetails": "*bool", "PaginationToken": "*string", "ResourceARNList": "[]string", "ResourceTypeFilters": "[]string", "ResourcesPerPage": "*int32", "TagFilters": "[]types.TagFilter", "TagsPerPage": "*int32"},
			"get-tag-keys":             {"PaginationToken": "*string"},
			"get-tag-values":           {"Key": "*string", "PaginationToken": "*string"},
			"list-required-tags":       {"MaxResults": "*int32", "NextToken": "*string"},
			"start-report-creation":    {"S3Bucket": "*string"},
			"tag-resources":            {"ResourceARNList": "[]string", "Tags": "map[string]string"},
			"untag-resources":          {"ResourceARNList": "[]string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"describe-report-creation": {},
			"get-compliance-summary":   {},
			"get-resources":            {},
			"get-tag-keys":             {},
			"get-tag-values":           {"Key"},
			"list-required-tags":       {},
			"start-report-creation":    {"S3Bucket"},
			"tag-resources":            {"ResourceARNList", "Tags"},
			"untag-resources":          {"ResourceARNList", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("resourcegroupstaggingapi", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
