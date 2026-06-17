package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// glueCmd represents the glue command
var _glueCmd = &cobra.Command{
	Use:   "glue",
	Short: "AWS glue CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := glue.NewFromConfig(cfg)
		if _glueBatchCreatePartition {
			glue_BatchCreatePartition(cfg, client)
			return
		}
		if _glueBatchDeleteConnection {
			glue_BatchDeleteConnection(cfg, client)
			return
		}
		if _glueBatchDeletePartition {
			glue_BatchDeletePartition(cfg, client)
			return
		}
		if _glueBatchDeleteTable {
			glue_BatchDeleteTable(cfg, client)
			return
		}
		if _glueBatchDeleteTableVersion {
			glue_BatchDeleteTableVersion(cfg, client)
			return
		}
		if _glueBatchGetBlueprints {
			glue_BatchGetBlueprints(cfg, client)
			return
		}
		if _glueBatchGetCrawlers {
			glue_BatchGetCrawlers(cfg, client)
			return
		}
		if _glueBatchGetCustomEntityTypes {
			glue_BatchGetCustomEntityTypes(cfg, client)
			return
		}
		if _glueBatchGetDataQualityResult {
			glue_BatchGetDataQualityResult(cfg, client)
			return
		}
		if _glueBatchGetDevEndpoints {
			glue_BatchGetDevEndpoints(cfg, client)
			return
		}
		if _glueBatchGetJobs {
			glue_BatchGetJobs(cfg, client)
			return
		}
		if _glueBatchGetPartition {
			glue_BatchGetPartition(cfg, client)
			return
		}
		if _glueBatchGetTableOptimizer {
			glue_BatchGetTableOptimizer(cfg, client)
			return
		}
		if _glueBatchGetTriggers {
			glue_BatchGetTriggers(cfg, client)
			return
		}
		if _glueBatchGetWorkflows {
			glue_BatchGetWorkflows(cfg, client)
			return
		}
		if _glueBatchPutDataQualityStatisticAnnotation {
			glue_BatchPutDataQualityStatisticAnnotation(cfg, client)
			return
		}
		if _glueBatchStopJobRun {
			glue_BatchStopJobRun(cfg, client)
			return
		}
		if _glueBatchUpdatePartition {
			glue_BatchUpdatePartition(cfg, client)
			return
		}
		if _glueCancelDataQualityRuleRecommendationRun {
			glue_CancelDataQualityRuleRecommendationRun(cfg, client)
			return
		}
		if _glueCancelDataQualityRulesetEvaluationRun {
			glue_CancelDataQualityRulesetEvaluationRun(cfg, client)
			return
		}
		if _glueCancelMLTaskRun {
			glue_CancelMLTaskRun(cfg, client)
			return
		}
		if _glueCancelStatement {
			glue_CancelStatement(cfg, client)
			return
		}
		if _glueCheckSchemaVersionValidity {
			glue_CheckSchemaVersionValidity(cfg, client)
			return
		}
		if _glueCreateBlueprint {
			glue_CreateBlueprint(cfg, client)
			return
		}
		if _glueCreateCatalog {
			glue_CreateCatalog(cfg, client)
			return
		}
		if _glueCreateClassifier {
			glue_CreateClassifier(cfg, client)
			return
		}
		if _glueCreateColumnStatisticsTaskSettings {
			glue_CreateColumnStatisticsTaskSettings(cfg, client)
			return
		}
		if _glueCreateConnection {
			glue_CreateConnection(cfg, client)
			return
		}
		if _glueCreateCrawler {
			glue_CreateCrawler(cfg, client)
			return
		}
		if _glueCreateCustomEntityType {
			glue_CreateCustomEntityType(cfg, client)
			return
		}
		if _glueCreateDataQualityRuleset {
			glue_CreateDataQualityRuleset(cfg, client)
			return
		}
		if _glueCreateDatabase {
			glue_CreateDatabase(cfg, client)
			return
		}
		if _glueCreateDevEndpoint {
			glue_CreateDevEndpoint(cfg, client)
			return
		}
		if _glueCreateGlueIdentityCenterConfiguration {
			glue_CreateGlueIdentityCenterConfiguration(cfg, client)
			return
		}
		if _glueCreateIntegration {
			glue_CreateIntegration(cfg, client)
			return
		}
		if _glueCreateIntegrationResourceProperty {
			glue_CreateIntegrationResourceProperty(cfg, client)
			return
		}
		if _glueCreateIntegrationTableProperties {
			glue_CreateIntegrationTableProperties(cfg, client)
			return
		}
		if _glueCreateJob {
			glue_CreateJob(cfg, client)
			return
		}
		if _glueCreateMLTransform {
			glue_CreateMLTransform(cfg, client)
			return
		}
		if _glueCreatePartition {
			glue_CreatePartition(cfg, client)
			return
		}
		if _glueCreatePartitionIndex {
			glue_CreatePartitionIndex(cfg, client)
			return
		}
		if _glueCreateRegistry {
			glue_CreateRegistry(cfg, client)
			return
		}
		if _glueCreateSchema {
			glue_CreateSchema(cfg, client)
			return
		}
		if _glueCreateScript {
			glue_CreateScript(cfg, client)
			return
		}
		if _glueCreateSecurityConfiguration {
			glue_CreateSecurityConfiguration(cfg, client)
			return
		}
		if _glueCreateSession {
			glue_CreateSession(cfg, client)
			return
		}
		if _glueCreateTable {
			glue_CreateTable(cfg, client)
			return
		}
		if _glueCreateTableOptimizer {
			glue_CreateTableOptimizer(cfg, client)
			return
		}
		if _glueCreateTrigger {
			glue_CreateTrigger(cfg, client)
			return
		}
		if _glueCreateUsageProfile {
			glue_CreateUsageProfile(cfg, client)
			return
		}
		if _glueCreateUserDefinedFunction {
			glue_CreateUserDefinedFunction(cfg, client)
			return
		}
		if _glueCreateWorkflow {
			glue_CreateWorkflow(cfg, client)
			return
		}
		if _glueDeleteBlueprint {
			glue_DeleteBlueprint(cfg, client)
			return
		}
		if _glueDeleteCatalog {
			glue_DeleteCatalog(cfg, client)
			return
		}
		if _glueDeleteClassifier {
			glue_DeleteClassifier(cfg, client)
			return
		}
		if _glueDeleteColumnStatisticsForPartition {
			glue_DeleteColumnStatisticsForPartition(cfg, client)
			return
		}
		if _glueDeleteColumnStatisticsForTable {
			glue_DeleteColumnStatisticsForTable(cfg, client)
			return
		}
		if _glueDeleteColumnStatisticsTaskSettings {
			glue_DeleteColumnStatisticsTaskSettings(cfg, client)
			return
		}
		if _glueDeleteConnection {
			glue_DeleteConnection(cfg, client)
			return
		}
		if _glueDeleteConnectionType {
			glue_DeleteConnectionType(cfg, client)
			return
		}
		if _glueDeleteCrawler {
			glue_DeleteCrawler(cfg, client)
			return
		}
		if _glueDeleteCustomEntityType {
			glue_DeleteCustomEntityType(cfg, client)
			return
		}
		if _glueDeleteDataQualityRuleset {
			glue_DeleteDataQualityRuleset(cfg, client)
			return
		}
		if _glueDeleteDatabase {
			glue_DeleteDatabase(cfg, client)
			return
		}
		if _glueDeleteDevEndpoint {
			glue_DeleteDevEndpoint(cfg, client)
			return
		}
		if _glueDeleteGlueIdentityCenterConfiguration {
			glue_DeleteGlueIdentityCenterConfiguration(cfg, client)
			return
		}
		if _glueDeleteIntegration {
			glue_DeleteIntegration(cfg, client)
			return
		}
		if _glueDeleteIntegrationResourceProperty {
			glue_DeleteIntegrationResourceProperty(cfg, client)
			return
		}
		if _glueDeleteIntegrationTableProperties {
			glue_DeleteIntegrationTableProperties(cfg, client)
			return
		}
		if _glueDeleteJob {
			glue_DeleteJob(cfg, client)
			return
		}
		if _glueDeleteMLTransform {
			glue_DeleteMLTransform(cfg, client)
			return
		}
		if _glueDeletePartition {
			glue_DeletePartition(cfg, client)
			return
		}
		if _glueDeletePartitionIndex {
			glue_DeletePartitionIndex(cfg, client)
			return
		}
		if _glueDeleteRegistry {
			glue_DeleteRegistry(cfg, client)
			return
		}
		if _glueDeleteResourcePolicy {
			glue_DeleteResourcePolicy(cfg, client)
			return
		}
		if _glueDeleteSchema {
			glue_DeleteSchema(cfg, client)
			return
		}
		if _glueDeleteSchemaVersions {
			glue_DeleteSchemaVersions(cfg, client)
			return
		}
		if _glueDeleteSecurityConfiguration {
			glue_DeleteSecurityConfiguration(cfg, client)
			return
		}
		if _glueDeleteSession {
			glue_DeleteSession(cfg, client)
			return
		}
		if _glueDeleteTable {
			glue_DeleteTable(cfg, client)
			return
		}
		if _glueDeleteTableOptimizer {
			glue_DeleteTableOptimizer(cfg, client)
			return
		}
		if _glueDeleteTableVersion {
			glue_DeleteTableVersion(cfg, client)
			return
		}
		if _glueDeleteTrigger {
			glue_DeleteTrigger(cfg, client)
			return
		}
		if _glueDeleteUsageProfile {
			glue_DeleteUsageProfile(cfg, client)
			return
		}
		if _glueDeleteUserDefinedFunction {
			glue_DeleteUserDefinedFunction(cfg, client)
			return
		}
		if _glueDeleteWorkflow {
			glue_DeleteWorkflow(cfg, client)
			return
		}
		if _glueDescribeConnectionType {
			glue_DescribeConnectionType(cfg, client)
			return
		}
		if _glueDescribeEntity {
			glue_DescribeEntity(cfg, client)
			return
		}
		if _glueDescribeInboundIntegrations {
			glue_DescribeInboundIntegrations(cfg, client)
			return
		}
		if _glueDescribeIntegrations {
			glue_DescribeIntegrations(cfg, client)
			return
		}
		if _glueGetBlueprint {
			glue_GetBlueprint(cfg, client)
			return
		}
		if _glueGetBlueprintRun {
			glue_GetBlueprintRun(cfg, client)
			return
		}
		if _glueGetBlueprintRuns {
			glue_GetBlueprintRuns(cfg, client)
			return
		}
		if _glueGetCatalog {
			glue_GetCatalog(cfg, client)
			return
		}
		if _glueGetCatalogImportStatus {
			glue_GetCatalogImportStatus(cfg, client)
			return
		}
		if _glueGetCatalogs {
			glue_GetCatalogs(cfg, client)
			return
		}
		if _glueGetClassifier {
			glue_GetClassifier(cfg, client)
			return
		}
		if _glueGetClassifiers {
			glue_GetClassifiers(cfg, client)
			return
		}
		if _glueGetColumnStatisticsForPartition {
			glue_GetColumnStatisticsForPartition(cfg, client)
			return
		}
		if _glueGetColumnStatisticsForTable {
			glue_GetColumnStatisticsForTable(cfg, client)
			return
		}
		if _glueGetColumnStatisticsTaskRun {
			glue_GetColumnStatisticsTaskRun(cfg, client)
			return
		}
		if _glueGetColumnStatisticsTaskRuns {
			glue_GetColumnStatisticsTaskRuns(cfg, client)
			return
		}
		if _glueGetColumnStatisticsTaskSettings {
			glue_GetColumnStatisticsTaskSettings(cfg, client)
			return
		}
		if _glueGetConnection {
			glue_GetConnection(cfg, client)
			return
		}
		if _glueGetConnections {
			glue_GetConnections(cfg, client)
			return
		}
		if _glueGetCrawler {
			glue_GetCrawler(cfg, client)
			return
		}
		if _glueGetCrawlerMetrics {
			glue_GetCrawlerMetrics(cfg, client)
			return
		}
		if _glueGetCrawlers {
			glue_GetCrawlers(cfg, client)
			return
		}
		if _glueGetCustomEntityType {
			glue_GetCustomEntityType(cfg, client)
			return
		}
		if _glueGetDataCatalogEncryptionSettings {
			glue_GetDataCatalogEncryptionSettings(cfg, client)
			return
		}
		if _glueGetDataQualityModel {
			glue_GetDataQualityModel(cfg, client)
			return
		}
		if _glueGetDataQualityModelResult {
			glue_GetDataQualityModelResult(cfg, client)
			return
		}
		if _glueGetDataQualityResult {
			glue_GetDataQualityResult(cfg, client)
			return
		}
		if _glueGetDataQualityRuleRecommendationRun {
			glue_GetDataQualityRuleRecommendationRun(cfg, client)
			return
		}
		if _glueGetDataQualityRuleset {
			glue_GetDataQualityRuleset(cfg, client)
			return
		}
		if _glueGetDataQualityRulesetEvaluationRun {
			glue_GetDataQualityRulesetEvaluationRun(cfg, client)
			return
		}
		if _glueGetDatabase {
			glue_GetDatabase(cfg, client)
			return
		}
		if _glueGetDatabases {
			glue_GetDatabases(cfg, client)
			return
		}
		if _glueGetDataflowGraph {
			glue_GetDataflowGraph(cfg, client)
			return
		}
		if _glueGetDevEndpoint {
			glue_GetDevEndpoint(cfg, client)
			return
		}
		if _glueGetDevEndpoints {
			glue_GetDevEndpoints(cfg, client)
			return
		}
		if _glueGetEntityRecords {
			glue_GetEntityRecords(cfg, client)
			return
		}
		if _glueGetGlueIdentityCenterConfiguration {
			glue_GetGlueIdentityCenterConfiguration(cfg, client)
			return
		}
		if _glueGetIntegrationResourceProperty {
			glue_GetIntegrationResourceProperty(cfg, client)
			return
		}
		if _glueGetIntegrationTableProperties {
			glue_GetIntegrationTableProperties(cfg, client)
			return
		}
		if _glueGetJob {
			glue_GetJob(cfg, client)
			return
		}
		if _glueGetJobBookmark {
			glue_GetJobBookmark(cfg, client)
			return
		}
		if _glueGetJobRun {
			glue_GetJobRun(cfg, client)
			return
		}
		if _glueGetJobRuns {
			glue_GetJobRuns(cfg, client)
			return
		}
		if _glueGetJobs {
			glue_GetJobs(cfg, client)
			return
		}
		if _glueGetMapping {
			glue_GetMapping(cfg, client)
			return
		}
		if _glueGetMaterializedViewRefreshTaskRun {
			glue_GetMaterializedViewRefreshTaskRun(cfg, client)
			return
		}
		if _glueGetMLTaskRun {
			glue_GetMLTaskRun(cfg, client)
			return
		}
		if _glueGetMLTaskRuns {
			glue_GetMLTaskRuns(cfg, client)
			return
		}
		if _glueGetMLTransform {
			glue_GetMLTransform(cfg, client)
			return
		}
		if _glueGetMLTransforms {
			glue_GetMLTransforms(cfg, client)
			return
		}
		if _glueGetPartition {
			glue_GetPartition(cfg, client)
			return
		}
		if _glueGetPartitionIndexes {
			glue_GetPartitionIndexes(cfg, client)
			return
		}
		if _glueGetPartitions {
			glue_GetPartitions(cfg, client)
			return
		}
		if _glueGetPlan {
			glue_GetPlan(cfg, client)
			return
		}
		if _glueGetRegistry {
			glue_GetRegistry(cfg, client)
			return
		}
		if _glueGetResourcePolicies {
			glue_GetResourcePolicies(cfg, client)
			return
		}
		if _glueGetResourcePolicy {
			glue_GetResourcePolicy(cfg, client)
			return
		}
		if _glueGetSchema {
			glue_GetSchema(cfg, client)
			return
		}
		if _glueGetSchemaByDefinition {
			glue_GetSchemaByDefinition(cfg, client)
			return
		}
		if _glueGetSchemaVersion {
			glue_GetSchemaVersion(cfg, client)
			return
		}
		if _glueGetSchemaVersionsDiff {
			glue_GetSchemaVersionsDiff(cfg, client)
			return
		}
		if _glueGetSecurityConfiguration {
			glue_GetSecurityConfiguration(cfg, client)
			return
		}
		if _glueGetSecurityConfigurations {
			glue_GetSecurityConfigurations(cfg, client)
			return
		}
		if _glueGetSession {
			glue_GetSession(cfg, client)
			return
		}
		if _glueGetStatement {
			glue_GetStatement(cfg, client)
			return
		}
		if _glueGetTable {
			glue_GetTable(cfg, client)
			return
		}
		if _glueGetTableOptimizer {
			glue_GetTableOptimizer(cfg, client)
			return
		}
		if _glueGetTableVersion {
			glue_GetTableVersion(cfg, client)
			return
		}
		if _glueGetTableVersions {
			glue_GetTableVersions(cfg, client)
			return
		}
		if _glueGetTables {
			glue_GetTables(cfg, client)
			return
		}
		if _glueGetTags {
			glue_GetTags(cfg, client)
			return
		}
		if _glueGetTrigger {
			glue_GetTrigger(cfg, client)
			return
		}
		if _glueGetTriggers {
			glue_GetTriggers(cfg, client)
			return
		}
		if _glueGetUnfilteredPartitionMetadata {
			glue_GetUnfilteredPartitionMetadata(cfg, client)
			return
		}
		if _glueGetUnfilteredPartitionsMetadata {
			glue_GetUnfilteredPartitionsMetadata(cfg, client)
			return
		}
		if _glueGetUnfilteredTableMetadata {
			glue_GetUnfilteredTableMetadata(cfg, client)
			return
		}
		if _glueGetUsageProfile {
			glue_GetUsageProfile(cfg, client)
			return
		}
		if _glueGetUserDefinedFunction {
			glue_GetUserDefinedFunction(cfg, client)
			return
		}
		if _glueGetUserDefinedFunctions {
			glue_GetUserDefinedFunctions(cfg, client)
			return
		}
		if _glueGetWorkflow {
			glue_GetWorkflow(cfg, client)
			return
		}
		if _glueGetWorkflowRun {
			glue_GetWorkflowRun(cfg, client)
			return
		}
		if _glueGetWorkflowRunProperties {
			glue_GetWorkflowRunProperties(cfg, client)
			return
		}
		if _glueGetWorkflowRuns {
			glue_GetWorkflowRuns(cfg, client)
			return
		}
		if _glueImportCatalogToGlue {
			glue_ImportCatalogToGlue(cfg, client)
			return
		}
		if _glueListBlueprints {
			glue_ListBlueprints(cfg, client)
			return
		}
		if _glueListColumnStatisticsTaskRuns {
			glue_ListColumnStatisticsTaskRuns(cfg, client)
			return
		}
		if _glueListConnectionTypes {
			glue_ListConnectionTypes(cfg, client)
			return
		}
		if _glueListCrawlers {
			glue_ListCrawlers(cfg, client)
			return
		}
		if _glueListCrawls {
			glue_ListCrawls(cfg, client)
			return
		}
		if _glueListCustomEntityTypes {
			glue_ListCustomEntityTypes(cfg, client)
			return
		}
		if _glueListDataQualityResults {
			glue_ListDataQualityResults(cfg, client)
			return
		}
		if _glueListDataQualityRuleRecommendationRuns {
			glue_ListDataQualityRuleRecommendationRuns(cfg, client)
			return
		}
		if _glueListDataQualityRulesetEvaluationRuns {
			glue_ListDataQualityRulesetEvaluationRuns(cfg, client)
			return
		}
		if _glueListDataQualityRulesets {
			glue_ListDataQualityRulesets(cfg, client)
			return
		}
		if _glueListDataQualityStatisticAnnotations {
			glue_ListDataQualityStatisticAnnotations(cfg, client)
			return
		}
		if _glueListDataQualityStatistics {
			glue_ListDataQualityStatistics(cfg, client)
			return
		}
		if _glueListDevEndpoints {
			glue_ListDevEndpoints(cfg, client)
			return
		}
		if _glueListEntities {
			glue_ListEntities(cfg, client)
			return
		}
		if _glueListIntegrationResourceProperties {
			glue_ListIntegrationResourceProperties(cfg, client)
			return
		}
		if _glueListJobs {
			glue_ListJobs(cfg, client)
			return
		}
		if _glueListMaterializedViewRefreshTaskRuns {
			glue_ListMaterializedViewRefreshTaskRuns(cfg, client)
			return
		}
		if _glueListMLTransforms {
			glue_ListMLTransforms(cfg, client)
			return
		}
		if _glueListRegistries {
			glue_ListRegistries(cfg, client)
			return
		}
		if _glueListSchemaVersions {
			glue_ListSchemaVersions(cfg, client)
			return
		}
		if _glueListSchemas {
			glue_ListSchemas(cfg, client)
			return
		}
		if _glueListSessions {
			glue_ListSessions(cfg, client)
			return
		}
		if _glueListStatements {
			glue_ListStatements(cfg, client)
			return
		}
		if _glueListTableOptimizerRuns {
			glue_ListTableOptimizerRuns(cfg, client)
			return
		}
		if _glueListTriggers {
			glue_ListTriggers(cfg, client)
			return
		}
		if _glueListUsageProfiles {
			glue_ListUsageProfiles(cfg, client)
			return
		}
		if _glueListWorkflows {
			glue_ListWorkflows(cfg, client)
			return
		}
		if _glueModifyIntegration {
			glue_ModifyIntegration(cfg, client)
			return
		}
		if _gluePutDataCatalogEncryptionSettings {
			glue_PutDataCatalogEncryptionSettings(cfg, client)
			return
		}
		if _gluePutDataQualityProfileAnnotation {
			glue_PutDataQualityProfileAnnotation(cfg, client)
			return
		}
		if _gluePutResourcePolicy {
			glue_PutResourcePolicy(cfg, client)
			return
		}
		if _gluePutSchemaVersionMetadata {
			glue_PutSchemaVersionMetadata(cfg, client)
			return
		}
		if _gluePutWorkflowRunProperties {
			glue_PutWorkflowRunProperties(cfg, client)
			return
		}
		if _glueQuerySchemaVersionMetadata {
			glue_QuerySchemaVersionMetadata(cfg, client)
			return
		}
		if _glueRegisterConnectionType {
			glue_RegisterConnectionType(cfg, client)
			return
		}
		if _glueRegisterSchemaVersion {
			glue_RegisterSchemaVersion(cfg, client)
			return
		}
		if _glueRemoveSchemaVersionMetadata {
			glue_RemoveSchemaVersionMetadata(cfg, client)
			return
		}
		if _glueResetJobBookmark {
			glue_ResetJobBookmark(cfg, client)
			return
		}
		if _glueResumeWorkflowRun {
			glue_ResumeWorkflowRun(cfg, client)
			return
		}
		if _glueRunStatement {
			glue_RunStatement(cfg, client)
			return
		}
		if _glueSearchTables {
			glue_SearchTables(cfg, client)
			return
		}
		if _glueStartBlueprintRun {
			glue_StartBlueprintRun(cfg, client)
			return
		}
		if _glueStartColumnStatisticsTaskRun {
			glue_StartColumnStatisticsTaskRun(cfg, client)
			return
		}
		if _glueStartColumnStatisticsTaskRunSchedule {
			glue_StartColumnStatisticsTaskRunSchedule(cfg, client)
			return
		}
		if _glueStartCrawler {
			glue_StartCrawler(cfg, client)
			return
		}
		if _glueStartCrawlerSchedule {
			glue_StartCrawlerSchedule(cfg, client)
			return
		}
		if _glueStartDataQualityRuleRecommendationRun {
			glue_StartDataQualityRuleRecommendationRun(cfg, client)
			return
		}
		if _glueStartDataQualityRulesetEvaluationRun {
			glue_StartDataQualityRulesetEvaluationRun(cfg, client)
			return
		}
		if _glueStartExportLabelsTaskRun {
			glue_StartExportLabelsTaskRun(cfg, client)
			return
		}
		if _glueStartImportLabelsTaskRun {
			glue_StartImportLabelsTaskRun(cfg, client)
			return
		}
		if _glueStartJobRun {
			glue_StartJobRun(cfg, client)
			return
		}
		if _glueStartMaterializedViewRefreshTaskRun {
			glue_StartMaterializedViewRefreshTaskRun(cfg, client)
			return
		}
		if _glueStartMLEvaluationTaskRun {
			glue_StartMLEvaluationTaskRun(cfg, client)
			return
		}
		if _glueStartMLLabelingSetGenerationTaskRun {
			glue_StartMLLabelingSetGenerationTaskRun(cfg, client)
			return
		}
		if _glueStartTrigger {
			glue_StartTrigger(cfg, client)
			return
		}
		if _glueStartWorkflowRun {
			glue_StartWorkflowRun(cfg, client)
			return
		}
		if _glueStopColumnStatisticsTaskRun {
			glue_StopColumnStatisticsTaskRun(cfg, client)
			return
		}
		if _glueStopColumnStatisticsTaskRunSchedule {
			glue_StopColumnStatisticsTaskRunSchedule(cfg, client)
			return
		}
		if _glueStopCrawler {
			glue_StopCrawler(cfg, client)
			return
		}
		if _glueStopCrawlerSchedule {
			glue_StopCrawlerSchedule(cfg, client)
			return
		}
		if _glueStopMaterializedViewRefreshTaskRun {
			glue_StopMaterializedViewRefreshTaskRun(cfg, client)
			return
		}
		if _glueStopSession {
			glue_StopSession(cfg, client)
			return
		}
		if _glueStopTrigger {
			glue_StopTrigger(cfg, client)
			return
		}
		if _glueStopWorkflowRun {
			glue_StopWorkflowRun(cfg, client)
			return
		}
		if _glueTagResource {
			glue_TagResource(cfg, client)
			return
		}
		if _glueTestConnection {
			glue_TestConnection(cfg, client)
			return
		}
		if _glueUntagResource {
			glue_UntagResource(cfg, client)
			return
		}
		if _glueUpdateBlueprint {
			glue_UpdateBlueprint(cfg, client)
			return
		}
		if _glueUpdateCatalog {
			glue_UpdateCatalog(cfg, client)
			return
		}
		if _glueUpdateClassifier {
			glue_UpdateClassifier(cfg, client)
			return
		}
		if _glueUpdateColumnStatisticsForPartition {
			glue_UpdateColumnStatisticsForPartition(cfg, client)
			return
		}
		if _glueUpdateColumnStatisticsForTable {
			glue_UpdateColumnStatisticsForTable(cfg, client)
			return
		}
		if _glueUpdateColumnStatisticsTaskSettings {
			glue_UpdateColumnStatisticsTaskSettings(cfg, client)
			return
		}
		if _glueUpdateConnection {
			glue_UpdateConnection(cfg, client)
			return
		}
		if _glueUpdateCrawler {
			glue_UpdateCrawler(cfg, client)
			return
		}
		if _glueUpdateCrawlerSchedule {
			glue_UpdateCrawlerSchedule(cfg, client)
			return
		}
		if _glueUpdateDataQualityRuleset {
			glue_UpdateDataQualityRuleset(cfg, client)
			return
		}
		if _glueUpdateDatabase {
			glue_UpdateDatabase(cfg, client)
			return
		}
		if _glueUpdateDevEndpoint {
			glue_UpdateDevEndpoint(cfg, client)
			return
		}
		if _glueUpdateGlueIdentityCenterConfiguration {
			glue_UpdateGlueIdentityCenterConfiguration(cfg, client)
			return
		}
		if _glueUpdateIntegrationResourceProperty {
			glue_UpdateIntegrationResourceProperty(cfg, client)
			return
		}
		if _glueUpdateIntegrationTableProperties {
			glue_UpdateIntegrationTableProperties(cfg, client)
			return
		}
		if _glueUpdateJob {
			glue_UpdateJob(cfg, client)
			return
		}
		if _glueUpdateJobFromSourceControl {
			glue_UpdateJobFromSourceControl(cfg, client)
			return
		}
		if _glueUpdateMLTransform {
			glue_UpdateMLTransform(cfg, client)
			return
		}
		if _glueUpdatePartition {
			glue_UpdatePartition(cfg, client)
			return
		}
		if _glueUpdateRegistry {
			glue_UpdateRegistry(cfg, client)
			return
		}
		if _glueUpdateSchema {
			glue_UpdateSchema(cfg, client)
			return
		}
		if _glueUpdateSourceControlFromJob {
			glue_UpdateSourceControlFromJob(cfg, client)
			return
		}
		if _glueUpdateTable {
			glue_UpdateTable(cfg, client)
			return
		}
		if _glueUpdateTableOptimizer {
			glue_UpdateTableOptimizer(cfg, client)
			return
		}
		if _glueUpdateTrigger {
			glue_UpdateTrigger(cfg, client)
			return
		}
		if _glueUpdateUsageProfile {
			glue_UpdateUsageProfile(cfg, client)
			return
		}
		if _glueUpdateUserDefinedFunction {
			glue_UpdateUserDefinedFunction(cfg, client)
			return
		}
		if _glueUpdateWorkflow {
			glue_UpdateWorkflow(cfg, client)
			return
		}

	},
}

