package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

var fields_apply_archive_rule = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_cancel_policy_generation = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_check_access_not_granted = []leanruntime.Field{
	{Name: "Access", Flag: "access", Type: "[]types.Access", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.AccessCheckPolicyType", Required: true},
}

var fields_check_no_new_access = []leanruntime.Field{
	{Name: "ExistingPolicyDocument", Flag: "existing-policy-document", Type: "*string", Required: true},
	{Name: "NewPolicyDocument", Flag: "new-policy-document", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.AccessCheckPolicyType", Required: true},
}

var fields_check_no_public_access = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.AccessCheckResourceType", Required: true},
}

var fields_create_access_preview = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configurations", Flag: "configurations", Type: "map[string]types.Configuration", Required: true},
}

var fields_create_analyzer = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "ArchiveRules", Flag: "archive-rules", Type: "[]types.InlineArchiveRule", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.AnalyzerConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_create_archive_rule = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "map[string]types.Criterion", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_delete_analyzer = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_archive_rule = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_generate_finding_recommendation = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_access_preview = []leanruntime.Field{
	{Name: "AccessPreviewId", Flag: "access-preview-id", Type: "*string", Required: true},
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
}

var fields_get_analyzed_resource = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_analyzer = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
}

var fields_get_archive_rule = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_get_finding = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_finding_recommendation = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_finding_v2 = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_findings_statistics = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
}

var fields_get_generated_policy = []leanruntime.Field{
	{Name: "IncludeResourcePlaceholders", Flag: "include-resource-placeholders", Type: "*bool", Required: false},
	{Name: "IncludeServiceLevelTemplate", Flag: "include-service-level-template", Type: "*bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_list_access_preview_findings = []leanruntime.Field{
	{Name: "AccessPreviewId", Flag: "access-preview-id", Type: "*string", Required: true},
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "map[string]types.Criterion", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_previews = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_analyzed_resources = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: false},
}

var fields_list_analyzers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: false},
}

var fields_list_archive_rules = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_findings = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "map[string]types.Criterion", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SortCriteria", Required: false},
}

var fields_list_findings_v2 = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "map[string]types.Criterion", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SortCriteria", Required: false},
}

var fields_list_policy_generations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_policy_generation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CloudTrailDetails", Flag: "cloud-trail-details", Type: "*types.CloudTrailDetails", Required: false},
	{Name: "PolicyGenerationDetails", Flag: "policy-generation-details", Type: "*types.PolicyGenerationDetails", Required: true},
}

var fields_start_resource_scan = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceOwnerAccount", Flag: "resource-owner-account", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_analyzer = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.AnalyzerConfiguration", Required: false},
}

