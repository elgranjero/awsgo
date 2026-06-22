package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/elasticache"
)

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_authorize_cache_security_group_ingress = []leanruntime.Field{
	{Name: "CacheSecurityGroupName", Flag: "cache-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupName", Flag: "ec2-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupOwnerId", Flag: "ec2-security-group-owner-id", Type: "*string", Required: true},
}

var fields_batch_apply_update_action = []leanruntime.Field{
	{Name: "CacheClusterIds", Flag: "cache-cluster-ids", Type: "[]string", Required: false},
	{Name: "ReplicationGroupIds", Flag: "replication-group-ids", Type: "[]string", Required: false},
	{Name: "ServiceUpdateName", Flag: "service-update-name", Type: "*string", Required: true},
}

var fields_batch_stop_update_action = []leanruntime.Field{
	{Name: "CacheClusterIds", Flag: "cache-cluster-ids", Type: "[]string", Required: false},
	{Name: "ReplicationGroupIds", Flag: "replication-group-ids", Type: "[]string", Required: false},
	{Name: "ServiceUpdateName", Flag: "service-update-name", Type: "*string", Required: true},
}

var fields_complete_migration = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
}

var fields_copy_serverless_cache_snapshot = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceServerlessCacheSnapshotName", Flag: "source-serverless-cache-snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetServerlessCacheSnapshotName", Flag: "target-serverless-cache-snapshot-name", Type: "*string", Required: true},
}

var fields_copy_snapshot = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceSnapshotName", Flag: "source-snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetBucket", Flag: "target-bucket", Type: "*string", Required: false},
	{Name: "TargetSnapshotName", Flag: "target-snapshot-name", Type: "*string", Required: true},
}

