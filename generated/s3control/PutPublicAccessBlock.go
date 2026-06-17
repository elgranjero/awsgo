package s3control

// PutPublicAccessBlock is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Creates or modifies the PublicAccessBlock configuration for an Amazon Web
// Services account. This operation may be restricted when the account is managed
// by organization-level Block Public Access policies. You might get an Access
// Denied (403) error when the account is managed by organization-level Block
// Public Access policies. Organization-level policies override account-level
// settings, preventing direct account-level modifications. For this operation,
// users must have the s3:PutAccountPublicAccessBlock permission. For more
// information, see [Using Amazon S3 block public access].
//
// Related actions include:
//
// [GetPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html
// [Using Amazon S3 block public access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
