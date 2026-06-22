package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/glue"
)

var fields_batch_create_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionInputList", Flag: "partition-input-list", Type: "[]types.PartitionInput", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_batch_delete_connection = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionNameList", Flag: "connection-name-list", Type: "[]string", Required: true},
}

var fields_batch_delete_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionsToDelete", Flag: "partitions-to-delete", Type: "[]types.PartitionValueList", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_batch_delete_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TablesToDelete", Flag: "tables-to-delete", Type: "[]string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_batch_delete_table_version = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "VersionIds", Flag: "version-ids", Type: "[]string", Required: true},
}

var fields_batch_get_blueprints = []leanruntime.Field{
	{Name: "IncludeBlueprint", Flag: "include-blueprint", Type: "*bool", Required: false},
	{Name: "IncludeParameterSpec", Flag: "include-parameter-spec", Type: "*bool", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_batch_get_crawlers = []leanruntime.Field{
	{Name: "CrawlerNames", Flag: "crawler-names", Type: "[]string", Required: true},
}

var fields_batch_get_custom_entity_types = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_batch_get_data_quality_result = []leanruntime.Field{
	{Name: "ResultIds", Flag: "result-ids", Type: "[]string", Required: true},
}

var fields_batch_get_dev_endpoints = []leanruntime.Field{
	{Name: "DevEndpointNames", Flag: "dev-endpoint-names", Type: "[]string", Required: true},
}

var fields_batch_get_jobs = []leanruntime.Field{
	{Name: "JobNames", Flag: "job-names", Type: "[]string", Required: true},
}

var fields_batch_get_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionsToGet", Flag: "partitions-to-get", Type: "[]types.PartitionValueList", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_batch_get_table_optimizer = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchGetTableOptimizerEntry", Required: true},
}

var fields_batch_get_triggers = []leanruntime.Field{
	{Name: "TriggerNames", Flag: "trigger-names", Type: "[]string", Required: true},
}

var fields_batch_get_workflows = []leanruntime.Field{
	{Name: "IncludeGraph", Flag: "include-graph", Type: "*bool", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_batch_put_data_quality_statistic_annotation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InclusionAnnotations", Flag: "inclusion-annotations", Type: "[]types.DatapointInclusionAnnotation", Required: true},
}

var fields_batch_stop_job_run = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobRunIds", Flag: "job-run-ids", Type: "[]string", Required: true},
}

var fields_batch_update_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchUpdatePartitionRequestEntry", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_cancel_data_quality_rule_recommendation_run = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_cancel_data_quality_ruleset_evaluation_run = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_cancel_ml_task_run = []leanruntime.Field{
	{Name: "TaskRunId", Flag: "task-run-id", Type: "*string", Required: true},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_cancel_statement = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "int32", Required: true},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_check_schema_version_validity = []leanruntime.Field{
	{Name: "DataFormat", Flag: "data-format", Type: "types.DataFormat", Required: true},
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*string", Required: true},
}

var fields_create_blueprint = []leanruntime.Field{
	{Name: "BlueprintLocation", Flag: "blueprint-location", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_catalog = []leanruntime.Field{
	{Name: "CatalogInput", Flag: "catalog-input", Type: "*types.CatalogInput", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_classifier = []leanruntime.Field{
	{Name: "CsvClassifier", Flag: "csv-classifier", Type: "*types.CreateCsvClassifierRequest", Required: false},
	{Name: "GrokClassifier", Flag: "grok-classifier", Type: "*types.CreateGrokClassifierRequest", Required: false},
	{Name: "JsonClassifier", Flag: "json-classifier", Type: "*types.CreateJsonClassifierRequest", Required: false},
	{Name: "XMLClassifier", Flag: "xml-classifier", Type: "*types.CreateXMLClassifierRequest", Required: false},
}

var fields_create_column_statistics_task_settings = []leanruntime.Field{
	{Name: "CatalogID", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnNameList", Flag: "column-name-list", Type: "[]string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "SampleSize", Flag: "sample-size", Type: "float64", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_connection = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionInput", Flag: "connection-input", Type: "*types.ConnectionInput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_crawler = []leanruntime.Field{
	{Name: "Classifiers", Flag: "classifiers", Type: "[]string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: false},
	{Name: "CrawlerSecurityConfiguration", Flag: "crawler-security-configuration", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LakeFormationConfiguration", Flag: "lake-formation-configuration", Type: "*types.LakeFormationConfiguration", Required: false},
	{Name: "LineageConfiguration", Flag: "lineage-configuration", Type: "*types.LineageConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecrawlPolicy", Flag: "recrawl-policy", Type: "*types.RecrawlPolicy", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "SchemaChangePolicy", Flag: "schema-change-policy", Type: "*types.SchemaChangePolicy", Required: false},
	{Name: "TablePrefix", Flag: "table-prefix", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "*types.CrawlerTargets", Required: true},
}

var fields_create_custom_entity_type = []leanruntime.Field{
	{Name: "ContextWords", Flag: "context-words", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RegexString", Flag: "regex-string", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_quality_ruleset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataQualitySecurityConfiguration", Flag: "data-quality-security-configuration", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Ruleset", Flag: "ruleset", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetTable", Flag: "target-table", Type: "*types.DataQualityTargetTable", Required: false},
}

var fields_create_database = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseInput", Flag: "database-input", Type: "*types.DatabaseInput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dev_endpoint = []leanruntime.Field{
	{Name: "Arguments", Flag: "arguments", Type: "map[string]string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "ExtraJarsS3Path", Flag: "extra-jars-s3-path", Type: "*string", Required: false},
	{Name: "ExtraPythonLibsS3Path", Flag: "extra-python-libs-s3-path", Type: "*string", Required: false},
	{Name: "GlueVersion", Flag: "glue-version", Type: "*string", Required: false},
	{Name: "NumberOfNodes", Flag: "number-of-nodes", Type: "int32", Required: false},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "PublicKey", Flag: "public-key", Type: "*string", Required: false},
	{Name: "PublicKeys", Flag: "public-keys", Type: "[]string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkerType", Flag: "worker-type", Type: "types.WorkerType", Required: false},
}

var fields_create_glue_identity_center_configuration = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Scopes", Flag: "scopes", Type: "[]string", Required: false},
	{Name: "UserBackgroundSessionsEnabled", Flag: "user-background-sessions-enabled", Type: "*bool", Required: false},
}

var fields_create_integration = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "DataFilter", Flag: "data-filter", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationConfig", Flag: "integration-config", Type: "*types.IntegrationConfig", Required: false},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_create_integration_resource_property = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceProcessingProperties", Flag: "source-processing-properties", Type: "*types.SourceProcessingProperties", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetProcessingProperties", Flag: "target-processing-properties", Type: "*types.TargetProcessingProperties", Required: false},
}

var fields_create_integration_table_properties = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceTableConfig", Flag: "source-table-config", Type: "*types.SourceTableConfig", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TargetTableConfig", Flag: "target-table-config", Type: "*types.TargetTableConfig", Required: false},
}

var fields_create_job = []leanruntime.Field{
	{Name: "AllocatedCapacity", Flag: "allocated-capacity", Type: "int32", Required: false},
	{Name: "CodeGenConfigurationNodes", Flag: "code-gen-configuration-nodes", Type: "map[string]types.CodeGenConfigurationNode", Required: false},
	{Name: "Command", Flag: "command", Type: "*types.JobCommand", Required: true},
	{Name: "Connections", Flag: "connections", Type: "*types.ConnectionsList", Required: false},
	{Name: "DefaultArguments", Flag: "default-arguments", Type: "map[string]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionClass", Flag: "execution-class", Type: "types.ExecutionClass", Required: false},
	{Name: "ExecutionProperty", Flag: "execution-property", Type: "*types.ExecutionProperty", Required: false},
	{Name: "GlueVersion", Flag: "glue-version", Type: "*string", Required: false},
	{Name: "JobMode", Flag: "job-mode", Type: "types.JobMode", Required: false},
	{Name: "JobRunQueuingEnabled", Flag: "job-run-queuing-enabled", Type: "*bool", Required: false},
	{Name: "LogUri", Flag: "log-uri", Type: "*string", Required: false},
	{Name: "MaintenanceWindow", Flag: "maintenance-window", Type: "*string", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*float64", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NonOverridableArguments", Flag: "non-overridable-arguments", Type: "map[string]string", Required: false},
	{Name: "NotificationProperty", Flag: "notification-property", Type: "*types.NotificationProperty", Required: false},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "SourceControlDetails", Flag: "source-control-details", Type: "*types.SourceControlDetails", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "WorkerType", Flag: "worker-type", Type: "types.WorkerType", Required: false},
}

var fields_create_ml_transform = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlueVersion", Flag: "glue-version", Type: "*string", Required: false},
	{Name: "InputRecordTables", Flag: "input-record-tables", Type: "[]types.GlueTable", Required: true},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*float64", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "*types.TransformParameters", Required: true},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "TransformEncryption", Flag: "transform-encryption", Type: "*types.TransformEncryption", Required: false},
	{Name: "WorkerType", Flag: "worker-type", Type: "types.WorkerType", Required: false},
}

var fields_create_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionInput", Flag: "partition-input", Type: "*types.PartitionInput", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_create_partition_index = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionIndex", Flag: "partition-index", Type: "*types.PartitionIndex", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_create_registry = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_schema = []leanruntime.Field{
	{Name: "Compatibility", Flag: "compatibility", Type: "types.Compatibility", Required: false},
	{Name: "DataFormat", Flag: "data-format", Type: "types.DataFormat", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*types.RegistryId", Required: false},
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*string", Required: false},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_script = []leanruntime.Field{
	{Name: "DagEdges", Flag: "dag-edges", Type: "[]types.CodeGenEdge", Required: false},
	{Name: "DagNodes", Flag: "dag-nodes", Type: "[]types.CodeGenNode", Required: false},
	{Name: "Language", Flag: "language", Type: "types.Language", Required: false},
}

var fields_create_security_configuration = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_session = []leanruntime.Field{
	{Name: "Command", Flag: "command", Type: "*types.SessionCommand", Required: true},
	{Name: "Connections", Flag: "connections", Type: "*types.ConnectionsList", Required: false},
	{Name: "DefaultArguments", Flag: "default-arguments", Type: "map[string]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlueVersion", Flag: "glue-version", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IdleTimeout", Flag: "idle-timeout", Type: "*int32", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*float64", Required: false},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "WorkerType", Flag: "worker-type", Type: "types.WorkerType", Required: false},
}

var fields_create_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OpenTableFormatInput", Flag: "open-table-format-input", Type: "*types.OpenTableFormatInput", Required: false},
	{Name: "PartitionIndexes", Flag: "partition-indexes", Type: "[]types.PartitionIndex", Required: false},
	{Name: "TableInput", Flag: "table-input", Type: "*types.TableInput", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_create_table_optimizer = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TableOptimizerConfiguration", Flag: "table-optimizer-configuration", Type: "*types.TableOptimizerConfiguration", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableOptimizerType", Required: true},
}

var fields_create_trigger = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.Action", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBatchingCondition", Flag: "event-batching-condition", Type: "*types.EventBatchingCondition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Predicate", Flag: "predicate", Type: "*types.Predicate", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "StartOnCreation", Flag: "start-on-creation", Type: "bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.TriggerType", Required: true},
	{Name: "WorkflowName", Flag: "workflow-name", Type: "*string", Required: false},
}

