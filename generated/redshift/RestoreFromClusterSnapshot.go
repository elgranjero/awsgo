package redshift

// RestoreFromClusterSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Creates a new cluster from a snapshot. By default, Amazon Redshift creates the
// resulting cluster with the same configuration as the original cluster from which
// the snapshot was created, except that the new cluster is created with the
// default cluster security and parameter groups. After Amazon Redshift creates the
// cluster, you can use the ModifyClusterAPI to associate a different security group and
// different parameter group with the restored cluster. If you are using a DS node
// type, you can also choose to change to another DS node type of the same size
// during restore.
//
// If you restore a cluster into a VPC, you must provide a cluster subnet group
// where you want the cluster restored.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// subnet group for a provisioned cluster is in an account with VPC BPA turned on,
// the following capabilities are blocked:
//
// - Creating a public cluster
//
// - Restoring a public cluster
//
// - Modifying a private cluster to be public
//
// - Adding a subnet with VPC BPA turned on to the subnet group when there's at
// least one public cluster within the group
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
