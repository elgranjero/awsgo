package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_apply_pending_maintenance_action = []leanruntime.Field{
	{Name: "ApplyAction", Flag: "apply-action", Type: "*string", Required: true},
	{Name: "OptInType", Flag: "opt-in-type", Type: "*string", Required: true},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_batch_start_recommendations = []leanruntime.Field{
	{Name: "Data", Flag: "data", Type: "[]types.StartRecommendationsRequestEntry", Required: false},
}

var fields_cancel_metadata_model_conversion = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "RequestIdentifier", Flag: "request-identifier", Type: "*string", Required: true},
}

var fields_cancel_metadata_model_creation = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "RequestIdentifier", Flag: "request-identifier", Type: "*string", Required: true},
}

var fields_cancel_replication_task_assessment_run = []leanruntime.Field{
	{Name: "ReplicationTaskAssessmentRunArn", Flag: "replication-task-assessment-run-arn", Type: "*string", Required: true},
}

var fields_create_data_migration = []leanruntime.Field{
	{Name: "DataMigrationName", Flag: "data-migration-name", Type: "*string", Required: false},
	{Name: "DataMigrationType", Flag: "data-migration-type", Type: "types.MigrationTypeValue", Required: true},
	{Name: "EnableCloudwatchLogs", Flag: "enable-cloudwatch-logs", Type: "*bool", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "NumberOfJobs", Flag: "number-of-jobs", Type: "*int32", Required: false},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: false},
	{Name: "ServiceAccessRoleArn", Flag: "service-access-role-arn", Type: "*string", Required: true},
	{Name: "SourceDataSettings", Flag: "source-data-settings", Type: "[]types.SourceDataSetting", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDataSettings", Flag: "target-data-settings", Type: "[]types.TargetDataSetting", Required: false},
}

var fields_create_data_provider = []leanruntime.Field{
	{Name: "DataProviderName", Flag: "data-provider-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "types.DataProviderSettings", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Virtual", Flag: "virtual", Type: "*bool", Required: false},
}

var fields_create_endpoint = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DmsTransferSettings", Flag: "dms-transfer-settings", Type: "*types.DmsTransferSettings", Required: false},
	{Name: "DocDbSettings", Flag: "doc-db-settings", Type: "*types.DocDbSettings", Required: false},
	{Name: "DynamoDbSettings", Flag: "dynamo-db-settings", Type: "*types.DynamoDbSettings", Required: false},
	{Name: "ElasticsearchSettings", Flag: "elasticsearch-settings", Type: "*types.ElasticsearchSettings", Required: false},
	{Name: "EndpointIdentifier", Flag: "endpoint-identifier", Type: "*string", Required: true},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.ReplicationEndpointTypeValue", Required: true},
	{Name: "EngineName", Flag: "engine-name", Type: "*string", Required: true},
	{Name: "ExternalTableDefinition", Flag: "external-table-definition", Type: "*string", Required: false},
	{Name: "ExtraConnectionAttributes", Flag: "extra-connection-attributes", Type: "*string", Required: false},
	{Name: "GcpMySQLSettings", Flag: "gcp-my-sql-settings", Type: "*types.GcpMySQLSettings", Required: false},
	{Name: "IBMDb2Settings", Flag: "ibmdb2-settings", Type: "*types.IBMDb2Settings", Required: false},
	{Name: "KafkaSettings", Flag: "kafka-settings", Type: "*types.KafkaSettings", Required: false},
	{Name: "KinesisSettings", Flag: "kinesis-settings", Type: "*types.KinesisSettings", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MicrosoftSQLServerSettings", Flag: "microsoft-sql-server-settings", Type: "*types.MicrosoftSQLServerSettings", Required: false},
	{Name: "MongoDbSettings", Flag: "mongo-db-settings", Type: "*types.MongoDbSettings", Required: false},
	{Name: "MySQLSettings", Flag: "my-sql-settings", Type: "*types.MySQLSettings", Required: false},
	{Name: "NeptuneSettings", Flag: "neptune-settings", Type: "*types.NeptuneSettings", Required: false},
	{Name: "OracleSettings", Flag: "oracle-settings", Type: "*types.OracleSettings", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PostgreSQLSettings", Flag: "postgre-sql-settings", Type: "*types.PostgreSQLSettings", Required: false},
	{Name: "RedisSettings", Flag: "redis-settings", Type: "*types.RedisSettings", Required: false},
	{Name: "RedshiftSettings", Flag: "redshift-settings", Type: "*types.RedshiftSettings", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "S3Settings", Flag: "s3-settings", Type: "*types.S3Settings", Required: false},
	{Name: "ServerName", Flag: "server-name", Type: "*string", Required: false},
	{Name: "ServiceAccessRoleArn", Flag: "service-access-role-arn", Type: "*string", Required: false},
	{Name: "SslMode", Flag: "ssl-mode", Type: "types.DmsSslModeValue", Required: false},
	{Name: "SybaseSettings", Flag: "sybase-settings", Type: "*types.SybaseSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimestreamSettings", Flag: "timestream-settings", Type: "*types.TimestreamSettings", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_create_event_subscription = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventCategories", Flag: "event-categories", Type: "[]string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: true},
	{Name: "SourceIds", Flag: "source-ids", Type: "[]string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_fleet_advisor_collector = []leanruntime.Field{
	{Name: "CollectorName", Flag: "collector-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "ServiceAccessRoleArn", Flag: "service-access-role-arn", Type: "*string", Required: true},
}

var fields_create_instance_profile = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "SubnetGroupIdentifier", Flag: "subnet-group-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroups", Flag: "vpc-security-groups", Type: "[]string", Required: false},
}

var fields_create_migration_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceProfileIdentifier", Flag: "instance-profile-identifier", Type: "*string", Required: true},
	{Name: "MigrationProjectName", Flag: "migration-project-name", Type: "*string", Required: false},
	{Name: "SchemaConversionApplicationAttributes", Flag: "schema-conversion-application-attributes", Type: "*types.SCApplicationAttributes", Required: false},
	{Name: "SourceDataProviderDescriptors", Flag: "source-data-provider-descriptors", Type: "[]types.DataProviderDescriptorDefinition", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDataProviderDescriptors", Flag: "target-data-provider-descriptors", Type: "[]types.DataProviderDescriptorDefinition", Required: true},
	{Name: "TransformationRules", Flag: "transformation-rules", Type: "*string", Required: false},
}

var fields_create_replication_config = []leanruntime.Field{
	{Name: "ComputeConfig", Flag: "compute-config", Type: "*types.ComputeConfig", Required: true},
	{Name: "ReplicationConfigIdentifier", Flag: "replication-config-identifier", Type: "*string", Required: true},
	{Name: "ReplicationSettings", Flag: "replication-settings", Type: "*string", Required: false},
	{Name: "ReplicationType", Flag: "replication-type", Type: "types.MigrationTypeValue", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "SourceEndpointArn", Flag: "source-endpoint-arn", Type: "*string", Required: true},
	{Name: "SupplementalSettings", Flag: "supplemental-settings", Type: "*string", Required: false},
	{Name: "TableMappings", Flag: "table-mappings", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetEndpointArn", Flag: "target-endpoint-arn", Type: "*string", Required: true},
}

var fields_create_replication_instance = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "DnsNameServers", Flag: "dns-name-servers", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "KerberosAuthenticationSettings", Flag: "kerberos-authentication-settings", Type: "*types.KerberosAuthenticationSettings", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "ReplicationInstanceClass", Flag: "replication-instance-class", Type: "*string", Required: true},
	{Name: "ReplicationInstanceIdentifier", Flag: "replication-instance-identifier", Type: "*string", Required: true},
	{Name: "ReplicationSubnetGroupIdentifier", Flag: "replication-subnet-group-identifier", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_create_replication_subnet_group = []leanruntime.Field{
	{Name: "ReplicationSubnetGroupDescription", Flag: "replication-subnet-group-description", Type: "*string", Required: true},
	{Name: "ReplicationSubnetGroupIdentifier", Flag: "replication-subnet-group-identifier", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_replication_task = []leanruntime.Field{
	{Name: "CdcStartPosition", Flag: "cdc-start-position", Type: "*string", Required: false},
	{Name: "CdcStartTime", Flag: "cdc-start-time", Type: "*time.Time", Required: false},
	{Name: "CdcStopPosition", Flag: "cdc-stop-position", Type: "*string", Required: false},
	{Name: "MigrationType", Flag: "migration-type", Type: "types.MigrationTypeValue", Required: true},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
	{Name: "ReplicationTaskIdentifier", Flag: "replication-task-identifier", Type: "*string", Required: true},
	{Name: "ReplicationTaskSettings", Flag: "replication-task-settings", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "SourceEndpointArn", Flag: "source-endpoint-arn", Type: "*string", Required: true},
	{Name: "TableMappings", Flag: "table-mappings", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetEndpointArn", Flag: "target-endpoint-arn", Type: "*string", Required: true},
	{Name: "TaskData", Flag: "task-data", Type: "*string", Required: false},
}

var fields_delete_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_delete_data_migration = []leanruntime.Field{
	{Name: "DataMigrationIdentifier", Flag: "data-migration-identifier", Type: "*string", Required: true},
}

var fields_delete_data_provider = []leanruntime.Field{
	{Name: "DataProviderIdentifier", Flag: "data-provider-identifier", Type: "*string", Required: true},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_delete_event_subscription = []leanruntime.Field{
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_delete_fleet_advisor_collector = []leanruntime.Field{
	{Name: "CollectorReferencedId", Flag: "collector-referenced-id", Type: "*string", Required: true},
}

var fields_delete_fleet_advisor_databases = []leanruntime.Field{
	{Name: "DatabaseIds", Flag: "database-ids", Type: "[]string", Required: true},
}

var fields_delete_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileIdentifier", Flag: "instance-profile-identifier", Type: "*string", Required: true},
}

var fields_delete_migration_project = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_delete_replication_config = []leanruntime.Field{
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: true},
}

var fields_delete_replication_instance = []leanruntime.Field{
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_delete_replication_subnet_group = []leanruntime.Field{
	{Name: "ReplicationSubnetGroupIdentifier", Flag: "replication-subnet-group-identifier", Type: "*string", Required: true},
}

var fields_delete_replication_task = []leanruntime.Field{
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
}

var fields_delete_replication_task_assessment_run = []leanruntime.Field{
	{Name: "ReplicationTaskAssessmentRunArn", Flag: "replication-task-assessment-run-arn", Type: "*string", Required: true},
}

var fields_describe_account_attributes = []leanruntime.Field{}

var fields_describe_applicable_individual_assessments = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationType", Flag: "migration-type", Type: "types.MigrationTypeValue", Required: false},
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: false},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: false},
	{Name: "SourceEngineName", Flag: "source-engine-name", Type: "*string", Required: false},
	{Name: "TargetEngineName", Flag: "target-engine-name", Type: "*string", Required: false},
}

var fields_describe_certificates = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_connections = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_conversion_configuration = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_data_migrations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "WithoutSettings", Flag: "without-settings", Type: "*bool", Required: false},
	{Name: "WithoutStatistics", Flag: "without-statistics", Type: "*bool", Required: false},
}

var fields_describe_data_providers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_endpoint_settings = []leanruntime.Field{
	{Name: "EngineName", Flag: "engine-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_endpoint_types = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_endpoints = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_engine_versions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_event_categories = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
}

var fields_describe_event_subscriptions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: false},
}

