package iam

// UpdateOpenIDConnectProviderThumbprint is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Replaces the existing list of server certificate thumbprints associated with an
// OpenID Connect (OIDC) provider resource object with a new list of thumbprints.
//
// The list that you pass with this operation completely replaces the existing
// list of thumbprints. (The lists are not merged.)
//
// Typically, you need to update a thumbprint only when the identity provider
// certificate changes, which occurs rarely. However, if the provider's certificate
// does change, any attempt to assume an IAM role that specifies the OIDC provider
// as a principal fails until the certificate thumbprint is updated.
//
// Amazon Web Services secures communication with OIDC identity providers (IdPs)
// using our library of trusted root certificate authorities (CAs) to verify the
// JSON Web Key Set (JWKS) endpoint's TLS certificate. If your OIDC IdP relies on a
// certificate that is not signed by one of these trusted CAs, only then we secure
// communication using the thumbprints set in the IdP's configuration.
//
// Trust for the OIDC provider is derived from the provider certificate and is
// validated by the thumbprint. Therefore, it is best to limit access to the
// UpdateOpenIDConnectProviderThumbprint operation to highly privileged users.
