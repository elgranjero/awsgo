package eks

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Associates the specified tags to an Amazon EKS resource with the specified
// resourceArn . If existing tags on a resource are not specified in the request
// parameters, they aren't changed. When a resource is deleted, the tags associated
// with that resource are also deleted. Tags that you create for Amazon EKS
// resources don't propagate to any other resources associated with the cluster.
// For example, if you tag a cluster with this operation, that tag doesn't
// automatically propagate to the subnets and nodes associated with the cluster.
