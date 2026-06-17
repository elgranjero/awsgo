package ec2

// AllocateAddress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Allocates an Elastic IP address to your Amazon Web Services account. After you
// allocate the Elastic IP address you can associate it with an instance or network
// interface. After you release an Elastic IP address, it is released to the IP
// address pool and can be allocated to a different Amazon Web Services account.
//
// You can allocate an Elastic IP address from one of the following address pools:
//
// - Amazon's pool of IPv4 addresses
//
// - Public IPv4 address range that you own and bring to your Amazon Web
// Services account using [Bring Your Own IP Addresses (BYOIP)]
//
// - An IPv4 IPAM pool with an Amazon-provided or BYOIP public IPv4 address range
//
// - IPv4 addresses from your on-premises network made available for use with an
// Outpost using a [customer-owned IP address pool](CoIP pool)
//
// For more information, see [Elastic IP Addresses] in the Amazon EC2 User Guide.
//
// If you release an Elastic IP address, you might be able to recover it. You
// cannot recover an Elastic IP address that you released after it is allocated to
// another Amazon Web Services account. To attempt to recover an Elastic IP address
// that you released, specify it in this operation.
//
// You can allocate a carrier IP address which is a public IP address from a
// telecommunication carrier, to a network interface which resides in a subnet in a
// Wavelength Zone (for example an EC2 instance).
//
// [Bring Your Own IP Addresses (BYOIP)]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html
// [Elastic IP Addresses]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html
// [customer-owned IP address pool]: https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing
