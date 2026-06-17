package s3

// PutBucketIntelligentTieringConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Puts a S3 Intelligent-Tiering configuration to the specified bucket. You can
// have up to 1,000 S3 Intelligent-Tiering configurations per bucket.
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
// Operations related to PutBucketIntelligentTieringConfiguration include:
//
// [DeleteBucketIntelligentTieringConfiguration]
//
// [GetBucketIntelligentTieringConfiguration]
//
// [ListBucketIntelligentTieringConfigurations]
//
// You only need S3 Intelligent-Tiering enabled on a bucket if you want to
// automatically move objects stored in the S3 Intelligent-Tiering storage class to
// the Archive Access or Deep Archive Access tier.
//
// PutBucketIntelligentTieringConfiguration has the following special errors:
//
// HTTP 400 Bad Request Error  Code: InvalidArgument
//
// Cause: Invalid Argument
//
// HTTP 400 Bad Request Error  Code: TooManyConfigurations
//
// Cause: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// HTTP 403 Forbidden Error  Cause: You are not the owner of the specified bucket,
// or you do not have the s3:PutIntelligentTieringConfiguration bucket permission
// to set the configuration on the bucket.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListBucketIntelligentTieringConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketIntelligentTieringConfigurations.html
// [GetBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html
// [Storage class for automatically optimizing frequently and infrequently accessed objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access
// [DeleteBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html