var fields_create_cache_cluster = []leanruntime.Field{
	{Name: "AZMode", Flag: "az-mode", Type: "types.AZMode", Required: false},
	{Name: "AuthToken", Flag: "auth-token", Type: "*string", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: true},
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: false},
	{Name: "CacheSecurityGroupNames", Flag: "cache-security-group-names", Type: "[]string", Required: false},
	{Name: "CacheSubnetGroupName", Flag: "cache-subnet-group-name", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "IpDiscovery", Flag: "ip-discovery", Type: "types.IpDiscovery", Required: false},
	{Name: "LogDeliveryConfigurations", Flag: "log-delivery-configurations", Type: "[]types.LogDeliveryConfigurationRequest", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "NotificationTopicArn", Flag: "notification-topic-arn", Type: "*string", Required: false},
	{Name: "NumCacheNodes", Flag: "num-cache-nodes", Type: "*int32", Required: false},
	{Name: "OutpostMode", Flag: "outpost-mode", Type: "types.OutpostMode", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredAvailabilityZone", Flag: "preferred-availability-zone", Type: "*string", Required: false},
	{Name: "PreferredAvailabilityZones", Flag: "preferred-availability-zones", Type: "[]string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PreferredOutpostArn", Flag: "preferred-outpost-arn", Type: "*string", Required: false},
	{Name: "PreferredOutpostArns", Flag: "preferred-outpost-arns", Type: "[]string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SnapshotArns", Flag: "snapshot-arns", Type: "[]string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SnapshotWindow", Flag: "snapshot-window", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransitEncryptionEnabled", Flag: "transit-encryption-enabled", Type: "*bool", Required: false},
}

var fields_create_cache_parameter_group = []leanruntime.Field{
	{Name: "CacheParameterGroupFamily", Flag: "cache-parameter-group-family", Type: "*string", Required: true},
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cache_security_group = []leanruntime.Field{
	{Name: "CacheSecurityGroupName", Flag: "cache-security-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cache_subnet_group = []leanruntime.Field{
	{Name: "CacheSubnetGroupDescription", Flag: "cache-subnet-group-description", Type: "*string", Required: true},
	{Name: "CacheSubnetGroupName", Flag: "cache-subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_global_replication_group = []leanruntime.Field{
	{Name: "GlobalReplicationGroupDescription", Flag: "global-replication-group-description", Type: "*string", Required: false},
	{Name: "GlobalReplicationGroupIdSuffix", Flag: "global-replication-group-id-suffix", Type: "*string", Required: true},
	{Name: "PrimaryReplicationGroupId", Flag: "primary-replication-group-id", Type: "*string", Required: true},
}

var fields_create_replication_group = []leanruntime.Field{
	{Name: "AtRestEncryptionEnabled", Flag: "at-rest-encryption-enabled", Type: "*bool", Required: false},
	{Name: "AuthToken", Flag: "auth-token", Type: "*string", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AutomaticFailoverEnabled", Flag: "automatic-failover-enabled", Type: "*bool", Required: false},
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: false},
	{Name: "CacheSecurityGroupNames", Flag: "cache-security-group-names", Type: "[]string", Required: false},
	{Name: "CacheSubnetGroupName", Flag: "cache-subnet-group-name", Type: "*string", Required: false},
	{Name: "ClusterMode", Flag: "cluster-mode", Type: "types.ClusterMode", Required: false},
	{Name: "DataTieringEnabled", Flag: "data-tiering-enabled", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: false},
	{Name: "IpDiscovery", Flag: "ip-discovery", Type: "types.IpDiscovery", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LogDeliveryConfigurations", Flag: "log-delivery-configurations", Type: "[]types.LogDeliveryConfigurationRequest", Required: false},
	{Name: "MultiAZEnabled", Flag: "multi-az-enabled", Type: "*bool", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "NodeGroupConfiguration", Flag: "node-group-configuration", Type: "[]types.NodeGroupConfiguration", Required: false},
	{Name: "NotificationTopicArn", Flag: "notification-topic-arn", Type: "*string", Required: false},
	{Name: "NumCacheClusters", Flag: "num-cache-clusters", Type: "*int32", Required: false},
	{Name: "NumNodeGroups", Flag: "num-node-groups", Type: "*int32", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PreferredCacheClusterAZs", Flag: "preferred-cache-cluster-azs", Type: "[]string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PrimaryClusterId", Flag: "primary-cluster-id", Type: "*string", Required: false},
	{Name: "ReplicasPerNodeGroup", Flag: "replicas-per-node-group", Type: "*int32", Required: false},
	{Name: "ReplicationGroupDescription", Flag: "replication-group-description", Type: "*string", Required: true},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "ServerlessCacheSnapshotName", Flag: "serverless-cache-snapshot-name", Type: "*string", Required: false},
	{Name: "SnapshotArns", Flag: "snapshot-arns", Type: "[]string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SnapshotWindow", Flag: "snapshot-window", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransitEncryptionEnabled", Flag: "transit-encryption-enabled", Type: "*bool", Required: false},
	{Name: "TransitEncryptionMode", Flag: "transit-encryption-mode", Type: "types.TransitEncryptionMode", Required: false},
	{Name: "UserGroupIds", Flag: "user-group-ids", Type: "[]string", Required: false},
}

var fields_create_serverless_cache = []leanruntime.Field{
	{Name: "CacheUsageLimits", Flag: "cache-usage-limits", Type: "*types.CacheUsageLimits", Required: false},
	{Name: "DailySnapshotTime", Flag: "daily-snapshot-time", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MajorEngineVersion", Flag: "major-engine-version", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "ServerlessCacheName", Flag: "serverless-cache-name", Type: "*string", Required: true},
	{Name: "SnapshotArnsToRestore", Flag: "snapshot-arns-to-restore", Type: "[]string", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserGroupId", Flag: "user-group-id", Type: "*string", Required: false},
}

var fields_create_serverless_cache_snapshot = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "ServerlessCacheName", Flag: "serverless-cache-name", Type: "*string", Required: true},
	{Name: "ServerlessCacheSnapshotName", Flag: "serverless-cache-snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_user = []leanruntime.Field{
	{Name: "AccessString", Flag: "access-string", Type: "*string", Required: true},
	{Name: "AuthenticationMode", Flag: "authentication-mode", Type: "*types.AuthenticationMode", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "NoPasswordRequired", Flag: "no-password-required", Type: "*bool", Required: false},
	{Name: "Passwords", Flag: "passwords", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_create_user_group = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserGroupId", Flag: "user-group-id", Type: "*string", Required: true},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: false},
}

var fields_decrease_node_groups_in_global_replication_group = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "GlobalNodeGroupsToRemove", Flag: "global-node-groups-to-remove", Type: "[]string", Required: false},
	{Name: "GlobalNodeGroupsToRetain", Flag: "global-node-groups-to-retain", Type: "[]string", Required: false},
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
	{Name: "NodeGroupCount", Flag: "node-group-count", Type: "*int32", Required: true},
}

var fields_decrease_replica_count = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "NewReplicaCount", Flag: "new-replica-count", Type: "*int32", Required: false},
	{Name: "ReplicaConfiguration", Flag: "replica-configuration", Type: "[]types.ConfigureShard", Required: false},
	{Name: "ReplicasToRemove", Flag: "replicas-to-remove", Type: "[]string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
}

var fields_delete_cache_cluster = []leanruntime.Field{
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: true},
	{Name: "FinalSnapshotIdentifier", Flag: "final-snapshot-identifier", Type: "*string", Required: false},
}

var fields_delete_cache_parameter_group = []leanruntime.Field{
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_cache_security_group = []leanruntime.Field{
	{Name: "CacheSecurityGroupName", Flag: "cache-security-group-name", Type: "*string", Required: true},
}

var fields_delete_cache_subnet_group = []leanruntime.Field{
	{Name: "CacheSubnetGroupName", Flag: "cache-subnet-group-name", Type: "*string", Required: true},
}

var fields_delete_global_replication_group = []leanruntime.Field{
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
	{Name: "RetainPrimaryReplicationGroup", Flag: "retain-primary-replication-group", Type: "*bool", Required: true},
}

var fields_delete_replication_group = []leanruntime.Field{
	{Name: "FinalSnapshotIdentifier", Flag: "final-snapshot-identifier", Type: "*string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
	{Name: "RetainPrimaryCluster", Flag: "retain-primary-cluster", Type: "*bool", Required: false},
}

var fields_delete_serverless_cache = []leanruntime.Field{
	{Name: "FinalSnapshotName", Flag: "final-snapshot-name", Type: "*string", Required: false},
	{Name: "ServerlessCacheName", Flag: "serverless-cache-name", Type: "*string", Required: true},
}

var fields_delete_serverless_cache_snapshot = []leanruntime.Field{
	{Name: "ServerlessCacheSnapshotName", Flag: "serverless-cache-snapshot-name", Type: "*string", Required: true},
}

var fields_delete_snapshot = []leanruntime.Field{
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_user_group = []leanruntime.Field{
	{Name: "UserGroupId", Flag: "user-group-id", Type: "*string", Required: true},
}

var fields_describe_cache_clusters = []leanruntime.Field{
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ShowCacheClustersNotInReplicationGroups", Flag: "show-cache-clusters-not-in-replication-groups", Type: "*bool", Required: false},
	{Name: "ShowCacheNodeInfo", Flag: "show-cache-node-info", Type: "*bool", Required: false},
}

var fields_describe_cache_engine_versions = []leanruntime.Field{
	{Name: "CacheParameterGroupFamily", Flag: "cache-parameter-group-family", Type: "*string", Required: false},
	{Name: "DefaultOnly", Flag: "default-only", Type: "*bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_cache_parameter_groups = []leanruntime.Field{
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_cache_parameters = []leanruntime.Field{
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_cache_security_groups = []leanruntime.Field{
	{Name: "CacheSecurityGroupName", Flag: "cache-security-group-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_cache_subnet_groups = []leanruntime.Field{
	{Name: "CacheSubnetGroupName", Flag: "cache-subnet-group-name", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
}

var fields_describe_engine_default_parameters = []leanruntime.Field{
	{Name: "CacheParameterGroupFamily", Flag: "cache-parameter-group-family", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
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

var fields_describe_global_replication_groups = []leanruntime.Field{
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ShowMemberInfo", Flag: "show-member-info", Type: "*bool", Required: false},
}

var fields_describe_replication_groups = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: false},
}

var fields_describe_reserved_cache_nodes = []leanruntime.Field{
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "*string", Required: false},
	{Name: "ProductDescription", Flag: "product-description", Type: "*string", Required: false},
	{Name: "ReservedCacheNodeId", Flag: "reserved-cache-node-id", Type: "*string", Required: false},
	{Name: "ReservedCacheNodesOfferingId", Flag: "reserved-cache-nodes-offering-id", Type: "*string", Required: false},
}

var fields_describe_reserved_cache_nodes_offerings = []leanruntime.Field{
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "*string", Required: false},
	{Name: "ProductDescription", Flag: "product-description", Type: "*string", Required: false},
	{Name: "ReservedCacheNodesOfferingId", Flag: "reserved-cache-nodes-offering-id", Type: "*string", Required: false},
}

var fields_describe_serverless_cache_snapshots = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerlessCacheName", Flag: "serverless-cache-name", Type: "*string", Required: false},
	{Name: "ServerlessCacheSnapshotName", Flag: "serverless-cache-snapshot-name", Type: "*string", Required: false},
	{Name: "SnapshotType", Flag: "snapshot-type", Type: "*string", Required: false},
}

var fields_describe_serverless_caches = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerlessCacheName", Flag: "serverless-cache-name", Type: "*string", Required: false},
}

var fields_describe_service_updates = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ServiceUpdateName", Flag: "service-update-name", Type: "*string", Required: false},
	{Name: "ServiceUpdateStatus", Flag: "service-update-status", Type: "[]types.ServiceUpdateStatus", Required: false},
}

var fields_describe_snapshots = []leanruntime.Field{
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: false},
	{Name: "ShowNodeGroupConfig", Flag: "show-node-group-config", Type: "*bool", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
	{Name: "SnapshotSource", Flag: "snapshot-source", Type: "*string", Required: false},
}

var fields_describe_update_actions = []leanruntime.Field{
	{Name: "CacheClusterIds", Flag: "cache-cluster-ids", Type: "[]string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "ReplicationGroupIds", Flag: "replication-group-ids", Type: "[]string", Required: false},
	{Name: "ServiceUpdateName", Flag: "service-update-name", Type: "*string", Required: false},
	{Name: "ServiceUpdateStatus", Flag: "service-update-status", Type: "[]types.ServiceUpdateStatus", Required: false},
	{Name: "ServiceUpdateTimeRange", Flag: "service-update-time-range", Type: "*types.TimeRangeFilter", Required: false},
	{Name: "ShowNodeLevelUpdateStatus", Flag: "show-node-level-update-status", Type: "*bool", Required: false},
	{Name: "UpdateActionStatus", Flag: "update-action-status", Type: "[]types.UpdateActionStatus", Required: false},
}

var fields_describe_user_groups = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "UserGroupId", Flag: "user-group-id", Type: "*string", Required: false},
}

var fields_describe_users = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_disassociate_global_replication_group = []leanruntime.Field{
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
	{Name: "ReplicationGroupRegion", Flag: "replication-group-region", Type: "*string", Required: true},
}

var fields_export_serverless_cache_snapshot = []leanruntime.Field{
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "ServerlessCacheSnapshotName", Flag: "serverless-cache-snapshot-name", Type: "*string", Required: true},
}

var fields_failover_global_replication_group = []leanruntime.Field{
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
	{Name: "PrimaryRegion", Flag: "primary-region", Type: "*string", Required: true},
	{Name: "PrimaryReplicationGroupId", Flag: "primary-replication-group-id", Type: "*string", Required: true},
}

var fields_increase_node_groups_in_global_replication_group = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
	{Name: "NodeGroupCount", Flag: "node-group-count", Type: "*int32", Required: true},
	{Name: "RegionalConfigurations", Flag: "regional-configurations", Type: "[]types.RegionalConfiguration", Required: false},
}

var fields_increase_replica_count = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "NewReplicaCount", Flag: "new-replica-count", Type: "*int32", Required: false},
	{Name: "ReplicaConfiguration", Flag: "replica-configuration", Type: "[]types.ConfigureShard", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
}

var fields_list_allowed_node_type_modifications = []leanruntime.Field{
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_modify_cache_cluster = []leanruntime.Field{
	{Name: "AZMode", Flag: "az-mode", Type: "types.AZMode", Required: false},
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "AuthToken", Flag: "auth-token", Type: "*string", Required: false},
	{Name: "AuthTokenUpdateStrategy", Flag: "auth-token-update-strategy", Type: "types.AuthTokenUpdateStrategyType", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: true},
	{Name: "CacheNodeIdsToRemove", Flag: "cache-node-ids-to-remove", Type: "[]string", Required: false},
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: false},
	{Name: "CacheSecurityGroupNames", Flag: "cache-security-group-names", Type: "[]string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "IpDiscovery", Flag: "ip-discovery", Type: "types.IpDiscovery", Required: false},
	{Name: "LogDeliveryConfigurations", Flag: "log-delivery-configurations", Type: "[]types.LogDeliveryConfigurationRequest", Required: false},
	{Name: "NewAvailabilityZones", Flag: "new-availability-zones", Type: "[]string", Required: false},
	{Name: "NotificationTopicArn", Flag: "notification-topic-arn", Type: "*string", Required: false},
	{Name: "NotificationTopicStatus", Flag: "notification-topic-status", Type: "*string", Required: false},
	{Name: "NumCacheNodes", Flag: "num-cache-nodes", Type: "*int32", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ScaleConfig", Flag: "scale-config", Type: "*types.ScaleConfig", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SnapshotWindow", Flag: "snapshot-window", Type: "*string", Required: false},
}

var fields_modify_cache_parameter_group = []leanruntime.Field{
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: true},
	{Name: "ParameterNameValues", Flag: "parameter-name-values", Type: "[]types.ParameterNameValue", Required: true},
}

var fields_modify_cache_subnet_group = []leanruntime.Field{
	{Name: "CacheSubnetGroupDescription", Flag: "cache-subnet-group-description", Type: "*string", Required: false},
	{Name: "CacheSubnetGroupName", Flag: "cache-subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
}

var fields_modify_global_replication_group = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "AutomaticFailoverEnabled", Flag: "automatic-failover-enabled", Type: "*bool", Required: false},
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "GlobalReplicationGroupDescription", Flag: "global-replication-group-description", Type: "*string", Required: false},
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
}

