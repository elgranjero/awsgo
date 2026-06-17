package iam

// CreateSAMLProvider is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Creates an IAM resource that describes an identity provider (IdP) that supports
// SAML 2.0.
//
// The SAML provider resource that you create with this operation can be used as a
// principal in an IAM role's trust policy. Such a policy can enable federated
// users who sign in using the SAML IdP to assume the role. You can create an IAM
// role that supports Web-based single sign-on (SSO) to the Amazon Web Services
// Management Console or one that supports API access to Amazon Web Services.
//
// When you create the SAML provider resource, you upload a SAML metadata document
// that you get from your IdP. That document includes the issuer's name, expiration
// information, and keys that can be used to validate the SAML authentication
// response (assertions) that the IdP sends. You must generate the metadata
// document using the identity management software that is used as your
// organization's IdP.
//
// This operation requires [Signature Version 4].
//
// For more information, see [Enabling SAML 2.0 federated users to access the Amazon Web Services Management Console] and [About SAML 2.0-based federation] in the IAM User Guide.
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [About SAML 2.0-based federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Enabling SAML 2.0 federated users to access the Amazon Web Services Management Console]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-saml.html
