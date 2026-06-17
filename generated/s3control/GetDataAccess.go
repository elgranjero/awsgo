package s3control

// GetDataAccess is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Returns a temporary access credential from S3 Access Grants to the grantee or
// client application. The [temporary credential]is an Amazon Web Services STS token that grants them
// access to the S3 data.
//
// Permissions You must have the s3:GetDataAccess permission to use this
// operation.
//
// Additional Permissions The IAM role that S3 Access Grants assumes must have the
// following permissions specified in the trust policy when registering the
// location: sts:AssumeRole , for directory users or groups sts:SetContext , and
// for IAM users or roles sts:SetSourceIdentity .
//
// [temporary credential]: https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html
