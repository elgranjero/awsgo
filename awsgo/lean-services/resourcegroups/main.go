package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
)

var fields_cancel_tag_sync_task = []leanruntime.Field{
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
}

var fields_create_group = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "[]types.GroupConfigurationItem", Required: false},
	{Name: "Criticality", Flag: "criticality", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
	{Name: "ResourceQuery", Flag: "resource-query", Type: "*types.ResourceQuery", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
}

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_group = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
}

var fields_get_group_configuration = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
}

var fields_get_group_query = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
}

var fields_get_tag_sync_task = []leanruntime.Field{
	{Name: "TaskArn", Flag: "task-arn", Type: "*string", Required: true},
}

var fields_get_tags = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_group_resources = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: true},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
}

var fields_list_group_resources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ResourceFilter", Required: false},
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_grouping_statuses = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListGroupingStatusesFilter", Required: false},
	{Name: "Group", Flag: "group", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.GroupFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tag_sync_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListTagSyncTasksFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_group_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "[]types.GroupConfigurationItem", Required: false},
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
}

var fields_search_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceQuery", Flag: "resource-query", Type: "*types.ResourceQuery", Required: true},
}

var fields_start_tag_sync_task = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: true},
	{Name: "ResourceQuery", Flag: "resource-query", Type: "*types.ResourceQuery", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: false},
	{Name: "TagValue", Flag: "tag-value", Type: "*string", Required: false},
}

var fields_tag = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_ungroup_resources = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: true},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
}

var fields_untag = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Keys", Flag: "keys", Type: "[]string", Required: true},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "GroupLifecycleEventsDesiredStatus", Flag: "group-lifecycle-events-desired-status", Type: "types.GroupLifecycleEventsDesiredStatus", Required: false},
}

var fields_update_group = []leanruntime.Field{
	{Name: "Criticality", Flag: "criticality", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
}

var fields_update_group_query = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "ResourceQuery", Flag: "resource-query", Type: "*types.ResourceQuery", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-tag-sync-task": {
			Name:   "cancel-tag-sync-task",
			Fields: fields_cancel_tag_sync_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTagSyncTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_tag_sync_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTagSyncTask(ctx, input)
			},
		},
		"create-group": {
			Name:   "create-group",
			Fields: fields_create_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroup(ctx, input)
			},
		},
		"delete-group": {
			Name:   "delete-group",
			Fields: fields_delete_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroup(ctx, input)
			},
		},
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
			},
		},
		"get-group": {
			Name:   "get-group",
			Fields: fields_get_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroup(ctx, input)
			},
		},
		"get-group-configuration": {
			Name:   "get-group-configuration",
			Fields: fields_get_group_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupConfiguration(ctx, input)
			},
		},
		"get-group-query": {
			Name:   "get-group-query",
			Fields: fields_get_group_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupQuery(ctx, input)
			},
		},
		"get-tag-sync-task": {
			Name:   "get-tag-sync-task",
			Fields: fields_get_tag_sync_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagSyncTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tag_sync_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTagSyncTask(ctx, input)
			},
		},
		"get-tags": {
			Name:   "get-tags",
			Fields: fields_get_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTags(ctx, input)
			},
		},
		"group-resources": {
			Name:   "group-resources",
			Fields: fields_group_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GroupResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_group_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GroupResources(ctx, input)
			},
		},
		"list-group-resources": {
			Name:   "list-group-resources",
			Fields: fields_list_group_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupResources(ctx, input)
				}
				var results []*svc.ListGroupResourcesOutput
				p := svc.NewListGroupResourcesPaginator(client, input)
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
		"list-grouping-statuses": {
			Name:   "list-grouping-statuses",
			Fields: fields_list_grouping_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupingStatusesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_grouping_statuses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupingStatuses(ctx, input)
				}
				var results []*svc.ListGroupingStatusesOutput
				p := svc.NewListGroupingStatusesPaginator(client, input)
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
		"list-groups": {
			Name:   "list-groups",
			Fields: fields_list_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroups(ctx, input)
				}
				var results []*svc.ListGroupsOutput
				p := svc.NewListGroupsPaginator(client, input)
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
		"list-tag-sync-tasks": {
			Name:   "list-tag-sync-tasks",
			Fields: fields_list_tag_sync_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagSyncTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tag_sync_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagSyncTasks(ctx, input)
				}
				var results []*svc.ListTagSyncTasksOutput
				p := svc.NewListTagSyncTasksPaginator(client, input)
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
		"put-group-configuration": {
			Name:   "put-group-configuration",
			Fields: fields_put_group_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGroupConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_group_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGroupConfiguration(ctx, input)
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
		"start-tag-sync-task": {
			Name:   "start-tag-sync-task",
			Fields: fields_start_tag_sync_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTagSyncTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_tag_sync_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTagSyncTask(ctx, input)
			},
		},
		"tag": {
			Name:   "tag",
			Fields: fields_tag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Tag(ctx, input)
			},
		},
		"ungroup-resources": {
			Name:   "ungroup-resources",
			Fields: fields_ungroup_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UngroupResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_ungroup_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UngroupResources(ctx, input)
			},
		},
		"untag": {
			Name:   "untag",
			Fields: fields_untag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Untag(ctx, input)
			},
		},
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-group": {
			Name:   "update-group",
			Fields: fields_update_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroup(ctx, input)
			},
		},
		"update-group-query": {
			Name:   "update-group-query",
			Fields: fields_update_group_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroupQuery(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("resourcegroups", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
