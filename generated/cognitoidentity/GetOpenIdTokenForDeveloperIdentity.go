package cognitoidentity

// GetOpenIdTokenForDeveloperIdentity is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentity.go.
//
// Registers (or retrieves) a Cognito IdentityId and an OpenID Connect token for a
// user authenticated by your backend authentication process. Supplying multiple
// logins will create an implicit linked account. You can only specify one
// developer provider as part of the Logins map, which is linked to the identity
// pool. The developer provider is the "domain" by which Cognito will refer to your
// users.
//
// You can use GetOpenIdTokenForDeveloperIdentity to create a new identity and to
// link new logins (that is, user credentials issued by a public provider or
// developer provider) to an existing identity. When you want to create a new
// identity, the IdentityId should be null. When you want to associate a new login
// with an existing authenticated/unauthenticated identity, you can do so by
// providing the existing IdentityId . This API will create the identity in the
// specified IdentityPoolId .
//
// You must use Amazon Web Services developer credentials to call this operation.
