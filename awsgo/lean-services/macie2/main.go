package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/macie2"
)

var fields_accept_invitation = []leanruntime.Field{
	{Name: "AdministratorAccountId", Flag: "administrator-account-id", Type: "*string", Required: false},
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
	{Name: "MasterAccount", Flag: "master-account", Type: "*string", Required: false},
}

var fields_batch_get_custom_data_identifiers = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
}

var fields_batch_update_automated_discovery_accounts = []leanruntime.Field{
	{Name: "Accounts", Flag: "accounts", Type: "[]types.AutomatedDiscoveryAccountUpdate", Required: false},
}

var fields_create_allow_list = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Criteria", Flag: "criteria", Type: "*types.AllowListCriteria", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_classification_job = []leanruntime.Field{
	{Name: "AllowListIds", Flag: "allow-list-ids", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "CustomDataIdentifierIds", Flag: "custom-data-identifier-ids", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InitialRun", Flag: "initial-run", Type: "*bool", Required: false},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: true},
	{Name: "ManagedDataIdentifierIds", Flag: "managed-data-identifier-ids", Type: "[]string", Required: false},
	{Name: "ManagedDataIdentifierSelector", Flag: "managed-data-identifier-selector", Type: "types.ManagedDataIdentifierSelector", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "S3JobDefinition", Flag: "s3-job-definition", Type: "*types.S3JobDefinition", Required: true},
	{Name: "SamplingPercentage", Flag: "sampling-percentage", Type: "*int32", Required: false},
	{Name: "ScheduleFrequency", Flag: "schedule-frequency", Type: "*types.JobScheduleFrequency", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_custom_data_identifier = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IgnoreWords", Flag: "ignore-words", Type: "[]string", Required: false},
	{Name: "Keywords", Flag: "keywords", Type: "[]string", Required: false},
	{Name: "MaximumMatchDistance", Flag: "maximum-match-distance", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Regex", Flag: "regex", Type: "*string", Required: true},
	{Name: "SeverityLevels", Flag: "severity-levels", Type: "[]types.SeverityLevel", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_findings_filter = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FindingsFilterAction", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Position", Flag: "position", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "DisableEmailNotification", Flag: "disable-email-notification", Type: "*bool", Required: false},
	{Name: "Message", Flag: "message", Type: "*string", Required: false},
}

var fields_create_member = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*types.AccountDetail", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sample_findings = []leanruntime.Field{
	{Name: "FindingTypes", Flag: "finding-types", Type: "[]types.FindingType", Required: false},
}

var fields_decline_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_delete_allow_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IgnoreJobChecks", Flag: "ignore-job-checks", Type: "*string", Required: false},
}

var fields_delete_custom_data_identifier = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_findings_filter = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_invitations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_delete_member = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_buckets = []leanruntime.Field{
	{Name: "Criteria", Flag: "criteria", Type: "map[string]types.BucketCriteriaAdditionalProperties", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.BucketSortCriteria", Required: false},
}

var fields_describe_classification_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_organization_configuration = []leanruntime.Field{}

var fields_disable_macie = []leanruntime.Field{}

var fields_disable_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
}

var fields_disassociate_from_administrator_account = []leanruntime.Field{}

var fields_disassociate_from_master_account = []leanruntime.Field{}

var fields_disassociate_member = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_enable_macie = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FindingPublishingFrequency", Flag: "finding-publishing-frequency", Type: "types.FindingPublishingFrequency", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MacieStatus", Required: false},
}

var fields_enable_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_get_administrator_account = []leanruntime.Field{}

var fields_get_allow_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_automated_discovery_configuration = []leanruntime.Field{}

var fields_get_bucket_statistics = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
}

var fields_get_classification_export_configuration = []leanruntime.Field{}

var fields_get_classification_scope = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_custom_data_identifier = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_finding_statistics = []leanruntime.Field{
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "types.GroupBy", Required: true},
	{Name: "Size", Flag: "size", Type: "*int32", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.FindingStatisticsSortCriteria", Required: false},
}

