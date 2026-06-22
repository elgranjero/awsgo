package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/s3control"
)

var fields_associate_access_grants_identity_center = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "IdentityCenterArn", Flag: "identity-center-arn", Type: "*string", Required: true},
}

var fields_create_access_grant = []leanruntime.Field{
	{Name: "AccessGrantsLocationConfiguration", Flag: "access-grants-location-configuration", Type: "*types.AccessGrantsLocationConfiguration", Required: false},
	{Name: "AccessGrantsLocationId", Flag: "access-grants-location-id", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: false},
	{Name: "Grantee", Flag: "grantee", Type: "*types.Grantee", Required: true},
	{Name: "Permission", Flag: "permission", Type: "types.Permission", Required: true},
	{Name: "S3PrefixType", Flag: "s3-prefix-type", Type: "types.S3PrefixType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_access_grants_instance = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "IdentityCenterArn", Flag: "identity-center-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_access_grants_location = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "IAMRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "LocationScope", Flag: "location-scope", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_access_point = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BucketAccountId", Flag: "bucket-account-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PublicAccessBlockConfiguration", Flag: "public-access-block-configuration", Type: "*types.PublicAccessBlockConfiguration", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: false},
}

var fields_create_access_point_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ObjectLambdaConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_bucket = []leanruntime.Field{
	{Name: "ACL", Flag: "acl", Type: "types.BucketCannedACL", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "CreateBucketConfiguration", Flag: "create-bucket-configuration", Type: "*types.CreateBucketConfiguration", Required: false},
	{Name: "GrantFullControl", Flag: "grant-full-control", Type: "*string", Required: false},
	{Name: "GrantRead", Flag: "grant-read", Type: "*string", Required: false},
	{Name: "GrantReadACP", Flag: "grant-read-acp", Type: "*string", Required: false},
	{Name: "GrantWrite", Flag: "grant-write", Type: "*string", Required: false},
	{Name: "GrantWriteACP", Flag: "grant-write-acp", Type: "*string", Required: false},
	{Name: "ObjectLockEnabledForBucket", Flag: "object-lock-enabled-for-bucket", Type: "bool", Required: false},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: false},
}

var fields_create_job = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ConfirmationRequired", Flag: "confirmation-required", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Manifest", Flag: "manifest", Type: "*types.JobManifest", Required: false},
	{Name: "ManifestGenerator", Flag: "manifest-generator", Type: "types.JobManifestGenerator", Required: false},
	{Name: "Operation", Flag: "operation", Type: "*types.JobOperation", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "Report", Flag: "report", Type: "*types.JobReport", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.S3Tag", Required: false},
}

var fields_create_multi_region_access_point = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Details", Flag: "details", Type: "*types.CreateMultiRegionAccessPointInput", Required: true},
}

var fields_create_storage_lens_group = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "StorageLensGroup", Flag: "storage-lens-group", Type: "*types.StorageLensGroup", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_access_grant = []leanruntime.Field{
	{Name: "AccessGrantId", Flag: "access-grant-id", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_delete_access_grants_instance = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_delete_access_grants_instance_resource_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_delete_access_grants_location = []leanruntime.Field{
	{Name: "AccessGrantsLocationId", Flag: "access-grants-location-id", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_delete_access_point = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_access_point_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_access_point_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_access_point_policy_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_access_point_scope = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_bucket = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_delete_bucket_lifecycle_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_delete_bucket_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_delete_bucket_replication = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_delete_bucket_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_delete_job_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_delete_multi_region_access_point = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Details", Flag: "details", Type: "*types.DeleteMultiRegionAccessPointInput", Required: true},
}

var fields_delete_public_access_block = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_delete_storage_lens_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
}

var fields_delete_storage_lens_configuration_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
}

var fields_delete_storage_lens_group = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_job = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_multi_region_access_point_operation = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "RequestTokenARN", Flag: "request-token-arn", Type: "*string", Required: true},
}

var fields_dissociate_access_grants_identity_center = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_access_grant = []leanruntime.Field{
	{Name: "AccessGrantId", Flag: "access-grant-id", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_access_grants_instance = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_access_grants_instance_for_prefix = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: true},
}

