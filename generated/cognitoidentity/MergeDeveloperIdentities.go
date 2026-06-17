package cognitoidentity

// MergeDeveloperIdentities is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentity.go.
//
// Merges two users having different IdentityId s, existing in the same identity
// pool, and identified by the same developer provider. You can use this action to
// request that discrete users be merged and identified as a single user in the
// Cognito environment. Cognito associates the given source user (
// SourceUserIdentifier ) with the IdentityId of the DestinationUserIdentifier .
// Only developer-authenticated users can be merged. If the users to be merged are
// associated with the same public provider, but as two different users, an
// exception will be thrown.
//
// The number of linked logins is limited to 20. So, the number of linked logins
// for the source user, SourceUserIdentifier , and the destination user,
// DestinationUserIdentifier , together should not be larger than 20. Otherwise, an
// exception will be thrown.
//
// You must use Amazon Web Services developer credentials to call this operation.
