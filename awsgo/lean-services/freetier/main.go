package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/freetier"
)

var fields_get_account_activity = []leanruntime.Field{
	{Name: "ActivityId", Flag: "activity-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
}

var fields_get_account_plan_state = []leanruntime.Field{}

var fields_get_free_tier_usage = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Expression", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_account_activities = []leanruntime.Field{
	{Name: "FilterActivityStatuses", Flag: "filter-activity-statuses", Type: "[]types.ActivityStatus", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_upgrade_account_plan = []leanruntime.Field{
	{Name: "AccountPlanType", Flag: "account-plan-type", Type: "types.AccountPlanType", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-account-activity": {
			Name:   "get-account-activity",
			Fields: fields_get_account_activity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountActivityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_activity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountActivity(ctx, input)
			},
		},
		"get-account-plan-state": {
			Name:   "get-account-plan-state",
			Fields: fields_get_account_plan_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountPlanStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_plan_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountPlanState(ctx, input)
			},
		},
		"get-free-tier-usage": {
			Name:   "get-free-tier-usage",
			Fields: fields_get_free_tier_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFreeTierUsageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_free_tier_usage, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFreeTierUsage(ctx, input)
				}
				var results []*svc.GetFreeTierUsageOutput
				p := svc.NewGetFreeTierUsagePaginator(client, input)
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
		"list-account-activities": {
			Name:   "list-account-activities",
			Fields: fields_list_account_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountActivitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_activities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountActivities(ctx, input)
				}
				var results []*svc.ListAccountActivitiesOutput
				p := svc.NewListAccountActivitiesPaginator(client, input)
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
		"upgrade-account-plan": {
			Name:   "upgrade-account-plan",
			Fields: fields_upgrade_account_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeAccountPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_account_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeAccountPlan(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("freetier", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
