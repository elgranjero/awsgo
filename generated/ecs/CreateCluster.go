package ecs

// CreateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Creates a new Amazon ECS cluster. By default, your account receives a default
// cluster when you launch your first container instance. However, you can create
// your own cluster with a unique name.
//
// When you call the [CreateCluster] API operation, Amazon ECS attempts to create the Amazon ECS
// service-linked role for your account. This is so that it can manage required
// resources in other Amazon Web Services services on your behalf. However, if the
// user that makes the call doesn't have permissions to create the service-linked
// role, it isn't created. For more information, see [Using service-linked roles for Amazon ECS]in the Amazon Elastic
// Container Service Developer Guide.
//
// [Using service-linked roles for Amazon ECS]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html
// [CreateCluster]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCluster.html
