package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// databasemigrationserviceCmd represents the databasemigrationservice command
var _databasemigrationserviceCmd = &cobra.Command{
	Use:   "databasemigrationservice",
	Short: "AWS databasemigrationservice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := databasemigrationservice.NewFromConfig(cfg)
		if _databasemigrationserviceAddTagsToResource {
			databasemigrationservice_AddTagsToResource(cfg, client)
			return
		}
		if _databasemigrationserviceApplyPendingMaintenanceAction {
			databasemigrationservice_ApplyPendingMaintenanceAction(cfg, client)
			return
		}
		if _databasemigrationserviceBatchStartRecommendations {
			databasemigrationservice_BatchStartRecommendations(cfg, client)
			return
		}
		if _databasemigrationserviceCancelMetadataModelConversion {
			databasemigrationservice_CancelMetadataModelConversion(cfg, client)
			return
		}
		if _databasemigrationserviceCancelMetadataModelCreation {
			databasemigrationservice_CancelMetadataModelCreation(cfg, client)
			return
		}
		if _databasemigrationserviceCancelReplicationTaskAssessmentRun {
			databasemigrationservice_CancelReplicationTaskAssessmentRun(cfg, client)
			return
		}
		if _databasemigrationserviceCreateDataMigration {
			databasemigrationservice_CreateDataMigration(cfg, client)
			return
		}
		if _databasemigrationserviceCreateDataProvider {
			databasemigrationservice_CreateDataProvider(cfg, client)
			return
		}
		if _databasemigrationserviceCreateEndpoint {
			databasemigrationservice_CreateEndpoint(cfg, client)
			return
		}
		if _databasemigrationserviceCreateEventSubscription {
			databasemigrationservice_CreateEventSubscription(cfg, client)
			return
		}
		if _databasemigrationserviceCreateFleetAdvisorCollector {
			databasemigrationservice_CreateFleetAdvisorCollector(cfg, client)
			return
		}
		if _databasemigrationserviceCreateInstanceProfile {
			databasemigrationservice_CreateInstanceProfile(cfg, client)
			return
		}
		if _databasemigrationserviceCreateMigrationProject {
			databasemigrationservice_CreateMigrationProject(cfg, client)
			return
		}
		if _databasemigrationserviceCreateReplicationConfig {
			databasemigrationservice_CreateReplicationConfig(cfg, client)
			return
		}
		if _databasemigrationserviceCreateReplicationInstance {
			databasemigrationservice_CreateReplicationInstance(cfg, client)
			return
		}
		if _databasemigrationserviceCreateReplicationSubnetGroup {
			databasemigrationservice_CreateReplicationSubnetGroup(cfg, client)
			return
		}
		if _databasemigrationserviceCreateReplicationTask {
			databasemigrationservice_CreateReplicationTask(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteCertificate {
			databasemigrationservice_DeleteCertificate(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteConnection {
			databasemigrationservice_DeleteConnection(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteDataMigration {
			databasemigrationservice_DeleteDataMigration(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteDataProvider {
			databasemigrationservice_DeleteDataProvider(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteEndpoint {
			databasemigrationservice_DeleteEndpoint(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteEventSubscription {
			databasemigrationservice_DeleteEventSubscription(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteFleetAdvisorCollector {
			databasemigrationservice_DeleteFleetAdvisorCollector(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteFleetAdvisorDatabases {
			databasemigrationservice_DeleteFleetAdvisorDatabases(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteInstanceProfile {
			databasemigrationservice_DeleteInstanceProfile(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteMigrationProject {
			databasemigrationservice_DeleteMigrationProject(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteReplicationConfig {
			databasemigrationservice_DeleteReplicationConfig(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteReplicationInstance {
			databasemigrationservice_DeleteReplicationInstance(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteReplicationSubnetGroup {
			databasemigrationservice_DeleteReplicationSubnetGroup(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteReplicationTask {
			databasemigrationservice_DeleteReplicationTask(cfg, client)
			return
		}
		if _databasemigrationserviceDeleteReplicationTaskAssessmentRun {
			databasemigrationservice_DeleteReplicationTaskAssessmentRun(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeAccountAttributes {
			databasemigrationservice_DescribeAccountAttributes(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeApplicableIndividualAssessments {
			databasemigrationservice_DescribeApplicableIndividualAssessments(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeCertificates {
			databasemigrationservice_DescribeCertificates(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeConnections {
			databasemigrationservice_DescribeConnections(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeConversionConfiguration {
			databasemigrationservice_DescribeConversionConfiguration(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeDataMigrations {
			databasemigrationservice_DescribeDataMigrations(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeDataProviders {
			databasemigrationservice_DescribeDataProviders(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEndpointSettings {
			databasemigrationservice_DescribeEndpointSettings(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEndpointTypes {
			databasemigrationservice_DescribeEndpointTypes(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEndpoints {
			databasemigrationservice_DescribeEndpoints(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEngineVersions {
			databasemigrationservice_DescribeEngineVersions(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEventCategories {
			databasemigrationservice_DescribeEventCategories(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEventSubscriptions {
			databasemigrationservice_DescribeEventSubscriptions(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeEvents {
			databasemigrationservice_DescribeEvents(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeExtensionPackAssociations {
			databasemigrationservice_DescribeExtensionPackAssociations(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeFleetAdvisorCollectors {
			databasemigrationservice_DescribeFleetAdvisorCollectors(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeFleetAdvisorDatabases {
			databasemigrationservice_DescribeFleetAdvisorDatabases(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeFleetAdvisorLsaAnalysis {
			databasemigrationservice_DescribeFleetAdvisorLsaAnalysis(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeFleetAdvisorSchemaObjectSummary {
			databasemigrationservice_DescribeFleetAdvisorSchemaObjectSummary(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeFleetAdvisorSchemas {
			databasemigrationservice_DescribeFleetAdvisorSchemas(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeInstanceProfiles {
			databasemigrationservice_DescribeInstanceProfiles(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModel {
			databasemigrationservice_DescribeMetadataModel(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelAssessments {
			databasemigrationservice_DescribeMetadataModelAssessments(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelChildren {
			databasemigrationservice_DescribeMetadataModelChildren(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelConversions {
			databasemigrationservice_DescribeMetadataModelConversions(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelCreations {
			databasemigrationservice_DescribeMetadataModelCreations(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelExportsAsScript {
			databasemigrationservice_DescribeMetadataModelExportsAsScript(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelExportsToTarget {
			databasemigrationservice_DescribeMetadataModelExportsToTarget(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMetadataModelImports {
			databasemigrationservice_DescribeMetadataModelImports(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeMigrationProjects {
			databasemigrationservice_DescribeMigrationProjects(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeOrderableReplicationInstances {
			databasemigrationservice_DescribeOrderableReplicationInstances(cfg, client)
			return
		}
		if _databasemigrationserviceDescribePendingMaintenanceActions {
			databasemigrationservice_DescribePendingMaintenanceActions(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeRecommendationLimitations {
			databasemigrationservice_DescribeRecommendationLimitations(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeRecommendations {
			databasemigrationservice_DescribeRecommendations(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeRefreshSchemasStatus {
			databasemigrationservice_DescribeRefreshSchemasStatus(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationConfigs {
			databasemigrationservice_DescribeReplicationConfigs(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationInstanceTaskLogs {
			databasemigrationservice_DescribeReplicationInstanceTaskLogs(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationInstances {
			databasemigrationservice_DescribeReplicationInstances(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationSubnetGroups {
			databasemigrationservice_DescribeReplicationSubnetGroups(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationTableStatistics {
			databasemigrationservice_DescribeReplicationTableStatistics(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationTaskAssessmentResults {
			databasemigrationservice_DescribeReplicationTaskAssessmentResults(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationTaskAssessmentRuns {
			databasemigrationservice_DescribeReplicationTaskAssessmentRuns(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationTaskIndividualAssessments {
			databasemigrationservice_DescribeReplicationTaskIndividualAssessments(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplicationTasks {
			databasemigrationservice_DescribeReplicationTasks(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeReplications {
			databasemigrationservice_DescribeReplications(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeSchemas {
			databasemigrationservice_DescribeSchemas(cfg, client)
			return
		}
		if _databasemigrationserviceDescribeTableStatistics {
			databasemigrationservice_DescribeTableStatistics(cfg, client)
			return
		}
		if _databasemigrationserviceExportMetadataModelAssessment {
			databasemigrationservice_ExportMetadataModelAssessment(cfg, client)
			return
		}
		if _databasemigrationserviceGetTargetSelectionRules {
			databasemigrationservice_GetTargetSelectionRules(cfg, client)
			return
		}
		if _databasemigrationserviceImportCertificate {
			databasemigrationservice_ImportCertificate(cfg, client)
			return
		}
		if _databasemigrationserviceListTagsForResource {
			databasemigrationservice_ListTagsForResource(cfg, client)
			return
		}
		if _databasemigrationserviceModifyConversionConfiguration {
			databasemigrationservice_ModifyConversionConfiguration(cfg, client)
			return
		}
		if _databasemigrationserviceModifyDataMigration {
			databasemigrationservice_ModifyDataMigration(cfg, client)
			return
		}
		if _databasemigrationserviceModifyDataProvider {
			databasemigrationservice_ModifyDataProvider(cfg, client)
			return
		}
		if _databasemigrationserviceModifyEndpoint {
			databasemigrationservice_ModifyEndpoint(cfg, client)
			return
		}
		if _databasemigrationserviceModifyEventSubscription {
			databasemigrationservice_ModifyEventSubscription(cfg, client)
			return
		}
		if _databasemigrationserviceModifyInstanceProfile {
			databasemigrationservice_ModifyInstanceProfile(cfg, client)
			return
		}
		if _databasemigrationserviceModifyMigrationProject {
			databasemigrationservice_ModifyMigrationProject(cfg, client)
			return
		}
		if _databasemigrationserviceModifyReplicationConfig {
			databasemigrationservice_ModifyReplicationConfig(cfg, client)
			return
		}
		if _databasemigrationserviceModifyReplicationInstance {
			databasemigrationservice_ModifyReplicationInstance(cfg, client)
			return
		}
		if _databasemigrationserviceModifyReplicationSubnetGroup {
			databasemigrationservice_ModifyReplicationSubnetGroup(cfg, client)
			return
		}
		if _databasemigrationserviceModifyReplicationTask {
			databasemigrationservice_ModifyReplicationTask(cfg, client)
			return
		}
		if _databasemigrationserviceMoveReplicationTask {
			databasemigrationservice_MoveReplicationTask(cfg, client)
			return
		}
		if _databasemigrationserviceRebootReplicationInstance {
			databasemigrationservice_RebootReplicationInstance(cfg, client)
			return
		}
		if _databasemigrationserviceRefreshSchemas {
			databasemigrationservice_RefreshSchemas(cfg, client)
			return
		}
		if _databasemigrationserviceReloadReplicationTables {
			databasemigrationservice_ReloadReplicationTables(cfg, client)
			return
		}
		if _databasemigrationserviceReloadTables {
			databasemigrationservice_ReloadTables(cfg, client)
			return
		}
		if _databasemigrationserviceRemoveTagsFromResource {
			databasemigrationservice_RemoveTagsFromResource(cfg, client)
			return
		}
		if _databasemigrationserviceRunFleetAdvisorLsaAnalysis {
			databasemigrationservice_RunFleetAdvisorLsaAnalysis(cfg, client)
			return
		}
		if _databasemigrationserviceStartDataMigration {
			databasemigrationservice_StartDataMigration(cfg, client)
			return
		}
		if _databasemigrationserviceStartExtensionPackAssociation {
			databasemigrationservice_StartExtensionPackAssociation(cfg, client)
			return
		}
		if _databasemigrationserviceStartMetadataModelAssessment {
			databasemigrationservice_StartMetadataModelAssessment(cfg, client)
			return
		}
		if _databasemigrationserviceStartMetadataModelConversion {
			databasemigrationservice_StartMetadataModelConversion(cfg, client)
			return
		}
		if _databasemigrationserviceStartMetadataModelCreation {
			databasemigrationservice_StartMetadataModelCreation(cfg, client)
			return
		}
		if _databasemigrationserviceStartMetadataModelExportAsScript {
			databasemigrationservice_StartMetadataModelExportAsScript(cfg, client)
			return
		}
		if _databasemigrationserviceStartMetadataModelExportToTarget {
			databasemigrationservice_StartMetadataModelExportToTarget(cfg, client)
			return
		}
		if _databasemigrationserviceStartMetadataModelImport {
			databasemigrationservice_StartMetadataModelImport(cfg, client)
			return
		}
		if _databasemigrationserviceStartRecommendations {
			databasemigrationservice_StartRecommendations(cfg, client)
			return
		}
		if _databasemigrationserviceStartReplication {
			databasemigrationservice_StartReplication(cfg, client)
			return
		}
		if _databasemigrationserviceStartReplicationTask {
			databasemigrationservice_StartReplicationTask(cfg, client)
			return
		}
		if _databasemigrationserviceStartReplicationTaskAssessment {
			databasemigrationservice_StartReplicationTaskAssessment(cfg, client)
			return
		}
		if _databasemigrationserviceStartReplicationTaskAssessmentRun {
			databasemigrationservice_StartReplicationTaskAssessmentRun(cfg, client)
			return
		}
		if _databasemigrationserviceStopDataMigration {
			databasemigrationservice_StopDataMigration(cfg, client)
			return
		}
		if _databasemigrationserviceStopReplication {
			databasemigrationservice_StopReplication(cfg, client)
			return
		}
		if _databasemigrationserviceStopReplicationTask {
			databasemigrationservice_StopReplicationTask(cfg, client)
			return
		}
		if _databasemigrationserviceTestConnection {
			databasemigrationservice_TestConnection(cfg, client)
			return
		}
		if _databasemigrationserviceUpdateSubscriptionsToEventBridge {
			databasemigrationservice_UpdateSubscriptionsToEventBridge(cfg, client)
			return
		}

	},
}

var (
	_databasemigrationserviceAddTagsToResource                            bool
	_databasemigrationserviceApplyPendingMaintenanceAction                bool
	_databasemigrationserviceBatchStartRecommendations                    bool
	_databasemigrationserviceCancelMetadataModelConversion                bool
	_databasemigrationserviceCancelMetadataModelCreation                  bool
	_databasemigrationserviceCancelReplicationTaskAssessmentRun           bool
	_databasemigrationserviceCreateDataMigration                          bool
	_databasemigrationserviceCreateDataProvider                           bool
	_databasemigrationserviceCreateEndpoint                               bool
	_databasemigrationserviceCreateEventSubscription                      bool
	_databasemigrationserviceCreateFleetAdvisorCollector                  bool
	_databasemigrationserviceCreateInstanceProfile                        bool
	_databasemigrationserviceCreateMigrationProject                       bool
	_databasemigrationserviceCreateReplicationConfig                      bool
	_databasemigrationserviceCreateReplicationInstance                    bool
	_databasemigrationserviceCreateReplicationSubnetGroup                 bool
	_databasemigrationserviceCreateReplicationTask                        bool
	_databasemigrationserviceDeleteCertificate                            bool
	_databasemigrationserviceDeleteConnection                             bool
	_databasemigrationserviceDeleteDataMigration                          bool
	_databasemigrationserviceDeleteDataProvider                           bool
	_databasemigrationserviceDeleteEndpoint                               bool
	_databasemigrationserviceDeleteEventSubscription                      bool
	_databasemigrationserviceDeleteFleetAdvisorCollector                  bool
	_databasemigrationserviceDeleteFleetAdvisorDatabases                  bool
	_databasemigrationserviceDeleteInstanceProfile                        bool
	_databasemigrationserviceDeleteMigrationProject                       bool
	_databasemigrationserviceDeleteReplicationConfig                      bool
	_databasemigrationserviceDeleteReplicationInstance                    bool
	_databasemigrationserviceDeleteReplicationSubnetGroup                 bool
	_databasemigrationserviceDeleteReplicationTask                        bool
	_databasemigrationserviceDeleteReplicationTaskAssessmentRun           bool
	_databasemigrationserviceDescribeAccountAttributes                    bool
	_databasemigrationserviceDescribeApplicableIndividualAssessments      bool
	_databasemigrationserviceDescribeCertificates                         bool
	_databasemigrationserviceDescribeConnections                          bool
	_databasemigrationserviceDescribeConversionConfiguration              bool
	_databasemigrationserviceDescribeDataMigrations                       bool
	_databasemigrationserviceDescribeDataProviders                        bool
	_databasemigrationserviceDescribeEndpointSettings                     bool
	_databasemigrationserviceDescribeEndpointTypes                        bool
	_databasemigrationserviceDescribeEndpoints                            bool
	_databasemigrationserviceDescribeEngineVersions                       bool
	_databasemigrationserviceDescribeEventCategories                      bool
	_databasemigrationserviceDescribeEventSubscriptions                   bool
	_databasemigrationserviceDescribeEvents                               bool
	_databasemigrationserviceDescribeExtensionPackAssociations            bool
	_databasemigrationserviceDescribeFleetAdvisorCollectors               bool
	_databasemigrationserviceDescribeFleetAdvisorDatabases                bool
	_databasemigrationserviceDescribeFleetAdvisorLsaAnalysis              bool
	_databasemigrationserviceDescribeFleetAdvisorSchemaObjectSummary      bool
	_databasemigrationserviceDescribeFleetAdvisorSchemas                  bool
	_databasemigrationserviceDescribeInstanceProfiles                     bool
	_databasemigrationserviceDescribeMetadataModel                        bool
	_databasemigrationserviceDescribeMetadataModelAssessments             bool
	_databasemigrationserviceDescribeMetadataModelChildren                bool
	_databasemigrationserviceDescribeMetadataModelConversions             bool
	_databasemigrationserviceDescribeMetadataModelCreations               bool
	_databasemigrationserviceDescribeMetadataModelExportsAsScript         bool
	_databasemigrationserviceDescribeMetadataModelExportsToTarget         bool
	_databasemigrationserviceDescribeMetadataModelImports                 bool
	_databasemigrationserviceDescribeMigrationProjects                    bool
	_databasemigrationserviceDescribeOrderableReplicationInstances        bool
	_databasemigrationserviceDescribePendingMaintenanceActions            bool
	_databasemigrationserviceDescribeRecommendationLimitations            bool
	_databasemigrationserviceDescribeRecommendations                      bool
	_databasemigrationserviceDescribeRefreshSchemasStatus                 bool
	_databasemigrationserviceDescribeReplicationConfigs                   bool
	_databasemigrationserviceDescribeReplicationInstanceTaskLogs          bool
	_databasemigrationserviceDescribeReplicationInstances                 bool
	_databasemigrationserviceDescribeReplicationSubnetGroups              bool
	_databasemigrationserviceDescribeReplicationTableStatistics           bool
	_databasemigrationserviceDescribeReplicationTaskAssessmentResults     bool
	_databasemigrationserviceDescribeReplicationTaskAssessmentRuns        bool
	_databasemigrationserviceDescribeReplicationTaskIndividualAssessments bool
	_databasemigrationserviceDescribeReplicationTasks                     bool
	_databasemigrationserviceDescribeReplications                         bool
	_databasemigrationserviceDescribeSchemas                              bool
	_databasemigrationserviceDescribeTableStatistics                      bool
	_databasemigrationserviceExportMetadataModelAssessment                bool
	_databasemigrationserviceGetTargetSelectionRules                      bool
	_databasemigrationserviceImportCertificate                            bool
	_databasemigrationserviceListTagsForResource                          bool
	_databasemigrationserviceModifyConversionConfiguration                bool
	_databasemigrationserviceModifyDataMigration                          bool
	_databasemigrationserviceModifyDataProvider                           bool
	_databasemigrationserviceModifyEndpoint                               bool
	_databasemigrationserviceModifyEventSubscription                      bool
	_databasemigrationserviceModifyInstanceProfile                        bool
	_databasemigrationserviceModifyMigrationProject                       bool
	_databasemigrationserviceModifyReplicationConfig                      bool
	_databasemigrationserviceModifyReplicationInstance                    bool
	_databasemigrationserviceModifyReplicationSubnetGroup                 bool
	_databasemigrationserviceModifyReplicationTask                        bool
	_databasemigrationserviceMoveReplicationTask                          bool
	_databasemigrationserviceRebootReplicationInstance                    bool
	_databasemigrationserviceRefreshSchemas                               bool
	_databasemigrationserviceReloadReplicationTables                      bool
	_databasemigrationserviceReloadTables                                 bool
	_databasemigrationserviceRemoveTagsFromResource                       bool
	_databasemigrationserviceRunFleetAdvisorLsaAnalysis                   bool
	_databasemigrationserviceStartDataMigration                           bool
	_databasemigrationserviceStartExtensionPackAssociation                bool
	_databasemigrationserviceStartMetadataModelAssessment                 bool
	_databasemigrationserviceStartMetadataModelConversion                 bool
	_databasemigrationserviceStartMetadataModelCreation                   bool
	_databasemigrationserviceStartMetadataModelExportAsScript             bool
	_databasemigrationserviceStartMetadataModelExportToTarget             bool
	_databasemigrationserviceStartMetadataModelImport                     bool
	_databasemigrationserviceStartRecommendations                         bool
	_databasemigrationserviceStartReplication                             bool
	_databasemigrationserviceStartReplicationTask                         bool
	_databasemigrationserviceStartReplicationTaskAssessment               bool
	_databasemigrationserviceStartReplicationTaskAssessmentRun            bool
	_databasemigrationserviceStopDataMigration                            bool
	_databasemigrationserviceStopReplication                              bool
	_databasemigrationserviceStopReplicationTask                          bool
	_databasemigrationserviceTestConnection                               bool
	_databasemigrationserviceUpdateSubscriptionsToEventBridge             bool

	_databasemigrationserviceAllocatedStorage                      string
	_databasemigrationserviceAllowMajorVersionUpgrade              string
	_databasemigrationserviceApplyAction                           string
	_databasemigrationserviceApplyImmediately                      string
	_databasemigrationserviceAssessmentReportTypes                 string
	_databasemigrationserviceAssessmentRunName                     string
	_databasemigrationserviceAutoMinorVersionUpgrade               string
	_databasemigrationserviceAvailabilityZone                      string
	_databasemigrationserviceCdcStartPosition                      string
	_databasemigrationserviceCdcStartTime                          string
	_databasemigrationserviceCdcStopPosition                       string
	_databasemigrationserviceCertificateArn                        string
	_databasemigrationserviceCertificateIdentifier                 string
	_databasemigrationserviceCertificatePem                        string
	_databasemigrationserviceCertificateWallet                     string
	_databasemigrationserviceCollectorName                         string
	_databasemigrationserviceCollectorReferencedId                 string
	_databasemigrationserviceComputeConfig                         string
	_databasemigrationserviceConversionConfiguration               string
	_databasemigrationserviceData                                  string
	_databasemigrationserviceDataMigrationIdentifier               string
	_databasemigrationserviceDataMigrationName                     string
	_databasemigrationserviceDataMigrationType                     string
	_databasemigrationserviceDataProviderIdentifier                string
	_databasemigrationserviceDataProviderName                      string
	_databasemigrationserviceDatabaseId                            string
	_databasemigrationserviceDatabaseIds                           []string
	_databasemigrationserviceDatabaseName                          string
	_databasemigrationserviceDescription                           string
	_databasemigrationserviceDmsTransferSettings                   string
	_databasemigrationserviceDnsNameServers                        string
	_databasemigrationserviceDocDbSettings                         string
	_databasemigrationserviceDuration                              string
	_databasemigrationserviceDynamoDbSettings                      string
	_databasemigrationserviceElasticsearchSettings                 string
	_databasemigrationserviceEnableCloudwatchLogs                  string
	_databasemigrationserviceEnabled                               string
	_databasemigrationserviceEndTime                               string
	_databasemigrationserviceEndpointArn                           string
	_databasemigrationserviceEndpointIdentifier                    string
	_databasemigrationserviceEndpointType                          string
	_databasemigrationserviceEngine                                string
	_databasemigrationserviceEngineName                            string
	_databasemigrationserviceEngineVersion                         string
	_databasemigrationserviceEventCategories                       []string
	_databasemigrationserviceExactSettings                         string
	_databasemigrationserviceExclude                               []string
	_databasemigrationserviceExternalTableDefinition               string
	_databasemigrationserviceExtraConnectionAttributes             string
	_databasemigrationserviceFileName                              string
	_databasemigrationserviceFilters                               string
	_databasemigrationserviceForceFailover                         string
	_databasemigrationserviceForceMove                             string
	_databasemigrationserviceForcePlannedFailover                  string
	_databasemigrationserviceGcpMySQLSettings                      string
	_databasemigrationserviceIBMDb2Settings                        string
	_databasemigrationserviceIncludeOnly                           []string
	_databasemigrationserviceInstanceProfileIdentifier             string
	_databasemigrationserviceInstanceProfileName                   string
	_databasemigrationserviceKafkaSettings                         string
	_databasemigrationserviceKerberosAuthenticationSettings        string
	_databasemigrationserviceKinesisSettings                       string
	_databasemigrationserviceKmsKeyArn                             string
	_databasemigrationserviceKmsKeyId                              string
	_databasemigrationserviceMarker                                string
	_databasemigrationserviceMaxRecords                            string
	_databasemigrationserviceMetadataModelName                     string
	_databasemigrationserviceMicrosoftSQLServerSettings            string
	_databasemigrationserviceMigrationProjectIdentifier            string
	_databasemigrationserviceMigrationProjectName                  string
	_databasemigrationserviceMigrationType                         string
	_databasemigrationserviceMongoDbSettings                       string
	_databasemigrationserviceMultiAZ                               string
	_databasemigrationserviceMySQLSettings                         string
	_databasemigrationserviceNeptuneSettings                       string
	_databasemigrationserviceNetworkType                           string
	_databasemigrationserviceNextToken                             string
	_databasemigrationserviceNumberOfJobs                          string
	_databasemigrationserviceOptInType                             string
	_databasemigrationserviceOracleSettings                        string
	_databasemigrationserviceOrigin                                string
	_databasemigrationserviceOverwriteExtensionPack                string
	_databasemigrationservicePassword                              string
	_databasemigrationservicePort                                  string
	_databasemigrationservicePostgreSQLSettings                    string
	_databasemigrationservicePreferredMaintenanceWindow            string
	_databasemigrationservicePremigrationAssessmentSettings        string
	_databasemigrationserviceProperties                            string
	_databasemigrationservicePubliclyAccessible                    string
	_databasemigrationserviceRedisSettings                         string
	_databasemigrationserviceRedshiftSettings                      string
	_databasemigrationserviceRefresh                               string
	_databasemigrationserviceReloadOption                          string
	_databasemigrationserviceReplicationConfigArn                  string
	_databasemigrationserviceReplicationConfigIdentifier           string
	_databasemigrationserviceReplicationInstanceArn                string
	_databasemigrationserviceReplicationInstanceClass              string
	_databasemigrationserviceReplicationInstanceIdentifier         string
	_databasemigrationserviceReplicationSettings                   string
	_databasemigrationserviceReplicationSubnetGroupDescription     string
	_databasemigrationserviceReplicationSubnetGroupIdentifier      string
	_databasemigrationserviceReplicationTaskArn                    string
	_databasemigrationserviceReplicationTaskAssessmentRunArn       string
	_databasemigrationserviceReplicationTaskIdentifier             string
	_databasemigrationserviceReplicationTaskSettings               string
	_databasemigrationserviceReplicationType                       string
	_databasemigrationserviceRequestIdentifier                     string
	_databasemigrationserviceResourceArn                           string
	_databasemigrationserviceResourceArnList                       []string
	_databasemigrationserviceResourceIdentifier                    string
	_databasemigrationserviceResultEncryptionMode                  string
	_databasemigrationserviceResultKmsKeyArn                       string
	_databasemigrationserviceResultLocationBucket                  string
	_databasemigrationserviceResultLocationFolder                  string
	_databasemigrationserviceS3BucketName                          string
	_databasemigrationserviceS3Settings                            string
	_databasemigrationserviceSchemaConversionApplicationAttributes string
	_databasemigrationserviceSelectionRules                        string
	_databasemigrationserviceServerName                            string
	_databasemigrationserviceServiceAccessRoleArn                  string
	_databasemigrationserviceSettings                              string
	_databasemigrationserviceSnsTopicArn                           string
	_databasemigrationserviceSourceDataProviderDescriptors         string
	_databasemigrationserviceSourceDataSettings                    string
	_databasemigrationserviceSourceEndpointArn                     string
	_databasemigrationserviceSourceEngineName                      string
	_databasemigrationserviceSourceIdentifier                      string
	_databasemigrationserviceSourceIds                             []string
	_databasemigrationserviceSourceType                            string
	_databasemigrationserviceSslMode                               string
	_databasemigrationserviceStartReplicationTaskType              string
	_databasemigrationserviceStartReplicationType                  string
	_databasemigrationserviceStartTime                             string
	_databasemigrationserviceStartType                             string
	_databasemigrationserviceSubnetGroupIdentifier                 string
	_databasemigrationserviceSubnetIds                             []string
	_databasemigrationserviceSubscriptionName                      string
	_databasemigrationserviceSupplementalSettings                  string
	_databasemigrationserviceSybaseSettings                        string
	_databasemigrationserviceTableMappings                         string
	_databasemigrationserviceTablesToReload                        string
	_databasemigrationserviceTagKeys                               []string
	_databasemigrationserviceTags                                  string
	_databasemigrationserviceTargetDataProviderDescriptors         string
	_databasemigrationserviceTargetDataSettings                    string
	_databasemigrationserviceTargetEndpointArn                     string
	_databasemigrationserviceTargetEngineName                      string
	_databasemigrationserviceTargetReplicationInstanceArn          string
	_databasemigrationserviceTaskData                              string
	_databasemigrationserviceTimestreamSettings                    string
	_databasemigrationserviceTransformationRules                   string
	_databasemigrationserviceUsername                              string
	_databasemigrationserviceVirtual                               string
	_databasemigrationserviceVpcSecurityGroupIds                   []string
	_databasemigrationserviceVpcSecurityGroups                     []string
	_databasemigrationserviceWithoutSettings                       string
	_databasemigrationserviceWithoutStatistics                     string
)

// Adds metadata tags to an DMS resource, including replication instance,
// endpoint, subnet group, and migration task. These tags can also be used with
// cost allocation reporting to track cost associated with DMS resources, or used
// in a Condition statement in an IAM policy for DMS. For more information, see [Tag]Tag
// data type description.
//
// [Tag]: https://docs.aws.amazon.com/dms/latest/APIReference/API_Tag.html
func databasemigrationservice_AddTagsToResource(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.AddTagsToResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_databasemigrationserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_databasemigrationserviceResourceArn)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
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

// Applies a pending maintenance action to a resource (for example, to a
// replication instance).
func databasemigrationservice_ApplyPendingMaintenanceAction(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ApplyPendingMaintenanceActionInput{
		// ApplyAction: *string, // Required
		// OptInType: *string, // Required
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceApplyAction) > 0 {
		input.ApplyAction = aws.String(_databasemigrationserviceApplyAction)
	}
	if len(_databasemigrationserviceOptInType) > 0 {
		input.OptInType = aws.String(_databasemigrationserviceOptInType)
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}

	if resp, err := client.ApplyPendingMaintenanceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Starts the analysis of up to 20 source databases to recommend target engines
// for each source database. This is a batch version of [StartRecommendations].
//
// The result of analysis of each source database is reported individually in the
// response. Because the batch request can result in a combination of successful
// and unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
// [StartRecommendations]: https://docs.aws.amazon.com/dms/latest/APIReference/API_StartRecommendations.html
func databasemigrationservice_BatchStartRecommendations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.BatchStartRecommendationsInput{}

	if len(_databasemigrationserviceData) > 0 {
		if err := assignInputField(input, "Data", _databasemigrationserviceData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchStartRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a single metadata model conversion operation that was started with
// StartMetadataModelConversion .
func databasemigrationservice_CancelMetadataModelConversion(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CancelMetadataModelConversionInput{
		// MigrationProjectIdentifier: *string, // Required
		// RequestIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceRequestIdentifier) > 0 {
		input.RequestIdentifier = aws.String(_databasemigrationserviceRequestIdentifier)
	}

	if resp, err := client.CancelMetadataModelConversion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a single metadata model creation operation that was started with
// StartMetadataModelCreation .
func databasemigrationservice_CancelMetadataModelCreation(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CancelMetadataModelCreationInput{
		// MigrationProjectIdentifier: *string, // Required
		// RequestIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceRequestIdentifier) > 0 {
		input.RequestIdentifier = aws.String(_databasemigrationserviceRequestIdentifier)
	}

	if resp, err := client.CancelMetadataModelCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a single premigration assessment run.
// This operation prevents any individual assessments from running if they haven't
// started running. It also attempts to cancel any individual assessments that are
// currently running.
func databasemigrationservice_CancelReplicationTaskAssessmentRun(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CancelReplicationTaskAssessmentRunInput{
		// ReplicationTaskAssessmentRunArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskAssessmentRunArn) > 0 {
		input.ReplicationTaskAssessmentRunArn = aws.String(_databasemigrationserviceReplicationTaskAssessmentRunArn)
	}

	if resp, err := client.CancelReplicationTaskAssessmentRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data migration using the provided settings.
func databasemigrationservice_CreateDataMigration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateDataMigrationInput{
		// DataMigrationType: types.MigrationTypeValue, // Required
		// MigrationProjectIdentifier: *string, // Required
		// ServiceAccessRoleArn: *string, // Required
	}

	if len(_databasemigrationserviceDataMigrationType) > 0 {
		if err := assignInputField(input, "DataMigrationType", _databasemigrationserviceDataMigrationType); err != nil {
			log.Errorf("invalid --data-migration-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceServiceAccessRoleArn) > 0 {
		input.ServiceAccessRoleArn = aws.String(_databasemigrationserviceServiceAccessRoleArn)
	}
	if len(_databasemigrationserviceDataMigrationName) > 0 {
		input.DataMigrationName = aws.String(_databasemigrationserviceDataMigrationName)
	}
	if len(_databasemigrationserviceEnableCloudwatchLogs) > 0 {
		if err := assignInputField(input, "EnableCloudwatchLogs", _databasemigrationserviceEnableCloudwatchLogs); err != nil {
			log.Errorf("invalid --enable-cloudwatch-logs: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNumberOfJobs) > 0 {
		if err := assignInputField(input, "NumberOfJobs", _databasemigrationserviceNumberOfJobs); err != nil {
			log.Errorf("invalid --number-of-jobs: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceSourceDataSettings) > 0 {
		if err := assignInputField(input, "SourceDataSettings", _databasemigrationserviceSourceDataSettings); err != nil {
			log.Errorf("invalid --source-data-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTargetDataSettings) > 0 {
		if err := assignInputField(input, "TargetDataSettings", _databasemigrationserviceTargetDataSettings); err != nil {
			log.Errorf("invalid --target-data-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data provider using the provided settings. A data provider stores a
// data store type and location information about your database.
func databasemigrationservice_CreateDataProvider(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateDataProviderInput{
		// Engine: *string, // Required
		// Settings: types.DataProviderSettings, // Required
	}

	if len(_databasemigrationserviceEngine) > 0 {
		input.Engine = aws.String(_databasemigrationserviceEngine)
	}
	if len(_databasemigrationserviceSettings) > 0 {
		if err := assignInputField(input, "Settings", _databasemigrationserviceSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceDataProviderName) > 0 {
		input.DataProviderName = aws.String(_databasemigrationserviceDataProviderName)
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceVirtual) > 0 {
		if err := assignInputField(input, "Virtual", _databasemigrationserviceVirtual); err != nil {
			log.Errorf("invalid --virtual: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an endpoint using the provided settings.
// For a MySQL source or target endpoint, don't explicitly specify the database
// using the DatabaseName request parameter on the CreateEndpoint API call.
// Specifying DatabaseName when you create a MySQL endpoint replicates all the
// task tables to this single database. For MySQL endpoints, you specify the
// database only when you specify the schema in the table-mapping rules of the DMS
// task.
func databasemigrationservice_CreateEndpoint(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateEndpointInput{
		// EndpointIdentifier: *string, // Required
		// EndpointType: types.ReplicationEndpointTypeValue, // Required
		// EngineName: *string, // Required
	}

	if len(_databasemigrationserviceEndpointIdentifier) > 0 {
		input.EndpointIdentifier = aws.String(_databasemigrationserviceEndpointIdentifier)
	}
	if len(_databasemigrationserviceEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _databasemigrationserviceEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEngineName) > 0 {
		input.EngineName = aws.String(_databasemigrationserviceEngineName)
	}
	if len(_databasemigrationserviceCertificateArn) > 0 {
		input.CertificateArn = aws.String(_databasemigrationserviceCertificateArn)
	}
	if len(_databasemigrationserviceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_databasemigrationserviceDatabaseName)
	}
	if len(_databasemigrationserviceDmsTransferSettings) > 0 {
		if err := assignInputField(input, "DmsTransferSettings", _databasemigrationserviceDmsTransferSettings); err != nil {
			log.Errorf("invalid --dms-transfer-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceDocDbSettings) > 0 {
		if err := assignInputField(input, "DocDbSettings", _databasemigrationserviceDocDbSettings); err != nil {
			log.Errorf("invalid --doc-db-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceDynamoDbSettings) > 0 {
		if err := assignInputField(input, "DynamoDbSettings", _databasemigrationserviceDynamoDbSettings); err != nil {
			log.Errorf("invalid --dynamo-db-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceElasticsearchSettings) > 0 {
		if err := assignInputField(input, "ElasticsearchSettings", _databasemigrationserviceElasticsearchSettings); err != nil {
			log.Errorf("invalid --elasticsearch-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceExternalTableDefinition) > 0 {
		input.ExternalTableDefinition = aws.String(_databasemigrationserviceExternalTableDefinition)
	}
	if len(_databasemigrationserviceExtraConnectionAttributes) > 0 {
		input.ExtraConnectionAttributes = aws.String(_databasemigrationserviceExtraConnectionAttributes)
	}
	if len(_databasemigrationserviceGcpMySQLSettings) > 0 {
		if err := assignInputField(input, "GcpMySQLSettings", _databasemigrationserviceGcpMySQLSettings); err != nil {
			log.Errorf("invalid --gcp-my-sql-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceIBMDb2Settings) > 0 {
		if err := assignInputField(input, "IBMDb2Settings", _databasemigrationserviceIBMDb2Settings); err != nil {
			log.Errorf("invalid --ibmdb2-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKafkaSettings) > 0 {
		if err := assignInputField(input, "KafkaSettings", _databasemigrationserviceKafkaSettings); err != nil {
			log.Errorf("invalid --kafka-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKinesisSettings) > 0 {
		if err := assignInputField(input, "KinesisSettings", _databasemigrationserviceKinesisSettings); err != nil {
			log.Errorf("invalid --kinesis-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_databasemigrationserviceKmsKeyId)
	}
	if len(_databasemigrationserviceMicrosoftSQLServerSettings) > 0 {
		if err := assignInputField(input, "MicrosoftSQLServerSettings", _databasemigrationserviceMicrosoftSQLServerSettings); err != nil {
			log.Errorf("invalid --microsoft-sql-server-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMongoDbSettings) > 0 {
		if err := assignInputField(input, "MongoDbSettings", _databasemigrationserviceMongoDbSettings); err != nil {
			log.Errorf("invalid --mongo-db-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMySQLSettings) > 0 {
		if err := assignInputField(input, "MySQLSettings", _databasemigrationserviceMySQLSettings); err != nil {
			log.Errorf("invalid --my-sql-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNeptuneSettings) > 0 {
		if err := assignInputField(input, "NeptuneSettings", _databasemigrationserviceNeptuneSettings); err != nil {
			log.Errorf("invalid --neptune-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceOracleSettings) > 0 {
		if err := assignInputField(input, "OracleSettings", _databasemigrationserviceOracleSettings); err != nil {
			log.Errorf("invalid --oracle-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationservicePassword) > 0 {
		input.Password = aws.String(_databasemigrationservicePassword)
	}
	if len(_databasemigrationservicePort) > 0 {
		if err := assignInputField(input, "Port", _databasemigrationservicePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationservicePostgreSQLSettings) > 0 {
		if err := assignInputField(input, "PostgreSQLSettings", _databasemigrationservicePostgreSQLSettings); err != nil {
			log.Errorf("invalid --postgre-sql-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceRedisSettings) > 0 {
		if err := assignInputField(input, "RedisSettings", _databasemigrationserviceRedisSettings); err != nil {
			log.Errorf("invalid --redis-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceRedshiftSettings) > 0 {
		if err := assignInputField(input, "RedshiftSettings", _databasemigrationserviceRedshiftSettings); err != nil {
			log.Errorf("invalid --redshift-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_databasemigrationserviceResourceIdentifier)
	}
	if len(_databasemigrationserviceS3Settings) > 0 {
		if err := assignInputField(input, "S3Settings", _databasemigrationserviceS3Settings); err != nil {
			log.Errorf("invalid --s3-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceServerName) > 0 {
		input.ServerName = aws.String(_databasemigrationserviceServerName)
	}
	if len(_databasemigrationserviceServiceAccessRoleArn) > 0 {
		input.ServiceAccessRoleArn = aws.String(_databasemigrationserviceServiceAccessRoleArn)
	}
	if len(_databasemigrationserviceSslMode) > 0 {
		if err := assignInputField(input, "SslMode", _databasemigrationserviceSslMode); err != nil {
			log.Errorf("invalid --ssl-mode: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSybaseSettings) > 0 {
		if err := assignInputField(input, "SybaseSettings", _databasemigrationserviceSybaseSettings); err != nil {
			log.Errorf("invalid --sybase-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTimestreamSettings) > 0 {
		if err := assignInputField(input, "TimestreamSettings", _databasemigrationserviceTimestreamSettings); err != nil {
			log.Errorf("invalid --timestream-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceUsername) > 0 {
		input.Username = aws.String(_databasemigrationserviceUsername)
	}

	if resp, err := client.CreateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an DMS event notification subscription.
// You can specify the type of source ( SourceType ) you want to be notified of,
// provide a list of DMS source IDs ( SourceIds ) that triggers the events, and
// provide a list of event categories ( EventCategories ) for events you want to be
// notified of. If you specify both the SourceType and SourceIds , such as
// SourceType = replication-instance and SourceIdentifier = my-replinstance , you
// will be notified of all the replication instance events for the specified
// source. If you specify a SourceType but don't specify a SourceIdentifier , you
// receive notice of the events for that source type for all your DMS sources. If
// you don't specify either SourceType nor SourceIdentifier , you will be notified
// of events generated from all DMS sources belonging to your customer account.
//
// For more information about DMS events, see [Working with Events and Notifications] in the Database Migration Service
// User Guide.
//
// [Working with Events and Notifications]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html
func databasemigrationservice_CreateEventSubscription(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateEventSubscriptionInput{
		// SnsTopicArn: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_databasemigrationserviceSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_databasemigrationserviceSnsTopicArn)
	}
	if len(_databasemigrationserviceSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_databasemigrationserviceSubscriptionName)
	}
	if len(_databasemigrationserviceEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _databasemigrationserviceEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _databasemigrationserviceEventCategories...)
	}
	if len(_databasemigrationserviceSourceIds) > 0 {
		input.SourceIds = append([]string(nil), _databasemigrationserviceSourceIds...)
	}
	if len(_databasemigrationserviceSourceType) > 0 {
		input.SourceType = aws.String(_databasemigrationserviceSourceType)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Creates a Fleet Advisor collector using the specified parameters.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_CreateFleetAdvisorCollector(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateFleetAdvisorCollectorInput{
		// CollectorName: *string, // Required
		// S3BucketName: *string, // Required
		// ServiceAccessRoleArn: *string, // Required
	}

	if len(_databasemigrationserviceCollectorName) > 0 {
		input.CollectorName = aws.String(_databasemigrationserviceCollectorName)
	}
	if len(_databasemigrationserviceS3BucketName) > 0 {
		input.S3BucketName = aws.String(_databasemigrationserviceS3BucketName)
	}
	if len(_databasemigrationserviceServiceAccessRoleArn) > 0 {
		input.ServiceAccessRoleArn = aws.String(_databasemigrationserviceServiceAccessRoleArn)
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}

	if resp, err := client.CreateFleetAdvisorCollector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the instance profile using the specified parameters.
func databasemigrationservice_CreateInstanceProfile(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateInstanceProfileInput{}

	if len(_databasemigrationserviceAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_databasemigrationserviceAvailabilityZone)
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}
	if len(_databasemigrationserviceInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_databasemigrationserviceInstanceProfileName)
	}
	if len(_databasemigrationserviceKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_databasemigrationserviceKmsKeyArn)
	}
	if len(_databasemigrationserviceNetworkType) > 0 {
		input.NetworkType = aws.String(_databasemigrationserviceNetworkType)
	}
	if len(_databasemigrationservicePubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _databasemigrationservicePubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSubnetGroupIdentifier) > 0 {
		input.SubnetGroupIdentifier = aws.String(_databasemigrationserviceSubnetGroupIdentifier)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceVpcSecurityGroups) > 0 {
		input.VpcSecurityGroups = append([]string(nil), _databasemigrationserviceVpcSecurityGroups...)
	}

	if resp, err := client.CreateInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the migration project using the specified parameters.
// You can run this action only after you create an instance profile and data
// providers using [CreateInstanceProfile]and [CreateDataProvider].
//
// [CreateDataProvider]: https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateDataProvider.html
// [CreateInstanceProfile]: https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateInstanceProfile.html
func databasemigrationservice_CreateMigrationProject(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateMigrationProjectInput{
		// InstanceProfileIdentifier: *string, // Required
		// SourceDataProviderDescriptors: []types.DataProviderDescriptorDefinition, // Required
		// TargetDataProviderDescriptors: []types.DataProviderDescriptorDefinition, // Required
	}

	if len(_databasemigrationserviceInstanceProfileIdentifier) > 0 {
		input.InstanceProfileIdentifier = aws.String(_databasemigrationserviceInstanceProfileIdentifier)
	}
	if len(_databasemigrationserviceSourceDataProviderDescriptors) > 0 {
		if err := assignInputField(input, "SourceDataProviderDescriptors", _databasemigrationserviceSourceDataProviderDescriptors); err != nil {
			log.Errorf("invalid --source-data-provider-descriptors: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTargetDataProviderDescriptors) > 0 {
		if err := assignInputField(input, "TargetDataProviderDescriptors", _databasemigrationserviceTargetDataProviderDescriptors); err != nil {
			log.Errorf("invalid --target-data-provider-descriptors: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}
	if len(_databasemigrationserviceMigrationProjectName) > 0 {
		input.MigrationProjectName = aws.String(_databasemigrationserviceMigrationProjectName)
	}
	if len(_databasemigrationserviceSchemaConversionApplicationAttributes) > 0 {
		if err := assignInputField(input, "SchemaConversionApplicationAttributes", _databasemigrationserviceSchemaConversionApplicationAttributes); err != nil {
			log.Errorf("invalid --schema-conversion-application-attributes: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTransformationRules) > 0 {
		input.TransformationRules = aws.String(_databasemigrationserviceTransformationRules)
	}

	if resp, err := client.CreateMigrationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration that you can later provide to configure and start an
// DMS Serverless replication. You can also provide options to validate the
// configuration inputs before you start the replication.
func databasemigrationservice_CreateReplicationConfig(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateReplicationConfigInput{
		// ComputeConfig: *types.ComputeConfig, // Required
		// ReplicationConfigIdentifier: *string, // Required
		// ReplicationType: types.MigrationTypeValue, // Required
		// SourceEndpointArn: *string, // Required
		// TableMappings: *string, // Required
		// TargetEndpointArn: *string, // Required
	}

	if len(_databasemigrationserviceComputeConfig) > 0 {
		if err := assignInputField(input, "ComputeConfig", _databasemigrationserviceComputeConfig); err != nil {
			log.Errorf("invalid --compute-config: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationConfigIdentifier) > 0 {
		input.ReplicationConfigIdentifier = aws.String(_databasemigrationserviceReplicationConfigIdentifier)
	}
	if len(_databasemigrationserviceReplicationType) > 0 {
		if err := assignInputField(input, "ReplicationType", _databasemigrationserviceReplicationType); err != nil {
			log.Errorf("invalid --replication-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSourceEndpointArn) > 0 {
		input.SourceEndpointArn = aws.String(_databasemigrationserviceSourceEndpointArn)
	}
	if len(_databasemigrationserviceTableMappings) > 0 {
		input.TableMappings = aws.String(_databasemigrationserviceTableMappings)
	}
	if len(_databasemigrationserviceTargetEndpointArn) > 0 {
		input.TargetEndpointArn = aws.String(_databasemigrationserviceTargetEndpointArn)
	}
	if len(_databasemigrationserviceReplicationSettings) > 0 {
		input.ReplicationSettings = aws.String(_databasemigrationserviceReplicationSettings)
	}
	if len(_databasemigrationserviceResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_databasemigrationserviceResourceIdentifier)
	}
	if len(_databasemigrationserviceSupplementalSettings) > 0 {
		input.SupplementalSettings = aws.String(_databasemigrationserviceSupplementalSettings)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReplicationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the replication instance using the specified parameters.
// DMS requires that your account have certain roles with appropriate permissions
// before you can create a replication instance. For information on the required
// roles, see [Creating the IAM Roles to Use With the CLI and DMS API]. For information on the required permissions, see [IAM Permissions Needed to Use DMS].
//
// If you don't specify a version when creating a replication instance, DMS will
// create the instance using the default engine version. For information about the
// default engine version, see [Release Notes].
//
// [Release Notes]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReleaseNotes.html
// [Creating the IAM Roles to Use With the CLI and DMS API]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.APIRole
// [IAM Permissions Needed to Use DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.IAMPermissions
func databasemigrationservice_CreateReplicationInstance(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateReplicationInstanceInput{
		// ReplicationInstanceClass: *string, // Required
		// ReplicationInstanceIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceReplicationInstanceClass) > 0 {
		input.ReplicationInstanceClass = aws.String(_databasemigrationserviceReplicationInstanceClass)
	}
	if len(_databasemigrationserviceReplicationInstanceIdentifier) > 0 {
		input.ReplicationInstanceIdentifier = aws.String(_databasemigrationserviceReplicationInstanceIdentifier)
	}
	if len(_databasemigrationserviceAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _databasemigrationserviceAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _databasemigrationserviceAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_databasemigrationserviceAvailabilityZone)
	}
	if len(_databasemigrationserviceDnsNameServers) > 0 {
		input.DnsNameServers = aws.String(_databasemigrationserviceDnsNameServers)
	}
	if len(_databasemigrationserviceEngineVersion) > 0 {
		input.EngineVersion = aws.String(_databasemigrationserviceEngineVersion)
	}
	if len(_databasemigrationserviceKerberosAuthenticationSettings) > 0 {
		if err := assignInputField(input, "KerberosAuthenticationSettings", _databasemigrationserviceKerberosAuthenticationSettings); err != nil {
			log.Errorf("invalid --kerberos-authentication-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_databasemigrationserviceKmsKeyId)
	}
	if len(_databasemigrationserviceMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _databasemigrationserviceMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNetworkType) > 0 {
		input.NetworkType = aws.String(_databasemigrationserviceNetworkType)
	}
	if len(_databasemigrationservicePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_databasemigrationservicePreferredMaintenanceWindow)
	}
	if len(_databasemigrationservicePubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _databasemigrationservicePubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationSubnetGroupIdentifier) > 0 {
		input.ReplicationSubnetGroupIdentifier = aws.String(_databasemigrationserviceReplicationSubnetGroupIdentifier)
	}
	if len(_databasemigrationserviceResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_databasemigrationserviceResourceIdentifier)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _databasemigrationserviceVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateReplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a replication subnet group given a list of the subnet IDs in a VPC.
// The VPC needs to have at least one subnet in at least two availability zones in
// the Amazon Web Services Region, otherwise the service will throw a
// ReplicationSubnetGroupDoesNotCoverEnoughAZs exception.
//
// If a replication subnet group exists in your Amazon Web Services account, the
// CreateReplicationSubnetGroup action returns the following error message: The
// Replication Subnet Group already exists. In this case, delete the existing
// replication subnet group. To do so, use the [DeleteReplicationSubnetGroup]action. Optionally, choose Subnet
// groups in the DMS console, then choose your subnet group. Next, choose Delete
// from Actions.
//
// [DeleteReplicationSubnetGroup]: https://docs.aws.amazon.com/en_us/dms/latest/APIReference/API_DeleteReplicationSubnetGroup.html
func databasemigrationservice_CreateReplicationSubnetGroup(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateReplicationSubnetGroupInput{
		// ReplicationSubnetGroupDescription: *string, // Required
		// ReplicationSubnetGroupIdentifier: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_databasemigrationserviceReplicationSubnetGroupDescription) > 0 {
		input.ReplicationSubnetGroupDescription = aws.String(_databasemigrationserviceReplicationSubnetGroupDescription)
	}
	if len(_databasemigrationserviceReplicationSubnetGroupIdentifier) > 0 {
		input.ReplicationSubnetGroupIdentifier = aws.String(_databasemigrationserviceReplicationSubnetGroupIdentifier)
	}
	if len(_databasemigrationserviceSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _databasemigrationserviceSubnetIds...)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReplicationSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a replication task using the specified parameters.
func databasemigrationservice_CreateReplicationTask(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.CreateReplicationTaskInput{
		// MigrationType: types.MigrationTypeValue, // Required
		// ReplicationInstanceArn: *string, // Required
		// ReplicationTaskIdentifier: *string, // Required
		// SourceEndpointArn: *string, // Required
		// TableMappings: *string, // Required
		// TargetEndpointArn: *string, // Required
	}

	if len(_databasemigrationserviceMigrationType) > 0 {
		if err := assignInputField(input, "MigrationType", _databasemigrationserviceMigrationType); err != nil {
			log.Errorf("invalid --migration-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}
	if len(_databasemigrationserviceReplicationTaskIdentifier) > 0 {
		input.ReplicationTaskIdentifier = aws.String(_databasemigrationserviceReplicationTaskIdentifier)
	}
	if len(_databasemigrationserviceSourceEndpointArn) > 0 {
		input.SourceEndpointArn = aws.String(_databasemigrationserviceSourceEndpointArn)
	}
	if len(_databasemigrationserviceTableMappings) > 0 {
		input.TableMappings = aws.String(_databasemigrationserviceTableMappings)
	}
	if len(_databasemigrationserviceTargetEndpointArn) > 0 {
		input.TargetEndpointArn = aws.String(_databasemigrationserviceTargetEndpointArn)
	}
	if len(_databasemigrationserviceCdcStartPosition) > 0 {
		input.CdcStartPosition = aws.String(_databasemigrationserviceCdcStartPosition)
	}
	if len(_databasemigrationserviceCdcStartTime) > 0 {
		if err := assignInputField(input, "CdcStartTime", _databasemigrationserviceCdcStartTime); err != nil {
			log.Errorf("invalid --cdc-start-time: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceCdcStopPosition) > 0 {
		input.CdcStopPosition = aws.String(_databasemigrationserviceCdcStopPosition)
	}
	if len(_databasemigrationserviceReplicationTaskSettings) > 0 {
		input.ReplicationTaskSettings = aws.String(_databasemigrationserviceReplicationTaskSettings)
	}
	if len(_databasemigrationserviceResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_databasemigrationserviceResourceIdentifier)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTaskData) > 0 {
		input.TaskData = aws.String(_databasemigrationserviceTaskData)
	}

	if resp, err := client.CreateReplicationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified certificate.
func databasemigrationservice_DeleteCertificate(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteCertificateInput{
		// CertificateArn: *string, // Required
	}

	if len(_databasemigrationserviceCertificateArn) > 0 {
		input.CertificateArn = aws.String(_databasemigrationserviceCertificateArn)
	}

	if resp, err := client.DeleteCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the connection between a replication instance and an endpoint.
func databasemigrationservice_DeleteConnection(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteConnectionInput{
		// EndpointArn: *string, // Required
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified data migration.
func databasemigrationservice_DeleteDataMigration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteDataMigrationInput{
		// DataMigrationIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceDataMigrationIdentifier) > 0 {
		input.DataMigrationIdentifier = aws.String(_databasemigrationserviceDataMigrationIdentifier)
	}

	if resp, err := client.DeleteDataMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified data provider.
// All migration projects associated with the data provider must be deleted or
// modified before you can delete the data provider.
func databasemigrationservice_DeleteDataProvider(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteDataProviderInput{
		// DataProviderIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceDataProviderIdentifier) > 0 {
		input.DataProviderIdentifier = aws.String(_databasemigrationserviceDataProviderIdentifier)
	}

	if resp, err := client.DeleteDataProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified endpoint.
// All tasks associated with the endpoint must be deleted before you can delete
// the endpoint.
func databasemigrationservice_DeleteEndpoint(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an DMS event subscription.
func databasemigrationservice_DeleteEventSubscription(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_databasemigrationserviceSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_databasemigrationserviceSubscriptionName)
	}

	if resp, err := client.DeleteEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Deletes the specified Fleet Advisor collector.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DeleteFleetAdvisorCollector(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteFleetAdvisorCollectorInput{
		// CollectorReferencedId: *string, // Required
	}

	if len(_databasemigrationserviceCollectorReferencedId) > 0 {
		input.CollectorReferencedId = aws.String(_databasemigrationserviceCollectorReferencedId)
	}

	if resp, err := client.DeleteFleetAdvisorCollector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Deletes the specified Fleet Advisor collector databases.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DeleteFleetAdvisorDatabases(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteFleetAdvisorDatabasesInput{
		// DatabaseIds: []string, // Required
	}

	if len(_databasemigrationserviceDatabaseIds) > 0 {
		input.DatabaseIds = append([]string(nil), _databasemigrationserviceDatabaseIds...)
	}

	if resp, err := client.DeleteFleetAdvisorDatabases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified instance profile.
// All migration projects associated with the instance profile must be deleted or
// modified before you can delete the instance profile.
func databasemigrationservice_DeleteInstanceProfile(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteInstanceProfileInput{
		// InstanceProfileIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceInstanceProfileIdentifier) > 0 {
		input.InstanceProfileIdentifier = aws.String(_databasemigrationserviceInstanceProfileIdentifier)
	}

	if resp, err := client.DeleteInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified migration project.
// The migration project must be closed before you can delete it.
func databasemigrationservice_DeleteMigrationProject(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteMigrationProjectInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}

	if resp, err := client.DeleteMigrationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an DMS Serverless replication configuration. This effectively
// deprovisions any and all replications that use this configuration. You can't
// delete the configuration for an DMS Serverless replication that is ongoing. You
// can delete the configuration when the replication is in a non-RUNNING and
// non-STARTING state.
func databasemigrationservice_DeleteReplicationConfig(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteReplicationConfigInput{
		// ReplicationConfigArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}

	if resp, err := client.DeleteReplicationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified replication instance.
// You must delete any migration tasks that are associated with the replication
// instance before you can delete it.
func databasemigrationservice_DeleteReplicationInstance(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteReplicationInstanceInput{
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}

	if resp, err := client.DeleteReplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subnet group.
func databasemigrationservice_DeleteReplicationSubnetGroup(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteReplicationSubnetGroupInput{
		// ReplicationSubnetGroupIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceReplicationSubnetGroupIdentifier) > 0 {
		input.ReplicationSubnetGroupIdentifier = aws.String(_databasemigrationserviceReplicationSubnetGroupIdentifier)
	}

	if resp, err := client.DeleteReplicationSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified replication task.
func databasemigrationservice_DeleteReplicationTask(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteReplicationTaskInput{
		// ReplicationTaskArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}

	if resp, err := client.DeleteReplicationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the record of a single premigration assessment run.
// This operation removes all metadata that DMS maintains about this assessment
// run. However, the operation leaves untouched all information about this
// assessment run that is stored in your Amazon S3 bucket.
func databasemigrationservice_DeleteReplicationTaskAssessmentRun(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DeleteReplicationTaskAssessmentRunInput{
		// ReplicationTaskAssessmentRunArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskAssessmentRunArn) > 0 {
		input.ReplicationTaskAssessmentRunArn = aws.String(_databasemigrationserviceReplicationTaskAssessmentRunArn)
	}

	if resp, err := client.DeleteReplicationTaskAssessmentRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the DMS attributes for a customer account. These attributes
// include DMS quotas for the account and a unique account identifier in a
// particular DMS region. DMS quotas include a list of resource quotas supported by
// the account, such as the number of replication instances allowed. The
// description for each resource quota, includes the quota name, current usage
// toward that quota, and the quota's maximum value. DMS uses the unique account
// identifier to name each artifact used by DMS in the given region.
//
// This command does not take any parameters.
func databasemigrationservice_DescribeAccountAttributes(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeAccountAttributesInput{}

	if resp, err := client.DescribeAccountAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of individual assessments that you can specify for a new
// premigration assessment run, given one or more parameters.
//
// If you specify an existing migration task, this operation provides the default
// individual assessments you can specify for that task. Otherwise, the specified
// parameters model elements of a possible migration task on which to base a
// premigration assessment run.
//
// To use these migration task modeling parameters, you must specify an existing
// replication instance, a source database engine, a target database engine, and a
// migration type. This combination of parameters potentially limits the default
// individual assessments available for an assessment run created for a
// corresponding migration task.
//
// If you specify no parameters, this operation provides a list of all possible
// individual assessments that you can specify for an assessment run. If you
// specify any one of the task modeling parameters, you must specify all of them or
// the operation cannot provide a list of individual assessments. The only
// parameter that you can specify alone is for an existing migration task. The
// specified task definition then determines the default list of individual
// assessments that you can specify in an assessment run for the task.
func databasemigrationservice_DescribeApplicableIndividualAssessments(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeApplicableIndividualAssessmentsInput{}

	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMigrationType) > 0 {
		if err := assignInputField(input, "MigrationType", _databasemigrationserviceMigrationType); err != nil {
			log.Errorf("invalid --migration-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}
	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceSourceEngineName) > 0 {
		input.SourceEngineName = aws.String(_databasemigrationserviceSourceEngineName)
	}
	if len(_databasemigrationserviceTargetEngineName) > 0 {
		input.TargetEngineName = aws.String(_databasemigrationserviceTargetEngineName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeApplicableIndividualAssessments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput
	p := databasemigrationservice.NewDescribeApplicableIndividualAssessmentsPaginator(client, input)
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

// Provides a description of the certificate.
func databasemigrationservice_DescribeCertificates(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeCertificatesInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
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

	var results []*databasemigrationservice.DescribeCertificatesOutput
	p := databasemigrationservice.NewDescribeCertificatesPaginator(client, input)
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

// Describes the status of the connections that have been made between the
// replication instance and an endpoint. Connections are created when you test an
// endpoint.
func databasemigrationservice_DescribeConnections(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeConnectionsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeConnectionsOutput
	p := databasemigrationservice.NewDescribeConnectionsPaginator(client, input)
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

// Returns configuration parameters for a schema conversion project.
func databasemigrationservice_DescribeConversionConfiguration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeConversionConfigurationInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}

	if resp, err := client.DescribeConversionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about data migrations.
func databasemigrationservice_DescribeDataMigrations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeDataMigrationsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceWithoutSettings) > 0 {
		if err := assignInputField(input, "WithoutSettings", _databasemigrationserviceWithoutSettings); err != nil {
			log.Errorf("invalid --without-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceWithoutStatistics) > 0 {
		if err := assignInputField(input, "WithoutStatistics", _databasemigrationserviceWithoutStatistics); err != nil {
			log.Errorf("invalid --without-statistics: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataMigrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeDataMigrationsOutput
	p := databasemigrationservice.NewDescribeDataMigrationsPaginator(client, input)
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

// Returns a paginated list of data providers for your account in the current
// region.
func databasemigrationservice_DescribeDataProviders(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeDataProvidersInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeDataProvidersOutput
	p := databasemigrationservice.NewDescribeDataProvidersPaginator(client, input)
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

// Returns information about the possible endpoint settings available when you
// create an endpoint for a specific database engine.
func databasemigrationservice_DescribeEndpointSettings(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEndpointSettingsInput{
		// EngineName: *string, // Required
	}

	if len(_databasemigrationserviceEngineName) > 0 {
		input.EngineName = aws.String(_databasemigrationserviceEngineName)
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEndpointSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeEndpointSettingsOutput
	p := databasemigrationservice.NewDescribeEndpointSettingsPaginator(client, input)
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

// Returns information about the type of endpoints available.
func databasemigrationservice_DescribeEndpointTypes(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEndpointTypesInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEndpointTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeEndpointTypesOutput
	p := databasemigrationservice.NewDescribeEndpointTypesPaginator(client, input)
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

// Returns information about the endpoints for your account in the current region.
func databasemigrationservice_DescribeEndpoints(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEndpointsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeEndpointsOutput
	p := databasemigrationservice.NewDescribeEndpointsPaginator(client, input)
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

// Returns information about the replication instance versions used in the project.
func databasemigrationservice_DescribeEngineVersions(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEngineVersionsInput{}

	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
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

	var results []*databasemigrationservice.DescribeEngineVersionsOutput
	p := databasemigrationservice.NewDescribeEngineVersionsPaginator(client, input)
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

// Lists categories for all event source types, or, if specified, for a specified
// source type. You can see a list of the event categories and source types in [Working with Events and Notifications]in
// the Database Migration Service User Guide.
//
// [Working with Events and Notifications]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html
func databasemigrationservice_DescribeEventCategories(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEventCategoriesInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSourceType) > 0 {
		input.SourceType = aws.String(_databasemigrationserviceSourceType)
	}

	if resp, err := client.DescribeEventCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the event subscriptions for a customer account. The description of a
// subscription includes SubscriptionName , SNSTopicARN , CustomerID , SourceType ,
// SourceID , CreationTime , and Status .
//
// If you specify SubscriptionName , this action lists the description for that
// subscription.
func databasemigrationservice_DescribeEventSubscriptions(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEventSubscriptionsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_databasemigrationserviceSubscriptionName)
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

	var results []*databasemigrationservice.DescribeEventSubscriptionsOutput
	p := databasemigrationservice.NewDescribeEventSubscriptionsPaginator(client, input)
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

// Lists events for a given source identifier and source type. You can also
// specify a start and end time. For more information on DMS events, see [Working with Events and Notifications]in the
// Database Migration Service User Guide.
//
// [Working with Events and Notifications]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html
func databasemigrationservice_DescribeEvents(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeEventsInput{}

	if len(_databasemigrationserviceDuration) > 0 {
		if err := assignInputField(input, "Duration", _databasemigrationserviceDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _databasemigrationserviceEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _databasemigrationserviceEventCategories...)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_databasemigrationserviceSourceIdentifier)
	}
	if len(_databasemigrationserviceSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _databasemigrationserviceSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _databasemigrationserviceStartTime); err != nil {
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

	var results []*databasemigrationservice.DescribeEventsOutput
	p := databasemigrationservice.NewDescribeEventsPaginator(client, input)
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

// Returns a paginated list of extension pack associations for the specified
// migration project. An extension pack is an add-on module that emulates functions
// present in a source database that are required when converting objects to the
// target database.
func databasemigrationservice_DescribeExtensionPackAssociations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeExtensionPackAssociationsInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeExtensionPackAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeExtensionPackAssociationsOutput
	p := databasemigrationservice.NewDescribeExtensionPackAssociationsPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Returns a list of the Fleet Advisor collectors in your account.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeFleetAdvisorCollectors(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeFleetAdvisorCollectorsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFleetAdvisorCollectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput
	p := databasemigrationservice.NewDescribeFleetAdvisorCollectorsPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Returns a list of Fleet Advisor databases in your account.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeFleetAdvisorDatabases(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeFleetAdvisorDatabasesInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFleetAdvisorDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput
	p := databasemigrationservice.NewDescribeFleetAdvisorDatabasesPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Provides descriptions of large-scale assessment (LSA) analyses produced by your
// Fleet Advisor collectors.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeFleetAdvisorLsaAnalysis(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput{}

	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFleetAdvisorLsaAnalysis(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput
	p := databasemigrationservice.NewDescribeFleetAdvisorLsaAnalysisPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Provides descriptions of the schemas discovered by your Fleet Advisor
// collectors.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeFleetAdvisorSchemaObjectSummary(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFleetAdvisorSchemaObjectSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput
	p := databasemigrationservice.NewDescribeFleetAdvisorSchemaObjectSummaryPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Returns a list of schemas detected by Fleet Advisor Collectors in your account.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeFleetAdvisorSchemas(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeFleetAdvisorSchemasInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFleetAdvisorSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeFleetAdvisorSchemasOutput
	p := databasemigrationservice.NewDescribeFleetAdvisorSchemasPaginator(client, input)
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

// Returns a paginated list of instance profiles for your account in the current
// region.
func databasemigrationservice_DescribeInstanceProfiles(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeInstanceProfilesInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstanceProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeInstanceProfilesOutput
	p := databasemigrationservice.NewDescribeInstanceProfilesPaginator(client, input)
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

// Gets detailed information about the specified metadata model, including its
// definition and corresponding converted objects in the target database if
// applicable.
func databasemigrationservice_DescribeMetadataModel(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelInput{
		// MigrationProjectIdentifier: *string, // Required
		// Origin: types.OriginTypeValue, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceOrigin) > 0 {
		if err := assignInputField(input, "Origin", _databasemigrationserviceOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}

	if resp, err := client.DescribeMetadataModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of metadata model assessments for your account in the
// current region.
func databasemigrationservice_DescribeMetadataModelAssessments(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelAssessmentsInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelAssessments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelAssessmentsOutput
	p := databasemigrationservice.NewDescribeMetadataModelAssessmentsPaginator(client, input)
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

// Gets a list of child metadata models for the specified metadata model in the
// database hierarchy.
func databasemigrationservice_DescribeMetadataModelChildren(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelChildrenInput{
		// MigrationProjectIdentifier: *string, // Required
		// Origin: types.OriginTypeValue, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceOrigin) > 0 {
		if err := assignInputField(input, "Origin", _databasemigrationserviceOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelChildren(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelChildrenOutput
	p := databasemigrationservice.NewDescribeMetadataModelChildrenPaginator(client, input)
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

// Returns a paginated list of metadata model conversions for a migration project.
func databasemigrationservice_DescribeMetadataModelConversions(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelConversionsInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelConversions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelConversionsOutput
	p := databasemigrationservice.NewDescribeMetadataModelConversionsPaginator(client, input)
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

// Returns a paginated list of metadata model creation requests for a migration
// project.
func databasemigrationservice_DescribeMetadataModelCreations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelCreationsInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelCreations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelCreationsOutput
	p := databasemigrationservice.NewDescribeMetadataModelCreationsPaginator(client, input)
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

// Returns a paginated list of metadata model exports.
func databasemigrationservice_DescribeMetadataModelExportsAsScript(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelExportsAsScriptInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelExportsAsScript(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelExportsAsScriptOutput
	p := databasemigrationservice.NewDescribeMetadataModelExportsAsScriptPaginator(client, input)
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

// Returns a paginated list of metadata model exports.
func databasemigrationservice_DescribeMetadataModelExportsToTarget(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelExportsToTargetInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelExportsToTarget(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelExportsToTargetOutput
	p := databasemigrationservice.NewDescribeMetadataModelExportsToTargetPaginator(client, input)
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

// Returns a paginated list of metadata model imports.
func databasemigrationservice_DescribeMetadataModelImports(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMetadataModelImportsInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetadataModelImports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMetadataModelImportsOutput
	p := databasemigrationservice.NewDescribeMetadataModelImportsPaginator(client, input)
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

// Returns a paginated list of migration projects for your account in the current
// region.
func databasemigrationservice_DescribeMigrationProjects(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeMigrationProjectsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMigrationProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeMigrationProjectsOutput
	p := databasemigrationservice.NewDescribeMigrationProjectsPaginator(client, input)
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

// Returns information about the replication instance types that can be created in
// the specified region.
func databasemigrationservice_DescribeOrderableReplicationInstances(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeOrderableReplicationInstancesInput{}

	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrderableReplicationInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeOrderableReplicationInstancesOutput
	p := databasemigrationservice.NewDescribeOrderableReplicationInstancesPaginator(client, input)
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

// Returns a list of upcoming maintenance events for replication instances in your
// account in the current Region.
func databasemigrationservice_DescribePendingMaintenanceActions(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribePendingMaintenanceActionsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
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

	var results []*databasemigrationservice.DescribePendingMaintenanceActionsOutput
	p := databasemigrationservice.NewDescribePendingMaintenanceActionsPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Returns a paginated list of limitations for recommendations of target Amazon
// Web Services engines.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeRecommendationLimitations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeRecommendationLimitationsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRecommendationLimitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeRecommendationLimitationsOutput
	p := databasemigrationservice.NewDescribeRecommendationLimitationsPaginator(client, input)
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

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Returns a paginated list of target engine recommendations for your source
// databases.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_DescribeRecommendations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeRecommendationsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNextToken) > 0 {
		input.NextToken = aws.String(_databasemigrationserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeRecommendationsOutput
	p := databasemigrationservice.NewDescribeRecommendationsPaginator(client, input)
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

// Returns the status of the RefreshSchemas operation.
func databasemigrationservice_DescribeRefreshSchemasStatus(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeRefreshSchemasStatusInput{
		// EndpointArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}

	if resp, err := client.DescribeRefreshSchemasStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns one or more existing DMS Serverless replication configurations as a
// list of structures.
func databasemigrationservice_DescribeReplicationConfigs(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationConfigsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationConfigsOutput
	p := databasemigrationservice.NewDescribeReplicationConfigsPaginator(client, input)
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

// Returns information about the task logs for the specified task.
func databasemigrationservice_DescribeReplicationInstanceTaskLogs(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationInstanceTaskLogsInput{
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationInstanceTaskLogs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput
	p := databasemigrationservice.NewDescribeReplicationInstanceTaskLogsPaginator(client, input)
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

// Returns information about replication instances for your account in the current
// region.
func databasemigrationservice_DescribeReplicationInstances(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationInstancesInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationInstancesOutput
	p := databasemigrationservice.NewDescribeReplicationInstancesPaginator(client, input)
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

// Returns information about the replication subnet groups.
func databasemigrationservice_DescribeReplicationSubnetGroups(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationSubnetGroupsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationSubnetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationSubnetGroupsOutput
	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(client, input)
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

// Returns table and schema statistics for one or more provisioned replications
// that use a given DMS Serverless replication configuration.
func databasemigrationservice_DescribeReplicationTableStatistics(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationTableStatisticsInput{
		// ReplicationConfigArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationTableStatistics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationTableStatisticsOutput
	p := databasemigrationservice.NewDescribeReplicationTableStatisticsPaginator(client, input)
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

// Returns the task assessment results from the Amazon S3 bucket that DMS creates
// in your Amazon Web Services account. This action always returns the latest
// results.
//
// For more information about DMS task assessments, see [Creating a task assessment report] in the Database Migration
// Service User Guide.
//
// [Creating a task assessment report]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.AssessmentReport.html
func databasemigrationservice_DescribeReplicationTaskAssessmentResults(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput{}

	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationTaskAssessmentResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput
	p := databasemigrationservice.NewDescribeReplicationTaskAssessmentResultsPaginator(client, input)
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

// Returns a paginated list of premigration assessment runs based on filter
// settings.
//
// These filter settings can specify a combination of premigration assessment
// runs, migration tasks, replication instances, and assessment run status values.
//
// This operation doesn't return information about individual assessments. For
// this information, see the DescribeReplicationTaskIndividualAssessments
// operation.
func databasemigrationservice_DescribeReplicationTaskAssessmentRuns(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationTaskAssessmentRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput
	p := databasemigrationservice.NewDescribeReplicationTaskAssessmentRunsPaginator(client, input)
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

// Returns a paginated list of individual assessments based on filter settings.
// These filter settings can specify a combination of premigration assessment
// runs, migration tasks, and assessment status values.
func databasemigrationservice_DescribeReplicationTaskIndividualAssessments(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationTaskIndividualAssessments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput
	p := databasemigrationservice.NewDescribeReplicationTaskIndividualAssessmentsPaginator(client, input)
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

// Returns information about replication tasks for your account in the current
// region.
func databasemigrationservice_DescribeReplicationTasks(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationTasksInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceWithoutSettings) > 0 {
		if err := assignInputField(input, "WithoutSettings", _databasemigrationserviceWithoutSettings); err != nil {
			log.Errorf("invalid --without-settings: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationTasksOutput
	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(client, input)
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

// Provides details on replication progress by returning status information for
// one or more provisioned DMS Serverless replications.
func databasemigrationservice_DescribeReplications(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeReplicationsInput{}

	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeReplicationsOutput
	p := databasemigrationservice.NewDescribeReplicationsPaginator(client, input)
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

// Returns information about the schema for the specified endpoint.
func databasemigrationservice_DescribeSchemas(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeSchemasInput{
		// EndpointArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeSchemasOutput
	p := databasemigrationservice.NewDescribeSchemasPaginator(client, input)
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

// Returns table statistics on the database migration task, including table name,
// rows inserted, rows updated, and rows deleted.
//
// Note that the "last updated" column the DMS console only indicates the time
// that DMS last updated the table statistics record for a table. It does not
// indicate the time of the last update to the table.
func databasemigrationservice_DescribeTableStatistics(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.DescribeTableStatisticsInput{
		// ReplicationTaskArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _databasemigrationserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMarker) > 0 {
		input.Marker = aws.String(_databasemigrationserviceMarker)
	}
	if len(_databasemigrationserviceMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _databasemigrationserviceMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeTableStatistics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*databasemigrationservice.DescribeTableStatisticsOutput
	p := databasemigrationservice.NewDescribeTableStatisticsPaginator(client, input)
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

// Saves a copy of a database migration assessment report to your Amazon S3
// bucket. DMS can save your assessment report as a comma-separated value (CSV) or
// a PDF file.
func databasemigrationservice_ExportMetadataModelAssessment(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ExportMetadataModelAssessmentInput{
		// MigrationProjectIdentifier: *string, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceAssessmentReportTypes) > 0 {
		if err := assignInputField(input, "AssessmentReportTypes", _databasemigrationserviceAssessmentReportTypes); err != nil {
			log.Errorf("invalid --assessment-report-types: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceFileName) > 0 {
		input.FileName = aws.String(_databasemigrationserviceFileName)
	}

	if resp, err := client.ExportMetadataModelAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Converts source selection rules into their target counterparts for schema
// conversion operations.
func databasemigrationservice_GetTargetSelectionRules(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.GetTargetSelectionRulesInput{
		// MigrationProjectIdentifier: *string, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}

	if resp, err := client.GetTargetSelectionRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads the specified certificate.
func databasemigrationservice_ImportCertificate(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ImportCertificateInput{
		// CertificateIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceCertificateIdentifier) > 0 {
		input.CertificateIdentifier = aws.String(_databasemigrationserviceCertificateIdentifier)
	}
	if len(_databasemigrationserviceCertificatePem) > 0 {
		input.CertificatePem = aws.String(_databasemigrationserviceCertificatePem)
	}
	if len(_databasemigrationserviceCertificateWallet) > 0 {
		if err := assignInputField(input, "CertificateWallet", _databasemigrationserviceCertificateWallet); err != nil {
			log.Errorf("invalid --certificate-wallet: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_databasemigrationserviceKmsKeyId)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all metadata tags attached to an DMS resource, including replication
// instance, endpoint, subnet group, and migration task. For more information, see [Tag]
// Tag data type description.
//
// [Tag]: https://docs.aws.amazon.com/dms/latest/APIReference/API_Tag.html
func databasemigrationservice_ListTagsForResource(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ListTagsForResourceInput{}

	if len(_databasemigrationserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_databasemigrationserviceResourceArn)
	}
	if len(_databasemigrationserviceResourceArnList) > 0 {
		input.ResourceArnList = append([]string(nil), _databasemigrationserviceResourceArnList...)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified schema conversion configuration using the provided
// parameters.
func databasemigrationservice_ModifyConversionConfiguration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyConversionConfigurationInput{
		// ConversionConfiguration: *string, // Required
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceConversionConfiguration) > 0 {
		input.ConversionConfiguration = aws.String(_databasemigrationserviceConversionConfiguration)
	}
	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}

	if resp, err := client.ModifyConversionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing DMS data migration.
func databasemigrationservice_ModifyDataMigration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyDataMigrationInput{
		// DataMigrationIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceDataMigrationIdentifier) > 0 {
		input.DataMigrationIdentifier = aws.String(_databasemigrationserviceDataMigrationIdentifier)
	}
	if len(_databasemigrationserviceDataMigrationName) > 0 {
		input.DataMigrationName = aws.String(_databasemigrationserviceDataMigrationName)
	}
	if len(_databasemigrationserviceDataMigrationType) > 0 {
		if err := assignInputField(input, "DataMigrationType", _databasemigrationserviceDataMigrationType); err != nil {
			log.Errorf("invalid --data-migration-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEnableCloudwatchLogs) > 0 {
		if err := assignInputField(input, "EnableCloudwatchLogs", _databasemigrationserviceEnableCloudwatchLogs); err != nil {
			log.Errorf("invalid --enable-cloudwatch-logs: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNumberOfJobs) > 0 {
		if err := assignInputField(input, "NumberOfJobs", _databasemigrationserviceNumberOfJobs); err != nil {
			log.Errorf("invalid --number-of-jobs: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceServiceAccessRoleArn) > 0 {
		input.ServiceAccessRoleArn = aws.String(_databasemigrationserviceServiceAccessRoleArn)
	}
	if len(_databasemigrationserviceSourceDataSettings) > 0 {
		if err := assignInputField(input, "SourceDataSettings", _databasemigrationserviceSourceDataSettings); err != nil {
			log.Errorf("invalid --source-data-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTargetDataSettings) > 0 {
		if err := assignInputField(input, "TargetDataSettings", _databasemigrationserviceTargetDataSettings); err != nil {
			log.Errorf("invalid --target-data-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDataMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified data provider using the provided settings.
// You must remove the data provider from all migration projects before you can
// modify it.
func databasemigrationservice_ModifyDataProvider(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyDataProviderInput{
		// DataProviderIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceDataProviderIdentifier) > 0 {
		input.DataProviderIdentifier = aws.String(_databasemigrationserviceDataProviderIdentifier)
	}
	if len(_databasemigrationserviceDataProviderName) > 0 {
		input.DataProviderName = aws.String(_databasemigrationserviceDataProviderName)
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}
	if len(_databasemigrationserviceEngine) > 0 {
		input.Engine = aws.String(_databasemigrationserviceEngine)
	}
	if len(_databasemigrationserviceExactSettings) > 0 {
		if err := assignInputField(input, "ExactSettings", _databasemigrationserviceExactSettings); err != nil {
			log.Errorf("invalid --exact-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSettings) > 0 {
		if err := assignInputField(input, "Settings", _databasemigrationserviceSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceVirtual) > 0 {
		if err := assignInputField(input, "Virtual", _databasemigrationserviceVirtual); err != nil {
			log.Errorf("invalid --virtual: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDataProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified endpoint.
// For a MySQL source or target endpoint, don't explicitly specify the database
// using the DatabaseName request parameter on the ModifyEndpoint API call.
// Specifying DatabaseName when you modify a MySQL endpoint replicates all the
// task tables to this single database. For MySQL endpoints, you specify the
// database only when you specify the schema in the table-mapping rules of the DMS
// task.
func databasemigrationservice_ModifyEndpoint(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}
	if len(_databasemigrationserviceCertificateArn) > 0 {
		input.CertificateArn = aws.String(_databasemigrationserviceCertificateArn)
	}
	if len(_databasemigrationserviceDatabaseName) > 0 {
		input.DatabaseName = aws.String(_databasemigrationserviceDatabaseName)
	}
	if len(_databasemigrationserviceDmsTransferSettings) > 0 {
		if err := assignInputField(input, "DmsTransferSettings", _databasemigrationserviceDmsTransferSettings); err != nil {
			log.Errorf("invalid --dms-transfer-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceDocDbSettings) > 0 {
		if err := assignInputField(input, "DocDbSettings", _databasemigrationserviceDocDbSettings); err != nil {
			log.Errorf("invalid --doc-db-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceDynamoDbSettings) > 0 {
		if err := assignInputField(input, "DynamoDbSettings", _databasemigrationserviceDynamoDbSettings); err != nil {
			log.Errorf("invalid --dynamo-db-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceElasticsearchSettings) > 0 {
		if err := assignInputField(input, "ElasticsearchSettings", _databasemigrationserviceElasticsearchSettings); err != nil {
			log.Errorf("invalid --elasticsearch-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEndpointIdentifier) > 0 {
		input.EndpointIdentifier = aws.String(_databasemigrationserviceEndpointIdentifier)
	}
	if len(_databasemigrationserviceEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _databasemigrationserviceEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEngineName) > 0 {
		input.EngineName = aws.String(_databasemigrationserviceEngineName)
	}
	if len(_databasemigrationserviceExactSettings) > 0 {
		if err := assignInputField(input, "ExactSettings", _databasemigrationserviceExactSettings); err != nil {
			log.Errorf("invalid --exact-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceExternalTableDefinition) > 0 {
		input.ExternalTableDefinition = aws.String(_databasemigrationserviceExternalTableDefinition)
	}
	if len(_databasemigrationserviceExtraConnectionAttributes) > 0 {
		input.ExtraConnectionAttributes = aws.String(_databasemigrationserviceExtraConnectionAttributes)
	}
	if len(_databasemigrationserviceGcpMySQLSettings) > 0 {
		if err := assignInputField(input, "GcpMySQLSettings", _databasemigrationserviceGcpMySQLSettings); err != nil {
			log.Errorf("invalid --gcp-my-sql-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceIBMDb2Settings) > 0 {
		if err := assignInputField(input, "IBMDb2Settings", _databasemigrationserviceIBMDb2Settings); err != nil {
			log.Errorf("invalid --ibmdb2-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKafkaSettings) > 0 {
		if err := assignInputField(input, "KafkaSettings", _databasemigrationserviceKafkaSettings); err != nil {
			log.Errorf("invalid --kafka-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceKinesisSettings) > 0 {
		if err := assignInputField(input, "KinesisSettings", _databasemigrationserviceKinesisSettings); err != nil {
			log.Errorf("invalid --kinesis-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMicrosoftSQLServerSettings) > 0 {
		if err := assignInputField(input, "MicrosoftSQLServerSettings", _databasemigrationserviceMicrosoftSQLServerSettings); err != nil {
			log.Errorf("invalid --microsoft-sql-server-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMongoDbSettings) > 0 {
		if err := assignInputField(input, "MongoDbSettings", _databasemigrationserviceMongoDbSettings); err != nil {
			log.Errorf("invalid --mongo-db-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMySQLSettings) > 0 {
		if err := assignInputField(input, "MySQLSettings", _databasemigrationserviceMySQLSettings); err != nil {
			log.Errorf("invalid --my-sql-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNeptuneSettings) > 0 {
		if err := assignInputField(input, "NeptuneSettings", _databasemigrationserviceNeptuneSettings); err != nil {
			log.Errorf("invalid --neptune-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceOracleSettings) > 0 {
		if err := assignInputField(input, "OracleSettings", _databasemigrationserviceOracleSettings); err != nil {
			log.Errorf("invalid --oracle-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationservicePassword) > 0 {
		input.Password = aws.String(_databasemigrationservicePassword)
	}
	if len(_databasemigrationservicePort) > 0 {
		if err := assignInputField(input, "Port", _databasemigrationservicePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationservicePostgreSQLSettings) > 0 {
		if err := assignInputField(input, "PostgreSQLSettings", _databasemigrationservicePostgreSQLSettings); err != nil {
			log.Errorf("invalid --postgre-sql-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceRedisSettings) > 0 {
		if err := assignInputField(input, "RedisSettings", _databasemigrationserviceRedisSettings); err != nil {
			log.Errorf("invalid --redis-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceRedshiftSettings) > 0 {
		if err := assignInputField(input, "RedshiftSettings", _databasemigrationserviceRedshiftSettings); err != nil {
			log.Errorf("invalid --redshift-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceS3Settings) > 0 {
		if err := assignInputField(input, "S3Settings", _databasemigrationserviceS3Settings); err != nil {
			log.Errorf("invalid --s3-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceServerName) > 0 {
		input.ServerName = aws.String(_databasemigrationserviceServerName)
	}
	if len(_databasemigrationserviceServiceAccessRoleArn) > 0 {
		input.ServiceAccessRoleArn = aws.String(_databasemigrationserviceServiceAccessRoleArn)
	}
	if len(_databasemigrationserviceSslMode) > 0 {
		if err := assignInputField(input, "SslMode", _databasemigrationserviceSslMode); err != nil {
			log.Errorf("invalid --ssl-mode: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSybaseSettings) > 0 {
		if err := assignInputField(input, "SybaseSettings", _databasemigrationserviceSybaseSettings); err != nil {
			log.Errorf("invalid --sybase-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTimestreamSettings) > 0 {
		if err := assignInputField(input, "TimestreamSettings", _databasemigrationserviceTimestreamSettings); err != nil {
			log.Errorf("invalid --timestream-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceUsername) > 0 {
		input.Username = aws.String(_databasemigrationserviceUsername)
	}

	if resp, err := client.ModifyEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing DMS event notification subscription.
func databasemigrationservice_ModifyEventSubscription(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_databasemigrationserviceSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_databasemigrationserviceSubscriptionName)
	}
	if len(_databasemigrationserviceEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _databasemigrationserviceEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _databasemigrationserviceEventCategories...)
	}
	if len(_databasemigrationserviceSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_databasemigrationserviceSnsTopicArn)
	}
	if len(_databasemigrationserviceSourceType) > 0 {
		input.SourceType = aws.String(_databasemigrationserviceSourceType)
	}

	if resp, err := client.ModifyEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified instance profile using the provided parameters.
// All migration projects associated with the instance profile must be deleted or
// modified before you can modify the instance profile.
func databasemigrationservice_ModifyInstanceProfile(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyInstanceProfileInput{
		// InstanceProfileIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceInstanceProfileIdentifier) > 0 {
		input.InstanceProfileIdentifier = aws.String(_databasemigrationserviceInstanceProfileIdentifier)
	}
	if len(_databasemigrationserviceAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_databasemigrationserviceAvailabilityZone)
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}
	if len(_databasemigrationserviceInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_databasemigrationserviceInstanceProfileName)
	}
	if len(_databasemigrationserviceKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_databasemigrationserviceKmsKeyArn)
	}
	if len(_databasemigrationserviceNetworkType) > 0 {
		input.NetworkType = aws.String(_databasemigrationserviceNetworkType)
	}
	if len(_databasemigrationservicePubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _databasemigrationservicePubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSubnetGroupIdentifier) > 0 {
		input.SubnetGroupIdentifier = aws.String(_databasemigrationserviceSubnetGroupIdentifier)
	}
	if len(_databasemigrationserviceVpcSecurityGroups) > 0 {
		input.VpcSecurityGroups = append([]string(nil), _databasemigrationserviceVpcSecurityGroups...)
	}

	if resp, err := client.ModifyInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified migration project using the provided parameters.
// The migration project must be closed before you can modify it.
func databasemigrationservice_ModifyMigrationProject(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyMigrationProjectInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceDescription) > 0 {
		input.Description = aws.String(_databasemigrationserviceDescription)
	}
	if len(_databasemigrationserviceInstanceProfileIdentifier) > 0 {
		input.InstanceProfileIdentifier = aws.String(_databasemigrationserviceInstanceProfileIdentifier)
	}
	if len(_databasemigrationserviceMigrationProjectName) > 0 {
		input.MigrationProjectName = aws.String(_databasemigrationserviceMigrationProjectName)
	}
	if len(_databasemigrationserviceSchemaConversionApplicationAttributes) > 0 {
		if err := assignInputField(input, "SchemaConversionApplicationAttributes", _databasemigrationserviceSchemaConversionApplicationAttributes); err != nil {
			log.Errorf("invalid --schema-conversion-application-attributes: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSourceDataProviderDescriptors) > 0 {
		if err := assignInputField(input, "SourceDataProviderDescriptors", _databasemigrationserviceSourceDataProviderDescriptors); err != nil {
			log.Errorf("invalid --source-data-provider-descriptors: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTargetDataProviderDescriptors) > 0 {
		if err := assignInputField(input, "TargetDataProviderDescriptors", _databasemigrationserviceTargetDataProviderDescriptors); err != nil {
			log.Errorf("invalid --target-data-provider-descriptors: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceTransformationRules) > 0 {
		input.TransformationRules = aws.String(_databasemigrationserviceTransformationRules)
	}

	if resp, err := client.ModifyMigrationProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing DMS Serverless replication configuration that you can use
// to start a replication. This command includes input validation and logic to
// check the state of any replication that uses this configuration. You can only
// modify a replication configuration before any replication that uses it has
// started. As soon as you have initially started a replication with a given
// configuiration, you can't modify that configuration, even if you stop it.
//
// Other run statuses that allow you to run this command include FAILED and
// CREATED. A provisioning state that allows you to run this command is
// FAILED_PROVISION.
func databasemigrationservice_ModifyReplicationConfig(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyReplicationConfigInput{
		// ReplicationConfigArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}
	if len(_databasemigrationserviceComputeConfig) > 0 {
		if err := assignInputField(input, "ComputeConfig", _databasemigrationserviceComputeConfig); err != nil {
			log.Errorf("invalid --compute-config: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationConfigIdentifier) > 0 {
		input.ReplicationConfigIdentifier = aws.String(_databasemigrationserviceReplicationConfigIdentifier)
	}
	if len(_databasemigrationserviceReplicationSettings) > 0 {
		input.ReplicationSettings = aws.String(_databasemigrationserviceReplicationSettings)
	}
	if len(_databasemigrationserviceReplicationType) > 0 {
		if err := assignInputField(input, "ReplicationType", _databasemigrationserviceReplicationType); err != nil {
			log.Errorf("invalid --replication-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSourceEndpointArn) > 0 {
		input.SourceEndpointArn = aws.String(_databasemigrationserviceSourceEndpointArn)
	}
	if len(_databasemigrationserviceSupplementalSettings) > 0 {
		input.SupplementalSettings = aws.String(_databasemigrationserviceSupplementalSettings)
	}
	if len(_databasemigrationserviceTableMappings) > 0 {
		input.TableMappings = aws.String(_databasemigrationserviceTableMappings)
	}
	if len(_databasemigrationserviceTargetEndpointArn) > 0 {
		input.TargetEndpointArn = aws.String(_databasemigrationserviceTargetEndpointArn)
	}

	if resp, err := client.ModifyReplicationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the replication instance to apply new settings. You can change one or
// more parameters by specifying these parameters and the new values in the
// request.
//
// Some settings are applied during the maintenance window.
func databasemigrationservice_ModifyReplicationInstance(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyReplicationInstanceInput{
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}
	if len(_databasemigrationserviceAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _databasemigrationserviceAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _databasemigrationserviceAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _databasemigrationserviceApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _databasemigrationserviceAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceEngineVersion) > 0 {
		input.EngineVersion = aws.String(_databasemigrationserviceEngineVersion)
	}
	if len(_databasemigrationserviceKerberosAuthenticationSettings) > 0 {
		if err := assignInputField(input, "KerberosAuthenticationSettings", _databasemigrationserviceKerberosAuthenticationSettings); err != nil {
			log.Errorf("invalid --kerberos-authentication-settings: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _databasemigrationserviceMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceNetworkType) > 0 {
		input.NetworkType = aws.String(_databasemigrationserviceNetworkType)
	}
	if len(_databasemigrationservicePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_databasemigrationservicePreferredMaintenanceWindow)
	}
	if len(_databasemigrationserviceReplicationInstanceClass) > 0 {
		input.ReplicationInstanceClass = aws.String(_databasemigrationserviceReplicationInstanceClass)
	}
	if len(_databasemigrationserviceReplicationInstanceIdentifier) > 0 {
		input.ReplicationInstanceIdentifier = aws.String(_databasemigrationserviceReplicationInstanceIdentifier)
	}
	if len(_databasemigrationserviceVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _databasemigrationserviceVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyReplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for the specified replication subnet group.
func databasemigrationservice_ModifyReplicationSubnetGroup(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyReplicationSubnetGroupInput{
		// ReplicationSubnetGroupIdentifier: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_databasemigrationserviceReplicationSubnetGroupIdentifier) > 0 {
		input.ReplicationSubnetGroupIdentifier = aws.String(_databasemigrationserviceReplicationSubnetGroupIdentifier)
	}
	if len(_databasemigrationserviceSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _databasemigrationserviceSubnetIds...)
	}
	if len(_databasemigrationserviceReplicationSubnetGroupDescription) > 0 {
		input.ReplicationSubnetGroupDescription = aws.String(_databasemigrationserviceReplicationSubnetGroupDescription)
	}

	if resp, err := client.ModifyReplicationSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified replication task.
// You can't modify the task endpoints. The task must be stopped before you can
// modify it.
//
// For more information about DMS tasks, see [Working with Migration Tasks] in the Database Migration Service
// User Guide.
//
// [Working with Migration Tasks]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html
func databasemigrationservice_ModifyReplicationTask(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ModifyReplicationTaskInput{
		// ReplicationTaskArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceCdcStartPosition) > 0 {
		input.CdcStartPosition = aws.String(_databasemigrationserviceCdcStartPosition)
	}
	if len(_databasemigrationserviceCdcStartTime) > 0 {
		if err := assignInputField(input, "CdcStartTime", _databasemigrationserviceCdcStartTime); err != nil {
			log.Errorf("invalid --cdc-start-time: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceCdcStopPosition) > 0 {
		input.CdcStopPosition = aws.String(_databasemigrationserviceCdcStopPosition)
	}
	if len(_databasemigrationserviceMigrationType) > 0 {
		if err := assignInputField(input, "MigrationType", _databasemigrationserviceMigrationType); err != nil {
			log.Errorf("invalid --migration-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReplicationTaskIdentifier) > 0 {
		input.ReplicationTaskIdentifier = aws.String(_databasemigrationserviceReplicationTaskIdentifier)
	}
	if len(_databasemigrationserviceReplicationTaskSettings) > 0 {
		input.ReplicationTaskSettings = aws.String(_databasemigrationserviceReplicationTaskSettings)
	}
	if len(_databasemigrationserviceTableMappings) > 0 {
		input.TableMappings = aws.String(_databasemigrationserviceTableMappings)
	}
	if len(_databasemigrationserviceTaskData) > 0 {
		input.TaskData = aws.String(_databasemigrationserviceTaskData)
	}

	if resp, err := client.ModifyReplicationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves a replication task from its current replication instance to a different
// target replication instance using the specified parameters. The target
// replication instance must be created with the same or later DMS version as the
// current replication instance.
func databasemigrationservice_MoveReplicationTask(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.MoveReplicationTaskInput{
		// ReplicationTaskArn: *string, // Required
		// TargetReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceTargetReplicationInstanceArn) > 0 {
		input.TargetReplicationInstanceArn = aws.String(_databasemigrationserviceTargetReplicationInstanceArn)
	}

	if resp, err := client.MoveReplicationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots a replication instance. Rebooting results in a momentary outage, until
// the replication instance becomes available again.
func databasemigrationservice_RebootReplicationInstance(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.RebootReplicationInstanceInput{
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}
	if len(_databasemigrationserviceForceFailover) > 0 {
		if err := assignInputField(input, "ForceFailover", _databasemigrationserviceForceFailover); err != nil {
			log.Errorf("invalid --force-failover: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceForcePlannedFailover) > 0 {
		if err := assignInputField(input, "ForcePlannedFailover", _databasemigrationserviceForcePlannedFailover); err != nil {
			log.Errorf("invalid --force-planned-failover: %s", err.Error())
			return
		}
	}

	if resp, err := client.RebootReplicationInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Populates the schema for the specified endpoint. This is an asynchronous
// operation and can take several minutes. You can check the status of this
// operation by calling the DescribeRefreshSchemasStatus operation.
func databasemigrationservice_RefreshSchemas(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.RefreshSchemasInput{
		// EndpointArn: *string, // Required
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}

	if resp, err := client.RefreshSchemas(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reloads the target database table with the source data for a given DMS
// Serverless replication configuration.
//
// You can only use this operation with a task in the RUNNING state, otherwise the
// service will throw an InvalidResourceStateFault exception.
func databasemigrationservice_ReloadReplicationTables(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ReloadReplicationTablesInput{
		// ReplicationConfigArn: *string, // Required
		// TablesToReload: []types.TableToReload, // Required
	}

	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}
	if len(_databasemigrationserviceTablesToReload) > 0 {
		if err := assignInputField(input, "TablesToReload", _databasemigrationserviceTablesToReload); err != nil {
			log.Errorf("invalid --tables-to-reload: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReloadOption) > 0 {
		if err := assignInputField(input, "ReloadOption", _databasemigrationserviceReloadOption); err != nil {
			log.Errorf("invalid --reload-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReloadReplicationTables(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reloads the target database table with the source data.
// You can only use this operation with a task in the RUNNING state, otherwise the
// service will throw an InvalidResourceStateFault exception.
func databasemigrationservice_ReloadTables(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.ReloadTablesInput{
		// ReplicationTaskArn: *string, // Required
		// TablesToReload: []types.TableToReload, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceTablesToReload) > 0 {
		if err := assignInputField(input, "TablesToReload", _databasemigrationserviceTablesToReload); err != nil {
			log.Errorf("invalid --tables-to-reload: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceReloadOption) > 0 {
		if err := assignInputField(input, "ReloadOption", _databasemigrationserviceReloadOption); err != nil {
			log.Errorf("invalid --reload-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReloadTables(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes metadata tags from an DMS resource, including replication instance,
// endpoint, subnet group, and migration task. For more information, see [Tag]Tag data
// type description.
//
// [Tag]: https://docs.aws.amazon.com/dms/latest/APIReference/API_Tag.html
func databasemigrationservice_RemoveTagsFromResource(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.RemoveTagsFromResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_databasemigrationserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_databasemigrationserviceResourceArn)
	}
	if len(_databasemigrationserviceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _databasemigrationserviceTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Runs large-scale assessment (LSA) analysis on every Fleet Advisor collector in
// your account.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func databasemigrationservice_RunFleetAdvisorLsaAnalysis(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.RunFleetAdvisorLsaAnalysisInput{}

	if resp, err := client.RunFleetAdvisorLsaAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified data migration.
func databasemigrationservice_StartDataMigration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartDataMigrationInput{
		// DataMigrationIdentifier: *string, // Required
		// StartType: types.StartReplicationMigrationTypeValue, // Required
	}

	if len(_databasemigrationserviceDataMigrationIdentifier) > 0 {
		input.DataMigrationIdentifier = aws.String(_databasemigrationserviceDataMigrationIdentifier)
	}
	if len(_databasemigrationserviceStartType) > 0 {
		if err := assignInputField(input, "StartType", _databasemigrationserviceStartType); err != nil {
			log.Errorf("invalid --start-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDataMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies the extension pack to your target database. An extension pack is an
// add-on module that emulates functions present in a source database that are
// required when converting objects to the target database.
func databasemigrationservice_StartExtensionPackAssociation(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartExtensionPackAssociationInput{
		// MigrationProjectIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}

	if resp, err := client.StartExtensionPackAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a database migration assessment report by assessing the migration
// complexity for your source database. A database migration assessment report
// summarizes all of the schema conversion tasks. It also details the action items
// for database objects that can't be converted to the database engine of your
// target database instance.
func databasemigrationservice_StartMetadataModelAssessment(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartMetadataModelAssessmentInput{
		// MigrationProjectIdentifier: *string, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}

	if resp, err := client.StartMetadataModelAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Converts your source database objects to a format compatible with the target
// database.
func databasemigrationservice_StartMetadataModelConversion(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartMetadataModelConversionInput{
		// MigrationProjectIdentifier: *string, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}

	if resp, err := client.StartMetadataModelConversion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates source metadata model of the given type with the specified properties
// for schema conversion operations.
//
// This action supports only these directions: from SQL Server to Aurora
// PostgreSQL, or from SQL Server to RDS for PostgreSQL.
func databasemigrationservice_StartMetadataModelCreation(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartMetadataModelCreationInput{
		// MetadataModelName: *string, // Required
		// MigrationProjectIdentifier: *string, // Required
		// Properties: types.MetadataModelProperties, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMetadataModelName) > 0 {
		input.MetadataModelName = aws.String(_databasemigrationserviceMetadataModelName)
	}
	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceProperties) > 0 {
		if err := assignInputField(input, "Properties", _databasemigrationserviceProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}

	if resp, err := client.StartMetadataModelCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Saves your converted code to a file as a SQL script, and stores this file on
// your Amazon S3 bucket.
func databasemigrationservice_StartMetadataModelExportAsScript(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartMetadataModelExportAsScriptInput{
		// MigrationProjectIdentifier: *string, // Required
		// Origin: types.OriginTypeValue, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceOrigin) > 0 {
		if err := assignInputField(input, "Origin", _databasemigrationserviceOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceFileName) > 0 {
		input.FileName = aws.String(_databasemigrationserviceFileName)
	}

	if resp, err := client.StartMetadataModelExportAsScript(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies converted database objects to your target database.
func databasemigrationservice_StartMetadataModelExportToTarget(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartMetadataModelExportToTargetInput{
		// MigrationProjectIdentifier: *string, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceOverwriteExtensionPack) > 0 {
		if err := assignInputField(input, "OverwriteExtensionPack", _databasemigrationserviceOverwriteExtensionPack); err != nil {
			log.Errorf("invalid --overwrite-extension-pack: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMetadataModelExportToTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Loads the metadata for all the dependent database objects of the parent object.
// This operation uses your project's Amazon S3 bucket as a metadata cache to
// improve performance.
func databasemigrationservice_StartMetadataModelImport(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartMetadataModelImportInput{
		// MigrationProjectIdentifier: *string, // Required
		// Origin: types.OriginTypeValue, // Required
		// SelectionRules: *string, // Required
	}

	if len(_databasemigrationserviceMigrationProjectIdentifier) > 0 {
		input.MigrationProjectIdentifier = aws.String(_databasemigrationserviceMigrationProjectIdentifier)
	}
	if len(_databasemigrationserviceOrigin) > 0 {
		if err := assignInputField(input, "Origin", _databasemigrationserviceOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceSelectionRules) > 0 {
		input.SelectionRules = aws.String(_databasemigrationserviceSelectionRules)
	}
	if len(_databasemigrationserviceRefresh) > 0 {
		if err := assignInputField(input, "Refresh", _databasemigrationserviceRefresh); err != nil {
			log.Errorf("invalid --refresh: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMetadataModelImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On May 20, 2026, Amazon Web Services will end support
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Starts the analysis of your source database to provide recommendations of
// target engines.
//
// You can create recommendations for multiple source databases using [BatchStartRecommendations].
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
// [BatchStartRecommendations]: https://docs.aws.amazon.com/dms/latest/APIReference/API_BatchStartRecommendations.html
func databasemigrationservice_StartRecommendations(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartRecommendationsInput{
		// DatabaseId: *string, // Required
		// Settings: *types.RecommendationSettings, // Required
	}

	if len(_databasemigrationserviceDatabaseId) > 0 {
		input.DatabaseId = aws.String(_databasemigrationserviceDatabaseId)
	}
	if len(_databasemigrationserviceSettings) > 0 {
		if err := assignInputField(input, "Settings", _databasemigrationserviceSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a given DMS Serverless replication configuration, DMS connects to the
// source endpoint and collects the metadata to analyze the replication workload.
// Using this metadata, DMS then computes and provisions the required capacity and
// starts replicating to the target endpoint using the server resources that DMS
// has provisioned for the DMS Serverless replication.
func databasemigrationservice_StartReplication(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartReplicationInput{
		// ReplicationConfigArn: *string, // Required
		// StartReplicationType: *string, // Required
	}

	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}
	if len(_databasemigrationserviceStartReplicationType) > 0 {
		input.StartReplicationType = aws.String(_databasemigrationserviceStartReplicationType)
	}
	if len(_databasemigrationserviceCdcStartPosition) > 0 {
		input.CdcStartPosition = aws.String(_databasemigrationserviceCdcStartPosition)
	}
	if len(_databasemigrationserviceCdcStartTime) > 0 {
		if err := assignInputField(input, "CdcStartTime", _databasemigrationserviceCdcStartTime); err != nil {
			log.Errorf("invalid --cdc-start-time: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceCdcStopPosition) > 0 {
		input.CdcStopPosition = aws.String(_databasemigrationserviceCdcStopPosition)
	}
	if len(_databasemigrationservicePremigrationAssessmentSettings) > 0 {
		input.PremigrationAssessmentSettings = aws.String(_databasemigrationservicePremigrationAssessmentSettings)
	}

	if resp, err := client.StartReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the replication task.
// For more information about DMS tasks, see [Working with Migration Tasks] in the Database Migration Service
// User Guide.
//
// [Working with Migration Tasks]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html
func databasemigrationservice_StartReplicationTask(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartReplicationTaskInput{
		// ReplicationTaskArn: *string, // Required
		// StartReplicationTaskType: types.StartReplicationTaskTypeValue, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceStartReplicationTaskType) > 0 {
		if err := assignInputField(input, "StartReplicationTaskType", _databasemigrationserviceStartReplicationTaskType); err != nil {
			log.Errorf("invalid --start-replication-task-type: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceCdcStartPosition) > 0 {
		input.CdcStartPosition = aws.String(_databasemigrationserviceCdcStartPosition)
	}
	if len(_databasemigrationserviceCdcStartTime) > 0 {
		if err := assignInputField(input, "CdcStartTime", _databasemigrationserviceCdcStartTime); err != nil {
			log.Errorf("invalid --cdc-start-time: %s", err.Error())
			return
		}
	}
	if len(_databasemigrationserviceCdcStopPosition) > 0 {
		input.CdcStopPosition = aws.String(_databasemigrationserviceCdcStopPosition)
	}

	if resp, err := client.StartReplicationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the replication task assessment for unsupported data types in the
// source database.
//
// You can only use this operation for a task if the following conditions are true:
//
// - The task must be in the stopped state.
//
// - The task must have successful connections to the source and target.
//
// If either of these conditions are not met, an InvalidResourceStateFault error
// will result.
//
// For information about DMS task assessments, see [Creating a task assessment report] in the Database Migration
// Service User Guide.
//
// [Creating a task assessment report]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.AssessmentReport.html
func databasemigrationservice_StartReplicationTaskAssessment(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartReplicationTaskAssessmentInput{
		// ReplicationTaskArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}

	if resp, err := client.StartReplicationTaskAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new premigration assessment run for one or more individual assessments
// of a migration task.
//
// The assessments that you can specify depend on the source and target database
// engine and the migration type defined for the given task. To run this operation,
// your migration task must already be created. After you run this operation, you
// can review the status of each individual assessment. You can also run the
// migration task manually after the assessment run and its individual assessments
// complete.
func databasemigrationservice_StartReplicationTaskAssessmentRun(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StartReplicationTaskAssessmentRunInput{
		// AssessmentRunName: *string, // Required
		// ReplicationTaskArn: *string, // Required
		// ResultLocationBucket: *string, // Required
		// ServiceAccessRoleArn: *string, // Required
	}

	if len(_databasemigrationserviceAssessmentRunName) > 0 {
		input.AssessmentRunName = aws.String(_databasemigrationserviceAssessmentRunName)
	}
	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}
	if len(_databasemigrationserviceResultLocationBucket) > 0 {
		input.ResultLocationBucket = aws.String(_databasemigrationserviceResultLocationBucket)
	}
	if len(_databasemigrationserviceServiceAccessRoleArn) > 0 {
		input.ServiceAccessRoleArn = aws.String(_databasemigrationserviceServiceAccessRoleArn)
	}
	if len(_databasemigrationserviceExclude) > 0 {
		input.Exclude = append([]string(nil), _databasemigrationserviceExclude...)
	}
	if len(_databasemigrationserviceIncludeOnly) > 0 {
		input.IncludeOnly = append([]string(nil), _databasemigrationserviceIncludeOnly...)
	}
	if len(_databasemigrationserviceResultEncryptionMode) > 0 {
		input.ResultEncryptionMode = aws.String(_databasemigrationserviceResultEncryptionMode)
	}
	if len(_databasemigrationserviceResultKmsKeyArn) > 0 {
		input.ResultKmsKeyArn = aws.String(_databasemigrationserviceResultKmsKeyArn)
	}
	if len(_databasemigrationserviceResultLocationFolder) > 0 {
		input.ResultLocationFolder = aws.String(_databasemigrationserviceResultLocationFolder)
	}
	if len(_databasemigrationserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _databasemigrationserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartReplicationTaskAssessmentRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified data migration.
func databasemigrationservice_StopDataMigration(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StopDataMigrationInput{
		// DataMigrationIdentifier: *string, // Required
	}

	if len(_databasemigrationserviceDataMigrationIdentifier) > 0 {
		input.DataMigrationIdentifier = aws.String(_databasemigrationserviceDataMigrationIdentifier)
	}

	if resp, err := client.StopDataMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a given DMS Serverless replication configuration, DMS stops any and all
// ongoing DMS Serverless replications. This command doesn't deprovision the
// stopped replications.
func databasemigrationservice_StopReplication(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StopReplicationInput{
		// ReplicationConfigArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationConfigArn) > 0 {
		input.ReplicationConfigArn = aws.String(_databasemigrationserviceReplicationConfigArn)
	}

	if resp, err := client.StopReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the replication task.
func databasemigrationservice_StopReplicationTask(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.StopReplicationTaskInput{
		// ReplicationTaskArn: *string, // Required
	}

	if len(_databasemigrationserviceReplicationTaskArn) > 0 {
		input.ReplicationTaskArn = aws.String(_databasemigrationserviceReplicationTaskArn)
	}

	if resp, err := client.StopReplicationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests the connection between the replication instance and the endpoint.
func databasemigrationservice_TestConnection(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.TestConnectionInput{
		// EndpointArn: *string, // Required
		// ReplicationInstanceArn: *string, // Required
	}

	if len(_databasemigrationserviceEndpointArn) > 0 {
		input.EndpointArn = aws.String(_databasemigrationserviceEndpointArn)
	}
	if len(_databasemigrationserviceReplicationInstanceArn) > 0 {
		input.ReplicationInstanceArn = aws.String(_databasemigrationserviceReplicationInstanceArn)
	}

	if resp, err := client.TestConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Migrates 10 active and enabled Amazon SNS subscriptions at a time and converts
// them to corresponding Amazon EventBridge rules. By default, this operation
// migrates subscriptions only when all your replication instance versions are
// 3.4.5 or higher. If any replication instances are from versions earlier than
// 3.4.5, the operation raises an error and tells you to upgrade these instances to
// version 3.4.5 or higher. To enable migration regardless of version, set the
// Force option to true. However, if you don't upgrade instances earlier than
// version 3.4.5, some types of events might not be available when you use Amazon
// EventBridge.
//
// To call this operation, make sure that you have certain permissions added to
// your user account. For more information, see [Migrating event subscriptions to Amazon EventBridge]in the Amazon Web Services
// Database Migration Service User Guide.
//
// [Migrating event subscriptions to Amazon EventBridge]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html#CHAP_Events-migrate-to-eventbridge
func databasemigrationservice_UpdateSubscriptionsToEventBridge(cfg aws.Config, client *databasemigrationservice.Client) {
	input := &databasemigrationservice.UpdateSubscriptionsToEventBridgeInput{}

	if len(_databasemigrationserviceForceMove) > 0 {
		if err := assignInputField(input, "ForceMove", _databasemigrationserviceForceMove); err != nil {
			log.Errorf("invalid --force-move: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSubscriptionsToEventBridge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_databasemigrationserviceCmd)
	_databasemigrationserviceCmd.Flags().SortFlags = false

	_databasemigrationserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_databasemigrationserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_databasemigrationserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceAllocatedStorage, "allocated-storage", "", "", "Allocated Storage")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceAllowMajorVersionUpgrade, "allow-major-version-upgrade", "", "", "Allow Major Version Upgrade")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceApplyAction, "apply-action", "", "", "Apply Action")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceApplyImmediately, "apply-immediately", "", "", "Apply Immediately")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceAssessmentReportTypes, "assessment-report-types", "", "", "Assessment Report Types")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceAssessmentRunName, "assessment-run-name", "", "", "Assessment Run Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCdcStartPosition, "cdc-start-position", "", "", "Cdc Start Position")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCdcStartTime, "cdc-start-time", "", "", "Cdc Start Time")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCdcStopPosition, "cdc-stop-position", "", "", "Cdc Stop Position")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCertificateIdentifier, "certificate-identifier", "", "", "Certificate Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCertificatePem, "certificate-pem", "", "", "Certificate Pem")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCertificateWallet, "certificate-wallet", "", "", "Certificate Wallet")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCollectorName, "collector-name", "", "", "Collector Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceCollectorReferencedId, "collector-referenced-id", "", "", "Collector Referenced ID")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceComputeConfig, "compute-config", "", "", "Compute Config")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceConversionConfiguration, "conversion-configuration", "", "", "Conversion Configuration")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceData, "data", "", "", "Data")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDataMigrationIdentifier, "data-migration-identifier", "", "", "Data Migration Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDataMigrationName, "data-migration-name", "", "", "Data Migration Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDataMigrationType, "data-migration-type", "", "", "Data Migration Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDataProviderIdentifier, "data-provider-identifier", "", "", "Data Provider Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDataProviderName, "data-provider-name", "", "", "Data Provider Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDatabaseId, "database-id", "", "", "Database ID")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceDatabaseIds, "database-ids", "", nil, "Database Ids")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDatabaseName, "database-name", "", "", "Database Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDescription, "description", "", "", "Description")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDmsTransferSettings, "dms-transfer-settings", "", "", "Dms Transfer Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDnsNameServers, "dns-name-servers", "", "", "DNS Name Servers")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDocDbSettings, "doc-db-settings", "", "", "Doc DB Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDuration, "duration", "", "", "Duration")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceDynamoDbSettings, "dynamo-db-settings", "", "", "Dynamo DB Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceElasticsearchSettings, "elasticsearch-settings", "", "", "Elasticsearch Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEnableCloudwatchLogs, "enable-cloudwatch-logs", "", "", "Enable Cloudwatch Logs")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEnabled, "enabled", "", "", "Enabled")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEndTime, "end-time", "", "", "End Time")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEndpointArn, "endpoint-arn", "", "", "Endpoint ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEndpointIdentifier, "endpoint-identifier", "", "", "Endpoint Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEngine, "engine", "", "", "Engine")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEngineName, "engine-name", "", "", "Engine Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceEngineVersion, "engine-version", "", "", "Engine Version")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceEventCategories, "event-categories", "", nil, "Event Categories")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceExactSettings, "exact-settings", "", "", "Exact Settings")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceExclude, "exclude", "", nil, "Exclude")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceExternalTableDefinition, "external-table-definition", "", "", "External Table Definition")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceExtraConnectionAttributes, "extra-connection-attributes", "", "", "Extra Connection Attributes")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceFileName, "file-name", "", "", "File Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceFilters, "filters", "", "", "Filters")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceForceFailover, "force-failover", "", "", "Force Failover")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceForceMove, "force-move", "", "", "Force Move")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceForcePlannedFailover, "force-planned-failover", "", "", "Force Planned Failover")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceGcpMySQLSettings, "gcp-my-sql-settings", "", "", "Gcp My Sql Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceIBMDb2Settings, "ibmdb2-settings", "", "", "Ibmdb2 Settings")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceIncludeOnly, "include-only", "", nil, "Include Only")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceInstanceProfileIdentifier, "instance-profile-identifier", "", "", "Instance Profile Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceInstanceProfileName, "instance-profile-name", "", "", "Instance Profile Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceKafkaSettings, "kafka-settings", "", "", "Kafka Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceKerberosAuthenticationSettings, "kerberos-authentication-settings", "", "", "Kerberos Authentication Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceKinesisSettings, "kinesis-settings", "", "", "Kinesis Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMarker, "marker", "", "", "Marker")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMaxRecords, "max-records", "", "", "Max Records")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMetadataModelName, "metadata-model-name", "", "", "Metadata Model Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMicrosoftSQLServerSettings, "microsoft-sql-server-settings", "", "", "Microsoft Sql Server Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMigrationProjectIdentifier, "migration-project-identifier", "", "", "Migration Project Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMigrationProjectName, "migration-project-name", "", "", "Migration Project Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMigrationType, "migration-type", "", "", "Migration Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMongoDbSettings, "mongo-db-settings", "", "", "Mongo DB Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMultiAZ, "multi-az", "", "", "Multi AZ")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceMySQLSettings, "my-sql-settings", "", "", "My Sql Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceNeptuneSettings, "neptune-settings", "", "", "Neptune Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceNetworkType, "network-type", "", "", "Network Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceNextToken, "next-token", "", "", "Next Token")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceNumberOfJobs, "number-of-jobs", "", "", "Number Of Jobs")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceOptInType, "opt-in-type", "", "", "Opt In Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceOracleSettings, "oracle-settings", "", "", "Oracle Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceOrigin, "origin", "", "", "Origin")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceOverwriteExtensionPack, "overwrite-extension-pack", "", "", "Overwrite Extension Pack")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationservicePassword, "password", "", "", "Password")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationservicePort, "port", "", "", "Port")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationservicePostgreSQLSettings, "postgre-sql-settings", "", "", "Postgre Sql Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationservicePreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationservicePremigrationAssessmentSettings, "premigration-assessment-settings", "", "", "Premigration Assessment Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceProperties, "properties", "", "", "Properties")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationservicePubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceRedisSettings, "redis-settings", "", "", "Redis Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceRedshiftSettings, "redshift-settings", "", "", "Redshift Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceRefresh, "refresh", "", "", "Refresh")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReloadOption, "reload-option", "", "", "Reload Option")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationConfigArn, "replication-config-arn", "", "", "Replication Config ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationConfigIdentifier, "replication-config-identifier", "", "", "Replication Config Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationInstanceArn, "replication-instance-arn", "", "", "Replication Instance ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationInstanceClass, "replication-instance-class", "", "", "Replication Instance Class")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationInstanceIdentifier, "replication-instance-identifier", "", "", "Replication Instance Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationSettings, "replication-settings", "", "", "Replication Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationSubnetGroupDescription, "replication-subnet-group-description", "", "", "Replication Subnet Group Description")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationSubnetGroupIdentifier, "replication-subnet-group-identifier", "", "", "Replication Subnet Group Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationTaskArn, "replication-task-arn", "", "", "Replication Task ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationTaskAssessmentRunArn, "replication-task-assessment-run-arn", "", "", "Replication Task Assessment Run ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationTaskIdentifier, "replication-task-identifier", "", "", "Replication Task Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationTaskSettings, "replication-task-settings", "", "", "Replication Task Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceReplicationType, "replication-type", "", "", "Replication Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceRequestIdentifier, "request-identifier", "", "", "Request Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceResourceArn, "resource-arn", "", "", "Resource ARN")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceResourceArnList, "resource-arn-list", "", nil, "Resource ARN List")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceResultEncryptionMode, "result-encryption-mode", "", "", "Result Encryption Mode")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceResultKmsKeyArn, "result-kms-key-arn", "", "", "Result KMS Key ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceResultLocationBucket, "result-location-bucket", "", "", "Result Location Bucket")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceResultLocationFolder, "result-location-folder", "", "", "Result Location Folder")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceS3Settings, "s3-settings", "", "", "S3 Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSchemaConversionApplicationAttributes, "schema-conversion-application-attributes", "", "", "Schema Conversion Application Attributes")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSelectionRules, "selection-rules", "", "", "Selection Rules")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceServerName, "server-name", "", "", "Server Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceServiceAccessRoleArn, "service-access-role-arn", "", "", "Service Access Role ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSettings, "settings", "", "", "Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSourceDataProviderDescriptors, "source-data-provider-descriptors", "", "", "Source Data Provider Descriptors")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSourceDataSettings, "source-data-settings", "", "", "Source Data Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSourceEndpointArn, "source-endpoint-arn", "", "", "Source Endpoint ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSourceEngineName, "source-engine-name", "", "", "Source Engine Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSourceIdentifier, "source-identifier", "", "", "Source Identifier")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceSourceIds, "source-ids", "", nil, "Source Ids")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSourceType, "source-type", "", "", "Source Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSslMode, "ssl-mode", "", "", "SSL Mode")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceStartReplicationTaskType, "start-replication-task-type", "", "", "Start Replication Task Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceStartReplicationType, "start-replication-type", "", "", "Start Replication Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceStartTime, "start-time", "", "", "Start Time")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceStartType, "start-type", "", "", "Start Type")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSubnetGroupIdentifier, "subnet-group-identifier", "", "", "Subnet Group Identifier")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSubscriptionName, "subscription-name", "", "", "Subscription Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSupplementalSettings, "supplemental-settings", "", "", "Supplemental Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceSybaseSettings, "sybase-settings", "", "", "Sybase Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTableMappings, "table-mappings", "", "", "Table Mappings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTablesToReload, "tables-to-reload", "", "", "Tables To Reload")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTags, "tags", "", "", "Tags")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTargetDataProviderDescriptors, "target-data-provider-descriptors", "", "", "Target Data Provider Descriptors")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTargetDataSettings, "target-data-settings", "", "", "Target Data Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTargetEndpointArn, "target-endpoint-arn", "", "", "Target Endpoint ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTargetEngineName, "target-engine-name", "", "", "Target Engine Name")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTargetReplicationInstanceArn, "target-replication-instance-arn", "", "", "Target Replication Instance ARN")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTaskData, "task-data", "", "", "Task Data")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTimestreamSettings, "timestream-settings", "", "", "Timestream Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceTransformationRules, "transformation-rules", "", "", "Transformation Rules")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceUsername, "username", "", "", "Username")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceVirtual, "virtual", "", "", "Virtual")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")
	_databasemigrationserviceCmd.Flags().StringSliceVarP(&_databasemigrationserviceVpcSecurityGroups, "vpc-security-groups", "", nil, "VPC Security Groups")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceWithoutSettings, "without-settings", "", "", "Without Settings")
	_databasemigrationserviceCmd.Flags().StringVarP(&_databasemigrationserviceWithoutStatistics, "without-statistics", "", "", "Without Statistics")

	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceApplyPendingMaintenanceAction, "apply-pending-maintenance-action", "", false, "Apply Pending Maintenance Action")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceBatchStartRecommendations, "batch-start-recommendations", "", false, "Batch Start Recommendations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCancelMetadataModelConversion, "cancel-metadata-model-conversion", "", false, "Cancel Metadata Model Conversion")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCancelMetadataModelCreation, "cancel-metadata-model-creation", "", false, "Cancel Metadata Model Creation")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCancelReplicationTaskAssessmentRun, "cancel-replication-task-assessment-run", "", false, "Cancel Replication Task Assessment Run")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateDataMigration, "create-data-migration", "", false, "Create Data Migration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateDataProvider, "create-data-provider", "", false, "Create Data Provider")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateEndpoint, "create-endpoint", "", false, "Create Endpoint")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateEventSubscription, "create-event-subscription", "", false, "Create Event Subscription")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateFleetAdvisorCollector, "create-fleet-advisor-collector", "", false, "Create Fleet Advisor Collector")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateInstanceProfile, "create-instance-profile", "", false, "Create Instance Profile")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateMigrationProject, "create-migration-project", "", false, "Create Migration Project")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateReplicationConfig, "create-replication-config", "", false, "Create Replication Config")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateReplicationInstance, "create-replication-instance", "", false, "Create Replication Instance")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateReplicationSubnetGroup, "create-replication-subnet-group", "", false, "Create Replication Subnet Group")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceCreateReplicationTask, "create-replication-task", "", false, "Create Replication Task")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteCertificate, "delete-certificate", "", false, "Delete Certificate")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteDataMigration, "delete-data-migration", "", false, "Delete Data Migration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteDataProvider, "delete-data-provider", "", false, "Delete Data Provider")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteEventSubscription, "delete-event-subscription", "", false, "Delete Event Subscription")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteFleetAdvisorCollector, "delete-fleet-advisor-collector", "", false, "Delete Fleet Advisor Collector")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteFleetAdvisorDatabases, "delete-fleet-advisor-databases", "", false, "Delete Fleet Advisor Databases")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteInstanceProfile, "delete-instance-profile", "", false, "Delete Instance Profile")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteMigrationProject, "delete-migration-project", "", false, "Delete Migration Project")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteReplicationConfig, "delete-replication-config", "", false, "Delete Replication Config")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteReplicationInstance, "delete-replication-instance", "", false, "Delete Replication Instance")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteReplicationSubnetGroup, "delete-replication-subnet-group", "", false, "Delete Replication Subnet Group")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteReplicationTask, "delete-replication-task", "", false, "Delete Replication Task")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDeleteReplicationTaskAssessmentRun, "delete-replication-task-assessment-run", "", false, "Delete Replication Task Assessment Run")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeAccountAttributes, "describe-account-attributes", "", false, "Describe Account Attributes")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeApplicableIndividualAssessments, "describe-applicable-individual-assessments", "", false, "Describe Applicable Individual Assessments")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeCertificates, "describe-certificates", "", false, "Describe Certificates")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeConnections, "describe-connections", "", false, "Describe Connections")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeConversionConfiguration, "describe-conversion-configuration", "", false, "Describe Conversion Configuration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeDataMigrations, "describe-data-migrations", "", false, "Describe Data Migrations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeDataProviders, "describe-data-providers", "", false, "Describe Data Providers")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEndpointSettings, "describe-endpoint-settings", "", false, "Describe Endpoint Settings")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEndpointTypes, "describe-endpoint-types", "", false, "Describe Endpoint Types")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEndpoints, "describe-endpoints", "", false, "Describe Endpoints")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEngineVersions, "describe-engine-versions", "", false, "Describe Engine Versions")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEventCategories, "describe-event-categories", "", false, "Describe Event Categories")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEventSubscriptions, "describe-event-subscriptions", "", false, "Describe Event Subscriptions")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeEvents, "describe-events", "", false, "Describe Events")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeExtensionPackAssociations, "describe-extension-pack-associations", "", false, "Describe Extension Pack Associations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeFleetAdvisorCollectors, "describe-fleet-advisor-collectors", "", false, "Describe Fleet Advisor Collectors")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeFleetAdvisorDatabases, "describe-fleet-advisor-databases", "", false, "Describe Fleet Advisor Databases")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeFleetAdvisorLsaAnalysis, "describe-fleet-advisor-lsa-analysis", "", false, "Describe Fleet Advisor Lsa Analysis")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeFleetAdvisorSchemaObjectSummary, "describe-fleet-advisor-schema-object-summary", "", false, "Describe Fleet Advisor Schema Object Summary")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeFleetAdvisorSchemas, "describe-fleet-advisor-schemas", "", false, "Describe Fleet Advisor Schemas")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeInstanceProfiles, "describe-instance-profiles", "", false, "Describe Instance Profiles")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModel, "describe-metadata-model", "", false, "Describe Metadata Model")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelAssessments, "describe-metadata-model-assessments", "", false, "Describe Metadata Model Assessments")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelChildren, "describe-metadata-model-children", "", false, "Describe Metadata Model Children")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelConversions, "describe-metadata-model-conversions", "", false, "Describe Metadata Model Conversions")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelCreations, "describe-metadata-model-creations", "", false, "Describe Metadata Model Creations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelExportsAsScript, "describe-metadata-model-exports-as-script", "", false, "Describe Metadata Model Exports As Script")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelExportsToTarget, "describe-metadata-model-exports-to-target", "", false, "Describe Metadata Model Exports To Target")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMetadataModelImports, "describe-metadata-model-imports", "", false, "Describe Metadata Model Imports")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeMigrationProjects, "describe-migration-projects", "", false, "Describe Migration Projects")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeOrderableReplicationInstances, "describe-orderable-replication-instances", "", false, "Describe Orderable Replication Instances")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribePendingMaintenanceActions, "describe-pending-maintenance-actions", "", false, "Describe Pending Maintenance Actions")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeRecommendationLimitations, "describe-recommendation-limitations", "", false, "Describe Recommendation Limitations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeRecommendations, "describe-recommendations", "", false, "Describe Recommendations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeRefreshSchemasStatus, "describe-refresh-schemas-status", "", false, "Describe Refresh Schemas Status")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationConfigs, "describe-replication-configs", "", false, "Describe Replication Configs")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationInstanceTaskLogs, "describe-replication-instance-task-logs", "", false, "Describe Replication Instance Task Logs")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationInstances, "describe-replication-instances", "", false, "Describe Replication Instances")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationSubnetGroups, "describe-replication-subnet-groups", "", false, "Describe Replication Subnet Groups")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationTableStatistics, "describe-replication-table-statistics", "", false, "Describe Replication Table Statistics")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationTaskAssessmentResults, "describe-replication-task-assessment-results", "", false, "Describe Replication Task Assessment Results")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationTaskAssessmentRuns, "describe-replication-task-assessment-runs", "", false, "Describe Replication Task Assessment Runs")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationTaskIndividualAssessments, "describe-replication-task-individual-assessments", "", false, "Describe Replication Task Individual Assessments")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplicationTasks, "describe-replication-tasks", "", false, "Describe Replication Tasks")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeReplications, "describe-replications", "", false, "Describe Replications")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeSchemas, "describe-schemas", "", false, "Describe Schemas")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceDescribeTableStatistics, "describe-table-statistics", "", false, "Describe Table Statistics")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceExportMetadataModelAssessment, "export-metadata-model-assessment", "", false, "Export Metadata Model Assessment")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceGetTargetSelectionRules, "get-target-selection-rules", "", false, "Get Target Selection Rules")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceImportCertificate, "import-certificate", "", false, "Import Certificate")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyConversionConfiguration, "modify-conversion-configuration", "", false, "Modify Conversion Configuration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyDataMigration, "modify-data-migration", "", false, "Modify Data Migration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyDataProvider, "modify-data-provider", "", false, "Modify Data Provider")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyEndpoint, "modify-endpoint", "", false, "Modify Endpoint")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyEventSubscription, "modify-event-subscription", "", false, "Modify Event Subscription")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyInstanceProfile, "modify-instance-profile", "", false, "Modify Instance Profile")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyMigrationProject, "modify-migration-project", "", false, "Modify Migration Project")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyReplicationConfig, "modify-replication-config", "", false, "Modify Replication Config")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyReplicationInstance, "modify-replication-instance", "", false, "Modify Replication Instance")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyReplicationSubnetGroup, "modify-replication-subnet-group", "", false, "Modify Replication Subnet Group")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceModifyReplicationTask, "modify-replication-task", "", false, "Modify Replication Task")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceMoveReplicationTask, "move-replication-task", "", false, "Move Replication Task")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceRebootReplicationInstance, "reboot-replication-instance", "", false, "Reboot Replication Instance")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceRefreshSchemas, "refresh-schemas", "", false, "Refresh Schemas")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceReloadReplicationTables, "reload-replication-tables", "", false, "Reload Replication Tables")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceReloadTables, "reload-tables", "", false, "Reload Tables")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceRunFleetAdvisorLsaAnalysis, "run-fleet-advisor-lsa-analysis", "", false, "Run Fleet Advisor Lsa Analysis")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartDataMigration, "start-data-migration", "", false, "Start Data Migration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartExtensionPackAssociation, "start-extension-pack-association", "", false, "Start Extension Pack Association")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartMetadataModelAssessment, "start-metadata-model-assessment", "", false, "Start Metadata Model Assessment")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartMetadataModelConversion, "start-metadata-model-conversion", "", false, "Start Metadata Model Conversion")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartMetadataModelCreation, "start-metadata-model-creation", "", false, "Start Metadata Model Creation")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartMetadataModelExportAsScript, "start-metadata-model-export-as-script", "", false, "Start Metadata Model Export As Script")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartMetadataModelExportToTarget, "start-metadata-model-export-to-target", "", false, "Start Metadata Model Export To Target")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartMetadataModelImport, "start-metadata-model-import", "", false, "Start Metadata Model Import")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartRecommendations, "start-recommendations", "", false, "Start Recommendations")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartReplication, "start-replication", "", false, "Start Replication")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartReplicationTask, "start-replication-task", "", false, "Start Replication Task")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartReplicationTaskAssessment, "start-replication-task-assessment", "", false, "Start Replication Task Assessment")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStartReplicationTaskAssessmentRun, "start-replication-task-assessment-run", "", false, "Start Replication Task Assessment Run")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStopDataMigration, "stop-data-migration", "", false, "Stop Data Migration")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStopReplication, "stop-replication", "", false, "Stop Replication")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceStopReplicationTask, "stop-replication-task", "", false, "Stop Replication Task")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceTestConnection, "test-connection", "", false, "Test Connection")
	_databasemigrationserviceCmd.Flags().BoolVarP(&_databasemigrationserviceUpdateSubscriptionsToEventBridge, "update-subscriptions-to-event-bridge", "", false, "Update Subscriptions To Event Bridge")

}
