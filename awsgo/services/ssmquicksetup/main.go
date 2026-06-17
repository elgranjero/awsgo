package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ssmquicksetup/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-configuration-manager", "delete-configuration-manager", "get-configuration", "get-configuration-manager", "get-service-settings", "list-configuration-managers", "list-configurations", "list-quick-setup-types", "list-tags-for-resource", "tag-resource", "untag-resource", "update-configuration-definition", "update-configuration-manager", "update-service-settings"},
		OperationSet: map[string]bool{"create-configuration-manager": true, "delete-configuration-manager": true, "get-configuration": true, "get-configuration-manager": true, "get-service-settings": true, "list-configuration-managers": true, "list-configurations": true, "list-quick-setup-types": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-configuration-definition": true, "update-configuration-manager": true, "update-service-settings": true},
		OperationInputs: map[string][]string{
			"create-configuration-manager":    {"ConfigurationDefinitions", "Description", "Name", "Tags"},
			"delete-configuration-manager":    {"ManagerArn"},
			"get-configuration":               {"ConfigurationId"},
			"get-configuration-manager":       {"ManagerArn"},
			"get-service-settings":            {},
			"list-configuration-managers":     {"Filters", "MaxItems", "StartingToken"},
			"list-configurations":             {"ConfigurationDefinitionId", "Filters", "ManagerArn", "MaxItems", "StartingToken"},
			"list-quick-setup-types":          {},
			"list-tags-for-resource":          {"ResourceArn"},
			"tag-resource":                    {"ResourceArn", "Tags"},
			"untag-resource":                  {"ResourceArn", "TagKeys"},
			"update-configuration-definition": {"Id", "LocalDeploymentAdministrationRoleArn", "LocalDeploymentExecutionRoleName", "ManagerArn", "Parameters", "TypeVersion"},
			"update-configuration-manager":    {"Description", "ManagerArn", "Name"},
			"update-service-settings":         {"ExplorerEnablingRoleArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-configuration-manager":    {"ConfigurationDefinitions": "[]types.ConfigurationDefinitionInput", "Description": "*string", "Name": "*string", "Tags": "map[string]string"},
			"delete-configuration-manager":    {"ManagerArn": "*string"},
			"get-configuration":               {"ConfigurationId": "*string"},
			"get-configuration-manager":       {"ManagerArn": "*string"},
			"get-service-settings":            {},
			"list-configuration-managers":     {"Filters": "[]types.Filter", "MaxItems": "*int32", "StartingToken": "*string"},
			"list-configurations":             {"ConfigurationDefinitionId": "*string", "Filters": "[]types.Filter", "ManagerArn": "*string", "MaxItems": "*int32", "StartingToken": "*string"},
			"list-quick-setup-types":          {},
			"list-tags-for-resource":          {"ResourceArn": "*string"},
			"tag-resource":                    {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                  {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-configuration-definition": {"Id": "*string", "LocalDeploymentAdministrationRoleArn": "*string", "LocalDeploymentExecutionRoleName": "*string", "ManagerArn": "*string", "Parameters": "map[string]string", "TypeVersion": "*string"},
			"update-configuration-manager":    {"Description": "*string", "ManagerArn": "*string", "Name": "*string"},
			"update-service-settings":         {"ExplorerEnablingRoleArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-configuration-manager":    {"ConfigurationDefinitions"},
			"delete-configuration-manager":    {"ManagerArn"},
			"get-configuration":               {"ConfigurationId"},
			"get-configuration-manager":       {"ManagerArn"},
			"get-service-settings":            {},
			"list-configuration-managers":     {},
			"list-configurations":             {},
			"list-quick-setup-types":          {},
			"list-tags-for-resource":          {"ResourceArn"},
			"tag-resource":                    {"ResourceArn", "Tags"},
			"untag-resource":                  {"ResourceArn", "TagKeys"},
			"update-configuration-definition": {"Id", "ManagerArn"},
			"update-configuration-manager":    {"ManagerArn"},
			"update-service-settings":         {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ssmquicksetup", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
