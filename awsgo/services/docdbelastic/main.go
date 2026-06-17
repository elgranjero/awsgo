package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/docdbelastic/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"apply-pending-maintenance-action", "copy-cluster-snapshot", "create-cluster", "create-cluster-snapshot", "delete-cluster", "delete-cluster-snapshot", "get-cluster", "get-cluster-snapshot", "get-pending-maintenance-action", "list-cluster-snapshots", "list-clusters", "list-pending-maintenance-actions", "list-tags-for-resource", "restore-cluster-from-snapshot", "start-cluster", "stop-cluster", "tag-resource", "untag-resource", "update-cluster"},
		OperationSet: map[string]bool{"apply-pending-maintenance-action": true, "copy-cluster-snapshot": true, "create-cluster": true, "create-cluster-snapshot": true, "delete-cluster": true, "delete-cluster-snapshot": true, "get-cluster": true, "get-cluster-snapshot": true, "get-pending-maintenance-action": true, "list-cluster-snapshots": true, "list-clusters": true, "list-pending-maintenance-actions": true, "list-tags-for-resource": true, "restore-cluster-from-snapshot": true, "start-cluster": true, "stop-cluster": true, "tag-resource": true, "untag-resource": true, "update-cluster": true},
		OperationInputs: map[string][]string{
			"apply-pending-maintenance-action": {"ApplyAction", "ApplyOn", "OptInType", "ResourceArn"},
			"copy-cluster-snapshot":            {"CopyTags", "KmsKeyId", "SnapshotArn", "Tags", "TargetSnapshotName"},
			"create-cluster":                   {"AdminUserName", "AdminUserPassword", "AuthType", "BackupRetentionPeriod", "ClientToken", "ClusterName", "KmsKeyId", "PreferredBackupWindow", "PreferredMaintenanceWindow", "ShardCapacity", "ShardCount", "ShardInstanceCount", "SubnetIds", "Tags", "VpcSecurityGroupIds"},
			"create-cluster-snapshot":          {"ClusterArn", "SnapshotName", "Tags"},
			"delete-cluster":                   {"ClusterArn"},
			"delete-cluster-snapshot":          {"SnapshotArn"},
			"get-cluster":                      {"ClusterArn"},
			"get-cluster-snapshot":             {"SnapshotArn"},
			"get-pending-maintenance-action":   {"ResourceArn"},
			"list-cluster-snapshots":           {"ClusterArn", "MaxResults", "NextToken", "SnapshotType"},
			"list-clusters":                    {"MaxResults", "NextToken"},
			"list-pending-maintenance-actions": {"MaxResults", "NextToken"},
			"list-tags-for-resource":           {"ResourceArn"},
			"restore-cluster-from-snapshot":    {"ClusterName", "KmsKeyId", "ShardCapacity", "ShardInstanceCount", "SnapshotArn", "SubnetIds", "Tags", "VpcSecurityGroupIds"},
			"start-cluster":                    {"ClusterArn"},
			"stop-cluster":                     {"ClusterArn"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-cluster":                   {"AdminUserPassword", "AuthType", "BackupRetentionPeriod", "ClientToken", "ClusterArn", "PreferredBackupWindow", "PreferredMaintenanceWindow", "ShardCapacity", "ShardCount", "ShardInstanceCount", "SubnetIds", "VpcSecurityGroupIds"},
		},
		OperationInputTypes: map[string]map[string]string{
			"apply-pending-maintenance-action": {"ApplyAction": "*string", "ApplyOn": "*string", "OptInType": "types.OptInType", "ResourceArn": "*string"},
			"copy-cluster-snapshot":            {"CopyTags": "*bool", "KmsKeyId": "*string", "SnapshotArn": "*string", "Tags": "map[string]string", "TargetSnapshotName": "*string"},
			"create-cluster":                   {"AdminUserName": "*string", "AdminUserPassword": "*string", "AuthType": "types.Auth", "BackupRetentionPeriod": "*int32", "ClientToken": "*string", "ClusterName": "*string", "KmsKeyId": "*string", "PreferredBackupWindow": "*string", "PreferredMaintenanceWindow": "*string", "ShardCapacity": "*int32", "ShardCount": "*int32", "ShardInstanceCount": "*int32", "SubnetIds": "[]string", "Tags": "map[string]string", "VpcSecurityGroupIds": "[]string"},
			"create-cluster-snapshot":          {"ClusterArn": "*string", "SnapshotName": "*string", "Tags": "map[string]string"},
			"delete-cluster":                   {"ClusterArn": "*string"},
			"delete-cluster-snapshot":          {"SnapshotArn": "*string"},
			"get-cluster":                      {"ClusterArn": "*string"},
			"get-cluster-snapshot":             {"SnapshotArn": "*string"},
			"get-pending-maintenance-action":   {"ResourceArn": "*string"},
			"list-cluster-snapshots":           {"ClusterArn": "*string", "MaxResults": "*int32", "NextToken": "*string", "SnapshotType": "*string"},
			"list-clusters":                    {"MaxResults": "*int32", "NextToken": "*string"},
			"list-pending-maintenance-actions": {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":           {"ResourceArn": "*string"},
			"restore-cluster-from-snapshot":    {"ClusterName": "*string", "KmsKeyId": "*string", "ShardCapacity": "*int32", "ShardInstanceCount": "*int32", "SnapshotArn": "*string", "SubnetIds": "[]string", "Tags": "map[string]string", "VpcSecurityGroupIds": "[]string"},
			"start-cluster":                    {"ClusterArn": "*string"},
			"stop-cluster":                     {"ClusterArn": "*string"},
			"tag-resource":                     {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                   {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-cluster":                   {"AdminUserPassword": "*string", "AuthType": "types.Auth", "BackupRetentionPeriod": "*int32", "ClientToken": "*string", "ClusterArn": "*string", "PreferredBackupWindow": "*string", "PreferredMaintenanceWindow": "*string", "ShardCapacity": "*int32", "ShardCount": "*int32", "ShardInstanceCount": "*int32", "SubnetIds": "[]string", "VpcSecurityGroupIds": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"apply-pending-maintenance-action": {"ApplyAction", "OptInType", "ResourceArn"},
			"copy-cluster-snapshot":            {"SnapshotArn", "TargetSnapshotName"},
			"create-cluster":                   {"AdminUserName", "AdminUserPassword", "AuthType", "ClusterName", "ShardCapacity", "ShardCount"},
			"create-cluster-snapshot":          {"ClusterArn", "SnapshotName"},
			"delete-cluster":                   {"ClusterArn"},
			"delete-cluster-snapshot":          {"SnapshotArn"},
			"get-cluster":                      {"ClusterArn"},
			"get-cluster-snapshot":             {"SnapshotArn"},
			"get-pending-maintenance-action":   {"ResourceArn"},
			"list-cluster-snapshots":           {},
			"list-clusters":                    {},
			"list-pending-maintenance-actions": {},
			"list-tags-for-resource":           {"ResourceArn"},
			"restore-cluster-from-snapshot":    {"ClusterName", "SnapshotArn"},
			"start-cluster":                    {"ClusterArn"},
			"stop-cluster":                     {"ClusterArn"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-cluster":                   {"ClusterArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("docdbelastic", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
