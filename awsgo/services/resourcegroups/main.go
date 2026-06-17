package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/resourcegroups/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-tag-sync-task", "create-group", "delete-group", "get-account-settings", "get-group", "get-group-configuration", "get-group-query", "get-tag-sync-task", "get-tags", "group-resources", "list-group-resources", "list-grouping-statuses", "list-groups", "list-tag-sync-tasks", "put-group-configuration", "search-resources", "start-tag-sync-task", "tag", "ungroup-resources", "untag", "update-account-settings", "update-group", "update-group-query"},
		OperationSet: map[string]bool{"cancel-tag-sync-task": true, "create-group": true, "delete-group": true, "get-account-settings": true, "get-group": true, "get-group-configuration": true, "get-group-query": true, "get-tag-sync-task": true, "get-tags": true, "group-resources": true, "list-group-resources": true, "list-grouping-statuses": true, "list-groups": true, "list-tag-sync-tasks": true, "put-group-configuration": true, "search-resources": true, "start-tag-sync-task": true, "tag": true, "ungroup-resources": true, "untag": true, "update-account-settings": true, "update-group": true, "update-group-query": true},
		OperationInputs: map[string][]string{
			"cancel-tag-sync-task":    {"TaskArn"},
			"create-group":            {"Configuration", "Criticality", "Description", "DisplayName", "Name", "Owner", "ResourceQuery", "Tags"},
			"delete-group":            {"Group", "GroupName"},
			"get-account-settings":    {},
			"get-group":               {"Group", "GroupName"},
			"get-group-configuration": {"Group"},
			"get-group-query":         {"Group", "GroupName"},
			"get-tag-sync-task":       {"TaskArn"},
			"get-tags":                {"Arn"},
			"group-resources":         {"Group", "ResourceArns"},
			"list-group-resources":    {"Filters", "Group", "GroupName", "MaxResults", "NextToken"},
			"list-grouping-statuses":  {"Filters", "Group", "MaxResults", "NextToken"},
			"list-groups":             {"Filters", "MaxResults", "NextToken"},
			"list-tag-sync-tasks":     {"Filters", "MaxResults", "NextToken"},
			"put-group-configuration": {"Configuration", "Group"},
			"search-resources":        {"MaxResults", "NextToken", "ResourceQuery"},
			"start-tag-sync-task":     {"Group", "ResourceQuery", "RoleArn", "TagKey", "TagValue"},
			"tag":                     {"Arn", "Tags"},
			"ungroup-resources":       {"Group", "ResourceArns"},
			"untag":                   {"Arn", "Keys"},
			"update-account-settings": {"GroupLifecycleEventsDesiredStatus"},
			"update-group":            {"Criticality", "Description", "DisplayName", "Group", "GroupName", "Owner"},
			"update-group-query":      {"Group", "GroupName", "ResourceQuery"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-tag-sync-task":    {"TaskArn": "*string"},
			"create-group":            {"Configuration": "[]types.GroupConfigurationItem", "Criticality": "*int32", "Description": "*string", "DisplayName": "*string", "Name": "*string", "Owner": "*string", "ResourceQuery": "*types.ResourceQuery", "Tags": "map[string]string"},
			"delete-group":            {"Group": "*string", "GroupName": "*string"},
			"get-account-settings":    {},
			"get-group":               {"Group": "*string", "GroupName": "*string"},
			"get-group-configuration": {"Group": "*string"},
			"get-group-query":         {"Group": "*string", "GroupName": "*string"},
			"get-tag-sync-task":       {"TaskArn": "*string"},
			"get-tags":                {"Arn": "*string"},
			"group-resources":         {"Group": "*string", "ResourceArns": "[]string"},
			"list-group-resources":    {"Filters": "[]types.ResourceFilter", "Group": "*string", "GroupName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-grouping-statuses":  {"Filters": "[]types.ListGroupingStatusesFilter", "Group": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-groups":             {"Filters": "[]types.GroupFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tag-sync-tasks":     {"Filters": "[]types.ListTagSyncTasksFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"put-group-configuration": {"Configuration": "[]types.GroupConfigurationItem", "Group": "*string"},
			"search-resources":        {"MaxResults": "*int32", "NextToken": "*string", "ResourceQuery": "*types.ResourceQuery"},
			"start-tag-sync-task":     {"Group": "*string", "ResourceQuery": "*types.ResourceQuery", "RoleArn": "*string", "TagKey": "*string", "TagValue": "*string"},
			"tag":                     {"Arn": "*string", "Tags": "map[string]string"},
			"ungroup-resources":       {"Group": "*string", "ResourceArns": "[]string"},
			"untag":                   {"Arn": "*string", "Keys": "[]string"},
			"update-account-settings": {"GroupLifecycleEventsDesiredStatus": "types.GroupLifecycleEventsDesiredStatus"},
			"update-group":            {"Criticality": "*int32", "Description": "*string", "DisplayName": "*string", "Group": "*string", "GroupName": "*string", "Owner": "*string"},
			"update-group-query":      {"Group": "*string", "GroupName": "*string", "ResourceQuery": "*types.ResourceQuery"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-tag-sync-task":    {"TaskArn"},
			"create-group":            {"Name"},
			"delete-group":            {},
			"get-account-settings":    {},
			"get-group":               {},
			"get-group-configuration": {},
			"get-group-query":         {},
			"get-tag-sync-task":       {"TaskArn"},
			"get-tags":                {"Arn"},
			"group-resources":         {"Group", "ResourceArns"},
			"list-group-resources":    {},
			"list-grouping-statuses":  {"Group"},
			"list-groups":             {},
			"list-tag-sync-tasks":     {},
			"put-group-configuration": {},
			"search-resources":        {"ResourceQuery"},
			"start-tag-sync-task":     {"Group", "RoleArn"},
			"tag":                     {"Arn", "Tags"},
			"ungroup-resources":       {"Group", "ResourceArns"},
			"untag":                   {"Arn", "Keys"},
			"update-account-settings": {},
			"update-group":            {},
			"update-group-query":      {"ResourceQuery"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("resourcegroups", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
