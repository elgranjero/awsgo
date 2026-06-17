package directconnect

// AssociateConnectionWithLag is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Associates an existing connection with a link aggregation group (LAG). The
// connection is interrupted and re-established as a member of the LAG
// (connectivity to Amazon Web Services is interrupted). The connection must be
// hosted on the same Direct Connect endpoint as the LAG, and its bandwidth must
// match the bandwidth for the LAG. You can re-associate a connection that's
// currently associated with a different LAG; however, if removing the connection
// would cause the original LAG to fall below its setting for minimum number of
// operational connections, the request fails.
//
// Any virtual interfaces that are directly associated with the connection are
// automatically re-associated with the LAG. If the connection was originally
// associated with a different LAG, the virtual interfaces remain associated with
// the original LAG.
//
// For interconnects, any hosted connections are automatically re-associated with
// the LAG. If the interconnect was originally associated with a different LAG, the
// hosted connections remain associated with the original LAG.
