package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/xray"
)

var fields_batch_get_traces = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TraceIds", Flag: "trace-ids", Type: "[]string", Required: true},
}

var fields_cancel_trace_retrieval = []leanruntime.Field{
	{Name: "RetrievalToken", Flag: "retrieval-token", Type: "*string", Required: true},
}

var fields_create_group = []leanruntime.Field{
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "InsightsConfiguration", Flag: "insights-configuration", Type: "*types.InsightsConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_sampling_rule = []leanruntime.Field{
	{Name: "SamplingRule", Flag: "sampling-rule", Type: "*types.SamplingRule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupARN", Flag: "group-arn", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
}

var fields_delete_sampling_rule = []leanruntime.Field{
	{Name: "RuleARN", Flag: "rule-arn", Type: "*string", Required: false},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: false},
}

var fields_get_encryption_config = []leanruntime.Field{}

var fields_get_group = []leanruntime.Field{
	{Name: "GroupARN", Flag: "group-arn", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
}

var fields_get_groups = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_indexing_rules = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_insight = []leanruntime.Field{
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: true},
}

var fields_get_insight_events = []leanruntime.Field{
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_insight_impact_graph = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "InsightId", Flag: "insight-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_insight_summaries = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "GroupARN", Flag: "group-arn", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "States", Flag: "states", Type: "[]types.InsightState", Required: false},
}

var fields_get_retrieved_traces_graph = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RetrievalToken", Flag: "retrieval-token", Type: "*string", Required: true},
}

var fields_get_sampling_rules = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_sampling_statistic_summaries = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_sampling_targets = []leanruntime.Field{
	{Name: "SamplingBoostStatisticsDocuments", Flag: "sampling-boost-statistics-documents", Type: "[]types.SamplingBoostStatisticsDocument", Required: false},
	{Name: "SamplingStatisticsDocuments", Flag: "sampling-statistics-documents", Type: "[]types.SamplingStatisticsDocument", Required: true},
}

var fields_get_service_graph = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "GroupARN", Flag: "group-arn", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_time_series_service_statistics = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "EntitySelectorExpression", Flag: "entity-selector-expression", Type: "*string", Required: false},
	{Name: "ForecastStatistics", Flag: "forecast-statistics", Type: "*bool", Required: false},
	{Name: "GroupARN", Flag: "group-arn", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Period", Flag: "period", Type: "*int32", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_trace_graph = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TraceIds", Flag: "trace-ids", Type: "[]string", Required: true},
}

var fields_get_trace_segment_destination = []leanruntime.Field{}

var fields_get_trace_summaries = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sampling", Flag: "sampling", Type: "*bool", Required: false},
	{Name: "SamplingStrategy", Flag: "sampling-strategy", Type: "*types.SamplingStrategy", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "TimeRangeType", Flag: "time-range-type", Type: "types.TimeRangeType", Required: false},
}

var fields_list_resource_policies = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_retrieved_traces = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RetrievalToken", Flag: "retrieval-token", Type: "*string", Required: true},
	{Name: "TraceFormat", Flag: "trace-format", Type: "types.TraceFormatType", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_encryption_config = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.EncryptionType", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "BypassPolicyLockoutCheck", Flag: "bypass-policy-lockout-check", Type: "bool", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
}

var fields_put_telemetry_records = []leanruntime.Field{
	{Name: "EC2InstanceId", Flag: "ec2-instance-id", Type: "*string", Required: false},
	{Name: "Hostname", Flag: "hostname", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "TelemetryRecords", Flag: "telemetry-records", Type: "[]types.TelemetryRecord", Required: true},
}

var fields_put_trace_segments = []leanruntime.Field{
	{Name: "TraceSegmentDocuments", Flag: "trace-segment-documents", Type: "[]string", Required: true},
}

var fields_start_trace_retrieval = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "TraceIds", Flag: "trace-ids", Type: "[]string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_group = []leanruntime.Field{
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: false},
	{Name: "GroupARN", Flag: "group-arn", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "InsightsConfiguration", Flag: "insights-configuration", Type: "*types.InsightsConfiguration", Required: false},
}

var fields_update_indexing_rule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "types.IndexingRuleValueUpdate", Required: true},
}

var fields_update_sampling_rule = []leanruntime.Field{
	{Name: "SamplingRuleUpdate", Flag: "sampling-rule-update", Type: "*types.SamplingRuleUpdate", Required: true},
}

