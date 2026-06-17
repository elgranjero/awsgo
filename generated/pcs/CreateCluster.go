package pcs

// CreateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/pcs.go.
//
// Creates a cluster in your account. PCS creates the cluster controller in a
// service-owned account. The cluster controller communicates with the cluster
// resources in your account. The subnets and security groups for the cluster must
// already exist before you use this API action.
//
// It takes time for PCS to create the cluster. The cluster is in a Creating state
// until it is ready to use. There can only be 1 cluster in a Creating state per
// Amazon Web Services Region per Amazon Web Services account. CreateCluster fails
// with a ServiceQuotaExceededException if there is already a cluster in a Creating
// state.
