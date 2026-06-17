package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rdsCmd represents the rds command
var _rdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "AWS rds CLI",
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
		client := rds.NewFromConfig(cfg)
		if _rdsAddRoleToDBCluster {
			rds_AddRoleToDBCluster(cfg, client)
			return
		}
		if _rdsAddRoleToDBInstance {
			rds_AddRoleToDBInstance(cfg, client)
			return
		}
		if _rdsAddSourceIdentifierToSubscription {
			rds_AddSourceIdentifierToSubscription(cfg, client)
			return
		}
		if _rdsAddTagsToResource {
			rds_AddTagsToResource(cfg, client)
			return
		}
		if _rdsApplyPendingMaintenanceAction {
			rds_ApplyPendingMaintenanceAction(cfg, client)
			return
		}
		if _rdsAuthorizeDBSecurityGroupIngress {
			rds_AuthorizeDBSecurityGroupIngress(cfg, client)
			return
		}
		if _rdsBacktrackDBCluster {
			rds_BacktrackDBCluster(cfg, client)
			return
		}
		if _rdsCancelExportTask {
			rds_CancelExportTask(cfg, client)
			return
		}
		if _rdsCopyDBClusterParameterGroup {
			rds_CopyDBClusterParameterGroup(cfg, client)
			return
		}
		if _rdsCopyDBClusterSnapshot {
			rds_CopyDBClusterSnapshot(cfg, client)
			return
		}
		if _rdsCopyDBParameterGroup {
			rds_CopyDBParameterGroup(cfg, client)
			return
		}
		if _rdsCopyDBSnapshot {
			rds_CopyDBSnapshot(cfg, client)
			return
		}
		if _rdsCopyOptionGroup {
			rds_CopyOptionGroup(cfg, client)
			return
		}
		if _rdsCreateBlueGreenDeployment {
			rds_CreateBlueGreenDeployment(cfg, client)
			return
		}
		if _rdsCreateCustomDBEngineVersion {
			rds_CreateCustomDBEngineVersion(cfg, client)
			return
		}
		if _rdsCreateDBCluster {
			rds_CreateDBCluster(cfg, client)
			return
		}
		if _rdsCreateDBClusterEndpoint {
			rds_CreateDBClusterEndpoint(cfg, client)
			return
		}
		if _rdsCreateDBClusterParameterGroup {
			rds_CreateDBClusterParameterGroup(cfg, client)
			return
		}
		if _rdsCreateDBClusterSnapshot {
			rds_CreateDBClusterSnapshot(cfg, client)
			return
		}
		if _rdsCreateDBInstance {
			rds_CreateDBInstance(cfg, client)
			return
		}
		if _rdsCreateDBInstanceReadReplica {
			rds_CreateDBInstanceReadReplica(cfg, client)
			return
		}
		if _rdsCreateDBParameterGroup {
			rds_CreateDBParameterGroup(cfg, client)
			return
		}
		if _rdsCreateDBProxy {
			rds_CreateDBProxy(cfg, client)
			return
		}
		if _rdsCreateDBProxyEndpoint {
			rds_CreateDBProxyEndpoint(cfg, client)
			return
		}
		if _rdsCreateDBSecurityGroup {
			rds_CreateDBSecurityGroup(cfg, client)
			return
		}
		if _rdsCreateDBShardGroup {
			rds_CreateDBShardGroup(cfg, client)
			return
		}
		if _rdsCreateDBSnapshot {
			rds_CreateDBSnapshot(cfg, client)
			return
		}
		if _rdsCreateDBSubnetGroup {
			rds_CreateDBSubnetGroup(cfg, client)
			return
		}
		if _rdsCreateEventSubscription {
			rds_CreateEventSubscription(cfg, client)
			return
		}
		if _rdsCreateGlobalCluster {
			rds_CreateGlobalCluster(cfg, client)
			return
		}
		if _rdsCreateIntegration {
			rds_CreateIntegration(cfg, client)
			return
		}
		if _rdsCreateOptionGroup {
			rds_CreateOptionGroup(cfg, client)
			return
		}
		if _rdsCreateTenantDatabase {
			rds_CreateTenantDatabase(cfg, client)
			return
		}
		if _rdsDeleteBlueGreenDeployment {
			rds_DeleteBlueGreenDeployment(cfg, client)
			return
		}
		if _rdsDeleteCustomDBEngineVersion {
			rds_DeleteCustomDBEngineVersion(cfg, client)
			return
		}
		if _rdsDeleteDBCluster {
			rds_DeleteDBCluster(cfg, client)
			return
		}
		if _rdsDeleteDBClusterAutomatedBackup {
			rds_DeleteDBClusterAutomatedBackup(cfg, client)
			return
		}
		if _rdsDeleteDBClusterEndpoint {
			rds_DeleteDBClusterEndpoint(cfg, client)
			return
		}
		if _rdsDeleteDBClusterParameterGroup {
			rds_DeleteDBClusterParameterGroup(cfg, client)
			return
		}
		if _rdsDeleteDBClusterSnapshot {
			rds_DeleteDBClusterSnapshot(cfg, client)
			return
		}
		if _rdsDeleteDBInstance {
			rds_DeleteDBInstance(cfg, client)
			return
		}
		if _rdsDeleteDBInstanceAutomatedBackup {
			rds_DeleteDBInstanceAutomatedBackup(cfg, client)
			return
		}
		if _rdsDeleteDBParameterGroup {
			rds_DeleteDBParameterGroup(cfg, client)
			return
		}
		if _rdsDeleteDBProxy {
			rds_DeleteDBProxy(cfg, client)
			return
		}
		if _rdsDeleteDBProxyEndpoint {
			rds_DeleteDBProxyEndpoint(cfg, client)
			return
		}
		if _rdsDeleteDBSecurityGroup {
			rds_DeleteDBSecurityGroup(cfg, client)
			return
		}
		if _rdsDeleteDBShardGroup {
			rds_DeleteDBShardGroup(cfg, client)
			return
		}
		if _rdsDeleteDBSnapshot {
			rds_DeleteDBSnapshot(cfg, client)
			return
		}
		if _rdsDeleteDBSubnetGroup {
			rds_DeleteDBSubnetGroup(cfg, client)
			return
		}
		if _rdsDeleteEventSubscription {
			rds_DeleteEventSubscription(cfg, client)
			return
		}
		if _rdsDeleteGlobalCluster {
			rds_DeleteGlobalCluster(cfg, client)
			return
		}
		if _rdsDeleteIntegration {
			rds_DeleteIntegration(cfg, client)
			return
		}
		if _rdsDeleteOptionGroup {
			rds_DeleteOptionGroup(cfg, client)
			return
		}
		if _rdsDeleteTenantDatabase {
			rds_DeleteTenantDatabase(cfg, client)
			return
		}
		if _rdsDeregisterDBProxyTargets {
			rds_DeregisterDBProxyTargets(cfg, client)
			return
		}
		if _rdsDescribeAccountAttributes {
			rds_DescribeAccountAttributes(cfg, client)
			return
		}
		if _rdsDescribeBlueGreenDeployments {
			rds_DescribeBlueGreenDeployments(cfg, client)
			return
		}
		if _rdsDescribeCertificates {
			rds_DescribeCertificates(cfg, client)
			return
		}
		if _rdsDescribeDBClusterAutomatedBackups {
			rds_DescribeDBClusterAutomatedBackups(cfg, client)
			return
		}
		if _rdsDescribeDBClusterBacktracks {
			rds_DescribeDBClusterBacktracks(cfg, client)
			return
		}
		if _rdsDescribeDBClusterEndpoints {
			rds_DescribeDBClusterEndpoints(cfg, client)
			return
		}
		if _rdsDescribeDBClusterParameterGroups {
			rds_DescribeDBClusterParameterGroups(cfg, client)
			return
		}
		if _rdsDescribeDBClusterParameters {
			rds_DescribeDBClusterParameters(cfg, client)
			return
		}
		if _rdsDescribeDBClusterSnapshotAttributes {
			rds_DescribeDBClusterSnapshotAttributes(cfg, client)
			return
		}
		if _rdsDescribeDBClusterSnapshots {
			rds_DescribeDBClusterSnapshots(cfg, client)
			return
		}
		if _rdsDescribeDBClusters {
			rds_DescribeDBClusters(cfg, client)
			return
		}
		if _rdsDescribeDBEngineVersions {
			rds_DescribeDBEngineVersions(cfg, client)
			return
		}
		if _rdsDescribeDBInstanceAutomatedBackups {
			rds_DescribeDBInstanceAutomatedBackups(cfg, client)
			return
		}
		if _rdsDescribeDBInstances {
			rds_DescribeDBInstances(cfg, client)
			return
		}
		if _rdsDescribeDBLogFiles {
			rds_DescribeDBLogFiles(cfg, client)
			return
		}
		if _rdsDescribeDBMajorEngineVersions {
			rds_DescribeDBMajorEngineVersions(cfg, client)
			return
		}
		if _rdsDescribeDBParameterGroups {
			rds_DescribeDBParameterGroups(cfg, client)
			return
		}
		if _rdsDescribeDBParameters {
			rds_DescribeDBParameters(cfg, client)
			return
		}
		if _rdsDescribeDBProxies {
			rds_DescribeDBProxies(cfg, client)
			return
		}
		if _rdsDescribeDBProxyEndpoints {
			rds_DescribeDBProxyEndpoints(cfg, client)
			return
		}
		if _rdsDescribeDBProxyTargetGroups {
			rds_DescribeDBProxyTargetGroups(cfg, client)
			return
		}
		if _rdsDescribeDBProxyTargets {
			rds_DescribeDBProxyTargets(cfg, client)
			return
		}
		if _rdsDescribeDBRecommendations {
			rds_DescribeDBRecommendations(cfg, client)
			return
		}
		if _rdsDescribeDBSecurityGroups {
			rds_DescribeDBSecurityGroups(cfg, client)
			return
		}
		if _rdsDescribeDBShardGroups {
			rds_DescribeDBShardGroups(cfg, client)
			return
		}
		if _rdsDescribeDBSnapshotAttributes {
			rds_DescribeDBSnapshotAttributes(cfg, client)
			return
		}
		if _rdsDescribeDBSnapshotTenantDatabases {
			rds_DescribeDBSnapshotTenantDatabases(cfg, client)
			return
		}
		if _rdsDescribeDBSnapshots {
			rds_DescribeDBSnapshots(cfg, client)
			return
		}
		if _rdsDescribeDBSubnetGroups {
			rds_DescribeDBSubnetGroups(cfg, client)
			return
		}
		if _rdsDescribeEngineDefaultClusterParameters {
			rds_DescribeEngineDefaultClusterParameters(cfg, client)
			return
		}
		if _rdsDescribeEngineDefaultParameters {
			rds_DescribeEngineDefaultParameters(cfg, client)
			return
		}
		if _rdsDescribeEventCategories {
			rds_DescribeEventCategories(cfg, client)
			return
		}
		if _rdsDescribeEventSubscriptions {
			rds_DescribeEventSubscriptions(cfg, client)
			return
		}
		if _rdsDescribeEvents {
			rds_DescribeEvents(cfg, client)
			return
		}
		if _rdsDescribeExportTasks {
			rds_DescribeExportTasks(cfg, client)
			return
		}
		if _rdsDescribeGlobalClusters {
			rds_DescribeGlobalClusters(cfg, client)
			return
		}
		if _rdsDescribeIntegrations {
			rds_DescribeIntegrations(cfg, client)
			return
		}
		if _rdsDescribeOptionGroupOptions {
			rds_DescribeOptionGroupOptions(cfg, client)
			return
		}
		if _rdsDescribeOptionGroups {
			rds_DescribeOptionGroups(cfg, client)
			return
		}
		if _rdsDescribeOrderableDBInstanceOptions {
			rds_DescribeOrderableDBInstanceOptions(cfg, client)
			return
		}
		if _rdsDescribePendingMaintenanceActions {
			rds_DescribePendingMaintenanceActions(cfg, client)
			return
		}
		if _rdsDescribeReservedDBInstances {
			rds_DescribeReservedDBInstances(cfg, client)
			return
		}
		if _rdsDescribeReservedDBInstancesOfferings {
			rds_DescribeReservedDBInstancesOfferings(cfg, client)
			return
		}
		if _rdsDescribeSourceRegions {
			rds_DescribeSourceRegions(cfg, client)
			return
		}
		if _rdsDescribeTenantDatabases {
			rds_DescribeTenantDatabases(cfg, client)
			return
		}
		if _rdsDescribeValidDBInstanceModifications {
			rds_DescribeValidDBInstanceModifications(cfg, client)
			return
		}
		if _rdsDisableHttpEndpoint {
			rds_DisableHttpEndpoint(cfg, client)
			return
		}
		if _rdsDownloadDBLogFilePortion {
			rds_DownloadDBLogFilePortion(cfg, client)
			return
		}
		if _rdsEnableHttpEndpoint {
			rds_EnableHttpEndpoint(cfg, client)
			return
		}
		if _rdsFailoverDBCluster {
			rds_FailoverDBCluster(cfg, client)
			return
		}
		if _rdsFailoverGlobalCluster {
			rds_FailoverGlobalCluster(cfg, client)
			return
		}
		if _rdsListTagsForResource {
			rds_ListTagsForResource(cfg, client)
			return
		}
		if _rdsModifyActivityStream {
			rds_ModifyActivityStream(cfg, client)
			return
		}
		if _rdsModifyCertificates {
			rds_ModifyCertificates(cfg, client)
			return
		}
		if _rdsModifyCurrentDBClusterCapacity {
			rds_ModifyCurrentDBClusterCapacity(cfg, client)
			return
		}
		if _rdsModifyCustomDBEngineVersion {
			rds_ModifyCustomDBEngineVersion(cfg, client)
			return
		}
		if _rdsModifyDBCluster {
			rds_ModifyDBCluster(cfg, client)
			return
		}
		if _rdsModifyDBClusterEndpoint {
			rds_ModifyDBClusterEndpoint(cfg, client)
			return
		}
		if _rdsModifyDBClusterParameterGroup {
			rds_ModifyDBClusterParameterGroup(cfg, client)
			return
		}
		if _rdsModifyDBClusterSnapshotAttribute {
			rds_ModifyDBClusterSnapshotAttribute(cfg, client)
			return
		}
		if _rdsModifyDBInstance {
			rds_ModifyDBInstance(cfg, client)
			return
		}
		if _rdsModifyDBParameterGroup {
			rds_ModifyDBParameterGroup(cfg, client)
			return
		}
		if _rdsModifyDBProxy {
			rds_ModifyDBProxy(cfg, client)
			return
		}
		if _rdsModifyDBProxyEndpoint {
			rds_ModifyDBProxyEndpoint(cfg, client)
			return
		}
		if _rdsModifyDBProxyTargetGroup {
			rds_ModifyDBProxyTargetGroup(cfg, client)
			return
		}
		if _rdsModifyDBRecommendation {
			rds_ModifyDBRecommendation(cfg, client)
			return
		}
		if _rdsModifyDBShardGroup {
			rds_ModifyDBShardGroup(cfg, client)
			return
		}
		if _rdsModifyDBSnapshot {
			rds_ModifyDBSnapshot(cfg, client)
			return
		}
		if _rdsModifyDBSnapshotAttribute {
			rds_ModifyDBSnapshotAttribute(cfg, client)
			return
		}
		if _rdsModifyDBSubnetGroup {
			rds_ModifyDBSubnetGroup(cfg, client)
			return
		}
		if _rdsModifyEventSubscription {
			rds_ModifyEventSubscription(cfg, client)
			return
		}
		if _rdsModifyGlobalCluster {
			rds_ModifyGlobalCluster(cfg, client)
			return
		}
		if _rdsModifyIntegration {
			rds_ModifyIntegration(cfg, client)
			return
		}
		if _rdsModifyOptionGroup {
			rds_ModifyOptionGroup(cfg, client)
			return
		}
		if _rdsModifyTenantDatabase {
			rds_ModifyTenantDatabase(cfg, client)
			return
		}
		if _rdsPromoteReadReplica {
			rds_PromoteReadReplica(cfg, client)
			return
		}
		if _rdsPromoteReadReplicaDBCluster {
			rds_PromoteReadReplicaDBCluster(cfg, client)
			return
		}
		if _rdsPurchaseReservedDBInstancesOffering {
			rds_PurchaseReservedDBInstancesOffering(cfg, client)
			return
		}
		if _rdsRebootDBCluster {
			rds_RebootDBCluster(cfg, client)
			return
		}
		if _rdsRebootDBInstance {
			rds_RebootDBInstance(cfg, client)
			return
		}
		if _rdsRebootDBShardGroup {
			rds_RebootDBShardGroup(cfg, client)
			return
		}
		if _rdsRegisterDBProxyTargets {
			rds_RegisterDBProxyTargets(cfg, client)
			return
		}
		if _rdsRemoveFromGlobalCluster {
			rds_RemoveFromGlobalCluster(cfg, client)
			return
		}
		if _rdsRemoveRoleFromDBCluster {
			rds_RemoveRoleFromDBCluster(cfg, client)
			return
		}
		if _rdsRemoveRoleFromDBInstance {
			rds_RemoveRoleFromDBInstance(cfg, client)
			return
		}
		if _rdsRemoveSourceIdentifierFromSubscription {
			rds_RemoveSourceIdentifierFromSubscription(cfg, client)
			return
		}
		if _rdsRemoveTagsFromResource {
			rds_RemoveTagsFromResource(cfg, client)
			return
		}
		if _rdsResetDBClusterParameterGroup {
			rds_ResetDBClusterParameterGroup(cfg, client)
			return
		}
		if _rdsResetDBParameterGroup {
			rds_ResetDBParameterGroup(cfg, client)
			return
		}
		if _rdsRestoreDBClusterFromS3 {
			rds_RestoreDBClusterFromS3(cfg, client)
			return
		}
		if _rdsRestoreDBClusterFromSnapshot {
			rds_RestoreDBClusterFromSnapshot(cfg, client)
			return
		}
		if _rdsRestoreDBClusterToPointInTime {
			rds_RestoreDBClusterToPointInTime(cfg, client)
			return
		}
		if _rdsRestoreDBInstanceFromDBSnapshot {
			rds_RestoreDBInstanceFromDBSnapshot(cfg, client)
			return
		}
		if _rdsRestoreDBInstanceFromS3 {
			rds_RestoreDBInstanceFromS3(cfg, client)
			return
		}
		if _rdsRestoreDBInstanceToPointInTime {
			rds_RestoreDBInstanceToPointInTime(cfg, client)
			return
		}
		if _rdsRevokeDBSecurityGroupIngress {
			rds_RevokeDBSecurityGroupIngress(cfg, client)
			return
		}
		if _rdsStartActivityStream {
			rds_StartActivityStream(cfg, client)
			return
		}
		if _rdsStartDBCluster {
			rds_StartDBCluster(cfg, client)
			return
		}
		if _rdsStartDBInstance {
			rds_StartDBInstance(cfg, client)
			return
		}
		if _rdsStartDBInstanceAutomatedBackupsReplication {
			rds_StartDBInstanceAutomatedBackupsReplication(cfg, client)
			return
		}
		if _rdsStartExportTask {
			rds_StartExportTask(cfg, client)
			return
		}
		if _rdsStopActivityStream {
			rds_StopActivityStream(cfg, client)
			return
		}
		if _rdsStopDBCluster {
			rds_StopDBCluster(cfg, client)
			return
		}
		if _rdsStopDBInstance {
			rds_StopDBInstance(cfg, client)
			return
		}
		if _rdsStopDBInstanceAutomatedBackupsReplication {
			rds_StopDBInstanceAutomatedBackupsReplication(cfg, client)
			return
		}
		if _rdsSwitchoverBlueGreenDeployment {
			rds_SwitchoverBlueGreenDeployment(cfg, client)
			return
		}
		if _rdsSwitchoverGlobalCluster {
			rds_SwitchoverGlobalCluster(cfg, client)
			return
		}
		if _rdsSwitchoverReadReplica {
			rds_SwitchoverReadReplica(cfg, client)
			return
		}

	},
}

var (
	_rdsAddRoleToDBCluster                         bool
	_rdsAddRoleToDBInstance                        bool
	_rdsAddSourceIdentifierToSubscription          bool
	_rdsAddTagsToResource                          bool
	_rdsApplyPendingMaintenanceAction              bool
	_rdsAuthorizeDBSecurityGroupIngress            bool
	_rdsBacktrackDBCluster                         bool
	_rdsCancelExportTask                           bool
	_rdsCopyDBClusterParameterGroup                bool
	_rdsCopyDBClusterSnapshot                      bool
	_rdsCopyDBParameterGroup                       bool
	_rdsCopyDBSnapshot                             bool
	_rdsCopyOptionGroup                            bool
	_rdsCreateBlueGreenDeployment                  bool
	_rdsCreateCustomDBEngineVersion                bool
	_rdsCreateDBCluster                            bool
	_rdsCreateDBClusterEndpoint                    bool
	_rdsCreateDBClusterParameterGroup              bool
	_rdsCreateDBClusterSnapshot                    bool
	_rdsCreateDBInstance                           bool
	_rdsCreateDBInstanceReadReplica                bool
	_rdsCreateDBParameterGroup                     bool
	_rdsCreateDBProxy                              bool
	_rdsCreateDBProxyEndpoint                      bool
	_rdsCreateDBSecurityGroup                      bool
	_rdsCreateDBShardGroup                         bool
	_rdsCreateDBSnapshot                           bool
	_rdsCreateDBSubnetGroup                        bool
	_rdsCreateEventSubscription                    bool
	_rdsCreateGlobalCluster                        bool
	_rdsCreateIntegration                          bool
	_rdsCreateOptionGroup                          bool
	_rdsCreateTenantDatabase                       bool
	_rdsDeleteBlueGreenDeployment                  bool
	_rdsDeleteCustomDBEngineVersion                bool
	_rdsDeleteDBCluster                            bool
	_rdsDeleteDBClusterAutomatedBackup             bool
	_rdsDeleteDBClusterEndpoint                    bool
	_rdsDeleteDBClusterParameterGroup              bool
	_rdsDeleteDBClusterSnapshot                    bool
	_rdsDeleteDBInstance                           bool
	_rdsDeleteDBInstanceAutomatedBackup            bool
	_rdsDeleteDBParameterGroup                     bool
	_rdsDeleteDBProxy                              bool
	_rdsDeleteDBProxyEndpoint                      bool
	_rdsDeleteDBSecurityGroup                      bool
	_rdsDeleteDBShardGroup                         bool
	_rdsDeleteDBSnapshot                           bool
	_rdsDeleteDBSubnetGroup                        bool
	_rdsDeleteEventSubscription                    bool
	_rdsDeleteGlobalCluster                        bool
	_rdsDeleteIntegration                          bool
	_rdsDeleteOptionGroup                          bool
	_rdsDeleteTenantDatabase                       bool
	_rdsDeregisterDBProxyTargets                   bool
	_rdsDescribeAccountAttributes                  bool
	_rdsDescribeBlueGreenDeployments               bool
	_rdsDescribeCertificates                       bool
	_rdsDescribeDBClusterAutomatedBackups          bool
	_rdsDescribeDBClusterBacktracks                bool
	_rdsDescribeDBClusterEndpoints                 bool
	_rdsDescribeDBClusterParameterGroups           bool
	_rdsDescribeDBClusterParameters                bool
	_rdsDescribeDBClusterSnapshotAttributes        bool
	_rdsDescribeDBClusterSnapshots                 bool
	_rdsDescribeDBClusters                         bool
	_rdsDescribeDBEngineVersions                   bool
	_rdsDescribeDBInstanceAutomatedBackups         bool
	_rdsDescribeDBInstances                        bool
	_rdsDescribeDBLogFiles                         bool
	_rdsDescribeDBMajorEngineVersions              bool
	_rdsDescribeDBParameterGroups                  bool
	_rdsDescribeDBParameters                       bool
	_rdsDescribeDBProxies                          bool
	_rdsDescribeDBProxyEndpoints                   bool
	_rdsDescribeDBProxyTargetGroups                bool
	_rdsDescribeDBProxyTargets                     bool
	_rdsDescribeDBRecommendations                  bool
	_rdsDescribeDBSecurityGroups                   bool
	_rdsDescribeDBShardGroups                      bool
	_rdsDescribeDBSnapshotAttributes               bool
	_rdsDescribeDBSnapshotTenantDatabases          bool
	_rdsDescribeDBSnapshots                        bool
	_rdsDescribeDBSubnetGroups                     bool
	_rdsDescribeEngineDefaultClusterParameters     bool
	_rdsDescribeEngineDefaultParameters            bool
	_rdsDescribeEventCategories                    bool
	_rdsDescribeEventSubscriptions                 bool
	_rdsDescribeEvents                             bool
	_rdsDescribeExportTasks                        bool
	_rdsDescribeGlobalClusters                     bool
	_rdsDescribeIntegrations                       bool
	_rdsDescribeOptionGroupOptions                 bool
	_rdsDescribeOptionGroups                       bool
	_rdsDescribeOrderableDBInstanceOptions         bool
	_rdsDescribePendingMaintenanceActions          bool
	_rdsDescribeReservedDBInstances                bool
	_rdsDescribeReservedDBInstancesOfferings       bool
	_rdsDescribeSourceRegions                      bool
	_rdsDescribeTenantDatabases                    bool
	_rdsDescribeValidDBInstanceModifications       bool
	_rdsDisableHttpEndpoint                        bool
	_rdsDownloadDBLogFilePortion                   bool
	_rdsEnableHttpEndpoint                         bool
	_rdsFailoverDBCluster                          bool
	_rdsFailoverGlobalCluster                      bool
	_rdsListTagsForResource                        bool
	_rdsModifyActivityStream                       bool
	_rdsModifyCertificates                         bool
	_rdsModifyCurrentDBClusterCapacity             bool
	_rdsModifyCustomDBEngineVersion                bool
	_rdsModifyDBCluster                            bool
	_rdsModifyDBClusterEndpoint                    bool
	_rdsModifyDBClusterParameterGroup              bool
	_rdsModifyDBClusterSnapshotAttribute           bool
	_rdsModifyDBInstance                           bool
	_rdsModifyDBParameterGroup                     bool
	_rdsModifyDBProxy                              bool
	_rdsModifyDBProxyEndpoint                      bool
	_rdsModifyDBProxyTargetGroup                   bool
	_rdsModifyDBRecommendation                     bool
	_rdsModifyDBShardGroup                         bool
	_rdsModifyDBSnapshot                           bool
	_rdsModifyDBSnapshotAttribute                  bool
	_rdsModifyDBSubnetGroup                        bool
	_rdsModifyEventSubscription                    bool
	_rdsModifyGlobalCluster                        bool
	_rdsModifyIntegration                          bool
	_rdsModifyOptionGroup                          bool
	_rdsModifyTenantDatabase                       bool
	_rdsPromoteReadReplica                         bool
	_rdsPromoteReadReplicaDBCluster                bool
	_rdsPurchaseReservedDBInstancesOffering        bool
	_rdsRebootDBCluster                            bool
	_rdsRebootDBInstance                           bool
	_rdsRebootDBShardGroup                         bool
	_rdsRegisterDBProxyTargets                     bool
	_rdsRemoveFromGlobalCluster                    bool
	_rdsRemoveRoleFromDBCluster                    bool
	_rdsRemoveRoleFromDBInstance                   bool
	_rdsRemoveSourceIdentifierFromSubscription     bool
	_rdsRemoveTagsFromResource                     bool
	_rdsResetDBClusterParameterGroup               bool
	_rdsResetDBParameterGroup                      bool
	_rdsRestoreDBClusterFromS3                     bool
	_rdsRestoreDBClusterFromSnapshot               bool
	_rdsRestoreDBClusterToPointInTime              bool
	_rdsRestoreDBInstanceFromDBSnapshot            bool
	_rdsRestoreDBInstanceFromS3                    bool
	_rdsRestoreDBInstanceToPointInTime             bool
	_rdsRevokeDBSecurityGroupIngress               bool
	_rdsStartActivityStream                        bool
	_rdsStartDBCluster                             bool
	_rdsStartDBInstance                            bool
	_rdsStartDBInstanceAutomatedBackupsReplication bool
	_rdsStartExportTask                            bool
	_rdsStopActivityStream                         bool
	_rdsStopDBCluster                              bool
	_rdsStopDBInstance                             bool
	_rdsStopDBInstanceAutomatedBackupsReplication  bool
	_rdsSwitchoverBlueGreenDeployment              bool
	_rdsSwitchoverGlobalCluster                    bool
	_rdsSwitchoverReadReplica                      bool

	_rdsAdditionalEncryptionContext              string
	_rdsAdditionalStorageVolumes                 string
	_rdsAllocatedStorage                         string
	_rdsAllowDataLoss                            string
	_rdsAllowEngineModeChange                    string
	_rdsAllowMajorVersionUpgrade                 string
	_rdsApplyAction                              string
	_rdsApplyImmediately                         string
	_rdsAttributeName                            string
	_rdsAuditPolicyState                         string
	_rdsAuth                                     string
	_rdsAutoMinorVersionUpgrade                  string
	_rdsAutomationMode                           string
	_rdsAvailabilityZone                         string
	_rdsAvailabilityZoneGroup                    string
	_rdsAvailabilityZones                        []string
	_rdsAwsBackupRecoveryPointArn                string
	_rdsBacktrackIdentifier                      string
	_rdsBacktrackTo                              string
	_rdsBacktrackWindow                          string
	_rdsBackupRetentionPeriod                    string
	_rdsBackupTarget                             string
	_rdsBlueGreenDeploymentIdentifier            string
	_rdsBlueGreenDeploymentName                  string
	_rdsCACertificateIdentifier                  string
	_rdsCapacity                                 string
	_rdsCertificateIdentifier                    string
	_rdsCertificateRotationRestart               string
	_rdsCharacterSetName                         string
	_rdsCIDRIP                                   string
	_rdsCloudwatchLogsExportConfiguration        string
	_rdsClusterScalabilityType                   string
	_rdsComputeRedundancy                        string
	_rdsConnectionPoolConfig                     string
	_rdsCopyTags                                 string
	_rdsCopyTagsToSnapshot                       string
	_rdsCustomIamInstanceProfile                 string
	_rdsDataFilter                               string
	_rdsDatabaseInsightsMode                     string
	_rdsDatabaseInstallationFiles                []string
	_rdsDatabaseInstallationFilesS3BucketName    string
	_rdsDatabaseInstallationFilesS3Prefix        string
	_rdsDatabaseName                             string
	_rdsDBClusterEndpointIdentifier              string
	_rdsDBClusterIdentifier                      string
	_rdsDBClusterIdentifiers                     []string
	_rdsDBClusterInstanceClass                   string
	_rdsDBClusterParameterGroupName              string
	_rdsDbClusterResourceId                      string
	_rdsDBClusterSnapshotIdentifier              string
	_rdsDBInstanceAutomatedBackupsArn            string
	_rdsDBInstanceClass                          string
	_rdsDBInstanceCount                          string
	_rdsDBInstanceIdentifier                     string
	_rdsDBInstanceIdentifiers                    []string
	_rdsDBInstanceParameterGroupName             string
	_rdsDBName                                   string
	_rdsDBParameterGroupFamily                   string
	_rdsDBParameterGroupName                     string
	_rdsDBPortNumber                             string
	_rdsDBProxyEndpointName                      string
	_rdsDBProxyName                              string
	_rdsDBSecurityGroupDescription               string
	_rdsDBSecurityGroupName                      string
	_rdsDBSecurityGroups                         []string
	_rdsDBShardGroupIdentifier                   string
	_rdsDBSnapshotIdentifier                     string
	_rdsDBSubnetGroupDescription                 string
	_rdsDBSubnetGroupName                        string
	_rdsDBSystemId                               string
	_rdsDbiResourceId                            string
	_rdsDebugLogging                             string
	_rdsDedicatedLogVolume                       string
	_rdsDefaultAuthScheme                        string
	_rdsDefaultOnly                              string
	_rdsDeleteAutomatedBackups                   string
	_rdsDeleteTarget                             string
	_rdsDeletionProtection                       string
	_rdsDescription                              string
	_rdsDisableDomain                            string
	_rdsDomain                                   string
	_rdsDomainAuthSecretArn                      string
	_rdsDomainDnsIps                             []string
	_rdsDomainFqdn                               string
	_rdsDomainIAMRoleName                        string
	_rdsDomainOu                                 string
	_rdsDuration                                 string
	_rdsEC2SecurityGroupId                       string
	_rdsEC2SecurityGroupName                     string
	_rdsEC2SecurityGroupOwnerId                  string
	_rdsEnableCloudwatchLogsExports              []string
	_rdsEnableCustomerOwnedIp                    string
	_rdsEnableGlobalWriteForwarding              string
	_rdsEnableIAMDatabaseAuthentication          string
	_rdsEnableLimitlessDatabase                  string
	_rdsEnableLocalWriteForwarding               string
	_rdsEnablePerformanceInsights                string
	_rdsEnabled                                  string
	_rdsEndTime                                  string
	_rdsEndpointNetworkType                      string
	_rdsEndpointType                             string
	_rdsEngine                                   string
	_rdsEngineFamily                             string
	_rdsEngineLifecycleSupport                   string
	_rdsEngineMode                               string
	_rdsEngineName                               string
	_rdsEngineNativeAuditFieldsIncluded          string
	_rdsEngineVersion                            string
	_rdsEventCategories                          []string
	_rdsExcludedMembers                          []string
	_rdsExportOnly                               []string
	_rdsExportTaskIdentifier                     string
	_rdsFeatureName                              string
	_rdsFileLastWritten                          string
	_rdsFileSize                                 string
	_rdsFilenameContains                         string
	_rdsFilters                                  string
	_rdsFinalDBSnapshotIdentifier                string
	_rdsForce                                    string
	_rdsForceFailover                            string
	_rdsGlobalClusterIdentifier                  string
	_rdsIamRoleArn                               string
	_rdsIdleClientTimeout                        string
	_rdsImageId                                  string
	_rdsIncludeAll                               string
	_rdsIncludePublic                            string
	_rdsIncludeShared                            string
	_rdsIntegrationIdentifier                    string
	_rdsIntegrationName                          string
	_rdsIops                                     string
	_rdsKmsKeyId                                 string
	_rdsLastUpdatedAfter                         string
	_rdsLastUpdatedBefore                        string
	_rdsLeaseId                                  string
	_rdsLicenseModel                             string
	_rdsListSupportedCharacterSets               string
	_rdsListSupportedTimezones                   string
	_rdsLocale                                   string
	_rdsLogFileName                              string
	_rdsMajorEngineVersion                       string
	_rdsManageMasterUserPassword                 string
	_rdsManifest                                 string
	_rdsMarker                                   string
	_rdsMasterUserAuthenticationType             string
	_rdsMasterUserPassword                       string
	_rdsMasterUserSecretKmsKeyId                 string
	_rdsMasterUsername                           string
	_rdsMaxACU                                   string
	_rdsMaxAllocatedStorage                      string
	_rdsMaxRecords                               string
	_rdsMinACU                                   string
	_rdsMode                                     string
	_rdsMonitoringInterval                       string
	_rdsMonitoringRoleArn                        string
	_rdsMultiAZ                                  string
	_rdsMultiTenant                              string
	_rdsNcharCharacterSetName                    string
	_rdsNetworkType                              string
	_rdsNewDBClusterIdentifier                   string
	_rdsNewDBInstanceIdentifier                  string
	_rdsNewDBProxyEndpointName                   string
	_rdsNewDBProxyName                           string
	_rdsNewGlobalClusterIdentifier               string
	_rdsNewName                                  string
	_rdsNewTenantDBName                          string
	_rdsNumberOfLines                            string
	_rdsOfferingType                             string
	_rdsOptInType                                string
	_rdsOptionGroupDescription                   string
	_rdsOptionGroupName                          string
	_rdsOptionsToInclude                         string
	_rdsOptionsToRemove                          []string
	_rdsParameters                               string
	_rdsPerformanceInsightsKMSKeyId              string
	_rdsPerformanceInsightsRetentionPeriod       string
	_rdsPort                                     string
	_rdsPreSignedUrl                             string
	_rdsPreferredBackupWindow                    string
	_rdsPreferredMaintenanceWindow               string
	_rdsProcessorFeatures                        string
	_rdsProductDescription                       string
	_rdsPromotionTier                            string
	_rdsPubliclyAccessible                       string
	_rdsRdsCustomClusterConfiguration            string
	_rdsRecommendationId                         string
	_rdsRecommendedActionUpdates                 string
	_rdsRegionName                               string
	_rdsRemoveCustomerOverride                   string
	_rdsReplicaMode                              string
	_rdsReplicationSourceIdentifier              string
	_rdsRequireTLS                               string
	_rdsReservedDBInstanceId                     string
	_rdsReservedDBInstancesOfferingId            string
	_rdsResetAllParameters                       string
	_rdsResourceArn                              string
	_rdsResourceIdentifier                       string
	_rdsResourceName                             string
	_rdsRestoreTime                              string
	_rdsRestoreToTime                            string
	_rdsRestoreType                              string
	_rdsResumeFullAutomationModeMinutes          string
	_rdsRoleArn                                  string
	_rdsRotateMasterUserPassword                 string
	_rdsS3BucketName                             string
	_rdsS3IngestionRoleArn                       string
	_rdsS3Prefix                                 string
	_rdsScalingConfiguration                     string
	_rdsSecondsBeforeTimeout                     string
	_rdsSecurityGroups                           []string
	_rdsServerlessV2ScalingConfiguration         string
	_rdsSkipFinalSnapshot                        string
	_rdsSnapshotAvailabilityZone                 string
	_rdsSnapshotIdentifier                       string
	_rdsSnapshotTarget                           string
	_rdsSnapshotType                             string
	_rdsSnsTopicArn                              string
	_rdsSource                                   string
	_rdsSourceArn                                string
	_rdsSourceCustomDbEngineVersionIdentifier    string
	_rdsSourceDBClusterIdentifier                string
	_rdsSourceDBClusterParameterGroupIdentifier  string
	_rdsSourceDbClusterResourceId                string
	_rdsSourceDBClusterSnapshotIdentifier        string
	_rdsSourceDBInstanceArn                      string
	_rdsSourceDBInstanceAutomatedBackupsArn      string
	_rdsSourceDBInstanceIdentifier               string
	_rdsSourceDBParameterGroupIdentifier         string
	_rdsSourceDBSnapshotIdentifier               string
	_rdsSourceDbiResourceId                      string
	_rdsSourceEngine                             string
	_rdsSourceEngineVersion                      string
	_rdsSourceIdentifier                         string
	_rdsSourceIds                                []string
	_rdsSourceOptionGroupIdentifier              string
	_rdsSourceRegion                             string
	_rdsSourceType                               string
	_rdsStartTime                                string
	_rdsStaticMembers                            []string
	_rdsStatus                                   string
	_rdsStorageEncrypted                         string
	_rdsStorageThroughput                        string
	_rdsStorageType                              string
	_rdsSubnetIds                                []string
	_rdsSubscriptionName                         string
	_rdsSwitchover                               string
	_rdsSwitchoverTimeout                        string
	_rdsTagKeys                                  []string
	_rdsTagSpecifications                        string
	_rdsTags                                     string
	_rdsTargetAllocatedStorage                   string
	_rdsTargetArn                                string
	_rdsTargetConnectionNetworkType              string
	_rdsTargetCustomAvailabilityZone             string
	_rdsTargetDbClusterIdentifier                string
	_rdsTargetDBClusterParameterGroupDescription string
	_rdsTargetDBClusterParameterGroupIdentifier  string
	_rdsTargetDBClusterParameterGroupName        string
	_rdsTargetDBClusterSnapshotIdentifier        string
	_rdsTargetDBInstanceClass                    string
	_rdsTargetDBInstanceIdentifier               string
	_rdsTargetDBParameterGroupDescription        string
	_rdsTargetDBParameterGroupIdentifier         string
	_rdsTargetDBParameterGroupName               string
	_rdsTargetDBSnapshotIdentifier               string
	_rdsTargetEngineVersion                      string
	_rdsTargetGroupName                          string
	_rdsTargetIops                               string
	_rdsTargetOptionGroupDescription             string
	_rdsTargetOptionGroupIdentifier              string
	_rdsTargetRole                               string
	_rdsTargetStorageThroughput                  string
	_rdsTargetStorageType                        string
	_rdsTdeCredentialArn                         string
	_rdsTdeCredentialPassword                    string
	_rdsTenantDBName                             string
	_rdsTimeoutAction                            string
	_rdsTimezone                                 string
	_rdsUpgradeStorageConfig                     string
	_rdsUpgradeTargetStorageConfig               string
	_rdsUseAwsProvidedLatestImage                string
	_rdsUseDefaultProcessorFeatures              string
	_rdsUseEarliestTimeOnPointInTimeUnavailable  string
	_rdsUseLatestRestorableTime                  string
	_rdsValuesToAdd                              []string
	_rdsValuesToRemove                           []string
	_rdsVpc                                      string
	_rdsVpcSecurityGroupIds                      []string
	_rdsVpcSubnetIds                             []string
)

