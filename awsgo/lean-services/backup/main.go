package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/backup"
)

var fields_associate_backup_vault_mpa_approval_team = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "MpaApprovalTeamArn", Flag: "mpa-approval-team-arn", Type: "*string", Required: true},
	{Name: "RequesterComment", Flag: "requester-comment", Type: "*string", Required: false},
}

var fields_cancel_legal_hold = []leanruntime.Field{
	{Name: "CancelDescription", Flag: "cancel-description", Type: "*string", Required: true},
	{Name: "LegalHoldId", Flag: "legal-hold-id", Type: "*string", Required: true},
	{Name: "RetainRecordInDays", Flag: "retain-record-in-days", Type: "*int64", Required: false},
}

var fields_create_backup_plan = []leanruntime.Field{
	{Name: "BackupPlan", Flag: "backup-plan", Type: "*types.BackupPlanInput", Required: true},
	{Name: "BackupPlanTags", Flag: "backup-plan-tags", Type: "map[string]string", Required: false},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
}

var fields_create_backup_selection = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
	{Name: "BackupSelection", Flag: "backup-selection", Type: "*types.BackupSelection", Required: true},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
}

var fields_create_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "BackupVaultTags", Flag: "backup-vault-tags", Type: "map[string]string", Required: false},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
}

var fields_create_framework = []leanruntime.Field{
	{Name: "FrameworkControls", Flag: "framework-controls", Type: "[]types.FrameworkControl", Required: true},
	{Name: "FrameworkDescription", Flag: "framework-description", Type: "*string", Required: false},
	{Name: "FrameworkName", Flag: "framework-name", Type: "*string", Required: true},
	{Name: "FrameworkTags", Flag: "framework-tags", Type: "map[string]string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
}

var fields_create_legal_hold = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "RecoveryPointSelection", Flag: "recovery-point-selection", Type: "*types.RecoveryPointSelection", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_create_logically_air_gapped_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "BackupVaultTags", Flag: "backup-vault-tags", Type: "map[string]string", Required: false},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "MaxRetentionDays", Flag: "max-retention-days", Type: "*int64", Required: true},
	{Name: "MinRetentionDays", Flag: "min-retention-days", Type: "*int64", Required: true},
}

var fields_create_report_plan = []leanruntime.Field{
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ReportDeliveryChannel", Flag: "report-delivery-channel", Type: "*types.ReportDeliveryChannel", Required: true},
	{Name: "ReportPlanDescription", Flag: "report-plan-description", Type: "*string", Required: false},
	{Name: "ReportPlanName", Flag: "report-plan-name", Type: "*string", Required: true},
	{Name: "ReportPlanTags", Flag: "report-plan-tags", Type: "map[string]string", Required: false},
	{Name: "ReportSetting", Flag: "report-setting", Type: "*types.ReportSetting", Required: true},
}

var fields_create_restore_access_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: false},
	{Name: "BackupVaultTags", Flag: "backup-vault-tags", Type: "map[string]string", Required: false},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "RequesterComment", Flag: "requester-comment", Type: "*string", Required: false},
	{Name: "SourceBackupVaultArn", Flag: "source-backup-vault-arn", Type: "*string", Required: true},
}

var fields_create_restore_testing_plan = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "RestoreTestingPlan", Flag: "restore-testing-plan", Type: "*types.RestoreTestingPlanForCreate", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_restore_testing_selection = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
	{Name: "RestoreTestingSelection", Flag: "restore-testing-selection", Type: "*types.RestoreTestingSelectionForCreate", Required: true},
}

var fields_create_tiering_configuration = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "TieringConfiguration", Flag: "tiering-configuration", Type: "*types.TieringConfigurationInputForCreate", Required: true},
	{Name: "TieringConfigurationTags", Flag: "tiering-configuration-tags", Type: "map[string]string", Required: false},
}

var fields_delete_backup_plan = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
}

var fields_delete_backup_selection = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
	{Name: "SelectionId", Flag: "selection-id", Type: "*string", Required: true},
}

var fields_delete_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_delete_backup_vault_access_policy = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_delete_backup_vault_lock_configuration = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_delete_backup_vault_notifications = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_delete_framework = []leanruntime.Field{
	{Name: "FrameworkName", Flag: "framework-name", Type: "*string", Required: true},
}

var fields_delete_recovery_point = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_delete_report_plan = []leanruntime.Field{
	{Name: "ReportPlanName", Flag: "report-plan-name", Type: "*string", Required: true},
}

var fields_delete_restore_testing_plan = []leanruntime.Field{
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
}

var fields_delete_restore_testing_selection = []leanruntime.Field{
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
	{Name: "RestoreTestingSelectionName", Flag: "restore-testing-selection-name", Type: "*string", Required: true},
}

var fields_delete_tiering_configuration = []leanruntime.Field{
	{Name: "TieringConfigurationName", Flag: "tiering-configuration-name", Type: "*string", Required: true},
}

var fields_describe_backup_job = []leanruntime.Field{
	{Name: "BackupJobId", Flag: "backup-job-id", Type: "*string", Required: true},
}

var fields_describe_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultAccountId", Flag: "backup-vault-account-id", Type: "*string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_describe_copy_job = []leanruntime.Field{
	{Name: "CopyJobId", Flag: "copy-job-id", Type: "*string", Required: true},
}

var fields_describe_framework = []leanruntime.Field{
	{Name: "FrameworkName", Flag: "framework-name", Type: "*string", Required: true},
}

var fields_describe_global_settings = []leanruntime.Field{}

