package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var _kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "AWS kafka CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := kafka.NewFromConfig(cfg)
		if _kafkaBatchAssociateScramSecret {
			kafka_BatchAssociateScramSecret(cfg, client)
			return
		}
		if _kafkaBatchDisassociateScramSecret {
			kafka_BatchDisassociateScramSecret(cfg, client)
			return
		}
		if _kafkaCreateCluster {
			kafka_CreateCluster(cfg, client)
			return
		}
		if _kafkaCreateClusterV2 {
			kafka_CreateClusterV2(cfg, client)
			return
		}
		if _kafkaCreateConfiguration {
			kafka_CreateConfiguration(cfg, client)
			return
		}
		if _kafkaCreateReplicator {
			kafka_CreateReplicator(cfg, client)
			return
		}
		if _kafkaCreateTopic {
			kafka_CreateTopic(cfg, client)
			return
		}
		if _kafkaCreateVpcConnection {
			kafka_CreateVpcConnection(cfg, client)
			return
		}
		if _kafkaDeleteCluster {
			kafka_DeleteCluster(cfg, client)
			return
		}
		if _kafkaDeleteClusterPolicy {
			kafka_DeleteClusterPolicy(cfg, client)
			return
		}
		if _kafkaDeleteConfiguration {
			kafka_DeleteConfiguration(cfg, client)
			return
		}
		if _kafkaDeleteReplicator {
			kafka_DeleteReplicator(cfg, client)
			return
		}
		if _kafkaDeleteTopic {
			kafka_DeleteTopic(cfg, client)
			return
		}
		if _kafkaDeleteVpcConnection {
			kafka_DeleteVpcConnection(cfg, client)
			return
		}
		if _kafkaDescribeCluster {
			kafka_DescribeCluster(cfg, client)
			return
		}
		if _kafkaDescribeClusterOperation {
			kafka_DescribeClusterOperation(cfg, client)
			return
		}
		if _kafkaDescribeClusterOperationV2 {
			kafka_DescribeClusterOperationV2(cfg, client)
			return
		}
		if _kafkaDescribeClusterV2 {
			kafka_DescribeClusterV2(cfg, client)
			return
		}
		if _kafkaDescribeConfiguration {
			kafka_DescribeConfiguration(cfg, client)
			return
		}
		if _kafkaDescribeConfigurationRevision {
			kafka_DescribeConfigurationRevision(cfg, client)
			return
		}
		if _kafkaDescribeReplicator {
			kafka_DescribeReplicator(cfg, client)
			return
		}
		if _kafkaDescribeTopic {
			kafka_DescribeTopic(cfg, client)
			return
		}
		if _kafkaDescribeTopicPartitions {
			kafka_DescribeTopicPartitions(cfg, client)
			return
		}
		if _kafkaDescribeVpcConnection {
			kafka_DescribeVpcConnection(cfg, client)
			return
		}
		if _kafkaGetBootstrapBrokers {
			kafka_GetBootstrapBrokers(cfg, client)
			return
		}
		if _kafkaGetClusterPolicy {
			kafka_GetClusterPolicy(cfg, client)
			return
		}
		if _kafkaGetCompatibleKafkaVersions {
			kafka_GetCompatibleKafkaVersions(cfg, client)
			return
		}
		if _kafkaListClientVpcConnections {
			kafka_ListClientVpcConnections(cfg, client)
			return
		}
		if _kafkaListClusterOperations {
			kafka_ListClusterOperations(cfg, client)
			return
		}
		if _kafkaListClusterOperationsV2 {
			kafka_ListClusterOperationsV2(cfg, client)
			return
		}
		if _kafkaListClusters {
			kafka_ListClusters(cfg, client)
			return
		}
		if _kafkaListClustersV2 {
			kafka_ListClustersV2(cfg, client)
			return
		}
		if _kafkaListConfigurationRevisions {
			kafka_ListConfigurationRevisions(cfg, client)
			return
		}
		if _kafkaListConfigurations {
			kafka_ListConfigurations(cfg, client)
			return
		}
		if _kafkaListKafkaVersions {
			kafka_ListKafkaVersions(cfg, client)
			return
		}
		if _kafkaListNodes {
			kafka_ListNodes(cfg, client)
			return
		}
		if _kafkaListReplicators {
			kafka_ListReplicators(cfg, client)
			return
		}
		if _kafkaListScramSecrets {
			kafka_ListScramSecrets(cfg, client)
			return
		}
		if _kafkaListTagsForResource {
			kafka_ListTagsForResource(cfg, client)
			return
		}
		if _kafkaListTopics {
			kafka_ListTopics(cfg, client)
			return
		}
		if _kafkaListVpcConnections {
			kafka_ListVpcConnections(cfg, client)
			return
		}
		if _kafkaPutClusterPolicy {
			kafka_PutClusterPolicy(cfg, client)
			return
		}
		if _kafkaRebootBroker {
			kafka_RebootBroker(cfg, client)
			return
		}
		if _kafkaRejectClientVpcConnection {
			kafka_RejectClientVpcConnection(cfg, client)
			return
		}
		if _kafkaTagResource {
			kafka_TagResource(cfg, client)
			return
		}
		if _kafkaUntagResource {
			kafka_UntagResource(cfg, client)
			return
		}
		if _kafkaUpdateBrokerCount {
			kafka_UpdateBrokerCount(cfg, client)
			return
		}
		if _kafkaUpdateBrokerStorage {
			kafka_UpdateBrokerStorage(cfg, client)
			return
		}
		if _kafkaUpdateBrokerType {
			kafka_UpdateBrokerType(cfg, client)
			return
		}
		if _kafkaUpdateClusterConfiguration {
			kafka_UpdateClusterConfiguration(cfg, client)
			return
		}
		if _kafkaUpdateClusterKafkaVersion {
			kafka_UpdateClusterKafkaVersion(cfg, client)
			return
		}
		if _kafkaUpdateConfiguration {
			kafka_UpdateConfiguration(cfg, client)
			return
		}
		if _kafkaUpdateConnectivity {
			kafka_UpdateConnectivity(cfg, client)
			return
		}
		if _kafkaUpdateMonitoring {
			kafka_UpdateMonitoring(cfg, client)
			return
		}
		if _kafkaUpdateRebalancing {
			kafka_UpdateRebalancing(cfg, client)
			return
		}
		if _kafkaUpdateReplicationInfo {
			kafka_UpdateReplicationInfo(cfg, client)
			return
		}
		if _kafkaUpdateSecurity {
			kafka_UpdateSecurity(cfg, client)
			return
		}
		if _kafkaUpdateStorage {
			kafka_UpdateStorage(cfg, client)
			return
		}
		if _kafkaUpdateTopic {
			kafka_UpdateTopic(cfg, client)
			return
		}

	},
}

