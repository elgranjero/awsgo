package ec2

// RegisterTransitGatewayMulticastGroupSources is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Registers sources (network interfaces) with the specified transit gateway
// multicast group.
//
// A multicast source is a network interface attached to a supported instance that
// sends multicast traffic. For more information about supported instances, see [Multicast on transit gateways]in
// the Amazon Web Services Transit Gateways Guide.
//
// After you add the source, use [SearchTransitGatewayMulticastGroups] to verify that the source was added to the
// multicast group.
//
// [SearchTransitGatewayMulticastGroups]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayMulticastGroups.html
// [Multicast on transit gateways]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-multicast-overview.html
