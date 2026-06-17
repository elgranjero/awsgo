package sesv2

// GetEmailIdentityPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/sesv2.go.
//
// Returns the requested sending authorization policies for the given identity (an
// email address or a domain). The policies are returned as a map of policy names
// to policy contents. You can retrieve a maximum of 20 policies at a time.
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html
