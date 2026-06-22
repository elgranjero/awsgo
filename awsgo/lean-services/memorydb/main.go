package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/memorydb"
)

var fields_batch_update_cluster = []leanruntime.Field{
	{Name: "ClusterNames", Flag: "cluster-names", Type: "[]string", Required: true},
	{Name: "ServiceUpdate", Flag: "service-update", Type: "*types.ServiceUpdateRequest", Required: false},
}

var fields_copy_snapshot = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceSnapshotName", Flag: "source-snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetBucket", Flag: "target-bucket", Type: "*string", Required: false},
	{Name: "TargetSnapshotName", Flag: "target-snapshot-name", Type: "*string", Required: true},
}

var fields_create_acl = []leanruntime.Field{
	{Name: "ACLName", Flag: "acl-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserNames", Flag: "user-names", Type: "[]string", Required: false},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "ACLName", Flag: "acl-name", Type: "*string", Required: true},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "DataTiering", Flag: "data-tiering", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "IpDiscovery", Flag: "ip-discovery", Type: "types.IpDiscovery", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MaintenanceWindow", Flag: "maintenance-window", Type: "*string", Required: false},
	{Name: "MultiRegionClusterName", Flag: "multi-region-cluster-name", Type: "*string", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: true},
	{Name: "NumReplicasPerShard", Flag: "num-replicas-per-shard", Type: "*int32", Required: false},
	{Name: "NumShards", Flag: "num-shards", Type: "*int32", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SnapshotArns", Flag: "snapshot-arns", Type: "[]string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SnapshotWindow", Flag: "snapshot-window", Type: "*string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: false},
	{Name: "TLSEnabled", Flag: "tls-enabled", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_multi_region_cluster = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "MultiRegionClusterNameSuffix", Flag: "multi-region-cluster-name-suffix", Type: "*string", Required: true},
	{Name: "MultiRegionParameterGroupName", Flag: "multi-region-parameter-group-name", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: true},
	{Name: "NumShards", Flag: "num-shards", Type: "*int32", Required: false},
	{Name: "TLSEnabled", Flag: "tls-enabled", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_parameter_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Family", Flag: "family", Type: "*string", Required: true},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_subnet_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_user = []leanruntime.Field{
	{Name: "AccessString", Flag: "access-string", Type: "*string", Required: true},
	{Name: "AuthenticationMode", Flag: "authentication-mode", Type: "*types.AuthenticationMode", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_acl = []leanruntime.Field{
	{Name: "ACLName", Flag: "acl-name", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "FinalSnapshotName", Flag: "final-snapshot-name", Type: "*string", Required: false},
	{Name: "MultiRegionClusterName", Flag: "multi-region-cluster-name", Type: "*string", Required: false},
}

var fields_delete_multi_region_cluster = []leanruntime.Field{
	{Name: "MultiRegionClusterName", Flag: "multi-region-cluster-name", Type: "*string", Required: true},
}

var fields_delete_parameter_group = []leanruntime.Field{
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_snapshot = []leanruntime.Field{
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_delete_subnet_group = []leanruntime.Field{
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_describe_acls = []leanruntime.Field{
	{Name: "ACLName", Flag: "acl-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_clusters = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShowShardDetails", Flag: "show-shard-details", Type: "*bool", Required: false},
}

var fields_describe_engine_versions = []leanruntime.Field{
	{Name: "DefaultOnly", Flag: "default-only", Type: "bool", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterGroupFamily", Flag: "parameter-group-family", Type: "*string", Required: false},
}

var fields_describe_events = []leanruntime.Field{
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceName", Flag: "source-name", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.SourceType", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_multi_region_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MultiRegionClusterName", Flag: "multi-region-cluster-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShowClusterDetails", Flag: "show-cluster-details", Type: "*bool", Required: false},
}

var fields_describe_multi_region_parameter_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MultiRegionParameterGroupName", Flag: "multi-region-parameter-group-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_multi_region_parameters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MultiRegionParameterGroupName", Flag: "multi-region-parameter-group-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_parameter_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: false},
}

var fields_describe_parameters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
}

var fields_describe_reserved_nodes = []leanruntime.Field{
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "*string", Required: false},
	{Name: "ReservationId", Flag: "reservation-id", Type: "*string", Required: false},
	{Name: "ReservedNodesOfferingId", Flag: "reserved-nodes-offering-id", Type: "*string", Required: false},
}

var fields_describe_reserved_nodes_offerings = []leanruntime.Field{
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "*string", Required: false},
	{Name: "ReservedNodesOfferingId", Flag: "reserved-nodes-offering-id", Type: "*string", Required: false},
}

var fields_describe_service_updates = []leanruntime.Field{
	{Name: "ClusterNames", Flag: "cluster-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceUpdateName", Flag: "service-update-name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.ServiceUpdateStatus", Required: false},
}

var fields_describe_snapshots = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShowDetail", Flag: "show-detail", Type: "*bool", Required: false},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_subnet_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: false},
}

