package s3tables

// PutTableReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3tables.go.
//
// Creates or updates the replication configuration for a specific table. This
// operation allows you to define table-level replication independently of
// bucket-level replication, providing granular control over which tables are
// replicated and where.
//
// Permissions
//
// - You must have the s3tables:PutTableReplication permission to use this
// operation. The IAM role specified in the configuration must have permissions to
// read from the source table and write to all destination tables.
//
// - You must also have the following permissions:
//
// - s3tables:GetTable permission on the source table being replicated.
//
// - s3tables:CreateTable permission for the destination.
//
// - s3tables:CreateNamespace permission for the destination.
//
// - s3tables:GetTableMaintenanceConfig permission for the source table.
//
// - s3tables:PutTableMaintenanceConfig permission for the destination table.
//
// - You must have iam:PassRole permission with condition allowing roles to be
// passed to replication.s3tables.amazonaws.com .
