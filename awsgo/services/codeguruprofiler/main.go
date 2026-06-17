package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/codeguruprofiler/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-notification-channels", "batch-get-frame-metric-data", "configure-agent", "create-profiling-group", "delete-profiling-group", "describe-profiling-group", "get-findings-report-account-summary", "get-notification-configuration", "get-policy", "get-profile", "get-recommendations", "list-findings-reports", "list-profile-times", "list-profiling-groups", "list-tags-for-resource", "post-agent-profile", "put-permission", "remove-notification-channel", "remove-permission", "submit-feedback", "tag-resource", "untag-resource", "update-profiling-group"},
		OperationSet: map[string]bool{"add-notification-channels": true, "batch-get-frame-metric-data": true, "configure-agent": true, "create-profiling-group": true, "delete-profiling-group": true, "describe-profiling-group": true, "get-findings-report-account-summary": true, "get-notification-configuration": true, "get-policy": true, "get-profile": true, "get-recommendations": true, "list-findings-reports": true, "list-profile-times": true, "list-profiling-groups": true, "list-tags-for-resource": true, "post-agent-profile": true, "put-permission": true, "remove-notification-channel": true, "remove-permission": true, "submit-feedback": true, "tag-resource": true, "untag-resource": true, "update-profiling-group": true},
		OperationInputs: map[string][]string{
			"add-notification-channels":           {"Channels", "ProfilingGroupName"},
			"batch-get-frame-metric-data":         {"EndTime", "FrameMetrics", "Period", "ProfilingGroupName", "StartTime", "TargetResolution"},
			"configure-agent":                     {"FleetInstanceId", "Metadata", "ProfilingGroupName"},
			"create-profiling-group":              {"AgentOrchestrationConfig", "ClientToken", "ComputePlatform", "ProfilingGroupName", "Tags"},
			"delete-profiling-group":              {"ProfilingGroupName"},
			"describe-profiling-group":            {"ProfilingGroupName"},
			"get-findings-report-account-summary": {"DailyReportsOnly", "MaxResults", "NextToken"},
			"get-notification-configuration":      {"ProfilingGroupName"},
			"get-policy":                          {"ProfilingGroupName"},
			"get-profile":                         {"Accept", "EndTime", "MaxDepth", "Period", "ProfilingGroupName", "StartTime"},
			"get-recommendations":                 {"EndTime", "Locale", "ProfilingGroupName", "StartTime"},
			"list-findings-reports":               {"DailyReportsOnly", "EndTime", "MaxResults", "NextToken", "ProfilingGroupName", "StartTime"},
			"list-profile-times":                  {"EndTime", "MaxResults", "NextToken", "OrderBy", "Period", "ProfilingGroupName", "StartTime"},
			"list-profiling-groups":               {"IncludeDescription", "MaxResults", "NextToken"},
			"list-tags-for-resource":              {"ResourceArn"},
			"post-agent-profile":                  {"AgentProfile", "ContentType", "ProfileToken", "ProfilingGroupName"},
			"put-permission":                      {"ActionGroup", "Principals", "ProfilingGroupName", "RevisionId"},
			"remove-notification-channel":         {"ChannelId", "ProfilingGroupName"},
			"remove-permission":                   {"ActionGroup", "ProfilingGroupName", "RevisionId"},
			"submit-feedback":                     {"AnomalyInstanceId", "Comment", "ProfilingGroupName", "Type"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-profiling-group":              {"AgentOrchestrationConfig", "ProfilingGroupName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-notification-channels":           {"Channels": "[]types.Channel", "ProfilingGroupName": "*string"},
			"batch-get-frame-metric-data":         {"EndTime": "*time.Time", "FrameMetrics": "[]types.FrameMetric", "Period": "*string", "ProfilingGroupName": "*string", "StartTime": "*time.Time", "TargetResolution": "types.AggregationPeriod"},
			"configure-agent":                     {"FleetInstanceId": "*string", "Metadata": "map[string]string", "ProfilingGroupName": "*string"},
			"create-profiling-group":              {"AgentOrchestrationConfig": "*types.AgentOrchestrationConfig", "ClientToken": "*string", "ComputePlatform": "types.ComputePlatform", "ProfilingGroupName": "*string", "Tags": "map[string]string"},
			"delete-profiling-group":              {"ProfilingGroupName": "*string"},
			"describe-profiling-group":            {"ProfilingGroupName": "*string"},
			"get-findings-report-account-summary": {"DailyReportsOnly": "*bool", "MaxResults": "*int32", "NextToken": "*string"},
			"get-notification-configuration":      {"ProfilingGroupName": "*string"},
			"get-policy":                          {"ProfilingGroupName": "*string"},
			"get-profile":                         {"Accept": "*string", "EndTime": "*time.Time", "MaxDepth": "*int32", "Period": "*string", "ProfilingGroupName": "*string", "StartTime": "*time.Time"},
			"get-recommendations":                 {"EndTime": "*time.Time", "Locale": "*string", "ProfilingGroupName": "*string", "StartTime": "*time.Time"},
			"list-findings-reports":               {"DailyReportsOnly": "*bool", "EndTime": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "ProfilingGroupName": "*string", "StartTime": "*time.Time"},
			"list-profile-times":                  {"EndTime": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "OrderBy": "types.OrderBy", "Period": "types.AggregationPeriod", "ProfilingGroupName": "*string", "StartTime": "*time.Time"},
			"list-profiling-groups":               {"IncludeDescription": "*bool", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":              {"ResourceArn": "*string"},
			"post-agent-profile":                  {"AgentProfile": "[]byte", "ContentType": "*string", "ProfileToken": "*string", "ProfilingGroupName": "*string"},
			"put-permission":                      {"ActionGroup": "types.ActionGroup", "Principals": "[]string", "ProfilingGroupName": "*string", "RevisionId": "*string"},
			"remove-notification-channel":         {"ChannelId": "*string", "ProfilingGroupName": "*string"},
			"remove-permission":                   {"ActionGroup": "types.ActionGroup", "ProfilingGroupName": "*string", "RevisionId": "*string"},
			"submit-feedback":                     {"AnomalyInstanceId": "*string", "Comment": "*string", "ProfilingGroupName": "*string", "Type": "types.FeedbackType"},
			"tag-resource":                        {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                      {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-profiling-group":              {"AgentOrchestrationConfig": "*types.AgentOrchestrationConfig", "ProfilingGroupName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"add-notification-channels":           {"Channels", "ProfilingGroupName"},
			"batch-get-frame-metric-data":         {"ProfilingGroupName"},
			"configure-agent":                     {"ProfilingGroupName"},
			"create-profiling-group":              {"ClientToken", "ProfilingGroupName"},
			"delete-profiling-group":              {"ProfilingGroupName"},
			"describe-profiling-group":            {"ProfilingGroupName"},
			"get-findings-report-account-summary": {},
			"get-notification-configuration":      {"ProfilingGroupName"},
			"get-policy":                          {"ProfilingGroupName"},
			"get-profile":                         {"ProfilingGroupName"},
			"get-recommendations":                 {"EndTime", "ProfilingGroupName", "StartTime"},
			"list-findings-reports":               {"EndTime", "ProfilingGroupName", "StartTime"},
			"list-profile-times":                  {"EndTime", "Period", "ProfilingGroupName", "StartTime"},
			"list-profiling-groups":               {},
			"list-tags-for-resource":              {"ResourceArn"},
			"post-agent-profile":                  {"AgentProfile", "ContentType", "ProfilingGroupName"},
			"put-permission":                      {"ActionGroup", "Principals", "ProfilingGroupName"},
			"remove-notification-channel":         {"ChannelId", "ProfilingGroupName"},
			"remove-permission":                   {"ActionGroup", "ProfilingGroupName", "RevisionId"},
			"submit-feedback":                     {"AnomalyInstanceId", "ProfilingGroupName", "Type"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-profiling-group":              {"AgentOrchestrationConfig", "ProfilingGroupName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("codeguruprofiler", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
