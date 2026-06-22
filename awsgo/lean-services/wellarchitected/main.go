package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/wellarchitected"
)

var fields_associate_lenses = []leanruntime.Field{
	{Name: "LensAliases", Flag: "lens-aliases", Type: "[]string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_associate_profiles = []leanruntime.Field{
	{Name: "ProfileArns", Flag: "profile-arns", Type: "[]string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_create_lens_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "SharedWith", Flag: "shared-with", Type: "*string", Required: true},
}

var fields_create_lens_version = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "IsMajorVersion", Flag: "is-major-version", Type: "*bool", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "LensVersion", Flag: "lens-version", Type: "*string", Required: true},
}

var fields_create_milestone = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "MilestoneName", Flag: "milestone-name", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_create_profile = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ProfileDescription", Flag: "profile-description", Type: "*string", Required: true},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "ProfileQuestions", Flag: "profile-questions", Type: "[]types.ProfileQuestionUpdate", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_profile_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
	{Name: "SharedWith", Flag: "shared-with", Type: "*string", Required: true},
}

var fields_create_review_template = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Lenses", Flag: "lenses", Type: "[]string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_template_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "SharedWith", Flag: "shared-with", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_create_workload = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Applications", Flag: "applications", Type: "[]string", Required: false},
	{Name: "ArchitecturalDesign", Flag: "architectural-design", Type: "*string", Required: false},
	{Name: "AwsRegions", Flag: "aws-regions", Type: "[]string", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DiscoveryConfig", Flag: "discovery-config", Type: "*types.WorkloadDiscoveryConfig", Required: false},
	{Name: "Environment", Flag: "environment", Type: "types.WorkloadEnvironment", Required: true},
	{Name: "Industry", Flag: "industry", Type: "*string", Required: false},
	{Name: "IndustryType", Flag: "industry-type", Type: "*string", Required: false},
	{Name: "JiraConfiguration", Flag: "jira-configuration", Type: "*types.WorkloadJiraConfigurationInput", Required: false},
	{Name: "Lenses", Flag: "lenses", Type: "[]string", Required: true},
	{Name: "NonAwsRegions", Flag: "non-aws-regions", Type: "[]string", Required: false},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "PillarPriorities", Flag: "pillar-priorities", Type: "[]string", Required: false},
	{Name: "ProfileArns", Flag: "profile-arns", Type: "[]string", Required: false},
	{Name: "ReviewOwner", Flag: "review-owner", Type: "*string", Required: false},
	{Name: "ReviewTemplateArns", Flag: "review-template-arns", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_create_workload_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "PermissionType", Flag: "permission-type", Type: "types.PermissionType", Required: true},
	{Name: "SharedWith", Flag: "shared-with", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_delete_lens = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "LensStatus", Flag: "lens-status", Type: "types.LensStatusType", Required: true},
}

var fields_delete_lens_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
}

var fields_delete_profile = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
}

var fields_delete_profile_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
}

