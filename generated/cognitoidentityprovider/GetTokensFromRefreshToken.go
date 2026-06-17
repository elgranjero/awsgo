package cognitoidentityprovider

// GetTokensFromRefreshToken is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Given a refresh token, issues new ID, access, and optionally refresh tokens for
// the user who owns the submitted token. This operation issues a new refresh token
// and invalidates the original refresh token after an optional grace period when
// refresh token rotation is enabled. If refresh token rotation is disabled, issues
// new ID and access tokens only.
