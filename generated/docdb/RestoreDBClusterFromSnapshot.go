package docdb

// RestoreDBClusterFromSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Creates a new cluster from a snapshot or cluster snapshot.
//
// If a snapshot is specified, the target cluster is created from the source DB
// snapshot with a default configuration and default security group.
//
// If a cluster snapshot is specified, the target cluster is created from the
// source cluster restore point with the same configuration as the original source
// DB cluster, except that the new cluster is created with the default security
// group.
