package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var _backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "AWS backup CLI",
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
		client := backup.NewFromConfig(cfg)
		if _backupAssociateBackupVaultMpaApprovalTeam {
			backup_AssociateBackupVaultMpaApprovalTeam(cfg, client)
			return
		}
		if _backupCancelLegalHold {
			backup_CancelLegalHold(cfg, client)
			return
		}
		if _backupCreateBackupPlan {
			backup_CreateBackupPlan(cfg, client)
			return
		}
		if _backupCreateBackupSelection {
			backup_CreateBackupSelection(cfg, client)
			return
		}
		if _backupCreateBackupVault {
			backup_CreateBackupVault(cfg, client)
			return
		}
		if _backupCreateFramework {
			backup_CreateFramework(cfg, client)
			return
		}
		if _backupCreateLegalHold {
			backup_CreateLegalHold(cfg, client)
			return
		}
		if _backupCreateLogicallyAirGappedBackupVault {
			backup_CreateLogicallyAirGappedBackupVault(cfg, client)
			return
		}
		if _backupCreateReportPlan {
			backup_CreateReportPlan(cfg, client)
			return
		}
		if _backupCreateRestoreAccessBackupVault {
			backup_CreateRestoreAccessBackupVault(cfg, client)
			return
		}
		if _backupCreateRestoreTestingPlan {
			backup_CreateRestoreTestingPlan(cfg, client)
			return
		}
		if _backupCreateRestoreTestingSelection {
			backup_CreateRestoreTestingSelection(cfg, client)
			return
		}
		if _backupCreateTieringConfiguration {
			backup_CreateTieringConfiguration(cfg, client)
			return
		}
		if _backupDeleteBackupPlan {
			backup_DeleteBackupPlan(cfg, client)
			return
		}
		if _backupDeleteBackupSelection {
			backup_DeleteBackupSelection(cfg, client)
			return
		}
		if _backupDeleteBackupVault {
			backup_DeleteBackupVault(cfg, client)
			return
		}
		if _backupDeleteBackupVaultAccessPolicy {
			backup_DeleteBackupVaultAccessPolicy(cfg, client)
			return
		}
		if _backupDeleteBackupVaultLockConfiguration {
			backup_DeleteBackupVaultLockConfiguration(cfg, client)
			return
		}
		if _backupDeleteBackupVaultNotifications {
			backup_DeleteBackupVaultNotifications(cfg, client)
			return
		}
		if _backupDeleteFramework {
			backup_DeleteFramework(cfg, client)
			return
		}
		if _backupDeleteRecoveryPoint {
			backup_DeleteRecoveryPoint(cfg, client)
			return
		}
		if _backupDeleteReportPlan {
			backup_DeleteReportPlan(cfg, client)
			return
		}
		if _backupDeleteRestoreTestingPlan {
			backup_DeleteRestoreTestingPlan(cfg, client)
			return
		}
		if _backupDeleteRestoreTestingSelection {
			backup_DeleteRestoreTestingSelection(cfg, client)
			return
		}
		if _backupDeleteTieringConfiguration {
			backup_DeleteTieringConfiguration(cfg, client)
			return
		}
		if _backupDescribeBackupJob {
			backup_DescribeBackupJob(cfg, client)
			return
		}
		if _backupDescribeBackupVault {
			backup_DescribeBackupVault(cfg, client)
			return
		}
		if _backupDescribeCopyJob {
			backup_DescribeCopyJob(cfg, client)
			return
		}
		if _backupDescribeFramework {
			backup_DescribeFramework(cfg, client)
			return
		}
		if _backupDescribeGlobalSettings {
			backup_DescribeGlobalSettings(cfg, client)
			return
		}
		if _backupDescribeProtectedResource {
			backup_DescribeProtectedResource(cfg, client)
			return
		}
		if _backupDescribeRecoveryPoint {
			backup_DescribeRecoveryPoint(cfg, client)
			return
		}
		if _backupDescribeRegionSettings {
			backup_DescribeRegionSettings(cfg, client)
			return
		}
		if _backupDescribeReportJob {
			backup_DescribeReportJob(cfg, client)
			return
		}
		if _backupDescribeReportPlan {
			backup_DescribeReportPlan(cfg, client)
			return
		}
		if _backupDescribeRestoreJob {
			backup_DescribeRestoreJob(cfg, client)
			return
		}
		if _backupDescribeScanJob {
			backup_DescribeScanJob(cfg, client)
			return
		}
		if _backupDisassociateBackupVaultMpaApprovalTeam {
			backup_DisassociateBackupVaultMpaApprovalTeam(cfg, client)
			return
		}
		if _backupDisassociateRecoveryPoint {
			backup_DisassociateRecoveryPoint(cfg, client)
			return
		}
		if _backupDisassociateRecoveryPointFromParent {
			backup_DisassociateRecoveryPointFromParent(cfg, client)
			return
		}
		if _backupExportBackupPlanTemplate {
			backup_ExportBackupPlanTemplate(cfg, client)
			return
		}
		if _backupGetBackupPlan {
			backup_GetBackupPlan(cfg, client)
			return
		}
		if _backupGetBackupPlanFromJSON {
			backup_GetBackupPlanFromJSON(cfg, client)
			return
		}
		if _backupGetBackupPlanFromTemplate {
			backup_GetBackupPlanFromTemplate(cfg, client)
			return
		}
		if _backupGetBackupSelection {
			backup_GetBackupSelection(cfg, client)
			return
		}
		if _backupGetBackupVaultAccessPolicy {
			backup_GetBackupVaultAccessPolicy(cfg, client)
			return
		}
		if _backupGetBackupVaultNotifications {
			backup_GetBackupVaultNotifications(cfg, client)
			return
		}
		if _backupGetLegalHold {
			backup_GetLegalHold(cfg, client)
			return
		}
		if _backupGetRecoveryPointIndexDetails {
			backup_GetRecoveryPointIndexDetails(cfg, client)
			return
		}
		if _backupGetRecoveryPointRestoreMetadata {
			backup_GetRecoveryPointRestoreMetadata(cfg, client)
			return
		}
		if _backupGetRestoreJobMetadata {
			backup_GetRestoreJobMetadata(cfg, client)
			return
		}
		if _backupGetRestoreTestingInferredMetadata {
			backup_GetRestoreTestingInferredMetadata(cfg, client)
			return
		}
		if _backupGetRestoreTestingPlan {
			backup_GetRestoreTestingPlan(cfg, client)
			return
		}
		if _backupGetRestoreTestingSelection {
			backup_GetRestoreTestingSelection(cfg, client)
			return
		}
		if _backupGetSupportedResourceTypes {
			backup_GetSupportedResourceTypes(cfg, client)
			return
		}
		if _backupGetTieringConfiguration {
			backup_GetTieringConfiguration(cfg, client)
			return
		}
		if _backupListBackupJobSummaries {
			backup_ListBackupJobSummaries(cfg, client)
			return
		}
		if _backupListBackupJobs {
			backup_ListBackupJobs(cfg, client)
			return
		}
		if _backupListBackupPlanTemplates {
			backup_ListBackupPlanTemplates(cfg, client)
			return
		}
		if _backupListBackupPlanVersions {
			backup_ListBackupPlanVersions(cfg, client)
			return
		}
		if _backupListBackupPlans {
			backup_ListBackupPlans(cfg, client)
			return
		}
		if _backupListBackupSelections {
			backup_ListBackupSelections(cfg, client)
			return
		}
		if _backupListBackupVaults {
			backup_ListBackupVaults(cfg, client)
			return
		}
		if _backupListCopyJobSummaries {
			backup_ListCopyJobSummaries(cfg, client)
			return
		}
		if _backupListCopyJobs {
			backup_ListCopyJobs(cfg, client)
			return
		}
		if _backupListFrameworks {
			backup_ListFrameworks(cfg, client)
			return
		}
		if _backupListIndexedRecoveryPoints {
			backup_ListIndexedRecoveryPoints(cfg, client)
			return
		}
		if _backupListLegalHolds {
			backup_ListLegalHolds(cfg, client)
			return
		}
		if _backupListProtectedResources {
			backup_ListProtectedResources(cfg, client)
			return
		}
		if _backupListProtectedResourcesByBackupVault {
			backup_ListProtectedResourcesByBackupVault(cfg, client)
			return
		}
		if _backupListRecoveryPointsByBackupVault {
			backup_ListRecoveryPointsByBackupVault(cfg, client)
			return
		}
		if _backupListRecoveryPointsByLegalHold {
			backup_ListRecoveryPointsByLegalHold(cfg, client)
			return
		}
		if _backupListRecoveryPointsByResource {
			backup_ListRecoveryPointsByResource(cfg, client)
			return
		}
		if _backupListReportJobs {
			backup_ListReportJobs(cfg, client)
			return
		}
		if _backupListReportPlans {
			backup_ListReportPlans(cfg, client)
			return
		}
		if _backupListRestoreAccessBackupVaults {
			backup_ListRestoreAccessBackupVaults(cfg, client)
			return
		}
		if _backupListRestoreJobSummaries {
			backup_ListRestoreJobSummaries(cfg, client)
			return
		}
		if _backupListRestoreJobs {
			backup_ListRestoreJobs(cfg, client)
			return
		}
		if _backupListRestoreJobsByProtectedResource {
			backup_ListRestoreJobsByProtectedResource(cfg, client)
			return
		}
		if _backupListRestoreTestingPlans {
			backup_ListRestoreTestingPlans(cfg, client)
			return
		}
		if _backupListRestoreTestingSelections {
			backup_ListRestoreTestingSelections(cfg, client)
			return
		}
		if _backupListScanJobSummaries {
			backup_ListScanJobSummaries(cfg, client)
			return
		}
		if _backupListScanJobs {
			backup_ListScanJobs(cfg, client)
			return
		}
		if _backupListTags {
			backup_ListTags(cfg, client)
			return
		}
		if _backupListTieringConfigurations {
			backup_ListTieringConfigurations(cfg, client)
			return
		}
		if _backupPutBackupVaultAccessPolicy {
			backup_PutBackupVaultAccessPolicy(cfg, client)
			return
		}
		if _backupPutBackupVaultLockConfiguration {
			backup_PutBackupVaultLockConfiguration(cfg, client)
			return
		}
		if _backupPutBackupVaultNotifications {
			backup_PutBackupVaultNotifications(cfg, client)
			return
		}
		if _backupPutRestoreValidationResult {
			backup_PutRestoreValidationResult(cfg, client)
			return
		}
		if _backupRevokeRestoreAccessBackupVault {
			backup_RevokeRestoreAccessBackupVault(cfg, client)
			return
		}
		if _backupStartBackupJob {
			backup_StartBackupJob(cfg, client)
			return
		}
		if _backupStartCopyJob {
			backup_StartCopyJob(cfg, client)
			return
		}
		if _backupStartReportJob {
			backup_StartReportJob(cfg, client)
			return
		}
		if _backupStartRestoreJob {
			backup_StartRestoreJob(cfg, client)
			return
		}
		if _backupStartScanJob {
			backup_StartScanJob(cfg, client)
			return
		}
		if _backupStopBackupJob {
			backup_StopBackupJob(cfg, client)
			return
		}
		if _backupTagResource {
			backup_TagResource(cfg, client)
			return
		}
		if _backupUntagResource {
			backup_UntagResource(cfg, client)
			return
		}
		if _backupUpdateBackupPlan {
			backup_UpdateBackupPlan(cfg, client)
			return
		}
		if _backupUpdateFramework {
			backup_UpdateFramework(cfg, client)
			return
		}
		if _backupUpdateGlobalSettings {
			backup_UpdateGlobalSettings(cfg, client)
			return
		}
		if _backupUpdateRecoveryPointIndexSettings {
			backup_UpdateRecoveryPointIndexSettings(cfg, client)
			return
		}
		if _backupUpdateRecoveryPointLifecycle {
			backup_UpdateRecoveryPointLifecycle(cfg, client)
			return
		}
		if _backupUpdateRegionSettings {
			backup_UpdateRegionSettings(cfg, client)
			return
		}
		if _backupUpdateReportPlan {
			backup_UpdateReportPlan(cfg, client)
			return
		}
		if _backupUpdateRestoreTestingPlan {
			backup_UpdateRestoreTestingPlan(cfg, client)
			return
		}
		if _backupUpdateRestoreTestingSelection {
			backup_UpdateRestoreTestingSelection(cfg, client)
			return
		}
		if _backupUpdateTieringConfiguration {
			backup_UpdateTieringConfiguration(cfg, client)
			return
		}

	},
}

