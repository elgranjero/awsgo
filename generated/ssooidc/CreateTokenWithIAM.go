package ssooidc

// CreateTokenWithIAM is generated as a reference stub.
// Executable command wiring lives under cmd/ssooidc.go.
//
// Creates and returns access and refresh tokens for authorized client
// applications that are authenticated using any IAM entity, such as a service role
// or user. These tokens might contain defined scopes that specify permissions such
// as read:profile or write:data . Through downscoping, you can use the scopes
// parameter to request tokens with reduced permissions compared to the original
// client application's permissions or, if applicable, the refresh token's scopes.
// The access token can be used to fetch short-lived credentials for the assigned
// Amazon Web Services accounts or to access application APIs using bearer
// authentication.
//
// This API is used with Signature Version 4. For more information, see [Amazon Web Services Signature Version 4 for API Requests].
//
// [Amazon Web Services Signature Version 4 for API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html
