package codeartifact

// GetAuthorizationToken is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Generates a temporary authorization token for accessing repositories in the
//
// domain. This API requires the codeartifact:GetAuthorizationToken and
// sts:GetServiceBearerToken permissions. For more information about authorization
// tokens, see [CodeArtifact authentication and tokens].
//
// CodeArtifact authorization tokens are valid for a period of 12 hours when
// created with the login command. You can call login periodically to refresh the
// token. When you create an authorization token with the GetAuthorizationToken
// API, you can set a custom authorization period, up to a maximum of 12 hours,
// with the durationSeconds parameter.
//
// The authorization period begins after login or GetAuthorizationToken is called.
// If login or GetAuthorizationToken is called while assuming a role, the token
// lifetime is independent of the maximum session duration of the role. For
// example, if you call sts assume-role and specify a session duration of 15
// minutes, then generate a CodeArtifact authorization token, the token will be
// valid for the full authorization period even though this is longer than the
// 15-minute session duration.
//
// See [Using IAM Roles] for more information on controlling session duration.
//
// [Using IAM Roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html
// [CodeArtifact authentication and tokens]: https://docs.aws.amazon.com/codeartifact/latest/ug/tokens-authentication.html