var (
	_backupAssociateBackupVaultMpaApprovalTeam    bool
	_backupCancelLegalHold                        bool
	_backupCreateBackupPlan                       bool
	_backupCreateBackupSelection                  bool
	_backupCreateBackupVault                      bool
	_backupCreateFramework                        bool
	_backupCreateLegalHold                        bool
	_backupCreateLogicallyAirGappedBackupVault    bool
	_backupCreateReportPlan                       bool
	_backupCreateRestoreAccessBackupVault         bool
	_backupCreateRestoreTestingPlan               bool
	_backupCreateRestoreTestingSelection          bool
	_backupCreateTieringConfiguration             bool
	_backupDeleteBackupPlan                       bool
	_backupDeleteBackupSelection                  bool
	_backupDeleteBackupVault                      bool
	_backupDeleteBackupVaultAccessPolicy          bool
	_backupDeleteBackupVaultLockConfiguration     bool
	_backupDeleteBackupVaultNotifications         bool
	_backupDeleteFramework                        bool
	_backupDeleteRecoveryPoint                    bool
	_backupDeleteReportPlan                       bool
	_backupDeleteRestoreTestingPlan               bool
	_backupDeleteRestoreTestingSelection          bool
	_backupDeleteTieringConfiguration             bool
	_backupDescribeBackupJob                      bool
	_backupDescribeBackupVault                    bool
	_backupDescribeCopyJob                        bool
	_backupDescribeFramework                      bool
	_backupDescribeGlobalSettings                 bool
	_backupDescribeProtectedResource              bool
	_backupDescribeRecoveryPoint                  bool
	_backupDescribeRegionSettings                 bool
	_backupDescribeReportJob                      bool
	_backupDescribeReportPlan                     bool
	_backupDescribeRestoreJob                     bool
	_backupDescribeScanJob                        bool
	_backupDisassociateBackupVaultMpaApprovalTeam bool
	_backupDisassociateRecoveryPoint              bool
	_backupDisassociateRecoveryPointFromParent    bool
	_backupExportBackupPlanTemplate               bool
	_backupGetBackupPlan                          bool
	_backupGetBackupPlanFromJSON                  bool
	_backupGetBackupPlanFromTemplate              bool
	_backupGetBackupSelection                     bool
	_backupGetBackupVaultAccessPolicy             bool
	_backupGetBackupVaultNotifications            bool
	_backupGetLegalHold                           bool
	_backupGetRecoveryPointIndexDetails           bool
	_backupGetRecoveryPointRestoreMetadata        bool
	_backupGetRestoreJobMetadata                  bool
	_backupGetRestoreTestingInferredMetadata      bool
	_backupGetRestoreTestingPlan                  bool
	_backupGetRestoreTestingSelection             bool
	_backupGetSupportedResourceTypes              bool
	_backupGetTieringConfiguration                bool
	_backupListBackupJobSummaries                 bool
	_backupListBackupJobs                         bool
	_backupListBackupPlanTemplates                bool
	_backupListBackupPlanVersions                 bool
	_backupListBackupPlans                        bool
	_backupListBackupSelections                   bool
	_backupListBackupVaults                       bool
	_backupListCopyJobSummaries                   bool
	_backupListCopyJobs                           bool
	_backupListFrameworks                         bool
	_backupListIndexedRecoveryPoints              bool
	_backupListLegalHolds                         bool
	_backupListProtectedResources                 bool
	_backupListProtectedResourcesByBackupVault    bool
	_backupListRecoveryPointsByBackupVault        bool
	_backupListRecoveryPointsByLegalHold          bool
	_backupListRecoveryPointsByResource           bool
	_backupListReportJobs                         bool
	_backupListReportPlans                        bool
	_backupListRestoreAccessBackupVaults          bool
	_backupListRestoreJobSummaries                bool
	_backupListRestoreJobs                        bool
	_backupListRestoreJobsByProtectedResource     bool
	_backupListRestoreTestingPlans                bool
	_backupListRestoreTestingSelections           bool
	_backupListScanJobSummaries                   bool
	_backupListScanJobs                           bool
	_backupListTags                               bool
	_backupListTieringConfigurations              bool
	_backupPutBackupVaultAccessPolicy             bool
	_backupPutBackupVaultLockConfiguration        bool
	_backupPutBackupVaultNotifications            bool
	_backupPutRestoreValidationResult             bool
	_backupRevokeRestoreAccessBackupVault         bool
	_backupStartBackupJob                         bool
	_backupStartCopyJob                           bool
	_backupStartReportJob                         bool
	_backupStartRestoreJob                        bool
	_backupStartScanJob                           bool
	_backupStopBackupJob                          bool
	_backupTagResource                            bool
	_backupUntagResource                          bool
	_backupUpdateBackupPlan                       bool
	_backupUpdateFramework                        bool
	_backupUpdateGlobalSettings                   bool
	_backupUpdateRecoveryPointIndexSettings       bool
	_backupUpdateRecoveryPointLifecycle           bool
	_backupUpdateRegionSettings                   bool
	_backupUpdateReportPlan                       bool
	_backupUpdateRestoreTestingPlan               bool
	_backupUpdateRestoreTestingSelection          bool
	_backupUpdateTieringConfiguration             bool

	_backupAccountId                         string
	_backupAggregationPeriod                 string
	_backupBackupJobId                       string
	_backupBackupOptions                     string
	_backupBackupPlan                        string
	_backupBackupPlanId                      string
	_backupBackupPlanTags                    string
	_backupBackupPlanTemplateId              string
	_backupBackupPlanTemplateJson            string
	_backupBackupSelection                   string
	_backupBackupVaultAccountId              string
	_backupBackupVaultEvents                 string
	_backupBackupVaultName                   string
	_backupBackupVaultTags                   string
	_backupByAccountId                       string
	_backupByBackupPlanId                    string
	_backupByBackupVaultName                 string
	_backupByCompleteAfter                   string
	_backupByCompleteBefore                  string
	_backupByCreatedAfter                    string
	_backupByCreatedBefore                   string
	_backupByCreationAfter                   string
	_backupByCreationBefore                  string
	_backupByDestinationVaultArn             string
	_backupByMalwareScanner                  string
	_backupByMessageCategory                 string
	_backupByParentJobId                     string
	_backupByParentRecoveryPointArn          string
	_backupByRecoveryPointArn                string
	_backupByRecoveryPointCreationDateAfter  string
	_backupByRecoveryPointCreationDateBefore string
	_backupByReportPlanName                  string
	_backupByResourceArn                     string
	_backupByResourceType                    string
	_backupByRestoreTestingPlanArn           string
	_backupByScanResultStatus                string
	_backupByShared                          string
	_backupBySourceRecoveryPointArn          string
	_backupByState                           string
	_backupByStatus                          string
	_backupByVaultType                       string
	_backupCancelDescription                 string
	_backupChangeableForDays                 string
	_backupCompleteWindowMinutes             string
	_backupCopyJobId                         string
	_backupCopySourceTagsToRestoredResource  string
	_backupCreatedAfter                      string
	_backupCreatedBefore                     string
	_backupCreatorRequestId                  string
	_backupDescription                       string
	_backupDestinationBackupVaultArn         string
	_backupEncryptionKeyArn                  string
	_backupFrameworkControls                 string
	_backupFrameworkDescription              string
	_backupFrameworkName                     string
	_backupFrameworkTags                     string
	_backupGlobalSettings                    string
	_backupIamRoleArn                        string
	_backupIdempotencyToken                  string
	_backupIncludeDeleted                    string
	_backupIndex                             string
	_backupIndexStatus                       string
	_backupLegalHoldId                       string
	_backupLifecycle                         string
	_backupLogicallyAirGappedBackupVaultArn  string
	_backupMalwareScanner                    string
	_backupManagedByAWSBackupOnly            string
	_backupMaxResults                        string
	_backupMaxRetentionDays                  string
	_backupMaxScheduledRunsPreview           string
	_backupMessageCategory                   string
	_backupMetadata                          string
	_backupMinRetentionDays                  string
	_backupMpaApprovalTeamArn                string
	_backupNextToken                         string
	_backupPolicy                            string
	_backupRecoveryPointArn                  string
	_backupRecoveryPointSelection            string
	_backupRecoveryPointTags                 string
	_backupReportDeliveryChannel             string
	_backupReportJobId                       string
	_backupReportPlanDescription             string
	_backupReportPlanName                    string
	_backupReportPlanTags                    string
	_backupReportSetting                     string
	_backupRequesterComment                  string
	_backupResourceArn                       string
	_backupResourceType                      string
	_backupResourceTypeManagementPreference  string
	_backupResourceTypeOptInPreference       string
	_backupRestoreAccessBackupVaultArn       string
	_backupRestoreJobId                      string
	_backupRestoreTestingPlan                string
	_backupRestoreTestingPlanName            string
	_backupRestoreTestingSelection           string
	_backupRestoreTestingSelectionName       string
	_backupRetainRecordInDays                string
	_backupScanBaseRecoveryPointArn          string
	_backupScanJobId                         string
	_backupScanMode                          string
	_backupScanResultStatus                  string
	_backupScannerRoleArn                    string
	_backupSelectionId                       string
	_backupSNSTopicArn                       string
	_backupSourceBackupVaultArn              string
	_backupSourceBackupVaultName             string
	_backupSourceResourceArn                 string
	_backupStartWindowMinutes                string
	_backupState                             string
	_backupTagKeyList                        []string
	_backupTags                              string
	_backupTieringConfiguration              string
	_backupTieringConfigurationName          string
	_backupTieringConfigurationTags          string
	_backupTitle                             string
	_backupValidationStatus                  string
	_backupValidationStatusMessage           string
	_backupVersionId                         string
)

