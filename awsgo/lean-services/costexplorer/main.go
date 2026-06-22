package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

var fields_create_anomaly_monitor = []leanruntime.Field{
	{Name: "AnomalyMonitor", Flag: "anomaly-monitor", Type: "*types.AnomalyMonitor", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
}

var fields_create_anomaly_subscription = []leanruntime.Field{
	{Name: "AnomalySubscription", Flag: "anomaly-subscription", Type: "*types.AnomalySubscription", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
}

var fields_create_cost_category_definition = []leanruntime.Field{
	{Name: "DefaultValue", Flag: "default-value", Type: "*string", Required: false},
	{Name: "EffectiveStart", Flag: "effective-start", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "RuleVersion", Flag: "rule-version", Type: "types.CostCategoryRuleVersion", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.CostCategoryRule", Required: true},
	{Name: "SplitChargeRules", Flag: "split-charge-rules", Type: "[]types.CostCategorySplitChargeRule", Required: false},
}

var fields_delete_anomaly_monitor = []leanruntime.Field{
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: true},
}

var fields_delete_anomaly_subscription = []leanruntime.Field{
	{Name: "SubscriptionArn", Flag: "subscription-arn", Type: "*string", Required: true},
}

var fields_delete_cost_category_definition = []leanruntime.Field{
	{Name: "CostCategoryArn", Flag: "cost-category-arn", Type: "*string", Required: true},
}

var fields_describe_cost_category_definition = []leanruntime.Field{
	{Name: "CostCategoryArn", Flag: "cost-category-arn", Type: "*string", Required: true},
	{Name: "EffectiveOn", Flag: "effective-on", Type: "*string", Required: false},
}

var fields_get_anomalies = []leanruntime.Field{
	{Name: "DateInterval", Flag: "date-interval", Type: "*types.AnomalyDateInterval", Required: true},
	{Name: "Feedback", Flag: "feedback", Type: "types.AnomalyFeedbackType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "TotalImpact", Flag: "total-impact", Type: "*types.TotalImpactFilter", Required: false},
}

var fields_get_anomaly_monitors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorArnList", Flag: "monitor-arn-list", Type: "[]string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
}

var fields_get_anomaly_subscriptions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "SubscriptionArnList", Flag: "subscription-arn-list", Type: "[]string", Required: false},
}

var fields_get_approximate_usage_records = []leanruntime.Field{
	{Name: "ApproximationDimension", Flag: "approximation-dimension", Type: "types.ApproximationDimension", Required: true},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: true},
	{Name: "Services", Flag: "services", Type: "[]string", Required: false},
}

var fields_get_commitment_purchase_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
}

var fields_get_cost_and_usage = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: true},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: true},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_cost_and_usage_comparisons = []leanruntime.Field{
	{Name: "BaselineTimePeriod", Flag: "baseline-time-period", Type: "*types.DateInterval", Required: true},
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "ComparisonTimePeriod", Flag: "comparison-time-period", Type: "*types.DateInterval", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricForComparison", Flag: "metric-for-comparison", Type: "*string", Required: true},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
}

var fields_get_cost_and_usage_with_resources = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: true},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: true},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_cost_categories = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "CostCategoryName", Flag: "cost-category-name", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "SearchString", Flag: "search-string", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "[]types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_cost_comparison_drivers = []leanruntime.Field{
	{Name: "BaselineTimePeriod", Flag: "baseline-time-period", Type: "*types.DateInterval", Required: true},
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "ComparisonTimePeriod", Flag: "comparison-time-period", Type: "*types.DateInterval", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricForComparison", Flag: "metric-for-comparison", Type: "*string", Required: true},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
}

var fields_get_cost_forecast = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: true},
	{Name: "Metric", Flag: "metric", Type: "types.Metric", Required: true},
	{Name: "PredictionIntervalLevel", Flag: "prediction-interval-level", Type: "*int32", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_dimension_values = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "types.Context", Required: false},
	{Name: "Dimension", Flag: "dimension", Type: "types.Dimension", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "SearchString", Flag: "search-string", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "[]types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_reservation_coverage = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_reservation_purchase_recommendation = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AccountScope", Flag: "account-scope", Type: "types.AccountScope", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "LookbackPeriodInDays", Flag: "lookback-period-in-days", Type: "types.LookbackPeriodInDays", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PaymentOption", Flag: "payment-option", Type: "types.PaymentOption", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "ServiceSpecification", Flag: "service-specification", Type: "*types.ServiceSpecification", Required: false},
	{Name: "TermInYears", Flag: "term-in-years", Type: "types.TermInYears", Required: false},
}