var fields_describe_events = []leanruntime.Field{
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "EventCategories", Flag: "event-categories", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SourceIdentifier", Flag: "source-identifier", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.SourceType", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_extension_pack_associations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_fleet_advisor_collectors = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_advisor_databases = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_advisor_lsa_analysis = []leanruntime.Field{
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_advisor_schema_object_summary = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_advisor_schemas = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_profiles = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_metadata_model = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "Origin", Flag: "origin", Type: "types.OriginTypeValue", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_describe_metadata_model_assessments = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_metadata_model_children = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "Origin", Flag: "origin", Type: "types.OriginTypeValue", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_describe_metadata_model_conversions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_metadata_model_creations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_metadata_model_exports_as_script = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_metadata_model_exports_to_target = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_metadata_model_imports = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_describe_migration_projects = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_orderable_replication_instances = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_pending_maintenance_actions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: false},
}

var fields_describe_recommendation_limitations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_recommendations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_refresh_schemas_status = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_describe_replication_configs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_replication_instance_task_logs = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_describe_replication_instances = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_replication_subnet_groups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_replication_table_statistics = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: true},
}

var fields_describe_replication_task_assessment_results = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: false},
}

var fields_describe_replication_task_assessment_runs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_replication_task_individual_assessments = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_replication_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "WithoutSettings", Flag: "without-settings", Type: "*bool", Required: false},
}

var fields_describe_replications = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_schemas = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_table_statistics = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
}

