package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/efs"
)

var fields_create_access_point = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "PosixUser", Flag: "posix-user", Type: "*types.PosixUser", Required: false},
	{Name: "RootDirectory", Flag: "root-directory", Type: "*types.RootDirectory", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_file_system = []leanruntime.Field{
	{Name: "AvailabilityZoneName", Flag: "availability-zone-name", Type: "*string", Required: false},
	{Name: "Backup", Flag: "backup", Type: "*bool", Required: false},
	{Name: "CreationToken", Flag: "creation-token", Type: "*string", Required: true},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PerformanceMode", Flag: "performance-mode", Type: "types.PerformanceMode", Required: false},
	{Name: "ProvisionedThroughputInMibps", Flag: "provisioned-throughput-in-mibps", Type: "*float64", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThroughputMode", Flag: "throughput-mode", Type: "types.ThroughputMode", Required: false},
}

var fields_create_mount_target = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "IpAddress", Flag: "ip-address", Type: "*string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "Ipv6Address", Flag: "ipv6-address", Type: "*string", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_create_replication_configuration = []leanruntime.Field{
	{Name: "Destinations", Flag: "destinations", Type: "[]types.DestinationToCreate", Required: true},
	{Name: "SourceFileSystemId", Flag: "source-file-system-id", Type: "*string", Required: true},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_delete_access_point = []leanruntime.Field{
	{Name: "AccessPointId", Flag: "access-point-id", Type: "*string", Required: true},
}

var fields_delete_file_system = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_delete_file_system_policy = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_delete_mount_target = []leanruntime.Field{
	{Name: "MountTargetId", Flag: "mount-target-id", Type: "*string", Required: true},
}

var fields_delete_replication_configuration = []leanruntime.Field{
	{Name: "DeletionMode", Flag: "deletion-mode", Type: "types.DeletionMode", Required: false},
	{Name: "SourceFileSystemId", Flag: "source-file-system-id", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_describe_access_points = []leanruntime.Field{
	{Name: "AccessPointId", Flag: "access-point-id", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_account_preferences = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_backup_policy = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_describe_file_system_policy = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_describe_file_systems = []leanruntime.Field{
	{Name: "CreationToken", Flag: "creation-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_describe_lifecycle_configuration = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_describe_mount_target_security_groups = []leanruntime.Field{
	{Name: "MountTargetId", Flag: "mount-target-id", Type: "*string", Required: true},
}

var fields_describe_mount_targets = []leanruntime.Field{
	{Name: "AccessPointId", Flag: "access-point-id", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "MountTargetId", Flag: "mount-target-id", Type: "*string", Required: false},
}

var fields_describe_replication_configurations = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_modify_mount_target_security_groups = []leanruntime.Field{
	{Name: "MountTargetId", Flag: "mount-target-id", Type: "*string", Required: true},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
}

var fields_put_account_preferences = []leanruntime.Field{
	{Name: "ResourceIdType", Flag: "resource-id-type", Type: "types.ResourceIdType", Required: true},
}

var fields_put_backup_policy = []leanruntime.Field{
	{Name: "BackupPolicy", Flag: "backup-policy", Type: "*types.BackupPolicy", Required: true},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_put_file_system_policy = []leanruntime.Field{
	{Name: "BypassPolicyLockoutSafetyCheck", Flag: "bypass-policy-lockout-safety-check", Type: "bool", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_lifecycle_configuration = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "LifecyclePolicies", Flag: "lifecycle-policies", Type: "[]types.LifecyclePolicy", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_file_system = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "ProvisionedThroughputInMibps", Flag: "provisioned-throughput-in-mibps", Type: "*float64", Required: false},
	{Name: "ThroughputMode", Flag: "throughput-mode", Type: "types.ThroughputMode", Required: false},
}

var fields_update_file_system_protection = []leanruntime.Field{
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "ReplicationOverwriteProtection", Flag: "replication-overwrite-protection", Type: "types.ReplicationOverwriteProtection", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-access-point": {
			Name:   "create-access-point",
			Fields: fields_create_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessPoint(ctx, input)
			},
		},
		"create-file-system": {
			Name:   "create-file-system",
			Fields: fields_create_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFileSystem(ctx, input)
			},
		},
		"create-mount-target": {
			Name:   "create-mount-target",
			Fields: fields_create_mount_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMountTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mount_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMountTarget(ctx, input)
			},
		},
		"create-replication-configuration": {
			Name:   "create-replication-configuration",
			Fields: fields_create_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationConfiguration(ctx, input)
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
		"delete-access-point": {
			Name:   "delete-access-point",
			Fields: fields_delete_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPoint(ctx, input)
			},
		},
		"delete-file-system": {
			Name:   "delete-file-system",
			Fields: fields_delete_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFileSystem(ctx, input)
			},
		},
		"delete-file-system-policy": {
			Name:   "delete-file-system-policy",
			Fields: fields_delete_file_system_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFileSystemPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_file_system_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFileSystemPolicy(ctx, input)
			},
		},
		"delete-mount-target": {
			Name:   "delete-mount-target",
			Fields: fields_delete_mount_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMountTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mount_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMountTarget(ctx, input)
			},
		},
		"delete-replication-configuration": {
			Name:   "delete-replication-configuration",
			Fields: fields_delete_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationConfiguration(ctx, input)
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
		"describe-access-points": {
			Name:   "describe-access-points",
			Fields: fields_describe_access_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccessPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_access_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAccessPoints(ctx, input)
				}
				var results []*svc.DescribeAccessPointsOutput
				p := svc.NewDescribeAccessPointsPaginator(client, input)
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
		"describe-account-preferences": {
			Name:   "describe-account-preferences",
			Fields: fields_describe_account_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountPreferences(ctx, input)
			},
		},
		"describe-backup-policy": {
			Name:   "describe-backup-policy",
			Fields: fields_describe_backup_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBackupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_backup_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBackupPolicy(ctx, input)
			},
		},
		"describe-file-system-policy": {
			Name:   "describe-file-system-policy",
			Fields: fields_describe_file_system_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFileSystemPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_file_system_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFileSystemPolicy(ctx, input)
			},
		},
		"describe-file-systems": {
			Name:   "describe-file-systems",
			Fields: fields_describe_file_systems,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFileSystemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_file_systems, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFileSystems(ctx, input)
				}
				var results []*svc.DescribeFileSystemsOutput
				p := svc.NewDescribeFileSystemsPaginator(client, input)
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
		"describe-lifecycle-configuration": {
			Name:   "describe-lifecycle-configuration",
			Fields: fields_describe_lifecycle_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLifecycleConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_lifecycle_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLifecycleConfiguration(ctx, input)
			},
		},
		"describe-mount-target-security-groups": {
			Name:   "describe-mount-target-security-groups",
			Fields: fields_describe_mount_target_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMountTargetSecurityGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_mount_target_security_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMountTargetSecurityGroups(ctx, input)
			},
		},
		"describe-mount-targets": {
			Name:   "describe-mount-targets",
			Fields: fields_describe_mount_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMountTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_mount_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMountTargets(ctx, input)
				}
				var results []*svc.DescribeMountTargetsOutput
				p := svc.NewDescribeMountTargetsPaginator(client, input)
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
		"describe-replication-configurations": {
			Name:   "describe-replication-configurations",
			Fields: fields_describe_replication_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationConfigurations(ctx, input)
				}
				var results []*svc.DescribeReplicationConfigurationsOutput
				p := svc.NewDescribeReplicationConfigurationsPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"modify-mount-target-security-groups": {
			Name:   "modify-mount-target-security-groups",
			Fields: fields_modify_mount_target_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyMountTargetSecurityGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_mount_target_security_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyMountTargetSecurityGroups(ctx, input)
			},
		},
		"put-account-preferences": {
			Name:   "put-account-preferences",
			Fields: fields_put_account_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountPreferences(ctx, input)
			},
		},
		"put-backup-policy": {
			Name:   "put-backup-policy",
			Fields: fields_put_backup_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBackupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_backup_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBackupPolicy(ctx, input)
			},
		},
		"put-file-system-policy": {
			Name:   "put-file-system-policy",
			Fields: fields_put_file_system_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFileSystemPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_file_system_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFileSystemPolicy(ctx, input)
			},
		},
		"put-lifecycle-configuration": {
			Name:   "put-lifecycle-configuration",
			Fields: fields_put_lifecycle_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLifecycleConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_lifecycle_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLifecycleConfiguration(ctx, input)
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
		"update-file-system": {
			Name:   "update-file-system",
			Fields: fields_update_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFileSystem(ctx, input)
			},
		},
		"update-file-system-protection": {
			Name:   "update-file-system-protection",
			Fields: fields_update_file_system_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFileSystemProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_file_system_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFileSystemProtection(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("efs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