var fields_create_usage_profile = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.ProfileConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_user_defined_function = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "FunctionInput", Flag: "function-input", Type: "*types.UserDefinedFunctionInput", Required: true},
}

var fields_create_workflow = []leanruntime.Field{
	{Name: "DefaultRunProperties", Flag: "default-run-properties", Type: "map[string]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaxConcurrentRuns", Flag: "max-concurrent-runs", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_blueprint = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_catalog = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
}

var fields_delete_classifier = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_column_statistics_for_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnName", Flag: "column-name", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionValues", Flag: "partition-values", Type: "[]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_column_statistics_for_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnName", Flag: "column-name", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_column_statistics_task_settings = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
}

var fields_delete_connection_type = []leanruntime.Field{
	{Name: "ConnectionType", Flag: "connection-type", Type: "*string", Required: true},
}

var fields_delete_crawler = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_custom_entity_type = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_data_quality_ruleset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_database = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_dev_endpoint = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_delete_glue_identity_center_configuration = []leanruntime.Field{}

var fields_delete_integration = []leanruntime.Field{
	{Name: "IntegrationIdentifier", Flag: "integration-identifier", Type: "*string", Required: true},
}

var fields_delete_integration_resource_property = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_integration_table_properties = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_job = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
}

var fields_delete_ml_transform = []leanruntime.Field{
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_delete_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionValues", Flag: "partition-values", Type: "[]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_partition_index = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_registry = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*types.RegistryId", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "PolicyHashCondition", Flag: "policy-hash-condition", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_delete_schema = []leanruntime.Field{
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
}

var fields_delete_schema_versions = []leanruntime.Field{
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
	{Name: "Versions", Flag: "versions", Type: "*string", Required: true},
}

var fields_delete_security_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
}

var fields_delete_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_delete_table_optimizer = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableOptimizerType", Required: true},
}

var fields_delete_table_version = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_delete_trigger = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_usage_profile = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_user_defined_function = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_connection_type = []leanruntime.Field{
	{Name: "ConnectionType", Flag: "connection-type", Type: "*string", Required: true},
}

var fields_describe_entity = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
	{Name: "DataStoreApiVersion", Flag: "data-store-api-version", Type: "*string", Required: false},
	{Name: "EntityName", Flag: "entity-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_inbound_integrations = []leanruntime.Field{
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_describe_integrations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.IntegrationFilter", Required: false},
	{Name: "IntegrationIdentifier", Flag: "integration-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_get_blueprint = []leanruntime.Field{
	{Name: "IncludeBlueprint", Flag: "include-blueprint", Type: "*bool", Required: false},
	{Name: "IncludeParameterSpec", Flag: "include-parameter-spec", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_blueprint_run = []leanruntime.Field{
	{Name: "BlueprintName", Flag: "blueprint-name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_blueprint_runs = []leanruntime.Field{
	{Name: "BlueprintName", Flag: "blueprint-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_catalog = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
}

var fields_get_catalog_import_status = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
}

var fields_get_catalogs = []leanruntime.Field{
	{Name: "IncludeRoot", Flag: "include-root", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentCatalogId", Flag: "parent-catalog-id", Type: "*string", Required: false},
	{Name: "Recursive", Flag: "recursive", Type: "bool", Required: false},
}

var fields_get_classifier = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_classifiers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_column_statistics_for_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnNames", Flag: "column-names", Type: "[]string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionValues", Flag: "partition-values", Type: "[]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_column_statistics_for_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnNames", Flag: "column-names", Type: "[]string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_column_statistics_task_run = []leanruntime.Field{
	{Name: "ColumnStatisticsTaskRunId", Flag: "column-statistics-task-run-id", Type: "*string", Required: true},
}

var fields_get_column_statistics_task_runs = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_column_statistics_task_settings = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_connection = []leanruntime.Field{
	{Name: "ApplyOverrideForComputeEnvironment", Flag: "apply-override-for-compute-environment", Type: "types.ComputeEnvironment", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "HidePassword", Flag: "hide-password", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_connections = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.GetConnectionsFilter", Required: false},
	{Name: "HidePassword", Flag: "hide-password", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_crawler = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_crawler_metrics = []leanruntime.Field{
	{Name: "CrawlerNameList", Flag: "crawler-name-list", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_crawlers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_custom_entity_type = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_data_catalog_encryption_settings = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
}

var fields_get_data_quality_model = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "StatisticId", Flag: "statistic-id", Type: "*string", Required: false},
}

var fields_get_data_quality_model_result = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "StatisticId", Flag: "statistic-id", Type: "*string", Required: true},
}

var fields_get_data_quality_result = []leanruntime.Field{
	{Name: "ResultId", Flag: "result-id", Type: "*string", Required: true},
}

var fields_get_data_quality_rule_recommendation_run = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_data_quality_ruleset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_data_quality_ruleset_evaluation_run = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_database = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_databases = []leanruntime.Field{
	{Name: "AttributesToGet", Flag: "attributes-to-get", Type: "[]types.DatabaseAttributes", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceShareType", Flag: "resource-share-type", Type: "types.ResourceShareType", Required: false},
}

var fields_get_dataflow_graph = []leanruntime.Field{
	{Name: "PythonScript", Flag: "python-script", Type: "*string", Required: false},
}

var fields_get_dev_endpoint = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_get_dev_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_entity_records = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: false},
	{Name: "ConnectionOptions", Flag: "connection-options", Type: "map[string]string", Required: false},
	{Name: "DataStoreApiVersion", Flag: "data-store-api-version", Type: "*string", Required: false},
	{Name: "EntityName", Flag: "entity-name", Type: "*string", Required: true},
	{Name: "FilterPredicate", Flag: "filter-predicate", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int64", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "*string", Required: false},
	{Name: "SelectedFields", Flag: "selected-fields", Type: "[]string", Required: false},
}

var fields_get_glue_identity_center_configuration = []leanruntime.Field{}

var fields_get_integration_resource_property = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_integration_table_properties = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
}

var fields_get_job_bookmark = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
}

var fields_get_job_run = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "PredecessorsIncluded", Flag: "predecessors-included", Type: "bool", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_job_runs = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_mapping = []leanruntime.Field{
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "Sinks", Flag: "sinks", Type: "[]types.CatalogEntry", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.CatalogEntry", Required: true},
}

var fields_get_materialized_view_refresh_task_run = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "MaterializedViewRefreshTaskRunId", Flag: "materialized-view-refresh-task-run-id", Type: "*string", Required: true},
}

var fields_get_ml_task_run = []leanruntime.Field{
	{Name: "TaskRunId", Flag: "task-run-id", Type: "*string", Required: true},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_get_ml_task_runs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TaskRunFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.TaskRunSortCriteria", Required: false},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_get_ml_transform = []leanruntime.Field{
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_get_ml_transforms = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TransformFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.TransformSortCriteria", Required: false},
}

var fields_get_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionValues", Flag: "partition-values", Type: "[]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_partition_indexes = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_partitions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "ExcludeColumnSchema", Flag: "exclude-column-schema", Type: "*bool", Required: false},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryAsOfTime", Flag: "query-as-of-time", Type: "*time.Time", Required: false},
	{Name: "Segment", Flag: "segment", Type: "*types.Segment", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_get_plan = []leanruntime.Field{
	{Name: "AdditionalPlanOptionsMap", Flag: "additional-plan-options-map", Type: "map[string]string", Required: false},
	{Name: "Language", Flag: "language", Type: "types.Language", Required: false},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "Mapping", Flag: "mapping", Type: "[]types.MappingEntry", Required: true},
	{Name: "Sinks", Flag: "sinks", Type: "[]types.CatalogEntry", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.CatalogEntry", Required: true},
}

var fields_get_registry = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*types.RegistryId", Required: true},
}

var fields_get_resource_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_get_schema = []leanruntime.Field{
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
}

var fields_get_schema_by_definition = []leanruntime.Field{
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*string", Required: true},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
}

var fields_get_schema_version = []leanruntime.Field{
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: false},
	{Name: "SchemaVersionId", Flag: "schema-version-id", Type: "*string", Required: false},
	{Name: "SchemaVersionNumber", Flag: "schema-version-number", Type: "*types.SchemaVersionNumber", Required: false},
}

var fields_get_schema_versions_diff = []leanruntime.Field{
	{Name: "FirstSchemaVersionNumber", Flag: "first-schema-version-number", Type: "*types.SchemaVersionNumber", Required: true},
	{Name: "SchemaDiffType", Flag: "schema-diff-type", Type: "types.SchemaDiffType", Required: true},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
	{Name: "SecondSchemaVersionNumber", Flag: "second-schema-version-number", Type: "*types.SchemaVersionNumber", Required: true},
}

var fields_get_security_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_security_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
}

var fields_get_statement = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "int32", Required: true},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_table = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "IncludeStatusDetails", Flag: "include-status-details", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueryAsOfTime", Flag: "query-as-of-time", Type: "*time.Time", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_get_table_optimizer = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableOptimizerType", Required: true},
}

var fields_get_table_version = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_table_versions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_tables = []leanruntime.Field{
	{Name: "AttributesToGet", Flag: "attributes-to-get", Type: "[]types.TableAttributes", Required: false},
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: false},
	{Name: "IncludeStatusDetails", Flag: "include-status-details", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryAsOfTime", Flag: "query-as-of-time", Type: "*time.Time", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_get_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_trigger = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_triggers = []leanruntime.Field{
	{Name: "DependentJobName", Flag: "dependent-job-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_unfiltered_partition_metadata = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionValues", Flag: "partition-values", Type: "[]string", Required: true},
	{Name: "QuerySessionContext", Flag: "query-session-context", Type: "*types.QuerySessionContext", Required: false},
	{Name: "Region", Flag: "region", Type: "*string", Required: false},
	{Name: "SupportedPermissionTypes", Flag: "supported-permission-types", Type: "[]types.PermissionType", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_unfiltered_partitions_metadata = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QuerySessionContext", Flag: "query-session-context", Type: "*types.QuerySessionContext", Required: false},
	{Name: "Region", Flag: "region", Type: "*string", Required: false},
	{Name: "Segment", Flag: "segment", Type: "*types.Segment", Required: false},
	{Name: "SupportedPermissionTypes", Flag: "supported-permission-types", Type: "[]types.PermissionType", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_unfiltered_table_metadata = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentResourceArn", Flag: "parent-resource-arn", Type: "*string", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.Permission", Required: false},
	{Name: "QuerySessionContext", Flag: "query-session-context", Type: "*types.QuerySessionContext", Required: false},
	{Name: "Region", Flag: "region", Type: "*string", Required: false},
	{Name: "RootResourceArn", Flag: "root-resource-arn", Type: "*string", Required: false},
	{Name: "SupportedDialect", Flag: "supported-dialect", Type: "*types.SupportedDialect", Required: false},
	{Name: "SupportedPermissionTypes", Flag: "supported-permission-types", Type: "[]types.PermissionType", Required: true},
}

var fields_get_usage_profile = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_user_defined_function = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_get_user_defined_functions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "FunctionType", Flag: "function-type", Type: "types.FunctionType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Pattern", Flag: "pattern", Type: "*string", Required: true},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "IncludeGraph", Flag: "include-graph", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_workflow_run = []leanruntime.Field{
	{Name: "IncludeGraph", Flag: "include-graph", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_workflow_run_properties = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_workflow_runs = []leanruntime.Field{
	{Name: "IncludeGraph", Flag: "include-graph", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_import_catalog_to_glue = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
}

var fields_list_blueprints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_column_statistics_task_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connection_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_crawlers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_crawls = []leanruntime.Field{
	{Name: "CrawlerName", Flag: "crawler-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.CrawlsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_entity_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_data_quality_results = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DataQualityResultFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_quality_rule_recommendation_runs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DataQualityRuleRecommendationRunFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_quality_ruleset_evaluation_runs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DataQualityRulesetEvaluationRunFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_quality_rulesets = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DataQualityRulesetFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_data_quality_statistic_annotations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: false},
	{Name: "StatisticId", Flag: "statistic-id", Type: "*string", Required: false},
	{Name: "TimestampFilter", Flag: "timestamp-filter", Type: "*types.TimestampFilter", Required: false},
}

var fields_list_data_quality_statistics = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: false},
	{Name: "StatisticId", Flag: "statistic-id", Type: "*string", Required: false},
	{Name: "TimestampFilter", Flag: "timestamp-filter", Type: "*types.TimestampFilter", Required: false},
}

var fields_list_dev_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_entities = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: false},
	{Name: "DataStoreApiVersion", Flag: "data-store-api-version", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentEntityName", Flag: "parent-entity-name", Type: "*string", Required: false},
}

var fields_list_integration_resource_properties = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.IntegrationResourcePropertyFilter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_materialized_view_refresh_task_runs = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: false},
}

