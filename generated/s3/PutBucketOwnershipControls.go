package s3

// PutBucketOwnershipControls is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Creates or modifies OwnershipControls for an Amazon S3 bucket. To use this
// operation, you must have the s3:PutBucketOwnershipControls permission. For more
// information about Amazon S3 permissions, see [Specifying permissions in a policy].
//
// For information about Amazon S3 Object Ownership, see [Using object ownership].
//
// The following operations are related to PutBucketOwnershipControls :
//
// # GetBucketOwnershipControls
//
// # DeleteBucketOwnershipControls
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Specifying permissions in a policy]: https://docs.aws.amazon.com/AmazonS3/latest/user-guide/using-with-s3-actions.html
// [Using object ownership]: https://docs.aws.amazon.com/AmazonS3/latest/user-guide/about-object-ownership.html
