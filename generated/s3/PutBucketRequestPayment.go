package s3

// PutBucketRequestPayment is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Sets the request payment configuration for a bucket. By default, the bucket
// owner pays for downloads from the bucket. This configuration parameter enables
// the bucket owner (only) to specify that the person requesting the download will
// be charged for the download. For more information, see [Requester Pays Buckets].
//
// The following operations are related to PutBucketRequestPayment :
//
// [CreateBucket]
//
// [GetBucketRequestPayment]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketRequestPayment]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketRequestPayment.html
// [Requester Pays Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
