package cognitoidentityprovider

// VerifySoftwareToken is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Registers the current user's time-based one-time password (TOTP) authenticator
// with a code generated in their authenticator app from a private key that's
// supplied by your user pool. Marks the user's software token MFA status as
// "verified" if successful. The request takes an access token or a session string,
// but not both.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
