package rds

// RestoreDBClusterFromSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
//
// The target DB cluster is created from the source snapshot with a default
// configuration. If you don't specify a security group, the new DB cluster is
// associated with the default security group.
//
// This operation only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance operation to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier . You can create DB instances only after the
// RestoreDBClusterFromSnapshot operation has completed and the DB cluster is
// available.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
