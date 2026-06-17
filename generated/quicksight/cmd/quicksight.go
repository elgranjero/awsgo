package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// quicksightCmd represents the quicksight command
var _quicksightCmd = &cobra.Command{
	Use:   "quicksight",
	Short: "AWS quicksight CLI",
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
		client := quicksight.NewFromConfig(cfg)
		if _quicksightBatchCreateTopicReviewedAnswer {
			quicksight_BatchCreateTopicReviewedAnswer(cfg, client)
			return
		}
		if _quicksightBatchDeleteTopicReviewedAnswer {
			quicksight_BatchDeleteTopicReviewedAnswer(cfg, client)
			return
		}
		if _quicksightCancelIngestion {
			quicksight_CancelIngestion(cfg, client)
			return
		}
		if _quicksightCreateAccountCustomization {
			quicksight_CreateAccountCustomization(cfg, client)
			return
		}
		if _quicksightCreateAccountSubscription {
			quicksight_CreateAccountSubscription(cfg, client)
			return
		}
		if _quicksightCreateActionConnector {
			quicksight_CreateActionConnector(cfg, client)
			return
		}
		if _quicksightCreateAnalysis {
			quicksight_CreateAnalysis(cfg, client)
			return
		}
		if _quicksightCreateBrand {
			quicksight_CreateBrand(cfg, client)
			return
		}
		if _quicksightCreateCustomPermissions {
			quicksight_CreateCustomPermissions(cfg, client)
			return
		}
		if _quicksightCreateDashboard {
			quicksight_CreateDashboard(cfg, client)
			return
		}
		if _quicksightCreateDataSet {
			quicksight_CreateDataSet(cfg, client)
			return
		}
		if _quicksightCreateDataSource {
			quicksight_CreateDataSource(cfg, client)
			return
		}
		if _quicksightCreateFolder {
			quicksight_CreateFolder(cfg, client)
			return
		}
		if _quicksightCreateFolderMembership {
			quicksight_CreateFolderMembership(cfg, client)
			return
		}
		if _quicksightCreateGroup {
			quicksight_CreateGroup(cfg, client)
			return
		}
		if _quicksightCreateGroupMembership {
			quicksight_CreateGroupMembership(cfg, client)
			return
		}
		if _quicksightCreateIAMPolicyAssignment {
			quicksight_CreateIAMPolicyAssignment(cfg, client)
			return
		}
		if _quicksightCreateIngestion {
			quicksight_CreateIngestion(cfg, client)
			return
		}
		if _quicksightCreateNamespace {
			quicksight_CreateNamespace(cfg, client)
			return
		}
		if _quicksightCreateRefreshSchedule {
			quicksight_CreateRefreshSchedule(cfg, client)
			return
		}
		if _quicksightCreateRoleMembership {
			quicksight_CreateRoleMembership(cfg, client)
			return
		}
		if _quicksightCreateTemplate {
			quicksight_CreateTemplate(cfg, client)
			return
		}
		if _quicksightCreateTemplateAlias {
			quicksight_CreateTemplateAlias(cfg, client)
			return
		}
		if _quicksightCreateTheme {
			quicksight_CreateTheme(cfg, client)
			return
		}
		if _quicksightCreateThemeAlias {
			quicksight_CreateThemeAlias(cfg, client)
			return
		}
		if _quicksightCreateTopic {
			quicksight_CreateTopic(cfg, client)
			return
		}
		if _quicksightCreateTopicRefreshSchedule {
			quicksight_CreateTopicRefreshSchedule(cfg, client)
			return
		}
		if _quicksightCreateVPCConnection {
			quicksight_CreateVPCConnection(cfg, client)
			return
		}
		if _quicksightDeleteAccountCustomPermission {
			quicksight_DeleteAccountCustomPermission(cfg, client)
			return
		}
		if _quicksightDeleteAccountCustomization {
			quicksight_DeleteAccountCustomization(cfg, client)
			return
		}
		if _quicksightDeleteAccountSubscription {
			quicksight_DeleteAccountSubscription(cfg, client)
			return
		}
		if _quicksightDeleteActionConnector {
			quicksight_DeleteActionConnector(cfg, client)
			return
		}
		if _quicksightDeleteAnalysis {
			quicksight_DeleteAnalysis(cfg, client)
			return
		}
		if _quicksightDeleteBrand {
			quicksight_DeleteBrand(cfg, client)
			return
		}
		if _quicksightDeleteBrandAssignment {
			quicksight_DeleteBrandAssignment(cfg, client)
			return
		}
		if _quicksightDeleteCustomPermissions {
			quicksight_DeleteCustomPermissions(cfg, client)
			return
		}
		if _quicksightDeleteDashboard {
			quicksight_DeleteDashboard(cfg, client)
			return
		}
		if _quicksightDeleteDataSet {
			quicksight_DeleteDataSet(cfg, client)
			return
		}
		if _quicksightDeleteDataSetRefreshProperties {
			quicksight_DeleteDataSetRefreshProperties(cfg, client)
			return
		}
		if _quicksightDeleteDataSource {
			quicksight_DeleteDataSource(cfg, client)
			return
		}
		if _quicksightDeleteDefaultQBusinessApplication {
			quicksight_DeleteDefaultQBusinessApplication(cfg, client)
			return
		}
		if _quicksightDeleteFolder {
			quicksight_DeleteFolder(cfg, client)
			return
		}
		if _quicksightDeleteFolderMembership {
			quicksight_DeleteFolderMembership(cfg, client)
			return
		}
		if _quicksightDeleteGroup {
			quicksight_DeleteGroup(cfg, client)
			return
		}
		if _quicksightDeleteGroupMembership {
			quicksight_DeleteGroupMembership(cfg, client)
			return
		}
		if _quicksightDeleteIAMPolicyAssignment {
			quicksight_DeleteIAMPolicyAssignment(cfg, client)
			return
		}
		if _quicksightDeleteIdentityPropagationConfig {
			quicksight_DeleteIdentityPropagationConfig(cfg, client)
			return
		}
		if _quicksightDeleteNamespace {
			quicksight_DeleteNamespace(cfg, client)
			return
		}
		if _quicksightDeleteRefreshSchedule {
			quicksight_DeleteRefreshSchedule(cfg, client)
			return
		}
		if _quicksightDeleteRoleCustomPermission {
			quicksight_DeleteRoleCustomPermission(cfg, client)
			return
		}
		if _quicksightDeleteRoleMembership {
			quicksight_DeleteRoleMembership(cfg, client)
			return
		}
		if _quicksightDeleteTemplate {
			quicksight_DeleteTemplate(cfg, client)
			return
		}
		if _quicksightDeleteTemplateAlias {
			quicksight_DeleteTemplateAlias(cfg, client)
			return
		}
		if _quicksightDeleteTheme {
			quicksight_DeleteTheme(cfg, client)
			return
		}
		if _quicksightDeleteThemeAlias {
			quicksight_DeleteThemeAlias(cfg, client)
			return
		}
		if _quicksightDeleteTopic {
			quicksight_DeleteTopic(cfg, client)
			return
		}
		if _quicksightDeleteTopicRefreshSchedule {
			quicksight_DeleteTopicRefreshSchedule(cfg, client)
			return
		}
		if _quicksightDeleteUser {
			quicksight_DeleteUser(cfg, client)
			return
		}
		if _quicksightDeleteUserByPrincipalId {
			quicksight_DeleteUserByPrincipalId(cfg, client)
			return
		}
		if _quicksightDeleteUserCustomPermission {
			quicksight_DeleteUserCustomPermission(cfg, client)
			return
		}
		if _quicksightDeleteVPCConnection {
			quicksight_DeleteVPCConnection(cfg, client)
			return
		}
		if _quicksightDescribeAccountCustomPermission {
			quicksight_DescribeAccountCustomPermission(cfg, client)
			return
		}
		if _quicksightDescribeAccountCustomization {
			quicksight_DescribeAccountCustomization(cfg, client)
			return
		}
		if _quicksightDescribeAccountSettings {
			quicksight_DescribeAccountSettings(cfg, client)
			return
		}
		if _quicksightDescribeAccountSubscription {
			quicksight_DescribeAccountSubscription(cfg, client)
			return
		}
		if _quicksightDescribeActionConnector {
			quicksight_DescribeActionConnector(cfg, client)
			return
		}
		if _quicksightDescribeActionConnectorPermissions {
			quicksight_DescribeActionConnectorPermissions(cfg, client)
			return
		}
		if _quicksightDescribeAnalysis {
			quicksight_DescribeAnalysis(cfg, client)
			return
		}
		if _quicksightDescribeAnalysisDefinition {
			quicksight_DescribeAnalysisDefinition(cfg, client)
			return
		}
		if _quicksightDescribeAnalysisPermissions {
			quicksight_DescribeAnalysisPermissions(cfg, client)
			return
		}
		if _quicksightDescribeAssetBundleExportJob {
			quicksight_DescribeAssetBundleExportJob(cfg, client)
			return
		}
		if _quicksightDescribeAssetBundleImportJob {
			quicksight_DescribeAssetBundleImportJob(cfg, client)
			return
		}
		if _quicksightDescribeBrand {
			quicksight_DescribeBrand(cfg, client)
			return
		}
		if _quicksightDescribeBrandAssignment {
			quicksight_DescribeBrandAssignment(cfg, client)
			return
		}
		if _quicksightDescribeBrandPublishedVersion {
			quicksight_DescribeBrandPublishedVersion(cfg, client)
			return
		}
		if _quicksightDescribeCustomPermissions {
			quicksight_DescribeCustomPermissions(cfg, client)
			return
		}
		if _quicksightDescribeDashboard {
			quicksight_DescribeDashboard(cfg, client)
			return
		}
		if _quicksightDescribeDashboardDefinition {
			quicksight_DescribeDashboardDefinition(cfg, client)
			return
		}
		if _quicksightDescribeDashboardPermissions {
			quicksight_DescribeDashboardPermissions(cfg, client)
			return
		}
		if _quicksightDescribeDashboardSnapshotJob {
			quicksight_DescribeDashboardSnapshotJob(cfg, client)
			return
		}
		if _quicksightDescribeDashboardSnapshotJobResult {
			quicksight_DescribeDashboardSnapshotJobResult(cfg, client)
			return
		}
		if _quicksightDescribeDashboardsQAConfiguration {
			quicksight_DescribeDashboardsQAConfiguration(cfg, client)
			return
		}
		if _quicksightDescribeDataSet {
			quicksight_DescribeDataSet(cfg, client)
			return
		}
		if _quicksightDescribeDataSetPermissions {
			quicksight_DescribeDataSetPermissions(cfg, client)
			return
		}
		if _quicksightDescribeDataSetRefreshProperties {
			quicksight_DescribeDataSetRefreshProperties(cfg, client)
			return
		}
		if _quicksightDescribeDataSource {
			quicksight_DescribeDataSource(cfg, client)
			return
		}
		if _quicksightDescribeDataSourcePermissions {
			quicksight_DescribeDataSourcePermissions(cfg, client)
			return
		}
		if _quicksightDescribeDefaultQBusinessApplication {
			quicksight_DescribeDefaultQBusinessApplication(cfg, client)
			return
		}
		if _quicksightDescribeFolder {
			quicksight_DescribeFolder(cfg, client)
			return
		}
		if _quicksightDescribeFolderPermissions {
			quicksight_DescribeFolderPermissions(cfg, client)
			return
		}
		if _quicksightDescribeFolderResolvedPermissions {
			quicksight_DescribeFolderResolvedPermissions(cfg, client)
			return
		}
		if _quicksightDescribeGroup {
			quicksight_DescribeGroup(cfg, client)
			return
		}
		if _quicksightDescribeGroupMembership {
			quicksight_DescribeGroupMembership(cfg, client)
			return
		}
		if _quicksightDescribeIAMPolicyAssignment {
			quicksight_DescribeIAMPolicyAssignment(cfg, client)
			return
		}
		if _quicksightDescribeIngestion {
			quicksight_DescribeIngestion(cfg, client)
			return
		}
		if _quicksightDescribeIpRestriction {
			quicksight_DescribeIpRestriction(cfg, client)
			return
		}
		if _quicksightDescribeKeyRegistration {
			quicksight_DescribeKeyRegistration(cfg, client)
			return
		}
		if _quicksightDescribeNamespace {
			quicksight_DescribeNamespace(cfg, client)
			return
		}
		if _quicksightDescribeQPersonalizationConfiguration {
			quicksight_DescribeQPersonalizationConfiguration(cfg, client)
			return
		}
		if _quicksightDescribeQuickSightQSearchConfiguration {
			quicksight_DescribeQuickSightQSearchConfiguration(cfg, client)
			return
		}
		if _quicksightDescribeRefreshSchedule {
			quicksight_DescribeRefreshSchedule(cfg, client)
			return
		}
		if _quicksightDescribeRoleCustomPermission {
			quicksight_DescribeRoleCustomPermission(cfg, client)
			return
		}
		if _quicksightDescribeSelfUpgradeConfiguration {
			quicksight_DescribeSelfUpgradeConfiguration(cfg, client)
			return
		}
		if _quicksightDescribeTemplate {
			quicksight_DescribeTemplate(cfg, client)
			return
		}
		if _quicksightDescribeTemplateAlias {
			quicksight_DescribeTemplateAlias(cfg, client)
			return
		}
		if _quicksightDescribeTemplateDefinition {
			quicksight_DescribeTemplateDefinition(cfg, client)
			return
		}
		if _quicksightDescribeTemplatePermissions {
			quicksight_DescribeTemplatePermissions(cfg, client)
			return
		}
		if _quicksightDescribeTheme {
			quicksight_DescribeTheme(cfg, client)
			return
		}
		if _quicksightDescribeThemeAlias {
			quicksight_DescribeThemeAlias(cfg, client)
			return
		}
		if _quicksightDescribeThemePermissions {
			quicksight_DescribeThemePermissions(cfg, client)
			return
		}
		if _quicksightDescribeTopic {
			quicksight_DescribeTopic(cfg, client)
			return
		}
		if _quicksightDescribeTopicPermissions {
			quicksight_DescribeTopicPermissions(cfg, client)
			return
		}
		if _quicksightDescribeTopicRefresh {
			quicksight_DescribeTopicRefresh(cfg, client)
			return
		}
		if _quicksightDescribeTopicRefreshSchedule {
			quicksight_DescribeTopicRefreshSchedule(cfg, client)
			return
		}
		if _quicksightDescribeUser {
			quicksight_DescribeUser(cfg, client)
			return
		}
		if _quicksightDescribeVPCConnection {
			quicksight_DescribeVPCConnection(cfg, client)
			return
		}
		if _quicksightGenerateEmbedUrlForAnonymousUser {
			quicksight_GenerateEmbedUrlForAnonymousUser(cfg, client)
			return
		}
		if _quicksightGenerateEmbedUrlForRegisteredUser {
			quicksight_GenerateEmbedUrlForRegisteredUser(cfg, client)
			return
		}
		if _quicksightGenerateEmbedUrlForRegisteredUserWithIdentity {
			quicksight_GenerateEmbedUrlForRegisteredUserWithIdentity(cfg, client)
			return
		}
		if _quicksightGetDashboardEmbedUrl {
			quicksight_GetDashboardEmbedUrl(cfg, client)
			return
		}
		if _quicksightGetFlowMetadata {
			quicksight_GetFlowMetadata(cfg, client)
			return
		}
		if _quicksightGetFlowPermissions {
			quicksight_GetFlowPermissions(cfg, client)
			return
		}
		if _quicksightGetIdentityContext {
			quicksight_GetIdentityContext(cfg, client)
			return
		}
		if _quicksightGetSessionEmbedUrl {
			quicksight_GetSessionEmbedUrl(cfg, client)
			return
		}
		if _quicksightListActionConnectors {
			quicksight_ListActionConnectors(cfg, client)
			return
		}
		if _quicksightListAnalyses {
			quicksight_ListAnalyses(cfg, client)
			return
		}
		if _quicksightListAssetBundleExportJobs {
			quicksight_ListAssetBundleExportJobs(cfg, client)
			return
		}
		if _quicksightListAssetBundleImportJobs {
			quicksight_ListAssetBundleImportJobs(cfg, client)
			return
		}
		if _quicksightListBrands {
			quicksight_ListBrands(cfg, client)
			return
		}
		if _quicksightListCustomPermissions {
			quicksight_ListCustomPermissions(cfg, client)
			return
		}
		if _quicksightListDashboardVersions {
			quicksight_ListDashboardVersions(cfg, client)
			return
		}
		if _quicksightListDashboards {
			quicksight_ListDashboards(cfg, client)
			return
		}
		if _quicksightListDataSets {
			quicksight_ListDataSets(cfg, client)
			return
		}
		if _quicksightListDataSources {
			quicksight_ListDataSources(cfg, client)
			return
		}
		if _quicksightListFlows {
			quicksight_ListFlows(cfg, client)
			return
		}
		if _quicksightListFolderMembers {
			quicksight_ListFolderMembers(cfg, client)
			return
		}
		if _quicksightListFolders {
			quicksight_ListFolders(cfg, client)
			return
		}
		if _quicksightListFoldersForResource {
			quicksight_ListFoldersForResource(cfg, client)
			return
		}
		if _quicksightListGroupMemberships {
			quicksight_ListGroupMemberships(cfg, client)
			return
		}
		if _quicksightListGroups {
			quicksight_ListGroups(cfg, client)
			return
		}
		if _quicksightListIAMPolicyAssignments {
			quicksight_ListIAMPolicyAssignments(cfg, client)
			return
		}
		if _quicksightListIAMPolicyAssignmentsForUser {
			quicksight_ListIAMPolicyAssignmentsForUser(cfg, client)
			return
		}
		if _quicksightListIdentityPropagationConfigs {
			quicksight_ListIdentityPropagationConfigs(cfg, client)
			return
		}
		if _quicksightListIngestions {
			quicksight_ListIngestions(cfg, client)
			return
		}
		if _quicksightListNamespaces {
			quicksight_ListNamespaces(cfg, client)
			return
		}
		if _quicksightListRefreshSchedules {
			quicksight_ListRefreshSchedules(cfg, client)
			return
		}
		if _quicksightListRoleMemberships {
			quicksight_ListRoleMemberships(cfg, client)
			return
		}
		if _quicksightListSelfUpgrades {
			quicksight_ListSelfUpgrades(cfg, client)
			return
		}
		if _quicksightListTagsForResource {
			quicksight_ListTagsForResource(cfg, client)
			return
		}
		if _quicksightListTemplateAliases {
			quicksight_ListTemplateAliases(cfg, client)
			return
		}
		if _quicksightListTemplateVersions {
			quicksight_ListTemplateVersions(cfg, client)
			return
		}
		if _quicksightListTemplates {
			quicksight_ListTemplates(cfg, client)
			return
		}
		if _quicksightListThemeAliases {
			quicksight_ListThemeAliases(cfg, client)
			return
		}
		if _quicksightListThemeVersions {
			quicksight_ListThemeVersions(cfg, client)
			return
		}
		if _quicksightListThemes {
			quicksight_ListThemes(cfg, client)
			return
		}
		if _quicksightListTopicRefreshSchedules {
			quicksight_ListTopicRefreshSchedules(cfg, client)
			return
		}
		if _quicksightListTopicReviewedAnswers {
			quicksight_ListTopicReviewedAnswers(cfg, client)
			return
		}
		if _quicksightListTopics {
			quicksight_ListTopics(cfg, client)
			return
		}
		if _quicksightListUserGroups {
			quicksight_ListUserGroups(cfg, client)
			return
		}
		if _quicksightListUsers {
			quicksight_ListUsers(cfg, client)
			return
		}
		if _quicksightListVPCConnections {
			quicksight_ListVPCConnections(cfg, client)
			return
		}
		if _quicksightPredictQAResults {
			quicksight_PredictQAResults(cfg, client)
			return
		}
		if _quicksightPutDataSetRefreshProperties {
			quicksight_PutDataSetRefreshProperties(cfg, client)
			return
		}
		if _quicksightRegisterUser {
			quicksight_RegisterUser(cfg, client)
			return
		}
		if _quicksightRestoreAnalysis {
			quicksight_RestoreAnalysis(cfg, client)
			return
		}
		if _quicksightSearchActionConnectors {
			quicksight_SearchActionConnectors(cfg, client)
			return
		}
		if _quicksightSearchAnalyses {
			quicksight_SearchAnalyses(cfg, client)
			return
		}
		if _quicksightSearchDashboards {
			quicksight_SearchDashboards(cfg, client)
			return
		}
		if _quicksightSearchDataSets {
			quicksight_SearchDataSets(cfg, client)
			return
		}
		if _quicksightSearchDataSources {
			quicksight_SearchDataSources(cfg, client)
			return
		}
		if _quicksightSearchFlows {
			quicksight_SearchFlows(cfg, client)
			return
		}
		if _quicksightSearchFolders {
			quicksight_SearchFolders(cfg, client)
			return
		}
		if _quicksightSearchGroups {
			quicksight_SearchGroups(cfg, client)
			return
		}
		if _quicksightSearchTopics {
			quicksight_SearchTopics(cfg, client)
			return
		}
		if _quicksightStartAssetBundleExportJob {
			quicksight_StartAssetBundleExportJob(cfg, client)
			return
		}
		if _quicksightStartAssetBundleImportJob {
			quicksight_StartAssetBundleImportJob(cfg, client)
			return
		}
		if _quicksightStartDashboardSnapshotJob {
			quicksight_StartDashboardSnapshotJob(cfg, client)
			return
		}
		if _quicksightStartDashboardSnapshotJobSchedule {
			quicksight_StartDashboardSnapshotJobSchedule(cfg, client)
			return
		}
		if _quicksightTagResource {
			quicksight_TagResource(cfg, client)
			return
		}
		if _quicksightUntagResource {
			quicksight_UntagResource(cfg, client)
			return
		}
		if _quicksightUpdateAccountCustomPermission {
			quicksight_UpdateAccountCustomPermission(cfg, client)
			return
		}
		if _quicksightUpdateAccountCustomization {
			quicksight_UpdateAccountCustomization(cfg, client)
			return
		}
		if _quicksightUpdateAccountSettings {
			quicksight_UpdateAccountSettings(cfg, client)
			return
		}
		if _quicksightUpdateActionConnector {
			quicksight_UpdateActionConnector(cfg, client)
			return
		}
		if _quicksightUpdateActionConnectorPermissions {
			quicksight_UpdateActionConnectorPermissions(cfg, client)
			return
		}
		if _quicksightUpdateAnalysis {
			quicksight_UpdateAnalysis(cfg, client)
			return
		}
		if _quicksightUpdateAnalysisPermissions {
			quicksight_UpdateAnalysisPermissions(cfg, client)
			return
		}
		if _quicksightUpdateApplicationWithTokenExchangeGrant {
			quicksight_UpdateApplicationWithTokenExchangeGrant(cfg, client)
			return
		}
		if _quicksightUpdateBrand {
			quicksight_UpdateBrand(cfg, client)
			return
		}
		if _quicksightUpdateBrandAssignment {
			quicksight_UpdateBrandAssignment(cfg, client)
			return
		}
		if _quicksightUpdateBrandPublishedVersion {
			quicksight_UpdateBrandPublishedVersion(cfg, client)
			return
		}
		if _quicksightUpdateCustomPermissions {
			quicksight_UpdateCustomPermissions(cfg, client)
			return
		}
		if _quicksightUpdateDashboard {
			quicksight_UpdateDashboard(cfg, client)
			return
		}
		if _quicksightUpdateDashboardLinks {
			quicksight_UpdateDashboardLinks(cfg, client)
			return
		}
		if _quicksightUpdateDashboardPermissions {
			quicksight_UpdateDashboardPermissions(cfg, client)
			return
		}
		if _quicksightUpdateDashboardPublishedVersion {
			quicksight_UpdateDashboardPublishedVersion(cfg, client)
			return
		}
		if _quicksightUpdateDashboardsQAConfiguration {
			quicksight_UpdateDashboardsQAConfiguration(cfg, client)
			return
		}
		if _quicksightUpdateDataSet {
			quicksight_UpdateDataSet(cfg, client)
			return
		}
		if _quicksightUpdateDataSetPermissions {
			quicksight_UpdateDataSetPermissions(cfg, client)
			return
		}
		if _quicksightUpdateDataSource {
			quicksight_UpdateDataSource(cfg, client)
			return
		}
		if _quicksightUpdateDataSourcePermissions {
			quicksight_UpdateDataSourcePermissions(cfg, client)
			return
		}
		if _quicksightUpdateDefaultQBusinessApplication {
			quicksight_UpdateDefaultQBusinessApplication(cfg, client)
			return
		}
		if _quicksightUpdateFlowPermissions {
			quicksight_UpdateFlowPermissions(cfg, client)
			return
		}
		if _quicksightUpdateFolder {
			quicksight_UpdateFolder(cfg, client)
			return
		}
		if _quicksightUpdateFolderPermissions {
			quicksight_UpdateFolderPermissions(cfg, client)
			return
		}
		if _quicksightUpdateGroup {
			quicksight_UpdateGroup(cfg, client)
			return
		}
		if _quicksightUpdateIAMPolicyAssignment {
			quicksight_UpdateIAMPolicyAssignment(cfg, client)
			return
		}
		if _quicksightUpdateIdentityPropagationConfig {
			quicksight_UpdateIdentityPropagationConfig(cfg, client)
			return
		}
		if _quicksightUpdateIpRestriction {
			quicksight_UpdateIpRestriction(cfg, client)
			return
		}
		if _quicksightUpdateKeyRegistration {
			quicksight_UpdateKeyRegistration(cfg, client)
			return
		}
		if _quicksightUpdatePublicSharingSettings {
			quicksight_UpdatePublicSharingSettings(cfg, client)
			return
		}
		if _quicksightUpdateQPersonalizationConfiguration {
			quicksight_UpdateQPersonalizationConfiguration(cfg, client)
			return
		}
		if _quicksightUpdateQuickSightQSearchConfiguration {
			quicksight_UpdateQuickSightQSearchConfiguration(cfg, client)
			return
		}
		if _quicksightUpdateRefreshSchedule {
			quicksight_UpdateRefreshSchedule(cfg, client)
			return
		}
		if _quicksightUpdateRoleCustomPermission {
			quicksight_UpdateRoleCustomPermission(cfg, client)
			return
		}
		if _quicksightUpdateSelfUpgrade {
			quicksight_UpdateSelfUpgrade(cfg, client)
			return
		}
		if _quicksightUpdateSelfUpgradeConfiguration {
			quicksight_UpdateSelfUpgradeConfiguration(cfg, client)
			return
		}
		if _quicksightUpdateSPICECapacityConfiguration {
			quicksight_UpdateSPICECapacityConfiguration(cfg, client)
			return
		}
		if _quicksightUpdateTemplate {
			quicksight_UpdateTemplate(cfg, client)
			return
		}
		if _quicksightUpdateTemplateAlias {
			quicksight_UpdateTemplateAlias(cfg, client)
			return
		}
		if _quicksightUpdateTemplatePermissions {
			quicksight_UpdateTemplatePermissions(cfg, client)
			return
		}
		if _quicksightUpdateTheme {
			quicksight_UpdateTheme(cfg, client)
			return
		}
		if _quicksightUpdateThemeAlias {
			quicksight_UpdateThemeAlias(cfg, client)
			return
		}
		if _quicksightUpdateThemePermissions {
			quicksight_UpdateThemePermissions(cfg, client)
			return
		}
		if _quicksightUpdateTopic {
			quicksight_UpdateTopic(cfg, client)
			return
		}
		if _quicksightUpdateTopicPermissions {
			quicksight_UpdateTopicPermissions(cfg, client)
			return
		}
		if _quicksightUpdateTopicRefreshSchedule {
			quicksight_UpdateTopicRefreshSchedule(cfg, client)
			return
		}
		if _quicksightUpdateUser {
			quicksight_UpdateUser(cfg, client)
			return
		}
		if _quicksightUpdateUserCustomPermission {
			quicksight_UpdateUserCustomPermission(cfg, client)
			return
		}
		if _quicksightUpdateVPCConnection {
			quicksight_UpdateVPCConnection(cfg, client)
			return
		}

	},
}

var (
	_quicksightBatchCreateTopicReviewedAnswer                bool
	_quicksightBatchDeleteTopicReviewedAnswer                bool
	_quicksightCancelIngestion                               bool
	_quicksightCreateAccountCustomization                    bool
	_quicksightCreateAccountSubscription                     bool
	_quicksightCreateActionConnector                         bool
	_quicksightCreateAnalysis                                bool
	_quicksightCreateBrand                                   bool
	_quicksightCreateCustomPermissions                       bool
	_quicksightCreateDashboard                               bool
	_quicksightCreateDataSet                                 bool
	_quicksightCreateDataSource                              bool
	_quicksightCreateFolder                                  bool
	_quicksightCreateFolderMembership                        bool
	_quicksightCreateGroup                                   bool
	_quicksightCreateGroupMembership                         bool
	_quicksightCreateIAMPolicyAssignment                     bool
	_quicksightCreateIngestion                               bool
	_quicksightCreateNamespace                               bool
	_quicksightCreateRefreshSchedule                         bool
	_quicksightCreateRoleMembership                          bool
	_quicksightCreateTemplate                                bool
	_quicksightCreateTemplateAlias                           bool
	_quicksightCreateTheme                                   bool
	_quicksightCreateThemeAlias                              bool
	_quicksightCreateTopic                                   bool
	_quicksightCreateTopicRefreshSchedule                    bool
	_quicksightCreateVPCConnection                           bool
	_quicksightDeleteAccountCustomPermission                 bool
	_quicksightDeleteAccountCustomization                    bool
	_quicksightDeleteAccountSubscription                     bool
	_quicksightDeleteActionConnector                         bool
	_quicksightDeleteAnalysis                                bool
	_quicksightDeleteBrand                                   bool
	_quicksightDeleteBrandAssignment                         bool
	_quicksightDeleteCustomPermissions                       bool
	_quicksightDeleteDashboard                               bool
	_quicksightDeleteDataSet                                 bool
	_quicksightDeleteDataSetRefreshProperties                bool
	_quicksightDeleteDataSource                              bool
	_quicksightDeleteDefaultQBusinessApplication             bool
	_quicksightDeleteFolder                                  bool
	_quicksightDeleteFolderMembership                        bool
	_quicksightDeleteGroup                                   bool
	_quicksightDeleteGroupMembership                         bool
	_quicksightDeleteIAMPolicyAssignment                     bool
	_quicksightDeleteIdentityPropagationConfig               bool
	_quicksightDeleteNamespace                               bool
	_quicksightDeleteRefreshSchedule                         bool
	_quicksightDeleteRoleCustomPermission                    bool
	_quicksightDeleteRoleMembership                          bool
	_quicksightDeleteTemplate                                bool
	_quicksightDeleteTemplateAlias                           bool
	_quicksightDeleteTheme                                   bool
	_quicksightDeleteThemeAlias                              bool
	_quicksightDeleteTopic                                   bool
	_quicksightDeleteTopicRefreshSchedule                    bool
	_quicksightDeleteUser                                    bool
	_quicksightDeleteUserByPrincipalId                       bool
	_quicksightDeleteUserCustomPermission                    bool
	_quicksightDeleteVPCConnection                           bool
	_quicksightDescribeAccountCustomPermission               bool
	_quicksightDescribeAccountCustomization                  bool
	_quicksightDescribeAccountSettings                       bool
	_quicksightDescribeAccountSubscription                   bool
	_quicksightDescribeActionConnector                       bool
	_quicksightDescribeActionConnectorPermissions            bool
	_quicksightDescribeAnalysis                              bool
	_quicksightDescribeAnalysisDefinition                    bool
	_quicksightDescribeAnalysisPermissions                   bool
	_quicksightDescribeAssetBundleExportJob                  bool
	_quicksightDescribeAssetBundleImportJob                  bool
	_quicksightDescribeBrand                                 bool
	_quicksightDescribeBrandAssignment                       bool
	_quicksightDescribeBrandPublishedVersion                 bool
	_quicksightDescribeCustomPermissions                     bool
	_quicksightDescribeDashboard                             bool
	_quicksightDescribeDashboardDefinition                   bool
	_quicksightDescribeDashboardPermissions                  bool
	_quicksightDescribeDashboardSnapshotJob                  bool
	_quicksightDescribeDashboardSnapshotJobResult            bool
	_quicksightDescribeDashboardsQAConfiguration             bool
	_quicksightDescribeDataSet                               bool
	_quicksightDescribeDataSetPermissions                    bool
	_quicksightDescribeDataSetRefreshProperties              bool
	_quicksightDescribeDataSource                            bool
	_quicksightDescribeDataSourcePermissions                 bool
	_quicksightDescribeDefaultQBusinessApplication           bool
	_quicksightDescribeFolder                                bool
	_quicksightDescribeFolderPermissions                     bool
	_quicksightDescribeFolderResolvedPermissions             bool
	_quicksightDescribeGroup                                 bool
	_quicksightDescribeGroupMembership                       bool
	_quicksightDescribeIAMPolicyAssignment                   bool
	_quicksightDescribeIngestion                             bool
	_quicksightDescribeIpRestriction                         bool
	_quicksightDescribeKeyRegistration                       bool
	_quicksightDescribeNamespace                             bool
	_quicksightDescribeQPersonalizationConfiguration         bool
	_quicksightDescribeQuickSightQSearchConfiguration        bool
	_quicksightDescribeRefreshSchedule                       bool
	_quicksightDescribeRoleCustomPermission                  bool
	_quicksightDescribeSelfUpgradeConfiguration              bool
	_quicksightDescribeTemplate                              bool
	_quicksightDescribeTemplateAlias                         bool
	_quicksightDescribeTemplateDefinition                    bool
	_quicksightDescribeTemplatePermissions                   bool
	_quicksightDescribeTheme                                 bool
	_quicksightDescribeThemeAlias                            bool
	_quicksightDescribeThemePermissions                      bool
	_quicksightDescribeTopic                                 bool
	_quicksightDescribeTopicPermissions                      bool
	_quicksightDescribeTopicRefresh                          bool
	_quicksightDescribeTopicRefreshSchedule                  bool
	_quicksightDescribeUser                                  bool
	_quicksightDescribeVPCConnection                         bool
	_quicksightGenerateEmbedUrlForAnonymousUser              bool
	_quicksightGenerateEmbedUrlForRegisteredUser             bool
	_quicksightGenerateEmbedUrlForRegisteredUserWithIdentity bool
	_quicksightGetDashboardEmbedUrl                          bool
	_quicksightGetFlowMetadata                               bool
	_quicksightGetFlowPermissions                            bool
	_quicksightGetIdentityContext                            bool
	_quicksightGetSessionEmbedUrl                            bool
	_quicksightListActionConnectors                          bool
	_quicksightListAnalyses                                  bool
	_quicksightListAssetBundleExportJobs                     bool
	_quicksightListAssetBundleImportJobs                     bool
	_quicksightListBrands                                    bool
	_quicksightListCustomPermissions                         bool
	_quicksightListDashboardVersions                         bool
	_quicksightListDashboards                                bool
	_quicksightListDataSets                                  bool
	_quicksightListDataSources                               bool
	_quicksightListFlows                                     bool
	_quicksightListFolderMembers                             bool
	_quicksightListFolders                                   bool
	_quicksightListFoldersForResource                        bool
	_quicksightListGroupMemberships                          bool
	_quicksightListGroups                                    bool
	_quicksightListIAMPolicyAssignments                      bool
	_quicksightListIAMPolicyAssignmentsForUser               bool
	_quicksightListIdentityPropagationConfigs                bool
	_quicksightListIngestions                                bool
	_quicksightListNamespaces                                bool
	_quicksightListRefreshSchedules                          bool
	_quicksightListRoleMemberships                           bool
	_quicksightListSelfUpgrades                              bool
	_quicksightListTagsForResource                           bool
	_quicksightListTemplateAliases                           bool
	_quicksightListTemplateVersions                          bool
	_quicksightListTemplates                                 bool
	_quicksightListThemeAliases                              bool
	_quicksightListThemeVersions                             bool
	_quicksightListThemes                                    bool
	_quicksightListTopicRefreshSchedules                     bool
	_quicksightListTopicReviewedAnswers                      bool
	_quicksightListTopics                                    bool
	_quicksightListUserGroups                                bool
	_quicksightListUsers                                     bool
	_quicksightListVPCConnections                            bool
	_quicksightPredictQAResults                              bool
	_quicksightPutDataSetRefreshProperties                   bool
	_quicksightRegisterUser                                  bool
	_quicksightRestoreAnalysis                               bool
	_quicksightSearchActionConnectors                        bool
	_quicksightSearchAnalyses                                bool
	_quicksightSearchDashboards                              bool
	_quicksightSearchDataSets                                bool
	_quicksightSearchDataSources                             bool
	_quicksightSearchFlows                                   bool
	_quicksightSearchFolders                                 bool
	_quicksightSearchGroups                                  bool
	_quicksightSearchTopics                                  bool
	_quicksightStartAssetBundleExportJob                     bool
	_quicksightStartAssetBundleImportJob                     bool
	_quicksightStartDashboardSnapshotJob                     bool
	_quicksightStartDashboardSnapshotJobSchedule             bool
	_quicksightTagResource                                   bool
	_quicksightUntagResource                                 bool
	_quicksightUpdateAccountCustomPermission                 bool
	_quicksightUpdateAccountCustomization                    bool
	_quicksightUpdateAccountSettings                         bool
	_quicksightUpdateActionConnector                         bool
	_quicksightUpdateActionConnectorPermissions              bool
	_quicksightUpdateAnalysis                                bool
	_quicksightUpdateAnalysisPermissions                     bool
	_quicksightUpdateApplicationWithTokenExchangeGrant       bool
	_quicksightUpdateBrand                                   bool
	_quicksightUpdateBrandAssignment                         bool
	_quicksightUpdateBrandPublishedVersion                   bool
	_quicksightUpdateCustomPermissions                       bool
	_quicksightUpdateDashboard                               bool
	_quicksightUpdateDashboardLinks                          bool
	_quicksightUpdateDashboardPermissions                    bool
	_quicksightUpdateDashboardPublishedVersion               bool
	_quicksightUpdateDashboardsQAConfiguration               bool
	_quicksightUpdateDataSet                                 bool
	_quicksightUpdateDataSetPermissions                      bool
	_quicksightUpdateDataSource                              bool
	_quicksightUpdateDataSourcePermissions                   bool
	_quicksightUpdateDefaultQBusinessApplication             bool
	_quicksightUpdateFlowPermissions                         bool
	_quicksightUpdateFolder                                  bool
	_quicksightUpdateFolderPermissions                       bool
	_quicksightUpdateGroup                                   bool
	_quicksightUpdateIAMPolicyAssignment                     bool
	_quicksightUpdateIdentityPropagationConfig               bool
	_quicksightUpdateIpRestriction                           bool
	_quicksightUpdateKeyRegistration                         bool
	_quicksightUpdatePublicSharingSettings                   bool
	_quicksightUpdateQPersonalizationConfiguration           bool
	_quicksightUpdateQuickSightQSearchConfiguration          bool
	_quicksightUpdateRefreshSchedule                         bool
	_quicksightUpdateRoleCustomPermission                    bool
	_quicksightUpdateSelfUpgrade                             bool
	_quicksightUpdateSelfUpgradeConfiguration                bool
	_quicksightUpdateSPICECapacityConfiguration              bool
	_quicksightUpdateTemplate                                bool
	_quicksightUpdateTemplateAlias                           bool
	_quicksightUpdateTemplatePermissions                     bool
	_quicksightUpdateTheme                                   bool
	_quicksightUpdateThemeAlias                              bool
	_quicksightUpdateThemePermissions                        bool
	_quicksightUpdateTopic                                   bool
	_quicksightUpdateTopicPermissions                        bool
	_quicksightUpdateTopicRefreshSchedule                    bool
	_quicksightUpdateUser                                    bool
	_quicksightUpdateUserCustomPermission                    bool
	_quicksightUpdateVPCConnection                           bool

	_quicksightAccountCustomization                        string
	_quicksightAccountName                                 string
	_quicksightAction                                      string
	_quicksightActionConnectorId                           string
	_quicksightActiveDirectoryName                         string
	_quicksightAdditionalDashboardIds                      []string
	_quicksightAdminGroup                                  []string
	_quicksightAdminProGroup                               []string
	_quicksightAliasName                                   string
	_quicksightAllowedDomains                              []string
	_quicksightAnalysisId                                  string
	_quicksightAnswerIds                                   []string
	_quicksightAnswers                                     string
	_quicksightApplicationId                               string
	_quicksightAssetBundleExportJobId                      string
	_quicksightAssetBundleImportJobId                      string
	_quicksightAssetBundleImportSource                     string
	_quicksightAssignmentName                              string
	_quicksightAssignmentStatus                            string
	_quicksightAuthenticationConfig                        string
	_quicksightAuthenticationMethod                        string
	_quicksightAuthorGroup                                 []string
	_quicksightAuthorProGroup                              []string
	_quicksightAuthorizedResourceArns                      []string
	_quicksightAuthorizedTargets                           []string
	_quicksightAwsAccountId                                string
	_quicksightBaseThemeId                                 string
	_quicksightBrandArn                                    string
	_quicksightBrandDefinition                             string
	_quicksightBrandId                                     string
	_quicksightCapabilities                                string
	_quicksightCloudFormationOverridePropertyConfiguration string
	_quicksightColumnGroups                                string
	_quicksightColumnLevelPermissionRules                  string
	_quicksightConfiguration                               string
	_quicksightContactNumber                               string
	_quicksightCredentials                                 string
	_quicksightCustomFederationProviderUrl                 string
	_quicksightCustomInstructions                          string
	_quicksightCustomPermissionsName                       string
	_quicksightDashboardId                                 string
	_quicksightDashboardPublishOptions                     string
	_quicksightDashboardsQAStatus                          string
	_quicksightDataPrepConfiguration                       string
	_quicksightDataSetId                                   string
	_quicksightDataSetRefreshProperties                    string
	_quicksightDataSetUsageConfiguration                   string
	_quicksightDataSourceId                                string
	_quicksightDataSourceParameters                        string
	_quicksightDatasetArn                                  string
	_quicksightDatasetId                                   string
	_quicksightDatasetName                                 string
	_quicksightDatasetParameters                           string
	_quicksightDefaultKeyOnly                              string
	_quicksightDefaultNamespace                            string
	_quicksightDefinition                                  string
	_quicksightDescription                                 string
	_quicksightDirectoryId                                 string
	_quicksightDnsResolvers                                []string
	_quicksightEdition                                     string
	_quicksightEmail                                       string
	_quicksightEmailAddress                                string
	_quicksightEnabled                                     string
	_quicksightEntryPoint                                  string
	_quicksightExperienceConfiguration                     string
	_quicksightExportFormat                                string
	_quicksightExternalLoginFederationProviderType         string
	_quicksightExternalLoginId                             string
	_quicksightFailureAction                               string
	_quicksightFieldFolders                                string
	_quicksightFilters                                     string
	_quicksightFirstName                                   string
	_quicksightFlowId                                      string
	_quicksightFolderArns                                  []string
	_quicksightFolderId                                    string
	_quicksightFolderType                                  string
	_quicksightForceDeleteWithoutRecovery                  string
	_quicksightGrantLinkPermissions                        string
	_quicksightGrantPermissions                            string
	_quicksightGroupName                                   string
	_quicksightIamArn                                      string
	_quicksightIAMIdentityCenterInstanceArn                string
	_quicksightIdentities                                  string
	_quicksightIdentityStore                               string
	_quicksightIdentityType                                string
	_quicksightImportMode                                  string
	_quicksightIncludeAllDependencies                      string
	_quicksightIncludeFolderMembers                        string
	_quicksightIncludeFolderMemberships                    string
	_quicksightIncludeGeneratedAnswer                      string
	_quicksightIncludePermissions                          string
	_quicksightIncludeQuickSightQIndex                     string
	_quicksightIncludeTags                                 string
	_quicksightIngestionId                                 string
	_quicksightIngestionType                               string
	_quicksightIpRestrictionRuleMap                        string
	_quicksightKeyRegistration                             string
	_quicksightLastName                                    string
	_quicksightLinkEntities                                []string
	_quicksightLinkSharingConfiguration                    string
	_quicksightLogicalTableMap                             string
	_quicksightMaxResults                                  string
	_quicksightMaxTopicsToConsider                         string
	_quicksightMemberId                                    string
	_quicksightMemberName                                  string
	_quicksightMemberType                                  string
	_quicksightName                                        string
	_quicksightNamespace                                   string
	_quicksightNextToken                                   string
	_quicksightNotificationEmail                           string
	_quicksightOverrideParameters                          string
	_quicksightOverridePermissions                         string
	_quicksightOverrideTags                                string
	_quicksightOverrideValidationStrategy                  string
	_quicksightParameters                                  string
	_quicksightParentFolderArn                             string
	_quicksightPerformanceConfiguration                    string
	_quicksightPermissions                                 string
	_quicksightPersonalizationMode                         string
	_quicksightPhysicalTableMap                            string
	_quicksightPolicyArn                                   string
	_quicksightPrincipalId                                 string
	_quicksightPublicSharingEnabled                        string
	_quicksightPurchaseMode                                string
	_quicksightQSearchStatus                               string
	_quicksightQueryText                                   string
	_quicksightReaderGroup                                 []string
	_quicksightReaderProGroup                              []string
	_quicksightRealm                                       string
	_quicksightRecoveryWindowInDays                        string
	_quicksightRefreshId                                   string
	_quicksightRefreshSchedule                             string
	_quicksightResetDisabled                               string
	_quicksightResolved                                    string
	_quicksightResourceArn                                 string
	_quicksightResourceArns                                []string
	_quicksightRestoreToFolders                            string
	_quicksightRevokeLinkPermissions                       string
	_quicksightRevokePermissions                           string
	_quicksightRole                                        string
	_quicksightRoleArn                                     string
	_quicksightRowLevelPermissionDataSet                   string
	_quicksightRowLevelPermissionTagConfiguration          string
	_quicksightSchedule                                    string
	_quicksightScheduleId                                  string
	_quicksightSecurityGroupIds                            []string
	_quicksightSelfUpgradeStatus                           string
	_quicksightSemanticModelConfiguration                  string
	_quicksightService                                     string
	_quicksightSessionExpiresAt                            string
	_quicksightSessionLifetimeInMinutes                    string
	_quicksightSessionName                                 string
	_quicksightSessionTags                                 string
	_quicksightSharingModel                                string
	_quicksightSnapshotConfiguration                       string
	_quicksightSnapshotJobId                               string
	_quicksightSourceEntity                                string
	_quicksightSslProperties                               string
	_quicksightStatePersistenceEnabled                     string
	_quicksightSubnetIds                                   []string
	_quicksightTagKeys                                     []string
	_quicksightTags                                        string
	_quicksightTemplateId                                  string
	_quicksightTemplateVersionNumber                       string
	_quicksightTerminationProtectionEnabled                string
	_quicksightThemeArn                                    string
	_quicksightThemeId                                     string
	_quicksightThemeVersionNumber                          string
	_quicksightTopic                                       string
	_quicksightTopicId                                     string
	_quicksightType                                        string
	_quicksightUnapplyCustomPermissions                    string
	_quicksightUndoRedoDisabled                            string
	_quicksightUpgradeRequestId                            string
	_quicksightUseAs                                       string
	_quicksightUserArn                                     string
	_quicksightUserConfiguration                           string
	_quicksightUserIdentifier                              string
	_quicksightUserName                                    string
	_quicksightUserRole                                    string
	_quicksightValidationStrategy                          string
	_quicksightVersionDescription                          string
	_quicksightVersionId                                   string
	_quicksightVersionNumber                               string
	_quicksightVpcConnectionArn                            string
	_quicksightVPCConnectionId                             string
	_quicksightVpcConnectionProperties                     string
	_quicksightVpcEndpointIdRestrictionRuleMap             string
	_quicksightVpcIdRestrictionRuleMap                     string
)

