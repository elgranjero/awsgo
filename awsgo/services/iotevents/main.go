package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/iotevents/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-alarm-model", "create-detector-model", "create-input", "delete-alarm-model", "delete-detector-model", "delete-input", "describe-alarm-model", "describe-detector-model", "describe-detector-model-analysis", "describe-input", "describe-logging-options", "get-detector-model-analysis-results", "list-alarm-model-versions", "list-alarm-models", "list-detector-model-versions", "list-detector-models", "list-input-routings", "list-inputs", "list-tags-for-resource", "put-logging-options", "start-detector-model-analysis", "tag-resource", "untag-resource", "update-alarm-model", "update-detector-model", "update-input"},
		OperationSet: map[string]bool{"create-alarm-model": true, "create-detector-model": true, "create-input": true, "delete-alarm-model": true, "delete-detector-model": true, "delete-input": true, "describe-alarm-model": true, "describe-detector-model": true, "describe-detector-model-analysis": true, "describe-input": true, "describe-logging-options": true, "get-detector-model-analysis-results": true, "list-alarm-model-versions": true, "list-alarm-models": true, "list-detector-model-versions": true, "list-detector-models": true, "list-input-routings": true, "list-inputs": true, "list-tags-for-resource": true, "put-logging-options": true, "start-detector-model-analysis": true, "tag-resource": true, "untag-resource": true, "update-alarm-model": true, "update-detector-model": true, "update-input": true},
		OperationInputs: map[string][]string{
			"create-alarm-model":                  {"AlarmCapabilities", "AlarmEventActions", "AlarmModelDescription", "AlarmModelName", "AlarmNotification", "AlarmRule", "Key", "RoleArn", "Severity", "Tags"},
			"create-detector-model":               {"DetectorModelDefinition", "DetectorModelDescription", "DetectorModelName", "EvaluationMethod", "Key", "RoleArn", "Tags"},
			"create-input":                        {"InputDefinition", "InputDescription", "InputName", "Tags"},
			"delete-alarm-model":                  {"AlarmModelName"},
			"delete-detector-model":               {"DetectorModelName"},
			"delete-input":                        {"InputName"},
			"describe-alarm-model":                {"AlarmModelName", "AlarmModelVersion"},
			"describe-detector-model":             {"DetectorModelName", "DetectorModelVersion"},
			"describe-detector-model-analysis":    {"AnalysisId"},
			"describe-input":                      {"InputName"},
			"describe-logging-options":            {},
			"get-detector-model-analysis-results": {"AnalysisId", "MaxResults", "NextToken"},
			"list-alarm-model-versions":           {"AlarmModelName", "MaxResults", "NextToken"},
			"list-alarm-models":                   {"MaxResults", "NextToken"},
			"list-detector-model-versions":        {"DetectorModelName", "MaxResults", "NextToken"},
			"list-detector-models":                {"MaxResults", "NextToken"},
			"list-input-routings":                 {"InputIdentifier", "MaxResults", "NextToken"},
			"list-inputs":                         {"MaxResults", "NextToken"},
			"list-tags-for-resource":              {"ResourceArn"},
			"put-logging-options":                 {"LoggingOptions"},
			"start-detector-model-analysis":       {"DetectorModelDefinition"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-alarm-model":                  {"AlarmCapabilities", "AlarmEventActions", "AlarmModelDescription", "AlarmModelName", "AlarmNotification", "AlarmRule", "RoleArn", "Severity"},
			"update-detector-model":               {"DetectorModelDefinition", "DetectorModelDescription", "DetectorModelName", "EvaluationMethod", "RoleArn"},
			"update-input":                        {"InputDefinition", "InputDescription", "InputName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-alarm-model":                  {"AlarmCapabilities": "*types.AlarmCapabilities", "AlarmEventActions": "*types.AlarmEventActions", "AlarmModelDescription": "*string", "AlarmModelName": "*string", "AlarmNotification": "*types.AlarmNotification", "AlarmRule": "*types.AlarmRule", "Key": "*string", "RoleArn": "*string", "Severity": "*int32", "Tags": "[]types.Tag"},
			"create-detector-model":               {"DetectorModelDefinition": "*types.DetectorModelDefinition", "DetectorModelDescription": "*string", "DetectorModelName": "*string", "EvaluationMethod": "types.EvaluationMethod", "Key": "*string", "RoleArn": "*string", "Tags": "[]types.Tag"},
			"create-input":                        {"InputDefinition": "*types.InputDefinition", "InputDescription": "*string", "InputName": "*string", "Tags": "[]types.Tag"},
			"delete-alarm-model":                  {"AlarmModelName": "*string"},
			"delete-detector-model":               {"DetectorModelName": "*string"},
			"delete-input":                        {"InputName": "*string"},
			"describe-alarm-model":                {"AlarmModelName": "*string", "AlarmModelVersion": "*string"},
			"describe-detector-model":             {"DetectorModelName": "*string", "DetectorModelVersion": "*string"},
			"describe-detector-model-analysis":    {"AnalysisId": "*string"},
			"describe-input":                      {"InputName": "*string"},
			"describe-logging-options":            {},
			"get-detector-model-analysis-results": {"AnalysisId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-alarm-model-versions":           {"AlarmModelName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-alarm-models":                   {"MaxResults": "*int32", "NextToken": "*string"},
			"list-detector-model-versions":        {"DetectorModelName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-detector-models":                {"MaxResults": "*int32", "NextToken": "*string"},
			"list-input-routings":                 {"InputIdentifier": "*types.InputIdentifier", "MaxResults": "*int32", "NextToken": "*string"},
			"list-inputs":                         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":              {"ResourceArn": "*string"},
			"put-logging-options":                 {"LoggingOptions": "*types.LoggingOptions"},
			"start-detector-model-analysis":       {"DetectorModelDefinition": "*types.DetectorModelDefinition"},
			"tag-resource":                        {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                      {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-alarm-model":                  {"AlarmCapabilities": "*types.AlarmCapabilities", "AlarmEventActions": "*types.AlarmEventActions", "AlarmModelDescription": "*string", "AlarmModelName": "*string", "AlarmNotification": "*types.AlarmNotification", "AlarmRule": "*types.AlarmRule", "RoleArn": "*string", "Severity": "*int32"},
			"update-detector-model":               {"DetectorModelDefinition": "*types.DetectorModelDefinition", "DetectorModelDescription": "*string", "DetectorModelName": "*string", "EvaluationMethod": "types.EvaluationMethod", "RoleArn": "*string"},
			"update-input":                        {"InputDefinition": "*types.InputDefinition", "InputDescription": "*string", "InputName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-alarm-model":                  {"AlarmModelName", "AlarmRule", "RoleArn"},
			"create-detector-model":               {"DetectorModelDefinition", "DetectorModelName", "RoleArn"},
			"create-input":                        {"InputDefinition", "InputName"},
			"delete-alarm-model":                  {"AlarmModelName"},
			"delete-detector-model":               {"DetectorModelName"},
			"delete-input":                        {"InputName"},
			"describe-alarm-model":                {"AlarmModelName"},
			"describe-detector-model":             {"DetectorModelName"},
			"describe-detector-model-analysis":    {"AnalysisId"},
			"describe-input":                      {"InputName"},
			"describe-logging-options":            {},
			"get-detector-model-analysis-results": {"AnalysisId"},
			"list-alarm-model-versions":           {"AlarmModelName"},
			"list-alarm-models":                   {},
			"list-detector-model-versions":        {"DetectorModelName"},
			"list-detector-models":                {},
			"list-input-routings":                 {"InputIdentifier"},
			"list-inputs":                         {},
			"list-tags-for-resource":              {"ResourceArn"},
			"put-logging-options":                 {"LoggingOptions"},
			"start-detector-model-analysis":       {"DetectorModelDefinition"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-alarm-model":                  {"AlarmModelName", "AlarmRule", "RoleArn"},
			"update-detector-model":               {"DetectorModelDefinition", "DetectorModelName", "RoleArn"},
			"update-input":                        {"InputDefinition", "InputName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("iotevents", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
