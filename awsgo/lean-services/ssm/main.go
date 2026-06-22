package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssm"
)

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceTypeForTagging", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_associate_ops_item_related_item = []leanruntime.Field{
	{Name: "AssociationType", Flag: "association-type", Type: "*string", Required: true},
	{Name: "OpsItemId", Flag: "ops-item-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
	{Name: "ResourceUri", Flag: "resource-uri", Type: "*string", Required: true},
}

var fields_cancel_command = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
}

var fields_cancel_maintenance_window_execution = []leanruntime.Field{
	{Name: "WindowExecutionId", Flag: "window-execution-id", Type: "*string", Required: true},
}

var fields_create_activation = []leanruntime.Field{
	{Name: "DefaultInstanceName", Flag: "default-instance-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpirationDate", Flag: "expiration-date", Type: "*time.Time", Required: false},
	{Name: "IamRole", Flag: "iam-role", Type: "*string", Required: true},
	{Name: "RegistrationLimit", Flag: "registration-limit", Type: "*int32", Required: false},
	{Name: "RegistrationMetadata", Flag: "registration-metadata", Type: "[]types.RegistrationMetadataItem", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_association = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "ApplyOnlyAtCronInterval", Flag: "apply-only-at-cron-interval", Type: "bool", Required: false},
	{Name: "AssociationDispatchAssumeRole", Flag: "association-dispatch-assume-role", Type: "*string", Required: false},
	{Name: "AssociationName", Flag: "association-name", Type: "*string", Required: false},
	{Name: "AutomationTargetParameterName", Flag: "automation-target-parameter-name", Type: "*string", Required: false},
	{Name: "CalendarNames", Flag: "calendar-names", Type: "[]string", Required: false},
	{Name: "ComplianceSeverity", Flag: "compliance-severity", Type: "types.AssociationComplianceSeverity", Required: false},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*string", Required: false},
	{Name: "MaxErrors", Flag: "max-errors", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputLocation", Flag: "output-location", Type: "*types.InstanceAssociationOutputLocation", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: false},
	{Name: "ScheduleOffset", Flag: "schedule-offset", Type: "*int32", Required: false},
	{Name: "SyncCompliance", Flag: "sync-compliance", Type: "types.AssociationSyncCompliance", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetLocations", Flag: "target-locations", Type: "[]types.TargetLocation", Required: false},
	{Name: "TargetMaps", Flag: "target-maps", Type: "[]map[string][]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
}

var fields_create_association_batch = []leanruntime.Field{
	{Name: "AssociationDispatchAssumeRole", Flag: "association-dispatch-assume-role", Type: "*string", Required: false},
	{Name: "Entries", Flag: "entries", Type: "[]types.CreateAssociationBatchRequestEntry", Required: true},
}

var fields_create_document = []leanruntime.Field{
	{Name: "Attachments", Flag: "attachments", Type: "[]types.AttachmentsSource", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DocumentFormat", Flag: "document-format", Type: "types.DocumentFormat", Required: false},
	{Name: "DocumentType", Flag: "document-type", Type: "types.DocumentType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Requires", Flag: "requires", Type: "[]types.DocumentRequires", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetType", Flag: "target-type", Type: "*string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_create_maintenance_window = []leanruntime.Field{
	{Name: "AllowUnassociatedTargets", Flag: "allow-unassociated-targets", Type: "bool", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Cutoff", Flag: "cutoff", Type: "int32", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: true},
	{Name: "EndDate", Flag: "end-date", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: true},
	{Name: "ScheduleOffset", Flag: "schedule-offset", Type: "*int32", Required: false},
	{Name: "ScheduleTimezone", Flag: "schedule-timezone", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_ops_item = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ActualEndTime", Flag: "actual-end-time", Type: "*time.Time", Required: false},
	{Name: "ActualStartTime", Flag: "actual-start-time", Type: "*time.Time", Required: false},
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Notifications", Flag: "notifications", Type: "[]types.OpsItemNotification", Required: false},
	{Name: "OperationalData", Flag: "operational-data", Type: "map[string]types.OpsItemDataValue", Required: false},
	{Name: "OpsItemType", Flag: "ops-item-type", Type: "*string", Required: false},
	{Name: "PlannedEndTime", Flag: "planned-end-time", Type: "*time.Time", Required: false},
	{Name: "PlannedStartTime", Flag: "planned-start-time", Type: "*time.Time", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "RelatedOpsItems", Flag: "related-ops-items", Type: "[]types.RelatedOpsItem", Required: false},
	{Name: "Severity", Flag: "severity", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_create_ops_metadata = []leanruntime.Field{
	{Name: "Metadata", Flag: "metadata", Type: "map[string]types.MetadataValue", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_patch_baseline = []leanruntime.Field{
	{Name: "ApprovalRules", Flag: "approval-rules", Type: "*types.PatchRuleGroup", Required: false},
	{Name: "ApprovedPatches", Flag: "approved-patches", Type: "[]string", Required: false},
	{Name: "ApprovedPatchesComplianceLevel", Flag: "approved-patches-compliance-level", Type: "types.PatchComplianceLevel", Required: false},
	{Name: "ApprovedPatchesEnableNonSecurity", Flag: "approved-patches-enable-non-security", Type: "*bool", Required: false},
	{Name: "AvailableSecurityUpdatesComplianceStatus", Flag: "available-security-updates-compliance-status", Type: "types.PatchComplianceStatus", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalFilters", Flag: "global-filters", Type: "*types.PatchFilterGroup", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.OperatingSystem", Required: false},
	{Name: "RejectedPatches", Flag: "rejected-patches", Type: "[]string", Required: false},
	{Name: "RejectedPatchesAction", Flag: "rejected-patches-action", Type: "types.PatchAction", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.PatchSource", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_resource_data_sync = []leanruntime.Field{
	{Name: "S3Destination", Flag: "s3-destination", Type: "*types.ResourceDataSyncS3Destination", Required: false},
	{Name: "SyncName", Flag: "sync-name", Type: "*string", Required: true},
	{Name: "SyncSource", Flag: "sync-source", Type: "*types.ResourceDataSyncSource", Required: false},
	{Name: "SyncType", Flag: "sync-type", Type: "*string", Required: false},
}

var fields_delete_activation = []leanruntime.Field{
	{Name: "ActivationId", Flag: "activation-id", Type: "*string", Required: true},
}

var fields_delete_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_delete_document = []leanruntime.Field{
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_delete_inventory = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "SchemaDeleteOption", Flag: "schema-delete-option", Type: "types.InventorySchemaDeleteOption", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_delete_maintenance_window = []leanruntime.Field{
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_delete_ops_item = []leanruntime.Field{
	{Name: "OpsItemId", Flag: "ops-item-id", Type: "*string", Required: true},
}

var fields_delete_ops_metadata = []leanruntime.Field{
	{Name: "OpsMetadataArn", Flag: "ops-metadata-arn", Type: "*string", Required: true},
}

var fields_delete_parameter = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_parameters = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_delete_patch_baseline = []leanruntime.Field{
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
}

var fields_delete_resource_data_sync = []leanruntime.Field{
	{Name: "SyncName", Flag: "sync-name", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "*string", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "PolicyHash", Flag: "policy-hash", Type: "*string", Required: true},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_deregister_managed_instance = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_deregister_patch_baseline_for_patch_group = []leanruntime.Field{
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
	{Name: "PatchGroup", Flag: "patch-group", Type: "*string", Required: true},
}

var fields_deregister_target_from_maintenance_window = []leanruntime.Field{
	{Name: "Safe", Flag: "safe", Type: "*bool", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
	{Name: "WindowTargetId", Flag: "window-target-id", Type: "*string", Required: true},
}

var fields_deregister_task_from_maintenance_window = []leanruntime.Field{
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
	{Name: "WindowTaskId", Flag: "window-task-id", Type: "*string", Required: true},
}

var fields_describe_activations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DescribeActivationsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: false},
	{Name: "AssociationVersion", Flag: "association-version", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_describe_association_execution_targets = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AssociationExecutionTargetsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_association_executions = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AssociationExecutionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_automation_executions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.AutomationExecutionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_automation_step_executions = []leanruntime.Field{
	{Name: "AutomationExecutionId", Flag: "automation-execution-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.StepExecutionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "*bool", Required: false},
}

var fields_describe_available_patches = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PatchOrchestratorFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_document = []leanruntime.Field{
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_describe_document_permission = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionType", Flag: "permission-type", Type: "types.DocumentPermissionType", Required: true},
}

var fields_describe_effective_instance_associations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_effective_patches_for_patch_baseline = []leanruntime.Field{
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_associations_status = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_information = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.InstanceInformationStringFilter", Required: false},
	{Name: "InstanceInformationFilterList", Flag: "instance-information-filter-list", Type: "[]types.InstanceInformationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_patch_states = []leanruntime.Field{
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_patch_states_for_patch_group = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.InstancePatchStateFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PatchGroup", Flag: "patch-group", Type: "*string", Required: true},
}

var fields_describe_instance_patches = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PatchOrchestratorFilter", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_properties = []leanruntime.Field{
	{Name: "FiltersWithOperator", Flag: "filters-with-operator", Type: "[]types.InstancePropertyStringFilter", Required: false},
	{Name: "InstancePropertyFilterList", Flag: "instance-property-filter-list", Type: "[]types.InstancePropertyFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_inventory_deletions = []leanruntime.Field{
	{Name: "DeletionId", Flag: "deletion-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_maintenance_window_execution_task_invocations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.MaintenanceWindowFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
	{Name: "WindowExecutionId", Flag: "window-execution-id", Type: "*string", Required: true},
}

var fields_describe_maintenance_window_execution_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.MaintenanceWindowFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WindowExecutionId", Flag: "window-execution-id", Type: "*string", Required: true},
}

var fields_describe_maintenance_window_executions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.MaintenanceWindowFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_describe_maintenance_window_schedule = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PatchOrchestratorFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.MaintenanceWindowResourceType", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: false},
}

var fields_describe_maintenance_window_targets = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.MaintenanceWindowFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_describe_maintenance_window_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.MaintenanceWindowFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_describe_maintenance_windows = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.MaintenanceWindowFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_maintenance_windows_for_target = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.MaintenanceWindowResourceType", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
}

var fields_describe_ops_items = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OpsItemFilters", Flag: "ops-item-filters", Type: "[]types.OpsItemFilter", Required: false},
}

var fields_describe_parameters = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ParametersFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterFilters", Flag: "parameter-filters", Type: "[]types.ParameterStringFilter", Required: false},
	{Name: "Shared", Flag: "shared", Type: "*bool", Required: false},
}

var fields_describe_patch_baselines = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PatchOrchestratorFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_patch_group_state = []leanruntime.Field{
	{Name: "PatchGroup", Flag: "patch-group", Type: "*string", Required: true},
}

var fields_describe_patch_groups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PatchOrchestratorFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_patch_properties = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.OperatingSystem", Required: true},
	{Name: "PatchSet", Flag: "patch-set", Type: "types.PatchSet", Required: false},
	{Name: "Property", Flag: "property", Type: "types.PatchProperty", Required: true},
}

var fields_describe_sessions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SessionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.SessionState", Required: true},
}

var fields_disassociate_ops_item_related_item = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "OpsItemId", Flag: "ops-item-id", Type: "*string", Required: true},
}

var fields_get_access_token = []leanruntime.Field{
	{Name: "AccessRequestId", Flag: "access-request-id", Type: "*string", Required: true},
}

var fields_get_automation_execution = []leanruntime.Field{
	{Name: "AutomationExecutionId", Flag: "automation-execution-id", Type: "*string", Required: true},
}

var fields_get_calendar_state = []leanruntime.Field{
	{Name: "AtTime", Flag: "at-time", Type: "*string", Required: false},
	{Name: "CalendarNames", Flag: "calendar-names", Type: "[]string", Required: true},
}

var fields_get_command_invocation = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PluginName", Flag: "plugin-name", Type: "*string", Required: false},
}

var fields_get_connection_status = []leanruntime.Field{
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
}

var fields_get_default_patch_baseline = []leanruntime.Field{
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.OperatingSystem", Required: false},
}

var fields_get_deployable_patch_snapshot_for_instance = []leanruntime.Field{
	{Name: "BaselineOverride", Flag: "baseline-override", Type: "*types.BaselineOverride", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
	{Name: "UseS3DualStackEndpoint", Flag: "use-s3-dual-stack-endpoint", Type: "bool", Required: false},
}

var fields_get_document = []leanruntime.Field{
	{Name: "DocumentFormat", Flag: "document-format", Type: "types.DocumentFormat", Required: false},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_get_execution_preview = []leanruntime.Field{
	{Name: "ExecutionPreviewId", Flag: "execution-preview-id", Type: "*string", Required: true},
}

var fields_get_inventory = []leanruntime.Field{
	{Name: "Aggregators", Flag: "aggregators", Type: "[]types.InventoryAggregator", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.InventoryFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResultAttributes", Flag: "result-attributes", Type: "[]types.ResultAttribute", Required: false},
}

var fields_get_inventory_schema = []leanruntime.Field{
	{Name: "Aggregator", Flag: "aggregator", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubType", Flag: "sub-type", Type: "*bool", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
}

var fields_get_maintenance_window = []leanruntime.Field{
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_get_maintenance_window_execution = []leanruntime.Field{
	{Name: "WindowExecutionId", Flag: "window-execution-id", Type: "*string", Required: true},
}

var fields_get_maintenance_window_execution_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
	{Name: "WindowExecutionId", Flag: "window-execution-id", Type: "*string", Required: true},
}

var fields_get_maintenance_window_execution_task_invocation = []leanruntime.Field{
	{Name: "InvocationId", Flag: "invocation-id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
	{Name: "WindowExecutionId", Flag: "window-execution-id", Type: "*string", Required: true},
}

var fields_get_maintenance_window_task = []leanruntime.Field{
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
	{Name: "WindowTaskId", Flag: "window-task-id", Type: "*string", Required: true},
}

var fields_get_ops_item = []leanruntime.Field{
	{Name: "OpsItemArn", Flag: "ops-item-arn", Type: "*string", Required: false},
	{Name: "OpsItemId", Flag: "ops-item-id", Type: "*string", Required: true},
}

var fields_get_ops_metadata = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OpsMetadataArn", Flag: "ops-metadata-arn", Type: "*string", Required: true},
}

var fields_get_ops_summary = []leanruntime.Field{
	{Name: "Aggregators", Flag: "aggregators", Type: "[]types.OpsAggregator", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.OpsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResultAttributes", Flag: "result-attributes", Type: "[]types.OpsResultAttribute", Required: false},
	{Name: "SyncName", Flag: "sync-name", Type: "*string", Required: false},
}

var fields_get_parameter = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WithDecryption", Flag: "with-decryption", Type: "*bool", Required: false},
}

var fields_get_parameter_history = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WithDecryption", Flag: "with-decryption", Type: "*bool", Required: false},
}

var fields_get_parameters = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
	{Name: "WithDecryption", Flag: "with-decryption", Type: "*bool", Required: false},
}

var fields_get_parameters_by_path = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParameterFilters", Flag: "parameter-filters", Type: "[]types.ParameterStringFilter", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: true},
	{Name: "Recursive", Flag: "recursive", Type: "*bool", Required: false},
	{Name: "WithDecryption", Flag: "with-decryption", Type: "*bool", Required: false},
}

var fields_get_patch_baseline = []leanruntime.Field{
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
}

var fields_get_patch_baseline_for_patch_group = []leanruntime.Field{
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.OperatingSystem", Required: false},
	{Name: "PatchGroup", Flag: "patch-group", Type: "*string", Required: true},
}

var fields_get_resource_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_service_setting = []leanruntime.Field{
	{Name: "SettingId", Flag: "setting-id", Type: "*string", Required: true},
}

var fields_label_parameter_version = []leanruntime.Field{
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParameterVersion", Flag: "parameter-version", Type: "*int64", Required: false},
}

var fields_list_association_versions = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associations = []leanruntime.Field{
	{Name: "AssociationFilterList", Flag: "association-filter-list", Type: "[]types.AssociationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_command_invocations = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: false},
	{Name: "Details", Flag: "details", Type: "bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.CommandFilter", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_commands = []leanruntime.Field{
	{Name: "CommandId", Flag: "command-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.CommandFilter", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_compliance_items = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ComplianceStringFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
}

var fields_list_compliance_summaries = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ComplianceStringFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_document_metadata_history = []leanruntime.Field{
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "types.DocumentMetadataEnum", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_document_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_documents = []leanruntime.Field{
	{Name: "DocumentFilterList", Flag: "document-filter-list", Type: "[]types.DocumentFilter", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.DocumentKeyValuesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_inventory_entries = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.InventoryFilter", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_list_nodes = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.NodeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SyncName", Flag: "sync-name", Type: "*string", Required: false},
}

var fields_list_nodes_summary = []leanruntime.Field{
	{Name: "Aggregators", Flag: "aggregators", Type: "[]types.NodeAggregator", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.NodeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SyncName", Flag: "sync-name", Type: "*string", Required: false},
}

var fields_list_ops_item_events = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.OpsItemEventFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ops_item_related_items = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.OpsItemRelatedItemsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OpsItemId", Flag: "ops-item-id", Type: "*string", Required: false},
}

var fields_list_ops_metadata = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.OpsMetadataFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_compliance_summaries = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ComplianceStringFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_data_sync = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SyncType", Flag: "sync-type", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceTypeForTagging", Required: true},
}

var fields_modify_document_permission = []leanruntime.Field{
	{Name: "AccountIdsToAdd", Flag: "account-ids-to-add", Type: "[]string", Required: false},
	{Name: "AccountIdsToRemove", Flag: "account-ids-to-remove", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PermissionType", Flag: "permission-type", Type: "types.DocumentPermissionType", Required: true},
	{Name: "SharedDocumentVersion", Flag: "shared-document-version", Type: "*string", Required: false},
}

var fields_put_compliance_items = []leanruntime.Field{
	{Name: "ComplianceType", Flag: "compliance-type", Type: "*string", Required: true},
	{Name: "ExecutionSummary", Flag: "execution-summary", Type: "*types.ComplianceExecutionSummary", Required: true},
	{Name: "ItemContentHash", Flag: "item-content-hash", Type: "*string", Required: false},
	{Name: "Items", Flag: "items", Type: "[]types.ComplianceItemEntry", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
	{Name: "UploadType", Flag: "upload-type", Type: "types.ComplianceUploadType", Required: false},
}

var fields_put_inventory = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Items", Flag: "items", Type: "[]types.InventoryItem", Required: true},
}

var fields_put_parameter = []leanruntime.Field{
	{Name: "AllowedPattern", Flag: "allowed-pattern", Type: "*string", Required: false},
	{Name: "DataType", Flag: "data-type", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Overwrite", Flag: "overwrite", Type: "*bool", Required: false},
	{Name: "Policies", Flag: "policies", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.ParameterTier", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ParameterType", Required: false},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "PolicyHash", Flag: "policy-hash", Type: "*string", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_default_patch_baseline = []leanruntime.Field{
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
}

var fields_register_patch_baseline_for_patch_group = []leanruntime.Field{
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
	{Name: "PatchGroup", Flag: "patch-group", Type: "*string", Required: true},
}

var fields_register_target_with_maintenance_window = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OwnerInformation", Flag: "owner-information", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.MaintenanceWindowResourceType", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_register_task_with_maintenance_window = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CutoffBehavior", Flag: "cutoff-behavior", Type: "types.MaintenanceWindowTaskCutoffBehavior", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LoggingInfo", Flag: "logging-info", Type: "*types.LoggingInfo", Required: false},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*string", Required: false},
	{Name: "MaxErrors", Flag: "max-errors", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
	{Name: "TaskInvocationParameters", Flag: "task-invocation-parameters", Type: "*types.MaintenanceWindowTaskInvocationParameters", Required: false},
	{Name: "TaskParameters", Flag: "task-parameters", Type: "map[string]types.MaintenanceWindowTaskParameterValueExpression", Required: false},
	{Name: "TaskType", Flag: "task-type", Type: "types.MaintenanceWindowTaskType", Required: true},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceTypeForTagging", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_reset_service_setting = []leanruntime.Field{
	{Name: "SettingId", Flag: "setting-id", Type: "*string", Required: true},
}

var fields_resume_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_send_automation_signal = []leanruntime.Field{
	{Name: "AutomationExecutionId", Flag: "automation-execution-id", Type: "*string", Required: true},
	{Name: "Payload", Flag: "payload", Type: "map[string][]string", Required: false},
	{Name: "SignalType", Flag: "signal-type", Type: "types.SignalType", Required: true},
}

var fields_send_command = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "CloudWatchOutputConfig", Flag: "cloud-watch-output-config", Type: "*types.CloudWatchOutputConfig", Required: false},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "DocumentHash", Flag: "document-hash", Type: "*string", Required: false},
	{Name: "DocumentHashType", Flag: "document-hash-type", Type: "types.DocumentHashType", Required: false},
	{Name: "DocumentName", Flag: "document-name", Type: "*string", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*string", Required: false},
	{Name: "MaxErrors", Flag: "max-errors", Type: "*string", Required: false},
	{Name: "NotificationConfig", Flag: "notification-config", Type: "*types.NotificationConfig", Required: false},
	{Name: "OutputS3BucketName", Flag: "output-s3-bucket-name", Type: "*string", Required: false},
	{Name: "OutputS3KeyPrefix", Flag: "output-s3-key-prefix", Type: "*string", Required: false},
	{Name: "OutputS3Region", Flag: "output-s3-region", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
	{Name: "TimeoutSeconds", Flag: "timeout-seconds", Type: "*int32", Required: false},
}

var fields_start_access_request = []leanruntime.Field{
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
}

var fields_start_associations_once = []leanruntime.Field{
	{Name: "AssociationIds", Flag: "association-ids", Type: "[]string", Required: true},
}

var fields_start_automation_execution = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DocumentName", Flag: "document-name", Type: "*string", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*string", Required: false},
	{Name: "MaxErrors", Flag: "max-errors", Type: "*string", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.ExecutionMode", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetLocations", Flag: "target-locations", Type: "[]types.TargetLocation", Required: false},
	{Name: "TargetLocationsURL", Flag: "target-locations-url", Type: "*string", Required: false},
	{Name: "TargetMaps", Flag: "target-maps", Type: "[]map[string][]string", Required: false},
	{Name: "TargetParameterName", Flag: "target-parameter-name", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
}

var fields_start_change_request_execution = []leanruntime.Field{
	{Name: "AutoApprove", Flag: "auto-approve", Type: "bool", Required: false},
	{Name: "ChangeDetails", Flag: "change-details", Type: "*string", Required: false},
	{Name: "ChangeRequestName", Flag: "change-request-name", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DocumentName", Flag: "document-name", Type: "*string", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "Runbooks", Flag: "runbooks", Type: "[]types.Runbook", Required: true},
	{Name: "ScheduledEndTime", Flag: "scheduled-end-time", Type: "*time.Time", Required: false},
	{Name: "ScheduledTime", Flag: "scheduled-time", Type: "*time.Time", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_execution_preview = []leanruntime.Field{
	{Name: "DocumentName", Flag: "document-name", Type: "*string", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "ExecutionInputs", Flag: "execution-inputs", Type: "types.ExecutionInputs", Required: false},
}

var fields_start_session = []leanruntime.Field{
	{Name: "DocumentName", Flag: "document-name", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
}

var fields_stop_automation_execution = []leanruntime.Field{
	{Name: "AutomationExecutionId", Flag: "automation-execution-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.StopType", Required: false},
}

var fields_terminate_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_unlabel_parameter_version = []leanruntime.Field{
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParameterVersion", Flag: "parameter-version", Type: "*int64", Required: true},
}

var fields_update_association = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "ApplyOnlyAtCronInterval", Flag: "apply-only-at-cron-interval", Type: "bool", Required: false},
	{Name: "AssociationDispatchAssumeRole", Flag: "association-dispatch-assume-role", Type: "*string", Required: false},
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "AssociationName", Flag: "association-name", Type: "*string", Required: false},
	{Name: "AssociationVersion", Flag: "association-version", Type: "*string", Required: false},
	{Name: "AutomationTargetParameterName", Flag: "automation-target-parameter-name", Type: "*string", Required: false},
	{Name: "CalendarNames", Flag: "calendar-names", Type: "[]string", Required: false},
	{Name: "ComplianceSeverity", Flag: "compliance-severity", Type: "types.AssociationComplianceSeverity", Required: false},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*string", Required: false},
	{Name: "MaxErrors", Flag: "max-errors", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OutputLocation", Flag: "output-location", Type: "*types.InstanceAssociationOutputLocation", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: false},
	{Name: "ScheduleOffset", Flag: "schedule-offset", Type: "*int32", Required: false},
	{Name: "SyncCompliance", Flag: "sync-compliance", Type: "types.AssociationSyncCompliance", Required: false},
	{Name: "TargetLocations", Flag: "target-locations", Type: "[]types.TargetLocation", Required: false},
	{Name: "TargetMaps", Flag: "target-maps", Type: "[]map[string][]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
}

var fields_update_association_status = []leanruntime.Field{
	{Name: "AssociationStatus", Flag: "association-status", Type: "*types.AssociationStatus", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_document = []leanruntime.Field{
	{Name: "Attachments", Flag: "attachments", Type: "[]types.AttachmentsSource", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DocumentFormat", Flag: "document-format", Type: "types.DocumentFormat", Required: false},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "*string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_update_document_default_version = []leanruntime.Field{
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_document_metadata = []leanruntime.Field{
	{Name: "DocumentReviews", Flag: "document-reviews", Type: "*types.DocumentReviews", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_maintenance_window = []leanruntime.Field{
	{Name: "AllowUnassociatedTargets", Flag: "allow-unassociated-targets", Type: "*bool", Required: false},
	{Name: "Cutoff", Flag: "cutoff", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Replace", Flag: "replace", Type: "*bool", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "ScheduleOffset", Flag: "schedule-offset", Type: "*int32", Required: false},
	{Name: "ScheduleTimezone", Flag: "schedule-timezone", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*string", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
}

var fields_update_maintenance_window_target = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OwnerInformation", Flag: "owner-information", Type: "*string", Required: false},
	{Name: "Replace", Flag: "replace", Type: "*bool", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
	{Name: "WindowTargetId", Flag: "window-target-id", Type: "*string", Required: true},
}

var fields_update_maintenance_window_task = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "CutoffBehavior", Flag: "cutoff-behavior", Type: "types.MaintenanceWindowTaskCutoffBehavior", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LoggingInfo", Flag: "logging-info", Type: "*types.LoggingInfo", Required: false},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*string", Required: false},
	{Name: "MaxErrors", Flag: "max-errors", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "Replace", Flag: "replace", Type: "*bool", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: false},
	{Name: "TaskInvocationParameters", Flag: "task-invocation-parameters", Type: "*types.MaintenanceWindowTaskInvocationParameters", Required: false},
	{Name: "TaskParameters", Flag: "task-parameters", Type: "map[string]types.MaintenanceWindowTaskParameterValueExpression", Required: false},
	{Name: "WindowId", Flag: "window-id", Type: "*string", Required: true},
	{Name: "WindowTaskId", Flag: "window-task-id", Type: "*string", Required: true},
}

var fields_update_managed_instance_role = []leanruntime.Field{
	{Name: "IamRole", Flag: "iam-role", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_update_ops_item = []leanruntime.Field{
	{Name: "ActualEndTime", Flag: "actual-end-time", Type: "*time.Time", Required: false},
	{Name: "ActualStartTime", Flag: "actual-start-time", Type: "*time.Time", Required: false},
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Notifications", Flag: "notifications", Type: "[]types.OpsItemNotification", Required: false},
	{Name: "OperationalData", Flag: "operational-data", Type: "map[string]types.OpsItemDataValue", Required: false},
	{Name: "OperationalDataToDelete", Flag: "operational-data-to-delete", Type: "[]string", Required: false},
	{Name: "OpsItemArn", Flag: "ops-item-arn", Type: "*string", Required: false},
	{Name: "OpsItemId", Flag: "ops-item-id", Type: "*string", Required: true},
	{Name: "PlannedEndTime", Flag: "planned-end-time", Type: "*time.Time", Required: false},
	{Name: "PlannedStartTime", Flag: "planned-start-time", Type: "*time.Time", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "RelatedOpsItems", Flag: "related-ops-items", Type: "[]types.RelatedOpsItem", Required: false},
	{Name: "Severity", Flag: "severity", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.OpsItemStatus", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
}

var fields_update_ops_metadata = []leanruntime.Field{
	{Name: "KeysToDelete", Flag: "keys-to-delete", Type: "[]string", Required: false},
	{Name: "MetadataToUpdate", Flag: "metadata-to-update", Type: "map[string]types.MetadataValue", Required: false},
	{Name: "OpsMetadataArn", Flag: "ops-metadata-arn", Type: "*string", Required: true},
}

var fields_update_patch_baseline = []leanruntime.Field{
	{Name: "ApprovalRules", Flag: "approval-rules", Type: "*types.PatchRuleGroup", Required: false},
	{Name: "ApprovedPatches", Flag: "approved-patches", Type: "[]string", Required: false},
	{Name: "ApprovedPatchesComplianceLevel", Flag: "approved-patches-compliance-level", Type: "types.PatchComplianceLevel", Required: false},
	{Name: "ApprovedPatchesEnableNonSecurity", Flag: "approved-patches-enable-non-security", Type: "*bool", Required: false},
	{Name: "AvailableSecurityUpdatesComplianceStatus", Flag: "available-security-updates-compliance-status", Type: "types.PatchComplianceStatus", Required: false},
	{Name: "BaselineId", Flag: "baseline-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalFilters", Flag: "global-filters", Type: "*types.PatchFilterGroup", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RejectedPatches", Flag: "rejected-patches", Type: "[]string", Required: false},
	{Name: "RejectedPatchesAction", Flag: "rejected-patches-action", Type: "types.PatchAction", Required: false},
	{Name: "Replace", Flag: "replace", Type: "*bool", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.PatchSource", Required: false},
}

var fields_update_resource_data_sync = []leanruntime.Field{
	{Name: "SyncName", Flag: "sync-name", Type: "*string", Required: true},
	{Name: "SyncSource", Flag: "sync-source", Type: "*types.ResourceDataSyncSource", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "*string", Required: true},
}

var fields_update_service_setting = []leanruntime.Field{
	{Name: "SettingId", Flag: "setting-id", Type: "*string", Required: true},
	{Name: "SettingValue", Flag: "setting-value", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags-to-resource": {
			Name:   "add-tags-to-resource",
			Fields: fields_add_tags_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToResource(ctx, input)
			},
		},
		"associate-ops-item-related-item": {
			Name:   "associate-ops-item-related-item",
			Fields: fields_associate_ops_item_related_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateOpsItemRelatedItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_ops_item_related_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateOpsItemRelatedItem(ctx, input)
			},
		},
		"cancel-command": {
			Name:   "cancel-command",
			Fields: fields_cancel_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCommand(ctx, input)
			},
		},
		"cancel-maintenance-window-execution": {
			Name:   "cancel-maintenance-window-execution",
			Fields: fields_cancel_maintenance_window_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMaintenanceWindowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_maintenance_window_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMaintenanceWindowExecution(ctx, input)
			},
		},
		"create-activation": {
			Name:   "create-activation",
			Fields: fields_create_activation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateActivationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_activation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateActivation(ctx, input)
			},
		},
		"create-association": {
			Name:   "create-association",
			Fields: fields_create_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssociation(ctx, input)
			},
		},
		"create-association-batch": {
			Name:   "create-association-batch",
			Fields: fields_create_association_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssociationBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_association_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssociationBatch(ctx, input)
			},
		},
		"create-document": {
			Name:   "create-document",
			Fields: fields_create_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDocument(ctx, input)
			},
		},
		"create-maintenance-window": {
			Name:   "create-maintenance-window",
			Fields: fields_create_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMaintenanceWindow(ctx, input)
			},
		},
		"create-ops-item": {
			Name:   "create-ops-item",
			Fields: fields_create_ops_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOpsItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ops_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOpsItem(ctx, input)
			},
		},
		"create-ops-metadata": {
			Name:   "create-ops-metadata",
			Fields: fields_create_ops_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOpsMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ops_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOpsMetadata(ctx, input)
			},
		},
		"create-patch-baseline": {
			Name:   "create-patch-baseline",
			Fields: fields_create_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePatchBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_patch_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePatchBaseline(ctx, input)
			},
		},
		"create-resource-data-sync": {
			Name:   "create-resource-data-sync",
			Fields: fields_create_resource_data_sync,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceDataSyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_data_sync, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceDataSync(ctx, input)
			},
		},
		"delete-activation": {
			Name:   "delete-activation",
			Fields: fields_delete_activation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteActivationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_activation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteActivation(ctx, input)
			},
		},
		"delete-association": {
			Name:   "delete-association",
			Fields: fields_delete_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssociation(ctx, input)
			},
		},
		"delete-document": {
			Name:   "delete-document",
			Fields: fields_delete_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDocument(ctx, input)
			},
		},
		"delete-inventory": {
			Name:   "delete-inventory",
			Fields: fields_delete_inventory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInventoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inventory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInventory(ctx, input)
			},
		},
		"delete-maintenance-window": {
			Name:   "delete-maintenance-window",
			Fields: fields_delete_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMaintenanceWindow(ctx, input)
			},
		},
		"delete-ops-item": {
			Name:   "delete-ops-item",
			Fields: fields_delete_ops_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOpsItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ops_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOpsItem(ctx, input)
			},
		},
		"delete-ops-metadata": {
			Name:   "delete-ops-metadata",
			Fields: fields_delete_ops_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOpsMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ops_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOpsMetadata(ctx, input)
			},
		},
		"delete-parameter": {
			Name:   "delete-parameter",
			Fields: fields_delete_parameter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteParameterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_parameter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteParameter(ctx, input)
			},
		},
		"delete-parameters": {
			Name:   "delete-parameters",
			Fields: fields_delete_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteParameters(ctx, input)
			},
		},
		"delete-patch-baseline": {
			Name:   "delete-patch-baseline",
			Fields: fields_delete_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePatchBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_patch_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePatchBaseline(ctx, input)
			},
		},
		"delete-resource-data-sync": {
			Name:   "delete-resource-data-sync",
			Fields: fields_delete_resource_data_sync,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceDataSyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_data_sync, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceDataSync(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"deregister-managed-instance": {
			Name:   "deregister-managed-instance",
			Fields: fields_deregister_managed_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterManagedInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_managed_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterManagedInstance(ctx, input)
			},
		},
		"deregister-patch-baseline-for-patch-group": {
			Name:   "deregister-patch-baseline-for-patch-group",
			Fields: fields_deregister_patch_baseline_for_patch_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterPatchBaselineForPatchGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_patch_baseline_for_patch_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterPatchBaselineForPatchGroup(ctx, input)
			},
		},
		"deregister-target-from-maintenance-window": {
			Name:   "deregister-target-from-maintenance-window",
			Fields: fields_deregister_target_from_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTargetFromMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_target_from_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTargetFromMaintenanceWindow(ctx, input)
			},
		},
		"deregister-task-from-maintenance-window": {
			Name:   "deregister-task-from-maintenance-window",
			Fields: fields_deregister_task_from_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTaskFromMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_task_from_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTaskFromMaintenanceWindow(ctx, input)
			},
		},
		"describe-activations": {
			Name:   "describe-activations",
			Fields: fields_describe_activations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActivationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_activations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeActivations(ctx, input)
				}
				var results []*svc.DescribeActivationsOutput
				p := svc.NewDescribeActivationsPaginator(client, input)
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
		"describe-association": {
			Name:   "describe-association",
			Fields: fields_describe_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssociation(ctx, input)
			},
		},
		"describe-association-execution-targets": {
			Name:   "describe-association-execution-targets",
			Fields: fields_describe_association_execution_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssociationExecutionTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_association_execution_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAssociationExecutionTargets(ctx, input)
				}
				var results []*svc.DescribeAssociationExecutionTargetsOutput
				p := svc.NewDescribeAssociationExecutionTargetsPaginator(client, input)
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
		"describe-association-executions": {
			Name:   "describe-association-executions",
			Fields: fields_describe_association_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssociationExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_association_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAssociationExecutions(ctx, input)
				}
				var results []*svc.DescribeAssociationExecutionsOutput
				p := svc.NewDescribeAssociationExecutionsPaginator(client, input)
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
		"describe-automation-executions": {
			Name:   "describe-automation-executions",
			Fields: fields_describe_automation_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutomationExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_automation_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAutomationExecutions(ctx, input)
				}
				var results []*svc.DescribeAutomationExecutionsOutput
				p := svc.NewDescribeAutomationExecutionsPaginator(client, input)
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
		"describe-automation-step-executions": {
			Name:   "describe-automation-step-executions",
			Fields: fields_describe_automation_step_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutomationStepExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_automation_step_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAutomationStepExecutions(ctx, input)
				}
				var results []*svc.DescribeAutomationStepExecutionsOutput
				p := svc.NewDescribeAutomationStepExecutionsPaginator(client, input)
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
		"describe-available-patches": {
			Name:   "describe-available-patches",
			Fields: fields_describe_available_patches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAvailablePatchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_available_patches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAvailablePatches(ctx, input)
				}
				var results []*svc.DescribeAvailablePatchesOutput
				p := svc.NewDescribeAvailablePatchesPaginator(client, input)
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
		"describe-document": {
			Name:   "describe-document",
			Fields: fields_describe_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDocument(ctx, input)
			},
		},
		"describe-document-permission": {
			Name:   "describe-document-permission",
			Fields: fields_describe_document_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDocumentPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_document_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDocumentPermission(ctx, input)
			},
		},
		"describe-effective-instance-associations": {
			Name:   "describe-effective-instance-associations",
			Fields: fields_describe_effective_instance_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEffectiveInstanceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_effective_instance_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEffectiveInstanceAssociations(ctx, input)
				}
				var results []*svc.DescribeEffectiveInstanceAssociationsOutput
				p := svc.NewDescribeEffectiveInstanceAssociationsPaginator(client, input)
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
		"describe-effective-patches-for-patch-baseline": {
			Name:   "describe-effective-patches-for-patch-baseline",
			Fields: fields_describe_effective_patches_for_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEffectivePatchesForPatchBaselineInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_effective_patches_for_patch_baseline, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEffectivePatchesForPatchBaseline(ctx, input)
				}
				var results []*svc.DescribeEffectivePatchesForPatchBaselineOutput
				p := svc.NewDescribeEffectivePatchesForPatchBaselinePaginator(client, input)
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
		"describe-instance-associations-status": {
			Name:   "describe-instance-associations-status",
			Fields: fields_describe_instance_associations_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceAssociationsStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_associations_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceAssociationsStatus(ctx, input)
				}
				var results []*svc.DescribeInstanceAssociationsStatusOutput
				p := svc.NewDescribeInstanceAssociationsStatusPaginator(client, input)
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
		"describe-instance-information": {
			Name:   "describe-instance-information",
			Fields: fields_describe_instance_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceInformationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_information, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceInformation(ctx, input)
				}
				var results []*svc.DescribeInstanceInformationOutput
				p := svc.NewDescribeInstanceInformationPaginator(client, input)
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
		"describe-instance-patch-states": {
			Name:   "describe-instance-patch-states",
			Fields: fields_describe_instance_patch_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancePatchStatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_patch_states, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstancePatchStates(ctx, input)
				}
				var results []*svc.DescribeInstancePatchStatesOutput
				p := svc.NewDescribeInstancePatchStatesPaginator(client, input)
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
		"describe-instance-patch-states-for-patch-group": {
			Name:   "describe-instance-patch-states-for-patch-group",
			Fields: fields_describe_instance_patch_states_for_patch_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancePatchStatesForPatchGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_patch_states_for_patch_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstancePatchStatesForPatchGroup(ctx, input)
				}
				var results []*svc.DescribeInstancePatchStatesForPatchGroupOutput
				p := svc.NewDescribeInstancePatchStatesForPatchGroupPaginator(client, input)
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
		"describe-instance-patches": {
			Name:   "describe-instance-patches",
			Fields: fields_describe_instance_patches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancePatchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_patches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstancePatches(ctx, input)
				}
				var results []*svc.DescribeInstancePatchesOutput
				p := svc.NewDescribeInstancePatchesPaginator(client, input)
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
		"describe-instance-properties": {
			Name:   "describe-instance-properties",
			Fields: fields_describe_instance_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancePropertiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_properties, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceProperties(ctx, input)
				}
				var results []*svc.DescribeInstancePropertiesOutput
				p := svc.NewDescribeInstancePropertiesPaginator(client, input)
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
		"describe-inventory-deletions": {
			Name:   "describe-inventory-deletions",
			Fields: fields_describe_inventory_deletions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInventoryDeletionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_inventory_deletions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInventoryDeletions(ctx, input)
				}
				var results []*svc.DescribeInventoryDeletionsOutput
				p := svc.NewDescribeInventoryDeletionsPaginator(client, input)
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
		"describe-maintenance-window-execution-task-invocations": {
			Name:   "describe-maintenance-window-execution-task-invocations",
			Fields: fields_describe_maintenance_window_execution_task_invocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowExecutionTaskInvocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_window_execution_task_invocations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowExecutionTaskInvocations(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowExecutionTaskInvocationsOutput
				p := svc.NewDescribeMaintenanceWindowExecutionTaskInvocationsPaginator(client, input)
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
		"describe-maintenance-window-execution-tasks": {
			Name:   "describe-maintenance-window-execution-tasks",
			Fields: fields_describe_maintenance_window_execution_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowExecutionTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_window_execution_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowExecutionTasks(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowExecutionTasksOutput
				p := svc.NewDescribeMaintenanceWindowExecutionTasksPaginator(client, input)
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
		"describe-maintenance-window-executions": {
			Name:   "describe-maintenance-window-executions",
			Fields: fields_describe_maintenance_window_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_window_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowExecutions(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowExecutionsOutput
				p := svc.NewDescribeMaintenanceWindowExecutionsPaginator(client, input)
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
		"describe-maintenance-window-schedule": {
			Name:   "describe-maintenance-window-schedule",
			Fields: fields_describe_maintenance_window_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowScheduleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_window_schedule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowSchedule(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowScheduleOutput
				p := svc.NewDescribeMaintenanceWindowSchedulePaginator(client, input)
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
		"describe-maintenance-window-targets": {
			Name:   "describe-maintenance-window-targets",
			Fields: fields_describe_maintenance_window_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_window_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowTargets(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowTargetsOutput
				p := svc.NewDescribeMaintenanceWindowTargetsPaginator(client, input)
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
		"describe-maintenance-window-tasks": {
			Name:   "describe-maintenance-window-tasks",
			Fields: fields_describe_maintenance_window_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_window_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowTasks(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowTasksOutput
				p := svc.NewDescribeMaintenanceWindowTasksPaginator(client, input)
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
		"describe-maintenance-windows": {
			Name:   "describe-maintenance-windows",
			Fields: fields_describe_maintenance_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_windows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindows(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowsOutput
				p := svc.NewDescribeMaintenanceWindowsPaginator(client, input)
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
		"describe-maintenance-windows-for-target": {
			Name:   "describe-maintenance-windows-for-target",
			Fields: fields_describe_maintenance_windows_for_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceWindowsForTargetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_maintenance_windows_for_target, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMaintenanceWindowsForTarget(ctx, input)
				}
				var results []*svc.DescribeMaintenanceWindowsForTargetOutput
				p := svc.NewDescribeMaintenanceWindowsForTargetPaginator(client, input)
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
		"describe-ops-items": {
			Name:   "describe-ops-items",
			Fields: fields_describe_ops_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOpsItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ops_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOpsItems(ctx, input)
				}
				var results []*svc.DescribeOpsItemsOutput
				p := svc.NewDescribeOpsItemsPaginator(client, input)
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
		"describe-parameters": {
			Name:   "describe-parameters",
			Fields: fields_describe_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeParametersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_parameters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeParameters(ctx, input)
				}
				var results []*svc.DescribeParametersOutput
				p := svc.NewDescribeParametersPaginator(client, input)
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
		"describe-patch-baselines": {
			Name:   "describe-patch-baselines",
			Fields: fields_describe_patch_baselines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePatchBaselinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_patch_baselines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePatchBaselines(ctx, input)
				}
				var results []*svc.DescribePatchBaselinesOutput
				p := svc.NewDescribePatchBaselinesPaginator(client, input)
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
		"describe-patch-group-state": {
			Name:   "describe-patch-group-state",
			Fields: fields_describe_patch_group_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePatchGroupStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_patch_group_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePatchGroupState(ctx, input)
			},
		},
		"describe-patch-groups": {
			Name:   "describe-patch-groups",
			Fields: fields_describe_patch_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePatchGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_patch_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePatchGroups(ctx, input)
				}
				var results []*svc.DescribePatchGroupsOutput
				p := svc.NewDescribePatchGroupsPaginator(client, input)
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
		"describe-patch-properties": {
			Name:   "describe-patch-properties",
			Fields: fields_describe_patch_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePatchPropertiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_patch_properties, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePatchProperties(ctx, input)
				}
				var results []*svc.DescribePatchPropertiesOutput
				p := svc.NewDescribePatchPropertiesPaginator(client, input)
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
		"describe-sessions": {
			Name:   "describe-sessions",
			Fields: fields_describe_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSessions(ctx, input)
				}
				var results []*svc.DescribeSessionsOutput
				p := svc.NewDescribeSessionsPaginator(client, input)
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
		"disassociate-ops-item-related-item": {
			Name:   "disassociate-ops-item-related-item",
			Fields: fields_disassociate_ops_item_related_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateOpsItemRelatedItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_ops_item_related_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateOpsItemRelatedItem(ctx, input)
			},
		},
		"get-access-token": {
			Name:   "get-access-token",
			Fields: fields_get_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessToken(ctx, input)
			},
		},
		"get-automation-execution": {
			Name:   "get-automation-execution",
			Fields: fields_get_automation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomationExecution(ctx, input)
			},
		},
		"get-calendar-state": {
			Name:   "get-calendar-state",
			Fields: fields_get_calendar_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCalendarStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_calendar_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCalendarState(ctx, input)
			},
		},
		"get-command-invocation": {
			Name:   "get-command-invocation",
			Fields: fields_get_command_invocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommandInvocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_command_invocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCommandInvocation(ctx, input)
			},
		},
		"get-connection-status": {
			Name:   "get-connection-status",
			Fields: fields_get_connection_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionStatus(ctx, input)
			},
		},
		"get-default-patch-baseline": {
			Name:   "get-default-patch-baseline",
			Fields: fields_get_default_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultPatchBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_patch_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultPatchBaseline(ctx, input)
			},
		},
		"get-deployable-patch-snapshot-for-instance": {
			Name:   "get-deployable-patch-snapshot-for-instance",
			Fields: fields_get_deployable_patch_snapshot_for_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeployablePatchSnapshotForInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployable_patch_snapshot_for_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployablePatchSnapshotForInstance(ctx, input)
			},
		},
		"get-document": {
			Name:   "get-document",
			Fields: fields_get_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocument(ctx, input)
			},
		},
		"get-execution-preview": {
			Name:   "get-execution-preview",
			Fields: fields_get_execution_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExecutionPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_execution_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExecutionPreview(ctx, input)
			},
		},
		"get-inventory": {
			Name:   "get-inventory",
			Fields: fields_get_inventory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInventoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_inventory, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInventory(ctx, input)
				}
				var results []*svc.GetInventoryOutput
				p := svc.NewGetInventoryPaginator(client, input)
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
		"get-inventory-schema": {
			Name:   "get-inventory-schema",
			Fields: fields_get_inventory_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInventorySchemaInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_inventory_schema, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInventorySchema(ctx, input)
				}
				var results []*svc.GetInventorySchemaOutput
				p := svc.NewGetInventorySchemaPaginator(client, input)
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
		"get-maintenance-window": {
			Name:   "get-maintenance-window",
			Fields: fields_get_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMaintenanceWindow(ctx, input)
			},
		},
		"get-maintenance-window-execution": {
			Name:   "get-maintenance-window-execution",
			Fields: fields_get_maintenance_window_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMaintenanceWindowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_maintenance_window_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMaintenanceWindowExecution(ctx, input)
			},
		},
		"get-maintenance-window-execution-task": {
			Name:   "get-maintenance-window-execution-task",
			Fields: fields_get_maintenance_window_execution_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMaintenanceWindowExecutionTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_maintenance_window_execution_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMaintenanceWindowExecutionTask(ctx, input)
			},
		},
		"get-maintenance-window-execution-task-invocation": {
			Name:   "get-maintenance-window-execution-task-invocation",
			Fields: fields_get_maintenance_window_execution_task_invocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMaintenanceWindowExecutionTaskInvocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_maintenance_window_execution_task_invocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMaintenanceWindowExecutionTaskInvocation(ctx, input)
			},
		},
		"get-maintenance-window-task": {
			Name:   "get-maintenance-window-task",
			Fields: fields_get_maintenance_window_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMaintenanceWindowTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_maintenance_window_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMaintenanceWindowTask(ctx, input)
			},
		},
		"get-ops-item": {
			Name:   "get-ops-item",
			Fields: fields_get_ops_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpsItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ops_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpsItem(ctx, input)
			},
		},
		"get-ops-metadata": {
			Name:   "get-ops-metadata",
			Fields: fields_get_ops_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpsMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ops_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpsMetadata(ctx, input)
			},
		},
		"get-ops-summary": {
			Name:   "get-ops-summary",
			Fields: fields_get_ops_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpsSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ops_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOpsSummary(ctx, input)
				}
				var results []*svc.GetOpsSummaryOutput
				p := svc.NewGetOpsSummaryPaginator(client, input)
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
		"get-parameter": {
			Name:   "get-parameter",
			Fields: fields_get_parameter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParameterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_parameter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParameter(ctx, input)
			},
		},
		"get-parameter-history": {
			Name:   "get-parameter-history",
			Fields: fields_get_parameter_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParameterHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_parameter_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetParameterHistory(ctx, input)
				}
				var results []*svc.GetParameterHistoryOutput
				p := svc.NewGetParameterHistoryPaginator(client, input)
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
		"get-parameters": {
			Name:   "get-parameters",
			Fields: fields_get_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParameters(ctx, input)
			},
		},
		"get-parameters-by-path": {
			Name:   "get-parameters-by-path",
			Fields: fields_get_parameters_by_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParametersByPathInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_parameters_by_path, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetParametersByPath(ctx, input)
				}
				var results []*svc.GetParametersByPathOutput
				p := svc.NewGetParametersByPathPaginator(client, input)
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
		"get-patch-baseline": {
			Name:   "get-patch-baseline",
			Fields: fields_get_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPatchBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_patch_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPatchBaseline(ctx, input)
			},
		},
		"get-patch-baseline-for-patch-group": {
			Name:   "get-patch-baseline-for-patch-group",
			Fields: fields_get_patch_baseline_for_patch_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPatchBaselineForPatchGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_patch_baseline_for_patch_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPatchBaselineForPatchGroup(ctx, input)
			},
		},
		"get-resource-policies": {
			Name:   "get-resource-policies",
			Fields: fields_get_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourcePolicies(ctx, input)
				}
				var results []*svc.GetResourcePoliciesOutput
				p := svc.NewGetResourcePoliciesPaginator(client, input)
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
		"get-service-setting": {
			Name:   "get-service-setting",
			Fields: fields_get_service_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceSetting(ctx, input)
			},
		},
		"label-parameter-version": {
			Name:   "label-parameter-version",
			Fields: fields_label_parameter_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LabelParameterVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_label_parameter_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.LabelParameterVersion(ctx, input)
			},
		},
		"list-association-versions": {
			Name:   "list-association-versions",
			Fields: fields_list_association_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociationVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_association_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociationVersions(ctx, input)
				}
				var results []*svc.ListAssociationVersionsOutput
				p := svc.NewListAssociationVersionsPaginator(client, input)
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
		"list-associations": {
			Name:   "list-associations",
			Fields: fields_list_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociations(ctx, input)
				}
				var results []*svc.ListAssociationsOutput
				p := svc.NewListAssociationsPaginator(client, input)
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
		"list-command-invocations": {
			Name:   "list-command-invocations",
			Fields: fields_list_command_invocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommandInvocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_command_invocations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommandInvocations(ctx, input)
				}
				var results []*svc.ListCommandInvocationsOutput
				p := svc.NewListCommandInvocationsPaginator(client, input)
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
		"list-commands": {
			Name:   "list-commands",
			Fields: fields_list_commands,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommandsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_commands, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommands(ctx, input)
				}
				var results []*svc.ListCommandsOutput
				p := svc.NewListCommandsPaginator(client, input)
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
		"list-compliance-items": {
			Name:   "list-compliance-items",
			Fields: fields_list_compliance_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComplianceItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compliance_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComplianceItems(ctx, input)
				}
				var results []*svc.ListComplianceItemsOutput
				p := svc.NewListComplianceItemsPaginator(client, input)
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
		"list-compliance-summaries": {
			Name:   "list-compliance-summaries",
			Fields: fields_list_compliance_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComplianceSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compliance_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComplianceSummaries(ctx, input)
				}
				var results []*svc.ListComplianceSummariesOutput
				p := svc.NewListComplianceSummariesPaginator(client, input)
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
		"list-document-metadata-history": {
			Name:   "list-document-metadata-history",
			Fields: fields_list_document_metadata_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentMetadataHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_document_metadata_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDocumentMetadataHistory(ctx, input)
			},
		},
		"list-document-versions": {
			Name:   "list-document-versions",
			Fields: fields_list_document_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_document_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDocumentVersions(ctx, input)
				}
				var results []*svc.ListDocumentVersionsOutput
				p := svc.NewListDocumentVersionsPaginator(client, input)
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
		"list-documents": {
			Name:   "list-documents",
			Fields: fields_list_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_documents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDocuments(ctx, input)
				}
				var results []*svc.ListDocumentsOutput
				p := svc.NewListDocumentsPaginator(client, input)
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
		"list-inventory-entries": {
			Name:   "list-inventory-entries",
			Fields: fields_list_inventory_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInventoryEntriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_inventory_entries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListInventoryEntries(ctx, input)
			},
		},
		"list-nodes": {
			Name:   "list-nodes",
			Fields: fields_list_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodes(ctx, input)
				}
				var results []*svc.ListNodesOutput
				p := svc.NewListNodesPaginator(client, input)
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
		"list-nodes-summary": {
			Name:   "list-nodes-summary",
			Fields: fields_list_nodes_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodesSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodes_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodesSummary(ctx, input)
				}
				var results []*svc.ListNodesSummaryOutput
				p := svc.NewListNodesSummaryPaginator(client, input)
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
		"list-ops-item-events": {
			Name:   "list-ops-item-events",
			Fields: fields_list_ops_item_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpsItemEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ops_item_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpsItemEvents(ctx, input)
				}
				var results []*svc.ListOpsItemEventsOutput
				p := svc.NewListOpsItemEventsPaginator(client, input)
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
		"list-ops-item-related-items": {
			Name:   "list-ops-item-related-items",
			Fields: fields_list_ops_item_related_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpsItemRelatedItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ops_item_related_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpsItemRelatedItems(ctx, input)
				}
				var results []*svc.ListOpsItemRelatedItemsOutput
				p := svc.NewListOpsItemRelatedItemsPaginator(client, input)
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
		"list-ops-metadata": {
			Name:   "list-ops-metadata",
			Fields: fields_list_ops_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpsMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ops_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpsMetadata(ctx, input)
				}
				var results []*svc.ListOpsMetadataOutput
				p := svc.NewListOpsMetadataPaginator(client, input)
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
		"list-resource-compliance-summaries": {
			Name:   "list-resource-compliance-summaries",
			Fields: fields_list_resource_compliance_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceComplianceSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_compliance_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceComplianceSummaries(ctx, input)
				}
				var results []*svc.ListResourceComplianceSummariesOutput
				p := svc.NewListResourceComplianceSummariesPaginator(client, input)
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
		"list-resource-data-sync": {
			Name:   "list-resource-data-sync",
			Fields: fields_list_resource_data_sync,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceDataSyncInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_data_sync, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceDataSync(ctx, input)
				}
				var results []*svc.ListResourceDataSyncOutput
				p := svc.NewListResourceDataSyncPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"modify-document-permission": {
			Name:   "modify-document-permission",
			Fields: fields_modify_document_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDocumentPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_document_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDocumentPermission(ctx, input)
			},
		},
		"put-compliance-items": {
			Name:   "put-compliance-items",
			Fields: fields_put_compliance_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutComplianceItemsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_compliance_items, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutComplianceItems(ctx, input)
			},
		},
		"put-inventory": {
			Name:   "put-inventory",
			Fields: fields_put_inventory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInventoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_inventory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInventory(ctx, input)
			},
		},
		"put-parameter": {
			Name:   "put-parameter",
			Fields: fields_put_parameter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutParameterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_parameter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutParameter(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"register-default-patch-baseline": {
			Name:   "register-default-patch-baseline",
			Fields: fields_register_default_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDefaultPatchBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_default_patch_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDefaultPatchBaseline(ctx, input)
			},
		},
		"register-patch-baseline-for-patch-group": {
			Name:   "register-patch-baseline-for-patch-group",
			Fields: fields_register_patch_baseline_for_patch_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterPatchBaselineForPatchGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_patch_baseline_for_patch_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterPatchBaselineForPatchGroup(ctx, input)
			},
		},
		"register-target-with-maintenance-window": {
			Name:   "register-target-with-maintenance-window",
			Fields: fields_register_target_with_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTargetWithMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_target_with_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTargetWithMaintenanceWindow(ctx, input)
			},
		},
		"register-task-with-maintenance-window": {
			Name:   "register-task-with-maintenance-window",
			Fields: fields_register_task_with_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTaskWithMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_task_with_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTaskWithMaintenanceWindow(ctx, input)
			},
		},
		"remove-tags-from-resource": {
			Name:   "remove-tags-from-resource",
			Fields: fields_remove_tags_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromResource(ctx, input)
			},
		},
		"reset-service-setting": {
			Name:   "reset-service-setting",
			Fields: fields_reset_service_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetServiceSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_service_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetServiceSetting(ctx, input)
			},
		},
		"resume-session": {
			Name:   "resume-session",
			Fields: fields_resume_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeSession(ctx, input)
			},
		},
		"send-automation-signal": {
			Name:   "send-automation-signal",
			Fields: fields_send_automation_signal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendAutomationSignalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_automation_signal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendAutomationSignal(ctx, input)
			},
		},
		"send-command": {
			Name:   "send-command",
			Fields: fields_send_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendCommand(ctx, input)
			},
		},
		"start-access-request": {
			Name:   "start-access-request",
			Fields: fields_start_access_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAccessRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_access_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAccessRequest(ctx, input)
			},
		},
		"start-associations-once": {
			Name:   "start-associations-once",
			Fields: fields_start_associations_once,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAssociationsOnceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_associations_once, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAssociationsOnce(ctx, input)
			},
		},
		"start-automation-execution": {
			Name:   "start-automation-execution",
			Fields: fields_start_automation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAutomationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_automation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAutomationExecution(ctx, input)
			},
		},
		"start-change-request-execution": {
			Name:   "start-change-request-execution",
			Fields: fields_start_change_request_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartChangeRequestExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_change_request_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartChangeRequestExecution(ctx, input)
			},
		},
		"start-execution-preview": {
			Name:   "start-execution-preview",
			Fields: fields_start_execution_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExecutionPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_execution_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExecutionPreview(ctx, input)
			},
		},
		"start-session": {
			Name:   "start-session",
			Fields: fields_start_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSession(ctx, input)
			},
		},
		"stop-automation-execution": {
			Name:   "stop-automation-execution",
			Fields: fields_stop_automation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAutomationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_automation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAutomationExecution(ctx, input)
			},
		},
		"terminate-session": {
			Name:   "terminate-session",
			Fields: fields_terminate_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateSession(ctx, input)
			},
		},
		"unlabel-parameter-version": {
			Name:   "unlabel-parameter-version",
			Fields: fields_unlabel_parameter_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnlabelParameterVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unlabel_parameter_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnlabelParameterVersion(ctx, input)
			},
		},
		"update-association": {
			Name:   "update-association",
			Fields: fields_update_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssociation(ctx, input)
			},
		},
		"update-association-status": {
			Name:   "update-association-status",
			Fields: fields_update_association_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssociationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_association_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssociationStatus(ctx, input)
			},
		},
		"update-document": {
			Name:   "update-document",
			Fields: fields_update_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocument(ctx, input)
			},
		},
		"update-document-default-version": {
			Name:   "update-document-default-version",
			Fields: fields_update_document_default_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentDefaultVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_document_default_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocumentDefaultVersion(ctx, input)
			},
		},
		"update-document-metadata": {
			Name:   "update-document-metadata",
			Fields: fields_update_document_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_document_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocumentMetadata(ctx, input)
			},
		},
		"update-maintenance-window": {
			Name:   "update-maintenance-window",
			Fields: fields_update_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMaintenanceWindow(ctx, input)
			},
		},
		"update-maintenance-window-target": {
			Name:   "update-maintenance-window-target",
			Fields: fields_update_maintenance_window_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMaintenanceWindowTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_maintenance_window_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMaintenanceWindowTarget(ctx, input)
			},
		},
		"update-maintenance-window-task": {
			Name:   "update-maintenance-window-task",
			Fields: fields_update_maintenance_window_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMaintenanceWindowTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_maintenance_window_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMaintenanceWindowTask(ctx, input)
			},
		},
		"update-managed-instance-role": {
			Name:   "update-managed-instance-role",
			Fields: fields_update_managed_instance_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateManagedInstanceRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_managed_instance_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateManagedInstanceRole(ctx, input)
			},
		},
		"update-ops-item": {
			Name:   "update-ops-item",
			Fields: fields_update_ops_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOpsItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ops_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOpsItem(ctx, input)
			},
		},
		"update-ops-metadata": {
			Name:   "update-ops-metadata",
			Fields: fields_update_ops_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOpsMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ops_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOpsMetadata(ctx, input)
			},
		},
		"update-patch-baseline": {
			Name:   "update-patch-baseline",
			Fields: fields_update_patch_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePatchBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_patch_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePatchBaseline(ctx, input)
			},
		},
		"update-resource-data-sync": {
			Name:   "update-resource-data-sync",
			Fields: fields_update_resource_data_sync,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceDataSyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_data_sync, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceDataSync(ctx, input)
			},
		},
		"update-service-setting": {
			Name:   "update-service-setting",
			Fields: fields_update_service_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSetting(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssm", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
