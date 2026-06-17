package cognitoidentityprovider

// DeleteUserPool is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Deletes a user pool. After you delete a user pool, users can no longer sign in
// to any associated applications.
//
// When you delete a user pool, it's no longer visible or operational in your
// Amazon Web Services account. Amazon Cognito retains deleted user pools in an
// inactive state for 14 days, then begins a cleanup process that fully removes
// them from Amazon Web Services systems. In case of accidental deletion, contact
// Amazon Web Services Support within 14 days for restoration assistance.
//
// Amazon Cognito begins full deletion of all resources from deleted user pools
// after 14 days. In the case of large user pools, the cleanup process might take
// significant additional time before all user data is permanently deleted.