var fields_describe_users = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_failover_shard = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ShardName", Flag: "shard-name", Type: "*string", Required: true},
}

var fields_list_allowed_multi_region_cluster_updates = []leanruntime.Field{
	{Name: "MultiRegionClusterName", Flag: "multi-region-cluster-name", Type: "*string", Required: true},
}

var fields_list_allowed_node_type_updates = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_purchase_reserved_nodes_offering = []leanruntime.Field{
	{Name: "NodeCount", Flag: "node-count", Type: "*int32", Required: false},
	{Name: "ReservationId", Flag: "reservation-id", Type: "*string", Required: false},
	{Name: "ReservedNodesOfferingId", Flag: "reserved-nodes-offering-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_reset_parameter_group = []leanruntime.Field{
	{Name: "AllParameters", Flag: "all-parameters", Type: "bool", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "ParameterNames", Flag: "parameter-names", Type: "[]string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_acl = []leanruntime.Field{
	{Name: "ACLName", Flag: "acl-name", Type: "*string", Required: true},
	{Name: "UserNamesToAdd", Flag: "user-names-to-add", Type: "[]string", Required: false},
	{Name: "UserNamesToRemove", Flag: "user-names-to-remove", Type: "[]string", Required: false},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "ACLName", Flag: "acl-name", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "IpDiscovery", Flag: "ip-discovery", Type: "types.IpDiscovery", Required: false},
	{Name: "MaintenanceWindow", Flag: "maintenance-window", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: false},
	{Name: "ReplicaConfiguration", Flag: "replica-configuration", Type: "*types.ReplicaConfigurationRequest", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "ShardConfiguration", Flag: "shard-configuration", Type: "*types.ShardConfigurationRequest", Required: false},
	{Name: "SnapshotRetentionLimit", Flag: "snapshot-retention-limit", Type: "*int32", Required: false},
	{Name: "SnapshotWindow", Flag: "snapshot-window", Type: "*string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SnsTopicStatus", Flag: "sns-topic-status", Type: "*string", Required: false},
}

var fields_update_multi_region_cluster = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "MultiRegionClusterName", Flag: "multi-region-cluster-name", Type: "*string", Required: true},
	{Name: "MultiRegionParameterGroupName", Flag: "multi-region-parameter-group-name", Type: "*string", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: false},
	{Name: "ShardConfiguration", Flag: "shard-configuration", Type: "*types.ShardConfigurationRequest", Required: false},
	{Name: "UpdateStrategy", Flag: "update-strategy", Type: "types.UpdateStrategy", Required: false},
}

var fields_update_parameter_group = []leanruntime.Field{
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "ParameterNameValues", Flag: "parameter-name-values", Type: "[]types.ParameterNameValue", Required: true},
}

