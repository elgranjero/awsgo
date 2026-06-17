package route53domains

// DeleteDomain is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// This operation deletes the specified domain. This action is permanent. For more
// information, see [Deleting a domain name registration].
//
// To transfer the domain registration to another registrar, use the transfer
// process that’s provided by the registrar to which you want to transfer the
// registration. Otherwise, the following apply:
//
// - You can’t get a refund for the cost of a deleted domain registration.
//
// - The registry for the top-level domain might hold the domain name for a
// brief time before releasing it for other users to register (varies by registry).
//
// - When the registration has been deleted, we'll send you a confirmation to
// the registrant contact. The email will come from
// noreply(at)domainnameverification.net or noreply(at)registrar.amazon.com .
//
// [Deleting a domain name registration]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-delete.html