// Associates an Identity and Access Management (IAM) role with a DB cluster.
func rds_AddRoleToDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.AddRoleToDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsRoleArn) > 0 {
		input.RoleArn = aws.String(_rdsRoleArn)
	}
	if len(_rdsFeatureName) > 0 {
		input.FeatureName = aws.String(_rdsFeatureName)
	}

	if resp, err := client.AddRoleToDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an Amazon Web Services Identity and Access Management (IAM) role
// with a DB instance.
//
// To add a role to a DB instance, the status of the DB instance must be available .
//
// This command doesn't apply to RDS Custom.
func rds_AddRoleToDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.AddRoleToDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
		// FeatureName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsFeatureName) > 0 {
		input.FeatureName = aws.String(_rdsFeatureName)
	}
	if len(_rdsRoleArn) > 0 {
		input.RoleArn = aws.String(_rdsRoleArn)
	}

	if resp, err := client.AddRoleToDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a source identifier to an existing RDS event notification subscription.
func rds_AddSourceIdentifierToSubscription(cfg aws.Config, client *rds.Client) {
	input := &rds.AddSourceIdentifierToSubscriptionInput{
		// SourceIdentifier: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_rdsSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_rdsSourceIdentifier)
	}
	if len(_rdsSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_rdsSubscriptionName)
	}

	if resp, err := client.AddSourceIdentifierToSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to an Amazon RDS resource. These tags can also be used with