var (
	_glueBatchCreatePartition                   bool
	_glueBatchDeleteConnection                  bool
	_glueBatchDeletePartition                   bool
	_glueBatchDeleteTable                       bool
	_glueBatchDeleteTableVersion                bool
	_glueBatchGetBlueprints                     bool
	_glueBatchGetCrawlers                       bool
	_glueBatchGetCustomEntityTypes              bool
	_glueBatchGetDataQualityResult              bool
	_glueBatchGetDevEndpoints                   bool
	_glueBatchGetJobs                           bool
	_glueBatchGetPartition                      bool
	_glueBatchGetTableOptimizer                 bool
	_glueBatchGetTriggers                       bool
	_glueBatchGetWorkflows                      bool
	_glueBatchPutDataQualityStatisticAnnotation bool
	_glueBatchStopJobRun                        bool
	_glueBatchUpdatePartition                   bool
	_glueCancelDataQualityRuleRecommendationRun bool
	_glueCancelDataQualityRulesetEvaluationRun  bool
	_glueCancelMLTaskRun                        bool
	_glueCancelStatement                        bool
	_glueCheckSchemaVersionValidity             bool
	_glueCreateBlueprint                        bool
	_glueCreateCatalog                          bool
	_glueCreateClassifier                       bool
	_glueCreateColumnStatisticsTaskSettings     bool
	_glueCreateConnection                       bool
	_glueCreateCrawler                          bool
	_glueCreateCustomEntityType                 bool
	_glueCreateDataQualityRuleset               bool
	_glueCreateDatabase                         bool
	_glueCreateDevEndpoint                      bool
	_glueCreateGlueIdentityCenterConfiguration  bool
	_glueCreateIntegration                      bool
	_glueCreateIntegrationResourceProperty      bool
	_glueCreateIntegrationTableProperties       bool
	_glueCreateJob                              bool
	_glueCreateMLTransform                      bool
	_glueCreatePartition                        bool
	_glueCreatePartitionIndex                   bool
	_glueCreateRegistry                         bool
	_glueCreateSchema                           bool
	_glueCreateScript                           bool
	_glueCreateSecurityConfiguration            bool
	_glueCreateSession                          bool
	_glueCreateTable                            bool
	_glueCreateTableOptimizer                   bool
	_glueCreateTrigger                          bool
	_glueCreateUsageProfile                     bool
	_glueCreateUserDefinedFunction              bool
	_glueCreateWorkflow                         bool
	_glueDeleteBlueprint                        bool
	_glueDeleteCatalog                          bool
	_glueDeleteClassifier                       bool
	_glueDeleteColumnStatisticsForPartition     bool
	_glueDeleteColumnStatisticsForTable         bool
	_glueDeleteColumnStatisticsTaskSettings     bool
	_glueDeleteConnection                       bool
	_glueDeleteConnectionType                   bool
	_glueDeleteCrawler                          bool
	_glueDeleteCustomEntityType                 bool
	_glueDeleteDataQualityRuleset               bool
	_glueDeleteDatabase                         bool
	_glueDeleteDevEndpoint                      bool
	_glueDeleteGlueIdentityCenterConfiguration  bool
	_glueDeleteIntegration                      bool
	_glueDeleteIntegrationResourceProperty      bool
	_glueDeleteIntegrationTableProperties       bool
	_glueDeleteJob                              bool
	_glueDeleteMLTransform                      bool
	_glueDeletePartition                        bool
	_glueDeletePartitionIndex                   bool
	_glueDeleteRegistry                         bool
	_glueDeleteResourcePolicy                   bool
	_glueDeleteSchema                           bool
	_glueDeleteSchemaVersions                   bool
	_glueDeleteSecurityConfiguration            bool
	_glueDeleteSession                          bool
	_glueDeleteTable                            bool
	_glueDeleteTableOptimizer                   bool
	_glueDeleteTableVersion                     bool
	_glueDeleteTrigger                          bool
	_glueDeleteUsageProfile                     bool
	_glueDeleteUserDefinedFunction              bool
	_glueDeleteWorkflow                         bool
	_glueDescribeConnectionType                 bool
	_glueDescribeEntity                         bool
	_glueDescribeInboundIntegrations            bool
	_glueDescribeIntegrations                   bool
	_glueGetBlueprint                           bool
	_glueGetBlueprintRun                        bool
	_glueGetBlueprintRuns                       bool
	_glueGetCatalog                             bool
	_glueGetCatalogImportStatus                 bool
	_glueGetCatalogs                            bool
	_glueGetClassifier                          bool
	_glueGetClassifiers                         bool
	_glueGetColumnStatisticsForPartition        bool
	_glueGetColumnStatisticsForTable            bool
	_glueGetColumnStatisticsTaskRun             bool
	_glueGetColumnStatisticsTaskRuns            bool
	_glueGetColumnStatisticsTaskSettings        bool
	_glueGetConnection                          bool
	_glueGetConnections                         bool
	_glueGetCrawler                             bool
	_glueGetCrawlerMetrics                      bool
	_glueGetCrawlers                            bool
	_glueGetCustomEntityType                    bool
	_glueGetDataCatalogEncryptionSettings       bool
	_glueGetDataQualityModel                    bool
	_glueGetDataQualityModelResult              bool
	_glueGetDataQualityResult                   bool
	_glueGetDataQualityRuleRecommendationRun    bool
	_glueGetDataQualityRuleset                  bool
	_glueGetDataQualityRulesetEvaluationRun     bool
	_glueGetDatabase                            bool
	_glueGetDatabases                           bool
	_glueGetDataflowGraph                       bool
	_glueGetDevEndpoint                         bool
	_glueGetDevEndpoints                        bool
	_glueGetEntityRecords                       bool
	_glueGetGlueIdentityCenterConfiguration     bool
	_glueGetIntegrationResourceProperty         bool
	_glueGetIntegrationTableProperties          bool
	_glueGetJob                                 bool
	_glueGetJobBookmark                         bool
	_glueGetJobRun                              bool
	_glueGetJobRuns                             bool
	_glueGetJobs                                bool
	_glueGetMapping                             bool
	_glueGetMaterializedViewRefreshTaskRun      bool
	_glueGetMLTaskRun                           bool
	_glueGetMLTaskRuns                          bool
	_glueGetMLTransform                         bool
	_glueGetMLTransforms                        bool
	_glueGetPartition                           bool
	_glueGetPartitionIndexes                    bool
	_glueGetPartitions                          bool
	_glueGetPlan                                bool
	_glueGetRegistry                            bool
	_glueGetResourcePolicies                    bool
	_glueGetResourcePolicy                      bool
	_glueGetSchema                              bool
	_glueGetSchemaByDefinition                  bool
	_glueGetSchemaVersion                       bool
	_glueGetSchemaVersionsDiff                  bool
	_glueGetSecurityConfiguration               bool
	_glueGetSecurityConfigurations              bool
	_glueGetSession                             bool
	_glueGetStatement                           bool
	_glueGetTable                               bool
	_glueGetTableOptimizer                      bool
	_glueGetTableVersion                        bool
	_glueGetTableVersions                       bool
	_glueGetTables                              bool
	_glueGetTags                                bool
	_glueGetTrigger                             bool
	_glueGetTriggers                            bool
	_glueGetUnfilteredPartitionMetadata         bool
	_glueGetUnfilteredPartitionsMetadata        bool
	_glueGetUnfilteredTableMetadata             bool
	_glueGetUsageProfile                        bool
	_glueGetUserDefinedFunction                 bool
	_glueGetUserDefinedFunctions                bool
	_glueGetWorkflow                            bool
	_glueGetWorkflowRun                         bool
	_glueGetWorkflowRunProperties               bool
	_glueGetWorkflowRuns                        bool
	_glueImportCatalogToGlue                    bool
	_glueListBlueprints                         bool
	_glueListColumnStatisticsTaskRuns           bool
	_glueListConnectionTypes                    bool
	_glueListCrawlers                           bool
	_glueListCrawls                             bool
	_glueListCustomEntityTypes                  bool
	_glueListDataQualityResults                 bool
	_glueListDataQualityRuleRecommendationRuns  bool
	_glueListDataQualityRulesetEvaluationRuns   bool
	_glueListDataQualityRulesets                bool
	_glueListDataQualityStatisticAnnotations    bool
	_glueListDataQualityStatistics              bool
	_glueListDevEndpoints                       bool
	_glueListEntities                           bool
	_glueListIntegrationResourceProperties      bool
	_glueListJobs                               bool
	_glueListMaterializedViewRefreshTaskRuns    bool
	_glueListMLTransforms                       bool
	_glueListRegistries                         bool
	_glueListSchemaVersions                     bool
	_glueListSchemas                            bool
	_glueListSessions                           bool
	_glueListStatements                         bool
	_glueListTableOptimizerRuns                 bool
	_glueListTriggers                           bool
	_glueListUsageProfiles                      bool
	_glueListWorkflows                          bool
	_glueModifyIntegration                      bool
	_gluePutDataCatalogEncryptionSettings       bool
	_gluePutDataQualityProfileAnnotation        bool
	_gluePutResourcePolicy                      bool
	_gluePutSchemaVersionMetadata               bool
	_gluePutWorkflowRunProperties               bool
	_glueQuerySchemaVersionMetadata             bool
	_glueRegisterConnectionType                 bool
	_glueRegisterSchemaVersion                  bool
	_glueRemoveSchemaVersionMetadata            bool
	_glueResetJobBookmark                       bool
	_glueResumeWorkflowRun                      bool
	_glueRunStatement                           bool
	_glueSearchTables                           bool
	_glueStartBlueprintRun                      bool
	_glueStartColumnStatisticsTaskRun           bool
	_glueStartColumnStatisticsTaskRunSchedule   bool
	_glueStartCrawler                           bool
	_glueStartCrawlerSchedule                   bool
	_glueStartDataQualityRuleRecommendationRun  bool
	_glueStartDataQualityRulesetEvaluationRun   bool
	_glueStartExportLabelsTaskRun               bool
	_glueStartImportLabelsTaskRun               bool
	_glueStartJobRun                            bool
	_glueStartMaterializedViewRefreshTaskRun    bool
	_glueStartMLEvaluationTaskRun               bool
	_glueStartMLLabelingSetGenerationTaskRun    bool
	_glueStartTrigger                           bool
	_glueStartWorkflowRun                       bool
	_glueStopColumnStatisticsTaskRun            bool
	_glueStopColumnStatisticsTaskRunSchedule    bool
	_glueStopCrawler                            bool
	_glueStopCrawlerSchedule                    bool
	_glueStopMaterializedViewRefreshTaskRun     bool
	_glueStopSession                            bool
	_glueStopTrigger                            bool
	_glueStopWorkflowRun                        bool
	_glueTagResource                            bool
	_glueTestConnection                         bool
	_glueUntagResource                          bool
	_glueUpdateBlueprint                        bool
	_glueUpdateCatalog                          bool
	_glueUpdateClassifier                       bool
	_glueUpdateColumnStatisticsForPartition     bool
	_glueUpdateColumnStatisticsForTable         bool
	_glueUpdateColumnStatisticsTaskSettings     bool
	_glueUpdateConnection                       bool
	_glueUpdateCrawler                          bool
	_glueUpdateCrawlerSchedule                  bool
	_glueUpdateDataQualityRuleset               bool
	_glueUpdateDatabase                         bool
	_glueUpdateDevEndpoint                      bool
	_glueUpdateGlueIdentityCenterConfiguration  bool
	_glueUpdateIntegrationResourceProperty      bool
	_glueUpdateIntegrationTableProperties       bool
	_glueUpdateJob                              bool
	_glueUpdateJobFromSourceControl             bool
	_glueUpdateMLTransform                      bool
	_glueUpdatePartition                        bool
	_glueUpdateRegistry                         bool
	_glueUpdateSchema                           bool
	_glueUpdateSourceControlFromJob             bool
	_glueUpdateTable                            bool
	_glueUpdateTableOptimizer                   bool
	_glueUpdateTrigger                          bool
	_glueUpdateUsageProfile                     bool
	_glueUpdateUserDefinedFunction              bool
	_glueUpdateWorkflow                         bool

	_glueActions                              string
	_glueAddArguments                         string
	_glueAddPublicKeys                        []string
	_glueAdditionalDataSources                string
	_glueAdditionalEncryptionContext          string
	_glueAdditionalPlanOptionsMap             string
	_glueAdditionalRunOptions                 string
	_glueAllocatedCapacity                    string
	_glueApplyOverrideForComputeEnvironment   string
	_glueArguments                            string
	_glueAttributesToGet                      string
	_glueAuditContext                         string
	_glueAuthStrategy                         string
	_glueAuthToken                            string
	_glueBlueprintLocation                    string
	_glueBlueprintName                        string
	_glueBranchName                           string
	_glueCatalogId                            string
	_glueCatalogInput                         string
	_glueClassifiers                          []string
	_glueClientToken                          string
	_glueCode                                 string
	_glueCodeGenConfigurationNodes            string
	_glueColumnName                           string
	_glueColumnNameList                       []string
	_glueColumnNames                          []string
	_glueColumnStatisticsList                 string
	_glueColumnStatisticsTaskRunId            string
	_glueCommand                              string
	_glueCommitId                             string
	_glueCompatibility                        string
	_glueConfiguration                        string
	_glueConnectionInput                      string
	_glueConnectionName                       string
	_glueConnectionNameList                   []string
	_glueConnectionOptions                    string
	_glueConnectionProperties                 string
	_glueConnectionType                       string
	_glueConnections                          string
	_glueConnectorAuthenticationConfiguration string
	_glueContextWords                         []string
	_glueCrawlerName                          string
	_glueCrawlerNameList                      []string
	_glueCrawlerNames                         []string
	_glueCrawlerSecurityConfiguration         string
	_glueCreatedRulesetName                   string
	_glueCsvClassifier                        string
	_glueCustomLibraries                      string
	_glueDagEdges                             string
	_glueDagNodes                             string
	_glueDataCatalogEncryptionSettings        string
	_glueDataFilter                           string
	_glueDataFormat                           string
	_glueDataQualitySecurityConfiguration     string
	_glueDataSource                           string
	_glueDataStoreApiVersion                  string
	_glueDatabaseInput                        string
	_glueDatabaseName                         string
	_glueDefaultArguments                     string
	_glueDefaultRunProperties                 string
	_glueDeleteArguments                      []string
	_glueDeletePublicKeys                     []string
	_glueDependentJobName                     string
	_glueDescription                          string
	_glueDevEndpointNames                     []string
	_glueEnableHybrid                         string
	_glueEncryptionConfiguration              string
	_glueEndpointName                         string
	_glueEntityName                           string
	_glueEntries                              string
	_glueEventBatchingCondition               string
	_glueExcludeColumnSchema                  string
	_glueExecutionClass                       string
	_glueExecutionProperty                    string
	_glueExecutionRoleSessionPolicy           string
	_glueExpression                           string
	_glueExtraJarsS3Path                      string
	_glueExtraPythonLibsS3Path                string
	_glueFilter                               string
	_glueFilterPredicate                      string
	_glueFilters                              string
	_glueFirstSchemaVersionNumber             string
	_glueFolder                               string
	_glueForce                                string
	_glueFullRefresh                          string
	_glueFunctionInput                        string
	_glueFunctionName                         string
	_glueFunctionType                         string
	_glueGlueVersion                          string
	_glueGrokClassifier                       string
	_glueHidePassword                         string
	_glueId                                   string
	_glueIdleTimeout                          string
	_glueIncludeBlueprint                     string
	_glueIncludeGraph                         string
	_glueIncludeParameterSpec                 string
	_glueIncludeRoot                          string
	_glueIncludeStatusDetails                 string
	_glueInclusionAnnotation                  string
	_glueInclusionAnnotations                 string
	_glueIndexName                            string
	_glueInputRecordTables                    string
	_glueInputS3Path                          string
	_glueInstanceArn                          string
	_glueIntegrationArn                       string
	_glueIntegrationConfig                    string
	_glueIntegrationIdentifier                string
	_glueIntegrationName                      string
	_glueIntegrationType                      string
	_glueJobMode                              string
	_glueJobName                              string
	_glueJobNames                             []string
	_glueJobRunId                             string
	_glueJobRunIds                            []string
	_glueJobRunQueuingEnabled                 string
	_glueJobUpdate                            string
	_glueJsonClassifier                       string
	_glueKmsKeyId                             string
	_glueLakeFormationConfiguration           string
	_glueLanguage                             string
	_glueLimit                                string
	_glueLineageConfiguration                 string
	_glueLocation                             string
	_glueLogUri                               string
	_glueMaintenanceWindow                    string
	_glueMapping                              string
	_glueMarker                               string
	_glueMaterializedViewRefreshTaskRunId     string
	_glueMaxCapacity                          string
	_glueMaxConcurrentRuns                    string
	_glueMaxRecords                           string
	_glueMaxResults                           string
	_glueMaxRetries                           string
	_glueMetadataKeyValue                     string
	_glueMetadataList                         string
	_glueName                                 string
	_glueNames                                []string
	_glueNextToken                            string
	_glueNodeIds                              []string
	_glueNonOverridableArguments              string
	_glueNotificationProperty                 string
	_glueNumberOfNodes                        string
	_glueNumberOfWorkers                      string
	_glueOpenTableFormatInput                 string
	_glueOrderBy                              string
	_glueOutputS3Path                         string
	_glueParameters                           string
	_glueParentCatalogId                      string
	_glueParentEntityName                     string
	_glueParentResourceArn                    string
	_gluePartitionIndex                       string
	_gluePartitionIndexes                     string
	_gluePartitionInput                       string
	_gluePartitionInputList                   string
	_gluePartitionValueList                   []string
	_gluePartitionValues                      []string
	_gluePartitionsToDelete                   string
	_gluePartitionsToGet                      string
	_gluePattern                              string
	_gluePermissions                          string
	_gluePolicyExistsCondition                string
	_gluePolicyHashCondition                  string
	_gluePolicyInJson                         string
	_gluePredecessorsIncluded                 string
	_gluePredicate                            string
	_glueProfileId                            string
	_glueProvider                             string
	_gluePublicKey                            string
	_gluePublicKeys                           []string
	_gluePythonScript                         string
	_glueQueryAsOfTime                        string
	_glueQuerySessionContext                  string
	_glueRecrawlPolicy                        string
	_glueRecursive                            string
	_glueRegexString                          string
	_glueRegistryId                           string
	_glueRegistryName                         string
	_glueReplaceAllLabels                     string
	_glueRepositoryName                       string
	_glueRepositoryOwner                      string
	_glueRequestOrigin                        string
	_glueResourceArn                          string
	_glueResourceShareType                    string
	_glueRestConfiguration                    string
	_glueResultId                             string
	_glueResultIds                            []string
	_glueRole                                 string
	_glueRoleArn                              string
	_glueRootResourceArn                      string
	_glueRuleset                              string
	_glueRulesetNames                         []string
	_glueRunId                                string
	_glueRunProperties                        string
	_glueSampleSize                           string
	_glueSchedule                             string
	_glueSchemaChangePolicy                   string
	_glueSchemaDefinition                     string
	_glueSchemaDiffType                       string
	_glueSchemaId                             string
	_glueSchemaName                           string
	_glueSchemaVersionId                      string
	_glueSchemaVersionNumber                  string
	_glueScopes                               []string
	_glueSearchText                           string
	_glueSecondSchemaVersionNumber            string
	_glueSecurityConfiguration                string
	_glueSecurityGroupIds                     []string
	_glueSegment                              string
	_glueSelectedFields                       []string
	_glueSessionId                            string
	_glueSinks                                string
	_glueSkipArchive                          string
	_glueSort                                 string
	_glueSortCriteria                         string
	_glueSource                               string
	_glueSourceArn                            string
	_glueSourceControlDetails                 string
	_glueSourceProcessingProperties           string
	_glueSourceTableConfig                    string
	_glueStartOnCreation                      string
	_glueStatisticId                          string
	_glueSubnetId                             string
	_glueSupportedDialect                     string
	_glueSupportedPermissionTypes             string
	_glueTableInput                           string
	_glueTableName                            string
	_glueTableOptimizerConfiguration          string
	_glueTablePrefix                          string
	_glueTablesToDelete                       []string
	_glueTags                                 string
	_glueTagsToAdd                            string
	_glueTagsToRemove                         []string
	_glueTargetArn                            string
	_glueTargetProcessingProperties           string
	_glueTargetTable                          string
	_glueTargetTableConfig                    string
	_glueTargets                              string
	_glueTaskRunId                            string
	_glueTestConnectionInput                  string
	_glueTimeout                              string
	_glueTimestampFilter                      string
	_glueTransactionId                        string
	_glueTransformEncryption                  string
	_glueTransformId                          string
	_glueTriggerNames                         []string
	_glueTriggerUpdate                        string
	_glueType                                 string
	_glueUpdateEtlLibraries                   string
	_glueUpdateOpenTableFormatInput           string
	_glueUserBackgroundSessionsEnabled        string
	_glueVersionId                            string
	_glueVersionIds                           []string
	_glueVersions                             string
	_glueViewUpdateAction                     string
	_glueWorkerType                           string
	_glueWorkflowName                         string
	_glueXMLClassifier                        string
)

