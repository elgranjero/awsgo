package s3control

// DeletePublicAccessBlock is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Removes the PublicAccessBlock configuration for an Amazon Web Services account.
// This operation might be restricted when the account is managed by
// organization-level Block Public Access policies. You’ll get an Access Denied
// (403) error when the account is managed by organization-level Block Public
// Access policies. Organization-level policies override account-level settings,
// preventing direct account-level modifications. For more information, see [Using Amazon S3 block public access].
//
// Related actions include:
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html
// [Using Amazon S3 block public access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
