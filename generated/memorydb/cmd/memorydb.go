package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// memorydbCmd represents the memorydb command
var _memorydbCmd = &cobra.Command{
	Use:   "memorydb",
	Short: "AWS memorydb CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := memorydb.NewFromConfig(cfg)
		if _memorydbBatchUpdateCluster {
			memorydb_BatchUpdateCluster(cfg, client)
			return
		}
		if _memorydbCopySnapshot {
			memorydb_CopySnapshot(cfg, client)
			return
		}
		if _memorydbCreateACL {
			memorydb_CreateACL(cfg, client)
			return
		}
		if _memorydbCreateCluster {
			memorydb_CreateCluster(cfg, client)
			return
		}
		if _memorydbCreateMultiRegionCluster {
			memorydb_CreateMultiRegionCluster(cfg, client)
			return
		}
		if _memorydbCreateParameterGroup {
			memorydb_CreateParameterGroup(cfg, client)
			return
		}
		if _memorydbCreateSnapshot {
			memorydb_CreateSnapshot(cfg, client)
			return
		}
		if _memorydbCreateSubnetGroup {
			memorydb_CreateSubnetGroup(cfg, client)
			return
		}
		if _memorydbCreateUser {
			memorydb_CreateUser(cfg, client)
			return
		}
		if _memorydbDeleteACL {
			memorydb_DeleteACL(cfg, client)
			return
		}
		if _memorydbDeleteCluster {
			memorydb_DeleteCluster(cfg, client)
			return
		}
		if _memorydbDeleteMultiRegionCluster {
			memorydb_DeleteMultiRegionCluster(cfg, client)
			return
		}
		if _memorydbDeleteParameterGroup {
			memorydb_DeleteParameterGroup(cfg, client)
			return
		}
		if _memorydbDeleteSnapshot {
			memorydb_DeleteSnapshot(cfg, client)
			return
		}
		if _memorydbDeleteSubnetGroup {
			memorydb_DeleteSubnetGroup(cfg, client)
			return
		}
		if _memorydbDeleteUser {
			memorydb_DeleteUser(cfg, client)
			return
		}
		if _memorydbDescribeACLs {
			memorydb_DescribeACLs(cfg, client)
			return
		}
		if _memorydbDescribeClusters {
			memorydb_DescribeClusters(cfg, client)
			return
		}
		if _memorydbDescribeEngineVersions {
			memorydb_DescribeEngineVersions(cfg, client)
			return
		}
		if _memorydbDescribeEvents {
			memorydb_DescribeEvents(cfg, client)
			return
		}
		if _memorydbDescribeMultiRegionClusters {
			memorydb_DescribeMultiRegionClusters(cfg, client)
			return
		}
		if _memorydbDescribeMultiRegionParameterGroups {
			memorydb_DescribeMultiRegionParameterGroups(cfg, client)
			return
		}
		if _memorydbDescribeMultiRegionParameters {
			memorydb_DescribeMultiRegionParameters(cfg, client)
			return
		}
		if _memorydbDescribeParameterGroups {
			memorydb_DescribeParameterGroups(cfg, client)
			return
		}
		if _memorydbDescribeParameters {
			memorydb_DescribeParameters(cfg, client)
			return
		}
		if _memorydbDescribeReservedNodes {
			memorydb_DescribeReservedNodes(cfg, client)
			return
		}
		if _memorydbDescribeReservedNodesOfferings {
			memorydb_DescribeReservedNodesOfferings(cfg, client)
			return
		}
		if _memorydbDescribeServiceUpdates {
			memorydb_DescribeServiceUpdates(cfg, client)
			return
		}
		if _memorydbDescribeSnapshots {
			memorydb_DescribeSnapshots(cfg, client)
			return
		}
		if _memorydbDescribeSubnetGroups {
			memorydb_DescribeSubnetGroups(cfg, client)
			return
		}
		if _memorydbDescribeUsers {
			memorydb_DescribeUsers(cfg, client)
			return
		}
		if _memorydbFailoverShard {
			memorydb_FailoverShard(cfg, client)
			return
		}
		if _memorydbListAllowedMultiRegionClusterUpdates {
			memorydb_ListAllowedMultiRegionClusterUpdates(cfg, client)
			return
		}
		if _memorydbListAllowedNodeTypeUpdates {
			memorydb_ListAllowedNodeTypeUpdates(cfg, client)
			return
		}
		if _memorydbListTags {
			memorydb_ListTags(cfg, client)
			return
		}
		if _memorydbPurchaseReservedNodesOffering {
			memorydb_PurchaseReservedNodesOffering(cfg, client)
			return
		}
		if _memorydbResetParameterGroup {
			memorydb_ResetParameterGroup(cfg, client)
			return
		}
		if _memorydbTagResource {
			memorydb_TagResource(cfg, client)
			return
		}
		if _memorydbUntagResource {
			memorydb_UntagResource(cfg, client)
			return
		}
		if _memorydbUpdateACL {
			memorydb_UpdateACL(cfg, client)
			return
		}
		if _memorydbUpdateCluster {
			memorydb_UpdateCluster(cfg, client)
			return
		}
		if _memorydbUpdateMultiRegionCluster {
			memorydb_UpdateMultiRegionCluster(cfg, client)
			return
		}
		if _memorydbUpdateParameterGroup {
			memorydb_UpdateParameterGroup(cfg, client)
			return
		}
		if _memorydbUpdateSubnetGroup {
			memorydb_UpdateSubnetGroup(cfg, client)
			return
		}
		if _memorydbUpdateUser {
			memorydb_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_memorydbBatchUpdateCluster                   bool
	_memorydbCopySnapshot                         bool
	_memorydbCreateACL                            bool
	_memorydbCreateCluster                        bool
	_memorydbCreateMultiRegionCluster             bool
	_memorydbCreateParameterGroup                 bool
	_memorydbCreateSnapshot                       bool
	_memorydbCreateSubnetGroup                    bool
	_memorydbCreateUser                           bool
	_memorydbDeleteACL                            bool
	_memorydbDeleteCluster                        bool
	_memorydbDeleteMultiRegionCluster             bool
	_memorydbDeleteParameterGroup                 bool
	_memorydbDeleteSnapshot                       bool
	_memorydbDeleteSubnetGroup                    bool
	_memorydbDeleteUser                           bool
	_memorydbDescribeACLs                         bool
	_memorydbDescribeClusters                     bool
	_memorydbDescribeEngineVersions               bool
	_memorydbDescribeEvents                       bool
	_memorydbDescribeMultiRegionClusters          bool
	_memorydbDescribeMultiRegionParameterGroups   bool
	_memorydbDescribeMultiRegionParameters        bool
	_memorydbDescribeParameterGroups              bool
	_memorydbDescribeParameters                   bool
	_memorydbDescribeReservedNodes                bool
	_memorydbDescribeReservedNodesOfferings       bool
	_memorydbDescribeServiceUpdates               bool
	_memorydbDescribeSnapshots                    bool
	_memorydbDescribeSubnetGroups                 bool
	_memorydbDescribeUsers                        bool
	_memorydbFailoverShard                        bool
	_memorydbListAllowedMultiRegionClusterUpdates bool
	_memorydbListAllowedNodeTypeUpdates           bool
	_memorydbListTags                             bool
	_memorydbPurchaseReservedNodesOffering        bool
	_memorydbResetParameterGroup                  bool
	_memorydbTagResource                          bool
	_memorydbUntagResource                        bool
	_memorydbUpdateACL                            bool
	_memorydbUpdateCluster                        bool
	_memorydbUpdateMultiRegionCluster             bool
	_memorydbUpdateParameterGroup                 bool
	_memorydbUpdateSubnetGroup                    bool
	_memorydbUpdateUser                           bool

	_memorydbAccessString                  string
	_memorydbACLName                       string
	_memorydbAllParameters                 string
	_memorydbAuthenticationMode            string
	_memorydbAutoMinorVersionUpgrade       string
	_memorydbClusterName                   string
	_memorydbClusterNames                  []string
	_memorydbDataTiering                   string
	_memorydbDefaultOnly                   string
	_memorydbDescription                   string
	_memorydbDuration                      string
	_memorydbEndTime                       string
	_memorydbEngine                        string
	_memorydbEngineVersion                 string
	_memorydbFamily                        string
	_memorydbFilters                       string
	_memorydbFinalSnapshotName             string
	_memorydbIpDiscovery                   string
	_memorydbKmsKeyId                      string
	_memorydbMaintenanceWindow             string
	_memorydbMaxResults                    string
	_memorydbMultiRegionClusterName        string
	_memorydbMultiRegionClusterNameSuffix  string
	_memorydbMultiRegionParameterGroupName string
	_memorydbNetworkType                   string
	_memorydbNextToken                     string
	_memorydbNodeCount                     string
	_memorydbNodeType                      string
	_memorydbNumReplicasPerShard           string
	_memorydbNumShards                     string
	_memorydbOfferingType                  string
	_memorydbParameterGroupFamily          string
	_memorydbParameterGroupName            string
	_memorydbParameterNameValues           string
	_memorydbParameterNames                []string
	_memorydbPort                          string
	_memorydbReplicaConfiguration          string
	_memorydbReservationId                 string
	_memorydbReservedNodesOfferingId       string
	_memorydbResourceArn                   string
	_memorydbSecurityGroupIds              []string
	_memorydbServiceUpdate                 string
	_memorydbServiceUpdateName             string
	_memorydbShardConfiguration            string
	_memorydbShardName                     string
	_memorydbShowClusterDetails            string
	_memorydbShowDetail                    string
	_memorydbShowShardDetails              string
	_memorydbSnapshotArns                  []string
	_memorydbSnapshotName                  string
	_memorydbSnapshotRetentionLimit        string
	_memorydbSnapshotWindow                string
	_memorydbSnsTopicArn                   string
	_memorydbSnsTopicStatus                string
	_memorydbSource                        string
	_memorydbSourceName                    string
	_memorydbSourceSnapshotName            string
	_memorydbSourceType                    string
	_memorydbStartTime                     string
	_memorydbStatus                        string
	_memorydbSubnetGroupName               string
	_memorydbSubnetIds                     []string
	_memorydbTagKeys                       []string
	_memorydbTags                          string
	_memorydbTargetBucket                  string
	_memorydbTargetSnapshotName            string
	_memorydbTLSEnabled                    string
	_memorydbUpdateStrategy                string
	_memorydbUserName                      string
	_memorydbUserNames                     []string
	_memorydbUserNamesToAdd                []string
	_memorydbUserNamesToRemove             []string
)

// Apply the service update to a list of clusters supplied. For more information
// on service updates and applying them, see [Applying the service updates].
//
// [Applying the service updates]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/managing-updates.html#applying-updates
func memorydb_BatchUpdateCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.BatchUpdateClusterInput{
		// ClusterNames: []string, // Required
	}

	if len(_memorydbClusterNames) > 0 {
		input.ClusterNames = append([]string(nil), _memorydbClusterNames...)
	}
	if len(_memorydbServiceUpdate) > 0 {
		if err := assignInputField(input, "ServiceUpdate", _memorydbServiceUpdate); err != nil {
			log.Errorf("invalid --service-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Makes a copy of an existing snapshot.
func memorydb_CopySnapshot(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CopySnapshotInput{
		// SourceSnapshotName: *string, // Required
		// TargetSnapshotName: *string, // Required
	}

	if len(_memorydbSourceSnapshotName) > 0 {
		input.SourceSnapshotName = aws.String(_memorydbSourceSnapshotName)
	}
	if len(_memorydbTargetSnapshotName) > 0 {
		input.TargetSnapshotName = aws.String(_memorydbTargetSnapshotName)
	}
	if len(_memorydbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_memorydbKmsKeyId)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_memorydbTargetBucket) > 0 {
		input.TargetBucket = aws.String(_memorydbTargetBucket)
	}

	if resp, err := client.CopySnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Access Control List. For more information, see [Authenticating users with Access Contol Lists (ACLs)].
//
// [Authenticating users with Access Contol Lists (ACLs)]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/clusters.acls.html
func memorydb_CreateACL(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateACLInput{
		// ACLName: *string, // Required
	}

	if len(_memorydbACLName) > 0 {
		input.ACLName = aws.String(_memorydbACLName)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_memorydbUserNames) > 0 {
		input.UserNames = append([]string(nil), _memorydbUserNames...)
	}

	if resp, err := client.CreateACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cluster. All nodes in the cluster run the same protocol-compliant
// engine software.
func memorydb_CreateCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateClusterInput{
		// ACLName: *string, // Required
		// ClusterName: *string, // Required
		// NodeType: *string, // Required
	}

	if len(_memorydbACLName) > 0 {
		input.ACLName = aws.String(_memorydbACLName)
	}
	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbNodeType) > 0 {
		input.NodeType = aws.String(_memorydbNodeType)
	}
	if len(_memorydbAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _memorydbAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_memorydbDataTiering) > 0 {
		if err := assignInputField(input, "DataTiering", _memorydbDataTiering); err != nil {
			log.Errorf("invalid --data-tiering: %s", err.Error())
			return
		}
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbEngine) > 0 {
		input.Engine = aws.String(_memorydbEngine)
	}
	if len(_memorydbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_memorydbEngineVersion)
	}
	if len(_memorydbIpDiscovery) > 0 {
		if err := assignInputField(input, "IpDiscovery", _memorydbIpDiscovery); err != nil {
			log.Errorf("invalid --ip-discovery: %s", err.Error())
			return
		}
	}
	if len(_memorydbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_memorydbKmsKeyId)
	}
	if len(_memorydbMaintenanceWindow) > 0 {
		input.MaintenanceWindow = aws.String(_memorydbMaintenanceWindow)
	}
	if len(_memorydbMultiRegionClusterName) > 0 {
		input.MultiRegionClusterName = aws.String(_memorydbMultiRegionClusterName)
	}
	if len(_memorydbNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _memorydbNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_memorydbNumReplicasPerShard) > 0 {
		if err := assignInputField(input, "NumReplicasPerShard", _memorydbNumReplicasPerShard); err != nil {
			log.Errorf("invalid --num-replicas-per-shard: %s", err.Error())
			return
		}
	}
	if len(_memorydbNumShards) > 0 {
		if err := assignInputField(input, "NumShards", _memorydbNumShards); err != nil {
			log.Errorf("invalid --num-shards: %s", err.Error())
			return
		}
	}
	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}
	if len(_memorydbPort) > 0 {
		if err := assignInputField(input, "Port", _memorydbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_memorydbSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _memorydbSecurityGroupIds...)
	}
	if len(_memorydbSnapshotArns) > 0 {
		input.SnapshotArns = append([]string(nil), _memorydbSnapshotArns...)
	}
	if len(_memorydbSnapshotName) > 0 {
		input.SnapshotName = aws.String(_memorydbSnapshotName)
	}
	if len(_memorydbSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _memorydbSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_memorydbSnapshotWindow) > 0 {
		input.SnapshotWindow = aws.String(_memorydbSnapshotWindow)
	}
	if len(_memorydbSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_memorydbSnsTopicArn)
	}
	if len(_memorydbSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_memorydbSubnetGroupName)
	}
	if len(_memorydbTLSEnabled) > 0 {
		if err := assignInputField(input, "TLSEnabled", _memorydbTLSEnabled); err != nil {
			log.Errorf("invalid --tls-enabled: %s", err.Error())
			return
		}
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
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

// Creates a new multi-Region cluster.
func memorydb_CreateMultiRegionCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateMultiRegionClusterInput{
		// MultiRegionClusterNameSuffix: *string, // Required
		// NodeType: *string, // Required
	}

	if len(_memorydbMultiRegionClusterNameSuffix) > 0 {
		input.MultiRegionClusterNameSuffix = aws.String(_memorydbMultiRegionClusterNameSuffix)
	}
	if len(_memorydbNodeType) > 0 {
		input.NodeType = aws.String(_memorydbNodeType)
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbEngine) > 0 {
		input.Engine = aws.String(_memorydbEngine)
	}
	if len(_memorydbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_memorydbEngineVersion)
	}
	if len(_memorydbMultiRegionParameterGroupName) > 0 {
		input.MultiRegionParameterGroupName = aws.String(_memorydbMultiRegionParameterGroupName)
	}
	if len(_memorydbNumShards) > 0 {
		if err := assignInputField(input, "NumShards", _memorydbNumShards); err != nil {
			log.Errorf("invalid --num-shards: %s", err.Error())
			return
		}
	}
	if len(_memorydbTLSEnabled) > 0 {
		if err := assignInputField(input, "TLSEnabled", _memorydbTLSEnabled); err != nil {
			log.Errorf("invalid --tls-enabled: %s", err.Error())
			return
		}
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMultiRegionCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MemoryDB parameter group. A parameter group is a collection of
// parameters and their values that are applied to all of the nodes in any cluster.
// For more information, see [Configuring engine parameters using parameter groups].
//
// [Configuring engine parameters using parameter groups]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/parametergroups.html
func memorydb_CreateParameterGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateParameterGroupInput{
		// Family: *string, // Required
		// ParameterGroupName: *string, // Required
	}

	if len(_memorydbFamily) > 0 {
		input.Family = aws.String(_memorydbFamily)
	}
	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a copy of an entire cluster at a specific moment in time.
func memorydb_CreateSnapshot(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateSnapshotInput{
		// ClusterName: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbSnapshotName) > 0 {
		input.SnapshotName = aws.String(_memorydbSnapshotName)
	}
	if len(_memorydbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_memorydbKmsKeyId)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subnet group. A subnet group is a collection of subnets (typically
// private) that you can designate for your clusters running in an Amazon Virtual
// Private Cloud (VPC) environment.
//
// When you create a cluster in an Amazon VPC, you must specify a subnet group.
// MemoryDB uses that subnet group to choose a subnet and IP addresses within that
// subnet to associate with your nodes. For more information, see [Subnets and subnet groups].
//
// [Subnets and subnet groups]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/subnetgroups.html
func memorydb_CreateSubnetGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateSubnetGroupInput{
		// SubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_memorydbSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_memorydbSubnetGroupName)
	}
	if len(_memorydbSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _memorydbSubnetIds...)
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a MemoryDB user. For more information, see [Authenticating users with Access Contol Lists (ACLs)].
//
// [Authenticating users with Access Contol Lists (ACLs)]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/clusters.acls.html
func memorydb_CreateUser(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.CreateUserInput{
		// AccessString: *string, // Required
		// AuthenticationMode: *types.AuthenticationMode, // Required
		// UserName: *string, // Required
	}

	if len(_memorydbAccessString) > 0 {
		input.AccessString = aws.String(_memorydbAccessString)
	}
	if len(_memorydbAuthenticationMode) > 0 {
		if err := assignInputField(input, "AuthenticationMode", _memorydbAuthenticationMode); err != nil {
			log.Errorf("invalid --authentication-mode: %s", err.Error())
			return
		}
	}
	if len(_memorydbUserName) > 0 {
		input.UserName = aws.String(_memorydbUserName)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Access Control List. The ACL must first be disassociated from the
// cluster before it can be deleted. For more information, see [Authenticating users with Access Contol Lists (ACLs)].
//
// [Authenticating users with Access Contol Lists (ACLs)]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/clusters.acls.html
func memorydb_DeleteACL(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteACLInput{
		// ACLName: *string, // Required
	}

	if len(_memorydbACLName) > 0 {
		input.ACLName = aws.String(_memorydbACLName)
	}

	if resp, err := client.DeleteACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cluster. It also deletes all associated nodes and node endpoints.
// CreateSnapshot permission is required to create a final snapshot. Without this
// permission, the API call will fail with an Access Denied exception.
func memorydb_DeleteCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbFinalSnapshotName) > 0 {
		input.FinalSnapshotName = aws.String(_memorydbFinalSnapshotName)
	}
	if len(_memorydbMultiRegionClusterName) > 0 {
		input.MultiRegionClusterName = aws.String(_memorydbMultiRegionClusterName)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing multi-Region cluster.
func memorydb_DeleteMultiRegionCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteMultiRegionClusterInput{
		// MultiRegionClusterName: *string, // Required
	}

	if len(_memorydbMultiRegionClusterName) > 0 {
		input.MultiRegionClusterName = aws.String(_memorydbMultiRegionClusterName)
	}

	if resp, err := client.DeleteMultiRegionCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified parameter group. You cannot delete a parameter group if
// it is associated with any clusters. You cannot delete the default parameter
// groups in your account.
func memorydb_DeleteParameterGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteParameterGroupInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}

	if resp, err := client.DeleteParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing snapshot. When you receive a successful response from this
// operation, MemoryDB immediately begins deleting the snapshot; you cannot cancel
// or revert this operation.
func memorydb_DeleteSnapshot(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteSnapshotInput{
		// SnapshotName: *string, // Required
	}

	if len(_memorydbSnapshotName) > 0 {
		input.SnapshotName = aws.String(_memorydbSnapshotName)
	}

	if resp, err := client.DeleteSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subnet group. You cannot delete a default subnet group or one that is
// associated with any clusters.
func memorydb_DeleteSubnetGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteSubnetGroupInput{
		// SubnetGroupName: *string, // Required
	}

	if len(_memorydbSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_memorydbSubnetGroupName)
	}

	if resp, err := client.DeleteSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user. The user will be removed from all ACLs and in turn removed from
// all clusters.
func memorydb_DeleteUser(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DeleteUserInput{
		// UserName: *string, // Required
	}

	if len(_memorydbUserName) > 0 {
		input.UserName = aws.String(_memorydbUserName)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of ACLs.
func memorydb_DescribeACLs(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeACLsInput{}

	if len(_memorydbACLName) > 0 {
		input.ACLName = aws.String(_memorydbACLName)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeACLs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeACLsOutput
	p := memorydb.NewDescribeACLsPaginator(client, input)
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

// Returns information about all provisioned clusters if no cluster identifier is
// specified, or about a specific cluster if a cluster name is supplied.
func memorydb_DescribeClusters(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeClustersInput{}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbShowShardDetails) > 0 {
		if err := assignInputField(input, "ShowShardDetails", _memorydbShowShardDetails); err != nil {
			log.Errorf("invalid --show-shard-details: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeClustersOutput
	p := memorydb.NewDescribeClustersPaginator(client, input)
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

// Returns a list of the available Redis OSS engine versions.
func memorydb_DescribeEngineVersions(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeEngineVersionsInput{}

	if len(_memorydbDefaultOnly) > 0 {
		if err := assignInputField(input, "DefaultOnly", _memorydbDefaultOnly); err != nil {
			log.Errorf("invalid --default-only: %s", err.Error())
			return
		}
	}
	if len(_memorydbEngine) > 0 {
		input.Engine = aws.String(_memorydbEngine)
	}
	if len(_memorydbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_memorydbEngineVersion)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbParameterGroupFamily) > 0 {
		input.ParameterGroupFamily = aws.String(_memorydbParameterGroupFamily)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeEngineVersionsOutput
	p := memorydb.NewDescribeEngineVersionsPaginator(client, input)
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

// Returns events related to clusters, security groups, and parameter groups. You
// can obtain events specific to a particular cluster, security group, or parameter
// group by providing the name as a parameter.
//
// By default, only the events occurring within the last hour are returned;
// however, you can retrieve up to 14 days' worth of events if necessary.
func memorydb_DescribeEvents(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeEventsInput{}

	if len(_memorydbDuration) > 0 {
		if err := assignInputField(input, "Duration", _memorydbDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_memorydbEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _memorydbEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbSourceName) > 0 {
		input.SourceName = aws.String(_memorydbSourceName)
	}
	if len(_memorydbSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _memorydbSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_memorydbStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _memorydbStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeEventsOutput
	p := memorydb.NewDescribeEventsPaginator(client, input)
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

// Returns details about one or more multi-Region clusters.
func memorydb_DescribeMultiRegionClusters(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeMultiRegionClustersInput{}

	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbMultiRegionClusterName) > 0 {
		input.MultiRegionClusterName = aws.String(_memorydbMultiRegionClusterName)
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbShowClusterDetails) > 0 {
		if err := assignInputField(input, "ShowClusterDetails", _memorydbShowClusterDetails); err != nil {
			log.Errorf("invalid --show-cluster-details: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMultiRegionClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeMultiRegionClustersOutput
	p := memorydb.NewDescribeMultiRegionClustersPaginator(client, input)
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

// Returns a list of multi-region parameter groups.
func memorydb_DescribeMultiRegionParameterGroups(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeMultiRegionParameterGroupsInput{}

	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbMultiRegionParameterGroupName) > 0 {
		input.MultiRegionParameterGroupName = aws.String(_memorydbMultiRegionParameterGroupName)
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}

	if resp, err := client.DescribeMultiRegionParameterGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the detailed parameter list for a particular multi-region parameter
// group.
func memorydb_DescribeMultiRegionParameters(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeMultiRegionParametersInput{
		// MultiRegionParameterGroupName: *string, // Required
	}

	if len(_memorydbMultiRegionParameterGroupName) > 0 {
		input.MultiRegionParameterGroupName = aws.String(_memorydbMultiRegionParameterGroupName)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbSource) > 0 {
		input.Source = aws.String(_memorydbSource)
	}

	if resp, err := client.DescribeMultiRegionParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of parameter group descriptions. If a parameter group name is
// specified, the list contains only the descriptions for that group.
func memorydb_DescribeParameterGroups(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeParameterGroupsInput{}

	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeParameterGroupsOutput
	p := memorydb.NewDescribeParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular parameter group.
func memorydb_DescribeParameters(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeParametersInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeParametersOutput
	p := memorydb.NewDescribeParametersPaginator(client, input)
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

// Returns information about reserved nodes for this account, or about a specified
// reserved node.
func memorydb_DescribeReservedNodes(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeReservedNodesInput{}

	if len(_memorydbDuration) > 0 {
		input.Duration = aws.String(_memorydbDuration)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbNodeType) > 0 {
		input.NodeType = aws.String(_memorydbNodeType)
	}
	if len(_memorydbOfferingType) > 0 {
		input.OfferingType = aws.String(_memorydbOfferingType)
	}
	if len(_memorydbReservationId) > 0 {
		input.ReservationId = aws.String(_memorydbReservationId)
	}
	if len(_memorydbReservedNodesOfferingId) > 0 {
		input.ReservedNodesOfferingId = aws.String(_memorydbReservedNodesOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeReservedNodesOutput
	p := memorydb.NewDescribeReservedNodesPaginator(client, input)
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

// Lists available reserved node offerings.
func memorydb_DescribeReservedNodesOfferings(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeReservedNodesOfferingsInput{}

	if len(_memorydbDuration) > 0 {
		input.Duration = aws.String(_memorydbDuration)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbNodeType) > 0 {
		input.NodeType = aws.String(_memorydbNodeType)
	}
	if len(_memorydbOfferingType) > 0 {
		input.OfferingType = aws.String(_memorydbOfferingType)
	}
	if len(_memorydbReservedNodesOfferingId) > 0 {
		input.ReservedNodesOfferingId = aws.String(_memorydbReservedNodesOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedNodesOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeReservedNodesOfferingsOutput
	p := memorydb.NewDescribeReservedNodesOfferingsPaginator(client, input)
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

// Returns details of the service updates.
func memorydb_DescribeServiceUpdates(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeServiceUpdatesInput{}

	if len(_memorydbClusterNames) > 0 {
		input.ClusterNames = append([]string(nil), _memorydbClusterNames...)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbServiceUpdateName) > 0 {
		input.ServiceUpdateName = aws.String(_memorydbServiceUpdateName)
	}
	if len(_memorydbStatus) > 0 {
		if err := assignInputField(input, "Status", _memorydbStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeServiceUpdates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeServiceUpdatesOutput
	p := memorydb.NewDescribeServiceUpdatesPaginator(client, input)
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

// Returns information about cluster snapshots. By default, DescribeSnapshots
// lists all of your snapshots; it can optionally describe a single snapshot, or
// just the snapshots associated with a particular cluster.
func memorydb_DescribeSnapshots(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeSnapshotsInput{}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbShowDetail) > 0 {
		if err := assignInputField(input, "ShowDetail", _memorydbShowDetail); err != nil {
			log.Errorf("invalid --show-detail: %s", err.Error())
			return
		}
	}
	if len(_memorydbSnapshotName) > 0 {
		input.SnapshotName = aws.String(_memorydbSnapshotName)
	}
	if len(_memorydbSource) > 0 {
		input.Source = aws.String(_memorydbSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeSnapshotsOutput
	p := memorydb.NewDescribeSnapshotsPaginator(client, input)
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

// Returns a list of subnet group descriptions. If a subnet group name is
// specified, the list contains only the description of that group.
func memorydb_DescribeSubnetGroups(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeSubnetGroupsInput{}

	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_memorydbSubnetGroupName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSubnetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeSubnetGroupsOutput
	p := memorydb.NewDescribeSubnetGroupsPaginator(client, input)
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

// Returns a list of users.
func memorydb_DescribeUsers(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.DescribeUsersInput{}

	if len(_memorydbFilters) > 0 {
		if err := assignInputField(input, "Filters", _memorydbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_memorydbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _memorydbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_memorydbNextToken) > 0 {
		input.NextToken = aws.String(_memorydbNextToken)
	}
	if len(_memorydbUserName) > 0 {
		input.UserName = aws.String(_memorydbUserName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*memorydb.DescribeUsersOutput
	p := memorydb.NewDescribeUsersPaginator(client, input)
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

// Used to failover a shard. This API is designed for testing the behavior of your
// application in case of MemoryDB failover. It is not designed to be used as a
// production-level tool for initiating a failover to overcome a problem you may
// have with the cluster. Moreover, in certain conditions such as large scale
// operational events, Amazon may block this API.
func memorydb_FailoverShard(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.FailoverShardInput{
		// ClusterName: *string, // Required
		// ShardName: *string, // Required
	}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbShardName) > 0 {
		input.ShardName = aws.String(_memorydbShardName)
	}

	if resp, err := client.FailoverShard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the allowed updates for a multi-Region cluster.
func memorydb_ListAllowedMultiRegionClusterUpdates(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.ListAllowedMultiRegionClusterUpdatesInput{
		// MultiRegionClusterName: *string, // Required
	}

	if len(_memorydbMultiRegionClusterName) > 0 {
		input.MultiRegionClusterName = aws.String(_memorydbMultiRegionClusterName)
	}

	if resp, err := client.ListAllowedMultiRegionClusterUpdates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all available node types that you can scale to from your cluster's
// current node type.
//
// When you use the UpdateCluster operation to scale your cluster, the value of
// the NodeType parameter must be one of the node types returned by this operation.
func memorydb_ListAllowedNodeTypeUpdates(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.ListAllowedNodeTypeUpdatesInput{
		// ClusterName: *string, // Required
	}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}

	if resp, err := client.ListAllowedNodeTypeUpdates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all tags currently on a named resource. A tag is a key-value pair where
// the key and value are case-sensitive. You can use tags to categorize and track
// your MemoryDB resources. For more information, see [Tagging your MemoryDB resources].
//
// When you add or remove tags from multi region clusters, you might not
// immediately see the latest effective tags in the ListTags API response due to it
// being eventually consistent specifically for multi region clusters. For more
// information, see [Tagging your MemoryDB resources].
//
// [Tagging your MemoryDB resources]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/Tagging-Resources.html
func memorydb_ListTags(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.ListTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_memorydbResourceArn) > 0 {
		input.ResourceArn = aws.String(_memorydbResourceArn)
	}

	if resp, err := client.ListTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to purchase a reserved node offering. Reserved nodes are not
// eligible for cancellation and are non-refundable.
func memorydb_PurchaseReservedNodesOffering(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.PurchaseReservedNodesOfferingInput{
		// ReservedNodesOfferingId: *string, // Required
	}

	if len(_memorydbReservedNodesOfferingId) > 0 {
		input.ReservedNodesOfferingId = aws.String(_memorydbReservedNodesOfferingId)
	}
	if len(_memorydbNodeCount) > 0 {
		if err := assignInputField(input, "NodeCount", _memorydbNodeCount); err != nil {
			log.Errorf("invalid --node-count: %s", err.Error())
			return
		}
	}
	if len(_memorydbReservationId) > 0 {
		input.ReservationId = aws.String(_memorydbReservationId)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseReservedNodesOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a parameter group to the engine or system default
// value. You can reset specific parameters by submitting a list of parameter
// names. To reset the entire parameter group, specify the AllParameters and
// ParameterGroupName parameters.
func memorydb_ResetParameterGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.ResetParameterGroupInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}
	if len(_memorydbAllParameters) > 0 {
		if err := assignInputField(input, "AllParameters", _memorydbAllParameters); err != nil {
			log.Errorf("invalid --all-parameters: %s", err.Error())
			return
		}
	}
	if len(_memorydbParameterNames) > 0 {
		input.ParameterNames = append([]string(nil), _memorydbParameterNames...)
	}

	if resp, err := client.ResetParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to add tags to a resource. A tag is a key-value pair where
// the key and value are case-sensitive. You can use tags to categorize and track
// all your MemoryDB resources. For more information, see [Tagging your MemoryDB resources].
//
// When you add tags to multi region clusters, you might not immediately see the
// latest effective tags in the ListTags API response due to it being eventually
// consistent specifically for multi region clusters. For more information, see [Tagging your MemoryDB resources].
//
// You can specify cost-allocation tags for your MemoryDB resources, Amazon
// generates a cost allocation report as a comma-separated value (CSV) file with
// your usage and costs aggregated by your tags. You can apply tags that represent
// business categories (such as cost centers, application names, or owners) to
// organize your costs across multiple services.
//
// For more information, see [Using Cost Allocation Tags].
//
// [Tagging your MemoryDB resources]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/Tagging-Resources.html
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/tagging.html
func memorydb_TagResource(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_memorydbResourceArn) > 0 {
		input.ResourceArn = aws.String(_memorydbResourceArn)
	}
	if len(_memorydbTags) > 0 {
		if err := assignInputField(input, "Tags", _memorydbTags); err != nil {
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

// Use this operation to remove tags on a resource. A tag is a key-value pair
// where the key and value are case-sensitive. You can use tags to categorize and
// track all your MemoryDB resources. For more information, see [Tagging your MemoryDB resources].
//
// When you remove tags from multi region clusters, you might not immediately see
// the latest effective tags in the ListTags API response due to it being
// eventually consistent specifically for multi region clusters. For more
// information, see [Tagging your MemoryDB resources].
//
// You can specify cost-allocation tags for your MemoryDB resources, Amazon
// generates a cost allocation report as a comma-separated value (CSV) file with
// your usage and costs aggregated by your tags. You can apply tags that represent
// business categories (such as cost centers, application names, or owners) to
// organize your costs across multiple services.
//
// For more information, see [Using Cost Allocation Tags].
//
// [Tagging your MemoryDB resources]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/Tagging-Resources.html
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/tagging.html
func memorydb_UntagResource(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_memorydbResourceArn) > 0 {
		input.ResourceArn = aws.String(_memorydbResourceArn)
	}
	if len(_memorydbTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _memorydbTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the list of users that belong to the Access Control List.
func memorydb_UpdateACL(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UpdateACLInput{
		// ACLName: *string, // Required
	}

	if len(_memorydbACLName) > 0 {
		input.ACLName = aws.String(_memorydbACLName)
	}
	if len(_memorydbUserNamesToAdd) > 0 {
		input.UserNamesToAdd = append([]string(nil), _memorydbUserNamesToAdd...)
	}
	if len(_memorydbUserNamesToRemove) > 0 {
		input.UserNamesToRemove = append([]string(nil), _memorydbUserNamesToRemove...)
	}

	if resp, err := client.UpdateACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a cluster. You can use this operation to change one
// or more cluster configuration settings by specifying the settings and the new
// values.
func memorydb_UpdateCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UpdateClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_memorydbClusterName) > 0 {
		input.ClusterName = aws.String(_memorydbClusterName)
	}
	if len(_memorydbACLName) > 0 {
		input.ACLName = aws.String(_memorydbACLName)
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbEngine) > 0 {
		input.Engine = aws.String(_memorydbEngine)
	}
	if len(_memorydbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_memorydbEngineVersion)
	}
	if len(_memorydbIpDiscovery) > 0 {
		if err := assignInputField(input, "IpDiscovery", _memorydbIpDiscovery); err != nil {
			log.Errorf("invalid --ip-discovery: %s", err.Error())
			return
		}
	}
	if len(_memorydbMaintenanceWindow) > 0 {
		input.MaintenanceWindow = aws.String(_memorydbMaintenanceWindow)
	}
	if len(_memorydbNodeType) > 0 {
		input.NodeType = aws.String(_memorydbNodeType)
	}
	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}
	if len(_memorydbReplicaConfiguration) > 0 {
		if err := assignInputField(input, "ReplicaConfiguration", _memorydbReplicaConfiguration); err != nil {
			log.Errorf("invalid --replica-configuration: %s", err.Error())
			return
		}
	}
	if len(_memorydbSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _memorydbSecurityGroupIds...)
	}
	if len(_memorydbShardConfiguration) > 0 {
		if err := assignInputField(input, "ShardConfiguration", _memorydbShardConfiguration); err != nil {
			log.Errorf("invalid --shard-configuration: %s", err.Error())
			return
		}
	}
	if len(_memorydbSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _memorydbSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_memorydbSnapshotWindow) > 0 {
		input.SnapshotWindow = aws.String(_memorydbSnapshotWindow)
	}
	if len(_memorydbSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_memorydbSnsTopicArn)
	}
	if len(_memorydbSnsTopicStatus) > 0 {
		input.SnsTopicStatus = aws.String(_memorydbSnsTopicStatus)
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing multi-Region cluster.
func memorydb_UpdateMultiRegionCluster(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UpdateMultiRegionClusterInput{
		// MultiRegionClusterName: *string, // Required
	}

	if len(_memorydbMultiRegionClusterName) > 0 {
		input.MultiRegionClusterName = aws.String(_memorydbMultiRegionClusterName)
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_memorydbEngineVersion)
	}
	if len(_memorydbMultiRegionParameterGroupName) > 0 {
		input.MultiRegionParameterGroupName = aws.String(_memorydbMultiRegionParameterGroupName)
	}
	if len(_memorydbNodeType) > 0 {
		input.NodeType = aws.String(_memorydbNodeType)
	}
	if len(_memorydbShardConfiguration) > 0 {
		if err := assignInputField(input, "ShardConfiguration", _memorydbShardConfiguration); err != nil {
			log.Errorf("invalid --shard-configuration: %s", err.Error())
			return
		}
	}
	if len(_memorydbUpdateStrategy) > 0 {
		if err := assignInputField(input, "UpdateStrategy", _memorydbUpdateStrategy); err != nil {
			log.Errorf("invalid --update-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMultiRegionCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the parameters of a parameter group. You can modify up to 20 parameters
// in a single request by submitting a list parameter name and value pairs.
func memorydb_UpdateParameterGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UpdateParameterGroupInput{
		// ParameterGroupName: *string, // Required
		// ParameterNameValues: []types.ParameterNameValue, // Required
	}

	if len(_memorydbParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_memorydbParameterGroupName)
	}
	if len(_memorydbParameterNameValues) > 0 {
		if err := assignInputField(input, "ParameterNameValues", _memorydbParameterNameValues); err != nil {
			log.Errorf("invalid --parameter-name-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a subnet group. For more information, see [Updating a subnet group]
//
// [Updating a subnet group]: https://docs.aws.amazon.com/MemoryDB/latest/devguide/ubnetGroups.Modifying.html
func memorydb_UpdateSubnetGroup(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UpdateSubnetGroupInput{
		// SubnetGroupName: *string, // Required
	}

	if len(_memorydbSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_memorydbSubnetGroupName)
	}
	if len(_memorydbDescription) > 0 {
		input.Description = aws.String(_memorydbDescription)
	}
	if len(_memorydbSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _memorydbSubnetIds...)
	}

	if resp, err := client.UpdateSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes user password(s) and/or access string.
func memorydb_UpdateUser(cfg aws.Config, client *memorydb.Client) {
	input := &memorydb.UpdateUserInput{
		// UserName: *string, // Required
	}

	if len(_memorydbUserName) > 0 {
		input.UserName = aws.String(_memorydbUserName)
	}
	if len(_memorydbAccessString) > 0 {
		input.AccessString = aws.String(_memorydbAccessString)
	}
	if len(_memorydbAuthenticationMode) > 0 {
		if err := assignInputField(input, "AuthenticationMode", _memorydbAuthenticationMode); err != nil {
			log.Errorf("invalid --authentication-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_memorydbCmd)
	_memorydbCmd.Flags().SortFlags = false

	_memorydbCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_memorydbCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_memorydbCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_memorydbCmd.Flags().StringVarP(&_memorydbAccessString, "access-string", "", "", "Access String")
	_memorydbCmd.Flags().StringVarP(&_memorydbACLName, "acl-name", "", "", "ACL Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbAllParameters, "all-parameters", "", "", "All Parameters")
	_memorydbCmd.Flags().StringVarP(&_memorydbAuthenticationMode, "authentication-mode", "", "", "Authentication Mode")
	_memorydbCmd.Flags().StringVarP(&_memorydbAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_memorydbCmd.Flags().StringVarP(&_memorydbClusterName, "cluster-name", "", "", "Cluster Name")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbClusterNames, "cluster-names", "", nil, "Cluster Names")
	_memorydbCmd.Flags().StringVarP(&_memorydbDataTiering, "data-tiering", "", "", "Data Tiering")
	_memorydbCmd.Flags().StringVarP(&_memorydbDefaultOnly, "default-only", "", "", "Default Only")
	_memorydbCmd.Flags().StringVarP(&_memorydbDescription, "description", "", "", "Description")
	_memorydbCmd.Flags().StringVarP(&_memorydbDuration, "duration", "", "", "Duration")
	_memorydbCmd.Flags().StringVarP(&_memorydbEndTime, "end-time", "", "", "End Time")
	_memorydbCmd.Flags().StringVarP(&_memorydbEngine, "engine", "", "", "Engine")
	_memorydbCmd.Flags().StringVarP(&_memorydbEngineVersion, "engine-version", "", "", "Engine Version")
	_memorydbCmd.Flags().StringVarP(&_memorydbFamily, "family", "", "", "Family")
	_memorydbCmd.Flags().StringVarP(&_memorydbFilters, "filters", "", "", "Filters")
	_memorydbCmd.Flags().StringVarP(&_memorydbFinalSnapshotName, "final-snapshot-name", "", "", "Final Snapshot Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbIpDiscovery, "ip-discovery", "", "", "IP Discovery")
	_memorydbCmd.Flags().StringVarP(&_memorydbKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_memorydbCmd.Flags().StringVarP(&_memorydbMaintenanceWindow, "maintenance-window", "", "", "Maintenance Window")
	_memorydbCmd.Flags().StringVarP(&_memorydbMaxResults, "max-results", "", "", "Max Results")
	_memorydbCmd.Flags().StringVarP(&_memorydbMultiRegionClusterName, "multi-region-cluster-name", "", "", "Multi Region Cluster Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbMultiRegionClusterNameSuffix, "multi-region-cluster-name-suffix", "", "", "Multi Region Cluster Name Suffix")
	_memorydbCmd.Flags().StringVarP(&_memorydbMultiRegionParameterGroupName, "multi-region-parameter-group-name", "", "", "Multi Region Parameter Group Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbNetworkType, "network-type", "", "", "Network Type")
	_memorydbCmd.Flags().StringVarP(&_memorydbNextToken, "next-token", "", "", "Next Token")
	_memorydbCmd.Flags().StringVarP(&_memorydbNodeCount, "node-count", "", "", "Node Count")
	_memorydbCmd.Flags().StringVarP(&_memorydbNodeType, "node-type", "", "", "Node Type")
	_memorydbCmd.Flags().StringVarP(&_memorydbNumReplicasPerShard, "num-replicas-per-shard", "", "", "Num Replicas Per Shard")
	_memorydbCmd.Flags().StringVarP(&_memorydbNumShards, "num-shards", "", "", "Num Shards")
	_memorydbCmd.Flags().StringVarP(&_memorydbOfferingType, "offering-type", "", "", "Offering Type")
	_memorydbCmd.Flags().StringVarP(&_memorydbParameterGroupFamily, "parameter-group-family", "", "", "Parameter Group Family")
	_memorydbCmd.Flags().StringVarP(&_memorydbParameterGroupName, "parameter-group-name", "", "", "Parameter Group Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbParameterNameValues, "parameter-name-values", "", "", "Parameter Name Values")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbParameterNames, "parameter-names", "", nil, "Parameter Names")
	_memorydbCmd.Flags().StringVarP(&_memorydbPort, "port", "", "", "Port")
	_memorydbCmd.Flags().StringVarP(&_memorydbReplicaConfiguration, "replica-configuration", "", "", "Replica Configuration")
	_memorydbCmd.Flags().StringVarP(&_memorydbReservationId, "reservation-id", "", "", "Reservation ID")
	_memorydbCmd.Flags().StringVarP(&_memorydbReservedNodesOfferingId, "reserved-nodes-offering-id", "", "", "Reserved Nodes Offering ID")
	_memorydbCmd.Flags().StringVarP(&_memorydbResourceArn, "resource-arn", "", "", "Resource ARN")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_memorydbCmd.Flags().StringVarP(&_memorydbServiceUpdate, "service-update", "", "", "Service Update")
	_memorydbCmd.Flags().StringVarP(&_memorydbServiceUpdateName, "service-update-name", "", "", "Service Update Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbShardConfiguration, "shard-configuration", "", "", "Shard Configuration")
	_memorydbCmd.Flags().StringVarP(&_memorydbShardName, "shard-name", "", "", "Shard Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbShowClusterDetails, "show-cluster-details", "", "", "Show Cluster Details")
	_memorydbCmd.Flags().StringVarP(&_memorydbShowDetail, "show-detail", "", "", "Show Detail")
	_memorydbCmd.Flags().StringVarP(&_memorydbShowShardDetails, "show-shard-details", "", "", "Show Shard Details")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbSnapshotArns, "snapshot-arns", "", nil, "Snapshot Arns")
	_memorydbCmd.Flags().StringVarP(&_memorydbSnapshotName, "snapshot-name", "", "", "Snapshot Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbSnapshotRetentionLimit, "snapshot-retention-limit", "", "", "Snapshot Retention Limit")
	_memorydbCmd.Flags().StringVarP(&_memorydbSnapshotWindow, "snapshot-window", "", "", "Snapshot Window")
	_memorydbCmd.Flags().StringVarP(&_memorydbSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_memorydbCmd.Flags().StringVarP(&_memorydbSnsTopicStatus, "sns-topic-status", "", "", "SNS Topic Status")
	_memorydbCmd.Flags().StringVarP(&_memorydbSource, "source", "", "", "Source")
	_memorydbCmd.Flags().StringVarP(&_memorydbSourceName, "source-name", "", "", "Source Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbSourceSnapshotName, "source-snapshot-name", "", "", "Source Snapshot Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbSourceType, "source-type", "", "", "Source Type")
	_memorydbCmd.Flags().StringVarP(&_memorydbStartTime, "start-time", "", "", "Start Time")
	_memorydbCmd.Flags().StringVarP(&_memorydbStatus, "status", "", "", "Status")
	_memorydbCmd.Flags().StringVarP(&_memorydbSubnetGroupName, "subnet-group-name", "", "", "Subnet Group Name")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbTagKeys, "tag-keys", "", nil, "Tag Keys")
	_memorydbCmd.Flags().StringVarP(&_memorydbTags, "tags", "", "", "Tags")
	_memorydbCmd.Flags().StringVarP(&_memorydbTargetBucket, "target-bucket", "", "", "Target Bucket")
	_memorydbCmd.Flags().StringVarP(&_memorydbTargetSnapshotName, "target-snapshot-name", "", "", "Target Snapshot Name")
	_memorydbCmd.Flags().StringVarP(&_memorydbTLSEnabled, "tls-enabled", "", "", "TLS Enabled")
	_memorydbCmd.Flags().StringVarP(&_memorydbUpdateStrategy, "update-strategy", "", "", "Update Strategy")
	_memorydbCmd.Flags().StringVarP(&_memorydbUserName, "user-name", "", "", "User Name")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbUserNames, "user-names", "", nil, "User Names")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbUserNamesToAdd, "user-names-to-add", "", nil, "User Names To Add")
	_memorydbCmd.Flags().StringSliceVarP(&_memorydbUserNamesToRemove, "user-names-to-remove", "", nil, "User Names To Remove")

	_memorydbCmd.Flags().BoolVarP(&_memorydbBatchUpdateCluster, "batch-update-cluster", "", false, "Batch Update Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCopySnapshot, "copy-snapshot", "", false, "Copy Snapshot")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateACL, "create-acl", "", false, "Create ACL")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateCluster, "create-cluster", "", false, "Create Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateMultiRegionCluster, "create-multi-region-cluster", "", false, "Create Multi Region Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateParameterGroup, "create-parameter-group", "", false, "Create Parameter Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateSubnetGroup, "create-subnet-group", "", false, "Create Subnet Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbCreateUser, "create-user", "", false, "Create User")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteACL, "delete-acl", "", false, "Delete ACL")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteMultiRegionCluster, "delete-multi-region-cluster", "", false, "Delete Multi Region Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteParameterGroup, "delete-parameter-group", "", false, "Delete Parameter Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteSnapshot, "delete-snapshot", "", false, "Delete Snapshot")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteSubnetGroup, "delete-subnet-group", "", false, "Delete Subnet Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDeleteUser, "delete-user", "", false, "Delete User")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeACLs, "describe-acls", "", false, "Describe Acls")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeClusters, "describe-clusters", "", false, "Describe Clusters")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeEngineVersions, "describe-engine-versions", "", false, "Describe Engine Versions")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeEvents, "describe-events", "", false, "Describe Events")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeMultiRegionClusters, "describe-multi-region-clusters", "", false, "Describe Multi Region Clusters")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeMultiRegionParameterGroups, "describe-multi-region-parameter-groups", "", false, "Describe Multi Region Parameter Groups")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeMultiRegionParameters, "describe-multi-region-parameters", "", false, "Describe Multi Region Parameters")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeParameterGroups, "describe-parameter-groups", "", false, "Describe Parameter Groups")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeParameters, "describe-parameters", "", false, "Describe Parameters")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeReservedNodes, "describe-reserved-nodes", "", false, "Describe Reserved Nodes")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeReservedNodesOfferings, "describe-reserved-nodes-offerings", "", false, "Describe Reserved Nodes Offerings")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeServiceUpdates, "describe-service-updates", "", false, "Describe Service Updates")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeSnapshots, "describe-snapshots", "", false, "Describe Snapshots")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeSubnetGroups, "describe-subnet-groups", "", false, "Describe Subnet Groups")
	_memorydbCmd.Flags().BoolVarP(&_memorydbDescribeUsers, "describe-users", "", false, "Describe Users")
	_memorydbCmd.Flags().BoolVarP(&_memorydbFailoverShard, "failover-shard", "", false, "Failover Shard")
	_memorydbCmd.Flags().BoolVarP(&_memorydbListAllowedMultiRegionClusterUpdates, "list-allowed-multi-region-cluster-updates", "", false, "List Allowed Multi Region Cluster Updates")
	_memorydbCmd.Flags().BoolVarP(&_memorydbListAllowedNodeTypeUpdates, "list-allowed-node-type-updates", "", false, "List Allowed Node Type Updates")
	_memorydbCmd.Flags().BoolVarP(&_memorydbListTags, "list-tags", "", false, "List Tags")
	_memorydbCmd.Flags().BoolVarP(&_memorydbPurchaseReservedNodesOffering, "purchase-reserved-nodes-offering", "", false, "Purchase Reserved Nodes Offering")
	_memorydbCmd.Flags().BoolVarP(&_memorydbResetParameterGroup, "reset-parameter-group", "", false, "Reset Parameter Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbTagResource, "tag-resource", "", false, "Tag Resource")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUntagResource, "untag-resource", "", false, "Untag Resource")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUpdateACL, "update-acl", "", false, "Update ACL")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUpdateMultiRegionCluster, "update-multi-region-cluster", "", false, "Update Multi Region Cluster")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUpdateParameterGroup, "update-parameter-group", "", false, "Update Parameter Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUpdateSubnetGroup, "update-subnet-group", "", false, "Update Subnet Group")
	_memorydbCmd.Flags().BoolVarP(&_memorydbUpdateUser, "update-user", "", false, "Update User")

}
