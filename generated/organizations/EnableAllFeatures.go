package organizations

// EnableAllFeatures is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Enables all features in an organization. This enables the use of organization
// policies that can restrict the services and actions that can be called in each
// account. Until you enable all features, you have access only to consolidated
// billing, and you can't use any of the advanced account administration features
// that Organizations supports. For more information, see [Enabling all features in your organization]in the Organizations
// User Guide.
//
// This operation is required only for organizations that were created explicitly
// with only the consolidated billing features enabled. Calling this operation
// sends a handshake to every invited account in the organization. The feature set
// change can be finalized and the additional features enabled only after all
// administrators in the invited accounts approve the change by accepting the
// handshake.
//
// After you enable all features, you can separately enable or disable individual
// policy types in a root using EnablePolicyTypeand DisablePolicyType. To see the status of policy types in a root,
// use ListRoots.
//
// After all invited member accounts accept the handshake, you finalize the
// feature set change by accepting the handshake that contains "Action":
// "ENABLE_ALL_FEATURES" . This completes the change.
//
// After you enable all features in your organization, the management account in
// the organization can apply policies on all member accounts. These policies can
// restrict what users and even administrators in those accounts can do. The
// management account can apply policies that prevent accounts from leaving the
// organization. Ensure that your account administrators are aware of this.
//
// You can only call this operation from the management account.
//
// [Enabling all features in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
