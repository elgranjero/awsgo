package s3control

// CreateAccessGrantsLocation is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// The S3 data location that you would like to register in your S3 Access Grants
// instance. Your S3 data must be in the same Region as your S3 Access Grants
// instance. The location can be one of the following:
//
// - The default S3 location s3://
//
// - A bucket - S3://
//
// - A bucket and prefix - S3:///
//
// When you register a location, you must include the IAM role that has permission
// to manage the S3 location that you are registering. Give S3 Access Grants
// permission to assume this role [using a policy]. S3 Access Grants assumes this role to manage
// access to the location and to vend temporary credentials to grantees or client
// applications.
//
// Permissions You must have the s3:CreateAccessGrantsLocation permission to use
// this operation.
//
// Additional Permissions You must also have the following permission for the
// specified IAM role: iam:PassRole
//
// [using a policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-location.html
