package route53

// ListHostedZonesByVPC is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Lists all the private hosted zones that a specified VPC is associated with,
// regardless of which Amazon Web Services account or Amazon Web Services service
// owns the hosted zones. The HostedZoneOwner structure in the response contains
// one of the following values:
//
// - An OwningAccount element, which contains the account number of either the
// current Amazon Web Services account or another Amazon Web Services account. Some
// services, such as Cloud Map, create hosted zones using the current account.
//
// - An OwningService element, which identifies the Amazon Web Services service
// that created and owns the hosted zone. For example, if a hosted zone was created
// by Amazon Elastic File System (Amazon EFS), the value of Owner is
// efs.amazonaws.com .
//
// ListHostedZonesByVPC returns the hosted zones associated with the specified VPC
// and does not reflect the hosted zone associations to VPCs via Route 53 Profiles.
// To get the associations to a Profile, call the [ListProfileResourceAssociations]API.
//
// When listing private hosted zones, the hosted zone and the Amazon VPC must
// belong to the same partition where the hosted zones were created. A partition is
// a group of Amazon Web Services Regions. Each Amazon Web Services account is
// scoped to one partition.
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
// [ListProfileResourceAssociations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53profiles_ListProfileResourceAssociations.html
