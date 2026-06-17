package redshift

// ModifyCluster is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Modifies the settings for a cluster.
//
// You can also change node type and the number of nodes to scale up or down the
// cluster. When resizing a cluster, you must specify both the number of nodes and
// the node type even if one of the parameters does not change.
//
// You can add another security or parameter group, or change the admin user
// password. Resetting a cluster password or modifying the security groups
// associated with a cluster do not need a reboot. However, modifying a parameter
// group requires a reboot for parameters to take effect. For more information
// about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
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
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
