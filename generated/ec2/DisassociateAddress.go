package ec2

// DisassociateAddress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disassociates an Elastic IP address from the instance or network interface it's
// associated with.
//
// This is an idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error.
//
// An address cannot be disassociated if the all of the following conditions are
// met:
//
// - Network interface has a publicDualStackDnsName publicDnsName
//
// - Public IPv4 address is the primary public IPv4 address
//
// - Network interface only has one remaining public IPv4 address
