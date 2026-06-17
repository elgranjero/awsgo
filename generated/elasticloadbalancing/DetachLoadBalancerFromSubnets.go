package elasticloadbalancing

// DetachLoadBalancerFromSubnets is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancing.go.
//
// Removes the specified subnets from the set of configured subnets for the load
// balancer.
//
// After a subnet is removed, all EC2 instances registered with the load balancer
// in the removed subnet go into the OutOfService state. Then, the load balancer
// balances the traffic among the remaining routable subnets.
