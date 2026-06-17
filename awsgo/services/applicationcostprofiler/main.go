package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/applicationcostprofiler/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-report-definition", "get-report-definition", "import-application-usage", "list-report-definitions", "put-report-definition", "update-report-definition"},
		OperationSet: map[string]bool{"delete-report-definition": true, "get-report-definition": true, "import-application-usage": true, "list-report-definitions": true, "put-report-definition": true, "update-report-definition": true},
		OperationInputs: map[string][]string{
			"delete-report-definition": {"ReportId"},
			"get-report-definition":    {"ReportId"},
			"import-application-usage": {"SourceS3Location"},
			"list-report-definitions":  {"MaxResults", "NextToken"},
			"put-report-definition":    {"DestinationS3Location", "Format", "ReportDescription", "ReportFrequency", "ReportId"},
			"update-report-definition": {"DestinationS3Location", "Format", "ReportDescription", "ReportFrequency", "ReportId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-report-definition": {"ReportId": "*string"},
			"get-report-definition":    {"ReportId": "*string"},
			"import-application-usage": {"SourceS3Location": "*types.SourceS3Location"},
			"list-report-definitions":  {"MaxResults": "*int32", "NextToken": "*string"},
			"put-report-definition":    {"DestinationS3Location": "*types.S3Location", "Format": "types.Format", "ReportDescription": "*string", "ReportFrequency": "types.ReportFrequency", "ReportId": "*string"},
			"update-report-definition": {"DestinationS3Location": "*types.S3Location", "Format": "types.Format", "ReportDescription": "*string", "ReportFrequency": "types.ReportFrequency", "ReportId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-report-definition": {"ReportId"},
			"get-report-definition":    {"ReportId"},
			"import-application-usage": {"SourceS3Location"},
			"list-report-definitions":  {},
			"put-report-definition":    {"DestinationS3Location", "Format", "ReportDescription", "ReportFrequency", "ReportId"},
			"update-report-definition": {"DestinationS3Location", "Format", "ReportDescription", "ReportFrequency", "ReportId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("applicationcostprofiler", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
