package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/s3"
)

var fields_abort_multipart_upload = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "IfMatchInitiatedTime", Flag: "if-match-initiated-time", Type: "*time.Time", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_complete_multipart_upload = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumCRC32", Flag: "checksum-crc32", Type: "*string", Required: false},
	{Name: "ChecksumCRC32C", Flag: "checksum-crc32-c", Type: "*string", Required: false},
	{Name: "ChecksumCRC64NVME", Flag: "checksum-crc64-nvme", Type: "*string", Required: false},
	{Name: "ChecksumSHA1", Flag: "checksum-sha1", Type: "*string", Required: false},
	{Name: "ChecksumSHA256", Flag: "checksum-sha256", Type: "*string", Required: false},
	{Name: "ChecksumType", Flag: "checksum-type", Type: "types.ChecksumType", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "MpuObjectSize", Flag: "mpu-object-size", Type: "*int64", Required: false},
	{Name: "MultipartUpload", Flag: "multipart-upload", Type: "*types.CompletedMultipartUpload", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_copy_object = []leanruntime.Field{
	{Name: "ACL", Flag: "acl", Type: "types.ObjectCannedACL", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BucketKeyEnabled", Flag: "bucket-key-enabled", Type: "*bool", Required: false},
	{Name: "CacheControl", Flag: "cache-control", Type: "*string", Required: false},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentDisposition", Flag: "content-disposition", Type: "*string", Required: false},
	{Name: "ContentEncoding", Flag: "content-encoding", Type: "*string", Required: false},
	{Name: "ContentLanguage", Flag: "content-language", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "CopySource", Flag: "copy-source", Type: "*string", Required: true},
	{Name: "CopySourceIfMatch", Flag: "copy-source-if-match", Type: "*string", Required: false},
	{Name: "CopySourceIfModifiedSince", Flag: "copy-source-if-modified-since", Type: "*time.Time", Required: false},
	{Name: "CopySourceIfNoneMatch", Flag: "copy-source-if-none-match", Type: "*string", Required: false},
	{Name: "CopySourceIfUnmodifiedSince", Flag: "copy-source-if-unmodified-since", Type: "*time.Time", Required: false},
	{Name: "CopySourceSSECustomerAlgorithm", Flag: "copy-source-sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "CopySourceSSECustomerKey", Flag: "copy-source-sse-customer-key", Type: "*string", Required: false},
	{Name: "CopySourceSSECustomerKeyMD5", Flag: "copy-source-sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "ExpectedSourceBucketOwner", Flag: "expected-source-bucket-owner", Type: "*string", Required: false},
	{Name: "Expires", Flag: "expires", Type: "*time.Time", Required: false},
	{Name: "GrantFullControl", Flag: "grant-full-control", Type: "*string", Required: false},
	{Name: "GrantRead", Flag: "grant-read", Type: "*string", Required: false},
	{Name: "GrantReadACP", Flag: "grant-read-acp", Type: "*string", Required: false},
	{Name: "GrantWriteACP", Flag: "grant-write-acp", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "MetadataDirective", Flag: "metadata-directive", Type: "types.MetadataDirective", Required: false},
	{Name: "ObjectLockLegalHoldStatus", Flag: "object-lock-legal-hold-status", Type: "types.ObjectLockLegalHoldStatus", Required: false},
	{Name: "ObjectLockMode", Flag: "object-lock-mode", Type: "types.ObjectLockMode", Required: false},
	{Name: "ObjectLockRetainUntilDate", Flag: "object-lock-retain-until-date", Type: "*time.Time", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "SSEKMSEncryptionContext", Flag: "ssekms-encryption-context", Type: "*string", Required: false},
	{Name: "SSEKMSKeyId", Flag: "ssekms-key-id", Type: "*string", Required: false},
	{Name: "ServerSideEncryption", Flag: "server-side-encryption", Type: "types.ServerSideEncryption", Required: false},
	{Name: "StorageClass", Flag: "storage-class", Type: "types.StorageClass", Required: false},
	{Name: "Tagging", Flag: "tagging", Type: "*string", Required: false},
	{Name: "TaggingDirective", Flag: "tagging-directive", Type: "types.TaggingDirective", Required: false},
	{Name: "WebsiteRedirectLocation", Flag: "website-redirect-location", Type: "*string", Required: false},
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
	{Name: "ObjectLockEnabledForBucket", Flag: "object-lock-enabled-for-bucket", Type: "*bool", Required: false},
	{Name: "ObjectOwnership", Flag: "object-ownership", Type: "types.ObjectOwnership", Required: false},
}

var fields_create_bucket_metadata_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "MetadataConfiguration", Flag: "metadata-configuration", Type: "*types.MetadataConfiguration", Required: true},
}

var fields_create_bucket_metadata_table_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "MetadataTableConfiguration", Flag: "metadata-table-configuration", Type: "*types.MetadataTableConfiguration", Required: true},
}

