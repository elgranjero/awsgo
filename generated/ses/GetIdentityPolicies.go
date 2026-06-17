package ses

// GetIdentityPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Returns the requested sending authorization policies for the given identity (an
// email address or a domain). The policies are returned as a map of policy names
// to policy contents. You can retrieve a maximum of 20 policies at a time.
//
// This operation is for the identity owner only. If you have not verified the
// identity, it returns an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
