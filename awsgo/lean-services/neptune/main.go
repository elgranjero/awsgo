package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/neptune"
)

var fields_add_role_to_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_add_source_identifier_to_subscription = []leanruntime.Field{
	{Name: "SourceIdentifier", Flag: "source-identifier", Type: "*string", Required: true},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_apply_pending_maintenance_action = []leanruntime.Field{
	{Name: "ApplyAction", Flag: "apply-action", Type: "*string", Required: true},
	{Name: "OptInType", Flag: "opt-in-type", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_copy_db_cluster_parameter_group = []leanruntime.Field{
	{Name: "SourceDBClusterParameterGroupIdentifier", Flag: "source-db-cluster-parameter-group-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDBClusterParameterGroupDescription", Flag: "target-db-cluster-parameter-group-description", Type: "*string", Required: true},
	{Name: "TargetDBClusterParameterGroupIdentifier", Flag: "target-db-cluster-parameter-group-identifier", Type: "*string", Required: true},
}

var fields_copy_db_cluster_snapshot = []leanruntime.Field{
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PreSignedUrl", Flag: "pre-signed-url", Type: "*string", Required: false},
	{Name: "SourceDBClusterSnapshotIdentifier", Flag: "source-db-cluster-snapshot-identifier", Type: "*string", Required: true},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDBClusterSnapshotIdentifier", Flag: "target-db-cluster-snapshot-identifier", Type: "*string", Required: true},
	{Name: "destinationRegion", Flag: "destination-region", Type: "*string", Required: false},
}

var fields_copy_db_parameter_group = []leanruntime.Field{
	{Name: "SourceDBParameterGroupIdentifier", Flag: "source-db-parameter-group-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDBParameterGroupDescription", Flag: "target-db-parameter-group-description", Type: "*string", Required: true},
	{Name: "TargetDBParameterGroupIdentifier", Flag: "target-db-parameter-group-identifier", Type: "*string", Required: true},
}

var fields_create_db_cluster = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CharacterSetName", Flag: "character-set-name", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreSignedUrl", Flag: "pre-signed-url", Type: "*string", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ReplicationSourceIdentifier", Flag: "replication-source-identifier", Type: "*string", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: false},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
	{Name: "destinationRegion", Flag: "destination-region", Type: "*string", Required: false},
}

var fields_create_db_cluster_endpoint = []leanruntime.Field{
	{Name: "DBClusterEndpointIdentifier", Flag: "db-cluster-endpoint-identifier", Type: "*string", Required: true},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "*string", Required: true},
	{Name: "ExcludedMembers", Flag: "excluded-members", Type: "[]string", Required: false},
	{Name: "StaticMembers", Flag: "static-members", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_cluster_parameter_group = []leanruntime.Field{
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: true},
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_cluster_snapshot = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterSnapshotIdentifier", Flag: "db-cluster-snapshot-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_instance = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CharacterSetName", Flag: "character-set-name", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: true},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSecurityGroups", Flag: "db-security-groups", Type: "[]string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PromotionTier", Flag: "promotion-tier", Type: "*int32", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TdeCredentialArn", Flag: "tde-credential-arn", Type: "*string", Required: false},
	{Name: "TdeCredentialPassword", Flag: "tde-credential-password", Type: "*string", Required: false},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_create_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: true},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_subnet_group = []leanruntime.Field{
	{Name: "DBSubnetGroupDescription", Flag: "db-subnet-group-description", Type: "*string", Required: true},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
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

var fields_create_global_cluster = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
	{Name: "SourceDBClusterIdentifier", Flag: "source-db-cluster-identifier", Type: "*string", Required: false},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "FinalDBSnapshotIdentifier", Flag: "final-db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "SkipFinalSnapshot", Flag: "skip-final-snapshot", Type: "*bool", Required: false},
}

var fields_delete_db_cluster_endpoint = []leanruntime.Field{
	{Name: "DBClusterEndpointIdentifier", Flag: "db-cluster-endpoint-identifier", Type: "*string", Required: true},
}

var fields_delete_db_cluster_parameter_group = []leanruntime.Field{
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_db_cluster_snapshot = []leanruntime.Field{
	{Name: "DBClusterSnapshotIdentifier", Flag: "db-cluster-snapshot-identifier", Type: "*string", Required: true},
}

var fields_delete_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "FinalDBSnapshotIdentifier", Flag: "final-db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "SkipFinalSnapshot", Flag: "skip-final-snapshot", Type: "*bool", Required: false},
}

var fields_delete_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_db_subnet_group = []leanruntime.Field{
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: true},
}

var fields_delete_event_subscription = []leanruntime.Field{
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_delete_global_cluster = []leanruntime.Field{
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
}

var fields_describe_db_cluster_endpoints = []leanruntime.Field{
	{Name: "DBClusterEndpointIdentifier", Flag: "db-cluster-endpoint-identifier", Type: "*string", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_cluster_parameter_groups = []leanruntime.Field{
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_cluster_parameters = []leanruntime.Field{
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_db_cluster_snapshot_attributes = []leanruntime.Field{
	{Name: "DBClusterSnapshotIdentifier", Flag: "db-cluster-snapshot-identifier", Type: "*string", Required: true},
}

var fields_describe_db_cluster_snapshots = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: false},
	{Name: "DBClusterSnapshotIdentifier", Flag: "db-cluster-snapshot-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludePublic", Flag: "include-public", Type: "*bool", Required: false},
	{Name: "IncludeShared", Flag: "include-shared", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SnapshotType", Flag: "snapshot-type", Type: "*string", Required: false},
}

var fields_describe_db_clusters = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_engine_versions = []leanruntime.Field{
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: false},
	{Name: "DefaultOnly", Flag: "default-only", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ListSupportedCharacterSets", Flag: "list-supported-character-sets", Type: "*bool", Required: false},
	{Name: "ListSupportedTimezones", Flag: "list-supported-timezones", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_instances = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_parameter_groups = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_parameters = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_db_subnet_groups = []leanruntime.Field{
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_engine_default_cluster_parameters = []leanruntime.Field{
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_engine_default_parameters = []leanruntime.Field{
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
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

var fields_describe_global_clusters = []leanruntime.Field{
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_orderable_db_instance_options = []leanruntime.Field{
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "Vpc", Flag: "vpc", Type: "*bool", Required: false},
}

var fields_describe_pending_maintenance_actions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
}

var fields_describe_valid_db_instance_modifications = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
}

var fields_failover_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: false},
	{Name: "TargetDBInstanceIdentifier", Flag: "target-db-instance-identifier", Type: "*string", Required: false},
}

var fields_failover_global_cluster = []leanruntime.Field{
	{Name: "AllowDataLoss", Flag: "allow-data-loss", Type: "*bool", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
	{Name: "Switchover", Flag: "switchover", Type: "*bool", Required: false},
	{Name: "TargetDbClusterIdentifier", Flag: "target-db-cluster-identifier", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_modify_db_cluster = []leanruntime.Field{
	{Name: "AllowMajorVersionUpgrade", Flag: "allow-major-version-upgrade", Type: "*bool", Required: false},
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CloudwatchLogsExportConfiguration", Flag: "cloudwatch-logs-export-configuration", Type: "*types.CloudwatchLogsExportConfiguration", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBInstanceParameterGroupName", Flag: "db-instance-parameter-group-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "NewDBClusterIdentifier", Flag: "new-db-cluster-identifier", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_db_cluster_endpoint = []leanruntime.Field{
	{Name: "DBClusterEndpointIdentifier", Flag: "db-cluster-endpoint-identifier", Type: "*string", Required: true},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "*string", Required: false},
	{Name: "ExcludedMembers", Flag: "excluded-members", Type: "[]string", Required: false},
	{Name: "StaticMembers", Flag: "static-members", Type: "[]string", Required: false},
}

var fields_modify_db_cluster_parameter_group = []leanruntime.Field{
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: true},
}

var fields_modify_db_cluster_snapshot_attribute = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DBClusterSnapshotIdentifier", Flag: "db-cluster-snapshot-identifier", Type: "*string", Required: true},
	{Name: "ValuesToAdd", Flag: "values-to-add", Type: "[]string", Required: false},
	{Name: "ValuesToRemove", Flag: "values-to-remove", Type: "[]string", Required: false},
}

var fields_modify_db_instance = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AllowMajorVersionUpgrade", Flag: "allow-major-version-upgrade", Type: "*bool", Required: false},
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CloudwatchLogsExportConfiguration", Flag: "cloudwatch-logs-export-configuration", Type: "*types.CloudwatchLogsExportConfiguration", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBPortNumber", Flag: "db-port-number", Type: "*int32", Required: false},
	{Name: "DBSecurityGroups", Flag: "db-security-groups", Type: "[]string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NewDBInstanceIdentifier", Flag: "new-db-instance-identifier", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PromotionTier", Flag: "promotion-tier", Type: "*int32", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TdeCredentialArn", Flag: "tde-credential-arn", Type: "*string", Required: false},
	{Name: "TdeCredentialPassword", Flag: "tde-credential-password", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: true},
}

var fields_modify_db_subnet_group = []leanruntime.Field{
	{Name: "DBSubnetGroupDescription", Flag: "db-subnet-group-description", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
}

var fields_modify_event_subscription = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventCategories", Flag: "event-categories", Type: "[]string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_modify_global_cluster = []leanruntime.Field{
	{Name: "AllowMajorVersionUpgrade", Flag: "allow-major-version-upgrade", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
	{Name: "NewGlobalClusterIdentifier", Flag: "new-global-cluster-identifier", Type: "*string", Required: false},
}

var fields_promote_read_replica_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_reboot_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "ForceFailover", Flag: "force-failover", Type: "*bool", Required: false},
}

var fields_remove_from_global_cluster = []leanruntime.Field{
	{Name: "DbClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
}

var fields_remove_role_from_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_remove_source_identifier_from_subscription = []leanruntime.Field{
	{Name: "SourceIdentifier", Flag: "source-identifier", Type: "*string", Required: true},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_reset_db_cluster_parameter_group = []leanruntime.Field{
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "ResetAllParameters", Flag: "reset-all-parameters", Type: "*bool", Required: false},
}

var fields_reset_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "ResetAllParameters", Flag: "reset-all-parameters", Type: "*bool", Required: false},
}

var fields_restore_db_cluster_from_snapshot = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_db_cluster_to_point_in_time = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "RestoreToTime", Flag: "restore-to-time", Type: "*time.Time", Required: false},
	{Name: "RestoreType", Flag: "restore-type", Type: "*string", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SourceDBClusterIdentifier", Flag: "source-db-cluster-identifier", Type: "*string", Required: true},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseLatestRestorableTime", Flag: "use-latest-restorable-time", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_start_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_stop_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_switchover_global_cluster = []leanruntime.Field{
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
	{Name: "TargetDbClusterIdentifier", Flag: "target-db-cluster-identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-role-to-db-cluster": {
			Name:   "add-role-to-db-cluster",
			Fields: fields_add_role_to_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddRoleToDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_role_to_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddRoleToDBCluster(ctx, input)
			},
		},
		"add-source-identifier-to-subscription": {
			Name:   "add-source-identifier-to-subscription",
			Fields: fields_add_source_identifier_to_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddSourceIdentifierToSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_source_identifier_to_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddSourceIdentifierToSubscription(ctx, input)
			},
		},
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
		"copy-db-cluster-parameter-group": {
			Name:   "copy-db-cluster-parameter-group",
			Fields: fields_copy_db_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyDBClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_db_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyDBClusterParameterGroup(ctx, input)
			},
		},
		"copy-db-cluster-snapshot": {
			Name:   "copy-db-cluster-snapshot",
			Fields: fields_copy_db_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyDBClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_db_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyDBClusterSnapshot(ctx, input)
			},
		},
		"copy-db-parameter-group": {
			Name:   "copy-db-parameter-group",
			Fields: fields_copy_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyDBParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyDBParameterGroup(ctx, input)
			},
		},
		"create-db-cluster": {
			Name:   "create-db-cluster",
			Fields: fields_create_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBCluster(ctx, input)
			},
		},
		"create-db-cluster-endpoint": {
			Name:   "create-db-cluster-endpoint",
			Fields: fields_create_db_cluster_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBClusterEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_cluster_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBClusterEndpoint(ctx, input)
			},
		},
		"create-db-cluster-parameter-group": {
			Name:   "create-db-cluster-parameter-group",
			Fields: fields_create_db_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBClusterParameterGroup(ctx, input)
			},
		},
		"create-db-cluster-snapshot": {
			Name:   "create-db-cluster-snapshot",
			Fields: fields_create_db_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBClusterSnapshot(ctx, input)
			},
		},
		"create-db-instance": {
			Name:   "create-db-instance",
			Fields: fields_create_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBInstance(ctx, input)
			},
		},
		"create-db-parameter-group": {
			Name:   "create-db-parameter-group",
			Fields: fields_create_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBParameterGroup(ctx, input)
			},
		},
		"create-db-subnet-group": {
			Name:   "create-db-subnet-group",
			Fields: fields_create_db_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBSubnetGroup(ctx, input)
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
		"create-global-cluster": {
			Name:   "create-global-cluster",
			Fields: fields_create_global_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlobalClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_global_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlobalCluster(ctx, input)
			},
		},
		"delete-db-cluster": {
			Name:   "delete-db-cluster",
			Fields: fields_delete_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBCluster(ctx, input)
			},
		},
		"delete-db-cluster-endpoint": {
			Name:   "delete-db-cluster-endpoint",
			Fields: fields_delete_db_cluster_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBClusterEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_cluster_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBClusterEndpoint(ctx, input)
			},
		},
		"delete-db-cluster-parameter-group": {
			Name:   "delete-db-cluster-parameter-group",
			Fields: fields_delete_db_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBClusterParameterGroup(ctx, input)
			},
		},
		"delete-db-cluster-snapshot": {
			Name:   "delete-db-cluster-snapshot",
			Fields: fields_delete_db_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBClusterSnapshot(ctx, input)
			},
		},
		"delete-db-instance": {
			Name:   "delete-db-instance",
			Fields: fields_delete_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBInstance(ctx, input)
			},
		},
		"delete-db-parameter-group": {
			Name:   "delete-db-parameter-group",
			Fields: fields_delete_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBParameterGroup(ctx, input)
			},
		},
		"delete-db-subnet-group": {
			Name:   "delete-db-subnet-group",
			Fields: fields_delete_db_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBSubnetGroup(ctx, input)
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
		"delete-global-cluster": {
			Name:   "delete-global-cluster",
			Fields: fields_delete_global_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlobalClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_global_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlobalCluster(ctx, input)
			},
		},
		"describe-db-cluster-endpoints": {
			Name:   "describe-db-cluster-endpoints",
			Fields: fields_describe_db_cluster_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusterEndpoints(ctx, input)
				}
				var results []*svc.DescribeDBClusterEndpointsOutput
				p := svc.NewDescribeDBClusterEndpointsPaginator(client, input)
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
		"describe-db-cluster-parameter-groups": {
			Name:   "describe-db-cluster-parameter-groups",
			Fields: fields_describe_db_cluster_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterParameterGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_parameter_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusterParameterGroups(ctx, input)
				}
				var results []*svc.DescribeDBClusterParameterGroupsOutput
				p := svc.NewDescribeDBClusterParameterGroupsPaginator(client, input)
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
		"describe-db-cluster-parameters": {
			Name:   "describe-db-cluster-parameters",
			Fields: fields_describe_db_cluster_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusterParameters(ctx, input)
				}
				var results []*svc.DescribeDBClusterParametersOutput
				p := svc.NewDescribeDBClusterParametersPaginator(client, input)
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
		"describe-db-cluster-snapshot-attributes": {
			Name:   "describe-db-cluster-snapshot-attributes",
			Fields: fields_describe_db_cluster_snapshot_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterSnapshotAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_snapshot_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDBClusterSnapshotAttributes(ctx, input)
			},
		},
		"describe-db-cluster-snapshots": {
			Name:   "describe-db-cluster-snapshots",
			Fields: fields_describe_db_cluster_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusterSnapshots(ctx, input)
				}
				var results []*svc.DescribeDBClusterSnapshotsOutput
				p := svc.NewDescribeDBClusterSnapshotsPaginator(client, input)
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
		"describe-db-clusters": {
			Name:   "describe-db-clusters",
			Fields: fields_describe_db_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusters(ctx, input)
				}
				var results []*svc.DescribeDBClustersOutput
				p := svc.NewDescribeDBClustersPaginator(client, input)
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
		"describe-db-engine-versions": {
			Name:   "describe-db-engine-versions",
			Fields: fields_describe_db_engine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBEngineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_engine_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBEngineVersions(ctx, input)
				}
				var results []*svc.DescribeDBEngineVersionsOutput
				p := svc.NewDescribeDBEngineVersionsPaginator(client, input)
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
		"describe-db-instances": {
			Name:   "describe-db-instances",
			Fields: fields_describe_db_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBInstances(ctx, input)
				}
				var results []*svc.DescribeDBInstancesOutput
				p := svc.NewDescribeDBInstancesPaginator(client, input)
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
		"describe-db-parameter-groups": {
			Name:   "describe-db-parameter-groups",
			Fields: fields_describe_db_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBParameterGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_parameter_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBParameterGroups(ctx, input)
				}
				var results []*svc.DescribeDBParameterGroupsOutput
				p := svc.NewDescribeDBParameterGroupsPaginator(client, input)
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
		"describe-db-parameters": {
			Name:   "describe-db-parameters",
			Fields: fields_describe_db_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBParameters(ctx, input)
				}
				var results []*svc.DescribeDBParametersOutput
				p := svc.NewDescribeDBParametersPaginator(client, input)
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
		"describe-db-subnet-groups": {
			Name:   "describe-db-subnet-groups",
			Fields: fields_describe_db_subnet_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBSubnetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_subnet_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBSubnetGroups(ctx, input)
				}
				var results []*svc.DescribeDBSubnetGroupsOutput
				p := svc.NewDescribeDBSubnetGroupsPaginator(client, input)
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
		"describe-engine-default-cluster-parameters": {
			Name:   "describe-engine-default-cluster-parameters",
			Fields: fields_describe_engine_default_cluster_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEngineDefaultClusterParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_engine_default_cluster_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEngineDefaultClusterParameters(ctx, input)
			},
		},
		"describe-engine-default-parameters": {
			Name:   "describe-engine-default-parameters",
			Fields: fields_describe_engine_default_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEngineDefaultParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_engine_default_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEngineDefaultParameters(ctx, input)
				}
				var results []*svc.DescribeEngineDefaultParametersOutput
				p := svc.NewDescribeEngineDefaultParametersPaginator(client, input)
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
		"describe-global-clusters": {
			Name:   "describe-global-clusters",
			Fields: fields_describe_global_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGlobalClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_global_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGlobalClusters(ctx, input)
				}
				var results []*svc.DescribeGlobalClustersOutput
				p := svc.NewDescribeGlobalClustersPaginator(client, input)
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
		"describe-orderable-db-instance-options": {
			Name:   "describe-orderable-db-instance-options",
			Fields: fields_describe_orderable_db_instance_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrderableDBInstanceOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_orderable_db_instance_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrderableDBInstanceOptions(ctx, input)
				}
				var results []*svc.DescribeOrderableDBInstanceOptionsOutput
				p := svc.NewDescribeOrderableDBInstanceOptionsPaginator(client, input)
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
		"describe-valid-db-instance-modifications": {
			Name:   "describe-valid-db-instance-modifications",
			Fields: fields_describe_valid_db_instance_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeValidDBInstanceModificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_valid_db_instance_modifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeValidDBInstanceModifications(ctx, input)
			},
		},
		"failover-db-cluster": {
			Name:   "failover-db-cluster",
			Fields: fields_failover_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FailoverDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_failover_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FailoverDBCluster(ctx, input)
			},
		},
		"failover-global-cluster": {
			Name:   "failover-global-cluster",
			Fields: fields_failover_global_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FailoverGlobalClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_failover_global_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FailoverGlobalCluster(ctx, input)
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
		"modify-db-cluster": {
			Name:   "modify-db-cluster",
			Fields: fields_modify_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBCluster(ctx, input)
			},
		},
		"modify-db-cluster-endpoint": {
			Name:   "modify-db-cluster-endpoint",
			Fields: fields_modify_db_cluster_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBClusterEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_cluster_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBClusterEndpoint(ctx, input)
			},
		},
		"modify-db-cluster-parameter-group": {
			Name:   "modify-db-cluster-parameter-group",
			Fields: fields_modify_db_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBClusterParameterGroup(ctx, input)
			},
		},
		"modify-db-cluster-snapshot-attribute": {
			Name:   "modify-db-cluster-snapshot-attribute",
			Fields: fields_modify_db_cluster_snapshot_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBClusterSnapshotAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_cluster_snapshot_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBClusterSnapshotAttribute(ctx, input)
			},
		},
		"modify-db-instance": {
			Name:   "modify-db-instance",
			Fields: fields_modify_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBInstance(ctx, input)
			},
		},
		"modify-db-parameter-group": {
			Name:   "modify-db-parameter-group",
			Fields: fields_modify_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBParameterGroup(ctx, input)
			},
		},
		"modify-db-subnet-group": {
			Name:   "modify-db-subnet-group",
			Fields: fields_modify_db_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBSubnetGroup(ctx, input)
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
		"modify-global-cluster": {
			Name:   "modify-global-cluster",
			Fields: fields_modify_global_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyGlobalClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_global_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyGlobalCluster(ctx, input)
			},
		},
		"promote-read-replica-db-cluster": {
			Name:   "promote-read-replica-db-cluster",
			Fields: fields_promote_read_replica_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PromoteReadReplicaDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_promote_read_replica_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PromoteReadReplicaDBCluster(ctx, input)
			},
		},
		"reboot-db-instance": {
			Name:   "reboot-db-instance",
			Fields: fields_reboot_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootDBInstance(ctx, input)
			},
		},
		"remove-from-global-cluster": {
			Name:   "remove-from-global-cluster",
			Fields: fields_remove_from_global_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveFromGlobalClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_from_global_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveFromGlobalCluster(ctx, input)
			},
		},
		"remove-role-from-db-cluster": {
			Name:   "remove-role-from-db-cluster",
			Fields: fields_remove_role_from_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveRoleFromDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_role_from_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveRoleFromDBCluster(ctx, input)
			},
		},
		"remove-source-identifier-from-subscription": {
			Name:   "remove-source-identifier-from-subscription",
			Fields: fields_remove_source_identifier_from_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveSourceIdentifierFromSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_source_identifier_from_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveSourceIdentifierFromSubscription(ctx, input)
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
		"reset-db-cluster-parameter-group": {
			Name:   "reset-db-cluster-parameter-group",
			Fields: fields_reset_db_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetDBClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_db_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetDBClusterParameterGroup(ctx, input)
			},
		},
		"reset-db-parameter-group": {
			Name:   "reset-db-parameter-group",
			Fields: fields_reset_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetDBParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetDBParameterGroup(ctx, input)
			},
		},
		"restore-db-cluster-from-snapshot": {
			Name:   "restore-db-cluster-from-snapshot",
			Fields: fields_restore_db_cluster_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDBClusterFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_db_cluster_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDBClusterFromSnapshot(ctx, input)
			},
		},
		"restore-db-cluster-to-point-in-time": {
			Name:   "restore-db-cluster-to-point-in-time",
			Fields: fields_restore_db_cluster_to_point_in_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDBClusterToPointInTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_db_cluster_to_point_in_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDBClusterToPointInTime(ctx, input)
			},
		},
		"start-db-cluster": {
			Name:   "start-db-cluster",
			Fields: fields_start_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDBCluster(ctx, input)
			},
		},
		"stop-db-cluster": {
			Name:   "stop-db-cluster",
			Fields: fields_stop_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDBCluster(ctx, input)
			},
		},
		"switchover-global-cluster": {
			Name:   "switchover-global-cluster",
			Fields: fields_switchover_global_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SwitchoverGlobalClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_switchover_global_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SwitchoverGlobalCluster(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("neptune", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
