package s3tables

// DeleteTableBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3tables.go.
//
// Deletes the replication configuration for a table bucket. After deletion, new
// table updates will no longer be replicated to destination buckets, though
// existing replicated tables will remain in destination buckets.
//
// Permissions You must have the s3tables:DeleteTableBucketReplication permission
// to use this operation.
