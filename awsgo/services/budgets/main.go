package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/budgets/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-budget", "create-budget-action", "create-notification", "create-subscriber", "delete-budget", "delete-budget-action", "delete-notification", "delete-subscriber", "describe-budget", "describe-budget-action", "describe-budget-action-histories", "describe-budget-actions-for-account", "describe-budget-actions-for-budget", "describe-budget-notifications-for-account", "describe-budget-performance-history", "describe-budgets", "describe-notifications-for-budget", "describe-subscribers-for-notification", "execute-budget-action", "list-tags-for-resource", "tag-resource", "untag-resource", "update-budget", "update-budget-action", "update-notification", "update-subscriber"},
		OperationSet: map[string]bool{"create-budget": true, "create-budget-action": true, "create-notification": true, "create-subscriber": true, "delete-budget": true, "delete-budget-action": true, "delete-notification": true, "delete-subscriber": true, "describe-budget": true, "describe-budget-action": true, "describe-budget-action-histories": true, "describe-budget-actions-for-account": true, "describe-budget-actions-for-budget": true, "describe-budget-notifications-for-account": true, "describe-budget-performance-history": true, "describe-budgets": true, "describe-notifications-for-budget": true, "describe-subscribers-for-notification": true, "execute-budget-action": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-budget": true, "update-budget-action": true, "update-notification": true, "update-subscriber": true},
		OperationInputs: map[string][]string{
			"create-budget":                             {"AccountId", "Budget", "NotificationsWithSubscribers", "ResourceTags"},
			"create-budget-action":                      {"AccountId", "ActionThreshold", "ActionType", "ApprovalModel", "BudgetName", "Definition", "ExecutionRoleArn", "NotificationType", "ResourceTags", "Subscribers"},
			"create-notification":                       {"AccountId", "BudgetName", "Notification", "Subscribers"},
			"create-subscriber":                         {"AccountId", "BudgetName", "Notification", "Subscriber"},
			"delete-budget":                             {"AccountId", "BudgetName"},
			"delete-budget-action":                      {"AccountId", "ActionId", "BudgetName"},
			"delete-notification":                       {"AccountId", "BudgetName", "Notification"},
			"delete-subscriber":                         {"AccountId", "BudgetName", "Notification", "Subscriber"},
			"describe-budget":                           {"AccountId", "BudgetName", "ShowFilterExpression"},
			"describe-budget-action":                    {"AccountId", "ActionId", "BudgetName"},
			"describe-budget-action-histories":          {"AccountId", "ActionId", "BudgetName", "MaxResults", "NextToken", "TimePeriod"},
			"describe-budget-actions-for-account":       {"AccountId", "MaxResults", "NextToken"},
			"describe-budget-actions-for-budget":        {"AccountId", "BudgetName", "MaxResults", "NextToken"},
			"describe-budget-notifications-for-account": {"AccountId", "MaxResults", "NextToken"},
			"describe-budget-performance-history":       {"AccountId", "BudgetName", "MaxResults", "NextToken", "TimePeriod"},
			"describe-budgets":                          {"AccountId", "MaxResults", "NextToken", "ShowFilterExpression"},
			"describe-notifications-for-budget":         {"AccountId", "BudgetName", "MaxResults", "NextToken"},
			"describe-subscribers-for-notification":     {"AccountId", "BudgetName", "MaxResults", "NextToken", "Notification"},
			"execute-budget-action":                     {"AccountId", "ActionId", "BudgetName", "ExecutionType"},
			"list-tags-for-resource":                    {"ResourceARN"},
			"tag-resource":                              {"ResourceARN", "ResourceTags"},
			"untag-resource":                            {"ResourceARN", "ResourceTagKeys"},
			"update-budget":                             {"AccountId", "NewBudget"},
			"update-budget-action":                      {"AccountId", "ActionId", "ActionThreshold", "ApprovalModel", "BudgetName", "Definition", "ExecutionRoleArn", "NotificationType", "Subscribers"},
			"update-notification":                       {"AccountId", "BudgetName", "NewNotification", "OldNotification"},
			"update-subscriber":                         {"AccountId", "BudgetName", "NewSubscriber", "Notification", "OldSubscriber"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-budget":                             {"AccountId": "*string", "Budget": "*types.Budget", "NotificationsWithSubscribers": "[]types.NotificationWithSubscribers", "ResourceTags": "[]types.ResourceTag"},
			"create-budget-action":                      {"AccountId": "*string", "ActionThreshold": "*types.ActionThreshold", "ActionType": "types.ActionType", "ApprovalModel": "types.ApprovalModel", "BudgetName": "*string", "Definition": "*types.Definition", "ExecutionRoleArn": "*string", "NotificationType": "types.NotificationType", "ResourceTags": "[]types.ResourceTag", "Subscribers": "[]types.Subscriber"},
			"create-notification":                       {"AccountId": "*string", "BudgetName": "*string", "Notification": "*types.Notification", "Subscribers": "[]types.Subscriber"},
			"create-subscriber":                         {"AccountId": "*string", "BudgetName": "*string", "Notification": "*types.Notification", "Subscriber": "*types.Subscriber"},
			"delete-budget":                             {"AccountId": "*string", "BudgetName": "*string"},
			"delete-budget-action":                      {"AccountId": "*string", "ActionId": "*string", "BudgetName": "*string"},
			"delete-notification":                       {"AccountId": "*string", "BudgetName": "*string", "Notification": "*types.Notification"},
			"delete-subscriber":                         {"AccountId": "*string", "BudgetName": "*string", "Notification": "*types.Notification", "Subscriber": "*types.Subscriber"},
			"describe-budget":                           {"AccountId": "*string", "BudgetName": "*string", "ShowFilterExpression": "*bool"},
			"describe-budget-action":                    {"AccountId": "*string", "ActionId": "*string", "BudgetName": "*string"},
			"describe-budget-action-histories":          {"AccountId": "*string", "ActionId": "*string", "BudgetName": "*string", "MaxResults": "*int32", "NextToken": "*string", "TimePeriod": "*types.TimePeriod"},
			"describe-budget-actions-for-account":       {"AccountId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-budget-actions-for-budget":        {"AccountId": "*string", "BudgetName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-budget-notifications-for-account": {"AccountId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-budget-performance-history":       {"AccountId": "*string", "BudgetName": "*string", "MaxResults": "*int32", "NextToken": "*string", "TimePeriod": "*types.TimePeriod"},
			"describe-budgets":                          {"AccountId": "*string", "MaxResults": "*int32", "NextToken": "*string", "ShowFilterExpression": "*bool"},
			"describe-notifications-for-budget":         {"AccountId": "*string", "BudgetName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-subscribers-for-notification":     {"AccountId": "*string", "BudgetName": "*string", "MaxResults": "*int32", "NextToken": "*string", "Notification": "*types.Notification"},
			"execute-budget-action":                     {"AccountId": "*string", "ActionId": "*string", "BudgetName": "*string", "ExecutionType": "types.ExecutionType"},
			"list-tags-for-resource":                    {"ResourceARN": "*string"},
			"tag-resource":                              {"ResourceARN": "*string", "ResourceTags": "[]types.ResourceTag"},
			"untag-resource":                            {"ResourceARN": "*string", "ResourceTagKeys": "[]string"},
			"update-budget":                             {"AccountId": "*string", "NewBudget": "*types.Budget"},
			"update-budget-action":                      {"AccountId": "*string", "ActionId": "*string", "ActionThreshold": "*types.ActionThreshold", "ApprovalModel": "types.ApprovalModel", "BudgetName": "*string", "Definition": "*types.Definition", "ExecutionRoleArn": "*string", "NotificationType": "types.NotificationType", "Subscribers": "[]types.Subscriber"},
			"update-notification":                       {"AccountId": "*string", "BudgetName": "*string", "NewNotification": "*types.Notification", "OldNotification": "*types.Notification"},
			"update-subscriber":                         {"AccountId": "*string", "BudgetName": "*string", "NewSubscriber": "*types.Subscriber", "Notification": "*types.Notification", "OldSubscriber": "*types.Subscriber"},
		},
		OperationInputRequired: map[string][]string{
			"create-budget":                             {"AccountId", "Budget"},
			"create-budget-action":                      {"AccountId", "ActionThreshold", "ActionType", "ApprovalModel", "BudgetName", "Definition", "ExecutionRoleArn", "NotificationType", "Subscribers"},
			"create-notification":                       {"AccountId", "BudgetName", "Notification", "Subscribers"},
			"create-subscriber":                         {"AccountId", "BudgetName", "Notification", "Subscriber"},
			"delete-budget":                             {"AccountId", "BudgetName"},
			"delete-budget-action":                      {"AccountId", "ActionId", "BudgetName"},
			"delete-notification":                       {"AccountId", "BudgetName", "Notification"},
			"delete-subscriber":                         {"AccountId", "BudgetName", "Notification", "Subscriber"},
			"describe-budget":                           {"AccountId", "BudgetName"},
			"describe-budget-action":                    {"AccountId", "ActionId", "BudgetName"},
			"describe-budget-action-histories":          {"AccountId", "ActionId", "BudgetName"},
			"describe-budget-actions-for-account":       {"AccountId"},
			"describe-budget-actions-for-budget":        {"AccountId", "BudgetName"},
			"describe-budget-notifications-for-account": {"AccountId"},
			"describe-budget-performance-history":       {"AccountId", "BudgetName"},
			"describe-budgets":                          {"AccountId"},
			"describe-notifications-for-budget":         {"AccountId", "BudgetName"},
			"describe-subscribers-for-notification":     {"AccountId", "BudgetName", "Notification"},
			"execute-budget-action":                     {"AccountId", "ActionId", "BudgetName", "ExecutionType"},
			"list-tags-for-resource":                    {"ResourceARN"},
			"tag-resource":                              {"ResourceARN", "ResourceTags"},
			"untag-resource":                            {"ResourceARN", "ResourceTagKeys"},
			"update-budget":                             {"AccountId", "NewBudget"},
			"update-budget-action":                      {"AccountId", "ActionId", "BudgetName"},
			"update-notification":                       {"AccountId", "BudgetName", "NewNotification", "OldNotification"},
			"update-subscriber":                         {"AccountId", "BudgetName", "NewSubscriber", "Notification", "OldSubscriber"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("budgets", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
