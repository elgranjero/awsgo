package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/rum/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-create-rum-metric-definitions", "batch-delete-rum-metric-definitions", "batch-get-rum-metric-definitions", "create-app-monitor", "delete-app-monitor", "delete-resource-policy", "delete-rum-metrics-destination", "get-app-monitor", "get-app-monitor-data", "get-resource-policy", "list-app-monitors", "list-rum-metrics-destinations", "list-tags-for-resource", "put-resource-policy", "put-rum-events", "put-rum-metrics-destination", "tag-resource", "untag-resource", "update-app-monitor", "update-rum-metric-definition"},
		OperationSet: map[string]bool{"batch-create-rum-metric-definitions": true, "batch-delete-rum-metric-definitions": true, "batch-get-rum-metric-definitions": true, "create-app-monitor": true, "delete-app-monitor": true, "delete-resource-policy": true, "delete-rum-metrics-destination": true, "get-app-monitor": true, "get-app-monitor-data": true, "get-resource-policy": true, "list-app-monitors": true, "list-rum-metrics-destinations": true, "list-tags-for-resource": true, "put-resource-policy": true, "put-rum-events": true, "put-rum-metrics-destination": true, "tag-resource": true, "untag-resource": true, "update-app-monitor": true, "update-rum-metric-definition": true},
		OperationInputs: map[string][]string{
			"batch-create-rum-metric-definitions": {"AppMonitorName", "Destination", "DestinationArn", "MetricDefinitions"},
			"batch-delete-rum-metric-definitions": {"AppMonitorName", "Destination", "DestinationArn", "MetricDefinitionIds"},
			"batch-get-rum-metric-definitions":    {"AppMonitorName", "Destination", "DestinationArn", "MaxResults", "NextToken"},
			"create-app-monitor":                  {"AppMonitorConfiguration", "CustomEvents", "CwLogEnabled", "DeobfuscationConfiguration", "Domain", "DomainList", "Name", "Platform", "Tags"},
			"delete-app-monitor":                  {"Name"},
			"delete-resource-policy":              {"Name", "PolicyRevisionId"},
			"delete-rum-metrics-destination":      {"AppMonitorName", "Destination", "DestinationArn"},
			"get-app-monitor":                     {"Name"},
			"get-app-monitor-data":                {"Filters", "MaxResults", "Name", "NextToken", "TimeRange"},
			"get-resource-policy":                 {"Name"},
			"list-app-monitors":                   {"MaxResults", "NextToken"},
			"list-rum-metrics-destinations":       {"AppMonitorName", "MaxResults", "NextToken"},
			"list-tags-for-resource":              {"ResourceArn"},
			"put-resource-policy":                 {"Name", "PolicyDocument", "PolicyRevisionId"},
			"put-rum-events":                      {"Alias", "AppMonitorDetails", "BatchId", "Id", "RumEvents", "UserDetails"},
			"put-rum-metrics-destination":         {"AppMonitorName", "Destination", "DestinationArn", "IamRoleArn"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-app-monitor":                  {"AppMonitorConfiguration", "CustomEvents", "CwLogEnabled", "DeobfuscationConfiguration", "Domain", "DomainList", "Name"},
			"update-rum-metric-definition":        {"AppMonitorName", "Destination", "DestinationArn", "MetricDefinition", "MetricDefinitionId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-create-rum-metric-definitions": {"AppMonitorName": "*string", "Destination": "types.MetricDestination", "DestinationArn": "*string", "MetricDefinitions": "[]types.MetricDefinitionRequest"},
			"batch-delete-rum-metric-definitions": {"AppMonitorName": "*string", "Destination": "types.MetricDestination", "DestinationArn": "*string", "MetricDefinitionIds": "[]string"},
			"batch-get-rum-metric-definitions":    {"AppMonitorName": "*string", "Destination": "types.MetricDestination", "DestinationArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"create-app-monitor":                  {"AppMonitorConfiguration": "*types.AppMonitorConfiguration", "CustomEvents": "*types.CustomEvents", "CwLogEnabled": "*bool", "DeobfuscationConfiguration": "*types.DeobfuscationConfiguration", "Domain": "*string", "DomainList": "[]string", "Name": "*string", "Platform": "types.AppMonitorPlatform", "Tags": "map[string]string"},
			"delete-app-monitor":                  {"Name": "*string"},
			"delete-resource-policy":              {"Name": "*string", "PolicyRevisionId": "*string"},
			"delete-rum-metrics-destination":      {"AppMonitorName": "*string", "Destination": "types.MetricDestination", "DestinationArn": "*string"},
			"get-app-monitor":                     {"Name": "*string"},
			"get-app-monitor-data":                {"Filters": "[]types.QueryFilter", "MaxResults": "int32", "Name": "*string", "NextToken": "*string", "TimeRange": "*types.TimeRange"},
			"get-resource-policy":                 {"Name": "*string"},
			"list-app-monitors":                   {"MaxResults": "*int32", "NextToken": "*string"},
			"list-rum-metrics-destinations":       {"AppMonitorName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":              {"ResourceArn": "*string"},
			"put-resource-policy":                 {"Name": "*string", "PolicyDocument": "*string", "PolicyRevisionId": "*string"},
			"put-rum-events":                      {"Alias": "*string", "AppMonitorDetails": "*types.AppMonitorDetails", "BatchId": "*string", "Id": "*string", "RumEvents": "[]types.RumEvent", "UserDetails": "*types.UserDetails"},
			"put-rum-metrics-destination":         {"AppMonitorName": "*string", "Destination": "types.MetricDestination", "DestinationArn": "*string", "IamRoleArn": "*string"},
			"tag-resource":                        {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                      {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-app-monitor":                  {"AppMonitorConfiguration": "*types.AppMonitorConfiguration", "CustomEvents": "*types.CustomEvents", "CwLogEnabled": "*bool", "DeobfuscationConfiguration": "*types.DeobfuscationConfiguration", "Domain": "*string", "DomainList": "[]string", "Name": "*string"},
			"update-rum-metric-definition":        {"AppMonitorName": "*string", "Destination": "types.MetricDestination", "DestinationArn": "*string", "MetricDefinition": "*types.MetricDefinitionRequest", "MetricDefinitionId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-create-rum-metric-definitions": {"AppMonitorName", "Destination", "MetricDefinitions"},
			"batch-delete-rum-metric-definitions": {"AppMonitorName", "Destination", "MetricDefinitionIds"},
			"batch-get-rum-metric-definitions":    {"AppMonitorName", "Destination"},
			"create-app-monitor":                  {"Name"},
			"delete-app-monitor":                  {"Name"},
			"delete-resource-policy":              {"Name"},
			"delete-rum-metrics-destination":      {"AppMonitorName", "Destination"},
			"get-app-monitor":                     {"Name"},
			"get-app-monitor-data":                {"Name", "TimeRange"},
			"get-resource-policy":                 {"Name"},
			"list-app-monitors":                   {},
			"list-rum-metrics-destinations":       {"AppMonitorName"},
			"list-tags-for-resource":              {"ResourceArn"},
			"put-resource-policy":                 {"Name", "PolicyDocument"},
			"put-rum-events":                      {"AppMonitorDetails", "BatchId", "Id", "RumEvents", "UserDetails"},
			"put-rum-metrics-destination":         {"AppMonitorName", "Destination"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-app-monitor":                  {"Name"},
			"update-rum-metric-definition":        {"AppMonitorName", "Destination", "MetricDefinition", "MetricDefinitionId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("rum", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
