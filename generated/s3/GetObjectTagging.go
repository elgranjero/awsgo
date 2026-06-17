package s3

// GetObjectTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the tag-set of an object. You send the GET request against the tagging
// subresource associated with the object.
//
// To use this operation, you must have permission to perform the
// s3:GetObjectTagging action. By default, the GET action returns information about
// current version of an object. For a versioned bucket, you can have multiple
// versions of an object in your bucket. To retrieve tags of any other version, use
// the versionId query parameter. You also need permission for the
// s3:GetObjectVersionTagging action.
//
// By default, the bucket owner has this permission and can grant this permission
// to others.
//
// For information about the Amazon S3 object tagging feature, see [Object Tagging].
//
// The following actions are related to GetObjectTagging :
//
// [DeleteObjectTagging]
//
// [GetObjectAttributes]
//
// [PutObjectTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html
// [PutObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Object Tagging]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html
