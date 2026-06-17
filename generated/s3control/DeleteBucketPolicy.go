package s3control

// DeleteBucketPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action deletes an Amazon S3 on Outposts bucket policy. To delete an S3
// bucket policy, see [DeleteBucketPolicy]in the Amazon S3 API Reference.
//
// This implementation of the DELETE action uses the policy subresource to delete
// the policy of a specified Amazon S3 on Outposts bucket. If you are using an
// identity other than the root user of the Amazon Web Services account that owns
// the bucket, the calling identity must have the s3-outposts:DeleteBucketPolicy
// permissions on the specified Outposts bucket and belong to the bucket owner's
// account to use this action. For more information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// If you don't have DeleteBucketPolicy permissions, Amazon S3 returns a 403
// Access Denied error. If you have the correct permissions, but you're not using
// an identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
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
// The following actions are related to DeleteBucketPolicy :
//
// [GetBucketPolicy]
//
// [PutBucketPolicy]
//
// [PutBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html
// [DeleteBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [GetBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html#API_control_DeleteBucketPolicy_Examples
