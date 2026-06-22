package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/scheduler"
)

var fields_create_schedule = []leanruntime.Field{
	{Name: "ActionAfterCompletion", Flag: "action-after-completion", Type: "types.ActionAfterCompletion", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "FlexibleTimeWindow", Flag: "flexible-time-window", Type: "*types.FlexibleTimeWindow", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: true},
	{Name: "ScheduleExpressionTimezone", Flag: "schedule-expression-timezone", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScheduleState", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.Target", Required: true},
}

var fields_create_schedule_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_schedule = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_schedule_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_schedule = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_schedule_group = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_schedule_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_schedules = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScheduleState", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_schedule = []leanruntime.Field{
	{Name: "ActionAfterCompletion", Flag: "action-after-completion", Type: "types.ActionAfterCompletion", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "FlexibleTimeWindow", Flag: "flexible-time-window", Type: "*types.FlexibleTimeWindow", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: true},
	{Name: "ScheduleExpressionTimezone", Flag: "schedule-expression-timezone", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
	{Name: "State", Flag: "state", Type: "types.ScheduleState", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.Target", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-schedule": {
			Name:   "create-schedule",
			Fields: fields_create_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchedule(ctx, input)
			},
		},
		"create-schedule-group": {
			Name:   "create-schedule-group",
			Fields: fields_create_schedule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schedule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScheduleGroup(ctx, input)
			},
		},
		"delete-schedule": {
			Name:   "delete-schedule",
			Fields: fields_delete_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchedule(ctx, input)
			},
		},
		"delete-schedule-group": {
			Name:   "delete-schedule-group",
			Fields: fields_delete_schedule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schedule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduleGroup(ctx, input)
			},
		},
		"get-schedule": {
			Name:   "get-schedule",
			Fields: fields_get_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchedule(ctx, input)
			},
		},
		"get-schedule-group": {
			Name:   "get-schedule-group",
			Fields: fields_get_schedule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScheduleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schedule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScheduleGroup(ctx, input)
			},
		},
		"list-schedule-groups": {
			Name:   "list-schedule-groups",
			Fields: fields_list_schedule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScheduleGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schedule_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScheduleGroups(ctx, input)
				}
				var results []*svc.ListScheduleGroupsOutput
				p := svc.NewListScheduleGroupsPaginator(client, input)
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
		"list-schedules": {
			Name:   "list-schedules",
			Fields: fields_list_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchedulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schedules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchedules(ctx, input)
				}
				var results []*svc.ListSchedulesOutput
				p := svc.NewListSchedulesPaginator(client, input)
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
		"update-schedule": {
			Name:   "update-schedule",
			Fields: fields_update_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchedule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("scheduler", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
