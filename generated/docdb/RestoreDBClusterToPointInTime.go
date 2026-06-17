package docdb

// RestoreDBClusterToPointInTime is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Restores a cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target cluster is created from the source cluster with the same
// configuration as the original cluster, except that the new cluster is created
// with the default security group.
