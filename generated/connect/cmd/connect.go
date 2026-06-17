package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var _connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "AWS connect CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := connect.NewFromConfig(cfg)
		if _connectActivateEvaluationForm {
			connect_ActivateEvaluationForm(cfg, client)
			return
		}
		if _connectAssociateAnalyticsDataSet {
			connect_AssociateAnalyticsDataSet(cfg, client)
			return
		}
		if _connectAssociateApprovedOrigin {
			connect_AssociateApprovedOrigin(cfg, client)
			return
		}
		if _connectAssociateBot {
			connect_AssociateBot(cfg, client)
			return
		}
		if _connectAssociateContactWithUser {
			connect_AssociateContactWithUser(cfg, client)
			return
		}
		if _connectAssociateDefaultVocabulary {
			connect_AssociateDefaultVocabulary(cfg, client)
			return
		}
		if _connectAssociateEmailAddressAlias {
			connect_AssociateEmailAddressAlias(cfg, client)
			return
		}
		if _connectAssociateFlow {
			connect_AssociateFlow(cfg, client)
			return
		}
		if _connectAssociateHoursOfOperations {
			connect_AssociateHoursOfOperations(cfg, client)
			return
		}
		if _connectAssociateInstanceStorageConfig {
			connect_AssociateInstanceStorageConfig(cfg, client)
			return
		}
		if _connectAssociateLambdaFunction {
			connect_AssociateLambdaFunction(cfg, client)
			return
		}
		if _connectAssociateLexBot {
			connect_AssociateLexBot(cfg, client)
			return
		}
		if _connectAssociatePhoneNumberContactFlow {
			connect_AssociatePhoneNumberContactFlow(cfg, client)
			return
		}
		if _connectAssociateQueueQuickConnects {
			connect_AssociateQueueQuickConnects(cfg, client)
			return
		}
		if _connectAssociateRoutingProfileQueues {
			connect_AssociateRoutingProfileQueues(cfg, client)
			return
		}
		if _connectAssociateSecurityKey {
			connect_AssociateSecurityKey(cfg, client)
			return
		}
		if _connectAssociateSecurityProfiles {
			connect_AssociateSecurityProfiles(cfg, client)
			return
		}
		if _connectAssociateTrafficDistributionGroupUser {
			connect_AssociateTrafficDistributionGroupUser(cfg, client)
			return
		}
		if _connectAssociateUserProficiencies {
			connect_AssociateUserProficiencies(cfg, client)
			return
		}
		if _connectAssociateWorkspace {
			connect_AssociateWorkspace(cfg, client)
			return
		}
		if _connectBatchAssociateAnalyticsDataSet {
			connect_BatchAssociateAnalyticsDataSet(cfg, client)
			return
		}
		if _connectBatchCreateDataTableValue {
			connect_BatchCreateDataTableValue(cfg, client)
			return
		}
		if _connectBatchDeleteDataTableValue {
			connect_BatchDeleteDataTableValue(cfg, client)
			return
		}
		if _connectBatchDescribeDataTableValue {
			connect_BatchDescribeDataTableValue(cfg, client)
			return
		}
		if _connectBatchDisassociateAnalyticsDataSet {
			connect_BatchDisassociateAnalyticsDataSet(cfg, client)
			return
		}
		if _connectBatchGetAttachedFileMetadata {
			connect_BatchGetAttachedFileMetadata(cfg, client)
			return
		}
		if _connectBatchGetFlowAssociation {
			connect_BatchGetFlowAssociation(cfg, client)
			return
		}
		if _connectBatchPutContact {
			connect_BatchPutContact(cfg, client)
			return
		}
		if _connectBatchUpdateDataTableValue {
			connect_BatchUpdateDataTableValue(cfg, client)
			return
		}
		if _connectClaimPhoneNumber {
			connect_ClaimPhoneNumber(cfg, client)
			return
		}
		if _connectCompleteAttachedFileUpload {
			connect_CompleteAttachedFileUpload(cfg, client)
			return
		}
		if _connectCreateAgentStatus {
			connect_CreateAgentStatus(cfg, client)
			return
		}
		if _connectCreateContact {
			connect_CreateContact(cfg, client)
			return
		}
		if _connectCreateContactFlow {
			connect_CreateContactFlow(cfg, client)
			return
		}
		if _connectCreateContactFlowModule {
			connect_CreateContactFlowModule(cfg, client)
			return
		}
		if _connectCreateContactFlowModuleAlias {
			connect_CreateContactFlowModuleAlias(cfg, client)
			return
		}
		if _connectCreateContactFlowModuleVersion {
			connect_CreateContactFlowModuleVersion(cfg, client)
			return
		}
		if _connectCreateContactFlowVersion {
			connect_CreateContactFlowVersion(cfg, client)
			return
		}
		if _connectCreateDataTable {
			connect_CreateDataTable(cfg, client)
			return
		}
		if _connectCreateDataTableAttribute {
			connect_CreateDataTableAttribute(cfg, client)
			return
		}
		if _connectCreateEmailAddress {
			connect_CreateEmailAddress(cfg, client)
			return
		}
		if _connectCreateEvaluationForm {
			connect_CreateEvaluationForm(cfg, client)
			return
		}
		if _connectCreateHoursOfOperation {
			connect_CreateHoursOfOperation(cfg, client)
			return
		}
		if _connectCreateHoursOfOperationOverride {
			connect_CreateHoursOfOperationOverride(cfg, client)
			return
		}
		if _connectCreateInstance {
			connect_CreateInstance(cfg, client)
			return
		}
		if _connectCreateIntegrationAssociation {
			connect_CreateIntegrationAssociation(cfg, client)
			return
		}
		if _connectCreateNotification {
			connect_CreateNotification(cfg, client)
			return
		}
		if _connectCreateParticipant {
			connect_CreateParticipant(cfg, client)
			return
		}
		if _connectCreatePersistentContactAssociation {
			connect_CreatePersistentContactAssociation(cfg, client)
			return
		}
		if _connectCreatePredefinedAttribute {
			connect_CreatePredefinedAttribute(cfg, client)
			return
		}
		if _connectCreatePrompt {
			connect_CreatePrompt(cfg, client)
			return
		}
		if _connectCreatePushNotificationRegistration {
			connect_CreatePushNotificationRegistration(cfg, client)
			return
		}
		if _connectCreateQueue {
			connect_CreateQueue(cfg, client)
			return
		}
		if _connectCreateQuickConnect {
			connect_CreateQuickConnect(cfg, client)
			return
		}
		if _connectCreateRoutingProfile {
			connect_CreateRoutingProfile(cfg, client)
			return
		}
		if _connectCreateRule {
			connect_CreateRule(cfg, client)
			return
		}
		if _connectCreateSecurityProfile {
			connect_CreateSecurityProfile(cfg, client)
			return
		}
		if _connectCreateTaskTemplate {
			connect_CreateTaskTemplate(cfg, client)
			return
		}
		if _connectCreateTestCase {
			connect_CreateTestCase(cfg, client)
			return
		}
		if _connectCreateTrafficDistributionGroup {
			connect_CreateTrafficDistributionGroup(cfg, client)
			return
		}
		if _connectCreateUseCase {
			connect_CreateUseCase(cfg, client)
			return
		}
		if _connectCreateUser {
			connect_CreateUser(cfg, client)
			return
		}
		if _connectCreateUserHierarchyGroup {
			connect_CreateUserHierarchyGroup(cfg, client)
			return
		}
		if _connectCreateView {
			connect_CreateView(cfg, client)
			return
		}
		if _connectCreateViewVersion {
			connect_CreateViewVersion(cfg, client)
			return
		}
		if _connectCreateVocabulary {
			connect_CreateVocabulary(cfg, client)
			return
		}
		if _connectCreateWorkspace {
			connect_CreateWorkspace(cfg, client)
			return
		}
		if _connectCreateWorkspacePage {
			connect_CreateWorkspacePage(cfg, client)
			return
		}
		if _connectDeactivateEvaluationForm {
			connect_DeactivateEvaluationForm(cfg, client)
			return
		}
		if _connectDeleteAttachedFile {
			connect_DeleteAttachedFile(cfg, client)
			return
		}
		if _connectDeleteContactEvaluation {
			connect_DeleteContactEvaluation(cfg, client)
			return
		}
		if _connectDeleteContactFlow {
			connect_DeleteContactFlow(cfg, client)
			return
		}
		if _connectDeleteContactFlowModule {
			connect_DeleteContactFlowModule(cfg, client)
			return
		}
		if _connectDeleteContactFlowModuleAlias {
			connect_DeleteContactFlowModuleAlias(cfg, client)
			return
		}
		if _connectDeleteContactFlowModuleVersion {
			connect_DeleteContactFlowModuleVersion(cfg, client)
			return
		}
		if _connectDeleteContactFlowVersion {
			connect_DeleteContactFlowVersion(cfg, client)
			return
		}
		if _connectDeleteDataTable {
			connect_DeleteDataTable(cfg, client)
			return
		}
		if _connectDeleteDataTableAttribute {
			connect_DeleteDataTableAttribute(cfg, client)
			return
		}
		if _connectDeleteEmailAddress {
			connect_DeleteEmailAddress(cfg, client)
			return
		}
		if _connectDeleteEvaluationForm {
			connect_DeleteEvaluationForm(cfg, client)
			return
		}
		if _connectDeleteHoursOfOperation {
			connect_DeleteHoursOfOperation(cfg, client)
			return
		}
		if _connectDeleteHoursOfOperationOverride {
			connect_DeleteHoursOfOperationOverride(cfg, client)
			return
		}
		if _connectDeleteInstance {
			connect_DeleteInstance(cfg, client)
			return
		}
		if _connectDeleteIntegrationAssociation {
			connect_DeleteIntegrationAssociation(cfg, client)
			return
		}
		if _connectDeleteNotification {
			connect_DeleteNotification(cfg, client)
			return
		}
		if _connectDeletePredefinedAttribute {
			connect_DeletePredefinedAttribute(cfg, client)
			return
		}
		if _connectDeletePrompt {
			connect_DeletePrompt(cfg, client)
			return
		}
		if _connectDeletePushNotificationRegistration {
			connect_DeletePushNotificationRegistration(cfg, client)
			return
		}
		if _connectDeleteQueue {
			connect_DeleteQueue(cfg, client)
			return
		}
		if _connectDeleteQuickConnect {
			connect_DeleteQuickConnect(cfg, client)
			return
		}
		if _connectDeleteRoutingProfile {
			connect_DeleteRoutingProfile(cfg, client)
			return
		}
		if _connectDeleteRule {
			connect_DeleteRule(cfg, client)
			return
		}
		if _connectDeleteSecurityProfile {
			connect_DeleteSecurityProfile(cfg, client)
			return
		}
		if _connectDeleteTaskTemplate {
			connect_DeleteTaskTemplate(cfg, client)
			return
		}
		if _connectDeleteTestCase {
			connect_DeleteTestCase(cfg, client)
			return
		}
		if _connectDeleteTrafficDistributionGroup {
			connect_DeleteTrafficDistributionGroup(cfg, client)
			return
		}
		if _connectDeleteUseCase {
			connect_DeleteUseCase(cfg, client)
			return
		}
		if _connectDeleteUser {
			connect_DeleteUser(cfg, client)
			return
		}
		if _connectDeleteUserHierarchyGroup {
			connect_DeleteUserHierarchyGroup(cfg, client)
			return
		}
		if _connectDeleteView {
			connect_DeleteView(cfg, client)
			return
		}
		if _connectDeleteViewVersion {
			connect_DeleteViewVersion(cfg, client)
			return
		}
		if _connectDeleteVocabulary {
			connect_DeleteVocabulary(cfg, client)
			return
		}
		if _connectDeleteWorkspace {
			connect_DeleteWorkspace(cfg, client)
			return
		}
		if _connectDeleteWorkspaceMedia {
			connect_DeleteWorkspaceMedia(cfg, client)
			return
		}
		if _connectDeleteWorkspacePage {
			connect_DeleteWorkspacePage(cfg, client)
			return
		}
		if _connectDescribeAgentStatus {
			connect_DescribeAgentStatus(cfg, client)
			return
		}
		if _connectDescribeAuthenticationProfile {
			connect_DescribeAuthenticationProfile(cfg, client)
			return
		}
		if _connectDescribeContact {
			connect_DescribeContact(cfg, client)
			return
		}
		if _connectDescribeContactEvaluation {
			connect_DescribeContactEvaluation(cfg, client)
			return
		}
		if _connectDescribeContactFlow {
			connect_DescribeContactFlow(cfg, client)
			return
		}
		if _connectDescribeContactFlowModule {
			connect_DescribeContactFlowModule(cfg, client)
			return
		}
		if _connectDescribeContactFlowModuleAlias {
			connect_DescribeContactFlowModuleAlias(cfg, client)
			return
		}
		if _connectDescribeDataTable {
			connect_DescribeDataTable(cfg, client)
			return
		}
		if _connectDescribeDataTableAttribute {
			connect_DescribeDataTableAttribute(cfg, client)
			return
		}
		if _connectDescribeEmailAddress {
			connect_DescribeEmailAddress(cfg, client)
			return
		}
		if _connectDescribeEvaluationForm {
			connect_DescribeEvaluationForm(cfg, client)
			return
		}
		if _connectDescribeHoursOfOperation {
			connect_DescribeHoursOfOperation(cfg, client)
			return
		}
		if _connectDescribeHoursOfOperationOverride {
			connect_DescribeHoursOfOperationOverride(cfg, client)
			return
		}
		if _connectDescribeInstance {
			connect_DescribeInstance(cfg, client)
			return
		}
		if _connectDescribeInstanceAttribute {
			connect_DescribeInstanceAttribute(cfg, client)
			return
		}
		if _connectDescribeInstanceStorageConfig {
			connect_DescribeInstanceStorageConfig(cfg, client)
			return
		}
		if _connectDescribeNotification {
			connect_DescribeNotification(cfg, client)
			return
		}
		if _connectDescribePhoneNumber {
			connect_DescribePhoneNumber(cfg, client)
			return
		}
		if _connectDescribePredefinedAttribute {
			connect_DescribePredefinedAttribute(cfg, client)
			return
		}
		if _connectDescribePrompt {
			connect_DescribePrompt(cfg, client)
			return
		}
		if _connectDescribeQueue {
			connect_DescribeQueue(cfg, client)
			return
		}
		if _connectDescribeQuickConnect {
			connect_DescribeQuickConnect(cfg, client)
			return
		}
		if _connectDescribeRoutingProfile {
			connect_DescribeRoutingProfile(cfg, client)
			return
		}
		if _connectDescribeRule {
			connect_DescribeRule(cfg, client)
			return
		}
		if _connectDescribeSecurityProfile {
			connect_DescribeSecurityProfile(cfg, client)
			return
		}
		if _connectDescribeTestCase {
			connect_DescribeTestCase(cfg, client)
			return
		}
		if _connectDescribeTrafficDistributionGroup {
			connect_DescribeTrafficDistributionGroup(cfg, client)
			return
		}
		if _connectDescribeUser {
			connect_DescribeUser(cfg, client)
			return
		}
		if _connectDescribeUserHierarchyGroup {
			connect_DescribeUserHierarchyGroup(cfg, client)
			return
		}
		if _connectDescribeUserHierarchyStructure {
			connect_DescribeUserHierarchyStructure(cfg, client)
			return
		}
		if _connectDescribeView {
			connect_DescribeView(cfg, client)
			return
		}
		if _connectDescribeVocabulary {
			connect_DescribeVocabulary(cfg, client)
			return
		}
		if _connectDescribeWorkspace {
			connect_DescribeWorkspace(cfg, client)
			return
		}
		if _connectDisassociateAnalyticsDataSet {
			connect_DisassociateAnalyticsDataSet(cfg, client)
			return
		}
		if _connectDisassociateApprovedOrigin {
			connect_DisassociateApprovedOrigin(cfg, client)
			return
		}
		if _connectDisassociateBot {
			connect_DisassociateBot(cfg, client)
			return
		}
		if _connectDisassociateEmailAddressAlias {
			connect_DisassociateEmailAddressAlias(cfg, client)
			return
		}
		if _connectDisassociateFlow {
			connect_DisassociateFlow(cfg, client)
			return
		}
		if _connectDisassociateHoursOfOperations {
			connect_DisassociateHoursOfOperations(cfg, client)
			return
		}
		if _connectDisassociateInstanceStorageConfig {
			connect_DisassociateInstanceStorageConfig(cfg, client)
			return
		}
		if _connectDisassociateLambdaFunction {
			connect_DisassociateLambdaFunction(cfg, client)
			return
		}
		if _connectDisassociateLexBot {
			connect_DisassociateLexBot(cfg, client)
			return
		}
		if _connectDisassociatePhoneNumberContactFlow {
			connect_DisassociatePhoneNumberContactFlow(cfg, client)
			return
		}
		if _connectDisassociateQueueQuickConnects {
			connect_DisassociateQueueQuickConnects(cfg, client)
			return
		}
		if _connectDisassociateRoutingProfileQueues {
			connect_DisassociateRoutingProfileQueues(cfg, client)
			return
		}
		if _connectDisassociateSecurityKey {
			connect_DisassociateSecurityKey(cfg, client)
			return
		}
		if _connectDisassociateSecurityProfiles {
			connect_DisassociateSecurityProfiles(cfg, client)
			return
		}
		if _connectDisassociateTrafficDistributionGroupUser {
			connect_DisassociateTrafficDistributionGroupUser(cfg, client)
			return
		}
		if _connectDisassociateUserProficiencies {
			connect_DisassociateUserProficiencies(cfg, client)
			return
		}
		if _connectDisassociateWorkspace {
			connect_DisassociateWorkspace(cfg, client)
			return
		}
		if _connectDismissUserContact {
			connect_DismissUserContact(cfg, client)
			return
		}
		if _connectEvaluateDataTableValues {
			connect_EvaluateDataTableValues(cfg, client)
			return
		}
		if _connectGetAttachedFile {
			connect_GetAttachedFile(cfg, client)
			return
		}
		if _connectGetContactAttributes {
			connect_GetContactAttributes(cfg, client)
			return
		}
		if _connectGetContactMetrics {
			connect_GetContactMetrics(cfg, client)
			return
		}
		if _connectGetCurrentMetricData {
			connect_GetCurrentMetricData(cfg, client)
			return
		}
		if _connectGetCurrentUserData {
			connect_GetCurrentUserData(cfg, client)
			return
		}
		if _connectGetEffectiveHoursOfOperations {
			connect_GetEffectiveHoursOfOperations(cfg, client)
			return
		}
		if _connectGetFederationToken {
			connect_GetFederationToken(cfg, client)
			return
		}
		if _connectGetFlowAssociation {
			connect_GetFlowAssociation(cfg, client)
			return
		}
		if _connectGetMetricData {
			connect_GetMetricData(cfg, client)
			return
		}
		if _connectGetMetricDataV2 {
			connect_GetMetricDataV2(cfg, client)
			return
		}
		if _connectGetPromptFile {
			connect_GetPromptFile(cfg, client)
			return
		}
		if _connectGetTaskTemplate {
			connect_GetTaskTemplate(cfg, client)
			return
		}
		if _connectGetTestCaseExecutionSummary {
			connect_GetTestCaseExecutionSummary(cfg, client)
			return
		}
		if _connectGetTrafficDistribution {
			connect_GetTrafficDistribution(cfg, client)
			return
		}
		if _connectImportPhoneNumber {
			connect_ImportPhoneNumber(cfg, client)
			return
		}
		if _connectImportWorkspaceMedia {
			connect_ImportWorkspaceMedia(cfg, client)
			return
		}
		if _connectListAgentStatuses {
			connect_ListAgentStatuses(cfg, client)
			return
		}
		if _connectListAnalyticsDataAssociations {
			connect_ListAnalyticsDataAssociations(cfg, client)
			return
		}
		if _connectListAnalyticsDataLakeDataSets {
			connect_ListAnalyticsDataLakeDataSets(cfg, client)
			return
		}
		if _connectListApprovedOrigins {
			connect_ListApprovedOrigins(cfg, client)
			return
		}
		if _connectListAssociatedContacts {
			connect_ListAssociatedContacts(cfg, client)
			return
		}
		if _connectListAuthenticationProfiles {
			connect_ListAuthenticationProfiles(cfg, client)
			return
		}
		if _connectListBots {
			connect_ListBots(cfg, client)
			return
		}
		if _connectListChildHoursOfOperations {
			connect_ListChildHoursOfOperations(cfg, client)
			return
		}
		if _connectListContactEvaluations {
			connect_ListContactEvaluations(cfg, client)
			return
		}
		if _connectListContactFlowModuleAliases {
			connect_ListContactFlowModuleAliases(cfg, client)
			return
		}
		if _connectListContactFlowModuleVersions {
			connect_ListContactFlowModuleVersions(cfg, client)
			return
		}
		if _connectListContactFlowModules {
			connect_ListContactFlowModules(cfg, client)
			return
		}
		if _connectListContactFlowVersions {
			connect_ListContactFlowVersions(cfg, client)
			return
		}
		if _connectListContactFlows {
			connect_ListContactFlows(cfg, client)
			return
		}
		if _connectListContactReferences {
			connect_ListContactReferences(cfg, client)
			return
		}
		if _connectListDataTableAttributes {
			connect_ListDataTableAttributes(cfg, client)
			return
		}
		if _connectListDataTablePrimaryValues {
			connect_ListDataTablePrimaryValues(cfg, client)
			return
		}
		if _connectListDataTableValues {
			connect_ListDataTableValues(cfg, client)
			return
		}
		if _connectListDataTables {
			connect_ListDataTables(cfg, client)
			return
		}
		if _connectListDefaultVocabularies {
			connect_ListDefaultVocabularies(cfg, client)
			return
		}
		if _connectListEntitySecurityProfiles {
			connect_ListEntitySecurityProfiles(cfg, client)
			return
		}
		if _connectListEvaluationFormVersions {
			connect_ListEvaluationFormVersions(cfg, client)
			return
		}
		if _connectListEvaluationForms {
			connect_ListEvaluationForms(cfg, client)
			return
		}
		if _connectListFlowAssociations {
			connect_ListFlowAssociations(cfg, client)
			return
		}
		if _connectListHoursOfOperationOverrides {
			connect_ListHoursOfOperationOverrides(cfg, client)
			return
		}
		if _connectListHoursOfOperations {
			connect_ListHoursOfOperations(cfg, client)
			return
		}
		if _connectListInstanceAttributes {
			connect_ListInstanceAttributes(cfg, client)
			return
		}
		if _connectListInstanceStorageConfigs {
			connect_ListInstanceStorageConfigs(cfg, client)
			return
		}
		if _connectListInstances {
			connect_ListInstances(cfg, client)
			return
		}
		if _connectListIntegrationAssociations {
			connect_ListIntegrationAssociations(cfg, client)
			return
		}
		if _connectListLambdaFunctions {
			connect_ListLambdaFunctions(cfg, client)
			return
		}
		if _connectListLexBots {
			connect_ListLexBots(cfg, client)
			return
		}
		if _connectListNotifications {
			connect_ListNotifications(cfg, client)
			return
		}
		if _connectListPhoneNumbers {
			connect_ListPhoneNumbers(cfg, client)
			return
		}
		if _connectListPhoneNumbersV2 {
			connect_ListPhoneNumbersV2(cfg, client)
			return
		}
		if _connectListPredefinedAttributes {
			connect_ListPredefinedAttributes(cfg, client)
			return
		}
		if _connectListPrompts {
			connect_ListPrompts(cfg, client)
			return
		}
		if _connectListQueueQuickConnects {
			connect_ListQueueQuickConnects(cfg, client)
			return
		}
		if _connectListQueues {
			connect_ListQueues(cfg, client)
			return
		}
		if _connectListQuickConnects {
			connect_ListQuickConnects(cfg, client)
			return
		}
		if _connectListRealtimeContactAnalysisSegmentsV2 {
			connect_ListRealtimeContactAnalysisSegmentsV2(cfg, client)
			return
		}
		if _connectListRoutingProfileManualAssignmentQueues {
			connect_ListRoutingProfileManualAssignmentQueues(cfg, client)
			return
		}
		if _connectListRoutingProfileQueues {
			connect_ListRoutingProfileQueues(cfg, client)
			return
		}
		if _connectListRoutingProfiles {
			connect_ListRoutingProfiles(cfg, client)
			return
		}
		if _connectListRules {
			connect_ListRules(cfg, client)
			return
		}
		if _connectListSecurityKeys {
			connect_ListSecurityKeys(cfg, client)
			return
		}
		if _connectListSecurityProfileApplications {
			connect_ListSecurityProfileApplications(cfg, client)
			return
		}
		if _connectListSecurityProfileFlowModules {
			connect_ListSecurityProfileFlowModules(cfg, client)
			return
		}
		if _connectListSecurityProfilePermissions {
			connect_ListSecurityProfilePermissions(cfg, client)
			return
		}
		if _connectListSecurityProfiles {
			connect_ListSecurityProfiles(cfg, client)
			return
		}
		if _connectListTagsForResource {
			connect_ListTagsForResource(cfg, client)
			return
		}
		if _connectListTaskTemplates {
			connect_ListTaskTemplates(cfg, client)
			return
		}
		if _connectListTestCaseExecutionRecords {
			connect_ListTestCaseExecutionRecords(cfg, client)
			return
		}
		if _connectListTestCaseExecutions {
			connect_ListTestCaseExecutions(cfg, client)
			return
		}
		if _connectListTestCases {
			connect_ListTestCases(cfg, client)
			return
		}
		if _connectListTrafficDistributionGroupUsers {
			connect_ListTrafficDistributionGroupUsers(cfg, client)
			return
		}
		if _connectListTrafficDistributionGroups {
			connect_ListTrafficDistributionGroups(cfg, client)
			return
		}
		if _connectListUseCases {
			connect_ListUseCases(cfg, client)
			return
		}
		if _connectListUserHierarchyGroups {
			connect_ListUserHierarchyGroups(cfg, client)
			return
		}
		if _connectListUserNotifications {
			connect_ListUserNotifications(cfg, client)
			return
		}
		if _connectListUserProficiencies {
			connect_ListUserProficiencies(cfg, client)
			return
		}
		if _connectListUsers {
			connect_ListUsers(cfg, client)
			return
		}
		if _connectListViewVersions {
			connect_ListViewVersions(cfg, client)
			return
		}
		if _connectListViews {
			connect_ListViews(cfg, client)
			return
		}
		if _connectListWorkspaceMedia {
			connect_ListWorkspaceMedia(cfg, client)
			return
		}
		if _connectListWorkspacePages {
			connect_ListWorkspacePages(cfg, client)
			return
		}
		if _connectListWorkspaces {
			connect_ListWorkspaces(cfg, client)
			return
		}
		if _connectMonitorContact {
			connect_MonitorContact(cfg, client)
			return
		}
		if _connectPauseContact {
			connect_PauseContact(cfg, client)
			return
		}
		if _connectPutUserStatus {
			connect_PutUserStatus(cfg, client)
			return
		}
		if _connectReleasePhoneNumber {
			connect_ReleasePhoneNumber(cfg, client)
			return
		}
		if _connectReplicateInstance {
			connect_ReplicateInstance(cfg, client)
			return
		}
		if _connectResumeContact {
			connect_ResumeContact(cfg, client)
			return
		}
		if _connectResumeContactRecording {
			connect_ResumeContactRecording(cfg, client)
			return
		}
		if _connectSearchAgentStatuses {
			connect_SearchAgentStatuses(cfg, client)
			return
		}
		if _connectSearchAvailablePhoneNumbers {
			connect_SearchAvailablePhoneNumbers(cfg, client)
			return
		}
		if _connectSearchContactEvaluations {
			connect_SearchContactEvaluations(cfg, client)
			return
		}
		if _connectSearchContactFlowModules {
			connect_SearchContactFlowModules(cfg, client)
			return
		}
		if _connectSearchContactFlows {
			connect_SearchContactFlows(cfg, client)
			return
		}
		if _connectSearchContacts {
			connect_SearchContacts(cfg, client)
			return
		}
		if _connectSearchDataTables {
			connect_SearchDataTables(cfg, client)
			return
		}
		if _connectSearchEmailAddresses {
			connect_SearchEmailAddresses(cfg, client)
			return
		}
		if _connectSearchEvaluationForms {
			connect_SearchEvaluationForms(cfg, client)
			return
		}
		if _connectSearchHoursOfOperationOverrides {
			connect_SearchHoursOfOperationOverrides(cfg, client)
			return
		}
		if _connectSearchHoursOfOperations {
			connect_SearchHoursOfOperations(cfg, client)
			return
		}
		if _connectSearchNotifications {
			connect_SearchNotifications(cfg, client)
			return
		}
		if _connectSearchPredefinedAttributes {
			connect_SearchPredefinedAttributes(cfg, client)
			return
		}
		if _connectSearchPrompts {
			connect_SearchPrompts(cfg, client)
			return
		}
		if _connectSearchQueues {
			connect_SearchQueues(cfg, client)
			return
		}
		if _connectSearchQuickConnects {
			connect_SearchQuickConnects(cfg, client)
			return
		}
		if _connectSearchResourceTags {
			connect_SearchResourceTags(cfg, client)
			return
		}
		if _connectSearchRoutingProfiles {
			connect_SearchRoutingProfiles(cfg, client)
			return
		}
		if _connectSearchSecurityProfiles {
			connect_SearchSecurityProfiles(cfg, client)
			return
		}
		if _connectSearchTestCases {
			connect_SearchTestCases(cfg, client)
			return
		}
		if _connectSearchUserHierarchyGroups {
			connect_SearchUserHierarchyGroups(cfg, client)
			return
		}
		if _connectSearchUsers {
			connect_SearchUsers(cfg, client)
			return
		}
		if _connectSearchViews {
			connect_SearchViews(cfg, client)
			return
		}
		if _connectSearchVocabularies {
			connect_SearchVocabularies(cfg, client)
			return
		}
		if _connectSearchWorkspaceAssociations {
			connect_SearchWorkspaceAssociations(cfg, client)
			return
		}
		if _connectSearchWorkspaces {
			connect_SearchWorkspaces(cfg, client)
			return
		}
		if _connectSendChatIntegrationEvent {
			connect_SendChatIntegrationEvent(cfg, client)
			return
		}
		if _connectSendOutboundEmail {
			connect_SendOutboundEmail(cfg, client)
			return
		}
		if _connectStartAttachedFileUpload {
			connect_StartAttachedFileUpload(cfg, client)
			return
		}
		if _connectStartChatContact {
			connect_StartChatContact(cfg, client)
			return
		}
		if _connectStartContactEvaluation {
			connect_StartContactEvaluation(cfg, client)
			return
		}
		if _connectStartContactMediaProcessing {
			connect_StartContactMediaProcessing(cfg, client)
			return
		}
		if _connectStartContactRecording {
			connect_StartContactRecording(cfg, client)
			return
		}
		if _connectStartContactStreaming {
			connect_StartContactStreaming(cfg, client)
			return
		}
		if _connectStartEmailContact {
			connect_StartEmailContact(cfg, client)
			return
		}
		if _connectStartOutboundChatContact {
			connect_StartOutboundChatContact(cfg, client)
			return
		}
		if _connectStartOutboundEmailContact {
			connect_StartOutboundEmailContact(cfg, client)
			return
		}
		if _connectStartOutboundVoiceContact {
			connect_StartOutboundVoiceContact(cfg, client)
			return
		}
		if _connectStartScreenSharing {
			connect_StartScreenSharing(cfg, client)
			return
		}
		if _connectStartTaskContact {
			connect_StartTaskContact(cfg, client)
			return
		}
		if _connectStartTestCaseExecution {
			connect_StartTestCaseExecution(cfg, client)
			return
		}
		if _connectStartWebRTCContact {
			connect_StartWebRTCContact(cfg, client)
			return
		}
		if _connectStopContact {
			connect_StopContact(cfg, client)
			return
		}
		if _connectStopContactMediaProcessing {
			connect_StopContactMediaProcessing(cfg, client)
			return
		}
		if _connectStopContactRecording {
			connect_StopContactRecording(cfg, client)
			return
		}
		if _connectStopContactStreaming {
			connect_StopContactStreaming(cfg, client)
			return
		}
		if _connectStopTestCaseExecution {
			connect_StopTestCaseExecution(cfg, client)
			return
		}
		if _connectSubmitContactEvaluation {
			connect_SubmitContactEvaluation(cfg, client)
			return
		}
		if _connectSuspendContactRecording {
			connect_SuspendContactRecording(cfg, client)
			return
		}
		if _connectTagContact {
			connect_TagContact(cfg, client)
			return
		}
		if _connectTagResource {
			connect_TagResource(cfg, client)
			return
		}
		if _connectTransferContact {
			connect_TransferContact(cfg, client)
			return
		}
		if _connectUntagContact {
			connect_UntagContact(cfg, client)
			return
		}
		if _connectUntagResource {
			connect_UntagResource(cfg, client)
			return
		}
		if _connectUpdateAgentStatus {
			connect_UpdateAgentStatus(cfg, client)
			return
		}
		if _connectUpdateAuthenticationProfile {
			connect_UpdateAuthenticationProfile(cfg, client)
			return
		}
		if _connectUpdateContact {
			connect_UpdateContact(cfg, client)
			return
		}
		if _connectUpdateContactAttributes {
			connect_UpdateContactAttributes(cfg, client)
			return
		}
		if _connectUpdateContactEvaluation {
			connect_UpdateContactEvaluation(cfg, client)
			return
		}
		if _connectUpdateContactFlowContent {
			connect_UpdateContactFlowContent(cfg, client)
			return
		}
		if _connectUpdateContactFlowMetadata {
			connect_UpdateContactFlowMetadata(cfg, client)
			return
		}
		if _connectUpdateContactFlowModuleAlias {
			connect_UpdateContactFlowModuleAlias(cfg, client)
			return
		}
		if _connectUpdateContactFlowModuleContent {
			connect_UpdateContactFlowModuleContent(cfg, client)
			return
		}
		if _connectUpdateContactFlowModuleMetadata {
			connect_UpdateContactFlowModuleMetadata(cfg, client)
			return
		}
		if _connectUpdateContactFlowName {
			connect_UpdateContactFlowName(cfg, client)
			return
		}
		if _connectUpdateContactRoutingData {
			connect_UpdateContactRoutingData(cfg, client)
			return
		}
		if _connectUpdateContactSchedule {
			connect_UpdateContactSchedule(cfg, client)
			return
		}
		if _connectUpdateDataTableAttribute {
			connect_UpdateDataTableAttribute(cfg, client)
			return
		}
		if _connectUpdateDataTableMetadata {
			connect_UpdateDataTableMetadata(cfg, client)
			return
		}
		if _connectUpdateDataTablePrimaryValues {
			connect_UpdateDataTablePrimaryValues(cfg, client)
			return
		}
		if _connectUpdateEmailAddressMetadata {
			connect_UpdateEmailAddressMetadata(cfg, client)
			return
		}
		if _connectUpdateEvaluationForm {
			connect_UpdateEvaluationForm(cfg, client)
			return
		}
		if _connectUpdateHoursOfOperation {
			connect_UpdateHoursOfOperation(cfg, client)
			return
		}
		if _connectUpdateHoursOfOperationOverride {
			connect_UpdateHoursOfOperationOverride(cfg, client)
			return
		}
		if _connectUpdateInstanceAttribute {
			connect_UpdateInstanceAttribute(cfg, client)
			return
		}
		if _connectUpdateInstanceStorageConfig {
			connect_UpdateInstanceStorageConfig(cfg, client)
			return
		}
		if _connectUpdateNotificationContent {
			connect_UpdateNotificationContent(cfg, client)
			return
		}
		if _connectUpdateParticipantAuthentication {
			connect_UpdateParticipantAuthentication(cfg, client)
			return
		}
		if _connectUpdateParticipantRoleConfig {
			connect_UpdateParticipantRoleConfig(cfg, client)
			return
		}
		if _connectUpdatePhoneNumber {
			connect_UpdatePhoneNumber(cfg, client)
			return
		}
		if _connectUpdatePhoneNumberMetadata {
			connect_UpdatePhoneNumberMetadata(cfg, client)
			return
		}
		if _connectUpdatePredefinedAttribute {
			connect_UpdatePredefinedAttribute(cfg, client)
			return
		}
		if _connectUpdatePrompt {
			connect_UpdatePrompt(cfg, client)
			return
		}
		if _connectUpdateQueueHoursOfOperation {
			connect_UpdateQueueHoursOfOperation(cfg, client)
			return
		}
		if _connectUpdateQueueMaxContacts {
			connect_UpdateQueueMaxContacts(cfg, client)
			return
		}
		if _connectUpdateQueueName {
			connect_UpdateQueueName(cfg, client)
			return
		}
		if _connectUpdateQueueOutboundCallerConfig {
			connect_UpdateQueueOutboundCallerConfig(cfg, client)
			return
		}
		if _connectUpdateQueueOutboundEmailConfig {
			connect_UpdateQueueOutboundEmailConfig(cfg, client)
			return
		}
		if _connectUpdateQueueStatus {
			connect_UpdateQueueStatus(cfg, client)
			return
		}
		if _connectUpdateQuickConnectConfig {
			connect_UpdateQuickConnectConfig(cfg, client)
			return
		}
		if _connectUpdateQuickConnectName {
			connect_UpdateQuickConnectName(cfg, client)
			return
		}
		if _connectUpdateRoutingProfileAgentAvailabilityTimer {
			connect_UpdateRoutingProfileAgentAvailabilityTimer(cfg, client)
			return
		}
		if _connectUpdateRoutingProfileConcurrency {
			connect_UpdateRoutingProfileConcurrency(cfg, client)
			return
		}
		if _connectUpdateRoutingProfileDefaultOutboundQueue {
			connect_UpdateRoutingProfileDefaultOutboundQueue(cfg, client)
			return
		}
		if _connectUpdateRoutingProfileName {
			connect_UpdateRoutingProfileName(cfg, client)
			return
		}
		if _connectUpdateRoutingProfileQueues {
			connect_UpdateRoutingProfileQueues(cfg, client)
			return
		}
		if _connectUpdateRule {
			connect_UpdateRule(cfg, client)
			return
		}
		if _connectUpdateSecurityProfile {
			connect_UpdateSecurityProfile(cfg, client)
			return
		}
		if _connectUpdateTaskTemplate {
			connect_UpdateTaskTemplate(cfg, client)
			return
		}
		if _connectUpdateTestCase {
			connect_UpdateTestCase(cfg, client)
			return
		}
		if _connectUpdateTrafficDistribution {
			connect_UpdateTrafficDistribution(cfg, client)
			return
		}
		if _connectUpdateUserConfig {
			connect_UpdateUserConfig(cfg, client)
			return
		}
		if _connectUpdateUserHierarchy {
			connect_UpdateUserHierarchy(cfg, client)
			return
		}
		if _connectUpdateUserHierarchyGroupName {
			connect_UpdateUserHierarchyGroupName(cfg, client)
			return
		}
		if _connectUpdateUserHierarchyStructure {
			connect_UpdateUserHierarchyStructure(cfg, client)
			return
		}
		if _connectUpdateUserIdentityInfo {
			connect_UpdateUserIdentityInfo(cfg, client)
			return
		}
		if _connectUpdateUserNotificationStatus {
			connect_UpdateUserNotificationStatus(cfg, client)
			return
		}
		if _connectUpdateUserPhoneConfig {
			connect_UpdateUserPhoneConfig(cfg, client)
			return
		}
		if _connectUpdateUserProficiencies {
			connect_UpdateUserProficiencies(cfg, client)
			return
		}
		if _connectUpdateUserRoutingProfile {
			connect_UpdateUserRoutingProfile(cfg, client)
			return
		}
		if _connectUpdateUserSecurityProfiles {
			connect_UpdateUserSecurityProfiles(cfg, client)
			return
		}
		if _connectUpdateViewContent {
			connect_UpdateViewContent(cfg, client)
			return
		}
		if _connectUpdateViewMetadata {
			connect_UpdateViewMetadata(cfg, client)
			return
		}
		if _connectUpdateWorkspaceMetadata {
			connect_UpdateWorkspaceMetadata(cfg, client)
			return
		}
		if _connectUpdateWorkspacePage {
			connect_UpdateWorkspacePage(cfg, client)
			return
		}
		if _connectUpdateWorkspaceTheme {
			connect_UpdateWorkspaceTheme(cfg, client)
			return
		}
		if _connectUpdateWorkspaceVisibility {
			connect_UpdateWorkspaceVisibility(cfg, client)
			return
		}

	},
}

var (
	_connectActivateEvaluationForm                     bool
	_connectAssociateAnalyticsDataSet                  bool
	_connectAssociateApprovedOrigin                    bool
	_connectAssociateBot                               bool
	_connectAssociateContactWithUser                   bool
	_connectAssociateDefaultVocabulary                 bool
	_connectAssociateEmailAddressAlias                 bool
	_connectAssociateFlow                              bool
	_connectAssociateHoursOfOperations                 bool
	_connectAssociateInstanceStorageConfig             bool
	_connectAssociateLambdaFunction                    bool
	_connectAssociateLexBot                            bool
	_connectAssociatePhoneNumberContactFlow            bool
	_connectAssociateQueueQuickConnects                bool
	_connectAssociateRoutingProfileQueues              bool
	_connectAssociateSecurityKey                       bool
	_connectAssociateSecurityProfiles                  bool
	_connectAssociateTrafficDistributionGroupUser      bool
	_connectAssociateUserProficiencies                 bool
	_connectAssociateWorkspace                         bool
	_connectBatchAssociateAnalyticsDataSet             bool
	_connectBatchCreateDataTableValue                  bool
	_connectBatchDeleteDataTableValue                  bool
	_connectBatchDescribeDataTableValue                bool
	_connectBatchDisassociateAnalyticsDataSet          bool
	_connectBatchGetAttachedFileMetadata               bool
	_connectBatchGetFlowAssociation                    bool
	_connectBatchPutContact                            bool
	_connectBatchUpdateDataTableValue                  bool
	_connectClaimPhoneNumber                           bool
	_connectCompleteAttachedFileUpload                 bool
	_connectCreateAgentStatus                          bool
	_connectCreateContact                              bool
	_connectCreateContactFlow                          bool
	_connectCreateContactFlowModule                    bool
	_connectCreateContactFlowModuleAlias               bool
	_connectCreateContactFlowModuleVersion             bool
	_connectCreateContactFlowVersion                   bool
	_connectCreateDataTable                            bool
	_connectCreateDataTableAttribute                   bool
	_connectCreateEmailAddress                         bool
	_connectCreateEvaluationForm                       bool
	_connectCreateHoursOfOperation                     bool
	_connectCreateHoursOfOperationOverride             bool
	_connectCreateInstance                             bool
	_connectCreateIntegrationAssociation               bool
	_connectCreateNotification                         bool
	_connectCreateParticipant                          bool
	_connectCreatePersistentContactAssociation         bool
	_connectCreatePredefinedAttribute                  bool
	_connectCreatePrompt                               bool
	_connectCreatePushNotificationRegistration         bool
	_connectCreateQueue                                bool
	_connectCreateQuickConnect                         bool
	_connectCreateRoutingProfile                       bool
	_connectCreateRule                                 bool
	_connectCreateSecurityProfile                      bool
	_connectCreateTaskTemplate                         bool
	_connectCreateTestCase                             bool
	_connectCreateTrafficDistributionGroup             bool
	_connectCreateUseCase                              bool
	_connectCreateUser                                 bool
	_connectCreateUserHierarchyGroup                   bool
	_connectCreateView                                 bool
	_connectCreateViewVersion                          bool
	_connectCreateVocabulary                           bool
	_connectCreateWorkspace                            bool
	_connectCreateWorkspacePage                        bool
	_connectDeactivateEvaluationForm                   bool
	_connectDeleteAttachedFile                         bool
	_connectDeleteContactEvaluation                    bool
	_connectDeleteContactFlow                          bool
	_connectDeleteContactFlowModule                    bool
	_connectDeleteContactFlowModuleAlias               bool
	_connectDeleteContactFlowModuleVersion             bool
	_connectDeleteContactFlowVersion                   bool
	_connectDeleteDataTable                            bool
	_connectDeleteDataTableAttribute                   bool
	_connectDeleteEmailAddress                         bool
	_connectDeleteEvaluationForm                       bool
	_connectDeleteHoursOfOperation                     bool
	_connectDeleteHoursOfOperationOverride             bool
	_connectDeleteInstance                             bool
	_connectDeleteIntegrationAssociation               bool
	_connectDeleteNotification                         bool
	_connectDeletePredefinedAttribute                  bool
	_connectDeletePrompt                               bool
	_connectDeletePushNotificationRegistration         bool
	_connectDeleteQueue                                bool
	_connectDeleteQuickConnect                         bool
	_connectDeleteRoutingProfile                       bool
	_connectDeleteRule                                 bool
	_connectDeleteSecurityProfile                      bool
	_connectDeleteTaskTemplate                         bool
	_connectDeleteTestCase                             bool
	_connectDeleteTrafficDistributionGroup             bool
	_connectDeleteUseCase                              bool
	_connectDeleteUser                                 bool
	_connectDeleteUserHierarchyGroup                   bool
	_connectDeleteView                                 bool
	_connectDeleteViewVersion                          bool
	_connectDeleteVocabulary                           bool
	_connectDeleteWorkspace                            bool
	_connectDeleteWorkspaceMedia                       bool
	_connectDeleteWorkspacePage                        bool
	_connectDescribeAgentStatus                        bool
	_connectDescribeAuthenticationProfile              bool
	_connectDescribeContact                            bool
	_connectDescribeContactEvaluation                  bool
	_connectDescribeContactFlow                        bool
	_connectDescribeContactFlowModule                  bool
	_connectDescribeContactFlowModuleAlias             bool
	_connectDescribeDataTable                          bool
	_connectDescribeDataTableAttribute                 bool
	_connectDescribeEmailAddress                       bool
	_connectDescribeEvaluationForm                     bool
	_connectDescribeHoursOfOperation                   bool
	_connectDescribeHoursOfOperationOverride           bool
	_connectDescribeInstance                           bool
	_connectDescribeInstanceAttribute                  bool
	_connectDescribeInstanceStorageConfig              bool
	_connectDescribeNotification                       bool
	_connectDescribePhoneNumber                        bool
	_connectDescribePredefinedAttribute                bool
	_connectDescribePrompt                             bool
	_connectDescribeQueue                              bool
	_connectDescribeQuickConnect                       bool
	_connectDescribeRoutingProfile                     bool
	_connectDescribeRule                               bool
	_connectDescribeSecurityProfile                    bool
	_connectDescribeTestCase                           bool
	_connectDescribeTrafficDistributionGroup           bool
	_connectDescribeUser                               bool
	_connectDescribeUserHierarchyGroup                 bool
	_connectDescribeUserHierarchyStructure             bool
	_connectDescribeView                               bool
	_connectDescribeVocabulary                         bool
	_connectDescribeWorkspace                          bool
	_connectDisassociateAnalyticsDataSet               bool
	_connectDisassociateApprovedOrigin                 bool
	_connectDisassociateBot                            bool
	_connectDisassociateEmailAddressAlias              bool
	_connectDisassociateFlow                           bool
	_connectDisassociateHoursOfOperations              bool
	_connectDisassociateInstanceStorageConfig          bool
	_connectDisassociateLambdaFunction                 bool
	_connectDisassociateLexBot                         bool
	_connectDisassociatePhoneNumberContactFlow         bool
	_connectDisassociateQueueQuickConnects             bool
	_connectDisassociateRoutingProfileQueues           bool
	_connectDisassociateSecurityKey                    bool
	_connectDisassociateSecurityProfiles               bool
	_connectDisassociateTrafficDistributionGroupUser   bool
	_connectDisassociateUserProficiencies              bool
	_connectDisassociateWorkspace                      bool
	_connectDismissUserContact                         bool
	_connectEvaluateDataTableValues                    bool
	_connectGetAttachedFile                            bool
	_connectGetContactAttributes                       bool
	_connectGetContactMetrics                          bool
	_connectGetCurrentMetricData                       bool
	_connectGetCurrentUserData                         bool
	_connectGetEffectiveHoursOfOperations              bool
	_connectGetFederationToken                         bool
	_connectGetFlowAssociation                         bool
	_connectGetMetricData                              bool
	_connectGetMetricDataV2                            bool
	_connectGetPromptFile                              bool
	_connectGetTaskTemplate                            bool
	_connectGetTestCaseExecutionSummary                bool
	_connectGetTrafficDistribution                     bool
	_connectImportPhoneNumber                          bool
	_connectImportWorkspaceMedia                       bool
	_connectListAgentStatuses                          bool
	_connectListAnalyticsDataAssociations              bool
	_connectListAnalyticsDataLakeDataSets              bool
	_connectListApprovedOrigins                        bool
	_connectListAssociatedContacts                     bool
	_connectListAuthenticationProfiles                 bool
	_connectListBots                                   bool
	_connectListChildHoursOfOperations                 bool
	_connectListContactEvaluations                     bool
	_connectListContactFlowModuleAliases               bool
	_connectListContactFlowModuleVersions              bool
	_connectListContactFlowModules                     bool
	_connectListContactFlowVersions                    bool
	_connectListContactFlows                           bool
	_connectListContactReferences                      bool
	_connectListDataTableAttributes                    bool
	_connectListDataTablePrimaryValues                 bool
	_connectListDataTableValues                        bool
	_connectListDataTables                             bool
	_connectListDefaultVocabularies                    bool
	_connectListEntitySecurityProfiles                 bool
	_connectListEvaluationFormVersions                 bool
	_connectListEvaluationForms                        bool
	_connectListFlowAssociations                       bool
	_connectListHoursOfOperationOverrides              bool
	_connectListHoursOfOperations                      bool
	_connectListInstanceAttributes                     bool
	_connectListInstanceStorageConfigs                 bool
	_connectListInstances                              bool
	_connectListIntegrationAssociations                bool
	_connectListLambdaFunctions                        bool
	_connectListLexBots                                bool
	_connectListNotifications                          bool
	_connectListPhoneNumbers                           bool
	_connectListPhoneNumbersV2                         bool
	_connectListPredefinedAttributes                   bool
	_connectListPrompts                                bool
	_connectListQueueQuickConnects                     bool
	_connectListQueues                                 bool
	_connectListQuickConnects                          bool
	_connectListRealtimeContactAnalysisSegmentsV2      bool
	_connectListRoutingProfileManualAssignmentQueues   bool
	_connectListRoutingProfileQueues                   bool
	_connectListRoutingProfiles                        bool
	_connectListRules                                  bool
	_connectListSecurityKeys                           bool
	_connectListSecurityProfileApplications            bool
	_connectListSecurityProfileFlowModules             bool
	_connectListSecurityProfilePermissions             bool
	_connectListSecurityProfiles                       bool
	_connectListTagsForResource                        bool
	_connectListTaskTemplates                          bool
	_connectListTestCaseExecutionRecords               bool
	_connectListTestCaseExecutions                     bool
	_connectListTestCases                              bool
	_connectListTrafficDistributionGroupUsers          bool
	_connectListTrafficDistributionGroups              bool
	_connectListUseCases                               bool
	_connectListUserHierarchyGroups                    bool
	_connectListUserNotifications                      bool
	_connectListUserProficiencies                      bool
	_connectListUsers                                  bool
	_connectListViewVersions                           bool
	_connectListViews                                  bool
	_connectListWorkspaceMedia                         bool
	_connectListWorkspacePages                         bool
	_connectListWorkspaces                             bool
	_connectMonitorContact                             bool
	_connectPauseContact                               bool
	_connectPutUserStatus                              bool
	_connectReleasePhoneNumber                         bool
	_connectReplicateInstance                          bool
	_connectResumeContact                              bool
	_connectResumeContactRecording                     bool
	_connectSearchAgentStatuses                        bool
	_connectSearchAvailablePhoneNumbers                bool
	_connectSearchContactEvaluations                   bool
	_connectSearchContactFlowModules                   bool
	_connectSearchContactFlows                         bool
	_connectSearchContacts                             bool
	_connectSearchDataTables                           bool
	_connectSearchEmailAddresses                       bool
	_connectSearchEvaluationForms                      bool
	_connectSearchHoursOfOperationOverrides            bool
	_connectSearchHoursOfOperations                    bool
	_connectSearchNotifications                        bool
	_connectSearchPredefinedAttributes                 bool
	_connectSearchPrompts                              bool
	_connectSearchQueues                               bool
	_connectSearchQuickConnects                        bool
	_connectSearchResourceTags                         bool
	_connectSearchRoutingProfiles                      bool
	_connectSearchSecurityProfiles                     bool
	_connectSearchTestCases                            bool
	_connectSearchUserHierarchyGroups                  bool
	_connectSearchUsers                                bool
	_connectSearchViews                                bool
	_connectSearchVocabularies                         bool
	_connectSearchWorkspaceAssociations                bool
	_connectSearchWorkspaces                           bool
	_connectSendChatIntegrationEvent                   bool
	_connectSendOutboundEmail                          bool
	_connectStartAttachedFileUpload                    bool
	_connectStartChatContact                           bool
	_connectStartContactEvaluation                     bool
	_connectStartContactMediaProcessing                bool
	_connectStartContactRecording                      bool
	_connectStartContactStreaming                      bool
	_connectStartEmailContact                          bool
	_connectStartOutboundChatContact                   bool
	_connectStartOutboundEmailContact                  bool
	_connectStartOutboundVoiceContact                  bool
	_connectStartScreenSharing                         bool
	_connectStartTaskContact                           bool
	_connectStartTestCaseExecution                     bool
	_connectStartWebRTCContact                         bool
	_connectStopContact                                bool
	_connectStopContactMediaProcessing                 bool
	_connectStopContactRecording                       bool
	_connectStopContactStreaming                       bool
	_connectStopTestCaseExecution                      bool
	_connectSubmitContactEvaluation                    bool
	_connectSuspendContactRecording                    bool
	_connectTagContact                                 bool
	_connectTagResource                                bool
	_connectTransferContact                            bool
	_connectUntagContact                               bool
	_connectUntagResource                              bool
	_connectUpdateAgentStatus                          bool
	_connectUpdateAuthenticationProfile                bool
	_connectUpdateContact                              bool
	_connectUpdateContactAttributes                    bool
	_connectUpdateContactEvaluation                    bool
	_connectUpdateContactFlowContent                   bool
	_connectUpdateContactFlowMetadata                  bool
	_connectUpdateContactFlowModuleAlias               bool
	_connectUpdateContactFlowModuleContent             bool
	_connectUpdateContactFlowModuleMetadata            bool
	_connectUpdateContactFlowName                      bool
	_connectUpdateContactRoutingData                   bool
	_connectUpdateContactSchedule                      bool
	_connectUpdateDataTableAttribute                   bool
	_connectUpdateDataTableMetadata                    bool
	_connectUpdateDataTablePrimaryValues               bool
	_connectUpdateEmailAddressMetadata                 bool
	_connectUpdateEvaluationForm                       bool
	_connectUpdateHoursOfOperation                     bool
	_connectUpdateHoursOfOperationOverride             bool
	_connectUpdateInstanceAttribute                    bool
	_connectUpdateInstanceStorageConfig                bool
	_connectUpdateNotificationContent                  bool
	_connectUpdateParticipantAuthentication            bool
	_connectUpdateParticipantRoleConfig                bool
	_connectUpdatePhoneNumber                          bool
	_connectUpdatePhoneNumberMetadata                  bool
	_connectUpdatePredefinedAttribute                  bool
	_connectUpdatePrompt                               bool
	_connectUpdateQueueHoursOfOperation                bool
	_connectUpdateQueueMaxContacts                     bool
	_connectUpdateQueueName                            bool
	_connectUpdateQueueOutboundCallerConfig            bool
	_connectUpdateQueueOutboundEmailConfig             bool
	_connectUpdateQueueStatus                          bool
	_connectUpdateQuickConnectConfig                   bool
	_connectUpdateQuickConnectName                     bool
	_connectUpdateRoutingProfileAgentAvailabilityTimer bool
	_connectUpdateRoutingProfileConcurrency            bool
	_connectUpdateRoutingProfileDefaultOutboundQueue   bool
	_connectUpdateRoutingProfileName                   bool
	_connectUpdateRoutingProfileQueues                 bool
	_connectUpdateRule                                 bool
	_connectUpdateSecurityProfile                      bool
	_connectUpdateTaskTemplate                         bool
	_connectUpdateTestCase                             bool
	_connectUpdateTrafficDistribution                  bool
	_connectUpdateUserConfig                           bool
	_connectUpdateUserHierarchy                        bool
	_connectUpdateUserHierarchyGroupName               bool
	_connectUpdateUserHierarchyStructure               bool
	_connectUpdateUserIdentityInfo                     bool
	_connectUpdateUserNotificationStatus               bool
	_connectUpdateUserPhoneConfig                      bool
	_connectUpdateUserProficiencies                    bool
	_connectUpdateUserRoutingProfile                   bool
	_connectUpdateUserSecurityProfiles                 bool
	_connectUpdateViewContent                          bool
	_connectUpdateViewMetadata                         bool
	_connectUpdateWorkspaceMetadata                    bool
	_connectUpdateWorkspacePage                        bool
	_connectUpdateWorkspaceTheme                       bool
	_connectUpdateWorkspaceVisibility                  bool

	_connectActions                              string
	_connectAdditionalRecipients                 string
	_connectAfterContactWorkConfigs              string
	_connectAgentAvailabilityTimer               string
	_connectAgentConfig                          string
	_connectAgentStatusId                        string
	_connectAgentStatusTypes                     string
	_connectAliasConfiguration                   string
	_connectAliasId                              string
	_connectAliasName                            string
	_connectAllowedAccessControlHierarchyGroupId string
	_connectAllowedAccessControlTags             string
	_connectAllowedCapabilities                  string
	_connectAllowedFlowModules                   string
	_connectAllowedIps                           []string
	_connectAllowedMonitorCapabilities           string
	_connectAnswerMachineDetectionConfig         string
	_connectAnswers                              string
	_connectApplications                         string
	_connectAsDraft                              string
	_connectAssociatedResourceArn                string
	_connectAssociationId                        string
	_connectAttachments                          string
	_connectAttributeConfiguration               string
	_connectAttributeIds                         []string
	_connectAttributeName                        string
	_connectAttributeType                        string
	_connectAttributes                           string
	_connectAuthenticationProfileId              string
	_connectAutoAcceptConfigs                    string
	_connectAutoEvaluationConfiguration          string
	_connectBlockedIps                           []string
	_connectBotName                              string
	_connectCampaignId                           string
	_connectChannel                              string
	_connectChannelConfiguration                 string
	_connectChatDurationInMinutes                string
	_connectChatStreamingConfiguration           string
	_connectClientToken                          string
	_connectCode                                 string
	_connectConfig                               string
	_connectConstraints                          string
	_connectContactConfiguration                 string
	_connectContactDataRequestList               string
	_connectContactFlowId                        string
	_connectContactFlowModuleId                  string
	_connectContactFlowModuleState               string
	_connectContactFlowModuleVersion             string
	_connectContactFlowState                     string
	_connectContactFlowTypes                     string
	_connectContactFlowVersion                   string
	_connectContactId                            string
	_connectContactRecordingType                 string
	_connectContent                              string
	_connectCreateNewVersion                     string
	_connectCreatedBy                            string
	_connectCurrentMetrics                       string
	_connectCustomerEndpoint                     string
	_connectCustomerId                           string
	_connectDataSetId                            string
	_connectDataSetIds                           []string
	_connectDataTableId                          string
	_connectDefaultOutboundQueueId               string
	_connectDefaults                             string
	_connectDescription                          string
	_connectDestinationEmailAddress              string
	_connectDestinationEndpoint                  string
	_connectDestinationId                        string
	_connectDestinationPhoneNumber               string
	_connectDeviceToken                          string
	_connectDeviceType                           string
	_connectDirectoryId                          string
	_connectDirectoryUserId                      string
	_connectDisconnectOnCustomerExit             string
	_connectDisconnectReason                     string
	_connectDisplayName                          string
	_connectDisplayOrder                         string
	_connectEffectiveFrom                        string
	_connectEffectiveTill                        string
	_connectEmailAddress                         string
	_connectEmailAddressId                       string
	_connectEmailMessage                         string
	_connectEndTime                              string
	_connectEntityArn                            string
	_connectEntityType                           string
	_connectEntryPoint                           string
	_connectError                                string
	_connectErrorDescription                     string
	_connectEvaluationFormId                     string
	_connectEvaluationFormVersion                string
	_connectEvaluationId                         string
	_connectEvent                                string
	_connectEventSourceName                      string
	_connectExpiresAt                            string
	_connectExpiryDurationInMinutes              string
	_connectExternalInvocationConfiguration      string
	_connectFailureMode                          string
	_connectFields                               string
	_connectFileId                               string
	_connectFileIds                              []string
	_connectFileName                             string
	_connectFileSizeInBytes                      string
	_connectFileUseCaseType                      string
	_connectFilters                              string
	_connectFlowContentSha256                    string
	_connectFlowId                               string
	_connectFlowModuleContentSha256              string
	_connectFromDate                             string
	_connectFromEmailAddress                     string
	_connectFunction                             string
	_connectFunctionArn                          string
	_connectGranularAccessControlConfiguration   string
	_connectGroupings                            []string
	_connectHierarchyGroupId                     string
	_connectHierarchyRestrictedResources         []string
	_connectHierarchyStructure                   string
	_connectHistoricalMetrics                    string
	_connectHoursOfOperationId                   string
	_connectHoursOfOperationOverrideId           string
	_connectId                                   string
	_connectIdentityInfo                         string
	_connectIdentityManagementType               string
	_connectInboundCallsEnabled                  string
	_connectInitialContactId                     string
	_connectInitialMessage                       string
	_connectInitialSystemMessage                 string
	_connectInitialTemplatedSystemMessage        string
	_connectInitializationData                   string
	_connectInitiateAs                           string
	_connectInitiationMethod                     string
	_connectInputData                            string
	_connectInstanceAlias                        string
	_connectInstanceId                           string
	_connectIntegrationArn                       string
	_connectIntegrationAssociationId             string
	_connectIntegrationType                      string
	_connectInterval                             string
	_connectItems                                string
	_connectKey                                  string
	_connectLanguageCode                         string
	_connectLanguageConfiguration                string
	_connectLastModifiedRegion                   string
	_connectLastModifiedTime                     string
	_connectLexBot                               string
	_connectLexRegion                            string
	_connectLexV2Bot                             string
	_connectLexVersion                           string
	_connectLockVersion                          string
	_connectManualAssignmentQueueConfigs         string
	_connectManualAssignmentQueueReferences      string
	_connectMaxContacts                          string
	_connectMaxResults                           string
	_connectMediaConcurrencies                   string
	_connectMediaSource                          string
	_connectMediaType                            string
	_connectMetrics                              string
	_connectName                                 string
	_connectNameStartsWith                       string
	_connectNewPage                              string
	_connectNewPrimaryValues                     string
	_connectNewSessionDetails                    string
	_connectNextToken                            string
	_connectNotes                                string
	_connectNotificationId                       string
	_connectOrigin                               string
	_connectOutboundCallerConfig                 string
	_connectOutboundCallsEnabled                 string
	_connectOutboundEmailConfig                  string
	_connectOutboundStrategy                     string
	_connectOutputType                           string
	_connectOverrideType                         string
	_connectPage                                 string
	_connectParentGroupId                        string
	_connectParentHoursOfOperationConfigs        string
	_connectParentHoursOfOperationIds            []string
	_connectParticipantConfiguration             string
	_connectParticipantDetails                   string
	_connectPassword                             string
	_connectPeriodicSessionDuration              string
	_connectPermissions                          []string
	_connectPersistentChat                       string
	_connectPersistentConnectionConfigs          string
	_connectPhoneConfig                          string
	_connectPhoneNumber                          string
	_connectPhoneNumberConfigs                   string
	_connectPhoneNumberCountryCode               string
	_connectPhoneNumberCountryCodes              string
	_connectPhoneNumberDescription               string
	_connectPhoneNumberId                        string
	_connectPhoneNumberPrefix                    string
	_connectPhoneNumberType                      string
	_connectPhoneNumberTypes                     string
	_connectPinpointAppArn                       string
	_connectPredefinedNotificationId             string
	_connectPreviousContactId                    string
	_connectPrimary                              string
	_connectPrimaryAttributeValues               string
	_connectPrimaryValues                        string
	_connectPriority                             string
	_connectProcessorArn                         string
	_connectPromptId                             string
	_connectPublishStatus                        string
	_connectPurposes                             []string
	_connectQueueConfigs                         string
	_connectQueueId                              string
	_connectQueueInfo                            string
	_connectQueuePriority                        string
	_connectQueueReferences                      string
	_connectQueueTimeAdjustmentSeconds           string
	_connectQueueTypes                           string
	_connectQuickConnectConfig                   string
	_connectQuickConnectId                       string
	_connectQuickConnectIds                      []string
	_connectQuickConnectTypes                    string
	_connectRecipients                           []string
	_connectRecordIds                            []string
	_connectRecurrenceConfig                     string
	_connectReferenceTypes                       string
	_connectReferences                           string
	_connectRegistrationId                       string
	_connectRehydrationType                      string
	_connectRelatedContactId                     string
	_connectReplicaAlias                         string
	_connectReplicaRegion                        string
	_connectResetOrderNumber                     string
	_connectResourceArn                          string
	_connectResourceArns                         []string
	_connectResourceId                           string
	_connectResourceIds                          []string
	_connectResourceType                         string
	_connectResourceTypes                        []string
	_connectReviewConfiguration                  string
	_connectRingTimeoutInSeconds                 string
	_connectRoutingCriteria                      string
	_connectRoutingProfileId                     string
	_connectRuleId                               string
	_connectS3Uri                                string
	_connectScheduledTime                        string
	_connectScoringStrategy                      string
	_connectSearchCriteria                       string
	_connectSearchFilter                         string
	_connectSecurityProfileId                    string
	_connectSecurityProfileIds                   []string
	_connectSecurityProfileName                  string
	_connectSecurityProfiles                     string
	_connectSegmentAttributes                    string
	_connectSegmentTypes                         string
	_connectSelfAssignFlowId                     string
	_connectSessionInactivityDuration            string
	_connectSessionInactivityHandlingEnabled     string
	_connectSettings                             string
	_connectSignInConfig                         string
	_connectSlug                                 string
	_connectSnapshotVersion                      string
	_connectSort                                 string
	_connectSortCriteria                         string
	_connectSourceApplicationName                string
	_connectSourceApplicationUrl                 string
	_connectSourceCampaign                       string
	_connectSourceContactId                      string
	_connectSourceEndpoint                       string
	_connectSourceId                             string
	_connectSourcePhoneNumber                    string
	_connectSourcePhoneNumberArn                 string
	_connectSourceType                           string
	_connectStartTime                            string
	_connectState                                string
	_connectStatus                               string
	_connectStorageConfig                        string
	_connectStreamingId                          string
	_connectSubmittedBy                          string
	_connectSubtype                              string
	_connectSupportedMessagingContentTypes       []string
	_connectSystemEndpoint                       string
	_connectTagKeys                              []string
	_connectTagRestrictedResources               []string
	_connectTags                                 string
	_connectTargetAccountId                      string
	_connectTargetArn                            string
	_connectTargetConfiguration                  string
	_connectTaskTemplateId                       string
	_connectTelephonyConfig                      string
	_connectTestCaseExecutionId                  string
	_connectTestCaseId                           string
	_connectTestCaseName                         string
	_connectTheme                                string
	_connectTimeRange                            string
	_connectTimeZone                             string
	_connectTitle                                string
	_connectToDate                               string
	_connectTrafficDistributionGroupId           string
	_connectTrafficType                          string
	_connectTriggerEventSource                   string
	_connectType                                 string
	_connectUpdatedBy                            string
	_connectUrlExpiryInSeconds                   string
	_connectUseCaseId                            string
	_connectUseCaseType                          string
	_connectUserId                               string
	_connectUserInfo                             string
	_connectUserProficiencies                    string
	_connectUsername                             string
	_connectValidation                           string
	_connectValue                                string
	_connectValueLockLevel                       string
	_connectValueType                            string
	_connectValues                               string
	_connectVersionDescription                   string
	_connectViewContentSha256                    string
	_connectViewId                               string
	_connectViewVersion                          string
	_connectVisibility                           string
	_connectVocabularyId                         string
	_connectVocabularyName                       string
	_connectVoiceEnhancementConfigs              string
	_connectVoiceRecordingConfiguration          string
	_connectWorkspaceId                          string
)

// Activates an evaluation form in the specified Amazon Connect instance. After
// the evaluation form is activated, it is available to start new evaluations based
// on the form.
func connect_ActivateEvaluationForm(cfg aws.Config, client *connect.Client) {
	input := &connect.ActivateEvaluationFormInput{
		// EvaluationFormId: *string, // Required
		// EvaluationFormVersion: int32, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectEvaluationFormVersion) > 0 {
		if err := assignInputField(input, "EvaluationFormVersion", _connectEvaluationFormVersion); err != nil {
			log.Errorf("invalid --evaluation-form-version: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.ActivateEvaluationForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified dataset for a Amazon Connect instance with the target
// account. You can associate only one dataset in a single call.
func connect_AssociateAnalyticsDataSet(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateAnalyticsDataSetInput{
		// DataSetId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataSetId) > 0 {
		input.DataSetId = aws.String(_connectDataSetId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTargetAccountId) > 0 {
		input.TargetAccountId = aws.String(_connectTargetAccountId)
	}

	if resp, err := client.AssociateAnalyticsDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Associates an approved origin to an Amazon Connect instance.
func connect_AssociateApprovedOrigin(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateApprovedOriginInput{
		// InstanceId: *string, // Required
		// Origin: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectOrigin) > 0 {
		input.Origin = aws.String(_connectOrigin)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.AssociateApprovedOrigin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Allows the specified Amazon Connect instance to access the specified Amazon Lex
// or Amazon Lex V2 bot.
func connect_AssociateBot(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateBotInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectLexBot) > 0 {
		if err := assignInputField(input, "LexBot", _connectLexBot); err != nil {
			log.Errorf("invalid --lex-bot: %s", err.Error())
			return
		}
	}
	if len(_connectLexV2Bot) > 0 {
		if err := assignInputField(input, "LexV2Bot", _connectLexV2Bot); err != nil {
			log.Errorf("invalid --lex-v2-bot: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a queued contact with an agent.
// # Use cases
//
// Following are common uses cases for this API:
//
// - Programmatically assign queued contacts to available users.
//
// - Leverage the IAM context key connect:PreferredUserArn to restrict contact
// association to specific preferred user.
//
// # Important things to know
//
// - Use this API with chat, email, and task contacts. It does not support voice
// contacts.
//
// - Use it to associate contacts with users regardless of their current state,
// including custom states. Ensure your application logic accounts for user
// availability before making associations.
//
// - It honors the IAM context key connect:PreferredUserArn to prevent
// unauthorized contact associations.
//
// - It respects the IAM context key connect:PreferredUserArn to enforce
// authorization controls and prevent unauthorized contact associations. Verify
// that your IAM policies are properly configured to support your intended use
// cases.
//
// - The service quota Queues per routing profile per instance applies to
// manually assigned queues, too. For more information about this quota, see [Amazon Connect quotas]in
// the Amazon Connect Administrator Guide.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
func connect_AssociateContactWithUser(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateContactWithUserInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.AssociateContactWithUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an existing vocabulary as the default. Contact Lens for Amazon
// Connect uses the vocabulary in post-call and real-time analysis sessions for the
// given language.
func connect_AssociateDefaultVocabulary(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateDefaultVocabularyInput{
		// InstanceId: *string, // Required
		// LanguageCode: types.VocabularyLanguageCode, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _connectLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_connectVocabularyId) > 0 {
		input.VocabularyId = aws.String(_connectVocabularyId)
	}

	if resp, err := client.AssociateDefaultVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an email address alias with an existing email address in an Amazon
// Connect instance. This creates a forwarding relationship where emails sent to
// the alias email address are automatically forwarded to the primary email
// address.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Unified customer support: Create multiple entry points (for example,
// support(at)example.com, help(at)example.com, customercare(at)example.com) that all
// forward to a single agent queue for streamlined management.
//
// - Department consolidation: Forward emails from legacy department addresses
// (for example, sales(at)example.com, info(at)example.com) to a centralized customer
// service email during organizational restructuring.
//
// - Brand management: Enable you to use familiar brand-specific email addresses
// that forward to the appropriate Amazon Connect instance email address.
//
// # Important things to know
//
// - Each email address can have a maximum of one alias. You cannot create
// multiple aliases for the same email address.
//
// - If the alias email address already receives direct emails, it continues to
// receive direct emails plus forwarded emails.
//
// - You cannot chain email aliases together (that is, create an alias of an
// alias).
//
// AssociateEmailAddressAlias does not return the following information:
//
// - A confirmation of the alias relationship details (you must call [DescribeEmailAddress]to verify).
//
// - The timestamp of when the association occurred.
//
// - The status of the forwarding configuration.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// # Related operations
//
// [DisassociateEmailAddressAlias]
// - : Removes the alias association between two email addresses in an Amazon
// Connect instance.
//
// [DescribeEmailAddress]
// - : View current alias configurations for an email address.
//
// [SearchEmailAddresses]
// - : Find email addresses and their alias relationships across an instance.
//
// [CreateEmailAddress]
// - : Create new email addresses that can participate in alias relationships.
//
// [DeleteEmailAddress]
// - : Remove email addresses (automatically removes any alias relationships).
//
// [UpdateEmailAddressMetadata]
// - : Modify email address properties (does not affect alias relationships).
//
// [DescribeEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeEmailAddress.html
// [DeleteEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteEmailAddress.html
// [DisassociateEmailAddressAlias]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DisassociateEmailAddressAlias.html
// [SearchEmailAddresses]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchEmailAddresses.html
// [UpdateEmailAddressMetadata]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdateEmailAddressMetadata.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [CreateEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateEmailAddress.html
func connect_AssociateEmailAddressAlias(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateEmailAddressAliasInput{
		// AliasConfiguration: *types.AliasConfiguration, // Required
		// EmailAddressId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAliasConfiguration) > 0 {
		if err := assignInputField(input, "AliasConfiguration", _connectAliasConfiguration); err != nil {
			log.Errorf("invalid --alias-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectEmailAddressId) > 0 {
		input.EmailAddressId = aws.String(_connectEmailAddressId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.AssociateEmailAddressAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a connect resource to a flow.
func connect_AssociateFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateFlowInput{
		// FlowId: *string, // Required
		// InstanceId: *string, // Required
		// ResourceId: *string, // Required
		// ResourceType: types.FlowAssociationResourceType, // Required
	}

	if len(_connectFlowId) > 0 {
		input.FlowId = aws.String(_connectFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceId) > 0 {
		input.ResourceId = aws.String(_connectResourceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of hours of operations with another hours of operation. Refer
// to Administrator Guide [here]for more information on inheriting overrides from parent
// hours of operation(s).
//
// [here]: https://docs.aws.amazon.com/connect/latest/adminguide/hours-of-operation-overrides.html
func connect_AssociateHoursOfOperations(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateHoursOfOperationsInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
		// ParentHoursOfOperationConfigs: []types.ParentHoursOfOperationConfig, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectParentHoursOfOperationConfigs) > 0 {
		if err := assignInputField(input, "ParentHoursOfOperationConfigs", _connectParentHoursOfOperationConfigs); err != nil {
			log.Errorf("invalid --parent-hours-of-operation-configs: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateHoursOfOperations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Associates a storage resource type for the first time. You can only associate
// one type of storage configuration in a single call. This means, for example,
// that you can't define an instance with multiple S3 buckets for storing chat
// transcripts.
//
// This API does not create a resource that doesn't exist. It only associates it
// to the instance. Ensure that the resource being specified in the storage
// configuration, like an S3 bucket, exists when being used for association.
func connect_AssociateInstanceStorageConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateInstanceStorageConfigInput{
		// InstanceId: *string, // Required
		// ResourceType: types.InstanceStorageResourceType, // Required
		// StorageConfig: *types.InstanceStorageConfig, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_connectStorageConfig) > 0 {
		if err := assignInputField(input, "StorageConfig", _connectStorageConfig); err != nil {
			log.Errorf("invalid --storage-config: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.AssociateInstanceStorageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Allows the specified Amazon Connect instance to access the specified Lambda
// function.
func connect_AssociateLambdaFunction(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateLambdaFunctionInput{
		// FunctionArn: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectFunctionArn) > 0 {
		input.FunctionArn = aws.String(_connectFunctionArn)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.AssociateLambdaFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Allows the specified Amazon Connect instance to access the specified Amazon Lex
// V1 bot. This API only supports the association of Amazon Lex V1 bots.
func connect_AssociateLexBot(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateLexBotInput{
		// InstanceId: *string, // Required
		// LexBot: *types.LexBot, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLexBot) > 0 {
		if err := assignInputField(input, "LexBot", _connectLexBot); err != nil {
			log.Errorf("invalid --lex-bot: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.AssociateLexBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a flow with a phone number claimed to your Amazon Connect instance.
// If the number is claimed to a traffic distribution group, and you are calling
// this API using an instance in the Amazon Web Services Region where the traffic
// distribution group was created, you can use either a full phone number ARN or
// UUID value for the PhoneNumberId URI request parameter. However, if the number
// is claimed to a traffic distribution group and you are calling this API using an
// instance in the alternate Amazon Web Services Region associated with the traffic
// distribution group, you must provide a full phone number ARN. If a UUID is
// provided in this scenario, you will receive a ResourceNotFoundException .
func connect_AssociatePhoneNumberContactFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociatePhoneNumberContactFlowInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
		// PhoneNumberId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_connectPhoneNumberId)
	}

	if resp, err := client.AssociatePhoneNumberContactFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of quick connects with a queue.
func connect_AssociateQueueQuickConnects(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateQueueQuickConnectsInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
		// QuickConnectIds: []string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectQuickConnectIds) > 0 {
		input.QuickConnectIds = append([]string(nil), _connectQuickConnectIds...)
	}

	if resp, err := client.AssociateQueueQuickConnects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of queues with a routing profile.
func connect_AssociateRoutingProfileQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateRoutingProfileQueuesInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectManualAssignmentQueueConfigs) > 0 {
		if err := assignInputField(input, "ManualAssignmentQueueConfigs", _connectManualAssignmentQueueConfigs); err != nil {
			log.Errorf("invalid --manual-assignment-queue-configs: %s", err.Error())
			return
		}
	}
	if len(_connectQueueConfigs) > 0 {
		if err := assignInputField(input, "QueueConfigs", _connectQueueConfigs); err != nil {
			log.Errorf("invalid --queue-configs: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateRoutingProfileQueues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Associates a security key to the instance.
func connect_AssociateSecurityKey(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateSecurityKeyInput{
		// InstanceId: *string, // Required
		// Key: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectKey) > 0 {
		input.Key = aws.String(_connectKey)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.AssociateSecurityKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate security profiles with an Entity in an Amazon Connect instance.
func connect_AssociateSecurityProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateSecurityProfilesInput{
		// EntityArn: *string, // Required
		// EntityType: types.EntityType, // Required
		// InstanceId: *string, // Required
		// SecurityProfiles: []types.SecurityProfileItem, // Required
	}

	if len(_connectEntityArn) > 0 {
		input.EntityArn = aws.String(_connectEntityArn)
	}
	if len(_connectEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _connectEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfiles) > 0 {
		if err := assignInputField(input, "SecurityProfiles", _connectSecurityProfiles); err != nil {
			log.Errorf("invalid --security-profiles: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateSecurityProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an agent with a traffic distribution group. This API can be called
// only in the Region where the traffic distribution group is created.
func connect_AssociateTrafficDistributionGroupUser(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateTrafficDistributionGroupUserInput{
		// InstanceId: *string, // Required
		// TrafficDistributionGroupId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTrafficDistributionGroupId) > 0 {
		input.TrafficDistributionGroupId = aws.String(_connectTrafficDistributionGroupId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.AssociateTrafficDistributionGroupUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of proficiencies with a user.
func connect_AssociateUserProficiencies(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateUserProficienciesInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
		// UserProficiencies: []types.UserProficiency, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectUserProficiencies) > 0 {
		if err := assignInputField(input, "UserProficiencies", _connectUserProficiencies); err != nil {
			log.Errorf("invalid --user-proficiencies: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateUserProficiencies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a workspace with one or more users or routing profiles, allowing
// them to access the workspace's configured views and pages.
func connect_AssociateWorkspace(cfg aws.Config, client *connect.Client) {
	input := &connect.AssociateWorkspaceInput{
		// InstanceId: *string, // Required
		// ResourceArns: []string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _connectResourceArns...)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.AssociateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a list of analytics datasets for a given Amazon Connect instance to
// a target account. You can associate multiple datasets in a single call.
func connect_BatchAssociateAnalyticsDataSet(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchAssociateAnalyticsDataSetInput{
		// DataSetIds: []string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataSetIds) > 0 {
		input.DataSetIds = append([]string(nil), _connectDataSetIds...)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTargetAccountId) > 0 {
		input.TargetAccountId = aws.String(_connectTargetAccountId)
	}

	if resp, err := client.BatchAssociateAnalyticsDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates values for attributes in a data table. The value may be a default or it
// may be associated with a primary value. The value must pass all customer defined
// validation as well as the default validation for the value type. The operation
// must conform to Batch Operation API Standards. Although the standard specifies
// that successful and failed entities are listed separately in the response,
// authorization fails if any primary values or attributes are unauthorized. The
// combination of primary values and the attribute name serve as the identifier for
// the individual item request.
func connect_BatchCreateDataTableValue(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchCreateDataTableValueInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Values: []types.DataTableValue, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchCreateDataTableValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes multiple values from a data table. API users may delete values at any
// time. When deletion is requested from the admin website, a warning is shown
// alerting the user of the most recent time the attribute and its values were
// accessed. System managed values are not deletable by customers.
func connect_BatchDeleteDataTableValue(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchDeleteDataTableValueInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Values: []types.DataTableDeleteValueIdentifier, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteDataTableValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves multiple values from a data table without evaluating expressions.
// Returns the raw stored values along with metadata such as lock versions and
// modification timestamps. "Describe" is a deprecated term but is allowed to
// maintain consistency with existing operations.
func connect_BatchDescribeDataTableValue(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchDescribeDataTableValueInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Values: []types.DataTableValueIdentifier, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDescribeDataTableValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a list of analytics datasets associated with a given Amazon Connect
// instance. You can disassociate multiple datasets in a single call.
func connect_BatchDisassociateAnalyticsDataSet(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchDisassociateAnalyticsDataSetInput{
		// DataSetIds: []string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataSetIds) > 0 {
		input.DataSetIds = append([]string(nil), _connectDataSetIds...)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTargetAccountId) > 0 {
		input.TargetAccountId = aws.String(_connectTargetAccountId)
	}

	if resp, err := client.BatchDisassociateAnalyticsDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to retrieve metadata about multiple attached files on an associated
// resource. Each attached file provided in the input list must be associated with
// the input AssociatedResourceArn.
func connect_BatchGetAttachedFileMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchGetAttachedFileMetadataInput{
		// AssociatedResourceArn: *string, // Required
		// FileIds: []string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAssociatedResourceArn) > 0 {
		input.AssociatedResourceArn = aws.String(_connectAssociatedResourceArn)
	}
	if len(_connectFileIds) > 0 {
		input.FileIds = append([]string(nil), _connectFileIds...)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.BatchGetAttachedFileMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the flow associations for the given resources.
func connect_BatchGetFlowAssociation(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchGetFlowAssociationInput{
		// InstanceId: *string, // Required
		// ResourceIds: []string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceIds) > 0 {
		input.ResourceIds = append([]string(nil), _connectResourceIds...)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetFlowAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Only the Amazon Connect outbound campaigns service principal is allowed to
// assume a role in your account and call this API.
//
// Allows you to create a batch of contacts in Amazon Connect. The outbound
// campaigns capability ingests dial requests via the [PutDialRequestBatch]API. It then uses
// BatchPutContact to create contacts corresponding to those dial requests. If
// agents are available, the dial requests are dialed out, which results in a voice
// call. The resulting voice call uses the same contactId that was created by
// BatchPutContact.
//
// [PutDialRequestBatch]: https://docs.aws.amazon.com/connect-outbound/latest/APIReference/API_PutDialRequestBatch.html
func connect_BatchPutContact(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchPutContactInput{
		// ContactDataRequestList: []types.ContactDataRequest, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactDataRequestList) > 0 {
		if err := assignInputField(input, "ContactDataRequestList", _connectContactDataRequestList); err != nil {
			log.Errorf("invalid --contact-data-request-list: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.BatchPutContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates multiple data table values using all properties from
// BatchCreateDataTableValue. System managed values are not modifiable by
// customers. The operation requires proper lock versions to prevent concurrent
// modification conflicts.
func connect_BatchUpdateDataTableValue(cfg aws.Config, client *connect.Client) {
	input := &connect.BatchUpdateDataTableValueInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Values: []types.DataTableValue, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateDataTableValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Claims an available phone number to your Amazon Connect instance or traffic
// distribution group. You can call this API only in the same Amazon Web Services
// Region where the Amazon Connect instance or traffic distribution group was
// created.
//
// For more information about how to use this operation, see [Claim a phone number in your country] and [Claim phone numbers to traffic distribution groups] in the Amazon
// Connect Administrator Guide.
//
// You can call the [SearchAvailablePhoneNumbers] API for available phone numbers that you can claim. Call the [DescribePhoneNumber]
// API to verify the status of a previous [ClaimPhoneNumber]operation.
//
// If you plan to claim and release numbers frequently, contact us for a service
// quota exception. Otherwise, it is possible you will be blocked from claiming and
// releasing any more numbers until up to 180 days past the oldest number released
// has expired.
//
// By default you can claim and release up to 200% of your maximum number of
// active phone numbers. If you claim and release phone numbers using the UI or API
// during a rolling 180 day cycle that exceeds 200% of your phone number service
// level quota, you will be blocked from claiming any more numbers until 180 days
// past the oldest number released has expired.
//
// For example, if you already have 99 claimed numbers and a service level quota
// of 99 phone numbers, and in any 180 day period you release 99, claim 99, and
// then release 99, you will have exceeded the 200% limit. At that point you are
// blocked from claiming any more numbers until you open an Amazon Web Services
// support ticket.
//
// [Claim phone numbers to traffic distribution groups]: https://docs.aws.amazon.com/connect/latest/adminguide/claim-phone-numbers-traffic-distribution-groups.html
// [Claim a phone number in your country]: https://docs.aws.amazon.com/connect/latest/adminguide/claim-phone-number.html
// [SearchAvailablePhoneNumbers]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
// [ClaimPhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ClaimPhoneNumber.html
func connect_ClaimPhoneNumber(cfg aws.Config, client *connect.Client) {
	input := &connect.ClaimPhoneNumberInput{
		// PhoneNumber: *string, // Required
	}

	if len(_connectPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_connectPhoneNumber)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPhoneNumberDescription) > 0 {
		input.PhoneNumberDescription = aws.String(_connectPhoneNumberDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectTargetArn) > 0 {
		input.TargetArn = aws.String(_connectTargetArn)
	}

	if resp, err := client.ClaimPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to confirm that the attached file has been uploaded using the
// pre-signed URL provided in the StartAttachedFileUpload API.
func connect_CompleteAttachedFileUpload(cfg aws.Config, client *connect.Client) {
	input := &connect.CompleteAttachedFileUploadInput{
		// AssociatedResourceArn: *string, // Required
		// FileId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAssociatedResourceArn) > 0 {
		input.AssociatedResourceArn = aws.String(_connectAssociatedResourceArn)
	}
	if len(_connectFileId) > 0 {
		input.FileId = aws.String(_connectFileId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.CompleteAttachedFileUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an agent status for the specified Amazon Connect instance.
func connect_CreateAgentStatus(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateAgentStatusInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// State: types.AgentStatusState, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectState) > 0 {
		if err := assignInputField(input, "State", _connectState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectDisplayOrder) > 0 {
		if err := assignInputField(input, "DisplayOrder", _connectDisplayOrder); err != nil {
			log.Errorf("invalid --display-order: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Only the VOICE, EMAIL, and TASK channels are supported.
// - For VOICE: The supported initiation method is TRANSFER . The contacts
// created with this initiation method have a subtype connect:ExternalAudio .
//
// - For EMAIL: The supported initiation methods are OUTBOUND , AGENT_REPLY , and
// FLOW .
//
// - For TASK: The supported initiation method is API . Contacts created with
// this API have a sub-type of connect:ExternalTask .
//
// Creates a new VOICE, EMAIL, or TASK contact.
//
// After a contact is created, you can move it to the desired state by using the
// InitiateAs parameter. While you can use API to create task contacts that are in
// the COMPLETED state, you must contact Amazon Web Services Support before using
// it for bulk import use cases. Bulk import causes your requests to be throttled
// or fail if your CreateContact limits aren't high enough.
func connect_CreateContact(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateContactInput{
		// Channel: types.Channel, // Required
		// InitiationMethod: types.ContactInitiationMethod, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectChannel) > 0 {
		if err := assignInputField(input, "Channel", _connectChannel); err != nil {
			log.Errorf("invalid --channel: %s", err.Error())
			return
		}
	}
	if len(_connectInitiationMethod) > 0 {
		if err := assignInputField(input, "InitiationMethod", _connectInitiationMethod); err != nil {
			log.Errorf("invalid --initiation-method: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectExpiryDurationInMinutes) > 0 {
		if err := assignInputField(input, "ExpiryDurationInMinutes", _connectExpiryDurationInMinutes); err != nil {
			log.Errorf("invalid --expiry-duration-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_connectInitiateAs) > 0 {
		if err := assignInputField(input, "InitiateAs", _connectInitiateAs); err != nil {
			log.Errorf("invalid --initiate-as: %s", err.Error())
			return
		}
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectPreviousContactId) > 0 {
		input.PreviousContactId = aws.String(_connectPreviousContactId)
	}
	if len(_connectReferences) > 0 {
		if err := assignInputField(input, "References", _connectReferences); err != nil {
			log.Errorf("invalid --references: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}
	if len(_connectSegmentAttributes) > 0 {
		if err := assignInputField(input, "SegmentAttributes", _connectSegmentAttributes); err != nil {
			log.Errorf("invalid --segment-attributes: %s", err.Error())
			return
		}
	}
	if len(_connectUserInfo) > 0 {
		if err := assignInputField(input, "UserInfo", _connectUserInfo); err != nil {
			log.Errorf("invalid --user-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a flow for the specified Amazon Connect instance.
// You can also create and update flows using the [Amazon Connect Flow language].
//
// [Amazon Connect Flow language]: https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html
func connect_CreateContactFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateContactFlowInput{
		// Content: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Type: types.ContactFlowType, // Required
	}

	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectType) > 0 {
		if err := assignInputField(input, "Type", _connectType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContactFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a flow module for the specified Amazon Connect instance.
func connect_CreateContactFlowModule(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateContactFlowModuleInput{
		// Content: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectExternalInvocationConfiguration) > 0 {
		if err := assignInputField(input, "ExternalInvocationConfiguration", _connectExternalInvocationConfiguration); err != nil {
			log.Errorf("invalid --external-invocation-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectSettings) > 0 {
		input.Settings = aws.String(_connectSettings)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContactFlowModule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a named alias that points to a specific version of a contact flow
// module.
func connect_CreateContactFlowModuleAlias(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateContactFlowModuleAliasInput{
		// AliasName: *string, // Required
		// ContactFlowModuleId: *string, // Required
		// ContactFlowModuleVersion: *int64, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAliasName) > 0 {
		input.AliasName = aws.String(_connectAliasName)
	}
	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectContactFlowModuleVersion) > 0 {
		if err := assignInputField(input, "ContactFlowModuleVersion", _connectContactFlowModuleVersion); err != nil {
			log.Errorf("invalid --contact-flow-module-version: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}

	if resp, err := client.CreateContactFlowModuleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an immutable snapshot of a contact flow module, preserving its content
// and settings at a specific point in time for version control and rollback
// capabilities.
func connect_CreateContactFlowModuleVersion(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateContactFlowModuleVersionInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectFlowModuleContentSha256) > 0 {
		input.FlowModuleContentSha256 = aws.String(_connectFlowModuleContentSha256)
	}

	if resp, err := client.CreateContactFlowModuleVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a new version of the flow provided. Versions are immutable and
// monotonically increasing. If the FlowContentSha256 provided is different from
// the FlowContentSha256 of the $LATEST published flow content, then an error is
// returned. This API only supports creating versions for flows of type Campaign .
func connect_CreateContactFlowVersion(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateContactFlowVersionInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowVersion) > 0 {
		if err := assignInputField(input, "ContactFlowVersion", _connectContactFlowVersion); err != nil {
			log.Errorf("invalid --contact-flow-version: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectFlowContentSha256) > 0 {
		input.FlowContentSha256 = aws.String(_connectFlowContentSha256)
	}
	if len(_connectLastModifiedRegion) > 0 {
		input.LastModifiedRegion = aws.String(_connectLastModifiedRegion)
	}
	if len(_connectLastModifiedTime) > 0 {
		if err := assignInputField(input, "LastModifiedTime", _connectLastModifiedTime); err != nil {
			log.Errorf("invalid --last-modified-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContactFlowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new data table with the specified properties. Supports the creation
// of all table properties except for attributes and values. A table with no
// attributes and values is a valid state for a table. The number of tables per
// instance is limited to 100 per instance. Customers can request an increase by
// using Amazon Web Services Service Quotas.
func connect_CreateDataTable(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateDataTableInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Status: types.DataTableStatus, // Required
		// TimeZone: *string, // Required
		// ValueLockLevel: types.DataTableLockLevel, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectTimeZone) > 0 {
		input.TimeZone = aws.String(_connectTimeZone)
	}
	if len(_connectValueLockLevel) > 0 {
		if err := assignInputField(input, "ValueLockLevel", _connectValueLockLevel); err != nil {
			log.Errorf("invalid --value-lock-level: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an attribute to an existing data table. Creating a new primary attribute
// uses the empty value for the specified value type for all existing records. This
// should not affect uniqueness of published data tables since the existing primary
// values will already be unique. Creating attributes does not create any values.
// System managed tables may not allow customers to create new attributes.
func connect_CreateDataTableAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateDataTableAttributeInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// ValueType: types.DataTableAttributeValueType, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectValueType) > 0 {
		if err := assignInputField(input, "ValueType", _connectValueType); err != nil {
			log.Errorf("invalid --value-type: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectPrimary) > 0 {
		if err := assignInputField(input, "Primary", _connectPrimary); err != nil {
			log.Errorf("invalid --primary: %s", err.Error())
			return
		}
	}
	if len(_connectValidation) > 0 {
		if err := assignInputField(input, "Validation", _connectValidation); err != nil {
			log.Errorf("invalid --validation: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataTableAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create new email address in the specified Amazon Connect instance. For more
// information about email addresses, see [Create email addresses]in the Amazon Connect Administrator
// Guide.
//
// [Create email addresses]: https://docs.aws.amazon.com/connect/latest/adminguide/create-email-address1.html
func connect_CreateEmailAddress(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateEmailAddressInput{
		// EmailAddress: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEmailAddress) > 0 {
		input.EmailAddress = aws.String(_connectEmailAddress)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectDisplayName) > 0 {
		input.DisplayName = aws.String(_connectDisplayName)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEmailAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an evaluation form in the specified Amazon Connect instance. The form
// can be used to define questions related to agent performance, and create
// sections to organize such questions. Question and section identifiers cannot be
// duplicated within the same evaluation form.
func connect_CreateEvaluationForm(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateEvaluationFormInput{
		// InstanceId: *string, // Required
		// Items: []types.EvaluationFormItem, // Required
		// Title: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectItems) > 0 {
		if err := assignInputField(input, "Items", _connectItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}
	if len(_connectTitle) > 0 {
		input.Title = aws.String(_connectTitle)
	}
	if len(_connectAsDraft) > 0 {
		if err := assignInputField(input, "AsDraft", _connectAsDraft); err != nil {
			log.Errorf("invalid --as-draft: %s", err.Error())
			return
		}
	}
	if len(_connectAutoEvaluationConfiguration) > 0 {
		if err := assignInputField(input, "AutoEvaluationConfiguration", _connectAutoEvaluationConfiguration); err != nil {
			log.Errorf("invalid --auto-evaluation-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectLanguageConfiguration) > 0 {
		if err := assignInputField(input, "LanguageConfiguration", _connectLanguageConfiguration); err != nil {
			log.Errorf("invalid --language-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectReviewConfiguration) > 0 {
		if err := assignInputField(input, "ReviewConfiguration", _connectReviewConfiguration); err != nil {
			log.Errorf("invalid --review-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectScoringStrategy) > 0 {
		if err := assignInputField(input, "ScoringStrategy", _connectScoringStrategy); err != nil {
			log.Errorf("invalid --scoring-strategy: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectTargetConfiguration) > 0 {
		if err := assignInputField(input, "TargetConfiguration", _connectTargetConfiguration); err != nil {
			log.Errorf("invalid --target-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEvaluationForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates hours of operation.
func connect_CreateHoursOfOperation(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateHoursOfOperationInput{
		// Config: []types.HoursOfOperationConfig, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// TimeZone: *string, // Required
	}

	if len(_connectConfig) > 0 {
		if err := assignInputField(input, "Config", _connectConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectTimeZone) > 0 {
		input.TimeZone = aws.String(_connectTimeZone)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectParentHoursOfOperationConfigs) > 0 {
		if err := assignInputField(input, "ParentHoursOfOperationConfigs", _connectParentHoursOfOperationConfigs); err != nil {
			log.Errorf("invalid --parent-hours-of-operation-configs: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHoursOfOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an hours of operation override in an Amazon Connect hours of operation
// resource.
func connect_CreateHoursOfOperationOverride(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateHoursOfOperationOverrideInput{
		// Config: []types.HoursOfOperationOverrideConfig, // Required
		// EffectiveFrom: *string, // Required
		// EffectiveTill: *string, // Required
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectConfig) > 0 {
		if err := assignInputField(input, "Config", _connectConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_connectEffectiveFrom) > 0 {
		input.EffectiveFrom = aws.String(_connectEffectiveFrom)
	}
	if len(_connectEffectiveTill) > 0 {
		input.EffectiveTill = aws.String(_connectEffectiveTill)
	}
	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectOverrideType) > 0 {
		if err := assignInputField(input, "OverrideType", _connectOverrideType); err != nil {
			log.Errorf("invalid --override-type: %s", err.Error())
			return
		}
	}
	if len(_connectRecurrenceConfig) > 0 {
		if err := assignInputField(input, "RecurrenceConfig", _connectRecurrenceConfig); err != nil {
			log.Errorf("invalid --recurrence-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHoursOfOperationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Initiates an Amazon Connect instance with all the supported channels enabled.
// It does not attach any storage, such as Amazon Simple Storage Service (Amazon
// S3) or Amazon Kinesis. It also does not allow for any configurations on
// features, such as Contact Lens for Amazon Connect.
//
// For more information, see [Create an Amazon Connect instance] in the Amazon Connect Administrator Guide.
//
// Amazon Connect enforces a limit on the total number of instances that you can
// create or delete in 30 days. If you exceed this limit, you will get an error
// message indicating there has been an excessive number of attempts at creating or
// deleting instances. You must wait 30 days before you can restart creating and
// deleting instances in your account.
//
// [Create an Amazon Connect instance]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-instances.html
func connect_CreateInstance(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateInstanceInput{
		// IdentityManagementType: types.DirectoryType, // Required
		// InboundCallsEnabled: *bool, // Required
		// OutboundCallsEnabled: *bool, // Required
	}

	if len(_connectIdentityManagementType) > 0 {
		if err := assignInputField(input, "IdentityManagementType", _connectIdentityManagementType); err != nil {
			log.Errorf("invalid --identity-management-type: %s", err.Error())
			return
		}
	}
	if len(_connectInboundCallsEnabled) > 0 {
		if err := assignInputField(input, "InboundCallsEnabled", _connectInboundCallsEnabled); err != nil {
			log.Errorf("invalid --inbound-calls-enabled: %s", err.Error())
			return
		}
	}
	if len(_connectOutboundCallsEnabled) > 0 {
		if err := assignInputField(input, "OutboundCallsEnabled", _connectOutboundCallsEnabled); err != nil {
			log.Errorf("invalid --outbound-calls-enabled: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDirectoryId) > 0 {
		input.DirectoryId = aws.String(_connectDirectoryId)
	}
	if len(_connectInstanceAlias) > 0 {
		input.InstanceAlias = aws.String(_connectInstanceAlias)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services resource association with an Amazon Connect
// instance.
func connect_CreateIntegrationAssociation(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateIntegrationAssociationInput{
		// InstanceId: *string, // Required
		// IntegrationArn: *string, // Required
		// IntegrationType: types.IntegrationType, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_connectIntegrationArn)
	}
	if len(_connectIntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _connectIntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_connectSourceApplicationName) > 0 {
		input.SourceApplicationName = aws.String(_connectSourceApplicationName)
	}
	if len(_connectSourceApplicationUrl) > 0 {
		input.SourceApplicationUrl = aws.String(_connectSourceApplicationUrl)
	}
	if len(_connectSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _connectSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegrationAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new notification to be delivered to specified recipients.
// Notifications can include localized content with links, and an optional
// expiration time. Recipients can be specified as individual user ARNs or instance
// ARNs to target all users in an instance.
func connect_CreateNotification(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateNotificationInput{
		// Content: map[string]string, // Required
		// InstanceId: *string, // Required
		// Recipients: []string, // Required
	}

	if len(_connectContent) > 0 {
		if err := assignInputField(input, "Content", _connectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRecipients) > 0 {
		input.Recipients = append([]string(nil), _connectRecipients...)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectExpiresAt) > 0 {
		if err := assignInputField(input, "ExpiresAt", _connectExpiresAt); err != nil {
			log.Errorf("invalid --expires-at: %s", err.Error())
			return
		}
	}
	if len(_connectPredefinedNotificationId) > 0 {
		input.PredefinedNotificationId = aws.String(_connectPredefinedNotificationId)
	}
	if len(_connectPriority) > 0 {
		if err := assignInputField(input, "Priority", _connectPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new participant into an on-going chat contact or webRTC call. For more
// information, see [Customize chat flow experiences by integrating custom participants]or [Enable multi-user web, in-app, and video calling].
//
// [Enable multi-user web, in-app, and video calling]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-multiuser-inapp.html
// [Customize chat flow experiences by integrating custom participants]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-customize-flow.html
func connect_CreateParticipant(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateParticipantInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// ParticipantDetails: *types.ParticipantDetailsToAdd, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectParticipantDetails) > 0 {
		if err := assignInputField(input, "ParticipantDetails", _connectParticipantDetails); err != nil {
			log.Errorf("invalid --participant-details: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.CreateParticipant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables rehydration of chats for the lifespan of a contact. For more
// information about chat rehydration, see [Enable persistent chat]in the Amazon Connect Administrator
// Guide.
//
// [Enable persistent chat]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-persistence.html
func connect_CreatePersistentContactAssociation(cfg aws.Config, client *connect.Client) {
	input := &connect.CreatePersistentContactAssociationInput{
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
		// RehydrationType: types.RehydrationType, // Required
		// SourceContactId: *string, // Required
	}

	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRehydrationType) > 0 {
		if err := assignInputField(input, "RehydrationType", _connectRehydrationType); err != nil {
			log.Errorf("invalid --rehydration-type: %s", err.Error())
			return
		}
	}
	if len(_connectSourceContactId) > 0 {
		input.SourceContactId = aws.String(_connectSourceContactId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.CreatePersistentContactAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new predefined attribute for the specified Amazon Connect instance. A
// predefined attribute is made up of a name and a value.
//
// For the predefined attributes per instance quota, see [Amazon Connect quotas].
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Create an attribute for routing proficiency (for example, agent
// certification) that has predefined values (for example, a list of possible
// certifications). For more information, see [Create predefined attributes for routing contacts to agents].
//
// - Create an attribute for business unit name that has a list of predefined
// business unit names used in your organization. This is a use case where
// information for a contact varies between transfers or conferences. For more
// information, see [Use contact segment attributes].
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Use contact segment attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/use-contact-segment-attributes.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
func connect_CreatePredefinedAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.CreatePredefinedAttributeInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectAttributeConfiguration) > 0 {
		if err := assignInputField(input, "AttributeConfiguration", _connectAttributeConfiguration); err != nil {
			log.Errorf("invalid --attribute-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectPurposes) > 0 {
		input.Purposes = append([]string(nil), _connectPurposes...)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePredefinedAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a prompt. For more information about prompts, such as supported file
// types and maximum length, see [Create prompts]in the Amazon Connect Administrator Guide.
//
// [Create prompts]: https://docs.aws.amazon.com/connect/latest/adminguide/prompts.html
func connect_CreatePrompt(cfg aws.Config, client *connect.Client) {
	input := &connect.CreatePromptInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// S3Uri: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectS3Uri) > 0 {
		input.S3Uri = aws.String(_connectS3Uri)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates registration for a device token and a chat contact to receive real-time
// push notifications. For more information about push notifications, see [Set up push notifications in Amazon Connect for mobile chat]in the
// Amazon Connect Administrator Guide.
//
// [Set up push notifications in Amazon Connect for mobile chat]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-push-notifications-for-mobile-chat.html
func connect_CreatePushNotificationRegistration(cfg aws.Config, client *connect.Client) {
	input := &connect.CreatePushNotificationRegistrationInput{
		// ContactConfiguration: *types.ContactConfiguration, // Required
		// DeviceToken: *string, // Required
		// DeviceType: types.DeviceType, // Required
		// InstanceId: *string, // Required
		// PinpointAppArn: *string, // Required
	}

	if len(_connectContactConfiguration) > 0 {
		if err := assignInputField(input, "ContactConfiguration", _connectContactConfiguration); err != nil {
			log.Errorf("invalid --contact-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectDeviceToken) > 0 {
		input.DeviceToken = aws.String(_connectDeviceToken)
	}
	if len(_connectDeviceType) > 0 {
		if err := assignInputField(input, "DeviceType", _connectDeviceType); err != nil {
			log.Errorf("invalid --device-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPinpointAppArn) > 0 {
		input.PinpointAppArn = aws.String(_connectPinpointAppArn)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.CreatePushNotificationRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new queue for the specified Amazon Connect instance.
// - If the phone number is claimed to a traffic distribution group that was
// created in the same Region as the Amazon Connect instance where you are calling
// this API, then you can use a full phone number ARN or a UUID for
// OutboundCallerIdNumberId . However, if the phone number is claimed to a
// traffic distribution group that is in one Region, and you are calling this API
// from an instance in another Amazon Web Services Region that is associated with
// the traffic distribution group, you must provide a full phone number ARN. If a
// UUID is provided in this scenario, you will receive a
// ResourceNotFoundException .
//
// - Only use the phone number ARN format that doesn't contain instance in the
// path, for example, arn:aws:connect:us-east-1:1234567890:phone-number/uuid .
// This is the same ARN format that is returned when you call the [ListPhoneNumbersV2]API.
//
// - If you plan to use IAM policies to allow/deny access to this API for phone
// number resources claimed to a traffic distribution group, see [Allow or Deny queue API actions for phone numbers in a replica Region].
//
// [ListPhoneNumbersV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html
// [Allow or Deny queue API actions for phone numbers in a replica Region]: https://docs.aws.amazon.com/connect/latest/adminguide/security_iam_resource-level-policy-examples.html#allow-deny-queue-actions-replica-region
func connect_CreateQueue(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateQueueInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectMaxContacts) > 0 {
		if err := assignInputField(input, "MaxContacts", _connectMaxContacts); err != nil {
			log.Errorf("invalid --max-contacts: %s", err.Error())
			return
		}
	}
	if len(_connectOutboundCallerConfig) > 0 {
		if err := assignInputField(input, "OutboundCallerConfig", _connectOutboundCallerConfig); err != nil {
			log.Errorf("invalid --outbound-caller-config: %s", err.Error())
			return
		}
	}
	if len(_connectOutboundEmailConfig) > 0 {
		if err := assignInputField(input, "OutboundEmailConfig", _connectOutboundEmailConfig); err != nil {
			log.Errorf("invalid --outbound-email-config: %s", err.Error())
			return
		}
	}
	if len(_connectQuickConnectIds) > 0 {
		input.QuickConnectIds = append([]string(nil), _connectQuickConnectIds...)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a quick connect for the specified Amazon Connect instance.
func connect_CreateQuickConnect(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateQuickConnectInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// QuickConnectConfig: *types.QuickConnectConfig, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectQuickConnectConfig) > 0 {
		if err := assignInputField(input, "QuickConnectConfig", _connectQuickConnectConfig); err != nil {
			log.Errorf("invalid --quick-connect-config: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQuickConnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new routing profile.
func connect_CreateRoutingProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateRoutingProfileInput{
		// DefaultOutboundQueueId: *string, // Required
		// Description: *string, // Required
		// InstanceId: *string, // Required
		// MediaConcurrencies: []types.MediaConcurrency, // Required
		// Name: *string, // Required
	}

	if len(_connectDefaultOutboundQueueId) > 0 {
		input.DefaultOutboundQueueId = aws.String(_connectDefaultOutboundQueueId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMediaConcurrencies) > 0 {
		if err := assignInputField(input, "MediaConcurrencies", _connectMediaConcurrencies); err != nil {
			log.Errorf("invalid --media-concurrencies: %s", err.Error())
			return
		}
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectAgentAvailabilityTimer) > 0 {
		if err := assignInputField(input, "AgentAvailabilityTimer", _connectAgentAvailabilityTimer); err != nil {
			log.Errorf("invalid --agent-availability-timer: %s", err.Error())
			return
		}
	}
	if len(_connectManualAssignmentQueueConfigs) > 0 {
		if err := assignInputField(input, "ManualAssignmentQueueConfigs", _connectManualAssignmentQueueConfigs); err != nil {
			log.Errorf("invalid --manual-assignment-queue-configs: %s", err.Error())
			return
		}
	}
	if len(_connectQueueConfigs) > 0 {
		if err := assignInputField(input, "QueueConfigs", _connectQueueConfigs); err != nil {
			log.Errorf("invalid --queue-configs: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoutingProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a rule for the specified Amazon Connect instance.
// Use the [Rules Function language] to code conditions for the rule.
//
// [Rules Function language]: https://docs.aws.amazon.com/connect/latest/APIReference/connect-rules-language.html
func connect_CreateRule(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateRuleInput{
		// Actions: []types.RuleAction, // Required
		// Function: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// PublishStatus: types.RulePublishStatus, // Required
		// TriggerEventSource: *types.RuleTriggerEventSource, // Required
	}

	if len(_connectActions) > 0 {
		if err := assignInputField(input, "Actions", _connectActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_connectFunction) > 0 {
		input.Function = aws.String(_connectFunction)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectPublishStatus) > 0 {
		if err := assignInputField(input, "PublishStatus", _connectPublishStatus); err != nil {
			log.Errorf("invalid --publish-status: %s", err.Error())
			return
		}
	}
	if len(_connectTriggerEventSource) > 0 {
		if err := assignInputField(input, "TriggerEventSource", _connectTriggerEventSource); err != nil {
			log.Errorf("invalid --trigger-event-source: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a security profile.
// For information about security profiles, see [Security Profiles] in the Amazon Connect
// Administrator Guide. For a mapping of the API name and user interface name of
// the security profile permissions, see [List of security profile permissions].
//
// [Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html
// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
func connect_CreateSecurityProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateSecurityProfileInput{
		// InstanceId: *string, // Required
		// SecurityProfileName: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_connectSecurityProfileName)
	}
	if len(_connectAllowedAccessControlHierarchyGroupId) > 0 {
		input.AllowedAccessControlHierarchyGroupId = aws.String(_connectAllowedAccessControlHierarchyGroupId)
	}
	if len(_connectAllowedAccessControlTags) > 0 {
		if err := assignInputField(input, "AllowedAccessControlTags", _connectAllowedAccessControlTags); err != nil {
			log.Errorf("invalid --allowed-access-control-tags: %s", err.Error())
			return
		}
	}
	if len(_connectAllowedFlowModules) > 0 {
		if err := assignInputField(input, "AllowedFlowModules", _connectAllowedFlowModules); err != nil {
			log.Errorf("invalid --allowed-flow-modules: %s", err.Error())
			return
		}
	}
	if len(_connectApplications) > 0 {
		if err := assignInputField(input, "Applications", _connectApplications); err != nil {
			log.Errorf("invalid --applications: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectGranularAccessControlConfiguration) > 0 {
		if err := assignInputField(input, "GranularAccessControlConfiguration", _connectGranularAccessControlConfiguration); err != nil {
			log.Errorf("invalid --granular-access-control-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectHierarchyRestrictedResources) > 0 {
		input.HierarchyRestrictedResources = append([]string(nil), _connectHierarchyRestrictedResources...)
	}
	if len(_connectPermissions) > 0 {
		input.Permissions = append([]string(nil), _connectPermissions...)
	}
	if len(_connectTagRestrictedResources) > 0 {
		input.TagRestrictedResources = append([]string(nil), _connectTagRestrictedResources...)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new task template in the specified Amazon Connect instance.
func connect_CreateTaskTemplate(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateTaskTemplateInput{
		// Fields: []types.TaskTemplateField, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectFields) > 0 {
		if err := assignInputField(input, "Fields", _connectFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectConstraints) > 0 {
		if err := assignInputField(input, "Constraints", _connectConstraints); err != nil {
			log.Errorf("invalid --constraints: %s", err.Error())
			return
		}
	}
	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectDefaults) > 0 {
		if err := assignInputField(input, "Defaults", _connectDefaults); err != nil {
			log.Errorf("invalid --defaults: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectSelfAssignFlowId) > 0 {
		input.SelfAssignFlowId = aws.String(_connectSelfAssignFlowId)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTaskTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a test case with its content and metadata for the specified Amazon
// Connect instance.
func connect_CreateTestCase(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateTestCaseInput{
		// Content: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectEntryPoint) > 0 {
		if err := assignInputField(input, "EntryPoint", _connectEntryPoint); err != nil {
			log.Errorf("invalid --entry-point: %s", err.Error())
			return
		}
	}
	if len(_connectInitializationData) > 0 {
		input.InitializationData = aws.String(_connectInitializationData)
	}
	if len(_connectLastModifiedRegion) > 0 {
		input.LastModifiedRegion = aws.String(_connectLastModifiedRegion)
	}
	if len(_connectLastModifiedTime) > 0 {
		if err := assignInputField(input, "LastModifiedTime", _connectLastModifiedTime); err != nil {
			log.Errorf("invalid --last-modified-time: %s", err.Error())
			return
		}
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}

	if resp, err := client.CreateTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a traffic distribution group given an Amazon Connect instance that has
// been replicated.
//
// The SignInConfig distribution is available only on a default
// TrafficDistributionGroup (see the IsDefault parameter in the [TrafficDistributionGroup] data type). If
// you call UpdateTrafficDistribution with a modified SignInConfig and a
// non-default TrafficDistributionGroup , an InvalidRequestException is returned.
//
// For more information about creating traffic distribution groups, see [Set up traffic distribution groups] in the
// Amazon Connect Administrator Guide.
//
// [Set up traffic distribution groups]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-traffic-distribution-groups.html
// [TrafficDistributionGroup]: https://docs.aws.amazon.com/connect/latest/APIReference/API_TrafficDistributionGroup.html
func connect_CreateTrafficDistributionGroup(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateTrafficDistributionGroupInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrafficDistributionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a use case for an integration association.
func connect_CreateUseCase(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateUseCaseInput{
		// InstanceId: *string, // Required
		// IntegrationAssociationId: *string, // Required
		// UseCaseType: types.UseCaseType, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectIntegrationAssociationId) > 0 {
		input.IntegrationAssociationId = aws.String(_connectIntegrationAssociationId)
	}
	if len(_connectUseCaseType) > 0 {
		if err := assignInputField(input, "UseCaseType", _connectUseCaseType); err != nil {
			log.Errorf("invalid --use-case-type: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUseCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user account for the specified Amazon Connect instance.
// Certain [UserIdentityInfo] parameters are required in some situations. For example, Email ,
// FirstName and LastName are required if you are using Amazon Connect or SAML for
// identity management.
//
// Fields in PhoneConfig cannot be set simultaneously with their corresponding
// channel-specific configuration parameters. Specifically:
//
// - PhoneConfig.AutoAccept conflicts with AutoAcceptConfigs
//
// - PhoneConfig.AfterContactWorkTimeLimit conflicts with AfterContactWorkConfigs
//
// - PhoneConfig.PhoneType and PhoneConfig.PhoneNumber conflict with
// PhoneNumberConfigs
//
// - PhoneConfig.PersistentConnection conflicts with PersistentConnectionConfigs
//
// We recommend using channel-specific parameters such as AutoAcceptConfigs ,
// AfterContactWorkConfigs , PhoneNumberConfigs , PersistentConnectionConfigs , and
// VoiceEnhancementConfigs for per-channel configuration.
//
// For information about how to create users using the Amazon Connect admin
// website, see [Add Users]in the Amazon Connect Administrator Guide.
//
// [Add Users]: https://docs.aws.amazon.com/connect/latest/adminguide/user-management.html
// [UserIdentityInfo]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UserIdentityInfo.html
func connect_CreateUser(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateUserInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
		// SecurityProfileIds: []string, // Required
		// Username: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectSecurityProfileIds) > 0 {
		input.SecurityProfileIds = append([]string(nil), _connectSecurityProfileIds...)
	}
	if len(_connectUsername) > 0 {
		input.Username = aws.String(_connectUsername)
	}
	if len(_connectAfterContactWorkConfigs) > 0 {
		if err := assignInputField(input, "AfterContactWorkConfigs", _connectAfterContactWorkConfigs); err != nil {
			log.Errorf("invalid --after-contact-work-configs: %s", err.Error())
			return
		}
	}
	if len(_connectAutoAcceptConfigs) > 0 {
		if err := assignInputField(input, "AutoAcceptConfigs", _connectAutoAcceptConfigs); err != nil {
			log.Errorf("invalid --auto-accept-configs: %s", err.Error())
			return
		}
	}
	if len(_connectDirectoryUserId) > 0 {
		input.DirectoryUserId = aws.String(_connectDirectoryUserId)
	}
	if len(_connectHierarchyGroupId) > 0 {
		input.HierarchyGroupId = aws.String(_connectHierarchyGroupId)
	}
	if len(_connectIdentityInfo) > 0 {
		if err := assignInputField(input, "IdentityInfo", _connectIdentityInfo); err != nil {
			log.Errorf("invalid --identity-info: %s", err.Error())
			return
		}
	}
	if len(_connectPassword) > 0 {
		input.Password = aws.String(_connectPassword)
	}
	if len(_connectPersistentConnectionConfigs) > 0 {
		if err := assignInputField(input, "PersistentConnectionConfigs", _connectPersistentConnectionConfigs); err != nil {
			log.Errorf("invalid --persistent-connection-configs: %s", err.Error())
			return
		}
	}
	if len(_connectPhoneConfig) > 0 {
		if err := assignInputField(input, "PhoneConfig", _connectPhoneConfig); err != nil {
			log.Errorf("invalid --phone-config: %s", err.Error())
			return
		}
	}
	if len(_connectPhoneNumberConfigs) > 0 {
		if err := assignInputField(input, "PhoneNumberConfigs", _connectPhoneNumberConfigs); err != nil {
			log.Errorf("invalid --phone-number-configs: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectVoiceEnhancementConfigs) > 0 {
		if err := assignInputField(input, "VoiceEnhancementConfigs", _connectVoiceEnhancementConfigs); err != nil {
			log.Errorf("invalid --voice-enhancement-configs: %s", err.Error())
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

// Creates a new user hierarchy group.
func connect_CreateUserHierarchyGroup(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateUserHierarchyGroupInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectParentGroupId) > 0 {
		input.ParentGroupId = aws.String(_connectParentGroupId)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserHierarchyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new view with the possible status of SAVED or PUBLISHED .
// The views will have a unique name for each connect instance.
//
// It performs basic content validation if the status is SAVED or full content
// validation if the status is set to PUBLISHED . An error is returned if
// validation fails. It associates either the $SAVED qualifier or both of the
// $SAVED and $LATEST qualifiers with the provided view content based on the
// status. The view is idempotent if ClientToken is provided.
func connect_CreateView(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateViewInput{
		// Content: *types.ViewInputContent, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// Status: types.ViewStatus, // Required
	}

	if len(_connectContent) > 0 {
		if err := assignInputField(input, "Content", _connectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a new version of the view identifier.
// Versions are immutable and monotonically increasing.
//
// It returns the highest version if there is no change in content compared to
// that version. An error is displayed if the supplied ViewContentSha256 is
// different from the ViewContentSha256 of the $LATEST alias.
func connect_CreateViewVersion(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateViewVersionInput{
		// InstanceId: *string, // Required
		// ViewId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}
	if len(_connectVersionDescription) > 0 {
		input.VersionDescription = aws.String(_connectVersionDescription)
	}
	if len(_connectViewContentSha256) > 0 {
		input.ViewContentSha256 = aws.String(_connectViewContentSha256)
	}

	if resp, err := client.CreateViewVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom vocabulary associated with your Amazon Connect instance. You
// can set a custom vocabulary to be your default vocabulary for a given language.
// Contact Lens for Amazon Connect uses the default vocabulary in post-call and
// real-time contact analysis sessions for that language.
func connect_CreateVocabulary(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateVocabularyInput{
		// Content: *string, // Required
		// InstanceId: *string, // Required
		// LanguageCode: types.VocabularyLanguageCode, // Required
		// VocabularyName: *string, // Required
	}

	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _connectLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_connectVocabularyName) > 0 {
		input.VocabularyName = aws.String(_connectVocabularyName)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a workspace that defines the user experience by mapping views to pages.
// Workspaces can be assigned to users or routing profiles.
func connect_CreateWorkspace(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateWorkspaceInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectTheme) > 0 {
		if err := assignInputField(input, "Theme", _connectTheme); err != nil {
			log.Errorf("invalid --theme: %s", err.Error())
			return
		}
	}
	if len(_connectTitle) > 0 {
		input.Title = aws.String(_connectTitle)
	}

	if resp, err := client.CreateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a view with a page in a workspace, defining what users see when they
// navigate to that page.
func connect_CreateWorkspacePage(cfg aws.Config, client *connect.Client) {
	input := &connect.CreateWorkspacePageInput{
		// InstanceId: *string, // Required
		// Page: *string, // Required
		// ResourceArn: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPage) > 0 {
		input.Page = aws.String(_connectPage)
	}
	if len(_connectResourceArn) > 0 {
		input.ResourceArn = aws.String(_connectResourceArn)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}
	if len(_connectInputData) > 0 {
		input.InputData = aws.String(_connectInputData)
	}
	if len(_connectSlug) > 0 {
		input.Slug = aws.String(_connectSlug)
	}

	if resp, err := client.CreateWorkspacePage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates an evaluation form in the specified Amazon Connect instance. After
// a form is deactivated, it is no longer available for users to start new
// evaluations based on the form.
func connect_DeactivateEvaluationForm(cfg aws.Config, client *connect.Client) {
	input := &connect.DeactivateEvaluationFormInput{
		// EvaluationFormId: *string, // Required
		// EvaluationFormVersion: int32, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectEvaluationFormVersion) > 0 {
		if err := assignInputField(input, "EvaluationFormVersion", _connectEvaluationFormVersion); err != nil {
			log.Errorf("invalid --evaluation-form-version: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeactivateEvaluationForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an attached file along with the underlying S3 Object.
// The attached file is permanently deleted if S3 bucket versioning is not enabled.
func connect_DeleteAttachedFile(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteAttachedFileInput{
		// AssociatedResourceArn: *string, // Required
		// FileId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAssociatedResourceArn) > 0 {
		input.AssociatedResourceArn = aws.String(_connectAssociatedResourceArn)
	}
	if len(_connectFileId) > 0 {
		input.FileId = aws.String(_connectFileId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteAttachedFile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a contact evaluation in the specified Amazon Connect instance.
func connect_DeleteContactEvaluation(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteContactEvaluationInput{
		// EvaluationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationId) > 0 {
		input.EvaluationId = aws.String(_connectEvaluationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteContactEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a flow for the specified Amazon Connect instance.
func connect_DeleteContactFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteContactFlowInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteContactFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified flow module.
func connect_DeleteContactFlowModule(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteContactFlowModuleInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteContactFlowModule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an alias reference, breaking the named connection to the underlying
// module version without affecting the version itself.
func connect_DeleteContactFlowModuleAlias(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteContactFlowModuleAliasInput{
		// AliasId: *string, // Required
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAliasId) > 0 {
		input.AliasId = aws.String(_connectAliasId)
	}
	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteContactFlowModuleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a specific version of a contact flow module.
func connect_DeleteContactFlowModuleVersion(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteContactFlowModuleVersionInput{
		// ContactFlowModuleId: *string, // Required
		// ContactFlowModuleVersion: *int64, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectContactFlowModuleVersion) > 0 {
		if err := assignInputField(input, "ContactFlowModuleVersion", _connectContactFlowModuleVersion); err != nil {
			log.Errorf("invalid --contact-flow-module-version: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteContactFlowModuleVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the particular version specified in flow version identifier.
func connect_DeleteContactFlowVersion(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteContactFlowVersionInput{
		// ContactFlowId: *string, // Required
		// ContactFlowVersion: *int64, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectContactFlowVersion) > 0 {
		if err := assignInputField(input, "ContactFlowVersion", _connectContactFlowVersion); err != nil {
			log.Errorf("invalid --contact-flow-version: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteContactFlowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data table and all associated attributes, versions, audits, and
// values. Does not update any references to the data table, even from other data
// tables. This includes dynamic values and conditional validations. System managed
// data tables are not deletable by customers. API users may delete the table at
// any time. When deletion is requested from the admin website, a warning is shown
// alerting the user of the most recent time the table and its values were
// accessed.
func connect_DeleteDataTable(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteDataTableInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteDataTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an attribute and all its values from a data table.
func connect_DeleteDataTableAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteDataTableAttributeInput{
		// AttributeName: *string, // Required
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAttributeName) > 0 {
		input.AttributeName = aws.String(_connectAttributeName)
	}
	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteDataTableAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes email address from the specified Amazon Connect instance.
func connect_DeleteEmailAddress(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteEmailAddressInput{
		// EmailAddressId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEmailAddressId) > 0 {
		input.EmailAddressId = aws.String(_connectEmailAddressId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteEmailAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an evaluation form in the specified Amazon Connect instance.
// - If the version property is provided, only the specified version of the
// evaluation form is deleted.
//
// - If no version is provided, then the full form (all versions) is deleted.
func connect_DeleteEvaluationForm(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteEvaluationFormInput{
		// EvaluationFormId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectEvaluationFormVersion) > 0 {
		if err := assignInputField(input, "EvaluationFormVersion", _connectEvaluationFormVersion); err != nil {
			log.Errorf("invalid --evaluation-form-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteEvaluationForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an hours of operation.
func connect_DeleteHoursOfOperation(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteHoursOfOperationInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteHoursOfOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an hours of operation override in an Amazon Connect hours of operation
// resource.
func connect_DeleteHoursOfOperationOverride(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteHoursOfOperationOverrideInput{
		// HoursOfOperationId: *string, // Required
		// HoursOfOperationOverrideId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectHoursOfOperationOverrideId) > 0 {
		input.HoursOfOperationOverrideId = aws.String(_connectHoursOfOperationOverrideId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteHoursOfOperationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Deletes the Amazon Connect instance. For more information, see [Delete your Amazon Connect instance] in the Amazon
// Connect Administrator Guide.
//
// Amazon Connect enforces a limit on the total number of instances that you can
// create or delete in 30 days. If you exceed this limit, you will get an error
// message indicating there has been an excessive number of attempts at creating or
// deleting instances. You must wait 30 days before you can restart creating and
// deleting instances in your account.
//
// [Delete your Amazon Connect instance]: https://docs.aws.amazon.com/connect/latest/adminguide/delete-connect-instance.html
func connect_DeleteInstance(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteInstanceInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DeleteInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Web Services resource association from an Amazon Connect
// instance. The association must not have any use cases associated with it.
func connect_DeleteIntegrationAssociation(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteIntegrationAssociationInput{
		// InstanceId: *string, // Required
		// IntegrationAssociationId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectIntegrationAssociationId) > 0 {
		input.IntegrationAssociationId = aws.String(_connectIntegrationAssociationId)
	}

	if resp, err := client.DeleteIntegrationAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a notification. Once deleted, the notification is no longer visible to
// all users and cannot be managed through the Admin Website or APIs.
func connect_DeleteNotification(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteNotificationInput{
		// InstanceId: *string, // Required
		// NotificationId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectNotificationId) > 0 {
		input.NotificationId = aws.String(_connectNotificationId)
	}

	if resp, err := client.DeleteNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a predefined attribute from the specified Amazon Connect instance.
func connect_DeletePredefinedAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.DeletePredefinedAttributeInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.DeletePredefinedAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a prompt.
func connect_DeletePrompt(cfg aws.Config, client *connect.Client) {
	input := &connect.DeletePromptInput{
		// InstanceId: *string, // Required
		// PromptId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPromptId) > 0 {
		input.PromptId = aws.String(_connectPromptId)
	}

	if resp, err := client.DeletePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes registration for a device token and a chat contact.
func connect_DeletePushNotificationRegistration(cfg aws.Config, client *connect.Client) {
	input := &connect.DeletePushNotificationRegistrationInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// RegistrationId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRegistrationId) > 0 {
		input.RegistrationId = aws.String(_connectRegistrationId)
	}

	if resp, err := client.DeletePushNotificationRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a queue.
func connect_DeleteQueue(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteQueueInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}

	if resp, err := client.DeleteQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a quick connect.
// After calling [DeleteUser], it's important to call DeleteQuickConnect to delete any records
// related to the deleted users. This will help you:
//
// - Avoid dangling resources that impact your service quotas.
//
// - Remove deleted users so they don't appear to agents as transfer options.
//
// - Avoid the disruption of other Amazon Connect processes, such as instance
// replication and syncing if you're using [Amazon Connect Global Resiliency].
//
// [Amazon Connect Global Resiliency]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-connect-global-resiliency.html
// [DeleteUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteUser.html
func connect_DeleteQuickConnect(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteQuickConnectInput{
		// InstanceId: *string, // Required
		// QuickConnectId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQuickConnectId) > 0 {
		input.QuickConnectId = aws.String(_connectQuickConnectId)
	}

	if resp, err := client.DeleteQuickConnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a routing profile.
func connect_DeleteRoutingProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteRoutingProfileInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}

	if resp, err := client.DeleteRoutingProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a rule for the specified Amazon Connect instance.
func connect_DeleteRule(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteRuleInput{
		// InstanceId: *string, // Required
		// RuleId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRuleId) > 0 {
		input.RuleId = aws.String(_connectRuleId)
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a security profile.
func connect_DeleteSecurityProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteSecurityProfileInput{
		// InstanceId: *string, // Required
		// SecurityProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileId) > 0 {
		input.SecurityProfileId = aws.String(_connectSecurityProfileId)
	}

	if resp, err := client.DeleteSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the task template.
func connect_DeleteTaskTemplate(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteTaskTemplateInput{
		// InstanceId: *string, // Required
		// TaskTemplateId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTaskTemplateId) > 0 {
		input.TaskTemplateId = aws.String(_connectTaskTemplateId)
	}

	if resp, err := client.DeleteTaskTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the test case that has already been created for the specified Amazon
// Connect instance.
func connect_DeleteTestCase(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteTestCaseInput{
		// InstanceId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}

	if resp, err := client.DeleteTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a traffic distribution group. This API can be called only in the Region
// where the traffic distribution group is created.
//
// For more information about deleting traffic distribution groups, see [Delete traffic distribution groups] in the
// Amazon Connect Administrator Guide.
//
// [Delete traffic distribution groups]: https://docs.aws.amazon.com/connect/latest/adminguide/delete-traffic-distribution-groups.html
func connect_DeleteTrafficDistributionGroup(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteTrafficDistributionGroupInput{
		// TrafficDistributionGroupId: *string, // Required
	}

	if len(_connectTrafficDistributionGroupId) > 0 {
		input.TrafficDistributionGroupId = aws.String(_connectTrafficDistributionGroupId)
	}

	if resp, err := client.DeleteTrafficDistributionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a use case from an integration association.
func connect_DeleteUseCase(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteUseCaseInput{
		// InstanceId: *string, // Required
		// IntegrationAssociationId: *string, // Required
		// UseCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectIntegrationAssociationId) > 0 {
		input.IntegrationAssociationId = aws.String(_connectIntegrationAssociationId)
	}
	if len(_connectUseCaseId) > 0 {
		input.UseCaseId = aws.String(_connectUseCaseId)
	}

	if resp, err := client.DeleteUseCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user account from the specified Amazon Connect instance.
// For information about what happens to a user's data when their account is
// deleted, see [Delete Users from Your Amazon Connect Instance]in the Amazon Connect Administrator Guide.
//
// After calling DeleteUser, call [DeleteQuickConnect] to delete any records related to the deleted
// users. This will help you:
//
// - Avoid dangling resources that impact your service quotas.
//
// - Remove deleted users so they don't appear to agents as transfer options.
//
// - Avoid the disruption of other Amazon Connect processes, such as instance
// replication and syncing if you're using [Amazon Connect Global Resiliency].
//
// [Delete Users from Your Amazon Connect Instance]: https://docs.aws.amazon.com/connect/latest/adminguide/delete-users.html
// [Amazon Connect Global Resiliency]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-connect-global-resiliency.html
// [DeleteQuickConnect]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteQuickConnect.html
func connect_DeleteUser(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteUserInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing user hierarchy group. It must not be associated with any
// agents or have any active child groups.
func connect_DeleteUserHierarchyGroup(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteUserHierarchyGroupInput{
		// HierarchyGroupId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHierarchyGroupId) > 0 {
		input.HierarchyGroupId = aws.String(_connectHierarchyGroupId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DeleteUserHierarchyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the view entirely. It deletes the view and all associated qualifiers
// (versions and aliases).
func connect_DeleteView(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteViewInput{
		// InstanceId: *string, // Required
		// ViewId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}

	if resp, err := client.DeleteView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the particular version specified in ViewVersion identifier.
func connect_DeleteViewVersion(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteViewVersionInput{
		// InstanceId: *string, // Required
		// ViewId: *string, // Required
		// ViewVersion: *int32, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}
	if len(_connectViewVersion) > 0 {
		if err := assignInputField(input, "ViewVersion", _connectViewVersion); err != nil {
			log.Errorf("invalid --view-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteViewVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the vocabulary that has the given identifier.
func connect_DeleteVocabulary(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteVocabularyInput{
		// InstanceId: *string, // Required
		// VocabularyId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectVocabularyId) > 0 {
		input.VocabularyId = aws.String(_connectVocabularyId)
	}

	if resp, err := client.DeleteVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workspace and removes all associated view and resource assignments.
func connect_DeleteWorkspace(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteWorkspaceInput{
		// InstanceId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.DeleteWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a media asset (such as a logo) from a workspace.
func connect_DeleteWorkspaceMedia(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteWorkspaceMediaInput{
		// InstanceId: *string, // Required
		// MediaType: types.MediaType, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMediaType) > 0 {
		if err := assignInputField(input, "MediaType", _connectMediaType); err != nil {
			log.Errorf("invalid --media-type: %s", err.Error())
			return
		}
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.DeleteWorkspaceMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a view and a page in a workspace. The page will
// display the default view after deletion.
func connect_DeleteWorkspacePage(cfg aws.Config, client *connect.Client) {
	input := &connect.DeleteWorkspacePageInput{
		// InstanceId: *string, // Required
		// Page: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPage) > 0 {
		input.Page = aws.String(_connectPage)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.DeleteWorkspacePage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an agent status.
func connect_DescribeAgentStatus(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeAgentStatusInput{
		// AgentStatusId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAgentStatusId) > 0 {
		input.AgentStatusId = aws.String(_connectAgentStatusId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeAgentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change. To
// request access to this API, contact Amazon Web Services Support.
//
// Describes the target authentication profile.
func connect_DescribeAuthenticationProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeAuthenticationProfileInput{
		// AuthenticationProfileId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAuthenticationProfileId) > 0 {
		input.AuthenticationProfileId = aws.String(_connectAuthenticationProfileId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeAuthenticationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Describes the specified contact.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Retrieve contact information such as the caller's phone number and the
// specific number the caller dialed to integrate into custom monitoring or custom
// agent experience solutions.
//
// - Detect when a customer chat session disconnects due to a network issue on
// the agent's end. Use the DisconnectReason field in the [ContactTraceRecord]to detect this event
// and then re-queue the chat for followup.
//
// - Identify after contact work (ACW) duration and call recordings information
// when a COMPLETED event is received by using the [contact event stream].
//
// # Important things to know
//
// - SystemEndpoint is not populated for contacts with initiation method of
// MONITOR, QUEUE_TRANSFER, or CALLBACK
//
// - Contact information remains available in Amazon Connect for 24 months from
// the InitiationTimestamp , and then it is deleted. Only contact information
// that is available in Amazon Connect is returned by this API.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [ContactTraceRecord]: https://docs.aws.amazon.com/connect/latest/adminguide/ctr-data-model.html#ctr-ContactTraceRecord
// [contact event stream]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-events.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
func connect_DescribeContact(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a contact evaluation in the specified Amazon Connect instance.
func connect_DescribeContactEvaluation(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeContactEvaluationInput{
		// EvaluationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationId) > 0 {
		input.EvaluationId = aws.String(_connectEvaluationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeContactEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified flow.
// You can also create and update flows using the [Amazon Connect Flow language].
//
// Use the $SAVED alias in the request to describe the SAVED content of a Flow.
// For example, arn:aws:.../contact-flow/{id}:$SAVED . After a flow is published,
// $SAVED needs to be supplied to view saved content that has not been published.
//
// Use arn:aws:.../contact-flow/{id}:{version} to retrieve the content of a
// specific flow version.
//
// In the response, Status indicates the flow status as either SAVED or PUBLISHED .
// The PUBLISHED status will initiate validation on the content. SAVED does not
// initiate validation of the content. SAVED | PUBLISHED
//
// [Amazon Connect Flow language]: https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html
func connect_DescribeContactFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeContactFlowInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeContactFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified flow module.
// Use the $SAVED alias in the request to describe the SAVED content of a Flow.
// For example, arn:aws:.../contact-flow/{id}:$SAVED . After a flow is published,
// $SAVED needs to be supplied to view saved content that has not been published.
func connect_DescribeContactFlowModule(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeContactFlowModuleInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeContactFlowModule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific alias, including which version
// it currently points to and its metadata.
func connect_DescribeContactFlowModuleAlias(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeContactFlowModuleAliasInput{
		// AliasId: *string, // Required
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAliasId) > 0 {
		input.AliasId = aws.String(_connectAliasId)
	}
	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeContactFlowModuleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all properties for a data table except for attributes and values. All
// properties from CreateDataTable are returned as well as properties for region
// replication, versioning, and system tables. "Describe" is a deprecated term but
// is allowed to maintain consistency with existing operations.
func connect_DescribeDataTable(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeDataTableInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeDataTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information for a specific data table attribute including its
// configuration, validation rules, and metadata. "Describe" is a deprecated term
// but is allowed to maintain consistency with existing operations.
func connect_DescribeDataTableAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeDataTableAttributeInput{
		// AttributeName: *string, // Required
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAttributeName) > 0 {
		input.AttributeName = aws.String(_connectAttributeName)
	}
	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeDataTableAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe email address form the specified Amazon Connect instance.
func connect_DescribeEmailAddress(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeEmailAddressInput{
		// EmailAddressId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEmailAddressId) > 0 {
		input.EmailAddressId = aws.String(_connectEmailAddressId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeEmailAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an evaluation form in the specified Amazon Connect instance. If the
// version property is not provided, the latest version of the evaluation form is
// described.
func connect_DescribeEvaluationForm(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeEvaluationFormInput{
		// EvaluationFormId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectEvaluationFormVersion) > 0 {
		if err := assignInputField(input, "EvaluationFormVersion", _connectEvaluationFormVersion); err != nil {
			log.Errorf("invalid --evaluation-form-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeEvaluationForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the hours of operation.
func connect_DescribeHoursOfOperation(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeHoursOfOperationInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeHoursOfOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the hours of operation override.
func connect_DescribeHoursOfOperationOverride(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeHoursOfOperationOverrideInput{
		// HoursOfOperationId: *string, // Required
		// HoursOfOperationOverrideId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectHoursOfOperationOverrideId) > 0 {
		input.HoursOfOperationOverrideId = aws.String(_connectHoursOfOperationOverrideId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeHoursOfOperationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Returns the current state of the specified instance identifier. It tracks the
// instance while it is being created and returns an error status, if applicable.
//
// If an instance is not created successfully, the instance status reason field
// returns details relevant to the reason. The instance in a failed state is
// returned only for 24 hours after the CreateInstance API was invoked.
func connect_DescribeInstance(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeInstanceInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Describes the specified instance attribute.
func connect_DescribeInstanceAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeInstanceAttributeInput{
		// AttributeType: types.InstanceAttributeType, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAttributeType) > 0 {
		if err := assignInputField(input, "AttributeType", _connectAttributeType); err != nil {
			log.Errorf("invalid --attribute-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeInstanceAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Retrieves the current storage configurations for the specified resource type,
// association ID, and instance ID.
func connect_DescribeInstanceStorageConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeInstanceStorageConfigInput{
		// AssociationId: *string, // Required
		// InstanceId: *string, // Required
		// ResourceType: types.InstanceStorageResourceType, // Required
	}

	if len(_connectAssociationId) > 0 {
		input.AssociationId = aws.String(_connectAssociationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeInstanceStorageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific notification, including its
// content, priority, recipients, and metadata.
func connect_DescribeNotification(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeNotificationInput{
		// InstanceId: *string, // Required
		// NotificationId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectNotificationId) > 0 {
		input.NotificationId = aws.String(_connectNotificationId)
	}

	if resp, err := client.DescribeNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details and status of a phone number that’s claimed to your Amazon Connect
// instance or traffic distribution group.
//
// If the number is claimed to a traffic distribution group, and you are calling
// in the Amazon Web Services Region where the traffic distribution group was
// created, you can use either a phone number ARN or UUID value for the
// PhoneNumberId URI request parameter. However, if the number is claimed to a
// traffic distribution group and you are calling this API in the alternate Amazon
// Web Services Region associated with the traffic distribution group, you must
// provide a full phone number ARN. If a UUID is provided in this scenario, you
// receive a ResourceNotFoundException .
func connect_DescribePhoneNumber(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_connectPhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_connectPhoneNumberId)
	}

	if resp, err := client.DescribePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a predefined attribute for the specified Amazon Connect instance. A
// predefined attribute is made up of a name and a value. You can use predefined
// attributes for:
//
// - Routing proficiency (for example, agent certification) that has predefined
// values (for example, a list of possible certifications). For more information,
// see [Create predefined attributes for routing contacts to agents].
//
// - Contact information that varies between transfers or conferences, such as
// the name of the business unit handling the contact. For more information, see [Use contact segment attributes]
// .
//
// For the predefined attributes per instance quota, see [Amazon Connect quotas].
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Use contact segment attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/use-contact-segment-attributes.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
func connect_DescribePredefinedAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribePredefinedAttributeInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.DescribePredefinedAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the prompt.
func connect_DescribePrompt(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribePromptInput{
		// InstanceId: *string, // Required
		// PromptId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPromptId) > 0 {
		input.PromptId = aws.String(_connectPromptId)
	}

	if resp, err := client.DescribePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified queue.
func connect_DescribeQueue(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeQueueInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}

	if resp, err := client.DescribeQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the quick connect.
func connect_DescribeQuickConnect(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeQuickConnectInput{
		// InstanceId: *string, // Required
		// QuickConnectId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQuickConnectId) > 0 {
		input.QuickConnectId = aws.String(_connectQuickConnectId)
	}

	if resp, err := client.DescribeQuickConnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified routing profile.
// DescribeRoutingProfile does not populate AssociatedQueueIds in its response.
// The example Response Syntax shown on this page is incorrect; we are working to
// update it. [SearchRoutingProfiles]does include AssociatedQueueIds.
//
// [SearchRoutingProfiles]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchRoutingProfiles.html
func connect_DescribeRoutingProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeRoutingProfileInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}

	if resp, err := client.DescribeRoutingProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a rule for the specified Amazon Connect instance.
func connect_DescribeRule(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeRuleInput{
		// InstanceId: *string, // Required
		// RuleId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRuleId) > 0 {
		input.RuleId = aws.String(_connectRuleId)
	}

	if resp, err := client.DescribeRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets basic information about the security profile.
// For information about security profiles, see [Security Profiles] in the Amazon Connect
// Administrator Guide. For a mapping of the API name and user interface name of
// the security profile permissions, see [List of security profile permissions].
//
// [Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html
// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
func connect_DescribeSecurityProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeSecurityProfileInput{
		// InstanceId: *string, // Required
		// SecurityProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileId) > 0 {
		input.SecurityProfileId = aws.String(_connectSecurityProfileId)
	}

	if resp, err := client.DescribeSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified test case and allows you to get the content and
// metadata of the test case for the specified Amazon Connect instance.
func connect_DescribeTestCase(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeTestCaseInput{
		// InstanceId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details and status of a traffic distribution group.
func connect_DescribeTrafficDistributionGroup(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeTrafficDistributionGroupInput{
		// TrafficDistributionGroupId: *string, // Required
	}

	if len(_connectTrafficDistributionGroupId) > 0 {
		input.TrafficDistributionGroupId = aws.String(_connectTrafficDistributionGroupId)
	}

	if resp, err := client.DescribeTrafficDistributionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified user. You can [find the instance ID in the Amazon Connect console] (it’s the final part of the ARN). The
// console does not display the user IDs. Instead, list the users and note the IDs
// provided in the output.
//
// [find the instance ID in the Amazon Connect console]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
func connect_DescribeUser(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeUserInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified hierarchy group.
func connect_DescribeUserHierarchyGroup(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeUserHierarchyGroupInput{
		// HierarchyGroupId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHierarchyGroupId) > 0 {
		input.HierarchyGroupId = aws.String(_connectHierarchyGroupId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeUserHierarchyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the hierarchy structure of the specified Amazon Connect instance.
func connect_DescribeUserHierarchyStructure(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeUserHierarchyStructureInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.DescribeUserHierarchyStructure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the view for the specified Amazon Connect instance and view
// identifier.
//
// The view identifier can be supplied as a ViewId or ARN.
//
// $SAVED needs to be supplied if a view is unpublished.
//
// The view identifier can contain an optional qualifier, for example, :$SAVED ,
// which is either an actual version number or an Amazon Connect managed qualifier
// $SAVED | $LATEST . If it is not supplied, then $LATEST is assumed for customer
// managed views and an error is returned if there is no published content
// available. Version 1 is assumed for Amazon Web Services managed views.
func connect_DescribeView(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeViewInput{
		// InstanceId: *string, // Required
		// ViewId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}

	if resp, err := client.DescribeView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified vocabulary.
func connect_DescribeVocabulary(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeVocabularyInput{
		// InstanceId: *string, // Required
		// VocabularyId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectVocabularyId) > 0 {
		input.VocabularyId = aws.String(_connectVocabularyId)
	}

	if resp, err := client.DescribeVocabulary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a workspace, including its configuration and metadata.
func connect_DescribeWorkspace(cfg aws.Config, client *connect.Client) {
	input := &connect.DescribeWorkspaceInput{
		// InstanceId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.DescribeWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the dataset ID associated with a given Amazon Connect instance.
func connect_DisassociateAnalyticsDataSet(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateAnalyticsDataSetInput{
		// DataSetId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataSetId) > 0 {
		input.DataSetId = aws.String(_connectDataSetId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTargetAccountId) > 0 {
		input.TargetAccountId = aws.String(_connectTargetAccountId)
	}

	if resp, err := client.DisassociateAnalyticsDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Revokes access to integrated applications from Amazon Connect.
func connect_DisassociateApprovedOrigin(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateApprovedOriginInput{
		// InstanceId: *string, // Required
		// Origin: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectOrigin) > 0 {
		input.Origin = aws.String(_connectOrigin)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DisassociateApprovedOrigin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Revokes authorization from the specified instance to access the specified
// Amazon Lex or Amazon Lex V2 bot.
func connect_DisassociateBot(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateBotInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectLexBot) > 0 {
		if err := assignInputField(input, "LexBot", _connectLexBot); err != nil {
			log.Errorf("invalid --lex-bot: %s", err.Error())
			return
		}
	}
	if len(_connectLexV2Bot) > 0 {
		if err := assignInputField(input, "LexV2Bot", _connectLexV2Bot); err != nil {
			log.Errorf("invalid --lex-v2-bot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the alias association between two email addresses in an Amazon Connect
// instance. After disassociation, emails sent to the former alias email address
// are no longer forwarded to the primary email address. Both email addresses
// continue to exist independently and can receive emails directly.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Department separation: Remove alias relationships when splitting a
// consolidated support queue back into separate department-specific queues.
//
// - Email address retirement: Cleanly remove forwarding relationships before
// decommissioning old email addresses.
//
// - Organizational restructuring: Reconfigure email routing when business
// processes change and aliases are no longer needed.
//
// # Important things to know
//
// - Concurrent operations: This API uses distributed locking, so concurrent
// operations on the same email addresses may be temporarily blocked.
//
// - Emails sent to the former alias address are still delivered directly to
// that address if it exists.
//
// - You do not need to delete the email addresses after disassociation. Both
// addresses remain active independently.
//
// - After a successful disassociation, you can immediately create a new alias
// relationship with the same addresses.
//
// - 200 status means alias was successfully disassociated.
//
// DisassociateEmailAddressAlias does not return the following information:
//
// - Details in the response about the email that was disassociated. The
// response returns an empty body.
//
// - The timestamp of when the disassociation occurred.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// # Related operations
//
// [AssociateEmailAddressAlias]
// - : Associates an email address alias with an existing email address in an
// Amazon Connect instance.
//
// [DescribeEmailAddress]
// - : View current alias configurations for an email address.
//
// [SearchEmailAddresses]
// - : Find email addresses and their alias relationships across an instance.
//
// [CreateEmailAddress]
// - : Create new email addresses that can participate in alias relationships.
//
// [DeleteEmailAddress]
// - : Remove email addresses (automatically removes any alias relationships).
//
// [UpdateEmailAddressMetadata]
// - : Modify email address properties (does not affect alias relationships).
//
// [DescribeEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeEmailAddress.html
// [DeleteEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteEmailAddress.html
// [SearchEmailAddresses]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchEmailAddresses.html
// [UpdateEmailAddressMetadata]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdateEmailAddressMetadata.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [CreateEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateEmailAddress.html
// [AssociateEmailAddressAlias]: https://docs.aws.amazon.com/connect/latest/APIReference/API_AssociateEmailAddressAlias.html
func connect_DisassociateEmailAddressAlias(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateEmailAddressAliasInput{
		// AliasConfiguration: *types.AliasConfiguration, // Required
		// EmailAddressId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAliasConfiguration) > 0 {
		if err := assignInputField(input, "AliasConfiguration", _connectAliasConfiguration); err != nil {
			log.Errorf("invalid --alias-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectEmailAddressId) > 0 {
		input.EmailAddressId = aws.String(_connectEmailAddressId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DisassociateEmailAddressAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a connect resource from a flow.
func connect_DisassociateFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateFlowInput{
		// InstanceId: *string, // Required
		// ResourceId: *string, // Required
		// ResourceType: types.FlowAssociationResourceType, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceId) > 0 {
		input.ResourceId = aws.String(_connectResourceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a set of hours of operations with another hours of operation.
// Refer to Administrator Guide [here]for more information on inheriting overrides from
// parent hours of operation(s).
//
// [here]: https://docs.aws.amazon.com/connect/latest/adminguide/hours-of-operation-overrides.html
func connect_DisassociateHoursOfOperations(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateHoursOfOperationsInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
		// ParentHoursOfOperationIds: []string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectParentHoursOfOperationIds) > 0 {
		input.ParentHoursOfOperationIds = append([]string(nil), _connectParentHoursOfOperationIds...)
	}

	if resp, err := client.DisassociateHoursOfOperations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Removes the storage type configurations for the specified resource type and
// association ID.
func connect_DisassociateInstanceStorageConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateInstanceStorageConfigInput{
		// AssociationId: *string, // Required
		// InstanceId: *string, // Required
		// ResourceType: types.InstanceStorageResourceType, // Required
	}

	if len(_connectAssociationId) > 0 {
		input.AssociationId = aws.String(_connectAssociationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DisassociateInstanceStorageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Remove the Lambda function from the dropdown options available in the relevant
// flow blocks.
func connect_DisassociateLambdaFunction(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateLambdaFunctionInput{
		// FunctionArn: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectFunctionArn) > 0 {
		input.FunctionArn = aws.String(_connectFunctionArn)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DisassociateLambdaFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Revokes authorization from the specified instance to access the specified
// Amazon Lex bot.
func connect_DisassociateLexBot(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateLexBotInput{
		// BotName: *string, // Required
		// InstanceId: *string, // Required
		// LexRegion: *string, // Required
	}

	if len(_connectBotName) > 0 {
		input.BotName = aws.String(_connectBotName)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLexRegion) > 0 {
		input.LexRegion = aws.String(_connectLexRegion)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DisassociateLexBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the flow association from a phone number claimed to your Amazon Connect
// instance.
//
// If the number is claimed to a traffic distribution group, and you are calling
// this API using an instance in the Amazon Web Services Region where the traffic
// distribution group was created, you can use either a full phone number ARN or
// UUID value for the PhoneNumberId URI request parameter. However, if the number
// is claimed to a traffic distribution group and you are calling this API using an
// instance in the alternate Amazon Web Services Region associated with the traffic
// distribution group, you must provide a full phone number ARN. If a UUID is
// provided in this scenario, you will receive a ResourceNotFoundException .
func connect_DisassociatePhoneNumberContactFlow(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociatePhoneNumberContactFlowInput{
		// InstanceId: *string, // Required
		// PhoneNumberId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_connectPhoneNumberId)
	}

	if resp, err := client.DisassociatePhoneNumberContactFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a set of quick connects from a queue.
func connect_DisassociateQueueQuickConnects(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateQueueQuickConnectsInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
		// QuickConnectIds: []string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectQuickConnectIds) > 0 {
		input.QuickConnectIds = append([]string(nil), _connectQuickConnectIds...)
	}

	if resp, err := client.DisassociateQueueQuickConnects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a set of queues from a routing profile.
// Up to 10 queue references can be disassociated in a single API call. More than
// 10 queue references results in a single call results in an
// InvalidParameterException.
func connect_DisassociateRoutingProfileQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateRoutingProfileQueuesInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectManualAssignmentQueueReferences) > 0 {
		if err := assignInputField(input, "ManualAssignmentQueueReferences", _connectManualAssignmentQueueReferences); err != nil {
			log.Errorf("invalid --manual-assignment-queue-references: %s", err.Error())
			return
		}
	}
	if len(_connectQueueReferences) > 0 {
		if err := assignInputField(input, "QueueReferences", _connectQueueReferences); err != nil {
			log.Errorf("invalid --queue-references: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateRoutingProfileQueues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Deletes the specified security key.
func connect_DisassociateSecurityKey(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateSecurityKeyInput{
		// AssociationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAssociationId) > 0 {
		input.AssociationId = aws.String(_connectAssociationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.DisassociateSecurityKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a security profile attached to a Q in Connect AI Agent Entity in
// an Amazon Connect instance.
func connect_DisassociateSecurityProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateSecurityProfilesInput{
		// EntityArn: *string, // Required
		// EntityType: types.EntityType, // Required
		// InstanceId: *string, // Required
		// SecurityProfiles: []types.SecurityProfileItem, // Required
	}

	if len(_connectEntityArn) > 0 {
		input.EntityArn = aws.String(_connectEntityArn)
	}
	if len(_connectEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _connectEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfiles) > 0 {
		if err := assignInputField(input, "SecurityProfiles", _connectSecurityProfiles); err != nil {
			log.Errorf("invalid --security-profiles: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateSecurityProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an agent from a traffic distribution group. This API can be
// called only in the Region where the traffic distribution group is created.
func connect_DisassociateTrafficDistributionGroupUser(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateTrafficDistributionGroupUserInput{
		// InstanceId: *string, // Required
		// TrafficDistributionGroupId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTrafficDistributionGroupId) > 0 {
		input.TrafficDistributionGroupId = aws.String(_connectTrafficDistributionGroupId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.DisassociateTrafficDistributionGroupUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a set of proficiencies from a user.
func connect_DisassociateUserProficiencies(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateUserProficienciesInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
		// UserProficiencies: []types.UserProficiencyDisassociate, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectUserProficiencies) > 0 {
		if err := assignInputField(input, "UserProficiencies", _connectUserProficiencies); err != nil {
			log.Errorf("invalid --user-proficiencies: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateUserProficiencies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a workspace and one or more users or routing
// profiles.
func connect_DisassociateWorkspace(cfg aws.Config, client *connect.Client) {
	input := &connect.DisassociateWorkspaceInput{
		// InstanceId: *string, // Required
		// ResourceArns: []string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _connectResourceArns...)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.DisassociateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dismisses contacts from an agent’s CCP and returns the agent to an available
// state, which allows the agent to receive a new routed contact. Contacts can only
// be dismissed if they are in a MISSED , ERROR , ENDED , or REJECTED state in the [Agent Event Stream]
// .
//
// [Agent Event Stream]: https://docs.aws.amazon.com/connect/latest/adminguide/about-contact-states.html
func connect_DismissUserContact(cfg aws.Config, client *connect.Client) {
	input := &connect.DismissUserContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.DismissUserContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates values at the time of the request and returns them. It considers the
// request's timezone or the table's timezone, in that order, when accessing time
// based tables. When a value is accessed, the accessor's identity and the time of
// access are saved alongside the value to help identify values that are actively
// in use. The term "Batch" is not included in the operation name since it does not
// meet all the criteria for a batch operation as specified in Batch Operations:
// Amazon Web Services API Standards.
func connect_EvaluateDataTableValues(cfg aws.Config, client *connect.Client) {
	input := &connect.EvaluateDataTableValuesInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Values: []types.DataTableValueEvaluationSet, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectTimeZone) > 0 {
		input.TimeZone = aws.String(_connectTimeZone)
	}

	if disablePaginator() {
		if resp, err := client.EvaluateDataTableValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.EvaluateDataTableValuesOutput
	p := connect.NewEvaluateDataTableValuesPaginator(client, input)
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

// Provides a pre-signed URL for download of an approved attached file. This API
// also returns metadata about the attached file. It will only return a downloadURL
// if the status of the attached file is APPROVED .
func connect_GetAttachedFile(cfg aws.Config, client *connect.Client) {
	input := &connect.GetAttachedFileInput{
		// AssociatedResourceArn: *string, // Required
		// FileId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAssociatedResourceArn) > 0 {
		input.AssociatedResourceArn = aws.String(_connectAssociatedResourceArn)
	}
	if len(_connectFileId) > 0 {
		input.FileId = aws.String(_connectFileId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUrlExpiryInSeconds) > 0 {
		if err := assignInputField(input, "UrlExpiryInSeconds", _connectUrlExpiryInSeconds); err != nil {
			log.Errorf("invalid --url-expiry-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAttachedFile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the contact attributes for the specified contact.
func connect_GetContactAttributes(cfg aws.Config, client *connect.Client) {
	input := &connect.GetContactAttributesInput{
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.GetContactAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves contact metric data for a specified contact.
// # Use cases
//
// Following are common use cases for position in queue and estimated wait time:
//
// - Customer-Facing Wait Time Announcements - Display or announce the estimated
// wait time and position in queue to customers before or during their queue
// experience.
//
// - Callback Offerings - Offer customers a callback option when the estimated
// wait time or position in queue exceeds a defined threshold.
//
// - Queue Routing Decisions - Route incoming contacts to less congested queues
// by comparing estimated wait time and position in queue across multiple queues.
//
// - Self-Service Deflection - Redirect customers to self-service options like
// chatbots or FAQs when estimated wait time is high or position in queue is
// unfavorable.
//
// # Important things to know
//
// - Metrics are only available while the contact is actively in queue.
//
// - For more information, see the [Position in queue]metric in the Amazon Connect Administrator
// Guide.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Position in queue]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
func connect_GetContactMetrics(cfg aws.Config, client *connect.Client) {
	input := &connect.GetContactMetricsInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// Metrics: []types.ContactMetricInfo, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMetrics) > 0 {
		if err := assignInputField(input, "Metrics", _connectMetrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetContactMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the real-time metric data from the specified Amazon Connect instance.
// For a description of each metric, see [Metrics definitions] in the Amazon Connect Administrator
// Guide.
//
// When you make a successful API request, you can expect the following metric
// values in the response:
//
// - Metric value is null: The calculation cannot be performed due to divide by
// zero or insufficient data
//
// - Metric value is a number (including 0) of defined type: The number provided
// is the calculation result
//
// - MetricResult list is empty: The request cannot find any data in the system
//
// The following guidelines can help you work with the API:
//
// - Each dimension in the metric response must contain a value
//
// - Each item in MetricResult must include all requested metrics
//
// - If the response is slow due to large result sets, try these approaches:
//
// - Add filters to reduce the amount of data returned
//
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
func connect_GetCurrentMetricData(cfg aws.Config, client *connect.Client) {
	input := &connect.GetCurrentMetricDataInput{
		// CurrentMetrics: []types.CurrentMetric, // Required
		// Filters: *types.Filters, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectCurrentMetrics) > 0 {
		if err := assignInputField(input, "CurrentMetrics", _connectCurrentMetrics); err != nil {
			log.Errorf("invalid --current-metrics: %s", err.Error())
			return
		}
	}
	if len(_connectFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectGroupings) > 0 {
		if err := assignInputField(input, "Groupings", _connectGroupings[0]); err != nil {
			log.Errorf("invalid --groupings: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _connectSortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetCurrentMetricData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.GetCurrentMetricDataOutput
	p := connect.NewGetCurrentMetricDataPaginator(client, input)
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

// Gets the real-time active user data from the specified Amazon Connect instance.
func connect_GetCurrentUserData(cfg aws.Config, client *connect.Client) {
	input := &connect.GetCurrentUserDataInput{
		// Filters: *types.UserDataFilters, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCurrentUserData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.GetCurrentUserDataOutput
	p := connect.NewGetCurrentUserDataPaginator(client, input)
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

// Get the hours of operations with the effective override applied.
func connect_GetEffectiveHoursOfOperations(cfg aws.Config, client *connect.Client) {
	input := &connect.GetEffectiveHoursOfOperationsInput{
		// FromDate: *string, // Required
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
		// ToDate: *string, // Required
	}

	if len(_connectFromDate) > 0 {
		input.FromDate = aws.String(_connectFromDate)
	}
	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectToDate) > 0 {
		input.ToDate = aws.String(_connectToDate)
	}

	if resp, err := client.GetEffectiveHoursOfOperations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Supports SAML sign-in for Amazon Connect. Retrieves a token for federation. The
// token is for the Amazon Connect user which corresponds to the IAM credentials
// that were used to invoke this action.
//
// For more information about how SAML sign-in works in Amazon Connect, see [Configure SAML with IAM for Amazon Connect in the Amazon Connect Administrator Guide.]
//
// This API doesn't support root users. If you try to invoke GetFederationToken
// with root credentials, an error message similar to the following one appears:
//
// Provided identity: Principal: .... User: .... cannot be used for federation
// with Amazon Connect
//
// [Configure SAML with IAM for Amazon Connect in the Amazon Connect Administrator Guide.]: https://docs.aws.amazon.com/connect/latest/adminguide/configure-saml.html
func connect_GetFederationToken(cfg aws.Config, client *connect.Client) {
	input := &connect.GetFederationTokenInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.GetFederationToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the flow associated for a given resource.
func connect_GetFlowAssociation(cfg aws.Config, client *connect.Client) {
	input := &connect.GetFlowAssociationInput{
		// InstanceId: *string, // Required
		// ResourceId: *string, // Required
		// ResourceType: types.FlowAssociationResourceType, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceId) > 0 {
		input.ResourceId = aws.String(_connectResourceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFlowAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets historical metric data from the specified Amazon Connect instance.
// For a description of each historical metric, see [Metrics definitions] in the Amazon Connect
// Administrator Guide.
//
// We recommend using the [GetMetricDataV2] API. It provides more flexibility, features, and the
// ability to query longer time ranges than GetMetricData . Use it to retrieve
// historical agent and contact metrics for the last 3 months, at varying
// intervals. You can also use it to build custom dashboards to measure historical
// queue and agent performance. For example, you can track the number of incoming
// contacts for the last 7 days, with data split by day, to see how contact volume
// changed per day of the week.
//
// [GetMetricDataV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_GetMetricDataV2.html
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
func connect_GetMetricData(cfg aws.Config, client *connect.Client) {
	input := &connect.GetMetricDataInput{
		// EndTime: *time.Time, // Required
		// Filters: *types.Filters, // Required
		// HistoricalMetrics: []types.HistoricalMetric, // Required
		// InstanceId: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_connectEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _connectEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_connectFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectHistoricalMetrics) > 0 {
		if err := assignInputField(input, "HistoricalMetrics", _connectHistoricalMetrics); err != nil {
			log.Errorf("invalid --historical-metrics: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _connectStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_connectGroupings) > 0 {
		if err := assignInputField(input, "Groupings", _connectGroupings[0]); err != nil {
			log.Errorf("invalid --groupings: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetMetricData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.GetMetricDataOutput
	p := connect.NewGetMetricDataPaginator(client, input)
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

// Gets metric data from the specified Amazon Connect instance.
// GetMetricDataV2 offers more features than [GetMetricData], the previous version of this API.
// It has new metrics, offers filtering at a metric level, and offers the ability
// to filter and group data by channels, queues, routing profiles, agents, and
// agent hierarchy levels. It can retrieve historical data for the last 3 months,
// at varying intervals. It does not support agent queues.
//
// For a description of the historical metrics that are supported by
// GetMetricDataV2 and GetMetricData , see [Metrics definitions] in the Amazon Connect Administrator
// Guide.
//
// When you make a successful API request, you can expect the following metric
// values in the response:
//
// - Metric value is null: The calculation cannot be performed due to divide by
// zero or insufficient data
//
// - Metric value is a number (including 0) of defined type: The number provided
// is the calculation result
//
// - MetricResult list is empty: The request cannot find any data in the system
//
// The following guidelines can help you work with the API:
//
// - Each dimension in the metric response must contain a value
//
// - Each item in MetricResult must include all requested metrics
//
// - If the response is slow due to large result sets, try these approaches:
//
// - Narrow the time range of your request
//
// - Add filters to reduce the amount of data returned
//
// [GetMetricData]: https://docs.aws.amazon.com/connect/latest/APIReference/API_GetMetricData.html
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
func connect_GetMetricDataV2(cfg aws.Config, client *connect.Client) {
	input := &connect.GetMetricDataV2Input{
		// EndTime: *time.Time, // Required
		// Filters: []types.FilterV2, // Required
		// Metrics: []types.MetricV2, // Required
		// ResourceArn: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_connectEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _connectEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_connectFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectMetrics) > 0 {
		if err := assignInputField(input, "Metrics", _connectMetrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_connectResourceArn) > 0 {
		input.ResourceArn = aws.String(_connectResourceArn)
	}
	if len(_connectStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _connectStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_connectGroupings) > 0 {
		input.Groupings = append([]string(nil), _connectGroupings...)
	}
	if len(_connectInterval) > 0 {
		if err := assignInputField(input, "Interval", _connectInterval); err != nil {
			log.Errorf("invalid --interval: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetMetricDataV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.GetMetricDataV2Output
	p := connect.NewGetMetricDataV2Paginator(client, input)
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

// Gets the prompt file.
func connect_GetPromptFile(cfg aws.Config, client *connect.Client) {
	input := &connect.GetPromptFileInput{
		// InstanceId: *string, // Required
		// PromptId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPromptId) > 0 {
		input.PromptId = aws.String(_connectPromptId)
	}

	if resp, err := client.GetPromptFile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a specific task template in the specified Amazon Connect
// instance.
func connect_GetTaskTemplate(cfg aws.Config, client *connect.Client) {
	input := &connect.GetTaskTemplateInput{
		// InstanceId: *string, // Required
		// TaskTemplateId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTaskTemplateId) > 0 {
		input.TaskTemplateId = aws.String(_connectTaskTemplateId)
	}
	if len(_connectSnapshotVersion) > 0 {
		input.SnapshotVersion = aws.String(_connectSnapshotVersion)
	}

	if resp, err := client.GetTaskTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an overview of a test execution that includes the status of the
// execution, start and end time, and observation summary.
func connect_GetTestCaseExecutionSummary(cfg aws.Config, client *connect.Client) {
	input := &connect.GetTestCaseExecutionSummaryInput{
		// InstanceId: *string, // Required
		// TestCaseExecutionId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseExecutionId) > 0 {
		input.TestCaseExecutionId = aws.String(_connectTestCaseExecutionId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}

	if resp, err := client.GetTestCaseExecutionSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current traffic distribution for a given traffic distribution
// group.
func connect_GetTrafficDistribution(cfg aws.Config, client *connect.Client) {
	input := &connect.GetTrafficDistributionInput{
		// Id: *string, // Required
	}

	if len(_connectId) > 0 {
		input.Id = aws.String(_connectId)
	}

	if resp, err := client.GetTrafficDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a claimed phone number from an external service, such as Amazon Web
// Services End User Messaging, into an Amazon Connect instance. You can call this
// API only in the same Amazon Web Services Region where the Amazon Connect
// instance was created.
//
// Call the [DescribePhoneNumber] API to verify the status of a previous ImportPhoneNumber operation.
//
// If you plan to claim or import numbers and then release numbers frequently,
// contact us for a service quota exception. Otherwise, it is possible you will be
// blocked from claiming and releasing any more numbers until up to 180 days past
// the oldest number released has expired.
//
// By default you can claim or import and then release up to 200% of your maximum
// number of active phone numbers. If you claim or import and then release phone
// numbers using the UI or API during a rolling 180 day cycle that exceeds 200% of
// your phone number service level quota, you will be blocked from claiming or
// importing any more numbers until 180 days past the oldest number released has
// expired.
//
// For example, if you already have 99 claimed or imported numbers and a service
// level quota of 99 phone numbers, and in any 180 day period you release 99, claim
// 99, and then release 99, you will have exceeded the 200% limit. At that point
// you are blocked from claiming any more numbers until you open an Amazon Web
// Services Support ticket.
//
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
func connect_ImportPhoneNumber(cfg aws.Config, client *connect.Client) {
	input := &connect.ImportPhoneNumberInput{
		// InstanceId: *string, // Required
		// SourcePhoneNumberArn: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSourcePhoneNumberArn) > 0 {
		input.SourcePhoneNumberArn = aws.String(_connectSourcePhoneNumberArn)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectPhoneNumberDescription) > 0 {
		input.PhoneNumberDescription = aws.String(_connectPhoneNumberDescription)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a media asset (such as a logo) for use in a workspace.
func connect_ImportWorkspaceMedia(cfg aws.Config, client *connect.Client) {
	input := &connect.ImportWorkspaceMediaInput{
		// InstanceId: *string, // Required
		// MediaSource: *string, // Required
		// MediaType: types.MediaType, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMediaSource) > 0 {
		input.MediaSource = aws.String(_connectMediaSource)
	}
	if len(_connectMediaType) > 0 {
		if err := assignInputField(input, "MediaType", _connectMediaType); err != nil {
			log.Errorf("invalid --media-type: %s", err.Error())
			return
		}
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.ImportWorkspaceMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists agent statuses.
func connect_ListAgentStatuses(cfg aws.Config, client *connect.Client) {
	input := &connect.ListAgentStatusesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAgentStatusTypes) > 0 {
		if err := assignInputField(input, "AgentStatusTypes", _connectAgentStatusTypes); err != nil {
			log.Errorf("invalid --agent-status-types: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentStatuses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListAgentStatusesOutput
	p := connect.NewListAgentStatusesPaginator(client, input)
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

// Lists the association status of requested dataset ID for a given Amazon Connect
// instance.
func connect_ListAnalyticsDataAssociations(cfg aws.Config, client *connect.Client) {
	input := &connect.ListAnalyticsDataAssociationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDataSetId) > 0 {
		input.DataSetId = aws.String(_connectDataSetId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if resp, err := client.ListAnalyticsDataAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the data lake datasets available to associate with for a given Amazon
// Connect instance.
func connect_ListAnalyticsDataLakeDataSets(cfg aws.Config, client *connect.Client) {
	input := &connect.ListAnalyticsDataLakeDataSetsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if resp, err := client.ListAnalyticsDataLakeDataSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all approved origins associated with the instance.
func connect_ListApprovedOrigins(cfg aws.Config, client *connect.Client) {
	input := &connect.ListApprovedOriginsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApprovedOrigins(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListApprovedOriginsOutput
	p := connect.NewListApprovedOriginsPaginator(client, input)
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

// Provides information about contact tree, a list of associated contacts with a
// unique identifier.
func connect_ListAssociatedContacts(cfg aws.Config, client *connect.Client) {
	input := &connect.ListAssociatedContactsInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if resp, err := client.ListAssociatedContacts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change. To
// request access to this API, contact Amazon Web Services Support.
//
// Provides summary information about the authentication profiles in a specified
// Amazon Connect instance.
func connect_ListAuthenticationProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.ListAuthenticationProfilesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAuthenticationProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListAuthenticationProfilesOutput
	p := connect.NewListAuthenticationProfilesPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// For the specified version of Amazon Lex, returns a paginated list of all the
// Amazon Lex bots currently associated with the instance. Use this API to return
// both Amazon Lex V1 and V2 bots.
func connect_ListBots(cfg aws.Config, client *connect.Client) {
	input := &connect.ListBotsInput{
		// InstanceId: *string, // Required
		// LexVersion: types.LexVersion, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLexVersion) > 0 {
		if err := assignInputField(input, "LexVersion", _connectLexVersion); err != nil {
			log.Errorf("invalid --lex-version: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListBotsOutput
	p := connect.NewListBotsPaginator(client, input)
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

// Provides information about the child hours of operations for the specified
// parent hours of operation.
//
// For more information about child hours of operations, see [Link overrides from different hours of operation] in the Administrator
// Guide.
//
// [Link overrides from different hours of operation]: https://docs.aws.amazon.com/connect/latest/adminguide/
func connect_ListChildHoursOfOperations(cfg aws.Config, client *connect.Client) {
	input := &connect.ListChildHoursOfOperationsInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChildHoursOfOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListChildHoursOfOperationsOutput
	p := connect.NewListChildHoursOfOperationsPaginator(client, input)
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

// Lists contact evaluations in the specified Amazon Connect instance.
func connect_ListContactEvaluations(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactEvaluationsInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactEvaluations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactEvaluationsOutput
	p := connect.NewListContactEvaluationsPaginator(client, input)
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

// Lists all aliases associated with a contact flow module, showing their current
// version mappings and metadata.
func connect_ListContactFlowModuleAliases(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactFlowModuleAliasesInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactFlowModuleAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactFlowModuleAliasesOutput
	p := connect.NewListContactFlowModuleAliasesPaginator(client, input)
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

// Retrieves a paginated list of all versions for a specific contact flow module.
func connect_ListContactFlowModuleVersions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactFlowModuleVersionsInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactFlowModuleVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactFlowModuleVersionsOutput
	p := connect.NewListContactFlowModuleVersionsPaginator(client, input)
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

// Provides information about the flow modules for the specified Amazon Connect
// instance.
func connect_ListContactFlowModules(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactFlowModulesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowModuleState) > 0 {
		if err := assignInputField(input, "ContactFlowModuleState", _connectContactFlowModuleState); err != nil {
			log.Errorf("invalid --contact-flow-module-state: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactFlowModules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactFlowModulesOutput
	p := connect.NewListContactFlowModulesPaginator(client, input)
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

// Returns all the available versions for the specified Amazon Connect instance
// and flow identifier.
func connect_ListContactFlowVersions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactFlowVersionsInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactFlowVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactFlowVersionsOutput
	p := connect.NewListContactFlowVersionsPaginator(client, input)
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

// Provides information about the flows for the specified Amazon Connect instance.
// You can also create and update flows using the [Amazon Connect Flow language].
//
// For more information about flows, see [Flows] in the Amazon Connect Administrator
// Guide.
//
// [Flows]: https://docs.aws.amazon.com/connect/latest/adminguide/concepts-contact-flows.html
// [Amazon Connect Flow language]: https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html
func connect_ListContactFlows(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactFlowsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowTypes) > 0 {
		if err := assignInputField(input, "ContactFlowTypes", _connectContactFlowTypes); err != nil {
			log.Errorf("invalid --contact-flow-types: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactFlowsOutput
	p := connect.NewListContactFlowsPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// For the specified referenceTypes , returns a list of references associated with
// the contact. References are links to documents that are related to a contact,
// such as emails, attachments, or URLs.
func connect_ListContactReferences(cfg aws.Config, client *connect.Client) {
	input := &connect.ListContactReferencesInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// ReferenceTypes: []types.ReferenceType, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectReferenceTypes) > 0 {
		if err := assignInputField(input, "ReferenceTypes", _connectReferenceTypes); err != nil {
			log.Errorf("invalid --reference-types: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactReferences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListContactReferencesOutput
	p := connect.NewListContactReferencesPaginator(client, input)
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

// Returns all attributes for a specified data table. A maximum of 100 attributes
// per data table is allowed. Customers can request an increase by using Amazon Web
// Services Service Quotas. The response can be filtered by specific attribute IDs
// for CloudFormation integration.
func connect_ListDataTableAttributes(cfg aws.Config, client *connect.Client) {
	input := &connect.ListDataTableAttributesInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAttributeIds) > 0 {
		input.AttributeIds = append([]string(nil), _connectAttributeIds...)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataTableAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListDataTableAttributesOutput
	p := connect.NewListDataTableAttributesPaginator(client, input)
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

// Lists all primary value combinations for a given data table. Returns the unique
// combinations of primary attribute values that identify records in the table. Up
// to 100 records are returned per request.
func connect_ListDataTablePrimaryValues(cfg aws.Config, client *connect.Client) {
	input := &connect.ListDataTablePrimaryValuesInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectPrimaryAttributeValues) > 0 {
		if err := assignInputField(input, "PrimaryAttributeValues", _connectPrimaryAttributeValues); err != nil {
			log.Errorf("invalid --primary-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_connectRecordIds) > 0 {
		input.RecordIds = append([]string(nil), _connectRecordIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListDataTablePrimaryValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListDataTablePrimaryValuesOutput
	p := connect.NewListDataTablePrimaryValuesPaginator(client, input)
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

// Lists values stored in a data table with optional filtering by record IDs or
// primary attribute values. Returns the raw stored values along with metadata such
// as lock versions and modification timestamps.
func connect_ListDataTableValues(cfg aws.Config, client *connect.Client) {
	input := &connect.ListDataTableValuesInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectPrimaryAttributeValues) > 0 {
		if err := assignInputField(input, "PrimaryAttributeValues", _connectPrimaryAttributeValues); err != nil {
			log.Errorf("invalid --primary-attribute-values: %s", err.Error())
			return
		}
	}
	if len(_connectRecordIds) > 0 {
		input.RecordIds = append([]string(nil), _connectRecordIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListDataTableValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListDataTableValuesOutput
	p := connect.NewListDataTableValuesPaginator(client, input)
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

// Lists all data tables for the specified Amazon Connect instance. Returns
// summary information for each table including basic metadata and modification
// details.
func connect_ListDataTables(cfg aws.Config, client *connect.Client) {
	input := &connect.ListDataTablesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListDataTablesOutput
	p := connect.NewListDataTablesPaginator(client, input)
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

// Lists the default vocabularies for the specified Amazon Connect instance.
func connect_ListDefaultVocabularies(cfg aws.Config, client *connect.Client) {
	input := &connect.ListDefaultVocabulariesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _connectLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDefaultVocabularies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListDefaultVocabulariesOutput
	p := connect.NewListDefaultVocabulariesPaginator(client, input)
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

// Lists all security profiles attached to a Q in Connect AIAgent Entity in an
// Amazon Connect instance.
func connect_ListEntitySecurityProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.ListEntitySecurityProfilesInput{
		// EntityArn: *string, // Required
		// EntityType: types.EntityType, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEntityArn) > 0 {
		input.EntityArn = aws.String(_connectEntityArn)
	}
	if len(_connectEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _connectEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntitySecurityProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListEntitySecurityProfilesOutput
	p := connect.NewListEntitySecurityProfilesPaginator(client, input)
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

// Lists versions of an evaluation form in the specified Amazon Connect instance.
func connect_ListEvaluationFormVersions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListEvaluationFormVersionsInput{
		// EvaluationFormId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEvaluationFormVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListEvaluationFormVersionsOutput
	p := connect.NewListEvaluationFormVersionsPaginator(client, input)
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

// Lists evaluation forms in the specified Amazon Connect instance.
func connect_ListEvaluationForms(cfg aws.Config, client *connect.Client) {
	input := &connect.ListEvaluationFormsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEvaluationForms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListEvaluationFormsOutput
	p := connect.NewListEvaluationFormsPaginator(client, input)
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

// List the flow association based on the filters.
func connect_ListFlowAssociations(cfg aws.Config, client *connect.Client) {
	input := &connect.ListFlowAssociationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFlowAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListFlowAssociationsOutput
	p := connect.NewListFlowAssociationsPaginator(client, input)
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

// List the hours of operation overrides.
func connect_ListHoursOfOperationOverrides(cfg aws.Config, client *connect.Client) {
	input := &connect.ListHoursOfOperationOverridesInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHoursOfOperationOverrides(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListHoursOfOperationOverridesOutput
	p := connect.NewListHoursOfOperationOverridesPaginator(client, input)
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

// Provides information about the hours of operation for the specified Amazon
// Connect instance.
//
// For more information about hours of operation, see [Set the Hours of Operation for a Queue] in the Amazon Connect
// Administrator Guide.
//
// [Set the Hours of Operation for a Queue]: https://docs.aws.amazon.com/connect/latest/adminguide/set-hours-operation.html
func connect_ListHoursOfOperations(cfg aws.Config, client *connect.Client) {
	input := &connect.ListHoursOfOperationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHoursOfOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListHoursOfOperationsOutput
	p := connect.NewListHoursOfOperationsPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all attribute types for the given instance.
func connect_ListInstanceAttributes(cfg aws.Config, client *connect.Client) {
	input := &connect.ListInstanceAttributesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListInstanceAttributesOutput
	p := connect.NewListInstanceAttributesPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of storage configs for the identified instance and
// resource type.
func connect_ListInstanceStorageConfigs(cfg aws.Config, client *connect.Client) {
	input := &connect.ListInstanceStorageConfigsInput{
		// InstanceId: *string, // Required
		// ResourceType: types.InstanceStorageResourceType, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceStorageConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListInstanceStorageConfigsOutput
	p := connect.NewListInstanceStorageConfigsPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// Return a list of instances which are in active state, creation-in-progress
// state, and failed state. Instances that aren't successfully created (they are in
// a failed state) are returned only for 24 hours after the CreateInstance API was
// invoked.
func connect_ListInstances(cfg aws.Config, client *connect.Client) {
	input := &connect.ListInstancesInput{}

	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListInstancesOutput
	p := connect.NewListInstancesPaginator(client, input)
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

// Provides summary information about the Amazon Web Services resource
// associations for the specified Amazon Connect instance.
func connect_ListIntegrationAssociations(cfg aws.Config, client *connect.Client) {
	input := &connect.ListIntegrationAssociationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_connectIntegrationArn)
	}
	if len(_connectIntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _connectIntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIntegrationAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListIntegrationAssociationsOutput
	p := connect.NewListIntegrationAssociationsPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all Lambda functions that display in the dropdown
// options in the relevant flow blocks.
func connect_ListLambdaFunctions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListLambdaFunctionsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLambdaFunctions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListLambdaFunctionsOutput
	p := connect.NewListLambdaFunctionsPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all the Amazon Lex V1 bots currently associated
// with the instance. To return both Amazon Lex V1 and V2 bots, use the [ListBots]API.
//
// [ListBots]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListBots.html
func connect_ListLexBots(cfg aws.Config, client *connect.Client) {
	input := &connect.ListLexBotsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLexBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListLexBotsOutput
	p := connect.NewListLexBotsPaginator(client, input)
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

// Retrieves a paginated list of all notifications in the Amazon Connect instance.
func connect_ListNotifications(cfg aws.Config, client *connect.Client) {
	input := &connect.ListNotificationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if resp, err := client.ListNotifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the phone numbers for the specified Amazon Connect
// instance.
//
// For more information about phone numbers, see [Set Up Phone Numbers for Your Contact Center] in the Amazon Connect
// Administrator Guide.
//
// - We recommend using [ListPhoneNumbersV2]to return phone number types. ListPhoneNumbers doesn't
// support number types UIFN , SHARED , THIRD_PARTY_TF , and THIRD_PARTY_DID .
// While it returns numbers of those types, it incorrectly lists them as
// TOLL_FREE or DID .
//
// - The phone number Arn value that is returned from each of the items in the [PhoneNumberSummaryList]
// cannot be used to tag phone number resources. It will fail with a
// ResourceNotFoundException . Instead, use the [ListPhoneNumbersV2]API. It returns the new phone
// number ARN that can be used to tag phone number resources.
//
// [ListPhoneNumbersV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html
// [Set Up Phone Numbers for Your Contact Center]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-center-phone-number.html
// [PhoneNumberSummaryList]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbers.html#connect-ListPhoneNumbers-response-PhoneNumberSummaryList
func connect_ListPhoneNumbers(cfg aws.Config, client *connect.Client) {
	input := &connect.ListPhoneNumbersInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectPhoneNumberCountryCodes) > 0 {
		if err := assignInputField(input, "PhoneNumberCountryCodes", _connectPhoneNumberCountryCodes); err != nil {
			log.Errorf("invalid --phone-number-country-codes: %s", err.Error())
			return
		}
	}
	if len(_connectPhoneNumberTypes) > 0 {
		if err := assignInputField(input, "PhoneNumberTypes", _connectPhoneNumberTypes); err != nil {
			log.Errorf("invalid --phone-number-types: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListPhoneNumbersOutput
	p := connect.NewListPhoneNumbersPaginator(client, input)
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

// Lists phone numbers claimed to your Amazon Connect instance or traffic
// distribution group. If the provided TargetArn is a traffic distribution group,
// you can call this API in both Amazon Web Services Regions associated with
// traffic distribution group.
//
// For more information about phone numbers, see [Set Up Phone Numbers for Your Contact Center] in the Amazon Connect
// Administrator Guide.
//
// - When given an instance ARN, ListPhoneNumbersV2 returns only the phone
// numbers claimed to the instance.
//
// - When given a traffic distribution group ARN ListPhoneNumbersV2 returns only
// the phone numbers claimed to the traffic distribution group.
//
// [Set Up Phone Numbers for Your Contact Center]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-center-phone-number.html
func connect_ListPhoneNumbersV2(cfg aws.Config, client *connect.Client) {
	input := &connect.ListPhoneNumbersV2Input{}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectPhoneNumberCountryCodes) > 0 {
		if err := assignInputField(input, "PhoneNumberCountryCodes", _connectPhoneNumberCountryCodes); err != nil {
			log.Errorf("invalid --phone-number-country-codes: %s", err.Error())
			return
		}
	}
	if len(_connectPhoneNumberPrefix) > 0 {
		input.PhoneNumberPrefix = aws.String(_connectPhoneNumberPrefix)
	}
	if len(_connectPhoneNumberTypes) > 0 {
		if err := assignInputField(input, "PhoneNumberTypes", _connectPhoneNumberTypes); err != nil {
			log.Errorf("invalid --phone-number-types: %s", err.Error())
			return
		}
	}
	if len(_connectTargetArn) > 0 {
		input.TargetArn = aws.String(_connectTargetArn)
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumbersV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListPhoneNumbersV2Output
	p := connect.NewListPhoneNumbersV2Paginator(client, input)
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

// Lists predefined attributes for the specified Amazon Connect instance. A
// predefined attribute is made up of a name and a value. You can use predefined
// attributes for:
//
// - Routing proficiency (for example, agent certification) that has predefined
// values (for example, a list of possible certifications). For more information,
// see [Create predefined attributes for routing contacts to agents].
//
// - Contact information that varies between transfers or conferences, such as
// the name of the business unit handling the contact. For more information, see [Use contact segment attributes]
// .
//
// For the predefined attributes per instance quota, see [Amazon Connect quotas].
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Use contact segment attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/use-contact-segment-attributes.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
func connect_ListPredefinedAttributes(cfg aws.Config, client *connect.Client) {
	input := &connect.ListPredefinedAttributesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPredefinedAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListPredefinedAttributesOutput
	p := connect.NewListPredefinedAttributesPaginator(client, input)
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

// Provides information about the prompts for the specified Amazon Connect
// instance.
func connect_ListPrompts(cfg aws.Config, client *connect.Client) {
	input := &connect.ListPromptsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPrompts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListPromptsOutput
	p := connect.NewListPromptsPaginator(client, input)
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

// Lists the quick connects associated with a queue.
func connect_ListQueueQuickConnects(cfg aws.Config, client *connect.Client) {
	input := &connect.ListQueueQuickConnectsInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQueueQuickConnects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListQueueQuickConnectsOutput
	p := connect.NewListQueueQuickConnectsPaginator(client, input)
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

// Provides information about the queues for the specified Amazon Connect instance.
// If you do not specify a QueueTypes parameter, both standard and agent queues
// are returned. This might cause an unexpected truncation of results if you have
// more than 1000 agents and you limit the number of results of the API call in
// code.
//
// For more information about queues, see [Queues: Standard and Agent] in the Amazon Connect Administrator
// Guide.
//
// [Queues: Standard and Agent]: https://docs.aws.amazon.com/connect/latest/adminguide/concepts-queues-standard-and-agent.html
func connect_ListQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.ListQueuesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectQueueTypes) > 0 {
		if err := assignInputField(input, "QueueTypes", _connectQueueTypes); err != nil {
			log.Errorf("invalid --queue-types: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListQueuesOutput
	p := connect.NewListQueuesPaginator(client, input)
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

// Provides information about the quick connects for the specified Amazon Connect
// instance.
func connect_ListQuickConnects(cfg aws.Config, client *connect.Client) {
	input := &connect.ListQuickConnectsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectQuickConnectTypes) > 0 {
		if err := assignInputField(input, "QuickConnectTypes", _connectQuickConnectTypes); err != nil {
			log.Errorf("invalid --quick-connect-types: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListQuickConnects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListQuickConnectsOutput
	p := connect.NewListQuickConnectsPaginator(client, input)
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

// Provides a list of analysis segments for a real-time chat analysis session.
// This API supports CHAT channels only.
//
// This API does not support VOICE. If you attempt to use it for VOICE, an
// InvalidRequestException occurs.
func connect_ListRealtimeContactAnalysisSegmentsV2(cfg aws.Config, client *connect.Client) {
	input := &connect.ListRealtimeContactAnalysisSegmentsV2Input{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// OutputType: types.RealTimeContactAnalysisOutputType, // Required
		// SegmentTypes: []types.RealTimeContactAnalysisSegmentType, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectOutputType) > 0 {
		if err := assignInputField(input, "OutputType", _connectOutputType); err != nil {
			log.Errorf("invalid --output-type: %s", err.Error())
			return
		}
	}
	if len(_connectSegmentTypes) > 0 {
		if err := assignInputField(input, "SegmentTypes", _connectSegmentTypes); err != nil {
			log.Errorf("invalid --segment-types: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRealtimeContactAnalysisSegmentsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListRealtimeContactAnalysisSegmentsV2Output
	p := connect.NewListRealtimeContactAnalysisSegmentsV2Paginator(client, input)
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

// Lists the manual assignment queues associated with a routing profile.
// # Use cases
//
// Following are common uses cases for this API:
//
// - This API returns list of queues where contacts can be manually assigned or
// picked by an agent who has access to the Worklist app. The user can additionally
// filter on queues, if they have access to those queues (otherwise a invalid
// request exception will be thrown).
//
// # For information about how manual contact assignment works in the agent
//
// workspace, see the [Access the Worklist app in the Amazon Connect agent workspace]in the Amazon Connect Administrator Guide.
//
// # Important things to know
//
// - This API only returns the manual assignment queues associated with a
// routing profile. Use the ListRoutingProfileQueues API to list the auto
// assignment queues for the routing profile.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Access the Worklist app in the Amazon Connect agent workspace]: https://docs.aws.amazon.com/connect/latest/adminguide/worklist-app.html
func connect_ListRoutingProfileManualAssignmentQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.ListRoutingProfileManualAssignmentQueuesInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutingProfileManualAssignmentQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListRoutingProfileManualAssignmentQueuesOutput
	p := connect.NewListRoutingProfileManualAssignmentQueuesPaginator(client, input)
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

// Lists the queues associated with a routing profile.
func connect_ListRoutingProfileQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.ListRoutingProfileQueuesInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutingProfileQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListRoutingProfileQueuesOutput
	p := connect.NewListRoutingProfileQueuesPaginator(client, input)
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

// Provides summary information about the routing profiles for the specified
// Amazon Connect instance.
//
// For more information about routing profiles, see [Routing Profiles] and [Create a Routing Profile] in the Amazon Connect
// Administrator Guide.
//
// [Create a Routing Profile]: https://docs.aws.amazon.com/connect/latest/adminguide/routing-profiles.html
// [Routing Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing.html
func connect_ListRoutingProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.ListRoutingProfilesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutingProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListRoutingProfilesOutput
	p := connect.NewListRoutingProfilesPaginator(client, input)
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

// List all rules for the specified Amazon Connect instance.
func connect_ListRules(cfg aws.Config, client *connect.Client) {
	input := &connect.ListRulesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectEventSourceName) > 0 {
		if err := assignInputField(input, "EventSourceName", _connectEventSourceName); err != nil {
			log.Errorf("invalid --event-source-name: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectPublishStatus) > 0 {
		if err := assignInputField(input, "PublishStatus", _connectPublishStatus); err != nil {
			log.Errorf("invalid --publish-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListRulesOutput
	p := connect.NewListRulesPaginator(client, input)
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

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all security keys associated with the instance.
func connect_ListSecurityKeys(cfg aws.Config, client *connect.Client) {
	input := &connect.ListSecurityKeysInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListSecurityKeysOutput
	p := connect.NewListSecurityKeysPaginator(client, input)
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

// Returns a list of third-party applications or MCP Servers in a specific
// security profile.
func connect_ListSecurityProfileApplications(cfg aws.Config, client *connect.Client) {
	input := &connect.ListSecurityProfileApplicationsInput{
		// InstanceId: *string, // Required
		// SecurityProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileId) > 0 {
		input.SecurityProfileId = aws.String(_connectSecurityProfileId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityProfileApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListSecurityProfileApplicationsOutput
	p := connect.NewListSecurityProfileApplicationsPaginator(client, input)
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

// A list of Flow Modules an AI Agent can invoke as a tool
func connect_ListSecurityProfileFlowModules(cfg aws.Config, client *connect.Client) {
	input := &connect.ListSecurityProfileFlowModulesInput{
		// InstanceId: *string, // Required
		// SecurityProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileId) > 0 {
		input.SecurityProfileId = aws.String(_connectSecurityProfileId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityProfileFlowModules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListSecurityProfileFlowModulesOutput
	p := connect.NewListSecurityProfileFlowModulesPaginator(client, input)
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

// Lists the permissions granted to a security profile.
// For information about security profiles, see [Security Profiles] in the Amazon Connect
// Administrator Guide. For a mapping of the API name and user interface name of
// the security profile permissions, see [List of security profile permissions].
//
// [Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html
// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
func connect_ListSecurityProfilePermissions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListSecurityProfilePermissionsInput{
		// InstanceId: *string, // Required
		// SecurityProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileId) > 0 {
		input.SecurityProfileId = aws.String(_connectSecurityProfileId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityProfilePermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListSecurityProfilePermissionsOutput
	p := connect.NewListSecurityProfilePermissionsPaginator(client, input)
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

// Provides summary information about the security profiles for the specified
// Amazon Connect instance.
//
// For more information about security profiles, see [Security Profiles] in the Amazon Connect
// Administrator Guide. For a mapping of the API name and user interface name of
// the security profile permissions, see [List of security profile permissions].
//
// [Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html
// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
func connect_ListSecurityProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.ListSecurityProfilesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListSecurityProfilesOutput
	p := connect.NewListSecurityProfilesPaginator(client, input)
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

// Lists the tags for the specified resource.
// For sample policies that use tags, see [Amazon Connect Identity-Based Policy Examples] in the Amazon Connect Administrator
// Guide.
//
// [Amazon Connect Identity-Based Policy Examples]: https://docs.aws.amazon.com/connect/latest/adminguide/security_iam_id-based-policy-examples.html
func connect_ListTagsForResource(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_connectResourceArn) > 0 {
		input.ResourceArn = aws.String(_connectResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists task templates for the specified Amazon Connect instance.
func connect_ListTaskTemplates(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTaskTemplatesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTaskTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListTaskTemplatesOutput
	p := connect.NewListTaskTemplatesPaginator(client, input)
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

// Lists detailed steps of test case execution that includes all observations
// along with actions taken and data associated in the specified Amazon Connect
// instance.
func connect_ListTestCaseExecutionRecords(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTestCaseExecutionRecordsInput{
		// InstanceId: *string, // Required
		// TestCaseExecutionId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseExecutionId) > 0 {
		input.TestCaseExecutionId = aws.String(_connectTestCaseExecutionId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTestCaseExecutionRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all test case executions and allows filtering by test case id, test case
// name, start time, end time or status of the execution for the specified Amazon
// Connect instance.
func connect_ListTestCaseExecutions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTestCaseExecutionsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _connectEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _connectStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}
	if len(_connectTestCaseName) > 0 {
		input.TestCaseName = aws.String(_connectTestCaseName)
	}

	if resp, err := client.ListTestCaseExecutions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the test cases present in the specific Amazon Connect instance.
func connect_ListTestCases(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTestCasesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTestCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListTestCasesOutput
	p := connect.NewListTestCasesPaginator(client, input)
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

// Lists traffic distribution group users.
func connect_ListTrafficDistributionGroupUsers(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTrafficDistributionGroupUsersInput{
		// TrafficDistributionGroupId: *string, // Required
	}

	if len(_connectTrafficDistributionGroupId) > 0 {
		input.TrafficDistributionGroupId = aws.String(_connectTrafficDistributionGroupId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrafficDistributionGroupUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListTrafficDistributionGroupUsersOutput
	p := connect.NewListTrafficDistributionGroupUsersPaginator(client, input)
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

// Lists traffic distribution groups.
func connect_ListTrafficDistributionGroups(cfg aws.Config, client *connect.Client) {
	input := &connect.ListTrafficDistributionGroupsInput{}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrafficDistributionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListTrafficDistributionGroupsOutput
	p := connect.NewListTrafficDistributionGroupsPaginator(client, input)
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

// Lists the use cases for the integration association.
func connect_ListUseCases(cfg aws.Config, client *connect.Client) {
	input := &connect.ListUseCasesInput{
		// InstanceId: *string, // Required
		// IntegrationAssociationId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectIntegrationAssociationId) > 0 {
		input.IntegrationAssociationId = aws.String(_connectIntegrationAssociationId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUseCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListUseCasesOutput
	p := connect.NewListUseCasesPaginator(client, input)
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

// Provides summary information about the hierarchy groups for the specified
// Amazon Connect instance.
//
// For more information about agent hierarchies, see [Set Up Agent Hierarchies] in the Amazon Connect
// Administrator Guide.
//
// [Set Up Agent Hierarchies]: https://docs.aws.amazon.com/connect/latest/adminguide/agent-hierarchy.html
func connect_ListUserHierarchyGroups(cfg aws.Config, client *connect.Client) {
	input := &connect.ListUserHierarchyGroupsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserHierarchyGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListUserHierarchyGroupsOutput
	p := connect.NewListUserHierarchyGroupsPaginator(client, input)
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

// Retrieves a paginated list of notifications for a specific user, including the
// notification status for that user.
func connect_ListUserNotifications(cfg aws.Config, client *connect.Client) {
	input := &connect.ListUserNotificationsInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if resp, err := client.ListUserNotifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists proficiencies associated with a user.
func connect_ListUserProficiencies(cfg aws.Config, client *connect.Client) {
	input := &connect.ListUserProficienciesInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserProficiencies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListUserProficienciesOutput
	p := connect.NewListUserProficienciesPaginator(client, input)
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

// Provides summary information about the users for the specified Amazon Connect
// instance.
func connect_ListUsers(cfg aws.Config, client *connect.Client) {
	input := &connect.ListUsersInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
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

	var results []*connect.ListUsersOutput
	p := connect.NewListUsersPaginator(client, input)
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

// Returns all the available versions for the specified Amazon Connect instance
// and view identifier.
//
// Results will be sorted from highest to lowest.
func connect_ListViewVersions(cfg aws.Config, client *connect.Client) {
	input := &connect.ListViewVersionsInput{
		// InstanceId: *string, // Required
		// ViewId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListViewVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListViewVersionsOutput
	p := connect.NewListViewVersionsPaginator(client, input)
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

// Returns views in the given instance.
// Results are sorted primarily by type, and secondarily by name.
func connect_ListViews(cfg aws.Config, client *connect.Client) {
	input := &connect.ListViewsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectType) > 0 {
		if err := assignInputField(input, "Type", _connectType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListViewsOutput
	p := connect.NewListViewsPaginator(client, input)
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

// Lists media assets (such as logos) associated with a workspace.
func connect_ListWorkspaceMedia(cfg aws.Config, client *connect.Client) {
	input := &connect.ListWorkspaceMediaInput{
		// InstanceId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.ListWorkspaceMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the page configurations in a workspace, including the views assigned to
// each page.
func connect_ListWorkspacePages(cfg aws.Config, client *connect.Client) {
	input := &connect.ListWorkspacePagesInput{
		// InstanceId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspacePages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListWorkspacePagesOutput
	p := connect.NewListWorkspacePagesPaginator(client, input)
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

// Lists the workspaces in an Amazon Connect instance.
func connect_ListWorkspaces(cfg aws.Config, client *connect.Client) {
	input := &connect.ListWorkspacesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.ListWorkspacesOutput
	p := connect.NewListWorkspacesPaginator(client, input)
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

// Initiates silent monitoring of a contact. The Contact Control Panel (CCP) of
// the user specified by userId will be set to silent monitoring mode on the
// contact.
func connect_MonitorContact(cfg aws.Config, client *connect.Client) {
	input := &connect.MonitorContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectAllowedMonitorCapabilities) > 0 {
		if err := assignInputField(input, "AllowedMonitorCapabilities", _connectAllowedMonitorCapabilities); err != nil {
			log.Errorf("invalid --allowed-monitor-capabilities: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.MonitorContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows pausing an ongoing task contact.
func connect_PauseContact(cfg aws.Config, client *connect.Client) {
	input := &connect.PauseContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}

	if resp, err := client.PauseContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the current status of a user or agent in Amazon Connect. If the agent
// is currently handling a contact, this sets the agent's next status.
//
// For more information, see [Agent status] and [Set your next status] in the Amazon Connect Administrator Guide.
//
// [Agent status]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-agent-status.html
// [Set your next status]: https://docs.aws.amazon.com/connect/latest/adminguide/set-next-status.html
func connect_PutUserStatus(cfg aws.Config, client *connect.Client) {
	input := &connect.PutUserStatusInput{
		// AgentStatusId: *string, // Required
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectAgentStatusId) > 0 {
		input.AgentStatusId = aws.String(_connectAgentStatusId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.PutUserStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Releases a phone number previously claimed to an Amazon Connect instance or
// traffic distribution group. You can call this API only in the Amazon Web
// Services Region where the number was claimed.
//
// To release phone numbers from a traffic distribution group, use the
// ReleasePhoneNumber API, not the Amazon Connect admin website.
//
// After releasing a phone number, the phone number enters into a cooldown period
// for up to 180 days. It cannot be searched for or claimed again until the period
// has ended. If you accidentally release a phone number, contact Amazon Web
// Services Support.
//
// If you plan to claim and release numbers frequently, contact us for a service
// quota exception. Otherwise, it is possible you will be blocked from claiming and
// releasing any more numbers until up to 180 days past the oldest number released
// has expired.
//
// By default you can claim and release up to 200% of your maximum number of
// active phone numbers. If you claim and release phone numbers using the UI or API
// during a rolling 180 day cycle that exceeds 200% of your phone number service
// level quota, you will be blocked from claiming any more numbers until 180 days
// past the oldest number released has expired.
//
// For example, if you already have 99 claimed numbers and a service level quota
// of 99 phone numbers, and in any 180 day period you release 99, claim 99, and
// then release 99, you will have exceeded the 200% limit. At that point you are
// blocked from claiming any more numbers until you open an Amazon Web Services
// support ticket.
func connect_ReleasePhoneNumber(cfg aws.Config, client *connect.Client) {
	input := &connect.ReleasePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_connectPhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_connectPhoneNumberId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.ReleasePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replicates an Amazon Connect instance in the specified Amazon Web Services
// Region and copies configuration information for Amazon Connect resources across
// Amazon Web Services Regions.
//
// For more information about replicating an Amazon Connect instance, see [Create a replica of your existing Amazon Connect instance] in the
// Amazon Connect Administrator Guide.
//
// [Create a replica of your existing Amazon Connect instance]: https://docs.aws.amazon.com/connect/latest/adminguide/create-replica-connect-instance.html
func connect_ReplicateInstance(cfg aws.Config, client *connect.Client) {
	input := &connect.ReplicateInstanceInput{
		// InstanceId: *string, // Required
		// ReplicaAlias: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectReplicaAlias) > 0 {
		input.ReplicaAlias = aws.String(_connectReplicaAlias)
	}
	if len(_connectReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_connectReplicaRegion)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.ReplicateInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows resuming a task contact in a paused state.
func connect_ResumeContact(cfg aws.Config, client *connect.Client) {
	input := &connect.ResumeContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}

	if resp, err := client.ResumeContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When a contact is being recorded, and the recording has been suspended using
// SuspendContactRecording, this API resumes recording whatever recording is
// selected in the flow configuration: call, screen, or both. If only call
// recording or only screen recording is enabled, then it would resume.
//
// Voice and screen recordings are supported.
func connect_ResumeContactRecording(cfg aws.Config, client *connect.Client) {
	input := &connect.ResumeContactRecordingInput{
		// ContactId: *string, // Required
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactRecordingType) > 0 {
		if err := assignInputField(input, "ContactRecordingType", _connectContactRecordingType); err != nil {
			log.Errorf("invalid --contact-recording-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResumeContactRecording(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches AgentStatuses in an Amazon Connect instance, with optional filtering.
func connect_SearchAgentStatuses(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchAgentStatusesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchAgentStatuses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchAgentStatusesOutput
	p := connect.NewSearchAgentStatusesPaginator(client, input)
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

// Searches for available phone numbers that you can claim to your Amazon Connect
// instance or traffic distribution group. If the provided TargetArn is a traffic
// distribution group, you can call this API in both Amazon Web Services Regions
// associated with the traffic distribution group.
func connect_SearchAvailablePhoneNumbers(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchAvailablePhoneNumbersInput{
		// PhoneNumberCountryCode: types.PhoneNumberCountryCode, // Required
		// PhoneNumberType: types.PhoneNumberType, // Required
	}

	if len(_connectPhoneNumberCountryCode) > 0 {
		if err := assignInputField(input, "PhoneNumberCountryCode", _connectPhoneNumberCountryCode); err != nil {
			log.Errorf("invalid --phone-number-country-code: %s", err.Error())
			return
		}
	}
	if len(_connectPhoneNumberType) > 0 {
		if err := assignInputField(input, "PhoneNumberType", _connectPhoneNumberType); err != nil {
			log.Errorf("invalid --phone-number-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectPhoneNumberPrefix) > 0 {
		input.PhoneNumberPrefix = aws.String(_connectPhoneNumberPrefix)
	}
	if len(_connectTargetArn) > 0 {
		input.TargetArn = aws.String(_connectTargetArn)
	}

	if disablePaginator() {
		if resp, err := client.SearchAvailablePhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchAvailablePhoneNumbersOutput
	p := connect.NewSearchAvailablePhoneNumbersPaginator(client, input)
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

// Searches contact evaluations in an Amazon Connect instance, with optional
// filtering.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Find contact evaluations by using specific search criteria.
//
// - Find contact evaluations that are tagged with a specific set of tags.
//
// # Important things to know
//
// - A Search operation, unlike a List operation, takes time to index changes to
// resource (create, update or delete). If you don't see updated information for
// recently changed contact evaluations, try calling the API again in a few
// seconds.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
func connect_SearchContactEvaluations(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchContactEvaluationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchContactEvaluations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches the flow modules in an Amazon Connect instance, with optional
// filtering.
func connect_SearchContactFlowModules(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchContactFlowModulesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchContactFlowModules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchContactFlowModulesOutput
	p := connect.NewSearchContactFlowModulesPaginator(client, input)
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

// Searches the flows in an Amazon Connect instance, with optional filtering.
func connect_SearchContactFlows(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchContactFlowsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchContactFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchContactFlowsOutput
	p := connect.NewSearchContactFlowsPaginator(client, input)
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

// Searches contacts in an Amazon Connect instance.
func connect_SearchContacts(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchContactsInput{
		// InstanceId: *string, // Required
		// TimeRange: *types.SearchContactsTimeRange, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTimeRange) > 0 {
		if err := assignInputField(input, "TimeRange", _connectTimeRange); err != nil {
			log.Errorf("invalid --time-range: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSort) > 0 {
		if err := assignInputField(input, "Sort", _connectSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchContacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchContactsOutput
	p := connect.NewSearchContactsPaginator(client, input)
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

// Searches for data tables based on the table's ID, name, and description. In the
// future, this operation can support searching on attribute names and possibly
// primary values. Follows other search operations closely and supports both search
// criteria and filters.
func connect_SearchDataTables(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchDataTablesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchDataTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchDataTablesOutput
	p := connect.NewSearchDataTablesPaginator(client, input)
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

// Searches email address in an instance, with optional filtering.
func connect_SearchEmailAddresses(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchEmailAddressesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchEmailAddresses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches evaluation forms in an Amazon Connect instance, with optional
// filtering.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - List all evaluation forms in an instance.
//
// - Find all evaluation forms that meet specific criteria, such as Title,
// Description, Status, and more.
//
// - Find all evaluation forms that are tagged with a specific set of tags.
//
// # Important things to know
//
// - A Search operation, unlike a List operation, takes time to index changes to
// resource (create, update or delete). If you don't see updated information for
// recently changed contact evaluations, try calling the API again in a few
// seconds.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
func connect_SearchEvaluationForms(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchEvaluationFormsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchEvaluationForms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches the hours of operation overrides.
func connect_SearchHoursOfOperationOverrides(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchHoursOfOperationOverridesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchHoursOfOperationOverrides(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchHoursOfOperationOverridesOutput
	p := connect.NewSearchHoursOfOperationOverridesPaginator(client, input)
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

// Searches the hours of operation in an Amazon Connect instance, with optional
// filtering.
func connect_SearchHoursOfOperations(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchHoursOfOperationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchHoursOfOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchHoursOfOperationsOutput
	p := connect.NewSearchHoursOfOperationsPaginator(client, input)
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

// Searches for notifications based on specified criteria and filters. Returns a
// paginated list of notifications matching the search parameters, ordered by
// descending creation time. Supports filtering by content and tags.
func connect_SearchNotifications(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchNotificationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchNotifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches predefined attributes that meet certain criteria. A predefined
// attribute is made up of a name and a value. You can use predefined attributes
// for:
//
// - Routing proficiency (for example, agent certification) that has predefined
// values (for example, a list of possible certifications). For more information,
// see [Create predefined attributes for routing contacts to agents].
//
// - Contact information that varies between transfers or conferences, such as
// the name of the business unit handling the contact. For more information, see [Use contact segment attributes]
// .
//
// For the predefined attributes per instance quota, see [Amazon Connect quotas].
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Use contact segment attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/use-contact-segment-attributes.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
func connect_SearchPredefinedAttributes(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchPredefinedAttributesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchPredefinedAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchPredefinedAttributesOutput
	p := connect.NewSearchPredefinedAttributesPaginator(client, input)
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

// Searches prompts in an Amazon Connect instance, with optional filtering.
func connect_SearchPrompts(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchPromptsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchPrompts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchPromptsOutput
	p := connect.NewSearchPromptsPaginator(client, input)
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

// Searches queues in an Amazon Connect instance, with optional filtering.
func connect_SearchQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchQueuesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchQueuesOutput
	p := connect.NewSearchQueuesPaginator(client, input)
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

// Searches quick connects in an Amazon Connect instance, with optional filtering.
func connect_SearchQuickConnects(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchQuickConnectsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchQuickConnects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchQuickConnectsOutput
	p := connect.NewSearchQuickConnectsPaginator(client, input)
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

// Searches tags used in an Amazon Connect instance using optional search criteria.
func connect_SearchResourceTags(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchResourceTagsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _connectResourceTypes...)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchResourceTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchResourceTagsOutput
	p := connect.NewSearchResourceTagsPaginator(client, input)
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

// Searches routing profiles in an Amazon Connect instance, with optional
// filtering.
//
// SearchRoutingProfiles does not populate LastModifiedRegion, LastModifiedTime,
// MediaConcurrencies.CrossChannelBehavior, and AgentAvailabilityTimer in its
// response, but [DescribeRoutingProfile]does.
//
// [DescribeRoutingProfile]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeRoutingProfile.html
func connect_SearchRoutingProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchRoutingProfilesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchRoutingProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchRoutingProfilesOutput
	p := connect.NewSearchRoutingProfilesPaginator(client, input)
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

// Searches security profiles in an Amazon Connect instance, with optional
// filtering.
//
// For information about security profiles, see [Security Profiles] in the Amazon Connect
// Administrator Guide. For a mapping of the API name and user interface name of
// the security profile permissions, see [List of security profile permissions].
//
// [Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html
// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
func connect_SearchSecurityProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchSecurityProfilesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchSecurityProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchSecurityProfilesOutput
	p := connect.NewSearchSecurityProfilesPaginator(client, input)
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

// Searches for test cases in the specified Amazon Connect instance, with optional
// filtering.
func connect_SearchTestCases(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchTestCasesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchTestCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchTestCasesOutput
	p := connect.NewSearchTestCasesPaginator(client, input)
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

// Searches UserHierarchyGroups in an Amazon Connect instance, with optional
// filtering.
//
// The UserHierarchyGroup with "LevelId": "0" is the foundation for building
// levels on top of an instance. It is not user-definable, nor is it visible in the
// UI.
func connect_SearchUserHierarchyGroups(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchUserHierarchyGroupsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchUserHierarchyGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchUserHierarchyGroupsOutput
	p := connect.NewSearchUserHierarchyGroupsPaginator(client, input)
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

// Searches users in an Amazon Connect instance, with optional filtering.
// AfterContactWorkTimeLimit is returned in milliseconds.
func connect_SearchUsers(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchUsersInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchUsersOutput
	p := connect.NewSearchUsersPaginator(client, input)
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

// Searches views based on name, description, or tags.
func connect_SearchViews(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchViewsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchViewsOutput
	p := connect.NewSearchViewsPaginator(client, input)
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

// Searches for vocabularies within a specific Amazon Connect instance using State
// , NameStartsWith , and LanguageCode .
func connect_SearchVocabularies(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchVocabulariesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _connectLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNameStartsWith) > 0 {
		input.NameStartsWith = aws.String(_connectNameStartsWith)
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectState) > 0 {
		if err := assignInputField(input, "State", _connectState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchVocabularies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchVocabulariesOutput
	p := connect.NewSearchVocabulariesPaginator(client, input)
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

// Searches for workspace associations with users or routing profiles based on
// various criteria.
func connect_SearchWorkspaceAssociations(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchWorkspaceAssociationsInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchWorkspaceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchWorkspaceAssociationsOutput
	p := connect.NewSearchWorkspaceAssociationsPaginator(client, input)
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

// Searches workspaces based on name, description, visibility, or tags.
func connect_SearchWorkspaces(cfg aws.Config, client *connect.Client) {
	input := &connect.SearchWorkspacesInput{
		// InstanceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectNextToken) > 0 {
		input.NextToken = aws.String(_connectNextToken)
	}
	if len(_connectSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _connectSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_connectSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _connectSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connect.SearchWorkspacesOutput
	p := connect.NewSearchWorkspacesPaginator(client, input)
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

// Processes chat integration events from Amazon Web Services or external
// integrations to Amazon Connect. A chat integration event includes:
//
// - SourceId, DestinationId, and Subtype: a set of identifiers, uniquely
// representing a chat
//
// - ChatEvent: details of the chat action to perform such as sending a message,
// event, or disconnecting from a chat
//
// When a chat integration event is sent with chat identifiers that do not map to
// an active chat contact, a new chat contact is also created before handling chat
// action.
//
// Access to this API is currently restricted to Amazon Web Services End User
// Messaging for supporting SMS integration.
func connect_SendChatIntegrationEvent(cfg aws.Config, client *connect.Client) {
	input := &connect.SendChatIntegrationEventInput{
		// DestinationId: *string, // Required
		// Event: *types.ChatEvent, // Required
		// SourceId: *string, // Required
	}

	if len(_connectDestinationId) > 0 {
		input.DestinationId = aws.String(_connectDestinationId)
	}
	if len(_connectEvent) > 0 {
		if err := assignInputField(input, "Event", _connectEvent); err != nil {
			log.Errorf("invalid --event: %s", err.Error())
			return
		}
	}
	if len(_connectSourceId) > 0 {
		input.SourceId = aws.String(_connectSourceId)
	}
	if len(_connectNewSessionDetails) > 0 {
		if err := assignInputField(input, "NewSessionDetails", _connectNewSessionDetails); err != nil {
			log.Errorf("invalid --new-session-details: %s", err.Error())
			return
		}
	}
	if len(_connectSubtype) > 0 {
		input.Subtype = aws.String(_connectSubtype)
	}

	if resp, err := client.SendChatIntegrationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send outbound email for outbound campaigns. For more information about outbound
// campaigns, see [Set up Amazon Connect outbound campaigns].
//
// Only the Amazon Connect outbound campaigns service principal is allowed to
// assume a role in your account and call this API.
//
// [Set up Amazon Connect outbound campaigns]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-outbound-campaigns.html
func connect_SendOutboundEmail(cfg aws.Config, client *connect.Client) {
	input := &connect.SendOutboundEmailInput{
		// DestinationEmailAddress: *types.EmailAddressInfo, // Required
		// EmailMessage: *types.OutboundEmailContent, // Required
		// FromEmailAddress: *types.EmailAddressInfo, // Required
		// InstanceId: *string, // Required
		// TrafficType: types.TrafficType, // Required
	}

	if len(_connectDestinationEmailAddress) > 0 {
		if err := assignInputField(input, "DestinationEmailAddress", _connectDestinationEmailAddress); err != nil {
			log.Errorf("invalid --destination-email-address: %s", err.Error())
			return
		}
	}
	if len(_connectEmailMessage) > 0 {
		if err := assignInputField(input, "EmailMessage", _connectEmailMessage); err != nil {
			log.Errorf("invalid --email-message: %s", err.Error())
			return
		}
	}
	if len(_connectFromEmailAddress) > 0 {
		if err := assignInputField(input, "FromEmailAddress", _connectFromEmailAddress); err != nil {
			log.Errorf("invalid --from-email-address: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTrafficType) > 0 {
		if err := assignInputField(input, "TrafficType", _connectTrafficType); err != nil {
			log.Errorf("invalid --traffic-type: %s", err.Error())
			return
		}
	}
	if len(_connectAdditionalRecipients) > 0 {
		if err := assignInputField(input, "AdditionalRecipients", _connectAdditionalRecipients); err != nil {
			log.Errorf("invalid --additional-recipients: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectSourceCampaign) > 0 {
		if err := assignInputField(input, "SourceCampaign", _connectSourceCampaign); err != nil {
			log.Errorf("invalid --source-campaign: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendOutboundEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a pre-signed Amazon S3 URL in response for uploading your content.
// You may only use this API to upload attachments to an [Amazon Connect Case] or [Amazon Connect Email].
//
// [Amazon Connect Case]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_CreateCase.html
// [Amazon Connect Email]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-email-channel.html
func connect_StartAttachedFileUpload(cfg aws.Config, client *connect.Client) {
	input := &connect.StartAttachedFileUploadInput{
		// AssociatedResourceArn: *string, // Required
		// FileName: *string, // Required
		// FileSizeInBytes: *int64, // Required
		// FileUseCaseType: types.FileUseCaseType, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAssociatedResourceArn) > 0 {
		input.AssociatedResourceArn = aws.String(_connectAssociatedResourceArn)
	}
	if len(_connectFileName) > 0 {
		input.FileName = aws.String(_connectFileName)
	}
	if len(_connectFileSizeInBytes) > 0 {
		if err := assignInputField(input, "FileSizeInBytes", _connectFileSizeInBytes); err != nil {
			log.Errorf("invalid --file-size-in-bytes: %s", err.Error())
			return
		}
	}
	if len(_connectFileUseCaseType) > 0 {
		if err := assignInputField(input, "FileUseCaseType", _connectFileUseCaseType); err != nil {
			log.Errorf("invalid --file-use-case-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectCreatedBy) > 0 {
		if err := assignInputField(input, "CreatedBy", _connectCreatedBy); err != nil {
			log.Errorf("invalid --created-by: %s", err.Error())
			return
		}
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_connectUrlExpiryInSeconds) > 0 {
		if err := assignInputField(input, "UrlExpiryInSeconds", _connectUrlExpiryInSeconds); err != nil {
			log.Errorf("invalid --url-expiry-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAttachedFileUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a flow to start a new chat for the customer. Response of this API
// provides a token required to obtain credentials from the [CreateParticipantConnection]API in the Amazon
// Connect Participant Service.
//
// When a new chat contact is successfully created, clients must subscribe to the
// participant’s connection for the created chat within 5 minutes. This is achieved
// by invoking [CreateParticipantConnection]with WEBSOCKET and CONNECTION_CREDENTIALS.
//
// A 429 error occurs in the following situations:
//
// - API rate limit is exceeded. API TPS throttling returns a TooManyRequests
// exception.
//
// - The [quota for concurrent active chats]is exceeded. Active chat throttling returns a LimitExceededException .
//
// If you use the ChatDurationInMinutes parameter and receive a 400 error, your
// account may not support the ability to configure custom chat durations. For more
// information, contact Amazon Web Services Support.
//
// For more information about chat, see the following topics in the Amazon Connect
// Administrator Guide:
//
// [Concepts: Web and mobile messaging capabilities in Amazon Connect]
//
// [Amazon Connect Chat security best practices]
//
// [CreateParticipantConnection]: https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html
// [Concepts: Web and mobile messaging capabilities in Amazon Connect]: https://docs.aws.amazon.com/connect/latest/adminguide/web-and-mobile-chat.html
// [quota for concurrent active chats]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connect_StartChatContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartChatContactInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
		// ParticipantDetails: *types.ParticipantDetails, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectParticipantDetails) > 0 {
		if err := assignInputField(input, "ParticipantDetails", _connectParticipantDetails); err != nil {
			log.Errorf("invalid --participant-details: %s", err.Error())
			return
		}
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectChatDurationInMinutes) > 0 {
		if err := assignInputField(input, "ChatDurationInMinutes", _connectChatDurationInMinutes); err != nil {
			log.Errorf("invalid --chat-duration-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectCustomerId) > 0 {
		input.CustomerId = aws.String(_connectCustomerId)
	}
	if len(_connectDisconnectOnCustomerExit) > 0 {
		if err := assignInputField(input, "DisconnectOnCustomerExit", _connectDisconnectOnCustomerExit); err != nil {
			log.Errorf("invalid --disconnect-on-customer-exit: %s", err.Error())
			return
		}
	}
	if len(_connectInitialMessage) > 0 {
		if err := assignInputField(input, "InitialMessage", _connectInitialMessage); err != nil {
			log.Errorf("invalid --initial-message: %s", err.Error())
			return
		}
	}
	if len(_connectParticipantConfiguration) > 0 {
		if err := assignInputField(input, "ParticipantConfiguration", _connectParticipantConfiguration); err != nil {
			log.Errorf("invalid --participant-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectPersistentChat) > 0 {
		if err := assignInputField(input, "PersistentChat", _connectPersistentChat); err != nil {
			log.Errorf("invalid --persistent-chat: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}
	if len(_connectSegmentAttributes) > 0 {
		if err := assignInputField(input, "SegmentAttributes", _connectSegmentAttributes); err != nil {
			log.Errorf("invalid --segment-attributes: %s", err.Error())
			return
		}
	}
	if len(_connectSupportedMessagingContentTypes) > 0 {
		input.SupportedMessagingContentTypes = append([]string(nil), _connectSupportedMessagingContentTypes...)
	}

	if resp, err := client.StartChatContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an empty evaluation in the specified Amazon Connect instance, using the
// given evaluation form for the particular contact. The evaluation form version
// used for the contact evaluation corresponds to the currently activated version.
// If no version is activated for the evaluation form, the contact evaluation
// cannot be started.
//
// Evaluations created through the public API do not contain answer values
// suggested from automation.
func connect_StartContactEvaluation(cfg aws.Config, client *connect.Client) {
	input := &connect.StartContactEvaluationInput{
		// ContactId: *string, // Required
		// EvaluationFormId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAutoEvaluationConfiguration) > 0 {
		if err := assignInputField(input, "AutoEvaluationConfiguration", _connectAutoEvaluationConfiguration); err != nil {
			log.Errorf("invalid --auto-evaluation-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartContactEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables in-flight message processing for an ongoing chat session. Message
// processing will stay active for the rest of the chat, even if an individual
// contact segment ends.
func connect_StartContactMediaProcessing(cfg aws.Config, client *connect.Client) {
	input := &connect.StartContactMediaProcessingInput{}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectFailureMode) > 0 {
		if err := assignInputField(input, "FailureMode", _connectFailureMode); err != nil {
			log.Errorf("invalid --failure-mode: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectProcessorArn) > 0 {
		input.ProcessorArn = aws.String(_connectProcessorArn)
	}

	if resp, err := client.StartContactMediaProcessing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts recording the contact:
// - If the API is called before the agent joins the call, recording starts when
// the agent joins the call.
//
// - If the API is called after the agent joins the call, recording starts at
// the time of the API call.
//
// StartContactRecording is a one-time action. For example, if you use
// StopContactRecording to stop recording an ongoing call, you can't use
// StartContactRecording to restart it. For scenarios where the recording has
// started and you want to suspend and resume it, such as when collecting sensitive
// information (for example, a credit card number), use SuspendContactRecording and
// ResumeContactRecording.
//
// You can use this API to override the recording behavior configured in the [Set recording behavior]
// block.
//
// Only voice recordings are supported at this time.
//
// [Set recording behavior]: https://docs.aws.amazon.com/connect/latest/adminguide/set-recording-behavior.html
func connect_StartContactRecording(cfg aws.Config, client *connect.Client) {
	input := &connect.StartContactRecordingInput{
		// ContactId: *string, // Required
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
		// VoiceRecordingConfiguration: *types.VoiceRecordingConfiguration, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectVoiceRecordingConfiguration) > 0 {
		if err := assignInputField(input, "VoiceRecordingConfiguration", _connectVoiceRecordingConfiguration); err != nil {
			log.Errorf("invalid --voice-recording-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartContactRecording(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates real-time message streaming for a new chat contact.
// For more information about message streaming, see [Enable real-time chat message streaming] in the Amazon Connect
// Administrator Guide.
//
// For more information about chat, see the following topics in the Amazon Connect
// Administrator Guide:
//
// [Concepts: Web and mobile messaging capabilities in Amazon Connect]
//
// [Amazon Connect Chat security best practices]
//
// [Enable real-time chat message streaming]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-message-streaming.html
// [Concepts: Web and mobile messaging capabilities in Amazon Connect]: https://docs.aws.amazon.com/connect/latest/adminguide/web-and-mobile-chat.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connect_StartContactStreaming(cfg aws.Config, client *connect.Client) {
	input := &connect.StartContactStreamingInput{
		// ChatStreamingConfiguration: *types.ChatStreamingConfiguration, // Required
		// ClientToken: *string, // Required
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectChatStreamingConfiguration) > 0 {
		if err := assignInputField(input, "ChatStreamingConfiguration", _connectChatStreamingConfiguration); err != nil {
			log.Errorf("invalid --chat-streaming-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.StartContactStreaming(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an inbound email contact and initiates a flow to start the email
// contact for the customer. Response of this API provides the ContactId of the
// email contact created.
func connect_StartEmailContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartEmailContactInput{
		// DestinationEmailAddress: *string, // Required
		// EmailMessage: *types.InboundEmailContent, // Required
		// FromEmailAddress: *types.EmailAddressInfo, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectDestinationEmailAddress) > 0 {
		input.DestinationEmailAddress = aws.String(_connectDestinationEmailAddress)
	}
	if len(_connectEmailMessage) > 0 {
		if err := assignInputField(input, "EmailMessage", _connectEmailMessage); err != nil {
			log.Errorf("invalid --email-message: %s", err.Error())
			return
		}
	}
	if len(_connectFromEmailAddress) > 0 {
		if err := assignInputField(input, "FromEmailAddress", _connectFromEmailAddress); err != nil {
			log.Errorf("invalid --from-email-address: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAdditionalRecipients) > 0 {
		if err := assignInputField(input, "AdditionalRecipients", _connectAdditionalRecipients); err != nil {
			log.Errorf("invalid --additional-recipients: %s", err.Error())
			return
		}
	}
	if len(_connectAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _connectAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectReferences) > 0 {
		if err := assignInputField(input, "References", _connectReferences); err != nil {
			log.Errorf("invalid --references: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}
	if len(_connectSegmentAttributes) > 0 {
		if err := assignInputField(input, "SegmentAttributes", _connectSegmentAttributes); err != nil {
			log.Errorf("invalid --segment-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartEmailContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a new outbound SMS or WhatsApp contact to a customer. Response of
// this API provides the ContactId of the outbound SMS or WhatsApp contact created.
//
// SourceEndpoint only supports Endpoints with CONNECT_PHONENUMBER_ARN as Type and
// DestinationEndpoint only supports Endpoints with TELEPHONE_NUMBER as Type.
// ContactFlowId initiates the flow to manage the new contact created.
//
// This API can be used to initiate outbound SMS or WhatsApp contacts for an
// agent, or it can also deflect an ongoing contact to an outbound SMS or WhatsApp
// contact by using the [StartOutboundChatContact]Flow Action.
//
// For more information about using SMS or WhatsApp in Amazon Connect, see the
// following topics in the Amazon Connect Administrator Guide:
//
// [Set up SMS messaging]
//
// [Request an SMS-enabled phone number through Amazon Web Services End User Messaging SMS]
//
// [Set up WhatsApp Business messaging]
//
// [Set up SMS messaging]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-sms-messaging.html
// [Request an SMS-enabled phone number through Amazon Web Services End User Messaging SMS]: https://docs.aws.amazon.com/connect/latest/adminguide/sms-number.html
// [Set up WhatsApp Business messaging]: https://docs.aws.amazon.com/connect/latest/adminguide/whatsapp-integration.html
// [StartOutboundChatContact]: https://docs.aws.amazon.com/connect/latest/APIReference/API_StartOutboundChatContact.html
func connect_StartOutboundChatContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartOutboundChatContactInput{
		// ContactFlowId: *string, // Required
		// DestinationEndpoint: *types.Endpoint, // Required
		// InstanceId: *string, // Required
		// SegmentAttributes: map[string]types.SegmentAttributeValue, // Required
		// SourceEndpoint: *types.Endpoint, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectDestinationEndpoint) > 0 {
		if err := assignInputField(input, "DestinationEndpoint", _connectDestinationEndpoint); err != nil {
			log.Errorf("invalid --destination-endpoint: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSegmentAttributes) > 0 {
		if err := assignInputField(input, "SegmentAttributes", _connectSegmentAttributes); err != nil {
			log.Errorf("invalid --segment-attributes: %s", err.Error())
			return
		}
	}
	if len(_connectSourceEndpoint) > 0 {
		if err := assignInputField(input, "SourceEndpoint", _connectSourceEndpoint); err != nil {
			log.Errorf("invalid --source-endpoint: %s", err.Error())
			return
		}
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectChatDurationInMinutes) > 0 {
		if err := assignInputField(input, "ChatDurationInMinutes", _connectChatDurationInMinutes); err != nil {
			log.Errorf("invalid --chat-duration-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectInitialSystemMessage) > 0 {
		if err := assignInputField(input, "InitialSystemMessage", _connectInitialSystemMessage); err != nil {
			log.Errorf("invalid --initial-system-message: %s", err.Error())
			return
		}
	}
	if len(_connectInitialTemplatedSystemMessage) > 0 {
		if err := assignInputField(input, "InitialTemplatedSystemMessage", _connectInitialTemplatedSystemMessage); err != nil {
			log.Errorf("invalid --initial-templated-system-message: %s", err.Error())
			return
		}
	}
	if len(_connectParticipantDetails) > 0 {
		if err := assignInputField(input, "ParticipantDetails", _connectParticipantDetails); err != nil {
			log.Errorf("invalid --participant-details: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}
	if len(_connectSupportedMessagingContentTypes) > 0 {
		input.SupportedMessagingContentTypes = append([]string(nil), _connectSupportedMessagingContentTypes...)
	}

	if resp, err := client.StartOutboundChatContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a flow to send an agent reply or outbound email contact (created from
// the CreateContact API) to a customer.
func connect_StartOutboundEmailContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartOutboundEmailContactInput{
		// ContactId: *string, // Required
		// DestinationEmailAddress: *types.EmailAddressInfo, // Required
		// EmailMessage: *types.OutboundEmailContent, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectDestinationEmailAddress) > 0 {
		if err := assignInputField(input, "DestinationEmailAddress", _connectDestinationEmailAddress); err != nil {
			log.Errorf("invalid --destination-email-address: %s", err.Error())
			return
		}
	}
	if len(_connectEmailMessage) > 0 {
		if err := assignInputField(input, "EmailMessage", _connectEmailMessage); err != nil {
			log.Errorf("invalid --email-message: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAdditionalRecipients) > 0 {
		if err := assignInputField(input, "AdditionalRecipients", _connectAdditionalRecipients); err != nil {
			log.Errorf("invalid --additional-recipients: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectFromEmailAddress) > 0 {
		if err := assignInputField(input, "FromEmailAddress", _connectFromEmailAddress); err != nil {
			log.Errorf("invalid --from-email-address: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartOutboundEmailContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Places an outbound call to a contact, and then initiates the flow. It performs
// the actions in the flow that's specified (in ContactFlowId ).
//
// Agents do not initiate the outbound API, which means that they do not dial the
// contact. If the flow places an outbound call to a contact, and then puts the
// contact in queue, the call is then routed to the agent, like any other inbound
// case.
//
// Dialing timeout for this operation can be configured with the
// “RingTimeoutInSeconds” parameter. If not specified, the default dialing timeout
// will be 60 seconds which means if the call is not connected within 60 seconds,
// it fails.
//
// UK numbers with a 447 prefix are not allowed by default. Before you can dial
// these UK mobile numbers, you must submit a service quota increase request. For
// more information, see [Amazon Connect Service Quotas]in the Amazon Connect Administrator Guide.
//
// Campaign calls are not allowed by default. Before you can make a call with
// TrafficType = CAMPAIGN , you must submit a service quota increase request to the
// quota [Amazon Connect campaigns].
//
// For Preview dialing mode, only the Amazon Connect outbound campaigns service
// principal is allowed to assume a role in your account and call this API with
// OutboundStrategy.
//
// [Amazon Connect Service Quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Amazon Connect campaigns]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#outbound-communications-quotas
func connect_StartOutboundVoiceContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartOutboundVoiceContactInput{
		// ContactFlowId: *string, // Required
		// DestinationPhoneNumber: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectDestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_connectDestinationPhoneNumber)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAnswerMachineDetectionConfig) > 0 {
		if err := assignInputField(input, "AnswerMachineDetectionConfig", _connectAnswerMachineDetectionConfig); err != nil {
			log.Errorf("invalid --answer-machine-detection-config: %s", err.Error())
			return
		}
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectCampaignId) > 0 {
		input.CampaignId = aws.String(_connectCampaignId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectOutboundStrategy) > 0 {
		if err := assignInputField(input, "OutboundStrategy", _connectOutboundStrategy); err != nil {
			log.Errorf("invalid --outbound-strategy: %s", err.Error())
			return
		}
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectReferences) > 0 {
		if err := assignInputField(input, "References", _connectReferences); err != nil {
			log.Errorf("invalid --references: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}
	if len(_connectRingTimeoutInSeconds) > 0 {
		if err := assignInputField(input, "RingTimeoutInSeconds", _connectRingTimeoutInSeconds); err != nil {
			log.Errorf("invalid --ring-timeout-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_connectSourcePhoneNumber) > 0 {
		input.SourcePhoneNumber = aws.String(_connectSourcePhoneNumber)
	}
	if len(_connectTrafficType) > 0 {
		if err := assignInputField(input, "TrafficType", _connectTrafficType); err != nil {
			log.Errorf("invalid --traffic-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartOutboundVoiceContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts screen sharing for a contact. For more information about screen sharing,
// see [Set up in-app, web, video calling, and screen sharing capabilities]in the Amazon Connect Administrator Guide.
//
// [Set up in-app, web, video calling, and screen sharing capabilities]: https://docs.aws.amazon.com/connect/latest/adminguide/inapp-calling.html
func connect_StartScreenSharing(cfg aws.Config, client *connect.Client) {
	input := &connect.StartScreenSharingInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.StartScreenSharing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a flow to start a new task contact. For more information about task
// contacts, see [Concepts: Tasks in Amazon Connect]in the Amazon Connect Administrator Guide.
//
// When using PreviousContactId and RelatedContactId input parameters, note the
// following:
//
// - PreviousContactId
//
// - Any updates to user-defined task contact attributes on any contact linked
// through the same PreviousContactId will affect every contact in the chain.
//
// - There can be a maximum of 12 linked task contacts in a chain. That is, 12
// task contacts can be created that share the same PreviousContactId .
//
// - RelatedContactId
//
// - Copies contact attributes from the related task contact to the new contact.
//
// - Any update on attributes in a new task contact does not update attributes
// on previous contact.
//
// - There’s no limit on the number of task contacts that can be created that
// use the same RelatedContactId .
//
// In addition, when calling StartTaskContact include only one of these
// parameters: ContactFlowID , QuickConnectID , or TaskTemplateID . Only one
// parameter is required as long as the task template has a flow configured to run
// it. If more than one parameter is specified, or only the TaskTemplateID is
// specified but it does not have a flow configured, the request returns an error
// because Amazon Connect cannot identify the unique flow to run when the task is
// created.
//
// A ServiceQuotaExceededException occurs when the number of open tasks exceeds
// the active tasks quota or there are already 12 tasks referencing the same
// PreviousContactId . For more information about service quotas for task contacts,
// see [Amazon Connect service quotas]in the Amazon Connect Administrator Guide.
//
// [Amazon Connect service quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Concepts: Tasks in Amazon Connect]: https://docs.aws.amazon.com/connect/latest/adminguide/tasks.html
func connect_StartTaskContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartTaskContactInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _connectAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectPreviousContactId) > 0 {
		input.PreviousContactId = aws.String(_connectPreviousContactId)
	}
	if len(_connectQuickConnectId) > 0 {
		input.QuickConnectId = aws.String(_connectQuickConnectId)
	}
	if len(_connectReferences) > 0 {
		if err := assignInputField(input, "References", _connectReferences); err != nil {
			log.Errorf("invalid --references: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}
	if len(_connectScheduledTime) > 0 {
		if err := assignInputField(input, "ScheduledTime", _connectScheduledTime); err != nil {
			log.Errorf("invalid --scheduled-time: %s", err.Error())
			return
		}
	}
	if len(_connectSegmentAttributes) > 0 {
		if err := assignInputField(input, "SegmentAttributes", _connectSegmentAttributes); err != nil {
			log.Errorf("invalid --segment-attributes: %s", err.Error())
			return
		}
	}
	if len(_connectTaskTemplateId) > 0 {
		input.TaskTemplateId = aws.String(_connectTaskTemplateId)
	}

	if resp, err := client.StartTaskContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts executing a published test case.
func connect_StartTestCaseExecution(cfg aws.Config, client *connect.Client) {
	input := &connect.StartTestCaseExecutionInput{
		// InstanceId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.StartTestCaseExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Places an inbound in-app, web, or video call to a contact, and then initiates
// the flow. It performs the actions in the flow that are specified (in
// ContactFlowId) and present in the Amazon Connect instance (specified as
// InstanceId).
func connect_StartWebRTCContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StartWebRTCContactInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
		// ParticipantDetails: *types.ParticipantDetails, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectParticipantDetails) > 0 {
		if err := assignInputField(input, "ParticipantDetails", _connectParticipantDetails); err != nil {
			log.Errorf("invalid --participant-details: %s", err.Error())
			return
		}
	}
	if len(_connectAllowedCapabilities) > 0 {
		if err := assignInputField(input, "AllowedCapabilities", _connectAllowedCapabilities); err != nil {
			log.Errorf("invalid --allowed-capabilities: %s", err.Error())
			return
		}
	}
	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectReferences) > 0 {
		if err := assignInputField(input, "References", _connectReferences); err != nil {
			log.Errorf("invalid --references: %s", err.Error())
			return
		}
	}
	if len(_connectRelatedContactId) > 0 {
		input.RelatedContactId = aws.String(_connectRelatedContactId)
	}

	if resp, err := client.StartWebRTCContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ends the specified contact. Use this API to stop queued callbacks. It does not
// work for voice contacts that use the following initiation methods:
//
// - DISCONNECT
//
// - TRANSFER
//
// - QUEUE_TRANSFER
//
// - EXTERNAL_OUTBOUND
//
// - MONITOR
//
// Chat and task contacts can be terminated in any state, regardless of initiation
// method.
func connect_StopContact(cfg aws.Config, client *connect.Client) {
	input := &connect.StopContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDisconnectReason) > 0 {
		if err := assignInputField(input, "DisconnectReason", _connectDisconnectReason); err != nil {
			log.Errorf("invalid --disconnect-reason: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops in-flight message processing for an ongoing chat session.
func connect_StopContactMediaProcessing(cfg aws.Config, client *connect.Client) {
	input := &connect.StopContactMediaProcessingInput{}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.StopContactMediaProcessing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops recording a call when a contact is being recorded. StopContactRecording
// is a one-time action. If you use StopContactRecording to stop recording an
// ongoing call, you can't use StartContactRecording to restart it. For scenarios
// where the recording has started and you want to suspend it for sensitive
// information (for example, to collect a credit card number), and then restart it,
// use SuspendContactRecording and ResumeContactRecording.
//
// Only voice recordings are supported at this time.
func connect_StopContactRecording(cfg aws.Config, client *connect.Client) {
	input := &connect.StopContactRecordingInput{
		// ContactId: *string, // Required
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactRecordingType) > 0 {
		if err := assignInputField(input, "ContactRecordingType", _connectContactRecordingType); err != nil {
			log.Errorf("invalid --contact-recording-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopContactRecording(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ends message streaming on a specified contact. To restart message streaming on
// that contact, call the [StartContactStreaming]API.
//
// [StartContactStreaming]: https://docs.aws.amazon.com/connect/latest/APIReference/API_StartContactStreaming.html
func connect_StopContactStreaming(cfg aws.Config, client *connect.Client) {
	input := &connect.StopContactStreamingInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// StreamingId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectStreamingId) > 0 {
		input.StreamingId = aws.String(_connectStreamingId)
	}

	if resp, err := client.StopContactStreaming(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running test execution.
func connect_StopTestCaseExecution(cfg aws.Config, client *connect.Client) {
	input := &connect.StopTestCaseExecutionInput{
		// InstanceId: *string, // Required
		// TestCaseExecutionId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseExecutionId) > 0 {
		input.TestCaseExecutionId = aws.String(_connectTestCaseExecutionId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.StopTestCaseExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a contact evaluation in the specified Amazon Connect instance. Answers
// included in the request are merged with existing answers for the given
// evaluation. If no answers or notes are passed, the evaluation is submitted with
// the existing answers and notes. You can delete an answer or note by passing an
// empty object ( {} ) to the question identifier.
//
// If a contact evaluation is already in submitted state, this operation will
// trigger a resubmission.
func connect_SubmitContactEvaluation(cfg aws.Config, client *connect.Client) {
	input := &connect.SubmitContactEvaluationInput{
		// EvaluationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationId) > 0 {
		input.EvaluationId = aws.String(_connectEvaluationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAnswers) > 0 {
		if err := assignInputField(input, "Answers", _connectAnswers); err != nil {
			log.Errorf("invalid --answers: %s", err.Error())
			return
		}
	}
	if len(_connectNotes) > 0 {
		if err := assignInputField(input, "Notes", _connectNotes); err != nil {
			log.Errorf("invalid --notes: %s", err.Error())
			return
		}
	}
	if len(_connectSubmittedBy) > 0 {
		if err := assignInputField(input, "SubmittedBy", _connectSubmittedBy); err != nil {
			log.Errorf("invalid --submitted-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitContactEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When a contact is being recorded, this API suspends recording whatever is
// selected in the flow configuration: call (IVR or agent), screen, or both. If
// only call recording or only screen recording is enabled, then it would be
// suspended. For example, you might suspend the screen recording while collecting
// sensitive information, such as a credit card number. Then use [ResumeContactRecording]to restart
// recording the screen.
//
// The period of time that the recording is suspended is filled with silence in
// the final recording.
//
// Voice (IVR, agent) and screen recordings are supported.
//
// [ResumeContactRecording]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ResumeContactRecording.html
func connect_SuspendContactRecording(cfg aws.Config, client *connect.Client) {
	input := &connect.SuspendContactRecordingInput{
		// ContactId: *string, // Required
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactRecordingType) > 0 {
		if err := assignInputField(input, "ContactRecordingType", _connectContactRecordingType); err != nil {
			log.Errorf("invalid --contact-recording-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.SuspendContactRecording(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the contact resource. For more information about
// this API is used, see [Set up granular billing for a detailed view of your Amazon Connect usage].
//
// [Set up granular billing for a detailed view of your Amazon Connect usage]: https://docs.aws.amazon.com/connect/latest/adminguide/granular-billing.html
func connect_TagContact(cfg aws.Config, client *connect.Client) {
	input := &connect.TagContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource.
// Some of the supported resource types are agents, routing profiles, queues,
// quick connects, flows, agent statuses, hours of operation, phone numbers,
// security profiles, and task templates. For a complete list, see [Tagging resources in Amazon Connect].
//
// For sample policies that use tags, see [Amazon Connect Identity-Based Policy Examples] in the Amazon Connect Administrator
// Guide.
//
// [Amazon Connect Identity-Based Policy Examples]: https://docs.aws.amazon.com/connect/latest/adminguide/security_iam_id-based-policy-examples.html
// [Tagging resources in Amazon Connect]: https://docs.aws.amazon.com/connect/latest/adminguide/tagging.html
func connect_TagResource(cfg aws.Config, client *connect.Client) {
	input := &connect.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_connectResourceArn) > 0 {
		input.ResourceArn = aws.String(_connectResourceArn)
	}
	if len(_connectTags) > 0 {
		if err := assignInputField(input, "Tags", _connectTags); err != nil {
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

// Transfers TASK or EMAIL contacts from one agent or queue to another agent or
// queue at any point after a contact is created. You can transfer a contact to
// another queue by providing the flow which orchestrates the contact to the
// destination queue. This gives you more control over contact handling and helps
// you adhere to the service level agreement (SLA) guaranteed to your customers.
//
// Note the following requirements:
//
// - Transfer is only supported for TASK and EMAIL contacts.
//
// - Do not use both QueueId and UserId in the same call.
//
// - The following flow types are supported: Inbound flow, Transfer to agent
// flow, and Transfer to queue flow.
//
// - The TransferContact API can be called only on active contacts.
//
// - A contact cannot be transferred more than 11 times.
func connect_TransferContact(cfg aws.Config, client *connect.Client) {
	input := &connect.TransferContactInput{
		// ContactFlowId: *string, // Required
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.TransferContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the contact resource. For more information
// about this API is used, see [Set up granular billing for a detailed view of your Amazon Connect usage].
//
// [Set up granular billing for a detailed view of your Amazon Connect usage]: https://docs.aws.amazon.com/connect/latest/adminguide/granular-billing.html
func connect_UntagContact(cfg aws.Config, client *connect.Client) {
	input := &connect.UntagContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _connectTagKeys...)
	}

	if resp, err := client.UntagContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the specified resource.
func connect_UntagResource(cfg aws.Config, client *connect.Client) {
	input := &connect.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_connectResourceArn) > 0 {
		input.ResourceArn = aws.String(_connectResourceArn)
	}
	if len(_connectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _connectTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates agent status.
func connect_UpdateAgentStatus(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateAgentStatusInput{
		// AgentStatusId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAgentStatusId) > 0 {
		input.AgentStatusId = aws.String(_connectAgentStatusId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectDisplayOrder) > 0 {
		if err := assignInputField(input, "DisplayOrder", _connectDisplayOrder); err != nil {
			log.Errorf("invalid --display-order: %s", err.Error())
			return
		}
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectResetOrderNumber) > 0 {
		if err := assignInputField(input, "ResetOrderNumber", _connectResetOrderNumber); err != nil {
			log.Errorf("invalid --reset-order-number: %s", err.Error())
			return
		}
	}
	if len(_connectState) > 0 {
		if err := assignInputField(input, "State", _connectState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change. To
// request access to this API, contact Amazon Web Services Support.
//
// Updates the selected authentication profile.
func connect_UpdateAuthenticationProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateAuthenticationProfileInput{
		// AuthenticationProfileId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAuthenticationProfileId) > 0 {
		input.AuthenticationProfileId = aws.String(_connectAuthenticationProfileId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAllowedIps) > 0 {
		input.AllowedIps = append([]string(nil), _connectAllowedIps...)
	}
	if len(_connectBlockedIps) > 0 {
		input.BlockedIps = append([]string(nil), _connectBlockedIps...)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectPeriodicSessionDuration) > 0 {
		if err := assignInputField(input, "PeriodicSessionDuration", _connectPeriodicSessionDuration); err != nil {
			log.Errorf("invalid --periodic-session-duration: %s", err.Error())
			return
		}
	}
	if len(_connectSessionInactivityDuration) > 0 {
		if err := assignInputField(input, "SessionInactivityDuration", _connectSessionInactivityDuration); err != nil {
			log.Errorf("invalid --session-inactivity-duration: %s", err.Error())
			return
		}
	}
	if len(_connectSessionInactivityHandlingEnabled) > 0 {
		if err := assignInputField(input, "SessionInactivityHandlingEnabled", _connectSessionInactivityHandlingEnabled); err != nil {
			log.Errorf("invalid --session-inactivity-handling-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAuthenticationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Adds or updates user-defined contact information associated with the specified
// contact. At least one field to be updated must be present in the request.
//
// You can add or update user-defined contact information for both ongoing and
// completed contacts.
func connect_UpdateContact(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectCustomerEndpoint) > 0 {
		if err := assignInputField(input, "CustomerEndpoint", _connectCustomerEndpoint); err != nil {
			log.Errorf("invalid --customer-endpoint: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectQueueInfo) > 0 {
		if err := assignInputField(input, "QueueInfo", _connectQueueInfo); err != nil {
			log.Errorf("invalid --queue-info: %s", err.Error())
			return
		}
	}
	if len(_connectReferences) > 0 {
		if err := assignInputField(input, "References", _connectReferences); err != nil {
			log.Errorf("invalid --references: %s", err.Error())
			return
		}
	}
	if len(_connectSegmentAttributes) > 0 {
		if err := assignInputField(input, "SegmentAttributes", _connectSegmentAttributes); err != nil {
			log.Errorf("invalid --segment-attributes: %s", err.Error())
			return
		}
	}
	if len(_connectSystemEndpoint) > 0 {
		if err := assignInputField(input, "SystemEndpoint", _connectSystemEndpoint); err != nil {
			log.Errorf("invalid --system-endpoint: %s", err.Error())
			return
		}
	}
	if len(_connectUserInfo) > 0 {
		if err := assignInputField(input, "UserInfo", _connectUserInfo); err != nil {
			log.Errorf("invalid --user-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates user-defined contact attributes associated with the
// specified contact.
//
// You can create or update user-defined attributes for both ongoing and completed
// contacts. For example, while the call is active, you can update the customer's
// name or the reason the customer called. You can add notes about steps that the
// agent took during the call that display to the next agent that takes the call.
// You can also update attributes for a contact using data from your CRM
// application and save the data with the contact in Amazon Connect. You could also
// flag calls for additional analysis, such as legal review or to identify abusive
// callers.
//
// Contact attributes are available in Amazon Connect for 24 months, and are then
// deleted. For information about contact record retention and the maximum size of
// the contact record attributes section, see [Feature specifications]in the Amazon Connect Administrator
// Guide.
//
// [Feature specifications]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits
func connect_UpdateContactAttributes(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactAttributesInput{
		// Attributes: map[string]string, // Required
		// InitialContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectInitialContactId) > 0 {
		input.InitialContactId = aws.String(_connectInitialContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.UpdateContactAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates details about a contact evaluation in the specified Amazon Connect
// instance. A contact evaluation must be in draft state. Answers included in the
// request are merged with existing answers for the given evaluation. An answer or
// note can be deleted by passing an empty object ( {} ) to the question
// identifier.
func connect_UpdateContactEvaluation(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactEvaluationInput{
		// EvaluationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEvaluationId) > 0 {
		input.EvaluationId = aws.String(_connectEvaluationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectAnswers) > 0 {
		if err := assignInputField(input, "Answers", _connectAnswers); err != nil {
			log.Errorf("invalid --answers: %s", err.Error())
			return
		}
	}
	if len(_connectNotes) > 0 {
		if err := assignInputField(input, "Notes", _connectNotes); err != nil {
			log.Errorf("invalid --notes: %s", err.Error())
			return
		}
	}
	if len(_connectUpdatedBy) > 0 {
		if err := assignInputField(input, "UpdatedBy", _connectUpdatedBy); err != nil {
			log.Errorf("invalid --updated-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContactEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified flow.
// You can also create and update flows using the [Amazon Connect Flow language].
//
// Use the $SAVED alias in the request to describe the SAVED content of a Flow.
// For example, arn:aws:.../contact-flow/{id}:$SAVED . After a flow is published,
// $SAVED needs to be supplied to view saved content that has not been published.
//
// [Amazon Connect Flow language]: https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html
func connect_UpdateContactFlowContent(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactFlowContentInput{
		// ContactFlowId: *string, // Required
		// Content: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.UpdateContactFlowContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates metadata about specified flow.
func connect_UpdateContactFlowMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactFlowMetadataInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowState) > 0 {
		if err := assignInputField(input, "ContactFlowState", _connectContactFlowState); err != nil {
			log.Errorf("invalid --contact-flow-state: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateContactFlowMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specific Aliases metadata, including the version it’s tied to, it’s
// name, and description.
func connect_UpdateContactFlowModuleAlias(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactFlowModuleAliasInput{
		// AliasId: *string, // Required
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectAliasId) > 0 {
		input.AliasId = aws.String(_connectAliasId)
	}
	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContactFlowModuleVersion) > 0 {
		if err := assignInputField(input, "ContactFlowModuleVersion", _connectContactFlowModuleVersion); err != nil {
			log.Errorf("invalid --contact-flow-module-version: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateContactFlowModuleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates specified flow module for the specified Amazon Connect instance.
// Use the $SAVED alias in the request to describe the SAVED content of a Flow.
// For example, arn:aws:.../contact-flow/{id}:$SAVED . After a flow is published,
// $SAVED needs to be supplied to view saved content that has not been published.
func connect_UpdateContactFlowModuleContent(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactFlowModuleContentInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectSettings) > 0 {
		input.Settings = aws.String(_connectSettings)
	}

	if resp, err := client.UpdateContactFlowModuleContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates metadata about specified flow module.
func connect_UpdateContactFlowModuleMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactFlowModuleMetadataInput{
		// ContactFlowModuleId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowModuleId) > 0 {
		input.ContactFlowModuleId = aws.String(_connectContactFlowModuleId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectState) > 0 {
		if err := assignInputField(input, "State", _connectState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContactFlowModuleMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The name of the flow.
// You can also create and update flows using the [Amazon Connect Flow language].
//
// [Amazon Connect Flow language]: https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html
func connect_UpdateContactFlowName(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactFlowNameInput{
		// ContactFlowId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateContactFlowName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates routing priority and age on the contact (QueuePriority and
// QueueTimeAdjustmentInSeconds). These properties can be used to change a
// customer's position in the queue. For example, you can move a contact to the
// back of the queue by setting a lower routing priority relative to other contacts
// in queue; or you can move a contact to the front of the queue by increasing the
// routing age which will make the contact look artificially older and therefore
// higher up in the first-in-first-out routing order. Note that adjusting the
// routing age of a contact affects only its position in queue, and not its actual
// queue wait time as reported through metrics. These properties can also be
// updated by using [the Set routing priority / age flow block].
//
// Either QueuePriority or QueueTimeAdjustmentInSeconds should be provided within
// the request body, but not both.
//
// [the Set routing priority / age flow block]: https://docs.aws.amazon.com/connect/latest/adminguide/change-routing-priority.html
func connect_UpdateContactRoutingData(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactRoutingDataInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueuePriority) > 0 {
		if err := assignInputField(input, "QueuePriority", _connectQueuePriority); err != nil {
			log.Errorf("invalid --queue-priority: %s", err.Error())
			return
		}
	}
	if len(_connectQueueTimeAdjustmentSeconds) > 0 {
		if err := assignInputField(input, "QueueTimeAdjustmentSeconds", _connectQueueTimeAdjustmentSeconds); err != nil {
			log.Errorf("invalid --queue-time-adjustment-seconds: %s", err.Error())
			return
		}
	}
	if len(_connectRoutingCriteria) > 0 {
		if err := assignInputField(input, "RoutingCriteria", _connectRoutingCriteria); err != nil {
			log.Errorf("invalid --routing-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContactRoutingData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the scheduled time of a task contact that is already scheduled.
func connect_UpdateContactSchedule(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateContactScheduleInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
		// ScheduledTime: *time.Time, // Required
	}

	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectScheduledTime) > 0 {
		if err := assignInputField(input, "ScheduledTime", _connectScheduledTime); err != nil {
			log.Errorf("invalid --scheduled-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContactSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates all properties for an attribute using all properties from
// CreateDataTableAttribute. There are no other granular update endpoints. It does
// not act as a patch operation - all properties must be provided. System managed
// attributes are not mutable by customers. Changing an attribute's validation does
// not invalidate existing values since validation only runs when values are
// created or updated.
func connect_UpdateDataTableAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateDataTableAttributeInput{
		// AttributeName: *string, // Required
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// ValueType: types.DataTableAttributeValueType, // Required
	}

	if len(_connectAttributeName) > 0 {
		input.AttributeName = aws.String(_connectAttributeName)
	}
	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectValueType) > 0 {
		if err := assignInputField(input, "ValueType", _connectValueType); err != nil {
			log.Errorf("invalid --value-type: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectPrimary) > 0 {
		if err := assignInputField(input, "Primary", _connectPrimary); err != nil {
			log.Errorf("invalid --primary: %s", err.Error())
			return
		}
	}
	if len(_connectValidation) > 0 {
		if err := assignInputField(input, "Validation", _connectValidation); err != nil {
			log.Errorf("invalid --validation: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataTableAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the metadata properties of a data table. Accepts all fields similar to
// CreateDataTable, except for fields and tags. There are no other granular update
// endpoints. It does not act as a patch operation - all properties must be
// provided or defaults will be used. Fields follow the same requirements as
// CreateDataTable.
func connect_UpdateDataTableMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateDataTableMetadataInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// TimeZone: *string, // Required
		// ValueLockLevel: types.DataTableLockLevel, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectTimeZone) > 0 {
		input.TimeZone = aws.String(_connectTimeZone)
	}
	if len(_connectValueLockLevel) > 0 {
		if err := assignInputField(input, "ValueLockLevel", _connectValueLockLevel); err != nil {
			log.Errorf("invalid --value-lock-level: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}

	if resp, err := client.UpdateDataTableMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the primary values for a record. This operation affects all existing
// values that are currently associated to the record and its primary values. Users
// that have restrictions on attributes and/or primary values are not authorized to
// use this endpoint. The combination of new primary values must be unique within
// the table.
func connect_UpdateDataTablePrimaryValues(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateDataTablePrimaryValuesInput{
		// DataTableId: *string, // Required
		// InstanceId: *string, // Required
		// LockVersion: *types.DataTableLockVersion, // Required
		// NewPrimaryValues: []types.PrimaryValue, // Required
		// PrimaryValues: []types.PrimaryValue, // Required
	}

	if len(_connectDataTableId) > 0 {
		input.DataTableId = aws.String(_connectDataTableId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectLockVersion) > 0 {
		if err := assignInputField(input, "LockVersion", _connectLockVersion); err != nil {
			log.Errorf("invalid --lock-version: %s", err.Error())
			return
		}
	}
	if len(_connectNewPrimaryValues) > 0 {
		if err := assignInputField(input, "NewPrimaryValues", _connectNewPrimaryValues); err != nil {
			log.Errorf("invalid --new-primary-values: %s", err.Error())
			return
		}
	}
	if len(_connectPrimaryValues) > 0 {
		if err := assignInputField(input, "PrimaryValues", _connectPrimaryValues); err != nil {
			log.Errorf("invalid --primary-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataTablePrimaryValues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an email address metadata. For more information about email addresses,
// see [Create email addresses]in the Amazon Connect Administrator Guide.
//
// [Create email addresses]: https://docs.aws.amazon.com/connect/latest/adminguide/create-email-address1.html
func connect_UpdateEmailAddressMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateEmailAddressMetadataInput{
		// EmailAddressId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectEmailAddressId) > 0 {
		input.EmailAddressId = aws.String(_connectEmailAddressId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectDisplayName) > 0 {
		input.DisplayName = aws.String(_connectDisplayName)
	}

	if resp, err := client.UpdateEmailAddressMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates details about a specific evaluation form version in the specified
// Amazon Connect instance. Question and section identifiers cannot be duplicated
// within the same evaluation form.
//
// This operation does not support partial updates. Instead it does a full update
// of evaluation form content.
func connect_UpdateEvaluationForm(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateEvaluationFormInput{
		// EvaluationFormId: *string, // Required
		// EvaluationFormVersion: int32, // Required
		// InstanceId: *string, // Required
		// Items: []types.EvaluationFormItem, // Required
		// Title: *string, // Required
	}

	if len(_connectEvaluationFormId) > 0 {
		input.EvaluationFormId = aws.String(_connectEvaluationFormId)
	}
	if len(_connectEvaluationFormVersion) > 0 {
		if err := assignInputField(input, "EvaluationFormVersion", _connectEvaluationFormVersion); err != nil {
			log.Errorf("invalid --evaluation-form-version: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectItems) > 0 {
		if err := assignInputField(input, "Items", _connectItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}
	if len(_connectTitle) > 0 {
		input.Title = aws.String(_connectTitle)
	}
	if len(_connectAsDraft) > 0 {
		if err := assignInputField(input, "AsDraft", _connectAsDraft); err != nil {
			log.Errorf("invalid --as-draft: %s", err.Error())
			return
		}
	}
	if len(_connectAutoEvaluationConfiguration) > 0 {
		if err := assignInputField(input, "AutoEvaluationConfiguration", _connectAutoEvaluationConfiguration); err != nil {
			log.Errorf("invalid --auto-evaluation-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectCreateNewVersion) > 0 {
		if err := assignInputField(input, "CreateNewVersion", _connectCreateNewVersion); err != nil {
			log.Errorf("invalid --create-new-version: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectLanguageConfiguration) > 0 {
		if err := assignInputField(input, "LanguageConfiguration", _connectLanguageConfiguration); err != nil {
			log.Errorf("invalid --language-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectReviewConfiguration) > 0 {
		if err := assignInputField(input, "ReviewConfiguration", _connectReviewConfiguration); err != nil {
			log.Errorf("invalid --review-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectScoringStrategy) > 0 {
		if err := assignInputField(input, "ScoringStrategy", _connectScoringStrategy); err != nil {
			log.Errorf("invalid --scoring-strategy: %s", err.Error())
			return
		}
	}
	if len(_connectTargetConfiguration) > 0 {
		if err := assignInputField(input, "TargetConfiguration", _connectTargetConfiguration); err != nil {
			log.Errorf("invalid --target-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEvaluationForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the hours of operation.
func connect_UpdateHoursOfOperation(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateHoursOfOperationInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectConfig) > 0 {
		if err := assignInputField(input, "Config", _connectConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectTimeZone) > 0 {
		input.TimeZone = aws.String(_connectTimeZone)
	}

	if resp, err := client.UpdateHoursOfOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the hours of operation override.
func connect_UpdateHoursOfOperationOverride(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateHoursOfOperationOverrideInput{
		// HoursOfOperationId: *string, // Required
		// HoursOfOperationOverrideId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectHoursOfOperationOverrideId) > 0 {
		input.HoursOfOperationOverrideId = aws.String(_connectHoursOfOperationOverrideId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectConfig) > 0 {
		if err := assignInputField(input, "Config", _connectConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectEffectiveFrom) > 0 {
		input.EffectiveFrom = aws.String(_connectEffectiveFrom)
	}
	if len(_connectEffectiveTill) > 0 {
		input.EffectiveTill = aws.String(_connectEffectiveTill)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectOverrideType) > 0 {
		if err := assignInputField(input, "OverrideType", _connectOverrideType); err != nil {
			log.Errorf("invalid --override-type: %s", err.Error())
			return
		}
	}
	if len(_connectRecurrenceConfig) > 0 {
		if err := assignInputField(input, "RecurrenceConfig", _connectRecurrenceConfig); err != nil {
			log.Errorf("invalid --recurrence-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateHoursOfOperationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Updates the value for the specified attribute type.
func connect_UpdateInstanceAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateInstanceAttributeInput{
		// AttributeType: types.InstanceAttributeType, // Required
		// InstanceId: *string, // Required
		// Value: *string, // Required
	}

	if len(_connectAttributeType) > 0 {
		if err := assignInputField(input, "AttributeType", _connectAttributeType); err != nil {
			log.Errorf("invalid --attribute-type: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectValue) > 0 {
		input.Value = aws.String(_connectValue)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.UpdateInstanceAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is in preview release for Amazon Connect and is subject to change.
// Updates an existing configuration for a resource type. This API is idempotent.
func connect_UpdateInstanceStorageConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateInstanceStorageConfigInput{
		// AssociationId: *string, // Required
		// InstanceId: *string, // Required
		// ResourceType: types.InstanceStorageResourceType, // Required
		// StorageConfig: *types.InstanceStorageConfig, // Required
	}

	if len(_connectAssociationId) > 0 {
		input.AssociationId = aws.String(_connectAssociationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _connectResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_connectStorageConfig) > 0 {
		if err := assignInputField(input, "StorageConfig", _connectStorageConfig); err != nil {
			log.Errorf("invalid --storage-config: %s", err.Error())
			return
		}
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}

	if resp, err := client.UpdateInstanceStorageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the localized content of an existing notification. This operation
// applies to all users for whom the notification was sent.
func connect_UpdateNotificationContent(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateNotificationContentInput{
		// Content: map[string]string, // Required
		// InstanceId: *string, // Required
		// NotificationId: *string, // Required
	}

	if len(_connectContent) > 0 {
		if err := assignInputField(input, "Content", _connectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectNotificationId) > 0 {
		input.NotificationId = aws.String(_connectNotificationId)
	}

	if resp, err := client.UpdateNotificationContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instructs Amazon Connect to resume the authentication process. The subsequent
// actions depend on the request body contents:
//
// - If a code is provided: Connect retrieves the identity information from
// Amazon Cognito and imports it into Connect Customer Profiles.
//
// - If an error is provided: The error branch of the Authenticate Customer
// block is executed.
//
// The API returns a success response to acknowledge the request. However, the
// interaction and exchange of identity information occur asynchronously after the
// response is returned.
func connect_UpdateParticipantAuthentication(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateParticipantAuthenticationInput{
		// InstanceId: *string, // Required
		// State: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectState) > 0 {
		input.State = aws.String(_connectState)
	}
	if len(_connectCode) > 0 {
		input.Code = aws.String(_connectCode)
	}
	if len(_connectError) > 0 {
		input.Error = aws.String(_connectError)
	}
	if len(_connectErrorDescription) > 0 {
		input.ErrorDescription = aws.String(_connectErrorDescription)
	}

	if resp, err := client.UpdateParticipantAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates timeouts for when human chat participants are to be considered idle,
// and when agents are automatically disconnected from a chat due to idleness. You
// can set four timers:
//
// - Customer idle timeout
//
// - Customer auto-disconnect timeout
//
// - Agent idle timeout
//
// - Agent auto-disconnect timeout
//
// For more information about how chat timeouts work, see [Set up chat timeouts for human participants].
//
// [Set up chat timeouts for human participants]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-chat-timeouts.html
func connect_UpdateParticipantRoleConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateParticipantRoleConfigInput{
		// ChannelConfiguration: types.UpdateParticipantRoleConfigChannelInfo, // Required
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectChannelConfiguration) > 0 {
		if err := assignInputField(input, "ChannelConfiguration", _connectChannelConfiguration); err != nil {
			log.Errorf("invalid --channel-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectContactId) > 0 {
		input.ContactId = aws.String(_connectContactId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.UpdateParticipantRoleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates your claimed phone number from its current Amazon Connect instance or
// traffic distribution group to another Amazon Connect instance or traffic
// distribution group in the same Amazon Web Services Region.
//
// After using this API, you must verify that the phone number is attached to the
// correct flow in the target instance or traffic distribution group. You need to
// do this because the API switches only the phone number to a new instance or
// traffic distribution group. It doesn't migrate the flow configuration of the
// phone number, too.
//
// You can call [DescribePhoneNumber] API to verify the status of a previous [UpdatePhoneNumber] operation.
//
// [UpdatePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdatePhoneNumber.html
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
func connect_UpdatePhoneNumber(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdatePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_connectPhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_connectPhoneNumberId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTargetArn) > 0 {
		input.TargetArn = aws.String(_connectTargetArn)
	}

	if resp, err := client.UpdatePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a phone number’s metadata.
// To verify the status of a previous UpdatePhoneNumberMetadata operation, call
// the [DescribePhoneNumber]API.
//
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
func connect_UpdatePhoneNumberMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdatePhoneNumberMetadataInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_connectPhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_connectPhoneNumberId)
	}
	if len(_connectClientToken) > 0 {
		input.ClientToken = aws.String(_connectClientToken)
	}
	if len(_connectPhoneNumberDescription) > 0 {
		input.PhoneNumberDescription = aws.String(_connectPhoneNumberDescription)
	}

	if resp, err := client.UpdatePhoneNumberMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a predefined attribute for the specified Amazon Connect instance. A
// predefined attribute is made up of a name and a value.
//
// For the predefined attributes per instance quota, see [Amazon Connect quotas].
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Update routing proficiency (for example, agent certification) that has
// predefined values (for example, a list of possible certifications). For more
// information, see [Create predefined attributes for routing contacts to agents].
//
// - Update an attribute for business unit name that has a list of predefined
// business unit names used in your organization. This is a use case where
// information for a contact varies between transfers or conferences. For more
// information, see [Use contact segment attributes].
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Use contact segment attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/use-contact-segment-attributes.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
func connect_UpdatePredefinedAttribute(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdatePredefinedAttributeInput{
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectAttributeConfiguration) > 0 {
		if err := assignInputField(input, "AttributeConfiguration", _connectAttributeConfiguration); err != nil {
			log.Errorf("invalid --attribute-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectPurposes) > 0 {
		input.Purposes = append([]string(nil), _connectPurposes...)
	}
	if len(_connectValues) > 0 {
		if err := assignInputField(input, "Values", _connectValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePredefinedAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a prompt.
func connect_UpdatePrompt(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdatePromptInput{
		// InstanceId: *string, // Required
		// PromptId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPromptId) > 0 {
		input.PromptId = aws.String(_connectPromptId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectS3Uri) > 0 {
		input.S3Uri = aws.String(_connectS3Uri)
	}

	if resp, err := client.UpdatePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the hours of operation for the specified queue.
func connect_UpdateQueueHoursOfOperation(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQueueHoursOfOperationInput{
		// HoursOfOperationId: *string, // Required
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_connectHoursOfOperationId) > 0 {
		input.HoursOfOperationId = aws.String(_connectHoursOfOperationId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}

	if resp, err := client.UpdateQueueHoursOfOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the maximum number of contacts allowed in a queue before it is
// considered full.
func connect_UpdateQueueMaxContacts(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQueueMaxContactsInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectMaxContacts) > 0 {
		if err := assignInputField(input, "MaxContacts", _connectMaxContacts); err != nil {
			log.Errorf("invalid --max-contacts: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueueMaxContacts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and description of a queue. At least Name or Description must
// be provided.
func connect_UpdateQueueName(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQueueNameInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateQueueName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the outbound caller ID name, number, and outbound whisper flow for a
// specified queue.
//
// - If the phone number is claimed to a traffic distribution group that was
// created in the same Region as the Amazon Connect instance where you are calling
// this API, then you can use a full phone number ARN or a UUID for
// OutboundCallerIdNumberId . However, if the phone number is claimed to a
// traffic distribution group that is in one Region, and you are calling this API
// from an instance in another Amazon Web Services Region that is associated with
// the traffic distribution group, you must provide a full phone number ARN. If a
// UUID is provided in this scenario, you will receive a
// ResourceNotFoundException .
//
// - Only use the phone number ARN format that doesn't contain instance in the
// path, for example, arn:aws:connect:us-east-1:1234567890:phone-number/uuid .
// This is the same ARN format that is returned when you call the [ListPhoneNumbersV2]API.
//
// - If you plan to use IAM policies to allow/deny access to this API for phone
// number resources claimed to a traffic distribution group, see [Allow or Deny queue API actions for phone numbers in a replica Region].
//
// [ListPhoneNumbersV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html
// [Allow or Deny queue API actions for phone numbers in a replica Region]: https://docs.aws.amazon.com/connect/latest/adminguide/security_iam_resource-level-policy-examples.html#allow-deny-queue-actions-replica-region
func connect_UpdateQueueOutboundCallerConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQueueOutboundCallerConfigInput{
		// InstanceId: *string, // Required
		// OutboundCallerConfig: *types.OutboundCallerConfig, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectOutboundCallerConfig) > 0 {
		if err := assignInputField(input, "OutboundCallerConfig", _connectOutboundCallerConfig); err != nil {
			log.Errorf("invalid --outbound-caller-config: %s", err.Error())
			return
		}
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}

	if resp, err := client.UpdateQueueOutboundCallerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the outbound email address Id for a specified queue.
func connect_UpdateQueueOutboundEmailConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQueueOutboundEmailConfigInput{
		// InstanceId: *string, // Required
		// OutboundEmailConfig: *types.OutboundEmailConfig, // Required
		// QueueId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectOutboundEmailConfig) > 0 {
		if err := assignInputField(input, "OutboundEmailConfig", _connectOutboundEmailConfig); err != nil {
			log.Errorf("invalid --outbound-email-config: %s", err.Error())
			return
		}
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}

	if resp, err := client.UpdateQueueOutboundEmailConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of the queue.
func connect_UpdateQueueStatus(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQueueStatusInput{
		// InstanceId: *string, // Required
		// QueueId: *string, // Required
		// Status: types.QueueStatus, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueId) > 0 {
		input.QueueId = aws.String(_connectQueueId)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueueStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration settings for the specified quick connect.
func connect_UpdateQuickConnectConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQuickConnectConfigInput{
		// InstanceId: *string, // Required
		// QuickConnectConfig: *types.QuickConnectConfig, // Required
		// QuickConnectId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQuickConnectConfig) > 0 {
		if err := assignInputField(input, "QuickConnectConfig", _connectQuickConnectConfig); err != nil {
			log.Errorf("invalid --quick-connect-config: %s", err.Error())
			return
		}
	}
	if len(_connectQuickConnectId) > 0 {
		input.QuickConnectId = aws.String(_connectQuickConnectId)
	}

	if resp, err := client.UpdateQuickConnectConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and description of a quick connect. The request accepts the
// following data in JSON format. At least Name or Description must be provided.
func connect_UpdateQuickConnectName(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateQuickConnectNameInput{
		// InstanceId: *string, // Required
		// QuickConnectId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQuickConnectId) > 0 {
		input.QuickConnectId = aws.String(_connectQuickConnectId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateQuickConnectName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Whether agents with this routing profile will have their routing order
// calculated based on time since their last inbound contact or longest idle time.
func connect_UpdateRoutingProfileAgentAvailabilityTimer(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateRoutingProfileAgentAvailabilityTimerInput{
		// AgentAvailabilityTimer: types.AgentAvailabilityTimer, // Required
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectAgentAvailabilityTimer) > 0 {
		if err := assignInputField(input, "AgentAvailabilityTimer", _connectAgentAvailabilityTimer); err != nil {
			log.Errorf("invalid --agent-availability-timer: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}

	if resp, err := client.UpdateRoutingProfileAgentAvailabilityTimer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the channels that agents can handle in the Contact Control Panel (CCP)
// for a routing profile.
func connect_UpdateRoutingProfileConcurrency(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateRoutingProfileConcurrencyInput{
		// InstanceId: *string, // Required
		// MediaConcurrencies: []types.MediaConcurrency, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectMediaConcurrencies) > 0 {
		if err := assignInputField(input, "MediaConcurrencies", _connectMediaConcurrencies); err != nil {
			log.Errorf("invalid --media-concurrencies: %s", err.Error())
			return
		}
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}

	if resp, err := client.UpdateRoutingProfileConcurrency(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the default outbound queue of a routing profile.
func connect_UpdateRoutingProfileDefaultOutboundQueue(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateRoutingProfileDefaultOutboundQueueInput{
		// DefaultOutboundQueueId: *string, // Required
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectDefaultOutboundQueueId) > 0 {
		input.DefaultOutboundQueueId = aws.String(_connectDefaultOutboundQueueId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}

	if resp, err := client.UpdateRoutingProfileDefaultOutboundQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and description of a routing profile. The request accepts the
// following data in JSON format. At least Name or Description must be provided.
func connect_UpdateRoutingProfileName(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateRoutingProfileNameInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateRoutingProfileName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties associated with a set of queues for a routing profile.
func connect_UpdateRoutingProfileQueues(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateRoutingProfileQueuesInput{
		// InstanceId: *string, // Required
		// QueueConfigs: []types.RoutingProfileQueueConfig, // Required
		// RoutingProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectQueueConfigs) > 0 {
		if err := assignInputField(input, "QueueConfigs", _connectQueueConfigs); err != nil {
			log.Errorf("invalid --queue-configs: %s", err.Error())
			return
		}
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}

	if resp, err := client.UpdateRoutingProfileQueues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a rule for the specified Amazon Connect instance.
// Use the [Rules Function language] to code conditions for the rule.
//
// [Rules Function language]: https://docs.aws.amazon.com/connect/latest/APIReference/connect-rules-language.html
func connect_UpdateRule(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateRuleInput{
		// Actions: []types.RuleAction, // Required
		// Function: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
		// PublishStatus: types.RulePublishStatus, // Required
		// RuleId: *string, // Required
	}

	if len(_connectActions) > 0 {
		if err := assignInputField(input, "Actions", _connectActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_connectFunction) > 0 {
		input.Function = aws.String(_connectFunction)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectPublishStatus) > 0 {
		if err := assignInputField(input, "PublishStatus", _connectPublishStatus); err != nil {
			log.Errorf("invalid --publish-status: %s", err.Error())
			return
		}
	}
	if len(_connectRuleId) > 0 {
		input.RuleId = aws.String(_connectRuleId)
	}

	if resp, err := client.UpdateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a security profile.
// For information about security profiles, see [Security Profiles] in the Amazon Connect
// Administrator Guide. For a mapping of the API name and user interface name of
// the security profile permissions, see [List of security profile permissions].
//
// [Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html
// [List of security profile permissions]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html
func connect_UpdateSecurityProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateSecurityProfileInput{
		// InstanceId: *string, // Required
		// SecurityProfileId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileId) > 0 {
		input.SecurityProfileId = aws.String(_connectSecurityProfileId)
	}
	if len(_connectAllowedAccessControlHierarchyGroupId) > 0 {
		input.AllowedAccessControlHierarchyGroupId = aws.String(_connectAllowedAccessControlHierarchyGroupId)
	}
	if len(_connectAllowedAccessControlTags) > 0 {
		if err := assignInputField(input, "AllowedAccessControlTags", _connectAllowedAccessControlTags); err != nil {
			log.Errorf("invalid --allowed-access-control-tags: %s", err.Error())
			return
		}
	}
	if len(_connectAllowedFlowModules) > 0 {
		if err := assignInputField(input, "AllowedFlowModules", _connectAllowedFlowModules); err != nil {
			log.Errorf("invalid --allowed-flow-modules: %s", err.Error())
			return
		}
	}
	if len(_connectApplications) > 0 {
		if err := assignInputField(input, "Applications", _connectApplications); err != nil {
			log.Errorf("invalid --applications: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectGranularAccessControlConfiguration) > 0 {
		if err := assignInputField(input, "GranularAccessControlConfiguration", _connectGranularAccessControlConfiguration); err != nil {
			log.Errorf("invalid --granular-access-control-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectHierarchyRestrictedResources) > 0 {
		input.HierarchyRestrictedResources = append([]string(nil), _connectHierarchyRestrictedResources...)
	}
	if len(_connectPermissions) > 0 {
		input.Permissions = append([]string(nil), _connectPermissions...)
	}
	if len(_connectTagRestrictedResources) > 0 {
		input.TagRestrictedResources = append([]string(nil), _connectTagRestrictedResources...)
	}

	if resp, err := client.UpdateSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates details about a specific task template in the specified Amazon Connect
// instance. This operation does not support partial updates. Instead it does a
// full update of template content.
func connect_UpdateTaskTemplate(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateTaskTemplateInput{
		// InstanceId: *string, // Required
		// TaskTemplateId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTaskTemplateId) > 0 {
		input.TaskTemplateId = aws.String(_connectTaskTemplateId)
	}
	if len(_connectConstraints) > 0 {
		if err := assignInputField(input, "Constraints", _connectConstraints); err != nil {
			log.Errorf("invalid --constraints: %s", err.Error())
			return
		}
	}
	if len(_connectContactFlowId) > 0 {
		input.ContactFlowId = aws.String(_connectContactFlowId)
	}
	if len(_connectDefaults) > 0 {
		if err := assignInputField(input, "Defaults", _connectDefaults); err != nil {
			log.Errorf("invalid --defaults: %s", err.Error())
			return
		}
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectFields) > 0 {
		if err := assignInputField(input, "Fields", _connectFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectSelfAssignFlowId) > 0 {
		input.SelfAssignFlowId = aws.String(_connectSelfAssignFlowId)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTaskTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates any of the metadata for a test case, such as the name, description, and
// status or content of an existing test case. This API doesn't allow customers to
// update the tags of the test case resource for the specified Amazon Connect
// instance.
func connect_UpdateTestCase(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateTestCaseInput{
		// InstanceId: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectTestCaseId) > 0 {
		input.TestCaseId = aws.String(_connectTestCaseId)
	}
	if len(_connectContent) > 0 {
		input.Content = aws.String(_connectContent)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectEntryPoint) > 0 {
		if err := assignInputField(input, "EntryPoint", _connectEntryPoint); err != nil {
			log.Errorf("invalid --entry-point: %s", err.Error())
			return
		}
	}
	if len(_connectInitializationData) > 0 {
		input.InitializationData = aws.String(_connectInitializationData)
	}
	if len(_connectLastModifiedRegion) > 0 {
		input.LastModifiedRegion = aws.String(_connectLastModifiedRegion)
	}
	if len(_connectLastModifiedTime) > 0 {
		if err := assignInputField(input, "LastModifiedTime", _connectLastModifiedTime); err != nil {
			log.Errorf("invalid --last-modified-time: %s", err.Error())
			return
		}
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the traffic distribution for a given traffic distribution group.
// When you shift telephony traffic, also shift agents and/or agent sign-ins to
// ensure they can handle the calls in the other Region. If you don't shift the
// agents, voice calls will go to the shifted Region but there won't be any agents
// available to receive the calls.
//
// The SignInConfig distribution is available only on a default
// TrafficDistributionGroup (see the IsDefault parameter in the [TrafficDistributionGroup] data type). If
// you call UpdateTrafficDistribution with a modified SignInConfig and a
// non-default TrafficDistributionGroup , an InvalidRequestException is returned.
//
// For more information about updating a traffic distribution group, see [Update telephony traffic distribution across Amazon Web Services Regions] in the
// Amazon Connect Administrator Guide.
//
// [TrafficDistributionGroup]: https://docs.aws.amazon.com/connect/latest/APIReference/API_TrafficDistributionGroup.html
// [Update telephony traffic distribution across Amazon Web Services Regions]: https://docs.aws.amazon.com/connect/latest/adminguide/update-telephony-traffic-distribution.html
func connect_UpdateTrafficDistribution(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateTrafficDistributionInput{
		// Id: *string, // Required
	}

	if len(_connectId) > 0 {
		input.Id = aws.String(_connectId)
	}
	if len(_connectAgentConfig) > 0 {
		if err := assignInputField(input, "AgentConfig", _connectAgentConfig); err != nil {
			log.Errorf("invalid --agent-config: %s", err.Error())
			return
		}
	}
	if len(_connectSignInConfig) > 0 {
		if err := assignInputField(input, "SignInConfig", _connectSignInConfig); err != nil {
			log.Errorf("invalid --sign-in-config: %s", err.Error())
			return
		}
	}
	if len(_connectTelephonyConfig) > 0 {
		if err := assignInputField(input, "TelephonyConfig", _connectTelephonyConfig); err != nil {
			log.Errorf("invalid --telephony-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrafficDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration settings for the specified user, including
// per-channel auto-accept and after contact work (ACW) timeout settings.
//
// This operation replaces the UpdateUserPhoneConfig API. While
// UpdateUserPhoneConfig applies the same ACW timeout to all channels,
// UpdateUserConfig allows you to set different auto-accept and ACW timeout values
// for each channel type.
func connect_UpdateUserConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserConfigInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectAfterContactWorkConfigs) > 0 {
		if err := assignInputField(input, "AfterContactWorkConfigs", _connectAfterContactWorkConfigs); err != nil {
			log.Errorf("invalid --after-contact-work-configs: %s", err.Error())
			return
		}
	}
	if len(_connectAutoAcceptConfigs) > 0 {
		if err := assignInputField(input, "AutoAcceptConfigs", _connectAutoAcceptConfigs); err != nil {
			log.Errorf("invalid --auto-accept-configs: %s", err.Error())
			return
		}
	}
	if len(_connectPersistentConnectionConfigs) > 0 {
		if err := assignInputField(input, "PersistentConnectionConfigs", _connectPersistentConnectionConfigs); err != nil {
			log.Errorf("invalid --persistent-connection-configs: %s", err.Error())
			return
		}
	}
	if len(_connectPhoneNumberConfigs) > 0 {
		if err := assignInputField(input, "PhoneNumberConfigs", _connectPhoneNumberConfigs); err != nil {
			log.Errorf("invalid --phone-number-configs: %s", err.Error())
			return
		}
	}
	if len(_connectVoiceEnhancementConfigs) > 0 {
		if err := assignInputField(input, "VoiceEnhancementConfigs", _connectVoiceEnhancementConfigs); err != nil {
			log.Errorf("invalid --voice-enhancement-configs: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the specified hierarchy group to the specified user.
func connect_UpdateUserHierarchy(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserHierarchyInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectHierarchyGroupId) > 0 {
		input.HierarchyGroupId = aws.String(_connectHierarchyGroupId)
	}

	if resp, err := client.UpdateUserHierarchy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of the user hierarchy group.
func connect_UpdateUserHierarchyGroupName(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserHierarchyGroupNameInput{
		// HierarchyGroupId: *string, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectHierarchyGroupId) > 0 {
		input.HierarchyGroupId = aws.String(_connectHierarchyGroupId)
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateUserHierarchyGroupName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the user hierarchy structure: add, remove, and rename user hierarchy
// levels.
func connect_UpdateUserHierarchyStructure(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserHierarchyStructureInput{
		// HierarchyStructure: *types.HierarchyStructureUpdate, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectHierarchyStructure) > 0 {
		if err := assignInputField(input, "HierarchyStructure", _connectHierarchyStructure); err != nil {
			log.Errorf("invalid --hierarchy-structure: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}

	if resp, err := client.UpdateUserHierarchyStructure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the identity information for the specified user.
// We strongly recommend limiting who has the ability to invoke
// UpdateUserIdentityInfo . Someone with that ability can change the login
// credentials of other users by changing their email address. This poses a
// security risk to your organization. They can change the email address of a user
// to the attacker's email address, and then reset the password through email. For
// more information, see [Best Practices for Security Profiles]in the Amazon Connect Administrator Guide.
//
// [Best Practices for Security Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-best-practices.html
func connect_UpdateUserIdentityInfo(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserIdentityInfoInput{
		// IdentityInfo: *types.UserIdentityInfo, // Required
		// InstanceId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectIdentityInfo) > 0 {
		if err := assignInputField(input, "IdentityInfo", _connectIdentityInfo); err != nil {
			log.Errorf("invalid --identity-info: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.UpdateUserIdentityInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a notification for a specific user, such as marking it as
// read or hidden. Users can only update notification status for notifications that
// have been sent to them. READ status deprioritizes the notification and greys it
// out, while HIDDEN status removes it from the notification widget.
func connect_UpdateUserNotificationStatus(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserNotificationStatusInput{
		// InstanceId: *string, // Required
		// NotificationId: *string, // Required
		// Status: types.NotificationStatus, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectNotificationId) > 0 {
		input.NotificationId = aws.String(_connectNotificationId)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectLastModifiedRegion) > 0 {
		input.LastModifiedRegion = aws.String(_connectLastModifiedRegion)
	}
	if len(_connectLastModifiedTime) > 0 {
		if err := assignInputField(input, "LastModifiedTime", _connectLastModifiedTime); err != nil {
			log.Errorf("invalid --last-modified-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserNotificationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the phone configuration settings for the specified user.
// We recommend using the [UpdateUserConfig] API, which supports additional functionality that is
// not available in the UpdateUserPhoneConfig API, such as voice enhancement
// settings and per-channel configuration for auto-accept and After Contact Work
// (ACW) timeouts. In comparison, the UpdateUserPhoneConfig API will always set the
// same ACW timeouts to all channels the user handles.
//
// [UpdateUserConfig]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdateUserConfig.html
func connect_UpdateUserPhoneConfig(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserPhoneConfigInput{
		// InstanceId: *string, // Required
		// PhoneConfig: *types.UserPhoneConfig, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPhoneConfig) > 0 {
		if err := assignInputField(input, "PhoneConfig", _connectPhoneConfig); err != nil {
			log.Errorf("invalid --phone-config: %s", err.Error())
			return
		}
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.UpdateUserPhoneConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties associated with the proficiencies of a user.
func connect_UpdateUserProficiencies(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserProficienciesInput{
		// InstanceId: *string, // Required
		// UserId: *string, // Required
		// UserProficiencies: []types.UserProficiency, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}
	if len(_connectUserProficiencies) > 0 {
		if err := assignInputField(input, "UserProficiencies", _connectUserProficiencies); err != nil {
			log.Errorf("invalid --user-proficiencies: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserProficiencies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the specified routing profile to the specified user.
func connect_UpdateUserRoutingProfile(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserRoutingProfileInput{
		// InstanceId: *string, // Required
		// RoutingProfileId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectRoutingProfileId) > 0 {
		input.RoutingProfileId = aws.String(_connectRoutingProfileId)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.UpdateUserRoutingProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns the specified security profiles to the specified user.
func connect_UpdateUserSecurityProfiles(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateUserSecurityProfilesInput{
		// InstanceId: *string, // Required
		// SecurityProfileIds: []string, // Required
		// UserId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectSecurityProfileIds) > 0 {
		input.SecurityProfileIds = append([]string(nil), _connectSecurityProfileIds...)
	}
	if len(_connectUserId) > 0 {
		input.UserId = aws.String(_connectUserId)
	}

	if resp, err := client.UpdateUserSecurityProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the view content of the given view identifier in the specified Amazon
// Connect instance.
//
// It performs content validation if Status is set to SAVED and performs full
// content validation if Status is PUBLISHED . Note that the $SAVED alias' content
// will always be updated, but the $LATEST alias' content will only be updated if
// Status is PUBLISHED .
func connect_UpdateViewContent(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateViewContentInput{
		// Content: *types.ViewInputContent, // Required
		// InstanceId: *string, // Required
		// Status: types.ViewStatus, // Required
		// ViewId: *string, // Required
	}

	if len(_connectContent) > 0 {
		if err := assignInputField(input, "Content", _connectContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectStatus) > 0 {
		if err := assignInputField(input, "Status", _connectStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}

	if resp, err := client.UpdateViewContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the view metadata. Note that either Name or Description must be
// provided.
func connect_UpdateViewMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateViewMetadataInput{
		// InstanceId: *string, // Required
		// ViewId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectViewId) > 0 {
		input.ViewId = aws.String(_connectViewId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}

	if resp, err := client.UpdateViewMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the metadata of a workspace, such as its name and description.
func connect_UpdateWorkspaceMetadata(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateWorkspaceMetadataInput{
		// InstanceId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}
	if len(_connectDescription) > 0 {
		input.Description = aws.String(_connectDescription)
	}
	if len(_connectName) > 0 {
		input.Name = aws.String(_connectName)
	}
	if len(_connectTitle) > 0 {
		input.Title = aws.String(_connectTitle)
	}

	if resp, err := client.UpdateWorkspaceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a page in a workspace, including the associated
// view and input data.
func connect_UpdateWorkspacePage(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateWorkspacePageInput{
		// InstanceId: *string, // Required
		// Page: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectPage) > 0 {
		input.Page = aws.String(_connectPage)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}
	if len(_connectInputData) > 0 {
		input.InputData = aws.String(_connectInputData)
	}
	if len(_connectNewPage) > 0 {
		input.NewPage = aws.String(_connectNewPage)
	}
	if len(_connectResourceArn) > 0 {
		input.ResourceArn = aws.String(_connectResourceArn)
	}
	if len(_connectSlug) > 0 {
		input.Slug = aws.String(_connectSlug)
	}

	if resp, err := client.UpdateWorkspacePage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the theme configuration for a workspace, including colors and styling.
func connect_UpdateWorkspaceTheme(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateWorkspaceThemeInput{
		// InstanceId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}
	if len(_connectTheme) > 0 {
		if err := assignInputField(input, "Theme", _connectTheme); err != nil {
			log.Errorf("invalid --theme: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkspaceTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the visibility setting of a workspace, controlling whether it is
// available to all users, assigned users only, or none.
func connect_UpdateWorkspaceVisibility(cfg aws.Config, client *connect.Client) {
	input := &connect.UpdateWorkspaceVisibilityInput{
		// InstanceId: *string, // Required
		// Visibility: types.Visibility, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_connectInstanceId) > 0 {
		input.InstanceId = aws.String(_connectInstanceId)
	}
	if len(_connectVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _connectVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}
	if len(_connectWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_connectWorkspaceId)
	}

	if resp, err := client.UpdateWorkspaceVisibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_connectCmd)
	_connectCmd.Flags().SortFlags = false

	_connectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_connectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_connectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_connectCmd.Flags().StringVarP(&_connectActions, "actions", "", "", "Actions")
	_connectCmd.Flags().StringVarP(&_connectAdditionalRecipients, "additional-recipients", "", "", "Additional Recipients")
	_connectCmd.Flags().StringVarP(&_connectAfterContactWorkConfigs, "after-contact-work-configs", "", "", "After Contact Work Configs")
	_connectCmd.Flags().StringVarP(&_connectAgentAvailabilityTimer, "agent-availability-timer", "", "", "Agent Availability Timer")
	_connectCmd.Flags().StringVarP(&_connectAgentConfig, "agent-config", "", "", "Agent Config")
	_connectCmd.Flags().StringVarP(&_connectAgentStatusId, "agent-status-id", "", "", "Agent Status ID")
	_connectCmd.Flags().StringVarP(&_connectAgentStatusTypes, "agent-status-types", "", "", "Agent Status Types")
	_connectCmd.Flags().StringVarP(&_connectAliasConfiguration, "alias-configuration", "", "", "Alias Configuration")
	_connectCmd.Flags().StringVarP(&_connectAliasId, "alias-id", "", "", "Alias ID")
	_connectCmd.Flags().StringVarP(&_connectAliasName, "alias-name", "", "", "Alias Name")
	_connectCmd.Flags().StringVarP(&_connectAllowedAccessControlHierarchyGroupId, "allowed-access-control-hierarchy-group-id", "", "", "Allowed Access Control Hierarchy Group ID")
	_connectCmd.Flags().StringVarP(&_connectAllowedAccessControlTags, "allowed-access-control-tags", "", "", "Allowed Access Control Tags")
	_connectCmd.Flags().StringVarP(&_connectAllowedCapabilities, "allowed-capabilities", "", "", "Allowed Capabilities")
	_connectCmd.Flags().StringVarP(&_connectAllowedFlowModules, "allowed-flow-modules", "", "", "Allowed Flow Modules")
	_connectCmd.Flags().StringSliceVarP(&_connectAllowedIps, "allowed-ips", "", nil, "Allowed Ips")
	_connectCmd.Flags().StringVarP(&_connectAllowedMonitorCapabilities, "allowed-monitor-capabilities", "", "", "Allowed Monitor Capabilities")
	_connectCmd.Flags().StringVarP(&_connectAnswerMachineDetectionConfig, "answer-machine-detection-config", "", "", "Answer Machine Detection Config")
	_connectCmd.Flags().StringVarP(&_connectAnswers, "answers", "", "", "Answers")
	_connectCmd.Flags().StringVarP(&_connectApplications, "applications", "", "", "Applications")
	_connectCmd.Flags().StringVarP(&_connectAsDraft, "as-draft", "", "", "As Draft")
	_connectCmd.Flags().StringVarP(&_connectAssociatedResourceArn, "associated-resource-arn", "", "", "Associated Resource ARN")
	_connectCmd.Flags().StringVarP(&_connectAssociationId, "association-id", "", "", "Association ID")
	_connectCmd.Flags().StringVarP(&_connectAttachments, "attachments", "", "", "Attachments")
	_connectCmd.Flags().StringVarP(&_connectAttributeConfiguration, "attribute-configuration", "", "", "Attribute Configuration")
	_connectCmd.Flags().StringSliceVarP(&_connectAttributeIds, "attribute-ids", "", nil, "Attribute Ids")
	_connectCmd.Flags().StringVarP(&_connectAttributeName, "attribute-name", "", "", "Attribute Name")
	_connectCmd.Flags().StringVarP(&_connectAttributeType, "attribute-type", "", "", "Attribute Type")
	_connectCmd.Flags().StringVarP(&_connectAttributes, "attributes", "", "", "Attributes")
	_connectCmd.Flags().StringVarP(&_connectAuthenticationProfileId, "authentication-profile-id", "", "", "Authentication Profile ID")
	_connectCmd.Flags().StringVarP(&_connectAutoAcceptConfigs, "auto-accept-configs", "", "", "Auto Accept Configs")
	_connectCmd.Flags().StringVarP(&_connectAutoEvaluationConfiguration, "auto-evaluation-configuration", "", "", "Auto Evaluation Configuration")
	_connectCmd.Flags().StringSliceVarP(&_connectBlockedIps, "blocked-ips", "", nil, "Blocked Ips")
	_connectCmd.Flags().StringVarP(&_connectBotName, "bot-name", "", "", "Bot Name")
	_connectCmd.Flags().StringVarP(&_connectCampaignId, "campaign-id", "", "", "Campaign ID")
	_connectCmd.Flags().StringVarP(&_connectChannel, "channel", "", "", "Channel")
	_connectCmd.Flags().StringVarP(&_connectChannelConfiguration, "channel-configuration", "", "", "Channel Configuration")
	_connectCmd.Flags().StringVarP(&_connectChatDurationInMinutes, "chat-duration-in-minutes", "", "", "Chat Duration In Minutes")
	_connectCmd.Flags().StringVarP(&_connectChatStreamingConfiguration, "chat-streaming-configuration", "", "", "Chat Streaming Configuration")
	_connectCmd.Flags().StringVarP(&_connectClientToken, "client-token", "", "", "Client Token")
	_connectCmd.Flags().StringVarP(&_connectCode, "code", "", "", "Code")
	_connectCmd.Flags().StringVarP(&_connectConfig, "config", "", "", "Config")
	_connectCmd.Flags().StringVarP(&_connectConstraints, "constraints", "", "", "Constraints")
	_connectCmd.Flags().StringVarP(&_connectContactConfiguration, "contact-configuration", "", "", "Contact Configuration")
	_connectCmd.Flags().StringVarP(&_connectContactDataRequestList, "contact-data-request-list", "", "", "Contact Data Request List")
	_connectCmd.Flags().StringVarP(&_connectContactFlowId, "contact-flow-id", "", "", "Contact Flow ID")
	_connectCmd.Flags().StringVarP(&_connectContactFlowModuleId, "contact-flow-module-id", "", "", "Contact Flow Module ID")
	_connectCmd.Flags().StringVarP(&_connectContactFlowModuleState, "contact-flow-module-state", "", "", "Contact Flow Module State")
	_connectCmd.Flags().StringVarP(&_connectContactFlowModuleVersion, "contact-flow-module-version", "", "", "Contact Flow Module Version")
	_connectCmd.Flags().StringVarP(&_connectContactFlowState, "contact-flow-state", "", "", "Contact Flow State")
	_connectCmd.Flags().StringVarP(&_connectContactFlowTypes, "contact-flow-types", "", "", "Contact Flow Types")
	_connectCmd.Flags().StringVarP(&_connectContactFlowVersion, "contact-flow-version", "", "", "Contact Flow Version")
	_connectCmd.Flags().StringVarP(&_connectContactId, "contact-id", "", "", "Contact ID")
	_connectCmd.Flags().StringVarP(&_connectContactRecordingType, "contact-recording-type", "", "", "Contact Recording Type")
	_connectCmd.Flags().StringVarP(&_connectContent, "content", "", "", "Content")
	_connectCmd.Flags().StringVarP(&_connectCreateNewVersion, "create-new-version", "", "", "Create New Version")
	_connectCmd.Flags().StringVarP(&_connectCreatedBy, "created-by", "", "", "Created By")
	_connectCmd.Flags().StringVarP(&_connectCurrentMetrics, "current-metrics", "", "", "Current Metrics")
	_connectCmd.Flags().StringVarP(&_connectCustomerEndpoint, "customer-endpoint", "", "", "Customer Endpoint")
	_connectCmd.Flags().StringVarP(&_connectCustomerId, "customer-id", "", "", "Customer ID")
	_connectCmd.Flags().StringVarP(&_connectDataSetId, "data-set-id", "", "", "Data Set ID")
	_connectCmd.Flags().StringSliceVarP(&_connectDataSetIds, "data-set-ids", "", nil, "Data Set Ids")
	_connectCmd.Flags().StringVarP(&_connectDataTableId, "data-table-id", "", "", "Data Table ID")
	_connectCmd.Flags().StringVarP(&_connectDefaultOutboundQueueId, "default-outbound-queue-id", "", "", "Default Outbound Queue ID")
	_connectCmd.Flags().StringVarP(&_connectDefaults, "defaults", "", "", "Defaults")
	_connectCmd.Flags().StringVarP(&_connectDescription, "description", "", "", "Description")
	_connectCmd.Flags().StringVarP(&_connectDestinationEmailAddress, "destination-email-address", "", "", "Destination Email Address")
	_connectCmd.Flags().StringVarP(&_connectDestinationEndpoint, "destination-endpoint", "", "", "Destination Endpoint")
	_connectCmd.Flags().StringVarP(&_connectDestinationId, "destination-id", "", "", "Destination ID")
	_connectCmd.Flags().StringVarP(&_connectDestinationPhoneNumber, "destination-phone-number", "", "", "Destination Phone Number")
	_connectCmd.Flags().StringVarP(&_connectDeviceToken, "device-token", "", "", "Device Token")
	_connectCmd.Flags().StringVarP(&_connectDeviceType, "device-type", "", "", "Device Type")
	_connectCmd.Flags().StringVarP(&_connectDirectoryId, "directory-id", "", "", "Directory ID")
	_connectCmd.Flags().StringVarP(&_connectDirectoryUserId, "directory-user-id", "", "", "Directory User ID")
	_connectCmd.Flags().StringVarP(&_connectDisconnectOnCustomerExit, "disconnect-on-customer-exit", "", "", "Disconnect On Customer Exit")
	_connectCmd.Flags().StringVarP(&_connectDisconnectReason, "disconnect-reason", "", "", "Disconnect Reason")
	_connectCmd.Flags().StringVarP(&_connectDisplayName, "display-name", "", "", "Display Name")
	_connectCmd.Flags().StringVarP(&_connectDisplayOrder, "display-order", "", "", "Display Order")
	_connectCmd.Flags().StringVarP(&_connectEffectiveFrom, "effective-from", "", "", "Effective From")
	_connectCmd.Flags().StringVarP(&_connectEffectiveTill, "effective-till", "", "", "Effective Till")
	_connectCmd.Flags().StringVarP(&_connectEmailAddress, "email-address", "", "", "Email Address")
	_connectCmd.Flags().StringVarP(&_connectEmailAddressId, "email-address-id", "", "", "Email Address ID")
	_connectCmd.Flags().StringVarP(&_connectEmailMessage, "email-message", "", "", "Email Message")
	_connectCmd.Flags().StringVarP(&_connectEndTime, "end-time", "", "", "End Time")
	_connectCmd.Flags().StringVarP(&_connectEntityArn, "entity-arn", "", "", "Entity ARN")
	_connectCmd.Flags().StringVarP(&_connectEntityType, "entity-type", "", "", "Entity Type")
	_connectCmd.Flags().StringVarP(&_connectEntryPoint, "entry-point", "", "", "Entry Point")
	_connectCmd.Flags().StringVarP(&_connectError, "error", "", "", "Error")
	_connectCmd.Flags().StringVarP(&_connectErrorDescription, "error-description", "", "", "Error Description")
	_connectCmd.Flags().StringVarP(&_connectEvaluationFormId, "evaluation-form-id", "", "", "Evaluation Form ID")
	_connectCmd.Flags().StringVarP(&_connectEvaluationFormVersion, "evaluation-form-version", "", "", "Evaluation Form Version")
	_connectCmd.Flags().StringVarP(&_connectEvaluationId, "evaluation-id", "", "", "Evaluation ID")
	_connectCmd.Flags().StringVarP(&_connectEvent, "event", "", "", "Event")
	_connectCmd.Flags().StringVarP(&_connectEventSourceName, "event-source-name", "", "", "Event Source Name")
	_connectCmd.Flags().StringVarP(&_connectExpiresAt, "expires-at", "", "", "Expires At")
	_connectCmd.Flags().StringVarP(&_connectExpiryDurationInMinutes, "expiry-duration-in-minutes", "", "", "Expiry Duration In Minutes")
	_connectCmd.Flags().StringVarP(&_connectExternalInvocationConfiguration, "external-invocation-configuration", "", "", "External Invocation Configuration")
	_connectCmd.Flags().StringVarP(&_connectFailureMode, "failure-mode", "", "", "Failure Mode")
	_connectCmd.Flags().StringVarP(&_connectFields, "fields", "", "", "Fields")
	_connectCmd.Flags().StringVarP(&_connectFileId, "file-id", "", "", "File ID")
	_connectCmd.Flags().StringSliceVarP(&_connectFileIds, "file-ids", "", nil, "File Ids")
	_connectCmd.Flags().StringVarP(&_connectFileName, "file-name", "", "", "File Name")
	_connectCmd.Flags().StringVarP(&_connectFileSizeInBytes, "file-size-in-bytes", "", "", "File Size In Bytes")
	_connectCmd.Flags().StringVarP(&_connectFileUseCaseType, "file-use-case-type", "", "", "File Use Case Type")
	_connectCmd.Flags().StringVarP(&_connectFilters, "filters", "", "", "Filters")
	_connectCmd.Flags().StringVarP(&_connectFlowContentSha256, "flow-content-sha256", "", "", "Flow Content SHA256")
	_connectCmd.Flags().StringVarP(&_connectFlowId, "flow-id", "", "", "Flow ID")
	_connectCmd.Flags().StringVarP(&_connectFlowModuleContentSha256, "flow-module-content-sha256", "", "", "Flow Module Content SHA256")
	_connectCmd.Flags().StringVarP(&_connectFromDate, "from-date", "", "", "From Date")
	_connectCmd.Flags().StringVarP(&_connectFromEmailAddress, "from-email-address", "", "", "From Email Address")
	_connectCmd.Flags().StringVarP(&_connectFunction, "function", "", "", "Function")
	_connectCmd.Flags().StringVarP(&_connectFunctionArn, "function-arn", "", "", "Function ARN")
	_connectCmd.Flags().StringVarP(&_connectGranularAccessControlConfiguration, "granular-access-control-configuration", "", "", "Granular Access Control Configuration")
	_connectCmd.Flags().StringSliceVarP(&_connectGroupings, "groupings", "", nil, "Groupings")
	_connectCmd.Flags().StringVarP(&_connectHierarchyGroupId, "hierarchy-group-id", "", "", "Hierarchy Group ID")
	_connectCmd.Flags().StringSliceVarP(&_connectHierarchyRestrictedResources, "hierarchy-restricted-resources", "", nil, "Hierarchy Restricted Resources")
	_connectCmd.Flags().StringVarP(&_connectHierarchyStructure, "hierarchy-structure", "", "", "Hierarchy Structure")
	_connectCmd.Flags().StringVarP(&_connectHistoricalMetrics, "historical-metrics", "", "", "Historical Metrics")
	_connectCmd.Flags().StringVarP(&_connectHoursOfOperationId, "hours-of-operation-id", "", "", "Hours Of Operation ID")
	_connectCmd.Flags().StringVarP(&_connectHoursOfOperationOverrideId, "hours-of-operation-override-id", "", "", "Hours Of Operation Override ID")
	_connectCmd.Flags().StringVarP(&_connectId, "id", "", "", "ID")
	_connectCmd.Flags().StringVarP(&_connectIdentityInfo, "identity-info", "", "", "Identity Info")
	_connectCmd.Flags().StringVarP(&_connectIdentityManagementType, "identity-management-type", "", "", "Identity Management Type")
	_connectCmd.Flags().StringVarP(&_connectInboundCallsEnabled, "inbound-calls-enabled", "", "", "Inbound Calls Enabled")
	_connectCmd.Flags().StringVarP(&_connectInitialContactId, "initial-contact-id", "", "", "Initial Contact ID")
	_connectCmd.Flags().StringVarP(&_connectInitialMessage, "initial-message", "", "", "Initial Message")
	_connectCmd.Flags().StringVarP(&_connectInitialSystemMessage, "initial-system-message", "", "", "Initial System Message")
	_connectCmd.Flags().StringVarP(&_connectInitialTemplatedSystemMessage, "initial-templated-system-message", "", "", "Initial Templated System Message")
	_connectCmd.Flags().StringVarP(&_connectInitializationData, "initialization-data", "", "", "Initialization Data")
	_connectCmd.Flags().StringVarP(&_connectInitiateAs, "initiate-as", "", "", "Initiate As")
	_connectCmd.Flags().StringVarP(&_connectInitiationMethod, "initiation-method", "", "", "Initiation Method")
	_connectCmd.Flags().StringVarP(&_connectInputData, "input-data", "", "", "Input Data")
	_connectCmd.Flags().StringVarP(&_connectInstanceAlias, "instance-alias", "", "", "Instance Alias")
	_connectCmd.Flags().StringVarP(&_connectInstanceId, "instance-id", "", "", "Instance ID")
	_connectCmd.Flags().StringVarP(&_connectIntegrationArn, "integration-arn", "", "", "Integration ARN")
	_connectCmd.Flags().StringVarP(&_connectIntegrationAssociationId, "integration-association-id", "", "", "Integration Association ID")
	_connectCmd.Flags().StringVarP(&_connectIntegrationType, "integration-type", "", "", "Integration Type")
	_connectCmd.Flags().StringVarP(&_connectInterval, "interval", "", "", "Interval")
	_connectCmd.Flags().StringVarP(&_connectItems, "items", "", "", "Items")
	_connectCmd.Flags().StringVarP(&_connectKey, "key", "", "", "Key")
	_connectCmd.Flags().StringVarP(&_connectLanguageCode, "language-code", "", "", "Language Code")
	_connectCmd.Flags().StringVarP(&_connectLanguageConfiguration, "language-configuration", "", "", "Language Configuration")
	_connectCmd.Flags().StringVarP(&_connectLastModifiedRegion, "last-modified-region", "", "", "Last Modified Region")
	_connectCmd.Flags().StringVarP(&_connectLastModifiedTime, "last-modified-time", "", "", "Last Modified Time")
	_connectCmd.Flags().StringVarP(&_connectLexBot, "lex-bot", "", "", "Lex Bot")
	_connectCmd.Flags().StringVarP(&_connectLexRegion, "lex-region", "", "", "Lex Region")
	_connectCmd.Flags().StringVarP(&_connectLexV2Bot, "lex-v2-bot", "", "", "Lex V2 Bot")
	_connectCmd.Flags().StringVarP(&_connectLexVersion, "lex-version", "", "", "Lex Version")
	_connectCmd.Flags().StringVarP(&_connectLockVersion, "lock-version", "", "", "Lock Version")
	_connectCmd.Flags().StringVarP(&_connectManualAssignmentQueueConfigs, "manual-assignment-queue-configs", "", "", "Manual Assignment Queue Configs")
	_connectCmd.Flags().StringVarP(&_connectManualAssignmentQueueReferences, "manual-assignment-queue-references", "", "", "Manual Assignment Queue References")
	_connectCmd.Flags().StringVarP(&_connectMaxContacts, "max-contacts", "", "", "Max Contacts")
	_connectCmd.Flags().StringVarP(&_connectMaxResults, "max-results", "", "", "Max Results")
	_connectCmd.Flags().StringVarP(&_connectMediaConcurrencies, "media-concurrencies", "", "", "Media Concurrencies")
	_connectCmd.Flags().StringVarP(&_connectMediaSource, "media-source", "", "", "Media Source")
	_connectCmd.Flags().StringVarP(&_connectMediaType, "media-type", "", "", "Media Type")
	_connectCmd.Flags().StringVarP(&_connectMetrics, "metrics", "", "", "Metrics")
	_connectCmd.Flags().StringVarP(&_connectName, "name", "", "", "Name")
	_connectCmd.Flags().StringVarP(&_connectNameStartsWith, "name-starts-with", "", "", "Name Starts With")
	_connectCmd.Flags().StringVarP(&_connectNewPage, "new-page", "", "", "New Page")
	_connectCmd.Flags().StringVarP(&_connectNewPrimaryValues, "new-primary-values", "", "", "New Primary Values")
	_connectCmd.Flags().StringVarP(&_connectNewSessionDetails, "new-session-details", "", "", "New Session Details")
	_connectCmd.Flags().StringVarP(&_connectNextToken, "next-token", "", "", "Next Token")
	_connectCmd.Flags().StringVarP(&_connectNotes, "notes", "", "", "Notes")
	_connectCmd.Flags().StringVarP(&_connectNotificationId, "notification-id", "", "", "Notification ID")
	_connectCmd.Flags().StringVarP(&_connectOrigin, "origin", "", "", "Origin")
	_connectCmd.Flags().StringVarP(&_connectOutboundCallerConfig, "outbound-caller-config", "", "", "Outbound Caller Config")
	_connectCmd.Flags().StringVarP(&_connectOutboundCallsEnabled, "outbound-calls-enabled", "", "", "Outbound Calls Enabled")
	_connectCmd.Flags().StringVarP(&_connectOutboundEmailConfig, "outbound-email-config", "", "", "Outbound Email Config")
	_connectCmd.Flags().StringVarP(&_connectOutboundStrategy, "outbound-strategy", "", "", "Outbound Strategy")
	_connectCmd.Flags().StringVarP(&_connectOutputType, "output-type", "", "", "Output Type")
	_connectCmd.Flags().StringVarP(&_connectOverrideType, "override-type", "", "", "Override Type")
	_connectCmd.Flags().StringVarP(&_connectPage, "page", "", "", "Page")
	_connectCmd.Flags().StringVarP(&_connectParentGroupId, "parent-group-id", "", "", "Parent Group ID")
	_connectCmd.Flags().StringVarP(&_connectParentHoursOfOperationConfigs, "parent-hours-of-operation-configs", "", "", "Parent Hours Of Operation Configs")
	_connectCmd.Flags().StringSliceVarP(&_connectParentHoursOfOperationIds, "parent-hours-of-operation-ids", "", nil, "Parent Hours Of Operation Ids")
	_connectCmd.Flags().StringVarP(&_connectParticipantConfiguration, "participant-configuration", "", "", "Participant Configuration")
	_connectCmd.Flags().StringVarP(&_connectParticipantDetails, "participant-details", "", "", "Participant Details")
	_connectCmd.Flags().StringVarP(&_connectPassword, "password", "", "", "Password")
	_connectCmd.Flags().StringVarP(&_connectPeriodicSessionDuration, "periodic-session-duration", "", "", "Periodic Session Duration")
	_connectCmd.Flags().StringSliceVarP(&_connectPermissions, "permissions", "", nil, "Permissions")
	_connectCmd.Flags().StringVarP(&_connectPersistentChat, "persistent-chat", "", "", "Persistent Chat")
	_connectCmd.Flags().StringVarP(&_connectPersistentConnectionConfigs, "persistent-connection-configs", "", "", "Persistent Connection Configs")
	_connectCmd.Flags().StringVarP(&_connectPhoneConfig, "phone-config", "", "", "Phone Config")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumber, "phone-number", "", "", "Phone Number")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberConfigs, "phone-number-configs", "", "", "Phone Number Configs")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberCountryCode, "phone-number-country-code", "", "", "Phone Number Country Code")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberCountryCodes, "phone-number-country-codes", "", "", "Phone Number Country Codes")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberDescription, "phone-number-description", "", "", "Phone Number Description")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberId, "phone-number-id", "", "", "Phone Number ID")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberPrefix, "phone-number-prefix", "", "", "Phone Number Prefix")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberType, "phone-number-type", "", "", "Phone Number Type")
	_connectCmd.Flags().StringVarP(&_connectPhoneNumberTypes, "phone-number-types", "", "", "Phone Number Types")
	_connectCmd.Flags().StringVarP(&_connectPinpointAppArn, "pinpoint-app-arn", "", "", "Pinpoint App ARN")
	_connectCmd.Flags().StringVarP(&_connectPredefinedNotificationId, "predefined-notification-id", "", "", "Predefined Notification ID")
	_connectCmd.Flags().StringVarP(&_connectPreviousContactId, "previous-contact-id", "", "", "Previous Contact ID")
	_connectCmd.Flags().StringVarP(&_connectPrimary, "primary", "", "", "Primary")
	_connectCmd.Flags().StringVarP(&_connectPrimaryAttributeValues, "primary-attribute-values", "", "", "Primary Attribute Values")
	_connectCmd.Flags().StringVarP(&_connectPrimaryValues, "primary-values", "", "", "Primary Values")
	_connectCmd.Flags().StringVarP(&_connectPriority, "priority", "", "", "Priority")
	_connectCmd.Flags().StringVarP(&_connectProcessorArn, "processor-arn", "", "", "Processor ARN")
	_connectCmd.Flags().StringVarP(&_connectPromptId, "prompt-id", "", "", "Prompt ID")
	_connectCmd.Flags().StringVarP(&_connectPublishStatus, "publish-status", "", "", "Publish Status")
	_connectCmd.Flags().StringSliceVarP(&_connectPurposes, "purposes", "", nil, "Purposes")
	_connectCmd.Flags().StringVarP(&_connectQueueConfigs, "queue-configs", "", "", "Queue Configs")
	_connectCmd.Flags().StringVarP(&_connectQueueId, "queue-id", "", "", "Queue ID")
	_connectCmd.Flags().StringVarP(&_connectQueueInfo, "queue-info", "", "", "Queue Info")
	_connectCmd.Flags().StringVarP(&_connectQueuePriority, "queue-priority", "", "", "Queue Priority")
	_connectCmd.Flags().StringVarP(&_connectQueueReferences, "queue-references", "", "", "Queue References")
	_connectCmd.Flags().StringVarP(&_connectQueueTimeAdjustmentSeconds, "queue-time-adjustment-seconds", "", "", "Queue Time Adjustment Seconds")
	_connectCmd.Flags().StringVarP(&_connectQueueTypes, "queue-types", "", "", "Queue Types")
	_connectCmd.Flags().StringVarP(&_connectQuickConnectConfig, "quick-connect-config", "", "", "Quick Connect Config")
	_connectCmd.Flags().StringVarP(&_connectQuickConnectId, "quick-connect-id", "", "", "Quick Connect ID")
	_connectCmd.Flags().StringSliceVarP(&_connectQuickConnectIds, "quick-connect-ids", "", nil, "Quick Connect Ids")
	_connectCmd.Flags().StringVarP(&_connectQuickConnectTypes, "quick-connect-types", "", "", "Quick Connect Types")
	_connectCmd.Flags().StringSliceVarP(&_connectRecipients, "recipients", "", nil, "Recipients")
	_connectCmd.Flags().StringSliceVarP(&_connectRecordIds, "record-ids", "", nil, "Record Ids")
	_connectCmd.Flags().StringVarP(&_connectRecurrenceConfig, "recurrence-config", "", "", "Recurrence Config")
	_connectCmd.Flags().StringVarP(&_connectReferenceTypes, "reference-types", "", "", "Reference Types")
	_connectCmd.Flags().StringVarP(&_connectReferences, "references", "", "", "References")
	_connectCmd.Flags().StringVarP(&_connectRegistrationId, "registration-id", "", "", "Registration ID")
	_connectCmd.Flags().StringVarP(&_connectRehydrationType, "rehydration-type", "", "", "Rehydration Type")
	_connectCmd.Flags().StringVarP(&_connectRelatedContactId, "related-contact-id", "", "", "Related Contact ID")
	_connectCmd.Flags().StringVarP(&_connectReplicaAlias, "replica-alias", "", "", "Replica Alias")
	_connectCmd.Flags().StringVarP(&_connectReplicaRegion, "replica-region", "", "", "Replica Region")
	_connectCmd.Flags().StringVarP(&_connectResetOrderNumber, "reset-order-number", "", "", "Reset Order Number")
	_connectCmd.Flags().StringVarP(&_connectResourceArn, "resource-arn", "", "", "Resource ARN")
	_connectCmd.Flags().StringSliceVarP(&_connectResourceArns, "resource-arns", "", nil, "Resource Arns")
	_connectCmd.Flags().StringVarP(&_connectResourceId, "resource-id", "", "", "Resource ID")
	_connectCmd.Flags().StringSliceVarP(&_connectResourceIds, "resource-ids", "", nil, "Resource Ids")
	_connectCmd.Flags().StringVarP(&_connectResourceType, "resource-type", "", "", "Resource Type")
	_connectCmd.Flags().StringSliceVarP(&_connectResourceTypes, "resource-types", "", nil, "Resource Types")
	_connectCmd.Flags().StringVarP(&_connectReviewConfiguration, "review-configuration", "", "", "Review Configuration")
	_connectCmd.Flags().StringVarP(&_connectRingTimeoutInSeconds, "ring-timeout-in-seconds", "", "", "Ring Timeout In Seconds")
	_connectCmd.Flags().StringVarP(&_connectRoutingCriteria, "routing-criteria", "", "", "Routing Criteria")
	_connectCmd.Flags().StringVarP(&_connectRoutingProfileId, "routing-profile-id", "", "", "Routing Profile ID")
	_connectCmd.Flags().StringVarP(&_connectRuleId, "rule-id", "", "", "Rule ID")
	_connectCmd.Flags().StringVarP(&_connectS3Uri, "s3-uri", "", "", "S3 URI")
	_connectCmd.Flags().StringVarP(&_connectScheduledTime, "scheduled-time", "", "", "Scheduled Time")
	_connectCmd.Flags().StringVarP(&_connectScoringStrategy, "scoring-strategy", "", "", "Scoring Strategy")
	_connectCmd.Flags().StringVarP(&_connectSearchCriteria, "search-criteria", "", "", "Search Criteria")
	_connectCmd.Flags().StringVarP(&_connectSearchFilter, "search-filter", "", "", "Search Filter")
	_connectCmd.Flags().StringVarP(&_connectSecurityProfileId, "security-profile-id", "", "", "Security Profile ID")
	_connectCmd.Flags().StringSliceVarP(&_connectSecurityProfileIds, "security-profile-ids", "", nil, "Security Profile Ids")
	_connectCmd.Flags().StringVarP(&_connectSecurityProfileName, "security-profile-name", "", "", "Security Profile Name")
	_connectCmd.Flags().StringVarP(&_connectSecurityProfiles, "security-profiles", "", "", "Security Profiles")
	_connectCmd.Flags().StringVarP(&_connectSegmentAttributes, "segment-attributes", "", "", "Segment Attributes")
	_connectCmd.Flags().StringVarP(&_connectSegmentTypes, "segment-types", "", "", "Segment Types")
	_connectCmd.Flags().StringVarP(&_connectSelfAssignFlowId, "self-assign-flow-id", "", "", "Self Assign Flow ID")
	_connectCmd.Flags().StringVarP(&_connectSessionInactivityDuration, "session-inactivity-duration", "", "", "Session Inactivity Duration")
	_connectCmd.Flags().StringVarP(&_connectSessionInactivityHandlingEnabled, "session-inactivity-handling-enabled", "", "", "Session Inactivity Handling Enabled")
	_connectCmd.Flags().StringVarP(&_connectSettings, "settings", "", "", "Settings")
	_connectCmd.Flags().StringVarP(&_connectSignInConfig, "sign-in-config", "", "", "Sign In Config")
	_connectCmd.Flags().StringVarP(&_connectSlug, "slug", "", "", "Slug")
	_connectCmd.Flags().StringVarP(&_connectSnapshotVersion, "snapshot-version", "", "", "Snapshot Version")
	_connectCmd.Flags().StringVarP(&_connectSort, "sort", "", "", "Sort")
	_connectCmd.Flags().StringVarP(&_connectSortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_connectCmd.Flags().StringVarP(&_connectSourceApplicationName, "source-application-name", "", "", "Source Application Name")
	_connectCmd.Flags().StringVarP(&_connectSourceApplicationUrl, "source-application-url", "", "", "Source Application URL")
	_connectCmd.Flags().StringVarP(&_connectSourceCampaign, "source-campaign", "", "", "Source Campaign")
	_connectCmd.Flags().StringVarP(&_connectSourceContactId, "source-contact-id", "", "", "Source Contact ID")
	_connectCmd.Flags().StringVarP(&_connectSourceEndpoint, "source-endpoint", "", "", "Source Endpoint")
	_connectCmd.Flags().StringVarP(&_connectSourceId, "source-id", "", "", "Source ID")
	_connectCmd.Flags().StringVarP(&_connectSourcePhoneNumber, "source-phone-number", "", "", "Source Phone Number")
	_connectCmd.Flags().StringVarP(&_connectSourcePhoneNumberArn, "source-phone-number-arn", "", "", "Source Phone Number ARN")
	_connectCmd.Flags().StringVarP(&_connectSourceType, "source-type", "", "", "Source Type")
	_connectCmd.Flags().StringVarP(&_connectStartTime, "start-time", "", "", "Start Time")
	_connectCmd.Flags().StringVarP(&_connectState, "state", "", "", "State")
	_connectCmd.Flags().StringVarP(&_connectStatus, "status", "", "", "Status")
	_connectCmd.Flags().StringVarP(&_connectStorageConfig, "storage-config", "", "", "Storage Config")
	_connectCmd.Flags().StringVarP(&_connectStreamingId, "streaming-id", "", "", "Streaming ID")
	_connectCmd.Flags().StringVarP(&_connectSubmittedBy, "submitted-by", "", "", "Submitted By")
	_connectCmd.Flags().StringVarP(&_connectSubtype, "subtype", "", "", "Subtype")
	_connectCmd.Flags().StringSliceVarP(&_connectSupportedMessagingContentTypes, "supported-messaging-content-types", "", nil, "Supported Messaging Content Types")
	_connectCmd.Flags().StringVarP(&_connectSystemEndpoint, "system-endpoint", "", "", "System Endpoint")
	_connectCmd.Flags().StringSliceVarP(&_connectTagKeys, "tag-keys", "", nil, "Tag Keys")
	_connectCmd.Flags().StringSliceVarP(&_connectTagRestrictedResources, "tag-restricted-resources", "", nil, "Tag Restricted Resources")
	_connectCmd.Flags().StringVarP(&_connectTags, "tags", "", "", "Tags")
	_connectCmd.Flags().StringVarP(&_connectTargetAccountId, "target-account-id", "", "", "Target Account ID")
	_connectCmd.Flags().StringVarP(&_connectTargetArn, "target-arn", "", "", "Target ARN")
	_connectCmd.Flags().StringVarP(&_connectTargetConfiguration, "target-configuration", "", "", "Target Configuration")
	_connectCmd.Flags().StringVarP(&_connectTaskTemplateId, "task-template-id", "", "", "Task Template ID")
	_connectCmd.Flags().StringVarP(&_connectTelephonyConfig, "telephony-config", "", "", "Telephony Config")
	_connectCmd.Flags().StringVarP(&_connectTestCaseExecutionId, "test-case-execution-id", "", "", "Test Case Execution ID")
	_connectCmd.Flags().StringVarP(&_connectTestCaseId, "test-case-id", "", "", "Test Case ID")
	_connectCmd.Flags().StringVarP(&_connectTestCaseName, "test-case-name", "", "", "Test Case Name")
	_connectCmd.Flags().StringVarP(&_connectTheme, "theme", "", "", "Theme")
	_connectCmd.Flags().StringVarP(&_connectTimeRange, "time-range", "", "", "Time Range")
	_connectCmd.Flags().StringVarP(&_connectTimeZone, "time-zone", "", "", "Time Zone")
	_connectCmd.Flags().StringVarP(&_connectTitle, "title", "", "", "Title")
	_connectCmd.Flags().StringVarP(&_connectToDate, "to-date", "", "", "To Date")
	_connectCmd.Flags().StringVarP(&_connectTrafficDistributionGroupId, "traffic-distribution-group-id", "", "", "Traffic Distribution Group ID")
	_connectCmd.Flags().StringVarP(&_connectTrafficType, "traffic-type", "", "", "Traffic Type")
	_connectCmd.Flags().StringVarP(&_connectTriggerEventSource, "trigger-event-source", "", "", "Trigger Event Source")
	_connectCmd.Flags().StringVarP(&_connectType, "type", "", "", "Type")
	_connectCmd.Flags().StringVarP(&_connectUpdatedBy, "updated-by", "", "", "Updated By")
	_connectCmd.Flags().StringVarP(&_connectUrlExpiryInSeconds, "url-expiry-in-seconds", "", "", "URL Expiry In Seconds")
	_connectCmd.Flags().StringVarP(&_connectUseCaseId, "use-case-id", "", "", "Use Case ID")
	_connectCmd.Flags().StringVarP(&_connectUseCaseType, "use-case-type", "", "", "Use Case Type")
	_connectCmd.Flags().StringVarP(&_connectUserId, "user-id", "", "", "User ID")
	_connectCmd.Flags().StringVarP(&_connectUserInfo, "user-info", "", "", "User Info")
	_connectCmd.Flags().StringVarP(&_connectUserProficiencies, "user-proficiencies", "", "", "User Proficiencies")
	_connectCmd.Flags().StringVarP(&_connectUsername, "username", "", "", "Username")
	_connectCmd.Flags().StringVarP(&_connectValidation, "validation", "", "", "Validation")
	_connectCmd.Flags().StringVarP(&_connectValue, "value", "", "", "Value")
	_connectCmd.Flags().StringVarP(&_connectValueLockLevel, "value-lock-level", "", "", "Value Lock Level")
	_connectCmd.Flags().StringVarP(&_connectValueType, "value-type", "", "", "Value Type")
	_connectCmd.Flags().StringVarP(&_connectValues, "values", "", "", "Values")
	_connectCmd.Flags().StringVarP(&_connectVersionDescription, "version-description", "", "", "Version Description")
	_connectCmd.Flags().StringVarP(&_connectViewContentSha256, "view-content-sha256", "", "", "View Content SHA256")
	_connectCmd.Flags().StringVarP(&_connectViewId, "view-id", "", "", "View ID")
	_connectCmd.Flags().StringVarP(&_connectViewVersion, "view-version", "", "", "View Version")
	_connectCmd.Flags().StringVarP(&_connectVisibility, "visibility", "", "", "Visibility")
	_connectCmd.Flags().StringVarP(&_connectVocabularyId, "vocabulary-id", "", "", "Vocabulary ID")
	_connectCmd.Flags().StringVarP(&_connectVocabularyName, "vocabulary-name", "", "", "Vocabulary Name")
	_connectCmd.Flags().StringVarP(&_connectVoiceEnhancementConfigs, "voice-enhancement-configs", "", "", "Voice Enhancement Configs")
	_connectCmd.Flags().StringVarP(&_connectVoiceRecordingConfiguration, "voice-recording-configuration", "", "", "Voice Recording Configuration")
	_connectCmd.Flags().StringVarP(&_connectWorkspaceId, "workspace-id", "", "", "Workspace ID")

	_connectCmd.Flags().BoolVarP(&_connectActivateEvaluationForm, "activate-evaluation-form", "", false, "Activate Evaluation Form")
	_connectCmd.Flags().BoolVarP(&_connectAssociateAnalyticsDataSet, "associate-analytics-data-set", "", false, "Associate Analytics Data Set")
	_connectCmd.Flags().BoolVarP(&_connectAssociateApprovedOrigin, "associate-approved-origin", "", false, "Associate Approved Origin")
	_connectCmd.Flags().BoolVarP(&_connectAssociateBot, "associate-bot", "", false, "Associate Bot")
	_connectCmd.Flags().BoolVarP(&_connectAssociateContactWithUser, "associate-contact-with-user", "", false, "Associate Contact With User")
	_connectCmd.Flags().BoolVarP(&_connectAssociateDefaultVocabulary, "associate-default-vocabulary", "", false, "Associate Default Vocabulary")
	_connectCmd.Flags().BoolVarP(&_connectAssociateEmailAddressAlias, "associate-email-address-alias", "", false, "Associate Email Address Alias")
	_connectCmd.Flags().BoolVarP(&_connectAssociateFlow, "associate-flow", "", false, "Associate Flow")
	_connectCmd.Flags().BoolVarP(&_connectAssociateHoursOfOperations, "associate-hours-of-operations", "", false, "Associate Hours Of Operations")
	_connectCmd.Flags().BoolVarP(&_connectAssociateInstanceStorageConfig, "associate-instance-storage-config", "", false, "Associate Instance Storage Config")
	_connectCmd.Flags().BoolVarP(&_connectAssociateLambdaFunction, "associate-lambda-function", "", false, "Associate Lambda Function")
	_connectCmd.Flags().BoolVarP(&_connectAssociateLexBot, "associate-lex-bot", "", false, "Associate Lex Bot")
	_connectCmd.Flags().BoolVarP(&_connectAssociatePhoneNumberContactFlow, "associate-phone-number-contact-flow", "", false, "Associate Phone Number Contact Flow")
	_connectCmd.Flags().BoolVarP(&_connectAssociateQueueQuickConnects, "associate-queue-quick-connects", "", false, "Associate Queue Quick Connects")
	_connectCmd.Flags().BoolVarP(&_connectAssociateRoutingProfileQueues, "associate-routing-profile-queues", "", false, "Associate Routing Profile Queues")
	_connectCmd.Flags().BoolVarP(&_connectAssociateSecurityKey, "associate-security-key", "", false, "Associate Security Key")
	_connectCmd.Flags().BoolVarP(&_connectAssociateSecurityProfiles, "associate-security-profiles", "", false, "Associate Security Profiles")
	_connectCmd.Flags().BoolVarP(&_connectAssociateTrafficDistributionGroupUser, "associate-traffic-distribution-group-user", "", false, "Associate Traffic Distribution Group User")
	_connectCmd.Flags().BoolVarP(&_connectAssociateUserProficiencies, "associate-user-proficiencies", "", false, "Associate User Proficiencies")
	_connectCmd.Flags().BoolVarP(&_connectAssociateWorkspace, "associate-workspace", "", false, "Associate Workspace")
	_connectCmd.Flags().BoolVarP(&_connectBatchAssociateAnalyticsDataSet, "batch-associate-analytics-data-set", "", false, "Batch Associate Analytics Data Set")
	_connectCmd.Flags().BoolVarP(&_connectBatchCreateDataTableValue, "batch-create-data-table-value", "", false, "Batch Create Data Table Value")
	_connectCmd.Flags().BoolVarP(&_connectBatchDeleteDataTableValue, "batch-delete-data-table-value", "", false, "Batch Delete Data Table Value")
	_connectCmd.Flags().BoolVarP(&_connectBatchDescribeDataTableValue, "batch-describe-data-table-value", "", false, "Batch Describe Data Table Value")
	_connectCmd.Flags().BoolVarP(&_connectBatchDisassociateAnalyticsDataSet, "batch-disassociate-analytics-data-set", "", false, "Batch Disassociate Analytics Data Set")
	_connectCmd.Flags().BoolVarP(&_connectBatchGetAttachedFileMetadata, "batch-get-attached-file-metadata", "", false, "Batch Get Attached File Metadata")
	_connectCmd.Flags().BoolVarP(&_connectBatchGetFlowAssociation, "batch-get-flow-association", "", false, "Batch Get Flow Association")
	_connectCmd.Flags().BoolVarP(&_connectBatchPutContact, "batch-put-contact", "", false, "Batch Put Contact")
	_connectCmd.Flags().BoolVarP(&_connectBatchUpdateDataTableValue, "batch-update-data-table-value", "", false, "Batch Update Data Table Value")
	_connectCmd.Flags().BoolVarP(&_connectClaimPhoneNumber, "claim-phone-number", "", false, "Claim Phone Number")
	_connectCmd.Flags().BoolVarP(&_connectCompleteAttachedFileUpload, "complete-attached-file-upload", "", false, "Complete Attached File Upload")
	_connectCmd.Flags().BoolVarP(&_connectCreateAgentStatus, "create-agent-status", "", false, "Create Agent Status")
	_connectCmd.Flags().BoolVarP(&_connectCreateContact, "create-contact", "", false, "Create Contact")
	_connectCmd.Flags().BoolVarP(&_connectCreateContactFlow, "create-contact-flow", "", false, "Create Contact Flow")
	_connectCmd.Flags().BoolVarP(&_connectCreateContactFlowModule, "create-contact-flow-module", "", false, "Create Contact Flow Module")
	_connectCmd.Flags().BoolVarP(&_connectCreateContactFlowModuleAlias, "create-contact-flow-module-alias", "", false, "Create Contact Flow Module Alias")
	_connectCmd.Flags().BoolVarP(&_connectCreateContactFlowModuleVersion, "create-contact-flow-module-version", "", false, "Create Contact Flow Module Version")
	_connectCmd.Flags().BoolVarP(&_connectCreateContactFlowVersion, "create-contact-flow-version", "", false, "Create Contact Flow Version")
	_connectCmd.Flags().BoolVarP(&_connectCreateDataTable, "create-data-table", "", false, "Create Data Table")
	_connectCmd.Flags().BoolVarP(&_connectCreateDataTableAttribute, "create-data-table-attribute", "", false, "Create Data Table Attribute")
	_connectCmd.Flags().BoolVarP(&_connectCreateEmailAddress, "create-email-address", "", false, "Create Email Address")
	_connectCmd.Flags().BoolVarP(&_connectCreateEvaluationForm, "create-evaluation-form", "", false, "Create Evaluation Form")
	_connectCmd.Flags().BoolVarP(&_connectCreateHoursOfOperation, "create-hours-of-operation", "", false, "Create Hours Of Operation")
	_connectCmd.Flags().BoolVarP(&_connectCreateHoursOfOperationOverride, "create-hours-of-operation-override", "", false, "Create Hours Of Operation Override")
	_connectCmd.Flags().BoolVarP(&_connectCreateInstance, "create-instance", "", false, "Create Instance")
	_connectCmd.Flags().BoolVarP(&_connectCreateIntegrationAssociation, "create-integration-association", "", false, "Create Integration Association")
	_connectCmd.Flags().BoolVarP(&_connectCreateNotification, "create-notification", "", false, "Create Notification")
	_connectCmd.Flags().BoolVarP(&_connectCreateParticipant, "create-participant", "", false, "Create Participant")
	_connectCmd.Flags().BoolVarP(&_connectCreatePersistentContactAssociation, "create-persistent-contact-association", "", false, "Create Persistent Contact Association")
	_connectCmd.Flags().BoolVarP(&_connectCreatePredefinedAttribute, "create-predefined-attribute", "", false, "Create Predefined Attribute")
	_connectCmd.Flags().BoolVarP(&_connectCreatePrompt, "create-prompt", "", false, "Create Prompt")
	_connectCmd.Flags().BoolVarP(&_connectCreatePushNotificationRegistration, "create-push-notification-registration", "", false, "Create Push Notification Registration")
	_connectCmd.Flags().BoolVarP(&_connectCreateQueue, "create-queue", "", false, "Create Queue")
	_connectCmd.Flags().BoolVarP(&_connectCreateQuickConnect, "create-quick-connect", "", false, "Create Quick Connect")
	_connectCmd.Flags().BoolVarP(&_connectCreateRoutingProfile, "create-routing-profile", "", false, "Create Routing Profile")
	_connectCmd.Flags().BoolVarP(&_connectCreateRule, "create-rule", "", false, "Create Rule")
	_connectCmd.Flags().BoolVarP(&_connectCreateSecurityProfile, "create-security-profile", "", false, "Create Security Profile")
	_connectCmd.Flags().BoolVarP(&_connectCreateTaskTemplate, "create-task-template", "", false, "Create Task Template")
	_connectCmd.Flags().BoolVarP(&_connectCreateTestCase, "create-test-case", "", false, "Create Test Case")
	_connectCmd.Flags().BoolVarP(&_connectCreateTrafficDistributionGroup, "create-traffic-distribution-group", "", false, "Create Traffic Distribution Group")
	_connectCmd.Flags().BoolVarP(&_connectCreateUseCase, "create-use-case", "", false, "Create Use Case")
	_connectCmd.Flags().BoolVarP(&_connectCreateUser, "create-user", "", false, "Create User")
	_connectCmd.Flags().BoolVarP(&_connectCreateUserHierarchyGroup, "create-user-hierarchy-group", "", false, "Create User Hierarchy Group")
	_connectCmd.Flags().BoolVarP(&_connectCreateView, "create-view", "", false, "Create View")
	_connectCmd.Flags().BoolVarP(&_connectCreateViewVersion, "create-view-version", "", false, "Create View Version")
	_connectCmd.Flags().BoolVarP(&_connectCreateVocabulary, "create-vocabulary", "", false, "Create Vocabulary")
	_connectCmd.Flags().BoolVarP(&_connectCreateWorkspace, "create-workspace", "", false, "Create Workspace")
	_connectCmd.Flags().BoolVarP(&_connectCreateWorkspacePage, "create-workspace-page", "", false, "Create Workspace Page")
	_connectCmd.Flags().BoolVarP(&_connectDeactivateEvaluationForm, "deactivate-evaluation-form", "", false, "Deactivate Evaluation Form")
	_connectCmd.Flags().BoolVarP(&_connectDeleteAttachedFile, "delete-attached-file", "", false, "Delete Attached File")
	_connectCmd.Flags().BoolVarP(&_connectDeleteContactEvaluation, "delete-contact-evaluation", "", false, "Delete Contact Evaluation")
	_connectCmd.Flags().BoolVarP(&_connectDeleteContactFlow, "delete-contact-flow", "", false, "Delete Contact Flow")
	_connectCmd.Flags().BoolVarP(&_connectDeleteContactFlowModule, "delete-contact-flow-module", "", false, "Delete Contact Flow Module")
	_connectCmd.Flags().BoolVarP(&_connectDeleteContactFlowModuleAlias, "delete-contact-flow-module-alias", "", false, "Delete Contact Flow Module Alias")
	_connectCmd.Flags().BoolVarP(&_connectDeleteContactFlowModuleVersion, "delete-contact-flow-module-version", "", false, "Delete Contact Flow Module Version")
	_connectCmd.Flags().BoolVarP(&_connectDeleteContactFlowVersion, "delete-contact-flow-version", "", false, "Delete Contact Flow Version")
	_connectCmd.Flags().BoolVarP(&_connectDeleteDataTable, "delete-data-table", "", false, "Delete Data Table")
	_connectCmd.Flags().BoolVarP(&_connectDeleteDataTableAttribute, "delete-data-table-attribute", "", false, "Delete Data Table Attribute")
	_connectCmd.Flags().BoolVarP(&_connectDeleteEmailAddress, "delete-email-address", "", false, "Delete Email Address")
	_connectCmd.Flags().BoolVarP(&_connectDeleteEvaluationForm, "delete-evaluation-form", "", false, "Delete Evaluation Form")
	_connectCmd.Flags().BoolVarP(&_connectDeleteHoursOfOperation, "delete-hours-of-operation", "", false, "Delete Hours Of Operation")
	_connectCmd.Flags().BoolVarP(&_connectDeleteHoursOfOperationOverride, "delete-hours-of-operation-override", "", false, "Delete Hours Of Operation Override")
	_connectCmd.Flags().BoolVarP(&_connectDeleteInstance, "delete-instance", "", false, "Delete Instance")
	_connectCmd.Flags().BoolVarP(&_connectDeleteIntegrationAssociation, "delete-integration-association", "", false, "Delete Integration Association")
	_connectCmd.Flags().BoolVarP(&_connectDeleteNotification, "delete-notification", "", false, "Delete Notification")
	_connectCmd.Flags().BoolVarP(&_connectDeletePredefinedAttribute, "delete-predefined-attribute", "", false, "Delete Predefined Attribute")
	_connectCmd.Flags().BoolVarP(&_connectDeletePrompt, "delete-prompt", "", false, "Delete Prompt")
	_connectCmd.Flags().BoolVarP(&_connectDeletePushNotificationRegistration, "delete-push-notification-registration", "", false, "Delete Push Notification Registration")
	_connectCmd.Flags().BoolVarP(&_connectDeleteQueue, "delete-queue", "", false, "Delete Queue")
	_connectCmd.Flags().BoolVarP(&_connectDeleteQuickConnect, "delete-quick-connect", "", false, "Delete Quick Connect")
	_connectCmd.Flags().BoolVarP(&_connectDeleteRoutingProfile, "delete-routing-profile", "", false, "Delete Routing Profile")
	_connectCmd.Flags().BoolVarP(&_connectDeleteRule, "delete-rule", "", false, "Delete Rule")
	_connectCmd.Flags().BoolVarP(&_connectDeleteSecurityProfile, "delete-security-profile", "", false, "Delete Security Profile")
	_connectCmd.Flags().BoolVarP(&_connectDeleteTaskTemplate, "delete-task-template", "", false, "Delete Task Template")
	_connectCmd.Flags().BoolVarP(&_connectDeleteTestCase, "delete-test-case", "", false, "Delete Test Case")
	_connectCmd.Flags().BoolVarP(&_connectDeleteTrafficDistributionGroup, "delete-traffic-distribution-group", "", false, "Delete Traffic Distribution Group")
	_connectCmd.Flags().BoolVarP(&_connectDeleteUseCase, "delete-use-case", "", false, "Delete Use Case")
	_connectCmd.Flags().BoolVarP(&_connectDeleteUser, "delete-user", "", false, "Delete User")
	_connectCmd.Flags().BoolVarP(&_connectDeleteUserHierarchyGroup, "delete-user-hierarchy-group", "", false, "Delete User Hierarchy Group")
	_connectCmd.Flags().BoolVarP(&_connectDeleteView, "delete-view", "", false, "Delete View")
	_connectCmd.Flags().BoolVarP(&_connectDeleteViewVersion, "delete-view-version", "", false, "Delete View Version")
	_connectCmd.Flags().BoolVarP(&_connectDeleteVocabulary, "delete-vocabulary", "", false, "Delete Vocabulary")
	_connectCmd.Flags().BoolVarP(&_connectDeleteWorkspace, "delete-workspace", "", false, "Delete Workspace")
	_connectCmd.Flags().BoolVarP(&_connectDeleteWorkspaceMedia, "delete-workspace-media", "", false, "Delete Workspace Media")
	_connectCmd.Flags().BoolVarP(&_connectDeleteWorkspacePage, "delete-workspace-page", "", false, "Delete Workspace Page")
	_connectCmd.Flags().BoolVarP(&_connectDescribeAgentStatus, "describe-agent-status", "", false, "Describe Agent Status")
	_connectCmd.Flags().BoolVarP(&_connectDescribeAuthenticationProfile, "describe-authentication-profile", "", false, "Describe Authentication Profile")
	_connectCmd.Flags().BoolVarP(&_connectDescribeContact, "describe-contact", "", false, "Describe Contact")
	_connectCmd.Flags().BoolVarP(&_connectDescribeContactEvaluation, "describe-contact-evaluation", "", false, "Describe Contact Evaluation")
	_connectCmd.Flags().BoolVarP(&_connectDescribeContactFlow, "describe-contact-flow", "", false, "Describe Contact Flow")
	_connectCmd.Flags().BoolVarP(&_connectDescribeContactFlowModule, "describe-contact-flow-module", "", false, "Describe Contact Flow Module")
	_connectCmd.Flags().BoolVarP(&_connectDescribeContactFlowModuleAlias, "describe-contact-flow-module-alias", "", false, "Describe Contact Flow Module Alias")
	_connectCmd.Flags().BoolVarP(&_connectDescribeDataTable, "describe-data-table", "", false, "Describe Data Table")
	_connectCmd.Flags().BoolVarP(&_connectDescribeDataTableAttribute, "describe-data-table-attribute", "", false, "Describe Data Table Attribute")
	_connectCmd.Flags().BoolVarP(&_connectDescribeEmailAddress, "describe-email-address", "", false, "Describe Email Address")
	_connectCmd.Flags().BoolVarP(&_connectDescribeEvaluationForm, "describe-evaluation-form", "", false, "Describe Evaluation Form")
	_connectCmd.Flags().BoolVarP(&_connectDescribeHoursOfOperation, "describe-hours-of-operation", "", false, "Describe Hours Of Operation")
	_connectCmd.Flags().BoolVarP(&_connectDescribeHoursOfOperationOverride, "describe-hours-of-operation-override", "", false, "Describe Hours Of Operation Override")
	_connectCmd.Flags().BoolVarP(&_connectDescribeInstance, "describe-instance", "", false, "Describe Instance")
	_connectCmd.Flags().BoolVarP(&_connectDescribeInstanceAttribute, "describe-instance-attribute", "", false, "Describe Instance Attribute")
	_connectCmd.Flags().BoolVarP(&_connectDescribeInstanceStorageConfig, "describe-instance-storage-config", "", false, "Describe Instance Storage Config")
	_connectCmd.Flags().BoolVarP(&_connectDescribeNotification, "describe-notification", "", false, "Describe Notification")
	_connectCmd.Flags().BoolVarP(&_connectDescribePhoneNumber, "describe-phone-number", "", false, "Describe Phone Number")
	_connectCmd.Flags().BoolVarP(&_connectDescribePredefinedAttribute, "describe-predefined-attribute", "", false, "Describe Predefined Attribute")
	_connectCmd.Flags().BoolVarP(&_connectDescribePrompt, "describe-prompt", "", false, "Describe Prompt")
	_connectCmd.Flags().BoolVarP(&_connectDescribeQueue, "describe-queue", "", false, "Describe Queue")
	_connectCmd.Flags().BoolVarP(&_connectDescribeQuickConnect, "describe-quick-connect", "", false, "Describe Quick Connect")
	_connectCmd.Flags().BoolVarP(&_connectDescribeRoutingProfile, "describe-routing-profile", "", false, "Describe Routing Profile")
	_connectCmd.Flags().BoolVarP(&_connectDescribeRule, "describe-rule", "", false, "Describe Rule")
	_connectCmd.Flags().BoolVarP(&_connectDescribeSecurityProfile, "describe-security-profile", "", false, "Describe Security Profile")
	_connectCmd.Flags().BoolVarP(&_connectDescribeTestCase, "describe-test-case", "", false, "Describe Test Case")
	_connectCmd.Flags().BoolVarP(&_connectDescribeTrafficDistributionGroup, "describe-traffic-distribution-group", "", false, "Describe Traffic Distribution Group")
	_connectCmd.Flags().BoolVarP(&_connectDescribeUser, "describe-user", "", false, "Describe User")
	_connectCmd.Flags().BoolVarP(&_connectDescribeUserHierarchyGroup, "describe-user-hierarchy-group", "", false, "Describe User Hierarchy Group")
	_connectCmd.Flags().BoolVarP(&_connectDescribeUserHierarchyStructure, "describe-user-hierarchy-structure", "", false, "Describe User Hierarchy Structure")
	_connectCmd.Flags().BoolVarP(&_connectDescribeView, "describe-view", "", false, "Describe View")
	_connectCmd.Flags().BoolVarP(&_connectDescribeVocabulary, "describe-vocabulary", "", false, "Describe Vocabulary")
	_connectCmd.Flags().BoolVarP(&_connectDescribeWorkspace, "describe-workspace", "", false, "Describe Workspace")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateAnalyticsDataSet, "disassociate-analytics-data-set", "", false, "Disassociate Analytics Data Set")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateApprovedOrigin, "disassociate-approved-origin", "", false, "Disassociate Approved Origin")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateBot, "disassociate-bot", "", false, "Disassociate Bot")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateEmailAddressAlias, "disassociate-email-address-alias", "", false, "Disassociate Email Address Alias")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateFlow, "disassociate-flow", "", false, "Disassociate Flow")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateHoursOfOperations, "disassociate-hours-of-operations", "", false, "Disassociate Hours Of Operations")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateInstanceStorageConfig, "disassociate-instance-storage-config", "", false, "Disassociate Instance Storage Config")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateLambdaFunction, "disassociate-lambda-function", "", false, "Disassociate Lambda Function")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateLexBot, "disassociate-lex-bot", "", false, "Disassociate Lex Bot")
	_connectCmd.Flags().BoolVarP(&_connectDisassociatePhoneNumberContactFlow, "disassociate-phone-number-contact-flow", "", false, "Disassociate Phone Number Contact Flow")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateQueueQuickConnects, "disassociate-queue-quick-connects", "", false, "Disassociate Queue Quick Connects")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateRoutingProfileQueues, "disassociate-routing-profile-queues", "", false, "Disassociate Routing Profile Queues")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateSecurityKey, "disassociate-security-key", "", false, "Disassociate Security Key")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateSecurityProfiles, "disassociate-security-profiles", "", false, "Disassociate Security Profiles")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateTrafficDistributionGroupUser, "disassociate-traffic-distribution-group-user", "", false, "Disassociate Traffic Distribution Group User")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateUserProficiencies, "disassociate-user-proficiencies", "", false, "Disassociate User Proficiencies")
	_connectCmd.Flags().BoolVarP(&_connectDisassociateWorkspace, "disassociate-workspace", "", false, "Disassociate Workspace")
	_connectCmd.Flags().BoolVarP(&_connectDismissUserContact, "dismiss-user-contact", "", false, "Dismiss User Contact")
	_connectCmd.Flags().BoolVarP(&_connectEvaluateDataTableValues, "evaluate-data-table-values", "", false, "Evaluate Data Table Values")
	_connectCmd.Flags().BoolVarP(&_connectGetAttachedFile, "get-attached-file", "", false, "Get Attached File")
	_connectCmd.Flags().BoolVarP(&_connectGetContactAttributes, "get-contact-attributes", "", false, "Get Contact Attributes")
	_connectCmd.Flags().BoolVarP(&_connectGetContactMetrics, "get-contact-metrics", "", false, "Get Contact Metrics")
	_connectCmd.Flags().BoolVarP(&_connectGetCurrentMetricData, "get-current-metric-data", "", false, "Get Current Metric Data")
	_connectCmd.Flags().BoolVarP(&_connectGetCurrentUserData, "get-current-user-data", "", false, "Get Current User Data")
	_connectCmd.Flags().BoolVarP(&_connectGetEffectiveHoursOfOperations, "get-effective-hours-of-operations", "", false, "Get Effective Hours Of Operations")
	_connectCmd.Flags().BoolVarP(&_connectGetFederationToken, "get-federation-token", "", false, "Get Federation Token")
	_connectCmd.Flags().BoolVarP(&_connectGetFlowAssociation, "get-flow-association", "", false, "Get Flow Association")
	_connectCmd.Flags().BoolVarP(&_connectGetMetricData, "get-metric-data", "", false, "Get Metric Data")
	_connectCmd.Flags().BoolVarP(&_connectGetMetricDataV2, "get-metric-data-v2", "", false, "Get Metric Data V2")
	_connectCmd.Flags().BoolVarP(&_connectGetPromptFile, "get-prompt-file", "", false, "Get Prompt File")
	_connectCmd.Flags().BoolVarP(&_connectGetTaskTemplate, "get-task-template", "", false, "Get Task Template")
	_connectCmd.Flags().BoolVarP(&_connectGetTestCaseExecutionSummary, "get-test-case-execution-summary", "", false, "Get Test Case Execution Summary")
	_connectCmd.Flags().BoolVarP(&_connectGetTrafficDistribution, "get-traffic-distribution", "", false, "Get Traffic Distribution")
	_connectCmd.Flags().BoolVarP(&_connectImportPhoneNumber, "import-phone-number", "", false, "Import Phone Number")
	_connectCmd.Flags().BoolVarP(&_connectImportWorkspaceMedia, "import-workspace-media", "", false, "Import Workspace Media")
	_connectCmd.Flags().BoolVarP(&_connectListAgentStatuses, "list-agent-statuses", "", false, "List Agent Statuses")
	_connectCmd.Flags().BoolVarP(&_connectListAnalyticsDataAssociations, "list-analytics-data-associations", "", false, "List Analytics Data Associations")
	_connectCmd.Flags().BoolVarP(&_connectListAnalyticsDataLakeDataSets, "list-analytics-data-lake-data-sets", "", false, "List Analytics Data Lake Data Sets")
	_connectCmd.Flags().BoolVarP(&_connectListApprovedOrigins, "list-approved-origins", "", false, "List Approved Origins")
	_connectCmd.Flags().BoolVarP(&_connectListAssociatedContacts, "list-associated-contacts", "", false, "List Associated Contacts")
	_connectCmd.Flags().BoolVarP(&_connectListAuthenticationProfiles, "list-authentication-profiles", "", false, "List Authentication Profiles")
	_connectCmd.Flags().BoolVarP(&_connectListBots, "list-bots", "", false, "List Bots")
	_connectCmd.Flags().BoolVarP(&_connectListChildHoursOfOperations, "list-child-hours-of-operations", "", false, "List Child Hours Of Operations")
	_connectCmd.Flags().BoolVarP(&_connectListContactEvaluations, "list-contact-evaluations", "", false, "List Contact Evaluations")
	_connectCmd.Flags().BoolVarP(&_connectListContactFlowModuleAliases, "list-contact-flow-module-aliases", "", false, "List Contact Flow Module Aliases")
	_connectCmd.Flags().BoolVarP(&_connectListContactFlowModuleVersions, "list-contact-flow-module-versions", "", false, "List Contact Flow Module Versions")
	_connectCmd.Flags().BoolVarP(&_connectListContactFlowModules, "list-contact-flow-modules", "", false, "List Contact Flow Modules")
	_connectCmd.Flags().BoolVarP(&_connectListContactFlowVersions, "list-contact-flow-versions", "", false, "List Contact Flow Versions")
	_connectCmd.Flags().BoolVarP(&_connectListContactFlows, "list-contact-flows", "", false, "List Contact Flows")
	_connectCmd.Flags().BoolVarP(&_connectListContactReferences, "list-contact-references", "", false, "List Contact References")
	_connectCmd.Flags().BoolVarP(&_connectListDataTableAttributes, "list-data-table-attributes", "", false, "List Data Table Attributes")
	_connectCmd.Flags().BoolVarP(&_connectListDataTablePrimaryValues, "list-data-table-primary-values", "", false, "List Data Table Primary Values")
	_connectCmd.Flags().BoolVarP(&_connectListDataTableValues, "list-data-table-values", "", false, "List Data Table Values")
	_connectCmd.Flags().BoolVarP(&_connectListDataTables, "list-data-tables", "", false, "List Data Tables")
	_connectCmd.Flags().BoolVarP(&_connectListDefaultVocabularies, "list-default-vocabularies", "", false, "List Default Vocabularies")
	_connectCmd.Flags().BoolVarP(&_connectListEntitySecurityProfiles, "list-entity-security-profiles", "", false, "List Entity Security Profiles")
	_connectCmd.Flags().BoolVarP(&_connectListEvaluationFormVersions, "list-evaluation-form-versions", "", false, "List Evaluation Form Versions")
	_connectCmd.Flags().BoolVarP(&_connectListEvaluationForms, "list-evaluation-forms", "", false, "List Evaluation Forms")
	_connectCmd.Flags().BoolVarP(&_connectListFlowAssociations, "list-flow-associations", "", false, "List Flow Associations")
	_connectCmd.Flags().BoolVarP(&_connectListHoursOfOperationOverrides, "list-hours-of-operation-overrides", "", false, "List Hours Of Operation Overrides")
	_connectCmd.Flags().BoolVarP(&_connectListHoursOfOperations, "list-hours-of-operations", "", false, "List Hours Of Operations")
	_connectCmd.Flags().BoolVarP(&_connectListInstanceAttributes, "list-instance-attributes", "", false, "List Instance Attributes")
	_connectCmd.Flags().BoolVarP(&_connectListInstanceStorageConfigs, "list-instance-storage-configs", "", false, "List Instance Storage Configs")
	_connectCmd.Flags().BoolVarP(&_connectListInstances, "list-instances", "", false, "List Instances")
	_connectCmd.Flags().BoolVarP(&_connectListIntegrationAssociations, "list-integration-associations", "", false, "List Integration Associations")
	_connectCmd.Flags().BoolVarP(&_connectListLambdaFunctions, "list-lambda-functions", "", false, "List Lambda Functions")
	_connectCmd.Flags().BoolVarP(&_connectListLexBots, "list-lex-bots", "", false, "List Lex Bots")
	_connectCmd.Flags().BoolVarP(&_connectListNotifications, "list-notifications", "", false, "List Notifications")
	_connectCmd.Flags().BoolVarP(&_connectListPhoneNumbers, "list-phone-numbers", "", false, "List Phone Numbers")
	_connectCmd.Flags().BoolVarP(&_connectListPhoneNumbersV2, "list-phone-numbers-v2", "", false, "List Phone Numbers V2")
	_connectCmd.Flags().BoolVarP(&_connectListPredefinedAttributes, "list-predefined-attributes", "", false, "List Predefined Attributes")
	_connectCmd.Flags().BoolVarP(&_connectListPrompts, "list-prompts", "", false, "List Prompts")
	_connectCmd.Flags().BoolVarP(&_connectListQueueQuickConnects, "list-queue-quick-connects", "", false, "List Queue Quick Connects")
	_connectCmd.Flags().BoolVarP(&_connectListQueues, "list-queues", "", false, "List Queues")
	_connectCmd.Flags().BoolVarP(&_connectListQuickConnects, "list-quick-connects", "", false, "List Quick Connects")
	_connectCmd.Flags().BoolVarP(&_connectListRealtimeContactAnalysisSegmentsV2, "list-realtime-contact-analysis-segments-v2", "", false, "List Realtime Contact Analysis Segments V2")
	_connectCmd.Flags().BoolVarP(&_connectListRoutingProfileManualAssignmentQueues, "list-routing-profile-manual-assignment-queues", "", false, "List Routing Profile Manual Assignment Queues")
	_connectCmd.Flags().BoolVarP(&_connectListRoutingProfileQueues, "list-routing-profile-queues", "", false, "List Routing Profile Queues")
	_connectCmd.Flags().BoolVarP(&_connectListRoutingProfiles, "list-routing-profiles", "", false, "List Routing Profiles")
	_connectCmd.Flags().BoolVarP(&_connectListRules, "list-rules", "", false, "List Rules")
	_connectCmd.Flags().BoolVarP(&_connectListSecurityKeys, "list-security-keys", "", false, "List Security Keys")
	_connectCmd.Flags().BoolVarP(&_connectListSecurityProfileApplications, "list-security-profile-applications", "", false, "List Security Profile Applications")
	_connectCmd.Flags().BoolVarP(&_connectListSecurityProfileFlowModules, "list-security-profile-flow-modules", "", false, "List Security Profile Flow Modules")
	_connectCmd.Flags().BoolVarP(&_connectListSecurityProfilePermissions, "list-security-profile-permissions", "", false, "List Security Profile Permissions")
	_connectCmd.Flags().BoolVarP(&_connectListSecurityProfiles, "list-security-profiles", "", false, "List Security Profiles")
	_connectCmd.Flags().BoolVarP(&_connectListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_connectCmd.Flags().BoolVarP(&_connectListTaskTemplates, "list-task-templates", "", false, "List Task Templates")
	_connectCmd.Flags().BoolVarP(&_connectListTestCaseExecutionRecords, "list-test-case-execution-records", "", false, "List Test Case Execution Records")
	_connectCmd.Flags().BoolVarP(&_connectListTestCaseExecutions, "list-test-case-executions", "", false, "List Test Case Executions")
	_connectCmd.Flags().BoolVarP(&_connectListTestCases, "list-test-cases", "", false, "List Test Cases")
	_connectCmd.Flags().BoolVarP(&_connectListTrafficDistributionGroupUsers, "list-traffic-distribution-group-users", "", false, "List Traffic Distribution Group Users")
	_connectCmd.Flags().BoolVarP(&_connectListTrafficDistributionGroups, "list-traffic-distribution-groups", "", false, "List Traffic Distribution Groups")
	_connectCmd.Flags().BoolVarP(&_connectListUseCases, "list-use-cases", "", false, "List Use Cases")
	_connectCmd.Flags().BoolVarP(&_connectListUserHierarchyGroups, "list-user-hierarchy-groups", "", false, "List User Hierarchy Groups")
	_connectCmd.Flags().BoolVarP(&_connectListUserNotifications, "list-user-notifications", "", false, "List User Notifications")
	_connectCmd.Flags().BoolVarP(&_connectListUserProficiencies, "list-user-proficiencies", "", false, "List User Proficiencies")
	_connectCmd.Flags().BoolVarP(&_connectListUsers, "list-users", "", false, "List Users")
	_connectCmd.Flags().BoolVarP(&_connectListViewVersions, "list-view-versions", "", false, "List View Versions")
	_connectCmd.Flags().BoolVarP(&_connectListViews, "list-views", "", false, "List Views")
	_connectCmd.Flags().BoolVarP(&_connectListWorkspaceMedia, "list-workspace-media", "", false, "List Workspace Media")
	_connectCmd.Flags().BoolVarP(&_connectListWorkspacePages, "list-workspace-pages", "", false, "List Workspace Pages")
	_connectCmd.Flags().BoolVarP(&_connectListWorkspaces, "list-workspaces", "", false, "List Workspaces")
	_connectCmd.Flags().BoolVarP(&_connectMonitorContact, "monitor-contact", "", false, "Monitor Contact")
	_connectCmd.Flags().BoolVarP(&_connectPauseContact, "pause-contact", "", false, "Pause Contact")
	_connectCmd.Flags().BoolVarP(&_connectPutUserStatus, "put-user-status", "", false, "Put User Status")
	_connectCmd.Flags().BoolVarP(&_connectReleasePhoneNumber, "release-phone-number", "", false, "Release Phone Number")
	_connectCmd.Flags().BoolVarP(&_connectReplicateInstance, "replicate-instance", "", false, "Replicate Instance")
	_connectCmd.Flags().BoolVarP(&_connectResumeContact, "resume-contact", "", false, "Resume Contact")
	_connectCmd.Flags().BoolVarP(&_connectResumeContactRecording, "resume-contact-recording", "", false, "Resume Contact Recording")
	_connectCmd.Flags().BoolVarP(&_connectSearchAgentStatuses, "search-agent-statuses", "", false, "Search Agent Statuses")
	_connectCmd.Flags().BoolVarP(&_connectSearchAvailablePhoneNumbers, "search-available-phone-numbers", "", false, "Search Available Phone Numbers")
	_connectCmd.Flags().BoolVarP(&_connectSearchContactEvaluations, "search-contact-evaluations", "", false, "Search Contact Evaluations")
	_connectCmd.Flags().BoolVarP(&_connectSearchContactFlowModules, "search-contact-flow-modules", "", false, "Search Contact Flow Modules")
	_connectCmd.Flags().BoolVarP(&_connectSearchContactFlows, "search-contact-flows", "", false, "Search Contact Flows")
	_connectCmd.Flags().BoolVarP(&_connectSearchContacts, "search-contacts", "", false, "Search Contacts")
	_connectCmd.Flags().BoolVarP(&_connectSearchDataTables, "search-data-tables", "", false, "Search Data Tables")
	_connectCmd.Flags().BoolVarP(&_connectSearchEmailAddresses, "search-email-addresses", "", false, "Search Email Addresses")
	_connectCmd.Flags().BoolVarP(&_connectSearchEvaluationForms, "search-evaluation-forms", "", false, "Search Evaluation Forms")
	_connectCmd.Flags().BoolVarP(&_connectSearchHoursOfOperationOverrides, "search-hours-of-operation-overrides", "", false, "Search Hours Of Operation Overrides")
	_connectCmd.Flags().BoolVarP(&_connectSearchHoursOfOperations, "search-hours-of-operations", "", false, "Search Hours Of Operations")
	_connectCmd.Flags().BoolVarP(&_connectSearchNotifications, "search-notifications", "", false, "Search Notifications")
	_connectCmd.Flags().BoolVarP(&_connectSearchPredefinedAttributes, "search-predefined-attributes", "", false, "Search Predefined Attributes")
	_connectCmd.Flags().BoolVarP(&_connectSearchPrompts, "search-prompts", "", false, "Search Prompts")
	_connectCmd.Flags().BoolVarP(&_connectSearchQueues, "search-queues", "", false, "Search Queues")
	_connectCmd.Flags().BoolVarP(&_connectSearchQuickConnects, "search-quick-connects", "", false, "Search Quick Connects")
	_connectCmd.Flags().BoolVarP(&_connectSearchResourceTags, "search-resource-tags", "", false, "Search Resource Tags")
	_connectCmd.Flags().BoolVarP(&_connectSearchRoutingProfiles, "search-routing-profiles", "", false, "Search Routing Profiles")
	_connectCmd.Flags().BoolVarP(&_connectSearchSecurityProfiles, "search-security-profiles", "", false, "Search Security Profiles")
	_connectCmd.Flags().BoolVarP(&_connectSearchTestCases, "search-test-cases", "", false, "Search Test Cases")
	_connectCmd.Flags().BoolVarP(&_connectSearchUserHierarchyGroups, "search-user-hierarchy-groups", "", false, "Search User Hierarchy Groups")
	_connectCmd.Flags().BoolVarP(&_connectSearchUsers, "search-users", "", false, "Search Users")
	_connectCmd.Flags().BoolVarP(&_connectSearchViews, "search-views", "", false, "Search Views")
	_connectCmd.Flags().BoolVarP(&_connectSearchVocabularies, "search-vocabularies", "", false, "Search Vocabularies")
	_connectCmd.Flags().BoolVarP(&_connectSearchWorkspaceAssociations, "search-workspace-associations", "", false, "Search Workspace Associations")
	_connectCmd.Flags().BoolVarP(&_connectSearchWorkspaces, "search-workspaces", "", false, "Search Workspaces")
	_connectCmd.Flags().BoolVarP(&_connectSendChatIntegrationEvent, "send-chat-integration-event", "", false, "Send Chat Integration Event")
	_connectCmd.Flags().BoolVarP(&_connectSendOutboundEmail, "send-outbound-email", "", false, "Send Outbound Email")
	_connectCmd.Flags().BoolVarP(&_connectStartAttachedFileUpload, "start-attached-file-upload", "", false, "Start Attached File Upload")
	_connectCmd.Flags().BoolVarP(&_connectStartChatContact, "start-chat-contact", "", false, "Start Chat Contact")
	_connectCmd.Flags().BoolVarP(&_connectStartContactEvaluation, "start-contact-evaluation", "", false, "Start Contact Evaluation")
	_connectCmd.Flags().BoolVarP(&_connectStartContactMediaProcessing, "start-contact-media-processing", "", false, "Start Contact Media Processing")
	_connectCmd.Flags().BoolVarP(&_connectStartContactRecording, "start-contact-recording", "", false, "Start Contact Recording")
	_connectCmd.Flags().BoolVarP(&_connectStartContactStreaming, "start-contact-streaming", "", false, "Start Contact Streaming")
	_connectCmd.Flags().BoolVarP(&_connectStartEmailContact, "start-email-contact", "", false, "Start Email Contact")
	_connectCmd.Flags().BoolVarP(&_connectStartOutboundChatContact, "start-outbound-chat-contact", "", false, "Start Outbound Chat Contact")
	_connectCmd.Flags().BoolVarP(&_connectStartOutboundEmailContact, "start-outbound-email-contact", "", false, "Start Outbound Email Contact")
	_connectCmd.Flags().BoolVarP(&_connectStartOutboundVoiceContact, "start-outbound-voice-contact", "", false, "Start Outbound Voice Contact")
	_connectCmd.Flags().BoolVarP(&_connectStartScreenSharing, "start-screen-sharing", "", false, "Start Screen Sharing")
	_connectCmd.Flags().BoolVarP(&_connectStartTaskContact, "start-task-contact", "", false, "Start Task Contact")
	_connectCmd.Flags().BoolVarP(&_connectStartTestCaseExecution, "start-test-case-execution", "", false, "Start Test Case Execution")
	_connectCmd.Flags().BoolVarP(&_connectStartWebRTCContact, "start-web-rtc-contact", "", false, "Start Web Rtc Contact")
	_connectCmd.Flags().BoolVarP(&_connectStopContact, "stop-contact", "", false, "Stop Contact")
	_connectCmd.Flags().BoolVarP(&_connectStopContactMediaProcessing, "stop-contact-media-processing", "", false, "Stop Contact Media Processing")
	_connectCmd.Flags().BoolVarP(&_connectStopContactRecording, "stop-contact-recording", "", false, "Stop Contact Recording")
	_connectCmd.Flags().BoolVarP(&_connectStopContactStreaming, "stop-contact-streaming", "", false, "Stop Contact Streaming")
	_connectCmd.Flags().BoolVarP(&_connectStopTestCaseExecution, "stop-test-case-execution", "", false, "Stop Test Case Execution")
	_connectCmd.Flags().BoolVarP(&_connectSubmitContactEvaluation, "submit-contact-evaluation", "", false, "Submit Contact Evaluation")
	_connectCmd.Flags().BoolVarP(&_connectSuspendContactRecording, "suspend-contact-recording", "", false, "Suspend Contact Recording")
	_connectCmd.Flags().BoolVarP(&_connectTagContact, "tag-contact", "", false, "Tag Contact")
	_connectCmd.Flags().BoolVarP(&_connectTagResource, "tag-resource", "", false, "Tag Resource")
	_connectCmd.Flags().BoolVarP(&_connectTransferContact, "transfer-contact", "", false, "Transfer Contact")
	_connectCmd.Flags().BoolVarP(&_connectUntagContact, "untag-contact", "", false, "Untag Contact")
	_connectCmd.Flags().BoolVarP(&_connectUntagResource, "untag-resource", "", false, "Untag Resource")
	_connectCmd.Flags().BoolVarP(&_connectUpdateAgentStatus, "update-agent-status", "", false, "Update Agent Status")
	_connectCmd.Flags().BoolVarP(&_connectUpdateAuthenticationProfile, "update-authentication-profile", "", false, "Update Authentication Profile")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContact, "update-contact", "", false, "Update Contact")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactAttributes, "update-contact-attributes", "", false, "Update Contact Attributes")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactEvaluation, "update-contact-evaluation", "", false, "Update Contact Evaluation")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactFlowContent, "update-contact-flow-content", "", false, "Update Contact Flow Content")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactFlowMetadata, "update-contact-flow-metadata", "", false, "Update Contact Flow Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactFlowModuleAlias, "update-contact-flow-module-alias", "", false, "Update Contact Flow Module Alias")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactFlowModuleContent, "update-contact-flow-module-content", "", false, "Update Contact Flow Module Content")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactFlowModuleMetadata, "update-contact-flow-module-metadata", "", false, "Update Contact Flow Module Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactFlowName, "update-contact-flow-name", "", false, "Update Contact Flow Name")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactRoutingData, "update-contact-routing-data", "", false, "Update Contact Routing Data")
	_connectCmd.Flags().BoolVarP(&_connectUpdateContactSchedule, "update-contact-schedule", "", false, "Update Contact Schedule")
	_connectCmd.Flags().BoolVarP(&_connectUpdateDataTableAttribute, "update-data-table-attribute", "", false, "Update Data Table Attribute")
	_connectCmd.Flags().BoolVarP(&_connectUpdateDataTableMetadata, "update-data-table-metadata", "", false, "Update Data Table Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdateDataTablePrimaryValues, "update-data-table-primary-values", "", false, "Update Data Table Primary Values")
	_connectCmd.Flags().BoolVarP(&_connectUpdateEmailAddressMetadata, "update-email-address-metadata", "", false, "Update Email Address Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdateEvaluationForm, "update-evaluation-form", "", false, "Update Evaluation Form")
	_connectCmd.Flags().BoolVarP(&_connectUpdateHoursOfOperation, "update-hours-of-operation", "", false, "Update Hours Of Operation")
	_connectCmd.Flags().BoolVarP(&_connectUpdateHoursOfOperationOverride, "update-hours-of-operation-override", "", false, "Update Hours Of Operation Override")
	_connectCmd.Flags().BoolVarP(&_connectUpdateInstanceAttribute, "update-instance-attribute", "", false, "Update Instance Attribute")
	_connectCmd.Flags().BoolVarP(&_connectUpdateInstanceStorageConfig, "update-instance-storage-config", "", false, "Update Instance Storage Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdateNotificationContent, "update-notification-content", "", false, "Update Notification Content")
	_connectCmd.Flags().BoolVarP(&_connectUpdateParticipantAuthentication, "update-participant-authentication", "", false, "Update Participant Authentication")
	_connectCmd.Flags().BoolVarP(&_connectUpdateParticipantRoleConfig, "update-participant-role-config", "", false, "Update Participant Role Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdatePhoneNumber, "update-phone-number", "", false, "Update Phone Number")
	_connectCmd.Flags().BoolVarP(&_connectUpdatePhoneNumberMetadata, "update-phone-number-metadata", "", false, "Update Phone Number Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdatePredefinedAttribute, "update-predefined-attribute", "", false, "Update Predefined Attribute")
	_connectCmd.Flags().BoolVarP(&_connectUpdatePrompt, "update-prompt", "", false, "Update Prompt")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQueueHoursOfOperation, "update-queue-hours-of-operation", "", false, "Update Queue Hours Of Operation")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQueueMaxContacts, "update-queue-max-contacts", "", false, "Update Queue Max Contacts")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQueueName, "update-queue-name", "", false, "Update Queue Name")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQueueOutboundCallerConfig, "update-queue-outbound-caller-config", "", false, "Update Queue Outbound Caller Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQueueOutboundEmailConfig, "update-queue-outbound-email-config", "", false, "Update Queue Outbound Email Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQueueStatus, "update-queue-status", "", false, "Update Queue Status")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQuickConnectConfig, "update-quick-connect-config", "", false, "Update Quick Connect Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdateQuickConnectName, "update-quick-connect-name", "", false, "Update Quick Connect Name")
	_connectCmd.Flags().BoolVarP(&_connectUpdateRoutingProfileAgentAvailabilityTimer, "update-routing-profile-agent-availability-timer", "", false, "Update Routing Profile Agent Availability Timer")
	_connectCmd.Flags().BoolVarP(&_connectUpdateRoutingProfileConcurrency, "update-routing-profile-concurrency", "", false, "Update Routing Profile Concurrency")
	_connectCmd.Flags().BoolVarP(&_connectUpdateRoutingProfileDefaultOutboundQueue, "update-routing-profile-default-outbound-queue", "", false, "Update Routing Profile Default Outbound Queue")
	_connectCmd.Flags().BoolVarP(&_connectUpdateRoutingProfileName, "update-routing-profile-name", "", false, "Update Routing Profile Name")
	_connectCmd.Flags().BoolVarP(&_connectUpdateRoutingProfileQueues, "update-routing-profile-queues", "", false, "Update Routing Profile Queues")
	_connectCmd.Flags().BoolVarP(&_connectUpdateRule, "update-rule", "", false, "Update Rule")
	_connectCmd.Flags().BoolVarP(&_connectUpdateSecurityProfile, "update-security-profile", "", false, "Update Security Profile")
	_connectCmd.Flags().BoolVarP(&_connectUpdateTaskTemplate, "update-task-template", "", false, "Update Task Template")
	_connectCmd.Flags().BoolVarP(&_connectUpdateTestCase, "update-test-case", "", false, "Update Test Case")
	_connectCmd.Flags().BoolVarP(&_connectUpdateTrafficDistribution, "update-traffic-distribution", "", false, "Update Traffic Distribution")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserConfig, "update-user-config", "", false, "Update User Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserHierarchy, "update-user-hierarchy", "", false, "Update User Hierarchy")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserHierarchyGroupName, "update-user-hierarchy-group-name", "", false, "Update User Hierarchy Group Name")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserHierarchyStructure, "update-user-hierarchy-structure", "", false, "Update User Hierarchy Structure")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserIdentityInfo, "update-user-identity-info", "", false, "Update User Identity Info")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserNotificationStatus, "update-user-notification-status", "", false, "Update User Notification Status")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserPhoneConfig, "update-user-phone-config", "", false, "Update User Phone Config")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserProficiencies, "update-user-proficiencies", "", false, "Update User Proficiencies")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserRoutingProfile, "update-user-routing-profile", "", false, "Update User Routing Profile")
	_connectCmd.Flags().BoolVarP(&_connectUpdateUserSecurityProfiles, "update-user-security-profiles", "", false, "Update User Security Profiles")
	_connectCmd.Flags().BoolVarP(&_connectUpdateViewContent, "update-view-content", "", false, "Update View Content")
	_connectCmd.Flags().BoolVarP(&_connectUpdateViewMetadata, "update-view-metadata", "", false, "Update View Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdateWorkspaceMetadata, "update-workspace-metadata", "", false, "Update Workspace Metadata")
	_connectCmd.Flags().BoolVarP(&_connectUpdateWorkspacePage, "update-workspace-page", "", false, "Update Workspace Page")
	_connectCmd.Flags().BoolVarP(&_connectUpdateWorkspaceTheme, "update-workspace-theme", "", false, "Update Workspace Theme")
	_connectCmd.Flags().BoolVarP(&_connectUpdateWorkspaceVisibility, "update-workspace-visibility", "", false, "Update Workspace Visibility")

}
