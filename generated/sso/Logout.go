package sso

// Logout is generated as a reference stub.
// Executable command wiring lives under cmd/sso.go.
//
// Removes the locally stored SSO tokens from the client-side cache and sends an
// API call to the IAM Identity Center service to invalidate the corresponding
// server-side IAM Identity Center sign in session.
//
// If a user uses IAM Identity Center to access the AWS CLI, the user’s IAM
// Identity Center sign in session is used to obtain an IAM session, as specified
// in the corresponding IAM Identity Center permission set. More specifically, IAM
// Identity Center assumes an IAM role in the target account on behalf of the user,
// and the corresponding temporary AWS credentials are returned to the client.
//
// After user logout, any existing IAM role sessions that were created by using
// IAM Identity Center permission sets continue based on the duration configured in
// the permission set. For more information, see [User authentications]in the IAM Identity Center User
// Guide.
//
// [User authentications]: https://docs.aws.amazon.com/singlesignon/latest/userguide/authconcept.html
