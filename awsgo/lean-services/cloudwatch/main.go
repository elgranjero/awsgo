package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

var fields_delete_alarm_mute_rule = []leanruntime.Field{
	{Name: "AlarmMuteRuleName", Flag: "alarm-mute-rule-name", Type: "*string", Required: true},
}

var fields_delete_alarms = []leanruntime.Field{
	{Name: "AlarmNames", Flag: "alarm-names", Type: "[]string", Required: true},
}

var fields_delete_anomaly_detector = []leanruntime.Field{
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.Dimension", Required: false},
	{Name: "MetricMathAnomalyDetector", Flag: "metric-math-anomaly-detector", Type: "*types.MetricMathAnomalyDetector", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "SingleMetricAnomalyDetector", Flag: "single-metric-anomaly-detector", Type: "*types.SingleMetricAnomalyDetector", Required: false},
	{Name: "Stat", Flag: "stat", Type: "*string", Required: false},
}

var fields_delete_dashboards = []leanruntime.Field{
	{Name: "DashboardNames", Flag: "dashboard-names", Type: "[]string", Required: true},
}

var fields_delete_insight_rules = []leanruntime.Field{
	{Name: "RuleNames", Flag: "rule-names", Type: "[]string", Required: true},
}

var fields_delete_metric_stream = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_alarm_contributors = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_alarm_history = []leanruntime.Field{
	{Name: "AlarmContributorId", Flag: "alarm-contributor-id", Type: "*string", Required: false},
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: false},
	{Name: "AlarmTypes", Flag: "alarm-types", Type: "[]types.AlarmType", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "HistoryItemType", Flag: "history-item-type", Type: "types.HistoryItemType", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanBy", Flag: "scan-by", Type: "types.ScanBy", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
}

var fields_describe_alarms = []leanruntime.Field{
	{Name: "ActionPrefix", Flag: "action-prefix", Type: "*string", Required: false},
	{Name: "AlarmNamePrefix", Flag: "alarm-name-prefix", Type: "*string", Required: false},
	{Name: "AlarmNames", Flag: "alarm-names", Type: "[]string", Required: false},
	{Name: "AlarmTypes", Flag: "alarm-types", Type: "[]types.AlarmType", Required: false},
	{Name: "ChildrenOfAlarmName", Flag: "children-of-alarm-name", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentsOfAlarmName", Flag: "parents-of-alarm-name", Type: "*string", Required: false},
	{Name: "StateValue", Flag: "state-value", Type: "types.StateValue", Required: false},
}

var fields_describe_alarms_for_metric = []leanruntime.Field{
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.Dimension", Required: false},
	{Name: "ExtendedStatistic", Flag: "extended-statistic", Type: "*string", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.Statistic", Required: false},
	{Name: "Unit", Flag: "unit", Type: "types.StandardUnit", Required: false},
}

var fields_describe_anomaly_detectors = []leanruntime.Field{
	{Name: "AnomalyDetectorTypes", Flag: "anomaly-detector-types", Type: "[]types.AnomalyDetectorType", Required: false},
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.Dimension", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_insight_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disable_alarm_actions = []leanruntime.Field{
	{Name: "AlarmNames", Flag: "alarm-names", Type: "[]string", Required: true},
}

var fields_disable_insight_rules = []leanruntime.Field{
	{Name: "RuleNames", Flag: "rule-names", Type: "[]string", Required: true},
}

var fields_enable_alarm_actions = []leanruntime.Field{
	{Name: "AlarmNames", Flag: "alarm-names", Type: "[]string", Required: true},
}

var fields_enable_insight_rules = []leanruntime.Field{
	{Name: "RuleNames", Flag: "rule-names", Type: "[]string", Required: true},
}

var fields_get_alarm_mute_rule = []leanruntime.Field{
	{Name: "AlarmMuteRuleName", Flag: "alarm-mute-rule-name", Type: "*string", Required: true},
}

var fields_get_dashboard = []leanruntime.Field{
	{Name: "DashboardName", Flag: "dashboard-name", Type: "*string", Required: true},
}

var fields_get_insight_rule_report = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxContributorCount", Flag: "max-contributor-count", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "*string", Required: false},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "LabelOptions", Flag: "label-options", Type: "*types.LabelOptions", Required: false},
	{Name: "MaxDatapoints", Flag: "max-datapoints", Type: "*int32", Required: false},
	{Name: "MetricDataQueries", Flag: "metric-data-queries", Type: "[]types.MetricDataQuery", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanBy", Flag: "scan-by", Type: "types.ScanBy", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_metric_statistics = []leanruntime.Field{
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.Dimension", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "ExtendedStatistics", Flag: "extended-statistics", Type: "[]string", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.Statistic", Required: false},
	{Name: "Unit", Flag: "unit", Type: "types.StandardUnit", Required: false},
}

var fields_get_metric_stream = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_metric_widget_image = []leanruntime.Field{
	{Name: "MetricWidget", Flag: "metric-widget", Type: "*string", Required: true},
	{Name: "OutputFormat", Flag: "output-format", Type: "*string", Required: false},
}

var fields_list_alarm_mute_rules = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.AlarmMuteRuleStatus", Required: false},
}

