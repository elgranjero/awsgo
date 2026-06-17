package neptune

// DescribeDBClusterSnapshotAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Returns a list of DB cluster snapshot attribute names and values for a manual
// DB cluster snapshot.
//
// When sharing snapshots with other Amazon accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the Amazon accounts that are authorized to copy or restore the manual DB
// cluster snapshot. If all is included in the list of values for the restore
// attribute, then the manual DB cluster snapshot is public and can be copied or
// restored by all Amazon accounts.
//
// To add or remove access for an Amazon account to copy or restore a manual DB
// cluster snapshot, or to make the manual DB cluster snapshot public or private,
// use the ModifyDBClusterSnapshotAttributeAPI action.
