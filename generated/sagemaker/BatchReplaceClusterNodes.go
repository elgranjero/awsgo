package sagemaker

// BatchReplaceClusterNodes is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Replaces specific nodes within a SageMaker HyperPod cluster with new hardware.
// BatchReplaceClusterNodes terminates the specified instances and provisions new
// replacement instances with the same configuration but fresh hardware. The Amazon
// Machine Image (AMI) and instance configuration remain the same.
//
// This operation is useful for recovering from hardware failures or persistent
// issues that cannot be resolved through a reboot.
//
// - Data Loss Warning: Replacing nodes destroys all instance volumes, including
// both root and secondary volumes. All data stored on these volumes will be
// permanently lost and cannot be recovered.
//
// - To safeguard your work, back up your data to Amazon S3 or an FSx for Lustre
// file system before invoking the API on a worker node group. This will help
// prevent any potential data loss from the instance root volume. For more
// information about backup, see [Use the backup script provided by SageMaker HyperPod].
//
// - If you want to invoke this API on an existing cluster, you'll first need to
// patch the cluster by running the [UpdateClusterSoftware API]. For more information about patching a
// cluster, see [Update the SageMaker HyperPod platform software of a cluster].
//
// - You can replace up to 25 nodes in a single request.
//
// [UpdateClusterSoftware API]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateClusterSoftware.html
// [Use the backup script provided by SageMaker HyperPod]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate-cli-command.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software-backup
// [Update the SageMaker HyperPod platform software of a cluster]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate-cli-command.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software
