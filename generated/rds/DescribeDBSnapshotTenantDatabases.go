package rds

// DescribeDBSnapshotTenantDatabases is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Describes the tenant databases that exist in a DB snapshot. This command only
// applies to RDS for Oracle DB instances in the multi-tenant configuration.
//
// You can use this command to inspect the tenant databases within a snapshot
// before restoring it. You can't directly interact with the tenant databases in a
// DB snapshot. If you restore a snapshot that was taken from DB instance using the
// multi-tenant configuration, you restore all its tenant databases.