var fields_create_multipart_upload = []leanruntime.Field{
	{Name: "ACL", Flag: "acl", Type: "types.ObjectCannedACL", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BucketKeyEnabled", Flag: "bucket-key-enabled", Type: "*bool", Required: false},
	{Name: "CacheControl", Flag: "cache-control", Type: "*string", Required: false},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ChecksumType", Flag: "checksum-type", Type: "types.ChecksumType", Required: false},
	{Name: "ContentDisposition", Flag: "content-disposition", Type: "*string", Required: false},
	{Name: "ContentEncoding", Flag: "content-encoding", Type: "*string", Required: false},
	{Name: "ContentLanguage", Flag: "content-language", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Expires", Flag: "expires", Type: "*time.Time", Required: false},
	{Name: "GrantFullControl", Flag: "grant-full-control", Type: "*string", Required: false},
	{Name: "GrantRead", Flag: "grant-read", Type: "*string", Required: false},
	{Name: "GrantReadACP", Flag: "grant-read-acp", Type: "*string", Required: false},
	{Name: "GrantWriteACP", Flag: "grant-write-acp", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "ObjectLockLegalHoldStatus", Flag: "object-lock-legal-hold-status", Type: "types.ObjectLockLegalHoldStatus", Required: false},
	{Name: "ObjectLockMode", Flag: "object-lock-mode", Type: "types.ObjectLockMode", Required: false},
	{Name: "ObjectLockRetainUntilDate", Flag: "object-lock-retain-until-date", Type: "*time.Time", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "SSEKMSEncryptionContext", Flag: "ssekms-encryption-context", Type: "*string", Required: false},
	{Name: "SSEKMSKeyId", Flag: "ssekms-key-id", Type: "*string", Required: false},
	{Name: "ServerSideEncryption", Flag: "server-side-encryption", Type: "types.ServerSideEncryption", Required: false},
	{Name: "StorageClass", Flag: "storage-class", Type: "types.StorageClass", Required: false},
	{Name: "Tagging", Flag: "tagging", Type: "*string", Required: false},
	{Name: "WebsiteRedirectLocation", Flag: "website-redirect-location", Type: "*string", Required: false},
}

var fields_create_session = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BucketKeyEnabled", Flag: "bucket-key-enabled", Type: "*bool", Required: false},
	{Name: "SSEKMSEncryptionContext", Flag: "ssekms-encryption-context", Type: "*string", Required: false},
	{Name: "SSEKMSKeyId", Flag: "ssekms-key-id", Type: "*string", Required: false},
	{Name: "ServerSideEncryption", Flag: "server-side-encryption", Type: "types.ServerSideEncryption", Required: false},
	{Name: "SessionMode", Flag: "session-mode", Type: "types.SessionMode", Required: false},
}

var fields_delete_bucket = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_analytics_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_bucket_cors = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_encryption = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_intelligent_tiering_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_bucket_inventory_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_bucket_lifecycle = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_metadata_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_metadata_table_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_metrics_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_bucket_ownership_controls = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_policy = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_replication = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_tagging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_bucket_website = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_delete_object = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BypassGovernanceRetention", Flag: "bypass-governance-retention", Type: "*bool", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfMatchLastModifiedTime", Flag: "if-match-last-modified-time", Type: "*time.Time", Required: false},
	{Name: "IfMatchSize", Flag: "if-match-size", Type: "*int64", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "MFA", Flag: "mfa", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_delete_object_tagging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_delete_objects = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BypassGovernanceRetention", Flag: "bypass-governance-retention", Type: "*bool", Required: false},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "Delete", Flag: "delete", Type: "*types.Delete", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "MFA", Flag: "mfa", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
}

var fields_delete_public_access_block = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_abac = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_accelerate_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
}

var fields_get_bucket_acl = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_analytics_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_bucket_cors = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_encryption = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_intelligent_tiering_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_bucket_inventory_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_bucket_lifecycle_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_location = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_logging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_metadata_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_metadata_table_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_metrics_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_bucket_notification_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_ownership_controls = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_policy = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_policy_status = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_replication = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_request_payment = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_tagging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_versioning = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_bucket_website = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_object = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumMode", Flag: "checksum-mode", Type: "types.ChecksumMode", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfModifiedSince", Flag: "if-modified-since", Type: "*time.Time", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "IfUnmodifiedSince", Flag: "if-unmodified-since", Type: "*time.Time", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: false},
	{Name: "Range", Flag: "range", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "ResponseCacheControl", Flag: "response-cache-control", Type: "*string", Required: false},
	{Name: "ResponseContentDisposition", Flag: "response-content-disposition", Type: "*string", Required: false},
	{Name: "ResponseContentEncoding", Flag: "response-content-encoding", Type: "*string", Required: false},
	{Name: "ResponseContentLanguage", Flag: "response-content-language", Type: "*string", Required: false},
	{Name: "ResponseContentType", Flag: "response-content-type", Type: "*string", Required: false},
	{Name: "ResponseExpires", Flag: "response-expires", Type: "*time.Time", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_object_acl = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_object_attributes = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "MaxParts", Flag: "max-parts", Type: "*int32", Required: false},
	{Name: "ObjectAttributes", Flag: "object-attributes", Type: "[]types.ObjectAttributes", Required: true},
	{Name: "PartNumberMarker", Flag: "part-number-marker", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_object_legal_hold = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_object_lock_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_get_object_retention = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_object_tagging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_get_object_torrent = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
}

var fields_get_public_access_block = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_head_bucket = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_head_object = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumMode", Flag: "checksum-mode", Type: "types.ChecksumMode", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfModifiedSince", Flag: "if-modified-since", Type: "*time.Time", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "IfUnmodifiedSince", Flag: "if-unmodified-since", Type: "*time.Time", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: false},
	{Name: "Range", Flag: "range", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "ResponseCacheControl", Flag: "response-cache-control", Type: "*string", Required: false},
	{Name: "ResponseContentDisposition", Flag: "response-content-disposition", Type: "*string", Required: false},
	{Name: "ResponseContentEncoding", Flag: "response-content-encoding", Type: "*string", Required: false},
	{Name: "ResponseContentLanguage", Flag: "response-content-language", Type: "*string", Required: false},
	{Name: "ResponseContentType", Flag: "response-content-type", Type: "*string", Required: false},
	{Name: "ResponseExpires", Flag: "response-expires", Type: "*time.Time", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_list_bucket_analytics_configurations = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_list_bucket_intelligent_tiering_configurations = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_list_bucket_inventory_configurations = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_list_bucket_metrics_configurations = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_list_buckets = []leanruntime.Field{
	{Name: "BucketRegion", Flag: "bucket-region", Type: "*string", Required: false},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "MaxBuckets", Flag: "max-buckets", Type: "*int32", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
}

var fields_list_directory_buckets = []leanruntime.Field{
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "MaxDirectoryBuckets", Flag: "max-directory-buckets", Type: "*int32", Required: false},
}

var fields_list_multipart_uploads = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "Delimiter", Flag: "delimiter", Type: "*string", Required: false},
	{Name: "EncodingType", Flag: "encoding-type", Type: "types.EncodingType", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "KeyMarker", Flag: "key-marker", Type: "*string", Required: false},
	{Name: "MaxUploads", Flag: "max-uploads", Type: "*int32", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "UploadIdMarker", Flag: "upload-id-marker", Type: "*string", Required: false},
}

