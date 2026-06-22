package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/quicksight"
)

var fields_batch_create_topic_reviewed_answer = []leanruntime.Field{
	{Name: "Answers", Flag: "answers", Type: "[]types.CreateTopicReviewedAnswer", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_batch_delete_topic_reviewed_answer = []leanruntime.Field{
	{Name: "AnswerIds", Flag: "answer-ids", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_cancel_ingestion = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "IngestionId", Flag: "ingestion-id", Type: "*string", Required: true},
}

var fields_create_account_customization = []leanruntime.Field{
	{Name: "AccountCustomization", Flag: "account-customization", Type: "*types.AccountCustomization", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_account_subscription = []leanruntime.Field{
	{Name: "AccountName", Flag: "account-name", Type: "*string", Required: true},
	{Name: "ActiveDirectoryName", Flag: "active-directory-name", Type: "*string", Required: false},
	{Name: "AdminGroup", Flag: "admin-group", Type: "[]string", Required: false},
	{Name: "AdminProGroup", Flag: "admin-pro-group", Type: "[]string", Required: false},
	{Name: "AuthenticationMethod", Flag: "authentication-method", Type: "types.AuthenticationMethodOption", Required: true},
	{Name: "AuthorGroup", Flag: "author-group", Type: "[]string", Required: false},
	{Name: "AuthorProGroup", Flag: "author-pro-group", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ContactNumber", Flag: "contact-number", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Edition", Flag: "edition", Type: "types.Edition", Required: false},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: false},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "IAMIdentityCenterInstanceArn", Flag: "iam-identity-center-instance-arn", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "NotificationEmail", Flag: "notification-email", Type: "*string", Required: true},
	{Name: "ReaderGroup", Flag: "reader-group", Type: "[]string", Required: false},
	{Name: "ReaderProGroup", Flag: "reader-pro-group", Type: "[]string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
}

var fields_create_action_connector = []leanruntime.Field{
	{Name: "ActionConnectorId", Flag: "action-connector-id", Type: "*string", Required: true},
	{Name: "AuthenticationConfig", Flag: "authentication-config", Type: "*types.AuthConfig", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ActionConnectorType", Required: true},
	{Name: "VpcConnectionArn", Flag: "vpc-connection-arn", Type: "*string", Required: false},
}

var fields_create_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.AnalysisDefinition", Required: false},
	{Name: "FolderArns", Flag: "folder-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*types.Parameters", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "SourceEntity", Flag: "source-entity", Type: "*types.AnalysisSourceEntity", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThemeArn", Flag: "theme-arn", Type: "*string", Required: false},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.ValidationStrategy", Required: false},
}

var fields_create_brand = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandDefinition", Flag: "brand-definition", Type: "*types.BrandDefinition", Required: false},
	{Name: "BrandId", Flag: "brand-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Capabilities", Flag: "capabilities", Type: "*types.Capabilities", Required: false},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_dashboard = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "DashboardPublishOptions", Flag: "dashboard-publish-options", Type: "*types.DashboardPublishOptions", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.DashboardVersionDefinition", Required: false},
	{Name: "FolderArns", Flag: "folder-arns", Type: "[]string", Required: false},
	{Name: "LinkEntities", Flag: "link-entities", Type: "[]string", Required: false},
	{Name: "LinkSharingConfiguration", Flag: "link-sharing-configuration", Type: "*types.LinkSharingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*types.Parameters", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "SourceEntity", Flag: "source-entity", Type: "*types.DashboardSourceEntity", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThemeArn", Flag: "theme-arn", Type: "*string", Required: false},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.ValidationStrategy", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_data_set = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ColumnGroups", Flag: "column-groups", Type: "[]types.ColumnGroup", Required: false},
	{Name: "ColumnLevelPermissionRules", Flag: "column-level-permission-rules", Type: "[]types.ColumnLevelPermissionRule", Required: false},
	{Name: "DataPrepConfiguration", Flag: "data-prep-configuration", Type: "*types.DataPrepConfiguration", Required: false},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "DataSetUsageConfiguration", Flag: "data-set-usage-configuration", Type: "*types.DataSetUsageConfiguration", Required: false},
	{Name: "DatasetParameters", Flag: "dataset-parameters", Type: "[]types.DatasetParameter", Required: false},
	{Name: "FieldFolders", Flag: "field-folders", Type: "map[string]types.FieldFolder", Required: false},
	{Name: "FolderArns", Flag: "folder-arns", Type: "[]string", Required: false},
	{Name: "ImportMode", Flag: "import-mode", Type: "types.DataSetImportMode", Required: true},
	{Name: "LogicalTableMap", Flag: "logical-table-map", Type: "map[string]types.LogicalTable", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PerformanceConfiguration", Flag: "performance-configuration", Type: "*types.PerformanceConfiguration", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "PhysicalTableMap", Flag: "physical-table-map", Type: "map[string]types.PhysicalTable", Required: true},
	{Name: "RowLevelPermissionDataSet", Flag: "row-level-permission-data-set", Type: "*types.RowLevelPermissionDataSet", Required: false},
	{Name: "RowLevelPermissionTagConfiguration", Flag: "row-level-permission-tag-configuration", Type: "*types.RowLevelPermissionTagConfiguration", Required: false},
	{Name: "SemanticModelConfiguration", Flag: "semantic-model-configuration", Type: "*types.SemanticModelConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseAs", Flag: "use-as", Type: "types.DataSetUseAs", Required: false},
}

var fields_create_data_source = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Credentials", Flag: "credentials", Type: "*types.DataSourceCredentials", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DataSourceParameters", Flag: "data-source-parameters", Type: "types.DataSourceParameters", Required: false},
	{Name: "FolderArns", Flag: "folder-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "SslProperties", Flag: "ssl-properties", Type: "*types.SslProperties", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataSourceType", Required: true},
	{Name: "VpcConnectionProperties", Flag: "vpc-connection-properties", Type: "*types.VpcConnectionProperties", Required: false},
}

var fields_create_folder = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "FolderType", Flag: "folder-type", Type: "types.FolderType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentFolderArn", Flag: "parent-folder-arn", Type: "*string", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "SharingModel", Flag: "sharing-model", Type: "types.SharingModel", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_folder_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "MemberType", Flag: "member-type", Type: "types.MemberType", Required: true},
}

var fields_create_group = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_create_group_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_create_iam_policy_assignment = []leanruntime.Field{
	{Name: "AssignmentName", Flag: "assignment-name", Type: "*string", Required: true},
	{Name: "AssignmentStatus", Flag: "assignment-status", Type: "types.AssignmentStatus", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Identities", Flag: "identities", Type: "map[string][]string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: false},
}

var fields_create_ingestion = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "IngestionId", Flag: "ingestion-id", Type: "*string", Required: true},
	{Name: "IngestionType", Flag: "ingestion-type", Type: "types.IngestionType", Required: false},
}

var fields_create_namespace = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "IdentityStore", Flag: "identity-store", Type: "types.IdentityStore", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.RefreshSchedule", Required: true},
}

var fields_create_role_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
}

var fields_create_template = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.TemplateVersionDefinition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "SourceEntity", Flag: "source-entity", Type: "*types.TemplateSourceEntity", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.ValidationStrategy", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_template_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "TemplateVersionNumber", Flag: "template-version-number", Type: "*int64", Required: true},
}

var fields_create_theme = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BaseThemeId", Flag: "base-theme-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ThemeConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_theme_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
	{Name: "ThemeVersionNumber", Flag: "theme-version-number", Type: "*int64", Required: true},
}

var fields_create_topic = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomInstructions", Flag: "custom-instructions", Type: "*types.CustomInstructions", Required: false},
	{Name: "FolderArns", Flag: "folder-arns", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Topic", Flag: "topic", Type: "*types.TopicDetails", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_create_topic_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: false},
	{Name: "RefreshSchedule", Flag: "refresh-schedule", Type: "*types.TopicRefreshSchedule", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_create_vpc_connection = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DnsResolvers", Flag: "dns-resolvers", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VPCConnectionId", Flag: "vpc-connection-id", Type: "*string", Required: true},
}

var fields_delete_account_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_delete_account_customization = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
}

var fields_delete_account_subscription = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_delete_action_connector = []leanruntime.Field{
	{Name: "ActionConnectorId", Flag: "action-connector-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_delete_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ForceDeleteWithoutRecovery", Flag: "force-delete-without-recovery", Type: "bool", Required: false},
	{Name: "RecoveryWindowInDays", Flag: "recovery-window-in-days", Type: "*int64", Required: false},
}

var fields_delete_brand = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandId", Flag: "brand-id", Type: "*string", Required: true},
}

var fields_delete_brand_assignment = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_delete_custom_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
}

var fields_delete_dashboard = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_delete_data_set = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_delete_data_set_refresh_properties = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
}

var fields_delete_default_qbusiness_application = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
}

var fields_delete_folder = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
}

var fields_delete_folder_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "MemberType", Flag: "member-type", Type: "types.MemberType", Required: true},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_delete_group_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_delete_iam_policy_assignment = []leanruntime.Field{
	{Name: "AssignmentName", Flag: "assignment-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_delete_identity_propagation_config = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Service", Flag: "service", Type: "types.ServiceType", Required: true},
}

var fields_delete_namespace = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_delete_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "ScheduleId", Flag: "schedule-id", Type: "*string", Required: true},
}

var fields_delete_role_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
}

var fields_delete_role_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
}

var fields_delete_template = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_delete_template_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_delete_theme = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_delete_theme_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
}

var fields_delete_topic = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_delete_topic_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_user_by_principal_id = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
}

var fields_delete_user_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_vpc_connection = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "VPCConnectionId", Flag: "vpc-connection-id", Type: "*string", Required: true},
}

var fields_describe_account_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_account_customization = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Resolved", Flag: "resolved", Type: "bool", Required: false},
}

var fields_describe_account_settings = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_account_subscription = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_action_connector = []leanruntime.Field{
	{Name: "ActionConnectorId", Flag: "action-connector-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_action_connector_permissions = []leanruntime.Field{
	{Name: "ActionConnectorId", Flag: "action-connector-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_analysis_definition = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_analysis_permissions = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_asset_bundle_export_job = []leanruntime.Field{
	{Name: "AssetBundleExportJobId", Flag: "asset-bundle-export-job-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_asset_bundle_import_job = []leanruntime.Field{
	{Name: "AssetBundleImportJobId", Flag: "asset-bundle-import-job-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_brand = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandId", Flag: "brand-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_describe_brand_assignment = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_brand_published_version = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandId", Flag: "brand-id", Type: "*string", Required: true},
}

var fields_describe_custom_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
}

var fields_describe_dashboard = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_describe_dashboard_definition = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_describe_dashboard_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
}

var fields_describe_dashboard_snapshot_job = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "SnapshotJobId", Flag: "snapshot-job-id", Type: "*string", Required: true},
}

var fields_describe_dashboard_snapshot_job_result = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "SnapshotJobId", Flag: "snapshot-job-id", Type: "*string", Required: true},
}

var fields_describe_dashboards_qa_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_data_set = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_describe_data_set_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_describe_data_set_refresh_properties = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_describe_data_source = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
}

var fields_describe_data_source_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
}

var fields_describe_default_qbusiness_application = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
}

var fields_describe_folder = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
}

var fields_describe_folder_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_folder_resolved_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_group = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_describe_group_membership = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_describe_iam_policy_assignment = []leanruntime.Field{
	{Name: "AssignmentName", Flag: "assignment-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_describe_ingestion = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "IngestionId", Flag: "ingestion-id", Type: "*string", Required: true},
}

var fields_describe_ip_restriction = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_key_registration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DefaultKeyOnly", Flag: "default-key-only", Type: "bool", Required: false},
}

var fields_describe_namespace = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_describe_qpersonalization_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_quicksight_qsearch_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
}

var fields_describe_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "ScheduleId", Flag: "schedule-id", Type: "*string", Required: true},
}

var fields_describe_role_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
}

var fields_describe_self_upgrade_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_describe_template = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_describe_template_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_describe_template_definition = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_describe_template_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_describe_theme = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_describe_theme_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
}

var fields_describe_theme_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
}

var fields_describe_topic = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_describe_topic_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_describe_topic_refresh = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "RefreshId", Flag: "refresh-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_describe_topic_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_describe_vpc_connection = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "VPCConnectionId", Flag: "vpc-connection-id", Type: "*string", Required: true},
}

var fields_generate_embed_url_for_anonymous_user = []leanruntime.Field{
	{Name: "AllowedDomains", Flag: "allowed-domains", Type: "[]string", Required: false},
	{Name: "AuthorizedResourceArns", Flag: "authorized-resource-arns", Type: "[]string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ExperienceConfiguration", Flag: "experience-configuration", Type: "*types.AnonymousUserEmbeddingExperienceConfiguration", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "SessionLifetimeInMinutes", Flag: "session-lifetime-in-minutes", Type: "*int64", Required: false},
	{Name: "SessionTags", Flag: "session-tags", Type: "[]types.SessionTag", Required: false},
}

var fields_generate_embed_url_for_registered_user = []leanruntime.Field{
	{Name: "AllowedDomains", Flag: "allowed-domains", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ExperienceConfiguration", Flag: "experience-configuration", Type: "*types.RegisteredUserEmbeddingExperienceConfiguration", Required: true},
	{Name: "SessionLifetimeInMinutes", Flag: "session-lifetime-in-minutes", Type: "*int64", Required: false},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: true},
}

var fields_generate_embed_url_for_registered_user_with_identity = []leanruntime.Field{
	{Name: "AllowedDomains", Flag: "allowed-domains", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ExperienceConfiguration", Flag: "experience-configuration", Type: "*types.RegisteredUserEmbeddingExperienceConfiguration", Required: true},
	{Name: "SessionLifetimeInMinutes", Flag: "session-lifetime-in-minutes", Type: "*int64", Required: false},
}

var fields_get_dashboard_embed_url = []leanruntime.Field{
	{Name: "AdditionalDashboardIds", Flag: "additional-dashboard-ids", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.EmbeddingIdentityType", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "ResetDisabled", Flag: "reset-disabled", Type: "bool", Required: false},
	{Name: "SessionLifetimeInMinutes", Flag: "session-lifetime-in-minutes", Type: "*int64", Required: false},
	{Name: "StatePersistenceEnabled", Flag: "state-persistence-enabled", Type: "bool", Required: false},
	{Name: "UndoRedoDisabled", Flag: "undo-redo-disabled", Type: "bool", Required: false},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: false},
}

var fields_get_flow_metadata = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FlowId", Flag: "flow-id", Type: "*string", Required: true},
}

var fields_get_flow_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FlowId", Flag: "flow-id", Type: "*string", Required: true},
}

var fields_get_identity_context = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "SessionExpiresAt", Flag: "session-expires-at", Type: "*time.Time", Required: false},
	{Name: "UserIdentifier", Flag: "user-identifier", Type: "types.UserIdentifier", Required: true},
}

var fields_get_session_embed_url = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "EntryPoint", Flag: "entry-point", Type: "*string", Required: false},
	{Name: "SessionLifetimeInMinutes", Flag: "session-lifetime-in-minutes", Type: "*int64", Required: false},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: false},
}

var fields_list_action_connectors = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_analyses = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_bundle_export_jobs = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_bundle_import_jobs = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_brands = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dashboard_versions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dashboards = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_sets = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flows = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_folder_members = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_folders = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_folders_for_resource = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_group_memberships = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_iam_policy_assignments = []leanruntime.Field{
	{Name: "AssignmentStatus", Flag: "assignment-status", Type: "types.AssignmentStatus", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_iam_policy_assignments_for_user = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_identity_propagation_configs = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ingestions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_namespaces = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_refresh_schedules = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_list_role_memberships = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
}

var fields_list_self_upgrades = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_template_aliases = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_list_template_versions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_list_templates = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_theme_aliases = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
}

var fields_list_theme_versions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
}

var fields_list_themes = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ThemeType", Required: false},
}

var fields_list_topic_refresh_schedules = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_list_topic_reviewed_answers = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_list_topics = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_user_groups = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_connections = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_predict_qa_results = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "IncludeGeneratedAnswer", Flag: "include-generated-answer", Type: "types.IncludeGeneratedAnswer", Required: false},
	{Name: "IncludeQuickSightQIndex", Flag: "include-quicksight-q-index", Type: "types.IncludeQuickSightQIndex", Required: false},
	{Name: "MaxTopicsToConsider", Flag: "max-topics-to-consider", Type: "*int32", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
}

var fields_put_data_set_refresh_properties = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "DataSetRefreshProperties", Flag: "data-set-refresh-properties", Type: "*types.DataSetRefreshProperties", Required: true},
}

var fields_register_user = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomFederationProviderUrl", Flag: "custom-federation-provider-url", Type: "*string", Required: false},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "ExternalLoginFederationProviderType", Flag: "external-login-federation-provider-type", Type: "*string", Required: false},
	{Name: "ExternalLoginId", Flag: "external-login-id", Type: "*string", Required: false},
	{Name: "IamArn", Flag: "iam-arn", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "SessionName", Flag: "session-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
	{Name: "UserRole", Flag: "user-role", Type: "types.UserRole", Required: true},
}

var fields_restore_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "RestoreToFolders", Flag: "restore-to-folders", Type: "bool", Required: false},
}

