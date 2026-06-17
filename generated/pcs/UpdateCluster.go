package pcs

// UpdateCluster is generated as a reference stub.
// Executable command wiring lives under cmd/pcs.go.
//
// Updates a cluster configuration. You can modify Slurm scheduler settings,
// accounting configuration, and security groups for an existing cluster.
//
// You can only update clusters that are in ACTIVE , UPDATE_FAILED , or SUSPENDED
// state. All associated resources (queues and compute node groups) must be in
// ACTIVE state before you can update the cluster.
