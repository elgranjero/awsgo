package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
)

var fields_create_export = []leanruntime.Field{
	{Name: "Export", Flag: "export", Type: "*types.Export", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
}

var fields_delete_export = []leanruntime.Field{
	{Name: "ExportArn", Flag: "export-arn", Type: "*string", Required: true},
}

var fields_get_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "ExportArn", Flag: "export-arn", Type: "*string", Required: true},
}

var fields_get_export = []leanruntime.Field{
	{Name: "ExportArn", Flag: "export-arn", Type: "*string", Required: true},
}

var fields_get_table = []leanruntime.Field{
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TableProperties", Flag: "table-properties", Type: "map[string]string", Required: false},
}

var fields_list_executions = []leanruntime.Field{
	{Name: "ExportArn", Flag: "export-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_exports = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tables = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTagKeys", Flag: "resource-tag-keys", Type: "[]string", Required: true},
}

var fields_update_export = []leanruntime.Field{
	{Name: "Export", Flag: "export", Type: "*types.Export", Required: true},
	{Name: "ExportArn", Flag: "export-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-export": {
			Name:   "create-export",
			Fields: fields_create_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExport(ctx, input)
			},
		},
		"delete-export": {
			Name:   "delete-export",
			Fields: fields_delete_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExport(ctx, input)
			},
		},
		"get-execution": {
			Name:   "get-execution",
			Fields: fields_get_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExecution(ctx, input)
			},
		},
		"get-export": {
			Name:   "get-export",
			Fields: fields_get_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExport(ctx, input)
			},
		},
		"get-table": {
			Name:   "get-table",
			Fields: fields_get_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTable(ctx, input)
			},
		},
		"list-executions": {
			Name:   "list-executions",
			Fields: fields_list_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExecutions(ctx, input)
				}
				var results []*svc.ListExecutionsOutput
				p := svc.NewListExecutionsPaginator(client, input)
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
		"list-exports": {
			Name:   "list-exports",
			Fields: fields_list_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExports(ctx, input)
				}
				var results []*svc.ListExportsOutput
				p := svc.NewListExportsPaginator(client, input)
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
		"list-tables": {
			Name:   "list-tables",
			Fields: fields_list_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTables(ctx, input)
				}
				var results []*svc.ListTablesOutput
				p := svc.NewListTablesPaginator(client, input)
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
		"update-export": {
			Name:   "update-export",
			Fields: fields_update_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExport(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bcmdataexports", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
