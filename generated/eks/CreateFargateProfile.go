package eks

// CreateFargateProfile is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Creates an Fargate profile for your Amazon EKS cluster. You must have at least
// one Fargate profile in a cluster to be able to run pods on Fargate.
//
// The Fargate profile allows an administrator to declare which pods run on
// Fargate and specify which pods run on which Fargate profile. This declaration is
// done through the profile's selectors. Each profile can have up to five selectors
// that contain a namespace and labels. A namespace is required for every selector.
// The label field consists of multiple optional key-value pairs. Pods that match
// the selectors are scheduled on Fargate. If a to-be-scheduled pod matches any of
// the selectors in the Fargate profile, then that pod is run on Fargate.
//
// When you create a Fargate profile, you must specify a pod execution role to use
// with the pods that are scheduled with the profile. This role is added to the
// cluster's Kubernetes [Role Based Access Control](RBAC) for authorization so that the kubelet that is
// running on the Fargate infrastructure can register with your Amazon EKS cluster
// so that it can appear in your cluster as a node. The pod execution role also
// provides IAM permissions to the Fargate infrastructure to allow read access to
// Amazon ECR image repositories. For more information, see [Pod Execution Role]in the Amazon EKS User
// Guide.
//
// Fargate profiles are immutable. However, you can create a new updated profile
// to replace an existing profile and then delete the original after the updated
// profile has finished creating.
//
// If any Fargate profiles in a cluster are in the DELETING status, you must wait
// for that Fargate profile to finish deleting before you can create any other
// profiles in that cluster.
//
// For more information, see [Fargate profile] in the Amazon EKS User Guide.
//
// [Role Based Access Control]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
// [Fargate profile]: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
// [Pod Execution Role]: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