var fields_list_object_versions = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "Delimiter", Flag: "delimiter", Type: "*string", Required: false},
	{Name: "EncodingType", Flag: "encoding-type", Type: "types.EncodingType", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "KeyMarker", Flag: "key-marker", Type: "*string", Required: false},
	{Name: "MaxKeys", Flag: "max-keys", Type: "*int32", Required: false},
	{Name: "OptionalObjectAttributes", Flag: "optional-object-attributes", Type: "[]types.OptionalObjectAttributes", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionIdMarker", Flag: "version-id-marker", Type: "*string", Required: false},
}

var fields_list_objects = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "Delimiter", Flag: "delimiter", Type: "*string", Required: false},
	{Name: "EncodingType", Flag: "encoding-type", Type: "types.EncodingType", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxKeys", Flag: "max-keys", Type: "*int32", Required: false},
	{Name: "OptionalObjectAttributes", Flag: "optional-object-attributes", Type: "[]types.OptionalObjectAttributes", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
}

var fields_list_objects_v2 = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "Delimiter", Flag: "delimiter", Type: "*string", Required: false},
	{Name: "EncodingType", Flag: "encoding-type", Type: "types.EncodingType", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "FetchOwner", Flag: "fetch-owner", Type: "*bool", Required: false},
	{Name: "MaxKeys", Flag: "max-keys", Type: "*int32", Required: false},
	{Name: "OptionalObjectAttributes", Flag: "optional-object-attributes", Type: "[]types.OptionalObjectAttributes", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "StartAfter", Flag: "start-after", Type: "*string", Required: false},
}

var fields_list_parts = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "MaxParts", Flag: "max-parts", Type: "*int32", Required: false},
	{Name: "PartNumberMarker", Flag: "part-number-marker", Type: "*string", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_put_bucket_abac = []leanruntime.Field{
	{Name: "AbacStatus", Flag: "abac-status", Type: "*types.AbacStatus", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_put_bucket_accelerate_configuration = []leanruntime.Field{
	{Name: "AccelerateConfiguration", Flag: "accelerate-configuration", Type: "*types.AccelerateConfiguration", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_put_bucket_acl = []leanruntime.Field{
	{Name: "ACL", Flag: "acl", Type: "types.BucketCannedACL", Required: false},
	{Name: "AccessControlPolicy", Flag: "access-control-policy", Type: "*types.AccessControlPolicy", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "GrantFullControl", Flag: "grant-full-control", Type: "*string", Required: false},
	{Name: "GrantRead", Flag: "grant-read", Type: "*string", Required: false},
	{Name: "GrantReadACP", Flag: "grant-read-acp", Type: "*string", Required: false},
	{Name: "GrantWrite", Flag: "grant-write", Type: "*string", Required: false},
	{Name: "GrantWriteACP", Flag: "grant-write-acp", Type: "*string", Required: false},
}

var fields_put_bucket_analytics_configuration = []leanruntime.Field{
	{Name: "AnalyticsConfiguration", Flag: "analytics-configuration", Type: "*types.AnalyticsConfiguration", Required: true},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_put_bucket_cors = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "CORSConfiguration", Flag: "cors-configuration", Type: "*types.CORSConfiguration", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_put_bucket_encryption = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: true},
}

var fields_put_bucket_intelligent_tiering_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IntelligentTieringConfiguration", Flag: "intelligent-tiering-configuration", Type: "*types.IntelligentTieringConfiguration", Required: true},
}

var fields_put_bucket_inventory_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "InventoryConfiguration", Flag: "inventory-configuration", Type: "*types.InventoryConfiguration", Required: true},
}

var fields_put_bucket_lifecycle_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "LifecycleConfiguration", Flag: "lifecycle-configuration", Type: "*types.BucketLifecycleConfiguration", Required: false},
	{Name: "TransitionDefaultMinimumObjectSize", Flag: "transition-default-minimum-object-size", Type: "types.TransitionDefaultMinimumObjectSize", Required: false},
}

var fields_put_bucket_logging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BucketLoggingStatus", Flag: "bucket-logging-status", Type: "*types.BucketLoggingStatus", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
}

var fields_put_bucket_metrics_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MetricsConfiguration", Flag: "metrics-configuration", Type: "*types.MetricsConfiguration", Required: true},
}

var fields_put_bucket_notification_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "NotificationConfiguration", Flag: "notification-configuration", Type: "*types.NotificationConfiguration", Required: true},
	{Name: "SkipDestinationValidation", Flag: "skip-destination-validation", Type: "*bool", Required: false},
}

var fields_put_bucket_ownership_controls = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "OwnershipControls", Flag: "ownership-controls", Type: "*types.OwnershipControls", Required: true},
}

var fields_put_bucket_policy = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ConfirmRemoveSelfBucketAccess", Flag: "confirm-remove-self-bucket-access", Type: "*bool", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_bucket_replication = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "ReplicationConfiguration", Flag: "replication-configuration", Type: "*types.ReplicationConfiguration", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_put_bucket_request_payment = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "RequestPaymentConfiguration", Flag: "request-payment-configuration", Type: "*types.RequestPaymentConfiguration", Required: true},
}

var fields_put_bucket_tagging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Tagging", Flag: "tagging", Type: "*types.Tagging", Required: true},
}

var fields_put_bucket_versioning = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "MFA", Flag: "mfa", Type: "*string", Required: false},
	{Name: "VersioningConfiguration", Flag: "versioning-configuration", Type: "*types.VersioningConfiguration", Required: true},
}

var fields_put_bucket_website = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "WebsiteConfiguration", Flag: "website-configuration", Type: "*types.WebsiteConfiguration", Required: true},
}

