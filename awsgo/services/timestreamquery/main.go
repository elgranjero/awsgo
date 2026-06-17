package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/timestreamquery/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-query", "create-scheduled-query", "delete-scheduled-query", "describe-account-settings", "describe-endpoints", "describe-scheduled-query", "execute-scheduled-query", "list-scheduled-queries", "list-tags-for-resource", "prepare-query", "query", "tag-resource", "untag-resource", "update-account-settings", "update-scheduled-query"},
		OperationSet: map[string]bool{"cancel-query": true, "create-scheduled-query": true, "delete-scheduled-query": true, "describe-account-settings": true, "describe-endpoints": true, "describe-scheduled-query": true, "execute-scheduled-query": true, "list-scheduled-queries": true, "list-tags-for-resource": true, "prepare-query": true, "query": true, "tag-resource": true, "untag-resource": true, "update-account-settings": true, "update-scheduled-query": true},
		OperationInputs: map[string][]string{
			"cancel-query":              {"QueryId"},
			"create-scheduled-query":    {"ClientToken", "ErrorReportConfiguration", "KmsKeyId", "Name", "NotificationConfiguration", "QueryString", "ScheduleConfiguration", "ScheduledQueryExecutionRoleArn", "Tags", "TargetConfiguration"},
			"delete-scheduled-query":    {"ScheduledQueryArn"},
			"describe-account-settings": {},
			"describe-endpoints":        {},
			"describe-scheduled-query":  {"ScheduledQueryArn"},
			"execute-scheduled-query":   {"ClientToken", "InvocationTime", "QueryInsights", "ScheduledQueryArn"},
			"list-scheduled-queries":    {"MaxResults", "NextToken"},
			"list-tags-for-resource":    {"MaxResults", "NextToken", "ResourceARN"},
			"prepare-query":             {"QueryString", "ValidateOnly"},
			"query":                     {"ClientToken", "MaxRows", "NextToken", "QueryInsights", "QueryString"},
			"tag-resource":              {"ResourceARN", "Tags"},
			"untag-resource":            {"ResourceARN", "TagKeys"},
			"update-account-settings":   {"MaxQueryTCU", "QueryCompute", "QueryPricingModel"},
			"update-scheduled-query":    {"ScheduledQueryArn", "State"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-query":              {"QueryId": "*string"},
			"create-scheduled-query":    {"ClientToken": "*string", "ErrorReportConfiguration": "*types.ErrorReportConfiguration", "KmsKeyId": "*string", "Name": "*string", "NotificationConfiguration": "*types.NotificationConfiguration", "QueryString": "*string", "ScheduleConfiguration": "*types.ScheduleConfiguration", "ScheduledQueryExecutionRoleArn": "*string", "Tags": "[]types.Tag", "TargetConfiguration": "*types.TargetConfiguration"},
			"delete-scheduled-query":    {"ScheduledQueryArn": "*string"},
			"describe-account-settings": {},
			"describe-endpoints":        {},
			"describe-scheduled-query":  {"ScheduledQueryArn": "*string"},
			"execute-scheduled-query":   {"ClientToken": "*string", "InvocationTime": "*time.Time", "QueryInsights": "*types.ScheduledQueryInsights", "ScheduledQueryArn": "*string"},
			"list-scheduled-queries":    {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":    {"MaxResults": "*int32", "NextToken": "*string", "ResourceARN": "*string"},
			"prepare-query":             {"QueryString": "*string", "ValidateOnly": "*bool"},
			"query":                     {"ClientToken": "*string", "MaxRows": "*int32", "NextToken": "*string", "QueryInsights": "*types.QueryInsights", "QueryString": "*string"},
			"tag-resource":              {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":            {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-account-settings":   {"MaxQueryTCU": "*int32", "QueryCompute": "*types.QueryComputeRequest", "QueryPricingModel": "types.QueryPricingModel"},
			"update-scheduled-query":    {"ScheduledQueryArn": "*string", "State": "types.ScheduledQueryState"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-query":              {"QueryId"},
			"create-scheduled-query":    {"ErrorReportConfiguration", "Name", "NotificationConfiguration", "QueryString", "ScheduleConfiguration", "ScheduledQueryExecutionRoleArn"},
			"delete-scheduled-query":    {"ScheduledQueryArn"},
			"describe-account-settings": {},
			"describe-endpoints":        {},
			"describe-scheduled-query":  {"ScheduledQueryArn"},
			"execute-scheduled-query":   {"InvocationTime", "ScheduledQueryArn"},
			"list-scheduled-queries":    {},
			"list-tags-for-resource":    {"ResourceARN"},
			"prepare-query":             {"QueryString"},
			"query":                     {"QueryString"},
			"tag-resource":              {"ResourceARN", "Tags"},
			"untag-resource":            {"ResourceARN", "TagKeys"},
			"update-account-settings":   {},
			"update-scheduled-query":    {"ScheduledQueryArn", "State"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("timestreamquery", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
