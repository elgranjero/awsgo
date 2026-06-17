package s3

// PutBucketMetricsConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Sets a metrics configuration (specified by the metrics configuration ID) for
// the bucket. You can have up to 1,000 metrics configurations per bucket. If
// you're updating an existing metrics configuration, note that this is a full
// replacement of the existing metrics configuration. If you don't include the
// elements you want to keep, they are erased.
//
// To use this operation, you must have permissions to perform the
// s3:PutMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about CloudWatch request metrics for Amazon S3, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to PutBucketMetricsConfiguration :
//
// [DeleteBucketMetricsConfiguration]
//
// [GetBucketMetricsConfiguration]
//
// [ListBucketMetricsConfigurations]
//
// PutBucketMetricsConfiguration has the following special error:
//
// - Error code: TooManyConfigurations
//
// - Description: You are attempting to create a new configuration but have
// already reached the 1,000-configuration limit.
//
// - HTTP Status Code: HTTP 400 Bad Request
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [GetBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html
// [ListBucketMetricsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html
// [DeleteBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
