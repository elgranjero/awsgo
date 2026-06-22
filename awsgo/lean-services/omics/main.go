package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/omics"
)

var fields_abort_multipart_read_set_upload = []leanruntime.Field{
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_accept_share = []leanruntime.Field{
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
}

var fields_batch_delete_read_set = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_cancel_annotation_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_cancel_run = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_cancel_variant_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_complete_multipart_read_set_upload = []leanruntime.Field{
	{Name: "Parts", Flag: "parts", Type: "[]types.CompleteReadSetUploadPartListItem", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_create_annotation_store = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Reference", Flag: "reference", Type: "types.ReferenceItem", Required: false},
	{Name: "SseConfig", Flag: "sse-config", Type: "*types.SseConfig", Required: false},
	{Name: "StoreFormat", Flag: "store-format", Type: "types.StoreFormat", Required: true},
	{Name: "StoreOptions", Flag: "store-options", Type: "types.StoreOptions", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_create_annotation_store_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
	{Name: "VersionOptions", Flag: "version-options", Type: "types.VersionOptions", Required: false},
}

var fields_create_multipart_read_set_upload = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GeneratedFrom", Flag: "generated-from", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReferenceArn", Flag: "reference-arn", Type: "*string", Required: false},
	{Name: "SampleId", Flag: "sample-id", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "SourceFileType", Flag: "source-file-type", Type: "types.FileType", Required: true},
	{Name: "SubjectId", Flag: "subject-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_reference_store = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SseConfig", Flag: "sse-config", Type: "*types.SseConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_run_cache = []leanruntime.Field{
	{Name: "CacheBehavior", Flag: "cache-behavior", Type: "types.CacheBehavior", Required: false},
	{Name: "CacheBucketOwnerId", Flag: "cache-bucket-owner-id", Type: "*string", Required: false},
	{Name: "CacheS3Location", Flag: "cache-s3-location", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_run_group = []leanruntime.Field{
	{Name: "MaxCpus", Flag: "max-cpus", Type: "*int32", Required: false},
	{Name: "MaxDuration", Flag: "max-duration", Type: "*int32", Required: false},
	{Name: "MaxGpus", Flag: "max-gpus", Type: "*int32", Required: false},
	{Name: "MaxRuns", Flag: "max-runs", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sequence_store = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ETagAlgorithmFamily", Flag: "etag-algorithm-family", Type: "types.ETagAlgorithmFamily", Required: false},
	{Name: "FallbackLocation", Flag: "fallback-location", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PropagatedSetLevelTags", Flag: "propagated-set-level-tags", Type: "[]string", Required: false},
	{Name: "S3AccessConfig", Flag: "s3-access-config", Type: "*types.S3AccessConfig", Required: false},
	{Name: "SseConfig", Flag: "sse-config", Type: "*types.SseConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_share = []leanruntime.Field{
	{Name: "PrincipalSubscriber", Flag: "principal-subscriber", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ShareName", Flag: "share-name", Type: "*string", Required: false},
}

var fields_create_variant_store = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Reference", Flag: "reference", Type: "types.ReferenceItem", Required: true},
	{Name: "SseConfig", Flag: "sse-config", Type: "*types.SseConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_workflow = []leanruntime.Field{
	{Name: "Accelerators", Flag: "accelerators", Type: "types.Accelerators", Required: false},
	{Name: "ContainerRegistryMap", Flag: "container-registry-map", Type: "*types.ContainerRegistryMap", Required: false},
	{Name: "ContainerRegistryMapUri", Flag: "container-registry-map-uri", Type: "*string", Required: false},
	{Name: "DefinitionRepository", Flag: "definition-repository", Type: "*types.DefinitionRepository", Required: false},
	{Name: "DefinitionUri", Flag: "definition-uri", Type: "*string", Required: false},
	{Name: "DefinitionZip", Flag: "definition-zip", Type: "[]byte", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "types.WorkflowEngine", Required: false},
	{Name: "Main", Flag: "main", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParameterTemplate", Flag: "parameter-template", Type: "map[string]types.WorkflowParameter", Required: false},
	{Name: "ParameterTemplatePath", Flag: "parameter-template-path", Type: "*string", Required: false},
	{Name: "ReadmeMarkdown", Flag: "readme-markdown", Type: "*string", Required: false},
	{Name: "ReadmePath", Flag: "readme-path", Type: "*string", Required: false},
	{Name: "ReadmeUri", Flag: "readme-uri", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkflowBucketOwnerId", Flag: "workflow-bucket-owner-id", Type: "*string", Required: false},
}

var fields_create_workflow_version = []leanruntime.Field{
	{Name: "Accelerators", Flag: "accelerators", Type: "types.Accelerators", Required: false},
	{Name: "ContainerRegistryMap", Flag: "container-registry-map", Type: "*types.ContainerRegistryMap", Required: false},
	{Name: "ContainerRegistryMapUri", Flag: "container-registry-map-uri", Type: "*string", Required: false},
	{Name: "DefinitionRepository", Flag: "definition-repository", Type: "*types.DefinitionRepository", Required: false},
	{Name: "DefinitionUri", Flag: "definition-uri", Type: "*string", Required: false},
	{Name: "DefinitionZip", Flag: "definition-zip", Type: "[]byte", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Engine", Flag: "engine", Type: "types.WorkflowEngine", Required: false},
	{Name: "Main", Flag: "main", Type: "*string", Required: false},
	{Name: "ParameterTemplate", Flag: "parameter-template", Type: "map[string]types.WorkflowParameter", Required: false},
	{Name: "ParameterTemplatePath", Flag: "parameter-template-path", Type: "*string", Required: false},
	{Name: "ReadmeMarkdown", Flag: "readme-markdown", Type: "*string", Required: false},
	{Name: "ReadmePath", Flag: "readme-path", Type: "*string", Required: false},
	{Name: "ReadmeUri", Flag: "readme-uri", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
	{Name: "WorkflowBucketOwnerId", Flag: "workflow-bucket-owner-id", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_delete_annotation_store = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_annotation_store_versions = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: true},
}

var fields_delete_reference = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
}

var fields_delete_reference_store = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_run = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_run_cache = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_run_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_s3_access_policy = []leanruntime.Field{
	{Name: "S3AccessPointArn", Flag: "s3-access-point-arn", Type: "*string", Required: true},
}

var fields_delete_sequence_store = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_share = []leanruntime.Field{
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
}

var fields_delete_variant_store = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_workflow_version = []leanruntime.Field{
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_get_annotation_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_annotation_store = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_annotation_store_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_get_read_set = []leanruntime.Field{
	{Name: "File", Flag: "file", Type: "types.ReadSetFile", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_get_read_set_activation_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_get_read_set_export_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_get_read_set_import_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_get_read_set_metadata = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_get_reference = []leanruntime.Field{
	{Name: "File", Flag: "file", Type: "types.ReferenceFile", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: true},
	{Name: "Range", Flag: "range", Type: "*string", Required: false},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
}

var fields_get_reference_import_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
}

var fields_get_reference_metadata = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
}

var fields_get_reference_store = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_run = []leanruntime.Field{
	{Name: "Export", Flag: "export", Type: "[]types.RunExport", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_run_cache = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_run_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_run_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_get_s3_access_policy = []leanruntime.Field{
	{Name: "S3AccessPointArn", Flag: "s3-access-point-arn", Type: "*string", Required: true},
}

var fields_get_sequence_store = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_share = []leanruntime.Field{
	{Name: "ShareId", Flag: "share-id", Type: "*string", Required: true},
}

var fields_get_variant_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_variant_store = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "Export", Flag: "export", Type: "[]types.WorkflowExport", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.WorkflowType", Required: false},
	{Name: "WorkflowOwnerId", Flag: "workflow-owner-id", Type: "*string", Required: false},
}

var fields_get_workflow_version = []leanruntime.Field{
	{Name: "Export", Flag: "export", Type: "[]types.WorkflowExport", Required: false},
	{Name: "Type", Flag: "type", Type: "types.WorkflowType", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
	{Name: "WorkflowOwnerId", Flag: "workflow-owner-id", Type: "*string", Required: false},
}

var fields_list_annotation_import_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListAnnotationImportJobsFilter", Required: false},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_annotation_store_versions = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListAnnotationStoreVersionsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_annotation_stores = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListAnnotationStoresFilter", Required: false},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_multipart_read_set_uploads = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_list_read_set_activation_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ActivateReadSetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_list_read_set_export_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ExportReadSetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_list_read_set_import_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ImportReadSetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_list_read_set_upload_parts = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReadSetUploadPartListFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PartSource", Flag: "part-source", Type: "types.ReadSetPartSource", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_list_read_sets = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReadSetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
}

var fields_list_reference_import_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ImportReferenceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
}

var fields_list_reference_stores = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReferenceStoreFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_references = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReferenceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
}

var fields_list_run_caches = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
}

var fields_list_run_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
}

var fields_list_run_tasks = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskStatus", Required: false},
}

var fields_list_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RunGroupId", Flag: "run-group-id", Type: "*string", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RunStatus", Required: false},
}

var fields_list_sequence_stores = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.SequenceStoreFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_shares = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_variant_import_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListVariantImportJobsFilter", Required: false},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_variant_stores = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListVariantStoresFilter", Required: false},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflow_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.WorkflowType", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
	{Name: "WorkflowOwnerId", Flag: "workflow-owner-id", Type: "*string", Required: false},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StartingToken", Flag: "starting-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.WorkflowType", Required: false},
}

var fields_put_s3_access_policy = []leanruntime.Field{
	{Name: "S3AccessPointArn", Flag: "s3-access-point-arn", Type: "*string", Required: true},
	{Name: "S3AccessPolicy", Flag: "s3-access-policy", Type: "*string", Required: true},
}

var fields_start_annotation_import_job = []leanruntime.Field{
	{Name: "AnnotationFields", Flag: "annotation-fields", Type: "map[string]string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "FormatOptions", Flag: "format-options", Type: "types.FormatOptions", Required: false},
	{Name: "Items", Flag: "items", Type: "[]types.AnnotationImportItemSource", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "RunLeftNormalization", Flag: "run-left-normalization", Type: "bool", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_start_read_set_activation_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.StartReadSetActivationJobSourceItem", Required: true},
}

var fields_start_read_set_export_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.ExportReadSet", Required: true},
}

var fields_start_read_set_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.StartReadSetImportJobSourceItem", Required: true},
}

var fields_start_reference_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ReferenceStoreId", Flag: "reference-store-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.StartReferenceImportJobSourceItem", Required: true},
}

var fields_start_run = []leanruntime.Field{
	{Name: "CacheBehavior", Flag: "cache-behavior", Type: "types.CacheBehavior", Required: false},
	{Name: "CacheId", Flag: "cache-id", Type: "*string", Required: false},
	{Name: "LogLevel", Flag: "log-level", Type: "types.RunLogLevel", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OutputUri", Flag: "output-uri", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "document.Interface", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "RetentionMode", Flag: "retention-mode", Type: "types.RunRetentionMode", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "RunGroupId", Flag: "run-group-id", Type: "*string", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: false},
	{Name: "WorkflowOwnerId", Flag: "workflow-owner-id", Type: "*string", Required: false},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "types.WorkflowType", Required: false},
	{Name: "WorkflowVersionName", Flag: "workflow-version-name", Type: "*string", Required: false},
}

var fields_start_variant_import_job = []leanruntime.Field{
	{Name: "AnnotationFields", Flag: "annotation-fields", Type: "map[string]string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "Items", Flag: "items", Type: "[]types.VariantImportItemSource", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "RunLeftNormalization", Flag: "run-left-normalization", Type: "bool", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_annotation_store = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_annotation_store_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_update_run_cache = []leanruntime.Field{
	{Name: "CacheBehavior", Flag: "cache-behavior", Type: "types.CacheBehavior", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_run_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxCpus", Flag: "max-cpus", Type: "*int32", Required: false},
	{Name: "MaxDuration", Flag: "max-duration", Type: "*int32", Required: false},
	{Name: "MaxGpus", Flag: "max-gpus", Type: "*int32", Required: false},
	{Name: "MaxRuns", Flag: "max-runs", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_sequence_store = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FallbackLocation", Flag: "fallback-location", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PropagatedSetLevelTags", Flag: "propagated-set-level-tags", Type: "[]string", Required: false},
	{Name: "S3AccessConfig", Flag: "s3-access-config", Type: "*types.S3AccessConfig", Required: false},
}

var fields_update_variant_store = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ReadmeMarkdown", Flag: "readme-markdown", Type: "*string", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
}

var fields_update_workflow_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ReadmeMarkdown", Flag: "readme-markdown", Type: "*string", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_upload_read_set_part = []leanruntime.Field{
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: true},
	{Name: "PartSource", Flag: "part-source", Type: "types.ReadSetPartSource", Required: true},
	{Name: "Payload", Flag: "payload", Type: "io.Reader", Required: true},
	{Name: "SequenceStoreId", Flag: "sequence-store-id", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"abort-multipart-read-set-upload": {
			Name:   "abort-multipart-read-set-upload",
			Fields: fields_abort_multipart_read_set_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AbortMultipartReadSetUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_abort_multipart_read_set_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AbortMultipartReadSetUpload(ctx, input)
			},
		},
		"accept-share": {
			Name:   "accept-share",
			Fields: fields_accept_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptShare(ctx, input)
			},
		},
		"batch-delete-read-set": {
			Name:   "batch-delete-read-set",
			Fields: fields_batch_delete_read_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteReadSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_read_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteReadSet(ctx, input)
			},
		},
		"cancel-annotation-import-job": {
			Name:   "cancel-annotation-import-job",
			Fields: fields_cancel_annotation_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelAnnotationImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_annotation_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelAnnotationImportJob(ctx, input)
			},
		},
		"cancel-run": {
			Name:   "cancel-run",
			Fields: fields_cancel_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelRun(ctx, input)
			},
		},
		"cancel-variant-import-job": {
			Name:   "cancel-variant-import-job",
			Fields: fields_cancel_variant_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelVariantImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_variant_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelVariantImportJob(ctx, input)
			},
		},
		"complete-multipart-read-set-upload": {
			Name:   "complete-multipart-read-set-upload",
			Fields: fields_complete_multipart_read_set_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteMultipartReadSetUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_multipart_read_set_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteMultipartReadSetUpload(ctx, input)
			},
		},
		"create-annotation-store": {
			Name:   "create-annotation-store",
			Fields: fields_create_annotation_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnnotationStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_annotation_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnnotationStore(ctx, input)
			},
		},
		"create-annotation-store-version": {
			Name:   "create-annotation-store-version",
			Fields: fields_create_annotation_store_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnnotationStoreVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_annotation_store_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnnotationStoreVersion(ctx, input)
			},
		},
		"create-multipart-read-set-upload": {
			Name:   "create-multipart-read-set-upload",
			Fields: fields_create_multipart_read_set_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultipartReadSetUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multipart_read_set_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultipartReadSetUpload(ctx, input)
			},
		},
		"create-reference-store": {
			Name:   "create-reference-store",
			Fields: fields_create_reference_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReferenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_reference_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReferenceStore(ctx, input)
			},
		},
		"create-run-cache": {
			Name:   "create-run-cache",
			Fields: fields_create_run_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRunCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_run_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRunCache(ctx, input)
			},
		},
		"create-run-group": {
			Name:   "create-run-group",
			Fields: fields_create_run_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRunGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_run_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRunGroup(ctx, input)
			},
		},
		"create-sequence-store": {
			Name:   "create-sequence-store",
			Fields: fields_create_sequence_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSequenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sequence_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSequenceStore(ctx, input)
			},
		},
		"create-share": {
			Name:   "create-share",
			Fields: fields_create_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateShare(ctx, input)
			},
		},
		"create-variant-store": {
			Name:   "create-variant-store",
			Fields: fields_create_variant_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVariantStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_variant_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVariantStore(ctx, input)
			},
		},
		"create-workflow": {
			Name:   "create-workflow",
			Fields: fields_create_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflow(ctx, input)
			},
		},
		"create-workflow-version": {
			Name:   "create-workflow-version",
			Fields: fields_create_workflow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflowVersion(ctx, input)
			},
		},
		"delete-annotation-store": {
			Name:   "delete-annotation-store",
			Fields: fields_delete_annotation_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnnotationStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_annotation_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnnotationStore(ctx, input)
			},
		},
		"delete-annotation-store-versions": {
			Name:   "delete-annotation-store-versions",
			Fields: fields_delete_annotation_store_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnnotationStoreVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_annotation_store_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnnotationStoreVersions(ctx, input)
			},
		},
		"delete-reference": {
			Name:   "delete-reference",
			Fields: fields_delete_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReference(ctx, input)
			},
		},
		"delete-reference-store": {
			Name:   "delete-reference-store",
			Fields: fields_delete_reference_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReferenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_reference_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReferenceStore(ctx, input)
			},
		},
		"delete-run": {
			Name:   "delete-run",
			Fields: fields_delete_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRun(ctx, input)
			},
		},
		"delete-run-cache": {
			Name:   "delete-run-cache",
			Fields: fields_delete_run_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRunCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_run_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRunCache(ctx, input)
			},
		},
		"delete-run-group": {
			Name:   "delete-run-group",
			Fields: fields_delete_run_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRunGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_run_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRunGroup(ctx, input)
			},
		},
		"delete-s3-access-policy": {
			Name:   "delete-s3-access-policy",
			Fields: fields_delete_s3_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteS3AccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_s3_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteS3AccessPolicy(ctx, input)
			},
		},
		"delete-sequence-store": {
			Name:   "delete-sequence-store",
			Fields: fields_delete_sequence_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSequenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sequence_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSequenceStore(ctx, input)
			},
		},
		"delete-share": {
			Name:   "delete-share",
			Fields: fields_delete_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteShare(ctx, input)
			},
		},
		"delete-variant-store": {
			Name:   "delete-variant-store",
			Fields: fields_delete_variant_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVariantStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_variant_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVariantStore(ctx, input)
			},
		},
		"delete-workflow": {
			Name:   "delete-workflow",
			Fields: fields_delete_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflow(ctx, input)
			},
		},
		"delete-workflow-version": {
			Name:   "delete-workflow-version",
			Fields: fields_delete_workflow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflowVersion(ctx, input)
			},
		},
		"get-annotation-import-job": {
			Name:   "get-annotation-import-job",
			Fields: fields_get_annotation_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnnotationImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_annotation_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnnotationImportJob(ctx, input)
			},
		},
		"get-annotation-store": {
			Name:   "get-annotation-store",
			Fields: fields_get_annotation_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnnotationStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_annotation_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnnotationStore(ctx, input)
			},
		},
		"get-annotation-store-version": {
			Name:   "get-annotation-store-version",
			Fields: fields_get_annotation_store_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAnnotationStoreVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_annotation_store_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAnnotationStoreVersion(ctx, input)
			},
		},
		"get-read-set": {
			Name:   "get-read-set",
			Fields: fields_get_read_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_read_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReadSet(ctx, input)
			},
		},
		"get-read-set-activation-job": {
			Name:   "get-read-set-activation-job",
			Fields: fields_get_read_set_activation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadSetActivationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_read_set_activation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReadSetActivationJob(ctx, input)
			},
		},
		"get-read-set-export-job": {
			Name:   "get-read-set-export-job",
			Fields: fields_get_read_set_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadSetExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_read_set_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReadSetExportJob(ctx, input)
			},
		},
		"get-read-set-import-job": {
			Name:   "get-read-set-import-job",
			Fields: fields_get_read_set_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadSetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_read_set_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReadSetImportJob(ctx, input)
			},
		},
		"get-read-set-metadata": {
			Name:   "get-read-set-metadata",
			Fields: fields_get_read_set_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadSetMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_read_set_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReadSetMetadata(ctx, input)
			},
		},
		"get-reference": {
			Name:   "get-reference",
			Fields: fields_get_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReference(ctx, input)
			},
		},
		"get-reference-import-job": {
			Name:   "get-reference-import-job",
			Fields: fields_get_reference_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReferenceImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reference_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReferenceImportJob(ctx, input)
			},
		},
		"get-reference-metadata": {
			Name:   "get-reference-metadata",
			Fields: fields_get_reference_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReferenceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reference_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReferenceMetadata(ctx, input)
			},
		},
		"get-reference-store": {
			Name:   "get-reference-store",
			Fields: fields_get_reference_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReferenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reference_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReferenceStore(ctx, input)
			},
		},
		"get-run": {
			Name:   "get-run",
			Fields: fields_get_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRun(ctx, input)
			},
		},
		"get-run-cache": {
			Name:   "get-run-cache",
			Fields: fields_get_run_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRunCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_run_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRunCache(ctx, input)
			},
		},
		"get-run-group": {
			Name:   "get-run-group",
			Fields: fields_get_run_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRunGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_run_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRunGroup(ctx, input)
			},
		},
		"get-run-task": {
			Name:   "get-run-task",
			Fields: fields_get_run_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRunTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_run_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRunTask(ctx, input)
			},
		},
		"get-s3-access-policy": {
			Name:   "get-s3-access-policy",
			Fields: fields_get_s3_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetS3AccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_s3_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetS3AccessPolicy(ctx, input)
			},
		},
		"get-sequence-store": {
			Name:   "get-sequence-store",
			Fields: fields_get_sequence_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSequenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sequence_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSequenceStore(ctx, input)
			},
		},
		"get-share": {
			Name:   "get-share",
			Fields: fields_get_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetShare(ctx, input)
			},
		},
		"get-variant-import-job": {
			Name:   "get-variant-import-job",
			Fields: fields_get_variant_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVariantImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_variant_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVariantImportJob(ctx, input)
			},
		},
		"get-variant-store": {
			Name:   "get-variant-store",
			Fields: fields_get_variant_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVariantStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_variant_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVariantStore(ctx, input)
			},
		},
		"get-workflow": {
			Name:   "get-workflow",
			Fields: fields_get_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflow(ctx, input)
			},
		},
		"get-workflow-version": {
			Name:   "get-workflow-version",
			Fields: fields_get_workflow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowVersion(ctx, input)
			},
		},
		"list-annotation-import-jobs": {
			Name:   "list-annotation-import-jobs",
			Fields: fields_list_annotation_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnnotationImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_annotation_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnnotationImportJobs(ctx, input)
				}
				var results []*svc.ListAnnotationImportJobsOutput
				p := svc.NewListAnnotationImportJobsPaginator(client, input)
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
		"list-annotation-store-versions": {
			Name:   "list-annotation-store-versions",
			Fields: fields_list_annotation_store_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnnotationStoreVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_annotation_store_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnnotationStoreVersions(ctx, input)
				}
				var results []*svc.ListAnnotationStoreVersionsOutput
				p := svc.NewListAnnotationStoreVersionsPaginator(client, input)
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
		"list-annotation-stores": {
			Name:   "list-annotation-stores",
			Fields: fields_list_annotation_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnnotationStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_annotation_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnnotationStores(ctx, input)
				}
				var results []*svc.ListAnnotationStoresOutput
				p := svc.NewListAnnotationStoresPaginator(client, input)
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
		"list-multipart-read-set-uploads": {
			Name:   "list-multipart-read-set-uploads",
			Fields: fields_list_multipart_read_set_uploads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultipartReadSetUploadsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multipart_read_set_uploads, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultipartReadSetUploads(ctx, input)
				}
				var results []*svc.ListMultipartReadSetUploadsOutput
				p := svc.NewListMultipartReadSetUploadsPaginator(client, input)
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
		"list-read-set-activation-jobs": {
			Name:   "list-read-set-activation-jobs",
			Fields: fields_list_read_set_activation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReadSetActivationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_read_set_activation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReadSetActivationJobs(ctx, input)
				}
				var results []*svc.ListReadSetActivationJobsOutput
				p := svc.NewListReadSetActivationJobsPaginator(client, input)
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
		"list-read-set-export-jobs": {
			Name:   "list-read-set-export-jobs",
			Fields: fields_list_read_set_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReadSetExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_read_set_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReadSetExportJobs(ctx, input)
				}
				var results []*svc.ListReadSetExportJobsOutput
				p := svc.NewListReadSetExportJobsPaginator(client, input)
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
		"list-read-set-import-jobs": {
			Name:   "list-read-set-import-jobs",
			Fields: fields_list_read_set_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReadSetImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_read_set_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReadSetImportJobs(ctx, input)
				}
				var results []*svc.ListReadSetImportJobsOutput
				p := svc.NewListReadSetImportJobsPaginator(client, input)
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
		"list-read-set-upload-parts": {
			Name:   "list-read-set-upload-parts",
			Fields: fields_list_read_set_upload_parts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReadSetUploadPartsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_read_set_upload_parts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReadSetUploadParts(ctx, input)
				}
				var results []*svc.ListReadSetUploadPartsOutput
				p := svc.NewListReadSetUploadPartsPaginator(client, input)
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
		"list-read-sets": {
			Name:   "list-read-sets",
			Fields: fields_list_read_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReadSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_read_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReadSets(ctx, input)
				}
				var results []*svc.ListReadSetsOutput
				p := svc.NewListReadSetsPaginator(client, input)
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
		"list-reference-import-jobs": {
			Name:   "list-reference-import-jobs",
			Fields: fields_list_reference_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReferenceImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reference_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReferenceImportJobs(ctx, input)
				}
				var results []*svc.ListReferenceImportJobsOutput
				p := svc.NewListReferenceImportJobsPaginator(client, input)
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
		"list-reference-stores": {
			Name:   "list-reference-stores",
			Fields: fields_list_reference_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReferenceStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reference_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReferenceStores(ctx, input)
				}
				var results []*svc.ListReferenceStoresOutput
				p := svc.NewListReferenceStoresPaginator(client, input)
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
		"list-references": {
			Name:   "list-references",
			Fields: fields_list_references,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_references, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReferences(ctx, input)
				}
				var results []*svc.ListReferencesOutput
				p := svc.NewListReferencesPaginator(client, input)
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
		"list-run-caches": {
			Name:   "list-run-caches",
			Fields: fields_list_run_caches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRunCachesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_run_caches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRunCaches(ctx, input)
				}
				var results []*svc.ListRunCachesOutput
				p := svc.NewListRunCachesPaginator(client, input)
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
		"list-run-groups": {
			Name:   "list-run-groups",
			Fields: fields_list_run_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRunGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_run_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRunGroups(ctx, input)
				}
				var results []*svc.ListRunGroupsOutput
				p := svc.NewListRunGroupsPaginator(client, input)
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
		"list-run-tasks": {
			Name:   "list-run-tasks",
			Fields: fields_list_run_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRunTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_run_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRunTasks(ctx, input)
				}
				var results []*svc.ListRunTasksOutput
				p := svc.NewListRunTasksPaginator(client, input)
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
		"list-runs": {
			Name:   "list-runs",
			Fields: fields_list_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuns(ctx, input)
				}
				var results []*svc.ListRunsOutput
				p := svc.NewListRunsPaginator(client, input)
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
		"list-sequence-stores": {
			Name:   "list-sequence-stores",
			Fields: fields_list_sequence_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSequenceStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sequence_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSequenceStores(ctx, input)
				}
				var results []*svc.ListSequenceStoresOutput
				p := svc.NewListSequenceStoresPaginator(client, input)
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
		"list-shares": {
			Name:   "list-shares",
			Fields: fields_list_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListShares(ctx, input)
				}
				var results []*svc.ListSharesOutput
				p := svc.NewListSharesPaginator(client, input)
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
		"list-variant-import-jobs": {
			Name:   "list-variant-import-jobs",
			Fields: fields_list_variant_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVariantImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_variant_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVariantImportJobs(ctx, input)
				}
				var results []*svc.ListVariantImportJobsOutput
				p := svc.NewListVariantImportJobsPaginator(client, input)
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
		"list-variant-stores": {
			Name:   "list-variant-stores",
			Fields: fields_list_variant_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVariantStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_variant_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVariantStores(ctx, input)
				}
				var results []*svc.ListVariantStoresOutput
				p := svc.NewListVariantStoresPaginator(client, input)
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
		"list-workflow-versions": {
			Name:   "list-workflow-versions",
			Fields: fields_list_workflow_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowVersions(ctx, input)
				}
				var results []*svc.ListWorkflowVersionsOutput
				p := svc.NewListWorkflowVersionsPaginator(client, input)
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
		"list-workflows": {
			Name:   "list-workflows",
			Fields: fields_list_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflows(ctx, input)
				}
				var results []*svc.ListWorkflowsOutput
				p := svc.NewListWorkflowsPaginator(client, input)
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
		"put-s3-access-policy": {
			Name:   "put-s3-access-policy",
			Fields: fields_put_s3_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutS3AccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_s3_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutS3AccessPolicy(ctx, input)
			},
		},
		"start-annotation-import-job": {
			Name:   "start-annotation-import-job",
			Fields: fields_start_annotation_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAnnotationImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_annotation_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAnnotationImportJob(ctx, input)
			},
		},
		"start-read-set-activation-job": {
			Name:   "start-read-set-activation-job",
			Fields: fields_start_read_set_activation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReadSetActivationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_read_set_activation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReadSetActivationJob(ctx, input)
			},
		},
		"start-read-set-export-job": {
			Name:   "start-read-set-export-job",
			Fields: fields_start_read_set_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReadSetExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_read_set_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReadSetExportJob(ctx, input)
			},
		},
		"start-read-set-import-job": {
			Name:   "start-read-set-import-job",
			Fields: fields_start_read_set_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReadSetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_read_set_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReadSetImportJob(ctx, input)
			},
		},
		"start-reference-import-job": {
			Name:   "start-reference-import-job",
			Fields: fields_start_reference_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReferenceImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_reference_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReferenceImportJob(ctx, input)
			},
		},
		"start-run": {
			Name:   "start-run",
			Fields: fields_start_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRun(ctx, input)
			},
		},
		"start-variant-import-job": {
			Name:   "start-variant-import-job",
			Fields: fields_start_variant_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVariantImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_variant_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVariantImportJob(ctx, input)
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
		"update-annotation-store": {
			Name:   "update-annotation-store",
			Fields: fields_update_annotation_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnnotationStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_annotation_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnnotationStore(ctx, input)
			},
		},
		"update-annotation-store-version": {
			Name:   "update-annotation-store-version",
			Fields: fields_update_annotation_store_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAnnotationStoreVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_annotation_store_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAnnotationStoreVersion(ctx, input)
			},
		},
		"update-run-cache": {
			Name:   "update-run-cache",
			Fields: fields_update_run_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRunCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_run_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRunCache(ctx, input)
			},
		},
		"update-run-group": {
			Name:   "update-run-group",
			Fields: fields_update_run_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRunGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_run_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRunGroup(ctx, input)
			},
		},
		"update-sequence-store": {
			Name:   "update-sequence-store",
			Fields: fields_update_sequence_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSequenceStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sequence_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSequenceStore(ctx, input)
			},
		},
		"update-variant-store": {
			Name:   "update-variant-store",
			Fields: fields_update_variant_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVariantStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_variant_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVariantStore(ctx, input)
			},
		},
		"update-workflow": {
			Name:   "update-workflow",
			Fields: fields_update_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflow(ctx, input)
			},
		},
		"update-workflow-version": {
			Name:   "update-workflow-version",
			Fields: fields_update_workflow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflowVersion(ctx, input)
			},
		},
		"upload-read-set-part": {
			Name:   "upload-read-set-part",
			Fields: fields_upload_read_set_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadReadSetPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_read_set_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadReadSetPart(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("omics", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
