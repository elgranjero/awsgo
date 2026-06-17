package ses

// SetIdentityMailFromDomain is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Enables or disables the custom MAIL FROM domain setup for a verified identity
// (an email address or a domain).
//
// To send emails using the specified MAIL FROM domain, you must add an MX record
// to your MAIL FROM domain's DNS settings. To ensure that your emails pass Sender
// Policy Framework (SPF) checks, you must also add or update an SPF record. For
// more information, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/mail-from.html
