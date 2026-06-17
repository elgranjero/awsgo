package s3

// GetBucketAcl is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// This implementation of the GET action uses the acl subresource to return the
// access control list (ACL) of a bucket. To use GET to return the ACL of the
// bucket, you must have the READ_ACP access to the bucket. If READ_ACP permission
// is granted to the anonymous user, you can return the ACL of the bucket without
// using an authorization header.
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
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// requests to read ACLs are still supported and return the
// bucket-owner-full-control ACL with the owner being the account that created the
// bucket. For more information, see [Controlling object ownership and disabling ACLs]in the Amazon S3 User Guide.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// The following operations are related to GetBucketAcl :
//
// [ListObjects]
//
// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [Controlling object ownership and disabling ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
