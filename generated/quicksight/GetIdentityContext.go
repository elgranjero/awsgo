package quicksight

// GetIdentityContext is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Retrieves the identity context for a Quick Sight user in a specified namespace,
// allowing you to obtain identity tokens that can be used with identity-enhanced
// IAM role sessions to call identity-aware APIs.
//
// # Currently, you can call the following APIs with identity-enhanced Credentials
//
// [StartDashboardSnapshotJob]
//
// [DescribeDashboardSnapshotJob]
//
// [DescribeDashboardSnapshotJobResult]
//
// # Supported Authentication Methods
//
// This API supports Quick Sight native users, IAM federated users, and Active
// Directory users. For Quick Sight users authenticated by Amazon Web Services
// Identity Center, see [Identity Center documentation on identity-enhanced IAM role sessions].
//
// # Supported Regions
//
// The GetIdentityContext API works only in regions that support at least one of
// these identity types:
//
// - Amazon Quick Sight native identity
//
// - IAM federated identity
//
// - Active Directory
//
// To use this API successfully, call it in the same region where your user's
// identity resides. For example, if your user's identity is in us-east-1, make the
// API call in us-east-1. For more information about managing identities in Amazon
// Quick Sight, see [Identity and access management in Amazon Quick Sight]in the Amazon Quick Sight User Guide.
//
// # Getting Identity-Enhanced Credentials
//
// To obtain identity-enhanced credentials, follow these steps:
//
// - Call the GetIdentityContext API to retrieve an identity token for the
// specified user.
//
// - Use the identity token with the [STS AssumeRole API]to obtain identity-enhanced IAM role
// session credentials.
//
// # Usage with STS AssumeRole
//
// The identity token returned by this API should be used with the STS AssumeRole
// API to obtain credentials for an identity-enhanced IAM role session. When
// calling AssumeRole, include the identity token in the ProvidedContexts
// parameter with ProviderArn set to arn:aws:iam::aws:contextProvider/QuickSight
// and ContextAssertion set to the identity token received from this API.
//
// The assumed role must allow the sts:SetContext action in addition to
// sts:AssumeRole in its trust relationship policy. The trust policy should include
// both actions for the principal that will be assuming the role.
//
// [DescribeDashboardSnapshotJobResult]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeDashboardSnapshotJobResult.html
// [Identity Center documentation on identity-enhanced IAM role sessions]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-identity-enhanced-iam-role-sessions.html
// [DescribeDashboardSnapshotJob]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeDashboardSnapshotJob.html
// [STS AssumeRole API]: https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html
// [Identity and access management in Amazon Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/userguide/identity.html
// [StartDashboardSnapshotJob]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StartDashboardSnapshotJob.html
