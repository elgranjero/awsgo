package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/computeoptimizerautomation/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-accounts", "create-automation-rule", "delete-automation-rule", "disassociate-accounts", "get-automation-event", "get-automation-rule", "get-enrollment-configuration", "list-accounts", "list-automation-event-steps", "list-automation-event-summaries", "list-automation-events", "list-automation-rule-preview", "list-automation-rule-preview-summaries", "list-automation-rules", "list-recommended-action-summaries", "list-recommended-actions", "list-tags-for-resource", "rollback-automation-event", "start-automation-event", "tag-resource", "untag-resource", "update-automation-rule", "update-enrollment-configuration"},
		OperationSet: map[string]bool{"associate-accounts": true, "create-automation-rule": true, "delete-automation-rule": true, "disassociate-accounts": true, "get-automation-event": true, "get-automation-rule": true, "get-enrollment-configuration": true, "list-accounts": true, "list-automation-event-steps": true, "list-automation-event-summaries": true, "list-automation-events": true, "list-automation-rule-preview": true, "list-automation-rule-preview-summaries": true, "list-automation-rules": true, "list-recommended-action-summaries": true, "list-recommended-actions": true, "list-tags-for-resource": true, "rollback-automation-event": true, "start-automation-event": true, "tag-resource": true, "untag-resource": true, "update-automation-rule": true, "update-enrollment-configuration": true},
		OperationInputs: map[string][]string{
			"associate-accounts":                     {"AccountIds", "ClientToken"},
			"create-automation-rule":                 {"ClientToken", "Criteria", "Description", "Name", "OrganizationConfiguration", "Priority", "RecommendedActionTypes", "RuleType", "Schedule", "Status", "Tags"},
			"delete-automation-rule":                 {"ClientToken", "RuleArn", "RuleRevision"},
			"disassociate-accounts":                  {"AccountIds", "ClientToken"},
			"get-automation-event":                   {"EventId"},
			"get-automation-rule":                    {"RuleArn"},
			"get-enrollment-configuration":           {},
			"list-accounts":                          {"MaxResults", "NextToken"},
			"list-automation-event-steps":            {"EventId", "MaxResults", "NextToken"},
			"list-automation-event-summaries":        {"EndDateExclusive", "Filters", "MaxResults", "NextToken", "StartDateInclusive"},
			"list-automation-events":                 {"EndTimeExclusive", "Filters", "MaxResults", "NextToken", "StartTimeInclusive"},
			"list-automation-rule-preview":           {"Criteria", "MaxResults", "NextToken", "OrganizationScope", "RecommendedActionTypes", "RuleType"},
			"list-automation-rule-preview-summaries": {"Criteria", "MaxResults", "NextToken", "OrganizationScope", "RecommendedActionTypes", "RuleType"},
			"list-automation-rules":                  {"Filters", "MaxResults", "NextToken"},
			"list-recommended-action-summaries":      {"Filters", "MaxResults", "NextToken"},
			"list-recommended-actions":               {"Filters", "MaxResults", "NextToken"},
			"list-tags-for-resource":                 {"ResourceArn"},
			"rollback-automation-event":              {"ClientToken", "EventId"},
			"start-automation-event":                 {"ClientToken", "RecommendedActionId"},
			"tag-resource":                           {"ClientToken", "ResourceArn", "RuleRevision", "Tags"},
			"untag-resource":                         {"ClientToken", "ResourceArn", "RuleRevision", "TagKeys"},
			"update-automation-rule":                 {"ClientToken", "Criteria", "Description", "Name", "OrganizationConfiguration", "Priority", "RecommendedActionTypes", "RuleArn", "RuleRevision", "RuleType", "Schedule", "Status"},
			"update-enrollment-configuration":        {"ClientToken", "Status"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-accounts":                     {"AccountIds": "[]string", "ClientToken": "*string"},
			"create-automation-rule":                 {"ClientToken": "*string", "Criteria": "*types.Criteria", "Description": "*string", "Name": "*string", "OrganizationConfiguration": "*types.OrganizationConfiguration", "Priority": "*string", "RecommendedActionTypes": "[]types.RecommendedActionType", "RuleType": "types.RuleType", "Schedule": "*types.Schedule", "Status": "types.RuleStatus", "Tags": "[]types.Tag"},
			"delete-automation-rule":                 {"ClientToken": "*string", "RuleArn": "*string", "RuleRevision": "*int64"},
			"disassociate-accounts":                  {"AccountIds": "[]string", "ClientToken": "*string"},
			"get-automation-event":                   {"EventId": "*string"},
			"get-automation-rule":                    {"RuleArn": "*string"},
			"get-enrollment-configuration":           {},
			"list-accounts":                          {"MaxResults": "*int32", "NextToken": "*string"},
			"list-automation-event-steps":            {"EventId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-automation-event-summaries":        {"EndDateExclusive": "*string", "Filters": "[]types.AutomationEventFilter", "MaxResults": "*int32", "NextToken": "*string", "StartDateInclusive": "*string"},
			"list-automation-events":                 {"EndTimeExclusive": "*time.Time", "Filters": "[]types.AutomationEventFilter", "MaxResults": "*int32", "NextToken": "*string", "StartTimeInclusive": "*time.Time"},
			"list-automation-rule-preview":           {"Criteria": "*types.Criteria", "MaxResults": "*int32", "NextToken": "*string", "OrganizationScope": "*types.OrganizationScope", "RecommendedActionTypes": "[]types.RecommendedActionType", "RuleType": "types.RuleType"},
			"list-automation-rule-preview-summaries": {"Criteria": "*types.Criteria", "MaxResults": "*int32", "NextToken": "*string", "OrganizationScope": "*types.OrganizationScope", "RecommendedActionTypes": "[]types.RecommendedActionType", "RuleType": "types.RuleType"},
			"list-automation-rules":                  {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-recommended-action-summaries":      {"Filters": "[]types.RecommendedActionFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-recommended-actions":               {"Filters": "[]types.RecommendedActionFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                 {"ResourceArn": "*string"},
			"rollback-automation-event":              {"ClientToken": "*string", "EventId": "*string"},
			"start-automation-event":                 {"ClientToken": "*string", "RecommendedActionId": "*string"},
			"tag-resource":                           {"ClientToken": "*string", "ResourceArn": "*string", "RuleRevision": "*int64", "Tags": "[]types.Tag"},
			"untag-resource":                         {"ClientToken": "*string", "ResourceArn": "*string", "RuleRevision": "*int64", "TagKeys": "[]string"},
			"update-automation-rule":                 {"ClientToken": "*string", "Criteria": "*types.Criteria", "Description": "*string", "Name": "*string", "OrganizationConfiguration": "*types.OrganizationConfiguration", "Priority": "*string", "RecommendedActionTypes": "[]types.RecommendedActionType", "RuleArn": "*string", "RuleRevision": "*int64", "RuleType": "types.RuleType", "Schedule": "*types.Schedule", "Status": "types.RuleStatus"},
			"update-enrollment-configuration":        {"ClientToken": "*string", "Status": "types.EnrollmentStatus"},
		},
		OperationInputRequired: map[string][]string{
			"associate-accounts":                     {"AccountIds"},
			"create-automation-rule":                 {"Name", "RecommendedActionTypes", "RuleType", "Schedule", "Status"},
			"delete-automation-rule":                 {"RuleArn", "RuleRevision"},
			"disassociate-accounts":                  {"AccountIds"},
			"get-automation-event":                   {"EventId"},
			"get-automation-rule":                    {"RuleArn"},
			"get-enrollment-configuration":           {},
			"list-accounts":                          {},
			"list-automation-event-steps":            {"EventId"},
			"list-automation-event-summaries":        {},
			"list-automation-events":                 {},
			"list-automation-rule-preview":           {"RecommendedActionTypes", "RuleType"},
			"list-automation-rule-preview-summaries": {"RecommendedActionTypes", "RuleType"},
			"list-automation-rules":                  {},
			"list-recommended-action-summaries":      {},
			"list-recommended-actions":               {},
			"list-tags-for-resource":                 {"ResourceArn"},
			"rollback-automation-event":              {"EventId"},
			"start-automation-event":                 {"RecommendedActionId"},
			"tag-resource":                           {"ResourceArn", "RuleRevision", "Tags"},
			"untag-resource":                         {"ResourceArn", "RuleRevision", "TagKeys"},
			"update-automation-rule":                 {"RuleArn", "RuleRevision"},
			"update-enrollment-configuration":        {"Status"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("computeoptimizerautomation", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
