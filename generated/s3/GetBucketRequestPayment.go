package s3

// GetBucketRequestPayment is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the request payment configuration of a bucket. To use this version of
// the operation, you must be the bucket owner. For more information, see [Requester Pays Buckets].
//
// The following operations are related to GetBucketRequestPayment :
//
// [ListObjects]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
// [Requester Pays Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html
