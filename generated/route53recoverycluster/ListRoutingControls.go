package route53recoverycluster

// ListRoutingControls is generated as a reference stub.
// Executable command wiring lives under cmd/route53recoverycluster.go.
//
// List routing control names and Amazon Resource Names (ARNs), as well as the
// routing control state for each routing control, along with the control panel
// name and control panel ARN for the routing controls. If you specify a control
// panel ARN, this call lists the routing controls in the control panel. Otherwise,
// it lists all the routing controls in the cluster.
//
// A routing control is a simple on/off switch in Route 53 ARC that you can use to
// route traffic to cells. When a routing control state is set to ON, traffic flows
// to a cell. When the state is set to OFF, traffic does not flow.
//
// Before you can create a routing control, you must first create a cluster, and
// then host the control in a control panel on the cluster. For more information,
// see [Create routing control structures]in the Amazon Route 53 Application Recovery Controller Developer Guide. You
// access one of the endpoints for the cluster to get or update the routing control
// state to redirect traffic for your application.
//
// You must specify Regional endpoints when you work with API cluster operations
// to use this API operation to list routing controls in Route 53 ARC.
//
// Learn more about working with routing controls in the following topics in the
// Amazon Route 53 Application Recovery Controller Developer Guide:
//
// [Viewing and updating routing control states]
//
// [Working with routing controls in Route 53 ARC]
//
// [Working with routing controls in Route 53 ARC]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
// [Create routing control structures]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.create.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