var fields_update_trace_segment_destination = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "types.TraceSegmentDestination", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-traces": {
			Name:   "batch-get-traces",
			Fields: fields_batch_get_traces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetTracesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_batch_get_traces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.BatchGetTraces(ctx, input)
				}
				var results []*svc.BatchGetTracesOutput
				p := svc.NewBatchGetTracesPaginator(client, input)
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
		"cancel-trace-retrieval": {
			Name:   "cancel-trace-retrieval",
			Fields: fields_cancel_trace_retrieval,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTraceRetrievalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_trace_retrieval, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTraceRetrieval(ctx, input)
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
		"create-sampling-rule": {
			Name:   "create-sampling-rule",
			Fields: fields_create_sampling_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSamplingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sampling_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSamplingRule(ctx, input)
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
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-sampling-rule": {
			Name:   "delete-sampling-rule",
			Fields: fields_delete_sampling_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSamplingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sampling_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSamplingRule(ctx, input)
			},
		},
		"get-encryption-config": {
			Name:   "get-encryption-config",
			Fields: fields_get_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEncryptionConfig(ctx, input)
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
		"get-groups": {
			Name:   "get-groups",
			Fields: fields_get_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetGroups(ctx, input)
				}
				var results []*svc.GetGroupsOutput
				p := svc.NewGetGroupsPaginator(client, input)
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
		"get-indexing-rules": {
			Name:   "get-indexing-rules",
			Fields: fields_get_indexing_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexingRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_indexing_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndexingRules(ctx, input)
			},
		},
		"get-insight": {
			Name:   "get-insight",
			Fields: fields_get_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsight(ctx, input)
			},
		},
		"get-insight-events": {
			Name:   "get-insight-events",
			Fields: fields_get_insight_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_insight_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInsightEvents(ctx, input)
				}
				var results []*svc.GetInsightEventsOutput
				p := svc.NewGetInsightEventsPaginator(client, input)
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
		"get-insight-impact-graph": {
			Name:   "get-insight-impact-graph",
			Fields: fields_get_insight_impact_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightImpactGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insight_impact_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsightImpactGraph(ctx, input)
			},
		},
		"get-insight-summaries": {
			Name:   "get-insight-summaries",
			Fields: fields_get_insight_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_insight_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInsightSummaries(ctx, input)
				}
				var results []*svc.GetInsightSummariesOutput
				p := svc.NewGetInsightSummariesPaginator(client, input)
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
		"get-retrieved-traces-graph": {
			Name:   "get-retrieved-traces-graph",
			Fields: fields_get_retrieved_traces_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRetrievedTracesGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_retrieved_traces_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRetrievedTracesGraph(ctx, input)
			},
		},
		"get-sampling-rules": {
			Name:   "get-sampling-rules",
			Fields: fields_get_sampling_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSamplingRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_sampling_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSamplingRules(ctx, input)
				}
				var results []*svc.GetSamplingRulesOutput
				p := svc.NewGetSamplingRulesPaginator(client, input)
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
		"get-sampling-statistic-summaries": {
			Name:   "get-sampling-statistic-summaries",
			Fields: fields_get_sampling_statistic_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSamplingStatisticSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_sampling_statistic_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSamplingStatisticSummaries(ctx, input)
				}
				var results []*svc.GetSamplingStatisticSummariesOutput
				p := svc.NewGetSamplingStatisticSummariesPaginator(client, input)
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
		"get-sampling-targets": {
			Name:   "get-sampling-targets",
			Fields: fields_get_sampling_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSamplingTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sampling_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSamplingTargets(ctx, input)
			},
		},
		"get-service-graph": {
			Name:   "get-service-graph",
			Fields: fields_get_service_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceGraphInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_service_graph, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetServiceGraph(ctx, input)
				}
				var results []*svc.GetServiceGraphOutput
				p := svc.NewGetServiceGraphPaginator(client, input)
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
		"get-time-series-service-statistics": {
			Name:   "get-time-series-service-statistics",
			Fields: fields_get_time_series_service_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTimeSeriesServiceStatisticsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_time_series_service_statistics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTimeSeriesServiceStatistics(ctx, input)
				}
				var results []*svc.GetTimeSeriesServiceStatisticsOutput
				p := svc.NewGetTimeSeriesServiceStatisticsPaginator(client, input)
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
		"get-trace-graph": {
			Name:   "get-trace-graph",
			Fields: fields_get_trace_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTraceGraphInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_trace_graph, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTraceGraph(ctx, input)
				}
				var results []*svc.GetTraceGraphOutput
				p := svc.NewGetTraceGraphPaginator(client, input)
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
		"get-trace-segment-destination": {
			Name:   "get-trace-segment-destination",
			Fields: fields_get_trace_segment_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTraceSegmentDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trace_segment_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTraceSegmentDestination(ctx, input)
			},
		},
		"get-trace-summaries": {
			Name:   "get-trace-summaries",
			Fields: fields_get_trace_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTraceSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_trace_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTraceSummaries(ctx, input)
				}
				var results []*svc.GetTraceSummariesOutput
				p := svc.NewGetTraceSummariesPaginator(client, input)
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
		"list-resource-policies": {
			Name:   "list-resource-policies",
			Fields: fields_list_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourcePolicies(ctx, input)
				}
				var results []*svc.ListResourcePoliciesOutput
				p := svc.NewListResourcePoliciesPaginator(client, input)
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
		"list-retrieved-traces": {
			Name:   "list-retrieved-traces",
			Fields: fields_list_retrieved_traces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRetrievedTracesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_retrieved_traces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRetrievedTraces(ctx, input)
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"put-encryption-config": {
			Name:   "put-encryption-config",
			Fields: fields_put_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEncryptionConfig(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"put-telemetry-records": {
			Name:   "put-telemetry-records",
			Fields: fields_put_telemetry_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTelemetryRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_telemetry_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTelemetryRecords(ctx, input)
			},
		},
		"put-trace-segments": {
			Name:   "put-trace-segments",
			Fields: fields_put_trace_segments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTraceSegmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_trace_segments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTraceSegments(ctx, input)
			},
		},
		"start-trace-retrieval": {
			Name:   "start-trace-retrieval",
			Fields: fields_start_trace_retrieval,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTraceRetrievalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_trace_retrieval, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTraceRetrieval(ctx, input)
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
		"update-indexing-rule": {
			Name:   "update-indexing-rule",
			Fields: fields_update_indexing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_indexing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndexingRule(ctx, input)
			},
		},
		"update-sampling-rule": {
			Name:   "update-sampling-rule",
			Fields: fields_update_sampling_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSamplingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sampling_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSamplingRule(ctx, input)
			},
		},
		"update-trace-segment-destination": {
			Name:   "update-trace-segment-destination",
			Fields: fields_update_trace_segment_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTraceSegmentDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trace_segment_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTraceSegmentDestination(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("xray", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
