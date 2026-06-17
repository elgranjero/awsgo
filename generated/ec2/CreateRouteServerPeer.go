package ec2

// CreateRouteServerPeer is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a new BGP peer for a specified route server endpoint.
//
// A route server peer is a session between a route server endpoint and the device
// deployed in Amazon Web Services (such as a firewall appliance or other network
// security function running on an EC2 instance). The device must meet these
// requirements:
//
// - Have an elastic network interface in the VPC
//
// - Support BGP (Border Gateway Protocol)
//
// - Can initiate BGP sessions
//
// For more information see [Dynamic routing in your VPC with VPC Route Server] in the Amazon VPC User Guide.
//
// [Dynamic routing in your VPC with VPC Route Server]: https://docs.aws.amazon.com/vpc/latest/userguide/dynamic-routing-route-server.html
