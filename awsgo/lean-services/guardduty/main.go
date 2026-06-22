package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/guardduty"
)

var fields_accept_administrator_invitation = []leanruntime.Field{
	{Name: "AdministratorId", Flag: "administrator-id", Type: "*string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
}

var fields_accept_invitation = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
	{Name: "MasterId", Flag: "master-id", Type: "*string", Required: true},
}

var fields_archive_findings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingIds", Flag: "finding-ids", Type: "[]string", Required: true},
}

var fields_create_detector = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSources", Flag: "data-sources", Type: "*types.DataSourceConfigurations", Required: false},
	{Name: "Enable", Flag: "enable", Type: "*bool", Required: true},
	{Name: "Features", Flag: "features", Type: "[]types.DetectorFeatureConfiguration", Required: false},
	{Name: "FindingPublishingFrequency", Flag: "finding-publishing-frequency", Type: "types.FindingPublishingFrequency", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_filter = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FilterAction", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rank", Flag: "rank", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_ip_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.IpSetFormat", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_malware_protection_plan = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "*types.MalwareProtectionPlanActions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProtectedResource", Flag: "protected-resource", Type: "*types.CreateProtectedResource", Required: true},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_members = []leanruntime.Field{
	{Name: "AccountDetails", Flag: "account-details", Type: "[]types.AccountDetail", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_create_publishing_destination = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationProperties", Flag: "destination-properties", Type: "*types.DestinationProperties", Required: true},
	{Name: "DestinationType", Flag: "destination-type", Type: "types.DestinationType", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sample_findings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingTypes", Flag: "finding-types", Type: "[]string", Required: false},
}

var fields_create_threat_entity_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.ThreatEntitySetFormat", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_threat_intel_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.ThreatIntelSetFormat", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_trusted_entity_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.TrustedEntitySetFormat", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_decline_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_delete_detector = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_delete_filter = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
}

var fields_delete_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_delete_ip_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "IpSetId", Flag: "ip-set-id", Type: "*string", Required: true},
}

var fields_delete_malware_protection_plan = []leanruntime.Field{
	{Name: "MalwareProtectionPlanId", Flag: "malware-protection-plan-id", Type: "*string", Required: true},
}

var fields_delete_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_delete_publishing_destination = []leanruntime.Field{
	{Name: "DestinationId", Flag: "destination-id", Type: "*string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_delete_threat_entity_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ThreatEntitySetId", Flag: "threat-entity-set-id", Type: "*string", Required: true},
}

var fields_delete_threat_intel_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ThreatIntelSetId", Flag: "threat-intel-set-id", Type: "*string", Required: true},
}

var fields_delete_trusted_entity_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "TrustedEntitySetId", Flag: "trusted-entity-set-id", Type: "*string", Required: true},
}

var fields_describe_malware_scans = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.FilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_describe_organization_configuration = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_publishing_destination = []leanruntime.Field{
	{Name: "DestinationId", Flag: "destination-id", Type: "*string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_disable_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
}

var fields_disassociate_from_administrator_account = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_disassociate_from_master_account = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_disassociate_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_enable_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
}

var fields_get_administrator_account = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_coverage_statistics = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CoverageFilterCriteria", Required: false},
	{Name: "StatisticsType", Flag: "statistics-type", Type: "[]types.CoverageStatisticsType", Required: true},
}

var fields_get_detector = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_filter = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
}

var fields_get_findings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingIds", Flag: "finding-ids", Type: "[]string", Required: true},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_get_findings_statistics = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: false},
	{Name: "FindingStatisticTypes", Flag: "finding-statistic-types", Type: "[]types.FindingStatisticType", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "types.GroupByType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "types.OrderBy", Required: false},
}

var fields_get_invitations_count = []leanruntime.Field{}

var fields_get_ip_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "IpSetId", Flag: "ip-set-id", Type: "*string", Required: true},
}

var fields_get_malware_protection_plan = []leanruntime.Field{
	{Name: "MalwareProtectionPlanId", Flag: "malware-protection-plan-id", Type: "*string", Required: true},
}

var fields_get_malware_scan = []leanruntime.Field{
	{Name: "ScanId", Flag: "scan-id", Type: "*string", Required: true},
}

var fields_get_malware_scan_settings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_master_account = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_member_detectors = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_organization_statistics = []leanruntime.Field{}

