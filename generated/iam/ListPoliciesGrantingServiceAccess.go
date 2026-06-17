package iam

// ListPoliciesGrantingServiceAccess is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves a list of policies that the IAM identity (user, group, or role) can
// use to access each specified service.
//
// This operation does not use other policy types when determining whether a
// resource could access a service. These other policy types include resource-based
// policies, access control lists, Organizations policies, IAM permissions
// boundaries, and STS assume role policies. It only applies permissions policy
// logic. For more about the evaluation of policy types, see [Evaluating policies]in the IAM User Guide.
//
// The list of policies returned by the operation depends on the ARN of the
// identity that you provide.
//
// - User – The list of policies includes the managed and inline policies that
// are attached to the user directly. The list also includes any additional managed
// and inline policies that are attached to the group to which the user belongs.
//
// - Group – The list of policies includes only the managed and inline policies
// that are attached to the group directly. Policies that are attached to the
// group’s user are not included.
//
// - Role – The list of policies includes only the managed and inline policies
// that are attached to the role.
//
// For each managed policy, this operation returns the ARN and policy name. For
// each inline policy, it returns the policy name and the entity to which it is
// attached. Inline policies do not have an ARN. For more information about these
// policy types, see [Managed policies and inline policies]in the IAM User Guide.
//
// Policies that are attached to users and roles as permissions boundaries are not
// returned. To view which managed policy is currently used to set the permissions
// boundary for a user or role, use the [GetUser]or [GetRole] operations.
//
// [GetRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRole.html
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html
// [GetUser]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUser.html
