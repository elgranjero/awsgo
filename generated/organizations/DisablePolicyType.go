package organizations

// DisablePolicyType is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Disables an organizational policy type in a root. A policy of a certain type
// can be attached to entities in a root only if that type is enabled in the root.
// After you perform this operation, you no longer can attach policies of the
// specified type to that root or to any organizational unit (OU) or account in
// that root. You can undo this by using the EnablePolicyTypeoperation.
//
// This is an asynchronous request that Amazon Web Services performs in the
// background. If you disable a policy type for a root, it still appears enabled
// for the organization if [all features]are enabled for the organization. Amazon Web Services
// recommends that you first use ListRootsto see the status of policy types for a specified
// root, and then use this operation.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// To view the status of available policy types in the organization, use ListRoots.
//
// [all features]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
