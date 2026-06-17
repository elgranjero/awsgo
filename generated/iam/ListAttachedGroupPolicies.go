package iam

// ListAttachedGroupPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists all managed policies that are attached to the specified IAM group.
//
// An IAM group can also have inline policies embedded with it. To list the inline
// policies for a group, use [ListGroupPolicies]. For information about policies, see [Managed policies and inline policies] in the IAM
// User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You can
// use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to the
// specified group (or none that match the specified path prefix), the operation
// returns an empty list.
//
// [ListGroupPolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListGroupPolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
