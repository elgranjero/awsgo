package s3tables

// PutTableBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3tables.go.
//
// Creates or updates the replication configuration for a table bucket. This
// operation defines how tables in the source bucket are replicated to destination
// buckets. Replication helps ensure data availability and disaster recovery across
// regions or accounts.
//
// Permissions
//
// - You must have the s3tables:PutTableBucketReplication permission to use this
// operation. The IAM role specified in the configuration must have permissions to
// read from the source bucket and write permissions to all destination buckets.
//
// - You must also have the following permissions:
//
// - s3tables:GetTable permission on the source table.
//
// - s3tables:ListTables permission on the bucket containing the table.
//
// - s3tables:CreateTable permission for the destination.
//
// - s3tables:CreateNamespace permission for the destination.
//
// - s3tables:GetTableMaintenanceConfig permission for the source bucket.
//
// - s3tables:PutTableMaintenanceConfig permission for the destination bucket.
//
// - You must have iam:PassRole permission with condition allowing roles to be
// passed to replication.s3tables.amazonaws.com .
