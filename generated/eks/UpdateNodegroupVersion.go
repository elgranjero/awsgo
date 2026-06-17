package eks

// UpdateNodegroupVersion is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Updates the Kubernetes version or AMI version of an Amazon EKS managed node
// group.
//
// You can update a node group using a launch template only if the node group was
// originally deployed with a launch template. Additionally, the launch template ID
// or name must match what was used when the node group was created. You can update
// the launch template version with necessary changes.
//
// If you need to update a custom AMI in a node group that was deployed with a
// launch template, then update your custom AMI, specify the new ID in a new
// version of the launch template, and then update the node group to the new
// version of the launch template.
//
// If you update without a launch template, then you can update to the latest
// available AMI version of a node group's current Kubernetes version by not
// specifying a Kubernetes version in the request. You can update to the latest AMI
// version of your cluster's current Kubernetes version by specifying your
// cluster's Kubernetes version in the request. For information about Linux
// versions, see [Amazon EKS optimized Amazon Linux AMI versions]in the Amazon EKS User Guide. For information about Windows
// versions, see [Amazon EKS optimized Windows AMI versions]in the Amazon EKS User Guide.
//
// You cannot roll back a node group to an earlier Kubernetes version or AMI
// version.
//
// When a node in a managed node group is terminated due to a scaling action or
// update, every Pod on that node is drained first. Amazon EKS attempts to drain
// the nodes gracefully and will fail if it is unable to do so. You can force the
// update if Amazon EKS is unable to drain the nodes as a result of a Pod
// disruption budget issue.
//
// [Amazon EKS optimized Amazon Linux AMI versions]: https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html
// [Amazon EKS optimized Windows AMI versions]: https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-versions-windows.html
