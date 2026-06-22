package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kafka"
)

var fields_batch_associate_scram_secret = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "SecretArnList", Flag: "secret-arn-list", Type: "[]string", Required: true},
}

var fields_batch_disassociate_scram_secret = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "SecretArnList", Flag: "secret-arn-list", Type: "[]string", Required: true},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "BrokerNodeGroupInfo", Flag: "broker-node-group-info", Type: "*types.BrokerNodeGroupInfo", Required: true},
	{Name: "ClientAuthentication", Flag: "client-authentication", Type: "*types.ClientAuthentication", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ConfigurationInfo", Flag: "configuration-info", Type: "*types.ConfigurationInfo", Required: false},
	{Name: "EncryptionInfo", Flag: "encryption-info", Type: "*types.EncryptionInfo", Required: false},
	{Name: "EnhancedMonitoring", Flag: "enhanced-monitoring", Type: "types.EnhancedMonitoring", Required: false},
	{Name: "KafkaVersion", Flag: "kafka-version", Type: "*string", Required: true},
	{Name: "LoggingInfo", Flag: "logging-info", Type: "*types.LoggingInfo", Required: false},
	{Name: "NumberOfBrokerNodes", Flag: "number-of-broker-nodes", Type: "*int32", Required: true},
	{Name: "OpenMonitoring", Flag: "open-monitoring", Type: "*types.OpenMonitoringInfo", Required: false},
	{Name: "Rebalancing", Flag: "rebalancing", Type: "*types.Rebalancing", Required: false},
	{Name: "StorageMode", Flag: "storage-mode", Type: "types.StorageMode", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_cluster_v2 = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Provisioned", Flag: "provisioned", Type: "*types.ProvisionedRequest", Required: false},
	{Name: "Serverless", Flag: "serverless", Type: "*types.ServerlessRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KafkaVersions", Flag: "kafka-versions", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerProperties", Flag: "server-properties", Type: "[]byte", Required: true},
}

var fields_create_replicator = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KafkaClusters", Flag: "kafka-clusters", Type: "[]types.KafkaCluster", Required: true},
	{Name: "ReplicationInfoList", Flag: "replication-info-list", Type: "[]types.ReplicationInfo", Required: true},
	{Name: "ReplicatorName", Flag: "replicator-name", Type: "*string", Required: true},
	{Name: "ServiceExecutionRoleArn", Flag: "service-execution-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_topic = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "Configs", Flag: "configs", Type: "*string", Required: false},
	{Name: "PartitionCount", Flag: "partition-count", Type: "*int32", Required: true},
	{Name: "ReplicationFactor", Flag: "replication-factor", Type: "*int32", Required: true},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

var fields_create_vpc_connection = []leanruntime.Field{
	{Name: "Authentication", Flag: "authentication", Type: "*string", Required: true},
	{Name: "ClientSubnets", Flag: "client-subnets", Type: "[]string", Required: true},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetClusterArn", Flag: "target-cluster-arn", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: false},
}

var fields_delete_cluster_policy = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_delete_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_replicator = []leanruntime.Field{
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: false},
	{Name: "ReplicatorArn", Flag: "replicator-arn", Type: "*string", Required: true},
}

var fields_delete_topic = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

var fields_delete_vpc_connection = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_describe_cluster_operation = []leanruntime.Field{
	{Name: "ClusterOperationArn", Flag: "cluster-operation-arn", Type: "*string", Required: true},
}

var fields_describe_cluster_operation_v2 = []leanruntime.Field{
	{Name: "ClusterOperationArn", Flag: "cluster-operation-arn", Type: "*string", Required: true},
}

var fields_describe_cluster_v2 = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_describe_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_describe_configuration_revision = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*int64", Required: true},
}

var fields_describe_replicator = []leanruntime.Field{
	{Name: "ReplicatorArn", Flag: "replicator-arn", Type: "*string", Required: true},
}

var fields_describe_topic = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

var fields_describe_topic_partitions = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

var fields_describe_vpc_connection = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_bootstrap_brokers = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_get_cluster_policy = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_get_compatible_kafka_versions = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: false},
}

var fields_list_client_vpc_connections = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cluster_operations = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cluster_operations_v2 = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "ClusterNameFilter", Flag: "cluster-name-filter", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_clusters_v2 = []leanruntime.Field{
	{Name: "ClusterNameFilter", Flag: "cluster-name-filter", Type: "*string", Required: false},
	{Name: "ClusterTypeFilter", Flag: "cluster-type-filter", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_revisions = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kafka_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_nodes = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_replicators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReplicatorNameFilter", Flag: "replicator-name-filter", Type: "*string", Required: false},
}

var fields_list_scram_secrets = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_topics = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TopicNameFilter", Flag: "topic-name-filter", Type: "*string", Required: false},
}

