package redshift

// CreateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Creates a new cluster with the specified parameters.
//
// To create a cluster in Virtual Private Cloud (VPC), you must provide a cluster
// subnet group name. The cluster subnet group identifies the subnets of your VPC
// that Amazon Redshift uses when creating the cluster. For more information about
// managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
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
