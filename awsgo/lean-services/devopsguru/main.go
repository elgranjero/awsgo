package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/devopsguru"
)

var fields_add_notification_channel = []leanruntime.Field{
	{Name: "Config", Flag: "config", Type: "*types.NotificationChannelConfig", Required: true},
}

var fields_delete_insight = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_account_health = []leanruntime.Field{}

var fields_describe_account_overview = []leanruntime.Field{
	{Name: "FromTime", Flag: "from-time", Type: "*time.Time", Required: true},
	{Name: "ToTime", Flag: "to-time", Type: "*time.Time", Required: false},
}

var fields_describe_anomaly = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_event_sources_config = []leanruntime.Field{}

var fields_describe_feedback = []leanruntime.Field{
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: false},
}

var fields_describe_insight = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_organization_health = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "OrganizationalUnitIds", Flag: "organizational-unit-ids", Type: "[]string", Required: false},
}

var fields_describe_organization_overview = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FromTime", Flag: "from-time", Type: "*time.Time", Required: true},
	{Name: "OrganizationalUnitIds", Flag: "organizational-unit-ids", Type: "[]string", Required: false},
	{Name: "ToTime", Flag: "to-time", Type: "*time.Time", Required: false},
}

var fields_describe_organization_resource_collection_health = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationResourceCollectionType", Flag: "organization-resource-collection-type", Type: "types.OrganizationResourceCollectionType", Required: true},
	{Name: "OrganizationalUnitIds", Flag: "organizational-unit-ids", Type: "[]string", Required: false},
}

var fields_describe_resource_collection_health = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceCollectionType", Flag: "resource-collection-type", Type: "types.ResourceCollectionType", Required: true},
}

var fields_describe_service_integration = []leanruntime.Field{}

var fields_get_cost_estimation = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_resource_collection = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceCollectionType", Flag: "resource-collection-type", Type: "types.ResourceCollectionType", Required: true},
}

var fields_list_anomalies_for_insight = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListAnomaliesForInsightFilters", Required: false},
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTimeRange", Flag: "start-time-range", Type: "*types.StartTimeRange", Required: false},
}

var fields_list_anomalous_log_groups = []leanruntime.Field{
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_events = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListEventsFilters", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_insights = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "*types.ListInsightsStatusFilter", Required: true},
}

var fields_list_monitored_resources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListMonitoredResourcesFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_notification_channels = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_organization_insights = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationalUnitIds", Flag: "organizational-unit-ids", Type: "[]string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "*types.ListInsightsStatusFilter", Required: true},
}

var fields_list_recommendations = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_feedback = []leanruntime.Field{
	{Name: "InsightFeedback", Flag: "insight-feedback", Type: "*types.InsightFeedback", Required: false},
}

var fields_remove_notification_channel = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_search_insights = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.SearchInsightsFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTimeRange", Flag: "start-time-range", Type: "*types.StartTimeRange", Required: true},
	{Name: "Type", Flag: "type", Type: "types.InsightType", Required: true},
}

var fields_search_organization_insights = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.SearchOrganizationInsightsFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTimeRange", Flag: "start-time-range", Type: "*types.StartTimeRange", Required: true},
	{Name: "Type", Flag: "type", Type: "types.InsightType", Required: true},
}

var fields_start_cost_estimation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceCollection", Flag: "resource-collection", Type: "*types.CostEstimationResourceCollectionFilter", Required: true},
}

var fields_update_event_sources_config = []leanruntime.Field{
	{Name: "EventSources", Flag: "event-sources", Type: "*types.EventSourcesConfig", Required: false},
}

var fields_update_resource_collection = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.UpdateResourceCollectionAction", Required: true},
	{Name: "ResourceCollection", Flag: "resource-collection", Type: "*types.UpdateResourceCollectionFilter", Required: true},
}

