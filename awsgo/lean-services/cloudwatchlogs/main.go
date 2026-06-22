package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

var fields_associate_kms_key = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
}

var fields_associate_source_to_s3_table_integration = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: true},
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
}

var fields_cancel_export_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_cancel_import_task = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
}

var fields_create_delivery = []leanruntime.Field{
	{Name: "DeliveryDestinationArn", Flag: "delivery-destination-arn", Type: "*string", Required: true},
	{Name: "DeliverySourceName", Flag: "delivery-source-name", Type: "*string", Required: true},
	{Name: "FieldDelimiter", Flag: "field-delimiter", Type: "*string", Required: false},
	{Name: "RecordFields", Flag: "record-fields", Type: "[]string", Required: false},
	{Name: "S3DeliveryConfiguration", Flag: "s3-delivery-configuration", Type: "*types.S3DeliveryConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_export_task = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: true},
	{Name: "DestinationPrefix", Flag: "destination-prefix", Type: "*string", Required: false},
	{Name: "From", Flag: "from", Type: "*int64", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "LogStreamNamePrefix", Flag: "log-stream-name-prefix", Type: "*string", Required: false},
	{Name: "TaskName", Flag: "task-name", Type: "*string", Required: false},
	{Name: "To", Flag: "to", Type: "*int64", Required: true},
}

var fields_create_import_task = []leanruntime.Field{
	{Name: "ImportFilter", Flag: "import-filter", Type: "*types.ImportFilter", Required: false},
	{Name: "ImportRoleArn", Flag: "import-role-arn", Type: "*string", Required: true},
	{Name: "ImportSourceArn", Flag: "import-source-arn", Type: "*string", Required: true},
}

var fields_create_log_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyVisibilityTime", Flag: "anomaly-visibility-time", Type: "*int64", Required: false},
	{Name: "DetectorName", Flag: "detector-name", Type: "*string", Required: false},
	{Name: "EvaluationFrequency", Flag: "evaluation-frequency", Type: "types.EvaluationFrequency", Required: false},
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LogGroupArnList", Flag: "log-group-arn-list", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_log_group = []leanruntime.Field{
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LogGroupClass", Flag: "log-group-class", Type: "types.LogGroupClass", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_log_stream = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "LogStreamName", Flag: "log-stream-name", Type: "*string", Required: true},
}

var fields_create_scheduled_query = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "*types.DestinationConfiguration", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueryLanguage", Flag: "query-language", Type: "types.QueryLanguage", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "ScheduleEndTime", Flag: "schedule-end-time", Type: "*int64", Required: false},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: true},
	{Name: "ScheduleStartTime", Flag: "schedule-start-time", Type: "*int64", Required: false},
	{Name: "StartTimeOffset", Flag: "start-time-offset", Type: "*int64", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScheduledQueryState", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
}

var fields_delete_account_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
}

var fields_delete_data_protection_policy = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_delete_delivery = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_delivery_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_delivery_destination_policy = []leanruntime.Field{
	{Name: "DeliveryDestinationName", Flag: "delivery-destination-name", Type: "*string", Required: true},
}

var fields_delete_delivery_source = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_destination = []leanruntime.Field{
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
}

var fields_delete_index_policy = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_delete_integration = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: true},
}

var fields_delete_log_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyDetectorArn", Flag: "anomaly-detector-arn", Type: "*string", Required: true},
}

var fields_delete_log_group = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
}

var fields_delete_log_stream = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "LogStreamName", Flag: "log-stream-name", Type: "*string", Required: true},
}

var fields_delete_metric_filter = []leanruntime.Field{
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
}

