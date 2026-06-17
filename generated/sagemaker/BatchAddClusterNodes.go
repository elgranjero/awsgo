package sagemaker

// BatchAddClusterNodes is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Adds nodes to a HyperPod cluster by incrementing the target count for one or
// more instance groups. This operation returns a unique NodeLogicalId for each
// node being added, which can be used to track the provisioning status of the
// node. This API provides a safer alternative to UpdateCluster for scaling
// operations by avoiding unintended configuration changes.
//
// This API is only supported for clusters using Continuous as the
// NodeProvisioningMode .
