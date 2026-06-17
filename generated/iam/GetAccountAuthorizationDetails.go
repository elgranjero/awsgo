package iam

// GetAccountAuthorizationDetails is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves information about all IAM users, groups, roles, and policies in your
// Amazon Web Services account, including their relationships to one another. Use
// this operation to obtain a snapshot of the configuration of IAM permissions
// (users, groups, roles, and policies) in your account.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// You can optionally filter the results using the Filter parameter. You can
// paginate the results using the MaxItems and Marker parameters.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
