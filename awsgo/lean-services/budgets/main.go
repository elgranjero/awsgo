package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/budgets"
)

var fields_create_budget = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Budget", Flag: "budget", Type: "*types.Budget", Required: true},
	{Name: "NotificationsWithSubscribers", Flag: "notifications-with-subscribers", Type: "[]types.NotificationWithSubscribers", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
}

var fields_create_budget_action = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ActionThreshold", Flag: "action-threshold", Type: "*types.ActionThreshold", Required: true},
	{Name: "ActionType", Flag: "action-type", Type: "types.ActionType", Required: true},
	{Name: "ApprovalModel", Flag: "approval-model", Type: "types.ApprovalModel", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.Definition", Required: true},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "NotificationType", Flag: "notification-type", Type: "types.NotificationType", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "Subscribers", Flag: "subscribers", Type: "[]types.Subscriber", Required: true},
}

var fields_create_notification = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: true},
	{Name: "Subscribers", Flag: "subscribers", Type: "[]types.Subscriber", Required: true},
}

var fields_create_subscriber = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: true},
	{Name: "Subscriber", Flag: "subscriber", Type: "*types.Subscriber", Required: true},
}

var fields_delete_budget = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
}

var fields_delete_budget_action = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
}

var fields_delete_notification = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: true},
}

var fields_delete_subscriber = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: true},
	{Name: "Subscriber", Flag: "subscriber", Type: "*types.Subscriber", Required: true},
}

var fields_describe_budget = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "ShowFilterExpression", Flag: "show-filter-expression", Type: "*bool", Required: false},
}

var fields_describe_budget_action = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
}

var fields_describe_budget_action_histories = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.TimePeriod", Required: false},
}

var fields_describe_budget_actions_for_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_budget_actions_for_budget = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_budget_notifications_for_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_budget_performance_history = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.TimePeriod", Required: false},
}

var fields_describe_budgets = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShowFilterExpression", Flag: "show-filter-expression", Type: "*bool", Required: false},
}

var fields_describe_notifications_for_budget = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_subscribers_for_notification = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: true},
}

var fields_execute_budget_action = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "ExecutionType", Flag: "execution-type", Type: "types.ExecutionType", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTagKeys", Flag: "resource-tag-keys", Type: "[]string", Required: true},
}

var fields_update_budget = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "NewBudget", Flag: "new-budget", Type: "*types.Budget", Required: true},
}

var fields_update_budget_action = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "ActionThreshold", Flag: "action-threshold", Type: "*types.ActionThreshold", Required: false},
	{Name: "ApprovalModel", Flag: "approval-model", Type: "types.ApprovalModel", Required: false},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "*types.Definition", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "NotificationType", Flag: "notification-type", Type: "types.NotificationType", Required: false},
	{Name: "Subscribers", Flag: "subscribers", Type: "[]types.Subscriber", Required: false},
}

var fields_update_notification = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "NewNotification", Flag: "new-notification", Type: "*types.Notification", Required: true},
	{Name: "OldNotification", Flag: "old-notification", Type: "*types.Notification", Required: true},
}

