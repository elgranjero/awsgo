package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/redshift"
)

var fields_accept_reserved_node_exchange = []leanruntime.Field{
	{Name: "ReservedNodeId", Flag: "reserved-node-id", Type: "*string", Required: true},
	{Name: "TargetReservedNodeOfferingId", Flag: "target-reserved-node-offering-id", Type: "*string", Required: true},
}

var fields_add_partner = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartnerName", Flag: "partner-name", Type: "*string", Required: true},
}

var fields_associate_data_share_consumer = []leanruntime.Field{
	{Name: "AllowWrites", Flag: "allow-writes", Type: "*bool", Required: false},
	{Name: "AssociateEntireAccount", Flag: "associate-entire-account", Type: "*bool", Required: false},
	{Name: "ConsumerArn", Flag: "consumer-arn", Type: "*string", Required: false},
	{Name: "ConsumerRegion", Flag: "consumer-region", Type: "*string", Required: false},
	{Name: "DataShareArn", Flag: "data-share-arn", Type: "*string", Required: true},
}

var fields_authorize_cluster_security_group_ingress = []leanruntime.Field{
	{Name: "CIDRIP", Flag: "cidrip", Type: "*string", Required: false},
	{Name: "ClusterSecurityGroupName", Flag: "cluster-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupName", Flag: "ec2-security-group-name", Type: "*string", Required: false},
	{Name: "EC2SecurityGroupOwnerId", Flag: "ec2-security-group-owner-id", Type: "*string", Required: false},
}

var fields_authorize_data_share = []leanruntime.Field{
	{Name: "AllowWrites", Flag: "allow-writes", Type: "*bool", Required: false},
	{Name: "ConsumerIdentifier", Flag: "consumer-identifier", Type: "*string", Required: true},
	{Name: "DataShareArn", Flag: "data-share-arn", Type: "*string", Required: true},
}

var fields_authorize_endpoint_access = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_authorize_snapshot_access = []leanruntime.Field{
	{Name: "AccountWithRestoreAccess", Flag: "account-with-restore-access", Type: "*string", Required: true},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotClusterIdentifier", Flag: "snapshot-cluster-identifier", Type: "*string", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: false},
}

var fields_batch_delete_cluster_snapshots = []leanruntime.Field{
	{Name: "Identifiers", Flag: "identifiers", Type: "[]types.DeleteClusterSnapshotMessage", Required: true},
}

var fields_batch_modify_cluster_snapshots = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotIdentifierList", Flag: "snapshot-identifier-list", Type: "[]string", Required: true},
}

var fields_cancel_resize = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_copy_cluster_snapshot = []leanruntime.Field{
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "SourceSnapshotClusterIdentifier", Flag: "source-snapshot-cluster-identifier", Type: "*string", Required: false},
	{Name: "SourceSnapshotIdentifier", Flag: "source-snapshot-identifier", Type: "*string", Required: true},
	{Name: "TargetSnapshotIdentifier", Flag: "target-snapshot-identifier", Type: "*string", Required: true},
}