var fields_search_action_connectors = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.ActionConnectorSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_analyses = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.AnalysisSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_dashboards = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.DashboardSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_data_sets = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.DataSetSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_data_sources = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.DataSourceSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_flows = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.SearchFlowsFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_folders = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.FolderSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_groups = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.GroupSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_topics = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.TopicSearchFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_asset_bundle_export_job = []leanruntime.Field{
	{Name: "AssetBundleExportJobId", Flag: "asset-bundle-export-job-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CloudFormationOverridePropertyConfiguration", Flag: "cloud-formation-override-property-configuration", Type: "*types.AssetBundleCloudFormationOverridePropertyConfiguration", Required: false},
	{Name: "ExportFormat", Flag: "export-format", Type: "types.AssetBundleExportFormat", Required: true},
	{Name: "IncludeAllDependencies", Flag: "include-all-dependencies", Type: "bool", Required: false},
	{Name: "IncludeFolderMembers", Flag: "include-folder-members", Type: "types.IncludeFolderMembers", Required: false},
	{Name: "IncludeFolderMemberships", Flag: "include-folder-memberships", Type: "bool", Required: false},
	{Name: "IncludePermissions", Flag: "include-permissions", Type: "bool", Required: false},
	{Name: "IncludeTags", Flag: "include-tags", Type: "bool", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.AssetBundleExportJobValidationStrategy", Required: false},
}

var fields_start_asset_bundle_import_job = []leanruntime.Field{
	{Name: "AssetBundleImportJobId", Flag: "asset-bundle-import-job-id", Type: "*string", Required: true},
	{Name: "AssetBundleImportSource", Flag: "asset-bundle-import-source", Type: "*types.AssetBundleImportSource", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FailureAction", Flag: "failure-action", Type: "types.AssetBundleImportFailureAction", Required: false},
	{Name: "OverrideParameters", Flag: "override-parameters", Type: "*types.AssetBundleImportJobOverrideParameters", Required: false},
	{Name: "OverridePermissions", Flag: "override-permissions", Type: "*types.AssetBundleImportJobOverridePermissions", Required: false},
	{Name: "OverrideTags", Flag: "override-tags", Type: "*types.AssetBundleImportJobOverrideTags", Required: false},
	{Name: "OverrideValidationStrategy", Flag: "override-validation-strategy", Type: "*types.AssetBundleImportJobOverrideValidationStrategy", Required: false},
}

var fields_start_dashboard_snapshot_job = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "SnapshotConfiguration", Flag: "snapshot-configuration", Type: "*types.SnapshotConfiguration", Required: true},
	{Name: "SnapshotJobId", Flag: "snapshot-job-id", Type: "*string", Required: true},
	{Name: "UserConfiguration", Flag: "user-configuration", Type: "*types.SnapshotUserConfiguration", Required: false},
}

var fields_start_dashboard_snapshot_job_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "ScheduleId", Flag: "schedule-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
}

var fields_update_account_customization = []leanruntime.Field{
	{Name: "AccountCustomization", Flag: "account-customization", Type: "*types.AccountCustomization", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DefaultNamespace", Flag: "default-namespace", Type: "*string", Required: true},
	{Name: "NotificationEmail", Flag: "notification-email", Type: "*string", Required: false},
	{Name: "TerminationProtectionEnabled", Flag: "termination-protection-enabled", Type: "bool", Required: false},
}

var fields_update_action_connector = []leanruntime.Field{
	{Name: "ActionConnectorId", Flag: "action-connector-id", Type: "*string", Required: true},
	{Name: "AuthenticationConfig", Flag: "authentication-config", Type: "*types.AuthConfig", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VpcConnectionArn", Flag: "vpc-connection-arn", Type: "*string", Required: false},
}

var fields_update_action_connector_permissions = []leanruntime.Field{
	{Name: "ActionConnectorId", Flag: "action-connector-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
}

var fields_update_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.AnalysisDefinition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*types.Parameters", Required: false},
	{Name: "SourceEntity", Flag: "source-entity", Type: "*types.AnalysisSourceEntity", Required: false},
	{Name: "ThemeArn", Flag: "theme-arn", Type: "*string", Required: false},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.ValidationStrategy", Required: false},
}

var fields_update_analysis_permissions = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
}

var fields_update_application_with_token_exchange_grant = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_update_brand = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandDefinition", Flag: "brand-definition", Type: "*types.BrandDefinition", Required: false},
	{Name: "BrandId", Flag: "brand-id", Type: "*string", Required: true},
}

var fields_update_brand_assignment = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandArn", Flag: "brand-arn", Type: "*string", Required: true},
}

var fields_update_brand_published_version = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BrandId", Flag: "brand-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_update_custom_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Capabilities", Flag: "capabilities", Type: "*types.Capabilities", Required: false},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
}

var fields_update_dashboard = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "DashboardPublishOptions", Flag: "dashboard-publish-options", Type: "*types.DashboardPublishOptions", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.DashboardVersionDefinition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*types.Parameters", Required: false},
	{Name: "SourceEntity", Flag: "source-entity", Type: "*types.DashboardSourceEntity", Required: false},
	{Name: "ThemeArn", Flag: "theme-arn", Type: "*string", Required: false},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.ValidationStrategy", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_update_dashboard_links = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "LinkEntities", Flag: "link-entities", Type: "[]string", Required: true},
}

var fields_update_dashboard_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "GrantLinkPermissions", Flag: "grant-link-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokeLinkPermissions", Flag: "revoke-link-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
}

var fields_update_dashboard_published_version = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_update_dashboards_qa_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DashboardsQAStatus", Flag: "dashboards-qa-status", Type: "types.DashboardsQAStatus", Required: true},
}

var fields_update_data_set = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ColumnGroups", Flag: "column-groups", Type: "[]types.ColumnGroup", Required: false},
	{Name: "ColumnLevelPermissionRules", Flag: "column-level-permission-rules", Type: "[]types.ColumnLevelPermissionRule", Required: false},
	{Name: "DataPrepConfiguration", Flag: "data-prep-configuration", Type: "*types.DataPrepConfiguration", Required: false},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "DataSetUsageConfiguration", Flag: "data-set-usage-configuration", Type: "*types.DataSetUsageConfiguration", Required: false},
	{Name: "DatasetParameters", Flag: "dataset-parameters", Type: "[]types.DatasetParameter", Required: false},
	{Name: "FieldFolders", Flag: "field-folders", Type: "map[string]types.FieldFolder", Required: false},
	{Name: "ImportMode", Flag: "import-mode", Type: "types.DataSetImportMode", Required: true},
	{Name: "LogicalTableMap", Flag: "logical-table-map", Type: "map[string]types.LogicalTable", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PerformanceConfiguration", Flag: "performance-configuration", Type: "*types.PerformanceConfiguration", Required: false},
	{Name: "PhysicalTableMap", Flag: "physical-table-map", Type: "map[string]types.PhysicalTable", Required: true},
	{Name: "RowLevelPermissionDataSet", Flag: "row-level-permission-data-set", Type: "*types.RowLevelPermissionDataSet", Required: false},
	{Name: "RowLevelPermissionTagConfiguration", Flag: "row-level-permission-tag-configuration", Type: "*types.RowLevelPermissionTagConfiguration", Required: false},
	{Name: "SemanticModelConfiguration", Flag: "semantic-model-configuration", Type: "*types.SemanticModelConfiguration", Required: false},
}

var fields_update_data_set_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Credentials", Flag: "credentials", Type: "*types.DataSourceCredentials", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DataSourceParameters", Flag: "data-source-parameters", Type: "types.DataSourceParameters", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SslProperties", Flag: "ssl-properties", Type: "*types.SslProperties", Required: false},
	{Name: "VpcConnectionProperties", Flag: "vpc-connection-properties", Type: "*types.VpcConnectionProperties", Required: false},
}

var fields_update_data_source_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
}

var fields_update_default_qbusiness_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
}

var fields_update_flow_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FlowId", Flag: "flow-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.Permission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.Permission", Required: false},
}

var fields_update_folder = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_folder_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
}

var fields_update_group = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_update_iam_policy_assignment = []leanruntime.Field{
	{Name: "AssignmentName", Flag: "assignment-name", Type: "*string", Required: true},
	{Name: "AssignmentStatus", Flag: "assignment-status", Type: "types.AssignmentStatus", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Identities", Flag: "identities", Type: "map[string][]string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: false},
}

var fields_update_identity_propagation_config = []leanruntime.Field{
	{Name: "AuthorizedTargets", Flag: "authorized-targets", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Service", Flag: "service", Type: "types.ServiceType", Required: true},
}

var fields_update_ip_restriction = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "IpRestrictionRuleMap", Flag: "ip-restriction-rule-map", Type: "map[string]string", Required: false},
	{Name: "VpcEndpointIdRestrictionRuleMap", Flag: "vpc-endpoint-id-restriction-rule-map", Type: "map[string]string", Required: false},
	{Name: "VpcIdRestrictionRuleMap", Flag: "vpc-id-restriction-rule-map", Type: "map[string]string", Required: false},
}

var fields_update_key_registration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "KeyRegistration", Flag: "key-registration", Type: "[]types.RegisteredCustomerManagedKey", Required: true},
}

var fields_update_public_sharing_settings = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "PublicSharingEnabled", Flag: "public-sharing-enabled", Type: "bool", Required: false},
}

var fields_update_qpersonalization_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "PersonalizationMode", Flag: "personalization-mode", Type: "types.PersonalizationMode", Required: true},
}

var fields_update_quicksight_qsearch_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "QSearchStatus", Flag: "qsearch-status", Type: "types.QSearchStatus", Required: true},
}

var fields_update_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.RefreshSchedule", Required: true},
}

var fields_update_role_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
}

var fields_update_self_upgrade = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.SelfUpgradeAdminAction", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "UpgradeRequestId", Flag: "upgrade-request-id", Type: "*string", Required: true},
}

var fields_update_self_upgrade_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "SelfUpgradeStatus", Flag: "self-upgrade-status", Type: "types.SelfUpgradeStatus", Required: true},
}

var fields_update_spice_capacity_configuration = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "PurchaseMode", Flag: "purchase-mode", Type: "types.PurchaseMode", Required: true},
}

