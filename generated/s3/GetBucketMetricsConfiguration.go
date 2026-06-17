package s3

// GetBucketMetricsConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Gets a metrics configuration (specified by the metrics configuration ID) from
// the bucket. Note that this doesn't include the daily storage metrics.
//
// To use this operation, you must have permissions to perform the
// s3:GetMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about CloudWatch request metrics for Amazon S3, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to GetBucketMetricsConfiguration :
//
// [PutBucketMetricsConfiguration]
//
// [DeleteBucketMetricsConfiguration]
//
// [ListBucketMetricsConfigurations]
//
// [Monitoring Metrics with Amazon CloudWatch]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [ListBucketMetricsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html
// [PutBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html
// [DeleteBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
