package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var fields_batch_execute_statement = []leanruntime.Field{
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "Statements", Flag: "statements", Type: "[]types.BatchStatementRequest", Required: true},
}

var fields_batch_get_item = []leanruntime.Field{
	{Name: "RequestItems", Flag: "request-items", Type: "map[string]types.KeysAndAttributes", Required: true},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
}

var fields_batch_write_item = []leanruntime.Field{
	{Name: "RequestItems", Flag: "request-items", Type: "map[string][]types.WriteRequest", Required: true},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ReturnItemCollectionMetrics", Flag: "return-item-collection-metrics", Type: "types.ReturnItemCollectionMetrics", Required: false},
}

var fields_create_backup = []leanruntime.Field{
	{Name: "BackupName", Flag: "backup-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_create_global_table = []leanruntime.Field{
	{Name: "GlobalTableName", Flag: "global-table-name", Type: "*string", Required: true},
	{Name: "ReplicationGroup", Flag: "replication-group", Type: "[]types.Replica", Required: true},
}

var fields_create_table = []leanruntime.Field{
	{Name: "AttributeDefinitions", Flag: "attribute-definitions", Type: "[]types.AttributeDefinition", Required: false},
	{Name: "BillingMode", Flag: "billing-mode", Type: "types.BillingMode", Required: false},
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "GlobalSecondaryIndexes", Flag: "global-secondary-indexes", Type: "[]types.GlobalSecondaryIndex", Required: false},
	{Name: "GlobalTableSettingsReplicationMode", Flag: "global-table-settings-replication-mode", Type: "types.GlobalTableSettingsReplicationMode", Required: false},
	{Name: "GlobalTableSourceArn", Flag: "global-table-source-arn", Type: "*string", Required: false},
	{Name: "KeySchema", Flag: "key-schema", Type: "[]types.KeySchemaElement", Required: false},
	{Name: "LocalSecondaryIndexes", Flag: "local-secondary-indexes", Type: "[]types.LocalSecondaryIndex", Required: false},
	{Name: "OnDemandThroughput", Flag: "on-demand-throughput", Type: "*types.OnDemandThroughput", Required: false},
	{Name: "ProvisionedThroughput", Flag: "provisioned-throughput", Type: "*types.ProvisionedThroughput", Required: false},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: false},
	{Name: "SSESpecification", Flag: "sse-specification", Type: "*types.SSESpecification", Required: false},
	{Name: "StreamSpecification", Flag: "stream-specification", Type: "*types.StreamSpecification", Required: false},
	{Name: "TableClass", Flag: "table-class", Type: "types.TableClass", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WarmThroughput", Flag: "warm-throughput", Type: "*types.WarmThroughput", Required: false},
}

var fields_delete_backup = []leanruntime.Field{
	{Name: "BackupArn", Flag: "backup-arn", Type: "*string", Required: true},
}

