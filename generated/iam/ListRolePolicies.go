package iam

// ListRolePolicies is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists the names of the inline policies that are embedded in the specified IAM
// role.
//
// An IAM role can also have managed policies attached to it. To list the managed
// policies that are attached to a role, use [ListAttachedRolePolicies]. For more information about
// policies, see [Managed policies and inline policies]in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If there
// are no inline policies embedded with the specified role, the operation returns
// an empty list.
//
// [ListAttachedRolePolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListAttachedRolePolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
