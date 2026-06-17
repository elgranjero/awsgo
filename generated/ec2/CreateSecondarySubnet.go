package ec2

// CreateSecondarySubnet is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a secondary subnet in a secondary network.
//
// A secondary subnet CIDR block must not overlap with the CIDR block of an
// existing secondary subnet in the secondary network. After you create a secondary
// subnet, you can't change its CIDR block.
//
// The allowed size for a secondary subnet CIDR block is between /28 netmask (16
// IP addresses) and /12 netmask (1,048,576 IP addresses). Amazon reserves the
// first four IP addresses and the last IP address in each secondary subnet for
// internal use.
