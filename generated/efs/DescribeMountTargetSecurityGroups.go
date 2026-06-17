package efs

// DescribeMountTargetSecurityGroups is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Returns the security groups currently in effect for a mount target. This
// operation requires that the network interface of the mount target has been
// created and the lifecycle state of the mount target is not deleted .
//
// This operation requires permissions for the following actions:
//
// - elasticfilesystem:DescribeMountTargetSecurityGroups action on the mount
// target's file system.
//
// - ec2:DescribeNetworkInterfaceAttribute action on the mount target's network
// interface.
