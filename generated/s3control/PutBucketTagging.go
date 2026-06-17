package s3control

// PutBucketTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This action puts tags on an Amazon S3 on Outposts bucket. To put tags on an S3
// bucket, see [PutBucketTagging]in the Amazon S3 API Reference.
//
// Sets the tags for an S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts] in the
// Amazon S3 User Guide.
//
// Use tags to organize your Amazon Web Services bill to reflect your own cost
// structure. To do this, sign up to get your Amazon Web Services account bill with
// tag key values included. Then, to see the cost of combined resources, organize
// your billing information according to resources with the same tag key values.
// For example, you can tag several resources with a specific application name, and
// then organize your billing information to see the total cost of that application
// across several services. For more information, see [Cost allocation and tagging].
//
// Within a bucket, if you add a tag that has the same key as an existing tag, the
// new value overwrites the old value. For more information, see [Using cost allocation in Amazon S3 bucket tags].
//
// To use this action, you must have permissions to perform the
// s3-outposts:PutBucketTagging action. The Outposts bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing access permissions to your Amazon S3 resources].
//
// PutBucketTagging has the following special errors:
//
// - Error code: InvalidTagError
//
// - Description: The tag provided was not a valid tag. This error can occur if
// the tag did not pass input validation. For information about tag restrictions,
// see [User-Defined Tag Restrictions]and [Amazon Web Services-Generated Cost Allocation Tag Restrictions].
//
// - Error code: MalformedXMLError
//
// - Description: The XML provided does not match the schema.
//
// - Error code: OperationAbortedError
//
// - Description: A conflicting conditional action is currently in progress
// against this resource. Try again.
//
// - Error code: InternalError
//
// - Description: The service was unable to apply the provided tag to the bucket.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutBucketTagging :
//
// [GetBucketTagging]
//
// [DeleteBucketTagging]
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [Cost allocation and tagging]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html
// [Using cost allocation in Amazon S3 bucket tags]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html
// [User-Defined Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
// [Managing access permissions to your Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Amazon Web Services-Generated Cost Allocation Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html#API_control_PutBucketTagging_Examples
