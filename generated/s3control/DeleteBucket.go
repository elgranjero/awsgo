package s3control

// DeleteBucket is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action deletes an Amazon S3 on Outposts bucket. To delete an S3 bucket,
// see [DeleteBucket]in the Amazon S3 API Reference.
//
// Deletes the Amazon S3 on Outposts bucket. All objects (including all object
// versions and delete markers) in the bucket must be deleted before the bucket
// itself can be deleted. For more information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// # Related Resources
//
// [CreateBucket]
//
// [GetBucket]
//
// [DeleteObject]
//
// [GetBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html#API_control_DeleteBucket_Examples