var fields_delete_review_template = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_delete_template_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_delete_workload = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_delete_workload_share = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_disassociate_lenses = []leanruntime.Field{
	{Name: "LensAliases", Flag: "lens-aliases", Type: "[]string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_disassociate_profiles = []leanruntime.Field{
	{Name: "ProfileArns", Flag: "profile-arns", Type: "[]string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_export_lens = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "LensVersion", Flag: "lens-version", Type: "*string", Required: false},
}

var fields_get_answer = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: false},
	{Name: "QuestionId", Flag: "question-id", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_get_consolidated_report = []leanruntime.Field{
	{Name: "Format", Flag: "format", Type: "types.ReportFormat", Required: true},
	{Name: "IncludeSharedResources", Flag: "include-shared-resources", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_global_settings = []leanruntime.Field{}

var fields_get_lens = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "LensVersion", Flag: "lens-version", Type: "*string", Required: false},
}

var fields_get_lens_review = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_get_lens_review_report = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_get_lens_version_difference = []leanruntime.Field{
	{Name: "BaseLensVersion", Flag: "base-lens-version", Type: "*string", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "TargetLensVersion", Flag: "target-lens-version", Type: "*string", Required: false},
}

var fields_get_milestone = []leanruntime.Field{
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_get_profile = []leanruntime.Field{
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
	{Name: "ProfileVersion", Flag: "profile-version", Type: "*string", Required: false},
}

var fields_get_profile_template = []leanruntime.Field{}

var fields_get_review_template = []leanruntime.Field{
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_get_review_template_answer = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "QuestionId", Flag: "question-id", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_get_review_template_lens_review = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_get_workload = []leanruntime.Field{
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_import_lens = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "JSONString", Flag: "json-string", Type: "*string", Required: true},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_answers = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PillarId", Flag: "pillar-id", Type: "*string", Required: false},
	{Name: "QuestionPriority", Flag: "question-priority", Type: "types.QuestionPriority", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_check_details = []leanruntime.Field{
	{Name: "ChoiceId", Flag: "choice-id", Type: "*string", Required: true},
	{Name: "LensArn", Flag: "lens-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PillarId", Flag: "pillar-id", Type: "*string", Required: true},
	{Name: "QuestionId", Flag: "question-id", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_check_summaries = []leanruntime.Field{
	{Name: "ChoiceId", Flag: "choice-id", Type: "*string", Required: true},
	{Name: "LensArn", Flag: "lens-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PillarId", Flag: "pillar-id", Type: "*string", Required: true},
	{Name: "QuestionId", Flag: "question-id", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_lens_review_improvements = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PillarId", Flag: "pillar-id", Type: "*string", Required: false},
	{Name: "QuestionPriority", Flag: "question-priority", Type: "types.QuestionPriority", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_lens_reviews = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MilestoneNumber", Flag: "milestone-number", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_lens_shares = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SharedWithPrefix", Flag: "shared-with-prefix", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ShareStatus", Required: false},
}

var fields_list_lenses = []leanruntime.Field{
	{Name: "LensName", Flag: "lens-name", Type: "*string", Required: false},
	{Name: "LensStatus", Flag: "lens-status", Type: "types.LensStatusType", Required: false},
	{Name: "LensType", Flag: "lens-type", Type: "types.LensType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_milestones = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_notifications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: false},
}

var fields_list_profile_notifications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: false},
}

var fields_list_profile_shares = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
	{Name: "SharedWithPrefix", Flag: "shared-with-prefix", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ShareStatus", Required: false},
}

var fields_list_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileNamePrefix", Flag: "profile-name-prefix", Type: "*string", Required: false},
	{Name: "ProfileOwnerType", Flag: "profile-owner-type", Type: "types.ProfileOwnerType", Required: false},
}

var fields_list_review_template_answers = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PillarId", Flag: "pillar-id", Type: "*string", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_list_review_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_share_invitations = []leanruntime.Field{
	{Name: "LensNamePrefix", Flag: "lens-name-prefix", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileNamePrefix", Flag: "profile-name-prefix", Type: "*string", Required: false},
	{Name: "ShareResourceType", Flag: "share-resource-type", Type: "types.ShareResourceType", Required: false},
	{Name: "TemplateNamePrefix", Flag: "template-name-prefix", Type: "*string", Required: false},
	{Name: "WorkloadNamePrefix", Flag: "workload-name-prefix", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "WorkloadArn", Flag: "workload-arn", Type: "*string", Required: true},
}

var fields_list_template_shares = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SharedWithPrefix", Flag: "shared-with-prefix", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ShareStatus", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_list_workload_shares = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SharedWithPrefix", Flag: "shared-with-prefix", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ShareStatus", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_list_workloads = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadNamePrefix", Flag: "workload-name-prefix", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
	{Name: "WorkloadArn", Flag: "workload-arn", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
	{Name: "WorkloadArn", Flag: "workload-arn", Type: "*string", Required: true},
}

var fields_update_answer = []leanruntime.Field{
	{Name: "ChoiceUpdates", Flag: "choice-updates", Type: "map[string]types.ChoiceUpdate", Required: false},
	{Name: "IsApplicable", Flag: "is-applicable", Type: "*bool", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "QuestionId", Flag: "question-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "types.AnswerReason", Required: false},
	{Name: "SelectedChoices", Flag: "selected-choices", Type: "[]string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_update_global_settings = []leanruntime.Field{
	{Name: "DiscoveryIntegrationStatus", Flag: "discovery-integration-status", Type: "types.DiscoveryIntegrationStatus", Required: false},
	{Name: "JiraConfiguration", Flag: "jira-configuration", Type: "*types.AccountJiraConfigurationInput", Required: false},
	{Name: "OrganizationSharingStatus", Flag: "organization-sharing-status", Type: "types.OrganizationSharingStatus", Required: false},
}

var fields_update_integration = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "IntegratingService", Flag: "integrating-service", Type: "types.IntegratingService", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_update_lens_review = []leanruntime.Field{
	{Name: "JiraConfiguration", Flag: "jira-configuration", Type: "*types.JiraSelectedQuestionConfiguration", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "LensNotes", Flag: "lens-notes", Type: "*string", Required: false},
	{Name: "PillarNotes", Flag: "pillar-notes", Type: "map[string]string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_update_profile = []leanruntime.Field{
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
	{Name: "ProfileDescription", Flag: "profile-description", Type: "*string", Required: false},
	{Name: "ProfileQuestions", Flag: "profile-questions", Type: "[]types.ProfileQuestionUpdate", Required: false},
}

var fields_update_review_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LensesToAssociate", Flag: "lenses-to-associate", Type: "[]string", Required: false},
	{Name: "LensesToDisassociate", Flag: "lenses-to-disassociate", Type: "[]string", Required: false},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

var fields_update_review_template_answer = []leanruntime.Field{
	{Name: "ChoiceUpdates", Flag: "choice-updates", Type: "map[string]types.ChoiceUpdate", Required: false},
	{Name: "IsApplicable", Flag: "is-applicable", Type: "*bool", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "QuestionId", Flag: "question-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "types.AnswerReason", Required: false},
	{Name: "SelectedChoices", Flag: "selected-choices", Type: "[]string", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_update_review_template_lens_review = []leanruntime.Field{
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "LensNotes", Flag: "lens-notes", Type: "*string", Required: false},
	{Name: "PillarNotes", Flag: "pillar-notes", Type: "map[string]string", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_update_share_invitation = []leanruntime.Field{
	{Name: "ShareInvitationAction", Flag: "share-invitation-action", Type: "types.ShareInvitationAction", Required: true},
	{Name: "ShareInvitationId", Flag: "share-invitation-id", Type: "*string", Required: true},
}

var fields_update_workload = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Applications", Flag: "applications", Type: "[]string", Required: false},
	{Name: "ArchitecturalDesign", Flag: "architectural-design", Type: "*string", Required: false},
	{Name: "AwsRegions", Flag: "aws-regions", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiscoveryConfig", Flag: "discovery-config", Type: "*types.WorkloadDiscoveryConfig", Required: false},
	{Name: "Environment", Flag: "environment", Type: "types.WorkloadEnvironment", Required: false},
	{Name: "ImprovementStatus", Flag: "improvement-status", Type: "types.WorkloadImprovementStatus", Required: false},
	{Name: "Industry", Flag: "industry", Type: "*string", Required: false},
	{Name: "IndustryType", Flag: "industry-type", Type: "*string", Required: false},
	{Name: "IsReviewOwnerUpdateAcknowledged", Flag: "is-review-owner-update-acknowledged", Type: "*bool", Required: false},
	{Name: "JiraConfiguration", Flag: "jira-configuration", Type: "*types.WorkloadJiraConfigurationInput", Required: false},
	{Name: "NonAwsRegions", Flag: "non-aws-regions", Type: "[]string", Required: false},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "PillarPriorities", Flag: "pillar-priorities", Type: "[]string", Required: false},
	{Name: "ReviewOwner", Flag: "review-owner", Type: "*string", Required: false},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: false},
}

var fields_update_workload_share = []leanruntime.Field{
	{Name: "PermissionType", Flag: "permission-type", Type: "types.PermissionType", Required: true},
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_upgrade_lens_review = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "MilestoneName", Flag: "milestone-name", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_upgrade_profile_version = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "MilestoneName", Flag: "milestone-name", Type: "*string", Required: false},
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: true},
	{Name: "WorkloadId", Flag: "workload-id", Type: "*string", Required: true},
}

var fields_upgrade_review_template_lens_review = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LensAlias", Flag: "lens-alias", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-lenses": {
			Name:   "associate-lenses",
			Fields: fields_associate_lenses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateLensesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_lenses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateLenses(ctx, input)
			},
		},
		"associate-profiles": {
			Name:   "associate-profiles",
			Fields: fields_associate_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateProfiles(ctx, input)
			},
		},
		"create-lens-share": {
			Name:   "create-lens-share",
			Fields: fields_create_lens_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLensShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lens_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLensShare(ctx, input)
			},
		},
		"create-lens-version": {
			Name:   "create-lens-version",
			Fields: fields_create_lens_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLensVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lens_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLensVersion(ctx, input)
			},
		},
		"create-milestone": {
			Name:   "create-milestone",
			Fields: fields_create_milestone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMilestoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_milestone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMilestone(ctx, input)
			},
		},
		"create-profile": {
			Name:   "create-profile",
			Fields: fields_create_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfile(ctx, input)
			},
		},
		"create-profile-share": {
			Name:   "create-profile-share",
			Fields: fields_create_profile_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfileShare(ctx, input)
			},
		},
		"create-review-template": {
			Name:   "create-review-template",
			Fields: fields_create_review_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReviewTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_review_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReviewTemplate(ctx, input)
			},
		},
		"create-template-share": {
			Name:   "create-template-share",
			Fields: fields_create_template_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplateShare(ctx, input)
			},
		},
		"create-workload": {
			Name:   "create-workload",
			Fields: fields_create_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkload(ctx, input)
			},
		},
		"create-workload-share": {
			Name:   "create-workload-share",
			Fields: fields_create_workload_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkloadShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workload_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkloadShare(ctx, input)
			},
		},
		"delete-lens": {
			Name:   "delete-lens",
			Fields: fields_delete_lens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLens(ctx, input)
			},
		},
		"delete-lens-share": {
			Name:   "delete-lens-share",
			Fields: fields_delete_lens_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLensShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lens_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLensShare(ctx, input)
			},
		},
		"delete-profile": {
			Name:   "delete-profile",
			Fields: fields_delete_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfile(ctx, input)
			},
		},
		"delete-profile-share": {
			Name:   "delete-profile-share",
			Fields: fields_delete_profile_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfileShare(ctx, input)
			},
		},
		"delete-review-template": {
			Name:   "delete-review-template",
			Fields: fields_delete_review_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReviewTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_review_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReviewTemplate(ctx, input)
			},
		},
		"delete-template-share": {
			Name:   "delete-template-share",
			Fields: fields_delete_template_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplateShare(ctx, input)
			},
		},
		"delete-workload": {
			Name:   "delete-workload",
			Fields: fields_delete_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkload(ctx, input)
			},
		},
		"delete-workload-share": {
			Name:   "delete-workload-share",
			Fields: fields_delete_workload_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkloadShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workload_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkloadShare(ctx, input)
			},
		},
		"disassociate-lenses": {
			Name:   "disassociate-lenses",
			Fields: fields_disassociate_lenses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateLensesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_lenses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateLenses(ctx, input)
			},
		},
		"disassociate-profiles": {
			Name:   "disassociate-profiles",
			Fields: fields_disassociate_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateProfiles(ctx, input)
			},
		},
		"export-lens": {
			Name:   "export-lens",
			Fields: fields_export_lens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportLensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_lens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportLens(ctx, input)
			},
		},
		"get-answer": {
			Name:   "get-answer",
			Fields: fields_get_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnswer(ctx, input)
			},
		},
		"get-consolidated-report": {
			Name:   "get-consolidated-report",
			Fields: fields_get_consolidated_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConsolidatedReportInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_consolidated_report, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetConsolidatedReport(ctx, input)
				}
				var results []*svc.GetConsolidatedReportOutput
				p := svc.NewGetConsolidatedReportPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-global-settings": {
			Name:   "get-global-settings",
			Fields: fields_get_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlobalSettings(ctx, input)
			},
		},
		"get-lens": {
			Name:   "get-lens",
			Fields: fields_get_lens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLens(ctx, input)
			},
		},
		"get-lens-review": {
			Name:   "get-lens-review",
			Fields: fields_get_lens_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLensReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lens_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLensReview(ctx, input)
			},
		},
		"get-lens-review-report": {
			Name:   "get-lens-review-report",
			Fields: fields_get_lens_review_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLensReviewReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lens_review_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLensReviewReport(ctx, input)
			},
		},
		"get-lens-version-difference": {
			Name:   "get-lens-version-difference",
			Fields: fields_get_lens_version_difference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLensVersionDifferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lens_version_difference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLensVersionDifference(ctx, input)
			},
		},
		"get-milestone": {
			Name:   "get-milestone",
			Fields: fields_get_milestone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMilestoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_milestone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMilestone(ctx, input)
			},
		},
		"get-profile": {
			Name:   "get-profile",
			Fields: fields_get_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfile(ctx, input)
			},
		},
		"get-profile-template": {
			Name:   "get-profile-template",
			Fields: fields_get_profile_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileTemplate(ctx, input)
			},
		},
		"get-review-template": {
			Name:   "get-review-template",
			Fields: fields_get_review_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReviewTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_review_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReviewTemplate(ctx, input)
			},
		},
		"get-review-template-answer": {
			Name:   "get-review-template-answer",
			Fields: fields_get_review_template_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReviewTemplateAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_review_template_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReviewTemplateAnswer(ctx, input)
			},
		},
		"get-review-template-lens-review": {
			Name:   "get-review-template-lens-review",
			Fields: fields_get_review_template_lens_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReviewTemplateLensReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_review_template_lens_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReviewTemplateLensReview(ctx, input)
			},
		},
		"get-workload": {
			Name:   "get-workload",
			Fields: fields_get_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkload(ctx, input)
			},
		},
		"import-lens": {
			Name:   "import-lens",
			Fields: fields_import_lens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportLensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_lens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportLens(ctx, input)
			},
		},
		"list-answers": {
			Name:   "list-answers",
			Fields: fields_list_answers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnswersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_answers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnswers(ctx, input)
				}
				var results []*svc.ListAnswersOutput
				p := svc.NewListAnswersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-check-details": {
			Name:   "list-check-details",
			Fields: fields_list_check_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCheckDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_check_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCheckDetails(ctx, input)
				}
				var results []*svc.ListCheckDetailsOutput
				p := svc.NewListCheckDetailsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-check-summaries": {
			Name:   "list-check-summaries",
			Fields: fields_list_check_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCheckSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_check_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCheckSummaries(ctx, input)
				}
				var results []*svc.ListCheckSummariesOutput
				p := svc.NewListCheckSummariesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-lens-review-improvements": {
			Name:   "list-lens-review-improvements",
			Fields: fields_list_lens_review_improvements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLensReviewImprovementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lens_review_improvements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLensReviewImprovements(ctx, input)
				}
				var results []*svc.ListLensReviewImprovementsOutput
				p := svc.NewListLensReviewImprovementsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-lens-reviews": {
			Name:   "list-lens-reviews",
			Fields: fields_list_lens_reviews,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLensReviewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lens_reviews, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLensReviews(ctx, input)
				}
				var results []*svc.ListLensReviewsOutput
				p := svc.NewListLensReviewsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-lens-shares": {
			Name:   "list-lens-shares",
			Fields: fields_list_lens_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLensSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lens_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLensShares(ctx, input)
				}
				var results []*svc.ListLensSharesOutput
				p := svc.NewListLensSharesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-lenses": {
			Name:   "list-lenses",
			Fields: fields_list_lenses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLensesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lenses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLenses(ctx, input)
				}
				var results []*svc.ListLensesOutput
				p := svc.NewListLensesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-milestones": {
			Name:   "list-milestones",
			Fields: fields_list_milestones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMilestonesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_milestones, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMilestones(ctx, input)
				}
				var results []*svc.ListMilestonesOutput
				p := svc.NewListMilestonesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-notifications": {
			Name:   "list-notifications",
			Fields: fields_list_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotifications(ctx, input)
				}
				var results []*svc.ListNotificationsOutput
				p := svc.NewListNotificationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-profile-notifications": {
			Name:   "list-profile-notifications",
			Fields: fields_list_profile_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileNotificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profile_notifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfileNotifications(ctx, input)
				}
				var results []*svc.ListProfileNotificationsOutput
				p := svc.NewListProfileNotificationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-profile-shares": {
			Name:   "list-profile-shares",
			Fields: fields_list_profile_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profile_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfileShares(ctx, input)
				}
				var results []*svc.ListProfileSharesOutput
				p := svc.NewListProfileSharesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-profiles": {
			Name:   "list-profiles",
			Fields: fields_list_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfiles(ctx, input)
				}
				var results []*svc.ListProfilesOutput
				p := svc.NewListProfilesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-review-template-answers": {
			Name:   "list-review-template-answers",
			Fields: fields_list_review_template_answers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReviewTemplateAnswersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_review_template_answers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReviewTemplateAnswers(ctx, input)
				}
				var results []*svc.ListReviewTemplateAnswersOutput
				p := svc.NewListReviewTemplateAnswersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-review-templates": {
			Name:   "list-review-templates",
			Fields: fields_list_review_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReviewTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_review_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReviewTemplates(ctx, input)
				}
				var results []*svc.ListReviewTemplatesOutput
				p := svc.NewListReviewTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-share-invitations": {
			Name:   "list-share-invitations",
			Fields: fields_list_share_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListShareInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_share_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListShareInvitations(ctx, input)
				}
				var results []*svc.ListShareInvitationsOutput
				p := svc.NewListShareInvitationsPaginator(client, input)
				for p.HasMorePages() {
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
		"list-template-shares": {
			Name:   "list-template-shares",
			Fields: fields_list_template_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateShares(ctx, input)
				}
				var results []*svc.ListTemplateSharesOutput
				p := svc.NewListTemplateSharesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-workload-shares": {
			Name:   "list-workload-shares",
			Fields: fields_list_workload_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workload_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloadShares(ctx, input)
				}
				var results []*svc.ListWorkloadSharesOutput
				p := svc.NewListWorkloadSharesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-workloads": {
			Name:   "list-workloads",
			Fields: fields_list_workloads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workloads, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloads(ctx, input)
				}
				var results []*svc.ListWorkloadsOutput
				p := svc.NewListWorkloadsPaginator(client, input)
				for p.HasMorePages() {
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
		"update-answer": {
			Name:   "update-answer",
			Fields: fields_update_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnswer(ctx, input)
			},
		},
		"update-global-settings": {
			Name:   "update-global-settings",
			Fields: fields_update_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalSettings(ctx, input)
			},
		},
		"update-integration": {
			Name:   "update-integration",
			Fields: fields_update_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIntegration(ctx, input)
			},
		},
		"update-lens-review": {
			Name:   "update-lens-review",
			Fields: fields_update_lens_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLensReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lens_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLensReview(ctx, input)
			},
		},
		"update-profile": {
			Name:   "update-profile",
			Fields: fields_update_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfile(ctx, input)
			},
		},
		"update-review-template": {
			Name:   "update-review-template",
			Fields: fields_update_review_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReviewTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_review_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReviewTemplate(ctx, input)
			},
		},
		"update-review-template-answer": {
			Name:   "update-review-template-answer",
			Fields: fields_update_review_template_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReviewTemplateAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_review_template_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReviewTemplateAnswer(ctx, input)
			},
		},
		"update-review-template-lens-review": {
			Name:   "update-review-template-lens-review",
			Fields: fields_update_review_template_lens_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReviewTemplateLensReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_review_template_lens_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReviewTemplateLensReview(ctx, input)
			},
		},
		"update-share-invitation": {
			Name:   "update-share-invitation",
			Fields: fields_update_share_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateShareInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_share_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateShareInvitation(ctx, input)
			},
		},
		"update-workload": {
			Name:   "update-workload",
			Fields: fields_update_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkload(ctx, input)
			},
		},
		"update-workload-share": {
			Name:   "update-workload-share",
			Fields: fields_update_workload_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkloadShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workload_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkloadShare(ctx, input)
			},
		},
		"upgrade-lens-review": {
			Name:   "upgrade-lens-review",
			Fields: fields_upgrade_lens_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeLensReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_lens_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeLensReview(ctx, input)
			},
		},
		"upgrade-profile-version": {
			Name:   "upgrade-profile-version",
			Fields: fields_upgrade_profile_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeProfileVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_profile_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeProfileVersion(ctx, input)
			},
		},
		"upgrade-review-template-lens-review": {
			Name:   "upgrade-review-template-lens-review",
			Fields: fields_upgrade_review_template_lens_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeReviewTemplateLensReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_review_template_lens_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeReviewTemplateLensReview(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("wellarchitected", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
