package directconnect

// CreateLag is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Creates a link aggregation group (LAG) with the specified number of bundled
// physical dedicated connections between the customer network and a specific
// Direct Connect location. A LAG is a logical interface that uses the Link
// Aggregation Control Protocol (LACP) to aggregate multiple interfaces, enabling
// you to treat them as a single interface.
//
// All connections in a LAG must use the same bandwidth (either 1Gbps, 10Gbps,
// 100Gbps, or 400Gbps) and must terminate at the same Direct Connect endpoint.
//
// You can have up to 10 dedicated connections per location. Regardless of this
// limit, if you request more connections for the LAG than Direct Connect can
// allocate on a single endpoint, no LAG is created..
//
// You can specify an existing physical dedicated connection or interconnect to
// include in the LAG (which counts towards the total number of connections). Doing
// so interrupts the current physical dedicated connection, and re-establishes them
// as a member of the LAG. The LAG will be created on the same Direct Connect
// endpoint to which the dedicated connection terminates. Any virtual interfaces
// associated with the dedicated connection are automatically disassociated and
// re-associated with the LAG. The connection ID does not change.
//
// If the Amazon Web Services account used to create a LAG is a registered Direct
// Connect Partner, the LAG is automatically enabled to host sub-connections. For a
// LAG owned by a partner, any associated virtual interfaces cannot be directly
// configured.
