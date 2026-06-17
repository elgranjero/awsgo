package acm

// RequestCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acm.go.
//
// Requests an ACM certificate for use with other Amazon Web Services services. To
// request an ACM certificate, you must specify a fully qualified domain name
// (FQDN) in the DomainName parameter. You can also specify additional FQDNs in
// the SubjectAlternativeNames parameter.
//
// If you are requesting a private certificate, domain validation is not required.
// If you are requesting a public certificate, each domain name that you specify
// must be validated to verify that you own or control the domain. You can use [DNS validation]or [email validation]
// . We recommend that you use DNS validation.
//
// ACM behavior differs from the [RFC 6125] specification of the certificate validation
// process. ACM first checks for a Subject Alternative Name, and, if it finds one,
// ignores the common name (CN).
//
// After successful completion of the RequestCertificate action, there is a delay
// of several seconds before you can retrieve information about the new
// certificate.
//
// [email validation]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html
// [RFC 6125]: https://datatracker.ietf.org/doc/html/rfc6125#appendix-B.2
// [DNS validation]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html
