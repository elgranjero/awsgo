package directconnect

// CreateInterconnect is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Creates an interconnect between an Direct Connect Partner's network and a
// specific Direct Connect location.
//
// An interconnect is a connection that is capable of hosting other connections.
// The Direct Connect Partner can use an interconnect to provide Direct Connect
// hosted connections to customers through their own network services. Like a
// standard connection, an interconnect links the partner's network to an Direct
// Connect location over a standard Ethernet fiber-optic cable. One end is
// connected to the partner's router, the other to an Direct Connect router.
//
// You can automatically add the new interconnect to a link aggregation group
// (LAG) by specifying a LAG ID in the request. This ensures that the new
// interconnect is allocated on the same Direct Connect endpoint that hosts the
// specified LAG. If there are no available ports on the endpoint, the request
// fails and no interconnect is created.
//
// For each end customer, the Direct Connect Partner provisions a connection on
// their interconnect by calling AllocateHostedConnection. The end customer can then connect to Amazon Web
// Services resources by creating a virtual interface on their connection, using
// the VLAN assigned to them by the Direct Connect Partner.
//
// Intended for use by Direct Connect Partners only.
