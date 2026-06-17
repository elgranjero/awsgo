package elasticache

// CopySnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Makes a copy of an existing snapshot.
//
// This operation is valid for Valkey or Redis OSS only.
//
// Users or groups that have permissions to use the CopySnapshot operation can
// create their own Amazon S3 buckets and copy snapshots to it. To control access
// to your snapshots, use an IAM policy to control who has the ability to use the
// CopySnapshot operation. For more information about using IAM to control the use
// of ElastiCache operations, see [Exporting Snapshots]and [Authentication & Access Control].
//
// You could receive the following error messages.
//
// Error Messages
//
// - Error Message: The S3 bucket %s is outside of the region.
//
// Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//
// more information, see [Step 1: Create an Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: The S3 bucket %s does not exist.
//
// Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//
// more information, see [Step 1: Create an Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: The S3 bucket %s is not owned by the authenticated user.
//
// Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//
// more information, see [Step 1: Create an Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: The authenticated user does not have sufficient permissions
// to perform the desired activity.
//
// Solution: Contact your system administrator to get the needed permissions.
//
// - Error Message: The S3 bucket %s already contains an object with key %s.
//
// Solution: Give the TargetSnapshotName a new and unique value. If exporting a
//
// snapshot, you could alternatively create a new Amazon S3 bucket and use this
// same value for TargetSnapshotName .
//
// - Error Message: ElastiCache has not been granted READ permissions %s on the
// S3 Bucket.
//
// Solution: Add List and Read permissions on the bucket. For more information,
//
// see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: ElastiCache has not been granted WRITE permissions %s on the
// S3 Bucket.
//
// Solution: Add Upload/Delete permissions on the bucket. For more information,
//
// see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]in the ElastiCache User Guide.
//
// - Error Message: ElastiCache has not been granted READ_ACP permissions %s on
// the S3 Bucket.
//
// Solution: Add View Permissions on the bucket. For more information, see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]in the
//
// ElastiCache User Guide.
//
// [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-exporting.html#backups-exporting-grant-access
// [Exporting Snapshots]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-exporting.html
// [Authentication & Access Control]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/IAM.html
//
// [Step 1: Create an Amazon S3 Bucket]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-exporting.html#backups-exporting-create-s3-bucket
