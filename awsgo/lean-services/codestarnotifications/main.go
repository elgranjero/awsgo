package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
)

var fields_create_notification_rule = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DetailType", Flag: "detail-type", Type: "types.DetailType", Required: true},
	{Name: "EventTypeIds", Flag: "event-type-ids", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.NotificationRuleStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
}

var fields_delete_notification_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_target = []leanruntime.Field{
	{Name: "ForceUnsubscribeAll", Flag: "force-unsubscribe-all", Type: "bool", Required: false},
	{Name: "TargetAddress", Flag: "target-address", Type: "*string", Required: true},
}

var fields_describe_notification_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_event_types = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListEventTypesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_notification_rules = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListNotificationRulesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_targets = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListTargetsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_subscribe = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.Target", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_unsubscribe = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TargetAddress", Flag: "target-address", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_notification_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "DetailType", Flag: "detail-type", Type: "types.DetailType", Required: false},
	{Name: "EventTypeIds", Flag: "event-type-ids", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.NotificationRuleStatus", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-notification-rule": {
			Name:   "create-notification-rule",
			Fields: fields_create_notification_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotificationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notification_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotificationRule(ctx, input)
			},
		},
		"delete-notification-rule": {
			Name:   "delete-notification-rule",
			Fields: fields_delete_notification_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotificationRule(ctx, input)
			},
		},
		"delete-target": {
			Name:   "delete-target",
			Fields: fields_delete_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTarget(ctx, input)
			},
		},
		"describe-notification-rule": {
			Name:   "describe-notification-rule",
			Fields: fields_describe_notification_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotificationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_notification_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNotificationRule(ctx, input)
			},
		},
		"list-event-types": {
			Name:   "list-event-types",
			Fields: fields_list_event_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventTypes(ctx, input)
				}
				var results []*svc.ListEventTypesOutput
				p := svc.NewListEventTypesPaginator(client, input)
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
		"list-notification-rules": {
			Name:   "list-notification-rules",
			Fields: fields_list_notification_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notification_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotificationRules(ctx, input)
				}
				var results []*svc.ListNotificationRulesOutput
				p := svc.NewListNotificationRulesPaginator(client, input)
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
		"list-targets": {
			Name:   "list-targets",
			Fields: fields_list_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargets(ctx, input)
				}
				var results []*svc.ListTargetsOutput
				p := svc.NewListTargetsPaginator(client, input)
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
		"subscribe": {
			Name:   "subscribe",
			Fields: fields_subscribe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubscribeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_subscribe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Subscribe(ctx, input)
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
		"unsubscribe": {
			Name:   "unsubscribe",
			Fields: fields_unsubscribe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnsubscribeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unsubscribe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Unsubscribe(ctx, input)
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
		"update-notification-rule": {
			Name:   "update-notification-rule",
			Fields: fields_update_notification_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotificationRule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codestarnotifications", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