var fields_update_archive_rule = []leanruntime.Field{
	{Name: "AnalyzerName", Flag: "analyzer-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "map[string]types.Criterion", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
}

var fields_update_findings = []leanruntime.Field{
	{Name: "AnalyzerArn", Flag: "analyzer-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.FindingStatusUpdate", Required: true},
}

var fields_validate_policy = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
	{Name: "ValidatePolicyResourceType", Flag: "validate-policy-resource-type", Type: "types.ValidatePolicyResourceType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"apply-archive-rule": {
			Name:   "apply-archive-rule",
			Fields: fields_apply_archive_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplyArchiveRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_archive_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplyArchiveRule(ctx, input)
			},
		},
		"cancel-policy-generation": {
			Name:   "cancel-policy-generation",
			Fields: fields_cancel_policy_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelPolicyGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_policy_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelPolicyGeneration(ctx, input)
			},
		},
		"check-access-not-granted": {
			Name:   "check-access-not-granted",
			Fields: fields_check_access_not_granted,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckAccessNotGrantedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_access_not_granted, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckAccessNotGranted(ctx, input)
			},
		},
		"check-no-new-access": {
			Name:   "check-no-new-access",
			Fields: fields_check_no_new_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckNoNewAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_no_new_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckNoNewAccess(ctx, input)
			},
		},
		"check-no-public-access": {
			Name:   "check-no-public-access",
			Fields: fields_check_no_public_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckNoPublicAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_no_public_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckNoPublicAccess(ctx, input)
			},
		},
		"create-access-preview": {
			Name:   "create-access-preview",
			Fields: fields_create_access_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessPreview(ctx, input)
			},
		},
		"create-analyzer": {
			Name:   "create-analyzer",
			Fields: fields_create_analyzer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnalyzerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_analyzer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnalyzer(ctx, input)
			},
		},
		"create-archive-rule": {
			Name:   "create-archive-rule",
			Fields: fields_create_archive_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateArchiveRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_archive_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateArchiveRule(ctx, input)
			},
		},
		"delete-analyzer": {
			Name:   "delete-analyzer",
			Fields: fields_delete_analyzer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnalyzerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_analyzer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnalyzer(ctx, input)
			},
		},
		"delete-archive-rule": {
			Name:   "delete-archive-rule",
			Fields: fields_delete_archive_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteArchiveRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_archive_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteArchiveRule(ctx, input)
			},
		},
		"generate-finding-recommendation": {
			Name:   "generate-finding-recommendation",
			Fields: fields_generate_finding_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateFindingRecommendationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_finding_recommendation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateFindingRecommendation(ctx, input)
			},
		},
		"get-access-preview": {
			Name:   "get-access-preview",
			Fields: fields_get_access_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPreview(ctx, input)
			},
		},
		"get-analyzed-resource": {
			Name:   "get-analyzed-resource",
			Fields: fields_get_analyzed_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnalyzedResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_analyzed_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnalyzedResource(ctx, input)
			},
		},
		"get-analyzer": {
			Name:   "get-analyzer",
			Fields: fields_get_analyzer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnalyzerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_analyzer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnalyzer(ctx, input)
			},
		},
		"get-archive-rule": {
			Name:   "get-archive-rule",
			Fields: fields_get_archive_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchiveRule(ctx, input)
			},
		},
		"get-finding": {
			Name:   "get-finding",
			Fields: fields_get_finding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_finding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFinding(ctx, input)
			},
		},
		"get-finding-recommendation": {
			Name:   "get-finding-recommendation",
			Fields: fields_get_finding_recommendation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingRecommendationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_finding_recommendation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindingRecommendation(ctx, input)
				}
				var results []*svc.GetFindingRecommendationOutput
				p := svc.NewGetFindingRecommendationPaginator(client, input)
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
		"get-finding-v2": {
			Name:   "get-finding-v2",
			Fields: fields_get_finding_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_finding_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindingV2(ctx, input)
				}
				var results []*svc.GetFindingV2Output
				p := svc.NewGetFindingV2Paginator(client, input)
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
		"get-findings-statistics": {
			Name:   "get-findings-statistics",
			Fields: fields_get_findings_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_findings_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFindingsStatistics(ctx, input)
			},
		},
		"get-generated-policy": {
			Name:   "get-generated-policy",
			Fields: fields_get_generated_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGeneratedPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_generated_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGeneratedPolicy(ctx, input)
			},
		},
		"list-access-preview-findings": {
			Name:   "list-access-preview-findings",
			Fields: fields_list_access_preview_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPreviewFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_preview_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPreviewFindings(ctx, input)
				}
				var results []*svc.ListAccessPreviewFindingsOutput
				p := svc.NewListAccessPreviewFindingsPaginator(client, input)
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
		"list-access-previews": {
			Name:   "list-access-previews",
			Fields: fields_list_access_previews,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPreviewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_previews, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPreviews(ctx, input)
				}
				var results []*svc.ListAccessPreviewsOutput
				p := svc.NewListAccessPreviewsPaginator(client, input)
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
		"list-analyzed-resources": {
			Name:   "list-analyzed-resources",
			Fields: fields_list_analyzed_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalyzedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_analyzed_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnalyzedResources(ctx, input)
				}
				var results []*svc.ListAnalyzedResourcesOutput
				p := svc.NewListAnalyzedResourcesPaginator(client, input)
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
		"list-analyzers": {
			Name:   "list-analyzers",
			Fields: fields_list_analyzers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalyzersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_analyzers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnalyzers(ctx, input)
				}
				var results []*svc.ListAnalyzersOutput
				p := svc.NewListAnalyzersPaginator(client, input)
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
		"list-archive-rules": {
			Name:   "list-archive-rules",
			Fields: fields_list_archive_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArchiveRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_archive_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListArchiveRules(ctx, input)
				}
				var results []*svc.ListArchiveRulesOutput
				p := svc.NewListArchiveRulesPaginator(client, input)
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
		"list-findings": {
			Name:   "list-findings",
			Fields: fields_list_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindings(ctx, input)
				}
				var results []*svc.ListFindingsOutput
				p := svc.NewListFindingsPaginator(client, input)
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
		"list-findings-v2": {
			Name:   "list-findings-v2",
			Fields: fields_list_findings_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_findings_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindingsV2(ctx, input)
				}
				var results []*svc.ListFindingsV2Output
				p := svc.NewListFindingsV2Paginator(client, input)
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
		"list-policy-generations": {
			Name:   "list-policy-generations",
			Fields: fields_list_policy_generations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyGenerationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_generations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyGenerations(ctx, input)
				}
				var results []*svc.ListPolicyGenerationsOutput
				p := svc.NewListPolicyGenerationsPaginator(client, input)
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
		"start-policy-generation": {
			Name:   "start-policy-generation",
			Fields: fields_start_policy_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPolicyGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_policy_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPolicyGeneration(ctx, input)
			},
		},
		"start-resource-scan": {
			Name:   "start-resource-scan",
			Fields: fields_start_resource_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartResourceScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_resource_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartResourceScan(ctx, input)
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
		"update-analyzer": {
			Name:   "update-analyzer",
			Fields: fields_update_analyzer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnalyzerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_analyzer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnalyzer(ctx, input)
			},
		},
		"update-archive-rule": {
			Name:   "update-archive-rule",
			Fields: fields_update_archive_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateArchiveRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_archive_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateArchiveRule(ctx, input)
			},
		},
		"update-findings": {
			Name:   "update-findings",
			Fields: fields_update_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFindings(ctx, input)
			},
		},
		"validate-policy": {
			Name:   "validate-policy",
			Fields: fields_validate_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidatePolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_validate_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ValidatePolicy(ctx, input)
				}
				var results []*svc.ValidatePolicyOutput
				p := svc.NewValidatePolicyPaginator(client, input)
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
	}
	if err := leanruntime.Execute("accessanalyzer", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
