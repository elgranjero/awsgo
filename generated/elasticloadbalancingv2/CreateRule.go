package elasticloadbalancingv2

// CreateRule is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancingv2.go.
//
// Creates a rule for the specified listener. The listener must be associated with
// an Application Load Balancer.
//
// Each rule consists of a priority, one or more actions, one or more conditions,
// and up to two optional transforms. Rules are evaluated in priority order, from
// the lowest value to the highest value. When the conditions for a rule are met,
// its actions are performed. If the conditions for no rules are met, the actions
// for the default rule are performed. For more information, see [Listener rules]in the
// Application Load Balancers Guide.
//
// [Listener rules]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules
