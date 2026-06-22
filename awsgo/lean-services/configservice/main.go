package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/configservice"
)

var fields_associate_resource_types = []leanruntime.Field{
	{Name: "ConfigurationRecorderArn", Flag: "configuration-recorder-arn", Type: "*string", Required: true},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceType", Required: true},
}

var fields_batch_get_aggregate_resource_config = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "ResourceIdentifiers", Flag: "resource-identifiers", Type: "[]types.AggregateResourceIdentifier", Required: true},
}

var fields_batch_get_resource_config = []leanruntime.Field{
	{Name: "ResourceKeys", Flag: "resource-keys", Type: "[]types.ResourceKey", Required: true},
}

var fields_delete_aggregation_authorization = []leanruntime.Field{
	{Name: "AuthorizedAccountId", Flag: "authorized-account-id", Type: "*string", Required: true},
	{Name: "AuthorizedAwsRegion", Flag: "authorized-aws-region", Type: "*string", Required: true},
}

var fields_delete_config_rule = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
}

var fields_delete_configuration_aggregator = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
}

var fields_delete_configuration_recorder = []leanruntime.Field{
	{Name: "ConfigurationRecorderName", Flag: "configuration-recorder-name", Type: "*string", Required: true},
}

var fields_delete_conformance_pack = []leanruntime.Field{
	{Name: "ConformancePackName", Flag: "conformance-pack-name", Type: "*string", Required: true},
}

var fields_delete_delivery_channel = []leanruntime.Field{
	{Name: "DeliveryChannelName", Flag: "delivery-channel-name", Type: "*string", Required: true},
}

var fields_delete_evaluation_results = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
}

var fields_delete_organization_config_rule = []leanruntime.Field{
	{Name: "OrganizationConfigRuleName", Flag: "organization-config-rule-name", Type: "*string", Required: true},
}

var fields_delete_organization_conformance_pack = []leanruntime.Field{
	{Name: "OrganizationConformancePackName", Flag: "organization-conformance-pack-name", Type: "*string", Required: true},
}

var fields_delete_pending_aggregation_request = []leanruntime.Field{
	{Name: "RequesterAccountId", Flag: "requester-account-id", Type: "*string", Required: true},
	{Name: "RequesterAwsRegion", Flag: "requester-aws-region", Type: "*string", Required: true},
}

var fields_delete_remediation_configuration = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_delete_remediation_exceptions = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "ResourceKeys", Flag: "resource-keys", Type: "[]types.RemediationExceptionResourceKey", Required: true},
}

var fields_delete_resource_config = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_delete_retention_configuration = []leanruntime.Field{
	{Name: "RetentionConfigurationName", Flag: "retention-configuration-name", Type: "*string", Required: true},
}

var fields_delete_service_linked_configuration_recorder = []leanruntime.Field{
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: true},
}

var fields_delete_stored_query = []leanruntime.Field{
	{Name: "QueryName", Flag: "query-name", Type: "*string", Required: true},
}

var fields_deliver_config_snapshot = []leanruntime.Field{
	{Name: "DeliveryChannelName", Flag: "delivery-channel-name", Type: "*string", Required: true},
}

var fields_describe_aggregate_compliance_by_config_rules = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ConfigRuleComplianceFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_aggregate_compliance_by_conformance_packs = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.AggregateConformancePackComplianceFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_aggregation_authorizations = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_compliance_by_config_rule = []leanruntime.Field{
	{Name: "ComplianceTypes", Flag: "compliance-types", Type: "[]types.ComplianceType", Required: false},
	{Name: "ConfigRuleNames", Flag: "config-rule-names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_compliance_by_resource = []leanruntime.Field{
	{Name: "ComplianceTypes", Flag: "compliance-types", Type: "[]types.ComplianceType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_describe_config_rule_evaluation_status = []leanruntime.Field{
	{Name: "ConfigRuleNames", Flag: "config-rule-names", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_config_rules = []leanruntime.Field{
	{Name: "ConfigRuleNames", Flag: "config-rule-names", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeConfigRulesFilters", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_configuration_aggregator_sources_status = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UpdateStatus", Flag: "update-status", Type: "[]types.AggregatedSourceStatusType", Required: false},
}

var fields_describe_configuration_aggregators = []leanruntime.Field{
	{Name: "ConfigurationAggregatorNames", Flag: "configuration-aggregator-names", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_configuration_recorder_status = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "ConfigurationRecorderNames", Flag: "configuration-recorder-names", Type: "[]string", Required: false},
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: false},
}

var fields_describe_configuration_recorders = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "ConfigurationRecorderNames", Flag: "configuration-recorder-names", Type: "[]string", Required: false},
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: false},
}

var fields_describe_conformance_pack_compliance = []leanruntime.Field{
	{Name: "ConformancePackName", Flag: "conformance-pack-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ConformancePackComplianceFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_conformance_pack_status = []leanruntime.Field{
	{Name: "ConformancePackNames", Flag: "conformance-pack-names", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_conformance_packs = []leanruntime.Field{
	{Name: "ConformancePackNames", Flag: "conformance-pack-names", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_delivery_channel_status = []leanruntime.Field{
	{Name: "DeliveryChannelNames", Flag: "delivery-channel-names", Type: "[]string", Required: false},
}

var fields_describe_delivery_channels = []leanruntime.Field{
	{Name: "DeliveryChannelNames", Flag: "delivery-channel-names", Type: "[]string", Required: false},
}

var fields_describe_organization_config_rule_statuses = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationConfigRuleNames", Flag: "organization-config-rule-names", Type: "[]string", Required: false},
}

var fields_describe_organization_config_rules = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationConfigRuleNames", Flag: "organization-config-rule-names", Type: "[]string", Required: false},
}

var fields_describe_organization_conformance_pack_statuses = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationConformancePackNames", Flag: "organization-conformance-pack-names", Type: "[]string", Required: false},
}

var fields_describe_organization_conformance_packs = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationConformancePackNames", Flag: "organization-conformance-pack-names", Type: "[]string", Required: false},
}

var fields_describe_pending_aggregation_requests = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_remediation_configurations = []leanruntime.Field{
	{Name: "ConfigRuleNames", Flag: "config-rule-names", Type: "[]string", Required: true},
}

var fields_describe_remediation_exceptions = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceKeys", Flag: "resource-keys", Type: "[]types.RemediationExceptionResourceKey", Required: false},
}

