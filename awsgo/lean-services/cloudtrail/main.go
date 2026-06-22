package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

var fields_add_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagsList", Flag: "tags-list", Type: "[]types.Tag", Required: true},
}

var fields_cancel_query = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "EventDataStoreOwnerAccountId", Flag: "event-data-store-owner-account-id", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "Destinations", Flag: "destinations", Type: "[]types.Destination", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dashboard = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RefreshSchedule", Flag: "refresh-schedule", Type: "*types.RefreshSchedule", Required: false},
	{Name: "TagsList", Flag: "tags-list", Type: "[]types.Tag", Required: false},
	{Name: "TerminationProtectionEnabled", Flag: "termination-protection-enabled", Type: "*bool", Required: false},
	{Name: "Widgets", Flag: "widgets", Type: "[]types.RequestWidget", Required: false},
}

var fields_create_event_data_store = []leanruntime.Field{
	{Name: "AdvancedEventSelectors", Flag: "advanced-event-selectors", Type: "[]types.AdvancedEventSelector", Required: false},
	{Name: "BillingMode", Flag: "billing-mode", Type: "types.BillingMode", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MultiRegionEnabled", Flag: "multi-region-enabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationEnabled", Flag: "organization-enabled", Type: "*bool", Required: false},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: false},
	{Name: "StartIngestion", Flag: "start-ingestion", Type: "*bool", Required: false},
	{Name: "TagsList", Flag: "tags-list", Type: "[]types.Tag", Required: false},
	{Name: "TerminationProtectionEnabled", Flag: "termination-protection-enabled", Type: "*bool", Required: false},
}

var fields_create_trail = []leanruntime.Field{
	{Name: "CloudWatchLogsLogGroupArn", Flag: "cloud-watch-logs-log-group-arn", Type: "*string", Required: false},
	{Name: "CloudWatchLogsRoleArn", Flag: "cloud-watch-logs-role-arn", Type: "*string", Required: false},
	{Name: "EnableLogFileValidation", Flag: "enable-log-file-validation", Type: "*bool", Required: false},
	{Name: "IncludeGlobalServiceEvents", Flag: "include-global-service-events", Type: "*bool", Required: false},
	{Name: "IsMultiRegionTrail", Flag: "is-multi-region-trail", Type: "*bool", Required: false},
	{Name: "IsOrganizationTrail", Flag: "is-organization-trail", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "S3KeyPrefix", Flag: "s3-key-prefix", Type: "*string", Required: false},
	{Name: "SnsTopicName", Flag: "sns-topic-name", Type: "*string", Required: false},
	{Name: "TagsList", Flag: "tags-list", Type: "[]types.Tag", Required: false},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "Channel", Flag: "channel", Type: "*string", Required: true},
}

var fields_delete_dashboard = []leanruntime.Field{
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
}

var fields_delete_event_data_store = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_trail = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_deregister_organization_delegated_admin = []leanruntime.Field{
	{Name: "DelegatedAdminAccountId", Flag: "delegated-admin-account-id", Type: "*string", Required: true},
}

var fields_describe_query = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "EventDataStoreOwnerAccountId", Flag: "event-data-store-owner-account-id", Type: "*string", Required: false},
	{Name: "QueryAlias", Flag: "query-alias", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: false},
	{Name: "RefreshId", Flag: "refresh-id", Type: "*string", Required: false},
}

var fields_describe_trails = []leanruntime.Field{
	{Name: "IncludeShadowTrails", Flag: "include-shadow-trails", Type: "*bool", Required: false},
	{Name: "TrailNameList", Flag: "trail-name-list", Type: "[]string", Required: false},
}

var fields_disable_federation = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
}

var fields_enable_federation = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
	{Name: "FederationRoleArn", Flag: "federation-role-arn", Type: "*string", Required: true},
}

var fields_generate_query = []leanruntime.Field{
	{Name: "EventDataStores", Flag: "event-data-stores", Type: "[]string", Required: true},
	{Name: "Prompt", Flag: "prompt", Type: "*string", Required: true},
}

var fields_get_channel = []leanruntime.Field{
	{Name: "Channel", Flag: "channel", Type: "*string", Required: true},
}

