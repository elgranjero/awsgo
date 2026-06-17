package route53

// DeleteHostedZone is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Deletes a hosted zone.
//
// If the hosted zone was created by another service, such as Cloud Map, see [Deleting Public Hosted Zones That Were Created by Another Service] in
// the Amazon Route 53 Developer Guide for information about how to delete it. (The
// process is the same for public and private hosted zones that were created by
// another service.)
//
// If you want to keep your domain registration but you want to stop routing
// internet traffic to your website or web application, we recommend that you
// delete resource record sets in the hosted zone instead of deleting the hosted
// zone.
//
// If you delete a hosted zone, you can't undelete it. You must create a new
// hosted zone and update the name servers for your domain registration, which can
// require up to 48 hours to take effect. (If you delegated responsibility for a
// subdomain to a hosted zone and you delete the child hosted zone, you must update
// the name servers in the parent hosted zone.) In addition, if you delete a hosted
// zone, someone could hijack the domain and route traffic to their own resources
// using your domain name.
//
// If you want to avoid the monthly charge for the hosted zone, you can transfer
// DNS service for the domain to a free DNS service. When you transfer DNS service,
// you have to update the name servers for the domain registration. If the domain
// is registered with Route 53, see [UpdateDomainNameservers]for information about how to replace Route 53
// name servers with name servers for the new DNS service. If the domain is
// registered with another registrar, use the method provided by the registrar to
// update name servers for the domain registration. For more information, perform
// an internet search on "free DNS service."
//
// You can delete a hosted zone only if it contains only the default SOA and NS
// records and has DNSSEC signing disabled. If the hosted zone contains other
// records or has DNSSEC enabled, you must delete the records and disable DNSSEC
// before deletion. Attempting to delete a hosted zone with additional records or
// DNSSEC enabled returns a HostedZoneNotEmpty error. For information about
// deleting records, see [ChangeResourceRecordSets].
//
// To verify that the hosted zone has been deleted, do one of the following:
//
// - Use the GetHostedZone action to request information about the hosted zone.
//
// - Use the ListHostedZones action to get a list of the hosted zones associated
// with the current Amazon Web Services account.
//
// [ChangeResourceRecordSets]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html
// [Deleting Public Hosted Zones That Were Created by Another Service]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DeleteHostedZone.html#delete-public-hosted-zone-created-by-another-service
// [UpdateDomainNameservers]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainNameservers.html
