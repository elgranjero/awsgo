package s3control

// GetBucketLifecycleConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action gets an Amazon S3 on Outposts bucket's lifecycle configuration. To
// get an S3 bucket's lifecycle configuration, see [GetBucketLifecycleConfiguration]in the Amazon S3 API Reference.
//
// Returns the lifecycle configuration information set on the Outposts bucket. For
// more information, see [Using Amazon S3 on Outposts]and for information about lifecycle configuration, see [Object Lifecycle Management]
// in Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3-outposts:GetLifecycleConfiguration action. The Outposts bucket owner has this
// permission, by default. The bucket owner can grant this permission to others.
// For more information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// GetBucketLifecycleConfiguration has the following special error:
//
// - Error code: NoSuchLifecycleConfiguration
//
// - Description: The lifecycle configuration does not exist.
//
// - HTTP Status Code: 404 Not Found
//
// - SOAP Fault Code Prefix: Client
//
// The following actions are related to GetBucketLifecycleConfiguration :
//
// [PutBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [DeleteBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html#API_control_GetBucketLifecycleConfiguration_Examples
