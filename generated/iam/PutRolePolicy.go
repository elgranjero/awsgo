package iam

// PutRolePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Adds or updates an inline policy document that is embedded in the specified IAM
// role.
//
// When you embed an inline policy in a role, the inline policy is used as part of
// the role's access (permissions) policy. The role's trust policy is created at
// the same time as the role, using [CreateRole]CreateRole . You can update a role's trust
// policy using [UpdateAssumeRolePolicy]UpdateAssumeRolePolicy . For more information about roles, see [IAM roles] in
// the IAM User Guide.
//
// A role can also have a managed policy attached to it. To attach a managed
// policy to a role, use [AttachRolePolicy]AttachRolePolicy . To create a new managed policy, use [CreatePolicy]
// CreatePolicy . For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For information about the maximum number of inline policies that you can embed
// with a role, see [IAM and STS quotas]in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET when
// calling PutRolePolicy . For general information about using the Query API with
// IAM, see [Making query requests]in the IAM User Guide.
//
// [UpdateAssumeRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html
// [AttachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachRolePolicy.html
// [CreatePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html
// [CreateRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
