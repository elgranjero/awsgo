package route53domains

// TransferDomain is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// Transfers a domain from another registrar to Amazon Route 53.
//
// For more information about transferring domains, see the following topics:
//
// - For transfer requirements, a detailed procedure, and information about
// viewing the status of a domain that you're transferring to Route 53, see [Transferring Registration for a Domain to Amazon Route 53]in
// the Amazon Route 53 Developer Guide.
//
// - For information about how to transfer a domain from one Amazon Web Services
// account to another, see [TransferDomainToAnotherAwsAccount].
//
// - For information about how to transfer a domain to another domain registrar,
// see [Transferring a Domain from Amazon Route 53 to Another Registrar]in the Amazon Route 53 Developer Guide.
//
// During the transfer of any country code top-level domains (ccTLDs) to Route 53,
// except for .cc and .tv, updates to the owner contact are ignored and the owner
// contact data from the registry is used. You can update the owner contact after
// the transfer is complete. For more information, see [UpdateDomainContact].
//
// If the registrar for your domain is also the DNS service provider for the
// domain, we highly recommend that you transfer your DNS service to Route 53 or to
// another DNS service provider before you transfer your registration. Some
// registrars provide free DNS service when you purchase a domain registration.
// When you transfer the registration, the previous registrar will not renew your
// domain registration and could end your DNS service at any time.
//
// If the registrar for your domain is also the DNS service provider for the
// domain and you don't transfer DNS service to another provider, your website,
// email, and the web applications associated with the domain might become
// unavailable.
//
// If the transfer is successful, this method returns an operation ID that you can
// use to track the progress and completion of the action. If the transfer doesn't
// complete successfully, the domain registrant will be notified by email.
//
// [Transferring a Domain from Amazon Route 53 to Another Registrar]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-from-route-53.html
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [Transferring Registration for a Domain to Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html
// [UpdateDomainContact]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainContact.html
