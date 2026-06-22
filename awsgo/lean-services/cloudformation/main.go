package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var fields_activate_organizations_access = []leanruntime.Field{}

var fields_activate_type = []leanruntime.Field{
	{Name: "AutoUpdate", Flag: "auto-update", Type: "*bool", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.LoggingConfig", Required: false},
	{Name: "MajorVersion", Flag: "major-version", Type: "*int64", Required: false},
	{Name: "PublicTypeArn", Flag: "public-type-arn", Type: "*string", Required: false},
	{Name: "PublisherId", Flag: "publisher-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ThirdPartyType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
	{Name: "TypeNameAlias", Flag: "type-name-alias", Type: "*string", Required: false},
	{Name: "VersionBump", Flag: "version-bump", Type: "types.VersionBump", Required: false},
}

var fields_batch_describe_type_configurations = []leanruntime.Field{
	{Name: "TypeConfigurationIdentifiers", Flag: "type-configuration-identifiers", Type: "[]types.TypeConfigurationIdentifier", Required: true},
}

var fields_cancel_update_stack = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_continue_update_rollback = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ResourcesToSkip", Flag: "resources-to-skip", Type: "[]string", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_create_change_set = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: false},
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: true},
	{Name: "ChangeSetType", Flag: "change-set-type", Type: "types.ChangeSetType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeploymentMode", Flag: "deployment-mode", Type: "types.DeploymentMode", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ImportExistingResources", Flag: "import-existing-resources", Type: "*bool", Required: false},
	{Name: "IncludeNestedStacks", Flag: "include-nested-stacks", Type: "*bool", Required: false},
	{Name: "NotificationARNs", Flag: "notification-arns", Type: "[]string", Required: false},
	{Name: "OnStackFailure", Flag: "on-stack-failure", Type: "types.OnStackFailure", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
	{Name: "ResourcesToImport", Flag: "resources-to-import", Type: "[]types.ResourceToImport", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RollbackConfiguration", Flag: "rollback-configuration", Type: "*types.RollbackConfiguration", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
	{Name: "UsePreviousTemplate", Flag: "use-previous-template", Type: "*bool", Required: false},
}

var fields_create_generated_template = []leanruntime.Field{
	{Name: "GeneratedTemplateName", Flag: "generated-template-name", Type: "*string", Required: true},
	{Name: "Resources", Flag: "resources", Type: "[]types.ResourceDefinition", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
	{Name: "TemplateConfiguration", Flag: "template-configuration", Type: "*types.TemplateConfiguration", Required: false},
}

var fields_create_stack = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DisableRollback", Flag: "disable-rollback", Type: "*bool", Required: false},
	{Name: "EnableTerminationProtection", Flag: "enable-termination-protection", Type: "*bool", Required: false},
	{Name: "NotificationARNs", Flag: "notification-arns", Type: "[]string", Required: false},
	{Name: "OnFailure", Flag: "on-failure", Type: "types.OnFailure", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
	{Name: "RetainExceptOnCreate", Flag: "retain-except-on-create", Type: "*bool", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RollbackConfiguration", Flag: "rollback-configuration", Type: "*types.RollbackConfiguration", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "StackPolicyBody", Flag: "stack-policy-body", Type: "*string", Required: false},
	{Name: "StackPolicyURL", Flag: "stack-policy-url", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
	{Name: "TimeoutInMinutes", Flag: "timeout-in-minutes", Type: "*int32", Required: false},
}

var fields_create_stack_instances = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]string", Required: false},
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "DeploymentTargets", Flag: "deployment-targets", Type: "*types.DeploymentTargets", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "OperationPreferences", Flag: "operation-preferences", Type: "*types.StackSetOperationPreferences", Required: false},
	{Name: "ParameterOverrides", Flag: "parameter-overrides", Type: "[]types.Parameter", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_create_stack_refactor = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnableStackCreation", Flag: "enable-stack-creation", Type: "*bool", Required: false},
	{Name: "ResourceMappings", Flag: "resource-mappings", Type: "[]types.ResourceMapping", Required: false},
	{Name: "StackDefinitions", Flag: "stack-definitions", Type: "[]types.StackDefinition", Required: true},
}

var fields_create_stack_set = []leanruntime.Field{
	{Name: "AdministrationRoleARN", Flag: "administration-role-arn", Type: "*string", Required: false},
	{Name: "AutoDeployment", Flag: "auto-deployment", Type: "*types.AutoDeployment", Required: false},
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRoleName", Flag: "execution-role-name", Type: "*string", Required: false},
	{Name: "ManagedExecution", Flag: "managed-execution", Type: "*types.ManagedExecution", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "PermissionModel", Flag: "permission-model", Type: "types.PermissionModels", Required: false},
	{Name: "StackId", Flag: "stack-id", Type: "*string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
}

var fields_deactivate_organizations_access = []leanruntime.Field{}

var fields_deactivate_type = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ThirdPartyType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
}

var fields_delete_change_set = []leanruntime.Field{
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_delete_generated_template = []leanruntime.Field{
	{Name: "GeneratedTemplateName", Flag: "generated-template-name", Type: "*string", Required: true},
}

var fields_delete_stack = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DeletionMode", Flag: "deletion-mode", Type: "types.DeletionMode", Required: false},
	{Name: "RetainResources", Flag: "retain-resources", Type: "[]string", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_delete_stack_instances = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]string", Required: false},
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "DeploymentTargets", Flag: "deployment-targets", Type: "*types.DeploymentTargets", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "OperationPreferences", Flag: "operation-preferences", Type: "*types.StackSetOperationPreferences", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
	{Name: "RetainStacks", Flag: "retain-stacks", Type: "*bool", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_delete_stack_set = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_deregister_type = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_describe_account_limits = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_change_set = []leanruntime.Field{
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: true},
	{Name: "IncludePropertyValues", Flag: "include-property-values", Type: "*bool", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_describe_change_set_hooks = []leanruntime.Field{
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: true},
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_describe_events = []leanruntime.Field{
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.EventFilter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_describe_generated_template = []leanruntime.Field{
	{Name: "GeneratedTemplateName", Flag: "generated-template-name", Type: "*string", Required: true},
}

var fields_describe_organizations_access = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
}

var fields_describe_publisher = []leanruntime.Field{
	{Name: "PublisherId", Flag: "publisher-id", Type: "*string", Required: false},
}

var fields_describe_resource_scan = []leanruntime.Field{
	{Name: "ResourceScanId", Flag: "resource-scan-id", Type: "*string", Required: true},
}

var fields_describe_stack_drift_detection_status = []leanruntime.Field{
	{Name: "StackDriftDetectionId", Flag: "stack-drift-detection-id", Type: "*string", Required: true},
}

var fields_describe_stack_events = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_describe_stack_instance = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "StackInstanceAccount", Flag: "stack-instance-account", Type: "*string", Required: true},
	{Name: "StackInstanceRegion", Flag: "stack-instance-region", Type: "*string", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_describe_stack_refactor = []leanruntime.Field{
	{Name: "StackRefactorId", Flag: "stack-refactor-id", Type: "*string", Required: true},
}

var fields_describe_stack_resource = []leanruntime.Field{
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_describe_stack_resource_drifts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "StackResourceDriftStatusFilters", Flag: "stack-resource-drift-status-filters", Type: "[]types.StackResourceDriftStatus", Required: false},
}

var fields_describe_stack_resources = []leanruntime.Field{
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*string", Required: false},
	{Name: "PhysicalResourceId", Flag: "physical-resource-id", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_describe_stack_set = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_describe_stack_set_operation = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_describe_stacks = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_describe_type = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "PublicVersionNumber", Flag: "public-version-number", Type: "*string", Required: false},
	{Name: "PublisherId", Flag: "publisher-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_describe_type_registration = []leanruntime.Field{
	{Name: "RegistrationToken", Flag: "registration-token", Type: "*string", Required: true},
}

var fields_detect_stack_drift = []leanruntime.Field{
	{Name: "LogicalResourceIds", Flag: "logical-resource-ids", Type: "[]string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_detect_stack_resource_drift = []leanruntime.Field{
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_detect_stack_set_drift = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "OperationPreferences", Flag: "operation-preferences", Type: "*types.StackSetOperationPreferences", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_estimate_template_cost = []leanruntime.Field{
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
}

var fields_execute_change_set = []leanruntime.Field{
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DisableRollback", Flag: "disable-rollback", Type: "*bool", Required: false},
	{Name: "RetainExceptOnCreate", Flag: "retain-except-on-create", Type: "*bool", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_execute_stack_refactor = []leanruntime.Field{
	{Name: "StackRefactorId", Flag: "stack-refactor-id", Type: "*string", Required: true},
}

var fields_get_generated_template = []leanruntime.Field{
	{Name: "Format", Flag: "format", Type: "types.TemplateFormat", Required: false},
	{Name: "GeneratedTemplateName", Flag: "generated-template-name", Type: "*string", Required: true},
}

var fields_get_hook_result = []leanruntime.Field{
	{Name: "HookResultId", Flag: "hook-result-id", Type: "*string", Required: false},
}

var fields_get_stack_policy = []leanruntime.Field{
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_get_template = []leanruntime.Field{
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
	{Name: "TemplateStage", Flag: "template-stage", Type: "types.TemplateStage", Required: false},
}

var fields_get_template_summary = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateSummaryConfig", Flag: "template-summary-config", Type: "*types.TemplateSummaryConfig", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
}

var fields_import_stacks_to_stack_set = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "OperationPreferences", Flag: "operation-preferences", Type: "*types.StackSetOperationPreferences", Required: false},
	{Name: "OrganizationalUnitIds", Flag: "organizational-unit-ids", Type: "[]string", Required: false},
	{Name: "StackIds", Flag: "stack-ids", Type: "[]string", Required: false},
	{Name: "StackIdsUrl", Flag: "stack-ids-url", Type: "*string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_list_change_sets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_list_exports = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_generated_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hook_results = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.HookStatus", Required: false},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: false},
	{Name: "TargetType", Flag: "target-type", Type: "types.ListHookResultsTargetType", Required: false},
	{Name: "TypeArn", Flag: "type-arn", Type: "*string", Required: false},
}

var fields_list_imports = []leanruntime.Field{
	{Name: "ExportName", Flag: "export-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_scan_related_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceScanId", Flag: "resource-scan-id", Type: "*string", Required: true},
	{Name: "Resources", Flag: "resources", Type: "[]types.ScannedResourceIdentifier", Required: true},
}

var fields_list_resource_scan_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "ResourceScanId", Flag: "resource-scan-id", Type: "*string", Required: true},
	{Name: "ResourceTypePrefix", Flag: "resource-type-prefix", Type: "*string", Required: false},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: false},
	{Name: "TagValue", Flag: "tag-value", Type: "*string", Required: false},
}

var fields_list_resource_scans = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanTypeFilter", Flag: "scan-type-filter", Type: "types.ScanType", Required: false},
}

var fields_list_stack_instance_resource_drifts = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
	{Name: "StackInstanceAccount", Flag: "stack-instance-account", Type: "*string", Required: true},
	{Name: "StackInstanceRegion", Flag: "stack-instance-region", Type: "*string", Required: true},
	{Name: "StackInstanceResourceDriftStatuses", Flag: "stack-instance-resource-drift-statuses", Type: "[]types.StackResourceDriftStatus", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_list_stack_instances = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.StackInstanceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackInstanceAccount", Flag: "stack-instance-account", Type: "*string", Required: false},
	{Name: "StackInstanceRegion", Flag: "stack-instance-region", Type: "*string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_list_stack_refactor_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackRefactorId", Flag: "stack-refactor-id", Type: "*string", Required: true},
}

var fields_list_stack_refactors = []leanruntime.Field{
	{Name: "ExecutionStatusFilter", Flag: "execution-status-filter", Type: "[]types.StackRefactorExecutionStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stack_resources = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_list_stack_set_auto_deployment_targets = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_list_stack_set_operation_results = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.OperationResultFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_list_stack_set_operations = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_list_stack_sets = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.StackSetStatus", Required: false},
}

var fields_list_stacks = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackStatusFilter", Flag: "stack-status-filter", Type: "[]types.StackStatus", Required: false},
}

var fields_list_type_registrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationStatusFilter", Flag: "registration-status-filter", Type: "types.RegistrationStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "TypeArn", Flag: "type-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
}

var fields_list_type_versions = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "DeprecatedStatus", Flag: "deprecated-status", Type: "types.DeprecatedStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PublisherId", Flag: "publisher-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
}

var fields_list_types = []leanruntime.Field{
	{Name: "DeprecatedStatus", Flag: "deprecated-status", Type: "types.DeprecatedStatus", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.TypeFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProvisioningType", Flag: "provisioning-type", Type: "types.ProvisioningType", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "Visibility", Flag: "visibility", Type: "types.Visibility", Required: false},
}

var fields_publish_type = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "PublicVersionNumber", Flag: "public-version-number", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ThirdPartyType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
}

var fields_record_handler_progress = []leanruntime.Field{
	{Name: "BearerToken", Flag: "bearer-token", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CurrentOperationStatus", Flag: "current-operation-status", Type: "types.OperationStatus", Required: false},
	{Name: "ErrorCode", Flag: "error-code", Type: "types.HandlerErrorCode", Required: false},
	{Name: "OperationStatus", Flag: "operation-status", Type: "types.OperationStatus", Required: true},
	{Name: "ResourceModel", Flag: "resource-model", Type: "*string", Required: false},
	{Name: "StatusMessage", Flag: "status-message", Type: "*string", Required: false},
}

var fields_register_publisher = []leanruntime.Field{
	{Name: "AcceptTermsAndConditions", Flag: "accept-terms-and-conditions", Type: "*bool", Required: false},
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: false},
}

var fields_register_type = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "LoggingConfig", Flag: "logging-config", Type: "*types.LoggingConfig", Required: false},
	{Name: "SchemaHandlerPackage", Flag: "schema-handler-package", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_rollback_stack = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "RetainExceptOnCreate", Flag: "retain-except-on-create", Type: "*bool", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_set_stack_policy = []leanruntime.Field{
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "StackPolicyBody", Flag: "stack-policy-body", Type: "*string", Required: false},
	{Name: "StackPolicyURL", Flag: "stack-policy-url", Type: "*string", Required: false},
}

var fields_set_type_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: true},
	{Name: "ConfigurationAlias", Flag: "configuration-alias", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ThirdPartyType", Required: false},
	{Name: "TypeArn", Flag: "type-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
}

var fields_set_type_default_version = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RegistryType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_signal_resource = []leanruntime.Field{
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ResourceSignalStatus", Required: true},
	{Name: "UniqueId", Flag: "unique-id", Type: "*string", Required: true},
}

var fields_start_resource_scan = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ScanFilters", Flag: "scan-filters", Type: "[]types.ScanFilter", Required: false},
}

var fields_stop_stack_set_operation = []leanruntime.Field{
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_test_type = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "LogDeliveryBucket", Flag: "log-delivery-bucket", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ThirdPartyType", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_update_generated_template = []leanruntime.Field{
	{Name: "AddResources", Flag: "add-resources", Type: "[]types.ResourceDefinition", Required: false},
	{Name: "GeneratedTemplateName", Flag: "generated-template-name", Type: "*string", Required: true},
	{Name: "NewGeneratedTemplateName", Flag: "new-generated-template-name", Type: "*string", Required: false},
	{Name: "RefreshAllResources", Flag: "refresh-all-resources", Type: "*bool", Required: false},
	{Name: "RemoveResources", Flag: "remove-resources", Type: "[]string", Required: false},
	{Name: "TemplateConfiguration", Flag: "template-configuration", Type: "*types.TemplateConfiguration", Required: false},
}

