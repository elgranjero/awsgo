package cognitoidentityprovider

// AdminConfirmSignUp is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Confirms user sign-up as an administrator.
//
// This request sets a user account active in a user pool that [requires confirmation of new user accounts] before they can
// sign in. You can configure your user pool to not send confirmation codes to new
// users and instead confirm them with this API operation on the back end.
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
// To configure your user pool to require administrative confirmation of users,
// set AllowAdminCreateUserOnly to true in a CreateUserPool or UpdateUserPool
// request.
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [requires confirmation of new user accounts]: https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#signing-up-users-in-your-app-and-confirming-them-as-admin
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
