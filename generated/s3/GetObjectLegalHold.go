package s3

// GetObjectLegalHold is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Gets an object's current legal hold status. For more information, see [Locking Objects].
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// The following action is related to GetObjectLegalHold :
//
// [GetObjectAttributes]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
