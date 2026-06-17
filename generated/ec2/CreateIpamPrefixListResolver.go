package ec2

// CreateIpamPrefixListResolver is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an IPAM prefix list resolver.
//
// An IPAM prefix list resolver is a component that manages the synchronization
// between IPAM's CIDR selection rules and customer-managed prefix lists. It
// automates connectivity configurations by selecting CIDRs from IPAM's database
// based on your business logic and synchronizing them with prefix lists used in
// resources such as VPC route tables and security groups.
//
// For more information about IPAM prefix list resolver, see [Automate prefix list updates with IPAM] in the Amazon VPC
// IPAM User Guide.
//
// [Automate prefix list updates with IPAM]: https://docs.aws.amazon.com/vpc/latest/ipam/automate-prefix-list-updates.html