var fields_update_subnet_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "AccessString", Flag: "access-string", Type: "*string", Required: false},
	{Name: "AuthenticationMode", Flag: "authentication-mode", Type: "*types.AuthenticationMode", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-update-cluster": {
			Name:   "batch-update-cluster",
			Fields: fields_batch_update_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateCluster(ctx, input)
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
		"create-acl": {
			Name:   "create-acl",
			Fields: fields_create_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateACL(ctx, input)
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
		"create-multi-region-cluster": {
			Name:   "create-multi-region-cluster",
			Fields: fields_create_multi_region_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultiRegionClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multi_region_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultiRegionCluster(ctx, input)
			},
		},
		"create-parameter-group": {
			Name:   "create-parameter-group",
			Fields: fields_create_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateParameterGroup(ctx, input)
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
		"create-subnet-group": {
			Name:   "create-subnet-group",
			Fields: fields_create_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubnetGroup(ctx, input)
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
		"delete-acl": {
			Name:   "delete-acl",
			Fields: fields_delete_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteACL(ctx, input)
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
		"delete-multi-region-cluster": {
			Name:   "delete-multi-region-cluster",
			Fields: fields_delete_multi_region_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMultiRegionClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_multi_region_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMultiRegionCluster(ctx, input)
			},
		},
		"delete-parameter-group": {
			Name:   "delete-parameter-group",
			Fields: fields_delete_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteParameterGroup(ctx, input)
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
		"delete-subnet-group": {
			Name:   "delete-subnet-group",
			Fields: fields_delete_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubnetGroup(ctx, input)
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
		"describe-acls": {
			Name:   "describe-acls",
			Fields: fields_describe_acls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeACLsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_acls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeACLs(ctx, input)
				}
				var results []*svc.DescribeACLsOutput
				p := svc.NewDescribeACLsPaginator(client, input)
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
		"describe-multi-region-clusters": {
			Name:   "describe-multi-region-clusters",
			Fields: fields_describe_multi_region_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMultiRegionClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_multi_region_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMultiRegionClusters(ctx, input)
				}
				var results []*svc.DescribeMultiRegionClustersOutput
				p := svc.NewDescribeMultiRegionClustersPaginator(client, input)
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
		"describe-multi-region-parameter-groups": {
			Name:   "describe-multi-region-parameter-groups",
			Fields: fields_describe_multi_region_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMultiRegionParameterGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_multi_region_parameter_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMultiRegionParameterGroups(ctx, input)
			},
		},
		"describe-multi-region-parameters": {
			Name:   "describe-multi-region-parameters",
			Fields: fields_describe_multi_region_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMultiRegionParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_multi_region_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMultiRegionParameters(ctx, input)
			},
		},
		"describe-parameter-groups": {
			Name:   "describe-parameter-groups",
			Fields: fields_describe_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeParameterGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_parameter_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeParameterGroups(ctx, input)
				}
				var results []*svc.DescribeParameterGroupsOutput
				p := svc.NewDescribeParameterGroupsPaginator(client, input)
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
		"describe-parameters": {
			Name:   "describe-parameters",
			Fields: fields_describe_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeParameters(ctx, input)
				}
				var results []*svc.DescribeParametersOutput
				p := svc.NewDescribeParametersPaginator(client, input)
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
		"describe-reserved-nodes-offerings": {
			Name:   "describe-reserved-nodes-offerings",
			Fields: fields_describe_reserved_nodes_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedNodesOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_nodes_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedNodesOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedNodesOfferingsOutput
				p := svc.NewDescribeReservedNodesOfferingsPaginator(client, input)
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
		"describe-subnet-groups": {
			Name:   "describe-subnet-groups",
			Fields: fields_describe_subnet_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubnetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_subnet_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSubnetGroups(ctx, input)
				}
				var results []*svc.DescribeSubnetGroupsOutput
				p := svc.NewDescribeSubnetGroupsPaginator(client, input)
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
		"failover-shard": {
			Name:   "failover-shard",
			Fields: fields_failover_shard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FailoverShardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_failover_shard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FailoverShard(ctx, input)
			},
		},
		"list-allowed-multi-region-cluster-updates": {
			Name:   "list-allowed-multi-region-cluster-updates",
			Fields: fields_list_allowed_multi_region_cluster_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAllowedMultiRegionClusterUpdatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_allowed_multi_region_cluster_updates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAllowedMultiRegionClusterUpdates(ctx, input)
			},
		},
		"list-allowed-node-type-updates": {
			Name:   "list-allowed-node-type-updates",
			Fields: fields_list_allowed_node_type_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAllowedNodeTypeUpdatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_allowed_node_type_updates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAllowedNodeTypeUpdates(ctx, input)
			},
		},
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTags(ctx, input)
			},
		},
		"purchase-reserved-nodes-offering": {
			Name:   "purchase-reserved-nodes-offering",
			Fields: fields_purchase_reserved_nodes_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedNodesOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_nodes_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedNodesOffering(ctx, input)
			},
		},
		"reset-parameter-group": {
			Name:   "reset-parameter-group",
			Fields: fields_reset_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetParameterGroup(ctx, input)
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
		"update-acl": {
			Name:   "update-acl",
			Fields: fields_update_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateACL(ctx, input)
			},
		},
		"update-cluster": {
			Name:   "update-cluster",
			Fields: fields_update_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCluster(ctx, input)
			},
		},
		"update-multi-region-cluster": {
			Name:   "update-multi-region-cluster",
			Fields: fields_update_multi_region_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMultiRegionClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_multi_region_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMultiRegionCluster(ctx, input)
			},
		},
		"update-parameter-group": {
			Name:   "update-parameter-group",
			Fields: fields_update_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateParameterGroup(ctx, input)
			},
		},
		"update-subnet-group": {
			Name:   "update-subnet-group",
			Fields: fields_update_subnet_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubnetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subnet_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubnetGroup(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("memorydb", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