var fields_describe_protected_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_recovery_point = []leanruntime.Field{
	{Name: "BackupVaultAccountId", Flag: "backup-vault-account-id", Type: "*string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_describe_region_settings = []leanruntime.Field{}

var fields_describe_report_job = []leanruntime.Field{
	{Name: "ReportJobId", Flag: "report-job-id", Type: "*string", Required: true},
}

var fields_describe_report_plan = []leanruntime.Field{
	{Name: "ReportPlanName", Flag: "report-plan-name", Type: "*string", Required: true},
}

var fields_describe_restore_job = []leanruntime.Field{
	{Name: "RestoreJobId", Flag: "restore-job-id", Type: "*string", Required: true},
}

var fields_describe_scan_job = []leanruntime.Field{
	{Name: "ScanJobId", Flag: "scan-job-id", Type: "*string", Required: true},
}

var fields_disassociate_backup_vault_mpa_approval_team = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RequesterComment", Flag: "requester-comment", Type: "*string", Required: false},
}

var fields_disassociate_recovery_point = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_disassociate_recovery_point_from_parent = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_export_backup_plan_template = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
}

var fields_get_backup_plan = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
	{Name: "MaxScheduledRunsPreview", Flag: "max-scheduled-runs-preview", Type: "int32", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_backup_plan_from_json = []leanruntime.Field{
	{Name: "BackupPlanTemplateJson", Flag: "backup-plan-template-json", Type: "*string", Required: true},
}

var fields_get_backup_plan_from_template = []leanruntime.Field{
	{Name: "BackupPlanTemplateId", Flag: "backup-plan-template-id", Type: "*string", Required: true},
}

var fields_get_backup_selection = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
	{Name: "SelectionId", Flag: "selection-id", Type: "*string", Required: true},
}

var fields_get_backup_vault_access_policy = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_get_backup_vault_notifications = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
}

var fields_get_legal_hold = []leanruntime.Field{
	{Name: "LegalHoldId", Flag: "legal-hold-id", Type: "*string", Required: true},
}

