package ssoadmin

// PutApplicationGrant is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Creates a configuration for an application to use grants. Conceptually grants
// are authorization to request actions related to tokens. This configuration will
// be used when parties are requesting and receiving tokens during the trusted
// identity propagation process. For more information on the IAM Identity Center
// supported grant workflows, see [SAML 2.0 and OAuth 2.0].
//
// A grant is created between your applications and Identity Center instance which
// enables an application to use specified mechanisms to obtain tokens. These
// tokens are used by your applications to gain access to Amazon Web Services
// resources on behalf of users. The following elements are within these exchanges:
//
// - Requester - The application requesting access to Amazon Web Services
// resources.
//
// - Subject - Typically the user that is requesting access to Amazon Web
// Services resources.
//
// - Grant - Conceptually, a grant is authorization to access Amazon Web
// Services resources. These grants authorize token generation for authenticating
// access to the requester and for the request to make requests on behalf of the
// subjects. There are four types of grants:
//
// - AuthorizationCode - Allows an application to request authorization through
// a series of user-agent redirects.
//
// - JWT bearer - Authorizes an application to exchange a JSON Web Token that
// came from an external identity provider. To learn more, see [RFC 6479].
//
// - Refresh token - Enables application to request new access tokens to replace
// expiring or expired access tokens.
//
// - Exchange token - A grant that requests tokens from the authorization server
// by providing a ‘subject’ token with access scope authorizing trusted identity
// propagation to this application. To learn more, see [RFC 8693].
//
// - Authorization server - IAM Identity Center requests tokens.
//
// User credentials are never shared directly within these exchanges. Instead,
// applications use grants to request access tokens from IAM Identity Center. For
// more information, see [RFC 6479].
//
// Use cases
//
// - Connecting to custom applications.
//
// - Configuring an Amazon Web Services service to make calls to another Amazon
// Web Services services using JWT tokens.
//
// [RFC 6479]: https://datatracker.ietf.org/doc/html/rfc6749
// [SAML 2.0 and OAuth 2.0]: https://docs.aws.amazon.com/singlesignon/latest/userguide/customermanagedapps-saml2-oauth2.html
// [RFC 8693]: https://datatracker.ietf.org/doc/html/rfc8693
