package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bcmpricingcalculator"
)

var fields_batch_create_bill_scenario_commitment_modification = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CommitmentModifications", Flag: "commitment-modifications", Type: "[]types.BatchCreateBillScenarioCommitmentModificationEntry", Required: true},
}

var fields_batch_create_bill_scenario_usage_modification = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "UsageModifications", Flag: "usage-modifications", Type: "[]types.BatchCreateBillScenarioUsageModificationEntry", Required: true},
}

var fields_batch_create_workload_estimate_usage = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Usage", Flag: "usage", Type: "[]types.BatchCreateWorkloadEstimateUsageEntry", Required: true},
	{Name: "WorkloadEstimateId", Flag: "workload-estimate-id", Type: "*string", Required: true},
}

var fields_batch_delete_bill_scenario_commitment_modification = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_batch_delete_bill_scenario_usage_modification = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_batch_delete_workload_estimate_usage = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
	{Name: "WorkloadEstimateId", Flag: "workload-estimate-id", Type: "*string", Required: true},
}

var fields_batch_update_bill_scenario_commitment_modification = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "CommitmentModifications", Flag: "commitment-modifications", Type: "[]types.BatchUpdateBillScenarioCommitmentModificationEntry", Required: true},
}

var fields_batch_update_bill_scenario_usage_modification = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "UsageModifications", Flag: "usage-modifications", Type: "[]types.BatchUpdateBillScenarioUsageModificationEntry", Required: true},
}

var fields_batch_update_workload_estimate_usage = []leanruntime.Field{
	{Name: "Usage", Flag: "usage", Type: "[]types.BatchUpdateWorkloadEstimateUsageEntry", Required: true},
	{Name: "WorkloadEstimateId", Flag: "workload-estimate-id", Type: "*string", Required: true},
}

