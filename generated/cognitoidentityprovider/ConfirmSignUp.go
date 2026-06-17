package cognitoidentityprovider

// ConfirmSignUp is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Confirms the account of a new user. This public API operation submits a code
// that Amazon Cognito sent to your user when they signed up in your user pool.
// After your user enters their code, they confirm ownership of the email address
// or phone number that they provided, and their user account becomes active.
// Depending on your user pool configuration, your users will receive their
// confirmation code in an email or SMS message.
//
// Local users who signed up in your user pool are the only type of user who can
// confirm sign-up with a code. Users who federate through an external identity
// provider (IdP) have already been confirmed by their IdP.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
