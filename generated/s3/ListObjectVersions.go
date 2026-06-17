package s3

// ListObjectVersions is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns metadata about all versions of the objects in a bucket. You can also
// use request parameters as selection criteria to return metadata about a subset
// of all the object versions.
//
// To use this operation, you must have permission to perform the
// s3:ListBucketVersions action. Be aware of the name difference.
//
// A 200 OK response can contain valid or invalid XML. Make sure to design your
// application to parse the contents of the response and handle it appropriately.
//
// To use this operation, you must have READ access to the bucket.
//
// The following operations are related to ListObjectVersions :
//
// [ListObjectsV2]
//
// [GetObject]
//
// [PutObject]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [ListObjectsV2]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html