// Creates one or more partitions in a batch operation.
func glue_BatchCreatePartition(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchCreatePartitionInput{
		// DatabaseName: *string, // Required
		// PartitionInputList: []types.PartitionInput, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionInputList) > 0 {
		if err := assignInputField(input, "PartitionInputList", _gluePartitionInputList); err != nil {
			log.Errorf("invalid --partition-input-list: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.BatchCreatePartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a list of connection definitions from the Data Catalog.
func glue_BatchDeleteConnection(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchDeleteConnectionInput{
		// ConnectionNameList: []string, // Required
	}

	if len(_glueConnectionNameList) > 0 {
		input.ConnectionNameList = append([]string(nil), _glueConnectionNameList...)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.BatchDeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more partitions in a batch operation.
func glue_BatchDeletePartition(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchDeletePartitionInput{
		// DatabaseName: *string, // Required
		// PartitionsToDelete: []types.PartitionValueList, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionsToDelete) > 0 {
		if err := assignInputField(input, "PartitionsToDelete", _gluePartitionsToDelete); err != nil {
			log.Errorf("invalid --partitions-to-delete: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.BatchDeletePartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes multiple tables at once.
// After completing this operation, you no longer have access to the table
// versions and partitions that belong to the deleted table. Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion of the
// service.
//
// To ensure the immediate deletion of all related resources, before calling
// BatchDeleteTable , use DeleteTableVersion or BatchDeleteTableVersion , and
// DeletePartition or BatchDeletePartition , to delete any resources that belong to
// the table.
func glue_BatchDeleteTable(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchDeleteTableInput{
		// DatabaseName: *string, // Required
		// TablesToDelete: []string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTablesToDelete) > 0 {
		input.TablesToDelete = append([]string(nil), _glueTablesToDelete...)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}

	if resp, err := client.BatchDeleteTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified batch of versions of a table.
func glue_BatchDeleteTableVersion(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchDeleteTableVersionInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// VersionIds: []string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueVersionIds) > 0 {
		input.VersionIds = append([]string(nil), _glueVersionIds...)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.BatchDeleteTableVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a list of blueprints.
func glue_BatchGetBlueprints(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetBlueprintsInput{
		// Names: []string, // Required
	}

	if len(_glueNames) > 0 {
		input.Names = append([]string(nil), _glueNames...)
	}
	if len(_glueIncludeBlueprint) > 0 {
		if err := assignInputField(input, "IncludeBlueprint", _glueIncludeBlueprint); err != nil {
			log.Errorf("invalid --include-blueprint: %s", err.Error())
			return
		}
	}
	if len(_glueIncludeParameterSpec) > 0 {
		if err := assignInputField(input, "IncludeParameterSpec", _glueIncludeParameterSpec); err != nil {
			log.Errorf("invalid --include-parameter-spec: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetBlueprints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of resource metadata for a given list of crawler names. After
// calling the ListCrawlers operation, you can call this operation to access the
// data to which you have been granted permissions. This operation supports all IAM
// permissions, including permission conditions that uses tags.
func glue_BatchGetCrawlers(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetCrawlersInput{
		// CrawlerNames: []string, // Required
	}

	if len(_glueCrawlerNames) > 0 {
		input.CrawlerNames = append([]string(nil), _glueCrawlerNames...)
	}

	if resp, err := client.BatchGetCrawlers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details for the custom patterns specified by a list of names.
func glue_BatchGetCustomEntityTypes(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetCustomEntityTypesInput{
		// Names: []string, // Required
	}

	if len(_glueNames) > 0 {
		input.Names = append([]string(nil), _glueNames...)
	}

	if resp, err := client.BatchGetCustomEntityTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of data quality results for the specified result IDs.
func glue_BatchGetDataQualityResult(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetDataQualityResultInput{
		// ResultIds: []string, // Required
	}

	if len(_glueResultIds) > 0 {
		input.ResultIds = append([]string(nil), _glueResultIds...)
	}

	if resp, err := client.BatchGetDataQualityResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of resource metadata for a given list of development endpoint
// names. After calling the ListDevEndpoints operation, you can call this
// operation to access the data to which you have been granted permissions. This
// operation supports all IAM permissions, including permission conditions that
// uses tags.
func glue_BatchGetDevEndpoints(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetDevEndpointsInput{
		// DevEndpointNames: []string, // Required
	}

	if len(_glueDevEndpointNames) > 0 {
		input.DevEndpointNames = append([]string(nil), _glueDevEndpointNames...)
	}

	if resp, err := client.BatchGetDevEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of resource metadata for a given list of job names. After
// calling the ListJobs operation, you can call this operation to access the data
// to which you have been granted permissions. This operation supports all IAM
// permissions, including permission conditions that uses tags.
func glue_BatchGetJobs(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetJobsInput{
		// JobNames: []string, // Required
	}

	if len(_glueJobNames) > 0 {
		input.JobNames = append([]string(nil), _glueJobNames...)
	}

	if resp, err := client.BatchGetJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves partitions in a batch request.
func glue_BatchGetPartition(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetPartitionInput{
		// DatabaseName: *string, // Required
		// PartitionsToGet: []types.PartitionValueList, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionsToGet) > 0 {
		if err := assignInputField(input, "PartitionsToGet", _gluePartitionsToGet); err != nil {
			log.Errorf("invalid --partitions-to-get: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.BatchGetPartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the configuration for the specified table optimizers.
func glue_BatchGetTableOptimizer(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetTableOptimizerInput{
		// Entries: []types.BatchGetTableOptimizerEntry, // Required
	}

	if len(_glueEntries) > 0 {
		if err := assignInputField(input, "Entries", _glueEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetTableOptimizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of resource metadata for a given list of trigger names. After
// calling the ListTriggers operation, you can call this operation to access the
// data to which you have been granted permissions. This operation supports all IAM
// permissions, including permission conditions that uses tags.
func glue_BatchGetTriggers(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetTriggersInput{
		// TriggerNames: []string, // Required
	}

	if len(_glueTriggerNames) > 0 {
		input.TriggerNames = append([]string(nil), _glueTriggerNames...)
	}

	if resp, err := client.BatchGetTriggers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of resource metadata for a given list of workflow names. After
// calling the ListWorkflows operation, you can call this operation to access the
// data to which you have been granted permissions. This operation supports all IAM
// permissions, including permission conditions that uses tags.
func glue_BatchGetWorkflows(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchGetWorkflowsInput{
		// Names: []string, // Required
	}

	if len(_glueNames) > 0 {
		input.Names = append([]string(nil), _glueNames...)
	}
	if len(_glueIncludeGraph) > 0 {
		if err := assignInputField(input, "IncludeGraph", _glueIncludeGraph); err != nil {
			log.Errorf("invalid --include-graph: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetWorkflows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Annotate datapoints over time for a specific data quality statistic. The API
// requires both profileID and statisticID as part of the InclusionAnnotation
// input. The API only works for a single statisticId across multiple profiles.
func glue_BatchPutDataQualityStatisticAnnotation(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchPutDataQualityStatisticAnnotationInput{
		// InclusionAnnotations: []types.DatapointInclusionAnnotation, // Required
	}

	if len(_glueInclusionAnnotations) > 0 {
		if err := assignInputField(input, "InclusionAnnotations", _glueInclusionAnnotations); err != nil {
			log.Errorf("invalid --inclusion-annotations: %s", err.Error())
			return
		}
	}
	if len(_glueClientToken) > 0 {
		input.ClientToken = aws.String(_glueClientToken)
	}

	if resp, err := client.BatchPutDataQualityStatisticAnnotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops one or more job runs for a specified job definition.
func glue_BatchStopJobRun(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchStopJobRunInput{
		// JobName: *string, // Required
		// JobRunIds: []string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueJobRunIds) > 0 {
		input.JobRunIds = append([]string(nil), _glueJobRunIds...)
	}

	if resp, err := client.BatchStopJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates one or more partitions in a batch operation.
func glue_BatchUpdatePartition(cfg aws.Config, client *glue.Client) {
	input := &glue.BatchUpdatePartitionInput{
		// DatabaseName: *string, // Required
		// Entries: []types.BatchUpdatePartitionRequestEntry, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueEntries) > 0 {
		if err := assignInputField(input, "Entries", _glueEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.BatchUpdatePartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified recommendation run that was being used to generate rules.
func glue_CancelDataQualityRuleRecommendationRun(cfg aws.Config, client *glue.Client) {
	input := &glue.CancelDataQualityRuleRecommendationRunInput{
		// RunId: *string, // Required
	}

	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.CancelDataQualityRuleRecommendationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a run where a ruleset is being evaluated against a data source.
func glue_CancelDataQualityRulesetEvaluationRun(cfg aws.Config, client *glue.Client) {
	input := &glue.CancelDataQualityRulesetEvaluationRunInput{
		// RunId: *string, // Required
	}

	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.CancelDataQualityRulesetEvaluationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels (stops) a task run. Machine learning task runs are asynchronous tasks
// that Glue runs on your behalf as part of various machine learning workflows. You
// can cancel a machine learning task run at any time by calling CancelMLTaskRun
// with a task run's parent transform's TransformID and the task run's TaskRunId .
func glue_CancelMLTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.CancelMLTaskRunInput{
		// TaskRunId: *string, // Required
		// TransformId: *string, // Required
	}

	if len(_glueTaskRunId) > 0 {
		input.TaskRunId = aws.String(_glueTaskRunId)
	}
	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.CancelMLTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the statement.
func glue_CancelStatement(cfg aws.Config, client *glue.Client) {
	input := &glue.CancelStatementInput{
		// Id: int32, // Required
		// SessionId: *string, // Required
	}

	if len(_glueId) > 0 {
		if err := assignInputField(input, "Id", _glueId); err != nil {
			log.Errorf("invalid --id: %s", err.Error())
			return
		}
	}
	if len(_glueSessionId) > 0 {
		input.SessionId = aws.String(_glueSessionId)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.CancelStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates the supplied schema. This call has no side effects, it simply
// validates using the supplied schema using DataFormat as the format. Since it
// does not take a schema set name, no compatibility checks are performed.
func glue_CheckSchemaVersionValidity(cfg aws.Config, client *glue.Client) {
	input := &glue.CheckSchemaVersionValidityInput{
		// DataFormat: types.DataFormat, // Required
		// SchemaDefinition: *string, // Required
	}

	if len(_glueDataFormat) > 0 {
		if err := assignInputField(input, "DataFormat", _glueDataFormat); err != nil {
			log.Errorf("invalid --data-format: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaDefinition) > 0 {
		input.SchemaDefinition = aws.String(_glueSchemaDefinition)
	}

	if resp, err := client.CheckSchemaVersionValidity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a blueprint with Glue.
func glue_CreateBlueprint(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateBlueprintInput{
		// BlueprintLocation: *string, // Required
		// Name: *string, // Required
	}

	if len(_glueBlueprintLocation) > 0 {
		input.BlueprintLocation = aws.String(_glueBlueprintLocation)
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new catalog in the Glue Data Catalog.
func glue_CreateCatalog(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateCatalogInput{
		// CatalogInput: *types.CatalogInput, // Required
		// Name: *string, // Required
	}

	if len(_glueCatalogInput) > 0 {
		if err := assignInputField(input, "CatalogInput", _glueCatalogInput); err != nil {
			log.Errorf("invalid --catalog-input: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a classifier in the user's account. This can be a GrokClassifier , an
// XMLClassifier , a JsonClassifier , or a CsvClassifier , depending on which field
// of the request is present.
func glue_CreateClassifier(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateClassifierInput{}

	if len(_glueCsvClassifier) > 0 {
		if err := assignInputField(input, "CsvClassifier", _glueCsvClassifier); err != nil {
			log.Errorf("invalid --csv-classifier: %s", err.Error())
			return
		}
	}
	if len(_glueGrokClassifier) > 0 {
		if err := assignInputField(input, "GrokClassifier", _glueGrokClassifier); err != nil {
			log.Errorf("invalid --grok-classifier: %s", err.Error())
			return
		}
	}
	if len(_glueJsonClassifier) > 0 {
		if err := assignInputField(input, "JsonClassifier", _glueJsonClassifier); err != nil {
			log.Errorf("invalid --json-classifier: %s", err.Error())
			return
		}
	}
	if len(_glueXMLClassifier) > 0 {
		if err := assignInputField(input, "XMLClassifier", _glueXMLClassifier); err != nil {
			log.Errorf("invalid --xml-classifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates settings for a column statistics task.
func glue_CreateColumnStatisticsTaskSettings(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateColumnStatisticsTaskSettingsInput{
		// DatabaseName: *string, // Required
		// Role: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogID = aws.String(_glueCatalogId)
	}
	if len(_glueColumnNameList) > 0 {
		input.ColumnNameList = append([]string(nil), _glueColumnNameList...)
	}
	if len(_glueSampleSize) > 0 {
		if err := assignInputField(input, "SampleSize", _glueSampleSize); err != nil {
			log.Errorf("invalid --sample-size: %s", err.Error())
			return
		}
	}
	if len(_glueSchedule) > 0 {
		input.Schedule = aws.String(_glueSchedule)
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateColumnStatisticsTaskSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection definition in the Data Catalog.
// Connections used for creating federated resources require the IAM
// glue:PassConnection permission.
func glue_CreateConnection(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateConnectionInput{
		// ConnectionInput: *types.ConnectionInput, // Required
	}

	if len(_glueConnectionInput) > 0 {
		if err := assignInputField(input, "ConnectionInput", _glueConnectionInput); err != nil {
			log.Errorf("invalid --connection-input: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new crawler with specified targets, role, configuration, and optional
// schedule. At least one crawl target must be specified, in the s3Targets field,
// the jdbcTargets field, or the DynamoDBTargets field.
func glue_CreateCrawler(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateCrawlerInput{
		// Name: *string, // Required
		// Role: *string, // Required
		// Targets: *types.CrawlerTargets, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueTargets) > 0 {
		if err := assignInputField(input, "Targets", _glueTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_glueClassifiers) > 0 {
		input.Classifiers = append([]string(nil), _glueClassifiers...)
	}
	if len(_glueConfiguration) > 0 {
		input.Configuration = aws.String(_glueConfiguration)
	}
	if len(_glueCrawlerSecurityConfiguration) > 0 {
		input.CrawlerSecurityConfiguration = aws.String(_glueCrawlerSecurityConfiguration)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueLakeFormationConfiguration) > 0 {
		if err := assignInputField(input, "LakeFormationConfiguration", _glueLakeFormationConfiguration); err != nil {
			log.Errorf("invalid --lake-formation-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueLineageConfiguration) > 0 {
		if err := assignInputField(input, "LineageConfiguration", _glueLineageConfiguration); err != nil {
			log.Errorf("invalid --lineage-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueRecrawlPolicy) > 0 {
		if err := assignInputField(input, "RecrawlPolicy", _glueRecrawlPolicy); err != nil {
			log.Errorf("invalid --recrawl-policy: %s", err.Error())
			return
		}
	}
	if len(_glueSchedule) > 0 {
		input.Schedule = aws.String(_glueSchedule)
	}
	if len(_glueSchemaChangePolicy) > 0 {
		if err := assignInputField(input, "SchemaChangePolicy", _glueSchemaChangePolicy); err != nil {
			log.Errorf("invalid --schema-change-policy: %s", err.Error())
			return
		}
	}
	if len(_glueTablePrefix) > 0 {
		input.TablePrefix = aws.String(_glueTablePrefix)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCrawler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom pattern that is used to detect sensitive data across the
// columns and rows of your structured data.
//
// Each custom pattern you create specifies a regular expression and an optional
// list of context words. If no context words are passed only a regular expression
// is checked.
func glue_CreateCustomEntityType(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateCustomEntityTypeInput{
		// Name: *string, // Required
		// RegexString: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRegexString) > 0 {
		input.RegexString = aws.String(_glueRegexString)
	}
	if len(_glueContextWords) > 0 {
		input.ContextWords = append([]string(nil), _glueContextWords...)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomEntityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data quality ruleset with DQDL rules applied to a specified Glue
// table.
//
// You create the ruleset using the Data Quality Definition Language (DQDL). For
// more information, see the Glue developer guide.
func glue_CreateDataQualityRuleset(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateDataQualityRulesetInput{
		// Name: *string, // Required
		// Ruleset: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRuleset) > 0 {
		input.Ruleset = aws.String(_glueRuleset)
	}
	if len(_glueClientToken) > 0 {
		input.ClientToken = aws.String(_glueClientToken)
	}
	if len(_glueDataQualitySecurityConfiguration) > 0 {
		input.DataQualitySecurityConfiguration = aws.String(_glueDataQualitySecurityConfiguration)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueTargetTable) > 0 {
		if err := assignInputField(input, "TargetTable", _glueTargetTable); err != nil {
			log.Errorf("invalid --target-table: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataQualityRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new database in a Data Catalog.
func glue_CreateDatabase(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateDatabaseInput{
		// DatabaseInput: *types.DatabaseInput, // Required
	}

	if len(_glueDatabaseInput) > 0 {
		if err := assignInputField(input, "DatabaseInput", _glueDatabaseInput); err != nil {
			log.Errorf("invalid --database-input: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new development endpoint.
func glue_CreateDevEndpoint(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateDevEndpointInput{
		// EndpointName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_glueEndpointName) > 0 {
		input.EndpointName = aws.String(_glueEndpointName)
	}
	if len(_glueRoleArn) > 0 {
		input.RoleArn = aws.String(_glueRoleArn)
	}
	if len(_glueArguments) > 0 {
		if err := assignInputField(input, "Arguments", _glueArguments); err != nil {
			log.Errorf("invalid --arguments: %s", err.Error())
			return
		}
	}
	if len(_glueExtraJarsS3Path) > 0 {
		input.ExtraJarsS3Path = aws.String(_glueExtraJarsS3Path)
	}
	if len(_glueExtraPythonLibsS3Path) > 0 {
		input.ExtraPythonLibsS3Path = aws.String(_glueExtraPythonLibsS3Path)
	}
	if len(_glueGlueVersion) > 0 {
		input.GlueVersion = aws.String(_glueGlueVersion)
	}
	if len(_glueNumberOfNodes) > 0 {
		if err := assignInputField(input, "NumberOfNodes", _glueNumberOfNodes); err != nil {
			log.Errorf("invalid --number-of-nodes: %s", err.Error())
			return
		}
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_gluePublicKey) > 0 {
		input.PublicKey = aws.String(_gluePublicKey)
	}
	if len(_gluePublicKeys) > 0 {
		input.PublicKeys = append([]string(nil), _gluePublicKeys...)
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}
	if len(_glueSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _glueSecurityGroupIds...)
	}
	if len(_glueSubnetId) > 0 {
		input.SubnetId = aws.String(_glueSubnetId)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueWorkerType) > 0 {
		if err := assignInputField(input, "WorkerType", _glueWorkerType); err != nil {
			log.Errorf("invalid --worker-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDevEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Glue Identity Center configuration to enable integration between
// Glue and Amazon Web Services IAM Identity Center for authentication and
// authorization.
func glue_CreateGlueIdentityCenterConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateGlueIdentityCenterConfigurationInput{
		// InstanceArn: *string, // Required
	}

	if len(_glueInstanceArn) > 0 {
		input.InstanceArn = aws.String(_glueInstanceArn)
	}
	if len(_glueScopes) > 0 {
		input.Scopes = append([]string(nil), _glueScopes...)
	}
	if len(_glueUserBackgroundSessionsEnabled) > 0 {
		if err := assignInputField(input, "UserBackgroundSessionsEnabled", _glueUserBackgroundSessionsEnabled); err != nil {
			log.Errorf("invalid --user-background-sessions-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlueIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Zero-ETL integration in the caller's account between two resources
// with Amazon Resource Names (ARNs): the SourceArn and TargetArn .
func glue_CreateIntegration(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateIntegrationInput{
		// IntegrationName: *string, // Required
		// SourceArn: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_glueIntegrationName) > 0 {
		input.IntegrationName = aws.String(_glueIntegrationName)
	}
	if len(_glueSourceArn) > 0 {
		input.SourceArn = aws.String(_glueSourceArn)
	}
	if len(_glueTargetArn) > 0 {
		input.TargetArn = aws.String(_glueTargetArn)
	}
	if len(_glueAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _glueAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_glueDataFilter) > 0 {
		input.DataFilter = aws.String(_glueDataFilter)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueIntegrationConfig) > 0 {
		if err := assignInputField(input, "IntegrationConfig", _glueIntegrationConfig); err != nil {
			log.Errorf("invalid --integration-config: %s", err.Error())
			return
		}
	}
	if len(_glueKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_glueKmsKeyId)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
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

// This API can be used for setting up the ResourceProperty of the Glue connection
// (for the source) or Glue database ARN (for the target). These properties can
// include the role to access the connection or database. To set both source and
// target properties the same API needs to be invoked with the Glue connection ARN
// as ResourceArn with SourceProcessingProperties and the Glue database ARN as
// ResourceArn with TargetProcessingProperties respectively.
func glue_CreateIntegrationResourceProperty(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateIntegrationResourcePropertyInput{
		// ResourceArn: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueSourceProcessingProperties) > 0 {
		if err := assignInputField(input, "SourceProcessingProperties", _glueSourceProcessingProperties); err != nil {
			log.Errorf("invalid --source-processing-properties: %s", err.Error())
			return
		}
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueTargetProcessingProperties) > 0 {
		if err := assignInputField(input, "TargetProcessingProperties", _glueTargetProcessingProperties); err != nil {
			log.Errorf("invalid --target-processing-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegrationResourceProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is used to provide optional override properties for the the tables
// that need to be replicated. These properties can include properties for
// filtering and partitioning for the source and target tables. To set both source
// and target properties the same API need to be invoked with the Glue connection
// ARN as ResourceArn with SourceTableConfig , and the Glue database ARN as
// ResourceArn with TargetTableConfig respectively.
func glue_CreateIntegrationTableProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateIntegrationTablePropertiesInput{
		// ResourceArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueSourceTableConfig) > 0 {
		if err := assignInputField(input, "SourceTableConfig", _glueSourceTableConfig); err != nil {
			log.Errorf("invalid --source-table-config: %s", err.Error())
			return
		}
	}
	if len(_glueTargetTableConfig) > 0 {
		if err := assignInputField(input, "TargetTableConfig", _glueTargetTableConfig); err != nil {
			log.Errorf("invalid --target-table-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegrationTableProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new job definition.
func glue_CreateJob(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateJobInput{
		// Command: *types.JobCommand, // Required
		// Name: *string, // Required
		// Role: *string, // Required
	}

	if len(_glueCommand) > 0 {
		if err := assignInputField(input, "Command", _glueCommand); err != nil {
			log.Errorf("invalid --command: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueAllocatedCapacity) > 0 {
		if err := assignInputField(input, "AllocatedCapacity", _glueAllocatedCapacity); err != nil {
			log.Errorf("invalid --allocated-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueCodeGenConfigurationNodes) > 0 {
		if err := assignInputField(input, "CodeGenConfigurationNodes", _glueCodeGenConfigurationNodes); err != nil {
			log.Errorf("invalid --code-gen-configuration-nodes: %s", err.Error())
			return
		}
	}
	if len(_glueConnections) > 0 {
		if err := assignInputField(input, "Connections", _glueConnections); err != nil {
			log.Errorf("invalid --connections: %s", err.Error())
			return
		}
	}
	if len(_glueDefaultArguments) > 0 {
		if err := assignInputField(input, "DefaultArguments", _glueDefaultArguments); err != nil {
			log.Errorf("invalid --default-arguments: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueExecutionClass) > 0 {
		if err := assignInputField(input, "ExecutionClass", _glueExecutionClass); err != nil {
			log.Errorf("invalid --execution-class: %s", err.Error())
			return
		}
	}
	if len(_glueExecutionProperty) > 0 {
		if err := assignInputField(input, "ExecutionProperty", _glueExecutionProperty); err != nil {
			log.Errorf("invalid --execution-property: %s", err.Error())
			return
		}
	}
	if len(_glueGlueVersion) > 0 {
		input.GlueVersion = aws.String(_glueGlueVersion)
	}
	if len(_glueJobMode) > 0 {
		if err := assignInputField(input, "JobMode", _glueJobMode); err != nil {
			log.Errorf("invalid --job-mode: %s", err.Error())
			return
		}
	}
	if len(_glueJobRunQueuingEnabled) > 0 {
		if err := assignInputField(input, "JobRunQueuingEnabled", _glueJobRunQueuingEnabled); err != nil {
			log.Errorf("invalid --job-run-queuing-enabled: %s", err.Error())
			return
		}
	}
	if len(_glueLogUri) > 0 {
		input.LogUri = aws.String(_glueLogUri)
	}
	if len(_glueMaintenanceWindow) > 0 {
		input.MaintenanceWindow = aws.String(_glueMaintenanceWindow)
	}
	if len(_glueMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _glueMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _glueMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_glueNonOverridableArguments) > 0 {
		if err := assignInputField(input, "NonOverridableArguments", _glueNonOverridableArguments); err != nil {
			log.Errorf("invalid --non-overridable-arguments: %s", err.Error())
			return
		}
	}
	if len(_glueNotificationProperty) > 0 {
		if err := assignInputField(input, "NotificationProperty", _glueNotificationProperty); err != nil {
			log.Errorf("invalid --notification-property: %s", err.Error())
			return
		}
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}
	if len(_glueSourceControlDetails) > 0 {
		if err := assignInputField(input, "SourceControlDetails", _glueSourceControlDetails); err != nil {
			log.Errorf("invalid --source-control-details: %s", err.Error())
			return
		}
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_glueWorkerType) > 0 {
		if err := assignInputField(input, "WorkerType", _glueWorkerType); err != nil {
			log.Errorf("invalid --worker-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Glue machine learning transform. This operation creates the
// transform and all the necessary parameters to train it.
//
// Call this operation as the first step in the process of using a machine
// learning transform (such as the FindMatches transform) for deduplicating data.
// You can provide an optional Description , in addition to the parameters that you
// want to use for your algorithm.
//
// You must also specify certain parameters for the tasks that Glue runs on your
// behalf as part of learning from your data and creating a high-quality machine
// learning transform. These parameters include Role , and optionally,
// AllocatedCapacity , Timeout , and MaxRetries . For more information, see [Jobs].
//
// [Jobs]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html
func glue_CreateMLTransform(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateMLTransformInput{
		// InputRecordTables: []types.GlueTable, // Required
		// Name: *string, // Required
		// Parameters: *types.TransformParameters, // Required
		// Role: *string, // Required
	}

	if len(_glueInputRecordTables) > 0 {
		if err := assignInputField(input, "InputRecordTables", _glueInputRecordTables); err != nil {
			log.Errorf("invalid --input-record-tables: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueParameters) > 0 {
		if err := assignInputField(input, "Parameters", _glueParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueGlueVersion) > 0 {
		input.GlueVersion = aws.String(_glueGlueVersion)
	}
	if len(_glueMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _glueMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _glueMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_glueTransformEncryption) > 0 {
		if err := assignInputField(input, "TransformEncryption", _glueTransformEncryption); err != nil {
			log.Errorf("invalid --transform-encryption: %s", err.Error())
			return
		}
	}
	if len(_glueWorkerType) > 0 {
		if err := assignInputField(input, "WorkerType", _glueWorkerType); err != nil {
			log.Errorf("invalid --worker-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMLTransform(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new partition.
func glue_CreatePartition(cfg aws.Config, client *glue.Client) {
	input := &glue.CreatePartitionInput{
		// DatabaseName: *string, // Required
		// PartitionInput: *types.PartitionInput, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionInput) > 0 {
		if err := assignInputField(input, "PartitionInput", _gluePartitionInput); err != nil {
			log.Errorf("invalid --partition-input: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.CreatePartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a specified partition index in an existing table.
func glue_CreatePartitionIndex(cfg aws.Config, client *glue.Client) {
	input := &glue.CreatePartitionIndexInput{
		// DatabaseName: *string, // Required
		// PartitionIndex: *types.PartitionIndex, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionIndex) > 0 {
		if err := assignInputField(input, "PartitionIndex", _gluePartitionIndex); err != nil {
			log.Errorf("invalid --partition-index: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.CreatePartitionIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new registry which may be used to hold a collection of schemas.
func glue_CreateRegistry(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateRegistryInput{
		// RegistryName: *string, // Required
	}

	if len(_glueRegistryName) > 0 {
		input.RegistryName = aws.String(_glueRegistryName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new schema set and registers the schema definition. Returns an error
// if the schema set already exists without actually registering the version.
//
// When the schema set is created, a version checkpoint will be set to the first
// version. Compatibility mode "DISABLED" restricts any additional schema versions
// from being added after the first schema version. For all other compatibility
// modes, validation of compatibility settings will be applied only from the second
// version onwards when the RegisterSchemaVersion API is used.
//
// When this API is called without a RegistryId , this will create an entry for a
// "default-registry" in the registry database tables, if it is not already
// present.
func glue_CreateSchema(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateSchemaInput{
		// DataFormat: types.DataFormat, // Required
		// SchemaName: *string, // Required
	}

	if len(_glueDataFormat) > 0 {
		if err := assignInputField(input, "DataFormat", _glueDataFormat); err != nil {
			log.Errorf("invalid --data-format: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaName) > 0 {
		input.SchemaName = aws.String(_glueSchemaName)
	}
	if len(_glueCompatibility) > 0 {
		if err := assignInputField(input, "Compatibility", _glueCompatibility); err != nil {
			log.Errorf("invalid --compatibility: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueRegistryId) > 0 {
		if err := assignInputField(input, "RegistryId", _glueRegistryId); err != nil {
			log.Errorf("invalid --registry-id: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaDefinition) > 0 {
		input.SchemaDefinition = aws.String(_glueSchemaDefinition)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transforms a directed acyclic graph (DAG) into code.
func glue_CreateScript(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateScriptInput{}

	if len(_glueDagEdges) > 0 {
		if err := assignInputField(input, "DagEdges", _glueDagEdges); err != nil {
			log.Errorf("invalid --dag-edges: %s", err.Error())
			return
		}
	}
	if len(_glueDagNodes) > 0 {
		if err := assignInputField(input, "DagNodes", _glueDagNodes); err != nil {
			log.Errorf("invalid --dag-nodes: %s", err.Error())
			return
		}
	}
	if len(_glueLanguage) > 0 {
		if err := assignInputField(input, "Language", _glueLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScript(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new security configuration. A security configuration is a set of
// security properties that can be used by Glue. You can use a security
// configuration to encrypt data at rest. For information about using security
// configurations in Glue, see [Encrypting Data Written by Crawlers, Jobs, and Development Endpoints].
//
// [Encrypting Data Written by Crawlers, Jobs, and Development Endpoints]: https://docs.aws.amazon.com/glue/latest/dg/encryption-security-configuration.html
func glue_CreateSecurityConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateSecurityConfigurationInput{
		// EncryptionConfiguration: *types.EncryptionConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_glueEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _glueEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.CreateSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new session.
func glue_CreateSession(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateSessionInput{
		// Command: *types.SessionCommand, // Required
		// Id: *string, // Required
		// Role: *string, // Required
	}

	if len(_glueCommand) > 0 {
		if err := assignInputField(input, "Command", _glueCommand); err != nil {
			log.Errorf("invalid --command: %s", err.Error())
			return
		}
	}
	if len(_glueId) > 0 {
		input.Id = aws.String(_glueId)
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueConnections) > 0 {
		if err := assignInputField(input, "Connections", _glueConnections); err != nil {
			log.Errorf("invalid --connections: %s", err.Error())
			return
		}
	}
	if len(_glueDefaultArguments) > 0 {
		if err := assignInputField(input, "DefaultArguments", _glueDefaultArguments); err != nil {
			log.Errorf("invalid --default-arguments: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueGlueVersion) > 0 {
		input.GlueVersion = aws.String(_glueGlueVersion)
	}
	if len(_glueIdleTimeout) > 0 {
		if err := assignInputField(input, "IdleTimeout", _glueIdleTimeout); err != nil {
			log.Errorf("invalid --idle-timeout: %s", err.Error())
			return
		}
	}
	if len(_glueMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _glueMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_glueWorkerType) > 0 {
		if err := assignInputField(input, "WorkerType", _glueWorkerType); err != nil {
			log.Errorf("invalid --worker-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new table definition in the Data Catalog.
func glue_CreateTable(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateTableInput{
		// DatabaseName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueOpenTableFormatInput) > 0 {
		if err := assignInputField(input, "OpenTableFormatInput", _glueOpenTableFormatInput); err != nil {
			log.Errorf("invalid --open-table-format-input: %s", err.Error())
			return
		}
	}
	if len(_gluePartitionIndexes) > 0 {
		if err := assignInputField(input, "PartitionIndexes", _gluePartitionIndexes); err != nil {
			log.Errorf("invalid --partition-indexes: %s", err.Error())
			return
		}
	}
	if len(_glueTableInput) > 0 {
		if err := assignInputField(input, "TableInput", _glueTableInput); err != nil {
			log.Errorf("invalid --table-input: %s", err.Error())
			return
		}
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}

	if resp, err := client.CreateTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new table optimizer for a specific function.
func glue_CreateTableOptimizer(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateTableOptimizerInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// TableOptimizerConfiguration: *types.TableOptimizerConfiguration, // Required
		// Type: types.TableOptimizerType, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueTableOptimizerConfiguration) > 0 {
		if err := assignInputField(input, "TableOptimizerConfiguration", _glueTableOptimizerConfiguration); err != nil {
			log.Errorf("invalid --table-optimizer-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueType) > 0 {
		if err := assignInputField(input, "Type", _glueType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTableOptimizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new trigger.
// Job arguments may be logged. Do not pass plaintext secrets as arguments.
// Retrieve secrets from a Glue Connection, Amazon Web Services Secrets Manager or
// other secret management mechanism if you intend to keep them within the Job.
func glue_CreateTrigger(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateTriggerInput{
		// Actions: []types.Action, // Required
		// Name: *string, // Required
		// Type: types.TriggerType, // Required
	}

	if len(_glueActions) > 0 {
		if err := assignInputField(input, "Actions", _glueActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueType) > 0 {
		if err := assignInputField(input, "Type", _glueType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueEventBatchingCondition) > 0 {
		if err := assignInputField(input, "EventBatchingCondition", _glueEventBatchingCondition); err != nil {
			log.Errorf("invalid --event-batching-condition: %s", err.Error())
			return
		}
	}
	if len(_gluePredicate) > 0 {
		if err := assignInputField(input, "Predicate", _gluePredicate); err != nil {
			log.Errorf("invalid --predicate: %s", err.Error())
			return
		}
	}
	if len(_glueSchedule) > 0 {
		input.Schedule = aws.String(_glueSchedule)
	}
	if len(_glueStartOnCreation) > 0 {
		if err := assignInputField(input, "StartOnCreation", _glueStartOnCreation); err != nil {
			log.Errorf("invalid --start-on-creation: %s", err.Error())
			return
		}
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_glueWorkflowName) > 0 {
		input.WorkflowName = aws.String(_glueWorkflowName)
	}

	if resp, err := client.CreateTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Glue usage profile.
func glue_CreateUsageProfile(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateUsageProfileInput{
		// Configuration: *types.ProfileConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_glueConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _glueConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUsageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new function definition in the Data Catalog.
func glue_CreateUserDefinedFunction(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateUserDefinedFunctionInput{
		// DatabaseName: *string, // Required
		// FunctionInput: *types.UserDefinedFunctionInput, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueFunctionInput) > 0 {
		if err := assignInputField(input, "FunctionInput", _glueFunctionInput); err != nil {
			log.Errorf("invalid --function-input: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.CreateUserDefinedFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new workflow.
func glue_CreateWorkflow(cfg aws.Config, client *glue.Client) {
	input := &glue.CreateWorkflowInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDefaultRunProperties) > 0 {
		if err := assignInputField(input, "DefaultRunProperties", _glueDefaultRunProperties); err != nil {
			log.Errorf("invalid --default-run-properties: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueMaxConcurrentRuns) > 0 {
		if err := assignInputField(input, "MaxConcurrentRuns", _glueMaxConcurrentRuns); err != nil {
			log.Errorf("invalid --max-concurrent-runs: %s", err.Error())
			return
		}
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing blueprint.
func glue_DeleteBlueprint(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteBlueprintInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified catalog from the Glue Data Catalog.
// After completing this operation, you no longer have access to the databases,
// tables (and all table versions and partitions that might belong to the tables)
// and the user-defined functions in the deleted catalog. Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion of the
// service.
//
// To ensure the immediate deletion of all related resources before calling the
// DeleteCatalog operation, use DeleteTableVersion (or BatchDeleteTableVersion ),
// DeletePartition (or BatchDeletePartition ), DeleteTable (or BatchDeleteTable ),
// DeleteUserDefinedFunction and DeleteDatabase to delete any resources that
// belong to the catalog.
func glue_DeleteCatalog(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteCatalogInput{
		// CatalogId: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a classifier from the Data Catalog.
func glue_DeleteClassifier(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteClassifierInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the partition column statistics of a column.
// The Identity and Access Management (IAM) permission required for this operation
// is DeletePartition .
func glue_DeleteColumnStatisticsForPartition(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteColumnStatisticsForPartitionInput{
		// ColumnName: *string, // Required
		// DatabaseName: *string, // Required
		// PartitionValues: []string, // Required
		// TableName: *string, // Required
	}

	if len(_glueColumnName) > 0 {
		input.ColumnName = aws.String(_glueColumnName)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionValues) > 0 {
		input.PartitionValues = append([]string(nil), _gluePartitionValues...)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteColumnStatisticsForPartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves table statistics of columns.
// The Identity and Access Management (IAM) permission required for this operation
// is DeleteTable .
func glue_DeleteColumnStatisticsForTable(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteColumnStatisticsForTableInput{
		// ColumnName: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueColumnName) > 0 {
		input.ColumnName = aws.String(_glueColumnName)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteColumnStatisticsForTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes settings for a column statistics task.
func glue_DeleteColumnStatisticsTaskSettings(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteColumnStatisticsTaskSettingsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.DeleteColumnStatisticsTaskSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connection from the Data Catalog.
func glue_DeleteConnection(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteConnectionInput{
		// ConnectionName: *string, // Required
	}

	if len(_glueConnectionName) > 0 {
		input.ConnectionName = aws.String(_glueConnectionName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom connection type in Glue.
// The connection type must exist and be registered before it can be deleted. This
// operation supports cleanup of connection type resources and helps maintain
// proper lifecycle management of custom connection types.
func glue_DeleteConnectionType(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteConnectionTypeInput{
		// ConnectionType: *string, // Required
	}

	if len(_glueConnectionType) > 0 {
		input.ConnectionType = aws.String(_glueConnectionType)
	}

	if resp, err := client.DeleteConnectionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a specified crawler from the Glue Data Catalog, unless the crawler
// state is RUNNING .
func glue_DeleteCrawler(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteCrawlerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteCrawler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom pattern by specifying its name.
func glue_DeleteCustomEntityType(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteCustomEntityTypeInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteCustomEntityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data quality ruleset.
func glue_DeleteDataQualityRuleset(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteDataQualityRulesetInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteDataQualityRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a specified database from a Data Catalog.
// After completing this operation, you no longer have access to the tables (and
// all table versions and partitions that might belong to the tables) and the
// user-defined functions in the deleted database. Glue deletes these "orphaned"
// resources asynchronously in a timely manner, at the discretion of the service.
//
// To ensure the immediate deletion of all related resources, before calling
// DeleteDatabase , use DeleteTableVersion or BatchDeleteTableVersion ,
// DeletePartition or BatchDeletePartition , DeleteUserDefinedFunction , and
// DeleteTable or BatchDeleteTable , to delete any resources that belong to the
// database.
func glue_DeleteDatabase(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteDatabaseInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified development endpoint.
func glue_DeleteDevEndpoint(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteDevEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_glueEndpointName) > 0 {
		input.EndpointName = aws.String(_glueEndpointName)
	}

	if resp, err := client.DeleteDevEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the existing Glue Identity Center configuration, removing the
// integration between Glue and Amazon Web Services IAM Identity Center.
func glue_DeleteGlueIdentityCenterConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteGlueIdentityCenterConfigurationInput{}

	if resp, err := client.DeleteGlueIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Zero-ETL integration.
func glue_DeleteIntegration(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteIntegrationInput{
		// IntegrationIdentifier: *string, // Required
	}

	if len(_glueIntegrationIdentifier) > 0 {
		input.IntegrationIdentifier = aws.String(_glueIntegrationIdentifier)
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is used for deleting the ResourceProperty of the Glue connection (for
// the source) or Glue database ARN (for the target).
func glue_DeleteIntegrationResourceProperty(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteIntegrationResourcePropertyInput{
		// ResourceArn: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}

	if resp, err := client.DeleteIntegrationResourceProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the table properties that have been created for the tables that need to
// be replicated.
func glue_DeleteIntegrationTableProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteIntegrationTablePropertiesInput{
		// ResourceArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.DeleteIntegrationTableProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified job definition. If the job definition is not found, no
// exception is thrown.
func glue_DeleteJob(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteJobInput{
		// JobName: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}

	if resp, err := client.DeleteJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Glue machine learning transform. Machine learning transforms are a
// special type of transform that use machine learning to learn the details of the
// transformation to be performed by learning from examples provided by humans.
// These transformations are then saved by Glue. If you no longer need a transform,
// you can delete it by calling DeleteMLTransforms . However, any Glue jobs that
// still reference the deleted transform will no longer succeed.
func glue_DeleteMLTransform(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteMLTransformInput{
		// TransformId: *string, // Required
	}

	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.DeleteMLTransform(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified partition.
func glue_DeletePartition(cfg aws.Config, client *glue.Client) {
	input := &glue.DeletePartitionInput{
		// DatabaseName: *string, // Required
		// PartitionValues: []string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionValues) > 0 {
		input.PartitionValues = append([]string(nil), _gluePartitionValues...)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeletePartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified partition index from an existing table.
func glue_DeletePartitionIndex(cfg aws.Config, client *glue.Client) {
	input := &glue.DeletePartitionIndexInput{
		// DatabaseName: *string, // Required
		// IndexName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueIndexName) > 0 {
		input.IndexName = aws.String(_glueIndexName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeletePartitionIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the entire registry including schema and all of its versions. To get the
// status of the delete operation, you can call the GetRegistry API after the
// asynchronous call. Deleting a registry will deactivate all online operations for
// the registry such as the UpdateRegistry , CreateSchema , UpdateSchema , and
// RegisterSchemaVersion APIs.
func glue_DeleteRegistry(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteRegistryInput{
		// RegistryId: *types.RegistryId, // Required
	}

	if len(_glueRegistryId) > 0 {
		if err := assignInputField(input, "RegistryId", _glueRegistryId); err != nil {
			log.Errorf("invalid --registry-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified policy.
func glue_DeleteResourcePolicy(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteResourcePolicyInput{}

	if len(_gluePolicyHashCondition) > 0 {
		input.PolicyHashCondition = aws.String(_gluePolicyHashCondition)
	}
	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the entire schema set, including the schema set and all of its
// versions. To get the status of the delete operation, you can call GetSchema API
// after the asynchronous call. Deleting a registry will deactivate all online
// operations for the schema, such as the GetSchemaByDefinition , and
// RegisterSchemaVersion APIs.
func glue_DeleteSchema(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteSchemaInput{
		// SchemaId: *types.SchemaId, // Required
	}

	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove versions from the specified schema. A version number or range may be
// supplied. If the compatibility mode forbids deleting of a version that is
// necessary, such as BACKWARDS_FULL, an error is returned. Calling the
// GetSchemaVersions API after this call will list the status of the deleted
// versions.
//
// When the range of version numbers contain check pointed version, the API will
// return a 409 conflict and will not proceed with the deletion. You have to remove
// the checkpoint first using the DeleteSchemaCheckpoint API before using this API.
//
// You cannot use the DeleteSchemaVersions API to delete the first schema version
// in the schema set. The first schema version can only be deleted by the
// DeleteSchema API. This operation will also delete the attached
// SchemaVersionMetadata under the schema versions. Hard deletes will be enforced
// on the database.
//
// If the compatibility mode forbids deleting of a version that is necessary, such
// as BACKWARDS_FULL, an error is returned.
func glue_DeleteSchemaVersions(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteSchemaVersionsInput{
		// SchemaId: *types.SchemaId, // Required
		// Versions: *string, // Required
	}

	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueVersions) > 0 {
		input.Versions = aws.String(_glueVersions)
	}

	if resp, err := client.DeleteSchemaVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified security configuration.
func glue_DeleteSecurityConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteSecurityConfigurationInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the session.
func glue_DeleteSession(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteSessionInput{
		// Id: *string, // Required
	}

	if len(_glueId) > 0 {
		input.Id = aws.String(_glueId)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.DeleteSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a table definition from the Data Catalog.
// After completing this operation, you no longer have access to the table
// versions and partitions that belong to the deleted table. Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion of the
// service.
//
// To ensure the immediate deletion of all related resources, before calling
// DeleteTable , use DeleteTableVersion or BatchDeleteTableVersion , and
// DeletePartition or BatchDeletePartition , to delete any resources that belong to
// the table.
func glue_DeleteTable(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteTableInput{
		// DatabaseName: *string, // Required
		// Name: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}

	if resp, err := client.DeleteTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an optimizer and all associated metadata for a table. The optimization
// will no longer be performed on the table.
func glue_DeleteTableOptimizer(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteTableOptimizerInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// Type: types.TableOptimizerType, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueType) > 0 {
		if err := assignInputField(input, "Type", _glueType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTableOptimizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified version of a table.
func glue_DeleteTableVersion(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteTableVersionInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueVersionId) > 0 {
		input.VersionId = aws.String(_glueVersionId)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteTableVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified trigger. If the trigger is not found, no exception is
// thrown.
func glue_DeleteTrigger(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteTriggerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Glue specified usage profile.
func glue_DeleteUsageProfile(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteUsageProfileInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteUsageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing function definition from the Data Catalog.
func glue_DeleteUserDefinedFunction(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteUserDefinedFunctionInput{
		// DatabaseName: *string, // Required
		// FunctionName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueFunctionName) > 0 {
		input.FunctionName = aws.String(_glueFunctionName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.DeleteUserDefinedFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workflow.
func glue_DeleteWorkflow(cfg aws.Config, client *glue.Client) {
	input := &glue.DeleteWorkflowInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DescribeConnectionType API provides full details of the supported options
// for a given connection type in Glue. The response includes authentication
// configuration details that show supported authentication types and properties,
// and RestConfiguration for custom REST-based connection types registered via
// RegisterConnectionType .
//
// See also: ListConnectionTypes , RegisterConnectionType , DeleteConnectionType
func glue_DescribeConnectionType(cfg aws.Config, client *glue.Client) {
	input := &glue.DescribeConnectionTypeInput{
		// ConnectionType: *string, // Required
	}

	if len(_glueConnectionType) > 0 {
		input.ConnectionType = aws.String(_glueConnectionType)
	}

	if resp, err := client.DescribeConnectionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details regarding the entity used with the connection type, with a
// description of the data model for each field in the selected entity.
//
// The response includes all the fields which make up the entity.
func glue_DescribeEntity(cfg aws.Config, client *glue.Client) {
	input := &glue.DescribeEntityInput{
		// ConnectionName: *string, // Required
		// EntityName: *string, // Required
	}

	if len(_glueConnectionName) > 0 {
		input.ConnectionName = aws.String(_glueConnectionName)
	}
	if len(_glueEntityName) > 0 {
		input.EntityName = aws.String(_glueEntityName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDataStoreApiVersion) > 0 {
		input.DataStoreApiVersion = aws.String(_glueDataStoreApiVersion)
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEntity(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.DescribeEntityOutput
	p := glue.NewDescribeEntityPaginator(client, input)
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

// Returns a list of inbound integrations for the specified integration.
func glue_DescribeInboundIntegrations(cfg aws.Config, client *glue.Client) {
	input := &glue.DescribeInboundIntegrationsInput{}

	if len(_glueIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_glueIntegrationArn)
	}
	if len(_glueMarker) > 0 {
		input.Marker = aws.String(_glueMarker)
	}
	if len(_glueMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _glueMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_glueTargetArn) > 0 {
		input.TargetArn = aws.String(_glueTargetArn)
	}

	if resp, err := client.DescribeInboundIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The API is used to retrieve a list of integrations.
func glue_DescribeIntegrations(cfg aws.Config, client *glue.Client) {
	input := &glue.DescribeIntegrationsInput{}

	if len(_glueFilters) > 0 {
		if err := assignInputField(input, "Filters", _glueFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_glueIntegrationIdentifier) > 0 {
		input.IntegrationIdentifier = aws.String(_glueIntegrationIdentifier)
	}
	if len(_glueMarker) > 0 {
		input.Marker = aws.String(_glueMarker)
	}
	if len(_glueMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _glueMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a blueprint.
func glue_GetBlueprint(cfg aws.Config, client *glue.Client) {
	input := &glue.GetBlueprintInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueIncludeBlueprint) > 0 {
		if err := assignInputField(input, "IncludeBlueprint", _glueIncludeBlueprint); err != nil {
			log.Errorf("invalid --include-blueprint: %s", err.Error())
			return
		}
	}
	if len(_glueIncludeParameterSpec) > 0 {
		if err := assignInputField(input, "IncludeParameterSpec", _glueIncludeParameterSpec); err != nil {
			log.Errorf("invalid --include-parameter-spec: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a blueprint run.
func glue_GetBlueprintRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetBlueprintRunInput{
		// BlueprintName: *string, // Required
		// RunId: *string, // Required
	}

	if len(_glueBlueprintName) > 0 {
		input.BlueprintName = aws.String(_glueBlueprintName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.GetBlueprintRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of blueprint runs for a specified blueprint.
func glue_GetBlueprintRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.GetBlueprintRunsInput{
		// BlueprintName: *string, // Required
	}

	if len(_glueBlueprintName) > 0 {
		input.BlueprintName = aws.String(_glueBlueprintName)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBlueprintRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetBlueprintRunsOutput
	p := glue.NewGetBlueprintRunsPaginator(client, input)
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

// The name of the Catalog to retrieve. This should be all lowercase.
func glue_GetCatalog(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCatalogInput{
		// CatalogId: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of a migration operation.
func glue_GetCatalogImportStatus(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCatalogImportStatusInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetCatalogImportStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all catalogs defined in a catalog in the Glue Data Catalog. For a
// Redshift-federated catalog use case, this operation returns the list of catalogs
// mapped to Redshift databases in the Redshift namespace catalog.
func glue_GetCatalogs(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCatalogsInput{}

	if len(_glueIncludeRoot) > 0 {
		if err := assignInputField(input, "IncludeRoot", _glueIncludeRoot); err != nil {
			log.Errorf("invalid --include-root: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueParentCatalogId) > 0 {
		input.ParentCatalogId = aws.String(_glueParentCatalogId)
	}
	if len(_glueRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _glueRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCatalogs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a classifier by name.
func glue_GetClassifier(cfg aws.Config, client *glue.Client) {
	input := &glue.GetClassifierInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all classifier objects in the Data Catalog.
func glue_GetClassifiers(cfg aws.Config, client *glue.Client) {
	input := &glue.GetClassifiersInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetClassifiers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetClassifiersOutput
	p := glue.NewGetClassifiersPaginator(client, input)
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

// Retrieves partition statistics of columns.
// The Identity and Access Management (IAM) permission required for this operation
// is GetPartition .
func glue_GetColumnStatisticsForPartition(cfg aws.Config, client *glue.Client) {
	input := &glue.GetColumnStatisticsForPartitionInput{
		// ColumnNames: []string, // Required
		// DatabaseName: *string, // Required
		// PartitionValues: []string, // Required
		// TableName: *string, // Required
	}

	if len(_glueColumnNames) > 0 {
		input.ColumnNames = append([]string(nil), _glueColumnNames...)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionValues) > 0 {
		input.PartitionValues = append([]string(nil), _gluePartitionValues...)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetColumnStatisticsForPartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves table statistics of columns.
// The Identity and Access Management (IAM) permission required for this operation
// is GetTable .
func glue_GetColumnStatisticsForTable(cfg aws.Config, client *glue.Client) {
	input := &glue.GetColumnStatisticsForTableInput{
		// ColumnNames: []string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueColumnNames) > 0 {
		input.ColumnNames = append([]string(nil), _glueColumnNames...)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetColumnStatisticsForTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the associated metadata/information for a task run, given a task run ID.
func glue_GetColumnStatisticsTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetColumnStatisticsTaskRunInput{
		// ColumnStatisticsTaskRunId: *string, // Required
	}

	if len(_glueColumnStatisticsTaskRunId) > 0 {
		input.ColumnStatisticsTaskRunId = aws.String(_glueColumnStatisticsTaskRunId)
	}

	if resp, err := client.GetColumnStatisticsTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all runs associated with the specified table.
func glue_GetColumnStatisticsTaskRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.GetColumnStatisticsTaskRunsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetColumnStatisticsTaskRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetColumnStatisticsTaskRunsOutput
	p := glue.NewGetColumnStatisticsTaskRunsPaginator(client, input)
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

// Gets settings for a column statistics task.
func glue_GetColumnStatisticsTaskSettings(cfg aws.Config, client *glue.Client) {
	input := &glue.GetColumnStatisticsTaskSettingsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.GetColumnStatisticsTaskSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a connection definition from the Data Catalog.
func glue_GetConnection(cfg aws.Config, client *glue.Client) {
	input := &glue.GetConnectionInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueApplyOverrideForComputeEnvironment) > 0 {
		if err := assignInputField(input, "ApplyOverrideForComputeEnvironment", _glueApplyOverrideForComputeEnvironment); err != nil {
			log.Errorf("invalid --apply-override-for-compute-environment: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueHidePassword) > 0 {
		if err := assignInputField(input, "HidePassword", _glueHidePassword); err != nil {
			log.Errorf("invalid --hide-password: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of connection definitions from the Data Catalog.
func glue_GetConnections(cfg aws.Config, client *glue.Client) {
	input := &glue.GetConnectionsInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueHidePassword) > 0 {
		if err := assignInputField(input, "HidePassword", _glueHidePassword); err != nil {
			log.Errorf("invalid --hide-password: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetConnectionsOutput
	p := glue.NewGetConnectionsPaginator(client, input)
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

// Retrieves metadata for a specified crawler.
func glue_GetCrawler(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCrawlerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetCrawler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metrics about specified crawlers.
func glue_GetCrawlerMetrics(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCrawlerMetricsInput{}

	if len(_glueCrawlerNameList) > 0 {
		input.CrawlerNameList = append([]string(nil), _glueCrawlerNameList...)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCrawlerMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetCrawlerMetricsOutput
	p := glue.NewGetCrawlerMetricsPaginator(client, input)
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

// Retrieves metadata for all crawlers defined in the customer account.
func glue_GetCrawlers(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCrawlersInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCrawlers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetCrawlersOutput
	p := glue.NewGetCrawlersPaginator(client, input)
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

// Retrieves the details of a custom pattern by specifying its name.
func glue_GetCustomEntityType(cfg aws.Config, client *glue.Client) {
	input := &glue.GetCustomEntityTypeInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetCustomEntityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the security configuration for a specified catalog.
func glue_GetDataCatalogEncryptionSettings(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataCatalogEncryptionSettingsInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetDataCatalogEncryptionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the training status of the model along with more information
// (CompletedOn, StartedOn, FailureReason).
func glue_GetDataQualityModel(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataQualityModelInput{
		// ProfileId: *string, // Required
	}

	if len(_glueProfileId) > 0 {
		input.ProfileId = aws.String(_glueProfileId)
	}
	if len(_glueStatisticId) > 0 {
		input.StatisticId = aws.String(_glueStatisticId)
	}

	if resp, err := client.GetDataQualityModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a statistic's predictions for a given Profile ID.
func glue_GetDataQualityModelResult(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataQualityModelResultInput{
		// ProfileId: *string, // Required
		// StatisticId: *string, // Required
	}

	if len(_glueProfileId) > 0 {
		input.ProfileId = aws.String(_glueProfileId)
	}
	if len(_glueStatisticId) > 0 {
		input.StatisticId = aws.String(_glueStatisticId)
	}

	if resp, err := client.GetDataQualityModelResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the result of a data quality rule evaluation.
func glue_GetDataQualityResult(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataQualityResultInput{
		// ResultId: *string, // Required
	}

	if len(_glueResultId) > 0 {
		input.ResultId = aws.String(_glueResultId)
	}

	if resp, err := client.GetDataQualityResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified recommendation run that was used to generate rules.
func glue_GetDataQualityRuleRecommendationRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataQualityRuleRecommendationRunInput{
		// RunId: *string, // Required
	}

	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.GetDataQualityRuleRecommendationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an existing ruleset by identifier or name.
func glue_GetDataQualityRuleset(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataQualityRulesetInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetDataQualityRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specific run where a ruleset is evaluated against a data source.
func glue_GetDataQualityRulesetEvaluationRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataQualityRulesetEvaluationRunInput{
		// RunId: *string, // Required
	}

	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.GetDataQualityRulesetEvaluationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the definition of a specified database.
func glue_GetDatabase(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDatabaseInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all databases defined in a given Data Catalog.
func glue_GetDatabases(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDatabasesInput{}

	if len(_glueAttributesToGet) > 0 {
		if err := assignInputField(input, "AttributesToGet", _glueAttributesToGet); err != nil {
			log.Errorf("invalid --attributes-to-get: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueResourceShareType) > 0 {
		if err := assignInputField(input, "ResourceShareType", _glueResourceShareType); err != nil {
			log.Errorf("invalid --resource-share-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetDatabasesOutput
	p := glue.NewGetDatabasesPaginator(client, input)
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

// Transforms a Python script into a directed acyclic graph (DAG).
func glue_GetDataflowGraph(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDataflowGraphInput{}

	if len(_gluePythonScript) > 0 {
		input.PythonScript = aws.String(_gluePythonScript)
	}

	if resp, err := client.GetDataflowGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specified development endpoint.
// When you create a development endpoint in a virtual private cloud (VPC), Glue
// returns only a private IP address, and the public IP address field is not
// populated. When you create a non-VPC development endpoint, Glue returns only a
// public IP address.
func glue_GetDevEndpoint(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDevEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_glueEndpointName) > 0 {
		input.EndpointName = aws.String(_glueEndpointName)
	}

	if resp, err := client.GetDevEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all the development endpoints in this Amazon Web Services account.
// When you create a development endpoint in a virtual private cloud (VPC), Glue
// returns only a private IP address and the public IP address field is not
// populated. When you create a non-VPC development endpoint, Glue returns only a
// public IP address.
func glue_GetDevEndpoints(cfg aws.Config, client *glue.Client) {
	input := &glue.GetDevEndpointsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetDevEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetDevEndpointsOutput
	p := glue.NewGetDevEndpointsPaginator(client, input)
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

// This API is used to query preview data from a given connection type or from a
// native Amazon S3 based Glue Data Catalog.
//
// Returns records as an array of JSON blobs. Each record is formatted using
// Jackson JsonNode based on the field type defined by the DescribeEntity API.
//
// Spark connectors generate schemas according to the same data type mapping as in
// the DescribeEntity API. Spark connectors convert data to the appropriate data
// types matching the schema when returning rows.
func glue_GetEntityRecords(cfg aws.Config, client *glue.Client) {
	input := &glue.GetEntityRecordsInput{
		// EntityName: *string, // Required
		// Limit: *int64, // Required
	}

	if len(_glueEntityName) > 0 {
		input.EntityName = aws.String(_glueEntityName)
	}
	if len(_glueLimit) > 0 {
		if err := assignInputField(input, "Limit", _glueLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueConnectionName) > 0 {
		input.ConnectionName = aws.String(_glueConnectionName)
	}
	if len(_glueConnectionOptions) > 0 {
		if err := assignInputField(input, "ConnectionOptions", _glueConnectionOptions); err != nil {
			log.Errorf("invalid --connection-options: %s", err.Error())
			return
		}
	}
	if len(_glueDataStoreApiVersion) > 0 {
		input.DataStoreApiVersion = aws.String(_glueDataStoreApiVersion)
	}
	if len(_glueFilterPredicate) > 0 {
		input.FilterPredicate = aws.String(_glueFilterPredicate)
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueOrderBy) > 0 {
		input.OrderBy = aws.String(_glueOrderBy)
	}
	if len(_glueSelectedFields) > 0 {
		input.SelectedFields = append([]string(nil), _glueSelectedFields...)
	}

	if resp, err := client.GetEntityRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current Glue Identity Center configuration details, including the
// associated Identity Center instance and application information.
func glue_GetGlueIdentityCenterConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.GetGlueIdentityCenterConfigurationInput{}

	if resp, err := client.GetGlueIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is used for fetching the ResourceProperty of the Glue connection (for
// the source) or Glue database ARN (for the target)
func glue_GetIntegrationResourceProperty(cfg aws.Config, client *glue.Client) {
	input := &glue.GetIntegrationResourcePropertyInput{
		// ResourceArn: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}

	if resp, err := client.GetIntegrationResourceProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is used to retrieve optional override properties for the tables that
// need to be replicated. These properties can include properties for filtering and
// partition for source and target tables.
func glue_GetIntegrationTableProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.GetIntegrationTablePropertiesInput{
		// ResourceArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.GetIntegrationTableProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an existing job definition.
func glue_GetJob(cfg aws.Config, client *glue.Client) {
	input := &glue.GetJobInput{
		// JobName: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information on a job bookmark entry.
// For more information about enabling and using job bookmarks, see:
//
// [Tracking processed data using job bookmarks]
//
// [Job parameters used by Glue]
//
// [Job structure]
//
// [Job parameters used by Glue]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
// [Tracking processed data using job bookmarks]: https://docs.aws.amazon.com/glue/latest/dg/monitor-continuations.html
// [Job structure]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html#aws-glue-api-jobs-job-Job
func glue_GetJobBookmark(cfg aws.Config, client *glue.Client) {
	input := &glue.GetJobBookmarkInput{
		// JobName: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.GetJobBookmark(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata for a given job run. Job run history is accessible for
// 365 days for your workflow and job run.
func glue_GetJobRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetJobRunInput{
		// JobName: *string, // Required
		// RunId: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}
	if len(_gluePredecessorsIncluded) > 0 {
		if err := assignInputField(input, "PredecessorsIncluded", _gluePredecessorsIncluded); err != nil {
			log.Errorf("invalid --predecessors-included: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata for all runs of a given job definition.
// GetJobRuns returns the job runs in chronological order, with the newest jobs
// returned first.
func glue_GetJobRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.GetJobRunsInput{
		// JobName: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetJobRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetJobRunsOutput
	p := glue.NewGetJobRunsPaginator(client, input)
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

// Retrieves all current job definitions.
func glue_GetJobs(cfg aws.Config, client *glue.Client) {
	input := &glue.GetJobsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetJobsOutput
	p := glue.NewGetJobsPaginator(client, input)
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

// Creates mappings.
func glue_GetMapping(cfg aws.Config, client *glue.Client) {
	input := &glue.GetMappingInput{
		// Source: *types.CatalogEntry, // Required
	}

	if len(_glueSource) > 0 {
		if err := assignInputField(input, "Source", _glueSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_glueLocation) > 0 {
		if err := assignInputField(input, "Location", _glueLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_glueSinks) > 0 {
		if err := assignInputField(input, "Sinks", _glueSinks); err != nil {
			log.Errorf("invalid --sinks: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the associated metadata/information for a task run, given a task run ID.
func glue_GetMaterializedViewRefreshTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetMaterializedViewRefreshTaskRunInput{
		// CatalogId: *string, // Required
		// MaterializedViewRefreshTaskRunId: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueMaterializedViewRefreshTaskRunId) > 0 {
		input.MaterializedViewRefreshTaskRunId = aws.String(_glueMaterializedViewRefreshTaskRunId)
	}

	if resp, err := client.GetMaterializedViewRefreshTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details for a specific task run on a machine learning transform. Machine
// learning task runs are asynchronous tasks that Glue runs on your behalf as part
// of various machine learning workflows. You can check the stats of any task run
// by calling GetMLTaskRun with the TaskRunID and its parent transform's
// TransformID .
func glue_GetMLTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetMLTaskRunInput{
		// TaskRunId: *string, // Required
		// TransformId: *string, // Required
	}

	if len(_glueTaskRunId) > 0 {
		input.TaskRunId = aws.String(_glueTaskRunId)
	}
	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.GetMLTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of runs for a machine learning transform. Machine learning task
// runs are asynchronous tasks that Glue runs on your behalf as part of various
// machine learning workflows. You can get a sortable, filterable list of machine
// learning task runs by calling GetMLTaskRuns with their parent transform's
// TransformID and other optional parameters as documented in this section.
//
// This operation returns a list of historic runs and must be paginated.
func glue_GetMLTaskRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.GetMLTaskRunsInput{
		// TransformId: *string, // Required
	}

	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}
	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueSort) > 0 {
		if err := assignInputField(input, "Sort", _glueSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetMLTaskRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetMLTaskRunsOutput
	p := glue.NewGetMLTaskRunsPaginator(client, input)
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

// Gets an Glue machine learning transform artifact and all its corresponding
// metadata. Machine learning transforms are a special type of transform that use
// machine learning to learn the details of the transformation to be performed by
// learning from examples provided by humans. These transformations are then saved
// by Glue. You can retrieve their metadata by calling GetMLTransform .
func glue_GetMLTransform(cfg aws.Config, client *glue.Client) {
	input := &glue.GetMLTransformInput{
		// TransformId: *string, // Required
	}

	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.GetMLTransform(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a sortable, filterable list of existing Glue machine learning transforms.
// Machine learning transforms are a special type of transform that use machine
// learning to learn the details of the transformation to be performed by learning
// from examples provided by humans. These transformations are then saved by Glue,
// and you can retrieve their metadata by calling GetMLTransforms .
func glue_GetMLTransforms(cfg aws.Config, client *glue.Client) {
	input := &glue.GetMLTransformsInput{}

	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueSort) > 0 {
		if err := assignInputField(input, "Sort", _glueSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetMLTransforms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetMLTransformsOutput
	p := glue.NewGetMLTransformsPaginator(client, input)
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

// Retrieves information about a specified partition.
func glue_GetPartition(cfg aws.Config, client *glue.Client) {
	input := &glue.GetPartitionInput{
		// DatabaseName: *string, // Required
		// PartitionValues: []string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionValues) > 0 {
		input.PartitionValues = append([]string(nil), _gluePartitionValues...)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetPartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the partition indexes associated with a table.
func glue_GetPartitionIndexes(cfg aws.Config, client *glue.Client) {
	input := &glue.GetPartitionIndexesInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetPartitionIndexes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetPartitionIndexesOutput
	p := glue.NewGetPartitionIndexesPaginator(client, input)
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

// Retrieves information about the partitions in a table.
func glue_GetPartitions(cfg aws.Config, client *glue.Client) {
	input := &glue.GetPartitionsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueExcludeColumnSchema) > 0 {
		if err := assignInputField(input, "ExcludeColumnSchema", _glueExcludeColumnSchema); err != nil {
			log.Errorf("invalid --exclude-column-schema: %s", err.Error())
			return
		}
	}
	if len(_glueExpression) > 0 {
		input.Expression = aws.String(_glueExpression)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueQueryAsOfTime) > 0 {
		if err := assignInputField(input, "QueryAsOfTime", _glueQueryAsOfTime); err != nil {
			log.Errorf("invalid --query-as-of-time: %s", err.Error())
			return
		}
	}
	if len(_glueSegment) > 0 {
		if err := assignInputField(input, "Segment", _glueSegment); err != nil {
			log.Errorf("invalid --segment: %s", err.Error())
			return
		}
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}

	if disablePaginator() {
		if resp, err := client.GetPartitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetPartitionsOutput
	p := glue.NewGetPartitionsPaginator(client, input)
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

// Gets code to perform a specified mapping.
func glue_GetPlan(cfg aws.Config, client *glue.Client) {
	input := &glue.GetPlanInput{
		// Mapping: []types.MappingEntry, // Required
		// Source: *types.CatalogEntry, // Required
	}

	if len(_glueMapping) > 0 {
		if err := assignInputField(input, "Mapping", _glueMapping); err != nil {
			log.Errorf("invalid --mapping: %s", err.Error())
			return
		}
	}
	if len(_glueSource) > 0 {
		if err := assignInputField(input, "Source", _glueSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_glueAdditionalPlanOptionsMap) > 0 {
		if err := assignInputField(input, "AdditionalPlanOptionsMap", _glueAdditionalPlanOptionsMap); err != nil {
			log.Errorf("invalid --additional-plan-options-map: %s", err.Error())
			return
		}
	}
	if len(_glueLanguage) > 0 {
		if err := assignInputField(input, "Language", _glueLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_glueLocation) > 0 {
		if err := assignInputField(input, "Location", _glueLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_glueSinks) > 0 {
		if err := assignInputField(input, "Sinks", _glueSinks); err != nil {
			log.Errorf("invalid --sinks: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified registry in detail.
func glue_GetRegistry(cfg aws.Config, client *glue.Client) {
	input := &glue.GetRegistryInput{
		// RegistryId: *types.RegistryId, // Required
	}

	if len(_glueRegistryId) > 0 {
		if err := assignInputField(input, "RegistryId", _glueRegistryId); err != nil {
			log.Errorf("invalid --registry-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource policies set on individual resources by Resource Access
// Manager during cross-account permission grants. Also retrieves the Data Catalog
// resource policy.
//
// If you enabled metadata encryption in Data Catalog settings, and you do not
// have permission on the KMS key, the operation can't return the Data Catalog
// resource policy.
func glue_GetResourcePolicies(cfg aws.Config, client *glue.Client) {
	input := &glue.GetResourcePoliciesInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetResourcePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetResourcePoliciesOutput
	p := glue.NewGetResourcePoliciesPaginator(client, input)
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

// Retrieves a specified resource policy.
func glue_GetResourcePolicy(cfg aws.Config, client *glue.Client) {
	input := &glue.GetResourcePolicyInput{}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified schema in detail.
func glue_GetSchema(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSchemaInput{
		// SchemaId: *types.SchemaId, // Required
	}

	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a schema by the SchemaDefinition . The schema definition is sent to
// the Schema Registry, canonicalized, and hashed. If the hash is matched within
// the scope of the SchemaName or ARN (or the default registry, if none is
// supplied), that schema’s metadata is returned. Otherwise, a 404 or NotFound
// error is returned. Schema versions in Deleted statuses will not be included in
// the results.
func glue_GetSchemaByDefinition(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSchemaByDefinitionInput{
		// SchemaDefinition: *string, // Required
		// SchemaId: *types.SchemaId, // Required
	}

	if len(_glueSchemaDefinition) > 0 {
		input.SchemaDefinition = aws.String(_glueSchemaDefinition)
	}
	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSchemaByDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the specified schema by its unique ID assigned when a version of the schema
// is created or registered. Schema versions in Deleted status will not be included
// in the results.
func glue_GetSchemaVersion(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSchemaVersionInput{}

	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaVersionId) > 0 {
		input.SchemaVersionId = aws.String(_glueSchemaVersionId)
	}
	if len(_glueSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "SchemaVersionNumber", _glueSchemaVersionNumber); err != nil {
			log.Errorf("invalid --schema-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSchemaVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the schema version difference in the specified difference type between
// two stored schema versions in the Schema Registry.
//
// This API allows you to compare two schema versions between two schema
// definitions under the same schema.
func glue_GetSchemaVersionsDiff(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSchemaVersionsDiffInput{
		// FirstSchemaVersionNumber: *types.SchemaVersionNumber, // Required
		// SchemaDiffType: types.SchemaDiffType, // Required
		// SchemaId: *types.SchemaId, // Required
		// SecondSchemaVersionNumber: *types.SchemaVersionNumber, // Required
	}

	if len(_glueFirstSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "FirstSchemaVersionNumber", _glueFirstSchemaVersionNumber); err != nil {
			log.Errorf("invalid --first-schema-version-number: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaDiffType) > 0 {
		if err := assignInputField(input, "SchemaDiffType", _glueSchemaDiffType); err != nil {
			log.Errorf("invalid --schema-diff-type: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueSecondSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "SecondSchemaVersionNumber", _glueSecondSchemaVersionNumber); err != nil {
			log.Errorf("invalid --second-schema-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSchemaVersionsDiff(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified security configuration.
func glue_GetSecurityConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSecurityConfigurationInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all security configurations.
func glue_GetSecurityConfigurations(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSecurityConfigurationsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSecurityConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetSecurityConfigurationsOutput
	p := glue.NewGetSecurityConfigurationsPaginator(client, input)
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

// Retrieves the session.
func glue_GetSession(cfg aws.Config, client *glue.Client) {
	input := &glue.GetSessionInput{
		// Id: *string, // Required
	}

	if len(_glueId) > 0 {
		input.Id = aws.String(_glueId)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the statement.
func glue_GetStatement(cfg aws.Config, client *glue.Client) {
	input := &glue.GetStatementInput{
		// Id: int32, // Required
		// SessionId: *string, // Required
	}

	if len(_glueId) > 0 {
		if err := assignInputField(input, "Id", _glueId); err != nil {
			log.Errorf("invalid --id: %s", err.Error())
			return
		}
	}
	if len(_glueSessionId) > 0 {
		input.SessionId = aws.String(_glueSessionId)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.GetStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Table definition in a Data Catalog for a specified table.
func glue_GetTable(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTableInput{
		// DatabaseName: *string, // Required
		// Name: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _glueAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueIncludeStatusDetails) > 0 {
		if err := assignInputField(input, "IncludeStatusDetails", _glueIncludeStatusDetails); err != nil {
			log.Errorf("invalid --include-status-details: %s", err.Error())
			return
		}
	}
	if len(_glueQueryAsOfTime) > 0 {
		if err := assignInputField(input, "QueryAsOfTime", _glueQueryAsOfTime); err != nil {
			log.Errorf("invalid --query-as-of-time: %s", err.Error())
			return
		}
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}

	if resp, err := client.GetTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the configuration of all optimizers associated with a specified table.
func glue_GetTableOptimizer(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTableOptimizerInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// Type: types.TableOptimizerType, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueType) > 0 {
		if err := assignInputField(input, "Type", _glueType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTableOptimizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified version of a table.
func glue_GetTableVersion(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTableVersionInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueVersionId) > 0 {
		input.VersionId = aws.String(_glueVersionId)
	}

	if resp, err := client.GetTableVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of strings that identify available versions of a specified
// table.
func glue_GetTableVersions(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTableVersionsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetTableVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetTableVersionsOutput
	p := glue.NewGetTableVersionsPaginator(client, input)
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

// Retrieves the definitions of some or all of the tables in a given Database .
func glue_GetTables(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTablesInput{
		// DatabaseName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueAttributesToGet) > 0 {
		if err := assignInputField(input, "AttributesToGet", _glueAttributesToGet); err != nil {
			log.Errorf("invalid --attributes-to-get: %s", err.Error())
			return
		}
	}
	if len(_glueAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _glueAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueExpression) > 0 {
		input.Expression = aws.String(_glueExpression)
	}
	if len(_glueIncludeStatusDetails) > 0 {
		if err := assignInputField(input, "IncludeStatusDetails", _glueIncludeStatusDetails); err != nil {
			log.Errorf("invalid --include-status-details: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueQueryAsOfTime) > 0 {
		if err := assignInputField(input, "QueryAsOfTime", _glueQueryAsOfTime); err != nil {
			log.Errorf("invalid --query-as-of-time: %s", err.Error())
			return
		}
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}

	if disablePaginator() {
		if resp, err := client.GetTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetTablesOutput
	p := glue.NewGetTablesPaginator(client, input)
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

// Retrieves a list of tags associated with a resource.
func glue_GetTags(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}

	if resp, err := client.GetTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the definition of a trigger.
func glue_GetTrigger(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTriggerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all the triggers associated with a job.
func glue_GetTriggers(cfg aws.Config, client *glue.Client) {
	input := &glue.GetTriggersInput{}

	if len(_glueDependentJobName) > 0 {
		input.DependentJobName = aws.String(_glueDependentJobName)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetTriggers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetTriggersOutput
	p := glue.NewGetTriggersPaginator(client, input)
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

// Retrieves partition metadata from the Data Catalog that contains unfiltered
// metadata.
//
// For IAM authorization, the public IAM action associated with this API is
// glue:GetPartition .
func glue_GetUnfilteredPartitionMetadata(cfg aws.Config, client *glue.Client) {
	input := &glue.GetUnfilteredPartitionMetadataInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// PartitionValues: []string, // Required
		// SupportedPermissionTypes: []types.PermissionType, // Required
		// TableName: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionValues) > 0 {
		input.PartitionValues = append([]string(nil), _gluePartitionValues...)
	}
	if len(_glueSupportedPermissionTypes) > 0 {
		if err := assignInputField(input, "SupportedPermissionTypes", _glueSupportedPermissionTypes); err != nil {
			log.Errorf("invalid --supported-permission-types: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _glueAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_glueQuerySessionContext) > 0 {
		if err := assignInputField(input, "QuerySessionContext", _glueQuerySessionContext); err != nil {
			log.Errorf("invalid --query-session-context: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUnfilteredPartitionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves partition metadata from the Data Catalog that contains unfiltered
// metadata.
//
// For IAM authorization, the public IAM action associated with this API is
// glue:GetPartitions .
func glue_GetUnfilteredPartitionsMetadata(cfg aws.Config, client *glue.Client) {
	input := &glue.GetUnfilteredPartitionsMetadataInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// SupportedPermissionTypes: []types.PermissionType, // Required
		// TableName: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueSupportedPermissionTypes) > 0 {
		if err := assignInputField(input, "SupportedPermissionTypes", _glueSupportedPermissionTypes); err != nil {
			log.Errorf("invalid --supported-permission-types: %s", err.Error())
			return
		}
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _glueAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_glueExpression) > 0 {
		input.Expression = aws.String(_glueExpression)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueQuerySessionContext) > 0 {
		if err := assignInputField(input, "QuerySessionContext", _glueQuerySessionContext); err != nil {
			log.Errorf("invalid --query-session-context: %s", err.Error())
			return
		}
	}
	if len(_glueSegment) > 0 {
		if err := assignInputField(input, "Segment", _glueSegment); err != nil {
			log.Errorf("invalid --segment: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetUnfilteredPartitionsMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetUnfilteredPartitionsMetadataOutput
	p := glue.NewGetUnfilteredPartitionsMetadataPaginator(client, input)
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

// Allows a third-party analytical engine to retrieve unfiltered table metadata
// from the Data Catalog.
//
// For IAM authorization, the public IAM action associated with this API is
// glue:GetTable .
func glue_GetUnfilteredTableMetadata(cfg aws.Config, client *glue.Client) {
	input := &glue.GetUnfilteredTableMetadataInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// Name: *string, // Required
		// SupportedPermissionTypes: []types.PermissionType, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueSupportedPermissionTypes) > 0 {
		if err := assignInputField(input, "SupportedPermissionTypes", _glueSupportedPermissionTypes); err != nil {
			log.Errorf("invalid --supported-permission-types: %s", err.Error())
			return
		}
	}
	if len(_glueAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _glueAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_glueParentResourceArn) > 0 {
		input.ParentResourceArn = aws.String(_glueParentResourceArn)
	}
	if len(_gluePermissions) > 0 {
		if err := assignInputField(input, "Permissions", _gluePermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_glueQuerySessionContext) > 0 {
		if err := assignInputField(input, "QuerySessionContext", _glueQuerySessionContext); err != nil {
			log.Errorf("invalid --query-session-context: %s", err.Error())
			return
		}
	}
	if len(_glueRootResourceArn) > 0 {
		input.RootResourceArn = aws.String(_glueRootResourceArn)
	}
	if len(_glueSupportedDialect) > 0 {
		if err := assignInputField(input, "SupportedDialect", _glueSupportedDialect); err != nil {
			log.Errorf("invalid --supported-dialect: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUnfilteredTableMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified Glue usage profile.
func glue_GetUsageProfile(cfg aws.Config, client *glue.Client) {
	input := &glue.GetUsageProfileInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.GetUsageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified function definition from the Data Catalog.
func glue_GetUserDefinedFunction(cfg aws.Config, client *glue.Client) {
	input := &glue.GetUserDefinedFunctionInput{
		// DatabaseName: *string, // Required
		// FunctionName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueFunctionName) > 0 {
		input.FunctionName = aws.String(_glueFunctionName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.GetUserDefinedFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves multiple function definitions from the Data Catalog.
func glue_GetUserDefinedFunctions(cfg aws.Config, client *glue.Client) {
	input := &glue.GetUserDefinedFunctionsInput{
		// Pattern: *string, // Required
	}

	if len(_gluePattern) > 0 {
		input.Pattern = aws.String(_gluePattern)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueFunctionType) > 0 {
		if err := assignInputField(input, "FunctionType", _glueFunctionType); err != nil {
			log.Errorf("invalid --function-type: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetUserDefinedFunctions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetUserDefinedFunctionsOutput
	p := glue.NewGetUserDefinedFunctionsPaginator(client, input)
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

// Retrieves resource metadata for a workflow.
func glue_GetWorkflow(cfg aws.Config, client *glue.Client) {
	input := &glue.GetWorkflowInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueIncludeGraph) > 0 {
		if err := assignInputField(input, "IncludeGraph", _glueIncludeGraph); err != nil {
			log.Errorf("invalid --include-graph: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata for a given workflow run. Job run history is accessible
// for 90 days for your workflow and job run.
func glue_GetWorkflowRun(cfg aws.Config, client *glue.Client) {
	input := &glue.GetWorkflowRunInput{
		// Name: *string, // Required
		// RunId: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}
	if len(_glueIncludeGraph) > 0 {
		if err := assignInputField(input, "IncludeGraph", _glueIncludeGraph); err != nil {
			log.Errorf("invalid --include-graph: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the workflow run properties which were set during the run.
func glue_GetWorkflowRunProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.GetWorkflowRunPropertiesInput{
		// Name: *string, // Required
		// RunId: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.GetWorkflowRunProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata for all runs of a given workflow.
func glue_GetWorkflowRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.GetWorkflowRunsInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueIncludeGraph) > 0 {
		if err := assignInputField(input, "IncludeGraph", _glueIncludeGraph); err != nil {
			log.Errorf("invalid --include-graph: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetWorkflowRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.GetWorkflowRunsOutput
	p := glue.NewGetWorkflowRunsPaginator(client, input)
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

// Imports an existing Amazon Athena Data Catalog to Glue.
func glue_ImportCatalogToGlue(cfg aws.Config, client *glue.Client) {
	input := &glue.ImportCatalogToGlueInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.ImportCatalogToGlue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the blueprint names in an account.
func glue_ListBlueprints(cfg aws.Config, client *glue.Client) {
	input := &glue.ListBlueprintsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBlueprints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListBlueprintsOutput
	p := glue.NewListBlueprintsPaginator(client, input)
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

// List all task runs for a particular account.
func glue_ListColumnStatisticsTaskRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.ListColumnStatisticsTaskRunsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListColumnStatisticsTaskRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListColumnStatisticsTaskRunsOutput
	p := glue.NewListColumnStatisticsTaskRunsPaginator(client, input)
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

// The ListConnectionTypes API provides a discovery mechanism to learn available
// connection types in Glue. The response contains a list of connection types with
// high-level details of what is supported for each connection type, including both
// built-in connection types and custom connection types registered via
// RegisterConnectionType . The connection types listed are the set of supported
// options for the ConnectionType value in the CreateConnection API.
//
// See also: DescribeConnectionType , RegisterConnectionType , DeleteConnectionType
func glue_ListConnectionTypes(cfg aws.Config, client *glue.Client) {
	input := &glue.ListConnectionTypesInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectionTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListConnectionTypesOutput
	p := glue.NewListConnectionTypesPaginator(client, input)
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

// Retrieves the names of all crawler resources in this Amazon Web Services
// account, or the resources with the specified tag. This operation allows you to
// see which resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter on
// the response so that tagged resources can be retrieved as a group. If you choose
// to use tags filtering, only resources with the tag are retrieved.
func glue_ListCrawlers(cfg aws.Config, client *glue.Client) {
	input := &glue.ListCrawlersInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCrawlers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListCrawlersOutput
	p := glue.NewListCrawlersPaginator(client, input)
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

// Returns all the crawls of a specified crawler. Returns only the crawls that
// have occurred since the launch date of the crawler history feature, and only
// retains up to 12 months of crawls. Older crawls will not be returned.
//
// You may use this API to:
//
// - Retrive all the crawls of a specified crawler.
//
// - Retrieve all the crawls of a specified crawler within a limited count.
//
// - Retrieve all the crawls of a specified crawler in a specific time range.
//
// - Retrieve all the crawls of a specified crawler with a particular state,
// crawl ID, or DPU hour value.
func glue_ListCrawls(cfg aws.Config, client *glue.Client) {
	input := &glue.ListCrawlsInput{
		// CrawlerName: *string, // Required
	}

	if len(_glueCrawlerName) > 0 {
		input.CrawlerName = aws.String(_glueCrawlerName)
	}
	if len(_glueFilters) > 0 {
		if err := assignInputField(input, "Filters", _glueFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if resp, err := client.ListCrawls(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the custom patterns that have been created.
func glue_ListCustomEntityTypes(cfg aws.Config, client *glue.Client) {
	input := &glue.ListCustomEntityTypesInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCustomEntityTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListCustomEntityTypesOutput
	p := glue.NewListCustomEntityTypesPaginator(client, input)
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

// Returns all data quality execution results for your account.
func glue_ListDataQualityResults(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDataQualityResultsInput{}

	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataQualityResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListDataQualityResultsOutput
	p := glue.NewListDataQualityResultsPaginator(client, input)
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

// Lists the recommendation runs meeting the filter criteria.
func glue_ListDataQualityRuleRecommendationRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDataQualityRuleRecommendationRunsInput{}

	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataQualityRuleRecommendationRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListDataQualityRuleRecommendationRunsOutput
	p := glue.NewListDataQualityRuleRecommendationRunsPaginator(client, input)
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

// Lists all the runs meeting the filter criteria, where a ruleset is evaluated
// against a data source.
func glue_ListDataQualityRulesetEvaluationRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDataQualityRulesetEvaluationRunsInput{}

	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataQualityRulesetEvaluationRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListDataQualityRulesetEvaluationRunsOutput
	p := glue.NewListDataQualityRulesetEvaluationRunsPaginator(client, input)
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

// Returns a paginated list of rulesets for the specified list of Glue tables.
func glue_ListDataQualityRulesets(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDataQualityRulesetsInput{}

	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataQualityRulesets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListDataQualityRulesetsOutput
	p := glue.NewListDataQualityRulesetsPaginator(client, input)
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

// Retrieve annotations for a data quality statistic.
func glue_ListDataQualityStatisticAnnotations(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDataQualityStatisticAnnotationsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueProfileId) > 0 {
		input.ProfileId = aws.String(_glueProfileId)
	}
	if len(_glueStatisticId) > 0 {
		input.StatisticId = aws.String(_glueStatisticId)
	}
	if len(_glueTimestampFilter) > 0 {
		if err := assignInputField(input, "TimestampFilter", _glueTimestampFilter); err != nil {
			log.Errorf("invalid --timestamp-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDataQualityStatisticAnnotations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of data quality statistics.
func glue_ListDataQualityStatistics(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDataQualityStatisticsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueProfileId) > 0 {
		input.ProfileId = aws.String(_glueProfileId)
	}
	if len(_glueStatisticId) > 0 {
		input.StatisticId = aws.String(_glueStatisticId)
	}
	if len(_glueTimestampFilter) > 0 {
		if err := assignInputField(input, "TimestampFilter", _glueTimestampFilter); err != nil {
			log.Errorf("invalid --timestamp-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDataQualityStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the names of all DevEndpoint resources in this Amazon Web Services
// account, or the resources with the specified tag. This operation allows you to
// see which resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter on
// the response so that tagged resources can be retrieved as a group. If you choose
// to use tags filtering, only resources with the tag are retrieved.
func glue_ListDevEndpoints(cfg aws.Config, client *glue.Client) {
	input := &glue.ListDevEndpointsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDevEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListDevEndpointsOutput
	p := glue.NewListDevEndpointsPaginator(client, input)
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

// Returns the available entities supported by the connection type.
func glue_ListEntities(cfg aws.Config, client *glue.Client) {
	input := &glue.ListEntitiesInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueConnectionName) > 0 {
		input.ConnectionName = aws.String(_glueConnectionName)
	}
	if len(_glueDataStoreApiVersion) > 0 {
		input.DataStoreApiVersion = aws.String(_glueDataStoreApiVersion)
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueParentEntityName) > 0 {
		input.ParentEntityName = aws.String(_glueParentEntityName)
	}

	if disablePaginator() {
		if resp, err := client.ListEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListEntitiesOutput
	p := glue.NewListEntitiesPaginator(client, input)
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

// List integration resource properties for a single customer. It supports the
// filters, maxRecords and markers.
func glue_ListIntegrationResourceProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.ListIntegrationResourcePropertiesInput{}

	if len(_glueFilters) > 0 {
		if err := assignInputField(input, "Filters", _glueFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_glueMarker) > 0 {
		input.Marker = aws.String(_glueMarker)
	}
	if len(_glueMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _glueMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListIntegrationResourceProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the names of all job resources in this Amazon Web Services account,
// or the resources with the specified tag. This operation allows you to see which
// resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter on
// the response so that tagged resources can be retrieved as a group. If you choose
// to use tags filtering, only resources with the tag are retrieved.
func glue_ListJobs(cfg aws.Config, client *glue.Client) {
	input := &glue.ListJobsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListJobsOutput
	p := glue.NewListJobsPaginator(client, input)
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

// List all task runs for a particular account.
func glue_ListMaterializedViewRefreshTaskRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.ListMaterializedViewRefreshTaskRunsInput{
		// CatalogId: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if disablePaginator() {
		if resp, err := client.ListMaterializedViewRefreshTaskRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListMaterializedViewRefreshTaskRunsOutput
	p := glue.NewListMaterializedViewRefreshTaskRunsPaginator(client, input)
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

// Retrieves a sortable, filterable list of existing Glue machine learning
// transforms in this Amazon Web Services account, or the resources with the
// specified tag. This operation takes the optional Tags field, which you can use
// as a filter of the responses so that tagged resources can be retrieved as a
// group. If you choose to use tag filtering, only resources with the tags are
// retrieved.
func glue_ListMLTransforms(cfg aws.Config, client *glue.Client) {
	input := &glue.ListMLTransformsInput{}

	if len(_glueFilter) > 0 {
		if err := assignInputField(input, "Filter", _glueFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueSort) > 0 {
		if err := assignInputField(input, "Sort", _glueSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMLTransforms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListMLTransformsOutput
	p := glue.NewListMLTransformsPaginator(client, input)
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

// Returns a list of registries that you have created, with minimal registry
// information. Registries in the Deleting status will not be included in the
// results. Empty results will be returned if there are no registries available.
func glue_ListRegistries(cfg aws.Config, client *glue.Client) {
	input := &glue.ListRegistriesInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRegistries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListRegistriesOutput
	p := glue.NewListRegistriesPaginator(client, input)
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

// Returns a list of schema versions that you have created, with minimal
// information. Schema versions in Deleted status will not be included in the
// results. Empty results will be returned if there are no schema versions
// available.
func glue_ListSchemaVersions(cfg aws.Config, client *glue.Client) {
	input := &glue.ListSchemaVersionsInput{
		// SchemaId: *types.SchemaId, // Required
	}

	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemaVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListSchemaVersionsOutput
	p := glue.NewListSchemaVersionsPaginator(client, input)
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

// Returns a list of schemas with minimal details. Schemas in Deleting status will
// not be included in the results. Empty results will be returned if there are no
// schemas available.
//
// When the RegistryId is not provided, all the schemas across registries will be
// part of the API response.
func glue_ListSchemas(cfg aws.Config, client *glue.Client) {
	input := &glue.ListSchemasInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueRegistryId) > 0 {
		if err := assignInputField(input, "RegistryId", _glueRegistryId); err != nil {
			log.Errorf("invalid --registry-id: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListSchemasOutput
	p := glue.NewListSchemasPaginator(client, input)
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

// Retrieve a list of sessions.
func glue_ListSessions(cfg aws.Config, client *glue.Client) {
	input := &glue.ListSessionsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListSessionsOutput
	p := glue.NewListSessionsPaginator(client, input)
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

// Lists statements for the session.
func glue_ListStatements(cfg aws.Config, client *glue.Client) {
	input := &glue.ListStatementsInput{
		// SessionId: *string, // Required
	}

	if len(_glueSessionId) > 0 {
		input.SessionId = aws.String(_glueSessionId)
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.ListStatements(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the history of previous optimizer runs for a specific table.
func glue_ListTableOptimizerRuns(cfg aws.Config, client *glue.Client) {
	input := &glue.ListTableOptimizerRunsInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// Type: types.TableOptimizerType, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueType) > 0 {
		if err := assignInputField(input, "Type", _glueType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTableOptimizerRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListTableOptimizerRunsOutput
	p := glue.NewListTableOptimizerRunsPaginator(client, input)
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

// Retrieves the names of all trigger resources in this Amazon Web Services
// account, or the resources with the specified tag. This operation allows you to
// see which resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter on
// the response so that tagged resources can be retrieved as a group. If you choose
// to use tags filtering, only resources with the tag are retrieved.
func glue_ListTriggers(cfg aws.Config, client *glue.Client) {
	input := &glue.ListTriggersInput{}

	if len(_glueDependentJobName) > 0 {
		input.DependentJobName = aws.String(_glueDependentJobName)
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTriggers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListTriggersOutput
	p := glue.NewListTriggersPaginator(client, input)
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

// List all the Glue usage profiles.
func glue_ListUsageProfiles(cfg aws.Config, client *glue.Client) {
	input := &glue.ListUsageProfilesInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsageProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListUsageProfilesOutput
	p := glue.NewListUsageProfilesPaginator(client, input)
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

// Lists names of workflows created in the account.
func glue_ListWorkflows(cfg aws.Config, client *glue.Client) {
	input := &glue.ListWorkflowsInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.ListWorkflowsOutput
	p := glue.NewListWorkflowsPaginator(client, input)
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

// Modifies a Zero-ETL integration in the caller's account.
func glue_ModifyIntegration(cfg aws.Config, client *glue.Client) {
	input := &glue.ModifyIntegrationInput{
		// IntegrationIdentifier: *string, // Required
	}

	if len(_glueIntegrationIdentifier) > 0 {
		input.IntegrationIdentifier = aws.String(_glueIntegrationIdentifier)
	}
	if len(_glueDataFilter) > 0 {
		input.DataFilter = aws.String(_glueDataFilter)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueIntegrationConfig) > 0 {
		if err := assignInputField(input, "IntegrationConfig", _glueIntegrationConfig); err != nil {
			log.Errorf("invalid --integration-config: %s", err.Error())
			return
		}
	}
	if len(_glueIntegrationName) > 0 {
		input.IntegrationName = aws.String(_glueIntegrationName)
	}

	if resp, err := client.ModifyIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the security configuration for a specified catalog. After the
// configuration has been set, the specified encryption is applied to every catalog
// write thereafter.
func glue_PutDataCatalogEncryptionSettings(cfg aws.Config, client *glue.Client) {
	input := &glue.PutDataCatalogEncryptionSettingsInput{
		// DataCatalogEncryptionSettings: *types.DataCatalogEncryptionSettings, // Required
	}

	if len(_glueDataCatalogEncryptionSettings) > 0 {
		if err := assignInputField(input, "DataCatalogEncryptionSettings", _glueDataCatalogEncryptionSettings); err != nil {
			log.Errorf("invalid --data-catalog-encryption-settings: %s", err.Error())
			return
		}
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.PutDataCatalogEncryptionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Annotate all datapoints for a Profile.
func glue_PutDataQualityProfileAnnotation(cfg aws.Config, client *glue.Client) {
	input := &glue.PutDataQualityProfileAnnotationInput{
		// InclusionAnnotation: types.InclusionAnnotationValue, // Required
		// ProfileId: *string, // Required
	}

	if len(_glueInclusionAnnotation) > 0 {
		if err := assignInputField(input, "InclusionAnnotation", _glueInclusionAnnotation); err != nil {
			log.Errorf("invalid --inclusion-annotation: %s", err.Error())
			return
		}
	}
	if len(_glueProfileId) > 0 {
		input.ProfileId = aws.String(_glueProfileId)
	}

	if resp, err := client.PutDataQualityProfileAnnotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the Data Catalog resource policy for access control.
func glue_PutResourcePolicy(cfg aws.Config, client *glue.Client) {
	input := &glue.PutResourcePolicyInput{
		// PolicyInJson: *string, // Required
	}

	if len(_gluePolicyInJson) > 0 {
		input.PolicyInJson = aws.String(_gluePolicyInJson)
	}
	if len(_glueEnableHybrid) > 0 {
		if err := assignInputField(input, "EnableHybrid", _glueEnableHybrid); err != nil {
			log.Errorf("invalid --enable-hybrid: %s", err.Error())
			return
		}
	}
	if len(_gluePolicyExistsCondition) > 0 {
		if err := assignInputField(input, "PolicyExistsCondition", _gluePolicyExistsCondition); err != nil {
			log.Errorf("invalid --policy-exists-condition: %s", err.Error())
			return
		}
	}
	if len(_gluePolicyHashCondition) > 0 {
		input.PolicyHashCondition = aws.String(_gluePolicyHashCondition)
	}
	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts the metadata key value pair for a specified schema version ID. A maximum
// of 10 key value pairs will be allowed per schema version. They can be added over
// one or more calls.
func glue_PutSchemaVersionMetadata(cfg aws.Config, client *glue.Client) {
	input := &glue.PutSchemaVersionMetadataInput{
		// MetadataKeyValue: *types.MetadataKeyValuePair, // Required
	}

	if len(_glueMetadataKeyValue) > 0 {
		if err := assignInputField(input, "MetadataKeyValue", _glueMetadataKeyValue); err != nil {
			log.Errorf("invalid --metadata-key-value: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaVersionId) > 0 {
		input.SchemaVersionId = aws.String(_glueSchemaVersionId)
	}
	if len(_glueSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "SchemaVersionNumber", _glueSchemaVersionNumber); err != nil {
			log.Errorf("invalid --schema-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSchemaVersionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts the specified workflow run properties for the given workflow run. If a
// property already exists for the specified run, then it overrides the value
// otherwise adds the property to existing properties.
func glue_PutWorkflowRunProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.PutWorkflowRunPropertiesInput{
		// Name: *string, // Required
		// RunId: *string, // Required
		// RunProperties: map[string]string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}
	if len(_glueRunProperties) > 0 {
		if err := assignInputField(input, "RunProperties", _glueRunProperties); err != nil {
			log.Errorf("invalid --run-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutWorkflowRunProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Queries for the schema version metadata information.
func glue_QuerySchemaVersionMetadata(cfg aws.Config, client *glue.Client) {
	input := &glue.QuerySchemaVersionMetadataInput{}

	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueMetadataList) > 0 {
		if err := assignInputField(input, "MetadataList", _glueMetadataList); err != nil {
			log.Errorf("invalid --metadata-list: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaVersionId) > 0 {
		input.SchemaVersionId = aws.String(_glueSchemaVersionId)
	}
	if len(_glueSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "SchemaVersionNumber", _glueSchemaVersionNumber); err != nil {
			log.Errorf("invalid --schema-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.QuerySchemaVersionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a custom connection type in Glue based on the configuration provided.
// This operation enables customers to configure custom connectors for any data
// source with REST-based APIs, eliminating the need for building custom Lambda
// connectors.
//
// The registered connection type stores details about how requests and responses
// are interpreted by REST sources, including connection properties, authentication
// configuration, and REST configuration with entity definitions. Once registered,
// customers can create connections using this connection type and work with them
// the same way as natively supported Glue connectors.
//
// Supports multiple authentication types including Basic, OAuth2 (Client
// Credentials, JWT Bearer, Authorization Code), and Custom Auth configurations.
func glue_RegisterConnectionType(cfg aws.Config, client *glue.Client) {
	input := &glue.RegisterConnectionTypeInput{
		// ConnectionProperties: *types.ConnectionPropertiesConfiguration, // Required
		// ConnectionType: *string, // Required
		// ConnectorAuthenticationConfiguration: *types.ConnectorAuthenticationConfiguration, // Required
		// IntegrationType: types.IntegrationType, // Required
		// RestConfiguration: *types.RestConfiguration, // Required
	}

	if len(_glueConnectionProperties) > 0 {
		if err := assignInputField(input, "ConnectionProperties", _glueConnectionProperties); err != nil {
			log.Errorf("invalid --connection-properties: %s", err.Error())
			return
		}
	}
	if len(_glueConnectionType) > 0 {
		input.ConnectionType = aws.String(_glueConnectionType)
	}
	if len(_glueConnectorAuthenticationConfiguration) > 0 {
		if err := assignInputField(input, "ConnectorAuthenticationConfiguration", _glueConnectorAuthenticationConfiguration); err != nil {
			log.Errorf("invalid --connector-authentication-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueIntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _glueIntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_glueRestConfiguration) > 0 {
		if err := assignInputField(input, "RestConfiguration", _glueRestConfiguration); err != nil {
			log.Errorf("invalid --rest-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueTags) > 0 {
		if err := assignInputField(input, "Tags", _glueTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterConnectionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new version to the existing schema. Returns an error if new version of
// schema does not meet the compatibility requirements of the schema set. This API
// will not create a new schema set and will return a 404 error if the schema set
// is not already present in the Schema Registry.
//
// If this is the first schema definition to be registered in the Schema Registry,
// this API will store the schema version and return immediately. Otherwise, this
// call has the potential to run longer than other operations due to compatibility
// modes. You can call the GetSchemaVersion API with the SchemaVersionId to check
// compatibility modes.
//
// If the same schema definition is already stored in Schema Registry as a
// version, the schema ID of the existing schema is returned to the caller.
func glue_RegisterSchemaVersion(cfg aws.Config, client *glue.Client) {
	input := &glue.RegisterSchemaVersionInput{
		// SchemaDefinition: *string, // Required
		// SchemaId: *types.SchemaId, // Required
	}

	if len(_glueSchemaDefinition) > 0 {
		input.SchemaDefinition = aws.String(_glueSchemaDefinition)
	}
	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterSchemaVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a key value pair from the schema version metadata for the specified
// schema version ID.
func glue_RemoveSchemaVersionMetadata(cfg aws.Config, client *glue.Client) {
	input := &glue.RemoveSchemaVersionMetadataInput{
		// MetadataKeyValue: *types.MetadataKeyValuePair, // Required
	}

	if len(_glueMetadataKeyValue) > 0 {
		if err := assignInputField(input, "MetadataKeyValue", _glueMetadataKeyValue); err != nil {
			log.Errorf("invalid --metadata-key-value: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueSchemaVersionId) > 0 {
		input.SchemaVersionId = aws.String(_glueSchemaVersionId)
	}
	if len(_glueSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "SchemaVersionNumber", _glueSchemaVersionNumber); err != nil {
			log.Errorf("invalid --schema-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveSchemaVersionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets a bookmark entry.
// For more information about enabling and using job bookmarks, see:
//
// [Tracking processed data using job bookmarks]
//
// [Job parameters used by Glue]
//
// [Job structure]
//
// [Job parameters used by Glue]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
// [Tracking processed data using job bookmarks]: https://docs.aws.amazon.com/glue/latest/dg/monitor-continuations.html
// [Job structure]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html#aws-glue-api-jobs-job-Job
func glue_ResetJobBookmark(cfg aws.Config, client *glue.Client) {
	input := &glue.ResetJobBookmarkInput{
		// JobName: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.ResetJobBookmark(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts selected nodes of a previous partially completed workflow run and
// resumes the workflow run. The selected nodes and all nodes that are downstream
// from the selected nodes are run.
func glue_ResumeWorkflowRun(cfg aws.Config, client *glue.Client) {
	input := &glue.ResumeWorkflowRunInput{
		// Name: *string, // Required
		// NodeIds: []string, // Required
		// RunId: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueNodeIds) > 0 {
		input.NodeIds = append([]string(nil), _glueNodeIds...)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.ResumeWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes the statement.
func glue_RunStatement(cfg aws.Config, client *glue.Client) {
	input := &glue.RunStatementInput{
		// Code: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_glueCode) > 0 {
		input.Code = aws.String(_glueCode)
	}
	if len(_glueSessionId) > 0 {
		input.SessionId = aws.String(_glueSessionId)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.RunStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches a set of tables based on properties in the table metadata as well as
// on the parent database. You can search against text or filter conditions.
//
// You can only get tables that you have access to based on the security policies
// defined in Lake Formation. You need at least a read-only access to the table for
// it to be returned. If you do not have access to all the columns in the table,
// these columns will not be searched against when returning the list of tables
// back to you. If you have access to the columns but not the data in the columns,
// those columns and the associated metadata for those columns will be included in
// the search.
func glue_SearchTables(cfg aws.Config, client *glue.Client) {
	input := &glue.SearchTablesInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueFilters) > 0 {
		if err := assignInputField(input, "Filters", _glueFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_glueIncludeStatusDetails) > 0 {
		if err := assignInputField(input, "IncludeStatusDetails", _glueIncludeStatusDetails); err != nil {
			log.Errorf("invalid --include-status-details: %s", err.Error())
			return
		}
	}
	if len(_glueMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _glueMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_glueNextToken) > 0 {
		input.NextToken = aws.String(_glueNextToken)
	}
	if len(_glueResourceShareType) > 0 {
		if err := assignInputField(input, "ResourceShareType", _glueResourceShareType); err != nil {
			log.Errorf("invalid --resource-share-type: %s", err.Error())
			return
		}
	}
	if len(_glueSearchText) > 0 {
		input.SearchText = aws.String(_glueSearchText)
	}
	if len(_glueSortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _glueSortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*glue.SearchTablesOutput
	p := glue.NewSearchTablesPaginator(client, input)
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

// Starts a new run of the specified blueprint.
func glue_StartBlueprintRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartBlueprintRunInput{
		// BlueprintName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_glueBlueprintName) > 0 {
		input.BlueprintName = aws.String(_glueBlueprintName)
	}
	if len(_glueRoleArn) > 0 {
		input.RoleArn = aws.String(_glueRoleArn)
	}
	if len(_glueParameters) > 0 {
		input.Parameters = aws.String(_glueParameters)
	}

	if resp, err := client.StartBlueprintRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a column statistics task run, for a specified table and columns.
func glue_StartColumnStatisticsTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartColumnStatisticsTaskRunInput{
		// DatabaseName: *string, // Required
		// Role: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogID = aws.String(_glueCatalogId)
	}
	if len(_glueColumnNameList) > 0 {
		input.ColumnNameList = append([]string(nil), _glueColumnNameList...)
	}
	if len(_glueSampleSize) > 0 {
		if err := assignInputField(input, "SampleSize", _glueSampleSize); err != nil {
			log.Errorf("invalid --sample-size: %s", err.Error())
			return
		}
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}

	if resp, err := client.StartColumnStatisticsTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a column statistics task run schedule.
func glue_StartColumnStatisticsTaskRunSchedule(cfg aws.Config, client *glue.Client) {
	input := &glue.StartColumnStatisticsTaskRunScheduleInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.StartColumnStatisticsTaskRunSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a crawl using the specified crawler, regardless of what is scheduled. If
// the crawler is already running, returns a [CrawlerRunningException].
//
// [CrawlerRunningException]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-exceptions.html#aws-glue-api-exceptions-CrawlerRunningException
func glue_StartCrawler(cfg aws.Config, client *glue.Client) {
	input := &glue.StartCrawlerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.StartCrawler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the schedule state of the specified crawler to SCHEDULED , unless the
// crawler is already running or the schedule state is already SCHEDULED .
func glue_StartCrawlerSchedule(cfg aws.Config, client *glue.Client) {
	input := &glue.StartCrawlerScheduleInput{
		// CrawlerName: *string, // Required
	}

	if len(_glueCrawlerName) > 0 {
		input.CrawlerName = aws.String(_glueCrawlerName)
	}

	if resp, err := client.StartCrawlerSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a recommendation run that is used to generate rules when you don't know
// what rules to write. Glue Data Quality analyzes the data and comes up with
// recommendations for a potential ruleset. You can then triage the ruleset and
// modify the generated ruleset to your liking.
//
// Recommendation runs are automatically deleted after 90 days.
func glue_StartDataQualityRuleRecommendationRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartDataQualityRuleRecommendationRunInput{
		// DataSource: *types.DataSource, // Required
		// Role: *string, // Required
	}

	if len(_glueDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _glueDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueClientToken) > 0 {
		input.ClientToken = aws.String(_glueClientToken)
	}
	if len(_glueCreatedRulesetName) > 0 {
		input.CreatedRulesetName = aws.String(_glueCreatedRulesetName)
	}
	if len(_glueDataQualitySecurityConfiguration) > 0 {
		input.DataQualitySecurityConfiguration = aws.String(_glueDataQualitySecurityConfiguration)
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDataQualityRuleRecommendationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Once you have a ruleset definition (either recommended or your own), you call
// this operation to evaluate the ruleset against a data source (Glue table). The
// evaluation computes results which you can retrieve with the GetDataQualityResult
// API.
func glue_StartDataQualityRulesetEvaluationRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartDataQualityRulesetEvaluationRunInput{
		// DataSource: *types.DataSource, // Required
		// Role: *string, // Required
		// RulesetNames: []string, // Required
	}

	if len(_glueDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _glueDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueRulesetNames) > 0 {
		input.RulesetNames = append([]string(nil), _glueRulesetNames...)
	}
	if len(_glueAdditionalDataSources) > 0 {
		if err := assignInputField(input, "AdditionalDataSources", _glueAdditionalDataSources); err != nil {
			log.Errorf("invalid --additional-data-sources: %s", err.Error())
			return
		}
	}
	if len(_glueAdditionalRunOptions) > 0 {
		if err := assignInputField(input, "AdditionalRunOptions", _glueAdditionalRunOptions); err != nil {
			log.Errorf("invalid --additional-run-options: %s", err.Error())
			return
		}
	}
	if len(_glueClientToken) > 0 {
		input.ClientToken = aws.String(_glueClientToken)
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDataQualityRulesetEvaluationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins an asynchronous task to export all labeled data for a particular
// transform. This task is the only label-related API call that is not part of the
// typical active learning workflow. You typically use StartExportLabelsTaskRun
// when you want to work with all of your existing labels at the same time, such as
// when you want to remove or change labels that were previously submitted as
// truth. This API operation accepts the TransformId whose labels you want to
// export and an Amazon Simple Storage Service (Amazon S3) path to export the
// labels to. The operation returns a TaskRunId . You can check on the status of
// your task run by calling the GetMLTaskRun API.
func glue_StartExportLabelsTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartExportLabelsTaskRunInput{
		// OutputS3Path: *string, // Required
		// TransformId: *string, // Required
	}

	if len(_glueOutputS3Path) > 0 {
		input.OutputS3Path = aws.String(_glueOutputS3Path)
	}
	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.StartExportLabelsTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to provide additional labels (examples of truth) to be used to
// teach the machine learning transform and improve its quality. This API operation
// is generally used as part of the active learning workflow that starts with the
// StartMLLabelingSetGenerationTaskRun call and that ultimately results in
// improving the quality of your machine learning transform.
//
// After the StartMLLabelingSetGenerationTaskRun finishes, Glue machine learning
// will have generated a series of questions for humans to answer. (Answering these
// questions is often called 'labeling' in the machine learning workflows). In the
// case of the FindMatches transform, these questions are of the form, “What is
// the correct way to group these rows together into groups composed entirely of
// matching records?” After the labeling process is finished, users upload their
// answers/labels with a call to StartImportLabelsTaskRun . After
// StartImportLabelsTaskRun finishes, all future runs of the machine learning
// transform use the new and improved labels and perform a higher-quality
// transformation.
//
// By default, StartMLLabelingSetGenerationTaskRun continually learns from and
// combines all labels that you upload unless you set Replace to true. If you set
// Replace to true, StartImportLabelsTaskRun deletes and forgets all previously
// uploaded labels and learns only from the exact set that you upload. Replacing
// labels can be helpful if you realize that you previously uploaded incorrect
// labels, and you believe that they are having a negative effect on your transform
// quality.
//
// You can check on the status of your task run by calling the GetMLTaskRun
// operation.
func glue_StartImportLabelsTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartImportLabelsTaskRunInput{
		// InputS3Path: *string, // Required
		// TransformId: *string, // Required
	}

	if len(_glueInputS3Path) > 0 {
		input.InputS3Path = aws.String(_glueInputS3Path)
	}
	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}
	if len(_glueReplaceAllLabels) > 0 {
		if err := assignInputField(input, "ReplaceAllLabels", _glueReplaceAllLabels); err != nil {
			log.Errorf("invalid --replace-all-labels: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImportLabelsTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a job run using a job definition.
func glue_StartJobRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartJobRunInput{
		// JobName: *string, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueAllocatedCapacity) > 0 {
		if err := assignInputField(input, "AllocatedCapacity", _glueAllocatedCapacity); err != nil {
			log.Errorf("invalid --allocated-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueArguments) > 0 {
		if err := assignInputField(input, "Arguments", _glueArguments); err != nil {
			log.Errorf("invalid --arguments: %s", err.Error())
			return
		}
	}
	if len(_glueExecutionClass) > 0 {
		if err := assignInputField(input, "ExecutionClass", _glueExecutionClass); err != nil {
			log.Errorf("invalid --execution-class: %s", err.Error())
			return
		}
	}
	if len(_glueExecutionRoleSessionPolicy) > 0 {
		input.ExecutionRoleSessionPolicy = aws.String(_glueExecutionRoleSessionPolicy)
	}
	if len(_glueJobRunId) > 0 {
		input.JobRunId = aws.String(_glueJobRunId)
	}
	if len(_glueJobRunQueuingEnabled) > 0 {
		if err := assignInputField(input, "JobRunQueuingEnabled", _glueJobRunQueuingEnabled); err != nil {
			log.Errorf("invalid --job-run-queuing-enabled: %s", err.Error())
			return
		}
	}
	if len(_glueMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _glueMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueNotificationProperty) > 0 {
		if err := assignInputField(input, "NotificationProperty", _glueNotificationProperty); err != nil {
			log.Errorf("invalid --notification-property: %s", err.Error())
			return
		}
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_glueWorkerType) > 0 {
		if err := assignInputField(input, "WorkerType", _glueWorkerType); err != nil {
			log.Errorf("invalid --worker-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a materialized view refresh task run, for a specified table and columns.
func glue_StartMaterializedViewRefreshTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartMaterializedViewRefreshTaskRunInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueFullRefresh) > 0 {
		if err := assignInputField(input, "FullRefresh", _glueFullRefresh); err != nil {
			log.Errorf("invalid --full-refresh: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMaterializedViewRefreshTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a task to estimate the quality of the transform.
// When you provide label sets as examples of truth, Glue machine learning uses
// some of those examples to learn from them. The rest of the labels are used as a
// test to estimate quality.
//
// Returns a unique identifier for the run. You can call GetMLTaskRun to get more
// information about the stats of the EvaluationTaskRun .
func glue_StartMLEvaluationTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartMLEvaluationTaskRunInput{
		// TransformId: *string, // Required
	}

	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.StartMLEvaluationTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the active learning workflow for your machine learning transform to
// improve the transform's quality by generating label sets and adding labels.
//
// When the StartMLLabelingSetGenerationTaskRun finishes, Glue will have generated
// a "labeling set" or a set of questions for humans to answer.
//
// In the case of the FindMatches transform, these questions are of the form,
// “What is the correct way to group these rows together into groups composed
// entirely of matching records?”
//
// After the labeling process is finished, you can upload your labels with a call
// to StartImportLabelsTaskRun . After StartImportLabelsTaskRun finishes, all
// future runs of the machine learning transform will use the new and improved
// labels and perform a higher-quality transformation.
//
// Note: The role used to write the generated labeling set to the OutputS3Path is
// the role associated with the Machine Learning Transform, specified in the
// CreateMLTransform API.
func glue_StartMLLabelingSetGenerationTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartMLLabelingSetGenerationTaskRunInput{
		// OutputS3Path: *string, // Required
		// TransformId: *string, // Required
	}

	if len(_glueOutputS3Path) > 0 {
		input.OutputS3Path = aws.String(_glueOutputS3Path)
	}
	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}

	if resp, err := client.StartMLLabelingSetGenerationTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an existing trigger. See [Triggering Jobs] for information about how different types of
// trigger are started.
//
// [Triggering Jobs]: https://docs.aws.amazon.com/glue/latest/dg/trigger-job.html
func glue_StartTrigger(cfg aws.Config, client *glue.Client) {
	input := &glue.StartTriggerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.StartTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new run of the specified workflow.
func glue_StartWorkflowRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StartWorkflowRunInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRunProperties) > 0 {
		if err := assignInputField(input, "RunProperties", _glueRunProperties); err != nil {
			log.Errorf("invalid --run-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a task run for the specified table.
func glue_StopColumnStatisticsTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StopColumnStatisticsTaskRunInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.StopColumnStatisticsTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a column statistics task run schedule.
func glue_StopColumnStatisticsTaskRunSchedule(cfg aws.Config, client *glue.Client) {
	input := &glue.StopColumnStatisticsTaskRunScheduleInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.StopColumnStatisticsTaskRunSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If the specified crawler is running, stops the crawl.
func glue_StopCrawler(cfg aws.Config, client *glue.Client) {
	input := &glue.StopCrawlerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.StopCrawler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the schedule state of the specified crawler to NOT_SCHEDULED , but does not
// stop the crawler if it is already running.
func glue_StopCrawlerSchedule(cfg aws.Config, client *glue.Client) {
	input := &glue.StopCrawlerScheduleInput{
		// CrawlerName: *string, // Required
	}

	if len(_glueCrawlerName) > 0 {
		input.CrawlerName = aws.String(_glueCrawlerName)
	}

	if resp, err := client.StopCrawlerSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a materialized view refresh task run, for a specified table and columns.
func glue_StopMaterializedViewRefreshTaskRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StopMaterializedViewRefreshTaskRunInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}

	if resp, err := client.StopMaterializedViewRefreshTaskRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the session.
func glue_StopSession(cfg aws.Config, client *glue.Client) {
	input := &glue.StopSessionInput{
		// Id: *string, // Required
	}

	if len(_glueId) > 0 {
		input.Id = aws.String(_glueId)
	}
	if len(_glueRequestOrigin) > 0 {
		input.RequestOrigin = aws.String(_glueRequestOrigin)
	}

	if resp, err := client.StopSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a specified trigger.
func glue_StopTrigger(cfg aws.Config, client *glue.Client) {
	input := &glue.StopTriggerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}

	if resp, err := client.StopTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the execution of the specified workflow run.
func glue_StopWorkflowRun(cfg aws.Config, client *glue.Client) {
	input := &glue.StopWorkflowRunInput{
		// Name: *string, // Required
		// RunId: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueRunId) > 0 {
		input.RunId = aws.String(_glueRunId)
	}

	if resp, err := client.StopWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a resource. A tag is a label you can assign to an Amazon Web
// Services resource. In Glue, you can tag only certain resources. For information
// about what resources you can tag, see [Amazon Web Services Tags in Glue].
//
// [Amazon Web Services Tags in Glue]: https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html
func glue_TagResource(cfg aws.Config, client *glue.Client) {
	input := &glue.TagResourceInput{
		// ResourceArn: *string, // Required
		// TagsToAdd: map[string]string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueTagsToAdd) > 0 {
		if err := assignInputField(input, "TagsToAdd", _glueTagsToAdd); err != nil {
			log.Errorf("invalid --tags-to-add: %s", err.Error())
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

// Tests a connection to a service to validate the service credentials that you
// provide.
//
// You can either provide an existing connection name or a TestConnectionInput for
// testing a non-existing connection input. Providing both at the same time will
// cause an error.
//
// If the action is successful, the service sends back an HTTP 200 response.
func glue_TestConnection(cfg aws.Config, client *glue.Client) {
	input := &glue.TestConnectionInput{}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueConnectionName) > 0 {
		input.ConnectionName = aws.String(_glueConnectionName)
	}
	if len(_glueTestConnectionInput) > 0 {
		if err := assignInputField(input, "TestConnectionInput", _glueTestConnectionInput); err != nil {
			log.Errorf("invalid --test-connection-input: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a resource.
func glue_UntagResource(cfg aws.Config, client *glue.Client) {
	input := &glue.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagsToRemove: []string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueTagsToRemove) > 0 {
		input.TagsToRemove = append([]string(nil), _glueTagsToRemove...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a registered blueprint.
func glue_UpdateBlueprint(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateBlueprintInput{
		// BlueprintLocation: *string, // Required
		// Name: *string, // Required
	}

	if len(_glueBlueprintLocation) > 0 {
		input.BlueprintLocation = aws.String(_glueBlueprintLocation)
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}

	if resp, err := client.UpdateBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing catalog's properties in the Glue Data Catalog.
func glue_UpdateCatalog(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateCatalogInput{
		// CatalogId: *string, // Required
		// CatalogInput: *types.CatalogInput, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueCatalogInput) > 0 {
		if err := assignInputField(input, "CatalogInput", _glueCatalogInput); err != nil {
			log.Errorf("invalid --catalog-input: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing classifier (a GrokClassifier , an XMLClassifier , a
// JsonClassifier , or a CsvClassifier , depending on which field is present).
func glue_UpdateClassifier(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateClassifierInput{}

	if len(_glueCsvClassifier) > 0 {
		if err := assignInputField(input, "CsvClassifier", _glueCsvClassifier); err != nil {
			log.Errorf("invalid --csv-classifier: %s", err.Error())
			return
		}
	}
	if len(_glueGrokClassifier) > 0 {
		if err := assignInputField(input, "GrokClassifier", _glueGrokClassifier); err != nil {
			log.Errorf("invalid --grok-classifier: %s", err.Error())
			return
		}
	}
	if len(_glueJsonClassifier) > 0 {
		if err := assignInputField(input, "JsonClassifier", _glueJsonClassifier); err != nil {
			log.Errorf("invalid --json-classifier: %s", err.Error())
			return
		}
	}
	if len(_glueXMLClassifier) > 0 {
		if err := assignInputField(input, "XMLClassifier", _glueXMLClassifier); err != nil {
			log.Errorf("invalid --xml-classifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClassifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates partition statistics of columns.
// The Identity and Access Management (IAM) permission required for this operation
// is UpdatePartition .
func glue_UpdateColumnStatisticsForPartition(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateColumnStatisticsForPartitionInput{
		// ColumnStatisticsList: []types.ColumnStatistics, // Required
		// DatabaseName: *string, // Required
		// PartitionValues: []string, // Required
		// TableName: *string, // Required
	}

	if len(_glueColumnStatisticsList) > 0 {
		if err := assignInputField(input, "ColumnStatisticsList", _glueColumnStatisticsList); err != nil {
			log.Errorf("invalid --column-statistics-list: %s", err.Error())
			return
		}
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionValues) > 0 {
		input.PartitionValues = append([]string(nil), _gluePartitionValues...)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.UpdateColumnStatisticsForPartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates table statistics of columns.
// The Identity and Access Management (IAM) permission required for this operation
// is UpdateTable .
func glue_UpdateColumnStatisticsForTable(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateColumnStatisticsForTableInput{
		// ColumnStatisticsList: []types.ColumnStatistics, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueColumnStatisticsList) > 0 {
		if err := assignInputField(input, "ColumnStatisticsList", _glueColumnStatisticsList); err != nil {
			log.Errorf("invalid --column-statistics-list: %s", err.Error())
			return
		}
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.UpdateColumnStatisticsForTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates settings for a column statistics task.
func glue_UpdateColumnStatisticsTaskSettings(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateColumnStatisticsTaskSettingsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogID = aws.String(_glueCatalogId)
	}
	if len(_glueColumnNameList) > 0 {
		input.ColumnNameList = append([]string(nil), _glueColumnNameList...)
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueSampleSize) > 0 {
		if err := assignInputField(input, "SampleSize", _glueSampleSize); err != nil {
			log.Errorf("invalid --sample-size: %s", err.Error())
			return
		}
	}
	if len(_glueSchedule) > 0 {
		input.Schedule = aws.String(_glueSchedule)
	}
	if len(_glueSecurityConfiguration) > 0 {
		input.SecurityConfiguration = aws.String(_glueSecurityConfiguration)
	}

	if resp, err := client.UpdateColumnStatisticsTaskSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a connection definition in the Data Catalog.
func glue_UpdateConnection(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateConnectionInput{
		// ConnectionInput: *types.ConnectionInput, // Required
		// Name: *string, // Required
	}

	if len(_glueConnectionInput) > 0 {
		if err := assignInputField(input, "ConnectionInput", _glueConnectionInput); err != nil {
			log.Errorf("invalid --connection-input: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.UpdateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a crawler. If a crawler is running, you must stop it using StopCrawler
// before updating it.
func glue_UpdateCrawler(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateCrawlerInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueClassifiers) > 0 {
		input.Classifiers = append([]string(nil), _glueClassifiers...)
	}
	if len(_glueConfiguration) > 0 {
		input.Configuration = aws.String(_glueConfiguration)
	}
	if len(_glueCrawlerSecurityConfiguration) > 0 {
		input.CrawlerSecurityConfiguration = aws.String(_glueCrawlerSecurityConfiguration)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueLakeFormationConfiguration) > 0 {
		if err := assignInputField(input, "LakeFormationConfiguration", _glueLakeFormationConfiguration); err != nil {
			log.Errorf("invalid --lake-formation-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueLineageConfiguration) > 0 {
		if err := assignInputField(input, "LineageConfiguration", _glueLineageConfiguration); err != nil {
			log.Errorf("invalid --lineage-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueRecrawlPolicy) > 0 {
		if err := assignInputField(input, "RecrawlPolicy", _glueRecrawlPolicy); err != nil {
			log.Errorf("invalid --recrawl-policy: %s", err.Error())
			return
		}
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueSchedule) > 0 {
		input.Schedule = aws.String(_glueSchedule)
	}
	if len(_glueSchemaChangePolicy) > 0 {
		if err := assignInputField(input, "SchemaChangePolicy", _glueSchemaChangePolicy); err != nil {
			log.Errorf("invalid --schema-change-policy: %s", err.Error())
			return
		}
	}
	if len(_glueTablePrefix) > 0 {
		input.TablePrefix = aws.String(_glueTablePrefix)
	}
	if len(_glueTargets) > 0 {
		if err := assignInputField(input, "Targets", _glueTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCrawler(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the schedule of a crawler using a cron expression.
func glue_UpdateCrawlerSchedule(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateCrawlerScheduleInput{
		// CrawlerName: *string, // Required
	}

	if len(_glueCrawlerName) > 0 {
		input.CrawlerName = aws.String(_glueCrawlerName)
	}
	if len(_glueSchedule) > 0 {
		input.Schedule = aws.String(_glueSchedule)
	}

	if resp, err := client.UpdateCrawlerSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified data quality ruleset.
func glue_UpdateDataQualityRuleset(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateDataQualityRulesetInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueRuleset) > 0 {
		input.Ruleset = aws.String(_glueRuleset)
	}

	if resp, err := client.UpdateDataQualityRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing database definition in a Data Catalog.
func glue_UpdateDatabase(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateDatabaseInput{
		// DatabaseInput: *types.DatabaseInput, // Required
		// Name: *string, // Required
	}

	if len(_glueDatabaseInput) > 0 {
		if err := assignInputField(input, "DatabaseInput", _glueDatabaseInput); err != nil {
			log.Errorf("invalid --database-input: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.UpdateDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified development endpoint.
func glue_UpdateDevEndpoint(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateDevEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_glueEndpointName) > 0 {
		input.EndpointName = aws.String(_glueEndpointName)
	}
	if len(_glueAddArguments) > 0 {
		if err := assignInputField(input, "AddArguments", _glueAddArguments); err != nil {
			log.Errorf("invalid --add-arguments: %s", err.Error())
			return
		}
	}
	if len(_glueAddPublicKeys) > 0 {
		input.AddPublicKeys = append([]string(nil), _glueAddPublicKeys...)
	}
	if len(_glueCustomLibraries) > 0 {
		if err := assignInputField(input, "CustomLibraries", _glueCustomLibraries); err != nil {
			log.Errorf("invalid --custom-libraries: %s", err.Error())
			return
		}
	}
	if len(_glueDeleteArguments) > 0 {
		input.DeleteArguments = append([]string(nil), _glueDeleteArguments...)
	}
	if len(_glueDeletePublicKeys) > 0 {
		input.DeletePublicKeys = append([]string(nil), _glueDeletePublicKeys...)
	}
	if len(_gluePublicKey) > 0 {
		input.PublicKey = aws.String(_gluePublicKey)
	}
	if len(_glueUpdateEtlLibraries) > 0 {
		if err := assignInputField(input, "UpdateEtlLibraries", _glueUpdateEtlLibraries); err != nil {
			log.Errorf("invalid --update-etl-libraries: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDevEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the existing Glue Identity Center configuration, allowing modification
// of scopes and permissions for the integration.
func glue_UpdateGlueIdentityCenterConfiguration(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateGlueIdentityCenterConfigurationInput{}

	if len(_glueScopes) > 0 {
		input.Scopes = append([]string(nil), _glueScopes...)
	}
	if len(_glueUserBackgroundSessionsEnabled) > 0 {
		if err := assignInputField(input, "UserBackgroundSessionsEnabled", _glueUserBackgroundSessionsEnabled); err != nil {
			log.Errorf("invalid --user-background-sessions-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlueIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API can be used for updating the ResourceProperty of the Glue connection
// (for the source) or Glue database ARN (for the target). These properties can
// include the role to access the connection or database. Since the same resource
// can be used across multiple integrations, updating resource properties will
// impact all the integrations using it.
func glue_UpdateIntegrationResourceProperty(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateIntegrationResourcePropertyInput{
		// ResourceArn: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueSourceProcessingProperties) > 0 {
		if err := assignInputField(input, "SourceProcessingProperties", _glueSourceProcessingProperties); err != nil {
			log.Errorf("invalid --source-processing-properties: %s", err.Error())
			return
		}
	}
	if len(_glueTargetProcessingProperties) > 0 {
		if err := assignInputField(input, "TargetProcessingProperties", _glueTargetProcessingProperties); err != nil {
			log.Errorf("invalid --target-processing-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIntegrationResourceProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is used to provide optional override properties for the tables that
// need to be replicated. These properties can include properties for filtering and
// partitioning for the source and target tables. To set both source and target
// properties the same API need to be invoked with the Glue connection ARN as
// ResourceArn with SourceTableConfig , and the Glue database ARN as ResourceArn
// with TargetTableConfig respectively.
//
// The override will be reflected across all the integrations using same
// ResourceArn and source table.
func glue_UpdateIntegrationTableProperties(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateIntegrationTablePropertiesInput{
		// ResourceArn: *string, // Required
		// TableName: *string, // Required
	}

	if len(_glueResourceArn) > 0 {
		input.ResourceArn = aws.String(_glueResourceArn)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueSourceTableConfig) > 0 {
		if err := assignInputField(input, "SourceTableConfig", _glueSourceTableConfig); err != nil {
			log.Errorf("invalid --source-table-config: %s", err.Error())
			return
		}
	}
	if len(_glueTargetTableConfig) > 0 {
		if err := assignInputField(input, "TargetTableConfig", _glueTargetTableConfig); err != nil {
			log.Errorf("invalid --target-table-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIntegrationTableProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing job definition. The previous job definition is completely
// overwritten by this information.
func glue_UpdateJob(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateJobInput{
		// JobName: *string, // Required
		// JobUpdate: *types.JobUpdate, // Required
	}

	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueJobUpdate) > 0 {
		if err := assignInputField(input, "JobUpdate", _glueJobUpdate); err != nil {
			log.Errorf("invalid --job-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Synchronizes a job from the source control repository. This operation takes the
// job artifacts that are located in the remote repository and updates the Glue
// internal stores with these artifacts.
//
// This API supports optional parameters which take in the repository information.
func glue_UpdateJobFromSourceControl(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateJobFromSourceControlInput{}

	if len(_glueAuthStrategy) > 0 {
		if err := assignInputField(input, "AuthStrategy", _glueAuthStrategy); err != nil {
			log.Errorf("invalid --auth-strategy: %s", err.Error())
			return
		}
	}
	if len(_glueAuthToken) > 0 {
		input.AuthToken = aws.String(_glueAuthToken)
	}
	if len(_glueBranchName) > 0 {
		input.BranchName = aws.String(_glueBranchName)
	}
	if len(_glueCommitId) > 0 {
		input.CommitId = aws.String(_glueCommitId)
	}
	if len(_glueFolder) > 0 {
		input.Folder = aws.String(_glueFolder)
	}
	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueProvider) > 0 {
		if err := assignInputField(input, "Provider", _glueProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}
	if len(_glueRepositoryName) > 0 {
		input.RepositoryName = aws.String(_glueRepositoryName)
	}
	if len(_glueRepositoryOwner) > 0 {
		input.RepositoryOwner = aws.String(_glueRepositoryOwner)
	}

	if resp, err := client.UpdateJobFromSourceControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing machine learning transform. Call this operation to tune the
// algorithm parameters to achieve better results.
//
// After calling this operation, you can call the StartMLEvaluationTaskRun
// operation to assess how well your new parameters achieved your goals (such as
// improving the quality of your machine learning transform, or making it more
// cost-effective).
func glue_UpdateMLTransform(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateMLTransformInput{
		// TransformId: *string, // Required
	}

	if len(_glueTransformId) > 0 {
		input.TransformId = aws.String(_glueTransformId)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueGlueVersion) > 0 {
		input.GlueVersion = aws.String(_glueGlueVersion)
	}
	if len(_glueMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _glueMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_glueMaxRetries) > 0 {
		if err := assignInputField(input, "MaxRetries", _glueMaxRetries); err != nil {
			log.Errorf("invalid --max-retries: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueNumberOfWorkers) > 0 {
		if err := assignInputField(input, "NumberOfWorkers", _glueNumberOfWorkers); err != nil {
			log.Errorf("invalid --number-of-workers: %s", err.Error())
			return
		}
	}
	if len(_glueParameters) > 0 {
		if err := assignInputField(input, "Parameters", _glueParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_glueRole) > 0 {
		input.Role = aws.String(_glueRole)
	}
	if len(_glueTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _glueTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_glueWorkerType) > 0 {
		if err := assignInputField(input, "WorkerType", _glueWorkerType); err != nil {
			log.Errorf("invalid --worker-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMLTransform(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a partition.
func glue_UpdatePartition(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdatePartitionInput{
		// DatabaseName: *string, // Required
		// PartitionInput: *types.PartitionInput, // Required
		// PartitionValueList: []string, // Required
		// TableName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_gluePartitionInput) > 0 {
		if err := assignInputField(input, "PartitionInput", _gluePartitionInput); err != nil {
			log.Errorf("invalid --partition-input: %s", err.Error())
			return
		}
	}
	if len(_gluePartitionValueList) > 0 {
		input.PartitionValueList = append([]string(nil), _gluePartitionValueList...)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.UpdatePartition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing registry which is used to hold a collection of schemas. The
// updated properties relate to the registry, and do not modify any of the schemas
// within the registry.
func glue_UpdateRegistry(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateRegistryInput{
		// Description: *string, // Required
		// RegistryId: *types.RegistryId, // Required
	}

	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueRegistryId) > 0 {
		if err := assignInputField(input, "RegistryId", _glueRegistryId); err != nil {
			log.Errorf("invalid --registry-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description, compatibility setting, or version checkpoint for a
// schema set.
//
// For updating the compatibility setting, the call will not validate
// compatibility for the entire set of schema versions with the new compatibility
// setting. If the value for Compatibility is provided, the VersionNumber (a
// checkpoint) is also required. The API will validate the checkpoint version
// number for consistency.
//
// If the value for the VersionNumber (checkpoint) is provided, Compatibility is
// optional and this can be used to set/reset a checkpoint for the schema.
//
// This update will happen only if the schema is in the AVAILABLE state.
func glue_UpdateSchema(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateSchemaInput{
		// SchemaId: *types.SchemaId, // Required
	}

	if len(_glueSchemaId) > 0 {
		if err := assignInputField(input, "SchemaId", _glueSchemaId); err != nil {
			log.Errorf("invalid --schema-id: %s", err.Error())
			return
		}
	}
	if len(_glueCompatibility) > 0 {
		if err := assignInputField(input, "Compatibility", _glueCompatibility); err != nil {
			log.Errorf("invalid --compatibility: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueSchemaVersionNumber) > 0 {
		if err := assignInputField(input, "SchemaVersionNumber", _glueSchemaVersionNumber); err != nil {
			log.Errorf("invalid --schema-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Synchronizes a job to the source control repository. This operation takes the
// job artifacts from the Glue internal stores and makes a commit to the remote
// repository that is configured on the job.
//
// This API supports optional parameters which take in the repository information.
func glue_UpdateSourceControlFromJob(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateSourceControlFromJobInput{}

	if len(_glueAuthStrategy) > 0 {
		if err := assignInputField(input, "AuthStrategy", _glueAuthStrategy); err != nil {
			log.Errorf("invalid --auth-strategy: %s", err.Error())
			return
		}
	}
	if len(_glueAuthToken) > 0 {
		input.AuthToken = aws.String(_glueAuthToken)
	}
	if len(_glueBranchName) > 0 {
		input.BranchName = aws.String(_glueBranchName)
	}
	if len(_glueCommitId) > 0 {
		input.CommitId = aws.String(_glueCommitId)
	}
	if len(_glueFolder) > 0 {
		input.Folder = aws.String(_glueFolder)
	}
	if len(_glueJobName) > 0 {
		input.JobName = aws.String(_glueJobName)
	}
	if len(_glueProvider) > 0 {
		if err := assignInputField(input, "Provider", _glueProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}
	if len(_glueRepositoryName) > 0 {
		input.RepositoryName = aws.String(_glueRepositoryName)
	}
	if len(_glueRepositoryOwner) > 0 {
		input.RepositoryOwner = aws.String(_glueRepositoryOwner)
	}

	if resp, err := client.UpdateSourceControlFromJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a metadata table in the Data Catalog.
func glue_UpdateTable(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateTableInput{
		// DatabaseName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueForce) > 0 {
		if err := assignInputField(input, "Force", _glueForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueSkipArchive) > 0 {
		if err := assignInputField(input, "SkipArchive", _glueSkipArchive); err != nil {
			log.Errorf("invalid --skip-archive: %s", err.Error())
			return
		}
	}
	if len(_glueTableInput) > 0 {
		if err := assignInputField(input, "TableInput", _glueTableInput); err != nil {
			log.Errorf("invalid --table-input: %s", err.Error())
			return
		}
	}
	if len(_glueTransactionId) > 0 {
		input.TransactionId = aws.String(_glueTransactionId)
	}
	if len(_glueUpdateOpenTableFormatInput) > 0 {
		if err := assignInputField(input, "UpdateOpenTableFormatInput", _glueUpdateOpenTableFormatInput); err != nil {
			log.Errorf("invalid --update-open-table-format-input: %s", err.Error())
			return
		}
	}
	if len(_glueVersionId) > 0 {
		input.VersionId = aws.String(_glueVersionId)
	}
	if len(_glueViewUpdateAction) > 0 {
		if err := assignInputField(input, "ViewUpdateAction", _glueViewUpdateAction); err != nil {
			log.Errorf("invalid --view-update-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for an existing table optimizer.
func glue_UpdateTableOptimizer(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateTableOptimizerInput{
		// CatalogId: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// TableOptimizerConfiguration: *types.TableOptimizerConfiguration, // Required
		// Type: types.TableOptimizerType, // Required
	}

	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}
	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueTableName) > 0 {
		input.TableName = aws.String(_glueTableName)
	}
	if len(_glueTableOptimizerConfiguration) > 0 {
		if err := assignInputField(input, "TableOptimizerConfiguration", _glueTableOptimizerConfiguration); err != nil {
			log.Errorf("invalid --table-optimizer-configuration: %s", err.Error())
			return
		}
	}
	if len(_glueType) > 0 {
		if err := assignInputField(input, "Type", _glueType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTableOptimizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a trigger definition.
// Job arguments may be logged. Do not pass plaintext secrets as arguments.
// Retrieve secrets from a Glue Connection, Amazon Web Services Secrets Manager or
// other secret management mechanism if you intend to keep them within the Job.
func glue_UpdateTrigger(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateTriggerInput{
		// Name: *string, // Required
		// TriggerUpdate: *types.TriggerUpdate, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueTriggerUpdate) > 0 {
		if err := assignInputField(input, "TriggerUpdate", _glueTriggerUpdate); err != nil {
			log.Errorf("invalid --trigger-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an Glue usage profile.
func glue_UpdateUsageProfile(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateUsageProfileInput{
		// Configuration: *types.ProfileConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_glueConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _glueConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}

	if resp, err := client.UpdateUsageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing function definition in the Data Catalog.
func glue_UpdateUserDefinedFunction(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateUserDefinedFunctionInput{
		// DatabaseName: *string, // Required
		// FunctionInput: *types.UserDefinedFunctionInput, // Required
		// FunctionName: *string, // Required
	}

	if len(_glueDatabaseName) > 0 {
		input.DatabaseName = aws.String(_glueDatabaseName)
	}
	if len(_glueFunctionInput) > 0 {
		if err := assignInputField(input, "FunctionInput", _glueFunctionInput); err != nil {
			log.Errorf("invalid --function-input: %s", err.Error())
			return
		}
	}
	if len(_glueFunctionName) > 0 {
		input.FunctionName = aws.String(_glueFunctionName)
	}
	if len(_glueCatalogId) > 0 {
		input.CatalogId = aws.String(_glueCatalogId)
	}

	if resp, err := client.UpdateUserDefinedFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing workflow.
func glue_UpdateWorkflow(cfg aws.Config, client *glue.Client) {
	input := &glue.UpdateWorkflowInput{
		// Name: *string, // Required
	}

	if len(_glueName) > 0 {
		input.Name = aws.String(_glueName)
	}
	if len(_glueDefaultRunProperties) > 0 {
		if err := assignInputField(input, "DefaultRunProperties", _glueDefaultRunProperties); err != nil {
			log.Errorf("invalid --default-run-properties: %s", err.Error())
			return
		}
	}
	if len(_glueDescription) > 0 {
		input.Description = aws.String(_glueDescription)
	}
	if len(_glueMaxConcurrentRuns) > 0 {
		if err := assignInputField(input, "MaxConcurrentRuns", _glueMaxConcurrentRuns); err != nil {
			log.Errorf("invalid --max-concurrent-runs: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_glueCmd)
	_glueCmd.Flags().SortFlags = false

	_glueCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_glueCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_glueCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_glueCmd.Flags().StringVarP(&_glueActions, "actions", "", "", "Actions")
	_glueCmd.Flags().StringVarP(&_glueAddArguments, "add-arguments", "", "", "Add Arguments")
	_glueCmd.Flags().StringSliceVarP(&_glueAddPublicKeys, "add-public-keys", "", nil, "Add Public Keys")
	_glueCmd.Flags().StringVarP(&_glueAdditionalDataSources, "additional-data-sources", "", "", "Additional Data Sources")
	_glueCmd.Flags().StringVarP(&_glueAdditionalEncryptionContext, "additional-encryption-context", "", "", "Additional Encryption Context")
	_glueCmd.Flags().StringVarP(&_glueAdditionalPlanOptionsMap, "additional-plan-options-map", "", "", "Additional Plan Options Map")
	_glueCmd.Flags().StringVarP(&_glueAdditionalRunOptions, "additional-run-options", "", "", "Additional Run Options")
	_glueCmd.Flags().StringVarP(&_glueAllocatedCapacity, "allocated-capacity", "", "", "Allocated Capacity")
	_glueCmd.Flags().StringVarP(&_glueApplyOverrideForComputeEnvironment, "apply-override-for-compute-environment", "", "", "Apply Override For Compute Environment")
	_glueCmd.Flags().StringVarP(&_glueArguments, "arguments", "", "", "Arguments")
	_glueCmd.Flags().StringVarP(&_glueAttributesToGet, "attributes-to-get", "", "", "Attributes To Get")
	_glueCmd.Flags().StringVarP(&_glueAuditContext, "audit-context", "", "", "Audit Context")
	_glueCmd.Flags().StringVarP(&_glueAuthStrategy, "auth-strategy", "", "", "Auth Strategy")
	_glueCmd.Flags().StringVarP(&_glueAuthToken, "auth-token", "", "", "Auth Token")
	_glueCmd.Flags().StringVarP(&_glueBlueprintLocation, "blueprint-location", "", "", "Blueprint Location")
	_glueCmd.Flags().StringVarP(&_glueBlueprintName, "blueprint-name", "", "", "Blueprint Name")
	_glueCmd.Flags().StringVarP(&_glueBranchName, "branch-name", "", "", "Branch Name")
	_glueCmd.Flags().StringVarP(&_glueCatalogId, "catalog-id", "", "", "Catalog ID")
	_glueCmd.Flags().StringVarP(&_glueCatalogInput, "catalog-input", "", "", "Catalog Input")
	_glueCmd.Flags().StringSliceVarP(&_glueClassifiers, "classifiers", "", nil, "Classifiers")
	_glueCmd.Flags().StringVarP(&_glueClientToken, "client-token", "", "", "Client Token")
	_glueCmd.Flags().StringVarP(&_glueCode, "code", "", "", "Code")
	_glueCmd.Flags().StringVarP(&_glueCodeGenConfigurationNodes, "code-gen-configuration-nodes", "", "", "Code Gen Configuration Nodes")
	_glueCmd.Flags().StringVarP(&_glueColumnName, "column-name", "", "", "Column Name")
	_glueCmd.Flags().StringSliceVarP(&_glueColumnNameList, "column-name-list", "", nil, "Column Name List")
	_glueCmd.Flags().StringSliceVarP(&_glueColumnNames, "column-names", "", nil, "Column Names")
	_glueCmd.Flags().StringVarP(&_glueColumnStatisticsList, "column-statistics-list", "", "", "Column Statistics List")
	_glueCmd.Flags().StringVarP(&_glueColumnStatisticsTaskRunId, "column-statistics-task-run-id", "", "", "Column Statistics Task Run ID")
	_glueCmd.Flags().StringVarP(&_glueCommand, "command", "", "", "Command")
	_glueCmd.Flags().StringVarP(&_glueCommitId, "commit-id", "", "", "Commit ID")
	_glueCmd.Flags().StringVarP(&_glueCompatibility, "compatibility", "", "", "Compatibility")
	_glueCmd.Flags().StringVarP(&_glueConfiguration, "configuration", "", "", "Configuration")
	_glueCmd.Flags().StringVarP(&_glueConnectionInput, "connection-input", "", "", "Connection Input")
	_glueCmd.Flags().StringVarP(&_glueConnectionName, "connection-name", "", "", "Connection Name")
	_glueCmd.Flags().StringSliceVarP(&_glueConnectionNameList, "connection-name-list", "", nil, "Connection Name List")
	_glueCmd.Flags().StringVarP(&_glueConnectionOptions, "connection-options", "", "", "Connection Options")
	_glueCmd.Flags().StringVarP(&_glueConnectionProperties, "connection-properties", "", "", "Connection Properties")
	_glueCmd.Flags().StringVarP(&_glueConnectionType, "connection-type", "", "", "Connection Type")
	_glueCmd.Flags().StringVarP(&_glueConnections, "connections", "", "", "Connections")
	_glueCmd.Flags().StringVarP(&_glueConnectorAuthenticationConfiguration, "connector-authentication-configuration", "", "", "Connector Authentication Configuration")
	_glueCmd.Flags().StringSliceVarP(&_glueContextWords, "context-words", "", nil, "Context Words")
	_glueCmd.Flags().StringVarP(&_glueCrawlerName, "crawler-name", "", "", "Crawler Name")
	_glueCmd.Flags().StringSliceVarP(&_glueCrawlerNameList, "crawler-name-list", "", nil, "Crawler Name List")
	_glueCmd.Flags().StringSliceVarP(&_glueCrawlerNames, "crawler-names", "", nil, "Crawler Names")
	_glueCmd.Flags().StringVarP(&_glueCrawlerSecurityConfiguration, "crawler-security-configuration", "", "", "Crawler Security Configuration")
	_glueCmd.Flags().StringVarP(&_glueCreatedRulesetName, "created-ruleset-name", "", "", "Created Ruleset Name")
	_glueCmd.Flags().StringVarP(&_glueCsvClassifier, "csv-classifier", "", "", "CSV Classifier")
	_glueCmd.Flags().StringVarP(&_glueCustomLibraries, "custom-libraries", "", "", "Custom Libraries")
	_glueCmd.Flags().StringVarP(&_glueDagEdges, "dag-edges", "", "", "Dag Edges")
	_glueCmd.Flags().StringVarP(&_glueDagNodes, "dag-nodes", "", "", "Dag Nodes")
	_glueCmd.Flags().StringVarP(&_glueDataCatalogEncryptionSettings, "data-catalog-encryption-settings", "", "", "Data Catalog Encryption Settings")
	_glueCmd.Flags().StringVarP(&_glueDataFilter, "data-filter", "", "", "Data Filter")
	_glueCmd.Flags().StringVarP(&_glueDataFormat, "data-format", "", "", "Data Format")
	_glueCmd.Flags().StringVarP(&_glueDataQualitySecurityConfiguration, "data-quality-security-configuration", "", "", "Data Quality Security Configuration")
	_glueCmd.Flags().StringVarP(&_glueDataSource, "data-source", "", "", "Data Source")
	_glueCmd.Flags().StringVarP(&_glueDataStoreApiVersion, "data-store-api-version", "", "", "Data Store API Version")
	_glueCmd.Flags().StringVarP(&_glueDatabaseInput, "database-input", "", "", "Database Input")
	_glueCmd.Flags().StringVarP(&_glueDatabaseName, "database-name", "", "", "Database Name")
	_glueCmd.Flags().StringVarP(&_glueDefaultArguments, "default-arguments", "", "", "Default Arguments")
	_glueCmd.Flags().StringVarP(&_glueDefaultRunProperties, "default-run-properties", "", "", "Default Run Properties")
	_glueCmd.Flags().StringSliceVarP(&_glueDeleteArguments, "delete-arguments", "", nil, "Delete Arguments")
	_glueCmd.Flags().StringSliceVarP(&_glueDeletePublicKeys, "delete-public-keys", "", nil, "Delete Public Keys")
	_glueCmd.Flags().StringVarP(&_glueDependentJobName, "dependent-job-name", "", "", "Dependent Job Name")
	_glueCmd.Flags().StringVarP(&_glueDescription, "description", "", "", "Description")
	_glueCmd.Flags().StringSliceVarP(&_glueDevEndpointNames, "dev-endpoint-names", "", nil, "Dev Endpoint Names")
	_glueCmd.Flags().StringVarP(&_glueEnableHybrid, "enable-hybrid", "", "", "Enable Hybrid")
	_glueCmd.Flags().StringVarP(&_glueEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_glueCmd.Flags().StringVarP(&_glueEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_glueCmd.Flags().StringVarP(&_glueEntityName, "entity-name", "", "", "Entity Name")
	_glueCmd.Flags().StringVarP(&_glueEntries, "entries", "", "", "Entries")
	_glueCmd.Flags().StringVarP(&_glueEventBatchingCondition, "event-batching-condition", "", "", "Event Batching Condition")
	_glueCmd.Flags().StringVarP(&_glueExcludeColumnSchema, "exclude-column-schema", "", "", "Exclude Column Schema")
	_glueCmd.Flags().StringVarP(&_glueExecutionClass, "execution-class", "", "", "Execution Class")
	_glueCmd.Flags().StringVarP(&_glueExecutionProperty, "execution-property", "", "", "Execution Property")
	_glueCmd.Flags().StringVarP(&_glueExecutionRoleSessionPolicy, "execution-role-session-policy", "", "", "Execution Role Session Policy")
	_glueCmd.Flags().StringVarP(&_glueExpression, "expression", "", "", "Expression")
	_glueCmd.Flags().StringVarP(&_glueExtraJarsS3Path, "extra-jars-s3-path", "", "", "Extra Jars S3 Path")
	_glueCmd.Flags().StringVarP(&_glueExtraPythonLibsS3Path, "extra-python-libs-s3-path", "", "", "Extra Python Libs S3 Path")
	_glueCmd.Flags().StringVarP(&_glueFilter, "filter", "", "", "Filter")
	_glueCmd.Flags().StringVarP(&_glueFilterPredicate, "filter-predicate", "", "", "Filter Predicate")
	_glueCmd.Flags().StringVarP(&_glueFilters, "filters", "", "", "Filters")
	_glueCmd.Flags().StringVarP(&_glueFirstSchemaVersionNumber, "first-schema-version-number", "", "", "First Schema Version Number")
	_glueCmd.Flags().StringVarP(&_glueFolder, "folder", "", "", "Folder")
	_glueCmd.Flags().StringVarP(&_glueForce, "force", "", "", "Force")
	_glueCmd.Flags().StringVarP(&_glueFullRefresh, "full-refresh", "", "", "Full Refresh")
	_glueCmd.Flags().StringVarP(&_glueFunctionInput, "function-input", "", "", "Function Input")
	_glueCmd.Flags().StringVarP(&_glueFunctionName, "function-name", "", "", "Function Name")
	_glueCmd.Flags().StringVarP(&_glueFunctionType, "function-type", "", "", "Function Type")
	_glueCmd.Flags().StringVarP(&_glueGlueVersion, "glue-version", "", "", "Glue Version")
	_glueCmd.Flags().StringVarP(&_glueGrokClassifier, "grok-classifier", "", "", "Grok Classifier")
	_glueCmd.Flags().StringVarP(&_glueHidePassword, "hide-password", "", "", "Hide Password")
	_glueCmd.Flags().StringVarP(&_glueId, "id", "", "", "ID")
	_glueCmd.Flags().StringVarP(&_glueIdleTimeout, "idle-timeout", "", "", "Idle Timeout")
	_glueCmd.Flags().StringVarP(&_glueIncludeBlueprint, "include-blueprint", "", "", "Include Blueprint")
	_glueCmd.Flags().StringVarP(&_glueIncludeGraph, "include-graph", "", "", "Include Graph")
	_glueCmd.Flags().StringVarP(&_glueIncludeParameterSpec, "include-parameter-spec", "", "", "Include Parameter Spec")
	_glueCmd.Flags().StringVarP(&_glueIncludeRoot, "include-root", "", "", "Include Root")
	_glueCmd.Flags().StringVarP(&_glueIncludeStatusDetails, "include-status-details", "", "", "Include Status Details")
	_glueCmd.Flags().StringVarP(&_glueInclusionAnnotation, "inclusion-annotation", "", "", "Inclusion Annotation")
	_glueCmd.Flags().StringVarP(&_glueInclusionAnnotations, "inclusion-annotations", "", "", "Inclusion Annotations")
	_glueCmd.Flags().StringVarP(&_glueIndexName, "index-name", "", "", "Index Name")
	_glueCmd.Flags().StringVarP(&_glueInputRecordTables, "input-record-tables", "", "", "Input Record Tables")
	_glueCmd.Flags().StringVarP(&_glueInputS3Path, "input-s3-path", "", "", "Input S3 Path")
	_glueCmd.Flags().StringVarP(&_glueInstanceArn, "instance-arn", "", "", "Instance ARN")
	_glueCmd.Flags().StringVarP(&_glueIntegrationArn, "integration-arn", "", "", "Integration ARN")
	_glueCmd.Flags().StringVarP(&_glueIntegrationConfig, "integration-config", "", "", "Integration Config")
	_glueCmd.Flags().StringVarP(&_glueIntegrationIdentifier, "integration-identifier", "", "", "Integration Identifier")
	_glueCmd.Flags().StringVarP(&_glueIntegrationName, "integration-name", "", "", "Integration Name")
	_glueCmd.Flags().StringVarP(&_glueIntegrationType, "integration-type", "", "", "Integration Type")
	_glueCmd.Flags().StringVarP(&_glueJobMode, "job-mode", "", "", "Job Mode")
	_glueCmd.Flags().StringVarP(&_glueJobName, "job-name", "", "", "Job Name")
	_glueCmd.Flags().StringSliceVarP(&_glueJobNames, "job-names", "", nil, "Job Names")
	_glueCmd.Flags().StringVarP(&_glueJobRunId, "job-run-id", "", "", "Job Run ID")
	_glueCmd.Flags().StringSliceVarP(&_glueJobRunIds, "job-run-ids", "", nil, "Job Run Ids")
	_glueCmd.Flags().StringVarP(&_glueJobRunQueuingEnabled, "job-run-queuing-enabled", "", "", "Job Run Queuing Enabled")
	_glueCmd.Flags().StringVarP(&_glueJobUpdate, "job-update", "", "", "Job Update")
	_glueCmd.Flags().StringVarP(&_glueJsonClassifier, "json-classifier", "", "", "JSON Classifier")
	_glueCmd.Flags().StringVarP(&_glueKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_glueCmd.Flags().StringVarP(&_glueLakeFormationConfiguration, "lake-formation-configuration", "", "", "Lake Formation Configuration")
	_glueCmd.Flags().StringVarP(&_glueLanguage, "language", "", "", "Language")
	_glueCmd.Flags().StringVarP(&_glueLimit, "limit", "", "", "Limit")
	_glueCmd.Flags().StringVarP(&_glueLineageConfiguration, "lineage-configuration", "", "", "Lineage Configuration")
	_glueCmd.Flags().StringVarP(&_glueLocation, "location", "", "", "Location")
	_glueCmd.Flags().StringVarP(&_glueLogUri, "log-uri", "", "", "Log URI")
	_glueCmd.Flags().StringVarP(&_glueMaintenanceWindow, "maintenance-window", "", "", "Maintenance Window")
	_glueCmd.Flags().StringVarP(&_glueMapping, "mapping", "", "", "Mapping")
	_glueCmd.Flags().StringVarP(&_glueMarker, "marker", "", "", "Marker")
	_glueCmd.Flags().StringVarP(&_glueMaterializedViewRefreshTaskRunId, "materialized-view-refresh-task-run-id", "", "", "Materialized View Refresh Task Run ID")
	_glueCmd.Flags().StringVarP(&_glueMaxCapacity, "max-capacity", "", "", "Max Capacity")
	_glueCmd.Flags().StringVarP(&_glueMaxConcurrentRuns, "max-concurrent-runs", "", "", "Max Concurrent Runs")
	_glueCmd.Flags().StringVarP(&_glueMaxRecords, "max-records", "", "", "Max Records")
	_glueCmd.Flags().StringVarP(&_glueMaxResults, "max-results", "", "", "Max Results")
	_glueCmd.Flags().StringVarP(&_glueMaxRetries, "max-retries", "", "", "Max Retries")
	_glueCmd.Flags().StringVarP(&_glueMetadataKeyValue, "metadata-key-value", "", "", "Metadata Key Value")
	_glueCmd.Flags().StringVarP(&_glueMetadataList, "metadata-list", "", "", "Metadata List")
	_glueCmd.Flags().StringVarP(&_glueName, "name", "", "", "Name")
	_glueCmd.Flags().StringSliceVarP(&_glueNames, "names", "", nil, "Names")
	_glueCmd.Flags().StringVarP(&_glueNextToken, "next-token", "", "", "Next Token")
	_glueCmd.Flags().StringSliceVarP(&_glueNodeIds, "node-ids", "", nil, "Node Ids")
	_glueCmd.Flags().StringVarP(&_glueNonOverridableArguments, "non-overridable-arguments", "", "", "Non Overridable Arguments")
	_glueCmd.Flags().StringVarP(&_glueNotificationProperty, "notification-property", "", "", "Notification Property")
	_glueCmd.Flags().StringVarP(&_glueNumberOfNodes, "number-of-nodes", "", "", "Number Of Nodes")
	_glueCmd.Flags().StringVarP(&_glueNumberOfWorkers, "number-of-workers", "", "", "Number Of Workers")
	_glueCmd.Flags().StringVarP(&_glueOpenTableFormatInput, "open-table-format-input", "", "", "Open Table Format Input")
	_glueCmd.Flags().StringVarP(&_glueOrderBy, "order-by", "", "", "Order By")
	_glueCmd.Flags().StringVarP(&_glueOutputS3Path, "output-s3-path", "", "", "Output S3 Path")
	_glueCmd.Flags().StringVarP(&_glueParameters, "parameters", "", "", "Parameters")
	_glueCmd.Flags().StringVarP(&_glueParentCatalogId, "parent-catalog-id", "", "", "Parent Catalog ID")
	_glueCmd.Flags().StringVarP(&_glueParentEntityName, "parent-entity-name", "", "", "Parent Entity Name")
	_glueCmd.Flags().StringVarP(&_glueParentResourceArn, "parent-resource-arn", "", "", "Parent Resource ARN")
	_glueCmd.Flags().StringVarP(&_gluePartitionIndex, "partition-index", "", "", "Partition Index")
	_glueCmd.Flags().StringVarP(&_gluePartitionIndexes, "partition-indexes", "", "", "Partition Indexes")
	_glueCmd.Flags().StringVarP(&_gluePartitionInput, "partition-input", "", "", "Partition Input")
	_glueCmd.Flags().StringVarP(&_gluePartitionInputList, "partition-input-list", "", "", "Partition Input List")
	_glueCmd.Flags().StringSliceVarP(&_gluePartitionValueList, "partition-value-list", "", nil, "Partition Value List")
	_glueCmd.Flags().StringSliceVarP(&_gluePartitionValues, "partition-values", "", nil, "Partition Values")
	_glueCmd.Flags().StringVarP(&_gluePartitionsToDelete, "partitions-to-delete", "", "", "Partitions To Delete")
	_glueCmd.Flags().StringVarP(&_gluePartitionsToGet, "partitions-to-get", "", "", "Partitions To Get")
	_glueCmd.Flags().StringVarP(&_gluePattern, "pattern", "", "", "Pattern")
	_glueCmd.Flags().StringVarP(&_gluePermissions, "permissions", "", "", "Permissions")
	_glueCmd.Flags().StringVarP(&_gluePolicyExistsCondition, "policy-exists-condition", "", "", "Policy Exists Condition")
	_glueCmd.Flags().StringVarP(&_gluePolicyHashCondition, "policy-hash-condition", "", "", "Policy Hash Condition")
	_glueCmd.Flags().StringVarP(&_gluePolicyInJson, "policy-in-json", "", "", "Policy In JSON")
	_glueCmd.Flags().StringVarP(&_gluePredecessorsIncluded, "predecessors-included", "", "", "Predecessors Included")
	_glueCmd.Flags().StringVarP(&_gluePredicate, "predicate", "", "", "Predicate")
	_glueCmd.Flags().StringVarP(&_glueProfileId, "profile-id", "", "", "Profile ID")
	_glueCmd.Flags().StringVarP(&_glueProvider, "provider", "", "", "Provider")
	_glueCmd.Flags().StringVarP(&_gluePublicKey, "public-key", "", "", "Public Key")
	_glueCmd.Flags().StringSliceVarP(&_gluePublicKeys, "public-keys", "", nil, "Public Keys")
	_glueCmd.Flags().StringVarP(&_gluePythonScript, "python-script", "", "", "Python Script")
	_glueCmd.Flags().StringVarP(&_glueQueryAsOfTime, "query-as-of-time", "", "", "Query As Of Time")
	_glueCmd.Flags().StringVarP(&_glueQuerySessionContext, "query-session-context", "", "", "Query Session Context")
	_glueCmd.Flags().StringVarP(&_glueRecrawlPolicy, "recrawl-policy", "", "", "Recrawl Policy")
	_glueCmd.Flags().StringVarP(&_glueRecursive, "recursive", "", "", "Recursive")
	_glueCmd.Flags().StringVarP(&_glueRegexString, "regex-string", "", "", "Regex String")
	_glueCmd.Flags().StringVarP(&_glueRegistryId, "registry-id", "", "", "Registry ID")
	_glueCmd.Flags().StringVarP(&_glueRegistryName, "registry-name", "", "", "Registry Name")
	_glueCmd.Flags().StringVarP(&_glueReplaceAllLabels, "replace-all-labels", "", "", "Replace All Labels")
	_glueCmd.Flags().StringVarP(&_glueRepositoryName, "repository-name", "", "", "Repository Name")
	_glueCmd.Flags().StringVarP(&_glueRepositoryOwner, "repository-owner", "", "", "Repository Owner")
	_glueCmd.Flags().StringVarP(&_glueRequestOrigin, "request-origin", "", "", "Request Origin")
	_glueCmd.Flags().StringVarP(&_glueResourceArn, "resource-arn", "", "", "Resource ARN")
	_glueCmd.Flags().StringVarP(&_glueResourceShareType, "resource-share-type", "", "", "Resource Share Type")
	_glueCmd.Flags().StringVarP(&_glueRestConfiguration, "rest-configuration", "", "", "Rest Configuration")
	_glueCmd.Flags().StringVarP(&_glueResultId, "result-id", "", "", "Result ID")
	_glueCmd.Flags().StringSliceVarP(&_glueResultIds, "result-ids", "", nil, "Result Ids")
	_glueCmd.Flags().StringVarP(&_glueRole, "role", "", "", "Role")
	_glueCmd.Flags().StringVarP(&_glueRoleArn, "role-arn", "", "", "Role ARN")
	_glueCmd.Flags().StringVarP(&_glueRootResourceArn, "root-resource-arn", "", "", "Root Resource ARN")
	_glueCmd.Flags().StringVarP(&_glueRuleset, "ruleset", "", "", "Ruleset")
	_glueCmd.Flags().StringSliceVarP(&_glueRulesetNames, "ruleset-names", "", nil, "Ruleset Names")
	_glueCmd.Flags().StringVarP(&_glueRunId, "run-id", "", "", "Run ID")
	_glueCmd.Flags().StringVarP(&_glueRunProperties, "run-properties", "", "", "Run Properties")
	_glueCmd.Flags().StringVarP(&_glueSampleSize, "sample-size", "", "", "Sample Size")
	_glueCmd.Flags().StringVarP(&_glueSchedule, "schedule", "", "", "Schedule")
	_glueCmd.Flags().StringVarP(&_glueSchemaChangePolicy, "schema-change-policy", "", "", "Schema Change Policy")
	_glueCmd.Flags().StringVarP(&_glueSchemaDefinition, "schema-definition", "", "", "Schema Definition")
	_glueCmd.Flags().StringVarP(&_glueSchemaDiffType, "schema-diff-type", "", "", "Schema Diff Type")
	_glueCmd.Flags().StringVarP(&_glueSchemaId, "schema-id", "", "", "Schema ID")
	_glueCmd.Flags().StringVarP(&_glueSchemaName, "schema-name", "", "", "Schema Name")
	_glueCmd.Flags().StringVarP(&_glueSchemaVersionId, "schema-version-id", "", "", "Schema Version ID")
	_glueCmd.Flags().StringVarP(&_glueSchemaVersionNumber, "schema-version-number", "", "", "Schema Version Number")
	_glueCmd.Flags().StringSliceVarP(&_glueScopes, "scopes", "", nil, "Scopes")
	_glueCmd.Flags().StringVarP(&_glueSearchText, "search-text", "", "", "Search Text")
	_glueCmd.Flags().StringVarP(&_glueSecondSchemaVersionNumber, "second-schema-version-number", "", "", "Second Schema Version Number")
	_glueCmd.Flags().StringVarP(&_glueSecurityConfiguration, "security-configuration", "", "", "Security Configuration")
	_glueCmd.Flags().StringSliceVarP(&_glueSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_glueCmd.Flags().StringVarP(&_glueSegment, "segment", "", "", "Segment")
	_glueCmd.Flags().StringSliceVarP(&_glueSelectedFields, "selected-fields", "", nil, "Selected Fields")
	_glueCmd.Flags().StringVarP(&_glueSessionId, "session-id", "", "", "Session ID")
	_glueCmd.Flags().StringVarP(&_glueSinks, "sinks", "", "", "Sinks")
	_glueCmd.Flags().StringVarP(&_glueSkipArchive, "skip-archive", "", "", "Skip Archive")
	_glueCmd.Flags().StringVarP(&_glueSort, "sort", "", "", "Sort")
	_glueCmd.Flags().StringVarP(&_glueSortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_glueCmd.Flags().StringVarP(&_glueSource, "source", "", "", "Source")
	_glueCmd.Flags().StringVarP(&_glueSourceArn, "source-arn", "", "", "Source ARN")
	_glueCmd.Flags().StringVarP(&_glueSourceControlDetails, "source-control-details", "", "", "Source Control Details")
	_glueCmd.Flags().StringVarP(&_glueSourceProcessingProperties, "source-processing-properties", "", "", "Source Processing Properties")
	_glueCmd.Flags().StringVarP(&_glueSourceTableConfig, "source-table-config", "", "", "Source Table Config")
	_glueCmd.Flags().StringVarP(&_glueStartOnCreation, "start-on-creation", "", "", "Start On Creation")
	_glueCmd.Flags().StringVarP(&_glueStatisticId, "statistic-id", "", "", "Statistic ID")
	_glueCmd.Flags().StringVarP(&_glueSubnetId, "subnet-id", "", "", "Subnet ID")
	_glueCmd.Flags().StringVarP(&_glueSupportedDialect, "supported-dialect", "", "", "Supported Dialect")
	_glueCmd.Flags().StringVarP(&_glueSupportedPermissionTypes, "supported-permission-types", "", "", "Supported Permission Types")
	_glueCmd.Flags().StringVarP(&_glueTableInput, "table-input", "", "", "Table Input")
	_glueCmd.Flags().StringVarP(&_glueTableName, "table-name", "", "", "Table Name")
	_glueCmd.Flags().StringVarP(&_glueTableOptimizerConfiguration, "table-optimizer-configuration", "", "", "Table Optimizer Configuration")
	_glueCmd.Flags().StringVarP(&_glueTablePrefix, "table-prefix", "", "", "Table Prefix")
	_glueCmd.Flags().StringSliceVarP(&_glueTablesToDelete, "tables-to-delete", "", nil, "Tables To Delete")
	_glueCmd.Flags().StringVarP(&_glueTags, "tags", "", "", "Tags")
	_glueCmd.Flags().StringVarP(&_glueTagsToAdd, "tags-to-add", "", "", "Tags To Add")
	_glueCmd.Flags().StringSliceVarP(&_glueTagsToRemove, "tags-to-remove", "", nil, "Tags To Remove")
	_glueCmd.Flags().StringVarP(&_glueTargetArn, "target-arn", "", "", "Target ARN")
	_glueCmd.Flags().StringVarP(&_glueTargetProcessingProperties, "target-processing-properties", "", "", "Target Processing Properties")
	_glueCmd.Flags().StringVarP(&_glueTargetTable, "target-table", "", "", "Target Table")
	_glueCmd.Flags().StringVarP(&_glueTargetTableConfig, "target-table-config", "", "", "Target Table Config")
	_glueCmd.Flags().StringVarP(&_glueTargets, "targets", "", "", "Targets")
	_glueCmd.Flags().StringVarP(&_glueTaskRunId, "task-run-id", "", "", "Task Run ID")
	_glueCmd.Flags().StringVarP(&_glueTestConnectionInput, "test-connection-input", "", "", "Test Connection Input")
	_glueCmd.Flags().StringVarP(&_glueTimeout, "timeout", "", "", "Timeout")
	_glueCmd.Flags().StringVarP(&_glueTimestampFilter, "timestamp-filter", "", "", "Timestamp Filter")
	_glueCmd.Flags().StringVarP(&_glueTransactionId, "transaction-id", "", "", "Transaction ID")
	_glueCmd.Flags().StringVarP(&_glueTransformEncryption, "transform-encryption", "", "", "Transform Encryption")
	_glueCmd.Flags().StringVarP(&_glueTransformId, "transform-id", "", "", "Transform ID")
	_glueCmd.Flags().StringSliceVarP(&_glueTriggerNames, "trigger-names", "", nil, "Trigger Names")
	_glueCmd.Flags().StringVarP(&_glueTriggerUpdate, "trigger-update", "", "", "Trigger Update")
	_glueCmd.Flags().StringVarP(&_glueType, "type", "", "", "Type")
	_glueCmd.Flags().StringVarP(&_glueUpdateEtlLibraries, "update-etl-libraries", "", "", "Update Etl Libraries")
	_glueCmd.Flags().StringVarP(&_glueUpdateOpenTableFormatInput, "update-open-table-format-input", "", "", "Update Open Table Format Input")
	_glueCmd.Flags().StringVarP(&_glueUserBackgroundSessionsEnabled, "user-background-sessions-enabled", "", "", "User Background Sessions Enabled")
	_glueCmd.Flags().StringVarP(&_glueVersionId, "version-id", "", "", "Version ID")
	_glueCmd.Flags().StringSliceVarP(&_glueVersionIds, "version-ids", "", nil, "Version Ids")
	_glueCmd.Flags().StringVarP(&_glueVersions, "versions", "", "", "Versions")
	_glueCmd.Flags().StringVarP(&_glueViewUpdateAction, "view-update-action", "", "", "View Update Action")
	_glueCmd.Flags().StringVarP(&_glueWorkerType, "worker-type", "", "", "Worker Type")
	_glueCmd.Flags().StringVarP(&_glueWorkflowName, "workflow-name", "", "", "Workflow Name")
	_glueCmd.Flags().StringVarP(&_glueXMLClassifier, "xml-classifier", "", "", "XML Classifier")

	_glueCmd.Flags().BoolVarP(&_glueBatchCreatePartition, "batch-create-partition", "", false, "Batch Create Partition")
	_glueCmd.Flags().BoolVarP(&_glueBatchDeleteConnection, "batch-delete-connection", "", false, "Batch Delete Connection")
	_glueCmd.Flags().BoolVarP(&_glueBatchDeletePartition, "batch-delete-partition", "", false, "Batch Delete Partition")
	_glueCmd.Flags().BoolVarP(&_glueBatchDeleteTable, "batch-delete-table", "", false, "Batch Delete Table")
	_glueCmd.Flags().BoolVarP(&_glueBatchDeleteTableVersion, "batch-delete-table-version", "", false, "Batch Delete Table Version")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetBlueprints, "batch-get-blueprints", "", false, "Batch Get Blueprints")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetCrawlers, "batch-get-crawlers", "", false, "Batch Get Crawlers")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetCustomEntityTypes, "batch-get-custom-entity-types", "", false, "Batch Get Custom Entity Types")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetDataQualityResult, "batch-get-data-quality-result", "", false, "Batch Get Data Quality Result")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetDevEndpoints, "batch-get-dev-endpoints", "", false, "Batch Get Dev Endpoints")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetJobs, "batch-get-jobs", "", false, "Batch Get Jobs")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetPartition, "batch-get-partition", "", false, "Batch Get Partition")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetTableOptimizer, "batch-get-table-optimizer", "", false, "Batch Get Table Optimizer")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetTriggers, "batch-get-triggers", "", false, "Batch Get Triggers")
	_glueCmd.Flags().BoolVarP(&_glueBatchGetWorkflows, "batch-get-workflows", "", false, "Batch Get Workflows")
	_glueCmd.Flags().BoolVarP(&_glueBatchPutDataQualityStatisticAnnotation, "batch-put-data-quality-statistic-annotation", "", false, "Batch Put Data Quality Statistic Annotation")
	_glueCmd.Flags().BoolVarP(&_glueBatchStopJobRun, "batch-stop-job-run", "", false, "Batch Stop Job Run")
	_glueCmd.Flags().BoolVarP(&_glueBatchUpdatePartition, "batch-update-partition", "", false, "Batch Update Partition")
	_glueCmd.Flags().BoolVarP(&_glueCancelDataQualityRuleRecommendationRun, "cancel-data-quality-rule-recommendation-run", "", false, "Cancel Data Quality Rule Recommendation Run")
	_glueCmd.Flags().BoolVarP(&_glueCancelDataQualityRulesetEvaluationRun, "cancel-data-quality-ruleset-evaluation-run", "", false, "Cancel Data Quality Ruleset Evaluation Run")
	_glueCmd.Flags().BoolVarP(&_glueCancelMLTaskRun, "cancel-ml-task-run", "", false, "Cancel Ml Task Run")
	_glueCmd.Flags().BoolVarP(&_glueCancelStatement, "cancel-statement", "", false, "Cancel Statement")
	_glueCmd.Flags().BoolVarP(&_glueCheckSchemaVersionValidity, "check-schema-version-validity", "", false, "Check Schema Version Validity")
	_glueCmd.Flags().BoolVarP(&_glueCreateBlueprint, "create-blueprint", "", false, "Create Blueprint")
	_glueCmd.Flags().BoolVarP(&_glueCreateCatalog, "create-catalog", "", false, "Create Catalog")
	_glueCmd.Flags().BoolVarP(&_glueCreateClassifier, "create-classifier", "", false, "Create Classifier")
	_glueCmd.Flags().BoolVarP(&_glueCreateColumnStatisticsTaskSettings, "create-column-statistics-task-settings", "", false, "Create Column Statistics Task Settings")
	_glueCmd.Flags().BoolVarP(&_glueCreateConnection, "create-connection", "", false, "Create Connection")
	_glueCmd.Flags().BoolVarP(&_glueCreateCrawler, "create-crawler", "", false, "Create Crawler")
	_glueCmd.Flags().BoolVarP(&_glueCreateCustomEntityType, "create-custom-entity-type", "", false, "Create Custom Entity Type")
	_glueCmd.Flags().BoolVarP(&_glueCreateDataQualityRuleset, "create-data-quality-ruleset", "", false, "Create Data Quality Ruleset")
	_glueCmd.Flags().BoolVarP(&_glueCreateDatabase, "create-database", "", false, "Create Database")
	_glueCmd.Flags().BoolVarP(&_glueCreateDevEndpoint, "create-dev-endpoint", "", false, "Create Dev Endpoint")
	_glueCmd.Flags().BoolVarP(&_glueCreateGlueIdentityCenterConfiguration, "create-glue-identity-center-configuration", "", false, "Create Glue Identity Center Configuration")
	_glueCmd.Flags().BoolVarP(&_glueCreateIntegration, "create-integration", "", false, "Create Integration")
	_glueCmd.Flags().BoolVarP(&_glueCreateIntegrationResourceProperty, "create-integration-resource-property", "", false, "Create Integration Resource Property")
	_glueCmd.Flags().BoolVarP(&_glueCreateIntegrationTableProperties, "create-integration-table-properties", "", false, "Create Integration Table Properties")
	_glueCmd.Flags().BoolVarP(&_glueCreateJob, "create-job", "", false, "Create Job")
	_glueCmd.Flags().BoolVarP(&_glueCreateMLTransform, "create-ml-transform", "", false, "Create Ml Transform")
	_glueCmd.Flags().BoolVarP(&_glueCreatePartition, "create-partition", "", false, "Create Partition")
	_glueCmd.Flags().BoolVarP(&_glueCreatePartitionIndex, "create-partition-index", "", false, "Create Partition Index")
	_glueCmd.Flags().BoolVarP(&_glueCreateRegistry, "create-registry", "", false, "Create Registry")
	_glueCmd.Flags().BoolVarP(&_glueCreateSchema, "create-schema", "", false, "Create Schema")
	_glueCmd.Flags().BoolVarP(&_glueCreateScript, "create-script", "", false, "Create Script")
	_glueCmd.Flags().BoolVarP(&_glueCreateSecurityConfiguration, "create-security-configuration", "", false, "Create Security Configuration")
	_glueCmd.Flags().BoolVarP(&_glueCreateSession, "create-session", "", false, "Create Session")
	_glueCmd.Flags().BoolVarP(&_glueCreateTable, "create-table", "", false, "Create Table")
	_glueCmd.Flags().BoolVarP(&_glueCreateTableOptimizer, "create-table-optimizer", "", false, "Create Table Optimizer")
	_glueCmd.Flags().BoolVarP(&_glueCreateTrigger, "create-trigger", "", false, "Create Trigger")
	_glueCmd.Flags().BoolVarP(&_glueCreateUsageProfile, "create-usage-profile", "", false, "Create Usage Profile")
	_glueCmd.Flags().BoolVarP(&_glueCreateUserDefinedFunction, "create-user-defined-function", "", false, "Create User Defined Function")
	_glueCmd.Flags().BoolVarP(&_glueCreateWorkflow, "create-workflow", "", false, "Create Workflow")
	_glueCmd.Flags().BoolVarP(&_glueDeleteBlueprint, "delete-blueprint", "", false, "Delete Blueprint")
	_glueCmd.Flags().BoolVarP(&_glueDeleteCatalog, "delete-catalog", "", false, "Delete Catalog")
	_glueCmd.Flags().BoolVarP(&_glueDeleteClassifier, "delete-classifier", "", false, "Delete Classifier")
	_glueCmd.Flags().BoolVarP(&_glueDeleteColumnStatisticsForPartition, "delete-column-statistics-for-partition", "", false, "Delete Column Statistics For Partition")
	_glueCmd.Flags().BoolVarP(&_glueDeleteColumnStatisticsForTable, "delete-column-statistics-for-table", "", false, "Delete Column Statistics For Table")
	_glueCmd.Flags().BoolVarP(&_glueDeleteColumnStatisticsTaskSettings, "delete-column-statistics-task-settings", "", false, "Delete Column Statistics Task Settings")
	_glueCmd.Flags().BoolVarP(&_glueDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_glueCmd.Flags().BoolVarP(&_glueDeleteConnectionType, "delete-connection-type", "", false, "Delete Connection Type")
	_glueCmd.Flags().BoolVarP(&_glueDeleteCrawler, "delete-crawler", "", false, "Delete Crawler")
	_glueCmd.Flags().BoolVarP(&_glueDeleteCustomEntityType, "delete-custom-entity-type", "", false, "Delete Custom Entity Type")
	_glueCmd.Flags().BoolVarP(&_glueDeleteDataQualityRuleset, "delete-data-quality-ruleset", "", false, "Delete Data Quality Ruleset")
	_glueCmd.Flags().BoolVarP(&_glueDeleteDatabase, "delete-database", "", false, "Delete Database")
	_glueCmd.Flags().BoolVarP(&_glueDeleteDevEndpoint, "delete-dev-endpoint", "", false, "Delete Dev Endpoint")
	_glueCmd.Flags().BoolVarP(&_glueDeleteGlueIdentityCenterConfiguration, "delete-glue-identity-center-configuration", "", false, "Delete Glue Identity Center Configuration")
	_glueCmd.Flags().BoolVarP(&_glueDeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_glueCmd.Flags().BoolVarP(&_glueDeleteIntegrationResourceProperty, "delete-integration-resource-property", "", false, "Delete Integration Resource Property")
	_glueCmd.Flags().BoolVarP(&_glueDeleteIntegrationTableProperties, "delete-integration-table-properties", "", false, "Delete Integration Table Properties")
	_glueCmd.Flags().BoolVarP(&_glueDeleteJob, "delete-job", "", false, "Delete Job")
	_glueCmd.Flags().BoolVarP(&_glueDeleteMLTransform, "delete-ml-transform", "", false, "Delete Ml Transform")
	_glueCmd.Flags().BoolVarP(&_glueDeletePartition, "delete-partition", "", false, "Delete Partition")
	_glueCmd.Flags().BoolVarP(&_glueDeletePartitionIndex, "delete-partition-index", "", false, "Delete Partition Index")
	_glueCmd.Flags().BoolVarP(&_glueDeleteRegistry, "delete-registry", "", false, "Delete Registry")
	_glueCmd.Flags().BoolVarP(&_glueDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_glueCmd.Flags().BoolVarP(&_glueDeleteSchema, "delete-schema", "", false, "Delete Schema")
	_glueCmd.Flags().BoolVarP(&_glueDeleteSchemaVersions, "delete-schema-versions", "", false, "Delete Schema Versions")
	_glueCmd.Flags().BoolVarP(&_glueDeleteSecurityConfiguration, "delete-security-configuration", "", false, "Delete Security Configuration")
	_glueCmd.Flags().BoolVarP(&_glueDeleteSession, "delete-session", "", false, "Delete Session")
	_glueCmd.Flags().BoolVarP(&_glueDeleteTable, "delete-table", "", false, "Delete Table")
	_glueCmd.Flags().BoolVarP(&_glueDeleteTableOptimizer, "delete-table-optimizer", "", false, "Delete Table Optimizer")
	_glueCmd.Flags().BoolVarP(&_glueDeleteTableVersion, "delete-table-version", "", false, "Delete Table Version")
	_glueCmd.Flags().BoolVarP(&_glueDeleteTrigger, "delete-trigger", "", false, "Delete Trigger")
	_glueCmd.Flags().BoolVarP(&_glueDeleteUsageProfile, "delete-usage-profile", "", false, "Delete Usage Profile")
	_glueCmd.Flags().BoolVarP(&_glueDeleteUserDefinedFunction, "delete-user-defined-function", "", false, "Delete User Defined Function")
	_glueCmd.Flags().BoolVarP(&_glueDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_glueCmd.Flags().BoolVarP(&_glueDescribeConnectionType, "describe-connection-type", "", false, "Describe Connection Type")
	_glueCmd.Flags().BoolVarP(&_glueDescribeEntity, "describe-entity", "", false, "Describe Entity")
	_glueCmd.Flags().BoolVarP(&_glueDescribeInboundIntegrations, "describe-inbound-integrations", "", false, "Describe Inbound Integrations")
	_glueCmd.Flags().BoolVarP(&_glueDescribeIntegrations, "describe-integrations", "", false, "Describe Integrations")
	_glueCmd.Flags().BoolVarP(&_glueGetBlueprint, "get-blueprint", "", false, "Get Blueprint")
	_glueCmd.Flags().BoolVarP(&_glueGetBlueprintRun, "get-blueprint-run", "", false, "Get Blueprint Run")
	_glueCmd.Flags().BoolVarP(&_glueGetBlueprintRuns, "get-blueprint-runs", "", false, "Get Blueprint Runs")
	_glueCmd.Flags().BoolVarP(&_glueGetCatalog, "get-catalog", "", false, "Get Catalog")
	_glueCmd.Flags().BoolVarP(&_glueGetCatalogImportStatus, "get-catalog-import-status", "", false, "Get Catalog Import Status")
	_glueCmd.Flags().BoolVarP(&_glueGetCatalogs, "get-catalogs", "", false, "Get Catalogs")
	_glueCmd.Flags().BoolVarP(&_glueGetClassifier, "get-classifier", "", false, "Get Classifier")
	_glueCmd.Flags().BoolVarP(&_glueGetClassifiers, "get-classifiers", "", false, "Get Classifiers")
	_glueCmd.Flags().BoolVarP(&_glueGetColumnStatisticsForPartition, "get-column-statistics-for-partition", "", false, "Get Column Statistics For Partition")
	_glueCmd.Flags().BoolVarP(&_glueGetColumnStatisticsForTable, "get-column-statistics-for-table", "", false, "Get Column Statistics For Table")
	_glueCmd.Flags().BoolVarP(&_glueGetColumnStatisticsTaskRun, "get-column-statistics-task-run", "", false, "Get Column Statistics Task Run")
	_glueCmd.Flags().BoolVarP(&_glueGetColumnStatisticsTaskRuns, "get-column-statistics-task-runs", "", false, "Get Column Statistics Task Runs")
	_glueCmd.Flags().BoolVarP(&_glueGetColumnStatisticsTaskSettings, "get-column-statistics-task-settings", "", false, "Get Column Statistics Task Settings")
	_glueCmd.Flags().BoolVarP(&_glueGetConnection, "get-connection", "", false, "Get Connection")
	_glueCmd.Flags().BoolVarP(&_glueGetConnections, "get-connections", "", false, "Get Connections")
	_glueCmd.Flags().BoolVarP(&_glueGetCrawler, "get-crawler", "", false, "Get Crawler")
	_glueCmd.Flags().BoolVarP(&_glueGetCrawlerMetrics, "get-crawler-metrics", "", false, "Get Crawler Metrics")
	_glueCmd.Flags().BoolVarP(&_glueGetCrawlers, "get-crawlers", "", false, "Get Crawlers")
	_glueCmd.Flags().BoolVarP(&_glueGetCustomEntityType, "get-custom-entity-type", "", false, "Get Custom Entity Type")
	_glueCmd.Flags().BoolVarP(&_glueGetDataCatalogEncryptionSettings, "get-data-catalog-encryption-settings", "", false, "Get Data Catalog Encryption Settings")
	_glueCmd.Flags().BoolVarP(&_glueGetDataQualityModel, "get-data-quality-model", "", false, "Get Data Quality Model")
	_glueCmd.Flags().BoolVarP(&_glueGetDataQualityModelResult, "get-data-quality-model-result", "", false, "Get Data Quality Model Result")
	_glueCmd.Flags().BoolVarP(&_glueGetDataQualityResult, "get-data-quality-result", "", false, "Get Data Quality Result")
	_glueCmd.Flags().BoolVarP(&_glueGetDataQualityRuleRecommendationRun, "get-data-quality-rule-recommendation-run", "", false, "Get Data Quality Rule Recommendation Run")
	_glueCmd.Flags().BoolVarP(&_glueGetDataQualityRuleset, "get-data-quality-ruleset", "", false, "Get Data Quality Ruleset")
	_glueCmd.Flags().BoolVarP(&_glueGetDataQualityRulesetEvaluationRun, "get-data-quality-ruleset-evaluation-run", "", false, "Get Data Quality Ruleset Evaluation Run")
	_glueCmd.Flags().BoolVarP(&_glueGetDatabase, "get-database", "", false, "Get Database")
	_glueCmd.Flags().BoolVarP(&_glueGetDatabases, "get-databases", "", false, "Get Databases")
	_glueCmd.Flags().BoolVarP(&_glueGetDataflowGraph, "get-dataflow-graph", "", false, "Get Dataflow Graph")
	_glueCmd.Flags().BoolVarP(&_glueGetDevEndpoint, "get-dev-endpoint", "", false, "Get Dev Endpoint")
	_glueCmd.Flags().BoolVarP(&_glueGetDevEndpoints, "get-dev-endpoints", "", false, "Get Dev Endpoints")
	_glueCmd.Flags().BoolVarP(&_glueGetEntityRecords, "get-entity-records", "", false, "Get Entity Records")
	_glueCmd.Flags().BoolVarP(&_glueGetGlueIdentityCenterConfiguration, "get-glue-identity-center-configuration", "", false, "Get Glue Identity Center Configuration")
	_glueCmd.Flags().BoolVarP(&_glueGetIntegrationResourceProperty, "get-integration-resource-property", "", false, "Get Integration Resource Property")
	_glueCmd.Flags().BoolVarP(&_glueGetIntegrationTableProperties, "get-integration-table-properties", "", false, "Get Integration Table Properties")
	_glueCmd.Flags().BoolVarP(&_glueGetJob, "get-job", "", false, "Get Job")
	_glueCmd.Flags().BoolVarP(&_glueGetJobBookmark, "get-job-bookmark", "", false, "Get Job Bookmark")
	_glueCmd.Flags().BoolVarP(&_glueGetJobRun, "get-job-run", "", false, "Get Job Run")
	_glueCmd.Flags().BoolVarP(&_glueGetJobRuns, "get-job-runs", "", false, "Get Job Runs")
	_glueCmd.Flags().BoolVarP(&_glueGetJobs, "get-jobs", "", false, "Get Jobs")
	_glueCmd.Flags().BoolVarP(&_glueGetMapping, "get-mapping", "", false, "Get Mapping")
	_glueCmd.Flags().BoolVarP(&_glueGetMaterializedViewRefreshTaskRun, "get-materialized-view-refresh-task-run", "", false, "Get Materialized View Refresh Task Run")
	_glueCmd.Flags().BoolVarP(&_glueGetMLTaskRun, "get-ml-task-run", "", false, "Get Ml Task Run")
	_glueCmd.Flags().BoolVarP(&_glueGetMLTaskRuns, "get-ml-task-runs", "", false, "Get Ml Task Runs")
	_glueCmd.Flags().BoolVarP(&_glueGetMLTransform, "get-ml-transform", "", false, "Get Ml Transform")
	_glueCmd.Flags().BoolVarP(&_glueGetMLTransforms, "get-ml-transforms", "", false, "Get Ml Transforms")
	_glueCmd.Flags().BoolVarP(&_glueGetPartition, "get-partition", "", false, "Get Partition")
	_glueCmd.Flags().BoolVarP(&_glueGetPartitionIndexes, "get-partition-indexes", "", false, "Get Partition Indexes")
	_glueCmd.Flags().BoolVarP(&_glueGetPartitions, "get-partitions", "", false, "Get Partitions")
	_glueCmd.Flags().BoolVarP(&_glueGetPlan, "get-plan", "", false, "Get Plan")
	_glueCmd.Flags().BoolVarP(&_glueGetRegistry, "get-registry", "", false, "Get Registry")
	_glueCmd.Flags().BoolVarP(&_glueGetResourcePolicies, "get-resource-policies", "", false, "Get Resource Policies")
	_glueCmd.Flags().BoolVarP(&_glueGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_glueCmd.Flags().BoolVarP(&_glueGetSchema, "get-schema", "", false, "Get Schema")
	_glueCmd.Flags().BoolVarP(&_glueGetSchemaByDefinition, "get-schema-by-definition", "", false, "Get Schema By Definition")
	_glueCmd.Flags().BoolVarP(&_glueGetSchemaVersion, "get-schema-version", "", false, "Get Schema Version")
	_glueCmd.Flags().BoolVarP(&_glueGetSchemaVersionsDiff, "get-schema-versions-diff", "", false, "Get Schema Versions Diff")
	_glueCmd.Flags().BoolVarP(&_glueGetSecurityConfiguration, "get-security-configuration", "", false, "Get Security Configuration")
	_glueCmd.Flags().BoolVarP(&_glueGetSecurityConfigurations, "get-security-configurations", "", false, "Get Security Configurations")
	_glueCmd.Flags().BoolVarP(&_glueGetSession, "get-session", "", false, "Get Session")
	_glueCmd.Flags().BoolVarP(&_glueGetStatement, "get-statement", "", false, "Get Statement")
	_glueCmd.Flags().BoolVarP(&_glueGetTable, "get-table", "", false, "Get Table")
	_glueCmd.Flags().BoolVarP(&_glueGetTableOptimizer, "get-table-optimizer", "", false, "Get Table Optimizer")
	_glueCmd.Flags().BoolVarP(&_glueGetTableVersion, "get-table-version", "", false, "Get Table Version")
	_glueCmd.Flags().BoolVarP(&_glueGetTableVersions, "get-table-versions", "", false, "Get Table Versions")
	_glueCmd.Flags().BoolVarP(&_glueGetTables, "get-tables", "", false, "Get Tables")
	_glueCmd.Flags().BoolVarP(&_glueGetTags, "get-tags", "", false, "Get Tags")
	_glueCmd.Flags().BoolVarP(&_glueGetTrigger, "get-trigger", "", false, "Get Trigger")
	_glueCmd.Flags().BoolVarP(&_glueGetTriggers, "get-triggers", "", false, "Get Triggers")
	_glueCmd.Flags().BoolVarP(&_glueGetUnfilteredPartitionMetadata, "get-unfiltered-partition-metadata", "", false, "Get Unfiltered Partition Metadata")
	_glueCmd.Flags().BoolVarP(&_glueGetUnfilteredPartitionsMetadata, "get-unfiltered-partitions-metadata", "", false, "Get Unfiltered Partitions Metadata")
	_glueCmd.Flags().BoolVarP(&_glueGetUnfilteredTableMetadata, "get-unfiltered-table-metadata", "", false, "Get Unfiltered Table Metadata")
	_glueCmd.Flags().BoolVarP(&_glueGetUsageProfile, "get-usage-profile", "", false, "Get Usage Profile")
	_glueCmd.Flags().BoolVarP(&_glueGetUserDefinedFunction, "get-user-defined-function", "", false, "Get User Defined Function")
	_glueCmd.Flags().BoolVarP(&_glueGetUserDefinedFunctions, "get-user-defined-functions", "", false, "Get User Defined Functions")
	_glueCmd.Flags().BoolVarP(&_glueGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_glueCmd.Flags().BoolVarP(&_glueGetWorkflowRun, "get-workflow-run", "", false, "Get Workflow Run")
	_glueCmd.Flags().BoolVarP(&_glueGetWorkflowRunProperties, "get-workflow-run-properties", "", false, "Get Workflow Run Properties")
	_glueCmd.Flags().BoolVarP(&_glueGetWorkflowRuns, "get-workflow-runs", "", false, "Get Workflow Runs")
	_glueCmd.Flags().BoolVarP(&_glueImportCatalogToGlue, "import-catalog-to-glue", "", false, "Import Catalog To Glue")
	_glueCmd.Flags().BoolVarP(&_glueListBlueprints, "list-blueprints", "", false, "List Blueprints")
	_glueCmd.Flags().BoolVarP(&_glueListColumnStatisticsTaskRuns, "list-column-statistics-task-runs", "", false, "List Column Statistics Task Runs")
	_glueCmd.Flags().BoolVarP(&_glueListConnectionTypes, "list-connection-types", "", false, "List Connection Types")
	_glueCmd.Flags().BoolVarP(&_glueListCrawlers, "list-crawlers", "", false, "List Crawlers")
	_glueCmd.Flags().BoolVarP(&_glueListCrawls, "list-crawls", "", false, "List Crawls")
	_glueCmd.Flags().BoolVarP(&_glueListCustomEntityTypes, "list-custom-entity-types", "", false, "List Custom Entity Types")
	_glueCmd.Flags().BoolVarP(&_glueListDataQualityResults, "list-data-quality-results", "", false, "List Data Quality Results")
	_glueCmd.Flags().BoolVarP(&_glueListDataQualityRuleRecommendationRuns, "list-data-quality-rule-recommendation-runs", "", false, "List Data Quality Rule Recommendation Runs")
	_glueCmd.Flags().BoolVarP(&_glueListDataQualityRulesetEvaluationRuns, "list-data-quality-ruleset-evaluation-runs", "", false, "List Data Quality Ruleset Evaluation Runs")
	_glueCmd.Flags().BoolVarP(&_glueListDataQualityRulesets, "list-data-quality-rulesets", "", false, "List Data Quality Rulesets")
	_glueCmd.Flags().BoolVarP(&_glueListDataQualityStatisticAnnotations, "list-data-quality-statistic-annotations", "", false, "List Data Quality Statistic Annotations")
	_glueCmd.Flags().BoolVarP(&_glueListDataQualityStatistics, "list-data-quality-statistics", "", false, "List Data Quality Statistics")
	_glueCmd.Flags().BoolVarP(&_glueListDevEndpoints, "list-dev-endpoints", "", false, "List Dev Endpoints")
	_glueCmd.Flags().BoolVarP(&_glueListEntities, "list-entities", "", false, "List Entities")
	_glueCmd.Flags().BoolVarP(&_glueListIntegrationResourceProperties, "list-integration-resource-properties", "", false, "List Integration Resource Properties")
	_glueCmd.Flags().BoolVarP(&_glueListJobs, "list-jobs", "", false, "List Jobs")
	_glueCmd.Flags().BoolVarP(&_glueListMaterializedViewRefreshTaskRuns, "list-materialized-view-refresh-task-runs", "", false, "List Materialized View Refresh Task Runs")
	_glueCmd.Flags().BoolVarP(&_glueListMLTransforms, "list-ml-transforms", "", false, "List Ml Transforms")
	_glueCmd.Flags().BoolVarP(&_glueListRegistries, "list-registries", "", false, "List Registries")
	_glueCmd.Flags().BoolVarP(&_glueListSchemaVersions, "list-schema-versions", "", false, "List Schema Versions")
	_glueCmd.Flags().BoolVarP(&_glueListSchemas, "list-schemas", "", false, "List Schemas")
	_glueCmd.Flags().BoolVarP(&_glueListSessions, "list-sessions", "", false, "List Sessions")
	_glueCmd.Flags().BoolVarP(&_glueListStatements, "list-statements", "", false, "List Statements")
	_glueCmd.Flags().BoolVarP(&_glueListTableOptimizerRuns, "list-table-optimizer-runs", "", false, "List Table Optimizer Runs")
	_glueCmd.Flags().BoolVarP(&_glueListTriggers, "list-triggers", "", false, "List Triggers")
	_glueCmd.Flags().BoolVarP(&_glueListUsageProfiles, "list-usage-profiles", "", false, "List Usage Profiles")
	_glueCmd.Flags().BoolVarP(&_glueListWorkflows, "list-workflows", "", false, "List Workflows")
	_glueCmd.Flags().BoolVarP(&_glueModifyIntegration, "modify-integration", "", false, "Modify Integration")
	_glueCmd.Flags().BoolVarP(&_gluePutDataCatalogEncryptionSettings, "put-data-catalog-encryption-settings", "", false, "Put Data Catalog Encryption Settings")
	_glueCmd.Flags().BoolVarP(&_gluePutDataQualityProfileAnnotation, "put-data-quality-profile-annotation", "", false, "Put Data Quality Profile Annotation")
	_glueCmd.Flags().BoolVarP(&_gluePutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_glueCmd.Flags().BoolVarP(&_gluePutSchemaVersionMetadata, "put-schema-version-metadata", "", false, "Put Schema Version Metadata")
	_glueCmd.Flags().BoolVarP(&_gluePutWorkflowRunProperties, "put-workflow-run-properties", "", false, "Put Workflow Run Properties")
	_glueCmd.Flags().BoolVarP(&_glueQuerySchemaVersionMetadata, "query-schema-version-metadata", "", false, "Query Schema Version Metadata")
	_glueCmd.Flags().BoolVarP(&_glueRegisterConnectionType, "register-connection-type", "", false, "Register Connection Type")
	_glueCmd.Flags().BoolVarP(&_glueRegisterSchemaVersion, "register-schema-version", "", false, "Register Schema Version")
	_glueCmd.Flags().BoolVarP(&_glueRemoveSchemaVersionMetadata, "remove-schema-version-metadata", "", false, "Remove Schema Version Metadata")
	_glueCmd.Flags().BoolVarP(&_glueResetJobBookmark, "reset-job-bookmark", "", false, "Reset Job Bookmark")
	_glueCmd.Flags().BoolVarP(&_glueResumeWorkflowRun, "resume-workflow-run", "", false, "Resume Workflow Run")
	_glueCmd.Flags().BoolVarP(&_glueRunStatement, "run-statement", "", false, "Run Statement")
	_glueCmd.Flags().BoolVarP(&_glueSearchTables, "search-tables", "", false, "Search Tables")
	_glueCmd.Flags().BoolVarP(&_glueStartBlueprintRun, "start-blueprint-run", "", false, "Start Blueprint Run")
	_glueCmd.Flags().BoolVarP(&_glueStartColumnStatisticsTaskRun, "start-column-statistics-task-run", "", false, "Start Column Statistics Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStartColumnStatisticsTaskRunSchedule, "start-column-statistics-task-run-schedule", "", false, "Start Column Statistics Task Run Schedule")
	_glueCmd.Flags().BoolVarP(&_glueStartCrawler, "start-crawler", "", false, "Start Crawler")
	_glueCmd.Flags().BoolVarP(&_glueStartCrawlerSchedule, "start-crawler-schedule", "", false, "Start Crawler Schedule")
	_glueCmd.Flags().BoolVarP(&_glueStartDataQualityRuleRecommendationRun, "start-data-quality-rule-recommendation-run", "", false, "Start Data Quality Rule Recommendation Run")
	_glueCmd.Flags().BoolVarP(&_glueStartDataQualityRulesetEvaluationRun, "start-data-quality-ruleset-evaluation-run", "", false, "Start Data Quality Ruleset Evaluation Run")
	_glueCmd.Flags().BoolVarP(&_glueStartExportLabelsTaskRun, "start-export-labels-task-run", "", false, "Start Export Labels Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStartImportLabelsTaskRun, "start-import-labels-task-run", "", false, "Start Import Labels Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStartJobRun, "start-job-run", "", false, "Start Job Run")
	_glueCmd.Flags().BoolVarP(&_glueStartMaterializedViewRefreshTaskRun, "start-materialized-view-refresh-task-run", "", false, "Start Materialized View Refresh Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStartMLEvaluationTaskRun, "start-ml-evaluation-task-run", "", false, "Start Ml Evaluation Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStartMLLabelingSetGenerationTaskRun, "start-ml-labeling-set-generation-task-run", "", false, "Start Ml Labeling Set Generation Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStartTrigger, "start-trigger", "", false, "Start Trigger")
	_glueCmd.Flags().BoolVarP(&_glueStartWorkflowRun, "start-workflow-run", "", false, "Start Workflow Run")
	_glueCmd.Flags().BoolVarP(&_glueStopColumnStatisticsTaskRun, "stop-column-statistics-task-run", "", false, "Stop Column Statistics Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStopColumnStatisticsTaskRunSchedule, "stop-column-statistics-task-run-schedule", "", false, "Stop Column Statistics Task Run Schedule")
	_glueCmd.Flags().BoolVarP(&_glueStopCrawler, "stop-crawler", "", false, "Stop Crawler")
	_glueCmd.Flags().BoolVarP(&_glueStopCrawlerSchedule, "stop-crawler-schedule", "", false, "Stop Crawler Schedule")
	_glueCmd.Flags().BoolVarP(&_glueStopMaterializedViewRefreshTaskRun, "stop-materialized-view-refresh-task-run", "", false, "Stop Materialized View Refresh Task Run")
	_glueCmd.Flags().BoolVarP(&_glueStopSession, "stop-session", "", false, "Stop Session")
	_glueCmd.Flags().BoolVarP(&_glueStopTrigger, "stop-trigger", "", false, "Stop Trigger")
	_glueCmd.Flags().BoolVarP(&_glueStopWorkflowRun, "stop-workflow-run", "", false, "Stop Workflow Run")
	_glueCmd.Flags().BoolVarP(&_glueTagResource, "tag-resource", "", false, "Tag Resource")
	_glueCmd.Flags().BoolVarP(&_glueTestConnection, "test-connection", "", false, "Test Connection")
	_glueCmd.Flags().BoolVarP(&_glueUntagResource, "untag-resource", "", false, "Untag Resource")
	_glueCmd.Flags().BoolVarP(&_glueUpdateBlueprint, "update-blueprint", "", false, "Update Blueprint")
	_glueCmd.Flags().BoolVarP(&_glueUpdateCatalog, "update-catalog", "", false, "Update Catalog")
	_glueCmd.Flags().BoolVarP(&_glueUpdateClassifier, "update-classifier", "", false, "Update Classifier")
	_glueCmd.Flags().BoolVarP(&_glueUpdateColumnStatisticsForPartition, "update-column-statistics-for-partition", "", false, "Update Column Statistics For Partition")
	_glueCmd.Flags().BoolVarP(&_glueUpdateColumnStatisticsForTable, "update-column-statistics-for-table", "", false, "Update Column Statistics For Table")
	_glueCmd.Flags().BoolVarP(&_glueUpdateColumnStatisticsTaskSettings, "update-column-statistics-task-settings", "", false, "Update Column Statistics Task Settings")
	_glueCmd.Flags().BoolVarP(&_glueUpdateConnection, "update-connection", "", false, "Update Connection")
	_glueCmd.Flags().BoolVarP(&_glueUpdateCrawler, "update-crawler", "", false, "Update Crawler")
	_glueCmd.Flags().BoolVarP(&_glueUpdateCrawlerSchedule, "update-crawler-schedule", "", false, "Update Crawler Schedule")
	_glueCmd.Flags().BoolVarP(&_glueUpdateDataQualityRuleset, "update-data-quality-ruleset", "", false, "Update Data Quality Ruleset")
	_glueCmd.Flags().BoolVarP(&_glueUpdateDatabase, "update-database", "", false, "Update Database")
	_glueCmd.Flags().BoolVarP(&_glueUpdateDevEndpoint, "update-dev-endpoint", "", false, "Update Dev Endpoint")
	_glueCmd.Flags().BoolVarP(&_glueUpdateGlueIdentityCenterConfiguration, "update-glue-identity-center-configuration", "", false, "Update Glue Identity Center Configuration")
	_glueCmd.Flags().BoolVarP(&_glueUpdateIntegrationResourceProperty, "update-integration-resource-property", "", false, "Update Integration Resource Property")
	_glueCmd.Flags().BoolVarP(&_glueUpdateIntegrationTableProperties, "update-integration-table-properties", "", false, "Update Integration Table Properties")
	_glueCmd.Flags().BoolVarP(&_glueUpdateJob, "update-job", "", false, "Update Job")
	_glueCmd.Flags().BoolVarP(&_glueUpdateJobFromSourceControl, "update-job-from-source-control", "", false, "Update Job From Source Control")
	_glueCmd.Flags().BoolVarP(&_glueUpdateMLTransform, "update-ml-transform", "", false, "Update Ml Transform")
	_glueCmd.Flags().BoolVarP(&_glueUpdatePartition, "update-partition", "", false, "Update Partition")
	_glueCmd.Flags().BoolVarP(&_glueUpdateRegistry, "update-registry", "", false, "Update Registry")
	_glueCmd.Flags().BoolVarP(&_glueUpdateSchema, "update-schema", "", false, "Update Schema")
	_glueCmd.Flags().BoolVarP(&_glueUpdateSourceControlFromJob, "update-source-control-from-job", "", false, "Update Source Control From Job")
	_glueCmd.Flags().BoolVarP(&_glueUpdateTable, "update-table", "", false, "Update Table")
	_glueCmd.Flags().BoolVarP(&_glueUpdateTableOptimizer, "update-table-optimizer", "", false, "Update Table Optimizer")
	_glueCmd.Flags().BoolVarP(&_glueUpdateTrigger, "update-trigger", "", false, "Update Trigger")
	_glueCmd.Flags().BoolVarP(&_glueUpdateUsageProfile, "update-usage-profile", "", false, "Update Usage Profile")
	_glueCmd.Flags().BoolVarP(&_glueUpdateUserDefinedFunction, "update-user-defined-function", "", false, "Update User Defined Function")
	_glueCmd.Flags().BoolVarP(&_glueUpdateWorkflow, "update-workflow", "", false, "Update Workflow")

}