var fields_modify_replication_group = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "AuthToken", Flag: "auth-token", Type: "*string", Required: false},
	{Name: "AuthTokenUpdateStrategy", Flag: "auth-token-update-strategy", Type: "types.AuthTokenUpdateStrategyType", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "AutomaticFailoverEnabled", Flag: "automatic-failover-enabled", Type: "*bool", Required: false},
	{Name: "CacheNodeType", Flag: "cache-node-type", Type: "*string", Required: false},
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: false},
	{Name: "CacheSecurityGroupNames", Flag: "cache-security-group-names", Type: "[]string", Required: false},
	{Name: "ClusterMode", Flag: "cluster-mode", Type: "types.ClusterMode", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "IpDiscovery", Flag: "ip-discovery", Type: "types.IpDiscovery", Required: false},
	{Name: "LogDeliveryConfigurations", Flag: "log-delivery-configurations", Type: "[]types.LogDeliveryConfigurationRequest", Required: false},
	{Name: "MultiAZEnabled", Flag: "multi-az-enabled", Type: "*bool", Required: false},
	{Name: "NodeGroupId", Flag: "node-group-id", Type: "*string", Required: false},
	{Name: "NotificationTopicArn", Flag: "notification-topic-arn", Type: "*string", Required: false},
	{Name: "NotificationTopicStatus", Flag: "notification-topic-status", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PrimaryClusterId", Flag: "primary-cluster-id", Type: "*string", Required: false},
	{Name: "RemoveUserGroups", Flag: "remove-user-groups", Type: "*bool", Required: false},
	{Name: "ReplicationGroupDescription", Flag: "replication-group-description", Type: "*string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SnapshotWindow", Flag: "snapshot-window", Type: "*string", Required: false},
	{Name: "SnapshottingClusterId", Flag: "snapshotting-cluster-id", Type: "*string", Required: false},
	{Name: "TransitEncryptionEnabled", Flag: "transit-encryption-enabled", Type: "*bool", Required: false},
	{Name: "TransitEncryptionMode", Flag: "transit-encryption-mode", Type: "types.TransitEncryptionMode", Required: false},
	{Name: "UserGroupIdsToAdd", Flag: "user-group-ids-to-add", Type: "[]string", Required: false},
	{Name: "UserGroupIdsToRemove", Flag: "user-group-ids-to-remove", Type: "[]string", Required: false},
}

