package route53domains

// AssociateDelegationSignerToDomain is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// Creates a delegation signer (DS) record in the registry zone for this domain
//
// name.
//
// Note that creating DS record at the registry impacts DNSSEC validation of your
// DNS records. This action may render your domain name unavailable on the internet
// if the steps are completed in the wrong order, or with incorrect timing. For
// more information about DNSSEC signing, see [Configuring DNSSEC signing]in the Route 53 developer guide.
//
// [Configuring DNSSEC signing]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html
