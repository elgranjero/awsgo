package elasticloadbalancingv2

// RegisterTargets is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancingv2.go.
//
// Registers the specified targets with the specified target group.
//
// If the target is an EC2 instance, it must be in the running state when you
// register it.
//
// By default, the load balancer routes requests to registered targets using the
// protocol and port for the target group. Alternatively, you can override the port
// for a target when you register it. You can register each EC2 instance or IP
// address with the same target group multiple times using different ports.
//
// For more information, see the following:
//
// [Register targets for your Application Load Balancer]
//
// [Register targets for your Network Load Balancer]
//
// [Register targets for your Gateway Load Balancer]
//
// [Register targets for your Network Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/target-group-register-targets.html
// [Register targets for your Gateway Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-group-register-targets.html
// [Register targets for your Application Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-register-targets.html
