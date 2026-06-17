package emr

// SetUnhealthyNodeReplacement is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// Specify whether to enable unhealthy node replacement, which lets Amazon EMR
// gracefully replace core nodes on a cluster if any nodes become unhealthy. For
// example, a node becomes unhealthy if disk usage is above 90%. If unhealthy node
// replacement is on and TerminationProtected are off, Amazon EMR immediately
// terminates the unhealthy core nodes. To use unhealthy node replacement and
// retain unhealthy core nodes, use to turn on termination protection. In such
// cases, Amazon EMR adds the unhealthy nodes to a denylist, reducing job
// interruptions and failures.
//
// If unhealthy node replacement is on, Amazon EMR notifies YARN and other
// applications on the cluster to stop scheduling tasks with these nodes, moves the
// data, and then terminates the nodes.
//
// For more information, see [graceful node replacement] in the Amazon EMR Management Guide.
//
// [graceful node replacement]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-node-replacement.html
