package cognitoidentityprovider

// AdminDisableProviderForUser is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Prevents the user from signing in with the specified external (SAML or social)
// identity provider (IdP). If the user that you want to deactivate is a Amazon
// Cognito user pools native username + password user, they can't use their
// password to sign in. If the user to deactivate is a linked external IdP user,
// any link between that user and an existing user is removed. When the external
// user signs in again, and the user is no longer attached to the previously linked
// DestinationUser , the user must create a new user account.
//
// The value of ProviderName must match the name of a user pool IdP.
//
// To deactivate a local user, set ProviderName to Cognito and the
// ProviderAttributeName to Cognito_Subject . The ProviderAttributeValue must be
// user's local username.
//
// The ProviderAttributeName must always be Cognito_Subject for social IdPs. The
// ProviderAttributeValue must always be the exact subject that was used when the
// user was originally linked as a source user.
//
// For de-linking a SAML identity, there are two scenarios. If the linked identity
// has not yet been used to sign in, the ProviderAttributeName and
// ProviderAttributeValue must be the same values that were used for the SourceUser
// when the identities were originally linked using AdminLinkProviderForUser call.
// This is also true if the linking was done with ProviderAttributeName set to
// Cognito_Subject . If the user has already signed in, the ProviderAttributeName
// must be Cognito_Subject and ProviderAttributeValue must be the NameID from
// their SAML assertion.
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
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
