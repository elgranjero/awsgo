package s3

// DeleteBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Deletes the replication configuration from the bucket.
//
// To use this operation, you must have permissions to perform the
// s3:PutReplicationConfiguration action. The bucket owner has these permissions by
// default and can grant it to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations]
// and [Managing Access Permissions to Your Amazon S3 Resources].
//
// It can take a while for the deletion of a replication configuration to fully
// propagate.
//
// For information about replication configuration, see [Replication] in the Amazon S3 User
// Guide.
//
// The following operations are related to DeleteBucketReplication :
//
// [PutBucketReplication]
//
// [GetBucketReplication]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html
// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
