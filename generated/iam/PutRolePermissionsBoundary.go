package iam

// PutRolePermissionsBoundary is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Adds or updates the policy that is specified as the IAM role's permissions
// boundary. You can use an Amazon Web Services managed policy or a customer
// managed policy to set the boundary for a role. Use the boundary to control the
// maximum permissions that the role can have. Setting a permissions boundary is an
// advanced feature that can affect the permissions for the role.
//
// You cannot set the boundary for a service-linked role.
//
// Policies used as permissions boundaries do not provide permissions. You must
// also attach a permissions policy to the role. To learn how the effective
// permissions for a role are evaluated, see [IAM JSON policy evaluation logic]in the IAM User Guide.
//
// [IAM JSON policy evaluation logic]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html