var (
	_kafkaBatchAssociateScramSecret     bool
	_kafkaBatchDisassociateScramSecret  bool
	_kafkaCreateCluster                 bool
	_kafkaCreateClusterV2               bool
	_kafkaCreateConfiguration           bool
	_kafkaCreateReplicator              bool
	_kafkaCreateTopic                   bool
	_kafkaCreateVpcConnection           bool
	_kafkaDeleteCluster                 bool
	_kafkaDeleteClusterPolicy           bool
	_kafkaDeleteConfiguration           bool
	_kafkaDeleteReplicator              bool
	_kafkaDeleteTopic                   bool
	_kafkaDeleteVpcConnection           bool
	_kafkaDescribeCluster               bool
	_kafkaDescribeClusterOperation      bool
	_kafkaDescribeClusterOperationV2    bool
	_kafkaDescribeClusterV2             bool
	_kafkaDescribeConfiguration         bool
	_kafkaDescribeConfigurationRevision bool
	_kafkaDescribeReplicator            bool
	_kafkaDescribeTopic                 bool
	_kafkaDescribeTopicPartitions       bool
	_kafkaDescribeVpcConnection         bool
	_kafkaGetBootstrapBrokers           bool
	_kafkaGetClusterPolicy              bool
	_kafkaGetCompatibleKafkaVersions    bool
	_kafkaListClientVpcConnections      bool
	_kafkaListClusterOperations         bool
	_kafkaListClusterOperationsV2       bool
	_kafkaListClusters                  bool
	_kafkaListClustersV2                bool
	_kafkaListConfigurationRevisions    bool
	_kafkaListConfigurations            bool
	_kafkaListKafkaVersions             bool
	_kafkaListNodes                     bool
	_kafkaListReplicators               bool
	_kafkaListScramSecrets              bool
	_kafkaListTagsForResource           bool
	_kafkaListTopics                    bool
	_kafkaListVpcConnections            bool
	_kafkaPutClusterPolicy              bool
	_kafkaRebootBroker                  bool
	_kafkaRejectClientVpcConnection     bool
	_kafkaTagResource                   bool
	_kafkaUntagResource                 bool
	_kafkaUpdateBrokerCount             bool
	_kafkaUpdateBrokerStorage           bool
	_kafkaUpdateBrokerType              bool
	_kafkaUpdateClusterConfiguration    bool
	_kafkaUpdateClusterKafkaVersion     bool
	_kafkaUpdateConfiguration           bool
	_kafkaUpdateConnectivity            bool
	_kafkaUpdateMonitoring              bool
	_kafkaUpdateRebalancing             bool
	_kafkaUpdateReplicationInfo         bool
	_kafkaUpdateSecurity                bool
	_kafkaUpdateStorage                 bool
	_kafkaUpdateTopic                   bool

	_kafkaArn                       string
	_kafkaAuthentication            string
	_kafkaBrokerIds                 []string
	_kafkaBrokerNodeGroupInfo       string
	_kafkaClientAuthentication      string
	_kafkaClientSubnets             []string
	_kafkaClusterArn                string
	_kafkaClusterName               string
	_kafkaClusterNameFilter         string
	_kafkaClusterOperationArn       string
	_kafkaClusterTypeFilter         string
	_kafkaConfigs                   string
	_kafkaConfigurationInfo         string
	_kafkaConnectivityInfo          string
	_kafkaConsumerGroupReplication  string
	_kafkaCurrentVersion            string
	_kafkaDescription               string
	_kafkaEncryptionInfo            string
	_kafkaEnhancedMonitoring        string
	_kafkaKafkaClusters             string
	_kafkaKafkaVersion              string
	_kafkaKafkaVersions             []string
	_kafkaLoggingInfo               string
	_kafkaMaxResults                string
	_kafkaName                      string
	_kafkaNextToken                 string
	_kafkaNumberOfBrokerNodes       string
	_kafkaOpenMonitoring            string
	_kafkaPartitionCount            string
	_kafkaPolicy                    string
	_kafkaProvisioned               string
	_kafkaProvisionedThroughput     string
	_kafkaRebalancing               string
	_kafkaReplicationFactor         string
	_kafkaReplicationInfoList       string
	_kafkaReplicatorArn             string
	_kafkaReplicatorName            string
	_kafkaReplicatorNameFilter      string
	_kafkaResourceArn               string
	_kafkaRevision                  string
	_kafkaSecretArnList             []string
	_kafkaSecurityGroups            []string
	_kafkaServerProperties          string
	_kafkaServerless                string
	_kafkaServiceExecutionRoleArn   string
	_kafkaSourceKafkaClusterArn     string
	_kafkaStorageMode               string
	_kafkaTagKeys                   []string
	_kafkaTags                      string
	_kafkaTargetBrokerEBSVolumeInfo string
	_kafkaTargetClusterArn          string
	_kafkaTargetInstanceType        string
	_kafkaTargetKafkaClusterArn     string
	_kafkaTargetKafkaVersion        string
	_kafkaTargetNumberOfBrokerNodes string
	_kafkaTopicName                 string
	_kafkaTopicNameFilter           string
	_kafkaTopicReplication          string
	_kafkaVolumeSizeGB              string
	_kafkaVpcConnectionArn          string
	_kafkaVpcId                     string
)

