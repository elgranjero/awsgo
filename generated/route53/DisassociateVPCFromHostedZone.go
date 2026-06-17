package route53

// DisassociateVPCFromHostedZone is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Disassociates an Amazon Virtual Private Cloud (Amazon VPC) from an Amazon Route
// 53 private hosted zone. Note the following:
//
// - You can't disassociate the last Amazon VPC from a private hosted zone.
//
// - You can't convert a private hosted zone into a public hosted zone.
//
// - You can submit a DisassociateVPCFromHostedZone request using either the
// account that created the hosted zone or the account that created the Amazon VPC.
//
// - Some services, such as Cloud Map and Amazon Elastic File System (Amazon
// EFS) automatically create hosted zones and associate VPCs with the hosted zones.
// A service can create a hosted zone using your account or using its own account.
// You can disassociate a VPC from a hosted zone only if the service created the
// hosted zone using your account.
//
// When you run [DisassociateVPCFromHostedZone], if the hosted zone has a value for OwningAccount , you can use
//
// DisassociateVPCFromHostedZone . If the hosted zone has a value for
// OwningService , you can't use DisassociateVPCFromHostedZone .
//
// When revoking access, the hosted zone and the Amazon VPC must belong to the
// same partition. A partition is a group of Amazon Web Services Regions. Each
// Amazon Web Services account is scoped to one partition.
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
// [DisassociateVPCFromHostedZone]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListHostedZonesByVPC.html