var fields_get_findings = []leanruntime.Field{
	{Name: "FindingIds", Flag: "finding-ids", Type: "[]string", Required: true},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_get_findings_filter = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_findings_publication_configuration = []leanruntime.Field{}

var fields_get_invitations_count = []leanruntime.Field{}

var fields_get_macie_session = []leanruntime.Field{}

var fields_get_master_account = []leanruntime.Field{}

var fields_get_member = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_resource_profile = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_reveal_configuration = []leanruntime.Field{}

var fields_get_sensitive_data_occurrences = []leanruntime.Field{
	{Name: "FindingId", Flag: "finding-id", Type: "*string", Required: true},
}

var fields_get_sensitive_data_occurrences_availability = []leanruntime.Field{
	{Name: "FindingId", Flag: "finding-id", Type: "*string", Required: true},
}

var fields_get_sensitivity_inspection_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_usage_statistics = []leanruntime.Field{
	{Name: "FilterBy", Flag: "filter-by", Type: "[]types.UsageStatisticsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.UsageStatisticsSortBy", Required: false},
	{Name: "TimeRange", Flag: "time-range", Type: "types.TimeRange", Required: false},
}

var fields_get_usage_totals = []leanruntime.Field{
	{Name: "TimeRange", Flag: "time-range", Type: "*string", Required: false},
}

var fields_list_allow_lists = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_automated_discovery_accounts = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_classification_jobs = []leanruntime.Field{
	{Name: "FilterCriteria", Flag: "filter-criteria", Type: "*types.ListJobsFilterCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.ListJobsSortCriteria", Required: false},
}

var fields_list_classification_scopes = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_data_identifiers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_findings = []leanruntime.Field{
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SortCriteria", Required: false},
}

var fields_list_findings_filters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_invitations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_data_identifiers = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OnlyAssociated", Flag: "only-associated", Type: "*string", Required: false},
}

var fields_list_organization_admin_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_profile_artifacts = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_resource_profile_detections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_sensitivity_inspection_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_classification_export_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.ClassificationExportConfiguration", Required: true},
}

var fields_put_findings_publication_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SecurityHubConfiguration", Flag: "security-hub-configuration", Type: "*types.SecurityHubConfiguration", Required: false},
}

var fields_search_resources = []leanruntime.Field{
	{Name: "BucketCriteria", Flag: "bucket-criteria", Type: "*types.SearchResourcesBucketCriteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "*types.SearchResourcesSortCriteria", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_test_custom_data_identifier = []leanruntime.Field{
	{Name: "IgnoreWords", Flag: "ignore-words", Type: "[]string", Required: false},
	{Name: "Keywords", Flag: "keywords", Type: "[]string", Required: false},
	{Name: "MaximumMatchDistance", Flag: "maximum-match-distance", Type: "*int32", Required: false},
	{Name: "Regex", Flag: "regex", Type: "*string", Required: true},
	{Name: "SampleText", Flag: "sample-text", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_allow_list = []leanruntime.Field{
	{Name: "Criteria", Flag: "criteria", Type: "*types.AllowListCriteria", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_automated_discovery_configuration = []leanruntime.Field{
	{Name: "AutoEnableOrganizationMembers", Flag: "auto-enable-organization-members", Type: "types.AutoEnableMode", Required: false},
	{Name: "Status", Flag: "status", Type: "types.AutomatedDiscoveryStatus", Required: true},
}

var fields_update_classification_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "JobStatus", Flag: "job-status", Type: "types.JobStatus", Required: true},
}

var fields_update_classification_scope = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "S3", Flag: "s3", Type: "*types.S3ClassificationScopeUpdate", Required: false},
}

var fields_update_findings_filter = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.FindingsFilterAction", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FindingCriteria", Flag: "finding-criteria", Type: "*types.FindingCriteria", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Position", Flag: "position", Type: "*int32", Required: false},
}

var fields_update_macie_session = []leanruntime.Field{
	{Name: "FindingPublishingFrequency", Flag: "finding-publishing-frequency", Type: "types.FindingPublishingFrequency", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MacieStatus", Required: false},
}

var fields_update_member_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.MacieStatus", Required: true},
}

var fields_update_organization_configuration = []leanruntime.Field{
	{Name: "AutoEnable", Flag: "auto-enable", Type: "*bool", Required: true},
}

var fields_update_resource_profile = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SensitivityScoreOverride", Flag: "sensitivity-score-override", Type: "*int32", Required: false},
}

var fields_update_resource_profile_detections = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SuppressDataIdentifiers", Flag: "suppress-data-identifiers", Type: "[]types.SuppressDataIdentifier", Required: false},
}