// cost allocation reporting to track cost associated with Amazon RDS resources, or
// used in a Condition statement in an IAM policy for Amazon RDS.
//
// For an overview on tagging your relational database resources, see [Tagging Amazon RDS Resources] or [Tagging Amazon Aurora and Amazon RDS Resources].
//
// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
// [Tagging Amazon Aurora and Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html
func rds_AddTagsToResource(cfg aws.Config, client *rds.Client) {
	input := &rds.AddTagsToResourceInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_rdsResourceName) > 0 {
		input.ResourceName = aws.String(_rdsResourceName)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
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

// Applies a pending maintenance action to a resource (for example, to a DB
// instance).
func rds_ApplyPendingMaintenanceAction(cfg aws.Config, client *rds.Client) {
	input := &rds.ApplyPendingMaintenanceActionInput{
		// ApplyAction: *string, // Required
		// OptInType: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_rdsApplyAction) > 0 {
		input.ApplyAction = aws.String(_rdsApplyAction)
	}
	if len(_rdsOptInType) > 0 {
		input.OptInType = aws.String(_rdsOptInType)
	}
	if len(_rdsResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_rdsResourceIdentifier)
	}

	if resp, err := client.ApplyPendingMaintenanceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables ingress to a DBSecurityGroup using one of two forms of authorization.
// First, EC2 or VPC security groups can be added to the DBSecurityGroup if the
// application using the database is running on EC2 or VPC instances. Second, IP
// ranges are available if the application accessing your database is running on
// the internet. Required parameters for this API are one of CIDR range,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId for non-VPC).
//
// You can't authorize ingress from an EC2 security group in one Amazon Web
// Services Region to an Amazon RDS DB instance in another. You can't authorize
// ingress from a VPC security group in one VPC to an Amazon RDS DB instance in
// another.
//
// For an overview of CIDR ranges, go to the [Wikipedia Tutorial].
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [Wikipedia Tutorial]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func rds_AuthorizeDBSecurityGroupIngress(cfg aws.Config, client *rds.Client) {
	input := &rds.AuthorizeDBSecurityGroupIngressInput{
		// DBSecurityGroupName: *string, // Required
	}

	if len(_rdsDBSecurityGroupName) > 0 {
		input.DBSecurityGroupName = aws.String(_rdsDBSecurityGroupName)
	}
	if len(_rdsCIDRIP) > 0 {
		input.CIDRIP = aws.String(_rdsCIDRIP)
	}
	if len(_rdsEC2SecurityGroupId) > 0 {
		input.EC2SecurityGroupId = aws.String(_rdsEC2SecurityGroupId)
	}
	if len(_rdsEC2SecurityGroupName) > 0 {
		input.EC2SecurityGroupName = aws.String(_rdsEC2SecurityGroupName)
	}
	if len(_rdsEC2SecurityGroupOwnerId) > 0 {
		input.EC2SecurityGroupOwnerId = aws.String(_rdsEC2SecurityGroupOwnerId)
	}

	if resp, err := client.AuthorizeDBSecurityGroupIngress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Backtracks a DB cluster to a specific time, without creating a new DB cluster.
// For more information on backtracking, see [Backtracking an Aurora DB Cluster] in the Amazon Aurora User Guide.
//
// This action applies only to Aurora MySQL DB clusters.
//
// [Backtracking an Aurora DB Cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html
func rds_BacktrackDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.BacktrackDBClusterInput{
		// BacktrackTo: *time.Time, // Required
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsBacktrackTo) > 0 {
		if err := assignInputField(input, "BacktrackTo", _rdsBacktrackTo); err != nil {
			log.Errorf("invalid --backtrack-to: %s", err.Error())
			return
		}
	}
	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsForce) > 0 {
		if err := assignInputField(input, "Force", _rdsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_rdsUseEarliestTimeOnPointInTimeUnavailable) > 0 {
		if err := assignInputField(input, "UseEarliestTimeOnPointInTimeUnavailable", _rdsUseEarliestTimeOnPointInTimeUnavailable); err != nil {
			log.Errorf("invalid --use-earliest-time-on-point-in-time-unavailable: %s", err.Error())
			return
		}
	}

	if resp, err := client.BacktrackDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an export task in progress that is exporting a snapshot or cluster to
// Amazon S3. Any data that has already been written to the S3 bucket isn't
// removed.
func rds_CancelExportTask(cfg aws.Config, client *rds.Client) {
	input := &rds.CancelExportTaskInput{
		// ExportTaskIdentifier: *string, // Required
	}

	if len(_rdsExportTaskIdentifier) > 0 {
		input.ExportTaskIdentifier = aws.String(_rdsExportTaskIdentifier)
	}

	if resp, err := client.CancelExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified DB cluster parameter group.
// You can't copy a default DB cluster parameter group. Instead, create a new
// custom DB cluster parameter group, which copies the default parameters and
// values for the specified DB cluster parameter group family.
func rds_CopyDBClusterParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CopyDBClusterParameterGroupInput{
		// SourceDBClusterParameterGroupIdentifier: *string, // Required
		// TargetDBClusterParameterGroupDescription: *string, // Required
		// TargetDBClusterParameterGroupIdentifier: *string, // Required
	}

	if len(_rdsSourceDBClusterParameterGroupIdentifier) > 0 {
		input.SourceDBClusterParameterGroupIdentifier = aws.String(_rdsSourceDBClusterParameterGroupIdentifier)
	}
	if len(_rdsTargetDBClusterParameterGroupDescription) > 0 {
		input.TargetDBClusterParameterGroupDescription = aws.String(_rdsTargetDBClusterParameterGroupDescription)
	}
	if len(_rdsTargetDBClusterParameterGroupIdentifier) > 0 {
		input.TargetDBClusterParameterGroupIdentifier = aws.String(_rdsTargetDBClusterParameterGroupIdentifier)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies a snapshot of a DB cluster.
// To copy a DB cluster snapshot from a shared manual DB cluster snapshot,
// SourceDBClusterSnapshotIdentifier must be the Amazon Resource Name (ARN) of the
// shared DB cluster snapshot.
//
// You can copy an encrypted DB cluster snapshot from another Amazon Web Services
// Region. In that case, the Amazon Web Services Region where you call the
// CopyDBClusterSnapshot operation is the destination Amazon Web Services Region
// for the encrypted DB cluster snapshot to be copied to. To copy an encrypted DB
// cluster snapshot from another Amazon Web Services Region, you must provide the
// following values:
//
// - KmsKeyId - The Amazon Web Services Key Management System (Amazon Web
// Services KMS) key identifier for the key to use to encrypt the copy of the DB
// cluster snapshot in the destination Amazon Web Services Region.
//
// - TargetDBClusterSnapshotIdentifier - The identifier for the new copy of the
// DB cluster snapshot in the destination Amazon Web Services Region.
//
// - SourceDBClusterSnapshotIdentifier - The DB cluster snapshot identifier for
// the encrypted DB cluster snapshot to be copied. This identifier must be in the
// ARN format for the source Amazon Web Services Region and is the same value as
// the SourceDBClusterSnapshotIdentifier in the presigned URL.
//
// To cancel the copy operation once it is in progress, delete the target DB
// cluster snapshot identified by TargetDBClusterSnapshotIdentifier while that DB
// cluster snapshot is in "copying" status.
//
// For more information on copying encrypted Amazon Aurora DB cluster snapshots
// from one Amazon Web Services Region to another, see [Copying a Snapshot]in the Amazon Aurora User
// Guide.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [Copying a Snapshot]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_CopyDBClusterSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.CopyDBClusterSnapshotInput{
		// SourceDBClusterSnapshotIdentifier: *string, // Required
		// TargetDBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_rdsSourceDBClusterSnapshotIdentifier) > 0 {
		input.SourceDBClusterSnapshotIdentifier = aws.String(_rdsSourceDBClusterSnapshotIdentifier)
	}
	if len(_rdsTargetDBClusterSnapshotIdentifier) > 0 {
		input.TargetDBClusterSnapshotIdentifier = aws.String(_rdsTargetDBClusterSnapshotIdentifier)
	}
	if len(_rdsCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _rdsCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_rdsPreSignedUrl)
	}
	if len(_rdsSourceRegion) > 0 {
		input.SourceRegion = aws.String(_rdsSourceRegion)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified DB parameter group.
// You can't copy a default DB parameter group. Instead, create a new custom DB
// parameter group, which copies the default parameters and values for the
// specified DB parameter group family.
func rds_CopyDBParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CopyDBParameterGroupInput{
		// SourceDBParameterGroupIdentifier: *string, // Required
		// TargetDBParameterGroupDescription: *string, // Required
		// TargetDBParameterGroupIdentifier: *string, // Required
	}

	if len(_rdsSourceDBParameterGroupIdentifier) > 0 {
		input.SourceDBParameterGroupIdentifier = aws.String(_rdsSourceDBParameterGroupIdentifier)
	}
	if len(_rdsTargetDBParameterGroupDescription) > 0 {
		input.TargetDBParameterGroupDescription = aws.String(_rdsTargetDBParameterGroupDescription)
	}
	if len(_rdsTargetDBParameterGroupIdentifier) > 0 {
		input.TargetDBParameterGroupIdentifier = aws.String(_rdsTargetDBParameterGroupIdentifier)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified DB snapshot. The source DB snapshot must be in the
// available state.
//
// You can copy a snapshot from one Amazon Web Services Region to another. In that
// case, the Amazon Web Services Region where you call the CopyDBSnapshot
// operation is the destination Amazon Web Services Region for the DB snapshot
// copy.
//
// This command doesn't apply to RDS Custom.
//
// For more information about copying snapshots, see [Copying a DB Snapshot] in the Amazon RDS User Guide.
//
// [Copying a DB Snapshot]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopySnapshot.html#USER_CopyDBSnapshot
func rds_CopyDBSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.CopyDBSnapshotInput{
		// SourceDBSnapshotIdentifier: *string, // Required
		// TargetDBSnapshotIdentifier: *string, // Required
	}

	if len(_rdsSourceDBSnapshotIdentifier) > 0 {
		input.SourceDBSnapshotIdentifier = aws.String(_rdsSourceDBSnapshotIdentifier)
	}
	if len(_rdsTargetDBSnapshotIdentifier) > 0 {
		input.TargetDBSnapshotIdentifier = aws.String(_rdsTargetDBSnapshotIdentifier)
	}
	if len(_rdsCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _rdsCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_rdsPreSignedUrl)
	}
	if len(_rdsSnapshotAvailabilityZone) > 0 {
		input.SnapshotAvailabilityZone = aws.String(_rdsSnapshotAvailabilityZone)
	}
	if len(_rdsSnapshotTarget) > 0 {
		input.SnapshotTarget = aws.String(_rdsSnapshotTarget)
	}
	if len(_rdsSourceRegion) > 0 {
		input.SourceRegion = aws.String(_rdsSourceRegion)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetCustomAvailabilityZone) > 0 {
		input.TargetCustomAvailabilityZone = aws.String(_rdsTargetCustomAvailabilityZone)
	}

	if resp, err := client.CopyDBSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified option group.
func rds_CopyOptionGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CopyOptionGroupInput{
		// SourceOptionGroupIdentifier: *string, // Required
		// TargetOptionGroupDescription: *string, // Required
		// TargetOptionGroupIdentifier: *string, // Required
	}

	if len(_rdsSourceOptionGroupIdentifier) > 0 {
		input.SourceOptionGroupIdentifier = aws.String(_rdsSourceOptionGroupIdentifier)
	}
	if len(_rdsTargetOptionGroupDescription) > 0 {
		input.TargetOptionGroupDescription = aws.String(_rdsTargetOptionGroupDescription)
	}
	if len(_rdsTargetOptionGroupIdentifier) > 0 {
		input.TargetOptionGroupIdentifier = aws.String(_rdsTargetOptionGroupIdentifier)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyOptionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a blue/green deployment.
// A blue/green deployment creates a staging environment that copies the
// production environment. In a blue/green deployment, the blue environment is the
// current production environment. The green environment is the staging
// environment, and it stays in sync with the current production environment.
//
// You can make changes to the databases in the green environment without
// affecting production workloads. For example, you can upgrade the major or minor
// DB engine version, change database parameters, or make schema changes in the
// staging environment. You can thoroughly test changes in the green environment.
// When ready, you can switch over the environments to promote the green
// environment to be the new production environment. The switchover typically takes
// under a minute.
//
// For more information, see [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon RDS User Guide and [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon
// Aurora User Guide.
//
// [Using Amazon RDS Blue/Green Deployments for database updates]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html
func rds_CreateBlueGreenDeployment(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateBlueGreenDeploymentInput{
		// BlueGreenDeploymentName: *string, // Required
		// Source: *string, // Required
	}

	if len(_rdsBlueGreenDeploymentName) > 0 {
		input.BlueGreenDeploymentName = aws.String(_rdsBlueGreenDeploymentName)
	}
	if len(_rdsSource) > 0 {
		input.Source = aws.String(_rdsSource)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetAllocatedStorage) > 0 {
		if err := assignInputField(input, "TargetAllocatedStorage", _rdsTargetAllocatedStorage); err != nil {
			log.Errorf("invalid --target-allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetDBClusterParameterGroupName) > 0 {
		input.TargetDBClusterParameterGroupName = aws.String(_rdsTargetDBClusterParameterGroupName)
	}
	if len(_rdsTargetDBInstanceClass) > 0 {
		input.TargetDBInstanceClass = aws.String(_rdsTargetDBInstanceClass)
	}
	if len(_rdsTargetDBParameterGroupName) > 0 {
		input.TargetDBParameterGroupName = aws.String(_rdsTargetDBParameterGroupName)
	}
	if len(_rdsTargetEngineVersion) > 0 {
		input.TargetEngineVersion = aws.String(_rdsTargetEngineVersion)
	}
	if len(_rdsTargetIops) > 0 {
		if err := assignInputField(input, "TargetIops", _rdsTargetIops); err != nil {
			log.Errorf("invalid --target-iops: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetStorageThroughput) > 0 {
		if err := assignInputField(input, "TargetStorageThroughput", _rdsTargetStorageThroughput); err != nil {
			log.Errorf("invalid --target-storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetStorageType) > 0 {
		input.TargetStorageType = aws.String(_rdsTargetStorageType)
	}
	if len(_rdsUpgradeTargetStorageConfig) > 0 {
		if err := assignInputField(input, "UpgradeTargetStorageConfig", _rdsUpgradeTargetStorageConfig); err != nil {
			log.Errorf("invalid --upgrade-target-storage-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBlueGreenDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom DB engine version (CEV).
func rds_CreateCustomDBEngineVersion(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateCustomDBEngineVersionInput{
		// Engine: *string, // Required
		// EngineVersion: *string, // Required
	}

	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsDatabaseInstallationFiles) > 0 {
		input.DatabaseInstallationFiles = append([]string(nil), _rdsDatabaseInstallationFiles...)
	}
	if len(_rdsDatabaseInstallationFilesS3BucketName) > 0 {
		input.DatabaseInstallationFilesS3BucketName = aws.String(_rdsDatabaseInstallationFilesS3BucketName)
	}
	if len(_rdsDatabaseInstallationFilesS3Prefix) > 0 {
		input.DatabaseInstallationFilesS3Prefix = aws.String(_rdsDatabaseInstallationFilesS3Prefix)
	}
	if len(_rdsDescription) > 0 {
		input.Description = aws.String(_rdsDescription)
	}
	if len(_rdsImageId) > 0 {
		input.ImageId = aws.String(_rdsImageId)
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KMSKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsManifest) > 0 {
		input.Manifest = aws.String(_rdsManifest)
	}
	if len(_rdsSourceCustomDbEngineVersionIdentifier) > 0 {
		input.SourceCustomDbEngineVersionIdentifier = aws.String(_rdsSourceCustomDbEngineVersionIdentifier)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsUseAwsProvidedLatestImage) > 0 {
		if err := assignInputField(input, "UseAwsProvidedLatestImage", _rdsUseAwsProvidedLatestImage); err != nil {
			log.Errorf("invalid --use-aws-provided-latest-image: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomDBEngineVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Aurora DB cluster or Multi-AZ DB cluster.
// If you create an Aurora DB cluster, the request creates an empty cluster. You
// must explicitly create the writer instance for your DB cluster using the [CreateDBInstance]
// operation. If you create a Multi-AZ DB cluster, the request creates a writer and
// two reader DB instances for you, each in a different Availability Zone.
//
// You can use the ReplicationSourceIdentifier parameter to create an Amazon
// Aurora DB cluster as a read replica of another DB cluster or Amazon RDS for
// MySQL or PostgreSQL DB instance. For more information about Amazon Aurora, see [What is Amazon Aurora?]
// in the Amazon Aurora User Guide.
//
// You can also use the ReplicationSourceIdentifier parameter to create a Multi-AZ
// DB cluster read replica with an RDS for MySQL or PostgreSQL DB instance as the
// source. For more information about Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments]in the Amazon RDS
// User Guide.
//
// [CreateDBInstance]: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_CreateDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _rdsAvailabilityZones...)
	}
	if len(_rdsBacktrackWindow) > 0 {
		if err := assignInputField(input, "BacktrackWindow", _rdsBacktrackWindow); err != nil {
			log.Errorf("invalid --backtrack-window: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCharacterSetName) > 0 {
		input.CharacterSetName = aws.String(_rdsCharacterSetName)
	}
	if len(_rdsClusterScalabilityType) > 0 {
		if err := assignInputField(input, "ClusterScalabilityType", _rdsClusterScalabilityType); err != nil {
			log.Errorf("invalid --cluster-scalability-type: %s", err.Error())
			return
		}
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBClusterInstanceClass) > 0 {
		input.DBClusterInstanceClass = aws.String(_rdsDBClusterInstanceClass)
	}
	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDBSystemId) > 0 {
		input.DBSystemId = aws.String(_rdsDBSystemId)
	}
	if len(_rdsDatabaseInsightsMode) > 0 {
		if err := assignInputField(input, "DatabaseInsightsMode", _rdsDatabaseInsightsMode); err != nil {
			log.Errorf("invalid --database-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsDatabaseName) > 0 {
		input.DatabaseName = aws.String(_rdsDatabaseName)
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableGlobalWriteForwarding) > 0 {
		if err := assignInputField(input, "EnableGlobalWriteForwarding", _rdsEnableGlobalWriteForwarding); err != nil {
			log.Errorf("invalid --enable-global-write-forwarding: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableLimitlessDatabase) > 0 {
		if err := assignInputField(input, "EnableLimitlessDatabase", _rdsEnableLimitlessDatabase); err != nil {
			log.Errorf("invalid --enable-limitless-database: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableLocalWriteForwarding) > 0 {
		if err := assignInputField(input, "EnableLocalWriteForwarding", _rdsEnableLocalWriteForwarding); err != nil {
			log.Errorf("invalid --enable-local-write-forwarding: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineMode) > 0 {
		input.EngineMode = aws.String(_rdsEngineMode)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserAuthenticationType) > 0 {
		if err := assignInputField(input, "MasterUserAuthenticationType", _rdsMasterUserAuthenticationType); err != nil {
			log.Errorf("invalid --master-user-authentication-type: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMasterUsername) > 0 {
		input.MasterUsername = aws.String(_rdsMasterUsername)
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_rdsPreSignedUrl)
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_rdsPreferredMaintenanceWindow)
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsRdsCustomClusterConfiguration) > 0 {
		if err := assignInputField(input, "RdsCustomClusterConfiguration", _rdsRdsCustomClusterConfiguration); err != nil {
			log.Errorf("invalid --rds-custom-cluster-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsReplicationSourceIdentifier) > 0 {
		input.ReplicationSourceIdentifier = aws.String(_rdsReplicationSourceIdentifier)
	}
	if len(_rdsScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _rdsScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _rdsServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceRegion) > 0 {
		input.SourceRegion = aws.String(_rdsSourceRegion)
	}
	if len(_rdsStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _rdsStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom endpoint and associates it with an Amazon Aurora DB
// cluster.
//
// This action applies only to Aurora DB clusters.
func rds_CreateDBClusterEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBClusterEndpointInput{
		// DBClusterEndpointIdentifier: *string, // Required
		// DBClusterIdentifier: *string, // Required
		// EndpointType: *string, // Required
	}

	if len(_rdsDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_rdsDBClusterEndpointIdentifier)
	}
	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsEndpointType) > 0 {
		input.EndpointType = aws.String(_rdsEndpointType)
	}
	if len(_rdsExcludedMembers) > 0 {
		input.ExcludedMembers = append([]string(nil), _rdsExcludedMembers...)
	}
	if len(_rdsStaticMembers) > 0 {
		input.StaticMembers = append([]string(nil), _rdsStaticMembers...)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBClusterEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB cluster parameter group.
// Parameters in a DB cluster parameter group apply to all of the instances in a
// DB cluster.
//
// A DB cluster parameter group is initially created with the default parameters
// for the database engine used by instances in the DB cluster. To provide custom
// values for any of the parameters, you must modify the group after creating it
// using ModifyDBClusterParameterGroup . Once you've created a DB cluster parameter
// group, you need to associate it with your DB cluster using ModifyDBCluster .
//
// When you associate a new DB cluster parameter group with a running Aurora DB
// cluster, reboot the DB instances in the DB cluster without failover for the new
// DB cluster parameter group and associated settings to take effect.
//
// When you associate a new DB cluster parameter group with a running Multi-AZ DB
// cluster, reboot the DB cluster without failover for the new DB cluster parameter
// group and associated settings to take effect.
//
// After you create a DB cluster parameter group, you should wait at least 5
// minutes before creating your first DB cluster that uses that DB cluster
// parameter group as the default parameter group. This allows Amazon RDS to fully
// complete the create action before the DB cluster parameter group is used as the
// default for a new DB cluster. This is especially important for parameters that
// are critical when creating the default database for a DB cluster, such as the
// character set for the default database defined by the character_set_database
// parameter. You can use the Parameter Groups option of the [Amazon RDS console]or the
// DescribeDBClusterParameters operation to verify that your DB cluster parameter
// group has been created or modified.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [Amazon RDS console]: https://console.aws.amazon.com/rds/
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_CreateDBClusterParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
		// DBParameterGroupFamily: *string, // Required
		// Description: *string, // Required
	}

	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_rdsDBParameterGroupFamily)
	}
	if len(_rdsDescription) > 0 {
		input.Description = aws.String(_rdsDescription)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of a DB cluster.
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_CreateDBClusterSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBClusterSnapshotInput{
		// DBClusterIdentifier: *string, // Required
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_rdsDBClusterSnapshotIdentifier)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB instance.
// The new DB instance can be an RDS DB instance, or it can be a DB instance in an
// Aurora DB cluster. For an Aurora DB cluster, you can call this operation
// multiple times to add more than one DB instance to the cluster.
//
// For more information about creating an RDS DB instance, see [Creating an Amazon RDS DB instance] in the Amazon RDS
// User Guide.
//
// For more information about creating a DB instance in an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster]
// in the Amazon Aurora User Guide.
//
// [Creating an Amazon Aurora DB cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html
// [Creating an Amazon RDS DB instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html
func rds_CreateDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBInstanceInput{
		// DBInstanceClass: *string, // Required
		// DBInstanceIdentifier: *string, // Required
		// Engine: *string, // Required
	}

	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsAdditionalStorageVolumes) > 0 {
		if err := assignInputField(input, "AdditionalStorageVolumes", _rdsAdditionalStorageVolumes); err != nil {
			log.Errorf("invalid --additional-storage-volumes: %s", err.Error())
			return
		}
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_rdsAvailabilityZone)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupTarget) > 0 {
		input.BackupTarget = aws.String(_rdsBackupTarget)
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCharacterSetName) > 0 {
		input.CharacterSetName = aws.String(_rdsCharacterSetName)
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsCustomIamInstanceProfile) > 0 {
		input.CustomIamInstanceProfile = aws.String(_rdsCustomIamInstanceProfile)
	}
	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsDBName) > 0 {
		input.DBName = aws.String(_rdsDBName)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDBSecurityGroups) > 0 {
		input.DBSecurityGroups = append([]string(nil), _rdsDBSecurityGroups...)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDBSystemId) > 0 {
		input.DBSystemId = aws.String(_rdsDBSystemId)
	}
	if len(_rdsDatabaseInsightsMode) > 0 {
		if err := assignInputField(input, "DatabaseInsightsMode", _rdsDatabaseInsightsMode); err != nil {
			log.Errorf("invalid --database-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsDedicatedLogVolume) > 0 {
		if err := assignInputField(input, "DedicatedLogVolume", _rdsDedicatedLogVolume); err != nil {
			log.Errorf("invalid --dedicated-log-volume: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainAuthSecretArn) > 0 {
		input.DomainAuthSecretArn = aws.String(_rdsDomainAuthSecretArn)
	}
	if len(_rdsDomainDnsIps) > 0 {
		input.DomainDnsIps = append([]string(nil), _rdsDomainDnsIps...)
	}
	if len(_rdsDomainFqdn) > 0 {
		input.DomainFqdn = aws.String(_rdsDomainFqdn)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsDomainOu) > 0 {
		input.DomainOu = aws.String(_rdsDomainOu)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableCustomerOwnedIp) > 0 {
		if err := assignInputField(input, "EnableCustomerOwnedIp", _rdsEnableCustomerOwnedIp); err != nil {
			log.Errorf("invalid --enable-customer-owned-ip: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsLicenseModel) > 0 {
		input.LicenseModel = aws.String(_rdsLicenseModel)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserAuthenticationType) > 0 {
		if err := assignInputField(input, "MasterUserAuthenticationType", _rdsMasterUserAuthenticationType); err != nil {
			log.Errorf("invalid --master-user-authentication-type: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMasterUsername) > 0 {
		input.MasterUsername = aws.String(_rdsMasterUsername)
	}
	if len(_rdsMaxAllocatedStorage) > 0 {
		if err := assignInputField(input, "MaxAllocatedStorage", _rdsMaxAllocatedStorage); err != nil {
			log.Errorf("invalid --max-allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsMultiTenant) > 0 {
		if err := assignInputField(input, "MultiTenant", _rdsMultiTenant); err != nil {
			log.Errorf("invalid --multi-tenant: %s", err.Error())
			return
		}
	}
	if len(_rdsNcharCharacterSetName) > 0 {
		input.NcharCharacterSetName = aws.String(_rdsNcharCharacterSetName)
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_rdsPreferredMaintenanceWindow)
	}
	if len(_rdsProcessorFeatures) > 0 {
		if err := assignInputField(input, "ProcessorFeatures", _rdsProcessorFeatures); err != nil {
			log.Errorf("invalid --processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsPromotionTier) > 0 {
		if err := assignInputField(input, "PromotionTier", _rdsPromotionTier); err != nil {
			log.Errorf("invalid --promotion-tier: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _rdsStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageThroughput) > 0 {
		if err := assignInputField(input, "StorageThroughput", _rdsStorageThroughput); err != nil {
			log.Errorf("invalid --storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTdeCredentialArn) > 0 {
		input.TdeCredentialArn = aws.String(_rdsTdeCredentialArn)
	}
	if len(_rdsTdeCredentialPassword) > 0 {
		input.TdeCredentialPassword = aws.String(_rdsTdeCredentialPassword)
	}
	if len(_rdsTimezone) > 0 {
		input.Timezone = aws.String(_rdsTimezone)
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB instance that acts as a read replica for an existing source DB
// instance or Multi-AZ DB cluster. You can create a read replica for a DB instance
// running Db2, MariaDB, MySQL, Oracle, PostgreSQL, or SQL Server. You can create a
// read replica for a Multi-AZ DB cluster running MySQL or PostgreSQL. For more
// information, see [Working with read replicas]and [Migrating from a Multi-AZ DB cluster to a DB instance using a read replica] in the Amazon RDS User Guide.
//
// Amazon Aurora doesn't support this operation. To create a DB instance for an
// Aurora DB cluster, use the CreateDBInstance operation.
//
// RDS creates read replicas with backups disabled. All other attributes
// (including DB security groups and DB parameter groups) are inherited from the
// source DB instance or cluster, except as specified.
//
// Your source DB instance or cluster must have backup retention enabled.
//
// [Working with read replicas]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html
// [Migrating from a Multi-AZ DB cluster to a DB instance using a read replica]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html#multi-az-db-clusters-migrating-to-instance-with-read-replica
func rds_CreateDBInstanceReadReplica(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBInstanceReadReplicaInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsAdditionalStorageVolumes) > 0 {
		if err := assignInputField(input, "AdditionalStorageVolumes", _rdsAdditionalStorageVolumes); err != nil {
			log.Errorf("invalid --additional-storage-volumes: %s", err.Error())
			return
		}
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_rdsAvailabilityZone)
	}
	if len(_rdsBackupTarget) > 0 {
		input.BackupTarget = aws.String(_rdsBackupTarget)
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsCustomIamInstanceProfile) > 0 {
		input.CustomIamInstanceProfile = aws.String(_rdsCustomIamInstanceProfile)
	}
	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDatabaseInsightsMode) > 0 {
		if err := assignInputField(input, "DatabaseInsightsMode", _rdsDatabaseInsightsMode); err != nil {
			log.Errorf("invalid --database-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsDedicatedLogVolume) > 0 {
		if err := assignInputField(input, "DedicatedLogVolume", _rdsDedicatedLogVolume); err != nil {
			log.Errorf("invalid --dedicated-log-volume: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainAuthSecretArn) > 0 {
		input.DomainAuthSecretArn = aws.String(_rdsDomainAuthSecretArn)
	}
	if len(_rdsDomainDnsIps) > 0 {
		input.DomainDnsIps = append([]string(nil), _rdsDomainDnsIps...)
	}
	if len(_rdsDomainFqdn) > 0 {
		input.DomainFqdn = aws.String(_rdsDomainFqdn)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsDomainOu) > 0 {
		input.DomainOu = aws.String(_rdsDomainOu)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableCustomerOwnedIp) > 0 {
		if err := assignInputField(input, "EnableCustomerOwnedIp", _rdsEnableCustomerOwnedIp); err != nil {
			log.Errorf("invalid --enable-customer-owned-ip: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsMaxAllocatedStorage) > 0 {
		if err := assignInputField(input, "MaxAllocatedStorage", _rdsMaxAllocatedStorage); err != nil {
			log.Errorf("invalid --max-allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_rdsPreSignedUrl)
	}
	if len(_rdsProcessorFeatures) > 0 {
		if err := assignInputField(input, "ProcessorFeatures", _rdsProcessorFeatures); err != nil {
			log.Errorf("invalid --processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsReplicaMode) > 0 {
		if err := assignInputField(input, "ReplicaMode", _rdsReplicaMode); err != nil {
			log.Errorf("invalid --replica-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_rdsSourceDBClusterIdentifier)
	}
	if len(_rdsSourceDBInstanceIdentifier) > 0 {
		input.SourceDBInstanceIdentifier = aws.String(_rdsSourceDBInstanceIdentifier)
	}
	if len(_rdsSourceRegion) > 0 {
		input.SourceRegion = aws.String(_rdsSourceRegion)
	}
	if len(_rdsStorageThroughput) > 0 {
		if err := assignInputField(input, "StorageThroughput", _rdsStorageThroughput); err != nil {
			log.Errorf("invalid --storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsUpgradeStorageConfig) > 0 {
		if err := assignInputField(input, "UpgradeStorageConfig", _rdsUpgradeStorageConfig); err != nil {
			log.Errorf("invalid --upgrade-storage-config: %s", err.Error())
			return
		}
	}
	if len(_rdsUseDefaultProcessorFeatures) > 0 {
		if err := assignInputField(input, "UseDefaultProcessorFeatures", _rdsUseDefaultProcessorFeatures); err != nil {
			log.Errorf("invalid --use-default-processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBInstanceReadReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB parameter group.
// A DB parameter group is initially created with the default parameters for the
// database engine used by the DB instance. To provide custom values for any of the
// parameters, you must modify the group after creating it using
// ModifyDBParameterGroup . Once you've created a DB parameter group, you need to
// associate it with your DB instance using ModifyDBInstance . When you associate a
// new DB parameter group with a running DB instance, you need to reboot the DB
// instance without failover for the new DB parameter group and associated settings
// to take effect.
//
// This command doesn't apply to RDS Custom.
func rds_CreateDBParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBParameterGroupInput{
		// DBParameterGroupFamily: *string, // Required
		// DBParameterGroupName: *string, // Required
		// Description: *string, // Required
	}

	if len(_rdsDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_rdsDBParameterGroupFamily)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDescription) > 0 {
		input.Description = aws.String(_rdsDescription)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB proxy.
func rds_CreateDBProxy(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBProxyInput{
		// DBProxyName: *string, // Required
		// EngineFamily: types.EngineFamily, // Required
		// RoleArn: *string, // Required
		// VpcSubnetIds: []string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsEngineFamily) > 0 {
		if err := assignInputField(input, "EngineFamily", _rdsEngineFamily); err != nil {
			log.Errorf("invalid --engine-family: %s", err.Error())
			return
		}
	}
	if len(_rdsRoleArn) > 0 {
		input.RoleArn = aws.String(_rdsRoleArn)
	}
	if len(_rdsVpcSubnetIds) > 0 {
		input.VpcSubnetIds = append([]string(nil), _rdsVpcSubnetIds...)
	}
	if len(_rdsAuth) > 0 {
		if err := assignInputField(input, "Auth", _rdsAuth); err != nil {
			log.Errorf("invalid --auth: %s", err.Error())
			return
		}
	}
	if len(_rdsDebugLogging) > 0 {
		if err := assignInputField(input, "DebugLogging", _rdsDebugLogging); err != nil {
			log.Errorf("invalid --debug-logging: %s", err.Error())
			return
		}
	}
	if len(_rdsDefaultAuthScheme) > 0 {
		if err := assignInputField(input, "DefaultAuthScheme", _rdsDefaultAuthScheme); err != nil {
			log.Errorf("invalid --default-auth-scheme: %s", err.Error())
			return
		}
	}
	if len(_rdsEndpointNetworkType) > 0 {
		if err := assignInputField(input, "EndpointNetworkType", _rdsEndpointNetworkType); err != nil {
			log.Errorf("invalid --endpoint-network-type: %s", err.Error())
			return
		}
	}
	if len(_rdsIdleClientTimeout) > 0 {
		if err := assignInputField(input, "IdleClientTimeout", _rdsIdleClientTimeout); err != nil {
			log.Errorf("invalid --idle-client-timeout: %s", err.Error())
			return
		}
	}
	if len(_rdsRequireTLS) > 0 {
		if err := assignInputField(input, "RequireTLS", _rdsRequireTLS); err != nil {
			log.Errorf("invalid --require-tls: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetConnectionNetworkType) > 0 {
		if err := assignInputField(input, "TargetConnectionNetworkType", _rdsTargetConnectionNetworkType); err != nil {
			log.Errorf("invalid --target-connection-network-type: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DBProxyEndpoint . Only applies to proxies that are associated with
// Aurora DB clusters. You can use DB proxy endpoints to specify read/write or
// read-only access to the DB cluster. You can also use DB proxy endpoints to
// access a DB proxy through a different VPC than the proxy's default VPC.
func rds_CreateDBProxyEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBProxyEndpointInput{
		// DBProxyEndpointName: *string, // Required
		// DBProxyName: *string, // Required
		// VpcSubnetIds: []string, // Required
	}

	if len(_rdsDBProxyEndpointName) > 0 {
		input.DBProxyEndpointName = aws.String(_rdsDBProxyEndpointName)
	}
	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsVpcSubnetIds) > 0 {
		input.VpcSubnetIds = append([]string(nil), _rdsVpcSubnetIds...)
	}
	if len(_rdsEndpointNetworkType) > 0 {
		if err := assignInputField(input, "EndpointNetworkType", _rdsEndpointNetworkType); err != nil {
			log.Errorf("invalid --endpoint-network-type: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetRole) > 0 {
		if err := assignInputField(input, "TargetRole", _rdsTargetRole); err != nil {
			log.Errorf("invalid --target-role: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBProxyEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB security group. DB security groups control access to a DB
// instance.
//
// A DB security group controls access to EC2-Classic DB instances that are not in
// a VPC.
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func rds_CreateDBSecurityGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBSecurityGroupInput{
		// DBSecurityGroupDescription: *string, // Required
		// DBSecurityGroupName: *string, // Required
	}

	if len(_rdsDBSecurityGroupDescription) > 0 {
		input.DBSecurityGroupDescription = aws.String(_rdsDBSecurityGroupDescription)
	}
	if len(_rdsDBSecurityGroupName) > 0 {
		input.DBSecurityGroupName = aws.String(_rdsDBSecurityGroupName)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB shard group for Aurora Limitless Database. You must enable
// Aurora Limitless Database to create a DB shard group.
//
// Valid for: Aurora DB clusters only
func rds_CreateDBShardGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBShardGroupInput{
		// DBClusterIdentifier: *string, // Required
		// DBShardGroupIdentifier: *string, // Required
		// MaxACU: *float64, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsDBShardGroupIdentifier) > 0 {
		input.DBShardGroupIdentifier = aws.String(_rdsDBShardGroupIdentifier)
	}
	if len(_rdsMaxACU) > 0 {
		if err := assignInputField(input, "MaxACU", _rdsMaxACU); err != nil {
			log.Errorf("invalid --max-acu: %s", err.Error())
			return
		}
	}
	if len(_rdsComputeRedundancy) > 0 {
		if err := assignInputField(input, "ComputeRedundancy", _rdsComputeRedundancy); err != nil {
			log.Errorf("invalid --compute-redundancy: %s", err.Error())
			return
		}
	}
	if len(_rdsMinACU) > 0 {
		if err := assignInputField(input, "MinACU", _rdsMinACU); err != nil {
			log.Errorf("invalid --min-acu: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBShardGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of a DB instance. The source DB instance must be in the
// available or storage-optimization state.
func rds_CreateDBSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBSnapshotInput{
		// DBInstanceIdentifier: *string, // Required
		// DBSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB subnet group. DB subnet groups must contain at least one
// subnet in at least two AZs in the Amazon Web Services Region.
func rds_CreateDBSubnetGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateDBSubnetGroupInput{
		// DBSubnetGroupDescription: *string, // Required
		// DBSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_rdsDBSubnetGroupDescription) > 0 {
		input.DBSubnetGroupDescription = aws.String(_rdsDBSubnetGroupDescription)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _rdsSubnetIds...)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an RDS event notification subscription. This operation requires a topic
// Amazon Resource Name (ARN) created by either the RDS console, the SNS console,
// or the SNS API. To obtain an ARN with SNS, you must create a topic in Amazon SNS
// and subscribe to the topic. The ARN is displayed in the SNS console.
//
// You can specify the type of source ( SourceType ) that you want to be notified
// of and provide a list of RDS sources ( SourceIds ) that triggers the events. You
// can also provide a list of event categories ( EventCategories ) for events that
// you want to be notified of. For example, you can specify SourceType =
// db-instance , SourceIds = mydbinstance1 , mydbinstance2 and EventCategories =
// Availability , Backup .
//
// If you specify both the SourceType and SourceIds , such as SourceType =
// db-instance and SourceIds = myDBInstance1 , you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify SourceIds , you receive notice of the events for that source type
// for all your RDS sources. If you don't specify either the SourceType or the
// SourceIds , you are notified of events generated from all RDS sources belonging
// to your customer account.
//
// For more information about subscribing to an event for RDS DB engines, see [Subscribing to Amazon RDS event notification] in
// the Amazon RDS User Guide.
//
// For more information about subscribing to an event for Aurora DB engines, see [Subscribing to Amazon RDS event notification]
// in the Amazon Aurora User Guide.
//
// [Subscribing to Amazon RDS event notification]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Subscribing.html
func rds_CreateEventSubscription(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateEventSubscriptionInput{
		// SnsTopicArn: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_rdsSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_rdsSnsTopicArn)
	}
	if len(_rdsSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_rdsSubscriptionName)
	}
	if len(_rdsEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _rdsEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_rdsEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _rdsEventCategories...)
	}
	if len(_rdsSourceIds) > 0 {
		input.SourceIds = append([]string(nil), _rdsSourceIds...)
	}
	if len(_rdsSourceType) > 0 {
		input.SourceType = aws.String(_rdsSourceType)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Aurora global database spread across multiple Amazon Web Services
// Regions. The global database contains a single primary cluster with read-write
// capability, and a read-only secondary cluster that receives data from the
// primary cluster through high-speed replication performed by the Aurora storage
// subsystem.
//
// You can create a global database that is initially empty, and then create the
// primary and secondary DB clusters in the global database. Or you can specify an
// existing Aurora cluster during the create operation, and this cluster becomes
// the primary cluster of the global database.
//
// This operation applies only to Aurora DB clusters.
func rds_CreateGlobalCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}
	if len(_rdsDatabaseName) > 0 {
		input.DatabaseName = aws.String(_rdsDatabaseName)
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_rdsSourceDBClusterIdentifier)
	}
	if len(_rdsStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _rdsStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a zero-ETL integration with Amazon Redshift.
func rds_CreateIntegration(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateIntegrationInput{
		// IntegrationName: *string, // Required
		// SourceArn: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_rdsIntegrationName) > 0 {
		input.IntegrationName = aws.String(_rdsIntegrationName)
	}
	if len(_rdsSourceArn) > 0 {
		input.SourceArn = aws.String(_rdsSourceArn)
	}
	if len(_rdsTargetArn) > 0 {
		input.TargetArn = aws.String(_rdsTargetArn)
	}
	if len(_rdsAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _rdsAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_rdsDataFilter) > 0 {
		input.DataFilter = aws.String(_rdsDataFilter)
	}
	if len(_rdsDescription) > 0 {
		input.Description = aws.String(_rdsDescription)
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KMSKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new option group. You can create up to 20 option groups.
// This command doesn't apply to RDS Custom.
func rds_CreateOptionGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateOptionGroupInput{
		// EngineName: *string, // Required
		// MajorEngineVersion: *string, // Required
		// OptionGroupDescription: *string, // Required
		// OptionGroupName: *string, // Required
	}

	if len(_rdsEngineName) > 0 {
		input.EngineName = aws.String(_rdsEngineName)
	}
	if len(_rdsMajorEngineVersion) > 0 {
		input.MajorEngineVersion = aws.String(_rdsMajorEngineVersion)
	}
	if len(_rdsOptionGroupDescription) > 0 {
		input.OptionGroupDescription = aws.String(_rdsOptionGroupDescription)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOptionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a tenant database in a DB instance that uses the multi-tenant
// configuration. Only RDS for Oracle container database (CDB) instances are
// supported.
func rds_CreateTenantDatabase(cfg aws.Config, client *rds.Client) {
	input := &rds.CreateTenantDatabaseInput{
		// DBInstanceIdentifier: *string, // Required
		// MasterUsername: *string, // Required
		// TenantDBName: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsMasterUsername) > 0 {
		input.MasterUsername = aws.String(_rdsMasterUsername)
	}
	if len(_rdsTenantDBName) > 0 {
		input.TenantDBName = aws.String(_rdsTenantDBName)
	}
	if len(_rdsCharacterSetName) > 0 {
		input.CharacterSetName = aws.String(_rdsCharacterSetName)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsNcharCharacterSetName) > 0 {
		input.NcharCharacterSetName = aws.String(_rdsNcharCharacterSetName)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTenantDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a blue/green deployment.
// For more information, see [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon RDS User Guide and [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon
// Aurora User Guide.
//
// [Using Amazon RDS Blue/Green Deployments for database updates]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html
func rds_DeleteBlueGreenDeployment(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteBlueGreenDeploymentInput{
		// BlueGreenDeploymentIdentifier: *string, // Required
	}

	if len(_rdsBlueGreenDeploymentIdentifier) > 0 {
		input.BlueGreenDeploymentIdentifier = aws.String(_rdsBlueGreenDeploymentIdentifier)
	}
	if len(_rdsDeleteTarget) > 0 {
		if err := assignInputField(input, "DeleteTarget", _rdsDeleteTarget); err != nil {
			log.Errorf("invalid --delete-target: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBlueGreenDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom engine version. To run this command, make sure you meet the
// following prerequisites:
//
// - The CEV must not be the default for RDS Custom. If it is, change the
// default before running this command.
//
// - The CEV must not be associated with an RDS Custom DB instance, RDS Custom
// instance snapshot, or automated backup of your RDS Custom instance.
//
// Typically, deletion takes a few minutes.
//
// The MediaImport service that imports files from Amazon S3 to create CEVs isn't
// integrated with Amazon Web Services CloudTrail. If you turn on data logging for
// Amazon RDS in CloudTrail, calls to the DeleteCustomDbEngineVersion event aren't
// logged. However, you might see calls from the API gateway that accesses your
// Amazon S3 bucket. These calls originate from the MediaImport service for the
// DeleteCustomDbEngineVersion event.
//
// For more information, see [Deleting a CEV] in the Amazon RDS User Guide.
//
// [Deleting a CEV]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.delete
func rds_DeleteCustomDBEngineVersion(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteCustomDBEngineVersionInput{
		// Engine: *string, // Required
		// EngineVersion: *string, // Required
	}

	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}

	if resp, err := client.DeleteCustomDBEngineVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted.
//
// If you're deleting a Multi-AZ DB cluster with read replicas, all cluster
// members are terminated and read replicas are promoted to standalone instances.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DeleteDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsDeleteAutomatedBackups) > 0 {
		if err := assignInputField(input, "DeleteAutomatedBackups", _rdsDeleteAutomatedBackups); err != nil {
			log.Errorf("invalid --delete-automated-backups: %s", err.Error())
			return
		}
	}
	if len(_rdsFinalDBSnapshotIdentifier) > 0 {
		input.FinalDBSnapshotIdentifier = aws.String(_rdsFinalDBSnapshotIdentifier)
	}
	if len(_rdsSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _rdsSkipFinalSnapshot); err != nil {
			log.Errorf("invalid --skip-final-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes automated backups using the DbClusterResourceId value of the source DB
// cluster or the Amazon Resource Name (ARN) of the automated backups.
func rds_DeleteDBClusterAutomatedBackup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBClusterAutomatedBackupInput{
		// DbClusterResourceId: *string, // Required
	}

	if len(_rdsDbClusterResourceId) > 0 {
		input.DbClusterResourceId = aws.String(_rdsDbClusterResourceId)
	}

	if resp, err := client.DeleteDBClusterAutomatedBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom endpoint and removes it from an Amazon Aurora DB cluster.
// This action only applies to Aurora DB clusters.
func rds_DeleteDBClusterEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBClusterEndpointInput{
		// DBClusterEndpointIdentifier: *string, // Required
	}

	if len(_rdsDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_rdsDBClusterEndpointIdentifier)
	}

	if resp, err := client.DeleteDBClusterEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified DB cluster parameter group. The DB cluster parameter group
// to be deleted can't be associated with any DB clusters.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DeleteDBClusterParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}

	if resp, err := client.DeleteDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DB cluster snapshot. If the snapshot is being copied, the copy
// operation is terminated.
//
// The DB cluster snapshot must be in the available state to be deleted.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DeleteDBClusterSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBClusterSnapshotInput{
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_rdsDBClusterSnapshotIdentifier)
	}

	if resp, err := client.DeleteDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously provisioned DB instance. When you delete a DB instance,
// all automated backups for that instance are deleted and can't be recovered.
// However, manual DB snapshots of the DB instance aren't deleted.
//
// If you request a final DB snapshot, the status of the Amazon RDS DB instance is
// deleting until the DB snapshot is created. This operation can't be canceled or
// reverted after it begins. To monitor the status of this operation, use
// DescribeDBInstance .
//
// When a DB instance is in a failure state and has a status of failed ,
// incompatible-restore , or incompatible-network , you can only delete it when you
// skip creation of the final snapshot with the SkipFinalSnapshot parameter.
//
// If the specified DB instance is part of an Amazon Aurora DB cluster, you can't
// delete the DB instance if both of the following conditions are true:
//
// - The DB cluster is a read replica of another Amazon Aurora DB cluster.
//
// - The DB instance is the only instance in the DB cluster.
//
// To delete a DB instance in this case, first use the PromoteReadReplicaDBCluster
// operation to promote the DB cluster so that it's no longer a read replica. After
// the promotion completes, use the DeleteDBInstance operation to delete the final
// instance in the DB cluster.
//
// For RDS Custom DB instances, deleting the DB instance permanently deletes the
// EC2 instance and the associated EBS volumes. Make sure that you don't terminate
// or delete these resources before you delete the DB instance. Otherwise, deleting
// the DB instance and creation of the final snapshot might fail.
func rds_DeleteDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsDeleteAutomatedBackups) > 0 {
		if err := assignInputField(input, "DeleteAutomatedBackups", _rdsDeleteAutomatedBackups); err != nil {
			log.Errorf("invalid --delete-automated-backups: %s", err.Error())
			return
		}
	}
	if len(_rdsFinalDBSnapshotIdentifier) > 0 {
		input.FinalDBSnapshotIdentifier = aws.String(_rdsFinalDBSnapshotIdentifier)
	}
	if len(_rdsSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _rdsSkipFinalSnapshot); err != nil {
			log.Errorf("invalid --skip-final-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes automated backups using the DbiResourceId value of the source DB
// instance or the Amazon Resource Name (ARN) of the automated backups.
func rds_DeleteDBInstanceAutomatedBackup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBInstanceAutomatedBackupInput{}

	if len(_rdsDBInstanceAutomatedBackupsArn) > 0 {
		input.DBInstanceAutomatedBackupsArn = aws.String(_rdsDBInstanceAutomatedBackupsArn)
	}
	if len(_rdsDbiResourceId) > 0 {
		input.DbiResourceId = aws.String(_rdsDbiResourceId)
	}

	if resp, err := client.DeleteDBInstanceAutomatedBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified DB parameter group. The DB parameter group to be deleted
// can't be associated with any DB instances.
func rds_DeleteDBParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBParameterGroupInput{
		// DBParameterGroupName: *string, // Required
	}

	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}

	if resp, err := client.DeleteDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing DB proxy.
func rds_DeleteDBProxy(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBProxyInput{
		// DBProxyName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}

	if resp, err := client.DeleteDBProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DBProxyEndpoint . Doing so removes the ability to access the DB proxy
// using the endpoint that you defined. The endpoint that you delete might have
// provided capabilities such as read/write or read-only operations, or using a
// different VPC than the DB proxy's default VPC.
func rds_DeleteDBProxyEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBProxyEndpointInput{
		// DBProxyEndpointName: *string, // Required
	}

	if len(_rdsDBProxyEndpointName) > 0 {
		input.DBProxyEndpointName = aws.String(_rdsDBProxyEndpointName)
	}

	if resp, err := client.DeleteDBProxyEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DB security group.
// The specified DB security group must not be associated with any DB instances.
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func rds_DeleteDBSecurityGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBSecurityGroupInput{
		// DBSecurityGroupName: *string, // Required
	}

	if len(_rdsDBSecurityGroupName) > 0 {
		input.DBSecurityGroupName = aws.String(_rdsDBSecurityGroupName)
	}

	if resp, err := client.DeleteDBSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Aurora Limitless Database DB shard group.
func rds_DeleteDBShardGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBShardGroupInput{
		// DBShardGroupIdentifier: *string, // Required
	}

	if len(_rdsDBShardGroupIdentifier) > 0 {
		input.DBShardGroupIdentifier = aws.String(_rdsDBShardGroupIdentifier)
	}

	if resp, err := client.DeleteDBShardGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DB snapshot. If the snapshot is being copied, the copy operation is
// terminated.
//
// The DB snapshot must be in the available state to be deleted.
func rds_DeleteDBSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBSnapshotInput{
		// DBSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}

	if resp, err := client.DeleteDBSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DB subnet group.
// The specified database subnet group must not be associated with any DB
// instances.
func rds_DeleteDBSubnetGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteDBSubnetGroupInput{
		// DBSubnetGroupName: *string, // Required
	}

	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}

	if resp, err := client.DeleteDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an RDS event notification subscription.
func rds_DeleteEventSubscription(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_rdsSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_rdsSubscriptionName)
	}

	if resp, err := client.DeleteEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a global database cluster. The primary and secondary clusters must
// already be detached or destroyed first.
//
// This action only applies to Aurora DB clusters.
func rds_DeleteGlobalCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}

	if resp, err := client.DeleteGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a zero-ETL integration with Amazon Redshift.
func rds_DeleteIntegration(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteIntegrationInput{
		// IntegrationIdentifier: *string, // Required
	}

	if len(_rdsIntegrationIdentifier) > 0 {
		input.IntegrationIdentifier = aws.String(_rdsIntegrationIdentifier)
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing option group.
func rds_DeleteOptionGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteOptionGroupInput{
		// OptionGroupName: *string, // Required
	}

	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}

	if resp, err := client.DeleteOptionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a tenant database from your DB instance. This command only applies to
// RDS for Oracle container database (CDB) instances.
//
// You can't delete a tenant database when it is the only tenant in the DB
// instance.
func rds_DeleteTenantDatabase(cfg aws.Config, client *rds.Client) {
	input := &rds.DeleteTenantDatabaseInput{
		// DBInstanceIdentifier: *string, // Required
		// TenantDBName: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsTenantDBName) > 0 {
		input.TenantDBName = aws.String(_rdsTenantDBName)
	}
	if len(_rdsFinalDBSnapshotIdentifier) > 0 {
		input.FinalDBSnapshotIdentifier = aws.String(_rdsFinalDBSnapshotIdentifier)
	}
	if len(_rdsSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _rdsSkipFinalSnapshot); err != nil {
			log.Errorf("invalid --skip-final-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTenantDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove the association between one or more DBProxyTarget data structures and a
// DBProxyTargetGroup .
func rds_DeregisterDBProxyTargets(cfg aws.Config, client *rds.Client) {
	input := &rds.DeregisterDBProxyTargetsInput{
		// DBProxyName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsDBClusterIdentifiers) > 0 {
		input.DBClusterIdentifiers = append([]string(nil), _rdsDBClusterIdentifiers...)
	}
	if len(_rdsDBInstanceIdentifiers) > 0 {
		input.DBInstanceIdentifiers = append([]string(nil), _rdsDBInstanceIdentifiers...)
	}
	if len(_rdsTargetGroupName) > 0 {
		input.TargetGroupName = aws.String(_rdsTargetGroupName)
	}

	if resp, err := client.DeregisterDBProxyTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the attributes for a customer account. The attributes include
// Amazon RDS quotas for the account, such as the number of DB instances allowed.
// The description for a quota includes the quota name, current usage toward that
// quota, and the quota's maximum value.
//
// This command doesn't take any parameters.
func rds_DescribeAccountAttributes(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeAccountAttributesInput{}

	if resp, err := client.DescribeAccountAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more blue/green deployments.
// For more information, see [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon RDS User Guide and [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon
// Aurora User Guide.
//
// [Using Amazon RDS Blue/Green Deployments for database updates]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html
func rds_DescribeBlueGreenDeployments(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeBlueGreenDeploymentsInput{}

	if len(_rdsBlueGreenDeploymentIdentifier) > 0 {
		input.BlueGreenDeploymentIdentifier = aws.String(_rdsBlueGreenDeploymentIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBlueGreenDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeBlueGreenDeploymentsOutput
	p := rds.NewDescribeBlueGreenDeploymentsPaginator(client, input)
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

// Lists the set of certificate authority (CA) certificates provided by Amazon RDS
// for this Amazon Web Services account.
//
// For more information, see [Using SSL/TLS to encrypt a connection to a DB instance] in the Amazon RDS User Guide and [Using SSL/TLS to encrypt a connection to a DB cluster] in the Amazon
// Aurora User Guide.
//
// [Using SSL/TLS to encrypt a connection to a DB cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html
// [Using SSL/TLS to encrypt a connection to a DB instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html
func rds_DescribeCertificates(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeCertificatesInput{}

	if len(_rdsCertificateIdentifier) > 0 {
		input.CertificateIdentifier = aws.String(_rdsCertificateIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeCertificatesOutput
	p := rds.NewDescribeCertificatesPaginator(client, input)
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

// Displays backups for both current and deleted DB clusters. For example, use
// this operation to find details about automated backups for previously deleted
// clusters. Current clusters are returned for both the
// DescribeDBClusterAutomatedBackups and DescribeDBClusters operations.
//
// All parameters are optional.
func rds_DescribeDBClusterAutomatedBackups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterAutomatedBackupsInput{}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsDbClusterResourceId) > 0 {
		input.DbClusterResourceId = aws.String(_rdsDbClusterResourceId)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterAutomatedBackups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClusterAutomatedBackupsOutput
	p := rds.NewDescribeDBClusterAutomatedBackupsPaginator(client, input)
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

// Returns information about backtracks for a DB cluster.
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// This action only applies to Aurora MySQL DB clusters.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
func rds_DescribeDBClusterBacktracks(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterBacktracksInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsBacktrackIdentifier) > 0 {
		input.BacktrackIdentifier = aws.String(_rdsBacktrackIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterBacktracks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClusterBacktracksOutput
	p := rds.NewDescribeDBClusterBacktracksPaginator(client, input)
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

// Returns information about endpoints for an Amazon Aurora DB cluster.
// This action only applies to Aurora DB clusters.
func rds_DescribeDBClusterEndpoints(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterEndpointsInput{}

	if len(_rdsDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_rdsDBClusterEndpointIdentifier)
	}
	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClusterEndpointsOutput
	p := rds.NewDescribeDBClusterEndpointsPaginator(client, input)
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

// Returns a list of DBClusterParameterGroup descriptions. If a
// DBClusterParameterGroupName parameter is specified, the list will contain only
// the description of the specified DB cluster parameter group.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DescribeDBClusterParameterGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterParameterGroupsInput{}

	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClusterParameterGroupsOutput
	p := rds.NewDescribeDBClusterParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular DB cluster parameter group.
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DescribeDBClusterParameters(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterParametersInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSource) > 0 {
		input.Source = aws.String(_rdsSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClusterParametersOutput
	p := rds.NewDescribeDBClusterParametersPaginator(client, input)
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

// Returns a list of DB cluster snapshot attribute names and values for a manual
// DB cluster snapshot.
//
// When sharing snapshots with other Amazon Web Services accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the Amazon Web Services accounts that are authorized to copy or restore
// the manual DB cluster snapshot. If all is included in the list of values for
// the restore attribute, then the manual DB cluster snapshot is public and can be
// copied or restored by all Amazon Web Services accounts.
//
// To add or remove access for an Amazon Web Services account to copy or restore a
// manual DB cluster snapshot, or to make the manual DB cluster snapshot public or
// private, use the ModifyDBClusterSnapshotAttribute API action.
func rds_DescribeDBClusterSnapshotAttributes(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterSnapshotAttributesInput{
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_rdsDBClusterSnapshotIdentifier)
	}

	if resp, err := client.DescribeDBClusterSnapshotAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about DB cluster snapshots. This API action supports
// pagination.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DescribeDBClusterSnapshots(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClusterSnapshotsInput{}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_rdsDBClusterSnapshotIdentifier)
	}
	if len(_rdsDbClusterResourceId) > 0 {
		input.DbClusterResourceId = aws.String(_rdsDbClusterResourceId)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsIncludePublic) > 0 {
		if err := assignInputField(input, "IncludePublic", _rdsIncludePublic); err != nil {
			log.Errorf("invalid --include-public: %s", err.Error())
			return
		}
	}
	if len(_rdsIncludeShared) > 0 {
		if err := assignInputField(input, "IncludeShared", _rdsIncludeShared); err != nil {
			log.Errorf("invalid --include-shared: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSnapshotType) > 0 {
		input.SnapshotType = aws.String(_rdsSnapshotType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClusterSnapshotsOutput
	p := rds.NewDescribeDBClusterSnapshotsPaginator(client, input)
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

// Describes existing Amazon Aurora DB clusters and Multi-AZ DB clusters. This API
// supports pagination.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// This operation can also return information for Amazon Neptune DB instances and
// Amazon DocumentDB instances.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_DescribeDBClusters(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBClustersInput{}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsIncludeShared) > 0 {
		if err := assignInputField(input, "IncludeShared", _rdsIncludeShared); err != nil {
			log.Errorf("invalid --include-shared: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBClustersOutput
	p := rds.NewDescribeDBClustersPaginator(client, input)
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

// Describes the properties of specific versions of DB engines.
func rds_DescribeDBEngineVersions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBEngineVersionsInput{}

	if len(_rdsDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_rdsDBParameterGroupFamily)
	}
	if len(_rdsDefaultOnly) > 0 {
		if err := assignInputField(input, "DefaultOnly", _rdsDefaultOnly); err != nil {
			log.Errorf("invalid --default-only: %s", err.Error())
			return
		}
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsIncludeAll) > 0 {
		if err := assignInputField(input, "IncludeAll", _rdsIncludeAll); err != nil {
			log.Errorf("invalid --include-all: %s", err.Error())
			return
		}
	}
	if len(_rdsListSupportedCharacterSets) > 0 {
		if err := assignInputField(input, "ListSupportedCharacterSets", _rdsListSupportedCharacterSets); err != nil {
			log.Errorf("invalid --list-supported-character-sets: %s", err.Error())
			return
		}
	}
	if len(_rdsListSupportedTimezones) > 0 {
		if err := assignInputField(input, "ListSupportedTimezones", _rdsListSupportedTimezones); err != nil {
			log.Errorf("invalid --list-supported-timezones: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBEngineVersionsOutput
	p := rds.NewDescribeDBEngineVersionsPaginator(client, input)
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

// Displays backups for both current and deleted instances. For example, use this
// operation to find details about automated backups for previously deleted
// instances. Current instances with retention periods greater than zero (0) are
// returned for both the DescribeDBInstanceAutomatedBackups and DescribeDBInstances
// operations.
//
// All parameters are optional.
func rds_DescribeDBInstanceAutomatedBackups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBInstanceAutomatedBackupsInput{}

	if len(_rdsDBInstanceAutomatedBackupsArn) > 0 {
		input.DBInstanceAutomatedBackupsArn = aws.String(_rdsDBInstanceAutomatedBackupsArn)
	}
	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsDbiResourceId) > 0 {
		input.DbiResourceId = aws.String(_rdsDbiResourceId)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBInstanceAutomatedBackups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBInstanceAutomatedBackupsOutput
	p := rds.NewDescribeDBInstanceAutomatedBackupsPaginator(client, input)
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

// Describes provisioned RDS instances. This API supports pagination.
// This operation can also return information for Amazon Neptune DB instances and
// Amazon DocumentDB instances.
func rds_DescribeDBInstances(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBInstancesInput{}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBInstancesOutput
	p := rds.NewDescribeDBInstancesPaginator(client, input)
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

// Returns a list of DB log files for the DB instance.
// This command doesn't apply to RDS Custom.
func rds_DescribeDBLogFiles(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBLogFilesInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsFileLastWritten) > 0 {
		if err := assignInputField(input, "FileLastWritten", _rdsFileLastWritten); err != nil {
			log.Errorf("invalid --file-last-written: %s", err.Error())
			return
		}
	}
	if len(_rdsFileSize) > 0 {
		if err := assignInputField(input, "FileSize", _rdsFileSize); err != nil {
			log.Errorf("invalid --file-size: %s", err.Error())
			return
		}
	}
	if len(_rdsFilenameContains) > 0 {
		input.FilenameContains = aws.String(_rdsFilenameContains)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBLogFiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBLogFilesOutput
	p := rds.NewDescribeDBLogFilesPaginator(client, input)
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

// Describes the properties of specific major versions of DB engines.
func rds_DescribeDBMajorEngineVersions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBMajorEngineVersionsInput{}

	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsMajorEngineVersion) > 0 {
		input.MajorEngineVersion = aws.String(_rdsMajorEngineVersion)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBMajorEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBMajorEngineVersionsOutput
	p := rds.NewDescribeDBMajorEngineVersionsPaginator(client, input)
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

// Returns a list of DBParameterGroup descriptions. If a DBParameterGroupName is
// specified, the list will contain only the description of the specified DB
// parameter group.
func rds_DescribeDBParameterGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBParameterGroupsInput{}

	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBParameterGroupsOutput
	p := rds.NewDescribeDBParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular DB parameter group.
func rds_DescribeDBParameters(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBParametersInput{
		// DBParameterGroupName: *string, // Required
	}

	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSource) > 0 {
		input.Source = aws.String(_rdsSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBParametersOutput
	p := rds.NewDescribeDBParametersPaginator(client, input)
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

// Returns information about DB proxies.
func rds_DescribeDBProxies(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBProxiesInput{}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBProxies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBProxiesOutput
	p := rds.NewDescribeDBProxiesPaginator(client, input)
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

// Returns information about DB proxy endpoints.
func rds_DescribeDBProxyEndpoints(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBProxyEndpointsInput{}

	if len(_rdsDBProxyEndpointName) > 0 {
		input.DBProxyEndpointName = aws.String(_rdsDBProxyEndpointName)
	}
	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBProxyEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBProxyEndpointsOutput
	p := rds.NewDescribeDBProxyEndpointsPaginator(client, input)
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

// Returns information about DB proxy target groups, represented by
// DBProxyTargetGroup data structures.
func rds_DescribeDBProxyTargetGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBProxyTargetGroupsInput{
		// DBProxyName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetGroupName) > 0 {
		input.TargetGroupName = aws.String(_rdsTargetGroupName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBProxyTargetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBProxyTargetGroupsOutput
	p := rds.NewDescribeDBProxyTargetGroupsPaginator(client, input)
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

// Returns information about DBProxyTarget objects. This API supports pagination.
func rds_DescribeDBProxyTargets(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBProxyTargetsInput{
		// DBProxyName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsTargetGroupName) > 0 {
		input.TargetGroupName = aws.String(_rdsTargetGroupName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBProxyTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBProxyTargetsOutput
	p := rds.NewDescribeDBProxyTargetsPaginator(client, input)
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

// Describes the recommendations to resolve the issues for your DB instances, DB
// clusters, and DB parameter groups.
func rds_DescribeDBRecommendations(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBRecommendationsInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsLastUpdatedAfter) > 0 {
		if err := assignInputField(input, "LastUpdatedAfter", _rdsLastUpdatedAfter); err != nil {
			log.Errorf("invalid --last-updated-after: %s", err.Error())
			return
		}
	}
	if len(_rdsLastUpdatedBefore) > 0 {
		if err := assignInputField(input, "LastUpdatedBefore", _rdsLastUpdatedBefore); err != nil {
			log.Errorf("invalid --last-updated-before: %s", err.Error())
			return
		}
	}
	if len(_rdsLocale) > 0 {
		input.Locale = aws.String(_rdsLocale)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBRecommendationsOutput
	p := rds.NewDescribeDBRecommendationsPaginator(client, input)
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

// Returns a list of DBSecurityGroup descriptions. If a DBSecurityGroupName is
// specified, the list will contain only the descriptions of the specified DB
// security group.
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func rds_DescribeDBSecurityGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBSecurityGroupsInput{}

	if len(_rdsDBSecurityGroupName) > 0 {
		input.DBSecurityGroupName = aws.String(_rdsDBSecurityGroupName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBSecurityGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBSecurityGroupsOutput
	p := rds.NewDescribeDBSecurityGroupsPaginator(client, input)
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

// Describes existing Aurora Limitless Database DB shard groups.
func rds_DescribeDBShardGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBShardGroupsInput{}

	if len(_rdsDBShardGroupIdentifier) > 0 {
		input.DBShardGroupIdentifier = aws.String(_rdsDBShardGroupIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeDBShardGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of DB snapshot attribute names and values for a manual DB
// snapshot.
//
// When sharing snapshots with other Amazon Web Services accounts,
// DescribeDBSnapshotAttributes returns the restore attribute and a list of IDs
// for the Amazon Web Services accounts that are authorized to copy or restore the
// manual DB snapshot. If all is included in the list of values for the restore
// attribute, then the manual DB snapshot is public and can be copied or restored
// by all Amazon Web Services accounts.
//
// To add or remove access for an Amazon Web Services account to copy or restore a
// manual DB snapshot, or to make the manual DB snapshot public or private, use the
// ModifyDBSnapshotAttribute API action.
func rds_DescribeDBSnapshotAttributes(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBSnapshotAttributesInput{
		// DBSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}

	if resp, err := client.DescribeDBSnapshotAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the tenant databases that exist in a DB snapshot. This command only
// applies to RDS for Oracle DB instances in the multi-tenant configuration.
//
// You can use this command to inspect the tenant databases within a snapshot
// before restoring it. You can't directly interact with the tenant databases in a
// DB snapshot. If you restore a snapshot that was taken from DB instance using the
// multi-tenant configuration, you restore all its tenant databases.
func rds_DescribeDBSnapshotTenantDatabases(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBSnapshotTenantDatabasesInput{}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}
	if len(_rdsDbiResourceId) > 0 {
		input.DbiResourceId = aws.String(_rdsDbiResourceId)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSnapshotType) > 0 {
		input.SnapshotType = aws.String(_rdsSnapshotType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBSnapshotTenantDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBSnapshotTenantDatabasesOutput
	p := rds.NewDescribeDBSnapshotTenantDatabasesPaginator(client, input)
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

// Returns information about DB snapshots. This API action supports pagination.
func rds_DescribeDBSnapshots(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBSnapshotsInput{}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}
	if len(_rdsDbiResourceId) > 0 {
		input.DbiResourceId = aws.String(_rdsDbiResourceId)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsIncludePublic) > 0 {
		if err := assignInputField(input, "IncludePublic", _rdsIncludePublic); err != nil {
			log.Errorf("invalid --include-public: %s", err.Error())
			return
		}
	}
	if len(_rdsIncludeShared) > 0 {
		if err := assignInputField(input, "IncludeShared", _rdsIncludeShared); err != nil {
			log.Errorf("invalid --include-shared: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSnapshotType) > 0 {
		input.SnapshotType = aws.String(_rdsSnapshotType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBSnapshotsOutput
	p := rds.NewDescribeDBSnapshotsPaginator(client, input)
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

// Returns a list of DBSubnetGroup descriptions. If a DBSubnetGroupName is
// specified, the list will contain only the descriptions of the specified
// DBSubnetGroup.
//
// For an overview of CIDR ranges, go to the [Wikipedia Tutorial].
//
// [Wikipedia Tutorial]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
func rds_DescribeDBSubnetGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeDBSubnetGroupsInput{}

	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBSubnetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeDBSubnetGroupsOutput
	p := rds.NewDescribeDBSubnetGroupsPaginator(client, input)
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

// Returns the default engine and system parameter information for the cluster
// database engine.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
func rds_DescribeEngineDefaultClusterParameters(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeEngineDefaultClusterParametersInput{
		// DBParameterGroupFamily: *string, // Required
	}

	if len(_rdsDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_rdsDBParameterGroupFamily)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEngineDefaultClusterParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeEngineDefaultClusterParametersOutput
	p := rds.NewDescribeEngineDefaultClusterParametersPaginator(client, input)
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
// database engine.
func rds_DescribeEngineDefaultParameters(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeEngineDefaultParametersInput{
		// DBParameterGroupFamily: *string, // Required
	}

	if len(_rdsDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_rdsDBParameterGroupFamily)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
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

	var results []*rds.DescribeEngineDefaultParametersOutput
	p := rds.NewDescribeEngineDefaultParametersPaginator(client, input)
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

// Displays a list of categories for all event source types, or, if specified, for
// a specified source type. You can also see this list in the "Amazon RDS event
// categories and event messages" section of the [Amazon RDS User Guide]or the [Amazon Aurora User Guide].
//
// [Amazon RDS User Guide]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Messages.html
// [Amazon Aurora User Guide]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Messages.html
func rds_DescribeEventCategories(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeEventCategoriesInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceType) > 0 {
		input.SourceType = aws.String(_rdsSourceType)
	}

	if resp, err := client.DescribeEventCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the subscription descriptions for a customer account. The description
// for a subscription includes SubscriptionName , SNSTopicARN , CustomerID ,
// SourceType , SourceID , CreationTime , and Status .
//
// If you specify a SubscriptionName , lists the description for that subscription.
func rds_DescribeEventSubscriptions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeEventSubscriptionsInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_rdsSubscriptionName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEventSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeEventSubscriptionsOutput
	p := rds.NewDescribeEventSubscriptionsPaginator(client, input)
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

// Returns events related to DB instances, DB clusters, DB parameter groups, DB
// security groups, DB snapshots, DB cluster snapshots, and RDS Proxies for the
// past 14 days. Events specific to a particular DB instance, DB cluster, DB
// parameter group, DB security group, DB snapshot, DB cluster snapshot group, or
// RDS Proxy can be obtained by providing the name as a parameter.
//
// For more information on working with events, see [Monitoring Amazon RDS events] in the Amazon RDS User Guide
// and [Monitoring Amazon Aurora events]in the Amazon Aurora User Guide.
//
// By default, RDS returns events that were generated in the past hour.
//
// [Monitoring Amazon RDS events]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/working-with-events.html
// [Monitoring Amazon Aurora events]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/working-with-events.html
func rds_DescribeEvents(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeEventsInput{}

	if len(_rdsDuration) > 0 {
		if err := assignInputField(input, "Duration", _rdsDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_rdsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _rdsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_rdsEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _rdsEventCategories...)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_rdsSourceIdentifier)
	}
	if len(_rdsSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _rdsSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_rdsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _rdsStartTime); err != nil {
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

	var results []*rds.DescribeEventsOutput
	p := rds.NewDescribeEventsPaginator(client, input)
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

// Returns information about a snapshot or cluster export to Amazon S3. This API
// operation supports pagination.
func rds_DescribeExportTasks(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeExportTasksInput{}

	if len(_rdsExportTaskIdentifier) > 0 {
		input.ExportTaskIdentifier = aws.String(_rdsExportTaskIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceArn) > 0 {
		input.SourceArn = aws.String(_rdsSourceArn)
	}
	if len(_rdsSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _rdsSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeExportTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeExportTasksOutput
	p := rds.NewDescribeExportTasksPaginator(client, input)
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

// Returns information about Aurora global database clusters. This API supports
// pagination.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
func rds_DescribeGlobalClusters(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeGlobalClustersInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeGlobalClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeGlobalClustersOutput
	p := rds.NewDescribeGlobalClustersPaginator(client, input)
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

// Describe one or more zero-ETL integrations with Amazon Redshift.
func rds_DescribeIntegrations(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeIntegrationsInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsIntegrationIdentifier) > 0 {
		input.IntegrationIdentifier = aws.String(_rdsIntegrationIdentifier)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeIntegrationsOutput
	p := rds.NewDescribeIntegrationsPaginator(client, input)
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

// Describes all available options for the specified engine.
func rds_DescribeOptionGroupOptions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeOptionGroupOptionsInput{
		// EngineName: *string, // Required
	}

	if len(_rdsEngineName) > 0 {
		input.EngineName = aws.String(_rdsEngineName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMajorEngineVersion) > 0 {
		input.MajorEngineVersion = aws.String(_rdsMajorEngineVersion)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeOptionGroupOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeOptionGroupOptionsOutput
	p := rds.NewDescribeOptionGroupOptionsPaginator(client, input)
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

// Describes the available option groups.
func rds_DescribeOptionGroups(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeOptionGroupsInput{}

	if len(_rdsEngineName) > 0 {
		input.EngineName = aws.String(_rdsEngineName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMajorEngineVersion) > 0 {
		input.MajorEngineVersion = aws.String(_rdsMajorEngineVersion)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOptionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeOptionGroupsOutput
	p := rds.NewDescribeOptionGroupsPaginator(client, input)
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

// Describes the orderable DB instance options for a specified DB engine.
func rds_DescribeOrderableDBInstanceOptions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeOrderableDBInstanceOptionsInput{
		// Engine: *string, // Required
	}

	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsAvailabilityZoneGroup) > 0 {
		input.AvailabilityZoneGroup = aws.String(_rdsAvailabilityZoneGroup)
	}
	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsLicenseModel) > 0 {
		input.LicenseModel = aws.String(_rdsLicenseModel)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsVpc) > 0 {
		if err := assignInputField(input, "Vpc", _rdsVpc); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrderableDBInstanceOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeOrderableDBInstanceOptionsOutput
	p := rds.NewDescribeOrderableDBInstanceOptionsPaginator(client, input)
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

// Returns a list of resources (for example, DB instances) that have at least one
// pending maintenance action.
//
// This API follows an eventual consistency model. This means that the result of
// the DescribePendingMaintenanceActions command might not be immediately visible
// to all subsequent RDS commands. Keep this in mind when you use
// DescribePendingMaintenanceActions immediately after using a previous API command
// such as ApplyPendingMaintenanceActions .
func rds_DescribePendingMaintenanceActions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribePendingMaintenanceActionsInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_rdsResourceIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.DescribePendingMaintenanceActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribePendingMaintenanceActionsOutput
	p := rds.NewDescribePendingMaintenanceActionsPaginator(client, input)
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

// Returns information about reserved DB instances for this account, or about a
// specified reserved DB instance.
func rds_DescribeReservedDBInstances(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeReservedDBInstancesInput{}

	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDuration) > 0 {
		input.Duration = aws.String(_rdsDuration)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsLeaseId) > 0 {
		input.LeaseId = aws.String(_rdsLeaseId)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsOfferingType) > 0 {
		input.OfferingType = aws.String(_rdsOfferingType)
	}
	if len(_rdsProductDescription) > 0 {
		input.ProductDescription = aws.String(_rdsProductDescription)
	}
	if len(_rdsReservedDBInstanceId) > 0 {
		input.ReservedDBInstanceId = aws.String(_rdsReservedDBInstanceId)
	}
	if len(_rdsReservedDBInstancesOfferingId) > 0 {
		input.ReservedDBInstancesOfferingId = aws.String(_rdsReservedDBInstancesOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedDBInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeReservedDBInstancesOutput
	p := rds.NewDescribeReservedDBInstancesPaginator(client, input)
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

// Lists available reserved DB instance offerings.
func rds_DescribeReservedDBInstancesOfferings(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeReservedDBInstancesOfferingsInput{}

	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDuration) > 0 {
		input.Duration = aws.String(_rdsDuration)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsOfferingType) > 0 {
		input.OfferingType = aws.String(_rdsOfferingType)
	}
	if len(_rdsProductDescription) > 0 {
		input.ProductDescription = aws.String(_rdsProductDescription)
	}
	if len(_rdsReservedDBInstancesOfferingId) > 0 {
		input.ReservedDBInstancesOfferingId = aws.String(_rdsReservedDBInstancesOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedDBInstancesOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeReservedDBInstancesOfferingsOutput
	p := rds.NewDescribeReservedDBInstancesOfferingsPaginator(client, input)
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

// Returns a list of the source Amazon Web Services Regions where the current
// Amazon Web Services Region can create a read replica, copy a DB snapshot from,
// or replicate automated backups from.
//
// Use this operation to determine whether cross-Region features are supported
// between other Regions and your current Region. This operation supports
// pagination.
//
// To return information about the Regions that are enabled for your account, or
// all Regions, use the EC2 operation DescribeRegions . For more information, see [DescribeRegions]
// in the Amazon EC2 API Reference.
//
// [DescribeRegions]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeRegions.html
func rds_DescribeSourceRegions(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeSourceRegionsInput{}

	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsRegionName) > 0 {
		input.RegionName = aws.String(_rdsRegionName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSourceRegions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeSourceRegionsOutput
	p := rds.NewDescribeSourceRegionsPaginator(client, input)
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

// Describes the tenant databases in a DB instance that uses the multi-tenant
// configuration. Only RDS for Oracle CDB instances are supported.
func rds_DescribeTenantDatabases(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeTenantDatabasesInput{}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _rdsMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_rdsTenantDBName) > 0 {
		input.TenantDBName = aws.String(_rdsTenantDBName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTenantDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DescribeTenantDatabasesOutput
	p := rds.NewDescribeTenantDatabasesPaginator(client, input)
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

// You can call DescribeValidDBInstanceModifications to learn what modifications
// you can make to your DB instance. You can use this information when you call
// ModifyDBInstance .
//
// This command doesn't apply to RDS Custom.
func rds_DescribeValidDBInstanceModifications(cfg aws.Config, client *rds.Client) {
	input := &rds.DescribeValidDBInstanceModificationsInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}

	if resp, err := client.DescribeValidDBInstanceModifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the HTTP endpoint for the specified DB cluster. Disabling this
// endpoint disables RDS Data API.
//
// For more information, see [Using RDS Data API] in the Amazon Aurora User Guide.
//
// This operation applies only to Aurora Serverless v2 and provisioned DB
// clusters. To disable the HTTP endpoint for Aurora Serverless v1 DB clusters, use
// the EnableHttpEndpoint parameter of the ModifyDBCluster operation.
//
// [Using RDS Data API]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
func rds_DisableHttpEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.DisableHttpEndpointInput{
		// ResourceArn: *string, // Required
	}

	if len(_rdsResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsResourceArn)
	}

	if resp, err := client.DisableHttpEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Downloads all or a portion of the specified log file, up to 1 MB in size.
// This command doesn't apply to RDS Custom.
//
// This operation uses resources on database instances. Because of this, we
// recommend publishing database logs to CloudWatch and then using the GetLogEvents
// operation. For more information, see [GetLogEvents]in the Amazon CloudWatch Logs API
// Reference.
//
// [GetLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html
func rds_DownloadDBLogFilePortion(cfg aws.Config, client *rds.Client) {
	input := &rds.DownloadDBLogFilePortionInput{
		// DBInstanceIdentifier: *string, // Required
		// LogFileName: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsLogFileName) > 0 {
		input.LogFileName = aws.String(_rdsLogFileName)
	}
	if len(_rdsMarker) > 0 {
		input.Marker = aws.String(_rdsMarker)
	}
	if len(_rdsNumberOfLines) > 0 {
		if err := assignInputField(input, "NumberOfLines", _rdsNumberOfLines); err != nil {
			log.Errorf("invalid --number-of-lines: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DownloadDBLogFilePortion(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rds.DownloadDBLogFilePortionOutput
	p := rds.NewDownloadDBLogFilePortionPaginator(client, input)
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

// Enables the HTTP endpoint for the DB cluster. By default, the HTTP endpoint
// isn't enabled.
//
// When enabled, this endpoint provides a connectionless web service API (RDS Data
// API) for running SQL queries on the Aurora DB cluster. You can also query your
// database from inside the RDS console with the RDS query editor.
//
// For more information, see [Using RDS Data API] in the Amazon Aurora User Guide.
//
// This operation applies only to Aurora Serverless v2 and provisioned DB
// clusters. To enable the HTTP endpoint for Aurora Serverless v1 DB clusters, use
// the EnableHttpEndpoint parameter of the ModifyDBCluster operation.
//
// [Using RDS Data API]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
func rds_EnableHttpEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.EnableHttpEndpointInput{
		// ResourceArn: *string, // Required
	}

	if len(_rdsResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsResourceArn)
	}

	if resp, err := client.EnableHttpEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Forces a failover for a DB cluster.
// For an Aurora DB cluster, failover for a DB cluster promotes one of the Aurora
// Replicas (read-only instances) in the DB cluster to be the primary DB instance
// (the cluster writer).
//
// For a Multi-AZ DB cluster, after RDS terminates the primary DB instance, the
// internal monitoring system detects that the primary DB instance is unhealthy and
// promotes a readable standby (read-only instances) in the DB cluster to be the
// primary DB instance (the cluster writer). Failover times are typically less than
// 35 seconds.
//
// An Amazon Aurora DB cluster automatically fails over to an Aurora Replica, if
// one exists, when the primary DB instance fails. A Multi-AZ DB cluster
// automatically fails over to a readable standby DB instance when the primary DB
// instance fails.
//
// To simulate a failure of a primary instance for testing, you can force a
// failover. Because each instance in a DB cluster has its own endpoint address,
// make sure to clean up and re-establish any existing connections that use those
// endpoint addresses when the failover is complete.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_FailoverDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.FailoverDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsTargetDBInstanceIdentifier) > 0 {
		input.TargetDBInstanceIdentifier = aws.String(_rdsTargetDBInstanceIdentifier)
	}

	if resp, err := client.FailoverDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Promotes the specified secondary DB cluster to be the primary DB cluster in the
// global database cluster to fail over or switch over a global database.
// Switchover operations were previously called "managed planned failovers."
//
// Although this operation can be used either to fail over or to switch over a
// global database cluster, its intended use is for global database failover. To
// switch over a global database cluster, we recommend that you use the SwitchoverGlobalClusteroperation
// instead.
//
// How you use this operation depends on whether you are failing over or switching
// over your global database cluster:
//
// - Failing over - Specify the AllowDataLoss parameter and don't specify the
// Switchover parameter.
//
// - Switching over - Specify the Switchover parameter or omit it, but don't
// specify the AllowDataLoss parameter.
//
// # About failing over and switching over
//
// While failing over and switching over a global database cluster both change the
// primary DB cluster, you use these operations for different reasons:
//
// - Failing over - Use this operation to respond to an unplanned event, such as
// a Regional disaster in the primary Region. Failing over can result in a loss of
// write transaction data that wasn't replicated to the chosen secondary before the
// failover event occurred. However, the recovery process that promotes a DB
// instance on the chosen seconday DB cluster to be the primary writer DB instance
// guarantees that the data is in a transactionally consistent state.
//
// For more information about failing over an Amazon Aurora global database, see [Performing managed failovers for Aurora global databases]
//
// in the Amazon Aurora User Guide.
//
// - Switching over - Use this operation on a healthy global database cluster
// for planned events, such as Regional rotation or to fail back to the original
// primary DB cluster after a failover operation. With this operation, there is no
// data loss.
//
// For more information about switching over an Amazon Aurora global database, see [Performing switchovers for Aurora global databases]
//
// in the Amazon Aurora User Guide.
//
// [Performing managed failovers for Aurora global databases]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-failover.managed-unplanned
// [Performing switchovers for Aurora global databases]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-disaster-recovery.managed-failover
func rds_FailoverGlobalCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.FailoverGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
		// TargetDbClusterIdentifier: *string, // Required
	}

	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}
	if len(_rdsTargetDbClusterIdentifier) > 0 {
		input.TargetDbClusterIdentifier = aws.String(_rdsTargetDbClusterIdentifier)
	}
	if len(_rdsAllowDataLoss) > 0 {
		if err := assignInputField(input, "AllowDataLoss", _rdsAllowDataLoss); err != nil {
			log.Errorf("invalid --allow-data-loss: %s", err.Error())
			return
		}
	}
	if len(_rdsSwitchover) > 0 {
		if err := assignInputField(input, "Switchover", _rdsSwitchover); err != nil {
			log.Errorf("invalid --switchover: %s", err.Error())
			return
		}
	}

	if resp, err := client.FailoverGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all tags on an Amazon RDS resource.
// For an overview on tagging an Amazon RDS resource, see [Tagging Amazon RDS Resources] in the Amazon RDS User
// Guide or [Tagging Amazon Aurora and Amazon RDS Resources]in the Amazon Aurora User Guide.
//
// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
// [Tagging Amazon Aurora and Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html
func rds_ListTagsForResource(cfg aws.Config, client *rds.Client) {
	input := &rds.ListTagsForResourceInput{
		// ResourceName: *string, // Required
	}

	if len(_rdsResourceName) > 0 {
		input.ResourceName = aws.String(_rdsResourceName)
	}
	if len(_rdsFilters) > 0 {
		if err := assignInputField(input, "Filters", _rdsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the audit policy state of a database activity stream to either locked
// (default) or unlocked. A locked policy is read-only, whereas an unlocked policy
// is read/write. If your activity stream is started and locked, you can unlock it,
// customize your audit policy, and then lock your activity stream. Restarting the
// activity stream isn't required. For more information, see [Modifying a database activity stream]in the Amazon RDS
// User Guide.
//
// This operation is supported for RDS for Oracle and Microsoft SQL Server.
//
// [Modifying a database activity stream]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.Modifying.html
func rds_ModifyActivityStream(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyActivityStreamInput{}

	if len(_rdsAuditPolicyState) > 0 {
		if err := assignInputField(input, "AuditPolicyState", _rdsAuditPolicyState); err != nil {
			log.Errorf("invalid --audit-policy-state: %s", err.Error())
			return
		}
	}
	if len(_rdsResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsResourceArn)
	}

	if resp, err := client.ModifyActivityStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Override the system-default Secure Sockets Layer/Transport Layer Security
// (SSL/TLS) certificate for Amazon RDS for new DB instances, or remove the
// override.
//
// By using this operation, you can specify an RDS-approved SSL/TLS certificate
// for new DB instances that is different from the default certificate provided by
// RDS. You can also use this operation to remove the override, so that new DB
// instances use the default certificate provided by RDS.
//
// You might need to override the default certificate in the following situations:
//
// - You already migrated your applications to support the latest certificate
// authority (CA) certificate, but the new CA certificate is not yet the RDS
// default CA certificate for the specified Amazon Web Services Region.
//
// - RDS has already moved to a new default CA certificate for the specified
// Amazon Web Services Region, but you are still in the process of supporting the
// new CA certificate. In this case, you temporarily need additional time to finish
// your application changes.
//
// For more information about rotating your SSL/TLS certificate for RDS DB
// engines, see [Rotating Your SSL/TLS Certificate]in the Amazon RDS User Guide.
//
// For more information about rotating your SSL/TLS certificate for Aurora DB
// engines, see [Rotating Your SSL/TLS Certificate]in the Amazon Aurora User Guide.
//
// [Rotating Your SSL/TLS Certificate]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html
func rds_ModifyCertificates(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyCertificatesInput{}

	if len(_rdsCertificateIdentifier) > 0 {
		input.CertificateIdentifier = aws.String(_rdsCertificateIdentifier)
	}
	if len(_rdsRemoveCustomerOverride) > 0 {
		if err := assignInputField(input, "RemoveCustomerOverride", _rdsRemoveCustomerOverride); err != nil {
			log.Errorf("invalid --remove-customer-override: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyCertificates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the capacity of an Aurora Serverless v1 DB cluster to a specific value.
// Aurora Serverless v1 scales seamlessly based on the workload on the DB cluster.
// In some cases, the capacity might not scale fast enough to meet a sudden change
// in workload, such as a large number of new transactions. Call
// ModifyCurrentDBClusterCapacity to set the capacity explicitly.
//
// After this call sets the DB cluster capacity, Aurora Serverless v1 can
// automatically scale the DB cluster based on the cooldown period for scaling up
// and the cooldown period for scaling down.
//
// For more information about Aurora Serverless v1, see [Using Amazon Aurora Serverless v1] in the Amazon Aurora User
// Guide.
//
// If you call ModifyCurrentDBClusterCapacity with the default TimeoutAction ,
// connections that prevent Aurora Serverless v1 from finding a scaling point might
// be dropped. For more information about scaling points, see [Autoscaling for Aurora Serverless v1]in the Amazon Aurora
// User Guide.
//
// This operation only applies to Aurora Serverless v1 DB clusters.
//
// [Autoscaling for Aurora Serverless v1]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling
// [Using Amazon Aurora Serverless v1]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html
func rds_ModifyCurrentDBClusterCapacity(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyCurrentDBClusterCapacityInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _rdsCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_rdsSecondsBeforeTimeout) > 0 {
		if err := assignInputField(input, "SecondsBeforeTimeout", _rdsSecondsBeforeTimeout); err != nil {
			log.Errorf("invalid --seconds-before-timeout: %s", err.Error())
			return
		}
	}
	if len(_rdsTimeoutAction) > 0 {
		input.TimeoutAction = aws.String(_rdsTimeoutAction)
	}

	if resp, err := client.ModifyCurrentDBClusterCapacity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the status of a custom engine version (CEV). You can find CEVs to
// modify by calling DescribeDBEngineVersions .
//
// The MediaImport service that imports files from Amazon S3 to create CEVs isn't
// integrated with Amazon Web Services CloudTrail. If you turn on data logging for
// Amazon RDS in CloudTrail, calls to the ModifyCustomDbEngineVersion event aren't
// logged. However, you might see calls from the API gateway that accesses your
// Amazon S3 bucket. These calls originate from the MediaImport service for the
// ModifyCustomDbEngineVersion event.
//
// For more information, see [Modifying CEV status] in the Amazon RDS User Guide.
//
// [Modifying CEV status]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.modify
func rds_ModifyCustomDBEngineVersion(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyCustomDBEngineVersionInput{
		// Engine: *string, // Required
		// EngineVersion: *string, // Required
	}

	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsDescription) > 0 {
		input.Description = aws.String(_rdsDescription)
	}
	if len(_rdsStatus) > 0 {
		if err := assignInputField(input, "Status", _rdsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyCustomDBEngineVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings of an Amazon Aurora DB cluster or a Multi-AZ DB cluster.
// You can change one or more settings by specifying these parameters and the new
// values in the request.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_ModifyDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAllowEngineModeChange) > 0 {
		if err := assignInputField(input, "AllowEngineModeChange", _rdsAllowEngineModeChange); err != nil {
			log.Errorf("invalid --allow-engine-mode-change: %s", err.Error())
			return
		}
	}
	if len(_rdsAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _rdsAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _rdsApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAwsBackupRecoveryPointArn) > 0 {
		input.AwsBackupRecoveryPointArn = aws.String(_rdsAwsBackupRecoveryPointArn)
	}
	if len(_rdsBacktrackWindow) > 0 {
		if err := assignInputField(input, "BacktrackWindow", _rdsBacktrackWindow); err != nil {
			log.Errorf("invalid --backtrack-window: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCloudwatchLogsExportConfiguration) > 0 {
		if err := assignInputField(input, "CloudwatchLogsExportConfiguration", _rdsCloudwatchLogsExportConfiguration); err != nil {
			log.Errorf("invalid --cloudwatch-logs-export-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBClusterInstanceClass) > 0 {
		input.DBClusterInstanceClass = aws.String(_rdsDBClusterInstanceClass)
	}
	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsDBInstanceParameterGroupName) > 0 {
		input.DBInstanceParameterGroupName = aws.String(_rdsDBInstanceParameterGroupName)
	}
	if len(_rdsDatabaseInsightsMode) > 0 {
		if err := assignInputField(input, "DatabaseInsightsMode", _rdsDatabaseInsightsMode); err != nil {
			log.Errorf("invalid --database-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsEnableGlobalWriteForwarding) > 0 {
		if err := assignInputField(input, "EnableGlobalWriteForwarding", _rdsEnableGlobalWriteForwarding); err != nil {
			log.Errorf("invalid --enable-global-write-forwarding: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableLimitlessDatabase) > 0 {
		if err := assignInputField(input, "EnableLimitlessDatabase", _rdsEnableLimitlessDatabase); err != nil {
			log.Errorf("invalid --enable-limitless-database: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableLocalWriteForwarding) > 0 {
		if err := assignInputField(input, "EnableLocalWriteForwarding", _rdsEnableLocalWriteForwarding); err != nil {
			log.Errorf("invalid --enable-local-write-forwarding: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineMode) > 0 {
		input.EngineMode = aws.String(_rdsEngineMode)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserAuthenticationType) > 0 {
		if err := assignInputField(input, "MasterUserAuthenticationType", _rdsMasterUserAuthenticationType); err != nil {
			log.Errorf("invalid --master-user-authentication-type: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsNewDBClusterIdentifier) > 0 {
		input.NewDBClusterIdentifier = aws.String(_rdsNewDBClusterIdentifier)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_rdsPreferredMaintenanceWindow)
	}
	if len(_rdsRotateMasterUserPassword) > 0 {
		if err := assignInputField(input, "RotateMasterUserPassword", _rdsRotateMasterUserPassword); err != nil {
			log.Errorf("invalid --rotate-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _rdsScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _rdsServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the properties of an endpoint in an Amazon Aurora DB cluster.
// This operation only applies to Aurora DB clusters.
func rds_ModifyDBClusterEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBClusterEndpointInput{
		// DBClusterEndpointIdentifier: *string, // Required
	}

	if len(_rdsDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_rdsDBClusterEndpointIdentifier)
	}
	if len(_rdsEndpointType) > 0 {
		input.EndpointType = aws.String(_rdsEndpointType)
	}
	if len(_rdsExcludedMembers) > 0 {
		input.ExcludedMembers = append([]string(nil), _rdsExcludedMembers...)
	}
	if len(_rdsStaticMembers) > 0 {
		input.StaticMembers = append([]string(nil), _rdsStaticMembers...)
	}

	if resp, err := client.ModifyDBClusterEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB cluster parameter group. To modify more than
// one parameter, submit a list of the following: ParameterName , ParameterValue ,
// and ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// There are two types of parameters - dynamic parameters and static parameters.
// Changes to dynamic parameters are applied to the DB cluster immediately without
// a reboot. Changes to static parameters are applied only after the DB cluster is
// rebooted, which can be done using RebootDBCluster operation. You can use the
// Parameter Groups option of the [Amazon RDS console]or the DescribeDBClusterParameters operation to
// verify that your DB cluster parameter group has been created or modified.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [Amazon RDS console]: https://console.aws.amazon.com/rds/
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_ModifyDBClusterParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
		// Parameters: []types.Parameter, // Required
	}

	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _rdsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB cluster snapshot.
//
// To share a manual DB cluster snapshot with other Amazon Web Services accounts,
// specify restore as the AttributeName and use the ValuesToAdd parameter to add a
// list of IDs of the Amazon Web Services accounts that are authorized to restore
// the manual DB cluster snapshot. Use the value all to make the manual DB cluster
// snapshot public, which means that it can be copied or restored by all Amazon Web
// Services accounts.
//
// Don't add the all value for any manual DB cluster snapshots that contain
// private information that you don't want available to all Amazon Web Services
// accounts.
//
// If a manual DB cluster snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized Amazon Web Services account IDs for the
// ValuesToAdd parameter. You can't use all as a value for that parameter in this
// case.
//
// To view which Amazon Web Services accounts have access to copy or restore a
// manual DB cluster snapshot, or whether a manual DB cluster snapshot is public or
// private, use the DescribeDBClusterSnapshotAttributesAPI operation. The accounts are returned as values for the
// restore attribute.
func rds_ModifyDBClusterSnapshotAttribute(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBClusterSnapshotAttributeInput{
		// AttributeName: *string, // Required
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_rdsAttributeName) > 0 {
		input.AttributeName = aws.String(_rdsAttributeName)
	}
	if len(_rdsDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_rdsDBClusterSnapshotIdentifier)
	}
	if len(_rdsValuesToAdd) > 0 {
		input.ValuesToAdd = append([]string(nil), _rdsValuesToAdd...)
	}
	if len(_rdsValuesToRemove) > 0 {
		input.ValuesToRemove = append([]string(nil), _rdsValuesToRemove...)
	}

	if resp, err := client.ModifyDBClusterSnapshotAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies settings for a DB instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request. To learn what modifications you can make to your DB instance, call
// DescribeValidDBInstanceModifications before you call ModifyDBInstance .
func rds_ModifyDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsAdditionalStorageVolumes) > 0 {
		if err := assignInputField(input, "AdditionalStorageVolumes", _rdsAdditionalStorageVolumes); err != nil {
			log.Errorf("invalid --additional-storage-volumes: %s", err.Error())
			return
		}
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _rdsAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _rdsApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAutomationMode) > 0 {
		if err := assignInputField(input, "AutomationMode", _rdsAutomationMode); err != nil {
			log.Errorf("invalid --automation-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsAwsBackupRecoveryPointArn) > 0 {
		input.AwsBackupRecoveryPointArn = aws.String(_rdsAwsBackupRecoveryPointArn)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCertificateRotationRestart) > 0 {
		if err := assignInputField(input, "CertificateRotationRestart", _rdsCertificateRotationRestart); err != nil {
			log.Errorf("invalid --certificate-rotation-restart: %s", err.Error())
			return
		}
	}
	if len(_rdsCloudwatchLogsExportConfiguration) > 0 {
		if err := assignInputField(input, "CloudwatchLogsExportConfiguration", _rdsCloudwatchLogsExportConfiguration); err != nil {
			log.Errorf("invalid --cloudwatch-logs-export-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDBPortNumber) > 0 {
		if err := assignInputField(input, "DBPortNumber", _rdsDBPortNumber); err != nil {
			log.Errorf("invalid --db-port-number: %s", err.Error())
			return
		}
	}
	if len(_rdsDBSecurityGroups) > 0 {
		input.DBSecurityGroups = append([]string(nil), _rdsDBSecurityGroups...)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDatabaseInsightsMode) > 0 {
		if err := assignInputField(input, "DatabaseInsightsMode", _rdsDatabaseInsightsMode); err != nil {
			log.Errorf("invalid --database-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsDedicatedLogVolume) > 0 {
		if err := assignInputField(input, "DedicatedLogVolume", _rdsDedicatedLogVolume); err != nil {
			log.Errorf("invalid --dedicated-log-volume: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDisableDomain) > 0 {
		if err := assignInputField(input, "DisableDomain", _rdsDisableDomain); err != nil {
			log.Errorf("invalid --disable-domain: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainAuthSecretArn) > 0 {
		input.DomainAuthSecretArn = aws.String(_rdsDomainAuthSecretArn)
	}
	if len(_rdsDomainDnsIps) > 0 {
		input.DomainDnsIps = append([]string(nil), _rdsDomainDnsIps...)
	}
	if len(_rdsDomainFqdn) > 0 {
		input.DomainFqdn = aws.String(_rdsDomainFqdn)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsDomainOu) > 0 {
		input.DomainOu = aws.String(_rdsDomainOu)
	}
	if len(_rdsEnableCustomerOwnedIp) > 0 {
		if err := assignInputField(input, "EnableCustomerOwnedIp", _rdsEnableCustomerOwnedIp); err != nil {
			log.Errorf("invalid --enable-customer-owned-ip: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsLicenseModel) > 0 {
		input.LicenseModel = aws.String(_rdsLicenseModel)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserAuthenticationType) > 0 {
		if err := assignInputField(input, "MasterUserAuthenticationType", _rdsMasterUserAuthenticationType); err != nil {
			log.Errorf("invalid --master-user-authentication-type: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMaxAllocatedStorage) > 0 {
		if err := assignInputField(input, "MaxAllocatedStorage", _rdsMaxAllocatedStorage); err != nil {
			log.Errorf("invalid --max-allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsMultiTenant) > 0 {
		if err := assignInputField(input, "MultiTenant", _rdsMultiTenant); err != nil {
			log.Errorf("invalid --multi-tenant: %s", err.Error())
			return
		}
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsNewDBInstanceIdentifier) > 0 {
		input.NewDBInstanceIdentifier = aws.String(_rdsNewDBInstanceIdentifier)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_rdsPreferredMaintenanceWindow)
	}
	if len(_rdsProcessorFeatures) > 0 {
		if err := assignInputField(input, "ProcessorFeatures", _rdsProcessorFeatures); err != nil {
			log.Errorf("invalid --processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsPromotionTier) > 0 {
		if err := assignInputField(input, "PromotionTier", _rdsPromotionTier); err != nil {
			log.Errorf("invalid --promotion-tier: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsReplicaMode) > 0 {
		if err := assignInputField(input, "ReplicaMode", _rdsReplicaMode); err != nil {
			log.Errorf("invalid --replica-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsResumeFullAutomationModeMinutes) > 0 {
		if err := assignInputField(input, "ResumeFullAutomationModeMinutes", _rdsResumeFullAutomationModeMinutes); err != nil {
			log.Errorf("invalid --resume-full-automation-mode-minutes: %s", err.Error())
			return
		}
	}
	if len(_rdsRotateMasterUserPassword) > 0 {
		if err := assignInputField(input, "RotateMasterUserPassword", _rdsRotateMasterUserPassword); err != nil {
			log.Errorf("invalid --rotate-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageThroughput) > 0 {
		if err := assignInputField(input, "StorageThroughput", _rdsStorageThroughput); err != nil {
			log.Errorf("invalid --storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTdeCredentialArn) > 0 {
		input.TdeCredentialArn = aws.String(_rdsTdeCredentialArn)
	}
	if len(_rdsTdeCredentialPassword) > 0 {
		input.TdeCredentialPassword = aws.String(_rdsTdeCredentialPassword)
	}
	if len(_rdsUseDefaultProcessorFeatures) > 0 {
		if err := assignInputField(input, "UseDefaultProcessorFeatures", _rdsUseDefaultProcessorFeatures); err != nil {
			log.Errorf("invalid --use-default-processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName , ParameterValue , and
// ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// After you modify a DB parameter group, you should wait at least 5 minutes
// before creating your first DB instance that uses that DB parameter group as the
// default parameter group. This allows Amazon RDS to fully complete the modify
// operation before the parameter group is used as the default for a new DB
// instance. This is especially important for parameters that are critical when
// creating the default database for a DB instance, such as the character set for
// the default database defined by the character_set_database parameter. You can
// use the Parameter Groups option of the [Amazon RDS console]or the DescribeDBParameters command to
// verify that your DB parameter group has been created or modified.
//
// [Amazon RDS console]: https://console.aws.amazon.com/rds/
func rds_ModifyDBParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBParameterGroupInput{
		// DBParameterGroupName: *string, // Required
		// Parameters: []types.Parameter, // Required
	}

	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _rdsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the settings for an existing DB proxy.
func rds_ModifyDBProxy(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBProxyInput{
		// DBProxyName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsAuth) > 0 {
		if err := assignInputField(input, "Auth", _rdsAuth); err != nil {
			log.Errorf("invalid --auth: %s", err.Error())
			return
		}
	}
	if len(_rdsDebugLogging) > 0 {
		if err := assignInputField(input, "DebugLogging", _rdsDebugLogging); err != nil {
			log.Errorf("invalid --debug-logging: %s", err.Error())
			return
		}
	}
	if len(_rdsDefaultAuthScheme) > 0 {
		if err := assignInputField(input, "DefaultAuthScheme", _rdsDefaultAuthScheme); err != nil {
			log.Errorf("invalid --default-auth-scheme: %s", err.Error())
			return
		}
	}
	if len(_rdsIdleClientTimeout) > 0 {
		if err := assignInputField(input, "IdleClientTimeout", _rdsIdleClientTimeout); err != nil {
			log.Errorf("invalid --idle-client-timeout: %s", err.Error())
			return
		}
	}
	if len(_rdsNewDBProxyName) > 0 {
		input.NewDBProxyName = aws.String(_rdsNewDBProxyName)
	}
	if len(_rdsRequireTLS) > 0 {
		if err := assignInputField(input, "RequireTLS", _rdsRequireTLS); err != nil {
			log.Errorf("invalid --require-tls: %s", err.Error())
			return
		}
	}
	if len(_rdsRoleArn) > 0 {
		input.RoleArn = aws.String(_rdsRoleArn)
	}
	if len(_rdsSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _rdsSecurityGroups...)
	}

	if resp, err := client.ModifyDBProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the settings for an existing DB proxy endpoint.
func rds_ModifyDBProxyEndpoint(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBProxyEndpointInput{
		// DBProxyEndpointName: *string, // Required
	}

	if len(_rdsDBProxyEndpointName) > 0 {
		input.DBProxyEndpointName = aws.String(_rdsDBProxyEndpointName)
	}
	if len(_rdsNewDBProxyEndpointName) > 0 {
		input.NewDBProxyEndpointName = aws.String(_rdsNewDBProxyEndpointName)
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyDBProxyEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the properties of a DBProxyTargetGroup .
func rds_ModifyDBProxyTargetGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBProxyTargetGroupInput{
		// DBProxyName: *string, // Required
		// TargetGroupName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsTargetGroupName) > 0 {
		input.TargetGroupName = aws.String(_rdsTargetGroupName)
	}
	if len(_rdsConnectionPoolConfig) > 0 {
		if err := assignInputField(input, "ConnectionPoolConfig", _rdsConnectionPoolConfig); err != nil {
			log.Errorf("invalid --connection-pool-config: %s", err.Error())
			return
		}
	}
	if len(_rdsNewName) > 0 {
		input.NewName = aws.String(_rdsNewName)
	}

	if resp, err := client.ModifyDBProxyTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the recommendation status and recommended action status for the
// specified recommendation.
func rds_ModifyDBRecommendation(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBRecommendationInput{
		// RecommendationId: *string, // Required
	}

	if len(_rdsRecommendationId) > 0 {
		input.RecommendationId = aws.String(_rdsRecommendationId)
	}
	if len(_rdsLocale) > 0 {
		input.Locale = aws.String(_rdsLocale)
	}
	if len(_rdsRecommendedActionUpdates) > 0 {
		if err := assignInputField(input, "RecommendedActionUpdates", _rdsRecommendedActionUpdates); err != nil {
			log.Errorf("invalid --recommended-action-updates: %s", err.Error())
			return
		}
	}
	if len(_rdsStatus) > 0 {
		input.Status = aws.String(_rdsStatus)
	}

	if resp, err := client.ModifyDBRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings of an Aurora Limitless Database DB shard group. You can
// change one or more settings by specifying these parameters and the new values in
// the request.
func rds_ModifyDBShardGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBShardGroupInput{
		// DBShardGroupIdentifier: *string, // Required
	}

	if len(_rdsDBShardGroupIdentifier) > 0 {
		input.DBShardGroupIdentifier = aws.String(_rdsDBShardGroupIdentifier)
	}
	if len(_rdsComputeRedundancy) > 0 {
		if err := assignInputField(input, "ComputeRedundancy", _rdsComputeRedundancy); err != nil {
			log.Errorf("invalid --compute-redundancy: %s", err.Error())
			return
		}
	}
	if len(_rdsMaxACU) > 0 {
		if err := assignInputField(input, "MaxACU", _rdsMaxACU); err != nil {
			log.Errorf("invalid --max-acu: %s", err.Error())
			return
		}
	}
	if len(_rdsMinACU) > 0 {
		if err := assignInputField(input, "MinACU", _rdsMinACU); err != nil {
			log.Errorf("invalid --min-acu: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDBShardGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a manual DB snapshot with a new engine version. The snapshot can be
// encrypted or unencrypted, but not shared or public.
//
// Amazon RDS supports upgrading DB snapshots for MariaDB, MySQL, PostgreSQL, and
// Oracle. This operation doesn't apply to RDS Custom or RDS for Db2.
func rds_ModifyDBSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBSnapshotInput{
		// DBSnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}

	if resp, err := client.ModifyDBSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB snapshot.
//
// To share a manual DB snapshot with other Amazon Web Services accounts, specify
// restore as the AttributeName and use the ValuesToAdd parameter to add a list of
// IDs of the Amazon Web Services accounts that are authorized to restore the
// manual DB snapshot. Uses the value all to make the manual DB snapshot public,
// which means it can be copied or restored by all Amazon Web Services accounts.
//
// Don't add the all value for any manual DB snapshots that contain private
// information that you don't want available to all Amazon Web Services accounts.
//
// If the manual DB snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized Amazon Web Services account IDs for the
// ValuesToAdd parameter. You can't use all as a value for that parameter in this
// case.
//
// To view which Amazon Web Services accounts have access to copy or restore a
// manual DB snapshot, or whether a manual DB snapshot public or private, use the DescribeDBSnapshotAttributes
// API operation. The accounts are returned as values for the restore attribute.
func rds_ModifyDBSnapshotAttribute(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBSnapshotAttributeInput{
		// AttributeName: *string, // Required
		// DBSnapshotIdentifier: *string, // Required
	}

	if len(_rdsAttributeName) > 0 {
		input.AttributeName = aws.String(_rdsAttributeName)
	}
	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}
	if len(_rdsValuesToAdd) > 0 {
		input.ValuesToAdd = append([]string(nil), _rdsValuesToAdd...)
	}
	if len(_rdsValuesToRemove) > 0 {
		input.ValuesToRemove = append([]string(nil), _rdsValuesToRemove...)
	}

	if resp, err := client.ModifyDBSnapshotAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing DB subnet group. DB subnet groups must contain at least
// one subnet in at least two AZs in the Amazon Web Services Region.
func rds_ModifyDBSubnetGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyDBSubnetGroupInput{
		// DBSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _rdsSubnetIds...)
	}
	if len(_rdsDBSubnetGroupDescription) > 0 {
		input.DBSubnetGroupDescription = aws.String(_rdsDBSubnetGroupDescription)
	}

	if resp, err := client.ModifyDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing RDS event notification subscription. You can't modify the
// source identifiers using this call. To change source identifiers for a
// subscription, use the AddSourceIdentifierToSubscription and
// RemoveSourceIdentifierFromSubscription calls.
//
// You can see a list of the event categories for a given source type ( SourceType
// ) in [Events]in the Amazon RDS User Guide or by using the DescribeEventCategories
// operation.
//
// [Events]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html
func rds_ModifyEventSubscription(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_rdsSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_rdsSubscriptionName)
	}
	if len(_rdsEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _rdsEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_rdsEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _rdsEventCategories...)
	}
	if len(_rdsSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_rdsSnsTopicArn)
	}
	if len(_rdsSourceType) > 0 {
		input.SourceType = aws.String(_rdsSourceType)
	}

	if resp, err := client.ModifyEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a setting for an Amazon Aurora global database cluster. You can change
// one or more database configuration parameters by specifying these parameters and
// the new values in the request. For more information on Amazon Aurora, see [What is Amazon Aurora?]in
// the Amazon Aurora User Guide.
//
// This operation only applies to Aurora global database clusters.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
func rds_ModifyGlobalCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}
	if len(_rdsAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _rdsAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsNewGlobalClusterIdentifier) > 0 {
		input.NewGlobalClusterIdentifier = aws.String(_rdsNewGlobalClusterIdentifier)
	}

	if resp, err := client.ModifyGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a zero-ETL integration with Amazon Redshift.
func rds_ModifyIntegration(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyIntegrationInput{
		// IntegrationIdentifier: *string, // Required
	}

	if len(_rdsIntegrationIdentifier) > 0 {
		input.IntegrationIdentifier = aws.String(_rdsIntegrationIdentifier)
	}
	if len(_rdsDataFilter) > 0 {
		input.DataFilter = aws.String(_rdsDataFilter)
	}
	if len(_rdsDescription) > 0 {
		input.Description = aws.String(_rdsDescription)
	}
	if len(_rdsIntegrationName) > 0 {
		input.IntegrationName = aws.String(_rdsIntegrationName)
	}

	if resp, err := client.ModifyIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing option group.
func rds_ModifyOptionGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyOptionGroupInput{
		// OptionGroupName: *string, // Required
	}

	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _rdsApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_rdsOptionsToInclude) > 0 {
		if err := assignInputField(input, "OptionsToInclude", _rdsOptionsToInclude); err != nil {
			log.Errorf("invalid --options-to-include: %s", err.Error())
			return
		}
	}
	if len(_rdsOptionsToRemove) > 0 {
		input.OptionsToRemove = append([]string(nil), _rdsOptionsToRemove...)
	}

	if resp, err := client.ModifyOptionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing tenant database in a DB instance. You can change the
// tenant database name or the master user password. This operation is supported
// only for RDS for Oracle CDB instances using the multi-tenant configuration.
func rds_ModifyTenantDatabase(cfg aws.Config, client *rds.Client) {
	input := &rds.ModifyTenantDatabaseInput{
		// DBInstanceIdentifier: *string, // Required
		// TenantDBName: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsTenantDBName) > 0 {
		input.TenantDBName = aws.String(_rdsTenantDBName)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsNewTenantDBName) > 0 {
		input.NewTenantDBName = aws.String(_rdsNewTenantDBName)
	}
	if len(_rdsRotateMasterUserPassword) > 0 {
		if err := assignInputField(input, "RotateMasterUserPassword", _rdsRotateMasterUserPassword); err != nil {
			log.Errorf("invalid --rotate-master-user-password: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyTenantDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Promotes a read replica DB instance to a standalone DB instance.
// - Backup duration is a function of the amount of changes to the database
// since the previous backup. If you plan to promote a read replica to a standalone
// instance, we recommend that you enable backups and complete at least one backup
// prior to promotion. In addition, a read replica cannot be promoted to a
// standalone instance when it is in the backing-up status. If you have enabled
// backups on your read replica, configure the automated backup window so that
// daily backups do not interfere with read replica promotion.
//
// - This command doesn't apply to Aurora MySQL, Aurora PostgreSQL, or RDS
// Custom.
func rds_PromoteReadReplica(cfg aws.Config, client *rds.Client) {
	input := &rds.PromoteReadReplicaInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}

	if resp, err := client.PromoteReadReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Promotes a read replica DB cluster to a standalone DB cluster.
func rds_PromoteReadReplicaDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.PromoteReadReplicaDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}

	if resp, err := client.PromoteReadReplicaDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Purchases a reserved DB instance offering.
func rds_PurchaseReservedDBInstancesOffering(cfg aws.Config, client *rds.Client) {
	input := &rds.PurchaseReservedDBInstancesOfferingInput{
		// ReservedDBInstancesOfferingId: *string, // Required
	}

	if len(_rdsReservedDBInstancesOfferingId) > 0 {
		input.ReservedDBInstancesOfferingId = aws.String(_rdsReservedDBInstancesOfferingId)
	}
	if len(_rdsDBInstanceCount) > 0 {
		if err := assignInputField(input, "DBInstanceCount", _rdsDBInstanceCount); err != nil {
			log.Errorf("invalid --db-instance-count: %s", err.Error())
			return
		}
	}
	if len(_rdsReservedDBInstanceId) > 0 {
		input.ReservedDBInstanceId = aws.String(_rdsReservedDBInstanceId)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseReservedDBInstancesOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You might need to reboot your DB cluster, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB cluster
// parameter group associated with the DB cluster, reboot the DB cluster for the
// changes to take effect.
//
// Rebooting a DB cluster restarts the database engine service. Rebooting a DB
// cluster results in a momentary outage, during which the DB cluster status is set
// to rebooting.
//
// Use this operation only for a non-Aurora Multi-AZ DB cluster.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_RebootDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.RebootDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}

	if resp, err := client.RebootDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You might need to reboot your DB instance, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB parameter
// group associated with the DB instance, you must reboot the instance for the
// changes to take effect.
//
// Rebooting a DB instance restarts the database engine service. Rebooting a DB
// instance results in a momentary outage, during which the DB instance status is
// set to rebooting.
//
// For more information about rebooting, see [Rebooting a DB Instance] in the Amazon RDS User Guide.
//
// This command doesn't apply to RDS Custom.
//
// If your DB instance is part of a Multi-AZ DB cluster, you can reboot the DB
// cluster with the RebootDBCluster operation.
//
// [Rebooting a DB Instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RebootInstance.html
func rds_RebootDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.RebootDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsForceFailover) > 0 {
		if err := assignInputField(input, "ForceFailover", _rdsForceFailover); err != nil {
			log.Errorf("invalid --force-failover: %s", err.Error())
			return
		}
	}

	if resp, err := client.RebootDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You might need to reboot your DB shard group, usually for maintenance reasons.
// For example, if you make certain modifications, reboot the DB shard group for
// the changes to take effect.
//
// This operation applies only to Aurora Limitless Database DBb shard groups.
func rds_RebootDBShardGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.RebootDBShardGroupInput{
		// DBShardGroupIdentifier: *string, // Required
	}

	if len(_rdsDBShardGroupIdentifier) > 0 {
		input.DBShardGroupIdentifier = aws.String(_rdsDBShardGroupIdentifier)
	}

	if resp, err := client.RebootDBShardGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate one or more DBProxyTarget data structures with a DBProxyTargetGroup .
func rds_RegisterDBProxyTargets(cfg aws.Config, client *rds.Client) {
	input := &rds.RegisterDBProxyTargetsInput{
		// DBProxyName: *string, // Required
	}

	if len(_rdsDBProxyName) > 0 {
		input.DBProxyName = aws.String(_rdsDBProxyName)
	}
	if len(_rdsDBClusterIdentifiers) > 0 {
		input.DBClusterIdentifiers = append([]string(nil), _rdsDBClusterIdentifiers...)
	}
	if len(_rdsDBInstanceIdentifiers) > 0 {
		input.DBInstanceIdentifiers = append([]string(nil), _rdsDBInstanceIdentifiers...)
	}
	if len(_rdsTargetGroupName) > 0 {
		input.TargetGroupName = aws.String(_rdsTargetGroupName)
	}

	if resp, err := client.RegisterDBProxyTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches an Aurora secondary cluster from an Aurora global database cluster.
// The cluster becomes a standalone cluster with read-write capability instead of
// being read-only and receiving data from a primary cluster in a different Region.
//
// This operation only applies to Aurora DB clusters.
func rds_RemoveFromGlobalCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.RemoveFromGlobalClusterInput{
		// DbClusterIdentifier: *string, // Required
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DbClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}

	if resp, err := client.RemoveFromGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the asssociation of an Amazon Web Services Identity and Access
// Management (IAM) role from a DB cluster.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_RemoveRoleFromDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.RemoveRoleFromDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsRoleArn) > 0 {
		input.RoleArn = aws.String(_rdsRoleArn)
	}
	if len(_rdsFeatureName) > 0 {
		input.FeatureName = aws.String(_rdsFeatureName)
	}

	if resp, err := client.RemoveRoleFromDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Amazon Web Services Identity and Access Management (IAM) role
// from a DB instance.
func rds_RemoveRoleFromDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.RemoveRoleFromDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
		// FeatureName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsFeatureName) > 0 {
		input.FeatureName = aws.String(_rdsFeatureName)
	}
	if len(_rdsRoleArn) > 0 {
		input.RoleArn = aws.String(_rdsRoleArn)
	}

	if resp, err := client.RemoveRoleFromDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a source identifier from an existing RDS event notification
// subscription.
func rds_RemoveSourceIdentifierFromSubscription(cfg aws.Config, client *rds.Client) {
	input := &rds.RemoveSourceIdentifierFromSubscriptionInput{
		// SourceIdentifier: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_rdsSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_rdsSourceIdentifier)
	}
	if len(_rdsSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_rdsSubscriptionName)
	}

	if resp, err := client.RemoveSourceIdentifierFromSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes metadata tags from an Amazon RDS resource.
// For an overview on tagging an Amazon RDS resource, see [Tagging Amazon RDS Resources] in the Amazon RDS User
// Guide or [Tagging Amazon Aurora and Amazon RDS Resources]in the Amazon Aurora User Guide.
//
// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
// [Tagging Amazon Aurora and Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html
func rds_RemoveTagsFromResource(cfg aws.Config, client *rds.Client) {
	input := &rds.RemoveTagsFromResourceInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_rdsResourceName) > 0 {
		input.ResourceName = aws.String(_rdsResourceName)
	}
	if len(_rdsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _rdsTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB cluster parameter group to the default value.
// To reset specific parameters submit a list of the following: ParameterName and
// ApplyMethod . To reset the entire DB cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When resetting the entire group, dynamic parameters are updated immediately and
// static parameters are set to pending-reboot to take effect on the next DB
// instance restart or RebootDBInstance request. You must call RebootDBInstance
// for every DB instance in your DB cluster that you want the updated static
// parameter to apply to.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_ResetDBClusterParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ResetDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _rdsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_rdsResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _rdsResetAllParameters); err != nil {
			log.Errorf("invalid --reset-all-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB parameter group to the engine/system default
// value. To reset specific parameters, provide a list of the following:
// ParameterName and ApplyMethod . To reset the entire DB parameter group, specify
// the DBParameterGroup name and ResetAllParameters parameters. When resetting the
// entire group, dynamic parameters are updated immediately and static parameters
// are set to pending-reboot to take effect on the next DB instance restart or
// RebootDBInstance request.
func rds_ResetDBParameterGroup(cfg aws.Config, client *rds.Client) {
	input := &rds.ResetDBParameterGroupInput{
		// DBParameterGroupName: *string, // Required
	}

	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _rdsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_rdsResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _rdsResetAllParameters); err != nil {
			log.Errorf("invalid --reset-all-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Aurora DB cluster from MySQL data stored in an Amazon S3
// bucket. Amazon RDS must be authorized to access the Amazon S3 bucket and the
// data must be created using the Percona XtraBackup utility as described in [Migrating Data from MySQL by Using an Amazon S3 Bucket]in
// the Amazon Aurora User Guide.
//
// This operation only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance operation to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier . You can create DB instances only after the
// RestoreDBClusterFromS3 operation has completed and the DB cluster is available.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// This operation only applies to Aurora DB clusters. The source DB engine must be
// MySQL.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Migrating Data from MySQL by Using an Amazon S3 Bucket]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Migrating.ExtMySQL.html#AuroraMySQL.Migrating.ExtMySQL.S3
func rds_RestoreDBClusterFromS3(cfg aws.Config, client *rds.Client) {
	input := &rds.RestoreDBClusterFromS3Input{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
		// MasterUsername: *string, // Required
		// S3BucketName: *string, // Required
		// S3IngestionRoleArn: *string, // Required
		// SourceEngine: *string, // Required
		// SourceEngineVersion: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsMasterUsername) > 0 {
		input.MasterUsername = aws.String(_rdsMasterUsername)
	}
	if len(_rdsS3BucketName) > 0 {
		input.S3BucketName = aws.String(_rdsS3BucketName)
	}
	if len(_rdsS3IngestionRoleArn) > 0 {
		input.S3IngestionRoleArn = aws.String(_rdsS3IngestionRoleArn)
	}
	if len(_rdsSourceEngine) > 0 {
		input.SourceEngine = aws.String(_rdsSourceEngine)
	}
	if len(_rdsSourceEngineVersion) > 0 {
		input.SourceEngineVersion = aws.String(_rdsSourceEngineVersion)
	}
	if len(_rdsAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _rdsAvailabilityZones...)
	}
	if len(_rdsBacktrackWindow) > 0 {
		if err := assignInputField(input, "BacktrackWindow", _rdsBacktrackWindow); err != nil {
			log.Errorf("invalid --backtrack-window: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCharacterSetName) > 0 {
		input.CharacterSetName = aws.String(_rdsCharacterSetName)
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDatabaseName) > 0 {
		input.DatabaseName = aws.String(_rdsDatabaseName)
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_rdsPreferredMaintenanceWindow)
	}
	if len(_rdsS3Prefix) > 0 {
		input.S3Prefix = aws.String(_rdsS3Prefix)
	}
	if len(_rdsServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _rdsServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _rdsStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterFromS3(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
// The target DB cluster is created from the source snapshot with a default
// configuration. If you don't specify a security group, the new DB cluster is
// associated with the default security group.
//
// This operation only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance operation to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier . You can create DB instances only after the
// RestoreDBClusterFromSnapshot operation has completed and the DB cluster is
// available.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_RestoreDBClusterFromSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.RestoreDBClusterFromSnapshotInput{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
		// SnapshotIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_rdsSnapshotIdentifier)
	}
	if len(_rdsAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _rdsAvailabilityZones...)
	}
	if len(_rdsBacktrackWindow) > 0 {
		if err := assignInputField(input, "BacktrackWindow", _rdsBacktrackWindow); err != nil {
			log.Errorf("invalid --backtrack-window: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBClusterInstanceClass) > 0 {
		input.DBClusterInstanceClass = aws.String(_rdsDBClusterInstanceClass)
	}
	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDatabaseName) > 0 {
		input.DatabaseName = aws.String(_rdsDatabaseName)
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineMode) > 0 {
		input.EngineMode = aws.String(_rdsEngineMode)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsRdsCustomClusterConfiguration) > 0 {
		if err := assignInputField(input, "RdsCustomClusterConfiguration", _rdsRdsCustomClusterConfiguration); err != nil {
			log.Errorf("invalid --rds-custom-cluster-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _rdsScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _rdsServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group. Unless the RestoreType is set to
// copy-on-write , the restore may occur in a different Availability Zone (AZ) from
// the original DB cluster. The AZ where RDS restores the DB cluster depends on the
// AZs in the specified subnet group.
//
// For Aurora, this operation only restores the DB cluster, not the DB instances
// for that DB cluster. You must invoke the CreateDBInstance operation to create
// DB instances for the restored DB cluster, specifying the identifier of the
// restored DB cluster in DBClusterIdentifier . You can create DB instances only
// after the RestoreDBClusterToPointInTime operation has completed and the DB
// cluster is available.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func rds_RestoreDBClusterToPointInTime(cfg aws.Config, client *rds.Client) {
	input := &rds.RestoreDBClusterToPointInTimeInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}
	if len(_rdsBacktrackWindow) > 0 {
		if err := assignInputField(input, "BacktrackWindow", _rdsBacktrackWindow); err != nil {
			log.Errorf("invalid --backtrack-window: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBClusterInstanceClass) > 0 {
		input.DBClusterInstanceClass = aws.String(_rdsDBClusterInstanceClass)
	}
	if len(_rdsDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_rdsDBClusterParameterGroupName)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineMode) > 0 {
		input.EngineMode = aws.String(_rdsEngineMode)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsRdsCustomClusterConfiguration) > 0 {
		if err := assignInputField(input, "RdsCustomClusterConfiguration", _rdsRdsCustomClusterConfiguration); err != nil {
			log.Errorf("invalid --rds-custom-cluster-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsRestoreToTime) > 0 {
		if err := assignInputField(input, "RestoreToTime", _rdsRestoreToTime); err != nil {
			log.Errorf("invalid --restore-to-time: %s", err.Error())
			return
		}
	}
	if len(_rdsRestoreType) > 0 {
		input.RestoreType = aws.String(_rdsRestoreType)
	}
	if len(_rdsScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _rdsScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _rdsServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_rdsSourceDBClusterIdentifier)
	}
	if len(_rdsSourceDbClusterResourceId) > 0 {
		input.SourceDbClusterResourceId = aws.String(_rdsSourceDbClusterResourceId)
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsUseLatestRestorableTime) > 0 {
		if err := assignInputField(input, "UseLatestRestorableTime", _rdsUseLatestRestorableTime); err != nil {
			log.Errorf("invalid --use-latest-restorable-time: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterToPointInTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB instance from a DB snapshot. The target database is created
// from the source database restore point with most of the source's original
// configuration, including the default security group and DB parameter group. By
// default, the new DB instance is created as a Single-AZ deployment, except when
// the instance is a SQL Server instance that has an option group associated with
// mirroring. In this case, the instance becomes a Multi-AZ deployment, not a
// Single-AZ deployment.
//
// If you want to replace your original DB instance with the new, restored DB
// instance, then rename your original DB instance before you call the
// RestoreDBInstanceFromDBSnapshot operation. RDS doesn't allow two DB instances
// with the same name. After you have renamed your original DB instance with a
// different identifier, then you can pass the original name of the DB instance as
// the DBInstanceIdentifier in the call to the RestoreDBInstanceFromDBSnapshot
// operation. The result is that you replace the original DB instance with the DB
// instance created from the snapshot.
//
// If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier
// must be the ARN of the shared DB snapshot.
//
// To restore from a DB snapshot with an unsupported engine version, you must
// first upgrade the engine version of the snapshot. For more information about
// upgrading a RDS for MySQL DB snapshot engine version, see [Upgrading a MySQL DB snapshot engine version]. For more
// information about upgrading a RDS for PostgreSQL DB snapshot engine version, [Upgrading a PostgreSQL DB snapshot engine version].
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterFromSnapshot .
//
// [Upgrading a PostgreSQL DB snapshot engine version]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBSnapshot.PostgreSQL.html
// [Upgrading a MySQL DB snapshot engine version]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-upgrade-snapshot.html
func rds_RestoreDBInstanceFromDBSnapshot(cfg aws.Config, client *rds.Client) {
	input := &rds.RestoreDBInstanceFromDBSnapshotInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsAdditionalStorageVolumes) > 0 {
		if err := assignInputField(input, "AdditionalStorageVolumes", _rdsAdditionalStorageVolumes); err != nil {
			log.Errorf("invalid --additional-storage-volumes: %s", err.Error())
			return
		}
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_rdsAvailabilityZone)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupTarget) > 0 {
		input.BackupTarget = aws.String(_rdsBackupTarget)
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsCustomIamInstanceProfile) > 0 {
		input.CustomIamInstanceProfile = aws.String(_rdsCustomIamInstanceProfile)
	}
	if len(_rdsDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_rdsDBClusterSnapshotIdentifier)
	}
	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDBName) > 0 {
		input.DBName = aws.String(_rdsDBName)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDedicatedLogVolume) > 0 {
		if err := assignInputField(input, "DedicatedLogVolume", _rdsDedicatedLogVolume); err != nil {
			log.Errorf("invalid --dedicated-log-volume: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainAuthSecretArn) > 0 {
		input.DomainAuthSecretArn = aws.String(_rdsDomainAuthSecretArn)
	}
	if len(_rdsDomainDnsIps) > 0 {
		input.DomainDnsIps = append([]string(nil), _rdsDomainDnsIps...)
	}
	if len(_rdsDomainFqdn) > 0 {
		input.DomainFqdn = aws.String(_rdsDomainFqdn)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsDomainOu) > 0 {
		input.DomainOu = aws.String(_rdsDomainOu)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableCustomerOwnedIp) > 0 {
		if err := assignInputField(input, "EnableCustomerOwnedIp", _rdsEnableCustomerOwnedIp); err != nil {
			log.Errorf("invalid --enable-customer-owned-ip: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsLicenseModel) > 0 {
		input.LicenseModel = aws.String(_rdsLicenseModel)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsProcessorFeatures) > 0 {
		if err := assignInputField(input, "ProcessorFeatures", _rdsProcessorFeatures); err != nil {
			log.Errorf("invalid --processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageThroughput) > 0 {
		if err := assignInputField(input, "StorageThroughput", _rdsStorageThroughput); err != nil {
			log.Errorf("invalid --storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTdeCredentialArn) > 0 {
		input.TdeCredentialArn = aws.String(_rdsTdeCredentialArn)
	}
	if len(_rdsTdeCredentialPassword) > 0 {
		input.TdeCredentialPassword = aws.String(_rdsTdeCredentialPassword)
	}
	if len(_rdsUseDefaultProcessorFeatures) > 0 {
		if err := assignInputField(input, "UseDefaultProcessorFeatures", _rdsUseDefaultProcessorFeatures); err != nil {
			log.Errorf("invalid --use-default-processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBInstanceFromDBSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Relational Database Service (Amazon RDS) supports importing MySQL
// databases by using backup files. You can create a backup of your on-premises
// database, store it on Amazon Simple Storage Service (Amazon S3), and then
// restore the backup file onto a new Amazon RDS DB instance running MySQL. For
// more information, see [Restoring a backup into an Amazon RDS for MySQL DB instance]in the Amazon RDS User Guide.
//
// This operation doesn't apply to RDS Custom.
//
// [Restoring a backup into an Amazon RDS for MySQL DB instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.html
func rds_RestoreDBInstanceFromS3(cfg aws.Config, client *rds.Client) {
	input := &rds.RestoreDBInstanceFromS3Input{
		// DBInstanceClass: *string, // Required
		// DBInstanceIdentifier: *string, // Required
		// Engine: *string, // Required
		// S3BucketName: *string, // Required
		// S3IngestionRoleArn: *string, // Required
		// SourceEngine: *string, // Required
		// SourceEngineVersion: *string, // Required
	}

	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsS3BucketName) > 0 {
		input.S3BucketName = aws.String(_rdsS3BucketName)
	}
	if len(_rdsS3IngestionRoleArn) > 0 {
		input.S3IngestionRoleArn = aws.String(_rdsS3IngestionRoleArn)
	}
	if len(_rdsSourceEngine) > 0 {
		input.SourceEngine = aws.String(_rdsSourceEngine)
	}
	if len(_rdsSourceEngineVersion) > 0 {
		input.SourceEngineVersion = aws.String(_rdsSourceEngineVersion)
	}
	if len(_rdsAdditionalStorageVolumes) > 0 {
		if err := assignInputField(input, "AdditionalStorageVolumes", _rdsAdditionalStorageVolumes); err != nil {
			log.Errorf("invalid --additional-storage-volumes: %s", err.Error())
			return
		}
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_rdsAvailabilityZone)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsDBName) > 0 {
		input.DBName = aws.String(_rdsDBName)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDBSecurityGroups) > 0 {
		input.DBSecurityGroups = append([]string(nil), _rdsDBSecurityGroups...)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDatabaseInsightsMode) > 0 {
		if err := assignInputField(input, "DatabaseInsightsMode", _rdsDatabaseInsightsMode); err != nil {
			log.Errorf("invalid --database-insights-mode: %s", err.Error())
			return
		}
	}
	if len(_rdsDedicatedLogVolume) > 0 {
		if err := assignInputField(input, "DedicatedLogVolume", _rdsDedicatedLogVolume); err != nil {
			log.Errorf("invalid --dedicated-log-volume: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _rdsEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsEngineVersion) > 0 {
		input.EngineVersion = aws.String(_rdsEngineVersion)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsLicenseModel) > 0 {
		input.LicenseModel = aws.String(_rdsLicenseModel)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_rdsMasterUserPassword)
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMasterUsername) > 0 {
		input.MasterUsername = aws.String(_rdsMasterUsername)
	}
	if len(_rdsMaxAllocatedStorage) > 0 {
		if err := assignInputField(input, "MaxAllocatedStorage", _rdsMaxAllocatedStorage); err != nil {
			log.Errorf("invalid --max-allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _rdsMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_rdsMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_rdsMonitoringRoleArn)
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_rdsPerformanceInsightsKMSKeyId)
	}
	if len(_rdsPerformanceInsightsRetentionPeriod) > 0 {
		if err := assignInputField(input, "PerformanceInsightsRetentionPeriod", _rdsPerformanceInsightsRetentionPeriod); err != nil {
			log.Errorf("invalid --performance-insights-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_rdsPreferredMaintenanceWindow)
	}
	if len(_rdsProcessorFeatures) > 0 {
		if err := assignInputField(input, "ProcessorFeatures", _rdsProcessorFeatures); err != nil {
			log.Errorf("invalid --processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsS3Prefix) > 0 {
		input.S3Prefix = aws.String(_rdsS3Prefix)
	}
	if len(_rdsStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _rdsStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageThroughput) > 0 {
		if err := assignInputField(input, "StorageThroughput", _rdsStorageThroughput); err != nil {
			log.Errorf("invalid --storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsUseDefaultProcessorFeatures) > 0 {
		if err := assignInputField(input, "UseDefaultProcessorFeatures", _rdsUseDefaultProcessorFeatures); err != nil {
			log.Errorf("invalid --use-default-processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBInstanceFromS3(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a DB instance to an arbitrary point in time. You can restore to any
// point in time before the time identified by the LatestRestorableTime property.
// You can restore to a point up to the number of days specified by the
// BackupRetentionPeriod property.
//
// The target database is created with most of the original configuration, but in
// a system-selected Availability Zone, with the default security group, the
// default subnet group, and the default DB parameter group. By default, the new DB
// instance is created as a single-AZ deployment except when the instance is a SQL
// Server instance that has an option group that is associated with mirroring; in
// this case, the instance becomes a mirrored deployment and not a single-AZ
// deployment.
//
// This operation doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterToPointInTime .
func rds_RestoreDBInstanceToPointInTime(cfg aws.Config, client *rds.Client) {
	input := &rds.RestoreDBInstanceToPointInTimeInput{
		// TargetDBInstanceIdentifier: *string, // Required
	}

	if len(_rdsTargetDBInstanceIdentifier) > 0 {
		input.TargetDBInstanceIdentifier = aws.String(_rdsTargetDBInstanceIdentifier)
	}
	if len(_rdsAdditionalStorageVolumes) > 0 {
		if err := assignInputField(input, "AdditionalStorageVolumes", _rdsAdditionalStorageVolumes); err != nil {
			log.Errorf("invalid --additional-storage-volumes: %s", err.Error())
			return
		}
	}
	if len(_rdsAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _rdsAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _rdsAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_rdsAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_rdsAvailabilityZone)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsBackupTarget) > 0 {
		input.BackupTarget = aws.String(_rdsBackupTarget)
	}
	if len(_rdsCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_rdsCACertificateIdentifier)
	}
	if len(_rdsCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _rdsCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_rdsCustomIamInstanceProfile) > 0 {
		input.CustomIamInstanceProfile = aws.String(_rdsCustomIamInstanceProfile)
	}
	if len(_rdsDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_rdsDBInstanceClass)
	}
	if len(_rdsDBName) > 0 {
		input.DBName = aws.String(_rdsDBName)
	}
	if len(_rdsDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_rdsDBParameterGroupName)
	}
	if len(_rdsDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_rdsDBSubnetGroupName)
	}
	if len(_rdsDedicatedLogVolume) > 0 {
		if err := assignInputField(input, "DedicatedLogVolume", _rdsDedicatedLogVolume); err != nil {
			log.Errorf("invalid --dedicated-log-volume: %s", err.Error())
			return
		}
	}
	if len(_rdsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _rdsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_rdsDomain) > 0 {
		input.Domain = aws.String(_rdsDomain)
	}
	if len(_rdsDomainAuthSecretArn) > 0 {
		input.DomainAuthSecretArn = aws.String(_rdsDomainAuthSecretArn)
	}
	if len(_rdsDomainDnsIps) > 0 {
		input.DomainDnsIps = append([]string(nil), _rdsDomainDnsIps...)
	}
	if len(_rdsDomainFqdn) > 0 {
		input.DomainFqdn = aws.String(_rdsDomainFqdn)
	}
	if len(_rdsDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_rdsDomainIAMRoleName)
	}
	if len(_rdsDomainOu) > 0 {
		input.DomainOu = aws.String(_rdsDomainOu)
	}
	if len(_rdsEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _rdsEnableCloudwatchLogsExports...)
	}
	if len(_rdsEnableCustomerOwnedIp) > 0 {
		if err := assignInputField(input, "EnableCustomerOwnedIp", _rdsEnableCustomerOwnedIp); err != nil {
			log.Errorf("invalid --enable-customer-owned-ip: %s", err.Error())
			return
		}
	}
	if len(_rdsEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _rdsEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_rdsEngine) > 0 {
		input.Engine = aws.String(_rdsEngine)
	}
	if len(_rdsEngineLifecycleSupport) > 0 {
		input.EngineLifecycleSupport = aws.String(_rdsEngineLifecycleSupport)
	}
	if len(_rdsIops) > 0 {
		if err := assignInputField(input, "Iops", _rdsIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_rdsLicenseModel) > 0 {
		input.LicenseModel = aws.String(_rdsLicenseModel)
	}
	if len(_rdsManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _rdsManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_rdsMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_rdsMasterUserSecretKmsKeyId)
	}
	if len(_rdsMaxAllocatedStorage) > 0 {
		if err := assignInputField(input, "MaxAllocatedStorage", _rdsMaxAllocatedStorage); err != nil {
			log.Errorf("invalid --max-allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_rdsMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _rdsMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_rdsNetworkType) > 0 {
		input.NetworkType = aws.String(_rdsNetworkType)
	}
	if len(_rdsOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_rdsOptionGroupName)
	}
	if len(_rdsPort) > 0 {
		if err := assignInputField(input, "Port", _rdsPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rdsPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_rdsPreferredBackupWindow)
	}
	if len(_rdsProcessorFeatures) > 0 {
		if err := assignInputField(input, "ProcessorFeatures", _rdsProcessorFeatures); err != nil {
			log.Errorf("invalid --processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _rdsPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_rdsRestoreTime) > 0 {
		if err := assignInputField(input, "RestoreTime", _rdsRestoreTime); err != nil {
			log.Errorf("invalid --restore-time: %s", err.Error())
			return
		}
	}
	if len(_rdsSourceDBInstanceAutomatedBackupsArn) > 0 {
		input.SourceDBInstanceAutomatedBackupsArn = aws.String(_rdsSourceDBInstanceAutomatedBackupsArn)
	}
	if len(_rdsSourceDBInstanceIdentifier) > 0 {
		input.SourceDBInstanceIdentifier = aws.String(_rdsSourceDBInstanceIdentifier)
	}
	if len(_rdsSourceDbiResourceId) > 0 {
		input.SourceDbiResourceId = aws.String(_rdsSourceDbiResourceId)
	}
	if len(_rdsStorageThroughput) > 0 {
		if err := assignInputField(input, "StorageThroughput", _rdsStorageThroughput); err != nil {
			log.Errorf("invalid --storage-throughput: %s", err.Error())
			return
		}
	}
	if len(_rdsStorageType) > 0 {
		input.StorageType = aws.String(_rdsStorageType)
	}
	if len(_rdsTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _rdsTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rdsTdeCredentialArn) > 0 {
		input.TdeCredentialArn = aws.String(_rdsTdeCredentialArn)
	}
	if len(_rdsTdeCredentialPassword) > 0 {
		input.TdeCredentialPassword = aws.String(_rdsTdeCredentialPassword)
	}
	if len(_rdsUseDefaultProcessorFeatures) > 0 {
		if err := assignInputField(input, "UseDefaultProcessorFeatures", _rdsUseDefaultProcessorFeatures); err != nil {
			log.Errorf("invalid --use-default-processor-features: %s", err.Error())
			return
		}
	}
	if len(_rdsUseLatestRestorableTime) > 0 {
		if err := assignInputField(input, "UseLatestRestorableTime", _rdsUseLatestRestorableTime); err != nil {
			log.Errorf("invalid --use-latest-restorable-time: %s", err.Error())
			return
		}
	}
	if len(_rdsVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _rdsVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBInstanceToPointInTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes ingress from a DBSecurityGroup for previously authorized IP ranges or
// EC2 or VPC security groups. Required parameters for this API are one of CIDRIP,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId).
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func rds_RevokeDBSecurityGroupIngress(cfg aws.Config, client *rds.Client) {
	input := &rds.RevokeDBSecurityGroupIngressInput{
		// DBSecurityGroupName: *string, // Required
	}

	if len(_rdsDBSecurityGroupName) > 0 {
		input.DBSecurityGroupName = aws.String(_rdsDBSecurityGroupName)
	}
	if len(_rdsCIDRIP) > 0 {
		input.CIDRIP = aws.String(_rdsCIDRIP)
	}
	if len(_rdsEC2SecurityGroupId) > 0 {
		input.EC2SecurityGroupId = aws.String(_rdsEC2SecurityGroupId)
	}
	if len(_rdsEC2SecurityGroupName) > 0 {
		input.EC2SecurityGroupName = aws.String(_rdsEC2SecurityGroupName)
	}
	if len(_rdsEC2SecurityGroupOwnerId) > 0 {
		input.EC2SecurityGroupOwnerId = aws.String(_rdsEC2SecurityGroupOwnerId)
	}

	if resp, err := client.RevokeDBSecurityGroupIngress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a database activity stream to monitor activity on the database. For more
// information, see [Monitoring Amazon Aurora with Database Activity Streams]in the Amazon Aurora User Guide or [Monitoring Amazon RDS with Database Activity Streams] in the Amazon RDS User
// Guide.
//
// [Monitoring Amazon Aurora with Database Activity Streams]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html
// [Monitoring Amazon RDS with Database Activity Streams]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.html
func rds_StartActivityStream(cfg aws.Config, client *rds.Client) {
	input := &rds.StartActivityStreamInput{
		// KmsKeyId: *string, // Required
		// Mode: types.ActivityStreamMode, // Required
		// ResourceArn: *string, // Required
	}

	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsMode) > 0 {
		if err := assignInputField(input, "Mode", _rdsMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_rdsResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsResourceArn)
	}
	if len(_rdsApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _rdsApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_rdsEngineNativeAuditFieldsIncluded) > 0 {
		if err := assignInputField(input, "EngineNativeAuditFieldsIncluded", _rdsEngineNativeAuditFieldsIncluded); err != nil {
			log.Errorf("invalid --engine-native-audit-fields-included: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartActivityStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an Amazon Aurora DB cluster that was stopped using the Amazon Web
// Services console, the stop-db-cluster CLI command, or the StopDBCluster
// operation.
//
// For more information, see [Stopping and Starting an Aurora Cluster] in the Amazon Aurora User Guide.
//
// This operation only applies to Aurora DB clusters.
//
// [Stopping and Starting an Aurora Cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-cluster-stop-start.html
func rds_StartDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.StartDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}

	if resp, err := client.StartDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an Amazon RDS DB instance that was stopped using the Amazon Web Services
// console, the stop-db-instance CLI command, or the StopDBInstance operation.
//
// For more information, see [Starting an Amazon RDS DB instance That Was Previously Stopped] in the Amazon RDS User Guide.
//
// This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.
// For Aurora DB clusters, use StartDBCluster instead.
//
// [Starting an Amazon RDS DB instance That Was Previously Stopped]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StartInstance.html
func rds_StartDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.StartDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}

	if resp, err := client.StartDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables replication of automated backups to a different Amazon Web Services
// Region.
//
// This command doesn't apply to RDS Custom.
//
// For more information, see [Replicating Automated Backups to Another Amazon Web Services Region] in the Amazon RDS User Guide.
//
// [Replicating Automated Backups to Another Amazon Web Services Region]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReplicateBackups.html
func rds_StartDBInstanceAutomatedBackupsReplication(cfg aws.Config, client *rds.Client) {
	input := &rds.StartDBInstanceAutomatedBackupsReplicationInput{
		// SourceDBInstanceArn: *string, // Required
	}

	if len(_rdsSourceDBInstanceArn) > 0 {
		input.SourceDBInstanceArn = aws.String(_rdsSourceDBInstanceArn)
	}
	if len(_rdsBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _rdsBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_rdsPreSignedUrl)
	}
	if len(_rdsTags) > 0 {
		if err := assignInputField(input, "Tags", _rdsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDBInstanceAutomatedBackupsReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an export of DB snapshot or DB cluster data to Amazon S3. The provided
// IAM role must have access to the S3 bucket.
//
// You can't export snapshot data from RDS Custom DB instances. For more
// information, see [Supported Regions and DB engines for exporting snapshots to S3 in Amazon RDS].
//
// For more information on exporting DB snapshot data, see [Exporting DB snapshot data to Amazon S3] in the Amazon RDS User
// Guide or [Exporting DB cluster snapshot data to Amazon S3]in the Amazon Aurora User Guide.
//
// For more information on exporting DB cluster data, see [Exporting DB cluster data to Amazon S3] in the Amazon Aurora
// User Guide.
//
// [Exporting DB cluster snapshot data to Amazon S3]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-export-snapshot.html
// [Exporting DB cluster data to Amazon S3]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/export-cluster-data.html
// [Exporting DB snapshot data to Amazon S3]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ExportSnapshot.html
// [Supported Regions and DB engines for exporting snapshots to S3 in Amazon RDS]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RDS_Fea_Regions_DB-eng.Feature.ExportSnapshotToS3.html
func rds_StartExportTask(cfg aws.Config, client *rds.Client) {
	input := &rds.StartExportTaskInput{
		// ExportTaskIdentifier: *string, // Required
		// IamRoleArn: *string, // Required
		// KmsKeyId: *string, // Required
		// S3BucketName: *string, // Required
		// SourceArn: *string, // Required
	}

	if len(_rdsExportTaskIdentifier) > 0 {
		input.ExportTaskIdentifier = aws.String(_rdsExportTaskIdentifier)
	}
	if len(_rdsIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_rdsIamRoleArn)
	}
	if len(_rdsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rdsKmsKeyId)
	}
	if len(_rdsS3BucketName) > 0 {
		input.S3BucketName = aws.String(_rdsS3BucketName)
	}
	if len(_rdsSourceArn) > 0 {
		input.SourceArn = aws.String(_rdsSourceArn)
	}
	if len(_rdsExportOnly) > 0 {
		input.ExportOnly = append([]string(nil), _rdsExportOnly...)
	}
	if len(_rdsS3Prefix) > 0 {
		input.S3Prefix = aws.String(_rdsS3Prefix)
	}

	if resp, err := client.StartExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a database activity stream that was started using the Amazon Web Services
// console, the start-activity-stream CLI command, or the StartActivityStream
// operation.
//
// For more information, see [Monitoring Amazon Aurora with Database Activity Streams] in the Amazon Aurora User Guide or [Monitoring Amazon RDS with Database Activity Streams] in the Amazon
// RDS User Guide.
//
// [Monitoring Amazon Aurora with Database Activity Streams]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html
// [Monitoring Amazon RDS with Database Activity Streams]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.html
func rds_StopActivityStream(cfg aws.Config, client *rds.Client) {
	input := &rds.StopActivityStreamInput{
		// ResourceArn: *string, // Required
	}

	if len(_rdsResourceArn) > 0 {
		input.ResourceArn = aws.String(_rdsResourceArn)
	}
	if len(_rdsApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _rdsApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopActivityStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an Amazon Aurora DB cluster. When you stop a DB cluster, Aurora retains
// the DB cluster's metadata, including its endpoints and DB parameter groups.
// Aurora also retains the transaction logs so you can do a point-in-time restore
// if necessary.
//
// For more information, see [Stopping and Starting an Aurora Cluster] in the Amazon Aurora User Guide.
//
// This operation only applies to Aurora DB clusters.
//
// [Stopping and Starting an Aurora Cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-cluster-stop-start.html
func rds_StopDBCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.StopDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_rdsDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_rdsDBClusterIdentifier)
	}

	if resp, err := client.StopDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an Amazon RDS DB instance temporarily. When you stop a DB instance,
// Amazon RDS retains the DB instance's metadata, including its endpoint, DB
// parameter group, and option group membership. Amazon RDS also retains the
// transaction logs so you can do a point-in-time restore if necessary. The
// instance restarts automatically after 7 days.
//
// For more information, see [Stopping an Amazon RDS DB Instance Temporarily] in the Amazon RDS User Guide.
//
// This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.
// For Aurora clusters, use StopDBCluster instead.
//
// [Stopping an Amazon RDS DB Instance Temporarily]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StopInstance.html
func rds_StopDBInstance(cfg aws.Config, client *rds.Client) {
	input := &rds.StopDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}
	if len(_rdsDBSnapshotIdentifier) > 0 {
		input.DBSnapshotIdentifier = aws.String(_rdsDBSnapshotIdentifier)
	}

	if resp, err := client.StopDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops automated backup replication for a DB instance.
// This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.
//
// For more information, see [Replicating Automated Backups to Another Amazon Web Services Region] in the Amazon RDS User Guide.
//
// [Replicating Automated Backups to Another Amazon Web Services Region]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReplicateBackups.html
func rds_StopDBInstanceAutomatedBackupsReplication(cfg aws.Config, client *rds.Client) {
	input := &rds.StopDBInstanceAutomatedBackupsReplicationInput{
		// SourceDBInstanceArn: *string, // Required
	}

	if len(_rdsSourceDBInstanceArn) > 0 {
		input.SourceDBInstanceArn = aws.String(_rdsSourceDBInstanceArn)
	}

	if resp, err := client.StopDBInstanceAutomatedBackupsReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Switches over a blue/green deployment.
// Before you switch over, production traffic is routed to the databases in the
// blue environment. After you switch over, production traffic is routed to the
// databases in the green environment.
//
// For more information, see [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon RDS User Guide and [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon
// Aurora User Guide.
//
// [Using Amazon RDS Blue/Green Deployments for database updates]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html
func rds_SwitchoverBlueGreenDeployment(cfg aws.Config, client *rds.Client) {
	input := &rds.SwitchoverBlueGreenDeploymentInput{
		// BlueGreenDeploymentIdentifier: *string, // Required
	}

	if len(_rdsBlueGreenDeploymentIdentifier) > 0 {
		input.BlueGreenDeploymentIdentifier = aws.String(_rdsBlueGreenDeploymentIdentifier)
	}
	if len(_rdsSwitchoverTimeout) > 0 {
		if err := assignInputField(input, "SwitchoverTimeout", _rdsSwitchoverTimeout); err != nil {
			log.Errorf("invalid --switchover-timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.SwitchoverBlueGreenDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Switches over the specified secondary DB cluster to be the new primary DB
// cluster in the global database cluster. Switchover operations were previously
// called "managed planned failovers."
//
// Aurora promotes the specified secondary cluster to assume full read/write
// capabilities and demotes the current primary cluster to a secondary (read-only)
// cluster, maintaining the orginal replication topology. All secondary clusters
// are synchronized with the primary at the beginning of the process so the new
// primary continues operations for the Aurora global database without losing any
// data. Your database is unavailable for a short time while the primary and
// selected secondary clusters are assuming their new roles. For more information
// about switching over an Aurora global database, see [Performing switchovers for Amazon Aurora global databases]in the Amazon Aurora User
// Guide.
//
// This operation is intended for controlled environments, for operations such as
// "regional rotation" or to fall back to the original primary after a global
// database failover.
//
// [Performing switchovers for Amazon Aurora global databases]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-disaster-recovery.managed-failover
func rds_SwitchoverGlobalCluster(cfg aws.Config, client *rds.Client) {
	input := &rds.SwitchoverGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
		// TargetDbClusterIdentifier: *string, // Required
	}

	if len(_rdsGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_rdsGlobalClusterIdentifier)
	}
	if len(_rdsTargetDbClusterIdentifier) > 0 {
		input.TargetDbClusterIdentifier = aws.String(_rdsTargetDbClusterIdentifier)
	}

	if resp, err := client.SwitchoverGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Switches over an Oracle standby database in an Oracle Data Guard environment,
// making it the new primary database. Issue this command in the Region that hosts
// the current standby database.
func rds_SwitchoverReadReplica(cfg aws.Config, client *rds.Client) {
	input := &rds.SwitchoverReadReplicaInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_rdsDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_rdsDBInstanceIdentifier)
	}

	if resp, err := client.SwitchoverReadReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rdsCmd)
	_rdsCmd.Flags().SortFlags = false

	_rdsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_rdsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rdsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_rdsCmd.Flags().StringVarP(&_rdsAdditionalEncryptionContext, "additional-encryption-context", "", "", "Additional Encryption Context")
	_rdsCmd.Flags().StringVarP(&_rdsAdditionalStorageVolumes, "additional-storage-volumes", "", "", "Additional Storage Volumes")
	_rdsCmd.Flags().StringVarP(&_rdsAllocatedStorage, "allocated-storage", "", "", "Allocated Storage")
	_rdsCmd.Flags().StringVarP(&_rdsAllowDataLoss, "allow-data-loss", "", "", "Allow Data Loss")
	_rdsCmd.Flags().StringVarP(&_rdsAllowEngineModeChange, "allow-engine-mode-change", "", "", "Allow Engine Mode Change")
	_rdsCmd.Flags().StringVarP(&_rdsAllowMajorVersionUpgrade, "allow-major-version-upgrade", "", "", "Allow Major Version Upgrade")
	_rdsCmd.Flags().StringVarP(&_rdsApplyAction, "apply-action", "", "", "Apply Action")
	_rdsCmd.Flags().StringVarP(&_rdsApplyImmediately, "apply-immediately", "", "", "Apply Immediately")
	_rdsCmd.Flags().StringVarP(&_rdsAttributeName, "attribute-name", "", "", "Attribute Name")
	_rdsCmd.Flags().StringVarP(&_rdsAuditPolicyState, "audit-policy-state", "", "", "Audit Policy State")
	_rdsCmd.Flags().StringVarP(&_rdsAuth, "auth", "", "", "Auth")
	_rdsCmd.Flags().StringVarP(&_rdsAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_rdsCmd.Flags().StringVarP(&_rdsAutomationMode, "automation-mode", "", "", "Automation Mode")
	_rdsCmd.Flags().StringVarP(&_rdsAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_rdsCmd.Flags().StringVarP(&_rdsAvailabilityZoneGroup, "availability-zone-group", "", "", "Availability Zone Group")
	_rdsCmd.Flags().StringSliceVarP(&_rdsAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_rdsCmd.Flags().StringVarP(&_rdsAwsBackupRecoveryPointArn, "aws-backup-recovery-point-arn", "", "", "AWS Backup Recovery Point ARN")
	_rdsCmd.Flags().StringVarP(&_rdsBacktrackIdentifier, "backtrack-identifier", "", "", "Backtrack Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsBacktrackTo, "backtrack-to", "", "", "Backtrack To")
	_rdsCmd.Flags().StringVarP(&_rdsBacktrackWindow, "backtrack-window", "", "", "Backtrack Window")
	_rdsCmd.Flags().StringVarP(&_rdsBackupRetentionPeriod, "backup-retention-period", "", "", "Backup Retention Period")
	_rdsCmd.Flags().StringVarP(&_rdsBackupTarget, "backup-target", "", "", "Backup Target")
	_rdsCmd.Flags().StringVarP(&_rdsBlueGreenDeploymentIdentifier, "blue-green-deployment-identifier", "", "", "Blue Green Deployment Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsBlueGreenDeploymentName, "blue-green-deployment-name", "", "", "Blue Green Deployment Name")
	_rdsCmd.Flags().StringVarP(&_rdsCACertificateIdentifier, "ca-certificate-identifier", "", "", "Ca Certificate Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsCapacity, "capacity", "", "", "Capacity")
	_rdsCmd.Flags().StringVarP(&_rdsCertificateIdentifier, "certificate-identifier", "", "", "Certificate Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsCertificateRotationRestart, "certificate-rotation-restart", "", "", "Certificate Rotation Restart")
	_rdsCmd.Flags().StringVarP(&_rdsCharacterSetName, "character-set-name", "", "", "Character Set Name")
	_rdsCmd.Flags().StringVarP(&_rdsCIDRIP, "cidrip", "", "", "Cidrip")
	_rdsCmd.Flags().StringVarP(&_rdsCloudwatchLogsExportConfiguration, "cloudwatch-logs-export-configuration", "", "", "Cloudwatch Logs Export Configuration")
	_rdsCmd.Flags().StringVarP(&_rdsClusterScalabilityType, "cluster-scalability-type", "", "", "Cluster Scalability Type")
	_rdsCmd.Flags().StringVarP(&_rdsComputeRedundancy, "compute-redundancy", "", "", "Compute Redundancy")
	_rdsCmd.Flags().StringVarP(&_rdsConnectionPoolConfig, "connection-pool-config", "", "", "Connection Pool Config")
	_rdsCmd.Flags().StringVarP(&_rdsCopyTags, "copy-tags", "", "", "Copy Tags")
	_rdsCmd.Flags().StringVarP(&_rdsCopyTagsToSnapshot, "copy-tags-to-snapshot", "", "", "Copy Tags To Snapshot")
	_rdsCmd.Flags().StringVarP(&_rdsCustomIamInstanceProfile, "custom-iam-instance-profile", "", "", "Custom IAM Instance Profile")
	_rdsCmd.Flags().StringVarP(&_rdsDataFilter, "data-filter", "", "", "Data Filter")
	_rdsCmd.Flags().StringVarP(&_rdsDatabaseInsightsMode, "database-insights-mode", "", "", "Database Insights Mode")
	_rdsCmd.Flags().StringSliceVarP(&_rdsDatabaseInstallationFiles, "database-installation-files", "", nil, "Database Installation Files")
	_rdsCmd.Flags().StringVarP(&_rdsDatabaseInstallationFilesS3BucketName, "database-installation-files-s3-bucket-name", "", "", "Database Installation Files S3 Bucket Name")
	_rdsCmd.Flags().StringVarP(&_rdsDatabaseInstallationFilesS3Prefix, "database-installation-files-s3-prefix", "", "", "Database Installation Files S3 Prefix")
	_rdsCmd.Flags().StringVarP(&_rdsDatabaseName, "database-name", "", "", "Database Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBClusterEndpointIdentifier, "db-cluster-endpoint-identifier", "", "", "DB Cluster Endpoint Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsDBClusterIdentifier, "db-cluster-identifier", "", "", "DB Cluster Identifier")
	_rdsCmd.Flags().StringSliceVarP(&_rdsDBClusterIdentifiers, "db-cluster-identifiers", "", nil, "DB Cluster Identifiers")
	_rdsCmd.Flags().StringVarP(&_rdsDBClusterInstanceClass, "db-cluster-instance-class", "", "", "DB Cluster Instance Class")
	_rdsCmd.Flags().StringVarP(&_rdsDBClusterParameterGroupName, "db-cluster-parameter-group-name", "", "", "DB Cluster Parameter Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsDbClusterResourceId, "db-cluster-resource-id", "", "", "DB Cluster Resource ID")
	_rdsCmd.Flags().StringVarP(&_rdsDBClusterSnapshotIdentifier, "db-cluster-snapshot-identifier", "", "", "DB Cluster Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsDBInstanceAutomatedBackupsArn, "db-instance-automated-backups-arn", "", "", "DB Instance Automated Backups ARN")
	_rdsCmd.Flags().StringVarP(&_rdsDBInstanceClass, "db-instance-class", "", "", "DB Instance Class")
	_rdsCmd.Flags().StringVarP(&_rdsDBInstanceCount, "db-instance-count", "", "", "DB Instance Count")
	_rdsCmd.Flags().StringVarP(&_rdsDBInstanceIdentifier, "db-instance-identifier", "", "", "DB Instance Identifier")
	_rdsCmd.Flags().StringSliceVarP(&_rdsDBInstanceIdentifiers, "db-instance-identifiers", "", nil, "DB Instance Identifiers")
	_rdsCmd.Flags().StringVarP(&_rdsDBInstanceParameterGroupName, "db-instance-parameter-group-name", "", "", "DB Instance Parameter Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBName, "db-name", "", "", "DB Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBParameterGroupFamily, "db-parameter-group-family", "", "", "DB Parameter Group Family")
	_rdsCmd.Flags().StringVarP(&_rdsDBParameterGroupName, "db-parameter-group-name", "", "", "DB Parameter Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBPortNumber, "db-port-number", "", "", "DB Port Number")
	_rdsCmd.Flags().StringVarP(&_rdsDBProxyEndpointName, "db-proxy-endpoint-name", "", "", "DB Proxy Endpoint Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBProxyName, "db-proxy-name", "", "", "DB Proxy Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBSecurityGroupDescription, "db-security-group-description", "", "", "DB Security Group Description")
	_rdsCmd.Flags().StringVarP(&_rdsDBSecurityGroupName, "db-security-group-name", "", "", "DB Security Group Name")
	_rdsCmd.Flags().StringSliceVarP(&_rdsDBSecurityGroups, "db-security-groups", "", nil, "DB Security Groups")
	_rdsCmd.Flags().StringVarP(&_rdsDBShardGroupIdentifier, "db-shard-group-identifier", "", "", "DB Shard Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsDBSnapshotIdentifier, "db-snapshot-identifier", "", "", "DB Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsDBSubnetGroupDescription, "db-subnet-group-description", "", "", "DB Subnet Group Description")
	_rdsCmd.Flags().StringVarP(&_rdsDBSubnetGroupName, "db-subnet-group-name", "", "", "DB Subnet Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsDBSystemId, "db-system-id", "", "", "DB System ID")
	_rdsCmd.Flags().StringVarP(&_rdsDbiResourceId, "dbi-resource-id", "", "", "Dbi Resource ID")
	_rdsCmd.Flags().StringVarP(&_rdsDebugLogging, "debug-logging", "", "", "Debug Logging")
	_rdsCmd.Flags().StringVarP(&_rdsDedicatedLogVolume, "dedicated-log-volume", "", "", "Dedicated Log Volume")
	_rdsCmd.Flags().StringVarP(&_rdsDefaultAuthScheme, "default-auth-scheme", "", "", "Default Auth Scheme")
	_rdsCmd.Flags().StringVarP(&_rdsDefaultOnly, "default-only", "", "", "Default Only")
	_rdsCmd.Flags().StringVarP(&_rdsDeleteAutomatedBackups, "delete-automated-backups", "", "", "Delete Automated Backups")
	_rdsCmd.Flags().StringVarP(&_rdsDeleteTarget, "delete-target", "", "", "Delete Target")
	_rdsCmd.Flags().StringVarP(&_rdsDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_rdsCmd.Flags().StringVarP(&_rdsDescription, "description", "", "", "Description")
	_rdsCmd.Flags().StringVarP(&_rdsDisableDomain, "disable-domain", "", "", "Disable Domain")
	_rdsCmd.Flags().StringVarP(&_rdsDomain, "domain", "", "", "Domain")
	_rdsCmd.Flags().StringVarP(&_rdsDomainAuthSecretArn, "domain-auth-secret-arn", "", "", "Domain Auth Secret ARN")
	_rdsCmd.Flags().StringSliceVarP(&_rdsDomainDnsIps, "domain-dns-ips", "", nil, "Domain DNS Ips")
	_rdsCmd.Flags().StringVarP(&_rdsDomainFqdn, "domain-fqdn", "", "", "Domain Fqdn")
	_rdsCmd.Flags().StringVarP(&_rdsDomainIAMRoleName, "domain-iam-role-name", "", "", "Domain IAM Role Name")
	_rdsCmd.Flags().StringVarP(&_rdsDomainOu, "domain-ou", "", "", "Domain Ou")
	_rdsCmd.Flags().StringVarP(&_rdsDuration, "duration", "", "", "Duration")
	_rdsCmd.Flags().StringVarP(&_rdsEC2SecurityGroupId, "ec2-security-group-id", "", "", "EC2 Security Group ID")
	_rdsCmd.Flags().StringVarP(&_rdsEC2SecurityGroupName, "ec2-security-group-name", "", "", "EC2 Security Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsEC2SecurityGroupOwnerId, "ec2-security-group-owner-id", "", "", "EC2 Security Group Owner ID")
	_rdsCmd.Flags().StringSliceVarP(&_rdsEnableCloudwatchLogsExports, "enable-cloudwatch-logs-exports", "", nil, "Enable Cloudwatch Logs Exports")
	_rdsCmd.Flags().StringVarP(&_rdsEnableCustomerOwnedIp, "enable-customer-owned-ip", "", "", "Enable Customer Owned IP")
	_rdsCmd.Flags().StringVarP(&_rdsEnableGlobalWriteForwarding, "enable-global-write-forwarding", "", "", "Enable Global Write Forwarding")
	_rdsCmd.Flags().StringVarP(&_rdsEnableIAMDatabaseAuthentication, "enable-iam-database-authentication", "", "", "Enable IAM Database Authentication")
	_rdsCmd.Flags().StringVarP(&_rdsEnableLimitlessDatabase, "enable-limitless-database", "", "", "Enable Limitless Database")
	_rdsCmd.Flags().StringVarP(&_rdsEnableLocalWriteForwarding, "enable-local-write-forwarding", "", "", "Enable Local Write Forwarding")
	_rdsCmd.Flags().StringVarP(&_rdsEnablePerformanceInsights, "enable-performance-insights", "", "", "Enable Performance Insights")
	_rdsCmd.Flags().StringVarP(&_rdsEnabled, "enabled", "", "", "Enabled")
	_rdsCmd.Flags().StringVarP(&_rdsEndTime, "end-time", "", "", "End Time")
	_rdsCmd.Flags().StringVarP(&_rdsEndpointNetworkType, "endpoint-network-type", "", "", "Endpoint Network Type")
	_rdsCmd.Flags().StringVarP(&_rdsEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_rdsCmd.Flags().StringVarP(&_rdsEngine, "engine", "", "", "Engine")
	_rdsCmd.Flags().StringVarP(&_rdsEngineFamily, "engine-family", "", "", "Engine Family")
	_rdsCmd.Flags().StringVarP(&_rdsEngineLifecycleSupport, "engine-lifecycle-support", "", "", "Engine Lifecycle Support")
	_rdsCmd.Flags().StringVarP(&_rdsEngineMode, "engine-mode", "", "", "Engine Mode")
	_rdsCmd.Flags().StringVarP(&_rdsEngineName, "engine-name", "", "", "Engine Name")
	_rdsCmd.Flags().StringVarP(&_rdsEngineNativeAuditFieldsIncluded, "engine-native-audit-fields-included", "", "", "Engine Native Audit Fields Included")
	_rdsCmd.Flags().StringVarP(&_rdsEngineVersion, "engine-version", "", "", "Engine Version")
	_rdsCmd.Flags().StringSliceVarP(&_rdsEventCategories, "event-categories", "", nil, "Event Categories")
	_rdsCmd.Flags().StringSliceVarP(&_rdsExcludedMembers, "excluded-members", "", nil, "Excluded Members")
	_rdsCmd.Flags().StringSliceVarP(&_rdsExportOnly, "export-only", "", nil, "Export Only")
	_rdsCmd.Flags().StringVarP(&_rdsExportTaskIdentifier, "export-task-identifier", "", "", "Export Task Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsFeatureName, "feature-name", "", "", "Feature Name")
	_rdsCmd.Flags().StringVarP(&_rdsFileLastWritten, "file-last-written", "", "", "File Last Written")
	_rdsCmd.Flags().StringVarP(&_rdsFileSize, "file-size", "", "", "File Size")
	_rdsCmd.Flags().StringVarP(&_rdsFilenameContains, "filename-contains", "", "", "Filename Contains")
	_rdsCmd.Flags().StringVarP(&_rdsFilters, "filters", "", "", "Filters")
	_rdsCmd.Flags().StringVarP(&_rdsFinalDBSnapshotIdentifier, "final-db-snapshot-identifier", "", "", "Final DB Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsForce, "force", "", "", "Force")
	_rdsCmd.Flags().StringVarP(&_rdsForceFailover, "force-failover", "", "", "Force Failover")
	_rdsCmd.Flags().StringVarP(&_rdsGlobalClusterIdentifier, "global-cluster-identifier", "", "", "Global Cluster Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_rdsCmd.Flags().StringVarP(&_rdsIdleClientTimeout, "idle-client-timeout", "", "", "Idle Client Timeout")
	_rdsCmd.Flags().StringVarP(&_rdsImageId, "image-id", "", "", "Image ID")
	_rdsCmd.Flags().StringVarP(&_rdsIncludeAll, "include-all", "", "", "Include All")
	_rdsCmd.Flags().StringVarP(&_rdsIncludePublic, "include-public", "", "", "Include Public")
	_rdsCmd.Flags().StringVarP(&_rdsIncludeShared, "include-shared", "", "", "Include Shared")
	_rdsCmd.Flags().StringVarP(&_rdsIntegrationIdentifier, "integration-identifier", "", "", "Integration Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsIntegrationName, "integration-name", "", "", "Integration Name")
	_rdsCmd.Flags().StringVarP(&_rdsIops, "iops", "", "", "IOPS")
	_rdsCmd.Flags().StringVarP(&_rdsKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_rdsCmd.Flags().StringVarP(&_rdsLastUpdatedAfter, "last-updated-after", "", "", "Last Updated After")
	_rdsCmd.Flags().StringVarP(&_rdsLastUpdatedBefore, "last-updated-before", "", "", "Last Updated Before")
	_rdsCmd.Flags().StringVarP(&_rdsLeaseId, "lease-id", "", "", "Lease ID")
	_rdsCmd.Flags().StringVarP(&_rdsLicenseModel, "license-model", "", "", "License Model")
	_rdsCmd.Flags().StringVarP(&_rdsListSupportedCharacterSets, "list-supported-character-sets", "", "", "List Supported Character Sets")
	_rdsCmd.Flags().StringVarP(&_rdsListSupportedTimezones, "list-supported-timezones", "", "", "List Supported Timezones")
	_rdsCmd.Flags().StringVarP(&_rdsLocale, "locale", "", "", "Locale")
	_rdsCmd.Flags().StringVarP(&_rdsLogFileName, "log-file-name", "", "", "Log File Name")
	_rdsCmd.Flags().StringVarP(&_rdsMajorEngineVersion, "major-engine-version", "", "", "Major Engine Version")
	_rdsCmd.Flags().StringVarP(&_rdsManageMasterUserPassword, "manage-master-user-password", "", "", "Manage Master User Password")
	_rdsCmd.Flags().StringVarP(&_rdsManifest, "manifest", "", "", "Manifest")
	_rdsCmd.Flags().StringVarP(&_rdsMarker, "marker", "", "", "Marker")
	_rdsCmd.Flags().StringVarP(&_rdsMasterUserAuthenticationType, "master-user-authentication-type", "", "", "Master User Authentication Type")
	_rdsCmd.Flags().StringVarP(&_rdsMasterUserPassword, "master-user-password", "", "", "Master User Password")
	_rdsCmd.Flags().StringVarP(&_rdsMasterUserSecretKmsKeyId, "master-user-secret-kms-key-id", "", "", "Master User Secret KMS Key ID")
	_rdsCmd.Flags().StringVarP(&_rdsMasterUsername, "master-username", "", "", "Master Username")
	_rdsCmd.Flags().StringVarP(&_rdsMaxACU, "max-acu", "", "", "Max Acu")
	_rdsCmd.Flags().StringVarP(&_rdsMaxAllocatedStorage, "max-allocated-storage", "", "", "Max Allocated Storage")
	_rdsCmd.Flags().StringVarP(&_rdsMaxRecords, "max-records", "", "", "Max Records")
	_rdsCmd.Flags().StringVarP(&_rdsMinACU, "min-acu", "", "", "Min Acu")
	_rdsCmd.Flags().StringVarP(&_rdsMode, "mode", "", "", "Mode")
	_rdsCmd.Flags().StringVarP(&_rdsMonitoringInterval, "monitoring-interval", "", "", "Monitoring Interval")
	_rdsCmd.Flags().StringVarP(&_rdsMonitoringRoleArn, "monitoring-role-arn", "", "", "Monitoring Role ARN")
	_rdsCmd.Flags().StringVarP(&_rdsMultiAZ, "multi-az", "", "", "Multi AZ")
	_rdsCmd.Flags().StringVarP(&_rdsMultiTenant, "multi-tenant", "", "", "Multi Tenant")
	_rdsCmd.Flags().StringVarP(&_rdsNcharCharacterSetName, "nchar-character-set-name", "", "", "Nchar Character Set Name")
	_rdsCmd.Flags().StringVarP(&_rdsNetworkType, "network-type", "", "", "Network Type")
	_rdsCmd.Flags().StringVarP(&_rdsNewDBClusterIdentifier, "new-db-cluster-identifier", "", "", "New DB Cluster Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsNewDBInstanceIdentifier, "new-db-instance-identifier", "", "", "New DB Instance Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsNewDBProxyEndpointName, "new-db-proxy-endpoint-name", "", "", "New DB Proxy Endpoint Name")
	_rdsCmd.Flags().StringVarP(&_rdsNewDBProxyName, "new-db-proxy-name", "", "", "New DB Proxy Name")
	_rdsCmd.Flags().StringVarP(&_rdsNewGlobalClusterIdentifier, "new-global-cluster-identifier", "", "", "New Global Cluster Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsNewName, "new-name", "", "", "New Name")
	_rdsCmd.Flags().StringVarP(&_rdsNewTenantDBName, "new-tenant-db-name", "", "", "New Tenant DB Name")
	_rdsCmd.Flags().StringVarP(&_rdsNumberOfLines, "number-of-lines", "", "", "Number Of Lines")
	_rdsCmd.Flags().StringVarP(&_rdsOfferingType, "offering-type", "", "", "Offering Type")
	_rdsCmd.Flags().StringVarP(&_rdsOptInType, "opt-in-type", "", "", "Opt In Type")
	_rdsCmd.Flags().StringVarP(&_rdsOptionGroupDescription, "option-group-description", "", "", "Option Group Description")
	_rdsCmd.Flags().StringVarP(&_rdsOptionGroupName, "option-group-name", "", "", "Option Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsOptionsToInclude, "options-to-include", "", "", "Options To Include")
	_rdsCmd.Flags().StringSliceVarP(&_rdsOptionsToRemove, "options-to-remove", "", nil, "Options To Remove")
	_rdsCmd.Flags().StringVarP(&_rdsParameters, "parameters", "", "", "Parameters")
	_rdsCmd.Flags().StringVarP(&_rdsPerformanceInsightsKMSKeyId, "performance-insights-kms-key-id", "", "", "Performance Insights KMS Key ID")
	_rdsCmd.Flags().StringVarP(&_rdsPerformanceInsightsRetentionPeriod, "performance-insights-retention-period", "", "", "Performance Insights Retention Period")
	_rdsCmd.Flags().StringVarP(&_rdsPort, "port", "", "", "Port")
	_rdsCmd.Flags().StringVarP(&_rdsPreSignedUrl, "pre-signed-url", "", "", "Pre Signed URL")
	_rdsCmd.Flags().StringVarP(&_rdsPreferredBackupWindow, "preferred-backup-window", "", "", "Preferred Backup Window")
	_rdsCmd.Flags().StringVarP(&_rdsPreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_rdsCmd.Flags().StringVarP(&_rdsProcessorFeatures, "processor-features", "", "", "Processor Features")
	_rdsCmd.Flags().StringVarP(&_rdsProductDescription, "product-description", "", "", "Product Description")
	_rdsCmd.Flags().StringVarP(&_rdsPromotionTier, "promotion-tier", "", "", "Promotion Tier")
	_rdsCmd.Flags().StringVarP(&_rdsPubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_rdsCmd.Flags().StringVarP(&_rdsRdsCustomClusterConfiguration, "rds-custom-cluster-configuration", "", "", "RDS Custom Cluster Configuration")
	_rdsCmd.Flags().StringVarP(&_rdsRecommendationId, "recommendation-id", "", "", "Recommendation ID")
	_rdsCmd.Flags().StringVarP(&_rdsRecommendedActionUpdates, "recommended-action-updates", "", "", "Recommended Action Updates")
	_rdsCmd.Flags().StringVarP(&_rdsRegionName, "region-name", "", "", "Region Name")
	_rdsCmd.Flags().StringVarP(&_rdsRemoveCustomerOverride, "remove-customer-override", "", "", "Remove Customer Override")
	_rdsCmd.Flags().StringVarP(&_rdsReplicaMode, "replica-mode", "", "", "Replica Mode")
	_rdsCmd.Flags().StringVarP(&_rdsReplicationSourceIdentifier, "replication-source-identifier", "", "", "Replication Source Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsRequireTLS, "require-tls", "", "", "Require TLS")
	_rdsCmd.Flags().StringVarP(&_rdsReservedDBInstanceId, "reserved-db-instance-id", "", "", "Reserved DB Instance ID")
	_rdsCmd.Flags().StringVarP(&_rdsReservedDBInstancesOfferingId, "reserved-db-instances-offering-id", "", "", "Reserved DB Instances Offering ID")
	_rdsCmd.Flags().StringVarP(&_rdsResetAllParameters, "reset-all-parameters", "", "", "Reset All Parameters")
	_rdsCmd.Flags().StringVarP(&_rdsResourceArn, "resource-arn", "", "", "Resource ARN")
	_rdsCmd.Flags().StringVarP(&_rdsResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsResourceName, "resource-name", "", "", "Resource Name")
	_rdsCmd.Flags().StringVarP(&_rdsRestoreTime, "restore-time", "", "", "Restore Time")
	_rdsCmd.Flags().StringVarP(&_rdsRestoreToTime, "restore-to-time", "", "", "Restore To Time")
	_rdsCmd.Flags().StringVarP(&_rdsRestoreType, "restore-type", "", "", "Restore Type")
	_rdsCmd.Flags().StringVarP(&_rdsResumeFullAutomationModeMinutes, "resume-full-automation-mode-minutes", "", "", "Resume Full Automation Mode Minutes")
	_rdsCmd.Flags().StringVarP(&_rdsRoleArn, "role-arn", "", "", "Role ARN")
	_rdsCmd.Flags().StringVarP(&_rdsRotateMasterUserPassword, "rotate-master-user-password", "", "", "Rotate Master User Password")
	_rdsCmd.Flags().StringVarP(&_rdsS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_rdsCmd.Flags().StringVarP(&_rdsS3IngestionRoleArn, "s3-ingestion-role-arn", "", "", "S3 Ingestion Role ARN")
	_rdsCmd.Flags().StringVarP(&_rdsS3Prefix, "s3-prefix", "", "", "S3 Prefix")
	_rdsCmd.Flags().StringVarP(&_rdsScalingConfiguration, "scaling-configuration", "", "", "Scaling Configuration")
	_rdsCmd.Flags().StringVarP(&_rdsSecondsBeforeTimeout, "seconds-before-timeout", "", "", "Seconds Before Timeout")
	_rdsCmd.Flags().StringSliceVarP(&_rdsSecurityGroups, "security-groups", "", nil, "Security Groups")
	_rdsCmd.Flags().StringVarP(&_rdsServerlessV2ScalingConfiguration, "serverless-v2-scaling-configuration", "", "", "Serverless V2 Scaling Configuration")
	_rdsCmd.Flags().StringVarP(&_rdsSkipFinalSnapshot, "skip-final-snapshot", "", "", "Skip Final Snapshot")
	_rdsCmd.Flags().StringVarP(&_rdsSnapshotAvailabilityZone, "snapshot-availability-zone", "", "", "Snapshot Availability Zone")
	_rdsCmd.Flags().StringVarP(&_rdsSnapshotIdentifier, "snapshot-identifier", "", "", "Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSnapshotTarget, "snapshot-target", "", "", "Snapshot Target")
	_rdsCmd.Flags().StringVarP(&_rdsSnapshotType, "snapshot-type", "", "", "Snapshot Type")
	_rdsCmd.Flags().StringVarP(&_rdsSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_rdsCmd.Flags().StringVarP(&_rdsSource, "source", "", "", "Source")
	_rdsCmd.Flags().StringVarP(&_rdsSourceArn, "source-arn", "", "", "Source ARN")
	_rdsCmd.Flags().StringVarP(&_rdsSourceCustomDbEngineVersionIdentifier, "source-custom-db-engine-version-identifier", "", "", "Source Custom DB Engine Version Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBClusterIdentifier, "source-db-cluster-identifier", "", "", "Source DB Cluster Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBClusterParameterGroupIdentifier, "source-db-cluster-parameter-group-identifier", "", "", "Source DB Cluster Parameter Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDbClusterResourceId, "source-db-cluster-resource-id", "", "", "Source DB Cluster Resource ID")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBClusterSnapshotIdentifier, "source-db-cluster-snapshot-identifier", "", "", "Source DB Cluster Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBInstanceArn, "source-db-instance-arn", "", "", "Source DB Instance ARN")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBInstanceAutomatedBackupsArn, "source-db-instance-automated-backups-arn", "", "", "Source DB Instance Automated Backups ARN")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBInstanceIdentifier, "source-db-instance-identifier", "", "", "Source DB Instance Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBParameterGroupIdentifier, "source-db-parameter-group-identifier", "", "", "Source DB Parameter Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDBSnapshotIdentifier, "source-db-snapshot-identifier", "", "", "Source DB Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceDbiResourceId, "source-dbi-resource-id", "", "", "Source Dbi Resource ID")
	_rdsCmd.Flags().StringVarP(&_rdsSourceEngine, "source-engine", "", "", "Source Engine")
	_rdsCmd.Flags().StringVarP(&_rdsSourceEngineVersion, "source-engine-version", "", "", "Source Engine Version")
	_rdsCmd.Flags().StringVarP(&_rdsSourceIdentifier, "source-identifier", "", "", "Source Identifier")
	_rdsCmd.Flags().StringSliceVarP(&_rdsSourceIds, "source-ids", "", nil, "Source Ids")
	_rdsCmd.Flags().StringVarP(&_rdsSourceOptionGroupIdentifier, "source-option-group-identifier", "", "", "Source Option Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsSourceRegion, "source-region", "", "", "Source Region")
	_rdsCmd.Flags().StringVarP(&_rdsSourceType, "source-type", "", "", "Source Type")
	_rdsCmd.Flags().StringVarP(&_rdsStartTime, "start-time", "", "", "Start Time")
	_rdsCmd.Flags().StringSliceVarP(&_rdsStaticMembers, "static-members", "", nil, "Static Members")
	_rdsCmd.Flags().StringVarP(&_rdsStatus, "status", "", "", "Status")
	_rdsCmd.Flags().StringVarP(&_rdsStorageEncrypted, "storage-encrypted", "", "", "Storage Encrypted")
	_rdsCmd.Flags().StringVarP(&_rdsStorageThroughput, "storage-throughput", "", "", "Storage Throughput")
	_rdsCmd.Flags().StringVarP(&_rdsStorageType, "storage-type", "", "", "Storage Type")
	_rdsCmd.Flags().StringSliceVarP(&_rdsSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_rdsCmd.Flags().StringVarP(&_rdsSubscriptionName, "subscription-name", "", "", "Subscription Name")
	_rdsCmd.Flags().StringVarP(&_rdsSwitchover, "switchover", "", "", "Switchover")
	_rdsCmd.Flags().StringVarP(&_rdsSwitchoverTimeout, "switchover-timeout", "", "", "Switchover Timeout")
	_rdsCmd.Flags().StringSliceVarP(&_rdsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_rdsCmd.Flags().StringVarP(&_rdsTagSpecifications, "tag-specifications", "", "", "Tag Specifications")
	_rdsCmd.Flags().StringVarP(&_rdsTags, "tags", "", "", "Tags")
	_rdsCmd.Flags().StringVarP(&_rdsTargetAllocatedStorage, "target-allocated-storage", "", "", "Target Allocated Storage")
	_rdsCmd.Flags().StringVarP(&_rdsTargetArn, "target-arn", "", "", "Target ARN")
	_rdsCmd.Flags().StringVarP(&_rdsTargetConnectionNetworkType, "target-connection-network-type", "", "", "Target Connection Network Type")
	_rdsCmd.Flags().StringVarP(&_rdsTargetCustomAvailabilityZone, "target-custom-availability-zone", "", "", "Target Custom Availability Zone")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDbClusterIdentifier, "target-db-cluster-identifier", "", "", "Target DB Cluster Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBClusterParameterGroupDescription, "target-db-cluster-parameter-group-description", "", "", "Target DB Cluster Parameter Group Description")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBClusterParameterGroupIdentifier, "target-db-cluster-parameter-group-identifier", "", "", "Target DB Cluster Parameter Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBClusterParameterGroupName, "target-db-cluster-parameter-group-name", "", "", "Target DB Cluster Parameter Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBClusterSnapshotIdentifier, "target-db-cluster-snapshot-identifier", "", "", "Target DB Cluster Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBInstanceClass, "target-db-instance-class", "", "", "Target DB Instance Class")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBInstanceIdentifier, "target-db-instance-identifier", "", "", "Target DB Instance Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBParameterGroupDescription, "target-db-parameter-group-description", "", "", "Target DB Parameter Group Description")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBParameterGroupIdentifier, "target-db-parameter-group-identifier", "", "", "Target DB Parameter Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBParameterGroupName, "target-db-parameter-group-name", "", "", "Target DB Parameter Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsTargetDBSnapshotIdentifier, "target-db-snapshot-identifier", "", "", "Target DB Snapshot Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetEngineVersion, "target-engine-version", "", "", "Target Engine Version")
	_rdsCmd.Flags().StringVarP(&_rdsTargetGroupName, "target-group-name", "", "", "Target Group Name")
	_rdsCmd.Flags().StringVarP(&_rdsTargetIops, "target-iops", "", "", "Target IOPS")
	_rdsCmd.Flags().StringVarP(&_rdsTargetOptionGroupDescription, "target-option-group-description", "", "", "Target Option Group Description")
	_rdsCmd.Flags().StringVarP(&_rdsTargetOptionGroupIdentifier, "target-option-group-identifier", "", "", "Target Option Group Identifier")
	_rdsCmd.Flags().StringVarP(&_rdsTargetRole, "target-role", "", "", "Target Role")
	_rdsCmd.Flags().StringVarP(&_rdsTargetStorageThroughput, "target-storage-throughput", "", "", "Target Storage Throughput")
	_rdsCmd.Flags().StringVarP(&_rdsTargetStorageType, "target-storage-type", "", "", "Target Storage Type")
	_rdsCmd.Flags().StringVarP(&_rdsTdeCredentialArn, "tde-credential-arn", "", "", "Tde Credential ARN")
	_rdsCmd.Flags().StringVarP(&_rdsTdeCredentialPassword, "tde-credential-password", "", "", "Tde Credential Password")
	_rdsCmd.Flags().StringVarP(&_rdsTenantDBName, "tenant-db-name", "", "", "Tenant DB Name")
	_rdsCmd.Flags().StringVarP(&_rdsTimeoutAction, "timeout-action", "", "", "Timeout Action")
	_rdsCmd.Flags().StringVarP(&_rdsTimezone, "timezone", "", "", "Timezone")
	_rdsCmd.Flags().StringVarP(&_rdsUpgradeStorageConfig, "upgrade-storage-config", "", "", "Upgrade Storage Config")
	_rdsCmd.Flags().StringVarP(&_rdsUpgradeTargetStorageConfig, "upgrade-target-storage-config", "", "", "Upgrade Target Storage Config")
	_rdsCmd.Flags().StringVarP(&_rdsUseAwsProvidedLatestImage, "use-aws-provided-latest-image", "", "", "Use AWS Provided Latest Image")
	_rdsCmd.Flags().StringVarP(&_rdsUseDefaultProcessorFeatures, "use-default-processor-features", "", "", "Use Default Processor Features")
	_rdsCmd.Flags().StringVarP(&_rdsUseEarliestTimeOnPointInTimeUnavailable, "use-earliest-time-on-point-in-time-unavailable", "", "", "Use Earliest Time On Point In Time Unavailable")
	_rdsCmd.Flags().StringVarP(&_rdsUseLatestRestorableTime, "use-latest-restorable-time", "", "", "Use Latest Restorable Time")
	_rdsCmd.Flags().StringSliceVarP(&_rdsValuesToAdd, "values-to-add", "", nil, "Values To Add")
	_rdsCmd.Flags().StringSliceVarP(&_rdsValuesToRemove, "values-to-remove", "", nil, "Values To Remove")
	_rdsCmd.Flags().StringVarP(&_rdsVpc, "vpc", "", "", "VPC")
	_rdsCmd.Flags().StringSliceVarP(&_rdsVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")
	_rdsCmd.Flags().StringSliceVarP(&_rdsVpcSubnetIds, "vpc-subnet-ids", "", nil, "VPC Subnet Ids")

	_rdsCmd.Flags().BoolVarP(&_rdsAddRoleToDBCluster, "add-role-to-db-cluster", "", false, "Add Role To DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsAddRoleToDBInstance, "add-role-to-db-instance", "", false, "Add Role To DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsAddSourceIdentifierToSubscription, "add-source-identifier-to-subscription", "", false, "Add Source Identifier To Subscription")
	_rdsCmd.Flags().BoolVarP(&_rdsAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_rdsCmd.Flags().BoolVarP(&_rdsApplyPendingMaintenanceAction, "apply-pending-maintenance-action", "", false, "Apply Pending Maintenance Action")
	_rdsCmd.Flags().BoolVarP(&_rdsAuthorizeDBSecurityGroupIngress, "authorize-db-security-group-ingress", "", false, "Authorize DB Security Group Ingress")
	_rdsCmd.Flags().BoolVarP(&_rdsBacktrackDBCluster, "backtrack-db-cluster", "", false, "Backtrack DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsCancelExportTask, "cancel-export-task", "", false, "Cancel Export Task")
	_rdsCmd.Flags().BoolVarP(&_rdsCopyDBClusterParameterGroup, "copy-db-cluster-parameter-group", "", false, "Copy DB Cluster Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCopyDBClusterSnapshot, "copy-db-cluster-snapshot", "", false, "Copy DB Cluster Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsCopyDBParameterGroup, "copy-db-parameter-group", "", false, "Copy DB Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCopyDBSnapshot, "copy-db-snapshot", "", false, "Copy DB Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsCopyOptionGroup, "copy-option-group", "", false, "Copy Option Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateBlueGreenDeployment, "create-blue-green-deployment", "", false, "Create Blue Green Deployment")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateCustomDBEngineVersion, "create-custom-db-engine-version", "", false, "Create Custom DB Engine Version")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBCluster, "create-db-cluster", "", false, "Create DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBClusterEndpoint, "create-db-cluster-endpoint", "", false, "Create DB Cluster Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBClusterParameterGroup, "create-db-cluster-parameter-group", "", false, "Create DB Cluster Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBClusterSnapshot, "create-db-cluster-snapshot", "", false, "Create DB Cluster Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBInstance, "create-db-instance", "", false, "Create DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBInstanceReadReplica, "create-db-instance-read-replica", "", false, "Create DB Instance Read Replica")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBParameterGroup, "create-db-parameter-group", "", false, "Create DB Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBProxy, "create-db-proxy", "", false, "Create DB Proxy")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBProxyEndpoint, "create-db-proxy-endpoint", "", false, "Create DB Proxy Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBSecurityGroup, "create-db-security-group", "", false, "Create DB Security Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBShardGroup, "create-db-shard-group", "", false, "Create DB Shard Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBSnapshot, "create-db-snapshot", "", false, "Create DB Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateDBSubnetGroup, "create-db-subnet-group", "", false, "Create DB Subnet Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateEventSubscription, "create-event-subscription", "", false, "Create Event Subscription")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateGlobalCluster, "create-global-cluster", "", false, "Create Global Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateIntegration, "create-integration", "", false, "Create Integration")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateOptionGroup, "create-option-group", "", false, "Create Option Group")
	_rdsCmd.Flags().BoolVarP(&_rdsCreateTenantDatabase, "create-tenant-database", "", false, "Create Tenant Database")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteBlueGreenDeployment, "delete-blue-green-deployment", "", false, "Delete Blue Green Deployment")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteCustomDBEngineVersion, "delete-custom-db-engine-version", "", false, "Delete Custom DB Engine Version")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBCluster, "delete-db-cluster", "", false, "Delete DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBClusterAutomatedBackup, "delete-db-cluster-automated-backup", "", false, "Delete DB Cluster Automated Backup")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBClusterEndpoint, "delete-db-cluster-endpoint", "", false, "Delete DB Cluster Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBClusterParameterGroup, "delete-db-cluster-parameter-group", "", false, "Delete DB Cluster Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBClusterSnapshot, "delete-db-cluster-snapshot", "", false, "Delete DB Cluster Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBInstance, "delete-db-instance", "", false, "Delete DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBInstanceAutomatedBackup, "delete-db-instance-automated-backup", "", false, "Delete DB Instance Automated Backup")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBParameterGroup, "delete-db-parameter-group", "", false, "Delete DB Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBProxy, "delete-db-proxy", "", false, "Delete DB Proxy")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBProxyEndpoint, "delete-db-proxy-endpoint", "", false, "Delete DB Proxy Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBSecurityGroup, "delete-db-security-group", "", false, "Delete DB Security Group")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBShardGroup, "delete-db-shard-group", "", false, "Delete DB Shard Group")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBSnapshot, "delete-db-snapshot", "", false, "Delete DB Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteDBSubnetGroup, "delete-db-subnet-group", "", false, "Delete DB Subnet Group")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteEventSubscription, "delete-event-subscription", "", false, "Delete Event Subscription")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteGlobalCluster, "delete-global-cluster", "", false, "Delete Global Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteOptionGroup, "delete-option-group", "", false, "Delete Option Group")
	_rdsCmd.Flags().BoolVarP(&_rdsDeleteTenantDatabase, "delete-tenant-database", "", false, "Delete Tenant Database")
	_rdsCmd.Flags().BoolVarP(&_rdsDeregisterDBProxyTargets, "deregister-db-proxy-targets", "", false, "Deregister DB Proxy Targets")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeAccountAttributes, "describe-account-attributes", "", false, "Describe Account Attributes")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeBlueGreenDeployments, "describe-blue-green-deployments", "", false, "Describe Blue Green Deployments")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeCertificates, "describe-certificates", "", false, "Describe Certificates")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterAutomatedBackups, "describe-db-cluster-automated-backups", "", false, "Describe DB Cluster Automated Backups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterBacktracks, "describe-db-cluster-backtracks", "", false, "Describe DB Cluster Backtracks")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterEndpoints, "describe-db-cluster-endpoints", "", false, "Describe DB Cluster Endpoints")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterParameterGroups, "describe-db-cluster-parameter-groups", "", false, "Describe DB Cluster Parameter Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterParameters, "describe-db-cluster-parameters", "", false, "Describe DB Cluster Parameters")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterSnapshotAttributes, "describe-db-cluster-snapshot-attributes", "", false, "Describe DB Cluster Snapshot Attributes")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusterSnapshots, "describe-db-cluster-snapshots", "", false, "Describe DB Cluster Snapshots")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBClusters, "describe-db-clusters", "", false, "Describe DB Clusters")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBEngineVersions, "describe-db-engine-versions", "", false, "Describe DB Engine Versions")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBInstanceAutomatedBackups, "describe-db-instance-automated-backups", "", false, "Describe DB Instance Automated Backups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBInstances, "describe-db-instances", "", false, "Describe DB Instances")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBLogFiles, "describe-db-log-files", "", false, "Describe DB Log Files")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBMajorEngineVersions, "describe-db-major-engine-versions", "", false, "Describe DB Major Engine Versions")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBParameterGroups, "describe-db-parameter-groups", "", false, "Describe DB Parameter Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBParameters, "describe-db-parameters", "", false, "Describe DB Parameters")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBProxies, "describe-db-proxies", "", false, "Describe DB Proxies")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBProxyEndpoints, "describe-db-proxy-endpoints", "", false, "Describe DB Proxy Endpoints")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBProxyTargetGroups, "describe-db-proxy-target-groups", "", false, "Describe DB Proxy Target Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBProxyTargets, "describe-db-proxy-targets", "", false, "Describe DB Proxy Targets")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBRecommendations, "describe-db-recommendations", "", false, "Describe DB Recommendations")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBSecurityGroups, "describe-db-security-groups", "", false, "Describe DB Security Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBShardGroups, "describe-db-shard-groups", "", false, "Describe DB Shard Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBSnapshotAttributes, "describe-db-snapshot-attributes", "", false, "Describe DB Snapshot Attributes")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBSnapshotTenantDatabases, "describe-db-snapshot-tenant-databases", "", false, "Describe DB Snapshot Tenant Databases")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBSnapshots, "describe-db-snapshots", "", false, "Describe DB Snapshots")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeDBSubnetGroups, "describe-db-subnet-groups", "", false, "Describe DB Subnet Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeEngineDefaultClusterParameters, "describe-engine-default-cluster-parameters", "", false, "Describe Engine Default Cluster Parameters")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeEngineDefaultParameters, "describe-engine-default-parameters", "", false, "Describe Engine Default Parameters")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeEventCategories, "describe-event-categories", "", false, "Describe Event Categories")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeEventSubscriptions, "describe-event-subscriptions", "", false, "Describe Event Subscriptions")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeEvents, "describe-events", "", false, "Describe Events")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeExportTasks, "describe-export-tasks", "", false, "Describe Export Tasks")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeGlobalClusters, "describe-global-clusters", "", false, "Describe Global Clusters")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeIntegrations, "describe-integrations", "", false, "Describe Integrations")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeOptionGroupOptions, "describe-option-group-options", "", false, "Describe Option Group Options")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeOptionGroups, "describe-option-groups", "", false, "Describe Option Groups")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeOrderableDBInstanceOptions, "describe-orderable-db-instance-options", "", false, "Describe Orderable DB Instance Options")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribePendingMaintenanceActions, "describe-pending-maintenance-actions", "", false, "Describe Pending Maintenance Actions")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeReservedDBInstances, "describe-reserved-db-instances", "", false, "Describe Reserved DB Instances")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeReservedDBInstancesOfferings, "describe-reserved-db-instances-offerings", "", false, "Describe Reserved DB Instances Offerings")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeSourceRegions, "describe-source-regions", "", false, "Describe Source Regions")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeTenantDatabases, "describe-tenant-databases", "", false, "Describe Tenant Databases")
	_rdsCmd.Flags().BoolVarP(&_rdsDescribeValidDBInstanceModifications, "describe-valid-db-instance-modifications", "", false, "Describe Valid DB Instance Modifications")
	_rdsCmd.Flags().BoolVarP(&_rdsDisableHttpEndpoint, "disable-http-endpoint", "", false, "Disable HTTP Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsDownloadDBLogFilePortion, "download-db-log-file-portion", "", false, "Download DB Log File Portion")
	_rdsCmd.Flags().BoolVarP(&_rdsEnableHttpEndpoint, "enable-http-endpoint", "", false, "Enable HTTP Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsFailoverDBCluster, "failover-db-cluster", "", false, "Failover DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsFailoverGlobalCluster, "failover-global-cluster", "", false, "Failover Global Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyActivityStream, "modify-activity-stream", "", false, "Modify Activity Stream")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyCertificates, "modify-certificates", "", false, "Modify Certificates")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyCurrentDBClusterCapacity, "modify-current-db-cluster-capacity", "", false, "Modify Current DB Cluster Capacity")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyCustomDBEngineVersion, "modify-custom-db-engine-version", "", false, "Modify Custom DB Engine Version")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBCluster, "modify-db-cluster", "", false, "Modify DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBClusterEndpoint, "modify-db-cluster-endpoint", "", false, "Modify DB Cluster Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBClusterParameterGroup, "modify-db-cluster-parameter-group", "", false, "Modify DB Cluster Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBClusterSnapshotAttribute, "modify-db-cluster-snapshot-attribute", "", false, "Modify DB Cluster Snapshot Attribute")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBInstance, "modify-db-instance", "", false, "Modify DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBParameterGroup, "modify-db-parameter-group", "", false, "Modify DB Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBProxy, "modify-db-proxy", "", false, "Modify DB Proxy")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBProxyEndpoint, "modify-db-proxy-endpoint", "", false, "Modify DB Proxy Endpoint")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBProxyTargetGroup, "modify-db-proxy-target-group", "", false, "Modify DB Proxy Target Group")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBRecommendation, "modify-db-recommendation", "", false, "Modify DB Recommendation")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBShardGroup, "modify-db-shard-group", "", false, "Modify DB Shard Group")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBSnapshot, "modify-db-snapshot", "", false, "Modify DB Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBSnapshotAttribute, "modify-db-snapshot-attribute", "", false, "Modify DB Snapshot Attribute")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyDBSubnetGroup, "modify-db-subnet-group", "", false, "Modify DB Subnet Group")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyEventSubscription, "modify-event-subscription", "", false, "Modify Event Subscription")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyGlobalCluster, "modify-global-cluster", "", false, "Modify Global Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyIntegration, "modify-integration", "", false, "Modify Integration")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyOptionGroup, "modify-option-group", "", false, "Modify Option Group")
	_rdsCmd.Flags().BoolVarP(&_rdsModifyTenantDatabase, "modify-tenant-database", "", false, "Modify Tenant Database")
	_rdsCmd.Flags().BoolVarP(&_rdsPromoteReadReplica, "promote-read-replica", "", false, "Promote Read Replica")
	_rdsCmd.Flags().BoolVarP(&_rdsPromoteReadReplicaDBCluster, "promote-read-replica-db-cluster", "", false, "Promote Read Replica DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsPurchaseReservedDBInstancesOffering, "purchase-reserved-db-instances-offering", "", false, "Purchase Reserved DB Instances Offering")
	_rdsCmd.Flags().BoolVarP(&_rdsRebootDBCluster, "reboot-db-cluster", "", false, "Reboot DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsRebootDBInstance, "reboot-db-instance", "", false, "Reboot DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsRebootDBShardGroup, "reboot-db-shard-group", "", false, "Reboot DB Shard Group")
	_rdsCmd.Flags().BoolVarP(&_rdsRegisterDBProxyTargets, "register-db-proxy-targets", "", false, "Register DB Proxy Targets")
	_rdsCmd.Flags().BoolVarP(&_rdsRemoveFromGlobalCluster, "remove-from-global-cluster", "", false, "Remove From Global Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsRemoveRoleFromDBCluster, "remove-role-from-db-cluster", "", false, "Remove Role From DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsRemoveRoleFromDBInstance, "remove-role-from-db-instance", "", false, "Remove Role From DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsRemoveSourceIdentifierFromSubscription, "remove-source-identifier-from-subscription", "", false, "Remove Source Identifier From Subscription")
	_rdsCmd.Flags().BoolVarP(&_rdsRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_rdsCmd.Flags().BoolVarP(&_rdsResetDBClusterParameterGroup, "reset-db-cluster-parameter-group", "", false, "Reset DB Cluster Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsResetDBParameterGroup, "reset-db-parameter-group", "", false, "Reset DB Parameter Group")
	_rdsCmd.Flags().BoolVarP(&_rdsRestoreDBClusterFromS3, "restore-db-cluster-from-s3", "", false, "Restore DB Cluster From S3")
	_rdsCmd.Flags().BoolVarP(&_rdsRestoreDBClusterFromSnapshot, "restore-db-cluster-from-snapshot", "", false, "Restore DB Cluster From Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsRestoreDBClusterToPointInTime, "restore-db-cluster-to-point-in-time", "", false, "Restore DB Cluster To Point In Time")
	_rdsCmd.Flags().BoolVarP(&_rdsRestoreDBInstanceFromDBSnapshot, "restore-db-instance-from-db-snapshot", "", false, "Restore DB Instance From DB Snapshot")
	_rdsCmd.Flags().BoolVarP(&_rdsRestoreDBInstanceFromS3, "restore-db-instance-from-s3", "", false, "Restore DB Instance From S3")
	_rdsCmd.Flags().BoolVarP(&_rdsRestoreDBInstanceToPointInTime, "restore-db-instance-to-point-in-time", "", false, "Restore DB Instance To Point In Time")
	_rdsCmd.Flags().BoolVarP(&_rdsRevokeDBSecurityGroupIngress, "revoke-db-security-group-ingress", "", false, "Revoke DB Security Group Ingress")
	_rdsCmd.Flags().BoolVarP(&_rdsStartActivityStream, "start-activity-stream", "", false, "Start Activity Stream")
	_rdsCmd.Flags().BoolVarP(&_rdsStartDBCluster, "start-db-cluster", "", false, "Start DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsStartDBInstance, "start-db-instance", "", false, "Start DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsStartDBInstanceAutomatedBackupsReplication, "start-db-instance-automated-backups-replication", "", false, "Start DB Instance Automated Backups Replication")
	_rdsCmd.Flags().BoolVarP(&_rdsStartExportTask, "start-export-task", "", false, "Start Export Task")
	_rdsCmd.Flags().BoolVarP(&_rdsStopActivityStream, "stop-activity-stream", "", false, "Stop Activity Stream")
	_rdsCmd.Flags().BoolVarP(&_rdsStopDBCluster, "stop-db-cluster", "", false, "Stop DB Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsStopDBInstance, "stop-db-instance", "", false, "Stop DB Instance")
	_rdsCmd.Flags().BoolVarP(&_rdsStopDBInstanceAutomatedBackupsReplication, "stop-db-instance-automated-backups-replication", "", false, "Stop DB Instance Automated Backups Replication")
	_rdsCmd.Flags().BoolVarP(&_rdsSwitchoverBlueGreenDeployment, "switchover-blue-green-deployment", "", false, "Switchover Blue Green Deployment")
	_rdsCmd.Flags().BoolVarP(&_rdsSwitchoverGlobalCluster, "switchover-global-cluster", "", false, "Switchover Global Cluster")
	_rdsCmd.Flags().BoolVarP(&_rdsSwitchoverReadReplica, "switchover-read-replica", "", false, "Switchover Read Replica")

}
