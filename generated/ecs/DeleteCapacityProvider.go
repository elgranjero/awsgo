package ecs

// DeleteCapacityProvider is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deletes the specified capacity provider.
//
// The FARGATE and FARGATE_SPOT capacity providers are reserved and can't be
// deleted. You can disassociate them from a cluster using either [PutClusterCapacityProviders]or by deleting
// the cluster.
//
// Prior to a capacity provider being deleted, the capacity provider must be
// removed from the capacity provider strategy from all services. The [UpdateService]API can be
// used to remove a capacity provider from a service's capacity provider strategy.
// When updating a service, the forceNewDeployment option can be used to ensure
// that any tasks using the Amazon EC2 instance capacity provided by the capacity
// provider are transitioned to use the capacity from the remaining capacity
// providers. Only capacity providers that aren't associated with a cluster can be
// deleted. To remove a capacity provider from a cluster, you can either use [PutClusterCapacityProviders]or
// delete the cluster.
//
// [PutClusterCapacityProviders]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutClusterCapacityProviders.html
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
