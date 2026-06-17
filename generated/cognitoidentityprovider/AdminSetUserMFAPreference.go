package cognitoidentityprovider

// AdminSetUserMFAPreference is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Sets the user's multi-factor authentication (MFA) preference, including which
// MFA options are activated, and if any are preferred. Only one factor can be set
// as preferred. The preferred MFA factor will be used to authenticate a user if
// multiple factors are activated. If multiple options are activated and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign-in.
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
