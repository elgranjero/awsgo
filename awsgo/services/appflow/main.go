package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/appflow/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-flow-executions", "create-connector-profile", "create-flow", "delete-connector-profile", "delete-flow", "describe-connector", "describe-connector-entity", "describe-connector-profiles", "describe-connectors", "describe-flow", "describe-flow-execution-records", "list-connector-entities", "list-connectors", "list-flows", "list-tags-for-resource", "register-connector", "reset-connector-metadata-cache", "start-flow", "stop-flow", "tag-resource", "unregister-connector", "untag-resource", "update-connector-profile", "update-connector-registration", "update-flow"},
		OperationSet: map[string]bool{"cancel-flow-executions": true, "create-connector-profile": true, "create-flow": true, "delete-connector-profile": true, "delete-flow": true, "describe-connector": true, "describe-connector-entity": true, "describe-connector-profiles": true, "describe-connectors": true, "describe-flow": true, "describe-flow-execution-records": true, "list-connector-entities": true, "list-connectors": true, "list-flows": true, "list-tags-for-resource": true, "register-connector": true, "reset-connector-metadata-cache": true, "start-flow": true, "stop-flow": true, "tag-resource": true, "unregister-connector": true, "untag-resource": true, "update-connector-profile": true, "update-connector-registration": true, "update-flow": true},
		OperationInputs: map[string][]string{
			"cancel-flow-executions":          {"ExecutionIds", "FlowName"},
			"create-connector-profile":        {"ClientToken", "ConnectionMode", "ConnectorLabel", "ConnectorProfileConfig", "ConnectorProfileName", "ConnectorType", "KmsArn"},
			"create-flow":                     {"ClientToken", "Description", "DestinationFlowConfigList", "FlowName", "KmsArn", "MetadataCatalogConfig", "SourceFlowConfig", "Tags", "Tasks", "TriggerConfig"},
			"delete-connector-profile":        {"ConnectorProfileName", "ForceDelete"},
			"delete-flow":                     {"FlowName", "ForceDelete"},
			"describe-connector":              {"ConnectorLabel", "ConnectorType"},
			"describe-connector-entity":       {"ApiVersion", "ConnectorEntityName", "ConnectorProfileName", "ConnectorType"},
			"describe-connector-profiles":     {"ConnectorLabel", "ConnectorProfileNames", "ConnectorType", "MaxResults", "NextToken"},
			"describe-connectors":             {"ConnectorTypes", "MaxResults", "NextToken"},
			"describe-flow":                   {"FlowName"},
			"describe-flow-execution-records": {"FlowName", "MaxResults", "NextToken"},
			"list-connector-entities":         {"ApiVersion", "ConnectorProfileName", "ConnectorType", "EntitiesPath", "MaxResults", "NextToken"},
			"list-connectors":                 {"MaxResults", "NextToken"},
			"list-flows":                      {"MaxResults", "NextToken"},
			"list-tags-for-resource":          {"ResourceArn"},
			"register-connector":              {"ClientToken", "ConnectorLabel", "ConnectorProvisioningConfig", "ConnectorProvisioningType", "Description"},
			"reset-connector-metadata-cache":  {"ApiVersion", "ConnectorEntityName", "ConnectorProfileName", "ConnectorType", "EntitiesPath"},
			"start-flow":                      {"ClientToken", "FlowName"},
			"stop-flow":                       {"FlowName"},
			"tag-resource":                    {"ResourceArn", "Tags"},
			"unregister-connector":            {"ConnectorLabel", "ForceDelete"},
			"untag-resource":                  {"ResourceArn", "TagKeys"},
			"update-connector-profile":        {"ClientToken", "ConnectionMode", "ConnectorProfileConfig", "ConnectorProfileName"},
			"update-connector-registration":   {"ClientToken", "ConnectorLabel", "ConnectorProvisioningConfig", "Description"},
			"update-flow":                     {"ClientToken", "Description", "DestinationFlowConfigList", "FlowName", "MetadataCatalogConfig", "SourceFlowConfig", "Tasks", "TriggerConfig"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-flow-executions":          {"ExecutionIds": "[]string", "FlowName": "*string"},
			"create-connector-profile":        {"ClientToken": "*string", "ConnectionMode": "types.ConnectionMode", "ConnectorLabel": "*string", "ConnectorProfileConfig": "*types.ConnectorProfileConfig", "ConnectorProfileName": "*string", "ConnectorType": "types.ConnectorType", "KmsArn": "*string"},
			"create-flow":                     {"ClientToken": "*string", "Description": "*string", "DestinationFlowConfigList": "[]types.DestinationFlowConfig", "FlowName": "*string", "KmsArn": "*string", "MetadataCatalogConfig": "*types.MetadataCatalogConfig", "SourceFlowConfig": "*types.SourceFlowConfig", "Tags": "map[string]string", "Tasks": "[]types.Task", "TriggerConfig": "*types.TriggerConfig"},
			"delete-connector-profile":        {"ConnectorProfileName": "*string", "ForceDelete": "bool"},
			"delete-flow":                     {"FlowName": "*string", "ForceDelete": "bool"},
			"describe-connector":              {"ConnectorLabel": "*string", "ConnectorType": "types.ConnectorType"},
			"describe-connector-entity":       {"ApiVersion": "*string", "ConnectorEntityName": "*string", "ConnectorProfileName": "*string", "ConnectorType": "types.ConnectorType"},
			"describe-connector-profiles":     {"ConnectorLabel": "*string", "ConnectorProfileNames": "[]string", "ConnectorType": "types.ConnectorType", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-connectors":             {"ConnectorTypes": "[]types.ConnectorType", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-flow":                   {"FlowName": "*string"},
			"describe-flow-execution-records": {"FlowName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-connector-entities":         {"ApiVersion": "*string", "ConnectorProfileName": "*string", "ConnectorType": "types.ConnectorType", "EntitiesPath": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-connectors":                 {"MaxResults": "*int32", "NextToken": "*string"},
			"list-flows":                      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":          {"ResourceArn": "*string"},
			"register-connector":              {"ClientToken": "*string", "ConnectorLabel": "*string", "ConnectorProvisioningConfig": "*types.ConnectorProvisioningConfig", "ConnectorProvisioningType": "types.ConnectorProvisioningType", "Description": "*string"},
			"reset-connector-metadata-cache":  {"ApiVersion": "*string", "ConnectorEntityName": "*string", "ConnectorProfileName": "*string", "ConnectorType": "types.ConnectorType", "EntitiesPath": "*string"},
			"start-flow":                      {"ClientToken": "*string", "FlowName": "*string"},
			"stop-flow":                       {"FlowName": "*string"},
			"tag-resource":                    {"ResourceArn": "*string", "Tags": "map[string]string"},
			"unregister-connector":            {"ConnectorLabel": "*string", "ForceDelete": "bool"},
			"untag-resource":                  {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-connector-profile":        {"ClientToken": "*string", "ConnectionMode": "types.ConnectionMode", "ConnectorProfileConfig": "*types.ConnectorProfileConfig", "ConnectorProfileName": "*string"},
			"update-connector-registration":   {"ClientToken": "*string", "ConnectorLabel": "*string", "ConnectorProvisioningConfig": "*types.ConnectorProvisioningConfig", "Description": "*string"},
			"update-flow":                     {"ClientToken": "*string", "Description": "*string", "DestinationFlowConfigList": "[]types.DestinationFlowConfig", "FlowName": "*string", "MetadataCatalogConfig": "*types.MetadataCatalogConfig", "SourceFlowConfig": "*types.SourceFlowConfig", "Tasks": "[]types.Task", "TriggerConfig": "*types.TriggerConfig"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-flow-executions":          {"FlowName"},
			"create-connector-profile":        {"ConnectionMode", "ConnectorProfileConfig", "ConnectorProfileName", "ConnectorType"},
			"create-flow":                     {"DestinationFlowConfigList", "FlowName", "SourceFlowConfig", "Tasks", "TriggerConfig"},
			"delete-connector-profile":        {"ConnectorProfileName"},
			"delete-flow":                     {"FlowName"},
			"describe-connector":              {"ConnectorType"},
			"describe-connector-entity":       {"ConnectorEntityName"},
			"describe-connector-profiles":     {},
			"describe-connectors":             {},
			"describe-flow":                   {"FlowName"},
			"describe-flow-execution-records": {"FlowName"},
			"list-connector-entities":         {},
			"list-connectors":                 {},
			"list-flows":                      {},
			"list-tags-for-resource":          {"ResourceArn"},
			"register-connector":              {},
			"reset-connector-metadata-cache":  {},
			"start-flow":                      {"FlowName"},
			"stop-flow":                       {"FlowName"},
			"tag-resource":                    {"ResourceArn", "Tags"},
			"unregister-connector":            {"ConnectorLabel"},
			"untag-resource":                  {"ResourceArn", "TagKeys"},
			"update-connector-profile":        {"ConnectionMode", "ConnectorProfileConfig", "ConnectorProfileName"},
			"update-connector-registration":   {"ConnectorLabel"},
			"update-flow":                     {"DestinationFlowConfigList", "FlowName", "SourceFlowConfig", "Tasks", "TriggerConfig"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("appflow", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
