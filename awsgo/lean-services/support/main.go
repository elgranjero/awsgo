package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/support"
)

var fields_add_attachments_to_set = []leanruntime.Field{
	{Name: "AttachmentSetId", Flag: "attachment-set-id", Type: "*string", Required: false},
	{Name: "Attachments", Flag: "attachments", Type: "[]types.Attachment", Required: true},
}

var fields_add_communication_to_case = []leanruntime.Field{
	{Name: "AttachmentSetId", Flag: "attachment-set-id", Type: "*string", Required: false},
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: false},
	{Name: "CcEmailAddresses", Flag: "cc-email-addresses", Type: "[]string", Required: false},
	{Name: "CommunicationBody", Flag: "communication-body", Type: "*string", Required: true},
}

var fields_create_case = []leanruntime.Field{
	{Name: "AttachmentSetId", Flag: "attachment-set-id", Type: "*string", Required: false},
	{Name: "CategoryCode", Flag: "category-code", Type: "*string", Required: false},
	{Name: "CcEmailAddresses", Flag: "cc-email-addresses", Type: "[]string", Required: false},
	{Name: "CommunicationBody", Flag: "communication-body", Type: "*string", Required: true},
	{Name: "IssueType", Flag: "issue-type", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: false},
	{Name: "SeverityCode", Flag: "severity-code", Type: "*string", Required: false},
	{Name: "Subject", Flag: "subject", Type: "*string", Required: true},
}

var fields_describe_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_describe_cases = []leanruntime.Field{
	{Name: "AfterTime", Flag: "after-time", Type: "*string", Required: false},
	{Name: "BeforeTime", Flag: "before-time", Type: "*string", Required: false},
	{Name: "CaseIdList", Flag: "case-id-list", Type: "[]string", Required: false},
	{Name: "DisplayId", Flag: "display-id", Type: "*string", Required: false},
	{Name: "IncludeCommunications", Flag: "include-communications", Type: "*bool", Required: false},
	{Name: "IncludeResolvedCases", Flag: "include-resolved-cases", Type: "bool", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_communications = []leanruntime.Field{
	{Name: "AfterTime", Flag: "after-time", Type: "*string", Required: false},
	{Name: "BeforeTime", Flag: "before-time", Type: "*string", Required: false},
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_create_case_options = []leanruntime.Field{
	{Name: "CategoryCode", Flag: "category-code", Type: "*string", Required: true},
	{Name: "IssueType", Flag: "issue-type", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_describe_services = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "ServiceCodeList", Flag: "service-code-list", Type: "[]string", Required: false},
}

var fields_describe_severity_levels = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
}

var fields_describe_supported_languages = []leanruntime.Field{
	{Name: "CategoryCode", Flag: "category-code", Type: "*string", Required: true},
	{Name: "IssueType", Flag: "issue-type", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_describe_trusted_advisor_check_refresh_statuses = []leanruntime.Field{
	{Name: "CheckIds", Flag: "check-ids", Type: "[]*string", Required: true},
}

var fields_describe_trusted_advisor_check_result = []leanruntime.Field{
	{Name: "CheckId", Flag: "check-id", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
}

var fields_describe_trusted_advisor_check_summaries = []leanruntime.Field{
	{Name: "CheckIds", Flag: "check-ids", Type: "[]*string", Required: true},
}

var fields_describe_trusted_advisor_checks = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "*string", Required: true},
}

var fields_refresh_trusted_advisor_check = []leanruntime.Field{
	{Name: "CheckId", Flag: "check-id", Type: "*string", Required: true},
}

var fields_resolve_case = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-attachments-to-set": {
			Name:   "add-attachments-to-set",
			Fields: fields_add_attachments_to_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddAttachmentsToSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_attachments_to_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddAttachmentsToSet(ctx, input)
			},
		},
		"add-communication-to-case": {
			Name:   "add-communication-to-case",
			Fields: fields_add_communication_to_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddCommunicationToCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_communication_to_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddCommunicationToCase(ctx, input)
			},
		},
		"create-case": {
			Name:   "create-case",
			Fields: fields_create_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCase(ctx, input)
			},
		},
		"describe-attachment": {
			Name:   "describe-attachment",
			Fields: fields_describe_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAttachment(ctx, input)
			},
		},
		"describe-cases": {
			Name:   "describe-cases",
			Fields: fields_describe_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCases(ctx, input)
				}
				var results []*svc.DescribeCasesOutput
				p := svc.NewDescribeCasesPaginator(client, input)
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
		"describe-communications": {
			Name:   "describe-communications",
			Fields: fields_describe_communications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCommunicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_communications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCommunications(ctx, input)
				}
				var results []*svc.DescribeCommunicationsOutput
				p := svc.NewDescribeCommunicationsPaginator(client, input)
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
		"describe-create-case-options": {
			Name:   "describe-create-case-options",
			Fields: fields_describe_create_case_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCreateCaseOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_create_case_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCreateCaseOptions(ctx, input)
			},
		},
		"describe-services": {
			Name:   "describe-services",
			Fields: fields_describe_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_services, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServices(ctx, input)
			},
		},
		"describe-severity-levels": {
			Name:   "describe-severity-levels",
			Fields: fields_describe_severity_levels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSeverityLevelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_severity_levels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSeverityLevels(ctx, input)
			},
		},
		"describe-supported-languages": {
			Name:   "describe-supported-languages",
			Fields: fields_describe_supported_languages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSupportedLanguagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_supported_languages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSupportedLanguages(ctx, input)
			},
		},
		"describe-trusted-advisor-check-refresh-statuses": {
			Name:   "describe-trusted-advisor-check-refresh-statuses",
			Fields: fields_describe_trusted_advisor_check_refresh_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustedAdvisorCheckRefreshStatusesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trusted_advisor_check_refresh_statuses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrustedAdvisorCheckRefreshStatuses(ctx, input)
			},
		},
		"describe-trusted-advisor-check-result": {
			Name:   "describe-trusted-advisor-check-result",
			Fields: fields_describe_trusted_advisor_check_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustedAdvisorCheckResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trusted_advisor_check_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrustedAdvisorCheckResult(ctx, input)
			},
		},
		"describe-trusted-advisor-check-summaries": {
			Name:   "describe-trusted-advisor-check-summaries",
			Fields: fields_describe_trusted_advisor_check_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustedAdvisorCheckSummariesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trusted_advisor_check_summaries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrustedAdvisorCheckSummaries(ctx, input)
			},
		},
		"describe-trusted-advisor-checks": {
			Name:   "describe-trusted-advisor-checks",
			Fields: fields_describe_trusted_advisor_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustedAdvisorChecksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trusted_advisor_checks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrustedAdvisorChecks(ctx, input)
			},
		},
		"refresh-trusted-advisor-check": {
			Name:   "refresh-trusted-advisor-check",
			Fields: fields_refresh_trusted_advisor_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RefreshTrustedAdvisorCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_refresh_trusted_advisor_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RefreshTrustedAdvisorCheck(ctx, input)
			},
		},
		"resolve-case": {
			Name:   "resolve-case",
			Fields: fields_resolve_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResolveCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resolve_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResolveCase(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("support", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
