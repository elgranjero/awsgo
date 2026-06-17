package networkfirewall

// ListFlowOperations is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Returns a list of all flow operations ran in a specific firewall. You can
// optionally narrow the request scope by specifying the operation type or
// Availability Zone associated with a firewall's flow operations.
//
// Flow operations let you manage the flows tracked in the flow table, also known
// as the firewall table.
//
// A flow is network traffic that is monitored by a firewall, either by stateful
// or stateless rules. For traffic to be considered part of a flow, it must share
// Destination, DestinationPort, Direction, Protocol, Source, and SourcePort.
