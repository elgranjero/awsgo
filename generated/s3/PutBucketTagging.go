package s3

// PutBucketTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Sets the tags for a general purpose bucket if attribute based access control
// (ABAC) is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use the [TagResource]or [UntagResource] operations instead.
//
// Use tags to organize your Amazon Web Services bill to reflect your own cost
// structure. To do this, sign up to get your Amazon Web Services account bill with
// tag key values included. Then, to see the cost of combined resources, organize
// your billing information according to resources with the same tag key values.
// For example, you can tag several resources with a specific application name, and
// then organize your billing information to see the total cost of that application
// across several services. For more information, see [Cost Allocation and Tagging]and [Using Cost Allocation in Amazon S3 Bucket Tags].
//
// When this operation sets the tags for a bucket, it will overwrite any current
// tags the bucket already has. You cannot use this operation to add tags to an
// existing list of tags.
//
// To use this operation, you must have permissions to perform the
// s3:PutBucketTagging action. The bucket owner has this permission by default and
// can grant this permission to others. For more information about permissions, see
// [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// PutBucketTagging has the following special errors. For more Amazon S3 errors
// see, [Error Responses].
//
// - InvalidTag - The tag provided was not a valid tag. This error can occur if
// the tag did not pass input validation. For more information, see [Using Cost Allocation in Amazon S3 Bucket Tags].
//
// - MalformedXML - The XML provided does not match the schema.
//
// - OperationAborted - A conflicting conditional action is currently in progress
// against this resource. Please try again.
//
// - InternalError - The service was unable to apply the provided tag to the
// bucket.
//
// The following operations are related to PutBucketTagging :
//
// [GetBucketTagging]
//
// [DeleteBucketTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Error Responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [Cost Allocation and Tagging]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
// [enable ABAC for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [TagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UntagResource.html
// [Using Cost Allocation in Amazon S3 Bucket Tags]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
