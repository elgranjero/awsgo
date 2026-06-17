package ec2

// DisableRouteServerPropagation is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disables route propagation from a route server to a specified route table.
//
// When enabled, route server propagation installs the routes in the FIB on the
// route table you've specified. Route server supports IPv4 and IPv6 route
// propagation.
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
// For more information see [Dynamic routing in your VPC with VPC Route Server] in the Amazon VPC User Guide.
//
// [Dynamic routing in your VPC with VPC Route Server]: https://docs.aws.amazon.com/vpc/latest/userguide/dynamic-routing-route-server.html
// [Transit Gateway Connect]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-connect.html
