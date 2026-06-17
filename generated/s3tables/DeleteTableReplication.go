package s3tables

// DeleteTableReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3tables.go.
//
// Deletes the replication configuration for a specific table. After deletion, new
// updates to this table will no longer be replicated to destination tables, though
// existing replicated copies will remain in destination buckets.
//
// Permissions You must have the s3tables:DeleteTableReplication permission to use
// this operation.
