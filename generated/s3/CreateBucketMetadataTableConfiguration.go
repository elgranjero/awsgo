package s3

// CreateBucketMetadataTableConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// We recommend that you create your S3 Metadata configurations by using the V2 [CreateBucketMetadataConfiguration]
//
// API operation. We no longer recommend using the V1
// CreateBucketMetadataTableConfiguration API operation.
//
// If you created your S3 Metadata configuration before July 15, 2025, we
// recommend that you delete and re-create your configuration by using [CreateBucketMetadataConfiguration]so that you
// can expire journal table records and create a live inventory table.
//
// Creates a V1 S3 Metadata configuration for a general purpose bucket. For more
// information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the following permissions. For
// more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you want to encrypt your metadata tables with server-side encryption with
// Key Management Service (KMS) keys (SSE-KMS), you need additional permissions.
// For more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you also want to integrate your table bucket with Amazon Web Services
// analytics services so that you can query your metadata table, you need
// additional permissions. For more information, see [Integrating Amazon S3 Tables with Amazon Web Services analytics services]in the Amazon S3 User Guide.
//
// - s3:CreateBucketMetadataTableConfiguration
//
// - s3tables:CreateNamespace
//
// - s3tables:GetTable
//
// - s3tables:CreateTable
//
// - s3tables:PutTablePolicy
//
// The following operations are related to CreateBucketMetadataTableConfiguration :
//
// [DeleteBucketMetadataTableConfiguration]
//
// [GetBucketMetadataTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [GetBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html
// [DeleteBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataTableConfiguration.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [Integrating Amazon S3 Tables with Amazon Web Services analytics services]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-aws.html
