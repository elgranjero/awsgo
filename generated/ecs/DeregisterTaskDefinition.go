package ecs

// DeregisterTaskDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deregisters the specified task definition by family and revision. Upon
// deregistration, the task definition is marked as INACTIVE . Existing tasks and
// services that reference an INACTIVE task definition continue to run without
// disruption. Existing services that reference an INACTIVE task definition can
// still scale up or down by modifying the service's desired count. If you want to
// delete a task definition revision, you must first deregister the task definition
// revision.
//
// You can't use an INACTIVE task definition to run new tasks or create new
// services, and you can't update an existing service to reference an INACTIVE
// task definition. However, there may be up to a 10-minute window following
// deregistration where these restrictions have not yet taken effect.
//
// At this time, INACTIVE task definitions remain discoverable in your account
// indefinitely. However, this behavior is subject to change in the future. We
// don't recommend that you rely on INACTIVE task definitions persisting beyond
// the lifecycle of any associated tasks and services.
//
// You must deregister a task definition revision before you delete it. For more
// information, see [DeleteTaskDefinitions].
//
// [DeleteTaskDefinitions]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeleteTaskDefinitions.html
