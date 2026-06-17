package s3

// ListBucketMetricsConfigurations is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Lists the metrics configurations for the bucket. The metrics configurations are
// only for the request metrics of the bucket and do not provide information on
// daily storage metrics. You can have up to 1,000 configurations per bucket.
//
// This action supports list pagination and does not return more than 100
// configurations at a time. Always check the IsTruncated element in the response.
// If there are no more configurations to list, IsTruncated is set to false. If
// there are more configurations to list, IsTruncated is set to true, and there is
// a value in NextContinuationToken . You use the NextContinuationToken value to
// continue the pagination of the list by passing the value in continuation-token
// in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the
// s3:GetMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For more information about metrics configurations and CloudWatch request
// metrics, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to ListBucketMetricsConfigurations :
//
// [PutBucketMetricsConfiguration]
//
// [GetBucketMetricsConfiguration]
//
// [DeleteBucketMetricsConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [GetBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html
// [PutBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html
// [DeleteBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
