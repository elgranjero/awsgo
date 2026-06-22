package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codegurusecurity"
)

var fields_batch_get_findings = []leanruntime.Field{
	{Name: "FindingIdentifiers", Flag: "finding-identifiers", Type: "[]types.FindingIdentifier", Required: true},
}

var fields_create_scan = []leanruntime.Field{
	{Name: "AnalysisType", Flag: "analysis-type", Type: "types.AnalysisType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "types.ResourceId", Required: true},
	{Name: "ScanName", Flag: "scan-name", Type: "*string", Required: true},
	{Name: "ScanType", Flag: "scan-type", Type: "types.ScanType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_upload_url = []leanruntime.Field{
	{Name: "ScanName", Flag: "scan-name", Type: "*string", Required: true},
}

var fields_get_account_configuration = []leanruntime.Field{}

var fields_get_findings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanName", Flag: "scan-name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.Status", Required: false},
}

var fields_get_metrics_summary = []leanruntime.Field{
	{Name: "Date", Flag: "date", Type: "*time.Time", Required: true},
}

var fields_get_scan = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
	{Name: "ScanName", Flag: "scan-name", Type: "*string", Required: true},
}

var fields_list_findings_metrics = []leanruntime.Field{
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: true},
}

var fields_list_scans = []leanruntime.Field{
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

var fields_update_account_configuration = []leanruntime.Field{
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "*types.EncryptionConfig", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-findings": {
			Name:   "batch-get-findings",
			Fields: fields_batch_get_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFindingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_findings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetFindings(ctx, input)
			},
		},
		"create-scan": {
			Name:   "create-scan",
			Fields: fields_create_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScan(ctx, input)
			},
		},
		"create-upload-url": {
			Name:   "create-upload-url",
			Fields: fields_create_upload_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUploadUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_upload_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUploadUrl(ctx, input)
			},
		},
		"get-account-configuration": {
			Name:   "get-account-configuration",
			Fields: fields_get_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountConfiguration(ctx, input)
			},
		},
		"get-findings": {
			Name:   "get-findings",
			Fields: fields_get_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFindings(ctx, input)
				}
				var results []*svc.GetFindingsOutput
				p := svc.NewGetFindingsPaginator(client, input)
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
		"get-metrics-summary": {
			Name:   "get-metrics-summary",
			Fields: fields_get_metrics_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricsSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metrics_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetricsSummary(ctx, input)
			},
		},
		"get-scan": {
			Name:   "get-scan",
			Fields: fields_get_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScan(ctx, input)
			},
		},
		"list-findings-metrics": {
			Name:   "list-findings-metrics",
			Fields: fields_list_findings_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFindingsMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_findings_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFindingsMetrics(ctx, input)
				}
				var results []*svc.ListFindingsMetricsOutput
				p := svc.NewListFindingsMetricsPaginator(client, input)
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
		"list-scans": {
			Name:   "list-scans",
			Fields: fields_list_scans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScans(ctx, input)
				}
				var results []*svc.ListScansOutput
				p := svc.NewListScansPaginator(client, input)
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
		"update-account-configuration": {
			Name:   "update-account-configuration",
			Fields: fields_update_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codegurusecurity", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
