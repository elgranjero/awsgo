package redshift

// ModifyClusterSubnetGroup is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Modifies a cluster subnet group to include the specified list of VPC subnets.
// The operation replaces the existing list of subnets with the new list of
// subnets.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// subnet group for a provisioned cluster is in an account with VPC BPA turned on,
// the following capabilities are blocked:
//
// - Creating a public cluster
//
// - Restoring a public cluster
//
// - Modifying a private cluster to be public
//
// - Adding a subnet with VPC BPA turned on to the subnet group when there's at
// least one public cluster within the group
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
