package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cleanrooms"
)

var fields_batch_get_collaboration_analysis_template = []leanruntime.Field{
	{Name: "AnalysisTemplateArns", Flag: "analysis-template-arns", Type: "[]string", Required: true},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_batch_get_schema = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_batch_get_schema_analysis_rule = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "SchemaAnalysisRuleRequests", Flag: "schema-analysis-rule-requests", Type: "[]types.SchemaAnalysisRuleRequest", Required: true},
}

var fields_create_analysis_template = []leanruntime.Field{
	{Name: "AnalysisParameters", Flag: "analysis-parameters", Type: "[]types.AnalysisParameter", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ErrorMessageConfiguration", Flag: "error-message-configuration", Type: "*types.ErrorMessageConfiguration", Required: false},
	{Name: "Format", Flag: "format", Type: "types.AnalysisFormat", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "*types.AnalysisSchema", Required: false},
	{Name: "Source", Flag: "source", Type: "types.AnalysisSource", Required: true},
	{Name: "SyntheticDataParameters", Flag: "synthetic-data-parameters", Type: "types.SyntheticDataParameters", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_collaboration = []leanruntime.Field{
	{Name: "AllowedResultRegions", Flag: "allowed-result-regions", Type: "[]types.SupportedS3Region", Required: false},
	{Name: "AnalyticsEngine", Flag: "analytics-engine", Type: "types.AnalyticsEngine", Required: false},
	{Name: "AutoApprovedChangeRequestTypes", Flag: "auto-approved-change-request-types", Type: "[]types.AutoApprovedChangeType", Required: false},
	{Name: "CreatorDisplayName", Flag: "creator-display-name", Type: "*string", Required: true},
	{Name: "CreatorMLMemberAbilities", Flag: "creator-ml-member-abilities", Type: "*types.MLMemberAbilities", Required: false},
	{Name: "CreatorMemberAbilities", Flag: "creator-member-abilities", Type: "[]types.MemberAbility", Required: true},
	{Name: "CreatorPaymentConfiguration", Flag: "creator-payment-configuration", Type: "*types.PaymentConfiguration", Required: false},
	{Name: "DataEncryptionMetadata", Flag: "data-encryption-metadata", Type: "*types.DataEncryptionMetadata", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "IsMetricsEnabled", Flag: "is-metrics-enabled", Type: "*bool", Required: false},
	{Name: "JobLogStatus", Flag: "job-log-status", Type: "types.CollaborationJobLogStatus", Required: false},
	{Name: "Members", Flag: "members", Type: "[]types.MemberSpecification", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueryLogStatus", Flag: "query-log-status", Type: "types.CollaborationQueryLogStatus", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_collaboration_change_request = []leanruntime.Field{
	{Name: "Changes", Flag: "changes", Type: "[]types.ChangeInput", Required: true},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_create_configured_audience_model_association = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelArn", Flag: "configured-audience-model-arn", Type: "*string", Required: true},
	{Name: "ConfiguredAudienceModelAssociationName", Flag: "configured-audience-model-association-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ManageResourcePolicies", Flag: "manage-resource-policies", Type: "*bool", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configured_table = []leanruntime.Field{
	{Name: "AllowedColumns", Flag: "allowed-columns", Type: "[]string", Required: true},
	{Name: "AnalysisMethod", Flag: "analysis-method", Type: "types.AnalysisMethod", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SelectedAnalysisMethods", Flag: "selected-analysis-methods", Type: "[]types.SelectedAnalysisMethod", Required: false},
	{Name: "TableReference", Flag: "table-reference", Type: "types.TableReference", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configured_table_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRulePolicy", Flag: "analysis-rule-policy", Type: "types.ConfiguredTableAnalysisRulePolicy", Required: true},
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
}

var fields_create_configured_table_association = []leanruntime.Field{
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configured_table_association_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRulePolicy", Flag: "analysis-rule-policy", Type: "types.ConfiguredTableAssociationAnalysisRulePolicy", Required: true},
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAssociationAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_create_id_mapping_table = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InputReferenceConfig", Flag: "input-reference-config", Type: "*types.IdMappingTableInputReferenceConfig", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_id_namespace_association = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingConfig", Flag: "id-mapping-config", Type: "*types.IdMappingConfig", Required: false},
	{Name: "InputReferenceConfig", Flag: "input-reference-config", Type: "*types.IdNamespaceAssociationInputReferenceConfig", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_membership = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "DefaultJobResultConfiguration", Flag: "default-job-result-configuration", Type: "*types.MembershipProtectedJobResultConfiguration", Required: false},
	{Name: "DefaultResultConfiguration", Flag: "default-result-configuration", Type: "*types.MembershipProtectedQueryResultConfiguration", Required: false},
	{Name: "IsMetricsEnabled", Flag: "is-metrics-enabled", Type: "*bool", Required: false},
	{Name: "JobLogStatus", Flag: "job-log-status", Type: "types.MembershipJobLogStatus", Required: false},
	{Name: "PaymentConfiguration", Flag: "payment-configuration", Type: "*types.MembershipPaymentConfiguration", Required: false},
	{Name: "QueryLogStatus", Flag: "query-log-status", Type: "types.MembershipQueryLogStatus", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_privacy_budget_template = []leanruntime.Field{
	{Name: "AutoRefresh", Flag: "auto-refresh", Type: "types.PrivacyBudgetTemplateAutoRefresh", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "types.PrivacyBudgetTemplateParametersInput", Required: true},
	{Name: "PrivacyBudgetType", Flag: "privacy-budget-type", Type: "types.PrivacyBudgetType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_analysis_template = []leanruntime.Field{
	{Name: "AnalysisTemplateIdentifier", Flag: "analysis-template-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_collaboration = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_delete_configured_audience_model_association = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelAssociationIdentifier", Flag: "configured-audience-model-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_configured_table = []leanruntime.Field{
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
}

var fields_delete_configured_table_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
}

var fields_delete_configured_table_association = []leanruntime.Field{
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_configured_table_association_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAssociationAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_id_mapping_table = []leanruntime.Field{
	{Name: "IdMappingTableIdentifier", Flag: "id-mapping-table-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_id_namespace_association = []leanruntime.Field{
	{Name: "IdNamespaceAssociationIdentifier", Flag: "id-namespace-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_member = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_delete_membership = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_delete_privacy_budget_template = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "PrivacyBudgetTemplateIdentifier", Flag: "privacy-budget-template-identifier", Type: "*string", Required: true},
}

var fields_get_analysis_template = []leanruntime.Field{
	{Name: "AnalysisTemplateIdentifier", Flag: "analysis-template-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_collaboration = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_get_collaboration_analysis_template = []leanruntime.Field{
	{Name: "AnalysisTemplateArn", Flag: "analysis-template-arn", Type: "*string", Required: true},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_get_collaboration_change_request = []leanruntime.Field{
	{Name: "ChangeRequestIdentifier", Flag: "change-request-identifier", Type: "*string", Required: true},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_get_collaboration_configured_audience_model_association = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "ConfiguredAudienceModelAssociationIdentifier", Flag: "configured-audience-model-association-identifier", Type: "*string", Required: true},
}

var fields_get_collaboration_id_namespace_association = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "IdNamespaceAssociationIdentifier", Flag: "id-namespace-association-identifier", Type: "*string", Required: true},
}

var fields_get_collaboration_privacy_budget_template = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "PrivacyBudgetTemplateIdentifier", Flag: "privacy-budget-template-identifier", Type: "*string", Required: true},
}

var fields_get_configured_audience_model_association = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelAssociationIdentifier", Flag: "configured-audience-model-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_configured_table = []leanruntime.Field{
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
}

var fields_get_configured_table_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
}

var fields_get_configured_table_association = []leanruntime.Field{
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_configured_table_association_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAssociationAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_id_mapping_table = []leanruntime.Field{
	{Name: "IdMappingTableIdentifier", Flag: "id-mapping-table-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_id_namespace_association = []leanruntime.Field{
	{Name: "IdNamespaceAssociationIdentifier", Flag: "id-namespace-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_membership = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_get_privacy_budget_template = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "PrivacyBudgetTemplateIdentifier", Flag: "privacy-budget-template-identifier", Type: "*string", Required: true},
}

var fields_get_protected_job = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "ProtectedJobIdentifier", Flag: "protected-job-identifier", Type: "*string", Required: true},
}

var fields_get_protected_query = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "ProtectedQueryIdentifier", Flag: "protected-query-identifier", Type: "*string", Required: true},
}

var fields_get_schema = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_schema_analysis_rule = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AnalysisRuleType", Required: true},
}

var fields_list_analysis_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_analysis_templates = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_change_requests = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ChangeRequestStatus", Required: false},
}

var fields_list_collaboration_configured_audience_model_associations = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_id_namespace_associations = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_privacy_budget_templates = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collaboration_privacy_budgets = []leanruntime.Field{
	{Name: "AccessBudgetResourceArn", Flag: "access-budget-resource-arn", Type: "*string", Required: false},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrivacyBudgetType", Flag: "privacy-budget-type", Type: "types.PrivacyBudgetType", Required: true},
}

var fields_list_collaborations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberStatus", Flag: "member-status", Type: "types.FilterableMemberStatus", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configured_audience_model_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configured_table_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configured_tables = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_id_mapping_tables = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_id_namespace_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_memberships = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MembershipStatus", Required: false},
}

var fields_list_privacy_budget_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_privacy_budgets = []leanruntime.Field{
	{Name: "AccessBudgetResourceArn", Flag: "access-budget-resource-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrivacyBudgetType", Flag: "privacy-budget-type", Type: "types.PrivacyBudgetType", Required: true},
}

var fields_list_protected_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ProtectedJobStatus", Required: false},
}

var fields_list_protected_queries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ProtectedQueryStatus", Required: false},
}

var fields_list_schemas = []leanruntime.Field{
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaType", Flag: "schema-type", Type: "types.SchemaType", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_populate_id_mapping_table = []leanruntime.Field{
	{Name: "IdMappingTableIdentifier", Flag: "id-mapping-table-identifier", Type: "*string", Required: true},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_preview_privacy_impact = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "types.PreviewPrivacyImpactParametersInput", Required: true},
}

var fields_start_protected_job = []leanruntime.Field{
	{Name: "ComputeConfiguration", Flag: "compute-configuration", Type: "types.ProtectedJobComputeConfiguration", Required: false},
	{Name: "JobParameters", Flag: "job-parameters", Type: "*types.ProtectedJobParameters", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "ResultConfiguration", Flag: "result-configuration", Type: "*types.ProtectedJobResultConfigurationInput", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ProtectedJobType", Required: true},
}

var fields_start_protected_query = []leanruntime.Field{
	{Name: "ComputeConfiguration", Flag: "compute-configuration", Type: "types.ComputeConfiguration", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "ResultConfiguration", Flag: "result-configuration", Type: "*types.ProtectedQueryResultConfiguration", Required: false},
	{Name: "SqlParameters", Flag: "sql-parameters", Type: "*types.ProtectedQuerySQLParameters", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ProtectedQueryType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_analysis_template = []leanruntime.Field{
	{Name: "AnalysisTemplateIdentifier", Flag: "analysis-template-identifier", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_update_collaboration = []leanruntime.Field{
	{Name: "AnalyticsEngine", Flag: "analytics-engine", Type: "types.AnalyticsEngine", Required: false},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_collaboration_change_request = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.ChangeRequestAction", Required: true},
	{Name: "ChangeRequestIdentifier", Flag: "change-request-identifier", Type: "*string", Required: true},
	{Name: "CollaborationIdentifier", Flag: "collaboration-identifier", Type: "*string", Required: true},
}

var fields_update_configured_audience_model_association = []leanruntime.Field{
	{Name: "ConfiguredAudienceModelAssociationIdentifier", Flag: "configured-audience-model-association-identifier", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_configured_table = []leanruntime.Field{
	{Name: "AllowedColumns", Flag: "allowed-columns", Type: "[]string", Required: false},
	{Name: "AnalysisMethod", Flag: "analysis-method", Type: "types.AnalysisMethod", Required: false},
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SelectedAnalysisMethods", Flag: "selected-analysis-methods", Type: "[]types.SelectedAnalysisMethod", Required: false},
	{Name: "TableReference", Flag: "table-reference", Type: "types.TableReference", Required: false},
}

var fields_update_configured_table_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRulePolicy", Flag: "analysis-rule-policy", Type: "types.ConfiguredTableAnalysisRulePolicy", Required: true},
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableIdentifier", Flag: "configured-table-identifier", Type: "*string", Required: true},
}

var fields_update_configured_table_association = []leanruntime.Field{
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_configured_table_association_analysis_rule = []leanruntime.Field{
	{Name: "AnalysisRulePolicy", Flag: "analysis-rule-policy", Type: "types.ConfiguredTableAssociationAnalysisRulePolicy", Required: true},
	{Name: "AnalysisRuleType", Flag: "analysis-rule-type", Type: "types.ConfiguredTableAssociationAnalysisRuleType", Required: true},
	{Name: "ConfiguredTableAssociationIdentifier", Flag: "configured-table-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_update_id_mapping_table = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingTableIdentifier", Flag: "id-mapping-table-identifier", Type: "*string", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
}

var fields_update_id_namespace_association = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdMappingConfig", Flag: "id-mapping-config", Type: "*types.IdMappingConfig", Required: false},
	{Name: "IdNamespaceAssociationIdentifier", Flag: "id-namespace-association-identifier", Type: "*string", Required: true},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_membership = []leanruntime.Field{
	{Name: "DefaultJobResultConfiguration", Flag: "default-job-result-configuration", Type: "*types.MembershipProtectedJobResultConfiguration", Required: false},
	{Name: "DefaultResultConfiguration", Flag: "default-result-configuration", Type: "*types.MembershipProtectedQueryResultConfiguration", Required: false},
	{Name: "JobLogStatus", Flag: "job-log-status", Type: "types.MembershipJobLogStatus", Required: false},
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "QueryLogStatus", Flag: "query-log-status", Type: "types.MembershipQueryLogStatus", Required: false},
}

var fields_update_privacy_budget_template = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "types.PrivacyBudgetTemplateUpdateParameters", Required: false},
	{Name: "PrivacyBudgetTemplateIdentifier", Flag: "privacy-budget-template-identifier", Type: "*string", Required: true},
	{Name: "PrivacyBudgetType", Flag: "privacy-budget-type", Type: "types.PrivacyBudgetType", Required: true},
}

var fields_update_protected_job = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "ProtectedJobIdentifier", Flag: "protected-job-identifier", Type: "*string", Required: true},
	{Name: "TargetStatus", Flag: "target-status", Type: "types.TargetProtectedJobStatus", Required: true},
}

var fields_update_protected_query = []leanruntime.Field{
	{Name: "MembershipIdentifier", Flag: "membership-identifier", Type: "*string", Required: true},
	{Name: "ProtectedQueryIdentifier", Flag: "protected-query-identifier", Type: "*string", Required: true},
	{Name: "TargetStatus", Flag: "target-status", Type: "types.TargetProtectedQueryStatus", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-collaboration-analysis-template": {
			Name:   "batch-get-collaboration-analysis-template",
			Fields: fields_batch_get_collaboration_analysis_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCollaborationAnalysisTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_collaboration_analysis_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCollaborationAnalysisTemplate(ctx, input)
			},
		},
		"batch-get-schema": {
			Name:   "batch-get-schema",
			Fields: fields_batch_get_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetSchema(ctx, input)
			},
		},
		"batch-get-schema-analysis-rule": {
			Name:   "batch-get-schema-analysis-rule",
			Fields: fields_batch_get_schema_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetSchemaAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_schema_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetSchemaAnalysisRule(ctx, input)
			},
		},
		"create-analysis-template": {
			Name:   "create-analysis-template",
			Fields: fields_create_analysis_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnalysisTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_analysis_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnalysisTemplate(ctx, input)
			},
		},
		"create-collaboration": {
			Name:   "create-collaboration",
			Fields: fields_create_collaboration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCollaborationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_collaboration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCollaboration(ctx, input)
			},
		},
		"create-collaboration-change-request": {
			Name:   "create-collaboration-change-request",
			Fields: fields_create_collaboration_change_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCollaborationChangeRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_collaboration_change_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCollaborationChangeRequest(ctx, input)
			},
		},
		"create-configured-audience-model-association": {
			Name:   "create-configured-audience-model-association",
			Fields: fields_create_configured_audience_model_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredAudienceModelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_audience_model_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredAudienceModelAssociation(ctx, input)
			},
		},
		"create-configured-table": {
			Name:   "create-configured-table",
			Fields: fields_create_configured_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredTable(ctx, input)
			},
		},
		"create-configured-table-analysis-rule": {
			Name:   "create-configured-table-analysis-rule",
			Fields: fields_create_configured_table_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredTableAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_table_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredTableAnalysisRule(ctx, input)
			},
		},
		"create-configured-table-association": {
			Name:   "create-configured-table-association",
			Fields: fields_create_configured_table_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredTableAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_table_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredTableAssociation(ctx, input)
			},
		},
		"create-configured-table-association-analysis-rule": {
			Name:   "create-configured-table-association-analysis-rule",
			Fields: fields_create_configured_table_association_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfiguredTableAssociationAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configured_table_association_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguredTableAssociationAnalysisRule(ctx, input)
			},
		},
		"create-id-mapping-table": {
			Name:   "create-id-mapping-table",
			Fields: fields_create_id_mapping_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdMappingTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_id_mapping_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdMappingTable(ctx, input)
			},
		},
		"create-id-namespace-association": {
			Name:   "create-id-namespace-association",
			Fields: fields_create_id_namespace_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdNamespaceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_id_namespace_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdNamespaceAssociation(ctx, input)
			},
		},
		"create-membership": {
			Name:   "create-membership",
			Fields: fields_create_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMembership(ctx, input)
			},
		},
		"create-privacy-budget-template": {
			Name:   "create-privacy-budget-template",
			Fields: fields_create_privacy_budget_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePrivacyBudgetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_privacy_budget_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrivacyBudgetTemplate(ctx, input)
			},
		},
		"delete-analysis-template": {
			Name:   "delete-analysis-template",
			Fields: fields_delete_analysis_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnalysisTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_analysis_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnalysisTemplate(ctx, input)
			},
		},
		"delete-collaboration": {
			Name:   "delete-collaboration",
			Fields: fields_delete_collaboration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCollaborationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_collaboration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCollaboration(ctx, input)
			},
		},
		"delete-configured-audience-model-association": {
			Name:   "delete-configured-audience-model-association",
			Fields: fields_delete_configured_audience_model_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredAudienceModelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_audience_model_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredAudienceModelAssociation(ctx, input)
			},
		},
		"delete-configured-table": {
			Name:   "delete-configured-table",
			Fields: fields_delete_configured_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredTable(ctx, input)
			},
		},
		"delete-configured-table-analysis-rule": {
			Name:   "delete-configured-table-analysis-rule",
			Fields: fields_delete_configured_table_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredTableAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_table_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredTableAnalysisRule(ctx, input)
			},
		},
		"delete-configured-table-association": {
			Name:   "delete-configured-table-association",
			Fields: fields_delete_configured_table_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredTableAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_table_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredTableAssociation(ctx, input)
			},
		},
		"delete-configured-table-association-analysis-rule": {
			Name:   "delete-configured-table-association-analysis-rule",
			Fields: fields_delete_configured_table_association_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfiguredTableAssociationAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configured_table_association_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguredTableAssociationAnalysisRule(ctx, input)
			},
		},
		"delete-id-mapping-table": {
			Name:   "delete-id-mapping-table",
			Fields: fields_delete_id_mapping_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdMappingTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_id_mapping_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdMappingTable(ctx, input)
			},
		},
		"delete-id-namespace-association": {
			Name:   "delete-id-namespace-association",
			Fields: fields_delete_id_namespace_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdNamespaceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_id_namespace_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdNamespaceAssociation(ctx, input)
			},
		},
		"delete-member": {
			Name:   "delete-member",
			Fields: fields_delete_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMember(ctx, input)
			},
		},
		"delete-membership": {
			Name:   "delete-membership",
			Fields: fields_delete_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMembership(ctx, input)
			},
		},
		"delete-privacy-budget-template": {
			Name:   "delete-privacy-budget-template",
			Fields: fields_delete_privacy_budget_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePrivacyBudgetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_privacy_budget_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePrivacyBudgetTemplate(ctx, input)
			},
		},
		"get-analysis-template": {
			Name:   "get-analysis-template",
			Fields: fields_get_analysis_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnalysisTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_analysis_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnalysisTemplate(ctx, input)
			},
		},
		"get-collaboration": {
			Name:   "get-collaboration",
			Fields: fields_get_collaboration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaboration(ctx, input)
			},
		},
		"get-collaboration-analysis-template": {
			Name:   "get-collaboration-analysis-template",
			Fields: fields_get_collaboration_analysis_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationAnalysisTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_analysis_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationAnalysisTemplate(ctx, input)
			},
		},
		"get-collaboration-change-request": {
			Name:   "get-collaboration-change-request",
			Fields: fields_get_collaboration_change_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationChangeRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_change_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationChangeRequest(ctx, input)
			},
		},
		"get-collaboration-configured-audience-model-association": {
			Name:   "get-collaboration-configured-audience-model-association",
			Fields: fields_get_collaboration_configured_audience_model_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationConfiguredAudienceModelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_configured_audience_model_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationConfiguredAudienceModelAssociation(ctx, input)
			},
		},
		"get-collaboration-id-namespace-association": {
			Name:   "get-collaboration-id-namespace-association",
			Fields: fields_get_collaboration_id_namespace_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationIdNamespaceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_id_namespace_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationIdNamespaceAssociation(ctx, input)
			},
		},
		"get-collaboration-privacy-budget-template": {
			Name:   "get-collaboration-privacy-budget-template",
			Fields: fields_get_collaboration_privacy_budget_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCollaborationPrivacyBudgetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_collaboration_privacy_budget_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCollaborationPrivacyBudgetTemplate(ctx, input)
			},
		},
		"get-configured-audience-model-association": {
			Name:   "get-configured-audience-model-association",
			Fields: fields_get_configured_audience_model_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredAudienceModelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_audience_model_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredAudienceModelAssociation(ctx, input)
			},
		},
		"get-configured-table": {
			Name:   "get-configured-table",
			Fields: fields_get_configured_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredTable(ctx, input)
			},
		},
		"get-configured-table-analysis-rule": {
			Name:   "get-configured-table-analysis-rule",
			Fields: fields_get_configured_table_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredTableAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_table_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredTableAnalysisRule(ctx, input)
			},
		},
		"get-configured-table-association": {
			Name:   "get-configured-table-association",
			Fields: fields_get_configured_table_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredTableAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_table_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredTableAssociation(ctx, input)
			},
		},
		"get-configured-table-association-analysis-rule": {
			Name:   "get-configured-table-association-analysis-rule",
			Fields: fields_get_configured_table_association_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfiguredTableAssociationAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configured_table_association_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguredTableAssociationAnalysisRule(ctx, input)
			},
		},
		"get-id-mapping-table": {
			Name:   "get-id-mapping-table",
			Fields: fields_get_id_mapping_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdMappingTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_id_mapping_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdMappingTable(ctx, input)
			},
		},
		"get-id-namespace-association": {
			Name:   "get-id-namespace-association",
			Fields: fields_get_id_namespace_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdNamespaceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_id_namespace_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdNamespaceAssociation(ctx, input)
			},
		},
		"get-membership": {
			Name:   "get-membership",
			Fields: fields_get_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMembership(ctx, input)
			},
		},
		"get-privacy-budget-template": {
			Name:   "get-privacy-budget-template",
			Fields: fields_get_privacy_budget_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPrivacyBudgetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_privacy_budget_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPrivacyBudgetTemplate(ctx, input)
			},
		},
		"get-protected-job": {
			Name:   "get-protected-job",
			Fields: fields_get_protected_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProtectedJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_protected_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProtectedJob(ctx, input)
			},
		},
		"get-protected-query": {
			Name:   "get-protected-query",
			Fields: fields_get_protected_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProtectedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_protected_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProtectedQuery(ctx, input)
			},
		},
		"get-schema": {
			Name:   "get-schema",
			Fields: fields_get_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchema(ctx, input)
			},
		},
		"get-schema-analysis-rule": {
			Name:   "get-schema-analysis-rule",
			Fields: fields_get_schema_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaAnalysisRule(ctx, input)
			},
		},
		"list-analysis-templates": {
			Name:   "list-analysis-templates",
			Fields: fields_list_analysis_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalysisTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_analysis_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnalysisTemplates(ctx, input)
				}
				var results []*svc.ListAnalysisTemplatesOutput
				p := svc.NewListAnalysisTemplatesPaginator(client, input)
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
		"list-collaboration-analysis-templates": {
			Name:   "list-collaboration-analysis-templates",
			Fields: fields_list_collaboration_analysis_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationAnalysisTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_analysis_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationAnalysisTemplates(ctx, input)
				}
				var results []*svc.ListCollaborationAnalysisTemplatesOutput
				p := svc.NewListCollaborationAnalysisTemplatesPaginator(client, input)
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
		"list-collaboration-change-requests": {
			Name:   "list-collaboration-change-requests",
			Fields: fields_list_collaboration_change_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationChangeRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_change_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationChangeRequests(ctx, input)
				}
				var results []*svc.ListCollaborationChangeRequestsOutput
				p := svc.NewListCollaborationChangeRequestsPaginator(client, input)
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
		"list-collaboration-configured-audience-model-associations": {
			Name:   "list-collaboration-configured-audience-model-associations",
			Fields: fields_list_collaboration_configured_audience_model_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationConfiguredAudienceModelAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_configured_audience_model_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationConfiguredAudienceModelAssociations(ctx, input)
				}
				var results []*svc.ListCollaborationConfiguredAudienceModelAssociationsOutput
				p := svc.NewListCollaborationConfiguredAudienceModelAssociationsPaginator(client, input)
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
		"list-collaboration-id-namespace-associations": {
			Name:   "list-collaboration-id-namespace-associations",
			Fields: fields_list_collaboration_id_namespace_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationIdNamespaceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_id_namespace_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationIdNamespaceAssociations(ctx, input)
				}
				var results []*svc.ListCollaborationIdNamespaceAssociationsOutput
				p := svc.NewListCollaborationIdNamespaceAssociationsPaginator(client, input)
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
		"list-collaboration-privacy-budget-templates": {
			Name:   "list-collaboration-privacy-budget-templates",
			Fields: fields_list_collaboration_privacy_budget_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationPrivacyBudgetTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_privacy_budget_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationPrivacyBudgetTemplates(ctx, input)
				}
				var results []*svc.ListCollaborationPrivacyBudgetTemplatesOutput
				p := svc.NewListCollaborationPrivacyBudgetTemplatesPaginator(client, input)
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
		"list-collaboration-privacy-budgets": {
			Name:   "list-collaboration-privacy-budgets",
			Fields: fields_list_collaboration_privacy_budgets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationPrivacyBudgetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaboration_privacy_budgets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborationPrivacyBudgets(ctx, input)
				}
				var results []*svc.ListCollaborationPrivacyBudgetsOutput
				p := svc.NewListCollaborationPrivacyBudgetsPaginator(client, input)
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
		"list-collaborations": {
			Name:   "list-collaborations",
			Fields: fields_list_collaborations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollaborationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collaborations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollaborations(ctx, input)
				}
				var results []*svc.ListCollaborationsOutput
				p := svc.NewListCollaborationsPaginator(client, input)
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
		"list-configured-audience-model-associations": {
			Name:   "list-configured-audience-model-associations",
			Fields: fields_list_configured_audience_model_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfiguredAudienceModelAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configured_audience_model_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfiguredAudienceModelAssociations(ctx, input)
				}
				var results []*svc.ListConfiguredAudienceModelAssociationsOutput
				p := svc.NewListConfiguredAudienceModelAssociationsPaginator(client, input)
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
		"list-configured-table-associations": {
			Name:   "list-configured-table-associations",
			Fields: fields_list_configured_table_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfiguredTableAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configured_table_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfiguredTableAssociations(ctx, input)
				}
				var results []*svc.ListConfiguredTableAssociationsOutput
				p := svc.NewListConfiguredTableAssociationsPaginator(client, input)
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
		"list-configured-tables": {
			Name:   "list-configured-tables",
			Fields: fields_list_configured_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfiguredTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configured_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfiguredTables(ctx, input)
				}
				var results []*svc.ListConfiguredTablesOutput
				p := svc.NewListConfiguredTablesPaginator(client, input)
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
		"list-id-mapping-tables": {
			Name:   "list-id-mapping-tables",
			Fields: fields_list_id_mapping_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdMappingTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_id_mapping_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdMappingTables(ctx, input)
				}
				var results []*svc.ListIdMappingTablesOutput
				p := svc.NewListIdMappingTablesPaginator(client, input)
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
		"list-id-namespace-associations": {
			Name:   "list-id-namespace-associations",
			Fields: fields_list_id_namespace_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdNamespaceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_id_namespace_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdNamespaceAssociations(ctx, input)
				}
				var results []*svc.ListIdNamespaceAssociationsOutput
				p := svc.NewListIdNamespaceAssociationsPaginator(client, input)
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
		"list-members": {
			Name:   "list-members",
			Fields: fields_list_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMembers(ctx, input)
				}
				var results []*svc.ListMembersOutput
				p := svc.NewListMembersPaginator(client, input)
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
		"list-memberships": {
			Name:   "list-memberships",
			Fields: fields_list_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMemberships(ctx, input)
				}
				var results []*svc.ListMembershipsOutput
				p := svc.NewListMembershipsPaginator(client, input)
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
		"list-privacy-budget-templates": {
			Name:   "list-privacy-budget-templates",
			Fields: fields_list_privacy_budget_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrivacyBudgetTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_privacy_budget_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrivacyBudgetTemplates(ctx, input)
				}
				var results []*svc.ListPrivacyBudgetTemplatesOutput
				p := svc.NewListPrivacyBudgetTemplatesPaginator(client, input)
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
		"list-privacy-budgets": {
			Name:   "list-privacy-budgets",
			Fields: fields_list_privacy_budgets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrivacyBudgetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_privacy_budgets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrivacyBudgets(ctx, input)
				}
				var results []*svc.ListPrivacyBudgetsOutput
				p := svc.NewListPrivacyBudgetsPaginator(client, input)
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
		"list-protected-jobs": {
			Name:   "list-protected-jobs",
			Fields: fields_list_protected_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectedJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protected_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtectedJobs(ctx, input)
				}
				var results []*svc.ListProtectedJobsOutput
				p := svc.NewListProtectedJobsPaginator(client, input)
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
		"list-protected-queries": {
			Name:   "list-protected-queries",
			Fields: fields_list_protected_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectedQueriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protected_queries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtectedQueries(ctx, input)
				}
				var results []*svc.ListProtectedQueriesOutput
				p := svc.NewListProtectedQueriesPaginator(client, input)
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
		"list-schemas": {
			Name:   "list-schemas",
			Fields: fields_list_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemas(ctx, input)
				}
				var results []*svc.ListSchemasOutput
				p := svc.NewListSchemasPaginator(client, input)
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
		"populate-id-mapping-table": {
			Name:   "populate-id-mapping-table",
			Fields: fields_populate_id_mapping_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PopulateIdMappingTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_populate_id_mapping_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PopulateIdMappingTable(ctx, input)
			},
		},
		"preview-privacy-impact": {
			Name:   "preview-privacy-impact",
			Fields: fields_preview_privacy_impact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PreviewPrivacyImpactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_preview_privacy_impact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PreviewPrivacyImpact(ctx, input)
			},
		},
		"start-protected-job": {
			Name:   "start-protected-job",
			Fields: fields_start_protected_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartProtectedJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_protected_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartProtectedJob(ctx, input)
			},
		},
		"start-protected-query": {
			Name:   "start-protected-query",
			Fields: fields_start_protected_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartProtectedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_protected_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartProtectedQuery(ctx, input)
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
		"update-analysis-template": {
			Name:   "update-analysis-template",
			Fields: fields_update_analysis_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnalysisTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_analysis_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnalysisTemplate(ctx, input)
			},
		},
		"update-collaboration": {
			Name:   "update-collaboration",
			Fields: fields_update_collaboration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCollaborationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_collaboration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCollaboration(ctx, input)
			},
		},
		"update-collaboration-change-request": {
			Name:   "update-collaboration-change-request",
			Fields: fields_update_collaboration_change_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCollaborationChangeRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_collaboration_change_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCollaborationChangeRequest(ctx, input)
			},
		},
		"update-configured-audience-model-association": {
			Name:   "update-configured-audience-model-association",
			Fields: fields_update_configured_audience_model_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfiguredAudienceModelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configured_audience_model_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguredAudienceModelAssociation(ctx, input)
			},
		},
		"update-configured-table": {
			Name:   "update-configured-table",
			Fields: fields_update_configured_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfiguredTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configured_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguredTable(ctx, input)
			},
		},
		"update-configured-table-analysis-rule": {
			Name:   "update-configured-table-analysis-rule",
			Fields: fields_update_configured_table_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfiguredTableAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configured_table_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguredTableAnalysisRule(ctx, input)
			},
		},
		"update-configured-table-association": {
			Name:   "update-configured-table-association",
			Fields: fields_update_configured_table_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfiguredTableAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configured_table_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguredTableAssociation(ctx, input)
			},
		},
		"update-configured-table-association-analysis-rule": {
			Name:   "update-configured-table-association-analysis-rule",
			Fields: fields_update_configured_table_association_analysis_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfiguredTableAssociationAnalysisRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configured_table_association_analysis_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguredTableAssociationAnalysisRule(ctx, input)
			},
		},
		"update-id-mapping-table": {
			Name:   "update-id-mapping-table",
			Fields: fields_update_id_mapping_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdMappingTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_id_mapping_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdMappingTable(ctx, input)
			},
		},
		"update-id-namespace-association": {
			Name:   "update-id-namespace-association",
			Fields: fields_update_id_namespace_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdNamespaceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_id_namespace_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdNamespaceAssociation(ctx, input)
			},
		},
		"update-membership": {
			Name:   "update-membership",
			Fields: fields_update_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMembership(ctx, input)
			},
		},
		"update-privacy-budget-template": {
			Name:   "update-privacy-budget-template",
			Fields: fields_update_privacy_budget_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePrivacyBudgetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_privacy_budget_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePrivacyBudgetTemplate(ctx, input)
			},
		},
		"update-protected-job": {
			Name:   "update-protected-job",
			Fields: fields_update_protected_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProtectedJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_protected_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProtectedJob(ctx, input)
			},
		},
		"update-protected-query": {
			Name:   "update-protected-query",
			Fields: fields_update_protected_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProtectedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_protected_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProtectedQuery(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cleanrooms", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
