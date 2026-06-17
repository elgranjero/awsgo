package elasticloadbalancingv2

// DeleteLoadBalancer is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancingv2.go.
//
// Deletes the specified Application Load Balancer, Network Load Balancer, or
// Gateway Load Balancer. Deleting a load balancer also deletes its listeners.
//
// You can't delete a load balancer if deletion protection is enabled. If the load
// balancer does not exist or has already been deleted, the call succeeds.
//
// Deleting a load balancer does not affect its registered targets. For example,
// your EC2 instances continue to run and are still registered to their target
// groups. If you no longer need these EC2 instances, you can stop or terminate
// them.
