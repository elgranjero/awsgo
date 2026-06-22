package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/networkflowmonitor"
)

var fields_create_monitor = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LocalResources", Flag: "local-resources", Type: "[]types.MonitorLocalResource", Required: true},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "RemoteResources", Flag: "remote-resources", Type: "[]types.MonitorRemoteResource", Required: false},
	{Name: "ScopeArn", Flag: "scope-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_scope = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.TargetResource", Required: true},
}

var fields_delete_monitor = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_delete_scope = []leanruntime.Field{
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_get_monitor = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_get_query_results_monitor_top_contributors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_query_results_workload_insights_top_contributors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_get_query_results_workload_insights_top_contributors_data = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_get_query_status_monitor_top_contributors = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_query_status_workload_insights_top_contributors = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_get_query_status_workload_insights_top_contributors_data = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_get_scope = []leanruntime.Field{
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_list_monitors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorStatus", Flag: "monitor-status", Type: "types.MonitorStatus", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_scopes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_query_monitor_top_contributors = []leanruntime.Field{
	{Name: "DestinationCategory", Flag: "destination-category", Type: "types.DestinationCategory", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "types.MonitorMetric", Required: true},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_start_query_workload_insights_top_contributors = []leanruntime.Field{
	{Name: "DestinationCategory", Flag: "destination-category", Type: "types.DestinationCategory", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "types.WorkloadInsightsMetric", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_start_query_workload_insights_top_contributors_data = []leanruntime.Field{
	{Name: "DestinationCategory", Flag: "destination-category", Type: "types.DestinationCategory", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.WorkloadInsightsMetric", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_stop_query_monitor_top_contributors = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_stop_query_workload_insights_top_contributors = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_stop_query_workload_insights_top_contributors_data = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_monitor = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LocalResourcesToAdd", Flag: "local-resources-to-add", Type: "[]types.MonitorLocalResource", Required: false},
	{Name: "LocalResourcesToRemove", Flag: "local-resources-to-remove", Type: "[]types.MonitorLocalResource", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "RemoteResourcesToAdd", Flag: "remote-resources-to-add", Type: "[]types.MonitorRemoteResource", Required: false},
	{Name: "RemoteResourcesToRemove", Flag: "remote-resources-to-remove", Type: "[]types.MonitorRemoteResource", Required: false},
}

var fields_update_scope = []leanruntime.Field{
	{Name: "ResourcesToAdd", Flag: "resources-to-add", Type: "[]types.TargetResource", Required: false},
	{Name: "ResourcesToDelete", Flag: "resources-to-delete", Type: "[]types.TargetResource", Required: false},
	{Name: "ScopeId", Flag: "scope-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-monitor": {
			Name:   "create-monitor",
			Fields: fields_create_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMonitor(ctx, input)
			},
		},
		"create-scope": {
			Name:   "create-scope",
			Fields: fields_create_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScope(ctx, input)
			},
		},
		"delete-monitor": {
			Name:   "delete-monitor",
			Fields: fields_delete_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMonitor(ctx, input)
			},
		},
		"delete-scope": {
			Name:   "delete-scope",
			Fields: fields_delete_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScope(ctx, input)
			},
		},
		"get-monitor": {
			Name:   "get-monitor",
			Fields: fields_get_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMonitor(ctx, input)
			},
		},
		"get-query-results-monitor-top-contributors": {
			Name:   "get-query-results-monitor-top-contributors",
			Fields: fields_get_query_results_monitor_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryResultsMonitorTopContributorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_query_results_monitor_top_contributors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetQueryResultsMonitorTopContributors(ctx, input)
				}
				var results []*svc.GetQueryResultsMonitorTopContributorsOutput
				p := svc.NewGetQueryResultsMonitorTopContributorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-query-results-workload-insights-top-contributors": {
			Name:   "get-query-results-workload-insights-top-contributors",
			Fields: fields_get_query_results_workload_insights_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryResultsWorkloadInsightsTopContributorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_query_results_workload_insights_top_contributors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetQueryResultsWorkloadInsightsTopContributors(ctx, input)
				}
				var results []*svc.GetQueryResultsWorkloadInsightsTopContributorsOutput
				p := svc.NewGetQueryResultsWorkloadInsightsTopContributorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-query-results-workload-insights-top-contributors-data": {
			Name:   "get-query-results-workload-insights-top-contributors-data",
			Fields: fields_get_query_results_workload_insights_top_contributors_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryResultsWorkloadInsightsTopContributorsDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_query_results_workload_insights_top_contributors_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetQueryResultsWorkloadInsightsTopContributorsData(ctx, input)
				}
				var results []*svc.GetQueryResultsWorkloadInsightsTopContributorsDataOutput
				p := svc.NewGetQueryResultsWorkloadInsightsTopContributorsDataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-query-status-monitor-top-contributors": {
			Name:   "get-query-status-monitor-top-contributors",
			Fields: fields_get_query_status_monitor_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryStatusMonitorTopContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_status_monitor_top_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryStatusMonitorTopContributors(ctx, input)
			},
		},
		"get-query-status-workload-insights-top-contributors": {
			Name:   "get-query-status-workload-insights-top-contributors",
			Fields: fields_get_query_status_workload_insights_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryStatusWorkloadInsightsTopContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_status_workload_insights_top_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryStatusWorkloadInsightsTopContributors(ctx, input)
			},
		},
		"get-query-status-workload-insights-top-contributors-data": {
			Name:   "get-query-status-workload-insights-top-contributors-data",
			Fields: fields_get_query_status_workload_insights_top_contributors_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryStatusWorkloadInsightsTopContributorsDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_status_workload_insights_top_contributors_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryStatusWorkloadInsightsTopContributorsData(ctx, input)
			},
		},
		"get-scope": {
			Name:   "get-scope",
			Fields: fields_get_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScope(ctx, input)
			},
		},
		"list-monitors": {
			Name:   "list-monitors",
			Fields: fields_list_monitors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitors(ctx, input)
				}
				var results []*svc.ListMonitorsOutput
				p := svc.NewListMonitorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-scopes": {
			Name:   "list-scopes",
			Fields: fields_list_scopes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScopesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scopes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScopes(ctx, input)
				}
				var results []*svc.ListScopesOutput
				p := svc.NewListScopesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"start-query-monitor-top-contributors": {
			Name:   "start-query-monitor-top-contributors",
			Fields: fields_start_query_monitor_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQueryMonitorTopContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_query_monitor_top_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQueryMonitorTopContributors(ctx, input)
			},
		},
		"start-query-workload-insights-top-contributors": {
			Name:   "start-query-workload-insights-top-contributors",
			Fields: fields_start_query_workload_insights_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQueryWorkloadInsightsTopContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_query_workload_insights_top_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQueryWorkloadInsightsTopContributors(ctx, input)
			},
		},
		"start-query-workload-insights-top-contributors-data": {
			Name:   "start-query-workload-insights-top-contributors-data",
			Fields: fields_start_query_workload_insights_top_contributors_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQueryWorkloadInsightsTopContributorsDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_query_workload_insights_top_contributors_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQueryWorkloadInsightsTopContributorsData(ctx, input)
			},
		},
		"stop-query-monitor-top-contributors": {
			Name:   "stop-query-monitor-top-contributors",
			Fields: fields_stop_query_monitor_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopQueryMonitorTopContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_query_monitor_top_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopQueryMonitorTopContributors(ctx, input)
			},
		},
		"stop-query-workload-insights-top-contributors": {
			Name:   "stop-query-workload-insights-top-contributors",
			Fields: fields_stop_query_workload_insights_top_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopQueryWorkloadInsightsTopContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_query_workload_insights_top_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopQueryWorkloadInsightsTopContributors(ctx, input)
			},
		},
		"stop-query-workload-insights-top-contributors-data": {
			Name:   "stop-query-workload-insights-top-contributors-data",
			Fields: fields_stop_query_workload_insights_top_contributors_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopQueryWorkloadInsightsTopContributorsDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_query_workload_insights_top_contributors_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopQueryWorkloadInsightsTopContributorsData(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-monitor": {
			Name:   "update-monitor",
			Fields: fields_update_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMonitor(ctx, input)
			},
		},
		"update-scope": {
			Name:   "update-scope",
			Fields: fields_update_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScope(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("networkflowmonitor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