var fields_get_access_grants_instance_resource_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_access_grants_location = []leanruntime.Field{
	{Name: "AccessGrantsLocationId", Flag: "access-grants-location-id", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_access_point = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_configuration_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_policy_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_policy_status = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_policy_status_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_access_point_scope = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_bucket = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_get_bucket_lifecycle_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_get_bucket_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_get_bucket_replication = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_get_bucket_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_get_bucket_versioning = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
}

var fields_get_data_access = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Permission", Flag: "permission", Type: "types.Permission", Required: true},
	{Name: "Privilege", Flag: "privilege", Type: "types.Privilege", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "types.S3PrefixType", Required: false},
}

var fields_get_job_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_multi_region_access_point = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_multi_region_access_point_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_multi_region_access_point_policy_status = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_multi_region_access_point_routes = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Mrap", Flag: "mrap", Type: "*string", Required: true},
}

var fields_get_public_access_block = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_storage_lens_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
}

var fields_get_storage_lens_configuration_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
}

var fields_get_storage_lens_group = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_access_grants = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: false},
	{Name: "GrantScope", Flag: "grant-scope", Type: "*string", Required: false},
	{Name: "GranteeIdentifier", Flag: "grantee-identifier", Type: "*string", Required: false},
	{Name: "GranteeType", Flag: "grantee-type", Type: "types.GranteeType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Permission", Flag: "permission", Type: "types.Permission", Required: false},
}

var fields_list_access_grants_instances = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_grants_locations = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "LocationScope", Flag: "location-scope", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_points = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "DataSourceType", Flag: "data-source-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_points_for_directory_buckets = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "DirectoryBucket", Flag: "directory-bucket", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_points_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_caller_access_grants = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "AllowedByApplication", Flag: "allowed-by-application", Type: "bool", Required: false},
	{Name: "GrantScope", Flag: "grant-scope", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobStatuses", Flag: "job-statuses", Type: "[]types.JobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_multi_region_access_points = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_regional_buckets = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: false},
}

var fields_list_storage_lens_configurations = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_storage_lens_groups = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_access_grants_instance_resource_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Organization", Flag: "organization", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_access_point_configuration_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ObjectLambdaConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_put_access_point_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_access_point_policy_for_object_lambda = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_access_point_scope = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: true},
}

var fields_put_bucket_lifecycle_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "LifecycleConfiguration", Flag: "lifecycle-configuration", Type: "*types.LifecycleConfiguration", Required: false},
}

var fields_put_bucket_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ConfirmRemoveSelfBucketAccess", Flag: "confirm-remove-self-bucket-access", Type: "bool", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_bucket_replication = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ReplicationConfiguration", Flag: "replication-configuration", Type: "*types.ReplicationConfiguration", Required: true},
}

var fields_put_bucket_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "Tagging", Flag: "tagging", Type: "*types.Tagging", Required: true},
}

var fields_put_bucket_versioning = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "MFA", Flag: "mfa", Type: "*string", Required: false},
	{Name: "VersioningConfiguration", Flag: "versioning-configuration", Type: "*types.VersioningConfiguration", Required: true},
}

var fields_put_job_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.S3Tag", Required: true},
}

var fields_put_multi_region_access_point_policy = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Details", Flag: "details", Type: "*types.PutMultiRegionAccessPointPolicyInput", Required: true},
}

var fields_put_public_access_block = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "PublicAccessBlockConfiguration", Flag: "public-access-block-configuration", Type: "*types.PublicAccessBlockConfiguration", Required: true},
}

var fields_put_storage_lens_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
	{Name: "StorageLensConfiguration", Flag: "storage-lens-configuration", Type: "*types.StorageLensConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.StorageLensTag", Required: false},
}

var fields_put_storage_lens_configuration_tagging = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.StorageLensTag", Required: true},
}

var fields_submit_multi_region_access_point_routes = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Mrap", Flag: "mrap", Type: "*string", Required: true},
	{Name: "RouteUpdates", Flag: "route-updates", Type: "[]types.MultiRegionAccessPointRoute", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_grants_location = []leanruntime.Field{
	{Name: "AccessGrantsLocationId", Flag: "access-grants-location-id", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "IAMRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
}

var fields_update_job_priority = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "int32", Required: true},
}

var fields_update_job_status = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "RequestedJobStatus", Flag: "requested-job-status", Type: "types.RequestedJobStatus", Required: true},
	{Name: "StatusUpdateReason", Flag: "status-update-reason", Type: "*string", Required: false},
}

