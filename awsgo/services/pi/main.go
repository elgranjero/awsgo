package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pi/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-performance-analysis-report", "delete-performance-analysis-report", "describe-dimension-keys", "get-dimension-key-details", "get-performance-analysis-report", "get-resource-metadata", "get-resource-metrics", "list-available-resource-dimensions", "list-available-resource-metrics", "list-performance-analysis-reports", "list-tags-for-resource", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-performance-analysis-report": true, "delete-performance-analysis-report": true, "describe-dimension-keys": true, "get-dimension-key-details": true, "get-performance-analysis-report": true, "get-resource-metadata": true, "get-resource-metrics": true, "list-available-resource-dimensions": true, "list-available-resource-metrics": true, "list-performance-analysis-reports": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-performance-analysis-report": {"EndTime", "Identifier", "ServiceType", "StartTime", "Tags"},
			"delete-performance-analysis-report": {"AnalysisReportId", "Identifier", "ServiceType"},
			"describe-dimension-keys":            {"AdditionalMetrics", "EndTime", "Filter", "GroupBy", "Identifier", "MaxResults", "Metric", "NextToken", "PartitionBy", "PeriodInSeconds", "ServiceType", "StartTime"},
			"get-dimension-key-details":          {"Group", "GroupIdentifier", "Identifier", "RequestedDimensions", "ServiceType"},
			"get-performance-analysis-report":    {"AcceptLanguage", "AnalysisReportId", "Identifier", "ServiceType", "TextFormat"},
			"get-resource-metadata":              {"Identifier", "ServiceType"},
			"get-resource-metrics":               {"EndTime", "Identifier", "MaxResults", "MetricQueries", "NextToken", "PeriodAlignment", "PeriodInSeconds", "ServiceType", "StartTime"},
			"list-available-resource-dimensions": {"AuthorizedActions", "Identifier", "MaxResults", "Metrics", "NextToken", "ServiceType"},
			"list-available-resource-metrics":    {"Identifier", "MaxResults", "MetricTypes", "NextToken", "ServiceType"},
			"list-performance-analysis-reports":  {"Identifier", "ListTags", "MaxResults", "NextToken", "ServiceType"},
			"list-tags-for-resource":             {"ResourceARN", "ServiceType"},
			"tag-resource":                       {"ResourceARN", "ServiceType", "Tags"},
			"untag-resource":                     {"ResourceARN", "ServiceType", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-performance-analysis-report": {"EndTime": "*time.Time", "Identifier": "*string", "ServiceType": "types.ServiceType", "StartTime": "*time.Time", "Tags": "[]types.Tag"},
			"delete-performance-analysis-report": {"AnalysisReportId": "*string", "Identifier": "*string", "ServiceType": "types.ServiceType"},
			"describe-dimension-keys":            {"AdditionalMetrics": "[]string", "EndTime": "*time.Time", "Filter": "map[string]string", "GroupBy": "*types.DimensionGroup", "Identifier": "*string", "MaxResults": "*int32", "Metric": "*string", "NextToken": "*string", "PartitionBy": "*types.DimensionGroup", "PeriodInSeconds": "*int32", "ServiceType": "types.ServiceType", "StartTime": "*time.Time"},
			"get-dimension-key-details":          {"Group": "*string", "GroupIdentifier": "*string", "Identifier": "*string", "RequestedDimensions": "[]string", "ServiceType": "types.ServiceType"},
			"get-performance-analysis-report":    {"AcceptLanguage": "types.AcceptLanguage", "AnalysisReportId": "*string", "Identifier": "*string", "ServiceType": "types.ServiceType", "TextFormat": "types.TextFormat"},
			"get-resource-metadata":              {"Identifier": "*string", "ServiceType": "types.ServiceType"},
			"get-resource-metrics":               {"EndTime": "*time.Time", "Identifier": "*string", "MaxResults": "*int32", "MetricQueries": "[]types.MetricQuery", "NextToken": "*string", "PeriodAlignment": "types.PeriodAlignment", "PeriodInSeconds": "*int32", "ServiceType": "types.ServiceType", "StartTime": "*time.Time"},
			"list-available-resource-dimensions": {"AuthorizedActions": "[]types.FineGrainedAction", "Identifier": "*string", "MaxResults": "*int32", "Metrics": "[]string", "NextToken": "*string", "ServiceType": "types.ServiceType"},
			"list-available-resource-metrics":    {"Identifier": "*string", "MaxResults": "*int32", "MetricTypes": "[]string", "NextToken": "*string", "ServiceType": "types.ServiceType"},
			"list-performance-analysis-reports":  {"Identifier": "*string", "ListTags": "*bool", "MaxResults": "*int32", "NextToken": "*string", "ServiceType": "types.ServiceType"},
			"list-tags-for-resource":             {"ResourceARN": "*string", "ServiceType": "types.ServiceType"},
			"tag-resource":                       {"ResourceARN": "*string", "ServiceType": "types.ServiceType", "Tags": "[]types.Tag"},
			"untag-resource":                     {"ResourceARN": "*string", "ServiceType": "types.ServiceType", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-performance-analysis-report": {"EndTime", "Identifier", "ServiceType", "StartTime"},
			"delete-performance-analysis-report": {"AnalysisReportId", "Identifier", "ServiceType"},
			"describe-dimension-keys":            {"EndTime", "GroupBy", "Identifier", "Metric", "ServiceType", "StartTime"},
			"get-dimension-key-details":          {"Group", "GroupIdentifier", "Identifier", "ServiceType"},
			"get-performance-analysis-report":    {"AnalysisReportId", "Identifier", "ServiceType"},
			"get-resource-metadata":              {"Identifier", "ServiceType"},
			"get-resource-metrics":               {"EndTime", "Identifier", "MetricQueries", "ServiceType", "StartTime"},
			"list-available-resource-dimensions": {"Identifier", "Metrics", "ServiceType"},
			"list-available-resource-metrics":    {"Identifier", "MetricTypes", "ServiceType"},
			"list-performance-analysis-reports":  {"Identifier", "ServiceType"},
			"list-tags-for-resource":             {"ResourceARN", "ServiceType"},
			"tag-resource":                       {"ResourceARN", "ServiceType", "Tags"},
			"untag-resource":                     {"ResourceARN", "ServiceType", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pi", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