var fields_list_vpc_connections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_cluster_policy = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_reboot_broker = []leanruntime.Field{
	{Name: "BrokerIds", Flag: "broker-ids", Type: "[]string", Required: true},
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_reject_client_vpc_connection = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "VpcConnectionArn", Flag: "vpc-connection-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_broker_count = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "TargetNumberOfBrokerNodes", Flag: "target-number-of-broker-nodes", Type: "*int32", Required: true},
}

var fields_update_broker_storage = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "TargetBrokerEBSVolumeInfo", Flag: "target-broker-ebs-volume-info", Type: "[]types.BrokerEBSVolumeInfo", Required: true},
}

var fields_update_broker_type = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "TargetInstanceType", Flag: "target-instance-type", Type: "*string", Required: true},
}

var fields_update_cluster_configuration = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "ConfigurationInfo", Flag: "configuration-info", Type: "*types.ConfigurationInfo", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
}

var fields_update_cluster_kafka_version = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "ConfigurationInfo", Flag: "configuration-info", Type: "*types.ConfigurationInfo", Required: false},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "TargetKafkaVersion", Flag: "target-kafka-version", Type: "*string", Required: true},
}

var fields_update_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ServerProperties", Flag: "server-properties", Type: "[]byte", Required: true},
}

var fields_update_connectivity = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "ConnectivityInfo", Flag: "connectivity-info", Type: "*types.ConnectivityInfo", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
}

var fields_update_monitoring = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "EnhancedMonitoring", Flag: "enhanced-monitoring", Type: "types.EnhancedMonitoring", Required: false},
	{Name: "LoggingInfo", Flag: "logging-info", Type: "*types.LoggingInfo", Required: false},
	{Name: "OpenMonitoring", Flag: "open-monitoring", Type: "*types.OpenMonitoringInfo", Required: false},
}

var fields_update_rebalancing = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "Rebalancing", Flag: "rebalancing", Type: "*types.Rebalancing", Required: true},
}

var fields_update_replication_info = []leanruntime.Field{
	{Name: "ConsumerGroupReplication", Flag: "consumer-group-replication", Type: "*types.ConsumerGroupReplicationUpdate", Required: false},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "ReplicatorArn", Flag: "replicator-arn", Type: "*string", Required: true},
	{Name: "SourceKafkaClusterArn", Flag: "source-kafka-cluster-arn", Type: "*string", Required: true},
	{Name: "TargetKafkaClusterArn", Flag: "target-kafka-cluster-arn", Type: "*string", Required: true},
	{Name: "TopicReplication", Flag: "topic-replication", Type: "*types.TopicReplicationUpdate", Required: false},
}

var fields_update_security = []leanruntime.Field{
	{Name: "ClientAuthentication", Flag: "client-authentication", Type: "*types.ClientAuthentication", Required: false},
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "EncryptionInfo", Flag: "encryption-info", Type: "*types.EncryptionInfo", Required: false},
}

var fields_update_storage = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "ProvisionedThroughput", Flag: "provisioned-throughput", Type: "*types.ProvisionedThroughput", Required: false},
	{Name: "StorageMode", Flag: "storage-mode", Type: "types.StorageMode", Required: false},
	{Name: "VolumeSizeGB", Flag: "volume-size-gb", Type: "*int32", Required: false},
}

