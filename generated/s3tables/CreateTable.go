package s3tables

// CreateTable is generated as a reference stub.
// Executable command wiring lives under cmd/s3tables.go.
//
// Creates a new table associated with the given namespace in a table bucket. For
// more information, see [Creating an Amazon S3 table]in the Amazon Simple Storage Service User Guide.
//
// Permissions
//
// - You must have the s3tables:CreateTable permission to use this operation.
//
// - If you use this operation with the optional metadata request parameter you
// must have the s3tables:PutTableData permission.
//
// - If you use this operation with the optional encryptionConfiguration request
// parameter you must have the s3tables:PutTableEncryption permission.
//
// - If you use this operation with the storageClassConfiguration request
// parameter, you must have the s3tables:PutTableStorageClass permission.
//
// - To create a table with tags, you must have the s3tables:TagResource
// permission in addition to s3tables:CreateTable permission.
//
// Additionally, If you choose SSE-KMS encryption you must grant the S3 Tables
// maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption].
//
// [Creating an Amazon S3 table]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-create.html
// [Permissions requirements for S3 Tables SSE-KMS encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-kms-permissions.html
