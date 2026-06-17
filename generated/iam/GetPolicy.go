package iam

// GetPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves information about the specified managed policy, including the
// policy's default version and the total number of IAM users, groups, and roles to
// which the policy is attached. To retrieve the list of the specific users,
// groups, and roles that the policy is attached to, use [ListEntitiesForPolicy]. This operation returns
// metadata about the policy. To retrieve the actual policy document for a specific
// version of the policy, use [GetPolicyVersion].
//
// This operation retrieves information about managed policies. To retrieve
// information about an inline policy that is embedded with an IAM user, group, or
// role, use [GetUserPolicy], [GetGroupPolicy], or [GetRolePolicy].
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [ListEntitiesForPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListEntitiesForPolicy.html
// [GetRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html
// [GetPolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicyVersion.html
// [GetGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html
// [GetUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
