package ssm

// DescribeInstanceInformation is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Provides information about one or more of your managed nodes, including the
// operating system platform, SSM Agent version, association status, and IP
// address. This operation does not return information for nodes that are either
// Stopped or Terminated.
//
// If you specify one or more node IDs, the operation returns information for
// those managed nodes. If you don't specify node IDs, it returns information for
// all your managed nodes. If you specify a node ID that isn't valid or a node that
// you don't own, you receive an error.
//
// The IamRole field returned for this API operation is the role assigned to an
// Amazon EC2 instance configured with a Systems Manager Quick Setup host
// management configuration or the role assigned to an on-premises managed node.
