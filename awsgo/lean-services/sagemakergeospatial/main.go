package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial"
)

var fields_delete_earth_observation_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_vector_enrichment_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_export_earth_observation_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "ExportSourceImages", Flag: "export-source-images", Type: "*bool", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfigInput", Required: true},
}

var fields_export_vector_enrichment_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.ExportVectorEnrichmentJobOutputConfig", Required: true},
}

var fields_get_earth_observation_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_raster_data_collection = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_tile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "ImageAssets", Flag: "image-assets", Type: "[]string", Required: true},
	{Name: "ImageMask", Flag: "image-mask", Type: "*bool", Required: false},
	{Name: "OutputDataType", Flag: "output-data-type", Type: "types.OutputType", Required: false},
	{Name: "OutputFormat", Flag: "output-format", Type: "*string", Required: false},
	{Name: "PropertyFilters", Flag: "property-filters", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "types.TargetOptions", Required: true},
	{Name: "TimeRangeFilter", Flag: "time-range-filter", Type: "*string", Required: false},
	{Name: "X", Flag: "x", Type: "*int32", Required: true},
	{Name: "Y", Flag: "y", Type: "*int32", Required: true},
	{Name: "Z", Flag: "z", Type: "*int32", Required: true},
}

var fields_get_vector_enrichment_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_earth_observation_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.EarthObservationJobStatus", Required: false},
}

var fields_list_raster_data_collections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_vector_enrichment_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "*string", Required: false},
}

var fields_search_raster_data_collection = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RasterDataCollectionQuery", Flag: "raster-data-collection-query", Type: "*types.RasterDataCollectionQueryWithBandFilterInput", Required: true},
}

var fields_start_earth_observation_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "InputConfig", Flag: "input-config", Type: "*types.InputConfigInput", Required: true},
	{Name: "JobConfig", Flag: "job-config", Type: "types.JobConfigInput", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_vector_enrichment_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "InputConfig", Flag: "input-config", Type: "*types.VectorEnrichmentJobInputConfig", Required: true},
	{Name: "JobConfig", Flag: "job-config", Type: "types.VectorEnrichmentJobConfig", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_earth_observation_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_stop_vector_enrichment_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-earth-observation-job": {
			Name:   "delete-earth-observation-job",
			Fields: fields_delete_earth_observation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEarthObservationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_earth_observation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEarthObservationJob(ctx, input)
			},
		},
		"delete-vector-enrichment-job": {
			Name:   "delete-vector-enrichment-job",
			Fields: fields_delete_vector_enrichment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVectorEnrichmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vector_enrichment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVectorEnrichmentJob(ctx, input)
			},
		},
		"export-earth-observation-job": {
			Name:   "export-earth-observation-job",
			Fields: fields_export_earth_observation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportEarthObservationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_earth_observation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportEarthObservationJob(ctx, input)
			},
		},
		"export-vector-enrichment-job": {
			Name:   "export-vector-enrichment-job",
			Fields: fields_export_vector_enrichment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportVectorEnrichmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_vector_enrichment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportVectorEnrichmentJob(ctx, input)
			},
		},
		"get-earth-observation-job": {
			Name:   "get-earth-observation-job",
			Fields: fields_get_earth_observation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEarthObservationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_earth_observation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEarthObservationJob(ctx, input)
			},
		},
		"get-raster-data-collection": {
			Name:   "get-raster-data-collection",
			Fields: fields_get_raster_data_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRasterDataCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_raster_data_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRasterDataCollection(ctx, input)
			},
		},
		"get-tile": {
			Name:   "get-tile",
			Fields: fields_get_tile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTile(ctx, input)
			},
		},
		"get-vector-enrichment-job": {
			Name:   "get-vector-enrichment-job",
			Fields: fields_get_vector_enrichment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVectorEnrichmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vector_enrichment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVectorEnrichmentJob(ctx, input)
			},
		},
		"list-earth-observation-jobs": {
			Name:   "list-earth-observation-jobs",
			Fields: fields_list_earth_observation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEarthObservationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_earth_observation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEarthObservationJobs(ctx, input)
				}
				var results []*svc.ListEarthObservationJobsOutput
				p := svc.NewListEarthObservationJobsPaginator(client, input)
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
		"list-raster-data-collections": {
			Name:   "list-raster-data-collections",
			Fields: fields_list_raster_data_collections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRasterDataCollectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_raster_data_collections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRasterDataCollections(ctx, input)
				}
				var results []*svc.ListRasterDataCollectionsOutput
				p := svc.NewListRasterDataCollectionsPaginator(client, input)
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
		"list-vector-enrichment-jobs": {
			Name:   "list-vector-enrichment-jobs",
			Fields: fields_list_vector_enrichment_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVectorEnrichmentJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vector_enrichment_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVectorEnrichmentJobs(ctx, input)
				}
				var results []*svc.ListVectorEnrichmentJobsOutput
				p := svc.NewListVectorEnrichmentJobsPaginator(client, input)
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
		"search-raster-data-collection": {
			Name:   "search-raster-data-collection",
			Fields: fields_search_raster_data_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchRasterDataCollectionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_raster_data_collection, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchRasterDataCollection(ctx, input)
				}
				var results []*svc.SearchRasterDataCollectionOutput
				p := svc.NewSearchRasterDataCollectionPaginator(client, input)
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
		"start-earth-observation-job": {
			Name:   "start-earth-observation-job",
			Fields: fields_start_earth_observation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEarthObservationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_earth_observation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEarthObservationJob(ctx, input)
			},
		},
		"start-vector-enrichment-job": {
			Name:   "start-vector-enrichment-job",
			Fields: fields_start_vector_enrichment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVectorEnrichmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_vector_enrichment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVectorEnrichmentJob(ctx, input)
			},
		},
		"stop-earth-observation-job": {
			Name:   "stop-earth-observation-job",
			Fields: fields_stop_earth_observation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEarthObservationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_earth_observation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEarthObservationJob(ctx, input)
			},
		},
		"stop-vector-enrichment-job": {
			Name:   "stop-vector-enrichment-job",
			Fields: fields_stop_vector_enrichment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopVectorEnrichmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_vector_enrichment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopVectorEnrichmentJob(ctx, input)
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
	if err := leanruntime.Execute("sagemakergeospatial", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
