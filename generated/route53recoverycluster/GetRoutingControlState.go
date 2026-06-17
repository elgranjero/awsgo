package route53recoverycluster

// GetRoutingControlState is generated as a reference stub.
// Executable command wiring lives under cmd/route53recoverycluster.go.
//
// Get the state for a routing control. A routing control is a simple on/off
// switch that you can use to route traffic to cells. When a routing control state
// is set to ON, traffic flows to a cell. When the state is set to OFF, traffic
// does not flow.
//
// Before you can create a routing control, you must first create a cluster, and
// then host the control in a control panel on the cluster. For more information,
// see [Create routing control structures]in the Amazon Route 53 Application Recovery Controller Developer Guide. You
// access one of the endpoints for the cluster to get or update the routing control
// state to redirect traffic for your application.
//
// You must specify Regional endpoints when you work with API cluster operations
// to get or update routing control states in Route 53 ARC.
//
// To see a code example for getting a routing control state, including accessing
// Regional cluster endpoints in sequence, see [API examples]in the Amazon Route 53 Application
// Recovery Controller Developer Guide.
//
// Learn more about working with routing controls in the following topics in the
// Amazon Route 53 Application Recovery Controller Developer Guide:
//
// [Viewing and updating routing control states]
//
// [Working with routing controls in Route 53 ARC]
//
// [API examples]: https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html
// [Working with routing controls in Route 53 ARC]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
// [Create routing control structures]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.create.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
