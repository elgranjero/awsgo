package s3

// GetBucketPolicyStatus is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Retrieves the policy status for an Amazon S3 bucket, indicating whether the
// bucket is public. In order to use this operation, you must have the
// s3:GetBucketPolicyStatus permission. For more information about Amazon S3
// permissions, see [Specifying Permissions in a Policy].
//
// For more information about when Amazon S3 considers a bucket public, see [The Meaning of "Public"].
//
// The following operations are related to GetBucketPolicyStatus :
//
// [Using Amazon S3 Block Public Access]
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
