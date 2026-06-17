package s3

// ListBucketIntelligentTieringConfigurations is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Lists the S3 Intelligent-Tiering configuration from the specified bucket.
//
// The S3 Intelligent-Tiering storage class is designed to optimize storage costs
// by automatically moving data to the most cost-effective storage access tier,
// without performance impact or operational overhead. S3 Intelligent-Tiering
// delivers automatic cost savings in three low latency and high throughput access
// tiers. To get the lowest storage cost on data that can be accessed in minutes to
// hours, you can choose to activate additional archiving capabilities.
//
// The S3 Intelligent-Tiering storage class is the ideal storage class for data
// with unknown, changing, or unpredictable access patterns, independent of object
// size or retention period. If the size of an object is less than 128 KB, it is
// not monitored and not eligible for auto-tiering. Smaller objects can be stored,
// but they are always charged at the Frequent Access tier rates in the S3
// Intelligent-Tiering storage class.
//
// For more information, see [Storage class for automatically optimizing frequently and infrequently accessed objects].
//
// Operations related to ListBucketIntelligentTieringConfigurations include:
//
// [DeleteBucketIntelligentTieringConfiguration]
//
// [PutBucketIntelligentTieringConfiguration]
//
// [GetBucketIntelligentTieringConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html
// [PutBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html
// [Storage class for automatically optimizing frequently and infrequently accessed objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access
// [DeleteBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html
