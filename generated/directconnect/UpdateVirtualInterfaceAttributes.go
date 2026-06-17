package directconnect

// UpdateVirtualInterfaceAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Updates the specified attributes of the specified virtual private interface.
//
// Setting the MTU of a virtual interface to 8500 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To check
// whether your connection supports jumbo frames, call DescribeConnections. To check whether your
// virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
