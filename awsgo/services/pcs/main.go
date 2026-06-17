package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pcs/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-cluster", "create-compute-node-group", "create-queue", "delete-cluster", "delete-compute-node-group", "delete-queue", "get-cluster", "get-compute-node-group", "get-queue", "list-clusters", "list-compute-node-groups", "list-queues", "list-tags-for-resource", "register-compute-node-group-instance", "tag-resource", "untag-resource", "update-cluster", "update-compute-node-group", "update-queue"},
		OperationSet: map[string]bool{"create-cluster": true, "create-compute-node-group": true, "create-queue": true, "delete-cluster": true, "delete-compute-node-group": true, "delete-queue": true, "get-cluster": true, "get-compute-node-group": true, "get-queue": true, "list-clusters": true, "list-compute-node-groups": true, "list-queues": true, "list-tags-for-resource": true, "register-compute-node-group-instance": true, "tag-resource": true, "untag-resource": true, "update-cluster": true, "update-compute-node-group": true, "update-queue": true},
		OperationInputs: map[string][]string{
			"create-cluster":                       {"ClientToken", "ClusterName", "Networking", "Scheduler", "Size", "SlurmConfiguration", "Tags"},
			"create-compute-node-group":            {"AmiId", "ClientToken", "ClusterIdentifier", "ComputeNodeGroupName", "CustomLaunchTemplate", "IamInstanceProfileArn", "InstanceConfigs", "PurchaseOption", "ScalingConfiguration", "SlurmConfiguration", "SpotOptions", "SubnetIds", "Tags"},
			"create-queue":                         {"ClientToken", "ClusterIdentifier", "ComputeNodeGroupConfigurations", "QueueName", "SlurmConfiguration", "Tags"},
			"delete-cluster":                       {"ClientToken", "ClusterIdentifier"},
			"delete-compute-node-group":            {"ClientToken", "ClusterIdentifier", "ComputeNodeGroupIdentifier"},
			"delete-queue":                         {"ClientToken", "ClusterIdentifier", "QueueIdentifier"},
			"get-cluster":                          {"ClusterIdentifier"},
			"get-compute-node-group":               {"ClusterIdentifier", "ComputeNodeGroupIdentifier"},
			"get-queue":                            {"ClusterIdentifier", "QueueIdentifier"},
			"list-clusters":                        {"MaxResults", "NextToken"},
			"list-compute-node-groups":             {"ClusterIdentifier", "MaxResults", "NextToken"},
			"list-queues":                          {"ClusterIdentifier", "MaxResults", "NextToken"},
			"list-tags-for-resource":               {"ResourceArn"},
			"register-compute-node-group-instance": {"BootstrapId", "ClusterIdentifier"},
			"tag-resource":                         {"ResourceArn", "Tags"},
			"untag-resource":                       {"ResourceArn", "TagKeys"},
			"update-cluster":                       {"ClientToken", "ClusterIdentifier", "SlurmConfiguration"},
			"update-compute-node-group":            {"AmiId", "ClientToken", "ClusterIdentifier", "ComputeNodeGroupIdentifier", "CustomLaunchTemplate", "IamInstanceProfileArn", "PurchaseOption", "ScalingConfiguration", "SlurmConfiguration", "SpotOptions", "SubnetIds"},
			"update-queue":                         {"ClientToken", "ClusterIdentifier", "ComputeNodeGroupConfigurations", "QueueIdentifier", "SlurmConfiguration"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-cluster":                       {"ClientToken": "*string", "ClusterName": "*string", "Networking": "*types.NetworkingRequest", "Scheduler": "*types.SchedulerRequest", "Size": "types.Size", "SlurmConfiguration": "*types.ClusterSlurmConfigurationRequest", "Tags": "map[string]string"},
			"create-compute-node-group":            {"AmiId": "*string", "ClientToken": "*string", "ClusterIdentifier": "*string", "ComputeNodeGroupName": "*string", "CustomLaunchTemplate": "*types.CustomLaunchTemplate", "IamInstanceProfileArn": "*string", "InstanceConfigs": "[]types.InstanceConfig", "PurchaseOption": "types.PurchaseOption", "ScalingConfiguration": "*types.ScalingConfigurationRequest", "SlurmConfiguration": "*types.ComputeNodeGroupSlurmConfigurationRequest", "SpotOptions": "*types.SpotOptions", "SubnetIds": "[]string", "Tags": "map[string]string"},
			"create-queue":                         {"ClientToken": "*string", "ClusterIdentifier": "*string", "ComputeNodeGroupConfigurations": "[]types.ComputeNodeGroupConfiguration", "QueueName": "*string", "SlurmConfiguration": "*types.QueueSlurmConfigurationRequest", "Tags": "map[string]string"},
			"delete-cluster":                       {"ClientToken": "*string", "ClusterIdentifier": "*string"},
			"delete-compute-node-group":            {"ClientToken": "*string", "ClusterIdentifier": "*string", "ComputeNodeGroupIdentifier": "*string"},
			"delete-queue":                         {"ClientToken": "*string", "ClusterIdentifier": "*string", "QueueIdentifier": "*string"},
			"get-cluster":                          {"ClusterIdentifier": "*string"},
			"get-compute-node-group":               {"ClusterIdentifier": "*string", "ComputeNodeGroupIdentifier": "*string"},
			"get-queue":                            {"ClusterIdentifier": "*string", "QueueIdentifier": "*string"},
			"list-clusters":                        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-compute-node-groups":             {"ClusterIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-queues":                          {"ClusterIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":               {"ResourceArn": "*string"},
			"register-compute-node-group-instance": {"BootstrapId": "*string", "ClusterIdentifier": "*string"},
			"tag-resource":                         {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                       {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-cluster":                       {"ClientToken": "*string", "ClusterIdentifier": "*string", "SlurmConfiguration": "*types.UpdateClusterSlurmConfigurationRequest"},
			"update-compute-node-group":            {"AmiId": "*string", "ClientToken": "*string", "ClusterIdentifier": "*string", "ComputeNodeGroupIdentifier": "*string", "CustomLaunchTemplate": "*types.CustomLaunchTemplate", "IamInstanceProfileArn": "*string", "PurchaseOption": "types.PurchaseOption", "ScalingConfiguration": "*types.ScalingConfigurationRequest", "SlurmConfiguration": "*types.UpdateComputeNodeGroupSlurmConfigurationRequest", "SpotOptions": "*types.SpotOptions", "SubnetIds": "[]string"},
			"update-queue":                         {"ClientToken": "*string", "ClusterIdentifier": "*string", "ComputeNodeGroupConfigurations": "[]types.ComputeNodeGroupConfiguration", "QueueIdentifier": "*string", "SlurmConfiguration": "*types.UpdateQueueSlurmConfigurationRequest"},
		},
		OperationInputRequired: map[string][]string{
			"create-cluster":                       {"ClusterName", "Networking", "Scheduler", "Size"},
			"create-compute-node-group":            {"ClusterIdentifier", "ComputeNodeGroupName", "CustomLaunchTemplate", "IamInstanceProfileArn", "InstanceConfigs", "ScalingConfiguration", "SubnetIds"},
			"create-queue":                         {"ClusterIdentifier", "QueueName"},
			"delete-cluster":                       {"ClusterIdentifier"},
			"delete-compute-node-group":            {"ClusterIdentifier", "ComputeNodeGroupIdentifier"},
			"delete-queue":                         {"ClusterIdentifier", "QueueIdentifier"},
			"get-cluster":                          {"ClusterIdentifier"},
			"get-compute-node-group":               {"ClusterIdentifier", "ComputeNodeGroupIdentifier"},
			"get-queue":                            {"ClusterIdentifier", "QueueIdentifier"},
			"list-clusters":                        {},
			"list-compute-node-groups":             {"ClusterIdentifier"},
			"list-queues":                          {"ClusterIdentifier"},
			"list-tags-for-resource":               {"ResourceArn"},
			"register-compute-node-group-instance": {"BootstrapId", "ClusterIdentifier"},
			"tag-resource":                         {"ResourceArn", "Tags"},
			"untag-resource":                       {"ResourceArn", "TagKeys"},
			"update-cluster":                       {"ClusterIdentifier"},
			"update-compute-node-group":            {"ClusterIdentifier", "ComputeNodeGroupIdentifier"},
			"update-queue":                         {"ClusterIdentifier", "QueueIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pcs", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