var fields_update_service_integration = []leanruntime.Field{
	{Name: "ServiceIntegration", Flag: "service-integration", Type: "*types.UpdateServiceIntegrationConfig", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-notification-channel": {
			Name:   "add-notification-channel",
			Fields: fields_add_notification_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddNotificationChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_notification_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddNotificationChannel(ctx, input)
			},
		},
		"delete-insight": {
			Name:   "delete-insight",
			Fields: fields_delete_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInsight(ctx, input)
			},
		},
		"describe-account-health": {
			Name:   "describe-account-health",
			Fields: fields_describe_account_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountHealth(ctx, input)
			},
		},
		"describe-account-overview": {
			Name:   "describe-account-overview",
			Fields: fields_describe_account_overview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountOverviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_overview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountOverview(ctx, input)
			},
		},
		"describe-anomaly": {
			Name:   "describe-anomaly",
			Fields: fields_describe_anomaly,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnomalyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_anomaly, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAnomaly(ctx, input)
			},
		},
		"describe-event-sources-config": {
			Name:   "describe-event-sources-config",
			Fields: fields_describe_event_sources_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventSourcesConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_sources_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventSourcesConfig(ctx, input)
			},
		},
		"describe-feedback": {
			Name:   "describe-feedback",
			Fields: fields_describe_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFeedback(ctx, input)
			},
		},
		"describe-insight": {
			Name:   "describe-insight",
			Fields: fields_describe_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInsight(ctx, input)
			},
		},
		"describe-organization-health": {
			Name:   "describe-organization-health",
			Fields: fields_describe_organization_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organization_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganizationHealth(ctx, input)
			},
		},
		"describe-organization-overview": {
			Name:   "describe-organization-overview",
			Fields: fields_describe_organization_overview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationOverviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organization_overview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganizationOverview(ctx, input)
			},
		},
		"describe-organization-resource-collection-health": {
			Name:   "describe-organization-resource-collection-health",
			Fields: fields_describe_organization_resource_collection_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationResourceCollectionHealthInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_organization_resource_collection_health, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrganizationResourceCollectionHealth(ctx, input)
				}
				var results []*svc.DescribeOrganizationResourceCollectionHealthOutput
				p := svc.NewDescribeOrganizationResourceCollectionHealthPaginator(client, input)
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
		"describe-resource-collection-health": {
			Name:   "describe-resource-collection-health",
			Fields: fields_describe_resource_collection_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceCollectionHealthInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_resource_collection_health, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeResourceCollectionHealth(ctx, input)
				}
				var results []*svc.DescribeResourceCollectionHealthOutput
				p := svc.NewDescribeResourceCollectionHealthPaginator(client, input)
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
		"describe-service-integration": {
			Name:   "describe-service-integration",
			Fields: fields_describe_service_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceIntegration(ctx, input)
			},
		},
		"get-cost-estimation": {
			Name:   "get-cost-estimation",
			Fields: fields_get_cost_estimation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostEstimationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_cost_estimation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCostEstimation(ctx, input)
				}
				var results []*svc.GetCostEstimationOutput
				p := svc.NewGetCostEstimationPaginator(client, input)
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
		"get-resource-collection": {
			Name:   "get-resource-collection",
			Fields: fields_get_resource_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceCollectionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_collection, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceCollection(ctx, input)
				}
				var results []*svc.GetResourceCollectionOutput
				p := svc.NewGetResourceCollectionPaginator(client, input)
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
		"list-anomalies-for-insight": {
			Name:   "list-anomalies-for-insight",
			Fields: fields_list_anomalies_for_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnomaliesForInsightInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_anomalies_for_insight, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnomaliesForInsight(ctx, input)
				}
				var results []*svc.ListAnomaliesForInsightOutput
				p := svc.NewListAnomaliesForInsightPaginator(client, input)
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
		"list-anomalous-log-groups": {
			Name:   "list-anomalous-log-groups",
			Fields: fields_list_anomalous_log_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnomalousLogGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_anomalous_log_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnomalousLogGroups(ctx, input)
				}
				var results []*svc.ListAnomalousLogGroupsOutput
				p := svc.NewListAnomalousLogGroupsPaginator(client, input)
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
		"list-events": {
			Name:   "list-events",
			Fields: fields_list_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEvents(ctx, input)
				}
				var results []*svc.ListEventsOutput
				p := svc.NewListEventsPaginator(client, input)
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
		"list-insights": {
			Name:   "list-insights",
			Fields: fields_list_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInsights(ctx, input)
				}
				var results []*svc.ListInsightsOutput
				p := svc.NewListInsightsPaginator(client, input)
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
		"list-monitored-resources": {
			Name:   "list-monitored-resources",
			Fields: fields_list_monitored_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitoredResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitored_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitoredResources(ctx, input)
				}
				var results []*svc.ListMonitoredResourcesOutput
				p := svc.NewListMonitoredResourcesPaginator(client, input)
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
		"list-notification-channels": {
			Name:   "list-notification-channels",
			Fields: fields_list_notification_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notification_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotificationChannels(ctx, input)
				}
				var results []*svc.ListNotificationChannelsOutput
				p := svc.NewListNotificationChannelsPaginator(client, input)
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
		"list-organization-insights": {
			Name:   "list-organization-insights",
			Fields: fields_list_organization_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationInsights(ctx, input)
				}
				var results []*svc.ListOrganizationInsightsOutput
				p := svc.NewListOrganizationInsightsPaginator(client, input)
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
		"list-recommendations": {
			Name:   "list-recommendations",
			Fields: fields_list_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendations(ctx, input)
				}
				var results []*svc.ListRecommendationsOutput
				p := svc.NewListRecommendationsPaginator(client, input)
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
		"put-feedback": {
			Name:   "put-feedback",
			Fields: fields_put_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFeedback(ctx, input)
			},
		},
		"remove-notification-channel": {
			Name:   "remove-notification-channel",
			Fields: fields_remove_notification_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveNotificationChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_notification_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveNotificationChannel(ctx, input)
			},
		},
		"search-insights": {
			Name:   "search-insights",
			Fields: fields_search_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchInsights(ctx, input)
				}
				var results []*svc.SearchInsightsOutput
				p := svc.NewSearchInsightsPaginator(client, input)
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
		"search-organization-insights": {
			Name:   "search-organization-insights",
			Fields: fields_search_organization_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchOrganizationInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_organization_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchOrganizationInsights(ctx, input)
				}
				var results []*svc.SearchOrganizationInsightsOutput
				p := svc.NewSearchOrganizationInsightsPaginator(client, input)
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
		"start-cost-estimation": {
			Name:   "start-cost-estimation",
			Fields: fields_start_cost_estimation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCostEstimationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_cost_estimation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCostEstimation(ctx, input)
			},
		},
		"update-event-sources-config": {
			Name:   "update-event-sources-config",
			Fields: fields_update_event_sources_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventSourcesConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_sources_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventSourcesConfig(ctx, input)
			},
		},
		"update-resource-collection": {
			Name:   "update-resource-collection",
			Fields: fields_update_resource_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceCollection(ctx, input)
			},
		},
		"update-service-integration": {
			Name:   "update-service-integration",
			Fields: fields_update_service_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceIntegration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("devopsguru", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
