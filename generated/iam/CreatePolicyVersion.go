package iam

// CreatePolicyVersion is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Creates a new version of the specified managed policy. To update a managed
// policy, you create a new policy version. A managed policy can have up to five
// versions. If the policy has five versions, you must delete an existing version
// using [DeletePolicyVersion]before you create a new version.
//
// Optionally, you can set the new version as the policy's default version. The
// default version is the version that is in effect for the IAM users, groups, and
// roles to which the policy is attached.
//
// For more information about managed policy versions, see [Versioning for managed policies] in the IAM User Guide.
//
// [DeletePolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeletePolicyVersion.html
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
