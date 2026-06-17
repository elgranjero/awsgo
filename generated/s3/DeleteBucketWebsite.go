package s3

// DeleteBucketWebsite is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// This action removes the website configuration for a bucket. Amazon S3 returns a
// 200 OK response upon successfully deleting a website configuration on the
// specified bucket. You will get a 200 OK response if the website configuration
// you are trying to delete does not exist on the bucket. Amazon S3 returns a 404
// response if the bucket specified in the request does not exist.
//
// This DELETE action requires the S3:DeleteBucketWebsite permission. By default,
// only the bucket owner can delete the website configuration attached to a bucket.
// However, bucket owners can grant other users permission to delete the website
// configuration by writing a bucket policy granting them the
// S3:DeleteBucketWebsite permission.
//
// For more information about hosting websites, see [Hosting Websites on Amazon S3].
//
// The following operations are related to DeleteBucketWebsite :
//
// [GetBucketWebsite]
//
// [PutBucketWebsite]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketWebsite.html
// [PutBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html
// [Hosting Websites on Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
