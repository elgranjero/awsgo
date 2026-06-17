package rds

// RestoreDBInstanceToPointInTime is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Restores a DB instance to an arbitrary point in time. You can restore to any
// point in time before the time identified by the LatestRestorableTime property.
// You can restore to a point up to the number of days specified by the
// BackupRetentionPeriod property.
//
// The target database is created with most of the original configuration, but in
// a system-selected Availability Zone, with the default security group, the
// default subnet group, and the default DB parameter group. By default, the new DB
// instance is created as a single-AZ deployment except when the instance is a SQL
// Server instance that has an option group that is associated with mirroring; in
// this case, the instance becomes a mirrored deployment and not a single-AZ
// deployment.
//
// This operation doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterToPointInTime .
