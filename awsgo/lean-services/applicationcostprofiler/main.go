package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler"
)

var fields_delete_report_definition = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_get_report_definition = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_import_application_usage = []leanruntime.Field{
	{Name: "SourceS3Location", Flag: "source-s3-location", Type: "*types.SourceS3Location", Required: true},
}

var fields_list_report_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_report_definition = []leanruntime.Field{
	{Name: "DestinationS3Location", Flag: "destination-s3-location", Type: "*types.S3Location", Required: true},
	{Name: "Format", Flag: "format", Type: "types.Format", Required: true},
	{Name: "ReportDescription", Flag: "report-description", Type: "*string", Required: true},
	{Name: "ReportFrequency", Flag: "report-frequency", Type: "types.ReportFrequency", Required: true},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_update_report_definition = []leanruntime.Field{
	{Name: "DestinationS3Location", Flag: "destination-s3-location", Type: "*types.S3Location", Required: true},
	{Name: "Format", Flag: "format", Type: "types.Format", Required: true},
	{Name: "ReportDescription", Flag: "report-description", Type: "*string", Required: true},
	{Name: "ReportFrequency", Flag: "report-frequency", Type: "types.ReportFrequency", Required: true},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
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
		"get-report-definition": {
			Name:   "get-report-definition",
			Fields: fields_get_report_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReportDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_report_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReportDefinition(ctx, input)
			},
		},
		"import-application-usage": {
			Name:   "import-application-usage",
			Fields: fields_import_application_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportApplicationUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_application_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportApplicationUsage(ctx, input)
			},
		},
		"list-report-definitions": {
			Name:   "list-report-definitions",
			Fields: fields_list_report_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReportDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_report_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReportDefinitions(ctx, input)
				}
				var results []*svc.ListReportDefinitionsOutput
				p := svc.NewListReportDefinitionsPaginator(client, input)
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
		"update-report-definition": {
			Name:   "update-report-definition",
			Fields: fields_update_report_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReportDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_report_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReportDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("applicationcostprofiler", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
