package s3

// PutObjectLegalHold is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Applies a legal hold configuration to the specified object. For more
// information, see [Locking Objects].
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
