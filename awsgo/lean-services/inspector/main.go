package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/inspector"
)

var fields_add_attributes_to_findings = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.Attribute", Required: true},
	{Name: "FindingArns", Flag: "finding-arns", Type: "[]string", Required: true},
}

var fields_create_assessment_target = []leanruntime.Field{
	{Name: "AssessmentTargetName", Flag: "assessment-target-name", Type: "*string", Required: true},
	{Name: "ResourceGroupArn", Flag: "resource-group-arn", Type: "*string", Required: false},
}

var fields_create_assessment_template = []leanruntime.Field{
	{Name: "AssessmentTargetArn", Flag: "assessment-target-arn", Type: "*string", Required: true},
	{Name: "AssessmentTemplateName", Flag: "assessment-template-name", Type: "*string", Required: true},
	{Name: "DurationInSeconds", Flag: "duration-in-seconds", Type: "*int32", Required: true},
	{Name: "RulesPackageArns", Flag: "rules-package-arns", Type: "[]string", Required: true},
	{Name: "UserAttributesForFindings", Flag: "user-attributes-for-findings", Type: "[]types.Attribute", Required: false},
}

var fields_create_exclusions_preview = []leanruntime.Field{
	{Name: "AssessmentTemplateArn", Flag: "assessment-template-arn", Type: "*string", Required: true},
}

var fields_create_resource_group = []leanruntime.Field{
	{Name: "ResourceGroupTags", Flag: "resource-group-tags", Type: "[]types.ResourceGroupTag", Required: true},
}

var fields_delete_assessment_run = []leanruntime.Field{
	{Name: "AssessmentRunArn", Flag: "assessment-run-arn", Type: "*string", Required: true},
}

var fields_delete_assessment_target = []leanruntime.Field{
	{Name: "AssessmentTargetArn", Flag: "assessment-target-arn", Type: "*string", Required: true},
}

var fields_delete_assessment_template = []leanruntime.Field{
	{Name: "AssessmentTemplateArn", Flag: "assessment-template-arn", Type: "*string", Required: true},
}

var fields_describe_assessment_runs = []leanruntime.Field{
	{Name: "AssessmentRunArns", Flag: "assessment-run-arns", Type: "[]string", Required: true},
}

var fields_describe_assessment_targets = []leanruntime.Field{
	{Name: "AssessmentTargetArns", Flag: "assessment-target-arns", Type: "[]string", Required: true},
}

var fields_describe_assessment_templates = []leanruntime.Field{
	{Name: "AssessmentTemplateArns", Flag: "assessment-template-arns", Type: "[]string", Required: true},
}

var fields_describe_cross_account_access_role = []leanruntime.Field{}

