package s3control

// CreateAccessGrant is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Creates an access grant that gives a grantee access to your S3 data. The
// grantee can be an IAM user or role or a directory user, or group. Before you can
// create a grant, you must have an S3 Access Grants instance in the same Region as
// the S3 data. You can create an S3 Access Grants instance using the [CreateAccessGrantsInstance]. You must
// also have registered at least one S3 data location in your S3 Access Grants
// instance using [CreateAccessGrantsLocation].
//
// Permissions You must have the s3:CreateAccessGrant permission to use this
// operation.
//
// Additional Permissions For any directory identity - sso:DescribeInstance and
// sso:DescribeApplication
//
// For directory users - identitystore:DescribeUser
//
// For directory groups - identitystore:DescribeGroup
//
// [CreateAccessGrantsLocation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessGrantsLocation.html
// [CreateAccessGrantsInstance]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessGrantsInstance.html
