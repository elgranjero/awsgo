package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/artifact/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-account-settings", "get-report", "get-report-metadata", "get-term-for-report", "list-customer-agreements", "list-report-versions", "list-reports", "put-account-settings"},
		OperationSet: map[string]bool{"get-account-settings": true, "get-report": true, "get-report-metadata": true, "get-term-for-report": true, "list-customer-agreements": true, "list-report-versions": true, "list-reports": true, "put-account-settings": true},
		OperationInputs: map[string][]string{
			"get-account-settings":     {},
			"get-report":               {"ReportId", "ReportVersion", "TermToken"},
			"get-report-metadata":      {"ReportId", "ReportVersion"},
			"get-term-for-report":      {"ReportId", "ReportVersion"},
			"list-customer-agreements": {"MaxResults", "NextToken"},
			"list-report-versions":     {"MaxResults", "NextToken", "ReportId"},
			"list-reports":             {"MaxResults", "NextToken"},
			"put-account-settings":     {"NotificationSubscriptionStatus"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-account-settings":     {},
			"get-report":               {"ReportId": "*string", "ReportVersion": "*int64", "TermToken": "*string"},
			"get-report-metadata":      {"ReportId": "*string", "ReportVersion": "*int64"},
			"get-term-for-report":      {"ReportId": "*string", "ReportVersion": "*int64"},
			"list-customer-agreements": {"MaxResults": "*int32", "NextToken": "*string"},
			"list-report-versions":     {"MaxResults": "*int32", "NextToken": "*string", "ReportId": "*string"},
			"list-reports":             {"MaxResults": "*int32", "NextToken": "*string"},
			"put-account-settings":     {"NotificationSubscriptionStatus": "types.NotificationSubscriptionStatus"},
		},
		OperationInputRequired: map[string][]string{
			"get-account-settings":     {},
			"get-report":               {"ReportId", "TermToken"},
			"get-report-metadata":      {"ReportId"},
			"get-term-for-report":      {"ReportId"},
			"list-customer-agreements": {},
			"list-report-versions":     {"ReportId"},
			"list-reports":             {},
			"put-account-settings":     {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("artifact", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
