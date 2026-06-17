package s3control

// GetBucketReplication is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation gets an Amazon S3 on Outposts bucket's replication
// configuration. To get an S3 bucket's replication configuration, see [GetBucketReplication]in the
// Amazon S3 API Reference.
//
// Returns the replication configuration of an S3 on Outposts bucket. For more
// information about S3 on Outposts, see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide. For
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
// This action requires permissions for the s3-outposts:GetReplicationConfiguration
// action. The Outposts bucket owner has this permission by default and can grant
// it to others. For more information about permissions, see [Setting up IAM with S3 on Outposts]and [Managing access to S3 on Outposts bucket] in the Amazon S3
// User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// If you include the Filter element in a replication configuration, you must also
// include the DeleteMarkerReplication , Status , and Priority elements. The
// response also returns those elements.
//
// For information about S3 on Outposts replication failure reasons, see [Replication failure reasons] in the
// Amazon S3 User Guide.
//
// The following operations are related to GetBucketReplication :
//
// [PutBucketReplication]
//
// [DeleteBucketReplication]
//
// [Replicating objects for S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html
// [Replication failure reasons]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [Setting up IAM with S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html
// [Managing access to S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html#API_control_GetBucketReplication_Examples
