package organizations

// CreateOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Creates an Amazon Web Services organization. The account whose user is calling
// the CreateOrganization operation automatically becomes the [management account] of the new
// organization.
//
// This operation must be called using credentials from the account that is to
// become the new organization's management account. The principal must also have
// the relevant IAM permissions.
//
// By default (or if you set the FeatureSet parameter to ALL ), the new
// organization is created with all features enabled and service control policies
// automatically enabled in the root. If you instead choose to create the
// organization supporting only the consolidated billing features by setting the
// FeatureSet parameter to CONSOLIDATED_BILLING , no policy types are enabled by
// default and you can't use organization policies.
//
// [management account]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account
