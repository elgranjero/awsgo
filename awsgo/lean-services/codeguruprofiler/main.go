package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
)

var fields_add_notification_channels = []leanruntime.Field{
	{Name: "Channels", Flag: "channels", Type: "[]types.Channel", Required: true},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_batch_get_frame_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "FrameMetrics", Flag: "frame-metrics", Type: "[]types.FrameMetric", Required: false},
	{Name: "Period", Flag: "period", Type: "*string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TargetResolution", Flag: "target-resolution", Type: "types.AggregationPeriod", Required: false},
}

var fields_configure_agent = []leanruntime.Field{
	{Name: "FleetInstanceId", Flag: "fleet-instance-id", Type: "*string", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_create_profiling_group = []leanruntime.Field{
	{Name: "AgentOrchestrationConfig", Flag: "agent-orchestration-config", Type: "*types.AgentOrchestrationConfig", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ComputePlatform", Flag: "compute-platform", Type: "types.ComputePlatform", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_profiling_group = []leanruntime.Field{
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_describe_profiling_group = []leanruntime.Field{
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_get_findings_report_account_summary = []leanruntime.Field{
	{Name: "DailyReportsOnly", Flag: "daily-reports-only", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_notification_configuration = []leanruntime.Field{
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_get_profile = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxDepth", Flag: "max-depth", Type: "*int32", Required: false},
	{Name: "Period", Flag: "period", Type: "*string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_recommendations = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_findings_reports = []leanruntime.Field{
	{Name: "DailyReportsOnly", Flag: "daily-reports-only", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_profile_times = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "types.OrderBy", Required: false},
	{Name: "Period", Flag: "period", Type: "types.AggregationPeriod", Required: true},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_profiling_groups = []leanruntime.Field{
	{Name: "IncludeDescription", Flag: "include-description", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_post_agent_profile = []leanruntime.Field{
	{Name: "AgentProfile", Flag: "agent-profile", Type: "[]byte", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
	{Name: "ProfileToken", Flag: "profile-token", Type: "*string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_put_permission = []leanruntime.Field{
	{Name: "ActionGroup", Flag: "action-group", Type: "types.ActionGroup", Required: true},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: true},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
}

var fields_remove_notification_channel = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

var fields_remove_permission = []leanruntime.Field{
	{Name: "ActionGroup", Flag: "action-group", Type: "types.ActionGroup", Required: true},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_submit_feedback = []leanruntime.Field{
	{Name: "AnomalyInstanceId", Flag: "anomaly-instance-id", Type: "*string", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.FeedbackType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_profiling_group = []leanruntime.Field{
	{Name: "AgentOrchestrationConfig", Flag: "agent-orchestration-config", Type: "*types.AgentOrchestrationConfig", Required: true},
	{Name: "ProfilingGroupName", Flag: "profiling-group-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-notification-channels": {
			Name:   "add-notification-channels",
			Fields: fields_add_notification_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddNotificationChannelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_notification_channels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddNotificationChannels(ctx, input)
			},
		},
		"batch-get-frame-metric-data": {
			Name:   "batch-get-frame-metric-data",
			Fields: fields_batch_get_frame_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFrameMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_frame_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetFrameMetricData(ctx, input)
			},
		},
		"configure-agent": {
			Name:   "configure-agent",
			Fields: fields_configure_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfigureAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_configure_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfigureAgent(ctx, input)
			},
		},
		"create-profiling-group": {
			Name:   "create-profiling-group",
			Fields: fields_create_profiling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfilingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profiling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfilingGroup(ctx, input)
			},
		},
		"delete-profiling-group": {
			Name:   "delete-profiling-group",
			Fields: fields_delete_profiling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfilingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profiling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfilingGroup(ctx, input)
			},
		},
		"describe-profiling-group": {
			Name:   "describe-profiling-group",
			Fields: fields_describe_profiling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProfilingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_profiling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProfilingGroup(ctx, input)
			},
		},
		"get-findings-report-account-summary": {
			Name:   "get-findings-report-account-summary",
			Fields: fields_get_findings_report_account_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsReportAccountSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_findings_report_account_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindingsReportAccountSummary(ctx, input)
				}
				var results []*svc.GetFindingsReportAccountSummaryOutput
				p := svc.NewGetFindingsReportAccountSummaryPaginator(client, input)
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
		"get-notification-configuration": {
			Name:   "get-notification-configuration",
			Fields: fields_get_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotificationConfiguration(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-profile": {
			Name:   "get-profile",
			Fields: fields_get_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfile(ctx, input)
			},
		},
		"get-recommendations": {
			Name:   "get-recommendations",
			Fields: fields_get_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommendations(ctx, input)
			},
		},
		"list-findings-reports": {
			Name:   "list-findings-reports",
			Fields: fields_list_findings_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingsReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_findings_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindingsReports(ctx, input)
				}
				var results []*svc.ListFindingsReportsOutput
				p := svc.NewListFindingsReportsPaginator(client, input)
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
		"list-profile-times": {
			Name:   "list-profile-times",
			Fields: fields_list_profile_times,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileTimesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profile_times, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfileTimes(ctx, input)
				}
				var results []*svc.ListProfileTimesOutput
				p := svc.NewListProfileTimesPaginator(client, input)
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
		"list-profiling-groups": {
			Name:   "list-profiling-groups",
			Fields: fields_list_profiling_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfilingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profiling_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfilingGroups(ctx, input)
				}
				var results []*svc.ListProfilingGroupsOutput
				p := svc.NewListProfilingGroupsPaginator(client, input)
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
		"post-agent-profile": {
			Name:   "post-agent-profile",
			Fields: fields_post_agent_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostAgentProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_agent_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostAgentProfile(ctx, input)
			},
		},
		"put-permission": {
			Name:   "put-permission",
			Fields: fields_put_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPermission(ctx, input)
			},
		},
		"remove-notification-channel": {
			Name:   "remove-notification-channel",
			Fields: fields_remove_notification_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveNotificationChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_notification_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveNotificationChannel(ctx, input)
			},
		},
		"remove-permission": {
			Name:   "remove-permission",
			Fields: fields_remove_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemovePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemovePermission(ctx, input)
			},
		},
		"submit-feedback": {
			Name:   "submit-feedback",
			Fields: fields_submit_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitFeedback(ctx, input)
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
		"update-profiling-group": {
			Name:   "update-profiling-group",
			Fields: fields_update_profiling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfilingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profiling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfilingGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codeguruprofiler", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