var fields_delete_item = []leanruntime.Field{
	{Name: "ConditionExpression", Flag: "condition-expression", Type: "*string", Required: false},
	{Name: "ConditionalOperator", Flag: "conditional-operator", Type: "types.ConditionalOperator", Required: false},
	{Name: "Expected", Flag: "expected", Type: "map[string]types.ExpectedAttributeValue", Required: false},
	{Name: "ExpressionAttributeNames", Flag: "expression-attribute-names", Type: "map[string]string", Required: false},
	{Name: "ExpressionAttributeValues", Flag: "expression-attribute-values", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "Key", Flag: "key", Type: "map[string]types.AttributeValue", Required: true},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ReturnItemCollectionMetrics", Flag: "return-item-collection-metrics", Type: "types.ReturnItemCollectionMetrics", Required: false},
	{Name: "ReturnValues", Flag: "return-values", Type: "types.ReturnValue", Required: false},
	{Name: "ReturnValuesOnConditionCheckFailure", Flag: "return-values-on-condition-check-failure", Type: "types.ReturnValuesOnConditionCheckFailure", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_table = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_backup = []leanruntime.Field{
	{Name: "BackupArn", Flag: "backup-arn", Type: "*string", Required: true},
}

var fields_describe_continuous_backups = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_contributor_insights = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_endpoints = []leanruntime.Field{}

var fields_describe_export = []leanruntime.Field{
	{Name: "ExportArn", Flag: "export-arn", Type: "*string", Required: true},
}

var fields_describe_global_table = []leanruntime.Field{
	{Name: "GlobalTableName", Flag: "global-table-name", Type: "*string", Required: true},
}

var fields_describe_global_table_settings = []leanruntime.Field{
	{Name: "GlobalTableName", Flag: "global-table-name", Type: "*string", Required: true},
}

var fields_describe_import = []leanruntime.Field{
	{Name: "ImportArn", Flag: "import-arn", Type: "*string", Required: true},
}

var fields_describe_kinesis_streaming_destination = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_limits = []leanruntime.Field{}

var fields_describe_table = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_table_replica_auto_scaling = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_describe_time_to_live = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_disable_kinesis_streaming_destination = []leanruntime.Field{
	{Name: "EnableKinesisStreamingConfiguration", Flag: "enable-kinesis-streaming-configuration", Type: "*types.EnableKinesisStreamingConfiguration", Required: false},
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_enable_kinesis_streaming_destination = []leanruntime.Field{
	{Name: "EnableKinesisStreamingConfiguration", Flag: "enable-kinesis-streaming-configuration", Type: "*types.EnableKinesisStreamingConfiguration", Required: false},
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_execute_statement = []leanruntime.Field{
	{Name: "ConsistentRead", Flag: "consistent-read", Type: "*bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.AttributeValue", Required: false},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ReturnValuesOnConditionCheckFailure", Flag: "return-values-on-condition-check-failure", Type: "types.ReturnValuesOnConditionCheckFailure", Required: false},
	{Name: "Statement", Flag: "statement", Type: "*string", Required: true},
}

var fields_execute_transaction = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "TransactStatements", Flag: "transact-statements", Type: "[]types.ParameterizedStatement", Required: true},
}

var fields_export_table_to_point_in_time = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExportFormat", Flag: "export-format", Type: "types.ExportFormat", Required: false},
	{Name: "ExportTime", Flag: "export-time", Type: "*time.Time", Required: false},
	{Name: "ExportType", Flag: "export-type", Type: "types.ExportType", Required: false},
	{Name: "IncrementalExportSpecification", Flag: "incremental-export-specification", Type: "*types.IncrementalExportSpecification", Required: false},
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: true},
	{Name: "S3BucketOwner", Flag: "s3-bucket-owner", Type: "*string", Required: false},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: false},
	{Name: "S3SseAlgorithm", Flag: "s3-sse-algorithm", Type: "types.S3SseAlgorithm", Required: false},
	{Name: "S3SseKmsKeyId", Flag: "s3-sse-kms-key-id", Type: "*string", Required: false},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_item = []leanruntime.Field{
	{Name: "AttributesToGet", Flag: "attributes-to-get", Type: "[]string", Required: false},
	{Name: "ConsistentRead", Flag: "consistent-read", Type: "*bool", Required: false},
	{Name: "ExpressionAttributeNames", Flag: "expression-attribute-names", Type: "map[string]string", Required: false},
	{Name: "Key", Flag: "key", Type: "map[string]types.AttributeValue", Required: true},
	{Name: "ProjectionExpression", Flag: "projection-expression", Type: "*string", Required: false},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_import_table = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InputCompressionType", Flag: "input-compression-type", Type: "types.InputCompressionType", Required: false},
	{Name: "InputFormat", Flag: "input-format", Type: "types.InputFormat", Required: true},
	{Name: "InputFormatOptions", Flag: "input-format-options", Type: "*types.InputFormatOptions", Required: false},
	{Name: "S3BucketSource", Flag: "s3-bucket-source", Type: "*types.S3BucketSource", Required: true},
	{Name: "TableCreationParameters", Flag: "table-creation-parameters", Type: "*types.TableCreationParameters", Required: true},
}

var fields_list_backups = []leanruntime.Field{
	{Name: "BackupType", Flag: "backup-type", Type: "types.BackupTypeFilter", Required: false},
	{Name: "ExclusiveStartBackupArn", Flag: "exclusive-start-backup-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: false},
	{Name: "TimeRangeLowerBound", Flag: "time-range-lower-bound", Type: "*time.Time", Required: false},
	{Name: "TimeRangeUpperBound", Flag: "time-range-upper-bound", Type: "*time.Time", Required: false},
}

var fields_list_contributor_insights = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: false},
}

var fields_list_exports = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: false},
}

var fields_list_global_tables = []leanruntime.Field{
	{Name: "ExclusiveStartGlobalTableName", Flag: "exclusive-start-global-table-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
}

var fields_list_imports = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: false},
}