var fields_update_subscriber = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "NewSubscriber", Flag: "new-subscriber", Type: "*types.Subscriber", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: true},
	{Name: "OldSubscriber", Flag: "old-subscriber", Type: "*types.Subscriber", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-budget": {
			Name:   "create-budget",
			Fields: fields_create_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBudget(ctx, input)
			},
		},
		"create-budget-action": {
			Name:   "create-budget-action",
			Fields: fields_create_budget_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBudgetActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_budget_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBudgetAction(ctx, input)
			},
		},
		"create-notification": {
			Name:   "create-notification",
			Fields: fields_create_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotification(ctx, input)
			},
		},
		"create-subscriber": {
			Name:   "create-subscriber",
			Fields: fields_create_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriber(ctx, input)
			},
		},
		"delete-budget": {
			Name:   "delete-budget",
			Fields: fields_delete_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBudget(ctx, input)
			},
		},
		"delete-budget-action": {
			Name:   "delete-budget-action",
			Fields: fields_delete_budget_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBudgetActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_budget_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBudgetAction(ctx, input)
			},
		},
		"delete-notification": {
			Name:   "delete-notification",
			Fields: fields_delete_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotification(ctx, input)
			},
		},
		"delete-subscriber": {
			Name:   "delete-subscriber",
			Fields: fields_delete_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriber(ctx, input)
			},
		},
		"describe-budget": {
			Name:   "describe-budget",
			Fields: fields_describe_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBudget(ctx, input)
			},
		},
		"describe-budget-action": {
			Name:   "describe-budget-action",
			Fields: fields_describe_budget_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_budget_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBudgetAction(ctx, input)
			},
		},
		"describe-budget-action-histories": {
			Name:   "describe-budget-action-histories",
			Fields: fields_describe_budget_action_histories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetActionHistoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_budget_action_histories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBudgetActionHistories(ctx, input)
				}
				var results []*svc.DescribeBudgetActionHistoriesOutput
				p := svc.NewDescribeBudgetActionHistoriesPaginator(client, input)
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
		"describe-budget-actions-for-account": {
			Name:   "describe-budget-actions-for-account",
			Fields: fields_describe_budget_actions_for_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetActionsForAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_budget_actions_for_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBudgetActionsForAccount(ctx, input)
				}
				var results []*svc.DescribeBudgetActionsForAccountOutput
				p := svc.NewDescribeBudgetActionsForAccountPaginator(client, input)
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
		"describe-budget-actions-for-budget": {
			Name:   "describe-budget-actions-for-budget",
			Fields: fields_describe_budget_actions_for_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetActionsForBudgetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_budget_actions_for_budget, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBudgetActionsForBudget(ctx, input)
				}
				var results []*svc.DescribeBudgetActionsForBudgetOutput
				p := svc.NewDescribeBudgetActionsForBudgetPaginator(client, input)
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
		"describe-budget-notifications-for-account": {
			Name:   "describe-budget-notifications-for-account",
			Fields: fields_describe_budget_notifications_for_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetNotificationsForAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_budget_notifications_for_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBudgetNotificationsForAccount(ctx, input)
				}
				var results []*svc.DescribeBudgetNotificationsForAccountOutput
				p := svc.NewDescribeBudgetNotificationsForAccountPaginator(client, input)
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
		"describe-budget-performance-history": {
			Name:   "describe-budget-performance-history",
			Fields: fields_describe_budget_performance_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetPerformanceHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_budget_performance_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBudgetPerformanceHistory(ctx, input)
				}
				var results []*svc.DescribeBudgetPerformanceHistoryOutput
				p := svc.NewDescribeBudgetPerformanceHistoryPaginator(client, input)
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
		"describe-budgets": {
			Name:   "describe-budgets",
			Fields: fields_describe_budgets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBudgetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_budgets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBudgets(ctx, input)
				}
				var results []*svc.DescribeBudgetsOutput
				p := svc.NewDescribeBudgetsPaginator(client, input)
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
		"describe-notifications-for-budget": {
			Name:   "describe-notifications-for-budget",
			Fields: fields_describe_notifications_for_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotificationsForBudgetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_notifications_for_budget, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNotificationsForBudget(ctx, input)
				}
				var results []*svc.DescribeNotificationsForBudgetOutput
				p := svc.NewDescribeNotificationsForBudgetPaginator(client, input)
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
		"describe-subscribers-for-notification": {
			Name:   "describe-subscribers-for-notification",
			Fields: fields_describe_subscribers_for_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubscribersForNotificationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_subscribers_for_notification, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSubscribersForNotification(ctx, input)
				}
				var results []*svc.DescribeSubscribersForNotificationOutput
				p := svc.NewDescribeSubscribersForNotificationPaginator(client, input)
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
		"execute-budget-action": {
			Name:   "execute-budget-action",
			Fields: fields_execute_budget_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteBudgetActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_budget_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteBudgetAction(ctx, input)
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
		"update-budget": {
			Name:   "update-budget",
			Fields: fields_update_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBudget(ctx, input)
			},
		},
		"update-budget-action": {
			Name:   "update-budget-action",
			Fields: fields_update_budget_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBudgetActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_budget_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBudgetAction(ctx, input)
			},
		},
		"update-notification": {
			Name:   "update-notification",
			Fields: fields_update_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotification(ctx, input)
			},
		},
		"update-subscriber": {
			Name:   "update-subscriber",
			Fields: fields_update_subscriber,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscriber, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriber(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("budgets", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
