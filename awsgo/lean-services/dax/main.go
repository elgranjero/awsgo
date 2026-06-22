package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/dax"
)

var fields_create_cluster = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "ClusterEndpointEncryptionType", Flag: "cluster-endpoint-encryption-type", Type: "types.ClusterEndpointEncryptionType", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "NodeType", Flag: "node-type", Type: "*string", Required: true},
	{Name: "NotificationTopicArn", Flag: "notification-topic-arn", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "ReplicationFactor", Flag: "replication-factor", Type: "int32", Required: true},
	{Name: "SSESpecification", Flag: "sse-specification", Type: "*types.SSESpecification", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_parameter_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
}

var fields_create_subnet_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
}

var fields_decrease_replication_factor = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NewReplicationFactor", Flag: "new-replication-factor", Type: "int32", Required: true},
	{Name: "NodeIdsToRemove", Flag: "node-ids-to-remove", Type: "[]string", Required: false},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_delete_parameter_group = []leanruntime.Field{
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
}

var fields_delete_subnet_group = []leanruntime.Field{
	{Name: "SubnetGroupName", Flag: "subnet-group-name", Type: "*string", Required: true},
}

var fields_describe_clusters = []leanruntime.Field{
	{Name: "ClusterNames", Flag: "cluster-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_default_parameters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
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

var fields_describe_parameter_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterGroupNames", Flag: "parameter-group-names", Type: "[]string", Required: false},
}

var fields_describe_parameters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
}

var fields_describe_subnet_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubnetGroupNames", Flag: "subnet-group-names", Type: "[]string", Required: false},
}

var fields_increase_replication_factor = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NewReplicationFactor", Flag: "new-replication-factor", Type: "int32", Required: true},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_reboot_node = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "NotificationTopicArn", Flag: "notification-topic-arn", Type: "*string", Required: false},
	{Name: "NotificationTopicStatus", Flag: "notification-topic-status", Type: "*string", Required: false},
	{Name: "ParameterGroupName", Flag: "parameter-group-name", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
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
		"decrease-replication-factor": {
			Name:   "decrease-replication-factor",
			Fields: fields_decrease_replication_factor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecreaseReplicationFactorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decrease_replication_factor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DecreaseReplicationFactor(ctx, input)
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
		"describe-clusters": {
			Name:   "describe-clusters",
			Fields: fields_describe_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClustersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_clusters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusters(ctx, input)
			},
		},
		"describe-default-parameters": {
			Name:   "describe-default-parameters",
			Fields: fields_describe_default_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDefaultParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_default_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDefaultParameters(ctx, input)
			},
		},
		"describe-events": {
			Name:   "describe-events",
			Fields: fields_describe_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEvents(ctx, input)
			},
		},
		"describe-parameter-groups": {
			Name:   "describe-parameter-groups",
			Fields: fields_describe_parameter_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeParameterGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_parameter_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeParameterGroups(ctx, input)
			},
		},
		"describe-parameters": {
			Name:   "describe-parameters",
			Fields: fields_describe_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeParameters(ctx, input)
			},
		},
		"describe-subnet-groups": {
			Name:   "describe-subnet-groups",
			Fields: fields_describe_subnet_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubnetGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_subnet_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSubnetGroups(ctx, input)
			},
		},
		"increase-replication-factor": {
			Name:   "increase-replication-factor",
			Fields: fields_increase_replication_factor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IncreaseReplicationFactorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_increase_replication_factor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IncreaseReplicationFactor(ctx, input)
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
		"reboot-node": {
			Name:   "reboot-node",
			Fields: fields_reboot_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootNode(ctx, input)
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
	}
	if err := leanruntime.Execute("dax", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
