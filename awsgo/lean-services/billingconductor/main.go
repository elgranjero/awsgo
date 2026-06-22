package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/billingconductor"
)

var fields_associate_accounts = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_associate_pricing_rules = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "PricingRuleArns", Flag: "pricing-rule-arns", Type: "[]string", Required: true},
}

var fields_batch_associate_resources_to_custom_line_item = []leanruntime.Field{
	{Name: "BillingPeriodRange", Flag: "billing-period-range", Type: "*types.CustomLineItemBillingPeriodRange", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_batch_disassociate_resources_from_custom_line_item = []leanruntime.Field{
	{Name: "BillingPeriodRange", Flag: "billing-period-range", Type: "*types.CustomLineItemBillingPeriodRange", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_create_billing_group = []leanruntime.Field{
	{Name: "AccountGrouping", Flag: "account-grouping", Type: "*types.AccountGrouping", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComputationPreference", Flag: "computation-preference", Type: "*types.ComputationPreference", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PrimaryAccountId", Flag: "primary-account-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_custom_line_item = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "BillingGroupArn", Flag: "billing-group-arn", Type: "*string", Required: true},
	{Name: "BillingPeriodRange", Flag: "billing-period-range", Type: "*types.CustomLineItemBillingPeriodRange", Required: false},
	{Name: "ChargeDetails", Flag: "charge-details", Type: "*types.CustomLineItemChargeDetails", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComputationRule", Flag: "computation-rule", Type: "types.ComputationRuleEnum", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PresentationDetails", Flag: "presentation-details", Type: "*types.PresentationObject", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_pricing_plan = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PricingRuleArns", Flag: "pricing-rule-arns", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_pricing_rule = []leanruntime.Field{
	{Name: "BillingEntity", Flag: "billing-entity", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModifierPercentage", Flag: "modifier-percentage", Type: "*float64", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.PricingRuleScope", Required: true},
	{Name: "Service", Flag: "service", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tiering", Flag: "tiering", Type: "*types.CreateTieringInput", Required: false},
	{Name: "Type", Flag: "type", Type: "types.PricingRuleType", Required: true},
	{Name: "UsageType", Flag: "usage-type", Type: "*string", Required: false},
}

var fields_delete_billing_group = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_custom_line_item = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "BillingPeriodRange", Flag: "billing-period-range", Type: "*types.CustomLineItemBillingPeriodRange", Required: false},
}

var fields_delete_pricing_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_pricing_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_disassociate_accounts = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_disassociate_pricing_rules = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "PricingRuleArns", Flag: "pricing-rule-arns", Type: "[]string", Required: true},
}

var fields_get_billing_group_cost_report = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "BillingPeriodRange", Flag: "billing-period-range", Type: "*types.BillingPeriodRange", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupByAttributeName", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_account_associations = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListAccountAssociationsFilter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_billing_group_cost_reports = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListBillingGroupCostReportsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_billing_groups = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListBillingGroupsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_line_item_versions = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ListCustomLineItemVersionsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_line_items = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListCustomLineItemsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pricing_plans = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListPricingPlansFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pricing_plans_associated_with_pricing_rule = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PricingRuleArn", Flag: "pricing-rule-arn", Type: "*string", Required: true},
}

var fields_list_pricing_rules = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListPricingRulesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pricing_rules_associated_to_pricing_plan = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PricingPlanArn", Flag: "pricing-plan-arn", Type: "*string", Required: true},
}

var fields_list_resources_associated_to_custom_line_item = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListResourcesAssociatedToCustomLineItemFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_billing_group = []leanruntime.Field{
	{Name: "AccountGrouping", Flag: "account-grouping", Type: "*types.UpdateBillingGroupAccountGrouping", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ComputationPreference", Flag: "computation-preference", Type: "*types.ComputationPreference", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BillingGroupStatus", Required: false},
}

