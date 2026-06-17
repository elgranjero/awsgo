package pcs

// CreateComputeNodeGroup is generated as a reference stub.
// Executable command wiring lives under cmd/pcs.go.
//
// Creates a managed set of compute nodes. You associate a compute node group with
// a cluster through 1 or more PCS queues or as part of the login fleet. A compute
// node group includes the definition of the compute properties and lifecycle
// management. PCS uses the information you provide to this API action to launch
// compute nodes in your account. You can only specify subnets in the same Amazon
// VPC as your cluster. You receive billing charges for the compute nodes that PCS
// launches in your account. You must already have a launch template before you
// call this API. For more information, see [Launch an instance from a launch template]in the Amazon Elastic Compute Cloud
// User Guide for Linux Instances.
//
// [Launch an instance from a launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
