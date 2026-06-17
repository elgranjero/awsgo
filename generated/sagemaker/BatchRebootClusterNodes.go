package sagemaker

// BatchRebootClusterNodes is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Reboots specific nodes within a SageMaker HyperPod cluster using a soft
// recovery mechanism. BatchRebootClusterNodes performs a graceful reboot of the
// specified nodes by calling the Amazon Elastic Compute Cloud RebootInstances
// API, which attempts to cleanly shut down the operating system before restarting
// the instance.
//
// This operation is useful for recovering from transient issues or applying
// certain configuration changes that require a restart.
//
// - Rebooting a node may cause temporary service interruption for workloads
// running on that node. Ensure your workloads can handle node restarts or use
// appropriate scheduling to minimize impact.
//
// - You can reboot up to 25 nodes in a single request.
//
// - For SageMaker HyperPod clusters using the Slurm workload manager, ensure
// rebooting nodes will not disrupt critical cluster operations.
