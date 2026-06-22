package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/partnercentralselling"
)

var fields_accept_engagement_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_assign_opportunity = []leanruntime.Field{
	{Name: "Assignee", Flag: "assignee", Type: "*types.AssigneeContact", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_associate_opportunity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "OpportunityIdentifier", Flag: "opportunity-identifier", Type: "*string", Required: true},
	{Name: "RelatedEntityIdentifier", Flag: "related-entity-identifier", Type: "*string", Required: true},
	{Name: "RelatedEntityType", Flag: "related-entity-type", Type: "types.RelatedEntityType", Required: true},
}

var fields_create_engagement = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Contexts", Flag: "contexts", Type: "[]types.EngagementContextDetails", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_create_engagement_context = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "Payload", Flag: "payload", Type: "types.EngagementContextPayload", Required: true},
	{Name: "Type", Flag: "type", Type: "types.EngagementContextType", Required: true},
}

var fields_create_engagement_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "Invitation", Flag: "invitation", Type: "*types.Invitation", Required: true},
}

var fields_create_opportunity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Customer", Flag: "customer", Type: "*types.Customer", Required: false},
	{Name: "LifeCycle", Flag: "life-cycle", Type: "*types.LifeCycle", Required: false},
	{Name: "Marketing", Flag: "marketing", Type: "*types.Marketing", Required: false},
	{Name: "NationalSecurity", Flag: "national-security", Type: "types.NationalSecurity", Required: false},
	{Name: "OpportunityTeam", Flag: "opportunity-team", Type: "[]types.Contact", Required: false},
	{Name: "OpportunityType", Flag: "opportunity-type", Type: "types.OpportunityType", Required: false},
	{Name: "Origin", Flag: "origin", Type: "types.OpportunityOrigin", Required: false},
	{Name: "PartnerOpportunityIdentifier", Flag: "partner-opportunity-identifier", Type: "*string", Required: false},
	{Name: "PrimaryNeedsFromAws", Flag: "primary-needs-from-aws", Type: "[]types.PrimaryNeedFromAws", Required: false},
	{Name: "Project", Flag: "project", Type: "*types.Project", Required: false},
	{Name: "SoftwareRevenue", Flag: "software-revenue", Type: "*types.SoftwareRevenue", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_resource_snapshot = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceSnapshotTemplateIdentifier", Flag: "resource-snapshot-template-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_create_resource_snapshot_job = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceSnapshotTemplateIdentifier", Flag: "resource-snapshot-template-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_resource_snapshot_job = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceSnapshotJobIdentifier", Flag: "resource-snapshot-job-identifier", Type: "*string", Required: true},
}

var fields_disassociate_opportunity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "OpportunityIdentifier", Flag: "opportunity-identifier", Type: "*string", Required: true},
	{Name: "RelatedEntityIdentifier", Flag: "related-entity-identifier", Type: "*string", Required: true},
	{Name: "RelatedEntityType", Flag: "related-entity-type", Type: "types.RelatedEntityType", Required: true},
}

var fields_get_aws_opportunity_summary = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "RelatedOpportunityIdentifier", Flag: "related-opportunity-identifier", Type: "*string", Required: true},
}

var fields_get_engagement = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_engagement_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_opportunity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_resource_snapshot = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceSnapshotTemplateIdentifier", Flag: "resource-snapshot-template-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*int32", Required: false},
}

var fields_get_resource_snapshot_job = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceSnapshotJobIdentifier", Flag: "resource-snapshot-job-identifier", Type: "*string", Required: true},
}

var fields_get_selling_system_settings = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
}

var fields_list_engagement_by_accepting_invitation_tasks = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EngagementInvitationIdentifier", Flag: "engagement-invitation-identifier", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OpportunityIdentifier", Flag: "opportunity-identifier", Type: "[]string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListTasksSortBase", Required: false},
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "[]string", Required: false},
	{Name: "TaskStatus", Flag: "task-status", Type: "[]types.TaskStatus", Required: false},
}

var fields_list_engagement_from_opportunity_tasks = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OpportunityIdentifier", Flag: "opportunity-identifier", Type: "[]string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListTasksSortBase", Required: false},
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "[]string", Required: false},
	{Name: "TaskStatus", Flag: "task-status", Type: "[]types.TaskStatus", Required: false},
}

var fields_list_engagement_invitations = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParticipantType", Flag: "participant-type", Type: "types.ParticipantType", Required: true},
	{Name: "PayloadType", Flag: "payload-type", Type: "[]types.EngagementInvitationPayloadType", Required: false},
	{Name: "SenderAwsAccountId", Flag: "sender-aws-account-id", Type: "[]string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.OpportunityEngagementInvitationSort", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.InvitationStatus", Required: false},
}