var fields_update_topic = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "Configs", Flag: "configs", Type: "*string", Required: false},
	{Name: "PartitionCount", Flag: "partition-count", Type: "*int32", Required: false},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-associate-scram-secret": {
			Name:   "batch-associate-scram-secret",
			Fields: fields_batch_associate_scram_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateScramSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_scram_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateScramSecret(ctx, input)
			},
		},
		"batch-disassociate-scram-secret": {
			Name:   "batch-disassociate-scram-secret",
			Fields: fields_batch_disassociate_scram_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateScramSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_scram_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateScramSecret(ctx, input)
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
		"create-cluster-v2": {
			Name:   "create-cluster-v2",
			Fields: fields_create_cluster_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClusterV2(ctx, input)
			},
		},
		"create-configuration": {
			Name:   "create-configuration",
			Fields: fields_create_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguration(ctx, input)
			},
		},
		"create-replicator": {
			Name:   "create-replicator",
			Fields: fields_create_replicator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replicator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicator(ctx, input)
			},
		},
		"create-topic": {
			Name:   "create-topic",
			Fields: fields_create_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTopic(ctx, input)
			},
		},
		"create-vpc-connection": {
			Name:   "create-vpc-connection",
			Fields: fields_create_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcConnection(ctx, input)
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
		"delete-cluster-policy": {
			Name:   "delete-cluster-policy",
			Fields: fields_delete_cluster_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClusterPolicy(ctx, input)
			},
		},
		"delete-configuration": {
			Name:   "delete-configuration",
			Fields: fields_delete_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguration(ctx, input)
			},
		},
		"delete-replicator": {
			Name:   "delete-replicator",
			Fields: fields_delete_replicator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replicator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicator(ctx, input)
			},
		},
		"delete-topic": {
			Name:   "delete-topic",
			Fields: fields_delete_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTopic(ctx, input)
			},
		},
		"delete-vpc-connection": {
			Name:   "delete-vpc-connection",
			Fields: fields_delete_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcConnection(ctx, input)
			},
		},
		"describe-cluster": {
			Name:   "describe-cluster",
			Fields: fields_describe_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCluster(ctx, input)
			},
		},
		"describe-cluster-operation": {
			Name:   "describe-cluster-operation",
			Fields: fields_describe_cluster_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusterOperation(ctx, input)
			},
		},
		"describe-cluster-operation-v2": {
			Name:   "describe-cluster-operation-v2",
			Fields: fields_describe_cluster_operation_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterOperationV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster_operation_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusterOperationV2(ctx, input)
			},
		},
		"describe-cluster-v2": {
			Name:   "describe-cluster-v2",
			Fields: fields_describe_cluster_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusterV2(ctx, input)
			},
		},
		"describe-configuration": {
			Name:   "describe-configuration",
			Fields: fields_describe_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfiguration(ctx, input)
			},
		},
		"describe-configuration-revision": {
			Name:   "describe-configuration-revision",
			Fields: fields_describe_configuration_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationRevision(ctx, input)
			},
		},
		"describe-replicator": {
			Name:   "describe-replicator",
			Fields: fields_describe_replicator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_replicator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReplicator(ctx, input)
			},
		},
		"describe-topic": {
			Name:   "describe-topic",
			Fields: fields_describe_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTopic(ctx, input)
			},
		},
		"describe-topic-partitions": {
			Name:   "describe-topic-partitions",
			Fields: fields_describe_topic_partitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicPartitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_topic_partitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTopicPartitions(ctx, input)
				}
				var results []*svc.DescribeTopicPartitionsOutput
				p := svc.NewDescribeTopicPartitionsPaginator(client, input)
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
		"describe-vpc-connection": {
			Name:   "describe-vpc-connection",
			Fields: fields_describe_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcConnection(ctx, input)
			},
		},
		"get-bootstrap-brokers": {
			Name:   "get-bootstrap-brokers",
			Fields: fields_get_bootstrap_brokers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBootstrapBrokersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bootstrap_brokers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBootstrapBrokers(ctx, input)
			},
		},
		"get-cluster-policy": {
			Name:   "get-cluster-policy",
			Fields: fields_get_cluster_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClusterPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cluster_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClusterPolicy(ctx, input)
			},
		},
		"get-compatible-kafka-versions": {
			Name:   "get-compatible-kafka-versions",
			Fields: fields_get_compatible_kafka_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCompatibleKafkaVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compatible_kafka_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCompatibleKafkaVersions(ctx, input)
			},
		},
		"list-client-vpc-connections": {
			Name:   "list-client-vpc-connections",
			Fields: fields_list_client_vpc_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClientVpcConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_client_vpc_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClientVpcConnections(ctx, input)
				}
				var results []*svc.ListClientVpcConnectionsOutput
				p := svc.NewListClientVpcConnectionsPaginator(client, input)
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
		"list-cluster-operations": {
			Name:   "list-cluster-operations",
			Fields: fields_list_cluster_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterOperations(ctx, input)
				}
				var results []*svc.ListClusterOperationsOutput
				p := svc.NewListClusterOperationsPaginator(client, input)
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
		"list-cluster-operations-v2": {
			Name:   "list-cluster-operations-v2",
			Fields: fields_list_cluster_operations_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterOperationsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_operations_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterOperationsV2(ctx, input)
				}
				var results []*svc.ListClusterOperationsV2Output
				p := svc.NewListClusterOperationsV2Paginator(client, input)
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
		"list-clusters-v2": {
			Name:   "list-clusters-v2",
			Fields: fields_list_clusters_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClustersV2(ctx, input)
				}
				var results []*svc.ListClustersV2Output
				p := svc.NewListClustersV2Paginator(client, input)
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
		"list-configuration-revisions": {
			Name:   "list-configuration-revisions",
			Fields: fields_list_configuration_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationRevisions(ctx, input)
				}
				var results []*svc.ListConfigurationRevisionsOutput
				p := svc.NewListConfigurationRevisionsPaginator(client, input)
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
		"list-configurations": {
			Name:   "list-configurations",
			Fields: fields_list_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurations(ctx, input)
				}
				var results []*svc.ListConfigurationsOutput
				p := svc.NewListConfigurationsPaginator(client, input)
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
		"list-kafka-versions": {
			Name:   "list-kafka-versions",
			Fields: fields_list_kafka_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKafkaVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kafka_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKafkaVersions(ctx, input)
				}
				var results []*svc.ListKafkaVersionsOutput
				p := svc.NewListKafkaVersionsPaginator(client, input)
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
		"list-nodes": {
			Name:   "list-nodes",
			Fields: fields_list_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodes(ctx, input)
				}
				var results []*svc.ListNodesOutput
				p := svc.NewListNodesPaginator(client, input)
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
		"list-replicators": {
			Name:   "list-replicators",
			Fields: fields_list_replicators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReplicatorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_replicators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReplicators(ctx, input)
				}
				var results []*svc.ListReplicatorsOutput
				p := svc.NewListReplicatorsPaginator(client, input)
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
		"list-scram-secrets": {
			Name:   "list-scram-secrets",
			Fields: fields_list_scram_secrets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScramSecretsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scram_secrets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScramSecrets(ctx, input)
				}
				var results []*svc.ListScramSecretsOutput
				p := svc.NewListScramSecretsPaginator(client, input)
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
		"list-topics": {
			Name:   "list-topics",
			Fields: fields_list_topics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_topics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTopics(ctx, input)
				}
				var results []*svc.ListTopicsOutput
				p := svc.NewListTopicsPaginator(client, input)
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
		"list-vpc-connections": {
			Name:   "list-vpc-connections",
			Fields: fields_list_vpc_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vpc_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVpcConnections(ctx, input)
				}
				var results []*svc.ListVpcConnectionsOutput
				p := svc.NewListVpcConnectionsPaginator(client, input)
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
		"put-cluster-policy": {
			Name:   "put-cluster-policy",
			Fields: fields_put_cluster_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutClusterPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_cluster_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutClusterPolicy(ctx, input)
			},
		},
		"reboot-broker": {
			Name:   "reboot-broker",
			Fields: fields_reboot_broker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootBrokerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_broker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootBroker(ctx, input)
			},
		},
		"reject-client-vpc-connection": {
			Name:   "reject-client-vpc-connection",
			Fields: fields_reject_client_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectClientVpcConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_client_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectClientVpcConnection(ctx, input)
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
		"update-broker-count": {
			Name:   "update-broker-count",
			Fields: fields_update_broker_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrokerCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_broker_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrokerCount(ctx, input)
			},
		},
		"update-broker-storage": {
			Name:   "update-broker-storage",
			Fields: fields_update_broker_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrokerStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_broker_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrokerStorage(ctx, input)
			},
		},
		"update-broker-type": {
			Name:   "update-broker-type",
			Fields: fields_update_broker_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrokerTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_broker_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrokerType(ctx, input)
			},
		},
		"update-cluster-configuration": {
			Name:   "update-cluster-configuration",
			Fields: fields_update_cluster_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterConfiguration(ctx, input)
			},
		},
		"update-cluster-kafka-version": {
			Name:   "update-cluster-kafka-version",
			Fields: fields_update_cluster_kafka_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterKafkaVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_kafka_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterKafkaVersion(ctx, input)
			},
		},
		"update-configuration": {
			Name:   "update-configuration",
			Fields: fields_update_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguration(ctx, input)
			},
		},
		"update-connectivity": {
			Name:   "update-connectivity",
			Fields: fields_update_connectivity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectivityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connectivity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectivity(ctx, input)
			},
		},
		"update-monitoring": {
			Name:   "update-monitoring",
			Fields: fields_update_monitoring,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMonitoringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_monitoring, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMonitoring(ctx, input)
			},
		},
		"update-rebalancing": {
			Name:   "update-rebalancing",
			Fields: fields_update_rebalancing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRebalancingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rebalancing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRebalancing(ctx, input)
			},
		},
		"update-replication-info": {
			Name:   "update-replication-info",
			Fields: fields_update_replication_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReplicationInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_replication_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReplicationInfo(ctx, input)
			},
		},
		"update-security": {
			Name:   "update-security",
			Fields: fields_update_security,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurity(ctx, input)
			},
		},
		"update-storage": {
			Name:   "update-storage",
			Fields: fields_update_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStorage(ctx, input)
			},
		},
		"update-topic": {
			Name:   "update-topic",
			Fields: fields_update_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTopic(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kafka", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
