package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssmincidents"
)

var fields_batch_get_incident_findings = []leanruntime.Field{
	{Name: "FindingIds", Flag: "finding-ids", Type: "[]string", Required: true},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
}

var fields_create_replication_set = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "map[string]types.RegionMapInputValue", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_response_plan = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.Action", Required: false},
	{Name: "ChatChannel", Flag: "chat-channel", Type: "types.ChatChannel", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Engagements", Flag: "engagements", Type: "[]string", Required: false},
	{Name: "IncidentTemplate", Flag: "incident-template", Type: "*types.IncidentTemplate", Required: true},
	{Name: "Integrations", Flag: "integrations", Type: "[]types.Integration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_timeline_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EventData", Flag: "event-data", Type: "*string", Required: true},
	{Name: "EventReferences", Flag: "event-references", Type: "[]types.EventReference", Required: false},
	{Name: "EventTime", Flag: "event-time", Type: "*time.Time", Required: true},
	{Name: "EventType", Flag: "event-type", Type: "*string", Required: true},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
}

var fields_delete_incident_record = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_replication_set = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_response_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_timeline_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
}

var fields_get_incident_record = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_replication_set = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_resource_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_response_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_timeline_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
}

var fields_list_incident_findings = []leanruntime.Field{
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_incident_records = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_related_items = []leanruntime.Field{
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_replication_sets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_response_plans = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_timeline_events = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.TimelineEventSort", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_incident = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Impact", Flag: "impact", Type: "*int32", Required: false},
	{Name: "RelatedItems", Flag: "related-items", Type: "[]types.RelatedItem", Required: false},
	{Name: "ResponsePlanArn", Flag: "response-plan-arn", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "TriggerDetails", Flag: "trigger-details", Type: "*types.TriggerDetails", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_deletion_protection = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeletionProtected", Flag: "deletion-protected", Type: "*bool", Required: true},
}

var fields_update_incident_record = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ChatChannel", Flag: "chat-channel", Type: "types.ChatChannel", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Impact", Flag: "impact", Type: "*int32", Required: false},
	{Name: "NotificationTargets", Flag: "notification-targets", Type: "[]types.NotificationTargetItem", Required: false},
	{Name: "Status", Flag: "status", Type: "types.IncidentRecordStatus", Required: false},
	{Name: "Summary", Flag: "summary", Type: "*string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
}

var fields_update_related_items = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
	{Name: "RelatedItemsUpdate", Flag: "related-items-update", Type: "types.RelatedItemsUpdate", Required: true},
}

var fields_update_replication_set = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.UpdateReplicationSetAction", Required: true},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_update_response_plan = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.Action", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ChatChannel", Flag: "chat-channel", Type: "types.ChatChannel", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Engagements", Flag: "engagements", Type: "[]string", Required: false},
	{Name: "IncidentTemplateDedupeString", Flag: "incident-template-dedupe-string", Type: "*string", Required: false},
	{Name: "IncidentTemplateImpact", Flag: "incident-template-impact", Type: "*int32", Required: false},
	{Name: "IncidentTemplateNotificationTargets", Flag: "incident-template-notification-targets", Type: "[]types.NotificationTargetItem", Required: false},
	{Name: "IncidentTemplateSummary", Flag: "incident-template-summary", Type: "*string", Required: false},
	{Name: "IncidentTemplateTags", Flag: "incident-template-tags", Type: "map[string]string", Required: false},
	{Name: "IncidentTemplateTitle", Flag: "incident-template-title", Type: "*string", Required: false},
	{Name: "Integrations", Flag: "integrations", Type: "[]types.Integration", Required: false},
}

