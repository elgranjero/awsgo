package elasticloadbalancing

// DisableAvailabilityZonesForLoadBalancer is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancing.go.
//
// Removes the specified Availability Zones from the set of Availability Zones for
// the specified load balancer in EC2-Classic or a default VPC.
//
// For load balancers in a non-default VPC, use DetachLoadBalancerFromSubnets.
//
// There must be at least one Availability Zone registered with a load balancer at
// all times. After an Availability Zone is removed, all instances registered with
// the load balancer that are in the removed Availability Zone go into the
// OutOfService state. Then, the load balancer attempts to equally balance the
// traffic among its remaining Availability Zones.
//
// For more information, see [Add or Remove Availability Zones] in the Classic Load Balancers Guide.
//
// [Add or Remove Availability Zones]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-az.html
