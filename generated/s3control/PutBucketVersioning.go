package s3control

// PutBucketVersioning is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation sets the versioning state for S3 on Outposts buckets only. To
// set the versioning state for an S3 bucket, see [PutBucketVersioning]in the Amazon S3 API Reference.
//
// Sets the versioning state for an S3 on Outposts bucket. With S3 Versioning, you
// can save multiple distinct copies of your objects and recover from unintended
// user actions and application failures.
//
// You can set the versioning state to one of the following:
//
// - Enabled - Enables versioning for the objects in the bucket. All objects
// added to the bucket receive a unique version ID.
//
// - Suspended - Suspends versioning for the objects in the bucket. All objects
// added to the bucket receive the version ID null .
//
// If you've never set versioning on your bucket, it has no versioning state. In
// that case, a [GetBucketVersioning]request does not return a versioning state value.
//
// When you enable S3 Versioning, for each object in your bucket, you have a
// current version and zero or more noncurrent versions. You can configure your
// bucket S3 Lifecycle rules to expire noncurrent versions after a specified time
// period. For more information, see [Creating and managing a lifecycle configuration for your S3 on Outposts bucket]in the Amazon S3 User Guide.
//
// If you have an object expiration lifecycle configuration in your non-versioned
// bucket and you want to maintain the same permanent delete behavior when you
// enable versioning, you must add a noncurrent expiration policy. The noncurrent
// expiration lifecycle configuration will manage the deletes of the noncurrent
// object versions in the version-enabled bucket. For more information, see [Versioning]in the
// Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following operations are related to PutBucketVersioning for S3 on Outposts.
//
// [GetBucketVersioning]
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketLifecycleConfiguration]
//
// [Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [PutBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketVersioning.html
// [Creating and managing a lifecycle configuration for your S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsLifecycleManaging.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketVersioning.html#API_control_PutBucketVersioning_Examples
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html