var fields_delete_query_definition = []leanruntime.Field{
	{Name: "QueryDefinitionId", Flag: "query-definition-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_delete_retention_policy = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
}

var fields_delete_scheduled_query = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_subscription_filter = []leanruntime.Field{
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
}

var fields_delete_transformer = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_describe_account_policies = []leanruntime.Field{
	{Name: "AccountIdentifiers", Flag: "account-identifiers", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
}

var fields_describe_configuration_templates = []leanruntime.Field{
	{Name: "DeliveryDestinationTypes", Flag: "delivery-destination-types", Type: "[]types.DeliveryDestinationType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogTypes", Flag: "log-types", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: false},
}

var fields_describe_deliveries = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_delivery_destinations = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_delivery_sources = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_destinations = []leanruntime.Field{
	{Name: "DestinationNamePrefix", Flag: "destination-name-prefix", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_export_tasks = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusCode", Flag: "status-code", Type: "types.ExportTaskStatusCode", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: false},
}

var fields_describe_field_indexes = []leanruntime.Field{
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_import_task_batches = []leanruntime.Field{
	{Name: "BatchImportStatus", Flag: "batch-import-status", Type: "[]types.ImportStatus", Required: false},
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_import_tasks = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: false},
	{Name: "ImportSourceArn", Flag: "import-source-arn", Type: "*string", Required: false},
	{Name: "ImportStatus", Flag: "import-status", Type: "types.ImportStatus", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_index_policies = []leanruntime.Field{
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_log_groups = []leanruntime.Field{
	{Name: "AccountIdentifiers", Flag: "account-identifiers", Type: "[]string", Required: false},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupClass", Flag: "log-group-class", Type: "types.LogGroupClass", Required: false},
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: false},
	{Name: "LogGroupNamePattern", Flag: "log-group-name-pattern", Type: "*string", Required: false},
	{Name: "LogGroupNamePrefix", Flag: "log-group-name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_log_streams = []leanruntime.Field{
	{Name: "Descending", Flag: "descending", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "LogStreamNamePrefix", Flag: "log-stream-name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "types.OrderBy", Required: false},
}

var fields_describe_metric_filters = []leanruntime.Field{
	{Name: "FilterNamePrefix", Flag: "filter-name-prefix", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "MetricNamespace", Flag: "metric-namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_queries = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryLanguage", Flag: "query-language", Type: "types.QueryLanguage", Required: false},
	{Name: "Status", Flag: "status", Type: "types.QueryStatus", Required: false},
}

var fields_describe_query_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryDefinitionNamePrefix", Flag: "query-definition-name-prefix", Type: "*string", Required: false},
	{Name: "QueryLanguage", Flag: "query-language", Type: "types.QueryLanguage", Required: false},
}

var fields_describe_resource_policies = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyScope", Flag: "policy-scope", Type: "types.PolicyScope", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_describe_subscription_filters = []leanruntime.Field{
	{Name: "FilterNamePrefix", Flag: "filter-name-prefix", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_kms_key = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
}

var fields_disassociate_source_from_s3_table_integration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_filter_log_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*int64", Required: false},
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: false},
	{Name: "Interleaved", Flag: "interleaved", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "LogStreamNamePrefix", Flag: "log-stream-name-prefix", Type: "*string", Required: false},
	{Name: "LogStreamNames", Flag: "log-stream-names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*int64", Required: false},
	{Name: "Unmask", Flag: "unmask", Type: "bool", Required: false},
}

var fields_get_data_protection_policy = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_get_delivery = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_delivery_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_delivery_destination_policy = []leanruntime.Field{
	{Name: "DeliveryDestinationName", Flag: "delivery-destination-name", Type: "*string", Required: true},
}

var fields_get_delivery_source = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_integration = []leanruntime.Field{
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: true},
}

var fields_get_log_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyDetectorArn", Flag: "anomaly-detector-arn", Type: "*string", Required: true},
}

var fields_get_log_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*int64", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "LogStreamName", Flag: "log-stream-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartFromHead", Flag: "start-from-head", Type: "*bool", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*int64", Required: false},
	{Name: "Unmask", Flag: "unmask", Type: "bool", Required: false},
}

var fields_get_log_fields = []leanruntime.Field{
	{Name: "DataSourceName", Flag: "data-source-name", Type: "*string", Required: true},
	{Name: "DataSourceType", Flag: "data-source-type", Type: "*string", Required: true},
}

var fields_get_log_group_fields = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "Time", Flag: "time", Type: "*int64", Required: false},
}

var fields_get_log_object = []leanruntime.Field{
	{Name: "LogObjectPointer", Flag: "log-object-pointer", Type: "*string", Required: true},
	{Name: "Unmask", Flag: "unmask", Type: "bool", Required: false},
}

var fields_get_log_record = []leanruntime.Field{
	{Name: "LogRecordPointer", Flag: "log-record-pointer", Type: "*string", Required: true},
	{Name: "Unmask", Flag: "unmask", Type: "bool", Required: false},
}

var fields_get_query_results = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_scheduled_query = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_scheduled_query_history = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*int64", Required: true},
	{Name: "ExecutionStatuses", Flag: "execution-statuses", Type: "[]types.ExecutionStatus", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*int64", Required: true},
}

