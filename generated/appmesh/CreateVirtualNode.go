package appmesh

// CreateVirtualNode is generated as a reference stub.
// Executable command wiring lives under cmd/appmesh.go.
//
// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such as an
// Amazon ECS service or a Kubernetes deployment. When you create a virtual node,
// you can specify the service discovery information for your task group, and
// whether the proxy running in a task group will communicate with other proxies
// using Transport Layer Security (TLS).
//
// You define a listener for any inbound traffic that your virtual node expects.
// Any virtual service that your virtual node expects to communicate to is
// specified as a backend .
//
// The response metadata for your new virtual node contains the arn that is
// associated with the virtual node. Set this value to the full ARN; for example,
// arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp ) as the
// APPMESH_RESOURCE_ARN environment variable for your task group's Envoy proxy
// container in your task definition or pod spec. This is then mapped to the
// node.id and node.cluster Envoy parameters.
//
// By default, App Mesh uses the name of the resource you specified in
// APPMESH_RESOURCE_ARN when Envoy is referring to itself in metrics and traces.
// You can override this behavior by setting the APPMESH_RESOURCE_CLUSTER
// environment variable with your own name.
//
// For more information about virtual nodes, see [Virtual nodes]. You must be using 1.15.0 or
// later of the Envoy image when setting these variables. For more information
// aboutApp Mesh Envoy variables, see [Envoy image]in the App Mesh User Guide.
//
// [Virtual nodes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html
// [Envoy image]: https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html
