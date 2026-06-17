package ec2

// DeleteIpamPool is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Delete an IPAM pool.
//
// You cannot delete an IPAM pool if there are allocations in it or CIDRs
// provisioned to it. To release allocations, see [ReleaseIpamPoolAllocation]. To deprovision pool CIDRs, see [DeprovisionIpamPoolCidr]
// .
//
// For more information, see [Delete a pool] in the Amazon VPC IPAM User Guide.
//
// [ReleaseIpamPoolAllocation]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseIpamPoolAllocation.html
// [Delete a pool]: https://docs.aws.amazon.com/vpc/latest/ipam/delete-pool-ipam.html
// [DeprovisionIpamPoolCidr]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeprovisionIpamPoolCidr.html
