package ses

// GetIdentityDkimAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Returns the current status of Easy DKIM signing for an entity. For domain name
// identities, this operation also returns the DKIM tokens that are required for
// Easy DKIM signing, and whether Amazon SES has successfully verified that these
// tokens have been published.
//
// This operation takes a list of identities as input and returns the following
// information for each:
//
// - Whether Easy DKIM signing is enabled or disabled.
//
// - A set of DKIM tokens that represent the identity. If the identity is an
// email address, the tokens represent the domain of that address.
//
// - Whether Amazon SES has successfully verified the DKIM tokens published in
// the domain's DNS. This information is only returned for domain name identities,
// not for email addresses.
//
// This operation is throttled at one request per second and can only get DKIM
// attributes for up to 100 identities at a time.
//
// For more information about creating DNS records using DKIM tokens, go to the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy-managing.html
