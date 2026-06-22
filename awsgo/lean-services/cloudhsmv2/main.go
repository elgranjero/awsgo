package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

var fields_copy_backup_to_region = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
	{Name: "DestinationRegion", Flag: "destination-region", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "BackupRetentionPolicy", Flag: "backup-retention-policy", Type: "*types.BackupRetentionPolicy", Required: false},
	{Name: "HsmType", Flag: "hsm-type", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.ClusterMode", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "SourceBackupId", Flag: "source-backup-id", Type: "*string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
}

var fields_create_hsm = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: true},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "IpAddress", Flag: "ip-address", Type: "*string", Required: false},
}

var fields_delete_backup = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_delete_hsm = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "EniId", Flag: "eni-id", Type: "*string", Required: false},
	{Name: "EniIp", Flag: "eni-ip", Type: "*string", Required: false},
	{Name: "HsmId", Flag: "hsm-id", Type: "*string", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_describe_backups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Shared", Flag: "shared", Type: "*bool", Required: false},
	{Name: "SortAscending", Flag: "sort-ascending", Type: "*bool", Required: false},
}

var fields_describe_clusters = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_initialize_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "SignedCert", Flag: "signed-cert", Type: "*string", Required: true},
	{Name: "TrustAnchor", Flag: "trust-anchor", Type: "*string", Required: true},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_modify_backup_attributes = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
	{Name: "NeverExpires", Flag: "never-expires", Type: "*bool", Required: true},
}

var fields_modify_cluster = []leanruntime.Field{
	{Name: "BackupRetentionPolicy", Flag: "backup-retention-policy", Type: "*types.BackupRetentionPolicy", Required: false},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "HsmType", Flag: "hsm-type", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_restore_backup = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagKeyList", Flag: "tag-key-list", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"copy-backup-to-region": {
			Name:   "copy-backup-to-region",
			Fields: fields_copy_backup_to_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyBackupToRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_backup_to_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyBackupToRegion(ctx, input)
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
		"create-hsm": {
			Name:   "create-hsm",
			Fields: fields_create_hsm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHsmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hsm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHsm(ctx, input)
			},
		},
		"delete-backup": {
			Name:   "delete-backup",
			Fields: fields_delete_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackup(ctx, input)
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
		"delete-hsm": {
			Name:   "delete-hsm",
			Fields: fields_delete_hsm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHsmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hsm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHsm(ctx, input)
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
		"describe-backups": {
			Name:   "describe-backups",
			Fields: fields_describe_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBackupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_backups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBackups(ctx, input)
				}
				var results []*svc.DescribeBackupsOutput
				p := svc.NewDescribeBackupsPaginator(client, input)
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
		"initialize-cluster": {
			Name:   "initialize-cluster",
			Fields: fields_initialize_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitializeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initialize_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitializeCluster(ctx, input)
			},
		},
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTags(ctx, input)
				}
				var results []*svc.ListTagsOutput
				p := svc.NewListTagsPaginator(client, input)
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
		"modify-backup-attributes": {
			Name:   "modify-backup-attributes",
			Fields: fields_modify_backup_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyBackupAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_backup_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyBackupAttributes(ctx, input)
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
		"restore-backup": {
			Name:   "restore-backup",
			Fields: fields_restore_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreBackup(ctx, input)
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
	}
	if err := leanruntime.Execute("cloudhsmv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
