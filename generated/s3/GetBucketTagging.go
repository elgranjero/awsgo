package s3

// GetBucketTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the tag set associated with the general purpose bucket.
//
// if ABAC is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use [ListTagsForResource]instead.
//
// To use this operation, you must have permission to perform the
// s3:GetBucketTagging action. By default, the bucket owner has this permission and
// can grant this permission to others.
//
// GetBucketTagging has the following special error:
//
// - Error code: NoSuchTagSet
//
// - Description: There is no tag set associated with the bucket.
//
// The following operations are related to GetBucketTagging :
//
// [PutBucketTagging]
//
// [DeleteBucketTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [enable ABAC for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListTagsForResource.html
