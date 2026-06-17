package ec2

// AssociateNatGatewayAddress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Associates Elastic IP addresses (EIPs) and private IPv4 addresses with a public
// NAT gateway. For more information, see [Work with NAT gateways]in the Amazon VPC User Guide.
//
// By default, you can associate up to 2 Elastic IP addresses per public NAT
// gateway. You can increase the limit by requesting a quota adjustment. For more
// information, see [Elastic IP address quotas]in the Amazon VPC User Guide.
//
// When you associate an EIP or secondary EIPs with a public NAT gateway, the
// network border group of the EIPs must match the network border group of the
// Availability Zone (AZ) that the public NAT gateway is in. If it's not the same,
// the EIP will fail to associate. You can see the network border group for the
// subnet's AZ by viewing the details of the subnet. Similarly, you can view the
// network border group of an EIP by viewing the details of the EIP address. For
// more information about network border groups and EIPs, see [Allocate an Elastic IP address]in the Amazon VPC
// User Guide.
//
// [Elastic IP address quotas]: https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html#vpc-limits-eips
// [Work with NAT gateways]: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html
// [Allocate an Elastic IP address]: https://docs.aws.amazon.com/vpc/latest/userguide/WorkWithEIPs.html
