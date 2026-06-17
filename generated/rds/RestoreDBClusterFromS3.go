package rds

// RestoreDBClusterFromS3 is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates an Amazon Aurora DB cluster from MySQL data stored in an Amazon S3
// bucket. Amazon RDS must be authorized to access the Amazon S3 bucket and the
// data must be created using the Percona XtraBackup utility as described in [Migrating Data from MySQL by Using an Amazon S3 Bucket]in
// the Amazon Aurora User Guide.
//
// This operation only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance operation to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier . You can create DB instances only after the
// RestoreDBClusterFromS3 operation has completed and the DB cluster is available.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// This operation only applies to Aurora DB clusters. The source DB engine must be
// MySQL.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Migrating Data from MySQL by Using an Amazon S3 Bucket]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Migrating.ExtMySQL.html#AuroraMySQL.Migrating.ExtMySQL.S3
