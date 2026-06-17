package ses

// DeleteIdentityPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Deletes the specified sending authorization policy for the given identity (an
// email address or a domain). This operation returns successfully even if a policy
// with the specified name does not exist.
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
