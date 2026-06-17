package s3

// GetObjectAcl is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the access control list (ACL) of an object. To use this operation, you
// must have s3:GetObjectAcl permissions or READ_ACP access to the object. For
// more information, see [Mapping of ACL permissions and access policy permissions]in the Amazon S3 User Guide
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// By default, GET returns ACL information about the current version of an object.
// To return ACL information about a different version, use the versionId
// subresource.
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// requests to read ACLs are still supported and return the
// bucket-owner-full-control ACL with the owner being the account that created the
// bucket. For more information, see [Controlling object ownership and disabling ACLs]in the Amazon S3 User Guide.
//
// The following operations are related to GetObjectAcl :
//
// [GetObject]
//
// [GetObjectAttributes]
//
// [DeleteObject]
//
// [PutObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [Mapping of ACL permissions and access policy permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#acl-access-policy-permission-mapping
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Controlling object ownership and disabling ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
