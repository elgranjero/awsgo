package elasticloadbalancing

// DeleteLoadBalancer is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancing.go.
//
// Deletes the specified load balancer.
//
// If you are attempting to recreate a load balancer, you must reconfigure all
// settings. The DNS name associated with a deleted load balancer are no longer
// usable. The name and associated DNS record of the deleted load balancer no
// longer exist and traffic sent to any of its IP addresses is no longer delivered
// to your instances.
//
// If the load balancer does not exist or has already been deleted, the call to
// DeleteLoadBalancer still succeeds.