var fields_get_remaining_free_trial_days = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_get_threat_entity_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ThreatEntitySetId", Flag: "threat-entity-set-id", Type: "*string", Required: true},
}

var fields_get_threat_intel_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ThreatIntelSetId", Flag: "threat-intel-set-id", Type: "*string", Required: true},
}

var fields_get_trusted_entity_set = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "TrustedEntitySetId", Flag: "trusted-entity-set-id", Type: "*string", Required: true},
}

var fields_get_usage_statistics = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Unit", Flag: "unit", Type: "*string", Required: false},
	{Name: "UsageCriteria", Flag: "usage-criteria", Type: "*types.UsageCriteria", Required: true},
	{Name: "UsageStatisticType", Flag: "usage-statistic-type", Type: "types.UsageStatisticType", Required: true},
}

var fields_invite_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DisableEmailNotification", Flag: "disable-email-notification", Type: "*bool", Required: false},
	{Name: "Message", Flag: "message", Type: "*string", Required: false},
}

var fields_list_coverage = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.CoverageFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.CoverageSortCriteria", Required: false},
}

var fields_list_detectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_filters = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_findings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_list_invitations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ip_sets = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_malware_protection_plans = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_malware_scans = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.ListMalwareScansFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OnlyAssociated", Flag: "only-associated", Type: "*string", Required: false},
}

var fields_list_organization_admin_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_publishing_destinations = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_threat_entity_sets = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_threat_intel_sets = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_trusted_entity_sets = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_send_object_malware_scan = []leanruntime.Field{
	{Name: "S3Object", Flag: "s3-object", Type: "*types.S3ObjectForSendObjectMalwareScan", Required: false},
}

var fields_start_malware_scan = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ScanConfiguration", Flag: "scan-configuration", Type: "*types.StartMalwareScanConfiguration", Required: false},
}

var fields_start_monitoring_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_stop_monitoring_members = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_unarchive_findings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FindingIds", Flag: "finding-ids", Type: "[]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_detector = []leanruntime.Field{
	{Name: "DataSources", Flag: "data-sources", Type: "*types.DataSourceConfigurations", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "Enable", Flag: "enable", Type: "*bool", Required: false},
	{Name: "Features", Flag: "features", Type: "[]types.DetectorFeatureConfiguration", Required: false},
	{Name: "FindingPublishingFrequency", Flag: "finding-publishing-frequency", Type: "types.FindingPublishingFrequency", Required: false},
}

var fields_update_filter = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FilterAction", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: false},
	{Name: "Rank", Flag: "rank", Type: "*int32", Required: false},
}

var fields_update_findings_feedback = []leanruntime.Field{
	{Name: "Comments", Flag: "comments", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "Feedback", Flag: "feedback", Type: "types.Feedback", Required: true},
	{Name: "FindingIds", Flag: "finding-ids", Type: "[]string", Required: true},
}

var fields_update_ip_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "IpSetId", Flag: "ip-set-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_malware_protection_plan = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "*types.MalwareProtectionPlanActions", Required: false},
	{Name: "MalwareProtectionPlanId", Flag: "malware-protection-plan-id", Type: "*string", Required: true},
	{Name: "ProtectedResource", Flag: "protected-resource", Type: "*types.UpdateProtectedResource", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
}

var fields_update_malware_scan_settings = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "EbsSnapshotPreservation", Flag: "ebs-snapshot-preservation", Type: "types.EbsSnapshotPreservation", Required: false},
	{Name: "ScanResourceCriteria", Flag: "scan-resource-criteria", Type: "*types.ScanResourceCriteria", Required: false},
}

var fields_update_member_detectors = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DataSources", Flag: "data-sources", Type: "*types.DataSourceConfigurations", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "Features", Flag: "features", Type: "[]types.MemberFeaturesConfiguration", Required: false},
}

var fields_update_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnable", Flag: "auto-enable", Type: "*bool", Required: false},
	{Name: "AutoEnableOrganizationMembers", Flag: "auto-enable-organization-members", Type: "types.AutoEnableMembers", Required: false},
	{Name: "DataSources", Flag: "data-sources", Type: "*types.OrganizationDataSourceConfigurations", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "Features", Flag: "features", Type: "[]types.OrganizationFeatureConfiguration", Required: false},
}

