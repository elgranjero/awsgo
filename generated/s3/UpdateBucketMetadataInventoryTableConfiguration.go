package s3

// UpdateBucketMetadataInventoryTableConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Enables or disables a live inventory table for an S3 Metadata configuration on
// a general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the following permissions. For
// more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you want to encrypt your inventory table with server-side encryption with
// Key Management Service (KMS) keys (SSE-KMS), you need additional permissions in
// your KMS key policy. For more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// - s3:UpdateBucketMetadataInventoryTableConfiguration
//
// - s3tables:CreateTableBucket
//
// - s3tables:CreateNamespace
//
// - s3tables:GetTable
//
// - s3tables:CreateTable
//
// - s3tables:PutTablePolicy
//
// - s3tables:PutTableEncryption
//
// - kms:DescribeKey
//
// The following operations are related to
// UpdateBucketMetadataInventoryTableConfiguration :
//
// [CreateBucketMetadataConfiguration]
//
// [DeleteBucketMetadataConfiguration]
//
// [GetBucketMetadataConfiguration]
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
// [DeleteBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html