var fields_export_metadata_model_assessment = []leanruntime.Field{
	{Name: "AssessmentReportTypes", Flag: "assessment-report-types", Type: "[]types.AssessmentReportType", Required: false},
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_get_target_selection_rules = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_import_certificate = []leanruntime.Field{
	{Name: "CertificateIdentifier", Flag: "certificate-identifier", Type: "*string", Required: true},
	{Name: "CertificatePem", Flag: "certificate-pem", Type: "*string", Required: false},
	{Name: "CertificateWallet", Flag: "certificate-wallet", Type: "[]byte", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceArnList", Flag: "resource-arn-list", Type: "[]string", Required: false},
}

var fields_modify_conversion_configuration = []leanruntime.Field{
	{Name: "ConversionConfiguration", Flag: "conversion-configuration", Type: "*string", Required: true},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_modify_data_migration = []leanruntime.Field{
	{Name: "DataMigrationIdentifier", Flag: "data-migration-identifier", Type: "*string", Required: true},
	{Name: "DataMigrationName", Flag: "data-migration-name", Type: "*string", Required: false},
	{Name: "DataMigrationType", Flag: "data-migration-type", Type: "types.MigrationTypeValue", Required: false},
	{Name: "EnableCloudwatchLogs", Flag: "enable-cloudwatch-logs", Type: "*bool", Required: false},
	{Name: "NumberOfJobs", Flag: "number-of-jobs", Type: "*int32", Required: false},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: false},
	{Name: "ServiceAccessRoleArn", Flag: "service-access-role-arn", Type: "*string", Required: false},
	{Name: "SourceDataSettings", Flag: "source-data-settings", Type: "[]types.SourceDataSetting", Required: false},
	{Name: "TargetDataSettings", Flag: "target-data-settings", Type: "[]types.TargetDataSetting", Required: false},
}

var fields_modify_data_provider = []leanruntime.Field{
	{Name: "DataProviderIdentifier", Flag: "data-provider-identifier", Type: "*string", Required: true},
	{Name: "DataProviderName", Flag: "data-provider-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "ExactSettings", Flag: "exact-settings", Type: "*bool", Required: false},
	{Name: "Settings", Flag: "settings", Type: "types.DataProviderSettings", Required: false},
	{Name: "Virtual", Flag: "virtual", Type: "*bool", Required: false},
}

var fields_modify_endpoint = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DmsTransferSettings", Flag: "dms-transfer-settings", Type: "*types.DmsTransferSettings", Required: false},
	{Name: "DocDbSettings", Flag: "doc-db-settings", Type: "*types.DocDbSettings", Required: false},
	{Name: "DynamoDbSettings", Flag: "dynamo-db-settings", Type: "*types.DynamoDbSettings", Required: false},
	{Name: "ElasticsearchSettings", Flag: "elasticsearch-settings", Type: "*types.ElasticsearchSettings", Required: false},
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "EndpointIdentifier", Flag: "endpoint-identifier", Type: "*string", Required: false},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.ReplicationEndpointTypeValue", Required: false},
	{Name: "EngineName", Flag: "engine-name", Type: "*string", Required: false},
	{Name: "ExactSettings", Flag: "exact-settings", Type: "*bool", Required: false},
	{Name: "ExternalTableDefinition", Flag: "external-table-definition", Type: "*string", Required: false},
	{Name: "ExtraConnectionAttributes", Flag: "extra-connection-attributes", Type: "*string", Required: false},
	{Name: "GcpMySQLSettings", Flag: "gcp-my-sql-settings", Type: "*types.GcpMySQLSettings", Required: false},
	{Name: "IBMDb2Settings", Flag: "ibmdb2-settings", Type: "*types.IBMDb2Settings", Required: false},
	{Name: "KafkaSettings", Flag: "kafka-settings", Type: "*types.KafkaSettings", Required: false},
	{Name: "KinesisSettings", Flag: "kinesis-settings", Type: "*types.KinesisSettings", Required: false},
	{Name: "MicrosoftSQLServerSettings", Flag: "microsoft-sql-server-settings", Type: "*types.MicrosoftSQLServerSettings", Required: false},
	{Name: "MongoDbSettings", Flag: "mongo-db-settings", Type: "*types.MongoDbSettings", Required: false},
	{Name: "MySQLSettings", Flag: "my-sql-settings", Type: "*types.MySQLSettings", Required: false},
	{Name: "NeptuneSettings", Flag: "neptune-settings", Type: "*types.NeptuneSettings", Required: false},
	{Name: "OracleSettings", Flag: "oracle-settings", Type: "*types.OracleSettings", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PostgreSQLSettings", Flag: "postgre-sql-settings", Type: "*types.PostgreSQLSettings", Required: false},
	{Name: "RedisSettings", Flag: "redis-settings", Type: "*types.RedisSettings", Required: false},
	{Name: "RedshiftSettings", Flag: "redshift-settings", Type: "*types.RedshiftSettings", Required: false},
	{Name: "S3Settings", Flag: "s3-settings", Type: "*types.S3Settings", Required: false},
	{Name: "ServerName", Flag: "server-name", Type: "*string", Required: false},
	{Name: "ServiceAccessRoleArn", Flag: "service-access-role-arn", Type: "*string", Required: false},
	{Name: "SslMode", Flag: "ssl-mode", Type: "types.DmsSslModeValue", Required: false},
	{Name: "SybaseSettings", Flag: "sybase-settings", Type: "*types.SybaseSettings", Required: false},
	{Name: "TimestreamSettings", Flag: "timestream-settings", Type: "*types.TimestreamSettings", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_modify_event_subscription = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventCategories", Flag: "event-categories", Type: "[]string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_modify_instance_profile = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceProfileIdentifier", Flag: "instance-profile-identifier", Type: "*string", Required: true},
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "SubnetGroupIdentifier", Flag: "subnet-group-identifier", Type: "*string", Required: false},
	{Name: "VpcSecurityGroups", Flag: "vpc-security-groups", Type: "[]string", Required: false},
}

var fields_modify_migration_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceProfileIdentifier", Flag: "instance-profile-identifier", Type: "*string", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "MigrationProjectName", Flag: "migration-project-name", Type: "*string", Required: false},
	{Name: "SchemaConversionApplicationAttributes", Flag: "schema-conversion-application-attributes", Type: "*types.SCApplicationAttributes", Required: false},
	{Name: "SourceDataProviderDescriptors", Flag: "source-data-provider-descriptors", Type: "[]types.DataProviderDescriptorDefinition", Required: false},
	{Name: "TargetDataProviderDescriptors", Flag: "target-data-provider-descriptors", Type: "[]types.DataProviderDescriptorDefinition", Required: false},
	{Name: "TransformationRules", Flag: "transformation-rules", Type: "*string", Required: false},
}

var fields_modify_replication_config = []leanruntime.Field{
	{Name: "ComputeConfig", Flag: "compute-config", Type: "*types.ComputeConfig", Required: false},
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: true},
	{Name: "ReplicationConfigIdentifier", Flag: "replication-config-identifier", Type: "*string", Required: false},
	{Name: "ReplicationSettings", Flag: "replication-settings", Type: "*string", Required: false},
	{Name: "ReplicationType", Flag: "replication-type", Type: "types.MigrationTypeValue", Required: false},
	{Name: "SourceEndpointArn", Flag: "source-endpoint-arn", Type: "*string", Required: false},
	{Name: "SupplementalSettings", Flag: "supplemental-settings", Type: "*string", Required: false},
	{Name: "TableMappings", Flag: "table-mappings", Type: "*string", Required: false},
	{Name: "TargetEndpointArn", Flag: "target-endpoint-arn", Type: "*string", Required: false},
}

