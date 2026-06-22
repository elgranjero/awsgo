package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/docdbelastic"
)

var fields_apply_pending_maintenance_action = []leanruntime.Field{
	{Name: "ApplyAction", Flag: "apply-action", Type: "*string", Required: true},
	{Name: "ApplyOn", Flag: "apply-on", Type: "*string", Required: false},
	{Name: "OptInType", Flag: "opt-in-type", Type: "types.OptInType", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_copy_cluster_snapshot = []leanruntime.Field{
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetSnapshotName", Flag: "target-snapshot-name", Type: "*string", Required: true},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "AdminUserName", Flag: "admin-user-name", Type: "*string", Required: true},
	{Name: "AdminUserPassword", Flag: "admin-user-password", Type: "*string", Required: true},
	{Name: "AuthType", Flag: "auth-type", Type: "types.Auth", Required: true},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ShardCapacity", Flag: "shard-capacity", Type: "*int32", Required: true},
	{Name: "ShardCount", Flag: "shard-count", Type: "*int32", Required: true},
	{Name: "ShardInstanceCount", Flag: "shard-instance-count", Type: "*int32", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_create_cluster_snapshot = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_delete_cluster_snapshot = []leanruntime.Field{
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: true},
}

var fields_get_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_get_cluster_snapshot = []leanruntime.Field{
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: true},
}

var fields_get_pending_maintenance_action = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_cluster_snapshots = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SnapshotType", Flag: "snapshot-type", Type: "*string", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pending_maintenance_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_restore_cluster_from_snapshot = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "ShardCapacity", Flag: "shard-capacity", Type: "*int32", Required: false},
	{Name: "ShardInstanceCount", Flag: "shard-instance-count", Type: "*int32", Required: false},
	{Name: "SnapshotArn", Flag: "snapshot-arn", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_start_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_stop_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "AdminUserPassword", Flag: "admin-user-password", Type: "*string", Required: false},
	{Name: "AuthType", Flag: "auth-type", Type: "types.Auth", Required: false},
	{Name: "BackupRetentionPeriod", Flag: "backup-retention-period", Type: "*int32", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ShardCapacity", Flag: "shard-capacity", Type: "*int32", Required: false},
	{Name: "ShardCount", Flag: "shard-count", Type: "*int32", Required: false},
	{Name: "ShardInstanceCount", Flag: "shard-instance-count", Type: "*int32", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"get-cluster": {
			Name:   "get-cluster",
			Fields: fields_get_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCluster(ctx, input)
			},
		},
		"get-cluster-snapshot": {
			Name:   "get-cluster-snapshot",
			Fields: fields_get_cluster_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClusterSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cluster_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClusterSnapshot(ctx, input)
			},
		},
		"get-pending-maintenance-action": {
			Name:   "get-pending-maintenance-action",
			Fields: fields_get_pending_maintenance_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPendingMaintenanceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pending_maintenance_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPendingMaintenanceAction(ctx, input)
			},
		},
		"list-cluster-snapshots": {
			Name:   "list-cluster-snapshots",
			Fields: fields_list_cluster_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterSnapshots(ctx, input)
				}
				var results []*svc.ListClusterSnapshotsOutput
				p := svc.NewListClusterSnapshotsPaginator(client, input)
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
		"list-clusters": {
			Name:   "list-clusters",
			Fields: fields_list_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusters(ctx, input)
				}
				var results []*svc.ListClustersOutput
				p := svc.NewListClustersPaginator(client, input)
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
		"list-pending-maintenance-actions": {
			Name:   "list-pending-maintenance-actions",
			Fields: fields_list_pending_maintenance_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPendingMaintenanceActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pending_maintenance_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPendingMaintenanceActions(ctx, input)
				}
				var results []*svc.ListPendingMaintenanceActionsOutput
				p := svc.NewListPendingMaintenanceActionsPaginator(client, input)
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
		"restore-cluster-from-snapshot": {
			Name:   "restore-cluster-from-snapshot",
			Fields: fields_restore_cluster_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreClusterFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_cluster_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreClusterFromSnapshot(ctx, input)
			},
		},
		"start-cluster": {
			Name:   "start-cluster",
			Fields: fields_start_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCluster(ctx, input)
			},
		},
		"stop-cluster": {
			Name:   "stop-cluster",
			Fields: fields_stop_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCluster(ctx, input)
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
	}
	if err := leanruntime.Execute("docdbelastic", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
