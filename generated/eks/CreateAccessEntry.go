package eks

// CreateAccessEntry is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Creates an access entry.
//
// An access entry allows an IAM principal to access your cluster. Access entries
// can replace the need to maintain entries in the aws-auth ConfigMap for
// authentication. You have the following options for authorizing an IAM principal
// to access Kubernetes objects on your cluster: Kubernetes role-based access
// control (RBAC), Amazon EKS, or both. Kubernetes RBAC authorization requires you
// to create and manage Kubernetes Role , ClusterRole , RoleBinding , and
// ClusterRoleBinding objects, in addition to managing access entries. If you use
// Amazon EKS authorization exclusively, you don't need to create and manage
// Kubernetes Role , ClusterRole , RoleBinding , and ClusterRoleBinding objects.
//
// For more information about access entries, see [Access entries] in the Amazon EKS User Guide.
//
// [Access entries]: https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html
