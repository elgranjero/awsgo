package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/observabilityadmin"
)

var fields_create_centralization_rule_for_organization = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.CentralizationRule", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_s3_table_integration = []leanruntime.Field{
	{Name: "Encryption", Flag: "encryption", Type: "*types.Encryption", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_telemetry_pipeline = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.TelemetryPipelineConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_telemetry_rule = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.TelemetryRule", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_telemetry_rule_for_organization = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.TelemetryRule", Required: true},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_centralization_rule_for_organization = []leanruntime.Field{
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_delete_s3_table_integration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_telemetry_pipeline = []leanruntime.Field{
	{Name: "PipelineIdentifier", Flag: "pipeline-identifier", Type: "*string", Required: true},
}

var fields_delete_telemetry_rule = []leanruntime.Field{
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_delete_telemetry_rule_for_organization = []leanruntime.Field{
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_get_centralization_rule_for_organization = []leanruntime.Field{
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_get_s3_table_integration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_telemetry_enrichment_status = []leanruntime.Field{}

var fields_get_telemetry_evaluation_status = []leanruntime.Field{}

var fields_get_telemetry_evaluation_status_for_organization = []leanruntime.Field{}

var fields_get_telemetry_pipeline = []leanruntime.Field{
	{Name: "PipelineIdentifier", Flag: "pipeline-identifier", Type: "*string", Required: true},
}

var fields_get_telemetry_rule = []leanruntime.Field{
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_get_telemetry_rule_for_organization = []leanruntime.Field{
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_list_centralization_rules_for_organization = []leanruntime.Field{
	{Name: "AllRegions", Flag: "all-regions", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RuleNamePrefix", Flag: "rule-name-prefix", Type: "*string", Required: false},
}

var fields_list_resource_telemetry = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifierPrefix", Flag: "resource-identifier-prefix", Type: "*string", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "map[string]string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceType", Required: false},
	{Name: "TelemetryConfigurationState", Flag: "telemetry-configuration-state", Type: "map[string]types.TelemetryState", Required: false},
}

var fields_list_resource_telemetry_for_organization = []leanruntime.Field{
	{Name: "AccountIdentifiers", Flag: "account-identifiers", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifierPrefix", Flag: "resource-identifier-prefix", Type: "*string", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "map[string]string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceType", Required: false},
	{Name: "TelemetryConfigurationState", Flag: "telemetry-configuration-state", Type: "map[string]types.TelemetryState", Required: false},
}

var fields_list_s3_table_integrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_telemetry_pipelines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_telemetry_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RuleNamePrefix", Flag: "rule-name-prefix", Type: "*string", Required: false},
}

var fields_list_telemetry_rules_for_organization = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RuleNamePrefix", Flag: "rule-name-prefix", Type: "*string", Required: false},
	{Name: "SourceAccountIds", Flag: "source-account-ids", Type: "[]string", Required: false},
	{Name: "SourceOrganizationUnitIds", Flag: "source-organization-unit-ids", Type: "[]string", Required: false},
}

var fields_start_telemetry_enrichment = []leanruntime.Field{}

var fields_start_telemetry_evaluation = []leanruntime.Field{}

var fields_start_telemetry_evaluation_for_organization = []leanruntime.Field{}

var fields_stop_telemetry_enrichment = []leanruntime.Field{}

var fields_stop_telemetry_evaluation = []leanruntime.Field{}

var fields_stop_telemetry_evaluation_for_organization = []leanruntime.Field{}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_test_telemetry_pipeline = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.TelemetryPipelineConfiguration", Required: true},
	{Name: "Records", Flag: "records", Type: "[]types.Record", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_centralization_rule_for_organization = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.CentralizationRule", Required: true},
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_update_telemetry_pipeline = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.TelemetryPipelineConfiguration", Required: true},
	{Name: "PipelineIdentifier", Flag: "pipeline-identifier", Type: "*string", Required: true},
}

var fields_update_telemetry_rule = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.TelemetryRule", Required: true},
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_update_telemetry_rule_for_organization = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.TelemetryRule", Required: true},
	{Name: "RuleIdentifier", Flag: "rule-identifier", Type: "*string", Required: true},
}

var fields_validate_telemetry_pipeline_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.TelemetryPipelineConfiguration", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-centralization-rule-for-organization": {
			Name:   "create-centralization-rule-for-organization",
			Fields: fields_create_centralization_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCentralizationRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_centralization_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCentralizationRuleForOrganization(ctx, input)
			},
		},
		"create-s3-table-integration": {
			Name:   "create-s3-table-integration",
			Fields: fields_create_s3_table_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateS3TableIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_s3_table_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateS3TableIntegration(ctx, input)
			},
		},
		"create-telemetry-pipeline": {
			Name:   "create-telemetry-pipeline",
			Fields: fields_create_telemetry_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTelemetryPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_telemetry_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTelemetryPipeline(ctx, input)
			},
		},
		"create-telemetry-rule": {
			Name:   "create-telemetry-rule",
			Fields: fields_create_telemetry_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTelemetryRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_telemetry_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTelemetryRule(ctx, input)
			},
		},
		"create-telemetry-rule-for-organization": {
			Name:   "create-telemetry-rule-for-organization",
			Fields: fields_create_telemetry_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTelemetryRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_telemetry_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTelemetryRuleForOrganization(ctx, input)
			},
		},
		"delete-centralization-rule-for-organization": {
			Name:   "delete-centralization-rule-for-organization",
			Fields: fields_delete_centralization_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCentralizationRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_centralization_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCentralizationRuleForOrganization(ctx, input)
			},
		},
		"delete-s3-table-integration": {
			Name:   "delete-s3-table-integration",
			Fields: fields_delete_s3_table_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteS3TableIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_s3_table_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteS3TableIntegration(ctx, input)
			},
		},
		"delete-telemetry-pipeline": {
			Name:   "delete-telemetry-pipeline",
			Fields: fields_delete_telemetry_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTelemetryPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_telemetry_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTelemetryPipeline(ctx, input)
			},
		},
		"delete-telemetry-rule": {
			Name:   "delete-telemetry-rule",
			Fields: fields_delete_telemetry_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTelemetryRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_telemetry_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTelemetryRule(ctx, input)
			},
		},
		"delete-telemetry-rule-for-organization": {
			Name:   "delete-telemetry-rule-for-organization",
			Fields: fields_delete_telemetry_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTelemetryRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_telemetry_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTelemetryRuleForOrganization(ctx, input)
			},
		},
		"get-centralization-rule-for-organization": {
			Name:   "get-centralization-rule-for-organization",
			Fields: fields_get_centralization_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCentralizationRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_centralization_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCentralizationRuleForOrganization(ctx, input)
			},
		},
		"get-s3-table-integration": {
			Name:   "get-s3-table-integration",
			Fields: fields_get_s3_table_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetS3TableIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_s3_table_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetS3TableIntegration(ctx, input)
			},
		},
		"get-telemetry-enrichment-status": {
			Name:   "get-telemetry-enrichment-status",
			Fields: fields_get_telemetry_enrichment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryEnrichmentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_enrichment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryEnrichmentStatus(ctx, input)
			},
		},
		"get-telemetry-evaluation-status": {
			Name:   "get-telemetry-evaluation-status",
			Fields: fields_get_telemetry_evaluation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryEvaluationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_evaluation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryEvaluationStatus(ctx, input)
			},
		},
		"get-telemetry-evaluation-status-for-organization": {
			Name:   "get-telemetry-evaluation-status-for-organization",
			Fields: fields_get_telemetry_evaluation_status_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryEvaluationStatusForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_evaluation_status_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryEvaluationStatusForOrganization(ctx, input)
			},
		},
		"get-telemetry-pipeline": {
			Name:   "get-telemetry-pipeline",
			Fields: fields_get_telemetry_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryPipeline(ctx, input)
			},
		},
		"get-telemetry-rule": {
			Name:   "get-telemetry-rule",
			Fields: fields_get_telemetry_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryRule(ctx, input)
			},
		},
		"get-telemetry-rule-for-organization": {
			Name:   "get-telemetry-rule-for-organization",
			Fields: fields_get_telemetry_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryRuleForOrganization(ctx, input)
			},
		},
		"list-centralization-rules-for-organization": {
			Name:   "list-centralization-rules-for-organization",
			Fields: fields_list_centralization_rules_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCentralizationRulesForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_centralization_rules_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCentralizationRulesForOrganization(ctx, input)
				}
				var results []*svc.ListCentralizationRulesForOrganizationOutput
				p := svc.NewListCentralizationRulesForOrganizationPaginator(client, input)
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
		"list-resource-telemetry": {
			Name:   "list-resource-telemetry",
			Fields: fields_list_resource_telemetry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceTelemetryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_telemetry, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceTelemetry(ctx, input)
				}
				var results []*svc.ListResourceTelemetryOutput
				p := svc.NewListResourceTelemetryPaginator(client, input)
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
		"list-resource-telemetry-for-organization": {
			Name:   "list-resource-telemetry-for-organization",
			Fields: fields_list_resource_telemetry_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceTelemetryForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_telemetry_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceTelemetryForOrganization(ctx, input)
				}
				var results []*svc.ListResourceTelemetryForOrganizationOutput
				p := svc.NewListResourceTelemetryForOrganizationPaginator(client, input)
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
		"list-s3-table-integrations": {
			Name:   "list-s3-table-integrations",
			Fields: fields_list_s3_table_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListS3TableIntegrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_s3_table_integrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListS3TableIntegrations(ctx, input)
				}
				var results []*svc.ListS3TableIntegrationsOutput
				p := svc.NewListS3TableIntegrationsPaginator(client, input)
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
		"list-telemetry-pipelines": {
			Name:   "list-telemetry-pipelines",
			Fields: fields_list_telemetry_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTelemetryPipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_telemetry_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTelemetryPipelines(ctx, input)
				}
				var results []*svc.ListTelemetryPipelinesOutput
				p := svc.NewListTelemetryPipelinesPaginator(client, input)
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
		"list-telemetry-rules": {
			Name:   "list-telemetry-rules",
			Fields: fields_list_telemetry_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTelemetryRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_telemetry_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTelemetryRules(ctx, input)
				}
				var results []*svc.ListTelemetryRulesOutput
				p := svc.NewListTelemetryRulesPaginator(client, input)
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
		"list-telemetry-rules-for-organization": {
			Name:   "list-telemetry-rules-for-organization",
			Fields: fields_list_telemetry_rules_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTelemetryRulesForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_telemetry_rules_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTelemetryRulesForOrganization(ctx, input)
				}
				var results []*svc.ListTelemetryRulesForOrganizationOutput
				p := svc.NewListTelemetryRulesForOrganizationPaginator(client, input)
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
		"start-telemetry-enrichment": {
			Name:   "start-telemetry-enrichment",
			Fields: fields_start_telemetry_enrichment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTelemetryEnrichmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_telemetry_enrichment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTelemetryEnrichment(ctx, input)
			},
		},
		"start-telemetry-evaluation": {
			Name:   "start-telemetry-evaluation",
			Fields: fields_start_telemetry_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTelemetryEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_telemetry_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTelemetryEvaluation(ctx, input)
			},
		},
		"start-telemetry-evaluation-for-organization": {
			Name:   "start-telemetry-evaluation-for-organization",
			Fields: fields_start_telemetry_evaluation_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTelemetryEvaluationForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_telemetry_evaluation_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTelemetryEvaluationForOrganization(ctx, input)
			},
		},
		"stop-telemetry-enrichment": {
			Name:   "stop-telemetry-enrichment",
			Fields: fields_stop_telemetry_enrichment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTelemetryEnrichmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_telemetry_enrichment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTelemetryEnrichment(ctx, input)
			},
		},
		"stop-telemetry-evaluation": {
			Name:   "stop-telemetry-evaluation",
			Fields: fields_stop_telemetry_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTelemetryEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_telemetry_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTelemetryEvaluation(ctx, input)
			},
		},
		"stop-telemetry-evaluation-for-organization": {
			Name:   "stop-telemetry-evaluation-for-organization",
			Fields: fields_stop_telemetry_evaluation_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTelemetryEvaluationForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_telemetry_evaluation_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTelemetryEvaluationForOrganization(ctx, input)
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
		"test-telemetry-pipeline": {
			Name:   "test-telemetry-pipeline",
			Fields: fields_test_telemetry_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestTelemetryPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_telemetry_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestTelemetryPipeline(ctx, input)
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
		"update-centralization-rule-for-organization": {
			Name:   "update-centralization-rule-for-organization",
			Fields: fields_update_centralization_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCentralizationRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_centralization_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCentralizationRuleForOrganization(ctx, input)
			},
		},
		"update-telemetry-pipeline": {
			Name:   "update-telemetry-pipeline",
			Fields: fields_update_telemetry_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTelemetryPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_telemetry_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTelemetryPipeline(ctx, input)
			},
		},
		"update-telemetry-rule": {
			Name:   "update-telemetry-rule",
			Fields: fields_update_telemetry_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTelemetryRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_telemetry_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTelemetryRule(ctx, input)
			},
		},
		"update-telemetry-rule-for-organization": {
			Name:   "update-telemetry-rule-for-organization",
			Fields: fields_update_telemetry_rule_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTelemetryRuleForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_telemetry_rule_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTelemetryRuleForOrganization(ctx, input)
			},
		},
		"validate-telemetry-pipeline-configuration": {
			Name:   "validate-telemetry-pipeline-configuration",
			Fields: fields_validate_telemetry_pipeline_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateTelemetryPipelineConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_telemetry_pipeline_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateTelemetryPipelineConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("observabilityadmin", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
