package rds

// RestoreDBInstanceFromDBSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates a new DB instance from a DB snapshot. The target database is created
// from the source database restore point with most of the source's original
// configuration, including the default security group and DB parameter group. By
// default, the new DB instance is created as a Single-AZ deployment, except when
// the instance is a SQL Server instance that has an option group associated with
// mirroring. In this case, the instance becomes a Multi-AZ deployment, not a
// Single-AZ deployment.
//
// If you want to replace your original DB instance with the new, restored DB
// instance, then rename your original DB instance before you call the
// RestoreDBInstanceFromDBSnapshot operation. RDS doesn't allow two DB instances
// with the same name. After you have renamed your original DB instance with a
// different identifier, then you can pass the original name of the DB instance as
// the DBInstanceIdentifier in the call to the RestoreDBInstanceFromDBSnapshot
// operation. The result is that you replace the original DB instance with the DB
// instance created from the snapshot.
//
// If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier
// must be the ARN of the shared DB snapshot.
//
// To restore from a DB snapshot with an unsupported engine version, you must
// first upgrade the engine version of the snapshot. For more information about
// upgrading a RDS for MySQL DB snapshot engine version, see [Upgrading a MySQL DB snapshot engine version]. For more
// information about upgrading a RDS for PostgreSQL DB snapshot engine version, [Upgrading a PostgreSQL DB snapshot engine version].
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterFromSnapshot .
//
// [Upgrading a PostgreSQL DB snapshot engine version]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBSnapshot.PostgreSQL.html
// [Upgrading a MySQL DB snapshot engine version]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-upgrade-snapshot.html