var fields_update_publishing_destination = []leanruntime.Field{
	{Name: "DestinationId", Flag: "destination-id", Type: "*string", Required: true},
	{Name: "DestinationProperties", Flag: "destination-properties", Type: "*types.DestinationProperties", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_update_threat_entity_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ThreatEntitySetId", Flag: "threat-entity-set-id", Type: "*string", Required: true},
}

var fields_update_threat_intel_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ThreatIntelSetId", Flag: "threat-intel-set-id", Type: "*string", Required: true},
}

var fields_update_trusted_entity_set = []leanruntime.Field{
	{Name: "Activate", Flag: "activate", Type: "*bool", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TrustedEntitySetId", Flag: "trusted-entity-set-id", Type: "*string", Required: true},
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
		"archive-findings": {
			Name:   "archive-findings",
			Fields: fields_archive_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ArchiveFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_archive_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ArchiveFindings(ctx, input)
			},
		},
		"create-detector": {
			Name:   "create-detector",
			Fields: fields_create_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDetector(ctx, input)
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
		"create-ip-set": {
			Name:   "create-ip-set",
			Fields: fields_create_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIPSet(ctx, input)
			},
		},
		"create-malware-protection-plan": {
			Name:   "create-malware-protection-plan",
			Fields: fields_create_malware_protection_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMalwareProtectionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_malware_protection_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMalwareProtectionPlan(ctx, input)
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
		"create-publishing-destination": {
			Name:   "create-publishing-destination",
			Fields: fields_create_publishing_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePublishingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_publishing_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePublishingDestination(ctx, input)
			},
		},
		"create-sample-findings": {
			Name:   "create-sample-findings",
			Fields: fields_create_sample_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSampleFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sample_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSampleFindings(ctx, input)
			},
		},
		"create-threat-entity-set": {
			Name:   "create-threat-entity-set",
			Fields: fields_create_threat_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThreatEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_threat_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThreatEntitySet(ctx, input)
			},
		},
		"create-threat-intel-set": {
			Name:   "create-threat-intel-set",
			Fields: fields_create_threat_intel_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThreatIntelSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_threat_intel_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThreatIntelSet(ctx, input)
			},
		},
		"create-trusted-entity-set": {
			Name:   "create-trusted-entity-set",
			Fields: fields_create_trusted_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrustedEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trusted_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrustedEntitySet(ctx, input)
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
		"delete-detector": {
			Name:   "delete-detector",
			Fields: fields_delete_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDetector(ctx, input)
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
		"delete-ip-set": {
			Name:   "delete-ip-set",
			Fields: fields_delete_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIPSet(ctx, input)
			},
		},
		"delete-malware-protection-plan": {
			Name:   "delete-malware-protection-plan",
			Fields: fields_delete_malware_protection_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMalwareProtectionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_malware_protection_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMalwareProtectionPlan(ctx, input)
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
		"delete-publishing-destination": {
			Name:   "delete-publishing-destination",
			Fields: fields_delete_publishing_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePublishingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_publishing_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePublishingDestination(ctx, input)
			},
		},
		"delete-threat-entity-set": {
			Name:   "delete-threat-entity-set",
			Fields: fields_delete_threat_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThreatEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_threat_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThreatEntitySet(ctx, input)
			},
		},
		"delete-threat-intel-set": {
			Name:   "delete-threat-intel-set",
			Fields: fields_delete_threat_intel_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThreatIntelSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_threat_intel_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThreatIntelSet(ctx, input)
			},
		},
		"delete-trusted-entity-set": {
			Name:   "delete-trusted-entity-set",
			Fields: fields_delete_trusted_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrustedEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trusted_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrustedEntitySet(ctx, input)
			},
		},
		"describe-malware-scans": {
			Name:   "describe-malware-scans",
			Fields: fields_describe_malware_scans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMalwareScansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_malware_scans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMalwareScans(ctx, input)
				}
				var results []*svc.DescribeMalwareScansOutput
				p := svc.NewDescribeMalwareScansPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-organization-configuration": {
			Name:   "describe-organization-configuration",
			Fields: fields_describe_organization_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationConfigurationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_organization_configuration, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOrganizationConfiguration(ctx, input)
				}
				var results []*svc.DescribeOrganizationConfigurationOutput
				p := svc.NewDescribeOrganizationConfigurationPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-publishing-destination": {
			Name:   "describe-publishing-destination",
			Fields: fields_describe_publishing_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePublishingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_publishing_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePublishingDestination(ctx, input)
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
		"get-coverage-statistics": {
			Name:   "get-coverage-statistics",
			Fields: fields_get_coverage_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoverageStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_coverage_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoverageStatistics(ctx, input)
			},
		},
		"get-detector": {
			Name:   "get-detector",
			Fields: fields_get_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDetector(ctx, input)
			},
		},
		"get-filter": {
			Name:   "get-filter",
			Fields: fields_get_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFilter(ctx, input)
			},
		},
		"get-findings": {
			Name:   "get-findings",
			Fields: fields_get_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindings(ctx, input)
			},
		},
		"get-findings-statistics": {
			Name:   "get-findings-statistics",
			Fields: fields_get_findings_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_findings_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingsStatistics(ctx, input)
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
		"get-ip-set": {
			Name:   "get-ip-set",
			Fields: fields_get_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIPSet(ctx, input)
			},
		},
		"get-malware-protection-plan": {
			Name:   "get-malware-protection-plan",
			Fields: fields_get_malware_protection_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMalwareProtectionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_malware_protection_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMalwareProtectionPlan(ctx, input)
			},
		},
		"get-malware-scan": {
			Name:   "get-malware-scan",
			Fields: fields_get_malware_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMalwareScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_malware_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMalwareScan(ctx, input)
			},
		},
		"get-malware-scan-settings": {
			Name:   "get-malware-scan-settings",
			Fields: fields_get_malware_scan_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMalwareScanSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_malware_scan_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMalwareScanSettings(ctx, input)
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
		"get-member-detectors": {
			Name:   "get-member-detectors",
			Fields: fields_get_member_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMemberDetectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_member_detectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMemberDetectors(ctx, input)
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
		"get-organization-statistics": {
			Name:   "get-organization-statistics",
			Fields: fields_get_organization_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_organization_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOrganizationStatistics(ctx, input)
			},
		},
		"get-remaining-free-trial-days": {
			Name:   "get-remaining-free-trial-days",
			Fields: fields_get_remaining_free_trial_days,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRemainingFreeTrialDaysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_remaining_free_trial_days, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRemainingFreeTrialDays(ctx, input)
			},
		},
		"get-threat-entity-set": {
			Name:   "get-threat-entity-set",
			Fields: fields_get_threat_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThreatEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_threat_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThreatEntitySet(ctx, input)
			},
		},
		"get-threat-intel-set": {
			Name:   "get-threat-intel-set",
			Fields: fields_get_threat_intel_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThreatIntelSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_threat_intel_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThreatIntelSet(ctx, input)
			},
		},
		"get-trusted-entity-set": {
			Name:   "get-trusted-entity-set",
			Fields: fields_get_trusted_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustedEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trusted_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustedEntitySet(ctx, input)
			},
		},
		"get-usage-statistics": {
			Name:   "get-usage-statistics",
			Fields: fields_get_usage_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsageStatisticsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_usage_statistics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUsageStatistics(ctx, input)
				}
				var results []*svc.GetUsageStatisticsOutput
				p := svc.NewGetUsageStatisticsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
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
		"list-detectors": {
			Name:   "list-detectors",
			Fields: fields_list_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDetectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_detectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDetectors(ctx, input)
				}
				var results []*svc.ListDetectorsOutput
				p := svc.NewListDetectorsPaginator(client, input)
				for p.HasMorePages() {
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
		"list-ip-sets": {
			Name:   "list-ip-sets",
			Fields: fields_list_ip_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIPSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ip_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIPSets(ctx, input)
				}
				var results []*svc.ListIPSetsOutput
				p := svc.NewListIPSetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-malware-protection-plans": {
			Name:   "list-malware-protection-plans",
			Fields: fields_list_malware_protection_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMalwareProtectionPlansInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_malware_protection_plans, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMalwareProtectionPlans(ctx, input)
			},
		},
		"list-malware-scans": {
			Name:   "list-malware-scans",
			Fields: fields_list_malware_scans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMalwareScansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_malware_scans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMalwareScans(ctx, input)
				}
				var results []*svc.ListMalwareScansOutput
				p := svc.NewListMalwareScansPaginator(client, input)
				for p.HasMorePages() {
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
		"list-publishing-destinations": {
			Name:   "list-publishing-destinations",
			Fields: fields_list_publishing_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPublishingDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_publishing_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPublishingDestinations(ctx, input)
				}
				var results []*svc.ListPublishingDestinationsOutput
				p := svc.NewListPublishingDestinationsPaginator(client, input)
				for p.HasMorePages() {
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
		"list-threat-entity-sets": {
			Name:   "list-threat-entity-sets",
			Fields: fields_list_threat_entity_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThreatEntitySetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_threat_entity_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThreatEntitySets(ctx, input)
				}
				var results []*svc.ListThreatEntitySetsOutput
				p := svc.NewListThreatEntitySetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-threat-intel-sets": {
			Name:   "list-threat-intel-sets",
			Fields: fields_list_threat_intel_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThreatIntelSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_threat_intel_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThreatIntelSets(ctx, input)
				}
				var results []*svc.ListThreatIntelSetsOutput
				p := svc.NewListThreatIntelSetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-trusted-entity-sets": {
			Name:   "list-trusted-entity-sets",
			Fields: fields_list_trusted_entity_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrustedEntitySetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trusted_entity_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrustedEntitySets(ctx, input)
				}
				var results []*svc.ListTrustedEntitySetsOutput
				p := svc.NewListTrustedEntitySetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"send-object-malware-scan": {
			Name:   "send-object-malware-scan",
			Fields: fields_send_object_malware_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendObjectMalwareScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_object_malware_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendObjectMalwareScan(ctx, input)
			},
		},
		"start-malware-scan": {
			Name:   "start-malware-scan",
			Fields: fields_start_malware_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMalwareScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_malware_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMalwareScan(ctx, input)
			},
		},
		"start-monitoring-members": {
			Name:   "start-monitoring-members",
			Fields: fields_start_monitoring_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMonitoringMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_monitoring_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMonitoringMembers(ctx, input)
			},
		},
		"stop-monitoring-members": {
			Name:   "stop-monitoring-members",
			Fields: fields_stop_monitoring_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMonitoringMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_monitoring_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMonitoringMembers(ctx, input)
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
		"unarchive-findings": {
			Name:   "unarchive-findings",
			Fields: fields_unarchive_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnarchiveFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unarchive_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnarchiveFindings(ctx, input)
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
		"update-detector": {
			Name:   "update-detector",
			Fields: fields_update_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDetector(ctx, input)
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
		"update-findings-feedback": {
			Name:   "update-findings-feedback",
			Fields: fields_update_findings_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFindingsFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_findings_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFindingsFeedback(ctx, input)
			},
		},
		"update-ip-set": {
			Name:   "update-ip-set",
			Fields: fields_update_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIPSet(ctx, input)
			},
		},
		"update-malware-protection-plan": {
			Name:   "update-malware-protection-plan",
			Fields: fields_update_malware_protection_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMalwareProtectionPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_malware_protection_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMalwareProtectionPlan(ctx, input)
			},
		},
		"update-malware-scan-settings": {
			Name:   "update-malware-scan-settings",
			Fields: fields_update_malware_scan_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMalwareScanSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_malware_scan_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMalwareScanSettings(ctx, input)
			},
		},
		"update-member-detectors": {
			Name:   "update-member-detectors",
			Fields: fields_update_member_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMemberDetectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_member_detectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMemberDetectors(ctx, input)
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
		"update-publishing-destination": {
			Name:   "update-publishing-destination",
			Fields: fields_update_publishing_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePublishingDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_publishing_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePublishingDestination(ctx, input)
			},
		},
		"update-threat-entity-set": {
			Name:   "update-threat-entity-set",
			Fields: fields_update_threat_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThreatEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_threat_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThreatEntitySet(ctx, input)
			},
		},
		"update-threat-intel-set": {
			Name:   "update-threat-intel-set",
			Fields: fields_update_threat_intel_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThreatIntelSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_threat_intel_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThreatIntelSet(ctx, input)
			},
		},
		"update-trusted-entity-set": {
			Name:   "update-trusted-entity-set",
			Fields: fields_update_trusted_entity_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrustedEntitySetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trusted_entity_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrustedEntitySet(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("guardduty", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
