package efs

// ModifyMountTargetSecurityGroups is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Modifies the set of security groups in effect for a mount target.
//
// When you create a mount target, Amazon EFS also creates a new network
// interface. For more information, see CreateMountTarget. This operation replaces the security
// groups in effect for the network interface associated with a mount target, with
// the SecurityGroups provided in the request. This operation requires that the
// network interface of the mount target has been created and the lifecycle state
// of the mount target is not deleted .
//
// The operation requires permissions for the following actions:
//
// - elasticfilesystem:ModifyMountTargetSecurityGroups action on the mount
// target's file system.
//
// - ec2:ModifyNetworkInterfaceAttribute action on the mount target's network
// interface.
