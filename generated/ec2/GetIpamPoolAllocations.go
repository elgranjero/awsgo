package ec2

// GetIpamPoolAllocations is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Get a list of all the CIDR allocations in an IPAM pool. The Region you use
// should be the IPAM pool locale. The locale is the Amazon Web Services Region
// where this IPAM pool is available for allocations.
//
// If you use this action after [AllocateIpamPoolCidr] or [ReleaseIpamPoolAllocation], note that all EC2 API actions follow an [eventual consistency]
// model.
//
// [ReleaseIpamPoolAllocation]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseIpamPoolAllocation.html
// [AllocateIpamPoolCidr]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllocateIpamPoolCidr.html
// [eventual consistency]: https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html
