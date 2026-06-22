package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/securityhub"
)

var fields_accept_administrator_invitation = []leanruntime.Field{
	{Name: "AdministratorId", Flag: "administrator-id", Type: "*string", Required: true},
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
}

var fields_accept_invitation = []leanruntime.Field{
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
	{Name: "MasterId", Flag: "master-id", Type: "*string", Required: true},
}

var fields_batch_delete_automation_rules = []leanruntime.Field{
	{Name: "AutomationRulesArns", Flag: "automation-rules-arns", Type: "[]string", Required: true},
}

var fields_batch_disable_standards = []leanruntime.Field{
	{Name: "StandardsSubscriptionArns", Flag: "standards-subscription-arns", Type: "[]string", Required: true},
}

var fields_batch_enable_standards = []leanruntime.Field{
	{Name: "StandardsSubscriptionRequests", Flag: "standards-subscription-requests", Type: "[]types.StandardsSubscriptionRequest", Required: true},
}

var fields_batch_get_automation_rules = []leanruntime.Field{
	{Name: "AutomationRulesArns", Flag: "automation-rules-arns", Type: "[]string", Required: true},
}

var fields_batch_get_configuration_policy_associations = []leanruntime.Field{
	{Name: "ConfigurationPolicyAssociationIdentifiers", Flag: "configuration-policy-association-identifiers", Type: "[]types.ConfigurationPolicyAssociation", Required: true},
}

var fields_batch_get_security_controls = []leanruntime.Field{
	{Name: "SecurityControlIds", Flag: "security-control-ids", Type: "[]string", Required: true},
}

var fields_batch_get_standards_control_associations = []leanruntime.Field{
	{Name: "StandardsControlAssociationIds", Flag: "standards-control-association-ids", Type: "[]types.StandardsControlAssociationId", Required: true},
}

var fields_batch_import_findings = []leanruntime.Field{
	{Name: "Findings", Flag: "findings", Type: "[]types.AwsSecurityFinding", Required: true},
}

var fields_batch_update_automation_rules = []leanruntime.Field{
	{Name: "UpdateAutomationRulesRequestItems", Flag: "update-automation-rules-request-items", Type: "[]types.UpdateAutomationRulesRequestItem", Required: true},
}

var fields_batch_update_findings = []leanruntime.Field{
	{Name: "Confidence", Flag: "confidence", Type: "*int32", Required: false},
	{Name: "Criticality", Flag: "criticality", Type: "*int32", Required: false},
	{Name: "FindingIdentifiers", Flag: "finding-identifiers", Type: "[]types.AwsSecurityFindingIdentifier", Required: true},
	{Name: "Note", Flag: "note", Type: "*types.NoteUpdate", Required: false},
	{Name: "RelatedFindings", Flag: "related-findings", Type: "[]types.RelatedFinding", Required: false},
	{Name: "Severity", Flag: "severity", Type: "*types.SeverityUpdate", Required: false},
	{Name: "Types", Flag: "types", Type: "[]string", Required: false},
	{Name: "UserDefinedFields", Flag: "user-defined-fields", Type: "map[string]string", Required: false},
	{Name: "VerificationState", Flag: "verification-state", Type: "types.VerificationState", Required: false},
	{Name: "Workflow", Flag: "workflow", Type: "*types.WorkflowUpdate", Required: false},
}

var fields_batch_update_findings_v2 = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "FindingIdentifiers", Flag: "finding-identifiers", Type: "[]types.OcsfFindingIdentifier", Required: false},
	{Name: "MetadataUids", Flag: "metadata-uids", Type: "[]string", Required: false},
	{Name: "SeverityId", Flag: "severity-id", Type: "*int32", Required: false},
	{Name: "StatusId", Flag: "status-id", Type: "*int32", Required: false},
}

var fields_batch_update_standards_control_associations = []leanruntime.Field{
	{Name: "StandardsControlAssociationUpdates", Flag: "standards-control-association-updates", Type: "[]types.StandardsControlAssociationUpdate", Required: true},
}

