package cognitoidentityprovider

// AssociateSoftwareToken is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Begins setup of time-based one-time password (TOTP) multi-factor authentication
// (MFA) for a user, with a unique private key that Amazon Cognito generates and
// returns in the API response. You can authorize an AssociateSoftwareToken
// request with either the user's access token, or a session string from a
// challenge response that you received from Amazon Cognito.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
