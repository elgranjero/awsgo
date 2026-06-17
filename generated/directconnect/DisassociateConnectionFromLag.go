package directconnect

// DisassociateConnectionFromLag is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Disassociates a connection from a link aggregation group (LAG). The connection
// is interrupted and re-established as a standalone connection (the connection is
// not deleted; to delete the connection, use the DeleteConnectionrequest). If the LAG has
// associated virtual interfaces or hosted connections, they remain associated with
// the LAG. A disassociated connection owned by an Direct Connect Partner is
// automatically converted to an interconnect.
//
// If disassociating the connection would cause the LAG to fall below its setting
// for minimum number of operational connections, the request fails, except when
// it's the last member of the LAG. If all connections are disassociated, the LAG
// continues to exist as an empty LAG with no physical connections.
