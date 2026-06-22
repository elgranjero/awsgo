package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/medicalimaging"
)

var fields_copy_image_set = []leanruntime.Field{
	{Name: "CopyImageSetInformation", Flag: "copy-image-set-information", Type: "*types.CopyImageSetInformation", Required: true},
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "PromoteToPrimary", Flag: "promote-to-primary", Type: "*bool", Required: false},
	{Name: "SourceImageSetId", Flag: "source-image-set-id", Type: "*string", Required: true},
}

var fields_create_datastore = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatastoreName", Flag: "datastore-name", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "LambdaAuthorizerArn", Flag: "lambda-authorizer-arn", Type: "*string", Required: false},
	{Name: "LosslessStorageFormat", Flag: "lossless-storage-format", Type: "types.LosslessStorageFormat", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_datastore = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
}

var fields_delete_image_set = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "ImageSetId", Flag: "image-set-id", Type: "*string", Required: true},
}

var fields_get_datastore = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
}

var fields_get_dicom_import_job = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_image_frame = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "ImageFrameInformation", Flag: "image-frame-information", Type: "*types.ImageFrameInformation", Required: true},
	{Name: "ImageSetId", Flag: "image-set-id", Type: "*string", Required: true},
}

var fields_get_image_set = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "ImageSetId", Flag: "image-set-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_image_set_metadata = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "ImageSetId", Flag: "image-set-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_list_datastores = []leanruntime.Field{
	{Name: "DatastoreStatus", Flag: "datastore-status", Type: "types.DatastoreStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dicom_import_jobs = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "JobStatus", Flag: "job-status", Type: "types.JobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_image_set_versions = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "ImageSetId", Flag: "image-set-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_search_image_sets = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.SearchCriteria", Required: false},
}

var fields_start_dicom_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "InputOwnerAccountId", Flag: "input-owner-account-id", Type: "*string", Required: false},
	{Name: "InputS3Uri", Flag: "input-s3-uri", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputS3Uri", Flag: "output-s3-uri", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_image_set_metadata = []leanruntime.Field{
	{Name: "DatastoreId", Flag: "datastore-id", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "ImageSetId", Flag: "image-set-id", Type: "*string", Required: true},
	{Name: "LatestVersionId", Flag: "latest-version-id", Type: "*string", Required: true},
	{Name: "UpdateImageSetMetadataUpdates", Flag: "update-image-set-metadata-updates", Type: "types.MetadataUpdates", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"copy-image-set": {
			Name:   "copy-image-set",
			Fields: fields_copy_image_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyImageSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_image_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyImageSet(ctx, input)
			},
		},
		"create-datastore": {
			Name:   "create-datastore",
			Fields: fields_create_datastore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatastoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_datastore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDatastore(ctx, input)
			},
		},
		"delete-datastore": {
			Name:   "delete-datastore",
			Fields: fields_delete_datastore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatastoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_datastore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDatastore(ctx, input)
			},
		},
		"delete-image-set": {
			Name:   "delete-image-set",
			Fields: fields_delete_image_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImageSet(ctx, input)
			},
		},
		"get-datastore": {
			Name:   "get-datastore",
			Fields: fields_get_datastore,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDatastoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_datastore, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDatastore(ctx, input)
			},
		},
		"get-dicom-import-job": {
			Name:   "get-dicom-import-job",
			Fields: fields_get_dicom_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDICOMImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dicom_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDICOMImportJob(ctx, input)
			},
		},
		"get-image-frame": {
			Name:   "get-image-frame",
			Fields: fields_get_image_frame,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageFrameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_frame, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageFrame(ctx, input)
			},
		},
		"get-image-set": {
			Name:   "get-image-set",
			Fields: fields_get_image_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageSet(ctx, input)
			},
		},
		"get-image-set-metadata": {
			Name:   "get-image-set-metadata",
			Fields: fields_get_image_set_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageSetMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_set_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageSetMetadata(ctx, input)
			},
		},
		"list-datastores": {
			Name:   "list-datastores",
			Fields: fields_list_datastores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatastoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datastores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatastores(ctx, input)
				}
				var results []*svc.ListDatastoresOutput
				p := svc.NewListDatastoresPaginator(client, input)
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
		"list-dicom-import-jobs": {
			Name:   "list-dicom-import-jobs",
			Fields: fields_list_dicom_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDICOMImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dicom_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDICOMImportJobs(ctx, input)
				}
				var results []*svc.ListDICOMImportJobsOutput
				p := svc.NewListDICOMImportJobsPaginator(client, input)
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
		"list-image-set-versions": {
			Name:   "list-image-set-versions",
			Fields: fields_list_image_set_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageSetVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_image_set_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImageSetVersions(ctx, input)
				}
				var results []*svc.ListImageSetVersionsOutput
				p := svc.NewListImageSetVersionsPaginator(client, input)
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
		"search-image-sets": {
			Name:   "search-image-sets",
			Fields: fields_search_image_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchImageSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_image_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchImageSets(ctx, input)
				}
				var results []*svc.SearchImageSetsOutput
				p := svc.NewSearchImageSetsPaginator(client, input)
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
		"start-dicom-import-job": {
			Name:   "start-dicom-import-job",
			Fields: fields_start_dicom_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDICOMImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dicom_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDICOMImportJob(ctx, input)
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
		"update-image-set-metadata": {
			Name:   "update-image-set-metadata",
			Fields: fields_update_image_set_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImageSetMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image_set_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImageSetMetadata(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("medicalimaging", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
