package s3control

// GetBucketTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action gets an Amazon S3 on Outposts bucket's tags. To get an S3 bucket
// tags, see [GetBucketTagging]in the Amazon S3 API Reference.
//
// Returns the tag set associated with the Outposts bucket. For more information,
// see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the GetBucketTagging
// action. By default, the bucket owner has this permission and can grant this
// permission to others.
//
// GetBucketTagging has the following special error:
//
// - Error code: NoSuchTagSetError
//
// - Description: There is no tag set associated with the bucket.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to GetBucketTagging :
//
// [PutBucketTagging]
//
// [DeleteBucketTagging]
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html#API_control_GetBucketTagging_Examples
