package eks

// DeleteCluster is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Deletes an Amazon EKS cluster control plane.
//
// If you have active services and ingress resources in your cluster that are
// associated with a load balancer, you must delete those services before deleting
// the cluster so that the load balancers are deleted properly. Otherwise, you can
// have orphaned resources in your VPC that prevent you from being able to delete
// the VPC. For more information, see [Deleting a cluster]in the Amazon EKS User Guide.
//
// If you have managed node groups or Fargate profiles attached to the cluster,
// you must delete them first. For more information, see DeleteNodgroup and
// DeleteFargateProfile .
//
// [Deleting a cluster]: https://docs.aws.amazon.com/eks/latest/userguide/delete-cluster.html
