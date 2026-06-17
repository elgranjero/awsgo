package s3

// GetBucketAccelerateConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// This implementation of the GET action uses the accelerate subresource to return
// the Transfer Acceleration state of a bucket, which is either Enabled or
// Suspended . Amazon S3 Transfer Acceleration is a bucket-level feature that
// enables you to perform faster data transfers to and from Amazon S3.
//
// To use this operation, you must have permission to perform the
// s3:GetAccelerateConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to your Amazon S3 Resources] in the Amazon S3 User Guide.
//
// You set the Transfer Acceleration state of an existing bucket to Enabled or
// Suspended by using the [PutBucketAccelerateConfiguration] operation.
//
// A GET accelerate request does not return a state value for a bucket that has no
// transfer acceleration state. A bucket has no Transfer Acceleration state if a
// state has never been set on the bucket.
//
// For more information about transfer acceleration, see [Transfer Acceleration] in the Amazon S3 User
// Guide.
//
// The following operations are related to GetBucketAccelerateConfiguration :
//
// [PutBucketAccelerateConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketAccelerateConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAccelerateConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Managing Access Permissions to your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Transfer Acceleration]: https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html
