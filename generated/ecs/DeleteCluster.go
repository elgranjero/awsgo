package ecs

// DeleteCluster is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deletes the specified cluster. The cluster transitions to the INACTIVE state.
// Clusters with an INACTIVE status might remain discoverable in your account for
// a period of time. However, this behavior is subject to change in the future. We
// don't recommend that you rely on INACTIVE clusters persisting.
//
// You must deregister all container instances from this cluster before you may
// delete it. You can list the container instances in a cluster with [ListContainerInstances]and
// deregister them with [DeregisterContainerInstance].
//
// [ListContainerInstances]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListContainerInstances.html
// [DeregisterContainerInstance]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeregisterContainerInstance.html
