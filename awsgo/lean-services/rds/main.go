package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/rds"
)

var fields_add_role_to_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_add_role_to_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: true},
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

var fields_authorize_db_security_group_ingress = []leanruntime.Field{
	{Name: "CIDRIP", Flag: "cidrip", Type: "*string", Required: false},
	{Name: "DBSecurityGroupName", Flag: "db-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupId", Flag: "ec2-security-group-id", Type: "*string", Required: false},
	{Name: "EC2SecurityGroupName", Flag: "ec2-security-group-name", Type: "*string", Required: false},
	{Name: "EC2SecurityGroupOwnerId", Flag: "ec2-security-group-owner-id", Type: "*string", Required: false},
}

var fields_backtrack_db_cluster = []leanruntime.Field{
	{Name: "BacktrackTo", Flag: "backtrack-to", Type: "*time.Time", Required: true},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "UseEarliestTimeOnPointInTimeUnavailable", Flag: "use-earliest-time-on-point-in-time-unavailable", Type: "*bool", Required: false},
}

var fields_cancel_export_task = []leanruntime.Field{
	{Name: "ExportTaskIdentifier", Flag: "export-task-identifier", Type: "*string", Required: true},
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

var fields_copy_db_snapshot = []leanruntime.Field{
	{Name: "CopyOptionGroup", Flag: "copy-option-group", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PreSignedUrl", Flag: "pre-signed-url", Type: "*string", Required: false},
	{Name: "SnapshotAvailabilityZone", Flag: "snapshot-availability-zone", Type: "*string", Required: false},
	{Name: "SnapshotTarget", Flag: "snapshot-target", Type: "*string", Required: false},
	{Name: "SourceDBSnapshotIdentifier", Flag: "source-db-snapshot-identifier", Type: "*string", Required: true},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetCustomAvailabilityZone", Flag: "target-custom-availability-zone", Type: "*string", Required: false},
	{Name: "TargetDBSnapshotIdentifier", Flag: "target-db-snapshot-identifier", Type: "*string", Required: true},
	{Name: "destinationRegion", Flag: "destination-region", Type: "*string", Required: false},
}

var fields_copy_option_group = []leanruntime.Field{
	{Name: "SourceOptionGroupIdentifier", Flag: "source-option-group-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetOptionGroupDescription", Flag: "target-option-group-description", Type: "*string", Required: true},
	{Name: "TargetOptionGroupIdentifier", Flag: "target-option-group-identifier", Type: "*string", Required: true},
}

var fields_create_blue_green_deployment = []leanruntime.Field{
	{Name: "BlueGreenDeploymentName", Flag: "blue-green-deployment-name", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetAllocatedStorage", Flag: "target-allocated-storage", Type: "*int32", Required: false},
	{Name: "TargetDBClusterParameterGroupName", Flag: "target-db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "TargetDBInstanceClass", Flag: "target-db-instance-class", Type: "*string", Required: false},
	{Name: "TargetDBParameterGroupName", Flag: "target-db-parameter-group-name", Type: "*string", Required: false},
	{Name: "TargetEngineVersion", Flag: "target-engine-version", Type: "*string", Required: false},
	{Name: "TargetIops", Flag: "target-iops", Type: "*int32", Required: false},
	{Name: "TargetStorageThroughput", Flag: "target-storage-throughput", Type: "*int32", Required: false},
	{Name: "TargetStorageType", Flag: "target-storage-type", Type: "*string", Required: false},
	{Name: "UpgradeTargetStorageConfig", Flag: "upgrade-target-storage-config", Type: "*bool", Required: false},
}

var fields_create_custom_db_engine_version = []leanruntime.Field{
	{Name: "DatabaseInstallationFiles", Flag: "database-installation-files", Type: "[]string", Required: false},
	{Name: "DatabaseInstallationFilesS3BucketName", Flag: "database-installation-files-s3-bucket-name", Type: "*string", Required: false},
	{Name: "DatabaseInstallationFilesS3Prefix", Flag: "database-installation-files-s3-prefix", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: true},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Manifest", Flag: "manifest", Type: "*string", Required: false},
	{Name: "SourceCustomDbEngineVersionIdentifier", Flag: "source-custom-db-engine-version-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseAwsProvidedLatestImage", Flag: "use-aws-provided-latest-image", Type: "*bool", Required: false},
}

var fields_create_db_cluster = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "BacktrackWindow", Flag: "backtrack-window", Type: "*int64", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CharacterSetName", Flag: "character-set-name", Type: "*string", Required: false},
	{Name: "ClusterScalabilityType", Flag: "cluster-scalability-type", Type: "types.ClusterScalabilityType", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterInstanceClass", Flag: "db-cluster-instance-class", Type: "*string", Required: false},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DBSystemId", Flag: "db-system-id", Type: "*string", Required: false},
	{Name: "DatabaseInsightsMode", Flag: "database-insights-mode", Type: "types.DatabaseInsightsMode", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableGlobalWriteForwarding", Flag: "enable-global-write-forwarding", Type: "*bool", Required: false},
	{Name: "EnableHttpEndpoint", Flag: "enable-http-endpoint", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnableLimitlessDatabase", Flag: "enable-limitless-database", Type: "*bool", Required: false},
	{Name: "EnableLocalWriteForwarding", Flag: "enable-local-write-forwarding", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineMode", Flag: "engine-mode", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserAuthenticationType", Flag: "master-user-authentication-type", Type: "types.MasterUserAuthenticationType", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreSignedUrl", Flag: "pre-signed-url", Type: "*string", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RdsCustomClusterConfiguration", Flag: "rds-custom-cluster-configuration", Type: "*types.RdsCustomClusterConfiguration", Required: false},
	{Name: "ReplicationSourceIdentifier", Flag: "replication-source-identifier", Type: "*string", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfiguration", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: false},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
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
	{Name: "AdditionalStorageVolumes", Flag: "additional-storage-volumes", Type: "[]types.AdditionalStorageVolume", Required: false},
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "BackupTarget", Flag: "backup-target", Type: "*string", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CharacterSetName", Flag: "character-set-name", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "CustomIamInstanceProfile", Flag: "custom-iam-instance-profile", Type: "*string", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: true},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSecurityGroups", Flag: "db-security-groups", Type: "[]string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DBSystemId", Flag: "db-system-id", Type: "*string", Required: false},
	{Name: "DatabaseInsightsMode", Flag: "database-insights-mode", Type: "types.DatabaseInsightsMode", Required: false},
	{Name: "DedicatedLogVolume", Flag: "dedicated-log-volume", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainAuthSecretArn", Flag: "domain-auth-secret-arn", Type: "*string", Required: false},
	{Name: "DomainDnsIps", Flag: "domain-dns-ips", Type: "[]string", Required: false},
	{Name: "DomainFqdn", Flag: "domain-fqdn", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "DomainOu", Flag: "domain-ou", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableCustomerOwnedIp", Flag: "enable-customer-owned-ip", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserAuthenticationType", Flag: "master-user-authentication-type", Type: "types.MasterUserAuthenticationType", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: false},
	{Name: "MaxAllocatedStorage", Flag: "max-allocated-storage", Type: "*int32", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "MultiTenant", Flag: "multi-tenant", Type: "*bool", Required: false},
	{Name: "NcharCharacterSetName", Flag: "nchar-character-set-name", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ProcessorFeatures", Flag: "processor-features", Type: "[]types.ProcessorFeature", Required: false},
	{Name: "PromotionTier", Flag: "promotion-tier", Type: "*int32", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "StorageThroughput", Flag: "storage-throughput", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TdeCredentialArn", Flag: "tde-credential-arn", Type: "*string", Required: false},
	{Name: "TdeCredentialPassword", Flag: "tde-credential-password", Type: "*string", Required: false},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_create_db_instance_read_replica = []leanruntime.Field{
	{Name: "AdditionalStorageVolumes", Flag: "additional-storage-volumes", Type: "[]types.AdditionalStorageVolume", Required: false},
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "BackupTarget", Flag: "backup-target", Type: "*string", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "CustomIamInstanceProfile", Flag: "custom-iam-instance-profile", Type: "*string", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseInsightsMode", Flag: "database-insights-mode", Type: "types.DatabaseInsightsMode", Required: false},
	{Name: "DedicatedLogVolume", Flag: "dedicated-log-volume", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainAuthSecretArn", Flag: "domain-auth-secret-arn", Type: "*string", Required: false},
	{Name: "DomainDnsIps", Flag: "domain-dns-ips", Type: "[]string", Required: false},
	{Name: "DomainFqdn", Flag: "domain-fqdn", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "DomainOu", Flag: "domain-ou", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableCustomerOwnedIp", Flag: "enable-customer-owned-ip", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MaxAllocatedStorage", Flag: "max-allocated-storage", Type: "*int32", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreSignedUrl", Flag: "pre-signed-url", Type: "*string", Required: false},
	{Name: "ProcessorFeatures", Flag: "processor-features", Type: "[]types.ProcessorFeature", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "ReplicaMode", Flag: "replica-mode", Type: "types.ReplicaMode", Required: false},
	{Name: "SourceDBClusterIdentifier", Flag: "source-db-cluster-identifier", Type: "*string", Required: false},
	{Name: "SourceDBInstanceIdentifier", Flag: "source-db-instance-identifier", Type: "*string", Required: false},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: false},
	{Name: "StorageThroughput", Flag: "storage-throughput", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UpgradeStorageConfig", Flag: "upgrade-storage-config", Type: "*bool", Required: false},
	{Name: "UseDefaultProcessorFeatures", Flag: "use-default-processor-features", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
	{Name: "destinationRegion", Flag: "destination-region", Type: "*string", Required: false},
}

var fields_create_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: true},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_proxy = []leanruntime.Field{
	{Name: "Auth", Flag: "auth", Type: "[]types.UserAuthConfig", Required: false},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "DebugLogging", Flag: "debug-logging", Type: "*bool", Required: false},
	{Name: "DefaultAuthScheme", Flag: "default-auth-scheme", Type: "types.DefaultAuthScheme", Required: false},
	{Name: "EndpointNetworkType", Flag: "endpoint-network-type", Type: "types.EndpointNetworkType", Required: false},
	{Name: "EngineFamily", Flag: "engine-family", Type: "types.EngineFamily", Required: true},
	{Name: "IdleClientTimeout", Flag: "idle-client-timeout", Type: "*int32", Required: false},
	{Name: "RequireTLS", Flag: "require-tls", Type: "*bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetConnectionNetworkType", Flag: "target-connection-network-type", Type: "types.TargetConnectionNetworkType", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
	{Name: "VpcSubnetIds", Flag: "vpc-subnet-ids", Type: "[]string", Required: true},
}

var fields_create_db_proxy_endpoint = []leanruntime.Field{
	{Name: "DBProxyEndpointName", Flag: "db-proxy-endpoint-name", Type: "*string", Required: true},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "EndpointNetworkType", Flag: "endpoint-network-type", Type: "types.EndpointNetworkType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetRole", Flag: "target-role", Type: "types.DBProxyEndpointTargetRole", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
	{Name: "VpcSubnetIds", Flag: "vpc-subnet-ids", Type: "[]string", Required: true},
}

var fields_create_db_security_group = []leanruntime.Field{
	{Name: "DBSecurityGroupDescription", Flag: "db-security-group-description", Type: "*string", Required: true},
	{Name: "DBSecurityGroupName", Flag: "db-security-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_shard_group = []leanruntime.Field{
	{Name: "ComputeRedundancy", Flag: "compute-redundancy", Type: "*int32", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBShardGroupIdentifier", Flag: "db-shard-group-identifier", Type: "*string", Required: true},
	{Name: "MaxACU", Flag: "max-acu", Type: "*float64", Required: true},
	{Name: "MinACU", Flag: "min-acu", Type: "*float64", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_db_snapshot = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: true},
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
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
	{Name: "SourceDBClusterIdentifier", Flag: "source-db-cluster-identifier", Type: "*string", Required: false},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_integration = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "DataFilter", Flag: "data-filter", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: true},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_create_option_group = []leanruntime.Field{
	{Name: "EngineName", Flag: "engine-name", Type: "*string", Required: true},
	{Name: "MajorEngineVersion", Flag: "major-engine-version", Type: "*string", Required: true},
	{Name: "OptionGroupDescription", Flag: "option-group-description", Type: "*string", Required: true},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_tenant_database = []leanruntime.Field{
	{Name: "CharacterSetName", Flag: "character-set-name", Type: "*string", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: true},
	{Name: "NcharCharacterSetName", Flag: "nchar-character-set-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TenantDBName", Flag: "tenant-db-name", Type: "*string", Required: true},
}

var fields_delete_blue_green_deployment = []leanruntime.Field{
	{Name: "BlueGreenDeploymentIdentifier", Flag: "blue-green-deployment-identifier", Type: "*string", Required: true},
	{Name: "DeleteTarget", Flag: "delete-target", Type: "*bool", Required: false},
}

var fields_delete_custom_db_engine_version = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: true},
}

var fields_delete_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DeleteAutomatedBackups", Flag: "delete-automated-backups", Type: "*bool", Required: false},
	{Name: "FinalDBSnapshotIdentifier", Flag: "final-db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "SkipFinalSnapshot", Flag: "skip-final-snapshot", Type: "*bool", Required: false},
}

var fields_delete_db_cluster_automated_backup = []leanruntime.Field{
	{Name: "DbClusterResourceId", Flag: "db-cluster-resource-id", Type: "*string", Required: true},
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
	{Name: "DeleteAutomatedBackups", Flag: "delete-automated-backups", Type: "*bool", Required: false},
	{Name: "FinalDBSnapshotIdentifier", Flag: "final-db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "SkipFinalSnapshot", Flag: "skip-final-snapshot", Type: "*bool", Required: false},
}

var fields_delete_db_instance_automated_backup = []leanruntime.Field{
	{Name: "DBInstanceAutomatedBackupsArn", Flag: "db-instance-automated-backups-arn", Type: "*string", Required: false},
	{Name: "DbiResourceId", Flag: "dbi-resource-id", Type: "*string", Required: false},
}

var fields_delete_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_db_proxy = []leanruntime.Field{
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
}

var fields_delete_db_proxy_endpoint = []leanruntime.Field{
	{Name: "DBProxyEndpointName", Flag: "db-proxy-endpoint-name", Type: "*string", Required: true},
}

var fields_delete_db_security_group = []leanruntime.Field{
	{Name: "DBSecurityGroupName", Flag: "db-security-group-name", Type: "*string", Required: true},
}

var fields_delete_db_shard_group = []leanruntime.Field{
	{Name: "DBShardGroupIdentifier", Flag: "db-shard-group-identifier", Type: "*string", Required: true},
}

var fields_delete_db_snapshot = []leanruntime.Field{
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: true},
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

var fields_delete_integration = []leanruntime.Field{
	{Name: "IntegrationIdentifier", Flag: "integration-identifier", Type: "*string", Required: true},
}

var fields_delete_option_group = []leanruntime.Field{
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: true},
}

var fields_delete_tenant_database = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "FinalDBSnapshotIdentifier", Flag: "final-db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "SkipFinalSnapshot", Flag: "skip-final-snapshot", Type: "*bool", Required: false},
	{Name: "TenantDBName", Flag: "tenant-db-name", Type: "*string", Required: true},
}

var fields_deregister_db_proxy_targets = []leanruntime.Field{
	{Name: "DBClusterIdentifiers", Flag: "db-cluster-identifiers", Type: "[]string", Required: false},
	{Name: "DBInstanceIdentifiers", Flag: "db-instance-identifiers", Type: "[]string", Required: false},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "TargetGroupName", Flag: "target-group-name", Type: "*string", Required: false},
}

var fields_describe_account_attributes = []leanruntime.Field{}

var fields_describe_blue_green_deployments = []leanruntime.Field{
	{Name: "BlueGreenDeploymentIdentifier", Flag: "blue-green-deployment-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_certificates = []leanruntime.Field{
	{Name: "CertificateIdentifier", Flag: "certificate-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_cluster_automated_backups = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: false},
	{Name: "DbClusterResourceId", Flag: "db-cluster-resource-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_cluster_backtracks = []leanruntime.Field{
	{Name: "BacktrackIdentifier", Flag: "backtrack-identifier", Type: "*string", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
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
	{Name: "DbClusterResourceId", Flag: "db-cluster-resource-id", Type: "*string", Required: false},
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
	{Name: "IncludeShared", Flag: "include-shared", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_engine_versions = []leanruntime.Field{
	{Name: "DBParameterGroupFamily", Flag: "db-parameter-group-family", Type: "*string", Required: false},
	{Name: "DefaultOnly", Flag: "default-only", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeAll", Flag: "include-all", Type: "*bool", Required: false},
	{Name: "ListSupportedCharacterSets", Flag: "list-supported-character-sets", Type: "*bool", Required: false},
	{Name: "ListSupportedTimezones", Flag: "list-supported-timezones", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_instance_automated_backups = []leanruntime.Field{
	{Name: "DBInstanceAutomatedBackupsArn", Flag: "db-instance-automated-backups-arn", Type: "*string", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: false},
	{Name: "DbiResourceId", Flag: "dbi-resource-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_instances = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_log_files = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "FileLastWritten", Flag: "file-last-written", Type: "*int64", Required: false},
	{Name: "FileSize", Flag: "file-size", Type: "*int64", Required: false},
	{Name: "FilenameContains", Flag: "filename-contains", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_major_engine_versions = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "MajorEngineVersion", Flag: "major-engine-version", Type: "*string", Required: false},
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

var fields_describe_db_proxies = []leanruntime.Field{
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_proxy_endpoints = []leanruntime.Field{
	{Name: "DBProxyEndpointName", Flag: "db-proxy-endpoint-name", Type: "*string", Required: false},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_proxy_target_groups = []leanruntime.Field{
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TargetGroupName", Flag: "target-group-name", Type: "*string", Required: false},
}

var fields_describe_db_proxy_targets = []leanruntime.Field{
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TargetGroupName", Flag: "target-group-name", Type: "*string", Required: false},
}

var fields_describe_db_recommendations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LastUpdatedAfter", Flag: "last-updated-after", Type: "*time.Time", Required: false},
	{Name: "LastUpdatedBefore", Flag: "last-updated-before", Type: "*time.Time", Required: false},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_security_groups = []leanruntime.Field{
	{Name: "DBSecurityGroupName", Flag: "db-security-group-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_shard_groups = []leanruntime.Field{
	{Name: "DBShardGroupIdentifier", Flag: "db-shard-group-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_db_snapshot_attributes = []leanruntime.Field{
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: true},
}

var fields_describe_db_snapshot_tenant_databases = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: false},
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "DbiResourceId", Flag: "dbi-resource-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SnapshotType", Flag: "snapshot-type", Type: "*string", Required: false},
}

var fields_describe_db_snapshots = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: false},
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "DbiResourceId", Flag: "dbi-resource-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludePublic", Flag: "include-public", Type: "*bool", Required: false},
	{Name: "IncludeShared", Flag: "include-shared", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SnapshotType", Flag: "snapshot-type", Type: "*string", Required: false},
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

var fields_describe_export_tasks = []leanruntime.Field{
	{Name: "ExportTaskIdentifier", Flag: "export-task-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.ExportSourceType", Required: false},
}

var fields_describe_global_clusters = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_integrations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IntegrationIdentifier", Flag: "integration-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_option_group_options = []leanruntime.Field{
	{Name: "EngineName", Flag: "engine-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MajorEngineVersion", Flag: "major-engine-version", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_option_groups = []leanruntime.Field{
	{Name: "EngineName", Flag: "engine-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MajorEngineVersion", Flag: "major-engine-version", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
}

var fields_describe_orderable_db_instance_options = []leanruntime.Field{
	{Name: "AvailabilityZoneGroup", Flag: "availability-zone-group", Type: "*string", Required: false},
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

var fields_describe_reserved_db_instances = []leanruntime.Field{
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LeaseId", Flag: "lease-id", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "*string", Required: false},
	{Name: "ProductDescription", Flag: "product-description", Type: "*string", Required: false},
	{Name: "ReservedDBInstanceId", Flag: "reserved-db-instance-id", Type: "*string", Required: false},
	{Name: "ReservedDBInstancesOfferingId", Flag: "reserved-db-instances-offering-id", Type: "*string", Required: false},
}

var fields_describe_reserved_db_instances_offerings = []leanruntime.Field{
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "*string", Required: false},
	{Name: "ProductDescription", Flag: "product-description", Type: "*string", Required: false},
	{Name: "ReservedDBInstancesOfferingId", Flag: "reserved-db-instances-offering-id", Type: "*string", Required: false},
}

var fields_describe_source_regions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
}

var fields_describe_tenant_databases = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TenantDBName", Flag: "tenant-db-name", Type: "*string", Required: false},
}

var fields_describe_valid_db_instance_modifications = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
}

var fields_disable_http_endpoint = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_download_db_log_file_portion = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "LogFileName", Flag: "log-file-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "NumberOfLines", Flag: "number-of-lines", Type: "*int32", Required: false},
}

var fields_enable_http_endpoint = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_failover_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
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

var fields_modify_activity_stream = []leanruntime.Field{
	{Name: "AuditPolicyState", Flag: "audit-policy-state", Type: "types.AuditPolicyState", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_modify_certificates = []leanruntime.Field{
	{Name: "CertificateIdentifier", Flag: "certificate-identifier", Type: "*string", Required: false},
	{Name: "RemoveCustomerOverride", Flag: "remove-customer-override", Type: "*bool", Required: false},
}

var fields_modify_current_db_cluster_capacity = []leanruntime.Field{
	{Name: "Capacity", Flag: "capacity", Type: "*int32", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "SecondsBeforeTimeout", Flag: "seconds-before-timeout", Type: "*int32", Required: false},
	{Name: "TimeoutAction", Flag: "timeout-action", Type: "*string", Required: false},
}

var fields_modify_custom_db_engine_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.CustomEngineVersionStatus", Required: false},
}

var fields_modify_db_cluster = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AllowEngineModeChange", Flag: "allow-engine-mode-change", Type: "*bool", Required: false},
	{Name: "AllowMajorVersionUpgrade", Flag: "allow-major-version-upgrade", Type: "*bool", Required: false},
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AwsBackupRecoveryPointArn", Flag: "aws-backup-recovery-point-arn", Type: "*string", Required: false},
	{Name: "BacktrackWindow", Flag: "backtrack-window", Type: "*int64", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CloudwatchLogsExportConfiguration", Flag: "cloudwatch-logs-export-configuration", Type: "*types.CloudwatchLogsExportConfiguration", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterInstanceClass", Flag: "db-cluster-instance-class", Type: "*string", Required: false},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBInstanceParameterGroupName", Flag: "db-instance-parameter-group-name", Type: "*string", Required: false},
	{Name: "DatabaseInsightsMode", Flag: "database-insights-mode", Type: "types.DatabaseInsightsMode", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableGlobalWriteForwarding", Flag: "enable-global-write-forwarding", Type: "*bool", Required: false},
	{Name: "EnableHttpEndpoint", Flag: "enable-http-endpoint", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnableLimitlessDatabase", Flag: "enable-limitless-database", Type: "*bool", Required: false},
	{Name: "EnableLocalWriteForwarding", Flag: "enable-local-write-forwarding", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "EngineMode", Flag: "engine-mode", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserAuthenticationType", Flag: "master-user-authentication-type", Type: "types.MasterUserAuthenticationType", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "NewDBClusterIdentifier", Flag: "new-db-cluster-identifier", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "RotateMasterUserPassword", Flag: "rotate-master-user-password", Type: "*bool", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfiguration", Required: false},
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
	{Name: "AdditionalStorageVolumes", Flag: "additional-storage-volumes", Type: "[]types.ModifyAdditionalStorageVolume", Required: false},
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AllowMajorVersionUpgrade", Flag: "allow-major-version-upgrade", Type: "*bool", Required: false},
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AutomationMode", Flag: "automation-mode", Type: "types.AutomationMode", Required: false},
	{Name: "AwsBackupRecoveryPointArn", Flag: "aws-backup-recovery-point-arn", Type: "*string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CertificateRotationRestart", Flag: "certificate-rotation-restart", Type: "*bool", Required: false},
	{Name: "CloudwatchLogsExportConfiguration", Flag: "cloudwatch-logs-export-configuration", Type: "*types.CloudwatchLogsExportConfiguration", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBPortNumber", Flag: "db-port-number", Type: "*int32", Required: false},
	{Name: "DBSecurityGroups", Flag: "db-security-groups", Type: "[]string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseInsightsMode", Flag: "database-insights-mode", Type: "types.DatabaseInsightsMode", Required: false},
	{Name: "DedicatedLogVolume", Flag: "dedicated-log-volume", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "DisableDomain", Flag: "disable-domain", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainAuthSecretArn", Flag: "domain-auth-secret-arn", Type: "*string", Required: false},
	{Name: "DomainDnsIps", Flag: "domain-dns-ips", Type: "[]string", Required: false},
	{Name: "DomainFqdn", Flag: "domain-fqdn", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "DomainOu", Flag: "domain-ou", Type: "*string", Required: false},
	{Name: "EnableCustomerOwnedIp", Flag: "enable-customer-owned-ip", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserAuthenticationType", Flag: "master-user-authentication-type", Type: "types.MasterUserAuthenticationType", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MaxAllocatedStorage", Flag: "max-allocated-storage", Type: "*int32", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "MultiTenant", Flag: "multi-tenant", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "NewDBInstanceIdentifier", Flag: "new-db-instance-identifier", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ProcessorFeatures", Flag: "processor-features", Type: "[]types.ProcessorFeature", Required: false},
	{Name: "PromotionTier", Flag: "promotion-tier", Type: "*int32", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "ReplicaMode", Flag: "replica-mode", Type: "types.ReplicaMode", Required: false},
	{Name: "ResumeFullAutomationModeMinutes", Flag: "resume-full-automation-mode-minutes", Type: "*int32", Required: false},
	{Name: "RotateMasterUserPassword", Flag: "rotate-master-user-password", Type: "*bool", Required: false},
	{Name: "StorageThroughput", Flag: "storage-throughput", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TdeCredentialArn", Flag: "tde-credential-arn", Type: "*string", Required: false},
	{Name: "TdeCredentialPassword", Flag: "tde-credential-password", Type: "*string", Required: false},
	{Name: "UseDefaultProcessorFeatures", Flag: "use-default-processor-features", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_db_parameter_group = []leanruntime.Field{
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: true},
}

var fields_modify_db_proxy = []leanruntime.Field{
	{Name: "Auth", Flag: "auth", Type: "[]types.UserAuthConfig", Required: false},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "DebugLogging", Flag: "debug-logging", Type: "*bool", Required: false},
	{Name: "DefaultAuthScheme", Flag: "default-auth-scheme", Type: "types.DefaultAuthScheme", Required: false},
	{Name: "IdleClientTimeout", Flag: "idle-client-timeout", Type: "*int32", Required: false},
	{Name: "NewDBProxyName", Flag: "new-db-proxy-name", Type: "*string", Required: false},
	{Name: "RequireTLS", Flag: "require-tls", Type: "*bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
}

var fields_modify_db_proxy_endpoint = []leanruntime.Field{
	{Name: "DBProxyEndpointName", Flag: "db-proxy-endpoint-name", Type: "*string", Required: true},
	{Name: "NewDBProxyEndpointName", Flag: "new-db-proxy-endpoint-name", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_db_proxy_target_group = []leanruntime.Field{
	{Name: "ConnectionPoolConfig", Flag: "connection-pool-config", Type: "*types.ConnectionPoolConfiguration", Required: false},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "NewName", Flag: "new-name", Type: "*string", Required: false},
	{Name: "TargetGroupName", Flag: "target-group-name", Type: "*string", Required: true},
}

var fields_modify_db_recommendation = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "RecommendationId", Flag: "recommendation-id", Type: "*string", Required: true},
	{Name: "RecommendedActionUpdates", Flag: "recommended-action-updates", Type: "[]types.RecommendedActionUpdate", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_modify_db_shard_group = []leanruntime.Field{
	{Name: "ComputeRedundancy", Flag: "compute-redundancy", Type: "*int32", Required: false},
	{Name: "DBShardGroupIdentifier", Flag: "db-shard-group-identifier", Type: "*string", Required: true},
	{Name: "MaxACU", Flag: "max-acu", Type: "*float64", Required: false},
	{Name: "MinACU", Flag: "min-acu", Type: "*float64", Required: false},
}

var fields_modify_db_snapshot = []leanruntime.Field{
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
}

var fields_modify_db_snapshot_attribute = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: true},
	{Name: "ValuesToAdd", Flag: "values-to-add", Type: "[]string", Required: false},
	{Name: "ValuesToRemove", Flag: "values-to-remove", Type: "[]string", Required: false},
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

var fields_modify_integration = []leanruntime.Field{
	{Name: "DataFilter", Flag: "data-filter", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationIdentifier", Flag: "integration-identifier", Type: "*string", Required: true},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: false},
}

var fields_modify_option_group = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: true},
	{Name: "OptionsToInclude", Flag: "options-to-include", Type: "[]types.OptionConfiguration", Required: false},
	{Name: "OptionsToRemove", Flag: "options-to-remove", Type: "[]string", Required: false},
}

var fields_modify_tenant_database = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "NewTenantDBName", Flag: "new-tenant-db-name", Type: "*string", Required: false},
	{Name: "RotateMasterUserPassword", Flag: "rotate-master-user-password", Type: "*bool", Required: false},
	{Name: "TenantDBName", Flag: "tenant-db-name", Type: "*string", Required: true},
}

var fields_promote_read_replica = []leanruntime.Field{
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_promote_read_replica_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_purchase_reserved_db_instances_offering = []leanruntime.Field{
	{Name: "DBInstanceCount", Flag: "db-instance-count", Type: "*int32", Required: false},
	{Name: "ReservedDBInstanceId", Flag: "reserved-db-instance-id", Type: "*string", Required: false},
	{Name: "ReservedDBInstancesOfferingId", Flag: "reserved-db-instances-offering-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_reboot_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_reboot_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "ForceFailover", Flag: "force-failover", Type: "*bool", Required: false},
}

var fields_reboot_db_shard_group = []leanruntime.Field{
	{Name: "DBShardGroupIdentifier", Flag: "db-shard-group-identifier", Type: "*string", Required: true},
}

var fields_register_db_proxy_targets = []leanruntime.Field{
	{Name: "DBClusterIdentifiers", Flag: "db-cluster-identifiers", Type: "[]string", Required: false},
	{Name: "DBInstanceIdentifiers", Flag: "db-instance-identifiers", Type: "[]string", Required: false},
	{Name: "DBProxyName", Flag: "db-proxy-name", Type: "*string", Required: true},
	{Name: "TargetGroupName", Flag: "target-group-name", Type: "*string", Required: false},
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

var fields_remove_role_from_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: true},
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

var fields_restore_db_cluster_from_s3 = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "BacktrackWindow", Flag: "backtrack-window", Type: "*int64", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CharacterSetName", Flag: "character-set-name", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "S3IngestionRoleArn", Flag: "s3-ingestion-role-arn", Type: "*string", Required: true},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SourceEngine", Flag: "source-engine", Type: "*string", Required: true},
	{Name: "SourceEngineVersion", Flag: "source-engine-version", Type: "*string", Required: true},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_db_cluster_from_snapshot = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "BacktrackWindow", Flag: "backtrack-window", Type: "*int64", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterInstanceClass", Flag: "db-cluster-instance-class", Type: "*string", Required: false},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineMode", Flag: "engine-mode", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RdsCustomClusterConfiguration", Flag: "rds-custom-cluster-configuration", Type: "*types.RdsCustomClusterConfiguration", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfiguration", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_db_cluster_to_point_in_time = []leanruntime.Field{
	{Name: "BacktrackWindow", Flag: "backtrack-window", Type: "*int64", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
	{Name: "DBClusterInstanceClass", Flag: "db-cluster-instance-class", Type: "*string", Required: false},
	{Name: "DBClusterParameterGroupName", Flag: "db-cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineMode", Flag: "engine-mode", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RdsCustomClusterConfiguration", Flag: "rds-custom-cluster-configuration", Type: "*types.RdsCustomClusterConfiguration", Required: false},
	{Name: "RestoreToTime", Flag: "restore-to-time", Type: "*time.Time", Required: false},
	{Name: "RestoreType", Flag: "restore-type", Type: "*string", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfiguration", Required: false},
	{Name: "ServerlessV2ScalingConfiguration", Flag: "serverless-v2-scaling-configuration", Type: "*types.ServerlessV2ScalingConfiguration", Required: false},
	{Name: "SourceDBClusterIdentifier", Flag: "source-db-cluster-identifier", Type: "*string", Required: false},
	{Name: "SourceDbClusterResourceId", Flag: "source-db-cluster-resource-id", Type: "*string", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseLatestRestorableTime", Flag: "use-latest-restorable-time", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_db_instance_from_db_snapshot = []leanruntime.Field{
	{Name: "AdditionalStorageVolumes", Flag: "additional-storage-volumes", Type: "[]types.AdditionalStorageVolume", Required: false},
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "BackupTarget", Flag: "backup-target", Type: "*string", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "CustomIamInstanceProfile", Flag: "custom-iam-instance-profile", Type: "*string", Required: false},
	{Name: "DBClusterSnapshotIdentifier", Flag: "db-cluster-snapshot-identifier", Type: "*string", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DedicatedLogVolume", Flag: "dedicated-log-volume", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainAuthSecretArn", Flag: "domain-auth-secret-arn", Type: "*string", Required: false},
	{Name: "DomainDnsIps", Flag: "domain-dns-ips", Type: "[]string", Required: false},
	{Name: "DomainFqdn", Flag: "domain-fqdn", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "DomainOu", Flag: "domain-ou", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableCustomerOwnedIp", Flag: "enable-customer-owned-ip", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "ProcessorFeatures", Flag: "processor-features", Type: "[]types.ProcessorFeature", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "StorageThroughput", Flag: "storage-throughput", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TdeCredentialArn", Flag: "tde-credential-arn", Type: "*string", Required: false},
	{Name: "TdeCredentialPassword", Flag: "tde-credential-password", Type: "*string", Required: false},
	{Name: "UseDefaultProcessorFeatures", Flag: "use-default-processor-features", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_db_instance_from_s3 = []leanruntime.Field{
	{Name: "AdditionalStorageVolumes", Flag: "additional-storage-volumes", Type: "[]types.AdditionalStorageVolume", Required: false},
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: true},
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSecurityGroups", Flag: "db-security-groups", Type: "[]string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DatabaseInsightsMode", Flag: "database-insights-mode", Type: "types.DatabaseInsightsMode", Required: false},
	{Name: "DedicatedLogVolume", Flag: "dedicated-log-volume", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "EnablePerformanceInsights", Flag: "enable-performance-insights", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: false},
	{Name: "MaxAllocatedStorage", Flag: "max-allocated-storage", Type: "*int32", Required: false},
	{Name: "MonitoringInterval", Flag: "monitoring-interval", Type: "*int32", Required: false},
	{Name: "MonitoringRoleArn", Flag: "monitoring-role-arn", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "PerformanceInsightsKMSKeyId", Flag: "performance-insights-kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceInsightsRetentionPeriod", Flag: "performance-insights-retention-period", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ProcessorFeatures", Flag: "processor-features", Type: "[]types.ProcessorFeature", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "S3IngestionRoleArn", Flag: "s3-ingestion-role-arn", Type: "*string", Required: true},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: false},
	{Name: "SourceEngine", Flag: "source-engine", Type: "*string", Required: true},
	{Name: "SourceEngineVersion", Flag: "source-engine-version", Type: "*string", Required: true},
	{Name: "StorageEncrypted", Flag: "storage-encrypted", Type: "*bool", Required: false},
	{Name: "StorageThroughput", Flag: "storage-throughput", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseDefaultProcessorFeatures", Flag: "use-default-processor-features", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_db_instance_to_point_in_time = []leanruntime.Field{
	{Name: "AdditionalStorageVolumes", Flag: "additional-storage-volumes", Type: "[]types.AdditionalStorageVolume", Required: false},
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "BackupTarget", Flag: "backup-target", Type: "*string", Required: false},
	{Name: "CACertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "CopyTagsToSnapshot", Flag: "copy-tags-to-snapshot", Type: "*bool", Required: false},
	{Name: "CustomIamInstanceProfile", Flag: "custom-iam-instance-profile", Type: "*string", Required: false},
	{Name: "DBInstanceClass", Flag: "db-instance-class", Type: "*string", Required: false},
	{Name: "DBName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DBParameterGroupName", Flag: "db-parameter-group-name", Type: "*string", Required: false},
	{Name: "DBSubnetGroupName", Flag: "db-subnet-group-name", Type: "*string", Required: false},
	{Name: "DedicatedLogVolume", Flag: "dedicated-log-volume", Type: "*bool", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "DomainAuthSecretArn", Flag: "domain-auth-secret-arn", Type: "*string", Required: false},
	{Name: "DomainDnsIps", Flag: "domain-dns-ips", Type: "[]string", Required: false},
	{Name: "DomainFqdn", Flag: "domain-fqdn", Type: "*string", Required: false},
	{Name: "DomainIAMRoleName", Flag: "domain-iam-role-name", Type: "*string", Required: false},
	{Name: "DomainOu", Flag: "domain-ou", Type: "*string", Required: false},
	{Name: "EnableCloudwatchLogsExports", Flag: "enable-cloudwatch-logs-exports", Type: "[]string", Required: false},
	{Name: "EnableCustomerOwnedIp", Flag: "enable-customer-owned-ip", Type: "*bool", Required: false},
	{Name: "EnableIAMDatabaseAuthentication", Flag: "enable-iam-database-authentication", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineLifecycleSupport", Flag: "engine-lifecycle-support", Type: "*string", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "LicenseModel", Flag: "license-model", Type: "*string", Required: false},
	{Name: "ManageMasterUserPassword", Flag: "manage-master-user-password", Type: "*bool", Required: false},
	{Name: "MasterUserSecretKmsKeyId", Flag: "master-user-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MaxAllocatedStorage", Flag: "max-allocated-storage", Type: "*int32", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "*string", Required: false},
	{Name: "OptionGroupName", Flag: "option-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "ProcessorFeatures", Flag: "processor-features", Type: "[]types.ProcessorFeature", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RestoreTime", Flag: "restore-time", Type: "*time.Time", Required: false},
	{Name: "SourceDBInstanceAutomatedBackupsArn", Flag: "source-db-instance-automated-backups-arn", Type: "*string", Required: false},
	{Name: "SourceDBInstanceIdentifier", Flag: "source-db-instance-identifier", Type: "*string", Required: false},
	{Name: "SourceDbiResourceId", Flag: "source-dbi-resource-id", Type: "*string", Required: false},
	{Name: "StorageThroughput", Flag: "storage-throughput", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDBInstanceIdentifier", Flag: "target-db-instance-identifier", Type: "*string", Required: true},
	{Name: "TdeCredentialArn", Flag: "tde-credential-arn", Type: "*string", Required: false},
	{Name: "TdeCredentialPassword", Flag: "tde-credential-password", Type: "*string", Required: false},
	{Name: "UseDefaultProcessorFeatures", Flag: "use-default-processor-features", Type: "*bool", Required: false},
	{Name: "UseLatestRestorableTime", Flag: "use-latest-restorable-time", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_revoke_db_security_group_ingress = []leanruntime.Field{
	{Name: "CIDRIP", Flag: "cidrip", Type: "*string", Required: false},
	{Name: "DBSecurityGroupName", Flag: "db-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupId", Flag: "ec2-security-group-id", Type: "*string", Required: false},
	{Name: "EC2SecurityGroupName", Flag: "ec2-security-group-name", Type: "*string", Required: false},
	{Name: "EC2SecurityGroupOwnerId", Flag: "ec2-security-group-owner-id", Type: "*string", Required: false},
}

var fields_start_activity_stream = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "EngineNativeAuditFieldsIncluded", Flag: "engine-native-audit-fields-included", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.ActivityStreamMode", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_start_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
}

var fields_start_db_instance_automated_backups_replication = []leanruntime.Field{
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PreSignedUrl", Flag: "pre-signed-url", Type: "*string", Required: false},
	{Name: "SourceDBInstanceArn", Flag: "source-db-instance-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_export_task = []leanruntime.Field{
	{Name: "ExportOnly", Flag: "export-only", Type: "[]string", Required: false},
	{Name: "ExportTaskIdentifier", Flag: "export-task-identifier", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
}

var fields_stop_activity_stream = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_stop_db_cluster = []leanruntime.Field{
	{Name: "DBClusterIdentifier", Flag: "db-cluster-identifier", Type: "*string", Required: true},
}

var fields_stop_db_instance = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
	{Name: "DBSnapshotIdentifier", Flag: "db-snapshot-identifier", Type: "*string", Required: false},
}

var fields_stop_db_instance_automated_backups_replication = []leanruntime.Field{
	{Name: "SourceDBInstanceArn", Flag: "source-db-instance-arn", Type: "*string", Required: true},
}

var fields_switchover_blue_green_deployment = []leanruntime.Field{
	{Name: "BlueGreenDeploymentIdentifier", Flag: "blue-green-deployment-identifier", Type: "*string", Required: true},
	{Name: "SwitchoverTimeout", Flag: "switchover-timeout", Type: "*int32", Required: false},
}

var fields_switchover_global_cluster = []leanruntime.Field{
	{Name: "GlobalClusterIdentifier", Flag: "global-cluster-identifier", Type: "*string", Required: true},
	{Name: "TargetDbClusterIdentifier", Flag: "target-db-cluster-identifier", Type: "*string", Required: true},
}

var fields_switchover_read_replica = []leanruntime.Field{
	{Name: "DBInstanceIdentifier", Flag: "db-instance-identifier", Type: "*string", Required: true},
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
		"add-role-to-db-instance": {
			Name:   "add-role-to-db-instance",
			Fields: fields_add_role_to_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddRoleToDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_role_to_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddRoleToDBInstance(ctx, input)
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
		"authorize-db-security-group-ingress": {
			Name:   "authorize-db-security-group-ingress",
			Fields: fields_authorize_db_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeDBSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_db_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeDBSecurityGroupIngress(ctx, input)
			},
		},
		"backtrack-db-cluster": {
			Name:   "backtrack-db-cluster",
			Fields: fields_backtrack_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BacktrackDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_backtrack_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BacktrackDBCluster(ctx, input)
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
		"copy-db-snapshot": {
			Name:   "copy-db-snapshot",
			Fields: fields_copy_db_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyDBSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_db_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyDBSnapshot(ctx, input)
			},
		},
		"copy-option-group": {
			Name:   "copy-option-group",
			Fields: fields_copy_option_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyOptionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_option_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyOptionGroup(ctx, input)
			},
		},
		"create-blue-green-deployment": {
			Name:   "create-blue-green-deployment",
			Fields: fields_create_blue_green_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBlueGreenDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_blue_green_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBlueGreenDeployment(ctx, input)
			},
		},
		"create-custom-db-engine-version": {
			Name:   "create-custom-db-engine-version",
			Fields: fields_create_custom_db_engine_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomDBEngineVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_db_engine_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomDBEngineVersion(ctx, input)
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
		"create-db-instance-read-replica": {
			Name:   "create-db-instance-read-replica",
			Fields: fields_create_db_instance_read_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBInstanceReadReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_instance_read_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBInstanceReadReplica(ctx, input)
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
		"create-db-proxy": {
			Name:   "create-db-proxy",
			Fields: fields_create_db_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBProxy(ctx, input)
			},
		},
		"create-db-proxy-endpoint": {
			Name:   "create-db-proxy-endpoint",
			Fields: fields_create_db_proxy_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBProxyEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_proxy_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBProxyEndpoint(ctx, input)
			},
		},
		"create-db-security-group": {
			Name:   "create-db-security-group",
			Fields: fields_create_db_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBSecurityGroup(ctx, input)
			},
		},
		"create-db-shard-group": {
			Name:   "create-db-shard-group",
			Fields: fields_create_db_shard_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBShardGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_shard_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBShardGroup(ctx, input)
			},
		},
		"create-db-snapshot": {
			Name:   "create-db-snapshot",
			Fields: fields_create_db_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDBSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDBSnapshot(ctx, input)
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
		"create-option-group": {
			Name:   "create-option-group",
			Fields: fields_create_option_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOptionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_option_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOptionGroup(ctx, input)
			},
		},
		"create-tenant-database": {
			Name:   "create-tenant-database",
			Fields: fields_create_tenant_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTenantDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tenant_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTenantDatabase(ctx, input)
			},
		},
		"delete-blue-green-deployment": {
			Name:   "delete-blue-green-deployment",
			Fields: fields_delete_blue_green_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBlueGreenDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_blue_green_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBlueGreenDeployment(ctx, input)
			},
		},
		"delete-custom-db-engine-version": {
			Name:   "delete-custom-db-engine-version",
			Fields: fields_delete_custom_db_engine_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomDBEngineVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_db_engine_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomDBEngineVersion(ctx, input)
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
		"delete-db-cluster-automated-backup": {
			Name:   "delete-db-cluster-automated-backup",
			Fields: fields_delete_db_cluster_automated_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBClusterAutomatedBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_cluster_automated_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBClusterAutomatedBackup(ctx, input)
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
		"delete-db-instance-automated-backup": {
			Name:   "delete-db-instance-automated-backup",
			Fields: fields_delete_db_instance_automated_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBInstanceAutomatedBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_instance_automated_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBInstanceAutomatedBackup(ctx, input)
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
		"delete-db-proxy": {
			Name:   "delete-db-proxy",
			Fields: fields_delete_db_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBProxy(ctx, input)
			},
		},
		"delete-db-proxy-endpoint": {
			Name:   "delete-db-proxy-endpoint",
			Fields: fields_delete_db_proxy_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBProxyEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_proxy_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBProxyEndpoint(ctx, input)
			},
		},
		"delete-db-security-group": {
			Name:   "delete-db-security-group",
			Fields: fields_delete_db_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBSecurityGroup(ctx, input)
			},
		},
		"delete-db-shard-group": {
			Name:   "delete-db-shard-group",
			Fields: fields_delete_db_shard_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBShardGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_shard_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBShardGroup(ctx, input)
			},
		},
		"delete-db-snapshot": {
			Name:   "delete-db-snapshot",
			Fields: fields_delete_db_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDBSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDBSnapshot(ctx, input)
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
		"delete-option-group": {
			Name:   "delete-option-group",
			Fields: fields_delete_option_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOptionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_option_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOptionGroup(ctx, input)
			},
		},
		"delete-tenant-database": {
			Name:   "delete-tenant-database",
			Fields: fields_delete_tenant_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTenantDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tenant_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTenantDatabase(ctx, input)
			},
		},
		"deregister-db-proxy-targets": {
			Name:   "deregister-db-proxy-targets",
			Fields: fields_deregister_db_proxy_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterDBProxyTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_db_proxy_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterDBProxyTargets(ctx, input)
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
		"describe-blue-green-deployments": {
			Name:   "describe-blue-green-deployments",
			Fields: fields_describe_blue_green_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBlueGreenDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_blue_green_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBlueGreenDeployments(ctx, input)
				}
				var results []*svc.DescribeBlueGreenDeploymentsOutput
				p := svc.NewDescribeBlueGreenDeploymentsPaginator(client, input)
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
		"describe-db-cluster-automated-backups": {
			Name:   "describe-db-cluster-automated-backups",
			Fields: fields_describe_db_cluster_automated_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterAutomatedBackupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_automated_backups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusterAutomatedBackups(ctx, input)
				}
				var results []*svc.DescribeDBClusterAutomatedBackupsOutput
				p := svc.NewDescribeDBClusterAutomatedBackupsPaginator(client, input)
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
		"describe-db-cluster-backtracks": {
			Name:   "describe-db-cluster-backtracks",
			Fields: fields_describe_db_cluster_backtracks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBClusterBacktracksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_cluster_backtracks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBClusterBacktracks(ctx, input)
				}
				var results []*svc.DescribeDBClusterBacktracksOutput
				p := svc.NewDescribeDBClusterBacktracksPaginator(client, input)
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
		"describe-db-instance-automated-backups": {
			Name:   "describe-db-instance-automated-backups",
			Fields: fields_describe_db_instance_automated_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBInstanceAutomatedBackupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_instance_automated_backups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBInstanceAutomatedBackups(ctx, input)
				}
				var results []*svc.DescribeDBInstanceAutomatedBackupsOutput
				p := svc.NewDescribeDBInstanceAutomatedBackupsPaginator(client, input)
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
		"describe-db-log-files": {
			Name:   "describe-db-log-files",
			Fields: fields_describe_db_log_files,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBLogFilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_log_files, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBLogFiles(ctx, input)
				}
				var results []*svc.DescribeDBLogFilesOutput
				p := svc.NewDescribeDBLogFilesPaginator(client, input)
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
		"describe-db-major-engine-versions": {
			Name:   "describe-db-major-engine-versions",
			Fields: fields_describe_db_major_engine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBMajorEngineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_major_engine_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBMajorEngineVersions(ctx, input)
				}
				var results []*svc.DescribeDBMajorEngineVersionsOutput
				p := svc.NewDescribeDBMajorEngineVersionsPaginator(client, input)
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
		"describe-db-proxies": {
			Name:   "describe-db-proxies",
			Fields: fields_describe_db_proxies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBProxiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_proxies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBProxies(ctx, input)
				}
				var results []*svc.DescribeDBProxiesOutput
				p := svc.NewDescribeDBProxiesPaginator(client, input)
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
		"describe-db-proxy-endpoints": {
			Name:   "describe-db-proxy-endpoints",
			Fields: fields_describe_db_proxy_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBProxyEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_proxy_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBProxyEndpoints(ctx, input)
				}
				var results []*svc.DescribeDBProxyEndpointsOutput
				p := svc.NewDescribeDBProxyEndpointsPaginator(client, input)
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
		"describe-db-proxy-target-groups": {
			Name:   "describe-db-proxy-target-groups",
			Fields: fields_describe_db_proxy_target_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBProxyTargetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_proxy_target_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBProxyTargetGroups(ctx, input)
				}
				var results []*svc.DescribeDBProxyTargetGroupsOutput
				p := svc.NewDescribeDBProxyTargetGroupsPaginator(client, input)
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
		"describe-db-proxy-targets": {
			Name:   "describe-db-proxy-targets",
			Fields: fields_describe_db_proxy_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBProxyTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_proxy_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBProxyTargets(ctx, input)
				}
				var results []*svc.DescribeDBProxyTargetsOutput
				p := svc.NewDescribeDBProxyTargetsPaginator(client, input)
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
		"describe-db-recommendations": {
			Name:   "describe-db-recommendations",
			Fields: fields_describe_db_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBRecommendations(ctx, input)
				}
				var results []*svc.DescribeDBRecommendationsOutput
				p := svc.NewDescribeDBRecommendationsPaginator(client, input)
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
		"describe-db-security-groups": {
			Name:   "describe-db-security-groups",
			Fields: fields_describe_db_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBSecurityGroups(ctx, input)
				}
				var results []*svc.DescribeDBSecurityGroupsOutput
				p := svc.NewDescribeDBSecurityGroupsPaginator(client, input)
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
		"describe-db-shard-groups": {
			Name:   "describe-db-shard-groups",
			Fields: fields_describe_db_shard_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBShardGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_db_shard_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDBShardGroups(ctx, input)
			},
		},
		"describe-db-snapshot-attributes": {
			Name:   "describe-db-snapshot-attributes",
			Fields: fields_describe_db_snapshot_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBSnapshotAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_db_snapshot_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDBSnapshotAttributes(ctx, input)
			},
		},
		"describe-db-snapshot-tenant-databases": {
			Name:   "describe-db-snapshot-tenant-databases",
			Fields: fields_describe_db_snapshot_tenant_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBSnapshotTenantDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_snapshot_tenant_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBSnapshotTenantDatabases(ctx, input)
				}
				var results []*svc.DescribeDBSnapshotTenantDatabasesOutput
				p := svc.NewDescribeDBSnapshotTenantDatabasesPaginator(client, input)
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
		"describe-db-snapshots": {
			Name:   "describe-db-snapshots",
			Fields: fields_describe_db_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDBSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_db_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDBSnapshots(ctx, input)
				}
				var results []*svc.DescribeDBSnapshotsOutput
				p := svc.NewDescribeDBSnapshotsPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_engine_default_cluster_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEngineDefaultClusterParameters(ctx, input)
				}
				var results []*svc.DescribeEngineDefaultClusterParametersOutput
				p := svc.NewDescribeEngineDefaultClusterParametersPaginator(client, input)
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
		"describe-export-tasks": {
			Name:   "describe-export-tasks",
			Fields: fields_describe_export_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_export_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeExportTasks(ctx, input)
				}
				var results []*svc.DescribeExportTasksOutput
				p := svc.NewDescribeExportTasksPaginator(client, input)
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
		"describe-integrations": {
			Name:   "describe-integrations",
			Fields: fields_describe_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIntegrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_integrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIntegrations(ctx, input)
				}
				var results []*svc.DescribeIntegrationsOutput
				p := svc.NewDescribeIntegrationsPaginator(client, input)
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
		"describe-option-group-options": {
			Name:   "describe-option-group-options",
			Fields: fields_describe_option_group_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOptionGroupOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_option_group_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOptionGroupOptions(ctx, input)
				}
				var results []*svc.DescribeOptionGroupOptionsOutput
				p := svc.NewDescribeOptionGroupOptionsPaginator(client, input)
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
		"describe-option-groups": {
			Name:   "describe-option-groups",
			Fields: fields_describe_option_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOptionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_option_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOptionGroups(ctx, input)
				}
				var results []*svc.DescribeOptionGroupsOutput
				p := svc.NewDescribeOptionGroupsPaginator(client, input)
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
		"describe-reserved-db-instances": {
			Name:   "describe-reserved-db-instances",
			Fields: fields_describe_reserved_db_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedDBInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_db_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedDBInstances(ctx, input)
				}
				var results []*svc.DescribeReservedDBInstancesOutput
				p := svc.NewDescribeReservedDBInstancesPaginator(client, input)
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
		"describe-reserved-db-instances-offerings": {
			Name:   "describe-reserved-db-instances-offerings",
			Fields: fields_describe_reserved_db_instances_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedDBInstancesOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_db_instances_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedDBInstancesOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedDBInstancesOfferingsOutput
				p := svc.NewDescribeReservedDBInstancesOfferingsPaginator(client, input)
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
		"describe-source-regions": {
			Name:   "describe-source-regions",
			Fields: fields_describe_source_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSourceRegionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_source_regions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSourceRegions(ctx, input)
				}
				var results []*svc.DescribeSourceRegionsOutput
				p := svc.NewDescribeSourceRegionsPaginator(client, input)
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
		"describe-tenant-databases": {
			Name:   "describe-tenant-databases",
			Fields: fields_describe_tenant_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTenantDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tenant_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTenantDatabases(ctx, input)
				}
				var results []*svc.DescribeTenantDatabasesOutput
				p := svc.NewDescribeTenantDatabasesPaginator(client, input)
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
		"disable-http-endpoint": {
			Name:   "disable-http-endpoint",
			Fields: fields_disable_http_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableHttpEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_http_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableHttpEndpoint(ctx, input)
			},
		},
		"download-db-log-file-portion": {
			Name:   "download-db-log-file-portion",
			Fields: fields_download_db_log_file_portion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DownloadDBLogFilePortionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_download_db_log_file_portion, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DownloadDBLogFilePortion(ctx, input)
				}
				var results []*svc.DownloadDBLogFilePortionOutput
				p := svc.NewDownloadDBLogFilePortionPaginator(client, input)
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
		"enable-http-endpoint": {
			Name:   "enable-http-endpoint",
			Fields: fields_enable_http_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableHttpEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_http_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableHttpEndpoint(ctx, input)
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
		"modify-activity-stream": {
			Name:   "modify-activity-stream",
			Fields: fields_modify_activity_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyActivityStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_activity_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyActivityStream(ctx, input)
			},
		},
		"modify-certificates": {
			Name:   "modify-certificates",
			Fields: fields_modify_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCertificatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_certificates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCertificates(ctx, input)
			},
		},
		"modify-current-db-cluster-capacity": {
			Name:   "modify-current-db-cluster-capacity",
			Fields: fields_modify_current_db_cluster_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCurrentDBClusterCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_current_db_cluster_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCurrentDBClusterCapacity(ctx, input)
			},
		},
		"modify-custom-db-engine-version": {
			Name:   "modify-custom-db-engine-version",
			Fields: fields_modify_custom_db_engine_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCustomDBEngineVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_custom_db_engine_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCustomDBEngineVersion(ctx, input)
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
		"modify-db-proxy": {
			Name:   "modify-db-proxy",
			Fields: fields_modify_db_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBProxy(ctx, input)
			},
		},
		"modify-db-proxy-endpoint": {
			Name:   "modify-db-proxy-endpoint",
			Fields: fields_modify_db_proxy_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBProxyEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_proxy_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBProxyEndpoint(ctx, input)
			},
		},
		"modify-db-proxy-target-group": {
			Name:   "modify-db-proxy-target-group",
			Fields: fields_modify_db_proxy_target_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBProxyTargetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_proxy_target_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBProxyTargetGroup(ctx, input)
			},
		},
		"modify-db-recommendation": {
			Name:   "modify-db-recommendation",
			Fields: fields_modify_db_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBRecommendation(ctx, input)
			},
		},
		"modify-db-shard-group": {
			Name:   "modify-db-shard-group",
			Fields: fields_modify_db_shard_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBShardGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_shard_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBShardGroup(ctx, input)
			},
		},
		"modify-db-snapshot": {
			Name:   "modify-db-snapshot",
			Fields: fields_modify_db_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBSnapshot(ctx, input)
			},
		},
		"modify-db-snapshot-attribute": {
			Name:   "modify-db-snapshot-attribute",
			Fields: fields_modify_db_snapshot_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDBSnapshotAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_db_snapshot_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDBSnapshotAttribute(ctx, input)
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
		"modify-option-group": {
			Name:   "modify-option-group",
			Fields: fields_modify_option_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyOptionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_option_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyOptionGroup(ctx, input)
			},
		},
		"modify-tenant-database": {
			Name:   "modify-tenant-database",
			Fields: fields_modify_tenant_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTenantDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_tenant_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTenantDatabase(ctx, input)
			},
		},
		"promote-read-replica": {
			Name:   "promote-read-replica",
			Fields: fields_promote_read_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PromoteReadReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_promote_read_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PromoteReadReplica(ctx, input)
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
		"purchase-reserved-db-instances-offering": {
			Name:   "purchase-reserved-db-instances-offering",
			Fields: fields_purchase_reserved_db_instances_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedDBInstancesOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_db_instances_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedDBInstancesOffering(ctx, input)
			},
		},
		"reboot-db-cluster": {
			Name:   "reboot-db-cluster",
			Fields: fields_reboot_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootDBClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootDBCluster(ctx, input)
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
		"reboot-db-shard-group": {
			Name:   "reboot-db-shard-group",
			Fields: fields_reboot_db_shard_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootDBShardGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_db_shard_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootDBShardGroup(ctx, input)
			},
		},
		"register-db-proxy-targets": {
			Name:   "register-db-proxy-targets",
			Fields: fields_register_db_proxy_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDBProxyTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_db_proxy_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDBProxyTargets(ctx, input)
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
		"remove-role-from-db-instance": {
			Name:   "remove-role-from-db-instance",
			Fields: fields_remove_role_from_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveRoleFromDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_role_from_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveRoleFromDBInstance(ctx, input)
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
		"restore-db-cluster-from-s3": {
			Name:   "restore-db-cluster-from-s3",
			Fields: fields_restore_db_cluster_from_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDBClusterFromS3Input{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_db_cluster_from_s3, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDBClusterFromS3(ctx, input)
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
		"restore-db-instance-from-db-snapshot": {
			Name:   "restore-db-instance-from-db-snapshot",
			Fields: fields_restore_db_instance_from_db_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDBInstanceFromDBSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_db_instance_from_db_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDBInstanceFromDBSnapshot(ctx, input)
			},
		},
		"restore-db-instance-from-s3": {
			Name:   "restore-db-instance-from-s3",
			Fields: fields_restore_db_instance_from_s3,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDBInstanceFromS3Input{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_db_instance_from_s3, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDBInstanceFromS3(ctx, input)
			},
		},
		"restore-db-instance-to-point-in-time": {
			Name:   "restore-db-instance-to-point-in-time",
			Fields: fields_restore_db_instance_to_point_in_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDBInstanceToPointInTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_db_instance_to_point_in_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDBInstanceToPointInTime(ctx, input)
			},
		},
		"revoke-db-security-group-ingress": {
			Name:   "revoke-db-security-group-ingress",
			Fields: fields_revoke_db_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeDBSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_db_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeDBSecurityGroupIngress(ctx, input)
			},
		},
		"start-activity-stream": {
			Name:   "start-activity-stream",
			Fields: fields_start_activity_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartActivityStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_activity_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartActivityStream(ctx, input)
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
		"start-db-instance": {
			Name:   "start-db-instance",
			Fields: fields_start_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDBInstance(ctx, input)
			},
		},
		"start-db-instance-automated-backups-replication": {
			Name:   "start-db-instance-automated-backups-replication",
			Fields: fields_start_db_instance_automated_backups_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDBInstanceAutomatedBackupsReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_db_instance_automated_backups_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDBInstanceAutomatedBackupsReplication(ctx, input)
			},
		},
		"start-export-task": {
			Name:   "start-export-task",
			Fields: fields_start_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExportTask(ctx, input)
			},
		},
		"stop-activity-stream": {
			Name:   "stop-activity-stream",
			Fields: fields_stop_activity_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopActivityStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_activity_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopActivityStream(ctx, input)
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
		"stop-db-instance": {
			Name:   "stop-db-instance",
			Fields: fields_stop_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDBInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDBInstance(ctx, input)
			},
		},
		"stop-db-instance-automated-backups-replication": {
			Name:   "stop-db-instance-automated-backups-replication",
			Fields: fields_stop_db_instance_automated_backups_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDBInstanceAutomatedBackupsReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_db_instance_automated_backups_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDBInstanceAutomatedBackupsReplication(ctx, input)
			},
		},
		"switchover-blue-green-deployment": {
			Name:   "switchover-blue-green-deployment",
			Fields: fields_switchover_blue_green_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SwitchoverBlueGreenDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_switchover_blue_green_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SwitchoverBlueGreenDeployment(ctx, input)
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
		"switchover-read-replica": {
			Name:   "switchover-read-replica",
			Fields: fields_switchover_read_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SwitchoverReadReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_switchover_read_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SwitchoverReadReplica(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("rds", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