var fields_get_dashboard = []leanruntime.Field{
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
}

var fields_get_event_configuration = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: false},
}

var fields_get_event_data_store = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
}

var fields_get_event_selectors = []leanruntime.Field{
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: true},
}

var fields_get_import = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
}

var fields_get_insight_selectors = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: false},
}

var fields_get_query_results = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "EventDataStoreOwnerAccountId", Flag: "event-data-store-owner-account-id", Type: "*string", Required: false},
	{Name: "MaxQueryResults", Flag: "max-query-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_trail = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_trail_status = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dashboards = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DashboardType", Required: false},
}

var fields_list_event_data_stores = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_import_failures = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_imports = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "ImportStatus", Flag: "import-status", Type: "types.ImportStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_insights_data = []leanruntime.Field{
	{Name: "DataType", Flag: "data-type", Type: "types.ListInsightsDataType", Required: true},
	{Name: "Dimensions", Flag: "dimensions", Type: "map[string]string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "InsightSource", Flag: "insight-source", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_insights_metric_data = []leanruntime.Field{
	{Name: "DataType", Flag: "data-type", Type: "types.InsightsMetricDataType", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "ErrorCode", Flag: "error-code", Type: "*string", Required: false},
	{Name: "EventName", Flag: "event-name", Type: "*string", Required: true},
	{Name: "EventSource", Flag: "event-source", Type: "*string", Required: true},
	{Name: "InsightType", Flag: "insight-type", Type: "types.InsightType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Period", Flag: "period", Type: "*int32", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: false},
}

var fields_list_public_keys = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_queries = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryStatus", Flag: "query-status", Type: "types.QueryStatus", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdList", Flag: "resource-id-list", Type: "[]string", Required: true},
}

var fields_list_trails = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_lookup_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EventCategory", Flag: "event-category", Type: "types.EventCategory", Required: false},
	{Name: "LookupAttributes", Flag: "lookup-attributes", Type: "[]types.LookupAttribute", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_put_event_configuration = []leanruntime.Field{
	{Name: "AggregationConfigurations", Flag: "aggregation-configurations", Type: "[]types.AggregationConfiguration", Required: false},
	{Name: "ContextKeySelectors", Flag: "context-key-selectors", Type: "[]types.ContextKeySelector", Required: false},
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "MaxEventSize", Flag: "max-event-size", Type: "types.MaxEventSize", Required: false},
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: false},
}

var fields_put_event_selectors = []leanruntime.Field{
	{Name: "AdvancedEventSelectors", Flag: "advanced-event-selectors", Type: "[]types.AdvancedEventSelector", Required: false},
	{Name: "EventSelectors", Flag: "event-selectors", Type: "[]types.EventSelector", Required: false},
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: true},
}

var fields_put_insight_selectors = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: false},
	{Name: "InsightSelectors", Flag: "insight-selectors", Type: "[]types.InsightSelector", Required: true},
	{Name: "InsightsDestination", Flag: "insights-destination", Type: "*string", Required: false},
	{Name: "TrailName", Flag: "trail-name", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
}

var fields_register_organization_delegated_admin = []leanruntime.Field{
	{Name: "MemberAccountId", Flag: "member-account-id", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagsList", Flag: "tags-list", Type: "[]types.Tag", Required: true},
}

var fields_restore_event_data_store = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
}

var fields_search_sample_queries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchPhrase", Flag: "search-phrase", Type: "*string", Required: true},
}

var fields_start_dashboard_refresh = []leanruntime.Field{
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "QueryParameterValues", Flag: "query-parameter-values", Type: "map[string]string", Required: false},
}

var fields_start_event_data_store_ingestion = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
}

var fields_start_import = []leanruntime.Field{
	{Name: "Destinations", Flag: "destinations", Type: "[]string", Required: false},
	{Name: "EndEventTime", Flag: "end-event-time", Type: "*time.Time", Required: false},
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: false},
	{Name: "ImportSource", Flag: "import-source", Type: "*types.ImportSource", Required: false},
	{Name: "StartEventTime", Flag: "start-event-time", Type: "*time.Time", Required: false},
}

