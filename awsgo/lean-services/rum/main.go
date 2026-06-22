package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/rum"
)

var fields_batch_create_rum_metric_definitions = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "types.MetricDestination", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "MetricDefinitions", Flag: "metric-definitions", Type: "[]types.MetricDefinitionRequest", Required: true},
}

var fields_batch_delete_rum_metric_definitions = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "types.MetricDestination", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "MetricDefinitionIds", Flag: "metric-definition-ids", Type: "[]string", Required: true},
}

var fields_batch_get_rum_metric_definitions = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "types.MetricDestination", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_create_app_monitor = []leanruntime.Field{
	{Name: "AppMonitorConfiguration", Flag: "app-monitor-configuration", Type: "*types.AppMonitorConfiguration", Required: false},
	{Name: "CustomEvents", Flag: "custom-events", Type: "*types.CustomEvents", Required: false},
	{Name: "CwLogEnabled", Flag: "cw-log-enabled", Type: "*bool", Required: false},
	{Name: "DeobfuscationConfiguration", Flag: "deobfuscation-configuration", Type: "*types.DeobfuscationConfiguration", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainList", Flag: "domain-list", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.AppMonitorPlatform", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_app_monitor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
}

var fields_delete_rum_metrics_destination = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "types.MetricDestination", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
}

var fields_get_app_monitor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_app_monitor_data = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.QueryFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TimeRange", Flag: "time-range", Type: "*types.TimeRange", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_app_monitors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rum_metrics_destinations = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
}

var fields_put_rum_events = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "AppMonitorDetails", Flag: "app-monitor-details", Type: "*types.AppMonitorDetails", Required: true},
	{Name: "BatchId", Flag: "batch-id", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RumEvents", Flag: "rum-events", Type: "[]types.RumEvent", Required: true},
	{Name: "UserDetails", Flag: "user-details", Type: "*types.UserDetails", Required: true},
}

var fields_put_rum_metrics_destination = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "types.MetricDestination", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_app_monitor = []leanruntime.Field{
	{Name: "AppMonitorConfiguration", Flag: "app-monitor-configuration", Type: "*types.AppMonitorConfiguration", Required: false},
	{Name: "CustomEvents", Flag: "custom-events", Type: "*types.CustomEvents", Required: false},
	{Name: "CwLogEnabled", Flag: "cw-log-enabled", Type: "*bool", Required: false},
	{Name: "DeobfuscationConfiguration", Flag: "deobfuscation-configuration", Type: "*types.DeobfuscationConfiguration", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainList", Flag: "domain-list", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_rum_metric_definition = []leanruntime.Field{
	{Name: "AppMonitorName", Flag: "app-monitor-name", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "types.MetricDestination", Required: true},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "MetricDefinition", Flag: "metric-definition", Type: "*types.MetricDefinitionRequest", Required: true},
	{Name: "MetricDefinitionId", Flag: "metric-definition-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-rum-metric-definitions": {
			Name:   "batch-create-rum-metric-definitions",
			Fields: fields_batch_create_rum_metric_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateRumMetricDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_rum_metric_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateRumMetricDefinitions(ctx, input)
			},
		},
		"batch-delete-rum-metric-definitions": {
			Name:   "batch-delete-rum-metric-definitions",
			Fields: fields_batch_delete_rum_metric_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteRumMetricDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_rum_metric_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteRumMetricDefinitions(ctx, input)
			},
		},
		"batch-get-rum-metric-definitions": {
			Name:   "batch-get-rum-metric-definitions",
			Fields: fields_batch_get_rum_metric_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRumMetricDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_batch_get_rum_metric_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.BatchGetRumMetricDefinitions(ctx, input)
				}
				var results []*svc.BatchGetRumMetricDefinitionsOutput
				p := svc.NewBatchGetRumMetricDefinitionsPaginator(client, input)
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
		"create-app-monitor": {
			Name:   "create-app-monitor",
			Fields: fields_create_app_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppMonitor(ctx, input)
			},
		},
		"delete-app-monitor": {
			Name:   "delete-app-monitor",
			Fields: fields_delete_app_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppMonitor(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-rum-metrics-destination": {
			Name:   "delete-rum-metrics-destination",
			Fields: fields_delete_rum_metrics_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRumMetricsDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rum_metrics_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRumMetricsDestination(ctx, input)
			},
		},
		"get-app-monitor": {
			Name:   "get-app-monitor",
			Fields: fields_get_app_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_app_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAppMonitor(ctx, input)
			},
		},
		"get-app-monitor-data": {
			Name:   "get-app-monitor-data",
			Fields: fields_get_app_monitor_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppMonitorDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_app_monitor_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAppMonitorData(ctx, input)
				}
				var results []*svc.GetAppMonitorDataOutput
				p := svc.NewGetAppMonitorDataPaginator(client, input)
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
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"list-app-monitors": {
			Name:   "list-app-monitors",
			Fields: fields_list_app_monitors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppMonitorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_monitors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppMonitors(ctx, input)
				}
				var results []*svc.ListAppMonitorsOutput
				p := svc.NewListAppMonitorsPaginator(client, input)
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
		"list-rum-metrics-destinations": {
			Name:   "list-rum-metrics-destinations",
			Fields: fields_list_rum_metrics_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRumMetricsDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rum_metrics_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRumMetricsDestinations(ctx, input)
				}
				var results []*svc.ListRumMetricsDestinationsOutput
				p := svc.NewListRumMetricsDestinationsPaginator(client, input)
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
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"put-rum-events": {
			Name:   "put-rum-events",
			Fields: fields_put_rum_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRumEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_rum_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRumEvents(ctx, input)
			},
		},
		"put-rum-metrics-destination": {
			Name:   "put-rum-metrics-destination",
			Fields: fields_put_rum_metrics_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRumMetricsDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_rum_metrics_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRumMetricsDestination(ctx, input)
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
		"update-app-monitor": {
			Name:   "update-app-monitor",
			Fields: fields_update_app_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppMonitor(ctx, input)
			},
		},
		"update-rum-metric-definition": {
			Name:   "update-rum-metric-definition",
			Fields: fields_update_rum_metric_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRumMetricDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rum_metric_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRumMetricDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("rum", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