var fields_modify_replication_group_shard_configuration = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "NodeGroupCount", Flag: "node-group-count", Type: "*int32", Required: true},
	{Name: "NodeGroupsToRemove", Flag: "node-groups-to-remove", Type: "[]string", Required: false},
	{Name: "NodeGroupsToRetain", Flag: "node-groups-to-retain", Type: "[]string", Required: false},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
	{Name: "ReshardingConfiguration", Flag: "resharding-configuration", Type: "[]types.ReshardingConfiguration", Required: false},
}

var fields_modify_serverless_cache = []leanruntime.Field{
	{Name: "CacheUsageLimits", Flag: "cache-usage-limits", Type: "*types.CacheUsageLimits", Required: false},
	{Name: "DailySnapshotTime", Flag: "daily-snapshot-time", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "MajorEngineVersion", Flag: "major-engine-version", Type: "*string", Required: false},
	{Name: "RemoveUserGroup", Flag: "remove-user-group", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "ServerlessCacheName", Flag: "serverless-cache-name", Type: "*string", Required: true},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "UserGroupId", Flag: "user-group-id", Type: "*string", Required: false},
}

var fields_modify_user = []leanruntime.Field{
	{Name: "AccessString", Flag: "access-string", Type: "*string", Required: false},
	{Name: "AppendAccessString", Flag: "append-access-string", Type: "*string", Required: false},
	{Name: "AuthenticationMode", Flag: "authentication-mode", Type: "*types.AuthenticationMode", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "NoPasswordRequired", Flag: "no-password-required", Type: "*bool", Required: false},
	{Name: "Passwords", Flag: "passwords", Type: "[]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_modify_user_group = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "UserGroupId", Flag: "user-group-id", Type: "*string", Required: true},
	{Name: "UserIdsToAdd", Flag: "user-ids-to-add", Type: "[]string", Required: false},
	{Name: "UserIdsToRemove", Flag: "user-ids-to-remove", Type: "[]string", Required: false},
}

var fields_purchase_reserved_cache_nodes_offering = []leanruntime.Field{
	{Name: "CacheNodeCount", Flag: "cache-node-count", Type: "*int32", Required: false},
	{Name: "ReservedCacheNodeId", Flag: "reserved-cache-node-id", Type: "*string", Required: false},
	{Name: "ReservedCacheNodesOfferingId", Flag: "reserved-cache-nodes-offering-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_rebalance_slots_in_global_replication_group = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: true},
	{Name: "GlobalReplicationGroupId", Flag: "global-replication-group-id", Type: "*string", Required: true},
}

var fields_reboot_cache_cluster = []leanruntime.Field{
	{Name: "CacheClusterId", Flag: "cache-cluster-id", Type: "*string", Required: true},
	{Name: "CacheNodeIdsToReboot", Flag: "cache-node-ids-to-reboot", Type: "[]string", Required: true},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_reset_cache_parameter_group = []leanruntime.Field{
	{Name: "CacheParameterGroupName", Flag: "cache-parameter-group-name", Type: "*string", Required: true},
	{Name: "ParameterNameValues", Flag: "parameter-name-values", Type: "[]types.ParameterNameValue", Required: false},
	{Name: "ResetAllParameters", Flag: "reset-all-parameters", Type: "*bool", Required: false},
}

var fields_revoke_cache_security_group_ingress = []leanruntime.Field{
	{Name: "CacheSecurityGroupName", Flag: "cache-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupName", Flag: "ec2-security-group-name", Type: "*string", Required: true},
	{Name: "EC2SecurityGroupOwnerId", Flag: "ec2-security-group-owner-id", Type: "*string", Required: true},
}

var fields_start_migration = []leanruntime.Field{
	{Name: "CustomerNodeEndpointList", Flag: "customer-node-endpoint-list", Type: "[]types.CustomerNodeEndpoint", Required: true},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
}

var fields_test_failover = []leanruntime.Field{
	{Name: "NodeGroupId", Flag: "node-group-id", Type: "*string", Required: true},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
}

var fields_test_migration = []leanruntime.Field{
	{Name: "CustomerNodeEndpointList", Flag: "customer-node-endpoint-list", Type: "[]types.CustomerNodeEndpoint", Required: true},
	{Name: "ReplicationGroupId", Flag: "replication-group-id", Type: "*string", Required: true},
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
		"authorize-cache-security-group-ingress": {
			Name:   "authorize-cache-security-group-ingress",
			Fields: fields_authorize_cache_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeCacheSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_cache_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeCacheSecurityGroupIngress(ctx, input)
			},
		},
		"batch-apply-update-action": {
			Name:   "batch-apply-update-action",
			Fields: fields_batch_apply_update_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchApplyUpdateActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_apply_update_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchApplyUpdateAction(ctx, input)
			},
		},
		"batch-stop-update-action": {
			Name:   "batch-stop-update-action",
			Fields: fields_batch_stop_update_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchStopUpdateActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_stop_update_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchStopUpdateAction(ctx, input)
			},
		},
		"complete-migration": {
			Name:   "complete-migration",
			Fields: fields_complete_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteMigration(ctx, input)
			},
		},
		"copy-serverless-cache-snapshot": {
			Name:   "copy-serverless-cache-snapshot",
			Fields: fields_copy_serverless_cache_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyServerlessCacheSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_serverless_cache_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyServerlessCacheSnapshot(ctx, input)
			},
		},
		"copy-snapshot": {
			Name:   "copy-snapshot",
			Fields: fields_copy_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopySnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopySnapshot(ctx, input)
			},
		},
		"create-cache-cluster": {
			Name:   "create-cache-cluster",
			Fields: fields_create_cache_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCacheClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cache_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCacheCluster(ctx, input)
			},
		},
		"create-cache-parameter-group": {
			Name:   "create-cache-parameter-group",
			Fields: fields_create_cache_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCacheParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cache_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCacheParameterGroup(ctx, input)
			},
		},
		"create-cache-security-group": {
			Name:   "create-cache-security-group",
			Fields: fields_create_cache_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCacheSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cache_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCacheSecurityGroup(ctx, input)
			},
		},
		"create-cache-subnet-group": {
			Name:   "create-cache-subnet-group",
			Fields: fields_create_cache_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCacheSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cache_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCacheSubnetGroup(ctx, input)
			},
		},
		"create-global-replication-group": {
			Name:   "create-global-replication-group",
			Fields: fields_create_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlobalReplicationGroup(ctx, input)
			},
		},
		"create-replication-group": {
			Name:   "create-replication-group",
			Fields: fields_create_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationGroup(ctx, input)
			},
		},
		"create-serverless-cache": {
			Name:   "create-serverless-cache",
			Fields: fields_create_serverless_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServerlessCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_serverless_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServerlessCache(ctx, input)
			},
		},
		"create-serverless-cache-snapshot": {
			Name:   "create-serverless-cache-snapshot",
			Fields: fields_create_serverless_cache_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServerlessCacheSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_serverless_cache_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServerlessCacheSnapshot(ctx, input)
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
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"create-user-group": {
			Name:   "create-user-group",
			Fields: fields_create_user_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserGroup(ctx, input)
			},
		},
		"decrease-node-groups-in-global-replication-group": {
			Name:   "decrease-node-groups-in-global-replication-group",
			Fields: fields_decrease_node_groups_in_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecreaseNodeGroupsInGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decrease_node_groups_in_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DecreaseNodeGroupsInGlobalReplicationGroup(ctx, input)
			},
		},
		"decrease-replica-count": {
			Name:   "decrease-replica-count",
			Fields: fields_decrease_replica_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecreaseReplicaCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decrease_replica_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DecreaseReplicaCount(ctx, input)
			},
		},
		"delete-cache-cluster": {
			Name:   "delete-cache-cluster",
			Fields: fields_delete_cache_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCacheClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cache_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCacheCluster(ctx, input)
			},
		},
		"delete-cache-parameter-group": {
			Name:   "delete-cache-parameter-group",
			Fields: fields_delete_cache_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCacheParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cache_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCacheParameterGroup(ctx, input)
			},
		},
		"delete-cache-security-group": {
			Name:   "delete-cache-security-group",
			Fields: fields_delete_cache_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCacheSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cache_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCacheSecurityGroup(ctx, input)
			},
		},
		"delete-cache-subnet-group": {
			Name:   "delete-cache-subnet-group",
			Fields: fields_delete_cache_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCacheSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cache_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCacheSubnetGroup(ctx, input)
			},
		},
		"delete-global-replication-group": {
			Name:   "delete-global-replication-group",
			Fields: fields_delete_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlobalReplicationGroup(ctx, input)
			},
		},
		"delete-replication-group": {
			Name:   "delete-replication-group",
			Fields: fields_delete_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationGroup(ctx, input)
			},
		},
		"delete-serverless-cache": {
			Name:   "delete-serverless-cache",
			Fields: fields_delete_serverless_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServerlessCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_serverless_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServerlessCache(ctx, input)
			},
		},
		"delete-serverless-cache-snapshot": {
			Name:   "delete-serverless-cache-snapshot",
			Fields: fields_delete_serverless_cache_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServerlessCacheSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_serverless_cache_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServerlessCacheSnapshot(ctx, input)
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
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"delete-user-group": {
			Name:   "delete-user-group",
			Fields: fields_delete_user_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserGroup(ctx, input)
			},
		},
		"describe-cache-clusters": {
			Name:   "describe-cache-clusters",
			Fields: fields_describe_cache_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cache_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCacheClusters(ctx, input)
				}
				var results []*svc.DescribeCacheClustersOutput
				p := svc.NewDescribeCacheClustersPaginator(client, input)
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
		"describe-cache-engine-versions": {
			Name:   "describe-cache-engine-versions",
			Fields: fields_describe_cache_engine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheEngineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cache_engine_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCacheEngineVersions(ctx, input)
				}
				var results []*svc.DescribeCacheEngineVersionsOutput
				p := svc.NewDescribeCacheEngineVersionsPaginator(client, input)
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
		"describe-cache-parameter-groups": {
			Name:   "describe-cache-parameter-groups",
			Fields: fields_describe_cache_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheParameterGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cache_parameter_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCacheParameterGroups(ctx, input)
				}
				var results []*svc.DescribeCacheParameterGroupsOutput
				p := svc.NewDescribeCacheParameterGroupsPaginator(client, input)
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
		"describe-cache-parameters": {
			Name:   "describe-cache-parameters",
			Fields: fields_describe_cache_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cache_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCacheParameters(ctx, input)
				}
				var results []*svc.DescribeCacheParametersOutput
				p := svc.NewDescribeCacheParametersPaginator(client, input)
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
		"describe-cache-security-groups": {
			Name:   "describe-cache-security-groups",
			Fields: fields_describe_cache_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cache_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCacheSecurityGroups(ctx, input)
				}
				var results []*svc.DescribeCacheSecurityGroupsOutput
				p := svc.NewDescribeCacheSecurityGroupsPaginator(client, input)
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
		"describe-cache-subnet-groups": {
			Name:   "describe-cache-subnet-groups",
			Fields: fields_describe_cache_subnet_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheSubnetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cache_subnet_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCacheSubnetGroups(ctx, input)
				}
				var results []*svc.DescribeCacheSubnetGroupsOutput
				p := svc.NewDescribeCacheSubnetGroupsPaginator(client, input)
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
		"describe-global-replication-groups": {
			Name:   "describe-global-replication-groups",
			Fields: fields_describe_global_replication_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGlobalReplicationGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_global_replication_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGlobalReplicationGroups(ctx, input)
				}
				var results []*svc.DescribeGlobalReplicationGroupsOutput
				p := svc.NewDescribeGlobalReplicationGroupsPaginator(client, input)
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
		"describe-replication-groups": {
			Name:   "describe-replication-groups",
			Fields: fields_describe_replication_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationGroups(ctx, input)
				}
				var results []*svc.DescribeReplicationGroupsOutput
				p := svc.NewDescribeReplicationGroupsPaginator(client, input)
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
		"describe-reserved-cache-nodes": {
			Name:   "describe-reserved-cache-nodes",
			Fields: fields_describe_reserved_cache_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedCacheNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_cache_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedCacheNodes(ctx, input)
				}
				var results []*svc.DescribeReservedCacheNodesOutput
				p := svc.NewDescribeReservedCacheNodesPaginator(client, input)
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
		"describe-reserved-cache-nodes-offerings": {
			Name:   "describe-reserved-cache-nodes-offerings",
			Fields: fields_describe_reserved_cache_nodes_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedCacheNodesOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_cache_nodes_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedCacheNodesOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedCacheNodesOfferingsOutput
				p := svc.NewDescribeReservedCacheNodesOfferingsPaginator(client, input)
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
		"describe-serverless-cache-snapshots": {
			Name:   "describe-serverless-cache-snapshots",
			Fields: fields_describe_serverless_cache_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServerlessCacheSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_serverless_cache_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeServerlessCacheSnapshots(ctx, input)
				}
				var results []*svc.DescribeServerlessCacheSnapshotsOutput
				p := svc.NewDescribeServerlessCacheSnapshotsPaginator(client, input)
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
		"describe-serverless-caches": {
			Name:   "describe-serverless-caches",
			Fields: fields_describe_serverless_caches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServerlessCachesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_serverless_caches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeServerlessCaches(ctx, input)
				}
				var results []*svc.DescribeServerlessCachesOutput
				p := svc.NewDescribeServerlessCachesPaginator(client, input)
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
		"describe-service-updates": {
			Name:   "describe-service-updates",
			Fields: fields_describe_service_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceUpdatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_service_updates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeServiceUpdates(ctx, input)
				}
				var results []*svc.DescribeServiceUpdatesOutput
				p := svc.NewDescribeServiceUpdatesPaginator(client, input)
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
		"describe-snapshots": {
			Name:   "describe-snapshots",
			Fields: fields_describe_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSnapshots(ctx, input)
				}
				var results []*svc.DescribeSnapshotsOutput
				p := svc.NewDescribeSnapshotsPaginator(client, input)
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
		"describe-update-actions": {
			Name:   "describe-update-actions",
			Fields: fields_describe_update_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUpdateActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_update_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeUpdateActions(ctx, input)
				}
				var results []*svc.DescribeUpdateActionsOutput
				p := svc.NewDescribeUpdateActionsPaginator(client, input)
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
		"describe-user-groups": {
			Name:   "describe-user-groups",
			Fields: fields_describe_user_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_user_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeUserGroups(ctx, input)
				}
				var results []*svc.DescribeUserGroupsOutput
				p := svc.NewDescribeUserGroupsPaginator(client, input)
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
		"describe-users": {
			Name:   "describe-users",
			Fields: fields_describe_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeUsers(ctx, input)
				}
				var results []*svc.DescribeUsersOutput
				p := svc.NewDescribeUsersPaginator(client, input)
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
		"disassociate-global-replication-group": {
			Name:   "disassociate-global-replication-group",
			Fields: fields_disassociate_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateGlobalReplicationGroup(ctx, input)
			},
		},
		"export-serverless-cache-snapshot": {
			Name:   "export-serverless-cache-snapshot",
			Fields: fields_export_serverless_cache_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportServerlessCacheSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_serverless_cache_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportServerlessCacheSnapshot(ctx, input)
			},
		},
		"failover-global-replication-group": {
			Name:   "failover-global-replication-group",
			Fields: fields_failover_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FailoverGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_failover_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FailoverGlobalReplicationGroup(ctx, input)
			},
		},
		"increase-node-groups-in-global-replication-group": {
			Name:   "increase-node-groups-in-global-replication-group",
			Fields: fields_increase_node_groups_in_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IncreaseNodeGroupsInGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_increase_node_groups_in_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IncreaseNodeGroupsInGlobalReplicationGroup(ctx, input)
			},
		},
		"increase-replica-count": {
			Name:   "increase-replica-count",
			Fields: fields_increase_replica_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IncreaseReplicaCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_increase_replica_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IncreaseReplicaCount(ctx, input)
			},
		},
		"list-allowed-node-type-modifications": {
			Name:   "list-allowed-node-type-modifications",
			Fields: fields_list_allowed_node_type_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAllowedNodeTypeModificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_allowed_node_type_modifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAllowedNodeTypeModifications(ctx, input)
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
		"modify-cache-cluster": {
			Name:   "modify-cache-cluster",
			Fields: fields_modify_cache_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCacheClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cache_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCacheCluster(ctx, input)
			},
		},
		"modify-cache-parameter-group": {
			Name:   "modify-cache-parameter-group",
			Fields: fields_modify_cache_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCacheParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cache_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCacheParameterGroup(ctx, input)
			},
		},
		"modify-cache-subnet-group": {
			Name:   "modify-cache-subnet-group",
			Fields: fields_modify_cache_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCacheSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cache_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCacheSubnetGroup(ctx, input)
			},
		},
		"modify-global-replication-group": {
			Name:   "modify-global-replication-group",
			Fields: fields_modify_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyGlobalReplicationGroup(ctx, input)
			},
		},
		"modify-replication-group": {
			Name:   "modify-replication-group",
			Fields: fields_modify_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReplicationGroup(ctx, input)
			},
		},
		"modify-replication-group-shard-configuration": {
			Name:   "modify-replication-group-shard-configuration",
			Fields: fields_modify_replication_group_shard_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReplicationGroupShardConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_replication_group_shard_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReplicationGroupShardConfiguration(ctx, input)
			},
		},
		"modify-serverless-cache": {
			Name:   "modify-serverless-cache",
			Fields: fields_modify_serverless_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyServerlessCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_serverless_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyServerlessCache(ctx, input)
			},
		},
		"modify-user": {
			Name:   "modify-user",
			Fields: fields_modify_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyUser(ctx, input)
			},
		},
		"modify-user-group": {
			Name:   "modify-user-group",
			Fields: fields_modify_user_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyUserGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_user_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyUserGroup(ctx, input)
			},
		},
		"purchase-reserved-cache-nodes-offering": {
			Name:   "purchase-reserved-cache-nodes-offering",
			Fields: fields_purchase_reserved_cache_nodes_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedCacheNodesOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_cache_nodes_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedCacheNodesOffering(ctx, input)
			},
		},
		"rebalance-slots-in-global-replication-group": {
			Name:   "rebalance-slots-in-global-replication-group",
			Fields: fields_rebalance_slots_in_global_replication_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebalanceSlotsInGlobalReplicationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rebalance_slots_in_global_replication_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebalanceSlotsInGlobalReplicationGroup(ctx, input)
			},
		},
		"reboot-cache-cluster": {
			Name:   "reboot-cache-cluster",
			Fields: fields_reboot_cache_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootCacheClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_cache_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootCacheCluster(ctx, input)
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
		"reset-cache-parameter-group": {
			Name:   "reset-cache-parameter-group",
			Fields: fields_reset_cache_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetCacheParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_cache_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetCacheParameterGroup(ctx, input)
			},
		},
		"revoke-cache-security-group-ingress": {
			Name:   "revoke-cache-security-group-ingress",
			Fields: fields_revoke_cache_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeCacheSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_cache_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeCacheSecurityGroupIngress(ctx, input)
			},
		},
		"start-migration": {
			Name:   "start-migration",
			Fields: fields_start_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMigration(ctx, input)
			},
		},
		"test-failover": {
			Name:   "test-failover",
			Fields: fields_test_failover,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestFailoverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_failover, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestFailover(ctx, input)
			},
		},
		"test-migration": {
			Name:   "test-migration",
			Fields: fields_test_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestMigration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("elasticache", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