var fields_get_reservation_utilization = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_rightsizing_recommendation = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.RightsizingRecommendationConfiguration", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
}

var fields_get_savings_plan_purchase_recommendation_details = []leanruntime.Field{
	{Name: "RecommendationDetailId", Flag: "recommendation-detail-id", Type: "*string", Required: true},
}

var fields_get_savings_plans_coverage = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupDefinition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_savings_plans_purchase_recommendation = []leanruntime.Field{
	{Name: "AccountScope", Flag: "account-scope", Type: "types.AccountScope", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "LookbackPeriodInDays", Flag: "lookback-period-in-days", Type: "types.LookbackPeriodInDays", Required: true},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PaymentOption", Flag: "payment-option", Type: "types.PaymentOption", Required: true},
	{Name: "SavingsPlansType", Flag: "savings-plans-type", Type: "types.SupportedSavingsPlansType", Required: true},
	{Name: "TermInYears", Flag: "term-in-years", Type: "types.TermInYears", Required: true},
}

var fields_get_savings_plans_utilization = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_savings_plans_utilization_details = []leanruntime.Field{
	{Name: "DataType", Flag: "data-type", Type: "[]types.SavingsPlansDataType", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.SortDefinition", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_tags = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "SearchString", Flag: "search-string", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "[]types.SortDefinition", Required: false},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_get_usage_forecast = []leanruntime.Field{
	{Name: "BillingViewArn", Flag: "billing-view-arn", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "Granularity", Flag: "granularity", Type: "types.Granularity", Required: true},
	{Name: "Metric", Flag: "metric", Type: "types.Metric", Required: true},
	{Name: "PredictionIntervalLevel", Flag: "prediction-interval-level", Type: "*int32", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.DateInterval", Required: true},
}

var fields_list_commitment_purchase_analyses = []leanruntime.Field{
	{Name: "AnalysisIds", Flag: "analysis-ids", Type: "[]string", Required: false},
	{Name: "AnalysisStatus", Flag: "analysis-status", Type: "types.AnalysisStatus", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
}

var fields_list_cost_allocation_tag_backfill_history = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cost_allocation_tags = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CostAllocationTagStatus", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CostAllocationTagType", Required: false},
}

var fields_list_cost_category_definitions = []leanruntime.Field{
	{Name: "EffectiveOn", Flag: "effective-on", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SupportedResourceTypes", Flag: "supported-resource-types", Type: "[]string", Required: false},
}

var fields_list_cost_category_resource_associations = []leanruntime.Field{
	{Name: "CostCategoryArn", Flag: "cost-category-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_savings_plans_purchase_recommendation_generation = []leanruntime.Field{
	{Name: "GenerationStatus", Flag: "generation-status", Type: "types.GenerationStatus", Required: false},
	{Name: "NextPageToken", Flag: "next-page-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "RecommendationIds", Flag: "recommendation-ids", Type: "[]string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_provide_anomaly_feedback = []leanruntime.Field{
	{Name: "AnomalyId", Flag: "anomaly-id", Type: "*string", Required: true},
	{Name: "Feedback", Flag: "feedback", Type: "types.AnomalyFeedbackType", Required: true},
}

var fields_start_commitment_purchase_analysis = []leanruntime.Field{
	{Name: "CommitmentPurchaseAnalysisConfiguration", Flag: "commitment-purchase-analysis-configuration", Type: "*types.CommitmentPurchaseAnalysisConfiguration", Required: true},
}

var fields_start_cost_allocation_tag_backfill = []leanruntime.Field{
	{Name: "BackfillFrom", Flag: "backfill-from", Type: "*string", Required: true},
}

var fields_start_savings_plans_purchase_recommendation_generation = []leanruntime.Field{}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTagKeys", Flag: "resource-tag-keys", Type: "[]string", Required: true},
}

var fields_update_anomaly_monitor = []leanruntime.Field{
	{Name: "MonitorArn", Flag: "monitor-arn", Type: "*string", Required: true},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: false},
}

var fields_update_anomaly_subscription = []leanruntime.Field{
	{Name: "Frequency", Flag: "frequency", Type: "types.AnomalySubscriptionFrequency", Required: false},
	{Name: "MonitorArnList", Flag: "monitor-arn-list", Type: "[]string", Required: false},
	{Name: "Subscribers", Flag: "subscribers", Type: "[]types.Subscriber", Required: false},
	{Name: "SubscriptionArn", Flag: "subscription-arn", Type: "*string", Required: true},
	{Name: "SubscriptionName", Flag: "subscription-name", Type: "*string", Required: false},
	{Name: "Threshold", Flag: "threshold", Type: "*float64", Required: false},
	{Name: "ThresholdExpression", Flag: "threshold-expression", Type: "*types.Expression", Required: false},
}

var fields_update_cost_allocation_tags_status = []leanruntime.Field{
	{Name: "CostAllocationTagsStatus", Flag: "cost-allocation-tags-status", Type: "[]types.CostAllocationTagStatusEntry", Required: true},
}

var fields_update_cost_category_definition = []leanruntime.Field{
	{Name: "CostCategoryArn", Flag: "cost-category-arn", Type: "*string", Required: true},
	{Name: "DefaultValue", Flag: "default-value", Type: "*string", Required: false},
	{Name: "EffectiveStart", Flag: "effective-start", Type: "*string", Required: false},
	{Name: "RuleVersion", Flag: "rule-version", Type: "types.CostCategoryRuleVersion", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.CostCategoryRule", Required: true},
	{Name: "SplitChargeRules", Flag: "split-charge-rules", Type: "[]types.CostCategorySplitChargeRule", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-anomaly-monitor": {
			Name:   "create-anomaly-monitor",
			Fields: fields_create_anomaly_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnomalyMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_anomaly_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnomalyMonitor(ctx, input)
			},
		},
		"create-anomaly-subscription": {
			Name:   "create-anomaly-subscription",
			Fields: fields_create_anomaly_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnomalySubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_anomaly_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnomalySubscription(ctx, input)
			},
		},
		"create-cost-category-definition": {
			Name:   "create-cost-category-definition",
			Fields: fields_create_cost_category_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCostCategoryDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cost_category_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCostCategoryDefinition(ctx, input)
			},
		},
		"delete-anomaly-monitor": {
			Name:   "delete-anomaly-monitor",
			Fields: fields_delete_anomaly_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnomalyMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_anomaly_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnomalyMonitor(ctx, input)
			},
		},
		"delete-anomaly-subscription": {
			Name:   "delete-anomaly-subscription",
			Fields: fields_delete_anomaly_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnomalySubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_anomaly_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnomalySubscription(ctx, input)
			},
		},
		"delete-cost-category-definition": {
			Name:   "delete-cost-category-definition",
			Fields: fields_delete_cost_category_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCostCategoryDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cost_category_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCostCategoryDefinition(ctx, input)
			},
		},
		"describe-cost-category-definition": {
			Name:   "describe-cost-category-definition",
			Fields: fields_describe_cost_category_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCostCategoryDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cost_category_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCostCategoryDefinition(ctx, input)
			},
		},
		"get-anomalies": {
			Name:   "get-anomalies",
			Fields: fields_get_anomalies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnomaliesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_anomalies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAnomalies(ctx, input)
				}
				var results []*svc.GetAnomaliesOutput
				p := svc.NewGetAnomaliesPaginator(client, input)
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
		"get-anomaly-monitors": {
			Name:   "get-anomaly-monitors",
			Fields: fields_get_anomaly_monitors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnomalyMonitorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_anomaly_monitors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAnomalyMonitors(ctx, input)
				}
				var results []*svc.GetAnomalyMonitorsOutput
				p := svc.NewGetAnomalyMonitorsPaginator(client, input)
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
		"get-anomaly-subscriptions": {
			Name:   "get-anomaly-subscriptions",
			Fields: fields_get_anomaly_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnomalySubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_anomaly_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAnomalySubscriptions(ctx, input)
				}
				var results []*svc.GetAnomalySubscriptionsOutput
				p := svc.NewGetAnomalySubscriptionsPaginator(client, input)
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
		"get-approximate-usage-records": {
			Name:   "get-approximate-usage-records",
			Fields: fields_get_approximate_usage_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApproximateUsageRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_approximate_usage_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApproximateUsageRecords(ctx, input)
			},
		},
		"get-commitment-purchase-analysis": {
			Name:   "get-commitment-purchase-analysis",
			Fields: fields_get_commitment_purchase_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommitmentPurchaseAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_commitment_purchase_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCommitmentPurchaseAnalysis(ctx, input)
			},
		},
		"get-cost-and-usage": {
			Name:   "get-cost-and-usage",
			Fields: fields_get_cost_and_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostAndUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cost_and_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCostAndUsage(ctx, input)
			},
		},
		"get-cost-and-usage-comparisons": {
			Name:   "get-cost-and-usage-comparisons",
			Fields: fields_get_cost_and_usage_comparisons,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostAndUsageComparisonsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_cost_and_usage_comparisons, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCostAndUsageComparisons(ctx, input)
				}
				var results []*svc.GetCostAndUsageComparisonsOutput
				p := svc.NewGetCostAndUsageComparisonsPaginator(client, input)
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
		"get-cost-and-usage-with-resources": {
			Name:   "get-cost-and-usage-with-resources",
			Fields: fields_get_cost_and_usage_with_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostAndUsageWithResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cost_and_usage_with_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCostAndUsageWithResources(ctx, input)
			},
		},
		"get-cost-categories": {
			Name:   "get-cost-categories",
			Fields: fields_get_cost_categories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostCategoriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cost_categories, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCostCategories(ctx, input)
			},
		},
		"get-cost-comparison-drivers": {
			Name:   "get-cost-comparison-drivers",
			Fields: fields_get_cost_comparison_drivers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostComparisonDriversInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_cost_comparison_drivers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCostComparisonDrivers(ctx, input)
				}
				var results []*svc.GetCostComparisonDriversOutput
				p := svc.NewGetCostComparisonDriversPaginator(client, input)
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
		"get-cost-forecast": {
			Name:   "get-cost-forecast",
			Fields: fields_get_cost_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cost_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCostForecast(ctx, input)
			},
		},
		"get-dimension-values": {
			Name:   "get-dimension-values",
			Fields: fields_get_dimension_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDimensionValuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dimension_values, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDimensionValues(ctx, input)
			},
		},
		"get-reservation-coverage": {
			Name:   "get-reservation-coverage",
			Fields: fields_get_reservation_coverage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservationCoverageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reservation_coverage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReservationCoverage(ctx, input)
			},
		},
		"get-reservation-purchase-recommendation": {
			Name:   "get-reservation-purchase-recommendation",
			Fields: fields_get_reservation_purchase_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservationPurchaseRecommendationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_reservation_purchase_recommendation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetReservationPurchaseRecommendation(ctx, input)
				}
				var results []*svc.GetReservationPurchaseRecommendationOutput
				p := svc.NewGetReservationPurchaseRecommendationPaginator(client, input)
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
		"get-reservation-utilization": {
			Name:   "get-reservation-utilization",
			Fields: fields_get_reservation_utilization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservationUtilizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reservation_utilization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReservationUtilization(ctx, input)
			},
		},
		"get-rightsizing-recommendation": {
			Name:   "get-rightsizing-recommendation",
			Fields: fields_get_rightsizing_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRightsizingRecommendationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_rightsizing_recommendation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetRightsizingRecommendation(ctx, input)
				}
				var results []*svc.GetRightsizingRecommendationOutput
				p := svc.NewGetRightsizingRecommendationPaginator(client, input)
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
		"get-savings-plan-purchase-recommendation-details": {
			Name:   "get-savings-plan-purchase-recommendation-details",
			Fields: fields_get_savings_plan_purchase_recommendation_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSavingsPlanPurchaseRecommendationDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_savings_plan_purchase_recommendation_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSavingsPlanPurchaseRecommendationDetails(ctx, input)
			},
		},
		"get-savings-plans-coverage": {
			Name:   "get-savings-plans-coverage",
			Fields: fields_get_savings_plans_coverage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSavingsPlansCoverageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_savings_plans_coverage, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSavingsPlansCoverage(ctx, input)
				}
				var results []*svc.GetSavingsPlansCoverageOutput
				p := svc.NewGetSavingsPlansCoveragePaginator(client, input)
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
		"get-savings-plans-purchase-recommendation": {
			Name:   "get-savings-plans-purchase-recommendation",
			Fields: fields_get_savings_plans_purchase_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSavingsPlansPurchaseRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_savings_plans_purchase_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSavingsPlansPurchaseRecommendation(ctx, input)
			},
		},
		"get-savings-plans-utilization": {
			Name:   "get-savings-plans-utilization",
			Fields: fields_get_savings_plans_utilization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSavingsPlansUtilizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_savings_plans_utilization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSavingsPlansUtilization(ctx, input)
			},
		},
		"get-savings-plans-utilization-details": {
			Name:   "get-savings-plans-utilization-details",
			Fields: fields_get_savings_plans_utilization_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSavingsPlansUtilizationDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_savings_plans_utilization_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSavingsPlansUtilizationDetails(ctx, input)
				}
				var results []*svc.GetSavingsPlansUtilizationDetailsOutput
				p := svc.NewGetSavingsPlansUtilizationDetailsPaginator(client, input)
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
		"get-tags": {
			Name:   "get-tags",
			Fields: fields_get_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTags(ctx, input)
			},
		},
		"get-usage-forecast": {
			Name:   "get-usage-forecast",
			Fields: fields_get_usage_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUsageForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_usage_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUsageForecast(ctx, input)
			},
		},
		"list-commitment-purchase-analyses": {
			Name:   "list-commitment-purchase-analyses",
			Fields: fields_list_commitment_purchase_analyses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommitmentPurchaseAnalysesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_commitment_purchase_analyses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommitmentPurchaseAnalyses(ctx, input)
				}
				var results []*svc.ListCommitmentPurchaseAnalysesOutput
				p := svc.NewListCommitmentPurchaseAnalysesPaginator(client, input)
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
		"list-cost-allocation-tag-backfill-history": {
			Name:   "list-cost-allocation-tag-backfill-history",
			Fields: fields_list_cost_allocation_tag_backfill_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCostAllocationTagBackfillHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cost_allocation_tag_backfill_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCostAllocationTagBackfillHistory(ctx, input)
				}
				var results []*svc.ListCostAllocationTagBackfillHistoryOutput
				p := svc.NewListCostAllocationTagBackfillHistoryPaginator(client, input)
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
		"list-cost-allocation-tags": {
			Name:   "list-cost-allocation-tags",
			Fields: fields_list_cost_allocation_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCostAllocationTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cost_allocation_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCostAllocationTags(ctx, input)
				}
				var results []*svc.ListCostAllocationTagsOutput
				p := svc.NewListCostAllocationTagsPaginator(client, input)
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
		"list-cost-category-definitions": {
			Name:   "list-cost-category-definitions",
			Fields: fields_list_cost_category_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCostCategoryDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cost_category_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCostCategoryDefinitions(ctx, input)
				}
				var results []*svc.ListCostCategoryDefinitionsOutput
				p := svc.NewListCostCategoryDefinitionsPaginator(client, input)
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
		"list-cost-category-resource-associations": {
			Name:   "list-cost-category-resource-associations",
			Fields: fields_list_cost_category_resource_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCostCategoryResourceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cost_category_resource_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCostCategoryResourceAssociations(ctx, input)
				}
				var results []*svc.ListCostCategoryResourceAssociationsOutput
				p := svc.NewListCostCategoryResourceAssociationsPaginator(client, input)
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
		"list-savings-plans-purchase-recommendation-generation": {
			Name:   "list-savings-plans-purchase-recommendation-generation",
			Fields: fields_list_savings_plans_purchase_recommendation_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSavingsPlansPurchaseRecommendationGenerationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_savings_plans_purchase_recommendation_generation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSavingsPlansPurchaseRecommendationGeneration(ctx, input)
				}
				var results []*svc.ListSavingsPlansPurchaseRecommendationGenerationOutput
				p := svc.NewListSavingsPlansPurchaseRecommendationGenerationPaginator(client, input)
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
		"provide-anomaly-feedback": {
			Name:   "provide-anomaly-feedback",
			Fields: fields_provide_anomaly_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvideAnomalyFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provide_anomaly_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvideAnomalyFeedback(ctx, input)
			},
		},
		"start-commitment-purchase-analysis": {
			Name:   "start-commitment-purchase-analysis",
			Fields: fields_start_commitment_purchase_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCommitmentPurchaseAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_commitment_purchase_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCommitmentPurchaseAnalysis(ctx, input)
			},
		},
		"start-cost-allocation-tag-backfill": {
			Name:   "start-cost-allocation-tag-backfill",
			Fields: fields_start_cost_allocation_tag_backfill,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCostAllocationTagBackfillInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_cost_allocation_tag_backfill, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCostAllocationTagBackfill(ctx, input)
			},
		},
		"start-savings-plans-purchase-recommendation-generation": {
			Name:   "start-savings-plans-purchase-recommendation-generation",
			Fields: fields_start_savings_plans_purchase_recommendation_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSavingsPlansPurchaseRecommendationGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_savings_plans_purchase_recommendation_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSavingsPlansPurchaseRecommendationGeneration(ctx, input)
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
		"update-anomaly-monitor": {
			Name:   "update-anomaly-monitor",
			Fields: fields_update_anomaly_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnomalyMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_anomaly_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnomalyMonitor(ctx, input)
			},
		},
		"update-anomaly-subscription": {
			Name:   "update-anomaly-subscription",
			Fields: fields_update_anomaly_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnomalySubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_anomaly_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnomalySubscription(ctx, input)
			},
		},
		"update-cost-allocation-tags-status": {
			Name:   "update-cost-allocation-tags-status",
			Fields: fields_update_cost_allocation_tags_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCostAllocationTagsStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cost_allocation_tags_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCostAllocationTagsStatus(ctx, input)
			},
		},
		"update-cost-category-definition": {
			Name:   "update-cost-category-definition",
			Fields: fields_update_cost_category_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCostCategoryDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cost_category_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCostCategoryDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("costexplorer", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
