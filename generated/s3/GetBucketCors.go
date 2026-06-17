package s3

// GetBucketCors is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the Cross-Origin Resource Sharing (CORS) configuration information set
// for the bucket.
//
// To use this operation, you must have permission to perform the s3:GetBucketCORS
// action. By default, the bucket owner has this permission and can grant it to
// others.
//
// When you use this API operation with an access point, provide the alias of the
// access point in place of the bucket name.
//
// When you use this API operation with an Object Lambda access point, provide the
// alias of the Object Lambda access point in place of the bucket name. If the
// Object Lambda access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see [List of Error Codes].
//
// For more information about CORS, see [Enabling Cross-Origin Resource Sharing].
//
// The following operations are related to GetBucketCors :
//
// [PutBucketCors]
//
// [DeleteBucketCors]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [DeleteBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html
