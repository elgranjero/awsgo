package vpclattice

// CreateServiceNetworkVpcAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/vpclattice.go.
//
// Associates a VPC with a service network. When you associate a VPC with the
// service network, it enables all the resources within that VPC to be clients and
// communicate with other services in the service network. For more information,
// see [Manage VPC associations]in the Amazon VPC Lattice User Guide.
//
// You can't use this operation if there is a disassociation in progress. If the
// association fails, retry by deleting the association and recreating it.
//
// As a result of this operation, the association gets created in the service
// network account and the VPC owner account.
//
// If you add a security group to the service network and VPC association, the
// association must continue to always have at least one security group. You can
// add or edit security groups at any time. However, to remove all security groups,
// you must first delete the association and recreate it without security groups.
//
// [Manage VPC associations]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-network-associations.html#service-network-vpc-associations
