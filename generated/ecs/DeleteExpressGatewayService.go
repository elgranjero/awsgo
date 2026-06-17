package ecs

// DeleteExpressGatewayService is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deletes an Express service and removes all associated Amazon Web Services
// resources. This operation stops service tasks, removes the Application Load
// Balancer, target groups, security groups, auto-scaling policies, and other
// managed infrastructure components.
//
// The service enters a DRAINING state where existing tasks complete current
// requests without starting new tasks. After all tasks stop, the service and
// infrastructure are permanently removed.
//
// This operation cannot be reversed. Back up important data and verify the
// service is no longer needed before deletion.
