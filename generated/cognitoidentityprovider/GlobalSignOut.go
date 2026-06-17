package cognitoidentityprovider

// GlobalSignOut is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Invalidates the identity, access, and refresh tokens that Amazon Cognito issued
// to a user. Call this operation when your user signs out of your app. This
// results in the following behavior.
//
// - Amazon Cognito no longer accepts token-authorized user operations that you
// authorize with a signed-out user's access tokens. For more information, see [Using the Amazon Cognito user pools API and user pool endpoints].
//
// Amazon Cognito returns an Access Token has been revoked error when your app
//
// attempts to authorize a user pools API request with a revoked access token that
// contains the scope aws.cognito.signin.user.admin .
//
// - Amazon Cognito no longer accepts a signed-out user's ID token in a [GetId]request
// to an identity pool with ServerSideTokenCheck enabled for its user pool IdP
// configuration in [CognitoIdentityProvider].
//
// - Amazon Cognito no longer accepts a signed-out user's refresh tokens in
// refresh requests.
//
// Other requests might be valid until your user's token expires. This operation
// doesn't clear the [managed login]session cookie. To clear the session for a user who signed in
// with managed login or the classic hosted UI, direct their browser session to the
// [logout endpoint].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [logout endpoint]: https://docs.aws.amazon.com/cognito/latest/developerguide/logout-endpoint.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [managed login]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html
//
// [CognitoIdentityProvider]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CognitoIdentityProvider.html
// [GetId]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html
