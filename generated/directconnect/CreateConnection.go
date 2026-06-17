package directconnect

// CreateConnection is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Creates a connection between a customer network and a specific Direct Connect
// location.
//
// A connection links your internal network to an Direct Connect location over a
// standard Ethernet fiber-optic cable. One end of the cable is connected to your
// router, the other to an Direct Connect router.
//
// To find the locations for your Region, use DescribeLocations.
//
// You can automatically add the new connection to a link aggregation group (LAG)
// by specifying a LAG ID in the request. This ensures that the new connection is
// allocated on the same Direct Connect endpoint that hosts the specified LAG. If
// there are no available ports on the endpoint, the request fails and no
// connection is created.
