package verifiedpermissions

// CreateIdentitySource is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Adds an identity source to a policy store–an Amazon Cognito user pool or OpenID
// Connect (OIDC) identity provider (IdP).
//
// After you create an identity source, you can use the identities provided by the
// IdP as proxies for the principal in authorization queries that use the [IsAuthorizedWithToken]or [BatchIsAuthorizedWithToken] API
// operations. These identities take the form of tokens that contain claims about
// the user, such as IDs, attributes and group memberships. Identity sources
// provide identity (ID) tokens and access tokens. Verified Permissions derives
// information about your user and session from token claims. Access tokens provide
// action context to your policies, and ID tokens provide principal Attributes .
//
// Tokens from an identity source user continue to be usable until they expire.
// Token revocation and resource deletion have no effect on the validity of a token
// in your policy store
//
// To reference a user from this identity source in your Cedar policies, refer to
// the following syntax examples.
//
// - Amazon Cognito user pool: Namespace::[Entity type]::[User pool ID]|[user
// principal attribute] , for example
// MyCorp::User::us-east-1_EXAMPLE|a1b2c3d4-5678-90ab-cdef-EXAMPLE11111 .
//
// - OpenID Connect (OIDC) provider: Namespace::[Entity
// type]::[entityIdPrefix]|[user principal attribute] , for example
// MyCorp::User::MyOIDCProvider|a1b2c3d4-5678-90ab-cdef-EXAMPLE22222 .
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [IsAuthorizedWithToken]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [BatchIsAuthorizedWithToken]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorizedWithToken.html
