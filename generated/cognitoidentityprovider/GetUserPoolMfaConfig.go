package cognitoidentityprovider

// GetUserPoolMfaConfig is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Given a user pool ID, returns configuration for sign-in with WebAuthn
// authenticators and for multi-factor authentication (MFA). This operation
// describes the following:
//
// - The WebAuthn relying party (RP) ID and user-verification settings.
//
// - The required, optional, or disabled state of MFA for all user pool users.
//
// - The message templates for email and SMS MFA.
//
// - The enabled or disabled state of time-based one-time password (TOTP) MFA.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
