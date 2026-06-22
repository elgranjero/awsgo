package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb"
)

var fields_create_db_cluster = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: false},
	{Name: "DbInstanceType", Flag: "db-instance-type", Type: "types.DbInstanceType", Required: true},
	{Name: "DbParameterGroupIdentifier", Flag: "db-parameter-group-identifier", Type: "*string", Required: false},
	{Name: "DbStorageType", Flag: "db-storage-type", Type: "types.DbStorageType", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.ClusterDeploymentType", Required: false},
	{Name: "FailoverMode", Flag: "failover-mode", Type: "types.FailoverMode", Required: false},
	{Name: "LogDeliveryConfiguration", Flag: "log-delivery-configuration", Type: "*types.LogDeliveryConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Organization", Flag: "organization", Type: "*string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: true},
	{Name: "VpcSubnetIds", Flag: "vpc-subnet-ids", Type: "[]string", Required: true},
}

var fields_create_db_instance = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: false},
	{Name: "DbInstanceType", Flag: "db-instance-type", Type: "types.DbInstanceType", Required: true},
	{Name: "DbParameterGroupIdentifier", Flag: "db-parameter-group-identifier", Type: "*string", Required: false},
	{Name: "DbStorageType", Flag: "db-storage-type", Type: "types.DbStorageType", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.DeploymentType", Required: false},
	{Name: "LogDeliveryConfiguration", Flag: "log-delivery-configuration", Type: "*types.LogDeliveryConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Organization", Flag: "organization", Type: "*string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: true},
	{Name: "VpcSubnetIds", Flag: "vpc-subnet-ids", Type: "[]string", Required: true},
}

var fields_create_db_parameter_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "types.Parameters", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_db_cluster = []leanruntime.Field{
	{Name: "DbClusterId", Flag: "db-cluster-id", Type: "*string", Required: true},
}

var fields_delete_db_instance = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_db_cluster = []leanruntime.Field{
	{Name: "DbClusterId", Flag: "db-cluster-id", Type: "*string", Required: true},
}

var fields_get_db_instance = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_db_parameter_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_list_db_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_db_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_db_instances_for_cluster = []leanruntime.Field{
	{Name: "DbClusterId", Flag: "db-cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_db_parameter_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reboot_db_cluster = []leanruntime.Field{
	{Name: "DbClusterId", Flag: "db-cluster-id", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
}

var fields_reboot_db_instance = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_db_cluster = []leanruntime.Field{
	{Name: "DbClusterId", Flag: "db-cluster-id", Type: "*string", Required: true},
	{Name: "DbInstanceType", Flag: "db-instance-type", Type: "types.DbInstanceType", Required: false},
	{Name: "DbParameterGroupIdentifier", Flag: "db-parameter-group-identifier", Type: "*string", Required: false},
	{Name: "FailoverMode", Flag: "failover-mode", Type: "types.FailoverMode", Required: false},
	{Name: "LogDeliveryConfiguration", Flag: "log-delivery-configuration", Type: "*types.LogDeliveryConfiguration", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
}

var fields_update_db_instance = []leanruntime.Field{
	{Name: "AllocatedStorage", Flag: "allocated-storage", Type: "*int32", Required: false},
	{Name: "DbInstanceType", Flag: "db-instance-type", Type: "types.DbInstanceType", Required: false},
	{Name: "DbParameterGroupIdentifier", Flag: "db-parameter-group-identifier", Type: "*string", Required: false},
	{Name: "DbStorageType", Flag: "db-storage-type", Type: "types.DbStorageType", Required: false},
	{Name: "DeploymentType", Flag: "deployment-type", Type: "types.DeploymentType", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LogDeliveryConfiguration", Flag: "log-delivery-configuration", Type: "*types.LogDeliveryConfiguration", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-db-cluster": {
			Name:   "create-db-cluster",
			Fields: fields_create_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDbClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDbCluster(ctx, input)
			},
		},
		"create-db-instance": {
			Name:   "create-db-instance",
			Fields: fields_create_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDbInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDbInstance(ctx, input)
			},
		},
		"create-db-parameter-group": {
			Name:   "create-db-parameter-group",
			Fields: fields_create_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDbParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDbParameterGroup(ctx, input)
			},
		},
		"delete-db-cluster": {
			Name:   "delete-db-cluster",
			Fields: fields_delete_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDbClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDbCluster(ctx, input)
			},
		},
		"delete-db-instance": {
			Name:   "delete-db-instance",
			Fields: fields_delete_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDbInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDbInstance(ctx, input)
			},
		},
		"get-db-cluster": {
			Name:   "get-db-cluster",
			Fields: fields_get_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDbClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDbCluster(ctx, input)
			},
		},
		"get-db-instance": {
			Name:   "get-db-instance",
			Fields: fields_get_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDbInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDbInstance(ctx, input)
			},
		},
		"get-db-parameter-group": {
			Name:   "get-db-parameter-group",
			Fields: fields_get_db_parameter_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDbParameterGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_db_parameter_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDbParameterGroup(ctx, input)
			},
		},
		"list-db-clusters": {
			Name:   "list-db-clusters",
			Fields: fields_list_db_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbClusters(ctx, input)
				}
				var results []*svc.ListDbClustersOutput
				p := svc.NewListDbClustersPaginator(client, input)
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
		"list-db-instances": {
			Name:   "list-db-instances",
			Fields: fields_list_db_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbInstances(ctx, input)
				}
				var results []*svc.ListDbInstancesOutput
				p := svc.NewListDbInstancesPaginator(client, input)
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
		"list-db-instances-for-cluster": {
			Name:   "list-db-instances-for-cluster",
			Fields: fields_list_db_instances_for_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbInstancesForClusterInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_instances_for_cluster, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbInstancesForCluster(ctx, input)
				}
				var results []*svc.ListDbInstancesForClusterOutput
				p := svc.NewListDbInstancesForClusterPaginator(client, input)
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
		"list-db-parameter-groups": {
			Name:   "list-db-parameter-groups",
			Fields: fields_list_db_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDbParameterGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_db_parameter_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDbParameterGroups(ctx, input)
				}
				var results []*svc.ListDbParameterGroupsOutput
				p := svc.NewListDbParameterGroupsPaginator(client, input)
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
		"reboot-db-cluster": {
			Name:   "reboot-db-cluster",
			Fields: fields_reboot_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootDbClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootDbCluster(ctx, input)
			},
		},
		"reboot-db-instance": {
			Name:   "reboot-db-instance",
			Fields: fields_reboot_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootDbInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootDbInstance(ctx, input)
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
		"update-db-cluster": {
			Name:   "update-db-cluster",
			Fields: fields_update_db_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDbClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_db_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDbCluster(ctx, input)
			},
		},
		"update-db-instance": {
			Name:   "update-db-instance",
			Fields: fields_update_db_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDbInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_db_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDbInstance(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("timestreaminfluxdb", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