var fields_create_bill_estimate = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_bill_scenario = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CostCategoryGroupSharingPreferenceArn", Flag: "cost-category-group-sharing-preference-arn", Type: "*string", Required: false},
	{Name: "GroupSharingPreference", Flag: "group-sharing-preference", Type: "types.GroupSharingPreferenceEnum", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_workload_estimate = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RateType", Flag: "rate-type", Type: "types.WorkloadEstimateRateType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_bill_estimate = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_bill_scenario = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_workload_estimate = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_bill_estimate = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_bill_scenario = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_preferences = []leanruntime.Field{}

var fields_get_workload_estimate = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_list_bill_estimate_commitments = []leanruntime.Field{
	{Name: "BillEstimateId", Flag: "bill-estimate-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_estimate_input_commitment_modifications = []leanruntime.Field{
	{Name: "BillEstimateId", Flag: "bill-estimate-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_estimate_input_usage_modifications = []leanruntime.Field{
	{Name: "BillEstimateId", Flag: "bill-estimate-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListUsageFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_estimate_line_items = []leanruntime.Field{
	{Name: "BillEstimateId", Flag: "bill-estimate-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListBillEstimateLineItemsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_estimates = []leanruntime.Field{
	{Name: "CreatedAtFilter", Flag: "created-at-filter", Type: "*types.FilterTimestamp", Required: false},
	{Name: "ExpiresAtFilter", Flag: "expires-at-filter", Type: "*types.FilterTimestamp", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListBillEstimatesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_scenario_commitment_modifications = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_scenario_usage_modifications = []leanruntime.Field{
	{Name: "BillScenarioId", Flag: "bill-scenario-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListUsageFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bill_scenarios = []leanruntime.Field{
	{Name: "CreatedAtFilter", Flag: "created-at-filter", Type: "*types.FilterTimestamp", Required: false},
	{Name: "ExpiresAtFilter", Flag: "expires-at-filter", Type: "*types.FilterTimestamp", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListBillScenariosFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_workload_estimate_usage = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ListUsageFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadEstimateId", Flag: "workload-estimate-id", Type: "*string", Required: true},
}

var fields_list_workload_estimates = []leanruntime.Field{
	{Name: "CreatedAtFilter", Flag: "created-at-filter", Type: "*types.FilterTimestamp", Required: false},
	{Name: "ExpiresAtFilter", Flag: "expires-at-filter", Type: "*types.FilterTimestamp", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ListWorkloadEstimatesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_bill_estimate = []leanruntime.Field{
	{Name: "ExpiresAt", Flag: "expires-at", Type: "*time.Time", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_bill_scenario = []leanruntime.Field{
	{Name: "CostCategoryGroupSharingPreferenceArn", Flag: "cost-category-group-sharing-preference-arn", Type: "*string", Required: false},
	{Name: "ExpiresAt", Flag: "expires-at", Type: "*time.Time", Required: false},
	{Name: "GroupSharingPreference", Flag: "group-sharing-preference", Type: "types.GroupSharingPreferenceEnum", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_preferences = []leanruntime.Field{
	{Name: "ManagementAccountRateTypeSelections", Flag: "management-account-rate-type-selections", Type: "[]types.RateType", Required: false},
	{Name: "MemberAccountRateTypeSelections", Flag: "member-account-rate-type-selections", Type: "[]types.RateType", Required: false},
	{Name: "StandaloneAccountRateTypeSelections", Flag: "standalone-account-rate-type-selections", Type: "[]types.RateType", Required: false},
}

var fields_update_workload_estimate = []leanruntime.Field{
	{Name: "ExpiresAt", Flag: "expires-at", Type: "*time.Time", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-bill-scenario-commitment-modification": {
			Name:   "batch-create-bill-scenario-commitment-modification",
			Fields: fields_batch_create_bill_scenario_commitment_modification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateBillScenarioCommitmentModificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_bill_scenario_commitment_modification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateBillScenarioCommitmentModification(ctx, input)
			},
		},
		"batch-create-bill-scenario-usage-modification": {
			Name:   "batch-create-bill-scenario-usage-modification",
			Fields: fields_batch_create_bill_scenario_usage_modification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateBillScenarioUsageModificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_bill_scenario_usage_modification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateBillScenarioUsageModification(ctx, input)
			},
		},
		"batch-create-workload-estimate-usage": {
			Name:   "batch-create-workload-estimate-usage",
			Fields: fields_batch_create_workload_estimate_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateWorkloadEstimateUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_workload_estimate_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateWorkloadEstimateUsage(ctx, input)
			},
		},
		"batch-delete-bill-scenario-commitment-modification": {
			Name:   "batch-delete-bill-scenario-commitment-modification",
			Fields: fields_batch_delete_bill_scenario_commitment_modification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteBillScenarioCommitmentModificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_bill_scenario_commitment_modification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteBillScenarioCommitmentModification(ctx, input)
			},
		},
		"batch-delete-bill-scenario-usage-modification": {
			Name:   "batch-delete-bill-scenario-usage-modification",
			Fields: fields_batch_delete_bill_scenario_usage_modification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteBillScenarioUsageModificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_bill_scenario_usage_modification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteBillScenarioUsageModification(ctx, input)
			},
		},
		"batch-delete-workload-estimate-usage": {
			Name:   "batch-delete-workload-estimate-usage",
			Fields: fields_batch_delete_workload_estimate_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteWorkloadEstimateUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_workload_estimate_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteWorkloadEstimateUsage(ctx, input)
			},
		},
		"batch-update-bill-scenario-commitment-modification": {
			Name:   "batch-update-bill-scenario-commitment-modification",
			Fields: fields_batch_update_bill_scenario_commitment_modification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateBillScenarioCommitmentModificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_bill_scenario_commitment_modification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateBillScenarioCommitmentModification(ctx, input)
			},
		},
		"batch-update-bill-scenario-usage-modification": {
			Name:   "batch-update-bill-scenario-usage-modification",
			Fields: fields_batch_update_bill_scenario_usage_modification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateBillScenarioUsageModificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_bill_scenario_usage_modification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateBillScenarioUsageModification(ctx, input)
			},
		},
		"batch-update-workload-estimate-usage": {
			Name:   "batch-update-workload-estimate-usage",
			Fields: fields_batch_update_workload_estimate_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateWorkloadEstimateUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_workload_estimate_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateWorkloadEstimateUsage(ctx, input)
			},
		},
		"create-bill-estimate": {
			Name:   "create-bill-estimate",
			Fields: fields_create_bill_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBillEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bill_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBillEstimate(ctx, input)
			},
		},
		"create-bill-scenario": {
			Name:   "create-bill-scenario",
			Fields: fields_create_bill_scenario,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBillScenarioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bill_scenario, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBillScenario(ctx, input)
			},
		},
		"create-workload-estimate": {
			Name:   "create-workload-estimate",
			Fields: fields_create_workload_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkloadEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workload_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkloadEstimate(ctx, input)
			},
		},
		"delete-bill-estimate": {
			Name:   "delete-bill-estimate",
			Fields: fields_delete_bill_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBillEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bill_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBillEstimate(ctx, input)
			},
		},
		"delete-bill-scenario": {
			Name:   "delete-bill-scenario",
			Fields: fields_delete_bill_scenario,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBillScenarioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bill_scenario, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBillScenario(ctx, input)
			},
		},
		"delete-workload-estimate": {
			Name:   "delete-workload-estimate",
			Fields: fields_delete_workload_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkloadEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workload_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkloadEstimate(ctx, input)
			},
		},
		"get-bill-estimate": {
			Name:   "get-bill-estimate",
			Fields: fields_get_bill_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBillEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bill_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBillEstimate(ctx, input)
			},
		},
		"get-bill-scenario": {
			Name:   "get-bill-scenario",
			Fields: fields_get_bill_scenario,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBillScenarioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bill_scenario, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBillScenario(ctx, input)
			},
		},
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
		"get-workload-estimate": {
			Name:   "get-workload-estimate",
			Fields: fields_get_workload_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkloadEstimate(ctx, input)
			},
		},
		"list-bill-estimate-commitments": {
			Name:   "list-bill-estimate-commitments",
			Fields: fields_list_bill_estimate_commitments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillEstimateCommitmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_estimate_commitments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillEstimateCommitments(ctx, input)
				}
				var results []*svc.ListBillEstimateCommitmentsOutput
				p := svc.NewListBillEstimateCommitmentsPaginator(client, input)
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
		"list-bill-estimate-input-commitment-modifications": {
			Name:   "list-bill-estimate-input-commitment-modifications",
			Fields: fields_list_bill_estimate_input_commitment_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillEstimateInputCommitmentModificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_estimate_input_commitment_modifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillEstimateInputCommitmentModifications(ctx, input)
				}
				var results []*svc.ListBillEstimateInputCommitmentModificationsOutput
				p := svc.NewListBillEstimateInputCommitmentModificationsPaginator(client, input)
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
		"list-bill-estimate-input-usage-modifications": {
			Name:   "list-bill-estimate-input-usage-modifications",
			Fields: fields_list_bill_estimate_input_usage_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillEstimateInputUsageModificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_estimate_input_usage_modifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillEstimateInputUsageModifications(ctx, input)
				}
				var results []*svc.ListBillEstimateInputUsageModificationsOutput
				p := svc.NewListBillEstimateInputUsageModificationsPaginator(client, input)
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
		"list-bill-estimate-line-items": {
			Name:   "list-bill-estimate-line-items",
			Fields: fields_list_bill_estimate_line_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillEstimateLineItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_estimate_line_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillEstimateLineItems(ctx, input)
				}
				var results []*svc.ListBillEstimateLineItemsOutput
				p := svc.NewListBillEstimateLineItemsPaginator(client, input)
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
		"list-bill-estimates": {
			Name:   "list-bill-estimates",
			Fields: fields_list_bill_estimates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillEstimatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_estimates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillEstimates(ctx, input)
				}
				var results []*svc.ListBillEstimatesOutput
				p := svc.NewListBillEstimatesPaginator(client, input)
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
		"list-bill-scenario-commitment-modifications": {
			Name:   "list-bill-scenario-commitment-modifications",
			Fields: fields_list_bill_scenario_commitment_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillScenarioCommitmentModificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_scenario_commitment_modifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillScenarioCommitmentModifications(ctx, input)
				}
				var results []*svc.ListBillScenarioCommitmentModificationsOutput
				p := svc.NewListBillScenarioCommitmentModificationsPaginator(client, input)
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
		"list-bill-scenario-usage-modifications": {
			Name:   "list-bill-scenario-usage-modifications",
			Fields: fields_list_bill_scenario_usage_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillScenarioUsageModificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_scenario_usage_modifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillScenarioUsageModifications(ctx, input)
				}
				var results []*svc.ListBillScenarioUsageModificationsOutput
				p := svc.NewListBillScenarioUsageModificationsPaginator(client, input)
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
		"list-bill-scenarios": {
			Name:   "list-bill-scenarios",
			Fields: fields_list_bill_scenarios,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBillScenariosInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bill_scenarios, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBillScenarios(ctx, input)
				}
				var results []*svc.ListBillScenariosOutput
				p := svc.NewListBillScenariosPaginator(client, input)
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
		"list-workload-estimate-usage": {
			Name:   "list-workload-estimate-usage",
			Fields: fields_list_workload_estimate_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadEstimateUsageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workload_estimate_usage, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloadEstimateUsage(ctx, input)
				}
				var results []*svc.ListWorkloadEstimateUsageOutput
				p := svc.NewListWorkloadEstimateUsagePaginator(client, input)
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
		"list-workload-estimates": {
			Name:   "list-workload-estimates",
			Fields: fields_list_workload_estimates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadEstimatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workload_estimates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloadEstimates(ctx, input)
				}
				var results []*svc.ListWorkloadEstimatesOutput
				p := svc.NewListWorkloadEstimatesPaginator(client, input)
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
		"update-bill-estimate": {
			Name:   "update-bill-estimate",
			Fields: fields_update_bill_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBillEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bill_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBillEstimate(ctx, input)
			},
		},
		"update-bill-scenario": {
			Name:   "update-bill-scenario",
			Fields: fields_update_bill_scenario,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBillScenarioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bill_scenario, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBillScenario(ctx, input)
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
		"update-workload-estimate": {
			Name:   "update-workload-estimate",
			Fields: fields_update_workload_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkloadEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workload_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkloadEstimate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bcmpricingcalculator", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
