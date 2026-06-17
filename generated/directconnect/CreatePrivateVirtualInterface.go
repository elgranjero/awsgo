package directconnect

// CreatePrivateVirtualInterface is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Creates a private virtual interface. A virtual interface is the VLAN that
// transports Direct Connect traffic. A private virtual interface can be connected
// to either a Direct Connect gateway or a Virtual Private Gateway (VGW).
// Connecting the private virtual interface to a Direct Connect gateway enables the
// possibility for connecting to multiple VPCs, including VPCs in different Amazon
// Web Services Regions. Connecting the private virtual interface to a VGW only
// provides access to a single VPC within the same Region.
//
// Setting the MTU of a virtual interface to 8500 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To check
// whether your connection supports jumbo frames, call DescribeConnections. To check whether your
// virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
