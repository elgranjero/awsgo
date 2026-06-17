package eks

// CreateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Creates an Amazon EKS control plane.
//
// The Amazon EKS control plane consists of control plane instances that run the
// Kubernetes software, such as etcd and the API server. The control plane runs in
// an account managed by Amazon Web Services, and the Kubernetes API is exposed by
// the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane is
// single tenant and unique. It runs on its own set of Amazon EC2 instances.
//
// The cluster control plane is provisioned across multiple Availability Zones and
// fronted by an Elastic Load Balancing Network Load Balancer. Amazon EKS also
// provisions elastic network interfaces in your VPC subnets to provide
// connectivity from the control plane instances to the nodes (for example, to
// support kubectl exec , logs , and proxy data flows).
//
// Amazon EKS nodes run in your Amazon Web Services account and connect to your
// cluster's control plane over the Kubernetes API server endpoint and a
// certificate file that is created for your cluster.
//
// You can use the endpointPublicAccess and endpointPrivateAccess parameters to
// enable or disable public and private access to your cluster's Kubernetes API
// server endpoint. By default, public access is enabled, and private access is
// disabled. The endpoint domain name and IP address family depends on the value of
// the ipFamily for the cluster. For more information, see [Amazon EKS Cluster Endpoint Access Control] in the Amazon EKS User
// Guide .
//
// You can use the logging parameter to enable or disable exporting the Kubernetes
// control plane logs for your cluster to CloudWatch Logs. By default, cluster
// control plane logs aren't exported to CloudWatch Logs. For more information, see
// [Amazon EKS Cluster Control Plane Logs]in the Amazon EKS User Guide .
//
// CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
// exported control plane logs. For more information, see [CloudWatch Pricing].
//
// In most cases, it takes several minutes to create a cluster. After you create
// an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate
// with the API server and launch nodes into your cluster. For more information,
// see [Allowing users to access your cluster]and [Launching Amazon EKS nodes] in the Amazon EKS User Guide.
//
// [Allowing users to access your cluster]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-auth.html
// [CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
// [Amazon EKS Cluster Control Plane Logs]: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
// [Amazon EKS Cluster Endpoint Access Control]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
// [Launching Amazon EKS nodes]: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
