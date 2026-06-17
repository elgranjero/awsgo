package ec2

// AllocateIpamPoolCidr is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Allocate a CIDR from an IPAM pool. The Region you use should be the IPAM pool
// locale. The locale is the Amazon Web Services Region where this IPAM pool is
// available for allocations.
//
// In IPAM, an allocation is a CIDR assignment from an IPAM pool to another IPAM
// pool or to a resource. For more information, see [Allocate CIDRs]in the Amazon VPC IPAM User
// Guide.
//
// This action creates an allocation with strong consistency. The returned CIDR
// will not overlap with any other allocations from the same pool.
//
// [Allocate CIDRs]: https://docs.aws.amazon.com/vpc/latest/ipam/allocate-cidrs-ipam.html
