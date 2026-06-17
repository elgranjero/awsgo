package iam

// DeleteSAMLProvider is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes a SAML provider resource in IAM.
//
// Deleting the provider resource from IAM does not update any roles that
// reference the SAML provider resource's ARN as a principal in their trust
// policies. Any attempt to assume a role that references a non-existent provider
// resource ARN fails.
//
// This operation requires [Signature Version 4].
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
