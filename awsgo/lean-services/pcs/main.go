package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pcs"
)

var fields_create_cluster = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Networking", Flag: "networking", Type: "*types.NetworkingRequest", Required: true},
	{Name: "Scheduler", Flag: "scheduler", Type: "*types.SchedulerRequest", Required: true},
	{Name: "Size", Flag: "size", Type: "types.Size", Required: true},
	{Name: "SlurmConfiguration", Flag: "slurm-configuration", Type: "*types.ClusterSlurmConfigurationRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_compute_node_group = []leanruntime.Field{
	{Name: "AmiId", Flag: "ami-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ComputeNodeGroupName", Flag: "compute-node-group-name", Type: "*string", Required: true},
	{Name: "CustomLaunchTemplate", Flag: "custom-launch-template", Type: "*types.CustomLaunchTemplate", Required: true},
	{Name: "IamInstanceProfileArn", Flag: "iam-instance-profile-arn", Type: "*string", Required: true},
	{Name: "InstanceConfigs", Flag: "instance-configs", Type: "[]types.InstanceConfig", Required: true},
	{Name: "PurchaseOption", Flag: "purchase-option", Type: "types.PurchaseOption", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfigurationRequest", Required: true},
	{Name: "SlurmConfiguration", Flag: "slurm-configuration", Type: "*types.ComputeNodeGroupSlurmConfigurationRequest", Required: false},
	{Name: "SpotOptions", Flag: "spot-options", Type: "*types.SpotOptions", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_queue = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ComputeNodeGroupConfigurations", Flag: "compute-node-group-configurations", Type: "[]types.ComputeNodeGroupConfiguration", Required: false},
	{Name: "QueueName", Flag: "queue-name", Type: "*string", Required: true},
	{Name: "SlurmConfiguration", Flag: "slurm-configuration", Type: "*types.QueueSlurmConfigurationRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_delete_compute_node_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ComputeNodeGroupIdentifier", Flag: "compute-node-group-identifier", Type: "*string", Required: true},
}

var fields_delete_queue = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "QueueIdentifier", Flag: "queue-identifier", Type: "*string", Required: true},
}

var fields_get_cluster = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
}

var fields_get_compute_node_group = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ComputeNodeGroupIdentifier", Flag: "compute-node-group-identifier", Type: "*string", Required: true},
}

var fields_get_queue = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "QueueIdentifier", Flag: "queue-identifier", Type: "*string", Required: true},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_compute_node_groups = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_queues = []leanruntime.Field{
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_compute_node_group_instance = []leanruntime.Field{
	{Name: "BootstrapId", Flag: "bootstrap-id", Type: "*string", Required: true},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
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
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "SlurmConfiguration", Flag: "slurm-configuration", Type: "*types.UpdateClusterSlurmConfigurationRequest", Required: false},
}

var fields_update_compute_node_group = []leanruntime.Field{
	{Name: "AmiId", Flag: "ami-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ComputeNodeGroupIdentifier", Flag: "compute-node-group-identifier", Type: "*string", Required: true},
	{Name: "CustomLaunchTemplate", Flag: "custom-launch-template", Type: "*types.CustomLaunchTemplate", Required: false},
	{Name: "IamInstanceProfileArn", Flag: "iam-instance-profile-arn", Type: "*string", Required: false},
	{Name: "PurchaseOption", Flag: "purchase-option", Type: "types.PurchaseOption", Required: false},
	{Name: "ScalingConfiguration", Flag: "scaling-configuration", Type: "*types.ScalingConfigurationRequest", Required: false},
	{Name: "SlurmConfiguration", Flag: "slurm-configuration", Type: "*types.UpdateComputeNodeGroupSlurmConfigurationRequest", Required: false},
	{Name: "SpotOptions", Flag: "spot-options", Type: "*types.SpotOptions", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
}

var fields_update_queue = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterIdentifier", Flag: "cluster-identifier", Type: "*string", Required: true},
	{Name: "ComputeNodeGroupConfigurations", Flag: "compute-node-group-configurations", Type: "[]types.ComputeNodeGroupConfiguration", Required: false},
	{Name: "QueueIdentifier", Flag: "queue-identifier", Type: "*string", Required: true},
	{Name: "SlurmConfiguration", Flag: "slurm-configuration", Type: "*types.UpdateQueueSlurmConfigurationRequest", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"create-compute-node-group": {
			Name:   "create-compute-node-group",
			Fields: fields_create_compute_node_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComputeNodeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_compute_node_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComputeNodeGroup(ctx, input)
			},
		},
		"create-queue": {
			Name:   "create-queue",
			Fields: fields_create_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueue(ctx, input)
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
		"delete-compute-node-group": {
			Name:   "delete-compute-node-group",
			Fields: fields_delete_compute_node_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComputeNodeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_compute_node_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComputeNodeGroup(ctx, input)
			},
		},
		"delete-queue": {
			Name:   "delete-queue",
			Fields: fields_delete_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueue(ctx, input)
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
		"get-compute-node-group": {
			Name:   "get-compute-node-group",
			Fields: fields_get_compute_node_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComputeNodeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compute_node_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComputeNodeGroup(ctx, input)
			},
		},
		"get-queue": {
			Name:   "get-queue",
			Fields: fields_get_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueue(ctx, input)
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
		"list-compute-node-groups": {
			Name:   "list-compute-node-groups",
			Fields: fields_list_compute_node_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComputeNodeGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compute_node_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComputeNodeGroups(ctx, input)
				}
				var results []*svc.ListComputeNodeGroupsOutput
				p := svc.NewListComputeNodeGroupsPaginator(client, input)
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
		"list-queues": {
			Name:   "list-queues",
			Fields: fields_list_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueues(ctx, input)
				}
				var results []*svc.ListQueuesOutput
				p := svc.NewListQueuesPaginator(client, input)
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
		"register-compute-node-group-instance": {
			Name:   "register-compute-node-group-instance",
			Fields: fields_register_compute_node_group_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterComputeNodeGroupInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_compute_node_group_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterComputeNodeGroupInstance(ctx, input)
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
		"update-compute-node-group": {
			Name:   "update-compute-node-group",
			Fields: fields_update_compute_node_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComputeNodeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_compute_node_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComputeNodeGroup(ctx, input)
			},
		},
		"update-queue": {
			Name:   "update-queue",
			Fields: fields_update_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueue(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pcs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
