package docdb

// DescribeDBClusterSnapshotAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Returns a list of cluster snapshot attribute names and values for a manual DB
// cluster snapshot.
//
// When you share snapshots with other Amazon Web Services accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the Amazon Web Services accounts that are authorized to copy or restore
// the manual cluster snapshot. If all is included in the list of values for the
// restore attribute, then the manual cluster snapshot is public and can be copied
// or restored by all Amazon Web Services accounts.
