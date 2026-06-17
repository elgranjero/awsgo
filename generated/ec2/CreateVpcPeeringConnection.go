package ec2

// CreateVpcPeeringConnection is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Requests a VPC peering connection between two VPCs: a requester VPC that you
// own and an accepter VPC with which to create the connection. The accepter VPC
// can belong to another Amazon Web Services account and can be in a different
// Region to the requester VPC. The requester VPC and accepter VPC cannot have
// overlapping CIDR blocks.
//
// Limitations and rules apply to a VPC peering connection. For more information,
// see the [VPC peering limitations]in the VPC Peering Guide.
//
// The owner of the accepter VPC must accept the peering request to activate the
// peering connection. The VPC peering connection request expires after 7 days,
// after which it cannot be accepted or rejected.
//
// If you create a VPC peering connection request between VPCs with overlapping
// CIDR blocks, the VPC peering connection has a status of failed .
//
// [VPC peering limitations]: https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-limitations
