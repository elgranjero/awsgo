package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
)

var fields_describe_report_creation = []leanruntime.Field{}

var fields_get_compliance_summary = []leanruntime.Field{
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupByAttribute", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
	{Name: "RegionFilters", Flag: "region-filters", Type: "[]string", Required: false},
	{Name: "ResourceTypeFilters", Flag: "resource-type-filters", Type: "[]string", Required: false},
	{Name: "TagKeyFilters", Flag: "tag-key-filters", Type: "[]string", Required: false},
	{Name: "TargetIdFilters", Flag: "target-id-filters", Type: "[]string", Required: false},
}

var fields_get_resources = []leanruntime.Field{
	{Name: "ExcludeCompliantResources", Flag: "exclude-compliant-resources", Type: "*bool", Required: false},
	{Name: "IncludeComplianceDetails", Flag: "include-compliance-details", Type: "*bool", Required: false},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
	{Name: "ResourceARNList", Flag: "resource-arn-list", Type: "[]string", Required: false},
	{Name: "ResourceTypeFilters", Flag: "resource-type-filters", Type: "[]string", Required: false},
	{Name: "ResourcesPerPage", Flag: "resources-per-page", Type: "*int32", Required: false},
	{Name: "TagFilters", Flag: "tag-filters", Type: "[]types.TagFilter", Required: false},
	{Name: "TagsPerPage", Flag: "tags-per-page", Type: "*int32", Required: false},
}

var fields_get_tag_keys = []leanruntime.Field{
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
}

var fields_get_tag_values = []leanruntime.Field{
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
}

var fields_list_required_tags = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_report_creation = []leanruntime.Field{
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: true},
}

var fields_tag_resources = []leanruntime.Field{
	{Name: "ResourceARNList", Flag: "resource-arn-list", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resources = []leanruntime.Field{
	{Name: "ResourceARNList", Flag: "resource-arn-list", Type: "[]string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-report-creation": {
			Name:   "describe-report-creation",
			Fields: fields_describe_report_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReportCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_report_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReportCreation(ctx, input)
			},
		},
		"get-compliance-summary": {
			Name:   "get-compliance-summary",
			Fields: fields_get_compliance_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComplianceSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_compliance_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetComplianceSummary(ctx, input)
				}
				var results []*svc.GetComplianceSummaryOutput
				p := svc.NewGetComplianceSummaryPaginator(client, input)
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
		"get-resources": {
			Name:   "get-resources",
			Fields: fields_get_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResources(ctx, input)
				}
				var results []*svc.GetResourcesOutput
				p := svc.NewGetResourcesPaginator(client, input)
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
		"get-tag-keys": {
			Name:   "get-tag-keys",
			Fields: fields_get_tag_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_tag_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTagKeys(ctx, input)
				}
				var results []*svc.GetTagKeysOutput
				p := svc.NewGetTagKeysPaginator(client, input)
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
		"get-tag-values": {
			Name:   "get-tag-values",
			Fields: fields_get_tag_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTagValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_tag_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTagValues(ctx, input)
				}
				var results []*svc.GetTagValuesOutput
				p := svc.NewGetTagValuesPaginator(client, input)
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
		"list-required-tags": {
			Name:   "list-required-tags",
			Fields: fields_list_required_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRequiredTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_required_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRequiredTags(ctx, input)
				}
				var results []*svc.ListRequiredTagsOutput
				p := svc.NewListRequiredTagsPaginator(client, input)
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
		"start-report-creation": {
			Name:   "start-report-creation",
			Fields: fields_start_report_creation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReportCreationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_report_creation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReportCreation(ctx, input)
			},
		},
		"tag-resources": {
			Name:   "tag-resources",
			Fields: fields_tag_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResources(ctx, input)
			},
		},
		"untag-resources": {
			Name:   "untag-resources",
			Fields: fields_untag_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResources(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("resourcegroupstaggingapi", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
