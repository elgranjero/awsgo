package cognitoidentityprovider

// AdminUpdateDeviceStatus is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Updates the status of a user's device so that it is marked as remembered or not
// remembered for the purpose of device authentication. Device authentication is a
// "remember me" mechanism that silently completes sign-in from trusted devices
// with a device key instead of a user-provided MFA code. This operation changes
// the status of a device without deleting it, so you can enable it again later.
// For more information about device authentication, see [Working with devices].
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
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