var fields_update_template = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.TemplateVersionDefinition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SourceEntity", Flag: "source-entity", Type: "*types.TemplateSourceEntity", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "ValidationStrategy", Flag: "validation-strategy", Type: "*types.ValidationStrategy", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_update_template_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
	{Name: "TemplateVersionNumber", Flag: "template-version-number", Type: "*int64", Required: true},
}

var fields_update_template_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_update_theme = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "BaseThemeId", Flag: "base-theme-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ThemeConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_update_theme_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
	{Name: "ThemeVersionNumber", Flag: "theme-version-number", Type: "*int64", Required: true},
}

var fields_update_theme_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "ThemeId", Flag: "theme-id", Type: "*string", Required: true},
}

var fields_update_topic = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomInstructions", Flag: "custom-instructions", Type: "*types.CustomInstructions", Required: false},
	{Name: "Topic", Flag: "topic", Type: "*types.TopicDetails", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_update_topic_permissions = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "GrantPermissions", Flag: "grant-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "RevokePermissions", Flag: "revoke-permissions", Type: "[]types.ResourcePermission", Required: false},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_update_topic_refresh_schedule = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "RefreshSchedule", Flag: "refresh-schedule", Type: "*types.TopicRefreshSchedule", Required: true},
	{Name: "TopicId", Flag: "topic-id", Type: "*string", Required: true},
}