var fields_modify_replication_instance = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AllowMajorVersionUpgrade", Flag: "allow-major-version-upgrade", Type: "bool", Required: false},
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "bool", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "KerberosAuthenticationSettings", Flag: "kerberos-authentication-settings", Type: "*types.KerberosAuthenticationSettings", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
	{Name: "ReplicationInstanceClass", Flag: "replication-instance-class", Type: "*string", Required: false},
	{Name: "ReplicationInstanceIdentifier", Flag: "replication-instance-identifier", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_replication_subnet_group = []leanruntime.Field{
	{Name: "ReplicationSubnetGroupDescription", Flag: "replication-subnet-group-description", Type: "*string", Required: false},
	{Name: "ReplicationSubnetGroupIdentifier", Flag: "replication-subnet-group-identifier", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
}

var fields_modify_replication_task = []leanruntime.Field{
	{Name: "CdcStartPosition", Flag: "cdc-start-position", Type: "*string", Required: false},
	{Name: "CdcStartTime", Flag: "cdc-start-time", Type: "*time.Time", Required: false},
	{Name: "CdcStopPosition", Flag: "cdc-stop-position", Type: "*string", Required: false},
	{Name: "MigrationType", Flag: "migration-type", Type: "types.MigrationTypeValue", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
	{Name: "ReplicationTaskIdentifier", Flag: "replication-task-identifier", Type: "*string", Required: false},
	{Name: "ReplicationTaskSettings", Flag: "replication-task-settings", Type: "*string", Required: false},
	{Name: "TableMappings", Flag: "table-mappings", Type: "*string", Required: false},
	{Name: "TaskData", Flag: "task-data", Type: "*string", Required: false},
}

var fields_move_replication_task = []leanruntime.Field{
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
	{Name: "TargetReplicationInstanceArn", Flag: "target-replication-instance-arn", Type: "*string", Required: true},
}

var fields_reboot_replication_instance = []leanruntime.Field{
	{Name: "ForceFailover", Flag: "force-failover", Type: "*bool", Required: false},
	{Name: "ForcePlannedFailover", Flag: "force-planned-failover", Type: "*bool", Required: false},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_refresh_schemas = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_reload_replication_tables = []leanruntime.Field{
	{Name: "ReloadOption", Flag: "reload-option", Type: "types.ReloadOptionValue", Required: false},
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: true},
	{Name: "TablesToReload", Flag: "tables-to-reload", Type: "[]types.TableToReload", Required: true},
}

var fields_reload_tables = []leanruntime.Field{
	{Name: "ReloadOption", Flag: "reload-option", Type: "types.ReloadOptionValue", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
	{Name: "TablesToReload", Flag: "tables-to-reload", Type: "[]types.TableToReload", Required: true},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_run_fleet_advisor_lsa_analysis = []leanruntime.Field{}

var fields_start_data_migration = []leanruntime.Field{
	{Name: "DataMigrationIdentifier", Flag: "data-migration-identifier", Type: "*string", Required: true},
	{Name: "StartType", Flag: "start-type", Type: "types.StartReplicationMigrationTypeValue", Required: true},
}

var fields_start_extension_pack_association = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
}

var fields_start_metadata_model_assessment = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_start_metadata_model_conversion = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_start_metadata_model_creation = []leanruntime.Field{
	{Name: "MetadataModelName", Flag: "metadata-model-name", Type: "*string", Required: true},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "Properties", Flag: "properties", Type: "types.MetadataModelProperties", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_start_metadata_model_export_as_script = []leanruntime.Field{
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: false},
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "Origin", Flag: "origin", Type: "types.OriginTypeValue", Required: true},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_start_metadata_model_export_to_target = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "OverwriteExtensionPack", Flag: "overwrite-extension-pack", Type: "*bool", Required: false},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_start_metadata_model_import = []leanruntime.Field{
	{Name: "MigrationProjectIdentifier", Flag: "migration-project-identifier", Type: "*string", Required: true},
	{Name: "Origin", Flag: "origin", Type: "types.OriginTypeValue", Required: true},
	{Name: "Refresh", Flag: "refresh", Type: "bool", Required: false},
	{Name: "SelectionRules", Flag: "selection-rules", Type: "*string", Required: true},
}

var fields_start_recommendations = []leanruntime.Field{
	{Name: "DatabaseId", Flag: "database-id", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.RecommendationSettings", Required: true},
}

var fields_start_replication = []leanruntime.Field{
	{Name: "CdcStartPosition", Flag: "cdc-start-position", Type: "*string", Required: false},
	{Name: "CdcStartTime", Flag: "cdc-start-time", Type: "*time.Time", Required: false},
	{Name: "CdcStopPosition", Flag: "cdc-stop-position", Type: "*string", Required: false},
	{Name: "PremigrationAssessmentSettings", Flag: "premigration-assessment-settings", Type: "*string", Required: false},
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: true},
	{Name: "StartReplicationType", Flag: "start-replication-type", Type: "*string", Required: true},
}

var fields_start_replication_task = []leanruntime.Field{
	{Name: "CdcStartPosition", Flag: "cdc-start-position", Type: "*string", Required: false},
	{Name: "CdcStartTime", Flag: "cdc-start-time", Type: "*time.Time", Required: false},
	{Name: "CdcStopPosition", Flag: "cdc-stop-position", Type: "*string", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
	{Name: "StartReplicationTaskType", Flag: "start-replication-task-type", Type: "types.StartReplicationTaskTypeValue", Required: true},
}

var fields_start_replication_task_assessment = []leanruntime.Field{
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
}

var fields_start_replication_task_assessment_run = []leanruntime.Field{
	{Name: "AssessmentRunName", Flag: "assessment-run-name", Type: "*string", Required: true},
	{Name: "Exclude", Flag: "exclude", Type: "[]string", Required: false},
	{Name: "IncludeOnly", Flag: "include-only", Type: "[]string", Required: false},
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
	{Name: "ResultEncryptionMode", Flag: "result-encryption-mode", Type: "*string", Required: false},
	{Name: "ResultKmsKeyArn", Flag: "result-kms-key-arn", Type: "*string", Required: false},
	{Name: "ResultLocationBucket", Flag: "result-location-bucket", Type: "*string", Required: true},
	{Name: "ResultLocationFolder", Flag: "result-location-folder", Type: "*string", Required: false},
	{Name: "ServiceAccessRoleArn", Flag: "service-access-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_stop_data_migration = []leanruntime.Field{
	{Name: "DataMigrationIdentifier", Flag: "data-migration-identifier", Type: "*string", Required: true},
}

var fields_stop_replication = []leanruntime.Field{
	{Name: "ReplicationConfigArn", Flag: "replication-config-arn", Type: "*string", Required: true},
}

var fields_stop_replication_task = []leanruntime.Field{
	{Name: "ReplicationTaskArn", Flag: "replication-task-arn", Type: "*string", Required: true},
}

var fields_test_connection = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "ReplicationInstanceArn", Flag: "replication-instance-arn", Type: "*string", Required: true},
}

var fields_update_subscriptions_to_event_bridge = []leanruntime.Field{
	{Name: "ForceMove", Flag: "force-move", Type: "*bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags-to-resource": {
			Name:   "add-tags-to-resource",
			Fields: fields_add_tags_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToResource(ctx, input)
			},
		},
		"apply-pending-maintenance-action": {
			Name:   "apply-pending-maintenance-action",
			Fields: fields_apply_pending_maintenance_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplyPendingMaintenanceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_pending_maintenance_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplyPendingMaintenanceAction(ctx, input)
			},
		},
		"batch-start-recommendations": {
			Name:   "batch-start-recommendations",
			Fields: fields_batch_start_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchStartRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_start_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchStartRecommendations(ctx, input)
			},
		},
		"cancel-metadata-model-conversion": {
			Name:   "cancel-metadata-model-conversion",
			Fields: fields_cancel_metadata_model_conversion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMetadataModelConversionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_metadata_model_conversion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMetadataModelConversion(ctx, input)
			},
		},
		"cancel-metadata-model-creation": {
			Name:   "cancel-metadata-model-creation",
			Fields: fields_cancel_metadata_model_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMetadataModelCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_metadata_model_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMetadataModelCreation(ctx, input)
			},
		},
		"cancel-replication-task-assessment-run": {
			Name:   "cancel-replication-task-assessment-run",
			Fields: fields_cancel_replication_task_assessment_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelReplicationTaskAssessmentRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_replication_task_assessment_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelReplicationTaskAssessmentRun(ctx, input)
			},
		},
		"create-data-migration": {
			Name:   "create-data-migration",
			Fields: fields_create_data_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataMigration(ctx, input)
			},
		},
		"create-data-provider": {
			Name:   "create-data-provider",
			Fields: fields_create_data_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataProvider(ctx, input)
			},
		},
		"create-endpoint": {
			Name:   "create-endpoint",
			Fields: fields_create_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpoint(ctx, input)
			},
		},
		"create-event-subscription": {
			Name:   "create-event-subscription",
			Fields: fields_create_event_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventSubscription(ctx, input)
			},
		},
		"create-fleet-advisor-collector": {
			Name:   "create-fleet-advisor-collector",
			Fields: fields_create_fleet_advisor_collector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetAdvisorCollectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet_advisor_collector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleetAdvisorCollector(ctx, input)
			},
		},
		"create-instance-profile": {
			Name:   "create-instance-profile",
			Fields: fields_create_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceProfile(ctx, input)
			},
		},
		"create-migration-project": {
			Name:   "create-migration-project",
			Fields: fields_create_migration_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMigrationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_migration_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMigrationProject(ctx, input)
			},
		},
		"create-replication-config": {
			Name:   "create-replication-config",
			Fields: fields_create_replication_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationConfig(ctx, input)
			},
		},
		"create-replication-instance": {
			Name:   "create-replication-instance",
			Fields: fields_create_replication_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationInstance(ctx, input)
			},
		},
		"create-replication-subnet-group": {
			Name:   "create-replication-subnet-group",
			Fields: fields_create_replication_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationSubnetGroup(ctx, input)
			},
		},
		"create-replication-task": {
			Name:   "create-replication-task",
			Fields: fields_create_replication_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationTask(ctx, input)
			},
		},
		"delete-certificate": {
			Name:   "delete-certificate",
			Fields: fields_delete_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCertificate(ctx, input)
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
		"delete-data-migration": {
			Name:   "delete-data-migration",
			Fields: fields_delete_data_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataMigration(ctx, input)
			},
		},
		"delete-data-provider": {
			Name:   "delete-data-provider",
			Fields: fields_delete_data_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataProvider(ctx, input)
			},
		},
		"delete-endpoint": {
			Name:   "delete-endpoint",
			Fields: fields_delete_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpoint(ctx, input)
			},
		},
		"delete-event-subscription": {
			Name:   "delete-event-subscription",
			Fields: fields_delete_event_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventSubscription(ctx, input)
			},
		},
		"delete-fleet-advisor-collector": {
			Name:   "delete-fleet-advisor-collector",
			Fields: fields_delete_fleet_advisor_collector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetAdvisorCollectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet_advisor_collector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleetAdvisorCollector(ctx, input)
			},
		},
		"delete-fleet-advisor-databases": {
			Name:   "delete-fleet-advisor-databases",
			Fields: fields_delete_fleet_advisor_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetAdvisorDatabasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet_advisor_databases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleetAdvisorDatabases(ctx, input)
			},
		},
		"delete-instance-profile": {
			Name:   "delete-instance-profile",
			Fields: fields_delete_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceProfile(ctx, input)
			},
		},
		"delete-migration-project": {
			Name:   "delete-migration-project",
			Fields: fields_delete_migration_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMigrationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_migration_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMigrationProject(ctx, input)
			},
		},
		"delete-replication-config": {
			Name:   "delete-replication-config",
			Fields: fields_delete_replication_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationConfig(ctx, input)
			},
		},
		"delete-replication-instance": {
			Name:   "delete-replication-instance",
			Fields: fields_delete_replication_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationInstance(ctx, input)
			},
		},
		"delete-replication-subnet-group": {
			Name:   "delete-replication-subnet-group",
			Fields: fields_delete_replication_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationSubnetGroup(ctx, input)
			},
		},
		"delete-replication-task": {
			Name:   "delete-replication-task",
			Fields: fields_delete_replication_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationTask(ctx, input)
			},
		},
		"delete-replication-task-assessment-run": {
			Name:   "delete-replication-task-assessment-run",
			Fields: fields_delete_replication_task_assessment_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationTaskAssessmentRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_task_assessment_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationTaskAssessmentRun(ctx, input)
			},
		},
		"describe-account-attributes": {
			Name:   "describe-account-attributes",
			Fields: fields_describe_account_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountAttributes(ctx, input)
			},
		},
		"describe-applicable-individual-assessments": {
			Name:   "describe-applicable-individual-assessments",
			Fields: fields_describe_applicable_individual_assessments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicableIndividualAssessmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_applicable_individual_assessments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeApplicableIndividualAssessments(ctx, input)
				}
				var results []*svc.DescribeApplicableIndividualAssessmentsOutput
				p := svc.NewDescribeApplicableIndividualAssessmentsPaginator(client, input)
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
		"describe-certificates": {
			Name:   "describe-certificates",
			Fields: fields_describe_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCertificates(ctx, input)
				}
				var results []*svc.DescribeCertificatesOutput
				p := svc.NewDescribeCertificatesPaginator(client, input)
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
		"describe-connections": {
			Name:   "describe-connections",
			Fields: fields_describe_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConnections(ctx, input)
				}
				var results []*svc.DescribeConnectionsOutput
				p := svc.NewDescribeConnectionsPaginator(client, input)
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
		"describe-conversion-configuration": {
			Name:   "describe-conversion-configuration",
			Fields: fields_describe_conversion_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConversionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_conversion_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConversionConfiguration(ctx, input)
			},
		},
		"describe-data-migrations": {
			Name:   "describe-data-migrations",
			Fields: fields_describe_data_migrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataMigrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_migrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataMigrations(ctx, input)
				}
				var results []*svc.DescribeDataMigrationsOutput
				p := svc.NewDescribeDataMigrationsPaginator(client, input)
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
		"describe-data-providers": {
			Name:   "describe-data-providers",
			Fields: fields_describe_data_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataProviders(ctx, input)
				}
				var results []*svc.DescribeDataProvidersOutput
				p := svc.NewDescribeDataProvidersPaginator(client, input)
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
		"describe-endpoint-settings": {
			Name:   "describe-endpoint-settings",
			Fields: fields_describe_endpoint_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_endpoint_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEndpointSettings(ctx, input)
				}
				var results []*svc.DescribeEndpointSettingsOutput
				p := svc.NewDescribeEndpointSettingsPaginator(client, input)
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
		"describe-endpoint-types": {
			Name:   "describe-endpoint-types",
			Fields: fields_describe_endpoint_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_endpoint_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEndpointTypes(ctx, input)
				}
				var results []*svc.DescribeEndpointTypesOutput
				p := svc.NewDescribeEndpointTypesPaginator(client, input)
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
		"describe-endpoints": {
			Name:   "describe-endpoints",
			Fields: fields_describe_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEndpoints(ctx, input)
				}
				var results []*svc.DescribeEndpointsOutput
				p := svc.NewDescribeEndpointsPaginator(client, input)
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
		"describe-engine-versions": {
			Name:   "describe-engine-versions",
			Fields: fields_describe_engine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEngineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_engine_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEngineVersions(ctx, input)
				}
				var results []*svc.DescribeEngineVersionsOutput
				p := svc.NewDescribeEngineVersionsPaginator(client, input)
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
		"describe-event-categories": {
			Name:   "describe-event-categories",
			Fields: fields_describe_event_categories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventCategoriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_categories, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventCategories(ctx, input)
			},
		},
		"describe-event-subscriptions": {
			Name:   "describe-event-subscriptions",
			Fields: fields_describe_event_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_event_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEventSubscriptions(ctx, input)
				}
				var results []*svc.DescribeEventSubscriptionsOutput
				p := svc.NewDescribeEventSubscriptionsPaginator(client, input)
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
		"describe-events": {
			Name:   "describe-events",
			Fields: fields_describe_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEvents(ctx, input)
				}
				var results []*svc.DescribeEventsOutput
				p := svc.NewDescribeEventsPaginator(client, input)
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
		"describe-extension-pack-associations": {
			Name:   "describe-extension-pack-associations",
			Fields: fields_describe_extension_pack_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExtensionPackAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_extension_pack_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeExtensionPackAssociations(ctx, input)
				}
				var results []*svc.DescribeExtensionPackAssociationsOutput
				p := svc.NewDescribeExtensionPackAssociationsPaginator(client, input)
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
		"describe-fleet-advisor-collectors": {
			Name:   "describe-fleet-advisor-collectors",
			Fields: fields_describe_fleet_advisor_collectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetAdvisorCollectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_advisor_collectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetAdvisorCollectors(ctx, input)
				}
				var results []*svc.DescribeFleetAdvisorCollectorsOutput
				p := svc.NewDescribeFleetAdvisorCollectorsPaginator(client, input)
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
		"describe-fleet-advisor-databases": {
			Name:   "describe-fleet-advisor-databases",
			Fields: fields_describe_fleet_advisor_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetAdvisorDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_advisor_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetAdvisorDatabases(ctx, input)
				}
				var results []*svc.DescribeFleetAdvisorDatabasesOutput
				p := svc.NewDescribeFleetAdvisorDatabasesPaginator(client, input)
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
		"describe-fleet-advisor-lsa-analysis": {
			Name:   "describe-fleet-advisor-lsa-analysis",
			Fields: fields_describe_fleet_advisor_lsa_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetAdvisorLsaAnalysisInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_advisor_lsa_analysis, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetAdvisorLsaAnalysis(ctx, input)
				}
				var results []*svc.DescribeFleetAdvisorLsaAnalysisOutput
				p := svc.NewDescribeFleetAdvisorLsaAnalysisPaginator(client, input)
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
		"describe-fleet-advisor-schema-object-summary": {
			Name:   "describe-fleet-advisor-schema-object-summary",
			Fields: fields_describe_fleet_advisor_schema_object_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetAdvisorSchemaObjectSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_advisor_schema_object_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetAdvisorSchemaObjectSummary(ctx, input)
				}
				var results []*svc.DescribeFleetAdvisorSchemaObjectSummaryOutput
				p := svc.NewDescribeFleetAdvisorSchemaObjectSummaryPaginator(client, input)
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
		"describe-fleet-advisor-schemas": {
			Name:   "describe-fleet-advisor-schemas",
			Fields: fields_describe_fleet_advisor_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetAdvisorSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_advisor_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetAdvisorSchemas(ctx, input)
				}
				var results []*svc.DescribeFleetAdvisorSchemasOutput
				p := svc.NewDescribeFleetAdvisorSchemasPaginator(client, input)
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
		"describe-instance-profiles": {
			Name:   "describe-instance-profiles",
			Fields: fields_describe_instance_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceProfiles(ctx, input)
				}
				var results []*svc.DescribeInstanceProfilesOutput
				p := svc.NewDescribeInstanceProfilesPaginator(client, input)
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
		"describe-metadata-model": {
			Name:   "describe-metadata-model",
			Fields: fields_describe_metadata_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_metadata_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMetadataModel(ctx, input)
			},
		},
		"describe-metadata-model-assessments": {
			Name:   "describe-metadata-model-assessments",
			Fields: fields_describe_metadata_model_assessments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelAssessmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_assessments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelAssessments(ctx, input)
				}
				var results []*svc.DescribeMetadataModelAssessmentsOutput
				p := svc.NewDescribeMetadataModelAssessmentsPaginator(client, input)
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
		"describe-metadata-model-children": {
			Name:   "describe-metadata-model-children",
			Fields: fields_describe_metadata_model_children,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelChildrenInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_children, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelChildren(ctx, input)
				}
				var results []*svc.DescribeMetadataModelChildrenOutput
				p := svc.NewDescribeMetadataModelChildrenPaginator(client, input)
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
		"describe-metadata-model-conversions": {
			Name:   "describe-metadata-model-conversions",
			Fields: fields_describe_metadata_model_conversions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelConversionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_conversions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelConversions(ctx, input)
				}
				var results []*svc.DescribeMetadataModelConversionsOutput
				p := svc.NewDescribeMetadataModelConversionsPaginator(client, input)
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
		"describe-metadata-model-creations": {
			Name:   "describe-metadata-model-creations",
			Fields: fields_describe_metadata_model_creations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelCreationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_creations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelCreations(ctx, input)
				}
				var results []*svc.DescribeMetadataModelCreationsOutput
				p := svc.NewDescribeMetadataModelCreationsPaginator(client, input)
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
		"describe-metadata-model-exports-as-script": {
			Name:   "describe-metadata-model-exports-as-script",
			Fields: fields_describe_metadata_model_exports_as_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelExportsAsScriptInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_exports_as_script, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelExportsAsScript(ctx, input)
				}
				var results []*svc.DescribeMetadataModelExportsAsScriptOutput
				p := svc.NewDescribeMetadataModelExportsAsScriptPaginator(client, input)
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
		"describe-metadata-model-exports-to-target": {
			Name:   "describe-metadata-model-exports-to-target",
			Fields: fields_describe_metadata_model_exports_to_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelExportsToTargetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_exports_to_target, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelExportsToTarget(ctx, input)
				}
				var results []*svc.DescribeMetadataModelExportsToTargetOutput
				p := svc.NewDescribeMetadataModelExportsToTargetPaginator(client, input)
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
		"describe-metadata-model-imports": {
			Name:   "describe-metadata-model-imports",
			Fields: fields_describe_metadata_model_imports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetadataModelImportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_metadata_model_imports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMetadataModelImports(ctx, input)
				}
				var results []*svc.DescribeMetadataModelImportsOutput
				p := svc.NewDescribeMetadataModelImportsPaginator(client, input)
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
		"describe-migration-projects": {
			Name:   "describe-migration-projects",
			Fields: fields_describe_migration_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMigrationProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_migration_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMigrationProjects(ctx, input)
				}
				var results []*svc.DescribeMigrationProjectsOutput
				p := svc.NewDescribeMigrationProjectsPaginator(client, input)
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
		"describe-orderable-replication-instances": {
			Name:   "describe-orderable-replication-instances",
			Fields: fields_describe_orderable_replication_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrderableReplicationInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_orderable_replication_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrderableReplicationInstances(ctx, input)
				}
				var results []*svc.DescribeOrderableReplicationInstancesOutput
				p := svc.NewDescribeOrderableReplicationInstancesPaginator(client, input)
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
		"describe-pending-maintenance-actions": {
			Name:   "describe-pending-maintenance-actions",
			Fields: fields_describe_pending_maintenance_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePendingMaintenanceActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_pending_maintenance_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePendingMaintenanceActions(ctx, input)
				}
				var results []*svc.DescribePendingMaintenanceActionsOutput
				p := svc.NewDescribePendingMaintenanceActionsPaginator(client, input)
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
		"describe-recommendation-limitations": {
			Name:   "describe-recommendation-limitations",
			Fields: fields_describe_recommendation_limitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecommendationLimitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_recommendation_limitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRecommendationLimitations(ctx, input)
				}
				var results []*svc.DescribeRecommendationLimitationsOutput
				p := svc.NewDescribeRecommendationLimitationsPaginator(client, input)
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
		"describe-recommendations": {
			Name:   "describe-recommendations",
			Fields: fields_describe_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRecommendations(ctx, input)
				}
				var results []*svc.DescribeRecommendationsOutput
				p := svc.NewDescribeRecommendationsPaginator(client, input)
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
		"describe-refresh-schemas-status": {
			Name:   "describe-refresh-schemas-status",
			Fields: fields_describe_refresh_schemas_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRefreshSchemasStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_refresh_schemas_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRefreshSchemasStatus(ctx, input)
			},
		},
		"describe-replication-configs": {
			Name:   "describe-replication-configs",
			Fields: fields_describe_replication_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationConfigs(ctx, input)
				}
				var results []*svc.DescribeReplicationConfigsOutput
				p := svc.NewDescribeReplicationConfigsPaginator(client, input)
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
		"describe-replication-instance-task-logs": {
			Name:   "describe-replication-instance-task-logs",
			Fields: fields_describe_replication_instance_task_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationInstanceTaskLogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_instance_task_logs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationInstanceTaskLogs(ctx, input)
				}
				var results []*svc.DescribeReplicationInstanceTaskLogsOutput
				p := svc.NewDescribeReplicationInstanceTaskLogsPaginator(client, input)
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
		"describe-replication-instances": {
			Name:   "describe-replication-instances",
			Fields: fields_describe_replication_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationInstances(ctx, input)
				}
				var results []*svc.DescribeReplicationInstancesOutput
				p := svc.NewDescribeReplicationInstancesPaginator(client, input)
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
		"describe-replication-subnet-groups": {
			Name:   "describe-replication-subnet-groups",
			Fields: fields_describe_replication_subnet_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationSubnetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_subnet_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationSubnetGroups(ctx, input)
				}
				var results []*svc.DescribeReplicationSubnetGroupsOutput
				p := svc.NewDescribeReplicationSubnetGroupsPaginator(client, input)
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
		"describe-replication-table-statistics": {
			Name:   "describe-replication-table-statistics",
			Fields: fields_describe_replication_table_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationTableStatisticsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_table_statistics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationTableStatistics(ctx, input)
				}
				var results []*svc.DescribeReplicationTableStatisticsOutput
				p := svc.NewDescribeReplicationTableStatisticsPaginator(client, input)
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
		"describe-replication-task-assessment-results": {
			Name:   "describe-replication-task-assessment-results",
			Fields: fields_describe_replication_task_assessment_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationTaskAssessmentResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_task_assessment_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationTaskAssessmentResults(ctx, input)
				}
				var results []*svc.DescribeReplicationTaskAssessmentResultsOutput
				p := svc.NewDescribeReplicationTaskAssessmentResultsPaginator(client, input)
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
		"describe-replication-task-assessment-runs": {
			Name:   "describe-replication-task-assessment-runs",
			Fields: fields_describe_replication_task_assessment_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationTaskAssessmentRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_task_assessment_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationTaskAssessmentRuns(ctx, input)
				}
				var results []*svc.DescribeReplicationTaskAssessmentRunsOutput
				p := svc.NewDescribeReplicationTaskAssessmentRunsPaginator(client, input)
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
		"describe-replication-task-individual-assessments": {
			Name:   "describe-replication-task-individual-assessments",
			Fields: fields_describe_replication_task_individual_assessments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationTaskIndividualAssessmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_task_individual_assessments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationTaskIndividualAssessments(ctx, input)
				}
				var results []*svc.DescribeReplicationTaskIndividualAssessmentsOutput
				p := svc.NewDescribeReplicationTaskIndividualAssessmentsPaginator(client, input)
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
		"describe-replication-tasks": {
			Name:   "describe-replication-tasks",
			Fields: fields_describe_replication_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationTasks(ctx, input)
				}
				var results []*svc.DescribeReplicationTasksOutput
				p := svc.NewDescribeReplicationTasksPaginator(client, input)
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
		"describe-replications": {
			Name:   "describe-replications",
			Fields: fields_describe_replications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplications(ctx, input)
				}
				var results []*svc.DescribeReplicationsOutput
				p := svc.NewDescribeReplicationsPaginator(client, input)
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
		"describe-schemas": {
			Name:   "describe-schemas",
			Fields: fields_describe_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSchemas(ctx, input)
				}
				var results []*svc.DescribeSchemasOutput
				p := svc.NewDescribeSchemasPaginator(client, input)
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
		"describe-table-statistics": {
			Name:   "describe-table-statistics",
			Fields: fields_describe_table_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTableStatisticsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_table_statistics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTableStatistics(ctx, input)
				}
				var results []*svc.DescribeTableStatisticsOutput
				p := svc.NewDescribeTableStatisticsPaginator(client, input)
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
		"export-metadata-model-assessment": {
			Name:   "export-metadata-model-assessment",
			Fields: fields_export_metadata_model_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportMetadataModelAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_metadata_model_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportMetadataModelAssessment(ctx, input)
			},
		},
		"get-target-selection-rules": {
			Name:   "get-target-selection-rules",
			Fields: fields_get_target_selection_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTargetSelectionRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_target_selection_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTargetSelectionRules(ctx, input)
			},
		},
		"import-certificate": {
			Name:   "import-certificate",
			Fields: fields_import_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCertificate(ctx, input)
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
		"modify-conversion-configuration": {
			Name:   "modify-conversion-configuration",
			Fields: fields_modify_conversion_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyConversionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_conversion_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyConversionConfiguration(ctx, input)
			},
		},
		"modify-data-migration": {
			Name:   "modify-data-migration",
			Fields: fields_modify_data_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDataMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_data_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDataMigration(ctx, input)
			},
		},
		"modify-data-provider": {
			Name:   "modify-data-provider",
			Fields: fields_modify_data_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDataProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_data_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDataProvider(ctx, input)
			},
		},
		"modify-endpoint": {
			Name:   "modify-endpoint",
			Fields: fields_modify_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyEndpoint(ctx, input)
			},
		},
		"modify-event-subscription": {
			Name:   "modify-event-subscription",
			Fields: fields_modify_event_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyEventSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_event_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyEventSubscription(ctx, input)
			},
		},
		"modify-instance-profile": {
			Name:   "modify-instance-profile",
			Fields: fields_modify_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceProfile(ctx, input)
			},
		},
		"modify-migration-project": {
			Name:   "modify-migration-project",
			Fields: fields_modify_migration_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyMigrationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_migration_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyMigrationProject(ctx, input)
			},
		},
		"modify-replication-config": {
			Name:   "modify-replication-config",
			Fields: fields_modify_replication_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReplicationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_replication_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReplicationConfig(ctx, input)
			},
		},
		"modify-replication-instance": {
			Name:   "modify-replication-instance",
			Fields: fields_modify_replication_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_replication_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReplicationInstance(ctx, input)
			},
		},
		"modify-replication-subnet-group": {
			Name:   "modify-replication-subnet-group",
			Fields: fields_modify_replication_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReplicationSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_replication_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReplicationSubnetGroup(ctx, input)
			},
		},
		"modify-replication-task": {
			Name:   "modify-replication-task",
			Fields: fields_modify_replication_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReplicationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_replication_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReplicationTask(ctx, input)
			},
		},
		"move-replication-task": {
			Name:   "move-replication-task",
			Fields: fields_move_replication_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MoveReplicationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_move_replication_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MoveReplicationTask(ctx, input)
			},
		},
		"reboot-replication-instance": {
			Name:   "reboot-replication-instance",
			Fields: fields_reboot_replication_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootReplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_replication_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootReplicationInstance(ctx, input)
			},
		},
		"refresh-schemas": {
			Name:   "refresh-schemas",
			Fields: fields_refresh_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RefreshSchemasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_refresh_schemas, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RefreshSchemas(ctx, input)
			},
		},
		"reload-replication-tables": {
			Name:   "reload-replication-tables",
			Fields: fields_reload_replication_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReloadReplicationTablesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reload_replication_tables, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReloadReplicationTables(ctx, input)
			},
		},
		"reload-tables": {
			Name:   "reload-tables",
			Fields: fields_reload_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReloadTablesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reload_tables, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReloadTables(ctx, input)
			},
		},
		"remove-tags-from-resource": {
			Name:   "remove-tags-from-resource",
			Fields: fields_remove_tags_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromResource(ctx, input)
			},
		},
		"run-fleet-advisor-lsa-analysis": {
			Name:   "run-fleet-advisor-lsa-analysis",
			Fields: fields_run_fleet_advisor_lsa_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RunFleetAdvisorLsaAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_run_fleet_advisor_lsa_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RunFleetAdvisorLsaAnalysis(ctx, input)
			},
		},
		"start-data-migration": {
			Name:   "start-data-migration",
			Fields: fields_start_data_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataMigration(ctx, input)
			},
		},
		"start-extension-pack-association": {
			Name:   "start-extension-pack-association",
			Fields: fields_start_extension_pack_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExtensionPackAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_extension_pack_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExtensionPackAssociation(ctx, input)
			},
		},
		"start-metadata-model-assessment": {
			Name:   "start-metadata-model-assessment",
			Fields: fields_start_metadata_model_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataModelAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_model_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataModelAssessment(ctx, input)
			},
		},
		"start-metadata-model-conversion": {
			Name:   "start-metadata-model-conversion",
			Fields: fields_start_metadata_model_conversion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataModelConversionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_model_conversion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataModelConversion(ctx, input)
			},
		},
		"start-metadata-model-creation": {
			Name:   "start-metadata-model-creation",
			Fields: fields_start_metadata_model_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataModelCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_model_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataModelCreation(ctx, input)
			},
		},
		"start-metadata-model-export-as-script": {
			Name:   "start-metadata-model-export-as-script",
			Fields: fields_start_metadata_model_export_as_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataModelExportAsScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_model_export_as_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataModelExportAsScript(ctx, input)
			},
		},
		"start-metadata-model-export-to-target": {
			Name:   "start-metadata-model-export-to-target",
			Fields: fields_start_metadata_model_export_to_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataModelExportToTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_model_export_to_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataModelExportToTarget(ctx, input)
			},
		},
		"start-metadata-model-import": {
			Name:   "start-metadata-model-import",
			Fields: fields_start_metadata_model_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataModelImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_model_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataModelImport(ctx, input)
			},
		},
		"start-recommendations": {
			Name:   "start-recommendations",
			Fields: fields_start_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRecommendations(ctx, input)
			},
		},
		"start-replication": {
			Name:   "start-replication",
			Fields: fields_start_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReplication(ctx, input)
			},
		},
		"start-replication-task": {
			Name:   "start-replication-task",
			Fields: fields_start_replication_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReplicationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_replication_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReplicationTask(ctx, input)
			},
		},
		"start-replication-task-assessment": {
			Name:   "start-replication-task-assessment",
			Fields: fields_start_replication_task_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReplicationTaskAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_replication_task_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReplicationTaskAssessment(ctx, input)
			},
		},
		"start-replication-task-assessment-run": {
			Name:   "start-replication-task-assessment-run",
			Fields: fields_start_replication_task_assessment_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReplicationTaskAssessmentRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_replication_task_assessment_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReplicationTaskAssessmentRun(ctx, input)
			},
		},
		"stop-data-migration": {
			Name:   "stop-data-migration",
			Fields: fields_stop_data_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDataMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_data_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDataMigration(ctx, input)
			},
		},
		"stop-replication": {
			Name:   "stop-replication",
			Fields: fields_stop_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopReplication(ctx, input)
			},
		},
		"stop-replication-task": {
			Name:   "stop-replication-task",
			Fields: fields_stop_replication_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopReplicationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_replication_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopReplicationTask(ctx, input)
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
		"update-subscriptions-to-event-bridge": {
			Name:   "update-subscriptions-to-event-bridge",
			Fields: fields_update_subscriptions_to_event_bridge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionsToEventBridgeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscriptions_to_event_bridge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriptionsToEventBridge(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("databasemigrationservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
