package rds

// CreateDBInstanceReadReplica is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates a new DB instance that acts as a read replica for an existing source DB
// instance or Multi-AZ DB cluster. You can create a read replica for a DB instance
// running Db2, MariaDB, MySQL, Oracle, PostgreSQL, or SQL Server. You can create a
// read replica for a Multi-AZ DB cluster running MySQL or PostgreSQL. For more
// information, see [Working with read replicas]and [Migrating from a Multi-AZ DB cluster to a DB instance using a read replica] in the Amazon RDS User Guide.
//
// Amazon Aurora doesn't support this operation. To create a DB instance for an
// Aurora DB cluster, use the CreateDBInstance operation.
//
// RDS creates read replicas with backups disabled. All other attributes
// (including DB security groups and DB parameter groups) are inherited from the
// source DB instance or cluster, except as specified.
//
// Your source DB instance or cluster must have backup retention enabled.
//
// [Working with read replicas]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html
// [Migrating from a Multi-AZ DB cluster to a DB instance using a read replica]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html#multi-az-db-clusters-migrating-to-instance-with-read-replica
