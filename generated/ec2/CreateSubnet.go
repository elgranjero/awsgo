package ec2

// CreateSubnet is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a subnet in the specified VPC. For an IPv4 only subnet, specify an IPv4
// CIDR block. If the VPC has an IPv6 CIDR block, you can create an IPv6 only
// subnet or a dual stack subnet instead. For an IPv6 only subnet, specify an IPv6
// CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and an IPv6
// CIDR block.
//
// A subnet CIDR block must not overlap the CIDR block of an existing subnet in
// the VPC. After you create a subnet, you can't change its CIDR block.
//
// The allowed size for an IPv4 subnet is between a /28 netmask (16 IP addresses)
// and a /16 netmask (65,536 IP addresses). Amazon Web Services reserves both the
// first four and the last IPv4 address in each subnet's CIDR block. They're not
// available for your use.
//
// If you've associated an IPv6 CIDR block with your VPC, you can associate an
// IPv6 CIDR block with a subnet when you create it.
//
// If you add more than one subnet to a VPC, they're set up in a star topology
// with a logical router in the middle.
//
// When you stop an instance in a subnet, it retains its private IPv4 address.
// It's therefore possible to have a subnet with no running instances (they're all
// stopped), but no remaining IP addresses available.
//
// For more information, see [Subnets] in the Amazon VPC User Guide.
//
// [Subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html
