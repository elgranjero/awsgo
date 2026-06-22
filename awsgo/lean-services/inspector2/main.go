package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/inspector2"
)

var fields_associate_member = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_batch_associate_code_security_scan_configuration = []leanruntime.Field{
	{Name: "AssociateConfigurationRequests", Flag: "associate-configuration-requests", Type: "[]types.AssociateConfigurationRequest", Required: true},
}

var fields_batch_disassociate_code_security_scan_configuration = []leanruntime.Field{
	{Name: "DisassociateConfigurationRequests", Flag: "disassociate-configuration-requests", Type: "[]types.DisassociateConfigurationRequest", Required: true},
}

var fields_batch_get_account_status = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
}

var fields_batch_get_code_snippet = []leanruntime.Field{
	{Name: "FindingArns", Flag: "finding-arns", Type: "[]string", Required: true},
}

var fields_batch_get_finding_details = []leanruntime.Field{
	{Name: "FindingArns", Flag: "finding-arns", Type: "[]string", Required: true},
}

var fields_batch_get_free_trial_info = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_batch_get_member_ec2_deep_inspection_status = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
}

var fields_batch_update_member_ec2_deep_inspection_status = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]types.MemberAccountEc2DeepInspectionStatus", Required: true},
}

var fields_cancel_findings_report = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_cancel_sbom_export = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_create_cis_scan_configuration = []leanruntime.Field{
	{Name: "ScanName", Flag: "scan-name", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "types.Schedule", Required: true},
	{Name: "SecurityLevel", Flag: "security-level", Type: "types.CisSecurityLevel", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "*types.CreateCisTargets", Required: true},
}

var fields_create_code_security_integration = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "types.CreateIntegrationDetail", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.IntegrationType", Required: true},
}

var fields_create_code_security_scan_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.CodeSecurityScanConfiguration", Required: true},
	{Name: "Level", Flag: "level", Type: "types.ConfigurationLevel", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ScopeSettings", Flag: "scope-settings", Type: "*types.ScopeSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_filter = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FilterAction", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_findings_report = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "ReportFormat", Flag: "report-format", Type: "types.ReportFormat", Required: true},
	{Name: "S3Destination", Flag: "s3-destination", Type: "*types.Destination", Required: true},
}

var fields_create_sbom_export = []leanruntime.Field{
	{Name: "ReportFormat", Flag: "report-format", Type: "types.SbomReportFormat", Required: true},
	{Name: "ResourceFilterCriteria", Flag: "resource-filter-criteria", Type: "*types.ResourceFilterCriteria", Required: false},
	{Name: "S3Destination", Flag: "s3-destination", Type: "*types.Destination", Required: true},
}

var fields_delete_cis_scan_configuration = []leanruntime.Field{
	{Name: "ScanConfigurationArn", Flag: "scan-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_code_security_integration = []leanruntime.Field{
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
}

var fields_delete_code_security_scan_configuration = []leanruntime.Field{
	{Name: "ScanConfigurationArn", Flag: "scan-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_filter = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_describe_organization_configuration = []leanruntime.Field{}

var fields_disable = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceScanType", Required: false},
}

var fields_disable_delegated_admin_account = []leanruntime.Field{
	{Name: "DelegatedAdminAccountId", Flag: "delegated-admin-account-id", Type: "*string", Required: true},
}

var fields_disassociate_member = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_enable = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceScanType", Required: true},
}

var fields_enable_delegated_admin_account = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DelegatedAdminAccountId", Flag: "delegated-admin-account-id", Type: "*string", Required: true},
}

var fields_get_cis_scan_report = []leanruntime.Field{
	{Name: "ReportFormat", Flag: "report-format", Type: "types.CisReportFormat", Required: false},
	{Name: "ScanArn", Flag: "scan-arn", Type: "*string", Required: true},
	{Name: "TargetAccounts", Flag: "target-accounts", Type: "[]string", Required: false},
}

var fields_get_cis_scan_result_details = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CisScanResultDetailsFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanArn", Flag: "scan-arn", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CisScanResultDetailsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.CisSortOrder", Required: false},
	{Name: "TargetResourceId", Flag: "target-resource-id", Type: "*string", Required: true},
}

var fields_get_clusters_for_image = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ClusterForImageFilterCriteria", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_code_security_integration = []leanruntime.Field{
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_get_code_security_scan = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "types.CodeSecurityResource", Required: true},
	{Name: "ScanId", Flag: "scan-id", Type: "*string", Required: true},
}

