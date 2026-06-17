package ec2

// RegisterTransitGatewayMulticastGroupMembers is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Registers members (network interfaces) with the transit gateway multicast
// group. A member is a network interface associated with a supported EC2 instance
// that receives multicast traffic. For more information, see [Multicast on transit gateways]in the Amazon Web
// Services Transit Gateways Guide.
//
// After you add the members, use [SearchTransitGatewayMulticastGroups] to verify that the members were added to the
// transit gateway multicast group.
//
// [SearchTransitGatewayMulticastGroups]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayMulticastGroups.html
// [Multicast on transit gateways]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-multicast-overview.html
