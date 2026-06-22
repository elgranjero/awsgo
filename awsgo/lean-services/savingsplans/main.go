package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/savingsplans"
)

var fields_create_savings_plan = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Commitment", Flag: "commitment", Type: "*string", Required: true},
	{Name: "PurchaseTime", Flag: "purchase-time", Type: "*time.Time", Required: false},
	{Name: "SavingsPlanOfferingId", Flag: "savings-plan-offering-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UpfrontPaymentAmount", Flag: "upfront-payment-amount", Type: "*string", Required: false},
}

var fields_delete_queued_savings_plan = []leanruntime.Field{
	{Name: "SavingsPlanId", Flag: "savings-plan-id", Type: "*string", Required: true},
}

var fields_describe_savings_plan_rates = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SavingsPlanRateFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SavingsPlanId", Flag: "savings-plan-id", Type: "*string", Required: true},
}

var fields_describe_savings_plans = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SavingsPlanFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SavingsPlanArns", Flag: "savings-plan-arns", Type: "[]string", Required: false},
	{Name: "SavingsPlanIds", Flag: "savings-plan-ids", Type: "[]string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.SavingsPlanState", Required: false},
}

var fields_describe_savings_plans_offering_rates = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SavingsPlanOfferingRateFilterElement", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Operations", Flag: "operations", Type: "[]string", Required: false},
	{Name: "Products", Flag: "products", Type: "[]types.SavingsPlanProductType", Required: false},
	{Name: "SavingsPlanOfferingIds", Flag: "savings-plan-offering-ids", Type: "[]string", Required: false},
	{Name: "SavingsPlanPaymentOptions", Flag: "savings-plan-payment-options", Type: "[]types.SavingsPlanPaymentOption", Required: false},
	{Name: "SavingsPlanTypes", Flag: "savings-plan-types", Type: "[]types.SavingsPlanType", Required: false},
	{Name: "ServiceCodes", Flag: "service-codes", Type: "[]types.SavingsPlanRateServiceCode", Required: false},
	{Name: "UsageTypes", Flag: "usage-types", Type: "[]string", Required: false},
}

var fields_describe_savings_plans_offerings = []leanruntime.Field{
	{Name: "Currencies", Flag: "currencies", Type: "[]types.CurrencyCode", Required: false},
	{Name: "Descriptions", Flag: "descriptions", Type: "[]string", Required: false},
	{Name: "Durations", Flag: "durations", Type: "[]int64", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.SavingsPlanOfferingFilterElement", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OfferingIds", Flag: "offering-ids", Type: "[]string", Required: false},
	{Name: "Operations", Flag: "operations", Type: "[]string", Required: false},
	{Name: "PaymentOptions", Flag: "payment-options", Type: "[]types.SavingsPlanPaymentOption", Required: false},
	{Name: "PlanTypes", Flag: "plan-types", Type: "[]types.SavingsPlanType", Required: false},
	{Name: "ProductType", Flag: "product-type", Type: "types.SavingsPlanProductType", Required: false},
	{Name: "ServiceCodes", Flag: "service-codes", Type: "[]string", Required: false},
	{Name: "UsageTypes", Flag: "usage-types", Type: "[]string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_return_savings_plan = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SavingsPlanId", Flag: "savings-plan-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-savings-plan": {
			Name:   "create-savings-plan",
			Fields: fields_create_savings_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSavingsPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_savings_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSavingsPlan(ctx, input)
			},
		},
		"delete-queued-savings-plan": {
			Name:   "delete-queued-savings-plan",
			Fields: fields_delete_queued_savings_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueuedSavingsPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queued_savings_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueuedSavingsPlan(ctx, input)
			},
		},
		"describe-savings-plan-rates": {
			Name:   "describe-savings-plan-rates",
			Fields: fields_describe_savings_plan_rates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSavingsPlanRatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_savings_plan_rates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSavingsPlanRates(ctx, input)
			},
		},
		"describe-savings-plans": {
			Name:   "describe-savings-plans",
			Fields: fields_describe_savings_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSavingsPlansInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_savings_plans, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSavingsPlans(ctx, input)
			},
		},
		"describe-savings-plans-offering-rates": {
			Name:   "describe-savings-plans-offering-rates",
			Fields: fields_describe_savings_plans_offering_rates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSavingsPlansOfferingRatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_savings_plans_offering_rates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSavingsPlansOfferingRates(ctx, input)
			},
		},
		"describe-savings-plans-offerings": {
			Name:   "describe-savings-plans-offerings",
			Fields: fields_describe_savings_plans_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSavingsPlansOfferingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_savings_plans_offerings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSavingsPlansOfferings(ctx, input)
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
		"return-savings-plan": {
			Name:   "return-savings-plan",
			Fields: fields_return_savings_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReturnSavingsPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_return_savings_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReturnSavingsPlan(ctx, input)
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
	if err := leanruntime.Execute("savingsplans", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
