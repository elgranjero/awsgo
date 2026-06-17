package cognitoidentityprovider

// ConfirmDevice is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Confirms a device that a user wants to remember. A remembered device is a
// "Remember me on this device" option for user pools that perform authentication
// with the device key of a trusted device in the back end, instead of a
// user-provided MFA code. For more information about device authentication, see [Working with user devices in your user pool].
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
// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
