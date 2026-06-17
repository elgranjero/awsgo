package s3control

// PutBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action creates an Amazon S3 on Outposts bucket's replication
// configuration. To create an S3 bucket's replication configuration, see [PutBucketReplication]in the
// Amazon S3 API Reference.
//
// Creates a replication configuration or replaces an existing one. For
// information about S3 replication on Outposts configuration, see [Replicating objects for S3 on Outposts]in the Amazon
// S3 User Guide.
//
// It can take a while to propagate PUT or DELETE requests for a replication
// configuration to all S3 on Outposts systems. Therefore, the replication
// configuration that's returned by a GET request soon after a PUT or DELETE
// request might return a more recent result than what's on the Outpost. If an
// Outpost is offline, the delay in updating the replication configuration on that
// Outpost can be significant.
//
// Specify the replication configuration in the request body. In the replication
// configuration, you provide the following information:
//
// - The name of the destination bucket or buckets where you want S3 on Outposts
// to replicate objects
//
// - The Identity and Access Management (IAM) role that S3 on Outposts can
// assume to replicate objects on your behalf
//
// - Other relevant information, such as replication rules
//
// A replication configuration must include at least one rule and can contain a
// maximum of 100. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source Outposts bucket. To choose additional
// subsets of objects to replicate, add a rule for each subset.
//
// To specify a subset of the objects in the source Outposts bucket to apply a
// replication rule to, add the Filter element as a child of the Rule element. You
// can filter objects based on an object key prefix, one or more object tags, or
// both. When you add the Filter element in the configuration, you must also add
// the following elements: DeleteMarkerReplication , Status , and Priority .
//
// Using PutBucketReplication on Outposts requires that both the source and
// destination buckets must have versioning enabled. For information about enabling
// versioning on a bucket, see [Managing S3 Versioning for your S3 on Outposts bucket].
//
// For information about S3 on Outposts replication failure reasons, see [Replication failure reasons] in the
// Amazon S3 User Guide.
//
// # Handling Replication of Encrypted Objects
//
// Outposts buckets are encrypted at all times. All the objects in the source
// Outposts bucket are encrypted and can be replicated. Also, all the replicas in
// the destination Outposts bucket are encrypted with the same encryption key as
// the objects in the source Outposts bucket.
//
// # Permissions
//
// To create a PutBucketReplication request, you must have
// s3-outposts:PutReplicationConfiguration permissions for the bucket. The Outposts
// bucket owner has this permission by default and can grant it to others. For more
// information about permissions, see [Setting up IAM with S3 on Outposts]and [Managing access to S3 on Outposts buckets].
//
// To perform this operation, the user or role must also have the iam:CreateRole
// and iam:PassRole permissions. For more information, see [Granting a user permissions to pass a role to an Amazon Web Services service].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following operations are related to PutBucketReplication :
//
// [GetBucketReplication]
//
// [DeleteBucketReplication]
//
// [Replicating objects for S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html
// [Replication failure reasons]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html
// [Setting up IAM with S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html
// [Managing access to S3 on Outposts buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html
// [Managing S3 Versioning for your S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsManagingVersioning.html
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html
// [Granting a user permissions to pass a role to an Amazon Web Services service]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html#API_control_PutBucketReplication_Examples
