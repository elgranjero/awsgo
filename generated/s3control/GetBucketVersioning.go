package s3control

// GetBucketVersioning is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation returns the versioning state for S3 on Outposts buckets only. To
// return the versioning state for an S3 bucket, see [GetBucketVersioning]in the Amazon S3 API
// Reference.
//
// Returns the versioning state for an S3 on Outposts bucket. With S3 Versioning,
// you can save multiple distinct copies of your objects and recover from
// unintended user actions and application failures.
//
// If you've never set versioning on your bucket, it has no versioning state. In
// that case, the GetBucketVersioning request does not return a versioning state
// value.
//
// For more information about versioning, see [Versioning] in the Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following operations are related to GetBucketVersioning for S3 on Outposts.
//
// [PutBucketVersioning]
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketLifecycleConfiguration]
//
// [Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [PutBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketVersioning.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html#API_control_GetBucketVersioning_Examples
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html