var fields_update_user = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomFederationProviderUrl", Flag: "custom-federation-provider-url", Type: "*string", Required: false},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "ExternalLoginFederationProviderType", Flag: "external-login-federation-provider-type", Type: "*string", Required: false},
	{Name: "ExternalLoginId", Flag: "external-login-id", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.UserRole", Required: true},
	{Name: "UnapplyCustomPermissions", Flag: "unapply-custom-permissions", Type: "bool", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_user_custom_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "CustomPermissionsName", Flag: "custom-permissions-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_vpc_connection = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "DnsResolvers", Flag: "dns-resolvers", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "VPCConnectionId", Flag: "vpc-connection-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-topic-reviewed-answer": {
			Name:   "batch-create-topic-reviewed-answer",
			Fields: fields_batch_create_topic_reviewed_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateTopicReviewedAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_topic_reviewed_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateTopicReviewedAnswer(ctx, input)
			},
		},
		"batch-delete-topic-reviewed-answer": {
			Name:   "batch-delete-topic-reviewed-answer",
			Fields: fields_batch_delete_topic_reviewed_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteTopicReviewedAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_topic_reviewed_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteTopicReviewedAnswer(ctx, input)
			},
		},
		"cancel-ingestion": {
			Name:   "cancel-ingestion",
			Fields: fields_cancel_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelIngestion(ctx, input)
			},
		},
		"create-account-customization": {
			Name:   "create-account-customization",
			Fields: fields_create_account_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountCustomization(ctx, input)
			},
		},
		"create-account-subscription": {
			Name:   "create-account-subscription",
			Fields: fields_create_account_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountSubscription(ctx, input)
			},
		},
		"create-action-connector": {
			Name:   "create-action-connector",
			Fields: fields_create_action_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateActionConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_action_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateActionConnector(ctx, input)
			},
		},
		"create-analysis": {
			Name:   "create-analysis",
			Fields: fields_create_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnalysis(ctx, input)
			},
		},
		"create-brand": {
			Name:   "create-brand",
			Fields: fields_create_brand,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBrandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_brand, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBrand(ctx, input)
			},
		},
		"create-custom-permissions": {
			Name:   "create-custom-permissions",
			Fields: fields_create_custom_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomPermissions(ctx, input)
			},
		},
		"create-dashboard": {
			Name:   "create-dashboard",
			Fields: fields_create_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDashboard(ctx, input)
			},
		},
		"create-data-set": {
			Name:   "create-data-set",
			Fields: fields_create_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSet(ctx, input)
			},
		},
		"create-data-source": {
			Name:   "create-data-source",
			Fields: fields_create_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSource(ctx, input)
			},
		},
		"create-folder": {
			Name:   "create-folder",
			Fields: fields_create_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFolder(ctx, input)
			},
		},
		"create-folder-membership": {
			Name:   "create-folder-membership",
			Fields: fields_create_folder_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFolderMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_folder_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFolderMembership(ctx, input)
			},
		},
		"create-group": {
			Name:   "create-group",
			Fields: fields_create_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroup(ctx, input)
			},
		},
		"create-group-membership": {
			Name:   "create-group-membership",
			Fields: fields_create_group_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroupMembership(ctx, input)
			},
		},
		"create-iam-policy-assignment": {
			Name:   "create-iam-policy-assignment",
			Fields: fields_create_iam_policy_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIAMPolicyAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_iam_policy_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIAMPolicyAssignment(ctx, input)
			},
		},
		"create-ingestion": {
			Name:   "create-ingestion",
			Fields: fields_create_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIngestion(ctx, input)
			},
		},
		"create-namespace": {
			Name:   "create-namespace",
			Fields: fields_create_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNamespace(ctx, input)
			},
		},
		"create-refresh-schedule": {
			Name:   "create-refresh-schedule",
			Fields: fields_create_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRefreshSchedule(ctx, input)
			},
		},
		"create-role-membership": {
			Name:   "create-role-membership",
			Fields: fields_create_role_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoleMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_role_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoleMembership(ctx, input)
			},
		},
		"create-template": {
			Name:   "create-template",
			Fields: fields_create_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplate(ctx, input)
			},
		},
		"create-template-alias": {
			Name:   "create-template-alias",
			Fields: fields_create_template_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplateAlias(ctx, input)
			},
		},
		"create-theme": {
			Name:   "create-theme",
			Fields: fields_create_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTheme(ctx, input)
			},
		},
		"create-theme-alias": {
			Name:   "create-theme-alias",
			Fields: fields_create_theme_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThemeAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_theme_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThemeAlias(ctx, input)
			},
		},
		"create-topic": {
			Name:   "create-topic",
			Fields: fields_create_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTopic(ctx, input)
			},
		},
		"create-topic-refresh-schedule": {
			Name:   "create-topic-refresh-schedule",
			Fields: fields_create_topic_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTopicRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_topic_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTopicRefreshSchedule(ctx, input)
			},
		},
		"create-vpc-connection": {
			Name:   "create-vpc-connection",
			Fields: fields_create_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVPCConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVPCConnection(ctx, input)
			},
		},
		"delete-account-custom-permission": {
			Name:   "delete-account-custom-permission",
			Fields: fields_delete_account_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountCustomPermission(ctx, input)
			},
		},
		"delete-account-customization": {
			Name:   "delete-account-customization",
			Fields: fields_delete_account_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountCustomization(ctx, input)
			},
		},
		"delete-account-subscription": {
			Name:   "delete-account-subscription",
			Fields: fields_delete_account_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountSubscription(ctx, input)
			},
		},
		"delete-action-connector": {
			Name:   "delete-action-connector",
			Fields: fields_delete_action_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteActionConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_action_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteActionConnector(ctx, input)
			},
		},
		"delete-analysis": {
			Name:   "delete-analysis",
			Fields: fields_delete_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnalysis(ctx, input)
			},
		},
		"delete-brand": {
			Name:   "delete-brand",
			Fields: fields_delete_brand,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBrandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_brand, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBrand(ctx, input)
			},
		},
		"delete-brand-assignment": {
			Name:   "delete-brand-assignment",
			Fields: fields_delete_brand_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBrandAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_brand_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBrandAssignment(ctx, input)
			},
		},
		"delete-custom-permissions": {
			Name:   "delete-custom-permissions",
			Fields: fields_delete_custom_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomPermissions(ctx, input)
			},
		},
		"delete-dashboard": {
			Name:   "delete-dashboard",
			Fields: fields_delete_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDashboard(ctx, input)
			},
		},
		"delete-data-set": {
			Name:   "delete-data-set",
			Fields: fields_delete_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSet(ctx, input)
			},
		},
		"delete-data-set-refresh-properties": {
			Name:   "delete-data-set-refresh-properties",
			Fields: fields_delete_data_set_refresh_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSetRefreshPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_set_refresh_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSetRefreshProperties(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
			},
		},
		"delete-default-qbusiness-application": {
			Name:   "delete-default-qbusiness-application",
			Fields: fields_delete_default_qbusiness_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDefaultQBusinessApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_default_qbusiness_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDefaultQBusinessApplication(ctx, input)
			},
		},
		"delete-folder": {
			Name:   "delete-folder",
			Fields: fields_delete_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFolder(ctx, input)
			},
		},
		"delete-folder-membership": {
			Name:   "delete-folder-membership",
			Fields: fields_delete_folder_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFolderMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_folder_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFolderMembership(ctx, input)
			},
		},
		"delete-group": {
			Name:   "delete-group",
			Fields: fields_delete_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroup(ctx, input)
			},
		},
		"delete-group-membership": {
			Name:   "delete-group-membership",
			Fields: fields_delete_group_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroupMembership(ctx, input)
			},
		},
		"delete-iam-policy-assignment": {
			Name:   "delete-iam-policy-assignment",
			Fields: fields_delete_iam_policy_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIAMPolicyAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_iam_policy_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIAMPolicyAssignment(ctx, input)
			},
		},
		"delete-identity-propagation-config": {
			Name:   "delete-identity-propagation-config",
			Fields: fields_delete_identity_propagation_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityPropagationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_propagation_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentityPropagationConfig(ctx, input)
			},
		},
		"delete-namespace": {
			Name:   "delete-namespace",
			Fields: fields_delete_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNamespace(ctx, input)
			},
		},
		"delete-refresh-schedule": {
			Name:   "delete-refresh-schedule",
			Fields: fields_delete_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRefreshSchedule(ctx, input)
			},
		},
		"delete-role-custom-permission": {
			Name:   "delete-role-custom-permission",
			Fields: fields_delete_role_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoleCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_role_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoleCustomPermission(ctx, input)
			},
		},
		"delete-role-membership": {
			Name:   "delete-role-membership",
			Fields: fields_delete_role_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoleMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_role_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoleMembership(ctx, input)
			},
		},
		"delete-template": {
			Name:   "delete-template",
			Fields: fields_delete_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplate(ctx, input)
			},
		},
		"delete-template-alias": {
			Name:   "delete-template-alias",
			Fields: fields_delete_template_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplateAlias(ctx, input)
			},
		},
		"delete-theme": {
			Name:   "delete-theme",
			Fields: fields_delete_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTheme(ctx, input)
			},
		},
		"delete-theme-alias": {
			Name:   "delete-theme-alias",
			Fields: fields_delete_theme_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThemeAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_theme_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThemeAlias(ctx, input)
			},
		},
		"delete-topic": {
			Name:   "delete-topic",
			Fields: fields_delete_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTopic(ctx, input)
			},
		},
		"delete-topic-refresh-schedule": {
			Name:   "delete-topic-refresh-schedule",
			Fields: fields_delete_topic_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTopicRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_topic_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTopicRefreshSchedule(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"delete-user-by-principal-id": {
			Name:   "delete-user-by-principal-id",
			Fields: fields_delete_user_by_principal_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserByPrincipalIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_by_principal_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserByPrincipalId(ctx, input)
			},
		},
		"delete-user-custom-permission": {
			Name:   "delete-user-custom-permission",
			Fields: fields_delete_user_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserCustomPermission(ctx, input)
			},
		},
		"delete-vpc-connection": {
			Name:   "delete-vpc-connection",
			Fields: fields_delete_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVPCConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVPCConnection(ctx, input)
			},
		},
		"describe-account-custom-permission": {
			Name:   "describe-account-custom-permission",
			Fields: fields_describe_account_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountCustomPermission(ctx, input)
			},
		},
		"describe-account-customization": {
			Name:   "describe-account-customization",
			Fields: fields_describe_account_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountCustomization(ctx, input)
			},
		},
		"describe-account-settings": {
			Name:   "describe-account-settings",
			Fields: fields_describe_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountSettings(ctx, input)
			},
		},
		"describe-account-subscription": {
			Name:   "describe-account-subscription",
			Fields: fields_describe_account_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountSubscription(ctx, input)
			},
		},
		"describe-action-connector": {
			Name:   "describe-action-connector",
			Fields: fields_describe_action_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActionConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_action_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeActionConnector(ctx, input)
			},
		},
		"describe-action-connector-permissions": {
			Name:   "describe-action-connector-permissions",
			Fields: fields_describe_action_connector_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActionConnectorPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_action_connector_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeActionConnectorPermissions(ctx, input)
			},
		},
		"describe-analysis": {
			Name:   "describe-analysis",
			Fields: fields_describe_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAnalysis(ctx, input)
			},
		},
		"describe-analysis-definition": {
			Name:   "describe-analysis-definition",
			Fields: fields_describe_analysis_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnalysisDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_analysis_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAnalysisDefinition(ctx, input)
			},
		},
		"describe-analysis-permissions": {
			Name:   "describe-analysis-permissions",
			Fields: fields_describe_analysis_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnalysisPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_analysis_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAnalysisPermissions(ctx, input)
			},
		},
		"describe-asset-bundle-export-job": {
			Name:   "describe-asset-bundle-export-job",
			Fields: fields_describe_asset_bundle_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetBundleExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_bundle_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetBundleExportJob(ctx, input)
			},
		},
		"describe-asset-bundle-import-job": {
			Name:   "describe-asset-bundle-import-job",
			Fields: fields_describe_asset_bundle_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetBundleImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_bundle_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetBundleImportJob(ctx, input)
			},
		},
		"describe-brand": {
			Name:   "describe-brand",
			Fields: fields_describe_brand,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBrandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_brand, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBrand(ctx, input)
			},
		},
		"describe-brand-assignment": {
			Name:   "describe-brand-assignment",
			Fields: fields_describe_brand_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBrandAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_brand_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBrandAssignment(ctx, input)
			},
		},
		"describe-brand-published-version": {
			Name:   "describe-brand-published-version",
			Fields: fields_describe_brand_published_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBrandPublishedVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_brand_published_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBrandPublishedVersion(ctx, input)
			},
		},
		"describe-custom-permissions": {
			Name:   "describe-custom-permissions",
			Fields: fields_describe_custom_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomPermissions(ctx, input)
			},
		},
		"describe-dashboard": {
			Name:   "describe-dashboard",
			Fields: fields_describe_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboard(ctx, input)
			},
		},
		"describe-dashboard-definition": {
			Name:   "describe-dashboard-definition",
			Fields: fields_describe_dashboard_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboard_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboardDefinition(ctx, input)
			},
		},
		"describe-dashboard-permissions": {
			Name:   "describe-dashboard-permissions",
			Fields: fields_describe_dashboard_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboard_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboardPermissions(ctx, input)
			},
		},
		"describe-dashboard-snapshot-job": {
			Name:   "describe-dashboard-snapshot-job",
			Fields: fields_describe_dashboard_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboard_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboardSnapshotJob(ctx, input)
			},
		},
		"describe-dashboard-snapshot-job-result": {
			Name:   "describe-dashboard-snapshot-job-result",
			Fields: fields_describe_dashboard_snapshot_job_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardSnapshotJobResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboard_snapshot_job_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboardSnapshotJobResult(ctx, input)
			},
		},
		"describe-dashboards-qa-configuration": {
			Name:   "describe-dashboards-qa-configuration",
			Fields: fields_describe_dashboards_qa_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardsQAConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboards_qa_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboardsQAConfiguration(ctx, input)
			},
		},
		"describe-data-set": {
			Name:   "describe-data-set",
			Fields: fields_describe_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataSet(ctx, input)
			},
		},
		"describe-data-set-permissions": {
			Name:   "describe-data-set-permissions",
			Fields: fields_describe_data_set_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSetPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_set_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataSetPermissions(ctx, input)
			},
		},
		"describe-data-set-refresh-properties": {
			Name:   "describe-data-set-refresh-properties",
			Fields: fields_describe_data_set_refresh_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSetRefreshPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_set_refresh_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataSetRefreshProperties(ctx, input)
			},
		},
		"describe-data-source": {
			Name:   "describe-data-source",
			Fields: fields_describe_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataSource(ctx, input)
			},
		},
		"describe-data-source-permissions": {
			Name:   "describe-data-source-permissions",
			Fields: fields_describe_data_source_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSourcePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_source_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataSourcePermissions(ctx, input)
			},
		},
		"describe-default-qbusiness-application": {
			Name:   "describe-default-qbusiness-application",
			Fields: fields_describe_default_qbusiness_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDefaultQBusinessApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_default_qbusiness_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDefaultQBusinessApplication(ctx, input)
			},
		},
		"describe-folder": {
			Name:   "describe-folder",
			Fields: fields_describe_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFolder(ctx, input)
			},
		},
		"describe-folder-permissions": {
			Name:   "describe-folder-permissions",
			Fields: fields_describe_folder_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFolderPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_folder_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFolderPermissions(ctx, input)
				}
				var results []*svc.DescribeFolderPermissionsOutput
				p := svc.NewDescribeFolderPermissionsPaginator(client, input)
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
		"describe-folder-resolved-permissions": {
			Name:   "describe-folder-resolved-permissions",
			Fields: fields_describe_folder_resolved_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFolderResolvedPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_folder_resolved_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFolderResolvedPermissions(ctx, input)
				}
				var results []*svc.DescribeFolderResolvedPermissionsOutput
				p := svc.NewDescribeFolderResolvedPermissionsPaginator(client, input)
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
		"describe-group": {
			Name:   "describe-group",
			Fields: fields_describe_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGroup(ctx, input)
			},
		},
		"describe-group-membership": {
			Name:   "describe-group-membership",
			Fields: fields_describe_group_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGroupMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_group_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGroupMembership(ctx, input)
			},
		},
		"describe-iam-policy-assignment": {
			Name:   "describe-iam-policy-assignment",
			Fields: fields_describe_iam_policy_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIAMPolicyAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_iam_policy_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIAMPolicyAssignment(ctx, input)
			},
		},
		"describe-ingestion": {
			Name:   "describe-ingestion",
			Fields: fields_describe_ingestion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIngestionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ingestion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIngestion(ctx, input)
			},
		},
		"describe-ip-restriction": {
			Name:   "describe-ip-restriction",
			Fields: fields_describe_ip_restriction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpRestrictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ip_restriction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIpRestriction(ctx, input)
			},
		},
		"describe-key-registration": {
			Name:   "describe-key-registration",
			Fields: fields_describe_key_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKeyRegistration(ctx, input)
			},
		},
		"describe-namespace": {
			Name:   "describe-namespace",
			Fields: fields_describe_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNamespace(ctx, input)
			},
		},
		"describe-qpersonalization-configuration": {
			Name:   "describe-qpersonalization-configuration",
			Fields: fields_describe_qpersonalization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQPersonalizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_qpersonalization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQPersonalizationConfiguration(ctx, input)
			},
		},
		"describe-quicksight-qsearch-configuration": {
			Name:   "describe-quicksight-qsearch-configuration",
			Fields: fields_describe_quicksight_qsearch_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQuickSightQSearchConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_quicksight_qsearch_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQuickSightQSearchConfiguration(ctx, input)
			},
		},
		"describe-refresh-schedule": {
			Name:   "describe-refresh-schedule",
			Fields: fields_describe_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRefreshSchedule(ctx, input)
			},
		},
		"describe-role-custom-permission": {
			Name:   "describe-role-custom-permission",
			Fields: fields_describe_role_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRoleCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_role_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRoleCustomPermission(ctx, input)
			},
		},
		"describe-self-upgrade-configuration": {
			Name:   "describe-self-upgrade-configuration",
			Fields: fields_describe_self_upgrade_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSelfUpgradeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_self_upgrade_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSelfUpgradeConfiguration(ctx, input)
			},
		},
		"describe-template": {
			Name:   "describe-template",
			Fields: fields_describe_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTemplate(ctx, input)
			},
		},
		"describe-template-alias": {
			Name:   "describe-template-alias",
			Fields: fields_describe_template_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTemplateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_template_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTemplateAlias(ctx, input)
			},
		},
		"describe-template-definition": {
			Name:   "describe-template-definition",
			Fields: fields_describe_template_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTemplateDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_template_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTemplateDefinition(ctx, input)
			},
		},
		"describe-template-permissions": {
			Name:   "describe-template-permissions",
			Fields: fields_describe_template_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTemplatePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_template_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTemplatePermissions(ctx, input)
			},
		},
		"describe-theme": {
			Name:   "describe-theme",
			Fields: fields_describe_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTheme(ctx, input)
			},
		},
		"describe-theme-alias": {
			Name:   "describe-theme-alias",
			Fields: fields_describe_theme_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThemeAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_theme_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThemeAlias(ctx, input)
			},
		},
		"describe-theme-permissions": {
			Name:   "describe-theme-permissions",
			Fields: fields_describe_theme_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThemePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_theme_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThemePermissions(ctx, input)
			},
		},
		"describe-topic": {
			Name:   "describe-topic",
			Fields: fields_describe_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTopic(ctx, input)
			},
		},
		"describe-topic-permissions": {
			Name:   "describe-topic-permissions",
			Fields: fields_describe_topic_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_topic_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTopicPermissions(ctx, input)
			},
		},
		"describe-topic-refresh": {
			Name:   "describe-topic-refresh",
			Fields: fields_describe_topic_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_topic_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTopicRefresh(ctx, input)
			},
		},
		"describe-topic-refresh-schedule": {
			Name:   "describe-topic-refresh-schedule",
			Fields: fields_describe_topic_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_topic_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTopicRefreshSchedule(ctx, input)
			},
		},
		"describe-user": {
			Name:   "describe-user",
			Fields: fields_describe_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUser(ctx, input)
			},
		},
		"describe-vpc-connection": {
			Name:   "describe-vpc-connection",
			Fields: fields_describe_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVPCConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVPCConnection(ctx, input)
			},
		},
		"generate-embed-url-for-anonymous-user": {
			Name:   "generate-embed-url-for-anonymous-user",
			Fields: fields_generate_embed_url_for_anonymous_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateEmbedUrlForAnonymousUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_embed_url_for_anonymous_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateEmbedUrlForAnonymousUser(ctx, input)
			},
		},
		"generate-embed-url-for-registered-user": {
			Name:   "generate-embed-url-for-registered-user",
			Fields: fields_generate_embed_url_for_registered_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateEmbedUrlForRegisteredUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_embed_url_for_registered_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateEmbedUrlForRegisteredUser(ctx, input)
			},
		},
		"generate-embed-url-for-registered-user-with-identity": {
			Name:   "generate-embed-url-for-registered-user-with-identity",
			Fields: fields_generate_embed_url_for_registered_user_with_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateEmbedUrlForRegisteredUserWithIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_embed_url_for_registered_user_with_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateEmbedUrlForRegisteredUserWithIdentity(ctx, input)
			},
		},
		"get-dashboard-embed-url": {
			Name:   "get-dashboard-embed-url",
			Fields: fields_get_dashboard_embed_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDashboardEmbedUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dashboard_embed_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDashboardEmbedUrl(ctx, input)
			},
		},
		"get-flow-metadata": {
			Name:   "get-flow-metadata",
			Fields: fields_get_flow_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowMetadata(ctx, input)
			},
		},
		"get-flow-permissions": {
			Name:   "get-flow-permissions",
			Fields: fields_get_flow_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowPermissions(ctx, input)
			},
		},
		"get-identity-context": {
			Name:   "get-identity-context",
			Fields: fields_get_identity_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityContext(ctx, input)
			},
		},
		"get-session-embed-url": {
			Name:   "get-session-embed-url",
			Fields: fields_get_session_embed_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionEmbedUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session_embed_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSessionEmbedUrl(ctx, input)
			},
		},
		"list-action-connectors": {
			Name:   "list-action-connectors",
			Fields: fields_list_action_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActionConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_action_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActionConnectors(ctx, input)
				}
				var results []*svc.ListActionConnectorsOutput
				p := svc.NewListActionConnectorsPaginator(client, input)
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
		"list-analyses": {
			Name:   "list-analyses",
			Fields: fields_list_analyses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalysesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_analyses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnalyses(ctx, input)
				}
				var results []*svc.ListAnalysesOutput
				p := svc.NewListAnalysesPaginator(client, input)
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
		"list-asset-bundle-export-jobs": {
			Name:   "list-asset-bundle-export-jobs",
			Fields: fields_list_asset_bundle_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetBundleExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_bundle_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetBundleExportJobs(ctx, input)
				}
				var results []*svc.ListAssetBundleExportJobsOutput
				p := svc.NewListAssetBundleExportJobsPaginator(client, input)
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
		"list-asset-bundle-import-jobs": {
			Name:   "list-asset-bundle-import-jobs",
			Fields: fields_list_asset_bundle_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetBundleImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_bundle_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetBundleImportJobs(ctx, input)
				}
				var results []*svc.ListAssetBundleImportJobsOutput
				p := svc.NewListAssetBundleImportJobsPaginator(client, input)
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
		"list-brands": {
			Name:   "list-brands",
			Fields: fields_list_brands,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBrandsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_brands, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBrands(ctx, input)
				}
				var results []*svc.ListBrandsOutput
				p := svc.NewListBrandsPaginator(client, input)
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
		"list-custom-permissions": {
			Name:   "list-custom-permissions",
			Fields: fields_list_custom_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomPermissions(ctx, input)
				}
				var results []*svc.ListCustomPermissionsOutput
				p := svc.NewListCustomPermissionsPaginator(client, input)
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
		"list-dashboard-versions": {
			Name:   "list-dashboard-versions",
			Fields: fields_list_dashboard_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDashboardVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dashboard_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDashboardVersions(ctx, input)
				}
				var results []*svc.ListDashboardVersionsOutput
				p := svc.NewListDashboardVersionsPaginator(client, input)
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
		"list-dashboards": {
			Name:   "list-dashboards",
			Fields: fields_list_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDashboardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dashboards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDashboards(ctx, input)
				}
				var results []*svc.ListDashboardsOutput
				p := svc.NewListDashboardsPaginator(client, input)
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
		"list-data-sets": {
			Name:   "list-data-sets",
			Fields: fields_list_data_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSets(ctx, input)
				}
				var results []*svc.ListDataSetsOutput
				p := svc.NewListDataSetsPaginator(client, input)
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
		"list-data-sources": {
			Name:   "list-data-sources",
			Fields: fields_list_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSources(ctx, input)
				}
				var results []*svc.ListDataSourcesOutput
				p := svc.NewListDataSourcesPaginator(client, input)
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
		"list-flows": {
			Name:   "list-flows",
			Fields: fields_list_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlows(ctx, input)
				}
				var results []*svc.ListFlowsOutput
				p := svc.NewListFlowsPaginator(client, input)
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
		"list-folder-members": {
			Name:   "list-folder-members",
			Fields: fields_list_folder_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFolderMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_folder_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFolderMembers(ctx, input)
				}
				var results []*svc.ListFolderMembersOutput
				p := svc.NewListFolderMembersPaginator(client, input)
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
		"list-folders": {
			Name:   "list-folders",
			Fields: fields_list_folders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFoldersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_folders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFolders(ctx, input)
				}
				var results []*svc.ListFoldersOutput
				p := svc.NewListFoldersPaginator(client, input)
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
		"list-folders-for-resource": {
			Name:   "list-folders-for-resource",
			Fields: fields_list_folders_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFoldersForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_folders_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFoldersForResource(ctx, input)
				}
				var results []*svc.ListFoldersForResourceOutput
				p := svc.NewListFoldersForResourcePaginator(client, input)
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
		"list-group-memberships": {
			Name:   "list-group-memberships",
			Fields: fields_list_group_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupMemberships(ctx, input)
				}
				var results []*svc.ListGroupMembershipsOutput
				p := svc.NewListGroupMembershipsPaginator(client, input)
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
		"list-groups": {
			Name:   "list-groups",
			Fields: fields_list_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroups(ctx, input)
				}
				var results []*svc.ListGroupsOutput
				p := svc.NewListGroupsPaginator(client, input)
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
		"list-iam-policy-assignments": {
			Name:   "list-iam-policy-assignments",
			Fields: fields_list_iam_policy_assignments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIAMPolicyAssignmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_iam_policy_assignments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIAMPolicyAssignments(ctx, input)
				}
				var results []*svc.ListIAMPolicyAssignmentsOutput
				p := svc.NewListIAMPolicyAssignmentsPaginator(client, input)
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
		"list-iam-policy-assignments-for-user": {
			Name:   "list-iam-policy-assignments-for-user",
			Fields: fields_list_iam_policy_assignments_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIAMPolicyAssignmentsForUserInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_iam_policy_assignments_for_user, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIAMPolicyAssignmentsForUser(ctx, input)
				}
				var results []*svc.ListIAMPolicyAssignmentsForUserOutput
				p := svc.NewListIAMPolicyAssignmentsForUserPaginator(client, input)
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
		"list-identity-propagation-configs": {
			Name:   "list-identity-propagation-configs",
			Fields: fields_list_identity_propagation_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityPropagationConfigsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_identity_propagation_configs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIdentityPropagationConfigs(ctx, input)
			},
		},
		"list-ingestions": {
			Name:   "list-ingestions",
			Fields: fields_list_ingestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIngestionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ingestions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIngestions(ctx, input)
				}
				var results []*svc.ListIngestionsOutput
				p := svc.NewListIngestionsPaginator(client, input)
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
		"list-namespaces": {
			Name:   "list-namespaces",
			Fields: fields_list_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNamespaces(ctx, input)
				}
				var results []*svc.ListNamespacesOutput
				p := svc.NewListNamespacesPaginator(client, input)
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
		"list-refresh-schedules": {
			Name:   "list-refresh-schedules",
			Fields: fields_list_refresh_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRefreshSchedulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_refresh_schedules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRefreshSchedules(ctx, input)
			},
		},
		"list-role-memberships": {
			Name:   "list-role-memberships",
			Fields: fields_list_role_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoleMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_role_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoleMemberships(ctx, input)
				}
				var results []*svc.ListRoleMembershipsOutput
				p := svc.NewListRoleMembershipsPaginator(client, input)
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
		"list-self-upgrades": {
			Name:   "list-self-upgrades",
			Fields: fields_list_self_upgrades,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSelfUpgradesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_self_upgrades, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSelfUpgrades(ctx, input)
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
		"list-template-aliases": {
			Name:   "list-template-aliases",
			Fields: fields_list_template_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateAliases(ctx, input)
				}
				var results []*svc.ListTemplateAliasesOutput
				p := svc.NewListTemplateAliasesPaginator(client, input)
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
		"list-template-versions": {
			Name:   "list-template-versions",
			Fields: fields_list_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateVersions(ctx, input)
				}
				var results []*svc.ListTemplateVersionsOutput
				p := svc.NewListTemplateVersionsPaginator(client, input)
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
		"list-templates": {
			Name:   "list-templates",
			Fields: fields_list_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplates(ctx, input)
				}
				var results []*svc.ListTemplatesOutput
				p := svc.NewListTemplatesPaginator(client, input)
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
		"list-theme-aliases": {
			Name:   "list-theme-aliases",
			Fields: fields_list_theme_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThemeAliasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_theme_aliases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListThemeAliases(ctx, input)
			},
		},
		"list-theme-versions": {
			Name:   "list-theme-versions",
			Fields: fields_list_theme_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThemeVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_theme_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThemeVersions(ctx, input)
				}
				var results []*svc.ListThemeVersionsOutput
				p := svc.NewListThemeVersionsPaginator(client, input)
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
		"list-themes": {
			Name:   "list-themes",
			Fields: fields_list_themes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThemesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_themes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThemes(ctx, input)
				}
				var results []*svc.ListThemesOutput
				p := svc.NewListThemesPaginator(client, input)
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
		"list-topic-refresh-schedules": {
			Name:   "list-topic-refresh-schedules",
			Fields: fields_list_topic_refresh_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicRefreshSchedulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_topic_refresh_schedules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTopicRefreshSchedules(ctx, input)
			},
		},
		"list-topic-reviewed-answers": {
			Name:   "list-topic-reviewed-answers",
			Fields: fields_list_topic_reviewed_answers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicReviewedAnswersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_topic_reviewed_answers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTopicReviewedAnswers(ctx, input)
			},
		},
		"list-topics": {
			Name:   "list-topics",
			Fields: fields_list_topics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_topics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTopics(ctx, input)
				}
				var results []*svc.ListTopicsOutput
				p := svc.NewListTopicsPaginator(client, input)
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
		"list-user-groups": {
			Name:   "list-user-groups",
			Fields: fields_list_user_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserGroups(ctx, input)
				}
				var results []*svc.ListUserGroupsOutput
				p := svc.NewListUserGroupsPaginator(client, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"list-vpc-connections": {
			Name:   "list-vpc-connections",
			Fields: fields_list_vpc_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVPCConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vpc_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVPCConnections(ctx, input)
				}
				var results []*svc.ListVPCConnectionsOutput
				p := svc.NewListVPCConnectionsPaginator(client, input)
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
		"predict-qa-results": {
			Name:   "predict-qa-results",
			Fields: fields_predict_qa_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PredictQAResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_predict_qa_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PredictQAResults(ctx, input)
			},
		},
		"put-data-set-refresh-properties": {
			Name:   "put-data-set-refresh-properties",
			Fields: fields_put_data_set_refresh_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataSetRefreshPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_set_refresh_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataSetRefreshProperties(ctx, input)
			},
		},
		"register-user": {
			Name:   "register-user",
			Fields: fields_register_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterUser(ctx, input)
			},
		},
		"restore-analysis": {
			Name:   "restore-analysis",
			Fields: fields_restore_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreAnalysis(ctx, input)
			},
		},
		"search-action-connectors": {
			Name:   "search-action-connectors",
			Fields: fields_search_action_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchActionConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_action_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchActionConnectors(ctx, input)
				}
				var results []*svc.SearchActionConnectorsOutput
				p := svc.NewSearchActionConnectorsPaginator(client, input)
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
		"search-analyses": {
			Name:   "search-analyses",
			Fields: fields_search_analyses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAnalysesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_analyses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchAnalyses(ctx, input)
				}
				var results []*svc.SearchAnalysesOutput
				p := svc.NewSearchAnalysesPaginator(client, input)
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
		"search-dashboards": {
			Name:   "search-dashboards",
			Fields: fields_search_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchDashboardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_dashboards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchDashboards(ctx, input)
				}
				var results []*svc.SearchDashboardsOutput
				p := svc.NewSearchDashboardsPaginator(client, input)
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
		"search-data-sets": {
			Name:   "search-data-sets",
			Fields: fields_search_data_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchDataSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_data_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchDataSets(ctx, input)
				}
				var results []*svc.SearchDataSetsOutput
				p := svc.NewSearchDataSetsPaginator(client, input)
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
		"search-data-sources": {
			Name:   "search-data-sources",
			Fields: fields_search_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchDataSources(ctx, input)
				}
				var results []*svc.SearchDataSourcesOutput
				p := svc.NewSearchDataSourcesPaginator(client, input)
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
		"search-flows": {
			Name:   "search-flows",
			Fields: fields_search_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchFlows(ctx, input)
				}
				var results []*svc.SearchFlowsOutput
				p := svc.NewSearchFlowsPaginator(client, input)
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
		"search-folders": {
			Name:   "search-folders",
			Fields: fields_search_folders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchFoldersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_folders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchFolders(ctx, input)
				}
				var results []*svc.SearchFoldersOutput
				p := svc.NewSearchFoldersPaginator(client, input)
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
		"search-groups": {
			Name:   "search-groups",
			Fields: fields_search_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchGroups(ctx, input)
				}
				var results []*svc.SearchGroupsOutput
				p := svc.NewSearchGroupsPaginator(client, input)
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
		"search-topics": {
			Name:   "search-topics",
			Fields: fields_search_topics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTopicsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_topics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTopics(ctx, input)
				}
				var results []*svc.SearchTopicsOutput
				p := svc.NewSearchTopicsPaginator(client, input)
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
		"start-asset-bundle-export-job": {
			Name:   "start-asset-bundle-export-job",
			Fields: fields_start_asset_bundle_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAssetBundleExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_asset_bundle_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAssetBundleExportJob(ctx, input)
			},
		},
		"start-asset-bundle-import-job": {
			Name:   "start-asset-bundle-import-job",
			Fields: fields_start_asset_bundle_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAssetBundleImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_asset_bundle_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAssetBundleImportJob(ctx, input)
			},
		},
		"start-dashboard-snapshot-job": {
			Name:   "start-dashboard-snapshot-job",
			Fields: fields_start_dashboard_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDashboardSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dashboard_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDashboardSnapshotJob(ctx, input)
			},
		},
		"start-dashboard-snapshot-job-schedule": {
			Name:   "start-dashboard-snapshot-job-schedule",
			Fields: fields_start_dashboard_snapshot_job_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDashboardSnapshotJobScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dashboard_snapshot_job_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDashboardSnapshotJobSchedule(ctx, input)
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
		"update-account-custom-permission": {
			Name:   "update-account-custom-permission",
			Fields: fields_update_account_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountCustomPermission(ctx, input)
			},
		},
		"update-account-customization": {
			Name:   "update-account-customization",
			Fields: fields_update_account_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountCustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountCustomization(ctx, input)
			},
		},
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-action-connector": {
			Name:   "update-action-connector",
			Fields: fields_update_action_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActionConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_action_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateActionConnector(ctx, input)
			},
		},
		"update-action-connector-permissions": {
			Name:   "update-action-connector-permissions",
			Fields: fields_update_action_connector_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActionConnectorPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_action_connector_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateActionConnectorPermissions(ctx, input)
			},
		},
		"update-analysis": {
			Name:   "update-analysis",
			Fields: fields_update_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnalysis(ctx, input)
			},
		},
		"update-analysis-permissions": {
			Name:   "update-analysis-permissions",
			Fields: fields_update_analysis_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnalysisPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_analysis_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnalysisPermissions(ctx, input)
			},
		},
		"update-application-with-token-exchange-grant": {
			Name:   "update-application-with-token-exchange-grant",
			Fields: fields_update_application_with_token_exchange_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationWithTokenExchangeGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_with_token_exchange_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationWithTokenExchangeGrant(ctx, input)
			},
		},
		"update-brand": {
			Name:   "update-brand",
			Fields: fields_update_brand,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_brand, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrand(ctx, input)
			},
		},
		"update-brand-assignment": {
			Name:   "update-brand-assignment",
			Fields: fields_update_brand_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrandAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_brand_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrandAssignment(ctx, input)
			},
		},
		"update-brand-published-version": {
			Name:   "update-brand-published-version",
			Fields: fields_update_brand_published_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrandPublishedVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_brand_published_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrandPublishedVersion(ctx, input)
			},
		},
		"update-custom-permissions": {
			Name:   "update-custom-permissions",
			Fields: fields_update_custom_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomPermissions(ctx, input)
			},
		},
		"update-dashboard": {
			Name:   "update-dashboard",
			Fields: fields_update_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboard(ctx, input)
			},
		},
		"update-dashboard-links": {
			Name:   "update-dashboard-links",
			Fields: fields_update_dashboard_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardLinksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard_links, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboardLinks(ctx, input)
			},
		},
		"update-dashboard-permissions": {
			Name:   "update-dashboard-permissions",
			Fields: fields_update_dashboard_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboardPermissions(ctx, input)
			},
		},
		"update-dashboard-published-version": {
			Name:   "update-dashboard-published-version",
			Fields: fields_update_dashboard_published_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardPublishedVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard_published_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboardPublishedVersion(ctx, input)
			},
		},
		"update-dashboards-qa-configuration": {
			Name:   "update-dashboards-qa-configuration",
			Fields: fields_update_dashboards_qa_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardsQAConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboards_qa_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboardsQAConfiguration(ctx, input)
			},
		},
		"update-data-set": {
			Name:   "update-data-set",
			Fields: fields_update_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSet(ctx, input)
			},
		},
		"update-data-set-permissions": {
			Name:   "update-data-set-permissions",
			Fields: fields_update_data_set_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSetPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_set_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSetPermissions(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
			},
		},
		"update-data-source-permissions": {
			Name:   "update-data-source-permissions",
			Fields: fields_update_data_source_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourcePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSourcePermissions(ctx, input)
			},
		},
		"update-default-qbusiness-application": {
			Name:   "update-default-qbusiness-application",
			Fields: fields_update_default_qbusiness_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDefaultQBusinessApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_default_qbusiness_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDefaultQBusinessApplication(ctx, input)
			},
		},
		"update-flow-permissions": {
			Name:   "update-flow-permissions",
			Fields: fields_update_flow_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowPermissions(ctx, input)
			},
		},
		"update-folder": {
			Name:   "update-folder",
			Fields: fields_update_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFolder(ctx, input)
			},
		},
		"update-folder-permissions": {
			Name:   "update-folder-permissions",
			Fields: fields_update_folder_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFolderPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_folder_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFolderPermissions(ctx, input)
			},
		},
		"update-group": {
			Name:   "update-group",
			Fields: fields_update_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroup(ctx, input)
			},
		},
		"update-iam-policy-assignment": {
			Name:   "update-iam-policy-assignment",
			Fields: fields_update_iam_policy_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIAMPolicyAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_iam_policy_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIAMPolicyAssignment(ctx, input)
			},
		},
		"update-identity-propagation-config": {
			Name:   "update-identity-propagation-config",
			Fields: fields_update_identity_propagation_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdentityPropagationConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_identity_propagation_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdentityPropagationConfig(ctx, input)
			},
		},
		"update-ip-restriction": {
			Name:   "update-ip-restriction",
			Fields: fields_update_ip_restriction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIpRestrictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ip_restriction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIpRestriction(ctx, input)
			},
		},
		"update-key-registration": {
			Name:   "update-key-registration",
			Fields: fields_update_key_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeyRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_key_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKeyRegistration(ctx, input)
			},
		},
		"update-public-sharing-settings": {
			Name:   "update-public-sharing-settings",
			Fields: fields_update_public_sharing_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePublicSharingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_public_sharing_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePublicSharingSettings(ctx, input)
			},
		},
		"update-qpersonalization-configuration": {
			Name:   "update-qpersonalization-configuration",
			Fields: fields_update_qpersonalization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQPersonalizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_qpersonalization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQPersonalizationConfiguration(ctx, input)
			},
		},
		"update-quicksight-qsearch-configuration": {
			Name:   "update-quicksight-qsearch-configuration",
			Fields: fields_update_quicksight_qsearch_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQuickSightQSearchConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_quicksight_qsearch_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQuickSightQSearchConfiguration(ctx, input)
			},
		},
		"update-refresh-schedule": {
			Name:   "update-refresh-schedule",
			Fields: fields_update_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRefreshSchedule(ctx, input)
			},
		},
		"update-role-custom-permission": {
			Name:   "update-role-custom-permission",
			Fields: fields_update_role_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoleCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_role_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoleCustomPermission(ctx, input)
			},
		},
		"update-self-upgrade": {
			Name:   "update-self-upgrade",
			Fields: fields_update_self_upgrade,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSelfUpgradeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_self_upgrade, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSelfUpgrade(ctx, input)
			},
		},
		"update-self-upgrade-configuration": {
			Name:   "update-self-upgrade-configuration",
			Fields: fields_update_self_upgrade_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSelfUpgradeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_self_upgrade_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSelfUpgradeConfiguration(ctx, input)
			},
		},
		"update-spice-capacity-configuration": {
			Name:   "update-spice-capacity-configuration",
			Fields: fields_update_spice_capacity_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSPICECapacityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_spice_capacity_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSPICECapacityConfiguration(ctx, input)
			},
		},
		"update-template": {
			Name:   "update-template",
			Fields: fields_update_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplate(ctx, input)
			},
		},
		"update-template-alias": {
			Name:   "update-template-alias",
			Fields: fields_update_template_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplateAlias(ctx, input)
			},
		},
		"update-template-permissions": {
			Name:   "update-template-permissions",
			Fields: fields_update_template_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplatePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplatePermissions(ctx, input)
			},
		},
		"update-theme": {
			Name:   "update-theme",
			Fields: fields_update_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTheme(ctx, input)
			},
		},
		"update-theme-alias": {
			Name:   "update-theme-alias",
			Fields: fields_update_theme_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThemeAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_theme_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThemeAlias(ctx, input)
			},
		},
		"update-theme-permissions": {
			Name:   "update-theme-permissions",
			Fields: fields_update_theme_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThemePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_theme_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThemePermissions(ctx, input)
			},
		},
		"update-topic": {
			Name:   "update-topic",
			Fields: fields_update_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTopic(ctx, input)
			},
		},
		"update-topic-permissions": {
			Name:   "update-topic-permissions",
			Fields: fields_update_topic_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTopicPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_topic_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTopicPermissions(ctx, input)
			},
		},
		"update-topic-refresh-schedule": {
			Name:   "update-topic-refresh-schedule",
			Fields: fields_update_topic_refresh_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTopicRefreshScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_topic_refresh_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTopicRefreshSchedule(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
		"update-user-custom-permission": {
			Name:   "update-user-custom-permission",
			Fields: fields_update_user_custom_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserCustomPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_custom_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserCustomPermission(ctx, input)
			},
		},
		"update-vpc-connection": {
			Name:   "update-vpc-connection",
			Fields: fields_update_vpc_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVPCConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVPCConnection(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("quicksight", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
