package elasticloadbalancingv2

// DeregisterTargets is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancingv2.go.
//
// Deregisters the specified targets from the specified target group. After the
// targets are deregistered, they no longer receive traffic from the load balancer.
//
// The load balancer stops sending requests to targets that are deregistering, but
// uses connection draining to ensure that in-flight traffic completes on the
// existing connections. This deregistration delay is configured by default but can
// be updated for each target group.
//
// For more information, see the following:
//
// [Deregistration delay]
// - in the Application Load Balancers User Guide
//
// [Deregistration delay]
// - in the Network Load Balancers User Guide
//
// [Deregistration delay]
// - in the Gateway Load Balancers User Guide
//
// Note: If the specified target does not exist, the action returns successfully.
//
// [Deregistration delay]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/edit-target-group-attributes.html#deregistration-delay