// Associates an MPA approval team with a backup vault.
func backup_AssociateBackupVaultMpaApprovalTeam(cfg aws.Config, client *backup.Client) {
	input := &backup.AssociateBackupVaultMpaApprovalTeamInput{
		// BackupVaultName: *string, // Required
		// MpaApprovalTeamArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupMpaApprovalTeamArn) > 0 {
		input.MpaApprovalTeamArn = aws.String(_backupMpaApprovalTeamArn)
	}
	if len(_backupRequesterComment) > 0 {
		input.RequesterComment = aws.String(_backupRequesterComment)
	}

	if resp, err := client.AssociateBackupVaultMpaApprovalTeam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified legal hold on a recovery point. This action can only be
// performed by a user with sufficient permissions.
func backup_CancelLegalHold(cfg aws.Config, client *backup.Client) {
	input := &backup.CancelLegalHoldInput{
		// CancelDescription: *string, // Required
		// LegalHoldId: *string, // Required
	}

	if len(_backupCancelDescription) > 0 {
		input.CancelDescription = aws.String(_backupCancelDescription)
	}
	if len(_backupLegalHoldId) > 0 {
		input.LegalHoldId = aws.String(_backupLegalHoldId)
	}
	if len(_backupRetainRecordInDays) > 0 {
		if err := assignInputField(input, "RetainRecordInDays", _backupRetainRecordInDays); err != nil {
			log.Errorf("invalid --retain-record-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelLegalHold(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a backup plan using a backup plan name and backup rules. A backup plan
// is a document that contains information that Backup uses to schedule tasks that
// create recovery points for resources.
//
// If you call CreateBackupPlan with a plan that already exists, you receive an
// AlreadyExistsException exception.
func backup_CreateBackupPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateBackupPlanInput{
		// BackupPlan: *types.BackupPlanInput, // Required
	}

	if len(_backupBackupPlan) > 0 {
		if err := assignInputField(input, "BackupPlan", _backupBackupPlan); err != nil {
			log.Errorf("invalid --backup-plan: %s", err.Error())
			return
		}
	}
	if len(_backupBackupPlanTags) > 0 {
		if err := assignInputField(input, "BackupPlanTags", _backupBackupPlanTags); err != nil {
			log.Errorf("invalid --backup-plan-tags: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}

	if resp, err := client.CreateBackupPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a JSON document that specifies a set of resources to assign to a backup
// plan. For examples, see [Assigning resources programmatically].
//
// [Assigning resources programmatically]: https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-json
func backup_CreateBackupSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateBackupSelectionInput{
		// BackupPlanId: *string, // Required
		// BackupSelection: *types.BackupSelection, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}
	if len(_backupBackupSelection) > 0 {
		if err := assignInputField(input, "BackupSelection", _backupBackupSelection); err != nil {
			log.Errorf("invalid --backup-selection: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}

	if resp, err := client.CreateBackupSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a logical container where backups are stored. A CreateBackupVault
// request includes a name, optionally one or more resource tags, an encryption
// key, and a request ID.
//
// Do not include sensitive data, such as passport numbers, in the name of a
// backup vault.
func backup_CreateBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateBackupVaultInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupBackupVaultTags) > 0 {
		if err := assignInputField(input, "BackupVaultTags", _backupBackupVaultTags); err != nil {
			log.Errorf("invalid --backup-vault-tags: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}
	if len(_backupEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_backupEncryptionKeyArn)
	}

	if resp, err := client.CreateBackupVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a framework with one or more controls. A framework is a collection of
// controls that you can use to evaluate your backup practices. By using pre-built
// customizable controls to define your policies, you can evaluate whether your
// backup practices comply with your policies and which resources are not yet in
// compliance.
func backup_CreateFramework(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateFrameworkInput{
		// FrameworkControls: []types.FrameworkControl, // Required
		// FrameworkName: *string, // Required
	}

	if len(_backupFrameworkControls) > 0 {
		if err := assignInputField(input, "FrameworkControls", _backupFrameworkControls); err != nil {
			log.Errorf("invalid --framework-controls: %s", err.Error())
			return
		}
	}
	if len(_backupFrameworkName) > 0 {
		input.FrameworkName = aws.String(_backupFrameworkName)
	}
	if len(_backupFrameworkDescription) > 0 {
		input.FrameworkDescription = aws.String(_backupFrameworkDescription)
	}
	if len(_backupFrameworkTags) > 0 {
		if err := assignInputField(input, "FrameworkTags", _backupFrameworkTags); err != nil {
			log.Errorf("invalid --framework-tags: %s", err.Error())
			return
		}
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}

	if resp, err := client.CreateFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a legal hold on a recovery point (backup). A legal hold is a restraint
// on altering or deleting a backup until an authorized user cancels the legal
// hold. Any actions to delete or disassociate a recovery point will fail with an
// error if one or more active legal holds are on the recovery point.
func backup_CreateLegalHold(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateLegalHoldInput{
		// Description: *string, // Required
		// Title: *string, // Required
	}

	if len(_backupDescription) > 0 {
		input.Description = aws.String(_backupDescription)
	}
	if len(_backupTitle) > 0 {
		input.Title = aws.String(_backupTitle)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupRecoveryPointSelection) > 0 {
		if err := assignInputField(input, "RecoveryPointSelection", _backupRecoveryPointSelection); err != nil {
			log.Errorf("invalid --recovery-point-selection: %s", err.Error())
			return
		}
	}
	if len(_backupTags) > 0 {
		if err := assignInputField(input, "Tags", _backupTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLegalHold(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a logical container to where backups may be copied.
// This request includes a name, the Region, the maximum number of retention days,
// the minimum number of retention days, and optionally can include tags and a
// creator request ID.
//
// Do not include sensitive data, such as passport numbers, in the name of a
// backup vault.
func backup_CreateLogicallyAirGappedBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateLogicallyAirGappedBackupVaultInput{
		// BackupVaultName: *string, // Required
		// MaxRetentionDays: *int64, // Required
		// MinRetentionDays: *int64, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupMaxRetentionDays) > 0 {
		if err := assignInputField(input, "MaxRetentionDays", _backupMaxRetentionDays); err != nil {
			log.Errorf("invalid --max-retention-days: %s", err.Error())
			return
		}
	}
	if len(_backupMinRetentionDays) > 0 {
		if err := assignInputField(input, "MinRetentionDays", _backupMinRetentionDays); err != nil {
			log.Errorf("invalid --min-retention-days: %s", err.Error())
			return
		}
	}
	if len(_backupBackupVaultTags) > 0 {
		if err := assignInputField(input, "BackupVaultTags", _backupBackupVaultTags); err != nil {
			log.Errorf("invalid --backup-vault-tags: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}
	if len(_backupEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_backupEncryptionKeyArn)
	}

	if resp, err := client.CreateLogicallyAirGappedBackupVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a report plan. A report plan is a document that contains information
// about the contents of the report and where Backup will deliver it.
//
// If you call CreateReportPlan with a plan that already exists, you receive an
// AlreadyExistsException exception.
func backup_CreateReportPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateReportPlanInput{
		// ReportDeliveryChannel: *types.ReportDeliveryChannel, // Required
		// ReportPlanName: *string, // Required
		// ReportSetting: *types.ReportSetting, // Required
	}

	if len(_backupReportDeliveryChannel) > 0 {
		if err := assignInputField(input, "ReportDeliveryChannel", _backupReportDeliveryChannel); err != nil {
			log.Errorf("invalid --report-delivery-channel: %s", err.Error())
			return
		}
	}
	if len(_backupReportPlanName) > 0 {
		input.ReportPlanName = aws.String(_backupReportPlanName)
	}
	if len(_backupReportSetting) > 0 {
		if err := assignInputField(input, "ReportSetting", _backupReportSetting); err != nil {
			log.Errorf("invalid --report-setting: %s", err.Error())
			return
		}
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupReportPlanDescription) > 0 {
		input.ReportPlanDescription = aws.String(_backupReportPlanDescription)
	}
	if len(_backupReportPlanTags) > 0 {
		if err := assignInputField(input, "ReportPlanTags", _backupReportPlanTags); err != nil {
			log.Errorf("invalid --report-plan-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReportPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a restore access backup vault that provides temporary access to
// recovery points in a logically air-gapped backup vault, subject to MPA approval.
func backup_CreateRestoreAccessBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateRestoreAccessBackupVaultInput{
		// SourceBackupVaultArn: *string, // Required
	}

	if len(_backupSourceBackupVaultArn) > 0 {
		input.SourceBackupVaultArn = aws.String(_backupSourceBackupVaultArn)
	}
	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupBackupVaultTags) > 0 {
		if err := assignInputField(input, "BackupVaultTags", _backupBackupVaultTags); err != nil {
			log.Errorf("invalid --backup-vault-tags: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}
	if len(_backupRequesterComment) > 0 {
		input.RequesterComment = aws.String(_backupRequesterComment)
	}

	if resp, err := client.CreateRestoreAccessBackupVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a restore testing plan.
// The first of two steps to create a restore testing plan. After this request is
// successful, finish the procedure using CreateRestoreTestingSelection.
func backup_CreateRestoreTestingPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateRestoreTestingPlanInput{
		// RestoreTestingPlan: *types.RestoreTestingPlanForCreate, // Required
	}

	if len(_backupRestoreTestingPlan) > 0 {
		if err := assignInputField(input, "RestoreTestingPlan", _backupRestoreTestingPlan); err != nil {
			log.Errorf("invalid --restore-testing-plan: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}
	if len(_backupTags) > 0 {
		if err := assignInputField(input, "Tags", _backupTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRestoreTestingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request can be sent after CreateRestoreTestingPlan request returns
// successfully. This is the second part of creating a resource testing plan, and
// it must be completed sequentially.
//
// This consists of RestoreTestingSelectionName , ProtectedResourceType , and one
// of the following:
//
// - ProtectedResourceArns
//
// - ProtectedResourceConditions
//
// Each protected resource type can have one single value.
//
// A restore testing selection can include a wildcard value ("*") for
// ProtectedResourceArns along with ProtectedResourceConditions . Alternatively,
// you can include up to 30 specific protected resource ARNs in
// ProtectedResourceArns .
//
// Cannot select by both protected resource types AND specific ARNs. Request will
// fail if both are included.
func backup_CreateRestoreTestingSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateRestoreTestingSelectionInput{
		// RestoreTestingPlanName: *string, // Required
		// RestoreTestingSelection: *types.RestoreTestingSelectionForCreate, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}
	if len(_backupRestoreTestingSelection) > 0 {
		if err := assignInputField(input, "RestoreTestingSelection", _backupRestoreTestingSelection); err != nil {
			log.Errorf("invalid --restore-testing-selection: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}

	if resp, err := client.CreateRestoreTestingSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a tiering configuration.
// A tiering configuration enables automatic movement of backup data to a
// lower-cost storage tier based on the age of backed-up objects in the backup
// vault.
//
// Each vault can only have one vault-specific tiering configuration, in addition
// to any global configuration that applies to all vaults.
func backup_CreateTieringConfiguration(cfg aws.Config, client *backup.Client) {
	input := &backup.CreateTieringConfigurationInput{
		// TieringConfiguration: *types.TieringConfigurationInputForCreate, // Required
	}

	if len(_backupTieringConfiguration) > 0 {
		if err := assignInputField(input, "TieringConfiguration", _backupTieringConfiguration); err != nil {
			log.Errorf("invalid --tiering-configuration: %s", err.Error())
			return
		}
	}
	if len(_backupCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_backupCreatorRequestId)
	}
	if len(_backupTieringConfigurationTags) > 0 {
		if err := assignInputField(input, "TieringConfigurationTags", _backupTieringConfigurationTags); err != nil {
			log.Errorf("invalid --tiering-configuration-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a backup plan. A backup plan can only be deleted after all associated
// selections of resources have been deleted. Deleting a backup plan deletes the
// current version of a backup plan. Previous versions, if any, will still exist.
func backup_DeleteBackupPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteBackupPlanInput{
		// BackupPlanId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}

	if resp, err := client.DeleteBackupPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource selection associated with a backup plan that is specified
// by the SelectionId .
func backup_DeleteBackupSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteBackupSelectionInput{
		// BackupPlanId: *string, // Required
		// SelectionId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}
	if len(_backupSelectionId) > 0 {
		input.SelectionId = aws.String(_backupSelectionId)
	}

	if resp, err := client.DeleteBackupSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the backup vault identified by its name. A vault can be deleted only if
// it is empty.
func backup_DeleteBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteBackupVaultInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}

	if resp, err := client.DeleteBackupVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the policy document that manages permissions on a backup vault.
func backup_DeleteBackupVaultAccessPolicy(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteBackupVaultAccessPolicyInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}

	if resp, err := client.DeleteBackupVaultAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes Backup Vault Lock from a backup vault specified by a backup vault name.
// If the Vault Lock configuration is immutable, then you cannot delete Vault Lock
// using API operations, and you will receive an InvalidRequestException if you
// attempt to do so. For more information, see [Vault Lock]in the Backup Developer Guide.
//
// [Vault Lock]: https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html
func backup_DeleteBackupVaultLockConfiguration(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteBackupVaultLockConfigurationInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}

	if resp, err := client.DeleteBackupVaultLockConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes event notifications for the specified backup vault.
func backup_DeleteBackupVaultNotifications(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteBackupVaultNotificationsInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}

	if resp, err := client.DeleteBackupVaultNotifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the framework specified by a framework name.
func backup_DeleteFramework(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteFrameworkInput{
		// FrameworkName: *string, // Required
	}

	if len(_backupFrameworkName) > 0 {
		input.FrameworkName = aws.String(_backupFrameworkName)
	}

	if resp, err := client.DeleteFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the recovery point specified by a recovery point ID.
// If the recovery point ID belongs to a continuous backup, calling this endpoint
// deletes the existing continuous backup and stops future continuous backup.
//
// When an IAM role's permissions are insufficient to call this API, the service
// sends back an HTTP 200 response with an empty HTTP body, but the recovery point
// is not deleted. Instead, it enters an EXPIRED state.
//
// EXPIRED recovery points can be deleted with this API once the IAM role has the
// iam:CreateServiceLinkedRole action. To learn more about adding this role, see [Troubleshooting manual deletions].
//
// If the user or role is deleted or the permission within the role is removed,
// the deletion will not be successful and will enter an EXPIRED state.
//
// [Troubleshooting manual deletions]: https://docs.aws.amazon.com/aws-backup/latest/devguide/deleting-backups.html#deleting-backups-troubleshooting
func backup_DeleteRecoveryPoint(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteRecoveryPointInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}

	if resp, err := client.DeleteRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the report plan specified by a report plan name.
func backup_DeleteReportPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteReportPlanInput{
		// ReportPlanName: *string, // Required
	}

	if len(_backupReportPlanName) > 0 {
		input.ReportPlanName = aws.String(_backupReportPlanName)
	}

	if resp, err := client.DeleteReportPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request deletes the specified restore testing plan.
// Deletion can only successfully occur if all associated restore testing
// selections are deleted first.
func backup_DeleteRestoreTestingPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteRestoreTestingPlanInput{
		// RestoreTestingPlanName: *string, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}

	if resp, err := client.DeleteRestoreTestingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Input the Restore Testing Plan name and Restore Testing Selection name.
// All testing selections associated with a restore testing plan must be deleted
// before the restore testing plan can be deleted.
func backup_DeleteRestoreTestingSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteRestoreTestingSelectionInput{
		// RestoreTestingPlanName: *string, // Required
		// RestoreTestingSelectionName: *string, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}
	if len(_backupRestoreTestingSelectionName) > 0 {
		input.RestoreTestingSelectionName = aws.String(_backupRestoreTestingSelectionName)
	}

	if resp, err := client.DeleteRestoreTestingSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the tiering configuration specified by a tiering configuration name.
func backup_DeleteTieringConfiguration(cfg aws.Config, client *backup.Client) {
	input := &backup.DeleteTieringConfigurationInput{
		// TieringConfigurationName: *string, // Required
	}

	if len(_backupTieringConfigurationName) > 0 {
		input.TieringConfigurationName = aws.String(_backupTieringConfigurationName)
	}

	if resp, err := client.DeleteTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns backup job details for the specified BackupJobId .
func backup_DescribeBackupJob(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeBackupJobInput{
		// BackupJobId: *string, // Required
	}

	if len(_backupBackupJobId) > 0 {
		input.BackupJobId = aws.String(_backupBackupJobId)
	}

	if resp, err := client.DescribeBackupJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata about a backup vault specified by its name.
func backup_DescribeBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeBackupVaultInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupBackupVaultAccountId) > 0 {
		input.BackupVaultAccountId = aws.String(_backupBackupVaultAccountId)
	}

	if resp, err := client.DescribeBackupVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata associated with creating a copy of a resource.
func backup_DescribeCopyJob(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeCopyJobInput{
		// CopyJobId: *string, // Required
	}

	if len(_backupCopyJobId) > 0 {
		input.CopyJobId = aws.String(_backupCopyJobId)
	}

	if resp, err := client.DescribeCopyJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the framework details for the specified FrameworkName .
func backup_DescribeFramework(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeFrameworkInput{
		// FrameworkName: *string, // Required
	}

	if len(_backupFrameworkName) > 0 {
		input.FrameworkName = aws.String(_backupFrameworkName)
	}

	if resp, err := client.DescribeFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes whether the Amazon Web Services account is opted in to cross-account
// backup. Returns an error if the account is not a member of an Organizations
// organization. Example: describe-global-settings --region us-west-2
func backup_DescribeGlobalSettings(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeGlobalSettingsInput{}

	if resp, err := client.DescribeGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a saved resource, including the last time it was
// backed up, its Amazon Resource Name (ARN), and the Amazon Web Services service
// type of the saved resource.
func backup_DescribeProtectedResource(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeProtectedResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}

	if resp, err := client.DescribeProtectedResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata associated with a recovery point, including ID, status,
// encryption, and lifecycle.
func backup_DescribeRecoveryPoint(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeRecoveryPointInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupBackupVaultAccountId) > 0 {
		input.BackupVaultAccountId = aws.String(_backupBackupVaultAccountId)
	}

	if resp, err := client.DescribeRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current service opt-in settings for the Region. If service opt-in
// is enabled for a service, Backup tries to protect that service's resources in
// this Region, when the resource is included in an on-demand backup or scheduled
// backup plan. Otherwise, Backup does not try to protect that service's resources
// in this Region.
func backup_DescribeRegionSettings(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeRegionSettingsInput{}

	if resp, err := client.DescribeRegionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details associated with creating a report as specified by its
// ReportJobId .
func backup_DescribeReportJob(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeReportJobInput{
		// ReportJobId: *string, // Required
	}

	if len(_backupReportJobId) > 0 {
		input.ReportJobId = aws.String(_backupReportJobId)
	}

	if resp, err := client.DescribeReportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all report plans for an Amazon Web Services account and
// Amazon Web Services Region.
func backup_DescribeReportPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeReportPlanInput{
		// ReportPlanName: *string, // Required
	}

	if len(_backupReportPlanName) > 0 {
		input.ReportPlanName = aws.String(_backupReportPlanName)
	}

	if resp, err := client.DescribeReportPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata associated with a restore job that is specified by a job ID.
func backup_DescribeRestoreJob(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeRestoreJobInput{
		// RestoreJobId: *string, // Required
	}

	if len(_backupRestoreJobId) > 0 {
		input.RestoreJobId = aws.String(_backupRestoreJobId)
	}

	if resp, err := client.DescribeRestoreJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns scan job details for the specified ScanJobID.
func backup_DescribeScanJob(cfg aws.Config, client *backup.Client) {
	input := &backup.DescribeScanJobInput{
		// ScanJobId: *string, // Required
	}

	if len(_backupScanJobId) > 0 {
		input.ScanJobId = aws.String(_backupScanJobId)
	}

	if resp, err := client.DescribeScanJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between an MPA approval team and a backup vault,
// disabling the MPA approval workflow for restore operations.
func backup_DisassociateBackupVaultMpaApprovalTeam(cfg aws.Config, client *backup.Client) {
	input := &backup.DisassociateBackupVaultMpaApprovalTeamInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRequesterComment) > 0 {
		input.RequesterComment = aws.String(_backupRequesterComment)
	}

	if resp, err := client.DisassociateBackupVaultMpaApprovalTeam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified continuous backup recovery point from Backup and releases
// control of that continuous backup to the source service, such as Amazon RDS. The
// source service will continue to create and retain continuous backups using the
// lifecycle that you specified in your original backup plan.
//
// Does not support snapshot backup recovery points.
func backup_DisassociateRecoveryPoint(cfg aws.Config, client *backup.Client) {
	input := &backup.DisassociateRecoveryPointInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}

	if resp, err := client.DisassociateRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action to a specific child (nested) recovery point removes the
// relationship between the specified recovery point and its parent (composite)
// recovery point.
func backup_DisassociateRecoveryPointFromParent(cfg aws.Config, client *backup.Client) {
	input := &backup.DisassociateRecoveryPointFromParentInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}

	if resp, err := client.DisassociateRecoveryPointFromParent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the backup plan that is specified by the plan ID as a backup template.
func backup_ExportBackupPlanTemplate(cfg aws.Config, client *backup.Client) {
	input := &backup.ExportBackupPlanTemplateInput{
		// BackupPlanId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}

	if resp, err := client.ExportBackupPlanTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns BackupPlan details for the specified BackupPlanId . The details are the
// body of a backup plan in JSON format, in addition to plan metadata.
func backup_GetBackupPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.GetBackupPlanInput{
		// BackupPlanId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}
	if len(_backupMaxScheduledRunsPreview) > 0 {
		if err := assignInputField(input, "MaxScheduledRunsPreview", _backupMaxScheduledRunsPreview); err != nil {
			log.Errorf("invalid --max-scheduled-runs-preview: %s", err.Error())
			return
		}
	}
	if len(_backupVersionId) > 0 {
		input.VersionId = aws.String(_backupVersionId)
	}

	if resp, err := client.GetBackupPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a valid JSON document specifying a backup plan or an error.
func backup_GetBackupPlanFromJSON(cfg aws.Config, client *backup.Client) {
	input := &backup.GetBackupPlanFromJSONInput{
		// BackupPlanTemplateJson: *string, // Required
	}

	if len(_backupBackupPlanTemplateJson) > 0 {
		input.BackupPlanTemplateJson = aws.String(_backupBackupPlanTemplateJson)
	}

	if resp, err := client.GetBackupPlanFromJSON(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the template specified by its templateId as a backup plan.
func backup_GetBackupPlanFromTemplate(cfg aws.Config, client *backup.Client) {
	input := &backup.GetBackupPlanFromTemplateInput{
		// BackupPlanTemplateId: *string, // Required
	}

	if len(_backupBackupPlanTemplateId) > 0 {
		input.BackupPlanTemplateId = aws.String(_backupBackupPlanTemplateId)
	}

	if resp, err := client.GetBackupPlanFromTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns selection metadata and a document in JSON format that specifies a list
// of resources that are associated with a backup plan.
func backup_GetBackupSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.GetBackupSelectionInput{
		// BackupPlanId: *string, // Required
		// SelectionId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}
	if len(_backupSelectionId) > 0 {
		input.SelectionId = aws.String(_backupSelectionId)
	}

	if resp, err := client.GetBackupSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the access policy document that is associated with the named backup
// vault.
func backup_GetBackupVaultAccessPolicy(cfg aws.Config, client *backup.Client) {
	input := &backup.GetBackupVaultAccessPolicyInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}

	if resp, err := client.GetBackupVaultAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns event notifications for the specified backup vault.
func backup_GetBackupVaultNotifications(cfg aws.Config, client *backup.Client) {
	input := &backup.GetBackupVaultNotificationsInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}

	if resp, err := client.GetBackupVaultNotifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action returns details for a specified legal hold. The details are the
// body of a legal hold in JSON format, in addition to metadata.
func backup_GetLegalHold(cfg aws.Config, client *backup.Client) {
	input := &backup.GetLegalHoldInput{
		// LegalHoldId: *string, // Required
	}

	if len(_backupLegalHoldId) > 0 {
		input.LegalHoldId = aws.String(_backupLegalHoldId)
	}

	if resp, err := client.GetLegalHold(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns the metadata and details specific to the backup index
// associated with the specified recovery point.
func backup_GetRecoveryPointIndexDetails(cfg aws.Config, client *backup.Client) {
	input := &backup.GetRecoveryPointIndexDetailsInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}

	if resp, err := client.GetRecoveryPointIndexDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of metadata key-value pairs that were used to create the backup.
func backup_GetRecoveryPointRestoreMetadata(cfg aws.Config, client *backup.Client) {
	input := &backup.GetRecoveryPointRestoreMetadataInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupBackupVaultAccountId) > 0 {
		input.BackupVaultAccountId = aws.String(_backupBackupVaultAccountId)
	}

	if resp, err := client.GetRecoveryPointRestoreMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request returns the metadata for the specified restore job.
func backup_GetRestoreJobMetadata(cfg aws.Config, client *backup.Client) {
	input := &backup.GetRestoreJobMetadataInput{
		// RestoreJobId: *string, // Required
	}

	if len(_backupRestoreJobId) > 0 {
		input.RestoreJobId = aws.String(_backupRestoreJobId)
	}

	if resp, err := client.GetRestoreJobMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request returns the minimal required set of metadata needed to start a
// restore job with secure default settings. BackupVaultName and RecoveryPointArn
// are required parameters. BackupVaultAccountId is an optional parameter.
func backup_GetRestoreTestingInferredMetadata(cfg aws.Config, client *backup.Client) {
	input := &backup.GetRestoreTestingInferredMetadataInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupBackupVaultAccountId) > 0 {
		input.BackupVaultAccountId = aws.String(_backupBackupVaultAccountId)
	}

	if resp, err := client.GetRestoreTestingInferredMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns RestoreTestingPlan details for the specified RestoreTestingPlanName .
// The details are the body of a restore testing plan in JSON format, in addition
// to plan metadata.
func backup_GetRestoreTestingPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.GetRestoreTestingPlanInput{
		// RestoreTestingPlanName: *string, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}

	if resp, err := client.GetRestoreTestingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns RestoreTestingSelection, which displays resources and elements of the
// restore testing plan.
func backup_GetRestoreTestingSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.GetRestoreTestingSelectionInput{
		// RestoreTestingPlanName: *string, // Required
		// RestoreTestingSelectionName: *string, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}
	if len(_backupRestoreTestingSelectionName) > 0 {
		input.RestoreTestingSelectionName = aws.String(_backupRestoreTestingSelectionName)
	}

	if resp, err := client.GetRestoreTestingSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the Amazon Web Services resource types supported by Backup.
func backup_GetSupportedResourceTypes(cfg aws.Config, client *backup.Client) {
	input := &backup.GetSupportedResourceTypesInput{}

	if resp, err := client.GetSupportedResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns TieringConfiguration details for the specified TieringConfigurationName
// . The details are the body of a tiering configuration in JSON format, in
// addition to configuration metadata.
func backup_GetTieringConfiguration(cfg aws.Config, client *backup.Client) {
	input := &backup.GetTieringConfigurationInput{
		// TieringConfigurationName: *string, // Required
	}

	if len(_backupTieringConfigurationName) > 0 {
		input.TieringConfigurationName = aws.String(_backupTieringConfigurationName)
	}

	if resp, err := client.GetTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is a request for a summary of backup jobs created or running within the
// most recent 30 days. You can include parameters AccountID, State, ResourceType,
// MessageCategory, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// ResourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
func backup_ListBackupJobSummaries(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupJobSummariesInput{}

	if len(_backupAccountId) > 0 {
		input.AccountId = aws.String(_backupAccountId)
	}
	if len(_backupAggregationPeriod) > 0 {
		if err := assignInputField(input, "AggregationPeriod", _backupAggregationPeriod); err != nil {
			log.Errorf("invalid --aggregation-period: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupMessageCategory) > 0 {
		input.MessageCategory = aws.String(_backupMessageCategory)
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}
	if len(_backupResourceType) > 0 {
		input.ResourceType = aws.String(_backupResourceType)
	}
	if len(_backupState) > 0 {
		if err := assignInputField(input, "State", _backupState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBackupJobSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupJobSummariesOutput
	p := backup.NewListBackupJobSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of existing backup jobs for an authenticated account for the
// last 30 days. For a longer period of time, consider using these [monitoring tools].
//
// [monitoring tools]: https://docs.aws.amazon.com/aws-backup/latest/devguide/monitoring.html
func backup_ListBackupJobs(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupJobsInput{}

	if len(_backupByAccountId) > 0 {
		input.ByAccountId = aws.String(_backupByAccountId)
	}
	if len(_backupByBackupVaultName) > 0 {
		input.ByBackupVaultName = aws.String(_backupByBackupVaultName)
	}
	if len(_backupByCompleteAfter) > 0 {
		if err := assignInputField(input, "ByCompleteAfter", _backupByCompleteAfter); err != nil {
			log.Errorf("invalid --by-complete-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCompleteBefore) > 0 {
		if err := assignInputField(input, "ByCompleteBefore", _backupByCompleteBefore); err != nil {
			log.Errorf("invalid --by-complete-before: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedAfter) > 0 {
		if err := assignInputField(input, "ByCreatedAfter", _backupByCreatedAfter); err != nil {
			log.Errorf("invalid --by-created-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedBefore) > 0 {
		if err := assignInputField(input, "ByCreatedBefore", _backupByCreatedBefore); err != nil {
			log.Errorf("invalid --by-created-before: %s", err.Error())
			return
		}
	}
	if len(_backupByMessageCategory) > 0 {
		input.ByMessageCategory = aws.String(_backupByMessageCategory)
	}
	if len(_backupByParentJobId) > 0 {
		input.ByParentJobId = aws.String(_backupByParentJobId)
	}
	if len(_backupByResourceArn) > 0 {
		input.ByResourceArn = aws.String(_backupByResourceArn)
	}
	if len(_backupByResourceType) > 0 {
		input.ByResourceType = aws.String(_backupByResourceType)
	}
	if len(_backupByState) > 0 {
		if err := assignInputField(input, "ByState", _backupByState); err != nil {
			log.Errorf("invalid --by-state: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBackupJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupJobsOutput
	p := backup.NewListBackupJobsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the backup plan templates.
func backup_ListBackupPlanTemplates(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupPlanTemplatesInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBackupPlanTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupPlanTemplatesOutput
	p := backup.NewListBackupPlanTemplatesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns version metadata of your backup plans, including Amazon Resource Names
// (ARNs), backup plan IDs, creation and deletion dates, plan names, and version
// IDs.
func backup_ListBackupPlanVersions(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupPlanVersionsInput{
		// BackupPlanId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBackupPlanVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupPlanVersionsOutput
	p := backup.NewListBackupPlanVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the active backup plans for the account.
func backup_ListBackupPlans(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupPlansInput{}

	if len(_backupIncludeDeleted) > 0 {
		if err := assignInputField(input, "IncludeDeleted", _backupIncludeDeleted); err != nil {
			log.Errorf("invalid --include-deleted: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBackupPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupPlansOutput
	p := backup.NewListBackupPlansPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns an array containing metadata of the resources associated with the
// target backup plan.
func backup_ListBackupSelections(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupSelectionsInput{
		// BackupPlanId: *string, // Required
	}

	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBackupSelections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupSelectionsOutput
	p := backup.NewListBackupSelectionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of recovery point storage containers along with information
// about them.
func backup_ListBackupVaults(cfg aws.Config, client *backup.Client) {
	input := &backup.ListBackupVaultsInput{}

	if len(_backupByShared) > 0 {
		if err := assignInputField(input, "ByShared", _backupByShared); err != nil {
			log.Errorf("invalid --by-shared: %s", err.Error())
			return
		}
	}
	if len(_backupByVaultType) > 0 {
		if err := assignInputField(input, "ByVaultType", _backupByVaultType); err != nil {
			log.Errorf("invalid --by-vault-type: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBackupVaults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListBackupVaultsOutput
	p := backup.NewListBackupVaultsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This request obtains a list of copy jobs created or running within the the most
// recent 30 days. You can include parameters AccountID, State, ResourceType,
// MessageCategory, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// RestourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
func backup_ListCopyJobSummaries(cfg aws.Config, client *backup.Client) {
	input := &backup.ListCopyJobSummariesInput{}

	if len(_backupAccountId) > 0 {
		input.AccountId = aws.String(_backupAccountId)
	}
	if len(_backupAggregationPeriod) > 0 {
		if err := assignInputField(input, "AggregationPeriod", _backupAggregationPeriod); err != nil {
			log.Errorf("invalid --aggregation-period: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupMessageCategory) > 0 {
		input.MessageCategory = aws.String(_backupMessageCategory)
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}
	if len(_backupResourceType) > 0 {
		input.ResourceType = aws.String(_backupResourceType)
	}
	if len(_backupState) > 0 {
		if err := assignInputField(input, "State", _backupState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCopyJobSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListCopyJobSummariesOutput
	p := backup.NewListCopyJobSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns metadata about your copy jobs.
func backup_ListCopyJobs(cfg aws.Config, client *backup.Client) {
	input := &backup.ListCopyJobsInput{}

	if len(_backupByAccountId) > 0 {
		input.ByAccountId = aws.String(_backupByAccountId)
	}
	if len(_backupByCompleteAfter) > 0 {
		if err := assignInputField(input, "ByCompleteAfter", _backupByCompleteAfter); err != nil {
			log.Errorf("invalid --by-complete-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCompleteBefore) > 0 {
		if err := assignInputField(input, "ByCompleteBefore", _backupByCompleteBefore); err != nil {
			log.Errorf("invalid --by-complete-before: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedAfter) > 0 {
		if err := assignInputField(input, "ByCreatedAfter", _backupByCreatedAfter); err != nil {
			log.Errorf("invalid --by-created-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedBefore) > 0 {
		if err := assignInputField(input, "ByCreatedBefore", _backupByCreatedBefore); err != nil {
			log.Errorf("invalid --by-created-before: %s", err.Error())
			return
		}
	}
	if len(_backupByDestinationVaultArn) > 0 {
		input.ByDestinationVaultArn = aws.String(_backupByDestinationVaultArn)
	}
	if len(_backupByMessageCategory) > 0 {
		input.ByMessageCategory = aws.String(_backupByMessageCategory)
	}
	if len(_backupByParentJobId) > 0 {
		input.ByParentJobId = aws.String(_backupByParentJobId)
	}
	if len(_backupByResourceArn) > 0 {
		input.ByResourceArn = aws.String(_backupByResourceArn)
	}
	if len(_backupByResourceType) > 0 {
		input.ByResourceType = aws.String(_backupByResourceType)
	}
	if len(_backupBySourceRecoveryPointArn) > 0 {
		input.BySourceRecoveryPointArn = aws.String(_backupBySourceRecoveryPointArn)
	}
	if len(_backupByState) > 0 {
		if err := assignInputField(input, "ByState", _backupByState); err != nil {
			log.Errorf("invalid --by-state: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCopyJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListCopyJobsOutput
	p := backup.NewListCopyJobsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all frameworks for an Amazon Web Services account and Amazon
// Web Services Region.
func backup_ListFrameworks(cfg aws.Config, client *backup.Client) {
	input := &backup.ListFrameworksInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFrameworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListFrameworksOutput
	p := backup.NewListFrameworksPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This operation returns a list of recovery points that have an associated index,
// belonging to the specified account.
//
// Optional parameters you can include are: MaxResults; NextToken;
// SourceResourceArns; CreatedBefore; CreatedAfter; and ResourceType.
func backup_ListIndexedRecoveryPoints(cfg aws.Config, client *backup.Client) {
	input := &backup.ListIndexedRecoveryPointsInput{}

	if len(_backupCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _backupCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_backupCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _backupCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_backupIndexStatus) > 0 {
		if err := assignInputField(input, "IndexStatus", _backupIndexStatus); err != nil {
			log.Errorf("invalid --index-status: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}
	if len(_backupResourceType) > 0 {
		input.ResourceType = aws.String(_backupResourceType)
	}
	if len(_backupSourceResourceArn) > 0 {
		input.SourceResourceArn = aws.String(_backupSourceResourceArn)
	}

	if disablePaginator() {
		if resp, err := client.ListIndexedRecoveryPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListIndexedRecoveryPointsOutput
	p := backup.NewListIndexedRecoveryPointsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This action returns metadata about active and previous legal holds.
func backup_ListLegalHolds(cfg aws.Config, client *backup.Client) {
	input := &backup.ListLegalHoldsInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLegalHolds(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListLegalHoldsOutput
	p := backup.NewListLegalHoldsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns an array of resources successfully backed up by Backup, including the
// time the resource was saved, an Amazon Resource Name (ARN) of the resource, and
// a resource type.
func backup_ListProtectedResources(cfg aws.Config, client *backup.Client) {
	input := &backup.ListProtectedResourcesInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProtectedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListProtectedResourcesOutput
	p := backup.NewListProtectedResourcesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This request lists the protected resources corresponding to each backup vault.
func backup_ListProtectedResourcesByBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.ListProtectedResourcesByBackupVaultInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupBackupVaultAccountId) > 0 {
		input.BackupVaultAccountId = aws.String(_backupBackupVaultAccountId)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProtectedResourcesByBackupVault(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListProtectedResourcesByBackupVaultOutput
	p := backup.NewListProtectedResourcesByBackupVaultPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns detailed information about the recovery points stored in a backup vault.
func backup_ListRecoveryPointsByBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRecoveryPointsByBackupVaultInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupBackupVaultAccountId) > 0 {
		input.BackupVaultAccountId = aws.String(_backupBackupVaultAccountId)
	}
	if len(_backupByBackupPlanId) > 0 {
		input.ByBackupPlanId = aws.String(_backupByBackupPlanId)
	}
	if len(_backupByCreatedAfter) > 0 {
		if err := assignInputField(input, "ByCreatedAfter", _backupByCreatedAfter); err != nil {
			log.Errorf("invalid --by-created-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedBefore) > 0 {
		if err := assignInputField(input, "ByCreatedBefore", _backupByCreatedBefore); err != nil {
			log.Errorf("invalid --by-created-before: %s", err.Error())
			return
		}
	}
	if len(_backupByParentRecoveryPointArn) > 0 {
		input.ByParentRecoveryPointArn = aws.String(_backupByParentRecoveryPointArn)
	}
	if len(_backupByResourceArn) > 0 {
		input.ByResourceArn = aws.String(_backupByResourceArn)
	}
	if len(_backupByResourceType) > 0 {
		input.ByResourceType = aws.String(_backupByResourceType)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecoveryPointsByBackupVault(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRecoveryPointsByBackupVaultOutput
	p := backup.NewListRecoveryPointsByBackupVaultPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This action returns recovery point ARNs (Amazon Resource Names) of the
// specified legal hold.
func backup_ListRecoveryPointsByLegalHold(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRecoveryPointsByLegalHoldInput{
		// LegalHoldId: *string, // Required
	}

	if len(_backupLegalHoldId) > 0 {
		input.LegalHoldId = aws.String(_backupLegalHoldId)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecoveryPointsByLegalHold(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRecoveryPointsByLegalHoldOutput
	p := backup.NewListRecoveryPointsByLegalHoldPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// The information about the recovery points of the type specified by a resource
// Amazon Resource Name (ARN).
//
// For Amazon EFS and Amazon EC2, this action only lists recovery points created
// by Backup.
func backup_ListRecoveryPointsByResource(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRecoveryPointsByResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}
	if len(_backupManagedByAWSBackupOnly) > 0 {
		if err := assignInputField(input, "ManagedByAWSBackupOnly", _backupManagedByAWSBackupOnly); err != nil {
			log.Errorf("invalid --managed-by-aws-backup-only: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecoveryPointsByResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRecoveryPointsByResourceOutput
	p := backup.NewListRecoveryPointsByResourcePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns details about your report jobs.
func backup_ListReportJobs(cfg aws.Config, client *backup.Client) {
	input := &backup.ListReportJobsInput{}

	if len(_backupByCreationAfter) > 0 {
		if err := assignInputField(input, "ByCreationAfter", _backupByCreationAfter); err != nil {
			log.Errorf("invalid --by-creation-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCreationBefore) > 0 {
		if err := assignInputField(input, "ByCreationBefore", _backupByCreationBefore); err != nil {
			log.Errorf("invalid --by-creation-before: %s", err.Error())
			return
		}
	}
	if len(_backupByReportPlanName) > 0 {
		input.ByReportPlanName = aws.String(_backupByReportPlanName)
	}
	if len(_backupByStatus) > 0 {
		input.ByStatus = aws.String(_backupByStatus)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListReportJobsOutput
	p := backup.NewListReportJobsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of your report plans. For detailed information about a single
// report plan, use DescribeReportPlan .
func backup_ListReportPlans(cfg aws.Config, client *backup.Client) {
	input := &backup.ListReportPlansInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReportPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListReportPlansOutput
	p := backup.NewListReportPlansPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of restore access backup vaults associated with a specified
// backup vault.
func backup_ListRestoreAccessBackupVaults(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRestoreAccessBackupVaultsInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRestoreAccessBackupVaults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRestoreAccessBackupVaultsOutput
	p := backup.NewListRestoreAccessBackupVaultsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This request obtains a summary of restore jobs created or running within the
// the most recent 30 days. You can include parameters AccountID, State,
// ResourceType, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// RestourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
func backup_ListRestoreJobSummaries(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRestoreJobSummariesInput{}

	if len(_backupAccountId) > 0 {
		input.AccountId = aws.String(_backupAccountId)
	}
	if len(_backupAggregationPeriod) > 0 {
		if err := assignInputField(input, "AggregationPeriod", _backupAggregationPeriod); err != nil {
			log.Errorf("invalid --aggregation-period: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}
	if len(_backupResourceType) > 0 {
		input.ResourceType = aws.String(_backupResourceType)
	}
	if len(_backupState) > 0 {
		if err := assignInputField(input, "State", _backupState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRestoreJobSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRestoreJobSummariesOutput
	p := backup.NewListRestoreJobSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of jobs that Backup initiated to restore a saved resource,
// including details about the recovery process.
func backup_ListRestoreJobs(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRestoreJobsInput{}

	if len(_backupByAccountId) > 0 {
		input.ByAccountId = aws.String(_backupByAccountId)
	}
	if len(_backupByCompleteAfter) > 0 {
		if err := assignInputField(input, "ByCompleteAfter", _backupByCompleteAfter); err != nil {
			log.Errorf("invalid --by-complete-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCompleteBefore) > 0 {
		if err := assignInputField(input, "ByCompleteBefore", _backupByCompleteBefore); err != nil {
			log.Errorf("invalid --by-complete-before: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedAfter) > 0 {
		if err := assignInputField(input, "ByCreatedAfter", _backupByCreatedAfter); err != nil {
			log.Errorf("invalid --by-created-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCreatedBefore) > 0 {
		if err := assignInputField(input, "ByCreatedBefore", _backupByCreatedBefore); err != nil {
			log.Errorf("invalid --by-created-before: %s", err.Error())
			return
		}
	}
	if len(_backupByParentJobId) > 0 {
		input.ByParentJobId = aws.String(_backupByParentJobId)
	}
	if len(_backupByResourceType) > 0 {
		input.ByResourceType = aws.String(_backupByResourceType)
	}
	if len(_backupByRestoreTestingPlanArn) > 0 {
		input.ByRestoreTestingPlanArn = aws.String(_backupByRestoreTestingPlanArn)
	}
	if len(_backupByStatus) > 0 {
		if err := assignInputField(input, "ByStatus", _backupByStatus); err != nil {
			log.Errorf("invalid --by-status: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRestoreJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRestoreJobsOutput
	p := backup.NewListRestoreJobsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This returns restore jobs that contain the specified protected resource.
// You must include ResourceArn . You can optionally include NextToken , ByStatus ,
// MaxResults , ByRecoveryPointCreationDateAfter , and
// ByRecoveryPointCreationDateBefore .
func backup_ListRestoreJobsByProtectedResource(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRestoreJobsByProtectedResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}
	if len(_backupByRecoveryPointCreationDateAfter) > 0 {
		if err := assignInputField(input, "ByRecoveryPointCreationDateAfter", _backupByRecoveryPointCreationDateAfter); err != nil {
			log.Errorf("invalid --by-recovery-point-creation-date-after: %s", err.Error())
			return
		}
	}
	if len(_backupByRecoveryPointCreationDateBefore) > 0 {
		if err := assignInputField(input, "ByRecoveryPointCreationDateBefore", _backupByRecoveryPointCreationDateBefore); err != nil {
			log.Errorf("invalid --by-recovery-point-creation-date-before: %s", err.Error())
			return
		}
	}
	if len(_backupByStatus) > 0 {
		if err := assignInputField(input, "ByStatus", _backupByStatus); err != nil {
			log.Errorf("invalid --by-status: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRestoreJobsByProtectedResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRestoreJobsByProtectedResourceOutput
	p := backup.NewListRestoreJobsByProtectedResourcePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of restore testing plans.
func backup_ListRestoreTestingPlans(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRestoreTestingPlansInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRestoreTestingPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRestoreTestingPlansOutput
	p := backup.NewListRestoreTestingPlansPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of restore testing selections. Can be filtered by MaxResults and
// RestoreTestingPlanName .
func backup_ListRestoreTestingSelections(cfg aws.Config, client *backup.Client) {
	input := &backup.ListRestoreTestingSelectionsInput{
		// RestoreTestingPlanName: *string, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRestoreTestingSelections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListRestoreTestingSelectionsOutput
	p := backup.NewListRestoreTestingSelectionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This is a request for a summary of scan jobs created or running within the most
// recent 30 days.
func backup_ListScanJobSummaries(cfg aws.Config, client *backup.Client) {
	input := &backup.ListScanJobSummariesInput{}

	if len(_backupAccountId) > 0 {
		input.AccountId = aws.String(_backupAccountId)
	}
	if len(_backupAggregationPeriod) > 0 {
		if err := assignInputField(input, "AggregationPeriod", _backupAggregationPeriod); err != nil {
			log.Errorf("invalid --aggregation-period: %s", err.Error())
			return
		}
	}
	if len(_backupMalwareScanner) > 0 {
		if err := assignInputField(input, "MalwareScanner", _backupMalwareScanner); err != nil {
			log.Errorf("invalid --malware-scanner: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}
	if len(_backupResourceType) > 0 {
		input.ResourceType = aws.String(_backupResourceType)
	}
	if len(_backupScanResultStatus) > 0 {
		if err := assignInputField(input, "ScanResultStatus", _backupScanResultStatus); err != nil {
			log.Errorf("invalid --scan-result-status: %s", err.Error())
			return
		}
	}
	if len(_backupState) > 0 {
		if err := assignInputField(input, "State", _backupState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListScanJobSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListScanJobSummariesOutput
	p := backup.NewListScanJobSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of existing scan jobs for an authenticated account for the last
// 30 days.
func backup_ListScanJobs(cfg aws.Config, client *backup.Client) {
	input := &backup.ListScanJobsInput{}

	if len(_backupByAccountId) > 0 {
		input.ByAccountId = aws.String(_backupByAccountId)
	}
	if len(_backupByBackupVaultName) > 0 {
		input.ByBackupVaultName = aws.String(_backupByBackupVaultName)
	}
	if len(_backupByCompleteAfter) > 0 {
		if err := assignInputField(input, "ByCompleteAfter", _backupByCompleteAfter); err != nil {
			log.Errorf("invalid --by-complete-after: %s", err.Error())
			return
		}
	}
	if len(_backupByCompleteBefore) > 0 {
		if err := assignInputField(input, "ByCompleteBefore", _backupByCompleteBefore); err != nil {
			log.Errorf("invalid --by-complete-before: %s", err.Error())
			return
		}
	}
	if len(_backupByMalwareScanner) > 0 {
		if err := assignInputField(input, "ByMalwareScanner", _backupByMalwareScanner); err != nil {
			log.Errorf("invalid --by-malware-scanner: %s", err.Error())
			return
		}
	}
	if len(_backupByRecoveryPointArn) > 0 {
		input.ByRecoveryPointArn = aws.String(_backupByRecoveryPointArn)
	}
	if len(_backupByResourceArn) > 0 {
		input.ByResourceArn = aws.String(_backupByResourceArn)
	}
	if len(_backupByResourceType) > 0 {
		if err := assignInputField(input, "ByResourceType", _backupByResourceType); err != nil {
			log.Errorf("invalid --by-resource-type: %s", err.Error())
			return
		}
	}
	if len(_backupByScanResultStatus) > 0 {
		if err := assignInputField(input, "ByScanResultStatus", _backupByScanResultStatus); err != nil {
			log.Errorf("invalid --by-scan-result-status: %s", err.Error())
			return
		}
	}
	if len(_backupByState) > 0 {
		if err := assignInputField(input, "ByState", _backupByState); err != nil {
			log.Errorf("invalid --by-state: %s", err.Error())
			return
		}
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScanJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListScanJobsOutput
	p := backup.NewListScanJobsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns the tags assigned to the resource, such as a target recovery point,
// backup plan, or backup vault.
//
// This operation returns results depending on the resource type used in the value
// for resourceArn . For example, recovery points of Amazon DynamoDB with Advanced
// Settings have an ARN (Amazon Resource Name) that begins with arn:aws:backup .
// Recovery points (backups) of DynamoDB without Advanced Settings enabled have an
// ARN that begins with arn:aws:dynamodb .
//
// When this operation is called and when you include values of resourceArn that
// have an ARN other than arn:aws:backup , it may return one of the exceptions
// listed below. To prevent this exception, include only values representing
// resource types that are fully managed by Backup. These have an ARN that begins
// arn:aws:backup and they are noted in the [Feature availability by resource] table.
//
// [Feature availability by resource]: https://docs.aws.amazon.com/aws-backup/latest/devguide/backup-feature-availability.html#features-by-resource
func backup_ListTags(cfg aws.Config, client *backup.Client) {
	input := &backup.ListTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}
	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListTagsOutput
	p := backup.NewListTagsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of tiering configurations.
func backup_ListTieringConfigurations(cfg aws.Config, client *backup.Client) {
	input := &backup.ListTieringConfigurationsInput{}

	if len(_backupMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupNextToken) > 0 {
		input.NextToken = aws.String(_backupNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTieringConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backup.ListTieringConfigurationsOutput
	p := backup.NewListTieringConfigurationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Sets a resource-based policy that is used to manage access permissions on the
// target backup vault. Requires a backup vault name and an access policy document
// in JSON format.
func backup_PutBackupVaultAccessPolicy(cfg aws.Config, client *backup.Client) {
	input := &backup.PutBackupVaultAccessPolicyInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupPolicy) > 0 {
		input.Policy = aws.String(_backupPolicy)
	}

	if resp, err := client.PutBackupVaultAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies Backup Vault Lock to a backup vault, preventing attempts to delete any
// recovery point stored in or created in a backup vault. Vault Lock also prevents
// attempts to update the lifecycle policy that controls the retention period of
// any recovery point currently stored in a backup vault. If specified, Vault Lock
// enforces a minimum and maximum retention period for future backup and copy jobs
// that target a backup vault.
//
// Backup Vault Lock has been assessed by Cohasset Associates for use in
// environments that are subject to SEC 17a-4, CFTC, and FINRA regulations. For
// more information about how Backup Vault Lock relates to these regulations, see
// the [Cohasset Associates Compliance Assessment.]
//
// For more information, see [Backup Vault Lock].
//
// [Cohasset Associates Compliance Assessment.]: https://docs.aws.amazon.com/aws-backup/latest/devguide/samples/cohassetreport.zip
// [Backup Vault Lock]: https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html
func backup_PutBackupVaultLockConfiguration(cfg aws.Config, client *backup.Client) {
	input := &backup.PutBackupVaultLockConfigurationInput{
		// BackupVaultName: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupChangeableForDays) > 0 {
		if err := assignInputField(input, "ChangeableForDays", _backupChangeableForDays); err != nil {
			log.Errorf("invalid --changeable-for-days: %s", err.Error())
			return
		}
	}
	if len(_backupMaxRetentionDays) > 0 {
		if err := assignInputField(input, "MaxRetentionDays", _backupMaxRetentionDays); err != nil {
			log.Errorf("invalid --max-retention-days: %s", err.Error())
			return
		}
	}
	if len(_backupMinRetentionDays) > 0 {
		if err := assignInputField(input, "MinRetentionDays", _backupMinRetentionDays); err != nil {
			log.Errorf("invalid --min-retention-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBackupVaultLockConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Turns on notifications on a backup vault for the specified topic and events.
func backup_PutBackupVaultNotifications(cfg aws.Config, client *backup.Client) {
	input := &backup.PutBackupVaultNotificationsInput{
		// BackupVaultEvents: []types.BackupVaultEvent, // Required
		// BackupVaultName: *string, // Required
		// SNSTopicArn: *string, // Required
	}

	if len(_backupBackupVaultEvents) > 0 {
		if err := assignInputField(input, "BackupVaultEvents", _backupBackupVaultEvents); err != nil {
			log.Errorf("invalid --backup-vault-events: %s", err.Error())
			return
		}
	}
	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupSNSTopicArn) > 0 {
		input.SNSTopicArn = aws.String(_backupSNSTopicArn)
	}

	if resp, err := client.PutBackupVaultNotifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request allows you to send your independent self-run restore test
// validation results. RestoreJobId and ValidationStatus are required. Optionally,
// you can input a ValidationStatusMessage .
func backup_PutRestoreValidationResult(cfg aws.Config, client *backup.Client) {
	input := &backup.PutRestoreValidationResultInput{
		// RestoreJobId: *string, // Required
		// ValidationStatus: types.RestoreValidationStatus, // Required
	}

	if len(_backupRestoreJobId) > 0 {
		input.RestoreJobId = aws.String(_backupRestoreJobId)
	}
	if len(_backupValidationStatus) > 0 {
		if err := assignInputField(input, "ValidationStatus", _backupValidationStatus); err != nil {
			log.Errorf("invalid --validation-status: %s", err.Error())
			return
		}
	}
	if len(_backupValidationStatusMessage) > 0 {
		input.ValidationStatusMessage = aws.String(_backupValidationStatusMessage)
	}

	if resp, err := client.PutRestoreValidationResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes access to a restore access backup vault, removing the ability to
// restore from its recovery points and permanently deleting the vault.
func backup_RevokeRestoreAccessBackupVault(cfg aws.Config, client *backup.Client) {
	input := &backup.RevokeRestoreAccessBackupVaultInput{
		// BackupVaultName: *string, // Required
		// RestoreAccessBackupVaultArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRestoreAccessBackupVaultArn) > 0 {
		input.RestoreAccessBackupVaultArn = aws.String(_backupRestoreAccessBackupVaultArn)
	}
	if len(_backupRequesterComment) > 0 {
		input.RequesterComment = aws.String(_backupRequesterComment)
	}

	if resp, err := client.RevokeRestoreAccessBackupVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an on-demand backup job for the specified resource.
func backup_StartBackupJob(cfg aws.Config, client *backup.Client) {
	input := &backup.StartBackupJobInput{
		// BackupVaultName: *string, // Required
		// IamRoleArn: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_backupIamRoleArn)
	}
	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}
	if len(_backupBackupOptions) > 0 {
		if err := assignInputField(input, "BackupOptions", _backupBackupOptions); err != nil {
			log.Errorf("invalid --backup-options: %s", err.Error())
			return
		}
	}
	if len(_backupCompleteWindowMinutes) > 0 {
		if err := assignInputField(input, "CompleteWindowMinutes", _backupCompleteWindowMinutes); err != nil {
			log.Errorf("invalid --complete-window-minutes: %s", err.Error())
			return
		}
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupIndex) > 0 {
		if err := assignInputField(input, "Index", _backupIndex); err != nil {
			log.Errorf("invalid --index: %s", err.Error())
			return
		}
	}
	if len(_backupLifecycle) > 0 {
		if err := assignInputField(input, "Lifecycle", _backupLifecycle); err != nil {
			log.Errorf("invalid --lifecycle: %s", err.Error())
			return
		}
	}
	if len(_backupLogicallyAirGappedBackupVaultArn) > 0 {
		input.LogicallyAirGappedBackupVaultArn = aws.String(_backupLogicallyAirGappedBackupVaultArn)
	}
	if len(_backupRecoveryPointTags) > 0 {
		if err := assignInputField(input, "RecoveryPointTags", _backupRecoveryPointTags); err != nil {
			log.Errorf("invalid --recovery-point-tags: %s", err.Error())
			return
		}
	}
	if len(_backupStartWindowMinutes) > 0 {
		if err := assignInputField(input, "StartWindowMinutes", _backupStartWindowMinutes); err != nil {
			log.Errorf("invalid --start-window-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBackupJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a job to create a one-time copy of the specified resource.
// Does not support continuous backups.
//
// See [Copy job retry] for information on how Backup retries copy job operations.
//
// [Copy job retry]: https://docs.aws.amazon.com/aws-backup/latest/devguide/recov-point-create-a-copy.html#backup-copy-retry
func backup_StartCopyJob(cfg aws.Config, client *backup.Client) {
	input := &backup.StartCopyJobInput{
		// DestinationBackupVaultArn: *string, // Required
		// IamRoleArn: *string, // Required
		// RecoveryPointArn: *string, // Required
		// SourceBackupVaultName: *string, // Required
	}

	if len(_backupDestinationBackupVaultArn) > 0 {
		input.DestinationBackupVaultArn = aws.String(_backupDestinationBackupVaultArn)
	}
	if len(_backupIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_backupIamRoleArn)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupSourceBackupVaultName) > 0 {
		input.SourceBackupVaultName = aws.String(_backupSourceBackupVaultName)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupLifecycle) > 0 {
		if err := assignInputField(input, "Lifecycle", _backupLifecycle); err != nil {
			log.Errorf("invalid --lifecycle: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCopyJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an on-demand report job for the specified report plan.
func backup_StartReportJob(cfg aws.Config, client *backup.Client) {
	input := &backup.StartReportJobInput{
		// ReportPlanName: *string, // Required
	}

	if len(_backupReportPlanName) > 0 {
		input.ReportPlanName = aws.String(_backupReportPlanName)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}

	if resp, err := client.StartReportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Recovers the saved resource identified by an Amazon Resource Name (ARN).
func backup_StartRestoreJob(cfg aws.Config, client *backup.Client) {
	input := &backup.StartRestoreJobInput{
		// Metadata: map[string]string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _backupMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupCopySourceTagsToRestoredResource) > 0 {
		if err := assignInputField(input, "CopySourceTagsToRestoredResource", _backupCopySourceTagsToRestoredResource); err != nil {
			log.Errorf("invalid --copy-source-tags-to-restored-resource: %s", err.Error())
			return
		}
	}
	if len(_backupIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_backupIamRoleArn)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupResourceType) > 0 {
		input.ResourceType = aws.String(_backupResourceType)
	}

	if resp, err := client.StartRestoreJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts scanning jobs for specific resources.
func backup_StartScanJob(cfg aws.Config, client *backup.Client) {
	input := &backup.StartScanJobInput{
		// BackupVaultName: *string, // Required
		// IamRoleArn: *string, // Required
		// MalwareScanner: types.MalwareScanner, // Required
		// RecoveryPointArn: *string, // Required
		// ScanMode: types.ScanMode, // Required
		// ScannerRoleArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_backupIamRoleArn)
	}
	if len(_backupMalwareScanner) > 0 {
		if err := assignInputField(input, "MalwareScanner", _backupMalwareScanner); err != nil {
			log.Errorf("invalid --malware-scanner: %s", err.Error())
			return
		}
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupScanMode) > 0 {
		if err := assignInputField(input, "ScanMode", _backupScanMode); err != nil {
			log.Errorf("invalid --scan-mode: %s", err.Error())
			return
		}
	}
	if len(_backupScannerRoleArn) > 0 {
		input.ScannerRoleArn = aws.String(_backupScannerRoleArn)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupScanBaseRecoveryPointArn) > 0 {
		input.ScanBaseRecoveryPointArn = aws.String(_backupScanBaseRecoveryPointArn)
	}

	if resp, err := client.StartScanJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel a job to create a one-time backup of a resource.
// This action is not supported for the following services:
//
// - Amazon Aurora
//
// - Amazon DocumentDB (with MongoDB compatibility)
//
// - Amazon FSx for Lustre
//
// - Amazon FSx for NetApp ONTAP
//
// - Amazon FSx for OpenZFS
//
// - Amazon FSx for Windows File Server
//
// - Amazon Neptune
//
// - SAP HANA databases on Amazon EC2 instances
//
// - Amazon RDS
func backup_StopBackupJob(cfg aws.Config, client *backup.Client) {
	input := &backup.StopBackupJobInput{
		// BackupJobId: *string, // Required
	}

	if len(_backupBackupJobId) > 0 {
		input.BackupJobId = aws.String(_backupBackupJobId)
	}

	if resp, err := client.StopBackupJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a set of key-value pairs to a resource.
func backup_TagResource(cfg aws.Config, client *backup.Client) {
	input := &backup.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}
	if len(_backupTags) > 0 {
		if err := assignInputField(input, "Tags", _backupTags); err != nil {
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

// Removes a set of key-value pairs from a recovery point, backup plan, or backup
// vault identified by an Amazon Resource Name (ARN)
//
// This API is not supported for recovery points for resource types including
// Aurora, Amazon DocumentDB. Amazon EBS, Amazon FSx, Neptune, and Amazon RDS.
func backup_UntagResource(cfg aws.Config, client *backup.Client) {
	input := &backup.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeyList: []string, // Required
	}

	if len(_backupResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupResourceArn)
	}
	if len(_backupTagKeyList) > 0 {
		input.TagKeyList = append([]string(nil), _backupTagKeyList...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified backup plan. The new version is uniquely identified by
// its ID.
func backup_UpdateBackupPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateBackupPlanInput{
		// BackupPlan: *types.BackupPlanInput, // Required
		// BackupPlanId: *string, // Required
	}

	if len(_backupBackupPlan) > 0 {
		if err := assignInputField(input, "BackupPlan", _backupBackupPlan); err != nil {
			log.Errorf("invalid --backup-plan: %s", err.Error())
			return
		}
	}
	if len(_backupBackupPlanId) > 0 {
		input.BackupPlanId = aws.String(_backupBackupPlanId)
	}

	if resp, err := client.UpdateBackupPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified framework.
func backup_UpdateFramework(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateFrameworkInput{
		// FrameworkName: *string, // Required
	}

	if len(_backupFrameworkName) > 0 {
		input.FrameworkName = aws.String(_backupFrameworkName)
	}
	if len(_backupFrameworkControls) > 0 {
		if err := assignInputField(input, "FrameworkControls", _backupFrameworkControls); err != nil {
			log.Errorf("invalid --framework-controls: %s", err.Error())
			return
		}
	}
	if len(_backupFrameworkDescription) > 0 {
		input.FrameworkDescription = aws.String(_backupFrameworkDescription)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}

	if resp, err := client.UpdateFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates whether the Amazon Web Services account is opted in to cross-account
// backup. Returns an error if the account is not an Organizations management
// account. Use the DescribeGlobalSettings API to determine the current settings.
func backup_UpdateGlobalSettings(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateGlobalSettingsInput{}

	if len(_backupGlobalSettings) > 0 {
		if err := assignInputField(input, "GlobalSettings", _backupGlobalSettings); err != nil {
			log.Errorf("invalid --global-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates the settings of a recovery point index.
// Required: BackupVaultName, RecoveryPointArn, and IAMRoleArn
func backup_UpdateRecoveryPointIndexSettings(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateRecoveryPointIndexSettingsInput{
		// BackupVaultName: *string, // Required
		// Index: types.Index, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupIndex) > 0 {
		if err := assignInputField(input, "Index", _backupIndex); err != nil {
			log.Errorf("invalid --index: %s", err.Error())
			return
		}
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_backupIamRoleArn)
	}

	if resp, err := client.UpdateRecoveryPointIndexSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the transition lifecycle of a recovery point.
// The lifecycle defines when a protected resource is transitioned to cold storage
// and when it expires. Backup transitions and expires backups automatically
// according to the lifecycle that you define.
//
// Resource types that can transition to cold storage are listed in the [Feature availability by resource] table.
// Backup ignores this expression for other resource types.
//
// Backups transitioned to cold storage must be stored in cold storage for a
// minimum of 90 days. Therefore, the “retention” setting must be 90 days greater
// than the “transition to cold after days” setting. The “transition to cold after
// days” setting cannot be changed after a backup has been transitioned to cold.
//
// If your lifecycle currently uses the parameters DeleteAfterDays and
// MoveToColdStorageAfterDays , include these parameters and their values when you
// call this operation. Not including them may result in your plan updating with
// null values.
//
// This operation does not support continuous backups.
//
// [Feature availability by resource]: https://docs.aws.amazon.com/aws-backup/latest/devguide/backup-feature-availability.html#features-by-resource
func backup_UpdateRecoveryPointLifecycle(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateRecoveryPointLifecycleInput{
		// BackupVaultName: *string, // Required
		// RecoveryPointArn: *string, // Required
	}

	if len(_backupBackupVaultName) > 0 {
		input.BackupVaultName = aws.String(_backupBackupVaultName)
	}
	if len(_backupRecoveryPointArn) > 0 {
		input.RecoveryPointArn = aws.String(_backupRecoveryPointArn)
	}
	if len(_backupLifecycle) > 0 {
		if err := assignInputField(input, "Lifecycle", _backupLifecycle); err != nil {
			log.Errorf("invalid --lifecycle: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecoveryPointLifecycle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the current service opt-in settings for the Region.
// Use the DescribeRegionSettings API to determine the resource types that are
// supported.
func backup_UpdateRegionSettings(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateRegionSettingsInput{}

	if len(_backupResourceTypeManagementPreference) > 0 {
		if err := assignInputField(input, "ResourceTypeManagementPreference", _backupResourceTypeManagementPreference); err != nil {
			log.Errorf("invalid --resource-type-management-preference: %s", err.Error())
			return
		}
	}
	if len(_backupResourceTypeOptInPreference) > 0 {
		if err := assignInputField(input, "ResourceTypeOptInPreference", _backupResourceTypeOptInPreference); err != nil {
			log.Errorf("invalid --resource-type-opt-in-preference: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRegionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified report plan.
func backup_UpdateReportPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateReportPlanInput{
		// ReportPlanName: *string, // Required
	}

	if len(_backupReportPlanName) > 0 {
		input.ReportPlanName = aws.String(_backupReportPlanName)
	}
	if len(_backupIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_backupIdempotencyToken)
	}
	if len(_backupReportDeliveryChannel) > 0 {
		if err := assignInputField(input, "ReportDeliveryChannel", _backupReportDeliveryChannel); err != nil {
			log.Errorf("invalid --report-delivery-channel: %s", err.Error())
			return
		}
	}
	if len(_backupReportPlanDescription) > 0 {
		input.ReportPlanDescription = aws.String(_backupReportPlanDescription)
	}
	if len(_backupReportSetting) > 0 {
		if err := assignInputField(input, "ReportSetting", _backupReportSetting); err != nil {
			log.Errorf("invalid --report-setting: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReportPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request will send changes to your specified restore testing plan.
// RestoreTestingPlanName cannot be updated after it is created.
//
// RecoveryPointSelection can contain:
//
// - Algorithm
//
// - ExcludeVaults
//
// - IncludeVaults
//
// - RecoveryPointTypes
//
// - SelectionWindowDays
func backup_UpdateRestoreTestingPlan(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateRestoreTestingPlanInput{
		// RestoreTestingPlan: *types.RestoreTestingPlanForUpdate, // Required
		// RestoreTestingPlanName: *string, // Required
	}

	if len(_backupRestoreTestingPlan) > 0 {
		if err := assignInputField(input, "RestoreTestingPlan", _backupRestoreTestingPlan); err != nil {
			log.Errorf("invalid --restore-testing-plan: %s", err.Error())
			return
		}
	}
	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}

	if resp, err := client.UpdateRestoreTestingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified restore testing selection.
// Most elements except the RestoreTestingSelectionName can be updated with this
// request.
//
// You can use either protected resource ARNs or conditions, but not both.
func backup_UpdateRestoreTestingSelection(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateRestoreTestingSelectionInput{
		// RestoreTestingPlanName: *string, // Required
		// RestoreTestingSelection: *types.RestoreTestingSelectionForUpdate, // Required
		// RestoreTestingSelectionName: *string, // Required
	}

	if len(_backupRestoreTestingPlanName) > 0 {
		input.RestoreTestingPlanName = aws.String(_backupRestoreTestingPlanName)
	}
	if len(_backupRestoreTestingSelection) > 0 {
		if err := assignInputField(input, "RestoreTestingSelection", _backupRestoreTestingSelection); err != nil {
			log.Errorf("invalid --restore-testing-selection: %s", err.Error())
			return
		}
	}
	if len(_backupRestoreTestingSelectionName) > 0 {
		input.RestoreTestingSelectionName = aws.String(_backupRestoreTestingSelectionName)
	}

	if resp, err := client.UpdateRestoreTestingSelection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This request will send changes to your specified tiering configuration.
// TieringConfigurationName cannot be updated after it is created.
//
// ResourceSelection can contain:
//
// - Resources
//
// - TieringDownSettingsInDays
//
// - ResourceType
func backup_UpdateTieringConfiguration(cfg aws.Config, client *backup.Client) {
	input := &backup.UpdateTieringConfigurationInput{
		// TieringConfiguration: *types.TieringConfigurationInputForUpdate, // Required
		// TieringConfigurationName: *string, // Required
	}

	if len(_backupTieringConfiguration) > 0 {
		if err := assignInputField(input, "TieringConfiguration", _backupTieringConfiguration); err != nil {
			log.Errorf("invalid --tiering-configuration: %s", err.Error())
			return
		}
	}
	if len(_backupTieringConfigurationName) > 0 {
		input.TieringConfigurationName = aws.String(_backupTieringConfigurationName)
	}

	if resp, err := client.UpdateTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_backupCmd)
	_backupCmd.Flags().SortFlags = false

	_backupCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_backupCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_backupCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_backupCmd.Flags().StringVarP(&_backupAccountId, "account-id", "", "", "Account ID")
	_backupCmd.Flags().StringVarP(&_backupAggregationPeriod, "aggregation-period", "", "", "Aggregation Period")
	_backupCmd.Flags().StringVarP(&_backupBackupJobId, "backup-job-id", "", "", "Backup Job ID")
	_backupCmd.Flags().StringVarP(&_backupBackupOptions, "backup-options", "", "", "Backup Options")
	_backupCmd.Flags().StringVarP(&_backupBackupPlan, "backup-plan", "", "", "Backup Plan")
	_backupCmd.Flags().StringVarP(&_backupBackupPlanId, "backup-plan-id", "", "", "Backup Plan ID")
	_backupCmd.Flags().StringVarP(&_backupBackupPlanTags, "backup-plan-tags", "", "", "Backup Plan Tags")
	_backupCmd.Flags().StringVarP(&_backupBackupPlanTemplateId, "backup-plan-template-id", "", "", "Backup Plan Template ID")
	_backupCmd.Flags().StringVarP(&_backupBackupPlanTemplateJson, "backup-plan-template-json", "", "", "Backup Plan Template JSON")
	_backupCmd.Flags().StringVarP(&_backupBackupSelection, "backup-selection", "", "", "Backup Selection")
	_backupCmd.Flags().StringVarP(&_backupBackupVaultAccountId, "backup-vault-account-id", "", "", "Backup Vault Account ID")
	_backupCmd.Flags().StringVarP(&_backupBackupVaultEvents, "backup-vault-events", "", "", "Backup Vault Events")
	_backupCmd.Flags().StringVarP(&_backupBackupVaultName, "backup-vault-name", "", "", "Backup Vault Name")
	_backupCmd.Flags().StringVarP(&_backupBackupVaultTags, "backup-vault-tags", "", "", "Backup Vault Tags")
	_backupCmd.Flags().StringVarP(&_backupByAccountId, "by-account-id", "", "", "By Account ID")
	_backupCmd.Flags().StringVarP(&_backupByBackupPlanId, "by-backup-plan-id", "", "", "By Backup Plan ID")
	_backupCmd.Flags().StringVarP(&_backupByBackupVaultName, "by-backup-vault-name", "", "", "By Backup Vault Name")
	_backupCmd.Flags().StringVarP(&_backupByCompleteAfter, "by-complete-after", "", "", "By Complete After")
	_backupCmd.Flags().StringVarP(&_backupByCompleteBefore, "by-complete-before", "", "", "By Complete Before")
	_backupCmd.Flags().StringVarP(&_backupByCreatedAfter, "by-created-after", "", "", "By Created After")
	_backupCmd.Flags().StringVarP(&_backupByCreatedBefore, "by-created-before", "", "", "By Created Before")
	_backupCmd.Flags().StringVarP(&_backupByCreationAfter, "by-creation-after", "", "", "By Creation After")
	_backupCmd.Flags().StringVarP(&_backupByCreationBefore, "by-creation-before", "", "", "By Creation Before")
	_backupCmd.Flags().StringVarP(&_backupByDestinationVaultArn, "by-destination-vault-arn", "", "", "By Destination Vault ARN")
	_backupCmd.Flags().StringVarP(&_backupByMalwareScanner, "by-malware-scanner", "", "", "By Malware Scanner")
	_backupCmd.Flags().StringVarP(&_backupByMessageCategory, "by-message-category", "", "", "By Message Category")
	_backupCmd.Flags().StringVarP(&_backupByParentJobId, "by-parent-job-id", "", "", "By Parent Job ID")
	_backupCmd.Flags().StringVarP(&_backupByParentRecoveryPointArn, "by-parent-recovery-point-arn", "", "", "By Parent Recovery Point ARN")
	_backupCmd.Flags().StringVarP(&_backupByRecoveryPointArn, "by-recovery-point-arn", "", "", "By Recovery Point ARN")
	_backupCmd.Flags().StringVarP(&_backupByRecoveryPointCreationDateAfter, "by-recovery-point-creation-date-after", "", "", "By Recovery Point Creation Date After")
	_backupCmd.Flags().StringVarP(&_backupByRecoveryPointCreationDateBefore, "by-recovery-point-creation-date-before", "", "", "By Recovery Point Creation Date Before")
	_backupCmd.Flags().StringVarP(&_backupByReportPlanName, "by-report-plan-name", "", "", "By Report Plan Name")
	_backupCmd.Flags().StringVarP(&_backupByResourceArn, "by-resource-arn", "", "", "By Resource ARN")
	_backupCmd.Flags().StringVarP(&_backupByResourceType, "by-resource-type", "", "", "By Resource Type")
	_backupCmd.Flags().StringVarP(&_backupByRestoreTestingPlanArn, "by-restore-testing-plan-arn", "", "", "By Restore Testing Plan ARN")
	_backupCmd.Flags().StringVarP(&_backupByScanResultStatus, "by-scan-result-status", "", "", "By Scan Result Status")
	_backupCmd.Flags().StringVarP(&_backupByShared, "by-shared", "", "", "By Shared")
	_backupCmd.Flags().StringVarP(&_backupBySourceRecoveryPointArn, "by-source-recovery-point-arn", "", "", "By Source Recovery Point ARN")
	_backupCmd.Flags().StringVarP(&_backupByState, "by-state", "", "", "By State")
	_backupCmd.Flags().StringVarP(&_backupByStatus, "by-status", "", "", "By Status")
	_backupCmd.Flags().StringVarP(&_backupByVaultType, "by-vault-type", "", "", "By Vault Type")
	_backupCmd.Flags().StringVarP(&_backupCancelDescription, "cancel-description", "", "", "Cancel Description")
	_backupCmd.Flags().StringVarP(&_backupChangeableForDays, "changeable-for-days", "", "", "Changeable For Days")
	_backupCmd.Flags().StringVarP(&_backupCompleteWindowMinutes, "complete-window-minutes", "", "", "Complete Window Minutes")
	_backupCmd.Flags().StringVarP(&_backupCopyJobId, "copy-job-id", "", "", "Copy Job ID")
	_backupCmd.Flags().StringVarP(&_backupCopySourceTagsToRestoredResource, "copy-source-tags-to-restored-resource", "", "", "Copy Source Tags To Restored Resource")
	_backupCmd.Flags().StringVarP(&_backupCreatedAfter, "created-after", "", "", "Created After")
	_backupCmd.Flags().StringVarP(&_backupCreatedBefore, "created-before", "", "", "Created Before")
	_backupCmd.Flags().StringVarP(&_backupCreatorRequestId, "creator-request-id", "", "", "Creator Request ID")
	_backupCmd.Flags().StringVarP(&_backupDescription, "description", "", "", "Description")
	_backupCmd.Flags().StringVarP(&_backupDestinationBackupVaultArn, "destination-backup-vault-arn", "", "", "Destination Backup Vault ARN")
	_backupCmd.Flags().StringVarP(&_backupEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_backupCmd.Flags().StringVarP(&_backupFrameworkControls, "framework-controls", "", "", "Framework Controls")
	_backupCmd.Flags().StringVarP(&_backupFrameworkDescription, "framework-description", "", "", "Framework Description")
	_backupCmd.Flags().StringVarP(&_backupFrameworkName, "framework-name", "", "", "Framework Name")
	_backupCmd.Flags().StringVarP(&_backupFrameworkTags, "framework-tags", "", "", "Framework Tags")
	_backupCmd.Flags().StringVarP(&_backupGlobalSettings, "global-settings", "", "", "Global Settings")
	_backupCmd.Flags().StringVarP(&_backupIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_backupCmd.Flags().StringVarP(&_backupIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_backupCmd.Flags().StringVarP(&_backupIncludeDeleted, "include-deleted", "", "", "Include Deleted")
	_backupCmd.Flags().StringVarP(&_backupIndex, "index", "", "", "Index")
	_backupCmd.Flags().StringVarP(&_backupIndexStatus, "index-status", "", "", "Index Status")
	_backupCmd.Flags().StringVarP(&_backupLegalHoldId, "legal-hold-id", "", "", "Legal Hold ID")
	_backupCmd.Flags().StringVarP(&_backupLifecycle, "lifecycle", "", "", "Lifecycle")
	_backupCmd.Flags().StringVarP(&_backupLogicallyAirGappedBackupVaultArn, "logically-air-gapped-backup-vault-arn", "", "", "Logically Air Gapped Backup Vault ARN")
	_backupCmd.Flags().StringVarP(&_backupMalwareScanner, "malware-scanner", "", "", "Malware Scanner")
	_backupCmd.Flags().StringVarP(&_backupManagedByAWSBackupOnly, "managed-by-aws-backup-only", "", "", "Managed By AWS Backup Only")
	_backupCmd.Flags().StringVarP(&_backupMaxResults, "max-results", "", "", "Max Results")
	_backupCmd.Flags().StringVarP(&_backupMaxRetentionDays, "max-retention-days", "", "", "Max Retention Days")
	_backupCmd.Flags().StringVarP(&_backupMaxScheduledRunsPreview, "max-scheduled-runs-preview", "", "", "Max Scheduled Runs Preview")
	_backupCmd.Flags().StringVarP(&_backupMessageCategory, "message-category", "", "", "Message Category")
	_backupCmd.Flags().StringVarP(&_backupMetadata, "metadata", "", "", "Metadata")
	_backupCmd.Flags().StringVarP(&_backupMinRetentionDays, "min-retention-days", "", "", "Min Retention Days")
	_backupCmd.Flags().StringVarP(&_backupMpaApprovalTeamArn, "mpa-approval-team-arn", "", "", "Mpa Approval Team ARN")
	_backupCmd.Flags().StringVarP(&_backupNextToken, "next-token", "", "", "Next Token")
	_backupCmd.Flags().StringVarP(&_backupPolicy, "policy", "", "", "Policy")
	_backupCmd.Flags().StringVarP(&_backupRecoveryPointArn, "recovery-point-arn", "", "", "Recovery Point ARN")
	_backupCmd.Flags().StringVarP(&_backupRecoveryPointSelection, "recovery-point-selection", "", "", "Recovery Point Selection")
	_backupCmd.Flags().StringVarP(&_backupRecoveryPointTags, "recovery-point-tags", "", "", "Recovery Point Tags")
	_backupCmd.Flags().StringVarP(&_backupReportDeliveryChannel, "report-delivery-channel", "", "", "Report Delivery Channel")
	_backupCmd.Flags().StringVarP(&_backupReportJobId, "report-job-id", "", "", "Report Job ID")
	_backupCmd.Flags().StringVarP(&_backupReportPlanDescription, "report-plan-description", "", "", "Report Plan Description")
	_backupCmd.Flags().StringVarP(&_backupReportPlanName, "report-plan-name", "", "", "Report Plan Name")
	_backupCmd.Flags().StringVarP(&_backupReportPlanTags, "report-plan-tags", "", "", "Report Plan Tags")
	_backupCmd.Flags().StringVarP(&_backupReportSetting, "report-setting", "", "", "Report Setting")
	_backupCmd.Flags().StringVarP(&_backupRequesterComment, "requester-comment", "", "", "Requester Comment")
	_backupCmd.Flags().StringVarP(&_backupResourceArn, "resource-arn", "", "", "Resource ARN")
	_backupCmd.Flags().StringVarP(&_backupResourceType, "resource-type", "", "", "Resource Type")
	_backupCmd.Flags().StringVarP(&_backupResourceTypeManagementPreference, "resource-type-management-preference", "", "", "Resource Type Management Preference")
	_backupCmd.Flags().StringVarP(&_backupResourceTypeOptInPreference, "resource-type-opt-in-preference", "", "", "Resource Type Opt In Preference")
	_backupCmd.Flags().StringVarP(&_backupRestoreAccessBackupVaultArn, "restore-access-backup-vault-arn", "", "", "Restore Access Backup Vault ARN")
	_backupCmd.Flags().StringVarP(&_backupRestoreJobId, "restore-job-id", "", "", "Restore Job ID")
	_backupCmd.Flags().StringVarP(&_backupRestoreTestingPlan, "restore-testing-plan", "", "", "Restore Testing Plan")
	_backupCmd.Flags().StringVarP(&_backupRestoreTestingPlanName, "restore-testing-plan-name", "", "", "Restore Testing Plan Name")
	_backupCmd.Flags().StringVarP(&_backupRestoreTestingSelection, "restore-testing-selection", "", "", "Restore Testing Selection")
	_backupCmd.Flags().StringVarP(&_backupRestoreTestingSelectionName, "restore-testing-selection-name", "", "", "Restore Testing Selection Name")
	_backupCmd.Flags().StringVarP(&_backupRetainRecordInDays, "retain-record-in-days", "", "", "Retain Record In Days")
	_backupCmd.Flags().StringVarP(&_backupScanBaseRecoveryPointArn, "scan-base-recovery-point-arn", "", "", "Scan Base Recovery Point ARN")
	_backupCmd.Flags().StringVarP(&_backupScanJobId, "scan-job-id", "", "", "Scan Job ID")
	_backupCmd.Flags().StringVarP(&_backupScanMode, "scan-mode", "", "", "Scan Mode")
	_backupCmd.Flags().StringVarP(&_backupScanResultStatus, "scan-result-status", "", "", "Scan Result Status")
	_backupCmd.Flags().StringVarP(&_backupScannerRoleArn, "scanner-role-arn", "", "", "Scanner Role ARN")
	_backupCmd.Flags().StringVarP(&_backupSelectionId, "selection-id", "", "", "Selection ID")
	_backupCmd.Flags().StringVarP(&_backupSNSTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_backupCmd.Flags().StringVarP(&_backupSourceBackupVaultArn, "source-backup-vault-arn", "", "", "Source Backup Vault ARN")
	_backupCmd.Flags().StringVarP(&_backupSourceBackupVaultName, "source-backup-vault-name", "", "", "Source Backup Vault Name")
	_backupCmd.Flags().StringVarP(&_backupSourceResourceArn, "source-resource-arn", "", "", "Source Resource ARN")
	_backupCmd.Flags().StringVarP(&_backupStartWindowMinutes, "start-window-minutes", "", "", "Start Window Minutes")
	_backupCmd.Flags().StringVarP(&_backupState, "state", "", "", "State")
	_backupCmd.Flags().StringSliceVarP(&_backupTagKeyList, "tag-key-list", "", nil, "Tag Key List")
	_backupCmd.Flags().StringVarP(&_backupTags, "tags", "", "", "Tags")
	_backupCmd.Flags().StringVarP(&_backupTieringConfiguration, "tiering-configuration", "", "", "Tiering Configuration")
	_backupCmd.Flags().StringVarP(&_backupTieringConfigurationName, "tiering-configuration-name", "", "", "Tiering Configuration Name")
	_backupCmd.Flags().StringVarP(&_backupTieringConfigurationTags, "tiering-configuration-tags", "", "", "Tiering Configuration Tags")
	_backupCmd.Flags().StringVarP(&_backupTitle, "title", "", "", "Title")
	_backupCmd.Flags().StringVarP(&_backupValidationStatus, "validation-status", "", "", "Validation Status")
	_backupCmd.Flags().StringVarP(&_backupValidationStatusMessage, "validation-status-message", "", "", "Validation Status Message")
	_backupCmd.Flags().StringVarP(&_backupVersionId, "version-id", "", "", "Version ID")

	_backupCmd.Flags().BoolVarP(&_backupAssociateBackupVaultMpaApprovalTeam, "associate-backup-vault-mpa-approval-team", "", false, "Associate Backup Vault Mpa Approval Team")
	_backupCmd.Flags().BoolVarP(&_backupCancelLegalHold, "cancel-legal-hold", "", false, "Cancel Legal Hold")
	_backupCmd.Flags().BoolVarP(&_backupCreateBackupPlan, "create-backup-plan", "", false, "Create Backup Plan")
	_backupCmd.Flags().BoolVarP(&_backupCreateBackupSelection, "create-backup-selection", "", false, "Create Backup Selection")
	_backupCmd.Flags().BoolVarP(&_backupCreateBackupVault, "create-backup-vault", "", false, "Create Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupCreateFramework, "create-framework", "", false, "Create Framework")
	_backupCmd.Flags().BoolVarP(&_backupCreateLegalHold, "create-legal-hold", "", false, "Create Legal Hold")
	_backupCmd.Flags().BoolVarP(&_backupCreateLogicallyAirGappedBackupVault, "create-logically-air-gapped-backup-vault", "", false, "Create Logically Air Gapped Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupCreateReportPlan, "create-report-plan", "", false, "Create Report Plan")
	_backupCmd.Flags().BoolVarP(&_backupCreateRestoreAccessBackupVault, "create-restore-access-backup-vault", "", false, "Create Restore Access Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupCreateRestoreTestingPlan, "create-restore-testing-plan", "", false, "Create Restore Testing Plan")
	_backupCmd.Flags().BoolVarP(&_backupCreateRestoreTestingSelection, "create-restore-testing-selection", "", false, "Create Restore Testing Selection")
	_backupCmd.Flags().BoolVarP(&_backupCreateTieringConfiguration, "create-tiering-configuration", "", false, "Create Tiering Configuration")
	_backupCmd.Flags().BoolVarP(&_backupDeleteBackupPlan, "delete-backup-plan", "", false, "Delete Backup Plan")
	_backupCmd.Flags().BoolVarP(&_backupDeleteBackupSelection, "delete-backup-selection", "", false, "Delete Backup Selection")
	_backupCmd.Flags().BoolVarP(&_backupDeleteBackupVault, "delete-backup-vault", "", false, "Delete Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupDeleteBackupVaultAccessPolicy, "delete-backup-vault-access-policy", "", false, "Delete Backup Vault Access Policy")
	_backupCmd.Flags().BoolVarP(&_backupDeleteBackupVaultLockConfiguration, "delete-backup-vault-lock-configuration", "", false, "Delete Backup Vault Lock Configuration")
	_backupCmd.Flags().BoolVarP(&_backupDeleteBackupVaultNotifications, "delete-backup-vault-notifications", "", false, "Delete Backup Vault Notifications")
	_backupCmd.Flags().BoolVarP(&_backupDeleteFramework, "delete-framework", "", false, "Delete Framework")
	_backupCmd.Flags().BoolVarP(&_backupDeleteRecoveryPoint, "delete-recovery-point", "", false, "Delete Recovery Point")
	_backupCmd.Flags().BoolVarP(&_backupDeleteReportPlan, "delete-report-plan", "", false, "Delete Report Plan")
	_backupCmd.Flags().BoolVarP(&_backupDeleteRestoreTestingPlan, "delete-restore-testing-plan", "", false, "Delete Restore Testing Plan")
	_backupCmd.Flags().BoolVarP(&_backupDeleteRestoreTestingSelection, "delete-restore-testing-selection", "", false, "Delete Restore Testing Selection")
	_backupCmd.Flags().BoolVarP(&_backupDeleteTieringConfiguration, "delete-tiering-configuration", "", false, "Delete Tiering Configuration")
	_backupCmd.Flags().BoolVarP(&_backupDescribeBackupJob, "describe-backup-job", "", false, "Describe Backup Job")
	_backupCmd.Flags().BoolVarP(&_backupDescribeBackupVault, "describe-backup-vault", "", false, "Describe Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupDescribeCopyJob, "describe-copy-job", "", false, "Describe Copy Job")
	_backupCmd.Flags().BoolVarP(&_backupDescribeFramework, "describe-framework", "", false, "Describe Framework")
	_backupCmd.Flags().BoolVarP(&_backupDescribeGlobalSettings, "describe-global-settings", "", false, "Describe Global Settings")
	_backupCmd.Flags().BoolVarP(&_backupDescribeProtectedResource, "describe-protected-resource", "", false, "Describe Protected Resource")
	_backupCmd.Flags().BoolVarP(&_backupDescribeRecoveryPoint, "describe-recovery-point", "", false, "Describe Recovery Point")
	_backupCmd.Flags().BoolVarP(&_backupDescribeRegionSettings, "describe-region-settings", "", false, "Describe Region Settings")
	_backupCmd.Flags().BoolVarP(&_backupDescribeReportJob, "describe-report-job", "", false, "Describe Report Job")
	_backupCmd.Flags().BoolVarP(&_backupDescribeReportPlan, "describe-report-plan", "", false, "Describe Report Plan")
	_backupCmd.Flags().BoolVarP(&_backupDescribeRestoreJob, "describe-restore-job", "", false, "Describe Restore Job")
	_backupCmd.Flags().BoolVarP(&_backupDescribeScanJob, "describe-scan-job", "", false, "Describe Scan Job")
	_backupCmd.Flags().BoolVarP(&_backupDisassociateBackupVaultMpaApprovalTeam, "disassociate-backup-vault-mpa-approval-team", "", false, "Disassociate Backup Vault Mpa Approval Team")
	_backupCmd.Flags().BoolVarP(&_backupDisassociateRecoveryPoint, "disassociate-recovery-point", "", false, "Disassociate Recovery Point")
	_backupCmd.Flags().BoolVarP(&_backupDisassociateRecoveryPointFromParent, "disassociate-recovery-point-from-parent", "", false, "Disassociate Recovery Point From Parent")
	_backupCmd.Flags().BoolVarP(&_backupExportBackupPlanTemplate, "export-backup-plan-template", "", false, "Export Backup Plan Template")
	_backupCmd.Flags().BoolVarP(&_backupGetBackupPlan, "get-backup-plan", "", false, "Get Backup Plan")
	_backupCmd.Flags().BoolVarP(&_backupGetBackupPlanFromJSON, "get-backup-plan-from-json", "", false, "Get Backup Plan From JSON")
	_backupCmd.Flags().BoolVarP(&_backupGetBackupPlanFromTemplate, "get-backup-plan-from-template", "", false, "Get Backup Plan From Template")
	_backupCmd.Flags().BoolVarP(&_backupGetBackupSelection, "get-backup-selection", "", false, "Get Backup Selection")
	_backupCmd.Flags().BoolVarP(&_backupGetBackupVaultAccessPolicy, "get-backup-vault-access-policy", "", false, "Get Backup Vault Access Policy")
	_backupCmd.Flags().BoolVarP(&_backupGetBackupVaultNotifications, "get-backup-vault-notifications", "", false, "Get Backup Vault Notifications")
	_backupCmd.Flags().BoolVarP(&_backupGetLegalHold, "get-legal-hold", "", false, "Get Legal Hold")
	_backupCmd.Flags().BoolVarP(&_backupGetRecoveryPointIndexDetails, "get-recovery-point-index-details", "", false, "Get Recovery Point Index Details")
	_backupCmd.Flags().BoolVarP(&_backupGetRecoveryPointRestoreMetadata, "get-recovery-point-restore-metadata", "", false, "Get Recovery Point Restore Metadata")
	_backupCmd.Flags().BoolVarP(&_backupGetRestoreJobMetadata, "get-restore-job-metadata", "", false, "Get Restore Job Metadata")
	_backupCmd.Flags().BoolVarP(&_backupGetRestoreTestingInferredMetadata, "get-restore-testing-inferred-metadata", "", false, "Get Restore Testing Inferred Metadata")
	_backupCmd.Flags().BoolVarP(&_backupGetRestoreTestingPlan, "get-restore-testing-plan", "", false, "Get Restore Testing Plan")
	_backupCmd.Flags().BoolVarP(&_backupGetRestoreTestingSelection, "get-restore-testing-selection", "", false, "Get Restore Testing Selection")
	_backupCmd.Flags().BoolVarP(&_backupGetSupportedResourceTypes, "get-supported-resource-types", "", false, "Get Supported Resource Types")
	_backupCmd.Flags().BoolVarP(&_backupGetTieringConfiguration, "get-tiering-configuration", "", false, "Get Tiering Configuration")
	_backupCmd.Flags().BoolVarP(&_backupListBackupJobSummaries, "list-backup-job-summaries", "", false, "List Backup Job Summaries")
	_backupCmd.Flags().BoolVarP(&_backupListBackupJobs, "list-backup-jobs", "", false, "List Backup Jobs")
	_backupCmd.Flags().BoolVarP(&_backupListBackupPlanTemplates, "list-backup-plan-templates", "", false, "List Backup Plan Templates")
	_backupCmd.Flags().BoolVarP(&_backupListBackupPlanVersions, "list-backup-plan-versions", "", false, "List Backup Plan Versions")
	_backupCmd.Flags().BoolVarP(&_backupListBackupPlans, "list-backup-plans", "", false, "List Backup Plans")
	_backupCmd.Flags().BoolVarP(&_backupListBackupSelections, "list-backup-selections", "", false, "List Backup Selections")
	_backupCmd.Flags().BoolVarP(&_backupListBackupVaults, "list-backup-vaults", "", false, "List Backup Vaults")
	_backupCmd.Flags().BoolVarP(&_backupListCopyJobSummaries, "list-copy-job-summaries", "", false, "List Copy Job Summaries")
	_backupCmd.Flags().BoolVarP(&_backupListCopyJobs, "list-copy-jobs", "", false, "List Copy Jobs")
	_backupCmd.Flags().BoolVarP(&_backupListFrameworks, "list-frameworks", "", false, "List Frameworks")
	_backupCmd.Flags().BoolVarP(&_backupListIndexedRecoveryPoints, "list-indexed-recovery-points", "", false, "List Indexed Recovery Points")
	_backupCmd.Flags().BoolVarP(&_backupListLegalHolds, "list-legal-holds", "", false, "List Legal Holds")
	_backupCmd.Flags().BoolVarP(&_backupListProtectedResources, "list-protected-resources", "", false, "List Protected Resources")
	_backupCmd.Flags().BoolVarP(&_backupListProtectedResourcesByBackupVault, "list-protected-resources-by-backup-vault", "", false, "List Protected Resources By Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupListRecoveryPointsByBackupVault, "list-recovery-points-by-backup-vault", "", false, "List Recovery Points By Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupListRecoveryPointsByLegalHold, "list-recovery-points-by-legal-hold", "", false, "List Recovery Points By Legal Hold")
	_backupCmd.Flags().BoolVarP(&_backupListRecoveryPointsByResource, "list-recovery-points-by-resource", "", false, "List Recovery Points By Resource")
	_backupCmd.Flags().BoolVarP(&_backupListReportJobs, "list-report-jobs", "", false, "List Report Jobs")
	_backupCmd.Flags().BoolVarP(&_backupListReportPlans, "list-report-plans", "", false, "List Report Plans")
	_backupCmd.Flags().BoolVarP(&_backupListRestoreAccessBackupVaults, "list-restore-access-backup-vaults", "", false, "List Restore Access Backup Vaults")
	_backupCmd.Flags().BoolVarP(&_backupListRestoreJobSummaries, "list-restore-job-summaries", "", false, "List Restore Job Summaries")
	_backupCmd.Flags().BoolVarP(&_backupListRestoreJobs, "list-restore-jobs", "", false, "List Restore Jobs")
	_backupCmd.Flags().BoolVarP(&_backupListRestoreJobsByProtectedResource, "list-restore-jobs-by-protected-resource", "", false, "List Restore Jobs By Protected Resource")
	_backupCmd.Flags().BoolVarP(&_backupListRestoreTestingPlans, "list-restore-testing-plans", "", false, "List Restore Testing Plans")
	_backupCmd.Flags().BoolVarP(&_backupListRestoreTestingSelections, "list-restore-testing-selections", "", false, "List Restore Testing Selections")
	_backupCmd.Flags().BoolVarP(&_backupListScanJobSummaries, "list-scan-job-summaries", "", false, "List Scan Job Summaries")
	_backupCmd.Flags().BoolVarP(&_backupListScanJobs, "list-scan-jobs", "", false, "List Scan Jobs")
	_backupCmd.Flags().BoolVarP(&_backupListTags, "list-tags", "", false, "List Tags")
	_backupCmd.Flags().BoolVarP(&_backupListTieringConfigurations, "list-tiering-configurations", "", false, "List Tiering Configurations")
	_backupCmd.Flags().BoolVarP(&_backupPutBackupVaultAccessPolicy, "put-backup-vault-access-policy", "", false, "Put Backup Vault Access Policy")
	_backupCmd.Flags().BoolVarP(&_backupPutBackupVaultLockConfiguration, "put-backup-vault-lock-configuration", "", false, "Put Backup Vault Lock Configuration")
	_backupCmd.Flags().BoolVarP(&_backupPutBackupVaultNotifications, "put-backup-vault-notifications", "", false, "Put Backup Vault Notifications")
	_backupCmd.Flags().BoolVarP(&_backupPutRestoreValidationResult, "put-restore-validation-result", "", false, "Put Restore Validation Result")
	_backupCmd.Flags().BoolVarP(&_backupRevokeRestoreAccessBackupVault, "revoke-restore-access-backup-vault", "", false, "Revoke Restore Access Backup Vault")
	_backupCmd.Flags().BoolVarP(&_backupStartBackupJob, "start-backup-job", "", false, "Start Backup Job")
	_backupCmd.Flags().BoolVarP(&_backupStartCopyJob, "start-copy-job", "", false, "Start Copy Job")
	_backupCmd.Flags().BoolVarP(&_backupStartReportJob, "start-report-job", "", false, "Start Report Job")
	_backupCmd.Flags().BoolVarP(&_backupStartRestoreJob, "start-restore-job", "", false, "Start Restore Job")
	_backupCmd.Flags().BoolVarP(&_backupStartScanJob, "start-scan-job", "", false, "Start Scan Job")
	_backupCmd.Flags().BoolVarP(&_backupStopBackupJob, "stop-backup-job", "", false, "Stop Backup Job")
	_backupCmd.Flags().BoolVarP(&_backupTagResource, "tag-resource", "", false, "Tag Resource")
	_backupCmd.Flags().BoolVarP(&_backupUntagResource, "untag-resource", "", false, "Untag Resource")
	_backupCmd.Flags().BoolVarP(&_backupUpdateBackupPlan, "update-backup-plan", "", false, "Update Backup Plan")
	_backupCmd.Flags().BoolVarP(&_backupUpdateFramework, "update-framework", "", false, "Update Framework")
	_backupCmd.Flags().BoolVarP(&_backupUpdateGlobalSettings, "update-global-settings", "", false, "Update Global Settings")
	_backupCmd.Flags().BoolVarP(&_backupUpdateRecoveryPointIndexSettings, "update-recovery-point-index-settings", "", false, "Update Recovery Point Index Settings")
	_backupCmd.Flags().BoolVarP(&_backupUpdateRecoveryPointLifecycle, "update-recovery-point-lifecycle", "", false, "Update Recovery Point Lifecycle")
	_backupCmd.Flags().BoolVarP(&_backupUpdateRegionSettings, "update-region-settings", "", false, "Update Region Settings")
	_backupCmd.Flags().BoolVarP(&_backupUpdateReportPlan, "update-report-plan", "", false, "Update Report Plan")
	_backupCmd.Flags().BoolVarP(&_backupUpdateRestoreTestingPlan, "update-restore-testing-plan", "", false, "Update Restore Testing Plan")
	_backupCmd.Flags().BoolVarP(&_backupUpdateRestoreTestingSelection, "update-restore-testing-selection", "", false, "Update Restore Testing Selection")
	_backupCmd.Flags().BoolVarP(&_backupUpdateTieringConfiguration, "update-tiering-configuration", "", false, "Update Tiering Configuration")

}
