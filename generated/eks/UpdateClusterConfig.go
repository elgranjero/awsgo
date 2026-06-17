package eks

// UpdateClusterConfig is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Updates an Amazon EKS cluster configuration. Your cluster continues to function
// during the update. The response output includes an update ID that you can use to
// track the status of your cluster update with DescribeUpdate .
//
// You can use this operation to do the following actions:
//
// - You can use this API operation to enable or disable exporting the
// Kubernetes control plane logs for your cluster to CloudWatch Logs. By default,
// cluster control plane logs aren't exported to CloudWatch Logs. For more
// information, see [Amazon EKS Cluster control plane logs]in the Amazon EKS User Guide .
//
// CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
//
// exported control plane logs. For more information, see [CloudWatch Pricing].
//
// - You can also use this API operation to enable or disable public and private
// access to your cluster's Kubernetes API server endpoint. By default, public
// access is enabled, and private access is disabled. For more information, see [Cluster API server endpoint]
// in the Amazon EKS User Guide .
//
// - You can also use this API operation to choose different subnets and
// security groups for the cluster. You must specify at least two subnets that are
// in different Availability Zones. You can't change which VPC the subnets are
// from, the subnets must be in the same VPC as the subnets that the cluster was
// created with. For more information about the VPC requirements, see [https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html]in the
// Amazon EKS User Guide .
//
// - You can also use this API operation to enable or disable ARC zonal shift.
// If zonal shift is enabled, Amazon Web Services configures zonal autoshift for
// the cluster.
//
// - You can also use this API operation to add, change, or remove the
// configuration in the cluster for EKS Hybrid Nodes. To remove the configuration,
// use the remoteNetworkConfig key with an object containing both subkeys with
// empty arrays for each. Here is an inline example: "remoteNetworkConfig": {
// "remoteNodeNetworks": [], "remotePodNetworks": [] } .
//
// Cluster updates are asynchronous, and they should finish within a few minutes.
// During an update, the cluster status moves to UPDATING (this status transition
// is eventually consistent). When the update is complete (either Failed or
// Successful ), the cluster status moves to Active .
//
// [Amazon EKS Cluster control plane logs]: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
//
// [Cluster API server endpoint]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
// [CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
// [https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html]: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html
