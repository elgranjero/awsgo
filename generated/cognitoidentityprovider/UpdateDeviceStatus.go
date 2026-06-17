package cognitoidentityprovider

// UpdateDeviceStatus is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Updates the status of a the currently signed-in user's device so that it is
// marked as remembered or not remembered for the purpose of device authentication.
// Device authentication is a "remember me" mechanism that silently completes
// sign-in from trusted devices with a device key instead of a user-provided MFA
// code. This operation changes the status of a device without deleting it, so you
// can enable it again later. For more information about device authentication, see
// [Working with devices].
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
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
