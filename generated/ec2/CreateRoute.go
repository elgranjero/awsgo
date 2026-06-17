package ec2

// CreateRoute is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a route in a route table within a VPC.
//
// You must specify either a destination CIDR block or a prefix list ID. You must
// also specify exactly one of the resources from the parameter list.
//
// When determining how to route traffic, we use the route with the most specific
// match. For example, traffic is destined for the IPv4 address 192.0.2.3 , and the
// route table includes the following two IPv4 routes:
//
// - 192.0.2.0/24 (goes to some target A)
//
// - 192.0.2.0/28 (goes to some target B)
//
// Both routes apply to the traffic destined for 192.0.2.3 . However, the second
// route in the list covers a smaller number of IP addresses and is therefore more
// specific, so we use that route to determine where to target the traffic.
//
// For more information about route tables, see [Route tables] in the Amazon VPC User Guide.
//
// [Route tables]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html
