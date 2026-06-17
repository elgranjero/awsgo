package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/migrationhubstrategy/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-application-component-details", "get-application-component-strategies", "get-assessment", "get-import-file-task", "get-latest-assessment-id", "get-portfolio-preferences", "get-portfolio-summary", "get-recommendation-report-details", "get-server-details", "get-server-strategies", "list-analyzable-servers", "list-application-components", "list-collectors", "list-import-file-task", "list-servers", "put-portfolio-preferences", "start-assessment", "start-import-file-task", "start-recommendation-report-generation", "stop-assessment", "update-application-component-config", "update-server-config"},
		OperationSet: map[string]bool{"get-application-component-details": true, "get-application-component-strategies": true, "get-assessment": true, "get-import-file-task": true, "get-latest-assessment-id": true, "get-portfolio-preferences": true, "get-portfolio-summary": true, "get-recommendation-report-details": true, "get-server-details": true, "get-server-strategies": true, "list-analyzable-servers": true, "list-application-components": true, "list-collectors": true, "list-import-file-task": true, "list-servers": true, "put-portfolio-preferences": true, "start-assessment": true, "start-import-file-task": true, "start-recommendation-report-generation": true, "stop-assessment": true, "update-application-component-config": true, "update-server-config": true},
		OperationInputs: map[string][]string{
			"get-application-component-details":      {"ApplicationComponentId"},
			"get-application-component-strategies":   {"ApplicationComponentId"},
			"get-assessment":                         {"Id"},
			"get-import-file-task":                   {"Id"},
			"get-latest-assessment-id":               {},
			"get-portfolio-preferences":              {},
			"get-portfolio-summary":                  {},
			"get-recommendation-report-details":      {"Id"},
			"get-server-details":                     {"MaxResults", "NextToken", "ServerId"},
			"get-server-strategies":                  {"ServerId"},
			"list-analyzable-servers":                {"MaxResults", "NextToken", "Sort"},
			"list-application-components":            {"ApplicationComponentCriteria", "FilterValue", "GroupIdFilter", "MaxResults", "NextToken", "Sort"},
			"list-collectors":                        {"MaxResults", "NextToken"},
			"list-import-file-task":                  {"MaxResults", "NextToken"},
			"list-servers":                           {"FilterValue", "GroupIdFilter", "MaxResults", "NextToken", "ServerCriteria", "Sort"},
			"put-portfolio-preferences":              {"ApplicationMode", "ApplicationPreferences", "DatabasePreferences", "PrioritizeBusinessGoals"},
			"start-assessment":                       {"AssessmentDataSourceType", "AssessmentTargets", "S3bucketForAnalysisData", "S3bucketForReportData"},
			"start-import-file-task":                 {"DataSourceType", "GroupId", "Name", "S3Bucket", "S3bucketForReportData", "S3key"},
			"start-recommendation-report-generation": {"GroupIdFilter", "OutputFormat"},
			"stop-assessment":                        {"AssessmentId"},
			"update-application-component-config":    {"AppType", "ApplicationComponentId", "ConfigureOnly", "InclusionStatus", "SecretsManagerKey", "SourceCodeList", "StrategyOption"},
			"update-server-config":                   {"ServerId", "StrategyOption"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-application-component-details":      {"ApplicationComponentId": "*string"},
			"get-application-component-strategies":   {"ApplicationComponentId": "*string"},
			"get-assessment":                         {"Id": "*string"},
			"get-import-file-task":                   {"Id": "*string"},
			"get-latest-assessment-id":               {},
			"get-portfolio-preferences":              {},
			"get-portfolio-summary":                  {},
			"get-recommendation-report-details":      {"Id": "*string"},
			"get-server-details":                     {"MaxResults": "*int32", "NextToken": "*string", "ServerId": "*string"},
			"get-server-strategies":                  {"ServerId": "*string"},
			"list-analyzable-servers":                {"MaxResults": "*int32", "NextToken": "*string", "Sort": "types.SortOrder"},
			"list-application-components":            {"ApplicationComponentCriteria": "types.ApplicationComponentCriteria", "FilterValue": "*string", "GroupIdFilter": "[]types.Group", "MaxResults": "*int32", "NextToken": "*string", "Sort": "types.SortOrder"},
			"list-collectors":                        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-import-file-task":                  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-servers":                           {"FilterValue": "*string", "GroupIdFilter": "[]types.Group", "MaxResults": "*int32", "NextToken": "*string", "ServerCriteria": "types.ServerCriteria", "Sort": "types.SortOrder"},
			"put-portfolio-preferences":              {"ApplicationMode": "types.ApplicationMode", "ApplicationPreferences": "*types.ApplicationPreferences", "DatabasePreferences": "*types.DatabasePreferences", "PrioritizeBusinessGoals": "*types.PrioritizeBusinessGoals"},
			"start-assessment":                       {"AssessmentDataSourceType": "types.AssessmentDataSourceType", "AssessmentTargets": "[]types.AssessmentTarget", "S3bucketForAnalysisData": "*string", "S3bucketForReportData": "*string"},
			"start-import-file-task":                 {"DataSourceType": "types.DataSourceType", "GroupId": "[]types.Group", "Name": "*string", "S3Bucket": "*string", "S3bucketForReportData": "*string", "S3key": "*string"},
			"start-recommendation-report-generation": {"GroupIdFilter": "[]types.Group", "OutputFormat": "types.OutputFormat"},
			"stop-assessment":                        {"AssessmentId": "*string"},
			"update-application-component-config":    {"AppType": "types.AppType", "ApplicationComponentId": "*string", "ConfigureOnly": "*bool", "InclusionStatus": "types.InclusionStatus", "SecretsManagerKey": "*string", "SourceCodeList": "[]types.SourceCode", "StrategyOption": "*types.StrategyOption"},
			"update-server-config":                   {"ServerId": "*string", "StrategyOption": "*types.StrategyOption"},
		},
		OperationInputRequired: map[string][]string{
			"get-application-component-details":      {"ApplicationComponentId"},
			"get-application-component-strategies":   {"ApplicationComponentId"},
			"get-assessment":                         {"Id"},
			"get-import-file-task":                   {"Id"},
			"get-latest-assessment-id":               {},
			"get-portfolio-preferences":              {},
			"get-portfolio-summary":                  {},
			"get-recommendation-report-details":      {"Id"},
			"get-server-details":                     {"ServerId"},
			"get-server-strategies":                  {"ServerId"},
			"list-analyzable-servers":                {},
			"list-application-components":            {},
			"list-collectors":                        {},
			"list-import-file-task":                  {},
			"list-servers":                           {},
			"put-portfolio-preferences":              {},
			"start-assessment":                       {},
			"start-import-file-task":                 {"Name", "S3Bucket", "S3key"},
			"start-recommendation-report-generation": {},
			"stop-assessment":                        {"AssessmentId"},
			"update-application-component-config":    {"ApplicationComponentId"},
			"update-server-config":                   {"ServerId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("migrationhubstrategy", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