var fields_get_recovery_point_index_details = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_get_recovery_point_restore_metadata = []leanruntime.Field{
	{Name: "BackupVaultAccountId", Flag: "backup-vault-account-id", Type: "*string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_get_restore_job_metadata = []leanruntime.Field{
	{Name: "RestoreJobId", Flag: "restore-job-id", Type: "*string", Required: true},
}

var fields_get_restore_testing_inferred_metadata = []leanruntime.Field{
	{Name: "BackupVaultAccountId", Flag: "backup-vault-account-id", Type: "*string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_get_restore_testing_plan = []leanruntime.Field{
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
}

var fields_get_restore_testing_selection = []leanruntime.Field{
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
	{Name: "RestoreTestingSelectionName", Flag: "restore-testing-selection-name", Type: "*string", Required: true},
}

var fields_get_supported_resource_types = []leanruntime.Field{}

var fields_get_tiering_configuration = []leanruntime.Field{
	{Name: "TieringConfigurationName", Flag: "tiering-configuration-name", Type: "*string", Required: true},
}

var fields_list_backup_job_summaries = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AggregationPeriod", Flag: "aggregation-period", Type: "types.AggregationPeriod", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MessageCategory", Flag: "message-category", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.BackupJobStatus", Required: false},
}

var fields_list_backup_jobs = []leanruntime.Field{
	{Name: "ByAccountId", Flag: "by-account-id", Type: "*string", Required: false},
	{Name: "ByBackupVaultName", Flag: "by-backup-vault-name", Type: "*string", Required: false},
	{Name: "ByCompleteAfter", Flag: "by-complete-after", Type: "*time.Time", Required: false},
	{Name: "ByCompleteBefore", Flag: "by-complete-before", Type: "*time.Time", Required: false},
	{Name: "ByCreatedAfter", Flag: "by-created-after", Type: "*time.Time", Required: false},
	{Name: "ByCreatedBefore", Flag: "by-created-before", Type: "*time.Time", Required: false},
	{Name: "ByMessageCategory", Flag: "by-message-category", Type: "*string", Required: false},
	{Name: "ByParentJobId", Flag: "by-parent-job-id", Type: "*string", Required: false},
	{Name: "ByResourceArn", Flag: "by-resource-arn", Type: "*string", Required: false},
	{Name: "ByResourceType", Flag: "by-resource-type", Type: "*string", Required: false},
	{Name: "ByState", Flag: "by-state", Type: "types.BackupJobState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_backup_plan_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_backup_plan_versions = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_backup_plans = []leanruntime.Field{
	{Name: "IncludeDeleted", Flag: "include-deleted", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_backup_selections = []leanruntime.Field{
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_backup_vaults = []leanruntime.Field{
	{Name: "ByShared", Flag: "by-shared", Type: "bool", Required: false},
	{Name: "ByVaultType", Flag: "by-vault-type", Type: "types.VaultType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_copy_job_summaries = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AggregationPeriod", Flag: "aggregation-period", Type: "types.AggregationPeriod", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MessageCategory", Flag: "message-category", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.CopyJobStatus", Required: false},
}

var fields_list_copy_jobs = []leanruntime.Field{
	{Name: "ByAccountId", Flag: "by-account-id", Type: "*string", Required: false},
	{Name: "ByCompleteAfter", Flag: "by-complete-after", Type: "*time.Time", Required: false},
	{Name: "ByCompleteBefore", Flag: "by-complete-before", Type: "*time.Time", Required: false},
	{Name: "ByCreatedAfter", Flag: "by-created-after", Type: "*time.Time", Required: false},
	{Name: "ByCreatedBefore", Flag: "by-created-before", Type: "*time.Time", Required: false},
	{Name: "ByDestinationVaultArn", Flag: "by-destination-vault-arn", Type: "*string", Required: false},
	{Name: "ByMessageCategory", Flag: "by-message-category", Type: "*string", Required: false},
	{Name: "ByParentJobId", Flag: "by-parent-job-id", Type: "*string", Required: false},
	{Name: "ByResourceArn", Flag: "by-resource-arn", Type: "*string", Required: false},
	{Name: "ByResourceType", Flag: "by-resource-type", Type: "*string", Required: false},
	{Name: "BySourceRecoveryPointArn", Flag: "by-source-recovery-point-arn", Type: "*string", Required: false},
	{Name: "ByState", Flag: "by-state", Type: "types.CopyJobState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_frameworks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_indexed_recovery_points = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "IndexStatus", Flag: "index-status", Type: "types.IndexStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "SourceResourceArn", Flag: "source-resource-arn", Type: "*string", Required: false},
}

var fields_list_legal_holds = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_protected_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_protected_resources_by_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultAccountId", Flag: "backup-vault-account-id", Type: "*string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recovery_points_by_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultAccountId", Flag: "backup-vault-account-id", Type: "*string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "ByBackupPlanId", Flag: "by-backup-plan-id", Type: "*string", Required: false},
	{Name: "ByCreatedAfter", Flag: "by-created-after", Type: "*time.Time", Required: false},
	{Name: "ByCreatedBefore", Flag: "by-created-before", Type: "*time.Time", Required: false},
	{Name: "ByParentRecoveryPointArn", Flag: "by-parent-recovery-point-arn", Type: "*string", Required: false},
	{Name: "ByResourceArn", Flag: "by-resource-arn", Type: "*string", Required: false},
	{Name: "ByResourceType", Flag: "by-resource-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recovery_points_by_legal_hold = []leanruntime.Field{
	{Name: "LegalHoldId", Flag: "legal-hold-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recovery_points_by_resource = []leanruntime.Field{
	{Name: "ManagedByAWSBackupOnly", Flag: "managed-by-aws-backup-only", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_report_jobs = []leanruntime.Field{
	{Name: "ByCreationAfter", Flag: "by-creation-after", Type: "*time.Time", Required: false},
	{Name: "ByCreationBefore", Flag: "by-creation-before", Type: "*time.Time", Required: false},
	{Name: "ByReportPlanName", Flag: "by-report-plan-name", Type: "*string", Required: false},
	{Name: "ByStatus", Flag: "by-status", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_report_plans = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_restore_access_backup_vaults = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_restore_job_summaries = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AggregationPeriod", Flag: "aggregation-period", Type: "types.AggregationPeriod", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.RestoreJobState", Required: false},
}

var fields_list_restore_jobs = []leanruntime.Field{
	{Name: "ByAccountId", Flag: "by-account-id", Type: "*string", Required: false},
	{Name: "ByCompleteAfter", Flag: "by-complete-after", Type: "*time.Time", Required: false},
	{Name: "ByCompleteBefore", Flag: "by-complete-before", Type: "*time.Time", Required: false},
	{Name: "ByCreatedAfter", Flag: "by-created-after", Type: "*time.Time", Required: false},
	{Name: "ByCreatedBefore", Flag: "by-created-before", Type: "*time.Time", Required: false},
	{Name: "ByParentJobId", Flag: "by-parent-job-id", Type: "*string", Required: false},
	{Name: "ByResourceType", Flag: "by-resource-type", Type: "*string", Required: false},
	{Name: "ByRestoreTestingPlanArn", Flag: "by-restore-testing-plan-arn", Type: "*string", Required: false},
	{Name: "ByStatus", Flag: "by-status", Type: "types.RestoreJobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_restore_jobs_by_protected_resource = []leanruntime.Field{
	{Name: "ByRecoveryPointCreationDateAfter", Flag: "by-recovery-point-creation-date-after", Type: "*time.Time", Required: false},
	{Name: "ByRecoveryPointCreationDateBefore", Flag: "by-recovery-point-creation-date-before", Type: "*time.Time", Required: false},
	{Name: "ByStatus", Flag: "by-status", Type: "types.RestoreJobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_restore_testing_plans = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_restore_testing_selections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
}

var fields_list_scan_job_summaries = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AggregationPeriod", Flag: "aggregation-period", Type: "types.AggregationPeriod", Required: false},
	{Name: "MalwareScanner", Flag: "malware-scanner", Type: "types.MalwareScanner", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "ScanResultStatus", Flag: "scan-result-status", Type: "types.ScanResultStatus", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScanJobStatus", Required: false},
}

var fields_list_scan_jobs = []leanruntime.Field{
	{Name: "ByAccountId", Flag: "by-account-id", Type: "*string", Required: false},
	{Name: "ByBackupVaultName", Flag: "by-backup-vault-name", Type: "*string", Required: false},
	{Name: "ByCompleteAfter", Flag: "by-complete-after", Type: "*time.Time", Required: false},
	{Name: "ByCompleteBefore", Flag: "by-complete-before", Type: "*time.Time", Required: false},
	{Name: "ByMalwareScanner", Flag: "by-malware-scanner", Type: "types.MalwareScanner", Required: false},
	{Name: "ByRecoveryPointArn", Flag: "by-recovery-point-arn", Type: "*string", Required: false},
	{Name: "ByResourceArn", Flag: "by-resource-arn", Type: "*string", Required: false},
	{Name: "ByResourceType", Flag: "by-resource-type", Type: "types.ScanResourceType", Required: false},
	{Name: "ByScanResultStatus", Flag: "by-scan-result-status", Type: "types.ScanResultStatus", Required: false},
	{Name: "ByState", Flag: "by-state", Type: "types.ScanState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tiering_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_backup_vault_access_policy = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
}

var fields_put_backup_vault_lock_configuration = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "ChangeableForDays", Flag: "changeable-for-days", Type: "*int64", Required: false},
	{Name: "MaxRetentionDays", Flag: "max-retention-days", Type: "*int64", Required: false},
	{Name: "MinRetentionDays", Flag: "min-retention-days", Type: "*int64", Required: false},
}

var fields_put_backup_vault_notifications = []leanruntime.Field{
	{Name: "BackupVaultEvents", Flag: "backup-vault-events", Type: "[]types.BackupVaultEvent", Required: true},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "SNSTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: true},
}

var fields_put_restore_validation_result = []leanruntime.Field{
	{Name: "RestoreJobId", Flag: "restore-job-id", Type: "*string", Required: true},
	{Name: "ValidationStatus", Flag: "validation-status", Type: "types.RestoreValidationStatus", Required: true},
	{Name: "ValidationStatusMessage", Flag: "validation-status-message", Type: "*string", Required: false},
}

var fields_revoke_restore_access_backup_vault = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "RequesterComment", Flag: "requester-comment", Type: "*string", Required: false},
	{Name: "RestoreAccessBackupVaultArn", Flag: "restore-access-backup-vault-arn", Type: "*string", Required: true},
}

var fields_start_backup_job = []leanruntime.Field{
	{Name: "BackupOptions", Flag: "backup-options", Type: "map[string]string", Required: false},
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "CompleteWindowMinutes", Flag: "complete-window-minutes", Type: "*int64", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Index", Flag: "index", Type: "types.Index", Required: false},
	{Name: "Lifecycle", Flag: "lifecycle", Type: "*types.Lifecycle", Required: false},
	{Name: "LogicallyAirGappedBackupVaultArn", Flag: "logically-air-gapped-backup-vault-arn", Type: "*string", Required: false},
	{Name: "RecoveryPointTags", Flag: "recovery-point-tags", Type: "map[string]string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StartWindowMinutes", Flag: "start-window-minutes", Type: "*int64", Required: false},
}

var fields_start_copy_job = []leanruntime.Field{
	{Name: "DestinationBackupVaultArn", Flag: "destination-backup-vault-arn", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Lifecycle", Flag: "lifecycle", Type: "*types.Lifecycle", Required: false},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
	{Name: "SourceBackupVaultName", Flag: "source-backup-vault-name", Type: "*string", Required: true},
}

var fields_start_report_job = []leanruntime.Field{
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ReportPlanName", Flag: "report-plan-name", Type: "*string", Required: true},
}

var fields_start_restore_job = []leanruntime.Field{
	{Name: "CopySourceTagsToRestoredResource", Flag: "copy-source-tags-to-restored-resource", Type: "bool", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_start_scan_job = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "MalwareScanner", Flag: "malware-scanner", Type: "types.MalwareScanner", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
	{Name: "ScanBaseRecoveryPointArn", Flag: "scan-base-recovery-point-arn", Type: "*string", Required: false},
	{Name: "ScanMode", Flag: "scan-mode", Type: "types.ScanMode", Required: true},
	{Name: "ScannerRoleArn", Flag: "scanner-role-arn", Type: "*string", Required: true},
}

var fields_stop_backup_job = []leanruntime.Field{
	{Name: "BackupJobId", Flag: "backup-job-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeyList", Flag: "tag-key-list", Type: "[]string", Required: true},
}

var fields_update_backup_plan = []leanruntime.Field{
	{Name: "BackupPlan", Flag: "backup-plan", Type: "*types.BackupPlanInput", Required: true},
	{Name: "BackupPlanId", Flag: "backup-plan-id", Type: "*string", Required: true},
}

var fields_update_framework = []leanruntime.Field{
	{Name: "FrameworkControls", Flag: "framework-controls", Type: "[]types.FrameworkControl", Required: false},
	{Name: "FrameworkDescription", Flag: "framework-description", Type: "*string", Required: false},
	{Name: "FrameworkName", Flag: "framework-name", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
}

var fields_update_global_settings = []leanruntime.Field{
	{Name: "GlobalSettings", Flag: "global-settings", Type: "map[string]string", Required: false},
}

var fields_update_recovery_point_index_settings = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "Index", Flag: "index", Type: "types.Index", Required: true},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_update_recovery_point_lifecycle = []leanruntime.Field{
	{Name: "BackupVaultName", Flag: "backup-vault-name", Type: "*string", Required: true},
	{Name: "Lifecycle", Flag: "lifecycle", Type: "*types.Lifecycle", Required: false},
	{Name: "RecoveryPointArn", Flag: "recovery-point-arn", Type: "*string", Required: true},
}

var fields_update_region_settings = []leanruntime.Field{
	{Name: "ResourceTypeManagementPreference", Flag: "resource-type-management-preference", Type: "map[string]bool", Required: false},
	{Name: "ResourceTypeOptInPreference", Flag: "resource-type-opt-in-preference", Type: "map[string]bool", Required: false},
}

var fields_update_report_plan = []leanruntime.Field{
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ReportDeliveryChannel", Flag: "report-delivery-channel", Type: "*types.ReportDeliveryChannel", Required: false},
	{Name: "ReportPlanDescription", Flag: "report-plan-description", Type: "*string", Required: false},
	{Name: "ReportPlanName", Flag: "report-plan-name", Type: "*string", Required: true},
	{Name: "ReportSetting", Flag: "report-setting", Type: "*types.ReportSetting", Required: false},
}

var fields_update_restore_testing_plan = []leanruntime.Field{
	{Name: "RestoreTestingPlan", Flag: "restore-testing-plan", Type: "*types.RestoreTestingPlanForUpdate", Required: true},
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
}

var fields_update_restore_testing_selection = []leanruntime.Field{
	{Name: "RestoreTestingPlanName", Flag: "restore-testing-plan-name", Type: "*string", Required: true},
	{Name: "RestoreTestingSelection", Flag: "restore-testing-selection", Type: "*types.RestoreTestingSelectionForUpdate", Required: true},
	{Name: "RestoreTestingSelectionName", Flag: "restore-testing-selection-name", Type: "*string", Required: true},
}

var fields_update_tiering_configuration = []leanruntime.Field{
	{Name: "TieringConfiguration", Flag: "tiering-configuration", Type: "*types.TieringConfigurationInputForUpdate", Required: true},
	{Name: "TieringConfigurationName", Flag: "tiering-configuration-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-backup-vault-mpa-approval-team": {
			Name:   "associate-backup-vault-mpa-approval-team",
			Fields: fields_associate_backup_vault_mpa_approval_team,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateBackupVaultMpaApprovalTeamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_backup_vault_mpa_approval_team, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateBackupVaultMpaApprovalTeam(ctx, input)
			},
		},
		"cancel-legal-hold": {
			Name:   "cancel-legal-hold",
			Fields: fields_cancel_legal_hold,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelLegalHoldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_legal_hold, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelLegalHold(ctx, input)
			},
		},
		"create-backup-plan": {
			Name:   "create-backup-plan",
			Fields: fields_create_backup_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackupPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backup_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackupPlan(ctx, input)
			},
		},
		"create-backup-selection": {
			Name:   "create-backup-selection",
			Fields: fields_create_backup_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackupSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backup_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackupSelection(ctx, input)
			},
		},
		"create-backup-vault": {
			Name:   "create-backup-vault",
			Fields: fields_create_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackupVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backup_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackupVault(ctx, input)
			},
		},
		"create-framework": {
			Name:   "create-framework",
			Fields: fields_create_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFramework(ctx, input)
			},
		},
		"create-legal-hold": {
			Name:   "create-legal-hold",
			Fields: fields_create_legal_hold,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLegalHoldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_legal_hold, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLegalHold(ctx, input)
			},
		},
		"create-logically-air-gapped-backup-vault": {
			Name:   "create-logically-air-gapped-backup-vault",
			Fields: fields_create_logically_air_gapped_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLogicallyAirGappedBackupVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_logically_air_gapped_backup_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLogicallyAirGappedBackupVault(ctx, input)
			},
		},
		"create-report-plan": {
			Name:   "create-report-plan",
			Fields: fields_create_report_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReportPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_report_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReportPlan(ctx, input)
			},
		},
		"create-restore-access-backup-vault": {
			Name:   "create-restore-access-backup-vault",
			Fields: fields_create_restore_access_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRestoreAccessBackupVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_restore_access_backup_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRestoreAccessBackupVault(ctx, input)
			},
		},
		"create-restore-testing-plan": {
			Name:   "create-restore-testing-plan",
			Fields: fields_create_restore_testing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRestoreTestingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_restore_testing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRestoreTestingPlan(ctx, input)
			},
		},
		"create-restore-testing-selection": {
			Name:   "create-restore-testing-selection",
			Fields: fields_create_restore_testing_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRestoreTestingSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_restore_testing_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRestoreTestingSelection(ctx, input)
			},
		},
		"create-tiering-configuration": {
			Name:   "create-tiering-configuration",
			Fields: fields_create_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTieringConfiguration(ctx, input)
			},
		},
		"delete-backup-plan": {
			Name:   "delete-backup-plan",
			Fields: fields_delete_backup_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackupPlan(ctx, input)
			},
		},
		"delete-backup-selection": {
			Name:   "delete-backup-selection",
			Fields: fields_delete_backup_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackupSelection(ctx, input)
			},
		},
		"delete-backup-vault": {
			Name:   "delete-backup-vault",
			Fields: fields_delete_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackupVault(ctx, input)
			},
		},
		"delete-backup-vault-access-policy": {
			Name:   "delete-backup-vault-access-policy",
			Fields: fields_delete_backup_vault_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupVaultAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup_vault_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackupVaultAccessPolicy(ctx, input)
			},
		},
		"delete-backup-vault-lock-configuration": {
			Name:   "delete-backup-vault-lock-configuration",
			Fields: fields_delete_backup_vault_lock_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupVaultLockConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup_vault_lock_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackupVaultLockConfiguration(ctx, input)
			},
		},
		"delete-backup-vault-notifications": {
			Name:   "delete-backup-vault-notifications",
			Fields: fields_delete_backup_vault_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupVaultNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup_vault_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackupVaultNotifications(ctx, input)
			},
		},
		"delete-framework": {
			Name:   "delete-framework",
			Fields: fields_delete_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFramework(ctx, input)
			},
		},
		"delete-recovery-point": {
			Name:   "delete-recovery-point",
			Fields: fields_delete_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecoveryPoint(ctx, input)
			},
		},
		"delete-report-plan": {
			Name:   "delete-report-plan",
			Fields: fields_delete_report_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReportPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_report_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReportPlan(ctx, input)
			},
		},
		"delete-restore-testing-plan": {
			Name:   "delete-restore-testing-plan",
			Fields: fields_delete_restore_testing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRestoreTestingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_restore_testing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRestoreTestingPlan(ctx, input)
			},
		},
		"delete-restore-testing-selection": {
			Name:   "delete-restore-testing-selection",
			Fields: fields_delete_restore_testing_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRestoreTestingSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_restore_testing_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRestoreTestingSelection(ctx, input)
			},
		},
		"delete-tiering-configuration": {
			Name:   "delete-tiering-configuration",
			Fields: fields_delete_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTieringConfiguration(ctx, input)
			},
		},
		"describe-backup-job": {
			Name:   "describe-backup-job",
			Fields: fields_describe_backup_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBackupJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_backup_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBackupJob(ctx, input)
			},
		},
		"describe-backup-vault": {
			Name:   "describe-backup-vault",
			Fields: fields_describe_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBackupVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_backup_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBackupVault(ctx, input)
			},
		},
		"describe-copy-job": {
			Name:   "describe-copy-job",
			Fields: fields_describe_copy_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCopyJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_copy_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCopyJob(ctx, input)
			},
		},
		"describe-framework": {
			Name:   "describe-framework",
			Fields: fields_describe_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFramework(ctx, input)
			},
		},
		"describe-global-settings": {
			Name:   "describe-global-settings",
			Fields: fields_describe_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGlobalSettings(ctx, input)
			},
		},
		"describe-protected-resource": {
			Name:   "describe-protected-resource",
			Fields: fields_describe_protected_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProtectedResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_protected_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProtectedResource(ctx, input)
			},
		},
		"describe-recovery-point": {
			Name:   "describe-recovery-point",
			Fields: fields_describe_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRecoveryPoint(ctx, input)
			},
		},
		"describe-region-settings": {
			Name:   "describe-region-settings",
			Fields: fields_describe_region_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_region_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRegionSettings(ctx, input)
			},
		},
		"describe-report-job": {
			Name:   "describe-report-job",
			Fields: fields_describe_report_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_report_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReportJob(ctx, input)
			},
		},
		"describe-report-plan": {
			Name:   "describe-report-plan",
			Fields: fields_describe_report_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReportPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_report_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReportPlan(ctx, input)
			},
		},
		"describe-restore-job": {
			Name:   "describe-restore-job",
			Fields: fields_describe_restore_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRestoreJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_restore_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRestoreJob(ctx, input)
			},
		},
		"describe-scan-job": {
			Name:   "describe-scan-job",
			Fields: fields_describe_scan_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScanJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scan_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScanJob(ctx, input)
			},
		},
		"disassociate-backup-vault-mpa-approval-team": {
			Name:   "disassociate-backup-vault-mpa-approval-team",
			Fields: fields_disassociate_backup_vault_mpa_approval_team,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateBackupVaultMpaApprovalTeamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_backup_vault_mpa_approval_team, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateBackupVaultMpaApprovalTeam(ctx, input)
			},
		},
		"disassociate-recovery-point": {
			Name:   "disassociate-recovery-point",
			Fields: fields_disassociate_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRecoveryPoint(ctx, input)
			},
		},
		"disassociate-recovery-point-from-parent": {
			Name:   "disassociate-recovery-point-from-parent",
			Fields: fields_disassociate_recovery_point_from_parent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRecoveryPointFromParentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_recovery_point_from_parent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRecoveryPointFromParent(ctx, input)
			},
		},
		"export-backup-plan-template": {
			Name:   "export-backup-plan-template",
			Fields: fields_export_backup_plan_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportBackupPlanTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_backup_plan_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportBackupPlanTemplate(ctx, input)
			},
		},
		"get-backup-plan": {
			Name:   "get-backup-plan",
			Fields: fields_get_backup_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackupPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backup_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackupPlan(ctx, input)
			},
		},
		"get-backup-plan-from-json": {
			Name:   "get-backup-plan-from-json",
			Fields: fields_get_backup_plan_from_json,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackupPlanFromJSONInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backup_plan_from_json, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackupPlanFromJSON(ctx, input)
			},
		},
		"get-backup-plan-from-template": {
			Name:   "get-backup-plan-from-template",
			Fields: fields_get_backup_plan_from_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackupPlanFromTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backup_plan_from_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackupPlanFromTemplate(ctx, input)
			},
		},
		"get-backup-selection": {
			Name:   "get-backup-selection",
			Fields: fields_get_backup_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackupSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backup_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackupSelection(ctx, input)
			},
		},
		"get-backup-vault-access-policy": {
			Name:   "get-backup-vault-access-policy",
			Fields: fields_get_backup_vault_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackupVaultAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backup_vault_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackupVaultAccessPolicy(ctx, input)
			},
		},
		"get-backup-vault-notifications": {
			Name:   "get-backup-vault-notifications",
			Fields: fields_get_backup_vault_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackupVaultNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backup_vault_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackupVaultNotifications(ctx, input)
			},
		},
		"get-legal-hold": {
			Name:   "get-legal-hold",
			Fields: fields_get_legal_hold,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLegalHoldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_legal_hold, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLegalHold(ctx, input)
			},
		},
		"get-recovery-point-index-details": {
			Name:   "get-recovery-point-index-details",
			Fields: fields_get_recovery_point_index_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecoveryPointIndexDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recovery_point_index_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecoveryPointIndexDetails(ctx, input)
			},
		},
		"get-recovery-point-restore-metadata": {
			Name:   "get-recovery-point-restore-metadata",
			Fields: fields_get_recovery_point_restore_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecoveryPointRestoreMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recovery_point_restore_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecoveryPointRestoreMetadata(ctx, input)
			},
		},
		"get-restore-job-metadata": {
			Name:   "get-restore-job-metadata",
			Fields: fields_get_restore_job_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRestoreJobMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_restore_job_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRestoreJobMetadata(ctx, input)
			},
		},
		"get-restore-testing-inferred-metadata": {
			Name:   "get-restore-testing-inferred-metadata",
			Fields: fields_get_restore_testing_inferred_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRestoreTestingInferredMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_restore_testing_inferred_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRestoreTestingInferredMetadata(ctx, input)
			},
		},
		"get-restore-testing-plan": {
			Name:   "get-restore-testing-plan",
			Fields: fields_get_restore_testing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRestoreTestingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_restore_testing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRestoreTestingPlan(ctx, input)
			},
		},
		"get-restore-testing-selection": {
			Name:   "get-restore-testing-selection",
			Fields: fields_get_restore_testing_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRestoreTestingSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_restore_testing_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRestoreTestingSelection(ctx, input)
			},
		},
		"get-supported-resource-types": {
			Name:   "get-supported-resource-types",
			Fields: fields_get_supported_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSupportedResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_supported_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSupportedResourceTypes(ctx, input)
			},
		},
		"get-tiering-configuration": {
			Name:   "get-tiering-configuration",
			Fields: fields_get_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTieringConfiguration(ctx, input)
			},
		},
		"list-backup-job-summaries": {
			Name:   "list-backup-job-summaries",
			Fields: fields_list_backup_job_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupJobSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_job_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupJobSummaries(ctx, input)
				}
				var results []*svc.ListBackupJobSummariesOutput
				p := svc.NewListBackupJobSummariesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-backup-jobs": {
			Name:   "list-backup-jobs",
			Fields: fields_list_backup_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupJobs(ctx, input)
				}
				var results []*svc.ListBackupJobsOutput
				p := svc.NewListBackupJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-backup-plan-templates": {
			Name:   "list-backup-plan-templates",
			Fields: fields_list_backup_plan_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupPlanTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_plan_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupPlanTemplates(ctx, input)
				}
				var results []*svc.ListBackupPlanTemplatesOutput
				p := svc.NewListBackupPlanTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-backup-plan-versions": {
			Name:   "list-backup-plan-versions",
			Fields: fields_list_backup_plan_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupPlanVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_plan_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupPlanVersions(ctx, input)
				}
				var results []*svc.ListBackupPlanVersionsOutput
				p := svc.NewListBackupPlanVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-backup-plans": {
			Name:   "list-backup-plans",
			Fields: fields_list_backup_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupPlans(ctx, input)
				}
				var results []*svc.ListBackupPlansOutput
				p := svc.NewListBackupPlansPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-backup-selections": {
			Name:   "list-backup-selections",
			Fields: fields_list_backup_selections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupSelectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_selections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupSelections(ctx, input)
				}
				var results []*svc.ListBackupSelectionsOutput
				p := svc.NewListBackupSelectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-backup-vaults": {
			Name:   "list-backup-vaults",
			Fields: fields_list_backup_vaults,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackupVaultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_backup_vaults, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBackupVaults(ctx, input)
				}
				var results []*svc.ListBackupVaultsOutput
				p := svc.NewListBackupVaultsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-copy-job-summaries": {
			Name:   "list-copy-job-summaries",
			Fields: fields_list_copy_job_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCopyJobSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_copy_job_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCopyJobSummaries(ctx, input)
				}
				var results []*svc.ListCopyJobSummariesOutput
				p := svc.NewListCopyJobSummariesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-copy-jobs": {
			Name:   "list-copy-jobs",
			Fields: fields_list_copy_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCopyJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_copy_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCopyJobs(ctx, input)
				}
				var results []*svc.ListCopyJobsOutput
				p := svc.NewListCopyJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-frameworks": {
			Name:   "list-frameworks",
			Fields: fields_list_frameworks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFrameworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_frameworks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFrameworks(ctx, input)
				}
				var results []*svc.ListFrameworksOutput
				p := svc.NewListFrameworksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-indexed-recovery-points": {
			Name:   "list-indexed-recovery-points",
			Fields: fields_list_indexed_recovery_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndexedRecoveryPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indexed_recovery_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndexedRecoveryPoints(ctx, input)
				}
				var results []*svc.ListIndexedRecoveryPointsOutput
				p := svc.NewListIndexedRecoveryPointsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-legal-holds": {
			Name:   "list-legal-holds",
			Fields: fields_list_legal_holds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLegalHoldsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_legal_holds, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLegalHolds(ctx, input)
				}
				var results []*svc.ListLegalHoldsOutput
				p := svc.NewListLegalHoldsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-protected-resources": {
			Name:   "list-protected-resources",
			Fields: fields_list_protected_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protected_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtectedResources(ctx, input)
				}
				var results []*svc.ListProtectedResourcesOutput
				p := svc.NewListProtectedResourcesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-protected-resources-by-backup-vault": {
			Name:   "list-protected-resources-by-backup-vault",
			Fields: fields_list_protected_resources_by_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectedResourcesByBackupVaultInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protected_resources_by_backup_vault, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtectedResourcesByBackupVault(ctx, input)
				}
				var results []*svc.ListProtectedResourcesByBackupVaultOutput
				p := svc.NewListProtectedResourcesByBackupVaultPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-recovery-points-by-backup-vault": {
			Name:   "list-recovery-points-by-backup-vault",
			Fields: fields_list_recovery_points_by_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecoveryPointsByBackupVaultInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recovery_points_by_backup_vault, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecoveryPointsByBackupVault(ctx, input)
				}
				var results []*svc.ListRecoveryPointsByBackupVaultOutput
				p := svc.NewListRecoveryPointsByBackupVaultPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-recovery-points-by-legal-hold": {
			Name:   "list-recovery-points-by-legal-hold",
			Fields: fields_list_recovery_points_by_legal_hold,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecoveryPointsByLegalHoldInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recovery_points_by_legal_hold, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecoveryPointsByLegalHold(ctx, input)
				}
				var results []*svc.ListRecoveryPointsByLegalHoldOutput
				p := svc.NewListRecoveryPointsByLegalHoldPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-recovery-points-by-resource": {
			Name:   "list-recovery-points-by-resource",
			Fields: fields_list_recovery_points_by_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecoveryPointsByResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recovery_points_by_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecoveryPointsByResource(ctx, input)
				}
				var results []*svc.ListRecoveryPointsByResourceOutput
				p := svc.NewListRecoveryPointsByResourcePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-report-jobs": {
			Name:   "list-report-jobs",
			Fields: fields_list_report_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_report_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReportJobs(ctx, input)
				}
				var results []*svc.ListReportJobsOutput
				p := svc.NewListReportJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-report-plans": {
			Name:   "list-report-plans",
			Fields: fields_list_report_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_report_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReportPlans(ctx, input)
				}
				var results []*svc.ListReportPlansOutput
				p := svc.NewListReportPlansPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-restore-access-backup-vaults": {
			Name:   "list-restore-access-backup-vaults",
			Fields: fields_list_restore_access_backup_vaults,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRestoreAccessBackupVaultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_restore_access_backup_vaults, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRestoreAccessBackupVaults(ctx, input)
				}
				var results []*svc.ListRestoreAccessBackupVaultsOutput
				p := svc.NewListRestoreAccessBackupVaultsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-restore-job-summaries": {
			Name:   "list-restore-job-summaries",
			Fields: fields_list_restore_job_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRestoreJobSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_restore_job_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRestoreJobSummaries(ctx, input)
				}
				var results []*svc.ListRestoreJobSummariesOutput
				p := svc.NewListRestoreJobSummariesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-restore-jobs": {
			Name:   "list-restore-jobs",
			Fields: fields_list_restore_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRestoreJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_restore_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRestoreJobs(ctx, input)
				}
				var results []*svc.ListRestoreJobsOutput
				p := svc.NewListRestoreJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-restore-jobs-by-protected-resource": {
			Name:   "list-restore-jobs-by-protected-resource",
			Fields: fields_list_restore_jobs_by_protected_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRestoreJobsByProtectedResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_restore_jobs_by_protected_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRestoreJobsByProtectedResource(ctx, input)
				}
				var results []*svc.ListRestoreJobsByProtectedResourceOutput
				p := svc.NewListRestoreJobsByProtectedResourcePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-restore-testing-plans": {
			Name:   "list-restore-testing-plans",
			Fields: fields_list_restore_testing_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRestoreTestingPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_restore_testing_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRestoreTestingPlans(ctx, input)
				}
				var results []*svc.ListRestoreTestingPlansOutput
				p := svc.NewListRestoreTestingPlansPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-restore-testing-selections": {
			Name:   "list-restore-testing-selections",
			Fields: fields_list_restore_testing_selections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRestoreTestingSelectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_restore_testing_selections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRestoreTestingSelections(ctx, input)
				}
				var results []*svc.ListRestoreTestingSelectionsOutput
				p := svc.NewListRestoreTestingSelectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-scan-job-summaries": {
			Name:   "list-scan-job-summaries",
			Fields: fields_list_scan_job_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScanJobSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scan_job_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScanJobSummaries(ctx, input)
				}
				var results []*svc.ListScanJobSummariesOutput
				p := svc.NewListScanJobSummariesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-scan-jobs": {
			Name:   "list-scan-jobs",
			Fields: fields_list_scan_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScanJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scan_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScanJobs(ctx, input)
				}
				var results []*svc.ListScanJobsOutput
				p := svc.NewListScanJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTags(ctx, input)
				}
				var results []*svc.ListTagsOutput
				p := svc.NewListTagsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tiering-configurations": {
			Name:   "list-tiering-configurations",
			Fields: fields_list_tiering_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTieringConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tiering_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTieringConfigurations(ctx, input)
				}
				var results []*svc.ListTieringConfigurationsOutput
				p := svc.NewListTieringConfigurationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"put-backup-vault-access-policy": {
			Name:   "put-backup-vault-access-policy",
			Fields: fields_put_backup_vault_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBackupVaultAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_backup_vault_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBackupVaultAccessPolicy(ctx, input)
			},
		},
		"put-backup-vault-lock-configuration": {
			Name:   "put-backup-vault-lock-configuration",
			Fields: fields_put_backup_vault_lock_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBackupVaultLockConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_backup_vault_lock_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBackupVaultLockConfiguration(ctx, input)
			},
		},
		"put-backup-vault-notifications": {
			Name:   "put-backup-vault-notifications",
			Fields: fields_put_backup_vault_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBackupVaultNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_backup_vault_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBackupVaultNotifications(ctx, input)
			},
		},
		"put-restore-validation-result": {
			Name:   "put-restore-validation-result",
			Fields: fields_put_restore_validation_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRestoreValidationResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_restore_validation_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRestoreValidationResult(ctx, input)
			},
		},
		"revoke-restore-access-backup-vault": {
			Name:   "revoke-restore-access-backup-vault",
			Fields: fields_revoke_restore_access_backup_vault,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeRestoreAccessBackupVaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_restore_access_backup_vault, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeRestoreAccessBackupVault(ctx, input)
			},
		},
		"start-backup-job": {
			Name:   "start-backup-job",
			Fields: fields_start_backup_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBackupJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_backup_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBackupJob(ctx, input)
			},
		},
		"start-copy-job": {
			Name:   "start-copy-job",
			Fields: fields_start_copy_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCopyJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_copy_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCopyJob(ctx, input)
			},
		},
		"start-report-job": {
			Name:   "start-report-job",
			Fields: fields_start_report_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_report_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReportJob(ctx, input)
			},
		},
		"start-restore-job": {
			Name:   "start-restore-job",
			Fields: fields_start_restore_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRestoreJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_restore_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRestoreJob(ctx, input)
			},
		},
		"start-scan-job": {
			Name:   "start-scan-job",
			Fields: fields_start_scan_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartScanJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_scan_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartScanJob(ctx, input)
			},
		},
		"stop-backup-job": {
			Name:   "stop-backup-job",
			Fields: fields_stop_backup_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBackupJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_backup_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBackupJob(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-backup-plan": {
			Name:   "update-backup-plan",
			Fields: fields_update_backup_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBackupPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_backup_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBackupPlan(ctx, input)
			},
		},
		"update-framework": {
			Name:   "update-framework",
			Fields: fields_update_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFramework(ctx, input)
			},
		},
		"update-global-settings": {
			Name:   "update-global-settings",
			Fields: fields_update_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalSettings(ctx, input)
			},
		},
		"update-recovery-point-index-settings": {
			Name:   "update-recovery-point-index-settings",
			Fields: fields_update_recovery_point_index_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecoveryPointIndexSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recovery_point_index_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecoveryPointIndexSettings(ctx, input)
			},
		},
		"update-recovery-point-lifecycle": {
			Name:   "update-recovery-point-lifecycle",
			Fields: fields_update_recovery_point_lifecycle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecoveryPointLifecycleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recovery_point_lifecycle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecoveryPointLifecycle(ctx, input)
			},
		},
		"update-region-settings": {
			Name:   "update-region-settings",
			Fields: fields_update_region_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRegionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_region_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRegionSettings(ctx, input)
			},
		},
		"update-report-plan": {
			Name:   "update-report-plan",
			Fields: fields_update_report_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReportPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_report_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReportPlan(ctx, input)
			},
		},
		"update-restore-testing-plan": {
			Name:   "update-restore-testing-plan",
			Fields: fields_update_restore_testing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRestoreTestingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_restore_testing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRestoreTestingPlan(ctx, input)
			},
		},
		"update-restore-testing-selection": {
			Name:   "update-restore-testing-selection",
			Fields: fields_update_restore_testing_selection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRestoreTestingSelectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_restore_testing_selection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRestoreTestingSelection(ctx, input)
			},
		},
		"update-tiering-configuration": {
			Name:   "update-tiering-configuration",
			Fields: fields_update_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTieringConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("backup", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
