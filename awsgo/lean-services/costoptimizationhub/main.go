package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"
)

var fields_get_preferences = []leanruntime.Field{}

var fields_get_recommendation = []leanruntime.Field{
	{Name: "RecommendationId", Flag: "recommendation-id", Type: "*string", Required: true},
}

var fields_list_efficiency_metrics = []leanruntime.Field{
	{Name: "Granularity", Flag: "granularity", Type: "types.GranularityType", Required: true},
	{Name: "GroupBy", Flag: "group-by", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "*types.OrderBy", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.TimePeriod", Required: true},
}

var fields_list_enrollment_statuses = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "IncludeOrganizationInfo", Flag: "include-organization-info", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recommendation_summaries = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Filter", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.SummaryMetrics", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recommendations = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Filter", Required: false},
	{Name: "IncludeAllRecommendations", Flag: "include-all-recommendations", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "*types.OrderBy", Required: false},
}

var fields_update_enrollment_status = []leanruntime.Field{
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "*bool", Required: false},
	{Name: "Status", Flag: "status", Type: "types.EnrollmentStatus", Required: true},
}

var fields_update_preferences = []leanruntime.Field{
	{Name: "MemberAccountDiscountVisibility", Flag: "member-account-discount-visibility", Type: "types.MemberAccountDiscountVisibility", Required: false},
	{Name: "PreferredCommitment", Flag: "preferred-commitment", Type: "*types.PreferredCommitment", Required: false},
	{Name: "SavingsEstimationMode", Flag: "savings-estimation-mode", Type: "types.SavingsEstimationMode", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-preferences": {
			Name:   "get-preferences",
			Fields: fields_get_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPreferences(ctx, input)
			},
		},
		"get-recommendation": {
			Name:   "get-recommendation",
			Fields: fields_get_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommendation(ctx, input)
			},
		},
		"list-efficiency-metrics": {
			Name:   "list-efficiency-metrics",
			Fields: fields_list_efficiency_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEfficiencyMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_efficiency_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEfficiencyMetrics(ctx, input)
				}
				var results []*svc.ListEfficiencyMetricsOutput
				p := svc.NewListEfficiencyMetricsPaginator(client, input)
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
		"list-enrollment-statuses": {
			Name:   "list-enrollment-statuses",
			Fields: fields_list_enrollment_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnrollmentStatusesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_enrollment_statuses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnrollmentStatuses(ctx, input)
				}
				var results []*svc.ListEnrollmentStatusesOutput
				p := svc.NewListEnrollmentStatusesPaginator(client, input)
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
		"list-recommendation-summaries": {
			Name:   "list-recommendation-summaries",
			Fields: fields_list_recommendation_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendation_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendationSummaries(ctx, input)
				}
				var results []*svc.ListRecommendationSummariesOutput
				p := svc.NewListRecommendationSummariesPaginator(client, input)
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
		"list-recommendations": {
			Name:   "list-recommendations",
			Fields: fields_list_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendations(ctx, input)
				}
				var results []*svc.ListRecommendationsOutput
				p := svc.NewListRecommendationsPaginator(client, input)
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
		"update-enrollment-status": {
			Name:   "update-enrollment-status",
			Fields: fields_update_enrollment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnrollmentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_enrollment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnrollmentStatus(ctx, input)
			},
		},
		"update-preferences": {
			Name:   "update-preferences",
			Fields: fields_update_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePreferences(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("costoptimizationhub", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