// Creates new reviewed answers for a Q Topic.
func quicksight_BatchCreateTopicReviewedAnswer(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.BatchCreateTopicReviewedAnswerInput{
		// Answers: []types.CreateTopicReviewedAnswer, // Required
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAnswers) > 0 {
		if err := assignInputField(input, "Answers", _quicksightAnswers); err != nil {
			log.Errorf("invalid --answers: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.BatchCreateTopicReviewedAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes reviewed answers for Q Topic.
func quicksight_BatchDeleteTopicReviewedAnswer(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.BatchDeleteTopicReviewedAnswerInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}
	if len(_quicksightAnswerIds) > 0 {
		input.AnswerIds = append([]string(nil), _quicksightAnswerIds...)
	}

	if resp, err := client.BatchDeleteTopicReviewedAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an ongoing ingestion of data into SPICE.
func quicksight_CancelIngestion(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CancelIngestionInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// IngestionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightIngestionId) > 0 {
		input.IngestionId = aws.String(_quicksightIngestionId)
	}

	if resp, err := client.CancelIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates Amazon Quick Sight customizations. Currently, you can add a custom
// default theme by using the CreateAccountCustomization or
// UpdateAccountCustomization API operation. To further customize Amazon Quick
// Sight by removing Amazon Quick Sight sample assets and videos for all new users,
// see [Customizing Quick Sight]in the Amazon Quick Sight User Guide.
//
// You can create customizations for your Amazon Web Services account or, if you
// specify a namespace, for a Quick Sight namespace instead. Customizations that
// apply to a namespace always override customizations that apply to an Amazon Web
// Services account. To find out which customizations apply, use the
// DescribeAccountCustomization API operation.
//
// Before you use the CreateAccountCustomization API operation to add a theme as
// the namespace default, make sure that you first share the theme with the
// namespace. If you don't share it with the namespace, the theme isn't visible to
// your users even if you make it the default theme. To check if the theme is
// shared, view the current permissions by using the [DescribeThemePermissions]API operation. To share the
// theme, grant permissions by using the [UpdateThemePermissions]API operation.
//
// [UpdateThemePermissions]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateThemePermissions.html
// [DescribeThemePermissions]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeThemePermissions.html
// [Customizing Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/customizing-quicksight.html
func quicksight_CreateAccountCustomization(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateAccountCustomizationInput{
		// AccountCustomization: *types.AccountCustomization, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAccountCustomization) > 0 {
		if err := assignInputField(input, "AccountCustomization", _quicksightAccountCustomization); err != nil {
			log.Errorf("invalid --account-customization: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccountCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Quick Sight account, or subscribes to Amazon Quick Sight Q.
// The Amazon Web Services Region for the account is derived from what is
// configured in the CLI or SDK.
//
// Before you use this operation, make sure that you can connect to an existing
// Amazon Web Services account. If you don't have an Amazon Web Services account,
// see [Sign up for Amazon Web Services]in the Amazon Quick Sight User Guide. The person who signs up for Amazon
// Quick Sight needs to have the correct Identity and Access Management (IAM)
// permissions. For more information, see [IAM Policy Examples for Amazon Quick Sight]in the Amazon Quick Sight User Guide.
//
// If your IAM policy includes both the Subscribe and CreateAccountSubscription
// actions, make sure that both actions are set to Allow . If either action is set
// to Deny , the Deny action prevails and your API call fails.
//
// You can't pass an existing IAM role to access other Amazon Web Services
// services using this API operation. To pass your existing IAM role to Amazon
// Quick Sight, see [Passing IAM roles to Amazon Quick Sight]in the Amazon Quick Sight User Guide.
//
// You can't set default resource access on the new account from the Amazon Quick
// Sight API. Instead, add default resource access from the Amazon Quick Sight
// console. For more information about setting default resource access to Amazon
// Web Services services, see [Setting default resource access to Amazon Web Services services]in the Amazon Quick Sight User Guide.
//
// [Setting default resource access to Amazon Web Services services]: https://docs.aws.amazon.com/quicksight/latest/user/scoping-policies-defaults.html
// [IAM Policy Examples for Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/iam-policy-examples.html
// [Passing IAM roles to Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/security_iam_service-with-iam.html#security-create-iam-role
// [Sign up for Amazon Web Services]: https://docs.aws.amazon.com/quicksight/latest/user/setting-up-aws-sign-up.html
func quicksight_CreateAccountSubscription(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateAccountSubscriptionInput{
		// AccountName: *string, // Required
		// AuthenticationMethod: types.AuthenticationMethodOption, // Required
		// AwsAccountId: *string, // Required
		// NotificationEmail: *string, // Required
	}

	if len(_quicksightAccountName) > 0 {
		input.AccountName = aws.String(_quicksightAccountName)
	}
	if len(_quicksightAuthenticationMethod) > 0 {
		if err := assignInputField(input, "AuthenticationMethod", _quicksightAuthenticationMethod); err != nil {
			log.Errorf("invalid --authentication-method: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNotificationEmail) > 0 {
		input.NotificationEmail = aws.String(_quicksightNotificationEmail)
	}
	if len(_quicksightActiveDirectoryName) > 0 {
		input.ActiveDirectoryName = aws.String(_quicksightActiveDirectoryName)
	}
	if len(_quicksightAdminGroup) > 0 {
		input.AdminGroup = append([]string(nil), _quicksightAdminGroup...)
	}
	if len(_quicksightAdminProGroup) > 0 {
		input.AdminProGroup = append([]string(nil), _quicksightAdminProGroup...)
	}
	if len(_quicksightAuthorGroup) > 0 {
		input.AuthorGroup = append([]string(nil), _quicksightAuthorGroup...)
	}
	if len(_quicksightAuthorProGroup) > 0 {
		input.AuthorProGroup = append([]string(nil), _quicksightAuthorProGroup...)
	}
	if len(_quicksightContactNumber) > 0 {
		input.ContactNumber = aws.String(_quicksightContactNumber)
	}
	if len(_quicksightDirectoryId) > 0 {
		input.DirectoryId = aws.String(_quicksightDirectoryId)
	}
	if len(_quicksightEdition) > 0 {
		if err := assignInputField(input, "Edition", _quicksightEdition); err != nil {
			log.Errorf("invalid --edition: %s", err.Error())
			return
		}
	}
	if len(_quicksightEmailAddress) > 0 {
		input.EmailAddress = aws.String(_quicksightEmailAddress)
	}
	if len(_quicksightFirstName) > 0 {
		input.FirstName = aws.String(_quicksightFirstName)
	}
	if len(_quicksightIAMIdentityCenterInstanceArn) > 0 {
		input.IAMIdentityCenterInstanceArn = aws.String(_quicksightIAMIdentityCenterInstanceArn)
	}
	if len(_quicksightLastName) > 0 {
		input.LastName = aws.String(_quicksightLastName)
	}
	if len(_quicksightReaderGroup) > 0 {
		input.ReaderGroup = append([]string(nil), _quicksightReaderGroup...)
	}
	if len(_quicksightReaderProGroup) > 0 {
		input.ReaderProGroup = append([]string(nil), _quicksightReaderProGroup...)
	}
	if len(_quicksightRealm) > 0 {
		input.Realm = aws.String(_quicksightRealm)
	}

	if resp, err := client.CreateAccountSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an action connector that enables Amazon Quick Sight to connect to
// external services and perform actions. Action connectors support various
// authentication methods and can be configured with specific actions from
// supported connector types like Amazon S3, Salesforce, JIRA.
func quicksight_CreateActionConnector(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateActionConnectorInput{
		// ActionConnectorId: *string, // Required
		// AuthenticationConfig: *types.AuthConfig, // Required
		// AwsAccountId: *string, // Required
		// Name: *string, // Required
		// Type: types.ActionConnectorType, // Required
	}

	if len(_quicksightActionConnectorId) > 0 {
		input.ActionConnectorId = aws.String(_quicksightActionConnectorId)
	}
	if len(_quicksightAuthenticationConfig) > 0 {
		if err := assignInputField(input, "AuthenticationConfig", _quicksightAuthenticationConfig); err != nil {
			log.Errorf("invalid --authentication-config: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightType) > 0 {
		if err := assignInputField(input, "Type", _quicksightType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_quicksightDescription) > 0 {
		input.Description = aws.String(_quicksightDescription)
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightVpcConnectionArn) > 0 {
		input.VpcConnectionArn = aws.String(_quicksightVpcConnectionArn)
	}

	if resp, err := client.CreateActionConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an analysis in Amazon Quick Sight. Analyses can be created either from
// a template or from an AnalysisDefinition .
func quicksight_CreateAnalysis(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateAnalysisInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightDefinition) > 0 {
		if err := assignInputField(input, "Definition", _quicksightDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightFolderArns) > 0 {
		input.FolderArns = append([]string(nil), _quicksightFolderArns...)
	}
	if len(_quicksightParameters) > 0 {
		if err := assignInputField(input, "Parameters", _quicksightParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightSourceEntity) > 0 {
		if err := assignInputField(input, "SourceEntity", _quicksightSourceEntity); err != nil {
			log.Errorf("invalid --source-entity: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightThemeArn) > 0 {
		input.ThemeArn = aws.String(_quicksightThemeArn)
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Quick Sight brand.
func quicksight_CreateBrand(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateBrandInput{
		// AwsAccountId: *string, // Required
		// BrandId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandId) > 0 {
		input.BrandId = aws.String(_quicksightBrandId)
	}
	if len(_quicksightBrandDefinition) > 0 {
		if err := assignInputField(input, "BrandDefinition", _quicksightBrandDefinition); err != nil {
			log.Errorf("invalid --brand-definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBrand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom permissions profile.
func quicksight_CreateCustomPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateCustomPermissionsInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}
	if len(_quicksightCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _quicksightCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dashboard from either a template or directly with a
// DashboardDefinition . To first create a template, see the [CreateTemplate] API operation.
//
// A dashboard is an entity in Amazon Quick Sight that identifies Amazon Quick
// Sight reports, created from analyses. You can share Amazon Quick Sight
// dashboards. With the right permissions, you can create scheduled email reports
// from them. If you have the correct permissions, you can create a dashboard from
// a template that exists in a different Amazon Web Services account.
//
// [CreateTemplate]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateTemplate.html
func quicksight_CreateDashboard(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateDashboardInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightDashboardPublishOptions) > 0 {
		if err := assignInputField(input, "DashboardPublishOptions", _quicksightDashboardPublishOptions); err != nil {
			log.Errorf("invalid --dashboard-publish-options: %s", err.Error())
			return
		}
	}
	if len(_quicksightDefinition) > 0 {
		if err := assignInputField(input, "Definition", _quicksightDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightFolderArns) > 0 {
		input.FolderArns = append([]string(nil), _quicksightFolderArns...)
	}
	if len(_quicksightLinkEntities) > 0 {
		input.LinkEntities = append([]string(nil), _quicksightLinkEntities...)
	}
	if len(_quicksightLinkSharingConfiguration) > 0 {
		if err := assignInputField(input, "LinkSharingConfiguration", _quicksightLinkSharingConfiguration); err != nil {
			log.Errorf("invalid --link-sharing-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightParameters) > 0 {
		if err := assignInputField(input, "Parameters", _quicksightParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightSourceEntity) > 0 {
		if err := assignInputField(input, "SourceEntity", _quicksightSourceEntity); err != nil {
			log.Errorf("invalid --source-entity: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightThemeArn) > 0 {
		input.ThemeArn = aws.String(_quicksightThemeArn)
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}
	if len(_quicksightVersionDescription) > 0 {
		input.VersionDescription = aws.String(_quicksightVersionDescription)
	}

	if resp, err := client.CreateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dataset. This operation doesn't support datasets that include
// uploaded files as a source.
func quicksight_CreateDataSet(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateDataSetInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// ImportMode: types.DataSetImportMode, // Required
		// Name: *string, // Required
		// PhysicalTableMap: map[string]types.PhysicalTable, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightImportMode) > 0 {
		if err := assignInputField(input, "ImportMode", _quicksightImportMode); err != nil {
			log.Errorf("invalid --import-mode: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightPhysicalTableMap) > 0 {
		if err := assignInputField(input, "PhysicalTableMap", _quicksightPhysicalTableMap); err != nil {
			log.Errorf("invalid --physical-table-map: %s", err.Error())
			return
		}
	}
	if len(_quicksightColumnGroups) > 0 {
		if err := assignInputField(input, "ColumnGroups", _quicksightColumnGroups); err != nil {
			log.Errorf("invalid --column-groups: %s", err.Error())
			return
		}
	}
	if len(_quicksightColumnLevelPermissionRules) > 0 {
		if err := assignInputField(input, "ColumnLevelPermissionRules", _quicksightColumnLevelPermissionRules); err != nil {
			log.Errorf("invalid --column-level-permission-rules: %s", err.Error())
			return
		}
	}
	if len(_quicksightDataPrepConfiguration) > 0 {
		if err := assignInputField(input, "DataPrepConfiguration", _quicksightDataPrepConfiguration); err != nil {
			log.Errorf("invalid --data-prep-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightDataSetUsageConfiguration) > 0 {
		if err := assignInputField(input, "DataSetUsageConfiguration", _quicksightDataSetUsageConfiguration); err != nil {
			log.Errorf("invalid --data-set-usage-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightDatasetParameters) > 0 {
		if err := assignInputField(input, "DatasetParameters", _quicksightDatasetParameters); err != nil {
			log.Errorf("invalid --dataset-parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightFieldFolders) > 0 {
		if err := assignInputField(input, "FieldFolders", _quicksightFieldFolders); err != nil {
			log.Errorf("invalid --field-folders: %s", err.Error())
			return
		}
	}
	if len(_quicksightFolderArns) > 0 {
		input.FolderArns = append([]string(nil), _quicksightFolderArns...)
	}
	if len(_quicksightLogicalTableMap) > 0 {
		if err := assignInputField(input, "LogicalTableMap", _quicksightLogicalTableMap); err != nil {
			log.Errorf("invalid --logical-table-map: %s", err.Error())
			return
		}
	}
	if len(_quicksightPerformanceConfiguration) > 0 {
		if err := assignInputField(input, "PerformanceConfiguration", _quicksightPerformanceConfiguration); err != nil {
			log.Errorf("invalid --performance-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRowLevelPermissionDataSet) > 0 {
		if err := assignInputField(input, "RowLevelPermissionDataSet", _quicksightRowLevelPermissionDataSet); err != nil {
			log.Errorf("invalid --row-level-permission-data-set: %s", err.Error())
			return
		}
	}
	if len(_quicksightRowLevelPermissionTagConfiguration) > 0 {
		if err := assignInputField(input, "RowLevelPermissionTagConfiguration", _quicksightRowLevelPermissionTagConfiguration); err != nil {
			log.Errorf("invalid --row-level-permission-tag-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightSemanticModelConfiguration) > 0 {
		if err := assignInputField(input, "SemanticModelConfiguration", _quicksightSemanticModelConfiguration); err != nil {
			log.Errorf("invalid --semantic-model-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightUseAs) > 0 {
		if err := assignInputField(input, "UseAs", _quicksightUseAs); err != nil {
			log.Errorf("invalid --use-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data source.
func quicksight_CreateDataSource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateDataSourceInput{
		// AwsAccountId: *string, // Required
		// DataSourceId: *string, // Required
		// Name: *string, // Required
		// Type: types.DataSourceType, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSourceId) > 0 {
		input.DataSourceId = aws.String(_quicksightDataSourceId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightType) > 0 {
		if err := assignInputField(input, "Type", _quicksightType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_quicksightCredentials) > 0 {
		if err := assignInputField(input, "Credentials", _quicksightCredentials); err != nil {
			log.Errorf("invalid --credentials: %s", err.Error())
			return
		}
	}
	if len(_quicksightDataSourceParameters) > 0 {
		if err := assignInputField(input, "DataSourceParameters", _quicksightDataSourceParameters); err != nil {
			log.Errorf("invalid --data-source-parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightFolderArns) > 0 {
		input.FolderArns = append([]string(nil), _quicksightFolderArns...)
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightSslProperties) > 0 {
		if err := assignInputField(input, "SslProperties", _quicksightSslProperties); err != nil {
			log.Errorf("invalid --ssl-properties: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightVpcConnectionProperties) > 0 {
		if err := assignInputField(input, "VpcConnectionProperties", _quicksightVpcConnectionProperties); err != nil {
			log.Errorf("invalid --vpc-connection-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty shared folder.
func quicksight_CreateFolder(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateFolderInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightFolderType) > 0 {
		if err := assignInputField(input, "FolderType", _quicksightFolderType); err != nil {
			log.Errorf("invalid --folder-type: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightParentFolderArn) > 0 {
		input.ParentFolderArn = aws.String(_quicksightParentFolderArn)
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightSharingModel) > 0 {
		if err := assignInputField(input, "SharingModel", _quicksightSharingModel); err != nil {
			log.Errorf("invalid --sharing-model: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an asset, such as a dashboard, analysis, or dataset into a folder.
func quicksight_CreateFolderMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateFolderMembershipInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
		// MemberId: *string, // Required
		// MemberType: types.MemberType, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightMemberId) > 0 {
		input.MemberId = aws.String(_quicksightMemberId)
	}
	if len(_quicksightMemberType) > 0 {
		if err := assignInputField(input, "MemberType", _quicksightMemberType); err != nil {
			log.Errorf("invalid --member-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFolderMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use the CreateGroup operation to create a group in Quick Sight. You can create
// up to 10,000 groups in a namespace. If you want to create more than 10,000
// groups in a namespace, contact Amazon Web Services Support.
//
// The permissions resource is arn:aws:quicksight:::group/default/ .
//
// The response is a group object.
func quicksight_CreateGroup(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateGroupInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightDescription) > 0 {
		input.Description = aws.String(_quicksightDescription)
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an Amazon Quick Sight user to an Amazon Quick Sight group.
func quicksight_CreateGroupMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateGroupMembershipInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// MemberName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightMemberName) > 0 {
		input.MemberName = aws.String(_quicksightMemberName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.CreateGroupMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an assignment with one specified IAM policy, identified by its Amazon
// Resource Name (ARN). This policy assignment is attached to the specified groups
// or users of Amazon Quick Sight. Assignment names are unique per Amazon Web
// Services account. To avoid overwriting rules in other namespaces, use assignment
// names that are unique.
func quicksight_CreateIAMPolicyAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateIAMPolicyAssignmentInput{
		// AssignmentName: *string, // Required
		// AssignmentStatus: types.AssignmentStatus, // Required
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAssignmentName) > 0 {
		input.AssignmentName = aws.String(_quicksightAssignmentName)
	}
	if len(_quicksightAssignmentStatus) > 0 {
		if err := assignInputField(input, "AssignmentStatus", _quicksightAssignmentStatus); err != nil {
			log.Errorf("invalid --assignment-status: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightIdentities) > 0 {
		if err := assignInputField(input, "Identities", _quicksightIdentities); err != nil {
			log.Errorf("invalid --identities: %s", err.Error())
			return
		}
	}
	if len(_quicksightPolicyArn) > 0 {
		input.PolicyArn = aws.String(_quicksightPolicyArn)
	}

	if resp, err := client.CreateIAMPolicyAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and starts a new SPICE ingestion for a dataset. You can manually
// refresh datasets in an Enterprise edition account 32 times in a 24-hour period.
// You can manually refresh datasets in a Standard edition account 8 times in a
// 24-hour period. Each 24-hour period is measured starting 24 hours before the
// current date and time.
//
// Any ingestions operating on tagged datasets inherit the same tags automatically
// for use in access control. For an example, see [How do I create an IAM policy to control access to Amazon EC2 resources using tags?]in the Amazon Web Services
// Knowledge Center. Tags are visible on the tagged dataset, but not on the
// ingestion resource.
//
// [How do I create an IAM policy to control access to Amazon EC2 resources using tags?]: http://aws.amazon.com/premiumsupport/knowledge-center/iam-ec2-resource-tags/
func quicksight_CreateIngestion(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateIngestionInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// IngestionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightIngestionId) > 0 {
		input.IngestionId = aws.String(_quicksightIngestionId)
	}
	if len(_quicksightIngestionType) > 0 {
		if err := assignInputField(input, "IngestionType", _quicksightIngestionType); err != nil {
			log.Errorf("invalid --ingestion-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// (Enterprise edition only) Creates a new namespace for you to use with Amazon
// Quick Sight.
//
// A namespace allows you to isolate the Quick Sight users and groups that are
// registered for that namespace. Users that access the namespace can share assets
// only with other users or groups in the same namespace. They can't see users and
// groups in other namespaces. You can create a namespace after your Amazon Web
// Services account is subscribed to Quick Sight. The namespace must be unique
// within the Amazon Web Services account. By default, there is a limit of 100
// namespaces per Amazon Web Services account. To increase your limit, create a
// ticket with Amazon Web Services Support.
func quicksight_CreateNamespace(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateNamespaceInput{
		// AwsAccountId: *string, // Required
		// IdentityStore: types.IdentityStore, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightIdentityStore) > 0 {
		if err := assignInputField(input, "IdentityStore", _quicksightIdentityStore); err != nil {
			log.Errorf("invalid --identity-store: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a refresh schedule for a dataset. You can create up to 5 different
// schedules for a single dataset.
func quicksight_CreateRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// Schedule: *types.RefreshSchedule, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _quicksightSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use CreateRoleMembership to add an existing Quick Sight group to an existing
// role.
func quicksight_CreateRoleMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateRoleMembershipInput{
		// AwsAccountId: *string, // Required
		// MemberName: *string, // Required
		// Namespace: *string, // Required
		// Role: types.Role, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMemberName) > 0 {
		input.MemberName = aws.String(_quicksightMemberName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoleMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a template either from a TemplateDefinition or from an existing Quick
// Sight analysis or template. You can use the resulting template to create
// additional dashboards, templates, or analyses.
//
// A template is an entity in Quick Sight that encapsulates the metadata required
// to create an analysis and that you can use to create s dashboard. A template
// adds a layer of abstraction by using placeholders to replace the dataset
// associated with the analysis. You can use templates to create dashboards by
// replacing dataset placeholders with datasets that follow the same schema that
// was used to create the source analysis and template.
func quicksight_CreateTemplate(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateTemplateInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightDefinition) > 0 {
		if err := assignInputField(input, "Definition", _quicksightDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightSourceEntity) > 0 {
		if err := assignInputField(input, "SourceEntity", _quicksightSourceEntity); err != nil {
			log.Errorf("invalid --source-entity: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}
	if len(_quicksightVersionDescription) > 0 {
		input.VersionDescription = aws.String(_quicksightVersionDescription)
	}

	if resp, err := client.CreateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a template alias for a template.
func quicksight_CreateTemplateAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateTemplateAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
		// TemplateVersionNumber: *int64, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightTemplateVersionNumber) > 0 {
		if err := assignInputField(input, "TemplateVersionNumber", _quicksightTemplateVersionNumber); err != nil {
			log.Errorf("invalid --template-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTemplateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a theme.
// A theme is set of configuration options for color and layout. Themes apply to
// analyses and dashboards. For more information, see [Using Themes in Amazon Quick Sight]in the Amazon Quick Sight
// User Guide.
//
// [Using Themes in Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html
func quicksight_CreateTheme(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateThemeInput{
		// AwsAccountId: *string, // Required
		// BaseThemeId: *string, // Required
		// Configuration: *types.ThemeConfiguration, // Required
		// Name: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBaseThemeId) > 0 {
		input.BaseThemeId = aws.String(_quicksightBaseThemeId)
	}
	if len(_quicksightConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _quicksightConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _quicksightPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightVersionDescription) > 0 {
		input.VersionDescription = aws.String(_quicksightVersionDescription)
	}

	if resp, err := client.CreateTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a theme alias for a theme.
func quicksight_CreateThemeAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateThemeAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
		// ThemeVersionNumber: *int64, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightThemeVersionNumber) > 0 {
		if err := assignInputField(input, "ThemeVersionNumber", _quicksightThemeVersionNumber); err != nil {
			log.Errorf("invalid --theme-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThemeAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Q topic.
func quicksight_CreateTopic(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateTopicInput{
		// AwsAccountId: *string, // Required
		// Topic: *types.TopicDetails, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopic) > 0 {
		if err := assignInputField(input, "Topic", _quicksightTopic); err != nil {
			log.Errorf("invalid --topic: %s", err.Error())
			return
		}
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}
	if len(_quicksightCustomInstructions) > 0 {
		if err := assignInputField(input, "CustomInstructions", _quicksightCustomInstructions); err != nil {
			log.Errorf("invalid --custom-instructions: %s", err.Error())
			return
		}
	}
	if len(_quicksightFolderArns) > 0 {
		input.FolderArns = append([]string(nil), _quicksightFolderArns...)
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a topic refresh schedule.
func quicksight_CreateTopicRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateTopicRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DatasetArn: *string, // Required
		// RefreshSchedule: *types.TopicRefreshSchedule, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDatasetArn) > 0 {
		input.DatasetArn = aws.String(_quicksightDatasetArn)
	}
	if len(_quicksightRefreshSchedule) > 0 {
		if err := assignInputField(input, "RefreshSchedule", _quicksightRefreshSchedule); err != nil {
			log.Errorf("invalid --refresh-schedule: %s", err.Error())
			return
		}
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}
	if len(_quicksightDatasetName) > 0 {
		input.DatasetName = aws.String(_quicksightDatasetName)
	}

	if resp, err := client.CreateTopicRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new VPC connection.
func quicksight_CreateVPCConnection(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.CreateVPCConnectionInput{
		// AwsAccountId: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// SecurityGroupIds: []string, // Required
		// SubnetIds: []string, // Required
		// VPCConnectionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightRoleArn) > 0 {
		input.RoleArn = aws.String(_quicksightRoleArn)
	}
	if len(_quicksightSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _quicksightSecurityGroupIds...)
	}
	if len(_quicksightSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _quicksightSubnetIds...)
	}
	if len(_quicksightVPCConnectionId) > 0 {
		input.VPCConnectionId = aws.String(_quicksightVPCConnectionId)
	}
	if len(_quicksightDnsResolvers) > 0 {
		input.DnsResolvers = append([]string(nil), _quicksightDnsResolvers...)
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVPCConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unapplies a custom permissions profile from an account.
func quicksight_DeleteAccountCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteAccountCustomPermissionInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DeleteAccountCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API permanently deletes all Quick Sight customizations for the specified
// Amazon Web Services account and namespace. When you delete account
// customizations:
//
// - All customizations are removed including themes, branding, and visual
// settings
//
// - This action cannot be undone through the API
//
// - Users will see default Quick Sight styling after customizations are deleted
//
// Before proceeding: Ensure you have backups of any custom themes or branding
// elements you may want to recreate.
//
// Deletes all Amazon Quick Sight customizations for the specified Amazon Web
// Services account and Quick Sight namespace.
func quicksight_DeleteAccountCustomization(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteAccountCustomizationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DeleteAccountCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deleting your Quick Sight account subscription has permanent, irreversible
// consequences across all Amazon Web Services regions:
//
// - Global deletion – Running this operation from any single region will delete
// your Quick Sight account and all data in every Amazon Web Services region where
// you have Quick Sight resources.
//
// - Complete data loss – All dashboards, analyses, datasets, data sources, and
// custom visuals will be permanently deleted across all regions.
//
// - Embedded content failure – All embedded dashboards and visuals in your
// applications will immediately stop working and display errors to end users.
//
// - Shared resources removed – All shared dashboards, folders, and resources
// will become inaccessible to other users and external recipients.
//
// - User access terminated – All Quick Sight users in your account will lose
// access immediately, including authors, readers, and administrators.
//
// - No recovery possible – Once deleted, your Quick Sight account and all
// associated data cannot be restored.
//
// Consider exporting critical dashboards and data before proceeding with account
// deletion.
//
// Use the DeleteAccountSubscription operation to delete an Quick Sight account.
// This operation will result in an error message if you have configured your
// account termination protection settings to True . To change this setting and
// delete your account, call the UpdateAccountSettings API and set the value of
// the TerminationProtectionEnabled parameter to False , then make another call to
// the DeleteAccountSubscription API.
func quicksight_DeleteAccountSubscription(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteAccountSubscriptionInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DeleteAccountSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Hard deletes an action connector, making it unrecoverable. This operation
// removes the connector and all its associated configurations. Any resources
// currently using this action connector will no longer be able to perform actions
// through it.
func quicksight_DeleteActionConnector(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteActionConnectorInput{
		// ActionConnectorId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightActionConnectorId) > 0 {
		input.ActionConnectorId = aws.String(_quicksightActionConnectorId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DeleteActionConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an analysis from Amazon Quick Sight. You can optionally include a
// recovery window during which you can restore the analysis. If you don't specify
// a recovery window value, the operation defaults to 30 days. Amazon Quick Sight
// attaches a DeletionTime stamp to the response that specifies the end of the
// recovery window. At the end of the recovery window, Amazon Quick Sight deletes
// the analysis permanently.
//
// At any time before recovery window ends, you can use the RestoreAnalysis API
// operation to remove the DeletionTime stamp and cancel the deletion of the
// analysis. The analysis remains visible in the API until it's deleted, so you can
// describe it but you can't make a template from it.
//
// An analysis that's scheduled for deletion isn't accessible in the Amazon Quick
// Sight console. To access it in the console, restore it. Deleting an analysis
// doesn't delete the dashboards that you publish from it.
func quicksight_DeleteAnalysis(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteAnalysisInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightForceDeleteWithoutRecovery) > 0 {
		if err := assignInputField(input, "ForceDeleteWithoutRecovery", _quicksightForceDeleteWithoutRecovery); err != nil {
			log.Errorf("invalid --force-delete-without-recovery: %s", err.Error())
			return
		}
	}
	if len(_quicksightRecoveryWindowInDays) > 0 {
		if err := assignInputField(input, "RecoveryWindowInDays", _quicksightRecoveryWindowInDays); err != nil {
			log.Errorf("invalid --recovery-window-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API permanently deletes the specified Quick Sight brand. When you delete a
// brand:
//
// - The brand and all its associated branding elements are permanently removed
//
// - Any applications or dashboards using this brand will revert to default
// styling
//
// - This action cannot be undone through the API
//
// Before proceeding: Verify that the brand is no longer needed and consider the
// impact on any applications currently using this brand.
//
// Deletes an Quick Sight brand.
func quicksight_DeleteBrand(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteBrandInput{
		// AwsAccountId: *string, // Required
		// BrandId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandId) > 0 {
		input.BrandId = aws.String(_quicksightBrandId)
	}

	if resp, err := client.DeleteBrand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a brand assignment.
func quicksight_DeleteBrandAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteBrandAssignmentInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DeleteBrandAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom permissions profile.
func quicksight_DeleteCustomPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteCustomPermissionsInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}

	if resp, err := client.DeleteCustomPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dashboard.
func quicksight_DeleteDashboard(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteDashboardInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset.
func quicksight_DeleteDataSet(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteDataSetInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}

	if resp, err := client.DeleteDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the dataset refresh properties of the dataset.
func quicksight_DeleteDataSetRefreshProperties(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteDataSetRefreshPropertiesInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}

	if resp, err := client.DeleteDataSetRefreshProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the data source permanently. This operation breaks all the datasets
// that reference the deleted data source.
func quicksight_DeleteDataSource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteDataSourceInput{
		// AwsAccountId: *string, // Required
		// DataSourceId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSourceId) > 0 {
		input.DataSourceId = aws.String(_quicksightDataSourceId)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a linked Amazon Q Business application from an Quick Sight account
func quicksight_DeleteDefaultQBusinessApplication(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteDefaultQBusinessApplicationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DeleteDefaultQBusinessApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an empty folder.
func quicksight_DeleteFolder(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteFolderInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}

	if resp, err := client.DeleteFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an asset, such as a dashboard, analysis, or dataset, from a folder.
func quicksight_DeleteFolderMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteFolderMembershipInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
		// MemberId: *string, // Required
		// MemberType: types.MemberType, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightMemberId) > 0 {
		input.MemberId = aws.String(_quicksightMemberId)
	}
	if len(_quicksightMemberType) > 0 {
		if err := assignInputField(input, "MemberType", _quicksightMemberType); err != nil {
			log.Errorf("invalid --member-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFolderMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a user group from Amazon Quick Sight.
func quicksight_DeleteGroup(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteGroupInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a user from a group so that the user is no longer a member of the group.
func quicksight_DeleteGroupMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteGroupMembershipInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// MemberName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightMemberName) > 0 {
		input.MemberName = aws.String(_quicksightMemberName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DeleteGroupMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing IAM policy assignment.
func quicksight_DeleteIAMPolicyAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteIAMPolicyAssignmentInput{
		// AssignmentName: *string, // Required
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAssignmentName) > 0 {
		input.AssignmentName = aws.String(_quicksightAssignmentName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DeleteIAMPolicyAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all access scopes and authorized targets that are associated with a
// service from the Quick Sight IAM Identity Center application.
//
// This operation is only supported for Quick Sight accounts that use IAM Identity
// Center.
func quicksight_DeleteIdentityPropagationConfig(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteIdentityPropagationConfigInput{
		// AwsAccountId: *string, // Required
		// Service: types.ServiceType, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightService) > 0 {
		if err := assignInputField(input, "Service", _quicksightService); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteIdentityPropagationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a namespace and the users and groups that are associated with the
// namespace. This is an asynchronous process. Assets including dashboards,
// analyses, datasets and data sources are not deleted. To delete these assets, you
// use the API operations for the relevant asset.
func quicksight_DeleteNamespace(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteNamespaceInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DeleteNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a refresh schedule from a dataset.
func quicksight_DeleteRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// ScheduleId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightScheduleId) > 0 {
		input.ScheduleId = aws.String(_quicksightScheduleId)
	}

	if resp, err := client.DeleteRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes custom permissions from the role.
func quicksight_DeleteRoleCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteRoleCustomPermissionInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// Role: types.Role, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRoleCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a group from a role.
func quicksight_DeleteRoleMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteRoleMembershipInput{
		// AwsAccountId: *string, // Required
		// MemberName: *string, // Required
		// Namespace: *string, // Required
		// Role: types.Role, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMemberName) > 0 {
		input.MemberName = aws.String(_quicksightMemberName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRoleMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a template.
func quicksight_DeleteTemplate(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteTemplateInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the item that the specified template alias points to. If you provide a
// specific alias, you delete the version of the template that the alias points to.
func quicksight_DeleteTemplateAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteTemplateAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}

	if resp, err := client.DeleteTemplateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a theme.
func quicksight_DeleteTheme(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteThemeInput{
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the version of the theme that the specified theme alias points to. If
// you provide a specific alias, you delete the version of the theme that the alias
// points to.
func quicksight_DeleteThemeAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteThemeAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}

	if resp, err := client.DeleteThemeAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a topic.
func quicksight_DeleteTopic(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteTopicInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.DeleteTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a topic refresh schedule.
func quicksight_DeleteTopicRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteTopicRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DatasetId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDatasetId) > 0 {
		input.DatasetId = aws.String(_quicksightDatasetId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.DeleteTopicRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Amazon Quick Sight user that is associated with the identity of the
// IAM user or role that's making the call. The IAM user isn't deleted as a result
// of this call.
func quicksight_DeleteUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteUserInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user identified by its principal ID.
func quicksight_DeleteUserByPrincipalId(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteUserByPrincipalIdInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// PrincipalId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightPrincipalId) > 0 {
		input.PrincipalId = aws.String(_quicksightPrincipalId)
	}

	if resp, err := client.DeleteUserByPrincipalId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom permissions profile from a user.
func quicksight_DeleteUserCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteUserCustomPermissionInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}

	if resp, err := client.DeleteUserCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a VPC connection.
func quicksight_DeleteVPCConnection(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DeleteVPCConnectionInput{
		// AwsAccountId: *string, // Required
		// VPCConnectionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightVPCConnectionId) > 0 {
		input.VPCConnectionId = aws.String(_quicksightVPCConnectionId)
	}

	if resp, err := client.DeleteVPCConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the custom permissions profile that is applied to an account.
func quicksight_DescribeAccountCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAccountCustomPermissionInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAccountCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the customizations associated with the provided Amazon Web Services
// account and Amazon Quick Sight namespace. The Quick Sight console evaluates
// which customizations to apply by running this API operation with the Resolved
// flag included.
//
// To determine what customizations display when you run this command, it can help
// to visualize the relationship of the entities involved.
//
// - Amazon Web Services account - The Amazon Web Services account exists at the
// top of the hierarchy. It has the potential to use all of the Amazon Web Services
// Regions and Amazon Web Services Services. When you subscribe to Quick Sight, you
// choose one Amazon Web Services Region to use as your home Region. That's where
// your free SPICE capacity is located. You can use Quick Sight in any supported
// Amazon Web Services Region.
//
// - Amazon Web Services Region - You can sign in to Quick Sight in any Amazon
// Web Services Region. If you have a user directory, it resides in us-east-1,
// which is US East (N. Virginia). Generally speaking, these users have access to
// Quick Sight in any Amazon Web Services Region, unless they are constrained to a
// namespace.
//
// # To run the command in a different Amazon Web Services Region, you change your
//
// Region settings. If you're using the CLI, you can use one of the following
// options:
//
// - Use [command line options].
//
// - Use [named profiles].
//
// - Run aws configure to change your default Amazon Web Services Region. Use
// Enter to key the same settings for your keys. For more information, see [Configuring the CLI].
//
// - Namespace - A Quick Sight namespace is a partition that contains users and
// assets (data sources, datasets, dashboards, and so on). To access assets that
// are in a specific namespace, users and groups must also be part of the same
// namespace. People who share a namespace are completely isolated from users and
// assets in other namespaces, even if they are in the same Amazon Web Services
// account and Amazon Web Services Region.
//
// - Applied customizations - Quick Sight customizations can apply to an Amazon
// Web Services account or to a namespace. Settings that you apply to a namespace
// override settings that you apply to an Amazon Web Services account.
//
// [named profiles]: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
// [Configuring the CLI]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
// [command line options]: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-options.html
func quicksight_DescribeAccountCustomization(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAccountCustomizationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightResolved) > 0 {
		if err := assignInputField(input, "Resolved", _quicksightResolved); err != nil {
			log.Errorf("invalid --resolved: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAccountCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the settings that were used when your Quick Sight subscription was
// first created in this Amazon Web Services account.
func quicksight_DescribeAccountSettings(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAccountSettingsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use the DescribeAccountSubscription operation to receive a description of an
// Quick Sight account's subscription. A successful API call returns an AccountInfo
// object that includes an account's name, subscription status, authentication
// type, edition, and notification email address.
func quicksight_DescribeAccountSubscription(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAccountSubscriptionInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAccountSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about an action connector, including its
// configuration, authentication settings, enabled actions, and current status.
func quicksight_DescribeActionConnector(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeActionConnectorInput{
		// ActionConnectorId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightActionConnectorId) > 0 {
		input.ActionConnectorId = aws.String(_quicksightActionConnectorId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeActionConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the permissions configuration for an action connector, showing which
// users, groups, and namespaces have access and what operations they can perform.
func quicksight_DescribeActionConnectorPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeActionConnectorPermissionsInput{
		// ActionConnectorId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightActionConnectorId) > 0 {
		input.ActionConnectorId = aws.String(_quicksightActionConnectorId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeActionConnectorPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a summary of the metadata for an analysis.
func quicksight_DescribeAnalysis(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAnalysisInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a detailed description of the definition of an analysis.
// If you do not need to know details about the content of an Analysis, for
// instance if you are trying to check the status of a recently created or updated
// Analysis, use the [DescribeAnalysis]DescribeAnalysis instead.
//
// [DescribeAnalysis]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeAnalysis.html
func quicksight_DescribeAnalysisDefinition(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAnalysisDefinitionInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAnalysisDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the read and write permissions for an analysis.
func quicksight_DescribeAnalysisPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAnalysisPermissionsInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAnalysisPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing export job.
// Poll job descriptions after a job starts to know the status of the job. When a
// job succeeds, a URL is provided to download the exported assets' data from.
// Download URLs are valid for five minutes after they are generated. You can call
// the DescribeAssetBundleExportJob API for a new download URL as needed.
//
// Job descriptions are available for 14 days after the job starts.
func quicksight_DescribeAssetBundleExportJob(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAssetBundleExportJobInput{
		// AssetBundleExportJobId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAssetBundleExportJobId) > 0 {
		input.AssetBundleExportJobId = aws.String(_quicksightAssetBundleExportJobId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAssetBundleExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing import job.
// Poll job descriptions after starting a job to know when it has succeeded or
// failed. Job descriptions are available for 14 days after job starts.
func quicksight_DescribeAssetBundleImportJob(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeAssetBundleImportJobInput{
		// AssetBundleImportJobId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAssetBundleImportJobId) > 0 {
		input.AssetBundleImportJobId = aws.String(_quicksightAssetBundleImportJobId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeAssetBundleImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a brand.
func quicksight_DescribeBrand(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeBrandInput{
		// AwsAccountId: *string, // Required
		// BrandId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandId) > 0 {
		input.BrandId = aws.String(_quicksightBrandId)
	}
	if len(_quicksightVersionId) > 0 {
		input.VersionId = aws.String(_quicksightVersionId)
	}

	if resp, err := client.DescribeBrand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a brand assignment.
func quicksight_DescribeBrandAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeBrandAssignmentInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeBrandAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the published version of the brand.
func quicksight_DescribeBrandPublishedVersion(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeBrandPublishedVersionInput{
		// AwsAccountId: *string, // Required
		// BrandId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandId) > 0 {
		input.BrandId = aws.String(_quicksightBrandId)
	}

	if resp, err := client.DescribeBrandPublishedVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a custom permissions profile.
func quicksight_DescribeCustomPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeCustomPermissionsInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}

	if resp, err := client.DescribeCustomPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a summary for a dashboard.
func quicksight_DescribeDashboard(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDashboardInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a detailed description of the definition of a dashboard.
// If you do not need to know details about the content of a dashboard, for
// instance if you are trying to check the status of a recently created or updated
// dashboard, use the [DescribeDashboard]DescribeDashboard instead.
//
// [DescribeDashboard]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeDashboard.html
func quicksight_DescribeDashboardDefinition(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDashboardDefinitionInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeDashboardDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes read and write permissions for a dashboard.
func quicksight_DescribeDashboardPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDashboardPermissionsInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}

	if resp, err := client.DescribeDashboardPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing snapshot job.
// Poll job descriptions after a job starts to know the status of the job. For
// information on available status codes, see JobStatus .
//
// # Registered user support
//
// This API can be called as before to get status of a job started by the same
// Quick Sight user.
//
// # Possible error scenarios
//
// Request will fail with an Access Denied error in the following scenarios:
//
// - The credentials have expired.
//
// - Job has been started by a different user.
//
// - Impersonated Quick Sight user doesn't have access to the specified
// dashboard in the job.
func quicksight_DescribeDashboardSnapshotJob(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDashboardSnapshotJobInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// SnapshotJobId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightSnapshotJobId) > 0 {
		input.SnapshotJobId = aws.String(_quicksightSnapshotJobId)
	}

	if resp, err := client.DescribeDashboardSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the result of an existing snapshot job that has finished running.
// A finished snapshot job will return a COMPLETED or FAILED status when you poll
// the job with a DescribeDashboardSnapshotJob API call.
//
// If the job has not finished running, this operation returns a message that says
// Dashboard Snapshot Job with id has not reached a terminal state. .
//
// # Registered user support
//
// This API can be called as before to get the result of a job started by the same
// Quick Sight user. The result for the user will be returned in RegisteredUsers
// response attribute. The attribute will contain a list with at most one object in
// it.
//
// # Possible error scenarios
//
// The request fails with an Access Denied error in the following scenarios:
//
// - The credentials have expired.
//
// - The job was started by a different user.
//
// - The registered user doesn't have access to the specified dashboard.
//
// The request succeeds but the job fails in the following scenarios:
//
// - DASHBOARD_ACCESS_DENIED - The registered user lost access to the dashboard.
//
// - CAPABILITY_RESTRICTED - The registered user is restricted from exporting
// data in all selected formats.
//
// The request succeeds but the response contains an error code in the following
// scenarios:
//
// - CAPABILITY_RESTRICTED - The registered user is restricted from exporting
// data in some selected formats.
//
// - RLS_CHANGED - Row-level security settings have changed. Re-run the job with
// current settings.
//
// - CLS_CHANGED - Column-level security settings have changed. Re-run the job
// with current settings.
//
// - DATASET_DELETED - The dataset has been deleted. Verify the dataset exists
// before re-running the job.
func quicksight_DescribeDashboardSnapshotJobResult(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDashboardSnapshotJobResultInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// SnapshotJobId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightSnapshotJobId) > 0 {
		input.SnapshotJobId = aws.String(_quicksightSnapshotJobId)
	}

	if resp, err := client.DescribeDashboardSnapshotJobResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing dashboard QA configuration.
func quicksight_DescribeDashboardsQAConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDashboardsQAConfigurationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeDashboardsQAConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a dataset. This operation doesn't support datasets that include
// uploaded files as a source.
func quicksight_DescribeDataSet(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDataSetInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}

	if resp, err := client.DescribeDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the permissions on a dataset.
// The permissions resource is
// arn:aws:quicksight:region:aws-account-id:dataset/data-set-id .
func quicksight_DescribeDataSetPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDataSetPermissionsInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}

	if resp, err := client.DescribeDataSetPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the refresh properties of a dataset.
func quicksight_DescribeDataSetRefreshProperties(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDataSetRefreshPropertiesInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}

	if resp, err := client.DescribeDataSetRefreshProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a data source.
func quicksight_DescribeDataSource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDataSourceInput{
		// AwsAccountId: *string, // Required
		// DataSourceId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSourceId) > 0 {
		input.DataSourceId = aws.String(_quicksightDataSourceId)
	}

	if resp, err := client.DescribeDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the resource permissions for a data source.
func quicksight_DescribeDataSourcePermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDataSourcePermissionsInput{
		// AwsAccountId: *string, // Required
		// DataSourceId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSourceId) > 0 {
		input.DataSourceId = aws.String(_quicksightDataSourceId)
	}

	if resp, err := client.DescribeDataSourcePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a Amazon Q Business application that is linked to an Quick Sight
// account.
func quicksight_DescribeDefaultQBusinessApplication(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeDefaultQBusinessApplicationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DescribeDefaultQBusinessApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a folder.
func quicksight_DescribeFolder(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeFolderInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}

	if resp, err := client.DescribeFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes permissions for a folder.
func quicksight_DescribeFolderPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeFolderPermissionsInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFolderPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.DescribeFolderPermissionsOutput
	p := quicksight.NewDescribeFolderPermissionsPaginator(client, input)
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

// Describes the folder resolved permissions. Permissions consists of both folder
// direct permissions and the inherited permissions from the ancestor folders.
func quicksight_DescribeFolderResolvedPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeFolderResolvedPermissionsInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFolderResolvedPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.DescribeFolderResolvedPermissionsOutput
	p := quicksight.NewDescribeFolderResolvedPermissionsPaginator(client, input)
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

// Returns an Amazon Quick Sight group's description and Amazon Resource Name
// (ARN).
func quicksight_DescribeGroup(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeGroupInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DescribeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use the DescribeGroupMembership operation to determine if a user is a member of
// the specified group. If the user exists and is a member of the specified group,
// an associated GroupMember object is returned.
func quicksight_DescribeGroupMembership(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeGroupMembershipInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// MemberName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightMemberName) > 0 {
		input.MemberName = aws.String(_quicksightMemberName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DescribeGroupMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing IAM policy assignment, as specified by the assignment
// name.
func quicksight_DescribeIAMPolicyAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeIAMPolicyAssignmentInput{
		// AssignmentName: *string, // Required
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAssignmentName) > 0 {
		input.AssignmentName = aws.String(_quicksightAssignmentName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DescribeIAMPolicyAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a SPICE ingestion.
func quicksight_DescribeIngestion(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeIngestionInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// IngestionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightIngestionId) > 0 {
		input.IngestionId = aws.String(_quicksightIngestionId)
	}

	if resp, err := client.DescribeIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a summary and status of IP rules.
func quicksight_DescribeIpRestriction(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeIpRestrictionInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeIpRestriction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes all customer managed key registrations in a Quick Sight account.
func quicksight_DescribeKeyRegistration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeKeyRegistrationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDefaultKeyOnly) > 0 {
		if err := assignInputField(input, "DefaultKeyOnly", _quicksightDefaultKeyOnly); err != nil {
			log.Errorf("invalid --default-key-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeKeyRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the current namespace.
func quicksight_DescribeNamespace(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeNamespaceInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DescribeNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a personalization configuration.
func quicksight_DescribeQPersonalizationConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeQPersonalizationConfigurationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeQPersonalizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the state of a Quick Sight Q Search configuration.
func quicksight_DescribeQuickSightQSearchConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeQuickSightQSearchConfigurationInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}

	if resp, err := client.DescribeQuickSightQSearchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a summary of a refresh schedule.
func quicksight_DescribeRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// ScheduleId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightScheduleId) > 0 {
		input.ScheduleId = aws.String(_quicksightScheduleId)
	}

	if resp, err := client.DescribeRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes all custom permissions that are mapped to a role.
func quicksight_DescribeRoleCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeRoleCustomPermissionInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// Role: types.Role, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeRoleCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the self-upgrade configuration for a Quick Suite account.
func quicksight_DescribeSelfUpgradeConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeSelfUpgradeConfigurationInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.DescribeSelfUpgradeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a template's metadata.
func quicksight_DescribeTemplate(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTemplateInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the template alias for a template.
func quicksight_DescribeTemplateAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTemplateAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}

	if resp, err := client.DescribeTemplateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a detailed description of the definition of a template.
// If you do not need to know details about the content of a template, for
// instance if you are trying to check the status of a recently created or updated
// template, use the [DescribeTemplate]DescribeTemplate instead.
//
// [DescribeTemplate]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeTemplate.html
func quicksight_DescribeTemplateDefinition(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTemplateDefinitionInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTemplateDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes read and write permissions on a template.
func quicksight_DescribeTemplatePermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTemplatePermissionsInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}

	if resp, err := client.DescribeTemplatePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a theme.
func quicksight_DescribeTheme(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeThemeInput{
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the alias for a theme.
func quicksight_DescribeThemeAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeThemeAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}

	if resp, err := client.DescribeThemeAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the read and write permissions for a theme.
func quicksight_DescribeThemePermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeThemePermissionsInput{
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}

	if resp, err := client.DescribeThemePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a topic.
func quicksight_DescribeTopic(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTopicInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.DescribeTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the permissions of a topic.
func quicksight_DescribeTopicPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTopicPermissionsInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.DescribeTopicPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the status of a topic refresh.
func quicksight_DescribeTopicRefresh(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTopicRefreshInput{
		// AwsAccountId: *string, // Required
		// RefreshId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightRefreshId) > 0 {
		input.RefreshId = aws.String(_quicksightRefreshId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.DescribeTopicRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a topic refresh schedule.
func quicksight_DescribeTopicRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeTopicRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DatasetId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDatasetId) > 0 {
		input.DatasetId = aws.String(_quicksightDatasetId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.DescribeTopicRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a user, given the user name.
func quicksight_DescribeUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeUserInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a VPC connection.
func quicksight_DescribeVPCConnection(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.DescribeVPCConnectionInput{
		// AwsAccountId: *string, // Required
		// VPCConnectionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightVPCConnectionId) > 0 {
		input.VPCConnectionId = aws.String(_quicksightVPCConnectionId)
	}

	if resp, err := client.DescribeVPCConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an embed URL that you can use to embed an Amazon Quick Suite
// dashboard or visual in your website, without having to register any reader
// users. Before you use this action, make sure that you have configured the
// dashboards and permissions.
//
// The following rules apply to the generated URL:
//
// - It contains a temporary bearer token. It is valid for 5 minutes after it is
// generated. Once redeemed within this period, it cannot be re-used again.
//
// - The URL validity period should not be confused with the actual session
// lifetime that can be customized using the [SessionLifetimeInMinutes]parameter. The resulting user
// session is valid for 15 minutes (minimum) to 10 hours (maximum). The default
// session duration is 10 hours.
//
// - You are charged only when the URL is used or there is interaction with
// Amazon Quick Suite.
//
// For more information, see [Embedded Analytics] in the Amazon Quick Suite User Guide.
//
// For more information about the high-level steps for embedding and for an
// interactive demo of the ways you can customize embedding, visit the [Amazon Quick Suite Developer Portal].
//
// [Embedded Analytics]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html
// [Amazon Quick Suite Developer Portal]: https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html
// [SessionLifetimeInMinutes]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GenerateEmbedUrlForAnonymousUser.html#QS-GenerateEmbedUrlForAnonymousUser-request-SessionLifetimeInMinutes
func quicksight_GenerateEmbedUrlForAnonymousUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GenerateEmbedUrlForAnonymousUserInput{
		// AuthorizedResourceArns: []string, // Required
		// AwsAccountId: *string, // Required
		// ExperienceConfiguration: *types.AnonymousUserEmbeddingExperienceConfiguration, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAuthorizedResourceArns) > 0 {
		input.AuthorizedResourceArns = append([]string(nil), _quicksightAuthorizedResourceArns...)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightExperienceConfiguration) > 0 {
		if err := assignInputField(input, "ExperienceConfiguration", _quicksightExperienceConfiguration); err != nil {
			log.Errorf("invalid --experience-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightAllowedDomains) > 0 {
		input.AllowedDomains = append([]string(nil), _quicksightAllowedDomains...)
	}
	if len(_quicksightSessionLifetimeInMinutes) > 0 {
		if err := assignInputField(input, "SessionLifetimeInMinutes", _quicksightSessionLifetimeInMinutes); err != nil {
			log.Errorf("invalid --session-lifetime-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_quicksightSessionTags) > 0 {
		if err := assignInputField(input, "SessionTags", _quicksightSessionTags); err != nil {
			log.Errorf("invalid --session-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateEmbedUrlForAnonymousUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an embed URL that you can use to embed an Amazon Quick Suite
// experience in your website. This action can be used for any type of user
// registered in an Amazon Quick Suite account. Before you use this action, make
// sure that you have configured the relevant Amazon Quick Suite resource and
// permissions.
//
// The following rules apply to the generated URL:
//
// - It contains a temporary bearer token. It is valid for 5 minutes after it is
// generated. Once redeemed within this period, it cannot be re-used again.
//
// - The URL validity period should not be confused with the actual session
// lifetime that can be customized using the [SessionLifetimeInMinutes]parameter.
//
// # The resulting user session is valid for 15 minutes (minimum) to 10 hours
//
// (maximum). The default session duration is 10 hours.
//
// - You are charged only when the URL is used or there is interaction with
// Amazon Quick Suite.
//
// For more information, see [Embedded Analytics] in the Amazon Quick Suite User Guide.
//
// For more information about the high-level steps for embedding and for an
// interactive demo of the ways you can customize embedding, visit the [Amazon Quick Suite Developer Portal].
//
// [Embedded Analytics]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html
// [Amazon Quick Suite Developer Portal]: https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html
// [SessionLifetimeInMinutes]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GenerateEmbedUrlForRegisteredUser.html#QS-GenerateEmbedUrlForRegisteredUser-request-SessionLifetimeInMinutes
func quicksight_GenerateEmbedUrlForRegisteredUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GenerateEmbedUrlForRegisteredUserInput{
		// AwsAccountId: *string, // Required
		// ExperienceConfiguration: *types.RegisteredUserEmbeddingExperienceConfiguration, // Required
		// UserArn: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightExperienceConfiguration) > 0 {
		if err := assignInputField(input, "ExperienceConfiguration", _quicksightExperienceConfiguration); err != nil {
			log.Errorf("invalid --experience-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightUserArn) > 0 {
		input.UserArn = aws.String(_quicksightUserArn)
	}
	if len(_quicksightAllowedDomains) > 0 {
		input.AllowedDomains = append([]string(nil), _quicksightAllowedDomains...)
	}
	if len(_quicksightSessionLifetimeInMinutes) > 0 {
		if err := assignInputField(input, "SessionLifetimeInMinutes", _quicksightSessionLifetimeInMinutes); err != nil {
			log.Errorf("invalid --session-lifetime-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateEmbedUrlForRegisteredUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an embed URL that you can use to embed an Amazon Quick Sight
// experience in your website. This action can be used for any type of user that is
// registered in an Amazon Quick Sight account that uses IAM Identity Center for
// authentication. This API requires [identity-enhanced IAM Role sessions]for the authenticated user that the API call
// is being made for.
//
// This API uses [trusted identity propagation] to ensure that an end user is authenticated and receives the
// embed URL that is specific to that user. The IAM Identity Center application
// that the user has logged into needs to have [trusted Identity Propagation enabled for Amazon Quick Sight]with the scope value set to
// quicksight:read . Before you use this action, make sure that you have configured
// the relevant Amazon Quick Sight resource and permissions.
//
// [identity-enhanced IAM Role sessions]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-overview.html#types-identity-enhanced-iam-role-sessions
// [trusted Identity Propagation enabled for Amazon Quick Sight]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-using-customermanagedapps-specify-trusted-apps.html
// [trusted identity propagation]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation.html
func quicksight_GenerateEmbedUrlForRegisteredUserWithIdentity(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GenerateEmbedUrlForRegisteredUserWithIdentityInput{
		// AwsAccountId: *string, // Required
		// ExperienceConfiguration: *types.RegisteredUserEmbeddingExperienceConfiguration, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightExperienceConfiguration) > 0 {
		if err := assignInputField(input, "ExperienceConfiguration", _quicksightExperienceConfiguration); err != nil {
			log.Errorf("invalid --experience-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightAllowedDomains) > 0 {
		input.AllowedDomains = append([]string(nil), _quicksightAllowedDomains...)
	}
	if len(_quicksightSessionLifetimeInMinutes) > 0 {
		if err := assignInputField(input, "SessionLifetimeInMinutes", _quicksightSessionLifetimeInMinutes); err != nil {
			log.Errorf("invalid --session-lifetime-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateEmbedUrlForRegisteredUserWithIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a temporary session URL and authorization code(bearer token) that you
// can use to embed an Amazon Quick Sight read-only dashboard in your website or
// application. Before you use this command, make sure that you have configured the
// dashboards and permissions.
//
// Currently, you can use GetDashboardEmbedURL only from the server, not from the
// user's browser. The following rules apply to the generated URL:
//
// - They must be used together.
//
// - They can be used one time only.
//
// - They are valid for 5 minutes after you run this command.
//
// - You are charged only when the URL is used or there is interaction with
// Quick Suite.
//
// - The resulting user session is valid for 15 minutes (default) up to 10 hours
// (maximum). You can use the optional SessionLifetimeInMinutes parameter to
// customize session duration.
//
// For more information, see [Embedding Analytics Using GetDashboardEmbedUrl] in the Amazon Quick Suite User Guide.
//
// For more information about the high-level steps for embedding and for an
// interactive demo of the ways you can customize embedding, visit the [Amazon Quick Suite Developer Portal].
//
// [Amazon Quick Suite Developer Portal]: https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html
// [Embedding Analytics Using GetDashboardEmbedUrl]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics-deprecated.html
func quicksight_GetDashboardEmbedUrl(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GetDashboardEmbedUrlInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// IdentityType: types.EmbeddingIdentityType, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _quicksightIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_quicksightAdditionalDashboardIds) > 0 {
		input.AdditionalDashboardIds = append([]string(nil), _quicksightAdditionalDashboardIds...)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightResetDisabled) > 0 {
		if err := assignInputField(input, "ResetDisabled", _quicksightResetDisabled); err != nil {
			log.Errorf("invalid --reset-disabled: %s", err.Error())
			return
		}
	}
	if len(_quicksightSessionLifetimeInMinutes) > 0 {
		if err := assignInputField(input, "SessionLifetimeInMinutes", _quicksightSessionLifetimeInMinutes); err != nil {
			log.Errorf("invalid --session-lifetime-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_quicksightStatePersistenceEnabled) > 0 {
		if err := assignInputField(input, "StatePersistenceEnabled", _quicksightStatePersistenceEnabled); err != nil {
			log.Errorf("invalid --state-persistence-enabled: %s", err.Error())
			return
		}
	}
	if len(_quicksightUndoRedoDisabled) > 0 {
		if err := assignInputField(input, "UndoRedoDisabled", _quicksightUndoRedoDisabled); err != nil {
			log.Errorf("invalid --undo-redo-disabled: %s", err.Error())
			return
		}
	}
	if len(_quicksightUserArn) > 0 {
		input.UserArn = aws.String(_quicksightUserArn)
	}

	if resp, err := client.GetDashboardEmbedUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata of a flow, not including its definition specifying the
// steps.
func quicksight_GetFlowMetadata(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GetFlowMetadataInput{
		// AwsAccountId: *string, // Required
		// FlowId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFlowId) > 0 {
		input.FlowId = aws.String(_quicksightFlowId)
	}

	if resp, err := client.GetFlowMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get permissions for a flow.
func quicksight_GetFlowPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GetFlowPermissionsInput{
		// AwsAccountId: *string, // Required
		// FlowId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFlowId) > 0 {
		input.FlowId = aws.String(_quicksightFlowId)
	}

	if resp, err := client.GetFlowPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the identity context for a Quick Sight user in a specified namespace,
// allowing you to obtain identity tokens that can be used with identity-enhanced
// IAM role sessions to call identity-aware APIs.
//
// # Currently, you can call the following APIs with identity-enhanced Credentials
//
// [StartDashboardSnapshotJob]
//
// [DescribeDashboardSnapshotJob]
//
// [DescribeDashboardSnapshotJobResult]
//
// # Supported Authentication Methods
//
// This API supports Quick Sight native users, IAM federated users, and Active
// Directory users. For Quick Sight users authenticated by Amazon Web Services
// Identity Center, see [Identity Center documentation on identity-enhanced IAM role sessions].
//
// # Supported Regions
//
// The GetIdentityContext API works only in regions that support at least one of
// these identity types:
//
// - Amazon Quick Sight native identity
//
// - IAM federated identity
//
// - Active Directory
//
// To use this API successfully, call it in the same region where your user's
// identity resides. For example, if your user's identity is in us-east-1, make the
// API call in us-east-1. For more information about managing identities in Amazon
// Quick Sight, see [Identity and access management in Amazon Quick Sight]in the Amazon Quick Sight User Guide.
//
// # Getting Identity-Enhanced Credentials
//
// To obtain identity-enhanced credentials, follow these steps:
//
// - Call the GetIdentityContext API to retrieve an identity token for the
// specified user.
//
// - Use the identity token with the [STS AssumeRole API]to obtain identity-enhanced IAM role
// session credentials.
//
// # Usage with STS AssumeRole
//
// The identity token returned by this API should be used with the STS AssumeRole
// API to obtain credentials for an identity-enhanced IAM role session. When
// calling AssumeRole, include the identity token in the ProvidedContexts
// parameter with ProviderArn set to arn:aws:iam::aws:contextProvider/QuickSight
// and ContextAssertion set to the identity token received from this API.
//
// The assumed role must allow the sts:SetContext action in addition to
// sts:AssumeRole in its trust relationship policy. The trust policy should include
// both actions for the principal that will be assuming the role.
//
// [DescribeDashboardSnapshotJobResult]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeDashboardSnapshotJobResult.html
// [Identity Center documentation on identity-enhanced IAM role sessions]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-identity-enhanced-iam-role-sessions.html
// [DescribeDashboardSnapshotJob]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeDashboardSnapshotJob.html
// [STS AssumeRole API]: https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html
// [Identity and access management in Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/userguide/identity.html
// [StartDashboardSnapshotJob]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StartDashboardSnapshotJob.html
func quicksight_GetIdentityContext(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GetIdentityContextInput{
		// AwsAccountId: *string, // Required
		// UserIdentifier: types.UserIdentifier, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightUserIdentifier) > 0 {
		if err := assignInputField(input, "UserIdentifier", _quicksightUserIdentifier); err != nil {
			log.Errorf("invalid --user-identifier: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightSessionExpiresAt) > 0 {
		if err := assignInputField(input, "SessionExpiresAt", _quicksightSessionExpiresAt); err != nil {
			log.Errorf("invalid --session-expires-at: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetIdentityContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a session URL and authorization code that you can use to embed the
// Amazon Amazon Quick Sight console in your web server code. Use
// GetSessionEmbedUrl where you want to provide an authoring portal that allows
// users to create data sources, datasets, analyses, and dashboards. The users who
// access an embedded Amazon Quick Sight console need belong to the author or admin
// security cohort. If you want to restrict permissions to some of these features,
// add a custom permissions profile to the user with the [UpdateUser]API operation. Use [RegisterUser] API
// operation to add a new user with a custom permission profile attached. For more
// information, see the following sections in the Amazon Quick Suite User Guide:
//
// [Embedding Analytics]
//
// [Customizing Access to the Amazon Quick Suite Console]
//
// [UpdateUser]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateUser.html
// [Customizing Access to the Amazon Quick Suite Console]: https://docs.aws.amazon.com/quicksight/latest/user/customizing-permissions-to-the-quicksight-console.html
// [RegisterUser]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_RegisterUser.html
// [Embedding Analytics]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html
func quicksight_GetSessionEmbedUrl(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.GetSessionEmbedUrlInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightEntryPoint) > 0 {
		input.EntryPoint = aws.String(_quicksightEntryPoint)
	}
	if len(_quicksightSessionLifetimeInMinutes) > 0 {
		if err := assignInputField(input, "SessionLifetimeInMinutes", _quicksightSessionLifetimeInMinutes); err != nil {
			log.Errorf("invalid --session-lifetime-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_quicksightUserArn) > 0 {
		input.UserArn = aws.String(_quicksightUserArn)
	}

	if resp, err := client.GetSessionEmbedUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all action connectors in the specified Amazon Web Services account.
// Returns summary information for each connector including its name, type,
// creation time, and status.
func quicksight_ListActionConnectors(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListActionConnectorsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListActionConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListActionConnectorsOutput
	p := quicksight.NewListActionConnectorsPaginator(client, input)
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

// Lists Amazon Quick Sight analyses that exist in the specified Amazon Web
// Services account.
func quicksight_ListAnalyses(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListAnalysesInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnalyses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListAnalysesOutput
	p := quicksight.NewListAnalysesPaginator(client, input)
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

// Lists all asset bundle export jobs that have been taken place in the last 14
// days. Jobs created more than 14 days ago are deleted forever and are not
// returned. If you are using the same job ID for multiple jobs,
// ListAssetBundleExportJobs only returns the most recent job that uses the
// repeated job ID.
func quicksight_ListAssetBundleExportJobs(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListAssetBundleExportJobsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetBundleExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListAssetBundleExportJobsOutput
	p := quicksight.NewListAssetBundleExportJobsPaginator(client, input)
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

// Lists all asset bundle import jobs that have taken place in the last 14 days.
// Jobs created more than 14 days ago are deleted forever and are not returned. If
// you are using the same job ID for multiple jobs, ListAssetBundleImportJobs only
// returns the most recent job that uses the repeated job ID.
func quicksight_ListAssetBundleImportJobs(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListAssetBundleImportJobsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetBundleImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListAssetBundleImportJobsOutput
	p := quicksight.NewListAssetBundleImportJobsPaginator(client, input)
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

// Lists all brands in an Quick Sight account.
func quicksight_ListBrands(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListBrandsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBrands(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListBrandsOutput
	p := quicksight.NewListBrandsPaginator(client, input)
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

// Returns a list of all the custom permissions profiles.
func quicksight_ListCustomPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListCustomPermissionsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListCustomPermissionsOutput
	p := quicksight.NewListCustomPermissionsPaginator(client, input)
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

// Lists all the versions of the dashboards in the Amazon Quick Sight subscription.
func quicksight_ListDashboardVersions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListDashboardVersionsInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDashboardVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListDashboardVersionsOutput
	p := quicksight.NewListDashboardVersionsPaginator(client, input)
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

// Lists dashboards in an Amazon Web Services account.
func quicksight_ListDashboards(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListDashboardsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDashboards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListDashboardsOutput
	p := quicksight.NewListDashboardsPaginator(client, input)
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

// Lists all of the datasets belonging to the current Amazon Web Services account
// in an Amazon Web Services Region.
//
// The permissions resource is arn:aws:quicksight:region:aws-account-id:dataset/* .
func quicksight_ListDataSets(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListDataSetsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListDataSetsOutput
	p := quicksight.NewListDataSetsPaginator(client, input)
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

// Lists data sources in current Amazon Web Services Region that belong to this
// Amazon Web Services account.
func quicksight_ListDataSources(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListDataSourcesInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListDataSourcesOutput
	p := quicksight.NewListDataSourcesPaginator(client, input)
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

// Lists flows in an Amazon Web Services account.
func quicksight_ListFlows(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListFlowsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListFlowsOutput
	p := quicksight.NewListFlowsPaginator(client, input)
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

// List all assets ( DASHBOARD , ANALYSIS , and DATASET ) in a folder.
func quicksight_ListFolderMembers(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListFolderMembersInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFolderMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListFolderMembersOutput
	p := quicksight.NewListFolderMembersPaginator(client, input)
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

// Lists all folders in an account.
func quicksight_ListFolders(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListFoldersInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFolders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListFoldersOutput
	p := quicksight.NewListFoldersPaginator(client, input)
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

// List all folders that a resource is a member of.
func quicksight_ListFoldersForResource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListFoldersForResourceInput{
		// AwsAccountId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightResourceArn) > 0 {
		input.ResourceArn = aws.String(_quicksightResourceArn)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFoldersForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListFoldersForResourceOutput
	p := quicksight.NewListFoldersForResourcePaginator(client, input)
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

// Lists member users in a group.
func quicksight_ListGroupMemberships(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListGroupMembershipsInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListGroupMembershipsOutput
	p := quicksight.NewListGroupMembershipsPaginator(client, input)
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

// Lists all user groups in Amazon Quick Sight.
func quicksight_ListGroups(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListGroupsInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListGroupsOutput
	p := quicksight.NewListGroupsPaginator(client, input)
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

// Lists the IAM policy assignments in the current Amazon Quick Sight account.
func quicksight_ListIAMPolicyAssignments(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListIAMPolicyAssignmentsInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightAssignmentStatus) > 0 {
		if err := assignInputField(input, "AssignmentStatus", _quicksightAssignmentStatus); err != nil {
			log.Errorf("invalid --assignment-status: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIAMPolicyAssignments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListIAMPolicyAssignmentsOutput
	p := quicksight.NewListIAMPolicyAssignmentsPaginator(client, input)
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

// Lists all of the IAM policy assignments, including the Amazon Resource Names
// (ARNs), for the IAM policies assigned to the specified user and group, or groups
// that the user belongs to.
func quicksight_ListIAMPolicyAssignmentsForUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListIAMPolicyAssignmentsForUserInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIAMPolicyAssignmentsForUser(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListIAMPolicyAssignmentsForUserOutput
	p := quicksight.NewListIAMPolicyAssignmentsForUserPaginator(client, input)
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

// Lists all services and authorized targets that the Quick Sight IAM Identity
// Center application can access.
//
// This operation is only supported for Quick Sight accounts that use IAM Identity
// Center.
func quicksight_ListIdentityPropagationConfigs(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListIdentityPropagationConfigsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if resp, err := client.ListIdentityPropagationConfigs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the history of SPICE ingestions for a dataset. Limited to 5 TPS per user
// and 25 TPS per account.
func quicksight_ListIngestions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListIngestionsInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIngestions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListIngestionsOutput
	p := quicksight.NewListIngestionsPaginator(client, input)
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

// Lists the namespaces for the specified Amazon Web Services account. This
// operation doesn't list deleted namespaces.
func quicksight_ListNamespaces(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListNamespacesInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListNamespacesOutput
	p := quicksight.NewListNamespacesPaginator(client, input)
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

// Lists the refresh schedules of a dataset. Each dataset can have up to 5
// schedules.
func quicksight_ListRefreshSchedules(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListRefreshSchedulesInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}

	if resp, err := client.ListRefreshSchedules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all groups that are associated with a role.
func quicksight_ListRoleMemberships(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListRoleMembershipsInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// Role: types.Role, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoleMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListRoleMembershipsOutput
	p := quicksight.NewListRoleMembershipsPaginator(client, input)
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

// Lists all self-upgrade requests for a Quick Suite account.
func quicksight_ListSelfUpgrades(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListSelfUpgradesInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if resp, err := client.ListSelfUpgrades(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags assigned to a resource.
func quicksight_ListTagsForResource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_quicksightResourceArn) > 0 {
		input.ResourceArn = aws.String(_quicksightResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the aliases of a template.
func quicksight_ListTemplateAliases(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTemplateAliasesInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListTemplateAliasesOutput
	p := quicksight.NewListTemplateAliasesPaginator(client, input)
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

// Lists all the versions of the templates in the current Amazon Quick Sight
// account.
func quicksight_ListTemplateVersions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTemplateVersionsInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListTemplateVersionsOutput
	p := quicksight.NewListTemplateVersionsPaginator(client, input)
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

// Lists all the templates in the current Amazon Quick Sight account.
func quicksight_ListTemplates(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTemplatesInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListTemplatesOutput
	p := quicksight.NewListTemplatesPaginator(client, input)
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

// Lists all the aliases of a theme.
func quicksight_ListThemeAliases(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListThemeAliasesInput{
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if resp, err := client.ListThemeAliases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the versions of the themes in the current Amazon Web Services account.
func quicksight_ListThemeVersions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListThemeVersionsInput{
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThemeVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListThemeVersionsOutput
	p := quicksight.NewListThemeVersionsPaginator(client, input)
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

// Lists all the themes in the current Amazon Web Services account.
func quicksight_ListThemes(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListThemesInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}
	if len(_quicksightType) > 0 {
		if err := assignInputField(input, "Type", _quicksightType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListThemes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListThemesOutput
	p := quicksight.NewListThemesPaginator(client, input)
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

// Lists all of the refresh schedules for a topic.
func quicksight_ListTopicRefreshSchedules(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTopicRefreshSchedulesInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.ListTopicRefreshSchedules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all reviewed answers for a Q Topic.
func quicksight_ListTopicReviewedAnswers(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTopicReviewedAnswersInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.ListTopicReviewedAnswers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the topics within an account.
func quicksight_ListTopics(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListTopicsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
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

	var results []*quicksight.ListTopicsOutput
	p := quicksight.NewListTopicsPaginator(client, input)
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

// Lists the Amazon Quick Sight groups that an Amazon Quick Sight user is a member
// of.
func quicksight_ListUserGroups(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListUserGroupsInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListUserGroupsOutput
	p := quicksight.NewListUserGroupsPaginator(client, input)
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

// Returns a list of all of the Amazon Quick Sight users belonging to this
// account.
func quicksight_ListUsers(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListUsersInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListUsersOutput
	p := quicksight.NewListUsersPaginator(client, input)
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

// Lists all of the VPC connections in the current set Amazon Web Services Region
// of an Amazon Web Services account.
func quicksight_ListVPCConnections(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.ListVPCConnectionsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVPCConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.ListVPCConnectionsOutput
	p := quicksight.NewListVPCConnectionsPaginator(client, input)
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

// Predicts existing visuals or generates new visuals to answer a given query.
// This API uses [trusted identity propagation] to ensure that an end user is authenticated and receives the
// embed URL that is specific to that user. The IAM Identity Center application
// that the user has logged into needs to have [trusted Identity Propagation enabled for Quick Suite]with the scope value set to
// quicksight:read . Before you use this action, make sure that you have configured
// the relevant Quick Suite resource and permissions.
//
// We recommend enabling the QSearchStatus API to unlock the full potential of
// PredictQnA . When QSearchStatus is enabled, it first checks the specified
// dashboard for any existing visuals that match the question. If no matching
// visuals are found, PredictQnA uses generative Q&A to provide an answer. To
// update the QSearchStatus , see [UpdateQuickSightQSearchConfiguration].
//
// [trusted Identity Propagation enabled for Quick Suite]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-using-customermanagedapps-specify-trusted-apps.html
// [trusted identity propagation]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation.html
// [UpdateQuickSightQSearchConfiguration]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateQuickSightQSearchConfiguration.html
func quicksight_PredictQAResults(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.PredictQAResultsInput{
		// AwsAccountId: *string, // Required
		// QueryText: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightQueryText) > 0 {
		input.QueryText = aws.String(_quicksightQueryText)
	}
	if len(_quicksightIncludeGeneratedAnswer) > 0 {
		if err := assignInputField(input, "IncludeGeneratedAnswer", _quicksightIncludeGeneratedAnswer); err != nil {
			log.Errorf("invalid --include-generated-answer: %s", err.Error())
			return
		}
	}
	if len(_quicksightIncludeQuickSightQIndex) > 0 {
		if err := assignInputField(input, "IncludeQuickSightQIndex", _quicksightIncludeQuickSightQIndex); err != nil {
			log.Errorf("invalid --include-quicksight-q-index: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxTopicsToConsider) > 0 {
		if err := assignInputField(input, "MaxTopicsToConsider", _quicksightMaxTopicsToConsider); err != nil {
			log.Errorf("invalid --max-topics-to-consider: %s", err.Error())
			return
		}
	}

	if resp, err := client.PredictQAResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the dataset refresh properties for the dataset.
func quicksight_PutDataSetRefreshProperties(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.PutDataSetRefreshPropertiesInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// DataSetRefreshProperties: *types.DataSetRefreshProperties, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightDataSetRefreshProperties) > 0 {
		if err := assignInputField(input, "DataSetRefreshProperties", _quicksightDataSetRefreshProperties); err != nil {
			log.Errorf("invalid --data-set-refresh-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDataSetRefreshProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Quick Sight user whose identity is associated with the
// Identity and Access Management (IAM) identity or role specified in the request.
// When you register a new user from the Quick Sight API, Quick Sight generates a
// registration URL. The user accesses this registration URL to create their
// account. Quick Sight doesn't send a registration email to users who are
// registered from the Quick Sight API. If you want new users to receive a
// registration email, then add those users in the Quick Sight console. For more
// information on registering a new user in the Quick Sight console, see [Inviting users to access Quick Sight].
//
// [Inviting users to access Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/managing-users.html#inviting-users
func quicksight_RegisterUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.RegisterUserInput{
		// AwsAccountId: *string, // Required
		// Email: *string, // Required
		// IdentityType: types.IdentityType, // Required
		// Namespace: *string, // Required
		// UserRole: types.UserRole, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightEmail) > 0 {
		input.Email = aws.String(_quicksightEmail)
	}
	if len(_quicksightIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _quicksightIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserRole) > 0 {
		if err := assignInputField(input, "UserRole", _quicksightUserRole); err != nil {
			log.Errorf("invalid --user-role: %s", err.Error())
			return
		}
	}
	if len(_quicksightCustomFederationProviderUrl) > 0 {
		input.CustomFederationProviderUrl = aws.String(_quicksightCustomFederationProviderUrl)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}
	if len(_quicksightExternalLoginFederationProviderType) > 0 {
		input.ExternalLoginFederationProviderType = aws.String(_quicksightExternalLoginFederationProviderType)
	}
	if len(_quicksightExternalLoginId) > 0 {
		input.ExternalLoginId = aws.String(_quicksightExternalLoginId)
	}
	if len(_quicksightIamArn) > 0 {
		input.IamArn = aws.String(_quicksightIamArn)
	}
	if len(_quicksightSessionName) > 0 {
		input.SessionName = aws.String(_quicksightSessionName)
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}

	if resp, err := client.RegisterUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores an analysis.
func quicksight_RestoreAnalysis(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.RestoreAnalysisInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightRestoreToFolders) > 0 {
		if err := assignInputField(input, "RestoreToFolders", _quicksightRestoreToFolders); err != nil {
			log.Errorf("invalid --restore-to-folders: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for action connectors in the specified Amazon Web Services account
// using filters. You can search by connector name, type, or user permissions.
func quicksight_SearchActionConnectors(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchActionConnectorsInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.ActionConnectorSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchActionConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchActionConnectorsOutput
	p := quicksight.NewSearchActionConnectorsPaginator(client, input)
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

// Searches for analyses that belong to the user specified in the filter.
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func quicksight_SearchAnalyses(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchAnalysesInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.AnalysisSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchAnalyses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchAnalysesOutput
	p := quicksight.NewSearchAnalysesPaginator(client, input)
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

// Searches for dashboards that belong to a user.
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func quicksight_SearchDashboards(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchDashboardsInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.DashboardSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchDashboards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchDashboardsOutput
	p := quicksight.NewSearchDashboardsPaginator(client, input)
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

// Use the SearchDataSets operation to search for datasets that belong to an
// account.
func quicksight_SearchDataSets(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchDataSetsInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.DataSetSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchDataSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchDataSetsOutput
	p := quicksight.NewSearchDataSetsPaginator(client, input)
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

// Use the SearchDataSources operation to search for data sources that belong to
// an account.
func quicksight_SearchDataSources(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchDataSourcesInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.DataSourceSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchDataSourcesOutput
	p := quicksight.NewSearchDataSourcesPaginator(client, input)
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

// Search for the flows in an Amazon Web Services account.
func quicksight_SearchFlows(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchFlowsInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.SearchFlowsFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchFlowsOutput
	p := quicksight.NewSearchFlowsPaginator(client, input)
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

// Searches the subfolders in a folder.
func quicksight_SearchFolders(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchFoldersInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.FolderSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchFolders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchFoldersOutput
	p := quicksight.NewSearchFoldersPaginator(client, input)
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

// Use the SearchGroups operation to search groups in a specified Quick Sight
// namespace using the supplied filters.
func quicksight_SearchGroups(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchGroupsInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.GroupSearchFilter, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchGroupsOutput
	p := quicksight.NewSearchGroupsPaginator(client, input)
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

// Searches for any Q topic that exists in an Quick Suite account.
func quicksight_SearchTopics(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.SearchTopicsInput{
		// AwsAccountId: *string, // Required
		// Filters: []types.TopicSearchFilter, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFilters) > 0 {
		if err := assignInputField(input, "Filters", _quicksightFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_quicksightMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _quicksightMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_quicksightNextToken) > 0 {
		input.NextToken = aws.String(_quicksightNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchTopics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*quicksight.SearchTopicsOutput
	p := quicksight.NewSearchTopicsPaginator(client, input)
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

// Starts an Asset Bundle export job.
// An Asset Bundle export job exports specified Amazon Quick Sight assets. You can
// also choose to export any asset dependencies in the same job. Export jobs run
// asynchronously and can be polled with a DescribeAssetBundleExportJob API call.
// When a job is successfully completed, a download URL that contains the exported
// assets is returned. The URL is valid for 5 minutes and can be refreshed with a
// DescribeAssetBundleExportJob API call. Each Amazon Quick Sight account can run
// up to 5 export jobs concurrently.
//
// The API caller must have the necessary permissions in their IAM role to access
// each resource before the resources can be exported.
func quicksight_StartAssetBundleExportJob(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.StartAssetBundleExportJobInput{
		// AssetBundleExportJobId: *string, // Required
		// AwsAccountId: *string, // Required
		// ExportFormat: types.AssetBundleExportFormat, // Required
		// ResourceArns: []string, // Required
	}

	if len(_quicksightAssetBundleExportJobId) > 0 {
		input.AssetBundleExportJobId = aws.String(_quicksightAssetBundleExportJobId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightExportFormat) > 0 {
		if err := assignInputField(input, "ExportFormat", _quicksightExportFormat); err != nil {
			log.Errorf("invalid --export-format: %s", err.Error())
			return
		}
	}
	if len(_quicksightResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _quicksightResourceArns...)
	}
	if len(_quicksightCloudFormationOverridePropertyConfiguration) > 0 {
		if err := assignInputField(input, "CloudFormationOverridePropertyConfiguration", _quicksightCloudFormationOverridePropertyConfiguration); err != nil {
			log.Errorf("invalid --cloud-formation-override-property-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightIncludeAllDependencies) > 0 {
		if err := assignInputField(input, "IncludeAllDependencies", _quicksightIncludeAllDependencies); err != nil {
			log.Errorf("invalid --include-all-dependencies: %s", err.Error())
			return
		}
	}
	if len(_quicksightIncludeFolderMembers) > 0 {
		if err := assignInputField(input, "IncludeFolderMembers", _quicksightIncludeFolderMembers); err != nil {
			log.Errorf("invalid --include-folder-members: %s", err.Error())
			return
		}
	}
	if len(_quicksightIncludeFolderMemberships) > 0 {
		if err := assignInputField(input, "IncludeFolderMemberships", _quicksightIncludeFolderMemberships); err != nil {
			log.Errorf("invalid --include-folder-memberships: %s", err.Error())
			return
		}
	}
	if len(_quicksightIncludePermissions) > 0 {
		if err := assignInputField(input, "IncludePermissions", _quicksightIncludePermissions); err != nil {
			log.Errorf("invalid --include-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightIncludeTags) > 0 {
		if err := assignInputField(input, "IncludeTags", _quicksightIncludeTags); err != nil {
			log.Errorf("invalid --include-tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAssetBundleExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an Asset Bundle import job.
// An Asset Bundle import job imports specified Amazon Quick Sight assets into an
// Amazon Quick Sight account. You can also choose to import a naming prefix and
// specified configuration overrides. The assets that are contained in the bundle
// file that you provide are used to create or update a new or existing asset in
// your Amazon Quick Sight account. Each Amazon Quick Sight account can run up to 5
// import jobs concurrently.
//
// The API caller must have the necessary "create" , "describe" , and "update"
// permissions in their IAM role to access each resource type that is contained in
// the bundle file before the resources can be imported.
func quicksight_StartAssetBundleImportJob(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.StartAssetBundleImportJobInput{
		// AssetBundleImportJobId: *string, // Required
		// AssetBundleImportSource: *types.AssetBundleImportSource, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAssetBundleImportJobId) > 0 {
		input.AssetBundleImportJobId = aws.String(_quicksightAssetBundleImportJobId)
	}
	if len(_quicksightAssetBundleImportSource) > 0 {
		if err := assignInputField(input, "AssetBundleImportSource", _quicksightAssetBundleImportSource); err != nil {
			log.Errorf("invalid --asset-bundle-import-source: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFailureAction) > 0 {
		if err := assignInputField(input, "FailureAction", _quicksightFailureAction); err != nil {
			log.Errorf("invalid --failure-action: %s", err.Error())
			return
		}
	}
	if len(_quicksightOverrideParameters) > 0 {
		if err := assignInputField(input, "OverrideParameters", _quicksightOverrideParameters); err != nil {
			log.Errorf("invalid --override-parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightOverridePermissions) > 0 {
		if err := assignInputField(input, "OverridePermissions", _quicksightOverridePermissions); err != nil {
			log.Errorf("invalid --override-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightOverrideTags) > 0 {
		if err := assignInputField(input, "OverrideTags", _quicksightOverrideTags); err != nil {
			log.Errorf("invalid --override-tags: %s", err.Error())
			return
		}
	}
	if len(_quicksightOverrideValidationStrategy) > 0 {
		if err := assignInputField(input, "OverrideValidationStrategy", _quicksightOverrideValidationStrategy); err != nil {
			log.Errorf("invalid --override-validation-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAssetBundleImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous job that generates a snapshot of a dashboard's output.
// You can request one or several of the following format configurations in each
// API call.
//
// - 1 PDF
//
// - 1 Excel workbook that includes up to 5 table or pivot table visuals
//
// - 5 CSVs from table or pivot table visuals
//
// Exporting CSV, Excel, or Pixel Perfect PDF reports requires Pixel Perfect
// Report Add-on.
//
// The status of a submitted job can be polled with the
// DescribeDashboardSnapshotJob API. When you call the DescribeDashboardSnapshotJob
// API, check the JobStatus field in the response. Once the job reaches a COMPLETED
// or FAILED status, use the DescribeDashboardSnapshotJobResult API to obtain the
// URLs for the generated files. If the job fails, the
// DescribeDashboardSnapshotJobResult API returns detailed information about the
// error that occurred.
//
// # StartDashboardSnapshotJob API throttling
//
// Quick Sight utilizes API throttling to create a more consistent user experience
// within a time span for customers when they call the StartDashboardSnapshotJob .
// By default, 12 jobs can run simlutaneously in one Amazon Web Services account
// and users can submit up 10 API requests per second before an account is
// throttled. If an overwhelming number of API requests are made by the same user
// in a short period of time, Quick Sight throttles the API calls to maintin an
// optimal experience and reliability for all Quick Sight users.
//
// # Common throttling scenarios
//
// The following list provides information about the most commin throttling
// scenarios that can occur.
//
// - A large number of SnapshotExport API jobs are running simultaneously on an
// Amazon Web Services account. When a new StartDashboardSnapshotJob is created
// and there are already 12 jobs with the RUNNING status, the new job request
// fails and returns a LimitExceededException error. Wait for a current job to
// comlpete before you resubmit the new job.
//
// - A large number of API requests are submitted on an Amazon Web Services
// account. When a user makes more than 10 API calls to the Quick Sight API in one
// second, a ThrottlingException is returned.
//
// If your use case requires a higher throttling limit, contact your account admin
// or [Amazon Web ServicesSupport]to explore options to tailor a more optimal expereince for your account.
//
// # Best practices to handle throttling
//
// If your use case projects high levels of API traffic, try to reduce the degree
// of frequency and parallelism of API calls as much as you can to avoid
// throttling. You can also perform a timing test to calculate an estimate for the
// total processing time of your projected load that stays within the throttling
// limits of the Quick Sight APIs. For example, if your projected traffic is 100
// snapshot jobs before 12:00 PM per day, start 12 jobs in parallel and measure the
// amount of time it takes to proccess all 12 jobs. Once you obtain the result,
// multiply the duration by 9, for example (12 minutes * 9 = 108 minutes) . Use the
// new result to determine the latest time at which the jobs need to be started to
// meet your target deadline.
//
// The time that it takes to process a job can be impacted by the following
// factors:
//
// - The dataset type (Direct Query or SPICE).
//
// - The size of the dataset.
//
// - The complexity of the calculated fields that are used in the dashboard.
//
// - The number of visuals that are on a sheet.
//
// - The types of visuals that are on the sheet.
//
// - The number of formats and snapshots that are requested in the job
// configuration.
//
// - The size of the generated snapshots.
//
// # Registered user support
//
// You can generate snapshots for registered Quick Sight users by using the
// Snapshot Job APIs with [identity-enhanced IAM role session credentials]. This approach allows you to create snapshots on behalf
// of specific Quick Sight users while respecting their row-level security (RLS),
// column-level security (CLS), dynamic default parameters and dashboard
// parameter/filter settings.
//
// To generate snapshots for registered Quick Sight users, you need to:
//
// - Obtain identity-enhanced IAM role session credentials from Amazon Web
// Services Security Token Service (STS).
//
// - Use these credentials to call the Snapshot Job APIs.
//
// Identity-enhanced credentials are credentials that contain information about
// the end user (e.g., registered Quick Sight user).
//
// If your Quick Sight users are backed by [Amazon Web Services Identity Center], then you need to set up a [trusted token issuer]. Then,
// getting identity-enhanced IAM credentials for a Quick Sight user will look like
// the following:
//
// - Authenticate user with your OIDC compliant Identity Provider. You should
// get auth tokens back.
//
// - Use the OIDC API, [CreateTokenWithIAM], to exchange auth tokens to IAM tokens. One of the
// resulted tokens will be identity token.
//
// - Call STS AssumeRole API as you normally would, but provide an extra
// ProvidedContexts parameter in the API request. The list of contexts must have
// a single trusted context assertion. The ProviderArn should be
// arn:aws:iam::aws:contextProvider/IdentityCenter while ContextAssertion will be
// the identity token you received in response from CreateTokenWithIAM
//
// For more details, see [IdC documentation on Identity-enhanced IAM role sessions].
//
// To obtain Identity-enhanced credentials for Quick Sight native users, IAM
// federated users, or Active Directory users, follow the steps below:
//
// - Call Quick Sight [GetIdentityContext API]to get identity token.
//
// - Call STS AssumeRole API as you normally would, but provide extra
// ProvidedContexts parameter in the API request. The list of contexts must have
// a single trusted context assertion. The ProviderArn should be
// arn:aws:iam::aws:contextProvider/QuickSight while ContextAssertion will be the
// identity token you received in response from GetIdentityContext
//
// After obtaining the identity-enhanced IAM role session credentials, you can use
// them to start a job, describe the job and describe job result. You can use the
// same credentials as long as they haven't expired. All API requests made with
// these credentials are considered to be made by the impersonated Quick Sight
// user.
//
// When using identity-enhanced session credentials, set the UserConfiguration
// request attribute to null. Otherwise, the request will be invalid.
//
// # Possible error scenarios
//
// The request fails with an Access Denied error in the following scenarios:
//
// - The credentials have expired.
//
// - The impersonated Quick Sight user doesn't have access to the specified
// dashboard.
//
// - The impersonated Quick Sight user is restricted from exporting data in the
// selected formats. For more information about export restrictions, see [Customizing access to Amazon Quick Sight capabilities].
//
// [GetIdentityContext API]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GetIdentityContext.html
// [identity-enhanced IAM role session credentials]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-identity-enhanced-iam-role-sessions.html
// [CreateTokenWithIAM]: https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateTokenWithIAM.html
// [IdC documentation on Identity-enhanced IAM role sessions]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-identity-enhanced-iam-role-sessions.html
// [trusted token issuer]: https://docs.aws.amazon.com/singlesignon/latest/userguide/setuptrustedtokenissuer.html
// [Customizing access to Amazon Quick Sight capabilities]: https://docs.aws.amazon.com/quicksuite/latest/userguide/create-custom-permisions-profile.html
// [Amazon Web ServicesSupport]: http://aws.amazon.com/contact-us/
// [Amazon Web Services Identity Center]: https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html
func quicksight_StartDashboardSnapshotJob(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.StartDashboardSnapshotJobInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// SnapshotConfiguration: *types.SnapshotConfiguration, // Required
		// SnapshotJobId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightSnapshotConfiguration) > 0 {
		if err := assignInputField(input, "SnapshotConfiguration", _quicksightSnapshotConfiguration); err != nil {
			log.Errorf("invalid --snapshot-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightSnapshotJobId) > 0 {
		input.SnapshotJobId = aws.String(_quicksightSnapshotJobId)
	}
	if len(_quicksightUserConfiguration) > 0 {
		if err := assignInputField(input, "UserConfiguration", _quicksightUserConfiguration); err != nil {
			log.Errorf("invalid --user-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDashboardSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous job that runs an existing dashboard schedule and sends
// the dashboard snapshot through email.
//
// Only one job can run simultaneously in a given schedule. Repeated requests are
// skipped with a 202 HTTP status code.
//
// For more information, see [Scheduling and sending Amazon Quick Sight reports by email] and [Configuring email report settings for a Amazon Quick Sight dashboard] in the Amazon Quick Sight User Guide.
//
// [Configuring email report settings for a Amazon Quick Sight dashboard]: https://docs.aws.amazon.com/quicksight/latest/user/email-reports-from-dashboard.html
// [Scheduling and sending Amazon Quick Sight reports by email]: https://docs.aws.amazon.com/quicksight/latest/user/sending-reports.html
func quicksight_StartDashboardSnapshotJobSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.StartDashboardSnapshotJobScheduleInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// ScheduleId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightScheduleId) > 0 {
		input.ScheduleId = aws.String(_quicksightScheduleId)
	}

	if resp, err := client.StartDashboardSnapshotJobSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified Amazon Quick Sight
// resource.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions, by granting a user permission to access or change
// only resources with certain tag values. You can use the TagResource operation
// with a resource that already has tags. If you specify a new tag key for the
// resource, this tag is appended to the list of tags associated with the resource.
// If you specify a tag key that is already associated with the resource, the new
// tag value that you specify replaces the previous value for that tag.
//
// You can associate as many as 50 tags with a resource. Amazon Quick Sight
// supports tagging on data set, data source, dashboard, template, topic, and user.
//
// Tagging for Amazon Quick Sight works in a similar way to tagging for other
// Amazon Web Services services, except for the following:
//
// - Tags are used to track costs for users in Amazon Quick Sight. You can't tag
// other resources that Amazon Quick Sight costs are based on, such as storage
// capacoty (SPICE), session usage, alert consumption, or reporting units.
//
// - Amazon Quick Sight doesn't currently support the tag editor for Resource
// Groups.
func quicksight_TagResource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_quicksightResourceArn) > 0 {
		input.ResourceArn = aws.String(_quicksightResourceArn)
	}
	if len(_quicksightTags) > 0 {
		if err := assignInputField(input, "Tags", _quicksightTags); err != nil {
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

// Removes a tag or tags from a resource.
func quicksight_UntagResource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_quicksightResourceArn) > 0 {
		input.ResourceArn = aws.String(_quicksightResourceArn)
	}
	if len(_quicksightTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _quicksightTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a custom permissions profile to an account.
func quicksight_UpdateAccountCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateAccountCustomPermissionInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}

	if resp, err := client.UpdateAccountCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates Amazon Quick Sight customizations. Currently, the only customization
// that you can use is a theme.
//
// You can use customizations for your Amazon Web Services account or, if you
// specify a namespace, for a Quick Sight namespace instead. Customizations that
// apply to a namespace override customizations that apply to an Amazon Web
// Services account. To find out which customizations apply, use the
// DescribeAccountCustomization API operation.
func quicksight_UpdateAccountCustomization(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateAccountCustomizationInput{
		// AccountCustomization: *types.AccountCustomization, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAccountCustomization) > 0 {
		if err := assignInputField(input, "AccountCustomization", _quicksightAccountCustomization); err != nil {
			log.Errorf("invalid --account-customization: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.UpdateAccountCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Amazon Quick Sight settings in your Amazon Web Services account.
func quicksight_UpdateAccountSettings(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateAccountSettingsInput{
		// AwsAccountId: *string, // Required
		// DefaultNamespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDefaultNamespace) > 0 {
		input.DefaultNamespace = aws.String(_quicksightDefaultNamespace)
	}
	if len(_quicksightNotificationEmail) > 0 {
		input.NotificationEmail = aws.String(_quicksightNotificationEmail)
	}
	if len(_quicksightTerminationProtectionEnabled) > 0 {
		if err := assignInputField(input, "TerminationProtectionEnabled", _quicksightTerminationProtectionEnabled); err != nil {
			log.Errorf("invalid --termination-protection-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing action connector with new configuration details,
// authentication settings, or enabled actions. You can modify the connector's
// name, description, authentication configuration, and which actions are enabled.
// For more information, [https://docs.aws.amazon.com/quicksuite/latest/userguide/quick-action-auth.html].
//
// [https://docs.aws.amazon.com/quicksuite/latest/userguide/quick-action-auth.html]: https://docs.aws.amazon.com/quicksuite/latest/userguide/quick-action-auth.html
func quicksight_UpdateActionConnector(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateActionConnectorInput{
		// ActionConnectorId: *string, // Required
		// AuthenticationConfig: *types.AuthConfig, // Required
		// AwsAccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightActionConnectorId) > 0 {
		input.ActionConnectorId = aws.String(_quicksightActionConnectorId)
	}
	if len(_quicksightAuthenticationConfig) > 0 {
		if err := assignInputField(input, "AuthenticationConfig", _quicksightAuthenticationConfig); err != nil {
			log.Errorf("invalid --authentication-config: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightDescription) > 0 {
		input.Description = aws.String(_quicksightDescription)
	}
	if len(_quicksightVpcConnectionArn) > 0 {
		input.VpcConnectionArn = aws.String(_quicksightVpcConnectionArn)
	}

	if resp, err := client.UpdateActionConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the permissions for an action connector by granting or revoking access
// for specific users and groups. You can control who can view, use, or manage the
// action connector.
func quicksight_UpdateActionConnectorPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateActionConnectorPermissionsInput{
		// ActionConnectorId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightActionConnectorId) > 0 {
		input.ActionConnectorId = aws.String(_quicksightActionConnectorId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateActionConnectorPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an analysis in Amazon Quick Sight
func quicksight_UpdateAnalysis(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateAnalysisInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightDefinition) > 0 {
		if err := assignInputField(input, "Definition", _quicksightDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightParameters) > 0 {
		if err := assignInputField(input, "Parameters", _quicksightParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightSourceEntity) > 0 {
		if err := assignInputField(input, "SourceEntity", _quicksightSourceEntity); err != nil {
			log.Errorf("invalid --source-entity: %s", err.Error())
			return
		}
	}
	if len(_quicksightThemeArn) > 0 {
		input.ThemeArn = aws.String(_quicksightThemeArn)
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the read and write permissions for an analysis.
func quicksight_UpdateAnalysisPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateAnalysisPermissionsInput{
		// AnalysisId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAnalysisId) > 0 {
		input.AnalysisId = aws.String(_quicksightAnalysisId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAnalysisPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Quick Suite application with a token exchange grant. This operation
// only supports Quick Suite applications that are registered with IAM Identity
// Center.
func quicksight_UpdateApplicationWithTokenExchangeGrant(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateApplicationWithTokenExchangeGrantInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.UpdateApplicationWithTokenExchangeGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a brand.
func quicksight_UpdateBrand(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateBrandInput{
		// AwsAccountId: *string, // Required
		// BrandId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandId) > 0 {
		input.BrandId = aws.String(_quicksightBrandId)
	}
	if len(_quicksightBrandDefinition) > 0 {
		if err := assignInputField(input, "BrandDefinition", _quicksightBrandDefinition); err != nil {
			log.Errorf("invalid --brand-definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBrand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a brand assignment.
func quicksight_UpdateBrandAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateBrandAssignmentInput{
		// AwsAccountId: *string, // Required
		// BrandArn: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandArn) > 0 {
		input.BrandArn = aws.String(_quicksightBrandArn)
	}

	if resp, err := client.UpdateBrandAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the published version of a brand.
func quicksight_UpdateBrandPublishedVersion(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateBrandPublishedVersionInput{
		// AwsAccountId: *string, // Required
		// BrandId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBrandId) > 0 {
		input.BrandId = aws.String(_quicksightBrandId)
	}
	if len(_quicksightVersionId) > 0 {
		input.VersionId = aws.String(_quicksightVersionId)
	}

	if resp, err := client.UpdateBrandPublishedVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom permissions profile.
func quicksight_UpdateCustomPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateCustomPermissionsInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}
	if len(_quicksightCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _quicksightCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCustomPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a dashboard in an Amazon Web Services account.
// Updating a Dashboard creates a new dashboard version but does not immediately
// publish the new version. You can update the published version of a dashboard by
// using the [UpdateDashboardPublishedVersion]API operation.
//
// [UpdateDashboardPublishedVersion]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateDashboardPublishedVersion.html
func quicksight_UpdateDashboard(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDashboardInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightDashboardPublishOptions) > 0 {
		if err := assignInputField(input, "DashboardPublishOptions", _quicksightDashboardPublishOptions); err != nil {
			log.Errorf("invalid --dashboard-publish-options: %s", err.Error())
			return
		}
	}
	if len(_quicksightDefinition) > 0 {
		if err := assignInputField(input, "Definition", _quicksightDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightParameters) > 0 {
		if err := assignInputField(input, "Parameters", _quicksightParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightSourceEntity) > 0 {
		if err := assignInputField(input, "SourceEntity", _quicksightSourceEntity); err != nil {
			log.Errorf("invalid --source-entity: %s", err.Error())
			return
		}
	}
	if len(_quicksightThemeArn) > 0 {
		input.ThemeArn = aws.String(_quicksightThemeArn)
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}
	if len(_quicksightVersionDescription) > 0 {
		input.VersionDescription = aws.String(_quicksightVersionDescription)
	}

	if resp, err := client.UpdateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the linked analyses on a dashboard.
func quicksight_UpdateDashboardLinks(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDashboardLinksInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// LinkEntities: []string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightLinkEntities) > 0 {
		input.LinkEntities = append([]string(nil), _quicksightLinkEntities...)
	}

	if resp, err := client.UpdateDashboardLinks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates read and write permissions on a dashboard.
func quicksight_UpdateDashboardPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDashboardPermissionsInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightGrantLinkPermissions) > 0 {
		if err := assignInputField(input, "GrantLinkPermissions", _quicksightGrantLinkPermissions); err != nil {
			log.Errorf("invalid --grant-link-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokeLinkPermissions) > 0 {
		if err := assignInputField(input, "RevokeLinkPermissions", _quicksightRevokeLinkPermissions); err != nil {
			log.Errorf("invalid --revoke-link-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDashboardPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the published version of a dashboard.
func quicksight_UpdateDashboardPublishedVersion(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDashboardPublishedVersionInput{
		// AwsAccountId: *string, // Required
		// DashboardId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardId) > 0 {
		input.DashboardId = aws.String(_quicksightDashboardId)
	}
	if len(_quicksightVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _quicksightVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDashboardPublishedVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Dashboard QA configuration.
func quicksight_UpdateDashboardsQAConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDashboardsQAConfigurationInput{
		// AwsAccountId: *string, // Required
		// DashboardsQAStatus: types.DashboardsQAStatus, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDashboardsQAStatus) > 0 {
		if err := assignInputField(input, "DashboardsQAStatus", _quicksightDashboardsQAStatus); err != nil {
			log.Errorf("invalid --dashboards-qa-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDashboardsQAConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a dataset. This operation doesn't support datasets that include
// uploaded files as a source. Partial updates are not supported by this operation.
func quicksight_UpdateDataSet(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDataSetInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// ImportMode: types.DataSetImportMode, // Required
		// Name: *string, // Required
		// PhysicalTableMap: map[string]types.PhysicalTable, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightImportMode) > 0 {
		if err := assignInputField(input, "ImportMode", _quicksightImportMode); err != nil {
			log.Errorf("invalid --import-mode: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightPhysicalTableMap) > 0 {
		if err := assignInputField(input, "PhysicalTableMap", _quicksightPhysicalTableMap); err != nil {
			log.Errorf("invalid --physical-table-map: %s", err.Error())
			return
		}
	}
	if len(_quicksightColumnGroups) > 0 {
		if err := assignInputField(input, "ColumnGroups", _quicksightColumnGroups); err != nil {
			log.Errorf("invalid --column-groups: %s", err.Error())
			return
		}
	}
	if len(_quicksightColumnLevelPermissionRules) > 0 {
		if err := assignInputField(input, "ColumnLevelPermissionRules", _quicksightColumnLevelPermissionRules); err != nil {
			log.Errorf("invalid --column-level-permission-rules: %s", err.Error())
			return
		}
	}
	if len(_quicksightDataPrepConfiguration) > 0 {
		if err := assignInputField(input, "DataPrepConfiguration", _quicksightDataPrepConfiguration); err != nil {
			log.Errorf("invalid --data-prep-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightDataSetUsageConfiguration) > 0 {
		if err := assignInputField(input, "DataSetUsageConfiguration", _quicksightDataSetUsageConfiguration); err != nil {
			log.Errorf("invalid --data-set-usage-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightDatasetParameters) > 0 {
		if err := assignInputField(input, "DatasetParameters", _quicksightDatasetParameters); err != nil {
			log.Errorf("invalid --dataset-parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightFieldFolders) > 0 {
		if err := assignInputField(input, "FieldFolders", _quicksightFieldFolders); err != nil {
			log.Errorf("invalid --field-folders: %s", err.Error())
			return
		}
	}
	if len(_quicksightLogicalTableMap) > 0 {
		if err := assignInputField(input, "LogicalTableMap", _quicksightLogicalTableMap); err != nil {
			log.Errorf("invalid --logical-table-map: %s", err.Error())
			return
		}
	}
	if len(_quicksightPerformanceConfiguration) > 0 {
		if err := assignInputField(input, "PerformanceConfiguration", _quicksightPerformanceConfiguration); err != nil {
			log.Errorf("invalid --performance-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightRowLevelPermissionDataSet) > 0 {
		if err := assignInputField(input, "RowLevelPermissionDataSet", _quicksightRowLevelPermissionDataSet); err != nil {
			log.Errorf("invalid --row-level-permission-data-set: %s", err.Error())
			return
		}
	}
	if len(_quicksightRowLevelPermissionTagConfiguration) > 0 {
		if err := assignInputField(input, "RowLevelPermissionTagConfiguration", _quicksightRowLevelPermissionTagConfiguration); err != nil {
			log.Errorf("invalid --row-level-permission-tag-configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightSemanticModelConfiguration) > 0 {
		if err := assignInputField(input, "SemanticModelConfiguration", _quicksightSemanticModelConfiguration); err != nil {
			log.Errorf("invalid --semantic-model-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the permissions on a dataset.
// The permissions resource is
// arn:aws:quicksight:region:aws-account-id:dataset/data-set-id .
func quicksight_UpdateDataSetPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDataSetPermissionsInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSetPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a data source.
func quicksight_UpdateDataSource(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDataSourceInput{
		// AwsAccountId: *string, // Required
		// DataSourceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSourceId) > 0 {
		input.DataSourceId = aws.String(_quicksightDataSourceId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightCredentials) > 0 {
		if err := assignInputField(input, "Credentials", _quicksightCredentials); err != nil {
			log.Errorf("invalid --credentials: %s", err.Error())
			return
		}
	}
	if len(_quicksightDataSourceParameters) > 0 {
		if err := assignInputField(input, "DataSourceParameters", _quicksightDataSourceParameters); err != nil {
			log.Errorf("invalid --data-source-parameters: %s", err.Error())
			return
		}
	}
	if len(_quicksightSslProperties) > 0 {
		if err := assignInputField(input, "SslProperties", _quicksightSslProperties); err != nil {
			log.Errorf("invalid --ssl-properties: %s", err.Error())
			return
		}
	}
	if len(_quicksightVpcConnectionProperties) > 0 {
		if err := assignInputField(input, "VpcConnectionProperties", _quicksightVpcConnectionProperties); err != nil {
			log.Errorf("invalid --vpc-connection-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the permissions to a data source.
func quicksight_UpdateDataSourcePermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDataSourcePermissionsInput{
		// AwsAccountId: *string, // Required
		// DataSourceId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSourceId) > 0 {
		input.DataSourceId = aws.String(_quicksightDataSourceId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSourcePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Amazon Q Business application that is linked to a Quick Sight account.
func quicksight_UpdateDefaultQBusinessApplication(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateDefaultQBusinessApplicationInput{
		// ApplicationId: *string, // Required
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightApplicationId) > 0 {
		input.ApplicationId = aws.String(_quicksightApplicationId)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}

	if resp, err := client.UpdateDefaultQBusinessApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates permissions against principals on a flow.
func quicksight_UpdateFlowPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateFlowPermissionsInput{
		// AwsAccountId: *string, // Required
		// FlowId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFlowId) > 0 {
		input.FlowId = aws.String(_quicksightFlowId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFlowPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of a folder.
func quicksight_UpdateFolder(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateFolderInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
		// Name: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}

	if resp, err := client.UpdateFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates permissions of a folder.
func quicksight_UpdateFolderPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateFolderPermissionsInput{
		// AwsAccountId: *string, // Required
		// FolderId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightFolderId) > 0 {
		input.FolderId = aws.String(_quicksightFolderId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFolderPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes a group description.
func quicksight_UpdateGroup(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateGroupInput{
		// AwsAccountId: *string, // Required
		// GroupName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightGroupName) > 0 {
		input.GroupName = aws.String(_quicksightGroupName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightDescription) > 0 {
		input.Description = aws.String(_quicksightDescription)
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing IAM policy assignment. This operation updates only the
// optional parameter or parameters that are specified in the request. This
// overwrites all of the users included in Identities .
func quicksight_UpdateIAMPolicyAssignment(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateIAMPolicyAssignmentInput{
		// AssignmentName: *string, // Required
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_quicksightAssignmentName) > 0 {
		input.AssignmentName = aws.String(_quicksightAssignmentName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightAssignmentStatus) > 0 {
		if err := assignInputField(input, "AssignmentStatus", _quicksightAssignmentStatus); err != nil {
			log.Errorf("invalid --assignment-status: %s", err.Error())
			return
		}
	}
	if len(_quicksightIdentities) > 0 {
		if err := assignInputField(input, "Identities", _quicksightIdentities); err != nil {
			log.Errorf("invalid --identities: %s", err.Error())
			return
		}
	}
	if len(_quicksightPolicyArn) > 0 {
		input.PolicyArn = aws.String(_quicksightPolicyArn)
	}

	if resp, err := client.UpdateIAMPolicyAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates services and authorized targets to configure what the Quick
// Sight IAM Identity Center application can access.
//
// This operation is only supported for Quick Sight accounts using IAM Identity
// Center
func quicksight_UpdateIdentityPropagationConfig(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateIdentityPropagationConfigInput{
		// AwsAccountId: *string, // Required
		// Service: types.ServiceType, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightService) > 0 {
		if err := assignInputField(input, "Service", _quicksightService); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}
	if len(_quicksightAuthorizedTargets) > 0 {
		input.AuthorizedTargets = append([]string(nil), _quicksightAuthorizedTargets...)
	}

	if resp, err := client.UpdateIdentityPropagationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the content and status of IP rules. Traffic from a source is allowed
// when the source satisfies either the IpRestrictionRule , VpcIdRestrictionRule ,
// or VpcEndpointIdRestrictionRule . To use this operation, you must provide the
// entire map of rules. You can use the DescribeIpRestriction operation to get the
// current rule map.
func quicksight_UpdateIpRestriction(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateIpRestrictionInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _quicksightEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_quicksightIpRestrictionRuleMap) > 0 {
		if err := assignInputField(input, "IpRestrictionRuleMap", _quicksightIpRestrictionRuleMap); err != nil {
			log.Errorf("invalid --ip-restriction-rule-map: %s", err.Error())
			return
		}
	}
	if len(_quicksightVpcEndpointIdRestrictionRuleMap) > 0 {
		if err := assignInputField(input, "VpcEndpointIdRestrictionRuleMap", _quicksightVpcEndpointIdRestrictionRuleMap); err != nil {
			log.Errorf("invalid --vpc-endpoint-id-restriction-rule-map: %s", err.Error())
			return
		}
	}
	if len(_quicksightVpcIdRestrictionRuleMap) > 0 {
		if err := assignInputField(input, "VpcIdRestrictionRuleMap", _quicksightVpcIdRestrictionRuleMap); err != nil {
			log.Errorf("invalid --vpc-id-restriction-rule-map: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIpRestriction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a customer managed key in a Quick Sight account.
func quicksight_UpdateKeyRegistration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateKeyRegistrationInput{
		// AwsAccountId: *string, // Required
		// KeyRegistration: []types.RegisteredCustomerManagedKey, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightKeyRegistration) > 0 {
		if err := assignInputField(input, "KeyRegistration", _quicksightKeyRegistration); err != nil {
			log.Errorf("invalid --key-registration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKeyRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API controls public sharing settings for your entire Quick Sight account,
// affecting data security and access. When you enable public sharing:
//
// - Dashboards can be shared publicly
//
// - This setting affects your entire Amazon Web Services account and all Quick
// Sight users
//
// Before proceeding: Ensure you understand the security implications and have
// proper IAM permissions configured.
//
// Use the UpdatePublicSharingSettings operation to turn on or turn off the public
// sharing settings of an Amazon Quick Sight dashboard.
//
// To use this operation, turn on session capacity pricing for your Amazon Quick
// Sight account.
//
// Before you can turn on public sharing on your account, make sure to give public
// sharing permissions to an administrative user in the Identity and Access
// Management (IAM) console. For more information on using IAM with Amazon Quick
// Sight, see [Using Quick Suite with IAM]in the Amazon Quick Sight User Guide.
//
// [Using Quick Suite with IAM]: https://docs.aws.amazon.com/quicksight/latest/user/security_iam_service-with-iam.html
func quicksight_UpdatePublicSharingSettings(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdatePublicSharingSettingsInput{
		// AwsAccountId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightPublicSharingEnabled) > 0 {
		if err := assignInputField(input, "PublicSharingEnabled", _quicksightPublicSharingEnabled); err != nil {
			log.Errorf("invalid --public-sharing-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePublicSharingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a personalization configuration.
func quicksight_UpdateQPersonalizationConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateQPersonalizationConfigurationInput{
		// AwsAccountId: *string, // Required
		// PersonalizationMode: types.PersonalizationMode, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightPersonalizationMode) > 0 {
		if err := assignInputField(input, "PersonalizationMode", _quicksightPersonalizationMode); err != nil {
			log.Errorf("invalid --personalization-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQPersonalizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the state of a Quick Sight Q Search configuration.
func quicksight_UpdateQuickSightQSearchConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateQuickSightQSearchConfigurationInput{
		// AwsAccountId: *string, // Required
		// QSearchStatus: types.QSearchStatus, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightQSearchStatus) > 0 {
		if err := assignInputField(input, "QSearchStatus", _quicksightQSearchStatus); err != nil {
			log.Errorf("invalid --qsearch-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQuickSightQSearchConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a refresh schedule for a dataset.
func quicksight_UpdateRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DataSetId: *string, // Required
		// Schedule: *types.RefreshSchedule, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDataSetId) > 0 {
		input.DataSetId = aws.String(_quicksightDataSetId)
	}
	if len(_quicksightSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _quicksightSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the custom permissions that are associated with a role.
func quicksight_UpdateRoleCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateRoleCustomPermissionInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
		// Namespace: *string, // Required
		// Role: types.Role, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRoleCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a self-upgrade request for a Quick Suite user by approving, denying, or
// verifying the request.
func quicksight_UpdateSelfUpgrade(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateSelfUpgradeInput{
		// Action: types.SelfUpgradeAdminAction, // Required
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// UpgradeRequestId: *string, // Required
	}

	if len(_quicksightAction) > 0 {
		if err := assignInputField(input, "Action", _quicksightAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUpgradeRequestId) > 0 {
		input.UpgradeRequestId = aws.String(_quicksightUpgradeRequestId)
	}

	if resp, err := client.UpdateSelfUpgrade(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the self-upgrade configuration for a Quick Suite account.
func quicksight_UpdateSelfUpgradeConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateSelfUpgradeConfigurationInput{
		// AwsAccountId: *string, // Required
		// Namespace: *string, // Required
		// SelfUpgradeStatus: types.SelfUpgradeStatus, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightSelfUpgradeStatus) > 0 {
		if err := assignInputField(input, "SelfUpgradeStatus", _quicksightSelfUpgradeStatus); err != nil {
			log.Errorf("invalid --self-upgrade-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSelfUpgradeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the SPICE capacity configuration for a Quick Sight account.
func quicksight_UpdateSPICECapacityConfiguration(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateSPICECapacityConfigurationInput{
		// AwsAccountId: *string, // Required
		// PurchaseMode: types.PurchaseMode, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightPurchaseMode) > 0 {
		if err := assignInputField(input, "PurchaseMode", _quicksightPurchaseMode); err != nil {
			log.Errorf("invalid --purchase-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSPICECapacityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a template from an existing Amazon Quick Sight analysis or another
// template.
func quicksight_UpdateTemplate(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateTemplateInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightDefinition) > 0 {
		if err := assignInputField(input, "Definition", _quicksightDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightSourceEntity) > 0 {
		if err := assignInputField(input, "SourceEntity", _quicksightSourceEntity); err != nil {
			log.Errorf("invalid --source-entity: %s", err.Error())
			return
		}
	}
	if len(_quicksightValidationStrategy) > 0 {
		if err := assignInputField(input, "ValidationStrategy", _quicksightValidationStrategy); err != nil {
			log.Errorf("invalid --validation-strategy: %s", err.Error())
			return
		}
	}
	if len(_quicksightVersionDescription) > 0 {
		input.VersionDescription = aws.String(_quicksightVersionDescription)
	}

	if resp, err := client.UpdateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the template alias of a template.
func quicksight_UpdateTemplateAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateTemplateAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
		// TemplateVersionNumber: *int64, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightTemplateVersionNumber) > 0 {
		if err := assignInputField(input, "TemplateVersionNumber", _quicksightTemplateVersionNumber); err != nil {
			log.Errorf("invalid --template-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTemplateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource permissions for a template.
func quicksight_UpdateTemplatePermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateTemplatePermissionsInput{
		// AwsAccountId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTemplateId) > 0 {
		input.TemplateId = aws.String(_quicksightTemplateId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTemplatePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a theme.
func quicksight_UpdateTheme(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateThemeInput{
		// AwsAccountId: *string, // Required
		// BaseThemeId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightBaseThemeId) > 0 {
		input.BaseThemeId = aws.String(_quicksightBaseThemeId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _quicksightConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightVersionDescription) > 0 {
		input.VersionDescription = aws.String(_quicksightVersionDescription)
	}

	if resp, err := client.UpdateTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an alias of a theme.
func quicksight_UpdateThemeAlias(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateThemeAliasInput{
		// AliasName: *string, // Required
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
		// ThemeVersionNumber: *int64, // Required
	}

	if len(_quicksightAliasName) > 0 {
		input.AliasName = aws.String(_quicksightAliasName)
	}
	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightThemeVersionNumber) > 0 {
		if err := assignInputField(input, "ThemeVersionNumber", _quicksightThemeVersionNumber); err != nil {
			log.Errorf("invalid --theme-version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateThemeAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource permissions for a theme. Permissions apply to the action
// to grant or revoke permissions on, for example "quicksight:DescribeTheme" .
//
// Theme permissions apply in groupings. Valid groupings include the following for
// the three levels of permissions, which are user, owner, or no permissions:
//
// - User
//
// - "quicksight:DescribeTheme"
//
// - "quicksight:DescribeThemeAlias"
//
// - "quicksight:ListThemeAliases"
//
// - "quicksight:ListThemeVersions"
//
// - Owner
//
// - "quicksight:DescribeTheme"
//
// - "quicksight:DescribeThemeAlias"
//
// - "quicksight:ListThemeAliases"
//
// - "quicksight:ListThemeVersions"
//
// - "quicksight:DeleteTheme"
//
// - "quicksight:UpdateTheme"
//
// - "quicksight:CreateThemeAlias"
//
// - "quicksight:DeleteThemeAlias"
//
// - "quicksight:UpdateThemeAlias"
//
// - "quicksight:UpdateThemePermissions"
//
// - "quicksight:DescribeThemePermissions"
//
// - To specify no permissions, omit the permissions list.
func quicksight_UpdateThemePermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateThemePermissionsInput{
		// AwsAccountId: *string, // Required
		// ThemeId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightThemeId) > 0 {
		input.ThemeId = aws.String(_quicksightThemeId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateThemePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a topic.
func quicksight_UpdateTopic(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateTopicInput{
		// AwsAccountId: *string, // Required
		// Topic: *types.TopicDetails, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopic) > 0 {
		if err := assignInputField(input, "Topic", _quicksightTopic); err != nil {
			log.Errorf("invalid --topic: %s", err.Error())
			return
		}
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}
	if len(_quicksightCustomInstructions) > 0 {
		if err := assignInputField(input, "CustomInstructions", _quicksightCustomInstructions); err != nil {
			log.Errorf("invalid --custom-instructions: %s", err.Error())
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

// Updates the permissions of a topic.
func quicksight_UpdateTopicPermissions(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateTopicPermissionsInput{
		// AwsAccountId: *string, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}
	if len(_quicksightGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _quicksightGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_quicksightRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _quicksightRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTopicPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a topic refresh schedule.
func quicksight_UpdateTopicRefreshSchedule(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateTopicRefreshScheduleInput{
		// AwsAccountId: *string, // Required
		// DatasetId: *string, // Required
		// RefreshSchedule: *types.TopicRefreshSchedule, // Required
		// TopicId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightDatasetId) > 0 {
		input.DatasetId = aws.String(_quicksightDatasetId)
	}
	if len(_quicksightRefreshSchedule) > 0 {
		if err := assignInputField(input, "RefreshSchedule", _quicksightRefreshSchedule); err != nil {
			log.Errorf("invalid --refresh-schedule: %s", err.Error())
			return
		}
	}
	if len(_quicksightTopicId) > 0 {
		input.TopicId = aws.String(_quicksightTopicId)
	}

	if resp, err := client.UpdateTopicRefreshSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Quick Sight user.
func quicksight_UpdateUser(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateUserInput{
		// AwsAccountId: *string, // Required
		// Email: *string, // Required
		// Namespace: *string, // Required
		// Role: types.UserRole, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightEmail) > 0 {
		input.Email = aws.String(_quicksightEmail)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightRole) > 0 {
		if err := assignInputField(input, "Role", _quicksightRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}
	if len(_quicksightCustomFederationProviderUrl) > 0 {
		input.CustomFederationProviderUrl = aws.String(_quicksightCustomFederationProviderUrl)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}
	if len(_quicksightExternalLoginFederationProviderType) > 0 {
		input.ExternalLoginFederationProviderType = aws.String(_quicksightExternalLoginFederationProviderType)
	}
	if len(_quicksightExternalLoginId) > 0 {
		input.ExternalLoginId = aws.String(_quicksightExternalLoginId)
	}
	if len(_quicksightUnapplyCustomPermissions) > 0 {
		if err := assignInputField(input, "UnapplyCustomPermissions", _quicksightUnapplyCustomPermissions); err != nil {
			log.Errorf("invalid --unapply-custom-permissions: %s", err.Error())
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

// Updates a custom permissions profile for a user.
func quicksight_UpdateUserCustomPermission(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateUserCustomPermissionInput{
		// AwsAccountId: *string, // Required
		// CustomPermissionsName: *string, // Required
		// Namespace: *string, // Required
		// UserName: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightCustomPermissionsName) > 0 {
		input.CustomPermissionsName = aws.String(_quicksightCustomPermissionsName)
	}
	if len(_quicksightNamespace) > 0 {
		input.Namespace = aws.String(_quicksightNamespace)
	}
	if len(_quicksightUserName) > 0 {
		input.UserName = aws.String(_quicksightUserName)
	}

	if resp, err := client.UpdateUserCustomPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a VPC connection.
func quicksight_UpdateVPCConnection(cfg aws.Config, client *quicksight.Client) {
	input := &quicksight.UpdateVPCConnectionInput{
		// AwsAccountId: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// SecurityGroupIds: []string, // Required
		// SubnetIds: []string, // Required
		// VPCConnectionId: *string, // Required
	}

	if len(_quicksightAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_quicksightAwsAccountId)
	}
	if len(_quicksightName) > 0 {
		input.Name = aws.String(_quicksightName)
	}
	if len(_quicksightRoleArn) > 0 {
		input.RoleArn = aws.String(_quicksightRoleArn)
	}
	if len(_quicksightSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _quicksightSecurityGroupIds...)
	}
	if len(_quicksightSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _quicksightSubnetIds...)
	}
	if len(_quicksightVPCConnectionId) > 0 {
		input.VPCConnectionId = aws.String(_quicksightVPCConnectionId)
	}
	if len(_quicksightDnsResolvers) > 0 {
		input.DnsResolvers = append([]string(nil), _quicksightDnsResolvers...)
	}

	if resp, err := client.UpdateVPCConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_quicksightCmd)
	_quicksightCmd.Flags().SortFlags = false

	_quicksightCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_quicksightCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_quicksightCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_quicksightCmd.Flags().StringVarP(&_quicksightAccountCustomization, "account-customization", "", "", "Account Customization")
	_quicksightCmd.Flags().StringVarP(&_quicksightAccountName, "account-name", "", "", "Account Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightAction, "action", "", "", "Action")
	_quicksightCmd.Flags().StringVarP(&_quicksightActionConnectorId, "action-connector-id", "", "", "Action Connector ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightActiveDirectoryName, "active-directory-name", "", "", "Active Directory Name")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAdditionalDashboardIds, "additional-dashboard-ids", "", nil, "Additional Dashboard Ids")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAdminGroup, "admin-group", "", nil, "Admin Group")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAdminProGroup, "admin-pro-group", "", nil, "Admin Pro Group")
	_quicksightCmd.Flags().StringVarP(&_quicksightAliasName, "alias-name", "", "", "Alias Name")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAllowedDomains, "allowed-domains", "", nil, "Allowed Domains")
	_quicksightCmd.Flags().StringVarP(&_quicksightAnalysisId, "analysis-id", "", "", "Analysis ID")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAnswerIds, "answer-ids", "", nil, "Answer Ids")
	_quicksightCmd.Flags().StringVarP(&_quicksightAnswers, "answers", "", "", "Answers")
	_quicksightCmd.Flags().StringVarP(&_quicksightApplicationId, "application-id", "", "", "Application ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightAssetBundleExportJobId, "asset-bundle-export-job-id", "", "", "Asset Bundle Export Job ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightAssetBundleImportJobId, "asset-bundle-import-job-id", "", "", "Asset Bundle Import Job ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightAssetBundleImportSource, "asset-bundle-import-source", "", "", "Asset Bundle Import Source")
	_quicksightCmd.Flags().StringVarP(&_quicksightAssignmentName, "assignment-name", "", "", "Assignment Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightAssignmentStatus, "assignment-status", "", "", "Assignment Status")
	_quicksightCmd.Flags().StringVarP(&_quicksightAuthenticationConfig, "authentication-config", "", "", "Authentication Config")
	_quicksightCmd.Flags().StringVarP(&_quicksightAuthenticationMethod, "authentication-method", "", "", "Authentication Method")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAuthorGroup, "author-group", "", nil, "Author Group")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAuthorProGroup, "author-pro-group", "", nil, "Author Pro Group")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAuthorizedResourceArns, "authorized-resource-arns", "", nil, "Authorized Resource Arns")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightAuthorizedTargets, "authorized-targets", "", nil, "Authorized Targets")
	_quicksightCmd.Flags().StringVarP(&_quicksightAwsAccountId, "aws-account-id", "", "", "AWS Account ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightBaseThemeId, "base-theme-id", "", "", "Base Theme ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightBrandArn, "brand-arn", "", "", "Brand ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightBrandDefinition, "brand-definition", "", "", "Brand Definition")
	_quicksightCmd.Flags().StringVarP(&_quicksightBrandId, "brand-id", "", "", "Brand ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightCapabilities, "capabilities", "", "", "Capabilities")
	_quicksightCmd.Flags().StringVarP(&_quicksightCloudFormationOverridePropertyConfiguration, "cloud-formation-override-property-configuration", "", "", "Cloud Formation Override Property Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightColumnGroups, "column-groups", "", "", "Column Groups")
	_quicksightCmd.Flags().StringVarP(&_quicksightColumnLevelPermissionRules, "column-level-permission-rules", "", "", "Column Level Permission Rules")
	_quicksightCmd.Flags().StringVarP(&_quicksightConfiguration, "configuration", "", "", "Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightContactNumber, "contact-number", "", "", "Contact Number")
	_quicksightCmd.Flags().StringVarP(&_quicksightCredentials, "credentials", "", "", "Credentials")
	_quicksightCmd.Flags().StringVarP(&_quicksightCustomFederationProviderUrl, "custom-federation-provider-url", "", "", "Custom Federation Provider URL")
	_quicksightCmd.Flags().StringVarP(&_quicksightCustomInstructions, "custom-instructions", "", "", "Custom Instructions")
	_quicksightCmd.Flags().StringVarP(&_quicksightCustomPermissionsName, "custom-permissions-name", "", "", "Custom Permissions Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightDashboardId, "dashboard-id", "", "", "Dashboard ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightDashboardPublishOptions, "dashboard-publish-options", "", "", "Dashboard Publish Options")
	_quicksightCmd.Flags().StringVarP(&_quicksightDashboardsQAStatus, "dashboards-qa-status", "", "", "Dashboards Qa Status")
	_quicksightCmd.Flags().StringVarP(&_quicksightDataPrepConfiguration, "data-prep-configuration", "", "", "Data Prep Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightDataSetId, "data-set-id", "", "", "Data Set ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightDataSetRefreshProperties, "data-set-refresh-properties", "", "", "Data Set Refresh Properties")
	_quicksightCmd.Flags().StringVarP(&_quicksightDataSetUsageConfiguration, "data-set-usage-configuration", "", "", "Data Set Usage Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightDataSourceId, "data-source-id", "", "", "Data Source ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightDataSourceParameters, "data-source-parameters", "", "", "Data Source Parameters")
	_quicksightCmd.Flags().StringVarP(&_quicksightDatasetArn, "dataset-arn", "", "", "Dataset ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightDatasetId, "dataset-id", "", "", "Dataset ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightDatasetName, "dataset-name", "", "", "Dataset Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightDatasetParameters, "dataset-parameters", "", "", "Dataset Parameters")
	_quicksightCmd.Flags().StringVarP(&_quicksightDefaultKeyOnly, "default-key-only", "", "", "Default Key Only")
	_quicksightCmd.Flags().StringVarP(&_quicksightDefaultNamespace, "default-namespace", "", "", "Default Namespace")
	_quicksightCmd.Flags().StringVarP(&_quicksightDefinition, "definition", "", "", "Definition")
	_quicksightCmd.Flags().StringVarP(&_quicksightDescription, "description", "", "", "Description")
	_quicksightCmd.Flags().StringVarP(&_quicksightDirectoryId, "directory-id", "", "", "Directory ID")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightDnsResolvers, "dns-resolvers", "", nil, "DNS Resolvers")
	_quicksightCmd.Flags().StringVarP(&_quicksightEdition, "edition", "", "", "Edition")
	_quicksightCmd.Flags().StringVarP(&_quicksightEmail, "email", "", "", "Email")
	_quicksightCmd.Flags().StringVarP(&_quicksightEmailAddress, "email-address", "", "", "Email Address")
	_quicksightCmd.Flags().StringVarP(&_quicksightEnabled, "enabled", "", "", "Enabled")
	_quicksightCmd.Flags().StringVarP(&_quicksightEntryPoint, "entry-point", "", "", "Entry Point")
	_quicksightCmd.Flags().StringVarP(&_quicksightExperienceConfiguration, "experience-configuration", "", "", "Experience Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightExportFormat, "export-format", "", "", "Export Format")
	_quicksightCmd.Flags().StringVarP(&_quicksightExternalLoginFederationProviderType, "external-login-federation-provider-type", "", "", "External Login Federation Provider Type")
	_quicksightCmd.Flags().StringVarP(&_quicksightExternalLoginId, "external-login-id", "", "", "External Login ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightFailureAction, "failure-action", "", "", "Failure Action")
	_quicksightCmd.Flags().StringVarP(&_quicksightFieldFolders, "field-folders", "", "", "Field Folders")
	_quicksightCmd.Flags().StringVarP(&_quicksightFilters, "filters", "", "", "Filters")
	_quicksightCmd.Flags().StringVarP(&_quicksightFirstName, "first-name", "", "", "First Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightFlowId, "flow-id", "", "", "Flow ID")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightFolderArns, "folder-arns", "", nil, "Folder Arns")
	_quicksightCmd.Flags().StringVarP(&_quicksightFolderId, "folder-id", "", "", "Folder ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightFolderType, "folder-type", "", "", "Folder Type")
	_quicksightCmd.Flags().StringVarP(&_quicksightForceDeleteWithoutRecovery, "force-delete-without-recovery", "", "", "Force Delete Without Recovery")
	_quicksightCmd.Flags().StringVarP(&_quicksightGrantLinkPermissions, "grant-link-permissions", "", "", "Grant Link Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightGrantPermissions, "grant-permissions", "", "", "Grant Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightGroupName, "group-name", "", "", "Group Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightIamArn, "iam-arn", "", "", "IAM ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightIAMIdentityCenterInstanceArn, "iam-identity-center-instance-arn", "", "", "IAM Identity Center Instance ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightIdentities, "identities", "", "", "Identities")
	_quicksightCmd.Flags().StringVarP(&_quicksightIdentityStore, "identity-store", "", "", "Identity Store")
	_quicksightCmd.Flags().StringVarP(&_quicksightIdentityType, "identity-type", "", "", "Identity Type")
	_quicksightCmd.Flags().StringVarP(&_quicksightImportMode, "import-mode", "", "", "Import Mode")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludeAllDependencies, "include-all-dependencies", "", "", "Include All Dependencies")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludeFolderMembers, "include-folder-members", "", "", "Include Folder Members")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludeFolderMemberships, "include-folder-memberships", "", "", "Include Folder Memberships")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludeGeneratedAnswer, "include-generated-answer", "", "", "Include Generated Answer")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludePermissions, "include-permissions", "", "", "Include Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludeQuickSightQIndex, "include-quicksight-q-index", "", "", "Include Quicksight Q Index")
	_quicksightCmd.Flags().StringVarP(&_quicksightIncludeTags, "include-tags", "", "", "Include Tags")
	_quicksightCmd.Flags().StringVarP(&_quicksightIngestionId, "ingestion-id", "", "", "Ingestion ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightIngestionType, "ingestion-type", "", "", "Ingestion Type")
	_quicksightCmd.Flags().StringVarP(&_quicksightIpRestrictionRuleMap, "ip-restriction-rule-map", "", "", "IP Restriction Rule Map")
	_quicksightCmd.Flags().StringVarP(&_quicksightKeyRegistration, "key-registration", "", "", "Key Registration")
	_quicksightCmd.Flags().StringVarP(&_quicksightLastName, "last-name", "", "", "Last Name")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightLinkEntities, "link-entities", "", nil, "Link Entities")
	_quicksightCmd.Flags().StringVarP(&_quicksightLinkSharingConfiguration, "link-sharing-configuration", "", "", "Link Sharing Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightLogicalTableMap, "logical-table-map", "", "", "Logical Table Map")
	_quicksightCmd.Flags().StringVarP(&_quicksightMaxResults, "max-results", "", "", "Max Results")
	_quicksightCmd.Flags().StringVarP(&_quicksightMaxTopicsToConsider, "max-topics-to-consider", "", "", "Max Topics To Consider")
	_quicksightCmd.Flags().StringVarP(&_quicksightMemberId, "member-id", "", "", "Member ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightMemberName, "member-name", "", "", "Member Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightMemberType, "member-type", "", "", "Member Type")
	_quicksightCmd.Flags().StringVarP(&_quicksightName, "name", "", "", "Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightNamespace, "namespace", "", "", "Namespace")
	_quicksightCmd.Flags().StringVarP(&_quicksightNextToken, "next-token", "", "", "Next Token")
	_quicksightCmd.Flags().StringVarP(&_quicksightNotificationEmail, "notification-email", "", "", "Notification Email")
	_quicksightCmd.Flags().StringVarP(&_quicksightOverrideParameters, "override-parameters", "", "", "Override Parameters")
	_quicksightCmd.Flags().StringVarP(&_quicksightOverridePermissions, "override-permissions", "", "", "Override Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightOverrideTags, "override-tags", "", "", "Override Tags")
	_quicksightCmd.Flags().StringVarP(&_quicksightOverrideValidationStrategy, "override-validation-strategy", "", "", "Override Validation Strategy")
	_quicksightCmd.Flags().StringVarP(&_quicksightParameters, "parameters", "", "", "Parameters")
	_quicksightCmd.Flags().StringVarP(&_quicksightParentFolderArn, "parent-folder-arn", "", "", "Parent Folder ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightPerformanceConfiguration, "performance-configuration", "", "", "Performance Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightPermissions, "permissions", "", "", "Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightPersonalizationMode, "personalization-mode", "", "", "Personalization Mode")
	_quicksightCmd.Flags().StringVarP(&_quicksightPhysicalTableMap, "physical-table-map", "", "", "Physical Table Map")
	_quicksightCmd.Flags().StringVarP(&_quicksightPolicyArn, "policy-arn", "", "", "Policy ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightPrincipalId, "principal-id", "", "", "Principal ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightPublicSharingEnabled, "public-sharing-enabled", "", "", "Public Sharing Enabled")
	_quicksightCmd.Flags().StringVarP(&_quicksightPurchaseMode, "purchase-mode", "", "", "Purchase Mode")
	_quicksightCmd.Flags().StringVarP(&_quicksightQSearchStatus, "qsearch-status", "", "", "Qsearch Status")
	_quicksightCmd.Flags().StringVarP(&_quicksightQueryText, "query-text", "", "", "Query Text")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightReaderGroup, "reader-group", "", nil, "Reader Group")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightReaderProGroup, "reader-pro-group", "", nil, "Reader Pro Group")
	_quicksightCmd.Flags().StringVarP(&_quicksightRealm, "realm", "", "", "Realm")
	_quicksightCmd.Flags().StringVarP(&_quicksightRecoveryWindowInDays, "recovery-window-in-days", "", "", "Recovery Window In Days")
	_quicksightCmd.Flags().StringVarP(&_quicksightRefreshId, "refresh-id", "", "", "Refresh ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightRefreshSchedule, "refresh-schedule", "", "", "Refresh Schedule")
	_quicksightCmd.Flags().StringVarP(&_quicksightResetDisabled, "reset-disabled", "", "", "Reset Disabled")
	_quicksightCmd.Flags().StringVarP(&_quicksightResolved, "resolved", "", "", "Resolved")
	_quicksightCmd.Flags().StringVarP(&_quicksightResourceArn, "resource-arn", "", "", "Resource ARN")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightResourceArns, "resource-arns", "", nil, "Resource Arns")
	_quicksightCmd.Flags().StringVarP(&_quicksightRestoreToFolders, "restore-to-folders", "", "", "Restore To Folders")
	_quicksightCmd.Flags().StringVarP(&_quicksightRevokeLinkPermissions, "revoke-link-permissions", "", "", "Revoke Link Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightRevokePermissions, "revoke-permissions", "", "", "Revoke Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightRole, "role", "", "", "Role")
	_quicksightCmd.Flags().StringVarP(&_quicksightRoleArn, "role-arn", "", "", "Role ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightRowLevelPermissionDataSet, "row-level-permission-data-set", "", "", "Row Level Permission Data Set")
	_quicksightCmd.Flags().StringVarP(&_quicksightRowLevelPermissionTagConfiguration, "row-level-permission-tag-configuration", "", "", "Row Level Permission Tag Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightSchedule, "schedule", "", "", "Schedule")
	_quicksightCmd.Flags().StringVarP(&_quicksightScheduleId, "schedule-id", "", "", "Schedule ID")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_quicksightCmd.Flags().StringVarP(&_quicksightSelfUpgradeStatus, "self-upgrade-status", "", "", "Self Upgrade Status")
	_quicksightCmd.Flags().StringVarP(&_quicksightSemanticModelConfiguration, "semantic-model-configuration", "", "", "Semantic Model Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightService, "service", "", "", "Service")
	_quicksightCmd.Flags().StringVarP(&_quicksightSessionExpiresAt, "session-expires-at", "", "", "Session Expires At")
	_quicksightCmd.Flags().StringVarP(&_quicksightSessionLifetimeInMinutes, "session-lifetime-in-minutes", "", "", "Session Lifetime In Minutes")
	_quicksightCmd.Flags().StringVarP(&_quicksightSessionName, "session-name", "", "", "Session Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightSessionTags, "session-tags", "", "", "Session Tags")
	_quicksightCmd.Flags().StringVarP(&_quicksightSharingModel, "sharing-model", "", "", "Sharing Model")
	_quicksightCmd.Flags().StringVarP(&_quicksightSnapshotConfiguration, "snapshot-configuration", "", "", "Snapshot Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightSnapshotJobId, "snapshot-job-id", "", "", "Snapshot Job ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightSourceEntity, "source-entity", "", "", "Source Entity")
	_quicksightCmd.Flags().StringVarP(&_quicksightSslProperties, "ssl-properties", "", "", "SSL Properties")
	_quicksightCmd.Flags().StringVarP(&_quicksightStatePersistenceEnabled, "state-persistence-enabled", "", "", "State Persistence Enabled")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_quicksightCmd.Flags().StringSliceVarP(&_quicksightTagKeys, "tag-keys", "", nil, "Tag Keys")
	_quicksightCmd.Flags().StringVarP(&_quicksightTags, "tags", "", "", "Tags")
	_quicksightCmd.Flags().StringVarP(&_quicksightTemplateId, "template-id", "", "", "Template ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightTemplateVersionNumber, "template-version-number", "", "", "Template Version Number")
	_quicksightCmd.Flags().StringVarP(&_quicksightTerminationProtectionEnabled, "termination-protection-enabled", "", "", "Termination Protection Enabled")
	_quicksightCmd.Flags().StringVarP(&_quicksightThemeArn, "theme-arn", "", "", "Theme ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightThemeId, "theme-id", "", "", "Theme ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightThemeVersionNumber, "theme-version-number", "", "", "Theme Version Number")
	_quicksightCmd.Flags().StringVarP(&_quicksightTopic, "topic", "", "", "Topic")
	_quicksightCmd.Flags().StringVarP(&_quicksightTopicId, "topic-id", "", "", "Topic ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightType, "type", "", "", "Type")
	_quicksightCmd.Flags().StringVarP(&_quicksightUnapplyCustomPermissions, "unapply-custom-permissions", "", "", "Unapply Custom Permissions")
	_quicksightCmd.Flags().StringVarP(&_quicksightUndoRedoDisabled, "undo-redo-disabled", "", "", "Undo Redo Disabled")
	_quicksightCmd.Flags().StringVarP(&_quicksightUpgradeRequestId, "upgrade-request-id", "", "", "Upgrade Request ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightUseAs, "use-as", "", "", "Use As")
	_quicksightCmd.Flags().StringVarP(&_quicksightUserArn, "user-arn", "", "", "User ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightUserConfiguration, "user-configuration", "", "", "User Configuration")
	_quicksightCmd.Flags().StringVarP(&_quicksightUserIdentifier, "user-identifier", "", "", "User Identifier")
	_quicksightCmd.Flags().StringVarP(&_quicksightUserName, "user-name", "", "", "User Name")
	_quicksightCmd.Flags().StringVarP(&_quicksightUserRole, "user-role", "", "", "User Role")
	_quicksightCmd.Flags().StringVarP(&_quicksightValidationStrategy, "validation-strategy", "", "", "Validation Strategy")
	_quicksightCmd.Flags().StringVarP(&_quicksightVersionDescription, "version-description", "", "", "Version Description")
	_quicksightCmd.Flags().StringVarP(&_quicksightVersionId, "version-id", "", "", "Version ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightVersionNumber, "version-number", "", "", "Version Number")
	_quicksightCmd.Flags().StringVarP(&_quicksightVpcConnectionArn, "vpc-connection-arn", "", "", "VPC Connection ARN")
	_quicksightCmd.Flags().StringVarP(&_quicksightVPCConnectionId, "vpc-connection-id", "", "", "VPC Connection ID")
	_quicksightCmd.Flags().StringVarP(&_quicksightVpcConnectionProperties, "vpc-connection-properties", "", "", "VPC Connection Properties")
	_quicksightCmd.Flags().StringVarP(&_quicksightVpcEndpointIdRestrictionRuleMap, "vpc-endpoint-id-restriction-rule-map", "", "", "VPC Endpoint ID Restriction Rule Map")
	_quicksightCmd.Flags().StringVarP(&_quicksightVpcIdRestrictionRuleMap, "vpc-id-restriction-rule-map", "", "", "VPC ID Restriction Rule Map")

	_quicksightCmd.Flags().BoolVarP(&_quicksightBatchCreateTopicReviewedAnswer, "batch-create-topic-reviewed-answer", "", false, "Batch Create Topic Reviewed Answer")
	_quicksightCmd.Flags().BoolVarP(&_quicksightBatchDeleteTopicReviewedAnswer, "batch-delete-topic-reviewed-answer", "", false, "Batch Delete Topic Reviewed Answer")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCancelIngestion, "cancel-ingestion", "", false, "Cancel Ingestion")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateAccountCustomization, "create-account-customization", "", false, "Create Account Customization")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateAccountSubscription, "create-account-subscription", "", false, "Create Account Subscription")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateActionConnector, "create-action-connector", "", false, "Create Action Connector")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateAnalysis, "create-analysis", "", false, "Create Analysis")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateBrand, "create-brand", "", false, "Create Brand")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateCustomPermissions, "create-custom-permissions", "", false, "Create Custom Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateDashboard, "create-dashboard", "", false, "Create Dashboard")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateDataSet, "create-data-set", "", false, "Create Data Set")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateDataSource, "create-data-source", "", false, "Create Data Source")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateFolder, "create-folder", "", false, "Create Folder")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateFolderMembership, "create-folder-membership", "", false, "Create Folder Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateGroup, "create-group", "", false, "Create Group")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateGroupMembership, "create-group-membership", "", false, "Create Group Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateIAMPolicyAssignment, "create-iam-policy-assignment", "", false, "Create IAM Policy Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateIngestion, "create-ingestion", "", false, "Create Ingestion")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateNamespace, "create-namespace", "", false, "Create Namespace")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateRefreshSchedule, "create-refresh-schedule", "", false, "Create Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateRoleMembership, "create-role-membership", "", false, "Create Role Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateTemplate, "create-template", "", false, "Create Template")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateTemplateAlias, "create-template-alias", "", false, "Create Template Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateTheme, "create-theme", "", false, "Create Theme")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateThemeAlias, "create-theme-alias", "", false, "Create Theme Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateTopic, "create-topic", "", false, "Create Topic")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateTopicRefreshSchedule, "create-topic-refresh-schedule", "", false, "Create Topic Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightCreateVPCConnection, "create-vpc-connection", "", false, "Create VPC Connection")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteAccountCustomPermission, "delete-account-custom-permission", "", false, "Delete Account Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteAccountCustomization, "delete-account-customization", "", false, "Delete Account Customization")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteAccountSubscription, "delete-account-subscription", "", false, "Delete Account Subscription")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteActionConnector, "delete-action-connector", "", false, "Delete Action Connector")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteAnalysis, "delete-analysis", "", false, "Delete Analysis")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteBrand, "delete-brand", "", false, "Delete Brand")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteBrandAssignment, "delete-brand-assignment", "", false, "Delete Brand Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteCustomPermissions, "delete-custom-permissions", "", false, "Delete Custom Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteDashboard, "delete-dashboard", "", false, "Delete Dashboard")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteDataSet, "delete-data-set", "", false, "Delete Data Set")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteDataSetRefreshProperties, "delete-data-set-refresh-properties", "", false, "Delete Data Set Refresh Properties")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteDefaultQBusinessApplication, "delete-default-qbusiness-application", "", false, "Delete Default Qbusiness Application")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteFolder, "delete-folder", "", false, "Delete Folder")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteFolderMembership, "delete-folder-membership", "", false, "Delete Folder Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteGroup, "delete-group", "", false, "Delete Group")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteGroupMembership, "delete-group-membership", "", false, "Delete Group Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteIAMPolicyAssignment, "delete-iam-policy-assignment", "", false, "Delete IAM Policy Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteIdentityPropagationConfig, "delete-identity-propagation-config", "", false, "Delete Identity Propagation Config")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteNamespace, "delete-namespace", "", false, "Delete Namespace")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteRefreshSchedule, "delete-refresh-schedule", "", false, "Delete Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteRoleCustomPermission, "delete-role-custom-permission", "", false, "Delete Role Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteRoleMembership, "delete-role-membership", "", false, "Delete Role Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteTemplate, "delete-template", "", false, "Delete Template")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteTemplateAlias, "delete-template-alias", "", false, "Delete Template Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteTheme, "delete-theme", "", false, "Delete Theme")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteThemeAlias, "delete-theme-alias", "", false, "Delete Theme Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteTopic, "delete-topic", "", false, "Delete Topic")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteTopicRefreshSchedule, "delete-topic-refresh-schedule", "", false, "Delete Topic Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteUser, "delete-user", "", false, "Delete User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteUserByPrincipalId, "delete-user-by-principal-id", "", false, "Delete User By Principal ID")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteUserCustomPermission, "delete-user-custom-permission", "", false, "Delete User Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDeleteVPCConnection, "delete-vpc-connection", "", false, "Delete VPC Connection")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAccountCustomPermission, "describe-account-custom-permission", "", false, "Describe Account Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAccountCustomization, "describe-account-customization", "", false, "Describe Account Customization")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAccountSettings, "describe-account-settings", "", false, "Describe Account Settings")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAccountSubscription, "describe-account-subscription", "", false, "Describe Account Subscription")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeActionConnector, "describe-action-connector", "", false, "Describe Action Connector")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeActionConnectorPermissions, "describe-action-connector-permissions", "", false, "Describe Action Connector Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAnalysis, "describe-analysis", "", false, "Describe Analysis")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAnalysisDefinition, "describe-analysis-definition", "", false, "Describe Analysis Definition")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAnalysisPermissions, "describe-analysis-permissions", "", false, "Describe Analysis Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAssetBundleExportJob, "describe-asset-bundle-export-job", "", false, "Describe Asset Bundle Export Job")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeAssetBundleImportJob, "describe-asset-bundle-import-job", "", false, "Describe Asset Bundle Import Job")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeBrand, "describe-brand", "", false, "Describe Brand")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeBrandAssignment, "describe-brand-assignment", "", false, "Describe Brand Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeBrandPublishedVersion, "describe-brand-published-version", "", false, "Describe Brand Published Version")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeCustomPermissions, "describe-custom-permissions", "", false, "Describe Custom Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDashboard, "describe-dashboard", "", false, "Describe Dashboard")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDashboardDefinition, "describe-dashboard-definition", "", false, "Describe Dashboard Definition")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDashboardPermissions, "describe-dashboard-permissions", "", false, "Describe Dashboard Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDashboardSnapshotJob, "describe-dashboard-snapshot-job", "", false, "Describe Dashboard Snapshot Job")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDashboardSnapshotJobResult, "describe-dashboard-snapshot-job-result", "", false, "Describe Dashboard Snapshot Job Result")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDashboardsQAConfiguration, "describe-dashboards-qa-configuration", "", false, "Describe Dashboards Qa Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDataSet, "describe-data-set", "", false, "Describe Data Set")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDataSetPermissions, "describe-data-set-permissions", "", false, "Describe Data Set Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDataSetRefreshProperties, "describe-data-set-refresh-properties", "", false, "Describe Data Set Refresh Properties")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDataSource, "describe-data-source", "", false, "Describe Data Source")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDataSourcePermissions, "describe-data-source-permissions", "", false, "Describe Data Source Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeDefaultQBusinessApplication, "describe-default-qbusiness-application", "", false, "Describe Default Qbusiness Application")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeFolder, "describe-folder", "", false, "Describe Folder")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeFolderPermissions, "describe-folder-permissions", "", false, "Describe Folder Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeFolderResolvedPermissions, "describe-folder-resolved-permissions", "", false, "Describe Folder Resolved Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeGroup, "describe-group", "", false, "Describe Group")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeGroupMembership, "describe-group-membership", "", false, "Describe Group Membership")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeIAMPolicyAssignment, "describe-iam-policy-assignment", "", false, "Describe IAM Policy Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeIngestion, "describe-ingestion", "", false, "Describe Ingestion")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeIpRestriction, "describe-ip-restriction", "", false, "Describe IP Restriction")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeKeyRegistration, "describe-key-registration", "", false, "Describe Key Registration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeNamespace, "describe-namespace", "", false, "Describe Namespace")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeQPersonalizationConfiguration, "describe-qpersonalization-configuration", "", false, "Describe Qpersonalization Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeQuickSightQSearchConfiguration, "describe-quicksight-qsearch-configuration", "", false, "Describe Quicksight Qsearch Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeRefreshSchedule, "describe-refresh-schedule", "", false, "Describe Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeRoleCustomPermission, "describe-role-custom-permission", "", false, "Describe Role Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeSelfUpgradeConfiguration, "describe-self-upgrade-configuration", "", false, "Describe Self Upgrade Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTemplate, "describe-template", "", false, "Describe Template")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTemplateAlias, "describe-template-alias", "", false, "Describe Template Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTemplateDefinition, "describe-template-definition", "", false, "Describe Template Definition")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTemplatePermissions, "describe-template-permissions", "", false, "Describe Template Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTheme, "describe-theme", "", false, "Describe Theme")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeThemeAlias, "describe-theme-alias", "", false, "Describe Theme Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeThemePermissions, "describe-theme-permissions", "", false, "Describe Theme Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTopic, "describe-topic", "", false, "Describe Topic")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTopicPermissions, "describe-topic-permissions", "", false, "Describe Topic Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTopicRefresh, "describe-topic-refresh", "", false, "Describe Topic Refresh")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeTopicRefreshSchedule, "describe-topic-refresh-schedule", "", false, "Describe Topic Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeUser, "describe-user", "", false, "Describe User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightDescribeVPCConnection, "describe-vpc-connection", "", false, "Describe VPC Connection")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGenerateEmbedUrlForAnonymousUser, "generate-embed-url-for-anonymous-user", "", false, "Generate Embed URL For Anonymous User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGenerateEmbedUrlForRegisteredUser, "generate-embed-url-for-registered-user", "", false, "Generate Embed URL For Registered User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGenerateEmbedUrlForRegisteredUserWithIdentity, "generate-embed-url-for-registered-user-with-identity", "", false, "Generate Embed URL For Registered User With Identity")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGetDashboardEmbedUrl, "get-dashboard-embed-url", "", false, "Get Dashboard Embed URL")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGetFlowMetadata, "get-flow-metadata", "", false, "Get Flow Metadata")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGetFlowPermissions, "get-flow-permissions", "", false, "Get Flow Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGetIdentityContext, "get-identity-context", "", false, "Get Identity Context")
	_quicksightCmd.Flags().BoolVarP(&_quicksightGetSessionEmbedUrl, "get-session-embed-url", "", false, "Get Session Embed URL")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListActionConnectors, "list-action-connectors", "", false, "List Action Connectors")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListAnalyses, "list-analyses", "", false, "List Analyses")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListAssetBundleExportJobs, "list-asset-bundle-export-jobs", "", false, "List Asset Bundle Export Jobs")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListAssetBundleImportJobs, "list-asset-bundle-import-jobs", "", false, "List Asset Bundle Import Jobs")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListBrands, "list-brands", "", false, "List Brands")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListCustomPermissions, "list-custom-permissions", "", false, "List Custom Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListDashboardVersions, "list-dashboard-versions", "", false, "List Dashboard Versions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListDashboards, "list-dashboards", "", false, "List Dashboards")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListDataSets, "list-data-sets", "", false, "List Data Sets")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListDataSources, "list-data-sources", "", false, "List Data Sources")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListFlows, "list-flows", "", false, "List Flows")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListFolderMembers, "list-folder-members", "", false, "List Folder Members")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListFolders, "list-folders", "", false, "List Folders")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListFoldersForResource, "list-folders-for-resource", "", false, "List Folders For Resource")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListGroupMemberships, "list-group-memberships", "", false, "List Group Memberships")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListGroups, "list-groups", "", false, "List Groups")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListIAMPolicyAssignments, "list-iam-policy-assignments", "", false, "List IAM Policy Assignments")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListIAMPolicyAssignmentsForUser, "list-iam-policy-assignments-for-user", "", false, "List IAM Policy Assignments For User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListIdentityPropagationConfigs, "list-identity-propagation-configs", "", false, "List Identity Propagation Configs")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListIngestions, "list-ingestions", "", false, "List Ingestions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListNamespaces, "list-namespaces", "", false, "List Namespaces")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListRefreshSchedules, "list-refresh-schedules", "", false, "List Refresh Schedules")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListRoleMemberships, "list-role-memberships", "", false, "List Role Memberships")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListSelfUpgrades, "list-self-upgrades", "", false, "List Self Upgrades")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTemplateAliases, "list-template-aliases", "", false, "List Template Aliases")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTemplateVersions, "list-template-versions", "", false, "List Template Versions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTemplates, "list-templates", "", false, "List Templates")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListThemeAliases, "list-theme-aliases", "", false, "List Theme Aliases")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListThemeVersions, "list-theme-versions", "", false, "List Theme Versions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListThemes, "list-themes", "", false, "List Themes")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTopicRefreshSchedules, "list-topic-refresh-schedules", "", false, "List Topic Refresh Schedules")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTopicReviewedAnswers, "list-topic-reviewed-answers", "", false, "List Topic Reviewed Answers")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListTopics, "list-topics", "", false, "List Topics")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListUserGroups, "list-user-groups", "", false, "List User Groups")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListUsers, "list-users", "", false, "List Users")
	_quicksightCmd.Flags().BoolVarP(&_quicksightListVPCConnections, "list-vpc-connections", "", false, "List VPC Connections")
	_quicksightCmd.Flags().BoolVarP(&_quicksightPredictQAResults, "predict-qa-results", "", false, "Predict Qa Results")
	_quicksightCmd.Flags().BoolVarP(&_quicksightPutDataSetRefreshProperties, "put-data-set-refresh-properties", "", false, "Put Data Set Refresh Properties")
	_quicksightCmd.Flags().BoolVarP(&_quicksightRegisterUser, "register-user", "", false, "Register User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightRestoreAnalysis, "restore-analysis", "", false, "Restore Analysis")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchActionConnectors, "search-action-connectors", "", false, "Search Action Connectors")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchAnalyses, "search-analyses", "", false, "Search Analyses")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchDashboards, "search-dashboards", "", false, "Search Dashboards")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchDataSets, "search-data-sets", "", false, "Search Data Sets")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchDataSources, "search-data-sources", "", false, "Search Data Sources")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchFlows, "search-flows", "", false, "Search Flows")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchFolders, "search-folders", "", false, "Search Folders")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchGroups, "search-groups", "", false, "Search Groups")
	_quicksightCmd.Flags().BoolVarP(&_quicksightSearchTopics, "search-topics", "", false, "Search Topics")
	_quicksightCmd.Flags().BoolVarP(&_quicksightStartAssetBundleExportJob, "start-asset-bundle-export-job", "", false, "Start Asset Bundle Export Job")
	_quicksightCmd.Flags().BoolVarP(&_quicksightStartAssetBundleImportJob, "start-asset-bundle-import-job", "", false, "Start Asset Bundle Import Job")
	_quicksightCmd.Flags().BoolVarP(&_quicksightStartDashboardSnapshotJob, "start-dashboard-snapshot-job", "", false, "Start Dashboard Snapshot Job")
	_quicksightCmd.Flags().BoolVarP(&_quicksightStartDashboardSnapshotJobSchedule, "start-dashboard-snapshot-job-schedule", "", false, "Start Dashboard Snapshot Job Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightTagResource, "tag-resource", "", false, "Tag Resource")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUntagResource, "untag-resource", "", false, "Untag Resource")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateAccountCustomPermission, "update-account-custom-permission", "", false, "Update Account Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateAccountCustomization, "update-account-customization", "", false, "Update Account Customization")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateActionConnector, "update-action-connector", "", false, "Update Action Connector")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateActionConnectorPermissions, "update-action-connector-permissions", "", false, "Update Action Connector Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateAnalysis, "update-analysis", "", false, "Update Analysis")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateAnalysisPermissions, "update-analysis-permissions", "", false, "Update Analysis Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateApplicationWithTokenExchangeGrant, "update-application-with-token-exchange-grant", "", false, "Update Application With Token Exchange Grant")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateBrand, "update-brand", "", false, "Update Brand")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateBrandAssignment, "update-brand-assignment", "", false, "Update Brand Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateBrandPublishedVersion, "update-brand-published-version", "", false, "Update Brand Published Version")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateCustomPermissions, "update-custom-permissions", "", false, "Update Custom Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDashboard, "update-dashboard", "", false, "Update Dashboard")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDashboardLinks, "update-dashboard-links", "", false, "Update Dashboard Links")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDashboardPermissions, "update-dashboard-permissions", "", false, "Update Dashboard Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDashboardPublishedVersion, "update-dashboard-published-version", "", false, "Update Dashboard Published Version")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDashboardsQAConfiguration, "update-dashboards-qa-configuration", "", false, "Update Dashboards Qa Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDataSet, "update-data-set", "", false, "Update Data Set")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDataSetPermissions, "update-data-set-permissions", "", false, "Update Data Set Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDataSourcePermissions, "update-data-source-permissions", "", false, "Update Data Source Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateDefaultQBusinessApplication, "update-default-qbusiness-application", "", false, "Update Default Qbusiness Application")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateFlowPermissions, "update-flow-permissions", "", false, "Update Flow Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateFolder, "update-folder", "", false, "Update Folder")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateFolderPermissions, "update-folder-permissions", "", false, "Update Folder Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateGroup, "update-group", "", false, "Update Group")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateIAMPolicyAssignment, "update-iam-policy-assignment", "", false, "Update IAM Policy Assignment")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateIdentityPropagationConfig, "update-identity-propagation-config", "", false, "Update Identity Propagation Config")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateIpRestriction, "update-ip-restriction", "", false, "Update IP Restriction")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateKeyRegistration, "update-key-registration", "", false, "Update Key Registration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdatePublicSharingSettings, "update-public-sharing-settings", "", false, "Update Public Sharing Settings")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateQPersonalizationConfiguration, "update-qpersonalization-configuration", "", false, "Update Qpersonalization Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateQuickSightQSearchConfiguration, "update-quicksight-qsearch-configuration", "", false, "Update Quicksight Qsearch Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateRefreshSchedule, "update-refresh-schedule", "", false, "Update Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateRoleCustomPermission, "update-role-custom-permission", "", false, "Update Role Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateSelfUpgrade, "update-self-upgrade", "", false, "Update Self Upgrade")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateSelfUpgradeConfiguration, "update-self-upgrade-configuration", "", false, "Update Self Upgrade Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateSPICECapacityConfiguration, "update-spice-capacity-configuration", "", false, "Update Spice Capacity Configuration")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTemplate, "update-template", "", false, "Update Template")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTemplateAlias, "update-template-alias", "", false, "Update Template Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTemplatePermissions, "update-template-permissions", "", false, "Update Template Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTheme, "update-theme", "", false, "Update Theme")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateThemeAlias, "update-theme-alias", "", false, "Update Theme Alias")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateThemePermissions, "update-theme-permissions", "", false, "Update Theme Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTopic, "update-topic", "", false, "Update Topic")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTopicPermissions, "update-topic-permissions", "", false, "Update Topic Permissions")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateTopicRefreshSchedule, "update-topic-refresh-schedule", "", false, "Update Topic Refresh Schedule")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateUser, "update-user", "", false, "Update User")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateUserCustomPermission, "update-user-custom-permission", "", false, "Update User Custom Permission")
	_quicksightCmd.Flags().BoolVarP(&_quicksightUpdateVPCConnection, "update-vpc-connection", "", false, "Update VPC Connection")

}
