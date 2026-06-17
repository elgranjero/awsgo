package cognitoidentityprovider

// VerifyUserAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Submits a verification code for a signed-in user who has added or changed a
// value of an auto-verified attribute. When successful, the user's attribute
// becomes verified and the attribute email_verified or phone_number_verified
// becomes true .
//
// If your user pool requires verification before Amazon Cognito updates the
// attribute value, this operation updates the affected attribute to its pending
// value.
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
