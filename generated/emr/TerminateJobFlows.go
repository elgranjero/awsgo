package emr

// TerminateJobFlows is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// TerminateJobFlows shuts a list of clusters (job flows) down. When a job flow is
// shut down, any step not yet completed is canceled and the Amazon EC2 instances
// on which the cluster is running are stopped. Any log files not already saved are
// uploaded to Amazon S3 if a LogUri was specified when the cluster was created.
//
// The maximum number of clusters allowed is 10. The call to TerminateJobFlows is
// asynchronous. Depending on the configuration of the cluster, it may take up to
// 1-5 minutes for the cluster to completely terminate and release allocated
// resources, such as Amazon EC2 instances.
