package emr

// SetTerminationProtection is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// SetTerminationProtection locks a cluster (job flow) so the Amazon EC2 instances
// in the cluster cannot be terminated by user intervention, an API call, or in the
// event of a job-flow error. The cluster still terminates upon successful
// completion of the job flow. Calling SetTerminationProtection on a cluster is
// similar to calling the Amazon EC2 DisableAPITermination API on all Amazon EC2
// instances in a cluster.
//
// SetTerminationProtection is used to prevent accidental termination of a cluster
// and to ensure that in the event of an error, the instances persist so that you
// can recover any data stored in their ephemeral instance storage.
//
// To terminate a cluster that has been locked by setting SetTerminationProtection
// to true , you must first unlock the job flow by a subsequent call to
// SetTerminationProtection in which you set the value to false .
//
// For more information, see [Managing Cluster Termination] in the Amazon EMR Management Guide.
//
// [Managing Cluster Termination]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html
