package acmpca

// UpdateCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Updates the status or configuration of a private certificate authority (CA).
// Your private CA must be in the ACTIVE or DISABLED state before you can update
// it. You can disable a private CA that is in the ACTIVE state or make a CA that
// is in the DISABLED state active again.
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
