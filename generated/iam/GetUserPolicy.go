package iam

// GetUserPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves the specified inline policy document that is embedded in the
// specified IAM user.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// An IAM user can also have managed policies attached to it. To retrieve a
// managed policy document that is attached to a user, use [GetPolicy]to determine the
// policy's default version. Then use [GetPolicyVersion]to retrieve the policy document.
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [GetPolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicyVersion.html
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
