package iam

// ListPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists all the managed policies that are available in your Amazon Web Services
// account, including your own customer-defined managed policies and all Amazon Web
// Services managed policies.
//
// You can filter the list of policies that is returned using the optional
// OnlyAttached , Scope , and PathPrefix parameters. For example, to list only the
// customer managed policies in your Amazon Web Services account, set Scope to
// Local . To list only Amazon Web Services managed policies, set Scope to AWS .
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for a
// customer manged policy, see [GetPolicy].
//
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