var fields_get_transformer = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_list_aggregate_log_group_summaries = []leanruntime.Field{
	{Name: "AccountIdentifiers", Flag: "account-identifiers", Type: "[]string", Required: false},
	{Name: "DataSources", Flag: "data-sources", Type: "[]types.DataSourceFilter", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "types.ListAggregateLogGroupSummariesGroupBy", Required: true},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupClass", Flag: "log-group-class", Type: "types.LogGroupClass", Required: false},
	{Name: "LogGroupNamePattern", Flag: "log-group-name-pattern", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_anomalies = []leanruntime.Field{
	{Name: "AnomalyDetectorArn", Flag: "anomaly-detector-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SuppressionState", Flag: "suppression-state", Type: "types.SuppressionState", Required: false},
}

var fields_list_integrations = []leanruntime.Field{
	{Name: "IntegrationNamePrefix", Flag: "integration-name-prefix", Type: "*string", Required: false},
	{Name: "IntegrationStatus", Flag: "integration-status", Type: "types.IntegrationStatus", Required: false},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: false},
}

var fields_list_log_anomaly_detectors = []leanruntime.Field{
	{Name: "FilterLogGroupArn", Flag: "filter-log-group-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_log_groups = []leanruntime.Field{
	{Name: "AccountIdentifiers", Flag: "account-identifiers", Type: "[]string", Required: false},
	{Name: "DataSources", Flag: "data-sources", Type: "[]types.DataSourceFilter", Required: false},
	{Name: "FieldIndexNames", Flag: "field-index-names", Type: "[]string", Required: false},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupClass", Flag: "log-group-class", Type: "types.LogGroupClass", Required: false},
	{Name: "LogGroupNamePattern", Flag: "log-group-name-pattern", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_log_groups_for_query = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_list_scheduled_queries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScheduledQueryState", Required: false},
}

var fields_list_sources_for_s3_table_integration = []leanruntime.Field{
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tags_log_group = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
}

var fields_put_account_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: false},
	{Name: "SelectionCriteria", Flag: "selection-criteria", Type: "*string", Required: false},
}

var fields_put_bearer_token_authentication = []leanruntime.Field{
	{Name: "BearerTokenAuthenticationEnabled", Flag: "bearer-token-authentication-enabled", Type: "*bool", Required: true},
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_put_data_protection_policy = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
}

var fields_put_delivery_destination = []leanruntime.Field{
	{Name: "DeliveryDestinationConfiguration", Flag: "delivery-destination-configuration", Type: "*types.DeliveryDestinationConfiguration", Required: false},
	{Name: "DeliveryDestinationType", Flag: "delivery-destination-type", Type: "types.DeliveryDestinationType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_put_delivery_destination_policy = []leanruntime.Field{
	{Name: "DeliveryDestinationName", Flag: "delivery-destination-name", Type: "*string", Required: true},
	{Name: "DeliveryDestinationPolicy", Flag: "delivery-destination-policy", Type: "*string", Required: true},
}

var fields_put_delivery_source = []leanruntime.Field{
	{Name: "LogType", Flag: "log-type", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_put_destination = []leanruntime.Field{
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_put_destination_policy = []leanruntime.Field{
	{Name: "AccessPolicy", Flag: "access-policy", Type: "*string", Required: true},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "ForceUpdate", Flag: "force-update", Type: "*bool", Required: false},
}

var fields_put_index_policy = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
}

var fields_put_integration = []leanruntime.Field{
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: true},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: true},
	{Name: "ResourceConfig", Flag: "resource-config", Type: "types.ResourceConfig", Required: true},
}

var fields_put_log_events = []leanruntime.Field{
	{Name: "Entity", Flag: "entity", Type: "*types.Entity", Required: false},
	{Name: "LogEvents", Flag: "log-events", Type: "[]types.InputLogEvent", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "LogStreamName", Flag: "log-stream-name", Type: "*string", Required: true},
	{Name: "SequenceToken", Flag: "sequence-token", Type: "*string", Required: false},
}

var fields_put_log_group_deletion_protection = []leanruntime.Field{
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: true},
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
}

var fields_put_metric_filter = []leanruntime.Field{
	{Name: "ApplyOnTransformedLogs", Flag: "apply-on-transformed-logs", Type: "bool", Required: false},
	{Name: "EmitSystemFieldDimensions", Flag: "emit-system-field-dimensions", Type: "[]string", Required: false},
	{Name: "FieldSelectionCriteria", Flag: "field-selection-criteria", Type: "*string", Required: false},
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "MetricTransformations", Flag: "metric-transformations", Type: "[]types.MetricTransformation", Required: true},
}

var fields_put_query_definition = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LogGroupNames", Flag: "log-group-names", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueryDefinitionId", Flag: "query-definition-id", Type: "*string", Required: false},
	{Name: "QueryLanguage", Flag: "query-language", Type: "types.QueryLanguage", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_put_retention_policy = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "RetentionInDays", Flag: "retention-in-days", Type: "*int32", Required: true},
}

var fields_put_subscription_filter = []leanruntime.Field{
	{Name: "ApplyOnTransformedLogs", Flag: "apply-on-transformed-logs", Type: "bool", Required: false},
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: true},
	{Name: "Distribution", Flag: "distribution", Type: "types.Distribution", Required: false},
	{Name: "EmitSystemFields", Flag: "emit-system-fields", Type: "[]string", Required: false},
	{Name: "FieldSelectionCriteria", Flag: "field-selection-criteria", Type: "*string", Required: false},
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_put_transformer = []leanruntime.Field{
	{Name: "LogGroupIdentifier", Flag: "log-group-identifier", Type: "*string", Required: true},
	{Name: "TransformerConfig", Flag: "transformer-config", Type: "[]types.Processor", Required: true},
}

