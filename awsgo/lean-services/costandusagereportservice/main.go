package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

var fields_delete_report_definition = []leanruntime.Field{
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: true},
}

var fields_describe_report_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: true},
}

var fields_modify_report_definition = []leanruntime.Field{
	{Name: "ReportDefinition", Flag: "report-definition", Type: "*types.ReportDefinition", Required: true},
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: true},
}

var fields_put_report_definition = []leanruntime.Field{
	{Name: "ReportDefinition", Flag: "report-definition", Type: "*types.ReportDefinition", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-report-definition": {
			Name:   "delete-report-definition",
			Fields: fields_delete_report_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReportDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_report_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReportDefinition(ctx, input)
			},
		},
		"describe-report-definitions": {
			Name:   "describe-report-definitions",
			Fields: fields_describe_report_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReportDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_report_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReportDefinitions(ctx, input)
				}
				var results []*svc.DescribeReportDefinitionsOutput
				p := svc.NewDescribeReportDefinitionsPaginator(client, input)
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
		"modify-report-definition": {
			Name:   "modify-report-definition",
			Fields: fields_modify_report_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReportDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_report_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReportDefinition(ctx, input)
			},
		},
		"put-report-definition": {
			Name:   "put-report-definition",
			Fields: fields_put_report_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutReportDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_report_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutReportDefinition(ctx, input)
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
	}
	if err := leanruntime.Execute("costandusagereportservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
