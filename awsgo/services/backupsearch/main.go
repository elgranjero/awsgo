package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/backupsearch/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-search-job", "get-search-result-export-job", "list-search-job-backups", "list-search-job-results", "list-search-jobs", "list-search-result-export-jobs", "list-tags-for-resource", "start-search-job", "start-search-result-export-job", "stop-search-job", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"get-search-job": true, "get-search-result-export-job": true, "list-search-job-backups": true, "list-search-job-results": true, "list-search-jobs": true, "list-search-result-export-jobs": true, "list-tags-for-resource": true, "start-search-job": true, "start-search-result-export-job": true, "stop-search-job": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"get-search-job":                 {"SearchJobIdentifier"},
			"get-search-result-export-job":   {"ExportJobIdentifier"},
			"list-search-job-backups":        {"MaxResults", "NextToken", "SearchJobIdentifier"},
			"list-search-job-results":        {"MaxResults", "NextToken", "SearchJobIdentifier"},
			"list-search-jobs":               {"ByStatus", "MaxResults", "NextToken"},
			"list-search-result-export-jobs": {"MaxResults", "NextToken", "SearchJobIdentifier", "Status"},
			"list-tags-for-resource":         {"ResourceArn"},
			"start-search-job":               {"ClientToken", "EncryptionKeyArn", "ItemFilters", "Name", "SearchScope", "Tags"},
			"start-search-result-export-job": {"ClientToken", "ExportSpecification", "RoleArn", "SearchJobIdentifier", "Tags"},
			"stop-search-job":                {"SearchJobIdentifier"},
			"tag-resource":                   {"ResourceArn", "Tags"},
			"untag-resource":                 {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-search-job":                 {"SearchJobIdentifier": "*string"},
			"get-search-result-export-job":   {"ExportJobIdentifier": "*string"},
			"list-search-job-backups":        {"MaxResults": "*int32", "NextToken": "*string", "SearchJobIdentifier": "*string"},
			"list-search-job-results":        {"MaxResults": "*int32", "NextToken": "*string", "SearchJobIdentifier": "*string"},
			"list-search-jobs":               {"ByStatus": "types.SearchJobState", "MaxResults": "*int32", "NextToken": "*string"},
			"list-search-result-export-jobs": {"MaxResults": "*int32", "NextToken": "*string", "SearchJobIdentifier": "*string", "Status": "types.ExportJobStatus"},
			"list-tags-for-resource":         {"ResourceArn": "*string"},
			"start-search-job":               {"ClientToken": "*string", "EncryptionKeyArn": "*string", "ItemFilters": "*types.ItemFilters", "Name": "*string", "SearchScope": "*types.SearchScope", "Tags": "map[string]*string"},
			"start-search-result-export-job": {"ClientToken": "*string", "ExportSpecification": "types.ExportSpecification", "RoleArn": "*string", "SearchJobIdentifier": "*string", "Tags": "map[string]*string"},
			"stop-search-job":                {"SearchJobIdentifier": "*string"},
			"tag-resource":                   {"ResourceArn": "*string", "Tags": "map[string]*string"},
			"untag-resource":                 {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"get-search-job":                 {"SearchJobIdentifier"},
			"get-search-result-export-job":   {"ExportJobIdentifier"},
			"list-search-job-backups":        {"SearchJobIdentifier"},
			"list-search-job-results":        {"SearchJobIdentifier"},
			"list-search-jobs":               {},
			"list-search-result-export-jobs": {},
			"list-tags-for-resource":         {"ResourceArn"},
			"start-search-job":               {"SearchScope"},
			"start-search-result-export-job": {"ExportSpecification", "SearchJobIdentifier"},
			"stop-search-job":                {"SearchJobIdentifier"},
			"tag-resource":                   {"ResourceArn", "Tags"},
			"untag-resource":                 {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("backupsearch", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
