package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/iotdeviceadvisor/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-suite-definition", "delete-suite-definition", "get-endpoint", "get-suite-definition", "get-suite-run", "get-suite-run-report", "list-suite-definitions", "list-suite-runs", "list-tags-for-resource", "start-suite-run", "stop-suite-run", "tag-resource", "untag-resource", "update-suite-definition"},
		OperationSet: map[string]bool{"create-suite-definition": true, "delete-suite-definition": true, "get-endpoint": true, "get-suite-definition": true, "get-suite-run": true, "get-suite-run-report": true, "list-suite-definitions": true, "list-suite-runs": true, "list-tags-for-resource": true, "start-suite-run": true, "stop-suite-run": true, "tag-resource": true, "untag-resource": true, "update-suite-definition": true},
		OperationInputs: map[string][]string{
			"create-suite-definition": {"ClientToken", "SuiteDefinitionConfiguration", "Tags"},
			"delete-suite-definition": {"SuiteDefinitionId"},
			"get-endpoint":            {"AuthenticationMethod", "CertificateArn", "DeviceRoleArn", "ThingArn"},
			"get-suite-definition":    {"SuiteDefinitionId", "SuiteDefinitionVersion"},
			"get-suite-run":           {"SuiteDefinitionId", "SuiteRunId"},
			"get-suite-run-report":    {"SuiteDefinitionId", "SuiteRunId"},
			"list-suite-definitions":  {"MaxResults", "NextToken"},
			"list-suite-runs":         {"MaxResults", "NextToken", "SuiteDefinitionId", "SuiteDefinitionVersion"},
			"list-tags-for-resource":  {"ResourceArn"},
			"start-suite-run":         {"SuiteDefinitionId", "SuiteDefinitionVersion", "SuiteRunConfiguration", "Tags"},
			"stop-suite-run":          {"SuiteDefinitionId", "SuiteRunId"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
			"update-suite-definition": {"SuiteDefinitionConfiguration", "SuiteDefinitionId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-suite-definition": {"ClientToken": "*string", "SuiteDefinitionConfiguration": "*types.SuiteDefinitionConfiguration", "Tags": "map[string]string"},
			"delete-suite-definition": {"SuiteDefinitionId": "*string"},
			"get-endpoint":            {"AuthenticationMethod": "types.AuthenticationMethod", "CertificateArn": "*string", "DeviceRoleArn": "*string", "ThingArn": "*string"},
			"get-suite-definition":    {"SuiteDefinitionId": "*string", "SuiteDefinitionVersion": "*string"},
			"get-suite-run":           {"SuiteDefinitionId": "*string", "SuiteRunId": "*string"},
			"get-suite-run-report":    {"SuiteDefinitionId": "*string", "SuiteRunId": "*string"},
			"list-suite-definitions":  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-suite-runs":         {"MaxResults": "*int32", "NextToken": "*string", "SuiteDefinitionId": "*string", "SuiteDefinitionVersion": "*string"},
			"list-tags-for-resource":  {"ResourceArn": "*string"},
			"start-suite-run":         {"SuiteDefinitionId": "*string", "SuiteDefinitionVersion": "*string", "SuiteRunConfiguration": "*types.SuiteRunConfiguration", "Tags": "map[string]string"},
			"stop-suite-run":          {"SuiteDefinitionId": "*string", "SuiteRunId": "*string"},
			"tag-resource":            {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":          {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-suite-definition": {"SuiteDefinitionConfiguration": "*types.SuiteDefinitionConfiguration", "SuiteDefinitionId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-suite-definition": {"SuiteDefinitionConfiguration"},
			"delete-suite-definition": {"SuiteDefinitionId"},
			"get-endpoint":            {},
			"get-suite-definition":    {"SuiteDefinitionId"},
			"get-suite-run":           {"SuiteDefinitionId", "SuiteRunId"},
			"get-suite-run-report":    {"SuiteDefinitionId", "SuiteRunId"},
			"list-suite-definitions":  {},
			"list-suite-runs":         {},
			"list-tags-for-resource":  {"ResourceArn"},
			"start-suite-run":         {"SuiteDefinitionId", "SuiteRunConfiguration"},
			"stop-suite-run":          {"SuiteDefinitionId", "SuiteRunId"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
			"update-suite-definition": {"SuiteDefinitionConfiguration", "SuiteDefinitionId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("iotdeviceadvisor", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
