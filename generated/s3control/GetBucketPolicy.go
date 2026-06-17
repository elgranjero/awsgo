package s3control

// GetBucketPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action gets a bucket policy for an Amazon S3 on Outposts bucket. To get a
// policy for an S3 bucket, see [GetBucketPolicy]in the Amazon S3 API Reference.
//
// Returns the policy of a specified Outposts bucket. For more information, see [Using Amazon S3 on Outposts]
// in the Amazon S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the bucket, the calling identity must have the
// GetBucketPolicy permissions on the specified bucket and belong to the bucket
// owner's account in order to use this action.
//
// Only users from Outposts bucket owner account with the right permissions can
// perform actions on an Outposts bucket. If you don't have
// s3-outposts:GetBucketPolicy permissions or you're not using an identity that
// belongs to the bucket owner's account, Amazon S3 returns a 403 Access Denied
// error.
//
// As a security precaution, the root user of the Amazon Web Services account that
// owns a bucket can always use this action, even if the policy explicitly denies
// the root user the ability to perform this action.
//
// For more information about bucket policies, see [Using Bucket Policies and User Policies].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to GetBucketPolicy :
//
// [GetObject]
//
// [PutBucketPolicy]
//
// [DeleteBucketPolicy]
//
// [PutBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [DeleteBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html
// [GetBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html#API_control_GetBucketPolicy_Examples