// Associates one or more Scram Secrets with an Amazon MSK cluster.
func kafka_BatchAssociateScramSecret(cfg aws.Config, client *kafka.Client) {
	input := &kafka.BatchAssociateScramSecretInput{
		// ClusterArn: *string, // Required
		// SecretArnList: []string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaSecretArnList) > 0 {
		input.SecretArnList = append([]string(nil), _kafkaSecretArnList...)
	}

	if resp, err := client.BatchAssociateScramSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates one or more Scram Secrets from an Amazon MSK cluster.
func kafka_BatchDisassociateScramSecret(cfg aws.Config, client *kafka.Client) {
	input := &kafka.BatchDisassociateScramSecretInput{
		// ClusterArn: *string, // Required
		// SecretArnList: []string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaSecretArnList) > 0 {
		input.SecretArnList = append([]string(nil), _kafkaSecretArnList...)
	}

	if resp, err := client.BatchDisassociateScramSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MSK cluster.
func kafka_CreateCluster(cfg aws.Config, client *kafka.Client) {
	input := &kafka.CreateClusterInput{
		// BrokerNodeGroupInfo: *types.BrokerNodeGroupInfo, // Required
		// ClusterName: *string, // Required
		// KafkaVersion: *string, // Required
		// NumberOfBrokerNodes: *int32, // Required
	}

	if len(_kafkaBrokerNodeGroupInfo) > 0 {
		if err := assignInputField(input, "BrokerNodeGroupInfo", _kafkaBrokerNodeGroupInfo); err != nil {
			log.Errorf("invalid --broker-node-group-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaClusterName) > 0 {
		input.ClusterName = aws.String(_kafkaClusterName)
	}
	if len(_kafkaKafkaVersion) > 0 {
		input.KafkaVersion = aws.String(_kafkaKafkaVersion)
	}
	if len(_kafkaNumberOfBrokerNodes) > 0 {
		if err := assignInputField(input, "NumberOfBrokerNodes", _kafkaNumberOfBrokerNodes); err != nil {
			log.Errorf("invalid --number-of-broker-nodes: %s", err.Error())
			return
		}
	}
	if len(_kafkaClientAuthentication) > 0 {
		if err := assignInputField(input, "ClientAuthentication", _kafkaClientAuthentication); err != nil {
			log.Errorf("invalid --client-authentication: %s", err.Error())
			return
		}
	}
	if len(_kafkaConfigurationInfo) > 0 {
		if err := assignInputField(input, "ConfigurationInfo", _kafkaConfigurationInfo); err != nil {
			log.Errorf("invalid --configuration-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaEncryptionInfo) > 0 {
		if err := assignInputField(input, "EncryptionInfo", _kafkaEncryptionInfo); err != nil {
			log.Errorf("invalid --encryption-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaEnhancedMonitoring) > 0 {
		if err := assignInputField(input, "EnhancedMonitoring", _kafkaEnhancedMonitoring); err != nil {
			log.Errorf("invalid --enhanced-monitoring: %s", err.Error())
			return
		}
	}
	if len(_kafkaLoggingInfo) > 0 {
		if err := assignInputField(input, "LoggingInfo", _kafkaLoggingInfo); err != nil {
			log.Errorf("invalid --logging-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaOpenMonitoring) > 0 {
		if err := assignInputField(input, "OpenMonitoring", _kafkaOpenMonitoring); err != nil {
			log.Errorf("invalid --open-monitoring: %s", err.Error())
			return
		}
	}
	if len(_kafkaRebalancing) > 0 {
		if err := assignInputField(input, "Rebalancing", _kafkaRebalancing); err != nil {
			log.Errorf("invalid --rebalancing: %s", err.Error())
			return
		}
	}
	if len(_kafkaStorageMode) > 0 {
		if err := assignInputField(input, "StorageMode", _kafkaStorageMode); err != nil {
			log.Errorf("invalid --storage-mode: %s", err.Error())
			return
		}
	}
	if len(_kafkaTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MSK cluster.
func kafka_CreateClusterV2(cfg aws.Config, client *kafka.Client) {
	input := &kafka.CreateClusterV2Input{
		// ClusterName: *string, // Required
	}

	if len(_kafkaClusterName) > 0 {
		input.ClusterName = aws.String(_kafkaClusterName)
	}
	if len(_kafkaProvisioned) > 0 {
		if err := assignInputField(input, "Provisioned", _kafkaProvisioned); err != nil {
			log.Errorf("invalid --provisioned: %s", err.Error())
			return
		}
	}
	if len(_kafkaServerless) > 0 {
		if err := assignInputField(input, "Serverless", _kafkaServerless); err != nil {
			log.Errorf("invalid --serverless: %s", err.Error())
			return
		}
	}
	if len(_kafkaTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MSK configuration.
func kafka_CreateConfiguration(cfg aws.Config, client *kafka.Client) {
	input := &kafka.CreateConfigurationInput{
		// Name: *string, // Required
		// ServerProperties: []byte, // Required
	}

	if len(_kafkaName) > 0 {
		input.Name = aws.String(_kafkaName)
	}
	if len(_kafkaServerProperties) > 0 {
		if err := assignInputField(input, "ServerProperties", _kafkaServerProperties); err != nil {
			log.Errorf("invalid --server-properties: %s", err.Error())
			return
		}
	}
	if len(_kafkaDescription) > 0 {
		input.Description = aws.String(_kafkaDescription)
	}
	if len(_kafkaKafkaVersions) > 0 {
		input.KafkaVersions = append([]string(nil), _kafkaKafkaVersions...)
	}

	if resp, err := client.CreateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the replicator.
func kafka_CreateReplicator(cfg aws.Config, client *kafka.Client) {
	input := &kafka.CreateReplicatorInput{
		// KafkaClusters: []types.KafkaCluster, // Required
		// ReplicationInfoList: []types.ReplicationInfo, // Required
		// ReplicatorName: *string, // Required
		// ServiceExecutionRoleArn: *string, // Required
	}

	if len(_kafkaKafkaClusters) > 0 {
		if err := assignInputField(input, "KafkaClusters", _kafkaKafkaClusters); err != nil {
			log.Errorf("invalid --kafka-clusters: %s", err.Error())
			return
		}
	}
	if len(_kafkaReplicationInfoList) > 0 {
		if err := assignInputField(input, "ReplicationInfoList", _kafkaReplicationInfoList); err != nil {
			log.Errorf("invalid --replication-info-list: %s", err.Error())
			return
		}
	}
	if len(_kafkaReplicatorName) > 0 {
		input.ReplicatorName = aws.String(_kafkaReplicatorName)
	}
	if len(_kafkaServiceExecutionRoleArn) > 0 {
		input.ServiceExecutionRoleArn = aws.String(_kafkaServiceExecutionRoleArn)
	}
	if len(_kafkaDescription) > 0 {
		input.Description = aws.String(_kafkaDescription)
	}
	if len(_kafkaTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReplicator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a topic in the specified MSK cluster.
func kafka_CreateTopic(cfg aws.Config, client *kafka.Client) {
	input := &kafka.CreateTopicInput{
		// ClusterArn: *string, // Required
		// PartitionCount: *int32, // Required
		// ReplicationFactor: *int32, // Required
		// TopicName: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaPartitionCount) > 0 {
		if err := assignInputField(input, "PartitionCount", _kafkaPartitionCount); err != nil {
			log.Errorf("invalid --partition-count: %s", err.Error())
			return
		}
	}
	if len(_kafkaReplicationFactor) > 0 {
		if err := assignInputField(input, "ReplicationFactor", _kafkaReplicationFactor); err != nil {
			log.Errorf("invalid --replication-factor: %s", err.Error())
			return
		}
	}
	if len(_kafkaTopicName) > 0 {
		input.TopicName = aws.String(_kafkaTopicName)
	}
	if len(_kafkaConfigs) > 0 {
		input.Configs = aws.String(_kafkaConfigs)
	}

	if resp, err := client.CreateTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MSK VPC connection.
func kafka_CreateVpcConnection(cfg aws.Config, client *kafka.Client) {
	input := &kafka.CreateVpcConnectionInput{
		// Authentication: *string, // Required
		// ClientSubnets: []string, // Required
		// SecurityGroups: []string, // Required
		// TargetClusterArn: *string, // Required
		// VpcId: *string, // Required
	}

	if len(_kafkaAuthentication) > 0 {
		input.Authentication = aws.String(_kafkaAuthentication)
	}
	if len(_kafkaClientSubnets) > 0 {
		input.ClientSubnets = append([]string(nil), _kafkaClientSubnets...)
	}
	if len(_kafkaSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _kafkaSecurityGroups...)
	}
	if len(_kafkaTargetClusterArn) > 0 {
		input.TargetClusterArn = aws.String(_kafkaTargetClusterArn)
	}
	if len(_kafkaVpcId) > 0 {
		input.VpcId = aws.String(_kafkaVpcId)
	}
	if len(_kafkaTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the MSK cluster specified by the Amazon Resource Name (ARN) in the
// request.
func kafka_DeleteCluster(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DeleteClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the MSK cluster policy specified by the Amazon Resource Name (ARN) in
// the request.
func kafka_DeleteClusterPolicy(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DeleteClusterPolicyInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.DeleteClusterPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an MSK Configuration.
func kafka_DeleteConfiguration(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DeleteConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}

	if resp, err := client.DeleteConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a replicator.
func kafka_DeleteReplicator(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DeleteReplicatorInput{
		// ReplicatorArn: *string, // Required
	}

	if len(_kafkaReplicatorArn) > 0 {
		input.ReplicatorArn = aws.String(_kafkaReplicatorArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}

	if resp, err := client.DeleteReplicator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a topic in the specified MSK cluster.
func kafka_DeleteTopic(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DeleteTopicInput{
		// ClusterArn: *string, // Required
		// TopicName: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaTopicName) > 0 {
		input.TopicName = aws.String(_kafkaTopicName)
	}

	if resp, err := client.DeleteTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a MSK VPC connection.
func kafka_DeleteVpcConnection(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DeleteVpcConnectionInput{
		// Arn: *string, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}

	if resp, err := client.DeleteVpcConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the MSK cluster whose Amazon Resource Name (ARN) is
// specified in the request.
func kafka_DescribeCluster(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the cluster operation specified by the ARN.
func kafka_DescribeClusterOperation(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeClusterOperationInput{
		// ClusterOperationArn: *string, // Required
	}

	if len(_kafkaClusterOperationArn) > 0 {
		input.ClusterOperationArn = aws.String(_kafkaClusterOperationArn)
	}

	if resp, err := client.DescribeClusterOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the cluster operation specified by the ARN.
func kafka_DescribeClusterOperationV2(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeClusterOperationV2Input{
		// ClusterOperationArn: *string, // Required
	}

	if len(_kafkaClusterOperationArn) > 0 {
		input.ClusterOperationArn = aws.String(_kafkaClusterOperationArn)
	}

	if resp, err := client.DescribeClusterOperationV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the MSK cluster whose Amazon Resource Name (ARN) is
// specified in the request.
func kafka_DescribeClusterV2(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeClusterV2Input{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.DescribeClusterV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of this MSK configuration.
func kafka_DescribeConfiguration(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}

	if resp, err := client.DescribeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of this revision of the configuration.
func kafka_DescribeConfigurationRevision(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeConfigurationRevisionInput{
		// Arn: *string, // Required
		// Revision: *int64, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}
	if len(_kafkaRevision) > 0 {
		if err := assignInputField(input, "Revision", _kafkaRevision); err != nil {
			log.Errorf("invalid --revision: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeConfigurationRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a replicator.
func kafka_DescribeReplicator(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeReplicatorInput{
		// ReplicatorArn: *string, // Required
	}

	if len(_kafkaReplicatorArn) > 0 {
		input.ReplicatorArn = aws.String(_kafkaReplicatorArn)
	}

	if resp, err := client.DescribeReplicator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns topic details of this topic on a MSK cluster.
func kafka_DescribeTopic(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeTopicInput{
		// ClusterArn: *string, // Required
		// TopicName: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaTopicName) > 0 {
		input.TopicName = aws.String(_kafkaTopicName)
	}

	if resp, err := client.DescribeTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns partition details of this topic on a MSK cluster.
func kafka_DescribeTopicPartitions(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeTopicPartitionsInput{
		// ClusterArn: *string, // Required
		// TopicName: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaTopicName) > 0 {
		input.TopicName = aws.String(_kafkaTopicName)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTopicPartitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.DescribeTopicPartitionsOutput
	p := kafka.NewDescribeTopicPartitionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a description of this MSK VPC connection.
func kafka_DescribeVpcConnection(cfg aws.Config, client *kafka.Client) {
	input := &kafka.DescribeVpcConnectionInput{
		// Arn: *string, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}

	if resp, err := client.DescribeVpcConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of brokers that a client application can use to bootstrap. This list
// doesn't necessarily include all of the brokers in the cluster. The following
// Python 3.6 example shows how you can use the Amazon Resource Name (ARN) of a
// cluster to get its bootstrap brokers. If you don't know the ARN of your cluster,
// you can use the ListClusters operation to get the ARNs of all the clusters in
// this account and Region.
func kafka_GetBootstrapBrokers(cfg aws.Config, client *kafka.Client) {
	input := &kafka.GetBootstrapBrokersInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.GetBootstrapBrokers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the MSK cluster policy specified by the Amazon Resource Name (ARN) in the
// request.
func kafka_GetClusterPolicy(cfg aws.Config, client *kafka.Client) {
	input := &kafka.GetClusterPolicyInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.GetClusterPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Apache Kafka versions to which you can update the MSK cluster.
func kafka_GetCompatibleKafkaVersions(cfg aws.Config, client *kafka.Client) {
	input := &kafka.GetCompatibleKafkaVersionsInput{}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.GetCompatibleKafkaVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all the VPC connections in this Region.
func kafka_ListClientVpcConnections(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListClientVpcConnectionsInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClientVpcConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListClientVpcConnectionsOutput
	p := kafka.NewListClientVpcConnectionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the operations that have been performed on the specified
// MSK cluster.
func kafka_ListClusterOperations(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListClusterOperationsInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusterOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListClusterOperationsOutput
	p := kafka.NewListClusterOperationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the operations that have been performed on the specified
// MSK cluster.
func kafka_ListClusterOperationsV2(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListClusterOperationsV2Input{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusterOperationsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListClusterOperationsV2Output
	p := kafka.NewListClusterOperationsV2Paginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the MSK clusters in the current Region.
func kafka_ListClusters(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListClustersInput{}

	if len(_kafkaClusterNameFilter) > 0 {
		input.ClusterNameFilter = aws.String(_kafkaClusterNameFilter)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListClustersOutput
	p := kafka.NewListClustersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the MSK clusters in the current Region.
func kafka_ListClustersV2(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListClustersV2Input{}

	if len(_kafkaClusterNameFilter) > 0 {
		input.ClusterNameFilter = aws.String(_kafkaClusterNameFilter)
	}
	if len(_kafkaClusterTypeFilter) > 0 {
		input.ClusterTypeFilter = aws.String(_kafkaClusterTypeFilter)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClustersV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListClustersV2Output
	p := kafka.NewListClustersV2Paginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the MSK configurations in this Region.
func kafka_ListConfigurationRevisions(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListConfigurationRevisionsInput{
		// Arn: *string, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListConfigurationRevisionsOutput
	p := kafka.NewListConfigurationRevisionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the MSK configurations in this Region.
func kafka_ListConfigurations(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListConfigurationsInput{}

	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListConfigurationsOutput
	p := kafka.NewListConfigurationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of Apache Kafka versions.
func kafka_ListKafkaVersions(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListKafkaVersionsInput{}

	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKafkaVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListKafkaVersionsOutput
	p := kafka.NewListKafkaVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of the broker nodes in the cluster.
func kafka_ListNodes(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListNodesInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListNodesOutput
	p := kafka.NewListNodesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the replicators.
func kafka_ListReplicators(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListReplicatorsInput{}

	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}
	if len(_kafkaReplicatorNameFilter) > 0 {
		input.ReplicatorNameFilter = aws.String(_kafkaReplicatorNameFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListReplicators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListReplicatorsOutput
	p := kafka.NewListReplicatorsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of the Scram Secrets associated with an Amazon MSK cluster.
func kafka_ListScramSecrets(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListScramSecretsInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScramSecrets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListScramSecretsOutput
	p := kafka.NewListScramSecretsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of the tags associated with the specified resource.
func kafka_ListTagsForResource(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_kafkaResourceArn) > 0 {
		input.ResourceArn = aws.String(_kafkaResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List topics in a MSK cluster.
func kafka_ListTopics(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListTopicsInput{
		// ClusterArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}
	if len(_kafkaTopicNameFilter) > 0 {
		input.TopicNameFilter = aws.String(_kafkaTopicNameFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListTopics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListTopicsOutput
	p := kafka.NewListTopicsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all the VPC connections in this Region.
func kafka_ListVpcConnections(cfg aws.Config, client *kafka.Client) {
	input := &kafka.ListVpcConnectionsInput{}

	if len(_kafkaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kafkaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kafkaNextToken) > 0 {
		input.NextToken = aws.String(_kafkaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVpcConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kafka.ListVpcConnectionsOutput
	p := kafka.NewListVpcConnectionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Creates or updates the MSK cluster policy specified by the cluster Amazon
// Resource Name (ARN) in the request.
func kafka_PutClusterPolicy(cfg aws.Config, client *kafka.Client) {
	input := &kafka.PutClusterPolicyInput{
		// ClusterArn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaPolicy) > 0 {
		input.Policy = aws.String(_kafkaPolicy)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}

	if resp, err := client.PutClusterPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots brokers.
func kafka_RebootBroker(cfg aws.Config, client *kafka.Client) {
	input := &kafka.RebootBrokerInput{
		// BrokerIds: []string, // Required
		// ClusterArn: *string, // Required
	}

	if len(_kafkaBrokerIds) > 0 {
		input.BrokerIds = append([]string(nil), _kafkaBrokerIds...)
	}
	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}

	if resp, err := client.RebootBroker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns empty response.
func kafka_RejectClientVpcConnection(cfg aws.Config, client *kafka.Client) {
	input := &kafka.RejectClientVpcConnectionInput{
		// ClusterArn: *string, // Required
		// VpcConnectionArn: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaVpcConnectionArn) > 0 {
		input.VpcConnectionArn = aws.String(_kafkaVpcConnectionArn)
	}

	if resp, err := client.RejectClientVpcConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified MSK resource.
func kafka_TagResource(cfg aws.Config, client *kafka.Client) {
	input := &kafka.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_kafkaResourceArn) > 0 {
		input.ResourceArn = aws.String(_kafkaResourceArn)
	}
	if len(_kafkaTags) > 0 {
		if err := assignInputField(input, "Tags", _kafkaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the tags associated with the keys that are provided in the query.
func kafka_UntagResource(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kafkaResourceArn) > 0 {
		input.ResourceArn = aws.String(_kafkaResourceArn)
	}
	if len(_kafkaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kafkaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the number of broker nodes in the cluster.
func kafka_UpdateBrokerCount(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateBrokerCountInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
		// TargetNumberOfBrokerNodes: *int32, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaTargetNumberOfBrokerNodes) > 0 {
		if err := assignInputField(input, "TargetNumberOfBrokerNodes", _kafkaTargetNumberOfBrokerNodes); err != nil {
			log.Errorf("invalid --target-number-of-broker-nodes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBrokerCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the EBS storage associated with MSK brokers.
func kafka_UpdateBrokerStorage(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateBrokerStorageInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
		// TargetBrokerEBSVolumeInfo: []types.BrokerEBSVolumeInfo, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaTargetBrokerEBSVolumeInfo) > 0 {
		if err := assignInputField(input, "TargetBrokerEBSVolumeInfo", _kafkaTargetBrokerEBSVolumeInfo); err != nil {
			log.Errorf("invalid --target-broker-ebs-volume-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBrokerStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates EC2 instance type.
func kafka_UpdateBrokerType(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateBrokerTypeInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
		// TargetInstanceType: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaTargetInstanceType) > 0 {
		input.TargetInstanceType = aws.String(_kafkaTargetInstanceType)
	}

	if resp, err := client.UpdateBrokerType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the cluster with the configuration that is specified in the request
// body.
func kafka_UpdateClusterConfiguration(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateClusterConfigurationInput{
		// ClusterArn: *string, // Required
		// ConfigurationInfo: *types.ConfigurationInfo, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaConfigurationInfo) > 0 {
		if err := assignInputField(input, "ConfigurationInfo", _kafkaConfigurationInfo); err != nil {
			log.Errorf("invalid --configuration-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}

	if resp, err := client.UpdateClusterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Apache Kafka version for the cluster.
func kafka_UpdateClusterKafkaVersion(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateClusterKafkaVersionInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
		// TargetKafkaVersion: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaTargetKafkaVersion) > 0 {
		input.TargetKafkaVersion = aws.String(_kafkaTargetKafkaVersion)
	}
	if len(_kafkaConfigurationInfo) > 0 {
		if err := assignInputField(input, "ConfigurationInfo", _kafkaConfigurationInfo); err != nil {
			log.Errorf("invalid --configuration-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClusterKafkaVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an MSK configuration.
func kafka_UpdateConfiguration(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateConfigurationInput{
		// Arn: *string, // Required
		// ServerProperties: []byte, // Required
	}

	if len(_kafkaArn) > 0 {
		input.Arn = aws.String(_kafkaArn)
	}
	if len(_kafkaServerProperties) > 0 {
		if err := assignInputField(input, "ServerProperties", _kafkaServerProperties); err != nil {
			log.Errorf("invalid --server-properties: %s", err.Error())
			return
		}
	}
	if len(_kafkaDescription) > 0 {
		input.Description = aws.String(_kafkaDescription)
	}

	if resp, err := client.UpdateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the cluster's connectivity configuration.
func kafka_UpdateConnectivity(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateConnectivityInput{
		// ClusterArn: *string, // Required
		// ConnectivityInfo: *types.ConnectivityInfo, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaConnectivityInfo) > 0 {
		if err := assignInputField(input, "ConnectivityInfo", _kafkaConnectivityInfo); err != nil {
			log.Errorf("invalid --connectivity-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}

	if resp, err := client.UpdateConnectivity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the monitoring settings for the cluster. You can use this operation to
// specify which Apache Kafka metrics you want Amazon MSK to send to Amazon
// CloudWatch. You can also specify settings for open monitoring with Prometheus.
func kafka_UpdateMonitoring(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateMonitoringInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaEnhancedMonitoring) > 0 {
		if err := assignInputField(input, "EnhancedMonitoring", _kafkaEnhancedMonitoring); err != nil {
			log.Errorf("invalid --enhanced-monitoring: %s", err.Error())
			return
		}
	}
	if len(_kafkaLoggingInfo) > 0 {
		if err := assignInputField(input, "LoggingInfo", _kafkaLoggingInfo); err != nil {
			log.Errorf("invalid --logging-info: %s", err.Error())
			return
		}
	}
	if len(_kafkaOpenMonitoring) > 0 {
		if err := assignInputField(input, "OpenMonitoring", _kafkaOpenMonitoring); err != nil {
			log.Errorf("invalid --open-monitoring: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMonitoring(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this resource to update the intelligent rebalancing status of an Amazon MSK
// Provisioned cluster with Express brokers.
func kafka_UpdateRebalancing(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateRebalancingInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
		// Rebalancing: *types.Rebalancing, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaRebalancing) > 0 {
		if err := assignInputField(input, "Rebalancing", _kafkaRebalancing); err != nil {
			log.Errorf("invalid --rebalancing: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRebalancing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates replication info of a replicator.
func kafka_UpdateReplicationInfo(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateReplicationInfoInput{
		// CurrentVersion: *string, // Required
		// ReplicatorArn: *string, // Required
		// SourceKafkaClusterArn: *string, // Required
		// TargetKafkaClusterArn: *string, // Required
	}

	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaReplicatorArn) > 0 {
		input.ReplicatorArn = aws.String(_kafkaReplicatorArn)
	}
	if len(_kafkaSourceKafkaClusterArn) > 0 {
		input.SourceKafkaClusterArn = aws.String(_kafkaSourceKafkaClusterArn)
	}
	if len(_kafkaTargetKafkaClusterArn) > 0 {
		input.TargetKafkaClusterArn = aws.String(_kafkaTargetKafkaClusterArn)
	}
	if len(_kafkaConsumerGroupReplication) > 0 {
		if err := assignInputField(input, "ConsumerGroupReplication", _kafkaConsumerGroupReplication); err != nil {
			log.Errorf("invalid --consumer-group-replication: %s", err.Error())
			return
		}
	}
	if len(_kafkaTopicReplication) > 0 {
		if err := assignInputField(input, "TopicReplication", _kafkaTopicReplication); err != nil {
			log.Errorf("invalid --topic-replication: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReplicationInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the security settings for the cluster. You can use this operation to
// specify encryption and authentication on existing clusters.
func kafka_UpdateSecurity(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateSecurityInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaClientAuthentication) > 0 {
		if err := assignInputField(input, "ClientAuthentication", _kafkaClientAuthentication); err != nil {
			log.Errorf("invalid --client-authentication: %s", err.Error())
			return
		}
	}
	if len(_kafkaEncryptionInfo) > 0 {
		if err := assignInputField(input, "EncryptionInfo", _kafkaEncryptionInfo); err != nil {
			log.Errorf("invalid --encryption-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSecurity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates cluster broker volume size (or) sets cluster storage mode to TIERED.
func kafka_UpdateStorage(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateStorageInput{
		// ClusterArn: *string, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kafkaCurrentVersion)
	}
	if len(_kafkaProvisionedThroughput) > 0 {
		if err := assignInputField(input, "ProvisionedThroughput", _kafkaProvisionedThroughput); err != nil {
			log.Errorf("invalid --provisioned-throughput: %s", err.Error())
			return
		}
	}
	if len(_kafkaStorageMode) > 0 {
		if err := assignInputField(input, "StorageMode", _kafkaStorageMode); err != nil {
			log.Errorf("invalid --storage-mode: %s", err.Error())
			return
		}
	}
	if len(_kafkaVolumeSizeGB) > 0 {
		if err := assignInputField(input, "VolumeSizeGB", _kafkaVolumeSizeGB); err != nil {
			log.Errorf("invalid --volume-size-gb: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of the specified topic.
func kafka_UpdateTopic(cfg aws.Config, client *kafka.Client) {
	input := &kafka.UpdateTopicInput{
		// ClusterArn: *string, // Required
		// TopicName: *string, // Required
	}

	if len(_kafkaClusterArn) > 0 {
		input.ClusterArn = aws.String(_kafkaClusterArn)
	}
	if len(_kafkaTopicName) > 0 {
		input.TopicName = aws.String(_kafkaTopicName)
	}
	if len(_kafkaConfigs) > 0 {
		input.Configs = aws.String(_kafkaConfigs)
	}
	if len(_kafkaPartitionCount) > 0 {
		if err := assignInputField(input, "PartitionCount", _kafkaPartitionCount); err != nil {
			log.Errorf("invalid --partition-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kafkaCmd)
	_kafkaCmd.Flags().SortFlags = false

	_kafkaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_kafkaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kafkaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_kafkaCmd.Flags().StringVarP(&_kafkaArn, "arn", "", "", "ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaAuthentication, "authentication", "", "", "Authentication")
	_kafkaCmd.Flags().StringSliceVarP(&_kafkaBrokerIds, "broker-ids", "", nil, "Broker Ids")
	_kafkaCmd.Flags().StringVarP(&_kafkaBrokerNodeGroupInfo, "broker-node-group-info", "", "", "Broker Node Group Info")
	_kafkaCmd.Flags().StringVarP(&_kafkaClientAuthentication, "client-authentication", "", "", "Client Authentication")
	_kafkaCmd.Flags().StringSliceVarP(&_kafkaClientSubnets, "client-subnets", "", nil, "Client Subnets")
	_kafkaCmd.Flags().StringVarP(&_kafkaClusterArn, "cluster-arn", "", "", "Cluster ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaClusterName, "cluster-name", "", "", "Cluster Name")
	_kafkaCmd.Flags().StringVarP(&_kafkaClusterNameFilter, "cluster-name-filter", "", "", "Cluster Name Filter")
	_kafkaCmd.Flags().StringVarP(&_kafkaClusterOperationArn, "cluster-operation-arn", "", "", "Cluster Operation ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaClusterTypeFilter, "cluster-type-filter", "", "", "Cluster Type Filter")
	_kafkaCmd.Flags().StringVarP(&_kafkaConfigs, "configs", "", "", "Configs")
	_kafkaCmd.Flags().StringVarP(&_kafkaConfigurationInfo, "configuration-info", "", "", "Configuration Info")
	_kafkaCmd.Flags().StringVarP(&_kafkaConnectivityInfo, "connectivity-info", "", "", "Connectivity Info")
	_kafkaCmd.Flags().StringVarP(&_kafkaConsumerGroupReplication, "consumer-group-replication", "", "", "Consumer Group Replication")
	_kafkaCmd.Flags().StringVarP(&_kafkaCurrentVersion, "current-version", "", "", "Current Version")
	_kafkaCmd.Flags().StringVarP(&_kafkaDescription, "description", "", "", "Description")
	_kafkaCmd.Flags().StringVarP(&_kafkaEncryptionInfo, "encryption-info", "", "", "Encryption Info")
	_kafkaCmd.Flags().StringVarP(&_kafkaEnhancedMonitoring, "enhanced-monitoring", "", "", "Enhanced Monitoring")
	_kafkaCmd.Flags().StringVarP(&_kafkaKafkaClusters, "kafka-clusters", "", "", "Kafka Clusters")
	_kafkaCmd.Flags().StringVarP(&_kafkaKafkaVersion, "kafka-version", "", "", "Kafka Version")
	_kafkaCmd.Flags().StringSliceVarP(&_kafkaKafkaVersions, "kafka-versions", "", nil, "Kafka Versions")
	_kafkaCmd.Flags().StringVarP(&_kafkaLoggingInfo, "logging-info", "", "", "Logging Info")
	_kafkaCmd.Flags().StringVarP(&_kafkaMaxResults, "max-results", "", "", "Max Results")
	_kafkaCmd.Flags().StringVarP(&_kafkaName, "name", "", "", "Name")
	_kafkaCmd.Flags().StringVarP(&_kafkaNextToken, "next-token", "", "", "Next Token")
	_kafkaCmd.Flags().StringVarP(&_kafkaNumberOfBrokerNodes, "number-of-broker-nodes", "", "", "Number Of Broker Nodes")
	_kafkaCmd.Flags().StringVarP(&_kafkaOpenMonitoring, "open-monitoring", "", "", "Open Monitoring")
	_kafkaCmd.Flags().StringVarP(&_kafkaPartitionCount, "partition-count", "", "", "Partition Count")
	_kafkaCmd.Flags().StringVarP(&_kafkaPolicy, "policy", "", "", "Policy")
	_kafkaCmd.Flags().StringVarP(&_kafkaProvisioned, "provisioned", "", "", "Provisioned")
	_kafkaCmd.Flags().StringVarP(&_kafkaProvisionedThroughput, "provisioned-throughput", "", "", "Provisioned Throughput")
	_kafkaCmd.Flags().StringVarP(&_kafkaRebalancing, "rebalancing", "", "", "Rebalancing")
	_kafkaCmd.Flags().StringVarP(&_kafkaReplicationFactor, "replication-factor", "", "", "Replication Factor")
	_kafkaCmd.Flags().StringVarP(&_kafkaReplicationInfoList, "replication-info-list", "", "", "Replication Info List")
	_kafkaCmd.Flags().StringVarP(&_kafkaReplicatorArn, "replicator-arn", "", "", "Replicator ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaReplicatorName, "replicator-name", "", "", "Replicator Name")
	_kafkaCmd.Flags().StringVarP(&_kafkaReplicatorNameFilter, "replicator-name-filter", "", "", "Replicator Name Filter")
	_kafkaCmd.Flags().StringVarP(&_kafkaResourceArn, "resource-arn", "", "", "Resource ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaRevision, "revision", "", "", "Revision")
	_kafkaCmd.Flags().StringSliceVarP(&_kafkaSecretArnList, "secret-arn-list", "", nil, "Secret ARN List")
	_kafkaCmd.Flags().StringSliceVarP(&_kafkaSecurityGroups, "security-groups", "", nil, "Security Groups")
	_kafkaCmd.Flags().StringVarP(&_kafkaServerProperties, "server-properties", "", "", "Server Properties")
	_kafkaCmd.Flags().StringVarP(&_kafkaServerless, "serverless", "", "", "Serverless")
	_kafkaCmd.Flags().StringVarP(&_kafkaServiceExecutionRoleArn, "service-execution-role-arn", "", "", "Service Execution Role ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaSourceKafkaClusterArn, "source-kafka-cluster-arn", "", "", "Source Kafka Cluster ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaStorageMode, "storage-mode", "", "", "Storage Mode")
	_kafkaCmd.Flags().StringSliceVarP(&_kafkaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_kafkaCmd.Flags().StringVarP(&_kafkaTags, "tags", "", "", "Tags")
	_kafkaCmd.Flags().StringVarP(&_kafkaTargetBrokerEBSVolumeInfo, "target-broker-ebs-volume-info", "", "", "Target Broker Ebs Volume Info")
	_kafkaCmd.Flags().StringVarP(&_kafkaTargetClusterArn, "target-cluster-arn", "", "", "Target Cluster ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaTargetInstanceType, "target-instance-type", "", "", "Target Instance Type")
	_kafkaCmd.Flags().StringVarP(&_kafkaTargetKafkaClusterArn, "target-kafka-cluster-arn", "", "", "Target Kafka Cluster ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaTargetKafkaVersion, "target-kafka-version", "", "", "Target Kafka Version")
	_kafkaCmd.Flags().StringVarP(&_kafkaTargetNumberOfBrokerNodes, "target-number-of-broker-nodes", "", "", "Target Number Of Broker Nodes")
	_kafkaCmd.Flags().StringVarP(&_kafkaTopicName, "topic-name", "", "", "Topic Name")
	_kafkaCmd.Flags().StringVarP(&_kafkaTopicNameFilter, "topic-name-filter", "", "", "Topic Name Filter")
	_kafkaCmd.Flags().StringVarP(&_kafkaTopicReplication, "topic-replication", "", "", "Topic Replication")
	_kafkaCmd.Flags().StringVarP(&_kafkaVolumeSizeGB, "volume-size-gb", "", "", "Volume Size Gb")
	_kafkaCmd.Flags().StringVarP(&_kafkaVpcConnectionArn, "vpc-connection-arn", "", "", "VPC Connection ARN")
	_kafkaCmd.Flags().StringVarP(&_kafkaVpcId, "vpc-id", "", "", "VPC ID")

	_kafkaCmd.Flags().BoolVarP(&_kafkaBatchAssociateScramSecret, "batch-associate-scram-secret", "", false, "Batch Associate Scram Secret")
	_kafkaCmd.Flags().BoolVarP(&_kafkaBatchDisassociateScramSecret, "batch-disassociate-scram-secret", "", false, "Batch Disassociate Scram Secret")
	_kafkaCmd.Flags().BoolVarP(&_kafkaCreateCluster, "create-cluster", "", false, "Create Cluster")
	_kafkaCmd.Flags().BoolVarP(&_kafkaCreateClusterV2, "create-cluster-v2", "", false, "Create Cluster V2")
	_kafkaCmd.Flags().BoolVarP(&_kafkaCreateConfiguration, "create-configuration", "", false, "Create Configuration")
	_kafkaCmd.Flags().BoolVarP(&_kafkaCreateReplicator, "create-replicator", "", false, "Create Replicator")
	_kafkaCmd.Flags().BoolVarP(&_kafkaCreateTopic, "create-topic", "", false, "Create Topic")
	_kafkaCmd.Flags().BoolVarP(&_kafkaCreateVpcConnection, "create-vpc-connection", "", false, "Create VPC Connection")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDeleteClusterPolicy, "delete-cluster-policy", "", false, "Delete Cluster Policy")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDeleteConfiguration, "delete-configuration", "", false, "Delete Configuration")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDeleteReplicator, "delete-replicator", "", false, "Delete Replicator")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDeleteTopic, "delete-topic", "", false, "Delete Topic")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDeleteVpcConnection, "delete-vpc-connection", "", false, "Delete VPC Connection")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeClusterOperation, "describe-cluster-operation", "", false, "Describe Cluster Operation")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeClusterOperationV2, "describe-cluster-operation-v2", "", false, "Describe Cluster Operation V2")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeClusterV2, "describe-cluster-v2", "", false, "Describe Cluster V2")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeConfiguration, "describe-configuration", "", false, "Describe Configuration")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeConfigurationRevision, "describe-configuration-revision", "", false, "Describe Configuration Revision")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeReplicator, "describe-replicator", "", false, "Describe Replicator")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeTopic, "describe-topic", "", false, "Describe Topic")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeTopicPartitions, "describe-topic-partitions", "", false, "Describe Topic Partitions")
	_kafkaCmd.Flags().BoolVarP(&_kafkaDescribeVpcConnection, "describe-vpc-connection", "", false, "Describe VPC Connection")
	_kafkaCmd.Flags().BoolVarP(&_kafkaGetBootstrapBrokers, "get-bootstrap-brokers", "", false, "Get Bootstrap Brokers")
	_kafkaCmd.Flags().BoolVarP(&_kafkaGetClusterPolicy, "get-cluster-policy", "", false, "Get Cluster Policy")
	_kafkaCmd.Flags().BoolVarP(&_kafkaGetCompatibleKafkaVersions, "get-compatible-kafka-versions", "", false, "Get Compatible Kafka Versions")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListClientVpcConnections, "list-client-vpc-connections", "", false, "List Client VPC Connections")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListClusterOperations, "list-cluster-operations", "", false, "List Cluster Operations")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListClusterOperationsV2, "list-cluster-operations-v2", "", false, "List Cluster Operations V2")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListClusters, "list-clusters", "", false, "List Clusters")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListClustersV2, "list-clusters-v2", "", false, "List Clusters V2")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListConfigurationRevisions, "list-configuration-revisions", "", false, "List Configuration Revisions")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListConfigurations, "list-configurations", "", false, "List Configurations")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListKafkaVersions, "list-kafka-versions", "", false, "List Kafka Versions")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListNodes, "list-nodes", "", false, "List Nodes")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListReplicators, "list-replicators", "", false, "List Replicators")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListScramSecrets, "list-scram-secrets", "", false, "List Scram Secrets")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListTopics, "list-topics", "", false, "List Topics")
	_kafkaCmd.Flags().BoolVarP(&_kafkaListVpcConnections, "list-vpc-connections", "", false, "List VPC Connections")
	_kafkaCmd.Flags().BoolVarP(&_kafkaPutClusterPolicy, "put-cluster-policy", "", false, "Put Cluster Policy")
	_kafkaCmd.Flags().BoolVarP(&_kafkaRebootBroker, "reboot-broker", "", false, "Reboot Broker")
	_kafkaCmd.Flags().BoolVarP(&_kafkaRejectClientVpcConnection, "reject-client-vpc-connection", "", false, "Reject Client VPC Connection")
	_kafkaCmd.Flags().BoolVarP(&_kafkaTagResource, "tag-resource", "", false, "Tag Resource")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUntagResource, "untag-resource", "", false, "Untag Resource")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateBrokerCount, "update-broker-count", "", false, "Update Broker Count")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateBrokerStorage, "update-broker-storage", "", false, "Update Broker Storage")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateBrokerType, "update-broker-type", "", false, "Update Broker Type")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateClusterConfiguration, "update-cluster-configuration", "", false, "Update Cluster Configuration")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateClusterKafkaVersion, "update-cluster-kafka-version", "", false, "Update Cluster Kafka Version")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateConfiguration, "update-configuration", "", false, "Update Configuration")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateConnectivity, "update-connectivity", "", false, "Update Connectivity")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateMonitoring, "update-monitoring", "", false, "Update Monitoring")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateRebalancing, "update-rebalancing", "", false, "Update Rebalancing")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateReplicationInfo, "update-replication-info", "", false, "Update Replication Info")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateSecurity, "update-security", "", false, "Update Security")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateStorage, "update-storage", "", false, "Update Storage")
	_kafkaCmd.Flags().BoolVarP(&_kafkaUpdateTopic, "update-topic", "", false, "Update Topic")

}
