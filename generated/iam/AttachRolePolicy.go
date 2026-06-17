package iam

// AttachRolePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Attaches the specified managed policy to the specified IAM role. When you
// attach a managed policy to a role, the managed policy becomes part of the role's
// permission (access) policy.
//
// You cannot use a managed policy as the role's trust policy. The role's trust
// policy is created at the same time as the role, using [CreateRole]CreateRole . You can
// update a role's trust policy using [UpdateAssumerolePolicy]UpdateAssumerolePolicy .
//
// Use this operation to attach a managed policy to a role. To embed an inline
// policy in a role, use [PutRolePolicy]PutRolePolicy . For more information about policies, see [Managed policies and inline policies]
// in the IAM User Guide.
//
// As a best practice, you can validate your IAM policies. To learn more, see [Validating IAM policies] in
// the IAM User Guide.
//
// [Validating IAM policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html
// [UpdateAssumerolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html
// [PutRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutRolePolicy.html
// [CreateRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
