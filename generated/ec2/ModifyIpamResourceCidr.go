package ec2

// ModifyIpamResourceCidr is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modify a resource CIDR. You can use this action to transfer resource CIDRs
// between scopes and ignore resource CIDRs that you do not want to manage. If set
// to false, the resource will not be tracked for overlap, it cannot be
// auto-imported into a pool, and it will be removed from any pool it has an
// allocation in.
//
// For more information, see [Move resource CIDRs between scopes] and [Change the monitoring state of resource CIDRs] in the Amazon VPC IPAM User Guide.
//
// [Change the monitoring state of resource CIDRs]: https://docs.aws.amazon.com/vpc/latest/ipam/change-monitoring-state-ipam.html
// [Move resource CIDRs between scopes]: https://docs.aws.amazon.com/vpc/latest/ipam/move-resource-ipam.html
