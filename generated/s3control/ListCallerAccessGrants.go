package s3control

// ListCallerAccessGrants is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Use this API to list the access grants that grant the caller access to Amazon
// S3 data through S3 Access Grants. The caller (grantee) can be an Identity and
// Access Management (IAM) identity or Amazon Web Services Identity Center
// corporate directory identity. You must pass the Amazon Web Services account of
// the S3 data owner (grantor) in the request. You can, optionally, narrow the
// results by GrantScope , using a fragment of the data's S3 path, and S3 Access
// Grants will return only the grants with a path that contains the path fragment.
// You can also pass the AllowedByApplication filter in the request, which returns
// only the grants authorized for applications, whether the application is the
// caller's Identity Center application or any other application ( ALL ). For more
// information, see [List the caller's access grants]in the Amazon S3 User Guide.
//
// Permissions You must have the s3:ListCallerAccessGrants permission to use this
// operation.
//
// [List the caller's access grants]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-list-grants.html
