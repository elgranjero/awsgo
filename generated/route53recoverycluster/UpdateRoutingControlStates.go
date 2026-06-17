package route53recoverycluster

// UpdateRoutingControlStates is generated as a reference stub.
// Executable command wiring lives under cmd/route53recoverycluster.go.
//
// Set multiple routing control states. You can set the value for each state to be
// ON or OFF. When the state is ON, traffic flows to a cell. When it's OFF, traffic
// does not flow.
//
// With Route 53 ARC, you can add safety rules for routing controls, which are
// safeguards for routing control state updates that help prevent unexpected
// outcomes, like fail open traffic routing. However, there are scenarios when you
// might want to bypass the routing control safeguards that are enforced with
// safety rules that you've configured. For example, you might want to fail over
// quickly for disaster recovery, and one or more safety rules might be
// unexpectedly preventing you from updating a routing control state to reroute
// traffic. In a "break glass" scenario like this, you can override one or more
// safety rules to change a routing control state and fail over your application.
//
// The SafetyRulesToOverride property enables you override one or more safety
// rules and update routing control states. For more information, see [Override safety rules to reroute traffic]in the
// Amazon Route 53 Application Recovery Controller Developer Guide.
//
// You must specify Regional endpoints when you work with API cluster operations
// to get or update routing control states in Route 53 ARC.
//
// To see a code example for getting a routing control state, including accessing
// Regional cluster endpoints in sequence, see [API examples]in the Amazon Route 53 Application
// Recovery Controller Developer Guide.
//
// [Viewing and updating routing control states]
//
// [Working with routing controls overall]
//
// [API examples]: https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
// [Override safety rules to reroute traffic]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.override-safety-rule.html
// [Working with routing controls overall]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
