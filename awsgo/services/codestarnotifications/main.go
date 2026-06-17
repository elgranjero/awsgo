package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/codestarnotifications/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-notification-rule", "delete-notification-rule", "delete-target", "describe-notification-rule", "list-event-types", "list-notification-rules", "list-tags-for-resource", "list-targets", "subscribe", "tag-resource", "unsubscribe", "untag-resource", "update-notification-rule"},
		OperationSet: map[string]bool{"create-notification-rule": true, "delete-notification-rule": true, "delete-target": true, "describe-notification-rule": true, "list-event-types": true, "list-notification-rules": true, "list-tags-for-resource": true, "list-targets": true, "subscribe": true, "tag-resource": true, "unsubscribe": true, "untag-resource": true, "update-notification-rule": true},
		OperationInputs: map[string][]string{
			"create-notification-rule":   {"ClientRequestToken", "DetailType", "EventTypeIds", "Name", "Resource", "Status", "Tags", "Targets"},
			"delete-notification-rule":   {"Arn"},
			"delete-target":              {"ForceUnsubscribeAll", "TargetAddress"},
			"describe-notification-rule": {"Arn"},
			"list-event-types":           {"Filters", "MaxResults", "NextToken"},
			"list-notification-rules":    {"Filters", "MaxResults", "NextToken"},
			"list-tags-for-resource":     {"Arn"},
			"list-targets":               {"Filters", "MaxResults", "NextToken"},
			"subscribe":                  {"Arn", "ClientRequestToken", "Target"},
			"tag-resource":               {"Arn", "Tags"},
			"unsubscribe":                {"Arn", "TargetAddress"},
			"untag-resource":             {"Arn", "TagKeys"},
			"update-notification-rule":   {"Arn", "DetailType", "EventTypeIds", "Name", "Status", "Targets"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-notification-rule":   {"ClientRequestToken": "*string", "DetailType": "types.DetailType", "EventTypeIds": "[]string", "Name": "*string", "Resource": "*string", "Status": "types.NotificationRuleStatus", "Tags": "map[string]string", "Targets": "[]types.Target"},
			"delete-notification-rule":   {"Arn": "*string"},
			"delete-target":              {"ForceUnsubscribeAll": "bool", "TargetAddress": "*string"},
			"describe-notification-rule": {"Arn": "*string"},
			"list-event-types":           {"Filters": "[]types.ListEventTypesFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-notification-rules":    {"Filters": "[]types.ListNotificationRulesFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":     {"Arn": "*string"},
			"list-targets":               {"Filters": "[]types.ListTargetsFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"subscribe":                  {"Arn": "*string", "ClientRequestToken": "*string", "Target": "*types.Target"},
			"tag-resource":               {"Arn": "*string", "Tags": "map[string]string"},
			"unsubscribe":                {"Arn": "*string", "TargetAddress": "*string"},
			"untag-resource":             {"Arn": "*string", "TagKeys": "[]string"},
			"update-notification-rule":   {"Arn": "*string", "DetailType": "types.DetailType", "EventTypeIds": "[]string", "Name": "*string", "Status": "types.NotificationRuleStatus", "Targets": "[]types.Target"},
		},
		OperationInputRequired: map[string][]string{
			"create-notification-rule":   {"DetailType", "EventTypeIds", "Name", "Resource", "Targets"},
			"delete-notification-rule":   {"Arn"},
			"delete-target":              {"TargetAddress"},
			"describe-notification-rule": {"Arn"},
			"list-event-types":           {},
			"list-notification-rules":    {},
			"list-tags-for-resource":     {"Arn"},
			"list-targets":               {},
			"subscribe":                  {"Arn", "Target"},
			"tag-resource":               {"Arn", "Tags"},
			"unsubscribe":                {"Arn", "TargetAddress"},
			"untag-resource":             {"Arn", "TagKeys"},
			"update-notification-rule":   {"Arn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("codestarnotifications", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
