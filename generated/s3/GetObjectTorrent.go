package s3

// GetObjectTorrent is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns torrent files from a bucket. BitTorrent can save you bandwidth when
// you're distributing large files.
//
// You can get torrent only for objects that are less than 5 GB in size, and that
// are not encrypted using server-side encryption with a customer-provided
// encryption key.
//
// To use GET, you must have READ access to the object.
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// The following action is related to GetObjectTorrent :
//
// [GetObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
