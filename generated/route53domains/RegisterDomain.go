package route53domains

// RegisterDomain is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// This operation registers a domain. For some top-level domains (TLDs), this
// operation requires extra parameters.
//
// When you register a domain, Amazon Route 53 does the following:
//
// - Creates a Route 53 hosted zone that has the same name as the domain. Route
// 53 assigns four name servers to your hosted zone and automatically updates your
// domain registration with the names of these name servers.
//
// - Enables auto renew, so your domain registration will renew automatically
// each year. We'll notify you in advance of the renewal date so you can choose
// whether to renew the registration.
//
// - Optionally enables privacy protection, so WHOIS queries return contact for
// the registrar or the phrase "REDACTED FOR PRIVACY", or "On behalf of owner." If
// you don't enable privacy protection, WHOIS queries return the information that
// you entered for the administrative, registrant, and technical contacts.
//
// While some domains may allow different privacy settings per contact, we
//
// recommend specifying the same privacy setting for all contacts.
//
// - If registration is successful, returns an operation ID that you can use to
// track the progress and completion of the action. If the request is not completed
// successfully, the domain registrant is notified by email.
//
// - Charges your Amazon Web Services account an amount based on the top-level
// domain. For more information, see [Amazon Route 53 Pricing].
//
// [Amazon Route 53 Pricing]: http://aws.amazon.com/route53/pricing/
