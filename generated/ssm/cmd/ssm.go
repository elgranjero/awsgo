package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssmCmd represents the ssm command
var _ssmCmd = &cobra.Command{
	Use:   "ssm",
	Short: "AWS ssm CLI",
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
		client := ssm.NewFromConfig(cfg)
		if _ssmAddTagsToResource {
			ssm_AddTagsToResource(cfg, client)
			return
		}
		if _ssmAssociateOpsItemRelatedItem {
			ssm_AssociateOpsItemRelatedItem(cfg, client)
			return
		}
		if _ssmCancelCommand {
			ssm_CancelCommand(cfg, client)
			return
		}
		if _ssmCancelMaintenanceWindowExecution {
			ssm_CancelMaintenanceWindowExecution(cfg, client)
			return
		}
		if _ssmCreateActivation {
			ssm_CreateActivation(cfg, client)
			return
		}
		if _ssmCreateAssociation {
			ssm_CreateAssociation(cfg, client)
			return
		}
		if _ssmCreateAssociationBatch {
			ssm_CreateAssociationBatch(cfg, client)
			return
		}
		if _ssmCreateDocument {
			ssm_CreateDocument(cfg, client)
			return
		}
		if _ssmCreateMaintenanceWindow {
			ssm_CreateMaintenanceWindow(cfg, client)
			return
		}
		if _ssmCreateOpsItem {
			ssm_CreateOpsItem(cfg, client)
			return
		}
		if _ssmCreateOpsMetadata {
			ssm_CreateOpsMetadata(cfg, client)
			return
		}
		if _ssmCreatePatchBaseline {
			ssm_CreatePatchBaseline(cfg, client)
			return
		}
		if _ssmCreateResourceDataSync {
			ssm_CreateResourceDataSync(cfg, client)
			return
		}
		if _ssmDeleteActivation {
			ssm_DeleteActivation(cfg, client)
			return
		}
		if _ssmDeleteAssociation {
			ssm_DeleteAssociation(cfg, client)
			return
		}
		if _ssmDeleteDocument {
			ssm_DeleteDocument(cfg, client)
			return
		}
		if _ssmDeleteInventory {
			ssm_DeleteInventory(cfg, client)
			return
		}
		if _ssmDeleteMaintenanceWindow {
			ssm_DeleteMaintenanceWindow(cfg, client)
			return
		}
		if _ssmDeleteOpsItem {
			ssm_DeleteOpsItem(cfg, client)
			return
		}
		if _ssmDeleteOpsMetadata {
			ssm_DeleteOpsMetadata(cfg, client)
			return
		}
		if _ssmDeleteParameter {
			ssm_DeleteParameter(cfg, client)
			return
		}
		if _ssmDeleteParameters {
			ssm_DeleteParameters(cfg, client)
			return
		}
		if _ssmDeletePatchBaseline {
			ssm_DeletePatchBaseline(cfg, client)
			return
		}
		if _ssmDeleteResourceDataSync {
			ssm_DeleteResourceDataSync(cfg, client)
			return
		}
		if _ssmDeleteResourcePolicy {
			ssm_DeleteResourcePolicy(cfg, client)
			return
		}
		if _ssmDeregisterManagedInstance {
			ssm_DeregisterManagedInstance(cfg, client)
			return
		}
		if _ssmDeregisterPatchBaselineForPatchGroup {
			ssm_DeregisterPatchBaselineForPatchGroup(cfg, client)
			return
		}
		if _ssmDeregisterTargetFromMaintenanceWindow {
			ssm_DeregisterTargetFromMaintenanceWindow(cfg, client)
			return
		}
		if _ssmDeregisterTaskFromMaintenanceWindow {
			ssm_DeregisterTaskFromMaintenanceWindow(cfg, client)
			return
		}
		if _ssmDescribeActivations {
			ssm_DescribeActivations(cfg, client)
			return
		}
		if _ssmDescribeAssociation {
			ssm_DescribeAssociation(cfg, client)
			return
		}
		if _ssmDescribeAssociationExecutionTargets {
			ssm_DescribeAssociationExecutionTargets(cfg, client)
			return
		}
		if _ssmDescribeAssociationExecutions {
			ssm_DescribeAssociationExecutions(cfg, client)
			return
		}
		if _ssmDescribeAutomationExecutions {
			ssm_DescribeAutomationExecutions(cfg, client)
			return
		}
		if _ssmDescribeAutomationStepExecutions {
			ssm_DescribeAutomationStepExecutions(cfg, client)
			return
		}
		if _ssmDescribeAvailablePatches {
			ssm_DescribeAvailablePatches(cfg, client)
			return
		}
		if _ssmDescribeDocument {
			ssm_DescribeDocument(cfg, client)
			return
		}
		if _ssmDescribeDocumentPermission {
			ssm_DescribeDocumentPermission(cfg, client)
			return
		}
		if _ssmDescribeEffectiveInstanceAssociations {
			ssm_DescribeEffectiveInstanceAssociations(cfg, client)
			return
		}
		if _ssmDescribeEffectivePatchesForPatchBaseline {
			ssm_DescribeEffectivePatchesForPatchBaseline(cfg, client)
			return
		}
		if _ssmDescribeInstanceAssociationsStatus {
			ssm_DescribeInstanceAssociationsStatus(cfg, client)
			return
		}
		if _ssmDescribeInstanceInformation {
			ssm_DescribeInstanceInformation(cfg, client)
			return
		}
		if _ssmDescribeInstancePatchStates {
			ssm_DescribeInstancePatchStates(cfg, client)
			return
		}
		if _ssmDescribeInstancePatchStatesForPatchGroup {
			ssm_DescribeInstancePatchStatesForPatchGroup(cfg, client)
			return
		}
		if _ssmDescribeInstancePatches {
			ssm_DescribeInstancePatches(cfg, client)
			return
		}
		if _ssmDescribeInstanceProperties {
			ssm_DescribeInstanceProperties(cfg, client)
			return
		}
		if _ssmDescribeInventoryDeletions {
			ssm_DescribeInventoryDeletions(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowExecutionTaskInvocations {
			ssm_DescribeMaintenanceWindowExecutionTaskInvocations(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowExecutionTasks {
			ssm_DescribeMaintenanceWindowExecutionTasks(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowExecutions {
			ssm_DescribeMaintenanceWindowExecutions(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowSchedule {
			ssm_DescribeMaintenanceWindowSchedule(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowTargets {
			ssm_DescribeMaintenanceWindowTargets(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowTasks {
			ssm_DescribeMaintenanceWindowTasks(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindows {
			ssm_DescribeMaintenanceWindows(cfg, client)
			return
		}
		if _ssmDescribeMaintenanceWindowsForTarget {
			ssm_DescribeMaintenanceWindowsForTarget(cfg, client)
			return
		}
		if _ssmDescribeOpsItems {
			ssm_DescribeOpsItems(cfg, client)
			return
		}
		if _ssmDescribeParameters {
			ssm_DescribeParameters(cfg, client)
			return
		}
		if _ssmDescribePatchBaselines {
			ssm_DescribePatchBaselines(cfg, client)
			return
		}
		if _ssmDescribePatchGroupState {
			ssm_DescribePatchGroupState(cfg, client)
			return
		}
		if _ssmDescribePatchGroups {
			ssm_DescribePatchGroups(cfg, client)
			return
		}
		if _ssmDescribePatchProperties {
			ssm_DescribePatchProperties(cfg, client)
			return
		}
		if _ssmDescribeSessions {
			ssm_DescribeSessions(cfg, client)
			return
		}
		if _ssmDisassociateOpsItemRelatedItem {
			ssm_DisassociateOpsItemRelatedItem(cfg, client)
			return
		}
		if _ssmGetAccessToken {
			ssm_GetAccessToken(cfg, client)
			return
		}
		if _ssmGetAutomationExecution {
			ssm_GetAutomationExecution(cfg, client)
			return
		}
		if _ssmGetCalendarState {
			ssm_GetCalendarState(cfg, client)
			return
		}
		if _ssmGetCommandInvocation {
			ssm_GetCommandInvocation(cfg, client)
			return
		}
		if _ssmGetConnectionStatus {
			ssm_GetConnectionStatus(cfg, client)
			return
		}
		if _ssmGetDefaultPatchBaseline {
			ssm_GetDefaultPatchBaseline(cfg, client)
			return
		}
		if _ssmGetDeployablePatchSnapshotForInstance {
			ssm_GetDeployablePatchSnapshotForInstance(cfg, client)
			return
		}
		if _ssmGetDocument {
			ssm_GetDocument(cfg, client)
			return
		}
		if _ssmGetExecutionPreview {
			ssm_GetExecutionPreview(cfg, client)
			return
		}
		if _ssmGetInventory {
			ssm_GetInventory(cfg, client)
			return
		}
		if _ssmGetInventorySchema {
			ssm_GetInventorySchema(cfg, client)
			return
		}
		if _ssmGetMaintenanceWindow {
			ssm_GetMaintenanceWindow(cfg, client)
			return
		}
		if _ssmGetMaintenanceWindowExecution {
			ssm_GetMaintenanceWindowExecution(cfg, client)
			return
		}
		if _ssmGetMaintenanceWindowExecutionTask {
			ssm_GetMaintenanceWindowExecutionTask(cfg, client)
			return
		}
		if _ssmGetMaintenanceWindowExecutionTaskInvocation {
			ssm_GetMaintenanceWindowExecutionTaskInvocation(cfg, client)
			return
		}
		if _ssmGetMaintenanceWindowTask {
			ssm_GetMaintenanceWindowTask(cfg, client)
			return
		}
		if _ssmGetOpsItem {
			ssm_GetOpsItem(cfg, client)
			return
		}
		if _ssmGetOpsMetadata {
			ssm_GetOpsMetadata(cfg, client)
			return
		}
		if _ssmGetOpsSummary {
			ssm_GetOpsSummary(cfg, client)
			return
		}
		if _ssmGetParameter {
			ssm_GetParameter(cfg, client)
			return
		}
		if _ssmGetParameterHistory {
			ssm_GetParameterHistory(cfg, client)
			return
		}
		if _ssmGetParameters {
			ssm_GetParameters(cfg, client)
			return
		}
		if _ssmGetParametersByPath {
			ssm_GetParametersByPath(cfg, client)
			return
		}
		if _ssmGetPatchBaseline {
			ssm_GetPatchBaseline(cfg, client)
			return
		}
		if _ssmGetPatchBaselineForPatchGroup {
			ssm_GetPatchBaselineForPatchGroup(cfg, client)
			return
		}
		if _ssmGetResourcePolicies {
			ssm_GetResourcePolicies(cfg, client)
			return
		}
		if _ssmGetServiceSetting {
			ssm_GetServiceSetting(cfg, client)
			return
		}
		if _ssmLabelParameterVersion {
			ssm_LabelParameterVersion(cfg, client)
			return
		}
		if _ssmListAssociationVersions {
			ssm_ListAssociationVersions(cfg, client)
			return
		}
		if _ssmListAssociations {
			ssm_ListAssociations(cfg, client)
			return
		}
		if _ssmListCommandInvocations {
			ssm_ListCommandInvocations(cfg, client)
			return
		}
		if _ssmListCommands {
			ssm_ListCommands(cfg, client)
			return
		}
		if _ssmListComplianceItems {
			ssm_ListComplianceItems(cfg, client)
			return
		}
		if _ssmListComplianceSummaries {
			ssm_ListComplianceSummaries(cfg, client)
			return
		}
		if _ssmListDocumentMetadataHistory {
			ssm_ListDocumentMetadataHistory(cfg, client)
			return
		}
		if _ssmListDocumentVersions {
			ssm_ListDocumentVersions(cfg, client)
			return
		}
		if _ssmListDocuments {
			ssm_ListDocuments(cfg, client)
			return
		}
		if _ssmListInventoryEntries {
			ssm_ListInventoryEntries(cfg, client)
			return
		}
		if _ssmListNodes {
			ssm_ListNodes(cfg, client)
			return
		}
		if _ssmListNodesSummary {
			ssm_ListNodesSummary(cfg, client)
			return
		}
		if _ssmListOpsItemEvents {
			ssm_ListOpsItemEvents(cfg, client)
			return
		}
		if _ssmListOpsItemRelatedItems {
			ssm_ListOpsItemRelatedItems(cfg, client)
			return
		}
		if _ssmListOpsMetadata {
			ssm_ListOpsMetadata(cfg, client)
			return
		}
		if _ssmListResourceComplianceSummaries {
			ssm_ListResourceComplianceSummaries(cfg, client)
			return
		}
		if _ssmListResourceDataSync {
			ssm_ListResourceDataSync(cfg, client)
			return
		}
		if _ssmListTagsForResource {
			ssm_ListTagsForResource(cfg, client)
			return
		}
		if _ssmModifyDocumentPermission {
			ssm_ModifyDocumentPermission(cfg, client)
			return
		}
		if _ssmPutComplianceItems {
			ssm_PutComplianceItems(cfg, client)
			return
		}
		if _ssmPutInventory {
			ssm_PutInventory(cfg, client)
			return
		}
		if _ssmPutParameter {
			ssm_PutParameter(cfg, client)
			return
		}
		if _ssmPutResourcePolicy {
			ssm_PutResourcePolicy(cfg, client)
			return
		}
		if _ssmRegisterDefaultPatchBaseline {
			ssm_RegisterDefaultPatchBaseline(cfg, client)
			return
		}
		if _ssmRegisterPatchBaselineForPatchGroup {
			ssm_RegisterPatchBaselineForPatchGroup(cfg, client)
			return
		}
		if _ssmRegisterTargetWithMaintenanceWindow {
			ssm_RegisterTargetWithMaintenanceWindow(cfg, client)
			return
		}
		if _ssmRegisterTaskWithMaintenanceWindow {
			ssm_RegisterTaskWithMaintenanceWindow(cfg, client)
			return
		}
		if _ssmRemoveTagsFromResource {
			ssm_RemoveTagsFromResource(cfg, client)
			return
		}
		if _ssmResetServiceSetting {
			ssm_ResetServiceSetting(cfg, client)
			return
		}
		if _ssmResumeSession {
			ssm_ResumeSession(cfg, client)
			return
		}
		if _ssmSendAutomationSignal {
			ssm_SendAutomationSignal(cfg, client)
			return
		}
		if _ssmSendCommand {
			ssm_SendCommand(cfg, client)
			return
		}
		if _ssmStartAccessRequest {
			ssm_StartAccessRequest(cfg, client)
			return
		}
		if _ssmStartAssociationsOnce {
			ssm_StartAssociationsOnce(cfg, client)
			return
		}
		if _ssmStartAutomationExecution {
			ssm_StartAutomationExecution(cfg, client)
			return
		}
		if _ssmStartChangeRequestExecution {
			ssm_StartChangeRequestExecution(cfg, client)
			return
		}
		if _ssmStartExecutionPreview {
			ssm_StartExecutionPreview(cfg, client)
			return
		}
		if _ssmStartSession {
			ssm_StartSession(cfg, client)
			return
		}
		if _ssmStopAutomationExecution {
			ssm_StopAutomationExecution(cfg, client)
			return
		}
		if _ssmTerminateSession {
			ssm_TerminateSession(cfg, client)
			return
		}
		if _ssmUnlabelParameterVersion {
			ssm_UnlabelParameterVersion(cfg, client)
			return
		}
		if _ssmUpdateAssociation {
			ssm_UpdateAssociation(cfg, client)
			return
		}
		if _ssmUpdateAssociationStatus {
			ssm_UpdateAssociationStatus(cfg, client)
			return
		}
		if _ssmUpdateDocument {
			ssm_UpdateDocument(cfg, client)
			return
		}
		if _ssmUpdateDocumentDefaultVersion {
			ssm_UpdateDocumentDefaultVersion(cfg, client)
			return
		}
		if _ssmUpdateDocumentMetadata {
			ssm_UpdateDocumentMetadata(cfg, client)
			return
		}
		if _ssmUpdateMaintenanceWindow {
			ssm_UpdateMaintenanceWindow(cfg, client)
			return
		}
		if _ssmUpdateMaintenanceWindowTarget {
			ssm_UpdateMaintenanceWindowTarget(cfg, client)
			return
		}
		if _ssmUpdateMaintenanceWindowTask {
			ssm_UpdateMaintenanceWindowTask(cfg, client)
			return
		}
		if _ssmUpdateManagedInstanceRole {
			ssm_UpdateManagedInstanceRole(cfg, client)
			return
		}
		if _ssmUpdateOpsItem {
			ssm_UpdateOpsItem(cfg, client)
			return
		}
		if _ssmUpdateOpsMetadata {
			ssm_UpdateOpsMetadata(cfg, client)
			return
		}
		if _ssmUpdatePatchBaseline {
			ssm_UpdatePatchBaseline(cfg, client)
			return
		}
		if _ssmUpdateResourceDataSync {
			ssm_UpdateResourceDataSync(cfg, client)
			return
		}
		if _ssmUpdateServiceSetting {
			ssm_UpdateServiceSetting(cfg, client)
			return
		}

	},
}

var (
	_ssmAddTagsToResource                                 bool
	_ssmAssociateOpsItemRelatedItem                       bool
	_ssmCancelCommand                                     bool
	_ssmCancelMaintenanceWindowExecution                  bool
	_ssmCreateActivation                                  bool
	_ssmCreateAssociation                                 bool
	_ssmCreateAssociationBatch                            bool
	_ssmCreateDocument                                    bool
	_ssmCreateMaintenanceWindow                           bool
	_ssmCreateOpsItem                                     bool
	_ssmCreateOpsMetadata                                 bool
	_ssmCreatePatchBaseline                               bool
	_ssmCreateResourceDataSync                            bool
	_ssmDeleteActivation                                  bool
	_ssmDeleteAssociation                                 bool
	_ssmDeleteDocument                                    bool
	_ssmDeleteInventory                                   bool
	_ssmDeleteMaintenanceWindow                           bool
	_ssmDeleteOpsItem                                     bool
	_ssmDeleteOpsMetadata                                 bool
	_ssmDeleteParameter                                   bool
	_ssmDeleteParameters                                  bool
	_ssmDeletePatchBaseline                               bool
	_ssmDeleteResourceDataSync                            bool
	_ssmDeleteResourcePolicy                              bool
	_ssmDeregisterManagedInstance                         bool
	_ssmDeregisterPatchBaselineForPatchGroup              bool
	_ssmDeregisterTargetFromMaintenanceWindow             bool
	_ssmDeregisterTaskFromMaintenanceWindow               bool
	_ssmDescribeActivations                               bool
	_ssmDescribeAssociation                               bool
	_ssmDescribeAssociationExecutionTargets               bool
	_ssmDescribeAssociationExecutions                     bool
	_ssmDescribeAutomationExecutions                      bool
	_ssmDescribeAutomationStepExecutions                  bool
	_ssmDescribeAvailablePatches                          bool
	_ssmDescribeDocument                                  bool
	_ssmDescribeDocumentPermission                        bool
	_ssmDescribeEffectiveInstanceAssociations             bool
	_ssmDescribeEffectivePatchesForPatchBaseline          bool
	_ssmDescribeInstanceAssociationsStatus                bool
	_ssmDescribeInstanceInformation                       bool
	_ssmDescribeInstancePatchStates                       bool
	_ssmDescribeInstancePatchStatesForPatchGroup          bool
	_ssmDescribeInstancePatches                           bool
	_ssmDescribeInstanceProperties                        bool
	_ssmDescribeInventoryDeletions                        bool
	_ssmDescribeMaintenanceWindowExecutionTaskInvocations bool
	_ssmDescribeMaintenanceWindowExecutionTasks           bool
	_ssmDescribeMaintenanceWindowExecutions               bool
	_ssmDescribeMaintenanceWindowSchedule                 bool
	_ssmDescribeMaintenanceWindowTargets                  bool
	_ssmDescribeMaintenanceWindowTasks                    bool
	_ssmDescribeMaintenanceWindows                        bool
	_ssmDescribeMaintenanceWindowsForTarget               bool
	_ssmDescribeOpsItems                                  bool
	_ssmDescribeParameters                                bool
	_ssmDescribePatchBaselines                            bool
	_ssmDescribePatchGroupState                           bool
	_ssmDescribePatchGroups                               bool
	_ssmDescribePatchProperties                           bool
	_ssmDescribeSessions                                  bool
	_ssmDisassociateOpsItemRelatedItem                    bool
	_ssmGetAccessToken                                    bool
	_ssmGetAutomationExecution                            bool
	_ssmGetCalendarState                                  bool
	_ssmGetCommandInvocation                              bool
	_ssmGetConnectionStatus                               bool
	_ssmGetDefaultPatchBaseline                           bool
	_ssmGetDeployablePatchSnapshotForInstance             bool
	_ssmGetDocument                                       bool
	_ssmGetExecutionPreview                               bool
	_ssmGetInventory                                      bool
	_ssmGetInventorySchema                                bool
	_ssmGetMaintenanceWindow                              bool
	_ssmGetMaintenanceWindowExecution                     bool
	_ssmGetMaintenanceWindowExecutionTask                 bool
	_ssmGetMaintenanceWindowExecutionTaskInvocation       bool
	_ssmGetMaintenanceWindowTask                          bool
	_ssmGetOpsItem                                        bool
	_ssmGetOpsMetadata                                    bool
	_ssmGetOpsSummary                                     bool
	_ssmGetParameter                                      bool
	_ssmGetParameterHistory                               bool
	_ssmGetParameters                                     bool
	_ssmGetParametersByPath                               bool
	_ssmGetPatchBaseline                                  bool
	_ssmGetPatchBaselineForPatchGroup                     bool
	_ssmGetResourcePolicies                               bool
	_ssmGetServiceSetting                                 bool
	_ssmLabelParameterVersion                             bool
	_ssmListAssociationVersions                           bool
	_ssmListAssociations                                  bool
	_ssmListCommandInvocations                            bool
	_ssmListCommands                                      bool
	_ssmListComplianceItems                               bool
	_ssmListComplianceSummaries                           bool
	_ssmListDocumentMetadataHistory                       bool
	_ssmListDocumentVersions                              bool
	_ssmListDocuments                                     bool
	_ssmListInventoryEntries                              bool
	_ssmListNodes                                         bool
	_ssmListNodesSummary                                  bool
	_ssmListOpsItemEvents                                 bool
	_ssmListOpsItemRelatedItems                           bool
	_ssmListOpsMetadata                                   bool
	_ssmListResourceComplianceSummaries                   bool
	_ssmListResourceDataSync                              bool
	_ssmListTagsForResource                               bool
	_ssmModifyDocumentPermission                          bool
	_ssmPutComplianceItems                                bool
	_ssmPutInventory                                      bool
	_ssmPutParameter                                      bool
	_ssmPutResourcePolicy                                 bool
	_ssmRegisterDefaultPatchBaseline                      bool
	_ssmRegisterPatchBaselineForPatchGroup                bool
	_ssmRegisterTargetWithMaintenanceWindow               bool
	_ssmRegisterTaskWithMaintenanceWindow                 bool
	_ssmRemoveTagsFromResource                            bool
	_ssmResetServiceSetting                               bool
	_ssmResumeSession                                     bool
	_ssmSendAutomationSignal                              bool
	_ssmSendCommand                                       bool
	_ssmStartAccessRequest                                bool
	_ssmStartAssociationsOnce                             bool
	_ssmStartAutomationExecution                          bool
	_ssmStartChangeRequestExecution                       bool
	_ssmStartExecutionPreview                             bool
	_ssmStartSession                                      bool
	_ssmStopAutomationExecution                           bool
	_ssmTerminateSession                                  bool
	_ssmUnlabelParameterVersion                           bool
	_ssmUpdateAssociation                                 bool
	_ssmUpdateAssociationStatus                           bool
	_ssmUpdateDocument                                    bool
	_ssmUpdateDocumentDefaultVersion                      bool
	_ssmUpdateDocumentMetadata                            bool
	_ssmUpdateMaintenanceWindow                           bool
	_ssmUpdateMaintenanceWindowTarget                     bool
	_ssmUpdateMaintenanceWindowTask                       bool
	_ssmUpdateManagedInstanceRole                         bool
	_ssmUpdateOpsItem                                     bool
	_ssmUpdateOpsMetadata                                 bool
	_ssmUpdatePatchBaseline                               bool
	_ssmUpdateResourceDataSync                            bool
	_ssmUpdateServiceSetting                              bool

	_ssmAccessRequestId                          string
	_ssmAccountId                                string
	_ssmAccountIdsToAdd                          []string
	_ssmAccountIdsToRemove                       []string
	_ssmActivationId                             string
	_ssmActualEndTime                            string
	_ssmActualStartTime                          string
	_ssmAggregator                               string
	_ssmAggregators                              string
	_ssmAlarmConfiguration                       string
	_ssmAllowUnassociatedTargets                 string
	_ssmAllowedPattern                           string
	_ssmApplyOnlyAtCronInterval                  string
	_ssmApprovalRules                            string
	_ssmApprovedPatches                          []string
	_ssmApprovedPatchesComplianceLevel           string
	_ssmApprovedPatchesEnableNonSecurity         string
	_ssmAssociationDispatchAssumeRole            string
	_ssmAssociationFilterList                    string
	_ssmAssociationId                            string
	_ssmAssociationIds                           []string
	_ssmAssociationName                          string
	_ssmAssociationStatus                        string
	_ssmAssociationType                          string
	_ssmAssociationVersion                       string
	_ssmAtTime                                   string
	_ssmAttachments                              string
	_ssmAutoApprove                              string
	_ssmAutomationExecutionId                    string
	_ssmAutomationTargetParameterName            string
	_ssmAvailableSecurityUpdatesComplianceStatus string
	_ssmBaselineId                               string
	_ssmBaselineOverride                         string
	_ssmCalendarNames                            []string
	_ssmCategory                                 string
	_ssmChangeDetails                            string
	_ssmChangeRequestName                        string
	_ssmClientToken                              string
	_ssmCloudWatchOutputConfig                   string
	_ssmCommandId                                string
	_ssmComment                                  string
	_ssmComplianceSeverity                       string
	_ssmComplianceType                           string
	_ssmContent                                  string
	_ssmCutoff                                   string
	_ssmCutoffBehavior                           string
	_ssmDataType                                 string
	_ssmDefaultInstanceName                      string
	_ssmDeletionId                               string
	_ssmDescription                              string
	_ssmDetails                                  string
	_ssmDisplayName                              string
	_ssmDocumentFilterList                       string
	_ssmDocumentFormat                           string
	_ssmDocumentHash                             string
	_ssmDocumentHashType                         string
	_ssmDocumentName                             string
	_ssmDocumentReviews                          string
	_ssmDocumentType                             string
	_ssmDocumentVersion                          string
	_ssmDryRun                                   string
	_ssmDuration                                 string
	_ssmEnabled                                  string
	_ssmEndDate                                  string
	_ssmEntries                                  string
	_ssmExecutionId                              string
	_ssmExecutionInputs                          string
	_ssmExecutionPreviewId                       string
	_ssmExecutionSummary                         string
	_ssmExpirationDate                           string
	_ssmFilters                                  string
	_ssmFiltersWithOperator                      string
	_ssmForce                                    string
	_ssmGlobalFilters                            string
	_ssmIamRole                                  string
	_ssmInstanceId                               string
	_ssmInstanceIds                              []string
	_ssmInstanceInformationFilterList            string
	_ssmInstancePropertyFilterList               string
	_ssmInvocationId                             string
	_ssmItemContentHash                          string
	_ssmItems                                    string
	_ssmKeyId                                    string
	_ssmKeysToDelete                             []string
	_ssmLabels                                   []string
	_ssmLoggingInfo                              string
	_ssmMaxConcurrency                           string
	_ssmMaxErrors                                string
	_ssmMaxResults                               string
	_ssmMetadata                                 string
	_ssmMetadataToUpdate                         string
	_ssmMode                                     string
	_ssmName                                     string
	_ssmNames                                    []string
	_ssmNextToken                                string
	_ssmNotificationConfig                       string
	_ssmNotifications                            string
	_ssmOperatingSystem                          string
	_ssmOperationalData                          string
	_ssmOperationalDataToDelete                  []string
	_ssmOpsItemArn                               string
	_ssmOpsItemFilters                           string
	_ssmOpsItemId                                string
	_ssmOpsItemType                              string
	_ssmOpsMetadataArn                           string
	_ssmOutputLocation                           string
	_ssmOutputS3BucketName                       string
	_ssmOutputS3KeyPrefix                        string
	_ssmOutputS3Region                           string
	_ssmOverwrite                                string
	_ssmOwnerInformation                         string
	_ssmParameterFilters                         string
	_ssmParameterVersion                         string
	_ssmParameters                               string
	_ssmPatchGroup                               string
	_ssmPatchSet                                 string
	_ssmPath                                     string
	_ssmPayload                                  string
	_ssmPermissionType                           string
	_ssmPlannedEndTime                           string
	_ssmPlannedStartTime                         string
	_ssmPluginName                               string
	_ssmPolicies                                 string
	_ssmPolicy                                   string
	_ssmPolicyHash                               string
	_ssmPolicyId                                 string
	_ssmPriority                                 string
	_ssmProperty                                 string
	_ssmReason                                   string
	_ssmRecursive                                string
	_ssmRegistrationLimit                        string
	_ssmRegistrationMetadata                     string
	_ssmRejectedPatches                          []string
	_ssmRejectedPatchesAction                    string
	_ssmRelatedOpsItems                          string
	_ssmReplace                                  string
	_ssmRequires                                 string
	_ssmResourceArn                              string
	_ssmResourceId                               string
	_ssmResourceIds                              []string
	_ssmResourceType                             string
	_ssmResourceTypes                            []string
	_ssmResourceUri                              string
	_ssmResultAttributes                         string
	_ssmReverseOrder                             string
	_ssmRunbooks                                 string
	_ssmS3Destination                            string
	_ssmSafe                                     string
	_ssmSchedule                                 string
	_ssmScheduleExpression                       string
	_ssmScheduleOffset                           string
	_ssmScheduleTimezone                         string
	_ssmScheduledEndTime                         string
	_ssmScheduledTime                            string
	_ssmSchemaDeleteOption                       string
	_ssmServiceRoleArn                           string
	_ssmSessionId                                string
	_ssmSettingId                                string
	_ssmSettingValue                             string
	_ssmSeverity                                 string
	_ssmShared                                   string
	_ssmSharedDocumentVersion                    string
	_ssmSignalType                               string
	_ssmSnapshotId                               string
	_ssmSource                                   string
	_ssmSources                                  string
	_ssmStartDate                                string
	_ssmState                                    string
	_ssmStatus                                   string
	_ssmSubType                                  string
	_ssmSyncCompliance                           string
	_ssmSyncName                                 string
	_ssmSyncSource                               string
	_ssmSyncType                                 string
	_ssmTagKeys                                  []string
	_ssmTags                                     string
	_ssmTarget                                   string
	_ssmTargetLocations                          string
	_ssmTargetLocationsURL                       string
	_ssmTargetMaps                               string
	_ssmTargetParameterName                      string
	_ssmTargetType                               string
	_ssmTargets                                  string
	_ssmTaskArn                                  string
	_ssmTaskId                                   string
	_ssmTaskInvocationParameters                 string
	_ssmTaskParameters                           string
	_ssmTaskType                                 string
	_ssmTier                                     string
	_ssmTimeoutSeconds                           string
	_ssmTitle                                    string
	_ssmType                                     string
	_ssmTypeName                                 string
	_ssmUploadType                               string
	_ssmUseS3DualStackEndpoint                   string
	_ssmValue                                    string
	_ssmVersionName                              string
	_ssmWindowExecutionId                        string
	_ssmWindowId                                 string
	_ssmWindowTargetId                           string
	_ssmWindowTaskId                             string
	_ssmWithDecryption                           string
)

// Adds or overwrites one or more tags for the specified resource. Tags are
// metadata that you can assign to your automations, documents, managed nodes,
// maintenance windows, Parameter Store parameters, and patch baselines. Tags
// enable you to categorize your resources in different ways, for example, by
// purpose, owner, or environment. Each tag consists of a key and an optional
// value, both of which you define. For example, you could define a set of tags for
// your account's managed nodes that helps you track each node's owner and stack
// level. For example:
//
// - Key=Owner,Value=DbAdmin
//
// - Key=Owner,Value=SysAdmin
//
// - Key=Owner,Value=Dev
//
// - Key=Stack,Value=Production
//
// - Key=Stack,Value=Pre-Production
//
// - Key=Stack,Value=Test
//
// Most resources can have a maximum of 50 tags. Automations can have a maximum of
// 5 tags.
//
// We recommend that you devise a set of tag keys that meets your needs for each
// resource type. Using a consistent set of tag keys makes it easier for you to
// manage your resources. You can search and filter the resources based on the tags
// you add. Tags don't have any semantic meaning to and are interpreted strictly as
// a string of characters.
//
// For more information about using tags with Amazon Elastic Compute Cloud (Amazon
// EC2) instances, see [Tag your Amazon EC2 resources]in the Amazon EC2 User Guide.
//
// [Tag your Amazon EC2 resources]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html
func ssm_AddTagsToResource(cfg aws.Config, client *ssm.Client) {
	input := &ssm.AddTagsToResourceInput{
		// ResourceId: *string, // Required
		// ResourceType: types.ResourceTypeForTagging, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ssmResourceId) > 0 {
		input.ResourceId = aws.String(_ssmResourceId)
	}
	if len(_ssmResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _ssmResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
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

// Associates a related item to a Systems Manager OpsCenter OpsItem. For example,
// you can associate an Incident Manager incident or analysis with an OpsItem.
// Incident Manager and OpsCenter are tools in Amazon Web Services Systems Manager.
func ssm_AssociateOpsItemRelatedItem(cfg aws.Config, client *ssm.Client) {
	input := &ssm.AssociateOpsItemRelatedItemInput{
		// AssociationType: *string, // Required
		// OpsItemId: *string, // Required
		// ResourceType: *string, // Required
		// ResourceUri: *string, // Required
	}

	if len(_ssmAssociationType) > 0 {
		input.AssociationType = aws.String(_ssmAssociationType)
	}
	if len(_ssmOpsItemId) > 0 {
		input.OpsItemId = aws.String(_ssmOpsItemId)
	}
	if len(_ssmResourceType) > 0 {
		input.ResourceType = aws.String(_ssmResourceType)
	}
	if len(_ssmResourceUri) > 0 {
		input.ResourceUri = aws.String(_ssmResourceUri)
	}

	if resp, err := client.AssociateOpsItemRelatedItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel the command specified by the Command ID. There is no
// guarantee that the command will be terminated and the underlying process
// stopped.
func ssm_CancelCommand(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CancelCommandInput{
		// CommandId: *string, // Required
	}

	if len(_ssmCommandId) > 0 {
		input.CommandId = aws.String(_ssmCommandId)
	}
	if len(_ssmInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _ssmInstanceIds...)
	}

	if resp, err := client.CancelCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a maintenance window execution that is already in progress and cancels
// any tasks in the window that haven't already starting running. Tasks already in
// progress will continue to completion.
func ssm_CancelMaintenanceWindowExecution(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CancelMaintenanceWindowExecutionInput{
		// WindowExecutionId: *string, // Required
	}

	if len(_ssmWindowExecutionId) > 0 {
		input.WindowExecutionId = aws.String(_ssmWindowExecutionId)
	}

	if resp, err := client.CancelMaintenanceWindowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an activation code and activation ID you can use to register your
// on-premises servers, edge devices, or virtual machine (VM) with Amazon Web
// Services Systems Manager. Registering these machines with Systems Manager makes
// it possible to manage them using Systems Manager tools. You use the activation
// code and ID when installing SSM Agent on machines in your hybrid environment.
// For more information about requirements for managing on-premises machines using
// Systems Manager, see [Using Amazon Web Services Systems Manager in hybrid and multicloud environments]in the Amazon Web Services Systems Manager User Guide.
//
// Amazon Elastic Compute Cloud (Amazon EC2) instances, edge devices, and
// on-premises servers and VMs that are configured for Systems Manager are all
// called managed nodes.
//
// [Using Amazon Web Services Systems Manager in hybrid and multicloud environments]: https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-hybrid-multicloud.html
func ssm_CreateActivation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateActivationInput{
		// IamRole: *string, // Required
	}

	if len(_ssmIamRole) > 0 {
		input.IamRole = aws.String(_ssmIamRole)
	}
	if len(_ssmDefaultInstanceName) > 0 {
		input.DefaultInstanceName = aws.String(_ssmDefaultInstanceName)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmExpirationDate) > 0 {
		if err := assignInputField(input, "ExpirationDate", _ssmExpirationDate); err != nil {
			log.Errorf("invalid --expiration-date: %s", err.Error())
			return
		}
	}
	if len(_ssmRegistrationLimit) > 0 {
		if err := assignInputField(input, "RegistrationLimit", _ssmRegistrationLimit); err != nil {
			log.Errorf("invalid --registration-limit: %s", err.Error())
			return
		}
	}
	if len(_ssmRegistrationMetadata) > 0 {
		if err := assignInputField(input, "RegistrationMetadata", _ssmRegistrationMetadata); err != nil {
			log.Errorf("invalid --registration-metadata: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateActivation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A State Manager association defines the state that you want to maintain on your
// managed nodes. For example, an association can specify that anti-virus software
// must be installed and running on your managed nodes, or that certain ports must
// be closed. For static targets, the association specifies a schedule for when the
// configuration is reapplied. For dynamic targets, such as an Amazon Web Services
// resource group or an Amazon Web Services autoscaling group, State Manager, a
// tool in Amazon Web Services Systems Manager applies the configuration when new
// managed nodes are added to the group. The association also specifies actions to
// take when applying the configuration. For example, an association for anti-virus
// software might run once a day. If the software isn't installed, then State
// Manager installs it. If the software is installed, but the service isn't
// running, then the association might instruct State Manager to start the service.
func ssm_CreateAssociation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateAssociationInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _ssmAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssmApplyOnlyAtCronInterval) > 0 {
		if err := assignInputField(input, "ApplyOnlyAtCronInterval", _ssmApplyOnlyAtCronInterval); err != nil {
			log.Errorf("invalid --apply-only-at-cron-interval: %s", err.Error())
			return
		}
	}
	if len(_ssmAssociationDispatchAssumeRole) > 0 {
		input.AssociationDispatchAssumeRole = aws.String(_ssmAssociationDispatchAssumeRole)
	}
	if len(_ssmAssociationName) > 0 {
		input.AssociationName = aws.String(_ssmAssociationName)
	}
	if len(_ssmAutomationTargetParameterName) > 0 {
		input.AutomationTargetParameterName = aws.String(_ssmAutomationTargetParameterName)
	}
	if len(_ssmCalendarNames) > 0 {
		input.CalendarNames = append([]string(nil), _ssmCalendarNames...)
	}
	if len(_ssmComplianceSeverity) > 0 {
		if err := assignInputField(input, "ComplianceSeverity", _ssmComplianceSeverity); err != nil {
			log.Errorf("invalid --compliance-severity: %s", err.Error())
			return
		}
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmDuration) > 0 {
		if err := assignInputField(input, "Duration", _ssmDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmMaxConcurrency) > 0 {
		input.MaxConcurrency = aws.String(_ssmMaxConcurrency)
	}
	if len(_ssmMaxErrors) > 0 {
		input.MaxErrors = aws.String(_ssmMaxErrors)
	}
	if len(_ssmOutputLocation) > 0 {
		if err := assignInputField(input, "OutputLocation", _ssmOutputLocation); err != nil {
			log.Errorf("invalid --output-location: %s", err.Error())
			return
		}
	}
	if len(_ssmParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_ssmScheduleExpression)
	}
	if len(_ssmScheduleOffset) > 0 {
		if err := assignInputField(input, "ScheduleOffset", _ssmScheduleOffset); err != nil {
			log.Errorf("invalid --schedule-offset: %s", err.Error())
			return
		}
	}
	if len(_ssmSyncCompliance) > 0 {
		if err := assignInputField(input, "SyncCompliance", _ssmSyncCompliance); err != nil {
			log.Errorf("invalid --sync-compliance: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetLocations) > 0 {
		if err := assignInputField(input, "TargetLocations", _ssmTargetLocations); err != nil {
			log.Errorf("invalid --target-locations: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetMaps) > 0 {
		if err := assignInputField(input, "TargetMaps", _ssmTargetMaps); err != nil {
			log.Errorf("invalid --target-maps: %s", err.Error())
			return
		}
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified Amazon Web Services Systems Manager document (SSM
// document) with the specified managed nodes or targets.
//
// When you associate a document with one or more managed nodes using IDs or tags,
// Amazon Web Services Systems Manager Agent (SSM Agent) running on the managed
// node processes the document and configures the node as specified.
//
// If you associate a document with a managed node that already has an associated
// document, the system returns the AssociationAlreadyExists exception.
func ssm_CreateAssociationBatch(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateAssociationBatchInput{
		// Entries: []types.CreateAssociationBatchRequestEntry, // Required
	}

	if len(_ssmEntries) > 0 {
		if err := assignInputField(input, "Entries", _ssmEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_ssmAssociationDispatchAssumeRole) > 0 {
		input.AssociationDispatchAssumeRole = aws.String(_ssmAssociationDispatchAssumeRole)
	}

	if resp, err := client.CreateAssociationBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Amazon Web Services Systems Manager (SSM document). An SSM document
// defines the actions that Systems Manager performs on your managed nodes. For
// more information about SSM documents, including information about supported
// schemas, features, and syntax, see [Amazon Web Services Systems Manager Documents]in the Amazon Web Services Systems Manager
// User Guide.
//
// [Amazon Web Services Systems Manager Documents]: https://docs.aws.amazon.com/systems-manager/latest/userguide/documents.html
func ssm_CreateDocument(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateDocumentInput{
		// Content: *string, // Required
		// Name: *string, // Required
	}

	if len(_ssmContent) > 0 {
		input.Content = aws.String(_ssmContent)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _ssmAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_ssmDisplayName) > 0 {
		input.DisplayName = aws.String(_ssmDisplayName)
	}
	if len(_ssmDocumentFormat) > 0 {
		if err := assignInputField(input, "DocumentFormat", _ssmDocumentFormat); err != nil {
			log.Errorf("invalid --document-format: %s", err.Error())
			return
		}
	}
	if len(_ssmDocumentType) > 0 {
		if err := assignInputField(input, "DocumentType", _ssmDocumentType); err != nil {
			log.Errorf("invalid --document-type: %s", err.Error())
			return
		}
	}
	if len(_ssmRequires) > 0 {
		if err := assignInputField(input, "Requires", _ssmRequires); err != nil {
			log.Errorf("invalid --requires: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetType) > 0 {
		input.TargetType = aws.String(_ssmTargetType)
	}
	if len(_ssmVersionName) > 0 {
		input.VersionName = aws.String(_ssmVersionName)
	}

	if resp, err := client.CreateDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new maintenance window.
// The value you specify for Duration determines the specific end time for the
// maintenance window based on the time it begins. No maintenance window tasks are
// permitted to start after the resulting endtime minus the number of hours you
// specify for Cutoff . For example, if the maintenance window starts at 3 PM, the
// duration is three hours, and the value you specify for Cutoff is one hour, no
// maintenance window tasks can start after 5 PM.
func ssm_CreateMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateMaintenanceWindowInput{
		// AllowUnassociatedTargets: bool, // Required
		// Cutoff: int32, // Required
		// Duration: *int32, // Required
		// Name: *string, // Required
		// Schedule: *string, // Required
	}

	if len(_ssmAllowUnassociatedTargets) > 0 {
		if err := assignInputField(input, "AllowUnassociatedTargets", _ssmAllowUnassociatedTargets); err != nil {
			log.Errorf("invalid --allow-unassociated-targets: %s", err.Error())
			return
		}
	}
	if len(_ssmCutoff) > 0 {
		if err := assignInputField(input, "Cutoff", _ssmCutoff); err != nil {
			log.Errorf("invalid --cutoff: %s", err.Error())
			return
		}
	}
	if len(_ssmDuration) > 0 {
		if err := assignInputField(input, "Duration", _ssmDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmSchedule) > 0 {
		input.Schedule = aws.String(_ssmSchedule)
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmEndDate) > 0 {
		input.EndDate = aws.String(_ssmEndDate)
	}
	if len(_ssmScheduleOffset) > 0 {
		if err := assignInputField(input, "ScheduleOffset", _ssmScheduleOffset); err != nil {
			log.Errorf("invalid --schedule-offset: %s", err.Error())
			return
		}
	}
	if len(_ssmScheduleTimezone) > 0 {
		input.ScheduleTimezone = aws.String(_ssmScheduleTimezone)
	}
	if len(_ssmStartDate) > 0 {
		input.StartDate = aws.String(_ssmStartDate)
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new OpsItem. You must have permission in Identity and Access
// Management (IAM) to create a new OpsItem. For more information, see [Set up OpsCenter]in the
// Amazon Web Services Systems Manager User Guide.
//
// Operations engineers and IT professionals use Amazon Web Services Systems
// Manager OpsCenter to view, investigate, and remediate operational issues
// impacting the performance and health of their Amazon Web Services resources. For
// more information, see [Amazon Web Services Systems Manager OpsCenter]in the Amazon Web Services Systems Manager User Guide.
//
// [Amazon Web Services Systems Manager OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html
// [Set up OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html
func ssm_CreateOpsItem(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateOpsItemInput{
		// Description: *string, // Required
		// Source: *string, // Required
		// Title: *string, // Required
	}

	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmSource) > 0 {
		input.Source = aws.String(_ssmSource)
	}
	if len(_ssmTitle) > 0 {
		input.Title = aws.String(_ssmTitle)
	}
	if len(_ssmAccountId) > 0 {
		input.AccountId = aws.String(_ssmAccountId)
	}
	if len(_ssmActualEndTime) > 0 {
		if err := assignInputField(input, "ActualEndTime", _ssmActualEndTime); err != nil {
			log.Errorf("invalid --actual-end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmActualStartTime) > 0 {
		if err := assignInputField(input, "ActualStartTime", _ssmActualStartTime); err != nil {
			log.Errorf("invalid --actual-start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmCategory) > 0 {
		input.Category = aws.String(_ssmCategory)
	}
	if len(_ssmNotifications) > 0 {
		if err := assignInputField(input, "Notifications", _ssmNotifications); err != nil {
			log.Errorf("invalid --notifications: %s", err.Error())
			return
		}
	}
	if len(_ssmOperationalData) > 0 {
		if err := assignInputField(input, "OperationalData", _ssmOperationalData); err != nil {
			log.Errorf("invalid --operational-data: %s", err.Error())
			return
		}
	}
	if len(_ssmOpsItemType) > 0 {
		input.OpsItemType = aws.String(_ssmOpsItemType)
	}
	if len(_ssmPlannedEndTime) > 0 {
		if err := assignInputField(input, "PlannedEndTime", _ssmPlannedEndTime); err != nil {
			log.Errorf("invalid --planned-end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmPlannedStartTime) > 0 {
		if err := assignInputField(input, "PlannedStartTime", _ssmPlannedStartTime); err != nil {
			log.Errorf("invalid --planned-start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmPriority) > 0 {
		if err := assignInputField(input, "Priority", _ssmPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_ssmRelatedOpsItems) > 0 {
		if err := assignInputField(input, "RelatedOpsItems", _ssmRelatedOpsItems); err != nil {
			log.Errorf("invalid --related-ops-items: %s", err.Error())
			return
		}
	}
	if len(_ssmSeverity) > 0 {
		input.Severity = aws.String(_ssmSeverity)
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOpsItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If you create a new application in Application Manager, Amazon Web Services
// Systems Manager calls this API operation to specify information about the new
// application, including the application type.
func ssm_CreateOpsMetadata(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateOpsMetadataInput{
		// ResourceId: *string, // Required
	}

	if len(_ssmResourceId) > 0 {
		input.ResourceId = aws.String(_ssmResourceId)
	}
	if len(_ssmMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _ssmMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOpsMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a patch baseline.
// For information about valid key-value pairs in PatchFilters for each supported
// operating system type, see PatchFilter.
func ssm_CreatePatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreatePatchBaselineInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmApprovalRules) > 0 {
		if err := assignInputField(input, "ApprovalRules", _ssmApprovalRules); err != nil {
			log.Errorf("invalid --approval-rules: %s", err.Error())
			return
		}
	}
	if len(_ssmApprovedPatches) > 0 {
		input.ApprovedPatches = append([]string(nil), _ssmApprovedPatches...)
	}
	if len(_ssmApprovedPatchesComplianceLevel) > 0 {
		if err := assignInputField(input, "ApprovedPatchesComplianceLevel", _ssmApprovedPatchesComplianceLevel); err != nil {
			log.Errorf("invalid --approved-patches-compliance-level: %s", err.Error())
			return
		}
	}
	if len(_ssmApprovedPatchesEnableNonSecurity) > 0 {
		if err := assignInputField(input, "ApprovedPatchesEnableNonSecurity", _ssmApprovedPatchesEnableNonSecurity); err != nil {
			log.Errorf("invalid --approved-patches-enable-non-security: %s", err.Error())
			return
		}
	}
	if len(_ssmAvailableSecurityUpdatesComplianceStatus) > 0 {
		if err := assignInputField(input, "AvailableSecurityUpdatesComplianceStatus", _ssmAvailableSecurityUpdatesComplianceStatus); err != nil {
			log.Errorf("invalid --available-security-updates-compliance-status: %s", err.Error())
			return
		}
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmGlobalFilters) > 0 {
		if err := assignInputField(input, "GlobalFilters", _ssmGlobalFilters); err != nil {
			log.Errorf("invalid --global-filters: %s", err.Error())
			return
		}
	}
	if len(_ssmOperatingSystem) > 0 {
		if err := assignInputField(input, "OperatingSystem", _ssmOperatingSystem); err != nil {
			log.Errorf("invalid --operating-system: %s", err.Error())
			return
		}
	}
	if len(_ssmRejectedPatches) > 0 {
		input.RejectedPatches = append([]string(nil), _ssmRejectedPatches...)
	}
	if len(_ssmRejectedPatchesAction) > 0 {
		if err := assignInputField(input, "RejectedPatchesAction", _ssmRejectedPatchesAction); err != nil {
			log.Errorf("invalid --rejected-patches-action: %s", err.Error())
			return
		}
	}
	if len(_ssmSources) > 0 {
		if err := assignInputField(input, "Sources", _ssmSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePatchBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A resource data sync helps you view data from multiple sources in a single
// location. Amazon Web Services Systems Manager offers two types of resource data
// sync: SyncToDestination and SyncFromSource .
//
// You can configure Systems Manager Inventory to use the SyncToDestination type
// to synchronize Inventory data from multiple Amazon Web Services Regions to a
// single Amazon Simple Storage Service (Amazon S3) bucket. For more information,
// see [Creating a resource data sync for Inventory]in the Amazon Web Services Systems Manager User Guide.
//
// You can configure Systems Manager Explorer to use the SyncFromSource type to
// synchronize operational work items (OpsItems) and operational data (OpsData)
// from multiple Amazon Web Services Regions to a single Amazon S3 bucket. This
// type can synchronize OpsItems and OpsData from multiple Amazon Web Services
// accounts and Amazon Web Services Regions or EntireOrganization by using
// Organizations. For more information, see [Setting up Systems Manager Explorer to display data from multiple accounts and Regions]in the Amazon Web Services Systems
// Manager User Guide.
//
// A resource data sync is an asynchronous operation that returns immediately.
// After a successful initial sync is completed, the system continuously syncs
// data. To check the status of a sync, use the ListResourceDataSync.
//
// By default, data isn't encrypted in Amazon S3. We strongly recommend that you
// enable encryption in Amazon S3 to ensure secure data storage. We also recommend
// that you secure access to the Amazon S3 bucket by creating a restrictive bucket
// policy.
//
// [Setting up Systems Manager Explorer to display data from multiple accounts and Regions]: https://docs.aws.amazon.com/systems-manager/latest/userguide/Explorer-resource-data-sync.html
// [Creating a resource data sync for Inventory]: https://docs.aws.amazon.com/systems-manager/latest/userguide/inventory-create-resource-data-sync.html
func ssm_CreateResourceDataSync(cfg aws.Config, client *ssm.Client) {
	input := &ssm.CreateResourceDataSyncInput{
		// SyncName: *string, // Required
	}

	if len(_ssmSyncName) > 0 {
		input.SyncName = aws.String(_ssmSyncName)
	}
	if len(_ssmS3Destination) > 0 {
		if err := assignInputField(input, "S3Destination", _ssmS3Destination); err != nil {
			log.Errorf("invalid --s3-destination: %s", err.Error())
			return
		}
	}
	if len(_ssmSyncSource) > 0 {
		if err := assignInputField(input, "SyncSource", _ssmSyncSource); err != nil {
			log.Errorf("invalid --sync-source: %s", err.Error())
			return
		}
	}
	if len(_ssmSyncType) > 0 {
		input.SyncType = aws.String(_ssmSyncType)
	}

	if resp, err := client.CreateResourceDataSync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an activation. You aren't required to delete an activation. If you
// delete an activation, you can no longer use it to register additional managed
// nodes. Deleting an activation doesn't de-register managed nodes. You must
// manually de-register managed nodes.
func ssm_DeleteActivation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteActivationInput{
		// ActivationId: *string, // Required
	}

	if len(_ssmActivationId) > 0 {
		input.ActivationId = aws.String(_ssmActivationId)
	}

	if resp, err := client.DeleteActivation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified Amazon Web Services Systems Manager document (SSM
// document) from the specified managed node. If you created the association by
// using the Targets parameter, then you must delete the association by using the
// association ID.
//
// When you disassociate a document from a managed node, it doesn't change the
// configuration of the node. To change the configuration state of a managed node
// after you disassociate a document, you must create a new document with the
// desired configuration and associate it with the node.
func ssm_DeleteAssociation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteAssociationInput{}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}

	if resp, err := client.DeleteAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Amazon Web Services Systems Manager document (SSM document) and all
// managed node associations to the document.
//
// Before you delete the document, we recommend that you use DeleteAssociation to disassociate all
// managed nodes that are associated with the document.
func ssm_DeleteDocument(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteDocumentInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmForce) > 0 {
		if err := assignInputField(input, "Force", _ssmForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_ssmVersionName) > 0 {
		input.VersionName = aws.String(_ssmVersionName)
	}

	if resp, err := client.DeleteDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a custom inventory type or the data associated with a custom Inventory
// type. Deleting a custom inventory type is also referred to as deleting a custom
// inventory schema.
func ssm_DeleteInventory(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteInventoryInput{
		// TypeName: *string, // Required
	}

	if len(_ssmTypeName) > 0 {
		input.TypeName = aws.String(_ssmTypeName)
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _ssmDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_ssmSchemaDeleteOption) > 0 {
		if err := assignInputField(input, "SchemaDeleteOption", _ssmSchemaDeleteOption); err != nil {
			log.Errorf("invalid --schema-delete-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteInventory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a maintenance window.
func ssm_DeleteMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteMaintenanceWindowInput{
		// WindowId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}

	if resp, err := client.DeleteMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an OpsItem. You must have permission in Identity and Access Management
// (IAM) to delete an OpsItem.
//
// Note the following important information about this operation.
//
// - Deleting an OpsItem is irreversible. You can't restore a deleted OpsItem.
//
// - This operation uses an eventual consistency model, which means the system
// can take a few minutes to complete this operation. If you delete an OpsItem and
// immediately call, for example, GetOpsItem, the deleted OpsItem might still appear in
// the response.
//
// - This operation is idempotent. The system doesn't throw an exception if you
// repeatedly call this operation for the same OpsItem. If the first call is
// successful, all additional calls return the same successful response as the
// first call.
//
// - This operation doesn't support cross-account calls. A delegated
// administrator or management account can't delete OpsItems in other accounts,
// even if OpsCenter has been set up for cross-account administration. For more
// information about cross-account administration, see [Setting up OpsCenter to centrally manage OpsItems across accounts]in the Systems Manager
// User Guide.
//
// [Setting up OpsCenter to centrally manage OpsItems across accounts]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setting-up-cross-account.html
func ssm_DeleteOpsItem(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteOpsItemInput{
		// OpsItemId: *string, // Required
	}

	if len(_ssmOpsItemId) > 0 {
		input.OpsItemId = aws.String(_ssmOpsItemId)
	}

	if resp, err := client.DeleteOpsItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete OpsMetadata related to an application.
func ssm_DeleteOpsMetadata(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteOpsMetadataInput{
		// OpsMetadataArn: *string, // Required
	}

	if len(_ssmOpsMetadataArn) > 0 {
		input.OpsMetadataArn = aws.String(_ssmOpsMetadataArn)
	}

	if resp, err := client.DeleteOpsMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a parameter from the system. After deleting a parameter, wait for at
// least 30 seconds to create a parameter with the same name.
func ssm_DeleteParameter(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteParameterInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}

	if resp, err := client.DeleteParameter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a list of parameters. After deleting a parameter, wait for at least 30
// seconds to create a parameter with the same name.
func ssm_DeleteParameters(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteParametersInput{
		// Names: []string, // Required
	}

	if len(_ssmNames) > 0 {
		input.Names = append([]string(nil), _ssmNames...)
	}

	if resp, err := client.DeleteParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a patch baseline.
func ssm_DeletePatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeletePatchBaselineInput{
		// BaselineId: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}

	if resp, err := client.DeletePatchBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource data sync configuration. After the configuration is deleted,
// changes to data on managed nodes are no longer synced to or from the target.
// Deleting a sync configuration doesn't delete data.
func ssm_DeleteResourceDataSync(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteResourceDataSyncInput{
		// SyncName: *string, // Required
	}

	if len(_ssmSyncName) > 0 {
		input.SyncName = aws.String(_ssmSyncName)
	}
	if len(_ssmSyncType) > 0 {
		input.SyncType = aws.String(_ssmSyncType)
	}

	if resp, err := client.DeleteResourceDataSync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Systems Manager resource policy. A resource policy helps you to
// define the IAM entity (for example, an Amazon Web Services account) that can
// manage your Systems Manager resources. The following resources support Systems
// Manager resource policies.
//
// - OpsItemGroup - The resource policy for OpsItemGroup enables Amazon Web
// Services accounts to view and interact with OpsCenter operational work items
// (OpsItems).
//
// - Parameter - The resource policy is used to share a parameter with other
// accounts using Resource Access Manager (RAM). For more information about
// cross-account sharing of parameters, see [Working with shared parameters]in the Amazon Web Services Systems
// Manager User Guide.
//
// [Working with shared parameters]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-shared-parameters.html
func ssm_DeleteResourcePolicy(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeleteResourcePolicyInput{
		// PolicyHash: *string, // Required
		// PolicyId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_ssmPolicyHash) > 0 {
		input.PolicyHash = aws.String(_ssmPolicyHash)
	}
	if len(_ssmPolicyId) > 0 {
		input.PolicyId = aws.String(_ssmPolicyId)
	}
	if len(_ssmResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the server or virtual machine from the list of registered servers.
// If you want to reregister an on-premises server, edge device, or VM, you must
// use a different Activation Code and Activation ID than used to register the
// machine previously. The Activation Code and Activation ID must not have already
// been used on the maximum number of activations specified when they were created.
// For more information, see [Deregistering managed nodes in a hybrid and multicloud environment]in the Amazon Web Services Systems Manager User Guide.
//
// [Deregistering managed nodes in a hybrid and multicloud environment]: https://docs.aws.amazon.com/systems-manager/latest/userguide/fleet-manager-deregister-hybrid-nodes.html
func ssm_DeregisterManagedInstance(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeregisterManagedInstanceInput{
		// InstanceId: *string, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}

	if resp, err := client.DeregisterManagedInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a patch group from a patch baseline.
func ssm_DeregisterPatchBaselineForPatchGroup(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeregisterPatchBaselineForPatchGroupInput{
		// BaselineId: *string, // Required
		// PatchGroup: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}
	if len(_ssmPatchGroup) > 0 {
		input.PatchGroup = aws.String(_ssmPatchGroup)
	}

	if resp, err := client.DeregisterPatchBaselineForPatchGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a target from a maintenance window.
func ssm_DeregisterTargetFromMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeregisterTargetFromMaintenanceWindowInput{
		// WindowId: *string, // Required
		// WindowTargetId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmWindowTargetId) > 0 {
		input.WindowTargetId = aws.String(_ssmWindowTargetId)
	}
	if len(_ssmSafe) > 0 {
		if err := assignInputField(input, "Safe", _ssmSafe); err != nil {
			log.Errorf("invalid --safe: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterTargetFromMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a task from a maintenance window.
func ssm_DeregisterTaskFromMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DeregisterTaskFromMaintenanceWindowInput{
		// WindowId: *string, // Required
		// WindowTaskId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmWindowTaskId) > 0 {
		input.WindowTaskId = aws.String(_ssmWindowTaskId)
	}

	if resp, err := client.DeregisterTaskFromMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes details about the activation, such as the date and time the
// activation was created, its expiration date, the Identity and Access Management
// (IAM) role assigned to the managed nodes in the activation, and the number of
// nodes registered by using this activation.
func ssm_DescribeActivations(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeActivationsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeActivations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeActivationsOutput
	p := ssm.NewDescribeActivationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Describes the association for the specified target or managed node. If you
// created the association by using the Targets parameter, then you must retrieve
// the association by using the association ID.
func ssm_DescribeAssociation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeAssociationInput{}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmAssociationVersion) > 0 {
		input.AssociationVersion = aws.String(_ssmAssociationVersion)
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}

	if resp, err := client.DescribeAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Views information about a specific execution of a specific association.
func ssm_DescribeAssociationExecutionTargets(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeAssociationExecutionTargetsInput{
		// AssociationId: *string, // Required
		// ExecutionId: *string, // Required
	}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmExecutionId) > 0 {
		input.ExecutionId = aws.String(_ssmExecutionId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAssociationExecutionTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeAssociationExecutionTargetsOutput
	p := ssm.NewDescribeAssociationExecutionTargetsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Views all executions for a specific association ID.
func ssm_DescribeAssociationExecutions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeAssociationExecutionsInput{
		// AssociationId: *string, // Required
	}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAssociationExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeAssociationExecutionsOutput
	p := ssm.NewDescribeAssociationExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Provides details about all active and terminated Automation executions.
func ssm_DescribeAutomationExecutions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeAutomationExecutionsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAutomationExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeAutomationExecutionsOutput
	p := ssm.NewDescribeAutomationExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Information about all active and terminated step executions in an Automation
// workflow.
func ssm_DescribeAutomationStepExecutions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeAutomationStepExecutionsInput{
		// AutomationExecutionId: *string, // Required
	}

	if len(_ssmAutomationExecutionId) > 0 {
		input.AutomationExecutionId = aws.String(_ssmAutomationExecutionId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _ssmReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeAutomationStepExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeAutomationStepExecutionsOutput
	p := ssm.NewDescribeAutomationStepExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all patches eligible to be included in a patch baseline.
// Currently, DescribeAvailablePatches supports only the Amazon Linux 1, Amazon
// Linux 2, and Windows Server operating systems.
func ssm_DescribeAvailablePatches(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeAvailablePatchesInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAvailablePatches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeAvailablePatchesOutput
	p := ssm.NewDescribeAvailablePatchesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Describes the specified Amazon Web Services Systems Manager document (SSM
// document).
func ssm_DescribeDocument(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeDocumentInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmVersionName) > 0 {
		input.VersionName = aws.String(_ssmVersionName)
	}

	if resp, err := client.DescribeDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the permissions for a Amazon Web Services Systems Manager document
// (SSM document). If you created the document, you are the owner. If a document is
// shared, it can either be shared privately (by specifying a user's Amazon Web
// Services account ID) or publicly (All).
func ssm_DescribeDocumentPermission(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeDocumentPermissionInput{
		// Name: *string, // Required
		// PermissionType: types.DocumentPermissionType, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _ssmPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if resp, err := client.DescribeDocumentPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// All associations for the managed nodes.
func ssm_DescribeEffectiveInstanceAssociations(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeEffectiveInstanceAssociationsInput{
		// InstanceId: *string, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEffectiveInstanceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeEffectiveInstanceAssociationsOutput
	p := ssm.NewDescribeEffectiveInstanceAssociationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the current effective patches (the patch and the approval state) for
// the specified patch baseline. Applies to patch baselines for Windows only.
func ssm_DescribeEffectivePatchesForPatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeEffectivePatchesForPatchBaselineInput{
		// BaselineId: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEffectivePatchesForPatchBaseline(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeEffectivePatchesForPatchBaselineOutput
	p := ssm.NewDescribeEffectivePatchesForPatchBaselinePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// The status of the associations for the managed nodes.
func ssm_DescribeInstanceAssociationsStatus(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInstanceAssociationsStatusInput{
		// InstanceId: *string, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstanceAssociationsStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInstanceAssociationsStatusOutput
	p := ssm.NewDescribeInstanceAssociationsStatusPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Provides information about one or more of your managed nodes, including the
// operating system platform, SSM Agent version, association status, and IP
// address. This operation does not return information for nodes that are either
// Stopped or Terminated.
//
// If you specify one or more node IDs, the operation returns information for
// those managed nodes. If you don't specify node IDs, it returns information for
// all your managed nodes. If you specify a node ID that isn't valid or a node that
// you don't own, you receive an error.
//
// The IamRole field returned for this API operation is the role assigned to an
// Amazon EC2 instance configured with a Systems Manager Quick Setup host
// management configuration or the role assigned to an on-premises managed node.
func ssm_DescribeInstanceInformation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInstanceInformationInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmInstanceInformationFilterList) > 0 {
		if err := assignInputField(input, "InstanceInformationFilterList", _ssmInstanceInformationFilterList); err != nil {
			log.Errorf("invalid --instance-information-filter-list: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstanceInformation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInstanceInformationOutput
	p := ssm.NewDescribeInstanceInformationPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the high-level patch state of one or more managed nodes.
func ssm_DescribeInstancePatchStates(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInstancePatchStatesInput{
		// InstanceIds: []string, // Required
	}

	if len(_ssmInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _ssmInstanceIds...)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstancePatchStates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInstancePatchStatesOutput
	p := ssm.NewDescribeInstancePatchStatesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the high-level patch state for the managed nodes in the specified
// patch group.
func ssm_DescribeInstancePatchStatesForPatchGroup(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInstancePatchStatesForPatchGroupInput{
		// PatchGroup: *string, // Required
	}

	if len(_ssmPatchGroup) > 0 {
		input.PatchGroup = aws.String(_ssmPatchGroup)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstancePatchStatesForPatchGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInstancePatchStatesForPatchGroupOutput
	p := ssm.NewDescribeInstancePatchStatesForPatchGroupPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves information about the patches on the specified managed node and their
// state relative to the patch baseline being used for the node.
func ssm_DescribeInstancePatches(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInstancePatchesInput{
		// InstanceId: *string, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstancePatches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInstancePatchesOutput
	p := ssm.NewDescribeInstancePatchesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// An API operation used by the Systems Manager console to display information
// about Systems Manager managed nodes.
func ssm_DescribeInstanceProperties(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInstancePropertiesInput{}

	if len(_ssmFiltersWithOperator) > 0 {
		if err := assignInputField(input, "FiltersWithOperator", _ssmFiltersWithOperator); err != nil {
			log.Errorf("invalid --filters-with-operator: %s", err.Error())
			return
		}
	}
	if len(_ssmInstancePropertyFilterList) > 0 {
		if err := assignInputField(input, "InstancePropertyFilterList", _ssmInstancePropertyFilterList); err != nil {
			log.Errorf("invalid --instance-property-filter-list: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInstanceProperties(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInstancePropertiesOutput
	p := ssm.NewDescribeInstancePropertiesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Describes a specific delete inventory operation.
func ssm_DescribeInventoryDeletions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeInventoryDeletionsInput{}

	if len(_ssmDeletionId) > 0 {
		input.DeletionId = aws.String(_ssmDeletionId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInventoryDeletions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeInventoryDeletionsOutput
	p := ssm.NewDescribeInventoryDeletionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the individual task executions (one per target) for a particular task
// run as part of a maintenance window execution.
func ssm_DescribeMaintenanceWindowExecutionTaskInvocations(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput{
		// TaskId: *string, // Required
		// WindowExecutionId: *string, // Required
	}

	if len(_ssmTaskId) > 0 {
		input.TaskId = aws.String(_ssmTaskId)
	}
	if len(_ssmWindowExecutionId) > 0 {
		input.WindowExecutionId = aws.String(_ssmWindowExecutionId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowExecutionTaskInvocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput
	p := ssm.NewDescribeMaintenanceWindowExecutionTaskInvocationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// For a given maintenance window execution, lists the tasks that were run.
func ssm_DescribeMaintenanceWindowExecutionTasks(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowExecutionTasksInput{
		// WindowExecutionId: *string, // Required
	}

	if len(_ssmWindowExecutionId) > 0 {
		input.WindowExecutionId = aws.String(_ssmWindowExecutionId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowExecutionTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowExecutionTasksOutput
	p := ssm.NewDescribeMaintenanceWindowExecutionTasksPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the executions of a maintenance window. This includes information about
// when the maintenance window was scheduled to be active, and information about
// tasks registered and run with the maintenance window.
func ssm_DescribeMaintenanceWindowExecutions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowExecutionsInput{
		// WindowId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowExecutionsOutput
	p := ssm.NewDescribeMaintenanceWindowExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves information about upcoming executions of a maintenance window.
func ssm_DescribeMaintenanceWindowSchedule(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowScheduleInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _ssmResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowSchedule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowScheduleOutput
	p := ssm.NewDescribeMaintenanceWindowSchedulePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the targets registered with the maintenance window.
func ssm_DescribeMaintenanceWindowTargets(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowTargetsInput{
		// WindowId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowTargetsOutput
	p := ssm.NewDescribeMaintenanceWindowTargetsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the tasks in a maintenance window.
// For maintenance window tasks without a specified target, you can't supply
// values for --max-errors and --max-concurrency . Instead, the system inserts a
// placeholder value of 1 , which may be reported in the response to this command.
// These values don't affect the running of your task and can be ignored.
func ssm_DescribeMaintenanceWindowTasks(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowTasksInput{
		// WindowId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowTasksOutput
	p := ssm.NewDescribeMaintenanceWindowTasksPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the maintenance windows in an Amazon Web Services account.
func ssm_DescribeMaintenanceWindows(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowsOutput
	p := ssm.NewDescribeMaintenanceWindowsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves information about the maintenance window targets or tasks that a
// managed node is associated with.
func ssm_DescribeMaintenanceWindowsForTarget(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeMaintenanceWindowsForTargetInput{
		// ResourceType: types.MaintenanceWindowResourceType, // Required
		// Targets: []types.Target, // Required
	}

	if len(_ssmResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _ssmResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMaintenanceWindowsForTarget(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeMaintenanceWindowsForTargetOutput
	p := ssm.NewDescribeMaintenanceWindowsForTargetPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Query a set of OpsItems. You must have permission in Identity and Access
// Management (IAM) to query a list of OpsItems. For more information, see [Set up OpsCenter]in the
// Amazon Web Services Systems Manager User Guide.
//
// Operations engineers and IT professionals use Amazon Web Services Systems
// Manager OpsCenter to view, investigate, and remediate operational issues
// impacting the performance and health of their Amazon Web Services resources. For
// more information, see [Amazon Web Services Systems Manager OpsCenter]in the Amazon Web Services Systems Manager User Guide.
//
// [Amazon Web Services Systems Manager OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html
// [Set up OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html
func ssm_DescribeOpsItems(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeOpsItemsInput{}

	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmOpsItemFilters) > 0 {
		if err := assignInputField(input, "OpsItemFilters", _ssmOpsItemFilters); err != nil {
			log.Errorf("invalid --ops-item-filters: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeOpsItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeOpsItemsOutput
	p := ssm.NewDescribeOpsItemsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the parameters in your Amazon Web Services account or the parameters
// shared with you when you enable the [Shared]option.
//
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults . If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken . You can specify the NextToken in a subsequent call to get the
// next set of results.
//
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
//
// If you change the KMS key alias for the KMS key used to encrypt a parameter,
// then you must also update the key alias the parameter uses to reference KMS.
// Otherwise, DescribeParameters retrieves whatever the original key alias was
// referencing.
//
// [Shared]: https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribeParameters.html#systemsmanager-DescribeParameters-request-Shared
func ssm_DescribeParameters(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeParametersInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmParameterFilters) > 0 {
		if err := assignInputField(input, "ParameterFilters", _ssmParameterFilters); err != nil {
			log.Errorf("invalid --parameter-filters: %s", err.Error())
			return
		}
	}
	if len(_ssmShared) > 0 {
		if err := assignInputField(input, "Shared", _ssmShared); err != nil {
			log.Errorf("invalid --shared: %s", err.Error())
			return
		}
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

	var results []*ssm.DescribeParametersOutput
	p := ssm.NewDescribeParametersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the patch baselines in your Amazon Web Services account.
func ssm_DescribePatchBaselines(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribePatchBaselinesInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribePatchBaselines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribePatchBaselinesOutput
	p := ssm.NewDescribePatchBaselinesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns high-level aggregated patch compliance state information for a patch
// group.
func ssm_DescribePatchGroupState(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribePatchGroupStateInput{
		// PatchGroup: *string, // Required
	}

	if len(_ssmPatchGroup) > 0 {
		input.PatchGroup = aws.String(_ssmPatchGroup)
	}

	if resp, err := client.DescribePatchGroupState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all patch groups that have been registered with patch baselines.
func ssm_DescribePatchGroups(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribePatchGroupsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribePatchGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribePatchGroupsOutput
	p := ssm.NewDescribePatchGroupsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the properties of available patches organized by product, product family,
// classification, severity, and other properties of available patches. You can use
// the reported properties in the filters you specify in requests for operations
// such as CreatePatchBaseline, UpdatePatchBaseline, DescribeAvailablePatches, and DescribePatchBaselines.
//
// The following section lists the properties that can be used in filters for each
// major operating system type:
//
// AMAZON_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// AMAZON_LINUX_2 Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// AMAZON_LINUX_2023 Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// CENTOS Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// DEBIAN Valid properties: PRODUCT | PRIORITY
//
// MACOS Valid properties: PRODUCT | CLASSIFICATION
//
// ORACLE_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// REDHAT_ENTERPRISE_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// SUSE Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// UBUNTU Valid properties: PRODUCT | PRIORITY
//
// WINDOWS Valid properties: PRODUCT | PRODUCT_FAMILY | CLASSIFICATION |
// MSRC_SEVERITY
func ssm_DescribePatchProperties(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribePatchPropertiesInput{
		// OperatingSystem: types.OperatingSystem, // Required
		// Property: types.PatchProperty, // Required
	}

	if len(_ssmOperatingSystem) > 0 {
		if err := assignInputField(input, "OperatingSystem", _ssmOperatingSystem); err != nil {
			log.Errorf("invalid --operating-system: %s", err.Error())
			return
		}
	}
	if len(_ssmProperty) > 0 {
		if err := assignInputField(input, "Property", _ssmProperty); err != nil {
			log.Errorf("invalid --property: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmPatchSet) > 0 {
		if err := assignInputField(input, "PatchSet", _ssmPatchSet); err != nil {
			log.Errorf("invalid --patch-set: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribePatchProperties(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribePatchPropertiesOutput
	p := ssm.NewDescribePatchPropertiesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves a list of all active sessions (both connected and disconnected) or
// terminated sessions from the past 30 days.
func ssm_DescribeSessions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DescribeSessionsInput{
		// State: types.SessionState, // Required
	}

	if len(_ssmState) > 0 {
		if err := assignInputField(input, "State", _ssmState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.DescribeSessionsOutput
	p := ssm.NewDescribeSessionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Deletes the association between an OpsItem and a related item. For example,
// this API operation can delete an Incident Manager incident from an OpsItem.
// Incident Manager is a tool in Amazon Web Services Systems Manager.
func ssm_DisassociateOpsItemRelatedItem(cfg aws.Config, client *ssm.Client) {
	input := &ssm.DisassociateOpsItemRelatedItemInput{
		// AssociationId: *string, // Required
		// OpsItemId: *string, // Required
	}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmOpsItemId) > 0 {
		input.OpsItemId = aws.String(_ssmOpsItemId)
	}

	if resp, err := client.DisassociateOpsItemRelatedItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a credentials set to be used with just-in-time node access.
func ssm_GetAccessToken(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetAccessTokenInput{
		// AccessRequestId: *string, // Required
	}

	if len(_ssmAccessRequestId) > 0 {
		input.AccessRequestId = aws.String(_ssmAccessRequestId)
	}

	if resp, err := client.GetAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed information about a particular Automation execution.
func ssm_GetAutomationExecution(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetAutomationExecutionInput{
		// AutomationExecutionId: *string, // Required
	}

	if len(_ssmAutomationExecutionId) > 0 {
		input.AutomationExecutionId = aws.String(_ssmAutomationExecutionId)
	}

	if resp, err := client.GetAutomationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the state of a Amazon Web Services Systems Manager change calendar at the
// current time or a specified time. If you specify a time, GetCalendarState
// returns the state of the calendar at that specific time, and returns the next
// time that the change calendar state will transition. If you don't specify a
// time, GetCalendarState uses the current time. Change Calendar entries have two
// possible states: OPEN or CLOSED .
//
// If you specify more than one calendar in a request, the command returns the
// status of OPEN only if all calendars in the request are open. If one or more
// calendars in the request are closed, the status returned is CLOSED .
//
// For more information about Change Calendar, a tool in Amazon Web Services
// Systems Manager, see [Amazon Web Services Systems Manager Change Calendar]in the Amazon Web Services Systems Manager User Guide.
//
// [Amazon Web Services Systems Manager Change Calendar]: https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar.html
func ssm_GetCalendarState(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetCalendarStateInput{
		// CalendarNames: []string, // Required
	}

	if len(_ssmCalendarNames) > 0 {
		input.CalendarNames = append([]string(nil), _ssmCalendarNames...)
	}
	if len(_ssmAtTime) > 0 {
		input.AtTime = aws.String(_ssmAtTime)
	}

	if resp, err := client.GetCalendarState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about command execution for an invocation or
// plugin. The Run Command API follows an eventual consistency model, due to the
// distributed nature of the system supporting the API. This means that the result
// of an API command you run that affects your resources might not be immediately
// visible to all subsequent commands you run. You should keep this in mind when
// you carry out an API command that immediately follows a previous API command.
//
// GetCommandInvocation only gives the execution status of a plugin in a document.
// To get the command execution status on a specific managed node, use ListCommandInvocations. To get
// the command execution status across managed nodes, use ListCommands.
func ssm_GetCommandInvocation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetCommandInvocationInput{
		// CommandId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_ssmCommandId) > 0 {
		input.CommandId = aws.String(_ssmCommandId)
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmPluginName) > 0 {
		input.PluginName = aws.String(_ssmPluginName)
	}

	if resp, err := client.GetCommandInvocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Session Manager connection status for a managed node to determine
// whether it is running and ready to receive Session Manager connections.
func ssm_GetConnectionStatus(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetConnectionStatusInput{
		// Target: *string, // Required
	}

	if len(_ssmTarget) > 0 {
		input.Target = aws.String(_ssmTarget)
	}

	if resp, err := client.GetConnectionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the default patch baseline. Amazon Web Services Systems Manager
// supports creating multiple default patch baselines. For example, you can create
// a default patch baseline for each operating system.
//
// If you don't specify an operating system value, the default patch baseline for
// Windows is returned.
func ssm_GetDefaultPatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetDefaultPatchBaselineInput{}

	if len(_ssmOperatingSystem) > 0 {
		if err := assignInputField(input, "OperatingSystem", _ssmOperatingSystem); err != nil {
			log.Errorf("invalid --operating-system: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDefaultPatchBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current snapshot for the patch baseline the managed node uses.
// This API is primarily used by the AWS-RunPatchBaseline Systems Manager document
// (SSM document).
//
// If you run the command locally, such as with the Command Line Interface (CLI),
// the system attempts to use your local Amazon Web Services credentials and the
// operation fails. To avoid this, you can run the command in the Amazon Web
// Services Systems Manager console. Use Run Command, a tool in Amazon Web Services
// Systems Manager, with an SSM document that enables you to target a managed node
// with a script or command. For example, run the command using the
// AWS-RunShellScript document or the AWS-RunPowerShellScript document.
func ssm_GetDeployablePatchSnapshotForInstance(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetDeployablePatchSnapshotForInstanceInput{
		// InstanceId: *string, // Required
		// SnapshotId: *string, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmSnapshotId) > 0 {
		input.SnapshotId = aws.String(_ssmSnapshotId)
	}
	if len(_ssmBaselineOverride) > 0 {
		if err := assignInputField(input, "BaselineOverride", _ssmBaselineOverride); err != nil {
			log.Errorf("invalid --baseline-override: %s", err.Error())
			return
		}
	}
	if len(_ssmUseS3DualStackEndpoint) > 0 {
		if err := assignInputField(input, "UseS3DualStackEndpoint", _ssmUseS3DualStackEndpoint); err != nil {
			log.Errorf("invalid --use-s3-dual-stack-endpoint: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDeployablePatchSnapshotForInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the contents of the specified Amazon Web Services Systems Manager document
// (SSM document).
func ssm_GetDocument(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetDocumentInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmDocumentFormat) > 0 {
		if err := assignInputField(input, "DocumentFormat", _ssmDocumentFormat); err != nil {
			log.Errorf("invalid --document-format: %s", err.Error())
			return
		}
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmVersionName) > 0 {
		input.VersionName = aws.String(_ssmVersionName)
	}

	if resp, err := client.GetDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the process of retrieving an existing preview that shows the effects
// that running a specified Automation runbook would have on the targeted
// resources.
func ssm_GetExecutionPreview(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetExecutionPreviewInput{
		// ExecutionPreviewId: *string, // Required
	}

	if len(_ssmExecutionPreviewId) > 0 {
		input.ExecutionPreviewId = aws.String(_ssmExecutionPreviewId)
	}

	if resp, err := client.GetExecutionPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Query inventory information. This includes managed node status, such as Stopped
// or Terminated .
func ssm_GetInventory(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetInventoryInput{}

	if len(_ssmAggregators) > 0 {
		if err := assignInputField(input, "Aggregators", _ssmAggregators); err != nil {
			log.Errorf("invalid --aggregators: %s", err.Error())
			return
		}
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmResultAttributes) > 0 {
		if err := assignInputField(input, "ResultAttributes", _ssmResultAttributes); err != nil {
			log.Errorf("invalid --result-attributes: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetInventory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.GetInventoryOutput
	p := ssm.NewGetInventoryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Return a list of inventory type names for the account, or return a list of
// attribute names for a specific Inventory item type.
func ssm_GetInventorySchema(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetInventorySchemaInput{}

	if len(_ssmAggregator) > 0 {
		if err := assignInputField(input, "Aggregator", _ssmAggregator); err != nil {
			log.Errorf("invalid --aggregator: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmSubType) > 0 {
		if err := assignInputField(input, "SubType", _ssmSubType); err != nil {
			log.Errorf("invalid --sub-type: %s", err.Error())
			return
		}
	}
	if len(_ssmTypeName) > 0 {
		input.TypeName = aws.String(_ssmTypeName)
	}

	if disablePaginator() {
		if resp, err := client.GetInventorySchema(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.GetInventorySchemaOutput
	p := ssm.NewGetInventorySchemaPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves a maintenance window.
func ssm_GetMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetMaintenanceWindowInput{
		// WindowId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}

	if resp, err := client.GetMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific a maintenance window execution.
func ssm_GetMaintenanceWindowExecution(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetMaintenanceWindowExecutionInput{
		// WindowExecutionId: *string, // Required
	}

	if len(_ssmWindowExecutionId) > 0 {
		input.WindowExecutionId = aws.String(_ssmWindowExecutionId)
	}

	if resp, err := client.GetMaintenanceWindowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details about a specific task run as part of a maintenance window
// execution.
func ssm_GetMaintenanceWindowExecutionTask(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetMaintenanceWindowExecutionTaskInput{
		// TaskId: *string, // Required
		// WindowExecutionId: *string, // Required
	}

	if len(_ssmTaskId) > 0 {
		input.TaskId = aws.String(_ssmTaskId)
	}
	if len(_ssmWindowExecutionId) > 0 {
		input.WindowExecutionId = aws.String(_ssmWindowExecutionId)
	}

	if resp, err := client.GetMaintenanceWindowExecutionTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific task running on a specific target.
func ssm_GetMaintenanceWindowExecutionTaskInvocation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetMaintenanceWindowExecutionTaskInvocationInput{
		// InvocationId: *string, // Required
		// TaskId: *string, // Required
		// WindowExecutionId: *string, // Required
	}

	if len(_ssmInvocationId) > 0 {
		input.InvocationId = aws.String(_ssmInvocationId)
	}
	if len(_ssmTaskId) > 0 {
		input.TaskId = aws.String(_ssmTaskId)
	}
	if len(_ssmWindowExecutionId) > 0 {
		input.WindowExecutionId = aws.String(_ssmWindowExecutionId)
	}

	if resp, err := client.GetMaintenanceWindowExecutionTaskInvocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a maintenance window task.
// For maintenance window tasks without a specified target, you can't supply
// values for --max-errors and --max-concurrency . Instead, the system inserts a
// placeholder value of 1 , which may be reported in the response to this command.
// These values don't affect the running of your task and can be ignored.
//
// To retrieve a list of tasks in a maintenance window, instead use the DescribeMaintenanceWindowTasks command.
func ssm_GetMaintenanceWindowTask(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetMaintenanceWindowTaskInput{
		// WindowId: *string, // Required
		// WindowTaskId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmWindowTaskId) > 0 {
		input.WindowTaskId = aws.String(_ssmWindowTaskId)
	}

	if resp, err := client.GetMaintenanceWindowTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about an OpsItem by using the ID. You must have permission in
// Identity and Access Management (IAM) to view information about an OpsItem. For
// more information, see [Set up OpsCenter]in the Amazon Web Services Systems Manager User Guide.
//
// Operations engineers and IT professionals use Amazon Web Services Systems
// Manager OpsCenter to view, investigate, and remediate operational issues
// impacting the performance and health of their Amazon Web Services resources. For
// more information, see [Amazon Web Services Systems Manager OpsCenter]in the Amazon Web Services Systems Manager User Guide.
//
// [Amazon Web Services Systems Manager OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html
// [Set up OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html
func ssm_GetOpsItem(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetOpsItemInput{
		// OpsItemId: *string, // Required
	}

	if len(_ssmOpsItemId) > 0 {
		input.OpsItemId = aws.String(_ssmOpsItemId)
	}
	if len(_ssmOpsItemArn) > 0 {
		input.OpsItemArn = aws.String(_ssmOpsItemArn)
	}

	if resp, err := client.GetOpsItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// View operational metadata related to an application in Application Manager.
func ssm_GetOpsMetadata(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetOpsMetadataInput{
		// OpsMetadataArn: *string, // Required
	}

	if len(_ssmOpsMetadataArn) > 0 {
		input.OpsMetadataArn = aws.String(_ssmOpsMetadataArn)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if resp, err := client.GetOpsMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// View a summary of operations metadata (OpsData) based on specified filters and
// aggregators. OpsData can include information about Amazon Web Services Systems
// Manager OpsCenter operational workitems (OpsItems) as well as information about
// any Amazon Web Services resource or service configured to report OpsData to
// Amazon Web Services Systems Manager Explorer.
func ssm_GetOpsSummary(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetOpsSummaryInput{}

	if len(_ssmAggregators) > 0 {
		if err := assignInputField(input, "Aggregators", _ssmAggregators); err != nil {
			log.Errorf("invalid --aggregators: %s", err.Error())
			return
		}
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmResultAttributes) > 0 {
		if err := assignInputField(input, "ResultAttributes", _ssmResultAttributes); err != nil {
			log.Errorf("invalid --result-attributes: %s", err.Error())
			return
		}
	}
	if len(_ssmSyncName) > 0 {
		input.SyncName = aws.String(_ssmSyncName)
	}

	if disablePaginator() {
		if resp, err := client.GetOpsSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.GetOpsSummaryOutput
	p := ssm.NewGetOpsSummaryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Get information about a single parameter by specifying the parameter name.
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
//
// To get information about more than one parameter at a time, use the GetParameters operation.
func ssm_GetParameter(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetParameterInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmWithDecryption) > 0 {
		if err := assignInputField(input, "WithDecryption", _ssmWithDecryption); err != nil {
			log.Errorf("invalid --with-decryption: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetParameter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the history of all changes to a parameter.
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
//
// If you change the KMS key alias for the KMS key used to encrypt a parameter,
// then you must also update the key alias the parameter uses to reference KMS.
// Otherwise, GetParameterHistory retrieves whatever the original key alias was
// referencing.
func ssm_GetParameterHistory(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetParameterHistoryInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmWithDecryption) > 0 {
		if err := assignInputField(input, "WithDecryption", _ssmWithDecryption); err != nil {
			log.Errorf("invalid --with-decryption: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetParameterHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.GetParameterHistoryOutput
	p := ssm.NewGetParameterHistoryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Get information about one or more parameters by specifying multiple parameter
// names.
//
// To get information about a single parameter, you can use the GetParameter operation instead.
//
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
func ssm_GetParameters(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetParametersInput{
		// Names: []string, // Required
	}

	if len(_ssmNames) > 0 {
		input.Names = append([]string(nil), _ssmNames...)
	}
	if len(_ssmWithDecryption) > 0 {
		if err := assignInputField(input, "WithDecryption", _ssmWithDecryption); err != nil {
			log.Errorf("invalid --with-decryption: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve information about one or more parameters under a specified level in a
// hierarchy.
//
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults . If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken . You can specify the NextToken in a subsequent call to get the
// next set of results.
//
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
func ssm_GetParametersByPath(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetParametersByPathInput{
		// Path: *string, // Required
	}

	if len(_ssmPath) > 0 {
		input.Path = aws.String(_ssmPath)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmParameterFilters) > 0 {
		if err := assignInputField(input, "ParameterFilters", _ssmParameterFilters); err != nil {
			log.Errorf("invalid --parameter-filters: %s", err.Error())
			return
		}
	}
	if len(_ssmRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _ssmRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}
	if len(_ssmWithDecryption) > 0 {
		if err := assignInputField(input, "WithDecryption", _ssmWithDecryption); err != nil {
			log.Errorf("invalid --with-decryption: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetParametersByPath(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.GetParametersByPathOutput
	p := ssm.NewGetParametersByPathPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves information about a patch baseline.
func ssm_GetPatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetPatchBaselineInput{
		// BaselineId: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}

	if resp, err := client.GetPatchBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the patch baseline that should be used for the specified patch group.
func ssm_GetPatchBaselineForPatchGroup(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetPatchBaselineForPatchGroupInput{
		// PatchGroup: *string, // Required
	}

	if len(_ssmPatchGroup) > 0 {
		input.PatchGroup = aws.String(_ssmPatchGroup)
	}
	if len(_ssmOperatingSystem) > 0 {
		if err := assignInputField(input, "OperatingSystem", _ssmOperatingSystem); err != nil {
			log.Errorf("invalid --operating-system: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPatchBaselineForPatchGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of the Policy object.
func ssm_GetResourcePolicies(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetResourcePoliciesInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmResourceArn)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
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

	var results []*ssm.GetResourcePoliciesOutput
	p := ssm.NewGetResourcePoliciesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of false . This means the user can't use
// this feature unless they change the setting to true and intentionally opt in
// for a paid feature.
//
// Services map a SettingId object to a setting value. Amazon Web Services
// services teams define the default value for a SettingId . You can't create a new
// SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the UpdateServiceSetting API operation to
// change the default setting. Or use the ResetServiceSettingto change the value back to the original
// value defined by the Amazon Web Services service team.
//
// Query the current service setting for the Amazon Web Services account.
func ssm_GetServiceSetting(cfg aws.Config, client *ssm.Client) {
	input := &ssm.GetServiceSettingInput{
		// SettingId: *string, // Required
	}

	if len(_ssmSettingId) > 0 {
		input.SettingId = aws.String(_ssmSettingId)
	}

	if resp, err := client.GetServiceSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A parameter label is a user-defined alias to help you manage different versions
// of a parameter. When you modify a parameter, Amazon Web Services Systems Manager
// automatically saves a new version and increments the version number by one. A
// label can help you remember the purpose of a parameter when there are multiple
// versions.
//
// Parameter labels have the following requirements and restrictions.
//
// - A version of a parameter can have a maximum of 10 labels.
//
// - You can't attach the same label to different versions of the same
// parameter. For example, if version 1 has the label Production, then you can't
// attach Production to version 2.
//
// - You can move a label from one version of a parameter to another.
//
// - You can't create a label when you create a new parameter. You must attach a
// label to a specific version of a parameter.
//
// - If you no longer want to use a parameter label, then you can either delete
// it or move it to a different version of a parameter.
//
// - A label can have a maximum of 100 characters.
//
// - Labels can contain letters (case sensitive), numbers, periods (.), hyphens
// (-), or underscores (_).
//
// - Labels can't begin with a number, " aws " or " ssm " (not case sensitive).
// If a label fails to meet these requirements, then the label isn't associated
// with a parameter and the system displays it in the list of InvalidLabels.
//
// - Parameter names can't contain spaces. The service removes any spaces
// specified for the beginning or end of a parameter name. If the specified name
// for a parameter contains spaces between characters, the request fails with a
// ValidationException error.
func ssm_LabelParameterVersion(cfg aws.Config, client *ssm.Client) {
	input := &ssm.LabelParameterVersionInput{
		// Labels: []string, // Required
		// Name: *string, // Required
	}

	if len(_ssmLabels) > 0 {
		input.Labels = append([]string(nil), _ssmLabels...)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmParameterVersion) > 0 {
		if err := assignInputField(input, "ParameterVersion", _ssmParameterVersion); err != nil {
			log.Errorf("invalid --parameter-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.LabelParameterVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all versions of an association for a specific association ID.
func ssm_ListAssociationVersions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListAssociationVersionsInput{
		// AssociationId: *string, // Required
	}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociationVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListAssociationVersionsOutput
	p := ssm.NewListAssociationVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns all State Manager associations in the current Amazon Web Services
// account and Amazon Web Services Region. You can limit the results to a specific
// State Manager association document or managed node by specifying a filter. State
// Manager is a tool in Amazon Web Services Systems Manager.
func ssm_ListAssociations(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListAssociationsInput{}

	if len(_ssmAssociationFilterList) > 0 {
		if err := assignInputField(input, "AssociationFilterList", _ssmAssociationFilterList); err != nil {
			log.Errorf("invalid --association-filter-list: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListAssociationsOutput
	p := ssm.NewListAssociationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// An invocation is copy of a command sent to a specific managed node. A command
// can apply to one or more managed nodes. A command invocation applies to one
// managed node. For example, if a user runs SendCommand against three managed
// nodes, then a command invocation is created for each requested managed node ID.
// ListCommandInvocations provide status about command execution.
func ssm_ListCommandInvocations(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListCommandInvocationsInput{}

	if len(_ssmCommandId) > 0 {
		input.CommandId = aws.String(_ssmCommandId)
	}
	if len(_ssmDetails) > 0 {
		if err := assignInputField(input, "Details", _ssmDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCommandInvocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListCommandInvocationsOutput
	p := ssm.NewListCommandInvocationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the commands requested by users of the Amazon Web Services account.
func ssm_ListCommands(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListCommandsInput{}

	if len(_ssmCommandId) > 0 {
		input.CommandId = aws.String(_ssmCommandId)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCommands(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListCommandsOutput
	p := ssm.NewListCommandsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// For a specified resource ID, this API operation returns a list of compliance
// statuses for different resource types. Currently, you can only specify one
// resource ID per call. List results depend on the criteria specified in the
// filter.
func ssm_ListComplianceItems(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListComplianceItemsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmResourceIds) > 0 {
		input.ResourceIds = append([]string(nil), _ssmResourceIds...)
	}
	if len(_ssmResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _ssmResourceTypes...)
	}

	if disablePaginator() {
		if resp, err := client.ListComplianceItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListComplianceItemsOutput
	p := ssm.NewListComplianceItemsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a summary count of compliant and non-compliant resources for a
// compliance type. For example, this call can return State Manager associations,
// patches, or custom compliance types according to the filter criteria that you
// specify.
func ssm_ListComplianceSummaries(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListComplianceSummariesInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComplianceSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListComplianceSummariesOutput
	p := ssm.NewListComplianceSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Amazon Web Services Systems Manager Change Manager is no longer open to new
// customers. Existing customers can continue to use the service as normal. For
// more information, see [Amazon Web Services Systems Manager Change Manager availability change].
//
// Information about approval reviews for a version of a change template in Change
// Manager.
//
// [Amazon Web Services Systems Manager Change Manager availability change]: https://docs.aws.amazon.com/systems-manager/latest/userguide/change-manager-availability-change.html
func ssm_ListDocumentMetadataHistory(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListDocumentMetadataHistoryInput{
		// Metadata: types.DocumentMetadataEnum, // Required
		// Name: *string, // Required
	}

	if len(_ssmMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _ssmMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if resp, err := client.ListDocumentMetadataHistory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all versions for a document.
func ssm_ListDocumentVersions(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListDocumentVersionsInput{
		// Name: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDocumentVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListDocumentVersionsOutput
	p := ssm.NewListDocumentVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns all Systems Manager (SSM) documents in the current Amazon Web Services
// account and Amazon Web Services Region. You can limit the results of this
// request by using a filter.
func ssm_ListDocuments(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListDocumentsInput{}

	if len(_ssmDocumentFilterList) > 0 {
		if err := assignInputField(input, "DocumentFilterList", _ssmDocumentFilterList); err != nil {
			log.Errorf("invalid --document-filter-list: %s", err.Error())
			return
		}
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDocuments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListDocumentsOutput
	p := ssm.NewListDocumentsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// A list of inventory items returned by the request.
func ssm_ListInventoryEntries(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListInventoryEntriesInput{
		// InstanceId: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmTypeName) > 0 {
		input.TypeName = aws.String(_ssmTypeName)
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if resp, err := client.ListInventoryEntries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Takes in filters and returns a list of managed nodes matching the filter
// criteria.
func ssm_ListNodes(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListNodesInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmSyncName) > 0 {
		input.SyncName = aws.String(_ssmSyncName)
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

	var results []*ssm.ListNodesOutput
	p := ssm.NewListNodesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Generates a summary of managed instance/node metadata based on the filters and
// aggregators you specify. Results are grouped by the input aggregator you
// specify.
func ssm_ListNodesSummary(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListNodesSummaryInput{
		// Aggregators: []types.NodeAggregator, // Required
	}

	if len(_ssmAggregators) > 0 {
		if err := assignInputField(input, "Aggregators", _ssmAggregators); err != nil {
			log.Errorf("invalid --aggregators: %s", err.Error())
			return
		}
	}
	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmSyncName) > 0 {
		input.SyncName = aws.String(_ssmSyncName)
	}

	if disablePaginator() {
		if resp, err := client.ListNodesSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListNodesSummaryOutput
	p := ssm.NewListNodesSummaryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all OpsItem events in the current Amazon Web Services Region
// and Amazon Web Services account. You can limit the results to events associated
// with specific OpsItems by specifying a filter.
func ssm_ListOpsItemEvents(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListOpsItemEventsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOpsItemEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListOpsItemEventsOutput
	p := ssm.NewListOpsItemEventsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all related-item resources associated with a Systems Manager OpsCenter
// OpsItem. OpsCenter is a tool in Amazon Web Services Systems Manager.
func ssm_ListOpsItemRelatedItems(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListOpsItemRelatedItemsInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmOpsItemId) > 0 {
		input.OpsItemId = aws.String(_ssmOpsItemId)
	}

	if disablePaginator() {
		if resp, err := client.ListOpsItemRelatedItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListOpsItemRelatedItemsOutput
	p := ssm.NewListOpsItemRelatedItemsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Amazon Web Services Systems Manager calls this API operation when displaying
// all Application Manager OpsMetadata objects or blobs.
func ssm_ListOpsMetadata(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListOpsMetadataInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOpsMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListOpsMetadataOutput
	p := ssm.NewListOpsMetadataPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a resource-level summary count. The summary includes information about
// compliant and non-compliant statuses and detailed compliance-item severity
// counts, according to the filter criteria you specify.
func ssm_ListResourceComplianceSummaries(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListResourceComplianceSummariesInput{}

	if len(_ssmFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceComplianceSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListResourceComplianceSummariesOutput
	p := ssm.NewListResourceComplianceSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists your resource data sync configurations. Includes information about the
// last time a sync attempted to start, the last sync status, and the last time a
// sync successfully completed.
//
// The number of sync configurations might be too large to return using a single
// call to ListResourceDataSync . You can limit the number of sync configurations
// returned by using the MaxResults parameter. To determine whether there are more
// sync configurations to list, check the value of NextToken in the output. If
// there are more sync configurations to list, you can request them by specifying
// the NextToken returned in the call to the parameter of a subsequent call.
func ssm_ListResourceDataSync(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListResourceDataSyncInput{}

	if len(_ssmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmNextToken) > 0 {
		input.NextToken = aws.String(_ssmNextToken)
	}
	if len(_ssmSyncType) > 0 {
		input.SyncType = aws.String(_ssmSyncType)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceDataSync(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssm.ListResourceDataSyncOutput
	p := ssm.NewListResourceDataSyncPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of the tags assigned to the specified resource.
// For information about the ID format for each supported resource type, see AddTagsToResource.
func ssm_ListTagsForResource(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ListTagsForResourceInput{
		// ResourceId: *string, // Required
		// ResourceType: types.ResourceTypeForTagging, // Required
	}

	if len(_ssmResourceId) > 0 {
		input.ResourceId = aws.String(_ssmResourceId)
	}
	if len(_ssmResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _ssmResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
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

// Shares a Amazon Web Services Systems Manager document (SSM document)publicly or
// privately. If you share a document privately, you must specify the Amazon Web
// Services user IDs for those people who can use the document. If you share a
// document publicly, you must specify All as the account ID.
func ssm_ModifyDocumentPermission(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ModifyDocumentPermissionInput{
		// Name: *string, // Required
		// PermissionType: types.DocumentPermissionType, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _ssmPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_ssmAccountIdsToAdd) > 0 {
		input.AccountIdsToAdd = append([]string(nil), _ssmAccountIdsToAdd...)
	}
	if len(_ssmAccountIdsToRemove) > 0 {
		input.AccountIdsToRemove = append([]string(nil), _ssmAccountIdsToRemove...)
	}
	if len(_ssmSharedDocumentVersion) > 0 {
		input.SharedDocumentVersion = aws.String(_ssmSharedDocumentVersion)
	}

	if resp, err := client.ModifyDocumentPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a compliance type and other compliance details on a designated
// resource. This operation lets you register custom compliance details with a
// resource. This call overwrites existing compliance information on the resource,
// so you must provide a full list of compliance items each time that you send the
// request.
//
// ComplianceType can be one of the following:
//
// - ExecutionId: The execution ID when the patch, association, or custom
// compliance item was applied.
//
// - ExecutionType: Specify patch, association, or Custom: string .
//
// - ExecutionTime. The time the patch, association, or custom compliance item
// was applied to the managed node.
//
// # For State Manager associations, this represents the time when compliance status
//
// was captured by the Systems Manager service during its internal compliance
// aggregation workflow, not necessarily when the association was executed on the
// managed node. State Manager updates compliance information for all associations
// on an instance whenever any association executes, which may result in multiple
// associations showing the same execution time.
//
// - Id: The patch, association, or custom compliance ID.
//
// - Title: A title.
//
// - Status: The status of the compliance item. For example, approved for
// patches, or Failed for associations.
//
// - Severity: A patch severity. For example, Critical .
//
// - DocumentName: An SSM document name. For example, AWS-RunPatchBaseline .
//
// - DocumentVersion: An SSM document version number. For example, 4.
//
// - Classification: A patch classification. For example, security updates .
//
// - PatchBaselineId: A patch baseline ID.
//
// - PatchSeverity: A patch severity. For example, Critical .
//
// - PatchState: A patch state. For example, InstancesWithFailedPatches .
//
// - PatchGroup: The name of a patch group.
//
// - InstalledTime: The time the association, patch, or custom compliance item
// was applied to the resource. Specify the time by using the following format:
// yyyy-MM-dd'T'HH:mm:ss'Z'
func ssm_PutComplianceItems(cfg aws.Config, client *ssm.Client) {
	input := &ssm.PutComplianceItemsInput{
		// ComplianceType: *string, // Required
		// ExecutionSummary: *types.ComplianceExecutionSummary, // Required
		// Items: []types.ComplianceItemEntry, // Required
		// ResourceId: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_ssmComplianceType) > 0 {
		input.ComplianceType = aws.String(_ssmComplianceType)
	}
	if len(_ssmExecutionSummary) > 0 {
		if err := assignInputField(input, "ExecutionSummary", _ssmExecutionSummary); err != nil {
			log.Errorf("invalid --execution-summary: %s", err.Error())
			return
		}
	}
	if len(_ssmItems) > 0 {
		if err := assignInputField(input, "Items", _ssmItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}
	if len(_ssmResourceId) > 0 {
		input.ResourceId = aws.String(_ssmResourceId)
	}
	if len(_ssmResourceType) > 0 {
		input.ResourceType = aws.String(_ssmResourceType)
	}
	if len(_ssmItemContentHash) > 0 {
		input.ItemContentHash = aws.String(_ssmItemContentHash)
	}
	if len(_ssmUploadType) > 0 {
		if err := assignInputField(input, "UploadType", _ssmUploadType); err != nil {
			log.Errorf("invalid --upload-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutComplianceItems(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Bulk update custom inventory items on one or more managed nodes. The request
// adds an inventory item, if it doesn't already exist, or updates an inventory
// item, if it does exist.
func ssm_PutInventory(cfg aws.Config, client *ssm.Client) {
	input := &ssm.PutInventoryInput{
		// InstanceId: *string, // Required
		// Items: []types.InventoryItem, // Required
	}

	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmItems) > 0 {
		if err := assignInputField(input, "Items", _ssmItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutInventory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create or update a parameter in Parameter Store.
func ssm_PutParameter(cfg aws.Config, client *ssm.Client) {
	input := &ssm.PutParameterInput{
		// Name: *string, // Required
		// Value: *string, // Required
	}

	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmValue) > 0 {
		input.Value = aws.String(_ssmValue)
	}
	if len(_ssmAllowedPattern) > 0 {
		input.AllowedPattern = aws.String(_ssmAllowedPattern)
	}
	if len(_ssmDataType) > 0 {
		input.DataType = aws.String(_ssmDataType)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmKeyId) > 0 {
		input.KeyId = aws.String(_ssmKeyId)
	}
	if len(_ssmOverwrite) > 0 {
		if err := assignInputField(input, "Overwrite", _ssmOverwrite); err != nil {
			log.Errorf("invalid --overwrite: %s", err.Error())
			return
		}
	}
	if len(_ssmPolicies) > 0 {
		input.Policies = aws.String(_ssmPolicies)
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ssmTier) > 0 {
		if err := assignInputField(input, "Tier", _ssmTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_ssmType) > 0 {
		if err := assignInputField(input, "Type", _ssmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutParameter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a Systems Manager resource policy. A resource policy helps
// you to define the IAM entity (for example, an Amazon Web Services account) that
// can manage your Systems Manager resources. The following resources support
// Systems Manager resource policies.
//
// - OpsItemGroup - The resource policy for OpsItemGroup enables Amazon Web
// Services accounts to view and interact with OpsCenter operational work items
// (OpsItems).
//
// - Parameter - The resource policy is used to share a parameter with other
// accounts using Resource Access Manager (RAM).
//
// To share a parameter, it must be in the advanced parameter tier. For
//
// information about parameter tiers, see [Managing parameter tiers]. For information about changing an
// existing standard parameter to an advanced parameter, see [Changing a standard parameter to an advanced parameter].
//
// # To share a SecureString parameter, it must be encrypted with a customer managed
//
// key, and you must share the key separately through Key Management Service.
// Amazon Web Services managed keys cannot be shared. Parameters encrypted with the
// default Amazon Web Services managed key can be updated to use a customer managed
// key instead. For KMS key definitions, see [KMS concepts]in the Key Management Service
// Developer Guide.
//
// # While you can share a parameter using the Systems Manager PutResourcePolicy
//
// operation, we recommend using Resource Access Manager (RAM) instead. This is
// because using PutResourcePolicy requires the extra step of promoting the
// parameter to a standard RAM Resource Share using the RAM [PromoteResourceShareCreatedFromPolicy]API operation.
// Otherwise, the parameter won't be returned by the Systems Manager [DescribeParameters]API
// operation using the --shared option.
//
// For more information, see [Sharing a parameter]in the Amazon Web Services Systems Manager User Guide
//
// [Sharing a parameter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-shared-parameters.html#share
// [Managing parameter tiers]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html
// [Changing a standard parameter to an advanced parameter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html#parameter-store-advanced-parameters-enabling
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [KMS concepts]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html
// [DescribeParameters]: https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribeParameters.html
func ssm_PutResourcePolicy(cfg aws.Config, client *ssm.Client) {
	input := &ssm.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_ssmPolicy) > 0 {
		input.Policy = aws.String(_ssmPolicy)
	}
	if len(_ssmResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmResourceArn)
	}
	if len(_ssmPolicyHash) > 0 {
		input.PolicyHash = aws.String(_ssmPolicyHash)
	}
	if len(_ssmPolicyId) > 0 {
		input.PolicyId = aws.String(_ssmPolicyId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the default patch baseline for the relevant operating system.
// To reset the Amazon Web Services-predefined patch baseline as the default,
// specify the full patch baseline Amazon Resource Name (ARN) as the baseline ID
// value. For example, for CentOS, specify
// arn:aws:ssm:us-east-2:733109147000:patchbaseline/pb-0574b43a65ea646ed instead of
// pb-0574b43a65ea646ed .
func ssm_RegisterDefaultPatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.RegisterDefaultPatchBaselineInput{
		// BaselineId: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}

	if resp, err := client.RegisterDefaultPatchBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a patch baseline for a patch group.
func ssm_RegisterPatchBaselineForPatchGroup(cfg aws.Config, client *ssm.Client) {
	input := &ssm.RegisterPatchBaselineForPatchGroupInput{
		// BaselineId: *string, // Required
		// PatchGroup: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}
	if len(_ssmPatchGroup) > 0 {
		input.PatchGroup = aws.String(_ssmPatchGroup)
	}

	if resp, err := client.RegisterPatchBaselineForPatchGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a target with a maintenance window.
func ssm_RegisterTargetWithMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.RegisterTargetWithMaintenanceWindowInput{
		// ResourceType: types.MaintenanceWindowResourceType, // Required
		// Targets: []types.Target, // Required
		// WindowId: *string, // Required
	}

	if len(_ssmResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _ssmResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmOwnerInformation) > 0 {
		input.OwnerInformation = aws.String(_ssmOwnerInformation)
	}

	if resp, err := client.RegisterTargetWithMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new task to a maintenance window.
func ssm_RegisterTaskWithMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.RegisterTaskWithMaintenanceWindowInput{
		// TaskArn: *string, // Required
		// TaskType: types.MaintenanceWindowTaskType, // Required
		// WindowId: *string, // Required
	}

	if len(_ssmTaskArn) > 0 {
		input.TaskArn = aws.String(_ssmTaskArn)
	}
	if len(_ssmTaskType) > 0 {
		if err := assignInputField(input, "TaskType", _ssmTaskType); err != nil {
			log.Errorf("invalid --task-type: %s", err.Error())
			return
		}
	}
	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _ssmAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmCutoffBehavior) > 0 {
		if err := assignInputField(input, "CutoffBehavior", _ssmCutoffBehavior); err != nil {
			log.Errorf("invalid --cutoff-behavior: %s", err.Error())
			return
		}
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmLoggingInfo) > 0 {
		if err := assignInputField(input, "LoggingInfo", _ssmLoggingInfo); err != nil {
			log.Errorf("invalid --logging-info: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxConcurrency) > 0 {
		input.MaxConcurrency = aws.String(_ssmMaxConcurrency)
	}
	if len(_ssmMaxErrors) > 0 {
		input.MaxErrors = aws.String(_ssmMaxErrors)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmPriority) > 0 {
		if err := assignInputField(input, "Priority", _ssmPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_ssmServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_ssmServiceRoleArn)
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmTaskInvocationParameters) > 0 {
		if err := assignInputField(input, "TaskInvocationParameters", _ssmTaskInvocationParameters); err != nil {
			log.Errorf("invalid --task-invocation-parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmTaskParameters) > 0 {
		if err := assignInputField(input, "TaskParameters", _ssmTaskParameters); err != nil {
			log.Errorf("invalid --task-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterTaskWithMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tag keys from the specified resource.
func ssm_RemoveTagsFromResource(cfg aws.Config, client *ssm.Client) {
	input := &ssm.RemoveTagsFromResourceInput{
		// ResourceId: *string, // Required
		// ResourceType: types.ResourceTypeForTagging, // Required
		// TagKeys: []string, // Required
	}

	if len(_ssmResourceId) > 0 {
		input.ResourceId = aws.String(_ssmResourceId)
	}
	if len(_ssmResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _ssmResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_ssmTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ssmTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of "false". This means the user can't use
// this feature unless they change the setting to "true" and intentionally opt in
// for a paid feature.
//
// Services map a SettingId object to a setting value. Amazon Web Services
// services teams define the default value for a SettingId . You can't create a new
// SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the GetServiceSetting API operation to
// view the current value. Use the UpdateServiceSettingAPI operation to change the default setting.
//
// Reset the service setting for the account to the default value as provisioned
// by the Amazon Web Services service team.
func ssm_ResetServiceSetting(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ResetServiceSettingInput{
		// SettingId: *string, // Required
	}

	if len(_ssmSettingId) > 0 {
		input.SettingId = aws.String(_ssmSettingId)
	}

	if resp, err := client.ResetServiceSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reconnects a session to a managed node after it has been disconnected.
// Connections can be resumed for disconnected sessions, but not terminated
// sessions.
//
// This command is primarily for use by client machines to automatically reconnect
// during intermittent network issues. It isn't intended for any other use.
func ssm_ResumeSession(cfg aws.Config, client *ssm.Client) {
	input := &ssm.ResumeSessionInput{
		// SessionId: *string, // Required
	}

	if len(_ssmSessionId) > 0 {
		input.SessionId = aws.String(_ssmSessionId)
	}

	if resp, err := client.ResumeSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a signal to an Automation execution to change the current behavior or
// status of the execution.
func ssm_SendAutomationSignal(cfg aws.Config, client *ssm.Client) {
	input := &ssm.SendAutomationSignalInput{
		// AutomationExecutionId: *string, // Required
		// SignalType: types.SignalType, // Required
	}

	if len(_ssmAutomationExecutionId) > 0 {
		input.AutomationExecutionId = aws.String(_ssmAutomationExecutionId)
	}
	if len(_ssmSignalType) > 0 {
		if err := assignInputField(input, "SignalType", _ssmSignalType); err != nil {
			log.Errorf("invalid --signal-type: %s", err.Error())
			return
		}
	}
	if len(_ssmPayload) > 0 {
		if err := assignInputField(input, "Payload", _ssmPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendAutomationSignal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs commands on one or more managed nodes.
func ssm_SendCommand(cfg aws.Config, client *ssm.Client) {
	input := &ssm.SendCommandInput{
		// DocumentName: *string, // Required
	}

	if len(_ssmDocumentName) > 0 {
		input.DocumentName = aws.String(_ssmDocumentName)
	}
	if len(_ssmAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _ssmAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssmCloudWatchOutputConfig) > 0 {
		if err := assignInputField(input, "CloudWatchOutputConfig", _ssmCloudWatchOutputConfig); err != nil {
			log.Errorf("invalid --cloud-watch-output-config: %s", err.Error())
			return
		}
	}
	if len(_ssmComment) > 0 {
		input.Comment = aws.String(_ssmComment)
	}
	if len(_ssmDocumentHash) > 0 {
		input.DocumentHash = aws.String(_ssmDocumentHash)
	}
	if len(_ssmDocumentHashType) > 0 {
		if err := assignInputField(input, "DocumentHashType", _ssmDocumentHashType); err != nil {
			log.Errorf("invalid --document-hash-type: %s", err.Error())
			return
		}
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _ssmInstanceIds...)
	}
	if len(_ssmMaxConcurrency) > 0 {
		input.MaxConcurrency = aws.String(_ssmMaxConcurrency)
	}
	if len(_ssmMaxErrors) > 0 {
		input.MaxErrors = aws.String(_ssmMaxErrors)
	}
	if len(_ssmNotificationConfig) > 0 {
		if err := assignInputField(input, "NotificationConfig", _ssmNotificationConfig); err != nil {
			log.Errorf("invalid --notification-config: %s", err.Error())
			return
		}
	}
	if len(_ssmOutputS3BucketName) > 0 {
		input.OutputS3BucketName = aws.String(_ssmOutputS3BucketName)
	}
	if len(_ssmOutputS3KeyPrefix) > 0 {
		input.OutputS3KeyPrefix = aws.String(_ssmOutputS3KeyPrefix)
	}
	if len(_ssmOutputS3Region) > 0 {
		input.OutputS3Region = aws.String(_ssmOutputS3Region)
	}
	if len(_ssmParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_ssmServiceRoleArn)
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmTimeoutSeconds) > 0 {
		if err := assignInputField(input, "TimeoutSeconds", _ssmTimeoutSeconds); err != nil {
			log.Errorf("invalid --timeout-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the workflow for just-in-time node access sessions.
func ssm_StartAccessRequest(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StartAccessRequestInput{
		// Reason: *string, // Required
		// Targets: []types.Target, // Required
	}

	if len(_ssmReason) > 0 {
		input.Reason = aws.String(_ssmReason)
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAccessRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs an association immediately and only one time. This operation can be
// helpful when troubleshooting associations.
func ssm_StartAssociationsOnce(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StartAssociationsOnceInput{
		// AssociationIds: []string, // Required
	}

	if len(_ssmAssociationIds) > 0 {
		input.AssociationIds = append([]string(nil), _ssmAssociationIds...)
	}

	if resp, err := client.StartAssociationsOnce(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates execution of an Automation runbook.
func ssm_StartAutomationExecution(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StartAutomationExecutionInput{
		// DocumentName: *string, // Required
	}

	if len(_ssmDocumentName) > 0 {
		input.DocumentName = aws.String(_ssmDocumentName)
	}
	if len(_ssmAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _ssmAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmMaxConcurrency) > 0 {
		input.MaxConcurrency = aws.String(_ssmMaxConcurrency)
	}
	if len(_ssmMaxErrors) > 0 {
		input.MaxErrors = aws.String(_ssmMaxErrors)
	}
	if len(_ssmMode) > 0 {
		if err := assignInputField(input, "Mode", _ssmMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_ssmParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetLocations) > 0 {
		if err := assignInputField(input, "TargetLocations", _ssmTargetLocations); err != nil {
			log.Errorf("invalid --target-locations: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetLocationsURL) > 0 {
		input.TargetLocationsURL = aws.String(_ssmTargetLocationsURL)
	}
	if len(_ssmTargetMaps) > 0 {
		if err := assignInputField(input, "TargetMaps", _ssmTargetMaps); err != nil {
			log.Errorf("invalid --target-maps: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetParameterName) > 0 {
		input.TargetParameterName = aws.String(_ssmTargetParameterName)
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAutomationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services Systems Manager Change Manager is no longer open to new
// customers. Existing customers can continue to use the service as normal. For
// more information, see [Amazon Web Services Systems Manager Change Manager availability change].
//
// Creates a change request for Change Manager. The Automation runbooks specified
// in the change request run only after all required approvals for the change
// request have been received.
//
// [Amazon Web Services Systems Manager Change Manager availability change]: https://docs.aws.amazon.com/systems-manager/latest/userguide/change-manager-availability-change.html
func ssm_StartChangeRequestExecution(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StartChangeRequestExecutionInput{
		// DocumentName: *string, // Required
		// Runbooks: []types.Runbook, // Required
	}

	if len(_ssmDocumentName) > 0 {
		input.DocumentName = aws.String(_ssmDocumentName)
	}
	if len(_ssmRunbooks) > 0 {
		if err := assignInputField(input, "Runbooks", _ssmRunbooks); err != nil {
			log.Errorf("invalid --runbooks: %s", err.Error())
			return
		}
	}
	if len(_ssmAutoApprove) > 0 {
		if err := assignInputField(input, "AutoApprove", _ssmAutoApprove); err != nil {
			log.Errorf("invalid --auto-approve: %s", err.Error())
			return
		}
	}
	if len(_ssmChangeDetails) > 0 {
		input.ChangeDetails = aws.String(_ssmChangeDetails)
	}
	if len(_ssmChangeRequestName) > 0 {
		input.ChangeRequestName = aws.String(_ssmChangeRequestName)
	}
	if len(_ssmClientToken) > 0 {
		input.ClientToken = aws.String(_ssmClientToken)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmScheduledEndTime) > 0 {
		if err := assignInputField(input, "ScheduledEndTime", _ssmScheduledEndTime); err != nil {
			log.Errorf("invalid --scheduled-end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmScheduledTime) > 0 {
		if err := assignInputField(input, "ScheduledTime", _ssmScheduledTime); err != nil {
			log.Errorf("invalid --scheduled-time: %s", err.Error())
			return
		}
	}
	if len(_ssmTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartChangeRequestExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the process of creating a preview showing the effects that running a
// specified Automation runbook would have on the targeted resources.
func ssm_StartExecutionPreview(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StartExecutionPreviewInput{
		// DocumentName: *string, // Required
	}

	if len(_ssmDocumentName) > 0 {
		input.DocumentName = aws.String(_ssmDocumentName)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmExecutionInputs) > 0 {
		if err := assignInputField(input, "ExecutionInputs", _ssmExecutionInputs); err != nil {
			log.Errorf("invalid --execution-inputs: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartExecutionPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a connection to a target (for example, a managed node) for a Session
// Manager session. Returns a URL and token that can be used to open a WebSocket
// connection for sending input and receiving outputs.
//
// Amazon Web Services CLI usage: start-session is an interactive command that
// requires the Session Manager plugin to be installed on the client machine making
// the call. For information, see [Install the Session Manager plugin for the Amazon Web Services CLI]in the Amazon Web Services Systems Manager User
// Guide.
//
// Amazon Web Services Tools for PowerShell usage: Start-SSMSession isn't
// currently supported by Amazon Web Services Tools for PowerShell on Windows local
// machines.
//
// [Install the Session Manager plugin for the Amazon Web Services CLI]: https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html
func ssm_StartSession(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StartSessionInput{
		// Target: *string, // Required
	}

	if len(_ssmTarget) > 0 {
		input.Target = aws.String(_ssmTarget)
	}
	if len(_ssmDocumentName) > 0 {
		input.DocumentName = aws.String(_ssmDocumentName)
	}
	if len(_ssmParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmReason) > 0 {
		input.Reason = aws.String(_ssmReason)
	}

	if resp, err := client.StartSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop an Automation that is currently running.
func ssm_StopAutomationExecution(cfg aws.Config, client *ssm.Client) {
	input := &ssm.StopAutomationExecutionInput{
		// AutomationExecutionId: *string, // Required
	}

	if len(_ssmAutomationExecutionId) > 0 {
		input.AutomationExecutionId = aws.String(_ssmAutomationExecutionId)
	}
	if len(_ssmType) > 0 {
		if err := assignInputField(input, "Type", _ssmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopAutomationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently ends a session and closes the data connection between the Session
// Manager client and SSM Agent on the managed node. A terminated session can't be
// resumed.
func ssm_TerminateSession(cfg aws.Config, client *ssm.Client) {
	input := &ssm.TerminateSessionInput{
		// SessionId: *string, // Required
	}

	if len(_ssmSessionId) > 0 {
		input.SessionId = aws.String(_ssmSessionId)
	}

	if resp, err := client.TerminateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a label or labels from a parameter.
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
func ssm_UnlabelParameterVersion(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UnlabelParameterVersionInput{
		// Labels: []string, // Required
		// Name: *string, // Required
		// ParameterVersion: *int64, // Required
	}

	if len(_ssmLabels) > 0 {
		input.Labels = append([]string(nil), _ssmLabels...)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmParameterVersion) > 0 {
		if err := assignInputField(input, "ParameterVersion", _ssmParameterVersion); err != nil {
			log.Errorf("invalid --parameter-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UnlabelParameterVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an association. You can update the association name and version, the
// document version, schedule, parameters, and Amazon Simple Storage Service
// (Amazon S3) output. When you call UpdateAssociation , the system removes all
// optional parameters from the request and overwrites the association with null
// values for those parameters. This is by design. You must specify all optional
// parameters in the call, even if you are not changing the parameters. This
// includes the Name parameter. Before calling this API action, we recommend that
// you call the DescribeAssociationAPI operation and make a note of all optional parameters required
// for your UpdateAssociation call.
//
// In order to call this API operation, a user, group, or role must be granted
// permission to call the DescribeAssociationAPI operation. If you don't have permission to call
// DescribeAssociation , then you receive the following error: An error occurred
// (AccessDeniedException) when calling the UpdateAssociation operation: User:
// isn't authorized to perform: ssm:DescribeAssociation on resource:
//
// When you update an association, the association immediately runs against the
// specified targets. You can add the ApplyOnlyAtCronInterval parameter to run the
// association during the next schedule run.
func ssm_UpdateAssociation(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateAssociationInput{
		// AssociationId: *string, // Required
	}

	if len(_ssmAssociationId) > 0 {
		input.AssociationId = aws.String(_ssmAssociationId)
	}
	if len(_ssmAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _ssmAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssmApplyOnlyAtCronInterval) > 0 {
		if err := assignInputField(input, "ApplyOnlyAtCronInterval", _ssmApplyOnlyAtCronInterval); err != nil {
			log.Errorf("invalid --apply-only-at-cron-interval: %s", err.Error())
			return
		}
	}
	if len(_ssmAssociationDispatchAssumeRole) > 0 {
		input.AssociationDispatchAssumeRole = aws.String(_ssmAssociationDispatchAssumeRole)
	}
	if len(_ssmAssociationName) > 0 {
		input.AssociationName = aws.String(_ssmAssociationName)
	}
	if len(_ssmAssociationVersion) > 0 {
		input.AssociationVersion = aws.String(_ssmAssociationVersion)
	}
	if len(_ssmAutomationTargetParameterName) > 0 {
		input.AutomationTargetParameterName = aws.String(_ssmAutomationTargetParameterName)
	}
	if len(_ssmCalendarNames) > 0 {
		input.CalendarNames = append([]string(nil), _ssmCalendarNames...)
	}
	if len(_ssmComplianceSeverity) > 0 {
		if err := assignInputField(input, "ComplianceSeverity", _ssmComplianceSeverity); err != nil {
			log.Errorf("invalid --compliance-severity: %s", err.Error())
			return
		}
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmDuration) > 0 {
		if err := assignInputField(input, "Duration", _ssmDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxConcurrency) > 0 {
		input.MaxConcurrency = aws.String(_ssmMaxConcurrency)
	}
	if len(_ssmMaxErrors) > 0 {
		input.MaxErrors = aws.String(_ssmMaxErrors)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmOutputLocation) > 0 {
		if err := assignInputField(input, "OutputLocation", _ssmOutputLocation); err != nil {
			log.Errorf("invalid --output-location: %s", err.Error())
			return
		}
	}
	if len(_ssmParameters) > 0 {
		if err := assignInputField(input, "Parameters", _ssmParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_ssmScheduleExpression)
	}
	if len(_ssmScheduleOffset) > 0 {
		if err := assignInputField(input, "ScheduleOffset", _ssmScheduleOffset); err != nil {
			log.Errorf("invalid --schedule-offset: %s", err.Error())
			return
		}
	}
	if len(_ssmSyncCompliance) > 0 {
		if err := assignInputField(input, "SyncCompliance", _ssmSyncCompliance); err != nil {
			log.Errorf("invalid --sync-compliance: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetLocations) > 0 {
		if err := assignInputField(input, "TargetLocations", _ssmTargetLocations); err != nil {
			log.Errorf("invalid --target-locations: %s", err.Error())
			return
		}
	}
	if len(_ssmTargetMaps) > 0 {
		if err := assignInputField(input, "TargetMaps", _ssmTargetMaps); err != nil {
			log.Errorf("invalid --target-maps: %s", err.Error())
			return
		}
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of the Amazon Web Services Systems Manager document (SSM
// document) associated with the specified managed node.
//
// UpdateAssociationStatus is primarily used by the Amazon Web Services Systems
// Manager Agent (SSM Agent) to report status updates about your associations and
// is only used for associations created with the InstanceId legacy parameter.
func ssm_UpdateAssociationStatus(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateAssociationStatusInput{
		// AssociationStatus: *types.AssociationStatus, // Required
		// InstanceId: *string, // Required
		// Name: *string, // Required
	}

	if len(_ssmAssociationStatus) > 0 {
		if err := assignInputField(input, "AssociationStatus", _ssmAssociationStatus); err != nil {
			log.Errorf("invalid --association-status: %s", err.Error())
			return
		}
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}

	if resp, err := client.UpdateAssociationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates one or more values for an SSM document.
func ssm_UpdateDocument(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateDocumentInput{
		// Content: *string, // Required
		// Name: *string, // Required
	}

	if len(_ssmContent) > 0 {
		input.Content = aws.String(_ssmContent)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _ssmAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_ssmDisplayName) > 0 {
		input.DisplayName = aws.String(_ssmDisplayName)
	}
	if len(_ssmDocumentFormat) > 0 {
		if err := assignInputField(input, "DocumentFormat", _ssmDocumentFormat); err != nil {
			log.Errorf("invalid --document-format: %s", err.Error())
			return
		}
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmTargetType) > 0 {
		input.TargetType = aws.String(_ssmTargetType)
	}
	if len(_ssmVersionName) > 0 {
		input.VersionName = aws.String(_ssmVersionName)
	}

	if resp, err := client.UpdateDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the default version of a document.
// If you change a document version for a State Manager association, Systems
// Manager immediately runs the association unless you previously specifed the
// apply-only-at-cron-interval parameter.
func ssm_UpdateDocumentDefaultVersion(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateDocumentDefaultVersionInput{
		// DocumentVersion: *string, // Required
		// Name: *string, // Required
	}

	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}

	if resp, err := client.UpdateDocumentDefaultVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services Systems Manager Change Manager is no longer open to new
// customers. Existing customers can continue to use the service as normal. For
// more information, see [Amazon Web Services Systems Manager Change Manager availability change].
//
// Updates information related to approval reviews for a specific version of a
// change template in Change Manager.
//
// [Amazon Web Services Systems Manager Change Manager availability change]: https://docs.aws.amazon.com/systems-manager/latest/userguide/change-manager-availability-change.html
func ssm_UpdateDocumentMetadata(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateDocumentMetadataInput{
		// DocumentReviews: *types.DocumentReviews, // Required
		// Name: *string, // Required
	}

	if len(_ssmDocumentReviews) > 0 {
		if err := assignInputField(input, "DocumentReviews", _ssmDocumentReviews); err != nil {
			log.Errorf("invalid --document-reviews: %s", err.Error())
			return
		}
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmDocumentVersion) > 0 {
		input.DocumentVersion = aws.String(_ssmDocumentVersion)
	}

	if resp, err := client.UpdateDocumentMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing maintenance window. Only specified parameters are modified.
// The value you specify for Duration determines the specific end time for the
// maintenance window based on the time it begins. No maintenance window tasks are
// permitted to start after the resulting endtime minus the number of hours you
// specify for Cutoff . For example, if the maintenance window starts at 3 PM, the
// duration is three hours, and the value you specify for Cutoff is one hour, no
// maintenance window tasks can start after 5 PM.
func ssm_UpdateMaintenanceWindow(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateMaintenanceWindowInput{
		// WindowId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmAllowUnassociatedTargets) > 0 {
		if err := assignInputField(input, "AllowUnassociatedTargets", _ssmAllowUnassociatedTargets); err != nil {
			log.Errorf("invalid --allow-unassociated-targets: %s", err.Error())
			return
		}
	}
	if len(_ssmCutoff) > 0 {
		if err := assignInputField(input, "Cutoff", _ssmCutoff); err != nil {
			log.Errorf("invalid --cutoff: %s", err.Error())
			return
		}
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmDuration) > 0 {
		if err := assignInputField(input, "Duration", _ssmDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_ssmEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _ssmEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_ssmEndDate) > 0 {
		input.EndDate = aws.String(_ssmEndDate)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmReplace) > 0 {
		if err := assignInputField(input, "Replace", _ssmReplace); err != nil {
			log.Errorf("invalid --replace: %s", err.Error())
			return
		}
	}
	if len(_ssmSchedule) > 0 {
		input.Schedule = aws.String(_ssmSchedule)
	}
	if len(_ssmScheduleOffset) > 0 {
		if err := assignInputField(input, "ScheduleOffset", _ssmScheduleOffset); err != nil {
			log.Errorf("invalid --schedule-offset: %s", err.Error())
			return
		}
	}
	if len(_ssmScheduleTimezone) > 0 {
		input.ScheduleTimezone = aws.String(_ssmScheduleTimezone)
	}
	if len(_ssmStartDate) > 0 {
		input.StartDate = aws.String(_ssmStartDate)
	}

	if resp, err := client.UpdateMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the target of an existing maintenance window. You can change the
// following:
//
// - Name
//
// - Description
//
// - Owner
//
// - IDs for an ID target
//
// - Tags for a Tag target
//
// - From any supported tag type to another. The three supported tag types are
// ID target, Tag target, and resource group. For more information, see Target.
//
// If a parameter is null, then the corresponding field isn't modified.
func ssm_UpdateMaintenanceWindowTarget(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateMaintenanceWindowTargetInput{
		// WindowId: *string, // Required
		// WindowTargetId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmWindowTargetId) > 0 {
		input.WindowTargetId = aws.String(_ssmWindowTargetId)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmOwnerInformation) > 0 {
		input.OwnerInformation = aws.String(_ssmOwnerInformation)
	}
	if len(_ssmReplace) > 0 {
		if err := assignInputField(input, "Replace", _ssmReplace); err != nil {
			log.Errorf("invalid --replace: %s", err.Error())
			return
		}
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMaintenanceWindowTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a task assigned to a maintenance window. You can't change the task
// type, but you can change the following values:
//
// - TaskARN . For example, you can change a RUN_COMMAND task from
// AWS-RunPowerShellScript to AWS-RunShellScript .
//
// - ServiceRoleArn
//
// - TaskInvocationParameters
//
// - Priority
//
// - MaxConcurrency
//
// - MaxErrors
//
// One or more targets must be specified for maintenance window Run Command-type
// tasks. Depending on the task, targets are optional for other maintenance window
// task types (Automation, Lambda, and Step Functions). For more information about
// running tasks that don't specify targets, see [Registering maintenance window tasks without targets]in the Amazon Web Services
// Systems Manager User Guide.
//
// If the value for a parameter in UpdateMaintenanceWindowTask is null, then the
// corresponding field isn't modified. If you set Replace to true, then all fields
// required by the RegisterTaskWithMaintenanceWindowoperation are required for this request. Optional fields that
// aren't specified are set to null.
//
// When you update a maintenance window task that has options specified in
// TaskInvocationParameters , you must provide again all the
// TaskInvocationParameters values that you want to retain. The values you don't
// specify again are removed. For example, suppose that when you registered a Run
// Command task, you specified TaskInvocationParameters values for Comment ,
// NotificationConfig , and OutputS3BucketName . If you update the maintenance
// window task and specify only a different OutputS3BucketName value, the values
// for Comment and NotificationConfig are removed.
//
// [Registering maintenance window tasks without targets]: https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html
func ssm_UpdateMaintenanceWindowTask(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateMaintenanceWindowTaskInput{
		// WindowId: *string, // Required
		// WindowTaskId: *string, // Required
	}

	if len(_ssmWindowId) > 0 {
		input.WindowId = aws.String(_ssmWindowId)
	}
	if len(_ssmWindowTaskId) > 0 {
		input.WindowTaskId = aws.String(_ssmWindowTaskId)
	}
	if len(_ssmAlarmConfiguration) > 0 {
		if err := assignInputField(input, "AlarmConfiguration", _ssmAlarmConfiguration); err != nil {
			log.Errorf("invalid --alarm-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssmCutoffBehavior) > 0 {
		if err := assignInputField(input, "CutoffBehavior", _ssmCutoffBehavior); err != nil {
			log.Errorf("invalid --cutoff-behavior: %s", err.Error())
			return
		}
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmLoggingInfo) > 0 {
		if err := assignInputField(input, "LoggingInfo", _ssmLoggingInfo); err != nil {
			log.Errorf("invalid --logging-info: %s", err.Error())
			return
		}
	}
	if len(_ssmMaxConcurrency) > 0 {
		input.MaxConcurrency = aws.String(_ssmMaxConcurrency)
	}
	if len(_ssmMaxErrors) > 0 {
		input.MaxErrors = aws.String(_ssmMaxErrors)
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmPriority) > 0 {
		if err := assignInputField(input, "Priority", _ssmPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_ssmReplace) > 0 {
		if err := assignInputField(input, "Replace", _ssmReplace); err != nil {
			log.Errorf("invalid --replace: %s", err.Error())
			return
		}
	}
	if len(_ssmServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_ssmServiceRoleArn)
	}
	if len(_ssmTargets) > 0 {
		if err := assignInputField(input, "Targets", _ssmTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_ssmTaskArn) > 0 {
		input.TaskArn = aws.String(_ssmTaskArn)
	}
	if len(_ssmTaskInvocationParameters) > 0 {
		if err := assignInputField(input, "TaskInvocationParameters", _ssmTaskInvocationParameters); err != nil {
			log.Errorf("invalid --task-invocation-parameters: %s", err.Error())
			return
		}
	}
	if len(_ssmTaskParameters) > 0 {
		if err := assignInputField(input, "TaskParameters", _ssmTaskParameters); err != nil {
			log.Errorf("invalid --task-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMaintenanceWindowTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the Identity and Access Management (IAM) role that is assigned to the
// on-premises server, edge device, or virtual machines (VM). IAM roles are first
// assigned to these hybrid nodes during the activation process. For more
// information, see CreateActivation.
func ssm_UpdateManagedInstanceRole(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateManagedInstanceRoleInput{
		// IamRole: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_ssmIamRole) > 0 {
		input.IamRole = aws.String(_ssmIamRole)
	}
	if len(_ssmInstanceId) > 0 {
		input.InstanceId = aws.String(_ssmInstanceId)
	}

	if resp, err := client.UpdateManagedInstanceRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Edit or change an OpsItem. You must have permission in Identity and Access
// Management (IAM) to update an OpsItem. For more information, see [Set up OpsCenter]in the Amazon
// Web Services Systems Manager User Guide.
//
// Operations engineers and IT professionals use Amazon Web Services Systems
// Manager OpsCenter to view, investigate, and remediate operational issues
// impacting the performance and health of their Amazon Web Services resources. For
// more information, see [Amazon Web Services Systems Manager OpsCenter]in the Amazon Web Services Systems Manager User Guide.
//
// [Amazon Web Services Systems Manager OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html
// [Set up OpsCenter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html
func ssm_UpdateOpsItem(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateOpsItemInput{
		// OpsItemId: *string, // Required
	}

	if len(_ssmOpsItemId) > 0 {
		input.OpsItemId = aws.String(_ssmOpsItemId)
	}
	if len(_ssmActualEndTime) > 0 {
		if err := assignInputField(input, "ActualEndTime", _ssmActualEndTime); err != nil {
			log.Errorf("invalid --actual-end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmActualStartTime) > 0 {
		if err := assignInputField(input, "ActualStartTime", _ssmActualStartTime); err != nil {
			log.Errorf("invalid --actual-start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmCategory) > 0 {
		input.Category = aws.String(_ssmCategory)
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmNotifications) > 0 {
		if err := assignInputField(input, "Notifications", _ssmNotifications); err != nil {
			log.Errorf("invalid --notifications: %s", err.Error())
			return
		}
	}
	if len(_ssmOperationalData) > 0 {
		if err := assignInputField(input, "OperationalData", _ssmOperationalData); err != nil {
			log.Errorf("invalid --operational-data: %s", err.Error())
			return
		}
	}
	if len(_ssmOperationalDataToDelete) > 0 {
		input.OperationalDataToDelete = append([]string(nil), _ssmOperationalDataToDelete...)
	}
	if len(_ssmOpsItemArn) > 0 {
		input.OpsItemArn = aws.String(_ssmOpsItemArn)
	}
	if len(_ssmPlannedEndTime) > 0 {
		if err := assignInputField(input, "PlannedEndTime", _ssmPlannedEndTime); err != nil {
			log.Errorf("invalid --planned-end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmPlannedStartTime) > 0 {
		if err := assignInputField(input, "PlannedStartTime", _ssmPlannedStartTime); err != nil {
			log.Errorf("invalid --planned-start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmPriority) > 0 {
		if err := assignInputField(input, "Priority", _ssmPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_ssmRelatedOpsItems) > 0 {
		if err := assignInputField(input, "RelatedOpsItems", _ssmRelatedOpsItems); err != nil {
			log.Errorf("invalid --related-ops-items: %s", err.Error())
			return
		}
	}
	if len(_ssmSeverity) > 0 {
		input.Severity = aws.String(_ssmSeverity)
	}
	if len(_ssmStatus) > 0 {
		if err := assignInputField(input, "Status", _ssmStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_ssmTitle) > 0 {
		input.Title = aws.String(_ssmTitle)
	}

	if resp, err := client.UpdateOpsItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services Systems Manager calls this API operation when you edit
// OpsMetadata in Application Manager.
func ssm_UpdateOpsMetadata(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateOpsMetadataInput{
		// OpsMetadataArn: *string, // Required
	}

	if len(_ssmOpsMetadataArn) > 0 {
		input.OpsMetadataArn = aws.String(_ssmOpsMetadataArn)
	}
	if len(_ssmKeysToDelete) > 0 {
		input.KeysToDelete = append([]string(nil), _ssmKeysToDelete...)
	}
	if len(_ssmMetadataToUpdate) > 0 {
		if err := assignInputField(input, "MetadataToUpdate", _ssmMetadataToUpdate); err != nil {
			log.Errorf("invalid --metadata-to-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOpsMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing patch baseline. Fields not specified in the request are
// left unchanged.
//
// For information about valid key-value pairs in PatchFilters for each supported
// operating system type, see PatchFilter.
func ssm_UpdatePatchBaseline(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdatePatchBaselineInput{
		// BaselineId: *string, // Required
	}

	if len(_ssmBaselineId) > 0 {
		input.BaselineId = aws.String(_ssmBaselineId)
	}
	if len(_ssmApprovalRules) > 0 {
		if err := assignInputField(input, "ApprovalRules", _ssmApprovalRules); err != nil {
			log.Errorf("invalid --approval-rules: %s", err.Error())
			return
		}
	}
	if len(_ssmApprovedPatches) > 0 {
		input.ApprovedPatches = append([]string(nil), _ssmApprovedPatches...)
	}
	if len(_ssmApprovedPatchesComplianceLevel) > 0 {
		if err := assignInputField(input, "ApprovedPatchesComplianceLevel", _ssmApprovedPatchesComplianceLevel); err != nil {
			log.Errorf("invalid --approved-patches-compliance-level: %s", err.Error())
			return
		}
	}
	if len(_ssmApprovedPatchesEnableNonSecurity) > 0 {
		if err := assignInputField(input, "ApprovedPatchesEnableNonSecurity", _ssmApprovedPatchesEnableNonSecurity); err != nil {
			log.Errorf("invalid --approved-patches-enable-non-security: %s", err.Error())
			return
		}
	}
	if len(_ssmAvailableSecurityUpdatesComplianceStatus) > 0 {
		if err := assignInputField(input, "AvailableSecurityUpdatesComplianceStatus", _ssmAvailableSecurityUpdatesComplianceStatus); err != nil {
			log.Errorf("invalid --available-security-updates-compliance-status: %s", err.Error())
			return
		}
	}
	if len(_ssmDescription) > 0 {
		input.Description = aws.String(_ssmDescription)
	}
	if len(_ssmGlobalFilters) > 0 {
		if err := assignInputField(input, "GlobalFilters", _ssmGlobalFilters); err != nil {
			log.Errorf("invalid --global-filters: %s", err.Error())
			return
		}
	}
	if len(_ssmName) > 0 {
		input.Name = aws.String(_ssmName)
	}
	if len(_ssmRejectedPatches) > 0 {
		input.RejectedPatches = append([]string(nil), _ssmRejectedPatches...)
	}
	if len(_ssmRejectedPatchesAction) > 0 {
		if err := assignInputField(input, "RejectedPatchesAction", _ssmRejectedPatchesAction); err != nil {
			log.Errorf("invalid --rejected-patches-action: %s", err.Error())
			return
		}
	}
	if len(_ssmReplace) > 0 {
		if err := assignInputField(input, "Replace", _ssmReplace); err != nil {
			log.Errorf("invalid --replace: %s", err.Error())
			return
		}
	}
	if len(_ssmSources) > 0 {
		if err := assignInputField(input, "Sources", _ssmSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePatchBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a resource data sync. After you create a resource data sync for a
// Region, you can't change the account options for that sync. For example, if you
// create a sync in the us-east-2 (Ohio) Region and you choose the Include only
// the current account option, you can't edit that sync later and choose the
// Include all accounts from my Organizations configuration option. Instead, you
// must delete the first resource data sync, and create a new one.
//
// This API operation only supports a resource data sync that was created with a
// SyncFromSource SyncType .
func ssm_UpdateResourceDataSync(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateResourceDataSyncInput{
		// SyncName: *string, // Required
		// SyncSource: *types.ResourceDataSyncSource, // Required
		// SyncType: *string, // Required
	}

	if len(_ssmSyncName) > 0 {
		input.SyncName = aws.String(_ssmSyncName)
	}
	if len(_ssmSyncSource) > 0 {
		if err := assignInputField(input, "SyncSource", _ssmSyncSource); err != nil {
			log.Errorf("invalid --sync-source: %s", err.Error())
			return
		}
	}
	if len(_ssmSyncType) > 0 {
		input.SyncType = aws.String(_ssmSyncType)
	}

	if resp, err := client.UpdateResourceDataSync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of "false". This means the user can't use
// this feature unless they change the setting to "true" and intentionally opt in
// for a paid feature.
//
// Services map a SettingId object to a setting value. Amazon Web Services
// services teams define the default value for a SettingId . You can't create a new
// SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the GetServiceSetting API operation to
// view the current value. Or, use the ResetServiceSettingto change the value back to the original
// value defined by the Amazon Web Services service team.
//
// Update the service setting for the account.
func ssm_UpdateServiceSetting(cfg aws.Config, client *ssm.Client) {
	input := &ssm.UpdateServiceSettingInput{
		// SettingId: *string, // Required
		// SettingValue: *string, // Required
	}

	if len(_ssmSettingId) > 0 {
		input.SettingId = aws.String(_ssmSettingId)
	}
	if len(_ssmSettingValue) > 0 {
		input.SettingValue = aws.String(_ssmSettingValue)
	}

	if resp, err := client.UpdateServiceSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssmCmd)
	_ssmCmd.Flags().SortFlags = false

	_ssmCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ssmCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssmCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssmCmd.Flags().StringVarP(&_ssmAccessRequestId, "access-request-id", "", "", "Access Request ID")
	_ssmCmd.Flags().StringVarP(&_ssmAccountId, "account-id", "", "", "Account ID")
	_ssmCmd.Flags().StringSliceVarP(&_ssmAccountIdsToAdd, "account-ids-to-add", "", nil, "Account Ids To Add")
	_ssmCmd.Flags().StringSliceVarP(&_ssmAccountIdsToRemove, "account-ids-to-remove", "", nil, "Account Ids To Remove")
	_ssmCmd.Flags().StringVarP(&_ssmActivationId, "activation-id", "", "", "Activation ID")
	_ssmCmd.Flags().StringVarP(&_ssmActualEndTime, "actual-end-time", "", "", "Actual End Time")
	_ssmCmd.Flags().StringVarP(&_ssmActualStartTime, "actual-start-time", "", "", "Actual Start Time")
	_ssmCmd.Flags().StringVarP(&_ssmAggregator, "aggregator", "", "", "Aggregator")
	_ssmCmd.Flags().StringVarP(&_ssmAggregators, "aggregators", "", "", "Aggregators")
	_ssmCmd.Flags().StringVarP(&_ssmAlarmConfiguration, "alarm-configuration", "", "", "Alarm Configuration")
	_ssmCmd.Flags().StringVarP(&_ssmAllowUnassociatedTargets, "allow-unassociated-targets", "", "", "Allow Unassociated Targets")
	_ssmCmd.Flags().StringVarP(&_ssmAllowedPattern, "allowed-pattern", "", "", "Allowed Pattern")
	_ssmCmd.Flags().StringVarP(&_ssmApplyOnlyAtCronInterval, "apply-only-at-cron-interval", "", "", "Apply Only At Cron Interval")
	_ssmCmd.Flags().StringVarP(&_ssmApprovalRules, "approval-rules", "", "", "Approval Rules")
	_ssmCmd.Flags().StringSliceVarP(&_ssmApprovedPatches, "approved-patches", "", nil, "Approved Patches")
	_ssmCmd.Flags().StringVarP(&_ssmApprovedPatchesComplianceLevel, "approved-patches-compliance-level", "", "", "Approved Patches Compliance Level")
	_ssmCmd.Flags().StringVarP(&_ssmApprovedPatchesEnableNonSecurity, "approved-patches-enable-non-security", "", "", "Approved Patches Enable Non Security")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationDispatchAssumeRole, "association-dispatch-assume-role", "", "", "Association Dispatch Assume Role")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationFilterList, "association-filter-list", "", "", "Association Filter List")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationId, "association-id", "", "", "Association ID")
	_ssmCmd.Flags().StringSliceVarP(&_ssmAssociationIds, "association-ids", "", nil, "Association Ids")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationName, "association-name", "", "", "Association Name")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationStatus, "association-status", "", "", "Association Status")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationType, "association-type", "", "", "Association Type")
	_ssmCmd.Flags().StringVarP(&_ssmAssociationVersion, "association-version", "", "", "Association Version")
	_ssmCmd.Flags().StringVarP(&_ssmAtTime, "at-time", "", "", "At Time")
	_ssmCmd.Flags().StringVarP(&_ssmAttachments, "attachments", "", "", "Attachments")
	_ssmCmd.Flags().StringVarP(&_ssmAutoApprove, "auto-approve", "", "", "Auto Approve")
	_ssmCmd.Flags().StringVarP(&_ssmAutomationExecutionId, "automation-execution-id", "", "", "Automation Execution ID")
	_ssmCmd.Flags().StringVarP(&_ssmAutomationTargetParameterName, "automation-target-parameter-name", "", "", "Automation Target Parameter Name")
	_ssmCmd.Flags().StringVarP(&_ssmAvailableSecurityUpdatesComplianceStatus, "available-security-updates-compliance-status", "", "", "Available Security Updates Compliance Status")
	_ssmCmd.Flags().StringVarP(&_ssmBaselineId, "baseline-id", "", "", "Baseline ID")
	_ssmCmd.Flags().StringVarP(&_ssmBaselineOverride, "baseline-override", "", "", "Baseline Override")
	_ssmCmd.Flags().StringSliceVarP(&_ssmCalendarNames, "calendar-names", "", nil, "Calendar Names")
	_ssmCmd.Flags().StringVarP(&_ssmCategory, "category", "", "", "Category")
	_ssmCmd.Flags().StringVarP(&_ssmChangeDetails, "change-details", "", "", "Change Details")
	_ssmCmd.Flags().StringVarP(&_ssmChangeRequestName, "change-request-name", "", "", "Change Request Name")
	_ssmCmd.Flags().StringVarP(&_ssmClientToken, "client-token", "", "", "Client Token")
	_ssmCmd.Flags().StringVarP(&_ssmCloudWatchOutputConfig, "cloud-watch-output-config", "", "", "Cloud Watch Output Config")
	_ssmCmd.Flags().StringVarP(&_ssmCommandId, "command-id", "", "", "Command ID")
	_ssmCmd.Flags().StringVarP(&_ssmComment, "comment", "", "", "Comment")
	_ssmCmd.Flags().StringVarP(&_ssmComplianceSeverity, "compliance-severity", "", "", "Compliance Severity")
	_ssmCmd.Flags().StringVarP(&_ssmComplianceType, "compliance-type", "", "", "Compliance Type")
	_ssmCmd.Flags().StringVarP(&_ssmContent, "content", "", "", "Content")
	_ssmCmd.Flags().StringVarP(&_ssmCutoff, "cutoff", "", "", "Cutoff")
	_ssmCmd.Flags().StringVarP(&_ssmCutoffBehavior, "cutoff-behavior", "", "", "Cutoff Behavior")
	_ssmCmd.Flags().StringVarP(&_ssmDataType, "data-type", "", "", "Data Type")
	_ssmCmd.Flags().StringVarP(&_ssmDefaultInstanceName, "default-instance-name", "", "", "Default Instance Name")
	_ssmCmd.Flags().StringVarP(&_ssmDeletionId, "deletion-id", "", "", "Deletion ID")
	_ssmCmd.Flags().StringVarP(&_ssmDescription, "description", "", "", "Description")
	_ssmCmd.Flags().StringVarP(&_ssmDetails, "details", "", "", "Details")
	_ssmCmd.Flags().StringVarP(&_ssmDisplayName, "display-name", "", "", "Display Name")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentFilterList, "document-filter-list", "", "", "Document Filter List")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentFormat, "document-format", "", "", "Document Format")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentHash, "document-hash", "", "", "Document Hash")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentHashType, "document-hash-type", "", "", "Document Hash Type")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentName, "document-name", "", "", "Document Name")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentReviews, "document-reviews", "", "", "Document Reviews")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentType, "document-type", "", "", "Document Type")
	_ssmCmd.Flags().StringVarP(&_ssmDocumentVersion, "document-version", "", "", "Document Version")
	_ssmCmd.Flags().StringVarP(&_ssmDryRun, "dry-run", "", "", "Dry Run")
	_ssmCmd.Flags().StringVarP(&_ssmDuration, "duration", "", "", "Duration")
	_ssmCmd.Flags().StringVarP(&_ssmEnabled, "enabled", "", "", "Enabled")
	_ssmCmd.Flags().StringVarP(&_ssmEndDate, "end-date", "", "", "End Date")
	_ssmCmd.Flags().StringVarP(&_ssmEntries, "entries", "", "", "Entries")
	_ssmCmd.Flags().StringVarP(&_ssmExecutionId, "execution-id", "", "", "Execution ID")
	_ssmCmd.Flags().StringVarP(&_ssmExecutionInputs, "execution-inputs", "", "", "Execution Inputs")
	_ssmCmd.Flags().StringVarP(&_ssmExecutionPreviewId, "execution-preview-id", "", "", "Execution Preview ID")
	_ssmCmd.Flags().StringVarP(&_ssmExecutionSummary, "execution-summary", "", "", "Execution Summary")
	_ssmCmd.Flags().StringVarP(&_ssmExpirationDate, "expiration-date", "", "", "Expiration Date")
	_ssmCmd.Flags().StringVarP(&_ssmFilters, "filters", "", "", "Filters")
	_ssmCmd.Flags().StringVarP(&_ssmFiltersWithOperator, "filters-with-operator", "", "", "Filters With Operator")
	_ssmCmd.Flags().StringVarP(&_ssmForce, "force", "", "", "Force")
	_ssmCmd.Flags().StringVarP(&_ssmGlobalFilters, "global-filters", "", "", "Global Filters")
	_ssmCmd.Flags().StringVarP(&_ssmIamRole, "iam-role", "", "", "IAM Role")
	_ssmCmd.Flags().StringVarP(&_ssmInstanceId, "instance-id", "", "", "Instance ID")
	_ssmCmd.Flags().StringSliceVarP(&_ssmInstanceIds, "instance-ids", "", nil, "Instance Ids")
	_ssmCmd.Flags().StringVarP(&_ssmInstanceInformationFilterList, "instance-information-filter-list", "", "", "Instance Information Filter List")
	_ssmCmd.Flags().StringVarP(&_ssmInstancePropertyFilterList, "instance-property-filter-list", "", "", "Instance Property Filter List")
	_ssmCmd.Flags().StringVarP(&_ssmInvocationId, "invocation-id", "", "", "Invocation ID")
	_ssmCmd.Flags().StringVarP(&_ssmItemContentHash, "item-content-hash", "", "", "Item Content Hash")
	_ssmCmd.Flags().StringVarP(&_ssmItems, "items", "", "", "Items")
	_ssmCmd.Flags().StringVarP(&_ssmKeyId, "key-id", "", "", "Key ID")
	_ssmCmd.Flags().StringSliceVarP(&_ssmKeysToDelete, "keys-to-delete", "", nil, "Keys To Delete")
	_ssmCmd.Flags().StringSliceVarP(&_ssmLabels, "labels", "", nil, "Labels")
	_ssmCmd.Flags().StringVarP(&_ssmLoggingInfo, "logging-info", "", "", "Logging Info")
	_ssmCmd.Flags().StringVarP(&_ssmMaxConcurrency, "max-concurrency", "", "", "Max Concurrency")
	_ssmCmd.Flags().StringVarP(&_ssmMaxErrors, "max-errors", "", "", "Max Errors")
	_ssmCmd.Flags().StringVarP(&_ssmMaxResults, "max-results", "", "", "Max Results")
	_ssmCmd.Flags().StringVarP(&_ssmMetadata, "metadata", "", "", "Metadata")
	_ssmCmd.Flags().StringVarP(&_ssmMetadataToUpdate, "metadata-to-update", "", "", "Metadata To Update")
	_ssmCmd.Flags().StringVarP(&_ssmMode, "mode", "", "", "Mode")
	_ssmCmd.Flags().StringVarP(&_ssmName, "name", "", "", "Name")
	_ssmCmd.Flags().StringSliceVarP(&_ssmNames, "names", "", nil, "Names")
	_ssmCmd.Flags().StringVarP(&_ssmNextToken, "next-token", "", "", "Next Token")
	_ssmCmd.Flags().StringVarP(&_ssmNotificationConfig, "notification-config", "", "", "Notification Config")
	_ssmCmd.Flags().StringVarP(&_ssmNotifications, "notifications", "", "", "Notifications")
	_ssmCmd.Flags().StringVarP(&_ssmOperatingSystem, "operating-system", "", "", "Operating System")
	_ssmCmd.Flags().StringVarP(&_ssmOperationalData, "operational-data", "", "", "Operational Data")
	_ssmCmd.Flags().StringSliceVarP(&_ssmOperationalDataToDelete, "operational-data-to-delete", "", nil, "Operational Data To Delete")
	_ssmCmd.Flags().StringVarP(&_ssmOpsItemArn, "ops-item-arn", "", "", "Ops Item ARN")
	_ssmCmd.Flags().StringVarP(&_ssmOpsItemFilters, "ops-item-filters", "", "", "Ops Item Filters")
	_ssmCmd.Flags().StringVarP(&_ssmOpsItemId, "ops-item-id", "", "", "Ops Item ID")
	_ssmCmd.Flags().StringVarP(&_ssmOpsItemType, "ops-item-type", "", "", "Ops Item Type")
	_ssmCmd.Flags().StringVarP(&_ssmOpsMetadataArn, "ops-metadata-arn", "", "", "Ops Metadata ARN")
	_ssmCmd.Flags().StringVarP(&_ssmOutputLocation, "output-location", "", "", "Output Location")
	_ssmCmd.Flags().StringVarP(&_ssmOutputS3BucketName, "output-s3-bucket-name", "", "", "Output S3 Bucket Name")
	_ssmCmd.Flags().StringVarP(&_ssmOutputS3KeyPrefix, "output-s3-key-prefix", "", "", "Output S3 Key Prefix")
	_ssmCmd.Flags().StringVarP(&_ssmOutputS3Region, "output-s3-region", "", "", "Output S3 Region")
	_ssmCmd.Flags().StringVarP(&_ssmOverwrite, "overwrite", "", "", "Overwrite")
	_ssmCmd.Flags().StringVarP(&_ssmOwnerInformation, "owner-information", "", "", "Owner Information")
	_ssmCmd.Flags().StringVarP(&_ssmParameterFilters, "parameter-filters", "", "", "Parameter Filters")
	_ssmCmd.Flags().StringVarP(&_ssmParameterVersion, "parameter-version", "", "", "Parameter Version")
	_ssmCmd.Flags().StringVarP(&_ssmParameters, "parameters", "", "", "Parameters")
	_ssmCmd.Flags().StringVarP(&_ssmPatchGroup, "patch-group", "", "", "Patch Group")
	_ssmCmd.Flags().StringVarP(&_ssmPatchSet, "patch-set", "", "", "Patch Set")
	_ssmCmd.Flags().StringVarP(&_ssmPath, "path", "", "", "Path")
	_ssmCmd.Flags().StringVarP(&_ssmPayload, "payload", "", "", "Payload")
	_ssmCmd.Flags().StringVarP(&_ssmPermissionType, "permission-type", "", "", "Permission Type")
	_ssmCmd.Flags().StringVarP(&_ssmPlannedEndTime, "planned-end-time", "", "", "Planned End Time")
	_ssmCmd.Flags().StringVarP(&_ssmPlannedStartTime, "planned-start-time", "", "", "Planned Start Time")
	_ssmCmd.Flags().StringVarP(&_ssmPluginName, "plugin-name", "", "", "Plugin Name")
	_ssmCmd.Flags().StringVarP(&_ssmPolicies, "policies", "", "", "Policies")
	_ssmCmd.Flags().StringVarP(&_ssmPolicy, "policy", "", "", "Policy")
	_ssmCmd.Flags().StringVarP(&_ssmPolicyHash, "policy-hash", "", "", "Policy Hash")
	_ssmCmd.Flags().StringVarP(&_ssmPolicyId, "policy-id", "", "", "Policy ID")
	_ssmCmd.Flags().StringVarP(&_ssmPriority, "priority", "", "", "Priority")
	_ssmCmd.Flags().StringVarP(&_ssmProperty, "property", "", "", "Property")
	_ssmCmd.Flags().StringVarP(&_ssmReason, "reason", "", "", "Reason")
	_ssmCmd.Flags().StringVarP(&_ssmRecursive, "recursive", "", "", "Recursive")
	_ssmCmd.Flags().StringVarP(&_ssmRegistrationLimit, "registration-limit", "", "", "Registration Limit")
	_ssmCmd.Flags().StringVarP(&_ssmRegistrationMetadata, "registration-metadata", "", "", "Registration Metadata")
	_ssmCmd.Flags().StringSliceVarP(&_ssmRejectedPatches, "rejected-patches", "", nil, "Rejected Patches")
	_ssmCmd.Flags().StringVarP(&_ssmRejectedPatchesAction, "rejected-patches-action", "", "", "Rejected Patches Action")
	_ssmCmd.Flags().StringVarP(&_ssmRelatedOpsItems, "related-ops-items", "", "", "Related Ops Items")
	_ssmCmd.Flags().StringVarP(&_ssmReplace, "replace", "", "", "Replace")
	_ssmCmd.Flags().StringVarP(&_ssmRequires, "requires", "", "", "Requires")
	_ssmCmd.Flags().StringVarP(&_ssmResourceArn, "resource-arn", "", "", "Resource ARN")
	_ssmCmd.Flags().StringVarP(&_ssmResourceId, "resource-id", "", "", "Resource ID")
	_ssmCmd.Flags().StringSliceVarP(&_ssmResourceIds, "resource-ids", "", nil, "Resource Ids")
	_ssmCmd.Flags().StringVarP(&_ssmResourceType, "resource-type", "", "", "Resource Type")
	_ssmCmd.Flags().StringSliceVarP(&_ssmResourceTypes, "resource-types", "", nil, "Resource Types")
	_ssmCmd.Flags().StringVarP(&_ssmResourceUri, "resource-uri", "", "", "Resource URI")
	_ssmCmd.Flags().StringVarP(&_ssmResultAttributes, "result-attributes", "", "", "Result Attributes")
	_ssmCmd.Flags().StringVarP(&_ssmReverseOrder, "reverse-order", "", "", "Reverse Order")
	_ssmCmd.Flags().StringVarP(&_ssmRunbooks, "runbooks", "", "", "Runbooks")
	_ssmCmd.Flags().StringVarP(&_ssmS3Destination, "s3-destination", "", "", "S3 Destination")
	_ssmCmd.Flags().StringVarP(&_ssmSafe, "safe", "", "", "Safe")
	_ssmCmd.Flags().StringVarP(&_ssmSchedule, "schedule", "", "", "Schedule")
	_ssmCmd.Flags().StringVarP(&_ssmScheduleExpression, "schedule-expression", "", "", "Schedule Expression")
	_ssmCmd.Flags().StringVarP(&_ssmScheduleOffset, "schedule-offset", "", "", "Schedule Offset")
	_ssmCmd.Flags().StringVarP(&_ssmScheduleTimezone, "schedule-timezone", "", "", "Schedule Timezone")
	_ssmCmd.Flags().StringVarP(&_ssmScheduledEndTime, "scheduled-end-time", "", "", "Scheduled End Time")
	_ssmCmd.Flags().StringVarP(&_ssmScheduledTime, "scheduled-time", "", "", "Scheduled Time")
	_ssmCmd.Flags().StringVarP(&_ssmSchemaDeleteOption, "schema-delete-option", "", "", "Schema Delete Option")
	_ssmCmd.Flags().StringVarP(&_ssmServiceRoleArn, "service-role-arn", "", "", "Service Role ARN")
	_ssmCmd.Flags().StringVarP(&_ssmSessionId, "session-id", "", "", "Session ID")
	_ssmCmd.Flags().StringVarP(&_ssmSettingId, "setting-id", "", "", "Setting ID")
	_ssmCmd.Flags().StringVarP(&_ssmSettingValue, "setting-value", "", "", "Setting Value")
	_ssmCmd.Flags().StringVarP(&_ssmSeverity, "severity", "", "", "Severity")
	_ssmCmd.Flags().StringVarP(&_ssmShared, "shared", "", "", "Shared")
	_ssmCmd.Flags().StringVarP(&_ssmSharedDocumentVersion, "shared-document-version", "", "", "Shared Document Version")
	_ssmCmd.Flags().StringVarP(&_ssmSignalType, "signal-type", "", "", "Signal Type")
	_ssmCmd.Flags().StringVarP(&_ssmSnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_ssmCmd.Flags().StringVarP(&_ssmSource, "source", "", "", "Source")
	_ssmCmd.Flags().StringVarP(&_ssmSources, "sources", "", "", "Sources")
	_ssmCmd.Flags().StringVarP(&_ssmStartDate, "start-date", "", "", "Start Date")
	_ssmCmd.Flags().StringVarP(&_ssmState, "state", "", "", "State")
	_ssmCmd.Flags().StringVarP(&_ssmStatus, "status", "", "", "Status")
	_ssmCmd.Flags().StringVarP(&_ssmSubType, "sub-type", "", "", "Sub Type")
	_ssmCmd.Flags().StringVarP(&_ssmSyncCompliance, "sync-compliance", "", "", "Sync Compliance")
	_ssmCmd.Flags().StringVarP(&_ssmSyncName, "sync-name", "", "", "Sync Name")
	_ssmCmd.Flags().StringVarP(&_ssmSyncSource, "sync-source", "", "", "Sync Source")
	_ssmCmd.Flags().StringVarP(&_ssmSyncType, "sync-type", "", "", "Sync Type")
	_ssmCmd.Flags().StringSliceVarP(&_ssmTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ssmCmd.Flags().StringVarP(&_ssmTags, "tags", "", "", "Tags")
	_ssmCmd.Flags().StringVarP(&_ssmTarget, "target", "", "", "Target")
	_ssmCmd.Flags().StringVarP(&_ssmTargetLocations, "target-locations", "", "", "Target Locations")
	_ssmCmd.Flags().StringVarP(&_ssmTargetLocationsURL, "target-locations-url", "", "", "Target Locations URL")
	_ssmCmd.Flags().StringVarP(&_ssmTargetMaps, "target-maps", "", "", "Target Maps")
	_ssmCmd.Flags().StringVarP(&_ssmTargetParameterName, "target-parameter-name", "", "", "Target Parameter Name")
	_ssmCmd.Flags().StringVarP(&_ssmTargetType, "target-type", "", "", "Target Type")
	_ssmCmd.Flags().StringVarP(&_ssmTargets, "targets", "", "", "Targets")
	_ssmCmd.Flags().StringVarP(&_ssmTaskArn, "task-arn", "", "", "Task ARN")
	_ssmCmd.Flags().StringVarP(&_ssmTaskId, "task-id", "", "", "Task ID")
	_ssmCmd.Flags().StringVarP(&_ssmTaskInvocationParameters, "task-invocation-parameters", "", "", "Task Invocation Parameters")
	_ssmCmd.Flags().StringVarP(&_ssmTaskParameters, "task-parameters", "", "", "Task Parameters")
	_ssmCmd.Flags().StringVarP(&_ssmTaskType, "task-type", "", "", "Task Type")
	_ssmCmd.Flags().StringVarP(&_ssmTier, "tier", "", "", "Tier")
	_ssmCmd.Flags().StringVarP(&_ssmTimeoutSeconds, "timeout-seconds", "", "", "Timeout Seconds")
	_ssmCmd.Flags().StringVarP(&_ssmTitle, "title", "", "", "Title")
	_ssmCmd.Flags().StringVarP(&_ssmType, "type", "", "", "Type")
	_ssmCmd.Flags().StringVarP(&_ssmTypeName, "type-name", "", "", "Type Name")
	_ssmCmd.Flags().StringVarP(&_ssmUploadType, "upload-type", "", "", "Upload Type")
	_ssmCmd.Flags().StringVarP(&_ssmUseS3DualStackEndpoint, "use-s3-dual-stack-endpoint", "", "", "Use S3 Dual Stack Endpoint")
	_ssmCmd.Flags().StringVarP(&_ssmValue, "value", "", "", "Value")
	_ssmCmd.Flags().StringVarP(&_ssmVersionName, "version-name", "", "", "Version Name")
	_ssmCmd.Flags().StringVarP(&_ssmWindowExecutionId, "window-execution-id", "", "", "Window Execution ID")
	_ssmCmd.Flags().StringVarP(&_ssmWindowId, "window-id", "", "", "Window ID")
	_ssmCmd.Flags().StringVarP(&_ssmWindowTargetId, "window-target-id", "", "", "Window Target ID")
	_ssmCmd.Flags().StringVarP(&_ssmWindowTaskId, "window-task-id", "", "", "Window Task ID")
	_ssmCmd.Flags().StringVarP(&_ssmWithDecryption, "with-decryption", "", "", "With Decryption")

	_ssmCmd.Flags().BoolVarP(&_ssmAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_ssmCmd.Flags().BoolVarP(&_ssmAssociateOpsItemRelatedItem, "associate-ops-item-related-item", "", false, "Associate Ops Item Related Item")
	_ssmCmd.Flags().BoolVarP(&_ssmCancelCommand, "cancel-command", "", false, "Cancel Command")
	_ssmCmd.Flags().BoolVarP(&_ssmCancelMaintenanceWindowExecution, "cancel-maintenance-window-execution", "", false, "Cancel Maintenance Window Execution")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateActivation, "create-activation", "", false, "Create Activation")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateAssociation, "create-association", "", false, "Create Association")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateAssociationBatch, "create-association-batch", "", false, "Create Association Batch")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateDocument, "create-document", "", false, "Create Document")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateMaintenanceWindow, "create-maintenance-window", "", false, "Create Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateOpsItem, "create-ops-item", "", false, "Create Ops Item")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateOpsMetadata, "create-ops-metadata", "", false, "Create Ops Metadata")
	_ssmCmd.Flags().BoolVarP(&_ssmCreatePatchBaseline, "create-patch-baseline", "", false, "Create Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmCreateResourceDataSync, "create-resource-data-sync", "", false, "Create Resource Data Sync")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteActivation, "delete-activation", "", false, "Delete Activation")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteAssociation, "delete-association", "", false, "Delete Association")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteDocument, "delete-document", "", false, "Delete Document")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteInventory, "delete-inventory", "", false, "Delete Inventory")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteMaintenanceWindow, "delete-maintenance-window", "", false, "Delete Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteOpsItem, "delete-ops-item", "", false, "Delete Ops Item")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteOpsMetadata, "delete-ops-metadata", "", false, "Delete Ops Metadata")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteParameter, "delete-parameter", "", false, "Delete Parameter")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteParameters, "delete-parameters", "", false, "Delete Parameters")
	_ssmCmd.Flags().BoolVarP(&_ssmDeletePatchBaseline, "delete-patch-baseline", "", false, "Delete Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteResourceDataSync, "delete-resource-data-sync", "", false, "Delete Resource Data Sync")
	_ssmCmd.Flags().BoolVarP(&_ssmDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_ssmCmd.Flags().BoolVarP(&_ssmDeregisterManagedInstance, "deregister-managed-instance", "", false, "Deregister Managed Instance")
	_ssmCmd.Flags().BoolVarP(&_ssmDeregisterPatchBaselineForPatchGroup, "deregister-patch-baseline-for-patch-group", "", false, "Deregister Patch Baseline For Patch Group")
	_ssmCmd.Flags().BoolVarP(&_ssmDeregisterTargetFromMaintenanceWindow, "deregister-target-from-maintenance-window", "", false, "Deregister Target From Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmDeregisterTaskFromMaintenanceWindow, "deregister-task-from-maintenance-window", "", false, "Deregister Task From Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeActivations, "describe-activations", "", false, "Describe Activations")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeAssociation, "describe-association", "", false, "Describe Association")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeAssociationExecutionTargets, "describe-association-execution-targets", "", false, "Describe Association Execution Targets")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeAssociationExecutions, "describe-association-executions", "", false, "Describe Association Executions")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeAutomationExecutions, "describe-automation-executions", "", false, "Describe Automation Executions")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeAutomationStepExecutions, "describe-automation-step-executions", "", false, "Describe Automation Step Executions")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeAvailablePatches, "describe-available-patches", "", false, "Describe Available Patches")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeDocument, "describe-document", "", false, "Describe Document")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeDocumentPermission, "describe-document-permission", "", false, "Describe Document Permission")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeEffectiveInstanceAssociations, "describe-effective-instance-associations", "", false, "Describe Effective Instance Associations")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeEffectivePatchesForPatchBaseline, "describe-effective-patches-for-patch-baseline", "", false, "Describe Effective Patches For Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInstanceAssociationsStatus, "describe-instance-associations-status", "", false, "Describe Instance Associations Status")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInstanceInformation, "describe-instance-information", "", false, "Describe Instance Information")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInstancePatchStates, "describe-instance-patch-states", "", false, "Describe Instance Patch States")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInstancePatchStatesForPatchGroup, "describe-instance-patch-states-for-patch-group", "", false, "Describe Instance Patch States For Patch Group")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInstancePatches, "describe-instance-patches", "", false, "Describe Instance Patches")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInstanceProperties, "describe-instance-properties", "", false, "Describe Instance Properties")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeInventoryDeletions, "describe-inventory-deletions", "", false, "Describe Inventory Deletions")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowExecutionTaskInvocations, "describe-maintenance-window-execution-task-invocations", "", false, "Describe Maintenance Window Execution Task Invocations")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowExecutionTasks, "describe-maintenance-window-execution-tasks", "", false, "Describe Maintenance Window Execution Tasks")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowExecutions, "describe-maintenance-window-executions", "", false, "Describe Maintenance Window Executions")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowSchedule, "describe-maintenance-window-schedule", "", false, "Describe Maintenance Window Schedule")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowTargets, "describe-maintenance-window-targets", "", false, "Describe Maintenance Window Targets")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowTasks, "describe-maintenance-window-tasks", "", false, "Describe Maintenance Window Tasks")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindows, "describe-maintenance-windows", "", false, "Describe Maintenance Windows")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeMaintenanceWindowsForTarget, "describe-maintenance-windows-for-target", "", false, "Describe Maintenance Windows For Target")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeOpsItems, "describe-ops-items", "", false, "Describe Ops Items")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeParameters, "describe-parameters", "", false, "Describe Parameters")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribePatchBaselines, "describe-patch-baselines", "", false, "Describe Patch Baselines")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribePatchGroupState, "describe-patch-group-state", "", false, "Describe Patch Group State")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribePatchGroups, "describe-patch-groups", "", false, "Describe Patch Groups")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribePatchProperties, "describe-patch-properties", "", false, "Describe Patch Properties")
	_ssmCmd.Flags().BoolVarP(&_ssmDescribeSessions, "describe-sessions", "", false, "Describe Sessions")
	_ssmCmd.Flags().BoolVarP(&_ssmDisassociateOpsItemRelatedItem, "disassociate-ops-item-related-item", "", false, "Disassociate Ops Item Related Item")
	_ssmCmd.Flags().BoolVarP(&_ssmGetAccessToken, "get-access-token", "", false, "Get Access Token")
	_ssmCmd.Flags().BoolVarP(&_ssmGetAutomationExecution, "get-automation-execution", "", false, "Get Automation Execution")
	_ssmCmd.Flags().BoolVarP(&_ssmGetCalendarState, "get-calendar-state", "", false, "Get Calendar State")
	_ssmCmd.Flags().BoolVarP(&_ssmGetCommandInvocation, "get-command-invocation", "", false, "Get Command Invocation")
	_ssmCmd.Flags().BoolVarP(&_ssmGetConnectionStatus, "get-connection-status", "", false, "Get Connection Status")
	_ssmCmd.Flags().BoolVarP(&_ssmGetDefaultPatchBaseline, "get-default-patch-baseline", "", false, "Get Default Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmGetDeployablePatchSnapshotForInstance, "get-deployable-patch-snapshot-for-instance", "", false, "Get Deployable Patch Snapshot For Instance")
	_ssmCmd.Flags().BoolVarP(&_ssmGetDocument, "get-document", "", false, "Get Document")
	_ssmCmd.Flags().BoolVarP(&_ssmGetExecutionPreview, "get-execution-preview", "", false, "Get Execution Preview")
	_ssmCmd.Flags().BoolVarP(&_ssmGetInventory, "get-inventory", "", false, "Get Inventory")
	_ssmCmd.Flags().BoolVarP(&_ssmGetInventorySchema, "get-inventory-schema", "", false, "Get Inventory Schema")
	_ssmCmd.Flags().BoolVarP(&_ssmGetMaintenanceWindow, "get-maintenance-window", "", false, "Get Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmGetMaintenanceWindowExecution, "get-maintenance-window-execution", "", false, "Get Maintenance Window Execution")
	_ssmCmd.Flags().BoolVarP(&_ssmGetMaintenanceWindowExecutionTask, "get-maintenance-window-execution-task", "", false, "Get Maintenance Window Execution Task")
	_ssmCmd.Flags().BoolVarP(&_ssmGetMaintenanceWindowExecutionTaskInvocation, "get-maintenance-window-execution-task-invocation", "", false, "Get Maintenance Window Execution Task Invocation")
	_ssmCmd.Flags().BoolVarP(&_ssmGetMaintenanceWindowTask, "get-maintenance-window-task", "", false, "Get Maintenance Window Task")
	_ssmCmd.Flags().BoolVarP(&_ssmGetOpsItem, "get-ops-item", "", false, "Get Ops Item")
	_ssmCmd.Flags().BoolVarP(&_ssmGetOpsMetadata, "get-ops-metadata", "", false, "Get Ops Metadata")
	_ssmCmd.Flags().BoolVarP(&_ssmGetOpsSummary, "get-ops-summary", "", false, "Get Ops Summary")
	_ssmCmd.Flags().BoolVarP(&_ssmGetParameter, "get-parameter", "", false, "Get Parameter")
	_ssmCmd.Flags().BoolVarP(&_ssmGetParameterHistory, "get-parameter-history", "", false, "Get Parameter History")
	_ssmCmd.Flags().BoolVarP(&_ssmGetParameters, "get-parameters", "", false, "Get Parameters")
	_ssmCmd.Flags().BoolVarP(&_ssmGetParametersByPath, "get-parameters-by-path", "", false, "Get Parameters By Path")
	_ssmCmd.Flags().BoolVarP(&_ssmGetPatchBaseline, "get-patch-baseline", "", false, "Get Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmGetPatchBaselineForPatchGroup, "get-patch-baseline-for-patch-group", "", false, "Get Patch Baseline For Patch Group")
	_ssmCmd.Flags().BoolVarP(&_ssmGetResourcePolicies, "get-resource-policies", "", false, "Get Resource Policies")
	_ssmCmd.Flags().BoolVarP(&_ssmGetServiceSetting, "get-service-setting", "", false, "Get Service Setting")
	_ssmCmd.Flags().BoolVarP(&_ssmLabelParameterVersion, "label-parameter-version", "", false, "Label Parameter Version")
	_ssmCmd.Flags().BoolVarP(&_ssmListAssociationVersions, "list-association-versions", "", false, "List Association Versions")
	_ssmCmd.Flags().BoolVarP(&_ssmListAssociations, "list-associations", "", false, "List Associations")
	_ssmCmd.Flags().BoolVarP(&_ssmListCommandInvocations, "list-command-invocations", "", false, "List Command Invocations")
	_ssmCmd.Flags().BoolVarP(&_ssmListCommands, "list-commands", "", false, "List Commands")
	_ssmCmd.Flags().BoolVarP(&_ssmListComplianceItems, "list-compliance-items", "", false, "List Compliance Items")
	_ssmCmd.Flags().BoolVarP(&_ssmListComplianceSummaries, "list-compliance-summaries", "", false, "List Compliance Summaries")
	_ssmCmd.Flags().BoolVarP(&_ssmListDocumentMetadataHistory, "list-document-metadata-history", "", false, "List Document Metadata History")
	_ssmCmd.Flags().BoolVarP(&_ssmListDocumentVersions, "list-document-versions", "", false, "List Document Versions")
	_ssmCmd.Flags().BoolVarP(&_ssmListDocuments, "list-documents", "", false, "List Documents")
	_ssmCmd.Flags().BoolVarP(&_ssmListInventoryEntries, "list-inventory-entries", "", false, "List Inventory Entries")
	_ssmCmd.Flags().BoolVarP(&_ssmListNodes, "list-nodes", "", false, "List Nodes")
	_ssmCmd.Flags().BoolVarP(&_ssmListNodesSummary, "list-nodes-summary", "", false, "List Nodes Summary")
	_ssmCmd.Flags().BoolVarP(&_ssmListOpsItemEvents, "list-ops-item-events", "", false, "List Ops Item Events")
	_ssmCmd.Flags().BoolVarP(&_ssmListOpsItemRelatedItems, "list-ops-item-related-items", "", false, "List Ops Item Related Items")
	_ssmCmd.Flags().BoolVarP(&_ssmListOpsMetadata, "list-ops-metadata", "", false, "List Ops Metadata")
	_ssmCmd.Flags().BoolVarP(&_ssmListResourceComplianceSummaries, "list-resource-compliance-summaries", "", false, "List Resource Compliance Summaries")
	_ssmCmd.Flags().BoolVarP(&_ssmListResourceDataSync, "list-resource-data-sync", "", false, "List Resource Data Sync")
	_ssmCmd.Flags().BoolVarP(&_ssmListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ssmCmd.Flags().BoolVarP(&_ssmModifyDocumentPermission, "modify-document-permission", "", false, "Modify Document Permission")
	_ssmCmd.Flags().BoolVarP(&_ssmPutComplianceItems, "put-compliance-items", "", false, "Put Compliance Items")
	_ssmCmd.Flags().BoolVarP(&_ssmPutInventory, "put-inventory", "", false, "Put Inventory")
	_ssmCmd.Flags().BoolVarP(&_ssmPutParameter, "put-parameter", "", false, "Put Parameter")
	_ssmCmd.Flags().BoolVarP(&_ssmPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_ssmCmd.Flags().BoolVarP(&_ssmRegisterDefaultPatchBaseline, "register-default-patch-baseline", "", false, "Register Default Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmRegisterPatchBaselineForPatchGroup, "register-patch-baseline-for-patch-group", "", false, "Register Patch Baseline For Patch Group")
	_ssmCmd.Flags().BoolVarP(&_ssmRegisterTargetWithMaintenanceWindow, "register-target-with-maintenance-window", "", false, "Register Target With Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmRegisterTaskWithMaintenanceWindow, "register-task-with-maintenance-window", "", false, "Register Task With Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_ssmCmd.Flags().BoolVarP(&_ssmResetServiceSetting, "reset-service-setting", "", false, "Reset Service Setting")
	_ssmCmd.Flags().BoolVarP(&_ssmResumeSession, "resume-session", "", false, "Resume Session")
	_ssmCmd.Flags().BoolVarP(&_ssmSendAutomationSignal, "send-automation-signal", "", false, "Send Automation Signal")
	_ssmCmd.Flags().BoolVarP(&_ssmSendCommand, "send-command", "", false, "Send Command")
	_ssmCmd.Flags().BoolVarP(&_ssmStartAccessRequest, "start-access-request", "", false, "Start Access Request")
	_ssmCmd.Flags().BoolVarP(&_ssmStartAssociationsOnce, "start-associations-once", "", false, "Start Associations Once")
	_ssmCmd.Flags().BoolVarP(&_ssmStartAutomationExecution, "start-automation-execution", "", false, "Start Automation Execution")
	_ssmCmd.Flags().BoolVarP(&_ssmStartChangeRequestExecution, "start-change-request-execution", "", false, "Start Change Request Execution")
	_ssmCmd.Flags().BoolVarP(&_ssmStartExecutionPreview, "start-execution-preview", "", false, "Start Execution Preview")
	_ssmCmd.Flags().BoolVarP(&_ssmStartSession, "start-session", "", false, "Start Session")
	_ssmCmd.Flags().BoolVarP(&_ssmStopAutomationExecution, "stop-automation-execution", "", false, "Stop Automation Execution")
	_ssmCmd.Flags().BoolVarP(&_ssmTerminateSession, "terminate-session", "", false, "Terminate Session")
	_ssmCmd.Flags().BoolVarP(&_ssmUnlabelParameterVersion, "unlabel-parameter-version", "", false, "Unlabel Parameter Version")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateAssociation, "update-association", "", false, "Update Association")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateAssociationStatus, "update-association-status", "", false, "Update Association Status")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateDocument, "update-document", "", false, "Update Document")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateDocumentDefaultVersion, "update-document-default-version", "", false, "Update Document Default Version")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateDocumentMetadata, "update-document-metadata", "", false, "Update Document Metadata")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateMaintenanceWindow, "update-maintenance-window", "", false, "Update Maintenance Window")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateMaintenanceWindowTarget, "update-maintenance-window-target", "", false, "Update Maintenance Window Target")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateMaintenanceWindowTask, "update-maintenance-window-task", "", false, "Update Maintenance Window Task")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateManagedInstanceRole, "update-managed-instance-role", "", false, "Update Managed Instance Role")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateOpsItem, "update-ops-item", "", false, "Update Ops Item")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateOpsMetadata, "update-ops-metadata", "", false, "Update Ops Metadata")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdatePatchBaseline, "update-patch-baseline", "", false, "Update Patch Baseline")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateResourceDataSync, "update-resource-data-sync", "", false, "Update Resource Data Sync")
	_ssmCmd.Flags().BoolVarP(&_ssmUpdateServiceSetting, "update-service-setting", "", false, "Update Service Setting")

}
