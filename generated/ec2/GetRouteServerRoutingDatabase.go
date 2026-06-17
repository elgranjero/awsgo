package ec2

// GetRouteServerRoutingDatabase is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Gets the routing database for the specified route server. The [Routing Information Base (RIB)] serves as a
// database that stores all the routing information and network topology data
// collected by a router or routing system, such as routes learned from BGP peers.
// The RIB is constantly updated as new routing information is received or existing
// routes change. This ensures that the route server always has the most current
// view of the network topology and can make optimal routing decisions.
//
// Amazon VPC Route Server simplifies routing for traffic between workloads that
// are deployed within a VPC and its internet gateways. With this feature, VPC
// Route Server dynamically updates VPC and internet gateway route tables with your
// preferred IPv4 or IPv6 routes to achieve routing fault tolerance for those
// workloads. This enables you to automatically reroute traffic within a VPC, which
// increases the manageability of VPC routing and interoperability with third-party
// workloads.
//
// Route server supports the follow route table types:
//
// - VPC route tables not associated with subnets
//
// - Subnet route tables
//
// - Internet gateway route tables
//
// Route server does not support route tables associated with virtual private
// gateways. To propagate routes into a transit gateway route table, use [Transit Gateway Connect].
//
// [Routing Information Base (RIB)]: https://en.wikipedia.org/wiki/Routing_table
// [Transit Gateway Connect]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-connect.html
