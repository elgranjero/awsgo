package s3

// GetPublicAccessBlock is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Retrieves the PublicAccessBlock configuration for an Amazon S3 bucket. This
// operation returns the bucket-level configuration only. To understand the
// effective public access behavior, you must also consider account-level settings
// (which may inherit from organization-level policies). To use this operation, you
// must have the s3:GetBucketPublicAccessBlock permission. For more information
// about Amazon S3 permissions, see [Specifying Permissions in a Policy].
//
// When Amazon S3 evaluates the PublicAccessBlock configuration for a bucket or an
// object, it checks the PublicAccessBlock configuration for both the bucket (or
// the bucket that contains the object) and the bucket owner's account.
// Account-level settings automatically inherit from organization-level policies
// when present. If the PublicAccessBlock settings are different between the
// bucket and the account, Amazon S3 uses the most restrictive combination of the
// bucket-level and account-level settings.
//
// For more information about when Amazon S3 considers a bucket or an object
// public, see [The Meaning of "Public"].
//
// The following operations are related to GetPublicAccessBlock :
//
// [Using Amazon S3 Block Public Access]
//
// [PutPublicAccessBlock]
//
// [GetPublicAccessBlock]
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