var fields_describe_remediation_execution_status = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceKeys", Flag: "resource-keys", Type: "[]types.ResourceKey", Required: false},
}

var fields_describe_retention_configurations = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RetentionConfigurationNames", Flag: "retention-configuration-names", Type: "[]string", Required: false},
}

var fields_disassociate_resource_types = []leanruntime.Field{
	{Name: "ConfigurationRecorderArn", Flag: "configuration-recorder-arn", Type: "*string", Required: true},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceType", Required: true},
}

var fields_get_aggregate_compliance_details_by_config_rule = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: true},
	{Name: "ComplianceType", Flag: "compliance-type", Type: "types.ComplianceType", Required: false},
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_aggregate_config_rule_compliance_summary = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ConfigRuleComplianceSummaryFilters", Required: false},
	{Name: "GroupByKey", Flag: "group-by-key", Type: "types.ConfigRuleComplianceSummaryGroupKey", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_aggregate_conformance_pack_compliance_summary = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.AggregateConformancePackComplianceSummaryFilters", Required: false},
	{Name: "GroupByKey", Flag: "group-by-key", Type: "types.AggregateConformancePackComplianceSummaryGroupKey", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_aggregate_discovered_resource_counts = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ResourceCountFilters", Required: false},
	{Name: "GroupByKey", Flag: "group-by-key", Type: "types.ResourceCountGroupKey", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_aggregate_resource_config = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.AggregateResourceIdentifier", Required: true},
}

var fields_get_compliance_details_by_config_rule = []leanruntime.Field{
	{Name: "ComplianceTypes", Flag: "compliance-types", Type: "[]types.ComplianceType", Required: false},
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_compliance_details_by_resource = []leanruntime.Field{
	{Name: "ComplianceTypes", Flag: "compliance-types", Type: "[]types.ComplianceType", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceEvaluationId", Flag: "resource-evaluation-id", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_get_compliance_summary_by_config_rule = []leanruntime.Field{}

var fields_get_compliance_summary_by_resource_type = []leanruntime.Field{
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
}

var fields_get_conformance_pack_compliance_details = []leanruntime.Field{
	{Name: "ConformancePackName", Flag: "conformance-pack-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ConformancePackEvaluationFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_conformance_pack_compliance_summary = []leanruntime.Field{
	{Name: "ConformancePackNames", Flag: "conformance-pack-names", Type: "[]string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_custom_rule_policy = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: false},
}

var fields_get_discovered_resource_counts = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
}

var fields_get_organization_config_rule_detailed_status = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.StatusDetailFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationConfigRuleName", Flag: "organization-config-rule-name", Type: "*string", Required: true},
}

var fields_get_organization_conformance_pack_detailed_status = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.OrganizationResourceDetailedStatusFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationConformancePackName", Flag: "organization-conformance-pack-name", Type: "*string", Required: true},
}

var fields_get_organization_custom_rule_policy = []leanruntime.Field{
	{Name: "OrganizationConfigRuleName", Flag: "organization-config-rule-name", Type: "*string", Required: true},
}

var fields_get_resource_config_history = []leanruntime.Field{
	{Name: "ChronologicalOrder", Flag: "chronological-order", Type: "types.ChronologicalOrder", Required: false},
	{Name: "EarlierTime", Flag: "earlier-time", Type: "*time.Time", Required: false},
	{Name: "LaterTime", Flag: "later-time", Type: "*time.Time", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_get_resource_evaluation_summary = []leanruntime.Field{
	{Name: "ResourceEvaluationId", Flag: "resource-evaluation-id", Type: "*string", Required: true},
}

var fields_get_stored_query = []leanruntime.Field{
	{Name: "QueryName", Flag: "query-name", Type: "*string", Required: true},
}

var fields_list_aggregate_discovered_resources = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ResourceFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_list_configuration_recorders = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ConfigurationRecorderFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_conformance_pack_compliance_scores = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ConformancePackComplianceScoresFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_discovered_resources = []leanruntime.Field{
	{Name: "IncludeDeletedResources", Flag: "include-deleted-resources", Type: "bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_list_resource_evaluations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ResourceEvaluationFilters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stored_queries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_aggregation_authorization = []leanruntime.Field{
	{Name: "AuthorizedAccountId", Flag: "authorized-account-id", Type: "*string", Required: true},
	{Name: "AuthorizedAwsRegion", Flag: "authorized-aws-region", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_config_rule = []leanruntime.Field{
	{Name: "ConfigRule", Flag: "config-rule", Type: "*types.ConfigRule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_configuration_aggregator = []leanruntime.Field{
	{Name: "AccountAggregationSources", Flag: "account-aggregation-sources", Type: "[]types.AccountAggregationSource", Required: false},
	{Name: "AggregatorFilters", Flag: "aggregator-filters", Type: "*types.AggregatorFilters", Required: false},
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "OrganizationAggregationSource", Flag: "organization-aggregation-source", Type: "*types.OrganizationAggregationSource", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_configuration_recorder = []leanruntime.Field{
	{Name: "ConfigurationRecorder", Flag: "configuration-recorder", Type: "*types.ConfigurationRecorder", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_conformance_pack = []leanruntime.Field{
	{Name: "ConformancePackInputParameters", Flag: "conformance-pack-input-parameters", Type: "[]types.ConformancePackInputParameter", Required: false},
	{Name: "ConformancePackName", Flag: "conformance-pack-name", Type: "*string", Required: true},
	{Name: "DeliveryS3Bucket", Flag: "delivery-s3-bucket", Type: "*string", Required: false},
	{Name: "DeliveryS3KeyPrefix", Flag: "delivery-s3-key-prefix", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateS3Uri", Flag: "template-s3-uri", Type: "*string", Required: false},
	{Name: "TemplateSSMDocumentDetails", Flag: "template-ssm-document-details", Type: "*types.TemplateSSMDocumentDetails", Required: false},
}

var fields_put_delivery_channel = []leanruntime.Field{
	{Name: "DeliveryChannel", Flag: "delivery-channel", Type: "*types.DeliveryChannel", Required: true},
}

var fields_put_evaluations = []leanruntime.Field{
	{Name: "Evaluations", Flag: "evaluations", Type: "[]types.Evaluation", Required: false},
	{Name: "ResultToken", Flag: "result-token", Type: "*string", Required: true},
	{Name: "TestMode", Flag: "test-mode", Type: "bool", Required: false},
}

var fields_put_external_evaluation = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "ExternalEvaluation", Flag: "external-evaluation", Type: "*types.ExternalEvaluation", Required: true},
}

var fields_put_organization_config_rule = []leanruntime.Field{
	{Name: "ExcludedAccounts", Flag: "excluded-accounts", Type: "[]string", Required: false},
	{Name: "OrganizationConfigRuleName", Flag: "organization-config-rule-name", Type: "*string", Required: true},
	{Name: "OrganizationCustomPolicyRuleMetadata", Flag: "organization-custom-policy-rule-metadata", Type: "*types.OrganizationCustomPolicyRuleMetadata", Required: false},
	{Name: "OrganizationCustomRuleMetadata", Flag: "organization-custom-rule-metadata", Type: "*types.OrganizationCustomRuleMetadata", Required: false},
	{Name: "OrganizationManagedRuleMetadata", Flag: "organization-managed-rule-metadata", Type: "*types.OrganizationManagedRuleMetadata", Required: false},
}

var fields_put_organization_conformance_pack = []leanruntime.Field{
	{Name: "ConformancePackInputParameters", Flag: "conformance-pack-input-parameters", Type: "[]types.ConformancePackInputParameter", Required: false},
	{Name: "DeliveryS3Bucket", Flag: "delivery-s3-bucket", Type: "*string", Required: false},
	{Name: "DeliveryS3KeyPrefix", Flag: "delivery-s3-key-prefix", Type: "*string", Required: false},
	{Name: "ExcludedAccounts", Flag: "excluded-accounts", Type: "[]string", Required: false},
	{Name: "OrganizationConformancePackName", Flag: "organization-conformance-pack-name", Type: "*string", Required: true},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateS3Uri", Flag: "template-s3-uri", Type: "*string", Required: false},
}

var fields_put_remediation_configurations = []leanruntime.Field{
	{Name: "RemediationConfigurations", Flag: "remediation-configurations", Type: "[]types.RemediationConfiguration", Required: true},
}

var fields_put_remediation_exceptions = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "ExpirationTime", Flag: "expiration-time", Type: "*time.Time", Required: false},
	{Name: "Message", Flag: "message", Type: "*string", Required: false},
	{Name: "ResourceKeys", Flag: "resource-keys", Type: "[]types.RemediationExceptionResourceKey", Required: true},
}

var fields_put_resource_config = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
	{Name: "SchemaVersionId", Flag: "schema-version-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_put_retention_configuration = []leanruntime.Field{
	{Name: "RetentionPeriodInDays", Flag: "retention-period-in-days", Type: "*int32", Required: true},
}

var fields_put_service_linked_configuration_recorder = []leanruntime.Field{
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_stored_query = []leanruntime.Field{
	{Name: "StoredQuery", Flag: "stored-query", Type: "*types.StoredQuery", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_select_aggregate_resource_config = []leanruntime.Field{
	{Name: "ConfigurationAggregatorName", Flag: "configuration-aggregator-name", Type: "*string", Required: true},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_select_resource_config = []leanruntime.Field{
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_config_rules_evaluation = []leanruntime.Field{
	{Name: "ConfigRuleNames", Flag: "config-rule-names", Type: "[]string", Required: false},
}

var fields_start_configuration_recorder = []leanruntime.Field{
	{Name: "ConfigurationRecorderName", Flag: "configuration-recorder-name", Type: "*string", Required: true},
}

var fields_start_remediation_execution = []leanruntime.Field{
	{Name: "ConfigRuleName", Flag: "config-rule-name", Type: "*string", Required: true},
	{Name: "ResourceKeys", Flag: "resource-keys", Type: "[]types.ResourceKey", Required: true},
}

var fields_start_resource_evaluation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EvaluationContext", Flag: "evaluation-context", Type: "*types.EvaluationContext", Required: false},
	{Name: "EvaluationMode", Flag: "evaluation-mode", Type: "types.EvaluationMode", Required: true},
	{Name: "EvaluationTimeout", Flag: "evaluation-timeout", Type: "int32", Required: false},
	{Name: "ResourceDetails", Flag: "resource-details", Type: "*types.ResourceDetails", Required: true},
}

var fields_stop_configuration_recorder = []leanruntime.Field{
	{Name: "ConfigurationRecorderName", Flag: "configuration-recorder-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-resource-types": {
			Name:   "associate-resource-types",
			Fields: fields_associate_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResourceTypes(ctx, input)
			},
		},
		"batch-get-aggregate-resource-config": {
			Name:   "batch-get-aggregate-resource-config",
			Fields: fields_batch_get_aggregate_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAggregateResourceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_aggregate_resource_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetAggregateResourceConfig(ctx, input)
			},
		},
		"batch-get-resource-config": {
			Name:   "batch-get-resource-config",
			Fields: fields_batch_get_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetResourceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_resource_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetResourceConfig(ctx, input)
			},
		},
		"delete-aggregation-authorization": {
			Name:   "delete-aggregation-authorization",
			Fields: fields_delete_aggregation_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAggregationAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_aggregation_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAggregationAuthorization(ctx, input)
			},
		},
		"delete-config-rule": {
			Name:   "delete-config-rule",
			Fields: fields_delete_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_config_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigRule(ctx, input)
			},
		},
		"delete-configuration-aggregator": {
			Name:   "delete-configuration-aggregator",
			Fields: fields_delete_configuration_aggregator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationAggregatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_aggregator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationAggregator(ctx, input)
			},
		},
		"delete-configuration-recorder": {
			Name:   "delete-configuration-recorder",
			Fields: fields_delete_configuration_recorder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationRecorderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_recorder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationRecorder(ctx, input)
			},
		},
		"delete-conformance-pack": {
			Name:   "delete-conformance-pack",
			Fields: fields_delete_conformance_pack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConformancePackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_conformance_pack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConformancePack(ctx, input)
			},
		},
		"delete-delivery-channel": {
			Name:   "delete-delivery-channel",
			Fields: fields_delete_delivery_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeliveryChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_delivery_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeliveryChannel(ctx, input)
			},
		},
		"delete-evaluation-results": {
			Name:   "delete-evaluation-results",
			Fields: fields_delete_evaluation_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEvaluationResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_evaluation_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEvaluationResults(ctx, input)
			},
		},
		"delete-organization-config-rule": {
			Name:   "delete-organization-config-rule",
			Fields: fields_delete_organization_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOrganizationConfigRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_organization_config_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOrganizationConfigRule(ctx, input)
			},
		},
		"delete-organization-conformance-pack": {
			Name:   "delete-organization-conformance-pack",
			Fields: fields_delete_organization_conformance_pack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOrganizationConformancePackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_organization_conformance_pack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOrganizationConformancePack(ctx, input)
			},
		},
		"delete-pending-aggregation-request": {
			Name:   "delete-pending-aggregation-request",
			Fields: fields_delete_pending_aggregation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePendingAggregationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pending_aggregation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePendingAggregationRequest(ctx, input)
			},
		},
		"delete-remediation-configuration": {
			Name:   "delete-remediation-configuration",
			Fields: fields_delete_remediation_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRemediationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_remediation_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRemediationConfiguration(ctx, input)
			},
		},
		"delete-remediation-exceptions": {
			Name:   "delete-remediation-exceptions",
			Fields: fields_delete_remediation_exceptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRemediationExceptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_remediation_exceptions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRemediationExceptions(ctx, input)
			},
		},
		"delete-resource-config": {
			Name:   "delete-resource-config",
			Fields: fields_delete_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceConfig(ctx, input)
			},
		},
		"delete-retention-configuration": {
			Name:   "delete-retention-configuration",
			Fields: fields_delete_retention_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRetentionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_retention_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRetentionConfiguration(ctx, input)
			},
		},
		"delete-service-linked-configuration-recorder": {
			Name:   "delete-service-linked-configuration-recorder",
			Fields: fields_delete_service_linked_configuration_recorder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceLinkedConfigurationRecorderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_linked_configuration_recorder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceLinkedConfigurationRecorder(ctx, input)
			},
		},
		"delete-stored-query": {
			Name:   "delete-stored-query",
			Fields: fields_delete_stored_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStoredQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stored_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStoredQuery(ctx, input)
			},
		},
		"deliver-config-snapshot": {
			Name:   "deliver-config-snapshot",
			Fields: fields_deliver_config_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeliverConfigSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deliver_config_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeliverConfigSnapshot(ctx, input)
			},
		},
		"describe-aggregate-compliance-by-config-rules": {
			Name:   "describe-aggregate-compliance-by-config-rules",
			Fields: fields_describe_aggregate_compliance_by_config_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAggregateComplianceByConfigRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_aggregate_compliance_by_config_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAggregateComplianceByConfigRules(ctx, input)
				}
				var results []*svc.DescribeAggregateComplianceByConfigRulesOutput
				p := svc.NewDescribeAggregateComplianceByConfigRulesPaginator(client, input)
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
		"describe-aggregate-compliance-by-conformance-packs": {
			Name:   "describe-aggregate-compliance-by-conformance-packs",
			Fields: fields_describe_aggregate_compliance_by_conformance_packs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAggregateComplianceByConformancePacksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_aggregate_compliance_by_conformance_packs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAggregateComplianceByConformancePacks(ctx, input)
				}
				var results []*svc.DescribeAggregateComplianceByConformancePacksOutput
				p := svc.NewDescribeAggregateComplianceByConformancePacksPaginator(client, input)
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
		"describe-aggregation-authorizations": {
			Name:   "describe-aggregation-authorizations",
			Fields: fields_describe_aggregation_authorizations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAggregationAuthorizationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_aggregation_authorizations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAggregationAuthorizations(ctx, input)
				}
				var results []*svc.DescribeAggregationAuthorizationsOutput
				p := svc.NewDescribeAggregationAuthorizationsPaginator(client, input)
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
		"describe-compliance-by-config-rule": {
			Name:   "describe-compliance-by-config-rule",
			Fields: fields_describe_compliance_by_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComplianceByConfigRuleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_compliance_by_config_rule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeComplianceByConfigRule(ctx, input)
				}
				var results []*svc.DescribeComplianceByConfigRuleOutput
				p := svc.NewDescribeComplianceByConfigRulePaginator(client, input)
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
		"describe-compliance-by-resource": {
			Name:   "describe-compliance-by-resource",
			Fields: fields_describe_compliance_by_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComplianceByResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_compliance_by_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeComplianceByResource(ctx, input)
				}
				var results []*svc.DescribeComplianceByResourceOutput
				p := svc.NewDescribeComplianceByResourcePaginator(client, input)
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
		"describe-config-rule-evaluation-status": {
			Name:   "describe-config-rule-evaluation-status",
			Fields: fields_describe_config_rule_evaluation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigRuleEvaluationStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_config_rule_evaluation_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConfigRuleEvaluationStatus(ctx, input)
				}
				var results []*svc.DescribeConfigRuleEvaluationStatusOutput
				p := svc.NewDescribeConfigRuleEvaluationStatusPaginator(client, input)
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
		"describe-config-rules": {
			Name:   "describe-config-rules",
			Fields: fields_describe_config_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_config_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConfigRules(ctx, input)
				}
				var results []*svc.DescribeConfigRulesOutput
				p := svc.NewDescribeConfigRulesPaginator(client, input)
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
		"describe-configuration-aggregator-sources-status": {
			Name:   "describe-configuration-aggregator-sources-status",
			Fields: fields_describe_configuration_aggregator_sources_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationAggregatorSourcesStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_configuration_aggregator_sources_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConfigurationAggregatorSourcesStatus(ctx, input)
				}
				var results []*svc.DescribeConfigurationAggregatorSourcesStatusOutput
				p := svc.NewDescribeConfigurationAggregatorSourcesStatusPaginator(client, input)
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
		"describe-configuration-aggregators": {
			Name:   "describe-configuration-aggregators",
			Fields: fields_describe_configuration_aggregators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationAggregatorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_configuration_aggregators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConfigurationAggregators(ctx, input)
				}
				var results []*svc.DescribeConfigurationAggregatorsOutput
				p := svc.NewDescribeConfigurationAggregatorsPaginator(client, input)
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
		"describe-configuration-recorder-status": {
			Name:   "describe-configuration-recorder-status",
			Fields: fields_describe_configuration_recorder_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationRecorderStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_recorder_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationRecorderStatus(ctx, input)
			},
		},
		"describe-configuration-recorders": {
			Name:   "describe-configuration-recorders",
			Fields: fields_describe_configuration_recorders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationRecordersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_recorders, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationRecorders(ctx, input)
			},
		},
		"describe-conformance-pack-compliance": {
			Name:   "describe-conformance-pack-compliance",
			Fields: fields_describe_conformance_pack_compliance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConformancePackComplianceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_conformance_pack_compliance, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConformancePackCompliance(ctx, input)
				}
				var results []*svc.DescribeConformancePackComplianceOutput
				p := svc.NewDescribeConformancePackCompliancePaginator(client, input)
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
		"describe-conformance-pack-status": {
			Name:   "describe-conformance-pack-status",
			Fields: fields_describe_conformance_pack_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConformancePackStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_conformance_pack_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConformancePackStatus(ctx, input)
				}
				var results []*svc.DescribeConformancePackStatusOutput
				p := svc.NewDescribeConformancePackStatusPaginator(client, input)
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
		"describe-conformance-packs": {
			Name:   "describe-conformance-packs",
			Fields: fields_describe_conformance_packs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConformancePacksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_conformance_packs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConformancePacks(ctx, input)
				}
				var results []*svc.DescribeConformancePacksOutput
				p := svc.NewDescribeConformancePacksPaginator(client, input)
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
		"describe-delivery-channel-status": {
			Name:   "describe-delivery-channel-status",
			Fields: fields_describe_delivery_channel_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeliveryChannelStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_delivery_channel_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeliveryChannelStatus(ctx, input)
			},
		},
		"describe-delivery-channels": {
			Name:   "describe-delivery-channels",
			Fields: fields_describe_delivery_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeliveryChannelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_delivery_channels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeliveryChannels(ctx, input)
			},
		},
		"describe-organization-config-rule-statuses": {
			Name:   "describe-organization-config-rule-statuses",
			Fields: fields_describe_organization_config_rule_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConfigRuleStatusesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_organization_config_rule_statuses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrganizationConfigRuleStatuses(ctx, input)
				}
				var results []*svc.DescribeOrganizationConfigRuleStatusesOutput
				p := svc.NewDescribeOrganizationConfigRuleStatusesPaginator(client, input)
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
		"describe-organization-config-rules": {
			Name:   "describe-organization-config-rules",
			Fields: fields_describe_organization_config_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConfigRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_organization_config_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrganizationConfigRules(ctx, input)
				}
				var results []*svc.DescribeOrganizationConfigRulesOutput
				p := svc.NewDescribeOrganizationConfigRulesPaginator(client, input)
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
		"describe-organization-conformance-pack-statuses": {
			Name:   "describe-organization-conformance-pack-statuses",
			Fields: fields_describe_organization_conformance_pack_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConformancePackStatusesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_organization_conformance_pack_statuses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrganizationConformancePackStatuses(ctx, input)
				}
				var results []*svc.DescribeOrganizationConformancePackStatusesOutput
				p := svc.NewDescribeOrganizationConformancePackStatusesPaginator(client, input)
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
		"describe-organization-conformance-packs": {
			Name:   "describe-organization-conformance-packs",
			Fields: fields_describe_organization_conformance_packs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConformancePacksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_organization_conformance_packs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrganizationConformancePacks(ctx, input)
				}
				var results []*svc.DescribeOrganizationConformancePacksOutput
				p := svc.NewDescribeOrganizationConformancePacksPaginator(client, input)
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
		"describe-pending-aggregation-requests": {
			Name:   "describe-pending-aggregation-requests",
			Fields: fields_describe_pending_aggregation_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePendingAggregationRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_pending_aggregation_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePendingAggregationRequests(ctx, input)
				}
				var results []*svc.DescribePendingAggregationRequestsOutput
				p := svc.NewDescribePendingAggregationRequestsPaginator(client, input)
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
		"describe-remediation-configurations": {
			Name:   "describe-remediation-configurations",
			Fields: fields_describe_remediation_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRemediationConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_remediation_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRemediationConfigurations(ctx, input)
			},
		},
		"describe-remediation-exceptions": {
			Name:   "describe-remediation-exceptions",
			Fields: fields_describe_remediation_exceptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRemediationExceptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_remediation_exceptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRemediationExceptions(ctx, input)
				}
				var results []*svc.DescribeRemediationExceptionsOutput
				p := svc.NewDescribeRemediationExceptionsPaginator(client, input)
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
		"describe-remediation-execution-status": {
			Name:   "describe-remediation-execution-status",
			Fields: fields_describe_remediation_execution_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRemediationExecutionStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_remediation_execution_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRemediationExecutionStatus(ctx, input)
				}
				var results []*svc.DescribeRemediationExecutionStatusOutput
				p := svc.NewDescribeRemediationExecutionStatusPaginator(client, input)
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
		"describe-retention-configurations": {
			Name:   "describe-retention-configurations",
			Fields: fields_describe_retention_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRetentionConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_retention_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRetentionConfigurations(ctx, input)
				}
				var results []*svc.DescribeRetentionConfigurationsOutput
				p := svc.NewDescribeRetentionConfigurationsPaginator(client, input)
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
		"disassociate-resource-types": {
			Name:   "disassociate-resource-types",
			Fields: fields_disassociate_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResourceTypes(ctx, input)
			},
		},
		"get-aggregate-compliance-details-by-config-rule": {
			Name:   "get-aggregate-compliance-details-by-config-rule",
			Fields: fields_get_aggregate_compliance_details_by_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAggregateComplianceDetailsByConfigRuleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_aggregate_compliance_details_by_config_rule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAggregateComplianceDetailsByConfigRule(ctx, input)
				}
				var results []*svc.GetAggregateComplianceDetailsByConfigRuleOutput
				p := svc.NewGetAggregateComplianceDetailsByConfigRulePaginator(client, input)
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
		"get-aggregate-config-rule-compliance-summary": {
			Name:   "get-aggregate-config-rule-compliance-summary",
			Fields: fields_get_aggregate_config_rule_compliance_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAggregateConfigRuleComplianceSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_aggregate_config_rule_compliance_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAggregateConfigRuleComplianceSummary(ctx, input)
				}
				var results []*svc.GetAggregateConfigRuleComplianceSummaryOutput
				p := svc.NewGetAggregateConfigRuleComplianceSummaryPaginator(client, input)
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
		"get-aggregate-conformance-pack-compliance-summary": {
			Name:   "get-aggregate-conformance-pack-compliance-summary",
			Fields: fields_get_aggregate_conformance_pack_compliance_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAggregateConformancePackComplianceSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_aggregate_conformance_pack_compliance_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAggregateConformancePackComplianceSummary(ctx, input)
				}
				var results []*svc.GetAggregateConformancePackComplianceSummaryOutput
				p := svc.NewGetAggregateConformancePackComplianceSummaryPaginator(client, input)
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
		"get-aggregate-discovered-resource-counts": {
			Name:   "get-aggregate-discovered-resource-counts",
			Fields: fields_get_aggregate_discovered_resource_counts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAggregateDiscoveredResourceCountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_aggregate_discovered_resource_counts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAggregateDiscoveredResourceCounts(ctx, input)
				}
				var results []*svc.GetAggregateDiscoveredResourceCountsOutput
				p := svc.NewGetAggregateDiscoveredResourceCountsPaginator(client, input)
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
		"get-aggregate-resource-config": {
			Name:   "get-aggregate-resource-config",
			Fields: fields_get_aggregate_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAggregateResourceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_aggregate_resource_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAggregateResourceConfig(ctx, input)
			},
		},
		"get-compliance-details-by-config-rule": {
			Name:   "get-compliance-details-by-config-rule",
			Fields: fields_get_compliance_details_by_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComplianceDetailsByConfigRuleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_compliance_details_by_config_rule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetComplianceDetailsByConfigRule(ctx, input)
				}
				var results []*svc.GetComplianceDetailsByConfigRuleOutput
				p := svc.NewGetComplianceDetailsByConfigRulePaginator(client, input)
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
		"get-compliance-details-by-resource": {
			Name:   "get-compliance-details-by-resource",
			Fields: fields_get_compliance_details_by_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComplianceDetailsByResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_compliance_details_by_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetComplianceDetailsByResource(ctx, input)
				}
				var results []*svc.GetComplianceDetailsByResourceOutput
				p := svc.NewGetComplianceDetailsByResourcePaginator(client, input)
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
		"get-compliance-summary-by-config-rule": {
			Name:   "get-compliance-summary-by-config-rule",
			Fields: fields_get_compliance_summary_by_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComplianceSummaryByConfigRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compliance_summary_by_config_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComplianceSummaryByConfigRule(ctx, input)
			},
		},
		"get-compliance-summary-by-resource-type": {
			Name:   "get-compliance-summary-by-resource-type",
			Fields: fields_get_compliance_summary_by_resource_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComplianceSummaryByResourceTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compliance_summary_by_resource_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComplianceSummaryByResourceType(ctx, input)
			},
		},
		"get-conformance-pack-compliance-details": {
			Name:   "get-conformance-pack-compliance-details",
			Fields: fields_get_conformance_pack_compliance_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConformancePackComplianceDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_conformance_pack_compliance_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetConformancePackComplianceDetails(ctx, input)
				}
				var results []*svc.GetConformancePackComplianceDetailsOutput
				p := svc.NewGetConformancePackComplianceDetailsPaginator(client, input)
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
		"get-conformance-pack-compliance-summary": {
			Name:   "get-conformance-pack-compliance-summary",
			Fields: fields_get_conformance_pack_compliance_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConformancePackComplianceSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_conformance_pack_compliance_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetConformancePackComplianceSummary(ctx, input)
				}
				var results []*svc.GetConformancePackComplianceSummaryOutput
				p := svc.NewGetConformancePackComplianceSummaryPaginator(client, input)
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
		"get-custom-rule-policy": {
			Name:   "get-custom-rule-policy",
			Fields: fields_get_custom_rule_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomRulePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_rule_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomRulePolicy(ctx, input)
			},
		},
		"get-discovered-resource-counts": {
			Name:   "get-discovered-resource-counts",
			Fields: fields_get_discovered_resource_counts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDiscoveredResourceCountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_discovered_resource_counts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDiscoveredResourceCounts(ctx, input)
				}
				var results []*svc.GetDiscoveredResourceCountsOutput
				p := svc.NewGetDiscoveredResourceCountsPaginator(client, input)
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
		"get-organization-config-rule-detailed-status": {
			Name:   "get-organization-config-rule-detailed-status",
			Fields: fields_get_organization_config_rule_detailed_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationConfigRuleDetailedStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_organization_config_rule_detailed_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOrganizationConfigRuleDetailedStatus(ctx, input)
				}
				var results []*svc.GetOrganizationConfigRuleDetailedStatusOutput
				p := svc.NewGetOrganizationConfigRuleDetailedStatusPaginator(client, input)
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
		"get-organization-conformance-pack-detailed-status": {
			Name:   "get-organization-conformance-pack-detailed-status",
			Fields: fields_get_organization_conformance_pack_detailed_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationConformancePackDetailedStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_organization_conformance_pack_detailed_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOrganizationConformancePackDetailedStatus(ctx, input)
				}
				var results []*svc.GetOrganizationConformancePackDetailedStatusOutput
				p := svc.NewGetOrganizationConformancePackDetailedStatusPaginator(client, input)
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
		"get-organization-custom-rule-policy": {
			Name:   "get-organization-custom-rule-policy",
			Fields: fields_get_organization_custom_rule_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationCustomRulePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_organization_custom_rule_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOrganizationCustomRulePolicy(ctx, input)
			},
		},
		"get-resource-config-history": {
			Name:   "get-resource-config-history",
			Fields: fields_get_resource_config_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceConfigHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_config_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceConfigHistory(ctx, input)
				}
				var results []*svc.GetResourceConfigHistoryOutput
				p := svc.NewGetResourceConfigHistoryPaginator(client, input)
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
		"get-resource-evaluation-summary": {
			Name:   "get-resource-evaluation-summary",
			Fields: fields_get_resource_evaluation_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceEvaluationSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_evaluation_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceEvaluationSummary(ctx, input)
			},
		},
		"get-stored-query": {
			Name:   "get-stored-query",
			Fields: fields_get_stored_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStoredQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stored_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStoredQuery(ctx, input)
			},
		},
		"list-aggregate-discovered-resources": {
			Name:   "list-aggregate-discovered-resources",
			Fields: fields_list_aggregate_discovered_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAggregateDiscoveredResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aggregate_discovered_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAggregateDiscoveredResources(ctx, input)
				}
				var results []*svc.ListAggregateDiscoveredResourcesOutput
				p := svc.NewListAggregateDiscoveredResourcesPaginator(client, input)
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
		"list-configuration-recorders": {
			Name:   "list-configuration-recorders",
			Fields: fields_list_configuration_recorders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationRecordersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_recorders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationRecorders(ctx, input)
				}
				var results []*svc.ListConfigurationRecordersOutput
				p := svc.NewListConfigurationRecordersPaginator(client, input)
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
		"list-conformance-pack-compliance-scores": {
			Name:   "list-conformance-pack-compliance-scores",
			Fields: fields_list_conformance_pack_compliance_scores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConformancePackComplianceScoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_conformance_pack_compliance_scores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConformancePackComplianceScores(ctx, input)
				}
				var results []*svc.ListConformancePackComplianceScoresOutput
				p := svc.NewListConformancePackComplianceScoresPaginator(client, input)
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
		"list-discovered-resources": {
			Name:   "list-discovered-resources",
			Fields: fields_list_discovered_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDiscoveredResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_discovered_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDiscoveredResources(ctx, input)
				}
				var results []*svc.ListDiscoveredResourcesOutput
				p := svc.NewListDiscoveredResourcesPaginator(client, input)
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
		"list-resource-evaluations": {
			Name:   "list-resource-evaluations",
			Fields: fields_list_resource_evaluations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceEvaluationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_evaluations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceEvaluations(ctx, input)
				}
				var results []*svc.ListResourceEvaluationsOutput
				p := svc.NewListResourceEvaluationsPaginator(client, input)
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
		"list-stored-queries": {
			Name:   "list-stored-queries",
			Fields: fields_list_stored_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStoredQueriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stored_queries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStoredQueries(ctx, input)
				}
				var results []*svc.ListStoredQueriesOutput
				p := svc.NewListStoredQueriesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"put-aggregation-authorization": {
			Name:   "put-aggregation-authorization",
			Fields: fields_put_aggregation_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAggregationAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_aggregation_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAggregationAuthorization(ctx, input)
			},
		},
		"put-config-rule": {
			Name:   "put-config-rule",
			Fields: fields_put_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_config_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigRule(ctx, input)
			},
		},
		"put-configuration-aggregator": {
			Name:   "put-configuration-aggregator",
			Fields: fields_put_configuration_aggregator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationAggregatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_aggregator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationAggregator(ctx, input)
			},
		},
		"put-configuration-recorder": {
			Name:   "put-configuration-recorder",
			Fields: fields_put_configuration_recorder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationRecorderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_recorder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationRecorder(ctx, input)
			},
		},
		"put-conformance-pack": {
			Name:   "put-conformance-pack",
			Fields: fields_put_conformance_pack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConformancePackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_conformance_pack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConformancePack(ctx, input)
			},
		},
		"put-delivery-channel": {
			Name:   "put-delivery-channel",
			Fields: fields_put_delivery_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDeliveryChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_delivery_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDeliveryChannel(ctx, input)
			},
		},
		"put-evaluations": {
			Name:   "put-evaluations",
			Fields: fields_put_evaluations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEvaluationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_evaluations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEvaluations(ctx, input)
			},
		},
		"put-external-evaluation": {
			Name:   "put-external-evaluation",
			Fields: fields_put_external_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutExternalEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_external_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutExternalEvaluation(ctx, input)
			},
		},
		"put-organization-config-rule": {
			Name:   "put-organization-config-rule",
			Fields: fields_put_organization_config_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutOrganizationConfigRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_organization_config_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutOrganizationConfigRule(ctx, input)
			},
		},
		"put-organization-conformance-pack": {
			Name:   "put-organization-conformance-pack",
			Fields: fields_put_organization_conformance_pack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutOrganizationConformancePackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_organization_conformance_pack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutOrganizationConformancePack(ctx, input)
			},
		},
		"put-remediation-configurations": {
			Name:   "put-remediation-configurations",
			Fields: fields_put_remediation_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRemediationConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_remediation_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRemediationConfigurations(ctx, input)
			},
		},
		"put-remediation-exceptions": {
			Name:   "put-remediation-exceptions",
			Fields: fields_put_remediation_exceptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRemediationExceptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_remediation_exceptions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRemediationExceptions(ctx, input)
			},
		},
		"put-resource-config": {
			Name:   "put-resource-config",
			Fields: fields_put_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourceConfig(ctx, input)
			},
		},
		"put-retention-configuration": {
			Name:   "put-retention-configuration",
			Fields: fields_put_retention_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRetentionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_retention_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRetentionConfiguration(ctx, input)
			},
		},
		"put-service-linked-configuration-recorder": {
			Name:   "put-service-linked-configuration-recorder",
			Fields: fields_put_service_linked_configuration_recorder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutServiceLinkedConfigurationRecorderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_service_linked_configuration_recorder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutServiceLinkedConfigurationRecorder(ctx, input)
			},
		},
		"put-stored-query": {
			Name:   "put-stored-query",
			Fields: fields_put_stored_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutStoredQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_stored_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutStoredQuery(ctx, input)
			},
		},
		"select-aggregate-resource-config": {
			Name:   "select-aggregate-resource-config",
			Fields: fields_select_aggregate_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SelectAggregateResourceConfigInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_select_aggregate_resource_config, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SelectAggregateResourceConfig(ctx, input)
				}
				var results []*svc.SelectAggregateResourceConfigOutput
				p := svc.NewSelectAggregateResourceConfigPaginator(client, input)
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
		"select-resource-config": {
			Name:   "select-resource-config",
			Fields: fields_select_resource_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SelectResourceConfigInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_select_resource_config, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SelectResourceConfig(ctx, input)
				}
				var results []*svc.SelectResourceConfigOutput
				p := svc.NewSelectResourceConfigPaginator(client, input)
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
		"start-config-rules-evaluation": {
			Name:   "start-config-rules-evaluation",
			Fields: fields_start_config_rules_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConfigRulesEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_config_rules_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConfigRulesEvaluation(ctx, input)
			},
		},
		"start-configuration-recorder": {
			Name:   "start-configuration-recorder",
			Fields: fields_start_configuration_recorder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConfigurationRecorderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_configuration_recorder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConfigurationRecorder(ctx, input)
			},
		},
		"start-remediation-execution": {
			Name:   "start-remediation-execution",
			Fields: fields_start_remediation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRemediationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_remediation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRemediationExecution(ctx, input)
			},
		},
		"start-resource-evaluation": {
			Name:   "start-resource-evaluation",
			Fields: fields_start_resource_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartResourceEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_resource_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartResourceEvaluation(ctx, input)
			},
		},
		"stop-configuration-recorder": {
			Name:   "stop-configuration-recorder",
			Fields: fields_stop_configuration_recorder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopConfigurationRecorderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_configuration_recorder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopConfigurationRecorder(ctx, input)
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
	}
	if err := leanruntime.Execute("configservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
