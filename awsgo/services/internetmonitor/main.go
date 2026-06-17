package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/internetmonitor/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-monitor", "delete-monitor", "get-health-event", "get-internet-event", "get-monitor", "get-query-results", "get-query-status", "list-health-events", "list-internet-events", "list-monitors", "list-tags-for-resource", "start-query", "stop-query", "tag-resource", "untag-resource", "update-monitor"},
		OperationSet: map[string]bool{"create-monitor": true, "delete-monitor": true, "get-health-event": true, "get-internet-event": true, "get-monitor": true, "get-query-results": true, "get-query-status": true, "list-health-events": true, "list-internet-events": true, "list-monitors": true, "list-tags-for-resource": true, "start-query": true, "stop-query": true, "tag-resource": true, "untag-resource": true, "update-monitor": true},
		OperationInputs: map[string][]string{
			"create-monitor":         {"ClientToken", "HealthEventsConfig", "InternetMeasurementsLogDelivery", "MaxCityNetworksToMonitor", "MonitorName", "Resources", "Tags", "TrafficPercentageToMonitor"},
			"delete-monitor":         {"MonitorName"},
			"get-health-event":       {"EventId", "LinkedAccountId", "MonitorName"},
			"get-internet-event":     {"EventId"},
			"get-monitor":            {"LinkedAccountId", "MonitorName"},
			"get-query-results":      {"MaxResults", "MonitorName", "NextToken", "QueryId"},
			"get-query-status":       {"MonitorName", "QueryId"},
			"list-health-events":     {"EndTime", "EventStatus", "LinkedAccountId", "MaxResults", "MonitorName", "NextToken", "StartTime"},
			"list-internet-events":   {"EndTime", "EventStatus", "EventType", "MaxResults", "NextToken", "StartTime"},
			"list-monitors":          {"IncludeLinkedAccounts", "MaxResults", "MonitorStatus", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"start-query":            {"EndTime", "FilterParameters", "LinkedAccountId", "MonitorName", "QueryType", "StartTime"},
			"stop-query":             {"MonitorName", "QueryId"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-monitor":         {"ClientToken", "HealthEventsConfig", "InternetMeasurementsLogDelivery", "MaxCityNetworksToMonitor", "MonitorName", "ResourcesToAdd", "ResourcesToRemove", "Status", "TrafficPercentageToMonitor"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-monitor":         {"ClientToken": "*string", "HealthEventsConfig": "*types.HealthEventsConfig", "InternetMeasurementsLogDelivery": "*types.InternetMeasurementsLogDelivery", "MaxCityNetworksToMonitor": "*int32", "MonitorName": "*string", "Resources": "[]string", "Tags": "map[string]string", "TrafficPercentageToMonitor": "*int32"},
			"delete-monitor":         {"MonitorName": "*string"},
			"get-health-event":       {"EventId": "*string", "LinkedAccountId": "*string", "MonitorName": "*string"},
			"get-internet-event":     {"EventId": "*string"},
			"get-monitor":            {"LinkedAccountId": "*string", "MonitorName": "*string"},
			"get-query-results":      {"MaxResults": "*int32", "MonitorName": "*string", "NextToken": "*string", "QueryId": "*string"},
			"get-query-status":       {"MonitorName": "*string", "QueryId": "*string"},
			"list-health-events":     {"EndTime": "*time.Time", "EventStatus": "types.HealthEventStatus", "LinkedAccountId": "*string", "MaxResults": "*int32", "MonitorName": "*string", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-internet-events":   {"EndTime": "*time.Time", "EventStatus": "*string", "EventType": "*string", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-monitors":          {"IncludeLinkedAccounts": "*bool", "MaxResults": "*int32", "MonitorStatus": "*string", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"start-query":            {"EndTime": "*time.Time", "FilterParameters": "[]types.FilterParameter", "LinkedAccountId": "*string", "MonitorName": "*string", "QueryType": "types.QueryType", "StartTime": "*time.Time"},
			"stop-query":             {"MonitorName": "*string", "QueryId": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-monitor":         {"ClientToken": "*string", "HealthEventsConfig": "*types.HealthEventsConfig", "InternetMeasurementsLogDelivery": "*types.InternetMeasurementsLogDelivery", "MaxCityNetworksToMonitor": "*int32", "MonitorName": "*string", "ResourcesToAdd": "[]string", "ResourcesToRemove": "[]string", "Status": "types.MonitorConfigState", "TrafficPercentageToMonitor": "*int32"},
		},
		OperationInputRequired: map[string][]string{
			"create-monitor":         {"MonitorName"},
			"delete-monitor":         {"MonitorName"},
			"get-health-event":       {"EventId", "MonitorName"},
			"get-internet-event":     {"EventId"},
			"get-monitor":            {"MonitorName"},
			"get-query-results":      {"MonitorName", "QueryId"},
			"get-query-status":       {"MonitorName", "QueryId"},
			"list-health-events":     {"MonitorName"},
			"list-internet-events":   {},
			"list-monitors":          {},
			"list-tags-for-resource": {"ResourceArn"},
			"start-query":            {"EndTime", "MonitorName", "QueryType", "StartTime"},
			"stop-query":             {"MonitorName", "QueryId"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-monitor":         {"MonitorName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("internetmonitor", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