var fields_list_dashboards = []leanruntime.Field{
	{Name: "DashboardNamePrefix", Flag: "dashboard-name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_insight_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_metric_streams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_metrics = []leanruntime.Field{
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.DimensionFilter", Required: false},
	{Name: "IncludeLinkedAccounts", Flag: "include-linked-accounts", Type: "*bool", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwningAccount", Flag: "owning-account", Type: "*string", Required: false},
	{Name: "RecentlyActive", Flag: "recently-active", Type: "types.RecentlyActive", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_alarm_mute_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpireDate", Flag: "expire-date", Type: "*time.Time", Required: false},
	{Name: "MuteTargets", Flag: "mute-targets", Type: "*types.MuteTargets", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*types.Rule", Required: true},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_anomaly_detector = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.AnomalyDetectorConfiguration", Required: false},
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.Dimension", Required: false},
	{Name: "MetricCharacteristics", Flag: "metric-characteristics", Type: "*types.MetricCharacteristics", Required: false},
	{Name: "MetricMathAnomalyDetector", Flag: "metric-math-anomaly-detector", Type: "*types.MetricMathAnomalyDetector", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "SingleMetricAnomalyDetector", Flag: "single-metric-anomaly-detector", Type: "*types.SingleMetricAnomalyDetector", Required: false},
	{Name: "Stat", Flag: "stat", Type: "*string", Required: false},
}

var fields_put_composite_alarm = []leanruntime.Field{
	{Name: "ActionsEnabled", Flag: "actions-enabled", Type: "*bool", Required: false},
	{Name: "ActionsSuppressor", Flag: "actions-suppressor", Type: "*string", Required: false},
	{Name: "ActionsSuppressorExtensionPeriod", Flag: "actions-suppressor-extension-period", Type: "*int32", Required: false},
	{Name: "ActionsSuppressorWaitPeriod", Flag: "actions-suppressor-wait-period", Type: "*int32", Required: false},
	{Name: "AlarmActions", Flag: "alarm-actions", Type: "[]string", Required: false},
	{Name: "AlarmDescription", Flag: "alarm-description", Type: "*string", Required: false},
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
	{Name: "AlarmRule", Flag: "alarm-rule", Type: "*string", Required: true},
	{Name: "InsufficientDataActions", Flag: "insufficient-data-actions", Type: "[]string", Required: false},
	{Name: "OKActions", Flag: "ok-actions", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_dashboard = []leanruntime.Field{
	{Name: "DashboardBody", Flag: "dashboard-body", Type: "*string", Required: true},
	{Name: "DashboardName", Flag: "dashboard-name", Type: "*string", Required: true},
}

var fields_put_insight_rule = []leanruntime.Field{
	{Name: "ApplyOnTransformedLogs", Flag: "apply-on-transformed-logs", Type: "*bool", Required: false},
	{Name: "RuleDefinition", Flag: "rule-definition", Type: "*string", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "RuleState", Flag: "rule-state", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_managed_insight_rules = []leanruntime.Field{
	{Name: "ManagedRules", Flag: "managed-rules", Type: "[]types.ManagedRule", Required: true},
}

var fields_put_metric_alarm = []leanruntime.Field{
	{Name: "ActionsEnabled", Flag: "actions-enabled", Type: "*bool", Required: false},
	{Name: "AlarmActions", Flag: "alarm-actions", Type: "[]string", Required: false},
	{Name: "AlarmDescription", Flag: "alarm-description", Type: "*string", Required: false},
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
	{Name: "ComparisonOperator", Flag: "comparison-operator", Type: "types.ComparisonOperator", Required: true},
	{Name: "DatapointsToAlarm", Flag: "datapoints-to-alarm", Type: "*int32", Required: false},
	{Name: "Dimensions", Flag: "dimensions", Type: "[]types.Dimension", Required: false},
	{Name: "EvaluateLowSampleCountPercentile", Flag: "evaluate-low-sample-count-percentile", Type: "*string", Required: false},
	{Name: "EvaluationPeriods", Flag: "evaluation-periods", Type: "*int32", Required: true},
	{Name: "ExtendedStatistic", Flag: "extended-statistic", Type: "*string", Required: false},
	{Name: "InsufficientDataActions", Flag: "insufficient-data-actions", Type: "[]string", Required: false},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.MetricDataQuery", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "OKActions", Flag: "ok-actions", Type: "[]string", Required: false},
	{Name: "Period", Flag: "period", Type: "*int32", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.Statistic", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Threshold", Flag: "threshold", Type: "*float64", Required: false},
	{Name: "ThresholdMetricId", Flag: "threshold-metric-id", Type: "*string", Required: false},
	{Name: "TreatMissingData", Flag: "treat-missing-data", Type: "*string", Required: false},
	{Name: "Unit", Flag: "unit", Type: "types.StandardUnit", Required: false},
}

var fields_put_metric_data = []leanruntime.Field{
	{Name: "EntityMetricData", Flag: "entity-metric-data", Type: "[]types.EntityMetricData", Required: false},
	{Name: "MetricData", Flag: "metric-data", Type: "[]types.MetricDatum", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "StrictEntityValidation", Flag: "strict-entity-validation", Type: "*bool", Required: false},
}

var fields_put_metric_stream = []leanruntime.Field{
	{Name: "ExcludeFilters", Flag: "exclude-filters", Type: "[]types.MetricStreamFilter", Required: false},
	{Name: "FirehoseArn", Flag: "firehose-arn", Type: "*string", Required: true},
	{Name: "IncludeFilters", Flag: "include-filters", Type: "[]types.MetricStreamFilter", Required: false},
	{Name: "IncludeLinkedAccountsMetrics", Flag: "include-linked-accounts-metrics", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.MetricStreamOutputFormat", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StatisticsConfigurations", Flag: "statistics-configurations", Type: "[]types.MetricStreamStatisticsConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_set_alarm_state = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
	{Name: "StateReason", Flag: "state-reason", Type: "*string", Required: true},
	{Name: "StateReasonData", Flag: "state-reason-data", Type: "*string", Required: false},
	{Name: "StateValue", Flag: "state-value", Type: "types.StateValue", Required: true},
}

var fields_start_metric_streams = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_stop_metric_streams = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-alarm-mute-rule": {
			Name:   "delete-alarm-mute-rule",
			Fields: fields_delete_alarm_mute_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlarmMuteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alarm_mute_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlarmMuteRule(ctx, input)
			},
		},
		"delete-alarms": {
			Name:   "delete-alarms",
			Fields: fields_delete_alarms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlarmsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alarms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlarms(ctx, input)
			},
		},
		"delete-anomaly-detector": {
			Name:   "delete-anomaly-detector",
			Fields: fields_delete_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnomalyDetector(ctx, input)
			},
		},
		"delete-dashboards": {
			Name:   "delete-dashboards",
			Fields: fields_delete_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDashboardsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dashboards, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDashboards(ctx, input)
			},
		},
		"delete-insight-rules": {
			Name:   "delete-insight-rules",
			Fields: fields_delete_insight_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInsightRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_insight_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInsightRules(ctx, input)
			},
		},
		"delete-metric-stream": {
			Name:   "delete-metric-stream",
			Fields: fields_delete_metric_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMetricStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_metric_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMetricStream(ctx, input)
			},
		},
		"describe-alarm-contributors": {
			Name:   "describe-alarm-contributors",
			Fields: fields_describe_alarm_contributors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlarmContributorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_alarm_contributors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlarmContributors(ctx, input)
			},
		},
		"describe-alarm-history": {
			Name:   "describe-alarm-history",
			Fields: fields_describe_alarm_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlarmHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_alarm_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAlarmHistory(ctx, input)
				}
				var results []*svc.DescribeAlarmHistoryOutput
				p := svc.NewDescribeAlarmHistoryPaginator(client, input)
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
		"describe-alarms": {
			Name:   "describe-alarms",
			Fields: fields_describe_alarms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlarmsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_alarms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAlarms(ctx, input)
				}
				var results []*svc.DescribeAlarmsOutput
				p := svc.NewDescribeAlarmsPaginator(client, input)
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
		"describe-alarms-for-metric": {
			Name:   "describe-alarms-for-metric",
			Fields: fields_describe_alarms_for_metric,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlarmsForMetricInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_alarms_for_metric, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlarmsForMetric(ctx, input)
			},
		},
		"describe-anomaly-detectors": {
			Name:   "describe-anomaly-detectors",
			Fields: fields_describe_anomaly_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnomalyDetectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_anomaly_detectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAnomalyDetectors(ctx, input)
				}
				var results []*svc.DescribeAnomalyDetectorsOutput
				p := svc.NewDescribeAnomalyDetectorsPaginator(client, input)
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
		"describe-insight-rules": {
			Name:   "describe-insight-rules",
			Fields: fields_describe_insight_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInsightRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_insight_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInsightRules(ctx, input)
				}
				var results []*svc.DescribeInsightRulesOutput
				p := svc.NewDescribeInsightRulesPaginator(client, input)
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
		"disable-alarm-actions": {
			Name:   "disable-alarm-actions",
			Fields: fields_disable_alarm_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAlarmActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_alarm_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAlarmActions(ctx, input)
			},
		},
		"disable-insight-rules": {
			Name:   "disable-insight-rules",
			Fields: fields_disable_insight_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableInsightRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_insight_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableInsightRules(ctx, input)
			},
		},
		"enable-alarm-actions": {
			Name:   "enable-alarm-actions",
			Fields: fields_enable_alarm_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAlarmActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_alarm_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAlarmActions(ctx, input)
			},
		},
		"enable-insight-rules": {
			Name:   "enable-insight-rules",
			Fields: fields_enable_insight_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableInsightRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_insight_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableInsightRules(ctx, input)
			},
		},
		"get-alarm-mute-rule": {
			Name:   "get-alarm-mute-rule",
			Fields: fields_get_alarm_mute_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAlarmMuteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_alarm_mute_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAlarmMuteRule(ctx, input)
			},
		},
		"get-dashboard": {
			Name:   "get-dashboard",
			Fields: fields_get_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDashboard(ctx, input)
			},
		},
		"get-insight-rule-report": {
			Name:   "get-insight-rule-report",
			Fields: fields_get_insight_rule_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightRuleReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insight_rule_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsightRuleReport(ctx, input)
			},
		},
		"get-metric-data": {
			Name:   "get-metric-data",
			Fields: fields_get_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_metric_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMetricData(ctx, input)
				}
				var results []*svc.GetMetricDataOutput
				p := svc.NewGetMetricDataPaginator(client, input)
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
		"get-metric-statistics": {
			Name:   "get-metric-statistics",
			Fields: fields_get_metric_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metric_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetricStatistics(ctx, input)
			},
		},
		"get-metric-stream": {
			Name:   "get-metric-stream",
			Fields: fields_get_metric_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metric_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetricStream(ctx, input)
			},
		},
		"get-metric-widget-image": {
			Name:   "get-metric-widget-image",
			Fields: fields_get_metric_widget_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricWidgetImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metric_widget_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetricWidgetImage(ctx, input)
			},
		},
		"list-alarm-mute-rules": {
			Name:   "list-alarm-mute-rules",
			Fields: fields_list_alarm_mute_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlarmMuteRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_alarm_mute_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAlarmMuteRules(ctx, input)
				}
				var results []*svc.ListAlarmMuteRulesOutput
				p := svc.NewListAlarmMuteRulesPaginator(client, input)
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
		"list-dashboards": {
			Name:   "list-dashboards",
			Fields: fields_list_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDashboardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dashboards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDashboards(ctx, input)
				}
				var results []*svc.ListDashboardsOutput
				p := svc.NewListDashboardsPaginator(client, input)
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
		"list-managed-insight-rules": {
			Name:   "list-managed-insight-rules",
			Fields: fields_list_managed_insight_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedInsightRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_insight_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedInsightRules(ctx, input)
				}
				var results []*svc.ListManagedInsightRulesOutput
				p := svc.NewListManagedInsightRulesPaginator(client, input)
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
		"list-metric-streams": {
			Name:   "list-metric-streams",
			Fields: fields_list_metric_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetricStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metric_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetricStreams(ctx, input)
				}
				var results []*svc.ListMetricStreamsOutput
				p := svc.NewListMetricStreamsPaginator(client, input)
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
		"list-metrics": {
			Name:   "list-metrics",
			Fields: fields_list_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetrics(ctx, input)
				}
				var results []*svc.ListMetricsOutput
				p := svc.NewListMetricsPaginator(client, input)
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
		"put-alarm-mute-rule": {
			Name:   "put-alarm-mute-rule",
			Fields: fields_put_alarm_mute_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAlarmMuteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_alarm_mute_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAlarmMuteRule(ctx, input)
			},
		},
		"put-anomaly-detector": {
			Name:   "put-anomaly-detector",
			Fields: fields_put_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAnomalyDetector(ctx, input)
			},
		},
		"put-composite-alarm": {
			Name:   "put-composite-alarm",
			Fields: fields_put_composite_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCompositeAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_composite_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCompositeAlarm(ctx, input)
			},
		},
		"put-dashboard": {
			Name:   "put-dashboard",
			Fields: fields_put_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDashboard(ctx, input)
			},
		},
		"put-insight-rule": {
			Name:   "put-insight-rule",
			Fields: fields_put_insight_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInsightRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_insight_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInsightRule(ctx, input)
			},
		},
		"put-managed-insight-rules": {
			Name:   "put-managed-insight-rules",
			Fields: fields_put_managed_insight_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutManagedInsightRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_managed_insight_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutManagedInsightRules(ctx, input)
			},
		},
		"put-metric-alarm": {
			Name:   "put-metric-alarm",
			Fields: fields_put_metric_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetricAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metric_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetricAlarm(ctx, input)
			},
		},
		"put-metric-data": {
			Name:   "put-metric-data",
			Fields: fields_put_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetricData(ctx, input)
			},
		},
		"put-metric-stream": {
			Name:   "put-metric-stream",
			Fields: fields_put_metric_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetricStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metric_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetricStream(ctx, input)
			},
		},
		"set-alarm-state": {
			Name:   "set-alarm-state",
			Fields: fields_set_alarm_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetAlarmStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_alarm_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetAlarmState(ctx, input)
			},
		},
		"start-metric-streams": {
			Name:   "start-metric-streams",
			Fields: fields_start_metric_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetricStreamsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metric_streams, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetricStreams(ctx, input)
			},
		},
		"stop-metric-streams": {
			Name:   "stop-metric-streams",
			Fields: fields_stop_metric_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMetricStreamsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_metric_streams, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMetricStreams(ctx, input)
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
	}
	if err := leanruntime.Execute("cloudwatch", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
