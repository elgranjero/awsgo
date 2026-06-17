package s3

// DeleteBucketCors is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Deletes the cors configuration information set for the bucket.
//
// To use this operation, you must have permission to perform the s3:PutBucketCORS
// action. The bucket owner has this permission by default and can grant this
// permission to others.
//
// For information about cors , see [Enabling Cross-Origin Resource Sharing] in the Amazon S3 User Guide.
//
// # Related Resources
//
// [PutBucketCors]
//
// [RESTOPTIONSobject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [RESTOPTIONSobject]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html
