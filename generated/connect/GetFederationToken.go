package connect

// GetFederationToken is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Supports SAML sign-in for Amazon Connect. Retrieves a token for federation. The
// token is for the Amazon Connect user which corresponds to the IAM credentials
// that were used to invoke this action.
//
// For more information about how SAML sign-in works in Amazon Connect, see [Configure SAML with IAM for Amazon Connect in the Amazon Connect Administrator Guide.]
//
// This API doesn't support root users. If you try to invoke GetFederationToken
// with root credentials, an error message similar to the following one appears:
//
// Provided identity: Principal: .... User: .... cannot be used for federation
// with Amazon Connect
//
// [Configure SAML with IAM for Amazon Connect in the Amazon Connect Administrator Guide.]: https://docs.aws.amazon.com/connect/latest/adminguide/configure-saml.html