var fields_get_code_security_scan_configuration = []leanruntime.Field{
	{Name: "ScanConfigurationArn", Flag: "scan-configuration-arn", Type: "*string", Required: true},
}

var fields_get_configuration = []leanruntime.Field{}

var fields_get_delegated_admin_account = []leanruntime.Field{}

var fields_get_ec2_deep_inspection_configuration = []leanruntime.Field{}

var fields_get_encryption_key = []leanruntime.Field{
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "ScanType", Flag: "scan-type", Type: "types.ScanType", Required: true},
}

var fields_get_findings_report_status = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: false},
}

var fields_get_member = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_sbom_export = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_list_account_permissions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Service", Flag: "service", Type: "types.Service", Required: false},
}

var fields_list_cis_scan_configurations = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.ListCisScanConfigurationsFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CisScanConfigurationsSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.CisSortOrder", Required: false},
}

var fields_list_cis_scan_results_aggregated_by_checks = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CisScanResultsAggregatedByChecksFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanArn", Flag: "scan-arn", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CisScanResultsAggregatedByChecksSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.CisSortOrder", Required: false},
}

var fields_list_cis_scan_results_aggregated_by_target_resource = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CisScanResultsAggregatedByTargetResourceFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanArn", Flag: "scan-arn", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CisScanResultsAggregatedByTargetResourceSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.CisSortOrder", Required: false},
}

var fields_list_cis_scans = []leanruntime.Field{
	{Name: "DetailLevel", Flag: "detail-level", Type: "types.ListCisScansDetailLevel", Required: false},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.ListCisScansFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListCisScansSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.CisSortOrder", Required: false},
}

var fields_list_code_security_integrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_code_security_scan_configuration_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanConfigurationArn", Flag: "scan-configuration-arn", Type: "*string", Required: true},
}

var fields_list_code_security_scan_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_coverage = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CoverageFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_coverage_statistics = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CoverageFilterCriteria", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "types.GroupKey", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_delegated_admin_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_filters = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FilterAction", Required: false},
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_finding_aggregations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]types.StringFilter", Required: false},
	{Name: "AggregationRequest", Flag: "aggregation-request", Type: "types.AggregationRequest", Required: false},
	{Name: "AggregationType", Flag: "aggregation-type", Type: "types.AggregationType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_findings = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OnlyAssociated", Flag: "only-associated", Type: "*bool", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_usage_totals = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_reset_encryption_key = []leanruntime.Field{
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "ScanType", Flag: "scan-type", Type: "types.ScanType", Required: true},
}

var fields_search_vulnerabilities = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.SearchVulnerabilitiesFilterCriteria", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_send_cis_session_health = []leanruntime.Field{
	{Name: "ScanJobId", Flag: "scan-job-id", Type: "*string", Required: true},
	{Name: "SessionToken", Flag: "session-token", Type: "*string", Required: true},
}

var fields_send_cis_session_telemetry = []leanruntime.Field{
	{Name: "Messages", Flag: "messages", Type: "[]types.CisSessionMessage", Required: true},
	{Name: "ScanJobId", Flag: "scan-job-id", Type: "*string", Required: true},
	{Name: "SessionToken", Flag: "session-token", Type: "*string", Required: true},
}

var fields_start_cis_session = []leanruntime.Field{
	{Name: "Message", Flag: "message", Type: "*types.StartCisSessionMessage", Required: true},
	{Name: "ScanJobId", Flag: "scan-job-id", Type: "*string", Required: true},
}

var fields_start_code_security_scan = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Resource", Flag: "resource", Type: "types.CodeSecurityResource", Required: true},
}

var fields_stop_cis_session = []leanruntime.Field{
	{Name: "Message", Flag: "message", Type: "*types.StopCisSessionMessage", Required: true},
	{Name: "ScanJobId", Flag: "scan-job-id", Type: "*string", Required: true},
	{Name: "SessionToken", Flag: "session-token", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_cis_scan_configuration = []leanruntime.Field{
	{Name: "ScanConfigurationArn", Flag: "scan-configuration-arn", Type: "*string", Required: true},
	{Name: "ScanName", Flag: "scan-name", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "types.Schedule", Required: false},
	{Name: "SecurityLevel", Flag: "security-level", Type: "types.CisSecurityLevel", Required: false},
	{Name: "Targets", Flag: "targets", Type: "*types.UpdateCisTargets", Required: false},
}

var fields_update_code_security_integration = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "types.UpdateIntegrationDetails", Required: true},
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
}

