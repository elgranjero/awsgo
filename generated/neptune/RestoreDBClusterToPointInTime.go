package neptune

// RestoreDBClusterToPointInTime is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group.
//
// This action only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstanceaction to create DB instances for the restored DB
// cluster, specifying the identifier of the restored DB cluster in
// DBClusterIdentifier . You can create DB instances only after the
// RestoreDBClusterToPointInTime action has completed and the DB cluster is
// available.
