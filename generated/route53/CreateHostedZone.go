package route53

// CreateHostedZone is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Creates a new public or private hosted zone. You create records in a public
// hosted zone to define how you want to route traffic on the internet for a
// domain, such as example.com, and its subdomains (apex.example.com,
// acme.example.com). You create records in a private hosted zone to define how you
// want to route traffic for a domain and its subdomains within one or more Amazon
// Virtual Private Clouds (Amazon VPCs).
//
// You can't convert a public hosted zone to a private hosted zone or vice versa.
// Instead, you must create a new hosted zone with the same name and create new
// resource record sets.
//
// For more information about charges for hosted zones, see [Amazon Route 53 Pricing].
//
// Note the following:
//
// - You can't create a hosted zone for a top-level domain (TLD) such as .com.
//
// - For public hosted zones, Route 53 automatically creates a default SOA
// record and four NS records for the zone. For more information about SOA and NS
// records, see [NS and SOA Records that Route 53 Creates for a Hosted Zone]in the Amazon Route 53 Developer Guide.
//
// If you want to use the same name servers for multiple public hosted zones, you
//
// can optionally associate a reusable delegation set with the hosted zone. See the
// DelegationSetId element.
//
// - If your domain is registered with a registrar other than Route 53, you must
// update the name servers with your registrar to make Route 53 the DNS service for
// the domain. For more information, see [Migrating DNS Service for an Existing Domain to Amazon Route 53]in the Amazon Route 53 Developer Guide.
//
// When you submit a CreateHostedZone request, the initial status of the hosted
// zone is PENDING . For public hosted zones, this means that the NS and SOA
// records are not yet available on all Route 53 DNS servers. When the NS and SOA
// records are available, the status of the zone changes to INSYNC .
//
// The CreateHostedZone request requires the caller to have an ec2:DescribeVpcs
// permission.
//
// When creating private hosted zones, the Amazon VPC must belong to the same
// partition where the hosted zone is created. A partition is a group of Amazon Web
// Services Regions. Each Amazon Web Services account is scoped to one partition.
//
// The following are the supported partitions:
//
// - aws - Amazon Web Services Regions
//
// - aws-cn - China Regions
//
// - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see [Access Management] in the Amazon Web Services General Reference.
//
// [Access Management]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [NS and SOA Records that Route 53 Creates for a Hosted Zone]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/SOA-NSrecords.html
// [Amazon Route 53 Pricing]: http://aws.amazon.com/route53/pricing/
//
// [Migrating DNS Service for an Existing Domain to Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html
