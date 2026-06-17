package eks

// RegisterCluster is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Connects a Kubernetes cluster to the Amazon EKS control plane.
//
// Any Kubernetes cluster can be connected to the Amazon EKS control plane to view
// current information about the cluster and its nodes.
//
// Cluster connection requires two steps. First, send a [RegisterClusterRequest]RegisterClusterRequest to
// add it to the Amazon EKS control plane.
//
// Second, a [Manifest] containing the activationID and activationCode must be applied to
// the Kubernetes cluster through it's native provider to provide visibility.
//
// After the manifest is updated and applied, the connected cluster is visible to
// the Amazon EKS control plane. If the manifest isn't applied within three days,
// the connected cluster will no longer be visible and must be deregistered using
// DeregisterCluster .
//
// [RegisterClusterRequest]: https://docs.aws.amazon.com/eks/latest/APIReference/API_RegisterClusterRequest.html
// [Manifest]: https://amazon-eks.s3.us-west-2.amazonaws.com/eks-connector/manifests/eks-connector/latest/eks-connector.yaml