var fields_put_object = []leanruntime.Field{
	{Name: "ACL", Flag: "acl", Type: "types.ObjectCannedACL", Required: false},
	{Name: "Body", Flag: "body", Type: "io.Reader", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BucketKeyEnabled", Flag: "bucket-key-enabled", Type: "*bool", Required: false},
	{Name: "CacheControl", Flag: "cache-control", Type: "*string", Required: false},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ChecksumCRC32", Flag: "checksum-crc32", Type: "*string", Required: false},
	{Name: "ChecksumCRC32C", Flag: "checksum-crc32-c", Type: "*string", Required: false},
	{Name: "ChecksumCRC64NVME", Flag: "checksum-crc64-nvme", Type: "*string", Required: false},
	{Name: "ChecksumSHA1", Flag: "checksum-sha1", Type: "*string", Required: false},
	{Name: "ChecksumSHA256", Flag: "checksum-sha256", Type: "*string", Required: false},
	{Name: "ContentDisposition", Flag: "content-disposition", Type: "*string", Required: false},
	{Name: "ContentEncoding", Flag: "content-encoding", Type: "*string", Required: false},
	{Name: "ContentLanguage", Flag: "content-language", Type: "*string", Required: false},
	{Name: "ContentLength", Flag: "content-length", Type: "*int64", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Expires", Flag: "expires", Type: "*time.Time", Required: false},
	{Name: "GrantFullControl", Flag: "grant-full-control", Type: "*string", Required: false},
	{Name: "GrantRead", Flag: "grant-read", Type: "*string", Required: false},
	{Name: "GrantReadACP", Flag: "grant-read-acp", Type: "*string", Required: false},
	{Name: "GrantWriteACP", Flag: "grant-write-acp", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "ObjectLockLegalHoldStatus", Flag: "object-lock-legal-hold-status", Type: "types.ObjectLockLegalHoldStatus", Required: false},
	{Name: "ObjectLockMode", Flag: "object-lock-mode", Type: "types.ObjectLockMode", Required: false},
	{Name: "ObjectLockRetainUntilDate", Flag: "object-lock-retain-until-date", Type: "*time.Time", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "SSEKMSEncryptionContext", Flag: "ssekms-encryption-context", Type: "*string", Required: false},
	{Name: "SSEKMSKeyId", Flag: "ssekms-key-id", Type: "*string", Required: false},
	{Name: "ServerSideEncryption", Flag: "server-side-encryption", Type: "types.ServerSideEncryption", Required: false},
	{Name: "StorageClass", Flag: "storage-class", Type: "types.StorageClass", Required: false},
	{Name: "Tagging", Flag: "tagging", Type: "*string", Required: false},
	{Name: "WebsiteRedirectLocation", Flag: "website-redirect-location", Type: "*string", Required: false},
	{Name: "WriteOffsetBytes", Flag: "write-offset-bytes", Type: "*int64", Required: false},
}

var fields_put_object_acl = []leanruntime.Field{
	{Name: "ACL", Flag: "acl", Type: "types.ObjectCannedACL", Required: false},
	{Name: "AccessControlPolicy", Flag: "access-control-policy", Type: "*types.AccessControlPolicy", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "GrantFullControl", Flag: "grant-full-control", Type: "*string", Required: false},
	{Name: "GrantRead", Flag: "grant-read", Type: "*string", Required: false},
	{Name: "GrantReadACP", Flag: "grant-read-acp", Type: "*string", Required: false},
	{Name: "GrantWrite", Flag: "grant-write", Type: "*string", Required: false},
	{Name: "GrantWriteACP", Flag: "grant-write-acp", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_put_object_legal_hold = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "LegalHold", Flag: "legal-hold", Type: "*types.ObjectLockLegalHold", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_put_object_lock_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "ObjectLockConfiguration", Flag: "object-lock-configuration", Type: "*types.ObjectLockConfiguration", Required: false},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_put_object_retention = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "BypassGovernanceRetention", Flag: "bypass-governance-retention", Type: "*bool", Required: false},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "Retention", Flag: "retention", Type: "*types.ObjectLockRetention", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_put_object_tagging = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "Tagging", Flag: "tagging", Type: "*types.Tagging", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_put_public_access_block = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "PublicAccessBlockConfiguration", Flag: "public-access-block-configuration", Type: "*types.PublicAccessBlockConfiguration", Required: true},
}

var fields_rename_object = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationIfMatch", Flag: "destination-if-match", Type: "*string", Required: false},
	{Name: "DestinationIfModifiedSince", Flag: "destination-if-modified-since", Type: "*time.Time", Required: false},
	{Name: "DestinationIfNoneMatch", Flag: "destination-if-none-match", Type: "*string", Required: false},
	{Name: "DestinationIfUnmodifiedSince", Flag: "destination-if-unmodified-since", Type: "*time.Time", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RenameSource", Flag: "rename-source", Type: "*string", Required: true},
	{Name: "SourceIfMatch", Flag: "source-if-match", Type: "*string", Required: false},
	{Name: "SourceIfModifiedSince", Flag: "source-if-modified-since", Type: "*time.Time", Required: false},
	{Name: "SourceIfNoneMatch", Flag: "source-if-none-match", Type: "*string", Required: false},
	{Name: "SourceIfUnmodifiedSince", Flag: "source-if-unmodified-since", Type: "*time.Time", Required: false},
}

var fields_restore_object = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "RestoreRequest", Flag: "restore-request", Type: "*types.RestoreRequest", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_select_object_content = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "ExpressionType", Flag: "expression-type", Type: "types.ExpressionType", Required: true},
	{Name: "InputSerialization", Flag: "input-serialization", Type: "*types.InputSerialization", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "OutputSerialization", Flag: "output-serialization", Type: "*types.OutputSerialization", Required: true},
	{Name: "RequestProgress", Flag: "request-progress", Type: "*types.RequestProgress", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "ScanRange", Flag: "scan-range", Type: "*types.ScanRange", Required: false},
}

var fields_update_bucket_metadata_inventory_table_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "InventoryTableConfiguration", Flag: "inventory-table-configuration", Type: "*types.InventoryTableConfigurationUpdates", Required: true},
}

var fields_update_bucket_metadata_journal_table_configuration = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "JournalTableConfiguration", Flag: "journal-table-configuration", Type: "*types.JournalTableConfigurationUpdates", Required: true},
}

