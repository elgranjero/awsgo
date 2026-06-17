package s3

// DeleteObjectTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Removes the entire tag set from the specified object. For more information
// about managing object tags, see [Object Tagging].
//
// To use this operation, you must have permission to perform the
// s3:DeleteObjectTagging action.
//
// To delete tags of a specific object version, add the versionId query parameter
// in the request. You will need permission for the s3:DeleteObjectVersionTagging
// action.
//
// The following operations are related to DeleteObjectTagging :
//
// [PutObjectTagging]
//
// [GetObjectTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html
// [Object Tagging]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html
// [GetObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html
