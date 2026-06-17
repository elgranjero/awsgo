package iam

// DeleteOpenIDConnectProvider is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes an OpenID Connect identity provider (IdP) resource object in IAM.
//
// Deleting an IAM OIDC provider resource does not update any roles that reference
// the provider as a principal in their trust policies. Any attempt to assume a
// role that references a deleted provider fails.
//
// This operation is idempotent; it does not fail or return an error if you call
// the operation for a provider that does not exist.