var fields_update_custom_line_item = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "BillingPeriodRange", Flag: "billing-period-range", Type: "*types.CustomLineItemBillingPeriodRange", Required: false},
	{Name: "ChargeDetails", Flag: "charge-details", Type: "*types.UpdateCustomLineItemChargeDetails", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_pricing_plan = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_pricing_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModifierPercentage", Flag: "modifier-percentage", Type: "*float64", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tiering", Flag: "tiering", Type: "*types.UpdateTieringInput", Required: false},
	{Name: "Type", Flag: "type", Type: "types.PricingRuleType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-accounts": {
			Name:   "associate-accounts",
			Fields: fields_associate_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAccounts(ctx, input)
			},
		},
		"associate-pricing-rules": {
			Name:   "associate-pricing-rules",
			Fields: fields_associate_pricing_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePricingRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_pricing_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePricingRules(ctx, input)
			},
		},
		"batch-associate-resources-to-custom-line-item": {
			Name:   "batch-associate-resources-to-custom-line-item",
			Fields: fields_batch_associate_resources_to_custom_line_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateResourcesToCustomLineItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_resources_to_custom_line_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateResourcesToCustomLineItem(ctx, input)
			},
		},
		"batch-disassociate-resources-from-custom-line-item": {
			Name:   "batch-disassociate-resources-from-custom-line-item",
			Fields: fields_batch_disassociate_resources_from_custom_line_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateResourcesFromCustomLineItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_resources_from_custom_line_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateResourcesFromCustomLineItem(ctx, input)
			},
		},
		"create-billing-group": {
			Name:   "create-billing-group",
			Fields: fields_create_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBillingGroup(ctx, input)
			},
		},
		"create-custom-line-item": {
			Name:   "create-custom-line-item",
			Fields: fields_create_custom_line_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomLineItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_line_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomLineItem(ctx, input)
			},
		},
		"create-pricing-plan": {
			Name:   "create-pricing-plan",
			Fields: fields_create_pricing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePricingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pricing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePricingPlan(ctx, input)
			},
		},
		"create-pricing-rule": {
			Name:   "create-pricing-rule",
			Fields: fields_create_pricing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePricingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pricing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePricingRule(ctx, input)
			},
		},
		"delete-billing-group": {
			Name:   "delete-billing-group",
			Fields: fields_delete_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBillingGroup(ctx, input)
			},
		},
		"delete-custom-line-item": {
			Name:   "delete-custom-line-item",
			Fields: fields_delete_custom_line_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomLineItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_line_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomLineItem(ctx, input)
			},
		},
		"delete-pricing-plan": {
			Name:   "delete-pricing-plan",
			Fields: fields_delete_pricing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePricingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pricing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePricingPlan(ctx, input)
			},
		},
		"delete-pricing-rule": {
			Name:   "delete-pricing-rule",
			Fields: fields_delete_pricing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePricingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pricing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePricingRule(ctx, input)
			},
		},
		"disassociate-accounts": {
			Name:   "disassociate-accounts",
			Fields: fields_disassociate_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAccounts(ctx, input)
			},
		},
		"disassociate-pricing-rules": {
			Name:   "disassociate-pricing-rules",
			Fields: fields_disassociate_pricing_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePricingRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_pricing_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePricingRules(ctx, input)
			},
		},
		"get-billing-group-cost-report": {
			Name:   "get-billing-group-cost-report",
			Fields: fields_get_billing_group_cost_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBillingGroupCostReportInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_billing_group_cost_report, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBillingGroupCostReport(ctx, input)
				}
				var results []*svc.GetBillingGroupCostReportOutput
				p := svc.NewGetBillingGroupCostReportPaginator(client, input)
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
		"list-account-associations": {
			Name:   "list-account-associations",
			Fields: fields_list_account_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAssociations(ctx, input)
				}
				var results []*svc.ListAccountAssociationsOutput
				p := svc.NewListAccountAssociationsPaginator(client, input)
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
		"list-billing-group-cost-reports": {
			Name:   "list-billing-group-cost-reports",
			Fields: fields_list_billing_group_cost_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillingGroupCostReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_billing_group_cost_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillingGroupCostReports(ctx, input)
				}
				var results []*svc.ListBillingGroupCostReportsOutput
				p := svc.NewListBillingGroupCostReportsPaginator(client, input)
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
		"list-billing-groups": {
			Name:   "list-billing-groups",
			Fields: fields_list_billing_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_billing_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillingGroups(ctx, input)
				}
				var results []*svc.ListBillingGroupsOutput
				p := svc.NewListBillingGroupsPaginator(client, input)
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
		"list-custom-line-item-versions": {
			Name:   "list-custom-line-item-versions",
			Fields: fields_list_custom_line_item_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomLineItemVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_line_item_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomLineItemVersions(ctx, input)
				}
				var results []*svc.ListCustomLineItemVersionsOutput
				p := svc.NewListCustomLineItemVersionsPaginator(client, input)
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
		"list-custom-line-items": {
			Name:   "list-custom-line-items",
			Fields: fields_list_custom_line_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomLineItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_line_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomLineItems(ctx, input)
				}
				var results []*svc.ListCustomLineItemsOutput
				p := svc.NewListCustomLineItemsPaginator(client, input)
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
		"list-pricing-plans": {
			Name:   "list-pricing-plans",
			Fields: fields_list_pricing_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPricingPlansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pricing_plans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPricingPlans(ctx, input)
				}
				var results []*svc.ListPricingPlansOutput
				p := svc.NewListPricingPlansPaginator(client, input)
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
		"list-pricing-plans-associated-with-pricing-rule": {
			Name:   "list-pricing-plans-associated-with-pricing-rule",
			Fields: fields_list_pricing_plans_associated_with_pricing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPricingPlansAssociatedWithPricingRuleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pricing_plans_associated_with_pricing_rule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPricingPlansAssociatedWithPricingRule(ctx, input)
				}
				var results []*svc.ListPricingPlansAssociatedWithPricingRuleOutput
				p := svc.NewListPricingPlansAssociatedWithPricingRulePaginator(client, input)
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
		"list-pricing-rules": {
			Name:   "list-pricing-rules",
			Fields: fields_list_pricing_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPricingRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pricing_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPricingRules(ctx, input)
				}
				var results []*svc.ListPricingRulesOutput
				p := svc.NewListPricingRulesPaginator(client, input)
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
		"list-pricing-rules-associated-to-pricing-plan": {
			Name:   "list-pricing-rules-associated-to-pricing-plan",
			Fields: fields_list_pricing_rules_associated_to_pricing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPricingRulesAssociatedToPricingPlanInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pricing_rules_associated_to_pricing_plan, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPricingRulesAssociatedToPricingPlan(ctx, input)
				}
				var results []*svc.ListPricingRulesAssociatedToPricingPlanOutput
				p := svc.NewListPricingRulesAssociatedToPricingPlanPaginator(client, input)
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
		"list-resources-associated-to-custom-line-item": {
			Name:   "list-resources-associated-to-custom-line-item",
			Fields: fields_list_resources_associated_to_custom_line_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesAssociatedToCustomLineItemInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources_associated_to_custom_line_item, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourcesAssociatedToCustomLineItem(ctx, input)
				}
				var results []*svc.ListResourcesAssociatedToCustomLineItemOutput
				p := svc.NewListResourcesAssociatedToCustomLineItemPaginator(client, input)
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
		"update-billing-group": {
			Name:   "update-billing-group",
			Fields: fields_update_billing_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBillingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_billing_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBillingGroup(ctx, input)
			},
		},
		"update-custom-line-item": {
			Name:   "update-custom-line-item",
			Fields: fields_update_custom_line_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomLineItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_line_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomLineItem(ctx, input)
			},
		},
		"update-pricing-plan": {
			Name:   "update-pricing-plan",
			Fields: fields_update_pricing_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePricingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pricing_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePricingPlan(ctx, input)
			},
		},
		"update-pricing-rule": {
			Name:   "update-pricing-rule",
			Fields: fields_update_pricing_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePricingRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pricing_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePricingRule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("billingconductor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
