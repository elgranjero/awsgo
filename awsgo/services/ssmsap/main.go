package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ssmsap/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-resource-permission", "deregister-application", "get-application", "get-component", "get-configuration-check-operation", "get-database", "get-operation", "get-resource-permission", "list-applications", "list-components", "list-configuration-check-definitions", "list-configuration-check-operations", "list-databases", "list-operation-events", "list-operations", "list-sub-check-results", "list-sub-check-rule-results", "list-tags-for-resource", "put-resource-permission", "register-application", "start-application", "start-application-refresh", "start-configuration-checks", "stop-application", "tag-resource", "untag-resource", "update-application-settings"},
		OperationSet: map[string]bool{"delete-resource-permission": true, "deregister-application": true, "get-application": true, "get-component": true, "get-configuration-check-operation": true, "get-database": true, "get-operation": true, "get-resource-permission": true, "list-applications": true, "list-components": true, "list-configuration-check-definitions": true, "list-configuration-check-operations": true, "list-databases": true, "list-operation-events": true, "list-operations": true, "list-sub-check-results": true, "list-sub-check-rule-results": true, "list-tags-for-resource": true, "put-resource-permission": true, "register-application": true, "start-application": true, "start-application-refresh": true, "start-configuration-checks": true, "stop-application": true, "tag-resource": true, "untag-resource": true, "update-application-settings": true},
		OperationInputs: map[string][]string{
			"delete-resource-permission":           {"ActionType", "ResourceArn", "SourceResourceArn"},
			"deregister-application":               {"ApplicationId"},
			"get-application":                      {"AppRegistryArn", "ApplicationArn", "ApplicationId"},
			"get-component":                        {"ApplicationId", "ComponentId"},
			"get-configuration-check-operation":    {"OperationId"},
			"get-database":                         {"ApplicationId", "ComponentId", "DatabaseArn", "DatabaseId"},
			"get-operation":                        {"OperationId"},
			"get-resource-permission":              {"ActionType", "ResourceArn"},
			"list-applications":                    {"Filters", "MaxResults", "NextToken"},
			"list-components":                      {"ApplicationId", "MaxResults", "NextToken"},
			"list-configuration-check-definitions": {"MaxResults", "NextToken"},
			"list-configuration-check-operations":  {"ApplicationId", "Filters", "ListMode", "MaxResults", "NextToken"},
			"list-databases":                       {"ApplicationId", "ComponentId", "MaxResults", "NextToken"},
			"list-operation-events":                {"Filters", "MaxResults", "NextToken", "OperationId"},
			"list-operations":                      {"ApplicationId", "Filters", "MaxResults", "NextToken"},
			"list-sub-check-results":               {"MaxResults", "NextToken", "OperationId"},
			"list-sub-check-rule-results":          {"MaxResults", "NextToken", "SubCheckResultId"},
			"list-tags-for-resource":               {"ResourceArn"},
			"put-resource-permission":              {"ActionType", "ResourceArn", "SourceResourceArn"},
			"register-application":                 {"ApplicationId", "ApplicationType", "ComponentsInfo", "Credentials", "DatabaseArn", "Instances", "SapInstanceNumber", "Sid", "Tags"},
			"start-application":                    {"ApplicationId"},
			"start-application-refresh":            {"ApplicationId"},
			"start-configuration-checks":           {"ApplicationId", "ConfigurationCheckIds"},
			"stop-application":                     {"ApplicationId", "IncludeEc2InstanceShutdown", "StopConnectedEntity"},
			"tag-resource":                         {"ResourceArn", "Tags"},
			"untag-resource":                       {"ResourceArn", "TagKeys"},
			"update-application-settings":          {"ApplicationId", "Backint", "CredentialsToAddOrUpdate", "CredentialsToRemove", "DatabaseArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-resource-permission":           {"ActionType": "types.PermissionActionType", "ResourceArn": "*string", "SourceResourceArn": "*string"},
			"deregister-application":               {"ApplicationId": "*string"},
			"get-application":                      {"AppRegistryArn": "*string", "ApplicationArn": "*string", "ApplicationId": "*string"},
			"get-component":                        {"ApplicationId": "*string", "ComponentId": "*string"},
			"get-configuration-check-operation":    {"OperationId": "*string"},
			"get-database":                         {"ApplicationId": "*string", "ComponentId": "*string", "DatabaseArn": "*string", "DatabaseId": "*string"},
			"get-operation":                        {"OperationId": "*string"},
			"get-resource-permission":              {"ActionType": "types.PermissionActionType", "ResourceArn": "*string"},
			"list-applications":                    {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-components":                      {"ApplicationId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-configuration-check-definitions": {"MaxResults": "*int32", "NextToken": "*string"},
			"list-configuration-check-operations":  {"ApplicationId": "*string", "Filters": "[]types.Filter", "ListMode": "types.ConfigurationCheckOperationListingMode", "MaxResults": "*int32", "NextToken": "*string"},
			"list-databases":                       {"ApplicationId": "*string", "ComponentId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-operation-events":                {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string", "OperationId": "*string"},
			"list-operations":                      {"ApplicationId": "*string", "Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-sub-check-results":               {"MaxResults": "*int32", "NextToken": "*string", "OperationId": "*string"},
			"list-sub-check-rule-results":          {"MaxResults": "*int32", "NextToken": "*string", "SubCheckResultId": "*string"},
			"list-tags-for-resource":               {"ResourceArn": "*string"},
			"put-resource-permission":              {"ActionType": "types.PermissionActionType", "ResourceArn": "*string", "SourceResourceArn": "*string"},
			"register-application":                 {"ApplicationId": "*string", "ApplicationType": "types.ApplicationType", "ComponentsInfo": "[]types.ComponentInfo", "Credentials": "[]types.ApplicationCredential", "DatabaseArn": "*string", "Instances": "[]string", "SapInstanceNumber": "*string", "Sid": "*string", "Tags": "map[string]string"},
			"start-application":                    {"ApplicationId": "*string"},
			"start-application-refresh":            {"ApplicationId": "*string"},
			"start-configuration-checks":           {"ApplicationId": "*string", "ConfigurationCheckIds": "[]types.ConfigurationCheckType"},
			"stop-application":                     {"ApplicationId": "*string", "IncludeEc2InstanceShutdown": "*bool", "StopConnectedEntity": "types.ConnectedEntityType"},
			"tag-resource":                         {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                       {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-application-settings":          {"ApplicationId": "*string", "Backint": "*types.BackintConfig", "CredentialsToAddOrUpdate": "[]types.ApplicationCredential", "CredentialsToRemove": "[]types.ApplicationCredential", "DatabaseArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-resource-permission":           {"ResourceArn"},
			"deregister-application":               {"ApplicationId"},
			"get-application":                      {},
			"get-component":                        {"ApplicationId", "ComponentId"},
			"get-configuration-check-operation":    {"OperationId"},
			"get-database":                         {},
			"get-operation":                        {"OperationId"},
			"get-resource-permission":              {"ResourceArn"},
			"list-applications":                    {},
			"list-components":                      {},
			"list-configuration-check-definitions": {},
			"list-configuration-check-operations":  {"ApplicationId"},
			"list-databases":                       {},
			"list-operation-events":                {"OperationId"},
			"list-operations":                      {"ApplicationId"},
			"list-sub-check-results":               {"OperationId"},
			"list-sub-check-rule-results":          {"SubCheckResultId"},
			"list-tags-for-resource":               {"ResourceArn"},
			"put-resource-permission":              {"ActionType", "ResourceArn", "SourceResourceArn"},
			"register-application":                 {"ApplicationId", "ApplicationType", "Instances"},
			"start-application":                    {"ApplicationId"},
			"start-application-refresh":            {"ApplicationId"},
			"start-configuration-checks":           {"ApplicationId"},
			"stop-application":                     {"ApplicationId"},
			"tag-resource":                         {"ResourceArn", "Tags"},
			"untag-resource":                       {"ResourceArn", "TagKeys"},
			"update-application-settings":          {"ApplicationId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ssmsap", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
