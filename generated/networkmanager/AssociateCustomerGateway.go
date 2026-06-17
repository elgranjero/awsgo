package networkmanager

// AssociateCustomerGateway is generated as a reference stub.
// Executable command wiring lives under cmd/networkmanager.go.
//
// Associates a customer gateway with a device and optionally, with a link. If you
// specify a link, it must be associated with the specified device.
//
// You can only associate customer gateways that are connected to a VPN attachment
// on a transit gateway or core network registered in your global network. When you
// register a transit gateway or core network, customer gateways that are connected
// to the transit gateway are automatically included in the global network. To list
// customer gateways that are connected to a transit gateway, use the [DescribeVpnConnections]EC2 API and
// filter by transit-gateway-id .
//
// You cannot associate a customer gateway with more than one device and link.
//
// [DescribeVpnConnections]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnConnections.html
