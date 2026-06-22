package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
)

var fields_accept_network_firewall_transit_gateway_attachment = []leanruntime.Field{
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_associate_availability_zones = []leanruntime.Field{
	{Name: "AvailabilityZoneMappings", Flag: "availability-zone-mappings", Type: "[]types.AvailabilityZoneMapping", Required: true},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_associate_firewall_policy = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "FirewallPolicyArn", Flag: "firewall-policy-arn", Type: "*string", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_associate_subnets = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "SubnetMappings", Flag: "subnet-mappings", Type: "[]types.SubnetMapping", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_attach_rule_groups_to_proxy_configuration = []leanruntime.Field{
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
	{Name: "RuleGroups", Flag: "rule-groups", Type: "[]types.ProxyRuleGroupAttachment", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_create_firewall = []leanruntime.Field{
	{Name: "AvailabilityZoneChangeProtection", Flag: "availability-zone-change-protection", Type: "bool", Required: false},
	{Name: "AvailabilityZoneMappings", Flag: "availability-zone-mappings", Type: "[]types.AvailabilityZoneMapping", Required: false},
	{Name: "DeleteProtection", Flag: "delete-protection", Type: "bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnabledAnalysisTypes", Flag: "enabled-analysis-types", Type: "[]types.EnabledAnalysisType", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: true},
	{Name: "FirewallPolicyArn", Flag: "firewall-policy-arn", Type: "*string", Required: true},
	{Name: "FirewallPolicyChangeProtection", Flag: "firewall-policy-change-protection", Type: "bool", Required: false},
	{Name: "SubnetChangeProtection", Flag: "subnet-change-protection", Type: "bool", Required: false},
	{Name: "SubnetMappings", Flag: "subnet-mappings", Type: "[]types.SubnetMapping", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_create_firewall_policy = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "FirewallPolicy", Flag: "firewall-policy", Type: "*types.FirewallPolicy", Required: true},
	{Name: "FirewallPolicyName", Flag: "firewall-policy-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_proxy = []leanruntime.Field{
	{Name: "ListenerProperties", Flag: "listener-properties", Type: "[]types.ListenerPropertyRequest", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
	{Name: "ProxyName", Flag: "proxy-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TlsInterceptProperties", Flag: "tls-intercept-properties", Type: "*types.TlsInterceptPropertiesRequest", Required: true},
}

var fields_create_proxy_configuration = []leanruntime.Field{
	{Name: "DefaultRulePhaseActions", Flag: "default-rule-phase-actions", Type: "*types.ProxyConfigDefaultRulePhaseActionsRequest", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: true},
	{Name: "RuleGroupArns", Flag: "rule-group-arns", Type: "[]string", Required: false},
	{Name: "RuleGroupNames", Flag: "rule-group-names", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_proxy_rule_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "*types.ProxyRulesByRequestPhase", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_proxy_rules = []leanruntime.Field{
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
	{Name: "Rules", Flag: "rules", Type: "*types.CreateProxyRulesByRequestPhase", Required: true},
}

var fields_create_rule_group = []leanruntime.Field{
	{Name: "AnalyzeRuleGroup", Flag: "analyze-rule-group", Type: "bool", Required: false},
	{Name: "Capacity", Flag: "capacity", Type: "*int32", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "RuleGroup", Flag: "rule-group", Type: "*types.RuleGroup", Required: false},
	{Name: "RuleGroupName", Flag: "rule-group-name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "*string", Required: false},
	{Name: "SourceMetadata", Flag: "source-metadata", Type: "*types.SourceMetadata", Required: false},
	{Name: "SummaryConfiguration", Flag: "summary-configuration", Type: "*types.SummaryConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: true},
}

var fields_create_tls_inspection_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "TLSInspectionConfiguration", Flag: "tls-inspection-configuration", Type: "*types.TLSInspectionConfiguration", Required: true},
	{Name: "TLSInspectionConfigurationName", Flag: "tls-inspection-configuration-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_vpc_endpoint_association = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: true},
	{Name: "SubnetMapping", Flag: "subnet-mapping", Type: "*types.SubnetMapping", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_delete_firewall = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
}

var fields_delete_firewall_policy = []leanruntime.Field{
	{Name: "FirewallPolicyArn", Flag: "firewall-policy-arn", Type: "*string", Required: false},
	{Name: "FirewallPolicyName", Flag: "firewall-policy-name", Type: "*string", Required: false},
}

var fields_delete_network_firewall_transit_gateway_attachment = []leanruntime.Field{
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_delete_proxy = []leanruntime.Field{
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
	{Name: "ProxyArn", Flag: "proxy-arn", Type: "*string", Required: false},
	{Name: "ProxyName", Flag: "proxy-name", Type: "*string", Required: false},
}

var fields_delete_proxy_configuration = []leanruntime.Field{
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
}

var fields_delete_proxy_rule_group = []leanruntime.Field{
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
}

var fields_delete_proxy_rules = []leanruntime.Field{
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_rule_group = []leanruntime.Field{
	{Name: "RuleGroupArn", Flag: "rule-group-arn", Type: "*string", Required: false},
	{Name: "RuleGroupName", Flag: "rule-group-name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: false},
}

var fields_delete_tls_inspection_configuration = []leanruntime.Field{
	{Name: "TLSInspectionConfigurationArn", Flag: "tls-inspection-configuration-arn", Type: "*string", Required: false},
	{Name: "TLSInspectionConfigurationName", Flag: "tls-inspection-configuration-name", Type: "*string", Required: false},
}

var fields_delete_vpc_endpoint_association = []leanruntime.Field{
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: true},
}

var fields_describe_firewall = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
}

var fields_describe_firewall_metadata = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
}

var fields_describe_firewall_policy = []leanruntime.Field{
	{Name: "FirewallPolicyArn", Flag: "firewall-policy-arn", Type: "*string", Required: false},
	{Name: "FirewallPolicyName", Flag: "firewall-policy-name", Type: "*string", Required: false},
}

var fields_describe_flow_operation = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: true},
	{Name: "FlowOperationId", Flag: "flow-operation-id", Type: "*string", Required: true},
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_describe_logging_configuration = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
}

var fields_describe_proxy = []leanruntime.Field{
	{Name: "ProxyArn", Flag: "proxy-arn", Type: "*string", Required: false},
	{Name: "ProxyName", Flag: "proxy-name", Type: "*string", Required: false},
}

var fields_describe_proxy_configuration = []leanruntime.Field{
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
}

var fields_describe_proxy_rule = []leanruntime.Field{
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
	{Name: "ProxyRuleName", Flag: "proxy-rule-name", Type: "*string", Required: true},
}

var fields_describe_proxy_rule_group = []leanruntime.Field{
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
}

var fields_describe_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_rule_group = []leanruntime.Field{
	{Name: "AnalyzeRuleGroup", Flag: "analyze-rule-group", Type: "bool", Required: false},
	{Name: "RuleGroupArn", Flag: "rule-group-arn", Type: "*string", Required: false},
	{Name: "RuleGroupName", Flag: "rule-group-name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: false},
}

var fields_describe_rule_group_metadata = []leanruntime.Field{
	{Name: "RuleGroupArn", Flag: "rule-group-arn", Type: "*string", Required: false},
	{Name: "RuleGroupName", Flag: "rule-group-name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: false},
}

var fields_describe_rule_group_summary = []leanruntime.Field{
	{Name: "RuleGroupArn", Flag: "rule-group-arn", Type: "*string", Required: false},
	{Name: "RuleGroupName", Flag: "rule-group-name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: false},
}

var fields_describe_tls_inspection_configuration = []leanruntime.Field{
	{Name: "TLSInspectionConfigurationArn", Flag: "tls-inspection-configuration-arn", Type: "*string", Required: false},
	{Name: "TLSInspectionConfigurationName", Flag: "tls-inspection-configuration-name", Type: "*string", Required: false},
}

var fields_describe_vpc_endpoint_association = []leanruntime.Field{
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: true},
}

var fields_detach_rule_groups_from_proxy_configuration = []leanruntime.Field{
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
	{Name: "RuleGroupArns", Flag: "rule-group-arns", Type: "[]string", Required: false},
	{Name: "RuleGroupNames", Flag: "rule-group-names", Type: "[]string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_disassociate_availability_zones = []leanruntime.Field{
	{Name: "AvailabilityZoneMappings", Flag: "availability-zone-mappings", Type: "[]types.AvailabilityZoneMapping", Required: true},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_disassociate_subnets = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_get_analysis_report_results = []leanruntime.Field{
	{Name: "AnalysisReportId", Flag: "analysis-report-id", Type: "*string", Required: true},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_analysis_reports = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewall_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_firewalls = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_list_flow_operation_results = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: true},
	{Name: "FlowOperationId", Flag: "flow-operation-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_list_flow_operations = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: true},
	{Name: "FlowOperationType", Flag: "flow-operation-type", Type: "types.FlowOperationType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_list_proxies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_proxy_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_proxy_rule_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rule_groups = []leanruntime.Field{
	{Name: "ManagedType", Flag: "managed-type", Type: "types.ResourceManagedType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.ResourceManagedStatus", Required: false},
	{Name: "SubscriptionStatus", Flag: "subscription-status", Type: "types.SubscriptionStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tls_inspection_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_endpoint_associations = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reject_network_firewall_transit_gateway_attachment = []leanruntime.Field{
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_start_analysis_report = []leanruntime.Field{
	{Name: "AnalysisType", Flag: "analysis-type", Type: "types.EnabledAnalysisType", Required: true},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
}

var fields_start_flow_capture = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: true},
	{Name: "FlowFilters", Flag: "flow-filters", Type: "[]types.FlowFilter", Required: true},
	{Name: "MinimumFlowAgeInSeconds", Flag: "minimum-flow-age-in-seconds", Type: "*int32", Required: false},
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_start_flow_flush = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: true},
	{Name: "FlowFilters", Flag: "flow-filters", Type: "[]types.FlowFilter", Required: true},
	{Name: "MinimumFlowAgeInSeconds", Flag: "minimum-flow-age-in-seconds", Type: "*int32", Required: false},
	{Name: "VpcEndpointAssociationArn", Flag: "vpc-endpoint-association-arn", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_availability_zone_change_protection = []leanruntime.Field{
	{Name: "AvailabilityZoneChangeProtection", Flag: "availability-zone-change-protection", Type: "bool", Required: true},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_firewall_analysis_settings = []leanruntime.Field{
	{Name: "EnabledAnalysisTypes", Flag: "enabled-analysis-types", Type: "[]types.EnabledAnalysisType", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_firewall_delete_protection = []leanruntime.Field{
	{Name: "DeleteProtection", Flag: "delete-protection", Type: "bool", Required: true},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_firewall_description = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_firewall_encryption_configuration = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_firewall_policy = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "FirewallPolicy", Flag: "firewall-policy", Type: "*types.FirewallPolicy", Required: true},
	{Name: "FirewallPolicyArn", Flag: "firewall-policy-arn", Type: "*string", Required: false},
	{Name: "FirewallPolicyName", Flag: "firewall-policy-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_firewall_policy_change_protection = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "FirewallPolicyChangeProtection", Flag: "firewall-policy-change-protection", Type: "bool", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_logging_configuration = []leanruntime.Field{
	{Name: "EnableMonitoringDashboard", Flag: "enable-monitoring-dashboard", Type: "*bool", Required: false},
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: false},
}

var fields_update_proxy = []leanruntime.Field{
	{Name: "ListenerPropertiesToAdd", Flag: "listener-properties-to-add", Type: "[]types.ListenerPropertyRequest", Required: false},
	{Name: "ListenerPropertiesToRemove", Flag: "listener-properties-to-remove", Type: "[]types.ListenerPropertyRequest", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
	{Name: "ProxyArn", Flag: "proxy-arn", Type: "*string", Required: false},
	{Name: "ProxyName", Flag: "proxy-name", Type: "*string", Required: false},
	{Name: "TlsInterceptProperties", Flag: "tls-intercept-properties", Type: "*types.TlsInterceptPropertiesRequest", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_proxy_configuration = []leanruntime.Field{
	{Name: "DefaultRulePhaseActions", Flag: "default-rule-phase-actions", Type: "*types.ProxyConfigDefaultRulePhaseActionsRequest", Required: true},
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_proxy_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.ProxyRulePhaseAction", Required: false},
	{Name: "AddConditions", Flag: "add-conditions", Type: "[]types.ProxyRuleCondition", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
	{Name: "ProxyRuleName", Flag: "proxy-rule-name", Type: "*string", Required: true},
	{Name: "RemoveConditions", Flag: "remove-conditions", Type: "[]types.ProxyRuleCondition", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_proxy_rule_group_priorities = []leanruntime.Field{
	{Name: "ProxyConfigurationArn", Flag: "proxy-configuration-arn", Type: "*string", Required: false},
	{Name: "ProxyConfigurationName", Flag: "proxy-configuration-name", Type: "*string", Required: false},
	{Name: "RuleGroups", Flag: "rule-groups", Type: "[]types.ProxyRuleGroupPriority", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_proxy_rule_priorities = []leanruntime.Field{
	{Name: "ProxyRuleGroupArn", Flag: "proxy-rule-group-arn", Type: "*string", Required: false},
	{Name: "ProxyRuleGroupName", Flag: "proxy-rule-group-name", Type: "*string", Required: false},
	{Name: "RuleGroupRequestPhase", Flag: "rule-group-request-phase", Type: "types.RuleGroupRequestPhase", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.ProxyRulePriority", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_rule_group = []leanruntime.Field{
	{Name: "AnalyzeRuleGroup", Flag: "analyze-rule-group", Type: "bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "RuleGroup", Flag: "rule-group", Type: "*types.RuleGroup", Required: false},
	{Name: "RuleGroupArn", Flag: "rule-group-arn", Type: "*string", Required: false},
	{Name: "RuleGroupName", Flag: "rule-group-name", Type: "*string", Required: false},
	{Name: "Rules", Flag: "rules", Type: "*string", Required: false},
	{Name: "SourceMetadata", Flag: "source-metadata", Type: "*types.SourceMetadata", Required: false},
	{Name: "SummaryConfiguration", Flag: "summary-configuration", Type: "*types.SummaryConfiguration", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RuleGroupType", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_subnet_change_protection = []leanruntime.Field{
	{Name: "FirewallArn", Flag: "firewall-arn", Type: "*string", Required: false},
	{Name: "FirewallName", Flag: "firewall-name", Type: "*string", Required: false},
	{Name: "SubnetChangeProtection", Flag: "subnet-change-protection", Type: "bool", Required: true},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: false},
}

var fields_update_tls_inspection_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "TLSInspectionConfiguration", Flag: "tls-inspection-configuration", Type: "*types.TLSInspectionConfiguration", Required: true},
	{Name: "TLSInspectionConfigurationArn", Flag: "tls-inspection-configuration-arn", Type: "*string", Required: false},
	{Name: "TLSInspectionConfigurationName", Flag: "tls-inspection-configuration-name", Type: "*string", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-network-firewall-transit-gateway-attachment": {
			Name:   "accept-network-firewall-transit-gateway-attachment",
			Fields: fields_accept_network_firewall_transit_gateway_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptNetworkFirewallTransitGatewayAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_network_firewall_transit_gateway_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptNetworkFirewallTransitGatewayAttachment(ctx, input)
			},
		},
		"associate-availability-zones": {
			Name:   "associate-availability-zones",
			Fields: fields_associate_availability_zones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAvailabilityZonesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_availability_zones, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAvailabilityZones(ctx, input)
			},
		},
		"associate-firewall-policy": {
			Name:   "associate-firewall-policy",
			Fields: fields_associate_firewall_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFirewallPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_firewall_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFirewallPolicy(ctx, input)
			},
		},
		"associate-subnets": {
			Name:   "associate-subnets",
			Fields: fields_associate_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSubnetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_subnets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSubnets(ctx, input)
			},
		},
		"attach-rule-groups-to-proxy-configuration": {
			Name:   "attach-rule-groups-to-proxy-configuration",
			Fields: fields_attach_rule_groups_to_proxy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachRuleGroupsToProxyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_rule_groups_to_proxy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachRuleGroupsToProxyConfiguration(ctx, input)
			},
		},
		"create-firewall": {
			Name:   "create-firewall",
			Fields: fields_create_firewall,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFirewallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_firewall, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFirewall(ctx, input)
			},
		},
		"create-firewall-policy": {
			Name:   "create-firewall-policy",
			Fields: fields_create_firewall_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFirewallPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_firewall_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFirewallPolicy(ctx, input)
			},
		},
		"create-proxy": {
			Name:   "create-proxy",
			Fields: fields_create_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProxy(ctx, input)
			},
		},
		"create-proxy-configuration": {
			Name:   "create-proxy-configuration",
			Fields: fields_create_proxy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProxyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_proxy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProxyConfiguration(ctx, input)
			},
		},
		"create-proxy-rule-group": {
			Name:   "create-proxy-rule-group",
			Fields: fields_create_proxy_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProxyRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_proxy_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProxyRuleGroup(ctx, input)
			},
		},
		"create-proxy-rules": {
			Name:   "create-proxy-rules",
			Fields: fields_create_proxy_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProxyRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_proxy_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProxyRules(ctx, input)
			},
		},
		"create-rule-group": {
			Name:   "create-rule-group",
			Fields: fields_create_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRuleGroup(ctx, input)
			},
		},
		"create-tls-inspection-configuration": {
			Name:   "create-tls-inspection-configuration",
			Fields: fields_create_tls_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTLSInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tls_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTLSInspectionConfiguration(ctx, input)
			},
		},
		"create-vpc-endpoint-association": {
			Name:   "create-vpc-endpoint-association",
			Fields: fields_create_vpc_endpoint_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpointAssociation(ctx, input)
			},
		},
		"delete-firewall": {
			Name:   "delete-firewall",
			Fields: fields_delete_firewall,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFirewallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_firewall, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFirewall(ctx, input)
			},
		},
		"delete-firewall-policy": {
			Name:   "delete-firewall-policy",
			Fields: fields_delete_firewall_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFirewallPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_firewall_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFirewallPolicy(ctx, input)
			},
		},
		"delete-network-firewall-transit-gateway-attachment": {
			Name:   "delete-network-firewall-transit-gateway-attachment",
			Fields: fields_delete_network_firewall_transit_gateway_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkFirewallTransitGatewayAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_firewall_transit_gateway_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkFirewallTransitGatewayAttachment(ctx, input)
			},
		},
		"delete-proxy": {
			Name:   "delete-proxy",
			Fields: fields_delete_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProxy(ctx, input)
			},
		},
		"delete-proxy-configuration": {
			Name:   "delete-proxy-configuration",
			Fields: fields_delete_proxy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProxyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_proxy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProxyConfiguration(ctx, input)
			},
		},
		"delete-proxy-rule-group": {
			Name:   "delete-proxy-rule-group",
			Fields: fields_delete_proxy_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProxyRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_proxy_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProxyRuleGroup(ctx, input)
			},
		},
		"delete-proxy-rules": {
			Name:   "delete-proxy-rules",
			Fields: fields_delete_proxy_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProxyRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_proxy_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProxyRules(ctx, input)
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
		"delete-rule-group": {
			Name:   "delete-rule-group",
			Fields: fields_delete_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRuleGroup(ctx, input)
			},
		},
		"delete-tls-inspection-configuration": {
			Name:   "delete-tls-inspection-configuration",
			Fields: fields_delete_tls_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTLSInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tls_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTLSInspectionConfiguration(ctx, input)
			},
		},
		"delete-vpc-endpoint-association": {
			Name:   "delete-vpc-endpoint-association",
			Fields: fields_delete_vpc_endpoint_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoint_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpointAssociation(ctx, input)
			},
		},
		"describe-firewall": {
			Name:   "describe-firewall",
			Fields: fields_describe_firewall,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFirewallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_firewall, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFirewall(ctx, input)
			},
		},
		"describe-firewall-metadata": {
			Name:   "describe-firewall-metadata",
			Fields: fields_describe_firewall_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFirewallMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_firewall_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFirewallMetadata(ctx, input)
			},
		},
		"describe-firewall-policy": {
			Name:   "describe-firewall-policy",
			Fields: fields_describe_firewall_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFirewallPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_firewall_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFirewallPolicy(ctx, input)
			},
		},
		"describe-flow-operation": {
			Name:   "describe-flow-operation",
			Fields: fields_describe_flow_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flow_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlowOperation(ctx, input)
			},
		},
		"describe-logging-configuration": {
			Name:   "describe-logging-configuration",
			Fields: fields_describe_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoggingConfiguration(ctx, input)
			},
		},
		"describe-proxy": {
			Name:   "describe-proxy",
			Fields: fields_describe_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProxy(ctx, input)
			},
		},
		"describe-proxy-configuration": {
			Name:   "describe-proxy-configuration",
			Fields: fields_describe_proxy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProxyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_proxy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProxyConfiguration(ctx, input)
			},
		},
		"describe-proxy-rule": {
			Name:   "describe-proxy-rule",
			Fields: fields_describe_proxy_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProxyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_proxy_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProxyRule(ctx, input)
			},
		},
		"describe-proxy-rule-group": {
			Name:   "describe-proxy-rule-group",
			Fields: fields_describe_proxy_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProxyRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_proxy_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProxyRuleGroup(ctx, input)
			},
		},
		"describe-resource-policy": {
			Name:   "describe-resource-policy",
			Fields: fields_describe_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicy(ctx, input)
			},
		},
		"describe-rule-group": {
			Name:   "describe-rule-group",
			Fields: fields_describe_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRuleGroup(ctx, input)
			},
		},
		"describe-rule-group-metadata": {
			Name:   "describe-rule-group-metadata",
			Fields: fields_describe_rule_group_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuleGroupMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rule_group_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRuleGroupMetadata(ctx, input)
			},
		},
		"describe-rule-group-summary": {
			Name:   "describe-rule-group-summary",
			Fields: fields_describe_rule_group_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuleGroupSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rule_group_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRuleGroupSummary(ctx, input)
			},
		},
		"describe-tls-inspection-configuration": {
			Name:   "describe-tls-inspection-configuration",
			Fields: fields_describe_tls_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTLSInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tls_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTLSInspectionConfiguration(ctx, input)
			},
		},
		"describe-vpc-endpoint-association": {
			Name:   "describe-vpc-endpoint-association",
			Fields: fields_describe_vpc_endpoint_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcEndpointAssociation(ctx, input)
			},
		},
		"detach-rule-groups-from-proxy-configuration": {
			Name:   "detach-rule-groups-from-proxy-configuration",
			Fields: fields_detach_rule_groups_from_proxy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachRuleGroupsFromProxyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_rule_groups_from_proxy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachRuleGroupsFromProxyConfiguration(ctx, input)
			},
		},
		"disassociate-availability-zones": {
			Name:   "disassociate-availability-zones",
			Fields: fields_disassociate_availability_zones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAvailabilityZonesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_availability_zones, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAvailabilityZones(ctx, input)
			},
		},
		"disassociate-subnets": {
			Name:   "disassociate-subnets",
			Fields: fields_disassociate_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSubnetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_subnets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSubnets(ctx, input)
			},
		},
		"get-analysis-report-results": {
			Name:   "get-analysis-report-results",
			Fields: fields_get_analysis_report_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnalysisReportResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_analysis_report_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAnalysisReportResults(ctx, input)
				}
				var results []*svc.GetAnalysisReportResultsOutput
				p := svc.NewGetAnalysisReportResultsPaginator(client, input)
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
		"list-analysis-reports": {
			Name:   "list-analysis-reports",
			Fields: fields_list_analysis_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalysisReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_analysis_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnalysisReports(ctx, input)
				}
				var results []*svc.ListAnalysisReportsOutput
				p := svc.NewListAnalysisReportsPaginator(client, input)
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
		"list-firewall-policies": {
			Name:   "list-firewall-policies",
			Fields: fields_list_firewall_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewall_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewallPolicies(ctx, input)
				}
				var results []*svc.ListFirewallPoliciesOutput
				p := svc.NewListFirewallPoliciesPaginator(client, input)
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
		"list-firewalls": {
			Name:   "list-firewalls",
			Fields: fields_list_firewalls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFirewallsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_firewalls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFirewalls(ctx, input)
				}
				var results []*svc.ListFirewallsOutput
				p := svc.NewListFirewallsPaginator(client, input)
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
		"list-flow-operation-results": {
			Name:   "list-flow-operation-results",
			Fields: fields_list_flow_operation_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowOperationResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_operation_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowOperationResults(ctx, input)
				}
				var results []*svc.ListFlowOperationResultsOutput
				p := svc.NewListFlowOperationResultsPaginator(client, input)
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
		"list-flow-operations": {
			Name:   "list-flow-operations",
			Fields: fields_list_flow_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowOperations(ctx, input)
				}
				var results []*svc.ListFlowOperationsOutput
				p := svc.NewListFlowOperationsPaginator(client, input)
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
		"list-proxies": {
			Name:   "list-proxies",
			Fields: fields_list_proxies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProxiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_proxies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProxies(ctx, input)
				}
				var results []*svc.ListProxiesOutput
				p := svc.NewListProxiesPaginator(client, input)
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
		"list-proxy-configurations": {
			Name:   "list-proxy-configurations",
			Fields: fields_list_proxy_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProxyConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_proxy_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProxyConfigurations(ctx, input)
				}
				var results []*svc.ListProxyConfigurationsOutput
				p := svc.NewListProxyConfigurationsPaginator(client, input)
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
		"list-proxy-rule-groups": {
			Name:   "list-proxy-rule-groups",
			Fields: fields_list_proxy_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProxyRuleGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_proxy_rule_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProxyRuleGroups(ctx, input)
				}
				var results []*svc.ListProxyRuleGroupsOutput
				p := svc.NewListProxyRuleGroupsPaginator(client, input)
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
		"list-rule-groups": {
			Name:   "list-rule-groups",
			Fields: fields_list_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rule_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuleGroups(ctx, input)
				}
				var results []*svc.ListRuleGroupsOutput
				p := svc.NewListRuleGroupsPaginator(client, input)
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
		"list-tls-inspection-configurations": {
			Name:   "list-tls-inspection-configurations",
			Fields: fields_list_tls_inspection_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTLSInspectionConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tls_inspection_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTLSInspectionConfigurations(ctx, input)
				}
				var results []*svc.ListTLSInspectionConfigurationsOutput
				p := svc.NewListTLSInspectionConfigurationsPaginator(client, input)
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
		"list-vpc-endpoint-associations": {
			Name:   "list-vpc-endpoint-associations",
			Fields: fields_list_vpc_endpoint_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoint_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVpcEndpointAssociations(ctx, input)
				}
				var results []*svc.ListVpcEndpointAssociationsOutput
				p := svc.NewListVpcEndpointAssociationsPaginator(client, input)
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
		"reject-network-firewall-transit-gateway-attachment": {
			Name:   "reject-network-firewall-transit-gateway-attachment",
			Fields: fields_reject_network_firewall_transit_gateway_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectNetworkFirewallTransitGatewayAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_network_firewall_transit_gateway_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectNetworkFirewallTransitGatewayAttachment(ctx, input)
			},
		},
		"start-analysis-report": {
			Name:   "start-analysis-report",
			Fields: fields_start_analysis_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAnalysisReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_analysis_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAnalysisReport(ctx, input)
			},
		},
		"start-flow-capture": {
			Name:   "start-flow-capture",
			Fields: fields_start_flow_capture,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFlowCaptureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_flow_capture, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFlowCapture(ctx, input)
			},
		},
		"start-flow-flush": {
			Name:   "start-flow-flush",
			Fields: fields_start_flow_flush,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFlowFlushInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_flow_flush, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFlowFlush(ctx, input)
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
		"update-availability-zone-change-protection": {
			Name:   "update-availability-zone-change-protection",
			Fields: fields_update_availability_zone_change_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAvailabilityZoneChangeProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_availability_zone_change_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAvailabilityZoneChangeProtection(ctx, input)
			},
		},
		"update-firewall-analysis-settings": {
			Name:   "update-firewall-analysis-settings",
			Fields: fields_update_firewall_analysis_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallAnalysisSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_analysis_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallAnalysisSettings(ctx, input)
			},
		},
		"update-firewall-delete-protection": {
			Name:   "update-firewall-delete-protection",
			Fields: fields_update_firewall_delete_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallDeleteProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_delete_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallDeleteProtection(ctx, input)
			},
		},
		"update-firewall-description": {
			Name:   "update-firewall-description",
			Fields: fields_update_firewall_description,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallDescriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_description, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallDescription(ctx, input)
			},
		},
		"update-firewall-encryption-configuration": {
			Name:   "update-firewall-encryption-configuration",
			Fields: fields_update_firewall_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallEncryptionConfiguration(ctx, input)
			},
		},
		"update-firewall-policy": {
			Name:   "update-firewall-policy",
			Fields: fields_update_firewall_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallPolicy(ctx, input)
			},
		},
		"update-firewall-policy-change-protection": {
			Name:   "update-firewall-policy-change-protection",
			Fields: fields_update_firewall_policy_change_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFirewallPolicyChangeProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_firewall_policy_change_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFirewallPolicyChangeProtection(ctx, input)
			},
		},
		"update-logging-configuration": {
			Name:   "update-logging-configuration",
			Fields: fields_update_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLoggingConfiguration(ctx, input)
			},
		},
		"update-proxy": {
			Name:   "update-proxy",
			Fields: fields_update_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProxy(ctx, input)
			},
		},
		"update-proxy-configuration": {
			Name:   "update-proxy-configuration",
			Fields: fields_update_proxy_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProxyConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_proxy_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProxyConfiguration(ctx, input)
			},
		},
		"update-proxy-rule": {
			Name:   "update-proxy-rule",
			Fields: fields_update_proxy_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProxyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_proxy_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProxyRule(ctx, input)
			},
		},
		"update-proxy-rule-group-priorities": {
			Name:   "update-proxy-rule-group-priorities",
			Fields: fields_update_proxy_rule_group_priorities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProxyRuleGroupPrioritiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_proxy_rule_group_priorities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProxyRuleGroupPriorities(ctx, input)
			},
		},
		"update-proxy-rule-priorities": {
			Name:   "update-proxy-rule-priorities",
			Fields: fields_update_proxy_rule_priorities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProxyRulePrioritiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_proxy_rule_priorities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProxyRulePriorities(ctx, input)
			},
		},
		"update-rule-group": {
			Name:   "update-rule-group",
			Fields: fields_update_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuleGroup(ctx, input)
			},
		},
		"update-subnet-change-protection": {
			Name:   "update-subnet-change-protection",
			Fields: fields_update_subnet_change_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubnetChangeProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subnet_change_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubnetChangeProtection(ctx, input)
			},
		},
		"update-tls-inspection-configuration": {
			Name:   "update-tls-inspection-configuration",
			Fields: fields_update_tls_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTLSInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_tls_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTLSInspectionConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("networkfirewall", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
