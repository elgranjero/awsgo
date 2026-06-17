package organizations

// EnablePolicyType is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Enables a policy type in a root. After you enable a policy type in a root, you
// can attach policies of that type to the root, any organizational unit (OU), or
// account in that root. You can undo this by using the DisablePolicyTypeoperation.
//
// This is an asynchronous request that Amazon Web Services performs in the
// background. Amazon Web Services recommends that you first use ListRootsto see the status
// of policy types for a specified root, and then use this operation.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// You can enable a policy type in a root only if that policy type is available in
// the organization. To view the status of available policy types in the
// organization, use ListRoots.
