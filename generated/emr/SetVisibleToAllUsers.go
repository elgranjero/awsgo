package emr

// SetVisibleToAllUsers is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// The SetVisibleToAllUsers parameter is no longer supported. Your cluster may be
// visible to all users in your account. To restrict cluster access using an IAM
// policy, see [Identity and Access Management for Amazon EMR].
//
// Sets the Cluster$VisibleToAllUsers value for an Amazon EMR cluster. When true , IAM principals in the
// Amazon Web Services account can perform Amazon EMR cluster actions that their
// IAM policies allow. When false , only the IAM principal that created the cluster
// and the Amazon Web Services account root user can perform Amazon EMR actions on
// the cluster, regardless of IAM permissions policies attached to other IAM
// principals.
//
// This action works on running clusters. When you create a cluster, use the RunJobFlowInput$VisibleToAllUsers
// parameter.
//
// For more information, see [Understanding the Amazon EMR Cluster VisibleToAllUsers Setting] in the Amazon EMR Management Guide.
//
// [Understanding the Amazon EMR Cluster VisibleToAllUsers Setting]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/security_IAM_emr-with-IAM.html#security_set_visible_to_all_users
// [Identity and Access Management for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-access-IAM.html
