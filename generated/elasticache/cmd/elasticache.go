package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// elasticacheCmd represents the elasticache command
var _elasticacheCmd = &cobra.Command{
	Use:   "elasticache",
	Short: "AWS elasticache CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := elasticache.NewFromConfig(cfg)
		if _elasticacheAddTagsToResource {
			elasticache_AddTagsToResource(cfg, client)
			return
		}
		if _elasticacheAuthorizeCacheSecurityGroupIngress {
			elasticache_AuthorizeCacheSecurityGroupIngress(cfg, client)
			return
		}
		if _elasticacheBatchApplyUpdateAction {
			elasticache_BatchApplyUpdateAction(cfg, client)
			return
		}
		if _elasticacheBatchStopUpdateAction {
			elasticache_BatchStopUpdateAction(cfg, client)
			return
		}
		if _elasticacheCompleteMigration {
			elasticache_CompleteMigration(cfg, client)
			return
		}
		if _elasticacheCopyServerlessCacheSnapshot {
			elasticache_CopyServerlessCacheSnapshot(cfg, client)
			return
		}
		if _elasticacheCopySnapshot {
			elasticache_CopySnapshot(cfg, client)
			return
		}
		if _elasticacheCreateCacheCluster {
			elasticache_CreateCacheCluster(cfg, client)
			return
		}
		if _elasticacheCreateCacheParameterGroup {
			elasticache_CreateCacheParameterGroup(cfg, client)
			return
		}
		if _elasticacheCreateCacheSecurityGroup {
			elasticache_CreateCacheSecurityGroup(cfg, client)
			return
		}
		if _elasticacheCreateCacheSubnetGroup {
			elasticache_CreateCacheSubnetGroup(cfg, client)
			return
		}
		if _elasticacheCreateGlobalReplicationGroup {
			elasticache_CreateGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheCreateReplicationGroup {
			elasticache_CreateReplicationGroup(cfg, client)
			return
		}
		if _elasticacheCreateServerlessCache {
			elasticache_CreateServerlessCache(cfg, client)
			return
		}
		if _elasticacheCreateServerlessCacheSnapshot {
			elasticache_CreateServerlessCacheSnapshot(cfg, client)
			return
		}
		if _elasticacheCreateSnapshot {
			elasticache_CreateSnapshot(cfg, client)
			return
		}
		if _elasticacheCreateUser {
			elasticache_CreateUser(cfg, client)
			return
		}
		if _elasticacheCreateUserGroup {
			elasticache_CreateUserGroup(cfg, client)
			return
		}
		if _elasticacheDecreaseNodeGroupsInGlobalReplicationGroup {
			elasticache_DecreaseNodeGroupsInGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheDecreaseReplicaCount {
			elasticache_DecreaseReplicaCount(cfg, client)
			return
		}
		if _elasticacheDeleteCacheCluster {
			elasticache_DeleteCacheCluster(cfg, client)
			return
		}
		if _elasticacheDeleteCacheParameterGroup {
			elasticache_DeleteCacheParameterGroup(cfg, client)
			return
		}
		if _elasticacheDeleteCacheSecurityGroup {
			elasticache_DeleteCacheSecurityGroup(cfg, client)
			return
		}
		if _elasticacheDeleteCacheSubnetGroup {
			elasticache_DeleteCacheSubnetGroup(cfg, client)
			return
		}
		if _elasticacheDeleteGlobalReplicationGroup {
			elasticache_DeleteGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheDeleteReplicationGroup {
			elasticache_DeleteReplicationGroup(cfg, client)
			return
		}
		if _elasticacheDeleteServerlessCache {
			elasticache_DeleteServerlessCache(cfg, client)
			return
		}
		if _elasticacheDeleteServerlessCacheSnapshot {
			elasticache_DeleteServerlessCacheSnapshot(cfg, client)
			return
		}
		if _elasticacheDeleteSnapshot {
			elasticache_DeleteSnapshot(cfg, client)
			return
		}
		if _elasticacheDeleteUser {
			elasticache_DeleteUser(cfg, client)
			return
		}
		if _elasticacheDeleteUserGroup {
			elasticache_DeleteUserGroup(cfg, client)
			return
		}
		if _elasticacheDescribeCacheClusters {
			elasticache_DescribeCacheClusters(cfg, client)
			return
		}
		if _elasticacheDescribeCacheEngineVersions {
			elasticache_DescribeCacheEngineVersions(cfg, client)
			return
		}
		if _elasticacheDescribeCacheParameterGroups {
			elasticache_DescribeCacheParameterGroups(cfg, client)
			return
		}
		if _elasticacheDescribeCacheParameters {
			elasticache_DescribeCacheParameters(cfg, client)
			return
		}
		if _elasticacheDescribeCacheSecurityGroups {
			elasticache_DescribeCacheSecurityGroups(cfg, client)
			return
		}
		if _elasticacheDescribeCacheSubnetGroups {
			elasticache_DescribeCacheSubnetGroups(cfg, client)
			return
		}
		if _elasticacheDescribeEngineDefaultParameters {
			elasticache_DescribeEngineDefaultParameters(cfg, client)
			return
		}
		if _elasticacheDescribeEvents {
			elasticache_DescribeEvents(cfg, client)
			return
		}
		if _elasticacheDescribeGlobalReplicationGroups {
			elasticache_DescribeGlobalReplicationGroups(cfg, client)
			return
		}
		if _elasticacheDescribeReplicationGroups {
			elasticache_DescribeReplicationGroups(cfg, client)
			return
		}
		if _elasticacheDescribeReservedCacheNodes {
			elasticache_DescribeReservedCacheNodes(cfg, client)
			return
		}
		if _elasticacheDescribeReservedCacheNodesOfferings {
			elasticache_DescribeReservedCacheNodesOfferings(cfg, client)
			return
		}
		if _elasticacheDescribeServerlessCacheSnapshots {
			elasticache_DescribeServerlessCacheSnapshots(cfg, client)
			return
		}
		if _elasticacheDescribeServerlessCaches {
			elasticache_DescribeServerlessCaches(cfg, client)
			return
		}
		if _elasticacheDescribeServiceUpdates {
			elasticache_DescribeServiceUpdates(cfg, client)
			return
		}
		if _elasticacheDescribeSnapshots {
			elasticache_DescribeSnapshots(cfg, client)
			return
		}
		if _elasticacheDescribeUpdateActions {
			elasticache_DescribeUpdateActions(cfg, client)
			return
		}
		if _elasticacheDescribeUserGroups {
			elasticache_DescribeUserGroups(cfg, client)
			return
		}
		if _elasticacheDescribeUsers {
			elasticache_DescribeUsers(cfg, client)
			return
		}
		if _elasticacheDisassociateGlobalReplicationGroup {
			elasticache_DisassociateGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheExportServerlessCacheSnapshot {
			elasticache_ExportServerlessCacheSnapshot(cfg, client)
			return
		}
		if _elasticacheFailoverGlobalReplicationGroup {
			elasticache_FailoverGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheIncreaseNodeGroupsInGlobalReplicationGroup {
			elasticache_IncreaseNodeGroupsInGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheIncreaseReplicaCount {
			elasticache_IncreaseReplicaCount(cfg, client)
			return
		}
		if _elasticacheListAllowedNodeTypeModifications {
			elasticache_ListAllowedNodeTypeModifications(cfg, client)
			return
		}
		if _elasticacheListTagsForResource {
			elasticache_ListTagsForResource(cfg, client)
			return
		}
		if _elasticacheModifyCacheCluster {
			elasticache_ModifyCacheCluster(cfg, client)
			return
		}
		if _elasticacheModifyCacheParameterGroup {
			elasticache_ModifyCacheParameterGroup(cfg, client)
			return
		}
		if _elasticacheModifyCacheSubnetGroup {
			elasticache_ModifyCacheSubnetGroup(cfg, client)
			return
		}
		if _elasticacheModifyGlobalReplicationGroup {
			elasticache_ModifyGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheModifyReplicationGroup {
			elasticache_ModifyReplicationGroup(cfg, client)
			return
		}
		if _elasticacheModifyReplicationGroupShardConfiguration {
			elasticache_ModifyReplicationGroupShardConfiguration(cfg, client)
			return
		}
		if _elasticacheModifyServerlessCache {
			elasticache_ModifyServerlessCache(cfg, client)
			return
		}
		if _elasticacheModifyUser {
			elasticache_ModifyUser(cfg, client)
			return
		}
		if _elasticacheModifyUserGroup {
			elasticache_ModifyUserGroup(cfg, client)
			return
		}
		if _elasticachePurchaseReservedCacheNodesOffering {
			elasticache_PurchaseReservedCacheNodesOffering(cfg, client)
			return
		}
		if _elasticacheRebalanceSlotsInGlobalReplicationGroup {
			elasticache_RebalanceSlotsInGlobalReplicationGroup(cfg, client)
			return
		}
		if _elasticacheRebootCacheCluster {
			elasticache_RebootCacheCluster(cfg, client)
			return
		}
		if _elasticacheRemoveTagsFromResource {
			elasticache_RemoveTagsFromResource(cfg, client)
			return
		}
		if _elasticacheResetCacheParameterGroup {
			elasticache_ResetCacheParameterGroup(cfg, client)
			return
		}
		if _elasticacheRevokeCacheSecurityGroupIngress {
			elasticache_RevokeCacheSecurityGroupIngress(cfg, client)
			return
		}
		if _elasticacheStartMigration {
			elasticache_StartMigration(cfg, client)
			return
		}
		if _elasticacheTestFailover {
			elasticache_TestFailover(cfg, client)
			return
		}
		if _elasticacheTestMigration {
			elasticache_TestMigration(cfg, client)
			return
		}

	},
}

var (
	_elasticacheAddTagsToResource                          bool
	_elasticacheAuthorizeCacheSecurityGroupIngress         bool
	_elasticacheBatchApplyUpdateAction                     bool
	_elasticacheBatchStopUpdateAction                      bool
	_elasticacheCompleteMigration                          bool
	_elasticacheCopyServerlessCacheSnapshot                bool
	_elasticacheCopySnapshot                               bool
	_elasticacheCreateCacheCluster                         bool
	_elasticacheCreateCacheParameterGroup                  bool
	_elasticacheCreateCacheSecurityGroup                   bool
	_elasticacheCreateCacheSubnetGroup                     bool
	_elasticacheCreateGlobalReplicationGroup               bool
	_elasticacheCreateReplicationGroup                     bool
	_elasticacheCreateServerlessCache                      bool
	_elasticacheCreateServerlessCacheSnapshot              bool
	_elasticacheCreateSnapshot                             bool
	_elasticacheCreateUser                                 bool
	_elasticacheCreateUserGroup                            bool
	_elasticacheDecreaseNodeGroupsInGlobalReplicationGroup bool
	_elasticacheDecreaseReplicaCount                       bool
	_elasticacheDeleteCacheCluster                         bool
	_elasticacheDeleteCacheParameterGroup                  bool
	_elasticacheDeleteCacheSecurityGroup                   bool
	_elasticacheDeleteCacheSubnetGroup                     bool
	_elasticacheDeleteGlobalReplicationGroup               bool
	_elasticacheDeleteReplicationGroup                     bool
	_elasticacheDeleteServerlessCache                      bool
	_elasticacheDeleteServerlessCacheSnapshot              bool
	_elasticacheDeleteSnapshot                             bool
	_elasticacheDeleteUser                                 bool
	_elasticacheDeleteUserGroup                            bool
	_elasticacheDescribeCacheClusters                      bool
	_elasticacheDescribeCacheEngineVersions                bool
	_elasticacheDescribeCacheParameterGroups               bool
	_elasticacheDescribeCacheParameters                    bool
	_elasticacheDescribeCacheSecurityGroups                bool
	_elasticacheDescribeCacheSubnetGroups                  bool
	_elasticacheDescribeEngineDefaultParameters            bool
	_elasticacheDescribeEvents                             bool
	_elasticacheDescribeGlobalReplicationGroups            bool
	_elasticacheDescribeReplicationGroups                  bool
	_elasticacheDescribeReservedCacheNodes                 bool
	_elasticacheDescribeReservedCacheNodesOfferings        bool
	_elasticacheDescribeServerlessCacheSnapshots           bool
	_elasticacheDescribeServerlessCaches                   bool
	_elasticacheDescribeServiceUpdates                     bool
	_elasticacheDescribeSnapshots                          bool
	_elasticacheDescribeUpdateActions                      bool
	_elasticacheDescribeUserGroups                         bool
	_elasticacheDescribeUsers                              bool
	_elasticacheDisassociateGlobalReplicationGroup         bool
	_elasticacheExportServerlessCacheSnapshot              bool
	_elasticacheFailoverGlobalReplicationGroup             bool
	_elasticacheIncreaseNodeGroupsInGlobalReplicationGroup bool
	_elasticacheIncreaseReplicaCount                       bool
	_elasticacheListAllowedNodeTypeModifications           bool
	_elasticacheListTagsForResource                        bool
	_elasticacheModifyCacheCluster                         bool
	_elasticacheModifyCacheParameterGroup                  bool
	_elasticacheModifyCacheSubnetGroup                     bool
	_elasticacheModifyGlobalReplicationGroup               bool
	_elasticacheModifyReplicationGroup                     bool
	_elasticacheModifyReplicationGroupShardConfiguration   bool
	_elasticacheModifyServerlessCache                      bool
	_elasticacheModifyUser                                 bool
	_elasticacheModifyUserGroup                            bool
	_elasticachePurchaseReservedCacheNodesOffering         bool
	_elasticacheRebalanceSlotsInGlobalReplicationGroup     bool
	_elasticacheRebootCacheCluster                         bool
	_elasticacheRemoveTagsFromResource                     bool
	_elasticacheResetCacheParameterGroup                   bool
	_elasticacheRevokeCacheSecurityGroupIngress            bool
	_elasticacheStartMigration                             bool
	_elasticacheTestFailover                               bool
	_elasticacheTestMigration                              bool

	_elasticacheAccessString                            string
	_elasticacheAppendAccessString                      string
	_elasticacheApplyImmediately                        string
	_elasticacheAtRestEncryptionEnabled                 string
	_elasticacheAuthToken                               string
	_elasticacheAuthTokenUpdateStrategy                 string
	_elasticacheAuthenticationMode                      string
	_elasticacheAutoMinorVersionUpgrade                 string
	_elasticacheAutomaticFailoverEnabled                string
	_elasticacheAZMode                                  string
	_elasticacheCacheClusterId                          string
	_elasticacheCacheClusterIds                         []string
	_elasticacheCacheNodeCount                          string
	_elasticacheCacheNodeIdsToReboot                    []string
	_elasticacheCacheNodeIdsToRemove                    []string
	_elasticacheCacheNodeType                           string
	_elasticacheCacheParameterGroupFamily               string
	_elasticacheCacheParameterGroupName                 string
	_elasticacheCacheSecurityGroupName                  string
	_elasticacheCacheSecurityGroupNames                 []string
	_elasticacheCacheSubnetGroupDescription             string
	_elasticacheCacheSubnetGroupName                    string
	_elasticacheCacheUsageLimits                        string
	_elasticacheClusterMode                             string
	_elasticacheCustomerNodeEndpointList                string
	_elasticacheDailySnapshotTime                       string
	_elasticacheDataTieringEnabled                      string
	_elasticacheDefaultOnly                             string
	_elasticacheDescription                             string
	_elasticacheDuration                                string
	_elasticacheEC2SecurityGroupName                    string
	_elasticacheEC2SecurityGroupOwnerId                 string
	_elasticacheEndTime                                 string
	_elasticacheEngine                                  string
	_elasticacheEngineVersion                           string
	_elasticacheFilters                                 string
	_elasticacheFinalSnapshotIdentifier                 string
	_elasticacheFinalSnapshotName                       string
	_elasticacheForce                                   string
	_elasticacheGlobalNodeGroupsToRemove                []string
	_elasticacheGlobalNodeGroupsToRetain                []string
	_elasticacheGlobalReplicationGroupDescription       string
	_elasticacheGlobalReplicationGroupId                string
	_elasticacheGlobalReplicationGroupIdSuffix          string
	_elasticacheIpDiscovery                             string
	_elasticacheKmsKeyId                                string
	_elasticacheLogDeliveryConfigurations               string
	_elasticacheMajorEngineVersion                      string
	_elasticacheMarker                                  string
	_elasticacheMaxRecords                              string
	_elasticacheMaxResults                              string
	_elasticacheMultiAZEnabled                          string
	_elasticacheNetworkType                             string
	_elasticacheNewAvailabilityZones                    []string
	_elasticacheNewReplicaCount                         string
	_elasticacheNextToken                               string
	_elasticacheNoPasswordRequired                      string
	_elasticacheNodeGroupConfiguration                  string
	_elasticacheNodeGroupCount                          string
	_elasticacheNodeGroupId                             string
	_elasticacheNodeGroupsToRemove                      []string
	_elasticacheNodeGroupsToRetain                      []string
	_elasticacheNotificationTopicArn                    string
	_elasticacheNotificationTopicStatus                 string
	_elasticacheNumCacheClusters                        string
	_elasticacheNumCacheNodes                           string
	_elasticacheNumNodeGroups                           string
	_elasticacheOfferingType                            string
	_elasticacheOutpostMode                             string
	_elasticacheParameterNameValues                     string
	_elasticachePasswords                               []string
	_elasticachePort                                    string
	_elasticachePreferredAvailabilityZone               string
	_elasticachePreferredAvailabilityZones              []string
	_elasticachePreferredCacheClusterAZs                []string
	_elasticachePreferredMaintenanceWindow              string
	_elasticachePreferredOutpostArn                     string
	_elasticachePreferredOutpostArns                    []string
	_elasticachePrimaryClusterId                        string
	_elasticachePrimaryRegion                           string
	_elasticachePrimaryReplicationGroupId               string
	_elasticacheProductDescription                      string
	_elasticacheRegionalConfigurations                  string
	_elasticacheRemoveUserGroup                         string
	_elasticacheRemoveUserGroups                        string
	_elasticacheReplicaConfiguration                    string
	_elasticacheReplicasPerNodeGroup                    string
	_elasticacheReplicasToRemove                        []string
	_elasticacheReplicationGroupDescription             string
	_elasticacheReplicationGroupId                      string
	_elasticacheReplicationGroupIds                     []string
	_elasticacheReplicationGroupRegion                  string
	_elasticacheReservedCacheNodeId                     string
	_elasticacheReservedCacheNodesOfferingId            string
	_elasticacheResetAllParameters                      string
	_elasticacheReshardingConfiguration                 string
	_elasticacheResourceName                            string
	_elasticacheRetainPrimaryCluster                    string
	_elasticacheRetainPrimaryReplicationGroup           string
	_elasticacheS3BucketName                            string
	_elasticacheScaleConfig                             string
	_elasticacheSecurityGroupIds                        []string
	_elasticacheServerlessCacheName                     string
	_elasticacheServerlessCacheSnapshotName             string
	_elasticacheServiceUpdateName                       string
	_elasticacheServiceUpdateStatus                     string
	_elasticacheServiceUpdateTimeRange                  string
	_elasticacheShowCacheClustersNotInReplicationGroups string
	_elasticacheShowCacheNodeInfo                       string
	_elasticacheShowMemberInfo                          string
	_elasticacheShowNodeGroupConfig                     string
	_elasticacheShowNodeLevelUpdateStatus               string
	_elasticacheSnapshotArns                            []string
	_elasticacheSnapshotArnsToRestore                   []string
	_elasticacheSnapshotName                            string
	_elasticacheSnapshotRetentionLimit                  string
	_elasticacheSnapshotSource                          string
	_elasticacheSnapshotType                            string
	_elasticacheSnapshotWindow                          string
	_elasticacheSnapshottingClusterId                   string
	_elasticacheSource                                  string
	_elasticacheSourceIdentifier                        string
	_elasticacheSourceServerlessCacheSnapshotName       string
	_elasticacheSourceSnapshotName                      string
	_elasticacheSourceType                              string
	_elasticacheStartTime                               string
	_elasticacheSubnetIds                               []string
	_elasticacheTagKeys                                 []string
	_elasticacheTags                                    string
	_elasticacheTargetBucket                            string
	_elasticacheTargetServerlessCacheSnapshotName       string
	_elasticacheTargetSnapshotName                      string
	_elasticacheTransitEncryptionEnabled                string
	_elasticacheTransitEncryptionMode                   string
	_elasticacheUpdateActionStatus                      string
	_elasticacheUserGroupId                             string
	_elasticacheUserGroupIds                            []string
	_elasticacheUserGroupIdsToAdd                       []string
	_elasticacheUserGroupIdsToRemove                    []string
	_elasticacheUserId                                  string
	_elasticacheUserIds                                 []string
	_elasticacheUserIdsToAdd                            []string
	_elasticacheUserIdsToRemove                         []string
	_elasticacheUserName                                string
)

// A tag is a key-value pair where the key and value are case-sensitive. You can
// use tags to categorize and track all your ElastiCache resources, with the
// exception of global replication group. When you add or remove tags on
// replication groups, those actions will be replicated to all nodes in the
// replication group. For more information, see [Resource-level permissions].
//
// For example, you can use cost-allocation tags to your ElastiCache resources,
// Amazon generates a cost allocation report as a comma-separated value (CSV) file
// with your usage and costs aggregated by your tags. You can apply tags that
// represent business categories (such as cost centers, application names, or
// owners) to organize your costs across multiple services.
//
// For more information, see [Using Cost Allocation Tags in Amazon ElastiCache] in the ElastiCache User Guide.
//
// [Using Cost Allocation Tags in Amazon ElastiCache]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Tagging.html
// [Resource-level permissions]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/IAM.ResourceLevelPermissions.html
func elasticache_AddTagsToResource(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.AddTagsToResourceInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_elasticacheResourceName) > 0 {
		input.ResourceName = aws.String(_elasticacheResourceName)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows network ingress to a cache security group. Applications using
// ElastiCache must be running on Amazon EC2, and Amazon EC2 security groups are
// used as the authorization mechanism.
//
// You cannot authorize ingress from an Amazon EC2 security group in one region to
// an ElastiCache cluster in another region.
func elasticache_AuthorizeCacheSecurityGroupIngress(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.AuthorizeCacheSecurityGroupIngressInput{
		// CacheSecurityGroupName: *string, // Required
		// EC2SecurityGroupName: *string, // Required
		// EC2SecurityGroupOwnerId: *string, // Required
	}

	if len(_elasticacheCacheSecurityGroupName) > 0 {
		input.CacheSecurityGroupName = aws.String(_elasticacheCacheSecurityGroupName)
	}
	if len(_elasticacheEC2SecurityGroupName) > 0 {
		input.EC2SecurityGroupName = aws.String(_elasticacheEC2SecurityGroupName)
	}
	if len(_elasticacheEC2SecurityGroupOwnerId) > 0 {
		input.EC2SecurityGroupOwnerId = aws.String(_elasticacheEC2SecurityGroupOwnerId)
	}

	if resp, err := client.AuthorizeCacheSecurityGroupIngress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Apply the service update. For more information on service updates and applying
// them, see [Applying Service Updates].
//
// [Applying Service Updates]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/applying-updates.html
func elasticache_BatchApplyUpdateAction(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.BatchApplyUpdateActionInput{
		// ServiceUpdateName: *string, // Required
	}

	if len(_elasticacheServiceUpdateName) > 0 {
		input.ServiceUpdateName = aws.String(_elasticacheServiceUpdateName)
	}
	if len(_elasticacheCacheClusterIds) > 0 {
		input.CacheClusterIds = append([]string(nil), _elasticacheCacheClusterIds...)
	}
	if len(_elasticacheReplicationGroupIds) > 0 {
		input.ReplicationGroupIds = append([]string(nil), _elasticacheReplicationGroupIds...)
	}

	if resp, err := client.BatchApplyUpdateAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop the service update. For more information on service updates and stopping
// them, see [Stopping Service Updates].
//
// [Stopping Service Updates]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/stopping-self-service-updates.html
func elasticache_BatchStopUpdateAction(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.BatchStopUpdateActionInput{
		// ServiceUpdateName: *string, // Required
	}

	if len(_elasticacheServiceUpdateName) > 0 {
		input.ServiceUpdateName = aws.String(_elasticacheServiceUpdateName)
	}
	if len(_elasticacheCacheClusterIds) > 0 {
		input.CacheClusterIds = append([]string(nil), _elasticacheCacheClusterIds...)
	}
	if len(_elasticacheReplicationGroupIds) > 0 {
		input.ReplicationGroupIds = append([]string(nil), _elasticacheReplicationGroupIds...)
	}

	if resp, err := client.BatchStopUpdateAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Complete the migration of data.
func elasticache_CompleteMigration(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CompleteMigrationInput{
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheForce) > 0 {
		if err := assignInputField(input, "Force", _elasticacheForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.CompleteMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a copy of an existing serverless cache’s snapshot. Available for
// Valkey, Redis OSS and Serverless Memcached only.
func elasticache_CopyServerlessCacheSnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CopyServerlessCacheSnapshotInput{
		// SourceServerlessCacheSnapshotName: *string, // Required
		// TargetServerlessCacheSnapshotName: *string, // Required
	}

	if len(_elasticacheSourceServerlessCacheSnapshotName) > 0 {
		input.SourceServerlessCacheSnapshotName = aws.String(_elasticacheSourceServerlessCacheSnapshotName)
	}
	if len(_elasticacheTargetServerlessCacheSnapshotName) > 0 {
		input.TargetServerlessCacheSnapshotName = aws.String(_elasticacheTargetServerlessCacheSnapshotName)
	}
	if len(_elasticacheKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_elasticacheKmsKeyId)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyServerlessCacheSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Makes a copy of an existing snapshot.
// This operation is valid for Valkey or Redis OSS only.
//
// Users or groups that have permissions to use the CopySnapshot operation can
// create their own Amazon S3 buckets and copy snapshots to it. To control access
// to your snapshots, use an IAM policy to control who has the ability to use the
// CopySnapshot operation. For more information about using IAM to control the use
// of ElastiCache operations, see [Exporting Snapshots]and [Authentication & Access Control].
//
// You could receive the following error messages.
//
// # Error Messages
//
// - Error Message: The S3 bucket %s is outside of the region.
//
// Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//
// more information, see [Step 1: Create an Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: The S3 bucket %s does not exist.
//
// Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//
// more information, see [Step 1: Create an Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: The S3 bucket %s is not owned by the authenticated user.
//
// Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//
// more information, see [Step 1: Create an Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: The authenticated user does not have sufficient permissions
// to perform the desired activity.
//
// Solution: Contact your system administrator to get the needed permissions.
//
// - Error Message: The S3 bucket %s already contains an object with key %s.
//
// Solution: Give the TargetSnapshotName a new and unique value. If exporting a
//
// snapshot, you could alternatively create a new Amazon S3 bucket and use this
// same value for TargetSnapshotName .
//
// - Error Message: ElastiCache has not been granted READ permissions %s on the
// S3 Bucket.
//
// Solution: Add List and Read permissions on the bucket. For more information,
//
// see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: ElastiCache has not been granted WRITE permissions %s on the
// S3 Bucket.
//
// Solution: Add Upload/Delete permissions on the bucket. For more information,
//
// see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: ElastiCache has not been granted READ_ACP permissions %s on
// the S3 Bucket.
//
// Solution: Add View Permissions on the bucket. For more information, see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]in the
//
// ElastiCache User Guide.
//
// [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-exporting.html#backups-exporting-grant-access
// [Exporting Snapshots]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-exporting.html
// [Authentication & Access Control]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/IAM.html
// [Step 1: Create an Amazon S3 Bucket]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-exporting.html#backups-exporting-create-s3-bucket
func elasticache_CopySnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CopySnapshotInput{
		// SourceSnapshotName: *string, // Required
		// TargetSnapshotName: *string, // Required
	}

	if len(_elasticacheSourceSnapshotName) > 0 {
		input.SourceSnapshotName = aws.String(_elasticacheSourceSnapshotName)
	}
	if len(_elasticacheTargetSnapshotName) > 0 {
		input.TargetSnapshotName = aws.String(_elasticacheTargetSnapshotName)
	}
	if len(_elasticacheKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_elasticacheKmsKeyId)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticacheTargetBucket) > 0 {
		input.TargetBucket = aws.String(_elasticacheTargetBucket)
	}

	if resp, err := client.CopySnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cluster. All nodes in the cluster run the same protocol-compliant
// cache engine software, either Memcached, Valkey or Redis OSS.
//
// This operation is not supported for Valkey or Redis OSS (cluster mode enabled)
// clusters.
func elasticache_CreateCacheCluster(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateCacheClusterInput{
		// CacheClusterId: *string, // Required
	}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheAZMode) > 0 {
		if err := assignInputField(input, "AZMode", _elasticacheAZMode); err != nil {
			log.Errorf("invalid --az-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAuthToken) > 0 {
		input.AuthToken = aws.String(_elasticacheAuthToken)
	}
	if len(_elasticacheAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _elasticacheAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheCacheSecurityGroupNames) > 0 {
		input.CacheSecurityGroupNames = append([]string(nil), _elasticacheCacheSecurityGroupNames...)
	}
	if len(_elasticacheCacheSubnetGroupName) > 0 {
		input.CacheSubnetGroupName = aws.String(_elasticacheCacheSubnetGroupName)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheEngineVersion) > 0 {
		input.EngineVersion = aws.String(_elasticacheEngineVersion)
	}
	if len(_elasticacheIpDiscovery) > 0 {
		if err := assignInputField(input, "IpDiscovery", _elasticacheIpDiscovery); err != nil {
			log.Errorf("invalid --ip-discovery: %s", err.Error())
			return
		}
	}
	if len(_elasticacheLogDeliveryConfigurations) > 0 {
		if err := assignInputField(input, "LogDeliveryConfigurations", _elasticacheLogDeliveryConfigurations); err != nil {
			log.Errorf("invalid --log-delivery-configurations: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _elasticacheNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNotificationTopicArn) > 0 {
		input.NotificationTopicArn = aws.String(_elasticacheNotificationTopicArn)
	}
	if len(_elasticacheNumCacheNodes) > 0 {
		if err := assignInputField(input, "NumCacheNodes", _elasticacheNumCacheNodes); err != nil {
			log.Errorf("invalid --num-cache-nodes: %s", err.Error())
			return
		}
	}
	if len(_elasticacheOutpostMode) > 0 {
		if err := assignInputField(input, "OutpostMode", _elasticacheOutpostMode); err != nil {
			log.Errorf("invalid --outpost-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticachePort) > 0 {
		if err := assignInputField(input, "Port", _elasticachePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_elasticachePreferredAvailabilityZone) > 0 {
		input.PreferredAvailabilityZone = aws.String(_elasticachePreferredAvailabilityZone)
	}
	if len(_elasticachePreferredAvailabilityZones) > 0 {
		input.PreferredAvailabilityZones = append([]string(nil), _elasticachePreferredAvailabilityZones...)
	}
	if len(_elasticachePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_elasticachePreferredMaintenanceWindow)
	}
	if len(_elasticachePreferredOutpostArn) > 0 {
		input.PreferredOutpostArn = aws.String(_elasticachePreferredOutpostArn)
	}
	if len(_elasticachePreferredOutpostArns) > 0 {
		input.PreferredOutpostArns = append([]string(nil), _elasticachePreferredOutpostArns...)
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _elasticacheSecurityGroupIds...)
	}
	if len(_elasticacheSnapshotArns) > 0 {
		input.SnapshotArns = append([]string(nil), _elasticacheSnapshotArns...)
	}
	if len(_elasticacheSnapshotName) > 0 {
		input.SnapshotName = aws.String(_elasticacheSnapshotName)
	}
	if len(_elasticacheSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _elasticacheSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSnapshotWindow) > 0 {
		input.SnapshotWindow = aws.String(_elasticacheSnapshotWindow)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticacheTransitEncryptionEnabled) > 0 {
		if err := assignInputField(input, "TransitEncryptionEnabled", _elasticacheTransitEncryptionEnabled); err != nil {
			log.Errorf("invalid --transit-encryption-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCacheCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon ElastiCache cache parameter group. An ElastiCache cache
// parameter group is a collection of parameters and their values that are applied
// to all of the nodes in any cluster or replication group using the
// CacheParameterGroup.
//
// A newly created CacheParameterGroup is an exact duplicate of the default
// parameter group for the CacheParameterGroupFamily. To customize the newly
// created CacheParameterGroup you can change the values of specific parameters.
// For more information, see:
//
// [ModifyCacheParameterGroup]
// - in the ElastiCache API Reference.
//
// [Parameters and Parameter Groups]
// - in the ElastiCache User Guide.
//
// [ModifyCacheParameterGroup]: https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html
// [Parameters and Parameter Groups]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.html
func elasticache_CreateCacheParameterGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateCacheParameterGroupInput{
		// CacheParameterGroupFamily: *string, // Required
		// CacheParameterGroupName: *string, // Required
		// Description: *string, // Required
	}

	if len(_elasticacheCacheParameterGroupFamily) > 0 {
		input.CacheParameterGroupFamily = aws.String(_elasticacheCacheParameterGroupFamily)
	}
	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheDescription) > 0 {
		input.Description = aws.String(_elasticacheDescription)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCacheParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cache security group. Use a cache security group to control
// access to one or more clusters.
//
// Cache security groups are only used when you are creating a cluster outside of
// an Amazon Virtual Private Cloud (Amazon VPC). If you are creating a cluster
// inside of a VPC, use a cache subnet group instead. For more information, see [CreateCacheSubnetGroup].
//
// [CreateCacheSubnetGroup]: https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSubnetGroup.html
func elasticache_CreateCacheSecurityGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateCacheSecurityGroupInput{
		// CacheSecurityGroupName: *string, // Required
		// Description: *string, // Required
	}

	if len(_elasticacheCacheSecurityGroupName) > 0 {
		input.CacheSecurityGroupName = aws.String(_elasticacheCacheSecurityGroupName)
	}
	if len(_elasticacheDescription) > 0 {
		input.Description = aws.String(_elasticacheDescription)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCacheSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cache subnet group.
// Use this parameter only when you are creating a cluster in an Amazon Virtual
// Private Cloud (Amazon VPC).
func elasticache_CreateCacheSubnetGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateCacheSubnetGroupInput{
		// CacheSubnetGroupDescription: *string, // Required
		// CacheSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_elasticacheCacheSubnetGroupDescription) > 0 {
		input.CacheSubnetGroupDescription = aws.String(_elasticacheCacheSubnetGroupDescription)
	}
	if len(_elasticacheCacheSubnetGroupName) > 0 {
		input.CacheSubnetGroupName = aws.String(_elasticacheCacheSubnetGroupName)
	}
	if len(_elasticacheSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _elasticacheSubnetIds...)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCacheSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Global Datastore offers fully managed, fast, reliable and secure cross-region
// replication. Using Global Datastore with Valkey or Redis OSS, you can create
// cross-region read replica clusters for ElastiCache to enable low-latency reads
// and disaster recovery across regions. For more information, see [Replication Across Regions Using Global Datastore].
//
// - The GlobalReplicationGroupIdSuffix is the name of the Global datastore.
//
// - The PrimaryReplicationGroupId represents the name of the primary cluster
// that accepts writes and will replicate updates to the secondary cluster.
//
// [Replication Across Regions Using Global Datastore]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Redis-Global-Datastore.html
func elasticache_CreateGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateGlobalReplicationGroupInput{
		// GlobalReplicationGroupIdSuffix: *string, // Required
		// PrimaryReplicationGroupId: *string, // Required
	}

	if len(_elasticacheGlobalReplicationGroupIdSuffix) > 0 {
		input.GlobalReplicationGroupIdSuffix = aws.String(_elasticacheGlobalReplicationGroupIdSuffix)
	}
	if len(_elasticachePrimaryReplicationGroupId) > 0 {
		input.PrimaryReplicationGroupId = aws.String(_elasticachePrimaryReplicationGroupId)
	}
	if len(_elasticacheGlobalReplicationGroupDescription) > 0 {
		input.GlobalReplicationGroupDescription = aws.String(_elasticacheGlobalReplicationGroupDescription)
	}

	if resp, err := client.CreateGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Valkey or Redis OSS (cluster mode disabled) or a Valkey or Redis OSS
// (cluster mode enabled) replication group.
//
// This API can be used to create a standalone regional replication group or a
// secondary replication group associated with a Global datastore.
//
// A Valkey or Redis OSS (cluster mode disabled) replication group is a collection
// of nodes, where one of the nodes is a read/write primary and the others are
// read-only replicas. Writes to the primary are asynchronously propagated to the
// replicas.
//
// A Valkey or Redis OSS cluster-mode enabled cluster is comprised of from 1 to 90
// shards (API/CLI: node groups). Each shard has a primary node and up to 5
// read-only replica nodes. The configuration can range from 90 shards and 0
// replicas to 15 shards and 5 replicas, which is the maximum number or replicas
// allowed.
//
// The node or shard limit can be increased to a maximum of 500 per cluster if the
// Valkey or Redis OSS engine version is 5.0.6 or higher. For example, you can
// choose to configure a 500 node cluster that ranges between 83 shards (one
// primary and 5 replicas per shard) and 500 shards (single primary and no
// replicas). Make sure there are enough available IP addresses to accommodate the
// increase. Common pitfalls include the subnets in the subnet group have too small
// a CIDR range or the subnets are shared and heavily used by other clusters. For
// more information, see [Creating a Subnet Group]. For versions below 5.0.6, the limit is 250 per cluster.
//
// To request a limit increase, see [Amazon Service Limits] and choose the limit type Nodes per cluster
// per instance type.
//
// When a Valkey or Redis OSS (cluster mode disabled) replication group has been
// successfully created, you can add one or more read replicas to it, up to a total
// of 5 read replicas. If you need to increase or decrease the number of node
// groups (console: shards), you can use scaling. For more information, see [Scaling self-designed clusters]in the
// ElastiCache User Guide.
//
// This operation is valid for Valkey and Redis OSS only.
//
// [Scaling self-designed clusters]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Scaling.html
// [Creating a Subnet Group]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/SubnetGroups.Creating.html
// [Amazon Service Limits]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
func elasticache_CreateReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateReplicationGroupInput{
		// ReplicationGroupDescription: *string, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheReplicationGroupDescription) > 0 {
		input.ReplicationGroupDescription = aws.String(_elasticacheReplicationGroupDescription)
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheAtRestEncryptionEnabled) > 0 {
		if err := assignInputField(input, "AtRestEncryptionEnabled", _elasticacheAtRestEncryptionEnabled); err != nil {
			log.Errorf("invalid --at-rest-encryption-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAuthToken) > 0 {
		input.AuthToken = aws.String(_elasticacheAuthToken)
	}
	if len(_elasticacheAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _elasticacheAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAutomaticFailoverEnabled) > 0 {
		if err := assignInputField(input, "AutomaticFailoverEnabled", _elasticacheAutomaticFailoverEnabled); err != nil {
			log.Errorf("invalid --automatic-failover-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheCacheSecurityGroupNames) > 0 {
		input.CacheSecurityGroupNames = append([]string(nil), _elasticacheCacheSecurityGroupNames...)
	}
	if len(_elasticacheCacheSubnetGroupName) > 0 {
		input.CacheSubnetGroupName = aws.String(_elasticacheCacheSubnetGroupName)
	}
	if len(_elasticacheClusterMode) > 0 {
		if err := assignInputField(input, "ClusterMode", _elasticacheClusterMode); err != nil {
			log.Errorf("invalid --cluster-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheDataTieringEnabled) > 0 {
		if err := assignInputField(input, "DataTieringEnabled", _elasticacheDataTieringEnabled); err != nil {
			log.Errorf("invalid --data-tiering-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheEngineVersion) > 0 {
		input.EngineVersion = aws.String(_elasticacheEngineVersion)
	}
	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheIpDiscovery) > 0 {
		if err := assignInputField(input, "IpDiscovery", _elasticacheIpDiscovery); err != nil {
			log.Errorf("invalid --ip-discovery: %s", err.Error())
			return
		}
	}
	if len(_elasticacheKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_elasticacheKmsKeyId)
	}
	if len(_elasticacheLogDeliveryConfigurations) > 0 {
		if err := assignInputField(input, "LogDeliveryConfigurations", _elasticacheLogDeliveryConfigurations); err != nil {
			log.Errorf("invalid --log-delivery-configurations: %s", err.Error())
			return
		}
	}
	if len(_elasticacheMultiAZEnabled) > 0 {
		if err := assignInputField(input, "MultiAZEnabled", _elasticacheMultiAZEnabled); err != nil {
			log.Errorf("invalid --multi-az-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _elasticacheNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNodeGroupConfiguration) > 0 {
		if err := assignInputField(input, "NodeGroupConfiguration", _elasticacheNodeGroupConfiguration); err != nil {
			log.Errorf("invalid --node-group-configuration: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNotificationTopicArn) > 0 {
		input.NotificationTopicArn = aws.String(_elasticacheNotificationTopicArn)
	}
	if len(_elasticacheNumCacheClusters) > 0 {
		if err := assignInputField(input, "NumCacheClusters", _elasticacheNumCacheClusters); err != nil {
			log.Errorf("invalid --num-cache-clusters: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNumNodeGroups) > 0 {
		if err := assignInputField(input, "NumNodeGroups", _elasticacheNumNodeGroups); err != nil {
			log.Errorf("invalid --num-node-groups: %s", err.Error())
			return
		}
	}
	if len(_elasticachePort) > 0 {
		if err := assignInputField(input, "Port", _elasticachePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_elasticachePreferredCacheClusterAZs) > 0 {
		input.PreferredCacheClusterAZs = append([]string(nil), _elasticachePreferredCacheClusterAZs...)
	}
	if len(_elasticachePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_elasticachePreferredMaintenanceWindow)
	}
	if len(_elasticachePrimaryClusterId) > 0 {
		input.PrimaryClusterId = aws.String(_elasticachePrimaryClusterId)
	}
	if len(_elasticacheReplicasPerNodeGroup) > 0 {
		if err := assignInputField(input, "ReplicasPerNodeGroup", _elasticacheReplicasPerNodeGroup); err != nil {
			log.Errorf("invalid --replicas-per-node-group: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _elasticacheSecurityGroupIds...)
	}
	if len(_elasticacheServerlessCacheSnapshotName) > 0 {
		input.ServerlessCacheSnapshotName = aws.String(_elasticacheServerlessCacheSnapshotName)
	}
	if len(_elasticacheSnapshotArns) > 0 {
		input.SnapshotArns = append([]string(nil), _elasticacheSnapshotArns...)
	}
	if len(_elasticacheSnapshotName) > 0 {
		input.SnapshotName = aws.String(_elasticacheSnapshotName)
	}
	if len(_elasticacheSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _elasticacheSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSnapshotWindow) > 0 {
		input.SnapshotWindow = aws.String(_elasticacheSnapshotWindow)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticacheTransitEncryptionEnabled) > 0 {
		if err := assignInputField(input, "TransitEncryptionEnabled", _elasticacheTransitEncryptionEnabled); err != nil {
			log.Errorf("invalid --transit-encryption-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheTransitEncryptionMode) > 0 {
		if err := assignInputField(input, "TransitEncryptionMode", _elasticacheTransitEncryptionMode); err != nil {
			log.Errorf("invalid --transit-encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserGroupIds) > 0 {
		input.UserGroupIds = append([]string(nil), _elasticacheUserGroupIds...)
	}

	if resp, err := client.CreateReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a serverless cache.
func elasticache_CreateServerlessCache(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateServerlessCacheInput{
		// Engine: *string, // Required
		// ServerlessCacheName: *string, // Required
	}

	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheServerlessCacheName) > 0 {
		input.ServerlessCacheName = aws.String(_elasticacheServerlessCacheName)
	}
	if len(_elasticacheCacheUsageLimits) > 0 {
		if err := assignInputField(input, "CacheUsageLimits", _elasticacheCacheUsageLimits); err != nil {
			log.Errorf("invalid --cache-usage-limits: %s", err.Error())
			return
		}
	}
	if len(_elasticacheDailySnapshotTime) > 0 {
		input.DailySnapshotTime = aws.String(_elasticacheDailySnapshotTime)
	}
	if len(_elasticacheDescription) > 0 {
		input.Description = aws.String(_elasticacheDescription)
	}
	if len(_elasticacheKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_elasticacheKmsKeyId)
	}
	if len(_elasticacheMajorEngineVersion) > 0 {
		input.MajorEngineVersion = aws.String(_elasticacheMajorEngineVersion)
	}
	if len(_elasticacheSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _elasticacheSecurityGroupIds...)
	}
	if len(_elasticacheSnapshotArnsToRestore) > 0 {
		input.SnapshotArnsToRestore = append([]string(nil), _elasticacheSnapshotArnsToRestore...)
	}
	if len(_elasticacheSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _elasticacheSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _elasticacheSubnetIds...)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserGroupId) > 0 {
		input.UserGroupId = aws.String(_elasticacheUserGroupId)
	}

	if resp, err := client.CreateServerlessCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API creates a copy of an entire ServerlessCache at a specific moment in
// time. Available for Valkey, Redis OSS and Serverless Memcached only.
func elasticache_CreateServerlessCacheSnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateServerlessCacheSnapshotInput{
		// ServerlessCacheName: *string, // Required
		// ServerlessCacheSnapshotName: *string, // Required
	}

	if len(_elasticacheServerlessCacheName) > 0 {
		input.ServerlessCacheName = aws.String(_elasticacheServerlessCacheName)
	}
	if len(_elasticacheServerlessCacheSnapshotName) > 0 {
		input.ServerlessCacheSnapshotName = aws.String(_elasticacheServerlessCacheSnapshotName)
	}
	if len(_elasticacheKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_elasticacheKmsKeyId)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServerlessCacheSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a copy of an entire cluster or replication group at a specific moment
// in time.
//
// This operation is valid for Valkey or Redis OSS only.
func elasticache_CreateSnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateSnapshotInput{
		// SnapshotName: *string, // Required
	}

	if len(_elasticacheSnapshotName) > 0 {
		input.SnapshotName = aws.String(_elasticacheSnapshotName)
	}
	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_elasticacheKmsKeyId)
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
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

// For Valkey engine version 7.2 onwards and Redis OSS 6.0 to 7.1: Creates a user.
// For more information, see [Using Role Based Access Control (RBAC)].
//
// [Using Role Based Access Control (RBAC)]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html
func elasticache_CreateUser(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateUserInput{
		// AccessString: *string, // Required
		// Engine: *string, // Required
		// UserId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_elasticacheAccessString) > 0 {
		input.AccessString = aws.String(_elasticacheAccessString)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheUserId) > 0 {
		input.UserId = aws.String(_elasticacheUserId)
	}
	if len(_elasticacheUserName) > 0 {
		input.UserName = aws.String(_elasticacheUserName)
	}
	if len(_elasticacheAuthenticationMode) > 0 {
		if err := assignInputField(input, "AuthenticationMode", _elasticacheAuthenticationMode); err != nil {
			log.Errorf("invalid --authentication-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNoPasswordRequired) > 0 {
		if err := assignInputField(input, "NoPasswordRequired", _elasticacheNoPasswordRequired); err != nil {
			log.Errorf("invalid --no-password-required: %s", err.Error())
			return
		}
	}
	if len(_elasticachePasswords) > 0 {
		input.Passwords = append([]string(nil), _elasticachePasswords...)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
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

// For Valkey engine version 7.2 onwards and Redis OSS 6.0 to 7.1: Creates a user
// group. For more information, see [Using Role Based Access Control (RBAC)]
//
// [Using Role Based Access Control (RBAC)]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html
func elasticache_CreateUserGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.CreateUserGroupInput{
		// Engine: *string, // Required
		// UserGroupId: *string, // Required
	}

	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheUserGroupId) > 0 {
		input.UserGroupId = aws.String(_elasticacheUserGroupId)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserIds) > 0 {
		input.UserIds = append([]string(nil), _elasticacheUserIds...)
	}

	if resp, err := client.CreateUserGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Decreases the number of node groups in a Global datastore
func elasticache_DecreaseNodeGroupsInGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput{
		// ApplyImmediately: *bool, // Required
		// GlobalReplicationGroupId: *string, // Required
		// NodeGroupCount: *int32, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheNodeGroupCount) > 0 {
		if err := assignInputField(input, "NodeGroupCount", _elasticacheNodeGroupCount); err != nil {
			log.Errorf("invalid --node-group-count: %s", err.Error())
			return
		}
	}
	if len(_elasticacheGlobalNodeGroupsToRemove) > 0 {
		input.GlobalNodeGroupsToRemove = append([]string(nil), _elasticacheGlobalNodeGroupsToRemove...)
	}
	if len(_elasticacheGlobalNodeGroupsToRetain) > 0 {
		input.GlobalNodeGroupsToRetain = append([]string(nil), _elasticacheGlobalNodeGroupsToRetain...)
	}

	if resp, err := client.DecreaseNodeGroupsInGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dynamically decreases the number of replicas in a Valkey or Redis OSS (cluster
// mode disabled) replication group or the number of replica nodes in one or more
// node groups (shards) of a Valkey or Redis OSS (cluster mode enabled) replication
// group. This operation is performed with no cluster down time.
func elasticache_DecreaseReplicaCount(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DecreaseReplicaCountInput{
		// ApplyImmediately: *bool, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheNewReplicaCount) > 0 {
		if err := assignInputField(input, "NewReplicaCount", _elasticacheNewReplicaCount); err != nil {
			log.Errorf("invalid --new-replica-count: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicaConfiguration) > 0 {
		if err := assignInputField(input, "ReplicaConfiguration", _elasticacheReplicaConfiguration); err != nil {
			log.Errorf("invalid --replica-configuration: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicasToRemove) > 0 {
		input.ReplicasToRemove = append([]string(nil), _elasticacheReplicasToRemove...)
	}

	if resp, err := client.DecreaseReplicaCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously provisioned cluster. DeleteCacheCluster deletes all
// associated cache nodes, node endpoints and the cluster itself. When you receive
// a successful response from this operation, Amazon ElastiCache immediately begins
// deleting the cluster; you cannot cancel or revert this operation.
//
// This operation is not valid for:
//
// - Valkey or Redis OSS (cluster mode enabled) clusters
//
// - Valkey or Redis OSS (cluster mode disabled) clusters
//
// - A cluster that is the last read replica of a replication group
//
// - A cluster that is the primary node of a replication group
//
// - A node group (shard) that has Multi-AZ mode enabled
//
// - A cluster from a Valkey or Redis OSS (cluster mode enabled) replication
// group
//
// - A cluster that is not in the available state
func elasticache_DeleteCacheCluster(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteCacheClusterInput{
		// CacheClusterId: *string, // Required
	}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheFinalSnapshotIdentifier) > 0 {
		input.FinalSnapshotIdentifier = aws.String(_elasticacheFinalSnapshotIdentifier)
	}

	if resp, err := client.DeleteCacheCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified cache parameter group. You cannot delete a cache
// parameter group if it is associated with any cache clusters. You cannot delete
// the default cache parameter groups in your account.
func elasticache_DeleteCacheParameterGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteCacheParameterGroupInput{
		// CacheParameterGroupName: *string, // Required
	}

	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}

	if resp, err := client.DeleteCacheParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cache security group.
// You cannot delete a cache security group if it is associated with any clusters.
func elasticache_DeleteCacheSecurityGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteCacheSecurityGroupInput{
		// CacheSecurityGroupName: *string, // Required
	}

	if len(_elasticacheCacheSecurityGroupName) > 0 {
		input.CacheSecurityGroupName = aws.String(_elasticacheCacheSecurityGroupName)
	}

	if resp, err := client.DeleteCacheSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cache subnet group.
// You cannot delete a default cache subnet group or one that is associated with
// any clusters.
func elasticache_DeleteCacheSubnetGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteCacheSubnetGroupInput{
		// CacheSubnetGroupName: *string, // Required
	}

	if len(_elasticacheCacheSubnetGroupName) > 0 {
		input.CacheSubnetGroupName = aws.String(_elasticacheCacheSubnetGroupName)
	}

	if resp, err := client.DeleteCacheSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deleting a Global datastore is a two-step process:
// - First, you must DisassociateGlobalReplicationGroupto remove the secondary clusters in the Global datastore.
//
// - Once the Global datastore contains only the primary cluster, you can use
// the DeleteGlobalReplicationGroup API to delete the Global datastore while
// retainining the primary cluster using RetainPrimaryReplicationGroup=true .
//
// Since the Global Datastore has only a primary cluster, you can delete the
// Global Datastore while retaining the primary by setting
// RetainPrimaryReplicationGroup=true . The primary cluster is never deleted when
// deleting a Global Datastore. It can only be deleted when it no longer is
// associated with any Global Datastore.
//
// When you receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or revert
// this operation.
func elasticache_DeleteGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteGlobalReplicationGroupInput{
		// GlobalReplicationGroupId: *string, // Required
		// RetainPrimaryReplicationGroup: *bool, // Required
	}

	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheRetainPrimaryReplicationGroup) > 0 {
		if err := assignInputField(input, "RetainPrimaryReplicationGroup", _elasticacheRetainPrimaryReplicationGroup); err != nil {
			log.Errorf("invalid --retain-primary-replication-group: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing replication group. By default, this operation deletes the
// entire replication group, including the primary/primaries and all of the read
// replicas. If the replication group has only one primary, you can optionally
// delete only the read replicas, while retaining the primary by setting
// RetainPrimaryCluster=true .
//
// When you receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or revert
// this operation.
//
// - CreateSnapshot permission is required to create a final snapshot. Without
// this permission, the API call will fail with an Access Denied exception.
//
// - This operation is valid for Redis OSS only.
func elasticache_DeleteReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteReplicationGroupInput{
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheFinalSnapshotIdentifier) > 0 {
		input.FinalSnapshotIdentifier = aws.String(_elasticacheFinalSnapshotIdentifier)
	}
	if len(_elasticacheRetainPrimaryCluster) > 0 {
		if err := assignInputField(input, "RetainPrimaryCluster", _elasticacheRetainPrimaryCluster); err != nil {
			log.Errorf("invalid --retain-primary-cluster: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified existing serverless cache.
// CreateServerlessCacheSnapshot permission is required to create a final
// snapshot. Without this permission, the API call will fail with an Access Denied
// exception.
func elasticache_DeleteServerlessCache(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteServerlessCacheInput{
		// ServerlessCacheName: *string, // Required
	}

	if len(_elasticacheServerlessCacheName) > 0 {
		input.ServerlessCacheName = aws.String(_elasticacheServerlessCacheName)
	}
	if len(_elasticacheFinalSnapshotName) > 0 {
		input.FinalSnapshotName = aws.String(_elasticacheFinalSnapshotName)
	}

	if resp, err := client.DeleteServerlessCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing serverless cache snapshot. Available for Valkey, Redis OSS
// and Serverless Memcached only.
func elasticache_DeleteServerlessCacheSnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteServerlessCacheSnapshotInput{
		// ServerlessCacheSnapshotName: *string, // Required
	}

	if len(_elasticacheServerlessCacheSnapshotName) > 0 {
		input.ServerlessCacheSnapshotName = aws.String(_elasticacheServerlessCacheSnapshotName)
	}

	if resp, err := client.DeleteServerlessCacheSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing snapshot. When you receive a successful response from this
// operation, ElastiCache immediately begins deleting the snapshot; you cannot
// cancel or revert this operation.
//
// This operation is valid for Valkey or Redis OSS only.
func elasticache_DeleteSnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteSnapshotInput{
		// SnapshotName: *string, // Required
	}

	if len(_elasticacheSnapshotName) > 0 {
		input.SnapshotName = aws.String(_elasticacheSnapshotName)
	}

	if resp, err := client.DeleteSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For Valkey engine version 7.2 onwards and Redis OSS 6.0 onwards: Deletes a
// user. The user will be removed from all user groups and in turn removed from all
// replication groups. For more information, see [Using Role Based Access Control (RBAC)].
//
// [Using Role Based Access Control (RBAC)]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html
func elasticache_DeleteUser(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteUserInput{
		// UserId: *string, // Required
	}

	if len(_elasticacheUserId) > 0 {
		input.UserId = aws.String(_elasticacheUserId)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For Valkey engine version 7.2 onwards and Redis OSS 6.0 onwards: Deletes a user
// group. The user group must first be disassociated from the replication group
// before it can be deleted. For more information, see [Using Role Based Access Control (RBAC)].
//
// [Using Role Based Access Control (RBAC)]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html
func elasticache_DeleteUserGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DeleteUserGroupInput{
		// UserGroupId: *string, // Required
	}

	if len(_elasticacheUserGroupId) > 0 {
		input.UserGroupId = aws.String(_elasticacheUserGroupId)
	}

	if resp, err := client.DeleteUserGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all provisioned clusters if no cluster identifier is
// specified, or about a specific cache cluster if a cluster identifier is
// supplied.
//
// By default, abbreviated information about the clusters is returned. You can use
// the optional ShowCacheNodeInfo flag to retrieve detailed information about the
// cache nodes associated with the clusters. These details include the DNS address
// and port for the cache node endpoint.
//
// If the cluster is in the creating state, only cluster-level information is
// displayed until all of the nodes are successfully provisioned.
//
// If the cluster is in the deleting state, only cluster-level information is
// displayed.
//
// If cache nodes are currently being added to the cluster, node endpoint
// information and creation time for the additional nodes are not displayed until
// they are completely provisioned. When the cluster state is available, the
// cluster is ready for use.
//
// If cache nodes are currently being removed from the cluster, no endpoint
// information for the removed nodes is displayed.
func elasticache_DescribeCacheClusters(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeCacheClustersInput{}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheShowCacheClustersNotInReplicationGroups) > 0 {
		if err := assignInputField(input, "ShowCacheClustersNotInReplicationGroups", _elasticacheShowCacheClustersNotInReplicationGroups); err != nil {
			log.Errorf("invalid --show-cache-clusters-not-in-replication-groups: %s", err.Error())
			return
		}
	}
	if len(_elasticacheShowCacheNodeInfo) > 0 {
		if err := assignInputField(input, "ShowCacheNodeInfo", _elasticacheShowCacheNodeInfo); err != nil {
			log.Errorf("invalid --show-cache-node-info: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCacheClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeCacheClustersOutput
	p := elasticache.NewDescribeCacheClustersPaginator(client, input)
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

// Returns a list of the available cache engines and their versions.
func elasticache_DescribeCacheEngineVersions(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeCacheEngineVersionsInput{}

	if len(_elasticacheCacheParameterGroupFamily) > 0 {
		input.CacheParameterGroupFamily = aws.String(_elasticacheCacheParameterGroupFamily)
	}
	if len(_elasticacheDefaultOnly) > 0 {
		if err := assignInputField(input, "DefaultOnly", _elasticacheDefaultOnly); err != nil {
			log.Errorf("invalid --default-only: %s", err.Error())
			return
		}
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheEngineVersion) > 0 {
		input.EngineVersion = aws.String(_elasticacheEngineVersion)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCacheEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeCacheEngineVersionsOutput
	p := elasticache.NewDescribeCacheEngineVersionsPaginator(client, input)
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

// Returns a list of cache parameter group descriptions. If a cache parameter
// group name is specified, the list contains only the descriptions for that group.
func elasticache_DescribeCacheParameterGroups(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeCacheParameterGroupsInput{}

	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCacheParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeCacheParameterGroupsOutput
	p := elasticache.NewDescribeCacheParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular cache parameter group.
func elasticache_DescribeCacheParameters(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeCacheParametersInput{
		// CacheParameterGroupName: *string, // Required
	}

	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSource) > 0 {
		input.Source = aws.String(_elasticacheSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCacheParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeCacheParametersOutput
	p := elasticache.NewDescribeCacheParametersPaginator(client, input)
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

// Returns a list of cache security group descriptions. If a cache security group
// name is specified, the list contains only the description of that group. This
// applicable only when you have ElastiCache in Classic setup
func elasticache_DescribeCacheSecurityGroups(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeCacheSecurityGroupsInput{}

	if len(_elasticacheCacheSecurityGroupName) > 0 {
		input.CacheSecurityGroupName = aws.String(_elasticacheCacheSecurityGroupName)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCacheSecurityGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeCacheSecurityGroupsOutput
	p := elasticache.NewDescribeCacheSecurityGroupsPaginator(client, input)
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

// Returns a list of cache subnet group descriptions. If a subnet group name is
// specified, the list contains only the description of that group. This is
// applicable only when you have ElastiCache in VPC setup. All ElastiCache clusters
// now launch in VPC by default.
func elasticache_DescribeCacheSubnetGroups(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeCacheSubnetGroupsInput{}

	if len(_elasticacheCacheSubnetGroupName) > 0 {
		input.CacheSubnetGroupName = aws.String(_elasticacheCacheSubnetGroupName)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCacheSubnetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeCacheSubnetGroupsOutput
	p := elasticache.NewDescribeCacheSubnetGroupsPaginator(client, input)
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

// Returns the default engine and system parameter information for the specified
// cache engine.
func elasticache_DescribeEngineDefaultParameters(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeEngineDefaultParametersInput{
		// CacheParameterGroupFamily: *string, // Required
	}

	if len(_elasticacheCacheParameterGroupFamily) > 0 {
		input.CacheParameterGroupFamily = aws.String(_elasticacheCacheParameterGroupFamily)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEngineDefaultParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeEngineDefaultParametersOutput
	p := elasticache.NewDescribeEngineDefaultParametersPaginator(client, input)
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

// Returns events related to clusters, cache security groups, and cache parameter
// groups. You can obtain events specific to a particular cluster, cache security
// group, or cache parameter group by providing the name as a parameter.
//
// By default, only the events occurring within the last hour are returned;
// however, you can retrieve up to 14 days' worth of events if necessary.
func elasticache_DescribeEvents(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeEventsInput{}

	if len(_elasticacheDuration) > 0 {
		if err := assignInputField(input, "Duration", _elasticacheDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_elasticacheEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _elasticacheEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_elasticacheSourceIdentifier)
	}
	if len(_elasticacheSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _elasticacheSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_elasticacheStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _elasticacheStartTime); err != nil {
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

	var results []*elasticache.DescribeEventsOutput
	p := elasticache.NewDescribeEventsPaginator(client, input)
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

// Returns information about a particular global replication group. If no
// identifier is specified, returns information about all Global datastores.
func elasticache_DescribeGlobalReplicationGroups(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeGlobalReplicationGroupsInput{}

	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheShowMemberInfo) > 0 {
		if err := assignInputField(input, "ShowMemberInfo", _elasticacheShowMemberInfo); err != nil {
			log.Errorf("invalid --show-member-info: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeGlobalReplicationGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeGlobalReplicationGroupsOutput
	p := elasticache.NewDescribeGlobalReplicationGroupsPaginator(client, input)
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

// Returns information about a particular replication group. If no identifier is
// specified, DescribeReplicationGroups returns information about all replication
// groups.
//
// This operation is valid for Valkey or Redis OSS only.
func elasticache_DescribeReplicationGroups(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeReplicationGroupsInput{}

	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeReplicationGroupsOutput
	p := elasticache.NewDescribeReplicationGroupsPaginator(client, input)
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

// Returns information about reserved cache nodes for this account, or about a
// specified reserved cache node.
func elasticache_DescribeReservedCacheNodes(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeReservedCacheNodesInput{}

	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheDuration) > 0 {
		input.Duration = aws.String(_elasticacheDuration)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheOfferingType) > 0 {
		input.OfferingType = aws.String(_elasticacheOfferingType)
	}
	if len(_elasticacheProductDescription) > 0 {
		input.ProductDescription = aws.String(_elasticacheProductDescription)
	}
	if len(_elasticacheReservedCacheNodeId) > 0 {
		input.ReservedCacheNodeId = aws.String(_elasticacheReservedCacheNodeId)
	}
	if len(_elasticacheReservedCacheNodesOfferingId) > 0 {
		input.ReservedCacheNodesOfferingId = aws.String(_elasticacheReservedCacheNodesOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedCacheNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeReservedCacheNodesOutput
	p := elasticache.NewDescribeReservedCacheNodesPaginator(client, input)
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

// Lists available reserved cache node offerings.
func elasticache_DescribeReservedCacheNodesOfferings(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeReservedCacheNodesOfferingsInput{}

	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheDuration) > 0 {
		input.Duration = aws.String(_elasticacheDuration)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheOfferingType) > 0 {
		input.OfferingType = aws.String(_elasticacheOfferingType)
	}
	if len(_elasticacheProductDescription) > 0 {
		input.ProductDescription = aws.String(_elasticacheProductDescription)
	}
	if len(_elasticacheReservedCacheNodesOfferingId) > 0 {
		input.ReservedCacheNodesOfferingId = aws.String(_elasticacheReservedCacheNodesOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedCacheNodesOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeReservedCacheNodesOfferingsOutput
	p := elasticache.NewDescribeReservedCacheNodesOfferingsPaginator(client, input)
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

// Returns information about serverless cache snapshots. By default, this API
// lists all of the customer’s serverless cache snapshots. It can also describe a
// single serverless cache snapshot, or the snapshots associated with a particular
// serverless cache. Available for Valkey, Redis OSS and Serverless Memcached only.
func elasticache_DescribeServerlessCacheSnapshots(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeServerlessCacheSnapshotsInput{}

	if len(_elasticacheMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticacheMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNextToken) > 0 {
		input.NextToken = aws.String(_elasticacheNextToken)
	}
	if len(_elasticacheServerlessCacheName) > 0 {
		input.ServerlessCacheName = aws.String(_elasticacheServerlessCacheName)
	}
	if len(_elasticacheServerlessCacheSnapshotName) > 0 {
		input.ServerlessCacheSnapshotName = aws.String(_elasticacheServerlessCacheSnapshotName)
	}
	if len(_elasticacheSnapshotType) > 0 {
		input.SnapshotType = aws.String(_elasticacheSnapshotType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeServerlessCacheSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeServerlessCacheSnapshotsOutput
	p := elasticache.NewDescribeServerlessCacheSnapshotsPaginator(client, input)
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

// Returns information about a specific serverless cache. If no identifier is
// specified, then the API returns information on all the serverless caches
// belonging to this Amazon Web Services account.
func elasticache_DescribeServerlessCaches(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeServerlessCachesInput{}

	if len(_elasticacheMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticacheMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNextToken) > 0 {
		input.NextToken = aws.String(_elasticacheNextToken)
	}
	if len(_elasticacheServerlessCacheName) > 0 {
		input.ServerlessCacheName = aws.String(_elasticacheServerlessCacheName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeServerlessCaches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeServerlessCachesOutput
	p := elasticache.NewDescribeServerlessCachesPaginator(client, input)
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

// Returns details of the service updates
func elasticache_DescribeServiceUpdates(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeServiceUpdatesInput{}

	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheServiceUpdateName) > 0 {
		input.ServiceUpdateName = aws.String(_elasticacheServiceUpdateName)
	}
	if len(_elasticacheServiceUpdateStatus) > 0 {
		if err := assignInputField(input, "ServiceUpdateStatus", _elasticacheServiceUpdateStatus); err != nil {
			log.Errorf("invalid --service-update-status: %s", err.Error())
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

	var results []*elasticache.DescribeServiceUpdatesOutput
	p := elasticache.NewDescribeServiceUpdatesPaginator(client, input)
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

// Returns information about cluster or replication group snapshots. By default,
// DescribeSnapshots lists all of your snapshots; it can optionally describe a
// single snapshot, or just the snapshots associated with a particular cache
// cluster.
//
// This operation is valid for Valkey or Redis OSS only.
func elasticache_DescribeSnapshots(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeSnapshotsInput{}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheShowNodeGroupConfig) > 0 {
		if err := assignInputField(input, "ShowNodeGroupConfig", _elasticacheShowNodeGroupConfig); err != nil {
			log.Errorf("invalid --show-node-group-config: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSnapshotName) > 0 {
		input.SnapshotName = aws.String(_elasticacheSnapshotName)
	}
	if len(_elasticacheSnapshotSource) > 0 {
		input.SnapshotSource = aws.String(_elasticacheSnapshotSource)
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

	var results []*elasticache.DescribeSnapshotsOutput
	p := elasticache.NewDescribeSnapshotsPaginator(client, input)
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

// Returns details of the update actions
func elasticache_DescribeUpdateActions(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeUpdateActionsInput{}

	if len(_elasticacheCacheClusterIds) > 0 {
		input.CacheClusterIds = append([]string(nil), _elasticacheCacheClusterIds...)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupIds) > 0 {
		input.ReplicationGroupIds = append([]string(nil), _elasticacheReplicationGroupIds...)
	}
	if len(_elasticacheServiceUpdateName) > 0 {
		input.ServiceUpdateName = aws.String(_elasticacheServiceUpdateName)
	}
	if len(_elasticacheServiceUpdateStatus) > 0 {
		if err := assignInputField(input, "ServiceUpdateStatus", _elasticacheServiceUpdateStatus); err != nil {
			log.Errorf("invalid --service-update-status: %s", err.Error())
			return
		}
	}
	if len(_elasticacheServiceUpdateTimeRange) > 0 {
		if err := assignInputField(input, "ServiceUpdateTimeRange", _elasticacheServiceUpdateTimeRange); err != nil {
			log.Errorf("invalid --service-update-time-range: %s", err.Error())
			return
		}
	}
	if len(_elasticacheShowNodeLevelUpdateStatus) > 0 {
		if err := assignInputField(input, "ShowNodeLevelUpdateStatus", _elasticacheShowNodeLevelUpdateStatus); err != nil {
			log.Errorf("invalid --show-node-level-update-status: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUpdateActionStatus) > 0 {
		if err := assignInputField(input, "UpdateActionStatus", _elasticacheUpdateActionStatus); err != nil {
			log.Errorf("invalid --update-action-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeUpdateActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeUpdateActionsOutput
	p := elasticache.NewDescribeUpdateActionsPaginator(client, input)
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

// Returns a list of user groups.
func elasticache_DescribeUserGroups(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeUserGroupsInput{}

	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserGroupId) > 0 {
		input.UserGroupId = aws.String(_elasticacheUserGroupId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeUserGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticache.DescribeUserGroupsOutput
	p := elasticache.NewDescribeUserGroupsPaginator(client, input)
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
func elasticache_DescribeUsers(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DescribeUsersInput{}

	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheFilters) > 0 {
		if err := assignInputField(input, "Filters", _elasticacheFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_elasticacheMarker) > 0 {
		input.Marker = aws.String(_elasticacheMarker)
	}
	if len(_elasticacheMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _elasticacheMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserId) > 0 {
		input.UserId = aws.String(_elasticacheUserId)
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

	var results []*elasticache.DescribeUsersOutput
	p := elasticache.NewDescribeUsersPaginator(client, input)
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

// Remove a secondary cluster from the Global datastore using the Global datastore
// name. The secondary cluster will no longer receive updates from the primary
// cluster, but will remain as a standalone cluster in that Amazon region.
func elasticache_DisassociateGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.DisassociateGlobalReplicationGroupInput{
		// GlobalReplicationGroupId: *string, // Required
		// ReplicationGroupId: *string, // Required
		// ReplicationGroupRegion: *string, // Required
	}

	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheReplicationGroupRegion) > 0 {
		input.ReplicationGroupRegion = aws.String(_elasticacheReplicationGroupRegion)
	}

	if resp, err := client.DisassociateGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the functionality to export the serverless cache snapshot data to
// Amazon S3. Available for Valkey and Redis OSS only.
func elasticache_ExportServerlessCacheSnapshot(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ExportServerlessCacheSnapshotInput{
		// S3BucketName: *string, // Required
		// ServerlessCacheSnapshotName: *string, // Required
	}

	if len(_elasticacheS3BucketName) > 0 {
		input.S3BucketName = aws.String(_elasticacheS3BucketName)
	}
	if len(_elasticacheServerlessCacheSnapshotName) > 0 {
		input.ServerlessCacheSnapshotName = aws.String(_elasticacheServerlessCacheSnapshotName)
	}

	if resp, err := client.ExportServerlessCacheSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to failover the primary region to a secondary region. The secondary region
// will become primary, and all other clusters will become secondary.
func elasticache_FailoverGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.FailoverGlobalReplicationGroupInput{
		// GlobalReplicationGroupId: *string, // Required
		// PrimaryRegion: *string, // Required
		// PrimaryReplicationGroupId: *string, // Required
	}

	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticachePrimaryRegion) > 0 {
		input.PrimaryRegion = aws.String(_elasticachePrimaryRegion)
	}
	if len(_elasticachePrimaryReplicationGroupId) > 0 {
		input.PrimaryReplicationGroupId = aws.String(_elasticachePrimaryReplicationGroupId)
	}

	if resp, err := client.FailoverGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Increase the number of node groups in the Global datastore
func elasticache_IncreaseNodeGroupsInGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput{
		// ApplyImmediately: *bool, // Required
		// GlobalReplicationGroupId: *string, // Required
		// NodeGroupCount: *int32, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheNodeGroupCount) > 0 {
		if err := assignInputField(input, "NodeGroupCount", _elasticacheNodeGroupCount); err != nil {
			log.Errorf("invalid --node-group-count: %s", err.Error())
			return
		}
	}
	if len(_elasticacheRegionalConfigurations) > 0 {
		if err := assignInputField(input, "RegionalConfigurations", _elasticacheRegionalConfigurations); err != nil {
			log.Errorf("invalid --regional-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.IncreaseNodeGroupsInGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dynamically increases the number of replicas in a Valkey or Redis OSS (cluster
// mode disabled) replication group or the number of replica nodes in one or more
// node groups (shards) of a Valkey or Redis OSS (cluster mode enabled) replication
// group. This operation is performed with no cluster down time.
func elasticache_IncreaseReplicaCount(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.IncreaseReplicaCountInput{
		// ApplyImmediately: *bool, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheNewReplicaCount) > 0 {
		if err := assignInputField(input, "NewReplicaCount", _elasticacheNewReplicaCount); err != nil {
			log.Errorf("invalid --new-replica-count: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicaConfiguration) > 0 {
		if err := assignInputField(input, "ReplicaConfiguration", _elasticacheReplicaConfiguration); err != nil {
			log.Errorf("invalid --replica-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.IncreaseReplicaCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all available node types that you can scale with your cluster's
// replication group's current node type.
//
// When you use the ModifyCacheCluster or ModifyReplicationGroup operations to
// scale your cluster or replication group, the value of the CacheNodeType
// parameter must be one of the node types returned by this operation.
func elasticache_ListAllowedNodeTypeModifications(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ListAllowedNodeTypeModificationsInput{}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}

	if resp, err := client.ListAllowedNodeTypeModifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all tags currently on a named resource.
// A tag is a key-value pair where the key and value are case-sensitive. You can
// use tags to categorize and track all your ElastiCache resources, with the
// exception of global replication group. When you add or remove tags on
// replication groups, those actions will be replicated to all nodes in the
// replication group. For more information, see [Resource-level permissions].
//
// If the cluster is not in the available state, ListTagsForResource returns an
// error.
//
// [Resource-level permissions]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/IAM.ResourceLevelPermissions.html
func elasticache_ListTagsForResource(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ListTagsForResourceInput{
		// ResourceName: *string, // Required
	}

	if len(_elasticacheResourceName) > 0 {
		input.ResourceName = aws.String(_elasticacheResourceName)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a cluster. You can use this operation to change one
// or more cluster configuration parameters by specifying the parameters and the
// new values.
func elasticache_ModifyCacheCluster(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyCacheClusterInput{
		// CacheClusterId: *string, // Required
	}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheAZMode) > 0 {
		if err := assignInputField(input, "AZMode", _elasticacheAZMode); err != nil {
			log.Errorf("invalid --az-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAuthToken) > 0 {
		input.AuthToken = aws.String(_elasticacheAuthToken)
	}
	if len(_elasticacheAuthTokenUpdateStrategy) > 0 {
		if err := assignInputField(input, "AuthTokenUpdateStrategy", _elasticacheAuthTokenUpdateStrategy); err != nil {
			log.Errorf("invalid --auth-token-update-strategy: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _elasticacheAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_elasticacheCacheNodeIdsToRemove) > 0 {
		input.CacheNodeIdsToRemove = append([]string(nil), _elasticacheCacheNodeIdsToRemove...)
	}
	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheCacheSecurityGroupNames) > 0 {
		input.CacheSecurityGroupNames = append([]string(nil), _elasticacheCacheSecurityGroupNames...)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheEngineVersion) > 0 {
		input.EngineVersion = aws.String(_elasticacheEngineVersion)
	}
	if len(_elasticacheIpDiscovery) > 0 {
		if err := assignInputField(input, "IpDiscovery", _elasticacheIpDiscovery); err != nil {
			log.Errorf("invalid --ip-discovery: %s", err.Error())
			return
		}
	}
	if len(_elasticacheLogDeliveryConfigurations) > 0 {
		if err := assignInputField(input, "LogDeliveryConfigurations", _elasticacheLogDeliveryConfigurations); err != nil {
			log.Errorf("invalid --log-delivery-configurations: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNewAvailabilityZones) > 0 {
		input.NewAvailabilityZones = append([]string(nil), _elasticacheNewAvailabilityZones...)
	}
	if len(_elasticacheNotificationTopicArn) > 0 {
		input.NotificationTopicArn = aws.String(_elasticacheNotificationTopicArn)
	}
	if len(_elasticacheNotificationTopicStatus) > 0 {
		input.NotificationTopicStatus = aws.String(_elasticacheNotificationTopicStatus)
	}
	if len(_elasticacheNumCacheNodes) > 0 {
		if err := assignInputField(input, "NumCacheNodes", _elasticacheNumCacheNodes); err != nil {
			log.Errorf("invalid --num-cache-nodes: %s", err.Error())
			return
		}
	}
	if len(_elasticachePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_elasticachePreferredMaintenanceWindow)
	}
	if len(_elasticacheScaleConfig) > 0 {
		if err := assignInputField(input, "ScaleConfig", _elasticacheScaleConfig); err != nil {
			log.Errorf("invalid --scale-config: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _elasticacheSecurityGroupIds...)
	}
	if len(_elasticacheSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _elasticacheSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSnapshotWindow) > 0 {
		input.SnapshotWindow = aws.String(_elasticacheSnapshotWindow)
	}

	if resp, err := client.ModifyCacheCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a cache parameter group. You can modify up to 20
// parameters in a single request by submitting a list parameter name and value
// pairs.
func elasticache_ModifyCacheParameterGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyCacheParameterGroupInput{
		// CacheParameterGroupName: *string, // Required
		// ParameterNameValues: []types.ParameterNameValue, // Required
	}

	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheParameterNameValues) > 0 {
		if err := assignInputField(input, "ParameterNameValues", _elasticacheParameterNameValues); err != nil {
			log.Errorf("invalid --parameter-name-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyCacheParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing cache subnet group.
func elasticache_ModifyCacheSubnetGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyCacheSubnetGroupInput{
		// CacheSubnetGroupName: *string, // Required
	}

	if len(_elasticacheCacheSubnetGroupName) > 0 {
		input.CacheSubnetGroupName = aws.String(_elasticacheCacheSubnetGroupName)
	}
	if len(_elasticacheCacheSubnetGroupDescription) > 0 {
		input.CacheSubnetGroupDescription = aws.String(_elasticacheCacheSubnetGroupDescription)
	}
	if len(_elasticacheSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _elasticacheSubnetIds...)
	}

	if resp, err := client.ModifyCacheSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a Global datastore.
func elasticache_ModifyGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyGlobalReplicationGroupInput{
		// ApplyImmediately: *bool, // Required
		// GlobalReplicationGroupId: *string, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}
	if len(_elasticacheAutomaticFailoverEnabled) > 0 {
		if err := assignInputField(input, "AutomaticFailoverEnabled", _elasticacheAutomaticFailoverEnabled); err != nil {
			log.Errorf("invalid --automatic-failover-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheEngineVersion) > 0 {
		input.EngineVersion = aws.String(_elasticacheEngineVersion)
	}
	if len(_elasticacheGlobalReplicationGroupDescription) > 0 {
		input.GlobalReplicationGroupDescription = aws.String(_elasticacheGlobalReplicationGroupDescription)
	}

	if resp, err := client.ModifyGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a replication group. This is limited to Valkey and
// Redis OSS 7 and above.
//
// [Scaling for Valkey or Redis OSS (cluster mode enabled)]
// - in the ElastiCache User Guide
//
// [ModifyReplicationGroupShardConfiguration]
// - in the ElastiCache API Reference
//
// This operation is valid for Valkey or Redis OSS only.
//
// [ModifyReplicationGroupShardConfiguration]: https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyReplicationGroupShardConfiguration.html
// [Scaling for Valkey or Redis OSS (cluster mode enabled)]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/scaling-redis-cluster-mode-enabled.html
func elasticache_ModifyReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyReplicationGroupInput{
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAuthToken) > 0 {
		input.AuthToken = aws.String(_elasticacheAuthToken)
	}
	if len(_elasticacheAuthTokenUpdateStrategy) > 0 {
		if err := assignInputField(input, "AuthTokenUpdateStrategy", _elasticacheAuthTokenUpdateStrategy); err != nil {
			log.Errorf("invalid --auth-token-update-strategy: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _elasticacheAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_elasticacheAutomaticFailoverEnabled) > 0 {
		if err := assignInputField(input, "AutomaticFailoverEnabled", _elasticacheAutomaticFailoverEnabled); err != nil {
			log.Errorf("invalid --automatic-failover-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheCacheNodeType) > 0 {
		input.CacheNodeType = aws.String(_elasticacheCacheNodeType)
	}
	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheCacheSecurityGroupNames) > 0 {
		input.CacheSecurityGroupNames = append([]string(nil), _elasticacheCacheSecurityGroupNames...)
	}
	if len(_elasticacheClusterMode) > 0 {
		if err := assignInputField(input, "ClusterMode", _elasticacheClusterMode); err != nil {
			log.Errorf("invalid --cluster-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheEngineVersion) > 0 {
		input.EngineVersion = aws.String(_elasticacheEngineVersion)
	}
	if len(_elasticacheIpDiscovery) > 0 {
		if err := assignInputField(input, "IpDiscovery", _elasticacheIpDiscovery); err != nil {
			log.Errorf("invalid --ip-discovery: %s", err.Error())
			return
		}
	}
	if len(_elasticacheLogDeliveryConfigurations) > 0 {
		if err := assignInputField(input, "LogDeliveryConfigurations", _elasticacheLogDeliveryConfigurations); err != nil {
			log.Errorf("invalid --log-delivery-configurations: %s", err.Error())
			return
		}
	}
	if len(_elasticacheMultiAZEnabled) > 0 {
		if err := assignInputField(input, "MultiAZEnabled", _elasticacheMultiAZEnabled); err != nil {
			log.Errorf("invalid --multi-az-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNodeGroupId) > 0 {
		input.NodeGroupId = aws.String(_elasticacheNodeGroupId)
	}
	if len(_elasticacheNotificationTopicArn) > 0 {
		input.NotificationTopicArn = aws.String(_elasticacheNotificationTopicArn)
	}
	if len(_elasticacheNotificationTopicStatus) > 0 {
		input.NotificationTopicStatus = aws.String(_elasticacheNotificationTopicStatus)
	}
	if len(_elasticachePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_elasticachePreferredMaintenanceWindow)
	}
	if len(_elasticachePrimaryClusterId) > 0 {
		input.PrimaryClusterId = aws.String(_elasticachePrimaryClusterId)
	}
	if len(_elasticacheRemoveUserGroups) > 0 {
		if err := assignInputField(input, "RemoveUserGroups", _elasticacheRemoveUserGroups); err != nil {
			log.Errorf("invalid --remove-user-groups: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupDescription) > 0 {
		input.ReplicationGroupDescription = aws.String(_elasticacheReplicationGroupDescription)
	}
	if len(_elasticacheSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _elasticacheSecurityGroupIds...)
	}
	if len(_elasticacheSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _elasticacheSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSnapshotWindow) > 0 {
		input.SnapshotWindow = aws.String(_elasticacheSnapshotWindow)
	}
	if len(_elasticacheSnapshottingClusterId) > 0 {
		input.SnapshottingClusterId = aws.String(_elasticacheSnapshottingClusterId)
	}
	if len(_elasticacheTransitEncryptionEnabled) > 0 {
		if err := assignInputField(input, "TransitEncryptionEnabled", _elasticacheTransitEncryptionEnabled); err != nil {
			log.Errorf("invalid --transit-encryption-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticacheTransitEncryptionMode) > 0 {
		if err := assignInputField(input, "TransitEncryptionMode", _elasticacheTransitEncryptionMode); err != nil {
			log.Errorf("invalid --transit-encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserGroupIdsToAdd) > 0 {
		input.UserGroupIdsToAdd = append([]string(nil), _elasticacheUserGroupIdsToAdd...)
	}
	if len(_elasticacheUserGroupIdsToRemove) > 0 {
		input.UserGroupIdsToRemove = append([]string(nil), _elasticacheUserGroupIdsToRemove...)
	}

	if resp, err := client.ModifyReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a replication group's shards (node groups) by allowing you to add
// shards, remove shards, or rebalance the keyspaces among existing shards.
func elasticache_ModifyReplicationGroupShardConfiguration(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyReplicationGroupShardConfigurationInput{
		// ApplyImmediately: *bool, // Required
		// NodeGroupCount: *int32, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheNodeGroupCount) > 0 {
		if err := assignInputField(input, "NodeGroupCount", _elasticacheNodeGroupCount); err != nil {
			log.Errorf("invalid --node-group-count: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}
	if len(_elasticacheNodeGroupsToRemove) > 0 {
		input.NodeGroupsToRemove = append([]string(nil), _elasticacheNodeGroupsToRemove...)
	}
	if len(_elasticacheNodeGroupsToRetain) > 0 {
		input.NodeGroupsToRetain = append([]string(nil), _elasticacheNodeGroupsToRetain...)
	}
	if len(_elasticacheReshardingConfiguration) > 0 {
		if err := assignInputField(input, "ReshardingConfiguration", _elasticacheReshardingConfiguration); err != nil {
			log.Errorf("invalid --resharding-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyReplicationGroupShardConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API modifies the attributes of a serverless cache.
func elasticache_ModifyServerlessCache(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyServerlessCacheInput{
		// ServerlessCacheName: *string, // Required
	}

	if len(_elasticacheServerlessCacheName) > 0 {
		input.ServerlessCacheName = aws.String(_elasticacheServerlessCacheName)
	}
	if len(_elasticacheCacheUsageLimits) > 0 {
		if err := assignInputField(input, "CacheUsageLimits", _elasticacheCacheUsageLimits); err != nil {
			log.Errorf("invalid --cache-usage-limits: %s", err.Error())
			return
		}
	}
	if len(_elasticacheDailySnapshotTime) > 0 {
		input.DailySnapshotTime = aws.String(_elasticacheDailySnapshotTime)
	}
	if len(_elasticacheDescription) > 0 {
		input.Description = aws.String(_elasticacheDescription)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheMajorEngineVersion) > 0 {
		input.MajorEngineVersion = aws.String(_elasticacheMajorEngineVersion)
	}
	if len(_elasticacheRemoveUserGroup) > 0 {
		if err := assignInputField(input, "RemoveUserGroup", _elasticacheRemoveUserGroup); err != nil {
			log.Errorf("invalid --remove-user-group: %s", err.Error())
			return
		}
	}
	if len(_elasticacheSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _elasticacheSecurityGroupIds...)
	}
	if len(_elasticacheSnapshotRetentionLimit) > 0 {
		if err := assignInputField(input, "SnapshotRetentionLimit", _elasticacheSnapshotRetentionLimit); err != nil {
			log.Errorf("invalid --snapshot-retention-limit: %s", err.Error())
			return
		}
	}
	if len(_elasticacheUserGroupId) > 0 {
		input.UserGroupId = aws.String(_elasticacheUserGroupId)
	}

	if resp, err := client.ModifyServerlessCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes user password(s) and/or access string.
func elasticache_ModifyUser(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyUserInput{
		// UserId: *string, // Required
	}

	if len(_elasticacheUserId) > 0 {
		input.UserId = aws.String(_elasticacheUserId)
	}
	if len(_elasticacheAccessString) > 0 {
		input.AccessString = aws.String(_elasticacheAccessString)
	}
	if len(_elasticacheAppendAccessString) > 0 {
		input.AppendAccessString = aws.String(_elasticacheAppendAccessString)
	}
	if len(_elasticacheAuthenticationMode) > 0 {
		if err := assignInputField(input, "AuthenticationMode", _elasticacheAuthenticationMode); err != nil {
			log.Errorf("invalid --authentication-mode: %s", err.Error())
			return
		}
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheNoPasswordRequired) > 0 {
		if err := assignInputField(input, "NoPasswordRequired", _elasticacheNoPasswordRequired); err != nil {
			log.Errorf("invalid --no-password-required: %s", err.Error())
			return
		}
	}
	if len(_elasticachePasswords) > 0 {
		input.Passwords = append([]string(nil), _elasticachePasswords...)
	}

	if resp, err := client.ModifyUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the list of users that belong to the user group.
func elasticache_ModifyUserGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ModifyUserGroupInput{
		// UserGroupId: *string, // Required
	}

	if len(_elasticacheUserGroupId) > 0 {
		input.UserGroupId = aws.String(_elasticacheUserGroupId)
	}
	if len(_elasticacheEngine) > 0 {
		input.Engine = aws.String(_elasticacheEngine)
	}
	if len(_elasticacheUserIdsToAdd) > 0 {
		input.UserIdsToAdd = append([]string(nil), _elasticacheUserIdsToAdd...)
	}
	if len(_elasticacheUserIdsToRemove) > 0 {
		input.UserIdsToRemove = append([]string(nil), _elasticacheUserIdsToRemove...)
	}

	if resp, err := client.ModifyUserGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to purchase a reserved cache node offering. Reserved nodes are not
// eligible for cancellation and are non-refundable. For more information, see [Managing Costs with Reserved Nodes].
//
// [Managing Costs with Reserved Nodes]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/reserved-nodes.html
func elasticache_PurchaseReservedCacheNodesOffering(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.PurchaseReservedCacheNodesOfferingInput{
		// ReservedCacheNodesOfferingId: *string, // Required
	}

	if len(_elasticacheReservedCacheNodesOfferingId) > 0 {
		input.ReservedCacheNodesOfferingId = aws.String(_elasticacheReservedCacheNodesOfferingId)
	}
	if len(_elasticacheCacheNodeCount) > 0 {
		if err := assignInputField(input, "CacheNodeCount", _elasticacheCacheNodeCount); err != nil {
			log.Errorf("invalid --cache-node-count: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReservedCacheNodeId) > 0 {
		input.ReservedCacheNodeId = aws.String(_elasticacheReservedCacheNodeId)
	}
	if len(_elasticacheTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticacheTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseReservedCacheNodesOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Redistribute slots to ensure uniform distribution across existing shards in the
// cluster.
func elasticache_RebalanceSlotsInGlobalReplicationGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.RebalanceSlotsInGlobalReplicationGroupInput{
		// ApplyImmediately: *bool, // Required
		// GlobalReplicationGroupId: *string, // Required
	}

	if len(_elasticacheApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _elasticacheApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_elasticacheGlobalReplicationGroupId) > 0 {
		input.GlobalReplicationGroupId = aws.String(_elasticacheGlobalReplicationGroupId)
	}

	if resp, err := client.RebalanceSlotsInGlobalReplicationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots some, or all, of the cache nodes within a provisioned cluster. This
// operation applies any modified cache parameter groups to the cluster. The reboot
// operation takes place as soon as possible, and results in a momentary outage to
// the cluster. During the reboot, the cluster status is set to REBOOTING.
//
// The reboot causes the contents of the cache (for each cache node being
// rebooted) to be lost.
//
// When the reboot is complete, a cluster event is created.
//
// Rebooting a cluster is currently supported on Memcached, Valkey and Redis OSS
// (cluster mode disabled) clusters. Rebooting is not supported on Valkey or Redis
// OSS (cluster mode enabled) clusters.
//
// If you make changes to parameters that require a Valkey or Redis OSS (cluster
// mode enabled) cluster reboot for the changes to be applied, see [Rebooting a Cluster]for an
// alternate process.
//
// [Rebooting a Cluster]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/nodes.rebooting.html
func elasticache_RebootCacheCluster(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.RebootCacheClusterInput{
		// CacheClusterId: *string, // Required
		// CacheNodeIdsToReboot: []string, // Required
	}

	if len(_elasticacheCacheClusterId) > 0 {
		input.CacheClusterId = aws.String(_elasticacheCacheClusterId)
	}
	if len(_elasticacheCacheNodeIdsToReboot) > 0 {
		input.CacheNodeIdsToReboot = append([]string(nil), _elasticacheCacheNodeIdsToReboot...)
	}

	if resp, err := client.RebootCacheCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the tags identified by the TagKeys list from the named resource. A tag
// is a key-value pair where the key and value are case-sensitive. You can use tags
// to categorize and track all your ElastiCache resources, with the exception of
// global replication group. When you add or remove tags on replication groups,
// those actions will be replicated to all nodes in the replication group. For more
// information, see [Resource-level permissions].
//
// [Resource-level permissions]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/IAM.ResourceLevelPermissions.html
func elasticache_RemoveTagsFromResource(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.RemoveTagsFromResourceInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_elasticacheResourceName) > 0 {
		input.ResourceName = aws.String(_elasticacheResourceName)
	}
	if len(_elasticacheTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _elasticacheTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a cache parameter group to the engine or system
// default value. You can reset specific parameters by submitting a list of
// parameter names. To reset the entire cache parameter group, specify the
// ResetAllParameters and CacheParameterGroupName parameters.
func elasticache_ResetCacheParameterGroup(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.ResetCacheParameterGroupInput{
		// CacheParameterGroupName: *string, // Required
	}

	if len(_elasticacheCacheParameterGroupName) > 0 {
		input.CacheParameterGroupName = aws.String(_elasticacheCacheParameterGroupName)
	}
	if len(_elasticacheParameterNameValues) > 0 {
		if err := assignInputField(input, "ParameterNameValues", _elasticacheParameterNameValues); err != nil {
			log.Errorf("invalid --parameter-name-values: %s", err.Error())
			return
		}
	}
	if len(_elasticacheResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _elasticacheResetAllParameters); err != nil {
			log.Errorf("invalid --reset-all-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetCacheParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes ingress from a cache security group. Use this operation to disallow
// access from an Amazon EC2 security group that had been previously authorized.
func elasticache_RevokeCacheSecurityGroupIngress(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.RevokeCacheSecurityGroupIngressInput{
		// CacheSecurityGroupName: *string, // Required
		// EC2SecurityGroupName: *string, // Required
		// EC2SecurityGroupOwnerId: *string, // Required
	}

	if len(_elasticacheCacheSecurityGroupName) > 0 {
		input.CacheSecurityGroupName = aws.String(_elasticacheCacheSecurityGroupName)
	}
	if len(_elasticacheEC2SecurityGroupName) > 0 {
		input.EC2SecurityGroupName = aws.String(_elasticacheEC2SecurityGroupName)
	}
	if len(_elasticacheEC2SecurityGroupOwnerId) > 0 {
		input.EC2SecurityGroupOwnerId = aws.String(_elasticacheEC2SecurityGroupOwnerId)
	}

	if resp, err := client.RevokeCacheSecurityGroupIngress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start the migration of data.
func elasticache_StartMigration(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.StartMigrationInput{
		// CustomerNodeEndpointList: []types.CustomerNodeEndpoint, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheCustomerNodeEndpointList) > 0 {
		if err := assignInputField(input, "CustomerNodeEndpointList", _elasticacheCustomerNodeEndpointList); err != nil {
			log.Errorf("invalid --customer-node-endpoint-list: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}

	if resp, err := client.StartMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents the input of a TestFailover operation which tests automatic failover
// on a specified node group (called shard in the console) in a replication group
// (called cluster in the console).
//
// This API is designed for testing the behavior of your application in case of
// ElastiCache failover. It is not designed to be an operational tool for
// initiating a failover to overcome a problem you may have with the cluster.
// Moreover, in certain conditions such as large-scale operational events, Amazon
// may block this API.
//
// # Note the following
//
// - A customer can use this operation to test automatic failover on up to 15
// shards (called node groups in the ElastiCache API and Amazon CLI) in any rolling
// 24-hour period.
//
// - If calling this operation on shards in different clusters (called
// replication groups in the API and CLI), the calls can be made concurrently.
//
// - If calling this operation multiple times on different shards in the same
// Valkey or Redis OSS (cluster mode enabled) replication group, the first node
// replacement must complete before a subsequent call can be made.
//
// - To determine whether the node replacement is complete you can check Events
// using the Amazon ElastiCache console, the Amazon CLI, or the ElastiCache API.
// Look for the following automatic failover related events, listed here in order
// of occurrance:
//
// - Replication group message: Test Failover API called for node group
//
// - Cache cluster message: Failover from primary node to replica node completed
//
// - Replication group message: Failover from primary node to replica node
// completed
//
// - Cache cluster message: Recovering cache nodes
//
// - Cache cluster message: Finished recovery for cache nodes
//
// For more information see:
//
// [Viewing ElastiCache Events]
// - in the ElastiCache User Guide
//
// [DescribeEvents]
// - in the ElastiCache API Reference
//
// Also see, [Testing Multi-AZ] in the ElastiCache User Guide.
//
// [DescribeEvents]: https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_DescribeEvents.html
// [Testing Multi-AZ]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/AutoFailover.html#auto-failover-test
// [Viewing ElastiCache Events]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ECEvents.Viewing.html
func elasticache_TestFailover(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.TestFailoverInput{
		// NodeGroupId: *string, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheNodeGroupId) > 0 {
		input.NodeGroupId = aws.String(_elasticacheNodeGroupId)
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}

	if resp, err := client.TestFailover(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Async API to test connection between source and target replication group.
func elasticache_TestMigration(cfg aws.Config, client *elasticache.Client) {
	input := &elasticache.TestMigrationInput{
		// CustomerNodeEndpointList: []types.CustomerNodeEndpoint, // Required
		// ReplicationGroupId: *string, // Required
	}

	if len(_elasticacheCustomerNodeEndpointList) > 0 {
		if err := assignInputField(input, "CustomerNodeEndpointList", _elasticacheCustomerNodeEndpointList); err != nil {
			log.Errorf("invalid --customer-node-endpoint-list: %s", err.Error())
			return
		}
	}
	if len(_elasticacheReplicationGroupId) > 0 {
		input.ReplicationGroupId = aws.String(_elasticacheReplicationGroupId)
	}

	if resp, err := client.TestMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_elasticacheCmd)
	_elasticacheCmd.Flags().SortFlags = false

	_elasticacheCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_elasticacheCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_elasticacheCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_elasticacheCmd.Flags().StringVarP(&_elasticacheAccessString, "access-string", "", "", "Access String")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAppendAccessString, "append-access-string", "", "", "Append Access String")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheApplyImmediately, "apply-immediately", "", "", "Apply Immediately")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAtRestEncryptionEnabled, "at-rest-encryption-enabled", "", "", "At Rest Encryption Enabled")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAuthToken, "auth-token", "", "", "Auth Token")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAuthTokenUpdateStrategy, "auth-token-update-strategy", "", "", "Auth Token Update Strategy")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAuthenticationMode, "authentication-mode", "", "", "Authentication Mode")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAutomaticFailoverEnabled, "automatic-failover-enabled", "", "", "Automatic Failover Enabled")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheAZMode, "az-mode", "", "", "AZ Mode")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheClusterId, "cache-cluster-id", "", "", "Cache Cluster ID")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheCacheClusterIds, "cache-cluster-ids", "", nil, "Cache Cluster Ids")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheNodeCount, "cache-node-count", "", "", "Cache Node Count")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheCacheNodeIdsToReboot, "cache-node-ids-to-reboot", "", nil, "Cache Node Ids To Reboot")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheCacheNodeIdsToRemove, "cache-node-ids-to-remove", "", nil, "Cache Node Ids To Remove")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheNodeType, "cache-node-type", "", "", "Cache Node Type")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheParameterGroupFamily, "cache-parameter-group-family", "", "", "Cache Parameter Group Family")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheParameterGroupName, "cache-parameter-group-name", "", "", "Cache Parameter Group Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheSecurityGroupName, "cache-security-group-name", "", "", "Cache Security Group Name")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheCacheSecurityGroupNames, "cache-security-group-names", "", nil, "Cache Security Group Names")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheSubnetGroupDescription, "cache-subnet-group-description", "", "", "Cache Subnet Group Description")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheSubnetGroupName, "cache-subnet-group-name", "", "", "Cache Subnet Group Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCacheUsageLimits, "cache-usage-limits", "", "", "Cache Usage Limits")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheClusterMode, "cluster-mode", "", "", "Cluster Mode")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheCustomerNodeEndpointList, "customer-node-endpoint-list", "", "", "Customer Node Endpoint List")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheDailySnapshotTime, "daily-snapshot-time", "", "", "Daily Snapshot Time")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheDataTieringEnabled, "data-tiering-enabled", "", "", "Data Tiering Enabled")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheDefaultOnly, "default-only", "", "", "Default Only")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheDescription, "description", "", "", "Description")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheDuration, "duration", "", "", "Duration")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheEC2SecurityGroupName, "ec2-security-group-name", "", "", "EC2 Security Group Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheEC2SecurityGroupOwnerId, "ec2-security-group-owner-id", "", "", "EC2 Security Group Owner ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheEndTime, "end-time", "", "", "End Time")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheEngine, "engine", "", "", "Engine")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheEngineVersion, "engine-version", "", "", "Engine Version")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheFilters, "filters", "", "", "Filters")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheFinalSnapshotIdentifier, "final-snapshot-identifier", "", "", "Final Snapshot Identifier")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheFinalSnapshotName, "final-snapshot-name", "", "", "Final Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheForce, "force", "", "", "Force")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheGlobalNodeGroupsToRemove, "global-node-groups-to-remove", "", nil, "Global Node Groups To Remove")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheGlobalNodeGroupsToRetain, "global-node-groups-to-retain", "", nil, "Global Node Groups To Retain")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheGlobalReplicationGroupDescription, "global-replication-group-description", "", "", "Global Replication Group Description")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheGlobalReplicationGroupId, "global-replication-group-id", "", "", "Global Replication Group ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheGlobalReplicationGroupIdSuffix, "global-replication-group-id-suffix", "", "", "Global Replication Group ID Suffix")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheIpDiscovery, "ip-discovery", "", "", "IP Discovery")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheLogDeliveryConfigurations, "log-delivery-configurations", "", "", "Log Delivery Configurations")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheMajorEngineVersion, "major-engine-version", "", "", "Major Engine Version")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheMarker, "marker", "", "", "Marker")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheMaxRecords, "max-records", "", "", "Max Records")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheMaxResults, "max-results", "", "", "Max Results")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheMultiAZEnabled, "multi-az-enabled", "", "", "Multi AZ Enabled")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNetworkType, "network-type", "", "", "Network Type")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheNewAvailabilityZones, "new-availability-zones", "", nil, "New Availability Zones")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNewReplicaCount, "new-replica-count", "", "", "New Replica Count")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNextToken, "next-token", "", "", "Next Token")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNoPasswordRequired, "no-password-required", "", "", "No Password Required")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNodeGroupConfiguration, "node-group-configuration", "", "", "Node Group Configuration")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNodeGroupCount, "node-group-count", "", "", "Node Group Count")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNodeGroupId, "node-group-id", "", "", "Node Group ID")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheNodeGroupsToRemove, "node-groups-to-remove", "", nil, "Node Groups To Remove")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheNodeGroupsToRetain, "node-groups-to-retain", "", nil, "Node Groups To Retain")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNotificationTopicArn, "notification-topic-arn", "", "", "Notification Topic ARN")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNotificationTopicStatus, "notification-topic-status", "", "", "Notification Topic Status")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNumCacheClusters, "num-cache-clusters", "", "", "Num Cache Clusters")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNumCacheNodes, "num-cache-nodes", "", "", "Num Cache Nodes")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheNumNodeGroups, "num-node-groups", "", "", "Num Node Groups")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheOfferingType, "offering-type", "", "", "Offering Type")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheOutpostMode, "outpost-mode", "", "", "Outpost Mode")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheParameterNameValues, "parameter-name-values", "", "", "Parameter Name Values")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticachePasswords, "passwords", "", nil, "Passwords")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePort, "port", "", "", "Port")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePreferredAvailabilityZone, "preferred-availability-zone", "", "", "Preferred Availability Zone")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticachePreferredAvailabilityZones, "preferred-availability-zones", "", nil, "Preferred Availability Zones")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticachePreferredCacheClusterAZs, "preferred-cache-cluster-azs", "", nil, "Preferred Cache Cluster Azs")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePreferredOutpostArn, "preferred-outpost-arn", "", "", "Preferred Outpost ARN")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticachePreferredOutpostArns, "preferred-outpost-arns", "", nil, "Preferred Outpost Arns")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePrimaryClusterId, "primary-cluster-id", "", "", "Primary Cluster ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePrimaryRegion, "primary-region", "", "", "Primary Region")
	_elasticacheCmd.Flags().StringVarP(&_elasticachePrimaryReplicationGroupId, "primary-replication-group-id", "", "", "Primary Replication Group ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheProductDescription, "product-description", "", "", "Product Description")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheRegionalConfigurations, "regional-configurations", "", "", "Regional Configurations")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheRemoveUserGroup, "remove-user-group", "", "", "Remove User Group")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheRemoveUserGroups, "remove-user-groups", "", "", "Remove User Groups")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReplicaConfiguration, "replica-configuration", "", "", "Replica Configuration")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReplicasPerNodeGroup, "replicas-per-node-group", "", "", "Replicas Per Node Group")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheReplicasToRemove, "replicas-to-remove", "", nil, "Replicas To Remove")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReplicationGroupDescription, "replication-group-description", "", "", "Replication Group Description")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReplicationGroupId, "replication-group-id", "", "", "Replication Group ID")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheReplicationGroupIds, "replication-group-ids", "", nil, "Replication Group Ids")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReplicationGroupRegion, "replication-group-region", "", "", "Replication Group Region")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReservedCacheNodeId, "reserved-cache-node-id", "", "", "Reserved Cache Node ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReservedCacheNodesOfferingId, "reserved-cache-nodes-offering-id", "", "", "Reserved Cache Nodes Offering ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheResetAllParameters, "reset-all-parameters", "", "", "Reset All Parameters")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheReshardingConfiguration, "resharding-configuration", "", "", "Resharding Configuration")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheResourceName, "resource-name", "", "", "Resource Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheRetainPrimaryCluster, "retain-primary-cluster", "", "", "Retain Primary Cluster")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheRetainPrimaryReplicationGroup, "retain-primary-replication-group", "", "", "Retain Primary Replication Group")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheScaleConfig, "scale-config", "", "", "Scale Config")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheServerlessCacheName, "serverless-cache-name", "", "", "Serverless Cache Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheServerlessCacheSnapshotName, "serverless-cache-snapshot-name", "", "", "Serverless Cache Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheServiceUpdateName, "service-update-name", "", "", "Service Update Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheServiceUpdateStatus, "service-update-status", "", "", "Service Update Status")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheServiceUpdateTimeRange, "service-update-time-range", "", "", "Service Update Time Range")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheShowCacheClustersNotInReplicationGroups, "show-cache-clusters-not-in-replication-groups", "", "", "Show Cache Clusters Not In Replication Groups")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheShowCacheNodeInfo, "show-cache-node-info", "", "", "Show Cache Node Info")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheShowMemberInfo, "show-member-info", "", "", "Show Member Info")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheShowNodeGroupConfig, "show-node-group-config", "", "", "Show Node Group Config")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheShowNodeLevelUpdateStatus, "show-node-level-update-status", "", "", "Show Node Level Update Status")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheSnapshotArns, "snapshot-arns", "", nil, "Snapshot Arns")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheSnapshotArnsToRestore, "snapshot-arns-to-restore", "", nil, "Snapshot Arns To Restore")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSnapshotName, "snapshot-name", "", "", "Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSnapshotRetentionLimit, "snapshot-retention-limit", "", "", "Snapshot Retention Limit")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSnapshotSource, "snapshot-source", "", "", "Snapshot Source")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSnapshotType, "snapshot-type", "", "", "Snapshot Type")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSnapshotWindow, "snapshot-window", "", "", "Snapshot Window")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSnapshottingClusterId, "snapshotting-cluster-id", "", "", "Snapshotting Cluster ID")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSource, "source", "", "", "Source")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSourceIdentifier, "source-identifier", "", "", "Source Identifier")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSourceServerlessCacheSnapshotName, "source-serverless-cache-snapshot-name", "", "", "Source Serverless Cache Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSourceSnapshotName, "source-snapshot-name", "", "", "Source Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheSourceType, "source-type", "", "", "Source Type")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheStartTime, "start-time", "", "", "Start Time")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheTagKeys, "tag-keys", "", nil, "Tag Keys")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheTags, "tags", "", "", "Tags")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheTargetBucket, "target-bucket", "", "", "Target Bucket")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheTargetServerlessCacheSnapshotName, "target-serverless-cache-snapshot-name", "", "", "Target Serverless Cache Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheTargetSnapshotName, "target-snapshot-name", "", "", "Target Snapshot Name")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheTransitEncryptionEnabled, "transit-encryption-enabled", "", "", "Transit Encryption Enabled")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheTransitEncryptionMode, "transit-encryption-mode", "", "", "Transit Encryption Mode")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheUpdateActionStatus, "update-action-status", "", "", "Update Action Status")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheUserGroupId, "user-group-id", "", "", "User Group ID")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheUserGroupIds, "user-group-ids", "", nil, "User Group Ids")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheUserGroupIdsToAdd, "user-group-ids-to-add", "", nil, "User Group Ids To Add")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheUserGroupIdsToRemove, "user-group-ids-to-remove", "", nil, "User Group Ids To Remove")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheUserId, "user-id", "", "", "User ID")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheUserIds, "user-ids", "", nil, "User Ids")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheUserIdsToAdd, "user-ids-to-add", "", nil, "User Ids To Add")
	_elasticacheCmd.Flags().StringSliceVarP(&_elasticacheUserIdsToRemove, "user-ids-to-remove", "", nil, "User Ids To Remove")
	_elasticacheCmd.Flags().StringVarP(&_elasticacheUserName, "user-name", "", "", "User Name")

	_elasticacheCmd.Flags().BoolVarP(&_elasticacheAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheAuthorizeCacheSecurityGroupIngress, "authorize-cache-security-group-ingress", "", false, "Authorize Cache Security Group Ingress")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheBatchApplyUpdateAction, "batch-apply-update-action", "", false, "Batch Apply Update Action")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheBatchStopUpdateAction, "batch-stop-update-action", "", false, "Batch Stop Update Action")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCompleteMigration, "complete-migration", "", false, "Complete Migration")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCopyServerlessCacheSnapshot, "copy-serverless-cache-snapshot", "", false, "Copy Serverless Cache Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCopySnapshot, "copy-snapshot", "", false, "Copy Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateCacheCluster, "create-cache-cluster", "", false, "Create Cache Cluster")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateCacheParameterGroup, "create-cache-parameter-group", "", false, "Create Cache Parameter Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateCacheSecurityGroup, "create-cache-security-group", "", false, "Create Cache Security Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateCacheSubnetGroup, "create-cache-subnet-group", "", false, "Create Cache Subnet Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateGlobalReplicationGroup, "create-global-replication-group", "", false, "Create Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateReplicationGroup, "create-replication-group", "", false, "Create Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateServerlessCache, "create-serverless-cache", "", false, "Create Serverless Cache")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateServerlessCacheSnapshot, "create-serverless-cache-snapshot", "", false, "Create Serverless Cache Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateUser, "create-user", "", false, "Create User")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheCreateUserGroup, "create-user-group", "", false, "Create User Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDecreaseNodeGroupsInGlobalReplicationGroup, "decrease-node-groups-in-global-replication-group", "", false, "Decrease Node Groups In Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDecreaseReplicaCount, "decrease-replica-count", "", false, "Decrease Replica Count")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteCacheCluster, "delete-cache-cluster", "", false, "Delete Cache Cluster")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteCacheParameterGroup, "delete-cache-parameter-group", "", false, "Delete Cache Parameter Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteCacheSecurityGroup, "delete-cache-security-group", "", false, "Delete Cache Security Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteCacheSubnetGroup, "delete-cache-subnet-group", "", false, "Delete Cache Subnet Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteGlobalReplicationGroup, "delete-global-replication-group", "", false, "Delete Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteReplicationGroup, "delete-replication-group", "", false, "Delete Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteServerlessCache, "delete-serverless-cache", "", false, "Delete Serverless Cache")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteServerlessCacheSnapshot, "delete-serverless-cache-snapshot", "", false, "Delete Serverless Cache Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteSnapshot, "delete-snapshot", "", false, "Delete Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteUser, "delete-user", "", false, "Delete User")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDeleteUserGroup, "delete-user-group", "", false, "Delete User Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeCacheClusters, "describe-cache-clusters", "", false, "Describe Cache Clusters")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeCacheEngineVersions, "describe-cache-engine-versions", "", false, "Describe Cache Engine Versions")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeCacheParameterGroups, "describe-cache-parameter-groups", "", false, "Describe Cache Parameter Groups")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeCacheParameters, "describe-cache-parameters", "", false, "Describe Cache Parameters")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeCacheSecurityGroups, "describe-cache-security-groups", "", false, "Describe Cache Security Groups")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeCacheSubnetGroups, "describe-cache-subnet-groups", "", false, "Describe Cache Subnet Groups")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeEngineDefaultParameters, "describe-engine-default-parameters", "", false, "Describe Engine Default Parameters")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeEvents, "describe-events", "", false, "Describe Events")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeGlobalReplicationGroups, "describe-global-replication-groups", "", false, "Describe Global Replication Groups")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeReplicationGroups, "describe-replication-groups", "", false, "Describe Replication Groups")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeReservedCacheNodes, "describe-reserved-cache-nodes", "", false, "Describe Reserved Cache Nodes")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeReservedCacheNodesOfferings, "describe-reserved-cache-nodes-offerings", "", false, "Describe Reserved Cache Nodes Offerings")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeServerlessCacheSnapshots, "describe-serverless-cache-snapshots", "", false, "Describe Serverless Cache Snapshots")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeServerlessCaches, "describe-serverless-caches", "", false, "Describe Serverless Caches")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeServiceUpdates, "describe-service-updates", "", false, "Describe Service Updates")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeSnapshots, "describe-snapshots", "", false, "Describe Snapshots")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeUpdateActions, "describe-update-actions", "", false, "Describe Update Actions")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeUserGroups, "describe-user-groups", "", false, "Describe User Groups")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDescribeUsers, "describe-users", "", false, "Describe Users")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheDisassociateGlobalReplicationGroup, "disassociate-global-replication-group", "", false, "Disassociate Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheExportServerlessCacheSnapshot, "export-serverless-cache-snapshot", "", false, "Export Serverless Cache Snapshot")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheFailoverGlobalReplicationGroup, "failover-global-replication-group", "", false, "Failover Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheIncreaseNodeGroupsInGlobalReplicationGroup, "increase-node-groups-in-global-replication-group", "", false, "Increase Node Groups In Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheIncreaseReplicaCount, "increase-replica-count", "", false, "Increase Replica Count")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheListAllowedNodeTypeModifications, "list-allowed-node-type-modifications", "", false, "List Allowed Node Type Modifications")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyCacheCluster, "modify-cache-cluster", "", false, "Modify Cache Cluster")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyCacheParameterGroup, "modify-cache-parameter-group", "", false, "Modify Cache Parameter Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyCacheSubnetGroup, "modify-cache-subnet-group", "", false, "Modify Cache Subnet Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyGlobalReplicationGroup, "modify-global-replication-group", "", false, "Modify Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyReplicationGroup, "modify-replication-group", "", false, "Modify Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyReplicationGroupShardConfiguration, "modify-replication-group-shard-configuration", "", false, "Modify Replication Group Shard Configuration")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyServerlessCache, "modify-serverless-cache", "", false, "Modify Serverless Cache")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyUser, "modify-user", "", false, "Modify User")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheModifyUserGroup, "modify-user-group", "", false, "Modify User Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticachePurchaseReservedCacheNodesOffering, "purchase-reserved-cache-nodes-offering", "", false, "Purchase Reserved Cache Nodes Offering")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheRebalanceSlotsInGlobalReplicationGroup, "rebalance-slots-in-global-replication-group", "", false, "Rebalance Slots In Global Replication Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheRebootCacheCluster, "reboot-cache-cluster", "", false, "Reboot Cache Cluster")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheResetCacheParameterGroup, "reset-cache-parameter-group", "", false, "Reset Cache Parameter Group")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheRevokeCacheSecurityGroupIngress, "revoke-cache-security-group-ingress", "", false, "Revoke Cache Security Group Ingress")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheStartMigration, "start-migration", "", false, "Start Migration")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheTestFailover, "test-failover", "", false, "Test Failover")
	_elasticacheCmd.Flags().BoolVarP(&_elasticacheTestMigration, "test-migration", "", false, "Test Migration")

}
