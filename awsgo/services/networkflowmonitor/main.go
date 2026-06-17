package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/networkflowmonitor/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-monitor", "create-scope", "delete-monitor", "delete-scope", "get-monitor", "get-query-results-monitor-top-contributors", "get-query-results-workload-insights-top-contributors", "get-query-results-workload-insights-top-contributors-data", "get-query-status-monitor-top-contributors", "get-query-status-workload-insights-top-contributors", "get-query-status-workload-insights-top-contributors-data", "get-scope", "list-monitors", "list-scopes", "list-tags-for-resource", "start-query-monitor-top-contributors", "start-query-workload-insights-top-contributors", "start-query-workload-insights-top-contributors-data", "stop-query-monitor-top-contributors", "stop-query-workload-insights-top-contributors", "stop-query-workload-insights-top-contributors-data", "tag-resource", "untag-resource", "update-monitor", "update-scope"},
		OperationSet: map[string]bool{"create-monitor": true, "create-scope": true, "delete-monitor": true, "delete-scope": true, "get-monitor": true, "get-query-results-monitor-top-contributors": true, "get-query-results-workload-insights-top-contributors": true, "get-query-results-workload-insights-top-contributors-data": true, "get-query-status-monitor-top-contributors": true, "get-query-status-workload-insights-top-contributors": true, "get-query-status-workload-insights-top-contributors-data": true, "get-scope": true, "list-monitors": true, "list-scopes": true, "list-tags-for-resource": true, "start-query-monitor-top-contributors": true, "start-query-workload-insights-top-contributors": true, "start-query-workload-insights-top-contributors-data": true, "stop-query-monitor-top-contributors": true, "stop-query-workload-insights-top-contributors": true, "stop-query-workload-insights-top-contributors-data": true, "tag-resource": true, "untag-resource": true, "update-monitor": true, "update-scope": true},
		OperationInputs: map[string][]string{
			"create-monitor": {"ClientToken", "LocalResources", "MonitorName", "RemoteResources", "ScopeArn", "Tags"},
			"create-scope":   {"ClientToken", "Tags", "Targets"},
			"delete-monitor": {"MonitorName"},
			"delete-scope":   {"ScopeId"},
			"get-monitor":    {"MonitorName"},
			"get-query-results-monitor-top-contributors":                {"MaxResults", "MonitorName", "NextToken", "QueryId"},
			"get-query-results-workload-insights-top-contributors":      {"MaxResults", "NextToken", "QueryId", "ScopeId"},
			"get-query-results-workload-insights-top-contributors-data": {"MaxResults", "NextToken", "QueryId", "ScopeId"},
			"get-query-status-monitor-top-contributors":                 {"MonitorName", "QueryId"},
			"get-query-status-workload-insights-top-contributors":       {"QueryId", "ScopeId"},
			"get-query-status-workload-insights-top-contributors-data":  {"QueryId", "ScopeId"},
			"get-scope":                            {"ScopeId"},
			"list-monitors":                        {"MaxResults", "MonitorStatus", "NextToken"},
			"list-scopes":                          {"MaxResults", "NextToken"},
			"list-tags-for-resource":               {"ResourceArn"},
			"start-query-monitor-top-contributors": {"DestinationCategory", "EndTime", "Limit", "MetricName", "MonitorName", "StartTime"},
			"start-query-workload-insights-top-contributors":      {"DestinationCategory", "EndTime", "Limit", "MetricName", "ScopeId", "StartTime"},
			"start-query-workload-insights-top-contributors-data": {"DestinationCategory", "EndTime", "MetricName", "ScopeId", "StartTime"},
			"stop-query-monitor-top-contributors":                 {"MonitorName", "QueryId"},
			"stop-query-workload-insights-top-contributors":       {"QueryId", "ScopeId"},
			"stop-query-workload-insights-top-contributors-data":  {"QueryId", "ScopeId"},
			"tag-resource":   {"ResourceArn", "Tags"},
			"untag-resource": {"ResourceArn", "TagKeys"},
			"update-monitor": {"ClientToken", "LocalResourcesToAdd", "LocalResourcesToRemove", "MonitorName", "RemoteResourcesToAdd", "RemoteResourcesToRemove"},
			"update-scope":   {"ResourcesToAdd", "ResourcesToDelete", "ScopeId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-monitor": {"ClientToken": "*string", "LocalResources": "[]types.MonitorLocalResource", "MonitorName": "*string", "RemoteResources": "[]types.MonitorRemoteResource", "ScopeArn": "*string", "Tags": "map[string]string"},
			"create-scope":   {"ClientToken": "*string", "Tags": "map[string]string", "Targets": "[]types.TargetResource"},
			"delete-monitor": {"MonitorName": "*string"},
			"delete-scope":   {"ScopeId": "*string"},
			"get-monitor":    {"MonitorName": "*string"},
			"get-query-results-monitor-top-contributors":                {"MaxResults": "*int32", "MonitorName": "*string", "NextToken": "*string", "QueryId": "*string"},
			"get-query-results-workload-insights-top-contributors":      {"MaxResults": "*int32", "NextToken": "*string", "QueryId": "*string", "ScopeId": "*string"},
			"get-query-results-workload-insights-top-contributors-data": {"MaxResults": "*int32", "NextToken": "*string", "QueryId": "*string", "ScopeId": "*string"},
			"get-query-status-monitor-top-contributors":                 {"MonitorName": "*string", "QueryId": "*string"},
			"get-query-status-workload-insights-top-contributors":       {"QueryId": "*string", "ScopeId": "*string"},
			"get-query-status-workload-insights-top-contributors-data":  {"QueryId": "*string", "ScopeId": "*string"},
			"get-scope":                            {"ScopeId": "*string"},
			"list-monitors":                        {"MaxResults": "*int32", "MonitorStatus": "types.MonitorStatus", "NextToken": "*string"},
			"list-scopes":                          {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":               {"ResourceArn": "*string"},
			"start-query-monitor-top-contributors": {"DestinationCategory": "types.DestinationCategory", "EndTime": "*time.Time", "Limit": "*int32", "MetricName": "types.MonitorMetric", "MonitorName": "*string", "StartTime": "*time.Time"},
			"start-query-workload-insights-top-contributors":      {"DestinationCategory": "types.DestinationCategory", "EndTime": "*time.Time", "Limit": "*int32", "MetricName": "types.WorkloadInsightsMetric", "ScopeId": "*string", "StartTime": "*time.Time"},
			"start-query-workload-insights-top-contributors-data": {"DestinationCategory": "types.DestinationCategory", "EndTime": "*time.Time", "MetricName": "types.WorkloadInsightsMetric", "ScopeId": "*string", "StartTime": "*time.Time"},
			"stop-query-monitor-top-contributors":                 {"MonitorName": "*string", "QueryId": "*string"},
			"stop-query-workload-insights-top-contributors":       {"QueryId": "*string", "ScopeId": "*string"},
			"stop-query-workload-insights-top-contributors-data":  {"QueryId": "*string", "ScopeId": "*string"},
			"tag-resource":   {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource": {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-monitor": {"ClientToken": "*string", "LocalResourcesToAdd": "[]types.MonitorLocalResource", "LocalResourcesToRemove": "[]types.MonitorLocalResource", "MonitorName": "*string", "RemoteResourcesToAdd": "[]types.MonitorRemoteResource", "RemoteResourcesToRemove": "[]types.MonitorRemoteResource"},
			"update-scope":   {"ResourcesToAdd": "[]types.TargetResource", "ResourcesToDelete": "[]types.TargetResource", "ScopeId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-monitor": {"LocalResources", "MonitorName", "ScopeArn"},
			"create-scope":   {"Targets"},
			"delete-monitor": {"MonitorName"},
			"delete-scope":   {"ScopeId"},
			"get-monitor":    {"MonitorName"},
			"get-query-results-monitor-top-contributors":                {"MonitorName", "QueryId"},
			"get-query-results-workload-insights-top-contributors":      {"QueryId", "ScopeId"},
			"get-query-results-workload-insights-top-contributors-data": {"QueryId", "ScopeId"},
			"get-query-status-monitor-top-contributors":                 {"MonitorName", "QueryId"},
			"get-query-status-workload-insights-top-contributors":       {"QueryId", "ScopeId"},
			"get-query-status-workload-insights-top-contributors-data":  {"QueryId", "ScopeId"},
			"get-scope":                            {"ScopeId"},
			"list-monitors":                        {},
			"list-scopes":                          {},
			"list-tags-for-resource":               {"ResourceArn"},
			"start-query-monitor-top-contributors": {"DestinationCategory", "EndTime", "MetricName", "MonitorName", "StartTime"},
			"start-query-workload-insights-top-contributors":      {"DestinationCategory", "EndTime", "MetricName", "ScopeId", "StartTime"},
			"start-query-workload-insights-top-contributors-data": {"DestinationCategory", "EndTime", "MetricName", "ScopeId", "StartTime"},
			"stop-query-monitor-top-contributors":                 {"MonitorName", "QueryId"},
			"stop-query-workload-insights-top-contributors":       {"QueryId", "ScopeId"},
			"stop-query-workload-insights-top-contributors-data":  {"QueryId", "ScopeId"},
			"tag-resource":   {"ResourceArn", "Tags"},
			"untag-resource": {"ResourceArn", "TagKeys"},
			"update-monitor": {"MonitorName"},
			"update-scope":   {"ScopeId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("networkflowmonitor", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