var fields_create_action_target = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_aggregator_v2 = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LinkedRegions", Flag: "linked-regions", Type: "[]string", Required: false},
	{Name: "RegionLinkingMode", Flag: "region-linking-mode", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_automation_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.AutomationRulesAction", Required: true},
	{Name: "Criteria", Flag: "criteria", Type: "*types.AutomationRulesFindingFilters", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "IsTerminal", Flag: "is-terminal", Type: "*bool", Required: false},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "RuleOrder", Flag: "rule-order", Type: "*int32", Required: true},
	{Name: "RuleStatus", Flag: "rule-status", Type: "types.RuleStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_automation_rule_v2 = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.AutomationRulesActionV2", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Criteria", Flag: "criteria", Type: "types.Criteria", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "RuleOrder", Flag: "rule-order", Type: "*float32", Required: true},
	{Name: "RuleStatus", Flag: "rule-status", Type: "types.RuleStatusV2", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_configuration_policy = []leanruntime.Field{
	{Name: "ConfigurationPolicy", Flag: "configuration-policy", Type: "types.Policy", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_connector_v2 = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "types.ProviderConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_finding_aggregator = []leanruntime.Field{
	{Name: "RegionLinkingMode", Flag: "region-linking-mode", Type: "*string", Required: true},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
}

var fields_create_insight = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.AwsSecurityFindingFilters", Required: true},
	{Name: "GroupByAttribute", Flag: "group-by-attribute", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_members = []leanruntime.Field{
	{Name: "AccountDetails", Flag: "account-details", Type: "[]types.AccountDetails", Required: true},
}

var fields_create_ticket_v2 = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "FindingMetadataUid", Flag: "finding-metadata-uid", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.TicketCreationMode", Required: false},
}

var fields_decline_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_delete_action_target = []leanruntime.Field{
	{Name: "ActionTargetArn", Flag: "action-target-arn", Type: "*string", Required: true},
}

var fields_delete_aggregator_v2 = []leanruntime.Field{
	{Name: "AggregatorV2Arn", Flag: "aggregator-v2-arn", Type: "*string", Required: true},
}

var fields_delete_automation_rule_v2 = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_configuration_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_connector_v2 = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
}

var fields_delete_finding_aggregator = []leanruntime.Field{
	{Name: "FindingAggregatorArn", Flag: "finding-aggregator-arn", Type: "*string", Required: true},
}

var fields_delete_insight = []leanruntime.Field{
	{Name: "InsightArn", Flag: "insight-arn", Type: "*string", Required: true},
}

var fields_delete_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_delete_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_describe_action_targets = []leanruntime.Field{
	{Name: "ActionTargetArns", Flag: "action-target-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_hub = []leanruntime.Field{
	{Name: "HubArn", Flag: "hub-arn", Type: "*string", Required: false},
}

var fields_describe_organization_configuration = []leanruntime.Field{}

var fields_describe_products = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProductArn", Flag: "product-arn", Type: "*string", Required: false},
}

var fields_describe_products_v2 = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_security_hub_v2 = []leanruntime.Field{}

var fields_describe_standards = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_standards_controls = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StandardsSubscriptionArn", Flag: "standards-subscription-arn", Type: "*string", Required: true},
}

var fields_disable_import_findings_for_product = []leanruntime.Field{
	{Name: "ProductSubscriptionArn", Flag: "product-subscription-arn", Type: "*string", Required: true},
}

var fields_disable_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
	{Name: "Feature", Flag: "feature", Type: "types.SecurityHubFeature", Required: false},
}

var fields_disable_security_hub = []leanruntime.Field{}

var fields_disable_security_hub_v2 = []leanruntime.Field{}

var fields_disassociate_from_administrator_account = []leanruntime.Field{}

var fields_disassociate_from_master_account = []leanruntime.Field{}

var fields_disassociate_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_enable_import_findings_for_product = []leanruntime.Field{
	{Name: "ProductArn", Flag: "product-arn", Type: "*string", Required: true},
}

var fields_enable_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
	{Name: "Feature", Flag: "feature", Type: "types.SecurityHubFeature", Required: false},
}

var fields_enable_security_hub = []leanruntime.Field{
	{Name: "ControlFindingGenerator", Flag: "control-finding-generator", Type: "types.ControlFindingGenerator", Required: false},
	{Name: "EnableDefaultStandards", Flag: "enable-default-standards", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_enable_security_hub_v2 = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_get_administrator_account = []leanruntime.Field{}

var fields_get_aggregator_v2 = []leanruntime.Field{
	{Name: "AggregatorV2Arn", Flag: "aggregator-v2-arn", Type: "*string", Required: true},
}

var fields_get_automation_rule_v2 = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_configuration_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_configuration_policy_association = []leanruntime.Field{
	{Name: "Target", Flag: "target", Type: "types.Target", Required: true},
}

var fields_get_connector_v2 = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
}

var fields_get_enabled_standards = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StandardsSubscriptionArns", Flag: "standards-subscription-arns", Type: "[]string", Required: false},
}

var fields_get_finding_aggregator = []leanruntime.Field{
	{Name: "FindingAggregatorArn", Flag: "finding-aggregator-arn", Type: "*string", Required: true},
}

