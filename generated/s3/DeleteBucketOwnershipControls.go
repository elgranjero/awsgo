package s3

// DeleteBucketOwnershipControls is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Removes OwnershipControls for an Amazon S3 bucket. To use this operation, you
// must have the s3:PutBucketOwnershipControls permission. For more information
// about Amazon S3 permissions, see [Specifying Permissions in a Policy].
//
// For information about Amazon S3 Object Ownership, see [Using Object Ownership].
//
// The following operations are related to DeleteBucketOwnershipControls :
//
// # GetBucketOwnershipControls
//
// # PutBucketOwnershipControls
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Using Object Ownership]: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