var fields_list_ml_transforms = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TransformFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.TransformSortCriteria", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_registries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_schema_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
}

var fields_list_schemas = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*types.RegistryId", Required: false},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_statements = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_table_optimizer_runs = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableOptimizerType", Required: true},
}

var fields_list_triggers = []leanruntime.Field{
	{Name: "DependentJobName", Flag: "dependent-job-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_usage_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_modify_integration = []leanruntime.Field{
	{Name: "DataFilter", Flag: "data-filter", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationConfig", Flag: "integration-config", Type: "*types.IntegrationConfig", Required: false},
	{Name: "IntegrationIdentifier", Flag: "integration-identifier", Type: "*string", Required: true},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: false},
}

var fields_put_data_catalog_encryption_settings = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DataCatalogEncryptionSettings", Flag: "data-catalog-encryption-settings", Type: "*types.DataCatalogEncryptionSettings", Required: true},
}

var fields_put_data_quality_profile_annotation = []leanruntime.Field{
	{Name: "InclusionAnnotation", Flag: "inclusion-annotation", Type: "types.InclusionAnnotationValue", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "EnableHybrid", Flag: "enable-hybrid", Type: "types.EnableHybridValues", Required: false},
	{Name: "PolicyExistsCondition", Flag: "policy-exists-condition", Type: "types.ExistCondition", Required: false},
	{Name: "PolicyHashCondition", Flag: "policy-hash-condition", Type: "*string", Required: false},
	{Name: "PolicyInJson", Flag: "policy-in-json", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_put_schema_version_metadata = []leanruntime.Field{
	{Name: "MetadataKeyValue", Flag: "metadata-key-value", Type: "*types.MetadataKeyValuePair", Required: true},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: false},
	{Name: "SchemaVersionId", Flag: "schema-version-id", Type: "*string", Required: false},
	{Name: "SchemaVersionNumber", Flag: "schema-version-number", Type: "*types.SchemaVersionNumber", Required: false},
}

var fields_put_workflow_run_properties = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
	{Name: "RunProperties", Flag: "run-properties", Type: "map[string]string", Required: true},
}

var fields_query_schema_version_metadata = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetadataList", Flag: "metadata-list", Type: "[]types.MetadataKeyValuePair", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: false},
	{Name: "SchemaVersionId", Flag: "schema-version-id", Type: "*string", Required: false},
	{Name: "SchemaVersionNumber", Flag: "schema-version-number", Type: "*types.SchemaVersionNumber", Required: false},
}

var fields_register_connection_type = []leanruntime.Field{
	{Name: "ConnectionProperties", Flag: "connection-properties", Type: "*types.ConnectionPropertiesConfiguration", Required: true},
	{Name: "ConnectionType", Flag: "connection-type", Type: "*string", Required: true},
	{Name: "ConnectorAuthenticationConfiguration", Flag: "connector-authentication-configuration", Type: "*types.ConnectorAuthenticationConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: true},
	{Name: "RestConfiguration", Flag: "rest-configuration", Type: "*types.RestConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_register_schema_version = []leanruntime.Field{
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*string", Required: true},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
}

var fields_remove_schema_version_metadata = []leanruntime.Field{
	{Name: "MetadataKeyValue", Flag: "metadata-key-value", Type: "*types.MetadataKeyValuePair", Required: true},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: false},
	{Name: "SchemaVersionId", Flag: "schema-version-id", Type: "*string", Required: false},
	{Name: "SchemaVersionNumber", Flag: "schema-version-number", Type: "*types.SchemaVersionNumber", Required: false},
}

var fields_reset_job_bookmark = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
}

var fields_resume_workflow_run = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NodeIds", Flag: "node-ids", Type: "[]string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_run_statement = []leanruntime.Field{
	{Name: "Code", Flag: "code", Type: "*string", Required: true},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_search_tables = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.PropertyPredicate", Required: false},
	{Name: "IncludeStatusDetails", Flag: "include-status-details", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceShareType", Flag: "resource-share-type", Type: "types.ResourceShareType", Required: false},
	{Name: "SearchText", Flag: "search-text", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "[]types.SortCriterion", Required: false},
}

var fields_start_blueprint_run = []leanruntime.Field{
	{Name: "BlueprintName", Flag: "blueprint-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_start_column_statistics_task_run = []leanruntime.Field{
	{Name: "CatalogID", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnNameList", Flag: "column-name-list", Type: "[]string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "SampleSize", Flag: "sample-size", Type: "float64", Required: false},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_start_column_statistics_task_run_schedule = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_start_crawler = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_crawler_schedule = []leanruntime.Field{
	{Name: "CrawlerName", Flag: "crawler-name", Type: "*string", Required: true},
}

var fields_start_data_quality_rule_recommendation_run = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CreatedRulesetName", Flag: "created-ruleset-name", Type: "*string", Required: false},
	{Name: "DataQualitySecurityConfiguration", Flag: "data-quality-security-configuration", Type: "*string", Required: false},
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: true},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
}

var fields_start_data_quality_ruleset_evaluation_run = []leanruntime.Field{
	{Name: "AdditionalDataSources", Flag: "additional-data-sources", Type: "map[string]types.DataSource", Required: false},
	{Name: "AdditionalRunOptions", Flag: "additional-run-options", Type: "*types.DataQualityEvaluationRunAdditionalRunOptions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSource", Flag: "data-source", Type: "*types.DataSource", Required: true},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "RulesetNames", Flag: "ruleset-names", Type: "[]string", Required: true},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
}

var fields_start_export_labels_task_run = []leanruntime.Field{
	{Name: "OutputS3Path", Flag: "output-s3-path", Type: "*string", Required: true},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_start_import_labels_task_run = []leanruntime.Field{
	{Name: "InputS3Path", Flag: "input-s3-path", Type: "*string", Required: true},
	{Name: "ReplaceAllLabels", Flag: "replace-all-labels", Type: "bool", Required: false},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_start_job_run = []leanruntime.Field{
	{Name: "AllocatedCapacity", Flag: "allocated-capacity", Type: "int32", Required: false},
	{Name: "Arguments", Flag: "arguments", Type: "map[string]string", Required: false},
	{Name: "ExecutionClass", Flag: "execution-class", Type: "types.ExecutionClass", Required: false},
	{Name: "ExecutionRoleSessionPolicy", Flag: "execution-role-session-policy", Type: "*string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobRunId", Flag: "job-run-id", Type: "*string", Required: false},
	{Name: "JobRunQueuingEnabled", Flag: "job-run-queuing-enabled", Type: "*bool", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*float64", Required: false},
	{Name: "NotificationProperty", Flag: "notification-property", Type: "*types.NotificationProperty", Required: false},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "WorkerType", Flag: "worker-type", Type: "types.WorkerType", Required: false},
}

var fields_start_materialized_view_refresh_task_run = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "FullRefresh", Flag: "full-refresh", Type: "*bool", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_start_ml_evaluation_task_run = []leanruntime.Field{
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_start_ml_labeling_set_generation_task_run = []leanruntime.Field{
	{Name: "OutputS3Path", Flag: "output-s3-path", Type: "*string", Required: true},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
}

var fields_start_trigger = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_workflow_run = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunProperties", Flag: "run-properties", Type: "map[string]string", Required: false},
}

var fields_stop_column_statistics_task_run = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_stop_column_statistics_task_run_schedule = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_stop_crawler = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_crawler_schedule = []leanruntime.Field{
	{Name: "CrawlerName", Flag: "crawler-name", Type: "*string", Required: true},
}

var fields_stop_materialized_view_refresh_task_run = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_stop_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RequestOrigin", Flag: "request-origin", Type: "*string", Required: false},
}

var fields_stop_trigger = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_workflow_run = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagsToAdd", Flag: "tags-to-add", Type: "map[string]string", Required: true},
}

var fields_test_connection = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: false},
	{Name: "TestConnectionInput", Flag: "test-connection-input", Type: "*types.TestConnectionInput", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagsToRemove", Flag: "tags-to-remove", Type: "[]string", Required: true},
}

var fields_update_blueprint = []leanruntime.Field{
	{Name: "BlueprintLocation", Flag: "blueprint-location", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_catalog = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "CatalogInput", Flag: "catalog-input", Type: "*types.CatalogInput", Required: true},
}

var fields_update_classifier = []leanruntime.Field{
	{Name: "CsvClassifier", Flag: "csv-classifier", Type: "*types.UpdateCsvClassifierRequest", Required: false},
	{Name: "GrokClassifier", Flag: "grok-classifier", Type: "*types.UpdateGrokClassifierRequest", Required: false},
	{Name: "JsonClassifier", Flag: "json-classifier", Type: "*types.UpdateJsonClassifierRequest", Required: false},
	{Name: "XMLClassifier", Flag: "xml-classifier", Type: "*types.UpdateXMLClassifierRequest", Required: false},
}