var fields_update_code_security_scan_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.CodeSecurityScanConfiguration", Required: true},
	{Name: "ScanConfigurationArn", Flag: "scan-configuration-arn", Type: "*string", Required: true},
}

var fields_update_configuration = []leanruntime.Field{
	{Name: "Ec2Configuration", Flag: "ec2-configuration", Type: "*types.Ec2Configuration", Required: false},
	{Name: "EcrConfiguration", Flag: "ecr-configuration", Type: "*types.EcrConfiguration", Required: false},
}

var fields_update_ec2_deep_inspection_configuration = []leanruntime.Field{
	{Name: "ActivateDeepInspection", Flag: "activate-deep-inspection", Type: "*bool", Required: false},
	{Name: "PackagePaths", Flag: "package-paths", Type: "[]string", Required: false},
}

var fields_update_encryption_key = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "ScanType", Flag: "scan-type", Type: "types.ScanType", Required: true},
}

var fields_update_filter = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FilterAction", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: true},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_update_org_ec2_deep_inspection_configuration = []leanruntime.Field{
	{Name: "OrgPackagePaths", Flag: "org-package-paths", Type: "[]string", Required: true},
}

var fields_update_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnable", Flag: "auto-enable", Type: "*types.AutoEnable", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-member": {
			Name:   "associate-member",
			Fields: fields_associate_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMember(ctx, input)
			},
		},
		"batch-associate-code-security-scan-configuration": {
			Name:   "batch-associate-code-security-scan-configuration",
			Fields: fields_batch_associate_code_security_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateCodeSecurityScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_code_security_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateCodeSecurityScanConfiguration(ctx, input)
			},
		},
		"batch-disassociate-code-security-scan-configuration": {
			Name:   "batch-disassociate-code-security-scan-configuration",
			Fields: fields_batch_disassociate_code_security_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateCodeSecurityScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_code_security_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateCodeSecurityScanConfiguration(ctx, input)
			},
		},
		"batch-get-account-status": {
			Name:   "batch-get-account-status",
			Fields: fields_batch_get_account_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAccountStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_account_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetAccountStatus(ctx, input)
			},
		},
		"batch-get-code-snippet": {
			Name:   "batch-get-code-snippet",
			Fields: fields_batch_get_code_snippet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCodeSnippetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_code_snippet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCodeSnippet(ctx, input)
			},
		},
		"batch-get-finding-details": {
			Name:   "batch-get-finding-details",
			Fields: fields_batch_get_finding_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFindingDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_finding_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetFindingDetails(ctx, input)
			},
		},
		"batch-get-free-trial-info": {
			Name:   "batch-get-free-trial-info",
			Fields: fields_batch_get_free_trial_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFreeTrialInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_free_trial_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetFreeTrialInfo(ctx, input)
			},
		},
		"batch-get-member-ec2-deep-inspection-status": {
			Name:   "batch-get-member-ec2-deep-inspection-status",
			Fields: fields_batch_get_member_ec2_deep_inspection_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetMemberEc2DeepInspectionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_member_ec2_deep_inspection_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetMemberEc2DeepInspectionStatus(ctx, input)
			},
		},
		"batch-update-member-ec2-deep-inspection-status": {
			Name:   "batch-update-member-ec2-deep-inspection-status",
			Fields: fields_batch_update_member_ec2_deep_inspection_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateMemberEc2DeepInspectionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_member_ec2_deep_inspection_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateMemberEc2DeepInspectionStatus(ctx, input)
			},
		},
		"cancel-findings-report": {
			Name:   "cancel-findings-report",
			Fields: fields_cancel_findings_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelFindingsReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_findings_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelFindingsReport(ctx, input)
			},
		},
		"cancel-sbom-export": {
			Name:   "cancel-sbom-export",
			Fields: fields_cancel_sbom_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSbomExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_sbom_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSbomExport(ctx, input)
			},
		},
		"create-cis-scan-configuration": {
			Name:   "create-cis-scan-configuration",
			Fields: fields_create_cis_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCisScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cis_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCisScanConfiguration(ctx, input)
			},
		},
		"create-code-security-integration": {
			Name:   "create-code-security-integration",
			Fields: fields_create_code_security_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCodeSecurityIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_code_security_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCodeSecurityIntegration(ctx, input)
			},
		},
		"create-code-security-scan-configuration": {
			Name:   "create-code-security-scan-configuration",
			Fields: fields_create_code_security_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCodeSecurityScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_code_security_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCodeSecurityScanConfiguration(ctx, input)
			},
		},
		"create-filter": {
			Name:   "create-filter",
			Fields: fields_create_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFilter(ctx, input)
			},
		},
		"create-findings-report": {
			Name:   "create-findings-report",
			Fields: fields_create_findings_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFindingsReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_findings_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFindingsReport(ctx, input)
			},
		},
		"create-sbom-export": {
			Name:   "create-sbom-export",
			Fields: fields_create_sbom_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSbomExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sbom_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSbomExport(ctx, input)
			},
		},
		"delete-cis-scan-configuration": {
			Name:   "delete-cis-scan-configuration",
			Fields: fields_delete_cis_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCisScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cis_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCisScanConfiguration(ctx, input)
			},
		},
		"delete-code-security-integration": {
			Name:   "delete-code-security-integration",
			Fields: fields_delete_code_security_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCodeSecurityIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_code_security_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCodeSecurityIntegration(ctx, input)
			},
		},
		"delete-code-security-scan-configuration": {
			Name:   "delete-code-security-scan-configuration",
			Fields: fields_delete_code_security_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCodeSecurityScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_code_security_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCodeSecurityScanConfiguration(ctx, input)
			},
		},
		"delete-filter": {
			Name:   "delete-filter",
			Fields: fields_delete_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFilter(ctx, input)
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
		"disable": {
			Name:   "disable",
			Fields: fields_disable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Disable(ctx, input)
			},
		},
		"disable-delegated-admin-account": {
			Name:   "disable-delegated-admin-account",
			Fields: fields_disable_delegated_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDelegatedAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_delegated_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDelegatedAdminAccount(ctx, input)
			},
		},
		"disassociate-member": {
			Name:   "disassociate-member",
			Fields: fields_disassociate_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMember(ctx, input)
			},
		},
		"enable": {
			Name:   "enable",
			Fields: fields_enable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Enable(ctx, input)
			},
		},
		"enable-delegated-admin-account": {
			Name:   "enable-delegated-admin-account",
			Fields: fields_enable_delegated_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDelegatedAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_delegated_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDelegatedAdminAccount(ctx, input)
			},
		},
		"get-cis-scan-report": {
			Name:   "get-cis-scan-report",
			Fields: fields_get_cis_scan_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCisScanReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cis_scan_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCisScanReport(ctx, input)
			},
		},
		"get-cis-scan-result-details": {
			Name:   "get-cis-scan-result-details",
			Fields: fields_get_cis_scan_result_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCisScanResultDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_cis_scan_result_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCisScanResultDetails(ctx, input)
				}
				var results []*svc.GetCisScanResultDetailsOutput
				p := svc.NewGetCisScanResultDetailsPaginator(client, input)
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
		"get-clusters-for-image": {
			Name:   "get-clusters-for-image",
			Fields: fields_get_clusters_for_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClustersForImageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_clusters_for_image, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetClustersForImage(ctx, input)
				}
				var results []*svc.GetClustersForImageOutput
				p := svc.NewGetClustersForImagePaginator(client, input)
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
		"get-code-security-integration": {
			Name:   "get-code-security-integration",
			Fields: fields_get_code_security_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeSecurityIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_security_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeSecurityIntegration(ctx, input)
			},
		},
		"get-code-security-scan": {
			Name:   "get-code-security-scan",
			Fields: fields_get_code_security_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeSecurityScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_security_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeSecurityScan(ctx, input)
			},
		},
		"get-code-security-scan-configuration": {
			Name:   "get-code-security-scan-configuration",
			Fields: fields_get_code_security_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeSecurityScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_security_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeSecurityScanConfiguration(ctx, input)
			},
		},
		"get-configuration": {
			Name:   "get-configuration",
			Fields: fields_get_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguration(ctx, input)
			},
		},
		"get-delegated-admin-account": {
			Name:   "get-delegated-admin-account",
			Fields: fields_get_delegated_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDelegatedAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delegated_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDelegatedAdminAccount(ctx, input)
			},
		},
		"get-ec2-deep-inspection-configuration": {
			Name:   "get-ec2-deep-inspection-configuration",
			Fields: fields_get_ec2_deep_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEc2DeepInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ec2_deep_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEc2DeepInspectionConfiguration(ctx, input)
			},
		},
		"get-encryption-key": {
			Name:   "get-encryption-key",
			Fields: fields_get_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEncryptionKey(ctx, input)
			},
		},
		"get-findings-report-status": {
			Name:   "get-findings-report-status",
			Fields: fields_get_findings_report_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsReportStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_findings_report_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingsReportStatus(ctx, input)
			},
		},
		"get-member": {
			Name:   "get-member",
			Fields: fields_get_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMember(ctx, input)
			},
		},
		"get-sbom-export": {
			Name:   "get-sbom-export",
			Fields: fields_get_sbom_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSbomExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sbom_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSbomExport(ctx, input)
			},
		},
		"list-account-permissions": {
			Name:   "list-account-permissions",
			Fields: fields_list_account_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountPermissions(ctx, input)
				}
				var results []*svc.ListAccountPermissionsOutput
				p := svc.NewListAccountPermissionsPaginator(client, input)
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
		"list-cis-scan-configurations": {
			Name:   "list-cis-scan-configurations",
			Fields: fields_list_cis_scan_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCisScanConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cis_scan_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCisScanConfigurations(ctx, input)
				}
				var results []*svc.ListCisScanConfigurationsOutput
				p := svc.NewListCisScanConfigurationsPaginator(client, input)
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
		"list-cis-scan-results-aggregated-by-checks": {
			Name:   "list-cis-scan-results-aggregated-by-checks",
			Fields: fields_list_cis_scan_results_aggregated_by_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCisScanResultsAggregatedByChecksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cis_scan_results_aggregated_by_checks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCisScanResultsAggregatedByChecks(ctx, input)
				}
				var results []*svc.ListCisScanResultsAggregatedByChecksOutput
				p := svc.NewListCisScanResultsAggregatedByChecksPaginator(client, input)
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
		"list-cis-scan-results-aggregated-by-target-resource": {
			Name:   "list-cis-scan-results-aggregated-by-target-resource",
			Fields: fields_list_cis_scan_results_aggregated_by_target_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCisScanResultsAggregatedByTargetResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cis_scan_results_aggregated_by_target_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCisScanResultsAggregatedByTargetResource(ctx, input)
				}
				var results []*svc.ListCisScanResultsAggregatedByTargetResourceOutput
				p := svc.NewListCisScanResultsAggregatedByTargetResourcePaginator(client, input)
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
		"list-cis-scans": {
			Name:   "list-cis-scans",
			Fields: fields_list_cis_scans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCisScansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cis_scans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCisScans(ctx, input)
				}
				var results []*svc.ListCisScansOutput
				p := svc.NewListCisScansPaginator(client, input)
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
		"list-code-security-integrations": {
			Name:   "list-code-security-integrations",
			Fields: fields_list_code_security_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeSecurityIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_code_security_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCodeSecurityIntegrations(ctx, input)
			},
		},
		"list-code-security-scan-configuration-associations": {
			Name:   "list-code-security-scan-configuration-associations",
			Fields: fields_list_code_security_scan_configuration_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeSecurityScanConfigurationAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_code_security_scan_configuration_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCodeSecurityScanConfigurationAssociations(ctx, input)
			},
		},
		"list-code-security-scan-configurations": {
			Name:   "list-code-security-scan-configurations",
			Fields: fields_list_code_security_scan_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeSecurityScanConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_code_security_scan_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCodeSecurityScanConfigurations(ctx, input)
			},
		},
		"list-coverage": {
			Name:   "list-coverage",
			Fields: fields_list_coverage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoverageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_coverage, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoverage(ctx, input)
				}
				var results []*svc.ListCoverageOutput
				p := svc.NewListCoveragePaginator(client, input)
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
		"list-coverage-statistics": {
			Name:   "list-coverage-statistics",
			Fields: fields_list_coverage_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoverageStatisticsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_coverage_statistics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoverageStatistics(ctx, input)
				}
				var results []*svc.ListCoverageStatisticsOutput
				p := svc.NewListCoverageStatisticsPaginator(client, input)
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
		"list-delegated-admin-accounts": {
			Name:   "list-delegated-admin-accounts",
			Fields: fields_list_delegated_admin_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDelegatedAdminAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_delegated_admin_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDelegatedAdminAccounts(ctx, input)
				}
				var results []*svc.ListDelegatedAdminAccountsOutput
				p := svc.NewListDelegatedAdminAccountsPaginator(client, input)
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
		"list-filters": {
			Name:   "list-filters",
			Fields: fields_list_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFilters(ctx, input)
				}
				var results []*svc.ListFiltersOutput
				p := svc.NewListFiltersPaginator(client, input)
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
		"list-finding-aggregations": {
			Name:   "list-finding-aggregations",
			Fields: fields_list_finding_aggregations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingAggregationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_finding_aggregations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindingAggregations(ctx, input)
				}
				var results []*svc.ListFindingAggregationsOutput
				p := svc.NewListFindingAggregationsPaginator(client, input)
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
		"list-findings": {
			Name:   "list-findings",
			Fields: fields_list_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindings(ctx, input)
				}
				var results []*svc.ListFindingsOutput
				p := svc.NewListFindingsPaginator(client, input)
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
		"list-usage-totals": {
			Name:   "list-usage-totals",
			Fields: fields_list_usage_totals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsageTotalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_usage_totals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsageTotals(ctx, input)
				}
				var results []*svc.ListUsageTotalsOutput
				p := svc.NewListUsageTotalsPaginator(client, input)
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
		"reset-encryption-key": {
			Name:   "reset-encryption-key",
			Fields: fields_reset_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetEncryptionKey(ctx, input)
			},
		},
		"search-vulnerabilities": {
			Name:   "search-vulnerabilities",
			Fields: fields_search_vulnerabilities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchVulnerabilitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_vulnerabilities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchVulnerabilities(ctx, input)
				}
				var results []*svc.SearchVulnerabilitiesOutput
				p := svc.NewSearchVulnerabilitiesPaginator(client, input)
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
		"send-cis-session-health": {
			Name:   "send-cis-session-health",
			Fields: fields_send_cis_session_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendCisSessionHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_cis_session_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendCisSessionHealth(ctx, input)
			},
		},
		"send-cis-session-telemetry": {
			Name:   "send-cis-session-telemetry",
			Fields: fields_send_cis_session_telemetry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendCisSessionTelemetryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_cis_session_telemetry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendCisSessionTelemetry(ctx, input)
			},
		},
		"start-cis-session": {
			Name:   "start-cis-session",
			Fields: fields_start_cis_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCisSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_cis_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCisSession(ctx, input)
			},
		},
		"start-code-security-scan": {
			Name:   "start-code-security-scan",
			Fields: fields_start_code_security_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCodeSecurityScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_code_security_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCodeSecurityScan(ctx, input)
			},
		},
		"stop-cis-session": {
			Name:   "stop-cis-session",
			Fields: fields_stop_cis_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCisSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_cis_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCisSession(ctx, input)
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
		"update-cis-scan-configuration": {
			Name:   "update-cis-scan-configuration",
			Fields: fields_update_cis_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCisScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cis_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCisScanConfiguration(ctx, input)
			},
		},
		"update-code-security-integration": {
			Name:   "update-code-security-integration",
			Fields: fields_update_code_security_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCodeSecurityIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_code_security_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCodeSecurityIntegration(ctx, input)
			},
		},
		"update-code-security-scan-configuration": {
			Name:   "update-code-security-scan-configuration",
			Fields: fields_update_code_security_scan_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCodeSecurityScanConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_code_security_scan_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCodeSecurityScanConfiguration(ctx, input)
			},
		},
		"update-configuration": {
			Name:   "update-configuration",
			Fields: fields_update_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguration(ctx, input)
			},
		},
		"update-ec2-deep-inspection-configuration": {
			Name:   "update-ec2-deep-inspection-configuration",
			Fields: fields_update_ec2_deep_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEc2DeepInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ec2_deep_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEc2DeepInspectionConfiguration(ctx, input)
			},
		},
		"update-encryption-key": {
			Name:   "update-encryption-key",
			Fields: fields_update_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEncryptionKey(ctx, input)
			},
		},
		"update-filter": {
			Name:   "update-filter",
			Fields: fields_update_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFilter(ctx, input)
			},
		},
		"update-org-ec2-deep-inspection-configuration": {
			Name:   "update-org-ec2-deep-inspection-configuration",
			Fields: fields_update_org_ec2_deep_inspection_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOrgEc2DeepInspectionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_org_ec2_deep_inspection_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOrgEc2DeepInspectionConfiguration(ctx, input)
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
	}
	if err := leanruntime.Execute("inspector2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
