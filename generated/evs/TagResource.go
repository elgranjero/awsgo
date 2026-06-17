package evs

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/evs.go.
//
// Associates the specified tags to an Amazon EVS resource with the specified
// resourceArn . If existing tags on a resource are not specified in the request
// parameters, they aren't changed. When a resource is deleted, the tags associated
// with that resource are also deleted. Tags that you create for Amazon EVS
// resources don't propagate to any other resources associated with the
// environment. For example, if you tag an environment with this operation, that
// tag doesn't automatically propagate to the VLAN subnets and hosts associated
// with the environment.
