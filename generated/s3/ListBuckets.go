package s3

// ListBuckets is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns a list of all buckets owned by the authenticated sender of the request.
// To grant IAM permission to use this operation, you must add the
// s3:ListAllMyBuckets policy action.
//
// For information about Amazon S3 buckets, see [Creating, configuring, and working with Amazon S3 buckets].
//
// We strongly recommend using only paginated ListBuckets requests. Unpaginated
// ListBuckets requests are only supported for Amazon Web Services accounts set to
// the default general purpose bucket quota of 10,000. If you have an approved
// general purpose bucket quota above 10,000, you must send paginated ListBuckets
// requests to list your account’s buckets. All unpaginated ListBuckets requests
// will be rejected for Amazon Web Services accounts with a general purpose bucket
// quota greater than 10,000.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Creating, configuring, and working with Amazon S3 buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html
