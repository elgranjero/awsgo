package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
)

var fields_convert_recovery_point_to_snapshot = []leanruntime.Field{
	{Name: "RecoveryPointId", Flag: "recovery-point-id", Type: "*string", Required: true},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_domain_association = []leanruntime.Field{
	{Name: "CustomDomainCertificateArn", Flag: "custom-domain-certificate-arn", Type: "*string", Required: true},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_create_endpoint_access = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_create_namespace = []leanruntime.Field{
	{Name: "AdminPasswordSecretKmsKeyId", Flag: "admin-password-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "AdminUserPassword", Flag: "admin-user-password", Type: "*string", Required: false},
	{Name: "AdminUsername", Flag: "admin-username", Type: "*string", Required: false},
	{Name: "DbName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DefaultIamRoleArn", Flag: "default-iam-role-arn", Type: "*string", Required: false},
	{Name: "IamRoles", Flag: "iam-roles", Type: "[]string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LogExports", Flag: "log-exports", Type: "[]types.LogExport", Required: false},
	{Name: "ManageAdminPassword", Flag: "manage-admin-password", Type: "*bool", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "RedshiftIdcApplicationArn", Flag: "redshift-idc-application-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_reservation = []leanruntime.Field{
	{Name: "Capacity", Flag: "capacity", Type: "int32", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
}

var fields_create_scheduled_action = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "types.Schedule", Required: true},
	{Name: "ScheduledActionDescription", Flag: "scheduled-action-description", Type: "*string", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TargetAction", Flag: "target-action", Type: "types.TargetAction", Required: true},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_snapshot_copy_configuration = []leanruntime.Field{
	{Name: "DestinationKmsKeyId", Flag: "destination-kms-key-id", Type: "*string", Required: false},
	{Name: "DestinationRegion", Flag: "destination-region", Type: "*string", Required: true},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "SnapshotRetentionPeriod", Flag: "snapshot-retention-period", Type: "*int32", Required: false},
}

var fields_create_usage_limit = []leanruntime.Field{
	{Name: "Amount", Flag: "amount", Type: "*int64", Required: true},
	{Name: "BreachAction", Flag: "breach-action", Type: "types.UsageLimitBreachAction", Required: false},
	{Name: "Period", Flag: "period", Type: "types.UsageLimitPeriod", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "UsageType", Flag: "usage-type", Type: "types.UsageLimitUsageType", Required: true},
}

var fields_create_workgroup = []leanruntime.Field{
	{Name: "BaseCapacity", Flag: "base-capacity", Type: "*int32", Required: false},
	{Name: "ConfigParameters", Flag: "config-parameters", Type: "[]types.ConfigParameter", Required: false},
	{Name: "EnhancedVpcRouting", Flag: "enhanced-vpc-routing", Type: "*bool", Required: false},
	{Name: "ExtraComputeForAutomaticOptimization", Flag: "extra-compute-for-automatic-optimization", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "*string", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*int32", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PricePerformanceTarget", Flag: "price-performance-target", Type: "*types.PerformanceTarget", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrackName", Flag: "track-name", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_delete_custom_domain_association = []leanruntime.Field{
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_delete_endpoint_access = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_delete_namespace = []leanruntime.Field{
	{Name: "FinalSnapshotName", Flag: "final-snapshot-name", Type: "*string", Required: false},
	{Name: "FinalSnapshotRetentionPeriod", Flag: "final-snapshot-retention-period", Type: "*int32", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_scheduled_action = []leanruntime.Field{
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
}

var fields_delete_snapshot = []leanruntime.Field{
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_delete_snapshot_copy_configuration = []leanruntime.Field{
	{Name: "SnapshotCopyConfigurationId", Flag: "snapshot-copy-configuration-id", Type: "*string", Required: true},
}

var fields_delete_usage_limit = []leanruntime.Field{
	{Name: "UsageLimitId", Flag: "usage-limit-id", Type: "*string", Required: true},
}

var fields_delete_workgroup = []leanruntime.Field{
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_get_credentials = []leanruntime.Field{
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "DbName", Flag: "db-name", Type: "*string", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_get_custom_domain_association = []leanruntime.Field{
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_get_endpoint_access = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_get_identity_center_auth_token = []leanruntime.Field{
	{Name: "WorkgroupNames", Flag: "workgroup-names", Type: "[]string", Required: true},
}

var fields_get_namespace = []leanruntime.Field{
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
}

var fields_get_recovery_point = []leanruntime.Field{
	{Name: "RecoveryPointId", Flag: "recovery-point-id", Type: "*string", Required: true},
}

var fields_get_reservation = []leanruntime.Field{
	{Name: "ReservationId", Flag: "reservation-id", Type: "*string", Required: true},
}

var fields_get_reservation_offering = []leanruntime.Field{
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_scheduled_action = []leanruntime.Field{
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
}

var fields_get_snapshot = []leanruntime.Field{
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
}

var fields_get_table_restore_status = []leanruntime.Field{
	{Name: "TableRestoreRequestId", Flag: "table-restore-request-id", Type: "*string", Required: true},
}

var fields_get_track = []leanruntime.Field{
	{Name: "TrackName", Flag: "track-name", Type: "*string", Required: true},
}

var fields_get_usage_limit = []leanruntime.Field{
	{Name: "UsageLimitId", Flag: "usage-limit-id", Type: "*string", Required: true},
}

var fields_get_workgroup = []leanruntime.Field{
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_list_custom_domain_associations = []leanruntime.Field{
	{Name: "CustomDomainCertificateArn", Flag: "custom-domain-certificate-arn", Type: "*string", Required: false},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_endpoint_access = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_list_managed_workgroups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
}

var fields_list_namespaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recovery_points = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceArn", Flag: "namespace-arn", Type: "*string", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_reservation_offerings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_reservations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_scheduled_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_snapshot_copy_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_snapshots = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceArn", Flag: "namespace-arn", Type: "*string", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_table_restore_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tracks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_usage_limits = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "UsageType", Flag: "usage-type", Type: "types.UsageLimitUsageType", Required: false},
}

var fields_list_workgroups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_restore_from_recovery_point = []leanruntime.Field{
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "RecoveryPointId", Flag: "recovery-point-id", Type: "*string", Required: true},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_restore_from_snapshot = []leanruntime.Field{
	{Name: "AdminPasswordSecretKmsKeyId", Flag: "admin-password-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "ManageAdminPassword", Flag: "manage-admin-password", Type: "*bool", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_restore_table_from_recovery_point = []leanruntime.Field{
	{Name: "ActivateCaseSensitiveIdentifier", Flag: "activate-case-sensitive-identifier", Type: "*bool", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "NewTableName", Flag: "new-table-name", Type: "*string", Required: true},
	{Name: "RecoveryPointId", Flag: "recovery-point-id", Type: "*string", Required: true},
	{Name: "SourceDatabaseName", Flag: "source-database-name", Type: "*string", Required: true},
	{Name: "SourceSchemaName", Flag: "source-schema-name", Type: "*string", Required: false},
	{Name: "SourceTableName", Flag: "source-table-name", Type: "*string", Required: true},
	{Name: "TargetDatabaseName", Flag: "target-database-name", Type: "*string", Required: false},
	{Name: "TargetSchemaName", Flag: "target-schema-name", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_restore_table_from_snapshot = []leanruntime.Field{
	{Name: "ActivateCaseSensitiveIdentifier", Flag: "activate-case-sensitive-identifier", Type: "*bool", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "NewTableName", Flag: "new-table-name", Type: "*string", Required: true},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "SourceDatabaseName", Flag: "source-database-name", Type: "*string", Required: true},
	{Name: "SourceSchemaName", Flag: "source-schema-name", Type: "*string", Required: false},
	{Name: "SourceTableName", Flag: "source-table-name", Type: "*string", Required: true},
	{Name: "TargetDatabaseName", Flag: "target-database-name", Type: "*string", Required: false},
	{Name: "TargetSchemaName", Flag: "target-schema-name", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_custom_domain_association = []leanruntime.Field{
	{Name: "CustomDomainCertificateArn", Flag: "custom-domain-certificate-arn", Type: "*string", Required: true},
	{Name: "CustomDomainName", Flag: "custom-domain-name", Type: "*string", Required: true},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

var fields_update_endpoint_access = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_update_lakehouse_configuration = []leanruntime.Field{
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LakehouseIdcApplicationArn", Flag: "lakehouse-idc-application-arn", Type: "*string", Required: false},
	{Name: "LakehouseIdcRegistration", Flag: "lakehouse-idc-registration", Type: "types.LakehouseIdcRegistration", Required: false},
	{Name: "LakehouseRegistration", Flag: "lakehouse-registration", Type: "types.LakehouseRegistration", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
}

var fields_update_namespace = []leanruntime.Field{
	{Name: "AdminPasswordSecretKmsKeyId", Flag: "admin-password-secret-kms-key-id", Type: "*string", Required: false},
	{Name: "AdminUserPassword", Flag: "admin-user-password", Type: "*string", Required: false},
	{Name: "AdminUsername", Flag: "admin-username", Type: "*string", Required: false},
	{Name: "DefaultIamRoleArn", Flag: "default-iam-role-arn", Type: "*string", Required: false},
	{Name: "IamRoles", Flag: "iam-roles", Type: "[]string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LogExports", Flag: "log-exports", Type: "[]types.LogExport", Required: false},
	{Name: "ManageAdminPassword", Flag: "manage-admin-password", Type: "*bool", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
}

var fields_update_scheduled_action = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "types.Schedule", Required: false},
	{Name: "ScheduledActionDescription", Flag: "scheduled-action-description", Type: "*string", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TargetAction", Flag: "target-action", Type: "types.TargetAction", Required: false},
}

var fields_update_snapshot = []leanruntime.Field{
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*int32", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_update_snapshot_copy_configuration = []leanruntime.Field{
	{Name: "SnapshotCopyConfigurationId", Flag: "snapshot-copy-configuration-id", Type: "*string", Required: true},
	{Name: "SnapshotRetentionPeriod", Flag: "snapshot-retention-period", Type: "*int32", Required: false},
}

var fields_update_usage_limit = []leanruntime.Field{
	{Name: "Amount", Flag: "amount", Type: "*int64", Required: false},
	{Name: "BreachAction", Flag: "breach-action", Type: "types.UsageLimitBreachAction", Required: false},
	{Name: "UsageLimitId", Flag: "usage-limit-id", Type: "*string", Required: true},
}

var fields_update_workgroup = []leanruntime.Field{
	{Name: "BaseCapacity", Flag: "base-capacity", Type: "*int32", Required: false},
	{Name: "ConfigParameters", Flag: "config-parameters", Type: "[]types.ConfigParameter", Required: false},
	{Name: "EnhancedVpcRouting", Flag: "enhanced-vpc-routing", Type: "*bool", Required: false},
	{Name: "ExtraComputeForAutomaticOptimization", Flag: "extra-compute-for-automatic-optimization", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "*string", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PricePerformanceTarget", Flag: "price-performance-target", Type: "*types.PerformanceTarget", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "TrackName", Flag: "track-name", Type: "*string", Required: false},
	{Name: "WorkgroupName", Flag: "workgroup-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"convert-recovery-point-to-snapshot": {
			Name:   "convert-recovery-point-to-snapshot",
			Fields: fields_convert_recovery_point_to_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConvertRecoveryPointToSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_convert_recovery_point_to_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConvertRecoveryPointToSnapshot(ctx, input)
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
		"create-namespace": {
			Name:   "create-namespace",
			Fields: fields_create_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNamespace(ctx, input)
			},
		},
		"create-reservation": {
			Name:   "create-reservation",
			Fields: fields_create_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReservation(ctx, input)
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
		"create-snapshot": {
			Name:   "create-snapshot",
			Fields: fields_create_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshot(ctx, input)
			},
		},
		"create-snapshot-copy-configuration": {
			Name:   "create-snapshot-copy-configuration",
			Fields: fields_create_snapshot_copy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotCopyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot_copy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshotCopyConfiguration(ctx, input)
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
		"create-workgroup": {
			Name:   "create-workgroup",
			Fields: fields_create_workgroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkgroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workgroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkgroup(ctx, input)
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
		"delete-namespace": {
			Name:   "delete-namespace",
			Fields: fields_delete_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNamespace(ctx, input)
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
		"delete-snapshot": {
			Name:   "delete-snapshot",
			Fields: fields_delete_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshot(ctx, input)
			},
		},
		"delete-snapshot-copy-configuration": {
			Name:   "delete-snapshot-copy-configuration",
			Fields: fields_delete_snapshot_copy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotCopyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot_copy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshotCopyConfiguration(ctx, input)
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
		"delete-workgroup": {
			Name:   "delete-workgroup",
			Fields: fields_delete_workgroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkgroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workgroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkgroup(ctx, input)
			},
		},
		"get-credentials": {
			Name:   "get-credentials",
			Fields: fields_get_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCredentials(ctx, input)
			},
		},
		"get-custom-domain-association": {
			Name:   "get-custom-domain-association",
			Fields: fields_get_custom_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomDomainAssociation(ctx, input)
			},
		},
		"get-endpoint-access": {
			Name:   "get-endpoint-access",
			Fields: fields_get_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEndpointAccess(ctx, input)
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
		"get-namespace": {
			Name:   "get-namespace",
			Fields: fields_get_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNamespace(ctx, input)
			},
		},
		"get-recovery-point": {
			Name:   "get-recovery-point",
			Fields: fields_get_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecoveryPoint(ctx, input)
			},
		},
		"get-reservation": {
			Name:   "get-reservation",
			Fields: fields_get_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReservation(ctx, input)
			},
		},
		"get-reservation-offering": {
			Name:   "get-reservation-offering",
			Fields: fields_get_reservation_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservationOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reservation_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReservationOffering(ctx, input)
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
		"get-scheduled-action": {
			Name:   "get-scheduled-action",
			Fields: fields_get_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScheduledAction(ctx, input)
			},
		},
		"get-snapshot": {
			Name:   "get-snapshot",
			Fields: fields_get_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSnapshot(ctx, input)
			},
		},
		"get-table-restore-status": {
			Name:   "get-table-restore-status",
			Fields: fields_get_table_restore_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableRestoreStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_restore_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableRestoreStatus(ctx, input)
			},
		},
		"get-track": {
			Name:   "get-track",
			Fields: fields_get_track,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_track, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrack(ctx, input)
			},
		},
		"get-usage-limit": {
			Name:   "get-usage-limit",
			Fields: fields_get_usage_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsageLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_usage_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsageLimit(ctx, input)
			},
		},
		"get-workgroup": {
			Name:   "get-workgroup",
			Fields: fields_get_workgroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkgroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workgroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkgroup(ctx, input)
			},
		},
		"list-custom-domain-associations": {
			Name:   "list-custom-domain-associations",
			Fields: fields_list_custom_domain_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomDomainAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_domain_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomDomainAssociations(ctx, input)
				}
				var results []*svc.ListCustomDomainAssociationsOutput
				p := svc.NewListCustomDomainAssociationsPaginator(client, input)
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
		"list-endpoint-access": {
			Name:   "list-endpoint-access",
			Fields: fields_list_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointAccessInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoint_access, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpointAccess(ctx, input)
				}
				var results []*svc.ListEndpointAccessOutput
				p := svc.NewListEndpointAccessPaginator(client, input)
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
		"list-managed-workgroups": {
			Name:   "list-managed-workgroups",
			Fields: fields_list_managed_workgroups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedWorkgroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_workgroups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedWorkgroups(ctx, input)
				}
				var results []*svc.ListManagedWorkgroupsOutput
				p := svc.NewListManagedWorkgroupsPaginator(client, input)
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
		"list-namespaces": {
			Name:   "list-namespaces",
			Fields: fields_list_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNamespaces(ctx, input)
				}
				var results []*svc.ListNamespacesOutput
				p := svc.NewListNamespacesPaginator(client, input)
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
		"list-recovery-points": {
			Name:   "list-recovery-points",
			Fields: fields_list_recovery_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecoveryPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recovery_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecoveryPoints(ctx, input)
				}
				var results []*svc.ListRecoveryPointsOutput
				p := svc.NewListRecoveryPointsPaginator(client, input)
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
		"list-reservation-offerings": {
			Name:   "list-reservation-offerings",
			Fields: fields_list_reservation_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReservationOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reservation_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReservationOfferings(ctx, input)
				}
				var results []*svc.ListReservationOfferingsOutput
				p := svc.NewListReservationOfferingsPaginator(client, input)
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
		"list-reservations": {
			Name:   "list-reservations",
			Fields: fields_list_reservations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReservationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reservations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReservations(ctx, input)
				}
				var results []*svc.ListReservationsOutput
				p := svc.NewListReservationsPaginator(client, input)
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
		"list-scheduled-actions": {
			Name:   "list-scheduled-actions",
			Fields: fields_list_scheduled_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScheduledActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scheduled_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScheduledActions(ctx, input)
				}
				var results []*svc.ListScheduledActionsOutput
				p := svc.NewListScheduledActionsPaginator(client, input)
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
		"list-snapshot-copy-configurations": {
			Name:   "list-snapshot-copy-configurations",
			Fields: fields_list_snapshot_copy_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSnapshotCopyConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_snapshot_copy_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSnapshotCopyConfigurations(ctx, input)
				}
				var results []*svc.ListSnapshotCopyConfigurationsOutput
				p := svc.NewListSnapshotCopyConfigurationsPaginator(client, input)
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
		"list-snapshots": {
			Name:   "list-snapshots",
			Fields: fields_list_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSnapshots(ctx, input)
				}
				var results []*svc.ListSnapshotsOutput
				p := svc.NewListSnapshotsPaginator(client, input)
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
		"list-table-restore-status": {
			Name:   "list-table-restore-status",
			Fields: fields_list_table_restore_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTableRestoreStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_table_restore_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTableRestoreStatus(ctx, input)
				}
				var results []*svc.ListTableRestoreStatusOutput
				p := svc.NewListTableRestoreStatusPaginator(client, input)
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
		"list-tracks": {
			Name:   "list-tracks",
			Fields: fields_list_tracks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTracksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tracks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTracks(ctx, input)
				}
				var results []*svc.ListTracksOutput
				p := svc.NewListTracksPaginator(client, input)
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
		"list-usage-limits": {
			Name:   "list-usage-limits",
			Fields: fields_list_usage_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsageLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_usage_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsageLimits(ctx, input)
				}
				var results []*svc.ListUsageLimitsOutput
				p := svc.NewListUsageLimitsPaginator(client, input)
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
		"list-workgroups": {
			Name:   "list-workgroups",
			Fields: fields_list_workgroups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkgroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workgroups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkgroups(ctx, input)
				}
				var results []*svc.ListWorkgroupsOutput
				p := svc.NewListWorkgroupsPaginator(client, input)
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
		"restore-from-recovery-point": {
			Name:   "restore-from-recovery-point",
			Fields: fields_restore_from_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreFromRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_from_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreFromRecoveryPoint(ctx, input)
			},
		},
		"restore-from-snapshot": {
			Name:   "restore-from-snapshot",
			Fields: fields_restore_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreFromSnapshot(ctx, input)
			},
		},
		"restore-table-from-recovery-point": {
			Name:   "restore-table-from-recovery-point",
			Fields: fields_restore_table_from_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreTableFromRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_table_from_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreTableFromRecoveryPoint(ctx, input)
			},
		},
		"restore-table-from-snapshot": {
			Name:   "restore-table-from-snapshot",
			Fields: fields_restore_table_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreTableFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_table_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreTableFromSnapshot(ctx, input)
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
		"update-custom-domain-association": {
			Name:   "update-custom-domain-association",
			Fields: fields_update_custom_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomDomainAssociation(ctx, input)
			},
		},
		"update-endpoint-access": {
			Name:   "update-endpoint-access",
			Fields: fields_update_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpointAccess(ctx, input)
			},
		},
		"update-lakehouse-configuration": {
			Name:   "update-lakehouse-configuration",
			Fields: fields_update_lakehouse_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLakehouseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lakehouse_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLakehouseConfiguration(ctx, input)
			},
		},
		"update-namespace": {
			Name:   "update-namespace",
			Fields: fields_update_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNamespace(ctx, input)
			},
		},
		"update-scheduled-action": {
			Name:   "update-scheduled-action",
			Fields: fields_update_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScheduledAction(ctx, input)
			},
		},
		"update-snapshot": {
			Name:   "update-snapshot",
			Fields: fields_update_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSnapshot(ctx, input)
			},
		},
		"update-snapshot-copy-configuration": {
			Name:   "update-snapshot-copy-configuration",
			Fields: fields_update_snapshot_copy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSnapshotCopyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_snapshot_copy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSnapshotCopyConfiguration(ctx, input)
			},
		},
		"update-usage-limit": {
			Name:   "update-usage-limit",
			Fields: fields_update_usage_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUsageLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_usage_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUsageLimit(ctx, input)
			},
		},
		"update-workgroup": {
			Name:   "update-workgroup",
			Fields: fields_update_workgroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkgroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workgroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkgroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("redshiftserverless", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