var fields_update_stack = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DisableRollback", Flag: "disable-rollback", Type: "*bool", Required: false},
	{Name: "NotificationARNs", Flag: "notification-arns", Type: "[]string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
	{Name: "RetainExceptOnCreate", Flag: "retain-except-on-create", Type: "*bool", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RollbackConfiguration", Flag: "rollback-configuration", Type: "*types.RollbackConfiguration", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "StackPolicyBody", Flag: "stack-policy-body", Type: "*string", Required: false},
	{Name: "StackPolicyDuringUpdateBody", Flag: "stack-policy-during-update-body", Type: "*string", Required: false},
	{Name: "StackPolicyDuringUpdateURL", Flag: "stack-policy-during-update-url", Type: "*string", Required: false},
	{Name: "StackPolicyURL", Flag: "stack-policy-url", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
	{Name: "UsePreviousTemplate", Flag: "use-previous-template", Type: "*bool", Required: false},
}

var fields_update_stack_instances = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]string", Required: false},
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "DeploymentTargets", Flag: "deployment-targets", Type: "*types.DeploymentTargets", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "OperationPreferences", Flag: "operation-preferences", Type: "*types.StackSetOperationPreferences", Required: false},
	{Name: "ParameterOverrides", Flag: "parameter-overrides", Type: "[]types.Parameter", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
}

