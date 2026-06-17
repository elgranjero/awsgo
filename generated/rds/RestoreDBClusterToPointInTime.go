package rds

// RestoreDBClusterToPointInTime is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group. Unless the RestoreType is set to
// copy-on-write , the restore may occur in a different Availability Zone (AZ) from
// the original DB cluster. The AZ where RDS restores the DB cluster depends on the
// AZs in the specified subnet group.
//
// For Aurora, this operation only restores the DB cluster, not the DB instances
// for that DB cluster. You must invoke the CreateDBInstance operation to create
// DB instances for the restored DB cluster, specifying the identifier of the
// restored DB cluster in DBClusterIdentifier . You can create DB instances only
// after the RestoreDBClusterToPointInTime operation has completed and the DB
// cluster is available.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
