package s3control

// GetBucket is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Gets an Amazon S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts] in the Amazon
// S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the Outposts bucket, the calling identity must have
// the s3-outposts:GetBucket permissions on the specified Outposts bucket and
// belong to the Outposts bucket owner's account in order to use this action. Only
// users from Outposts bucket owner account with the right permissions can perform
// actions on an Outposts bucket.
//
// If you don't have s3-outposts:GetBucket permissions or you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 403
// Access Denied error.
//
// The following actions are related to GetBucket for Amazon S3 on Outposts:
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// [PutObject]
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html#API_control_GetBucket_Examples