var fields_update_timeline_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EventData", Flag: "event-data", Type: "*string", Required: false},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventReferences", Flag: "event-references", Type: "[]types.EventReference", Required: false},
	{Name: "EventTime", Flag: "event-time", Type: "*time.Time", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "*string", Required: false},
	{Name: "IncidentRecordArn", Flag: "incident-record-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-incident-findings": {
			Name:   "batch-get-incident-findings",
			Fields: fields_batch_get_incident_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetIncidentFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_incident_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetIncidentFindings(ctx, input)
			},
		},
		"create-replication-set": {
			Name:   "create-replication-set",
			Fields: fields_create_replication_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationSet(ctx, input)
			},
		},
		"create-response-plan": {
			Name:   "create-response-plan",
			Fields: fields_create_response_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResponsePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_response_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResponsePlan(ctx, input)
			},
		},
		"create-timeline-event": {
			Name:   "create-timeline-event",
			Fields: fields_create_timeline_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTimelineEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_timeline_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTimelineEvent(ctx, input)
			},
		},
		"delete-incident-record": {
			Name:   "delete-incident-record",
			Fields: fields_delete_incident_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIncidentRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_incident_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIncidentRecord(ctx, input)
			},
		},
		"delete-replication-set": {
			Name:   "delete-replication-set",
			Fields: fields_delete_replication_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationSet(ctx, input)
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
		"delete-response-plan": {
			Name:   "delete-response-plan",
			Fields: fields_delete_response_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResponsePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_response_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResponsePlan(ctx, input)
			},
		},
		"delete-timeline-event": {
			Name:   "delete-timeline-event",
			Fields: fields_delete_timeline_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTimelineEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_timeline_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTimelineEvent(ctx, input)
			},
		},
		"get-incident-record": {
			Name:   "get-incident-record",
			Fields: fields_get_incident_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIncidentRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_incident_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIncidentRecord(ctx, input)
			},
		},
		"get-replication-set": {
			Name:   "get-replication-set",
			Fields: fields_get_replication_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReplicationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_replication_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReplicationSet(ctx, input)
			},
		},
		"get-resource-policies": {
			Name:   "get-resource-policies",
			Fields: fields_get_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourcePolicies(ctx, input)
				}
				var results []*svc.GetResourcePoliciesOutput
				p := svc.NewGetResourcePoliciesPaginator(client, input)
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
		"get-response-plan": {
			Name:   "get-response-plan",
			Fields: fields_get_response_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResponsePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_response_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResponsePlan(ctx, input)
			},
		},
		"get-timeline-event": {
			Name:   "get-timeline-event",
			Fields: fields_get_timeline_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTimelineEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_timeline_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTimelineEvent(ctx, input)
			},
		},
		"list-incident-findings": {
			Name:   "list-incident-findings",
			Fields: fields_list_incident_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIncidentFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_incident_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIncidentFindings(ctx, input)
				}
				var results []*svc.ListIncidentFindingsOutput
				p := svc.NewListIncidentFindingsPaginator(client, input)
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
		"list-incident-records": {
			Name:   "list-incident-records",
			Fields: fields_list_incident_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIncidentRecordsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_incident_records, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIncidentRecords(ctx, input)
				}
				var results []*svc.ListIncidentRecordsOutput
				p := svc.NewListIncidentRecordsPaginator(client, input)
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
		"list-related-items": {
			Name:   "list-related-items",
			Fields: fields_list_related_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRelatedItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_related_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRelatedItems(ctx, input)
				}
				var results []*svc.ListRelatedItemsOutput
				p := svc.NewListRelatedItemsPaginator(client, input)
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
		"list-replication-sets": {
			Name:   "list-replication-sets",
			Fields: fields_list_replication_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReplicationSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_replication_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReplicationSets(ctx, input)
				}
				var results []*svc.ListReplicationSetsOutput
				p := svc.NewListReplicationSetsPaginator(client, input)
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
		"list-response-plans": {
			Name:   "list-response-plans",
			Fields: fields_list_response_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResponsePlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_response_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResponsePlans(ctx, input)
				}
				var results []*svc.ListResponsePlansOutput
				p := svc.NewListResponsePlansPaginator(client, input)
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
		"list-timeline-events": {
			Name:   "list-timeline-events",
			Fields: fields_list_timeline_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTimelineEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_timeline_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTimelineEvents(ctx, input)
				}
				var results []*svc.ListTimelineEventsOutput
				p := svc.NewListTimelineEventsPaginator(client, input)
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
		"start-incident": {
			Name:   "start-incident",
			Fields: fields_start_incident,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartIncidentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_incident, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartIncident(ctx, input)
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
		"update-deletion-protection": {
			Name:   "update-deletion-protection",
			Fields: fields_update_deletion_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeletionProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_deletion_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeletionProtection(ctx, input)
			},
		},
		"update-incident-record": {
			Name:   "update-incident-record",
			Fields: fields_update_incident_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIncidentRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_incident_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIncidentRecord(ctx, input)
			},
		},
		"update-related-items": {
			Name:   "update-related-items",
			Fields: fields_update_related_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRelatedItemsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_related_items, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRelatedItems(ctx, input)
			},
		},
		"update-replication-set": {
			Name:   "update-replication-set",
			Fields: fields_update_replication_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReplicationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_replication_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReplicationSet(ctx, input)
			},
		},
		"update-response-plan": {
			Name:   "update-response-plan",
			Fields: fields_update_response_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResponsePlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_response_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResponsePlan(ctx, input)
			},
		},
		"update-timeline-event": {
			Name:   "update-timeline-event",
			Fields: fields_update_timeline_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTimelineEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_timeline_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTimelineEvent(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssmincidents", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
