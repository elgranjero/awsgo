package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/artifact"
)

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_report = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
	{Name: "ReportVersion", Flag: "report-version", Type: "*int64", Required: false},
	{Name: "TermToken", Flag: "term-token", Type: "*string", Required: true},
}

var fields_get_report_metadata = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
	{Name: "ReportVersion", Flag: "report-version", Type: "*int64", Required: false},
}

var fields_get_term_for_report = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
	{Name: "ReportVersion", Flag: "report-version", Type: "*int64", Required: false},
}

var fields_list_customer_agreements = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_report_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_list_reports = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_account_settings = []leanruntime.Field{
	{Name: "NotificationSubscriptionStatus", Flag: "notification-subscription-status", Type: "types.NotificationSubscriptionStatus", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
			},
		},
		"get-report": {
			Name:   "get-report",
			Fields: fields_get_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReport(ctx, input)
			},
		},
		"get-report-metadata": {
			Name:   "get-report-metadata",
			Fields: fields_get_report_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReportMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_report_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReportMetadata(ctx, input)
			},
		},
		"get-term-for-report": {
			Name:   "get-term-for-report",
			Fields: fields_get_term_for_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTermForReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_term_for_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTermForReport(ctx, input)
			},
		},
		"list-customer-agreements": {
			Name:   "list-customer-agreements",
			Fields: fields_list_customer_agreements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomerAgreementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_customer_agreements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomerAgreements(ctx, input)
				}
				var results []*svc.ListCustomerAgreementsOutput
				p := svc.NewListCustomerAgreementsPaginator(client, input)
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
		"list-report-versions": {
			Name:   "list-report-versions",
			Fields: fields_list_report_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_report_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReportVersions(ctx, input)
				}
				var results []*svc.ListReportVersionsOutput
				p := svc.NewListReportVersionsPaginator(client, input)
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
		"list-reports": {
			Name:   "list-reports",
			Fields: fields_list_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReports(ctx, input)
				}
				var results []*svc.ListReportsOutput
				p := svc.NewListReportsPaginator(client, input)
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
		"put-account-settings": {
			Name:   "put-account-settings",
			Fields: fields_put_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("artifact", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
