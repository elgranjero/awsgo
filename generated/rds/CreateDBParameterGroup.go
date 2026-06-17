package rds

// CreateDBParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates a new DB parameter group.
//
// A DB parameter group is initially created with the default parameters for the
// database engine used by the DB instance. To provide custom values for any of the
// parameters, you must modify the group after creating it using
// ModifyDBParameterGroup . Once you've created a DB parameter group, you need to
// associate it with your DB instance using ModifyDBInstance . When you associate a
// new DB parameter group with a running DB instance, you need to reboot the DB
// instance without failover for the new DB parameter group and associated settings
// to take effect.
//
// This command doesn't apply to RDS Custom.
