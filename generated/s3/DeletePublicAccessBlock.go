package s3

// DeletePublicAccessBlock is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Removes the PublicAccessBlock configuration for an Amazon S3 bucket. This
// operation removes the bucket-level configuration only. The effective public
// access behavior will still be governed by account-level settings (which may
// inherit from organization-level policies). To use this operation, you must have
// the s3:PutBucketPublicAccessBlock permission. For more information about
// permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// The following operations are related to DeletePublicAccessBlock :
//
// [Using Amazon S3 Block Public Access]
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [GetBucketPolicyStatus]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [GetBucketPolicyStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html