var fields_update_stack_set = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]string", Required: false},
	{Name: "AdministrationRoleARN", Flag: "administration-role-arn", Type: "*string", Required: false},
	{Name: "AutoDeployment", Flag: "auto-deployment", Type: "*types.AutoDeployment", Required: false},
	{Name: "CallAs", Flag: "call-as", Type: "types.CallAs", Required: false},
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: false},
	{Name: "DeploymentTargets", Flag: "deployment-targets", Type: "*types.DeploymentTargets", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRoleName", Flag: "execution-role-name", Type: "*string", Required: false},
	{Name: "ManagedExecution", Flag: "managed-execution", Type: "*types.ManagedExecution", Required: false},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: false},
	{Name: "OperationPreferences", Flag: "operation-preferences", Type: "*types.StackSetOperationPreferences", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.Parameter", Required: false},
	{Name: "PermissionModel", Flag: "permission-model", Type: "types.PermissionModels", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
	{Name: "StackSetName", Flag: "stack-set-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
	{Name: "UsePreviousTemplate", Flag: "use-previous-template", Type: "*bool", Required: false},
}

var fields_update_termination_protection = []leanruntime.Field{
	{Name: "EnableTerminationProtection", Flag: "enable-termination-protection", Type: "*bool", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_validate_template = []leanruntime.Field{
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateURL", Flag: "template-url", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-organizations-access": {
			Name:   "activate-organizations-access",
			Fields: fields_activate_organizations_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateOrganizationsAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_organizations_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateOrganizationsAccess(ctx, input)
			},
		},
		"activate-type": {
			Name:   "activate-type",
			Fields: fields_activate_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateType(ctx, input)
			},
		},
		"batch-describe-type-configurations": {
			Name:   "batch-describe-type-configurations",
			Fields: fields_batch_describe_type_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDescribeTypeConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_describe_type_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDescribeTypeConfigurations(ctx, input)
			},
		},
		"cancel-update-stack": {
			Name:   "cancel-update-stack",
			Fields: fields_cancel_update_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelUpdateStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_update_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelUpdateStack(ctx, input)
			},
		},
		"continue-update-rollback": {
			Name:   "continue-update-rollback",
			Fields: fields_continue_update_rollback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ContinueUpdateRollbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_continue_update_rollback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ContinueUpdateRollback(ctx, input)
			},
		},
		"create-change-set": {
			Name:   "create-change-set",
			Fields: fields_create_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChangeSet(ctx, input)
			},
		},
		"create-generated-template": {
			Name:   "create-generated-template",
			Fields: fields_create_generated_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGeneratedTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_generated_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGeneratedTemplate(ctx, input)
			},
		},
		"create-stack": {
			Name:   "create-stack",
			Fields: fields_create_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStack(ctx, input)
			},
		},
		"create-stack-instances": {
			Name:   "create-stack-instances",
			Fields: fields_create_stack_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStackInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stack_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStackInstances(ctx, input)
			},
		},
		"create-stack-refactor": {
			Name:   "create-stack-refactor",
			Fields: fields_create_stack_refactor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStackRefactorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stack_refactor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStackRefactor(ctx, input)
			},
		},
		"create-stack-set": {
			Name:   "create-stack-set",
			Fields: fields_create_stack_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStackSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stack_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStackSet(ctx, input)
			},
		},
		"deactivate-organizations-access": {
			Name:   "deactivate-organizations-access",
			Fields: fields_deactivate_organizations_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateOrganizationsAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_organizations_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateOrganizationsAccess(ctx, input)
			},
		},
		"deactivate-type": {
			Name:   "deactivate-type",
			Fields: fields_deactivate_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateType(ctx, input)
			},
		},
		"delete-change-set": {
			Name:   "delete-change-set",
			Fields: fields_delete_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChangeSet(ctx, input)
			},
		},
		"delete-generated-template": {
			Name:   "delete-generated-template",
			Fields: fields_delete_generated_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGeneratedTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_generated_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGeneratedTemplate(ctx, input)
			},
		},
		"delete-stack": {
			Name:   "delete-stack",
			Fields: fields_delete_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStack(ctx, input)
			},
		},
		"delete-stack-instances": {
			Name:   "delete-stack-instances",
			Fields: fields_delete_stack_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStackInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stack_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStackInstances(ctx, input)
			},
		},
		"delete-stack-set": {
			Name:   "delete-stack-set",
			Fields: fields_delete_stack_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStackSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stack_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStackSet(ctx, input)
			},
		},
		"deregister-type": {
			Name:   "deregister-type",
			Fields: fields_deregister_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterType(ctx, input)
			},
		},
		"describe-account-limits": {
			Name:   "describe-account-limits",
			Fields: fields_describe_account_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_account_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAccountLimits(ctx, input)
				}
				var results []*svc.DescribeAccountLimitsOutput
				p := svc.NewDescribeAccountLimitsPaginator(client, input)
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
		"describe-change-set": {
			Name:   "describe-change-set",
			Fields: fields_describe_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChangeSetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_change_set, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeChangeSet(ctx, input)
				}
				var results []*svc.DescribeChangeSetOutput
				p := svc.NewDescribeChangeSetPaginator(client, input)
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
		"describe-change-set-hooks": {
			Name:   "describe-change-set-hooks",
			Fields: fields_describe_change_set_hooks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChangeSetHooksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_change_set_hooks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChangeSetHooks(ctx, input)
			},
		},
		"describe-events": {
			Name:   "describe-events",
			Fields: fields_describe_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEvents(ctx, input)
				}
				var results []*svc.DescribeEventsOutput
				p := svc.NewDescribeEventsPaginator(client, input)
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
		"describe-generated-template": {
			Name:   "describe-generated-template",
			Fields: fields_describe_generated_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGeneratedTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_generated_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGeneratedTemplate(ctx, input)
			},
		},
		"describe-organizations-access": {
			Name:   "describe-organizations-access",
			Fields: fields_describe_organizations_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationsAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organizations_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganizationsAccess(ctx, input)
			},
		},
		"describe-publisher": {
			Name:   "describe-publisher",
			Fields: fields_describe_publisher,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePublisherInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_publisher, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePublisher(ctx, input)
			},
		},
		"describe-resource-scan": {
			Name:   "describe-resource-scan",
			Fields: fields_describe_resource_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourceScan(ctx, input)
			},
		},
		"describe-stack-drift-detection-status": {
			Name:   "describe-stack-drift-detection-status",
			Fields: fields_describe_stack_drift_detection_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackDriftDetectionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_drift_detection_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackDriftDetectionStatus(ctx, input)
			},
		},
		"describe-stack-events": {
			Name:   "describe-stack-events",
			Fields: fields_describe_stack_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_stack_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStackEvents(ctx, input)
				}
				var results []*svc.DescribeStackEventsOutput
				p := svc.NewDescribeStackEventsPaginator(client, input)
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
		"describe-stack-instance": {
			Name:   "describe-stack-instance",
			Fields: fields_describe_stack_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackInstance(ctx, input)
			},
		},
		"describe-stack-refactor": {
			Name:   "describe-stack-refactor",
			Fields: fields_describe_stack_refactor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackRefactorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_refactor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackRefactor(ctx, input)
			},
		},
		"describe-stack-resource": {
			Name:   "describe-stack-resource",
			Fields: fields_describe_stack_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackResource(ctx, input)
			},
		},
		"describe-stack-resource-drifts": {
			Name:   "describe-stack-resource-drifts",
			Fields: fields_describe_stack_resource_drifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackResourceDriftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_stack_resource_drifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStackResourceDrifts(ctx, input)
				}
				var results []*svc.DescribeStackResourceDriftsOutput
				p := svc.NewDescribeStackResourceDriftsPaginator(client, input)
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
		"describe-stack-resources": {
			Name:   "describe-stack-resources",
			Fields: fields_describe_stack_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackResources(ctx, input)
			},
		},
		"describe-stack-set": {
			Name:   "describe-stack-set",
			Fields: fields_describe_stack_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackSet(ctx, input)
			},
		},
		"describe-stack-set-operation": {
			Name:   "describe-stack-set-operation",
			Fields: fields_describe_stack_set_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStackSetOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stack_set_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStackSetOperation(ctx, input)
			},
		},
		"describe-stacks": {
			Name:   "describe-stacks",
			Fields: fields_describe_stacks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStacksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_stacks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStacks(ctx, input)
				}
				var results []*svc.DescribeStacksOutput
				p := svc.NewDescribeStacksPaginator(client, input)
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
		"describe-type": {
			Name:   "describe-type",
			Fields: fields_describe_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeType(ctx, input)
			},
		},
		"describe-type-registration": {
			Name:   "describe-type-registration",
			Fields: fields_describe_type_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTypeRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_type_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTypeRegistration(ctx, input)
			},
		},
		"detect-stack-drift": {
			Name:   "detect-stack-drift",
			Fields: fields_detect_stack_drift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectStackDriftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_stack_drift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectStackDrift(ctx, input)
			},
		},
		"detect-stack-resource-drift": {
			Name:   "detect-stack-resource-drift",
			Fields: fields_detect_stack_resource_drift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectStackResourceDriftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_stack_resource_drift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectStackResourceDrift(ctx, input)
			},
		},
		"detect-stack-set-drift": {
			Name:   "detect-stack-set-drift",
			Fields: fields_detect_stack_set_drift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectStackSetDriftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_stack_set_drift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectStackSetDrift(ctx, input)
			},
		},
		"estimate-template-cost": {
			Name:   "estimate-template-cost",
			Fields: fields_estimate_template_cost,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EstimateTemplateCostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_estimate_template_cost, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EstimateTemplateCost(ctx, input)
			},
		},
		"execute-change-set": {
			Name:   "execute-change-set",
			Fields: fields_execute_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteChangeSet(ctx, input)
			},
		},
		"execute-stack-refactor": {
			Name:   "execute-stack-refactor",
			Fields: fields_execute_stack_refactor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteStackRefactorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_stack_refactor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteStackRefactor(ctx, input)
			},
		},
		"get-generated-template": {
			Name:   "get-generated-template",
			Fields: fields_get_generated_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGeneratedTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_generated_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGeneratedTemplate(ctx, input)
			},
		},
		"get-hook-result": {
			Name:   "get-hook-result",
			Fields: fields_get_hook_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHookResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hook_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHookResult(ctx, input)
			},
		},
		"get-stack-policy": {
			Name:   "get-stack-policy",
			Fields: fields_get_stack_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStackPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stack_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStackPolicy(ctx, input)
			},
		},
		"get-template": {
			Name:   "get-template",
			Fields: fields_get_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplate(ctx, input)
			},
		},
		"get-template-summary": {
			Name:   "get-template-summary",
			Fields: fields_get_template_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplateSummary(ctx, input)
			},
		},
		"import-stacks-to-stack-set": {
			Name:   "import-stacks-to-stack-set",
			Fields: fields_import_stacks_to_stack_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportStacksToStackSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_stacks_to_stack_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportStacksToStackSet(ctx, input)
			},
		},
		"list-change-sets": {
			Name:   "list-change-sets",
			Fields: fields_list_change_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChangeSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_change_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChangeSets(ctx, input)
				}
				var results []*svc.ListChangeSetsOutput
				p := svc.NewListChangeSetsPaginator(client, input)
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
		"list-exports": {
			Name:   "list-exports",
			Fields: fields_list_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExports(ctx, input)
				}
				var results []*svc.ListExportsOutput
				p := svc.NewListExportsPaginator(client, input)
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
		"list-generated-templates": {
			Name:   "list-generated-templates",
			Fields: fields_list_generated_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGeneratedTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_generated_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGeneratedTemplates(ctx, input)
				}
				var results []*svc.ListGeneratedTemplatesOutput
				p := svc.NewListGeneratedTemplatesPaginator(client, input)
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
		"list-hook-results": {
			Name:   "list-hook-results",
			Fields: fields_list_hook_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHookResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hook_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHookResults(ctx, input)
			},
		},
		"list-imports": {
			Name:   "list-imports",
			Fields: fields_list_imports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_imports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImports(ctx, input)
				}
				var results []*svc.ListImportsOutput
				p := svc.NewListImportsPaginator(client, input)
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
		"list-resource-scan-related-resources": {
			Name:   "list-resource-scan-related-resources",
			Fields: fields_list_resource_scan_related_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceScanRelatedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_scan_related_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceScanRelatedResources(ctx, input)
				}
				var results []*svc.ListResourceScanRelatedResourcesOutput
				p := svc.NewListResourceScanRelatedResourcesPaginator(client, input)
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
		"list-resource-scan-resources": {
			Name:   "list-resource-scan-resources",
			Fields: fields_list_resource_scan_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceScanResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_scan_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceScanResources(ctx, input)
				}
				var results []*svc.ListResourceScanResourcesOutput
				p := svc.NewListResourceScanResourcesPaginator(client, input)
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
		"list-resource-scans": {
			Name:   "list-resource-scans",
			Fields: fields_list_resource_scans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceScansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_scans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceScans(ctx, input)
				}
				var results []*svc.ListResourceScansOutput
				p := svc.NewListResourceScansPaginator(client, input)
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
		"list-stack-instance-resource-drifts": {
			Name:   "list-stack-instance-resource-drifts",
			Fields: fields_list_stack_instance_resource_drifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackInstanceResourceDriftsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_stack_instance_resource_drifts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStackInstanceResourceDrifts(ctx, input)
			},
		},
		"list-stack-instances": {
			Name:   "list-stack-instances",
			Fields: fields_list_stack_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackInstances(ctx, input)
				}
				var results []*svc.ListStackInstancesOutput
				p := svc.NewListStackInstancesPaginator(client, input)
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
		"list-stack-refactor-actions": {
			Name:   "list-stack-refactor-actions",
			Fields: fields_list_stack_refactor_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackRefactorActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_refactor_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackRefactorActions(ctx, input)
				}
				var results []*svc.ListStackRefactorActionsOutput
				p := svc.NewListStackRefactorActionsPaginator(client, input)
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
		"list-stack-refactors": {
			Name:   "list-stack-refactors",
			Fields: fields_list_stack_refactors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackRefactorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_refactors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackRefactors(ctx, input)
				}
				var results []*svc.ListStackRefactorsOutput
				p := svc.NewListStackRefactorsPaginator(client, input)
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
		"list-stack-resources": {
			Name:   "list-stack-resources",
			Fields: fields_list_stack_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackResources(ctx, input)
				}
				var results []*svc.ListStackResourcesOutput
				p := svc.NewListStackResourcesPaginator(client, input)
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
		"list-stack-set-auto-deployment-targets": {
			Name:   "list-stack-set-auto-deployment-targets",
			Fields: fields_list_stack_set_auto_deployment_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackSetAutoDeploymentTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_stack_set_auto_deployment_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStackSetAutoDeploymentTargets(ctx, input)
			},
		},
		"list-stack-set-operation-results": {
			Name:   "list-stack-set-operation-results",
			Fields: fields_list_stack_set_operation_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackSetOperationResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_set_operation_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackSetOperationResults(ctx, input)
				}
				var results []*svc.ListStackSetOperationResultsOutput
				p := svc.NewListStackSetOperationResultsPaginator(client, input)
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
		"list-stack-set-operations": {
			Name:   "list-stack-set-operations",
			Fields: fields_list_stack_set_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackSetOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_set_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackSetOperations(ctx, input)
				}
				var results []*svc.ListStackSetOperationsOutput
				p := svc.NewListStackSetOperationsPaginator(client, input)
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
		"list-stack-sets": {
			Name:   "list-stack-sets",
			Fields: fields_list_stack_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stack_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStackSets(ctx, input)
				}
				var results []*svc.ListStackSetsOutput
				p := svc.NewListStackSetsPaginator(client, input)
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
		"list-stacks": {
			Name:   "list-stacks",
			Fields: fields_list_stacks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStacksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stacks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStacks(ctx, input)
				}
				var results []*svc.ListStacksOutput
				p := svc.NewListStacksPaginator(client, input)
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
		"list-type-registrations": {
			Name:   "list-type-registrations",
			Fields: fields_list_type_registrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypeRegistrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_type_registrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypeRegistrations(ctx, input)
				}
				var results []*svc.ListTypeRegistrationsOutput
				p := svc.NewListTypeRegistrationsPaginator(client, input)
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
		"list-type-versions": {
			Name:   "list-type-versions",
			Fields: fields_list_type_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypeVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_type_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypeVersions(ctx, input)
				}
				var results []*svc.ListTypeVersionsOutput
				p := svc.NewListTypeVersionsPaginator(client, input)
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
		"list-types": {
			Name:   "list-types",
			Fields: fields_list_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypes(ctx, input)
				}
				var results []*svc.ListTypesOutput
				p := svc.NewListTypesPaginator(client, input)
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
		"publish-type": {
			Name:   "publish-type",
			Fields: fields_publish_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishType(ctx, input)
			},
		},
		"record-handler-progress": {
			Name:   "record-handler-progress",
			Fields: fields_record_handler_progress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecordHandlerProgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_record_handler_progress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecordHandlerProgress(ctx, input)
			},
		},
		"register-publisher": {
			Name:   "register-publisher",
			Fields: fields_register_publisher,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterPublisherInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_publisher, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterPublisher(ctx, input)
			},
		},
		"register-type": {
			Name:   "register-type",
			Fields: fields_register_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterType(ctx, input)
			},
		},
		"rollback-stack": {
			Name:   "rollback-stack",
			Fields: fields_rollback_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RollbackStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rollback_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RollbackStack(ctx, input)
			},
		},
		"set-stack-policy": {
			Name:   "set-stack-policy",
			Fields: fields_set_stack_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetStackPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_stack_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetStackPolicy(ctx, input)
			},
		},
		"set-type-configuration": {
			Name:   "set-type-configuration",
			Fields: fields_set_type_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTypeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_type_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTypeConfiguration(ctx, input)
			},
		},
		"set-type-default-version": {
			Name:   "set-type-default-version",
			Fields: fields_set_type_default_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTypeDefaultVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_type_default_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTypeDefaultVersion(ctx, input)
			},
		},
		"signal-resource": {
			Name:   "signal-resource",
			Fields: fields_signal_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SignalResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_signal_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SignalResource(ctx, input)
			},
		},
		"start-resource-scan": {
			Name:   "start-resource-scan",
			Fields: fields_start_resource_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartResourceScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_resource_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartResourceScan(ctx, input)
			},
		},
		"stop-stack-set-operation": {
			Name:   "stop-stack-set-operation",
			Fields: fields_stop_stack_set_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopStackSetOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_stack_set_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopStackSetOperation(ctx, input)
			},
		},
		"test-type": {
			Name:   "test-type",
			Fields: fields_test_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestType(ctx, input)
			},
		},
		"update-generated-template": {
			Name:   "update-generated-template",
			Fields: fields_update_generated_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGeneratedTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_generated_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGeneratedTemplate(ctx, input)
			},
		},
		"update-stack": {
			Name:   "update-stack",
			Fields: fields_update_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStack(ctx, input)
			},
		},
		"update-stack-instances": {
			Name:   "update-stack-instances",
			Fields: fields_update_stack_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStackInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stack_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStackInstances(ctx, input)
			},
		},
		"update-stack-set": {
			Name:   "update-stack-set",
			Fields: fields_update_stack_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStackSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stack_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStackSet(ctx, input)
			},
		},
		"update-termination-protection": {
			Name:   "update-termination-protection",
			Fields: fields_update_termination_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTerminationProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_termination_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTerminationProtection(ctx, input)
			},
		},
		"validate-template": {
			Name:   "validate-template",
			Fields: fields_validate_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateTemplate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudformation", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
