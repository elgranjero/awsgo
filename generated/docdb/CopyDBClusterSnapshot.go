package docdb

// CopyDBClusterSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Copies a snapshot of a cluster.
//
// To copy a cluster snapshot from a shared manual cluster snapshot,
// SourceDBClusterSnapshotIdentifier must be the Amazon Resource Name (ARN) of the
// shared cluster snapshot. You can only copy a shared DB cluster snapshot, whether
// encrypted or not, in the same Amazon Web Services Region.
//
// To cancel the copy operation after it is in progress, delete the target cluster
// snapshot identified by TargetDBClusterSnapshotIdentifier while that cluster
// snapshot is in the copying status.
