package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pi"
)

var fields_create_performance_analysis_report = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_performance_analysis_report = []leanruntime.Field{
	{Name: "AnalysisReportId", Flag: "analysis-report-id", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_describe_dimension_keys = []leanruntime.Field{
	{Name: "AdditionalMetrics", Flag: "additional-metrics", Type: "[]string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Filter", Flag: "filter", Type: "map[string]string", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "*types.DimensionGroup", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metric", Flag: "metric", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PartitionBy", Flag: "partition-by", Type: "*types.DimensionGroup", Required: false},
	{Name: "PeriodInSeconds", Flag: "period-in-seconds", Type: "*int32", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_dimension_key_details = []leanruntime.Field{
	{Name: "Group", Flag: "group", Type: "*string", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RequestedDimensions", Flag: "requested-dimensions", Type: "[]string", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_get_performance_analysis_report = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "types.AcceptLanguage", Required: false},
	{Name: "AnalysisReportId", Flag: "analysis-report-id", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
	{Name: "TextFormat", Flag: "text-format", Type: "types.TextFormat", Required: false},
}

var fields_get_resource_metadata = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_get_resource_metrics = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricQueries", Flag: "metric-queries", Type: "[]types.MetricQuery", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PeriodAlignment", Flag: "period-alignment", Type: "types.PeriodAlignment", Required: false},
	{Name: "PeriodInSeconds", Flag: "period-in-seconds", Type: "*int32", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_available_resource_dimensions = []leanruntime.Field{
	{Name: "AuthorizedActions", Flag: "authorized-actions", Type: "[]types.FineGrainedAction", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_list_available_resource_metrics = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricTypes", Flag: "metric-types", Type: "[]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_list_performance_analysis_reports = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ListTags", Flag: "list-tags", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ServiceType", Flag: "service-type", Type: "types.ServiceType", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-performance-analysis-report": {
			Name:   "create-performance-analysis-report",
			Fields: fields_create_performance_analysis_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePerformanceAnalysisReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_performance_analysis_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePerformanceAnalysisReport(ctx, input)
			},
		},
		"delete-performance-analysis-report": {
			Name:   "delete-performance-analysis-report",
			Fields: fields_delete_performance_analysis_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePerformanceAnalysisReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_performance_analysis_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePerformanceAnalysisReport(ctx, input)
			},
		},
		"describe-dimension-keys": {
			Name:   "describe-dimension-keys",
			Fields: fields_describe_dimension_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDimensionKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_dimension_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDimensionKeys(ctx, input)
				}
				var results []*svc.DescribeDimensionKeysOutput
				p := svc.NewDescribeDimensionKeysPaginator(client, input)
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
		"get-dimension-key-details": {
			Name:   "get-dimension-key-details",
			Fields: fields_get_dimension_key_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDimensionKeyDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dimension_key_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDimensionKeyDetails(ctx, input)
			},
		},
		"get-performance-analysis-report": {
			Name:   "get-performance-analysis-report",
			Fields: fields_get_performance_analysis_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPerformanceAnalysisReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_performance_analysis_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPerformanceAnalysisReport(ctx, input)
			},
		},
		"get-resource-metadata": {
			Name:   "get-resource-metadata",
			Fields: fields_get_resource_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceMetadata(ctx, input)
			},
		},
		"get-resource-metrics": {
			Name:   "get-resource-metrics",
			Fields: fields_get_resource_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceMetrics(ctx, input)
				}
				var results []*svc.GetResourceMetricsOutput
				p := svc.NewGetResourceMetricsPaginator(client, input)
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
		"list-available-resource-dimensions": {
			Name:   "list-available-resource-dimensions",
			Fields: fields_list_available_resource_dimensions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableResourceDimensionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_available_resource_dimensions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAvailableResourceDimensions(ctx, input)
				}
				var results []*svc.ListAvailableResourceDimensionsOutput
				p := svc.NewListAvailableResourceDimensionsPaginator(client, input)
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
		"list-available-resource-metrics": {
			Name:   "list-available-resource-metrics",
			Fields: fields_list_available_resource_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableResourceMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_available_resource_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAvailableResourceMetrics(ctx, input)
				}
				var results []*svc.ListAvailableResourceMetricsOutput
				p := svc.NewListAvailableResourceMetricsPaginator(client, input)
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
		"list-performance-analysis-reports": {
			Name:   "list-performance-analysis-reports",
			Fields: fields_list_performance_analysis_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPerformanceAnalysisReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_performance_analysis_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPerformanceAnalysisReports(ctx, input)
				}
				var results []*svc.ListPerformanceAnalysisReportsOutput
				p := svc.NewListPerformanceAnalysisReportsPaginator(client, input)
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
	}
	if err := leanruntime.Execute("pi", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