var fields_start_live_tail = []leanruntime.Field{
	{Name: "LogEventFilterPattern", Flag: "log-event-filter-pattern", Type: "*string", Required: false},
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: true},
	{Name: "LogStreamNamePrefixes", Flag: "log-stream-name-prefixes", Type: "[]string", Required: false},
	{Name: "LogStreamNames", Flag: "log-stream-names", Type: "[]string", Required: false},
}

var fields_start_query = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*int64", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "LogGroupNames", Flag: "log-group-names", Type: "[]string", Required: false},
	{Name: "QueryLanguage", Flag: "query-language", Type: "types.QueryLanguage", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*int64", Required: true},
}

var fields_stop_query = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_tag_log_group = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_test_metric_filter = []leanruntime.Field{
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: true},
	{Name: "LogEventMessages", Flag: "log-event-messages", Type: "[]string", Required: true},
}

var fields_test_transformer = []leanruntime.Field{
	{Name: "LogEventMessages", Flag: "log-event-messages", Type: "[]string", Required: true},
	{Name: "TransformerConfig", Flag: "transformer-config", Type: "[]types.Processor", Required: true},
}

var fields_untag_log_group = []leanruntime.Field{
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_anomaly = []leanruntime.Field{
	{Name: "AnomalyDetectorArn", Flag: "anomaly-detector-arn", Type: "*string", Required: true},
	{Name: "AnomalyId", Flag: "anomaly-id", Type: "*string", Required: false},
	{Name: "Baseline", Flag: "baseline", Type: "*bool", Required: false},
	{Name: "PatternId", Flag: "pattern-id", Type: "*string", Required: false},
	{Name: "SuppressionPeriod", Flag: "suppression-period", Type: "*types.SuppressionPeriod", Required: false},
	{Name: "SuppressionType", Flag: "suppression-type", Type: "types.SuppressionType", Required: false},
}

var fields_update_delivery_configuration = []leanruntime.Field{
	{Name: "FieldDelimiter", Flag: "field-delimiter", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RecordFields", Flag: "record-fields", Type: "[]string", Required: false},
	{Name: "S3DeliveryConfiguration", Flag: "s3-delivery-configuration", Type: "*types.S3DeliveryConfiguration", Required: false},
}

var fields_update_log_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyDetectorArn", Flag: "anomaly-detector-arn", Type: "*string", Required: true},
	{Name: "AnomalyVisibilityTime", Flag: "anomaly-visibility-time", Type: "*int64", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: true},
	{Name: "EvaluationFrequency", Flag: "evaluation-frequency", Type: "types.EvaluationFrequency", Required: false},
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: false},
}