var fields_update_reveal_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.RevealConfiguration", Required: true},
	{Name: "RetrievalConfiguration", Flag: "retrieval-configuration", Type: "*types.UpdateRetrievalConfiguration", Required: false},
}

var fields_update_sensitivity_inspection_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Excludes", Flag: "excludes", Type: "*types.SensitivityInspectionTemplateExcludes", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Includes", Flag: "includes", Type: "*types.SensitivityInspectionTemplateIncludes", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"batch-get-custom-data-identifiers": {
			Name:   "batch-get-custom-data-identifiers",
			Fields: fields_batch_get_custom_data_identifiers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCustomDataIdentifiersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_custom_data_identifiers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCustomDataIdentifiers(ctx, input)
			},
		},
		"batch-update-automated-discovery-accounts": {
			Name:   "batch-update-automated-discovery-accounts",
			Fields: fields_batch_update_automated_discovery_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateAutomatedDiscoveryAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_automated_discovery_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateAutomatedDiscoveryAccounts(ctx, input)
			},
		},
		"create-allow-list": {
			Name:   "create-allow-list",
			Fields: fields_create_allow_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAllowListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_allow_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAllowList(ctx, input)
			},
		},
		"create-classification-job": {
			Name:   "create-classification-job",
			Fields: fields_create_classification_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClassificationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_classification_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClassificationJob(ctx, input)
			},
		},
		"create-custom-data-identifier": {
			Name:   "create-custom-data-identifier",
			Fields: fields_create_custom_data_identifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomDataIdentifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_data_identifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomDataIdentifier(ctx, input)
			},
		},
		"create-findings-filter": {
			Name:   "create-findings-filter",
			Fields: fields_create_findings_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFindingsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_findings_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFindingsFilter(ctx, input)
			},
		},
		"create-invitations": {
			Name:   "create-invitations",
			Fields: fields_create_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInvitationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_invitations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInvitations(ctx, input)
			},
		},
		"create-member": {
			Name:   "create-member",
			Fields: fields_create_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMember(ctx, input)
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
		"delete-allow-list": {
			Name:   "delete-allow-list",
			Fields: fields_delete_allow_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAllowListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_allow_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAllowList(ctx, input)
			},
		},
		"delete-custom-data-identifier": {
			Name:   "delete-custom-data-identifier",
			Fields: fields_delete_custom_data_identifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomDataIdentifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_data_identifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomDataIdentifier(ctx, input)
			},
		},
		"delete-findings-filter": {
			Name:   "delete-findings-filter",
			Fields: fields_delete_findings_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFindingsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_findings_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFindingsFilter(ctx, input)
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
		"describe-buckets": {
			Name:   "describe-buckets",
			Fields: fields_describe_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBuckets(ctx, input)
				}
				var results []*svc.DescribeBucketsOutput
				p := svc.NewDescribeBucketsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-classification-job": {
			Name:   "describe-classification-job",
			Fields: fields_describe_classification_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClassificationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_classification_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClassificationJob(ctx, input)
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
		"disable-macie": {
			Name:   "disable-macie",
			Fields: fields_disable_macie,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableMacieInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_macie, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableMacie(ctx, input)
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
		"enable-macie": {
			Name:   "enable-macie",
			Fields: fields_enable_macie,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableMacieInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_macie, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableMacie(ctx, input)
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
		"get-allow-list": {
			Name:   "get-allow-list",
			Fields: fields_get_allow_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAllowListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_allow_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAllowList(ctx, input)
			},
		},
		"get-automated-discovery-configuration": {
			Name:   "get-automated-discovery-configuration",
			Fields: fields_get_automated_discovery_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomatedDiscoveryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automated_discovery_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomatedDiscoveryConfiguration(ctx, input)
			},
		},
		"get-bucket-statistics": {
			Name:   "get-bucket-statistics",
			Fields: fields_get_bucket_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketStatistics(ctx, input)
			},
		},
		"get-classification-export-configuration": {
			Name:   "get-classification-export-configuration",
			Fields: fields_get_classification_export_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClassificationExportConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_classification_export_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClassificationExportConfiguration(ctx, input)
			},
		},
		"get-classification-scope": {
			Name:   "get-classification-scope",
			Fields: fields_get_classification_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClassificationScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_classification_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClassificationScope(ctx, input)
			},
		},
		"get-custom-data-identifier": {
			Name:   "get-custom-data-identifier",
			Fields: fields_get_custom_data_identifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomDataIdentifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_data_identifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomDataIdentifier(ctx, input)
			},
		},
		"get-finding-statistics": {
			Name:   "get-finding-statistics",
			Fields: fields_get_finding_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_finding_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingStatistics(ctx, input)
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
		"get-findings-filter": {
			Name:   "get-findings-filter",
			Fields: fields_get_findings_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_findings_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingsFilter(ctx, input)
			},
		},
		"get-findings-publication-configuration": {
			Name:   "get-findings-publication-configuration",
			Fields: fields_get_findings_publication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsPublicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_findings_publication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingsPublicationConfiguration(ctx, input)
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
		"get-macie-session": {
			Name:   "get-macie-session",
			Fields: fields_get_macie_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMacieSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_macie_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMacieSession(ctx, input)
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
		"get-resource-profile": {
			Name:   "get-resource-profile",
			Fields: fields_get_resource_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceProfile(ctx, input)
			},
		},
		"get-reveal-configuration": {
			Name:   "get-reveal-configuration",
			Fields: fields_get_reveal_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRevealConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reveal_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRevealConfiguration(ctx, input)
			},
		},
		"get-sensitive-data-occurrences": {
			Name:   "get-sensitive-data-occurrences",
			Fields: fields_get_sensitive_data_occurrences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSensitiveDataOccurrencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sensitive_data_occurrences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSensitiveDataOccurrences(ctx, input)
			},
		},
		"get-sensitive-data-occurrences-availability": {
			Name:   "get-sensitive-data-occurrences-availability",
			Fields: fields_get_sensitive_data_occurrences_availability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSensitiveDataOccurrencesAvailabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sensitive_data_occurrences_availability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSensitiveDataOccurrencesAvailability(ctx, input)
			},
		},
		"get-sensitivity-inspection-template": {
			Name:   "get-sensitivity-inspection-template",
			Fields: fields_get_sensitivity_inspection_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSensitivityInspectionTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sensitivity_inspection_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSensitivityInspectionTemplate(ctx, input)
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
		"get-usage-totals": {
			Name:   "get-usage-totals",
			Fields: fields_get_usage_totals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsageTotalsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_usage_totals, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsageTotals(ctx, input)
			},
		},
		"list-allow-lists": {
			Name:   "list-allow-lists",
			Fields: fields_list_allow_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAllowListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_allow_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAllowLists(ctx, input)
				}
				var results []*svc.ListAllowListsOutput
				p := svc.NewListAllowListsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-automated-discovery-accounts": {
			Name:   "list-automated-discovery-accounts",
			Fields: fields_list_automated_discovery_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomatedDiscoveryAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automated_discovery_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomatedDiscoveryAccounts(ctx, input)
				}
				var results []*svc.ListAutomatedDiscoveryAccountsOutput
				p := svc.NewListAutomatedDiscoveryAccountsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-classification-jobs": {
			Name:   "list-classification-jobs",
			Fields: fields_list_classification_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClassificationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_classification_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClassificationJobs(ctx, input)
				}
				var results []*svc.ListClassificationJobsOutput
				p := svc.NewListClassificationJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-classification-scopes": {
			Name:   "list-classification-scopes",
			Fields: fields_list_classification_scopes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClassificationScopesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_classification_scopes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClassificationScopes(ctx, input)
				}
				var results []*svc.ListClassificationScopesOutput
				p := svc.NewListClassificationScopesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-custom-data-identifiers": {
			Name:   "list-custom-data-identifiers",
			Fields: fields_list_custom_data_identifiers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomDataIdentifiersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_data_identifiers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomDataIdentifiers(ctx, input)
				}
				var results []*svc.ListCustomDataIdentifiersOutput
				p := svc.NewListCustomDataIdentifiersPaginator(client, input)
				for p.HasMorePages() {
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
		"list-findings-filters": {
			Name:   "list-findings-filters",
			Fields: fields_list_findings_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingsFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_findings_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindingsFilters(ctx, input)
				}
				var results []*svc.ListFindingsFiltersOutput
				p := svc.NewListFindingsFiltersPaginator(client, input)
				for p.HasMorePages() {
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
		"list-managed-data-identifiers": {
			Name:   "list-managed-data-identifiers",
			Fields: fields_list_managed_data_identifiers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedDataIdentifiersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_data_identifiers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedDataIdentifiers(ctx, input)
				}
				var results []*svc.ListManagedDataIdentifiersOutput
				p := svc.NewListManagedDataIdentifiersPaginator(client, input)
				for p.HasMorePages() {
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
		"list-resource-profile-artifacts": {
			Name:   "list-resource-profile-artifacts",
			Fields: fields_list_resource_profile_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceProfileArtifactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_profile_artifacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceProfileArtifacts(ctx, input)
				}
				var results []*svc.ListResourceProfileArtifactsOutput
				p := svc.NewListResourceProfileArtifactsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-resource-profile-detections": {
			Name:   "list-resource-profile-detections",
			Fields: fields_list_resource_profile_detections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceProfileDetectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_profile_detections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceProfileDetections(ctx, input)
				}
				var results []*svc.ListResourceProfileDetectionsOutput
				p := svc.NewListResourceProfileDetectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-sensitivity-inspection-templates": {
			Name:   "list-sensitivity-inspection-templates",
			Fields: fields_list_sensitivity_inspection_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSensitivityInspectionTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sensitivity_inspection_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSensitivityInspectionTemplates(ctx, input)
				}
				var results []*svc.ListSensitivityInspectionTemplatesOutput
				p := svc.NewListSensitivityInspectionTemplatesPaginator(client, input)
				for p.HasMorePages() {
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
		"put-classification-export-configuration": {
			Name:   "put-classification-export-configuration",
			Fields: fields_put_classification_export_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutClassificationExportConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_classification_export_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutClassificationExportConfiguration(ctx, input)
			},
		},
		"put-findings-publication-configuration": {
			Name:   "put-findings-publication-configuration",
			Fields: fields_put_findings_publication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFindingsPublicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_findings_publication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFindingsPublicationConfiguration(ctx, input)
			},
		},
		"search-resources": {
			Name:   "search-resources",
			Fields: fields_search_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchResources(ctx, input)
				}
				var results []*svc.SearchResourcesOutput
				p := svc.NewSearchResourcesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
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
		"test-custom-data-identifier": {
			Name:   "test-custom-data-identifier",
			Fields: fields_test_custom_data_identifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestCustomDataIdentifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_custom_data_identifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestCustomDataIdentifier(ctx, input)
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
		"update-allow-list": {
			Name:   "update-allow-list",
			Fields: fields_update_allow_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAllowListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_allow_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAllowList(ctx, input)
			},
		},
		"update-automated-discovery-configuration": {
			Name:   "update-automated-discovery-configuration",
			Fields: fields_update_automated_discovery_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomatedDiscoveryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automated_discovery_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomatedDiscoveryConfiguration(ctx, input)
			},
		},
		"update-classification-job": {
			Name:   "update-classification-job",
			Fields: fields_update_classification_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClassificationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_classification_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClassificationJob(ctx, input)
			},
		},
		"update-classification-scope": {
			Name:   "update-classification-scope",
			Fields: fields_update_classification_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClassificationScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_classification_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClassificationScope(ctx, input)
			},
		},
		"update-findings-filter": {
			Name:   "update-findings-filter",
			Fields: fields_update_findings_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFindingsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_findings_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFindingsFilter(ctx, input)
			},
		},
		"update-macie-session": {
			Name:   "update-macie-session",
			Fields: fields_update_macie_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMacieSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_macie_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMacieSession(ctx, input)
			},
		},
		"update-member-session": {
			Name:   "update-member-session",
			Fields: fields_update_member_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMemberSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_member_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMemberSession(ctx, input)
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
		"update-resource-profile": {
			Name:   "update-resource-profile",
			Fields: fields_update_resource_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceProfile(ctx, input)
			},
		},
		"update-resource-profile-detections": {
			Name:   "update-resource-profile-detections",
			Fields: fields_update_resource_profile_detections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceProfileDetectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_profile_detections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceProfileDetections(ctx, input)
			},
		},
		"update-reveal-configuration": {
			Name:   "update-reveal-configuration",
			Fields: fields_update_reveal_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRevealConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_reveal_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRevealConfiguration(ctx, input)
			},
		},
		"update-sensitivity-inspection-template": {
			Name:   "update-sensitivity-inspection-template",
			Fields: fields_update_sensitivity_inspection_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSensitivityInspectionTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sensitivity_inspection_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSensitivityInspectionTemplate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("macie2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
