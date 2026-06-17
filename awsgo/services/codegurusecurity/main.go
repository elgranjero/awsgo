package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/codegurusecurity/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-findings", "create-scan", "create-upload-url", "get-account-configuration", "get-findings", "get-metrics-summary", "get-scan", "list-findings-metrics", "list-scans", "list-tags-for-resource", "tag-resource", "untag-resource", "update-account-configuration"},
		OperationSet: map[string]bool{"batch-get-findings": true, "create-scan": true, "create-upload-url": true, "get-account-configuration": true, "get-findings": true, "get-metrics-summary": true, "get-scan": true, "list-findings-metrics": true, "list-scans": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-account-configuration": true},
		OperationInputs: map[string][]string{
			"batch-get-findings":           {"FindingIdentifiers"},
			"create-scan":                  {"AnalysisType", "ClientToken", "ResourceId", "ScanName", "ScanType", "Tags"},
			"create-upload-url":            {"ScanName"},
			"get-account-configuration":    {},
			"get-findings":                 {"MaxResults", "NextToken", "ScanName", "Status"},
			"get-metrics-summary":          {"Date"},
			"get-scan":                     {"RunId", "ScanName"},
			"list-findings-metrics":        {"EndDate", "MaxResults", "NextToken", "StartDate"},
			"list-scans":                   {"MaxResults", "NextToken"},
			"list-tags-for-resource":       {"ResourceArn"},
			"tag-resource":                 {"ResourceArn", "Tags"},
			"untag-resource":               {"ResourceArn", "TagKeys"},
			"update-account-configuration": {"EncryptionConfig"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-findings":           {"FindingIdentifiers": "[]types.FindingIdentifier"},
			"create-scan":                  {"AnalysisType": "types.AnalysisType", "ClientToken": "*string", "ResourceId": "types.ResourceId", "ScanName": "*string", "ScanType": "types.ScanType", "Tags": "map[string]string"},
			"create-upload-url":            {"ScanName": "*string"},
			"get-account-configuration":    {},
			"get-findings":                 {"MaxResults": "*int32", "NextToken": "*string", "ScanName": "*string", "Status": "types.Status"},
			"get-metrics-summary":          {"Date": "*time.Time"},
			"get-scan":                     {"RunId": "*string", "ScanName": "*string"},
			"list-findings-metrics":        {"EndDate": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "StartDate": "*time.Time"},
			"list-scans":                   {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":       {"ResourceArn": "*string"},
			"tag-resource":                 {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":               {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-account-configuration": {"EncryptionConfig": "*types.EncryptionConfig"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-findings":           {"FindingIdentifiers"},
			"create-scan":                  {"ResourceId", "ScanName"},
			"create-upload-url":            {"ScanName"},
			"get-account-configuration":    {},
			"get-findings":                 {"ScanName"},
			"get-metrics-summary":          {"Date"},
			"get-scan":                     {"ScanName"},
			"list-findings-metrics":        {"EndDate", "StartDate"},
			"list-scans":                   {},
			"list-tags-for-resource":       {"ResourceArn"},
			"tag-resource":                 {"ResourceArn", "Tags"},
			"untag-resource":               {"ResourceArn", "TagKeys"},
			"update-account-configuration": {"EncryptionConfig"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("codegurusecurity", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
