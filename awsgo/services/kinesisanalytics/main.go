package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kinesisanalytics/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-application-cloud-watch-logging-option", "add-application-input", "add-application-input-processing-configuration", "add-application-output", "add-application-reference-data-source", "create-application", "delete-application", "delete-application-cloud-watch-logging-option", "delete-application-input-processing-configuration", "delete-application-output", "delete-application-reference-data-source", "describe-application", "discover-input-schema", "list-applications", "list-tags-for-resource", "start-application", "stop-application", "tag-resource", "untag-resource", "update-application"},
		OperationSet: map[string]bool{"add-application-cloud-watch-logging-option": true, "add-application-input": true, "add-application-input-processing-configuration": true, "add-application-output": true, "add-application-reference-data-source": true, "create-application": true, "delete-application": true, "delete-application-cloud-watch-logging-option": true, "delete-application-input-processing-configuration": true, "delete-application-output": true, "delete-application-reference-data-source": true, "describe-application": true, "discover-input-schema": true, "list-applications": true, "list-tags-for-resource": true, "start-application": true, "stop-application": true, "tag-resource": true, "untag-resource": true, "update-application": true},
		OperationInputs: map[string][]string{
			"add-application-cloud-watch-logging-option":        {"ApplicationName", "CloudWatchLoggingOption", "CurrentApplicationVersionId"},
			"add-application-input":                             {"ApplicationName", "CurrentApplicationVersionId", "Input"},
			"add-application-input-processing-configuration":    {"ApplicationName", "CurrentApplicationVersionId", "InputId", "InputProcessingConfiguration"},
			"add-application-output":                            {"ApplicationName", "CurrentApplicationVersionId", "Output"},
			"add-application-reference-data-source":             {"ApplicationName", "CurrentApplicationVersionId", "ReferenceDataSource"},
			"create-application":                                {"ApplicationCode", "ApplicationDescription", "ApplicationName", "CloudWatchLoggingOptions", "Inputs", "Outputs", "Tags"},
			"delete-application":                                {"ApplicationName", "CreateTimestamp"},
			"delete-application-cloud-watch-logging-option":     {"ApplicationName", "CloudWatchLoggingOptionId", "CurrentApplicationVersionId"},
			"delete-application-input-processing-configuration": {"ApplicationName", "CurrentApplicationVersionId", "InputId"},
			"delete-application-output":                         {"ApplicationName", "CurrentApplicationVersionId", "OutputId"},
			"delete-application-reference-data-source":          {"ApplicationName", "CurrentApplicationVersionId", "ReferenceId"},
			"describe-application":                              {"ApplicationName"},
			"discover-input-schema":                             {"InputProcessingConfiguration", "InputStartingPositionConfiguration", "ResourceARN", "RoleARN", "S3Configuration"},
			"list-applications":                                 {"ExclusiveStartApplicationName", "Limit"},
			"list-tags-for-resource":                            {"ResourceARN"},
			"start-application":                                 {"ApplicationName", "InputConfigurations"},
			"stop-application":                                  {"ApplicationName"},
			"tag-resource":                                      {"ResourceARN", "Tags"},
			"untag-resource":                                    {"ResourceARN", "TagKeys"},
			"update-application":                                {"ApplicationName", "ApplicationUpdate", "CurrentApplicationVersionId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-application-cloud-watch-logging-option":        {"ApplicationName": "*string", "CloudWatchLoggingOption": "*types.CloudWatchLoggingOption", "CurrentApplicationVersionId": "*int64"},
			"add-application-input":                             {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "Input": "*types.Input"},
			"add-application-input-processing-configuration":    {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "InputId": "*string", "InputProcessingConfiguration": "*types.InputProcessingConfiguration"},
			"add-application-output":                            {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "Output": "*types.Output"},
			"add-application-reference-data-source":             {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "ReferenceDataSource": "*types.ReferenceDataSource"},
			"create-application":                                {"ApplicationCode": "*string", "ApplicationDescription": "*string", "ApplicationName": "*string", "CloudWatchLoggingOptions": "[]types.CloudWatchLoggingOption", "Inputs": "[]types.Input", "Outputs": "[]types.Output", "Tags": "[]types.Tag"},
			"delete-application":                                {"ApplicationName": "*string", "CreateTimestamp": "*time.Time"},
			"delete-application-cloud-watch-logging-option":     {"ApplicationName": "*string", "CloudWatchLoggingOptionId": "*string", "CurrentApplicationVersionId": "*int64"},
			"delete-application-input-processing-configuration": {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "InputId": "*string"},
			"delete-application-output":                         {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "OutputId": "*string"},
			"delete-application-reference-data-source":          {"ApplicationName": "*string", "CurrentApplicationVersionId": "*int64", "ReferenceId": "*string"},
			"describe-application":                              {"ApplicationName": "*string"},
			"discover-input-schema":                             {"InputProcessingConfiguration": "*types.InputProcessingConfiguration", "InputStartingPositionConfiguration": "*types.InputStartingPositionConfiguration", "ResourceARN": "*string", "RoleARN": "*string", "S3Configuration": "*types.S3Configuration"},
			"list-applications":                                 {"ExclusiveStartApplicationName": "*string", "Limit": "*int32"},
			"list-tags-for-resource":                            {"ResourceARN": "*string"},
			"start-application":                                 {"ApplicationName": "*string", "InputConfigurations": "[]types.InputConfiguration"},
			"stop-application":                                  {"ApplicationName": "*string"},
			"tag-resource":                                      {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                                    {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-application":                                {"ApplicationName": "*string", "ApplicationUpdate": "*types.ApplicationUpdate", "CurrentApplicationVersionId": "*int64"},
		},
		OperationInputRequired: map[string][]string{
			"add-application-cloud-watch-logging-option":        {"ApplicationName", "CloudWatchLoggingOption", "CurrentApplicationVersionId"},
			"add-application-input":                             {"ApplicationName", "CurrentApplicationVersionId", "Input"},
			"add-application-input-processing-configuration":    {"ApplicationName", "CurrentApplicationVersionId", "InputId", "InputProcessingConfiguration"},
			"add-application-output":                            {"ApplicationName", "CurrentApplicationVersionId", "Output"},
			"add-application-reference-data-source":             {"ApplicationName", "CurrentApplicationVersionId", "ReferenceDataSource"},
			"create-application":                                {"ApplicationName"},
			"delete-application":                                {"ApplicationName", "CreateTimestamp"},
			"delete-application-cloud-watch-logging-option":     {"ApplicationName", "CloudWatchLoggingOptionId", "CurrentApplicationVersionId"},
			"delete-application-input-processing-configuration": {"ApplicationName", "CurrentApplicationVersionId", "InputId"},
			"delete-application-output":                         {"ApplicationName", "CurrentApplicationVersionId", "OutputId"},
			"delete-application-reference-data-source":          {"ApplicationName", "CurrentApplicationVersionId", "ReferenceId"},
			"describe-application":                              {"ApplicationName"},
			"discover-input-schema":                             {},
			"list-applications":                                 {},
			"list-tags-for-resource":                            {"ResourceARN"},
			"start-application":                                 {"ApplicationName", "InputConfigurations"},
			"stop-application":                                  {"ApplicationName"},
			"tag-resource":                                      {"ResourceARN", "Tags"},
			"untag-resource":                                    {"ResourceARN", "TagKeys"},
			"update-application":                                {"ApplicationName", "ApplicationUpdate", "CurrentApplicationVersionId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kinesisanalytics", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
