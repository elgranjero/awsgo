package directconnect

// AssociateVirtualInterface is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Associates a virtual interface with a specified link aggregation group (LAG) or
// connection. Connectivity to Amazon Web Services is temporarily interrupted as
// the virtual interface is being migrated. If the target connection or LAG has an
// associated virtual interface with a conflicting VLAN number or a conflicting IP
// address, the operation fails.
//
// Virtual interfaces associated with a hosted connection cannot be associated
// with a LAG; hosted connections must be migrated along with their virtual
// interfaces using AssociateHostedConnection.
//
// To reassociate a virtual interface to a new connection or LAG, the requester
// must own either the virtual interface itself or the connection to which the
// virtual interface is currently associated. Additionally, the requester must own
// the connection or LAG for the association.
