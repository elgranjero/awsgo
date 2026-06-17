package s3

// PutBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Creates a replication configuration or replaces an existing one. For more
// information, see [Replication]in the Amazon S3 User Guide.
//
// Specify the replication configuration in the request body. In the replication
// configuration, you provide the name of the destination bucket or buckets where
// you want Amazon S3 to replicate objects, the IAM role that Amazon S3 can assume
// to replicate objects on your behalf, and other relevant information. You can
// invoke this request for a specific Amazon Web Services Region by using the [aws:RequestedRegion]
// aws:RequestedRegion condition key.
//
// A replication configuration must include at least one rule, and can contain a
// maximum of 1,000. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source bucket. To choose additional subsets of
// objects to replicate, add a rule for each subset.
//
// To specify a subset of the objects in the source bucket to apply a replication
// rule to, add the Filter element as a child of the Rule element. You can filter
// objects based on an object key prefix, one or more object tags, or both. When
// you add the Filter element in the configuration, you must also add the following
// elements: DeleteMarkerReplication , Status , and Priority .
//
// If you are using an earlier version of the replication configuration, Amazon S3
// handles replication of delete markers differently. For more information, see [Backward Compatibility].
//
// For information about enabling versioning on a bucket, see [Using Versioning].
//
// Handling Replication of Encrypted Objects By default, Amazon S3 doesn't
// replicate objects that are stored at rest using server-side encryption with KMS
// keys. To replicate Amazon Web Services KMS-encrypted objects, add the following:
// SourceSelectionCriteria , SseKmsEncryptedObjects , Status ,
// EncryptionConfiguration , and ReplicaKmsKeyID . For information about
// replication configuration, see [Replicating Objects Created with SSE Using KMS keys].
//
// For information on PutBucketReplication errors, see [List of replication-related error codes]
//
// Permissions To create a PutBucketReplication request, you must have
// s3:PutReplicationConfiguration permissions for the bucket.
//
// By default, a resource owner, in this case the Amazon Web Services account that
// created the bucket, can perform this operation. The resource owner can also
// grant others permissions to perform the operation. For more information about
// permissions, see [Specifying Permissions in a Policy]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// To perform this operation, the user or role performing the action must have the [iam:PassRole]
// permission.
//
// The following operations are related to PutBucketReplication :
//
// [GetBucketReplication]
//
// [DeleteBucketReplication]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [iam:PassRole]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [aws:RequestedRegion]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-requestedregion
// [Replicating Objects Created with SSE Using KMS keys]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-config-for-kms-objects.html
// [Using Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html
// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html
// [List of replication-related error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ReplicationErrorCodeList
// [Backward Compatibility]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