var fields_update_column_statistics_for_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnStatisticsList", Flag: "column-statistics-list", Type: "[]types.ColumnStatistics", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionValues", Flag: "partition-values", Type: "[]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_column_statistics_for_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnStatisticsList", Flag: "column-statistics-list", Type: "[]types.ColumnStatistics", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_column_statistics_task_settings = []leanruntime.Field{
	{Name: "CatalogID", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ColumnNameList", Flag: "column-name-list", Type: "[]string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "SampleSize", Flag: "sample-size", Type: "float64", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_connection = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ConnectionInput", Flag: "connection-input", Type: "*types.ConnectionInput", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_crawler = []leanruntime.Field{
	{Name: "Classifiers", Flag: "classifiers", Type: "[]string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: false},
	{Name: "CrawlerSecurityConfiguration", Flag: "crawler-security-configuration", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LakeFormationConfiguration", Flag: "lake-formation-configuration", Type: "*types.LakeFormationConfiguration", Required: false},
	{Name: "LineageConfiguration", Flag: "lineage-configuration", Type: "*types.LineageConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecrawlPolicy", Flag: "recrawl-policy", Type: "*types.RecrawlPolicy", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "SchemaChangePolicy", Flag: "schema-change-policy", Type: "*types.SchemaChangePolicy", Required: false},
	{Name: "TablePrefix", Flag: "table-prefix", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "*types.CrawlerTargets", Required: false},
}

var fields_update_crawler_schedule = []leanruntime.Field{
	{Name: "CrawlerName", Flag: "crawler-name", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
}

var fields_update_data_quality_ruleset = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Ruleset", Flag: "ruleset", Type: "*string", Required: false},
}

var fields_update_database = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseInput", Flag: "database-input", Type: "*types.DatabaseInput", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_dev_endpoint = []leanruntime.Field{
	{Name: "AddArguments", Flag: "add-arguments", Type: "map[string]string", Required: false},
	{Name: "AddPublicKeys", Flag: "add-public-keys", Type: "[]string", Required: false},
	{Name: "CustomLibraries", Flag: "custom-libraries", Type: "*types.DevEndpointCustomLibraries", Required: false},
	{Name: "DeleteArguments", Flag: "delete-arguments", Type: "[]string", Required: false},
	{Name: "DeletePublicKeys", Flag: "delete-public-keys", Type: "[]string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "PublicKey", Flag: "public-key", Type: "*string", Required: false},
	{Name: "UpdateEtlLibraries", Flag: "update-etl-libraries", Type: "bool", Required: false},
}

var fields_update_glue_identity_center_configuration = []leanruntime.Field{
	{Name: "Scopes", Flag: "scopes", Type: "[]string", Required: false},
	{Name: "UserBackgroundSessionsEnabled", Flag: "user-background-sessions-enabled", Type: "*bool", Required: false},
}

var fields_update_integration_resource_property = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceProcessingProperties", Flag: "source-processing-properties", Type: "*types.SourceProcessingProperties", Required: false},
	{Name: "TargetProcessingProperties", Flag: "target-processing-properties", Type: "*types.TargetProcessingProperties", Required: false},
}

var fields_update_integration_table_properties = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceTableConfig", Flag: "source-table-config", Type: "*types.SourceTableConfig", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TargetTableConfig", Flag: "target-table-config", Type: "*types.TargetTableConfig", Required: false},
}

var fields_update_job = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobUpdate", Flag: "job-update", Type: "*types.JobUpdate", Required: true},
}

var fields_update_job_from_source_control = []leanruntime.Field{
	{Name: "AuthStrategy", Flag: "auth-strategy", Type: "types.SourceControlAuthStrategy", Required: false},
	{Name: "AuthToken", Flag: "auth-token", Type: "*string", Required: false},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: false},
	{Name: "CommitId", Flag: "commit-id", Type: "*string", Required: false},
	{Name: "Folder", Flag: "folder", Type: "*string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "Provider", Flag: "provider", Type: "types.SourceControlProvider", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: false},
	{Name: "RepositoryOwner", Flag: "repository-owner", Type: "*string", Required: false},
}

var fields_update_ml_transform = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlueVersion", Flag: "glue-version", Type: "*string", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*float64", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NumberOfWorkers", Flag: "number-of-workers", Type: "*int32", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "*types.TransformParameters", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*int32", Required: false},
	{Name: "TransformId", Flag: "transform-id", Type: "*string", Required: true},
	{Name: "WorkerType", Flag: "worker-type", Type: "types.WorkerType", Required: false},
}

var fields_update_partition = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartitionInput", Flag: "partition-input", Type: "*types.PartitionInput", Required: true},
	{Name: "PartitionValueList", Flag: "partition-value-list", Type: "[]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_update_registry = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*types.RegistryId", Required: true},
}

var fields_update_schema = []leanruntime.Field{
	{Name: "Compatibility", Flag: "compatibility", Type: "types.Compatibility", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SchemaId", Flag: "schema-id", Type: "*types.SchemaId", Required: true},
	{Name: "SchemaVersionNumber", Flag: "schema-version-number", Type: "*types.SchemaVersionNumber", Required: false},
}

var fields_update_source_control_from_job = []leanruntime.Field{
	{Name: "AuthStrategy", Flag: "auth-strategy", Type: "types.SourceControlAuthStrategy", Required: false},
	{Name: "AuthToken", Flag: "auth-token", Type: "*string", Required: false},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: false},
	{Name: "CommitId", Flag: "commit-id", Type: "*string", Required: false},
	{Name: "Folder", Flag: "folder", Type: "*string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "Provider", Flag: "provider", Type: "types.SourceControlProvider", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: false},
	{Name: "RepositoryOwner", Flag: "repository-owner", Type: "*string", Required: false},
}

var fields_update_table = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SkipArchive", Flag: "skip-archive", Type: "*bool", Required: false},
	{Name: "TableInput", Flag: "table-input", Type: "*types.TableInput", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
	{Name: "UpdateOpenTableFormatInput", Flag: "update-open-table-format-input", Type: "*types.UpdateOpenTableFormatInput", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
	{Name: "ViewUpdateAction", Flag: "view-update-action", Type: "types.ViewUpdateAction", Required: false},
}

var fields_update_table_optimizer = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TableOptimizerConfiguration", Flag: "table-optimizer-configuration", Type: "*types.TableOptimizerConfiguration", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableOptimizerType", Required: true},
}

var fields_update_trigger = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TriggerUpdate", Flag: "trigger-update", Type: "*types.TriggerUpdate", Required: true},
}

var fields_update_usage_profile = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.ProfileConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_user_defined_function = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "FunctionInput", Flag: "function-input", Type: "*types.UserDefinedFunctionInput", Required: true},
	{Name: "FunctionName", Flag: "function-name", Type: "*string", Required: true},
}

