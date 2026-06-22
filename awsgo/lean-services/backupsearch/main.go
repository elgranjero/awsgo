package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/backupsearch"
)

var fields_get_search_job = []leanruntime.Field{
	{Name: "SearchJobIdentifier", Flag: "search-job-identifier", Type: "*string", Required: true},
}

var fields_get_search_result_export_job = []leanruntime.Field{
	{Name: "ExportJobIdentifier", Flag: "export-job-identifier", Type: "*string", Required: true},
}

var fields_list_search_job_backups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchJobIdentifier", Flag: "search-job-identifier", Type: "*string", Required: true},
}

var fields_list_search_job_results = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchJobIdentifier", Flag: "search-job-identifier", Type: "*string", Required: true},
}

var fields_list_search_jobs = []leanruntime.Field{
	{Name: "ByStatus", Flag: "by-status", Type: "types.SearchJobState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_search_result_export_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchJobIdentifier", Flag: "search-job-identifier", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ExportJobStatus", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_search_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "ItemFilters", Flag: "item-filters", Type: "*types.ItemFilters", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SearchScope", Flag: "search-scope", Type: "*types.SearchScope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]*string", Required: false},
}

var fields_start_search_result_export_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExportSpecification", Flag: "export-specification", Type: "types.ExportSpecification", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SearchJobIdentifier", Flag: "search-job-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]*string", Required: false},
}

var fields_stop_search_job = []leanruntime.Field{
	{Name: "SearchJobIdentifier", Flag: "search-job-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-search-job": {
			Name:   "get-search-job",
			Fields: fields_get_search_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSearchJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_search_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSearchJob(ctx, input)
			},
		},
		"get-search-result-export-job": {
			Name:   "get-search-result-export-job",
			Fields: fields_get_search_result_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSearchResultExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_search_result_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSearchResultExportJob(ctx, input)
			},
		},
		"list-search-job-backups": {
			Name:   "list-search-job-backups",
			Fields: fields_list_search_job_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSearchJobBackupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_search_job_backups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSearchJobBackups(ctx, input)
				}
				var results []*svc.ListSearchJobBackupsOutput
				p := svc.NewListSearchJobBackupsPaginator(client, input)
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
		"list-search-job-results": {
			Name:   "list-search-job-results",
			Fields: fields_list_search_job_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSearchJobResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_search_job_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSearchJobResults(ctx, input)
				}
				var results []*svc.ListSearchJobResultsOutput
				p := svc.NewListSearchJobResultsPaginator(client, input)
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
		"list-search-jobs": {
			Name:   "list-search-jobs",
			Fields: fields_list_search_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSearchJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_search_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSearchJobs(ctx, input)
				}
				var results []*svc.ListSearchJobsOutput
				p := svc.NewListSearchJobsPaginator(client, input)
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
		"list-search-result-export-jobs": {
			Name:   "list-search-result-export-jobs",
			Fields: fields_list_search_result_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSearchResultExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_search_result_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSearchResultExportJobs(ctx, input)
				}
				var results []*svc.ListSearchResultExportJobsOutput
				p := svc.NewListSearchResultExportJobsPaginator(client, input)
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
		"start-search-job": {
			Name:   "start-search-job",
			Fields: fields_start_search_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSearchJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_search_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSearchJob(ctx, input)
			},
		},
		"start-search-result-export-job": {
			Name:   "start-search-result-export-job",
			Fields: fields_start_search_result_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSearchResultExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_search_result_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSearchResultExportJob(ctx, input)
			},
		},
		"stop-search-job": {
			Name:   "stop-search-job",
			Fields: fields_stop_search_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSearchJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_search_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSearchJob(ctx, input)
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
	if err := leanruntime.Execute("backupsearch", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
