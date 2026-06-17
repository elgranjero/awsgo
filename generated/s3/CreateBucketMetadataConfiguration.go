package s3

// CreateBucketMetadataConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Creates an S3 Metadata V2 metadata configuration for a general purpose bucket.
// For more information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the following permissions. For
// more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you want to encrypt your metadata tables with server-side encryption with
// Key Management Service (KMS) keys (SSE-KMS), you need additional permissions in
// your KMS key policy. For more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you also want to integrate your table bucket with Amazon Web Services
// analytics services so that you can query your metadata table, you need
// additional permissions. For more information, see [Integrating Amazon S3 Tables with Amazon Web Services analytics services]in the Amazon S3 User Guide.
//
// To query your metadata tables, you need additional permissions. For more
// information, see [Permissions for querying metadata tables]in the Amazon S3 User Guide.
//
// - s3:CreateBucketMetadataTableConfiguration
//
// The IAM policy action name is the same for the V1 and V2 API operations.
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
// The following operations are related to CreateBucketMetadataConfiguration :
//
// [DeleteBucketMetadataConfiguration]
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
// [Permissions for querying metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-bucket-query-permissions.html
// [UpdateBucketMetadataInventoryTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html
// [DeleteBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html
// [Integrating Amazon S3 Tables with Amazon Web Services analytics services]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-aws.html
