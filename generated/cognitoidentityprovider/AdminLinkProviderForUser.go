package cognitoidentityprovider

// AdminLinkProviderForUser is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Links an existing user account in a user pool, or DestinationUser , to an
// identity from an external IdP, or SourceUser , based on a specified attribute
// name and value from the external IdP.
//
// This operation connects a local user profile with a user identity who hasn't
// yet signed in from their third-party IdP. When the user signs in with their IdP,
// they get access-control configuration from the local user profile. Linked local
// users can also sign in with SDK-based API operations like InitiateAuth after
// they sign in at least once through their IdP. For more information, see [Linking federated users].
//
// The maximum number of federated identities linked to a user is five.
//
// Because this API allows a user with an external federated identity to sign in
// as a local user, it is critical that it only be used with external IdPs and
// linked attributes that you trust.
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
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Linking federated users]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
