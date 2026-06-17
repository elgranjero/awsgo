package redshift

// GetIdentityCenterAuthToken is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Generates an encrypted authentication token that propagates the caller's Amazon
// Web Services IAM Identity Center identity to Amazon Redshift clusters. This API
// extracts the Amazon Web Services IAM Identity Center identity from enhanced
// credentials and creates a secure token that Amazon Redshift drivers can use for
// authentication.
//
// The token is encrypted using Key Management Service (KMS) and can only be
// decrypted by the specified Amazon Redshift clusters. The token contains the
// caller's Amazon Web Services IAM Identity Center identity information and is
// valid for a limited time period.
//
// This API is exclusively for use with Amazon Web Services IAM Identity Center
// enhanced credentials. If the caller is not using enhanced credentials with
// embedded Amazon Web Services IAM Identity Center identity, the API will return
// an error.