var fields_describe_exclusions = []leanruntime.Field{
	{Name: "ExclusionArns", Flag: "exclusion-arns", Type: "[]string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
}

var fields_describe_findings = []leanruntime.Field{
	{Name: "FindingArns", Flag: "finding-arns", Type: "[]string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
}

var fields_describe_resource_groups = []leanruntime.Field{
	{Name: "ResourceGroupArns", Flag: "resource-group-arns", Type: "[]string", Required: true},
}

var fields_describe_rules_packages = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
	{Name: "RulesPackageArns", Flag: "rules-package-arns", Type: "[]string", Required: true},
}

var fields_get_assessment_report = []leanruntime.Field{
	{Name: "AssessmentRunArn", Flag: "assessment-run-arn", Type: "*string", Required: true},
	{Name: "ReportFileFormat", Flag: "report-file-format", Type: "types.ReportFileFormat", Required: true},
	{Name: "ReportType", Flag: "report-type", Type: "types.ReportType", Required: true},
}

var fields_get_exclusions_preview = []leanruntime.Field{
	{Name: "AssessmentTemplateArn", Flag: "assessment-template-arn", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PreviewToken", Flag: "preview-token", Type: "*string", Required: true},
}

var fields_get_telemetry_metadata = []leanruntime.Field{
	{Name: "AssessmentRunArn", Flag: "assessment-run-arn", Type: "*string", Required: true},
}

var fields_list_assessment_run_agents = []leanruntime.Field{
	{Name: "AssessmentRunArn", Flag: "assessment-run-arn", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.AgentFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assessment_runs = []leanruntime.Field{
	{Name: "AssessmentTemplateArns", Flag: "assessment-template-arns", Type: "[]string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.AssessmentRunFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assessment_targets = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.AssessmentTargetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assessment_templates = []leanruntime.Field{
	{Name: "AssessmentTargetArns", Flag: "assessment-target-arns", Type: "[]string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.AssessmentTemplateFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_subscriptions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_list_exclusions = []leanruntime.Field{
	{Name: "AssessmentRunArn", Flag: "assessment-run-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_findings = []leanruntime.Field{
	{Name: "AssessmentRunArns", Flag: "assessment-run-arns", Type: "[]string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.FindingFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rules_packages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_preview_agents = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PreviewAgentsArn", Flag: "preview-agents-arn", Type: "*string", Required: true},
}

var fields_register_cross_account_access_role = []leanruntime.Field{
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_remove_attributes_from_findings = []leanruntime.Field{
	{Name: "AttributeKeys", Flag: "attribute-keys", Type: "[]string", Required: true},
	{Name: "FindingArns", Flag: "finding-arns", Type: "[]string", Required: true},
}

var fields_set_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_assessment_run = []leanruntime.Field{
	{Name: "AssessmentRunName", Flag: "assessment-run-name", Type: "*string", Required: false},
	{Name: "AssessmentTemplateArn", Flag: "assessment-template-arn", Type: "*string", Required: true},
}

var fields_stop_assessment_run = []leanruntime.Field{
	{Name: "AssessmentRunArn", Flag: "assessment-run-arn", Type: "*string", Required: true},
	{Name: "StopAction", Flag: "stop-action", Type: "types.StopAction", Required: false},
}

var fields_subscribe_to_event = []leanruntime.Field{
	{Name: "Event", Flag: "event", Type: "types.InspectorEvent", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_unsubscribe_from_event = []leanruntime.Field{
	{Name: "Event", Flag: "event", Type: "types.InspectorEvent", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_update_assessment_target = []leanruntime.Field{
	{Name: "AssessmentTargetArn", Flag: "assessment-target-arn", Type: "*string", Required: true},
	{Name: "AssessmentTargetName", Flag: "assessment-target-name", Type: "*string", Required: true},
	{Name: "ResourceGroupArn", Flag: "resource-group-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-attributes-to-findings": {
			Name:   "add-attributes-to-findings",
			Fields: fields_add_attributes_to_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddAttributesToFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_attributes_to_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddAttributesToFindings(ctx, input)
			},
		},
		"create-assessment-target": {
			Name:   "create-assessment-target",
			Fields: fields_create_assessment_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssessmentTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assessment_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssessmentTarget(ctx, input)
			},
		},
		"create-assessment-template": {
			Name:   "create-assessment-template",
			Fields: fields_create_assessment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssessmentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assessment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssessmentTemplate(ctx, input)
			},
		},
		"create-exclusions-preview": {
			Name:   "create-exclusions-preview",
			Fields: fields_create_exclusions_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExclusionsPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_exclusions_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExclusionsPreview(ctx, input)
			},
		},
		"create-resource-group": {
			Name:   "create-resource-group",
			Fields: fields_create_resource_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceGroup(ctx, input)
			},
		},
		"delete-assessment-run": {
			Name:   "delete-assessment-run",
			Fields: fields_delete_assessment_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessmentRun(ctx, input)
			},
		},
		"delete-assessment-target": {
			Name:   "delete-assessment-target",
			Fields: fields_delete_assessment_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessmentTarget(ctx, input)
			},
		},
		"delete-assessment-template": {
			Name:   "delete-assessment-template",
			Fields: fields_delete_assessment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessmentTemplate(ctx, input)
			},
		},
		"describe-assessment-runs": {
			Name:   "describe-assessment-runs",
			Fields: fields_describe_assessment_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssessmentRunsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_assessment_runs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssessmentRuns(ctx, input)
			},
		},
		"describe-assessment-targets": {
			Name:   "describe-assessment-targets",
			Fields: fields_describe_assessment_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssessmentTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_assessment_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssessmentTargets(ctx, input)
			},
		},
		"describe-assessment-templates": {
			Name:   "describe-assessment-templates",
			Fields: fields_describe_assessment_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssessmentTemplatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_assessment_templates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssessmentTemplates(ctx, input)
			},
		},
		"describe-cross-account-access-role": {
			Name:   "describe-cross-account-access-role",
			Fields: fields_describe_cross_account_access_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCrossAccountAccessRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cross_account_access_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCrossAccountAccessRole(ctx, input)
			},
		},
		"describe-exclusions": {
			Name:   "describe-exclusions",
			Fields: fields_describe_exclusions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExclusionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_exclusions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExclusions(ctx, input)
			},
		},
		"describe-findings": {
			Name:   "describe-findings",
			Fields: fields_describe_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFindings(ctx, input)
			},
		},
		"describe-resource-groups": {
			Name:   "describe-resource-groups",
			Fields: fields_describe_resource_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourceGroups(ctx, input)
			},
		},
		"describe-rules-packages": {
			Name:   "describe-rules-packages",
			Fields: fields_describe_rules_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRulesPackagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rules_packages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRulesPackages(ctx, input)
			},
		},
		"get-assessment-report": {
			Name:   "get-assessment-report",
			Fields: fields_get_assessment_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssessmentReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assessment_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssessmentReport(ctx, input)
			},
		},
		"get-exclusions-preview": {
			Name:   "get-exclusions-preview",
			Fields: fields_get_exclusions_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExclusionsPreviewInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_exclusions_preview, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetExclusionsPreview(ctx, input)
				}
				var results []*svc.GetExclusionsPreviewOutput
				p := svc.NewGetExclusionsPreviewPaginator(client, input)
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
		"get-telemetry-metadata": {
			Name:   "get-telemetry-metadata",
			Fields: fields_get_telemetry_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTelemetryMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_telemetry_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTelemetryMetadata(ctx, input)
			},
		},
		"list-assessment-run-agents": {
			Name:   "list-assessment-run-agents",
			Fields: fields_list_assessment_run_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentRunAgentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_run_agents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentRunAgents(ctx, input)
				}
				var results []*svc.ListAssessmentRunAgentsOutput
				p := svc.NewListAssessmentRunAgentsPaginator(client, input)
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
		"list-assessment-runs": {
			Name:   "list-assessment-runs",
			Fields: fields_list_assessment_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentRuns(ctx, input)
				}
				var results []*svc.ListAssessmentRunsOutput
				p := svc.NewListAssessmentRunsPaginator(client, input)
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
		"list-assessment-targets": {
			Name:   "list-assessment-targets",
			Fields: fields_list_assessment_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentTargets(ctx, input)
				}
				var results []*svc.ListAssessmentTargetsOutput
				p := svc.NewListAssessmentTargetsPaginator(client, input)
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
		"list-assessment-templates": {
			Name:   "list-assessment-templates",
			Fields: fields_list_assessment_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentTemplates(ctx, input)
				}
				var results []*svc.ListAssessmentTemplatesOutput
				p := svc.NewListAssessmentTemplatesPaginator(client, input)
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
		"list-event-subscriptions": {
			Name:   "list-event-subscriptions",
			Fields: fields_list_event_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventSubscriptions(ctx, input)
				}
				var results []*svc.ListEventSubscriptionsOutput
				p := svc.NewListEventSubscriptionsPaginator(client, input)
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
		"list-exclusions": {
			Name:   "list-exclusions",
			Fields: fields_list_exclusions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExclusionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_exclusions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExclusions(ctx, input)
				}
				var results []*svc.ListExclusionsOutput
				p := svc.NewListExclusionsPaginator(client, input)
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
		"list-rules-packages": {
			Name:   "list-rules-packages",
			Fields: fields_list_rules_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rules_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRulesPackages(ctx, input)
				}
				var results []*svc.ListRulesPackagesOutput
				p := svc.NewListRulesPackagesPaginator(client, input)
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
		"preview-agents": {
			Name:   "preview-agents",
			Fields: fields_preview_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PreviewAgentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_preview_agents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.PreviewAgents(ctx, input)
				}
				var results []*svc.PreviewAgentsOutput
				p := svc.NewPreviewAgentsPaginator(client, input)
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
		"register-cross-account-access-role": {
			Name:   "register-cross-account-access-role",
			Fields: fields_register_cross_account_access_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterCrossAccountAccessRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_cross_account_access_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCrossAccountAccessRole(ctx, input)
			},
		},
		"remove-attributes-from-findings": {
			Name:   "remove-attributes-from-findings",
			Fields: fields_remove_attributes_from_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAttributesFromFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_attributes_from_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAttributesFromFindings(ctx, input)
			},
		},
		"set-tags-for-resource": {
			Name:   "set-tags-for-resource",
			Fields: fields_set_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTagsForResource(ctx, input)
			},
		},
		"start-assessment-run": {
			Name:   "start-assessment-run",
			Fields: fields_start_assessment_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAssessmentRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_assessment_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAssessmentRun(ctx, input)
			},
		},
		"stop-assessment-run": {
			Name:   "stop-assessment-run",
			Fields: fields_stop_assessment_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAssessmentRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_assessment_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAssessmentRun(ctx, input)
			},
		},
		"subscribe-to-event": {
			Name:   "subscribe-to-event",
			Fields: fields_subscribe_to_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubscribeToEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_subscribe_to_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubscribeToEvent(ctx, input)
			},
		},
		"unsubscribe-from-event": {
			Name:   "unsubscribe-from-event",
			Fields: fields_unsubscribe_from_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnsubscribeFromEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unsubscribe_from_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnsubscribeFromEvent(ctx, input)
			},
		},
		"update-assessment-target": {
			Name:   "update-assessment-target",
			Fields: fields_update_assessment_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessmentTarget(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("inspector", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
