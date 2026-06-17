package vpclattice

// CreateServiceNetworkServiceAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/vpclattice.go.
//
// Associates the specified service with the specified service network. For more
// information, see [Manage service associations]in the Amazon VPC Lattice User Guide.
//
// You can't use this operation if the service and service network are already
// associated or if there is a disassociation or deletion in progress. If the
// association fails, you can retry the operation by deleting the association and
// recreating it.
//
// You cannot associate a service and service network that are shared with a
// caller. The caller must own either the service or the service network.
//
// As a result of this operation, the association is created in the service
// network account and the association owner account.
//
// [Manage service associations]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-network-associations.html#service-network-service-associations
