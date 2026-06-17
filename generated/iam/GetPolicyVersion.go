package iam

// GetPolicyVersion is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves information about the specified version of the specified managed
// policy, including the policy document.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// To list the available versions for a policy, use [ListPolicyVersions].
//
// This operation retrieves information about managed policies. To retrieve
// information about an inline policy that is embedded in a user, group, or role,
// use [GetUserPolicy], [GetGroupPolicy], or [GetRolePolicy].
//
// For more information about the types of policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For more information about managed policy versions, see [Versioning for managed policies] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [GetRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html
// [GetGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html
// [GetUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
// [ListPolicyVersions]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPolicyVersions.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
