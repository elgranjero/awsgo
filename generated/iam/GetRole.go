package iam

// GetRole is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves information about the specified role, including the role's path,
// GUID, ARN, and the role's trust policy that grants permission to assume the
// role. For more information about roles, see [IAM roles]in the IAM User Guide.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