var fields_update_object_encryption = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "ObjectEncryption", Flag: "object-encryption", Type: "types.ObjectEncryption", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_upload_part = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "io.Reader", Required: false},
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "ChecksumAlgorithm", Flag: "checksum-algorithm", Type: "types.ChecksumAlgorithm", Required: false},
	{Name: "ChecksumCRC32", Flag: "checksum-crc32", Type: "*string", Required: false},
	{Name: "ChecksumCRC32C", Flag: "checksum-crc32-c", Type: "*string", Required: false},
	{Name: "ChecksumCRC64NVME", Flag: "checksum-crc64-nvme", Type: "*string", Required: false},
	{Name: "ChecksumSHA1", Flag: "checksum-sha1", Type: "*string", Required: false},
	{Name: "ChecksumSHA256", Flag: "checksum-sha256", Type: "*string", Required: false},
	{Name: "ContentLength", Flag: "content-length", Type: "*int64", Required: false},
	{Name: "ContentMD5", Flag: "content-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_upload_part_copy = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "CopySource", Flag: "copy-source", Type: "*string", Required: true},
	{Name: "CopySourceIfMatch", Flag: "copy-source-if-match", Type: "*string", Required: false},
	{Name: "CopySourceIfModifiedSince", Flag: "copy-source-if-modified-since", Type: "*time.Time", Required: false},
	{Name: "CopySourceIfNoneMatch", Flag: "copy-source-if-none-match", Type: "*string", Required: false},
	{Name: "CopySourceIfUnmodifiedSince", Flag: "copy-source-if-unmodified-since", Type: "*time.Time", Required: false},
	{Name: "CopySourceRange", Flag: "copy-source-range", Type: "*string", Required: false},
	{Name: "CopySourceSSECustomerAlgorithm", Flag: "copy-source-sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "CopySourceSSECustomerKey", Flag: "copy-source-sse-customer-key", Type: "*string", Required: false},
	{Name: "CopySourceSSECustomerKeyMD5", Flag: "copy-source-sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "ExpectedBucketOwner", Flag: "expected-bucket-owner", Type: "*string", Required: false},
	{Name: "ExpectedSourceBucketOwner", Flag: "expected-source-bucket-owner", Type: "*string", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "PartNumber", Flag: "part-number", Type: "*int32", Required: true},
	{Name: "RequestPayer", Flag: "request-payer", Type: "types.RequestPayer", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKey", Flag: "sse-customer-key", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_write_get_object_response = []leanruntime.Field{
	{Name: "AcceptRanges", Flag: "accept-ranges", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "io.Reader", Required: false},
	{Name: "BucketKeyEnabled", Flag: "bucket-key-enabled", Type: "*bool", Required: false},
	{Name: "CacheControl", Flag: "cache-control", Type: "*string", Required: false},
	{Name: "ChecksumCRC32", Flag: "checksum-crc32", Type: "*string", Required: false},
	{Name: "ChecksumCRC32C", Flag: "checksum-crc32-c", Type: "*string", Required: false},
	{Name: "ChecksumCRC64NVME", Flag: "checksum-crc64-nvme", Type: "*string", Required: false},
	{Name: "ChecksumSHA1", Flag: "checksum-sha1", Type: "*string", Required: false},
	{Name: "ChecksumSHA256", Flag: "checksum-sha256", Type: "*string", Required: false},
	{Name: "ContentDisposition", Flag: "content-disposition", Type: "*string", Required: false},
	{Name: "ContentEncoding", Flag: "content-encoding", Type: "*string", Required: false},
	{Name: "ContentLanguage", Flag: "content-language", Type: "*string", Required: false},
	{Name: "ContentLength", Flag: "content-length", Type: "*int64", Required: false},
	{Name: "ContentRange", Flag: "content-range", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "DeleteMarker", Flag: "delete-marker", Type: "*bool", Required: false},
	{Name: "ETag", Flag: "etag", Type: "*string", Required: false},
	{Name: "ErrorCode", Flag: "error-code", Type: "*string", Required: false},
	{Name: "ErrorMessage", Flag: "error-message", Type: "*string", Required: false},
	{Name: "Expiration", Flag: "expiration", Type: "*string", Required: false},
	{Name: "Expires", Flag: "expires", Type: "*time.Time", Required: false},
	{Name: "LastModified", Flag: "last-modified", Type: "*time.Time", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "MissingMeta", Flag: "missing-meta", Type: "*int32", Required: false},
	{Name: "ObjectLockLegalHoldStatus", Flag: "object-lock-legal-hold-status", Type: "types.ObjectLockLegalHoldStatus", Required: false},
	{Name: "ObjectLockMode", Flag: "object-lock-mode", Type: "types.ObjectLockMode", Required: false},
	{Name: "ObjectLockRetainUntilDate", Flag: "object-lock-retain-until-date", Type: "*time.Time", Required: false},
	{Name: "PartsCount", Flag: "parts-count", Type: "*int32", Required: false},
	{Name: "ReplicationStatus", Flag: "replication-status", Type: "types.ReplicationStatus", Required: false},
	{Name: "RequestCharged", Flag: "request-charged", Type: "types.RequestCharged", Required: false},
	{Name: "RequestRoute", Flag: "request-route", Type: "*string", Required: true},
	{Name: "RequestToken", Flag: "request-token", Type: "*string", Required: true},
	{Name: "Restore", Flag: "restore", Type: "*string", Required: false},
	{Name: "SSECustomerAlgorithm", Flag: "sse-customer-algorithm", Type: "*string", Required: false},
	{Name: "SSECustomerKeyMD5", Flag: "sse-customer-key-md5", Type: "*string", Required: false},
	{Name: "SSEKMSKeyId", Flag: "ssekms-key-id", Type: "*string", Required: false},
	{Name: "ServerSideEncryption", Flag: "server-side-encryption", Type: "types.ServerSideEncryption", Required: false},
	{Name: "StatusCode", Flag: "status-code", Type: "*int32", Required: false},
	{Name: "StorageClass", Flag: "storage-class", Type: "types.StorageClass", Required: false},
	{Name: "TagCount", Flag: "tag-count", Type: "*int32", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"abort-multipart-upload": {
			Name:   "abort-multipart-upload",
			Fields: fields_abort_multipart_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AbortMultipartUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_abort_multipart_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AbortMultipartUpload(ctx, input)
			},
		},
		"complete-multipart-upload": {
			Name:   "complete-multipart-upload",
			Fields: fields_complete_multipart_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteMultipartUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_multipart_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteMultipartUpload(ctx, input)
			},
		},
		"copy-object": {
			Name:   "copy-object",
			Fields: fields_copy_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyObject(ctx, input)
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
		"create-bucket-metadata-configuration": {
			Name:   "create-bucket-metadata-configuration",
			Fields: fields_create_bucket_metadata_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBucketMetadataConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bucket_metadata_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBucketMetadataConfiguration(ctx, input)
			},
		},
		"create-bucket-metadata-table-configuration": {
			Name:   "create-bucket-metadata-table-configuration",
			Fields: fields_create_bucket_metadata_table_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBucketMetadataTableConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bucket_metadata_table_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBucketMetadataTableConfiguration(ctx, input)
			},
		},
		"create-multipart-upload": {
			Name:   "create-multipart-upload",
			Fields: fields_create_multipart_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultipartUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multipart_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultipartUpload(ctx, input)
			},
		},
		"create-session": {
			Name:   "create-session",
			Fields: fields_create_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSession(ctx, input)
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
		"delete-bucket-analytics-configuration": {
			Name:   "delete-bucket-analytics-configuration",
			Fields: fields_delete_bucket_analytics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketAnalyticsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_analytics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketAnalyticsConfiguration(ctx, input)
			},
		},
		"delete-bucket-cors": {
			Name:   "delete-bucket-cors",
			Fields: fields_delete_bucket_cors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketCorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_cors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketCors(ctx, input)
			},
		},
		"delete-bucket-encryption": {
			Name:   "delete-bucket-encryption",
			Fields: fields_delete_bucket_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketEncryption(ctx, input)
			},
		},
		"delete-bucket-intelligent-tiering-configuration": {
			Name:   "delete-bucket-intelligent-tiering-configuration",
			Fields: fields_delete_bucket_intelligent_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketIntelligentTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_intelligent_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketIntelligentTieringConfiguration(ctx, input)
			},
		},
		"delete-bucket-inventory-configuration": {
			Name:   "delete-bucket-inventory-configuration",
			Fields: fields_delete_bucket_inventory_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketInventoryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_inventory_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketInventoryConfiguration(ctx, input)
			},
		},
		"delete-bucket-lifecycle": {
			Name:   "delete-bucket-lifecycle",
			Fields: fields_delete_bucket_lifecycle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketLifecycleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_lifecycle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketLifecycle(ctx, input)
			},
		},
		"delete-bucket-metadata-configuration": {
			Name:   "delete-bucket-metadata-configuration",
			Fields: fields_delete_bucket_metadata_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketMetadataConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_metadata_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketMetadataConfiguration(ctx, input)
			},
		},
		"delete-bucket-metadata-table-configuration": {
			Name:   "delete-bucket-metadata-table-configuration",
			Fields: fields_delete_bucket_metadata_table_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketMetadataTableConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_metadata_table_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketMetadataTableConfiguration(ctx, input)
			},
		},
		"delete-bucket-metrics-configuration": {
			Name:   "delete-bucket-metrics-configuration",
			Fields: fields_delete_bucket_metrics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketMetricsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_metrics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketMetricsConfiguration(ctx, input)
			},
		},
		"delete-bucket-ownership-controls": {
			Name:   "delete-bucket-ownership-controls",
			Fields: fields_delete_bucket_ownership_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketOwnershipControlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_ownership_controls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketOwnershipControls(ctx, input)
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
		"delete-bucket-website": {
			Name:   "delete-bucket-website",
			Fields: fields_delete_bucket_website,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketWebsiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_website, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketWebsite(ctx, input)
			},
		},
		"delete-object": {
			Name:   "delete-object",
			Fields: fields_delete_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObject(ctx, input)
			},
		},
		"delete-object-tagging": {
			Name:   "delete-object-tagging",
			Fields: fields_delete_object_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObjectTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_object_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObjectTagging(ctx, input)
			},
		},
		"delete-objects": {
			Name:   "delete-objects",
			Fields: fields_delete_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObjectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_objects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObjects(ctx, input)
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
		"get-bucket-abac": {
			Name:   "get-bucket-abac",
			Fields: fields_get_bucket_abac,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketAbacInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_abac, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketAbac(ctx, input)
			},
		},
		"get-bucket-accelerate-configuration": {
			Name:   "get-bucket-accelerate-configuration",
			Fields: fields_get_bucket_accelerate_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketAccelerateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_accelerate_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketAccelerateConfiguration(ctx, input)
			},
		},
		"get-bucket-acl": {
			Name:   "get-bucket-acl",
			Fields: fields_get_bucket_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketAclInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketAcl(ctx, input)
			},
		},
		"get-bucket-analytics-configuration": {
			Name:   "get-bucket-analytics-configuration",
			Fields: fields_get_bucket_analytics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketAnalyticsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_analytics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketAnalyticsConfiguration(ctx, input)
			},
		},
		"get-bucket-cors": {
			Name:   "get-bucket-cors",
			Fields: fields_get_bucket_cors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketCorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_cors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketCors(ctx, input)
			},
		},
		"get-bucket-encryption": {
			Name:   "get-bucket-encryption",
			Fields: fields_get_bucket_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketEncryption(ctx, input)
			},
		},
		"get-bucket-intelligent-tiering-configuration": {
			Name:   "get-bucket-intelligent-tiering-configuration",
			Fields: fields_get_bucket_intelligent_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketIntelligentTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_intelligent_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketIntelligentTieringConfiguration(ctx, input)
			},
		},
		"get-bucket-inventory-configuration": {
			Name:   "get-bucket-inventory-configuration",
			Fields: fields_get_bucket_inventory_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketInventoryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_inventory_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketInventoryConfiguration(ctx, input)
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
		"get-bucket-location": {
			Name:   "get-bucket-location",
			Fields: fields_get_bucket_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketLocation(ctx, input)
			},
		},
		"get-bucket-logging": {
			Name:   "get-bucket-logging",
			Fields: fields_get_bucket_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketLogging(ctx, input)
			},
		},
		"get-bucket-metadata-configuration": {
			Name:   "get-bucket-metadata-configuration",
			Fields: fields_get_bucket_metadata_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketMetadataConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_metadata_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketMetadataConfiguration(ctx, input)
			},
		},
		"get-bucket-metadata-table-configuration": {
			Name:   "get-bucket-metadata-table-configuration",
			Fields: fields_get_bucket_metadata_table_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketMetadataTableConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_metadata_table_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketMetadataTableConfiguration(ctx, input)
			},
		},
		"get-bucket-metrics-configuration": {
			Name:   "get-bucket-metrics-configuration",
			Fields: fields_get_bucket_metrics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketMetricsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_metrics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketMetricsConfiguration(ctx, input)
			},
		},
		"get-bucket-notification-configuration": {
			Name:   "get-bucket-notification-configuration",
			Fields: fields_get_bucket_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketNotificationConfiguration(ctx, input)
			},
		},
		"get-bucket-ownership-controls": {
			Name:   "get-bucket-ownership-controls",
			Fields: fields_get_bucket_ownership_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketOwnershipControlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_ownership_controls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketOwnershipControls(ctx, input)
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
		"get-bucket-policy-status": {
			Name:   "get-bucket-policy-status",
			Fields: fields_get_bucket_policy_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketPolicyStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_policy_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketPolicyStatus(ctx, input)
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
		"get-bucket-request-payment": {
			Name:   "get-bucket-request-payment",
			Fields: fields_get_bucket_request_payment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketRequestPaymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_request_payment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketRequestPayment(ctx, input)
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
		"get-bucket-website": {
			Name:   "get-bucket-website",
			Fields: fields_get_bucket_website,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketWebsiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_website, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketWebsite(ctx, input)
			},
		},
		"get-object": {
			Name:   "get-object",
			Fields: fields_get_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObject(ctx, input)
			},
		},
		"get-object-acl": {
			Name:   "get-object-acl",
			Fields: fields_get_object_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectAclInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectAcl(ctx, input)
			},
		},
		"get-object-attributes": {
			Name:   "get-object-attributes",
			Fields: fields_get_object_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectAttributes(ctx, input)
			},
		},
		"get-object-legal-hold": {
			Name:   "get-object-legal-hold",
			Fields: fields_get_object_legal_hold,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectLegalHoldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_legal_hold, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectLegalHold(ctx, input)
			},
		},
		"get-object-lock-configuration": {
			Name:   "get-object-lock-configuration",
			Fields: fields_get_object_lock_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectLockConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_lock_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectLockConfiguration(ctx, input)
			},
		},
		"get-object-retention": {
			Name:   "get-object-retention",
			Fields: fields_get_object_retention,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectRetentionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_retention, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectRetention(ctx, input)
			},
		},
		"get-object-tagging": {
			Name:   "get-object-tagging",
			Fields: fields_get_object_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectTagging(ctx, input)
			},
		},
		"get-object-torrent": {
			Name:   "get-object-torrent",
			Fields: fields_get_object_torrent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectTorrentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_torrent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectTorrent(ctx, input)
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
		"head-bucket": {
			Name:   "head-bucket",
			Fields: fields_head_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.HeadBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_head_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.HeadBucket(ctx, input)
			},
		},
		"head-object": {
			Name:   "head-object",
			Fields: fields_head_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.HeadObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_head_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.HeadObject(ctx, input)
			},
		},
		"list-bucket-analytics-configurations": {
			Name:   "list-bucket-analytics-configurations",
			Fields: fields_list_bucket_analytics_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBucketAnalyticsConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bucket_analytics_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBucketAnalyticsConfigurations(ctx, input)
			},
		},
		"list-bucket-intelligent-tiering-configurations": {
			Name:   "list-bucket-intelligent-tiering-configurations",
			Fields: fields_list_bucket_intelligent_tiering_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBucketIntelligentTieringConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bucket_intelligent_tiering_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBucketIntelligentTieringConfigurations(ctx, input)
			},
		},
		"list-bucket-inventory-configurations": {
			Name:   "list-bucket-inventory-configurations",
			Fields: fields_list_bucket_inventory_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBucketInventoryConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bucket_inventory_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBucketInventoryConfigurations(ctx, input)
			},
		},
		"list-bucket-metrics-configurations": {
			Name:   "list-bucket-metrics-configurations",
			Fields: fields_list_bucket_metrics_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBucketMetricsConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_bucket_metrics_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBucketMetricsConfigurations(ctx, input)
			},
		},
		"list-buckets": {
			Name:   "list-buckets",
			Fields: fields_list_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuckets(ctx, input)
				}
				var results []*svc.ListBucketsOutput
				p := svc.NewListBucketsPaginator(client, input)
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
		"list-directory-buckets": {
			Name:   "list-directory-buckets",
			Fields: fields_list_directory_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDirectoryBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_directory_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDirectoryBuckets(ctx, input)
				}
				var results []*svc.ListDirectoryBucketsOutput
				p := svc.NewListDirectoryBucketsPaginator(client, input)
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
		"list-multipart-uploads": {
			Name:   "list-multipart-uploads",
			Fields: fields_list_multipart_uploads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultipartUploadsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_multipart_uploads, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMultipartUploads(ctx, input)
			},
		},
		"list-object-versions": {
			Name:   "list-object-versions",
			Fields: fields_list_object_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_object_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListObjectVersions(ctx, input)
			},
		},
		"list-objects": {
			Name:   "list-objects",
			Fields: fields_list_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_objects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListObjects(ctx, input)
			},
		},
		"list-objects-v2": {
			Name:   "list-objects-v2",
			Fields: fields_list_objects_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_objects_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectsV2(ctx, input)
				}
				var results []*svc.ListObjectsV2Output
				p := svc.NewListObjectsV2Paginator(client, input)
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
		"list-parts": {
			Name:   "list-parts",
			Fields: fields_list_parts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_parts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParts(ctx, input)
				}
				var results []*svc.ListPartsOutput
				p := svc.NewListPartsPaginator(client, input)
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
		"put-bucket-abac": {
			Name:   "put-bucket-abac",
			Fields: fields_put_bucket_abac,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketAbacInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_abac, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketAbac(ctx, input)
			},
		},
		"put-bucket-accelerate-configuration": {
			Name:   "put-bucket-accelerate-configuration",
			Fields: fields_put_bucket_accelerate_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketAccelerateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_accelerate_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketAccelerateConfiguration(ctx, input)
			},
		},
		"put-bucket-acl": {
			Name:   "put-bucket-acl",
			Fields: fields_put_bucket_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketAclInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketAcl(ctx, input)
			},
		},
		"put-bucket-analytics-configuration": {
			Name:   "put-bucket-analytics-configuration",
			Fields: fields_put_bucket_analytics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketAnalyticsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_analytics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketAnalyticsConfiguration(ctx, input)
			},
		},
		"put-bucket-cors": {
			Name:   "put-bucket-cors",
			Fields: fields_put_bucket_cors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketCorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_cors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketCors(ctx, input)
			},
		},
		"put-bucket-encryption": {
			Name:   "put-bucket-encryption",
			Fields: fields_put_bucket_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketEncryption(ctx, input)
			},
		},
		"put-bucket-intelligent-tiering-configuration": {
			Name:   "put-bucket-intelligent-tiering-configuration",
			Fields: fields_put_bucket_intelligent_tiering_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketIntelligentTieringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_intelligent_tiering_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketIntelligentTieringConfiguration(ctx, input)
			},
		},
		"put-bucket-inventory-configuration": {
			Name:   "put-bucket-inventory-configuration",
			Fields: fields_put_bucket_inventory_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketInventoryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_inventory_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketInventoryConfiguration(ctx, input)
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
		"put-bucket-logging": {
			Name:   "put-bucket-logging",
			Fields: fields_put_bucket_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketLogging(ctx, input)
			},
		},
		"put-bucket-metrics-configuration": {
			Name:   "put-bucket-metrics-configuration",
			Fields: fields_put_bucket_metrics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketMetricsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_metrics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketMetricsConfiguration(ctx, input)
			},
		},
		"put-bucket-notification-configuration": {
			Name:   "put-bucket-notification-configuration",
			Fields: fields_put_bucket_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketNotificationConfiguration(ctx, input)
			},
		},
		"put-bucket-ownership-controls": {
			Name:   "put-bucket-ownership-controls",
			Fields: fields_put_bucket_ownership_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketOwnershipControlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_ownership_controls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketOwnershipControls(ctx, input)
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
		"put-bucket-request-payment": {
			Name:   "put-bucket-request-payment",
			Fields: fields_put_bucket_request_payment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketRequestPaymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_request_payment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketRequestPayment(ctx, input)
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
		"put-bucket-website": {
			Name:   "put-bucket-website",
			Fields: fields_put_bucket_website,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBucketWebsiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bucket_website, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBucketWebsite(ctx, input)
			},
		},
		"put-object": {
			Name:   "put-object",
			Fields: fields_put_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObject(ctx, input)
			},
		},
		"put-object-acl": {
			Name:   "put-object-acl",
			Fields: fields_put_object_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectAclInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObjectAcl(ctx, input)
			},
		},
		"put-object-legal-hold": {
			Name:   "put-object-legal-hold",
			Fields: fields_put_object_legal_hold,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectLegalHoldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object_legal_hold, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObjectLegalHold(ctx, input)
			},
		},
		"put-object-lock-configuration": {
			Name:   "put-object-lock-configuration",
			Fields: fields_put_object_lock_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectLockConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object_lock_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObjectLockConfiguration(ctx, input)
			},
		},
		"put-object-retention": {
			Name:   "put-object-retention",
			Fields: fields_put_object_retention,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectRetentionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object_retention, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObjectRetention(ctx, input)
			},
		},
		"put-object-tagging": {
			Name:   "put-object-tagging",
			Fields: fields_put_object_tagging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutObjectTaggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_object_tagging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutObjectTagging(ctx, input)
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
		"rename-object": {
			Name:   "rename-object",
			Fields: fields_rename_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenameObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rename_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenameObject(ctx, input)
			},
		},
		"restore-object": {
			Name:   "restore-object",
			Fields: fields_restore_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreObject(ctx, input)
			},
		},
		"select-object-content": {
			Name:   "select-object-content",
			Fields: fields_select_object_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SelectObjectContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_select_object_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SelectObjectContent(ctx, input)
			},
		},
		"update-bucket-metadata-inventory-table-configuration": {
			Name:   "update-bucket-metadata-inventory-table-configuration",
			Fields: fields_update_bucket_metadata_inventory_table_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBucketMetadataInventoryTableConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bucket_metadata_inventory_table_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBucketMetadataInventoryTableConfiguration(ctx, input)
			},
		},
		"update-bucket-metadata-journal-table-configuration": {
			Name:   "update-bucket-metadata-journal-table-configuration",
			Fields: fields_update_bucket_metadata_journal_table_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBucketMetadataJournalTableConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bucket_metadata_journal_table_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBucketMetadataJournalTableConfiguration(ctx, input)
			},
		},
		"update-object-encryption": {
			Name:   "update-object-encryption",
			Fields: fields_update_object_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateObjectEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_object_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateObjectEncryption(ctx, input)
			},
		},
		"upload-part": {
			Name:   "upload-part",
			Fields: fields_upload_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadPart(ctx, input)
			},
		},
		"upload-part-copy": {
			Name:   "upload-part-copy",
			Fields: fields_upload_part_copy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadPartCopyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_part_copy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadPartCopy(ctx, input)
			},
		},
		"write-get-object-response": {
			Name:   "write-get-object-response",
			Fields: fields_write_get_object_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.WriteGetObjectResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_write_get_object_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.WriteGetObjectResponse(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("s3", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
