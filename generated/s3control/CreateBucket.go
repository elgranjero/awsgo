package s3control

// CreateBucket is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action creates an Amazon S3 on Outposts bucket. To create an S3 bucket,
// see [Create Bucket]in the Amazon S3 API Reference.
//
// Creates a new Outposts bucket. By creating the bucket, you become the bucket
// owner. To create an Outposts bucket, you must have S3 on Outposts. For more
// information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// Not every string is an acceptable bucket name. For information on bucket naming
// restrictions, see [Working with Amazon S3 Buckets].
//
// S3 on Outposts buckets support:
//
// - Tags
//
// - LifecycleConfigurations for deleting expired objects
//
// For a complete list of restrictions and Amazon S3 feature limitations on S3 on
// Outposts, see [Amazon S3 on Outposts Restrictions and Limitations].
//
// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
// on Outposts endpoint hostname prefix and x-amz-outpost-id in your API request,
// see the [Examples]section.
//
// The following actions are related to CreateBucket for Amazon S3 on Outposts:
//
// [PutObject]
//
// [GetBucket]
//
// [DeleteBucket]
//
// [CreateAccessPoint]
//
// [PutAccessPointPolicy]
//
// [GetBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [Working with Amazon S3 Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html#bucketnamingrules
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html
// [Create Bucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [PutAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html#API_control_CreateBucket_Examples
// [Amazon S3 on Outposts Restrictions and Limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OnOutpostsRestrictionsLimitations.html
