package iam

// PutUserPermissionsBoundary is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Adds or updates the policy that is specified as the IAM user's permissions
// boundary. You can use an Amazon Web Services managed policy or a customer
// managed policy to set the boundary for a user. Use the boundary to control the
// maximum permissions that the user can have. Setting a permissions boundary is an
// advanced feature that can affect the permissions for the user.
//
// Policies that are used as permissions boundaries do not provide permissions.
// You must also attach a permissions policy to the user. To learn how the
// effective permissions for a user are evaluated, see [IAM JSON policy evaluation logic]in the IAM User Guide.
//
// [IAM JSON policy evaluation logic]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html
