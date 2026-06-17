package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/fis/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-experiment-template", "create-target-account-configuration", "delete-experiment-template", "delete-target-account-configuration", "get-action", "get-experiment", "get-experiment-target-account-configuration", "get-experiment-template", "get-safety-lever", "get-target-account-configuration", "get-target-resource-type", "list-actions", "list-experiment-resolved-targets", "list-experiment-target-account-configurations", "list-experiment-templates", "list-experiments", "list-tags-for-resource", "list-target-account-configurations", "list-target-resource-types", "start-experiment", "stop-experiment", "tag-resource", "untag-resource", "update-experiment-template", "update-safety-lever-state", "update-target-account-configuration"},
		OperationSet: map[string]bool{"create-experiment-template": true, "create-target-account-configuration": true, "delete-experiment-template": true, "delete-target-account-configuration": true, "get-action": true, "get-experiment": true, "get-experiment-target-account-configuration": true, "get-experiment-template": true, "get-safety-lever": true, "get-target-account-configuration": true, "get-target-resource-type": true, "list-actions": true, "list-experiment-resolved-targets": true, "list-experiment-target-account-configurations": true, "list-experiment-templates": true, "list-experiments": true, "list-tags-for-resource": true, "list-target-account-configurations": true, "list-target-resource-types": true, "start-experiment": true, "stop-experiment": true, "tag-resource": true, "untag-resource": true, "update-experiment-template": true, "update-safety-lever-state": true, "update-target-account-configuration": true},
		OperationInputs: map[string][]string{
			"create-experiment-template":                    {"Actions", "ClientToken", "Description", "ExperimentOptions", "ExperimentReportConfiguration", "LogConfiguration", "RoleArn", "StopConditions", "Tags", "Targets"},
			"create-target-account-configuration":           {"AccountId", "ClientToken", "Description", "ExperimentTemplateId", "RoleArn"},
			"delete-experiment-template":                    {"Id"},
			"delete-target-account-configuration":           {"AccountId", "ExperimentTemplateId"},
			"get-action":                                    {"Id"},
			"get-experiment":                                {"Id"},
			"get-experiment-target-account-configuration":   {"AccountId", "ExperimentId"},
			"get-experiment-template":                       {"Id"},
			"get-safety-lever":                              {"Id"},
			"get-target-account-configuration":              {"AccountId", "ExperimentTemplateId"},
			"get-target-resource-type":                      {"ResourceType"},
			"list-actions":                                  {"MaxResults", "NextToken"},
			"list-experiment-resolved-targets":              {"ExperimentId", "MaxResults", "NextToken", "TargetName"},
			"list-experiment-target-account-configurations": {"ExperimentId", "NextToken"},
			"list-experiment-templates":                     {"MaxResults", "NextToken"},
			"list-experiments":                              {"ExperimentTemplateId", "MaxResults", "NextToken"},
			"list-tags-for-resource":                        {"ResourceArn"},
			"list-target-account-configurations":            {"ExperimentTemplateId", "MaxResults", "NextToken"},
			"list-target-resource-types":                    {"MaxResults", "NextToken"},
			"start-experiment":                              {"ClientToken", "ExperimentOptions", "ExperimentTemplateId", "Tags"},
			"stop-experiment":                               {"Id"},
			"tag-resource":                                  {"ResourceArn", "Tags"},
			"untag-resource":                                {"ResourceArn", "TagKeys"},
			"update-experiment-template":                    {"Actions", "Description", "ExperimentOptions", "ExperimentReportConfiguration", "Id", "LogConfiguration", "RoleArn", "StopConditions", "Targets"},
			"update-safety-lever-state":                     {"Id", "State"},
			"update-target-account-configuration":           {"AccountId", "Description", "ExperimentTemplateId", "RoleArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-experiment-template":                    {"Actions": "map[string]types.CreateExperimentTemplateActionInput", "ClientToken": "*string", "Description": "*string", "ExperimentOptions": "*types.CreateExperimentTemplateExperimentOptionsInput", "ExperimentReportConfiguration": "*types.CreateExperimentTemplateReportConfigurationInput", "LogConfiguration": "*types.CreateExperimentTemplateLogConfigurationInput", "RoleArn": "*string", "StopConditions": "[]types.CreateExperimentTemplateStopConditionInput", "Tags": "map[string]string", "Targets": "map[string]types.CreateExperimentTemplateTargetInput"},
			"create-target-account-configuration":           {"AccountId": "*string", "ClientToken": "*string", "Description": "*string", "ExperimentTemplateId": "*string", "RoleArn": "*string"},
			"delete-experiment-template":                    {"Id": "*string"},
			"delete-target-account-configuration":           {"AccountId": "*string", "ExperimentTemplateId": "*string"},
			"get-action":                                    {"Id": "*string"},
			"get-experiment":                                {"Id": "*string"},
			"get-experiment-target-account-configuration":   {"AccountId": "*string", "ExperimentId": "*string"},
			"get-experiment-template":                       {"Id": "*string"},
			"get-safety-lever":                              {"Id": "*string"},
			"get-target-account-configuration":              {"AccountId": "*string", "ExperimentTemplateId": "*string"},
			"get-target-resource-type":                      {"ResourceType": "*string"},
			"list-actions":                                  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-experiment-resolved-targets":              {"ExperimentId": "*string", "MaxResults": "*int32", "NextToken": "*string", "TargetName": "*string"},
			"list-experiment-target-account-configurations": {"ExperimentId": "*string", "NextToken": "*string"},
			"list-experiment-templates":                     {"MaxResults": "*int32", "NextToken": "*string"},
			"list-experiments":                              {"ExperimentTemplateId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                        {"ResourceArn": "*string"},
			"list-target-account-configurations":            {"ExperimentTemplateId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-target-resource-types":                    {"MaxResults": "*int32", "NextToken": "*string"},
			"start-experiment":                              {"ClientToken": "*string", "ExperimentOptions": "*types.StartExperimentExperimentOptionsInput", "ExperimentTemplateId": "*string", "Tags": "map[string]string"},
			"stop-experiment":                               {"Id": "*string"},
			"tag-resource":                                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                                {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-experiment-template":                    {"Actions": "map[string]types.UpdateExperimentTemplateActionInputItem", "Description": "*string", "ExperimentOptions": "*types.UpdateExperimentTemplateExperimentOptionsInput", "ExperimentReportConfiguration": "*types.UpdateExperimentTemplateReportConfigurationInput", "Id": "*string", "LogConfiguration": "*types.UpdateExperimentTemplateLogConfigurationInput", "RoleArn": "*string", "StopConditions": "[]types.UpdateExperimentTemplateStopConditionInput", "Targets": "map[string]types.UpdateExperimentTemplateTargetInput"},
			"update-safety-lever-state":                     {"Id": "*string", "State": "*types.UpdateSafetyLeverStateInput"},
			"update-target-account-configuration":           {"AccountId": "*string", "Description": "*string", "ExperimentTemplateId": "*string", "RoleArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-experiment-template":                    {"Actions", "ClientToken", "Description", "RoleArn", "StopConditions"},
			"create-target-account-configuration":           {"AccountId", "ExperimentTemplateId", "RoleArn"},
			"delete-experiment-template":                    {"Id"},
			"delete-target-account-configuration":           {"AccountId", "ExperimentTemplateId"},
			"get-action":                                    {"Id"},
			"get-experiment":                                {"Id"},
			"get-experiment-target-account-configuration":   {"AccountId", "ExperimentId"},
			"get-experiment-template":                       {"Id"},
			"get-safety-lever":                              {"Id"},
			"get-target-account-configuration":              {"AccountId", "ExperimentTemplateId"},
			"get-target-resource-type":                      {"ResourceType"},
			"list-actions":                                  {},
			"list-experiment-resolved-targets":              {"ExperimentId"},
			"list-experiment-target-account-configurations": {"ExperimentId"},
			"list-experiment-templates":                     {},
			"list-experiments":                              {},
			"list-tags-for-resource":                        {"ResourceArn"},
			"list-target-account-configurations":            {"ExperimentTemplateId"},
			"list-target-resource-types":                    {},
			"start-experiment":                              {"ClientToken", "ExperimentTemplateId"},
			"stop-experiment":                               {"Id"},
			"tag-resource":                                  {"ResourceArn", "Tags"},
			"untag-resource":                                {"ResourceArn"},
			"update-experiment-template":                    {"Id"},
			"update-safety-lever-state":                     {"Id", "State"},
			"update-target-account-configuration":           {"AccountId", "ExperimentTemplateId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("fis", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
