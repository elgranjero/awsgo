package s3control

// PutBucketLifecycleConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action puts a lifecycle configuration to an Amazon S3 on Outposts bucket.
// To put a lifecycle configuration to an S3 bucket, see [PutBucketLifecycleConfiguration]in the Amazon S3 API
// Reference.
//
// Creates a new lifecycle configuration for the S3 on Outposts bucket or replaces
// an existing lifecycle configuration. Outposts buckets only support lifecycle
// configurations that delete/expire objects after a certain period of time and
// abort incomplete multipart uploads.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutBucketLifecycleConfiguration :
//
// [GetBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [DeleteBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html#API_control_PutBucketLifecycleConfiguration_Examples
