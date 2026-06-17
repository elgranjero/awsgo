package ecs

// UpdateExpressGatewayService is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Updates an existing Express service configuration. Modifies container settings,
// resource allocation, auto-scaling configuration, and other service parameters
// without recreating the service.
//
// Amazon ECS creates a new service revision with updated configuration and
// performs a rolling deployment to replace existing tasks. The service remains
// available during updates, ensuring zero-downtime deployments.
//
// Some parameters like the infrastructure role cannot be modified after service
// creation and require creating a new service.
