package ec2

// CreateNatGateway is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a NAT gateway in the specified subnet. This action creates a network
// interface in the specified subnet with a private IP address from the IP address
// range of the subnet. You can create either a public NAT gateway or a private NAT
// gateway.
//
// With a public NAT gateway, internet-bound traffic from a private subnet can be
// routed to the NAT gateway, so that instances in a private subnet can connect to
// the internet.
//
// With a private NAT gateway, private communication is routed across VPCs and
// on-premises networks through a transit gateway or virtual private gateway.
// Common use cases include running large workloads behind a small pool of
// allowlisted IPv4 addresses, preserving private IPv4 addresses, and communicating
// between overlapping networks.
//
// For more information, see [NAT gateways] in the Amazon VPC User Guide.
//
// When you create a public NAT gateway and assign it an EIP or secondary EIPs,
// the network border group of the EIPs must match the network border group of the
// Availability Zone (AZ) that the public NAT gateway is in. If it's not the same,
// the NAT gateway will fail to launch. You can see the network border group for
// the subnet's AZ by viewing the details of the subnet. Similarly, you can view
// the network border group of an EIP by viewing the details of the EIP address.
// For more information about network border groups and EIPs, see [Allocate an Elastic IP address]in the Amazon
// VPC User Guide.
//
// [NAT gateways]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
// [Allocate an Elastic IP address]: https://docs.aws.amazon.com/vpc/latest/userguide/WorkWithEIPs.html
