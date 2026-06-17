package qapps

// CreatePresignedUrl is generated as a reference stub.
// Executable command wiring lives under cmd/qapps.go.
//
// Creates a presigned URL for an S3 POST operation to upload a file. You can use
// this URL to set a default file for a FileUploadCard in a Q App definition or to
// provide a file for a single Q App run. The scope parameter determines how the
// file will be used, either at the app definition level or the app session level.
//
// The IAM permissions are derived from the qapps:ImportDocument action. For more
// information on the IAM policy for Amazon Q Apps, see [IAM permissions for using Amazon Q Apps].
//
// [IAM permissions for using Amazon Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/deploy-q-apps-iam-permissions.html