var fields_update_storage_lens_group = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StorageLensGroup", Flag: "storage-lens-group", Type: "*types.StorageLensGroup", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-access-grants-identity-center": {
			Name:   "associate-access-grants-identity-center",
			Fields: fields_associate_access_grants_identity_center,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAccessGrantsIdentityCenterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_access_grants_identity_center, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAccessGrantsIdentityCenter(ctx, input)
			},
		},
		"create-access-grant": {
			Name:   "create-access-grant",
			Fields: fields_create_access_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessGrant(ctx, input)
			},
		},
		"create-access-grants-instance": {
			Name:   "create-access-grants-instance",
			Fields: fields_create_access_grants_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessGrantsInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_grants_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessGrantsInstance(ctx, input)
			},
		},
		"create-access-grants-location": {
			Name:   "create-access-grants-location",
			Fields: fields_create_access_grants_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessGrantsLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_grants_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessGrantsLocation(ctx, input)
			},
		},
		"create-access-point": {
			Name:   "create-access-point",
			Fields: fields_create_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessPoint(ctx, input)
			},
		},
		"create-access-point-for-object-lambda": {
			Name:   "create-access-point-for-object-lambda",
			Fields: fields_create_access_point_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessPointForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_point_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessPointForObjectLambda(ctx, input)
			},
		},
		"create-bucket": {
			Name:   "create-bucket",
			Fields: fields_create_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBucket(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-multi-region-access-point": {
			Name:   "create-multi-region-access-point",
			Fields: fields_create_multi_region_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultiRegionAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multi_region_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultiRegionAccessPoint(ctx, input)
			},
		},
		"create-storage-lens-group": {
			Name:   "create-storage-lens-group",
			Fields: fields_create_storage_lens_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStorageLensGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_storage_lens_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStorageLensGroup(ctx, input)
			},
		},
		"delete-access-grant": {
			Name:   "delete-access-grant",
			Fields: fields_delete_access_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessGrant(ctx, input)
			},
		},
		"delete-access-grants-instance": {
			Name:   "delete-access-grants-instance",
			Fields: fields_delete_access_grants_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessGrantsInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_grants_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessGrantsInstance(ctx, input)
			},
		},
		"delete-access-grants-instance-resource-policy": {
			Name:   "delete-access-grants-instance-resource-policy",
			Fields: fields_delete_access_grants_instance_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessGrantsInstanceResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_grants_instance_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessGrantsInstanceResourcePolicy(ctx, input)
			},
		},
		"delete-access-grants-location": {
			Name:   "delete-access-grants-location",
			Fields: fields_delete_access_grants_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessGrantsLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_grants_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessGrantsLocation(ctx, input)
			},
		},
		"delete-access-point": {
			Name:   "delete-access-point",
			Fields: fields_delete_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPoint(ctx, input)
			},
		},
		"delete-access-point-for-object-lambda": {
			Name:   "delete-access-point-for-object-lambda",
			Fields: fields_delete_access_point_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPointForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_point_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPointForObjectLambda(ctx, input)
			},
		},
		"delete-access-point-policy": {
			Name:   "delete-access-point-policy",
			Fields: fields_delete_access_point_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_point_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPointPolicy(ctx, input)
			},
		},
		"delete-access-point-policy-for-object-lambda": {
			Name:   "delete-access-point-policy-for-object-lambda",
			Fields: fields_delete_access_point_policy_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPointPolicyForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_point_policy_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPointPolicyForObjectLambda(ctx, input)
			},
		},
		"delete-access-point-scope": {
			Name:   "delete-access-point-scope",
			Fields: fields_delete_access_point_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPointScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_point_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPointScope(ctx, input)
			},
		},
		"delete-bucket": {
			Name:   "delete-bucket",
			Fields: fields_delete_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucket(ctx, input)
			},
		},
		"delete-bucket-lifecycle-configuration": {
			Name:   "delete-bucket-lifecycle-configuration",
			Fields: fields_delete_bucket_lifecycle_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketLifecycleConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_lifecycle_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketLifecycleConfiguration(ctx, input)
			},
		},
		"delete-bucket-policy": {
			Name:   "delete-bucket-policy",
			Fields: fields_delete_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketPolicy(ctx, input)
			},
		},
		"delete-bucket-replication": {
			Name:   "delete-bucket-replication",
			Fields: fields_delete_bucket_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketReplication(ctx, input)
			},
		},
		"delete-bucket-tagging": {
			Name:   "delete-bucket-tagging",
			Fields: fields_delete_bucket_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketTagging(ctx, input)
			},
		},
		"delete-job-tagging": {
			Name:   "delete-job-tagging",
			Fields: fields_delete_job_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJobTagging(ctx, input)
			},
		},
		"delete-multi-region-access-point": {
			Name:   "delete-multi-region-access-point",
			Fields: fields_delete_multi_region_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMultiRegionAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_multi_region_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMultiRegionAccessPoint(ctx, input)
			},
		},
		"delete-public-access-block": {
			Name:   "delete-public-access-block",
			Fields: fields_delete_public_access_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePublicAccessBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_public_access_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePublicAccessBlock(ctx, input)
			},
		},
		"delete-storage-lens-configuration": {
			Name:   "delete-storage-lens-configuration",
			Fields: fields_delete_storage_lens_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStorageLensConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_storage_lens_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStorageLensConfiguration(ctx, input)
			},
		},
		"delete-storage-lens-configuration-tagging": {
			Name:   "delete-storage-lens-configuration-tagging",
			Fields: fields_delete_storage_lens_configuration_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStorageLensConfigurationTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_storage_lens_configuration_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStorageLensConfigurationTagging(ctx, input)
			},
		},
		"delete-storage-lens-group": {
			Name:   "delete-storage-lens-group",
			Fields: fields_delete_storage_lens_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStorageLensGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_storage_lens_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStorageLensGroup(ctx, input)
			},
		},
		"describe-job": {
			Name:   "describe-job",
			Fields: fields_describe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJob(ctx, input)
			},
		},
		"describe-multi-region-access-point-operation": {
			Name:   "describe-multi-region-access-point-operation",
			Fields: fields_describe_multi_region_access_point_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMultiRegionAccessPointOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_multi_region_access_point_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMultiRegionAccessPointOperation(ctx, input)
			},
		},
		"dissociate-access-grants-identity-center": {
			Name:   "dissociate-access-grants-identity-center",
			Fields: fields_dissociate_access_grants_identity_center,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DissociateAccessGrantsIdentityCenterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dissociate_access_grants_identity_center, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DissociateAccessGrantsIdentityCenter(ctx, input)
			},
		},
		"get-access-grant": {
			Name:   "get-access-grant",
			Fields: fields_get_access_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessGrant(ctx, input)
			},
		},
		"get-access-grants-instance": {
			Name:   "get-access-grants-instance",
			Fields: fields_get_access_grants_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessGrantsInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_grants_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessGrantsInstance(ctx, input)
			},
		},
		"get-access-grants-instance-for-prefix": {
			Name:   "get-access-grants-instance-for-prefix",
			Fields: fields_get_access_grants_instance_for_prefix,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessGrantsInstanceForPrefixInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_grants_instance_for_prefix, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessGrantsInstanceForPrefix(ctx, input)
			},
		},
		"get-access-grants-instance-resource-policy": {
			Name:   "get-access-grants-instance-resource-policy",
			Fields: fields_get_access_grants_instance_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessGrantsInstanceResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_grants_instance_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessGrantsInstanceResourcePolicy(ctx, input)
			},
		},
		"get-access-grants-location": {
			Name:   "get-access-grants-location",
			Fields: fields_get_access_grants_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessGrantsLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_grants_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessGrantsLocation(ctx, input)
			},
		},
		"get-access-point": {
			Name:   "get-access-point",
			Fields: fields_get_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPoint(ctx, input)
			},
		},
		"get-access-point-configuration-for-object-lambda": {
			Name:   "get-access-point-configuration-for-object-lambda",
			Fields: fields_get_access_point_configuration_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointConfigurationForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_configuration_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointConfigurationForObjectLambda(ctx, input)
			},
		},
		"get-access-point-for-object-lambda": {
			Name:   "get-access-point-for-object-lambda",
			Fields: fields_get_access_point_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointForObjectLambda(ctx, input)
			},
		},
		"get-access-point-policy": {
			Name:   "get-access-point-policy",
			Fields: fields_get_access_point_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointPolicy(ctx, input)
			},
		},
		"get-access-point-policy-for-object-lambda": {
			Name:   "get-access-point-policy-for-object-lambda",
			Fields: fields_get_access_point_policy_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointPolicyForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_policy_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointPolicyForObjectLambda(ctx, input)
			},
		},
		"get-access-point-policy-status": {
			Name:   "get-access-point-policy-status",
			Fields: fields_get_access_point_policy_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointPolicyStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_policy_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointPolicyStatus(ctx, input)
			},
		},
		"get-access-point-policy-status-for-object-lambda": {
			Name:   "get-access-point-policy-status-for-object-lambda",
			Fields: fields_get_access_point_policy_status_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointPolicyStatusForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_policy_status_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointPolicyStatusForObjectLambda(ctx, input)
			},
		},
		"get-access-point-scope": {
			Name:   "get-access-point-scope",
			Fields: fields_get_access_point_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPointScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_point_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPointScope(ctx, input)
			},
		},
		"get-bucket": {
			Name:   "get-bucket",
			Fields: fields_get_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucket(ctx, input)
			},
		},
		"get-bucket-lifecycle-configuration": {
			Name:   "get-bucket-lifecycle-configuration",
			Fields: fields_get_bucket_lifecycle_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketLifecycleConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_lifecycle_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketLifecycleConfiguration(ctx, input)
			},
		},
		"get-bucket-policy": {
			Name:   "get-bucket-policy",
			Fields: fields_get_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketPolicy(ctx, input)
			},
		},
		"get-bucket-replication": {
			Name:   "get-bucket-replication",
			Fields: fields_get_bucket_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketReplication(ctx, input)
			},
		},
		"get-bucket-tagging": {
			Name:   "get-bucket-tagging",
			Fields: fields_get_bucket_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketTagging(ctx, input)
			},
		},
		"get-bucket-versioning": {
			Name:   "get-bucket-versioning",
			Fields: fields_get_bucket_versioning,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketVersioningInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_versioning, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketVersioning(ctx, input)
			},
		},
		"get-data-access": {
			Name:   "get-data-access",
			Fields: fields_get_data_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataAccess(ctx, input)
			},
		},
		"get-job-tagging": {
			Name:   "get-job-tagging",
			Fields: fields_get_job_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobTagging(ctx, input)
			},
		},
		"get-multi-region-access-point": {
			Name:   "get-multi-region-access-point",
			Fields: fields_get_multi_region_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMultiRegionAccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multi_region_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMultiRegionAccessPoint(ctx, input)
			},
		},
		"get-multi-region-access-point-policy": {
			Name:   "get-multi-region-access-point-policy",
			Fields: fields_get_multi_region_access_point_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMultiRegionAccessPointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multi_region_access_point_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMultiRegionAccessPointPolicy(ctx, input)
			},
		},
		"get-multi-region-access-point-policy-status": {
			Name:   "get-multi-region-access-point-policy-status",
			Fields: fields_get_multi_region_access_point_policy_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMultiRegionAccessPointPolicyStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multi_region_access_point_policy_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMultiRegionAccessPointPolicyStatus(ctx, input)
			},
		},
		"get-multi-region-access-point-routes": {
			Name:   "get-multi-region-access-point-routes",
			Fields: fields_get_multi_region_access_point_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMultiRegionAccessPointRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multi_region_access_point_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMultiRegionAccessPointRoutes(ctx, input)
			},
		},
		"get-public-access-block": {
			Name:   "get-public-access-block",
			Fields: fields_get_public_access_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPublicAccessBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_public_access_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPublicAccessBlock(ctx, input)
			},
		},
		"get-storage-lens-configuration": {
			Name:   "get-storage-lens-configuration",
			Fields: fields_get_storage_lens_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStorageLensConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_storage_lens_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStorageLensConfiguration(ctx, input)
			},
		},
		"get-storage-lens-configuration-tagging": {
			Name:   "get-storage-lens-configuration-tagging",
			Fields: fields_get_storage_lens_configuration_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStorageLensConfigurationTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_storage_lens_configuration_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStorageLensConfigurationTagging(ctx, input)
			},
		},
		"get-storage-lens-group": {
			Name:   "get-storage-lens-group",
			Fields: fields_get_storage_lens_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStorageLensGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_storage_lens_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStorageLensGroup(ctx, input)
			},
		},
		"list-access-grants": {
			Name:   "list-access-grants",
			Fields: fields_list_access_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessGrants(ctx, input)
				}
				var results []*svc.ListAccessGrantsOutput
				p := svc.NewListAccessGrantsPaginator(client, input)
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
		"list-access-grants-instances": {
			Name:   "list-access-grants-instances",
			Fields: fields_list_access_grants_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessGrantsInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_grants_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessGrantsInstances(ctx, input)
				}
				var results []*svc.ListAccessGrantsInstancesOutput
				p := svc.NewListAccessGrantsInstancesPaginator(client, input)
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
		"list-access-grants-locations": {
			Name:   "list-access-grants-locations",
			Fields: fields_list_access_grants_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessGrantsLocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_grants_locations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessGrantsLocations(ctx, input)
				}
				var results []*svc.ListAccessGrantsLocationsOutput
				p := svc.NewListAccessGrantsLocationsPaginator(client, input)
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
		"list-access-points": {
			Name:   "list-access-points",
			Fields: fields_list_access_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPoints(ctx, input)
				}
				var results []*svc.ListAccessPointsOutput
				p := svc.NewListAccessPointsPaginator(client, input)
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
		"list-access-points-for-directory-buckets": {
			Name:   "list-access-points-for-directory-buckets",
			Fields: fields_list_access_points_for_directory_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPointsForDirectoryBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_points_for_directory_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPointsForDirectoryBuckets(ctx, input)
				}
				var results []*svc.ListAccessPointsForDirectoryBucketsOutput
				p := svc.NewListAccessPointsForDirectoryBucketsPaginator(client, input)
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
		"list-access-points-for-object-lambda": {
			Name:   "list-access-points-for-object-lambda",
			Fields: fields_list_access_points_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPointsForObjectLambdaInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_points_for_object_lambda, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPointsForObjectLambda(ctx, input)
				}
				var results []*svc.ListAccessPointsForObjectLambdaOutput
				p := svc.NewListAccessPointsForObjectLambdaPaginator(client, input)
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
		"list-caller-access-grants": {
			Name:   "list-caller-access-grants",
			Fields: fields_list_caller_access_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCallerAccessGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_caller_access_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCallerAccessGrants(ctx, input)
				}
				var results []*svc.ListCallerAccessGrantsOutput
				p := svc.NewListCallerAccessGrantsPaginator(client, input)
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
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-multi-region-access-points": {
			Name:   "list-multi-region-access-points",
			Fields: fields_list_multi_region_access_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultiRegionAccessPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multi_region_access_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultiRegionAccessPoints(ctx, input)
				}
				var results []*svc.ListMultiRegionAccessPointsOutput
				p := svc.NewListMultiRegionAccessPointsPaginator(client, input)
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
		"list-regional-buckets": {
			Name:   "list-regional-buckets",
			Fields: fields_list_regional_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegionalBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_regional_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegionalBuckets(ctx, input)
				}
				var results []*svc.ListRegionalBucketsOutput
				p := svc.NewListRegionalBucketsPaginator(client, input)
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
		"list-storage-lens-configurations": {
			Name:   "list-storage-lens-configurations",
			Fields: fields_list_storage_lens_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStorageLensConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_storage_lens_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStorageLensConfigurations(ctx, input)
				}
				var results []*svc.ListStorageLensConfigurationsOutput
				p := svc.NewListStorageLensConfigurationsPaginator(client, input)
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
		"list-storage-lens-groups": {
			Name:   "list-storage-lens-groups",
			Fields: fields_list_storage_lens_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStorageLensGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_storage_lens_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStorageLensGroups(ctx, input)
				}
				var results []*svc.ListStorageLensGroupsOutput
				p := svc.NewListStorageLensGroupsPaginator(client, input)
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
		"put-access-grants-instance-resource-policy": {
			Name:   "put-access-grants-instance-resource-policy",
			Fields: fields_put_access_grants_instance_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccessGrantsInstanceResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_access_grants_instance_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccessGrantsInstanceResourcePolicy(ctx, input)
			},
		},
		"put-access-point-configuration-for-object-lambda": {
			Name:   "put-access-point-configuration-for-object-lambda",
			Fields: fields_put_access_point_configuration_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccessPointConfigurationForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_access_point_configuration_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccessPointConfigurationForObjectLambda(ctx, input)
			},
		},
		"put-access-point-policy": {
			Name:   "put-access-point-policy",
			Fields: fields_put_access_point_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccessPointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_access_point_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccessPointPolicy(ctx, input)
			},
		},
		"put-access-point-policy-for-object-lambda": {
			Name:   "put-access-point-policy-for-object-lambda",
			Fields: fields_put_access_point_policy_for_object_lambda,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccessPointPolicyForObjectLambdaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_access_point_policy_for_object_lambda, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccessPointPolicyForObjectLambda(ctx, input)
			},
		},
		"put-access-point-scope": {
			Name:   "put-access-point-scope",
			Fields: fields_put_access_point_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccessPointScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_access_point_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccessPointScope(ctx, input)
			},
		},
		"put-bucket-lifecycle-configuration": {
			Name:   "put-bucket-lifecycle-configuration",
			Fields: fields_put_bucket_lifecycle_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketLifecycleConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_lifecycle_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketLifecycleConfiguration(ctx, input)
			},
		},
		"put-bucket-policy": {
			Name:   "put-bucket-policy",
			Fields: fields_put_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketPolicy(ctx, input)
			},
		},
		"put-bucket-replication": {
			Name:   "put-bucket-replication",
			Fields: fields_put_bucket_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketReplication(ctx, input)
			},
		},
		"put-bucket-tagging": {
			Name:   "put-bucket-tagging",
			Fields: fields_put_bucket_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketTagging(ctx, input)
			},
		},
		"put-bucket-versioning": {
			Name:   "put-bucket-versioning",
			Fields: fields_put_bucket_versioning,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketVersioningInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_versioning, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketVersioning(ctx, input)
			},
		},
		"put-job-tagging": {
			Name:   "put-job-tagging",
			Fields: fields_put_job_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutJobTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_job_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutJobTagging(ctx, input)
			},
		},
		"put-multi-region-access-point-policy": {
			Name:   "put-multi-region-access-point-policy",
			Fields: fields_put_multi_region_access_point_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMultiRegionAccessPointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_multi_region_access_point_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMultiRegionAccessPointPolicy(ctx, input)
			},
		},
		"put-public-access-block": {
			Name:   "put-public-access-block",
			Fields: fields_put_public_access_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPublicAccessBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_public_access_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPublicAccessBlock(ctx, input)
			},
		},
		"put-storage-lens-configuration": {
			Name:   "put-storage-lens-configuration",
			Fields: fields_put_storage_lens_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutStorageLensConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_storage_lens_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutStorageLensConfiguration(ctx, input)
			},
		},
		"put-storage-lens-configuration-tagging": {
			Name:   "put-storage-lens-configuration-tagging",
			Fields: fields_put_storage_lens_configuration_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutStorageLensConfigurationTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_storage_lens_configuration_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutStorageLensConfigurationTagging(ctx, input)
			},
		},
		"submit-multi-region-access-point-routes": {
			Name:   "submit-multi-region-access-point-routes",
			Fields: fields_submit_multi_region_access_point_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitMultiRegionAccessPointRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_multi_region_access_point_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitMultiRegionAccessPointRoutes(ctx, input)
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
		"update-access-grants-location": {
			Name:   "update-access-grants-location",
			Fields: fields_update_access_grants_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessGrantsLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_grants_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessGrantsLocation(ctx, input)
			},
		},
		"update-job-priority": {
			Name:   "update-job-priority",
			Fields: fields_update_job_priority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobPriorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_priority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobPriority(ctx, input)
			},
		},
		"update-job-status": {
			Name:   "update-job-status",
			Fields: fields_update_job_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobStatus(ctx, input)
			},
		},
		"update-storage-lens-group": {
			Name:   "update-storage-lens-group",
			Fields: fields_update_storage_lens_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStorageLensGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_storage_lens_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStorageLensGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("s3control", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
