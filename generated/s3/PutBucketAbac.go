package s3

// PutBucketAbac is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Sets the attribute-based access control (ABAC) property of the general purpose
// bucket. You must have s3:PutBucketABAC permission to perform this action. When
// you enable ABAC, you can use tags for access control on your buckets.
// Additionally, when ABAC is enabled, you must use the [TagResource]and [UntagResource] actions to manage
// tags on your buckets. You can nolonger use the [PutBucketTagging]and [DeleteBucketTagging] actions to tag your bucket.
// For more information, see [Enabling ABAC in general purpose buckets].
//
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [TagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UntagResource.html
// [Enabling ABAC in general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
