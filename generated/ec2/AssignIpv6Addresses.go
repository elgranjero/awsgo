package ec2

// AssignIpv6Addresses is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Assigns the specified IPv6 addresses to the specified network interface. You
// can specify specific IPv6 addresses, or you can specify the number of IPv6
// addresses to be automatically assigned from the subnet's IPv6 CIDR block range.
// You can assign as many IPv6 addresses to a network interface as you can assign
// private IPv4 addresses, and the limit varies by instance type.
//
// You must specify either the IPv6 addresses or the IPv6 address count in the
// request.
//
// You can optionally use Prefix Delegation on the network interface. You must
// specify either the IPV6 Prefix Delegation prefixes, or the IPv6 Prefix
// Delegation count. For information, see [Assigning prefixes to network interfaces]in the Amazon EC2 User Guide.
//
// [Assigning prefixes to network interfaces]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html
