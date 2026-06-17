package ec2

// ReleaseIpamPoolAllocation is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Release an allocation within an IPAM pool. The Region you use should be the
// IPAM pool locale. The locale is the Amazon Web Services Region where this IPAM
// pool is available for allocations. You can only use this action to release
// manual allocations. To remove an allocation for a resource without deleting the
// resource, set its monitored state to false using [ModifyIpamResourceCidr]. For more information, see [Release an allocation]
// in the Amazon VPC IPAM User Guide.
//
// All EC2 API actions follow an [eventual consistency] model.
//
// [Release an allocation]: https://docs.aws.amazon.com/vpc/latest/ipam/release-alloc-ipam.html
// [eventual consistency]: https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html
// [ModifyIpamResourceCidr]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyIpamResourceCidr.html
