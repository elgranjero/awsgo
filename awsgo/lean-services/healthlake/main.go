package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/healthlake"
)

var fields_create_fhir_datastore = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatastoreName", Flag: "datastore-name", Type: "*string", Required: false},
	{Name: "DatastoreTypeVersion", Flag: "datastore-type-version", Type: "types.FHIRVersion", Required: true},
	{Name: "IdentityProviderConfiguration", Flag: "identity-provider-configuration", Type: "*types.IdentityProviderConfiguration", Required: false},
	{Name: "PreloadDataConfig", Flag: "preload-data-config", Type: "*types.PreloadDataConfig", Required: false},
	{Name: "SseConfiguration", Flag: "sse-configuration", Type: "*types.SseConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_fhir_datastore = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
}

var fields_describe_fhir_datastore = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
}

var fields_describe_fhir_export_job = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_fhir_import_job = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_list_fhir_datastores = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DatastoreFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fhir_export_jobs = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "JobStatus", Flag: "job-status", Type: "types.JobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubmittedAfter", Flag: "submitted-after", Type: "*time.Time", Required: false},
	{Name: "SubmittedBefore", Flag: "submitted-before", Type: "*time.Time", Required: false},
}

var fields_list_fhir_import_jobs = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "JobStatus", Flag: "job-status", Type: "types.JobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubmittedAfter", Flag: "submitted-after", Type: "*time.Time", Required: false},
	{Name: "SubmittedBefore", Flag: "submitted-before", Type: "*time.Time", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_fhir_export_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "types.OutputDataConfig", Required: true},
}

var fields_start_fhir_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "JobOutputDataConfig", Flag: "job-output-data-config", Type: "types.OutputDataConfig", Required: true},
	{Name: "ValidationLevel", Flag: "validation-level", Type: "types.ValidationLevel", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-fhir-datastore": {
			Name:   "create-fhir-datastore",
			Fields: fields_create_fhir_datastore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFHIRDatastoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fhir_datastore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFHIRDatastore(ctx, input)
			},
		},
		"delete-fhir-datastore": {
			Name:   "delete-fhir-datastore",
			Fields: fields_delete_fhir_datastore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFHIRDatastoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fhir_datastore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFHIRDatastore(ctx, input)
			},
		},
		"describe-fhir-datastore": {
			Name:   "describe-fhir-datastore",
			Fields: fields_describe_fhir_datastore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFHIRDatastoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fhir_datastore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFHIRDatastore(ctx, input)
			},
		},
		"describe-fhir-export-job": {
			Name:   "describe-fhir-export-job",
			Fields: fields_describe_fhir_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFHIRExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fhir_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFHIRExportJob(ctx, input)
			},
		},
		"describe-fhir-import-job": {
			Name:   "describe-fhir-import-job",
			Fields: fields_describe_fhir_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFHIRImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fhir_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFHIRImportJob(ctx, input)
			},
		},
		"list-fhir-datastores": {
			Name:   "list-fhir-datastores",
			Fields: fields_list_fhir_datastores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFHIRDatastoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fhir_datastores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFHIRDatastores(ctx, input)
				}
				var results []*svc.ListFHIRDatastoresOutput
				p := svc.NewListFHIRDatastoresPaginator(client, input)
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
		"list-fhir-export-jobs": {
			Name:   "list-fhir-export-jobs",
			Fields: fields_list_fhir_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFHIRExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fhir_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFHIRExportJobs(ctx, input)
				}
				var results []*svc.ListFHIRExportJobsOutput
				p := svc.NewListFHIRExportJobsPaginator(client, input)
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
		"list-fhir-import-jobs": {
			Name:   "list-fhir-import-jobs",
			Fields: fields_list_fhir_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFHIRImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fhir_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFHIRImportJobs(ctx, input)
				}
				var results []*svc.ListFHIRImportJobsOutput
				p := svc.NewListFHIRImportJobsPaginator(client, input)
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
		"start-fhir-export-job": {
			Name:   "start-fhir-export-job",
			Fields: fields_start_fhir_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFHIRExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_fhir_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFHIRExportJob(ctx, input)
			},
		},
		"start-fhir-import-job": {
			Name:   "start-fhir-import-job",
			Fields: fields_start_fhir_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFHIRImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_fhir_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFHIRImportJob(ctx, input)
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
	if err := leanruntime.Execute("healthlake", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
