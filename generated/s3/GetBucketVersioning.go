package s3

// GetBucketVersioning is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the versioning state of a bucket.
//
// To retrieve the versioning state of a bucket, you must be the bucket owner.
//
// This implementation also returns the MFA Delete status of the versioning state.
// If the MFA Delete status is enabled , the bucket owner must use an
// authentication device to change the versioning state of the bucket.
//
// The following operations are related to GetBucketVersioning :
//
// [GetObject]
//
// [PutObject]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
