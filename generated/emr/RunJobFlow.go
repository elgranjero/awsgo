package emr

// RunJobFlow is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// RunJobFlow creates and starts running a new cluster (job flow). The cluster
// runs the steps specified. After the steps complete, the cluster stops and the
// HDFS partition is lost. To prevent loss of data, configure the last step of the
// job flow to store results in Amazon S3. If the JobFlowInstancesConfigKeepJobFlowAliveWhenNoSteps
// parameter is set to TRUE , the cluster transitions to the WAITING state rather
// than shutting down after the steps have completed.
//
// For additional protection, you can set the JobFlowInstancesConfigTerminationProtected parameter to
// TRUE to lock the cluster and prevent it from being terminated by API call, user
// intervention, or in the event of a job flow error.
//
// A maximum of 256 steps are allowed in each job flow.
//
// If your cluster is long-running (such as a Hive data warehouse) or complex, you
// may require more than 256 steps to process your data. You can bypass the
// 256-step limitation in various ways, including using the SSH shell to connect to
// the master node and submitting queries directly to the software running on the
// master node, such as Hive and Hadoop.
//
// For long-running clusters, we recommend that you periodically store your
// results.
//
// The instance fleets configuration is available only in Amazon EMR releases
// 4.8.0 and later, excluding 5.0.x versions. The RunJobFlow request can contain
// InstanceFleets parameters or InstanceGroups parameters, but not both.
