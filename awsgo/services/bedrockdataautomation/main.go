package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/bedrockdataautomation/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"copy-blueprint-stage", "create-blueprint", "create-blueprint-version", "create-data-automation-project", "delete-blueprint", "delete-data-automation-project", "get-blueprint", "get-blueprint-optimization-status", "get-data-automation-project", "invoke-blueprint-optimization-async", "list-blueprints", "list-data-automation-projects", "list-tags-for-resource", "tag-resource", "untag-resource", "update-blueprint", "update-data-automation-project"},
		OperationSet: map[string]bool{"copy-blueprint-stage": true, "create-blueprint": true, "create-blueprint-version": true, "create-data-automation-project": true, "delete-blueprint": true, "delete-data-automation-project": true, "get-blueprint": true, "get-blueprint-optimization-status": true, "get-data-automation-project": true, "invoke-blueprint-optimization-async": true, "list-blueprints": true, "list-data-automation-projects": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-blueprint": true, "update-data-automation-project": true},
		OperationInputs: map[string][]string{
			"copy-blueprint-stage":                {"BlueprintArn", "ClientToken", "SourceStage", "TargetStage"},
			"create-blueprint":                    {"BlueprintName", "BlueprintStage", "ClientToken", "EncryptionConfiguration", "Schema", "Tags", "Type"},
			"create-blueprint-version":            {"BlueprintArn", "ClientToken"},
			"create-data-automation-project":      {"ClientToken", "CustomOutputConfiguration", "EncryptionConfiguration", "OverrideConfiguration", "ProjectDescription", "ProjectName", "ProjectStage", "ProjectType", "StandardOutputConfiguration", "Tags"},
			"delete-blueprint":                    {"BlueprintArn", "BlueprintVersion"},
			"delete-data-automation-project":      {"ProjectArn"},
			"get-blueprint":                       {"BlueprintArn", "BlueprintStage", "BlueprintVersion"},
			"get-blueprint-optimization-status":   {"InvocationArn"},
			"get-data-automation-project":         {"ProjectArn", "ProjectStage"},
			"invoke-blueprint-optimization-async": {"Blueprint", "DataAutomationProfileArn", "EncryptionConfiguration", "OutputConfiguration", "Samples", "Tags"},
			"list-blueprints":                     {"BlueprintArn", "BlueprintStageFilter", "MaxResults", "NextToken", "ProjectFilter", "ResourceOwner"},
			"list-data-automation-projects":       {"BlueprintFilter", "MaxResults", "NextToken", "ProjectStageFilter", "ResourceOwner"},
			"list-tags-for-resource":              {"ResourceARN"},
			"tag-resource":                        {"ResourceARN", "Tags"},
			"untag-resource":                      {"ResourceARN", "TagKeys"},
			"update-blueprint":                    {"BlueprintArn", "BlueprintStage", "EncryptionConfiguration", "Schema"},
			"update-data-automation-project":      {"CustomOutputConfiguration", "EncryptionConfiguration", "OverrideConfiguration", "ProjectArn", "ProjectDescription", "ProjectStage", "StandardOutputConfiguration"},
		},
		OperationInputTypes: map[string]map[string]string{
			"copy-blueprint-stage":                {"BlueprintArn": "*string", "ClientToken": "*string", "SourceStage": "types.BlueprintStage", "TargetStage": "types.BlueprintStage"},
			"create-blueprint":                    {"BlueprintName": "*string", "BlueprintStage": "types.BlueprintStage", "ClientToken": "*string", "EncryptionConfiguration": "*types.EncryptionConfiguration", "Schema": "*string", "Tags": "[]types.Tag", "Type": "types.Type"},
			"create-blueprint-version":            {"BlueprintArn": "*string", "ClientToken": "*string"},
			"create-data-automation-project":      {"ClientToken": "*string", "CustomOutputConfiguration": "*types.CustomOutputConfiguration", "EncryptionConfiguration": "*types.EncryptionConfiguration", "OverrideConfiguration": "*types.OverrideConfiguration", "ProjectDescription": "*string", "ProjectName": "*string", "ProjectStage": "types.DataAutomationProjectStage", "ProjectType": "types.DataAutomationProjectType", "StandardOutputConfiguration": "*types.StandardOutputConfiguration", "Tags": "[]types.Tag"},
			"delete-blueprint":                    {"BlueprintArn": "*string", "BlueprintVersion": "*string"},
			"delete-data-automation-project":      {"ProjectArn": "*string"},
			"get-blueprint":                       {"BlueprintArn": "*string", "BlueprintStage": "types.BlueprintStage", "BlueprintVersion": "*string"},
			"get-blueprint-optimization-status":   {"InvocationArn": "*string"},
			"get-data-automation-project":         {"ProjectArn": "*string", "ProjectStage": "types.DataAutomationProjectStage"},
			"invoke-blueprint-optimization-async": {"Blueprint": "*types.BlueprintOptimizationObject", "DataAutomationProfileArn": "*string", "EncryptionConfiguration": "*types.EncryptionConfiguration", "OutputConfiguration": "*types.BlueprintOptimizationOutputConfiguration", "Samples": "[]types.BlueprintOptimizationSample", "Tags": "[]types.Tag"},
			"list-blueprints":                     {"BlueprintArn": "*string", "BlueprintStageFilter": "types.BlueprintStageFilter", "MaxResults": "*int32", "NextToken": "*string", "ProjectFilter": "*types.DataAutomationProjectFilter", "ResourceOwner": "types.ResourceOwner"},
			"list-data-automation-projects":       {"BlueprintFilter": "*types.BlueprintFilter", "MaxResults": "*int32", "NextToken": "*string", "ProjectStageFilter": "types.DataAutomationProjectStageFilter", "ResourceOwner": "types.ResourceOwner"},
			"list-tags-for-resource":              {"ResourceARN": "*string"},
			"tag-resource":                        {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                      {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-blueprint":                    {"BlueprintArn": "*string", "BlueprintStage": "types.BlueprintStage", "EncryptionConfiguration": "*types.EncryptionConfiguration", "Schema": "*string"},
			"update-data-automation-project":      {"CustomOutputConfiguration": "*types.CustomOutputConfiguration", "EncryptionConfiguration": "*types.EncryptionConfiguration", "OverrideConfiguration": "*types.OverrideConfiguration", "ProjectArn": "*string", "ProjectDescription": "*string", "ProjectStage": "types.DataAutomationProjectStage", "StandardOutputConfiguration": "*types.StandardOutputConfiguration"},
		},
		OperationInputRequired: map[string][]string{
			"copy-blueprint-stage":                {"BlueprintArn", "SourceStage", "TargetStage"},
			"create-blueprint":                    {"BlueprintName", "Schema", "Type"},
			"create-blueprint-version":            {"BlueprintArn"},
			"create-data-automation-project":      {"ProjectName", "StandardOutputConfiguration"},
			"delete-blueprint":                    {"BlueprintArn"},
			"delete-data-automation-project":      {"ProjectArn"},
			"get-blueprint":                       {"BlueprintArn"},
			"get-blueprint-optimization-status":   {"InvocationArn"},
			"get-data-automation-project":         {"ProjectArn"},
			"invoke-blueprint-optimization-async": {"Blueprint", "DataAutomationProfileArn", "OutputConfiguration", "Samples"},
			"list-blueprints":                     {},
			"list-data-automation-projects":       {},
			"list-tags-for-resource":              {"ResourceARN"},
			"tag-resource":                        {"ResourceARN", "Tags"},
			"untag-resource":                      {"ResourceARN", "TagKeys"},
			"update-blueprint":                    {"BlueprintArn", "Schema"},
			"update-data-automation-project":      {"ProjectArn", "StandardOutputConfiguration"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("bedrockdataautomation", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
