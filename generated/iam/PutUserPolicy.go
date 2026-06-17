package iam

// PutUserPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Adds or updates an inline policy document that is embedded in the specified IAM
// user.
//
// An IAM user can also have a managed policy attached to it. To attach a managed
// policy to a user, use [AttachUserPolicy]AttachUserPolicy . To create a new managed policy, use [CreatePolicy]
// CreatePolicy . For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For information about the maximum number of inline policies that you can embed
// in a user, see [IAM and STS quotas]in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET when
// calling PutUserPolicy . For general information about using the Query API with
// IAM, see [Making query requests]in the IAM User Guide.
//
// [CreatePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [AttachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
