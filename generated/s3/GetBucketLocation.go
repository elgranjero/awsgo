package s3

// GetBucketLocation is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Using the GetBucketLocation operation is no longer a best practice. To return
// the Region that a bucket resides in, we recommend that you use the [HeadBucket]operation
// instead. For backward compatibility, Amazon S3 continues to support the
// GetBucketLocation operation.
//
// Returns the Region the bucket resides in. You set the bucket's Region using the
// LocationConstraint request parameter in a CreateBucket request. For more
// information, see [CreateBucket].
//
// In a bucket's home Region, calls to the GetBucketLocation operation are
// governed by the bucket's policy. In other Regions, the bucket policy doesn't
// apply, which means that cross-account access won't be authorized. However, calls
// to the HeadBucket operation always return the bucket’s location through an HTTP
// response header, whether access to the bucket is authorized or not. Therefore,
// we recommend using the HeadBucket operation for bucket Region discovery and to
// avoid using the GetBucketLocation operation.
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
// This operation is not supported for directory buckets.
//
// The following operations are related to GetBucketLocation :
//
// [GetObject]
//
// [CreateBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [HeadBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html
