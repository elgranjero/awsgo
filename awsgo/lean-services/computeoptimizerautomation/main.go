package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/computeoptimizerautomation"
)

var fields_associate_accounts = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_create_automation_rule = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Criteria", Flag: "criteria", Type: "*types.Criteria", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationConfiguration", Flag: "organization-configuration", Type: "*types.OrganizationConfiguration", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*string", Required: false},
	{Name: "RecommendedActionTypes", Flag: "recommended-action-types", Type: "[]types.RecommendedActionType", Required: true},
	{Name: "RuleType", Flag: "rule-type", Type: "types.RuleType", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.Schedule", Required: true},
	{Name: "Status", Flag: "status", Type: "types.RuleStatus", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_automation_rule = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RuleArn", Flag: "rule-arn", Type: "*string", Required: true},
	{Name: "RuleRevision", Flag: "rule-revision", Type: "*int64", Required: true},
}

var fields_disassociate_accounts = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_get_automation_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
}

var fields_get_automation_rule = []leanruntime.Field{
	{Name: "RuleArn", Flag: "rule-arn", Type: "*string", Required: true},
}

var fields_get_enrollment_configuration = []leanruntime.Field{}

var fields_list_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_automation_event_steps = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_automation_event_summaries = []leanruntime.Field{
	{Name: "EndDateExclusive", Flag: "end-date-exclusive", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.AutomationEventFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDateInclusive", Flag: "start-date-inclusive", Type: "*string", Required: false},
}

var fields_list_automation_events = []leanruntime.Field{
	{Name: "EndTimeExclusive", Flag: "end-time-exclusive", Type: "*time.Time", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.AutomationEventFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTimeInclusive", Flag: "start-time-inclusive", Type: "*time.Time", Required: false},
}

var fields_list_automation_rule_preview = []leanruntime.Field{
	{Name: "Criteria", Flag: "criteria", Type: "*types.Criteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationScope", Flag: "organization-scope", Type: "*types.OrganizationScope", Required: false},
	{Name: "RecommendedActionTypes", Flag: "recommended-action-types", Type: "[]types.RecommendedActionType", Required: true},
	{Name: "RuleType", Flag: "rule-type", Type: "types.RuleType", Required: true},
}

var fields_list_automation_rule_preview_summaries = []leanruntime.Field{
	{Name: "Criteria", Flag: "criteria", Type: "*types.Criteria", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationScope", Flag: "organization-scope", Type: "*types.OrganizationScope", Required: false},
	{Name: "RecommendedActionTypes", Flag: "recommended-action-types", Type: "[]types.RecommendedActionType", Required: true},
	{Name: "RuleType", Flag: "rule-type", Type: "types.RuleType", Required: true},
}

var fields_list_automation_rules = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recommended_action_summaries = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RecommendedActionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recommended_actions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RecommendedActionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_rollback_automation_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
}

var fields_start_automation_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RecommendedActionId", Flag: "recommended-action-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "RuleRevision", Flag: "rule-revision", Type: "*int64", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "RuleRevision", Flag: "rule-revision", Type: "*int64", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_automation_rule = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Criteria", Flag: "criteria", Type: "*types.Criteria", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OrganizationConfiguration", Flag: "organization-configuration", Type: "*types.OrganizationConfiguration", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*string", Required: false},
	{Name: "RecommendedActionTypes", Flag: "recommended-action-types", Type: "[]types.RecommendedActionType", Required: false},
	{Name: "RuleArn", Flag: "rule-arn", Type: "*string", Required: true},
	{Name: "RuleRevision", Flag: "rule-revision", Type: "*int64", Required: true},
	{Name: "RuleType", Flag: "rule-type", Type: "types.RuleType", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.Schedule", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RuleStatus", Required: false},
}

var fields_update_enrollment_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.EnrollmentStatus", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-accounts": {
			Name:   "associate-accounts",
			Fields: fields_associate_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAccounts(ctx, input)
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
		"delete-automation-rule": {
			Name:   "delete-automation-rule",
			Fields: fields_delete_automation_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutomationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_automation_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutomationRule(ctx, input)
			},
		},
		"disassociate-accounts": {
			Name:   "disassociate-accounts",
			Fields: fields_disassociate_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAccounts(ctx, input)
			},
		},
		"get-automation-event": {
			Name:   "get-automation-event",
			Fields: fields_get_automation_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automation_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomationEvent(ctx, input)
			},
		},
		"get-automation-rule": {
			Name:   "get-automation-rule",
			Fields: fields_get_automation_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutomationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_automation_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutomationRule(ctx, input)
			},
		},
		"get-enrollment-configuration": {
			Name:   "get-enrollment-configuration",
			Fields: fields_get_enrollment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnrollmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_enrollment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnrollmentConfiguration(ctx, input)
			},
		},
		"list-accounts": {
			Name:   "list-accounts",
			Fields: fields_list_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccounts(ctx, input)
				}
				var results []*svc.ListAccountsOutput
				p := svc.NewListAccountsPaginator(client, input)
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
		"list-automation-event-steps": {
			Name:   "list-automation-event-steps",
			Fields: fields_list_automation_event_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationEventStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automation_event_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomationEventSteps(ctx, input)
				}
				var results []*svc.ListAutomationEventStepsOutput
				p := svc.NewListAutomationEventStepsPaginator(client, input)
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
		"list-automation-event-summaries": {
			Name:   "list-automation-event-summaries",
			Fields: fields_list_automation_event_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationEventSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automation_event_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomationEventSummaries(ctx, input)
				}
				var results []*svc.ListAutomationEventSummariesOutput
				p := svc.NewListAutomationEventSummariesPaginator(client, input)
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
		"list-automation-events": {
			Name:   "list-automation-events",
			Fields: fields_list_automation_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automation_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomationEvents(ctx, input)
				}
				var results []*svc.ListAutomationEventsOutput
				p := svc.NewListAutomationEventsPaginator(client, input)
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
		"list-automation-rule-preview": {
			Name:   "list-automation-rule-preview",
			Fields: fields_list_automation_rule_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationRulePreviewInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automation_rule_preview, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomationRulePreview(ctx, input)
				}
				var results []*svc.ListAutomationRulePreviewOutput
				p := svc.NewListAutomationRulePreviewPaginator(client, input)
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
		"list-automation-rule-preview-summaries": {
			Name:   "list-automation-rule-preview-summaries",
			Fields: fields_list_automation_rule_preview_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomationRulePreviewSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automation_rule_preview_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomationRulePreviewSummaries(ctx, input)
				}
				var results []*svc.ListAutomationRulePreviewSummariesOutput
				p := svc.NewListAutomationRulePreviewSummariesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_automation_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutomationRules(ctx, input)
				}
				var results []*svc.ListAutomationRulesOutput
				p := svc.NewListAutomationRulesPaginator(client, input)
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
		"list-recommended-action-summaries": {
			Name:   "list-recommended-action-summaries",
			Fields: fields_list_recommended_action_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendedActionSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommended_action_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendedActionSummaries(ctx, input)
				}
				var results []*svc.ListRecommendedActionSummariesOutput
				p := svc.NewListRecommendedActionSummariesPaginator(client, input)
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
		"list-recommended-actions": {
			Name:   "list-recommended-actions",
			Fields: fields_list_recommended_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendedActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommended_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendedActions(ctx, input)
				}
				var results []*svc.ListRecommendedActionsOutput
				p := svc.NewListRecommendedActionsPaginator(client, input)
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
		"rollback-automation-event": {
			Name:   "rollback-automation-event",
			Fields: fields_rollback_automation_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RollbackAutomationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rollback_automation_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RollbackAutomationEvent(ctx, input)
			},
		},
		"start-automation-event": {
			Name:   "start-automation-event",
			Fields: fields_start_automation_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAutomationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_automation_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAutomationEvent(ctx, input)
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
		"update-automation-rule": {
			Name:   "update-automation-rule",
			Fields: fields_update_automation_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automation_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomationRule(ctx, input)
			},
		},
		"update-enrollment-configuration": {
			Name:   "update-enrollment-configuration",
			Fields: fields_update_enrollment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnrollmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_enrollment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnrollmentConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("computeoptimizerautomation", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