var fields_update_workflow = []leanruntime.Field{
	{Name: "DefaultRunProperties", Flag: "default-run-properties", Type: "map[string]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaxConcurrentRuns", Flag: "max-concurrent-runs", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-partition": {
			Name:   "batch-create-partition",
			Fields: fields_batch_create_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreatePartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreatePartition(ctx, input)
			},
		},
		"batch-delete-connection": {
			Name:   "batch-delete-connection",
			Fields: fields_batch_delete_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteConnection(ctx, input)
			},
		},
		"batch-delete-partition": {
			Name:   "batch-delete-partition",
			Fields: fields_batch_delete_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeletePartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeletePartition(ctx, input)
			},
		},
		"batch-delete-table": {
			Name:   "batch-delete-table",
			Fields: fields_batch_delete_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteTable(ctx, input)
			},
		},
		"batch-delete-table-version": {
			Name:   "batch-delete-table-version",
			Fields: fields_batch_delete_table_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteTableVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_table_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteTableVersion(ctx, input)
			},
		},
		"batch-get-blueprints": {
			Name:   "batch-get-blueprints",
			Fields: fields_batch_get_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetBlueprintsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_blueprints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetBlueprints(ctx, input)
			},
		},
		"batch-get-crawlers": {
			Name:   "batch-get-crawlers",
			Fields: fields_batch_get_crawlers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCrawlersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_crawlers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCrawlers(ctx, input)
			},
		},
		"batch-get-custom-entity-types": {
			Name:   "batch-get-custom-entity-types",
			Fields: fields_batch_get_custom_entity_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCustomEntityTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_custom_entity_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCustomEntityTypes(ctx, input)
			},
		},
		"batch-get-data-quality-result": {
			Name:   "batch-get-data-quality-result",
			Fields: fields_batch_get_data_quality_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDataQualityResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_data_quality_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDataQualityResult(ctx, input)
			},
		},
		"batch-get-dev-endpoints": {
			Name:   "batch-get-dev-endpoints",
			Fields: fields_batch_get_dev_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDevEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_dev_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDevEndpoints(ctx, input)
			},
		},
		"batch-get-jobs": {
			Name:   "batch-get-jobs",
			Fields: fields_batch_get_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetJobs(ctx, input)
			},
		},
		"batch-get-partition": {
			Name:   "batch-get-partition",
			Fields: fields_batch_get_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetPartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetPartition(ctx, input)
			},
		},
		"batch-get-table-optimizer": {
			Name:   "batch-get-table-optimizer",
			Fields: fields_batch_get_table_optimizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetTableOptimizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_table_optimizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetTableOptimizer(ctx, input)
			},
		},
		"batch-get-triggers": {
			Name:   "batch-get-triggers",
			Fields: fields_batch_get_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetTriggersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_triggers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetTriggers(ctx, input)
			},
		},
		"batch-get-workflows": {
			Name:   "batch-get-workflows",
			Fields: fields_batch_get_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetWorkflowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_workflows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetWorkflows(ctx, input)
			},
		},
		"batch-put-data-quality-statistic-annotation": {
			Name:   "batch-put-data-quality-statistic-annotation",
			Fields: fields_batch_put_data_quality_statistic_annotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutDataQualityStatisticAnnotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_data_quality_statistic_annotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutDataQualityStatisticAnnotation(ctx, input)
			},
		},
		"batch-stop-job-run": {
			Name:   "batch-stop-job-run",
			Fields: fields_batch_stop_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchStopJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_stop_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchStopJobRun(ctx, input)
			},
		},
		"batch-update-partition": {
			Name:   "batch-update-partition",
			Fields: fields_batch_update_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdatePartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdatePartition(ctx, input)
			},
		},
		"cancel-data-quality-rule-recommendation-run": {
			Name:   "cancel-data-quality-rule-recommendation-run",
			Fields: fields_cancel_data_quality_rule_recommendation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDataQualityRuleRecommendationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_data_quality_rule_recommendation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDataQualityRuleRecommendationRun(ctx, input)
			},
		},
		"cancel-data-quality-ruleset-evaluation-run": {
			Name:   "cancel-data-quality-ruleset-evaluation-run",
			Fields: fields_cancel_data_quality_ruleset_evaluation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDataQualityRulesetEvaluationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_data_quality_ruleset_evaluation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDataQualityRulesetEvaluationRun(ctx, input)
			},
		},
		"cancel-ml-task-run": {
			Name:   "cancel-ml-task-run",
			Fields: fields_cancel_ml_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMLTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_ml_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMLTaskRun(ctx, input)
			},
		},
		"cancel-statement": {
			Name:   "cancel-statement",
			Fields: fields_cancel_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelStatement(ctx, input)
			},
		},
		"check-schema-version-validity": {
			Name:   "check-schema-version-validity",
			Fields: fields_check_schema_version_validity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckSchemaVersionValidityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_schema_version_validity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckSchemaVersionValidity(ctx, input)
			},
		},
		"create-blueprint": {
			Name:   "create-blueprint",
			Fields: fields_create_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBlueprint(ctx, input)
			},
		},
		"create-catalog": {
			Name:   "create-catalog",
			Fields: fields_create_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCatalog(ctx, input)
			},
		},
		"create-classifier": {
			Name:   "create-classifier",
			Fields: fields_create_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClassifier(ctx, input)
			},
		},
		"create-column-statistics-task-settings": {
			Name:   "create-column-statistics-task-settings",
			Fields: fields_create_column_statistics_task_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateColumnStatisticsTaskSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_column_statistics_task_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateColumnStatisticsTaskSettings(ctx, input)
			},
		},
		"create-connection": {
			Name:   "create-connection",
			Fields: fields_create_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnection(ctx, input)
			},
		},
		"create-crawler": {
			Name:   "create-crawler",
			Fields: fields_create_crawler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCrawlerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_crawler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCrawler(ctx, input)
			},
		},
		"create-custom-entity-type": {
			Name:   "create-custom-entity-type",
			Fields: fields_create_custom_entity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomEntityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_entity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomEntityType(ctx, input)
			},
		},
		"create-data-quality-ruleset": {
			Name:   "create-data-quality-ruleset",
			Fields: fields_create_data_quality_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataQualityRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_quality_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataQualityRuleset(ctx, input)
			},
		},
		"create-database": {
			Name:   "create-database",
			Fields: fields_create_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDatabase(ctx, input)
			},
		},
		"create-dev-endpoint": {
			Name:   "create-dev-endpoint",
			Fields: fields_create_dev_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDevEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dev_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDevEndpoint(ctx, input)
			},
		},
		"create-glue-identity-center-configuration": {
			Name:   "create-glue-identity-center-configuration",
			Fields: fields_create_glue_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlueIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_glue_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlueIdentityCenterConfiguration(ctx, input)
			},
		},
		"create-integration": {
			Name:   "create-integration",
			Fields: fields_create_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegration(ctx, input)
			},
		},
		"create-integration-resource-property": {
			Name:   "create-integration-resource-property",
			Fields: fields_create_integration_resource_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationResourcePropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration_resource_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegrationResourceProperty(ctx, input)
			},
		},
		"create-integration-table-properties": {
			Name:   "create-integration-table-properties",
			Fields: fields_create_integration_table_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationTablePropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration_table_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegrationTableProperties(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-ml-transform": {
			Name:   "create-ml-transform",
			Fields: fields_create_ml_transform,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMLTransformInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ml_transform, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMLTransform(ctx, input)
			},
		},
		"create-partition": {
			Name:   "create-partition",
			Fields: fields_create_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartition(ctx, input)
			},
		},
		"create-partition-index": {
			Name:   "create-partition-index",
			Fields: fields_create_partition_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartitionIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partition_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartitionIndex(ctx, input)
			},
		},
		"create-registry": {
			Name:   "create-registry",
			Fields: fields_create_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegistry(ctx, input)
			},
		},
		"create-schema": {
			Name:   "create-schema",
			Fields: fields_create_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchema(ctx, input)
			},
		},
		"create-script": {
			Name:   "create-script",
			Fields: fields_create_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScript(ctx, input)
			},
		},
		"create-security-configuration": {
			Name:   "create-security-configuration",
			Fields: fields_create_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityConfiguration(ctx, input)
			},
		},
		"create-session": {
			Name:   "create-session",
			Fields: fields_create_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSession(ctx, input)
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
		"create-table-optimizer": {
			Name:   "create-table-optimizer",
			Fields: fields_create_table_optimizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTableOptimizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_table_optimizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTableOptimizer(ctx, input)
			},
		},
		"create-trigger": {
			Name:   "create-trigger",
			Fields: fields_create_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrigger(ctx, input)
			},
		},
		"create-usage-profile": {
			Name:   "create-usage-profile",
			Fields: fields_create_usage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUsageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_usage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUsageProfile(ctx, input)
			},
		},
		"create-user-defined-function": {
			Name:   "create-user-defined-function",
			Fields: fields_create_user_defined_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserDefinedFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_defined_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserDefinedFunction(ctx, input)
			},
		},
		"create-workflow": {
			Name:   "create-workflow",
			Fields: fields_create_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflow(ctx, input)
			},
		},
		"delete-blueprint": {
			Name:   "delete-blueprint",
			Fields: fields_delete_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBlueprint(ctx, input)
			},
		},
		"delete-catalog": {
			Name:   "delete-catalog",
			Fields: fields_delete_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCatalog(ctx, input)
			},
		},
		"delete-classifier": {
			Name:   "delete-classifier",
			Fields: fields_delete_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClassifier(ctx, input)
			},
		},
		"delete-column-statistics-for-partition": {
			Name:   "delete-column-statistics-for-partition",
			Fields: fields_delete_column_statistics_for_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteColumnStatisticsForPartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_column_statistics_for_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteColumnStatisticsForPartition(ctx, input)
			},
		},
		"delete-column-statistics-for-table": {
			Name:   "delete-column-statistics-for-table",
			Fields: fields_delete_column_statistics_for_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteColumnStatisticsForTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_column_statistics_for_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteColumnStatisticsForTable(ctx, input)
			},
		},
		"delete-column-statistics-task-settings": {
			Name:   "delete-column-statistics-task-settings",
			Fields: fields_delete_column_statistics_task_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteColumnStatisticsTaskSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_column_statistics_task_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteColumnStatisticsTaskSettings(ctx, input)
			},
		},
		"delete-connection": {
			Name:   "delete-connection",
			Fields: fields_delete_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnection(ctx, input)
			},
		},
		"delete-connection-type": {
			Name:   "delete-connection-type",
			Fields: fields_delete_connection_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectionType(ctx, input)
			},
		},
		"delete-crawler": {
			Name:   "delete-crawler",
			Fields: fields_delete_crawler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCrawlerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_crawler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCrawler(ctx, input)
			},
		},
		"delete-custom-entity-type": {
			Name:   "delete-custom-entity-type",
			Fields: fields_delete_custom_entity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomEntityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_entity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomEntityType(ctx, input)
			},
		},
		"delete-data-quality-ruleset": {
			Name:   "delete-data-quality-ruleset",
			Fields: fields_delete_data_quality_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataQualityRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_quality_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataQualityRuleset(ctx, input)
			},
		},
		"delete-database": {
			Name:   "delete-database",
			Fields: fields_delete_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDatabase(ctx, input)
			},
		},
		"delete-dev-endpoint": {
			Name:   "delete-dev-endpoint",
			Fields: fields_delete_dev_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDevEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dev_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDevEndpoint(ctx, input)
			},
		},
		"delete-glue-identity-center-configuration": {
			Name:   "delete-glue-identity-center-configuration",
			Fields: fields_delete_glue_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlueIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_glue_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlueIdentityCenterConfiguration(ctx, input)
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
		"delete-integration-resource-property": {
			Name:   "delete-integration-resource-property",
			Fields: fields_delete_integration_resource_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationResourcePropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration_resource_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegrationResourceProperty(ctx, input)
			},
		},
		"delete-integration-table-properties": {
			Name:   "delete-integration-table-properties",
			Fields: fields_delete_integration_table_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationTablePropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration_table_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegrationTableProperties(ctx, input)
			},
		},
		"delete-job": {
			Name:   "delete-job",
			Fields: fields_delete_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJob(ctx, input)
			},
		},
		"delete-ml-transform": {
			Name:   "delete-ml-transform",
			Fields: fields_delete_ml_transform,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMLTransformInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ml_transform, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMLTransform(ctx, input)
			},
		},
		"delete-partition": {
			Name:   "delete-partition",
			Fields: fields_delete_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePartition(ctx, input)
			},
		},
		"delete-partition-index": {
			Name:   "delete-partition-index",
			Fields: fields_delete_partition_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePartitionIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_partition_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePartitionIndex(ctx, input)
			},
		},
		"delete-registry": {
			Name:   "delete-registry",
			Fields: fields_delete_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistry(ctx, input)
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
		"delete-schema": {
			Name:   "delete-schema",
			Fields: fields_delete_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchema(ctx, input)
			},
		},
		"delete-schema-versions": {
			Name:   "delete-schema-versions",
			Fields: fields_delete_schema_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchemaVersions(ctx, input)
			},
		},
		"delete-security-configuration": {
			Name:   "delete-security-configuration",
			Fields: fields_delete_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityConfiguration(ctx, input)
			},
		},
		"delete-session": {
			Name:   "delete-session",
			Fields: fields_delete_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSession(ctx, input)
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
		"delete-table-optimizer": {
			Name:   "delete-table-optimizer",
			Fields: fields_delete_table_optimizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableOptimizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_optimizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableOptimizer(ctx, input)
			},
		},
		"delete-table-version": {
			Name:   "delete-table-version",
			Fields: fields_delete_table_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableVersion(ctx, input)
			},
		},
		"delete-trigger": {
			Name:   "delete-trigger",
			Fields: fields_delete_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrigger(ctx, input)
			},
		},
		"delete-usage-profile": {
			Name:   "delete-usage-profile",
			Fields: fields_delete_usage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUsageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_usage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUsageProfile(ctx, input)
			},
		},
		"delete-user-defined-function": {
			Name:   "delete-user-defined-function",
			Fields: fields_delete_user_defined_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserDefinedFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_defined_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserDefinedFunction(ctx, input)
			},
		},
		"delete-workflow": {
			Name:   "delete-workflow",
			Fields: fields_delete_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflow(ctx, input)
			},
		},
		"describe-connection-type": {
			Name:   "describe-connection-type",
			Fields: fields_describe_connection_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connection_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectionType(ctx, input)
			},
		},
		"describe-entity": {
			Name:   "describe-entity",
			Fields: fields_describe_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntityInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_entity, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEntity(ctx, input)
				}
				var results []*svc.DescribeEntityOutput
				p := svc.NewDescribeEntityPaginator(client, input)
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
		"describe-inbound-integrations": {
			Name:   "describe-inbound-integrations",
			Fields: fields_describe_inbound_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInboundIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_inbound_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInboundIntegrations(ctx, input)
			},
		},
		"describe-integrations": {
			Name:   "describe-integrations",
			Fields: fields_describe_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIntegrations(ctx, input)
			},
		},
		"get-blueprint": {
			Name:   "get-blueprint",
			Fields: fields_get_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlueprint(ctx, input)
			},
		},
		"get-blueprint-run": {
			Name:   "get-blueprint-run",
			Fields: fields_get_blueprint_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlueprintRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blueprint_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlueprintRun(ctx, input)
			},
		},
		"get-blueprint-runs": {
			Name:   "get-blueprint-runs",
			Fields: fields_get_blueprint_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlueprintRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_blueprint_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBlueprintRuns(ctx, input)
				}
				var results []*svc.GetBlueprintRunsOutput
				p := svc.NewGetBlueprintRunsPaginator(client, input)
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
		"get-catalog": {
			Name:   "get-catalog",
			Fields: fields_get_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCatalog(ctx, input)
			},
		},
		"get-catalog-import-status": {
			Name:   "get-catalog-import-status",
			Fields: fields_get_catalog_import_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCatalogImportStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_catalog_import_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCatalogImportStatus(ctx, input)
			},
		},
		"get-catalogs": {
			Name:   "get-catalogs",
			Fields: fields_get_catalogs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCatalogsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_catalogs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCatalogs(ctx, input)
			},
		},
		"get-classifier": {
			Name:   "get-classifier",
			Fields: fields_get_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClassifier(ctx, input)
			},
		},
		"get-classifiers": {
			Name:   "get-classifiers",
			Fields: fields_get_classifiers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClassifiersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_classifiers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetClassifiers(ctx, input)
				}
				var results []*svc.GetClassifiersOutput
				p := svc.NewGetClassifiersPaginator(client, input)
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
		"get-column-statistics-for-partition": {
			Name:   "get-column-statistics-for-partition",
			Fields: fields_get_column_statistics_for_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetColumnStatisticsForPartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_column_statistics_for_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetColumnStatisticsForPartition(ctx, input)
			},
		},
		"get-column-statistics-for-table": {
			Name:   "get-column-statistics-for-table",
			Fields: fields_get_column_statistics_for_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetColumnStatisticsForTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_column_statistics_for_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetColumnStatisticsForTable(ctx, input)
			},
		},
		"get-column-statistics-task-run": {
			Name:   "get-column-statistics-task-run",
			Fields: fields_get_column_statistics_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetColumnStatisticsTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_column_statistics_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetColumnStatisticsTaskRun(ctx, input)
			},
		},
		"get-column-statistics-task-runs": {
			Name:   "get-column-statistics-task-runs",
			Fields: fields_get_column_statistics_task_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetColumnStatisticsTaskRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_column_statistics_task_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetColumnStatisticsTaskRuns(ctx, input)
				}
				var results []*svc.GetColumnStatisticsTaskRunsOutput
				p := svc.NewGetColumnStatisticsTaskRunsPaginator(client, input)
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
		"get-column-statistics-task-settings": {
			Name:   "get-column-statistics-task-settings",
			Fields: fields_get_column_statistics_task_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetColumnStatisticsTaskSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_column_statistics_task_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetColumnStatisticsTaskSettings(ctx, input)
			},
		},
		"get-connection": {
			Name:   "get-connection",
			Fields: fields_get_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnection(ctx, input)
			},
		},
		"get-connections": {
			Name:   "get-connections",
			Fields: fields_get_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetConnections(ctx, input)
				}
				var results []*svc.GetConnectionsOutput
				p := svc.NewGetConnectionsPaginator(client, input)
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
		"get-crawler": {
			Name:   "get-crawler",
			Fields: fields_get_crawler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCrawlerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_crawler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCrawler(ctx, input)
			},
		},
		"get-crawler-metrics": {
			Name:   "get-crawler-metrics",
			Fields: fields_get_crawler_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCrawlerMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_crawler_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCrawlerMetrics(ctx, input)
				}
				var results []*svc.GetCrawlerMetricsOutput
				p := svc.NewGetCrawlerMetricsPaginator(client, input)
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
		"get-crawlers": {
			Name:   "get-crawlers",
			Fields: fields_get_crawlers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCrawlersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_crawlers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCrawlers(ctx, input)
				}
				var results []*svc.GetCrawlersOutput
				p := svc.NewGetCrawlersPaginator(client, input)
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
		"get-custom-entity-type": {
			Name:   "get-custom-entity-type",
			Fields: fields_get_custom_entity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomEntityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_entity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomEntityType(ctx, input)
			},
		},
		"get-data-catalog-encryption-settings": {
			Name:   "get-data-catalog-encryption-settings",
			Fields: fields_get_data_catalog_encryption_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataCatalogEncryptionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_catalog_encryption_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataCatalogEncryptionSettings(ctx, input)
			},
		},
		"get-data-quality-model": {
			Name:   "get-data-quality-model",
			Fields: fields_get_data_quality_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataQualityModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_quality_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataQualityModel(ctx, input)
			},
		},
		"get-data-quality-model-result": {
			Name:   "get-data-quality-model-result",
			Fields: fields_get_data_quality_model_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataQualityModelResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_quality_model_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataQualityModelResult(ctx, input)
			},
		},
		"get-data-quality-result": {
			Name:   "get-data-quality-result",
			Fields: fields_get_data_quality_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataQualityResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_quality_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataQualityResult(ctx, input)
			},
		},
		"get-data-quality-rule-recommendation-run": {
			Name:   "get-data-quality-rule-recommendation-run",
			Fields: fields_get_data_quality_rule_recommendation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataQualityRuleRecommendationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_quality_rule_recommendation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataQualityRuleRecommendationRun(ctx, input)
			},
		},
		"get-data-quality-ruleset": {
			Name:   "get-data-quality-ruleset",
			Fields: fields_get_data_quality_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataQualityRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_quality_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataQualityRuleset(ctx, input)
			},
		},
		"get-data-quality-ruleset-evaluation-run": {
			Name:   "get-data-quality-ruleset-evaluation-run",
			Fields: fields_get_data_quality_ruleset_evaluation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataQualityRulesetEvaluationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_quality_ruleset_evaluation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataQualityRulesetEvaluationRun(ctx, input)
			},
		},
		"get-database": {
			Name:   "get-database",
			Fields: fields_get_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDatabase(ctx, input)
			},
		},
		"get-databases": {
			Name:   "get-databases",
			Fields: fields_get_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDatabases(ctx, input)
				}
				var results []*svc.GetDatabasesOutput
				p := svc.NewGetDatabasesPaginator(client, input)
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
		"get-dataflow-graph": {
			Name:   "get-dataflow-graph",
			Fields: fields_get_dataflow_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataflowGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dataflow_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataflowGraph(ctx, input)
			},
		},
		"get-dev-endpoint": {
			Name:   "get-dev-endpoint",
			Fields: fields_get_dev_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dev_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevEndpoint(ctx, input)
			},
		},
		"get-dev-endpoints": {
			Name:   "get-dev-endpoints",
			Fields: fields_get_dev_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_dev_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDevEndpoints(ctx, input)
				}
				var results []*svc.GetDevEndpointsOutput
				p := svc.NewGetDevEndpointsPaginator(client, input)
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
		"get-entity-records": {
			Name:   "get-entity-records",
			Fields: fields_get_entity_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEntityRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_entity_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEntityRecords(ctx, input)
			},
		},
		"get-glue-identity-center-configuration": {
			Name:   "get-glue-identity-center-configuration",
			Fields: fields_get_glue_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlueIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_glue_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlueIdentityCenterConfiguration(ctx, input)
			},
		},
		"get-integration-resource-property": {
			Name:   "get-integration-resource-property",
			Fields: fields_get_integration_resource_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationResourcePropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration_resource_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegrationResourceProperty(ctx, input)
			},
		},
		"get-integration-table-properties": {
			Name:   "get-integration-table-properties",
			Fields: fields_get_integration_table_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationTablePropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration_table_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegrationTableProperties(ctx, input)
			},
		},
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-job-bookmark": {
			Name:   "get-job-bookmark",
			Fields: fields_get_job_bookmark,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobBookmarkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_bookmark, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobBookmark(ctx, input)
			},
		},
		"get-job-run": {
			Name:   "get-job-run",
			Fields: fields_get_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobRun(ctx, input)
			},
		},
		"get-job-runs": {
			Name:   "get-job-runs",
			Fields: fields_get_job_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_job_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetJobRuns(ctx, input)
				}
				var results []*svc.GetJobRunsOutput
				p := svc.NewGetJobRunsPaginator(client, input)
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
		"get-jobs": {
			Name:   "get-jobs",
			Fields: fields_get_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetJobs(ctx, input)
				}
				var results []*svc.GetJobsOutput
				p := svc.NewGetJobsPaginator(client, input)
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
		"get-mapping": {
			Name:   "get-mapping",
			Fields: fields_get_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMapping(ctx, input)
			},
		},
		"get-materialized-view-refresh-task-run": {
			Name:   "get-materialized-view-refresh-task-run",
			Fields: fields_get_materialized_view_refresh_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMaterializedViewRefreshTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_materialized_view_refresh_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMaterializedViewRefreshTaskRun(ctx, input)
			},
		},
		"get-ml-task-run": {
			Name:   "get-ml-task-run",
			Fields: fields_get_ml_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLTaskRun(ctx, input)
			},
		},
		"get-ml-task-runs": {
			Name:   "get-ml-task-runs",
			Fields: fields_get_ml_task_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLTaskRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ml_task_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMLTaskRuns(ctx, input)
				}
				var results []*svc.GetMLTaskRunsOutput
				p := svc.NewGetMLTaskRunsPaginator(client, input)
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
		"get-ml-transform": {
			Name:   "get-ml-transform",
			Fields: fields_get_ml_transform,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLTransformInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ml_transform, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMLTransform(ctx, input)
			},
		},
		"get-ml-transforms": {
			Name:   "get-ml-transforms",
			Fields: fields_get_ml_transforms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMLTransformsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ml_transforms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMLTransforms(ctx, input)
				}
				var results []*svc.GetMLTransformsOutput
				p := svc.NewGetMLTransformsPaginator(client, input)
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
		"get-partition": {
			Name:   "get-partition",
			Fields: fields_get_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPartition(ctx, input)
			},
		},
		"get-partition-indexes": {
			Name:   "get-partition-indexes",
			Fields: fields_get_partition_indexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPartitionIndexesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_partition_indexes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPartitionIndexes(ctx, input)
				}
				var results []*svc.GetPartitionIndexesOutput
				p := svc.NewGetPartitionIndexesPaginator(client, input)
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
		"get-partitions": {
			Name:   "get-partitions",
			Fields: fields_get_partitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPartitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_partitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPartitions(ctx, input)
				}
				var results []*svc.GetPartitionsOutput
				p := svc.NewGetPartitionsPaginator(client, input)
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
		"get-plan": {
			Name:   "get-plan",
			Fields: fields_get_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlan(ctx, input)
			},
		},
		"get-registry": {
			Name:   "get-registry",
			Fields: fields_get_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegistry(ctx, input)
			},
		},
		"get-resource-policies": {
			Name:   "get-resource-policies",
			Fields: fields_get_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourcePolicies(ctx, input)
				}
				var results []*svc.GetResourcePoliciesOutput
				p := svc.NewGetResourcePoliciesPaginator(client, input)
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
		"get-schema": {
			Name:   "get-schema",
			Fields: fields_get_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchema(ctx, input)
			},
		},
		"get-schema-by-definition": {
			Name:   "get-schema-by-definition",
			Fields: fields_get_schema_by_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaByDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_by_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaByDefinition(ctx, input)
			},
		},
		"get-schema-version": {
			Name:   "get-schema-version",
			Fields: fields_get_schema_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaVersion(ctx, input)
			},
		},
		"get-schema-versions-diff": {
			Name:   "get-schema-versions-diff",
			Fields: fields_get_schema_versions_diff,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaVersionsDiffInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_versions_diff, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaVersionsDiff(ctx, input)
			},
		},
		"get-security-configuration": {
			Name:   "get-security-configuration",
			Fields: fields_get_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSecurityConfiguration(ctx, input)
			},
		},
		"get-security-configurations": {
			Name:   "get-security-configurations",
			Fields: fields_get_security_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_security_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSecurityConfigurations(ctx, input)
				}
				var results []*svc.GetSecurityConfigurationsOutput
				p := svc.NewGetSecurityConfigurationsPaginator(client, input)
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
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"get-statement": {
			Name:   "get-statement",
			Fields: fields_get_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStatement(ctx, input)
			},
		},
		"get-table": {
			Name:   "get-table",
			Fields: fields_get_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTable(ctx, input)
			},
		},
		"get-table-optimizer": {
			Name:   "get-table-optimizer",
			Fields: fields_get_table_optimizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableOptimizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_optimizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableOptimizer(ctx, input)
			},
		},
		"get-table-version": {
			Name:   "get-table-version",
			Fields: fields_get_table_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableVersion(ctx, input)
			},
		},
		"get-table-versions": {
			Name:   "get-table-versions",
			Fields: fields_get_table_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_table_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTableVersions(ctx, input)
				}
				var results []*svc.GetTableVersionsOutput
				p := svc.NewGetTableVersionsPaginator(client, input)
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
		"get-tables": {
			Name:   "get-tables",
			Fields: fields_get_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTables(ctx, input)
				}
				var results []*svc.GetTablesOutput
				p := svc.NewGetTablesPaginator(client, input)
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
		"get-tags": {
			Name:   "get-tags",
			Fields: fields_get_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTags(ctx, input)
			},
		},
		"get-trigger": {
			Name:   "get-trigger",
			Fields: fields_get_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrigger(ctx, input)
			},
		},
		"get-triggers": {
			Name:   "get-triggers",
			Fields: fields_get_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTriggersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_triggers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTriggers(ctx, input)
				}
				var results []*svc.GetTriggersOutput
				p := svc.NewGetTriggersPaginator(client, input)
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
		"get-unfiltered-partition-metadata": {
			Name:   "get-unfiltered-partition-metadata",
			Fields: fields_get_unfiltered_partition_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUnfilteredPartitionMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_unfiltered_partition_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUnfilteredPartitionMetadata(ctx, input)
			},
		},
		"get-unfiltered-partitions-metadata": {
			Name:   "get-unfiltered-partitions-metadata",
			Fields: fields_get_unfiltered_partitions_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUnfilteredPartitionsMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_unfiltered_partitions_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUnfilteredPartitionsMetadata(ctx, input)
				}
				var results []*svc.GetUnfilteredPartitionsMetadataOutput
				p := svc.NewGetUnfilteredPartitionsMetadataPaginator(client, input)
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
		"get-unfiltered-table-metadata": {
			Name:   "get-unfiltered-table-metadata",
			Fields: fields_get_unfiltered_table_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUnfilteredTableMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_unfiltered_table_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUnfilteredTableMetadata(ctx, input)
			},
		},
		"get-usage-profile": {
			Name:   "get-usage-profile",
			Fields: fields_get_usage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_usage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsageProfile(ctx, input)
			},
		},
		"get-user-defined-function": {
			Name:   "get-user-defined-function",
			Fields: fields_get_user_defined_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserDefinedFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_defined_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserDefinedFunction(ctx, input)
			},
		},
		"get-user-defined-functions": {
			Name:   "get-user-defined-functions",
			Fields: fields_get_user_defined_functions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserDefinedFunctionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_user_defined_functions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUserDefinedFunctions(ctx, input)
				}
				var results []*svc.GetUserDefinedFunctionsOutput
				p := svc.NewGetUserDefinedFunctionsPaginator(client, input)
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
		"get-workflow": {
			Name:   "get-workflow",
			Fields: fields_get_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflow(ctx, input)
			},
		},
		"get-workflow-run": {
			Name:   "get-workflow-run",
			Fields: fields_get_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowRun(ctx, input)
			},
		},
		"get-workflow-run-properties": {
			Name:   "get-workflow-run-properties",
			Fields: fields_get_workflow_run_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowRunPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_run_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowRunProperties(ctx, input)
			},
		},
		"get-workflow-runs": {
			Name:   "get-workflow-runs",
			Fields: fields_get_workflow_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_workflow_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetWorkflowRuns(ctx, input)
				}
				var results []*svc.GetWorkflowRunsOutput
				p := svc.NewGetWorkflowRunsPaginator(client, input)
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
		"import-catalog-to-glue": {
			Name:   "import-catalog-to-glue",
			Fields: fields_import_catalog_to_glue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCatalogToGlueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_catalog_to_glue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCatalogToGlue(ctx, input)
			},
		},
		"list-blueprints": {
			Name:   "list-blueprints",
			Fields: fields_list_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBlueprintsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_blueprints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBlueprints(ctx, input)
				}
				var results []*svc.ListBlueprintsOutput
				p := svc.NewListBlueprintsPaginator(client, input)
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
		"list-column-statistics-task-runs": {
			Name:   "list-column-statistics-task-runs",
			Fields: fields_list_column_statistics_task_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListColumnStatisticsTaskRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_column_statistics_task_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListColumnStatisticsTaskRuns(ctx, input)
				}
				var results []*svc.ListColumnStatisticsTaskRunsOutput
				p := svc.NewListColumnStatisticsTaskRunsPaginator(client, input)
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
		"list-connection-types": {
			Name:   "list-connection-types",
			Fields: fields_list_connection_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connection_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectionTypes(ctx, input)
				}
				var results []*svc.ListConnectionTypesOutput
				p := svc.NewListConnectionTypesPaginator(client, input)
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
		"list-crawlers": {
			Name:   "list-crawlers",
			Fields: fields_list_crawlers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrawlersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_crawlers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCrawlers(ctx, input)
				}
				var results []*svc.ListCrawlersOutput
				p := svc.NewListCrawlersPaginator(client, input)
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
		"list-crawls": {
			Name:   "list-crawls",
			Fields: fields_list_crawls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrawlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_crawls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCrawls(ctx, input)
			},
		},
		"list-custom-entity-types": {
			Name:   "list-custom-entity-types",
			Fields: fields_list_custom_entity_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomEntityTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_entity_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomEntityTypes(ctx, input)
				}
				var results []*svc.ListCustomEntityTypesOutput
				p := svc.NewListCustomEntityTypesPaginator(client, input)
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
		"list-data-quality-results": {
			Name:   "list-data-quality-results",
			Fields: fields_list_data_quality_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_quality_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataQualityResults(ctx, input)
				}
				var results []*svc.ListDataQualityResultsOutput
				p := svc.NewListDataQualityResultsPaginator(client, input)
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
		"list-data-quality-rule-recommendation-runs": {
			Name:   "list-data-quality-rule-recommendation-runs",
			Fields: fields_list_data_quality_rule_recommendation_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityRuleRecommendationRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_quality_rule_recommendation_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataQualityRuleRecommendationRuns(ctx, input)
				}
				var results []*svc.ListDataQualityRuleRecommendationRunsOutput
				p := svc.NewListDataQualityRuleRecommendationRunsPaginator(client, input)
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
		"list-data-quality-ruleset-evaluation-runs": {
			Name:   "list-data-quality-ruleset-evaluation-runs",
			Fields: fields_list_data_quality_ruleset_evaluation_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityRulesetEvaluationRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_quality_ruleset_evaluation_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataQualityRulesetEvaluationRuns(ctx, input)
				}
				var results []*svc.ListDataQualityRulesetEvaluationRunsOutput
				p := svc.NewListDataQualityRulesetEvaluationRunsPaginator(client, input)
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
		"list-data-quality-rulesets": {
			Name:   "list-data-quality-rulesets",
			Fields: fields_list_data_quality_rulesets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityRulesetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_quality_rulesets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataQualityRulesets(ctx, input)
				}
				var results []*svc.ListDataQualityRulesetsOutput
				p := svc.NewListDataQualityRulesetsPaginator(client, input)
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
		"list-data-quality-statistic-annotations": {
			Name:   "list-data-quality-statistic-annotations",
			Fields: fields_list_data_quality_statistic_annotations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityStatisticAnnotationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_data_quality_statistic_annotations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDataQualityStatisticAnnotations(ctx, input)
			},
		},
		"list-data-quality-statistics": {
			Name:   "list-data-quality-statistics",
			Fields: fields_list_data_quality_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataQualityStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_data_quality_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDataQualityStatistics(ctx, input)
			},
		},
		"list-dev-endpoints": {
			Name:   "list-dev-endpoints",
			Fields: fields_list_dev_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dev_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevEndpoints(ctx, input)
				}
				var results []*svc.ListDevEndpointsOutput
				p := svc.NewListDevEndpointsPaginator(client, input)
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
		"list-entities": {
			Name:   "list-entities",
			Fields: fields_list_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntities(ctx, input)
				}
				var results []*svc.ListEntitiesOutput
				p := svc.NewListEntitiesPaginator(client, input)
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
		"list-integration-resource-properties": {
			Name:   "list-integration-resource-properties",
			Fields: fields_list_integration_resource_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntegrationResourcePropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_integration_resource_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIntegrationResourceProperties(ctx, input)
			},
		},
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-materialized-view-refresh-task-runs": {
			Name:   "list-materialized-view-refresh-task-runs",
			Fields: fields_list_materialized_view_refresh_task_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMaterializedViewRefreshTaskRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_materialized_view_refresh_task_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMaterializedViewRefreshTaskRuns(ctx, input)
				}
				var results []*svc.ListMaterializedViewRefreshTaskRunsOutput
				p := svc.NewListMaterializedViewRefreshTaskRunsPaginator(client, input)
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
		"list-ml-transforms": {
			Name:   "list-ml-transforms",
			Fields: fields_list_ml_transforms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMLTransformsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ml_transforms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMLTransforms(ctx, input)
				}
				var results []*svc.ListMLTransformsOutput
				p := svc.NewListMLTransformsPaginator(client, input)
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
		"list-registries": {
			Name:   "list-registries",
			Fields: fields_list_registries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegistriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_registries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegistries(ctx, input)
				}
				var results []*svc.ListRegistriesOutput
				p := svc.NewListRegistriesPaginator(client, input)
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
		"list-schema-versions": {
			Name:   "list-schema-versions",
			Fields: fields_list_schema_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemaVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schema_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemaVersions(ctx, input)
				}
				var results []*svc.ListSchemaVersionsOutput
				p := svc.NewListSchemaVersionsPaginator(client, input)
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
		"list-schemas": {
			Name:   "list-schemas",
			Fields: fields_list_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemas(ctx, input)
				}
				var results []*svc.ListSchemasOutput
				p := svc.NewListSchemasPaginator(client, input)
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
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"list-statements": {
			Name:   "list-statements",
			Fields: fields_list_statements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStatementsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_statements, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStatements(ctx, input)
			},
		},
		"list-table-optimizer-runs": {
			Name:   "list-table-optimizer-runs",
			Fields: fields_list_table_optimizer_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTableOptimizerRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_table_optimizer_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTableOptimizerRuns(ctx, input)
				}
				var results []*svc.ListTableOptimizerRunsOutput
				p := svc.NewListTableOptimizerRunsPaginator(client, input)
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
		"list-triggers": {
			Name:   "list-triggers",
			Fields: fields_list_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTriggersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_triggers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTriggers(ctx, input)
				}
				var results []*svc.ListTriggersOutput
				p := svc.NewListTriggersPaginator(client, input)
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
		"list-usage-profiles": {
			Name:   "list-usage-profiles",
			Fields: fields_list_usage_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsageProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_usage_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsageProfiles(ctx, input)
				}
				var results []*svc.ListUsageProfilesOutput
				p := svc.NewListUsageProfilesPaginator(client, input)
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
		"list-workflows": {
			Name:   "list-workflows",
			Fields: fields_list_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflows(ctx, input)
				}
				var results []*svc.ListWorkflowsOutput
				p := svc.NewListWorkflowsPaginator(client, input)
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
		"modify-integration": {
			Name:   "modify-integration",
			Fields: fields_modify_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIntegration(ctx, input)
			},
		},
		"put-data-catalog-encryption-settings": {
			Name:   "put-data-catalog-encryption-settings",
			Fields: fields_put_data_catalog_encryption_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataCatalogEncryptionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_catalog_encryption_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataCatalogEncryptionSettings(ctx, input)
			},
		},
		"put-data-quality-profile-annotation": {
			Name:   "put-data-quality-profile-annotation",
			Fields: fields_put_data_quality_profile_annotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataQualityProfileAnnotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_quality_profile_annotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataQualityProfileAnnotation(ctx, input)
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
		"put-schema-version-metadata": {
			Name:   "put-schema-version-metadata",
			Fields: fields_put_schema_version_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSchemaVersionMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_schema_version_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSchemaVersionMetadata(ctx, input)
			},
		},
		"put-workflow-run-properties": {
			Name:   "put-workflow-run-properties",
			Fields: fields_put_workflow_run_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutWorkflowRunPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_workflow_run_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutWorkflowRunProperties(ctx, input)
			},
		},
		"query-schema-version-metadata": {
			Name:   "query-schema-version-metadata",
			Fields: fields_query_schema_version_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QuerySchemaVersionMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_query_schema_version_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.QuerySchemaVersionMetadata(ctx, input)
			},
		},
		"register-connection-type": {
			Name:   "register-connection-type",
			Fields: fields_register_connection_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterConnectionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_connection_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterConnectionType(ctx, input)
			},
		},
		"register-schema-version": {
			Name:   "register-schema-version",
			Fields: fields_register_schema_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterSchemaVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_schema_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterSchemaVersion(ctx, input)
			},
		},
		"remove-schema-version-metadata": {
			Name:   "remove-schema-version-metadata",
			Fields: fields_remove_schema_version_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveSchemaVersionMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_schema_version_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveSchemaVersionMetadata(ctx, input)
			},
		},
		"reset-job-bookmark": {
			Name:   "reset-job-bookmark",
			Fields: fields_reset_job_bookmark,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetJobBookmarkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_job_bookmark, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetJobBookmark(ctx, input)
			},
		},
		"resume-workflow-run": {
			Name:   "resume-workflow-run",
			Fields: fields_resume_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeWorkflowRun(ctx, input)
			},
		},
		"run-statement": {
			Name:   "run-statement",
			Fields: fields_run_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RunStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_run_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RunStatement(ctx, input)
			},
		},
		"search-tables": {
			Name:   "search-tables",
			Fields: fields_search_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTables(ctx, input)
				}
				var results []*svc.SearchTablesOutput
				p := svc.NewSearchTablesPaginator(client, input)
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
		"start-blueprint-run": {
			Name:   "start-blueprint-run",
			Fields: fields_start_blueprint_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBlueprintRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_blueprint_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBlueprintRun(ctx, input)
			},
		},
		"start-column-statistics-task-run": {
			Name:   "start-column-statistics-task-run",
			Fields: fields_start_column_statistics_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartColumnStatisticsTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_column_statistics_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartColumnStatisticsTaskRun(ctx, input)
			},
		},
		"start-column-statistics-task-run-schedule": {
			Name:   "start-column-statistics-task-run-schedule",
			Fields: fields_start_column_statistics_task_run_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartColumnStatisticsTaskRunScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_column_statistics_task_run_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartColumnStatisticsTaskRunSchedule(ctx, input)
			},
		},
		"start-crawler": {
			Name:   "start-crawler",
			Fields: fields_start_crawler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCrawlerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_crawler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCrawler(ctx, input)
			},
		},
		"start-crawler-schedule": {
			Name:   "start-crawler-schedule",
			Fields: fields_start_crawler_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCrawlerScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_crawler_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCrawlerSchedule(ctx, input)
			},
		},
		"start-data-quality-rule-recommendation-run": {
			Name:   "start-data-quality-rule-recommendation-run",
			Fields: fields_start_data_quality_rule_recommendation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataQualityRuleRecommendationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_quality_rule_recommendation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataQualityRuleRecommendationRun(ctx, input)
			},
		},
		"start-data-quality-ruleset-evaluation-run": {
			Name:   "start-data-quality-ruleset-evaluation-run",
			Fields: fields_start_data_quality_ruleset_evaluation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataQualityRulesetEvaluationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_quality_ruleset_evaluation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataQualityRulesetEvaluationRun(ctx, input)
			},
		},
		"start-export-labels-task-run": {
			Name:   "start-export-labels-task-run",
			Fields: fields_start_export_labels_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExportLabelsTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_export_labels_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExportLabelsTaskRun(ctx, input)
			},
		},
		"start-import-labels-task-run": {
			Name:   "start-import-labels-task-run",
			Fields: fields_start_import_labels_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportLabelsTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import_labels_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImportLabelsTaskRun(ctx, input)
			},
		},
		"start-job-run": {
			Name:   "start-job-run",
			Fields: fields_start_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartJobRun(ctx, input)
			},
		},
		"start-materialized-view-refresh-task-run": {
			Name:   "start-materialized-view-refresh-task-run",
			Fields: fields_start_materialized_view_refresh_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMaterializedViewRefreshTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_materialized_view_refresh_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMaterializedViewRefreshTaskRun(ctx, input)
			},
		},
		"start-ml-evaluation-task-run": {
			Name:   "start-ml-evaluation-task-run",
			Fields: fields_start_ml_evaluation_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMLEvaluationTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ml_evaluation_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMLEvaluationTaskRun(ctx, input)
			},
		},
		"start-ml-labeling-set-generation-task-run": {
			Name:   "start-ml-labeling-set-generation-task-run",
			Fields: fields_start_ml_labeling_set_generation_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMLLabelingSetGenerationTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ml_labeling_set_generation_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMLLabelingSetGenerationTaskRun(ctx, input)
			},
		},
		"start-trigger": {
			Name:   "start-trigger",
			Fields: fields_start_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTrigger(ctx, input)
			},
		},
		"start-workflow-run": {
			Name:   "start-workflow-run",
			Fields: fields_start_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkflowRun(ctx, input)
			},
		},
		"stop-column-statistics-task-run": {
			Name:   "stop-column-statistics-task-run",
			Fields: fields_stop_column_statistics_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopColumnStatisticsTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_column_statistics_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopColumnStatisticsTaskRun(ctx, input)
			},
		},
		"stop-column-statistics-task-run-schedule": {
			Name:   "stop-column-statistics-task-run-schedule",
			Fields: fields_stop_column_statistics_task_run_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopColumnStatisticsTaskRunScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_column_statistics_task_run_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopColumnStatisticsTaskRunSchedule(ctx, input)
			},
		},
		"stop-crawler": {
			Name:   "stop-crawler",
			Fields: fields_stop_crawler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCrawlerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_crawler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCrawler(ctx, input)
			},
		},
		"stop-crawler-schedule": {
			Name:   "stop-crawler-schedule",
			Fields: fields_stop_crawler_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCrawlerScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_crawler_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCrawlerSchedule(ctx, input)
			},
		},
		"stop-materialized-view-refresh-task-run": {
			Name:   "stop-materialized-view-refresh-task-run",
			Fields: fields_stop_materialized_view_refresh_task_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMaterializedViewRefreshTaskRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_materialized_view_refresh_task_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMaterializedViewRefreshTaskRun(ctx, input)
			},
		},
		"stop-session": {
			Name:   "stop-session",
			Fields: fields_stop_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSession(ctx, input)
			},
		},
		"stop-trigger": {
			Name:   "stop-trigger",
			Fields: fields_stop_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTrigger(ctx, input)
			},
		},
		"stop-workflow-run": {
			Name:   "stop-workflow-run",
			Fields: fields_stop_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopWorkflowRun(ctx, input)
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
		"test-connection": {
			Name:   "test-connection",
			Fields: fields_test_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestConnection(ctx, input)
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
		"update-blueprint": {
			Name:   "update-blueprint",
			Fields: fields_update_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBlueprint(ctx, input)
			},
		},
		"update-catalog": {
			Name:   "update-catalog",
			Fields: fields_update_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCatalog(ctx, input)
			},
		},
		"update-classifier": {
			Name:   "update-classifier",
			Fields: fields_update_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClassifier(ctx, input)
			},
		},
		"update-column-statistics-for-partition": {
			Name:   "update-column-statistics-for-partition",
			Fields: fields_update_column_statistics_for_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateColumnStatisticsForPartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_column_statistics_for_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateColumnStatisticsForPartition(ctx, input)
			},
		},
		"update-column-statistics-for-table": {
			Name:   "update-column-statistics-for-table",
			Fields: fields_update_column_statistics_for_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateColumnStatisticsForTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_column_statistics_for_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateColumnStatisticsForTable(ctx, input)
			},
		},
		"update-column-statistics-task-settings": {
			Name:   "update-column-statistics-task-settings",
			Fields: fields_update_column_statistics_task_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateColumnStatisticsTaskSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_column_statistics_task_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateColumnStatisticsTaskSettings(ctx, input)
			},
		},
		"update-connection": {
			Name:   "update-connection",
			Fields: fields_update_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnection(ctx, input)
			},
		},
		"update-crawler": {
			Name:   "update-crawler",
			Fields: fields_update_crawler,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCrawlerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_crawler, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCrawler(ctx, input)
			},
		},
		"update-crawler-schedule": {
			Name:   "update-crawler-schedule",
			Fields: fields_update_crawler_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCrawlerScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_crawler_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCrawlerSchedule(ctx, input)
			},
		},
		"update-data-quality-ruleset": {
			Name:   "update-data-quality-ruleset",
			Fields: fields_update_data_quality_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataQualityRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_quality_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataQualityRuleset(ctx, input)
			},
		},
		"update-database": {
			Name:   "update-database",
			Fields: fields_update_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDatabase(ctx, input)
			},
		},
		"update-dev-endpoint": {
			Name:   "update-dev-endpoint",
			Fields: fields_update_dev_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDevEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dev_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDevEndpoint(ctx, input)
			},
		},
		"update-glue-identity-center-configuration": {
			Name:   "update-glue-identity-center-configuration",
			Fields: fields_update_glue_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlueIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_glue_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlueIdentityCenterConfiguration(ctx, input)
			},
		},
		"update-integration-resource-property": {
			Name:   "update-integration-resource-property",
			Fields: fields_update_integration_resource_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIntegrationResourcePropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_integration_resource_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIntegrationResourceProperty(ctx, input)
			},
		},
		"update-integration-table-properties": {
			Name:   "update-integration-table-properties",
			Fields: fields_update_integration_table_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIntegrationTablePropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_integration_table_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIntegrationTableProperties(ctx, input)
			},
		},
		"update-job": {
			Name:   "update-job",
			Fields: fields_update_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJob(ctx, input)
			},
		},
		"update-job-from-source-control": {
			Name:   "update-job-from-source-control",
			Fields: fields_update_job_from_source_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobFromSourceControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_from_source_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobFromSourceControl(ctx, input)
			},
		},
		"update-ml-transform": {
			Name:   "update-ml-transform",
			Fields: fields_update_ml_transform,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMLTransformInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ml_transform, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMLTransform(ctx, input)
			},
		},
		"update-partition": {
			Name:   "update-partition",
			Fields: fields_update_partition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePartitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_partition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePartition(ctx, input)
			},
		},
		"update-registry": {
			Name:   "update-registry",
			Fields: fields_update_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRegistry(ctx, input)
			},
		},
		"update-schema": {
			Name:   "update-schema",
			Fields: fields_update_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchema(ctx, input)
			},
		},
		"update-source-control-from-job": {
			Name:   "update-source-control-from-job",
			Fields: fields_update_source_control_from_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSourceControlFromJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_source_control_from_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSourceControlFromJob(ctx, input)
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
		"update-table-optimizer": {
			Name:   "update-table-optimizer",
			Fields: fields_update_table_optimizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableOptimizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table_optimizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTableOptimizer(ctx, input)
			},
		},
		"update-trigger": {
			Name:   "update-trigger",
			Fields: fields_update_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrigger(ctx, input)
			},
		},
		"update-usage-profile": {
			Name:   "update-usage-profile",
			Fields: fields_update_usage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUsageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_usage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUsageProfile(ctx, input)
			},
		},
		"update-user-defined-function": {
			Name:   "update-user-defined-function",
			Fields: fields_update_user_defined_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserDefinedFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_defined_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserDefinedFunction(ctx, input)
			},
		},
		"update-workflow": {
			Name:   "update-workflow",
			Fields: fields_update_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflow(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("glue", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