var fields_create_authentication_profile = []leanruntime.Field{
	{Name: "AuthenticationProfileContent", Flag: "authentication-profile-content", Type: "*string", Required: true},
	{Name: "AuthenticationProfileName", Flag: "authentication-profile-name", Type: "*string", Required: true},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "*string", Required: false},
	{Name: "AllowVersionUpgrade", Flag: "allow-version-upgrade", Type: "*bool", Required: false},
	{Name: "AquaConfigurationStatus", Flag: "aqua-configuration-status", Type: "types.AquaConfigurationStatus", Required: false},
	{Name: "AutomatedSnapshotRetentionPeriod", Flag: "automated-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneRelocation", Flag: "availability-zone-relocation", Type: "*bool", Required: false},
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ClusterParameterGroupName", Flag: "cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "ClusterSecurityGroups", Flag: "cluster-security-groups", Type: "[]string", Required: false},
	{Name: "ClusterSubnetGroupName", Flag: "cluster-subnet-group-name", Type: "*string", Required: false},
	{Name: "ClusterType", Flag: "cluster-type", Type: "*string", Required: false},
	{Name: "ClusterVersion", Flag: "cluster-version", Type: "*string", Required: false},
	{Name: "DBName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DefaultIamRoleArn", Flag: "default-iam-role-arn", Type: "*string", Required: false},
	{Name: "ElasticIp", Flag: "elastic-ip", Type: "*string", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "EnhancedVpcRouting", Flag: "enhanced-vpc-routing", Type: "*bool", Required: false},
	{Name: "ExtraComputeForAutomaticOptimization", Flag: "extra-compute-for-automatic-optimization", Type: "*bool", Required: false},
	{Name: "HsmClientCertificateIdentifier", Flag: "hsm-client-certificate-identifier", Type: "*string", Required: false},
	{Name: "HsmConfigurationIdentifier", Flag: "hsm-configuration-identifier", Type: "*string", Required: false},
	{Name: "IamRoles", Flag: "iam-roles", Type: "[]string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LoadSampleData", Flag: "load-sample-data", Type: "*string", Required: false},
	{Name: "MaintenanceTrackName", Flag: "maintenance-track-name", Type: "*string", Required: false},
	{Name: "ManageMasterPassword", Flag: "manage-master-password", Type: "*bool", Required: false},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "MasterPasswordSecretKmsKeyId", Flag: "master-password-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: true},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: true},
	{Name: "NumberOfNodes", Flag: "number-of-nodes", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RedshiftIdcApplicationArn", Flag: "redshift-idc-application-arn", Type: "*string", Required: false},
	{Name: "SnapshotScheduleIdentifier", Flag: "snapshot-schedule-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_create_cluster_parameter_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "ParameterGroupFamily", Flag: "parameter-group-family", Type: "*string", Required: true},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cluster_security_group = []leanruntime.Field{
	{Name: "ClusterSecurityGroupName", Flag: "cluster-security-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cluster_snapshot = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cluster_subnet_group = []leanruntime.Field{
	{Name: "ClusterSubnetGroupName", Flag: "cluster-subnet-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_domain_association = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "CustomDomainCertificateArn", Flag: "custom-domain-certificate-arn", Type: "*string", Required: true},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
}

var fields_create_endpoint_access = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "*string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: true},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_create_event_subscription = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventCategories", Flag: "event-categories", Type: "[]string", Required: false},
	{Name: "Severity", Flag: "severity", Type: "*string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: true},
	{Name: "SourceIds", Flag: "source-ids", Type: "[]string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_hsm_client_certificate = []leanruntime.Field{
	{Name: "HsmClientCertificateIdentifier", Flag: "hsm-client-certificate-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_hsm_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "HsmConfigurationIdentifier", Flag: "hsm-configuration-identifier", Type: "*string", Required: true},
	{Name: "HsmIpAddress", Flag: "hsm-ip-address", Type: "*string", Required: true},
	{Name: "HsmPartitionName", Flag: "hsm-partition-name", Type: "*string", Required: true},
	{Name: "HsmPartitionPassword", Flag: "hsm-partition-password", Type: "*string", Required: true},
	{Name: "HsmServerPublicCertificate", Flag: "hsm-server-public-certificate", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_integration = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: true},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_create_redshift_idc_application = []leanruntime.Field{
	{Name: "ApplicationType", Flag: "application-type", Type: "types.ApplicationType", Required: false},
	{Name: "AuthorizedTokenIssuerList", Flag: "authorized-token-issuer-list", Type: "[]types.AuthorizedTokenIssuer", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "IdcDisplayName", Flag: "idc-display-name", Type: "*string", Required: true},
	{Name: "IdcInstanceArn", Flag: "idc-instance-arn", Type: "*string", Required: true},
	{Name: "IdentityNamespace", Flag: "identity-namespace", Type: "*string", Required: false},
	{Name: "RedshiftIdcApplicationName", Flag: "redshift-idc-application-name", Type: "*string", Required: true},
	{Name: "ServiceIntegrations", Flag: "service-integrations", Type: "[]types.ServiceIntegrationsUnion", Required: false},
	{Name: "SsoTagKeys", Flag: "sso-tag-keys", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_scheduled_action = []leanruntime.Field{
	{Name: "Enable", Flag: "enable", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "IamRole", Flag: "iam-role", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: true},
	{Name: "ScheduledActionDescription", Flag: "scheduled-action-description", Type: "*string", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TargetAction", Flag: "target-action", Type: "*types.ScheduledActionType", Required: true},
}

var fields_create_snapshot_copy_grant = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SnapshotCopyGrantName", Flag: "snapshot-copy-grant-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_snapshot_schedule = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NextInvocations", Flag: "next-invocations", Type: "*int32", Required: false},
	{Name: "ScheduleDefinitions", Flag: "schedule-definitions", Type: "[]string", Required: false},
	{Name: "ScheduleDescription", Flag: "schedule-description", Type: "*string", Required: false},
	{Name: "ScheduleIdentifier", Flag: "schedule-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_create_usage_limit = []leanruntime.Field{
	{Name: "Amount", Flag: "amount", Type: "*int64", Required: true},
	{Name: "BreachAction", Flag: "breach-action", Type: "types.UsageLimitBreachAction", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "FeatureType", Flag: "feature-type", Type: "types.UsageLimitFeatureType", Required: true},
	{Name: "LimitType", Flag: "limit-type", Type: "types.UsageLimitLimitType", Required: true},
	{Name: "Period", Flag: "period", Type: "types.UsageLimitPeriod", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_deauthorize_data_share = []leanruntime.Field{
	{Name: "ConsumerIdentifier", Flag: "consumer-identifier", Type: "*string", Required: true},
	{Name: "DataShareArn", Flag: "data-share-arn", Type: "*string", Required: true},
}

var fields_delete_authentication_profile = []leanruntime.Field{
	{Name: "AuthenticationProfileName", Flag: "authentication-profile-name", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "FinalClusterSnapshotIdentifier", Flag: "final-cluster-snapshot-identifier", Type: "*string", Required: false},
	{Name: "FinalClusterSnapshotRetentionPeriod", Flag: "final-cluster-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "SkipFinalClusterSnapshot", Flag: "skip-final-cluster-snapshot", Type: "*bool", Required: false},
}

var fields_delete_cluster_parameter_group = []leanruntime.Field{
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_cluster_security_group = []leanruntime.Field{
	{Name: "ClusterSecurityGroupName", Flag: "cluster-security-group-name", Type: "*string", Required: true},
}

var fields_delete_cluster_snapshot = []leanruntime.Field{
	{Name: "SnapshotClusterIdentifier", Flag: "snapshot-cluster-identifier", Type: "*string", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
}

var fields_delete_cluster_subnet_group = []leanruntime.Field{
	{Name: "ClusterSubnetGroupName", Flag: "cluster-subnet-group-name", Type: "*string", Required: true},
}

var fields_delete_custom_domain_association = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
}

var fields_delete_endpoint_access = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_delete_event_subscription = []leanruntime.Field{
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_delete_hsm_client_certificate = []leanruntime.Field{
	{Name: "HsmClientCertificateIdentifier", Flag: "hsm-client-certificate-identifier", Type: "*string", Required: true},
}

var fields_delete_hsm_configuration = []leanruntime.Field{
	{Name: "HsmConfigurationIdentifier", Flag: "hsm-configuration-identifier", Type: "*string", Required: true},
}

var fields_delete_integration = []leanruntime.Field{
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
}

var fields_delete_partner = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartnerName", Flag: "partner-name", Type: "*string", Required: true},
}

var fields_delete_redshift_idc_application = []leanruntime.Field{
	{Name: "RedshiftIdcApplicationArn", Flag: "redshift-idc-application-arn", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_scheduled_action = []leanruntime.Field{
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
}

var fields_delete_snapshot_copy_grant = []leanruntime.Field{
	{Name: "SnapshotCopyGrantName", Flag: "snapshot-copy-grant-name", Type: "*string", Required: true},
}

var fields_delete_snapshot_schedule = []leanruntime.Field{
	{Name: "ScheduleIdentifier", Flag: "schedule-identifier", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_delete_usage_limit = []leanruntime.Field{
	{Name: "UsageLimitId", Flag: "usage-limit-id", Type: "*string", Required: true},
}

var fields_deregister_namespace = []leanruntime.Field{
	{Name: "ConsumerIdentifiers", Flag: "consumer-identifiers", Type: "[]string", Required: true},
	{Name: "NamespaceIdentifier", Flag: "namespace-identifier", Type: "types.NamespaceIdentifierUnion", Required: true},
}

var fields_describe_account_attributes = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]string", Required: false},
}

var fields_describe_authentication_profiles = []leanruntime.Field{
	{Name: "AuthenticationProfileName", Flag: "authentication-profile-name", Type: "*string", Required: false},
}

var fields_describe_cluster_db_revisions = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_cluster_parameter_groups = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_cluster_parameters = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_cluster_security_groups = []leanruntime.Field{
	{Name: "ClusterSecurityGroupName", Flag: "cluster-security-group-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_cluster_snapshots = []leanruntime.Field{
	{Name: "ClusterExists", Flag: "cluster-exists", Type: "*bool", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: false},
	{Name: "SnapshotType", Flag: "snapshot-type", Type: "*string", Required: false},
	{Name: "SortingEntities", Flag: "sorting-entities", Type: "[]types.SnapshotSortingEntity", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_cluster_subnet_groups = []leanruntime.Field{
	{Name: "ClusterSubnetGroupName", Flag: "cluster-subnet-group-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_cluster_tracks = []leanruntime.Field{
	{Name: "MaintenanceTrackName", Flag: "maintenance-track-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_cluster_versions = []leanruntime.Field{
	{Name: "ClusterParameterGroupFamily", Flag: "cluster-parameter-group-family", Type: "*string", Required: false},
	{Name: "ClusterVersion", Flag: "cluster-version", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_clusters = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_custom_domain_associations = []leanruntime.Field{
	{Name: "CustomDomainCertificateArn", Flag: "custom-domain-certificate-arn", Type: "*string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_data_shares = []leanruntime.Field{
	{Name: "DataShareArn", Flag: "data-share-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_data_shares_for_consumer = []leanruntime.Field{
	{Name: "ConsumerArn", Flag: "consumer-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "Status", Flag: "status", Type: "types.DataShareStatusForConsumer", Required: false},
}

var fields_describe_data_shares_for_producer = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ProducerArn", Flag: "producer-arn", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.DataShareStatusForProducer", Required: false},
}

var fields_describe_default_cluster_parameters = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ParameterGroupFamily", Flag: "parameter-group-family", Type: "*string", Required: true},
}

var fields_describe_endpoint_access = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_describe_endpoint_authorization = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Grantee", Flag: "grantee", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_event_categories = []leanruntime.Field{
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
}

var fields_describe_event_subscriptions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_events = []leanruntime.Field{
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SourceIdentifier", Flag: "source-identifier", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.SourceType", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_hsm_client_certificates = []leanruntime.Field{
	{Name: "HsmClientCertificateIdentifier", Flag: "hsm-client-certificate-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_hsm_configurations = []leanruntime.Field{
	{Name: "HsmConfigurationIdentifier", Flag: "hsm-configuration-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_inbound_integrations = []leanruntime.Field{
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_describe_integrations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DescribeIntegrationsFilter", Required: false},
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_logging_status = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_describe_node_configuration_options = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.ActionType", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.NodeConfigurationOptionsFilter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: false},
}

var fields_describe_orderable_cluster_options = []leanruntime.Field{
	{Name: "ClusterVersion", Flag: "cluster-version", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
}

var fields_describe_partners = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "PartnerName", Flag: "partner-name", Type: "*string", Required: false},
}

var fields_describe_redshift_idc_applications = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "RedshiftIdcApplicationArn", Flag: "redshift-idc-application-arn", Type: "*string", Required: false},
}

var fields_describe_reserved_node_exchange_status = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReservedNodeExchangeRequestId", Flag: "reserved-node-exchange-request-id", Type: "*string", Required: false},
	{Name: "ReservedNodeId", Flag: "reserved-node-id", Type: "*string", Required: false},
}

var fields_describe_reserved_node_offerings = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReservedNodeOfferingId", Flag: "reserved-node-offering-id", Type: "*string", Required: false},
}

var fields_describe_reserved_nodes = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReservedNodeId", Flag: "reserved-node-id", Type: "*string", Required: false},
}

var fields_describe_resize = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_describe_scheduled_actions = []leanruntime.Field{
	{Name: "Active", Flag: "active", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ScheduledActionFilter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TargetActionType", Flag: "target-action-type", Type: "types.ScheduledActionTypeValues", Required: false},
}

var fields_describe_snapshot_copy_grants = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SnapshotCopyGrantName", Flag: "snapshot-copy-grant-name", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_snapshot_schedules = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ScheduleIdentifier", Flag: "schedule-identifier", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_storage = []leanruntime.Field{}

var fields_describe_table_restore_status = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TableRestoreRequestId", Flag: "table-restore-request-id", Type: "*string", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
}

var fields_describe_usage_limits = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "FeatureType", Flag: "feature-type", Type: "types.UsageLimitFeatureType", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: false},
	{Name: "UsageLimitId", Flag: "usage-limit-id", Type: "*string", Required: false},
}

var fields_disable_logging = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_disable_snapshot_copy = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_disassociate_data_share_consumer = []leanruntime.Field{
	{Name: "ConsumerArn", Flag: "consumer-arn", Type: "*string", Required: false},
	{Name: "ConsumerRegion", Flag: "consumer-region", Type: "*string", Required: false},
	{Name: "DataShareArn", Flag: "data-share-arn", Type: "*string", Required: true},
	{Name: "DisassociateEntireAccount", Flag: "disassociate-entire-account", Type: "*bool", Required: false},
}

var fields_enable_logging = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "LogDestinationType", Flag: "log-destination-type", Type: "types.LogDestinationType", Required: false},
	{Name: "LogExports", Flag: "log-exports", Type: "[]string", Required: false},
	{Name: "S3KeyPrefix", Flag: "s3-key-prefix", Type: "*string", Required: false},
}

var fields_enable_snapshot_copy = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DestinationRegion", Flag: "destination-region", Type: "*string", Required: true},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotCopyGrantName", Flag: "snapshot-copy-grant-name", Type: "*string", Required: false},
}

var fields_failover_primary_compute = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_get_cluster_credentials = []leanruntime.Field{
	{Name: "AutoCreate", Flag: "auto-create", Type: "*bool", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "DbGroups", Flag: "db-groups", Type: "[]string", Required: false},
	{Name: "DbName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DbUser", Flag: "db-user", Type: "*string", Required: true},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
}

var fields_get_cluster_credentials_with_iam = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "DbName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
}

var fields_get_identity_center_auth_token = []leanruntime.Field{
	{Name: "ClusterIds", Flag: "cluster-ids", Type: "[]string", Required: true},
}

var fields_get_reserved_node_exchange_configuration_options = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.ReservedNodeExchangeActionType", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: false},
}

var fields_get_reserved_node_exchange_offerings = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReservedNodeId", Flag: "reserved-node-id", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_recommendations = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NamespaceArn", Flag: "namespace-arn", Type: "*string", Required: false},
}

var fields_modify_aqua_configuration = []leanruntime.Field{
	{Name: "AquaConfigurationStatus", Flag: "aqua-configuration-status", Type: "types.AquaConfigurationStatus", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_modify_authentication_profile = []leanruntime.Field{
	{Name: "AuthenticationProfileContent", Flag: "authentication-profile-content", Type: "*string", Required: true},
	{Name: "AuthenticationProfileName", Flag: "authentication-profile-name", Type: "*string", Required: true},
}

var fields_modify_cluster = []leanruntime.Field{
	{Name: "AllowVersionUpgrade", Flag: "allow-version-upgrade", Type: "*bool", Required: false},
	{Name: "AutomatedSnapshotRetentionPeriod", Flag: "automated-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneRelocation", Flag: "availability-zone-relocation", Type: "*bool", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ClusterParameterGroupName", Flag: "cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "ClusterSecurityGroups", Flag: "cluster-security-groups", Type: "[]string", Required: false},
	{Name: "ClusterType", Flag: "cluster-type", Type: "*string", Required: false},
	{Name: "ClusterVersion", Flag: "cluster-version", Type: "*string", Required: false},
	{Name: "ElasticIp", Flag: "elastic-ip", Type: "*string", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "EnhancedVpcRouting", Flag: "enhanced-vpc-routing", Type: "*bool", Required: false},
	{Name: "ExtraComputeForAutomaticOptimization", Flag: "extra-compute-for-automatic-optimization", Type: "*bool", Required: false},
	{Name: "HsmClientCertificateIdentifier", Flag: "hsm-client-certificate-identifier", Type: "*string", Required: false},
	{Name: "HsmConfigurationIdentifier", Flag: "hsm-configuration-identifier", Type: "*string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MaintenanceTrackName", Flag: "maintenance-track-name", Type: "*string", Required: false},
	{Name: "ManageMasterPassword", Flag: "manage-master-password", Type: "*bool", Required: false},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "MasterPasswordSecretKmsKeyId", Flag: "master-password-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NewClusterIdentifier", Flag: "new-cluster-identifier", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "NumberOfNodes", Flag: "number-of-nodes", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_cluster_db_revision = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "RevisionTarget", Flag: "revision-target", Type: "*string", Required: true},
}

var fields_modify_cluster_iam_roles = []leanruntime.Field{
	{Name: "AddIamRoles", Flag: "add-iam-roles", Type: "[]string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DefaultIamRoleArn", Flag: "default-iam-role-arn", Type: "*string", Required: false},
	{Name: "RemoveIamRoles", Flag: "remove-iam-roles", Type: "[]string", Required: false},
}

var fields_modify_cluster_maintenance = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DeferMaintenance", Flag: "defer-maintenance", Type: "*bool", Required: false},
	{Name: "DeferMaintenanceDuration", Flag: "defer-maintenance-duration", Type: "*int32", Required: false},
	{Name: "DeferMaintenanceEndTime", Flag: "defer-maintenance-end-time", Type: "*time.Time", Required: false},
	{Name: "DeferMaintenanceIdentifier", Flag: "defer-maintenance-identifier", Type: "*string", Required: false},
	{Name: "DeferMaintenanceStartTime", Flag: "defer-maintenance-start-time", Type: "*time.Time", Required: false},
}

var fields_modify_cluster_parameter_group = []leanruntime.Field{
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: true},
}

var fields_modify_cluster_snapshot = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
}

var fields_modify_cluster_snapshot_schedule = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DisassociateSchedule", Flag: "disassociate-schedule", Type: "*bool", Required: false},
	{Name: "ScheduleIdentifier", Flag: "schedule-identifier", Type: "*string", Required: false},
}

var fields_modify_cluster_subnet_group = []leanruntime.Field{
	{Name: "ClusterSubnetGroupName", Flag: "cluster-subnet-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
}

var fields_modify_custom_domain_association = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "CustomDomainCertificateArn", Flag: "custom-domain-certificate-arn", Type: "*string", Required: true},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
}

var fields_modify_endpoint_access = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_event_subscription = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventCategories", Flag: "event-categories", Type: "[]string", Required: false},
	{Name: "Severity", Flag: "severity", Type: "*string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SourceIds", Flag: "source-ids", Type: "[]string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: true},
}

var fields_modify_integration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
	{Name: "IntegrationName", Flag: "integration-name", Type: "*string", Required: false},
}

var fields_modify_lakehouse_configuration = []leanruntime.Field{
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LakehouseIdcApplicationArn", Flag: "lakehouse-idc-application-arn", Type: "*string", Required: false},
	{Name: "LakehouseIdcRegistration", Flag: "lakehouse-idc-registration", Type: "types.LakehouseIdcRegistration", Required: false},
	{Name: "LakehouseRegistration", Flag: "lakehouse-registration", Type: "types.LakehouseRegistration", Required: false},
}

var fields_modify_redshift_idc_application = []leanruntime.Field{
	{Name: "AuthorizedTokenIssuerList", Flag: "authorized-token-issuer-list", Type: "[]types.AuthorizedTokenIssuer", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "IdcDisplayName", Flag: "idc-display-name", Type: "*string", Required: false},
	{Name: "IdentityNamespace", Flag: "identity-namespace", Type: "*string", Required: false},
	{Name: "RedshiftIdcApplicationArn", Flag: "redshift-idc-application-arn", Type: "*string", Required: true},
	{Name: "ServiceIntegrations", Flag: "service-integrations", Type: "[]types.ServiceIntegrationsUnion", Required: false},
}

var fields_modify_scheduled_action = []leanruntime.Field{
	{Name: "Enable", Flag: "enable", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "IamRole", Flag: "iam-role", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "ScheduledActionDescription", Flag: "scheduled-action-description", Type: "*string", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TargetAction", Flag: "target-action", Type: "*types.ScheduledActionType", Required: false},
}

var fields_modify_snapshot_copy_retention_period = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "Manual", Flag: "manual", Type: "*bool", Required: false},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: true},
}

