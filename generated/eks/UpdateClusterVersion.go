package eks

// UpdateClusterVersion is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Updates an Amazon EKS cluster to the specified Kubernetes version. Your cluster
// continues to function during the update. The response output includes an update
// ID that you can use to track the status of your cluster update with the [DescribeUpdate]
// DescribeUpdate API operation.
//
// Cluster updates are asynchronous, and they should finish within a few minutes.
// During an update, the cluster status moves to UPDATING (this status transition
// is eventually consistent). When the update is complete (either Failed or
// Successful ), the cluster status moves to Active .
//
// If your cluster has managed node groups attached to it, all of your node
// groups' Kubernetes versions must match the cluster's Kubernetes version in order
// to update the cluster to a new Kubernetes version.
//
// [DescribeUpdate]: https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeUpdate.html
