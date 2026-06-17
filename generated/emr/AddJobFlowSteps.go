package emr

// AddJobFlowSteps is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// AddJobFlowSteps adds new steps to a running cluster. A maximum of 256 steps are
// allowed in each job flow.
//
// If your cluster is long-running (such as a Hive data warehouse) or complex, you
// may require more than 256 steps to process your data. You can bypass the
// 256-step limitation in various ways, including using SSH to connect to the
// master node and submitting queries directly to the software running on the
// master node, such as Hive and Hadoop.
//
// A step specifies the location of a JAR file stored either on the master node of
// the cluster or in Amazon S3. Each step is performed by the main function of the
// main class of the JAR file. The main class can be specified either in the
// manifest of the JAR or by using the MainFunction parameter of the step.
//
// Amazon EMR executes each step in the order listed. For a step to be considered
// complete, the main function must exit with a zero exit code and all Hadoop jobs
// started while the step was running must have completed and run successfully.
//
// You can only add steps to a cluster that is in one of the following states:
// STARTING, BOOTSTRAPPING, RUNNING, or WAITING.
//
// The string values passed into HadoopJarStep object cannot exceed a total of
// 10240 characters.
