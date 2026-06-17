package s3

// GetBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the replication configuration of a bucket.
//
// It can take a while to propagate the put or delete a replication configuration
// to all Amazon S3 systems. Therefore, a get request soon after put or delete can
// return a wrong result.
//
// For information about replication configuration, see [Replication] in the Amazon S3 User
// Guide.
//
// This action requires permissions for the s3:GetReplicationConfiguration action.
// For more information about permissions, see [Using Bucket Policies and User Policies].
//
// If you include the Filter element in a replication configuration, you must also
// include the DeleteMarkerReplication and Priority elements. The response also
// returns those elements.
//
// For information about GetBucketReplication errors, see [List of replication-related error codes]
//
// The following operations are related to GetBucketReplication :
//
// [PutBucketReplication]
//
// [DeleteBucketReplication]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html
// [List of replication-related error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ReplicationErrorCodeList
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html
