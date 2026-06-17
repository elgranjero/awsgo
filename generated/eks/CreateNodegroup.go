package eks

// CreateNodegroup is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Creates a managed node group for an Amazon EKS cluster.
//
// You can only create a node group for your cluster that is equal to the current
// Kubernetes version for the cluster. All node groups are created with the latest
// AMI release version for the respective minor Kubernetes version of the cluster,
// unless you deploy a custom AMI using a launch template.
//
// For later updates, you will only be able to update a node group using a launch
// template only if it was originally deployed with a launch template.
// Additionally, the launch template ID or name must match what was used when the
// node group was created. You can update the launch template version with
// necessary changes. For more information about using launch templates, see [Customizing managed nodes with launch templates].
//
// An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and
// associated Amazon EC2 instances that are managed by Amazon Web Services for an
// Amazon EKS cluster. For more information, see [Managed node groups]in the Amazon EKS User Guide.
//
// Windows AMI types are only supported for commercial Amazon Web Services Regions
// that support Windows on Amazon EKS.
//
// [Customizing managed nodes with launch templates]: https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
// [Managed node groups]: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