var fields_modify_snapshot_schedule = []leanruntime.Field{
	{Name: "ScheduleDefinitions", Flag: "schedule-definitions", Type: "[]string", Required: true},
	{Name: "ScheduleIdentifier", Flag: "schedule-identifier", Type: "*string", Required: true},
}

var fields_modify_usage_limit = []leanruntime.Field{
	{Name: "Amount", Flag: "amount", Type: "*int64", Required: false},
	{Name: "BreachAction", Flag: "breach-action", Type: "types.UsageLimitBreachAction", Required: false},
	{Name: "UsageLimitId", Flag: "usage-limit-id", Type: "*string", Required: true},
}

var fields_pause_cluster = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_purchase_reserved_node_offering = []leanruntime.Field{
	{Name: "NodeCount", Flag: "node-count", Type: "*int32", Required: false},
	{Name: "ReservedNodeOfferingId", Flag: "reserved-node-offering-id", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reboot_cluster = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_register_namespace = []leanruntime.Field{
	{Name: "ConsumerIdentifiers", Flag: "consumer-identifiers", Type: "[]string", Required: true},
	{Name: "NamespaceIdentifier", Flag: "namespace-identifier", Type: "types.NamespaceIdentifierUnion", Required: true},
}

var fields_reject_data_share = []leanruntime.Field{
	{Name: "DataShareArn", Flag: "data-share-arn", Type: "*string", Required: true},
}

var fields_reset_cluster_parameter_group = []leanruntime.Field{
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "ResetAllParameters", Flag: "reset-all-parameters", Type: "*bool", Required: false},
}

var fields_resize_cluster = []leanruntime.Field{
	{Name: "Classic", Flag: "classic", Type: "*bool", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ClusterType", Flag: "cluster-type", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "NumberOfNodes", Flag: "number-of-nodes", Type: "*int32", Required: false},
	{Name: "ReservedNodeId", Flag: "reserved-node-id", Type: "*string", Required: false},
	{Name: "TargetReservedNodeOfferingId", Flag: "target-reserved-node-offering-id", Type: "*string", Required: false},
}

var fields_restore_from_cluster_snapshot = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "*string", Required: false},
	{Name: "AllowVersionUpgrade", Flag: "allow-version-upgrade", Type: "*bool", Required: false},
	{Name: "AquaConfigurationStatus", Flag: "aqua-configuration-status", Type: "types.AquaConfigurationStatus", Required: false},
	{Name: "AutomatedSnapshotRetentionPeriod", Flag: "automated-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneRelocation", Flag: "availability-zone-relocation", Type: "*bool", Required: false},
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ClusterParameterGroupName", Flag: "cluster-parameter-group-name", Type: "*string", Required: false},
	{Name: "ClusterSecurityGroups", Flag: "cluster-security-groups", Type: "[]string", Required: false},
	{Name: "ClusterSubnetGroupName", Flag: "cluster-subnet-group-name", Type: "*string", Required: false},
	{Name: "DefaultIamRoleArn", Flag: "default-iam-role-arn", Type: "*string", Required: false},
	{Name: "ElasticIp", Flag: "elastic-ip", Type: "*string", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "EnhancedVpcRouting", Flag: "enhanced-vpc-routing", Type: "*bool", Required: false},
	{Name: "HsmClientCertificateIdentifier", Flag: "hsm-client-certificate-identifier", Type: "*string", Required: false},
	{Name: "HsmConfigurationIdentifier", Flag: "hsm-configuration-identifier", Type: "*string", Required: false},
	{Name: "IamRoles", Flag: "iam-roles", Type: "[]string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MaintenanceTrackName", Flag: "maintenance-track-name", Type: "*string", Required: false},
	{Name: "ManageMasterPassword", Flag: "manage-master-password", Type: "*bool", Required: false},
	{Name: "ManualSnapshotRetentionPeriod", Flag: "manual-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "MasterPasswordSecretKmsKeyId", Flag: "master-password-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "NumberOfNodes", Flag: "number-of-nodes", Type: "*int32", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RedshiftIdcApplicationArn", Flag: "redshift-idc-application-arn", Type: "*string", Required: false},
	{Name: "ReservedNodeId", Flag: "reserved-node-id", Type: "*string", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotClusterIdentifier", Flag: "snapshot-cluster-identifier", Type: "*string", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: false},
	{Name: "SnapshotScheduleIdentifier", Flag: "snapshot-schedule-identifier", Type: "*string", Required: false},
	{Name: "TargetReservedNodeOfferingId", Flag: "target-reserved-node-offering-id", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_restore_table_from_cluster_snapshot = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "EnableCaseSensitiveIdentifier", Flag: "enable-case-sensitive-identifier", Type: "*bool", Required: false},
	{Name: "NewTableName", Flag: "new-table-name", Type: "*string", Required: true},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
	{Name: "SourceDatabaseName", Flag: "source-database-name", Type: "*string", Required: true},
	{Name: "SourceSchemaName", Flag: "source-schema-name", Type: "*string", Required: false},
	{Name: "SourceTableName", Flag: "source-table-name", Type: "*string", Required: true},
	{Name: "TargetDatabaseName", Flag: "target-database-name", Type: "*string", Required: false},
	{Name: "TargetSchemaName", Flag: "target-schema-name", Type: "*string", Required: false},
}

var fields_resume_cluster = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_revoke_cluster_security_group_ingress = []leanruntime.Field{
	{Name: "CIDRIP", Flag: "cidrip", Type: "*string", Required: false},
	{Name: "ClusterSecurityGroupName", Flag: "cluster-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupName", Flag: "ec2-security-group-name", Type: "*string", Required: false},
	{Name: "EC2SecurityGroupOwnerId", Flag: "ec2-security-group-owner-id", Type: "*string", Required: false},
}

var fields_revoke_endpoint_access = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_revoke_snapshot_access = []leanruntime.Field{
	{Name: "AccountWithRestoreAccess", Flag: "account-with-restore-access", Type: "*string", Required: true},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotClusterIdentifier", Flag: "snapshot-cluster-identifier", Type: "*string", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: false},
}

var fields_rotate_encryption_key = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_update_partner_status = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "PartnerName", Flag: "partner-name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.PartnerIntegrationStatus", Required: true},
	{Name: "StatusMessage", Flag: "status-message", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-reserved-node-exchange": {
			Name:   "accept-reserved-node-exchange",
			Fields: fields_accept_reserved_node_exchange,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptReservedNodeExchangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_reserved_node_exchange, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptReservedNodeExchange(ctx, input)
			},
		},
		"add-partner": {
			Name:   "add-partner",
			Fields: fields_add_partner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddPartnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_partner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddPartner(ctx, input)
			},
		},
		"associate-data-share-consumer": {
			Name:   "associate-data-share-consumer",
			Fields: fields_associate_data_share_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDataShareConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_data_share_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDataShareConsumer(ctx, input)
			},
		},
		"authorize-cluster-security-group-ingress": {
			Name:   "authorize-cluster-security-group-ingress",
			Fields: fields_authorize_cluster_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeClusterSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_cluster_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeClusterSecurityGroupIngress(ctx, input)
			},
		},
		"authorize-data-share": {
			Name:   "authorize-data-share",
			Fields: fields_authorize_data_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeDataShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_data_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeDataShare(ctx, input)
			},
		},
		"authorize-endpoint-access": {
			Name:   "authorize-endpoint-access",
			Fields: fields_authorize_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeEndpointAccess(ctx, input)
			},
		},
		"authorize-snapshot-access": {
			Name:   "authorize-snapshot-access",
			Fields: fields_authorize_snapshot_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeSnapshotAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_snapshot_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeSnapshotAccess(ctx, input)
			},
		},
		"batch-delete-cluster-snapshots": {
			Name:   "batch-delete-cluster-snapshots",
			Fields: fields_batch_delete_cluster_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteClusterSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_cluster_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteClusterSnapshots(ctx, input)
			},
		},
		"batch-modify-cluster-snapshots": {
			Name:   "batch-modify-cluster-snapshots",
			Fields: fields_batch_modify_cluster_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchModifyClusterSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_modify_cluster_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchModifyClusterSnapshots(ctx, input)
			},
		},
		"cancel-resize": {
			Name:   "cancel-resize",
			Fields: fields_cancel_resize,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelResizeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_resize, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelResize(ctx, input)
			},
		},
		"copy-cluster-snapshot": {
			Name:   "copy-cluster-snapshot",
			Fields: fields_copy_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyClusterSnapshot(ctx, input)
			},
		},
		"create-authentication-profile": {
			Name:   "create-authentication-profile",
			Fields: fields_create_authentication_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAuthenticationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_authentication_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAuthenticationProfile(ctx, input)
			},
		},
		"create-cluster": {
			Name:   "create-cluster",
			Fields: fields_create_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCluster(ctx, input)
			},
		},
		"create-cluster-parameter-group": {
			Name:   "create-cluster-parameter-group",
			Fields: fields_create_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClusterParameterGroup(ctx, input)
			},
		},
		"create-cluster-security-group": {
			Name:   "create-cluster-security-group",
			Fields: fields_create_cluster_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClusterSecurityGroup(ctx, input)
			},
		},
		"create-cluster-snapshot": {
			Name:   "create-cluster-snapshot",
			Fields: fields_create_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClusterSnapshot(ctx, input)
			},
		},
		"create-cluster-subnet-group": {
			Name:   "create-cluster-subnet-group",
			Fields: fields_create_cluster_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClusterSubnetGroup(ctx, input)
			},
		},
		"create-custom-domain-association": {
			Name:   "create-custom-domain-association",
			Fields: fields_create_custom_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomDomainAssociation(ctx, input)
			},
		},
		"create-endpoint-access": {
			Name:   "create-endpoint-access",
			Fields: fields_create_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpointAccess(ctx, input)
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
		"create-hsm-client-certificate": {
			Name:   "create-hsm-client-certificate",
			Fields: fields_create_hsm_client_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHsmClientCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hsm_client_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHsmClientCertificate(ctx, input)
			},
		},
		"create-hsm-configuration": {
			Name:   "create-hsm-configuration",
			Fields: fields_create_hsm_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHsmConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hsm_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHsmConfiguration(ctx, input)
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
		"create-redshift-idc-application": {
			Name:   "create-redshift-idc-application",
			Fields: fields_create_redshift_idc_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRedshiftIdcApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_redshift_idc_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRedshiftIdcApplication(ctx, input)
			},
		},
		"create-scheduled-action": {
			Name:   "create-scheduled-action",
			Fields: fields_create_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScheduledAction(ctx, input)
			},
		},
		"create-snapshot-copy-grant": {
			Name:   "create-snapshot-copy-grant",
			Fields: fields_create_snapshot_copy_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotCopyGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot_copy_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshotCopyGrant(ctx, input)
			},
		},
		"create-snapshot-schedule": {
			Name:   "create-snapshot-schedule",
			Fields: fields_create_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshotSchedule(ctx, input)
			},
		},
		"create-tags": {
			Name:   "create-tags",
			Fields: fields_create_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTags(ctx, input)
			},
		},
		"create-usage-limit": {
			Name:   "create-usage-limit",
			Fields: fields_create_usage_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUsageLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_usage_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUsageLimit(ctx, input)
			},
		},
		"deauthorize-data-share": {
			Name:   "deauthorize-data-share",
			Fields: fields_deauthorize_data_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeauthorizeDataShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deauthorize_data_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeauthorizeDataShare(ctx, input)
			},
		},
		"delete-authentication-profile": {
			Name:   "delete-authentication-profile",
			Fields: fields_delete_authentication_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAuthenticationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_authentication_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAuthenticationProfile(ctx, input)
			},
		},
		"delete-cluster": {
			Name:   "delete-cluster",
			Fields: fields_delete_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCluster(ctx, input)
			},
		},
		"delete-cluster-parameter-group": {
			Name:   "delete-cluster-parameter-group",
			Fields: fields_delete_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClusterParameterGroup(ctx, input)
			},
		},
		"delete-cluster-security-group": {
			Name:   "delete-cluster-security-group",
			Fields: fields_delete_cluster_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClusterSecurityGroup(ctx, input)
			},
		},
		"delete-cluster-snapshot": {
			Name:   "delete-cluster-snapshot",
			Fields: fields_delete_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClusterSnapshot(ctx, input)
			},
		},
		"delete-cluster-subnet-group": {
			Name:   "delete-cluster-subnet-group",
			Fields: fields_delete_cluster_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClusterSubnetGroup(ctx, input)
			},
		},
		"delete-custom-domain-association": {
			Name:   "delete-custom-domain-association",
			Fields: fields_delete_custom_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomDomainAssociation(ctx, input)
			},
		},
		"delete-endpoint-access": {
			Name:   "delete-endpoint-access",
			Fields: fields_delete_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpointAccess(ctx, input)
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
		"delete-hsm-client-certificate": {
			Name:   "delete-hsm-client-certificate",
			Fields: fields_delete_hsm_client_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHsmClientCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hsm_client_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHsmClientCertificate(ctx, input)
			},
		},
		"delete-hsm-configuration": {
			Name:   "delete-hsm-configuration",
			Fields: fields_delete_hsm_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHsmConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hsm_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHsmConfiguration(ctx, input)
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
		"delete-partner": {
			Name:   "delete-partner",
			Fields: fields_delete_partner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePartnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_partner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePartner(ctx, input)
			},
		},
		"delete-redshift-idc-application": {
			Name:   "delete-redshift-idc-application",
			Fields: fields_delete_redshift_idc_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRedshiftIdcApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_redshift_idc_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRedshiftIdcApplication(ctx, input)
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
		"delete-scheduled-action": {
			Name:   "delete-scheduled-action",
			Fields: fields_delete_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduledAction(ctx, input)
			},
		},
		"delete-snapshot-copy-grant": {
			Name:   "delete-snapshot-copy-grant",
			Fields: fields_delete_snapshot_copy_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotCopyGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot_copy_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshotCopyGrant(ctx, input)
			},
		},
		"delete-snapshot-schedule": {
			Name:   "delete-snapshot-schedule",
			Fields: fields_delete_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshotSchedule(ctx, input)
			},
		},
		"delete-tags": {
			Name:   "delete-tags",
			Fields: fields_delete_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTags(ctx, input)
			},
		},
		"delete-usage-limit": {
			Name:   "delete-usage-limit",
			Fields: fields_delete_usage_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUsageLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_usage_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUsageLimit(ctx, input)
			},
		},
		"deregister-namespace": {
			Name:   "deregister-namespace",
			Fields: fields_deregister_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterNamespace(ctx, input)
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
		"describe-authentication-profiles": {
			Name:   "describe-authentication-profiles",
			Fields: fields_describe_authentication_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuthenticationProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_authentication_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuthenticationProfiles(ctx, input)
			},
		},
		"describe-cluster-db-revisions": {
			Name:   "describe-cluster-db-revisions",
			Fields: fields_describe_cluster_db_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterDbRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_db_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterDbRevisions(ctx, input)
				}
				var results []*svc.DescribeClusterDbRevisionsOutput
				p := svc.NewDescribeClusterDbRevisionsPaginator(client, input)
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
		"describe-cluster-parameter-groups": {
			Name:   "describe-cluster-parameter-groups",
			Fields: fields_describe_cluster_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterParameterGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_parameter_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterParameterGroups(ctx, input)
				}
				var results []*svc.DescribeClusterParameterGroupsOutput
				p := svc.NewDescribeClusterParameterGroupsPaginator(client, input)
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
		"describe-cluster-parameters": {
			Name:   "describe-cluster-parameters",
			Fields: fields_describe_cluster_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterParameters(ctx, input)
				}
				var results []*svc.DescribeClusterParametersOutput
				p := svc.NewDescribeClusterParametersPaginator(client, input)
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
		"describe-cluster-security-groups": {
			Name:   "describe-cluster-security-groups",
			Fields: fields_describe_cluster_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterSecurityGroups(ctx, input)
				}
				var results []*svc.DescribeClusterSecurityGroupsOutput
				p := svc.NewDescribeClusterSecurityGroupsPaginator(client, input)
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
		"describe-cluster-snapshots": {
			Name:   "describe-cluster-snapshots",
			Fields: fields_describe_cluster_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterSnapshots(ctx, input)
				}
				var results []*svc.DescribeClusterSnapshotsOutput
				p := svc.NewDescribeClusterSnapshotsPaginator(client, input)
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
		"describe-cluster-subnet-groups": {
			Name:   "describe-cluster-subnet-groups",
			Fields: fields_describe_cluster_subnet_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterSubnetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_subnet_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterSubnetGroups(ctx, input)
				}
				var results []*svc.DescribeClusterSubnetGroupsOutput
				p := svc.NewDescribeClusterSubnetGroupsPaginator(client, input)
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
		"describe-cluster-tracks": {
			Name:   "describe-cluster-tracks",
			Fields: fields_describe_cluster_tracks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterTracksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_tracks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterTracks(ctx, input)
				}
				var results []*svc.DescribeClusterTracksOutput
				p := svc.NewDescribeClusterTracksPaginator(client, input)
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
		"describe-cluster-versions": {
			Name:   "describe-cluster-versions",
			Fields: fields_describe_cluster_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterVersions(ctx, input)
				}
				var results []*svc.DescribeClusterVersionsOutput
				p := svc.NewDescribeClusterVersionsPaginator(client, input)
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
		"describe-clusters": {
			Name:   "describe-clusters",
			Fields: fields_describe_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusters(ctx, input)
				}
				var results []*svc.DescribeClustersOutput
				p := svc.NewDescribeClustersPaginator(client, input)
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
		"describe-custom-domain-associations": {
			Name:   "describe-custom-domain-associations",
			Fields: fields_describe_custom_domain_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomDomainAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_custom_domain_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCustomDomainAssociations(ctx, input)
				}
				var results []*svc.DescribeCustomDomainAssociationsOutput
				p := svc.NewDescribeCustomDomainAssociationsPaginator(client, input)
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
		"describe-data-shares": {
			Name:   "describe-data-shares",
			Fields: fields_describe_data_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataShares(ctx, input)
				}
				var results []*svc.DescribeDataSharesOutput
				p := svc.NewDescribeDataSharesPaginator(client, input)
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
		"describe-data-shares-for-consumer": {
			Name:   "describe-data-shares-for-consumer",
			Fields: fields_describe_data_shares_for_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSharesForConsumerInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_shares_for_consumer, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataSharesForConsumer(ctx, input)
				}
				var results []*svc.DescribeDataSharesForConsumerOutput
				p := svc.NewDescribeDataSharesForConsumerPaginator(client, input)
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
		"describe-data-shares-for-producer": {
			Name:   "describe-data-shares-for-producer",
			Fields: fields_describe_data_shares_for_producer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSharesForProducerInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_shares_for_producer, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataSharesForProducer(ctx, input)
				}
				var results []*svc.DescribeDataSharesForProducerOutput
				p := svc.NewDescribeDataSharesForProducerPaginator(client, input)
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
		"describe-default-cluster-parameters": {
			Name:   "describe-default-cluster-parameters",
			Fields: fields_describe_default_cluster_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDefaultClusterParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_default_cluster_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDefaultClusterParameters(ctx, input)
				}
				var results []*svc.DescribeDefaultClusterParametersOutput
				p := svc.NewDescribeDefaultClusterParametersPaginator(client, input)
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
		"describe-endpoint-access": {
			Name:   "describe-endpoint-access",
			Fields: fields_describe_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointAccessInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_endpoint_access, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEndpointAccess(ctx, input)
				}
				var results []*svc.DescribeEndpointAccessOutput
				p := svc.NewDescribeEndpointAccessPaginator(client, input)
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
		"describe-endpoint-authorization": {
			Name:   "describe-endpoint-authorization",
			Fields: fields_describe_endpoint_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointAuthorizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_endpoint_authorization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEndpointAuthorization(ctx, input)
				}
				var results []*svc.DescribeEndpointAuthorizationOutput
				p := svc.NewDescribeEndpointAuthorizationPaginator(client, input)
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
		"describe-hsm-client-certificates": {
			Name:   "describe-hsm-client-certificates",
			Fields: fields_describe_hsm_client_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHsmClientCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_hsm_client_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeHsmClientCertificates(ctx, input)
				}
				var results []*svc.DescribeHsmClientCertificatesOutput
				p := svc.NewDescribeHsmClientCertificatesPaginator(client, input)
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
		"describe-hsm-configurations": {
			Name:   "describe-hsm-configurations",
			Fields: fields_describe_hsm_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHsmConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_hsm_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeHsmConfigurations(ctx, input)
				}
				var results []*svc.DescribeHsmConfigurationsOutput
				p := svc.NewDescribeHsmConfigurationsPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_inbound_integrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInboundIntegrations(ctx, input)
				}
				var results []*svc.DescribeInboundIntegrationsOutput
				p := svc.NewDescribeInboundIntegrationsPaginator(client, input)
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
		"describe-logging-status": {
			Name:   "describe-logging-status",
			Fields: fields_describe_logging_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoggingStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_logging_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoggingStatus(ctx, input)
			},
		},
		"describe-node-configuration-options": {
			Name:   "describe-node-configuration-options",
			Fields: fields_describe_node_configuration_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNodeConfigurationOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_node_configuration_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNodeConfigurationOptions(ctx, input)
				}
				var results []*svc.DescribeNodeConfigurationOptionsOutput
				p := svc.NewDescribeNodeConfigurationOptionsPaginator(client, input)
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
		"describe-orderable-cluster-options": {
			Name:   "describe-orderable-cluster-options",
			Fields: fields_describe_orderable_cluster_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrderableClusterOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_orderable_cluster_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrderableClusterOptions(ctx, input)
				}
				var results []*svc.DescribeOrderableClusterOptionsOutput
				p := svc.NewDescribeOrderableClusterOptionsPaginator(client, input)
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
		"describe-partners": {
			Name:   "describe-partners",
			Fields: fields_describe_partners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePartnersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_partners, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePartners(ctx, input)
			},
		},
		"describe-redshift-idc-applications": {
			Name:   "describe-redshift-idc-applications",
			Fields: fields_describe_redshift_idc_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRedshiftIdcApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_redshift_idc_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRedshiftIdcApplications(ctx, input)
				}
				var results []*svc.DescribeRedshiftIdcApplicationsOutput
				p := svc.NewDescribeRedshiftIdcApplicationsPaginator(client, input)
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
		"describe-reserved-node-exchange-status": {
			Name:   "describe-reserved-node-exchange-status",
			Fields: fields_describe_reserved_node_exchange_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedNodeExchangeStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_node_exchange_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedNodeExchangeStatus(ctx, input)
				}
				var results []*svc.DescribeReservedNodeExchangeStatusOutput
				p := svc.NewDescribeReservedNodeExchangeStatusPaginator(client, input)
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
		"describe-reserved-node-offerings": {
			Name:   "describe-reserved-node-offerings",
			Fields: fields_describe_reserved_node_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedNodeOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_node_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedNodeOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedNodeOfferingsOutput
				p := svc.NewDescribeReservedNodeOfferingsPaginator(client, input)
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
		"describe-reserved-nodes": {
			Name:   "describe-reserved-nodes",
			Fields: fields_describe_reserved_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedNodes(ctx, input)
				}
				var results []*svc.DescribeReservedNodesOutput
				p := svc.NewDescribeReservedNodesPaginator(client, input)
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
		"describe-resize": {
			Name:   "describe-resize",
			Fields: fields_describe_resize,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResizeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resize, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResize(ctx, input)
			},
		},
		"describe-scheduled-actions": {
			Name:   "describe-scheduled-actions",
			Fields: fields_describe_scheduled_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scheduled_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScheduledActions(ctx, input)
				}
				var results []*svc.DescribeScheduledActionsOutput
				p := svc.NewDescribeScheduledActionsPaginator(client, input)
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
		"describe-snapshot-copy-grants": {
			Name:   "describe-snapshot-copy-grants",
			Fields: fields_describe_snapshot_copy_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotCopyGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_snapshot_copy_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSnapshotCopyGrants(ctx, input)
				}
				var results []*svc.DescribeSnapshotCopyGrantsOutput
				p := svc.NewDescribeSnapshotCopyGrantsPaginator(client, input)
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
		"describe-snapshot-schedules": {
			Name:   "describe-snapshot-schedules",
			Fields: fields_describe_snapshot_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotSchedulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_snapshot_schedules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSnapshotSchedules(ctx, input)
				}
				var results []*svc.DescribeSnapshotSchedulesOutput
				p := svc.NewDescribeSnapshotSchedulesPaginator(client, input)
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
		"describe-storage": {
			Name:   "describe-storage",
			Fields: fields_describe_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStorage(ctx, input)
			},
		},
		"describe-table-restore-status": {
			Name:   "describe-table-restore-status",
			Fields: fields_describe_table_restore_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTableRestoreStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_table_restore_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTableRestoreStatus(ctx, input)
				}
				var results []*svc.DescribeTableRestoreStatusOutput
				p := svc.NewDescribeTableRestoreStatusPaginator(client, input)
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
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTags(ctx, input)
				}
				var results []*svc.DescribeTagsOutput
				p := svc.NewDescribeTagsPaginator(client, input)
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
		"describe-usage-limits": {
			Name:   "describe-usage-limits",
			Fields: fields_describe_usage_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUsageLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_usage_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeUsageLimits(ctx, input)
				}
				var results []*svc.DescribeUsageLimitsOutput
				p := svc.NewDescribeUsageLimitsPaginator(client, input)
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
		"disable-logging": {
			Name:   "disable-logging",
			Fields: fields_disable_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableLogging(ctx, input)
			},
		},
		"disable-snapshot-copy": {
			Name:   "disable-snapshot-copy",
			Fields: fields_disable_snapshot_copy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSnapshotCopyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_snapshot_copy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSnapshotCopy(ctx, input)
			},
		},
		"disassociate-data-share-consumer": {
			Name:   "disassociate-data-share-consumer",
			Fields: fields_disassociate_data_share_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDataShareConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_data_share_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDataShareConsumer(ctx, input)
			},
		},
		"enable-logging": {
			Name:   "enable-logging",
			Fields: fields_enable_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableLogging(ctx, input)
			},
		},
		"enable-snapshot-copy": {
			Name:   "enable-snapshot-copy",
			Fields: fields_enable_snapshot_copy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSnapshotCopyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_snapshot_copy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSnapshotCopy(ctx, input)
			},
		},
		"failover-primary-compute": {
			Name:   "failover-primary-compute",
			Fields: fields_failover_primary_compute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FailoverPrimaryComputeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_failover_primary_compute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FailoverPrimaryCompute(ctx, input)
			},
		},
		"get-cluster-credentials": {
			Name:   "get-cluster-credentials",
			Fields: fields_get_cluster_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClusterCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cluster_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClusterCredentials(ctx, input)
			},
		},
		"get-cluster-credentials-with-iam": {
			Name:   "get-cluster-credentials-with-iam",
			Fields: fields_get_cluster_credentials_with_iam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClusterCredentialsWithIAMInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cluster_credentials_with_iam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClusterCredentialsWithIAM(ctx, input)
			},
		},
		"get-identity-center-auth-token": {
			Name:   "get-identity-center-auth-token",
			Fields: fields_get_identity_center_auth_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityCenterAuthTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_center_auth_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityCenterAuthToken(ctx, input)
			},
		},
		"get-reserved-node-exchange-configuration-options": {
			Name:   "get-reserved-node-exchange-configuration-options",
			Fields: fields_get_reserved_node_exchange_configuration_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservedNodeExchangeConfigurationOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_reserved_node_exchange_configuration_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetReservedNodeExchangeConfigurationOptions(ctx, input)
				}
				var results []*svc.GetReservedNodeExchangeConfigurationOptionsOutput
				p := svc.NewGetReservedNodeExchangeConfigurationOptionsPaginator(client, input)
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
		"get-reserved-node-exchange-offerings": {
			Name:   "get-reserved-node-exchange-offerings",
			Fields: fields_get_reserved_node_exchange_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservedNodeExchangeOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_reserved_node_exchange_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetReservedNodeExchangeOfferings(ctx, input)
				}
				var results []*svc.GetReservedNodeExchangeOfferingsOutput
				p := svc.NewGetReservedNodeExchangeOfferingsPaginator(client, input)
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
		"modify-aqua-configuration": {
			Name:   "modify-aqua-configuration",
			Fields: fields_modify_aqua_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyAquaConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_aqua_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyAquaConfiguration(ctx, input)
			},
		},
		"modify-authentication-profile": {
			Name:   "modify-authentication-profile",
			Fields: fields_modify_authentication_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyAuthenticationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_authentication_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyAuthenticationProfile(ctx, input)
			},
		},
		"modify-cluster": {
			Name:   "modify-cluster",
			Fields: fields_modify_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCluster(ctx, input)
			},
		},
		"modify-cluster-db-revision": {
			Name:   "modify-cluster-db-revision",
			Fields: fields_modify_cluster_db_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterDbRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_db_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterDbRevision(ctx, input)
			},
		},
		"modify-cluster-iam-roles": {
			Name:   "modify-cluster-iam-roles",
			Fields: fields_modify_cluster_iam_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterIamRolesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_iam_roles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterIamRoles(ctx, input)
			},
		},
		"modify-cluster-maintenance": {
			Name:   "modify-cluster-maintenance",
			Fields: fields_modify_cluster_maintenance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterMaintenanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_maintenance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterMaintenance(ctx, input)
			},
		},
		"modify-cluster-parameter-group": {
			Name:   "modify-cluster-parameter-group",
			Fields: fields_modify_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterParameterGroup(ctx, input)
			},
		},
		"modify-cluster-snapshot": {
			Name:   "modify-cluster-snapshot",
			Fields: fields_modify_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterSnapshot(ctx, input)
			},
		},
		"modify-cluster-snapshot-schedule": {
			Name:   "modify-cluster-snapshot-schedule",
			Fields: fields_modify_cluster_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterSnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterSnapshotSchedule(ctx, input)
			},
		},
		"modify-cluster-subnet-group": {
			Name:   "modify-cluster-subnet-group",
			Fields: fields_modify_cluster_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClusterSubnetGroup(ctx, input)
			},
		},
		"modify-custom-domain-association": {
			Name:   "modify-custom-domain-association",
			Fields: fields_modify_custom_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCustomDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_custom_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCustomDomainAssociation(ctx, input)
			},
		},
		"modify-endpoint-access": {
			Name:   "modify-endpoint-access",
			Fields: fields_modify_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyEndpointAccess(ctx, input)
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
		"modify-lakehouse-configuration": {
			Name:   "modify-lakehouse-configuration",
			Fields: fields_modify_lakehouse_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyLakehouseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_lakehouse_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyLakehouseConfiguration(ctx, input)
			},
		},
		"modify-redshift-idc-application": {
			Name:   "modify-redshift-idc-application",
			Fields: fields_modify_redshift_idc_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyRedshiftIdcApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_redshift_idc_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyRedshiftIdcApplication(ctx, input)
			},
		},
		"modify-scheduled-action": {
			Name:   "modify-scheduled-action",
			Fields: fields_modify_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyScheduledAction(ctx, input)
			},
		},
		"modify-snapshot-copy-retention-period": {
			Name:   "modify-snapshot-copy-retention-period",
			Fields: fields_modify_snapshot_copy_retention_period,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySnapshotCopyRetentionPeriodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_snapshot_copy_retention_period, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySnapshotCopyRetentionPeriod(ctx, input)
			},
		},
		"modify-snapshot-schedule": {
			Name:   "modify-snapshot-schedule",
			Fields: fields_modify_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySnapshotSchedule(ctx, input)
			},
		},
		"modify-usage-limit": {
			Name:   "modify-usage-limit",
			Fields: fields_modify_usage_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyUsageLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_usage_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyUsageLimit(ctx, input)
			},
		},
		"pause-cluster": {
			Name:   "pause-cluster",
			Fields: fields_pause_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PauseClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_pause_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PauseCluster(ctx, input)
			},
		},
		"purchase-reserved-node-offering": {
			Name:   "purchase-reserved-node-offering",
			Fields: fields_purchase_reserved_node_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedNodeOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_node_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedNodeOffering(ctx, input)
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
		"reboot-cluster": {
			Name:   "reboot-cluster",
			Fields: fields_reboot_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootCluster(ctx, input)
			},
		},
		"register-namespace": {
			Name:   "register-namespace",
			Fields: fields_register_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterNamespace(ctx, input)
			},
		},
		"reject-data-share": {
			Name:   "reject-data-share",
			Fields: fields_reject_data_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectDataShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_data_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectDataShare(ctx, input)
			},
		},
		"reset-cluster-parameter-group": {
			Name:   "reset-cluster-parameter-group",
			Fields: fields_reset_cluster_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetClusterParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_cluster_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetClusterParameterGroup(ctx, input)
			},
		},
		"resize-cluster": {
			Name:   "resize-cluster",
			Fields: fields_resize_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResizeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resize_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResizeCluster(ctx, input)
			},
		},
		"restore-from-cluster-snapshot": {
			Name:   "restore-from-cluster-snapshot",
			Fields: fields_restore_from_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreFromClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_from_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreFromClusterSnapshot(ctx, input)
			},
		},
		"restore-table-from-cluster-snapshot": {
			Name:   "restore-table-from-cluster-snapshot",
			Fields: fields_restore_table_from_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreTableFromClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_table_from_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreTableFromClusterSnapshot(ctx, input)
			},
		},
		"resume-cluster": {
			Name:   "resume-cluster",
			Fields: fields_resume_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeCluster(ctx, input)
			},
		},
		"revoke-cluster-security-group-ingress": {
			Name:   "revoke-cluster-security-group-ingress",
			Fields: fields_revoke_cluster_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeClusterSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_cluster_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeClusterSecurityGroupIngress(ctx, input)
			},
		},
		"revoke-endpoint-access": {
			Name:   "revoke-endpoint-access",
			Fields: fields_revoke_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeEndpointAccess(ctx, input)
			},
		},
		"revoke-snapshot-access": {
			Name:   "revoke-snapshot-access",
			Fields: fields_revoke_snapshot_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeSnapshotAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_snapshot_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeSnapshotAccess(ctx, input)
			},
		},
		"rotate-encryption-key": {
			Name:   "rotate-encryption-key",
			Fields: fields_rotate_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RotateEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rotate_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RotateEncryptionKey(ctx, input)
			},
		},
		"update-partner-status": {
			Name:   "update-partner-status",
			Fields: fields_update_partner_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePartnerStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_partner_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePartnerStatus(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("redshift", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
