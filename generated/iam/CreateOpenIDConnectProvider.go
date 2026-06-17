package iam

// CreateOpenIDConnectProvider is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Creates an IAM entity to describe an identity provider (IdP) that supports [OpenID Connect (OIDC)].
//
// The OIDC provider that you create with this operation can be used as a
// principal in a role's trust policy. Such a policy establishes a trust
// relationship between Amazon Web Services and the OIDC provider.
//
// If you are using an OIDC identity provider from Google, Facebook, or Amazon
// Cognito, you don't need to create a separate IAM identity provider. These OIDC
// identity providers are already built-in to Amazon Web Services and are available
// for your use. Instead, you can move directly to creating new roles using your
// identity provider. To learn more, see [Creating a role for web identity or OpenID connect federation]in the IAM User Guide.
//
// When you create the IAM OIDC provider, you specify the following:
//
// - The URL of the OIDC identity provider (IdP) to trust
//
// - A list of client IDs (also known as audiences) that identify the
// application or applications allowed to authenticate using the OIDC provider
//
// - A list of tags that are attached to the specified IAM OIDC provider
//
// - A list of thumbprints of one or more server certificates that the IdP uses
//
// You get all of this information from the OIDC IdP you want to use to access
// Amazon Web Services.
//
// Amazon Web Services secures communication with OIDC identity providers (IdPs)
// using our library of trusted root certificate authorities (CAs) to verify the
// JSON Web Key Set (JWKS) endpoint's TLS certificate. If your OIDC IdP relies on a
// certificate that is not signed by one of these trusted CAs, only then we secure
// communication using the thumbprints set in the IdP's configuration.
//
// The trust for the OIDC provider is derived from the IAM provider that this
// operation creates. Therefore, it is best to limit access to the [CreateOpenIDConnectProvider]operation to
// highly privileged users.
//
// [OpenID Connect (OIDC)]: http://openid.net/connect/
// [Creating a role for web identity or OpenID connect federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_oidc.html
// [CreateOpenIDConnectProvider]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html
