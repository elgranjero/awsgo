package ecs

// DeregisterContainerInstance is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deregisters an Amazon ECS container instance from the specified cluster. This
// instance is no longer available to run tasks.
//
// If you intend to use the container instance for some other purpose after
// deregistration, we recommend that you stop all of the tasks running on the
// container instance before deregistration. That prevents any orphaned tasks from
// consuming resources.
//
// Deregistering a container instance removes the instance from a cluster, but it
// doesn't terminate the EC2 instance. If you are finished using the instance, be
// sure to terminate it in the Amazon EC2 console to stop billing.
//
// If you terminate a running container instance, Amazon ECS automatically
// deregisters the instance from your cluster (stopped container instances or
// instances with disconnected agents aren't automatically deregistered when
// terminated).
