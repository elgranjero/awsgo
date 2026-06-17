package s3control

// PutBucketPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action puts a bucket policy to an Amazon S3 on Outposts bucket. To put a
// policy on an S3 bucket, see [PutBucketPolicy]in the Amazon S3 API Reference.
//
// Applies an Amazon S3 bucket policy to an Outposts bucket. For more information,
// see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the Outposts bucket, the calling identity must have
// the PutBucketPolicy permissions on the specified Outposts bucket and belong to
// the bucket owner's account in order to use this action.
//
// If you don't have PutBucketPolicy permissions, Amazon S3 returns a 403 Access
// Denied error. If you have the correct permissions, but you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 405
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
// The following actions are related to PutBucketPolicy :
//
// [GetBucketPolicy]
//
// [DeleteBucketPolicy]
//
// [PutBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [DeleteBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html
// [GetBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html#API_control_PutBucketPolicy_Examples