var fields_list_engagement_members = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_engagement_resource_associations = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "CreatedBy", Flag: "created-by", Type: "*string", Required: false},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: false},
}

var fields_list_engagements = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ContextTypes", Flag: "context-types", Type: "[]types.EngagementContextType", Required: false},
	{Name: "CreatedBy", Flag: "created-by", Type: "[]string", Required: false},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "[]string", Required: false},
	{Name: "ExcludeContextTypes", Flag: "exclude-context-types", Type: "[]types.EngagementContextType", Required: false},
	{Name: "ExcludeCreatedBy", Flag: "exclude-created-by", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.EngagementSort", Required: false},
}

var fields_list_opportunities = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "CreatedDate", Flag: "created-date", Type: "*types.CreatedDateFilter", Required: false},
	{Name: "CustomerCompanyName", Flag: "customer-company-name", Type: "[]string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "[]string", Required: false},
	{Name: "LastModifiedDate", Flag: "last-modified-date", Type: "*types.LastModifiedDate", Required: false},
	{Name: "LifeCycleReviewStatus", Flag: "life-cycle-review-status", Type: "[]types.ReviewStatus", Required: false},
	{Name: "LifeCycleStage", Flag: "life-cycle-stage", Type: "[]types.Stage", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.OpportunitySort", Required: false},
	{Name: "TargetCloseDate", Flag: "target-close-date", Type: "*types.TargetCloseDateFilter", Required: false},
}

var fields_list_opportunity_from_engagement_tasks = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ContextIdentifier", Flag: "context-identifier", Type: "[]string", Required: false},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OpportunityIdentifier", Flag: "opportunity-identifier", Type: "[]string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListTasksSortBase", Required: false},
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "[]string", Required: false},
	{Name: "TaskStatus", Flag: "task-status", Type: "[]types.TaskStatus", Required: false},
}

var fields_list_resource_snapshot_jobs = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SortObject", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ResourceSnapshotJobStatus", Required: false},
}

var fields_list_resource_snapshots = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "CreatedBy", Flag: "created-by", Type: "*string", Required: false},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "ResourceSnapshotTemplateIdentifier", Flag: "resource-snapshot-template-identifier", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: false},
}

var fields_list_solutions = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Category", Flag: "category", Type: "[]string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SolutionSort", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.SolutionStatus", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_selling_system_settings = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceSnapshotJobRoleIdentifier", Flag: "resource-snapshot-job-role-identifier", Type: "*string", Required: false},
}

var fields_reject_engagement_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RejectionReason", Flag: "rejection-reason", Type: "*string", Required: false},
}

var fields_start_engagement_by_accepting_invitation_task = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_engagement_from_opportunity_task = []leanruntime.Field{
	{Name: "AwsSubmission", Flag: "aws-submission", Type: "*types.AwsSubmission", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_opportunity_from_engagement_task = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContextIdentifier", Flag: "context-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_resource_snapshot_job = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceSnapshotJobIdentifier", Flag: "resource-snapshot-job-identifier", Type: "*string", Required: true},
}

var fields_stop_resource_snapshot_job = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceSnapshotJobIdentifier", Flag: "resource-snapshot-job-identifier", Type: "*string", Required: true},
}

var fields_submit_opportunity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "InvolvementType", Flag: "involvement-type", Type: "types.SalesInvolvementType", Required: true},
	{Name: "Visibility", Flag: "visibility", Type: "types.Visibility", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_engagement_context = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ContextIdentifier", Flag: "context-identifier", Type: "*string", Required: true},
	{Name: "EngagementIdentifier", Flag: "engagement-identifier", Type: "*string", Required: true},
	{Name: "EngagementLastModifiedAt", Flag: "engagement-last-modified-at", Type: "*time.Time", Required: true},
	{Name: "Payload", Flag: "payload", Type: "types.UpdateEngagementContextPayload", Required: true},
	{Name: "Type", Flag: "type", Type: "types.EngagementContextType", Required: true},
}

