package s3

// DeleteBucketMetadataConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Deletes an S3 Metadata configuration from a general purpose bucket. For more
//
// information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// You can use the V2 DeleteBucketMetadataConfiguration API operation with V1 or
// V2 metadata configurations. However, if you try to use the V1
// DeleteBucketMetadataTableConfiguration API operation with V2 configurations, you
// will receive an HTTP 405 Method Not Allowed error.
//
// Permissions To use this operation, you must have the
// s3:DeleteBucketMetadataTableConfiguration permission. For more information, see [Setting up permissions for configuring metadata tables]
// in the Amazon S3 User Guide.
//
// The IAM policy action name is the same for the V1 and V2 API operations.
//
// The following operations are related to DeleteBucketMetadataConfiguration :
//
// [CreateBucketMetadataConfiguration]
//
// [GetBucketMetadataConfiguration]
//
// [UpdateBucketMetadataInventoryTableConfiguration]
//
// [UpdateBucketMetadataJournalTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [UpdateBucketMetadataJournalTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [UpdateBucketMetadataInventoryTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html
