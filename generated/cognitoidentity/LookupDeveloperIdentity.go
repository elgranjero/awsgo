package cognitoidentity

// LookupDeveloperIdentity is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentity.go.
//
// Retrieves the IdentityID associated with a DeveloperUserIdentifier or the list
// of DeveloperUserIdentifier values associated with an IdentityId for an existing
// identity. Either IdentityID or DeveloperUserIdentifier must not be null. If you
// supply only one of these values, the other value will be searched in the
// database and returned as a part of the response. If you supply both,
// DeveloperUserIdentifier will be matched against IdentityID . If the values are
// verified against the database, the response returns both values and is the same
// as the request. Otherwise, a ResourceConflictException is thrown.
//
// LookupDeveloperIdentity is intended for low-throughput control plane
// operations: for example, to enable customer service to locate an identity ID by
// username. If you are using it for higher-volume operations such as user
// authentication, your requests are likely to be throttled. GetOpenIdTokenForDeveloperIdentityis a better option
// for higher-volume operations for user authentication.
//
// You must use Amazon Web Services developer credentials to call this operation.