var fields_update_opportunity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Customer", Flag: "customer", Type: "*types.Customer", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LastModifiedDate", Flag: "last-modified-date", Type: "*time.Time", Required: true},
	{Name: "LifeCycle", Flag: "life-cycle", Type: "*types.LifeCycle", Required: false},
	{Name: "Marketing", Flag: "marketing", Type: "*types.Marketing", Required: false},
	{Name: "NationalSecurity", Flag: "national-security", Type: "types.NationalSecurity", Required: false},
	{Name: "OpportunityType", Flag: "opportunity-type", Type: "types.OpportunityType", Required: false},
	{Name: "PartnerOpportunityIdentifier", Flag: "partner-opportunity-identifier", Type: "*string", Required: false},
	{Name: "PrimaryNeedsFromAws", Flag: "primary-needs-from-aws", Type: "[]types.PrimaryNeedFromAws", Required: false},
	{Name: "Project", Flag: "project", Type: "*types.Project", Required: false},
	{Name: "SoftwareRevenue", Flag: "software-revenue", Type: "*types.SoftwareRevenue", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-engagement-invitation": {
			Name:   "accept-engagement-invitation",
			Fields: fields_accept_engagement_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptEngagementInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_engagement_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptEngagementInvitation(ctx, input)
			},
		},
		"assign-opportunity": {
			Name:   "assign-opportunity",
			Fields: fields_assign_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssignOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assign_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssignOpportunity(ctx, input)
			},
		},
		"associate-opportunity": {
			Name:   "associate-opportunity",
			Fields: fields_associate_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateOpportunity(ctx, input)
			},
		},
		"create-engagement": {
			Name:   "create-engagement",
			Fields: fields_create_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEngagement(ctx, input)
			},
		},
		"create-engagement-context": {
			Name:   "create-engagement-context",
			Fields: fields_create_engagement_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEngagementContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_engagement_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEngagementContext(ctx, input)
			},
		},
		"create-engagement-invitation": {
			Name:   "create-engagement-invitation",
			Fields: fields_create_engagement_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEngagementInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_engagement_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEngagementInvitation(ctx, input)
			},
		},
		"create-opportunity": {
			Name:   "create-opportunity",
			Fields: fields_create_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOpportunity(ctx, input)
			},
		},
		"create-resource-snapshot": {
			Name:   "create-resource-snapshot",
			Fields: fields_create_resource_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceSnapshot(ctx, input)
			},
		},
		"create-resource-snapshot-job": {
			Name:   "create-resource-snapshot-job",
			Fields: fields_create_resource_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceSnapshotJob(ctx, input)
			},
		},
		"delete-resource-snapshot-job": {
			Name:   "delete-resource-snapshot-job",
			Fields: fields_delete_resource_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceSnapshotJob(ctx, input)
			},
		},
		"disassociate-opportunity": {
			Name:   "disassociate-opportunity",
			Fields: fields_disassociate_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateOpportunity(ctx, input)
			},
		},
		"get-aws-opportunity-summary": {
			Name:   "get-aws-opportunity-summary",
			Fields: fields_get_aws_opportunity_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAwsOpportunitySummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_aws_opportunity_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAwsOpportunitySummary(ctx, input)
			},
		},
		"get-engagement": {
			Name:   "get-engagement",
			Fields: fields_get_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEngagement(ctx, input)
			},
		},
		"get-engagement-invitation": {
			Name:   "get-engagement-invitation",
			Fields: fields_get_engagement_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEngagementInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_engagement_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEngagementInvitation(ctx, input)
			},
		},
		"get-opportunity": {
			Name:   "get-opportunity",
			Fields: fields_get_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpportunity(ctx, input)
			},
		},
		"get-resource-snapshot": {
			Name:   "get-resource-snapshot",
			Fields: fields_get_resource_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceSnapshot(ctx, input)
			},
		},
		"get-resource-snapshot-job": {
			Name:   "get-resource-snapshot-job",
			Fields: fields_get_resource_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceSnapshotJob(ctx, input)
			},
		},
		"get-selling-system-settings": {
			Name:   "get-selling-system-settings",
			Fields: fields_get_selling_system_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSellingSystemSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_selling_system_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSellingSystemSettings(ctx, input)
			},
		},
		"list-engagement-by-accepting-invitation-tasks": {
			Name:   "list-engagement-by-accepting-invitation-tasks",
			Fields: fields_list_engagement_by_accepting_invitation_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementByAcceptingInvitationTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagement_by_accepting_invitation_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagementByAcceptingInvitationTasks(ctx, input)
				}
				var results []*svc.ListEngagementByAcceptingInvitationTasksOutput
				p := svc.NewListEngagementByAcceptingInvitationTasksPaginator(client, input)
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
		"list-engagement-from-opportunity-tasks": {
			Name:   "list-engagement-from-opportunity-tasks",
			Fields: fields_list_engagement_from_opportunity_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementFromOpportunityTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagement_from_opportunity_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagementFromOpportunityTasks(ctx, input)
				}
				var results []*svc.ListEngagementFromOpportunityTasksOutput
				p := svc.NewListEngagementFromOpportunityTasksPaginator(client, input)
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
		"list-engagement-invitations": {
			Name:   "list-engagement-invitations",
			Fields: fields_list_engagement_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagement_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagementInvitations(ctx, input)
				}
				var results []*svc.ListEngagementInvitationsOutput
				p := svc.NewListEngagementInvitationsPaginator(client, input)
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
		"list-engagement-members": {
			Name:   "list-engagement-members",
			Fields: fields_list_engagement_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagement_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagementMembers(ctx, input)
				}
				var results []*svc.ListEngagementMembersOutput
				p := svc.NewListEngagementMembersPaginator(client, input)
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
		"list-engagement-resource-associations": {
			Name:   "list-engagement-resource-associations",
			Fields: fields_list_engagement_resource_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementResourceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagement_resource_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagementResourceAssociations(ctx, input)
				}
				var results []*svc.ListEngagementResourceAssociationsOutput
				p := svc.NewListEngagementResourceAssociationsPaginator(client, input)
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
		"list-engagements": {
			Name:   "list-engagements",
			Fields: fields_list_engagements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagements(ctx, input)
				}
				var results []*svc.ListEngagementsOutput
				p := svc.NewListEngagementsPaginator(client, input)
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
		"list-opportunities": {
			Name:   "list-opportunities",
			Fields: fields_list_opportunities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpportunitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_opportunities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpportunities(ctx, input)
				}
				var results []*svc.ListOpportunitiesOutput
				p := svc.NewListOpportunitiesPaginator(client, input)
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
		"list-opportunity-from-engagement-tasks": {
			Name:   "list-opportunity-from-engagement-tasks",
			Fields: fields_list_opportunity_from_engagement_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpportunityFromEngagementTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_opportunity_from_engagement_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpportunityFromEngagementTasks(ctx, input)
				}
				var results []*svc.ListOpportunityFromEngagementTasksOutput
				p := svc.NewListOpportunityFromEngagementTasksPaginator(client, input)
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
		"list-resource-snapshot-jobs": {
			Name:   "list-resource-snapshot-jobs",
			Fields: fields_list_resource_snapshot_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceSnapshotJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_snapshot_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceSnapshotJobs(ctx, input)
				}
				var results []*svc.ListResourceSnapshotJobsOutput
				p := svc.NewListResourceSnapshotJobsPaginator(client, input)
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
		"list-resource-snapshots": {
			Name:   "list-resource-snapshots",
			Fields: fields_list_resource_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceSnapshots(ctx, input)
				}
				var results []*svc.ListResourceSnapshotsOutput
				p := svc.NewListResourceSnapshotsPaginator(client, input)
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
		"list-solutions": {
			Name:   "list-solutions",
			Fields: fields_list_solutions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_solutions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolutions(ctx, input)
				}
				var results []*svc.ListSolutionsOutput
				p := svc.NewListSolutionsPaginator(client, input)
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
		"put-selling-system-settings": {
			Name:   "put-selling-system-settings",
			Fields: fields_put_selling_system_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSellingSystemSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_selling_system_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSellingSystemSettings(ctx, input)
			},
		},
		"reject-engagement-invitation": {
			Name:   "reject-engagement-invitation",
			Fields: fields_reject_engagement_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectEngagementInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_engagement_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectEngagementInvitation(ctx, input)
			},
		},
		"start-engagement-by-accepting-invitation-task": {
			Name:   "start-engagement-by-accepting-invitation-task",
			Fields: fields_start_engagement_by_accepting_invitation_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEngagementByAcceptingInvitationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_engagement_by_accepting_invitation_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEngagementByAcceptingInvitationTask(ctx, input)
			},
		},
		"start-engagement-from-opportunity-task": {
			Name:   "start-engagement-from-opportunity-task",
			Fields: fields_start_engagement_from_opportunity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEngagementFromOpportunityTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_engagement_from_opportunity_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEngagementFromOpportunityTask(ctx, input)
			},
		},
		"start-opportunity-from-engagement-task": {
			Name:   "start-opportunity-from-engagement-task",
			Fields: fields_start_opportunity_from_engagement_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOpportunityFromEngagementTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_opportunity_from_engagement_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOpportunityFromEngagementTask(ctx, input)
			},
		},
		"start-resource-snapshot-job": {
			Name:   "start-resource-snapshot-job",
			Fields: fields_start_resource_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartResourceSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_resource_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartResourceSnapshotJob(ctx, input)
			},
		},
		"stop-resource-snapshot-job": {
			Name:   "stop-resource-snapshot-job",
			Fields: fields_stop_resource_snapshot_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopResourceSnapshotJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_resource_snapshot_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopResourceSnapshotJob(ctx, input)
			},
		},
		"submit-opportunity": {
			Name:   "submit-opportunity",
			Fields: fields_submit_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitOpportunity(ctx, input)
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
		"update-engagement-context": {
			Name:   "update-engagement-context",
			Fields: fields_update_engagement_context,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEngagementContextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_engagement_context, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEngagementContext(ctx, input)
			},
		},
		"update-opportunity": {
			Name:   "update-opportunity",
			Fields: fields_update_opportunity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOpportunityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_opportunity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOpportunity(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("partnercentralselling", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
