package s3

// GetBucketAnalyticsConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// This implementation of the GET action returns an analytics configuration
// (identified by the analytics configuration ID) from the bucket.
//
// To use this operation, you must have permissions to perform the
// s3:GetAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources] in the Amazon S3 User Guide.
//
// For information about Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class Analysis] in the Amazon S3 User
// Guide.
//
// The following operations are related to GetBucketAnalyticsConfiguration :
//
// [DeleteBucketAnalyticsConfiguration]
//
// [ListBucketAnalyticsConfigurations]
//
// [PutBucketAnalyticsConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Analytics – Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html
// [DeleteBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [ListBucketAnalyticsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketAnalyticsConfigurations.html
// [PutBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
