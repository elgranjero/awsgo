package directconnect

// CreateTransitVirtualInterface is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Creates a transit virtual interface. A transit virtual interface should be used
// to access one or more transit gateways associated with Direct Connect gateways.
// A transit virtual interface enables the connection of multiple VPCs attached to
// a transit gateway to a Direct Connect gateway.
//
// If you associate your transit gateway with one or more Direct Connect gateways,
// the Autonomous System Number (ASN) used by the transit gateway and the Direct
// Connect gateway must be different. For example, if you use the default ASN 64512
// for both your the transit gateway and Direct Connect gateway, the association
// request fails.
//
// A jumbo MTU value must be either 1500 or 8500. No other values will be
// accepted. Setting the MTU of a virtual interface to 8500 (jumbo frames) can
// cause an update to the underlying physical connection if it wasn't updated to
// support jumbo frames. Updating the connection disrupts network connectivity for
// all virtual interfaces associated with the connection for up to 30 seconds. To
// check whether your connection supports jumbo frames, call DescribeConnections. To check whether
// your virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
