package cognitoidentityprovider

// AdminSetUserPassword is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Sets the specified user's password in a user pool. This operation
// administratively sets a temporary or permanent password for a user. With this
// operation, you can bypass self-service password changes and permit immediate
// sign-in with the password that you set. To do this, set Permanent to true .
//
// You can also set a new temporary password in this request, send it to a user,
// and require them to choose a new password on their next sign-in. To do this, set
// Permanent to false .
//
// If the password is temporary, the user's Status becomes FORCE_CHANGE_PASSWORD .
// When the user next tries to sign in, the InitiateAuth or AdminInitiateAuth
// response includes the NEW_PASSWORD_REQUIRED challenge. If the user doesn't sign
// in before the temporary password expires, they can no longer sign in and you
// must repeat this operation to set a temporary or permanent password for them.
//
// After the user sets a new password, or if you set a permanent password, their
// status becomes Confirmed .
//
// AdminSetUserPassword can set a password for the user profile that Amazon
// Cognito creates for third-party federated users. When you set a password, the
// federated user's status changes from EXTERNAL_PROVIDER to CONFIRMED . A user in
// this state can sign in as a federated user, and initiate authentication flows in
// the API like a linked native user. They can also modify their password and
// attributes in token-authenticated API requests like ChangePassword and
// UpdateUserAttributes . As a best security practice and to keep users in sync
// with your external IdP, don't set passwords on federated user profiles. To set
// up a federated user for native sign-in with a linked native user, refer to [Linking federated users to an existing user profile].
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
// [Linking federated users to an existing user profile]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
