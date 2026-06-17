package cognitoidentityprovider

// SetRiskConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Configures threat protection for a user pool or app client. Sets configuration
// for the following.
//
// - Responses to risks with adaptive authentication
//
// - Responses to vulnerable passwords with compromised-credentials detection
//
// - Notifications to users who have had risky activity detected
//
// - IP-address denylist and allowlist
//
// To set the risk configuration for the user pool to defaults, send this request
// with only the UserPoolId parameter. To reset the threat protection settings of
// an app client to be inherited from the user pool, send UserPoolId and ClientId
// parameters only. To change threat protection to audit-only or off, update the
// value of UserPoolAddOns in an UpdateUserPool request. To activate this setting,
// your user pool must be on the [Plus tier].
//
// [Plus tier]: https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-plus.html
