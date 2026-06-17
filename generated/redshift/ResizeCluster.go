package redshift

// ResizeCluster is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Changes the size of the cluster. You can change the cluster's type, or change
// the number or type of nodes. The default behavior is to use the elastic resize
// method. With an elastic resize, your cluster is available for read and write
// operations more quickly than with the classic resize method.
//
// Elastic resize operations have the following restrictions:
//
// - You can only resize clusters of the following types:
//
// - dc2.large
//
// - dc2.8xlarge
//
// - ra3.large
//
// - ra3.xlplus
//
// - ra3.4xlarge
//
// - ra3.16xlarge
//
// - The type of nodes that you add must match the node type for the cluster.
