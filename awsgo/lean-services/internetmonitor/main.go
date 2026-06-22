package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/internetmonitor"
)

var fields_create_monitor = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "HealthEventsConfig", Flag: "health-events-config", Type: "*types.HealthEventsConfig", Required: false},
	{Name: "InternetMeasurementsLogDelivery", Flag: "internet-measurements-log-delivery", Type: "*types.InternetMeasurementsLogDelivery", Required: false},
	{Name: "MaxCityNetworksToMonitor", Flag: "max-city-networks-to-monitor", Type: "*int32", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "Resources", Flag: "resources", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrafficPercentageToMonitor", Flag: "traffic-percentage-to-monitor", Type: "*int32", Required: false},
}

var fields_delete_monitor = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_get_health_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "LinkedAccountId", Flag: "linked-account-id", Type: "*string", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_get_internet_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
}

var fields_get_monitor = []leanruntime.Field{
	{Name: "LinkedAccountId", Flag: "linked-account-id", Type: "*string", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_get_query_results = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_query_status = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_list_health_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EventStatus", Flag: "event-status", Type: "types.HealthEventStatus", Required: false},
	{Name: "LinkedAccountId", Flag: "linked-account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_internet_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EventStatus", Flag: "event-status", Type: "*string", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_monitors = []leanruntime.Field{
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorStatus", Flag: "monitor-status", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_query = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "FilterParameters", Flag: "filter-parameters", Type: "[]types.FilterParameter", Required: false},
	{Name: "LinkedAccountId", Flag: "linked-account-id", Type: "*string", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "QueryType", Flag: "query-type", Type: "types.QueryType", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_stop_query = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
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
	{Name: "HealthEventsConfig", Flag: "health-events-config", Type: "*types.HealthEventsConfig", Required: false},
	{Name: "InternetMeasurementsLogDelivery", Flag: "internet-measurements-log-delivery", Type: "*types.InternetMeasurementsLogDelivery", Required: false},
	{Name: "MaxCityNetworksToMonitor", Flag: "max-city-networks-to-monitor", Type: "*int32", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "ResourcesToAdd", Flag: "resources-to-add", Type: "[]string", Required: false},
	{Name: "ResourcesToRemove", Flag: "resources-to-remove", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MonitorConfigState", Required: false},
	{Name: "TrafficPercentageToMonitor", Flag: "traffic-percentage-to-monitor", Type: "*int32", Required: false},
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
		"get-health-event": {
			Name:   "get-health-event",
			Fields: fields_get_health_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHealthEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_health_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHealthEvent(ctx, input)
			},
		},
		"get-internet-event": {
			Name:   "get-internet-event",
			Fields: fields_get_internet_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInternetEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_internet_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInternetEvent(ctx, input)
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
		"get-query-results": {
			Name:   "get-query-results",
			Fields: fields_get_query_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_query_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetQueryResults(ctx, input)
				}
				var results []*svc.GetQueryResultsOutput
				p := svc.NewGetQueryResultsPaginator(client, input)
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
		"get-query-status": {
			Name:   "get-query-status",
			Fields: fields_get_query_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryStatus(ctx, input)
			},
		},
		"list-health-events": {
			Name:   "list-health-events",
			Fields: fields_list_health_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHealthEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_health_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHealthEvents(ctx, input)
				}
				var results []*svc.ListHealthEventsOutput
				p := svc.NewListHealthEventsPaginator(client, input)
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
		"list-internet-events": {
			Name:   "list-internet-events",
			Fields: fields_list_internet_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInternetEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_internet_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInternetEvents(ctx, input)
				}
				var results []*svc.ListInternetEventsOutput
				p := svc.NewListInternetEventsPaginator(client, input)
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
		"start-query": {
			Name:   "start-query",
			Fields: fields_start_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQuery(ctx, input)
			},
		},
		"stop-query": {
			Name:   "stop-query",
			Fields: fields_stop_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopQuery(ctx, input)
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
	}
	if err := leanruntime.Execute("internetmonitor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
