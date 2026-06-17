package s3

// PutObjectTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Sets the supplied tag-set to an object that already exists in a bucket. A tag
// is a key-value pair. For more information, see [Object Tagging].
//
// You can associate tags with an object by sending a PUT request against the
// tagging subresource that is associated with the object. You can retrieve tags by
// sending a GET request. For more information, see [GetObjectTagging].
//
// For tagging-related restrictions related to characters and encodings, see [Tag Restrictions].
// Note that Amazon S3 limits the maximum number of tags to 10 tags per object.
//
// To use this operation, you must have permission to perform the
// s3:PutObjectTagging action. By default, the bucket owner has this permission and
// can grant this permission to others.
//
// To put tags of any other version, use the versionId query parameter. You also
// need permission for the s3:PutObjectVersionTagging action.
//
// PutObjectTagging has the following special errors. For more Amazon S3 errors
// see, [Error Responses].
//
// - InvalidTag - The tag provided was not a valid tag. This error can occur if
// the tag did not pass input validation. For more information, see [Object Tagging].
//
// - MalformedXML - The XML provided does not match the schema.
//
// - OperationAborted - A conflicting conditional action is currently in progress
// against this resource. Please try again.
//
// - InternalError - The service was unable to apply the provided tag to the
// object.
//
// The following operations are related to PutObjectTagging :
//
// [GetObjectTagging]
//
// [DeleteObjectTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Error Responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
// [DeleteObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html
// [Object Tagging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-tagging.html
// [Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
// [GetObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html
