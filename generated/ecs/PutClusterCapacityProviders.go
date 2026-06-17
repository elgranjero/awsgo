package ecs

// PutClusterCapacityProviders is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Modifies the available capacity providers and the default capacity provider
// strategy for a cluster.
//
// You must specify both the available capacity providers and a default capacity
// provider strategy for the cluster. If the specified cluster has existing
// capacity providers associated with it, you must specify all existing capacity
// providers in addition to any new ones you want to add. Any existing capacity
// providers that are associated with a cluster that are omitted from a [PutClusterCapacityProviders]API call
// will be disassociated with the cluster. You can only disassociate an existing
// capacity provider from a cluster if it's not being used by any existing tasks.
//
// When creating a service or running a task on a cluster, if no capacity provider
// or launch type is specified, then the cluster's default capacity provider
// strategy is used. We recommend that you define a default capacity provider
// strategy for your cluster. However, you must specify an empty array ( [] ) to
// bypass defining a default strategy.
//
// Amazon ECS Managed Instances doesn't support this, because when you create a
// capacity provider with Amazon ECS Managed Instances, it becomes available only
// within the specified cluster.
//
// [PutClusterCapacityProviders]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutClusterCapacityProviders.html
