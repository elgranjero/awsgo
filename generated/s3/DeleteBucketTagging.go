package s3

// DeleteBucketTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Deletes tags from the general purpose bucket if attribute based access control
// (ABAC) is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use [UntagResource]instead.
//
// if ABAC is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use [UntagResource]instead.
//
// To use this operation, you must have permission to perform the
// s3:PutBucketTagging action. By default, the bucket owner has this permission and
// can grant this permission to others.
//
// The following operations are related to DeleteBucketTagging :
//
// [GetBucketTagging]
//
// [PutBucketTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [enable ABAC for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UntagResource.html
