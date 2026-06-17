package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sagemakergeospatial/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-earth-observation-job", "delete-vector-enrichment-job", "export-earth-observation-job", "export-vector-enrichment-job", "get-earth-observation-job", "get-raster-data-collection", "get-tile", "get-vector-enrichment-job", "list-earth-observation-jobs", "list-raster-data-collections", "list-tags-for-resource", "list-vector-enrichment-jobs", "search-raster-data-collection", "start-earth-observation-job", "start-vector-enrichment-job", "stop-earth-observation-job", "stop-vector-enrichment-job", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"delete-earth-observation-job": true, "delete-vector-enrichment-job": true, "export-earth-observation-job": true, "export-vector-enrichment-job": true, "get-earth-observation-job": true, "get-raster-data-collection": true, "get-tile": true, "get-vector-enrichment-job": true, "list-earth-observation-jobs": true, "list-raster-data-collections": true, "list-tags-for-resource": true, "list-vector-enrichment-jobs": true, "search-raster-data-collection": true, "start-earth-observation-job": true, "start-vector-enrichment-job": true, "stop-earth-observation-job": true, "stop-vector-enrichment-job": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"delete-earth-observation-job":  {"Arn"},
			"delete-vector-enrichment-job":  {"Arn"},
			"export-earth-observation-job":  {"Arn", "ClientToken", "ExecutionRoleArn", "ExportSourceImages", "OutputConfig"},
			"export-vector-enrichment-job":  {"Arn", "ClientToken", "ExecutionRoleArn", "OutputConfig"},
			"get-earth-observation-job":     {"Arn"},
			"get-raster-data-collection":    {"Arn"},
			"get-tile":                      {"Arn", "ExecutionRoleArn", "ImageAssets", "ImageMask", "OutputDataType", "OutputFormat", "PropertyFilters", "Target", "TimeRangeFilter", "X", "Y", "Z"},
			"get-vector-enrichment-job":     {"Arn"},
			"list-earth-observation-jobs":   {"MaxResults", "NextToken", "SortBy", "SortOrder", "StatusEquals"},
			"list-raster-data-collections":  {"MaxResults", "NextToken"},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-vector-enrichment-jobs":   {"MaxResults", "NextToken", "SortBy", "SortOrder", "StatusEquals"},
			"search-raster-data-collection": {"Arn", "NextToken", "RasterDataCollectionQuery"},
			"start-earth-observation-job":   {"ClientToken", "ExecutionRoleArn", "InputConfig", "JobConfig", "KmsKeyId", "Name", "Tags"},
			"start-vector-enrichment-job":   {"ClientToken", "ExecutionRoleArn", "InputConfig", "JobConfig", "KmsKeyId", "Name", "Tags"},
			"stop-earth-observation-job":    {"Arn"},
			"stop-vector-enrichment-job":    {"Arn"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-earth-observation-job":  {"Arn": "*string"},
			"delete-vector-enrichment-job":  {"Arn": "*string"},
			"export-earth-observation-job":  {"Arn": "*string", "ClientToken": "*string", "ExecutionRoleArn": "*string", "ExportSourceImages": "*bool", "OutputConfig": "*types.OutputConfigInput"},
			"export-vector-enrichment-job":  {"Arn": "*string", "ClientToken": "*string", "ExecutionRoleArn": "*string", "OutputConfig": "*types.ExportVectorEnrichmentJobOutputConfig"},
			"get-earth-observation-job":     {"Arn": "*string"},
			"get-raster-data-collection":    {"Arn": "*string"},
			"get-tile":                      {"Arn": "*string", "ExecutionRoleArn": "*string", "ImageAssets": "[]string", "ImageMask": "*bool", "OutputDataType": "types.OutputType", "OutputFormat": "*string", "PropertyFilters": "*string", "Target": "types.TargetOptions", "TimeRangeFilter": "*string", "X": "*int32", "Y": "*int32", "Z": "*int32"},
			"get-vector-enrichment-job":     {"Arn": "*string"},
			"list-earth-observation-jobs":   {"MaxResults": "*int32", "NextToken": "*string", "SortBy": "*string", "SortOrder": "types.SortOrder", "StatusEquals": "types.EarthObservationJobStatus"},
			"list-raster-data-collections":  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"list-vector-enrichment-jobs":   {"MaxResults": "*int32", "NextToken": "*string", "SortBy": "*string", "SortOrder": "types.SortOrder", "StatusEquals": "*string"},
			"search-raster-data-collection": {"Arn": "*string", "NextToken": "*string", "RasterDataCollectionQuery": "*types.RasterDataCollectionQueryWithBandFilterInput"},
			"start-earth-observation-job":   {"ClientToken": "*string", "ExecutionRoleArn": "*string", "InputConfig": "*types.InputConfigInput", "JobConfig": "types.JobConfigInput", "KmsKeyId": "*string", "Name": "*string", "Tags": "map[string]string"},
			"start-vector-enrichment-job":   {"ClientToken": "*string", "ExecutionRoleArn": "*string", "InputConfig": "*types.VectorEnrichmentJobInputConfig", "JobConfig": "types.VectorEnrichmentJobConfig", "KmsKeyId": "*string", "Name": "*string", "Tags": "map[string]string"},
			"stop-earth-observation-job":    {"Arn": "*string"},
			"stop-vector-enrichment-job":    {"Arn": "*string"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-earth-observation-job":  {"Arn"},
			"delete-vector-enrichment-job":  {"Arn"},
			"export-earth-observation-job":  {"Arn", "ExecutionRoleArn", "OutputConfig"},
			"export-vector-enrichment-job":  {"Arn", "ExecutionRoleArn", "OutputConfig"},
			"get-earth-observation-job":     {"Arn"},
			"get-raster-data-collection":    {"Arn"},
			"get-tile":                      {"Arn", "ImageAssets", "Target", "X", "Y", "Z"},
			"get-vector-enrichment-job":     {"Arn"},
			"list-earth-observation-jobs":   {},
			"list-raster-data-collections":  {},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-vector-enrichment-jobs":   {},
			"search-raster-data-collection": {"Arn", "RasterDataCollectionQuery"},
			"start-earth-observation-job":   {"ExecutionRoleArn", "InputConfig", "JobConfig", "Name"},
			"start-vector-enrichment-job":   {"ExecutionRoleArn", "InputConfig", "JobConfig", "Name"},
			"stop-earth-observation-job":    {"Arn"},
			"stop-vector-enrichment-job":    {"Arn"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sagemakergeospatial", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
