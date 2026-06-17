package cognitoidentityprovider

// DeleteUserAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Deletes attributes from the currently signed-in user. For example, your
// application can submit a request to this operation when a user wants to remove
// their birthdate attribute value.
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
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
