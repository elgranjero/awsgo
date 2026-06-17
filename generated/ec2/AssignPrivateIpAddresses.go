package ec2

// AssignPrivateIpAddresses is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Assigns the specified secondary private IP addresses to the specified network
// interface.
//
// You can specify specific secondary IP addresses, or you can specify the number
// of secondary IP addresses to be automatically assigned from the subnet's CIDR
// block range. The number of secondary IP addresses that you can assign to an
// instance varies by instance type. For more information about Elastic IP
// addresses, see [Elastic IP Addresses]in the Amazon EC2 User Guide.
//
// When you move a secondary private IP address to another network interface, any
// Elastic IP address that is associated with the IP address is also moved.
//
// Remapping an IP address is an asynchronous operation. When you move an IP
// address from one network interface to another, check
// network/interfaces/macs/mac/local-ipv4s in the instance metadata to confirm that
// the remapping is complete.
//
// You must specify either the IP addresses or the IP address count in the request.
//
// You can optionally use Prefix Delegation on the network interface. You must
// specify either the IPv4 Prefix Delegation prefixes, or the IPv4 Prefix
// Delegation count. For information, see [Assigning prefixes to network interfaces]in the Amazon EC2 User Guide.
//
// [Elastic IP Addresses]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html
// [Assigning prefixes to network interfaces]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html