var fields_update_scheduled_query = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "*types.DestinationConfiguration", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LogGroupIdentifiers", Flag: "log-group-identifiers", Type: "[]string", Required: false},
	{Name: "QueryLanguage", Flag: "query-language", Type: "types.QueryLanguage", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "ScheduleEndTime", Flag: "schedule-end-time", Type: "*int64", Required: false},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: true},
	{Name: "ScheduleStartTime", Flag: "schedule-start-time", Type: "*int64", Required: false},
	{Name: "StartTimeOffset", Flag: "start-time-offset", Type: "*int64", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScheduledQueryState", Required: false},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-kms-key": {
			Name:   "associate-kms-key",
			Fields: fields_associate_kms_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateKmsKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_kms_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateKmsKey(ctx, input)
			},
		},
		"associate-source-to-s3-table-integration": {
			Name:   "associate-source-to-s3-table-integration",
			Fields: fields_associate_source_to_s3_table_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSourceToS3TableIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_source_to_s3_table_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSourceToS3TableIntegration(ctx, input)
			},
		},
		"cancel-export-task": {
			Name:   "cancel-export-task",
			Fields: fields_cancel_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelExportTask(ctx, input)
			},
		},
		"cancel-import-task": {
			Name:   "cancel-import-task",
			Fields: fields_cancel_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelImportTask(ctx, input)
			},
		},
		"create-delivery": {
			Name:   "create-delivery",
			Fields: fields_create_delivery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeliveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_delivery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDelivery(ctx, input)
			},
		},
		"create-export-task": {
			Name:   "create-export-task",
			Fields: fields_create_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExportTask(ctx, input)
			},
		},
		"create-import-task": {
			Name:   "create-import-task",
			Fields: fields_create_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImportTask(ctx, input)
			},
		},
		"create-log-anomaly-detector": {
			Name:   "create-log-anomaly-detector",
			Fields: fields_create_log_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLogAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_log_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLogAnomalyDetector(ctx, input)
			},
		},
		"create-log-group": {
			Name:   "create-log-group",
			Fields: fields_create_log_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLogGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_log_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLogGroup(ctx, input)
			},
		},
		"create-log-stream": {
			Name:   "create-log-stream",
			Fields: fields_create_log_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLogStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_log_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLogStream(ctx, input)
			},
		},
		"create-scheduled-query": {
			Name:   "create-scheduled-query",
			Fields: fields_create_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScheduledQuery(ctx, input)
			},
		},
		"delete-account-policy": {
			Name:   "delete-account-policy",
			Fields: fields_delete_account_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountPolicy(ctx, input)
			},
		},
		"delete-data-protection-policy": {
			Name:   "delete-data-protection-policy",
			Fields: fields_delete_data_protection_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataProtectionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_protection_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataProtectionPolicy(ctx, input)
			},
		},
		"delete-delivery": {
			Name:   "delete-delivery",
			Fields: fields_delete_delivery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeliveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_delivery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDelivery(ctx, input)
			},
		},
		"delete-delivery-destination": {
			Name:   "delete-delivery-destination",
			Fields: fields_delete_delivery_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeliveryDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_delivery_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeliveryDestination(ctx, input)
			},
		},
		"delete-delivery-destination-policy": {
			Name:   "delete-delivery-destination-policy",
			Fields: fields_delete_delivery_destination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeliveryDestinationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_delivery_destination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeliveryDestinationPolicy(ctx, input)
			},
		},
		"delete-delivery-source": {
			Name:   "delete-delivery-source",
			Fields: fields_delete_delivery_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeliverySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_delivery_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeliverySource(ctx, input)
			},
		},
		"delete-destination": {
			Name:   "delete-destination",
			Fields: fields_delete_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDestination(ctx, input)
			},
		},
		"delete-index-policy": {
			Name:   "delete-index-policy",
			Fields: fields_delete_index_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndexPolicy(ctx, input)
			},
		},
		"delete-integration": {
			Name:   "delete-integration",
			Fields: fields_delete_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegration(ctx, input)
			},
		},
		"delete-log-anomaly-detector": {
			Name:   "delete-log-anomaly-detector",
			Fields: fields_delete_log_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLogAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_log_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLogAnomalyDetector(ctx, input)
			},
		},
		"delete-log-group": {
			Name:   "delete-log-group",
			Fields: fields_delete_log_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLogGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_log_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLogGroup(ctx, input)
			},
		},
		"delete-log-stream": {
			Name:   "delete-log-stream",
			Fields: fields_delete_log_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLogStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_log_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLogStream(ctx, input)
			},
		},
		"delete-metric-filter": {
			Name:   "delete-metric-filter",
			Fields: fields_delete_metric_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMetricFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_metric_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMetricFilter(ctx, input)
			},
		},
		"delete-query-definition": {
			Name:   "delete-query-definition",
			Fields: fields_delete_query_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueryDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_query_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueryDefinition(ctx, input)
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
		"delete-retention-policy": {
			Name:   "delete-retention-policy",
			Fields: fields_delete_retention_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRetentionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_retention_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRetentionPolicy(ctx, input)
			},
		},
		"delete-scheduled-query": {
			Name:   "delete-scheduled-query",
			Fields: fields_delete_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduledQuery(ctx, input)
			},
		},
		"delete-subscription-filter": {
			Name:   "delete-subscription-filter",
			Fields: fields_delete_subscription_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriptionFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscription_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriptionFilter(ctx, input)
			},
		},
		"delete-transformer": {
			Name:   "delete-transformer",
			Fields: fields_delete_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransformer(ctx, input)
			},
		},
		"describe-account-policies": {
			Name:   "describe-account-policies",
			Fields: fields_describe_account_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountPolicies(ctx, input)
			},
		},
		"describe-configuration-templates": {
			Name:   "describe-configuration-templates",
			Fields: fields_describe_configuration_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_configuration_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConfigurationTemplates(ctx, input)
				}
				var results []*svc.DescribeConfigurationTemplatesOutput
				p := svc.NewDescribeConfigurationTemplatesPaginator(client, input)
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
		"describe-deliveries": {
			Name:   "describe-deliveries",
			Fields: fields_describe_deliveries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeliveriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_deliveries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDeliveries(ctx, input)
				}
				var results []*svc.DescribeDeliveriesOutput
				p := svc.NewDescribeDeliveriesPaginator(client, input)
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
		"describe-delivery-destinations": {
			Name:   "describe-delivery-destinations",
			Fields: fields_describe_delivery_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeliveryDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_delivery_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDeliveryDestinations(ctx, input)
				}
				var results []*svc.DescribeDeliveryDestinationsOutput
				p := svc.NewDescribeDeliveryDestinationsPaginator(client, input)
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
		"describe-delivery-sources": {
			Name:   "describe-delivery-sources",
			Fields: fields_describe_delivery_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeliverySourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_delivery_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDeliverySources(ctx, input)
				}
				var results []*svc.DescribeDeliverySourcesOutput
				p := svc.NewDescribeDeliverySourcesPaginator(client, input)
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
		"describe-destinations": {
			Name:   "describe-destinations",
			Fields: fields_describe_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDestinations(ctx, input)
				}
				var results []*svc.DescribeDestinationsOutput
				p := svc.NewDescribeDestinationsPaginator(client, input)
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
		"describe-export-tasks": {
			Name:   "describe-export-tasks",
			Fields: fields_describe_export_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_export_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExportTasks(ctx, input)
			},
		},
		"describe-field-indexes": {
			Name:   "describe-field-indexes",
			Fields: fields_describe_field_indexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFieldIndexesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_field_indexes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFieldIndexes(ctx, input)
			},
		},
		"describe-import-task-batches": {
			Name:   "describe-import-task-batches",
			Fields: fields_describe_import_task_batches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportTaskBatchesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_import_task_batches, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImportTaskBatches(ctx, input)
			},
		},
		"describe-import-tasks": {
			Name:   "describe-import-tasks",
			Fields: fields_describe_import_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_import_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImportTasks(ctx, input)
			},
		},
		"describe-index-policies": {
			Name:   "describe-index-policies",
			Fields: fields_describe_index_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIndexPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_index_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIndexPolicies(ctx, input)
			},
		},
		"describe-log-groups": {
			Name:   "describe-log-groups",
			Fields: fields_describe_log_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLogGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_log_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLogGroups(ctx, input)
				}
				var results []*svc.DescribeLogGroupsOutput
				p := svc.NewDescribeLogGroupsPaginator(client, input)
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
		"describe-log-streams": {
			Name:   "describe-log-streams",
			Fields: fields_describe_log_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLogStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_log_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLogStreams(ctx, input)
				}
				var results []*svc.DescribeLogStreamsOutput
				p := svc.NewDescribeLogStreamsPaginator(client, input)
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
		"describe-metric-filters": {
			Name:   "describe-metric-filters",
			Fields: fields_describe_metric_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetricFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metric_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetricFilters(ctx, input)
				}
				var results []*svc.DescribeMetricFiltersOutput
				p := svc.NewDescribeMetricFiltersPaginator(client, input)
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
		"describe-queries": {
			Name:   "describe-queries",
			Fields: fields_describe_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQueriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_queries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQueries(ctx, input)
			},
		},
		"describe-query-definitions": {
			Name:   "describe-query-definitions",
			Fields: fields_describe_query_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQueryDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_query_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQueryDefinitions(ctx, input)
			},
		},
		"describe-resource-policies": {
			Name:   "describe-resource-policies",
			Fields: fields_describe_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicies(ctx, input)
			},
		},
		"describe-subscription-filters": {
			Name:   "describe-subscription-filters",
			Fields: fields_describe_subscription_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubscriptionFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_subscription_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSubscriptionFilters(ctx, input)
				}
				var results []*svc.DescribeSubscriptionFiltersOutput
				p := svc.NewDescribeSubscriptionFiltersPaginator(client, input)
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
		"disassociate-kms-key": {
			Name:   "disassociate-kms-key",
			Fields: fields_disassociate_kms_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateKmsKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_kms_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateKmsKey(ctx, input)
			},
		},
		"disassociate-source-from-s3-table-integration": {
			Name:   "disassociate-source-from-s3-table-integration",
			Fields: fields_disassociate_source_from_s3_table_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSourceFromS3TableIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_source_from_s3_table_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSourceFromS3TableIntegration(ctx, input)
			},
		},
		"filter-log-events": {
			Name:   "filter-log-events",
			Fields: fields_filter_log_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FilterLogEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_filter_log_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.FilterLogEvents(ctx, input)
				}
				var results []*svc.FilterLogEventsOutput
				p := svc.NewFilterLogEventsPaginator(client, input)
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
		"get-data-protection-policy": {
			Name:   "get-data-protection-policy",
			Fields: fields_get_data_protection_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataProtectionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_protection_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataProtectionPolicy(ctx, input)
			},
		},
		"get-delivery": {
			Name:   "get-delivery",
			Fields: fields_get_delivery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeliveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delivery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDelivery(ctx, input)
			},
		},
		"get-delivery-destination": {
			Name:   "get-delivery-destination",
			Fields: fields_get_delivery_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeliveryDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delivery_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeliveryDestination(ctx, input)
			},
		},
		"get-delivery-destination-policy": {
			Name:   "get-delivery-destination-policy",
			Fields: fields_get_delivery_destination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeliveryDestinationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delivery_destination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeliveryDestinationPolicy(ctx, input)
			},
		},
		"get-delivery-source": {
			Name:   "get-delivery-source",
			Fields: fields_get_delivery_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeliverySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delivery_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeliverySource(ctx, input)
			},
		},
		"get-integration": {
			Name:   "get-integration",
			Fields: fields_get_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegration(ctx, input)
			},
		},
		"get-log-anomaly-detector": {
			Name:   "get-log-anomaly-detector",
			Fields: fields_get_log_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogAnomalyDetector(ctx, input)
			},
		},
		"get-log-events": {
			Name:   "get-log-events",
			Fields: fields_get_log_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_log_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLogEvents(ctx, input)
				}
				var results []*svc.GetLogEventsOutput
				p := svc.NewGetLogEventsPaginator(client, input)
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
		"get-log-fields": {
			Name:   "get-log-fields",
			Fields: fields_get_log_fields,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogFieldsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_fields, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogFields(ctx, input)
			},
		},
		"get-log-group-fields": {
			Name:   "get-log-group-fields",
			Fields: fields_get_log_group_fields,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogGroupFieldsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_group_fields, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogGroupFields(ctx, input)
			},
		},
		"get-log-object": {
			Name:   "get-log-object",
			Fields: fields_get_log_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogObject(ctx, input)
			},
		},
		"get-log-record": {
			Name:   "get-log-record",
			Fields: fields_get_log_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogRecord(ctx, input)
			},
		},
		"get-query-results": {
			Name:   "get-query-results",
			Fields: fields_get_query_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryResults(ctx, input)
			},
		},
		"get-scheduled-query": {
			Name:   "get-scheduled-query",
			Fields: fields_get_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScheduledQuery(ctx, input)
			},
		},
		"get-scheduled-query-history": {
			Name:   "get-scheduled-query-history",
			Fields: fields_get_scheduled_query_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScheduledQueryHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_scheduled_query_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetScheduledQueryHistory(ctx, input)
				}
				var results []*svc.GetScheduledQueryHistoryOutput
				p := svc.NewGetScheduledQueryHistoryPaginator(client, input)
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
		"get-transformer": {
			Name:   "get-transformer",
			Fields: fields_get_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransformer(ctx, input)
			},
		},
		"list-aggregate-log-group-summaries": {
			Name:   "list-aggregate-log-group-summaries",
			Fields: fields_list_aggregate_log_group_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAggregateLogGroupSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aggregate_log_group_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAggregateLogGroupSummaries(ctx, input)
				}
				var results []*svc.ListAggregateLogGroupSummariesOutput
				p := svc.NewListAggregateLogGroupSummariesPaginator(client, input)
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
		"list-anomalies": {
			Name:   "list-anomalies",
			Fields: fields_list_anomalies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnomaliesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_anomalies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnomalies(ctx, input)
				}
				var results []*svc.ListAnomaliesOutput
				p := svc.NewListAnomaliesPaginator(client, input)
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
		"list-integrations": {
			Name:   "list-integrations",
			Fields: fields_list_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIntegrations(ctx, input)
			},
		},
		"list-log-anomaly-detectors": {
			Name:   "list-log-anomaly-detectors",
			Fields: fields_list_log_anomaly_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogAnomalyDetectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_log_anomaly_detectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLogAnomalyDetectors(ctx, input)
				}
				var results []*svc.ListLogAnomalyDetectorsOutput
				p := svc.NewListLogAnomalyDetectorsPaginator(client, input)
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
		"list-log-groups": {
			Name:   "list-log-groups",
			Fields: fields_list_log_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_log_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLogGroups(ctx, input)
			},
		},
		"list-log-groups-for-query": {
			Name:   "list-log-groups-for-query",
			Fields: fields_list_log_groups_for_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogGroupsForQueryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_log_groups_for_query, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLogGroupsForQuery(ctx, input)
				}
				var results []*svc.ListLogGroupsForQueryOutput
				p := svc.NewListLogGroupsForQueryPaginator(client, input)
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
		"list-scheduled-queries": {
			Name:   "list-scheduled-queries",
			Fields: fields_list_scheduled_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScheduledQueriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scheduled_queries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScheduledQueries(ctx, input)
				}
				var results []*svc.ListScheduledQueriesOutput
				p := svc.NewListScheduledQueriesPaginator(client, input)
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
		"list-sources-for-s3-table-integration": {
			Name:   "list-sources-for-s3-table-integration",
			Fields: fields_list_sources_for_s3_table_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourcesForS3TableIntegrationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sources_for_s3_table_integration, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourcesForS3TableIntegration(ctx, input)
				}
				var results []*svc.ListSourcesForS3TableIntegrationOutput
				p := svc.NewListSourcesForS3TableIntegrationPaginator(client, input)
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
		"list-tags-log-group": {
			Name:   "list-tags-log-group",
			Fields: fields_list_tags_log_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsLogGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_log_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsLogGroup(ctx, input)
			},
		},
		"put-account-policy": {
			Name:   "put-account-policy",
			Fields: fields_put_account_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountPolicy(ctx, input)
			},
		},
		"put-bearer-token-authentication": {
			Name:   "put-bearer-token-authentication",
			Fields: fields_put_bearer_token_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBearerTokenAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bearer_token_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBearerTokenAuthentication(ctx, input)
			},
		},
		"put-data-protection-policy": {
			Name:   "put-data-protection-policy",
			Fields: fields_put_data_protection_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataProtectionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_protection_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataProtectionPolicy(ctx, input)
			},
		},
		"put-delivery-destination": {
			Name:   "put-delivery-destination",
			Fields: fields_put_delivery_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDeliveryDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_delivery_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDeliveryDestination(ctx, input)
			},
		},
		"put-delivery-destination-policy": {
			Name:   "put-delivery-destination-policy",
			Fields: fields_put_delivery_destination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDeliveryDestinationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_delivery_destination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDeliveryDestinationPolicy(ctx, input)
			},
		},
		"put-delivery-source": {
			Name:   "put-delivery-source",
			Fields: fields_put_delivery_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDeliverySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_delivery_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDeliverySource(ctx, input)
			},
		},
		"put-destination": {
			Name:   "put-destination",
			Fields: fields_put_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDestination(ctx, input)
			},
		},
		"put-destination-policy": {
			Name:   "put-destination-policy",
			Fields: fields_put_destination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDestinationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_destination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDestinationPolicy(ctx, input)
			},
		},
		"put-index-policy": {
			Name:   "put-index-policy",
			Fields: fields_put_index_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIndexPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_index_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIndexPolicy(ctx, input)
			},
		},
		"put-integration": {
			Name:   "put-integration",
			Fields: fields_put_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIntegration(ctx, input)
			},
		},
		"put-log-events": {
			Name:   "put-log-events",
			Fields: fields_put_log_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLogEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_log_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLogEvents(ctx, input)
			},
		},
		"put-log-group-deletion-protection": {
			Name:   "put-log-group-deletion-protection",
			Fields: fields_put_log_group_deletion_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLogGroupDeletionProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_log_group_deletion_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLogGroupDeletionProtection(ctx, input)
			},
		},
		"put-metric-filter": {
			Name:   "put-metric-filter",
			Fields: fields_put_metric_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetricFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metric_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetricFilter(ctx, input)
			},
		},
		"put-query-definition": {
			Name:   "put-query-definition",
			Fields: fields_put_query_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutQueryDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_query_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutQueryDefinition(ctx, input)
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
		"put-retention-policy": {
			Name:   "put-retention-policy",
			Fields: fields_put_retention_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRetentionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_retention_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRetentionPolicy(ctx, input)
			},
		},
		"put-subscription-filter": {
			Name:   "put-subscription-filter",
			Fields: fields_put_subscription_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSubscriptionFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_subscription_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSubscriptionFilter(ctx, input)
			},
		},
		"put-transformer": {
			Name:   "put-transformer",
			Fields: fields_put_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTransformer(ctx, input)
			},
		},
		"start-live-tail": {
			Name:   "start-live-tail",
			Fields: fields_start_live_tail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartLiveTailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_live_tail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartLiveTail(ctx, input)
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
		"tag-log-group": {
			Name:   "tag-log-group",
			Fields: fields_tag_log_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagLogGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_log_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagLogGroup(ctx, input)
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
		"test-metric-filter": {
			Name:   "test-metric-filter",
			Fields: fields_test_metric_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestMetricFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_metric_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestMetricFilter(ctx, input)
			},
		},
		"test-transformer": {
			Name:   "test-transformer",
			Fields: fields_test_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestTransformer(ctx, input)
			},
		},
		"untag-log-group": {
			Name:   "untag-log-group",
			Fields: fields_untag_log_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagLogGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_log_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagLogGroup(ctx, input)
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
		"update-anomaly": {
			Name:   "update-anomaly",
			Fields: fields_update_anomaly,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnomalyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_anomaly, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnomaly(ctx, input)
			},
		},
		"update-delivery-configuration": {
			Name:   "update-delivery-configuration",
			Fields: fields_update_delivery_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeliveryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_delivery_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeliveryConfiguration(ctx, input)
			},
		},
		"update-log-anomaly-detector": {
			Name:   "update-log-anomaly-detector",
			Fields: fields_update_log_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLogAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_log_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLogAnomalyDetector(ctx, input)
			},
		},
		"update-scheduled-query": {
			Name:   "update-scheduled-query",
			Fields: fields_update_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScheduledQuery(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudwatchlogs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