var fields_start_logging = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_query = []leanruntime.Field{
	{Name: "DeliveryS3Uri", Flag: "delivery-s3-uri", Type: "*string", Required: false},
	{Name: "EventDataStoreOwnerAccountId", Flag: "event-data-store-owner-account-id", Type: "*string", Required: false},
	{Name: "QueryAlias", Flag: "query-alias", Type: "*string", Required: false},
	{Name: "QueryParameters", Flag: "query-parameters", Type: "[]string", Required: false},
	{Name: "QueryStatement", Flag: "query-statement", Type: "*string", Required: false},
}

var fields_stop_event_data_store_ingestion = []leanruntime.Field{
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
}

var fields_stop_import = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
}

var fields_stop_logging = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_channel = []leanruntime.Field{
	{Name: "Channel", Flag: "channel", Type: "*string", Required: true},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.Destination", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_dashboard = []leanruntime.Field{
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "RefreshSchedule", Flag: "refresh-schedule", Type: "*types.RefreshSchedule", Required: false},
	{Name: "TerminationProtectionEnabled", Flag: "termination-protection-enabled", Type: "*bool", Required: false},
	{Name: "Widgets", Flag: "widgets", Type: "[]types.RequestWidget", Required: false},
}

var fields_update_event_data_store = []leanruntime.Field{
	{Name: "AdvancedEventSelectors", Flag: "advanced-event-selectors", Type: "[]types.AdvancedEventSelector", Required: false},
	{Name: "BillingMode", Flag: "billing-mode", Type: "types.BillingMode", Required: false},
	{Name: "EventDataStore", Flag: "event-data-store", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MultiRegionEnabled", Flag: "multi-region-enabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OrganizationEnabled", Flag: "organization-enabled", Type: "*bool", Required: false},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: false},
	{Name: "TerminationProtectionEnabled", Flag: "termination-protection-enabled", Type: "*bool", Required: false},
}

var fields_update_trail = []leanruntime.Field{
	{Name: "CloudWatchLogsLogGroupArn", Flag: "cloud-watch-logs-log-group-arn", Type: "*string", Required: false},
	{Name: "CloudWatchLogsRoleArn", Flag: "cloud-watch-logs-role-arn", Type: "*string", Required: false},
	{Name: "EnableLogFileValidation", Flag: "enable-log-file-validation", Type: "*bool", Required: false},
	{Name: "IncludeGlobalServiceEvents", Flag: "include-global-service-events", Type: "*bool", Required: false},
	{Name: "IsMultiRegionTrail", Flag: "is-multi-region-trail", Type: "*bool", Required: false},
	{Name: "IsOrganizationTrail", Flag: "is-organization-trail", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: false},
	{Name: "S3KeyPrefix", Flag: "s3-key-prefix", Type: "*string", Required: false},
	{Name: "SnsTopicName", Flag: "sns-topic-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"cancel-query": {
			Name:   "cancel-query",
			Fields: fields_cancel_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelQuery(ctx, input)
			},
		},
		"create-channel": {
			Name:   "create-channel",
			Fields: fields_create_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannel(ctx, input)
			},
		},
		"create-dashboard": {
			Name:   "create-dashboard",
			Fields: fields_create_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDashboard(ctx, input)
			},
		},
		"create-event-data-store": {
			Name:   "create-event-data-store",
			Fields: fields_create_event_data_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventDataStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_data_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventDataStore(ctx, input)
			},
		},
		"create-trail": {
			Name:   "create-trail",
			Fields: fields_create_trail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrail(ctx, input)
			},
		},
		"delete-channel": {
			Name:   "delete-channel",
			Fields: fields_delete_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannel(ctx, input)
			},
		},
		"delete-dashboard": {
			Name:   "delete-dashboard",
			Fields: fields_delete_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDashboard(ctx, input)
			},
		},
		"delete-event-data-store": {
			Name:   "delete-event-data-store",
			Fields: fields_delete_event_data_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventDataStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_data_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventDataStore(ctx, input)
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
		"delete-trail": {
			Name:   "delete-trail",
			Fields: fields_delete_trail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrail(ctx, input)
			},
		},
		"deregister-organization-delegated-admin": {
			Name:   "deregister-organization-delegated-admin",
			Fields: fields_deregister_organization_delegated_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterOrganizationDelegatedAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_organization_delegated_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterOrganizationDelegatedAdmin(ctx, input)
			},
		},
		"describe-query": {
			Name:   "describe-query",
			Fields: fields_describe_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQuery(ctx, input)
			},
		},
		"describe-trails": {
			Name:   "describe-trails",
			Fields: fields_describe_trails,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trails, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrails(ctx, input)
			},
		},
		"disable-federation": {
			Name:   "disable-federation",
			Fields: fields_disable_federation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableFederationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_federation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableFederation(ctx, input)
			},
		},
		"enable-federation": {
			Name:   "enable-federation",
			Fields: fields_enable_federation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableFederationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_federation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableFederation(ctx, input)
			},
		},
		"generate-query": {
			Name:   "generate-query",
			Fields: fields_generate_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateQuery(ctx, input)
			},
		},
		"get-channel": {
			Name:   "get-channel",
			Fields: fields_get_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannel(ctx, input)
			},
		},
		"get-dashboard": {
			Name:   "get-dashboard",
			Fields: fields_get_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDashboard(ctx, input)
			},
		},
		"get-event-configuration": {
			Name:   "get-event-configuration",
			Fields: fields_get_event_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventConfiguration(ctx, input)
			},
		},
		"get-event-data-store": {
			Name:   "get-event-data-store",
			Fields: fields_get_event_data_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventDataStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_data_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventDataStore(ctx, input)
			},
		},
		"get-event-selectors": {
			Name:   "get-event-selectors",
			Fields: fields_get_event_selectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventSelectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_selectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventSelectors(ctx, input)
			},
		},
		"get-import": {
			Name:   "get-import",
			Fields: fields_get_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImport(ctx, input)
			},
		},
		"get-insight-selectors": {
			Name:   "get-insight-selectors",
			Fields: fields_get_insight_selectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightSelectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insight_selectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsightSelectors(ctx, input)
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
		"get-trail": {
			Name:   "get-trail",
			Fields: fields_get_trail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrail(ctx, input)
			},
		},
		"get-trail-status": {
			Name:   "get-trail-status",
			Fields: fields_get_trail_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrailStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trail_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrailStatus(ctx, input)
			},
		},
		"list-channels": {
			Name:   "list-channels",
			Fields: fields_list_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannels(ctx, input)
				}
				var results []*svc.ListChannelsOutput
				p := svc.NewListChannelsPaginator(client, input)
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
		"list-dashboards": {
			Name:   "list-dashboards",
			Fields: fields_list_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDashboardsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_dashboards, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDashboards(ctx, input)
			},
		},
		"list-event-data-stores": {
			Name:   "list-event-data-stores",
			Fields: fields_list_event_data_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventDataStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_data_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventDataStores(ctx, input)
				}
				var results []*svc.ListEventDataStoresOutput
				p := svc.NewListEventDataStoresPaginator(client, input)
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
		"list-import-failures": {
			Name:   "list-import-failures",
			Fields: fields_list_import_failures,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportFailuresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_import_failures, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportFailures(ctx, input)
				}
				var results []*svc.ListImportFailuresOutput
				p := svc.NewListImportFailuresPaginator(client, input)
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
		"list-imports": {
			Name:   "list-imports",
			Fields: fields_list_imports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_imports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImports(ctx, input)
				}
				var results []*svc.ListImportsOutput
				p := svc.NewListImportsPaginator(client, input)
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
		"list-insights-data": {
			Name:   "list-insights-data",
			Fields: fields_list_insights_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInsightsDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_insights_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInsightsData(ctx, input)
				}
				var results []*svc.ListInsightsDataOutput
				p := svc.NewListInsightsDataPaginator(client, input)
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
		"list-insights-metric-data": {
			Name:   "list-insights-metric-data",
			Fields: fields_list_insights_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInsightsMetricDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_insights_metric_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInsightsMetricData(ctx, input)
				}
				var results []*svc.ListInsightsMetricDataOutput
				p := svc.NewListInsightsMetricDataPaginator(client, input)
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
		"list-public-keys": {
			Name:   "list-public-keys",
			Fields: fields_list_public_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPublicKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_public_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPublicKeys(ctx, input)
				}
				var results []*svc.ListPublicKeysOutput
				p := svc.NewListPublicKeysPaginator(client, input)
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
		"list-queries": {
			Name:   "list-queries",
			Fields: fields_list_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueries(ctx, input)
				}
				var results []*svc.ListQueriesOutput
				p := svc.NewListQueriesPaginator(client, input)
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
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTags(ctx, input)
				}
				var results []*svc.ListTagsOutput
				p := svc.NewListTagsPaginator(client, input)
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
		"list-trails": {
			Name:   "list-trails",
			Fields: fields_list_trails,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trails, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrails(ctx, input)
				}
				var results []*svc.ListTrailsOutput
				p := svc.NewListTrailsPaginator(client, input)
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
		"lookup-events": {
			Name:   "lookup-events",
			Fields: fields_lookup_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LookupEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_lookup_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.LookupEvents(ctx, input)
				}
				var results []*svc.LookupEventsOutput
				p := svc.NewLookupEventsPaginator(client, input)
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
		"put-event-configuration": {
			Name:   "put-event-configuration",
			Fields: fields_put_event_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_event_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEventConfiguration(ctx, input)
			},
		},
		"put-event-selectors": {
			Name:   "put-event-selectors",
			Fields: fields_put_event_selectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventSelectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_event_selectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEventSelectors(ctx, input)
			},
		},
		"put-insight-selectors": {
			Name:   "put-insight-selectors",
			Fields: fields_put_insight_selectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInsightSelectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_insight_selectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInsightSelectors(ctx, input)
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
		"register-organization-delegated-admin": {
			Name:   "register-organization-delegated-admin",
			Fields: fields_register_organization_delegated_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterOrganizationDelegatedAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_organization_delegated_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterOrganizationDelegatedAdmin(ctx, input)
			},
		},
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"restore-event-data-store": {
			Name:   "restore-event-data-store",
			Fields: fields_restore_event_data_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreEventDataStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_event_data_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreEventDataStore(ctx, input)
			},
		},
		"search-sample-queries": {
			Name:   "search-sample-queries",
			Fields: fields_search_sample_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSampleQueriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_sample_queries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchSampleQueries(ctx, input)
			},
		},
		"start-dashboard-refresh": {
			Name:   "start-dashboard-refresh",
			Fields: fields_start_dashboard_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDashboardRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dashboard_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDashboardRefresh(ctx, input)
			},
		},
		"start-event-data-store-ingestion": {
			Name:   "start-event-data-store-ingestion",
			Fields: fields_start_event_data_store_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEventDataStoreIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_event_data_store_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEventDataStoreIngestion(ctx, input)
			},
		},
		"start-import": {
			Name:   "start-import",
			Fields: fields_start_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImport(ctx, input)
			},
		},
		"start-logging": {
			Name:   "start-logging",
			Fields: fields_start_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartLogging(ctx, input)
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
		"stop-event-data-store-ingestion": {
			Name:   "stop-event-data-store-ingestion",
			Fields: fields_stop_event_data_store_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEventDataStoreIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_event_data_store_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEventDataStoreIngestion(ctx, input)
			},
		},
		"stop-import": {
			Name:   "stop-import",
			Fields: fields_stop_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopImport(ctx, input)
			},
		},
		"stop-logging": {
			Name:   "stop-logging",
			Fields: fields_stop_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopLogging(ctx, input)
			},
		},
		"update-channel": {
			Name:   "update-channel",
			Fields: fields_update_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannel(ctx, input)
			},
		},
		"update-dashboard": {
			Name:   "update-dashboard",
			Fields: fields_update_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboard(ctx, input)
			},
		},
		"update-event-data-store": {
			Name:   "update-event-data-store",
			Fields: fields_update_event_data_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventDataStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_data_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventDataStore(ctx, input)
			},
		},
		"update-trail": {
			Name:   "update-trail",
			Fields: fields_update_trail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrail(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudtrail", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
