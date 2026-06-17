package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/appintegrations/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-application", "create-data-integration", "create-data-integration-association", "create-event-integration", "delete-application", "delete-data-integration", "delete-event-integration", "get-application", "get-data-integration", "get-event-integration", "list-application-associations", "list-applications", "list-data-integration-associations", "list-data-integrations", "list-event-integration-associations", "list-event-integrations", "list-tags-for-resource", "tag-resource", "untag-resource", "update-application", "update-data-integration", "update-data-integration-association", "update-event-integration"},
		OperationSet: map[string]bool{"create-application": true, "create-data-integration": true, "create-data-integration-association": true, "create-event-integration": true, "delete-application": true, "delete-data-integration": true, "delete-event-integration": true, "get-application": true, "get-data-integration": true, "get-event-integration": true, "list-application-associations": true, "list-applications": true, "list-data-integration-associations": true, "list-data-integrations": true, "list-event-integration-associations": true, "list-event-integrations": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-application": true, "update-data-integration": true, "update-data-integration-association": true, "update-event-integration": true},
		OperationInputs: map[string][]string{
			"create-application":                  {"ApplicationConfig", "ApplicationSourceConfig", "ApplicationType", "ClientToken", "Description", "IframeConfig", "InitializationTimeout", "IsService", "Name", "Namespace", "Permissions", "Publications", "Subscriptions", "Tags"},
			"create-data-integration":             {"ClientToken", "Description", "FileConfiguration", "KmsKey", "Name", "ObjectConfiguration", "ScheduleConfig", "SourceURI", "Tags"},
			"create-data-integration-association": {"ClientAssociationMetadata", "ClientId", "ClientToken", "DataIntegrationIdentifier", "DestinationURI", "ExecutionConfiguration", "ObjectConfiguration"},
			"create-event-integration":            {"ClientToken", "Description", "EventBridgeBus", "EventFilter", "Name", "Tags"},
			"delete-application":                  {"Arn"},
			"delete-data-integration":             {"DataIntegrationIdentifier"},
			"delete-event-integration":            {"Name"},
			"get-application":                     {"Arn"},
			"get-data-integration":                {"Identifier"},
			"get-event-integration":               {"Name"},
			"list-application-associations":       {"ApplicationId", "MaxResults", "NextToken"},
			"list-applications":                   {"ApplicationType", "MaxResults", "NextToken"},
			"list-data-integration-associations":  {"DataIntegrationIdentifier", "MaxResults", "NextToken"},
			"list-data-integrations":              {"MaxResults", "NextToken"},
			"list-event-integration-associations": {"EventIntegrationName", "MaxResults", "NextToken"},
			"list-event-integrations":             {"MaxResults", "NextToken"},
			"list-tags-for-resource":              {"ResourceArn"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-application":                  {"ApplicationConfig", "ApplicationSourceConfig", "ApplicationType", "Arn", "Description", "IframeConfig", "InitializationTimeout", "IsService", "Name", "Permissions", "Publications", "Subscriptions"},
			"update-data-integration":             {"Description", "Identifier", "Name"},
			"update-data-integration-association": {"DataIntegrationAssociationIdentifier", "DataIntegrationIdentifier", "ExecutionConfiguration"},
			"update-event-integration":            {"Description", "Name"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-application":                  {"ApplicationConfig": "*types.ApplicationConfig", "ApplicationSourceConfig": "*types.ApplicationSourceConfig", "ApplicationType": "types.ApplicationType", "ClientToken": "*string", "Description": "*string", "IframeConfig": "*types.IframeConfig", "InitializationTimeout": "*int32", "IsService": "bool", "Name": "*string", "Namespace": "*string", "Permissions": "[]string", "Publications": "[]types.Publication", "Subscriptions": "[]types.Subscription", "Tags": "map[string]string"},
			"create-data-integration":             {"ClientToken": "*string", "Description": "*string", "FileConfiguration": "*types.FileConfiguration", "KmsKey": "*string", "Name": "*string", "ObjectConfiguration": "map[string]map[string][]string", "ScheduleConfig": "*types.ScheduleConfiguration", "SourceURI": "*string", "Tags": "map[string]string"},
			"create-data-integration-association": {"ClientAssociationMetadata": "map[string]string", "ClientId": "*string", "ClientToken": "*string", "DataIntegrationIdentifier": "*string", "DestinationURI": "*string", "ExecutionConfiguration": "*types.ExecutionConfiguration", "ObjectConfiguration": "map[string]map[string][]string"},
			"create-event-integration":            {"ClientToken": "*string", "Description": "*string", "EventBridgeBus": "*string", "EventFilter": "*types.EventFilter", "Name": "*string", "Tags": "map[string]string"},
			"delete-application":                  {"Arn": "*string"},
			"delete-data-integration":             {"DataIntegrationIdentifier": "*string"},
			"delete-event-integration":            {"Name": "*string"},
			"get-application":                     {"Arn": "*string"},
			"get-data-integration":                {"Identifier": "*string"},
			"get-event-integration":               {"Name": "*string"},
			"list-application-associations":       {"ApplicationId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-applications":                   {"ApplicationType": "types.ApplicationType", "MaxResults": "*int32", "NextToken": "*string"},
			"list-data-integration-associations":  {"DataIntegrationIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-data-integrations":              {"MaxResults": "*int32", "NextToken": "*string"},
			"list-event-integration-associations": {"EventIntegrationName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-event-integrations":             {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":              {"ResourceArn": "*string"},
			"tag-resource":                        {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                      {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-application":                  {"ApplicationConfig": "*types.ApplicationConfig", "ApplicationSourceConfig": "*types.ApplicationSourceConfig", "ApplicationType": "types.ApplicationType", "Arn": "*string", "Description": "*string", "IframeConfig": "*types.IframeConfig", "InitializationTimeout": "*int32", "IsService": "*bool", "Name": "*string", "Permissions": "[]string", "Publications": "[]types.Publication", "Subscriptions": "[]types.Subscription"},
			"update-data-integration":             {"Description": "*string", "Identifier": "*string", "Name": "*string"},
			"update-data-integration-association": {"DataIntegrationAssociationIdentifier": "*string", "DataIntegrationIdentifier": "*string", "ExecutionConfiguration": "*types.ExecutionConfiguration"},
			"update-event-integration":            {"Description": "*string", "Name": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-application":                  {"ApplicationSourceConfig", "Name", "Namespace"},
			"create-data-integration":             {"KmsKey", "Name"},
			"create-data-integration-association": {"DataIntegrationIdentifier"},
			"create-event-integration":            {"EventBridgeBus", "EventFilter", "Name"},
			"delete-application":                  {"Arn"},
			"delete-data-integration":             {"DataIntegrationIdentifier"},
			"delete-event-integration":            {"Name"},
			"get-application":                     {"Arn"},
			"get-data-integration":                {"Identifier"},
			"get-event-integration":               {"Name"},
			"list-application-associations":       {"ApplicationId"},
			"list-applications":                   {},
			"list-data-integration-associations":  {"DataIntegrationIdentifier"},
			"list-data-integrations":              {},
			"list-event-integration-associations": {"EventIntegrationName"},
			"list-event-integrations":             {},
			"list-tags-for-resource":              {"ResourceArn"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-application":                  {"Arn"},
			"update-data-integration":             {"Identifier"},
			"update-data-integration-association": {"DataIntegrationAssociationIdentifier", "DataIntegrationIdentifier", "ExecutionConfiguration"},
			"update-event-integration":            {"Name"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("appintegrations", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
