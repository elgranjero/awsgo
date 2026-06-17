package ec2

// CreateVpc is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a VPC with the specified CIDR blocks.
//
// A VPC must have an associated IPv4 CIDR block. You can choose an IPv4 CIDR
// block or an IPAM-allocated IPv4 CIDR block. You can optionally associate an IPv6
// CIDR block with a VPC. You can choose an IPv6 CIDR block, an Amazon-provided
// IPv6 CIDR block, an IPAM-allocated IPv6 CIDR block, or an IPv6 CIDR block that
// you brought to Amazon Web Services. For more information, see [IP addressing for your VPCs and subnets]in the Amazon VPC
// User Guide.
//
// By default, each instance that you launch in the VPC has the default DHCP
// options, which include only a default DNS server that we provide
// (AmazonProvidedDNS). For more information, see [DHCP option sets]in the Amazon VPC User Guide.
//
// You can specify DNS options and tenancy for a VPC when you create it. You can't
// change the tenancy of a VPC after you create it. For more information, see [VPC configuration options]in
// the Amazon VPC User Guide.
//
// [VPC configuration options]: https://docs.aws.amazon.com/vpc/latest/userguide/create-vpc-options.html
// [DHCP option sets]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html
// [IP addressing for your VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-ip-addressing.html