var fields_get_finding_history = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "FindingIdentifier", Flag: "finding-identifier", Type: "*types.AwsSecurityFindingIdentifier", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_finding_statistics_v2 = []leanruntime.Field{
	{Name: "GroupByRules", Flag: "group-by-rules", Type: "[]types.GroupByRule", Required: true},
	{Name: "MaxStatisticResults", Flag: "max-statistic-results", Type: "*int32", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_get_findings = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.AwsSecurityFindingFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "[]types.SortCriterion", Required: false},
}

var fields_get_findings_trends_v2 = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.FindingsTrendsFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_findings_v2 = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.OcsfFindingFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "[]types.SortCriterion", Required: false},
}

var fields_get_insight_results = []leanruntime.Field{
	{Name: "InsightArn", Flag: "insight-arn", Type: "*string", Required: true},
}

var fields_get_insights = []leanruntime.Field{
	{Name: "InsightArns", Flag: "insight-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_invitations_count = []leanruntime.Field{}

var fields_get_master_account = []leanruntime.Field{}

var fields_get_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_get_resources_statistics_v2 = []leanruntime.Field{
	{Name: "GroupByRules", Flag: "group-by-rules", Type: "[]types.ResourceGroupByRule", Required: true},
	{Name: "MaxStatisticResults", Flag: "max-statistic-results", Type: "*int32", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_get_resources_trends_v2 = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ResourcesTrendsFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_resources_v2 = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ResourcesFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "[]types.SortCriterion", Required: false},
}

var fields_get_security_control_definition = []leanruntime.Field{
	{Name: "SecurityControlId", Flag: "security-control-id", Type: "*string", Required: true},
}

var fields_invite_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_list_aggregators_v2 = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_automation_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_automation_rules_v2 = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_policy_associations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.AssociationFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connectors_v2 = []leanruntime.Field{
	{Name: "ConnectorStatus", Flag: "connector-status", Type: "types.ConnectorStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProviderName", Flag: "provider-name", Type: "types.ConnectorProviderName", Required: false},
}

var fields_list_enabled_products_for_import = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_finding_aggregators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_invitations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OnlyAssociated", Flag: "only-associated", Type: "*bool", Required: false},
}

var fields_list_organization_admin_accounts = []leanruntime.Field{
	{Name: "Feature", Flag: "feature", Type: "types.SecurityHubFeature", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_security_control_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StandardsArn", Flag: "standards-arn", Type: "*string", Required: false},
}

var fields_list_standards_control_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityControlId", Flag: "security-control-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_connector_v2 = []leanruntime.Field{
	{Name: "AuthCode", Flag: "auth-code", Type: "*string", Required: true},
	{Name: "AuthState", Flag: "auth-state", Type: "*string", Required: true},
}

var fields_start_configuration_policy_association = []leanruntime.Field{
	{Name: "ConfigurationPolicyIdentifier", Flag: "configuration-policy-identifier", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "types.Target", Required: true},
}

var fields_start_configuration_policy_disassociation = []leanruntime.Field{
	{Name: "ConfigurationPolicyIdentifier", Flag: "configuration-policy-identifier", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "types.Target", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_action_target = []leanruntime.Field{
	{Name: "ActionTargetArn", Flag: "action-target-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_aggregator_v2 = []leanruntime.Field{
	{Name: "AggregatorV2Arn", Flag: "aggregator-v2-arn", Type: "*string", Required: true},
	{Name: "LinkedRegions", Flag: "linked-regions", Type: "[]string", Required: false},
	{Name: "RegionLinkingMode", Flag: "region-linking-mode", Type: "*string", Required: true},
}

var fields_update_automation_rule_v2 = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.AutomationRulesActionV2", Required: false},
	{Name: "Criteria", Flag: "criteria", Type: "types.Criteria", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: false},
	{Name: "RuleOrder", Flag: "rule-order", Type: "*float32", Required: false},
	{Name: "RuleStatus", Flag: "rule-status", Type: "types.RuleStatusV2", Required: false},
}

var fields_update_configuration_policy = []leanruntime.Field{
	{Name: "ConfigurationPolicy", Flag: "configuration-policy", Type: "types.Policy", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "UpdatedReason", Flag: "updated-reason", Type: "*string", Required: false},
}

var fields_update_connector_v2 = []leanruntime.Field{
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Provider", Flag: "provider", Type: "types.ProviderUpdateConfiguration", Required: false},
}

var fields_update_finding_aggregator = []leanruntime.Field{
	{Name: "FindingAggregatorArn", Flag: "finding-aggregator-arn", Type: "*string", Required: true},
	{Name: "RegionLinkingMode", Flag: "region-linking-mode", Type: "*string", Required: true},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
}

var fields_update_findings = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.AwsSecurityFindingFilters", Required: true},
	{Name: "Note", Flag: "note", Type: "*types.NoteUpdate", Required: false},
	{Name: "RecordState", Flag: "record-state", Type: "types.RecordState", Required: false},
}

var fields_update_insight = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.AwsSecurityFindingFilters", Required: false},
	{Name: "GroupByAttribute", Flag: "group-by-attribute", Type: "*string", Required: false},
	{Name: "InsightArn", Flag: "insight-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnable", Flag: "auto-enable", Type: "*bool", Required: true},
	{Name: "AutoEnableStandards", Flag: "auto-enable-standards", Type: "types.AutoEnableStandards", Required: false},
	{Name: "OrganizationConfiguration", Flag: "organization-configuration", Type: "*types.OrganizationConfiguration", Required: false},
}

var fields_update_security_control = []leanruntime.Field{
	{Name: "LastUpdateReason", Flag: "last-update-reason", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.ParameterConfiguration", Required: true},
	{Name: "SecurityControlId", Flag: "security-control-id", Type: "*string", Required: true},
}

var fields_update_security_hub_configuration = []leanruntime.Field{
	{Name: "AutoEnableControls", Flag: "auto-enable-controls", Type: "*bool", Required: false},
	{Name: "ControlFindingGenerator", Flag: "control-finding-generator", Type: "types.ControlFindingGenerator", Required: false},
}

var fields_update_standards_control = []leanruntime.Field{
	{Name: "ControlStatus", Flag: "control-status", Type: "types.ControlStatus", Required: false},
	{Name: "DisabledReason", Flag: "disabled-reason", Type: "*string", Required: false},
	{Name: "StandardsControlArn", Flag: "standards-control-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-administrator-invitation": {
			Name:   "accept-administrator-invitation",
			Fields: fields_accept_administrator_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptAdministratorInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_administrator_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptAdministratorInvitation(ctx, input)
			},
		},
		"accept-invitation": {
			Name:   "accept-invitation",
			Fields: fields_accept_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptInvitation(ctx, input)
			},
		},
		"batch-delete-automation-rules": {
			Name:   "batch-delete-automation-rules",
			Fields: fields_batch_delete_automation_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteAutomationRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_automation_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteAutomationRules(ctx, input)
			},
		},
		"batch-disable-standards": {
			Name:   "batch-disable-standards",
			Fields: fields_batch_disable_standards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisableStandardsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disable_standards, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisableStandards(ctx, input)
			},
		},
		"batch-enable-standards": {
			Name:   "batch-enable-standards",
			Fields: fields_batch_enable_standards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchEnableStandardsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_enable_standards, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchEnableStandards(ctx, input)
			},
		},
		"batch-get-automation-rules": {
			Name:   "batch-get-automation-rules",
			Fields: fields_batch_get_automation_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAutomationRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_automation_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetAutomationRules(ctx, input)
			},
		},
		"batch-get-configuration-policy-associations": {
			Name:   "batch-get-configuration-policy-associations",
			Fields: fields_batch_get_configuration_policy_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetConfigurationPolicyAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_configuration_policy_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetConfigurationPolicyAssociations(ctx, input)
			},
		},
		"batch-get-security-controls": {
			Name:   "batch-get-security-controls",
			Fields: fields_batch_get_security_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetSecurityControlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_security_controls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetSecurityControls(ctx, input)
			},
		},
		"batch-get-standards-control-associations": {
			Name:   "batch-get-standards-control-associations",
			Fields: fields_batch_get_standards_control_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetStandardsControlAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_standards_control_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetStandardsControlAssociations(ctx, input)
			},
		},
		"batch-import-findings": {
			Name:   "batch-import-findings",
			Fields: fields_batch_import_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchImportFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_import_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchImportFindings(ctx, input)
			},
		},
		"batch-update-automation-rules": {
			Name:   "batch-update-automation-rules",
			Fields: fields_batch_update_automation_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateAutomationRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_automation_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateAutomationRules(ctx, input)
			},
		},
		"batch-update-findings": {
			Name:   "batch-update-findings",
			Fields: fields_batch_update_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateFindings(ctx, input)
			},
		},
		"batch-update-findings-v2": {
			Name:   "batch-update-findings-v2",
			Fields: fields_batch_update_findings_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateFindingsV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_findings_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateFindingsV2(ctx, input)
			},
		},
		"batch-update-standards-control-associations": {
			Name:   "batch-update-standards-control-associations",
			Fields: fields_batch_update_standards_control_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateStandardsControlAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_standards_control_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateStandardsControlAssociations(ctx, input)
			},
		},
		"create-action-target": {
			Name:   "create-action-target",
			Fields: fields_create_action_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateActionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_action_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateActionTarget(ctx, input)
			},
		},
		"create-aggregator-v2": {
			Name:   "create-aggregator-v2",
			Fields: fields_create_aggregator_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAggregatorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_aggregator_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAggregatorV2(ctx, input)
			},
		},
		"create-automation-rule": {
			Name:   "create-automation-rule",
			Fields: fields_create_automation_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutomationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_automation_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutomationRule(ctx, input)
			},
		},
		"create-automation-rule-v2": {
			Name:   "create-automation-rule-v2",
			Fields: fields_create_automation_rule_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutomationRuleV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_automation_rule_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutomationRuleV2(ctx, input)
			},
		},
		"create-configuration-policy": {
			Name:   "create-configuration-policy",
			Fields: fields_create_configuration_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationPolicy(ctx, input)
			},
		},
		"create-connector-v2": {
			Name:   "create-connector-v2",
			Fields: fields_create_connector_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectorV2(ctx, input)
			},
		},
		"create-finding-aggregator": {
			Name:   "create-finding-aggregator",
			Fields: fields_create_finding_aggregator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFindingAggregatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_finding_aggregator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFindingAggregator(ctx, input)
			},
		},
		"create-insight": {
			Name:   "create-insight",
			Fields: fields_create_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInsight(ctx, input)
			},
		},
		"create-members": {
			Name:   "create-members",
			Fields: fields_create_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMembers(ctx, input)
			},
		},
		"create-ticket-v2": {
			Name:   "create-ticket-v2",
			Fields: fields_create_ticket_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTicketV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ticket_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTicketV2(ctx, input)
			},
		},
		"decline-invitations": {
			Name:   "decline-invitations",
			Fields: fields_decline_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeclineInvitationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decline_invitations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeclineInvitations(ctx, input)
			},
		},
		"delete-action-target": {
			Name:   "delete-action-target",
			Fields: fields_delete_action_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteActionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_action_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteActionTarget(ctx, input)
			},
		},
		"delete-aggregator-v2": {
			Name:   "delete-aggregator-v2",
			Fields: fields_delete_aggregator_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAggregatorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_aggregator_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAggregatorV2(ctx, input)
			},
		},
		"delete-automation-rule-v2": {
			Name:   "delete-automation-rule-v2",
			Fields: fields_delete_automation_rule_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutomationRuleV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_automation_rule_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutomationRuleV2(ctx, input)
			},
		},
		"delete-configuration-policy": {
			Name:   "delete-configuration-policy",
			Fields: fields_delete_configuration_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationPolicy(ctx, input)
			},
		},
		"delete-connector-v2": {
			Name:   "delete-connector-v2",
			Fields: fields_delete_connector_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectorV2(ctx, input)
			},
		},
		"delete-finding-aggregator": {
			Name:   "delete-finding-aggregator",
			Fields: fields_delete_finding_aggregator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFindingAggregatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_finding_aggregator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFindingAggregator(ctx, input)
			},
		},
		"delete-insight": {
			Name:   "delete-insight",
			Fields: fields_delete_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInsight(ctx, input)
			},
		},
		"delete-invitations": {
			Name:   "delete-invitations",
			Fields: fields_delete_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInvitationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_invitations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInvitations(ctx, input)
			},
		},
		"delete-members": {
			Name:   "delete-members",
			Fields: fields_delete_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMembers(ctx, input)
			},
		},
		"describe-action-targets": {
			Name:   "describe-action-targets",
			Fields: fields_describe_action_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActionTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_action_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeActionTargets(ctx, input)
				}
				var results []*svc.DescribeActionTargetsOutput
				p := svc.NewDescribeActionTargetsPaginator(client, input)
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
		"describe-hub": {
			Name:   "describe-hub",
			Fields: fields_describe_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHub(ctx, input)
			},
		},
		"describe-organization-configuration": {
			Name:   "describe-organization-configuration",
			Fields: fields_describe_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganizationConfiguration(ctx, input)
			},
		},
		"describe-products": {
			Name:   "describe-products",
			Fields: fields_describe_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProductsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_products, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeProducts(ctx, input)
				}
				var results []*svc.DescribeProductsOutput
				p := svc.NewDescribeProductsPaginator(client, input)
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
		"describe-products-v2": {
			Name:   "describe-products-v2",
			Fields: fields_describe_products_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProductsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_products_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeProductsV2(ctx, input)
				}
				var results []*svc.DescribeProductsV2Output
				p := svc.NewDescribeProductsV2Paginator(client, input)
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
		"describe-security-hub-v2": {
			Name:   "describe-security-hub-v2",
			Fields: fields_describe_security_hub_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityHubV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_hub_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityHubV2(ctx, input)
			},
		},
		"describe-standards": {
			Name:   "describe-standards",
			Fields: fields_describe_standards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStandardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_standards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStandards(ctx, input)
				}
				var results []*svc.DescribeStandardsOutput
				p := svc.NewDescribeStandardsPaginator(client, input)
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
		"describe-standards-controls": {
			Name:   "describe-standards-controls",
			Fields: fields_describe_standards_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStandardsControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_standards_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStandardsControls(ctx, input)
				}
				var results []*svc.DescribeStandardsControlsOutput
				p := svc.NewDescribeStandardsControlsPaginator(client, input)
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
		"disable-import-findings-for-product": {
			Name:   "disable-import-findings-for-product",
			Fields: fields_disable_import_findings_for_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableImportFindingsForProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_import_findings_for_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableImportFindingsForProduct(ctx, input)
			},
		},
		"disable-organization-admin-account": {
			Name:   "disable-organization-admin-account",
			Fields: fields_disable_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableOrganizationAdminAccount(ctx, input)
			},
		},
		"disable-security-hub": {
			Name:   "disable-security-hub",
			Fields: fields_disable_security_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSecurityHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_security_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSecurityHub(ctx, input)
			},
		},
		"disable-security-hub-v2": {
			Name:   "disable-security-hub-v2",
			Fields: fields_disable_security_hub_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSecurityHubV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_security_hub_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSecurityHubV2(ctx, input)
			},
		},
		"disassociate-from-administrator-account": {
			Name:   "disassociate-from-administrator-account",
			Fields: fields_disassociate_from_administrator_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFromAdministratorAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_from_administrator_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFromAdministratorAccount(ctx, input)
			},
		},
		"disassociate-from-master-account": {
			Name:   "disassociate-from-master-account",
			Fields: fields_disassociate_from_master_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFromMasterAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_from_master_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFromMasterAccount(ctx, input)
			},
		},
		"disassociate-members": {
			Name:   "disassociate-members",
			Fields: fields_disassociate_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMembers(ctx, input)
			},
		},
		"enable-import-findings-for-product": {
			Name:   "enable-import-findings-for-product",
			Fields: fields_enable_import_findings_for_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableImportFindingsForProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_import_findings_for_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableImportFindingsForProduct(ctx, input)
			},
		},
		"enable-organization-admin-account": {
			Name:   "enable-organization-admin-account",
			Fields: fields_enable_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableOrganizationAdminAccount(ctx, input)
			},
		},
		"enable-security-hub": {
			Name:   "enable-security-hub",
			Fields: fields_enable_security_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSecurityHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_security_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSecurityHub(ctx, input)
			},
		},
		"enable-security-hub-v2": {
			Name:   "enable-security-hub-v2",
			Fields: fields_enable_security_hub_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSecurityHubV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_security_hub_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSecurityHubV2(ctx, input)
			},
		},
		"get-administrator-account": {
			Name:   "get-administrator-account",
			Fields: fields_get_administrator_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAdministratorAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_administrator_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAdministratorAccount(ctx, input)
			},
		},
		"get-aggregator-v2": {
			Name:   "get-aggregator-v2",
			Fields: fields_get_aggregator_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAggregatorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_get_aggregator_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAggregatorV2(ctx, input)
			},
		},
		"get-automation-rule-v2": {
			Name:   "get-automation-rule-v2",
			Fields: fields_get_automation_rule_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomationRuleV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automation_rule_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomationRuleV2(ctx, input)
			},
		},
		"get-configuration-policy": {
			Name:   "get-configuration-policy",
			Fields: fields_get_configuration_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationPolicy(ctx, input)
			},
		},
		"get-configuration-policy-association": {
			Name:   "get-configuration-policy-association",
			Fields: fields_get_configuration_policy_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationPolicyAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_policy_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationPolicyAssociation(ctx, input)
			},
		},
		"get-connector-v2": {
			Name:   "get-connector-v2",
			Fields: fields_get_connector_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connector_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectorV2(ctx, input)
			},
		},
		"get-enabled-standards": {
			Name:   "get-enabled-standards",
			Fields: fields_get_enabled_standards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnabledStandardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_enabled_standards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEnabledStandards(ctx, input)
				}
				var results []*svc.GetEnabledStandardsOutput
				p := svc.NewGetEnabledStandardsPaginator(client, input)
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
		"get-finding-aggregator": {
			Name:   "get-finding-aggregator",
			Fields: fields_get_finding_aggregator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingAggregatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_finding_aggregator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingAggregator(ctx, input)
			},
		},
		"get-finding-history": {
			Name:   "get-finding-history",
			Fields: fields_get_finding_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_finding_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindingHistory(ctx, input)
				}
				var results []*svc.GetFindingHistoryOutput
				p := svc.NewGetFindingHistoryPaginator(client, input)
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
		"get-finding-statistics-v2": {
			Name:   "get-finding-statistics-v2",
			Fields: fields_get_finding_statistics_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingStatisticsV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_get_finding_statistics_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingStatisticsV2(ctx, input)
			},
		},
		"get-findings": {
			Name:   "get-findings",
			Fields: fields_get_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindings(ctx, input)
				}
				var results []*svc.GetFindingsOutput
				p := svc.NewGetFindingsPaginator(client, input)
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
		"get-findings-trends-v2": {
			Name:   "get-findings-trends-v2",
			Fields: fields_get_findings_trends_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsTrendsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_findings_trends_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindingsTrendsV2(ctx, input)
				}
				var results []*svc.GetFindingsTrendsV2Output
				p := svc.NewGetFindingsTrendsV2Paginator(client, input)
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
		"get-findings-v2": {
			Name:   "get-findings-v2",
			Fields: fields_get_findings_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_findings_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindingsV2(ctx, input)
				}
				var results []*svc.GetFindingsV2Output
				p := svc.NewGetFindingsV2Paginator(client, input)
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
		"get-insight-results": {
			Name:   "get-insight-results",
			Fields: fields_get_insight_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insight_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsightResults(ctx, input)
			},
		},
		"get-insights": {
			Name:   "get-insights",
			Fields: fields_get_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInsights(ctx, input)
				}
				var results []*svc.GetInsightsOutput
				p := svc.NewGetInsightsPaginator(client, input)
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
		"get-invitations-count": {
			Name:   "get-invitations-count",
			Fields: fields_get_invitations_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvitationsCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_invitations_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvitationsCount(ctx, input)
			},
		},
		"get-master-account": {
			Name:   "get-master-account",
			Fields: fields_get_master_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMasterAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_master_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMasterAccount(ctx, input)
			},
		},
		"get-members": {
			Name:   "get-members",
			Fields: fields_get_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMembers(ctx, input)
			},
		},
		"get-resources-statistics-v2": {
			Name:   "get-resources-statistics-v2",
			Fields: fields_get_resources_statistics_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesStatisticsV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resources_statistics_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcesStatisticsV2(ctx, input)
			},
		},
		"get-resources-trends-v2": {
			Name:   "get-resources-trends-v2",
			Fields: fields_get_resources_trends_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesTrendsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resources_trends_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourcesTrendsV2(ctx, input)
				}
				var results []*svc.GetResourcesTrendsV2Output
				p := svc.NewGetResourcesTrendsV2Paginator(client, input)
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
		"get-resources-v2": {
			Name:   "get-resources-v2",
			Fields: fields_get_resources_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resources_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourcesV2(ctx, input)
				}
				var results []*svc.GetResourcesV2Output
				p := svc.NewGetResourcesV2Paginator(client, input)
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
		"get-security-control-definition": {
			Name:   "get-security-control-definition",
			Fields: fields_get_security_control_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityControlDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_security_control_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSecurityControlDefinition(ctx, input)
			},
		},
		"invite-members": {
			Name:   "invite-members",
			Fields: fields_invite_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InviteMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invite_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InviteMembers(ctx, input)
			},
		},
		"list-aggregators-v2": {
			Name:   "list-aggregators-v2",
			Fields: fields_list_aggregators_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAggregatorsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aggregators_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAggregatorsV2(ctx, input)
				}
				var results []*svc.ListAggregatorsV2Output
				p := svc.NewListAggregatorsV2Paginator(client, input)
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
		"list-automation-rules": {
			Name:   "list-automation-rules",
			Fields: fields_list_automation_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_automation_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAutomationRules(ctx, input)
			},
		},
		"list-automation-rules-v2": {
			Name:   "list-automation-rules-v2",
			Fields: fields_list_automation_rules_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationRulesV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_list_automation_rules_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAutomationRulesV2(ctx, input)
			},
		},
		"list-configuration-policies": {
			Name:   "list-configuration-policies",
			Fields: fields_list_configuration_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationPolicies(ctx, input)
				}
				var results []*svc.ListConfigurationPoliciesOutput
				p := svc.NewListConfigurationPoliciesPaginator(client, input)
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
		"list-configuration-policy-associations": {
			Name:   "list-configuration-policy-associations",
			Fields: fields_list_configuration_policy_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationPolicyAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_policy_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationPolicyAssociations(ctx, input)
				}
				var results []*svc.ListConfigurationPolicyAssociationsOutput
				p := svc.NewListConfigurationPolicyAssociationsPaginator(client, input)
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
		"list-connectors-v2": {
			Name:   "list-connectors-v2",
			Fields: fields_list_connectors_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorsV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_list_connectors_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConnectorsV2(ctx, input)
			},
		},
		"list-enabled-products-for-import": {
			Name:   "list-enabled-products-for-import",
			Fields: fields_list_enabled_products_for_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnabledProductsForImportInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_enabled_products_for_import, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnabledProductsForImport(ctx, input)
				}
				var results []*svc.ListEnabledProductsForImportOutput
				p := svc.NewListEnabledProductsForImportPaginator(client, input)
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
		"list-finding-aggregators": {
			Name:   "list-finding-aggregators",
			Fields: fields_list_finding_aggregators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingAggregatorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_finding_aggregators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindingAggregators(ctx, input)
				}
				var results []*svc.ListFindingAggregatorsOutput
				p := svc.NewListFindingAggregatorsPaginator(client, input)
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
		"list-invitations": {
			Name:   "list-invitations",
			Fields: fields_list_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvitations(ctx, input)
				}
				var results []*svc.ListInvitationsOutput
				p := svc.NewListInvitationsPaginator(client, input)
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
		"list-organization-admin-accounts": {
			Name:   "list-organization-admin-accounts",
			Fields: fields_list_organization_admin_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationAdminAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_admin_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationAdminAccounts(ctx, input)
				}
				var results []*svc.ListOrganizationAdminAccountsOutput
				p := svc.NewListOrganizationAdminAccountsPaginator(client, input)
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
		"list-security-control-definitions": {
			Name:   "list-security-control-definitions",
			Fields: fields_list_security_control_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityControlDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_control_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityControlDefinitions(ctx, input)
				}
				var results []*svc.ListSecurityControlDefinitionsOutput
				p := svc.NewListSecurityControlDefinitionsPaginator(client, input)
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
		"list-standards-control-associations": {
			Name:   "list-standards-control-associations",
			Fields: fields_list_standards_control_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStandardsControlAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_standards_control_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStandardsControlAssociations(ctx, input)
				}
				var results []*svc.ListStandardsControlAssociationsOutput
				p := svc.NewListStandardsControlAssociationsPaginator(client, input)
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
		"register-connector-v2": {
			Name:   "register-connector-v2",
			Fields: fields_register_connector_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterConnectorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_register_connector_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterConnectorV2(ctx, input)
			},
		},
		"start-configuration-policy-association": {
			Name:   "start-configuration-policy-association",
			Fields: fields_start_configuration_policy_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConfigurationPolicyAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_configuration_policy_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConfigurationPolicyAssociation(ctx, input)
			},
		},
		"start-configuration-policy-disassociation": {
			Name:   "start-configuration-policy-disassociation",
			Fields: fields_start_configuration_policy_disassociation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConfigurationPolicyDisassociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_configuration_policy_disassociation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConfigurationPolicyDisassociation(ctx, input)
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
		"update-action-target": {
			Name:   "update-action-target",
			Fields: fields_update_action_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_action_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateActionTarget(ctx, input)
			},
		},
		"update-aggregator-v2": {
			Name:   "update-aggregator-v2",
			Fields: fields_update_aggregator_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAggregatorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_update_aggregator_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAggregatorV2(ctx, input)
			},
		},
		"update-automation-rule-v2": {
			Name:   "update-automation-rule-v2",
			Fields: fields_update_automation_rule_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomationRuleV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automation_rule_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomationRuleV2(ctx, input)
			},
		},
		"update-configuration-policy": {
			Name:   "update-configuration-policy",
			Fields: fields_update_configuration_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationPolicy(ctx, input)
			},
		},
		"update-connector-v2": {
			Name:   "update-connector-v2",
			Fields: fields_update_connector_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectorV2(ctx, input)
			},
		},
		"update-finding-aggregator": {
			Name:   "update-finding-aggregator",
			Fields: fields_update_finding_aggregator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFindingAggregatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_finding_aggregator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFindingAggregator(ctx, input)
			},
		},
		"update-findings": {
			Name:   "update-findings",
			Fields: fields_update_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFindings(ctx, input)
			},
		},
		"update-insight": {
			Name:   "update-insight",
			Fields: fields_update_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInsight(ctx, input)
			},
		},
		"update-organization-configuration": {
			Name:   "update-organization-configuration",
			Fields: fields_update_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOrganizationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_organization_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOrganizationConfiguration(ctx, input)
			},
		},
		"update-security-control": {
			Name:   "update-security-control",
			Fields: fields_update_security_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityControl(ctx, input)
			},
		},
		"update-security-hub-configuration": {
			Name:   "update-security-hub-configuration",
			Fields: fields_update_security_hub_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityHubConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_hub_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityHubConfiguration(ctx, input)
			},
		},
		"update-standards-control": {
			Name:   "update-standards-control",
			Fields: fields_update_standards_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStandardsControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_standards_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStandardsControl(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("securityhub", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
