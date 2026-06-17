package iam

// DeletePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the specified managed policy.
//
// Before you can delete a managed policy, you must first detach the policy from
// all users, groups, and roles that it is attached to. In addition, you must
// delete all the policy's versions. The following steps describe the process for
// deleting a managed policy:
//
// - Detach the policy from all users, groups, and roles that the policy is
// attached to, using [DetachUserPolicy], [DetachGroupPolicy], or [DetachRolePolicy]. To list all the users, groups, and roles that a
// policy is attached to, use [ListEntitiesForPolicy].
//
// - Delete all versions of the policy using [DeletePolicyVersion]. To list the policy's versions,
// use [ListPolicyVersions]. You cannot use [DeletePolicyVersion]to delete the version that is marked as the default
// version. You delete the policy's default version in the next step of the
// process.
//
// - Delete the policy (this automatically deletes the policy's default version)
// using this operation.
//
// For information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [DetachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachUserPolicy.html
// [DetachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachRolePolicy.html
// [ListEntitiesForPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListEntitiesForPolicy.html
// [DeletePolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeletePolicyVersion.html
// [DetachGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachGroupPolicy.html
// [ListPolicyVersions]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPolicyVersions.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
