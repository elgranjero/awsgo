package s3control

// DeleteAccessGrantsInstance is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Deletes your S3 Access Grants instance. You must first delete the access grants
// and locations before S3 Access Grants can delete the instance. See [DeleteAccessGrant]and [DeleteAccessGrantsLocation]. If you
// have associated an IAM Identity Center instance with your S3 Access Grants
// instance, you must first dissassociate the Identity Center instance from the S3
// Access Grants instance before you can delete the S3 Access Grants instance. See [AssociateAccessGrantsIdentityCenter]
// and [DissociateAccessGrantsIdentityCenter].
//
// Permissions You must have the s3:DeleteAccessGrantsInstance permission to use
// this operation.
//
// [DeleteAccessGrant]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessGrant.html
// [AssociateAccessGrantsIdentityCenter]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AssociateAccessGrantsIdentityCenter.html
// [DeleteAccessGrantsLocation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessGrantsLocation.html
// [DissociateAccessGrantsIdentityCenter]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DissociateAccessGrantsIdentityCenter.html