var fields_list_tables = []leanruntime.Field{
	{Name: "ExclusiveStartTableName", Flag: "exclusive-start-table-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
}

var fields_list_tags_of_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_item = []leanruntime.Field{
	{Name: "ConditionExpression", Flag: "condition-expression", Type: "*string", Required: false},
	{Name: "ConditionalOperator", Flag: "conditional-operator", Type: "types.ConditionalOperator", Required: false},
	{Name: "Expected", Flag: "expected", Type: "map[string]types.ExpectedAttributeValue", Required: false},
	{Name: "ExpressionAttributeNames", Flag: "expression-attribute-names", Type: "map[string]string", Required: false},
	{Name: "ExpressionAttributeValues", Flag: "expression-attribute-values", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "Item", Flag: "item", Type: "map[string]types.AttributeValue", Required: true},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ReturnItemCollectionMetrics", Flag: "return-item-collection-metrics", Type: "types.ReturnItemCollectionMetrics", Required: false},
	{Name: "ReturnValues", Flag: "return-values", Type: "types.ReturnValue", Required: false},
	{Name: "ReturnValuesOnConditionCheckFailure", Flag: "return-values-on-condition-check-failure", Type: "types.ReturnValuesOnConditionCheckFailure", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "ConfirmRemoveSelfResourceAccess", Flag: "confirm-remove-self-resource-access", Type: "bool", Required: false},
	{Name: "ExpectedRevisionId", Flag: "expected-revision-id", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_query = []leanruntime.Field{
	{Name: "AttributesToGet", Flag: "attributes-to-get", Type: "[]string", Required: false},
	{Name: "ConditionalOperator", Flag: "conditional-operator", Type: "types.ConditionalOperator", Required: false},
	{Name: "ConsistentRead", Flag: "consistent-read", Type: "*bool", Required: false},
	{Name: "ExclusiveStartKey", Flag: "exclusive-start-key", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "ExpressionAttributeNames", Flag: "expression-attribute-names", Type: "map[string]string", Required: false},
	{Name: "ExpressionAttributeValues", Flag: "expression-attribute-values", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "KeyConditionExpression", Flag: "key-condition-expression", Type: "*string", Required: false},
	{Name: "KeyConditions", Flag: "key-conditions", Type: "map[string]types.Condition", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "ProjectionExpression", Flag: "projection-expression", Type: "*string", Required: false},
	{Name: "QueryFilter", Flag: "query-filter", Type: "map[string]types.Condition", Required: false},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ScanIndexForward", Flag: "scan-index-forward", Type: "*bool", Required: false},
	{Name: "Select", Flag: "select", Type: "types.Select", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_restore_table_from_backup = []leanruntime.Field{
	{Name: "BackupArn", Flag: "backup-arn", Type: "*string", Required: true},
	{Name: "BillingModeOverride", Flag: "billing-mode-override", Type: "types.BillingMode", Required: false},
	{Name: "GlobalSecondaryIndexOverride", Flag: "global-secondary-index-override", Type: "[]types.GlobalSecondaryIndex", Required: false},
	{Name: "LocalSecondaryIndexOverride", Flag: "local-secondary-index-override", Type: "[]types.LocalSecondaryIndex", Required: false},
	{Name: "OnDemandThroughputOverride", Flag: "on-demand-throughput-override", Type: "*types.OnDemandThroughput", Required: false},
	{Name: "ProvisionedThroughputOverride", Flag: "provisioned-throughput-override", Type: "*types.ProvisionedThroughput", Required: false},
	{Name: "SSESpecificationOverride", Flag: "sse-specification-override", Type: "*types.SSESpecification", Required: false},
	{Name: "TargetTableName", Flag: "target-table-name", Type: "*string", Required: true},
}

var fields_restore_table_to_point_in_time = []leanruntime.Field{
	{Name: "BillingModeOverride", Flag: "billing-mode-override", Type: "types.BillingMode", Required: false},
	{Name: "GlobalSecondaryIndexOverride", Flag: "global-secondary-index-override", Type: "[]types.GlobalSecondaryIndex", Required: false},
	{Name: "LocalSecondaryIndexOverride", Flag: "local-secondary-index-override", Type: "[]types.LocalSecondaryIndex", Required: false},
	{Name: "OnDemandThroughputOverride", Flag: "on-demand-throughput-override", Type: "*types.OnDemandThroughput", Required: false},
	{Name: "ProvisionedThroughputOverride", Flag: "provisioned-throughput-override", Type: "*types.ProvisionedThroughput", Required: false},
	{Name: "RestoreDateTime", Flag: "restore-date-time", Type: "*time.Time", Required: false},
	{Name: "SSESpecificationOverride", Flag: "sse-specification-override", Type: "*types.SSESpecification", Required: false},
	{Name: "SourceTableArn", Flag: "source-table-arn", Type: "*string", Required: false},
	{Name: "SourceTableName", Flag: "source-table-name", Type: "*string", Required: false},
	{Name: "TargetTableName", Flag: "target-table-name", Type: "*string", Required: true},
	{Name: "UseLatestRestorableTime", Flag: "use-latest-restorable-time", Type: "*bool", Required: false},
}

var fields_scan = []leanruntime.Field{
	{Name: "AttributesToGet", Flag: "attributes-to-get", Type: "[]string", Required: false},
	{Name: "ConditionalOperator", Flag: "conditional-operator", Type: "types.ConditionalOperator", Required: false},
	{Name: "ConsistentRead", Flag: "consistent-read", Type: "*bool", Required: false},
	{Name: "ExclusiveStartKey", Flag: "exclusive-start-key", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "ExpressionAttributeNames", Flag: "expression-attribute-names", Type: "map[string]string", Required: false},
	{Name: "ExpressionAttributeValues", Flag: "expression-attribute-values", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "ProjectionExpression", Flag: "projection-expression", Type: "*string", Required: false},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ScanFilter", Flag: "scan-filter", Type: "map[string]types.Condition", Required: false},
	{Name: "Segment", Flag: "segment", Type: "*int32", Required: false},
	{Name: "Select", Flag: "select", Type: "types.Select", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TotalSegments", Flag: "total-segments", Type: "*int32", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_transact_get_items = []leanruntime.Field{
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "TransactItems", Flag: "transact-items", Type: "[]types.TransactGetItem", Required: true},
}

var fields_transact_write_items = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ReturnItemCollectionMetrics", Flag: "return-item-collection-metrics", Type: "types.ReturnItemCollectionMetrics", Required: false},
	{Name: "TransactItems", Flag: "transact-items", Type: "[]types.TransactWriteItem", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_continuous_backups = []leanruntime.Field{
	{Name: "PointInTimeRecoverySpecification", Flag: "point-in-time-recovery-specification", Type: "*types.PointInTimeRecoverySpecification", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_contributor_insights = []leanruntime.Field{
	{Name: "ContributorInsightsAction", Flag: "contributor-insights-action", Type: "types.ContributorInsightsAction", Required: true},
	{Name: "ContributorInsightsMode", Flag: "contributor-insights-mode", Type: "types.ContributorInsightsMode", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_global_table = []leanruntime.Field{
	{Name: "GlobalTableName", Flag: "global-table-name", Type: "*string", Required: true},
	{Name: "ReplicaUpdates", Flag: "replica-updates", Type: "[]types.ReplicaUpdate", Required: true},
}

var fields_update_global_table_settings = []leanruntime.Field{
	{Name: "GlobalTableBillingMode", Flag: "global-table-billing-mode", Type: "types.BillingMode", Required: false},
	{Name: "GlobalTableGlobalSecondaryIndexSettingsUpdate", Flag: "global-table-global-secondary-index-settings-update", Type: "[]types.GlobalTableGlobalSecondaryIndexSettingsUpdate", Required: false},
	{Name: "GlobalTableName", Flag: "global-table-name", Type: "*string", Required: true},
	{Name: "GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate", Flag: "global-table-provisioned-write-capacity-auto-scaling-settings-update", Type: "*types.AutoScalingSettingsUpdate", Required: false},
	{Name: "GlobalTableProvisionedWriteCapacityUnits", Flag: "global-table-provisioned-write-capacity-units", Type: "*int64", Required: false},
	{Name: "ReplicaSettingsUpdate", Flag: "replica-settings-update", Type: "[]types.ReplicaSettingsUpdate", Required: false},
}

var fields_update_item = []leanruntime.Field{
	{Name: "AttributeUpdates", Flag: "attribute-updates", Type: "map[string]types.AttributeValueUpdate", Required: false},
	{Name: "ConditionExpression", Flag: "condition-expression", Type: "*string", Required: false},
	{Name: "ConditionalOperator", Flag: "conditional-operator", Type: "types.ConditionalOperator", Required: false},
	{Name: "Expected", Flag: "expected", Type: "map[string]types.ExpectedAttributeValue", Required: false},
	{Name: "ExpressionAttributeNames", Flag: "expression-attribute-names", Type: "map[string]string", Required: false},
	{Name: "ExpressionAttributeValues", Flag: "expression-attribute-values", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "Key", Flag: "key", Type: "map[string]types.AttributeValue", Required: true},
	{Name: "ReturnConsumedCapacity", Flag: "return-consumed-capacity", Type: "types.ReturnConsumedCapacity", Required: false},
	{Name: "ReturnItemCollectionMetrics", Flag: "return-item-collection-metrics", Type: "types.ReturnItemCollectionMetrics", Required: false},
	{Name: "ReturnValues", Flag: "return-values", Type: "types.ReturnValue", Required: false},
	{Name: "ReturnValuesOnConditionCheckFailure", Flag: "return-values-on-condition-check-failure", Type: "types.ReturnValuesOnConditionCheckFailure", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "UpdateExpression", Flag: "update-expression", Type: "*string", Required: false},
}

var fields_update_kinesis_streaming_destination = []leanruntime.Field{
	{Name: "StreamArn", Flag: "stream-arn", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "UpdateKinesisStreamingConfiguration", Flag: "update-kinesis-streaming-configuration", Type: "*types.UpdateKinesisStreamingConfiguration", Required: false},
}

var fields_update_table = []leanruntime.Field{
	{Name: "AttributeDefinitions", Flag: "attribute-definitions", Type: "[]types.AttributeDefinition", Required: false},
	{Name: "BillingMode", Flag: "billing-mode", Type: "types.BillingMode", Required: false},
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "GlobalSecondaryIndexUpdates", Flag: "global-secondary-index-updates", Type: "[]types.GlobalSecondaryIndexUpdate", Required: false},
	{Name: "GlobalTableSettingsReplicationMode", Flag: "global-table-settings-replication-mode", Type: "types.GlobalTableSettingsReplicationMode", Required: false},
	{Name: "GlobalTableWitnessUpdates", Flag: "global-table-witness-updates", Type: "[]types.GlobalTableWitnessGroupUpdate", Required: false},
	{Name: "MultiRegionConsistency", Flag: "multi-region-consistency", Type: "types.MultiRegionConsistency", Required: false},
	{Name: "OnDemandThroughput", Flag: "on-demand-throughput", Type: "*types.OnDemandThroughput", Required: false},
	{Name: "ProvisionedThroughput", Flag: "provisioned-throughput", Type: "*types.ProvisionedThroughput", Required: false},
	{Name: "ReplicaUpdates", Flag: "replica-updates", Type: "[]types.ReplicationGroupUpdate", Required: false},
	{Name: "SSESpecification", Flag: "sse-specification", Type: "*types.SSESpecification", Required: false},
	{Name: "StreamSpecification", Flag: "stream-specification", Type: "*types.StreamSpecification", Required: false},
	{Name: "TableClass", Flag: "table-class", Type: "types.TableClass", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "WarmThroughput", Flag: "warm-throughput", Type: "*types.WarmThroughput", Required: false},
}

var fields_update_table_replica_auto_scaling = []leanruntime.Field{
	{Name: "GlobalSecondaryIndexUpdates", Flag: "global-secondary-index-updates", Type: "[]types.GlobalSecondaryIndexAutoScalingUpdate", Required: false},
	{Name: "ProvisionedWriteCapacityAutoScalingUpdate", Flag: "provisioned-write-capacity-auto-scaling-update", Type: "*types.AutoScalingSettingsUpdate", Required: false},
	{Name: "ReplicaUpdates", Flag: "replica-updates", Type: "[]types.ReplicaAutoScalingUpdate", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_time_to_live = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TimeToLiveSpecification", Flag: "time-to-live-specification", Type: "*types.TimeToLiveSpecification", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-execute-statement": {
			Name:   "batch-execute-statement",
			Fields: fields_batch_execute_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchExecuteStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_execute_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchExecuteStatement(ctx, input)
			},
		},
		"batch-get-item": {
			Name:   "batch-get-item",
			Fields: fields_batch_get_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetItem(ctx, input)
			},
		},
		"batch-write-item": {
			Name:   "batch-write-item",
			Fields: fields_batch_write_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchWriteItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_write_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchWriteItem(ctx, input)
			},
		},
		"create-backup": {
			Name:   "create-backup",
			Fields: fields_create_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackup(ctx, input)
			},
		},
		"create-global-table": {
			Name:   "create-global-table",
			Fields: fields_create_global_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlobalTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_global_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlobalTable(ctx, input)
			},
		},
		"create-table": {
			Name:   "create-table",
			Fields: fields_create_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTable(ctx, input)
			},
		},
		"delete-backup": {
			Name:   "delete-backup",
			Fields: fields_delete_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackup(ctx, input)
			},
		},
		"delete-item": {
			Name:   "delete-item",
			Fields: fields_delete_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteItem(ctx, input)
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
		"delete-table": {
			Name:   "delete-table",
			Fields: fields_delete_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTable(ctx, input)
			},
		},
		"describe-backup": {
			Name:   "describe-backup",
			Fields: fields_describe_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBackup(ctx, input)
			},
		},
		"describe-continuous-backups": {
			Name:   "describe-continuous-backups",
			Fields: fields_describe_continuous_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContinuousBackupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_continuous_backups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContinuousBackups(ctx, input)
			},
		},
		"describe-contributor-insights": {
			Name:   "describe-contributor-insights",
			Fields: fields_describe_contributor_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContributorInsightsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_contributor_insights, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContributorInsights(ctx, input)
			},
		},
		"describe-endpoints": {
			Name:   "describe-endpoints",
			Fields: fields_describe_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoints(ctx, input)
			},
		},
		"describe-export": {
			Name:   "describe-export",
			Fields: fields_describe_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExport(ctx, input)
			},
		},
		"describe-global-table": {
			Name:   "describe-global-table",
			Fields: fields_describe_global_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGlobalTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_global_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGlobalTable(ctx, input)
			},
		},
		"describe-global-table-settings": {
			Name:   "describe-global-table-settings",
			Fields: fields_describe_global_table_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGlobalTableSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_global_table_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGlobalTableSettings(ctx, input)
			},
		},
		"describe-import": {
			Name:   "describe-import",
			Fields: fields_describe_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImport(ctx, input)
			},
		},
		"describe-kinesis-streaming-destination": {
			Name:   "describe-kinesis-streaming-destination",
			Fields: fields_describe_kinesis_streaming_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKinesisStreamingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_kinesis_streaming_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKinesisStreamingDestination(ctx, input)
			},
		},
		"describe-limits": {
			Name:   "describe-limits",
			Fields: fields_describe_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLimits(ctx, input)
			},
		},
		"describe-table": {
			Name:   "describe-table",
			Fields: fields_describe_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTable(ctx, input)
			},
		},
		"describe-table-replica-auto-scaling": {
			Name:   "describe-table-replica-auto-scaling",
			Fields: fields_describe_table_replica_auto_scaling,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTableReplicaAutoScalingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_table_replica_auto_scaling, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTableReplicaAutoScaling(ctx, input)
			},
		},
		"describe-time-to-live": {
			Name:   "describe-time-to-live",
			Fields: fields_describe_time_to_live,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTimeToLiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_time_to_live, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTimeToLive(ctx, input)
			},
		},
		"disable-kinesis-streaming-destination": {
			Name:   "disable-kinesis-streaming-destination",
			Fields: fields_disable_kinesis_streaming_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableKinesisStreamingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_kinesis_streaming_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableKinesisStreamingDestination(ctx, input)
			},
		},
		"enable-kinesis-streaming-destination": {
			Name:   "enable-kinesis-streaming-destination",
			Fields: fields_enable_kinesis_streaming_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableKinesisStreamingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_kinesis_streaming_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableKinesisStreamingDestination(ctx, input)
			},
		},
		"execute-statement": {
			Name:   "execute-statement",
			Fields: fields_execute_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteStatement(ctx, input)
			},
		},
		"execute-transaction": {
			Name:   "execute-transaction",
			Fields: fields_execute_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteTransaction(ctx, input)
			},
		},
		"export-table-to-point-in-time": {
			Name:   "export-table-to-point-in-time",
			Fields: fields_export_table_to_point_in_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportTableToPointInTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_table_to_point_in_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportTableToPointInTime(ctx, input)
			},
		},
		"get-item": {
			Name:   "get-item",
			Fields: fields_get_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetItem(ctx, input)
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
		"import-table": {
			Name:   "import-table",
			Fields: fields_import_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportTable(ctx, input)
			},
		},
		"list-backups": {
			Name:   "list-backups",
			Fields: fields_list_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_backups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBackups(ctx, input)
			},
		},
		"list-contributor-insights": {
			Name:   "list-contributor-insights",
			Fields: fields_list_contributor_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContributorInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contributor_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContributorInsights(ctx, input)
				}
				var results []*svc.ListContributorInsightsOutput
				p := svc.NewListContributorInsightsPaginator(client, input)
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
		"list-exports": {
			Name:   "list-exports",
			Fields: fields_list_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExports(ctx, input)
				}
				var results []*svc.ListExportsOutput
				p := svc.NewListExportsPaginator(client, input)
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
		"list-global-tables": {
			Name:   "list-global-tables",
			Fields: fields_list_global_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGlobalTablesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_global_tables, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGlobalTables(ctx, input)
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
		"list-tables": {
			Name:   "list-tables",
			Fields: fields_list_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTables(ctx, input)
				}
				var results []*svc.ListTablesOutput
				p := svc.NewListTablesPaginator(client, input)
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
		"list-tags-of-resource": {
			Name:   "list-tags-of-resource",
			Fields: fields_list_tags_of_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsOfResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_of_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsOfResource(ctx, input)
			},
		},
		"put-item": {
			Name:   "put-item",
			Fields: fields_put_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutItem(ctx, input)
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
		"query": {
			Name:   "query",
			Fields: fields_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_query, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Query(ctx, input)
				}
				var results []*svc.QueryOutput
				p := svc.NewQueryPaginator(client, input)
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
		"restore-table-from-backup": {
			Name:   "restore-table-from-backup",
			Fields: fields_restore_table_from_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreTableFromBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_table_from_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreTableFromBackup(ctx, input)
			},
		},
		"restore-table-to-point-in-time": {
			Name:   "restore-table-to-point-in-time",
			Fields: fields_restore_table_to_point_in_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreTableToPointInTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_table_to_point_in_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreTableToPointInTime(ctx, input)
			},
		},
		"scan": {
			Name:   "scan",
			Fields: fields_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ScanInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_scan, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Scan(ctx, input)
				}
				var results []*svc.ScanOutput
				p := svc.NewScanPaginator(client, input)
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
		"transact-get-items": {
			Name:   "transact-get-items",
			Fields: fields_transact_get_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransactGetItemsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transact_get_items, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransactGetItems(ctx, input)
			},
		},
		"transact-write-items": {
			Name:   "transact-write-items",
			Fields: fields_transact_write_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransactWriteItemsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transact_write_items, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransactWriteItems(ctx, input)
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
		"update-continuous-backups": {
			Name:   "update-continuous-backups",
			Fields: fields_update_continuous_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContinuousBackupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_continuous_backups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContinuousBackups(ctx, input)
			},
		},
		"update-contributor-insights": {
			Name:   "update-contributor-insights",
			Fields: fields_update_contributor_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContributorInsightsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contributor_insights, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContributorInsights(ctx, input)
			},
		},
		"update-global-table": {
			Name:   "update-global-table",
			Fields: fields_update_global_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalTable(ctx, input)
			},
		},
		"update-global-table-settings": {
			Name:   "update-global-table-settings",
			Fields: fields_update_global_table_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalTableSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_table_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalTableSettings(ctx, input)
			},
		},
		"update-item": {
			Name:   "update-item",
			Fields: fields_update_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateItem(ctx, input)
			},
		},
		"update-kinesis-streaming-destination": {
			Name:   "update-kinesis-streaming-destination",
			Fields: fields_update_kinesis_streaming_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKinesisStreamingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kinesis_streaming_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKinesisStreamingDestination(ctx, input)
			},
		},
		"update-table": {
			Name:   "update-table",
			Fields: fields_update_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTable(ctx, input)
			},
		},
		"update-table-replica-auto-scaling": {
			Name:   "update-table-replica-auto-scaling",
			Fields: fields_update_table_replica_auto_scaling,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableReplicaAutoScalingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table_replica_auto_scaling, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTableReplicaAutoScaling(ctx, input)
			},
		},
		"update-time-to-live": {
			Name:   "update-time-to-live",
			Fields: fields_update_time_to_live,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTimeToLiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_time_to_live, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTimeToLive(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("dynamodb", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
