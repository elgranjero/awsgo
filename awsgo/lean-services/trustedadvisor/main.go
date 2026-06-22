package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/trustedadvisor"
)

var fields_batch_update_recommendation_resource_exclusion = []leanruntime.Field{
	{Name: "RecommendationResourceExclusions", Flag: "recommendation-resource-exclusions", Type: "[]types.RecommendationResourceExclusion", Required: true},
}

var fields_get_organization_recommendation = []leanruntime.Field{
	{Name: "OrganizationRecommendationIdentifier", Flag: "organization-recommendation-identifier", Type: "*string", Required: true},
}

var fields_get_recommendation = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "types.RecommendationLanguage", Required: false},
	{Name: "RecommendationIdentifier", Flag: "recommendation-identifier", Type: "*string", Required: true},
}

var fields_list_checks = []leanruntime.Field{
	{Name: "AwsService", Flag: "aws-service", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "types.RecommendationLanguage", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Pillar", Flag: "pillar", Type: "types.RecommendationPillar", Required: false},
	{Name: "Source", Flag: "source", Type: "types.RecommendationSource", Required: false},
}

var fields_list_organization_recommendation_accounts = []leanruntime.Field{
	{Name: "AffectedAccountId", Flag: "affected-account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationRecommendationIdentifier", Flag: "organization-recommendation-identifier", Type: "*string", Required: true},
}

var fields_list_organization_recommendation_resources = []leanruntime.Field{
	{Name: "AffectedAccountId", Flag: "affected-account-id", Type: "*string", Required: false},
	{Name: "ExclusionStatus", Flag: "exclusion-status", Type: "types.ExclusionStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationRecommendationIdentifier", Flag: "organization-recommendation-identifier", Type: "*string", Required: true},
	{Name: "RegionCode", Flag: "region-code", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ResourceStatus", Required: false},
}

var fields_list_organization_recommendations = []leanruntime.Field{
	{Name: "AfterLastUpdatedAt", Flag: "after-last-updated-at", Type: "*time.Time", Required: false},
	{Name: "AwsService", Flag: "aws-service", Type: "*string", Required: false},
	{Name: "BeforeLastUpdatedAt", Flag: "before-last-updated-at", Type: "*time.Time", Required: false},
	{Name: "CheckIdentifier", Flag: "check-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Pillar", Flag: "pillar", Type: "types.RecommendationPillar", Required: false},
	{Name: "Source", Flag: "source", Type: "types.RecommendationSource", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RecommendationStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RecommendationType", Required: false},
}

var fields_list_recommendation_resources = []leanruntime.Field{
	{Name: "ExclusionStatus", Flag: "exclusion-status", Type: "types.ExclusionStatus", Required: false},
	{Name: "Language", Flag: "language", Type: "types.RecommendationLanguage", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecommendationIdentifier", Flag: "recommendation-identifier", Type: "*string", Required: true},
	{Name: "RegionCode", Flag: "region-code", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ResourceStatus", Required: false},
}

var fields_list_recommendations = []leanruntime.Field{
	{Name: "AfterLastUpdatedAt", Flag: "after-last-updated-at", Type: "*time.Time", Required: false},
	{Name: "AwsService", Flag: "aws-service", Type: "*string", Required: false},
	{Name: "BeforeLastUpdatedAt", Flag: "before-last-updated-at", Type: "*time.Time", Required: false},
	{Name: "CheckIdentifier", Flag: "check-identifier", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "types.RecommendationLanguage", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Pillar", Flag: "pillar", Type: "types.RecommendationPillar", Required: false},
	{Name: "Source", Flag: "source", Type: "types.RecommendationSource", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RecommendationStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RecommendationType", Required: false},
}

var fields_update_organization_recommendation_lifecycle = []leanruntime.Field{
	{Name: "LifecycleStage", Flag: "lifecycle-stage", Type: "types.UpdateRecommendationLifecycleStage", Required: true},
	{Name: "OrganizationRecommendationIdentifier", Flag: "organization-recommendation-identifier", Type: "*string", Required: true},
	{Name: "UpdateReason", Flag: "update-reason", Type: "*string", Required: false},
	{Name: "UpdateReasonCode", Flag: "update-reason-code", Type: "types.UpdateRecommendationLifecycleStageReasonCode", Required: false},
}

var fields_update_recommendation_lifecycle = []leanruntime.Field{
	{Name: "LifecycleStage", Flag: "lifecycle-stage", Type: "types.UpdateRecommendationLifecycleStage", Required: true},
	{Name: "RecommendationIdentifier", Flag: "recommendation-identifier", Type: "*string", Required: true},
	{Name: "UpdateReason", Flag: "update-reason", Type: "*string", Required: false},
	{Name: "UpdateReasonCode", Flag: "update-reason-code", Type: "types.UpdateRecommendationLifecycleStageReasonCode", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-update-recommendation-resource-exclusion": {
			Name:   "batch-update-recommendation-resource-exclusion",
			Fields: fields_batch_update_recommendation_resource_exclusion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateRecommendationResourceExclusionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_recommendation_resource_exclusion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateRecommendationResourceExclusion(ctx, input)
			},
		},
		"get-organization-recommendation": {
			Name:   "get-organization-recommendation",
			Fields: fields_get_organization_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_organization_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOrganizationRecommendation(ctx, input)
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
		"list-checks": {
			Name:   "list-checks",
			Fields: fields_list_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChecksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_checks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChecks(ctx, input)
				}
				var results []*svc.ListChecksOutput
				p := svc.NewListChecksPaginator(client, input)
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
		"list-organization-recommendation-accounts": {
			Name:   "list-organization-recommendation-accounts",
			Fields: fields_list_organization_recommendation_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationRecommendationAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_recommendation_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationRecommendationAccounts(ctx, input)
				}
				var results []*svc.ListOrganizationRecommendationAccountsOutput
				p := svc.NewListOrganizationRecommendationAccountsPaginator(client, input)
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
		"list-organization-recommendation-resources": {
			Name:   "list-organization-recommendation-resources",
			Fields: fields_list_organization_recommendation_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationRecommendationResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_recommendation_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationRecommendationResources(ctx, input)
				}
				var results []*svc.ListOrganizationRecommendationResourcesOutput
				p := svc.NewListOrganizationRecommendationResourcesPaginator(client, input)
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
		"list-organization-recommendations": {
			Name:   "list-organization-recommendations",
			Fields: fields_list_organization_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationRecommendations(ctx, input)
				}
				var results []*svc.ListOrganizationRecommendationsOutput
				p := svc.NewListOrganizationRecommendationsPaginator(client, input)
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
		"list-recommendation-resources": {
			Name:   "list-recommendation-resources",
			Fields: fields_list_recommendation_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendation_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendationResources(ctx, input)
				}
				var results []*svc.ListRecommendationResourcesOutput
				p := svc.NewListRecommendationResourcesPaginator(client, input)
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
		"update-organization-recommendation-lifecycle": {
			Name:   "update-organization-recommendation-lifecycle",
			Fields: fields_update_organization_recommendation_lifecycle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOrganizationRecommendationLifecycleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_organization_recommendation_lifecycle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOrganizationRecommendationLifecycle(ctx, input)
			},
		},
		"update-recommendation-lifecycle": {
			Name:   "update-recommendation-lifecycle",
			Fields: fields_update_recommendation_lifecycle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecommendationLifecycleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recommendation_lifecycle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecommendationLifecycle(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("trustedadvisor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
